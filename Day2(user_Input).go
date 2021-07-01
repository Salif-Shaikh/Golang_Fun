package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	/* simple user intput for string type*/
	var name string
	fmt.Scanln(&name)
	fmt.Println("Output:", name)
	/* but take only the input before space.*/

	//user input in string type
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Input: ")
	myoutput, _ := reader.ReadString('\n')
	fmt.Print(myoutput)

}
