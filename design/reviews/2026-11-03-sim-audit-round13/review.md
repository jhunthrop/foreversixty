# Round 13: wow-player review fixes (lane review)

Site wt/r13-review (from int/r13), fork lane/r13-review (from int/r13). Every guide edit passes `web/src/content/guides/_prose.test.ts`. No nightly numbers in prose.

| # | reviewer ask | what changed | evidence | sha |
|---|---|---|---|---|
| 1 | Elemental leveling paragraph claims Tidal Focus and Mindfulness for the whole climb | The paragraph now says every point goes to Elemental until 50; from 50 Tidal Focus (5% hit) and all three Mindfulness points (half your regeneration while casting); at 60 Natural Grace, Totemic Focus and Improved Reincarnation. | The leveling walk (`leveling.LadderTalentString` on the guide build): 20 to 40 take no Restoration point, 50 `-0530000000000000`, 60 `-0533520000000000`. Mindfulness 104734 rank 3: "Allows 50% of your Mana regeneration to continue while casting." Natural Grace, Totemic Focus and Improved Reincarnation rank texts in talents/shaman.json. | S dc86a6e2e |
| 2 | Resto druid "Forever keeps it whole on low ranks" | Now: rank 1 Rejuvenation carries the same bonus healing as the top rank in Forever, so a geared druid gets a large share of the top rank's heal for a small fraction of the mana. The tank line says why a middle rank while leveling: healing power is low, so most of the heal is its base, and rank 1 barely dents a tank's damage. I did not use the reviewer's "nearly as much healing": at 600 healing power rank 1 is about 128 a tick against about 281 for rank 10. | SpellEffect 774, 1058, 9841 and 25299: aura 8 at 0.2 bonus coefficient on every rank (bases 8 to 194). The engine applies no level penalty to heals (`sim/core/spell.go` `levelPenalty`: `ProcMaskSpellHealing` returns 1; round 11's beta logs). | S dc86a6e2e |
| 3 | Affliction "for three things, and puts one point in Improved Bane of Agony" reads as four | Split in two: Nightfall, Soul Siphon, Wrack and Soul Harvest buy three talents. Then a separate sentence: the Improved Drains point moves to Improved Bane of Agony for 5% more Bane of Agony damage. | Improved Bane of Agony rank 1: "by 5%" (talents/warlock.json). | S dc86a6e2e |
| 4 | Mage "share one 2 minute cooldown" needs "in Forever" | All three guides now read "In Forever, mana gems and Demonic Runes share one 2 minute cooldown". | The gems and the runes are both SpellCategory 1153 at 120000 ms (auras.md C2). | S dc86a6e2e |
| 5 | Enhancement stat paragraph ends on Strength | Order is now hit, Intellect, crit, Agility, Strength. That is the raid-60 scale factors' order (haste is not on items). No numbers. | bis/shaman-enhancement.json raid 60 (both factions): melee_haste > hit > intellect > crit > agility > strength > attack_power > spell_power. | S dc86a6e2e |
| 6 | Shadow: say the sim keeps Vampiric Embrace up and counts damage only | One line added to the Vampiric Embrace paragraph. | `data/curated/apl/priest-shadow.json` casts 15286 whenever it is down, and its note says the damage figure does not count the healing. | S dc86a6e2e |
| 7 | Read Leyline: the sim does not cast it on its own | Added to the Balance and Restoration race paragraphs ("a druid healer can press every fight" replaces "will use every fight") and to the curated race note in races.json. | auras.md R1 (E 2fde66add): the spell is registered and castable from an APL, never autocast. No curated APL casts 1259705. | S dc86a6e2e |
| 8a | BM Aimed Shot on the shared cooldown is within error | BM guide: Multi-Shot on cooldown is fine, and Aimed Shot is a small extra right after an auto shot fires. The BM APL line note says the same. | Reviewer measurement; the APL itself is unchanged. | S dc86a6e2e |
| 8b | MM: keep the instruction, call it a small gain | MM guide and APL note say "a small gain over Multi-Shot when the cast fits". | - | S dc86a6e2e |
| 8c | "more than about a second away" timing | Both guides and both APL line notes now say "right after an auto shot fires" or "partway to your next auto shot". The 1.2 s gate in the APL is unchanged. The top-level history notes keep the engineering wording; they are not shown to players. | grep of the hunter guides and APL notes: no other occurrence. | S dc86a6e2e |
| 8d | MM raid rotation: dismiss your pet before the pull | Added as the first sentence of the first shown MM step note (Hunter's Mark). Prepull actions carry no rendered notes (`web/scripts/rotation-apl.mjs` reads only priorityList notes). Notes are curated-only: `make apl-check` compares parsed JSON with notes stripped, so no fork sync is needed. | Lone Wolf (MM tier 2): "while you do not have an active pet". | S dc86a6e2e |
| 8e | Can BM's raid tree reach Lone Wolf? | **It can.** Lone Wolf is Marksmanship tier 2 (needs 10 points in the tree). BM's raid build `5420001505001251/00523500014/0` has 20 Marksmanship points (31/20/0), and so does its leveling build `00524510012`. The build doesn't take Lone Wolf because BM's damage comes through the pet. Nothing to change. The published BiS raid band still shows the leveling Marksmanship string (`0052451001200000`); that is nightly data, so it is listed under "Measure at refresh". | talents/hunter.json tiers; guide frontmatter; bis/hunter-beast-mastery.json. | - |
| 8f | "Hawk 120" or rank-7 Hawk claims in hunter prose | None found (I grepped the guides, curated JSON and request Go for Hawk with 110/120 or rank 6/7). The APL history notes name rank 6 (14322). | grep | - |
| 9 | Fury Raging Blows: does it cheapen Cleave too? | Yes. The bullet now reads "Whirlwind and Cleave cost 3 less rage". The engine already discounts both (`sim/warrior/talents.go` `ragingBlowsRageDiscount`). | 1310315 effect 2: aura 107 (add flat modifier), misc 14 (cost), -30, mask [4194304, 4, 0, 0]. 4194304 is Cleave's SpellClassMask_0 (845, 11609, 20569) and mask_1 4 is Whirlwind 1680's. The rank text: "Reduces the Rage cost of your Cleave and Whirlwind abilities by 3." | S dc86a6e2e |
| 10 | Arms raid Charge at -2 s, before the tank | The APL cannot tell the raid preset from bare (no APL value names it). Charge also cannot be cast in combat (engine `ExtraCastCondition` is `CurrentTime <= 0`, which matches the client). So the simplest correct form is Charge at 0 s: in a raid that is the tank's pull, and in the bare bands it is the warrior's own pull. The pre-pull Battle Shout moves to 0 s after it and is still paid by the Charge's rage; it is still skipped under a raid shout. Bloodrage stays in the priority list. arms.md now says to Charge on the pull: in a raid once the tank has engaged, never before; out in the world Charge is how you pull. The fork copy was synced by hand. | `TestArmsChargesOnThePullNotBefore` (sim/request/warrior_opener_test.go) failed on the old opener, then passed. The other warrior opener and shout tests pass. Charge is still cast once per fight. | S 28d257996, E 5dc93c822 |

## Arms before/after (10000 iterations, published entries, the -rotation old vs new curated file)

| preset | faction | seed | before | after | change |
|---|---|---|---|---|---|
| raid 60 | horde | 1 | 731.2 | 729.7 | -0.21% |
| raid 60 | horde | 50001 | 731.8 | 730.3 | -0.20% |
| raid 60 | alliance | 1 | 719.0 | 717.5 | -0.21% |
| raid 60 | alliance | 50001 | 719.6 | 718.1 | -0.21% |
| bare 60 | horde | 1 | 249.6 | 249.1 | -0.20% |
| bare 60 | horde | 50001 | 249.8 | 249.4 | -0.16% |
| bare 60 | alliance | 1 | 251.3 | 250.7 | -0.24% |
| bare 60 | alliance | 50001 | 251.3 | 250.7 | -0.24% |

The raid error is about ±0.5 DPS. The loss is the two seconds of pre-pull rage the old opener got by Charging early. It is the price of the correct opener.

## Files to regenerate
- sim/request/apl/warrior-arms.apl.json, hunter-beast-mastery.apl.json, hunter-marksmanship.apl.json (`make apl-sync`; notes and the Arms opener). Not committed here. The druid, hunter and elemental copies in this worktree are also stale from earlier lanes.
- sim/request ladder goldens for warrior-arms (TestRotationLadder): the opener moves every band.
- Arms BiS/weights entries and dps-per-point (the raid and bare numbers move about 0.2%).
- data/builds/1.60.1.70753/races.json (the Read Leyline race note).
- web rotation notes (web/scripts/sync-rotations.mjs output) for hunter-beast-mastery and hunter-marksmanship.

## Measure at refresh
- hunter-beast-mastery: the published raid-60 band carries the leveling Marksmanship string `0052451001200000`, while the guide's raidBuild is `00523500014`. Check this after the next BiS run.

## Still open
- Restoration druid: check Rejuvenation rank 1's tooltip in game against a top-rank one at the same healing power (does the client show the full bonus?).
- Protection paladin band 20: not touched (owner call).

## Not run
- `web/src/lib/sim/rotations.test.ts` cannot load in this worktree (it needs `web/src/data/generated/combos.json`, which is not generated here). I checked the changed notes with `rotationNotesOf` from web/scripts/rotation-apl.mjs, which runs the engineering-leak guard; they pass.
- `data/tests/test_curated.py::test_every_class_and_race_documents_a_forever_change` fails on the clean base too: it reads a missing builds/1.60.1.69893/classes.json.
