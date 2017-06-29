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

// Operator operator interface.
type Operator interface {

	// Compare compares operators precedence.
	// Returns 1 if this has higher precedence than the other,
	// 0 if it has the same precedence and
	// -1 if lower.
	Compare(o Operator) int

	// HasLowerPrecedenceThan returns true this operator has lower precedence than the other operator; false otherwise.
	HasLowerPrecedenceThan(o *Generic) bool

	// HasSamePrecedenceThan returns true this operator has the same precedence than the other operator; false otherwise.
	HasSamePrecedenceThan(o *Generic) bool

	// HasHigherPrecedenceThan returns true this operator has higher precedence than the other operator; false otherwise.
	HasHigherPrecedenceThan(o *Generic) bool

	// Precedence returns operator precedence.
	Precedence() uint8

	// IsLeftAssociative returns true if the operator is left-associative; false otherwise.
	IsLeftAssociative() bool

	// IsRightAssociate returns true if the operator right-associative; false otherwise.
	IsRightAssociate() bool
}
