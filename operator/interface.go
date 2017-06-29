package operator

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
