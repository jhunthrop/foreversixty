# Sim-accuracy round 9: integration (2026-11-02)

The owner's round-9 rulings, the integration, the pin and the review. The lane reports beside this file have the per-spec rows: warlock.md, casters.md, physical.md, hybrids.md, healers.md, auras.md, items.md, crit.md, ranks.md, ranker.md and platform.md. Crit and ranks are new lanes this round, one for each of the owner's two engine-wide asks.

## Owner rulings, and where each landed

| # | ruling | outcome | where |
|---|---|---|---|
| R1 | Open the "battleground-reputation" and "darkspear-islands" gates at launch | **Done.** Both gates are phase "launch" in data/curated/content-phases.json, citing Blizzard's posts. An offline loot-merge shows the held refresh: 34 sources and 52 quest entries lose "later", and 347 equippable items gain an open source. The ranker found and fixed three side effects the open gates exposed (K8, K9, K13 below). | ranker.md; S f94b40cf2 |
| R2 | Keep the rank-10 cap reading `title_rank`; record the switch to honor cost | **Done.** The comment at `pvpRankExceedsCap` records it: ItemExtendedCost has honor costs but no item link, so the cap can move to honor cost only once the client links them. | ranker.md; S c8ab58049 |
| R3 | Rotations may downrank when mana-bound; curated downranks are respected | **Done.** Below 60 a ranked spell casts the highest learned rank not above the rank the rotation names, so Holy and Discipline's Renew, Discipline's Shield and Holy paladin's Flash of Light and Holy Shock keep their authored ranks. Per-band and per-preset choices live in data/curated/rank-choices.json; /sim, the ranker, spec-breakdown, the ladder and the addon all resolve through it, and /sim and the ranker now send the preset. 43 rows are adopted beyond error at seeds 1 and 50001; five raid-60 rotations change rank. Every adopted rank was checked to deal or heal its own client amount and cost its own mana. Healer rows below 60 are measured on the level-60 heal profile, and the player notes say so. | ranks.md; S 5c63f6719 to 3929ed358, 49356dbcf, dcd25dee3; E 0a9ef0d0c |
| R4 | Demonology's sacrifice then summon is a rotation bug; builds keep their talents | **Done.** Prepull actions honour their condition and Demonic Pact is a readable aura, so the APL sacrifices the Imp only when the build knows Pact, and otherwise summons the Succubus and casts Soul Link. +0.1 to +0.4% at bands 20, 40, 50 and bare 60. | warlock.md; E 17cc0ec01, S ddea9cfcc |
| O1 | actionlint accepts `xcode-27` | **Done.** .github/actionlint.yaml declares the label; actionlint 1.7.12 exited 1 before and exits 0 over all 17 workflows; tools.yml lints every workflow. The held data.yml is untouched (its shellcheck findings are ignored by path). | platform.md; S f02694e28 |
| O2 | `real-engine` gates deploy | **Done.** Deploy needs build, real-engine-scope and real-engine. A failed or cancelled real-engine blocks deploy; one the scope job skipped does not. A vitest evaluates deploy's condition over every result (3 of 7 cases failed before). | platform.md; S 8d9ea7ec5 |
| O3 | A sharpening stone must not raise spell crit; separate melee and spell crit | **Done.** Agility feeds melee and ranged crit, Intellect spell crit, each class keeps Era's two base-crit columns; stones are melee only, wizard oil spell only; item crit (stat 32) and aura-290 sources reach every table. Forever beta logs support two pools (poisons 2.2% against 19% melee, Consecration 6.1% against 12.4%, hunters 6.3% logged against 5.7% predicted). The weights vocabulary keeps one `crit` id (I2). | crit.md; E c14bf02a5; S 67f5e9d9b, 8235fc1b9 |
| O4 | Tank weight sweeps deterministic | **Done, and round 8's evidence withdrawn.** The sweeps were deterministic; round 8's "stamina weighs nothing" and 3-to-15-slot spread came from a probe that ran tanks through the damage sweep. On the tank sweep the same seed gives identical weights in all 30 cases, two seeds differ by 0 to 3 slots, and none fails. A real-engine test and a named seed hold it. | ranker.md; S d47e460f5 |
| O5 | The ranker's run time | **Done.** A sim memo and concurrent weight sets and factions: all 336 rows identical to origin/main, 17% fewer sims, wall -36% and CPU -20% under the same load, race detector clean. | ranker.md; S fd2fa8ecd, cc3c4c627 |

## Confirmed and fixed, by lane

