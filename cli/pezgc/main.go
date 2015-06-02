package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/codegangsta/cli"

	cf "github.com/pivotalservices/pezdispenser/cloudfoundryclient"
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

func main1() {
	baseURI := os.Getenv("CF_DOMAIN")
	user := os.Getenv("CF_USER")
	pass := os.Getenv("CF_PASS")
	loginURI := fmt.Sprintf("https://%s.%s", "login", baseURI)
	apiURI := fmt.Sprintf("https://%s.%s", "api", baseURI)
	heritageClient := &heritage{
		Client:   ccclient.New(loginURI, user, pass, new(http.Client)),
		ccTarget: apiURI,
	}
	heritageClient.Login()
	cfclient := cf.NewCloudFoundryClient(heritageClient, new(logger))
	cfclient.QueryAPIInfo()
	u, _ := cfclient.QueryUsers(1, 1, "id", "")
	users, _ := cfclient.QueryUsers(1, u.TotalResults, "userName,meta", url.QueryEscape("origin eq 'uaa' and userName eq 'calabrese.john@gmail.com'"))

	for _, v := range users.Resources {
		fmt.Printf("Created: %s Modified: %s User: %s\n", v.Meta["created"], v.Meta["lastModified"], v.UserName)
	}
	fmt.Println("Users Found: ", users.TotalResults)
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
