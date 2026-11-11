# Info density, lane G2 (home foot, /rankings snapshot)

Reviewed home and /rankings at 1440 and 390, plus the no-JS shots. Ignored the feed load error and the empty "Leaders by spec" as briefed.

## Jobs
- Home routes a player to the right tool. The foot is search copy plus a spec index.
- /rankings ranks logged kills. The no-JS snapshot gives three encounters with three rows each, so it is real text for a crawler.

## What works
- The "Leveling BiS by spec" index covers every spec by role in class colour. It is dense and holds on phone, where it wraps and does not collapse.
- Every snapshot row carries rank, crest, name and class, DPS, percentile (colour-banded) and date. It matches the JS table and survives at 390, where percentile stacks under DPS.
- Every prose link goes to the page that owns the detail, with boss links to `/rankings/<slug>`.

## Findings (all achievable with data we already have)
1. **Dead space at 1440 on /rankings.** The board is 880 px wide and the right 370 px is empty on both the snapshot and the boss page. Fix: add a rail with the fields the rows already carry. For the boss page that is the fastest kill length, kills ranked, and the top spec if it is present. Link the other encounters from it. Source: `web/src/lib/rankings/snapshot.ts` rows (length, date). On phone it folds below the table.
2. **Empty "Leaders by spec" card.** It is a heading plus one grey line, which the brief already marks as ignorable. Once the live rows carry no spec, hide the whole card instead of showing the empty line.
3. **"2 ranked" above three rows.** The count reads as the number of rows and contradicts what sits under it. Label it as kills, or show the real count. If it is a kills count, "2 kills" would be unambiguous.
4. **"Executed" column is all em-dashes on the boss page.** Drop a column that is empty for every row. The Length column carries the same width.
5. **Zero-DPS rows (ranks 6 to 10, 0 DPS with an empty bar).** If DPS is 0 for rows with a 10 to 50 percentile, the figures are suspect or incomplete. Check the data (sync-rankings.mjs) before it ships.
6. **Home foot is uneven.** "What's here" ends about 120 px above the BiS index on the left. The feed card also has about 110 px of dead space under its last row. Balance the two columns, or put the data-build line in the footer strip.
7. **The raid paragraph repeats the First raids card** (date and the three raid names, shown twice on screen). That is acceptable for search, but keep it as the crawler-facing text only, or drop the raid names and keep the "most logged bosses" sentence, which is new information.
8. **The BiS index is names only.** Optionally add the spec crest or the role-count per band. It is not needed to ship.

## Needs data
- Spec per live row (blocks "Leaders by spec" and spec-level context in the snapshot).
- Kill counts per boss in the home facts (`about-load.ts` returns name and slug only), so home could show "Relic Guardian 6 kills".

## Verdict
PASS WITH FIXES (1 rail or reflow for the empty right side of /rankings, 2 hide the empty Leaders card, 3 "ranked" label, 4 drop the all-dash Executed column, 5 verify zero-DPS rows, 6 balance the home foot columns)

## Round 2

Checked home-1440-nojs, rankings-1440-nojs and rankings-saltspine-1440.

- Fixed: the home foot is two equal halves ending together at the same y, with no dead gap. "N kills" now labels the encounter rows. The no-JS snapshot shows 3 nonzero rows per boss, so the zero-DPS rows are gone from the crawler text. The raid sentence now carries the "most logged dungeon bosses" links, which is new information.
- Not in this lane, accepted: the rankings right-rail space (about 370 px empty at 1440), the all-dash Executed column, and the feed panel height. The skeleton rows under Latest leave the panel mostly empty, which belongs to the feed.
- Still open, API follow-up: the live board (rankings-saltspine-1440) still shows ranks 6 to 10 at 0 DPS with an empty bar, and that is visible to players. It must not ship as the final board.
- Leaders by spec is hidden on the live board, as intended.

Verdict: PASS WITH FIXES (zero-DPS rows on the live board via the API follow-up; Executed column and the rankings rail, carried to the board lane)

## Round 3

Checked rankings-saltspine-1440 and rankings-saltspine-1440-nojs.

- Fixed: the live board and the no-JS snapshot agree. Both read "5 ranked" over 5 rows, and no 0-DPS rows remain.
- Follow-ups for the board lane, not this lane: the all-dash Executed column, and the empty right rail at 1440 (about 370 px).
- Observation: the no-JS shot shows "Filters (1)" while the JS shot shows "Filters". Check that this is intended.

Verdict: PASS (follow-ups for the board lane: Executed column, rankings rail)
