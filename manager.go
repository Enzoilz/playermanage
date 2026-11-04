package main

type PlayerManager struct {
	Players []Player 
}

func (m *PlayerManager) addPlayer(player Player)  {
	m.Players = append(m.Players, player)
}

func (m *PlayerManager) getPlayerID(id int) *Player  {
	for i := range m.Players {
		if m.Players[i].ID == id {
			return &m.Players[i]
		}
	}

	return nil
}

func (m *PlayerManager) updatePlayerHP(id int, hp int) bool  {
	player := m.getPlayerID(id)
	if player != nil {
		player.HP = hp
		return true
	} else {
		return false
	}
}