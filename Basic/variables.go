package main

import "fmt"

const LoggedInToken string="ghabbbhhjd" //public   

func main() {
	var username string = "ashutosh"
    fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool=true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8=255    //8 bits integer
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n",smallVal)

	var smallFloat float64=455.3286843843   //64 bit float
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n",smallFloat)

	//default values and same aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n",anotherVariable)

	//implicit
	var website="learncode.in"
	fmt.Println(website)

	//no var style
	numberOfUser := 300000.0   //without giving the data type (no var style) it automatically assign data type and we cant change it later
	fmt.Println(numberOfUser)

	fmt.Println(LoggedInToken)
}
