
# Bootcam Golang "I can haz Badges ?"

Ce projet est un atelier de programmation en langage Go guidé par étapes progressives,
devant aboutir à une API minimaliste de gestion de badges.

## Prérequis

Il est supposé que : 

* vous avez Go 1.20 installé sur votre machine
* vous savez opérer dans un terminal
* vous savez éditer du code source
* vous avez déja une première expérience avec un langage procédural ou orienté objet
* vous ne portez pas de manteau de fourrure


## badges-api : un microservice CRUD HTTP REST

L'idée de `badges-api` est de fournir un moyen gérer les badges associés à plusieurs  
utilisateurs reposant sur une API REST CRUD minimale.  

C'est quoi un [CRUD](https://developer.mozilla.org/en-US/docs/Glossary/CRUD) ? L'acronyme CRUD correspond à une liste d'opérations habituelles  
quand on manipule des données : Création, Lecture, Modification, Supression.


## A propos des badges

Les badges sont de réelles images générées par un service tiers : [Shields.io](shields.io)

Nous utiliserons des badges statiques ne nécessitant pas de paramètres fonctionnels mais  
tout simplement des paramètres d'apparence :

Exemples d'URL de badges statiques :

* go    : https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white
* steam : https://img.shields.io/badge/Steam-000000?style=for-the-badge&logo=steam&logoColor=white

## Les étapes du projet

Les principales étapes de ce projet :

* Mise en place d'un type utilisateur équipé de méthodes permettant de gérer les badges associés à un utilisateur.

* Mise en place d'une interface Storage permettant de brancher diverses implémentations respectant l'interface

* Implémentation de l'interface Storage dans une base SQL avec `sqlite`

* Implémentation de l'interface Storage avec une base SQL avec `PostgreSQL`

* En cherchant un peu vous trouverez l'étape 42 : elle contient une mini webapp permettant  
de visualiser les badges.


## Les entités que vous serez amené à manipuler

L'API est simple : elle permet de gérer (créer, modifier, accéder, retirer) des badges (type `Badge`) associés à des utilisateurs (type `User`).

Le type `User` a pour membres :

* un simple ID entier

Le type `Badge` a pour membres : 

* un nom (string)
* une URL (string, contenant les paramètres de contenu et de style)


## Vous êtes Go ?

➡️ Rendez-vous sur la page [ETAPES](ETAPES.md)
