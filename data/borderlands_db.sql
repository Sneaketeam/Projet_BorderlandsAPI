-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Hôte : 127.0.0.1
-- Généré le : dim. 23 nov. 2025 à 05:04
-- Version du serveur : 10.4.32-MariaDB
-- Version de PHP : 8.2.12

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Base de données : `borderlands_db`
--

-- --------------------------------------------------------

--
-- Structure de la table `weapons`
--

CREATE TABLE `weapons` (
  `id` int(11) NOT NULL,
  `category` varchar(50) DEFAULT NULL,
  `name` varchar(100) DEFAULT NULL,
  `manufacturer` varchar(50) DEFAULT NULL,
  `rarity` varchar(50) DEFAULT NULL,
  `flavor_text` text DEFAULT NULL,
  `details` text DEFAULT NULL,
  `source` varchar(100) DEFAULT NULL,
  `image_url` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Déchargement des données de la table `weapons`
--

INSERT INTO `weapons` (`id`, `category`, `name`, `manufacturer`, `rarity`, `flavor_text`, `details`, `source`, `image_url`) VALUES
(1, 'Pistolet', 'Unkempt Harold', 'Torgue', 'Légendaire', 'Did I fire six shots...?', 'Tire une rangée de projectiles explosifs.', 'Savage Lee', '/images/harold.jpg'),
(2, 'Fusil', 'KerBlaster', 'Torgue', 'Légendaire', 'Torgue got more BOOM!', 'Tire une roquette qui libère une grenade.', 'Midge-Mong', '/images/kerblaster.jpg'),
(3, 'Bouclier', 'The Bee', 'Hyperion', 'Légendaire', 'Float like a butterfly...', 'Dégâts d\'amplification énormes.', 'Hunter Hellquist', '/images/the_bee.jpg'),
(4, 'Pistolet', 'Hornet', 'Dahl', 'Légendaire', 'Fear the swarm!', 'Toujours corrosif. Tire une rafale de projectiles acides.', 'Knuckle Dragger', '/images/hornet.jpg'),
(5, 'Pistolet', 'Gunerang', 'Tediore', 'Légendaire', 'Give it a burl!', 'Se comporte comme un boomerang au rechargement.', 'Rakkman', '/images/gunerang.jpg'),
(6, 'Pistolet', 'Hector\'s Paradise', 'Dahl', 'Légendaire', 'High tech. Low life.', 'Dégâts électriques élevés, très précis.', 'Hector (DLC Lilith)', '/images/hectors_paradise.jpg'),
(7, 'Pistolet', 'Stalker', 'Vladof', 'Perlé', 'You can run, but you can\'t hide.', 'Les balles rebondissent sur les murs.', 'Butin Mondial (UVHM)', '/images/stalker.jpg'),
(8, 'Pistolet', 'Unforgiven', 'Jakobs', 'Perlé', 'It\'s a helluva thing...', 'Dégâts critiques massifs, cadence très lente.', 'Butin Mondial (UVHM)', '/images/unforgiven.jpg'),
(9, 'Pistolet', 'Wanderlust', 'Maliwan', 'Perlé', 'I have no idea where this is going.', 'Les balles cherchent les ennemis automatiquement (homing).', 'Tubby Enemies', '/images/wanderlust.jpg'),
(10, 'Pistolet', 'Devastator', 'Torgue', 'Séraphin', 'Hail.', 'Tire deux lignes de gyrojets supplémentaires.', 'Master Gee', '/images/devastator.jpg'),
(11, 'Pistolet', 'Infection', 'Maliwan', 'Séraphin', 'It\'s definitely lupus.', 'Dégâts corrosifs sur la durée très élevés.', 'Seraph Vendor', '/images/infection.jpg'),
(12, 'Pistolet', 'Stinger', 'Vladof', 'Séraphin', 'It\'s like an angry angry bee!', 'Les balles ricochent. Cadence de tir extrême.', 'Dragons of Destruction', '/images/stinger.jpg'),
(13, 'Fusil d\'assaut', 'Madhous!', 'Bandit', 'Légendaire', 'It\'s a Madhouse! A MADHOUSE!!!', 'Les balles partent dans tous les sens et ricochent.', 'Mad Dog', '/images/madhous.jpg'),
(14, 'Fusil d\'assaut', 'Ogre', 'Torgue', 'Légendaire', 'Ogres chew their food.', 'Tire des explosifs aléatoires, bonus de vitesse de tir.', 'Warlord Slog', '/images/ogre.jpg'),
(15, 'Fusil d\'assaut', 'M2828 Thumpson', 'Jakobs', 'Légendaire', 'I can do this all day.', 'Gros chargeur, tire très vite pour du Jakobs.', 'Graves (DLC Lilith)', '/images/thumpson.jpg'),
(16, 'Fusil d\'assaut', 'Sawbar', 'Bandit', 'Perlé', 'Suppressing Fires!!!', 'Les balles se divisent en trois boules de feu après une certaine distance.', 'Butin Mondial (UVHM)', '/images/sawbar.jpg'),
(17, 'Fusil d\'assaut', 'Bearcat', 'Dahl', 'Perlé', 'I love the smell of popcorn...', 'Tire des grenades qui rebondissent. Consomme beaucoup de munitions.', 'Butin Mondial (UVHM)', '/images/bearcat.jpg'),
(18, 'Fusil d\'assaut', 'Bekah', 'Jakobs', 'Perlé', 'Shoot them in the face.', 'La balle se divise en 3 après un certain temps. Dégâts énormes.', 'Tubby Enemies', '/images/bekah.jpg'),
(19, 'Fusil d\'assaut', 'Lead Storm', 'Vladof', 'Séraphin', 'Say \"Ow\".', 'Les balles se divisent en deux en cloche.', 'Seraph Vendor', '/images/lead_storm.jpg'),
(20, 'Fusil d\'assaut', 'Seraphim', 'Dahl', 'Séraphin', 'Angels would be proud.', 'Tire des projectiles incendiaires lents mais puissants.', 'Seraph Vendor', '/images/seraphim.jpg'),
(21, 'Fusil d\'assaut', 'Seeker', 'Torgue', 'Séraphin', 'Get \'em, boy!', 'Tire des missiles guidés vers les ennemis.', 'Seraph Vendor', '/images/seeker.jpg'),
(22, 'Bouclier', 'The Transformer', 'Vladof', 'Légendaire', 'More than your eye can see.', 'Immunisé à l\'électricité, recharge le bouclier avec les dégâts électriques.', 'Pimon', '/images/transformer.jpg'),
(23, 'Bouclier', 'Whisky Tango Foxtrot', 'Dahl', 'Légendaire', 'Situation normal...', 'Lâche des boosters explosifs qui blessent tout le monde.', 'Chubby Enemies', '/images/wtf_shield.jpg'),
(24, 'Bouclier', 'Fabled Tortoise', 'Pangolin', 'Légendaire', 'Win by a hare.', 'Énorme capacité, mais réduit la vitesse de déplacement.', 'Blue', '/images/tortoise.jpg'),
(25, 'Bouclier', 'Impaler', 'Vladof', 'Légendaire', 'Vlad would be proud.', 'Lance des piques corrosives sur les ennemis qui vous tirent dessus.', 'The Warrior', '/images/impaler.jpg'),
(26, 'Bouclier', 'The Cradle', 'Tediore', 'Légendaire', '...to the grave.', 'Le bouclier est jeté comme une grenade quand il est vide.', 'Henry', '/images/cradle.jpg'),
(27, 'Bouclier', 'Hide of Terramorphous', 'Bandit', 'Légendaire', 'His hide turned the mightiest tame...', 'Dégâts de feu, Nova de feu et Dégâts au corps à corps.', 'Terramorphous', '/images/hide_of_terra.jpg'),
(28, 'Bouclier', 'Antagonist', 'Maliwan', 'Séraphin', 'I\'m rubber, you\'re glue.', 'Renvoie les balles ennemies et lance des boules de slag.', 'Seraph Vendor', '/images/antagonist.jpg'),
(29, 'Bouclier', 'Blockade', 'Anshin', 'Séraphin', 'Am I preventing you from something?', 'Réduit les dégâts subis tant que le bouclier est actif.', 'Dragons of Destruction', '/images/blockade.jpg'),
(30, 'Bouclier', 'Big Boom Blaster', 'Torgue', 'Séraphin', 'For all your big boom needs.', 'Lâche des boosters qui redonnent du bouclier et des grenades.', 'Pyro Pete', '/images/big_boom_blaster.jpg');

--
-- Index pour les tables déchargées
--

--
-- Index pour la table `weapons`
--
ALTER TABLE `weapons`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT pour les tables déchargées
--

--
-- AUTO_INCREMENT pour la table `weapons`
--
ALTER TABLE `weapons`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=31;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
