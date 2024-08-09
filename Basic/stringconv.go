package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
// as reader takes input in string datatype and when we have to do some operations on input for example here rating, we use strconv to convert it

func main(){
	fmt.Println("Welcome to our Pizza App");
	fmt.Println("Please rate our pizza between 1 and 5")

	reader:=bufio.NewReader(os.Stdin)

	input,_:=reader.ReadString('\n')
    fmt.Println("Thanks for rating, ",input)

	numRating, err :=strconv.ParseFloat(strings.TrimSpace(input),64)  

	if err!=nil{
		fmt.Println(err)
	}else{
		fmt.Println("Added 1 to your rating: ",numRating+1)
	}


}
