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
	cli.StringFlag{
		Name:   strings.Join(cfdomainFlag, ", "),
		Value:  "",
		Usage:  cfdomainDescr,
		EnvVar: cfdomainEnv,
	},
	cli.StringFlag{
		Name:   strings.Join(cfuserFlag, ", "),
		Value:  "",
		Usage:  cfuserDescr,
		EnvVar: cfuserEnv,
	},
	cli.StringFlag{
		Name:   strings.Join(cfpassFlag, ", "),
		Value:  "",
		Usage:  cfpassDescr,
		EnvVar: cfpassEnv,
	},
}
