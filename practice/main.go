package main

import (
	"bufio"   // bufio is a package in the Go standard library that provides buffered I/O (input/output) functionality
	"fmt"     // fmt is a package in the Go standard library that provides formatting for input and output
	"os"      // os is a package in the Go standard library that provides a platform-independent interface to operating system functionality (like file I/O)
	"strconv" // strconv is a package in the Go standard library that provides functions for converting strings to other types (like integers and floats)
	"strings" // strings is a package in the Go standard library that provides string manipulation functions (like searching, replacing, and splitting)
)

func main() {

	reader := bufio.NewReader(os.Stdin) // Create a new reader that reads from the standard input (keyboard)
	fmt.Print("Enter your name: ")      // Print a message to the console. The user will see this message and know to enter their name
	input, _ := reader.ReadString('\n') // Read the user's input until they press the Enter key. The input is stored in the variable input
	fmt.Println("You entered:", input)  // Print a message to the console that shows the user's input

	fmt.Print("Enter a number: ")                                      // Print a message to the console. The user will see this message and know to enter a number (integer)
	numInput, _ := reader.ReadString('\n')                             // Read the user's input until they press the Enter key. The input is stored in the variable numInput as a string. the line feed character '\n' is the delimiter that tells the reader to stop reading input
	aFloat, err := strconv.ParseFloat(strings.TrimSpace(numInput), 64) // Convert the string input to a float64. The strings.TrimSpace function removes any leading or trailing whitespace from the input. The 64 is the bit size of the float64 type meaning that the float64 type is 64 bits in size
	if err != nil {                                                    // Check if there was an error converting the string to a float64
		fmt.Println(err) // Print the error message to the console
	} else {
		fmt.Println("You entered:", aFloat) // Print a message to the console that shows the user's input
	}

}