| lane | confirmed | the bugs |
|---|---|---|
| warlock | 1 | Demonology's sacrifice then summon (R4) |
| casters | 4 | Arcane and Frost never cast Evocation; Clearcasting kept by an instant at the proc's moment (shaman, mage); the raid's Curse of Shadow (Forever deleted it); base mana from the planner, not BaseMP, at 8 class levels |
| physical | 10 | Fury in Battle Stance to 36; Subtlety pools for Ghostly Strike; Savage Strikes on Survival's melee; the cat's Bite; ammo and quiver by level; Beast Mastery's Serpent Sting clip; Mortal Shots on Serpent Sting; the execute weapons below 25%; Expose Armor at 450 a point; hunter pets at level-60 stats and damage at every band |
| hybrids | 6 | Omen of Clarity and Maelstrom Weapon kept at the proc's moment; Windfury and Rockbiter attack power by level; Enhancement's own Windfury Totem under Windfury Weapon; a lower Strength of Earth or Grace of Air rank gave the top rank's amount; four talent buffs' display ids |
| healers | 6 | Twin Disciplines on Prayer of Mending's heal; priest racial heals had no class mask; Nature's Swiftness covered Wild Growth; the addon rotation table named ranks the sim never casts; two note fixes |
| auras | 11 | the raid's Winter's Chill handed a mage stacks; Thorns, Retribution Aura and Stoneskin by rank (and Stoneskin growing a rank every iteration); party Mana Tide by rank and level; Smoked Desert Dumplings from level 1; two missing Forever salmon foods; Blasted Lands buffs per spec; Survival stones both weapons; Orc Command reaches pets only |
| items | 5 | Aegis of the Scarlet Commander fired on every block (5 sec cooldown); Dormant Heart of the Mountain did nothing; the ranker's list of engine-modelled items was a stale Oct 7 scrape (fixed by the ranker, K11); terrain-only bonuses counted in the item files (Rune of the Guard Captain read 42 attack power, the sim gives 14; 16 items); the Alliance Darkspear Islands (Theramore) rewards had no source, now client reputation sources with each faction's side |
| crit | 8 | one crit pool; stones and wizard oil on both tables; Cruelty; spell-only set bonuses, Atiesh, Firestone and the set-bonus decoder; simdb dropping spell-only and every-table crit enchants |
| ranks | 3 | the rewrite lifted curated downranks below 60; the addon did the same; a fork test priced Holy Light rank 8 at the book rank's 660 mana (client 580) |
| ranker | 6 | Darkspear vendor stock with no standing or side (K8); two Unique-Equipped family items in one pair (K9); Alliance Sentinel legs on Horde rows in 21 specs (K13, two commits); the bare pass did not name its preset (K14); set completion tried only the cheapest combination (K15, Fury 30 Alliance picked 0.6% worse than origin/main); verify never measured a wand past the runner-up (K16, Shadow 40 Alliance 1.0% worse). The new client reputation sources are gated by content-phases: Powderfuse opens at launch (inferred: the Riverglades is a launch zone), The Watchers is "later", uncertain (K17) |
| platform | 3 | the Discord bot's unranked PvP line; the coach read the 10 s eating aura, not Well Fed; the coach's old foods (plus the panel's six inert flags, P7's preset wiring and the two CI items) |
| integrator | 4 | see below |
| **total** | **67** | |

## Integrator findings

| # | finding | evidence | confirmed? | fix sha |
|---|---|---|---|---|
| I1 | **Four player rotation notes named a review path, its date or a data file.** The rotation sync's leak guard stopped the web build (`npm run sync:rotations`, prebuild) and rotation-apl.test.ts failed. Caught by the platform lane. | web/scripts/rotation-apl.mjs assertNoEngineeringLeak | CONFIRMED | S a2cd0dd35; E 497deb4fc |
| I2 | **The crit split leaked `melee_crit` and `spell_crit` into the weights vocabulary**, which the contract's 10.8 does not list: the page and API would offer weight controls no Forever item carries. `statid.TableOnly` leaves them out. | sim/request vocabulary tests failed on the pin | CONFIRMED | S 8235fc1b9 |
| I3 | **addon-builds.json was stale** against Restoration druid's Heart of the Wild build. | TestCommittedAddonBuildsFollowTheGuides | CONFIRMED | S c89c41fa2 |
| I4 | Frost's expected ability names lacked Evocation. | data tests/test_apl.py | test-only | S 0e2f73f8c |
| I5 | The panel's inert list gains Winter's Chill after the auras lane's engine fix. | buffs.test.ts failed first | follow-on | S b57e15723 |
| I7 | **The coach's food catalogue missed round 9's preset salmon foods**, and its preset test read food labels from the groups alone; logs CI failed on main after the first push. The logs module was not in the integration gates; it is now. | logs CI on 5d2b17b3d: TestCatalogue_FoodsAreThePublishedPresetsFoods | CONFIRMED | S 07cf5c454 |
| I6 | The casters and hybrids lanes each carried the Clearcasting fix (hybrids cherry-picked ff4bae666); merged as one copy. Golden conflicts in TestP1Mage, TestFeralTank, TestProtection, TestP1DPSWarrior and TestP1TankWarrior were regenerated together; debuffs.go keeps both the Curse of Shadow and the Winter's Chill removals. | int/r9 merges | merge | E f42d92c50, 282a87fed |

