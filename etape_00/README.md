
# Etape 0 - Mise en place du module et du programme

## Objectif

A la fin de cette étape vous devriez aboutir à un programme Go minimaliste  
affichant un message, et se trouvant à la racine d'un module.


## 00. Choix de l'IDE

Sublime Text, Visual Code, Vim + vim-go, Emacs, nano... pour faire votre premiers pas  
à peu près tout éditeur de texte vous permettra de coder en Go.

Le plus important est qu'une extension Go soit prévue pour votre éditeur, afin qu'il  
puisse au moins lancer l'outil de formattage de code `go fmt`.

Si vous ne savez pas lequel choisir et voulez un outil répandu, Visual Code est ok  
pour démarrer. Attention, il n'est pas sans surprises.


## 01. Initialisation du projet et du module

Le bootcamp démarre dans le répertoire etape_00 où vous allez "initialiser"  
un nouveau projet, qui revient tout simplement à créer un fichier `go.mod`  
pour déclarer un module :

```sh
cd badges/etape_00

# Initialisez le fichier de module

go mod init {nom_du_module}
```

nom_du_module = badges (ou autre chose si vous préferez)


## 02. Edition du fichier Go contenant le point d'entrée

Le point d'entrée est l'endroit où votre programme va démarrer.  
En Go, c'est la fonction main(). Il ne doit y avoir qu'une seule  
fonction main dans un programme :

Avec Visual Code ou votre IDE favori :

```sh
# avec Visual Code
code main.go
```

Note : le nom main.go n'est pas une obligation.

## 03. Ajouter la directive de package en tête de fichier

```go
package main
```

## 04. Ajouter l'import de la bibliothèque intégrée fmt

```go
import "fmt"
```

## 05. Ajouter la fonction main

Une fonction simple se déclare comme suit :

```go
func Operation() {

}
```

Cf. [practical-go-lessons.com/chap-10-functions](https://www.practical-go-lessons.com/chap-10-functions)

Il vous faut simplement ajouter une fonction nommée `main` sans paramètre ni code de retour.


## 06. Ajouter un affichage de message sur la sortie standard

Son corps devra contenir simplement :

```golang
fmt.Println("I can haz Badges ?")
```

## 07. Construire le binaire

```sh
$ go build
```

## 08. Lancer le binaire produit

Le nom du binaire est celui du module :

```sh
./badges
```

➡️ [Etape 1](etape_01/README.md)

