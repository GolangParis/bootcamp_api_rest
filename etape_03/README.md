
# Etape 3 -

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

## 32. Ajoutez les routes HTTP suivantes a l'API

Pour l'instant ces routes ne font rien et ne retournent pas d'erreurs

|Méthode |  Route               | Fonction réponse| Type du corp de la requête (ou réponse pour GET)        |
| ---    | ---                  | ---             | ---                                                     |
|GET     | /api/badges/:userid  | GetUserBadges   | JSON contenant 1 badge: `{"name": "go", "URL": "xyz"}`  |
|POST    | /api/badge/:userid   | PostBadge       | JSON contenant 1 badge: `{"name": "go", "URL": "xyz"}`  |
|DELETE  | /api/badge/:userid   | DeleteBadge     | JSON avec le nom d'un badge: `{"badgeName": "go"}`      |
|PATCH   | /api/badge/:userid   | UpdateBadge     | JSON contenant 1 badge: `{"name": "go", "URL": "xyz"}`  |

Déclarez une map de Badges stoquée en mémoire (type `StorageInMemory`) en variable
globale afin de stoquer le contenue reçue via l'API

## 33. PostBadge

* Convertir le userID vers le type int avec strconv.ParseInt
* décoder le corps de la requête
* stoquer le corps de la requête dans la base en mémoire

## 33. DeleteBadge et UpdateBadge

* Décoder le param user ID de l'URL
* Convertir le userID vers le type int avec strconv.ParseInt
* supprimer le badge fournit dans le corp de la requête

## 34. GetBadge

* Décoder le param user ID de l'URL
* Retourner la liste de gadges ascociés a un user ID.

<br>

➡️ [Etape 4](../etape_04/README.md)
