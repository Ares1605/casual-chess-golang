package models

import (
	"database/sql"
	"errors"

	"github.com/Ares1605/casual-chess-backend/game"
	"github.com/Ares1605/casual-chess-backend/oauth/googleuser"
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
  ID int64 `json:"id"`
  DisplayName string `json:"display_name"`
  UUID string `json:"uuid"`
  Email string `json:"email"`
  ProfileURL string `json:"profile_url"`
}

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
func CreateUser(dbConn *sql.DB, googleUser *googleuser.GoogleUser) (*User, error) {
  statement := "INSERT INTO users (display_name, uuid, email, profile_url) VALUES (?, ?, ?, ?)"
  result, err := dbConn.Exec(statement, googleUser.Name, googleUser.UUID, googleUser.Email, googleUser.Profile)
  if err != nil {
    return nil, err
  }

  id, err := result.LastInsertId()
  if err != nil {
    return nil, err
  }
  // return the user we ASSUME was inserted into db
  return &User{
    ID: id,
    UUID: googleUser.UUID,
    Email: googleUser.Email,
    DisplayName: googleUser.Email,
  }, nil
}
func GetUserFromID(dbConn *sql.DB, id int) (*User, error) {
  user := User{}
  err := dbConn.QueryRow("SELECT id, display_name, uuid, email, profile_url FROM users WHERE id=?", id).Scan(
		&user.ID,
		&user.DisplayName,
		&user.UUID,
		&user.Email,
		&user.ProfileURL,
    )
  if err != nil {
    return nil, err
  }
  return &user, nil
}
func GetUserFromUUID(dbConn *sql.DB, uuid string) (*User, error) {
  user := User{}
  err := dbConn.QueryRow("SELECT id, display_name, uuid, email, profile_url FROM users WHERE uuid=?", uuid).Scan(
		&user.ID,
		&user.DisplayName,
		&user.UUID,
		&user.Email,
		&user.ProfileURL,
    )
  if err != nil {
    return nil, err
  }
  return &user, nil
}
