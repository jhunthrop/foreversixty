# Tank data follow-ups, 2026-10-08

Branch `tank-data` in both repositories (site and engine fork). Nothing was
pushed or merged. Follow-ups 1 to 3 of the Protection paladin tank audit
(`2026-10-08-paladin-tank-audit.md`, sections 2 and "Follow-ups").

## 1. Block value is its own stat

Cause, as the audit found it: `STAT_BY_MODIFIER_ID` filed both
`ITEM_MOD_BLOCK_RATING` (15) and `ITEM_MOD_BLOCK_VALUE` (48) under `block`;
the simdb conversion divides `block` by the rating factor (5) and the engine
applies it as block chance. The Immovable Object's 27 block value became
+5.4% block chance and no block value.

Fix:

- `pipeline/normalize/gear.py`: modifier 48 maps to `block_value`. Rating 15
  stays `block`.
- `pipeline/classicdb_items.py`: aura 158 (`MOD_SHIELD_BLOCKVALUE`, the flat
  "Block Value NN" equip line the supplement used to ignore) maps to
  `block_value`, with its tooltip label.
- Nothing in `pipeline/simdb` needed to change: `statmap` already maps
  `block_value` to the engine's `StatBlockValue`, and the rating conversion
  only divides the rating family (`block` is in it, `block_value` is not).
- Web: `block_value` joins `STAT_KEYS` / `STAT_LABELS` (and the data suite's
  transcribed list); the enchant stat type no longer names it twice.
- Ranker: `paladin-protection` and `warrior-protection` weigh `block_value`
  beside `block` (`data/curated/specs.json`, regenerated `specs.go` and
  `specs.ts`). The weights API already accepted the stat.
- Engine: no change. The engine does not compute block chance from the shield;
  every character gets a base 5% `Block` (`character.go`), `CanBlock` is set
  by a shield in the off hand, and block value is the stat plus strength.
  Giving a shield block value therefore does not double count anything.

Tests, written first and failing before the fix: a shield row with modifier 48
gives `block_value` and no `block` (and 15 still gives `block`), the classic-db
aura 158 case, a simdb `SimItem` with `StatBlockValue` and no `StatBlock`,
the ranker keeps `block_value` flat, and The Immovable Object scores 27 block
value and no block chance through the ranker's reader. The simdb reader test
skips while the committed `simdb.bin` predates the split (the nightly
regeneration refreshes it; `simdb.bin` is not committed here).

Regeneration. `python -m pipeline normalize` run on this machine does not
reproduce the committed `items/*.json`: the raw wowhead and classic-db
payloads that the nightly fetches are not in the checkout, so a plain
regeneration rewrote about 32,000 lines (new classic-db placements, icon
origins) unrelated to this change. Committing that was out of scope. The
committed files instead carry only the block change, taken from that
pipeline run item by item: for each item whose regenerated row has
`block_value`, the committed row gets the regenerated `block` and
`block_value` and nothing else. The diff is 185 lines added and 114 removed
over nine class files and touches no other key.

What moved, counted over the class files (an item appears in several): 94 rows
had their summed `block` changed (46 plain moves of the whole amount to
`block_value`, for example The Immovable Object 27 and Earthen
Guard 12 now `block_value`; 48 keep a smaller `block` rating beside a
`block_value`, for example Gluth's Missing Collar and the Dreadnaught and
Inquisition pieces), and 20 rows that had lost their block value gained it
(classic-db shields such as Aegis of the Blood God and Barrier Shield). Not
regenerated, nightly owned: `bis/*.json`, `loot.json`, `Data.lua`,
`addon-data.json`, `simdb.bin`, `manifest.json`, `simitems.json`. Two data
tests that compare the addon files with the items (`test_addondata`,
`test_audit_addon`) fail until the nightly `addon-data` step runs.

The tank ranking itself was not re-run. The audit priced the phantom 5.4% at
about 1% of damage taken, and the nightly bis run will rank the shields on
block value from now on.

## 2. Benediction

The client states 2% a rank, 10% at five (talent text; spell 20101 is one row
whose cost modifier is -10). The engine took 3% a rank, 15%. It is now
`100 - 2 * rank`, with tests against the client row and by rank. No golden
moves: neither the Retribution nor the Protection preset takes the talent, and
`conformance` is unchanged.

## 3. The band ladder reaches the build's deepest talent

`leveling.LadderTalentString` spent a build top row down, so the wide middle
rows used the points before Holy Shield. It now reserves the guide build's
deepest talent and its prerequisite chain when the band can legally carry it
(the tier gate is 5 points a tier above, the same rule the planner enforces),
then walks top down as before. Two guards keep it from changing anything else:
it does nothing if the plain walk already reaches that talent, and it is
dropped for a band where paying for it would take a talent entirely away from
the plain walk's build.

Protection paladin: band 40 now has Templar's Bulwark and Holy Shield
(`5530513321101001`, One-Handed Weapon Specialization 3 to 1) and no unresolved
ability. Band 38 still has neither, and cannot: Holy Shield is a tier 6 talent
and needs 30 points above it plus its own, 31, and band 38 has 29 (it is also
first learned at level 40). Templar's Bulwark at 38 would need a second rule
that reserves partial paths; trying that moved bands 20 and 30 of most specs
and dropped DPS, so it was left out.

Ladder goldens regenerated, every diff read: 12 specs move, each by taking its
deepest talent at band 30 to 50 in place of the last ranks of the plain walk.
No band loses a talent (checked digit by digit: no digit goes from nonzero to
zero in any row). Bands that gain: druid-balance 40, druid-feral-bear 50,
druid-restoration 30, mage-fire 40, paladin-protection 40, priest-discipline
38, rogue-assassination 40, shaman-enhancement 40, shaman-restoration 50,
warlock-affliction 40, warlock-destruction 40, warrior-fury 40. Warlock
affliction 40 gains most (63.9 to 103.1 DPS); paladin-protection 40 (35.5 to
33.6), rogue-assassination 40, warlock-destruction 40 and warrior-fury 40 trade
a little DPS for the deepest talent. The
duplicate walk in `request/ladder.go` (`ladderTalentPoints`) is gone; the test
reads the ladder's own string through `TalentRanksFromString`.

## Verification

Fork, `go test --tags=with_db ./sim/paladin/... ./sim/core/ ./sim/conformance/`:
pass. Site `sim/`, `go test ./...`: pass. `data`, `uv run pytest tests -q
--no-cov`: 1423 pass, 2 fail (the addon drift pair above). Web `vitest` on
`lib/items`, `lib/planner`, `lib/sim`, `lib/addon`: 922 pass.
