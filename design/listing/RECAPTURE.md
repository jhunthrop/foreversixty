# Gallery captures: the addon's mocks, and the in-game shots that could replace them

The store plates (`gen_gallery.mjs`, 3200x1800 PNG) draw each capture at most one capture pixel
per plate pixel and never enlarge it. As of 2026-11-03 every plate's capture is one of the
site's small JPEG crops in `web/public/addon/` (504 to 1166 px wide, saved at web quality,
with block artifacts on text edges). The settings capture is also two shots stitched together,
with a visible seam. No lossless originals of them exist anywhere on disk. A small crop drawn
1:1 looks small on the plate, and the old 1.5x enlargement made it blurry. The only real fix is
new lossless captures from the game.

## As of 2026-11-13: the gallery comes from the addon's mocks

The owner asked for the addon's mocks as the store screenshots. Every plate now comes from a mock, so
no plate waits on the game. `gen_captures.py` renders the 2.0.48 review boards again, the ones each
lane of the one-window release shipped against (`design/mocks/addon_hub_*.py` and
`addon_builds_window.py`, every word and figure from the addon's own Locale.lua, Data.lua and
models), hides the boards' review notes, cuts out the frame and writes `captures/<name>.png`.
`gen_gallery.mjs` then puts each on its plate. Every character and guild in them is made up
(Velisandre in Thornwake Company and her raid team). The script's store fixture makes the plates
one evening of one character: one raid day, one sign-up, one last raid, one set of gains (its
docstring lists what it aligns across the lanes).

| # | Plate | Mock board |
|---|---|---|
| 1 | `chip` | `addon_hub_l6.py` `l6-hover`, the first scene: the line and its hover |
| 2 | `window-me` | `addon_hub_me.py` `me-raid-day` |
| 3 | `window-builds` | `addon_builds_window.py`'s window at 60, following the Beast Mastery raid build |
| 4 | `window-gear` | `addon_hub_l3.py` `gear-hunter-60-raid`, the window only |
| 5 | `window-rotation` | `addon_hub_rotation.py` `rotation-60` |
| 6 | `window-guild` | `addon_hub_l5.py`'s raid day, signed up In |
| 7 | `window-officer` | `addon_hub_l7.py` `l7-raid-day` |
| 8 | `bag-arrow` | `addon_hub_l10.py` `l10-gear-marks`, without the bag item's tooltip |
| 9 | `quest-reward-marks` | `addon_hub_l10.py` `l10-rewards`, without the "nothing helps" frame and the repeated roll |
| 10 | `player-tooltip` | `addon_hub_l10.py` `l10-player`, the stale record and the Shift view |

When a tab changes, fix its lane generator and rerun `gen_captures.py` and `gen_gallery.mjs`;
never edit a capture by hand. Not in the gallery, because no mock draws them whole: the chat
badges, the full talent window (the L10 board draws only the strip and two talents) and the
Settings page (taller than a plate can show readably). The in-game shots below would add them, and
a lossless shot of any plate's surface may replace its mock capture under the same name (take
that name out of `gen_captures.py`'s `CAPTURES` too, or the next run draws over it).

## In-game shots, optional

Claude never drives the game client; the owner takes every screenshot.

## Client setup, once

1. **Resolution.** Settings > Graphics: Display Mode *Windowed (Fullscreen)* at the screen's
   native resolution, which is 3456x2234 on the MacBook's built-in Retina display or
   3840x2160 on a 4K monitor. Make sure the list shows the full Retina size, not 1728x1117.
   Set Render Scale to **100%**.
2. **UI scale 1.0.** Settings > Graphics (or System) > *Use UI Scale* on, slider at **1.0**.
   The same in chat: `/console useUiScale 1` then `/console uiScale 1`. At 2160 to 2234 px
   tall, one UI unit is about 2.8 to 2.9 screen pixels, so text is drawn about three times
   larger than the current crops.
3. **Lossless screenshots.** In chat: `/console screenshotFormat png`. If the next screenshot
   still comes out as `.jpg`, use `/console screenshotFormat tga` (also lossless). Never use JPEG.
4. **A key for the shot.** Settings > Keybindings > Miscellaneous > *Screen Shot*. Bind it to
   a key you can press while the mouse is hovering, such as F13 or Ctrl-P, because Mac
   keyboards have no Print Screen key. Do not use macOS Cmd-Shift-4: it captures the scaled
   desktop, not the game's frame buffer.
