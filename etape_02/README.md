
## Etape 2 - 

## Objectif

A la fin de cette étape vous devriez aboutir à

## 20. Déclarer une map associant à un `uint` une slice contenant 0 ou plusieurs Badges

## 21. Définir un type StorageInMemory basé sur ce type map

```go
type StorageInMemory map[uint][]Badge
```


## 22. Equiper StorageInMemory de méthodes pour gérer les badges :

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


## 23. Vérifier quelques fonctionnements de base

Créer un ou plusieurs users avec CreateUser
Appeler GetUsers et imprimer le résultat obbtenu sur la sortie standard

=> vous devriez obtenir les utilisateurs précédemment créés.

## 24. Lancer le test unitaire

```
go test
```

<br>

➡️ [Etape 3](../etape_03/README.md)

