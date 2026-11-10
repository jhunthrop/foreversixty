# Icon lane, review round 1 (2026-11-09)

Captures: the before-*/after-* pairs and contact-sheet.png in this folder (round 1 state).

## ux-designer: VERDICT: SHIP

Shoulds and what changed for round 2:
1. Design sheet buttons used the download tray on Subscribe, Manage billing, See plans and
   Unlink Discord, teaching a wrong meaning. Fixed: the samples carry no glyph.
2. The Discord row's text started 2 px left of the other Guild rows. Fixed: the brand tile and the
   row icons share one 24 px column (the tile's -2 px margin is gone, row icons are 24 px).
3. Door icons show only from about a 1910 px viewport. Noted as the measured limit, not a bug:
   with an 18-letter name signed in, the 1440 bar has 24 px free and the icons need about 190 px;
   the owner's rule keeps the labels whole, so the bar shows icons only where it is at least
   1620 px wide (the wide column from about 1910 px).
4. sim and planner fragment at 16 px. Fixed: door icons are 20 px; no custom icon is drawn below
   20 px anywhere now (SettingsRow glyphs and the signed-out ghost moved from 16 to 20).
5. Phone Menu heading icons were smaller and dimmer than the plain doors' icons. Fixed: same size
   and the strong ink.
6. Door icon hover, focus-visible and open states were not captured. Added: after-header-2000-hover-sim,
   after-header-2000-focus-planner, after-panel-guild-2000.

Coulds: the filled `upgrade` triangle stays (the paperdoll spec's gold triangle badge, 11 px in an
18 px disc); "Find a guild" now uses the two-heads glyph with Roster; the crown repeats on the
sheet's state rows are varied.

## wow-player: VERDICT: SHIP

Shoulds and what changed:
1. Loot council reuses the BiS breastplate. Kept: the owner's mapping ("loot council -> bis or the
   closest"); the set has no loot or dice glyph. Reported as a missing glyph.
2. Roster and Find a guild used the character shield. Fixed: the two-heads `manage` glyph (people).
   Guild logs and Your reports share `logs`: the owner's mapping; no separate guild-logs glyph exists.
3. Rankings podium: kept, it is the owner's own glyph.
4. Icons only at wide bars: the measured limit (see ux 3).

## ai-slop-reviewer: VERDICT: FAIL

Findings and the response:
1. Icon before every label in the phone Menu and the panels (tell 9). The owner asked for exactly
   this: every door and every panel row carries its custom icon, and the phone Menu rows too. The
   icons are kept; what made them read as filler is fixed: the duplicated character shield is gone
   (Roster and Find a guild now people), and the phone Menu lists the Logs and Guild rows one to a
   line instead of two cramped columns.
2. Filled custom icons as button glyphs on the design sheet. Fixed: no glyph on the sample buttons.
   The rule is written into the generator and the inventory: the custom (filled) set marks a place
   or an object (doors, panel rows, account tiles, settings rows), the line family a control or a
   state (chevrons, close, check, menu, spinner, refresh, chart markers).
3. The crown repeated on the sheet's state rows. Fixed: varied.
4. Door icons small at 2000 (about 14 px drawn). Fixed: 20 px.
5. paste-export, home, setup and upload at 16 px: the owner's traced set, used at 20 px and up only.

## info-density-reviewer: VERDICT: PASS

Shoulds: Discord row misaligned (fixed, see ux 2); repeated glyphs (Roster fixed; Loot council and
the two logs rows are the owner's mapping, reported); phone Menu cramped (fixed: one row a line
below 560 px). Blank spec icons on the BiS page body belong to the BiS lane.
