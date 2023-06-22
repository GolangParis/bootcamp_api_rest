
## Etape 02 - 

* Déclarer une map associant à un `uint` une slice contenant 0 ou plusieurs Badges

* Définir un type StorageInMemory basé sur ce type map

type StorageInMemory map[uint][]Badge


* Equiper StorageInMemory de méthodes pour gérer les badges :

```go
Connect() error

CreateUser(User) (User, error)

GetUsers() ([]User, error)

GetUserBadges(userID uint) ([]Badge, error)
AddUserBadge(userID uint, b Badge) (uint, error)
UpdateUserBadge(userID uint, b Badge) error
DeleteUserBadge(userID uint, badgeName string) error
```

Cf. [practical-go-lessons.com/chap-14-methods](https://www.practical-go-lessons.com/chap-14-methods)


* Vérifier quelques fonctionnements de base

Créer un ou plusieurs users avec CreateUser
Appeler GetUsers et imprimer le résultat obbtenu sur la sortie standard

=> vous devriez obtenir les utilisateurs précédemment créés.

