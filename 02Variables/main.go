package main

import "fmt"

const MinimumAgeToVote = 18

func main() {
	fmt.Println("Variables.........")
	var username string = "PunitPoonia"
	fmt.Println(username)
	fmt.Printf("Type of variable is : %T \n", username)

	//Boolean

	var isMale bool = true
	fmt.Println(isMale)
	fmt.Printf("Type of variable is : %T \n", isMale)

	// Float

	var decimalValue float64 = 56.78765433456765432
	fmt.Println(decimalValue)
	fmt.Printf("Type of variable is : %T \n", decimalValue)

	// Default values
	var notInitializedInt int
	fmt.Println(notInitializedInt)
	fmt.Printf("Type of variable is : %T \n", notInitializedInt)

	// Implicit type not need to define which type of variable is this

	var name = "Punit Poonia"
	fmt.Println(name)
	fmt.Printf("Type of variable is : %T \n", name)

	// No Var Way to declare variable

	ageToVote := 18
	fmt.Println(ageToVote)

	// Const Variable

	fmt.Println("Public Variable declared here ", MinimumAgeToVote)

}
