package main

import (
	"fmt"

	"badges/types"

	//"badges/datasources/inmemory"
	sqlpg "badges/datasources/sql_pg"
)

func CheckDB() {
	dataSource := sqlpg.DataSource{}

	dataSource.Connect()

	userID, err := dataSource.CreateUser(types.User{})
	if err != nil {
		fmt.Println("CreateUser err:", err)
		return
	}

	fmt.Println("user id:", userID)

	bid, err := dataSource.AddUserBadge(userID, types.Badge{Name: "go", URL: "somewher"})
	if err != nil {
		fmt.Println("CreateBadge err:", err)
		return
	}

	fmt.Println("badge id:", bid)

	bb, err := dataSource.GetUserBadges(userID)
	if err != nil {
		fmt.Println("GetUserBadges err:", err)
		return
	}

	fmt.Printf("user badges: %+v\n", bb)

	if _, err := dataSource.AddUserBadge(userID, types.Badge{Name: "ts", URL: "ts.dev"}); err != nil {
		fmt.Println("AddUserBadge err:", err)
		return	
	}

	bb, err = dataSource.GetUserBadges(userID)
	if err != nil {
		fmt.Println("GetUserBadges err:", err)
		return
	}

	fmt.Printf("user badges after addition : %+v\n", bb)

}
