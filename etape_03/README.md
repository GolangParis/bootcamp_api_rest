
## Etape 3 - 

## Objectif

A la fin de cette étape vous devriez aboutir à 


## 30. Mise en place de l'utilisation de gin

```go
import "github.com/gin-gonic/gin"
```

## 31. Création d'un routeur Gin basique

```go
ginRouter := gin.Default()
```


## 32. En se basant sur la route existante `GET /users` ajouter ces routes :

```sh
Méthode  Route                Fonction réponse

GET      /api/badges/:userid  GetUserBadges
POST     /api/badge/:userid   PostBadge
DELETE   /api/badge/:userid   DeleteBadge
PATCH    /api/badge/:userid   UpdateBadge
```

## 33. PostBadge

* décoder le corps de la requête


## 33. DeleteBadge et UpdateBadge

* Décoder le param user ID de l'URL
* Convertir le userID vers le type int avec strconv.ParseInt

<br>

➡️ [Etape 4](../etape_04/README.md)

