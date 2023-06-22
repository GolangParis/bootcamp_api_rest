
# Etape 01 - Mise en place du module et du programme

* Déclarer les types struct User et Badge
  La struct `Badge` contient: un nom `Name` et un url `Url`, les deux de type `string`.
  La struct `User` contient: un ID `ID` de type `uint` et une liste de badges `Badges`.
* Initialiser des variables struct User et Badge et les imprimer

* Déclarer un type BadgeAPI basé sur  struct User et Badge

* Ajouter un fichier utils.go dans lequel sera ajoutée :
* Une fonction AddToSet prenant en entrée une slice de Badges, et un nouveau Badge,
  et n'ajoutera ce badge à la slice seulement si elle ne contient pas déja.
  Cette fonction retournera la slice mise à jour.

