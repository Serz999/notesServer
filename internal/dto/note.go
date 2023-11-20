package dto

type Id string

type Note struct {
    Id Id `json:"id"`
    AuthorFirstName string `json:"authorFirstName"`
    AuthorLastName string `json:"authorLastName"`
    Note string `json:"note"`
}
