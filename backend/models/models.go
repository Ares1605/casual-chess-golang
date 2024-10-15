package models

import (
	"database/sql"
	"errors"

	"github.com/Ares1605/casual-chess-backend/game"
	"github.com/Ares1605/casual-chess-backend/oauth/googleuser"
	_ "github.com/mattn/go-sqlite3"
)

type BasicUser struct {
  ID int64 `json:"id"`
  DisplayName string `json:"display_name"`
  ProfileURL string `json:"profile_url"`
}
type User struct {
  ID int64 `json:"id"`
  DisplayName string `json:"display_name"`
  googleID string `json:"google_id"`
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
  statement := "INSERT INTO users (display_name, google_id, email, profile_url) VALUES (?, ?, ?, ?)"
  result, err := dbConn.Exec(statement, googleUser.Name, googleUser.ID, googleUser.Email, googleUser.Profile)
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
    googleID: googleUser.ID,
    Email: googleUser.Email,
    DisplayName: googleUser.Email,
  }, nil
}
func GetUserFromID(dbConn *sql.DB, id int) (*User, error) {
  user := User{}
  err := dbConn.QueryRow("SELECT id, display_name, google_id, email, profile_url FROM users WHERE id=?", id).Scan(
		&user.ID,
		&user.DisplayName,
		&user.googleID,
		&user.Email,
		&user.ProfileURL,
    )
  if err != nil {
    return nil, err
  }
  return &user, nil
}
func GetFriends(dbConn *sql.DB, googleID string) (*[]BasicUser, error) {
  var friends []BasicUser
  rows, err := dbConn.Query("SELECT u.id, u.display_name, u.profile_url FROM friends f INNER JOIN users u on IF(f.invitee_google_id=?, f.invited_google_id, f.invitee_google_id)=u.google_id WHERE f.invitee_id=? or f.invited_google_id=?", googleID, googleID, googleID)
  defer rows.Close()
  if err != nil {
    return nil, err
  }
  for rows.Next() {
    friend := BasicUser{}
    if err := rows.Scan(&friend.ID, &friend.DisplayName, &friend.ProfileURL); err != nil {
      return nil, errors.New("Issue processing row in the GetFriends query result")
    }
    friends = append(friends, friend)
  }
  return &friends, nil
}
func GetUser(dbConn *sql.DB, googleID string) (*User, error) {
  user := User{}
  err := dbConn.QueryRow("SELECT id, display_name, google_id, email, profile_url FROM users WHERE google_id=?", googleID).Scan(
		&user.ID,
		&user.DisplayName,
		&user.googleID,
		&user.Email,
		&user.ProfileURL,
    )
  if err != nil {
    return nil, err
  }
  return &user, nil
}
