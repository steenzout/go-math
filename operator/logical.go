package operator

// And && logical operator.
var And = Generic{
	precedence:    4,
	associativity: true,
}

// Not ! logical operator.
var Not = Generic{
	precedence:    13,
	associativity: false,
}

// Or || logical operator.
var Or = Generic{
	precedence:    3,
	associativity: true,
}
