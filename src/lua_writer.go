package src

import (
	"errors"
	"fmt"
)

type LuaWriter struct {
	keywords                map[string]bool
	overrideVariableType    map[string]string
	overrideTableTypeByName map[string]string
}

func NewLuaWriter() *LuaWriter {
	return &LuaWriter{
		keywords: map[string]bool{
			"end":      true,
			"return":   true,
			"function": true,
		},
		overrideVariableType: map[string]string{
			"int":     "number",
			"float":   "number",
			"value":   "number",
			"varying": "any",
		},
	}
}

func (luaWriter LuaWriter) Write(api Api) (string, error) {
	if len(api.Functions) == 0 {
		return "", errors.New("API is empty")
	}

	luaWriter.prepareData(&api)
	stub := luaWriter.getStubHeader()

	// we iterate through every function within our api
	for i := 0; i < len(api.Functions); i++ {
		stub += luaWriter.getFunctionStub(api.Functions[i])
		stub += "\n"
	}

	return stub, nil
}

func (luaWriter LuaWriter) prepareData(api *Api) {
	for i := 0; i < len(api.Functions); i++ {
		function := api.Functions[i]

		for i := 0; i < len(function.Inputs); i++ {
			input := function.Inputs[i]

			if luaWriter.keywords[input.Name] {
				input.Name += "Value"
			}

			replacement := luaWriter.overrideVariableType[input.Type]

			if replacement != "" {
				input.Type = replacement
			}
		}

		for i := 0; i < len(function.Outputs); i++ {
			output := function.Outputs[i]

			if luaWriter.keywords[output.Name] {
				output.Name += "Value"
			}

			replacement := luaWriter.overrideVariableType[output.Type]

			if replacement != "" {
				output.Type = replacement
			}
		}
	}
}

func (luaWriter LuaWriter) getFunctionStub(function Function) string {
	functionStub := ""
	parameterNames := ""

	for i := 0; i < len(function.Inputs); i++ {
		parameterNames += function.Inputs[i].Name

		if i < (len(function.Inputs) - 1) {
			parameterNames += ", "
		}
	}

	functionStub += luaWriter.getInputStub(function.Inputs)
	functionStub += luaWriter.getOutputStub(function.Outputs)
	functionStub += fmt.Sprintf("---@see @https://www.teardowngame.com/modding/api.html#%s\n", function.Name)
	functionStub += fmt.Sprintf("function %s(%s) end\n", function.Name, parameterNames)

	return functionStub
}

func (luaWriter LuaWriter) getInputStub(inputs []*Input) string {
	var inputStub string

	for i := 0; i < len(inputs); i++ {
		input := inputs[i]

		if input.Optional {
			inputStub += fmt.Sprintf("---@param %s %s %s %s\n", input.Name, input.Type, "_optional_", input.Description)
		} else {
			inputStub += fmt.Sprintf("---@param %s %s %s\n", input.Name, input.Type, input.Description)
		}
	}

	return inputStub
}

func (luaWriter LuaWriter) getOutputStub(outputs []*Output) string {
	if len(outputs) == 0 {
		return ""
	}

	returnTypes := ""
	returnDescriptions := ""

	for i := 0; i < len(outputs); i++ {
		output := outputs[i]

		returnTypes += output.Type
		returnDescriptions += fmt.Sprintf("%s: %s", output.Name, output.Description)

		if i < (len(outputs) - 1) {
			returnTypes += "|"
			returnDescriptions += ", "
		}
	}

	return fmt.Sprintf("---@return %s %s\n", returnTypes, returnDescriptions)
}

func (luaWriter LuaWriter) getStubHeader() string {
	return `
--[[
Teardown uses Lua version 5.1 as scripting language.
The Lua 5.1 reference manual can be found here. @https://www.lua.org/manual/5.1/

Created with HypnoTox's Teardown API Stub Generator, available at https://github.com/hypnotox/teardown-api-stub-generator
]]

--[[ Functions ]]

`
}
