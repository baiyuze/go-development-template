package dto

type Role struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Users       []int  `json:"users,omitempty"`
	Permissions []int  `json:"permissions,omitempty"`
}
