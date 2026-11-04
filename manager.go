package main

type PlayerManager struct {
	Players []Player 
}

func (m *PlayerManager) addPlayer(player Player)  {
	m.Players = append(m.Players, player)
}