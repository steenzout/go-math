package operator_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

// PackageTestSuite package test suite.
type PackageTestSuite struct {
	suite.Suite
}

// TestPackage run package test suites.
func TestPackage(t *testing.T) {
	suite.Run(t, new(GenericTestSuite))
	suite.Run(t, new(ComparisonTestSuite))
	suite.Run(t, new(LogicalTestSuite))
	suite.Run(t, new(MathTestSuite))
}
