// Copyright © 2021 Dell Inc., or its subsidiaries. All Rights Reserved.
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

package powerflex

import (
	"sync"
	"time"

	"context"

	"github.com/dell/goscaleio"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/trace"
)

type PowerFlexTokenGetter struct {
	Config       Config
	sem          chan struct{}
	mu           sync.Mutex // projects currentToken
	currentToken string
}

type Config struct {
	PowerFlexClient      *goscaleio.Client
	TokenRefreshInterval time.Duration
	ConfigConnect        *goscaleio.ConfigConnect
	Logger               *logrus.Entry
}

func NewTokenGetter(c Config) *PowerFlexTokenGetter {
	return &PowerFlexTokenGetter{
		Config: c,
		sem:    make(chan struct{}, 1),
	}
}

func (tg *PowerFlexTokenGetter) Start(ctx context.Context) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}
	// Update the token one time on startup, then update on timer interval after that
	tg.mu.Lock()
	tg.currentToken = ""
	tg.mu.Unlock()
	tg.updateTokenFromPowerFlex()

	timer := time.NewTimer(tg.Config.TokenRefreshInterval)
	defer timer.Stop()
	for {
		select {
		case <-timer.C:
			tg.updateTokenFromPowerFlex()
			timer.Reset(tg.Config.TokenRefreshInterval)
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func (tg *PowerFlexTokenGetter) GetToken(ctx context.Context) (string, error) {
	ctx, span := trace.SpanFromContext(ctx).Tracer().Start(ctx, "GetToken")
	defer span.End()

	select {
	case tg.sem <- struct{}{}:
	case <-ctx.Done():
		return "", ctx.Err()
	}
	defer func() { <-tg.sem }()
	return tg.getToken(), nil
}

func (tg *PowerFlexTokenGetter) getToken() string {
	tg.mu.Lock()
	defer tg.mu.Unlock()
	return tg.currentToken
}

func (tg *PowerFlexTokenGetter) updateTokenFromPowerFlex() {
	tg.sem <- struct{}{}
	defer func() {
		<-tg.sem
	}()

	if _, err := tg.Config.PowerFlexClient.Authenticate(tg.Config.ConfigConnect); err != nil {
		tg.Config.Logger.Errorf("PowerFlex Auth error: %+v", err)
	}
	tg.mu.Lock()
	tg.currentToken = tg.Config.PowerFlexClient.GetToken()
	tg.mu.Unlock()
}
