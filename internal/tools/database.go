package tools

import (
	log "github.com/sirupsen/logrus"
)
type loginDetails struct{
	AuthToken string
	Username string
}
type CoinDetails struct{
	Coins int64
	Username string
}

type DatabaseInterface struct{
	Get UserLoginDetails(username string) *LoginDetails
	GetUserCoins(username string) *CoinDetails
	SetupDatabase() error
}

func NewDatabase() (*DatabaseInterface, error){
	var database DatabaseInterface = &mockDB{}

	var err error=database.SetupDatabase()
	if err!=nil{
		log.Error(err)
		return nil, err
	}

	return &database, nil
}