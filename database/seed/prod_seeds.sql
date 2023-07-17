-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: arcadia_23_db
-- Generation Time: Mar 14, 2023 at 10:28 PM
-- Server version: 8.0.31
-- PHP Version: 8.0.26

SET FOREIGN_KEY_CHECKS=0;
SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `arcadia_23`
--

--
-- Dumping data for table `characters`
--

INSERT INTO `characters` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `description`, `image_url`) VALUES
(1, '2023-03-04 13:14:01.195', '2023-03-04 13:14:01.195', NULL, 'Adventurer', 'Adventurer is a brave and determined explorer driven by a thirst for knowledge and discovery. Adventurer roams the map with a keen eye for the elusive minicons. Adventurer can often overcome seemingly insurmountable obstacles with a quick wit and strategic mind.', 'adventurer.webp'),
(2, '2023-03-04 13:14:01.197', '2023-03-04 13:14:01.197', NULL, 'Bandit', 'Bandit is stealthy and cunning, deeply connected to the art of thievery and acquiring valuable treasures. With a mask to hide their identity, Bandit moves through the map with purpose in search of the priceless treasures that lie hidden throughout the terrain.', 'bandit.webp'),
(3, '2023-03-04 13:14:01.200', '2023-03-04 13:14:01.200', NULL, 'Pirate', 'Pirate is a cunning and elusive voyager with a fierce reputation that precedes him. Dressed in traditional pirate garb, complete with a pirate hat and a mask, Pirate moves stealthily through the map in search of minicons. Using his keen senses and strategic thinking, Pirate can navigate the map with ease.', 'pirate.webp'),
(4, '2023-03-04 13:14:01.201', '2023-03-04 13:14:01.201', NULL, 'Prince', 'Prince is a noble and regal pioneer, with a confident demeanour and a sharp eye for detail. With his striking blue hair and elegant crown, Prince moves through the map looking for more treasures to fill his coffers.', 'prince.webp'),
(5, '2023-03-04 13:14:01.203', '2023-03-04 13:14:01.203', NULL, 'Robin', 'Robin is quirky and energetic, with an insatiable curiosity and a knack for finding hidden treasures. Wearing a distinctive green hat and sporting a comical moustache, Robin flits about the map with lightning-fast reflexes and a sharp eye for detail.', 'robin.webp'),
(6, '2023-03-04 13:14:01.205', '2023-03-04 13:14:01.205', NULL, 'Santa', 'Santa is jolly and cheerful, with a heart full of holiday spirit and a knack for finding hidden treasures. Wearing his traditional red and white suit, complete with a big black belt and fluffy white beard, Santa roams the map in search of the valuable minicons.', 'santa.webp'),
(7, '2023-03-04 13:14:01.207', '2023-03-04 13:14:01.207', NULL, 'Sheriff', 'Sheriff is a straight-laced, no-nonsense lawman with a stern gaze and steady hand. Wearing a traditional sheriff hat and sporting a distinguished moustache, Sheriff patrols the map with a sense of purpose, seeking out the elusive treasures.', 'sheriff.webp'),
(8, '2023-03-04 13:14:01.207', '2023-03-04 13:14:01.207', NULL, 'Forager', 'Forager is resourceful and determined, with a keen eye for valuable items and a deep connection to nature. With her earthy demeanour, Forager moves through the map with purpose, searching for treasure hidden throughout the terrain.', 'forager.webp'),
(9, '2023-03-04 13:14:01.207', '2023-03-04 13:14:01.207', NULL, 'Villager', 'Villager is a brave and independent voyager who has left her home behind to seek new experiences and discover the world beyond. Determined, Villager carefully moves through the map, searching for its incredible treasures.', 'villager.webp');

--
-- Dumping data for table `constants`
--

