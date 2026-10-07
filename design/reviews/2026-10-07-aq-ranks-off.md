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

## Round 2: class-package tests and the book-rank flag

Fork:
- `core.MaxTrainerRank(totalRanks)` (sim/core/config.go) is the one place that says "the last rank is a book rank". The mage, druid balance, rogue, hunter and warlock rank tests use it or `core.TernaryInt32(core.IncludeAQ, book, trainer)`, so they pass with the constant either way. With `IncludeAQ = true` the hunter, warlock, rogue, mage and druid suites pass; the P1 mage golden and the warrior Heroic Strike test then fail only because the reference APLs name trainer ids, which is the intended pairing.
- `buffs_debuffs.ts` now shows Blessing of Might 19838 and Battle Shout 11551. `timeline.tsx` already listed every Heroic Strike rank, so it needed nothing.
- Battle Shout rank 6 is 11551 in `core` and 27578 in the warrior package's generated table (a dedup artifact the warrior lane noted); the engine registers the core id, so rotations and the UI use 11551.

Site:
- `data/curated/book-ranks.json` lists 36 book ids, each with a basis: `engine` (the fork gates it on IncludeAQ), `vanilla` (the vanilla 1.9 book list, not gated by the fork), `owner` (named in the brief: Greater Blessing of Might 25916 and Gift of the Wild 21850, whose book status I could not confirm; 21850 is the trainer's level-60 rank in vanilla, so check it), `copy` (a Forever copy of a book rank: Rejuvenation 417068, Renew 425277, Backstab 462717). It is a floor: an id missing from it is treated as a trainer rank.
- The `spellranks` emitter writes `book: true` for those ids (only when true) and `data/builds/1.60.1.70009/spellranks.json` carries 36 flags. `manifest.json` was not regenerated, so its spellranks.json hash is stale until the nightly refreshes it.
- `sim/internal/spellranks` and the ladder skip book rows unless `core.IncludeAQ` is on, through one function (`spellranks.RankAvailable`) that reads the engine's constant. `TestRewriteRotationRanksAtMaxLevelIsUnchanged` and `TestRotationAtMaxLevelParsesTheEmbedDirectly` pass unchanged: the highest learnable rank at 60 is the trainer rank the rotations name.
- Shaman Enhancement named Strength of Earth 25361 (a book rank that the engine's IncludeAQ switched on); it now names 10442, rank 4.
- Ladder goldens regenerated: the learned lists lose the book ranks and the level-60 violation rows that the first pass produced are gone, except Ferocious Bite (feral), Backstab (combat) and Shadow Bolt (destruction), which the ladder rotations do not cast at 60 under any id (the same rows exist at lower levels).

Still red: `data` `test_every_spell_the_rotations_name_exists_with_that_rank` fails because the curated druid-feral rotation names Lacerate 1322605, which the 1.60.1.70009 spell table lacks; it fails on main too and is unrelated to this change. `sim/web` still fails to build on the missing `binary_dist` import.

## Open question

If a Forever raid ever drops these books, the ranks flip back. That should be a phase switch (a per-phase predicate replacing the constant), not a flip of the constant. The conformance goldens and ladder goldens would regenerate under it.
