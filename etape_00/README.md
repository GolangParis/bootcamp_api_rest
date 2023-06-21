
# Etape 00 - Mise en place du module et du programme

## Initialisation du projet et du module

Si vous le n'avez pas, et que vous n'avez pas encode adopté d'IDE particulier,
installez Visual Code.

Travailler dans ce répertoire

```
cd badges/etape_00
```

Initialisez le fichier de module

```
go mod init {nom_du_module}
```

nom_du_module = badges (ou autre chose si vous préferez)


## Edition du fichier Go

Avec Visual Code ou votre IDE favori :

```
# avec Visual Code 
code main.go
```

Ajouter la directive de package en tête de fichier

```package main```

Ajouter l'import de la bibliothèque intégrée fmt

Ajouter la fonction main :

Nom de la fonction : main

Son corps devra contenir simplement :

```
fmt.Println("I can haz Badges ?")
```

