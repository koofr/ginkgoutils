package ginkgoutils

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
)

type SuiteConfig struct {
	name          string
	testRoot      string
	autocleanRoot bool
	suiteFailed   bool
	testNum       int
	testNumLock   sync.Mutex
}

func NewSuiteConfig(name string) *SuiteConfig {
	return &SuiteConfig{
		name:          name,
		autocleanRoot: true,
	}
}

func (sc *SuiteConfig) SetupSuite() {
	tmpDir := filepath.Join(os.TempDir(), filepath.FromSlash(filepath.Dir(sc.name)))
	err := os.MkdirAll(tmpDir, os.ModePerm|os.ModePerm)
	gomega.Expect(err).NotTo(gomega.HaveOccurred())
	dir, err := ioutil.TempDir(tmpDir, filepath.Base(sc.name))
	gomega.Expect(err).NotTo(gomega.HaveOccurred())
	err = os.MkdirAll(dir, os.ModeDir|os.ModePerm)
	gomega.Expect(err).NotTo(gomega.HaveOccurred())
	sc.testRoot = dir
}

func (sc *SuiteConfig) CleanupSuite() {
	isTemp := strings.HasPrefix(sc.testRoot, os.TempDir())
	if !sc.suiteFailed && sc.autocleanRoot && isTemp {
		os.RemoveAll(sc.testRoot)
	}
}

func (sc *SuiteConfig) newTest() (dir string, num int) {
	sc.testNumLock.Lock()
	defer sc.testNumLock.Unlock()
	num = sc.testNum
	dir = filepath.Join(sc.testRoot, fmt.Sprintf("test%d", num))
	os.Mkdir(dir, os.ModeDir|os.ModePerm)
	sc.testNum++
	return
}

func (sc *SuiteConfig) Fail(msg string, callerSkip ...int) {
	sc.suiteFailed = true
	ginkgo.Fail(msg, callerSkip...)
}
