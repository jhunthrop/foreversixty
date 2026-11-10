# Round 8 UX verdicts: sim pages (1440 + 390)

Handed-off items (h-doll ilvl rows, gear stat strip, SimTimeline toggle label) and fixture artefacts (blank healer icons, Orc/Troll mismatch, "0 valid combinations") do not fail a page.

| Page | Verdict |
|---|---|
| sim-landing-out | SHIP |
| sim-landing-in | SHIP |
| sim-result | SHIP |
| sim-saved | SHIP |
| sim-gear | SHIP |
| sim-talents | NOT SHIPPABLE |
| sim-drops | SHIP |
| sim-weights | SHIP |
| sim-specs | SHIP |

## Blocker: sim-talents (1440 and 390)
- Element: the "Guide build, level 60 · 17/32/2" row in Talent builds. The per-talent diff list ("+5 Unbridled Wrath", "+3 Improved Heroic Strike", "-5 Booming Voice", "-1 Piercing Howl" ...) prints on top of the three mini talent trees. Spec labels (Arms 17, Fury 32, Protection 2), icons and rank badges are overprinted; both the trees and the list are unreadable.
- Seen at 1440 (y 1040-1500) and 390 (y 1240-1700). "Your current build" above is clean, so the fault is the diff column only.
- Fix: flow layout, no overlap. 1440: trees left, diff list in its own column to the right (the plate has 800 px empty). 390: diff list under the trees. Group as Gains / Drops.
- Must-fix: the page's one job (compare two builds) fails.

## Notes, not failures
- sim-talents 1440: trees are 90 px wide in a 1240 px plate; fine once the diff sits beside them.
- sim-result 390: Talents, Rage, Buffs, Rotation fold; This run stays open for sim-phone.
