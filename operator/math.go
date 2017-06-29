package operator

// Addition + math operator.
var Addition = Generic{
	precedence:    11,
	associativity: true,
}

// Division / math operator.
var Division = Generic{
	precedence:    12,
	associativity: true,
}

// Modulus % math operator.
var Modulus = Generic{
	precedence:    12,
	associativity: true,
}

// Multiplication * math operator.
var Multiplication = Generic{
	precedence:    12,
	associativity: true,
}

// Remainder alias for Modulus.
var Remainder = Modulus

// Subtraction - math operator.
var Subtraction = Generic{
	precedence:    11,
	associativity: true,
}
