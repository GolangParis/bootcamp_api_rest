
## Etape 2 - 

## Objectif

A la fin de cette étape vous devriez aboutir à un programme permettant de gérer des badges associés à  
des identifiants utilisateurs.

C-a-d pour un identifiant utilisateur donné : il devrait être possible de lister ses badges, et de lui ajouter,  modifier, ou retirer un badge.

Ressources utiles :

* [practical-go-lessons.com/chap-22-maps](https://www.practical-go-lessons.com/chap-22-maps)

## 20. Rassembler des identifiants de Badges en une liste avec le type slice

Dans un premier temps, pour vous familiariser avec les slices, déclarez en une contenant des int.  
Et imprimez les élements de la slice avec fmt.Println

Cf. : [practical-go-lessons.com/chap-21-slices](https://www.practical-go-lessons.com/chap-21-slices)

## 21. Déclarer une map associant des identifiants de Badges à des Badges

Déclarez une variable de type map nommée BadgesMap associant à un `int` une instance de type Badge.


## 22. Associer des Badges à un ID utilisateur avec une map

Ensuite on va vouloir associer une slice d'identifiants de Badges à un identifiant utilisateur.

Pour ce faire, déclarez une variable de type map associant à un `int` une slice d'int.


## 22. Définir un type StorageInMemory basé sur ce type map

Pour manipuler une slice de badges, nous allons avoir besoin de plusieurs méthodes et  
nous voudrons les rattacher à un type donné.

Commençons par définir un type utilisateur basé sur notre map :

```go
type StorageInMemory map[int][]int
```

La clé de type `int` représente un User ID. La valeur de type `[]int` est une liste de Badge IDs.

Déclarez ensuite une instance de ce type, ajoutez lui des utilisateurs, et affichez les  
ensuite sur la sortie standard.


## 23. Equiper StorageInMemory de méthodes pour gérer les badges

La gestion des utilisateurs, l'accès aux badges d'un utilisateur, l'ajout, la mise à jour  
et la suppression de badge va se faire par le biais de ces méthodes  :


```go
Connect() error

CreateUser(user User) (User, error)
GetUsers() ([]User, error)

GetUserBadges(userID int) ([]Badge, error)
AddUserBadge(userID int, b Badge) (int, error)
UpdateUserBadge(userID int, b Badge) error
DeleteUserBadge(userID int, badgeName string) error
```

Equipez le type StorageInMemory de ces méthodes et implémentez les.

Cf. [practical-go-lessons.com/chap-14-methods](https://www.practical-go-lessons.com/chap-14-methods)


## 24. Vérifier quelques fonctionnements de base

Ajoutez un ou plusieurs utilisateur à l'instance de StorageInMemory avec CreateUser.  
Appeler GetUsers et imprimez le résultat obbtenu sur la sortie standard.

&rarr; vous devriez obtenir les utilisateurs précédemment créés.

## 25. Lancer le test unitaire

```
go test
```

<br>

➡️ [Etape 3](../etape_03/README.md)

