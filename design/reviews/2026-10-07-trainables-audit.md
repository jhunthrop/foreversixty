# Trainables audit, 2026-10-07

Build 1.60.1.70009. Question: which abilities does a Forever player learn that the engine never registers?

## Method

1. `SkillLineAbility` was fetched for the build (it was already in `wago.TABLES`; the local `raw/` never had it) and committed at `data/builds/1.60.1.70009/raw/SkillLineAbility.csv`.
2. `pipeline trainables --build <b>` (`data/pipeline/trainables.py`) writes `data/builds/<b>/trainables/<class>.json`. It lives in its own directory because `spellranks` globs `spellconst/*.json`.
   - Rows on the 24 class skill lines (`SkillLine.CategoryID == 7`, the lines named in the lane brief; Engraving, Pet - *, Beast Training and Runes lines are excluded).
   - Class is the single `ClassMask` bit, or, when the mask is 0 (the Affliction Curse of the Elements ranks, hunter Lacerate), the spell's `SpellClassOptions.SpellClassSet`. A row naming neither is dropped (pet abilities under Survival).
   - Ranks of one ability are grouped by name and by `SupercedesSpell` links (Frostfire Bolt's three ranks are not linked to each other). Learn level is `SpellLevels.SpellLevel`.
   - Active means any rank has a power cost, a cast time or a cooldown.
   - `(DNT)`, `Copy of`, and `Test`/`Testing` spells are dropped.
3. **Gap in the brief's premise.** Unstable Affliction (427717) and Hydra Shot (1293020) have no `SkillLineAbility` row at all in this build, so the table alone cannot see them. A second source tags them `source: "class_spell"`: a class-family spell with rank 1 or more, a level, an active shape, and a name not listed on a pet or engraving skill line. This heuristic also admits a few non-spellbook spells (rogue poisons, mage Volleys and Sleep, hunter Widow Bite and Sonic Blast, warrior Howling Blade and Recycle, warlock Haunt and Dispel Magic). Every row carries its `source` so they can be told apart.
4. The fork's `sim/conformance` loads `client/trainables/<class>.json` (a copy, like the spellconst fixtures) and lists, per class golden, every active trainable with no rank id in any preset's spellbook at levels 10..60 (empty-talent build) or in a talent-gated build. `SUMMARY.md` carries the per-class counts.

## Counts

| Class | Active trainables | Not registered |
|---|---|---|
| Hunter | 58 | 40 |
| Mage | 75 | 55 |
| Warlock | 65 | 41 |
| Paladin | 57 | 44 |
| Warrior | 51 | 21 |
| Druid | 64 | 44 |
| Priest | 67 | 57 |
| Shaman | 64 | 42 |
| Rogue | 63 | 33 |

## Reading the list

The comparison is by spell id, so it over-reports. Many entries are the level-1 rune-style copy of an ability (`SkillLineAbility.AcquireMethod` 3: Raging Blow, Starsurge, Chimera Shot, Lava Lash, Mind Spike, Chaos Bolt) which the engine may register under another id or as a talent-granted ability. Before treating any row below as missing, check the class package for the ability by name. The ids in the goldens are the ones to grep for.

## Unregistered active trainables that plausibly affect DPS

The full per-class tables (level, ranks, cost, cooldown) are the last section of each `sim/core/testdata/conformance/<class>.golden.md` in the fork.

- **Mage**: Frostfire Bolt (40), Icy Veins, Arcane Surge, Arcane Barrage, Living Bomb, Living Flame, Balefire Bolt, Spellfrost Bolt, Frozen Orb, Cone of Cold, Felfire, Fire Volley and the other Volleys (source `class_spell`, likely not player spells).
- **Warlock**: Unstable Affliction (40), Haunt (40), Chaos Bolt, Curse of the Elements (new, 4 ranks from 20), Shadow Cleave, Shadowflame, Hellfire, Metamorphosis, Demonic Grace, Demon Charge, Menace, Vengeance, Curse of Weakness, Curse of Tongues, Curse of Exhaustion (debuff value to the raid).
- **Hunter**: Hydra Shot (60), Chimera Shot, Kill Command, Carve, Flanking Strike, Wyvern Strike, Lacerate (4 ranks, 30 to 60), Readiness, Trueshot Aura, Aspect of the Falcon, Scorpid Sting, Viper Sting, Intimidation, Widow Bite and Sonic Blast (`class_spell`).
- **Warrior**: Devastate, Raging Blow, Meathook, Quick Strike, Shockwave, Victory Rush, Concussion Blow, Mocking Blow, Charge and Intercept (gap closers).
- **Rogue**: Saber Slash, Shiv, Shadowstrike, Blade Dance, Quick Draw, Between the Eyes, Shuriken Toss, Poisoned Knife, Main Gauche, Venom, Gouge, Kidney Shot, Kick (interrupt), Cheap Shot.
- **Druid**: Starsurge, Starfall, Sunfire, Lacerate, Swipe, Maul, Primal Bite, Skull Bash, Feral Charge, Pounce, Survival Instincts, Enrage, Wild Growth and Lifebloom (heal throughput).
- **Paladin**: Avenging Wrath, Divine Storm, Hammer of the Righteous, Shield of Righteousness, Seal of Fury, Seal of Martyrdom, Seal of Light and Wisdom (utility), Swift Judgement, Horn of Lordaeron, Rebuke, Hammer of Justice.
- **Priest**: Mind Spike, Mind Sear, Vampiric Touch, Void Plague, Shadowfiend, Dispersion, Starshards, Penance, Power Infusion, Silence, Mana Burn, Holy Nova, Chastise.
- **Shaman**: Lava Lash, Molten Blast, Fire Nova (and Fire Nova Totem), Riptide and Earth Shield (healing), Windfury Weapon, Flametongue Weapon and Totem, Rockbiter and Frostbrand Weapon, Stoneclaw Totem, Mana Tide Totem.

Pure utility (Polymorph, Blink, teleports and portals, conjure spells, resurrections, cleanses, forms, travel, wards, stealth tools) is listed in the goldens and does not affect DPS.

## Follow-ups

- Rune-style ids versus the ids the engine registers: a name-based second pass in the conformance report would remove the false positives.
- The `class_spell` heuristic could be replaced by whatever the client uses to teach Unstable Affliction and Hydra Shot (a trainer table the pipeline does not fetch).
- `manifest.json` was not refreshed (the lane may not commit it); the next `simconst` or `levels` run will cover `trainables/`.
- `tests/test_apl.py::test_every_spell_the_rotations_name_exists_with_that_rank` fails on this branch for druid-feral spell 1322605. It reads only curated rotations and spellconst, which this lane did not touch.
