# Live talents overlay: Wowhead's hotfix-aware trees over the client's trait tables

2026-10-07, lane `talents-live`.

## Problem

Blizzard ships most Forever talent-tree changes as hotfixes. Hotfixes live in the client's
DBCache and never reach the CDN DB2 tables wago.tools exports, so the trait tables
(`TraitNode`, `TraitNodeEntry`, `TraitDefinition`, `TraitEdge`, `TraitCond`) that
`pipeline normalize` reads describe the tree as shipped, not as live. Every consumer of
`data/builds/<build>/talents/<class>.json` (the planner, the guides, the engine fork's
`tools/talentgen`) inherited the gap.

The 1 October 2026 Fury rebuild is the proof: Lingering Rage, Furious Precision and Gore
Drinker, the removal of Boundless Rage, Improved Cleave and Precision, and Flurry moving
behind Death Wish are in neither build 1.60.1.70009 nor 1.60.1.70245, but are in Wowhead's
live payload (`nether.wowhead.com/forever/data/talents-classic?dv=100`). The same payload
shows the gap is wider than Fury: Feral Combat and Protection also differ, and so does tooltip
text across most trees (below).

## Design

1. `pipeline fetch-wowhead-talents --build <build>` (`pipeline/wowhead_talents.py`)
   downloads the payload with the pipeline's User-Agent, a 120 s timeout and three attempts
   with backoff on transport errors and 429/5xx. The body is a `WH.setPageData("...", {...});`
   envelope followed by other statements, so the payload is the first JSON object
   (`raw_decode`). It is saved to `data/builds/<build>/raw/wowhead-talents.json` with a
   `_meta` object (`url`, `fetched_at`). A body with no object, or without `trees` and
   `talents`, is refused before anything is written.
2. `normalize` prefers that file (`pipeline/normalize/wowhead_overlay.py`). The trait reader
   still builds the trees; the overlay keeps each tree's id, name, position, background and tab
   icon and replaces its talents with the payload's, field for field in the existing schema:
   `id` is the node id, `row`/`col` are tier/column, `requires` is the single prerequisite,
   `ranks` gives one client spell id per rank and `descriptions` one tooltip per rank. Without
   the file, normalize is unchanged (trait tables only). The overlay refuses a payload whose
   tree set differs from ours, two talents in one cell, a multi-prerequisite talent, a rank
   count that disagrees with the tooltips, or a prerequisite outside the tree.
3. Output schema: same keys and order. The one addition is `hotfix_only: true`, appended last
   and emitted only when true, so a build with no such talent is byte-identical to before. It
   is set when any rank's spell id is in no `SpellName` row of the build, meaning a
   hotfix-added spell; the Wowhead id and text are kept and the talent is not dropped.
   `talents.json` is refolded from the overlaid trees. `trees/` is background art and is
   untouched. The manifest records `talent_source: raw/wowhead-talents.json` and
   `provenance` for `talents.json` and `talents/`, and `refresh_manifest` now preserves such
   extra keys.
4. Tooltip text. Wowhead's `descriptions` are the HTML its page injects (`<br />`, colour
   spans, value markers in comments, `singular:one:many` pairs, an embedded spell tooltip table
   on Feral Charge). The planner renders plain text, so `clean_description` keeps the shown
   value, picks the singular or plural word by the number before it, drops the embedded table
   and strips tags. Tested.
5. CI: `data.yml` runs `fetch-wowhead-talents` right after `fetch-wowhead`. `bis.yml` does not
   normalize (it runs `quest-levels`, `item-sources`, `loot-merge`, `fetch`/`loot` and
   `addon-data`), so it needs no fetch step; its `addon-data` regeneration reads the committed
   `talents/` and picks the overlay up on its own.

The definition id (`definition`) is not used by the overlay: the node id is the join key and
the per-rank spell ids carry the client mapping.

## Result on 1.60.1.70009

The snapshot is `data/raw-forever/wowhead-talents-2026-10-07.json` (466 talents). Regenerated
`talents/*.json` and `talents.json` are committed. Five talents are `hotfix_only` (spell ids
1322605, 1322670, 1323963, 1323964, 1323967).

Per class against the committed trees. Rank-cap and prerequisite changes are listed with the
talent; classes not listed have none of that kind.

