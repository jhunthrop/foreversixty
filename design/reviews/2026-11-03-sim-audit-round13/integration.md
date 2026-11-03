# Sim-accuracy round 13: integration (2026-11-03)

Round 13 on client 1.60.1.70753, built on round 12's unlanded branches (site int/r12 93c5d2510, fork int/r12 076a5a517). The lane reports sit beside this file:
- core, regmelee and regcast (the regression lanes over round 12);
- warlock, casters, physical, hybrids and healers (the class re-audits);
- auras, ranks, trainables, guides and platform;
- hotfix, a lane added mid-round for the 2026-11-03 server hotfix push;
- walk, a lane added mid-round for the leveling talent walk's gaps;
- review, which applied the wow-player reviewers' fixes.

**Not landed.** Following the batch-landing practice, nothing is pushed or pinned. Site `int/r13` and fork `int/r13` wait for the integrator, who lands rounds 12 and 13 together or one after the other (see "For the integrator").

## Owner rulings this round, and where they landed

| # | ruling | outcome | where |
|---|---|---|---|
| 1 | Elemental Tidal Focus at bare 50/60: adopt | Done. The leveling build takes Tidal Focus for Improved Healing Wave from 50. Improved Reincarnation stays and nothing cast is lost. Bare 50/60 +5.7 to +6.1%. | casters.md, S on wt/r13-casters |
| 2 | Improved Bane of Agony: adopt (Shadow-build precedent) | Done on the Affliction raid build (Improved Drains, which the raid rotation never uses, gives the point). Raid +0.72% on the guide build. The published number moves at the refresh, since spec-breakdown runs the published BiS talents. Every tier gate checked. | warlock.md S 66fa0d042 |
| 3 | Rank rows keyed by race | Done end to end: sim, spec-breakdown `-race`, rotation-search, the ladder, the coach and the addon export (`raceSpellId`, read through `Compat.playerRaceSlug`). The faction key is deleted. A row names races only where one loses beyond error; an unnamed race casts its highest learned rank. | ranks.md S 33d1a01e4 |
| 4 | Randomize the regen tick's phase per fight | Done for mana: the first tick is drawn per iteration from the seeded RNG. Energy has no tick in Forever (it regenerates continuously). Every spec moves within ±0.3%. Fork goldens regenerated. | core.md E e88a33c97 |
| 5 | Arcane Shot spell power (E 05b72a3b5) | Kept, **pending the owner's in-game tooltip check.** regmelee re-checked that the removal is complete and consistent (Arcane Shot, Serpent Sting, Volley). | regmelee.md |
| 6 | Level-60 heal split | Mid-round the owner's Warcraft Logs key arrived. The healers lane wrote a fetcher (vanilla.warcraftlogs.com: MC, Onyxia, BWL as a provisional proxy), but gcloud's login expired after 12 cached responses. 60 stays 0.659 under its stated rule. Finishing it needs `gcloud auth login`, then healers/raidsplit.py, then a heal-profile change: the profile refuses a logged split at 60 today. | healers.md |

## The hotfix push (new data, mid-round)

A server hotfix push for 1.60.1.70753 went live on wago on 2026-11-03 at 14:14 UTC. The committed hotfix-push.json, fetched 11-02 22:28, has none of it; the trainables lane found it. It buffs hunters:

