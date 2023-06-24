
# Etape 3 -

## Objectif

A la fin de cette étape vous devriez aboutir à une API HTTP répondant à des requêtes REST  
permettant de gérer les badges associés à des identifiants utilisateurs.


## 30. Mise en place de l'utilisation de gin

```go
import "github.com/gin-gonic/gin"
```

Cf. [github.com/gin-gonic/gin](https://github.com/gin-gonic/gin) 

Puis mettre à jour les dépendances à l'endroit où se trouve le module avec :

```sh
go get .
```

&rarr; votre fichier go.mod devrait maintenant mentionner gin-gonic


## 31. Création d'un routeur Gin basique

```go
ginRouter := gin.Default()

// ... setup des routes ...

// Puis finalement lancement du serveur :

ginRouter.Run()

```

## 32. Ajoutez les routes HTTP suivantes a l'API

Pour l'instant ces routes ne font rien et ne retournent pas d'erreurs

|Méthode |  Route               | Fonction réponse| Type du corp de la requête (ou réponse pour GET)        |
| ---    | ---                  | ---             | ---                                                     |
|GET     | /api/badges/:userid  | GetUserBadges   | JSON contenant 1 badge: `{"name": "go", "URL": "xyz"}`  |
|POST    | /api/badge/:userid   | PostBadge       | JSON contenant 1 badge: `{"name": "go", "URL": "xyz"}`  |
|DELETE  | /api/badge/:userid   | DeleteBadge     | JSON avec le nom d'un badge: `{"badgeName": "go"}`      |
|PATCH   | /api/badge/:userid   | UpdateBadge     | JSON contenant 1 badge: `{"name": "go", "URL": "xyz"}`  |


## 33. Utilisation de l'implémentation du stockage des données StorageInMemory

Suite à la précédente étape vous devriez avoir une map de Badges déja déclarée.  
Pour une première implémentation naïve faites en sorte qu'elle soit dans la portée des fonctions  
réponses de l'API, afin que celles-ci puisse y accéder et appeler ses méthodes.


## 34. Implémenter la réponse GetBadge

GetBadge retourne le badge associé à un identifiant utilisateur passé en tant  
que paramètre du chemin URL.

Cf. https://gin-gonic.com/docs/examples/param-in-path

* Extraire l'ID utilisateur du chemin URL
* Convertir l'ID utilisateur vers le type int avec la fonction ParseInt de la lib strconv

* Former une slice contenant les badges associés à l'ID utilisateur.

* Retourner cette slice en tant que membre nommé badges d'un objet JSON :

Cf. : `c.JSON(http.StatusOK, gin.H{ "badges" : bb})`


## 35. La réponse PostBadge

* Extraire l'ID utilisateur du chemin URL
* Convertir l'ID utilisateur vers le type int avec la fonction ParseInt de la lib strconv

* Décoder le corps de la requête avec le [struct binding](https://gin-gonic.com/docs/examples/bind-form-data-request-with-custom-struct/) 

* Ajouter le Badge à l'ID utilisateur avec l'instance de StorageInMemory


## 36. Implémenter la réponse DeleteBadge 

* Extraire l'ID utilisateur du chemin URL
* Convertir l'ID utilisateur vers le type int avec la fonction ParseInt de la lib strconv


## 37. Implémenter la réponse UpdateBadge

* Décoder le param user ID de l'URL
* Convertir le userID vers le type int avec strconv.ParseInt
* Supprimer le badge fournit dans le corp de la requête



<br>

➡️ [Etape 4](../etape_04/README.md)

