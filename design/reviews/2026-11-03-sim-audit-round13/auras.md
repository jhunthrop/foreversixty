# Sim audit round 13: auras lane

Scope: racials (`sim/core/racials.go` and its files), raid buffs and debuffs in the raid preset, consumables, item effects and set bonuses on the published BiS files, and the crit plumbing. This was a fresh re-audit. Round 12's auras and itemscrit rows were re-read, not trusted.

Sources:
- The 1.60.1.70753 client tables (wago CSVs in `scratchpad/r11/c70753/builds/1.60.1.70753/raw`), read with round 12's `sp.py` and `$R/auras/racesla.py`. The 70753 hotfix push (`data/builds/1.60.1.70753/hotfix-push.json`) moves no table, so the shipped rows are the truth.
- `data/builds/1.60.1.70753/combos.json` for the race and class pairs.

Branches: engine `lane/r13-auras` (from 076a5a517), site `wt/r13-auras` (from 93c5d2510).

**Found: 2 CONFIRMED-BUG rows, both fixed. By origin: R12 1, old 1.** There are also three note fixes. The third renames the Skyborne racial Read Leyline everywhere. No published number moves.

## Racials

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| R1 | **The Alliance Skyborne's Read Leyline was not modelled, and its note was wrong.**<br>- **What was wrong:** the engine registered nothing. Its comment said no cooldown was published and the effect was "100% health and mana regeneration".<br>- **What the client says:** the spell is named Read Leyline (1259705). It casts Energized (1270842): aura 110 (mana regeneration percent) at 50 and aura 88 (health regeneration percent) at 50, for 15 s. SpellCooldowns gives a 120000 ms recovery with a 1500 ms global cooldown. The Spell text reads "increasing your Health and Mana regeneration by 50% for 15 sec". Standing at a leyline makes the caster Attuned (1259691) for more. A boss room has no leyline, so only the base 50% is modelled.<br>- **What changed:** the spell is registered. Its aura adds 0.5 to the spirit regeneration multiplier, the same lever Innervate uses. An APL can cast it, but the sim never fires it on its own: it costs a global cooldown, and spirit regeneration is mostly off inside the five-second rule. The site's race note, which said the client has no such row, is corrected too. | SkillLineAbility (Skyborne Racial 2980), SpellEffect, SpellMisc and SpellCooldowns 1259705 and 1270842, Spell text. `TestReadLeylineIsFiftyPercentSpiritRegenForFifteenSeconds` failed first (the spell was not registered), then passed. | CONFIRMED-BUG (latent: no published entry is Skyborne) | old (the racial table of 2026-10-07; no round opened the Skyborne actives) | E 2fde66add, S 9203c7753 |
| R2 | **The table called three racials unread, but the client names all three.** Tauren's skill line 124 carries War Stomp, Cultivation 20552 (an active, subtext "Racial"), Endurance and Plainsrunning 1259918 (a passive, subtext "Racial Passive"). Undead's skill line 220 has a fourth row: Underwater Breathing 5227 (aura 155, breath lasts 300% longer). The entries are now confirmed, with their spell ids. None of them has a combat effect. `UnconfirmedRacials()` is now empty. No site code reads it. The Wind Blessed, Human Spirit and Touch of the Grave notes now cite the 70753 rows (2, 10 and 15000 ms) instead of older shipped values. | SkillLineAbility, Spell subtext; `TestTheClientNamesEveryRacial` | note fix | old | E 2fde66add |
| R3 | Every SkillLineAbility row on the ten racial skill lines was re-read against the engine. All match:<br>- **Class versions:** Expansive Mind is mana for priest, mage and warlock (20591, mask 400), rage for warrior (1259802) and energy for rogue (1259803). These are exactly the gnome classes in `combos.json`. Touch of the Grave is 20% for warrior, paladin and rogue (mask 11) and 30% for priest, mage and warlock (mask 400). That covers every undead class, including the new paladin. Command is 20575 for warlocks and 20576 for hunters; 21563 is a dummy (aura 0) for the other orc classes, the new orc mage among them. Eureka! has one row for each gnome class.<br>- **Amounts:** Blood Fury is aura 166, 167 and 317 at 10, with a 2 min cooldown and no global cooldown. Berserking is aura 319, 140 and 65 at 10, 15 s, 3 min, no global cooldown. Elune's Light is aura 290 at 10, 15 s, 3 min. Ruthless is aura 290 at 1. Artisan's Precision is aura 54 and 55 at 1. Big Game Hunter, Beast Slaying and Elemental Insight are aura 168 at 5. Wind Blessed is aura 342 and 65 at 2. Quickness is 1% dodge. Endurance is aura 133 at 5. Stoneform is aura 87 at -10 physical for 8 s, 3 min.<br>- **Touch of the Grave's drain** (1260198) ignores damage modifiers on both sides, as the engine has it: Attributes_4 0x100 ignores damage-taken modifiers, Attributes_6 0x20000000 ignores the caster's. The rest (Will to Survive, Perception, Shatter Curse, Hardiness, Find Treasure, Shadowmeld, Wisp Spirit, Will of the Forsaken, Cannibalize, War Stomp, Escape Artist, Engineering Specialization, Rapid Regeneration, Regeneration, Walk on Air, Skysight) has no damage, threat or mana effect. Their cooldowns match SpellCooldowns. | client rows above | Aligned | - | - |
| R4 | **The racial's name is Read Leyline.** The client (1259705) and its tooltip spell it that way. Blizzard's notes wrote "Read Ley Line", and so did the engine, `race-table.ts`, the curated race notes and the Balance and Restoration druid guides. All of them now say Read Leyline. The guide's race table matches a racial name with or without a space inside a word ("Read Leyline" finds "Read Ley Line"), so it keeps finding the effect until the build's `races.json` is regenerated from the curated file. | Spell 1259705; `race-table.test.ts` and `_prose.test.ts` pass; the guide-bis and guide-page tests pass after the fixture sync | note fix | old | E 1c2190b4d, S 7aa6b12ce |

