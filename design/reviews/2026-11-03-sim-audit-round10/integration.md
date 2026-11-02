# Sim-accuracy round 10: integration (2026-11-02)

The owner's round-10 rulings, the integration, the pin and the review. The lane reports beside this file have the per-spec rows: warlock.md, casters.md, physical.md, hybrids.md, healers.md, healprofile.md, ranks.md, crit.md, auras.md, items.md, ranker.md and platform.md. Healprofile is new this round, for ruling 1.

## Owner rulings, and where each landed

| # | ruling | outcome | where |
|---|---|---|---|
| 1 | Scale the heal profile by band; re-measure every sub-60 healer rank and build on its own band | **Done.** `bands` in data/curated/heal-profile.json, applied only by sim/request/healprofile_bands.go, so /sim, the API, the ranker, spec-breakdown, heal-builds and dps-per-point all fight the healer's own band. The rule: a band's content is the open dungeons whose last boss is from the band's level to five above it (Onyxia at 60); its damage is the median of those bosses' melee (less cast time) plus the spells their AI casts while tanked (cmangos creature_template, spell_template, creature_spell_list, EventAI, ScriptDevAI; each row cites its source), times the client's x1.2 Classic-dungeon creature damage modifier; health follows the client's expected player health at the level. 60 is unchanged. Every sub-60 healer rank row was re-measured on its band; healer band builds can now take curated ranks first (data/curated/ladder-first.json), adopted for Restoration druid (band 20 +24%, band 30 +2%), Discipline (band 20 +2%) and Holy (band 40 +7%). DPS specs are byte-identical. | healprofile.md |
| 2 | Retribution's Consecration: remove the needless mana gate, re-measure the rank | **Done, with one correction to the premise.** The 65% gate was what kept Retribution from running dry: with no gate the paladin spends 76 to 115 s of a 180 s fight out of mana at bare 20 to 50. The flat line is gone; Consecration now casts while the mana left after it still pays for seal, Judgement and Holy Strike and mana keeps pace with the fight. Rank 1 at 30 to 59 still wins beyond error on the new line AND with no gate at all, so it is a real mana-bound win and stays. New row: rank 4 at bare 60 (+1.6%). | ranks.md |
| 3 | The Watchers: launch content? | **Not launch content; the gate stays "later", now "inferred".** Blizzard's raid overview (news.blizzard.com article 24303009) makes the Barrow Deeps the wardens' prison, a raid that opens 9 December; the 70707 client ties the faction to it (item 277136 "Grants access to the Vord'anaar Prison within Barrow Deeps"; map 2817 Starfall Barrow Den). No Blizzard post, client quest or reputation row, or Wowhead Forever page names a launch zone or a way to earn it. Kept at "later" rather than "raids-1" because nothing proves the rewards come only from the raid. The awaiting pass for 2819 stays until the data run. | ranker.md; S 3741c143f |
| 4 | Brotli for sim.wasm | **Done.** The Worker answers /_sim/*/sim.wasm.gz with the sim.wasm.br bytes and Content-Encoding: br for a browser that accepts br (encodeBody 'manual'), and the gzip for everyone else and on any miss; Vary: Accept-Encoding. 3.10 MB against 4.31 MB. CI budgets both files (WASM_BR_BUDGET_MB 3.5). Production's brotli response is unconfirmed until the first deploy: `curl -sI -H 'Accept-Encoding: br' https://foreversixty.gg/_sim/<dir>/sim.wasm.gz` should show content-encoding: br. A failure costs only the saving, never the page. | platform.md; S 4a493bcd1 |
| 5 | The Survival guide's time-bound line | **Done.** No date, no number, no "fell"; a test rejects that wording. | platform.md; S 15767bbd6, 135d6bfb6 |
| 6a | Coach persona test's failures | **Done.** 12 failures, not 7 (two code bugs, the rest stale tests). All 25 starters regenerated; new .github/workflows/coach.yml runs the test on every change to its inputs; the integration gates now run it. | platform.md; S 31debfef2 |
| 6b | "Apply to the page" drops the preset | **Done**, on the browser, drawer and server paths; Top Gear and the weights tool sent no preset either (fixed). | platform.md; S c4d905f6b, bf3236777 |
| 6c | Icons on the custom /sim panel | **Done in the pipeline; live at the first data run after the hold.** Three causes: the icon step never fetched the panel's icons (98 missing), the 96 oil and stone rows were never read, 29 rows had the wrong art. After the wow-player HOLD the weapon section is built per class: one slot per hand, off hand only for warriors, rogues and hunters, grouped by kind, full names (S 414b0d9b5). Shots in shots/. | platform.md |
| 6d | Warsong Gulch faction ids | **Done.** The fork's names table had 889/890 backwards and the pipeline trusted the names; ids now stay as mined (889 Warsong Outriders = Horde, 890 Silverwing Sentinels = Alliance), names come from the client, and a reward mined under the other side's faction stops the build. Offline loot-merge: only the 8 WSG sources change, all 124 placements on the right side. The ranker's repSide and the content-phases labels follow. | items.md (S 34e994332); ranker.md (S cdb749068) |
| - | Extra attention: the crit split | **Swept engine-wide.** Every crit roll in every published spec and band traced to its client defense type; 41 crit talents, the crit sets on published gear, simdb's crit mapping and the 22 places left from before the split checked. Fixed: Conviction (melee and ranged, not Holy damage spells), Gloves of Holy Might and Thornspike (ranged table), school crit to heals, Emerald Dragon Whelp, the stone on hunter pets, Weaponmaster and Hack and Slash (weapon attacks only), demon crit (the Succubus), hunter pet spell crit, Healing Stream (every-table crit only), Combustion off Ignite, Inner Focus off Renew, Libram of Infusion on the Holy Shock heal, Exorcism never crits. | crit.md and the class reports |

## Confirmed and fixed, by lane

| lane | confirmed | the bugs |
|---|---|---|
| warlock | 3 | demon crit (melee 0.7% with no Agility, Lash of Pain 6.4%, Firebolt 2.2%); Spellstone is Forever's; Touch of the Grave rolls spell hit |
| casters | 5 | Arcane Missiles spent Fingers of Frost and took Shatter; Frost's Ice Barrier note; Combustion crit on Ignite; Improved Flamestrike counted twice; Precision on item spells and the wand |
| physical | 6 | Weaponmaster crit on every table; Hack and Slash on poisons and procs; Lacerate and Lacerating Strikes missed Savage Strikes and Predator's Edge; Summon Hawk's talents; hunter pets' spell crit; Web's table. (Arena Grand Master was settled Aligned: its -6% works only in the Battle Ring and The Maul.) |
| hybrids | 5 | Retribution's twist swapped Seals every global; Elemental Devastation's id; Exorcism crit (0 of 641 logged); the bear never drank its potion; stones never reached a feral's weapon |
| healers | 11 | Swiftmend consumed the HoT; Swiftmend's line; Gift of Nature on Tranquility; Nature's Grace on heal crits; Healing Stream's crit; Totem of Urgency unregistered; Libram of Infusion; Mental Agility on Inner Fire; Inner Focus on Renew; two note and guide fixes |
| healprofile | 3 | sub-60 bands fought the level-60 fight; the BiS page showed the level-60 profile on every band; sub-60 healer builds were never chosen for their band |
| ranks | 3 | Retribution's Consecration gate; a stale guide; stale APL notes |
| crit | 5 | Conviction on Holy spells; Gloves of Holy Might's table; school crit to heals; Emerald Dragon Whelp; the stone on hunter pets |
| auras | 3 | Forever's 30-second potions missing; druids kept Era's shapeshift rules; tanks carried no weapon imbue (Protection warrior Elemental Sharpening Stone, Protection paladin Brilliant Wizard Oil) |
| items | 30 | Warsong Gulch ids; Thornspike's table; Mindtap Talisman; five Forever use effects; Blazing Rapier and twenty more retuned or Forever-replaced procs (Gloves of Holy Might counted under crit) |
| ranker | 4 | gated vendor stock opened by round 9's client reputation sources (27 Undermine items, picked in 16 local rows); the WSG ids in repSide and the content-phases labels; the healer trinket and relic re-measure asked for 14,000 iterations and was refused, so a gain inside error published empty; the BiS report's stale one-crit text |
| platform | 9 | the three owner page items; the panel's icons, oil/stone rows and wrong art; Top Gear and weights sent no preset; the coach flagged the site's own downranks; the coach's potion list |
| integrator | 1 | see below |

## Integrator findings

| # | finding | evidence | confirmed? | fix sha |
|---|---|---|---|---|
| I1 | Retribution's top rotation note still said the list never twists and the build skips Unsealed Wrath, and the Consecration line named the review round. | the merged APL | CONFIRMED (copy) | S on int/r10 "docs(apl): Retribution's notes ..."; E 4d25f54a8 |
| I2 | Two lanes fixed Gloves of Holy Might's table (crit E 495779597, items); one copy kept, counted once. | client_procs.go conflict | merge | E int/r10 merge |
| I3 | The crit lane's per-spell stone and the auras lane's druid-forms fix conflicted in core/consumes.go; kept the per-spell stone without its feral gate (Forever stones work in forms), and both new functions. The auras form test read the stone as a stat; it now checks the claw's crit bonus. | consumes.go conflict; TestWeaponConsumablesReachAShapeshiftedDruid | merge | E int/r10 merge, test commit |
| I4 | The bear's potion test looked for Greater Stoneshield after the auras tank pass settled the bear's 60 potion on Mighty Rage. | TestTheBearDrinksItsStoneshieldPotion | test-only | S 6d8500f4f |
| I5 | rank-choices.json notes edited by ranks and healprofile; both kept, `make rank-choices-sync`. Restoration druid's coach starter regenerated after the merge. | merge | merge | S 40a7929ea and the starter commit |

## Integration

- **Fork `int/r10`** merges lane/r10-crit, -warlock, -casters, -physical, -items, -ranks, -auras, -hybrids, -healers and -healprofile, then the APL sync and TestProtection regenerated on the merged state (Conviction, Exorcism and Sanctified Orb each moved it in their own lane). Fork suite with `-tags=with_db` passes except `sim/web`, which needs a built binary_dist at every pin.
- **The pin is cfcb76e80.** data/proto follows it (four new potions). The rotation ladder goldens are regenerated at it with `make simdb`: fifteen specs move, each from a lane fix.
- **The site** merges wt/r10-crit, -warlock, -casters, -physical, -items, -ranks, -platform, -auras, -hybrids, -healers, -healprofile and -ranker.
- **Gates, all exit 0:** go vet and go test ./... in sim; go test ./... in logs; api's coach and sims packages; `make apl-check` against the pin; `make styles-check`; tools/tests/test_wasm_compress.sh; tools/coach/test_persona_starter.py; the full data pytest; web sync, the full vitest, astro check, lint and format:check.
- **Held-output exceptions:** the data gate still deselects `test_simbuffs_names_every_id_the_engine_lets_a_request_send` (simbuffs.json is data.yml output; it also now lacks the four potions). The ranker's `TestRepSideMatchesTheBuildsFactionIDs` skips with "awaiting the next data run" while the committed loot.json still carries Warsong Outriders under 890.

## Before/after

Pin 4d3129cae with origin/main, against pin cfcb76e80 with this branch, on today's published data and the simdb built from the committed inputs. Raid preset, band 60, 10,000 iterations. Healers in effective HPS on the heal profile, tanks in threat per second.

| Spec | Alliance, seed 1 | Alliance % (1 / 50001) | Horde, seed 1 | Horde % (1 / 50001) |
|---|---|---|---|---|
| druid-balance | 561.4 to 566.5 | +0.9 / +0.8 | 565.2 to 571.6 | +1.1 / +1.1 |
| druid-feral-bear | 831.6 to 1001.0 | +20.4 / +20.3 | 836.3 to 992.4 | +18.7 / +18.6 |
| druid-feral | 597.0 to 674.7 | +13.0 / +12.9 | 593.0 to 670.0 | +13.0 / +13.0 |
| druid-restoration | 780.3 to 830.9 | +6.5 / +6.4 | 779.2 to 830.0 | +6.5 / +6.4 |
| hunter-beast-mastery | 629.4 to 629.4 | +0.0 / +0.0 | 633.4 to 633.4 | +0.0 / +0.0 |
| hunter-marksmanship | 618.5 to 618.5 | +0.0 / +0.0 | 623.0 to 623.0 | +0.0 / +0.0 |
| hunter-survival | 731.4 to 741.6 | +1.4 / +1.4 | 734.9 to 745.3 | +1.4 / +1.4 |
| mage-arcane | 594.3 to 594.3 | +0.0 / +0.0 | 600.2 to 600.4 | +0.0 / +0.0 |
| mage-fire | 652.5 to 651.9 | -0.1 / -0.1 | 660.8 to 660.1 | -0.1 / -0.1 |
| mage-frost | 587.1 to 617.0 | +5.1 / +5.0 | 594.1 to 623.0 | +4.9 / +5.0 |
| paladin-holy | 653.6 to 653.6 | +0.0 / +0.0 | 641.4 to 641.4 | +0.0 / +0.0 |
| paladin-protection | 976.9 to 1000.7 | +2.4 / +2.5 | 993.8 to 1013.5 | +2.0 / +2.0 |
| paladin-retribution | 623.3 to 688.7 | +10.5 / +10.6 | 621.3 to 685.7 | +10.4 / +10.5 |
| priest-discipline | 622.4 to 623.0 | +0.1 / +0.1 | 624.9 to 625.3 | +0.1 / +0.1 |
| priest-holy | 705.5 to 705.8 | +0.0 / +0.1 | 693.4 to 693.8 | +0.1 / +0.0 |
| priest-shadow | 703.8 to 710.8 | +1.0 / +1.0 | 712.2 to 717.8 | +0.8 / +0.8 |
| rogue-assassination | 616.3 to 622.5 | +1.0 / +1.0 | 613.6 to 619.7 | +1.0 / +1.0 |
| rogue-combat | 611.5 to 617.1 | +0.9 / +0.9 | 611.8 to 617.4 | +0.9 / +0.9 |
| rogue-subtlety | 538.5 to 544.7 | +1.2 / +1.2 | 538.3 to 544.5 | +1.2 / +1.2 |
| shaman-elemental | 459.6 to 459.6 | +0.0 / +0.0 | 466.9 to 466.9 | +0.0 / +0.0 |
| shaman-enhancement | 575.5 to 580.3 | +0.8 / +0.8 | 575.8 to 580.8 | +0.9 / +0.9 |
| shaman-restoration | 839.9 to 836.4 | -0.4 / -0.5 | 838.0 to 834.0 | -0.5 / -0.5 |
| warlock-affliction | 571.5 to 569.3 | -0.4 / -0.4 | 573.2 to 571.0 | -0.4 / -0.4 |
| warlock-demonology | 637.8 to 635.6 | -0.3 / -0.3 | 642.6 to 640.4 | -0.3 / -0.3 |
| warlock-destruction | 636.6 to 634.6 | -0.3 / -0.3 | 638.9 to 636.9 | -0.3 / -0.3 |
| warrior-arms | 708.9 to 708.9 | +0.0 / +0.0 | 720.8 to 720.8 | +0.0 / +0.0 |
| warrior-fury | 753.9 to 753.9 | +0.0 / +0.0 | 757.8 to 757.8 | +0.0 / +0.0 |
| warrior-protection | 1062.9 to 1079.0 | +1.5 / +1.5 | 1051.4 to 1067.5 | +1.5 / +1.7 |

Where the moves come from:
- **Bear +19 to +20%, cat +13%:** Forever druids keep potions, Thistle Tea, Windfury Totem and weapon enchants in forms (Blizzard's druid deep dive, forever-bugs #109); the bear now drinks Mighty Rage Potion and carries an Elemental Sharpening Stone. How much a stone adds to a form's swing is open (below).
- **Retribution +10.5%:** the twist's wait (about +9.4%), the paced Consecration, less Conviction on Consecration and Touch of the Grave's misses.
- **Frost +5%:** Arcane Missiles no longer spend Fingers of Frost (+3.4%), plus a Spellblasting potion.
- **Restoration druid +6.5%:** Swiftmend keeps the HoT and is cast under 75% tank health; Gift of Nature off Tranquility takes about 1% back.
- **Protection paladin +2.0 to +2.5%:** Brilliant Wizard Oil and a Frenzy potion, less Conviction off Holy damage and Touch of the Grave. **Protection warrior +1.5%:** Elemental Sharpening Stone.
- **Rogues, Survival, Enhancement, Balance, Shadow about +1%:** Forever's 30-second potions (Survival also gains Lacerate's talents).
- **Warlocks -0.3 to -0.4%:** demon melee crit. **Restoration shaman -0.5%:** Healing Stream's crit. **Fire -0.1%:** Combustion off Ignite.
- Holy paladin moves only once the libram is equipped (the ranker re-measures it at the refresh).

Below 60 the healers move most: each band now fights its own content. Effective HPS falls by up to half at bands 30 and 40 because the fight is easier, not because the healer is weaker; healprofile.md has each table.

## wow-player review

Verdicts on the lane reports, the before/after table and the shots:
- **SHIP:** warlock, casters, physical, healers, ranks, items.
- **SHIP WITH FIXES:** crit, hybrids, auras. Each fix was a note, settled as follows:
  - The bear and cat gains rest on one Blizzard post for the rules and on the engine for the size; how a stone adds to a form's swing is open, and the release note should name the consumables behind the bear's number.
  - Conviction (off Holy damage, kept on heals) and Thundering Strikes (reaching shaman spells) carry the same client restriction; both follow the beta logs, which disagree between the two talents. Recorded open, "per the beta logs".
  - The stone no longer reaching the Survival pet moves no published number (Survival takes Lone Wolf).
- **HOLD, then SHIP:**
  - **platform:** the weapon-imbue panel listed 96 rows with "Main hand stone: Deadly Poison" and truncated names; rebuilt per class (S 414b0d9b5). The four potions were absent only because the local regen predated the auras merge; re-shot. The Survival line drops "than an Era guide suggests" (S 135d6bfb6).
  - **healprofile:** kept as the owner ruled. The reviewer's concerns are recorded for the owner: the sub-60 boss data is cmangos 1.12, the x1.2 dungeon modifier rests on one client table, bands 30 and 40 come out so easy that bare healers cover 86 to 95% (gear picks will tie), and the sub-60 numbers fall in a way players may read as weaker healers. None of it is published until the held refresh.
- **Numbers players will question first**, with the release-note line:
  - Bear and cat: druids keep potions, Thistle Tea, Windfury Totem and weapon stones in forms.
  - Retribution: the Unsealed Wrath twist waits for the Righteousness swing; Consecration is paced to the fight.
  - Frost: Arcane Missiles no longer eat Fingers of Frost, so Ice Lance gets the charge.
  - Restoration druid: Swiftmend leaves the Rejuvenation or Regrowth running.
  - Healers below 60: each band fights a fight at its own level; a lower number means an easier fight.

## Still open

- **Owner decisions:**
  - The sub-60 heal profile: add a pull-size or band multiplier (bare coverage 72 to 98% below 60 against 29 to 40% at 60; healprofile.md has the sweep), and whether the BiS page should show sub-60 healers as coverage rather than HPS.
  - Conviction's crit on paladin heals: removing it would cut Holy paladin's published healing by 5.6%; one Retribution paladin's heals in the logs.
  - Retribution bare 20 loses 0.3 to 0.5% to the paced Consecration line against the old 65% gate.
- **Blocked on level-60 or launch logs:** a stone's damage on a form's swing; Holy Power's crit-heal size; Oath of Light; Holy Shield crit; Master of Elements on a Clearcasting crit; Lightning Overload's rates; demons taking the warlock's plain crit; Thornspike and poisons; normalized rage's off-hand factor and crit multiplier; Lacerating Strikes' refresh; the one-a-minute proc rate; Force of Will; Era's base-crit columns (the 70707 tables are empty; a log fit puts druid, priest and shaman within 1.5 points, paladin's spell base looks higher).
- **Held under the adoption rule:** Marksmanship, Affliction's Improved Bane of Agony (+0.72%), Destruction's Immolate line, Fire's Improved Scorch line, Discipline's Inspiration, Frost's +1.0% search winner (+0.3 to +0.5% at 10,000), Holy paladin's band-20 build (missed 1% at seed 1), a druid band-50 move the per-band mechanism cannot express.
- **Single source or open:** warlock demon sheets against Era's table; Concentration Aura; forever-bugs #419; Kings with Heart of the Lion; the new elixirs' stacking with Mongoose-style buffs; Rallying Cry and Slip'kik's Forever versions (world buffs are gated off); Blood Siphon's crit table.
- **Found late, for round 11 (not yet verified against the client):**
  - Blackhand's Breadth (13965) and Eye of the Beast (13968) give 2% crit on every table in our data; the ranker reads the client as 1% melee and ranged crit and 1% spell crit respectively, plus use effects. Neither is picked today.
  - Restoration druid band 20: the ranker measured 5 Furor points against a 5051 build at +5.9%, measured before the healprofile lane's band-20 build (Heart of the Wild, Reflection, Bloom, Nature's Swiftness first) landed; re-measure on the merged ladder.
- **The data run needs `--allow-shrink`:** 165 item ids are gone from wago's 70707 export (105 are item-level-30 copies of Molten Core gear, the rest recipes and consumables); confirm they are really gone before re-baselining.
- **Page and pipeline (at the refresh):** simbuffs.json (the four potions, the salmon foods, the panel's icons and imbue rows); loot.json's WSG ids; Data.lua and addon-data.json; the ranker re-ranks Frost after Fingers of Frost, the healer trinket and relic slots, the race picks of Undead Shadow, Retribution and Protection paladin (Touch of the Grave); then the brotli curl check on the first deploy.
