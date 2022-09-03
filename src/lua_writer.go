package src

import (
	"errors"
	"fmt"
)

type LuaWriter struct {
	overrideVariableType map[string]string
	vectorTypeName       string
	quaternionTypeName   string
	transformTypeName    string
}

func NewLuaWriter() *LuaWriter {
	vectorTypeName := "VectorType"
	quaternionTypeName := "QuaternionType"
	transformTypeName := "TransformType"

	return &LuaWriter{
		overrideVariableType: map[string]string{
			"number<integer>": "number",
			"Vector":          vectorTypeName,
			"Quaternion":      quaternionTypeName,
			"Transform":       transformTypeName,
		},
		vectorTypeName:     vectorTypeName,
		quaternionTypeName: quaternionTypeName,
		transformTypeName:  transformTypeName,
	}
}

func (luaWriter LuaWriter) Write(api Api) (string, error) {
	if len(api.Functions) == 0 {
		return "", errors.New("API is empty")
	}

	stub := luaWriter.getStubHeader()

	// we iterate through every function within our api
	for i := 0; i < len(api.Functions); i++ {
		var function = api.Functions[i]
		stub += luaWriter.getFunctionStub(function)
		stub += "\n"
	}

	return stub, nil
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

func (luaWriter LuaWriter) getInputStub(inputs []Input) string {
	var inputStub string

	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		replacement := luaWriter.overrideVariableType[input.Type]

		if replacement != "" {
			input.Type = replacement
		}

		if input.Optional {
			input.Type += "|nil"
		}

		inputStub += fmt.Sprintf("---@param %s %s %s\n", input.Name, input.Type, input.Description)
	}

	return inputStub
}

func (luaWriter LuaWriter) getOutputStub(outputs []Output) string {
	if len(outputs) == 0 {
		return ""
	}

	returns := ""

	for i := 0; i < len(outputs); i++ {
		output := outputs[i]
		replacement := luaWriter.overrideVariableType[output.Type]

		if replacement != "" {
			output.Type = replacement
		}

		returns += fmt.Sprintf("---@return %s %s %s\n", output.Type, output.Name, output.Description)
	}

	return returns
}

func (luaWriter LuaWriter) getStubHeader() string {
	return fmt.Sprintf(`
--[[
Teardown uses Lua version 5.1 as scripting language.
The API can be found here: https://www.teardowngame.com/modding/api.html
The Lua 5.1 reference manual can be found here: https://www.lua.org/manual/5.1/

Created with HypnoTox's Teardown API Stub Generator, available at: https://github.com/hypnotox/teardown-api-stub-generator
]]

--[[ Classes ]]

---@class %s:table
---@see @https://www.teardowngame.com/modding/api.html#Vec
local defaultVector = {0, 0, 0}

---@class %s:table
---@see @https://www.teardowngame.com/modding/api.html#Quat
local defaultQuaternion = {0, 0, 0, 0}

---@class %s:table
---@field pos %[1]s Vector representing transform position
---@field rot %[2]s Quaternion representing transform rotation
---@see @https://www.teardowngame.com/modding/api.html#Transform
local defaultTransform = {
    pos = Vec(),
    rot = Quat(),
}

--[[ Functions ]]

`, luaWriter.vectorTypeName, luaWriter.quaternionTypeName, luaWriter.transformTypeName)
}
