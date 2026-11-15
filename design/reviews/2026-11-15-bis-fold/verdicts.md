# BiS fold: persona verdicts (2026-11-15)

The 30-39, 40-49 and 50-59 BiS bands now sit on /bis/<class>/<spec> as a server-rendered "Leveling 30 to 59" section, and their old URLs 301 to `#levels-<band>`. The captures in build/ come from a production build of real data (capture.mjs), at 1440 and 390, for fury, frost, fury Horde and fury #levels-40-49.

## Round 1: all four FIX
- **ux-designer, ai-slop-reviewer, wow-player and info-density-reviewer:** the empty trinket line "Nothing here either until 40." read wrong with nothing above it.
- **All four:** "Picks for a Night Elf (Alliance)." repeated in every band column.
- **wow-player and info-density-reviewer:** blank icons at 390. This was a capture artifact from lazy loading; the capture now walks the page first.
- **wow-player:** check that Barbaric Iron Helm is mail. It is: item 7915 is class 4, subclass 3.

Fixes (18e45febe):
- The shared later-empty line is now "None until 40." / "None at this level.".
- GuideBisPanel takes `caption={false}` in the fold. The guides keep their caption.

## Round 2: all four SHIP
- **ux-designer:** SHIP.
- **wow-player:** SHIP.
- **ai-slop-reviewer:** SHIP. The dark empty socket is accepted as the shared ItemRow empty state.
- **info-density-reviewer:** SHIP.

## Open shoulds (follow-ups, not blockers)
- Mark the band column that matches the doll's band.
- On phones, a jump row or collapse for the three bands (the page is about 7,100 px tall).
- Merge two empty trinket rows into one line.
- Show the required level on each row, a DPS figure per band, runner-up deltas and a carry-over marker.
- Link to the level 60 page from the section.
