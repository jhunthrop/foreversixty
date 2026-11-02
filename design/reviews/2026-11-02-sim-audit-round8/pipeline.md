# Sim-accuracy round 8: pipeline lane

Scope: owner ruling 3, which the integrator widened mid-round (2026-11-01):
battlegrounds (AV, WSG, AB, Darkspear Islands) and PvP ranks are in Forever at
launch. Site branch `wt/r8-pipeline` (base 170b7b84c). No engine changes. No
regenerated loot.json is committed.

## 1. Loot generators: one source of truth

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| 1 | `loot` (data.yml) and `loot-merge` (bis.yml nightly) wrote different loot.json from the same inputs, so it flipped on every pair of runs. | `loot` read instance types from Map.csv (AV map 30, AB map 529 are InstanceType 3), and `instanced_zones` gave those zones no source. The merge had no Map.csv. It rebuilt types from the committed loot.json (`types_from_committed_loot`, which knows only dungeons and raids), so the same spawns became `world:*` sources and world-drop bands. Reproduced locally on 70707: the full run gives 3,199 sources and the merge 3,213. The committed file (45875a020) is the merge's answer. | CONFIRMED-BUG | S 52b3f1eb3 |
| 2 | The merge's PvP ranks were also rebuilt from loot.json, which loses the client's rank on re-itemised items. | ItemSparse.RequiredPVPRank: 15196 Private's Tabard = 5, 15197 Scout's Tabard = 5, 15198 Knight's Colors = 10. Their Forever twins 272694/272449/272695 have 0. loot.json lists only the twins (the classic ids are superseded). So the merge ranked the twins "by title", while the full run ranked the classic ids and re-itemised them (`reitemised_from`, `faction_source: classic-db`). The merge also wrote `superseded_by: null` for 15196-15198 into items.json, but the committed items.json has the full run's 272694/272449/272695. | CONFIRMED-BUG (provenance only; same items in the same rank sources) | S 52b3f1eb3 |

**Design.** `loot` writes the two client facts loot.json needs to
`data/builds/<build>/raw/client/loot-classification.json`
(`pipeline/loot/classification.py`): zone id to Map.InstanceType (battlegrounds
included, unchanged) and item id to RequiredPVPRank. Both commands then call
one function, `pipeline.loot._build_and_write_loot`, which runs the whole chain
from build_loot to npc ranks. So identical inputs give identical bytes by
construction. Any later drift can only come from the inputs, not from the
code. Why this design:
- The merge never needs Map.csv or ItemSparse.csv, and it no longer reads
  loot.json back as an input. That self-reference is what let a stale answer
  keep itself alive.
