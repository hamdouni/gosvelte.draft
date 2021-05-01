# WebToolKit

Ceci est un modèle d'application web. Il utilise Go pour la partie serveur et Svelte pour la partie cliente. Il peut servir de base pour construire rapidement une application utilisant ces technologies, en recopiant simplement ces fichiers et en lançant les commandes suivantes :

```
# installation des éléments javascript
npm i
```

Pour développer, il faut lancer en parallèle les commandes pour chaque partie, par exemple dans 2 terminaux séparées, ou en tâche de fond comme suit :

```
npm run dev & go run .
```

Le dossier **ihm** contient le code source pour la partie cliente (Svelte). La construction de cette partie génère les fichiers app.js et app.css dans le sous-dossier "public".

L'architecture côté serveur respecte les principes de séparation des responsabilités :
- **biz** est en charge de votre logique métier et est censée être agnostique de la façon d'interagir avec le monde extérieur, que ce soit l'interface web ou le stockage des données. On y trouvera toutes les fonctions purement métiers, que l'on pourrait réutiliser dans d'autres projets.
- **api** regroupe l'ensemble des fonctions en interaction avec l'extérieur (par exemple, l'application Svelte), et est responsable des échanges de données à travers le protocol HTTP. On y trouvera tous les points d'entrées, avec la mécanique pour décoder les demandes (request) et retourner les données en réponses (response au format JSON)

## Pré-requis

- nodejs & npm pour développer la partie cliente (javascript et Svelte)
- Go (golang) pour développer la partie serveur

## Reste à faire 

- [x] découpler l'api du biz en mettant une interface listant les fonctions attendues par api 
- [ ] montrer comment on gère les données persistantes avec un store en mémoire et un store dans une base de données
- [ ] ajouter des tests et montrer l'intérêt des interfaces