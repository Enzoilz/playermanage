package main

import "slices"

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

func (m *PlayerManager) PlayerRemove(id int) bool {
	for i, v := range m.Players {
		if v.ID == id {
			m.Players = slices.Delete(m.Players, i, i+1)
			return true
		}
	}

	return false
}