package general

import (
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/devstream-io/devstream/internal/pkg/configmanager"
)

var _ = Describe("validate func", func() {
	var (
		options                  configmanager.RawOptions
		cloneURL, configLocation string
	)
	BeforeEach(func() {
		cloneURL = "git@github.com/root/test-exmaple.git"
		configLocation = "workflows"
		options = configmanager.RawOptions{
			"scm": map[string]interface{}{
				"cloneURL": cloneURL,
			},
			"action": map[string]interface{}{
				"configLocation": configLocation,
			},
			"projectRepo": map[string]interface{}{
				"repoType": "github",
				"owner":    "test",
				"repo":     "gg",
			},
		}
	})
	When("scm token is not setted", func() {
		It("should return err", func() {
			_, err := validate(options)
			Expect(err).Error().Should(HaveOccurred())
		})
	})
	When("scm repo is gitlab", func() {
		BeforeEach(func() {
			options["scm"] = map[string]interface{}{
				"cloneURL": "http://exmaple.gitlab.com",
			}
		})
		It("should return error", func() {
			_, err := validate(options)
			Expect(err).Error().Should(HaveOccurred())
		})
	})
	When("all valid", func() {
		BeforeEach(func() {
			os.Setenv("GITHUB_TOKEN", "test")
		})
		It("should not raise error", func() {
			_, err := validate(options)
			Expect(err).Error().ShouldNot(HaveOccurred())
		})
		AfterEach(func() {
			os.Unsetenv("GITHUB_TOKEN")
		})
	})
})

var _ = Describe("setDefault func", func() {
	var (
		options                  configmanager.RawOptions
		cloneURL, configLocation string
	)
	BeforeEach(func() {
		cloneURL = "http://github.com/root/test-exmaple.git"
		configLocation = "workflows"
		options = configmanager.RawOptions{
			"scm": map[string]interface{}{
				"cloneURL": cloneURL,
			},
			"action": map[string]interface{}{
				"configLocation": configLocation,
			},
		}
	})
	BeforeEach(func() {
		os.Setenv("GITHUB_TOKEN", "test")
	})
	It("should work normal", func() {
		opts, err := setDefault(options)
		Expect(err).ShouldNot(HaveOccurred())
		projectRepo, ok := opts["projectRepo"]
		Expect(ok).Should(BeTrue())
		Expect(projectRepo).ShouldNot(BeNil())
		ciConfig, ok := opts["ci"]
		Expect(ok).Should(BeTrue())
		Expect(ciConfig).ShouldNot(BeNil())
	})
	AfterEach(func() {
		os.Unsetenv("GITHUB_TOKEN")
	})
})
