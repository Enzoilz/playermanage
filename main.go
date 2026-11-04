package main

import "fmt"

func main()  {
	manager := &PlayerManager{}

	player1 := Player{
        ID:    1,
        Name:  "Link",
        HP:    100,
        Level: 2,
    }

	player2 := Player{
        ID:    2,
        Name:  "Ganondorf",
        HP:   200,
        Level: 30,
    }

	manager.addPlayer(player1)
	manager.addPlayer(player2)

	for _, p := range manager.Players {
		fmt.Printf("ID : %d | Nom : %s | HP : %d | Niveau : %d\n", 
		p.ID,
		p.Name,
		p.HP,
		p.Level,
	)
	}
}