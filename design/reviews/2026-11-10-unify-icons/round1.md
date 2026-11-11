# Unify icons, review round 1 (2026-11-10)

The captures this round read were replaced by round 3's (round3.md); the folder holds those.

Captures: `before/` at `e4de576d2` (the lane's base, after merging lane/e2e-parity-fix and
lane/header-width) and `after/` at `90c07220e`, 16 pages at 1440 and 390, first screen and full page,
written by `tests/e2e/frame-capture.spec.ts` (FRAME_CAPTURE_DIR=design/reviews/2026-11-10-unify-icons).

| Reviewer | Verdict |
|---|---|
| ux-designer | HOLD |
| wow-player | HOLD |
| ai-slop-reviewer | FAIL |

## The one must-fix all three raised

Empty sockets where the old code drew nothing: they read as unticked checkboxes.
- Coach: the drill chips (Timelines, Resources) and the "Rage full for 14 seconds in total" heading.
- Guild report: the event lines (Tail Sweep, Deep Breath, Flame Breath).
- Guides and BiS: the Beast Mastery and Marksmanship spec tabs, the spec rows and race rows on
  `/bis` and the class guide. The fixture build has no file for these icons.
- Home feed: the Protection Warrior BiS row drew a white-framed box (ItemIcon), unchanged from before.

Fix (`22afb3237`): SpellIcon draws nothing for no icon or a failed image. A caller whose icons line up
in a column (the report tables and rails, the sim enchant list) sets `socket` and keeps the dark
socket. The static twin hides a failed image's frame. The home feed draws a spec or guide row's icon
as a spell icon.

## Should-fix

- ux-designer: role and spec glyphs at 14 px read as a checkbox with an X when framed, and the inline
  16 px coach icons look busy. Fixed in `22afb3237`: an icon of 16 px or less is an inline glyph with
  no frame.
- ux-designer: guild-loot looks unchanged. That is expected. RaidLootItem is drawn in the guild report
  and attendance views, not on the Loot tab, which already used ItemIcon.
- wow-player: the sim result header says item level 62.9 and Top Gear says 62.0. They are different
  characters: Thandrel's saved run, and the tool fixture's Orc wearing two items. Both now read
  `lib/gear/paperdoll` averageItemLevel.
- wow-player: the small boxed glyphs under report player names are the role marks (tank, healer,
  dps). They are unframed now (the 16 px rule).
- ux-designer: the chart lane markers were not in a capture.

## Kept as right by all three

- Raid zone icons now sit in the warm frame, not a white or legendary item border.
- Raid loot items keep their quality border and colour.
- Consumables and enchants use the spell frame, as on the action and buff bars.
- Report ability and buff icons have one frame and do not shift alignment.
- Item names are quality coloured.
