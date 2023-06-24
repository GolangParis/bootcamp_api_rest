
# Etape 1 - Mise en place du module et du programme

## Objectif

Cette étape va consister à déclarer 2 types utilisateurs : les type User et Badges.  
A la fin de celle-ci votre programme devrait être en mesure d'afficher des variables  
de type User et Badges.

Voir : [practical-go-lessons.com/chap-13-types#type-struct-variable-creation](https://www.practical-go-lessons.com/chap-13-types#type-struct-variable-creation)


## 10. Déclarer les types struct User et Badge

La struct `Badge` devrait contenir :  

* un champs nom nommé `Name`
* un champs url nommé `URL`  

Tous deux de type `string`

<br>

La struct `User` devrait contenir :  

* un identifiant nommé `ID` de type `uint`
* une liste (slice) d'identifiants de badges nommée `BadgeIDs`

<br>

## 11. Afficher des variables sur la sortie standard

* Initialiser des variables struct User et Badge et les imprimer

<br>

## 12. Gérer un ensemble de valeurs uniques stockées dans une slice

* Ajouter un fichier utils.go dans lequel sera ajoutée :

* Une fonction AddToSet prenant en entrée une slice d'identifiant entiers, et un nouvel identifiant,  
  et n'ajoutera cet identifiant à la slice seulement si elle ne le contient pas déja.  
  Cette fonction retournera la slice mise à jour.

<br>

➡️ [Etape 2](../etape_02/README.md)

