package main

import (
	"fmt"
	"net/http"
)

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

	//player := manager.getPlayerID(1)

	// removed := manager.PlayerRemove(1)

	//fmt.Println("suppression réussite :", removed)

	//if player != nil {
	//	fmt.Printf("Joueur trouvé : %s / HP : %d\n", player.Name, player.HP)
	//} else {
	//	fmt.Println("Joueur introuvable\n")
	//}

	for _, p := range manager.Players {
		fmt.Printf("ID : %d | Nom : %s | HP : %d | Niveau : %d\n", 
		p.ID,
		p.Name,
		p.HP,
		p.Level,
	)
	}

	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/players", manager.PlayerHandler)

	fmt.Println("Serveur démarré sur le port 8080")

	err := http.ListenAndServe(":8080", nil)
	
	if err != nil {
		panic(err)
	}

}