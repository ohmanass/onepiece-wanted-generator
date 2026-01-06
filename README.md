# 🏴‍☠️ One Piece – Wanted Generator (Go)

Un générateur d’avis de recherche (PDF) pour les pirates de l’univers **One Piece**, développé en **Golang**.  
Ce projet transforme une liste de pirates en fichiers PDF WANTED, avec leur nom, leur prime et leur image.

---

## 🎯 Objectif du projet

- Automatiser la création des avis de recherche pour les pirates.  
- Générer des PDFs “WANTED” stylés avec image et prime.  
- Tester le parsing CSV et la génération PDF en Go.  

---

## 🛠️ Fonctionnalités

- ✅ Lecture d’un fichier CSV listant les pirates  
- ✅ Validation du nom des pirates (doit être en majuscules)  
- ✅ Génération de PDF avec :
  - Titre **WANTED**
  - Nom du pirate
  - Prime affichée en Berrys
  - Image du pirate
  - Fond “Wanted” optionnel
- ✅ Projet structuré en Go simple : `main.go` + `pirate.go`  
- ✅ Compatible avec Go modules (`go.mod`)

---

## 🚀 Installation et exécution

1. Cloner le projet :

```bash
git clone https://github.com/TON_PSEUDO/onepiece-wanted-generator.git
cd onepiece-wanted-generator
```

2. Installer les dépendances :

```bash
go mod tidy
```

3. Lancer le programme :

```bash
go run . 
```