| spell | before | after |
|---|---|---|
| Aimed Shot (top trainer rank) | 166 | 350 |
| Arcane Shot | 217 | 280 |
| Serpent Sting (rank 8, per tick) | 83 | 100 |
| Sniper Shot | 295 | 590 |
| Aspect of the Hawk (rank 6) | 55 | 110 |
| Trueshot Aura | 50 | 100 (ranged attack power only, as the client's single aura-124 effect) |

Lacerating Strikes' new attribute bit is an aura-bounce flag the engine never reads. The other rows are Super Hogger event NPCs and item effects on no item. The new "Endless" elixirs are on no item yet, so they stay open. **The push is the only source.** Wowhead's Forever tooltips still show the old numbers (checked 2026-11-03); re-check once Wowhead re-scrapes.

## Confirmed and fixed, by lane, with origin

Origin key: **R12** caused by a round-12 fix or merge; **landing** caused by the launch-day landing c6bdc855c and missed by round 12; **R9-R11** caused by those rounds; **old** predates round 9 and was not found before; **new-data** data changed under us (the 70707/70753 pushes). Counted once, where fixed: Frostfire Bolt's speed under casters (found by regcast); the hotfix push under hotfix (found by trainables); the Lacerate landing line under walk.

| lane | confirmed | R12 | landing | R9-R11 | old | new-data | bugs |
|---|---|---|---|---|---|---|---|
| guides | 25 | 2 | 10 | | 13 | | prose and landing facts the client contradicts: feral Claw costs, the warrior index's Bloodrage, paladin landing talents, Holy stat order, Twilight Focus, shaman hit/Concussion/stat order, Ghost Wolf's 3 s cast, Drain Hope, the Incubus talents, Holy Strike, Restless Blades, Cure Poison, Survival's aspect, the mage gem/rune cooldown, Cone of Cold at 26, Heating Up, Frostfire Bolt talents |
| walk | 10 | | | | 9 | 1 | leveling bands missing a cast talent that fits: Arms 40 Mortal Strike (+20%), Survival 30 Strider Kick (+10%), Assassination 30 Mutilate (+8%), Subtlety 20/30 Ghostly Strike and Hemorrhage, Fury 30 Death Wish, Demonology 30 Soul Link, Fire 20 Pyroblast, Restoration shaman 30 Mana Tide, Affliction 30 Siphon Life; the druid landing's Lacerate line (new-data) |
| hotfix | 6 | | | | 2 | 4 | the push's hunter values and Trueshot Aura; BM and MM Aimed Shot on the shared cooldown; MM keeps Serpent Sting (old); MM raid build with Lone Wolf and no pet (old; the guide's "the cat does more" came from the landing's 1feda0686 but the build predates it) |
| casters | 5 | | | | 5 | | Elemental waits for Lava Burst (haste tie); Searing Totem's 2.437 s attack (beta logs); Balance's Nature's Grace on landing (beta logs); Balance's pre-pull Wrath; Frostfire Bolt 24 yd/s |
| ranks | 5 | 1 | | | 4 | | Lightwell rank 1 at bare 60 now loses (removed); Shadow Mind Blast rank 3 at 40 (Dwarf, Night Elf), Devouring Plague rank 2 at 50 (Night Elf), Shadow Word: Death top rank at 50 (Human); Protection paladin Holy Strike rank 1 at 40 (+0.6 to +0.9% threat) |
| healers | 4 | | | 1 | 3 | | Restoration druid tank Rejuvenation rank 10 in raid (R9 dcd25dee3); party Rejuvenation rank 1, not mana-paced; Holy's Prayer of Mending whenever ready; Shadow keeps Vampiric Embrace up |
| platform | 3 | 2 | | | 1 | | the mobile planners kept round 12's old tier-first order (their golden vectors would have turned web CI red); three "lowest tier first" notes; spec-breakdown and the sims API ran talent strings the game refuses (now ErrIllegalBuild / 400 illegal_talents) |
| trainables | 2 | | | | 2 | | auto-learned rank 1s, stances and Auto Shot were skipped; seven abilities counted their triggered spell as a rank (zeroed costs such as Lightning Shield's) |
| hybrids | 2 | | | | 1 | 1 | bear bands 30/40 without Lacerate and Primal Bite (threat +41% / +19%); Templar's Bulwark's stale level-30 gate (70707 push 113139) |
| physical | 2 | | | | 2 | | Serrated Blades and Hack and Slash's mace clause ignore a share of armor, not a flat amount (Subtlety raid -1.1%, bare +0.5 to +1.5%); Demoralizing Shout's threat grows 0.8 a level |
| auras | 2 | 1 | | | 1 | | the int/r12 regmelee merge (8850cbc26) undid four preset reasons; the Skyborne racial Read Leyline (1259705) was not modelled |
| regmelee | 1 | | | | 1 | | the Fury build's Improved Intercept point, which raid Fury never uses, moved to Raging Blows (+0.4 to +0.6% on the guide build) |
| core, warlock, regcast | 0 | | | | | | aligned (regcast's one find is counted under casters) |
| **total** | **67** | **6** | **10** | **1** | **44** | **6** | |

Bugs by origin, as the owner asked:
- **Caused by round 12: 6.** Two guide fixes that contradicted other pages, round 12's Lightwell row, two mobile-planner leftovers of round 12's strip-order change, and the int/r12 merge that dropped four preset reasons.
- **Caused by the landing: 10**, all guide prose round 12 missed.
- **Newly found: 51.** That is 44 older, 6 from new data (5 of them today's hotfix push and its effects) and 1 from round 9.

Of the 67, 26 are guide or landing copy and 41 are sim, data or tool bugs. No round-12 *sim* fix caused a sim bug: the regression lanes re-checked every round-12 change (Poised, the Arms opener, Thunder Clap 20%, Windfury ending with the totem, missile speeds, Elixir of Holy Power, Juju Might, Enrage armor, Primal Bite, the heal split, Vampiric Embrace 20%, Blood Fury on heals, Orc Command, Evocation, the cast-timing fix, the faction rows, trainables) and found them aligned.

Not counted: the review lane's wording fixes; my Restoration druid guide line (S 9f2743f27), which the healers' fix had made wrong; the healer test moved to the race key at merge; the five fork goldens that Trueshot Aura's push moves (E 4f81a2a92); the Demonology build proposal the warlock lane withdrew as illegal.

**Is the count falling? No: 69 in round 12, 67 in round 13.** But the mix changed:
- 16 came from two lanes added mid-round (hotfix 6, walk 10), on ground no earlier round covered.
- Guide copy is 26 of 67 (16 of 69 last round).
- Bugs caused by round 12's own sim fixes: 0. Earlier rounds' fixes caused 1, down from 4.

## Merge notes

| # | conflict | resolution |
|---|---|---|
| M1 | fork TestElemental and TestEnhancement goldens (core's regen phase, casters' Searing Totem and Lava Burst wait) | regenerated once on the merged engine (E c555683de) |
| M2 | data/curated/classes.json druid line (walk's Lacerate, guides' Cure Poison removal) | both kept |
| M3 | sim/request/caster_audit_rotation_test.go (casters' Balance pre-pull test, healers' Vampiric Embrace test) | both kept |
| M4 | healer_ranks_test.go used the faction key the ranks lane deleted | moved to the race key (night-elf, tauren) |
| M5 | spec-breakdown main.go/heal.go/tank.go (ranks' `-race`, platform's legality check) | auto-merged; both kept |
| M6 | five fork CharacterStats goldens (druid tank, mage, Protection paladin, both warriors) failed after the hotfix merge | Trueshot Aura 100 in the test party buffs: ranged attack power +50, nothing else; regenerated (E 4f81a2a92) |

## Before/after

Before: round 12's heads (site 93c5d2510, fork 076a5a517). After: site int/r13 with fork int/r13 (before the review lane's Arms Charge change, which the review report measures). Raid preset, band 60, 10,000 iterations, both factions; healers in effective HPS on the heal profile, tanks in threat per second. Druid and hunter rotation copies were synced to the curated files for both sides (the integrator's apl-sync).

| Spec | Faction | seed 1 | % (seed 1 / 50001) |
|---|---|---|---|
| druid-balance | alliance | 566.9 to 569.5 | +0.46 / +0.53 |
| druid-balance | horde | 572.2 to 574.9 | +0.47 / +0.44 |
| druid-feral | alliance | 665.0 to 664.9 | -0.02 / -0.11 |
| druid-feral | horde | 660.0 to 660.4 | +0.06 / -0.05 |
| druid-feral-bear | alliance | 1020.1 to 1020.2 | +0.01 / -0.01 |
| druid-feral-bear | horde | 1010.7 to 1011.2 | +0.05 / -0.01 |
| druid-restoration | alliance | 834.6 to 934.2 | +11.93 / +12.05 |
| druid-restoration | horde | 833.7 to 933.7 | +11.99 / +11.97 |
| hunter-beast-mastery | alliance | 619.4 to 659.1 | +6.41 / +6.24 |
| hunter-beast-mastery | horde | 623.5 to 663.5 | +6.42 / +6.33 |
| hunter-marksmanship | alliance | 617.3 to 697.2 | +12.94 / +12.94 |
| hunter-marksmanship | horde | 622.5 to 701.3 | +12.66 / +12.74 |
| hunter-survival | alliance | 741.1 to 741.0 | -0.01 / -0.07 |
| hunter-survival | horde | 744.7 to 744.0 | -0.09 / -0.08 |
| mage-arcane | alliance | 613.7 to 614.2 | +0.08 / -0.08 |
| mage-arcane | horde | 619.8 to 619.2 | -0.10 / -0.03 |
| mage-fire | alliance | 665.2 to 666.4 | +0.18 / +0.21 |
| mage-fire | horde | 673.0 to 673.0 | +0.00 / +0.10 |
| mage-frost | alliance | 627.3 to 628.0 | +0.11 / -0.18 |
| mage-frost | horde | 632.0 to 632.3 | +0.05 / -0.11 |
| paladin-holy | alliance | 656.5 to 656.8 | +0.05 / +0.03 |
| paladin-holy | horde | 645.1 to 644.6 | -0.08 / -0.06 |
| paladin-protection | alliance | 1113.1 to 1113.8 | +0.06 / -0.01 |
| paladin-protection | horde | 1125.7 to 1125.5 | -0.02 / +0.07 |
| paladin-retribution | alliance | 704.3 to 704.2 | -0.01 / +0.01 |
| paladin-retribution | horde | 701.5 to 701.0 | -0.07 / -0.11 |
| priest-discipline | alliance | 652.4 to 652.1 | -0.05 / -0.11 |
| priest-discipline | horde | 655.9 to 655.7 | -0.03 / -0.03 |
| priest-holy | alliance | 710.5 to 720.8 | +1.45 / +1.62 |
| priest-holy | horde | 698.9 to 708.7 | +1.40 / +1.43 |
| priest-shadow | alliance | 712.7 to 684.2 | -4.00 / -3.93 |
| priest-shadow | horde | 720.0 to 691.8 | -3.92 / -3.94 |
| rogue-assassination | alliance | 626.4 to 626.4 | +0.00 / +0.00 |
| rogue-assassination | horde | 622.1 to 622.1 | +0.00 / +0.00 |
| rogue-combat | alliance | 627.0 to 627.0 | +0.00 / +0.00 |
| rogue-combat | horde | 628.2 to 628.2 | +0.00 / +0.00 |
| rogue-subtlety | alliance | 550.8 to 544.7 | -1.11 / -1.13 |
| rogue-subtlety | horde | 550.2 to 544.0 | -1.13 / -1.11 |
| shaman-elemental | alliance | 471.6 to 472.4 | +0.17 / +0.19 |
| shaman-elemental | horde | 479.6 to 480.3 | +0.15 / +0.08 |
| shaman-enhancement | alliance | 585.4 to 586.1 | +0.12 / +0.00 |
| shaman-enhancement | horde | 586.0 to 587.0 | +0.17 / +0.02 |
| shaman-restoration | alliance | 842.8 to 842.3 | -0.06 / +0.04 |
| shaman-restoration | horde | 840.2 to 840.0 | -0.02 / +0.01 |
| warlock-affliction | alliance | 570.0 to 569.9 | -0.02 / +0.07 |
| warlock-affliction | horde | 571.9 to 572.0 | +0.02 / +0.02 |
| warlock-demonology | alliance | 635.5 to 635.4 | -0.02 / +0.06 |
| warlock-demonology | horde | 639.3 to 639.3 | +0.00 / +0.02 |
| warlock-destruction | alliance | 634.0 to 633.8 | -0.03 / +0.03 |
| warlock-destruction | horde | 636.2 to 636.8 | +0.09 / +0.09 |
| warrior-arms | alliance | 719.0 to 719.0 | +0.00 / +0.00 |
| warrior-arms | horde | 731.2 to 731.2 | +0.00 / +0.00 |
| warrior-fury | alliance | 755.3 to 755.3 | +0.00 / +0.00 |
| warrior-fury | horde | 758.2 to 758.2 | +0.00 / +0.00 |
| warrior-protection | alliance | 1081.2 to 1081.2 | +0.00 / +0.00 |
| warrior-protection | horde | 1068.8 to 1068.8 | +0.00 / +0.00 |

Where the moves over 1% at raid 60 come from:
- **Restoration druid +12.0%:** the tank's Rejuvenation at rank 10 and the party's at rank 1, no longer mana-paced. The druid now covers about 95% of incoming damage, above the retune rule's ~85% ceiling: the 1.85 scale wants an owner retune at the refresh.
- **Marksmanship +12.7 to +12.9%, Beast Mastery +6.2 to +6.4%:** the hotfix push, Aimed Shot on the shared cooldown, and Serpent Sting back for Marksmanship. Marksmanship's new Lone Wolf raid build adds about +5% more at the refresh (737.7 at seed 1, horde, on the guide build).
- **Holy priest +1.4 to +1.6%:** Prayer of Mending whenever ready.
- **Shadow priest -3.9 to -4.0%:** Vampiric Embrace kept up; its group healing is not counted in DPS.
- **Subtlety -1.1%:** Serrated Blades ignores 9% of the boss's armor, not a flat 126.

The review lane's Arms opener (Charge and the pre-pull shout at the pull, not 2 s before it) costs Arms about -0.2% at raid 60 and bare 60, both factions and seeds; it is not in the table above.

Under 1% at raid 60: Balance +0.4 to +0.5% (pre-pull Wrath, Nature's Grace on landing); Fire, Elemental and Enhancement within +0.2% (regen phase, Searing Totem); every other spec within ±0.2%. Warriors and rogues other than Subtlety are identical to the decimal. Fury's Raging Blows (+0.4 to +0.6%) and Affliction's Improved Bane of Agony (+0.72%) are on the guide builds and reach the published numbers at the refresh.

Moves over 1% below 60 (lane reports have every band):
- Restoration druid bare 60 +40%, bands 20 to 50 +2 to +14%; Holy priest bare 60 +22%, bare 40/50 +10 to +11%.
- Bear threat: band 20 +4.8%, band 30 +41%, band 40 +19%.
- Arms 40 +20%, Survival 30 +10%, Assassination 30 +8%, Subtlety 30 +6.7% and 20 +2.6%, Fury 30 +3.8%, Demonology 30 +2.3%, Fire 20 +2%, Restoration shaman 30 +1.9%, Affliction 30 +1.3%.
- Elemental bare 50/60 +5.7 to +6.1% (Tidal Focus); Balance bare 30 to 60 +1.1 to +5.2% (leveling build), bare 60 about +5.8% overall.
- Hunters bare bands +2.4 to +12.6%; Subtlety bare 40 to 60 +1.0 to +1.5%; Shadow bare 40 to 60 -3.1 to -3.9%.

## wow-player review

Two reviewers (physical/hybrid/hunter; casters/healers/guides/auras/trainables/platform).
- **SHIP:** physical, ranks, core, regcast, platform.
- **SHIP after fixes:** the review lane made every one (review.md):
  - hotfix: Beast Mastery's Aimed Shot advice softened (its gain is within error), "right after an auto shot fires", dismiss the pet before the pull; regenerate addon-data so the priority card shows Aimed Shot and Serpent Sting, and re-take the hunter captures hydrated;
  - regmelee: raid Arms no longer Charges the boss before the tank;
  - casters: Elemental's leveling prose says Tidal Focus from 50;
  - healers: the Restoration druid's Ranks paragraph; the Shadow guide says the sim keeps Vampiric Embrace up and counts only damage;
  - warlock: Affliction's raid-talents sentence split;
  - guides: "in Forever" on the shared gem/rune cooldown; Enhancement's stat order;
  - auras: Read Leyline is not cast automatically; races.json must be regenerated before release;
  - trainables: land the hotfix push before launch (regen step 1).
- **Owner calls the reviewers made:**
  - Protection paladin band 20: the reviewer would take Templar's Bulwark over the second Sacred Duty point (chance of death 77% to 66%, threat -2%). Not adopted: it fails the threat metric the adoption rule uses. **Owner decision.**
  - Balance leveling build without Improved Entangling Roots: the reviewer agrees (a leveling Moonkin does not kite with Roots).
  - Marksmanship raid with Lone Wolf: agreed. Beast Mastery's raid build can reach Lone Wolf (20 Marksmanship points; it needs 10) and skips it because its damage is the pet's; not measured this round.
- **Release-note lines** (the reviewers'):
  - "The 3 November hotfix raised Aimed Shot, Arcane Shot, Serpent Sting, Sniper Shot, Aspect of the Hawk and Trueshot Aura; hunter rotations are updated."
  - "Marksmanship raids with no pet: Lone Wolf. Keep the pet while leveling. Marksmanship keeps Serpent Sting up."
  - "Restoration druids cast Rejuvenation rank 1 on the party and the top rank on the raid tank. Holy priests cast Prayer of Mending every time it is ready."
  - "Shadow priests keep Vampiric Embrace up; published DPS counts only damage."
  - "Elemental shamans at 50 and 60 take Tidal Focus; Affliction raid builds take Improved Bane of Agony; Fury takes Raging Blows."
  - "Level 20 to 40 builds now take Mortal Strike (Arms 40), Death Wish (Fury 30), Strider Kick (Survival 30), Mutilate (Assassination 30), Hemorrhage (Subtlety 30) and Pyroblast (Fire 20). Bears take Lacerate at 25 and Primal Bite at 40."
  - "Subtlety's armor ignore is a share of the target's armor."
  - "Rank choices now follow your race."
  - "The Skyborne racial is Read Leyline; the sim does not cast it on its own."
  - "Guides: Lacerate is a Feral talent at 25, mana gems and Demonic Runes share a cooldown in Forever, and more fixes."

## For the integrator

**Branches.**
- Site `int/r13`: all 16 site lane branches plus the coordinator commits. Head: see the last line of this section.
- Fork `int/r13`: 9 lane branches plus the goldens commits. Head: see the last line of this section.
- Both sit on round 12's unlanded heads. Land round 12's list first (its integration.md), then this one; or merge both and run one regen. No fork master fast-forward.

**Pin first.** The site pins the fork's int/r13 head, as at every pin:
- data/proto/ENGINE_SHA, sim/enginever/version.go;
- web/src/lib/sim/version.ts and the web sim fixtures;
- round 12's proto items (common.proto holy_power_buff = 38, common_pb2.py).

**Regenerate, in this order.**
1. **layer-hotfix-push + spell generators** (the hotfix lane's list):
   - `python -m pipeline fetch --product wowf --build 1.60.1.70753` on a fresh checkout;
   - `python -m pipeline layer-hotfix-push --build 1.60.1.70753`: the push of 2026-11-03 14:14 UTC lands in data/builds/1.60.1.70753/hotfix-push.json;
   - `simconst` (only spellconst/hunter.json moves), then `levels` (spellranks came out unchanged), then `trainables --build 1.60.1.70753` (it also carries the trainables lane's fixes), then `make simdb`.
   - Until then sim/leveling TestTrueshotAuraRanksMatchTheClient fails.
2. **apl-sync** (`make apl-sync`, then `make apl-check`):
   - druid-feral-bear (round 12's Primal Bite) and druid-feral;
   - hunter-beast-mastery, hunter-marksmanship;
   - druid-balance, shaman-elemental, druid-restoration, priest-holy, priest-shadow;
   - warrior-arms (review).
3. **The committed build copy of curated data:** data/builds/1.60.1.70753/classes.json and races.json (round 12's paladin line plus this round's Lacerate, landing and Read Leyline text). test_curated's committed-build check fails until then.
4. **addon-builds.json** (`go run ./sim/cmd/addon-builds -repo-root .`):
   - Fury (Raging Blows), Affliction (Improved Bane of Agony), Elemental (Tidal Focus), Balance (leveling build), Marksmanship (raid build);
   - the walk's ten band rows and the bear;
   - round 12's list.
   TestCommittedAddonBuildsFollowTheGuides and leveling-builds.test.ts fail until then.
5. **sim/request TestRotationLadder goldens** at the pin. These move:
   - every mana spec (regen phase);
   - the hunters, Subtlety (30 to 60), the bear (20 to 40);
   - Elemental, Balance and Enhancement;
   - priest-holy, priest-shadow, druid-restoration;
   - the walk's nine specs at bands 20 to 40;
   - mage-fire 30, warlock-destruction 40, paladin-protection 40 (rank rows);
   - round 12's list.
   Also sim/adapter golden_test if it sims a mana user.
6. **addon-data** (`pipeline addon-data` → addon-data.json and Data.lua): `raceSpellId` rank rows (replaces round 12's factionSpellId, which no shipped Data.lua carried), every changed rotation, trainables' rank lists, round 12's list.
7. **Coach rules and starters:** pipeline.coachrules (bear and three rogue files); `python tools/coach/persona_starter.py --all`: rogues ×3, hunters BM and MM, Arms, Fury, round 12's list. test_persona_starter.py and coach.yml fail until then.
8. **dps-per-point.json** for every spec (regen phase, Elemental's haste fix), then Data.lua.
9. **Fork, at the next engine follow:** the conformance trainables copies and goldens.
10. **Nightly-owned, at the refresh:** every BiS file (see "Measure at refresh").
11. Two app fixtures still carry the old 70707 Elemental build code (casters.md).

Already done on the branches, not to redo: the fork goldens (E c555683de, E 4f81a2a92 and the lanes' own), planner-rules.json and the mobile copy (platform, as the fix's test vectors), sim/rankchoices/rank-choices.json.

**Gates.** No lane ran a full suite. On int/r13 I ran the fork's sim/core, druid, hunter, mage, paladin, priest, rogue, shaman and warrior packages (green after M6), the merged sim/request tests M3/M4 touched, and the guide prose test. The full gate set runs once, at the landing. Not run: iOS `xcodebuild test` (platform used a swiftc harness over the same vectors).

Heads at the end of round 13: fork `int/r13` 919616e88; site `int/r13`: the commit that records this line (the one before it adds this report).

## Measure at refresh (added this round)

- **MR20:** hunters on the hotfix push: every band; Marksmanship's Lone Wolf raid build; gear without spell power (MR14) still holds.
- **MR21:** Restoration druid and Holy priest picks after the rotation fixes; the druid's raid scale (coverage about 95%) for an owner retune.
- **MR22:** the walk's band builds and the bear's: picks at bands 20 to 40 for the ten specs.
- **MR23:** Subtlety picks and weights (armor ignore as a share); Shadow priest with Vampiric Embrace.
- **MR24:** Elemental bare 50/60 (Tidal Focus) and Balance bare 30 to 60 (leveling build); Affliction raid (Improved Bane of Agony); Fury raid (Raging Blows).
- **MR25:** dps-per-point and haste weights after the regen phase and Elemental's Lava Burst tie.

## Still open

- **Owner decisions:**
  - Protection paladin band 20: Templar's Bulwark (survival) over the second Sacred Duty point (-2% threat).
  - Same-instant ties in core (Penance's cooldown ending with a cast; channel pushback on a pulse): Discipline reads about 1% high at exactly 0% haste. Break ties fairly or jitter, in the spirit of ruling 4?
  - The Restoration druid's heal scale (1.85) at about 95% coverage.
  - The tank stand-in boss has no attack power for Demoralizing Shout/Roar to remove; changing tank-encounter.json moves every tank.
  - The Balance leveling build drops Improved Entangling Roots (the reviewer agrees).
- **Owner checks in game, before launch where noted:**
  - Arcane Shot's tooltip with spell power up (ruling 5).
  - Gcloud reauth for the level-60 split.
  - Era book ranks (33 across nine classes, e.g. Immolate 25309, Heroic Strike 25286): does a level-60 trainer teach them? The client lists them like trainer ranks.
  - Six talents Wowhead tags baseline (Toughness, Improved Cleave, One-Handed Weapon Specialization, Piercing Ice, Arcane Focus, Magic Attunement): a level-1 spellbook check.
  - Heart of the Lion 280 or 276 at 60; Rejuvenation rank 1's tooltip.
  - The push's values once Wowhead re-scrapes.
- **Logs (level 30+ / 60):** Deep Wounds' double multiplier, Survival's Raptor Strike queue miss gap, Holy Shield crit, Eclipse timing, Elemental Mastery's charge, Improved Moonfire, mage cost stacking, Master of Elements on free crits, Ice Lance's coefficient, demon crit, Ignition, Siphon Life tick crit, Gift of Nature on Tranquility, Kings with Heart of the Lion, the Judgement of Wisdom/Light proc roll.
- **Held under the adoption rule:** Arms raid build (+0.2%); Destruction's Immolate line (+0.12 to +0.23%); Elemental Natural Grace to Restorative Totems (+0.65%); cat Clearcasting Shred at 20 (+0.4 to +0.7%); Demonology's in-tree Improved Sayaad (+0.08%); Lava Burst, Cold Blood, Shadowburn, Last Stand and Riptide at their bands (under 1%).
- **Not done:** Discipline 35 and Restoration shaman 46 to 49 lose their capstone between bands (addon order and strips only); Demonology's Soul Link at -2 s starts the pull 8 to 12% short of mana (+0.14 to +0.24%); the "Endless" elixirs; the demoralizing_shout debuff option's Era 146; band 38 has no BiS entry.
- **The round was not zero:** round 14 follows the landing.
