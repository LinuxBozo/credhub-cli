package actions_test

import (
	"bytes"
	"io/ioutil"
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pivotal-cf/cm-cli/actions"
	"github.com/pivotal-cf/cm-cli/client"
	"github.com/pivotal-cf/cm-cli/client/clientfakes"
	"github.com/pivotal-cf/cm-cli/config"
)

var _ = Describe("Token", func() {
	var (
		subject    actions.Version
		httpClient clientfakes.FakeHttpClient
	)

	BeforeEach(func() {
		config := config.Config{AuthURL: "example.com"}
		subject = actions.NewToken(&httpClient, config)
	})

	Describe("Authorization", func() {
		It("returns the token from the authorization server", func() {
			request := client.NewTokenRequest("example.com", "userName", "password")

			responseObj := http.Response{
				StatusCode: 200,
				Body: ioutil.NopCloser(bytes.NewBufferString(`{
					"access_token":"2YotnFZFEjr1zCsicMWpAA",
					"token_type":"bearer",
					"expires_in":3600}`)),
			}

			httpClient.DoStub = func(req *http.Request) (resp *http.Response, err error) {
				Expect(req).To(Equal(request))

				return &responseObj, nil
			}

			token, _ := subject.GetToken("userName", "password")
			Expect(token.AccessToken).To(Equal("2YotnFZFEjr1zCsicMWpAA"))
		})


	})
})