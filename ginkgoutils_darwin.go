// +build darwin

package ginkgoutils

import (
	"path/filepath"
)

func (sc *SuiteConfig) NewTest() (dir string, num int) {
	dir, num = sc.newTest()
	dir = filepath.Join("/private", dir)
	return
}
