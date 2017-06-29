package operator_test

import (
	"github.com/steenzout/go-math/operator"
	"github.com/stretchr/testify/suite"
)

// GenericTestSuite package test suite.
type GenericTestSuite struct {
	suite.Suite
}

func (s GenericTestSuite) TestAssociativity() {
	r := operator.New(1, true)
	l := operator.New(1, false)

	s.True(r.IsRightAssociate())
	s.False(r.IsLeftAssociative())

	s.False(l.IsRightAssociate())
	s.True(l.IsLeftAssociative())
}

func (s GenericTestSuite) TestCompare() {
	p1 := operator.New(1, true)
	p2 := operator.New(2, true)

	s.Equal(operator.PrecedenceLower, p1.Compare(p2))
	s.Equal(operator.PrecedenceSame, p1.Compare(p1))
	s.Equal(operator.PredecendeHigher, p2.Compare(p1))
}

func (s GenericTestSuite) TestHasPrecedence() {
	p1 := operator.New(1, true)
	p2 := operator.New(2, true)

	s.True(p1.HasLowerPrecedenceThan(p2))
	s.False(p1.HasHigherPrecedenceThan(p2))
	s.True(p1.HasSamePrecedenceThan(p1))
	s.True(p2.HasHigherPrecedenceThan(p1))
	s.False(p2.HasLowerPrecedenceThan(p1))
}

func (s GenericTestSuite) TestPrecedence() {
	p := uint8(1)
	p1 := operator.New(p, true)

	s.Equal(p, p1.Precedence())
}

func (s GenericTestSuite) TestImplementsInterface() {
	var _ operator.Operator = (*operator.Generic)(nil)
}
