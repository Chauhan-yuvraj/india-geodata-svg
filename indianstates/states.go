package indianstates

import "fmt"

type District struct {
	Id   string
	Name string
	D    string
}

type State struct {
	Id       string
	Name     string
	District []District
}

var allStates map[string]State

func init(){
	
	allStates = make(map[string]State)
	allStates[]
}