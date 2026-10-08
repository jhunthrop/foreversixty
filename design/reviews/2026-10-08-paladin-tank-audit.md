# Protection paladin tank audit, 2026-10-08

Branch `pal-tank` in both repositories (site and engine fork). Nothing here was
pushed or merged. Scratch ranker runs only; no `bis/*.json`, `loot.json`,
`Data.lua` or `simdb.bin` is committed.

## The question

Against the same curated boss (2.0 s swings, 2,500 minimum before armor, parry
haste, 440 healing per second), the raid-ready level 60 Protection warrior
barely takes damage and the Protection paladin died in about one fight in
seven (the 2026-10-07 review printed 23%, and 90% bare). Is the cause the
paladin model, its gear pool, its rotation or the encounter?

Answer: the paladin's **build** and a missing **aura**, both of ours; not the
encounter, not the rotation, and the gear pool is complete. The reference build
spent the 20 points past Holy Shield on the Holy tree's healing talents, and the
kit gave the paladin none of the armor a Protection paladin brings itself. With
both fixed the same boss kills the raid-ready paladin 0.1% of fights (the
warrior: 0%) and the bare paladin 31% (the warrior: 1%). What is left is a real
class difference the client and the engine's class tables carry (health and
block), measured below.

## Before and after (ranker, level 60, Alliance, 3,000 iterations, scratch `-out`)

Baseline is branch `main` at `fa5274ec` with the fork at `921b7f8a4`. The
2026-10-07 review's figures (warrior 29.9k effective health, paladin 20.9k) were
taken before later data changes; this table is the same code run today.

| Spec | Preset | | Effective health | Damage taken /s | Chance of death | Threat /s | TMI |
|---|---|---|---|---|---|---|---|
| Protection warrior | bare | before and after | 22,758 | 360 | 1.0% | 528 | 54.8 |
| Protection warrior | raid | before and after | 35,336 | 313 | 0.0% | 726 | 33.7 |
| Protection paladin | bare | before | 16,613 | 452 | 90.0% | 278 | 80.3 |
| Protection paladin | bare | after | 19,090 | 403 | 30.7% | 290 | 67.9 |
| Protection paladin | raid | before | 23,884 | 411 | 15.3% | 604 | 52.9 |
| Protection paladin | raid | after | 28,961 | 333 | 0.1% | 611 | 44.5 |

The warrior row is unchanged because nothing in its code moved (checked by
re-running it on the final engine). Feral bear on the same engine, for the
encounter question below: raid 3.4%, bare 13.9%.

## 1. Where the gap came from

`spec-breakdown` gained `-tank` (and `-gear slot:id,...`): it builds the request
the ranker builds and prints the character's stats before the fight's auras,
the tank figures, the boss's swing outcomes against the character, absorbs and
aura uptimes. Raid-ready published sets, 10,000 iterations, dwarf:

| | Warrior | Paladin before | Paladin after |
|---|---|---|---|
| Maximum health | 7,589 | 6,741 | 6,619 |
| Stamina | 476 | 422 | 410 |
| Armor from items and buffs (before auras) | 7,819 | 8,103 | 7,409 |
| Armor aura in the fight | none | none | Devotion Aura +735 |
| Defense skill | 80 | 99 | 106 |
| Dodge / parry / block chance (before Shield Block or Holy Shield) | 22.7 / 13.2 / 15.2 | 19.8 / 9.0 / 14.4 | 24.7 / 14.2 / 14.6 |
| Block value (0.05 a strength, less 1; items add none, see 2) | 17 | 18 | 16 |
| Boss swings that miss / dodge / parry | 7.8 / 22.3 / 12.8 | 8.3 / 19.1 / 8.4 | 8.6 / 24.1 / 13.7 |
| ...are blocked | 52.9 | 47.4 | 45.0 |
| ...crit / crush / plain hit | 0.2 / 1.5 / 2.5 | 1.3 / 11.4 / 4.0 | 1.1 / 6.8 / 0.8 |
| Damage taken /s | 313 | 411 | 333 |

Reading it. A blocked swing takes its full damage less the block value (about
17 here), so blocking mitigates almost nothing by itself; what block chance
buys is room in the attack table: a swing that is blocked cannot crit or crush,
and a crush is 150% damage. The warrior's Shield Block (+75% block, two charges,
7 s, 84% uptime) fills its table until only 4% of swings are neither
avoided nor blocked; the old paladin's Holy Shield (+30%, 94% uptime) left 17%
of its swings in the table, 11.4% of them crushes. That, plus 12% less health,
plus 4.4 points less parry and 3.2 less dodge, was the 31% more damage taken.
Deflection and the re-ranked avoidance gear (24.7% dodge, 14.2% parry) took the
paladin's unblocked remainder down to 9%, 6.8% of it crushes.

