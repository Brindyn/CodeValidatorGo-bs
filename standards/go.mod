module github.com/Brindyn/CodeValidatorGo-bs/tree/master/standards

go 1.17

replace github.com/Brindyn/CodeValidatorGo-bs/tree/master/test1 => ../test1

replace github.com/Brindyn/CodeValidatorGo-bs/tree/master/test2 => ../test2

replace github.com/Brindyn/CodeValidatorGo-bs/tree/master/test3 => ../test3

replace github.com/Brindyn/CodeValidatorGo-bs/tree/master/test4 => ../test4

require (
	github.com/Brindyn/CodeValidatorGo-bs/tree/master/test1 v0.0.0-00010101000000-000000000000
	github.com/Brindyn/CodeValidatorGo-bs/tree/master/test2 v0.0.0-00010101000000-000000000000
	github.com/Brindyn/CodeValidatorGo-bs/tree/master/test3 v0.0.0-00010101000000-000000000000
	github.com/Brindyn/CodeValidatorGo-bs/tree/master/test4 v0.0.0-00010101000000-000000000000
)
