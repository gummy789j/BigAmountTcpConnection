package main

import (
	"fmt"
	"go_code/MultiusersChatRoom/client/processes"
	"log"
)

func main() {

	var key int

	// loop := true

	var userId int

	var userPwd string

	for {

		fmt.Println("---------------Welcome to Million Chat Room System------------------")
		fmt.Println("\t\t\t 1 Log in")
		fmt.Println("\t\t\t 2 Sign up")
		fmt.Println("\t\t\t 3 Log out")
		fmt.Println("Please Enter Your Choice(1-3): ")

		fmt.Scanf("%d\n", &key)

		switch key {
		case 1:
			fmt.Println("Log in the Million Chat Room !")
			fmt.Println("Please enter User ID: ")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("Please enter User Password: ")
			fmt.Scanf("%s\n", &userPwd)
			up := &processes.UserProcess{}
			err := up.Login(userId, userPwd)
			if err != nil {
				log.Println("Login Fail")
			}

			//loop = false
		case 2:
			fmt.Println("Sign up an account")
			//loop = false
		case 3:
			fmt.Println("3.Leave")
			//loop = false
		default:
			fmt.Println("The Choice does not Exit. Please CHOOSE AGAIN....")
		}

	}

	// 選擇後的用戶介面
	// if key == 1 {

	// 	fmt.Println("Please enter User ID: ")
	// 	fmt.Scanf("%d\n", &userId)
	// 	fmt.Println("Please enter User Password: ")
	// 	fmt.Scanf("%s\n", &userPwd)
	// 	err := login(userId, userPwd)
	// 	if err != nil {
	// 		log.Println("login fail")
	// 	}
	// } else if key == 2 {
	// 	fmt.Println("sign up program")
	// }
}
