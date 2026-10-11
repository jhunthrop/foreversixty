# Shared character header, Rankings and Logs: WoW player review (2026-10-11, round 1)

Specs: design/specs/2026-10-11-{character-header,rankings,logs}.md.
Renders: design/mocks/renders/{character-header-states,rankings*,logs*}.png.
Canvas: https://claude.ai/artifact/7yYKwphH3s7BdkJAHCQBmA

**Verdicts:** header SHIP WITH FIXES (NOT YET only on the home change);
Rankings NOT YET; Logs NOT YET.

## Header
One band (identity, answer, one action) is the right shape; the states are
right; a header switch swaps identity in place with no second switcher.
1. Home: the 34 px name becomes 24 px and the Best in slot card folds into
   the band with "+11.2 DPS together" as a tiny footnote; the DPS gain is
   what a player repeats, so it needs figure size (owner call).
2. Labels collide: "BEST IN SLOT" (home, upgrades to get) vs "BEST-IN-SLOT
   GEAR YOU WEAR" (BiS, slots worn) under near-identical labels.
3. Stale line wraps to three lines and makes the identity column taller;
   "/reload" is wrong for players; trim to "Log out in the game to sync".
4. 390 loading skeleton is taller than the ready box.
5. Signed out renders nothing, so every page has a different signed-out
   front, which is where new players land.
6. Descriptor lacks race ("Human Fury Warrior"); the 16 px faction logo is
   hard to read as a faction.
7. "CURRENT CHARACTER" / "YOUR CHARACTER" changes meaning when an alt is
   added.
8. The planner/BiS-spec card is a third look; within "same(ish)".

## Rankings
1. "14th of 212 Fury Warriors" is false while the API ranks kills; the mock
   shows the confident wording; ship "kills" wording at the same size until
   per-character best kill and own-rank lookups exist.
2. The pinned row repeats directly above the same row at 14; pin only when
   the own row is off the visible page.
3. The purple "14th" follows the WCL ladder by rank percentile; keep, with a
   legend.
4. Encounter-first is right; Region not realm is right for the data; seven
   controls plus a checkbox, encounter the same width as Faction (1024 is
   better).
5. "Another class or spec" is not a second switcher; the wording is clear.
6. SIZE 5 and LENGTH 4:12 on every row look fake.
7. BUILD "31/20/0" and Report are both gold links per row.
8. Phone: the board starts four screens down; collapse the filters to one
   summary line on phone.
9. Signed out: the page adds a second sign-in in prose beside the selector.
10. The heading repeats "212 … ranked" from the header.

## Logs
Latest parse is the latest ranked kill: nothing on screen says "kill", so a
wipe night reads as a stale header (largest risk).
1. The DPS figure is a hook, not the page's answer.
2. Three upload buttons in one screen (header gold, "or choose a file", the
   outline Upload before any file is chosen).
3. Companion below Upload though it is the primary route for live loggers.
4. Phone: a 48 px gold "Upload a log" for a file that lives on a PC; the
   phone primary should be "Open your latest report".
5. The sample report row duplicates the own-report row (same dungeon, date).
6. ISO dates in mono vs "Nov 11" on Rankings.
7. Report rows show "63 fights · 12 kills" with no zone or boss names.
8. 1024: "Open the fight →" wraps, the arrow orphaned.
9. Signed out: two identical sign-in buttons plus the selector's.
10. Gear and sim hand-offs gone from the landing: right.

Questions for the architect are in the transcript summary; the decisions
taken for round two are in the next section of the spec.

## Round 2

**Header: SHIP.** Fixed: figure sizes, distinct labels, stale line, loading
height, race, "Your character". Partly: signed-out band everywhere except the
home front door (accepted); faction logo still 16 px (the word carries it).
New: the 390 descriptor wraps to a lone "·" on its own line before the
faction (rankings-phone, logs-phone). Carry forward: the BiS landing must
re-measure its first card row with the 44 px signed-out band.

**Rankings: SHIP**, holding the merge for: the pinned off-page example (rank
204) contradicting the header (14th); Rankings "2 days ago" vs Logs "3 days
ago" for the same kill; "Raids, phas..." truncating in the Phase select and
the phone summary. Also: Onyxia with 10 and 20 sizes looks odd; class icons
on signed-out rows only; "kills" wraps at 1024 (accepted).

**Logs: SHIP.** Watch: "Last report 1 day ago" is unlinked; signed-out panel
order is the reverse of the paired order; the sample report appears twice.
