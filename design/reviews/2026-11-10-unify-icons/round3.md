# Unify icons, review round 3 (2026-11-10): SHIP

Captures:
- `before/` at `e4de576d2`, the lane's base, retaken with the current capture tool.
- `after/` at `ca0d5f0c9`.

Both sets serve every game icon from the real build in `data/builds`
(`tests/e2e/support/real-icons.ts`). The fixture build ships most icons as 1x1 black stubs, and the
new frame showed those stubs as the empty boxes rounds 1 and 2 flagged. Production has the real art.

| Reviewer | Verdict |
|---|---|
| ux-designer | SHIP |
| wow-player | SHIP |
| ai-slop-reviewer | PASS (optional fixes only) |

## Verified

- **home:** the Protection Warrior feed row shows its icon in the same frame as its siblings. The
  raids keep the plain frame, so a raid never looks like a legendary item. The "First raids open"
  row matches its before capture (round 2's dimming was the fresh-row animation).
- **bis-landing at 1440 and 390:** all 27 spec rows have their art. Row heights and left edges are
  identical before and after.
- **guides-class and guides-spec:** the spec tiles, the race portraits and the spec tabs are framed
  and aligned. The talent grid is untouched.
- **coach:** the Bloodthirst icon is framed. The chips and the "Rage full" heading draw no icon. The
  inline glyphs are unframed.
- **Newly read by ai-slop-reviewer, no tells:** sim-gear, sim-result, bis, character, report-deaths,
  report-damage-done, guild.

## After the verdicts

`30c2e9ee2`: a failed icon image now keeps its box, invisible, instead of removing it. The full
Playwright run found that removing the box reflowed rows when a fixture icon 404'd late (the guides
parity states at 390). Nothing visible changes: an empty frame still never shows, and no icon failed
in the round 3 captures.

## Open, non-blocking

- ux-designer: confirm 16 px as the frame cut-off. The 20 px chip icons keep the frame and the 16 px
  inline glyphs drop it.
- ux-designer: the report chart's lane markers (a SpellIcon at 18 px inside the marker's coloured
  face) appear in no capture. The sample report's chart carries no event markers. Check them on the
  next capture that includes a fight with deaths or casts on the lanes.
- ai-slop-reviewer: the guild raid nights repeat one zone icon on every row of a one-zone tier. It is
  accurate; consider keeping it only in the header.
