package main

import "fmt"

func main() {
	fmt.Println("Welcome to class on slices")

	var fruitList=[]string{"Apple","Tomato","Peach"}
	fmt.Printf("Type of fruitList is %T\n",fruitList)
	fmt.Println(fruitList);

	fruitList=append(fruitList,"Mango","Banana")
	fmt.Println(fruitList);

	fruitList=append(fruitList[1:3]) // printing starts from 1 and stops at 3
	fmt.Println(fruitList);
	fruitList=append(fruitList[:3]) //printing all till 3

	highScores:=make([]int,4) //make is a part of memory management which helps to create memory

	highScores[0]=234
	highScores[1]=945
	highScores[2]=455
	highScores[3]=889
	//highScores[4]=888  //when we try to print this we will get an error, showing index out of range[4] with length4

	highScores=append(highScores, 555,666,321) //but when we use append it does not show any error

	fmt.Println(highScores)

	sort.Ints(highScores); //help to sor
	fmt.Println(highScores);

        fmt.Println(sort.IntsAreSorted(highScores)) //Return true or false

	//how to remove a value from slices based on index

	var courses = []string{"ReactJs", "Javascript", "swift", "python", "ruby"}
	fmt.Println(courses)

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	index, err := strconv.Atoi(strings.TrimSpace(input))

	if err != nil {
		fmt.Println("Error:", err)
	} else if index < 0 || index >= len(courses) {
		fmt.Println("Index out of range")
	} else {
		courses = append(courses[:index], courses[index+1:]...)
		fmt.Println(courses)

}
}
