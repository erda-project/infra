// Copyright (c) 2021 Terminus, Inc.
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

package cardlist

import (
	"github.com/erda-project/erda-infra/providers/component-protocol/cptype"
)

type (
	// OpCardListGoto .
	OpCardListGoto struct {
		cptype.Operation
		ServerData OpCardListGotoData `json:"serverData,omitempty"`
	}

	// OpCardListGotoData .
	OpCardListGotoData struct {
		JumpOut bool
		Target  string
		Query   cptype.ExtraMap `json:"query,omitempty"`
		Params  cptype.ExtraMap `json:"params,omitempty"`
	}
)

type (
	// OpCardListIconGoto .
	OpCardListIconGoto struct {
		cptype.Operation
		Icon       string
		Tip        string
		ServerData OpCardListGotoData `json:"serverData,omitempty"`
	}
)

// OpKey .
func (o OpCardListGoto) OpKey() cptype.OperationKey { return "clickGoto" }

// OpKey .
func (o OpCardListIconGoto) OpKey() cptype.OperationKey { return "clickGoto" }
