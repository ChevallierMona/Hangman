# Jeu du pendu !

Bienvenue dans le jeu du pendu ! Le but du jeu est de deviner un mot choisi au hasard en proposant des lettres ou en devinant le mot entier.
Si la lettre que le joueur propose n'est pas bonne alors il perd une vie, s'il propose un mot et que le mot n'est pas le mot à deviner alors le joeur perd deux vies, Le nombre de vie est de dix au départ. Le nombre de vie est représnté par une représentation du pendu.

## Fonctionnalités

- **Jouer au jeu du pendu!** : Vous pouvez deviner une lettre ou le mot entier.
- **Sauvegarde et reprise** : Si vous devez interrompre la partie, vous pouvez taper `stop` pour sauvegarder votre progression. Pour reprendre la partie, tapez `go`.
- **Affichage ASCII** : Le mot actuel est affiché en ASCII Art pour une expérience visuelle amusante.
- **Gestion des vies** : Vous commencez avec dix vies. Chaque erreur pour une lettre proposer vous coûte une vies, chaque erreur pour un mot en entier proposer vous coûtera deux vies.

## Prérequis

- Go (version go 1.23.0)
- Un fichier `words.txt` contenant une liste de mots (un mot par ligne).
- Un fichier `standard.txt` pour les représentations ASCII des lettres.

## Installation

1. git clone https://github.com/ChevallierMona/Hangman.git
   cd hangman