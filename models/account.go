package model

import "github.com/spf13/cobra"

type Account struct {
	User User `json:"user"`
	Token string `json:"token"`
}

type User struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type AccountUseCases interface {
	SetEmail(string) error
	GetEmail(string) error
	SetPassword(string) error
	GetToken() (Account, error)
}

type AccountAdapter interface {
	SetEmail(cmd *cobra.Command, args []string)
	GetEmail(cmd *cobra.Command, args []string)
	SetPassword(cmd *cobra.Command, args []string)
	GetToken(cmd *cobra.Command, args []string)
}