| Class | Added | Removed | Moved | Rank cap | Prerequisite | Text |
|---|---|---|---|---|---|---|
| Druid | Improved Shifting Power (Feral, 4,0, 2 ranks, hotfix_only) | none | Shredding Attacks (3,0 to 2,0), Predatory Instincts (4,0 to 4,3), Shifting Power (4,3 to 3,0) | Shifting Power 3 to 1 | Shifting Power gains Shredding Attacks x3 | 9 talents; "King of the Jungle" becomes "Shifting Power" (name, icon, spell 417046 to 1322605) |
| Hunter | none | Improved Serpent Sting (Marksmanship, 3,3) | none | none | none | 8 talents; icon swaps on Improved Stings and Predator's Edge |
| Mage | none | none | none | none | none | 7 talents; "Hot Streak" becomes "Heating Up" (name, icon, text) |
| Paladin | none | none | none | none | none | 14 talents |
| Priest | none | none | none | none | none | 9 talents; one icon spelling fix |
| Rogue | none | none | none | none | none | 8 talents |
| Shaman | none | none | none | none | none | 8 talents |
| Warlock | none | none | none | none | none | 10 talents; "Soul Harvesting" becomes "Soul Harvest" |
| Warrior Arms | none | none | none | none | none | 3 talents |
| Warrior Fury | Gore Drinker (5,2, 2 ranks, needs Enrage x5, hotfix_only) | Improved Cleave (2,0), Precision (4,0) | Furious Precision (2,3 to 2,0), Flurry (5,2 to 5,1), Improved Berserker Rage (5,0 to 4,0) | none | Flurry: Enrage x5 becomes Death Wish x1; Bloodthirst loses its Death Wish prerequisite | 7 talents; Boundless Rage becomes Furious Precision, Iron Will becomes Lingering Rage (both reuse the node id, new hotfix spell ids 1323963 and 1323964) |
| Warrior Protection | Iron Will (0,2, 5 ranks) | Toughness (1,2) | Focused Rage, Bastion, Improved Shield Bash, Vanguard, Improved Disarm, Improved Revenge, Improved Bloodrage, Anticipation | none | Last Stand loses its Improved Bloodrage x2 prerequisite | 2 talents |

Text changes in total: 85 talents. They are a mix, not all hotfixes:

* real hotfixed content: Deflection 2% to 1% parry, Redoubt 6% to 4% block, Resourcefulness
  50% to 30% chance, Combustion 4 to 3 critical strikes, Sniper Shot rewritten, Shifting
  Power, Primal Bite, Lacerating Strikes and Heating Up reworded;
* presentation differences: Wowhead states damage and healing as a range ("28 to 32 Frost
  damage") where our spell-text renderer states one value ("28"), and keeps trailing zeros
  ("0.50 sec", "2.00 base Armor");
* our renderer's unresolved tokens that Wowhead resolves: `$a1 yards` (Wild Growth, now 43.5),
  `${32+($rap*(5/100))}` (Summon Hawk), `$?a5487...` (Feral Charge, Berserk).

Unresolved on Wowhead's side too: Summon Hawk and Prayer of Mending print a bracketed formula
(`[32 / Ferocity: 48 / ... + (Ranged Attack Power * (0.05))]`, `[(172 + (Healing * 0.43)) ...]`).

Tests: the full data suite is 1,358 passed, 8 skipped, 3 failed.

* `test_apl.py::test_every_named_spell_matches_its_own_label_by_name` fails the same way on
  `main` (mage-frost's sidecar); not related to this lane.
* `test_addondata.py::test_the_cli_check_passes_on_the_committed_file` and
  `test_audit_addon.py::test_clean_build_has_no_drift_findings` fail because the committed
  `addon-data.json` and `Data.lua` were derived from the pre-overlay trees. They are drift
  gates, not tree expectations, so no test was edited. Regenerating `addon-data.json` and
  `Data.lua` was out of bounds for this lane (nightly owns them), so they clear once
  `bis.yml` or `data.yml` runs `addon-data` against the new trees.
* No existing test pinned the pre-overlay tree, so none was changed.

## Open questions

1. Numbers inside ability tooltips. Wowhead's payload carries talent text, and the numbers in
   it come from Wowhead's own spell data, not necessarily the hotfix cache. Bloodthirst still
   reads "35% of your Attack Power plus 30" there while Blizzard's 1 October notes say 45%.
   Deflection, Redoubt and Resourcefulness did move, so some numbers are hotfix-aware. Which
   are not is unknown without a per-spell comparison against a client DBCache; the same doubt
   applies to the engine's own spell values (the DB2 gap is the same gap).
2. Node ids are reused for different talents (105953 was Boundless Rage and is Furious
   Precision, 110857 was Iron Will and is Lingering Rage, 104951 was King of the Jungle and is
   Shifting Power). Anything keyed by node id (saved builds, guide build strings, the engine's
   talent generator, curated adoption lists) now means a different talent for those ids. The
   fs1 build string encodes by tree position, so a moved talent also shifts old strings.
3. Six talents' icons are not in `builds/<build>/icons` (`ability_warrior_secondwind`,
   `ability_warrior_incite`, `racial_troll_berserk`, `spell_druid_displacement`,
   `ability_hunter_aspectmastery`, `spell_fire_firebolt`). They were not committed here
   (outside the allowed paths); the `icons` step of `data.yml` downloads them. Until then
   those cells render as broken images.
4. Name-based consumers: the mage guide and curated data refer to "Hot Streak", "Soul
   Harvesting", "Improved Serpent Sting", "Boundless Rage" and "Improved Cleave" by name. Not
   audited here.
5. Wowhead is the single source for the live tree, with no second check. A cheap guard would
   be to fail the fetch when the tree count, talent total or a per-class shape drops sharply
   against the previous snapshot. Not added.
6. The payload carries hotfixed talents but not hotfixed spell effects: a `hotfix_only` talent
   has Wowhead's text and ids but no value in the client's curve tables, so the engine fork
   must model these (Lingering Rage, Furious Precision, Gore Drinker, Shifting Power,
   Improved Shifting Power) from the text.
7. `wowhead-diff --snapshot data/raw-forever/wowhead-talents-2026-10-07.json --build 1.60.1.70009`
   reports 0 differences (466 against 466) by construction; a diff against the trait tables needs the
   pre-overlay trees, which are in git at the parent of the overlay commit.
