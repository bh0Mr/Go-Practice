package main

import "fmt"



func main() {
    
    person := &Person{Name: "Alice", Age: 30}



    fmt.Println("Before Birthday:", person.Age) 

    person.Birthday() 

    fmt.Println("After Birthday:", person.Age)   
	fmt.Println("=================================")

	Data := User{"rin",20,"test@test123"}
	
	Data.GetData()

	fmt.Println("=================================")

	Data.NewData()

	fmt.Println("=================================")

	Data.GetData()


	fmt.Println("=================================")


    
//composit type bystructure embedding 

	myCar := Car{
        Engine: Engine{horsepower: 200},  
        make:   "Honda",
        model:  "Civic",
    }

    fmt.Println("Car make:", myCar.make)
    fmt.Println("Engine horsepower:", myCar.horsepower)


   
	fmt.Println("=================================")

	//Method Values and Expressions

	g := Greeter{name: "Test1"}

  
    greetAction := g.greet

   
    greetAction()  




	fmt.Println("=================================")

	//Encapsulation

	account := BankAccount{balance: 1000}
    account.deposit(500)
    account.withdraw(200)
    fmt.Println("Current Balance:", account.getBalance())


	
}







type Person struct {
	 Name string
	 Age int
	  } 
	  
 func (p *Person) Birthday() { 
	 p.Age += 1 
	
	}



//user 


type User struct {

    Name string
    Age int
    Email string 
}

func (u User) GetData(){

	fmt.Println("Name of the user ",u.Name)
	fmt.Println("Name of the user ",u.Age)
	fmt.Println("Name of the user ",u.Email)


}

func (u *User) NewData(){


	u.Name = "kin"
	u.Age = 22
	u.Email = "test2@test.com"

	fmt.Println("Name of the user ",u.Name)
	fmt.Println("Name of the user ",u.Age)
	fmt.Println("Name of the user ",u.Email)

}


//composit type by structure embedding 


type Engine struct {
    horsepower int
}

type Car struct {
    Engine  
    make    string
    model   string
}




//Method Values and Expressions


type Greeter struct {
    name string
}


func (g Greeter) greet() {
    fmt.Println("Hello, " + g.name)
}





//encapsulation 


type BankAccount struct {
    balance float64
}

func (b *BankAccount) deposit(amount float64) {
    if amount > 0 {
        b.balance += amount
    }
}

func (b *BankAccount) withdraw(amount float64) {
    if amount > 0 && amount <= b.balance {
        b.balance -= amount
    }
}

func (b BankAccount) getBalance() float64 {
    return b.balance
}








    











