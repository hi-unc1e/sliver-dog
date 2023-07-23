package portfwd

/*
	Sliver Implant Framework
	Copyright (C) 2021  Bishop Fox

	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.

	This program is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU General Public License for more details.

	You should have received a copy of the GNU General Public License
	along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

import (
	"github.com/bishopfox/sliver/client/console"
	"github.com/bishopfox/sliver/client/core"
	"github.com/spf13/cobra"
)

// PortfwdRmCmd - Remove an existing tunneled port forward.
func PortfwdRmCmd(cmd *cobra.Command, con *console.SliverConsoleClient, args []string) {
	portfwdID, _ := cmd.Flags().GetInt("id")
	if portfwdID < 1 {
		con.PrintErrorf("Must specify a valid portfwd id\n")
		return
	}
	found := core.Portfwds.Remove(portfwdID)
	if !found {
		con.PrintErrorf("No portfwd with id %d\n", portfwdID)
	} else {
		con.PrintInfof("Removed portfwd\n")
	}
}
