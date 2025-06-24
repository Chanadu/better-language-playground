# <a href="#better-language">Better-Language</a>

## <a href="#description">Description</a>
This is a project that creates a new programming language by creating a tree-walking interpreter. 

## <a href="#how-to-run">How To Run</a>
1. Clone the repository
2. Open the terminal and navigate to the directory where the repository is located
3. Run the following command in the terminal to build to:
``` bash
go build -o gbpl . 
```
4. Create a .bpl file to write code in and run the following command in the terminal:
``` bash
./gbpl <filename>.bpl
```
or run this for a very limited REPL environment:
``` bash
./gbpl
```

## <a href="#why">Why</a>
I created this project to learn more about how programming languages are created and how interpreters work. This was a personal project not connected to any other organization or activity.

## <a href="#syntax">Syntax</a>
The language syntax is similar to the C programming language. <br>
Currently, the language supports the following:
- Variables
  - Declaration
    - `var x = 5` 
  - Assignment
	- `x = 10`
- Functions 
	- Declaration with Arguments
      - `function add(x, y) {}`
	- Calls
		- `add(1, 2)`
	- Return Values
		- `return x + y`
    - Function Recursion
- If & else statements
	- `if (x > 5) {} else {}` 
- For loops
	- `for (var i = 0; i < 5; i = i + 1) {}`
- While loops
  - `while (x < 5) {}`
- Print statements
  - `print(x)`
- Scope
  - `var x = 5; { var x = 10; }`
- Arithmetic operations
	 - Addition
		 - `x + y` 
	 - Subtraction
		 - `x - y`
	 - Multiplication 
		 - `x * y` 
	 - Division
		 - `x / y` 
	 - Modulus
		- `x % y`
	- Logical operations
	  - And
		- `x && y`
	  - Or
		  - `x || y`
	  - Not
		  - `!x`
- Comparison operations
	- Greater than
		- `x > y`
   - Greater than or equal to
        - `x >= y`
   - Less than
        - `x < y`
   - Less than or equal to
        - `x <= y`
    - Equal to
        - `x == y`
	- Not equal to
        - `x != y`
- Bitwise Operators
  - Bitwise AND
	- `x & y`
  - Bitwise OR
      - `x | y`
  - Bitwise XOR
      - `x ^ y`
  - Bitwise NOT
      - `~x`
  - Bitwise Left Shift
      - `x << y`
  - Bitwise Right Shift
      - `x >> y`
- Data Types
  - Integers
	- `var x = 5`
  - Booleans
	- `var x = true`
  - Strings
	- `var x = "Hello, World!"`
- Ternary
	  - `x > 5 ? 10 : 20`
- Comments
  - `// This is a comment`
- Builtin Functions
  - `clock()` - Returns the current time from Unix in milliseconds

## <a href="#features">Features</a>
- [x] Scanner
- [x] Parser
- [x] Interpreter