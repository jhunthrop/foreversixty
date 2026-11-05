# Sim-accuracy round 21: integration (2026-11-05)

Round 21 started from the landed round 20 (site 93c66ebc8, fork d5c708834) on the live client 1.60.1.70790. One lane per class, plus platform and a late items lane. Leads came from the seeds, guide-accuracy rounds 19 to 21, and a tier-list comparison against MythicSim. Guide rounds 19, 20, 21 and 22 landed on main during the round and are merged in.

**Not a zero round: 41 confirmed bugs, all fixed** (round 15 about 77, 16 about 43, 17 about 15, 18 about 26, 19 about 39, 20 was 31). The count rose because this round read Blizzard's notes line by line against the engine, added item effects, and took late leads from the guide lanes and the comparison.

| Lane | Confirmed | Origin | What |
|---|---|---|---|
| hunter | 8 | 6 old, 1 R20, 1 R21 | Auto Shot fires through Aimed and Sniper Shot casts; the Aimed and Sniper gates; Hunter's Mark before the pull on all three cards; Survival Raptor Strike rank rows; the low-mana stall at every level |
| druid | 6 | 5 old, 1 new-data | Bear Form threat 115%; Swipe attack power; no single-target Swipe; Maul at 30 rage at 40 to 53; Shred from 21; Balance Wrath filler below 20 |
| priest | 5 | 5 old | Holy Heal rows at 28 to 29 and 34 to 39; Lesser Heal filler below 16; Discipline tank Heal rows; Starshards under Twin Disciplines |
| items | 4 | 2 old, 2 new-data | Hand of Justice 1%; Storm Gauntlets; Warblade of Caer Darrow; Blade of Hanna |
| mage | 3 | 3 old | Fire filler Fireball to 29; Arcane Presence of Mind on Pyroblast from 54; Evocation from its trainer level |
| paladin | 3 | 3 old | Consecration rank 4 at 53 for Dwarf and Undead; Sanctified Judgement's mana makes no threat; the Seal twist paced on mana, not attack power |
| platform | 3 | 2 old, 1 new-data | Eureka!'s saved rage makes no Execute damage (single source); Defensive Stance tooltip override; demons by trainer level |
| shaman | 3 | 1 old, 2 new-data | Searing Totem on packs below 26; Maelstrom Weapon chance doubled; Lightning Bolt and the reserve below 5/5 |
| warlock | 3 | 3 old | Conflagrate hold before Shadow and Flame; Amplify Curse on each Agony below 60; Demonic Brand modeled |
| rogue | 2 | 2 old | Subtlety Shank hold only at 50 energy or more; finishers below 20 without Rupture |
| warrior | 1 | 1 new-data | Defensive Stance threat 80% |

