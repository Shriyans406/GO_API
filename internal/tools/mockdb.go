package tools

import{
	"time"
}

type mockDB struct{}

var mockLoginDeatils=map[string]loginDetails{
	"alex": {
		AuthToken: : "123ABC",
		Username: "alex",
	}
	"jason": {
		AuthToken: : "456DEF",
		Username: "jason",
	}
	"marie": {
		AuthToken: : "789GHI",
		Username: "marie",
	},
}

var mockCoinDetails=map[string]CoinDetails{
	"alex":{
		Coins: 100,
		Username: "alex",
	},
	"jason":{
		Coins: 200,
		Username: "jason",
	},
	"marie":{
		Coins: 300,
		Username: "marie",
	},
}

func(d *mockDB) GetUserLoginDetails(username string) *loginDetails {
	time.Sleep(time.Second*1)

	var clientData=LoginDetails{}
	if !ok{
		return nil
	}

	return &clientData
}

func(d *mockDB) GetUserCoins(username string) *CoinDetails {
	time.Sleep(time.Second*1)

	var clientData=CoinDetails{}
	if !ok{
		return nil
	}
	return &clientData
}

func(d *mockDB) SetupDatabase() error{
	time.Sleep(time.Second*1)
	return nil
}