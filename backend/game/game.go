package game

import (
  "database/sql"
)

type Piece string

const (
	Pawn   Piece = "P"
	Knight Piece = "N"
	Bishop Piece = "B"
	Rook   Piece = "R"
	Queen  Piece = "Q"
	King   Piece = "K"
)
type Game struct {
	ID            int    `json:"id"`
	WhiteID int    `json:"white_id"`
	BlackID int    `json:"black_id"`
	Result        string `json:"result"`
	DateCreated   string `json:"date_created"`
	DateFinished  sql.NullString `json:"date_finished"`
}
type Coord struct {
  X int `json:"x"`
  Y int `json:"y"`
}
type Move struct {
  ID int `json:"id"`
  GameID int `json:"game_id"`
  From *Coord `json:"from"`
  To *Coord `json:"to"`
  IsWhite bool `json:"is_white"`
  Piece *Piece `json:"piece"`
}
