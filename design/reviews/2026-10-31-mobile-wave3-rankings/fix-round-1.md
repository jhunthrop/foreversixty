# F11 build review, fix round 1

Both reviewers said HOLD (`ux-verdict-build.md`, `player-verdict-build.md`). This is the work list.

## Done by the lead (data and copy)
- Fixtures regenerated (`f11_fixtures.py`): Voltaris wears his own, nearly complete list (one gold
  dot, Trinket 2); Lightningbolltz is a week behind in the Elements set (six dots, gap 36.0). Each
  card's Onyxia kill is the same fight as its board row (time, duration, value, report).
- Copy keys `items.unknownItem` ("{slot} · Item {id}", the owner's ruling) and `paperdoll.twoHand`
  ("Two-hand in main hand").

## Both apps
1. Home's Rankings row takes its spec and class words from the current character, never from the
   rows; its capture uses Lightningbolltz (Elemental Shaman) as the current character.
2. The answer card names the character in both states: `rank.metaNamed`, the name in gold, on the
   page and off it.
3. The big ordinal is in the rank's parse colour on both platforms.
4. Paperdoll: an off hand under a two-hander is an empty Off hand socket with `paperdoll.twoHand`;
   the gold dot only where the slot has a pick that gains (never on an empty or unknown slot
   without one); an unknown item reads `items.unknownItem`.
5. A row's age never wraps (`2 d ago` stays on one line).
6. Re-record goldens for the new fixtures; recapture.

## iOS
7. Watch star and share in the navigation bar on the character page (on, off, failed).
8. Filters sheet: chips without the stray tick marks; `Today only` a real switch inside the margin.
9. Encounter list sheet anchored to the top.
10. Weapons row as on Android (socket above the name), no four-line wraps.
11. The pinned row's gold edge stays inside the card's rounded corner.

## Android
12. The filters sheet sits above the tab bar and scrolls to `Reset filters`.
13. The unranked answer card keeps `Your rank`.
14. Score and build totals in mono.

## New captures
iOS rankings error; character error, loading, watch failed; the long-press menu (both); Android
largest-text rankings; filters with `Today only` on and off (both).
