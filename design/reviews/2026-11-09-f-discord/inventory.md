# Discord-visible surfaces (lane f-discord, 2026-11-09)

Every place a player or officer sees Forever Sixty inside Discord, what draws it and the data it shows. Captures: `discord/<id>-desktop.png` (640) and `discord/<id>-phone.png` (390), drawn from the real handler payloads in `payloads/` by `design/mocks/2026-11-09-f-discord/render.mjs`; rendered images at their real size beside this file.

Look before this lane: gold `#e5b955` stripes and the flat gold seal, Barlow and Cinzel in every image. After: the teal accent (`color.forever.accent`) on normal posts, the seal v4 as avatar and before every footer, Open Sans and Marcellus, the brown ground, blue band and framed card in the images.

## 1. Rendered images (PNG, drawn by the API)

| Surface | Where it shows | Drawn by | Data |
|---|---|---|---|
| Character sheet (`sheet-char.png`) | /char, Share in channel (Components V2 media gallery, `attachment://sheet.png`) | `api/internal/sheet` | name, class, spec, level, guild, faction; 17 slots: item name, quality, icon, item level; per-slot DPS or points gain over what is worn at the band (`upgrades`), "Upgrade available" / "Gain not modelled yet"; average item level |
| BiS sheet (`sheet-bis.png`) | /bis, Share in channel | `api/internal/sheet` | class, spec, faction, band level and range (phase), 17 picks: name, quality, icon, source line (shortened to fit: quest, vendor, instance, crafted, PvP rank), band file's date |
| Item tiles (`payloads/tiles/*.png`) | thumbnails of every /bis and /char item row (`/v1/discord/tiles/{build}/{item}.png`) | `sheet.Tile` | item icon, quality ring |
| Quality dots (`payloads/emoji/*.png`) | application emoji `fs_q_*` before item names | `sheet.QualityDots` | quality colour |
| Spec marks | application emoji `fs_{class}_{spec}` in raid rosters | `discordbot.EmojiImages` | the spec's game icon |
| Build link preview (`card-build.png`) | unfurl of `/b/{id}` links | `api/internal/card.Render` | build title, race, class (class colour), talent split, level |
| Report link preview (`card-report*.png`) | unfurl of report links | `card.RenderReport` | title, zone, bosses killed and wiped on, kill count, night length |
| Fallback preview (`card-fallback.png`) | a build or report whose card cannot be drawn | `card.Fallback` | brand only |
| Site page previews | unfurl of every site page | `web/src/lib/og.ts` (satori) | page title and kicker; already round 5 before this lane, unchanged |

## 2. Slash commands (bot replies, ephemeral unless shared)

| Command | Frames | Format | Data |
|---|---|---|---|
| `/bis` | `bis-card-*` (guided card: class, spec, range, ready, nudge, stale), `bis-result`, `bis-panel-p1/p2/slot/sharefail`, `share-bis`, `share-bis-nosheet` | Components V2, class-colour accent | class and spec lists, the linked main, bands; per slot pick name, quality dot, slot, item level, source, DPS over the next best; the band date |
| `/char` | `char-card-*` (unlinked, own, modal, unknown), `char-panel`, `char-panel-nogear`, `share-char` | Components V2, class-colour accent | name link, level, spec, class, guild, item level, readiness, last sync; top 3 upgrades with tile, slot, DPS gain, quality, item level, source, what it replaces; count of the rest |
| `/raid` | `raid-answer` | embed, accent stripe, seal footer, buttons In/Late/Bench/Out/Raid page | next raid title and guild date, Discord timestamp, signed count, the caller's own sign-up |
| `/signup` | `signup-card`, `signup-card-in` | embed + raid and character selects + status buttons | the guild's coming raids, the caller's characters, their status |
| `/report` | `report-select`, `report-pick` | embed, kill-green stripe, then one embed per award in its quality colour, seal footer | night: came of signed, kills, wipes, deaths, length; kill order with times and wipes before; deaths and consumables (officers); loot awards with item, quality, item level, winner, boss |
| `/attendance` | `attendance-select`, `attendance-modal`, `attendance-pick` | embeds in class colour, seal footer | a raider's attendance rate, raids came of signed, last nights; officer select of lowest attendance |
| `/readiness` | `readiness-select`, `readiness-pick` | embed, accent stripe, seal footer | the raid, signed In or Late, who has something to fix and what |
| `/link` | `link-ask`, `link-done` | text + link button | link state |
| `/help` | `help` | embed, accent stripe, seal footer | the command list with descriptions (registration text, unchanged) |
| Spec card | `spec-ask-card`, `spec-ask-done` | embed in class colour + spec select, seal footer | the raid, the character, its specs and roles |

## 3. Channel posts

| Post | Frames | Sender | Data |
|---|---|---|---|
| Raid post (open, partial, full, tomorrow, locked, moved, cancelled, ended, done) | `post-scheduled`, `post-partial`, `post-full`, `post-tomorrow`, `post-reminder` | bot (drainer), edited as it changes | title, guild date, Discord timestamp, guild time, signed of size; roster by role with spec marks (tanks, healers, melee, ranged, late, bench, out); buttons |
| Raid post, webhook guild | `webhook-raid-posted`, `webhook-raid-moved`, `webhook-raid-cancelled`, `webhook-reminder`, `post-reminder-webhook` | guild webhook | the same, roster as text lines; cancelled in the wipe red |
| Raid report post | `webhook-report` | officer's Post to Discord, webhook | kills with times, wipes, time, came of signed, deaths, top parses, loot to winners, coach sentence |
| Guild test message | `webhook-guild-test` | webhook | guild name, Raids link |
| Personal webhook: test, sim finished, log parsed | `webhook-test`, `webhook-sim`, `webhook-log` | the player's own webhook | sim result and run time with the class crest; log zone, fights and kills, visibility |

## 4. Not drawn by this lane

- The bot's own avatar, banner and the server's icon and banner are uploaded by the owner.
- Command names and descriptions are the registered command set, unchanged.
- The site's page previews (`/og/*.png`) were already round 5 (`web/src/lib/og.ts`).
- The mobile apps bundle `web/public/seal-400.png` directly (`apps/ios/project.yml` lines 71 and 178, `apps/ios/App/Common/BrandImages.swift`, `apps/android/build-logic/convention/src/main/kotlin/BundledDataConventionPlugin.kt`), so their next build picks up the seal v4 with no copy to update.
