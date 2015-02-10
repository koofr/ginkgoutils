// +build !darwin

package ginkgoutils

func (sc *SuiteConfig) NewTest() (dir string, num int) {
	return sc.newTest()
}
