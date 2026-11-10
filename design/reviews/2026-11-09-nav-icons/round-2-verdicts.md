# Icon lane, review round 2 (2026-11-09)

Captures: the after-* files and contact-sheet.png as committed with this file.

## ux-designer: VERDICT: HOLD

Musts and the round 3 fixes:
1. The keyboard focus ring touched the icon and the last letter (header-2000-focus-planner).
   Fixed: each door carries 4 px of inline padding taken back by a matching negative margin, so the
   ring clears the icon and label without moving the bar. Round 3 captures focus on Planner,
   Guides (a door with a chevron) and Best in slot (the current door).
2. account-menu-1440 and search-1440 were cropped. Fixed: the account menu is the full 1440
   viewport; the lookup is the full bar width at 1440 and 1024.

Shoulds:
1. No 1920 capture. Added header-1920 (on the 1640 px bar line: the 1920 wide-column bar is 1632 px,
   so the doors stay labels there; the icons need the room).
2. The sheet's settings rows mixed the open chain-link glyph with filled ones. Fixed: setup and guild.
3. "Claim your guild" / "All guides" foot read as dim text. Fixed: 13 px gold link, gold-hover.
4. Phone Menu group headings were 13 px and dim beside 16 px doors. Fixed: a heading is a door too,
   with the plain doors' type, ink and 44 px row. Loot council's line trimmed to "Who gains most
   from each drop, by the sim" (no orphan).

## wow-player: VERDICT: SHIP

Shoulds: focus ring (fixed, see ux must 1); Account and Roster both showed two heads. Fixed by
moving Roster and Find a guild back to the owner's suggested `character` glyph; the account menu's
Account tile keeps `manage` as it shipped. Get set up and Plan both carry a crown motif: the
owner's own glyphs, kept. "Claim your guild" now reads as a link.

## ai-slop-reviewer: VERDICT: PASS

The sheet's two chain-link rows broke the written family rule: fixed (see ux should 2).

## info-density-reviewer: VERDICT: FAIL

Must: the clipped focus ring (fixed, see ux must 1). Shoulds: the phone Menu's Logs and Guild
headings repeat their first row's glyph. Kept: the heading is the door and carries the door's icon
(owner: every door gets its icon, phone rows too); Your reports and Your guild carry the owner's
mapping. Account and Roster sharing two heads: fixed (wow-player should). The doors-only-on-the-wide-
bar rule is now written in inventory.md.
