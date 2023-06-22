# WIP Bootcamp Golang Paris API REST

## Proposition de fonctionnement

Au contraire des précédents éditions guidées par des présentations théoriques, on 
pourrait cette fois-ci conduire la progression par de la découverte code, commentée oralement, 
avec des liens pour compléments d'information, et inviter les participants à réparer les tests 
en corrigeant le code existant et/ou en implémentant le code nécessaire.

Après avoir passé en revue avec le groupe un helloworld présentant le fonctionnement général 
du langage (Difficulté ici : où s'arrête-t-on ?) on invite les participants à git cloner ce projet, 
on leur explique le fonctionnement de go test, et on leur suggère un parcours, spécifié dans un 
fichier README et/ou par l'ordre des tests unitaires.


## La thématique

Il serait bon que la thématique du sujet résonne avec l'actualité, une tendance en cours, 
ou soit foncièrement stupide, et donc amusante. (NDLR : tout le monde n'est pas fan de l'humour  
absurde).

Vu le temps imparti : proposition d'associer à des user ID des badges fournit par [Shields](https://shields.io)


## Le projet badges-api

L'idée est de gérer les badges de plusieurs utilisateurs avec une API REST CRUD minimale.  
CRUD : Création, Lecture, Modification, Supression.

Une API REST de type CRUD, manipulant du JSON, stocké initialement en RAM dans une map.

Evolutions possibles ensuite :

* Mise en place d'une interface Storage permettant de brancher diverses implémentation respectant  
l'interface

* Implémentation dans une base NoSQL  type BoltDB

* Implémentation dans une base SQL avec PostGres


## Le parcours badges-api proposé aux participants

badges-api est un microservice (API REST CRUD) de gestion de badges

Exemples d'URL de badges 

* go    : https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white
* steam : https://img.shields.io/badge/Steam-000000?style=for-the-badge&logo=steam&logoColor=white
* vue   : https://img.shields.io/badge/Vue.js-35495E?style=for-the-badge&logo=vue.js&logoColor=4FC08D

### L'objectif général

TODO


## Les entités que l'on va manipuler

Un utilisateur a :
* un simple ID entier

Un badge a :
* un nom (string)
* une URL avec des caractéristiques de contenu et de style

Exemple d'URL de badge : 

TODO


## Progression étape par étape

## Mise en place du programme

Cf. [Etape 00](etape_00/README.md)

## Etape 01 : Types et fonctions

Cf. [Etape 01](etape_01/README.md)

## Etape 02 : Stockage de données en mémoire

Cf. [Etape 02](etape_02/README.md)


## Etape 03 : Mise en place de l'API HTTP avec Gin

Cf. [Etape 03](etape_03/README.md)


## Etape 04 : Abstraction des opérations de stockage de données

Cf. [Etape 04](etape_04/README.md)


## Points à statuer & discussions

Cf. les [issues](https://github.com/GolangParis/bootcamp_api_rest/issues)
