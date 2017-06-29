package operator

// Different != comparison operator.
var Different = Generic{
	precedence:    8,
	associativity: true,
}

// Equal == comparison operator.
var Equal = Generic{
	precedence:    8,
	associativity: true,
}

// GreaterThan > comparison operator.
var GreaterThan = Generic{
	precedence:    9,
	associativity: true,
}

// GreaterThanEqual >= comparison operator.
var GreaterThanEqual = Generic{
	precedence:    9,
	associativity: true,
}

// LessThan < comparison operator.
var LessThan = Generic{
	precedence:    9,
	associativity: true,
}

// LessThanEqual <= comparison operator.
var LessThanEqual = Generic{
	precedence:    9,
	associativity: true,
}