By origin: 32 old, 7 new-data, 1 R20, 1 R21 (the hunter lane's own below-34 scope went stale when Aimed Shot was ungated). The lanes' reports carry the evidence.

## New evidence class: Blizzard's notes against the cached rows

Blizzard's 3 November notes carry server changes the cached 70790 rows do not (Bear Form 115%, Defensive Stance 80%, Maelstrom Weapon doubled). Every lane read its class's lines from 24 September to 4 November. Defiance's 7/13/20 is the exception: the same note, but the client and Wowhead still show 5/10/15, so it stays OPEN.

## Integration

- **Fork:** int/r21 = 170a59c39 (pushed). **Site:** pin 8d70ef3e5.
- **Merges from main:** e039e1490 (guides 19 and 20), 987d08a7c (21) and 24ee457d0 (22). Conflicts kept the lanes' logic on the guide rounds' wording. One carried decision: main's Destruction "until Shadow and Flame reaches 5 points" gives way to the warlock lane's measured first-point rule.
- **Integrator fixes:** the Arcane Pyroblast gate carries its client rank (test_apl); Evocation is level-gated in the ladder and spellranks tables, not declared inert (the smoke test expected it unresolved at 60); the smoke build explains Shadow and Flame's aura; the Holy Heal rank test follows the new rows; a guard test pins three weapon speeds.
- **Weapon speeds (guide round 21 priority lead): not a bug.** The sim and the BiS ranker read simdb.bin. All 761 weapons with a 70790 ItemSparse row match its ItemDelay; the other 1175 come from Wowhead's Forever data, with no client row to compare. The fork's vanilla table feeds only the fork's own tests. Damage ranges were not compared with the client.
- **spellIsKnown:** the platform's RequiredLevel check was shipped and reverted in this round. RequiredLevel is the client spell level, not a learn level; registration is the only meaning of known.
- **Regen at the pin (clean disk, 369342be0):** ladder goldens, addon-data and Data.lua, dps-per-point, coach starters. Addon-builds and the planner vectors came out unchanged.
- **Disk:** the Mac's disk filled around 05:14 to 05:25 UTC. The first regen ended at 05:12:53, and the rerun on the clean disk changed only what the late hunter and card fixes explain.
- **Checks run locally:** data test_apl, test_addonrotation, test_addondata, test_genproto; addon-data and coach-rules `--check`; luac on Data.lua; sim request, leveling, rankchoices, internal; sim and api vet; prose and rotation vitest; prettier on the guides; apl-check. LAUNCH_MODE is on, so the land branch runs the smoke gates only.
- **Fork suite, failures that predate this round or are host-only:** `sim/web` needs a built binary_dist; TestElemental's stat weights fail the same on d5c708834; the conformance goldens differ in the fourth decimal on this arm64 host. Only rows the fixes moved were rewritten (d8d71c2a6), so the amd64 values stay. That is expected to pass on amd64 and was not run there.
- **Not regenerated:** the BiS (its own run), and the Defensive Stance tooltip, which needs the nightly data job's `spelltips` step. The hover reads 30% until that runs.

## Before/after, raid 60

Before: round 20's landed engine and site (`before.tsv`). After: int/r21 at the pin. Both on the same frozen published BiS, each side's guide raid builds, 10,000 iterations at seeds 1 and 50001, both factions. Seed 1 value, then the change at seed 1 / seed 50001. Tanks are threat per second.

| Spec | Faction | Seed 1 | % |
|---|---|---|---|
| druid-feral-bear (threat) | alliance | 1209.1 to 1756.3 | +45.26 / +45.38 |
| druid-feral-bear (threat) | horde | 1206.5 to 1750.6 | +45.10 / +45.27 |
| warrior-protection (threat) | alliance | 1088.6 to 1501.1 | +37.89 / +37.90 |
| warrior-protection (threat) | horde | 1079.1 to 1488.0 | +37.89 / +37.88 |
| hunter-marksmanship | horde | 934.4 to 989.5 | +5.90 / +5.84 |
| hunter-marksmanship | alliance | 934.2 to 988.5 | +5.81 / +5.84 |
| shaman-enhancement | horde | 859.1 to 864.5 | +0.63 / +0.76 |
| paladin-retribution | alliance | 789.7 to 792.3 | +0.33 / +0.34 |
| paladin-retribution | horde | 797.0 to 799.4 | +0.30 / +0.33 |
| hunter-beast-mastery | horde | 893.0 to 891.0 | -0.22 / -0.25 |
| hunter-beast-mastery | alliance | 895.0 to 893.3 | -0.19 / -0.04 |
| shaman-enhancement | alliance | 856.6 to 858.0 | +0.16 / +0.07 |
| rogue-subtlety | horde | 804.3 to 805.3 | +0.12 / +0.10 |
| rogue-subtlety | alliance | 809.9 to 810.5 | +0.07 / +0.11 |
| hunter-survival | alliance | 877.0 to 876.7 | -0.03 / +0.07 |
| hunter-survival | horde | 876.6 to 876.5 | -0.01 / -0.09 |

Every other spec is unchanged to the decimal at raid 60. Most of this round's fixes act below 60, on bare bands or off the published gear.

## Open, for round 22

- **Live logs.** At 05:27 UTC the site API still listed only the 20 October report and no Warcraft Logs host listed a Forever zone (v1 API with the key in Secret Manager; its web pages answer 403). The platform lane found foreverlogs.gg's public list empty at 03:02 UTC, and later lanes got no further. Each lane's "Still open" names the log it needs. Single-source fixes to re-check on a log: Eureka! with Execute, Demonic Brand's damage, Maelstrom Weapon 100% against a rate per minute, Flurry on a missed Heroic Strike, Immolation Trap 2%.
- **Threat constants with no client column:** Maul, Swipe, Lacerate, Revenge, Shield Slam, Thunder Clap. Defiance 7/13/20 waits on the client.
- **Shared, for platform:** Rend and Tear on other players' bleeds in the opener; Arcanite Champion's proc rate; Skyborne launch offsets; spells registered early but never cast by a shipped card (rogue Feint and Vanish, mage Counterspell, shaman Tremor Totem; the platform report has the rest); the "Fireball RANK 11" beside "Below 30" on the Fire raid card (mage lane flag; not yet checked after the regen).
- **Hunter:** Beast Mastery's order (Multi-Shot above Aimed Shot is +0.9% on the raid build and loses 1.5 to 5% on the leveling bands; a rotation cannot tell the builds apart at 60).
- **Talent and rotation leads from the MythicSim comparison, under the adoption rule:** Affliction +2.7%, Demonology +3.5%, Fire +1.7%, Enhancement +1.25%, Balance +2.35%, Assassination +2.6%, Subtlety +2.2%, Shadow +4.0% without Vampiric Embrace upkeep. Unverified here (`scratchpad/competitor/lanes/`).
- **BiS and method, for the coordinator:** the published band-60 BiS carries stale talents and races; the ranker sims with no permanent enchants (casters look under-ranked about 4 to 5% against melee, by the comparison's figures); the ranker drops client_unconfirmed items. The `fix/bis-unconfirmed` branch appeared while this ran.
- **Measure at refresh:** each lane's list. The biggest: bear and Protection weights and picks on the new threat, Hunter weights on Auto Shot through casts, Retribution gear under the paced twist, Arms Warblade and Retribution Hanna at bare 60, Enhancement bands 40 to 60 and its relic.
- **Guide lane:** fury.md:112 "edges Night Elf by a hair" should read level (Skyborne +0.02/+0.03%); Elemental's three wording SHOULDs from the card review.
- **Pending owner (not reopened):** as in the brief.

## Rendered cards after the regen (wow-player)

Eight cards were captured at 1440 on int/r21 after the regen. Two wow-player passes.
- **Enhancement, Elemental, Bear: SHIP** on the first pass. The Enhancement re-capture was the condition on the shaman lane's SHIP: the half-mana reserve is gone from the Lightning Bolt and Earth Shock rows.
- **Survival: SHIP.** Its Hunter's Mark note now names Expose Prey in the part the card prints.
- **Destruction and Retribution: NO-SHIP, then SHIP.** The Conflagrate row now carries the hold before Shadow and Flame; the Seal of Righteousness and Consecration rows carry the mana pace.
- **Beast Mastery and Marksmanship: NO-SHIP twice, then fixed.** The first pass found the low-mana contradiction, which exposed a real stall (hunter C1). The second found the card printing only the first sentence of each note, so the low-mana rule never reached the Serpent Sting and Arcane Shot rows. The integrator moved the reviewer's wording into those sentences and read the re-captured Beast Mastery card; the same text is on Marksmanship. **No third wow-player pass was run on these two cards.**
- **Not re-reviewed after their last edits:** the Destruction comma and the Survival note wording.
- **Not applied:** Elemental's three wording SHOULDs (guide lane).
