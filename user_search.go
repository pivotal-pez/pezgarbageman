package garbageman

import (
	"fmt"
	"net/url"
	"strings"

	cf "github.com/pivotalservices/pezdispenser/cloudfoundryclient"
)

type UserSearch struct {
	Client cf.CloudFoundryClient
}

func (s *UserSearch) Init(client cf.CloudFoundryClient) *UserSearch {
	s.Client = client
	s.Client.QueryAPIInfo()
	return s
}

func (s *UserSearch) BuildQuery(usertype, username string) (query string) {
	l := []string{}

	if usertype != "" {
		l = append(l, fmt.Sprintf("origin eq '%s'", usertype))
	}

	if username != "" {
		l = append(l, fmt.Sprintf("userName co '%s'", username))
	}
	query = url.QueryEscape(strings.Join(l, " and "))
	return
}

func (s *UserSearch) List(usertype, username string) (users cf.UserAPIResponse, err error) {
	var u cf.UserAPIResponse

	if u, err = s.Client.QueryUsers(1, 1, "id", ""); err == nil {
		query := s.BuildQuery(usertype, username)
		users, err = s.Client.QueryUsers(1, u.TotalResults, "userName,meta", query)
	}
	return
}