## Raid buffs and debuffs

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| B1 | **A merge undid round 12's preset reasons.** Round 12's S f8b327ab8 corrected four reasons in `data/curated/presets.json`. The int/r12 merge of `wt/r12-regmelee` (8850cbc26) resolved its conflict in that file with the older text, so the BiS page again said:<br>- Divine Spirit is "a Discipline talent". It is a trainer spell in 70753.<br>- Trueshot Aura gives "Attack power". It is aura 124 only, so ranged attack power.<br>- Heart of the Lion is "276 at 60". The client gives 40 + 4 × 60 = 280.<br>- Also lost: the Windfury Totem line saying the bonus ends with the totem.<br>The four reasons are restored, and a test now pins them. No number moves. | `git show 3290c0a2f:data/curated/presets.json` against 8850cbc26. `TestRaidPresetReasonsMatchTheClient` failed first on all four, then passed. The preset and published-preset tests pass. | CONFIRMED-BUG (published text) | R12 8850cbc26 (coordinator merge) | S df60eff5a |
| B2 | Debuffs re-read on 70753:<br>- Hunter's Mark 14325 is aura 127 at 143, ranged only. The engine has 143.<br>- Sunder Armor 11597 is aura 674 at -450, 5 stacks.<br>- Faerie Fire 9907 and Curse of Recklessness 11717 are aura 674 at -505. Curse of Recklessness has no attack power half.<br>- Judgement of the Crusader 20303 is aura 14 at 161, 40 s.<br>- Judgement of Wisdom 20355 has proc type 0x222A8: melee, ranged and harmful spells, as the engine filters them. Judgement of Light 20346 has 0x28: melee only, as the engine has it.<br>- Sunder and Expose share the major armor category. Faerie Fire is minor, and Curse of Recklessness stacks with both, which matches the client: there is no exclusivity row.<br>- The 50% Judgement roll is still the stated assumption (the intermediate 1826 and 5373 are dummies). | SpellEffect, SpellAuraOptions; `core/debuffs.go`, `raid_debuff_ranks.go` | Aligned | - | - |
| B3 | Raid buffs:<br>- Leader of the Pack (24932) and Moonkin Aura (24907) add 3% crit to every table, once each. A feral druid's own talent sets the same raid flag, and a Moonkin's form checks `raidHasMoonkinAura`, so neither counts twice.<br>- The totem, blessing and buff ranks are as round 12's B1 found. Nothing changed in 70753.<br>- Kings with Heart of the Lion is still open (round 7). | `core/buffs.go:231-245`, `druid/druid.go:241-245`, `druid/forms.go:316` | Aligned | - | - |

