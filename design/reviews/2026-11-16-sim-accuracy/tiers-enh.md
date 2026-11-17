# Sim accuracy 2026-11-16: the tier list, Enhancement at raid 60 and Feral on top

Lane brief: `design/reviews/sim-accuracy/brief.md`. Site from origin/main eced5df0b, engine pin f9385b498, client 1.60.1.70938. Trigger: a player on /tiers, "Enhancement shaman #10? More like #1-3". Scope: is the raid 60 order wrong, Enhancement's card, Feral's place at the top, and whether the comparison is fair.

**Totals: 0 engine bugs, 1 text fix held (the log coach still told a raid shaman to use Rockbiter Weapon under a Windfury Totem and had no Flame Shock line). No sims ran: the Windows box went unreachable mid-lane (ping and port 22 dead, the Mac's gateway too), and the brief forbids sims and tests on the Mac. The ranking holds on every source that exists; nothing measured at level 60 contradicts it.**

## Evidence

| source | what it is | Enhancement | Feral | verdict |
|---|---|---|---|---|
| Warcraft Logs | no Forever partition found (every subdomain 403 to scripts; no zone page indexed) | none | none | no data |
| foreverlogs.gg `/api/statistics?phase=2` | measured parses | Zul'Farrak median above Marksmanship and Retribution (n 7 to 13 a spec); Maraudon second behind Demonology (n up to 8); one Enhancement tops the site's single-kill list in a mid-level dungeon | no cat parses at those levels | small leveling-dungeon groups; no rogue, warrior or cat to compare; every level 60 zone (BRD, LBRS, UBRS, Scholomance, Stratholme, Dire Maul) has zero parses |
| our reports (Cloud SQL and the logs bucket, read only, after the login was renewed) | measured, advanced logging | none at 55 to 60: all 11 reports (10 complete, 1 live, 20 Oct to 16 Nov, about 1,380 fights) read with `tiers-enh/own60.py`; level is each player's own advanced-log level; the highest player in any fight is 51, and only 10 damage-dealer rows in 3 reports sit at 40 or above (one rogue at 47, two warriors at 44 and 51, a Retribution paladin at 45, a Fury warrior at 51, a priest and a druid at 41 to 42, no shaman, hunter or cat) | - | no data at 55 to 60; nothing to rank |
| brokenmeta.gg tier list | another sim, level 60, 180 s, level 63 boss, best datamined gear from dungeons, raids and crafting, client 70753 | 10th, two-hander and Windfury Weapon, its notes still quote Stormstrike's pre-hotfix 150% | 11th | agrees on Enhancement; its Feral predates 70755 |
| kami-labs Forever Enhancement page | theory | the 10 Nov hotfix roughly halves Stormstrike; no measured DPS anywhere | - | argues against top 3 |

The player's feel matches leveling dungeons: Enhancement leads small Zul'Farrak and Maraudon groups, where there is trash to cleave, no raid buffs and no rogues, warriors or cats logged. That is not the raid 60 single-target table /tiers ranks. Our own uploads have no player above 51. Both level 60 sims that exist put Enhancement 10th; BrokenMeta's would be lower still on the 8 s, 100% Stormstrike our engine carries.

## Enhancement's raid card

- Weapon: The Unstoppable Force (two-hander, Stormpike exalted), no off hand, so the shamans-never-dual-wield rule holds. Feral wears the same weapon.
- Imbue and totems: the kit's `main_hand_imbue:windfury_weapon` wins the slot over the preset's sharpening stone (`ResolvedPreset.Layer`, kit first), and the raid's Windfury Totem reaches the shaman beside it (the 10 Nov hotfix, measured +10.8% over Rockbiter on 13 Nov).
- Talents (raid): Concussion 5, Elemental Devastation 3, Reverberation 5, Elemental Focus; Thundering Strikes 5, Ancestral Knowledge 4, Mental Dexterity 3, Elemental Weapons 3, Shamanistic Focus, Flurry 5, Stormstrike, Mental Quickness 2, Maelstrom Weapon 5, Rage of the Farseer; Tidal Focus 5 (the hit). Every Forever rework the build takes is registered in the fork (Rage of the Farseer as a cooldown, Mental Dexterity, Maelstrom Weapon's per-hit rate). Improved Stormstrike, Anticipation and Spirit Weapons do nothing on a boss that does not hit the shaman.
- Hit: 9% with Tidal Focus, at the two-hander's cap, so the hit weight is 0 for melee; spell hit is separate.
- Rotation: Stormstrike, Flame Shock above 30% mana, five-stack Lightning Bolt, Searing Totem, Fire Nova on Clearcasting, Earth Shock above 15%; the 13 Nov lane's searches found nothing better beyond error.
- Expertise: Forever has it (item stat 37), and the shaman sits at 0 against the boss's 6.5% dodge, as do Fury, Arms, Feral, Retribution and Subtlety (Combat and Assassination start at 2 from talents). No damage spec's weights include expertise, only the three tanks', so any pre-raid expertise item is valued at zero for every melee spec alike: fair between them, open as O5.

## Feral on top

No bug found. The cat leads because Forever's druid is tuned up: Naturalist +10% to all damage, Rend and Tear +15% on bleeding targets, Genesis +10% periodic, and Shifting Power (40 energy on an 8 s cooldown for 55% of base mana; `shifting_power.go`, SpellCooldowns 101953). BrokenMeta's 70753 tables predate the 8 s cooldown (70755), which by itself explains most of why it has the cat 11th. The 11 to 13 Nov lanes checked Rake, Ferocious Bite, Rip, Shred and the energy against our logs. Open: Shifting Power's cast count over a raid 60 fight depends on the cat's mana (Judgement of Wisdom, Mana Spring, Blessing of Wisdom in form); no level 60 cat has logged it.

## Fairness

Every DPS spec's raid entry carries the same 14 buffs and 8 debuffs, the 180 s fight, a level 63 target, one target, 20% variation and 25% execute (`run` block identical). Consumables differ by role only as the preset notes state (potion per spec by measurement, one Blasted Lands buff per spec).

## Fixed, held, open

- **Held (text, needs the box for the coach tests):** `data/curated/coach/shaman-enhancement.md`, the log coach's persona, still said Rockbiter Weapon under a raid Windfury Totem and lacked the Flame Shock line the card casts since 13 Nov. Committed on branch `lane/tiers-enh-coach` (pushed, not on main); land it once `go test ./internal/coach/` passes on the box.
- **Open O1:** the box is unreachable; the raid 60 breakdowns for Enhancement, Feral, Fury and Combat (`tiers-enh/breakdown.sh`) did not run.
- **Closed O2:** our own uploads read once the login was renewed: no player at 55 to 60 (highest 51).
- **Open O3:** Shifting Power casts and the cat's mana in a level 60 log.
- **Open O4 (13 Nov O1):** Windfury Totem rank 2's hit 35% above an extra swing.
- **Open O5:** damage specs carry no expertise weight (`specs.go` WeightStats); measure whether any pre-raid item with expertise would change a melee pick.

## A reply the owner could paste

Enhancement feels great while leveling, and the dungeon logs agree: in Zul'Farrak and Maraudon it tops small groups, where there is trash to cleave and no raid around it. The tier list ranks something different: one level 63 boss for three minutes, full raid buffs, pre-raid gear. There Stormstrike is 100% weapon damage on an 8 second cooldown since the 10 November hotfix, and the rogues, warriors, hunters and cats scale harder off the same buffs. The only other level 60 sim out there has Enhancement 10th too. Nobody has logged a level 60 raid boss yet; when raids open we will check the list against real kills.