- One battleground rule, `loot_instance_types`, read by both commands. The
  switch is `pipeline/loot/constants.py: BATTLEGROUND_SPAWNS_ARE_SOURCES = True`.
  True files AV/AB spawns as open-world sources (today's committed state).
  False gives a battleground zone no source. Flipping it is one line, and it
  works in the merge too, because the file keeps the raw type 3.
- The merge fails fast (SystemExit naming `python -m pipeline loot`) when the
  file is missing. There is no fallback to guessing.
- The file lives under raw/ (manifest-exempt, like raw/classicdb/sources.json)
  with a .gitignore exception. The 70707 copy (265 zones, 449 ranks) is
  committed. Without it the nightly merge would refuse to run. It is an input
  cache derived from the 70707 Map.csv/ItemSparse.csv, not a regenerated
  output. Re-running `loot` reproduces it byte-for-byte.
- Deleted: `types_from_committed_loot`, `pvp_ranks_from_committed_loot`,
  `_load_committed_loot`, `pvp_faction.ranks_with_item_evidence` and their
  tests, plus the second copy of the gate chain in `merge_loot_files`.

**Tests** (data/tests/test_loot_merge_parity.py):
- `test_loot_and_loot_merge_write_byte_identical_loot_json[True|False]`:
  fixture with AV creature and object spawns (map 30, type 3) and the tabard
  shape (classic id with RequiredPVPRank 5, Forever twin). Runs loot, merge,
  loot, merge; all four loot.json files must be byte-identical. The switch
  alone decides whether the AV items are sourced. Before the fix both
  parameters failed ("loot-merge disagrees with loot").
- `test_loot_merge_fails_fast_without_the_full_runs_classification`: failed
  before (DID NOT RAISE).
- `test_battleground_spawns_are_loot_sources`: holds the switch at True
  (owner ruling).
- Committed build, needs FOREVER_ENGINE_PATH:
  `test_the_committed_loot_json_is_a_fixed_point_of_the_merge` (merge twice,
  same bytes) and `test_the_full_run_writes_the_merges_bytes_on_the_committed_build`
  (skips without raw CSVs; failed before the fix). Both pass.
- **Why there is no strict "merge reproduces committed loot.json" test:** the
  committed file carries the old merge's tabard provenance and the rank-10
  bug in section 2. A strict test would be red until the held data/bis runs
  regenerate loot.json, and this lane may not commit that regeneration.

**Local runs on 70707** (engine db at cd4dfdcc0, raw CSVs from the r7
worktree):
- New code: `loot` and `loot-merge` write byte-identical loot.json. The
  classification file is stable across runs. items.json is unchanged against
  committed in both runs.
- Old code: merge = committed (fixed point), full run differs by 18 sources.
- Against committed loot.json, the new output has no source added or removed
  (3,213 sources), and every battleground source is unchanged:
  `world:alterac-ram` ... `world:stormpike-guardsman`, `world:grunnda-wolfheart`
  (opens later), `world:anvil` with 3906.

  What a regeneration would change against committed:

  | source | change |
  |---|---|
  | pvp:rank-5:alliance, pvp:rank-5:horde, pvp:rank-10:alliance | `faction_source` title -> classic-db; `reitemised_from` {272694:15196}, {272449:15197}, {272695:15198} |
  | pvp:rank-10:horde | 21 -> 1 item (section 2); `reitemised_from` {272450:15199} |
  | pvp:rank-N:<side> (ranks 9, 11, 12, 14, 16, 17, 18) | 467 title-only listings move to 14 new `pvp:title-rank-N:<side>` sources: 11 A/H 47/49, 12 26/30, 14 26/29, 16 40/43, 17 41/42, 18 26/26, 9 20/22. 244 ranked pvp listings remain. Sources 3,213 -> 3,227. |

**Eligibility proof for the relabel** (`$R/pipeline/proof.py`, real 70707 regen
before and after 970ec9e90). For every item, compare the sorted list of
listings: (kind, faction, opens, rank or title_rank, and source id for
non-pvp kinds). Result: **0 items differ**. PvP listings at or under cap 10:
97 before, 97 after, same set. So once the ranker reads `rank or title_rank`
wherever it reads `Rank` for the cap, its default eligibility input is
byte-identical. That covers pvp sources and the vendor rows that inherit
through `vendorInheritsPvpRankGate`; each item keeps the same number of pvp
listings per side. Against the committed loot.json, the only difference is
the 20 Stone Guard's items from #4 (they move from rank-10 pvp to their
ungated vendor row). Both readings sit under the cap, so they stay
default-eligible. Full run == merge on the new code, items.json is unchanged,
and the Go CDN contract test passes on the regenerated file (checked with the
regen swapped in, then restored).

## 2. PvP rank labels (integrator's widened item 2)

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| 3 | Rank 5/10 rewards: the labels are right, so nothing is relabelled. | The client rows exist: 272694, 272449, 272695 (ItemSparse RequiredPVPRank 0, RequiredPVPMedal 0, BuyPrice 1g/1g/4g). Their classic twins 15196/15197/15198 carry rank 5/5/10. Forever's server text says "Requires Private/Scout (Rank 1)" on 272694 and 272449, and "Requires Knight / Stone Guard (Rank 6)" on 272695 (nether.wowhead.com/forever/tooltip/item/<id>, 2026-11-01). Client code 5 = displayed rank 1, 10 = rank 6. So these really are rank-locked, and `kind: pvp, rank: N` stays. ItemExtendedCost for 70707 (wago, 1,101 rows) does use Honor Points (CurrencyTypes 1792) on 23 costs. But the client has no item-to-cost link (vendor stock is server-side), so it cannot show which items cost honor. Single source for the rank (server tooltip; the client rank sits on the classic twin). | Aligned | - |
| 4 | Rank 10 was lopsided: 20 Horde armour pieces in pvp:rank-10:horde against one Alliance tabard. | `HONOR_RANK_BY_TITLE` ranked "Stone Guard's" (10) but leaves Alliance's "Knight's" out, because ordinary items also start with "Knight's" (classic world-drop set 7454-7465). So Lady Palanseer's ilvl 53 Stone Guard's sets (220798 ... 220906) were rank-locked, while Captain Dirgehammer's Knight's twins (220794 ...) were vendor-only. The server text on both says only "Requires Level 50" (220801, 220794, 220798, 220797). Fix: drop "Stone Guard's" from the table. Both sets stay on their vendors (vendor:12792, vendor:12777, both ungated). Stone Guard's Herald 272450 keeps rank 10 by re-itemisation from 15199 (RequiredPVPRank 10). Tests: every title-ranked rank covers both factions (failed: `{10: {'horde'}}`), and neither rank-10 armour name ranks. | CONFIRMED-BUG | S e1b402d61 |
| 7 | 467 listings of Forever PvP gear (ranks 9, 11, 12, 14, 16, 17 and 18; Forever ids only) were rank-locked only by their title. | Neither ItemSparse.RequiredPVPRank (no Forever id has one) nor classic-db (classic ids only) ranks them. Forever's server text on a sample (227198, 231630, 234555, 234584, 231681, 220870, 227122, 227147) shows only "Requires Level N". Owner ruling: where the client data does not show how an item is sold, label it "PvP reward" with no rank lock. Fix: `unlock_title_only_ranks` moves them to `pvp:title-rank-N:<side>`: kind `pvp`, `rank` absent, name "PvP reward (<Side>)", old rank in the new field **`title_rank`** ("Era honor rank the item's title implies; eligibility only, not a lock"). The step runs after supersession. Client- or classic-db-ranked ids, and ids re-itemised from them (the 4 tabards), keep `pvp:rank-N`. `title_rank` is added to models.py, the openapi LootSource, the CDN contract's optional fields and the regenerated Kotlin/Swift models. | CONFIRMED-BUG (owner ruling) | S 970ec9e90 |
| 5 | Vocabulary: no new kind or label string. | Kinds stay `pvp`/`vendor`. `faction_source` stays within `Literal["classic-db","title"]` (models.py:454). openapi has `faction_source: string`. No web, Go, Kotlin or Swift code branches on it. | Aligned | - |

