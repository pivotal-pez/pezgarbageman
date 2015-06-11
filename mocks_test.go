package garbageman_test

import (
	"net/http"

	cf "github.com/pivotalservices/pezdispenser/cloudfoundryclient"
)

type mockCFClient struct {
	UserList cf.UserAPIResponse
	Error    error
}

func (s *mockCFClient) QueryAPIInfo() (a *cf.CloudFoundryAPIInfo, b error) {
	return
}

func (s *mockCFClient) QueryUserGUID(username string) (a string, b error) {
	return
}

func (s *mockCFClient) AddRole(rolePathPrefix string, targetGUID string, roleType string, userGUID string) (a error) {
	return
}

func (s *mockCFClient) AddOrg(orgName string) (orgGUID string, err error) {
	return
}

func (s *mockCFClient) AddSpace(spaceName string, orgGUID string) (spaceGUID string, err error) {
	return
}

func (s *mockCFClient) AddUser(username string) (a error) {
	return
}

func (s *mockCFClient) RemoveOrg(orgGUID string) (err error) {
	return
}

func (s *mockCFClient) QueryUsers(int, int, string, string) (userList cf.UserAPIResponse, err error) {
	userList = s.UserList
	err = s.Error
	return
}

func (s *mockCFClient) Query(verb string, domain string, path string, args interface{}) (response *http.Response) {
	return
}
