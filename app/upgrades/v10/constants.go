// Copyright 2022 Gurux Foundation
// This file is part of the Gurux Network packages.
//
// Gurux is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The Gurux packages are distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the Gurux packages. If not, see https://github.com/gurux/gurux/blob/main/LICENSE

package v10

const (
	// UpgradeName is the shared upgrade plan name for mainnet
	UpgradeName = "v10.0.0"
	// UpgradeInfo defines the binaries that will be used for the upgrade
	UpgradeInfo = `'{"binaries":{"darwin/arm64":"https://github.com/gurux/gurux/releases/download/v10.0.0/gurux_10.0.0_Darwin_arm64.tar.gz","darwin/amd64":"https://github.com/gurux/gurux/releases/download/v10.0.0/gurux_10.0.0_Darwin_amd64.tar.gz","linux/arm64":"https://github.com/gurux/gurux/releases/download/v10.0.0/gurux_10.0.0_Linux_arm64.tar.gz","linux/amd64":"https://github.com/gurux/gurux/releases/download/v10.0.0/gurux_10.0.0_Linux_amd64.tar.gz","windows/x86_64":"https://github.com/gurux/gurux/releases/download/v10.0.0/gurux_10.0.0_Windows_x86_64.zip"}}'`
)
