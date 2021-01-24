package storage_plugin_test

import (
	"fmt"
	"os"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	storage_plugin "github.com/nitric-dev/membrane/plugins/dev/storage"
)

type MockStorageDriver struct {
	storage_plugin.StorageDriver
	ensureDirExistsError error
	writeFileError       error
	directories          []string
	storedItems          map[string][]byte
}

func (m *MockStorageDriver) EnsureDirExists(dir string) error {
	if m.ensureDirExistsError != nil {
		return m.ensureDirExistsError
	}

	if m.directories == nil {
		m.directories = make([]string, 0)
	}

	m.directories = append(m.directories, dir)

	return nil
}

func (m *MockStorageDriver) WriteFile(file string, contents []byte, fileMode os.FileMode) error {
	if m.writeFileError != nil {
		return m.writeFileError
	}

	// Capture for later eval
	if m.storedItems == nil {
		m.storedItems = make(map[string][]byte)
	}

	pathParts := strings.Split(file, "/")

	directory := strings.Join(pathParts[:len(pathParts)-1], "/") + "/"
	for _, dir := range m.directories {
		if dir == directory {
			m.storedItems[file] = contents
			return nil
		}
	}

	return fmt.Errorf("Cannot create file as directory does not exist")
}

var _ = Describe("Storage", func() {
	Context("Put", func() {
		// Test Put methods...
		Context("Given the provided storage driver is functioning without error", func() {
			workingDriver := &MockStorageDriver{}
			mockStoragePlugin, _ := storage_plugin.NewWithStorageDriver(workingDriver)
			It("Should store the provided byte array", func() {
				err := mockStoragePlugin.Put("test", "test", []byte("Test"))
				By("Not returning an error")
				Expect(err).To(BeNil())

				Expect(workingDriver.storedItems["/nitric/buckets/test/test"]).To(BeEquivalentTo([]byte("Test")))
			})
		})

		Context("Given the storage driver cannot create directories", func() {
			faultyDriver := &MockStorageDriver{
				ensureDirExistsError: fmt.Errorf("error creating directory"),
			}
			mockStoragePlugin, _ := storage_plugin.NewWithStorageDriver(faultyDriver)
			It("Should store the provided byte array", func() {

				err := mockStoragePlugin.Put("test", "test", []byte("Test"))
				By("By returning an error")
				Expect(err).ToNot(BeNil())

				By("Not creating any directories")
				Expect(faultyDriver.directories).To(BeNil())
			})
		})

		Context("Given the storage driver cannot create files", func() {
			faultyDriver := &MockStorageDriver{
				writeFileError: fmt.Errorf("error creating file"),
			}
			mockStoragePlugin, _ := storage_plugin.NewWithStorageDriver(faultyDriver)
			It("Should store the provided byte array", func() {
				err := mockStoragePlugin.Put("test", "test", []byte("Test"))
				By("By returning an error")
				Expect(err).ToNot(BeNil())

				By("By creating the requested directory")
				Expect(faultyDriver.directories).ToNot(BeNil())

				By("Not storing any items")
				Expect(faultyDriver.storedItems).To(BeNil())
			})
		})
	})
})