## Integration

- **Fork `int/r9`** merges lane/r9-crit, -casters, -hybrids, -healers, -items, -warlock, -auras, -physical and -ranks. The ranker and platform lanes made no engine change. The suite passes with `-tags=with_db` except `sim/web`, which also fails setup at the old pin.
- **The pin is 4d3129cae** (fork master fast-forwarded from 40021374a). data/proto/ENGINE_SHA follows it; common.proto gains the crit stats (crit lane) and the two salmon foods. The rotation ladder goldens are regenerated at it with `make simdb`.
- **The site** merges wt/r9-crit, -casters, -healers, -items, -warlock, -platform, -auras, -hybrids, -physical, -ranks and -ranker.
- **Gates, all exit 0:** go vet and go test ./... in sim; go test ./... in logs (added after I7); api's coach and sims packages; `make apl-check` against the pin; `make styles-check`; tools' test_wasm_gzip.sh; the full data pytest; web sync, the full vitest, astro check, lint and format:check; a real-data build.
- **One held-output exception in the data gate:** `test_simbuffs_names_every_id_the_engine_lets_a_request_send` is deselected. It holds data.yml's simbuffs.json to IDS.md, which now names the two salmon foods; simbuffs.json is data.yml output, held by the owner, so it is not regenerated here. Only the disabled data and engine workflows run that test. Until the refresh, /sim names the two foods with its humanised fallback.

## Before/after

Pin 40021374a with origin/main, against pin 4d3129cae with this branch. Both run on today's published data: the BiS sets and builds, and the simdb built from the committed inputs. Raid preset, band 60, 10,000 iterations. Healers are in effective HPS on the heal profile, tanks in threat per second.

| Spec | Alliance, seed 1 | Alliance % (1 / 50001) | Horde, seed 1 | Horde % (1 / 50001) |
|---|---|---|---|---|
| druid-balance | 578.1 to 561.4 | -2.9 / -2.9 | 579.5 to 565.2 | -2.5 / -2.5 |
| druid-feral | 613.8 to 597.0 | -2.7 / -2.7 | 608.9 to 593.0 | -2.6 / -2.7 |
| druid-feral-bear | 851.9 to 831.6 | -2.4 / -2.4 | 856.3 to 836.3 | -2.3 / -2.3 |
| druid-restoration | 779.6 to 780.3 | +0.1 / +0.2 | 777.3 to 779.2 | +0.2 / +0.3 |
| hunter-beast-mastery | 658.1 to 629.4 | -4.4 / -4.4 | 661.8 to 633.4 | -4.3 / -4.3 |
| hunter-marksmanship | 649.2 to 618.5 | -4.7 / -4.7 | 653.1 to 623.0 | -4.6 / -4.6 |
| hunter-survival | 756.3 to 731.4 | -3.3 / -3.3 | 759.5 to 734.9 | -3.2 / -3.3 |
| mage-arcane | 597.6 to 594.3 | -0.6 / -0.5 | 601.4 to 600.2 | -0.2 / -0.3 |
| mage-fire | 663.0 to 652.5 | -1.6 / -1.7 | 670.8 to 660.8 | -1.5 / -1.6 |
| mage-frost | 594.7 to 587.1 | -1.3 / -1.3 | 601.5 to 594.1 | -1.2 / -1.2 |
| paladin-holy | 662.7 to 653.6 | -1.4 / -1.4 | 648.2 to 641.4 | -1.0 / -1.1 |
| paladin-protection | 1001.2 to 976.9 | -2.4 / -2.5 | 1018.2 to 993.8 | -2.4 / -2.4 |
| paladin-retribution | 661.0 to 623.3 | -5.7 / -5.7 | 658.8 to 621.3 | -5.7 / -5.7 |
| priest-discipline | 619.6 to 622.4 | +0.5 / +0.5 | 620.0 to 624.9 | +0.8 / +0.7 |
| priest-holy | 687.0 to 705.5 | +2.7 / +2.6 | 674.1 to 693.4 | +2.9 / +2.8 |
| priest-shadow | 717.5 to 703.8 | -1.9 / -1.9 | 724.2 to 712.2 | -1.7 / -1.6 |
| rogue-assassination | 611.7 to 616.3 | +0.8 / +0.7 | 609.3 to 613.6 | +0.7 / +0.7 |
| rogue-combat | 605.8 to 611.5 | +0.9 / +1.0 | 606.0 to 611.8 | +1.0 / +1.0 |
| rogue-subtlety | 532.6 to 538.5 | +1.1 / +1.1 | 532.6 to 538.3 | +1.1 / +1.1 |
| shaman-elemental | 468.2 to 459.6 | -1.8 / -1.8 | 476.8 to 466.9 | -2.1 / -2.1 |
| shaman-enhancement | 594.7 to 575.5 | -3.2 / -3.4 | 594.0 to 575.8 | -3.1 / -3.1 |
| shaman-restoration | 841.5 to 839.9 | -0.2 / -0.2 | 838.9 to 838.0 | -0.1 / -0.1 |
| warlock-affliction | 583.7 to 571.5 | -2.1 / -2.1 | 585.2 to 573.2 | -2.1 / -2.0 |
| warlock-demonology | 649.2 to 637.8 | -1.8 / -1.8 | 654.9 to 642.6 | -1.9 / -1.9 |
| warlock-destruction | 647.0 to 636.6 | -1.6 / -1.7 | 649.7 to 638.9 | -1.7 / -1.6 |
| warrior-arms | 708.9 to 708.9 | +0.0 / +0.0 | 720.8 to 720.8 | +0.0 / +0.0 |
| warrior-fury | 754.4 to 753.9 | -0.1 / -0.1 | 757.8 to 757.8 | +0.0 / +0.0 |
| warrior-protection | 1049.0 to 1062.9 | +1.3 / +1.4 | 1036.9 to 1051.4 | +1.4 / +1.5 |

