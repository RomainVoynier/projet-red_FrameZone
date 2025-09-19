# projet-red_FrameZone
Projet Red Augustin Benoit Romain Voynier Randy Tembe

Character.go implémente le système de création et de gestion d’un personnage pour un jeu de rôle textuel. Le joueur choisit une classe (Chevalier, Archer ou Magicien), chaque classe ayant ses propres points de vie, description et sort spécial. Le personnage possède des caractéristiques comme l’attaque, les HP, un inventaire, des sorts, de l’équipement (tête, torse, pieds) et de l’argent. Des fonctions permettent d’afficher ses infos, d’apprendre des sorts, de gérer l’équipement et de détecter la mort du personnage. Ce système constitue la base du gameplay d’un RPG textuel.

Le code "combat.go" gère le combat tour par tour entre le joueur et un monstre dans un RPG textuel. Le joueur peut attaquer, lancer un sort ou utiliser un objet, tandis que l’adversaire riposte avec une attaque normale ou spéciale. En cas de victoire, le joueur obtient de l’XP, peut monter de niveau et améliorer ses statistiques. Les sorts sont utilisables une seule fois par combat, et les objets consommés sont retirés de l’inventaire.

Equipement.go contrôle l’interface d’achat du forgeron dans un RPG textuel. Le joueur peut y dépenser sa monnaie (Smic) pour acheter des pièces d’équipement (tête, torse, pieds), chacune offrant un bonus de points de vie. En cas d’achat, l’équipement est automatiquement équipé, les HP max et actuels sont ajustés, et l’objet est affiché comme acquis. Si le joueur est trop pauvre, l’achat est refusé. Un menu permet de quitter à tout moment.

Le code "menu.go" constitue le point d’entrée d’un jeu de rôle textuel en Go. Il affiche un prologue narratif, initialise le personnage joueur, puis lance un menu principal interactif. Le joueur peut consulter ses infos, se rendre chez le forgeron, accéder à un combat d'entraînement, visiter un marchand ou quitter le jeu. Le tout est structuré autour d’un système de boucle de jeu textuelle avec des choix numériques.

Le code merchant.go ne marche pas, petit problème du côté de notre troisième coéquipier.

Takepot.go définit une structure de joueur (Joueur) avec ses points de vie et un inventaire. La fonction takePot permet au joueur d'utiliser une potion de soin (si une "potion" est présente dans l'inventaire) sans la retirer. Elle restaure jusqu'à 50 PV sans dépasser le maximum (PVMax), et affiche le résultat à l'écran. (ne fonctionne pas)

upgrademyInventoryslot a également été fait par notre troisième coéquipier, non fonctionnel.

Pour lancer le programme il faut utiliser les commandes suivantes:

"cd Jeu"

"go run character.go combat.go equipement.go menu.go merchant.go takepot.go"
