package main

import (
	"fmt"
	"os"
	"strconv"
)

type CommandType string

const (
	int64uint64 CommandType = "int64_uint64"
	uint64int64             = "uint64_int64"
	int32uint32             = "int32_uint3"
	uint32int32             = "uint32_int3"
	int16uint16             = "int16_uint16"
	uint16int16             = "uint16_int16"
	int8uint8               = "int8_uint8"
	uint8int8               = "uint8_int8"
	unknown                 = "unknown"
)

func main() {

	var arguments []string = os.Args[1:]

	if len(arguments) < 2 {
		PrintHelp()
	} else {
		var cmd CommandType = GetCommandType(arguments[0])
		if cmd == unknown {
			PrintHelp()
		}
		var nums []string = arguments[1:]
		for _, str := range nums {
			switch cmd {
			case int64uint64:
				fmt.Printf("Converting INT64 %v to UINT64 %v\n", str, int64_to_uint64(str))
			case uint64int64:
				fmt.Printf("Converting UINT64 %v to INT64 %v\n", str, uint64_to_int64(str))
			case int32uint32:
				fmt.Printf("Converting INT32 %v to UINT32 %v\n", str, int32_to_uint32(str))
			case uint32int32:
				fmt.Printf("Converting UINT32 %v to INT32 %v\n", str, uint32_to_int32(str))
			case int16uint16:
				fmt.Printf("Converting INT16 %v to UINT16 %v\n", str, int16_to_uint16(str))
			case uint16int16:
				fmt.Printf("Converting UINT16 %v to INT16 %v\n", str, uint16_to_int16(str))
			case int8uint8:
				fmt.Printf("Converting INT8 %v to UINT8 %v\n", str, int8_to_uint8(str))
			case uint8int8:
				fmt.Printf("Converting UINT8 %v to INT8 %v\n", str, uint8_to_int8(str))
			}
		}
	}
}

func int64_to_uint64(param string) uint64 {
	origVal, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		return 0
	}
	var newVal = uint64(origVal)
	return newVal
}
func uint64_to_int64(param string) int64 {
	origVal, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		return 0
	}
	var newVal = int64(origVal)
	return newVal
}
func int32_to_uint32(param string) uint32 {
	origVal, err := strconv.ParseInt(param, 10, 32)
	if err != nil {
		return 0
	}
	var newVal = uint32(origVal)
	return newVal
}
func uint32_to_int32(param string) int32 {
	origVal, err := strconv.ParseUint(param, 10, 32)
	if err != nil {
		return 0
	}
	var newVal = int32(origVal)
	return newVal
}
func int16_to_uint16(param string) uint16 {
	origVal, err := strconv.ParseInt(param, 10, 16)
	if err != nil {
		return 0
	}
	var newVal = uint16(origVal)
	return newVal
}
func uint16_to_int16(param string) int16 {
	origVal, err := strconv.ParseUint(param, 10, 16)
	if err != nil {
		return 0
	}
	var newVal = int16(origVal)
	return newVal
}
func int8_to_uint8(param string) uint8 {
	origVal, err := strconv.ParseInt(param, 10, 8)
	if err != nil {
		return 0
	}
	var newVal = uint8(origVal)
	return newVal
}
func uint8_to_int8(param string) int8 {
	origVal, err := strconv.ParseUint(param, 10, 8)
	if err != nil {
		return 0
	}
	var newVal = int8(origVal)
	return newVal
}

func GetCommandType(param string) CommandType {

	if param == "int64_uint64" {
		return int64uint64
	} else if param == "uint64_int64" {
		return uint64int64
	} else if param == "int32_uint32" {
		return int32uint32
	} else if param == "uint32_int32" {
		return uint32int32
	} else if param == "int16_uint16" {
		return int16uint16
	} else if param == "uint16_int16" {
		return uint16int16
	} else if param == "int8_uint8" {
		return int8uint8
	} else if param == "uint8_int8" {
		return uint8int8
	} else {
		return unknown
	}
}

func PrintHelp() {
	fmt.Println("Command Line Parameters For Sign Conversion: conversion_type number [number]...")
	fmt.Println("All numbers should be in base 10 format.  Enter as many numbers as needes to convert.")
	fmt.Println(int64uint64)
	fmt.Println(uint64int64)
	fmt.Println(int32uint32)
	fmt.Println(uint32int32)
	fmt.Println(int16uint16)
	fmt.Println(uint16int16)
	fmt.Println(int8uint8)
	fmt.Println(uint8int8)
}
