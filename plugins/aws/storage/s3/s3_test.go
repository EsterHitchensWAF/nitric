package s3_plugin_test

import (
	"github.com/nitric-dev/membrane/plugins/aws/mocks"
	s3_plugin "github.com/nitric-dev/membrane/plugins/aws/storage/s3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("S3", func() {
	When("Put", func() {
		When("Given the S3 backend is available", func() {
			When("Creating an object in an existing bucket", func() {
				testPayload := []byte("Test")
				storage := make(map[string]map[string][]byte)
				mockStorageClient := mocks.NewStorageClient([]*mocks.MockBucket{
					&mocks.MockBucket{
						Name: "my-bucket",
						Tags: map[string]string{
							"x-nitric-name": "my-bucket",
						},
					},
				}, &storage)

				storagePlugin, _ := s3_plugin.NewWithClient(mockStorageClient)
				It("Should successfully store the object", func() {
					err := storagePlugin.Put("my-bucket", "test-item", testPayload)
					By("Not returning an error")
					Expect(err).ShouldNot(HaveOccurred())

					By("Storing the item")
					Expect(storage["my-bucket"]["test-item"]).To(BeEquivalentTo(testPayload))
				})
			})

			When("Creating an object in a non-existant bucket", func() {
				storage := make(map[string]map[string][]byte)
				mockStorageClient := mocks.NewStorageClient([]*mocks.MockBucket{}, &storage)
				storagePlugin, _ := s3_plugin.NewWithClient(mockStorageClient)
				It("Should fail to store the item", func() {
					err := storagePlugin.Put("my-bucket", "test-item", []byte("Test"))
					By("Returning an error")
					Expect(err).Should(HaveOccurred())
				})
			})
		})
	})
})