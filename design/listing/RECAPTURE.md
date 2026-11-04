# Gallery recapture: the twelve in-game shots the owner takes

The store plates (`gen_gallery.mjs`, 3200x1800 PNG) draw each capture at most one capture pixel
per plate pixel and never enlarge it. As of 2026-11-03 every plate's capture is one of the
site's small JPEG crops in `web/public/addon/` (504 to 1166 px wide, saved at web quality,
with block artifacts on text edges). The settings capture is also two shots stitched together,
with a visible seam. No lossless originals of them exist anywhere on disk. A small crop drawn
1:1 looks small on the plate, and the old 1.5x enlargement made it blurry. The only real fix is
new lossless captures from the game.

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

The client saves screenshots to `/Applications/World of Warcraft/_classic_beta_/Screenshots/`.
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
| 1 | `chip-tooltip-legend` | The chip beside the minimap, hovered | The chip and the whole minimap, with the chip's tooltip open: the build line, the next talent, the green-arrow and gold-corner legend, and the click hints. |
| 2 | `talent-marks-trees` | Talent window (N) | All three trees with the addon's numbers on the build's talents, the next one ringed, and the progress strip at the bottom. Use a build with some points off plan, so every kind of mark appears. |
| 3 | `item-tooltip-ranged` | Character panel (C), hovering an equipped item | The paper doll, the item tooltip with the Forever Sixty block (estimate, next upgrade, source, "Also good" with two items), and the comparison tooltip beside it. |
| 4 | `bag-arrow` | Character panel (C) and bags (B) open together | The green arrow on a paper-doll slot and on the bag item that beats it, plus a gold corner on a slot worth chasing. The current capture shows only the bag; the caption needs both. |
| 5 | `character-panel-tab` | The addon's tab on the character panel | Next targets with their sources, the rotation at your level, and the build's progress, all fully visible with nothing scrolled off. |
| 6 | `player-tooltip-score` | Hovering another player | The unit tooltip with the Forever Score line. If you can, take a second shot with Shift held so the five parts show (`player-tooltip-score-shift.png`). Claude will use whichever reads better. |
| 7 | `chat-badge-guild` | Chat frame | Make the chat window wider and taller first. Show several guild, trade or whisper lines with score badges in front of the senders, not one line on its own. |
| 8 | `guild-tab` | The addon's guild window | Online only unticked, sorted by Score, with several rows that have scores and specs (not mostly "no logs"). |
| 9 | `group-panel-solo` | Social (O) > Raid, the addon's group panel | Ideally in a party with addon users, so member rows (spec, gear gap, score) show and a ready-check line names what someone lacks. Solo only shows "Not in a group", so shoot solo only as a last resort. |
| 10 | `build-dialog-paste` | `/fs`, Paste a build tab | The "Following" line, the next talent and the paste box. If you can, paste a code so the build preview shows. |
| 11 | `build-dialog-copy` | `/fs`, Copy my character tab | The export selected, the "Press Ctrl+C" line and the companion status lines. |
| 12 | `settings-switches` | Settings > AddOns > Forever Sixty | All nine switches with their one-line explanations, in one shot with no stitching. If the page scrolls, take the top view and a second shot scrolled to the end (`settings-switches-2.png`); Claude will pick the one shot that holds the switches, not join the two. |

## After the shots

Claude crops them and reruns `node design/listing/gen_gallery.mjs`. Each line of its output
must end in `ok`. `RECAPTURE` flags a JPEG source or a crop under 60% of the stage. Then
follow "Uploading the gallery" in `curseforge-description.md`.

The `/setup` page still serves the old JPEGs in `web/public/addon/`. Once the PNGs exist, they
can replace them as a separate change.