INSERT INTO `constants` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `value`) VALUES
(1, '2023-03-04 13:19:23.777', '2023-03-04 13:19:23.777', NULL, 'minicons_in_lineup', 5),
(2, '2023-03-04 13:19:23.782', '2023-03-04 13:19:23.782', NULL, 'matchmaking_rank_range', 10),
(3, '2023-03-04 13:19:23.787', '2023-03-04 13:19:23.787', NULL, 'min_trophy_gain', 20),
(4, '2023-03-04 13:19:23.792', '2023-03-04 13:19:23.792', NULL, 'trophy_gain_range', 10),
(5, '2023-03-04 13:19:23.797', '2023-03-04 13:19:23.797', NULL, 'trophy_diff_loser', 3),
(6, '2023-03-04 13:19:23.801', '2023-03-04 13:19:23.801', NULL, 'survivor_trophy_range', 4),
(7, '2023-03-04 13:19:23.806', '2023-03-04 13:19:23.806', NULL, 'default_trophy_count', 1000),
(8, '2023-03-04 13:19:23.810', '2023-03-04 13:19:23.810', NULL, 'successive_duplicate_match_limit', 5),
(9, '2023-03-04 13:19:23.816', '2023-03-04 13:19:23.816', NULL, 'daily_attack_limit', 40),
(10, '2023-03-04 13:19:23.821', '2023-03-04 13:19:23.821', NULL, 'xp_base_count', 1000),
(11, '2023-03-04 13:19:23.826', '2023-03-04 13:19:23.826', NULL, 'xp_level_multiplier', 2),
(12, '2023-03-04 13:19:23.831', '2023-03-04 13:19:23.831', NULL, 'incr_xp_minicon', 38),
(13, '2023-03-04 13:19:23.836', '2023-03-04 13:19:23.836', NULL, 'incr_xp_user', 76),
(14, '2023-03-04 13:19:23.840', '2023-03-04 13:19:23.840', NULL, 'level_up_stat_multiplier_numerator', 110),
(15, '2023-03-04 13:19:23.840', '2023-03-04 13:19:23.840', NULL, 'level_up_stat_multiplier_denominator', 100),
(16, '2023-03-04 14:19:23.840', '2023-03-04 14:19:23.840', NULL, 'max_minicon_level', 5),
(17, '2023-03-04 15:19:23.840', '2023-03-04 15:19:23.840', NULL, 'max_unlocked_minicons', 12),
(18, '2023-03-14 21:59:05.531', '2023-03-14 21:59:05.531', NULL, 'type_multiplier_numerator', 120),
(19, '2023-03-14 21:59:05.532', '2023-03-14 21:59:05.532', NULL, 'type_multiplier_denominator', 100),
(20, '2023-03-14 21:59:05.532', '2023-03-14 21:59:05.532', NULL, 'is_arena_open', 1);

--
-- Dumping data for table `lootboxes`
--

INSERT INTO `lootboxes` (`id`, `created_at`, `updated_at`, `deleted_at`, `x`, `y`, `unlocks_id`, `region_id`) VALUES
(1, '2023-03-04 13:19:23.918', '2023-03-04 13:19:23.918', NULL, 19, 38, 14, 1),
(2, '2023-03-04 13:19:23.920', '2023-03-04 13:19:23.920', NULL, 19, 23, 11, 1),
(3, '2023-03-04 13:19:23.923', '2023-03-04 13:19:23.923', NULL, 11, 5, 17, 1),
(4, '2023-03-04 13:19:23.925', '2023-03-04 13:19:23.925', NULL, 37, 1, 7, 1),
(5, '2023-03-04 13:19:23.927', '2023-03-04 13:19:23.927', NULL, 4, 29, 8, 1),
(6, '2023-03-04 13:19:23.930', '2023-03-04 13:19:23.930', NULL, 36, 34, 1, 2),
(7, '2023-03-04 13:19:23.933', '2023-03-04 13:19:23.933', NULL, 27, 11, 6, 2),
(8, '2023-03-04 13:19:23.936', '2023-03-04 13:19:23.936', NULL, 14, 24, 2, 2),
(9, '2023-03-04 13:19:23.938', '2023-03-04 13:19:23.938', NULL, 38, 4, 18, 2),
(10, '2023-03-04 13:19:23.941', '2023-03-04 13:19:23.941', NULL, 18, 36, 5, 2),
(11, '2023-03-04 13:19:23.943', '2023-03-04 13:19:23.943', NULL, 9, 27, 3, 3),
(12, '2023-03-04 13:19:23.945', '2023-03-04 13:19:23.945', NULL, 33, 7, 13, 3),
(13, '2023-03-04 13:19:23.947', '2023-03-04 13:19:23.947', NULL, 4, 19, 19, 3),
(14, '2023-03-04 13:19:23.950', '2023-03-04 13:19:23.950', NULL, 14, 8, 15, 3),
(15, '2023-03-04 13:19:23.952', '2023-03-04 13:19:23.952', NULL, 27, 19, 9, 3),
(16, '2023-03-04 13:19:23.954', '2023-03-04 13:19:23.954', NULL, 19, 24, 4, 4),
(17, '2023-03-04 13:19:23.957', '2023-03-04 13:19:23.957', NULL, 6, 33, 10, 4),
(18, '2023-03-04 13:19:23.960', '2023-03-04 13:19:23.960', NULL, 10, 4, 20, 4),
(19, '2023-03-04 13:19:23.962', '2023-03-04 13:19:23.962', NULL, 36, 3, 12, 4),
(20, '2023-03-04 13:19:23.965', '2023-03-04 13:19:23.965', NULL, 3, 18, 16, 4);


--
-- Dumping data for table `minicons`
--

