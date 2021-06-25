package main

import "fmt" //implements formatted I/O with functions analogous to C's printf and scanf. The format 'verbs' are derived from C's but are simpler.

func main() {
	//Declare the variable with keyword 'var' and its type
	var a string = "Declaring variable"
	var b int = 100
	//int type comes with size 8, 16, 32, 64 bits
	var c float32 = 3.00
	// float type comes with size 32, 64 bits

	// Other method to declare var
	//this method of declaration gets the type automaticaly
	d := "string"
	e := 200
	f := 2.

	fmt.Print(a, "\n", b, "\n", c, "\n", d, "\n", e, "\n", f)
	// to get the type of var
	fmt.Printf("\n%T", d)

	//Declaring 2 or more var at single time
	var sa, lif string = "hi!", "user"
	fmt.Println("\n", sa, lif)
	// Other method
	var (
		site = "github.com"
		num  = 1101
		dec  = "empty string"
	)

	fmt.Println(site, num, dec)

}
