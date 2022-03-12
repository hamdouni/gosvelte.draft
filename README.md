# WebToolKit

Ceci est un modèle d'application web en Go pour la partie serveur et Svelte pour la partie cliente, avec Docker Compose comme orchestrateur. Il peut servir de base pour construire rapidement une application utilisant ces technologies.

## Pré-requis

- Go
- Docker et Docker Compose
- npm

## Installation

Pour la partie cliente Svelte uniquement :

```
cd client && npm i && cd -
```

## Utilisation

Le client et le serveur sont automatiquement recompilés à chaque modification de leur code respectif.

```
docker-compose up -d && cd client && npm run dev & cd -
```

Le dossier **client** contient le code source pour la partie cliente (Svelte). La construction de cette partie génère les fichiers app.js et app.css dans le sous-dossier "static".

L'architecture côté serveur respecte les principes de séparation des responsabilités :

- **model** est en charge de la logique métier et de la structuration des données. On y trouve toutes les fonctions purement métiers, que l'on pourrait réutiliser dans d'autres projets. 
Des interfaces sont définies pour intéragir avec le stockage des données et une injection de cette dépendance est nécessaire au démarrage. Une implémentation en RAM est fournie dans le paquet `store/ram`.
Cette couche internalise les fonctions de sécurité dans le sous paquet `secure`. 

- **api** regroupe l'ensemble des fonctions en interaction avec l'extérieur (par exemple, l'application Svelte), et est responsable des échanges de données à travers le protocol HTTP. On y trouvera tous les points d'entrées, avec la mécanique pour décoder les demandes (request), retourner les données en réponses (response au format JSON) et utilise le chiffrement des données (cookie et hash mot de passe) mis à disposition par le métier.

- **store** contient les différents magasins de données possibles. Un magasin en mémoire (ram) est utilisé en exemple et permet de remplir le contrat d'interface du métier pour le stockage des utilisateurs et de l'historique.

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
    - [x] implémente la stratégie d'authentification avec jeton (cf plus bas)
        - [x] enregistre les infos dans le cookie
        - [x] contrôle à chaque requête sa validité
- [ ] ajoute un Makefile pour faciliter l'installation et le lancement

## Stratégie d'authentification avec jeton (token)

Pour gérer l'authentification des utilisateurs, on leur demande un identifiant et un mot de passe, dont on contrôle la validité. Si les données sont correctes, on génère un jeton signé et crypté qui contient la structure suivante :

```
identifiant|adresse IP|horodatage
```

Le jeton est envoyé au navigateur de l'utilisateur sous la forme d'un 'cookie' : le navigateur va donc nous renvoyer ce 'cookie' à chaque requête. On pourra alors contrôler l'accès légitime de cette requête :

1. on décrypte le contenu du 'cookie' : si on y arrive c'est que la signature est bonne.
2. on vérifie que l'adresse IP est toujours celle de connexion
3. on vérifie que cela ne fait pas trop longtemps que la connexion a eu lieu

## Docker

Docker va nous servir à la fois de simulateur d'infrastructure de production et de stratégie de déploiement. 

## Test de l'API avec curl

```sh
# On s'authentifie en POST et on sauvegarde le cookie dans un fichier
curl -v -c /tmp/cookie.txt -d 'username=test&password=test' http://localhost:80/login
# On peut appeler un service en POST en réutilisant le fichier cookie
curl -v -b /tmp/cookie.txt -d 'nom=la%20galaxy' http://localhost:80/upper
# Ou alors en simple GET (parametres dans l'URL)
curl -v -b /tmp/cookie.txt http://localhost:80/hello\?nom\=le%20monde
curl -v -b /tmp/cookie.txt http://localhost:80/lower\?nom\=The%20Universe
curl -v -b /tmp/cookie.txt http://localhost:80/historic
# On se déconnecte en modifiant le fichier cookie
curl -v -c /tmp/cookie.txt http://localhost:80/logout
```