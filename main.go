package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"teardownApiStubGenerator/src"
)

func main() {
	// Open xmlFile
	xmlFile, fileOpenError := os.Open("api.xml")

	if fileOpenError != nil {
		fmt.Println(fileOpenError)
		return
	}

	fmt.Println("Successfully opened api.xml")

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(xmlFile)

	xmlReader := src.XmlReader{}
	var api, xmlReaderError = xmlReader.Read(byteValue)

	if xmlReaderError != nil {
		fmt.Println("xmlReaderError: " + xmlReaderError.Error())
		return
	}

	fmt.Println("Successfully parsed api.xml")
	fileCloseError := xmlFile.Close()

	if fileCloseError != nil {
		fmt.Println("fileCloseError: " + fileCloseError.Error())
		return
	}

	// Convert api object to Lua stub file string
	luaWriter := src.NewLuaWriter()
	apiStub, luaWriterError := luaWriter.Write(api)

	if luaWriterError != nil {
		fmt.Println("luaWriterError: " + luaWriterError.Error())
		return
	}

	fmt.Println("Successfully generated api stub")
	_ = os.Remove("stub.lua")
	stubFile, fileOpenError := os.Create("stub.lua")

	if fileOpenError != nil {
		fmt.Println("fileOpenError: " + fileOpenError.Error())
		return
	}

	fmt.Println("Successfully opened stub.lua")
	_, stubWriteError := stubFile.Write([]byte(apiStub))
	if stubWriteError != nil {
		fmt.Println(stubWriteError)
		return
	}

	// defer the closing of our stubFile so that we can parse it later on
	fileCloseError = stubFile.Close()

	if fileCloseError != nil {
		fmt.Println(fileCloseError)
	}

	fmt.Println("Successfully written stub.lua")
}
