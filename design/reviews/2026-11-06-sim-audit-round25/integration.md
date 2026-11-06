# Sim-accuracy round 25: integration

Scoped round, in build mode. The coordinator set four seeds and added two items during the round:
- the race sweep in the BiS ranker;
- the coach items (F1 talent-shortened cooldowns, F7 Evasion, the Shadowburn shard floor);
- the `internal/rating` presence test;
- rogues and hunters unarmed below 19 to 20;
- added: healer bare-band drops (healers-bare.md);
- added: Shadowform at 40 against Wand Specialization (shadow.md).

Bases: site int/r25 from main ab2029586, fork int/r25 from bf727699f. Engine pin after: **f871a65ae**.

## Confirmed bugs: 23

| lane | fixed | by origin | report |
|---|---|---|---|
| coach | 9 | 9 old (4 guides-handoff) | coach.md |
| platform | 13 | 7 R21 559e39867, 2 R24 2ae0fe3b8, 4 old | platform.md |
| race | 1 | 1 old | race.md |
| shadow | 0 (its one bug, casters with no wand below 17, is platform P4) | - | shadow.md |
| healers-bare (integrator) | 0: deliberate, fork 8d9ca6aae | - | healers-bare.md |
| **total** | **23** | **14 old, 7 R21, 2 R24** | |

What changed, in a player's words:
- **Coach, F1:** the core-cast count now uses the cooldown your own talents give, read from your stored build (addon export or Battle.net import). The five affected spells:
  - Fury Whirlwind, 8 s with Improved Whirlwind;
  - Rapid Fire on the hunter leveling builds;
  - Fire Blast, 6 s from 26;
  - Retribution Judgement, 8 s from 16;
  - Protection paladin Holy Strike, 6 s from 40.

  Shadowburn is only judged once a warlock has Shadow and Flame 5/5, or from 44 when the talents are unknown, so missed casts without a shard no longer count.
- **Coach, F7:** a melee-only defensive (Evasion, Deterrence) is only offered when the killing blow was physical.
- **The `internal/rating` presence test:** the test was wrong, not the handler. eaecf8cc3 added two visibility checks that read no evidence and never put them in the allowlist.
- **Platform:** below 20 the band-20 gear lost every item too high-level for the character and put nothing in its place, so rogues fought unarmed, hunters had no bow and casters had no wand below 17.
  - The refill now draws from the client tables: weapons, shields and wands the level can hold, without the other faction's quest rewards.
  - A gate on a spell the character has not learned now reads "cannot cast". Before, it dropped out of its gate, so hunters below 18 never cast Serpent Sting or Arcane Shot.
  - The rank rows judged on the new gear: Beast Mastery and Marksmanship rank 1 shots at 13 to 19; Feral Wrath and Moonfire at 10 to 15, keyed by race.
  - The hotfix-push alert can open its issue now (its label text was over GitHub's limit), and a push record that layered nothing is flagged.
- **Race:** Protection warrior's Horde pick is Undead, not Tauren. Touch of the Grave's drain counts as threat: +1.65 / +1.97% in the raid and +2.29 / +2.42% bare at 60 (seeds 1 / 50001, error about 0.08 points). Tauren stays named as the health pick.

## Not bugs

- **Healer bare bands** (Restoration shaman bare 40 to 60 -26 to 33%, Restoration druid bare 60 -23%, Discipline -17%): fork 8d9ca6aae, round 15's level-gap penalty (owner law). The published files predate it (`engine_version d1d9f2775`). Restoration shaman had lived on Healing Wave rank 2 with its whole bonus. Bisect table in healers-bare.md.
- **Shadowform at 40:** no order gives both Wand Specialization before 40 and Shadowform at 40. A priest has 31 points at 40 and the form needs 30 in Shadow first. Every Shadowform-40 order is slower over levels 10 to 41 (0.2 to 0.8% by kills, about 4% by level), so the order stays. The guide now tells the player to reset at 40 for Shadowform and buy the wand points back at 41 and 42, which is what Forever's leveling guides do.

## Before/after

- **Raid 60, 28 specs, both factions, seeds 1 and 50001, 10,000 iterations, frozen gear, guide builds:** 112 values, none moved more than 0.05% (before.tsv against after.tsv). Every change this round sits below 20, in coach code or in a guide race line. The frozen gear keeps Tauren for Protection warrior, so its race gain is the race lane's own measurement above.
- **Below 20 (platform):** Shadow Horde at 10 / 13 / 16 goes from 5.77 / 6.42 / 8.58 to 13.09 / 16.28 / 30.48 DPS. Hunters and rogues now fight armed. Per-row tables are in platform.md.

## Regenerated (S bc41d4f2f, regen790.sh at pin f871a65ae)

The Beast Mastery and Marksmanship ladder goldens and the addon rotation rows (`addon-data.json` and the rotation part of Data.lua). No BiS pick, coach starter or planner file moved.

## Landing notes

- **`TestPublishedBISMatchesGuides`:** warrior-protection is in `knownStaleBISFiles` (S 74d009e9c) until its Horde set is re-ranked on Undead. Priest-discipline was already red on main and clears when land/bis-final2 lands.
- **Coach rules version:** `RulesVersion` is now coach-rules-v3, so each cached coach pack rebuilds once on its next view, one model call each.
- **Data.lua:** after land/bis-final2 lands, merge main and rerun addon-data so the BiS picks in Data.lua are theirs.

## Measure at refresh

- MR32: the published `race_sweep` fields (best race, margin, error) in every band report. Check that no free race beats a guide race beyond error by over 1% at 60.
- MR33: Protection warrior Horde re-ranked on Undead (and the known-stale entry removed).
- MR34: Shadow's published band 20 to 40 talents follow the leveling order, and the Alliance race is Human.

## Open, for round 26

- 20 caster and healer rank rows at 10 to 19 were judged on the old below-band gear and need re-judging: Balance 4, Restoration druid 8, Holy paladin 1, Protection paladin 2, Discipline 1, Holy priest 1, Elemental 2, Restoration shaman 1.
- Combat and Assassination at 14 to 17 never reach a five-point Eviscerate: Slice and Dice is refreshed at one or two points.
- The ladder's own gear pickers ignore faction (moves goldens). Armor slots below 20 are not refilled.
- Leveling-band race leads over 1%: Fire Gnome, Shadow Gnome, Holy priest Night Elf at 30 (driver untraced), Protection paladin Human at 30, Combat Night Elf. Warlock Undead at 20 stays pending with the owner.
- Fury execute hold (WARRIOR-R26-10) and Bloodrage ceilings (WARRIOR-R26-11).
- The fork's `go vet ./sim/core` fails on `test_generators.go:535` (pre-existing).
- Retribution's band-20 gear at 16 runs out of mana for 20 s.
- Logs: GET https://api.foreversixty.gg/v1/reports/recent at 16:43 UTC 6 Nov lists only the 20 Oct test report.
