package model

type Role int

const (
	Customer Role = iota
	Collaborator
	Manager
	Administrator
)
