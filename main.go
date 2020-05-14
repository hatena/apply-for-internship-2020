package main

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
)

func main() {
	confirm := false
	prompt := &survey.Confirm{
		Message: "Do you want to apply for our internship?",
	}
	err := survey.AskOne(prompt, &confirm)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if !confirm {
		fmt.Println("See you again.")
		return
	}
	fmt.Println("Thank you!")
}
