# Ahn'Qiraj ranks off at level 60 (2026-10-07)

Decision: the engine constant `core.IncludeAQ` is now `false`. A level-60 character casts the top rank a trainer teaches, not the Ahn'Qiraj book rank.

## Evidence

- Ahn'Qiraj is not on Forever's roadmap. The 9 December 2026 raids are Barrow Deeps, Hyjal Summit and Onyxia's Lair.
- The book ranks come from the Ahn'Qiraj spell books. Community trainer lists at level 60 top out one rank below them for every spell below, so a launch character cannot learn the book rank.
- The upstream wowsims engine inherited `IncludeAQ = true` for its own phase model; Forever has no such phase.

## Spells whose level-60 rank changed (book rank to trainer rank)

| Spell | Book rank, id | Trainer rank, id |
|---|---|---|
| Shadow Bolt | 10, 25307 | 9, 11661 |
| Corruption | 7, 25311 | 6, 11672 |
| Immolate | 8, 25309 | 7, 11668 |
| Fireball | 12, 25306 | 11, 10151 |
| Frostbolt | 11, 25304 | 10, 10181 |
| Eviscerate | 9, 31016 | 8, 11300 |
| Backstab | 9, 25300 | 8, 11281 |
| Deadly Poison | 5, 25347 | 4, 11356 |
| Serpent Sting | 9, 25295 | 8, 13555 |
| Aspect of the Hawk | 7, 25296 | 6, 14322 |
| Heroic Strike | 9, 25286 | 8, 11567 |
| Revenge | 6, 25288 | 5, 11601 |
| Ferocious Bite | 5, 31018 | 4, 22829 |
| Starfire | 7, 25298 | 6, 9876 |
| Battle Shout | 7, 25289 | 6, 11551 |
| Blessing of Might | 7, 25291 | 6, 19838 |

Raid-buff values that follow the flag (sim/core/buffs.go, unchanged by this lane): Battle Shout AP 232 to 193, Blessing of Might AP 185 to 155, Blessing of Wisdom MP5 33 to 30, Horn of Lordaeron Strength and Agility 89 to 70.15, Grace of Air Agility 77 to 67, Strength of Earth Strength 77 to 61 (totem ranks 5 and 3 to 4 and 2), Heroic Strike threat 173 to 145, Rogue client damage effect 34 to 27.

## Level-60 ladder movement per spec (relative, ladder DPS at level 60)

| Spec | Change |
|---|---|
| rogue-assassination | -6.7% |
| mage-frost | -6.3% |
| mage-fire | -3.5% |
| hunter-marksmanship | -2.5% |
| warlock-affliction | -2.2% |
| hunter-beast-mastery | -2.1% |
| warlock-demonology | -2.1% |
| warlock-destruction | -1.9% |
| warrior-fury | -1.8% |
| shaman-enhancement | -1.7% |
| warrior-arms | -1.5% |
| druid-balance | -1.1% |
| rogue-combat | -0.8% |
| mage-arcane | -0.5% |
| rogue-subtlety | -0.4% |
| druid-feral, hunter-survival, paladin-retribution, priest-shadow, shaman-elemental | 0.0% |

Levels below 60 are untouched. Fork suite goldens: mage P1 -1.6% to -9.0% (median -7.7%), warrior DPS -0.9% to -3.9% (median -2.8%), elemental shaman -0.1% to -0.8%, enhancement shaman +5% to +200% (the fixture sits at 100 to 300 DPS and swings with the Strength of Earth and Grace of Air totem ranks; treat it as unreliable until its rewrite).

## What this change touched

- Fork: `sim/core/config.go`, four `.results` goldens, the conformance goldens and SUMMARY (the level-60 rank rows for the changed spells leave the tables, because the trainer ranks are below level 60), and the reference APLs under `ui/*/apls` (spell ids moved from book to trainer ranks; `forever_*` files come from `make apl-sync`).
- Site: the 15 curated rotations that named a book id (id and rank number), `sim/request/apl/*` (synced copies), and the ladder goldens. No guide names a book rank, so no guide changed.

## Follow-ups for other lanes (not done here)

- Class-package tests still hardcode book ranks and fail: `sim/hunter` spellconst checks (Aspect of the Hawk 25296, Serpent Sting 25295), `sim/warlock` Corruption rank 7, `sim/rogue` Deadly Poison rank 5, `sim/mage` Frostfire Bolt resist test (indexes `Frostbolt[FrostboltRanks]`), `sim/druid/balance` Moonglow test (indexes `Starfire[7]`).
- `sim/core/spellconst/gen/generated_test.go` and `ui/core/components/inputs/buffs_debuffs.ts` (Blessing of Might 25291, Battle Shout 25289) and `ui/core/components/detailed_results/timeline.tsx` (Heroic Strike 25286) name book ids.
- `data/builds/*/spellranks.json` still lists the book ranks as learned at level 60, so `sim/request` `TestRewriteRotationRanksAtMaxLevelIsUnchanged` and `TestRotationAtMaxLevelParsesTheEmbedDirectly` fail: their premise, that a rotation names the highest data rank, no longer holds. `rotation()` skips the rewrite at level 60, so runtime is unaffected. The proper fix is a book-rank flag in the spell-rank data.
- Battle Shout rank 6 is id 11551 in `core.BattleShoutSpellId` but 27578 in the class-package generated table (a dedup artifact); the APLs use 11551, which is what the engine registers.

## Open question

If a Forever raid ever drops these books, the ranks flip back. That should be a phase switch (a per-phase predicate replacing the constant), not a flip of the constant. The conformance goldens and ladder goldens would regenerate under it.
