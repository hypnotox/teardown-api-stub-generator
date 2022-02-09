package src

import (
	"errors"
	"fmt"
)

type TealWriter struct {
	overrideVariableType map[string]string
	vectorTypeName       string
	quaternionTypeName   string
	transformTypeName    string
}

func NewTealWriter() *TealWriter {
	vectorTypeName := "VectorType"
	quaternionTypeName := "QuaternionType"
	transformTypeName := "TransformType"

	return &TealWriter{
		overrideVariableType: map[string]string{
			"number<integer>": "integer",
			"table<number>":   "{number}",
			"Vector":          vectorTypeName,
			"Quaternion":      quaternionTypeName,
			"Transform":       transformTypeName,
		},
		vectorTypeName:     vectorTypeName,
		quaternionTypeName: quaternionTypeName,
		transformTypeName:  transformTypeName,
	}
}

func (tealWriter TealWriter) Write(api Api) (string, error) {
	if len(api.Functions) == 0 {
		return "", errors.New("API is empty")
	}

	stub := tealWriter.getStubHeader()

	// we iterate through every function within our api
	for i := 0; i < len(api.Functions); i++ {
		stub += tealWriter.getFunctionStub(api.Functions[i])
		stub += "\n"
	}

	return stub, nil
}

func (tealWriter TealWriter) getFunctionStub(function Function) string {
	// Desired
	// global Test: function(x: number, y: number): number, number

	functionStub := ""
	parameterDefinitions := ""
	var returnTypes string

	if len(function.Outputs) > 0 {
		returnTypes = ": "
	} else {
		returnTypes = ""
	}

	for i := 0; i < len(function.Inputs); i++ {
		input := function.Inputs[i]
		replacement := tealWriter.overrideVariableType[input.Type]

		if replacement != "" {
			input.Type = replacement
		}

		argument := ""

		argument += input.Name
		argument += ": "
		argument += input.Type

		if i < (len(function.Inputs) - 1) {
			argument += ", "
		}

		parameterDefinitions += argument
	}

	for i := 0; i < len(function.Outputs); i++ {
		output := function.Outputs[i]
		replacement := tealWriter.overrideVariableType[output.Type]

		if replacement != "" {
			output.Type = replacement
		}

		outputType := output.Type

		if i < (len(function.Outputs) - 1) {
			outputType += ", "
		}

		returnTypes += outputType
	}

	functionStub += fmt.Sprintf("---@see @https://www.teardowngame.com/modding/api.html#%s\n", function.Name)
	functionStub += fmt.Sprintf("global %s: function(%s)%s\n", function.Name, parameterDefinitions, returnTypes)

	return functionStub
}

func (tealWriter TealWriter) getStubHeader() string {
	return fmt.Sprintf(`
--[[
Teardown uses Lua version 5.1 as scripting language.
The Lua 5.1 reference manual can be found here. @https://www.lua.org/manual/5.1/

Created with HypnoTox's Teardown API Stub Generator, available at https://github.com/hypnotox/teardown-api-stub-generator
]]

--[[ Classes ]]

global type %s = {number, number, number}
global type %s = {number, number, number, number}

global record %s
    pos: %[1]s
    rot: %[2]s
end

--[[ Functions ]]

`, tealWriter.vectorTypeName, tealWriter.quaternionTypeName, tealWriter.transformTypeName)
}
