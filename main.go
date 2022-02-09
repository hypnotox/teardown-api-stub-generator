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

	luaApi := api
	tealApi := api

	writeStubFile(src.NewLuaWriter(), luaApi, "teardown.d.lua")
	writeStubFile(src.NewTealWriter(), tealApi, "teardown.d.tl")
}

func writeStubFile(writer src.Writer, api src.Api, fileName string) {
	apiStub, writerError := writer.Write(api)

	if writerError != nil {
		fmt.Println("writerError: " + writerError.Error())
		return
	}

	_ = os.Remove(fileName)
	stubFile, fileOpenError := os.Create(fileName)

	if fileOpenError != nil {
		fmt.Println("fileOpenError: " + fileOpenError.Error())
		return
	}

	_, stubWriteError := stubFile.Write([]byte(apiStub))
	if stubWriteError != nil {
		fmt.Println(stubWriteError)
		return
	}

	fileCloseError := stubFile.Close()

	if fileCloseError != nil {
		fmt.Println(fileCloseError)
	}

	fmt.Println("Successfully wrote " + fileName)
}
