# Sim-accuracy round 20: integration (2026-11-05)

Round 20 started from the landed round 19 (site 04d7da663, fork da54b7b06) on client 1.60.1.70755. One lane per class plus platform. Midway, main activated client 1.60.1.70790 (main 6beb919c2, engine pin abb4e6b8e, guides rounds 17 and 18); int/r20 took both before the pin. The lane reports sit beside this file.

**Not a zero round: 31 confirmed bugs** (round 15 about 77, 16 about 43, 17 about 15, 18 about 26, 19 about 39). All fixed.

| Lane | Confirmed | Origin |
|---|---|---|
| druid | 6 | old (band 20 and 30 rank rows never re-judged since rounds 15 and 16) |
| hunter | 4 | old (Serpent Sting and Arcane Shot ranged attack power, Immolation Trap attack power, the low-mana order) |
| mage | 4 | old |
| platform | 4 | old |
| paladin | 3 | old (Seal of Righteousness hand split and rank 8, Oath of Light), R18 for the rank-7 row |
| priest | 3 | old (Lesser Heal at 60, Renewed Hope per Penance), guides-handoff (Devouring Plague rank 1 at 28 and 29) |
| shaman | 2 | old (Elemental Devastation toll point), R19 73b7e850e (Maelstrom Weapon tag aura) |
| warlock | 2 | old (Shadow Bolt rank 4 at 30 and 31, the level-50 Succubus weapon) |
| warrior | 2 | old (Flurry charge on a missed Heroic Strike, single-source log), guides-handoff (Fury Berserker Rage and Overpower level gates) |
| rogue | 1 | old plus new-data (poisons scale with attack power, Blizzard 3 November) |

None was caused by the 70755 activation or the 70790 move. Most of what remains comes from two sources the client tables cannot show: server-side coefficients Blizzard posts in its notes (poisons, stings, Arcane Shot, traps), and rank rows written on the gear of their day. Three more lanes (warlock, rogue, warrior) were finished by a second integrator after the first hit an API rate limit; nothing merged was redone.

Also fixed by the integrator: **Thorns on the bear** (druid B4, a preset gap rather than an engine bug): the bear tank's kit now carries his own Thorns from level 6 (S eba3eb7cc), +0.7 to +1.0% threat at raid 60.

## Branches and regen

- **Fork `int/r20`:** **d5c708834**, pushed. It merges fork master abb4e6b8e (the stamp-only 70790 follow) and every lane. Fork master is fast-forwarded at landing.
- **Site `int/r20`:** every lane, plus origin/main 6beb919c2 (merge ac7638941).
  - Merge conflicts: the hunter and Retribution card notes kept round 20's logic and notes, with guide round 18's level qualifiers re-applied ("from 18", Resourcefulness at 33 and 34, Expose Prey at 35 and 36, Clever Traps "once in", the Wing Clip line). Retribution keeps round 20's split opener; main's "rank 7 from 58" note is stale after the paladin lane's rank-8 finding.
  - Guides: Arcane and Elemental keep round 20's talent digits under the 70790 prefix; Discipline keeps round 20; Restoration keeps round 20's Healing Touch ranks with round 17's Nature's Swiftness sentence; Enhancement keeps round 18's Mana paragraph plus round 20's "never needs to wait" clause; Destruction keeps the warlock lane's 30 to 31 sentence (prefix restored, dee593c7a).
  - apl-sync and apl-check are clean on every copy.
- **Pin:** 78f038b47 (engine d5c708834). The vendored proto is unchanged; only ENGINE_SHA moved.
- **Regen once at the pin (0598dea66), for 1.60.1.70790:** ladder goldens, addon-builds (Arcane and Elemental orders), addon-data and Data.lua (card notes), dps-per-point, coach starters. The planner vectors came out unchanged. web `sync:raid-preset` runs before every web build (generated files are ignored).
- **Test fixed by the integrator (effbda8d0):** the enginetalents reposition fixture named node 104765 "elemental_alacrity"; 104765 is Elemental Warding. It passed while the two talents had the same rank and failed once Elemental Warding went to 2 points.
- **Checks run locally:** Go vet (sim, api), the sim leveling, rankchoices, internal and request packages (ladder goldens regenerated, not re-run), data test_apl, test_addonrotation and test_addondata (110 passed), addon-data and coach-rules `--check` clean, the guide prose test. LAUNCH_MODE is on, so the land branch runs the smoke gates only.
- **Not regenerated:** the BiS. A BiS run for 70790 is on its own branch; the rogue BiS (attack power weighs more now), hunter BiS and Holy paladin picks move at that re-rank.

## Before/after, raid 60

Before: round 19's landed engine and site (`before.tsv`). After: int/r20 at the pin, client 1.60.1.70790. Both on the same frozen published BiS, each side's guide raid builds, 10,000 iterations at seeds 1 and 50001, both factions. Seed 1 value, then the change at seed 1 / seed 50001.

