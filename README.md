# WebToolKit

Ceci est un modèle d'application web. Il utilise Go pour la partie serveur et Svelte pour la partie cliente. Il peut servir de base pour construire rapidement une application utilisant ces technologies, en recopiant simplement ces fichiers et en lançant les commandes suivantes :

```
# installation des éléments javascript
cd ihm && npm i
```

Pour développer, il faut lancer en parallèle les commandes pour la partie serveur et la partie cliente :

```
# ici on lance le serveur
go run . &

# et ici le client
cd ihm && npm run dev & 
```

Le dossier **ihm** contient le code source pour la partie cliente (Svelte). La construction de cette partie génère les fichiers app.js et app.css dans le sous-dossier "public".

L'architecture côté serveur respecte les principes de séparation des responsabilités :
- **biz** est en charge de la logique métier et est censée être agnostique de la façon d'interagir avec le monde extérieur, que ce soit l'interface web ou le stockage des données. On y trouvera toutes les fonctions purement métiers, que l'on pourrait réutiliser dans d'autres projets.
- **api** regroupe l'ensemble des fonctions en interaction avec l'extérieur (par exemple, l'application Svelte), et est responsable des échanges de données à travers le protocol HTTP. On y trouvera tous les points d'entrées, avec la mécanique pour décoder les demandes (request) et retourner les données en réponses (response au format JSON)
- **bdd** contient les différents magasins de données possibles. Un magasin en mémoire est utilisé en exemple pour sauvegarder l'historique des demandes.
- **sec** centralise la stratégie d'encryption.

## Pré-requis

- nodejs & npm pour développer la partie cliente (javascript et Svelte)
- Go (golang) pour développer la partie serveur

## Reste à faire 

- [x] découpler l'api du biz en mettant une interface listant les fonctions attendues par api 
- [x] montrer comment on gère les données persistantes avec un store en mémoire
- [x] ajoute un middleware côté client pour centraliser tous les accès résseaux et pouvoir catcher les erreurs et les déconnexions
- [ ] ajouter des tests et montrer l'intérêt des interfaces
- [ ] gérer la connexion/déconnexion et les zones publiques/privées
    - [x] étoffer le storage pour stocker les users
    - [x] implémenter une 1ère vérification d'identifiant/mot de passe
    - [x] créer un package séparé pour la sécurité (par ex. "sec")
    - [/] gérer la connexion côté IHM [en cours...]