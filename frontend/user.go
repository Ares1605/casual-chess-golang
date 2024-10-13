package main
type User struct {
    ID          int    `json:"id"`
    DisplayName string `json:"display_name"`
    UUID        string `json:"uuid"`
    Email       string `json:"email"`
    ProfileURL  string `json:"profile_url"`
}
