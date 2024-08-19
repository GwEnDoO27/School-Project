
# Push_swap

A brief description of what this project does and who it's for

Push Swap

Push Swap est un programme simple de tri par manipulation d'une pile. Le but est de trier une pile d'entiers en utilisant un ensemble d'instructions prédéfinies, tout en utilisant une deuxième pile comme espace temporaire.
Table des matières

    Introduction
    Instructions
    Compilation
    Utilisation
    Exemples
    Auteur
    Licence

Introduction

Le projet Push Swap consiste à trier une pile d'entiers en utilisant deux piles et un ensemble d'instructions restreint. Le but est de minimiser le nombre d'instructions nécessaires pour trier la pile initiale.
Instructions

Les instructions autorisées sont les suivantes :

    sa : swap a - échanger les deux premiers éléments de la pile A.
    sb : swap b - échanger les deux premiers éléments de la pile B.
    ss : sa et sb en même temps.
    pa : push a - déplacer le premier élément de la pile B vers la pile A.
    pb : push b - déplacer le premier élément de la pile A vers la pile B.
    ra : rotate a - déplacer le premier élément de la pile A à la fin.
    rb : rotate b - déplacer le premier élément de la pile B à la fin.
    rr : ra et rb en même temps.
    rra : reverse rotate a - déplacer le dernier élément de la pile A au début.
    rrb : reverse rotate b - déplacer le dernier élément de la pile B au début.
    rrr : rra et rrb en même temps.

Compilation

Pour compiler le programme, utilisez la commande make :

cd push-swap/
go build -o push-swap

cd ../checker/
go build -o checker


(Déplacer l'executable checker dans le dossier où se trouve l'executable 'push-swap')
Cela générera l'exécutable "push_swap".
Utilisation

L'exécution du programme se fait comme suit :

bash

./push_swap [liste d'entiers]
echo -e [liste des fonctions de modifications de listes] | ./checker [liste d'entiers]
ARG=[liste d'entier]; ./push-swap "$ARG" | ./checker "$ARG".

Exemples

bash

./push_swap 3 1 4 2
echo -e "pb\nra\npb\nra\nsa\nra\npa\npa\n" | ./checker "0 9 1 8 2"
ARG="4 67 3 87 23"; ./push-swap "$ARG" | ./checker "$ARG"