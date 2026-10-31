# UX verdict: logs engine counts, fight-list rows (round 2)

Verdict: SHIP

Reviewed: sample-fights-1440/390, folded-zone-row-1440, no-mobs-died-1440, old-format-no-count-1440. The "1440" captures are rail crops (600-708 px wide), not full-page. sample-page-* were not judged. The Everyone-in-range dagger state and the tooltip were not captured.

## Blockers

None.

## Shoulds

1. "your group killed 0" repeats on most trash rows (about 9 of 12 in the 390 capture). The zeros bury the rows where the group killed something. Dim the whole "· your group killed 0" clause, or drop it and keep "N mobs died". Keep "No mobs died" as is.
2. The count line is one grey run, the same colour and weight as the time. Make the two numbers the only brighter text (tertiary to secondary text token) so "177 ... 26" scans. The row title stays primary.
3. At 390, the folded zone row puts "79 pulls ‹" on its own right-aligned line, so the row is a line taller than the fixture folded row, where it sits beside the counts. Keep the same layout in both cases. Pin the pull count to the title line on the right.
4. Dagger call: keep no dagger when only outsiders died. It matches the "your group killed" line and the "Your group" scope default. Two conditions:
   - Switching the scope to "Everyone in range" must show the in-range daggers. Attach a capture of that state.
   - The tooltip is hover-only. On a phone the "3 among everyone in range" half is unreachable. Make the dagger tappable to show the tooltip, or accept that the scope toggle is the touch path and say so in the report.
5. Old-format fights: dropping the count line is right. Never show 0 or a dash. But the dagger there is the in-range count, while the page scope reads "Your group". Either word that dagger's tooltip "N deaths among everyone in range (this report predates group data)", or note it as a known mismatch. A reader who sees "6†" next to a "Your group" subtitle will compare them.
6. The header reads "0 pulls" above 79 trash pulls on the sample. It is technically boss pulls, but it reads as a bug. Label it "0 boss pulls" or hide it at 0.
7. The 390 sheet title "Trash · 12:43 · 177 mobs died · your group killed 26" fills the full width and nearly hits "Hide". It is fine now. A two-digit kill count plus "1 mob died" variants will truncate, so confirm it ellipsizes and does not wrap.

## Checks

- Hierarchy and density: pass with shoulds 1-3.
- Folded zone row: the sum is right (0 + 1 = 1) and dimmer than boss rows. Pass.
- Old-format case: pass, with should 5.
- Scope and wording consistency ("Your group / Everyone in range", "players seen"): pass.
