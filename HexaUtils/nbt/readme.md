# 📦 nbt-go: A Powerful NBT Library for Go

[![Go Reference](https://pkg.go.dev/badge/github.com/Mansitoh/nbt-go.svg)](https://pkg.go.dev/github.com/Mansitoh/nbt-go)
[![Build Status](https://img.shields.io/github/actions/workflow/status/your-username/nbt-go/go.yml?branch=main)](https://github.com/Mansitoh/nbt-go/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/Mansitoh/nbt-go)](https://goreportcard.com/report/github.com/Mansitoh/nbt-go)

`nbt-go` is a robust and efficient library for working with NBT (Named Binary Tag) data in Go. This library provides all the necessary tools to serialize, deserialize, and manipulate NBT data structures commonly used in Minecraft and other applications. It supports reading and writing both named and unnamed NBT structures.

## ✨ Features

-   **Comprehensive NBT Support:** Handles all NBT tag types, including:
    -   End, Byte, Short, Int, Long
    -   Float, Double
    -   Byte Array, String
    -   List, Compound
    -   Int Array, Long Array
-   **Serialization:** Converts NBT data to binary format.
-   **Deserialization:** Reconstructs NBT data from binary format.
-   **Named and Unnamed NBT:** Supports both standard NBT (with root name) and Network NBT (without root name).
-   **Easy to Use API:** Provides a clean and intuitive interface.
-   **Error Handling:** Returns descriptive errors for easy debugging.
-   **CESU-8 Support:** Handles Java's modified UTF-8 encoding for strings.
-  **Helper functions:** Includes helpers for creating NBT structures


## 🚀 Usage

### Creating NBT Data
```go
package main

import (
	"fmt"
    "log"
	"nbt-go/nbt"
)

func main() {
	// Creating NBT data with a root name
	nbtData := nbt.NewNbt(
		"root_tag",
        nbt.NbtCompoundFromInterfaceMap(map[string]interface{}{
            "int_value": int32(12345),
            "string_value": "hello",
            "compound_value": map[string]interface{}{
                "float_value": float32(3.14),
            },
            "int_array_value": []int32{1, 2, 3},
        }),
	)


    bytes, err := nbtData.Write()
	if err != nil {
        log.Fatalf("Error writing nbt: %v", err)
    }

    println("Nbt Data: %v\n", bytes)

    nbtRead, err := nbt.DeserializeNbt(bytes)
    if err != nil{
        log.Fatalf("Error deserializing nbt: %v", err)
    }
    println("Nbt Read: %+v\n", nbtRead)

    intValue, ok := nbtRead.GetInt("int_value")
    if !ok {
        log.Fatalf("Error reading int_value")
    }
    fmt.Println("Int Value:", intValue)

    stringValue, ok := nbtRead.GetString("string_value")
    if !ok {
        log.Fatalf("Error reading string_value")
    }
     fmt.Println("String Value:", stringValue)

    compoundValue, ok := nbtRead.GetCompound("compound_value")
    if !ok{
        log.Fatalf("Error reading compound_value")
    }

    floatValue, ok := compoundValue.GetFloat("float_value")
     if !ok {
        log.Fatalf("Error reading float_value")
    }
     fmt.Println("Float Value:", floatValue)

    intArrayValue, ok := nbtRead.GetIntArray("int_array_value")
    if !ok{
        log.Fatalf("Error reading int_array_value")
    }
      fmt.Println("Int Array Value:", intArrayValue)



     unnamedBytes, err := nbtData.WriteUnnamed()
     if err != nil {
        log.Fatalf("Error writing unnamed nbt: %v", err)
    }

    println("Unnamed Nbt Data: %v\n", unnamedBytes)
    unnamedNbt, err := nbt.DeserializeUnnamedNbt(unnamedBytes)
     if err != nil {
        log.Fatalf("Error deserializing unnamed nbt: %v", err)
    }
    println("Unnamed Nbt Read: %+v\n", unnamedNbt)
}
```

### Serializing and Deserializing NBT Data

#### Named NBT (Standard)

```go
    // Serialize
    bytes, err := nbtData.Write()
    if err != nil {
        log.Fatalf("Error writing nbt: %v", err)
    }

	// Deserialize
    readNbt, err := nbt.DeserializeNbt(bytes)
	if err != nil{
        log.Fatalf("Error deserializing nbt: %v", err)
    }

    // Get the root name
    println("Root Name: %s\n", readNbt.GetName())

```

#### Unnamed NBT (Network NBT)

```go
    // Serialize
    unnamedBytes, err := nbtData.WriteUnnamed()
     if err != nil {
        log.Fatalf("Error writing unnamed nbt: %v", err)
    }

	// Deserialize
	unnamedNbt, err := nbt.DeserializeUnnamedNbt(unnamedBytes)
    if err != nil {
        log.Fatalf("Error reading unnamed nbt: %v", err)
    }
    // Root name will be ""
    println("Root Name: %s\n", unnamedNbt.GetName())
```

### Accessing NBT Tag Values

```go
    // Reading values from Nbt struct
    stringValue, ok := readNbt.GetString("string_value")
    if !ok {
        log.Fatalf("Error reading string_value")
    }
     fmt.Println("String Value:", stringValue)

    // Reading a nested compound
    compoundValue, ok := readNbt.GetCompound("compound_value")
    if !ok{
        log.Fatalf("Error reading compound_value")
    }

    floatValue, ok := compoundValue.GetFloat("float_value")
     if !ok {
        log.Fatalf("Error reading float_value")
    }
     fmt.Println("Float Value:", floatValue)

    intArrayValue, ok := readNbt.GetIntArray("int_array_value")
    if !ok{
        log.Fatalf("Error reading int_array_value")
    }
      fmt.Println("Int Array Value:", intArrayValue)

    // Use Get... methods to safely access NBT tag values.
    // Returns a value and a bool to check if it exists.
```

### Generating NBT with Helper Function

```go
import "nbt-go/nbt"

    nbtData := nbt.GenerateNbt(
		"my_root_nbt",
		"some_message_id",
		"linear",
		3.14,
	)

     // Serialize and access like before
```

## 🧰 API Reference

### `nbt.Nbt`

-   `NewNbt(name string, tag *NbtCompound) *Nbt`: Creates a new `Nbt` structure.
-   `Write() ([]byte, error)`: Serializes the NBT data with the root name.
-   `WriteUnnamed() ([]byte, error)`: Serializes the NBT data without the root name (Network NBT).
-  `GetName() string`: Returns the root name
-   `DeserializeNbt(r io.Reader) (*Nbt, error)`: Deserializes NBT data with the root name from `io.Reader`.
-   `DeserializeUnnamedNbt(r io.Reader) (*Nbt, error)`: Deserializes NBT data without the root name (Network NBT) from `io.Reader`.
-   `GetByte(name string) (int8, bool)`: Gets a byte tag by name.
-   `GetShort(name string) (int16, bool)`: Gets a short tag by name.
-   `GetInt(name string) (int32, bool)`: Gets an int tag by name.
-   `GetLong(name string) (int64, bool)`: Gets a long tag by name.
-   `GetFloat(name string) (float32, bool)`: Gets a float tag by name.
-   `GetDouble(name string) (float64, bool)`: Gets a double tag by name.
-   `GetBool(name string) (bool, bool)`: Gets a bool tag by name.
-   `GetString(name string) (string, bool)`: Gets a string tag by name.
-  `GetList(name string) ([]NbtTag, bool)`: Gets a list tag by name.
-   `GetCompound(name string) (*NbtCompound, bool)`: Gets a compound tag by name.
-   `GetIntArray(name string) ([]int32, bool)`: Gets an int array tag by name.
-   `GetLongArray(name string) ([]int64, bool)`: Gets a long array tag by name.

### `nbt.NbtCompound`

-  `NewNbtCompound() *NbtCompound`: Creates a new `NbtCompound` structure
-   `GetByte(name string) (int8, bool)`: Gets a byte tag by name.
-   `GetShort(name string) (int16, bool)`: Gets a short tag by name.
-   `GetInt(name string) (int32, bool)`: Gets an int tag by name.
-   `GetLong(name string) (int64, bool)`: Gets a long tag by name.
-   `GetFloat(name string) (float32, bool)`: Gets a float tag by name.
-   `GetDouble(name string) (float64, bool)`: Gets a double tag by name.
-   `GetBool(name string) (bool, bool)`: Gets a bool tag by name.
-   `GetString(name string) (string, bool)`: Gets a string tag by name.
-  `GetList(name string) ([]NbtTag, bool)`: Gets a list tag by name.
-   `GetCompound(name string) (*NbtCompound, bool)`: Gets a compound tag by name.
-   `GetIntArray(name string) ([]int32, bool)`: Gets an int array tag by name.
-   `GetLongArray(name string) ([]int64, bool)`: Gets a long array tag by name.
-   `serializeContent() ([]byte, error)`: Serializes the compound without tag and name.
-   `Put(key string, tag NbtTag)`: Put the tag with specified key

### Helper functions

-  `NbtCompoundFromMap(data map[string]NbtTag) *NbtCompound`: Creates an NbtCompound from a map
-  `NbtCompoundFromInterfaceMap(data map[string]interface{}) *NbtCompound`: Creates an NbtCompound with values that can be converted to NbtTags
-   `GenerateNbt(nbtName string, messageID string, scaling string, exhaustion float32) *Nbt`:  Generates a NBT struct with the specified data

## 🐛 Contributing

Contributions are welcome! Please feel free to submit issues and pull requests.

## ⚖️ License

This library is licensed under the [MIT License](LICENSE).

