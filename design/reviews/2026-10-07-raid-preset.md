# Raid-ready preset for the level-60 headline (2026-10-07)

Branch `raid-preset`. The published level-60 numbers came from a bare character. Level 60 now also
carries a "raid" entry per faction, measured under one Phase 1 raid context that is the same for every
spec; the bare entries are unchanged and stay as a toggle.

## Contract (what the web lane builds against)

- Every entry in a spec's `bands[]` has a string `preset`: `"bare"` for every existing entry (all
  bands, both factions) and `"raid"` for the new level-60 entries.
- Level 60 has four entries per spec: bare alliance, bare horde, raid alliance, raid horde. Bare entries
  are always written before their raid twins (the API, talent-search, rotation-search and seedguild
  readers take the first match for a band and faction, and a test holds that order).
- A raid entry has exactly the bare level-60 entry's shape: race, talents, weights (swept on the buffed
  character), set_dps, slots, coverage and the rest. Nothing else in a band entry changes.
- The file gains a top-level `presets`:
  `{"raid": {"label", "buffs": [{"id","label"}], "debuffs": [...], "consumes": [...], "notes"}}`.
  The lists are the ids applied to that spec's raid entries, resolved for its role, with the class kit
  already layered on top (a kit slot is not listed twice). Ids are the request vocabulary
  (`sim/request/IDS.md`). Only the preset's own additions are listed, not the kit.
- Fight: unchanged, the default encounter. No world buffs.

Source of truth: `data/curated/presets.json` (one statement, a `raid` preset, a reason per entry).
Loader and resolver: `sim/request/preset.go` (`LoadPresets`, `Resolve`, `Layer`). An unknown id, a
buff filed as a debuff, or an entry without a label stops the run.

## The preset

Raid buffs, every spec, plain (unimproved) forms: Arcane Intellect, Mark of the Wild, Power Word:
Fortitude, Divine Spirit, Blessing of Might, Blessing of Wisdom, Battle Shout, Trueshot Aura, Leader of
the Pack, Moonkin Aura, Sanctity Aura, Strength of Earth Totem, Grace of Air Totem, Mana Spring Totem.

- Plain forms: none of the improved forms exist in the live trees (there is no Improved Mark of the
  Wild, Power Word: Fortitude, Blessing of Might or Wisdom, Battle Shout, Enhancing Totems or Improved
  Hunter's Mark in `data/builds/1.60.1.70009/talents`), so a Phase 1 raid cannot bring them.
- Moonkin Aura is in: the live Balance tree has Moonkin Form and the engine models the aura as a raid
  buff (+3% crit). Trueshot Aura and Leader of the Pack are live talents.
- Blessing of Kings is out: the paladin trees do not grant it (it is a Protection skill-line spell in
  the trainables, not a talent), and the engine applies it to Alliance characters only, so the factions
  would differ.
- Windfury Totem is the engine's main-hand imbue `windfury` (there is no totem buff id), so it takes
  the main-hand imbue slot of melee and hunter specs.

Target debuffs: Curse of the Elements (the Forever all-magic version), Sunder Armor (the engine
applies five stacks), Faerie Fire, Judgement of Wisdom, Judgement of Light, Hunter's Mark,
Improved Scorch, Winter's Chill, Curse of Recklessness. Left out: Shadow Weaving (a self-buff now),
Expose Armor, Thunder Clap, Demoralizing Shout, Curse of Weakness and the rest (no Phase 1 raid needs
them on top of Sunder, and they would double-count armor and attack power the preset already has),
Curse of Shadow (Curse of the Elements covers magic in Forever), Thunderfury and Gift of Arthas
(items the preset does not assume).

Consumables, by the spec's group (class first, then reference stat) plus class and spec additions:

| Group | Gets |
| --- | --- |
| caster (spell power) | Flask of Supreme Power, Greater Arcane Elixir, Runn Tum Tuber Surprise, Brilliant Wizard Oil (main hand), Major Mana Potion, Demonic Rune |
| melee (attack power) | Elixir of the Mongoose, Juju Power, R.O.I.D.S., Grilled Squid, Juju Might, Windfury (main-hand imbue) |
| melee, dual wielders only | Elemental Sharpening Stone in the off hand |
| hunter | the melee list without the stone |
| warrior | Mighty Rage Potion |
| paladin, shaman, hunter | Major Mana Potion |
| mage-fire, warlock-destruction | Elixir of Fire Power |
| mage-frost | Elixir of Frost Power |
| warlock-affliction, warlock-demonology, priest-shadow | Elixir of Shadow Power |

- Potions and the Rune are used by the engine on cooldown. Healers and tanks take the caster and melee
  groups by reference stat (the ranker measures damage, not healing).
- Elixir of Giants is not in the 1.60 client's item table, so Juju Power is the strength item; Juju
  Might beats Winterfall Firewater for the shared attack-power slot. Greater Firepower and Greater
  Frost Power do not exist in the client either, so the school elixirs are the lesser ones. Elemental
  Sharpening Stone is in the table (level 50), so Dense is not needed.
- Windfury occupies the main-hand imbue, so a single-weapon melee spec has no free imbue slot for a
  stone, and the class kit still wins any slot it holds (a rogue keeps its poisons, an enhancement
  shaman its Windfury Weapon). Nothing from Ahn'Qiraj or Naxxramas is used (no Zanza buffs, no Titans
  flask). No world buffs.
