# Sim-accuracy round 26: integration

Scoped round in ship mode: round 25's open items plus a fresh audit of the two least recently audited specs. Subtlety and Arms had the fewest mentions in the round 23 to 25 reports, 19 and 20; every other melee spec had more. Local checks covered only the touched tests and the guide guard; CI runs the full suites.

Bases: site int/r26 from main eb64ba7e8, fork int/r26 from f871a65ae. Engine pin after: **227d95bd1**.

## Confirmed bugs: 19

| lane | fixed | by origin | report |
|---|---|---|---|
| platform | 5 | 3 old, 1 R21 559e39867, 1 R25 f5d1cd239 | platform.md |
| melee | 3 | 3 old | melee.md |
| ranks | 11 | 6 R24 (rows judged on the old below-band gear), 5 old | ranks.md |
| **total** | **19** | **11 old, 6 R24, 1 R21, 1 R25** | |

What changed, in a player's words:
- **Gear below 20 (platform):**
  - The leveling ladder handed out items only the other faction can earn, such as Crescent Staff to Alliance casters and Salbac Shield to Horde shamans. It now reads the faction from the race.
  - Empty armor, ring and trinket slots at 10 to 19 are filled from the class item table, by level, armor type and faction.
  - Quests whose faction the loot data marks unknown are open to both factions, as the BiS ranker already treated them.
- **Other platform fixes:**
  - The fork's `go vet ./sim/core` is clean.
  - Deep Wounds at one and two talent points (12834, 12849) had no icon on the result page. The icon itself was right.
- **Subtlety:**
  - Ghostly Strike uses normalized weapon damage, as the 70755 client row says (SpellEffect 1368764, the same shape as Hemorrhage).
  - Initiative comes at 24, 25 and 27, ahead of two points that do nothing in a fight.
- **Arms:** Improved Overpower comes at 15 and 16, ahead of Improved Tactical Mastery, which does nothing before the stance swap at 50.
- **Combat:** Slice and Dice waits for 60. Below raid gear, skipping it and spending every five points on Eviscerate wins at every bare level from 10 to 60, by 0.5 to 6%, beyond error at both seeds and factions, with daggers or a sword. Raid 60 keeps Slice and Dice, where it wins by 2.4%.
- **Assassination:** below 17 it casts Slice and Dice at four and Eviscerate at five.
- **Caster and healer rank rows at 10 to 19**, re-judged on the corrected gear:
  - Balance: Wrath and Moonfire.
  - Restoration druid: tank Rejuvenation rank 1 from 13, plus party and Healing Touch rows.
  - Protection paladin: Holy Strike rank 2 at 12 to 15.
  - Holy priest: Renew rank 2 at 16.
  - Discipline: new Power Word: Shield and Lesser Heal rows.
  - Elemental: Flame Shock rank 1 at 18 and 19, Dwarf only.

  Holy paladin and Restoration shaman are aligned.

## Before/after

Raid 60, frozen gear, guide builds, 10,000 iterations, seeds 1 / 50001. Only the specs the round's engine and rotation changes touch were run (ship mode):

| spec | alliance s1 | alliance s50001 | horde s1 | horde s50001 |
|---|---|---|---|---|
| rogue-subtlety | 808.7 to 804.1 (-0.57%) | -0.57% | 802.9 to 798.4 (-0.56%) | -0.56% |
| rogue-combat | 792.1, unchanged | unchanged | 794.5, unchanged | unchanged |
| rogue-assassination | 807.0, unchanged | unchanged | 804.6, unchanged | unchanged |
| warrior-arms | 844.3, unchanged | unchanged | 846.8, unchanged | unchanged |

Below 60 the gains are in the lane reports. Examples: Arms bare 20 +3.3%, Subtlety bare 30 +0.9 to 1.1%, Combat bare 30 +4.1 to 4.8%, and the rank rows +0.2 to 5.9%.

## Regenerated (S 1dd256eb1)

regen790.sh at pin 227d95bd1, then `pipeline icons` on the fresh 70790 tables. Outputs:
- the ladder goldens for 20 specs;
- addon-builds;
- dps-per-point;
- addon-data and the rotation and build rows in Data.lua (no BiS pick moves);
- the Combat coach rules;
- the three Deep Wounds rank icons in spellicons.json.

## Re-rank needed (one-spec bis.yml runs after landing)

- **rogue-subtlety, every band:** Ghostly Strike's damage changes, so weapon picks can move.
- **rogue-combat, bare bands 20 to 50:** Slice and Dice no longer runs there.
- **warrior-arms, band 20:** the leveling order changes at 15 and 16.

## Measure at refresh

- MR35: Combat bare 20 to 50 numbers rise by the margins in ranks.md once re-ranked; Subtlety picks after the Ghostly Strike change.

## Open, for round 27

- Subtlety's Opportunity against Master of Deception at 10 to 12, and the Arms Deflection order at 11 to 13.
- Arms A5: an Overpower clears its own Bloodthrill proc. It needs a log.
- Combat at bare 60 loses 0.5% with Slice and Dice: an APL cannot tell bare from raid at the same level.
- Skyborne druids below 20 were not re-judged.
- `spec-breakdown -race` accepts a race the class cannot be.
- The fork's class test suites crash on Era test items missing from the Forever item database (12620, 12640). Pre-existing.
- 54 engine-only spell ids have no icon source on the result page.
- The `tools/coach/starter_overrides.py`:489 comment names the old Initiative level. It is a generator file, so it goes to the coach owner.
- Logs: at 21:49 UTC 6 Nov there was still only the 20 Oct test report.
