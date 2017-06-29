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