Where the moves come from:
- **The crit split** lowers everyone whose Intellect gave melee crit or whose Agility gave spell crit: Retribution, hunters, Enhancement, the paladins, bear and cat, then every caster. Warriors are flat. The lane's own table, with no other fix, is in crit.md.
- **Hunters** win back part of their crit loss from Mortal Shots on Serpent Sting, Savage Strikes and Beast Mastery's Serpent Sting refresh.
- **Enhancement** nets its crit loss against the Windfury Totem fix, which is +4.0% alone.
- **Elemental** loses the doubled Clearcasting, gains Lightning Bolt rank 9 (+0.6%), and takes the crit split.
- **Holy priest** +2.7 / +2.9: the tank's Flash Heal at rank 2 and Twin Disciplines on Prayer of Mending.
- **Rogues and Protection warrior** rise on the per-spec Blasted Lands buffs and foods.

The bare bands move far more. The lane reports have each table; the largest:
- Holy paladin bare 30 to 60, +20 to +58%; Elemental bare 30 to 50, +10 to +46%; Restoration shaman bare 20 to 50, +9 to +20%: the rank choices.
- Frost bare +20 to +28% and Arcane bare +4 to +22%: Evocation.
- Hunters bare 20 to 50, -2 to -28%: pets at their level, plus ammo by level.
- Enhancement bare 50 and 60, +3.7 and +4.4%: the Windfury Totem fix.

What this table does not hold. Each needs data.yml's regeneration and the bis.yml re-rank, both held by the owner until a zero round:
- **simdb.bin** predates round 8's item fixes (2,212 items change: stat 45 healing, bonus armor, held off-hand typing, paladin axes) and round 9's crit enchant mapping. The refresh must rebuild it before re-ranking: on the committed simdb a held off-hand replaces the main hand, and Feral and bear sets read 38 to 90% high at band 30 (ranker K10).
- **The re-rank**: the open battleground and Darkspear sources, the client reputation sources, the ranker's fixes, the engine-modelled items, Dormant Heart, the terrain-rider stats, the bare-60 rank choices. On the local re-rank (engine lane/r9-items), 274 of 336 rows rise by more than 0.5% against origin/main and only the five tank rows fall, on the tank's own damage; on the tank score they rank on, all five rise 3.9 to 4.9%. Shadow 40 and Fury 30 Alliance first fell 1.0 and 0.6%; both were real, both are fixed (K15, K16).
- **Item files** predate round 8's Wowhead fill: Glimmering Staff's +34 spell power and healing, and 1,752 items' stats in all (mostly healing), plus 901 class placements (368 paladin axes).
- **simbuffs.json** with the two salmon foods; **Data.lua** and addon-data.json with the corrected rotation ranks.

## wow-player review

