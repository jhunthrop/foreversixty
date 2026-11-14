# UX designer verdict: addon 2.0.49 shared Stat weights (Gear 4.3, Rotation 4.4)

Reviewed: gallery 04-window-gear.png and 05-window-rotation.png against the before captures and spec sections 4.3 and 4.4.

Verdict: HOLD (one must-fix, two should)

## Consistent across the tabs (confirmed on the captures)
- Same heading, same muted Stat / Weight / DPS per point head row, same row labels ("Hit rating", "Agility", "Crit rating", "Ranged attack power", "Strength"), same numbers (1.00 / 0.42 down to 0.07 / 0.03).
- Weight column bright, DPS per point column muted, both right aligned; the per point column sits on the block's right edge on both tabs (Gear 1836 with the Upgrades notes above it, Rotation 1345 with the Level stepper above it). Column gap scales with the window (148 vs 161 px on the plates, the same ratio as the row pitch).
- One attack power row (the old doubled "Attack Power" / "Ranged Attack Power" pair is gone). Title-case leftovers ("Critical Strike", "Hit", "Attack Power") are gone from Gear. Gear's old "per rating point" two-word column header is replaced by the spec's "DPS per point".
- Rotation drops the Intellect 0.03 row the old build showed; Gear never showed it. Both tabs now apply the same floor. Good.
- Gear hit line reads "Ranged hit: 2% to the 9% cap"; the spurious "Spell hit: 9% to the 16% cap" is gone. Correct for a ranged hunter.
- Narrow 288 px Gear column: nothing wraps or truncates; the longest label ("Ranged attack power") ends 79 px before the Weight column. Rotation 388 px column: block ends 35 px above the footer rule, clear of it.

## Must fix
1. Gear's Stat weights heading has no right note. Spec 4.3 says the note is always present: `your sim, per rating point` or `shipped, per rating point`. The before capture had "per rating point"; the after plate has nothing, so the unit and the source of the numbers are lost on the tab that prices gear. Pass the source note to the mock for the Gear plate (shipped weights: `shipped, per rating point`) and confirm the Lua call site passes the same string. Rotation is correct to show no note here: at Level 60 with the shipped band equal to the level shown the spec says none (`at level N` only when they differ, `your sim, 7 Dec` when a sim result exists). Re-shoot 04 after the fix.

## Should fix
1. Gear hit line spacing: "Ranged hit: 2% to the 9% cap" sits 25 px under the Strength row, a hair under the row pitch (27.6), so it reads as a seventh stat row. Give it the haste line's gap (LAYOUT.hasteGap, 2 px) plus 4 px, or a 6 px total top margin, so it is visibly a footnote. Keep its muted colour.
2. Add one acceptance capture of a spec with haste (a caster) on both tabs, and one with a capped stat marked (spec 4.4 "a capped stat marked"), so the "haste line apart" and cap marking are proven on the two column widths. Neither state appears in the hunter plates.

## Polish
- The Gear heading note, once restored, should align right at the same 1836 edge as the Upgrades "+21.2 DPS in all" note.

---

# Round 2

Reviewed: re-rendered 04-window-gear.png and spec 4.3 item 2 as updated 2026-11-14.

Verdict: SHIP

- Must-fix 1 withdrawn. The owner's rule (Rotation's block is the reference, one component, source-only note) is now written into spec 4.3, and the plate matches it: no heading note at level 60 with the shipped band equal to the band shown, the unit carried by "Hit rating", "Crit rating" and the "DPS per point" column, as on Rotation. Gear and Rotation now follow one note rule.
- Should-fix 1 closed. "Ranged hit: 2% to the 9% cap" sits 35 px under the Strength row (row pitch 27.6), a visible 7 px break, muted, reads as a footnote. Sources moved down about 10 px and its last row ends 70 px above the footer rule; nothing crowds or clips.
- Should-fix 2 accepted as proposed: haste and capped-stat states are covered by the busted specs and stay out of the store plates. Condition: the capped-stat mark and the haste line must be checked on a render the next time a caster or capped plate is shot.