INSERT INTO `minicons` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `base_health`, `base_attack`, `description`, `image_link`, `type_id`) VALUES
(1, '2023-03-14 21:59:05.556', '2023-03-14 21:59:05.556', NULL, 'Krunk', 1040, 450, 'Krunk is a mischievous and sly goblin minicon, with a sinister appearance. It is crafty and resourceful with a talent for erupting chaos. It uses its agility to dart in and out of battle and strikes mercilessly with his sword. Krunk is infused with a mix of cunningness, and unpredictability that keeps his opponents on their toes.', 'krunk.webp', 4),
(2, '2023-03-14 21:59:05.558', '2023-03-14 21:59:05.558', NULL, 'Ember', 980, 450, 'Ember is a formidable goblin minicon that wreaks havoc wherever it goes. A being born from flame, Ember is a valorous and powerful adversary, with a host of deadly abilities at its disposal. If you find yourself facing off against Ember, be prepared for a fierce and fiery battle.', 'ember.webp', 1),
(3, '2023-03-14 21:59:05.560', '2023-03-14 21:59:05.560', NULL, 'Riptide', 1000, 450, 'Riptide is a malicious goblin minicon and is master of the water arts. It is small and blue with venomous skin. Its intense, piercing gaze can be both alluring and intimidating. Riptide is playful and curious, but also opportunistic and greedy.', 'riptide.webp', 2),
(4, '2023-03-14 21:59:05.561', '2023-03-14 21:59:05.561', NULL, 'Keranos', 1000, 468, 'Keranos is a rugged and tough goblin minicon with sharp fangs. It is fast and mercurial , delighting in causing chaos and confusion wherever it goes. It is nimble and quick, able to dodge and evade attacks with ease. Overall, Keranos is a dangerous and unpredictable foe, not to be underestimated.', 'keranos.webp', 3),
(5, '2023-03-14 21:59:05.563', '2023-03-14 21:59:05.563', NULL, 'Deimos', 1040, 450, 'Deimos is your worst nightmare. It is a menacing skeletal minicon that looms over its enemies, radiating an aura of death and decay. Its bones are ancient, and its empty eye sockets are filled with a green glow that burns with intensity. It wields a massive battle axe, and its attacks are swift and deadly. Those who cross his path are doomed to suffer a fate worse than death.', 'deimos.webp', 4),
(6, '2023-03-14 21:59:05.564', '2023-03-14 21:59:05.564', NULL, 'Adranus', 1000, 450, 'Adranus is infused with the power of fire, granting it incredible strength and resilience. It is driven by its desire to spread destruction and chaos wherever it goes, Its eyes burn with an intense heat, glowing bright red.', 'adranus.webp', 1),
(7, '2023-03-14 21:59:05.566', '2023-03-14 21:59:05.566', NULL, 'Osiris', 1000, 468, 'Osiris is a terrifying and gruesome minicon that strikes fear into the hearts of all who encounter it. Osiris posses a skeletal appearance, with bones made of ice-blue water. Its movements are swift as a stream and sharp as ice. Osiris will ensure those who encounter it don`t forget it anytime soon.', 'osiris.webp', 2),
(8, '2023-03-14 21:59:05.567', '2023-03-14 21:59:05.567', NULL, 'Thunderbolt', 980, 450, 'Thunderbolt is a fearsome and imposing creature, possessing both the power of electricity and the eerie appearance of a skeletal figure. Adorned in jagged, metallic armour, with glowing grey eyes that seem to burn with otherworldly energy, Thunderbolt is a ruthless and aggressive fighter, driven by a fierce desire to dominate and destroy its enemies.', 'thunderbolt.webp', 3),
(9, '2023-03-14 21:59:05.569', '2023-03-14 21:59:05.569', NULL, 'Daredevil', 1030, 465, 'Daredevil wields a sharp, blood-red club used to deliver deadly blows to opponents. This minicon is known for its daring moves and fearless approach to combat, making them a formidable opponent in any battle. Daredevil is fearless and always ready to take on any challenge, no matter how dangerous it may seem.', 'daredevil.webp', 4),
(10, '2023-03-14 21:59:05.571', '2023-03-14 21:59:05.571', NULL, 'Noir', 1000, 450, 'Noir is a dark armoured figure with a sword as its signature weapon. It is a master of stealth and deception and is always looking for its next score. This minicon is ruthless and calculating, with a deep sense of cynicism that comes from years of living on the fringes of society. Not much is known about Noir, and not much will ever be known about it.', 'noir.webp', 4),
(11, '2023-03-14 21:59:05.573', '2023-03-14 21:59:05.573', NULL, 'Samurai', 1035, 465, 'Samurai is a fearsome warrior, skilled in the art of combat and revered for its honour and bravery. It moves with lightning-fast speed and precision, wielding its trusty katana with deadly accuracy. It is a worthy opponent in any battle. With its deadly sword skills and strategic mind, it is sure to leave its mark on the battlefield and earn the respect of its fellow minicons.', 'samurai.webp', 4),
(12, '2023-03-14 21:59:05.575', '2023-03-14 21:59:05.575', NULL, 'Nightcrawler', 1000, 450, 'Nightcrawler is one of the most tactile and stealthy opponents on the battlefield. With its sleek, shadowy form and lightning-fast reflexes, it is known for its agility and stealth, making it a formidable opponent in battle.', 'nightcrawler.webp', 4),
(13, '2023-03-14 21:59:05.576', '2023-03-14 21:59:05.576', NULL, 'Achaeus', 1040, 450, 'Achaeus is a powerful warrior with a strong sense of honour. It is highly skilled in combat, wielding a sword with incredible proficiency. Years of training and experience have made it an intimidating opponent in combat.', 'achaeus.webp', 4),
(14, '2023-03-14 21:59:05.578', '2023-03-14 21:59:05.578', NULL, 'Laurence', 1000, 468, 'A valorant minicon renowned for its golden armour and the fiery labrys he wields in battle. In battle, Laurence is a frightening opponent. Its strength and agility make it difficult to hit, while its armour protects it from even the most powerful attacks. Its code of chivalry and loyalty are often admired.', 'laurence.webp', 1),
(15, '2023-03-14 21:59:05.580', '2023-03-14 21:59:05.580', NULL, 'Glaucus', 980, 450, 'Glaucus is a harrowing minicon that wears navy blue armour and carries a blue axe, making it instantly recognizable. Glaucus is an expert in battle, having honed its skills over many years of fighting. It is a valuable asset, and its combat prowess can often turn the tide of a battle.', 'glaucus.webp', 2),
(16, '2023-03-14 21:59:05.582', '2023-03-14 21:59:05.582', NULL, 'Taranis', 1000, 450, 'Taranis is a fearsome warrior who strikes fear into the hearts of their enemies on the battlefield. Clad in a suit of imposing black armour and wielding a massive black axe, the Taranis is one of the most skilled minicons in the game when it comes to combat. With lightning-quick reactions, Taranis can dodge incoming attacks and strike back with devastating force. This allows it to take on even the most powerful opponents and emerge victorious.', 'taranis.webp', 3),
(17, '2023-03-14 21:59:05.583', '2023-03-14 21:59:05.583', NULL, 'Sensei', 1040, 450, 'Sensei is a skilled warrior that is highly regarded as a mentor and guide. Clad in a traditional Kasa and wielding a powerful sword, Sensei is a force to be reckoned with in battle. Sensei`s ability to read its opponent`s movements and anticipate its next move is next to none.', 'sensei.webp', 4),
(18, '2023-03-14 21:59:05.585', '2023-03-14 21:59:05.585', NULL, 'Phobos', 1000, 433, 'Phobos  is a formidable minicon, equipped with a dark armour and wielding a sword that can cut through anything in its path. It is a fierce opponent that can take down enemies quickly and efficiently. Its sword strikes are fast and precise, it can dodge attacks with ease, making it difficult to hit. Additionally, its armour provides a level of protection that allows them to withstand even the most powerful attacks.', 'phobos.webp', 1),
(19, '2023-03-14 21:59:05.587', '2023-03-14 21:59:05.587', NULL, 'Njord', 1000, 433, 'A fierce viking warrior who is known for its remarkable battle skills, Njord exudes strength and intimidation. It is dressed in a brown and grey armour that protects it from any kind of attack in battle. It uses its strength and skill to outmanoeuvre its opponents and strike them down with devastating blows.', 'njord.webp', 2),
(20, '2023-03-14 21:59:05.588', '2023-03-14 21:59:05.588', NULL, 'Shockwave', 1000, 433, 'Shockwave is known for its impressive battle skills and intimidating appearance. This minicon wears sleek black armour that gleams in the light, giving it an air of mystery and danger. Equipped with a massive axe, this minicon is not to be messed with in combat. Its axe is well-balanced and expertly crafted, allowing it to deliver swift and powerful strikes that can take down even the most resilient opponents.', 'shockwave.webp', 3);

