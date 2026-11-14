# BiS verification round, 1.60.1.70938

Lane `bis-verify3`, ledger step 3 ("A verification round audits the regenerated output"). Scope: the first ranking on client 1.60.1.70938 with the corrected server-hotfix data (data run 43468132768, `273623f0f`): `f8b971e14` "chore(bis): nightly leveling BiS for 1.60.1.70938 @ 5603877e4" (27 specs, run 43507507056 dispatched on `275422dec`) plus priest-holy from `275422dec` (run 43487819219 on `273623f0f`'s head, generated 20:13 UTC). Every file names engine 5603877e4 and was generated between 20:13 and 23:40 UTC, after the data fix landed (18:58 UTC); simdb.bin is unchanged since `273623f0f` (sha256 ea39a6a1...). Baseline for the moves: the carried picks at `7f37bb364` (engine 4c595ee6b, 70790 data). Scripts and job files: `bis-verification-70938/`.

**Verdict: every check passes; no ranker bug confirmed; nothing to re-rank beyond what main's newer engine already owes (MR108, MR109).** Three `knownStaleBISFiles` entries dropped (shaman-elemental, warlock-affliction, mage-frost); shaman-enhancement stays, now for MR108's raid build.

## Method

As in `design/reviews/2026-11-12-sim-accuracy/bis-verification.md`: `checks.py` re-derives slot legality, weapon shape, tier order, the band series and the MR94/MR100 fields from the build's own data; `TestPublishedBISFilesPassSanityChecks` (FOREVER_BIS_SANITY=1) and `TestPublishedBISMatchesGuides` on the box; spot sims through `tools/remote-test.sh --dir sim --engine-sha 5603877e4` at `ce4d0fb5b` (origin/main plus the lane's scripts; the box printed head, build 1.60.1.70938, both simdb.bin hashes ea39a6a1... and the engine 5603877e4 before simming), one `spec-breakdown` at a time, `nice -n 15`, GOMAXPROCS=5, 3,000 iterations, seed 7: 336 published baselines, 224 previous-band sets, 690 emptied-slot sets, then 10 Retribution weapon and legs swaps (`extras.txt`, `extras2.txt`). 1,260 sims in about 600 s.

## Checks

