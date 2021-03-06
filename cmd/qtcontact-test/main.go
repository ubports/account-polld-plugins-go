/*
 Copyright 2014 Canonical Ltd.

 This program is free software: you can redistribute it and/or modify it
 under the terms of the GNU General Public License version 3, as published
 by the Free Software Foundation.

 This program is distributed in the hope that it will be useful, but
 WITHOUT ANY WARRANTY; without even the implied warranties of
 MERCHANTABILITY, SATISFACTORY QUALITY, or FITNESS FOR A PARTICULAR
 PURPOSE.  See the GNU General Public License for more details.

 You should have received a copy of the GNU General Public License along
 with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package main

import (
	"fmt"
	"os"

	"launchpad.net/account-polld/qtcontact"
)

func main() {
	qtcontact.MainLoopStart()

	if len(os.Args) != 2 {
		fmt.Println("usage:", os.Args[0], "[email address]")
		os.Exit(1)
	}

	path := qtcontact.GetAvatar(os.Args[1])
	fmt.Println("Avatar found:", path)
}