--
-- Dumping data for table `minicon_types`
--

INSERT INTO `minicon_types` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `description`) VALUES
(1, '2023-03-04 13:19:23.757', '2023-03-04 13:19:23.757', NULL, 'FIRE', 'Fire is a dangerous element. It can burn you, and it can burn your house down. It can also be used to cook food, and to keep you warm.'),
(2, '2023-03-04 13:19:23.762', '2023-03-04 13:19:23.762', NULL, 'WATER', 'Water is a very useful element. It can be used to drink, to wash, and to cook. It can also be used to put out fires.'),
(3, '2023-03-04 13:19:23.766', '2023-03-04 13:19:23.766', NULL, 'THUNDER', 'Thunder is as dangerous as it sounds but not just dangerous it is fast in attacking opponents'),
(4, '2023-03-04 13:19:23.771', '2023-03-04 13:19:23.771', NULL, 'NORMAL', 'Normal might seem basic, but it has no intrinsic weakness, unlike other types. Who`s basic now?');

--
-- Dumping data for table `perks`
--

INSERT INTO `perks` (`id`, `created_at`, `updated_at`, `deleted_at`, `perk_trigger`, `perk_name`, `effect`, `description`, `base_value`, `minicon_id`) VALUES
(1, '2023-03-14 21:59:05.591', '2023-03-14 21:59:05.591', NULL, 'HURT', 'Enraged', 'ATKBUFF', 'When Krunk is hurt, it boosts the attack of itself, the minicon behind it and the minicon in front of it, in its own team', 12, 1),
(2, '2023-03-14 21:59:05.592', '2023-03-14 21:59:05.592', NULL, 'HURT', 'Enraged', 'ATKBUFF', 'When Ember is hurt, it boosts the attack of itself, the minicon behind it and the minicon in front of it, in its own team', 7, 2),
(3, '2023-03-14 21:59:05.594', '2023-03-14 21:59:05.594', NULL, 'START', 'Trataka', 'HEAL', 'At the start of each round, Ember heals the minicon behind it and the minicon in front of it in its own team', 50, 2),
(4, '2023-03-14 21:59:05.595', '2023-03-14 21:59:05.595', NULL, 'HURT', 'Enraged', 'ATKBUFF', 'When Riptide is hurt, it boosts the attack of itself, the minicon behind it and the minicon in front of it in its own team', 9, 3),
(5, '2023-03-14 21:59:05.597', '2023-03-14 21:59:05.597', NULL, 'DEAD', 'Bubble Blast ', 'DMG', 'When Riptide dies, it damages the first minicon in the opposing team', 63, 3),
(6, '2023-03-14 21:59:05.598', '2023-03-14 21:59:05.598', NULL, 'HURT', 'Enraged', 'ATKBUFF', 'When Keranos is hurt, it boosts the attack of itself and the minicon in front of it in its own team', 7, 4),
(7, '2023-03-14 21:59:05.600', '2023-03-14 21:59:05.600', NULL, 'DEAD', 'Drain', 'ATKDEBUFF', 'When Keranos dies, it drains the attack of first, second and third minicons in the opposing team', 13, 4),
(8, '2023-03-14 21:59:05.602', '2023-03-14 21:59:05.602', NULL, 'HURT', 'Hex', 'ATKDEBUFF', 'When Deimos is hurt, it drains the attack of the second and third minicons in the opposing team', 11, 5),
(9, '2023-03-14 21:59:05.604', '2023-03-14 21:59:05.604', NULL, 'HURT', 'Hex', 'ATKDEBUFF', 'When Adranus is hurt, it drains the attack of the second and third minicons in the opposing team', 8, 6),
(10, '2023-03-14 21:59:05.606', '2023-03-14 21:59:05.606', NULL, 'DEAD', 'War Cry', 'ATKBUFF', 'When Adranus dies, it boosts the attack of minicon that is two places behind it, and the one behind that in its own team', 15, 6),
(11, '2023-03-14 21:59:05.608', '2023-03-14 21:59:05.608', NULL, 'HURT', 'Hex', 'ATKDEBUFF', 'When Orisis is hurt, it drains the attack of the third minicon in the opposing team', 9, 7),
(12, '2023-03-14 21:59:05.610', '2023-03-14 21:59:05.610', NULL, 'DEAD', 'Drain', 'ATKDEBUFF', 'When Osiris dies, it drains the attack of the first, second and third minicons in the opposing team', 13, 7),
(13, '2023-03-14 21:59:05.611', '2023-03-14 21:59:05.611', NULL, 'HURT', 'Hex', 'ATKDEBUFF', 'When Thunderbolt gets hurt, it drains the attack of the second and third minicons in the opposing team', 7, 8),
(14, '2023-03-14 21:59:05.612', '2023-03-14 21:59:05.612', NULL, 'START', 'Brewing Storm', 'ATKBUFF', 'At the start of the round, Thunderbolt boosts the attack of minicon three places in front of it, in its own team', 6, 8),
(15, '2023-03-14 21:59:05.614', '2023-03-14 21:59:05.614', NULL, 'DEAD', 'Team Player', 'ATKBUFF', 'When Daredevil dies, it boosts the attack of minicon two places behind it and the one after that in its own team', 10, 9),
(16, '2023-03-14 21:59:05.616', '2023-03-14 21:59:05.616', NULL, 'START', 'Sharp Shooter', 'DMG', 'At the start of the round, Daredevil damages the second and third minicons of the opposing team', 25, 9),
(17, '2023-03-14 21:59:05.617', '2023-03-14 21:59:05.617', NULL, 'START', 'Sharp Shooter', 'DMG', 'At the start of the round, Noir damages the second and third minicons of the opposing team', 25, 10),
(18, '2023-03-14 21:59:05.619', '2023-03-14 21:59:05.619', NULL, 'DEAD', 'Transitive Regen', 'HEAL', 'When Noir dies, it heals the three minicons right behind it in its own team', 75, 10),
(19, '2023-03-14 21:59:05.620', '2023-03-14 21:59:05.620', NULL, 'START', 'Bushido', 'ATKDEBUFF', 'At the start of the round, Samurai drains the attack of the third, fourth and fifth minicons of the opposing team', 7, 11),
(20, '2023-03-14 21:59:05.623', '2023-03-14 21:59:05.623', NULL, 'KILL', 'Zantetsuken Kaeshi', 'HEAL', 'When Samurai gets a kill, it heals itself', 55, 11),
(21, '2023-03-14 21:59:05.625', '2023-03-14 21:59:05.625', NULL, 'START', 'Speed strike', 'DMG', 'At the start of the round, Nightcrawler damages the second, third, fourth and fifth minicons of the opposing team', 15, 12),
(22, '2023-03-14 21:59:05.627', '2023-03-14 21:59:05.627', NULL, 'KILL', 'Clean Finish', 'ATKBUFF', 'When Nightcrawler gets a kill, it boosts the attack of the minicon two places behind it in its own team', 20, 12),
(23, '2023-03-14 21:59:05.628', '2023-03-14 21:59:05.628', NULL, 'KILL', 'Glory', 'HEAL', 'When Achaeus gets a kill, it heals itself', 90, 13),
(24, '2023-03-14 21:59:05.630', '2023-03-14 21:59:05.630', NULL, 'KILL', 'Glory', 'HEAL', 'When Laurence gets a kill, it heals itself', 45, 14),
(25, '2023-03-14 21:59:05.633', '2023-03-14 21:59:05.633', NULL, 'HURT', 'Heat Stroke', 'DMG', 'When Laurence gets hurt, it damages the second minicon of the opposing team', 43, 14),
(26, '2023-03-14 21:59:05.635', '2023-03-14 21:59:05.635', NULL, 'KILL', 'Glory', 'HEAL', 'When Glaucus gets a kill, it heals itself', 55, 15),
(27, '2023-03-14 21:59:05.637', '2023-03-14 21:59:05.637', NULL, 'DEAD', 'Tidal Fury', 'ATKBUFF', 'When Glaucus dies, it boosts the attack of the minicon two places behind it and the two minicons after that in its own team', 15, 15),
(28, '2023-03-14 21:59:05.639', '2023-03-14 21:59:05.639', NULL, 'KILL', 'Glory', 'HEAL', 'When Taranis gets a kill, it heals itself and the minicon behind it in its own team', 45, 16),
(29, '2023-03-14 21:59:05.640', '2023-03-14 21:59:05.640', NULL, 'START', 'StunOver', 'DMG', 'At the start of each round, Taranis damages the second and third minicons in the opposing team', 25, 16),
(30, '2023-03-14 21:59:05.642', '2023-03-14 21:59:05.642', NULL, 'DEAD', 'Poison Strike', 'DMG', 'When Sensei dies, it damages the first minicon in the opposing team', 160, 17),
(31, '2023-03-14 21:59:05.644', '2023-03-14 21:59:05.644', NULL, 'DEAD', 'Poison Strike', 'DMG', 'When Phobos dies, it damages the first and second minicon in the opposing team', 68, 18),
(32, '2023-03-14 21:59:05.646', '2023-03-14 21:59:05.646', NULL, 'KILL', 'Flame Rush', 'ATKBUFF', 'When Phobos gets a kill, it boosts its own attack', 20, 18),
(33, '2023-03-14 21:59:05.647', '2023-03-14 21:59:05.647', NULL, 'DEAD', 'Poison Strike', 'DMG', 'When Njord dies, it damages the first minicon in the opposing team', 98, 19),
(34, '2023-03-14 21:59:05.648', '2023-03-14 21:59:05.648', NULL, 'HURT', 'Calm Water', 'HEAL', 'When Njord is hurt, it to heals itself, the minicon behind it and the minicon infront of it in its own team', 35, 19),
(35, '2023-03-14 21:59:05.650', '2023-03-14 21:59:05.650', NULL, 'DEAD', 'Poison Strike', 'DMG', 'When Shockwave dies, it damages the first minicon in the opposing team', 85, 20),
(36, '2023-03-14 21:59:05.652', '2023-03-14 21:59:05.652', NULL, 'START', 'thunger surge', 'DMG', 'At the start of the round, Shockwave damages the second and third minicons in the opposing team', 35, 20);

