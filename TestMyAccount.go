package main //输出包的名字
// import the library
import "fmt"

func main(){

	userInput := ""

	isExit := false
	//render the main menu
	for{

		fmt.Println("------------ Family Balance Sheet -----------")
		fmt.Println("             1. Balance Sheet")
		fmt.Println("             2. Income")
		fmt.Println("             3. Expense")
		fmt.Println("             4. Exit")
		fmt.Println("Please select(1-4)")

		fmt.Scanln(&userInput) //scan user input

		switch userInput{
			case "1":
				fmt.Println("------------ Family Balance Sheet -----------")
			case "2":
			case "3":
				fmt.Println("------------ Regsiter Expense -----------")
			case "4":
				isExit = true
			default:
				fmt.Println("------------ Please input 1-4 -----------")
				
		}
		if isExit{
			break
		}
	}
	//after exit for loop
	fmt.Println("------------ Thank you for using this tool -----------")
}