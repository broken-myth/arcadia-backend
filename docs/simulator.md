# Simulator
Simulator is started when the [start match](../server/controller/match/start_match.go) function invokes the Init function
## [Workflow](../simulator/simulator.go)

The Simulator runs in four stages, as described below.

### Preload

1. Fetches the complete lineup data from the database for each Player.

2. Processes the data from the database into objects the game can work upon in the update function.

### Update

1. The Update function executes a round and updates the lineup of the minicon at the end of a round

2. The log of each round is appended at the end of each iteration

3. The game ends when either the Attacker or Defender has no minicons alive, or the number of rounds exceeds 100.

4. If the game ends due to the number of rounds exceeding 100, the team with the most minicons alive wins.

### End 

1. The end function updates the DB with the logs

After the successful execution of the above functions, the Simulator returns the number of survivors.

## [Minicons](../simulator/functions/minicons.go)

Players' lineups comprise minicons that attack and cast perks that update the stats of the player's or opponent's lineup.

Minicons comprise several stats which alter through the course of the game.

| Attribute           |  Description    |       Type       |
|---------------------|-----------------|------------------|
| **Health**      |     The current health of the minicon     |       uint      |
| **Attack**         |     The current attack of the minicon       |       string      |
| **Type**          |    The type of the minicon   |       string      |
| **Dead**  |      Whether a minicon is alive or dead      |       []int      |
| **Base Health**      |      The Health at the start of the Game     |       uint      |
| **Perks**      |      Minicon Perks     |       uint      |

### Types

The game comprises the following types that have their strengths and weaknesses.

| _Types_     | **Normal** | **Fire** | **Water** | **Thunder** |
|-------------|------------|----------|-----------|-------------|
| **Normal**  | Equal      | Equal    | Equal     | Equal       |
| **Fire**    | Equal      | Equal    | Weak      | Strong      |
| **Water**   | Equal      | Strong   | Equal     | Weak        |
| **Thunder** | Equal      | Weak     | Strong    | Equal       |

The Nature of each type results in a type multiplier:
* Equal = 1
* Strong = 1/Weak 
  
The Above Type multipliers are reflected only in the attack, not the perks.

### [Perks](../simulator/functions/perks.go)

Each Minicon comprises 1-2 unique perks, which get triggered under certain conditions and affect the minicon stats.

The Stats involved in perks are

| Attribute           |  Description    |       Type       |
|---------------------|-----------------|------------------|
| **Perk Value**      |      The Value of the perk stat     |       uint      |
| **Trigger**         |       The Stage/Situation under which Perk Gets triggered     |       string      |
| **Effect**          |     What the Perk Does   |       string      |
| **Target Indices**  |      The Minicon's affected      |       []int      |




Perks are classified under two types based on effect.

1. Self Perks: Heal, Attack Buff -> Affects Player's lineup
2. Non-Self Perks: Damage, Attack Debuff -> Affects Opponents Lineup 


Perks get triggered under certain triggers.

**Start Perks** : These perks get triggered at the Start stage of the round.

**Hurt Perks** : These Perks get triggered at the Hurt stage of the round.

**Kill Perks** : These Perks get triggered when a minicon gets a kill.

**Dead Perks** : These Perks get triggered when a minicon dies.

Target Indices:

1. For Self perk, the target indices are relative w.r.t to the index of minicon casting the Perk
2. For Non Self perk, the target indices are relative w.r.t to the index of the leading opponent perk
3. Indices refer to the lineup's next or previous alive minicon. 

   For Self Perks:
   * If the target index of a self perk being cast is [0], it heals the minicon casting the Perk
   * If the target index of a self perk being cast is [1,2], it heals the first and second alive minicon behind the minicon casting the Perk
   * If the target index of a self perk being cast is [-2], it heals the second alive minicon ahead of the minicon casting the Perk
   <br> </br> 

   For Non-Self Perks:
   * If the target index of a non-self perk being cast is [0], it damages the leading opponent minicon
   * If the target index of a non-self perk being cast is [1,2], it damages the first and second alive minicon behind the leading opponent minicon
   <br> </br> 

4. If the target indices are out of bounds, the Perk is not cast

Notes:
 * When a minicon gets hurt and dies later, only the Kill perk and not the hurt Perk gets triggered
 * Heal can only heal upto the max health of the minicon and not beyond
 * Attack Debuff reduces attack stat  to a minimum of 1

## Stashing of Attacks and Perks
    
1. The Game is designed in such a way that simultaneous updation of Minicons takes place when multiple minicons are in a situation to attack each other
2. This is implemented via stashing the minicon attacks and perks when multiple minicons are to apply their perks
3. This ensures that a minicon can successfully cast its perk/ attack and not die before it can apply it 

## [Round](../simulator/functions/stages.go)

Each game comprises several rounds that update and reflect the changes in each minicon

Each round comprises three stages determining when a minicon lineup can cast a perk and attack other minicons.

### Start Perk Stage
   - In this stage, the start perks of the minicons are cast pairwise between the attacker and defender lineups
   - If a minicon in the pair has a start perk, it is triggered
   - If both the minicons do not have a pair, the turn is skipped
   - If both the minicons have a start perk, the priority is flipped, and both the start perks get triggered
   -When a minicons dies, the dead and kill perks get triggered for the minicon that died and the minicon that got the kill, respectively (if the Perk exists), and this process repeats
   In Essence, it creates a chain of dead and kill Perks, which repeats whenever a minicon dies, and minicon gets a kill.
   - In case a minicon does not have a pair, the minicon's Perk is triggered, but there is no flipping of priority
   

### Attack Stage
   - During this stage of the round, both the leading minicon clash against each other and inflict damage 
 - When a minicons dies, the dead and kill perks get triggered for the minicon that died and the minicon that got the kill, respectively (if the Perk exists), and this process repeats

### Hurt Stage
   - During this stage, the minicons which got hurt and are alive cast dead perks (if present)
 -When a minicons dies, the dead and kill perks get triggered for the minicon that died and the minicon that got the kill, respectively (if the Perk exists), and this process repeats. The Hurt Perk for this minicon, if not cast yet, will not get cast since the minicon has died and they cannot cast its hurt Perk


## Note 

- The order of hurt perks is in the order the minicons got hurt
- The order of minicons that get their dead kill perks triggered is first, dead minicons cast their perks, then the minicons who got the kill
- The order in which minicon goes first in case there is a choice between attacker and defender is determined by the priority

## Priority
   1. To invoke randomness and prevent biases towards attacker and defender, priority is included in the game, which decides which Player goes first in each Stage
   2. This is purely to determine the order in which hurt minicons and dead minicons trigger their stats
   3. Each time priority is used in a function, the priority gets flipped
   4. The initial priority value is towards the Player with lower trophies