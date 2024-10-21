package models

import (
	"database/sql"
	"errors"

	"github.com/Ares1605/casual-chess-golang/backend/game"
	"github.com/Ares1605/casual-chess-golang/backend/oauth/googleuser"
	_ "github.com/mattn/go-sqlite3"
	"github.com/google/uuid"
)

type BasicUser struct {
  ID int64 `json:"id"`
  Username string `json:"username"`
  ProfileURL string `json:"profile_url"`
  GoogleID string `json:"google_id"`
}
type User struct {
  ID int64 `json:"id"`
  Username string `json:"username"`
  GoogleName string `json:"google_name"`
  GoogleID string `json:"google_id"`
  Email string `json:"email"`
  ProfileURL string `json:"profile_url"`
  SetupComplete bool `json:"setup_complete"`
}
type PendingRow struct {
  ID int64
  InvitedGoogleID string
  InviteeGoogleID string
  DateCreated string
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
  username := uuid.New().String()
  statement := "INSERT INTO users (google_name, username, google_id, email, profile_url) VALUES (?, ?, ?, ?, ?)"
  result, err := dbConn.Exec(statement, googleUser.Name, username, googleUser.ID, googleUser.Email, googleUser.ProfileURL)
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
    GoogleID: googleUser.ID,
    Email: googleUser.Email,
    GoogleName: googleUser.Email,
  }, nil
}
func GetUserFromID(dbConn *sql.DB, id int) (*User, error) {
  user := User{}
  err := dbConn.QueryRow("SELECT id, google_name, google_id, email, profile_url, username FROM users WHERE id=?", id).Scan(
		&user.ID,
		&user.GoogleName,
		&user.GoogleID,
		&user.Email,
		&user.ProfileURL,
		&user.Username,
    )
  if err != nil {
    return nil, err
  }
  return &user, nil
}
func GetFriends(dbConn *sql.DB, googleID string) (*[]BasicUser, error) {
  friends := []BasicUser{}
  rows, err := dbConn.Query("SELECT u.id, u.username, u.profile_url, u.google_id FROM friends f INNER JOIN users u on (CASE WHEN f.invitee_google_id=? THEN f.invited_google_id ELSE f.invitee_google_id END)=u.google_id WHERE f.invitee_google_id=? or f.invited_google_id=?", googleID, googleID, googleID)
  defer rows.Close()
  if err != nil {
    return nil, err
  }
  for rows.Next() {
    friend := BasicUser{}
    if err := rows.Scan(&friend.ID, &friend.Username, &friend.ProfileURL, &friend.GoogleID); err != nil {
      return nil, errors.New("Issue processing row in the GetFriends query result")
    }
    friends = append(friends, friend)
  }
  for i := 0; i < 15; i++ {
    friends = append(friends, BasicUser{
      ID: 1,
      Username: "gabster123",
      ProfileURL: "https://play-lh.googleusercontent.com/U6z-kQNP24tjHIjHgJPrkVhfJDxAeVFyKBuuV4C9g2YNPKBgw6M_GrGAjsbhQFx0SI4",
      GoogleID: "12345",
    })
  }
  return &friends, nil
}
func GetUser(dbConn *sql.DB, googleID string) (*User, error) {
  user := User{}
  var setupCompleteTinyInt int8
  err := dbConn.QueryRow("SELECT id, google_name, google_id, email, profile_url, username, setup_complete FROM users WHERE google_id=?", googleID).Scan(
		&user.ID,
		&user.GoogleName,
		&user.GoogleID,
		&user.Email,
		&user.ProfileURL,
		&user.Username,
		&setupCompleteTinyInt,
    )
  if err != nil {
    return nil, err
  }
  user.SetupComplete = setupCompleteTinyInt != 0
  return &user, nil
}
func GetPendingRow(dbConn *sql.DB, invitedGoogleID string, inviteeGoogleID string) (*PendingRow, error) {
  pendingRow := PendingRow{}
  statement := "SELECT id, invited_google_id, invitee_google_id, date_created FROM pending_friends WHERE invited_google_id=? AND invitee_google_id=?"
  err := dbConn.QueryRow(statement, invitedGoogleID, inviteeGoogleID).Scan(
    &pendingRow.ID,
    &pendingRow.InvitedGoogleID,
    &pendingRow.InviteeGoogleID,
    &pendingRow.DateCreated,
    )
  if err != nil {
    return nil, err
  }
  return &pendingRow, nil
}
func IsFriends(dbConn *sql.DB, googleIDOne string, googleIDTwo string) (bool, error) {
	statement := "SELECT COUNT(*) FROM friends WHERE invitee_google_id=? AND invited_google_id=? OR invited_google_id=? AND invitee_google_id=? LIMIT 1"
	var count uint8
	err := dbConn.QueryRow(statement, googleIDOne, googleIDTwo, googleIDOne, googleIDTwo).Scan(
		&count,
		)
	if err != nil {
	  return false, err
	}
	return count != 0, nil
}
func AddFriend(dbConn *sql.DB, invitedGoogleID string, inviteeGoogleID string) error {
  statement := "INSERT INTO friends (invited_google_id, invitee_google_id) VALUES (?, ?)"
  _, err := dbConn.Exec(statement, invitedGoogleID, inviteeGoogleID)
  return err
}
func DeletePendingFriendRequest(dbConn *sql.DB, rowID int64) error {
  _, err := dbConn.Exec("DELETE FROM pending_friends WHERE id=?", rowID)
	return err
}
func SetupUser(dbConn *sql.DB, googleUser *googleuser.GoogleUser, username string) (error) {
  statement := "UPDATE users SET username=?, setup_complete=1 WHERE google_id=?"
  result, err := dbConn.Exec(statement, username, googleUser.ID)
  if err != nil {
    return err
  }
  affected, err := result.RowsAffected()
  if err != nil {
    return err
  }
  if affected == 0 {
    return errors.New("No user was updated from the execute statement!")
  }
  return nil
}
func UsernameExists(dbConn *sql.DB, username string) (bool, error) {
  statement := "SELECT COUNT(*) FROM users WHERE username=? LIMIT 1"	
  var count uint8
  err := dbConn.QueryRow(statement, username).Scan(
    &count,
    )
  if err != nil {
    return false, err
  }
  return count != 0, nil
}
