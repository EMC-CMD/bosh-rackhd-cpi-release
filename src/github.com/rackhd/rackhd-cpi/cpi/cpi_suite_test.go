package cpi_test

import (
	"io/ioutil"

	log "github.com/Sirupsen/logrus"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestCpi(t *testing.T) {
	// where did my logs go
	// disable logging
	log.SetOutput(ioutil.Discard)

	RegisterFailHandler(Fail)
	RunSpecs(t, "CPI Suite")
}
