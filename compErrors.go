package main

import "errors"

func createError(source string, subSource string, message string) error {
	return errors.New(
		source + " Error:\n" +
			"Found in: " + subSource + "\n" +
			message,
	)
}
