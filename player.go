package main

type player struct {
	name         string
	purse        int
	place        int
	inPenaltyBox bool
}

func NewPlayer(name string) *player {
	return &player{
		name:         name,
		purse:        0,
		place:        0,
		inPenaltyBox: false,
	}
}
