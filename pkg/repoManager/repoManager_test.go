package repoManager

import (
	"os"
	"path"

	"github.com/multiGit/pkg/helpers"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

const baseDir = "tmp/test-multi-git"

var repoList = []string{}

var _ = Describe("Repo manager tests", Ordered, func() {
	var err error

	removeAll := func() {
		err = os.RemoveAll(baseDir)
		Expect(err).Should(BeNil())
	}

	BeforeAll(func() {
		removeAll()

		err = helpers.CreateDir(baseDir, "dir-1", true)
		Expect(err).Should(BeNil())
		repoList = []string{"dir-1"}
	})

	AfterAll(removeAll)

	Context("Test for success cases", func() {
		It("Should get git repo list successfully", func() {
			repos, err := getGitRepos(baseDir)
			Expect(err).Should(BeNil())

			Expect(repos).Should(HaveLen(1))
			Expect(repos[0] == path.Join(baseDir, repoList[0])).Should(BeTrue())
		})

		It("Should successfully create a new branch", func() {
			outputs, err := ExecGitCommand(baseDir, "checkout -b test-branch", true)
			Expect(err).Should(BeNil())

			for _, output := range outputs {
				Expect(output).Should(Equal("Switched to a new branch 'test-branch'\n"))
			}
		})
	})

	Context("Test for failure cases", func() {
		It("Should fail with invalid base dir", func() {
			_, err := getGitRepos("/no-such-dir")

			Expect(err).ShouldNot(BeNil())
		})
	})
})
