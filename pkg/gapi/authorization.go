// Copyright 2022 Buf Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gapi

import (
	"fmt"
	"net/http"
	"strings"

	auth "github.com/darwishdev/obd_api/pkg/auth"
	"github.com/rs/zerolog/log"
)

const (
	authorizationHeader = "authorization"
	authorizationBearer = "bearer"
)

func (server *Server) authorizeUser(header http.Header) (*auth.Payload, error) {
	authHeader := header.Get("Authorization")
	if authHeader == "" {
		return nil, fmt.Errorf("missing metadata")
	}

	fields := strings.Fields(authHeader)
	if len(fields) < 2 {
		return nil, fmt.Errorf("invalid authorization header format")
	}

	log.Debug().Interface("fields", fields).Msg("fields")
	// // authHeader := values[0]
	// fields := strings.Fields(authHeader)
	// if len(fields) < 2 {
	// 	return nil, fmt.Errorf("invalid authorization header format")
	// }

	authType := strings.ToLower(fields[0])
	if authType != authorizationBearer {
		return nil, fmt.Errorf("unsupported authorization type: %s", authType)
	}

	accessToken := fields[1]
	payload, err := server.tokenMaker.VerifyToken(accessToken)
	if err != nil {
		return nil, fmt.Errorf("invalid access token: %s", err)
	}

	return payload, nil
}
