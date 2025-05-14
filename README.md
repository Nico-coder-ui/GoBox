🟩 Étape 1 – Initialisation du projet (Jour 1–2)
🎯 Objectif :

    Créer la structure de base du projet Go.

    Préparer l'environnement CLI.

✅ Tâches :

    Initialiser le repo Git + go mod init gobox.

    Créer les dossiers :

/cmd        → CLI
/core       → logique d'exécution
/logs       → gestion des logs
/sandbox    → (placeholder pour V2)
/utils      → helpers

Ajouter une première commande CLI :

    gobox run ./script.sh

    Utiliser une lib comme spf13/cobra pour une CLI propre.

🟨 Étape 2 – Exécution simple de fichier (Jour 3–4)
🎯 Objectif :

    Lancer un script/binaire et capturer stdout, stderr.

✅ Tâches :

    Créer ExecRunner dans /core/runner.go.

    Implémenter :

        lancement du process (exec.Command),

        capture stdout / stderr,

        retour du code de sortie.

    Enregistrer :

        début / fin d’exécution (timestamp),

        PID,

        fichier exécuté,

        répertoire courant.

🟧 Étape 3 – Timeout d’exécution (Jour 5)
🎯 Objectif :

    Arrêter proprement un processus trop long.

✅ Tâches :

    Ajouter un flag --timeout=5s.

    Implémenter un context.WithTimeout() ou un timer dans le runner.

    Si le timeout est dépassé :

        tuer le process,

        logger qu’il a été tué,

        renvoyer un code d’erreur clair.

🟦 Étape 4 – Logs détaillés (Jour 6–7)
🎯 Objectif :

    Enregistrer toutes les infos d’exécution dans un log lisible (JSON ou texte).

✅ Tâches :

    Créer une struct ExecutionLog dans /logs/log.go.

    Stocker :

        stdout / stderr

        code de sortie

        timestamp début/fin

        durée

        PID

        nom du fichier exécuté

    Sauver dans :

        un fichier .json dans /tmp/gobox_logs/

        et/ou afficher avec une option --verbose

🟪 Étape 5 – Limitation mémoire (Jour 8–10)
🎯 Objectif :

    Empêcher un script de consommer trop de mémoire.

✅ Tâches :

    Ajout flag --memory=100MB.

    Pour Linux :

        soit via cgroups (v2 recommandé),

        soit wrapper le processus avec un ulimit si temporairement plus simple.

    Ajouter une alerte/log si la mémoire max est atteinte.

🟥 Étape 6 – Répertoire temporaire isolé (Jour 11)
🎯 Objectif :

    Exécuter le fichier dans un dossier temporaire dédié.

✅ Tâches :

    Créer un dossier os.MkdirTemp() dans /tmp/gobox_XXXX/

    Copier le script dedans si nécessaire.

    Changer Cmd.Dir pour pointer vers ce dossier.

    Supprimer le dossier après exécution sauf si --keep.

⬛ Étape 7 – Structuration & finalisation MVP (Jour 12–14)
🎯 Objectif :

    Nettoyer, documenter, finaliser une release V1 stable.

✅ Tâches :

    Créer un README avec :

        description,

        installation (go install ou binaire),

        exemples de commande (gobox run test.sh --timeout=3s)

    Ajouter :

        un logo ASCII ou image si tu veux styliser,

        un exemple d'exécution loguée.

    Publier la première version GitHub avec une release v0.1.0.

✅ Bonus (optionnels mais utiles)

Ajout d’un flag --logfile=xxx.json pour écrire ailleurs.

Ajout d’une option --dry-run (simule sans exécuter).

Intégration d’un flag --env pour passer des variables personnalisées.

Générer une empreinte SHA256 du fichier lancé.

Affichage CLI stylisé avec spinner, success, error comme une mini UX terminal.