Verdicts on the lane reports, the before/after table and the shots:
- **SHIP:** crit, warlock, casters, healers, items and platform.
- **SHIP WITH FIXES:** ranks, physical, hybrids and auras. Each fix, and how it was settled:
  - **Sub-60 healer rows are tuned on the level-60 heal profile.** The healer rotations' notes now say each band's rank is chosen for its mana over a fight sized for level 60 (S b696998ac). Whether the profile should scale below 60 is an owner decision (below).
  - **Elemental's lower Lightning Bolt.** The note says it is a choice for a mana-bound fight, and a shaman not short of mana casts the top rank (S b696998ac).
  - **Holy paladin "+60%" against "+10%".** Both are right: +58.5% at bare 50 is the whole change (the new default plus the band's choice); +10.1% is the band's Holy Light rank 3 over the new default.
  - **Enhancement's Windfury Totem.** The note now says that in a group with no other shaman, Windfury Totem goes down in place of Grace of Air (S b696998ac).
  - **Hunter pets rest on Era's tables** (mangos, vmangos), not a Forever client table; physical.md says so. Warlock demon sheets differ from Era's table and stay open; nothing claims they are verified.
  - **The panel's debuff names** ("Curse of elements", "Hunters mark") are the fixture build's humanised fallback. The committed simbuffs.json spells them as the client does ("Curse of the Elements", "Hunter's Mark"), so the live page is right. Icons on the custom panel are a page follow-up, not a sim bug.
  - **Forever foods.** Baked Salmon and Poached Sunscale Salmon are Forever client rows; the preset list should mark them as Forever's when the refresh names them.
- **Numbers players will question first:**
  - Retribution -5.7%
  - hunters -3 to -5%, and the pets at low bands
  - Enhancement -3%
  - Elemental's lower Lightning Bolt
  - Holy paladin's Holy Light rank 3 at 40 to 59
  - Holy priest above Discipline at raid 60

  None is a bug. The release note for the refresh should say why: Intellect no longer gives a hunter or paladin melee crit, and Agility no longer gives spell crit.

## Still open

- **Owner decisions:**
  - Should the heal profile scale below 60? Every sub-60 healer rank choice is measured on the level-60 fight.
  - Consecration rank 1 for Retribution at 30 and 40 is adopted because the rotation rations Consecration behind a mana line, though Retribution never runs dry there.
  - wasm brotli or no-inlining; the time-bound Survival line after launch (carried from round 8).
- **Held under the adoption rule:** Marksmanship +0.8 to +0.9%, Affliction's Improved Bane of Agony +0.7%, Destruction's Immolate below Conflagrate +0.1 to +0.2%, Fire's Improved Scorch +0.1 to +0.6%, Discipline's Inspiration.
- **Blocked on logs** (launch is 2026-11-04):
  - Holy Shield's damage critting (aura 43, a damage shield; about 2 to 3% of Protection threat)
  - Faerie Fire's form threat (needs an in-game threat reading)
  - armor-debuff stacking under aura 674
  - the one-a-minute rate for unstated "chance on" procs
  - Holy Power's crit-heal size, Swiftmend, Oath of Light, Gift of Nature on Tranquility
  - Era's base-crit columns (Forever's crit game tables 404)
  - the Succubus's Lash of Pain crit after the split (2.7% simmed against 7.1% logged at band 30; under 0.1% of DPS)
- **Single source or open:**
  - warlock demon sheets against Era's pet table
  - which hands an aura-52 weapon enchant reaches; a Survival stone reaching the pet
  - Concentration Aura, forever-bugs #419
  - Forever's new elixirs and 30-second potions (no recipe source or phase)
  - Lacerating Strikes, the normalized-rage off hand, Unbridled Wrath on Heroic Strike
  - Toughness on bonus armor; 1,351 unpicked candidates on Era stats; Mindtap Talisman and Force of Will unmodelled
- **Page and pipeline:**
  - simbuffs.json's two foods (held output; see Integration)
  - `tools/coach/test_persona_starter.py` has 7 failures that also fail on origin/main, and no workflow runs it
  - "Apply to the page" from the request drawer drops the preset
  - icons on the custom /sim panel
  - `allow_shrink` stays while data.yml passes it to `loot`
  - after the data run, empty the active-build banned test's AWAITING_THE_NEXT_DATA_RUN pass for The Watchers (2819)
  - loot's Warsong Gulch faction ids are swapped against the client (client 889 is Warsong Outriders, 890 Silverwing Sentinels). Every item sits under the right side today, but a client reputation source naming 889 or 890 would be sided wrong
  - is The Watchers launch content? It is gated "later" until a source names its zone
