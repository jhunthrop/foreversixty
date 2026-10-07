# Inferior trainer ranks (2026-10-07)

Decision: a rank whose effect is weaker than the rank before it in the client is never the highest learned rank. `data/curated/inferior-ranks.json` names them, the `spellranks` emitter writes `inferior: true`, and `sim/internal/spellranks` and the ladder skip them. A player keeps the stronger rank, and so does the sim.

## The trigger

With Ahn'Qiraj ranks off, a level-60 hunter's top trainer Aspect of the Hawk is rank 6 (14322, level 58). The client gives it 55 ranged attack power against rank 5's 90 (14321, level 48); the live Forever tooltip says 55 as well. The ladder rewrite and the engine's own registration both took the highest rank, so the hunter cast the weaker aspect from level 58. Both now keep rank 5.

## Sweep

Every trainer spell with ranked rows, all nine classes (`data/builds/1.60.1.70009/trainables` for the rank list, `spellconst` for the client's effect amounts). A rank is flagged when the amount of its predecessor's first non-zero effect is larger in absolute value, compared inside one id family (classic ids, and Forever copies above 100000, each against their own predecessors). A second pass compared every effect index. The CSV sources named in the brief are not in this checkout's `raw/`; `spellconst` is the same client data, already normalised.

Inferior, added to the list:

| Class | Spell | Rank, id | Client amount | Predecessor |
|---|---|---|---|---|
| Hunter | Aspect of the Hawk | 6, 14322 | 55 ranged attack power | rank 5 (14321): 90 |
| Hunter | Trueshot Aura | 5, 20906 | 50 ranged attack power | rank 4 (20905): 75 |
| Hunter | Hunter's Mark | 4, 14325 | 71 ranged attack power taken | rank 3 (14324): 98 |

Hunter's Mark rank 4 also has a Forever copy, 1213268, at 110; the engine already keeps that one (`HunterSMarkSpellId`), so a rank-4 tier holds one inferior id and one sound id, and the rewrite resolves to the sound one.

Flagged by the sweep and not added:

- Priest Power Word: Fortitude 23948 (54 against rank 5's 56): a scripted copy at rank 6; the real rank 6 (10938) is 70.
- Priest Holy Fire rank 6: direct damage rises 106 to 139 while the per-tick amount falls 13 to 10 (five ticks); the rank is stronger in total (189 against 171), so it stays.
- Druid Prowl, Rogue Stealth: the movement slow gets smaller (-40 to -30), an improvement.
- Mage Polymorph effect 2, Warrior Improved Pummel: a dummy effect and a talent, not trainer damage or healing.
- Druid Rejuvenation, Mage Fire Ward and Frost Ward, Shaman Earth Shock: appear weaker only when a classic id is compared with a Forever copy of the preceding rank. Inside each family the ranks rise.

No healing or damage rank other than the three above is weaker than its predecessor.

Finding outside this lane: the engine's raid buff `core.TrueshotAura` (`sim/core/buffs.go`) is hard-coded to 100 melee and ranged attack power under spell 20906, which the client gives 50. Not changed here.

## Fork

`sim/hunter/aspects.go` registered only the highest learned rank, so rank 5 was never castable at level 58 to 60. `getMaxHawkRank` now picks the learned rank with the most attack power (rank 5 while the book is off, rank 7 with `IncludeAQ`). Tests: level 48/58/60 register 14321 at 90 mana, 14322 is never registered. `ui/hunter/apls/p1.apl.json` and the synced `forever_*` rotations name 14321.

## Level-60 movement (relative, ladder DPS)

| Spec | Change |
|---|---|
| hunter-marksmanship | +1.4% |
| hunter-beast-mastery | +0.6% |

Other specs and levels below 60 are unchanged.
