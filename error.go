package main

import "errors"

var (
	ErrJsonObject = errors.New("input is not JSON object or array")
)