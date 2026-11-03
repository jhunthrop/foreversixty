# Round 14, lane build: client 1.60.1.70755 against 1.60.1.70753 with the 11-03 push

Read-only diff. Nothing was activated, fetched into data/builds, pinned or fixed. A separate lane is
activating 70755 (engine follow, build carry, guides), so every row below is **OPEN, pending the 70755
pin**, origin `new-data`.

## Method and what the build is

- wago.tools /api/builds lists `wowf` 1.60.1.70755, ingested 2026-11-03 18:16 UTC. Its product_config
  (0d21f2f2...) is the same as 70753's.
- I downloaded 42 tables for both builds with the pipeline's own `wago.download_table`, then layered each
  build's push with `hotfix_push.layer_hotfix_push` into a scratch root. My 70753 layer matches the committed
  `data/builds/1.60.1.70753/hotfix-push.json` table for table (same sha256 and counts; wago still lists
  6365 Spell entries for 70753, so 70753 has had no push since 14:14 UTC). Scripts are in
  `$R/build/scripts`, outputs in `$R/build/out` (`diff.json`, `detail.txt`, `spellconst_diff.json`,
  `talents.txt`).
- **The shipped DB2 files are identical** across the two builds (all 42 tables). Everything that changes
  comes from the server push: wago captured the 70755 push at 2026-11-03 20:06 UTC. Push ids 113483
  (rogue) and 113539 (warrior) were also sent to **70707** clients at 17:34 and 18:29 UTC. So these are
  server hotfixes, not baked client data, and they are probably live for every client build. The shaman
  pushes (113282, 113380) and the Binding Heal/Lava Burst removal (113007) show up only on 70755 and 70707.
