# Variables and Types
* Go is a statically typed language: It means that the type of variable should be explicitly declared.
* Go is a strongly typed language: all functions, variables, and expressions must have types

## What is `:=` in go
When declaring a variable using `:=`, the type is inferred from the value on the right-hand side. This means that you do not need to explicitly state its type when declaring.

## Data Types
Category          | Type                                          | Zero Value | Description                                                                  |
|-------------------|-----------------------------------------------|------------|------------------------------------------------------------------------------|
| Boolean           | bool                                          | FALSE      | true or false                                                                |
| Signed Integers   | int, int8, int16, int32, int64                | 0          | Whole numbers. int defaults to 32 or 64 bits depending on hardware.          |
| Unsigned Integers | uint, uint8, uint16, uint32, uint64, uintptr  | 0          | Non-negative whole numbers.                                                  |
| Byte & Rune       | byte (alias for uint8),rune (alias for int32) | 0          | byte represents raw ASCII/bytes; s rune represents a single Unicode character. Similar to char in some languages.|
| Floats            | float32, float64                              | 0          | Decimals/floating-point numbers (float64 is standard).                       |
| Complex           | complex64, complex128                         | (0+0i)     | Numbers with real and imaginary parts.                                       |
| Strings           | string                                        |    ""        | Immutable sequence of UTF-8 encoded bytes.                                   |
