package dummyserver

import (
	"github.com/pkg/errors"
)

func Perform(input map[string]interface{}) (interface{}, error) {
	numbers, ok := input["numbers"]
	if !ok {
		return nil, errors.New("'numbers' parameter was not given")
	}
	numbersList, ok := numbers.([]interface{})
	if !ok {
		return nil, errors.New("'numbers' is not a list")
	}
	inputFloats := make([]float64, len(numbersList))
	var err error
	for i, arg := range numbersList {
		inputFloats[i] = arg.(float64)
		if err != nil {
			return nil, errors.Wrap(err, "converting argument to float64")
		}
	}

	var sum float64 = 0
	for _, number := range inputFloats {
		sum += number
	}

	return sum, nil
}
