package main //输出包的名字
// import the library
import "fmt"

func main(){

	userInput := ""
	//init variable
	balance := 0.00
	amount := 0.00
	notes := ""

	hasRecords := false
	// \t for alignment
	details := "income/outcome\tamount\tbalace\tnotice" // here is inti the title of details
	isExit := false
	//render the main menu
	for{

		fmt.Println("\n------------ Family Balance Sheet -----------")
		fmt.Println("             1. Check Balance Sheet")
		fmt.Println("             2. Register Income")
		fmt.Println("             3. Register Expense")
		fmt.Println("             4. Exit")
		fmt.Println("Please select(1-4)")

		fmt.Scanln(&userInput) //scan user input

		switch userInput{
			case "1":
				fmt.Println("\n------------ Family Balance Sheet -----------")
				if !hasRecords {
					fmt.Println("\nNothing is registered")
				} else{
					fmt.Println(details)
				}
				
			case "2":
				fmt.Println("the income is: " )
				fmt.Scanln(&amount); // point to the address 
				balance += amount;
				fmt.Println("details: ")
				fmt.Scanln(&notes)
				details += fmt.Sprintf("\nin\t       %v\t%v\t%v", amount,balance,notes)
				hasRecords = true;

			case "3":
				fmt.Println("the expense is: " )
				fmt.Scanln(&amount); // point to the address 
				if amount > balance {
					fmt.Println("Expense can't bigger than balance")
					break; // only exit the switch
				}else{
					balance -= amount;
				}
				fmt.Println("details: ")
				fmt.Scanln(&notes)
				details += fmt.Sprintf("\nout\t       %v\t%v\t%v", amount,balance,notes)
			case "4":
				fmt.Println("Do you want to exit? Y/N")
				choice := ""
				for {
					fmt.Scanln(&choice);
					if choice == "Y" || choice == "N" {
						break; //break the for loop
					}
					fmt.Println("Only Y/N is accecptable")
				}
				if choice == "Y" {
					isExit = true
				}
				
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