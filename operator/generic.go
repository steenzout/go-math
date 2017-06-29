package operator

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

const (
	// PrecedenceLower operator has lower precedence then the other.
	PrecedenceLower = -1
	// PrecedenceSame operator has same precedence as the other.
	PrecedenceSame = 0
	// PredecendeHigher operator has higher precedence then the other.
	PredecendeHigher = 1
)

// Generic generic operator.
type Generic struct {
	// precedence order by which operators are evaluated.
	// Operators with higher precedence are evaluated first.
	precedence uint8
	// associativity True for operators with left-to-right associativity; False otherwise.
	associativity bool
}

// New returns a new Generic struct pointer.
func New(p uint8, assoc bool) *Generic {
	op := Generic{
		precedence:    p,
		associativity: assoc,
	}
	return &op
}

// Compare compares operators precedence.
// Returns 1 if this has higher precedence than the other,
// 0 if it has the same precedence and
// -1 if lower.
func (g Generic) Compare(o Operator) int {
	other := o.Precedence()

	if g.precedence == other {
		return PrecedenceSame
	} else if g.precedence > other {
		return PredecendeHigher
	} else {
		return PrecedenceLower
	}
}

// HasLowerPrecedenceThan returns true this operator has lower precedence than the other operator; false otherwise.
func (g Generic) HasLowerPrecedenceThan(o *Generic) bool {
	return g.precedence < o.precedence
}

// HasSamePrecedenceThan returns true this operator has the same precedence than the other operator; false otherwise.
func (g Generic) HasSamePrecedenceThan(o *Generic) bool {
	return g.precedence == o.precedence
}

// HasHigherPrecedenceThan returns true this operator has higher precedence than the other operator; false otherwise.
func (g Generic) HasHigherPrecedenceThan(o *Generic) bool {
	return g.precedence > o.precedence
}

// Precedence returns operator precedence.
func (g Generic) Precedence() uint8 {
	return g.precedence
}

// IsLeftAssociative returns true if the operator is left-associative; false otherwise.
func (g Generic) IsLeftAssociative() bool {
	return !g.associativity
}

// IsRightAssociate returns true if the operator right-associative; false otherwise.
func (g Generic) IsRightAssociate() bool {
	return g.associativity
}
