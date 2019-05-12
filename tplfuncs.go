package main

import (
	"strings"

	"github.com/Kretech/xgo/word"
)

var funcs map[string]interface{}

func init() {
	if funcs == nil {
		funcs = make(map[string]interface{}, 32)
	}

	addStrings(funcs)
}

func addStrings(fs map[string]interface{}) {
	mergeFrom(fs, map[string]interface{}{
		"ToUpper":    strings.ToUpper,
		"ToLower":    strings.ToLower,
		"LowerFirst": word.LowerFirst,
		"UpperFirst": word.UpperFirst,
	})
}

func mergeFrom(to, from map[string]interface{}) {
	for key, value := range from {
		to[key] = value
	}
}