| check | verdict | detail |
|---|---|---|
| slot legality | Pass | 28 specs, 336 entries, 20,140 picks and alternatives: class mask, slot, armor gate at 40, weapon subclass, faction, required level, no duplicated ring or trinket |
| weapon picks | Pass | a main hand everywhere, no two-hander beside an off hand, no shaman off-hand weapon, Survival on melee weapons and melee weights, ranged slots legal |
| set_dps across bands | Pass | rises band to band on both factions for every spec; raid 60 above bare 60 everywhere |
| previous published set (ranker guard) | Pass, 47 holds | `kept_alternative` re-measured on the new engine and data and kept; largest: bear raid 60 Horde (new set -3.9%), Protection paladin bare 60 Horde (-2.8%), Arms bare 60 Horde (-2.1%), Subtlety bare 60 Alliance (-1.3%), Ret raid 60 Alliance (-0.5%) |
| previous band's set at the next band | Pass, 224 of 224 | every bare 30 to 60 entry, both factions, tanks on the full tank score: the band below's set scores under the band's own |
| one slot emptied | Pass | 655 below, 17 ties inside the error, 18 no figure (a hunter without a ranged weapon has no fight). The ties: Ankh of Life in trinket1 at bare 40 for the rogues, Arms, Fury and Retribution (the only trinket the pool holds there; it adds nothing and nothing better exists), Balance bare 40 trinket2, three relic/ranged slots worth under 0.15% |
| tier order | Pass | 5,712 json slots and md rows descend; no measured alternative above its pick |
| published-file tests | Pass | sanity: all 28 active-build subtests; guides: 27 pass, shaman-enhancement fails only on raid 60 talents (`0533001/053030310051202051/053` against the guide's `0535001/054030310051202051/05`, MR108, moved at `0f77fad48` after the ranking started), so its stale entry stays with that reason |
| MR94 health | Pass | `metrics.health` on all 36 tank entries, always below effective health: warrior raid 60 7,572 / 7,275, paladin 7,286 / 7,121, bear 9,171 / 8,882; the tier list's Health column shows |
| MR100 raid_fight | Pass | all five healers, bare and raid 60, both factions (20 of 20) |

### Calibration

The published runs reproduce `set_dps` to the hundredth for every DPS entry except: shaman-enhancement, +0.5 to +5.4 DPS on every band (main's `data/curated/apl/shaman-enhancement.json` gained the Flame Shock card at `0f77fad48`, after the ranking's checkout; MR108's +0.7 to +2.3% bare), and two bare 20 Horde entries 1 to 1.7 sigma off (Balance, Elemental). Healers sit inside 2 sigma (0.3% at most) except Holy priest bare 40 and 50 Horde (-0.27%, -0.29%, 3.3 sigma; the file is the earlier run's on the same engine and data). The Enhancement spot sims compare like with like (both sides on main's card).

### Weapon spot checks (Retribution)

| entry | set | DPS | vs published |
|---|---|---|---|
| bare 60 Alliance | published (Blade of Hanna) | 356.97 | |
| | The Unstoppable Force | 341.40 | -4.4% |
| | Hammer of the Grand Crusader 18717 | 345.96 | -3.1% |
| | Blacksmithing 250617 | 339.37 | -4.9% |
| bare 60 Horde | published (Blade of Hanna) | 357.87 | |
| | The Unstoppable Force | 341.63 | -4.5% |
| | 11921 / 250619 / 19106 | 340.69 / 340.59 / 334.94 | -4.8 / -4.8 / -6.4% |
| | High Warlord's Greatsword (rank 14) | 360.14 | +0.6% |
| bare 50 Alliance | published (Infernal Trickster Leggings) | 209.24 | |
| | Knight's Plate Leggings | 207.21 | -1.0% |
| bare 50 Horde | published | 204.52 | |
| | Stone Guard's Plate Leggings | 203.08 | -0.7% |

Blade of Hanna (2.0 s, no stats, Improving Form) beats every default-eligible two-hander. The rank 14 Greatsword beating it is by design: rewards above in-game rank 6 are labelled alternatives, never defaults (`pvpRankCap`, owner ruling round 11), which is why every rank 14 weapon reads 0.00 in the json.

## Restored hotfix items (issue #27)

Every file was ranked after `273623f0f` on its simdb (hash checked on the box), and the calibration reproduces the entries below to the hundredth (Enhancement: see above). Values in `items/<class>.json`: Fiery War Axe 870 at 3.6 s (116 to 175, 15 Strength); Glowing Brightwood Staff 812 (3.1 s, 94 to 142); Hand of Edward the Odd 2243 (1.6 s, 33 to 62, 90 spell power); Boots of Avoidance 14549 (6 Strength, 11 Stamina, 18 dodge); Prelacy Cape 7004 (7 healing, 2 spell power, 3 Spirit); Shilly Mitts 9609 (5 spell power and healing, 6 Spirit).

| item | picks |
|---|---|
| Fiery War Axe | Retribution, Enhancement, Arms bare 40, both factions (6) |
| Glowing Brightwood Staff | Restoration druid bare 50 and raid 60; Discipline bare 50, 60 and raid 60; Holy priest bare 50 and raid 60; Restoration shaman bare 50, 60 and raid 60; both factions (20). Discipline bare 60 Alliance moved onto it from 22406 |
| Hand of Edward the Odd | Balance and Shadow bare 60, both factions (4) |
| Boots of Avoidance | Protection paladin and Protection warrior bare 40, both factions (4); every 50 to raid 60 tank entry left it (Agility 14 and 6 dodge gone) |
| Prelacy Cape, Shilly Mitts | no picks: every bare 30 Alliance entry that wore them moved off (Restoration druid and Discipline back, Balance and Destruction hands); they stay alternatives |

## Measure at refresh, MR72 to MR109

Engine 4c595ee6b (the baseline) already carried every fix through MR89; 5603877e4 adds MR90 to MR107's engine side. f17711960 (MR108) and 037798db1 (MR109) are not in this ranking.

| row | in f8b971e14 | verdict |
|---|---|---|
| MR72 Frost 60 build | bare 50 and 60 talents are the guide's (since the baseline); guide test passes | confirmed; stale entry dropped |
| MR73 Fire rows at 25, 27 | leveling walk; no band entry moved | addon-data row |
| MR74 Earthen Silk Slippers | no entry wears 254013 | confirmed |
| MR75, MR76 wand windup | every Shadow and mage bare band holds a wand (in place since the baseline) | confirmed |
| MR77 Judgement as ranged | in place since the baseline; raid 60 834.3 / 840.4 | confirmed |
| MR78 Ret rank rows | addon data | addon-data row |
| MR79 Protection Consecration row | addon data | addon-data row |
| MR80 Ret without ranged-hit items | bare 60 left Infernal Trickster Leggings (Ironfeather Leggings); bare 50 keeps them on a measured margin over the next legs (-1.0% / -0.7% above: Agility and the mana stats, not the hit) | confirmed |
| MR81, MR82 tank rage, Bloodrage | in place since the baseline; tank metrics published | confirmed |
| MR83 Defensive Stance row | settled (engine-70938) | n/a |
| MR84 Protection walk 10 to 18 | leveling walk | addon-data row |
| MR85 Holy paladin auras | in place since the baseline | confirmed |
| MR86 Holy rank rows | addon data | addon-data row |
| MR87 Elemental 60 build | talents are the guide's; guide test passes | confirmed; stale entry dropped |
| MR88 Enhancement 60 build | bare 50 and 60 talents are the guide's; raid 60 waits on MR108 | confirmed; entry kept for MR108 |
| MR89 Chain Lightning on Stormstrike | nothing published | n/a |
| MR90 Deep Wounds, Arms and Fury | Arms bare 60 -3.9 / -3.2%, raid 60 -3.1 / -2.5% (lane -3.3 to -3.7, -2.6 to -2.8); Fury bare 60 -3.4 / -2.5%, raid -0.9 / -1.7% with re-picks | confirmed |
| MR91 Deep Wounds, Protection warrior | threat moves under the tank lanes' other changes (bare 60 923 to 930, raid 1,677 to 1,675) | confirmed in direction |
| MR92 Bane of Havoc's global | settled; nothing published | n/a |
| MR93 Affliction 60 build | talents moved at bare 50 and 60 (+4.4% / +5.6 to +5.9%, lane +5.6% at 60); guide test passes | confirmed; stale entry dropped |
| MR94 tank health | see Checks | confirmed |
| MR95 priests' tank Flash Heal | Discipline raid 60 +3.4 / +4.0%, bare 60 +2.1 / +3.2%; Holy raid +1.1 / +1.3%, bare 60 +0.4 / +1.1% | confirmed |
| MR96 Devilsaur Eye | 6 of 20 hunter 50/60 entries keep it (all six are `kept_alternative` incumbents re-measured on the new engine; its equip stats still beat the field, and the emptied-trinket sims lose without it); Marksmanship raid -0.8%, Beast Mastery -0.7%, Survival -0.8 / -1.2% with the re-picks | confirmed |
| MR97, MR106 Gloves of Holy Might | 51 entries wear 867 on push 114008's trigger (rogues 40 to 60, Arms 40 to raid, Fury 40 to 60, Ret 40 and 50, hunters, Enhancement 40 and 50 Alliance, Feral 40 and 50); Survival 60 and raid and Fury bare 60 Horde moved off it | confirmed |
| MR98 client proc rows | settled | n/a |
| MR99 Thunder Clap on every tank's boss | raid 60 damage taken: bear 351 to 297, paladin 352 to 296, warrior 230 to 233; tank order bear 497.9, paladin 487.0, warrior 425.8 (Alliance) | confirmed |
| MR100, MR103 raid fight | raid 60 Alliance: shaman 135.0, paladin 134.8, druid 133.2, Holy 130.8, Discipline 126.8 (Horde 135.0, 134.8, 133.1, 130.8, 127.1): MR103's expected figures exactly; /tiers/healer switches since all 20 carry it | confirmed |
| MR101 Ignite | Fire raid 60 -1.0%, Arcane raid -0.5% (lane -0.9 to -1.1, -0.4 to -0.5); Fire bare 30 and 40 fall 1.0 to 1.6%: bare 30 wears Snake Hoop 6750 and Runemender's Seal 273031, both cut in the 70938 data (Snake Hoop lost 7 Spirit and 2 Intellect; the casters' bare 30 entries all fall 0.2 to 1.6%), bare 40 Alliance still wears Runemender's Seal | confirmed |
| MR102 ground channel ticks | nothing published | n/a |
| MR104 Rake's hit | cat bare 30 to 60 -0.9 to -2.4%, raid 60 -0.7 / -0.8% | confirmed |
| MR105 | superseded by MR107 | n/a |
| MR107 restored hotfix items | see above; its Enhancement part (Stormstrike, Maelstrom, Mental Dexterity read from the push) is engine 16e5eb6bf, after 5603877e4 | confirmed for items; Enhancement goes with MR108 |
| MR108 Enhancement card and raid build | not in this ranking (f17711960; calibration shows main's card +0.6 to +2.0% bare) | next ranking |
| MR109 mage gems | not in this ranking (037798db1) | next ranking |

Other moves: 20 melee and hunter entries moved onto Icemail Jerkin 1981 (24 Agility, 10 hit in 70938); Band of Mending 22334 (+10 spell power) joined 16 healer entries; Runemender's Seal and Snake Hoop lost Stamina/Spirit and left several 30 to 50 entries.

## Open (not ranker bugs)

- Holy priest bare 40 and 50 Horde reproduce 0.3% under the published figure (3.3 sigma); the next ranking replaces the file.
- The next ranking (on 037798db1) owes MR108, MR109 and Mental Agility on Shadowform; it also clears the Enhancement stale entry. It started landing while this round ran (5060a7eb7 warlocks, 5c468f999 mages, 00:07 to 00:10 UTC 14 Nov); those files are not verified here, and the published-file tests pass on main with them (sanity 28 of 28, guides 27 plus Enhancement skipped).
