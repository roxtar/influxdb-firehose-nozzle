package uaatokenfetcher_test

import (
	"github.com/evoila/influxdb-firehose-nozzle/uaatokenfetcher"

	"github.com/evoila/influxdb-firehose-nozzle/testhelpers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("UaaTokenFetcher", func() {
	var tokenFetcher *uaatokenfetcher.UAATokenFetcher
	var fakeUAA *testhelpers.FakeUAA
	var fakeToken string

	BeforeEach(func() {
		fakeUAA = testhelpers.NewFakeUAA("bearer", "123456789")
		fakeToken = fakeUAA.AuthToken()
		fakeUAA.Start()

		tokenFetcher = &uaatokenfetcher.UAATokenFetcher{
			UaaUrl: fakeUAA.URL(),
		}
	})

	It("fetches a token from the UAA", func() {
		receivedAuthToken := tokenFetcher.FetchAuthToken()
		Expect(fakeUAA.Requested()).To(BeTrue())
		Expect(receivedAuthToken).To(Equal(fakeToken))
	})
})
