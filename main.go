package main

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
)

func askDoYouWantToApply() (bool, error) {
	confirmation := false
	prompt := &survey.Confirm{
		Message: "Do you want to apply for our internship?",
	}
	if err := survey.AskOne(prompt, &confirmation); err != nil {
		return confirmation, err
	}
	return confirmation, nil
}

func askName() (string, error) {
	name := ""
	prompt := &survey.Input{
		Message:  "What is your name?",
	}
	if err := survey.AskOne(prompt, &name); err != nil {
		return name, err
	}
	return name, nil
}

func main() {
	wantToApply, err := askDoYouWantToApply()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if !wantToApply {
		fmt.Println("See you again.")
		return
	}
	fmt.Println("Thank you!")

	name, err := askName()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("Hello, %s.", name)
}