## 3. Darkspear Islands (widened item 3)

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| 6 | No new source is needed. | Map.csv 2997 "Darkspear Islands", InstanceType 3, zones 16606-16610. No classic-db or fork spawn or loot is on map 2997 (0 records). Its rewards are already a source: `vendor:263570` Creeg Bothunk, 51 items, wowhead-sourced, `opens: later` from curated/content-phases.json `darkspear-islands`. That gate is not changed (integrator d: the owner decides). | Aligned | - |

## 4. `loot --allow-shrink`

It is still a no-op. It is kept (and now documented on `write_loot_files`)
because data.yml:292 passes `$SHRINK` to `loot`. Removing the flag would make
a re-baseline dispatch fail in argparse. client.yml:324's hint names data.yml's
`allow_shrink` input, which still matters to `normalize`, so the hint is
correct as written.

## Before/after

No sim or DPS change: the lane touches loot.json generation only. The loot
before/after is the table in section 1.

## For other lanes

- **Ranker:** after the next regen, Horde leveling picks of Stone Guard's
  armour (20 ids; cited in 20 bis/*.json, e.g. 4 in shaman-enhancement,
  hunter-survival and druid-feral-bear; 36 Data.lua labels) are sourced from
  vendor:12792 "Lady Palanseer" instead of "PvP rank 10 · Stone Guard ·
  Horde". The Alliance twins are already picked as vendor:12777 "Captain
  Dirgehammer" (42 citations), so eligibility should match across factions.
  Please confirm that no band's rank cap was what admitted them.
- **Ranker (required before the next regen ships):** read `title_rank` as the
  rank for `pvpRankExceedsCap` and in `vendorInheritsPvpRankGate`, i.e.
  effective rank = `rank` if set, else `title_rank`. Without that, the 467
  rank-11+ listings would quietly pass the cap. Label: "PvP reward ·
  <faction>" for a pvp source with no `rank`.
- **Web / Discord bot:** `web/src/lib/bis/source-cell.ts` and
  `api/internal/discordbot/item_source.go` fall back (no crash) for a pvp
  source without `rank`: the web shows the slot's own label, and the bot
  skips the PvP line or shows the vendor. A "PvP reward · <faction>" cell is
  a page or API follow-up.
- **Settled by the owner ruling (was an owner decision):** 491 Forever-id items sit in
  `pvp:rank-N` sources only because their title implies a rank (ranks 7-9 and
  11-18). Forever's server tooltips for a sample of 8 (227198, 231630, 234555,
  234584, 231681, 220870, 227122, 227147) show no rank requirement, only a
  level. That fits "Forever's PvP gear isn't rank-locked". Proposal: rank only
  what the client or the server text ranks, and label the rest "PvP reward".
  This changes eligibility, so the owner rules on it.
- **Owner decision:** content-phases `battleground-reputation` (AV/WSG/AB
  quartermasters, zone 2597, AV quests) and `darkspear-islands` are still
  phase `later`, even though battlegrounds are in at launch. Not flipped,
  per the integrator.
- **Integrator:** data.yml must drop `$SHRINK` from the `loot` step before the
  flag can go.

## Still open

- The honor-cost link for PvP items: the client has no item-to-ItemExtendedCost
  mapping. Settling it needs vendor stock (server or wowhead "sold by" with
  cost).
- The committed loot.json carries the section 1 table's differences until the
  held data/bis runs regenerate it.
