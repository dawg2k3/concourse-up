package iaas_test

import (
	. "github.com/EngineerBetter/concourse-up/iaas"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Client#FindLongestMatchingHostedZone", func() {
	Context("When the hosted zone exists", func() {
		It("Returns the hosted zone details", func() {
			awsClient, err := New("AWS", "eu-west-1")
			Expect(err).To(Succeed())
			zoneName, zoneID, err := (awsClient).FindLongestMatchingHostedZone("integration-test.concourse-up.engineerbetter.com")
			Expect(err).ToNot(HaveOccurred())
			Expect(zoneName).To(Equal("concourse-up.engineerbetter.com"))
			Expect(zoneID).To(Equal("Z2NEMKRYH9QASG"))
		})
	})

	Context("When the hosted zone does not exist", func() {
		It("Returns a meaningful error", func() {
			awsClient, err := New("AWS", "eu-west-1")
			Expect(err).To(Succeed())
			_, _, err = (awsClient).FindLongestMatchingHostedZone("abc.google.com")
			Expect(err).To(MatchError("No matching hosted zone found for domain abc.google.com"))
		})
	})
})
