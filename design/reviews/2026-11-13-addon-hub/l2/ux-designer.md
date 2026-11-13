## Round 1

Verdict: FIX (two small items, then SHIP). Reviewed on the six renders at 740 px against hub-me.png and hub-empty.png boards 1-3, region by region.

What matches the boards: header (crest, name, figures line, X), tab strip with gold dots, Forever Score block (24 px tier-coloured figure, tier word, `316 to Epic (1500)`, five bars on the 700 scale in quality colours, `of` figures right-aligned, `2 flasks missed` on the lowest part), Last raid table (column widths, parse colours, died in warning, `No flask` warning, median/death/best/worst line, gold `Full log`), Right now rows (icon, two lines, gold tab link right-aligned), Your sim, footer states (`Companion synced 14 min ago`, `No companion`, `No ratings pack · companion synced 14 min ago`, `Ratings 8 days old` in warning), stale `as of 2 Dec` in warning, guild-less header and strip, Sim again / Queued states. Empty states leave out the right sections as spec 8 says.

Departures judged:
- Title-face section headings: accepted. One look across tabs beats the board's sans; amend spec 3.5 (it says Marcellus only for the window title) so the next tab lanes do not undo it.
- `Raid-ready list 984.0` omitted: accepted, no faked figure. Re-add when Data.lua carries list DPS.
- `Biggest gains` on its own line: accepted; the three icon + gain pairs fit 288 px on one row and read better than the board's wrapped `+11.2 / DPS`.
- Right now stand-in lines from the other tabs' `lead()`: accepted until L3-L5 land; the re-review must check they still read the same.
- `Not in a guild.` not drawn on Me: accepted, it follows hub-empty board 3, which is the approved board. Amend spec 8 to drop it.

Must fix:
1. Builds row copy. Shipped `51 of 51 on plan`; spec 4.1 and both boards say `51 of 51 points on plan`. Restore the verbatim words in `BuildDialog.lead` (and the stale/no-companion renders follow).
2. `me-new-character.png`: the `Pick your spec` row has no icon, so the text sits in an empty 56 px gutter with the gold link far right. Give it an icon (the class crest or the spec-choice glyph the Builds tab uses) at the same 32 px as other rows, or left-align the text at the column edge. Spec 4.1: every Right now row is an icon, two lines and the tab link. Also give it a second line (muted, e.g. the specs it chooses between) so the row keeps the two-line rhythm.

Polish (optional):
3. Section heading to first content gap is about 4 px tighter than the board (Right now, Your sim, Forever Score: about 22-25 px vs 27-28 px). Add 4 px under each heading so the larger title face breathes.
4. Biggest-gains icons show a stray 1 px tick at the top-left corner of each quality edge (Your sim row, all three). Check the edge texture crop.
5. Me's first-line rows in Right now sit directly under the heading with no gap, while rows are spaced 12 px apart; match the row spacing.

Checklist: no-scroll answer on first paint ticked; board layout ticked; copy verbatim not ticked (item 1); every state captured ticked (raid day, no pack, no companion, stale, new character, queued); hover/focus/tooltip states not in the renders (static), verify in-game.

## Round 2

Verdict: SHIP. Checked the regenerated me-raid-day, me-sim-queued and me-new-character renders.

1. `51 of 51 points on plan`: fixed, verbatim.
2. New character (level 12): `Pick your spec` now starts at the column edge in client gold with its `Builds` link; no empty gutter. Forever Score reads `No logs yet.` and Last raid keeps the app line, both per spec 8. Accepted as one-line row; it needs no icon.
3. Heading gaps: Forever Score, Last raid, Right now and Your sim now breathe like the board.
4. Gain icon tick: gone.
5. Right now rows sit under the heading at the row rhythm.
6. Biggest gains as three rows (icon, quality-coloured item name, right-aligned `+11.2 DPS`) is better than the board: names are readable and the tooltip target is larger. Accepted. `Simmed Mon 7 Dec` is clearer than the bare date. `Queued` stays muted beside it.

Notes, not blocking: the Rotation stand-in's second line now reads `Tue 8 Dec · parse 58` instead of the board's `Your weakest boss on Tuesday`; fine as a stand-in, and L4's `lead()` must settle the final words (re-check on that lane). Hover, focus and item-tooltip states are not visible in static renders; verify in game.
