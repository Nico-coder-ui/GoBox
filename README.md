# GoBox

**Statut du projet : en cours**<br>

GoBox est un outil de type sandbox en ligne de commande léger qui exécute des scripts externes dans des conteneurs Docker isolés et sécurisés.<br>
Il permet différentes statistiques et retour du programme demandé sans pour autant exposer sa machine et son environnement.<br>

## Compilation

```
go build -o gobox
```

## Utilisation

```
./gobox [script] --args="[arguments]"
```

**Exemples**<br>
```
./gobox test.sh
./gobox test.sh --args="val1"
```

## Fonctionnalités
- Isolation via Docker  
- Transmission des arguments au script  
- Capture de la sortie et des erreurs  
- Mesure du temps d’exécution  
- Indication de la réussite ou de l’échec du programme  
- Suppression automatique du conteneur après exécution

## Futurs ajouts
- Possibilité de limiter la mémoire  
- Timeout modifiable  
- Données transférables dans un fichier