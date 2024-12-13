package apperror

import "net/http"

const (
	ErrRegisteredEmail    = "E-mail already registered"
	ErrEmailWrong         = "Email not found"
	ErrQuery              = "Error query"
	ErrHashing            = "Error hashing"
	ErrUsernameNotFound   = "Username or password is wrong"
	ErrTokenWrong         = "Token not found"
	ErrBinding            = "Error binding"
	ErrAuthorization      = "Unauthorized"
	ErrWalletNotFound     = "Wallet not found"
	ErrTopUpQuantity      = "Please insert minimum amount of Rp. 50.000 and maximum of Rp. 10.000.000"
	ErrTransferQuantity   = "Please insert minimum amount of Rp. 1.000 and maximum of Rp. 50.000.000"
	ErrInsuficientFund    = "Insuficient funds"
	ErrCannotTransferSelf = "Cannot transfer to your own wallet"
	ErrPasswordPolicy     = "Please follow our password policy"
	ErrInsuficientChance  = "You don't have chance to play the game"
	ErrBoxIsUnavailable   = "Box number is unavailable"
	ErrServer             = "Server error"
	ErrSourceNotFound     = "Source of funds not found"
	ErrDateFormat         = "Date format incorrect"
	ErrFormat             = "Format incorrect"
)

var StatusCode = map[string]int{
	ErrRegisteredEmail:    http.StatusBadRequest,
	ErrEmailWrong:         http.StatusBadRequest,
	ErrQuery:              http.StatusInternalServerError,
	ErrHashing:            http.StatusInternalServerError,
	ErrUsernameNotFound:   http.StatusBadRequest,
	ErrTokenWrong:         http.StatusBadRequest,
	ErrBinding:            http.StatusInternalServerError,
	ErrAuthorization:      http.StatusUnauthorized,
	ErrWalletNotFound:     http.StatusBadRequest,
	ErrTopUpQuantity:      http.StatusBadRequest,
	ErrTransferQuantity:   http.StatusBadRequest,
	ErrPasswordPolicy:     http.StatusBadRequest,
	ErrInsuficientFund:    http.StatusBadRequest,
	ErrCannotTransferSelf: http.StatusBadRequest,
	ErrInsuficientChance:  http.StatusBadRequest,
	ErrBoxIsUnavailable:   http.StatusBadRequest,
	ErrServer:             http.StatusInternalServerError,
	ErrSourceNotFound:     http.StatusBadRequest,
	ErrDateFormat:         http.StatusBadRequest,
	ErrFormat:             http.StatusBadRequest,
}
