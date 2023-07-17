# Database Seeding

Populates the database with data from the seeds directory.

### Generation:
##### Any of the Below:
* USE PMA (export->custom. Only data not structure. Select `Disable foreign key checks`)
* create_mysql_dump.sh

###  Prod vs Dev:

* Prod Seeds = characters, regions, minicon_types, constants, minicons, perks, targets, lootboxes
* Dev Seeds = **Prod Seeds** + users, user_registrations, owned_minicons, owned_perks, lineups, matchmaking_details, battle _results, admins

##### Validate Foriegn Key Integrity of seeds in this order: 
1. characters
1. regions
1. minicon_types
1. constants
1. minicons
1. perks
1. targets
1. lootboxes
1. user_registrations
1. users
1. admins
1. owned_minicons
1. owned_perks
1. lineups
1. matchmaking_details
1. battle_results
