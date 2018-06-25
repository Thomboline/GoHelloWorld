/***
To run: go run main.go
To build: go build main.go (then execute ./main)
***/

/***
##ALWAYS IN THE TOP##

Just like in Java, but packages are only stated in the sourcecode.
This is an executable package, and as such, when build, leaves us with a "main" executable.
Reusable packages are more like dependencies and libraries.

You determine whether you're making an executable or reusable package with the name.
Specifically "main" is the executable keyword. Everything but "main" would create a reusable package.
**/
package main

/*** 
##ALWAYS AFTER PACKAGE##

Imports could be standard libraries or reusable packages authored by other engineers.
Just like in Java.
fmt is short for "format"
***/
import "fmt"

/***
Standard function declaration. Argument list in parentheses like any other language.
The main method, like in Java, is the one method where everything is run.
**/
func main() {
	fmt.Println("Hi there!")
}
