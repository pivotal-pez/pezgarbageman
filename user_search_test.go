package garbageman_test

import (
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	cf "github.com/pivotalservices/pezdispenser/cloudfoundryclient"
	. "github.com/pivotalservices/pezgarbageman"
)

var _ = Describe("UserSearch", func() {
	Context("callint .List with in-valid arguments", func() {
		var (
			myUserSearch     *UserSearch
			controlListError = errors.New("my mock error")
		)

		BeforeEach(func() {
			mockClient := &mockCFClient{
				Error: controlListError,
			}
			myUserSearch = new(UserSearch)
			myUserSearch.Init(mockClient)
		})

		It("should return an error", func() {
			_, err := myUserSearch.List("wrong", "invalid")
			Ω(err).ShouldNot(BeNil())
			Ω(err).Should(Equal(controlListError))
		})
	})

	Context("callint .List with valid arguments", func() {
		var (
			myUserSearch    *UserSearch
			controlUserList = cf.UserAPIResponse{
				Resources: []cf.UserResource{
					cf.UserResource{
						ID: "mytestid",
					},
				},
			}
		)

		BeforeEach(func() {
			mockClient := &mockCFClient{
				UserList: controlUserList,
			}
			myUserSearch = new(UserSearch)
			myUserSearch.Init(mockClient)
		})

		It("shouldreturn the expected userlist object and nil error", func() {
			users, err := myUserSearch.List("uaa", "testu")
			Ω(err).Should(BeNil())
			Ω(users).Should(Equal(controlUserList))
		})
	})

	Context("calling .Init with valid arguments", func() {
		var (
			myUserSearch *UserSearch
		)

		BeforeEach(func() {
			myUserSearch = new(UserSearch)
		})

		It("Should return a fully initialized UserSearch Object", func() {
			controlClient := myUserSearch.Client
			mockClient := new(mockCFClient)
			myUserSearch.Init(mockClient)
			Ω(myUserSearch.Client).ShouldNot(BeNil())
			Ω(myUserSearch.Client).ShouldNot(Equal(controlClient))
		})
	})

	Context("calling .BuildQuery with valid arguments", func() {
		var (
			myUserSearch *UserSearch
		)

		BeforeEach(func() {
			myUserSearch = new(UserSearch)
		})

		It("should combine multiple non empty arguments", func() {
			controlQuery := "origin+eq+%27okta%27+and+userName+co+%27testu%27"
			query := myUserSearch.BuildQuery("okta", "testu")
			Ω(query).ShouldNot(BeNil())
			Ω(query).Should(Equal(controlQuery))
		})

		It("should drop empty 1st argument", func() {
			controlQuery := "userName+co+%27testu%27"
			query := myUserSearch.BuildQuery("", "testu")
			Ω(query).ShouldNot(BeNil())
			Ω(query).Should(Equal(controlQuery))
		})

		It("should drop empty 2nd argument", func() {
			controlQuery := "origin+eq+%27okta%27"
			query := myUserSearch.BuildQuery("okta", "")
			Ω(query).ShouldNot(BeNil())
			Ω(query).Should(Equal(controlQuery))
		})
	})
})