## Consumables

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| C1 | Shared cooldowns from ItemEffect on 70753:<br>- Every potion in the presets is category 4, 120 s: Major Mana, Mighty Rage, the Frenzy and Spellblasting 30-second potions, and Greater Stoneshield.<br>- Demonic Rune 12662, Dark Rune 20520 and Thistle Tea 7676 share category 1153, 120 s. The engine's conjured timer matches, and no spec carries both a rune and Thistle Tea.<br>- Greater Mageblood and Elixir of Holy Power are category 2593, a 3 s elixir lockout.<br>- Nightfin Soup has no category.<br>- Amounts are as round 12 found (auras C1, itemscrit K3): nothing in 70753 moved. | ItemXItemEffect, ItemEffect | Aligned | - | - |
| C2 | **Mana gems share the rune category, and the engine already has them that way.** The guides lane questioned C1.<br>- **What the client says:** Mana Agate 5514, Mana Jade 5513, Mana Citrine 8007 and Mana Ruby 8008 are SpellCategory 1153 with a 120000 ms category cooldown, the same as Demonic Rune 12662 and Dark Rune 20520. The gem amounts (400, 600, 850 and 1100, with variance) match the engine's 375-425, 550-650, 775-925 and 1000-1200.<br>- **What the engine does:** `mage/mana_gem.go` gives every gem a 2 min cooldown on `GetConjuredCD()`, the same timer `registerConjuredCD` gives the Demonic Rune. The 60 s category-30 hold (`HoldRestoreCategory`) is applied on top. So a gem and a rune hold each other for two minutes, as the client says.<br>- A pinning test passes on the unchanged engine, so this is not a bug and no number moves. No before/after was needed: nothing changed in the mage or warlock code. | ItemXItemEffect, ItemEffect, SpellEffect 5405, 10052, 10057, 10058, 16666; `TestManaGemAndDemonicRuneShareTheTwoMinuteCategory` (pins, passes on the old engine) | Aligned | - | E 1c2190b4d (test only) |

## Items, set bonuses and crit

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| I1 | The census of the published sets (round 12 `census.txt`) was re-read against 70753:<br>- Frozen and Molten Heart of the Mountain share category 1141 with a 15 s lockout. The engine puts both on the offensive-trinket timer for 15 s. Chthonic Power is aura 13 at 29 on mask 48 (Frost and Shadow); Molten Fury is 55 attack power and 55 ranged attack power.<br>- Ankh of Life's self-heal also sits in category 1153 with the runes. The engine does not model it, and a self-heal changes no output.<br>- The set-bonus flat decoder (`core/client_sets.go`) reads aura 54 and 55 into one Hit stat. That is the Forever hit merge (research/08-stats.md 12.1), not a bug. Crit splits 52 (melee), 57 (spell) and 290 (every table) as the law says. | `core/client_sets.go:351-492`, `common/item_effects.go:3692-3760` | Aligned | - | - |

## Before/after

Engine 2fde66add with site 9203c7753, against `$R/before.tsv`. Raid 60, horde, 10,000 iterations. Identical:

| spec | s1 | s50001 |
|---|---|---|
| mage-fire | 673.0 → 673.0 | 672.8 → 672.8 |
| warrior-arms | 731.2 → 731.2 | 731.8 → 731.8 |
| priest-shadow | 720.0 → 720.0 | 720.0 → 720.0 |

Read Leyline is only cast from an APL, and no published entry is Skyborne. B1 is text only.

## Files to regenerate

- `data/builds/1.60.1.70753/races.json` (the race notes and the Read Leyline name, from the curated file, by the data run).
- No fork goldens and no ladder goldens.

## Measure at refresh

- None.

## For other lanes

- **Coordinator:** check presets.json after every merge. 8850cbc26 took one side of a conflict and dropped another lane's edits (B1).
- **Guides:** Read Leyline is 50% mana and health regeneration for 15 s on a 2 minute cooldown and the global cooldown (1259705 → 1270842).

## Still open

1. Kings with Heart of the Lion (round 7).
2. The Judgement of Wisdom and Light 50% roll (server script).
3. The Read Leyline Attuned bonus at a leyline (1259691) is not modelled. A raid room has none.
