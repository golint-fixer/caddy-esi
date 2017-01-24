// Copyright 2016-2017, Cyrill @ Schumacher.fm and the CaddyESI Contributors
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may not
// use this file except in compliance with the License. You may obtain a copy of
// the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations under
// the License.

package main

import (
	"fmt"

	"github.com/vdobler/ht/ht"
)

func init() {
	RegisterTest(headerCmd())
}

func headerCmd() (t *ht.Test) {

	req := makeRequestGET("page_cart_tiny.html")
	req.Header.Set("X-Esi-Cmd", "purge")

	t = &ht.Test{
		Name:        fmt.Sprint("Header Command Purge 1"),
		Description: `Sends a special header to purge the ESI tag cache`,
		Request:     req,
		Checks: makeChecklist200(
			&ht.Header{
				Header: `X-Esi-Cmd`,
				Condition: ht.Condition{
					// Whenever we change the number of tests and the cached
					// entries grows more we must change here the number of
					// items in the cache before purging it.
					Equals: `purge-ok-6`,
				},
			},
			&ht.Body{
				Contains: "demo-store.shop/autumn-pullie.html",
				Count:    2,
			},
		),
	}
	return
}
