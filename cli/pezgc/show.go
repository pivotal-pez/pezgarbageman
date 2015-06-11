package main

import (
	"fmt"

	"github.com/codegangsta/cli"
)

var showCli = cli.Command{
	Name:        show_full_name,
	ShortName:   show_short_name,
	Usage:       show_usage,
	Description: show_descr,
	Flags:       vacuumFlags,
	Action: func(c *cli.Context) {
		var (
			err      error
			userType = c.String(usertypeFlag[0])
			userName = c.String(usernameFlag[0])
			cfDomain = c.String(cfdomainFlag[0])
			cfUser   = c.String(cfuserFlag[0])
			cfPass   = c.String(cfpassFlag[0])
		)

		if userType != "" || userName != "" {
			getList(userType, userName, cfDomain, cfUser, cfPass)

		} else {
			cli.ShowCommandHelp(c, show_full_name)
		}

		if err != nil {
			fmt.Println(err)

		} else {
			fmt.Println(show_full_name, " completed successfully.")
		}
	},
}
