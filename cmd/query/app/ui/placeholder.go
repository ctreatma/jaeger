// Copyright (c) 2018 The Jaeger Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build !ui
// +build !ui

package ui

import (
	"embed"
	"net/http"

	"github.com/jaegertracing/jaeger/pkg/httpfs"
)

//go:embed placeholder/index.html
var assetsFS embed.FS

// StaticFiles provides http filesystem with static files for UI
var StaticFiles = httpfs.PrefixedFS("placeholder", http.FS(assetsFS))
