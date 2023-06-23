package main

import (
	"fmt"
	"os"
)


var (
	// TODO  map[uint !
	UsersMap = map[string]*User {
		"132654" : &User{"132654", []int{1}},
		"655554" : &User{"655554", []int{2, 3}},
	}

	BadgesMap = map[int]*Badge {
		1 : &Badge{"go",    "https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white"},
		2 : &Badge{"steam", "https://img.shields.io/badge/Steam-000000?style=for-the-badge&logo=steam&logoColor=white"},
		3 : &Badge{"vue",   "https://img.shields.io/badge/Vue.js-35495E?style=for-the-badge&logo=vue.js&logoColor=4FC08D"},
		4 : &Badge{"docker","https://img.shields.io/badge/Docker-00005E?style=for-the-badge&logo=docker&logoColor=4FC08D"},
	}
)

func main() {
	
	// Voir check_db.go pour vérifier la bonne exécution des requêtes SQL
	// *Prérequis* : un serveur Postgres doit être lancé localement
	if len(os.Args) > 1 {
		if os.Args[1] == "checkdb" {
			CheckDB()
			return
		}
	}

	port := "8090"
	fmt.Println("serving on :" + port)

	router := SetupRouter()
	router.Run(":"+port)
}