Absorbs and cooldowns over a 180 s fight (the standard boss rarely drops either
tank low enough to use them): warrior Shield Wall 0.03 casts, Last Stand 0.02;
paladin Templar's Bulwark 0.28 casts (3.6 damage a second absorbed), Divine
Shield 0.01, Seal of Fury's shield 6.8 a second, Holy Shield 18.2 casts. Both
sides have emergency tools the boss does not make them use.

Gear. The ranker's pick per slot (Alliance, raid-ready, band 60):

| Slot | Warrior | Paladin before | Paladin after |
|---|---|---|---|
| head | Faceguard of Heroism | Enchanted Thorium Helm | Helm of Awareness |
| neck | Rage of Mugamba | Amulet of the Darkmoon | Evil Eye Pendant |
| shoulder | Highlander's Plate Spaulders | Highlander's Plate Spaulders | Glowing Mantle of the Dawn |
| chest | Chestguard of Heroism | Deathbone Chestplate | Deathbone Chestplate |
| wrist | Wristguards of Heroism | Sentinel's Wristguards | Sentinel's Wristguards |
| hands | Handguards of Heroism | Heavy Thorium Gauntlets | Heavy Thorium Gauntlets |
| waist | Waistguard of Heroism | Sentinel's Waistguard | Sentinel's Waistguard |
| legs | Sentinel's Plate Legguards | Soulforge Legguards | Soulforge Legguards |
| trinkets | Vigilance Charm, Mark of Tyranny | Mark of Tyranny, Talisman of Arathor | Vigilance Charm, Stormpike Insignia Rank 6 |
| weapon, shield | Quel'Serrar, Earthen Guard | Ravencrest's Legacy, The Immovable Object | same |
| ranged | Precisely Calibrated Boomstick | none | none |

The pool is not missing pieces. The Mokvar vendor sells a tank variant for each
class (the "of Heroism" warrior set, the "Soulforge" paladin set: Faceguard,
Chestguards, Handguards, Pauldrons, Waistguard, Wristguards, Sabatons,
Legguards), and the paladin one trades the warrior's defense for intellect, so
the paladin takes only the legs from it and the ranker prefers crafted and
rep pieces with more defense and avoidance. The paladin's ranged slot is a
libram and the item tables carry none with a tank stat. The ranker never picks a
spell-power plate piece for the paladin. Where the paladin trails on gear
(about 66 stamina, 660 health) is the class sets, not the pool.

What moved the number, one change at a time on the same set (6,000 to 10,000
iterations, the spec-breakdown tank report):

| Change | Damage taken /s | Chance of death |
|---|---|---|
| Published set, old build | 411 | 14.8% |
| Deflection 5 instead of 5 Holy points | 388 | 5.8% |
| + Sacred Duty 2 | 389 | 5.3% |
| + Iron Creed 5 | 373 | 1.4% |
| + Devotion Aura | 358 | 0.3% |
| Re-ranked gear for the new build and aura | 333 | 0.1% |

The Holy points were the build's worst spend: Healing Light, Reverence and
Spiritual Focus (11 of the 20) do nothing for a tank. Deflection is the only
parry a paladin can take and lives in the Retribution tree.

## 2. Client parity

Sources: `data/builds/1.60.1.70009/raw` (Spell, SpellEffect, SpellAuraOptions,
SpellCooldowns, SpellMisc, SpellPower), the live trees in `talents/paladin.json`
and the vendored `spellconst`. The talent text is Wowhead's live overlay and in
three places it is newer than the wago tables (marked "hotfix").