- Live on servers? I found no Blizzard note about these changes. The search turned up only beta-era
  coverage
  ([Massively OP 10-17](https://massivelyop.com/2026/10/17/world-of-warcraft-forever-beta-patch-makes-sweeping-changes-to-racials-and-classes/),
  [kami-labs 10-30, unofficial](https://kami-labs.fr/en/wow-forever/patch-notes-wow-forever-30-octobre-gagnants-perdants/)).
  The wago evidence says the server is pushing them today (11-03), the day before launch.
- I regenerated spellconst for both layers with `pipeline.simconst.build_spell_constants` in scratch
  only. The 70753 output matches the committed spellconst for all 9 classes. The talent strings were
  decoded with `pipeline.normalize.traits.read_trait_trees` against both layouts.

Tables checked: SpellName, Spell, SpellEffect, SpellMisc, SpellCooldowns, SpellCategories,
SpellCastTimes, SpellDuration, SpellPower, SpellLevels, SpellAuraOptions, SpellClassOptions,
SpellShapeshift, SpellEquippedItems, SpellRadius, SpellRange, SkillLineAbility, SkillLine,
SkillLineXTraitTree, Talent, TalentTab, the 11 Trait* tables, ItemSparse, Item, ItemEffect,
ItemXItemEffect, ItemSet, ItemSetSpell, SpellItemEnchantment, ChrRaces, ChrClasses, CurvePoint,
RandPropPoints. No change in: SpellCastTimes, SpellDuration, SpellRange, Talent, TalentTab, ChrRaces,
ChrClasses, ItemEffect, ItemSet, ItemSetSpell, RandPropPoints. **No racial changed.**

## Shaman (Elemental, Enhancement, Restoration)

The engine ability files hold the 70753 numbers by hand (e.g. sim/shaman/lightning_bolt.go
`LightningBoltDamage` rank 10 = 196 +1.2/level), and the spellconst tests will fail on the move.

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| S1 | Lightning Bolt, all 10 ranks, base damage up: R10 15208 196 to 353 (+80%), per level 1.2 to 2.1; R9 178 to 295; R8 162 to 240. Elemental and Enhancement cast it | SpellEffect 691780/691075/688310 (push 113380) | OPEN, pending pin | new-data | - |
| S2 | Chain Lightning, all 4 ranks: R4 10605 base 123 to 510, coefficient .571 to .714, cast 2.0 to 2.5 s (CastingTimeIndex 5 to 19), mana 315 to 495, per level 0.8 to 3.3. Elemental and Enhancement APLs cast it | SpellEffect 688920, SpellMisc 322672, SpellPower 177727 | OPEN | new-data | - |
| S3 | Earth Shock: R7 10414 301 to 531, per level 1.9 to 3.4 (every rank up) | SpellEffect 688559 etc. | OPEN | new-data | - |
| S4 | Flame Shock: R6 29228 direct 166 to 292, tick 44 to 80 (every rank up) | SpellEffect 704945/704946 | OPEN | new-data | - |
| S5 | Frost Shock: R4 10473 283 to 506 (every rank up). Elemental APL | SpellEffect 688022 | OPEN | new-data | - |
| S6 | Stormstrike cooldown 8 to 6 s, plus a new effect 2 (Effect 31, weapon damage percent, 150) | SpellCooldowns 55241, SpellEffect 1368308. Enhancement | OPEN | new-data | - |
| S7 | Riptide (3 ranks) changes from a direct heal (Effect 10, 480/612/841) to Effect 3 (dummy, 0). Category cooldown 6 to 12 s, mana 245/300/385 to 150/180/230 | SpellEffect 1072707/1236793/1236796, SpellCooldowns 84834/94614/94615. Restoration APL casts Riptide | OPEN | new-data | - |
| S8 | Lava Burst is no longer learnable: all 7 SkillLineAbility rows are status 2 (removed) in push 113007. The Elemental talent node 104760 is now Call of Thunder (TraitDefinition 134221: 408490 to 16120). Call of Thunder 16120 effect 3 to 5. Elemental mastery node 104758 is renamed Lightning Mastery (16166). The Elemental APL casts Lava Burst. **This contradicts the activation lane's "Lava Burst trainable" note: please recheck** | SkillLineAbility 49537/49538/53175-53178/59659 status 2; TraitDefinition 134221 | OPEN | new-data | - |
| S9 | Elemental Reach 28999 effect 2: aura 107 misc 11 becomes aura 4 (dummy) | SpellEffect 1363729 | OPEN | new-data | - |
| S10 | Totemic Projection 437009 renamed Totemic Call, cooldown 60 to 30 s. Not cast by any APL | SpellCooldowns 88785 | OPEN (no sim effect expected) | new-data | - |
| S11 | Enhancement tree layout: Improved Lightning Shield and Improved Ghost Wolf swap cells, so the raid-60 string `...-253030031005102051-...` decodes to different talents on 70755 (Shamanistic Focus at 3 of 1, Stormstrike at 5 of 1). Every Enhancement string band 30-60 shifts | TraitNode 104751/104752 PosX/PosY; out/talents.txt | OPEN: the build carry must re-encode by name | new-data | - |

## Rogue (Assassination, Combat, Subtlety)

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| R1 | Mutilate (all ranks) weapon damage 33/48/67 to 44/63/88 and bonus 75 to 90. The 1310705-07 reissues are removed and 399956 is the talent again | SpellEffect 1044484/1240651/1240653..., Spell 1310705-07 status 2. Assassination | OPEN | new-data | - |
| R2 | Venom 1310703 reworked: no cost, no GCD, effect 0 becomes a 30% aura 108 (class mask 8192/8), effect 2 becomes 20% aura 107; the talent is now 1329328 (proc trigger). Assassination casts Venom | SpellEffect 1340416-18, SpellPower 314521, SpellCooldowns 101043 removed | OPEN | new-data | - |
| R3 | Poisons up about 40-50%: Instant VI 175 to 250, Deadly IV 160 to 240, Deadly V 180 to 240, Wound 160 to 240, Mind-numbing 150 to 225; enchant charges match (SpellItemEnchantment 625/627/643/706/2630) | SpellEffect 1340483..., SpellItemEnchantment Charges. Every rogue spec | OPEN | new-data | - |
| R4 | Improved Eviscerate is renamed Coup de Grace and gains effect 1 (aura 108, 20%, mask word 2 = 4). Improved Sinister Strike is renamed Efficient Strikes, mask 2 to 33554434 | SpellEffect 1369147, 691146 | OPEN | new-data | - |
| R5 | Flawless Execution 1310711 -10 to -15. Hemorrhage 16511 mask 1048576 to 1048832. Aggression 18427 mask 131078 to 131074. Ghostly Strike 14278 gains a trigger to 1329326 (new buff) | SpellEffect 1340441, 693309, 695769, 1368766 | OPEN | new-data | - |
| R6 | Tree reworked: Murder, Improved Distract and Thousand Cuts gone; Vigor/Blood Splatter, Composure, Relentless Strikes (now 5 ranks), Dual Wield Spec (now 2 ranks), Flash of Steel, Boot Knife, Quick Getaway and Death Mark added. Max ranks: Seal Fate 5 to 3, Lethality 5 to 2, Opportunity 2 to 5, Camouflage 5 to 3. All 17 published strings decode to illegal builds on 70755 (e.g. Assassination raid: Lethality 3 of 2, Seal Fate 5 of 3) | Trait* diff, out/talents.txt | OPEN: activation lane owns it | new-data | - |
| R7 | New baseline abilities: Shank 1329283/1329549 (Combat skill line, levels 24/44, 10 s cooldown, 50% weapon damage) and Death Mark 1316709-1325758 (Subtlety line, 4 ranks). No APL casts them yet | SkillLineAbility 59899-59904 | OPEN: rotation question once live | new-data | - |

## Warrior (Arms, Fury, Protection)

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| W1 | Slam (all ranks) category cooldown 18 to 15 s. Arms and Protection APLs cast Slam | SpellCooldowns 53310...101002 | OPEN | new-data | - |
| W2 | Improved Slam 12862 effect 2 -3000 to -5000 ms (2 ranks, Arms takes 2) | SpellEffect 1339521 | OPEN | new-data | - |
| W3 | Raging Blows 1310315 becomes Improved Whirlwind (2 ranks): -2 s Whirlwind cooldown, +30% damage | SpellEffect 1339744/1345686. Fury APL casts Whirlwind | OPEN | new-data | - |
| W4 | Reckless Swings 1329555 is new (Fury talent plus SkillLineAbility 256): 15 s cooldown, 25% rage cost, next 4 swings +rage | added rows | OPEN | new-data | - |
| W5 | Max ranks: Improved Tactical Mastery 5 to 3, Bloodthrill 5 to 2, Flurry 5 to 3. Arms raid takes 5/5 and 5/5, which is illegal on 70755; Fury strings shift | TraitNodeEntry 130684/137214/130658 | OPEN: activation lane | new-data | - |
| W6 | Booming Voice 12321 mask word 1 0 to 32; Sweeping Strikes and Concussion Blow Attributes_14 bit 64; Spearing Strike 1310222 effect 1 40 to 75, effect 2 2 to 1 | SpellEffect 1360666, 1339573/1339768 | OPEN | new-data | - |

## Other classes

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | Warlock Wrack 1316697: coefficient 0.143 to 0.286, per level 0 to 2.5. sim/warlock/wrack.go; Affliction | SpellEffect 1349556 | OPEN | new-data | - |
| X2 | Paladin Vengeance 20050 stacks 3 to 5 (sim/paladin/talents.go); Retribution | SpellAuraOptions 247421 | OPEN | new-data | - |
| X3 | Paladin Two-Handed Weapon Specialization 20111 effect misc 1 to 3 (an item subclass mask: adds a second weapon class); Retribution, Protection | SpellEffect 697263 | OPEN: check engine gating | new-data | - |
| X4 | Priest Binding Heal: all 6 SkillLineAbility rows removed (status 2, push 113007). The Discipline and Holy APLs reference Binding Heal | SLA 50041, 53923-53927 | OPEN | new-data | - |
| X5 | Hunter Sniper Shot 1310786 SkillLineAbility AcquireMethod 2 to 0 (becomes a normal trainable) | SLA 59257 | OPEN: trainables | new-data | - |
| X6 | Items: no ItemSparse change on a published BiS item. The item pair layers by `hotfix_merge`'s rule, not mine, so 45 BiS ids in "added" and 6 in "withheld" (13048, 3205, 3748, 5191, 5202, 6392) are layering artefacts to recheck after the real fetch | out/diff.json ItemSparse | OPEN | new-data | - |
| X7 | Druid, mage, hunter and priest spell values: no change in spellconst. Mangle! 57596 and Righteous Fire 48198 are NPC/event spells | spellconst_diff.json | Aligned | - | - |

## Before/after

Not measured: no 70755 pin exists yet, and the activation lane owns it.

## Files to regenerate (activation lane)

The 70755 spellconst, trainables, talents/trees, spellranks, simdb; the engine's constants_auto_gen,
talents_auto_gen, proto and periodic-crit/eureka gens for shaman, rogue, warrior and warlock; the
hand-typed ability tables (lightning_bolt.go, chain_lightning.go, earth/flame/frost_shock.go,
stormstrike.go, mutilate.go, venom.go, poisons.go, slam.go, wrack.go); the published talent strings for
rogue, warrior, Elemental and Enhancement; ladder goldens.

## Measure at refresh

Every shaman, rogue and warrior spec, Affliction and Retribution, once 70755 is pinned.

## For other lanes

- Activation lane: S8 (Lava Burst SkillLineAbility rows removed, node is Call of Thunder) and X4
  (Binding Heal removed) disagree with or are missing from your summary.
- Rogue/warrior/shaman lanes: R7 (Shank, Death Mark) and W4 (Reckless Swings) need rotation decisions.

## Still open

- Whether servers require 70755 at launch. The rogue and warrior pushes already reach 70707 clients.
- The item layer needs a real `pipeline fetch` and hotfix_merge.