5. Keep the UI on (no Alt-Z) and the addon on its current release. Turn off other addons that
   restyle tooltips or frames, so every frame looks like the default UI.

The client saves screenshots to `/Applications/World of Warcraft/_forever_/Screenshots/`.
That folder appears after the first shot, and the path follows whichever client folder you
play from.

## What to shoot

Take each shot full screen; do not crop it yourself. Copy each file into
`design/listing/captures/raw/<name>.png` (or `.tga`), using the names below, then ask Claude
to crop them. Claude crops each one losslessly into `design/listing/captures/<name>.png`
and reruns `gen_gallery.mjs`, which uses a PNG there in place of the site's JPEG.

**Crop size.** The plate's capture area is 2880x1200 px. A crop fills it well at about
1730 px wide or 720 px tall. A single tooltip at UI scale 1.0 is only about 900 px wide, so
frame each shot with some of the scene around the frame (the character, the minimap, the
world). That gives Claude room to crop wide.

**Character.** Use a character with a pasted build that is partly followed, an upgrade in
the bags, a Forever Score and a guild. Other players' names in a shot must not be readable on a
store plate, so shoot alone or crop them out. Stand somewhere with a
calm, dark backdrop (an inn or a city at night) so the frames read clearly.

| # | `<name>` | Frame | What must show |
|---|---|---|---|
| 1 | `forever-sixty-line` | The Forever Sixty line beside the minimap, hovered | The line and the whole minimap, with its two-line hover open. Shoot on raid day if you can, so it leads with tonight's raid; otherwise a next talent ("Next talent at 25: ..."). |
| 2 | `window-me` | The Forever Sixty window on Me (`/fs me`) | The header (crest, name, level, guild, build), the tab strip, the Forever Score with its five parts, the last raid boss by boss, Right now and Your sim, nothing cut off. |
| 3 | `item-tooltip-verdict` | Bags (B), hovering an upgrade | The item tooltip with its one verdict line (the green arrow and "+2.2 DPS") and its one support line. If you can, a second shot with Shift held (`item-tooltip-verdict-shift.png`). |
| 4 | `talent-window-marks` | Talent window (N) | All three trees with numbers on the points still to take, the next one ringed, and the strip ("Next at 25: ... · 12 of 15 · 3 off plan"). Use a build with some points off plan. |
| 5 | `bag-arrow` | Character pane (C) and bags (B) open together | The green arrow on a paperdoll slot and on the bag item that beats it, plus a gold corner on a slot with an upgrade to chase. |
| 6 | `quest-reward-marks` | A quest's reward choices | The green arrow on the reward that is an upgrade and "Best in slot" under the pick; or, on a quest where nothing helps, "Sells for most". A group loot roll with its line is a good second shot (`roll-line.png`). |
| 7 | `player-tooltip` | Hovering another player | The unit tooltip with "Forever Score 1240" in the tier's colour. A second shot with Shift held shows the five parts (`player-tooltip-shift.png`); Claude uses whichever reads better. |
| 8 | `chat-badge-guild` | Chat frame | Make the chat window wider and taller first. Several guild, Trade or whisper lines with "FS 1240 ·" before the senders, not one line on its own. |
| 9 | `window-guild` | The window's Guild tab (`/fs guild`) on raid day | Tonight's raid with your sign-up and job, your prep, the sign-ups and the next raids. Needs the companion running and a raid on the site. |
| 10 | `window-officer` | The window's Officer tab (`/fs officer`), as an officer | The roster with scores and sign-ups and tonight's readiness, ideally in a raid with addon users. A master loot session is a good second shot (`window-officer-loot.png`). |
| 11 | `window-builds` | The window's Builds tab (`/fs builds`) | Your class's builds with the followed one gold, the next talent points, and a build from the site or a pasted code with its preview. |
| 12 | `settings-page` | Settings > AddOns > Forever Sixty | The commands at the top and as many of the twelve switches as fit, in one shot with no stitching. If the page scrolls, take the top view and a second shot scrolled to the end (`settings-page-2.png`); Claude picks the one shot that reads best, never joins the two. |

## After the shots

Claude crops them and reruns `node design/listing/gen_gallery.mjs`. Each line of its output
must end in `ok`. `RECAPTURE` flags a JPEG source or a crop under 60% of the stage. Then
follow "Uploading the gallery" in `curseforge-description.md`.

The `/setup` page still serves the old JPEGs in `web/public/addon/`. Once the PNGs exist, they
can replace them as a separate change.
