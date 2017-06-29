package operator_test

//
// Copyright 2017 Pedro Salgado
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

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