| Ability or talent | Client | Engine | Verdict |
|---|---|---|---|
| Holy Shield (20925, 20927, 20928) | rank spells: block +20, 4 charges, 10 s, 10 s category cooldown, 150/195/240 mana, levels 40/50/60, damage 110/153/221 at 0.08 spell power, +20% threat; live talent text: block +30 | 30 block, 4 charges, 10 s, same damage, cost, coefficient, 1.2 threat; needs a shield | matches the live text (hotfix, 20 to 30); at 20 the tank takes 3.7% more damage and dies 0.5% of fights instead of 0.1%. Unverified until a build carries it |
| Righteous Fury (25780) | +60% threat on Holy only, no other effect (the damage-taken effect is Improved Righteous Fury's), 30 minutes | +60% on Holy-school spells as a static mod, permanent | matches; white hits unraised, as the client states |
| Improved Righteous Fury (20468) | damage taken -2% a rank, -6% at 3 | same | matches |
| Seal of Fury, ranks and Judgement of Fury | conformance golden | `declared, matches` at every rank; absorb half the damage with a shield | matches; the 10 s absorb life and the Judgement taunt are unconfirmed or unmodeled (single boss) |
| Hammer of the Righteous (407632) | text: "Holy damage to each equal to $s3 times the damage per second of your main hand weapon", third effect 3, a flat 1 and a dummy 120; 6 s, 6% of base mana, up to 3 targets | was 120% of one swing | **fixed**: 3 times weapon damage per second (3 / speed of a swing), whether attack power counts as the character sheet does is unconfirmed; the 120 dummy is not read |
| Holy Strike (10333) | 93 flat plus normalized weapon at 50%, 10 s, 20 mana | conformance `match`, damage `declared, matches` | matches |
| Consecration (20924) | 8 s; damage on a companion row per rank (1280345 to 1280349): 12 a second to all and 27 more (0.095 spell power) to the first 4 enemies at rank 5 | Classic's 48 a second, no coefficient | **fixed**: a single target takes 39 a second plus 0.095 spell power, 312 over the 8 s instead of 384; pinned by `TestConsecrationDamageMatchesClient` |
| Blessing of Sanctuary | not in this client (no spell by the name) | not modeled; `blessing_of_sanctuary` is still a request id and a no-op | matches; stale id noted below |
| Devotion Aura (465 to 10293) | 55, 160, 275, 390, 505, 620, 735 armor at levels 1, 10, 20, 30, 40, 50, 60 | always 735 | **fixed**: the seven ranks (`core.DevotionAuraRanks`, pinned against the client by the site); and the Protection kit now keeps the aura on, see 3 |
| Redoubt | 10% to proc on a damaging melee hit, +4% block a rank for 10 s or 5 blocks | same | matches |
| Toughness | armor from items +2% a rank | `ApplyEquipScaling` | matches |
| Shield Specialization (20150) | shield absorb +10% a rank (aura 272), blocks restore 6% of mana at 33/66/100%, 3 s ICD | block value multiplier +0.1 a rank, same mana rule | matches the text; the multiplier scales item block value only (strength's share is not scaled), so with no item block value (see below) it does nothing for a tank |
| Anticipation | defense +4 a rank | same | matches |
| Reckoning (20177) | 8% a rank to extra attack on block, 20% a rank (2.5 times) on a taken crit | same | matches |
| Templar's Bulwark (1311015) | absorb of 100% of maximum health, 8 s, 5 min cooldown, 110 mana, Forbearance 1 min | same | matches |
| Iron Creed (1311034) | live text: Holy Strike threat +5% a rank, damage taken -2% a rank for 6 s with Righteous Fury; wago tables read 30% and 15% at rank 5 | text numbers (25%, 10%) | matches the live text (hotfix, nerf from 30 and 15) |
| Sacred Duty | stamina +2% a rank, defensive cooldowns -30 s a rank | same | matches |
| Swift Judgement, Improved Seal of Fury | script / "restore 0 mana" (unresolved in the tooltip) | the engine assumes 1% of maximum mana, scaled 15% a level up to 45% | Improved Seal of Fury's mana is unconfirmed, as before |
| One-Handed Weapon Specialization | 3 / 7 / 10% | same | matches |
| Deflection (Retribution) | parry +1% a rank | same | matches |
| Benediction (Retribution) | instant spells' mana -2% a rank, 10% at 5 | -3% a rank, 15% | **mismatch, not fixed** (Retribution's ladder and goldens move; no tank gain either way, measured) |
| Guardian's Favor, Improved Hammer of Justice, Unyielding Faith (Holy, fear and disorient -30%) | utility | not modeled | the boss never stuns, fears or needs a Blessing of Protection; nothing a tank ranking can price |

Block value from gear. `items.json` lists The Immovable Object with `block: 27`.
That is `ITEM_MOD_BLOCK_VALUE` (StatModifier 48, 27 flat), not a block rating.
`pipeline/normalize/gear.py` maps both modifier 15 (rating) and 48 (value) to the
one planner stat `block`, and the simdb conversion then divides it by 5 ratings
per percent, so the engine applies **+5.4% block chance and no block value**.
The classic-db supplement drops shield block value too (aura 158 is on its
ignored list; Earthen Guard's "Block Value 12" is lost). Priced by swapping
shields on the published paladin set: the Immovable Object's phantom 5.4% is
worth about 1% of damage taken (333 against 336 for Earthen Guard); flat block
value 27 (35 with Shield Specialization) is of the same size, so the pick and
the sign of the error are both small. I did not fix it: a clean fix adds
`block_value` to the planner's `STAT_KEYS`/`STAT_LABELS`
(`web/src/lib/planner/types.ts`, and `test_build_stat_keys.py`), maps modifier
48 and aura 158 to it, and regenerates every class's items and simdb; that is
the data lane's change, with a regen, not an edit made inside a tank audit. The
shield's own block value is absent from this client's item rows altogether (no
`ItemSparse` field), so the engine's block value is only strength / 20 less 1.

Forever-new protection talents, unmodeled or only approximated: Guardian's Favor
(Blessing of Protection cooldown, Blessing of Freedom duration) and Improved
Hammer of Justice (stun cooldown) are not modeled and cannot matter against a
boss that is never stunned; Unyielding Faith likewise. Seal of Fury's Judgement
taunt (10 yards, 4 s) and its absorb's life are unmodeled and assumed. Improved
Seal of Fury's mana is assumed. Everything else in the tree is modeled.

## 3. Changes

Engine fork, branch `pal-tank` (`7bde0a766` and the comment edit after it):

- `ForeverProtectionTalents` is now the guide's build: Deflection 5, Sacred
  Duty 2, Iron Creed 5, Divine Strength 5, Improved Seals 3 (protection 38,
  Retribution 5, Holy 8). New tests: legality against the client's tree, and
  that the points past Holy Shield are spent on mitigation. `TestProtection`
  regenerated.
- Devotion Aura follows the client's seven ranks (`DevotionAuraRanks`,
  `DevotionAuraArmor`, a level-table test); the fixed 735 and its unused enum
  entry are gone.
- Hammer of the Righteous is 3 times the weapon's damage per second.
- Consecration is the client's two-part damage, first four enemies heavier.

Site, branch `pal-tank`:

- The paladin-protection guide build is `FS1:...:50003/5530513321301051/5:` and
  its Talents section says why; the ranker, the ladder and the planner read it.
- The class kit gives `paladin-protection` Devotion Aura from level 1
  (`sim/leveling/kit.go`; the other paladin specs are unchanged), with a test
  and a client rank pin (`TestDevotionAuraRanksMatchTheClient`).
- The paladin-protection ladder golden is regenerated (it changes because the
  build and kit changed).
- `spec-breakdown -tank` and `-gear`, with tests.
- The rotation notes are rewritten around what was measured (below).

## 4. Rotation

`rotation-search` scores damage only and skips tank specs (`setup.go`: "DPS is
not its objective"), so the comparison is by hand on the tank fight with
`spec-breakdown -tank -rotation` on the final published set, 10,000 iterations,
damage taken per second / threat per second:

| Variant | Damage taken | Threat | Chance of death |
|---|---|---|---|
| Curated | 333.0 | 611 | 0.1% |
| Holy Strike before Holy Shield | 333.0 | 611 | 0.1% |
| Hammer before Holy Strike | 333.5 | 609 | 0.1% |
| Judgement before Hammer | 332.7 | 611 | 0.1% |
| Consecration floor 30% / 80% mana | 333.4 / 333.1 | 615 / 598 | 0.1% |
| no Consecration | 332.2 | 537 | 0.1% |
| no Hammer | 329.3 | 504 | 0.1% |
| no Judgement | 333.0 | 567 | 0.2% |
| no Holy Strike | 347.1 | 570 | 0.8% |
| no Holy Shield | 353.7 | 482 | 1.6% |
| no Seal of Fury recast | 337.8 | 529 | 0.2% |

The list is right: Righteous Fury is on all fight (uptime 99.9%, a request
option, not a line), Holy Shield is up 94% of the time on its 10 s cooldown, Seal
of Fury is recast every 30 s, Holy Strike and Hammer sit ahead of Judgement and
Consecration, and nothing was a better order within noise. With Iron Creed in
the build Holy Strike is now mitigation as well as threat (58% uptime of the
damage cut), which the notes now say. Devotion Aura is supplied by the kit;
Sanctuary does not exist in the client. No line changed.

## 5. Encounter

After the fixes the geared paladin is not "far more" at risk than the warrior
(0.1% against 0%), so the 2,500 minimum is not above what a Phase 1 paladin can
hold with 440 healing. The bare gap (31% against 1%) is the real class
difference: 12.8% less health (the class base health in the engine's tables is
1,381 against 1,689, plus 66 less stamina from the class sets), a block that
fills less of the table (Holy Shield +30% against Shield Block +75%), and no
rage-fed cooldowns. All three come from carried Era tables and the client's own
ability text; none is a model error I could find.

The encounter is already as hard as the weakest geared tank allows. The bear is
the one that moves first:

| Minimum, healing | Warrior raid / bare | Paladin raid / bare | Bear raid / bare |
|---|---|---|---|
| 2,500, 440 (shipped) | 0.0 / 1.3 | 0.1 / 31.5 | 3.4 / 13.7 |
| 2,600, 440 | 0.0 / 4.0 | 0.3 / 50.2 | 8.8 / 27.3 |
| 2,700, 440 | 0.0 / 10.7 | 1.1 / 67.4 | 19.2 / 46.6 |
| 2,600, 480 | 0.0 / 0.7 | 0.0 / 23.2 | 1.9 / 8.0 |
| 2,800, 500 | 0.0 / 3.0 | 0.3 / 44.2 | 5.6 / 19.5 |

(chance of death, percent, 8,000 to 10,000 iterations.) Making the warrior die
bare (2,700) would make a geared bear die one fight in five. So: **no change to
`tank-encounter.json`**. Its reasons stand; the file is untouched.

## 6. What remains

- Holy Shield's 30% is the live text and 20% is in the wago spell. Confirm on the
  next build; the sensitivity is above.
- Hammer of the Righteous: whether "damage per second" includes attack power,
  and what the 120 dummy is, are unconfirmed (flagged in the code and the
  rotation notes).
- Block value from gear (section 2) and the empty shield block value: data lane.
- Benediction is 3% a rank in the engine and 2% in the client (Retribution
  goldens move when it is fixed).
- `blessing_of_sanctuary` is still a request id in `simbuffs.json` for a spell
  this client does not have; it applies nothing. Remove it with the next buff
  table change.
- The ladder walks a guide build top row down, and Iron Creed (tier 5) and
  Sacred Duty (tier 2) sit above Holy Shield, so the band 38 and 40 characters
  spend their 29 and 31 points without Holy Shield (and without Templar's Bulwark
  at 38). Their tank figures describe a build no player would choose. Keeping
  Holy Shield at 40 would mean dropping Iron Creed from the level 60 build,
  which costs the 60 band far more; the right fix is a prerequisite-first walk
  in `LadderTalentString` (the leveling lane).
- The warrior receives no Devotion Aura in the raid preset (the preset's one
  paladin group brings a Retribution paladin's Sanctity Aura). A real raid also
  has a tank paladin running Devotion, worth about 4% damage taken to the
  warrior; this is a raid-composition decision for the preset, which is another
  lane's file, so I left it. It would narrow the geared gap, not close the bare
  one.
- Final stats in the tank report are before the fight's auras; Devotion Aura's
  735 armor appears in the uptime table, not in the armor line.
- Fork `apl-check` still reports seven specs whose fork copies are older than
  the curated rotations (druid-feral, hunter-survival, both rogues, two warlocks,
  warrior-arms); not mine, and the paladin-protection copy matches.

## Verification

- Fork: `go test --tags=with_db ./sim/paladin/... ./sim/core/ ./sim/conformance/`
  green; the whole `./sim/...` run fails only `sim/web` (no `binary_dist`) and the
  mage P1 stat-weights golden, both pre-existing. Conformance rows for Holy
  Strike, Hammer of the Righteous and Seal of Fury read `declared, matches`.
- Site: `go test ./...` under `sim/` fails only the ladder goldens of other specs
  (druid-balance, priest-shadow, shaman-elemental, three warlocks) and the
  long-standing ladder violations, identical on the base fork commit
  (`921b7f8a4`, checked with a scratch worktree). `make apl-check` for
  paladin-protection matches. `web` `builds.test.ts` and `_sections.test.ts`
  accept the new build string.
