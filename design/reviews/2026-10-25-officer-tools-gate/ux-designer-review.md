# Officer tools gate: ux-designer review

Reviewer: ux-designer, 2026-10-25. Spec: `design/specs/2026-10-25-officer-tools-gate.md`.
Judged on the rendered captures in this folder and the `heights-*.json` numbers, never the diff.
Fixture names and dates (Kraggor awarding to Kraggor, "progression tools") are not defects.

## Verdict: SHIP WITH FIXES

Three must-fix items, all small. No blocker. Every state in §8 is captured and reads as specified.

## State by state

| # | State | Spec | What renders | Result |
|---|---|---|---|---|
| 1 | Loot, tools on | controls unchanged, no sentence, no label | Award bordered and bright; no `Read-only`; boss row at y 84 (1440 crop) | pass |
| 2 | Loot, officer-read-only | §3.2 sentence under header row, above boss picker; §3.3 label; §3.4 dimmed Award | sentence 13 px `--text`, gold underlined link; `Read-only` muted at h2 right edge; every Award in place, same size, ~50 % opacity; boss row moves 84 to 120 (36 px) | pass, see (a) |
| 3 | Loot, read-only, empty | sentence above `No bosses killed yet…` | exactly that, one line at 390 | pass |
| 4 | Loot, awarded item | Undo disabled, `Awarded · Undo` muted, no underline | as specified at 390 and 1440 | pass |
| 5 | Loot, member | no sentence, no controls, member `Read-only` | as before; no sentence | pass |
| 6 | Loot after 403 | one line under the clicked card's rows; no bottom error line | `Not awarded. This guild's officer tools have ended.` under Stormrage Cover's rows; both Awards dimmed; sentence shown at top | pass |
| 7 | Readiness, read-only | sentence under h2, above column header / first card; label; Nudge dimmed | as specified at 360, 390, 1440, 2000; every Nudge in place and dimmed; 12 px gap held | pass, see (a) |
| 8 | Readiness, member | no sentence, no Nudge, no label | as specified | pass |
| 9 | Focus-visible on link | 2 px gold ring at 2 px offset around the link | 1440: ring hugs the text. 390: ring is a 44 px box over the boss picker's top edge | must fix, see (b) |
| 10 | Settings, scoped, claim holder | §5.1 rows 1-4 in order | plan line 14 px `--text`; explanation 13 px muted; price line; gold `Subscribe the guild` button `w-fit`; no Manage billing | pass |
| 11 | Settings, scoped, no claim | §5.2: plan line, explanation, `guildMasterOnly` line | exactly that; no button | pass |
| 12 | Settings, granted (regression) | `Guild plan · until 31 Dec 2026` unchanged | unchanged | pass |
| 13 | Account, Guild panel | §6 muted 13 px line, muted inline link `Guild settings` | line placement, copy and single-line fit at 360 are right; the link is gold | must fix, see (c) |

## Heights (§3.5)

All eight numbers equal or sit under the cap. 360: Loot 56/56, Readiness 52/52, sentence 40 px on 2 lines.
390: 36 and 32, one line. 1440 and 2000: 36/36 and 32/32. Nothing else moved: the 01 vs 02 crops differ
only by the 36 px shift. Pass.

## 403 path (§4)

One line, in the card that was clicked, the tab-bottom `guild-loot-error` absent. The line has no colour
class; that matches the tab's error-line recipe, which also has none. Accepted as specified.
No visible ring on the programmatically focused line: accepted. `tabindex="-1"` focus shows no ring by
`:focus-visible` rules, the line is not interactive, and `role="alert"` announces it.

## Rulings on the capturer's notes

**(a) Link breaks as `Subscribe the` / `guild.` at 360. Must fix.** Add `whitespace-nowrap` to the link.
A split link draws two underlined fragments; with the 44 px hit area the first fragment's target also
overlaps line 1. With `nowrap` the break moves to after `plan.`: line 1 `Award and Nudge are part of the
guild plan.` (about 270 px at 13 px Barlow, inside 324), line 2 `Subscribe the guild.`. Still 2 lines,
still 40 px, still 56/52 at 360. **The cap is not exceeded.** 390 and up are unchanged (one line). The
implementer re-records `heights-*-360.json` and confirms 56 and 52.

**(b) Phone focus ring is a 44 px box over the boss picker. Must fix.** The ring must mark the 20 px line,
not the hit area. Remove `py-3` from the link. Achieve the 44 px target with a pseudo-element:
`relative` on the link, plus `after:absolute after:inset-x-0 after:-top-3 after:-bottom-3 after:content-['']`
(12 px above and below a 20 px line = 44 px), phone only (`md:after:hidden` or the same breakpoint `py-3` used).
The `focus-visible` outline stays on the link itself, so it hugs the text exactly as at 1440.
Do not use an inner span: it puts the ring on a non-focusable element and needs `:focus-visible` on the
parent plus a descendant selector for no gain. Recapture 09 at 390; the ring must not touch the boss picker.

**(c) Account row `Guild settings` link is gold. Must fix; confirmed.** §6: muted inline link.
At rest `text-muted underline`, `hover:text-strong active:text-strong`, 2 px gold `focus-visible` ring
at 2 px offset. Set the colour on the element; do not rely on the global `a` rule. Recapture 13 at 360,
390 and 1440, plus one focused state at 1440.

**(d) 2000 px captures. Pass.** The guild page container is centred at 2000 (wide-screen rule). The 02 and
07 panel crops at 2000 are pixel-identical in layout to 1440: sentence one line, 20 px, label at the right
edge, adds 36 and 32. No stretched row, no orphaned control.

## Other notes (no change required)

- Disabled Undo reads as a plain muted label beside `Already awarded`. That is the intent (§3.4): the
  greyed action button stays in its slot. The `Undo` word stays. Accepted.
- The `GuildViewerView.officer_tools` doc comment in `web/src/lib/guild/api.ts` reads inverted. Not a
  render issue; the implementer should correct it in the same fix commit.
- `notOnPlanLines().detail` still says "progression tools". Out of lane per §7; already flagged.
- The capturer noted no mock side-by-side. None is required here: this lane adds one line and a label to
  existing regions and no new region; the 01 vs 02 and 05 vs 02 pairs are the comparison.

## Checklist

- [x] §1 answer visible without scrolling at 360, 390, 1440, 2000 (inside the panel's first 160 px)
- [x] §3.2 sentence: copy verbatim, 13/20 `--text`, gold underlined link, placement on both tabs and above the empty state
- [ ] §3.2 link stays whole at 360 (fix a)
- [x] §3.3 `Read-only` label on both tabs in officer-read-only, absent for officer with tools
- [x] §3.4 Award, Undo, Nudge keep position, label, size; dimmed; no hover
- [x] §3.5 heights at or under cap at every width
- [x] §4 refused line in the clicked card, one message, sentence shown
- [x] §5.1 scoped settings, four rows in order
- [x] §5.2 officer without claim sees plan line, explanation, guild-master line
- [x] §5.1 granted regression unchanged
- [x] §6 account row placement, copy, one line at 360
- [ ] §6 account link muted (fix c)
- [ ] #9 focus ring hugs the link at every width (fix b)
- [x] 2000 px centred, no stretch

## Fixes, in one list

1. Capture 02/07 at 360: `whitespace-nowrap` on the gate link. Heights stay 56/52.
2. Capture 09 at 390: drop `py-3`; 44 px hit area via a 12 px top/bottom `::after`; ring on the link.
3. Capture 13 at 360/390/1440: `Guild settings` muted at rest, `--text-strong` on hover and active.

Ship once 02-360, 07-360, 09-390 and 13 (three widths plus focus) are recaptured and match the above.
