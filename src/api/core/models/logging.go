package models

type ErrorLogging struct {
	User  User        `json:"user"`
	Error interface{} `json:"error"`
}
