package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"

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
	users, _ := cfclient.QueryUsers(1, u.TotalResults, "userName,meta", url.QueryEscape("origin eq 'uaa'"))

	for _, v := range users.Resources {
		fmt.Printf("Created: %s Modified: %s User: %s\n", v.Meta["created"], v.Meta["lastModified"], v.UserName)
	}
	fmt.Println("UAA Users Found: ", users.TotalResults)
}
