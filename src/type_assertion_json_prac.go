package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	fmt.Println("=================================")
	 //interface and assertion


	 jsonData4 := `{"name":"Atul","age":21}`

	 var data2 interface{}
	 err := json.Unmarshal([]byte(jsonData4),&data2)

	 if err != nil {

		 fmt.Println("Error While Unmarshling")


		 return
	 }

	 fmt.Println(data2)

	 if m,ok := data2.(map[string]interface{}); ok {

		fmt.Println("Name:",m["name"])
		fmt.Println("Age:",m["age"])
	 }



	 //2nd example 
	 fmt.Println("=================================")


	 var OnlineGame interface{}

	 OnlineGame = Game {
		 GameName:"Free fire",
		 CharacterName:"Adam",
		 CharacterId:12345,
	 }

	 fmt.Println(OnlineGame)


	 if G,ok := OnlineGame.(Game); ok {

		fmt.Println(G.GameName)
		fmt.Println(G.CharacterName)
		fmt.Println(G.CharacterId)
	 }











}

type Game struct{

	GameName string

	CharacterName string

	CharacterId int

}

