# Sim-accuracy round 11: integration (2026-11-02)

The owner's round-11 rulings, the integration, the pin and the review. The lane reports beside this file have the per-spec rows: regression.md (new this round), warlock.md, casters.md, physical.md, hybrids.md, healers.md, healprofile.md, crit.md, auras.md, items.md (with items-165.md), ranks.md, ranker.md and platform.md.

Client 1.60.1.70753 went live on main during the round (ddb7d4601, engine follow 9f42ff1b3). The fork branch merges that follow, and every number below is measured on 70753.

## Owner rulings, and where each landed

| # | ruling | outcome | where |
|---|---|---|---|
| 1 | Sub-60 heal profile: a pull-size multiplier from real dungeon pulls | **Done.** A band below 60 fights the harder of its median last boss and its 75th-percentile trash pull, weighted by mobs (a quarter of a run's mobs come in pulls at least that hard). A pull is a creature plus every same-faction creature in its assistance radius, chained as mangos-classic pulls; the radius is the creature's own CallForHelp, else 10 yards. Every pull is cited spawn by spawn from cmangos classic-db. Multipliers: 20 ×1.61, 30 ×2.06, 40 ×3.05, 50 ×1.85; 60 is unchanged. Bare coverage below 60 falls from 72-98% to 42-90%, and a small gear-swap check separates picks by 1 to 15% where the old profile moved them 0 to 0.2%. Every sub-60 healer rank row was re-measured; five band builds adopted under the adoption rule. The BiS page's disclosure names the hard pull and says it is modelled, not average. The full pick check is "measure at refresh". | healprofile.md |
| 2 | Conviction on paladin heals follows the client | **Keep.** Conviction's effect, SpellEffect 698173, shipped as aura 52 (melee and ranged crit). The hotfix push 113139 (first seen 2026-10-31) makes it aura 290, crit on every table, with no school mask (EffectMiscValue_0 = 0), no class mask (EffectSpellClassMask 0,0,0,0) and no SpellClassOptions row; SpellEquippedItems 180908 only requires a melee weapon. 70753 ships the same row. So heals keep the crit, and Holy damage spells get it back: round 10 had removed it on logs that all predate the push. | crit.md CR1 |
| 3 | Retribution bare 20's 0.3 to 0.5% under the paced Consecration line | **Accepted**, recorded; no rank choice can apply at 20 (only Consecration rank 1 is learned). | ranks.md |
| 4 | The 165 item ids gone from wago's export | **Not gone.** None is in the shipped 70707 ItemSparse; every one is in the item-hotfix layer that `fetch-wago-hotfixes` writes (126 from 70532's push, 39 from 70707's). Round 10's local run skipped that step. On 70753 the 39 ship in the client itself and the 126 survive through 70532's push. No BiS set, guide or addon names any; loot names 2 (Precision Arrow, the Scarlet Crusade tabard). The refresh must NOT accept a drop of these; `fetch` now fetches the layer itself (S 3484405dd). The data run still needs `--allow-shrink` for the separate wowhead-stats gate (289 against 689). | items.md, items-165.md |
| - | Late finds | Blackhand's Breadth and Eye of the Beast: 1% melee/ranged and 1% spell crit, with their use effects (crit CR2, CR3). Restoration druid band 20 "5051": the ranker's "Furor" was the Restoration tree's first digits; that build loses 24% to the current ladder, not adopted. | crit.md, healprofile.md |
| - | PvP cap (ruled during the round) | Defaults are capped at in-game rank 6 (the item field's 10), the line the published picks always drew; ranks 7 and above stay labelled alternatives at every band. The cap now reads the in-game rank (70753's CharTitles has 14 honor titles a side). No published default pick moves; the ranker's K22 is withdrawn as a bug, its label fix (K23) stands. | ranker.md; S 1e1b85d8d |
| - | Main red from the 70707 addon-data re-render | Seven api tests used Overpower as an off-rotation spell after Fury's lines gained it. Fixed on main on its own (S cd0cfbb0e, Cleave rank 5). | S cd0cfbb0e |
| - | The fork's CI red since 40021374a | `TestConformance/warlock`: Destruction's Immolate rank 1 at 10 to 40 has 0.200 × 0.2875 = 0.0575, exactly on a three-place boundary; arm64's fused multiply-add rounds it to 0.058, amd64 (CI) to 0.057, so a golden written on a Mac never matched CI. The cell prints a penalised coefficient to four places; every golden regenerated; the whole fork suite passes as amd64 (GOARCH=amd64 under Rosetta). | E 47df4a89d |

## Confirmed and fixed, by lane, with origin

Origin: R9 or R10 = caused by that round's fix (the causing sha is in the lane row); old = predates round 9 and was not found before; R11 = a side effect of another round-11 fix, caught before merge.

| lane | confirmed | R9 | R10 | old | other | the bugs |
|---|---|---|---|---|---|---|
| regression | 5 | 3 | 2 | 0 | | talent-search and rotation-search sent no preset (bare-60 rank choices ignored); the coach flagged Elemental's own Lightning Bolt rank 9; a fork UI preset used Redoubt's old id; Healing Stream took a school's crit; bear notes named Stoneshield |
| warlock | 3 | | | 3 | | banes and curses in separate slots; the warlock's Curse of the Elements stacked on the raid's; a missed Conflagrate ate Immolate |
| casters | 6 | | | 6 | | Arcane Clearcasting -1000%; Improved Scorch on a hit; Arcane Missiles rank 1; Undead Shadow never cast Dark Sacrifice; Elemental's own Mana Spring; Robe of the Archmage and the mana gem's minute |
| physical | 4 | | | 4 | | Execute's per-rage damage by rank; Whirlwind's off hand rolled Windfury Totem; rage from damage taken at the boss's level (found by hybrids); a Survival note |
| hybrids | 1 | | | 1 | | Benediction did not discount Holy Shield |
| crit | 5 | | 1 | 4 | | Conviction on every table; Blackhand's Breadth and Eye of the Beast tables and use effects; Atiesh's spell crit; Forever's scopes and Biznicks |
| items | 6 | | | 6 | | four melee set procs on any hit; Thunderbrew's Boot Flask never crit; classic-db stats where the client states Forever's; ranged-only hit; Blackfathom stone's hit reached spells; the fetch skipped the item push |
| auras | 5 | | | 4 | 1 mixed | the plain Mana Spring under Restorative Totems; the warrior tank never shouted; Battle Shout's threat; a Commanding Shout note; no main-hand stone at 35-49 (found by regression; R9 for Survival, old for the rest) |
| healers | 4 | | 1 | 3 | +1 R11 | heals learned below 20 lost coefficient (beta-log fit, seven characters); Mender's potions; 30-second potions never fired for healers (R10); Holy priest's races (found by ranker); R11: Discipline fills with Lesser Heal rank 3 |
| ranks | 2 | 1 | 1 | | +2 R11 | Lava Burst rank 2 at 50-59 (R9); Restoration druid's Rejuvenation 10 row (R10); R11: Shadow's Devouring Plague row after Dark Sacrifice, Restoration shaman's Healing Wave rank 2 after the heal fix |
| ranker | 1 | | | 1 | | PvP source lines named the client field, not the rank |
| platform | 9 | 2 | 1 | 6 | | the request drawer dropped the preset (R9); the addon's rotation ids (R9); the Enhancement guide's totems (R10); weights tool against the SimC ruling; server runs turned execute and variation off into defaults; the coach's book ranks; Affliction's coach rules; the Restoration druid starter; the addon's talent spells |
| integrator | 1 | | | 1 | | the fork CI's arch-dependent golden (above) |
| **total** | **55** | **6** | **6** | **39** | **1 mixed, 3 R11** | |
| healprofile (ruling 1 follow-on) | 11 | | 9 | 2 | | the pull multiplier, its page line, re-measured rank rows and five band builds |

**Is the count falling?** Yes. Round 10 fixed 88; round 11 fixed 55 (66 counting the ruling-1 follow-on). Of the 55:
- 13 were caused by rounds 9 and 10: 6 by round 9, 6 by round 10, and the stone fix, mixed (old for most melee specs, round 9 for Survival).
- 39 were older bugs no round had found.
- 3 were side effects of other round-11 fixes, caught before merge.

The regression lane alone found 5 of the 13, the other lanes the rest.

## Integrator findings

| # | finding | evidence | fix |
|---|---|---|---|
| I1 | The regression lane (K2) and the platform lane (C2) fixed the same coach bug; counted once (under regression). Platform's BandRanks dropped, its test kept on Rotation.Naming. | merge conflict f5_ranks.go | S int/r11 platform merge |
| I2 | The healprofile lane measured Discipline's sub-60 rows before the Lesser Heal swap; its Heal rank-4 rows at 40-59 named a spell the rotation no longer casts. Dropped; Heal rank 4 instead of Lesser Heal rank 3 measures +0.4/+0.5% at 40 (under the 1% bar) and -0.9% at 50. | spec-breakdown, 10,000 iterations, seeds 1 and 50001 | S int/r11 healprofile merge |
| I3 | rank-choices.json merged by band: ranks' level-60 rows beside healprofile's sub-60 rows. | merge conflict | S int/r11 ranks merge |
| I4 | Merged-state test drift: rotation-search's preset test anchored on the dropped Devouring Plague row; Survival's APL copy; coach starters; the presets test against a graded id; the guides' Dense stone line; the mono-face registry; a weights test type error. | gates | S 1e1b85d8d |
| I5 | The fork CI's conformance golden (above). | GOARCH=amd64 reproduces | E 47df4a89d |

## Integration

- **Fork `int/r11`** merges lane/r11-regression, -warlock, -casters, -physical, -hybrids, -crit, -items, -auras and -healers (healprofile is healers fast-forwarded; ranks and ranker have no engine commits), then 9f42ff1b3 (the 70753 follow), five goldens regenerated on the merged state, and the conformance fix. The fork suite with `-tags=with_db` passes as arm64 and as amd64; the wasm cross-compile builds.
- **The pin is 47df4a89d.** data/proto follows it (Mender's potions). The rotation ladder goldens are regenerated at it with `make simdb` on 70753: seventeen specs move, each from a lane fix.
- **The site** merges all thirteen wt/r11 branches and main through ddb7d4601 and later.
- **Gates** (`nice -n 15`, GOMAXPROCS 9, one at a time): sim vet and test; logs vet and test; api vet and the coach and sims packages; the api on a checkout with no `make simdb` copy (vet, build, and no engine package linked); `make apl-check` against the pin; `make styles-check`; tools/tests/test_wasm_compress.sh; tools/coach/test_persona_starter.py; the data pytest; web sync, vitest, astro check, lint and format:check. Results in "Gates" below.
- **No regens.** Under the owner's loop structure no data run, simdb rebuild or BiS ranking ran this round; picks that need one are on the ledger's "Measure at refresh" list. `addon-builds.json` follows the healer band builds through its committed-copy test (as in round 10), not a regen.

## What 70753 changes for the sim

Fetched read-only from wago.tools for 70707 and 70753, each layered with its build's push, diffed row by row (scratchpad r11/c70753/):
- **No hotfix push for 70753 yet.** Its shipped tables already include 70707's push: 1,467 push changes to SpellMisc come down to 5 differences.
- **Talents:** no Talent, TalentTab or Trait row changed. SkillLineAbility only gains Poised.
- **Poised (new):** rogues learn it at 18, druids at 32. At full Energy out of combat, the next ability costs 10 less (aura 107, cost modifier -10): the opener for rogues and cat. The engine has its numbers but does not simulate it. Round 12.
- **Holy Shield 1239570:** now a whole level-60 spell (240 mana, 10 s cooldown, 0.08 coefficient). The activation reads it as rank 4; whether the Protection paladin rotation casts it is round 12.
- **Orc Command (20575, 20576, 21563):** gains hunter and warlock class-option rows and a pet-damage modifier with mask 12288; round 9's "Orc Command on pets" is re-checked in round 12.
- **Removed:** Salvation, Mass Salvation, Blessing of Light, Greater Blessing of Light, Instrument of Law, Revoke. No preset or engine buff uses them.
- **Items:** 40 new ItemSparse rows; 37 change price only; no stats.
- **Effect on this round's fixes:** none is made wrong or unnecessary. Conviction's aura-290 row ships in 70753 itself, the 165-id verdict holds, and the heal coefficient fix rests on logs, not a table row. Every number below is on 70753.

## Before/after

Engine 9f42ff1b3 with main (the 70753 activation, no round-11 changes), against pin 47df4a89d with this branch, both on 70753 with `make simdb`. Raid preset, band 60, 10,000 iterations. Healers in effective HPS on the heal profile, tanks in threat per second.

| Spec | Alliance, seed 1 | Alliance % (1 / 50001) | Horde, seed 1 | Horde % (1 / 50001) |
|---|---|---|---|---|
| druid-balance | 566.5 to 566.9 | +0.1 / +0.2 | 571.6 to 572.2 | +0.1 / +0.0 |
| druid-feral-bear | 1001.0 to 1014.2 | +1.3 / +1.3 | 992.4 to 1004.7 | +1.2 / +1.2 |
| druid-feral | 674.7 to 665.0 | -1.4 / -1.3 | 670.0 to 660.0 | -1.5 / -1.4 |
| druid-restoration | 830.9 to 834.6 | +0.4 / +0.4 | 830.0 to 833.7 | +0.4 / +0.4 |
| hunter-beast-mastery | 629.4 to 629.3 | -0.0 / +0.0 | 633.4 to 633.5 | +0.0 / -0.0 |
| hunter-marksmanship | 618.5 to 618.4 | -0.0 / -0.0 | 623.0 to 623.0 | +0.0 / +0.0 |
| hunter-survival | 741.1 to 741.1 | +0.0 / +0.0 | 744.7 to 744.7 | +0.0 / +0.0 |
| mage-arcane | 594.3 to 605.1 | +1.8 / +1.8 | 600.4 to 610.9 | +1.7 / +1.7 |
| mage-fire | 651.9 to 655.7 | +0.6 / +0.5 | 660.1 to 662.1 | +0.3 / +0.3 |
| mage-frost | 617.0 to 617.0 | +0.0 / +0.0 | 623.0 to 623.2 | +0.0 / +0.0 |
| paladin-holy | 653.6 to 656.5 | +0.4 / +0.4 | 641.4 to 645.1 | +0.6 / +0.6 |
| paladin-protection | 1000.7 to 1009.0 | +0.8 / +0.8 | 1013.5 to 1021.9 | +0.8 / +0.8 |
| paladin-retribution | 688.7 to 690.6 | +0.3 / +0.2 | 685.7 to 687.7 | +0.3 / +0.2 |
| priest-discipline | 623.0 to 649.8 | +4.3 / +4.3 | 625.3 to 653.5 | +4.5 / +4.5 |
| priest-holy | 705.8 to 708.9 | +0.4 / +0.4 | 693.8 to 696.7 | +0.4 / +0.4 |
| priest-shadow | 710.8 to 711.0 | +0.0 / +0.0 | 717.8 to 717.7 | -0.0 / -0.0 |
| rogue-assassination | 621.8 to 621.8 | +0.0 / +0.0 | 618.8 to 618.8 | +0.0 / +0.0 |
| rogue-combat | 616.8 to 616.8 | +0.0 / +0.0 | 617.0 to 617.0 | +0.0 / +0.0 |
| rogue-subtlety | 544.1 to 544.1 | +0.0 / +0.0 | 544.2 to 544.2 | +0.0 / +0.0 |
| shaman-elemental | 459.6 to 462.8 | +0.7 / +0.7 | 466.9 to 470.2 | +0.7 / +0.7 |
| shaman-enhancement | 580.3 to 580.3 | +0.0 / +0.0 | 580.8 to 580.8 | +0.0 / +0.0 |
| shaman-restoration | 836.4 to 842.8 | +0.8 / +0.7 | 834.0 to 840.2 | +0.7 / +0.8 |
| warlock-affliction | 569.3 to 570.2 | +0.2 / +0.2 | 571.0 to 572.1 | +0.2 / +0.2 |
| warlock-demonology | 635.6 to 636.5 | +0.1 / +0.3 | 640.4 to 641.4 | +0.2 / +0.2 |
| warlock-destruction | 634.6 to 635.0 | +0.1 / +0.1 | 636.9 to 638.0 | +0.2 / +0.1 |
| warrior-arms | 707.7 to 707.7 | +0.0 / +0.0 | 719.6 to 719.6 | +0.0 / +0.0 |
| warrior-fury | 752.8 to 749.4 | -0.5 / -0.4 | 756.8 to 754.8 | -0.3 / -0.4 |
| warrior-protection | 1079.0 to 1082.8 | +0.4 / +0.4 | 1067.5 to 1071.5 | +0.4 / +0.3 |

Where the moves over 1% come from:
- **Discipline +4.3 to +4.5%:** heals learned below 20 keep their coefficient (about half), and Lesser Heal rank 3 fills the party.
- **Arcane +1.7 / +1.8%:** Clearcasting makes a stacked Arcane Blast and an Arcane Power cast free, as the client's -1000% says.
- **Bear +1.2 / +1.3% threat:** rage from damage taken uses the bear's own level.
- **Cat -1.3 to -1.5%:** the four melee set procs roll on auto-attacks only.

Under 1% at 60:
- Protection paladin and Restoration shaman gain about 0.8%; Elemental 0.7%; Fire 0.3 to 0.6%; Holy paladin, Restoration druid, Holy priest and Protection warrior about 0.4%; Retribution and the warlocks 0.1 to 0.3%.
- Fury loses 0.3 to 0.5% (Whirlwind's off hand no longer rolls Windfury Totem).
- Hunters, rogues, Arms, Enhancement, Frost, Shadow and Balance are within 0.2% at raid.

Their changes land below 60:
- Warlock bare bands gain 2 to 8% (Curse of the Elements in the rotation).
- Arms and Fury lose up to 6.8% at 30 to 50 (Execute).
- Horde Shadow gains 3 to 16% at the bare bands (Dark Sacrifice).
- Rogues, Retribution and Arms lose 1 to 3% at bare 20 (set procs).
- The healers move with the pull profile (healprofile.md has every band).

## wow-player review

Verdicts on the lane reports, the screenshots and the follow-up evidence:
- **SHIP:** regression, warlock, casters, physical, hybrids, items, ranks.
- **SHIP after fixes**, every fix done:
  - **healers:** say plainly that Lesser Heal is a mana filler (guide reworded).
  - **healprofile:**
    - The BiS row says the hard pull is modelled, not average.
    - Healing Wave rank 2 at 40 and 50 was held until its coverage and healing power were shown: 261 healing power, about 238 a cast at 5.6 healing per mana, coverage 68% to 90%. Accepted as following from the heal coefficient fix, with "verify on launch logs".
  - **crit:**
    - The guides say Conviction needs a melee weapon.
    - Conviction on Consecration, Holy Shock and Holy Wrath is listed to verify on the first post-push log.
  - **auras:**
    - No shaman row carries an off-hand consumable, now held by a test.
    - The Mana Spring note watches forever-bugs #215.
  - **ranker:**
    - Its HOLD on the PvP labels was withdrawn once the client's CharTitles showed 14 titles a side.
    - The cap is ruled at rank 6.
  - **platform:**
    - The Enhancement loop was rewritten as player advice.
    - The weights concern was withdrawn: the shot was a fixture build, and real factors show Strength = 2 attack power.
- **Release-note lines** (the reviewer's):
  - "Heals learned before level 20 now keep their full spell power scaling; low-level healers heal more than before."
  - "Holy priest guide now recommends human or undead, and Discipline fills with Lesser Heal rank 3."
  - "Conviction now improves crit on all spells and attacks, including heals and Holy damage spells."
  - "Rage from damage taken now uses your level, not the boss's. Bears generate a little more threat at 60."
  - "Execute's extra damage per rage now follows the rank, so Arms and Fury below 56 hit slightly softer."
  - "Banes and curses are tracked separately. Warlocks can run Agony or Doom alongside a curse."
  - "Arcane Clearcasting now makes your next damage spell fully free, even with Arcane Blast stacks or Arcane Power."
  - "Defias, Shadowcraft, Valor and Bloodmail procs only roll on auto-attacks."
  - "Below 60, healer figures assume a hard dungeon pull as well as the last boss."
  - "PvP gear sources now show the in-game rank."
- **Still to check before release:** one rendered Enhancement guide and one weights page from a real-data build with the pinned wasm.

## Gates

All exit 0 on the pinned state (int/r11 with engine 47df4a89d, client 70753):
- **sim:** go vet; go test ./...
- **logs:** go vet; go test ./...
- **api:** go vet; the coach and sims packages.
- **api on a clean checkout:** go vet and go build on a checkout with no `make simdb` copy, and no engine package linked.
- **Repo checks:** make apl-check against the pin; make styles-check; tools/tests/test_wasm_compress.sh; tools/coach/test_persona_starter.py.
- **data:** pytest, deselecting only test_simbuffs_names_every_id_the_engine_lets_a_request_send (simbuffs.json is held output).
- **web:** sync; vitest (8,066 tests); astro check; lint; format:check.

The first vitest run failed once in src/sim-island.test.ts: the min-h reservation test fails about one run in five under load (5 local reruns: 4 pass). It predates the round (1b932d73b) and passed on the rerun; it is listed below.

The fork suite with `-tags=with_db` passes as arm64 and as amd64, and the wasm cross-compile builds.

## Still open

- **Round 12 scope (new in 70753):** Poised for rogues and the cat; Holy Shield rank 4 in the Protection paladin rotation; the Orc Command rows (round 9's pet fix).
- **Owner decisions:**
  - Do rank-choice rows also need the adoption rule's 1% bar? Six kept rows beat error by less: three Aimed Shot rows, Fire Blast rank 4 at 40, Shadow Bolt rank 4 at 30, Prayer of Healing at bare 60. Lava Burst rank 2 is +0.92% at seed 1.
  - Rank rows have no faction key: Alliance Shadow gives up 2.1% at bare 60 to the dropped Devouring Plague row.
  - Elemental's bare 50 and 60 build takes Improved Healing Wave over Tidal Focus (+5.7% if swapped; the guides call it a leveling choice).
  - The heal profile's tank-versus-group split and line of sight in the pull rule, and whether sub-60 healers show coverage as their headline.
- **Verify on logs:**
  - Conviction on Consecration, Holy Shock and Holy Wrath after the push.
  - The heal coefficient fix at 40 and above, including Healing Wave rank 2.
  - Improved Moonfire's reach; Searing Totem's cadence; Elemental Mastery's charges; Ice Lance's pushed cooldown; Fire Power on DoT ticks.
  - A queued Raptor Strike's off-hand penalty (Survival -0.8% raid if it applies to hunters).
  - Healing Stream crit without Thundering Strikes.
  - Percent spell modifiers adding or multiplying.
  - Round 10's log-blocked list.
- **Held under the adoption rule:** Marksmanship (+0.87 to +0.93%), Affliction's Improved Bane of Agony (+0.72%), Destruction's Immolate line, Fire's Improved Scorch line, Frost's search winner, Frost bare 20's Ice Shards, the druid, Holy and shaman band builds the reservation could not express.
- **Refresh:** the "Measure at refresh" list in the ledger.
- **Main's mage guide rewrite (e2b3a9a53, not this round)** landed after the before/after was measured. It changed Frost's leveling build (it now casts Arcane Missiles) and three mage APL notes. The merge regenerates the mage ladder goldens, addon-builds.json and the mage coach starters, and points rotations.test.ts at the new notes. Frost's bare-60 ladder rung moves +14%: round 12 should audit that build under the adoption rule.
- **A flaky web test:** src/sim-island.test.ts's "removes the shell min-h reservation" fails about one run in five under CPU load; likely an earlier test's real-timer boot landing in the next test's document. Predates the round (1b932d73b).
- **Healers lane coverage:** it hit the subagent limit and tested round 10's client-row conclusions against new evidence rather than re-reading every row; round 12 should give the healers a full row sweep.
