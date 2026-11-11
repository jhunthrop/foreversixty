# Gallery recapture: the twelve in-game shots the owner takes

The store plates (`gen_gallery.mjs`, 3200x1800 PNG) draw each capture at most one capture pixel
per plate pixel and never enlarge it. As of 2026-11-03 every plate's capture is one of the
site's small JPEG crops in `web/public/addon/` (504 to 1166 px wide, saved at web quality,
with block artifacts on text edges). The settings capture is also two shots stitched together,
with a visible seam. No lossless originals of them exist anywhere on disk. A small crop drawn
1:1 looks small on the plate, and the old 1.5x enlargement made it blurry. The only real fix is
new lossless captures from the game.

As of 2026-11-10 the plate list follows addon 2.0.38. Nine of the twelve surfaces changed since
those crops (the panel replaced the character pane's flyout, the tooltips, talent strip, player
tooltip, build dialog tabs and Settings page all read differently), so those plates have no
capture at all and `gen_gallery.mjs` stops at the first one. Only `bag-arrow`,
`chat-badge-guild` and `guild-tab` still have a JPEG stand-in.

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
the bags, a Forever Score and a guild: Bow Jackzon or Obnoxious Yell. Stand somewhere with a
calm, dark backdrop (an inn or a city at night) so the frames read clearly.

| # | `<name>` | Frame | What must show |
|---|---|---|---|
| 1 | `forever-sixty-line` | The Forever Sixty line beside the minimap, hovered | The line and the whole minimap, with its two-line hover open. Shoot on raid day if you can, so it leads with tonight's raid; otherwise a next talent ("Next talent at 25: ..."). |
| 2 | `addon-panel` | The panel (click the line) | The title band (crest, name, level, score) and every region with something in it: Tonight or Last raid, Upgrades with an Equip button, Talents with Change and Talents, and the rotation, nothing cut off. |
| 3 | `item-tooltip-verdict` | Bags (B), hovering an upgrade | The item tooltip with its one verdict line (the green arrow and "+2.2 DPS") and its one support line. If you can, a second shot with Shift held (`item-tooltip-verdict-shift.png`). |
| 4 | `talent-window-marks` | Talent window (N) | All three trees with numbers on the points still to take, the next one ringed, and the strip ("Next at 25: ... · 12 of 15 · 3 off plan"). Use a build with some points off plan. |
| 5 | `bag-arrow` | Character pane (C) and bags (B) open together | The green arrow on a paperdoll slot and on the bag item that beats it, plus a gold corner on a slot with an upgrade to chase. |
| 6 | `quest-reward-marks` | A quest's reward choices | The green arrow on the reward that is an upgrade and "Best in slot" under the pick; or, on a quest where nothing helps, "Sells for most". A group loot roll with its line is a good second shot (`roll-line.png`). |
| 7 | `player-tooltip` | Hovering another player | The unit tooltip with "Forever Score 1240" in the tier's colour. A second shot with Shift held shows the five parts (`player-tooltip-shift.png`); Claude uses whichever reads better. |
| 8 | `chat-badge-guild` | Chat frame | Make the chat window wider and taller first. Several guild, Trade or whisper lines with "FS 1240 ·" before the senders, not one line on its own. |
| 9 | `guild-tab` | The guild window's Forever Sixty button | Online only unticked, sorted by Score, the count line ("97 members · 3 online"), and several rows with scores and specs. |
| 10 | `group-panel` | Social (O) > Raid, the Forever Sixty button | In a party or raid with addon users, so member rows (spec, item level, gap, score, enchants) show, ideally after a ready check that names what someone lacks. Solo shows only "Join a group to see its specs, gear and buffs.", so never shoot it solo. |
| 11 | `build-dialog-builds` | `/fs`, the Builds tab | The "Following" line, the next talent, a build from the site in the list or a pasted code with its preview, and the Copy my character tab label. |
| 12 | `settings-page` | Settings > AddOns > Forever Sixty | The commands at the top and as many of the twelve switches as fit, in one shot with no stitching. If the page scrolls, take the top view and a second shot scrolled to the end (`settings-page-2.png`); Claude picks the one shot that reads best, never joins the two. |

## After the shots

Claude crops them and reruns `node design/listing/gen_gallery.mjs`. Each line of its output
must end in `ok`. `RECAPTURE` flags a JPEG source or a crop under 60% of the stage. Then
follow "Uploading the gallery" in `curseforge-description.md`.

The `/setup` page still serves the old JPEGs in `web/public/addon/`. Once the PNGs exist, they
can replace them as a separate change.