--
-- Dumping data for table `regions`
--

INSERT INTO `regions` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `description`) VALUES
(1, '2023-03-04 13:14:01.208', '2023-03-04 13:14:01.208', NULL, 'CENTRAL', 'Citadel is a peaceful and serene area that connects to three other regions: Pyralia, Atlantia, and Valhalla. It is home to several mystic gardens and a beautiful lake that adds to the peaceful ambiance of the region. The central focus of the map is an ancient minicon statue located in the centre of the region. The statue represents the original minicon. Players can explore this region seeking out lootboxes hidden throughout the region.'),
(2, '2023-03-04 13:14:01.210', '2023-03-04 13:14:01.210', NULL, 'FIRE', 'Pyralia is a dangerous and intense area that is characterised by its harsh terrain and scorching temperatures. It is characterised by two cemeteries. The cemeteries add an air of mystery and intrigue to the region, drawing players deeper into the fire region`s danger. Pyralia is also home to a few lava lakes that bubble and seethe with molten lava, making the terrain difficult to traverse.'),
(3, '2023-03-04 13:14:01.212', '2023-03-04 13:14:01.212', NULL, 'WATER', 'Atlantia is a refreshing area that is characterised by its tranquil waters and lush greenery. It is predominantly covered by water, with small grassy areas scattered throughout. Atlantia is filled with small boats and lily pads that adds to its beauty. Atlantia is a calming and relaxing area that provides a stark contrast to the intensity of Pyralia.'),
(4, '2023-03-04 13:14:01.214', '2023-03-04 13:14:01.214', NULL, 'THUNDER', 'Valhalla is a stunning area that is filled to the brim with lush grasslands and intricate stone pathways. It is imbued with a powerful energy that crackles in the air and electrifies everything around it. Valhalla is a thrilling and exciting area that provides a stark contrast to the calming tranquillity of the Atlantia. It is a place where players can test their skills and explore the wonders of nature while facing the dangers of the elements.');

