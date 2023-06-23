
# Etape 1 - Mise en place du module et du programme

## Objectif

Cette étape va consister à déclarer 2 types utilisateurs : les type User et Badges.  
A la fin de celle-ci votre programme devrait être en mesure d'afficher des variables  
de type User et Badges.

Voir : [practical-go-lessons.com/chap-13-types#type-struct-variable-creation](https://www.practical-go-lessons.com/chap-13-types#type-struct-variable-creation)


## 10. Déclarer les types struct User et Badge

La struct `Badge` devrait contenir :  

* un champs nom nommé `Name`
* un champs url nommé `Url`  

Tous deux de type `string`

<br>

La struct `User` devrait contenir :  

* un identifiant nommé `ID` de type `uint`
* une liste de badges nommée `Badges`

<br>

## 11. Afficher des variables sur la sortie standard

* Initialiser des variables struct User et Badge et les imprimer

<br>

TODO : fix gap de progression logique sur la prochaine étape dont l'utilité n'est pas évidente

## 12. Gérer un ensemble de valeurs unique stockées dans une slice

* Ajouter un fichier utils.go dans lequel sera ajoutée :

* Une fonction AddToSet prenant en entrée une slice de Badges, et un nouveau Badge,  
  et n'ajoutera ce badge à la slice seulement si elle ne contient pas déja.  
  Cette fonction retournera la slice mise à jour.

<br>

➡️ [Etape 2](../etape_02/README.md)

