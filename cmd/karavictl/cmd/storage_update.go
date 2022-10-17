// Copyright © 2021-2022 Dell Inc., or its subsidiaries. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"karavi-authorization/pb"
	"net/url"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"sigs.k8s.io/yaml"
)

// NewStorageUpdateCmd creates a new update command
func NewStorageUpdateCmd() *cobra.Command {
	storageUpdateCmd := &cobra.Command{
		Use:   "update",
		Short: "Update a registered storage system.",
		Long:  `Updates a registered storage system.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Get the storage systems and update it in place?
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			errAndExit := func(err error) {
				fmt.Fprintf(cmd.ErrOrStderr(), "error: %+v\n", err)
				osExit(1)
			}

			// Convenience functions for ignoring errors whilst
			// getting flag values.
			flagStringValue := func(v string, err error) string {
				if err != nil {
					errAndExit(err)
				}
				return v
			}
			flagBoolValue := func(v bool, err error) bool {
				if err != nil {
					errAndExit(err)
				}
				return v
			}
			verifyInput := func(v string) string {
				inputText := flagStringValue(cmd.Flags().GetString(v))
				if strings.TrimSpace(inputText) == "" {
					errAndExit(fmt.Errorf("no input provided: %s", v))
				}
				return inputText
			}

			// Gather the inputs

			addr := flagStringValue(cmd.Flags().GetString("addr"))
			insecure := flagBoolValue(cmd.Flags().GetBool("insecure"))

			input := input{
				Type:          verifyInput("type"),
				Endpoint:      verifyInput("endpoint"),
				SystemID:      verifyInput("system-id"),
				User:          verifyInput("user"),
				Password:      flagStringValue(cmd.Flags().GetString("password")),
				ArrayInsecure: flagBoolValue(cmd.Flags().GetBool("array-insecure")),
			}

			// Parse the URL and prepare for a password prompt.
			urlWithUser, err := url.Parse(input.Endpoint)
			if err != nil {
				errAndExit(err)
			}

			urlWithUser.Scheme = "https"
			urlWithUser.User = url.User(input.User)

			// If the password was not provided...
			prompt := fmt.Sprintf("Enter password for %v: ", urlWithUser)
			// If the password was not provided...
			if pf := cmd.Flags().Lookup("password"); !pf.Changed {
				// Get password from stdin
				readPassword(cmd.ErrOrStderr(), prompt, &input.Password)
			}

			// Sanitize the endpoint
			epURL, err := url.Parse(input.Endpoint)
			if err != nil {
				errAndExit(err)
			}
			epURL.Scheme = "https"

			if addr != "" {
				err := doStorageUpdateRequest(ctx, addr, input, insecure)
				if err != nil {
					errAndExit(err)
				}
			} else {
				k3sCmd := execCommandContext(ctx, K3sPath, "kubectl", "get",
					"--namespace=karavi",
					"--output=json",
					"secret/karavi-storage-secret")

				b, err := k3sCmd.Output()
				if err != nil {
					errAndExit(err)
				}

				base64Systems := struct {
					Data map[string]string
				}{}
				if err := json.Unmarshal(b, &base64Systems); err != nil {
					errAndExit(err)
				}
				decodedSystems, err := base64.StdEncoding.DecodeString(base64Systems.Data["storage-systems.yaml"])
				if err != nil {
					errAndExit(err)
				}

				var listData map[string]Storage
				if err := yaml.Unmarshal(decodedSystems, &listData); err != nil {
					errAndExit(err)
				}
				if listData == nil || listData["storage"] == nil {
					listData = make(map[string]Storage)
					listData["storage"] = make(Storage)
				}
				var storage = listData["storage"]

				var didUpdate bool
				for k := range storage {
					if k != input.Type {
						continue
					}
					_, ok := storage[k][input.SystemID]
					if !ok {
						continue
					}

					storage[k][input.SystemID] = System{
						User:     input.User,
						Password: input.Password,
						Endpoint: input.Endpoint,
						Insecure: input.ArrayInsecure,
					}
					didUpdate = true
					break
				}
				if !didUpdate {
					errAndExit(fmt.Errorf("no matching storage systems to update"))
				}

				listData["storage"] = storage
				b, err = yaml.Marshal(&listData)
				if err != nil {
					errAndExit(err)
				}

				tmpFile, err := ioutil.TempFile("", "karavi")
				if err != nil {
					errAndExit(err)
				}
				defer func() {
					if err := tmpFile.Close(); err != nil {
						fmt.Fprintf(os.Stderr, "error: %+v\n", err)
					}
					if err := os.Remove(tmpFile.Name()); err != nil {
						fmt.Fprintf(os.Stderr, "error: %+v\n", err)
					}
				}()
				_, err = tmpFile.WriteString(string(b))
				if err != nil {
					errAndExit(err)
				}

				crtCmd := execCommandContext(ctx, K3sPath, "kubectl", "create",
					"--namespace=karavi",
					"secret", "generic", "karavi-storage-secret",
					fmt.Sprintf("--from-file=storage-systems.yaml=%s", tmpFile.Name()),
					"--output=yaml",
					"--dry-run=client")
				appCmd := execCommandContext(ctx, K3sPath, "kubectl", "apply", "-f", "-")

				if err := pipeCommands(crtCmd, appCmd); err != nil {
					errAndExit(err)
				}
			}
		},
	}
	storageUpdateCmd.Flags().StringP("type", "t", "", "Type of storage system")
	err := storageUpdateCmd.MarkFlagRequired("type")
	if err != nil {
		reportErrorAndExit(JSONOutput, storageUpdateCmd.ErrOrStderr(), err)
	}
	storageUpdateCmd.Flags().StringP("endpoint", "e", "", "Endpoint of REST API gateway")
	err = storageUpdateCmd.MarkFlagRequired("endpoint")
	if err != nil {
		reportErrorAndExit(JSONOutput, storageUpdateCmd.ErrOrStderr(), err)
	}
	storageUpdateCmd.Flags().StringP("system-id", "s", "", "System identifier")
	err = storageUpdateCmd.MarkFlagRequired("system-id")
	if err != nil {
		reportErrorAndExit(JSONOutput, storageUpdateCmd.ErrOrStderr(), err)
	}
	storageUpdateCmd.Flags().StringP("user", "u", "", "Username")
	err = storageUpdateCmd.MarkFlagRequired("user")
	if err != nil {
		reportErrorAndExit(JSONOutput, storageUpdateCmd.ErrOrStderr(), err)
	}
	storageUpdateCmd.Flags().StringP("password", "p", "", "Specify password, or omit to use stdin")
	storageUpdateCmd.Flags().BoolP("array-insecure", "a", false, "Array insecure skip verify")

	return storageUpdateCmd
}

func doStorageUpdateRequest(ctx context.Context, addr string, system input, grpcInsecure bool) error {
	client, conn, err := CreateStorageServiceClient(addr, grpcInsecure)
	if err != nil {
		return err
	}
	defer conn.Close()

	req := &pb.StorageUpdateRequest{
		StorageType: system.Type,
		Endpoint:    system.Endpoint,
		SystemId:    system.SystemID,
		UserName:    system.User,
		Password:    system.Password,
		Insecure:    system.ArrayInsecure,
	}

	_, err = client.Update(ctx, req)
	if err != nil {
		return err
	}

	return nil
}
