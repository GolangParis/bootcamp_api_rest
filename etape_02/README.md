
## Etape 02 - 

* Déclarer une map associant à un `uint` une slice contenant 0 ou plusieurs Badges

* Définir un type StorageInMemory basé sur ce type map

type StorageInMemory map[uint][]Badge


* Equiper StorageInMemory de méthodes pour gérer les badges :

```
Connect()

CreateUser(types.User) (uint, error)
	
GetUsers() ([]types.User, error)

GetUserBadges(userID uint) ([]types.Badge, error)
AddUserBadge(userID uint, b types.Badge) (uint, error)
UpdateUserBadge(userID uint, b types.Badge) error
DeleteUserBadge(userID uint, badgeName string) error
```

Cf. [practical-go-lessons.com/chap-14-methods](https://www.practical-go-lessons.com/chap-14-methods)

