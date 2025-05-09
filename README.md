# Tools Text Processor

## Description
Tools is a command-line Go application designed to process text files by applying a series of transformations and formatting rules. It reads an input `.txt` file, performs conversions such as hexadecimal and binary to decimal, adjusts articles ("a" to "an"), applies case transformations based on embedded commands, and formats punctuation spacing. The processed output is saved to a specified output `.txt` file.

## Prerequisites
- Go 1.23.3 or higher installed on your system.

## Installation
1. Clone the repository or download the source code.
2. Navigate to the project directory.
3. Build the executable using the Go tool:

```bash
go build -o tools main.go
```

This will create an executable named `tools` in the project directory.

## Usage
Run the program from the command line with two arguments: the input file and the output file. Both files must have a `.txt` extension.

```bash
./tools <input_file.txt> <output_file.txt>
```

Example:

```bash
./tools input.txt output.txt
```

## Features and Transformations

- **Hexadecimal to Decimal Conversion:** Converts any hexadecimal numbers followed by `(hex)` to their decimal equivalents.
- **Binary to Decimal Conversion:** Converts any binary numbers followed by `(bin)` to their decimal equivalents.
- **Article Adjustment:** Changes the article "a" to "an" when followed by words starting with vowels or 'h'.
- **Bracketed Commands:** Supports inline commands enclosed in parentheses to modify text:
  - `(up, N)`: Converts the last N words before the command to uppercase.
  - `(cap, N)`: Capitalizes the first letter of the last N words before the command.
  - `(low, N)`: Converts the last N words before the command to lowercase.
- **Punctuation Formatting:** Adjusts spacing around punctuation marks for cleaner output.

## Input File Format
- The input file should be a plain text file with `.txt` extension.
- Commands must be enclosed in parentheses and follow the format `(command, number)`.
- Supported commands are `up`, `cap`, and `low` with a numeric argument.

## Example
Input line:

```
this is a test (up, 2)
```

Output line:

```
this is A TEST
```

## License
This project is licensed under the MIT License.

## Author
El Mahoudi Abderrahman (Mars)