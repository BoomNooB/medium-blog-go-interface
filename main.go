package main

import (
	"go-interface/couple"
	"go-interface/diane"
)

func main() {
	diane := diane.NewDiane()

	// Create a new match
	lover := couple.NewCouple(diane)
	lover.DoWhatLoverDo()
}
