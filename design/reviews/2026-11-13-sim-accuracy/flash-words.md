# Sim accuracy 2026-11-13: the priests' tank Flash Heal in words (wow-player)

Since ebb2fa545 the Holy and Discipline tank Flash Heal fires under 20% always, or under 35% while mana is on pace (an `or` in the APL). The panel line still read "on the tank below 35%", and the card-text generator (`web/src/lib/bis/apl-condition.ts`) gave up on the `or`, so a collapsed row kept only its note. The generator now reads an `or` of the target's own health thresholds, each branch with its own mana clause, and sets the next target apart with a semicolon.

## Lines reviewed

1. Addon panel and live site card (curated `panel_when`, Holy step 4, Discipline step 9):
   "on the tank below 20%, or below 35% while mana holds"
2. Generated card words, Holy (party step not paced):
   "Cast on the tank below 20%, or below 35% while your mana is at least 90% of the share of the fight left; then a party member below 30%."
3. Generated card words, Discipline (party step paced, its own group):
   "Cast on the tank below 20%, or below 35% while your mana is at least 90% of the share of the fight left; or on a party member below 30%, when your mana is at least 90% of the share of the fight left."

## Verdict

- Line 1: **SHIP.** 52 characters, the neighbours' "while mana holds" voice; the mana clause reads as gating only the 35% branch, and the 20% branch is plainly unconditional.
- Line 2: **SHIP** (after one NO SHIP: the first draft joined the party step with ", then", which could read as sitting inside the mana clause; the semicolon ends the tank step).
- Line 3: **SHIP.** Optional nit, not taken: "; then" would match neighbouring rows better than "; or", but "; or" is the generator's existing join between qualifier groups on every card.