| Spec | Faction | Seed 1 | % |
|---|---|---|---|
| hunter-marksmanship | alliance | 801.5 to 934.2 | +16.56 / +16.51 |
| hunter-marksmanship | horde | 801.2 to 934.4 | +16.63 / +16.55 |
| hunter-beast-mastery | alliance | 782.8 to 895.0 | +14.33 / +14.35 |
| hunter-beast-mastery | horde | 779.6 to 893.0 | +14.55 / +14.57 |
| paladin-holy (HPS) | alliance | 671.7 to 727.1 | +8.25 / +8.34 |
| paladin-holy (HPS) | horde | 655.6 to 709.3 | +8.19 / +8.29 |
| priest-discipline (HPS) | alliance | 506.5 to 539.9 | +6.59 / +6.62 |
| priest-discipline (HPS) | horde | 495.7 to 536.2 | +8.17 / +8.15 |
| rogue-assassination | alliance | 726.6 to 777.3 | +6.98 / +7.00 |
| rogue-assassination | horde | 729.1 to 780.7 | +7.08 / +7.22 |
| rogue-combat | alliance | 774.0 to 798.1 | +3.11 / +3.14 |
| rogue-combat | horde | 776.6 to 800.9 | +3.13 / +3.13 |
| rogue-subtlety | alliance | 788.4 to 809.9 | +2.73 / +2.74 |
| rogue-subtlety | horde | 782.7 to 804.3 | +2.76 / +2.75 |
| hunter-survival | alliance | 860.2 to 877.0 | +1.95 / +1.94 |
| hunter-survival | horde | 860.5 to 876.6 | +1.87 / +1.88 |
| mage-arcane | alliance | 721.0 to 729.5 | +1.18 / +1.28 |
| mage-arcane | horde | 708.3 to 716.8 | +1.20 / +1.14 |
| druid-feral-bear (threat) | alliance | 1200.3 to 1209.1 | +0.73 / +0.83 |
| druid-feral-bear (threat) | horde | 1194.2 to 1206.5 | +1.03 / +0.85 |
| paladin-retribution | alliance | 785.0 to 789.7 | +0.60 / +0.60 |
| paladin-retribution | horde | 792.4 to 797.0 | +0.58 / +0.59 |

Every other spec is unchanged to the decimal at raid 60 (the 70790 move bakes in hotfixes the 70755 re-layer already carried). Round 20's other fixes act below 60: rank rows (druid, priest, warlock), the level-50 demon, the Fury card gates.

## Open, for round 21

- **Live logs.** Launch was 23:00 UTC; at 02:00 UTC api.foreversixty.gg still listed only the 20 October report and foreverlogs.gg answered "Report not found" for every id probed, the beta archive included. Warcraft Logs stays blocked (gcloud login). The items waiting on a log: Ice Lance at 60, the hawk's swing, Heart of the Lion on the pet, Eureka! on Incinerate and on Gnome Shadow, demon crit, Soul Siphon, Death Mark and the Mutilate/Shank off hand, Maelstrom Weapon's proc rate, Eclipse on a missed Wrath, Ferocious Bite's attack power share, Deep Wounds' double multiplier, Seal of Righteousness rank 8 procs, Oath of Light on a low rank, the doubled poison coefficients, and the single-source Immolation Trap 2% and Flurry-on-miss.
- **Threat constants** with no client column: bear Maul/Swipe/Primal Bite/Lacerate, warrior Revenge/Shield Slam/Thunder Clap.
- **Shared, for platform:** Eureka! with Execute (forever-bugs #439, one reporter); Arcanite Champion's proc rate; the engine reports a spell as known below its trainer level (warrior X1), which breaks any future `spellIsKnown` level gate; Skyborne launch offsets (P1a); a fork comment in `sim/shaman/talents.go` names Elemental Alacrity's node as 104765 (it is 104771).
- **Guide lane:** the Feral bear page could tell the bear to put Thorns on himself before the pull; the Fury Bloodrage card wording (warrior lane, optional).
- **Measure at refresh:** each lane's list (rogue Instant main hand on the new BiS, Affliction Malevolence string, Subtlety band 20, Enhancement band 40 reserve and relic near tie, Discipline Gnome vs Human, the band 20 and 30 rank rows on re-ranked gear).
- **Pending owner (not reopened):** as in the brief.

## Rendered cards after the regen (wow-player)

The five cards round 20 changed most were captured at 1440 on the int/r20 build after the regen: Beast Mastery, Marksmanship, Survival, Discipline and Fury.
- **Beast Mastery, Marksmanship and Fury: SHIP.** The hunter reviewer's request for one more look at the three hunter cards after the regen is done here.
- **Survival: SHIP with one MUST that predates this round.** The page says a Survival hunter marks the boss for Expose Prey, but neither the opener nor the priority list shows Hunter's Mark. The engine already applies the mark through the Survival kit, so no number is wrong; the opener line goes to the hunter lane in round 21 with its note.
- **Discipline: the reviewer's two data checks are aligned.** Greater Heal rank 5 (25314) is a level-60 book rank, which stays with the owner, so the card's rank 4 is the shipped rule. Fear Ward has no race restriction in the 70790 trainables (Holy, level 20). Lesser Heal is gone from the card and the "Low ranks" paragraph agrees.
- **Fury:** "Berserker Rage From 36" is deliberate. The spell trains at 30, but the button needs Berserker Stance, which a Fury warrior uses only from 36 (warrior lane, WARRIOR-R19-1).
- **Shoulds for the guide lane:** BM's sting and Arcane Shot lines could say "the shared cooldown" as Marksmanship does; Intimidation, Inner Focus, Rapid Fire and Deterrence rows carry only a cooldown; the Marksmanship raid-only note could name Sniper Shot; the Rapid Fire row and paragraph word the same moment two ways; Discipline rows 9 and 10 could give a health threshold; Fury's Execute "First" could read "Before any other attack".