- Assumption the owner should confirm: Flask of Supreme Power is obtainable at launch (alchemy, client
  level 50). The brief asked for it; whether Black Lotus-style ingredients exist is not in the data.

## Two-spec comparison (level 60, scratch run, 100 weights iterations, default verify)

| Spec | Faction | Bare set_dps | Raid set_dps | Ratio |
| --- | --- | --- | --- | --- |
| mage-fire (caster) | alliance | 324.7 | 622.3 | 1.92 |
| mage-fire | horde | 312.3 | 608.7 | 1.95 |
| warrior-fury (melee) | alliance | 288.0 | 702.3 | 2.44 |
| warrior-fury | horde | 292.1 | 736.5 | 2.52 |

Ablation on the raid-pick gear (3000 iterations, alliance character, preset parts applied alone to a
no-kit character): fire mage 323 bare, 379 buffs, 398 debuffs, 471 consumables, 462 buffs plus debuffs,
624 all. Warrior 290 bare, 351 buffs, 454 debuffs, 386 consumables, 551 buffs plus debuffs, 706 all. The
largest single parts are Sunder Armor (+31% for the warrior), the Flask and Major Mana Potion (+15%
each for the mage), Windfury (+11%) and Curse of the Elements (+11%). Run time: about 32 s for a
level-60-only mage-fire run (both passes); the second pass touches band 60 only, so a nightly run
grows by roughly one band's share per spec, and `bis.yml` needs no new flag.

## What the engine lacks or gates (a Phase 1 raid would bring it)

1. Faction gating. The engine applies Blessing of Might, Blessing of Wisdom, Blessing of Kings and
   Sanctity Aura to Alliance characters only and Strength of Earth, Grace of Air, Stoneskin and Mana
   Spring Totem to Horde characters only (`sim/core/buffs.go`, `isAlliance`/`isHorde`). Forever has
   paladins and shamans on both sides, so both factions should get all of them. Today an alliance raid
   entry has no totems and a horde raid entry has no blessings or Sanctity Aura, which is part of why
   the two factions' raid numbers differ more than their bare ones (warrior horde 736.5 vs alliance
   702.3). The ids are applied to both factions; the engine drops what its gate refuses. The published
   `notes` say so. Dropping the gate in the fork is the fix, and the preset needs no change when it
   lands.
2. No Windfury Totem raid buff: it is modelled as the main-hand imbue, so it competes with sharpening
   stones and the rogue/shaman kit for one slot (a rogue never sees it, because poisons win the slot).
3. No Elixir of Giants, Greater Firepower or Greater Frost Power in the client table; no nature-power
   consumable slot for elemental shamans (Elixir of Nature Power is in the client but not the engine).
4. Single consumable slots: Juju Might and Winterfall Firewater share one; Mighty Rage and Major Mana
   Potion share the potion slot, so a mana-using warrior does not exist and a hybrid takes its mana
   potion rather than a rage one.
5. Improved Scorch alone lowered a fire mage's DPS by about 4% in the ablation (310.6 vs 323.4); in
   the full preset it is part of the 624. Worth a look at the fire rotation's Scorch handling when the
   debuff is already up; not changed here.
6. Healers are ranked on damage, so they carry the caster group; their healing consumables (Flask of
   Distilled Wisdom, Mageblood) are not in the preset.

## Other readers of the bis files

- `data/pipeline/addonbis.py` keyed (band, faction) with last-wins, which would have shipped raid gear
  in the addon's Data.lua; it now skips non-bare entries (test added).
- The sanity pass (`published_test.go`) checks a raid entry against its bare twin and does not use it as
  the next band's baseline.
- The API's `LoadBand`, talent-search, rotation-search and seedguild take the first entry for a band
  and faction, which is the bare one by write order.

## Verification

`go test ./...` under `sim/` passes; `cd data && uv run pytest -q --no-cov` passes except
`tests/test_apl.py::test_every_named_spell_matches_its_own_label_by_name` (mage-frost's sidecar is out
of date with its rotation; unrelated to this branch). The ladder and guide smoke tests run bare and are
unchanged. Data tests: `data/tests/test_presets.py`.
