package main

func isBuiltInFunc(word string) bool {
	for key := range inBuiltFuncs {
		if key == word {
			return true
		}
	}
	return false
}

var inBuiltFuncs map[string][]string = map[string][]string{
	"out": {"callable"}, // Print
}
