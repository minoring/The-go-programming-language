// Each type T has an underlying type: If T is one of the predeclared boolean, numeric, or string types, or a type literal, 
// the corresponding underlying type is T itself. Otherwise, T's underlying type is the underlying type of the type to which 
// T refers in its type declaration.

type (
	A1 = string
	A2 = A1
)

type (
	B1 string
	B2 B1
	B3 []B1
	B4 B3
)

//The underlying type of string, A1, A2, B1, and B2 is string. The underlying type of []B1, B3, and B4 is []B1.
