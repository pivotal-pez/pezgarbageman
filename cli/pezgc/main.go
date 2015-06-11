package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/codegangsta/cli"

	cf "github.com/pivotalservices/pezdispenser/cloudfoundryclient"
	gm "github.com/pivotalservices/pezgarbageman"
	"github.com/xchapter7x/cloudcontroller-client"
)

type logger struct{}

func (s *logger) Println(l ...interface{}) {
	if os.Getenv("LOG_DEBUG") == "true" {
		fmt.Println(l...)
	}
}

type heritage struct {
	*ccclient.Client
	ccTarget string
}

func (s *heritage) CCTarget() string {
	return s.ccTarget
}

func main() {
	app := NewApp()
	app.Run(os.Args)
}

// NewApp creates a new cli app
func NewApp() *cli.App {

	app := cli.NewApp()
	app.Name = "pezgc"
	app.Usage = "allows for auditing and purging assets in pez. a system wide garbage man service"
	app.Commands = append(app.Commands, []cli.Command{
		showCli,
		purgeCli,
	}...)

	return app
}

func getList(usertype, username, cfDomain, cfUser, cfPass string) {
	baseURI := cfDomain
	user := cfUser
	pass := cfPass
	loginURI := fmt.Sprintf("https://%s.%s", "login", baseURI)
	apiURI := fmt.Sprintf("https://%s.%s", "api", baseURI)
	heritageClient := &heritage{
		Client:   ccclient.New(loginURI, user, pass, new(http.Client)),
		ccTarget: apiURI,
	}
	heritageClient.Login()
	cfclient := cf.NewCloudFoundryClient(heritageClient, new(logger))
	userSearch := new(gm.UserSearch).Init(cfclient)
	users, _ := userSearch.List(usertype, username)

	for _, v := range users.Resources {
		fmt.Printf("Type:%s User: %s\n", v.Origin, v.UserName)
	}
	fmt.Println("Users Found: ", users.TotalResults)
}
