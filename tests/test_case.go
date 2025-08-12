package tests

import (
	"github.com/goravel/framework/testing"

	"pixel/bootstrap"
)

func init() {
	bootstrap.Boot()
}

type TestCase struct {
	testing.TestCase
}
