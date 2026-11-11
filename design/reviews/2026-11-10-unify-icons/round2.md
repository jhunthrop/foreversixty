# Unify icons, review round 2 (2026-11-10)

The captures this round read were replaced by round 3's (round3.md); the folder holds those.

Captures: `after/` at `22afb3237`.

| Reviewer | Verdict |
|---|---|
| ux-designer | HOLD |
| wow-player | HOLD ("once the two empty boxes are gone, I would call this SHIP") |
| ai-slop-reviewer | FAIL |

## Verified fixed

- Coach: the drill chips (Timelines, Resources) and the "Rage full" heading draw no box. Inline
  16 px icons are unframed glyphs and the line height is unchanged. Page height at 390 is back to
  4736 px.
- Guild report: the event lines (Tail Sweep, Deep Breath, Flame Breath) draw no box.
- Report: the role marks are unframed glyphs, and the socket stays in the table's icon column.
- Zone plaque at 390: the warm frame, with an identical layout.

## Must-fix that remained

These icons are missing from the fixture build (Beast Mastery, Marksmanship, Protection Warrior and
others, and the race portraits), so their files 404:
- the guide and BiS spec tabs;
- the bis-landing and guide class spec rows;
- the guide race rows;
- the home feed's Protection Warrior row.

On a server-rendered page, the component's error handler never runs before hydration, and
socket.css drew the dark socket over the broken image. bis-landing and the race rows looked the same
before this lane.

Fix (`971d1c589`):
- Base.astro inlines a capture-phase error listener at the top of `<head>` (`lib/ui/icon-fail.ts`).
  On every page, hydrated or not, it hides a SpellIcon frame that does not hold a column.
- `ui/RaceIcon`, a fifth copy of the frame, is deleted. The race rows draw SpellIcon.astro, and the
  icon guard fails if RaceIcon comes back.

## Should-fix noted

- ux-designer: home's "First raids open" row looked dimmed. The newest feed row, when it is fresh
  against the clock, slides in from opacity 0 (ui/FeedRow `fresh`), so the capture caught it
  mid-animation. This lane does not touch it.
