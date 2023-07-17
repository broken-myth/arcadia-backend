# Helper Algorithms

## [Matchmaking Algorithm](../server/helper/match/matchmaker.go)

1. Get the userIds of users of ranks within +-10 of the attacker. 

1. Select a random user that has not fought a battle against the attacker in the attacker's last `successive_duplicate_match_limit(5)` matches.

## [Trophy Algorithm](../server/helper/general/trophies.go)

Assume values of constants to be:
 * minicons_in_lineup= 5   (number of minicons battling per side)
 * matchmaking_rank_range= 20   (range of ranks to search in, i.e, 20 = upto +10 and -10 ranks)
 * min_trophy_gain= 20   (winners gain at least 20 trophies)
 * trophy_gain_range= 10   (max - min possible trophies won)
 * trophy_diff_loser= 5   (if winner gets x trophies, loser loses x-5 trophies)
 * survivor_trophy_range= 4  (the weightage range given to secondary factor - number of surviving minicons)

These constants are fetched first

**Primary factor:** Difference in ranks. Depending on the ratio of difference of ranks to matchmaking_rank_range, trophies are adjusted. 
Assume attacker wins and has rank 100:
 * If defender was of rank <93 - winnerTrophies = 1
 * If defender was of rank 93-96 - winnerTrophies = 3
 * If defender was of rank 97-103 - winnerTrophies = 5
 * If defender was of rank 104-107 - winnerTrophies = 7
 * If defender was of rank >108 - winnerTrophies = 9

**Secondary Factor**: proportional to number of surviving minicons (survivors/total * survivor_trophy_range). So:
* If attacker survivors = 1, winnerTrophies+= 0
* If attacker survivors = 2, winnerTrophies+= 1
* If attacker survivors = 3, winnerTrophies+= 2
* If attacker survivors = 4, winnerTrophies+= 3
* If attacker survivors = 5, winnerTrophies+= 4

Attacker gains `winnerTrophies` + `min_trophy_gain` trophies. <br>
Defender would lose `winnerTrophies` + `min_trophy_gain - trophy_diff_loser` trophies.

So technically, sum total of all trophies in game = 1000\*total_users + 5\*(number of battles that were not a draw)

## [XP Algorithm](../server/helper/general/xp_levels.go)
Assume values of constants to be:
 * incr_xp_minicon = 100 (fixed amount of xp gained by minicon for playing a battle)
 * incr_xp_user = 200 (fixed amount of xp gained by user for playing a battle, apart from minicon xp)
 * xp_level_multiplier = 2 (to go to the next level, you need 2x the xp of the previous level)
 * xp_base_count = 1000 (base xp required to reach level 2)
 * level_up_stat_multiplier = 1.2 (multiplier to increase stats of minicon when it levels up. stored as 120)

These constants are fetched first.

When UpdateXpUserLineup is called, the following happens:
1. For each minicon in the lineup, the xp is incremented by `incr_xp_minicon` and the level is updated if required.

1. Levels work like:
    * 0-1k XP = level 1
    * 1k-2k = level 2
    * 2k-4k = level 3
    * 4k-8k = level 4 and so on

1. The user's xp is incremented by `incr_xp_user` + (total amount of XP his minicons gained in previous step)


## [Leaderboard Algorithm](../server/helper/general/leaderboard.go)

1. Redis is used to store the leaderboard using a sorted set to maximise efficiency. The value is the userId (value to identify entries) and the score (value to sort based on) is the number of trophies the user has.

1. The trophies of attacker and defender are updated every time a battle is completed. The leaderboard is sorted by the number of trophies a user has. 
