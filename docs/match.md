# Match
A match is started when user makes a GET-request to `/api/match/start`.

## [Workflow](../server/controller/match/start_match.go)
1. If there is any error encountered anywhere along the pipeline, a default error ("Error in finding a match currently! Try again later!") is returned to the user as the entire pipeline is estimated to happen across <1 second. 

1. Ensure the user (attacker) has not exhausted his attacks. Find an opponent (defender) of similar skill level using the [Matchmaking Algorithm](helper_algos.md#matchmaking-algorithm).

1. Obtain the lineups of both attacker and defender and [simulate]() the battle.

1. Create an entry in `matchmaing_details` that outlines the initial state of the battle.

1. Create an entry in `simulation_details` that contains all details required for frontend to render the batle simulation.

1. Call apropriate [functions](../docs/helper_algos.md) to update XPs and calculate trophy difference.

1. Create an entry in `battle_results` that contains details like battle winner, trophy difference etc.
