package main

import "fmt"

// enumeratred variables/types there no inbuilt enum in go
// we implement enums using const

func changeStatus(status string) {
	fmt.Println("Status changed to ", status)
}

func main() {
	changeStatus("confirmed")
}
