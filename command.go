package main

import (
	"errors"
	"fmt"
)

type Arguments map[string]string

func ParseArguments(arguments []string) (Arguments, error) {
	container, err := mapArguments(arguments)
	if err != nil {
		return container, err
	}

	err = validateArguments(container)
	if err != nil {
		return container, err
	}

	return container, nil
}

func mapArguments(arguments []string) (Arguments, error) {
	container := Arguments{}
	for i := 0; i < len(arguments); i += 2 {
		if arguments[i][:2] != "--" {
			return nil, fmt.Errorf("invalid argument %s", arguments[i])
		}

		key := arguments[i][2:]
		container[key] = arguments[i+1]
	}

	return container, nil
}

func validateArguments(arguments Arguments) error {
	_, ok := arguments["subject"]
	if !ok {
		return errors.New("subject is required")
	}

	_, ok = arguments["receivers"]
	if !ok {
		return errors.New("receivers is required")
	}

	_, ok = arguments["template"]
	if !ok {
		return errors.New("template is required")
	}

	return nil
}
