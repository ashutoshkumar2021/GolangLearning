package main

import "fmt"

func main() {
	languages:=make(map[string]string)

	languages["JS"]="JavaScript"
	languages["TS"]="TypeScript"
	languages["PY"]="Python"

	fmt.Println("List of Languages: ",languages)
	fmt.Println("JS shorts for: ",languages["JS"])

	//remove element from map
	delete(languages,"TS")
	fmt.Println("List of Languages: ",languages)

	//loops are interesting in golang

	for _,value:=range languages{
		fmt.Printf("For key v,value is %v\n", value)
	}

}
