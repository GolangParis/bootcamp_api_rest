
# Etape 01 - Mise en place du module et du programme

## Déclarer les types struct User et Badge

La struct `Badge` devrait contenir :  

* un champs nom `Name`
* un champs url `Url`  
(tous deux de type `string`)

<br>

La struct `User` devrait contenir :  

* un ID `ID` de type `uint`
* une liste de badges `Badges`

<br>

## Afficher des variables sur la sortie standard

* Initialiser des variables struct User et Badge et les imprimer

<br>


## Gérer un ensemble de valeurs unique stockées dans une slice

* Ajouter un fichier utils.go dans lequel sera ajoutée :

* Une fonction AddToSet prenant en entrée une slice de Badges, et un nouveau Badge,  
  et n'ajoutera ce badge à la slice seulement si elle ne contient pas déja.  
  Cette fonction retournera la slice mise à jour.

