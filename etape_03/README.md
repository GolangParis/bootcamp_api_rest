
## Etape 03 - 

* Mise en place de l'utilisation de gin

import "github.com/gin-gonic/gin"

* Création d'un routeur Gin basique

```
ginRouter := gin.Default()
```

* En se basant sur la route existante `/get/users` ajouter ces routes :

```
Méthode  Route                Fonction réponse

GET      /api/badges/:userid  GetUserBadges
POST     /api/badge/:userid   PostBadge
DELETE   /api/badge/:userid   DeleteBadge
PATCH    /api/badge/:userid   UpdateBadge
```

* PostBadge

- décoder le corps de la requête

* DeleteBadge et UpdateBadge

- Décoder le param user ID de l'URL
- Convertir le userID vers le type int avec strconv.ParseInt