--
-- Dumping data for table `targets`
--

INSERT INTO `targets` (`id`, `created_at`, `updated_at`, `deleted_at`, `target_val`, `perk_id`) VALUES
(1, '2023-03-14 21:59:05.654', '2023-03-14 21:59:05.654', NULL, -1, 1),
(2, '2023-03-14 21:59:05.655', '2023-03-14 21:59:05.655', NULL, 0, 1),
(3, '2023-03-14 21:59:05.657', '2023-03-14 21:59:05.657', NULL, 1, 1),
(4, '2023-03-14 21:59:05.658', '2023-03-14 21:59:05.658', NULL, -1, 2),
(5, '2023-03-14 21:59:05.660', '2023-03-14 21:59:05.660', NULL, 0, 2),
(6, '2023-03-14 21:59:05.661', '2023-03-14 21:59:05.661', NULL, 1, 2),
(7, '2023-03-14 21:59:05.663', '2023-03-14 21:59:05.663', NULL, -1, 3),
(8, '2023-03-14 21:59:05.664', '2023-03-14 21:59:05.664', NULL, 1, 3),
(9, '2023-03-14 21:59:05.666', '2023-03-14 21:59:05.666', NULL, -1, 4),
(10, '2023-03-14 21:59:05.668', '2023-03-14 21:59:05.668', NULL, 0, 4),
(11, '2023-03-14 21:59:05.669', '2023-03-14 21:59:05.669', NULL, 1, 4),
(12, '2023-03-14 21:59:05.671', '2023-03-14 21:59:05.671', NULL, 0, 5),
(13, '2023-03-14 21:59:05.673', '2023-03-14 21:59:05.673', NULL, -1, 6),
(14, '2023-03-14 21:59:05.675', '2023-03-14 21:59:05.675', NULL, 0, 6),
(15, '2023-03-14 21:59:05.676', '2023-03-14 21:59:05.676', NULL, 0, 7),
(16, '2023-03-14 21:59:05.678', '2023-03-14 21:59:05.678', NULL, 1, 7),
(17, '2023-03-14 21:59:05.679', '2023-03-14 21:59:05.679', NULL, 2, 7),
(18, '2023-03-14 21:59:05.681', '2023-03-14 21:59:05.681', NULL, 1, 8),
(19, '2023-03-14 21:59:05.683', '2023-03-14 21:59:05.683', NULL, 2, 8),
(20, '2023-03-14 21:59:05.684', '2023-03-14 21:59:05.684', NULL, 1, 9),
(21, '2023-03-14 21:59:05.685', '2023-03-14 21:59:05.685', NULL, 2, 9),
(22, '2023-03-14 21:59:05.687', '2023-03-14 21:59:05.687', NULL, 2, 10),
(23, '2023-03-14 21:59:05.688', '2023-03-14 21:59:05.688', NULL, 3, 10),
(24, '2023-03-14 21:59:05.690', '2023-03-14 21:59:05.690', NULL, 2, 11),
(25, '2023-03-14 21:59:05.691', '2023-03-14 21:59:05.691', NULL, 0, 12),
(26, '2023-03-14 21:59:05.693', '2023-03-14 21:59:05.693', NULL, 1, 12),
(27, '2023-03-14 21:59:05.695', '2023-03-14 21:59:05.695', NULL, 2, 12),
(28, '2023-03-14 21:59:05.697', '2023-03-14 21:59:05.697', NULL, 1, 13),
(29, '2023-03-14 21:59:05.698', '2023-03-14 21:59:05.698', NULL, 2, 13),
(30, '2023-03-14 21:59:05.700', '2023-03-14 21:59:05.700', NULL, -3, 14),
(31, '2023-03-14 21:59:05.702', '2023-03-14 21:59:05.702', NULL, 2, 15),
(32, '2023-03-14 21:59:05.704', '2023-03-14 21:59:05.704', NULL, 3, 15),
(33, '2023-03-14 21:59:05.705', '2023-03-14 21:59:05.705', NULL, 1, 16),
(34, '2023-03-14 21:59:05.707', '2023-03-14 21:59:05.707', NULL, 2, 16),
(35, '2023-03-14 21:59:05.708', '2023-03-14 21:59:05.708', NULL, 1, 17),
(36, '2023-03-14 21:59:05.710', '2023-03-14 21:59:05.710', NULL, 2, 17),
(37, '2023-03-14 21:59:05.711', '2023-03-14 21:59:05.711', NULL, 1, 18),
(38, '2023-03-14 21:59:05.713', '2023-03-14 21:59:05.713', NULL, 2, 18),
(39, '2023-03-14 21:59:05.714', '2023-03-14 21:59:05.714', NULL, 3, 18),
(40, '2023-03-14 21:59:05.716', '2023-03-14 21:59:05.716', NULL, 2, 19),
(41, '2023-03-14 21:59:05.717', '2023-03-14 21:59:05.717', NULL, 3, 19),
(42, '2023-03-14 21:59:05.720', '2023-03-14 21:59:05.720', NULL, 4, 19),
(43, '2023-03-14 21:59:05.722', '2023-03-14 21:59:05.722', NULL, 0, 20),
(44, '2023-03-14 21:59:05.724', '2023-03-14 21:59:05.724', NULL, 1, 21),
(45, '2023-03-14 21:59:05.727', '2023-03-14 21:59:05.727', NULL, 2, 21),
(46, '2023-03-14 21:59:05.729', '2023-03-14 21:59:05.729', NULL, 3, 21),
(47, '2023-03-14 21:59:05.732', '2023-03-14 21:59:05.732', NULL, 4, 21),
(48, '2023-03-14 21:59:05.734', '2023-03-14 21:59:05.734', NULL, 2, 22),
(49, '2023-03-14 21:59:05.736', '2023-03-14 21:59:05.736', NULL, 0, 23),
(50, '2023-03-14 21:59:05.740', '2023-03-14 21:59:05.740', NULL, 0, 24),
(51, '2023-03-14 21:59:05.742', '2023-03-14 21:59:05.742', NULL, 1, 25),
(52, '2023-03-14 21:59:05.745', '2023-03-14 21:59:05.745', NULL, 0, 26),
(53, '2023-03-14 21:59:05.747', '2023-03-14 21:59:05.747', NULL, 2, 27),
(54, '2023-03-14 21:59:05.748', '2023-03-14 21:59:05.748', NULL, 3, 27),
(55, '2023-03-14 21:59:05.750', '2023-03-14 21:59:05.750', NULL, 4, 27),
(56, '2023-03-14 21:59:05.753', '2023-03-14 21:59:05.753', NULL, 0, 28),
(57, '2023-03-14 21:59:05.756', '2023-03-14 21:59:05.756', NULL, 1, 28),
(58, '2023-03-14 21:59:05.758', '2023-03-14 21:59:05.758', NULL, 1, 29),
(59, '2023-03-14 21:59:05.761', '2023-03-14 21:59:05.761', NULL, 2, 29),
(60, '2023-03-14 21:59:05.763', '2023-03-14 21:59:05.763', NULL, 0, 30),
(61, '2023-03-14 21:59:05.764', '2023-03-14 21:59:05.764', NULL, 0, 31),
(62, '2023-03-14 21:59:05.766', '2023-03-14 21:59:05.766', NULL, 1, 31),
(63, '2023-03-14 21:59:05.768', '2023-03-14 21:59:05.768', NULL, 0, 32),
(64, '2023-03-14 21:59:05.769', '2023-03-14 21:59:05.769', NULL, 0, 33),
(65, '2023-03-14 21:59:05.771', '2023-03-14 21:59:05.771', NULL, -1, 34),
(66, '2023-03-14 21:59:05.772', '2023-03-14 21:59:05.772', NULL, 0, 34),
(67, '2023-03-14 21:59:05.773', '2023-03-14 21:59:05.773', NULL, 1, 34),
(68, '2023-03-14 21:59:05.775', '2023-03-14 21:59:05.775', NULL, 0, 35),
(69, '2023-03-14 21:59:05.776', '2023-03-14 21:59:05.776', NULL, 1, 36),
(70, '2023-03-14 21:59:05.778', '2023-03-14 21:59:05.778', NULL, 2, 36);
SET FOREIGN_KEY_CHECKS=1;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
