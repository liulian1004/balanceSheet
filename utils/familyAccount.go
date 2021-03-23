package utils

import (
	"fmt"
)

type FamilyAccount struct {
	//field
	userInput string
	//init variable
	balance float64
	amount float64
	notes string

	hasRecords bool;
	details string;
	isExit bool;
}
//factory => construct
func NewFamilyAccount () *FamilyAccount{
	return &FamilyAccount{
		userInput: "",
		balance: 0.00,
		amount: 0.00,
		notes:"",
		hasRecords: false,
		details: "",
		isExit: false,

	}
}
func(this *FamilyAccount) checkDetails(){
	fmt.Println("\n------------ Family Balance Sheet -----------")
				if !this.hasRecords {
					fmt.Println("\nNothing is registered")
				} else{
					fmt.Println(this.details)
				}
				
}
func (this *FamilyAccount) registerIncome(){
	fmt.Println("the income is: " )
				fmt.Scanln(&this.amount); // point to the address 
				this.balance += this.amount;
				fmt.Println("details: ")
				fmt.Scanln(&this.notes)
				this.details += fmt.Sprintf("\nin\t       %v\t%v\t%v", this.amount,this.balance,this.notes)
				this.hasRecords = true;

}

func (this *FamilyAccount) registerExpense(){
	fmt.Println("the expense is: " )
				fmt.Scanln(&this.amount); // point to the address 
				if this.amount > this.balance {
					fmt.Println("Expense can't bigger than balance")
				}else{
					this.balance -= this.amount;
					fmt.Println("details: ")
					fmt.Scanln(&this.notes)
					this.details += fmt.Sprintf("\nout\t       %v\t%v\t%v", this.amount,this.balance,this.notes)
				}
				

}

func (this *FamilyAccount) exit(){
	fmt.Println("Do you want to exit? Y/N")
				choice := ""
				for {
					fmt.Scanln(&choice);
					if choice == "Y" || choice == "N" {
						break; //break the for loop
					}
					fmt.Println("Only Y/N is acceptable")
				}
				if choice == "Y" {
					this.isExit = true
				}

}
//implement the method
func (this *FamilyAccount) DisplayMenu() {
	for{
		fmt.Println("\n------------ Family Balance Sheet -----------")
		fmt.Println("             1. Check Balance Sheet")
		fmt.Println("             2. Register Income")
		fmt.Println("             3. Register Expense")
		fmt.Println("             4. Exit")
		fmt.Println("Please select(1-4)")
		fmt.Scanln(&this.userInput) //scan user input
		switch this.userInput{
		case "1":
			this.checkDetails()
		case "2":
			this.registerIncome()
		case "3":
			this.registerExpense()
		case "4":
			this.exit()
		}
		if this.isExit{
			break
		}
	}
		
}