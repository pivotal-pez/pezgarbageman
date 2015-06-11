package main

import (
	"fmt"

	"github.com/codegangsta/cli"
)

var purgeCli = cli.Command{
	Name:        purge_full_name,
	ShortName:   purge_short_name,
	Usage:       purge_usage,
	Description: purge_descr,
	Flags:       vacuumFlags,
	Action: func(c *cli.Context) {
		var (
			err      error
			userType = c.String(usertypeFlag[0])
			userName = c.String(usernameFlag[0])
		)

		if userType != "" || userName != "" {
			fmt.Println(userType, userName)

		} else {
			cli.ShowCommandHelp(c, purge_full_name)
		}

		if err != nil {
			fmt.Println(err)

		} else {
			fmt.Println(purge_full_name, " completed successfully.")
		}
	},
}
