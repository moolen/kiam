// Copyright 2017 uSwitch
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
package main

import (
	"github.com/vishvananda/netlink"
)

func newConntrackFilter() (uint, error) {
	return netlink.ConntrackDeleteFilter(netlink.ConntrackTable, netlink.FAMILY_V4, dropMetadataAPI{})
}

type dropMetadataAPI struct{}

func (c dropMetadataAPI) MatchConntrackFlow(flow *netlink.ConntrackFlow) bool {
	if flow.FamilyType == netlink.FAMILY_V4 && flow.Forward.DstIP.String() == "169.254.169.254" {
		return true
	}
	return false
}
