# WebToolKit

> Travaux en cours : le projet n'est pas utilisable et nécessite d'importantes modifications. La réflexion en cours est que le modèle logique est trop simpliste pour permettre de mettre en évidence la mécanique générale : une piste serait de développer un tableau de bord d'administrateur avec la gestion des utilisateurs. Ce serait un bon candidat pour une application minimum et réutilisable dans un contexte professionnel.

Ceci est un modèle d'application web. Il utilise Go pour la partie serveur et Svelte pour la partie cliente. Il peut servir de base pour construire rapidement une application utilisant ces technologies, en recopiant simplement ces fichiers et en lançant les commandes suivantes :

```
# installation des éléments javascript
cd cmd/client && npm i
```

Pour développer, il faut lancer en parallèle les commandes pour la partie serveur et la partie cliente :

```
# ici on lance le serveur
go run . &

# et ici le client
cd cmd/client && npm run dev & 
```

Le dossier **client** contient le code source pour la partie cliente (Svelte). La construction de cette partie génère les fichiers app.js et app.css dans le sous-dossier "html".

L'architecture côté serveur respecte les principes de séparation des responsabilités :
- **biz** est en charge de la logique métier et est agnostique de la façon d'interagir avec le monde extérieur, que ce soit l'interface web ou le stockage des données. On y trouvera toutes les fonctions purement métiers, que l'on pourrait réutiliser dans d'autres projets.
- **api** regroupe l'ensemble des fonctions en interaction avec l'extérieur (par exemple, l'application Svelte), et est responsable des échanges de données à travers le protocol HTTP. On y trouvera tous les points d'entrées, avec la mécanique pour décoder les demandes (request) et retourner les données en réponses (response au format JSON)
- **infra** regroupe **bdd** qui contient les différents magasins de données possibles (un magasin en mémoire est utilisé en exemple)  et **sec** pour le chiffrement des données (en exemple, une version AES 256).

## Pré-requis

- nodejs & npm pour développer la partie cliente (javascript et Svelte)
- Go (golang) pour développer la partie serveur

## Reste à faire 

- [x] découple l'api du biz en mettant une interface listant les fonctions attendues par api 
- [x] montre comment on gère les données persistantes avec un store en mémoire
- [x] ajoute un middleware côté client pour centraliser tous les accès réseaux et pouvoir capturer les erreurs et les déconnexions
- [ ] ajoute des tests et montrer l'intérêt des interfaces
- [ ] gére la connexion/déconnexion et les zones publiques/privées
    - [x] étoffe le storage pour stocker les users
    - [x] implémente une 1ère vérification d'identifiant/mot de passe
    - [x] crée un package séparé pour la sécurité (par ex. "sec")
    - [/] gére la connexion côté IHM [en cours...]
        - [x] connexion
        - [x] déconnexion
        - [ ] détection déconnexion backend ?
    - [/] implémente la stratégie d'authentification avec jeton (cf plus bas)
        - [x] enregistre les infos dans le cookie
        - [ ] contrôle à chaque requête sa validité

## Stratégie d'authentification avec jeton (token)

Pour gérer l'authentification des utilisateurs, on leur demande un identifiant et un mot de passe, dont on contrôle la validité. Si les données sont correctes, on génère un jeton crypté qui contient la structure suivante :

```
identifiant|adresse IP|horodatage
```

Le jeton est envoyé au navigateur de l'utilisateur sous la forme d'un 'cookie' : le navigateur va donc nous renvoyer ce 'cookie' à chaque requête. On pourra alors contrôler l'accès légitime de cette requête :

1. on décrypte le contenu du 'cookie'
2. on vérifie que l'adresse IP est toujours celle de connexion
3. on vérifie que cela ne fait pas trop longtemps que la connexion a eu lieu

## Docker

Docker va nous servir à la fois de simulateur d'infrastructure de production et de stratégie de déploiement. 

Nous utilisons Docker pour simuler une architecture avec un reverse-proxy (Traefik) ce qui nous permet de tester la mécanique de récupération de l'adresse IP de l'utilisateur.

Plus tard, nous utiliserons Docker (et Docker Compose) pour organiser notre infrastructure et l'envoyer sur un serveur de production.

## Test de l'API avec curl

```sh
# On s'authentifie en POST et on sauvegarde le cookie dans un fichier
curl -v -c /tmp/cookie.txt -d 'username=maximilien&password=motdepasse' http://localhost:8000/login
# On peut appeler un service en POST en réutilisant le fichier cookie
curl -v -b /tmp/cookie.txt -d 'nom=la%20galaxy' http://localhost:8000/upper
# Ou alors en simple GET (parametres dans l'URL)
curl -v -b /tmp/cookie.txt http://localhost:8000/hello\?nom\=le%20monde
curl -v -b /tmp/cookie.txt http://localhost:8000/lower\?nom\=The%20Universe
curl -v -b /tmp/cookie.txt http://localhost:8000/historic
# On se déconnecte en modifiant le fichier cookie
curl -v -c /tmp/cookie.txt http://localhost:8000/logout
```