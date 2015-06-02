package main

import (
	"strings"

	"github.com/codegangsta/cli"
)

var vacuumFlags = []cli.Flag{
	cli.StringFlag{
		Name:   strings.Join(usertypeFlag, ", "),
		Value:  "",
		Usage:  usertypeDescr,
		EnvVar: "",
	},
	cli.StringFlag{
		Name:   strings.Join(usernameFlag, ", "),
		Value:  "",
		Usage:  usernameDescr,
		EnvVar: "",
	},
}
