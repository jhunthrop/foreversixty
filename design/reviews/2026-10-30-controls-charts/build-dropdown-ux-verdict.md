# Build verdict: shared dropdown (UX designer)

Spec: design/specs/2026-10-30-controls-charts.md Part A. Judged on the build captures and the lane report.

**Verdict: HOLD**

## What is right
- The open list reads as the /account family: card surface, 12 px radius, gold inset top edge, tracked group heads, gold check, gold left edge on the active row, scroll fade (report Source, sim Fight style, planner Race, rankings Region, More menu).
- Report Source: crest plus class colour, role groups, 36 px field with 13 px rows, gold open glow. Matches dropdowns-open-1440.
- Fight style: descriptions never clipped, check and active state read as two signals.
- More menu moved onto the same surface and row look.
- Closed field: 8 px radius, chevron, open and error/loading/empty/placeholder states in the sheet at 52/44/36 match the states board.
- Phone: closed fields are the native control at 44 px with the same face (report Source, Table).

## Blockers
1. Rankings filter bar truncates the picked value at 1440 and 2000: "Every cl...", "Every s...", "All rules...", "All regio...", "No encounter...". Spec A9.3: the picked value never collapses. At 2000 there is spare width and it still truncates. Fix: size the nine columns to content (minimum the longest option label), let the bar wrap to a second row before it truncates. Re-capture at 1440, 2000 and 1024.
2. No build capture shows an open searchable list or an open multi-select list (filter row, count, gold match, sticky group head, no-match row, checkbox rows, foot row, "Top 8 of N"). The captures list only plain and icon lists; the sheet shows closed fields. These variants cannot be passed on the lane report alone. Add open captures of: report Source with 15+ players, sim Find an item Source typed "black", no-match, the Players multi-select, and a flip-above case, at 1440 and 390 (390 shows the coarse behaviour).

## Shoulds
1. Merge condition: the report still shows native, unstyled selects (Window in the chart head, Enemy under the threat chart) next to the new fields. They are the "basic html" the owner named. They are the chart lane's, so this lane may land, but the report must not go out to the owner until the chart lane converts them.
2. Report Source list: the last row "Elyra Duskvale · Hardcore" has no crest and no class colour, unlike the rows above. Give it the crest and colour, or explain it as a non-roster entry.
3. The More trigger shows no open state while its menu is open (no gold border or pressed look). Give it the same open treatment as the fields.
4. Rankings at 390 collapses to a Filters sheet; no capture shows the fields inside it. Capture the sheet open and one field open in it.
5. Planner capture: the Class field crest and gold colour are right, but the talent tree icons beside it are broken images. Not this lane's change; flag to the planner owner.
6. Home-search capture shows the closed fields only; add Region open.

## Lane deviations
- No `title` tooltip: acceptable. Hover-only titles are banned by repo rule and the accessible name carries the label. It becomes a real loss only while blocker 1 stands; fixing blocker 1 closes it.
- Raid form selects at 44 px: acceptable. They align with the 44 px inputs in the same grid. Seat and transfer pickers stay 52.
- No Enemies group in Source: acceptable. The `All enemies` scope is kept and the roster has no enemy rows. Reinstate the group when the roster carries enemies.
- Coarse-pointer fields at 44 px on a touch laptop: acceptable. It matches A9.2 and the 44 px touch target rule.

## Re-check (fix round 1)

**Verdict: SHIP**

Blockers
1. Rankings truncation: resolved. "Every class", "Every spec", "No encounters yet", "All rulesets", "All regions" read whole at 1024 (bar wraps to a second row), 1440 and 2000. At 390 the fields sit in the Filters panel, two per row, nothing truncated.
2. Open-state captures: resolved. Seen and correct: searchable report Source with the 40-player pull (filter row, tracked group heads, crests, gold check), sim Source typed "black" (1 match, gold match text, group head), no-match ("Nothing matches "naxxblack"." with 0 matches), Players multi-select (Compare group with descriptions, checkbox rows, figures, foot presets "Top 4 + bottom 4 / Top 8 / Bottom 8 / Custom", "10 drawn"), flip-above (Enemy list opens upward, ×8 meta), rankings Class list with crests and class colours.

Chart controls: Window and Enemy now use the shared dropdown face (dense, 36 px, chevron); Players field reads "Top 4, bottom 4 of 40". No regression seen. Shoulds 1 and 2 closed; 3 and 4 closed.

Remaining polish, not blocking
- The sim Source list opens above its field at 1440 and is narrower than the field (about 280 px against 240 px field plus clipped right edge of the "Only items" label beside it). Fine as shipped; watch that it never covers its own field.
- Players list is cut mid-row at the foot by the scroll fade: intended.
- Planner broken talent icons remain the planner owner's.
