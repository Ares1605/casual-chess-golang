package models

import (
	"database/sql"
	"errors"

	_ "github.com/mattn/go-sqlite3"
	"github.com/Ares1605/casual-chess-backend/game"
)

func parsePiece(pieceStr string) (game.Piece, error) {
  switch pieceStr {
  case "P":
    return game.Pawn, nil
  case "K":
    return game.King, nil
  case "N":
    return game.Knight, nil
  case "Q":
    return game.Queen, nil
  case "B":
    return game.Bishop, nil
  case "R":
    return game.Rook, nil
  default:
    return "", errors.New("Failed to parse piece")
  }
}
func GetMoves(dbConn *sql.DB, gameID int) (*[]game.Move, error) {
  var moves []game.Move

  rows, err := dbConn.Query("SELECT id, game_id, from_x, from_y, to_x, to_y, is_white, piece FROM moves WHERE game_id=?", gameID)
  defer rows.Close()
  if err != nil {
    return nil, err
  }
  for rows.Next() {
    move := game.Move{}
    var pieceStr string
    if err := rows.Scan(&move.ID, &move.GameID, &move.From.X, &move.From.Y, &move.To.X, &move.To.Y, &move.IsWhite, pieceStr); err != nil {
      return nil, errors.New("Issue processing row in the getMoves query result")
    }
    piece, err := parsePiece(pieceStr)
    if err != nil {
      return nil, err
    }
    move.Piece = &piece
    moves = append(moves, move)
  }
  return &moves, nil
}
func GetGame(dbConn *sql.DB, id int) (*game.Game, error) {
  game := game.Game{}
  err := dbConn.QueryRow("SELECT id, white_id, black_id, result, date_created, date_finished FROM games WHERE id=?", id).Scan(
		&game.ID,
		&game.WhiteID,
		&game.BlackID,
		&game.Result,
		&game.DateCreated,
		&game.DateFinished,
    )
  if err != nil {
    return nil, err
  }
  return &game, nil
}
