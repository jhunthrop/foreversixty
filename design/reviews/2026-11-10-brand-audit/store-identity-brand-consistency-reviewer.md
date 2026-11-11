# brand-consistency-reviewer: store, listing and identity (2026-11-10)

Sources: `captures/identity-sheet.png`, `design/reviews/2026-11-10-app-icons`, `2026-11-10-seal-flat`, `design/listing/*`, `web/src/lib/og.ts` and `design/specs/2026-11-09-directions.md`.

## DRIFT

**BLOCKER**
1. **On main, the mobile app icon is a third seal.**
   - The black disc with a gold ring, a dotted ring and "LX" appears in `apps/shared/brand/ios/Brand.xcassets/AppIcon.appiconset/AppIcon-1024.png` and in Android's `ic_launcher_foreground.png` (identity sheet).
   - It matches neither the painted seal nor the flat seal.
   - `design/specs/2026-11-09-directions.md:69` says the app icon is the flat seal at 62% inside the hammered-gold frame.
   - On main the icon dates from `44349ebfb` (2026-10-30).
   - The unmerged `origin/lane/app-brand-polish` carries `c669fb8a4`, "the painted app icon, seal v4 on the painted blue ground". The fix exists and has not landed.
   - The three variants (light, dark and tinted) are also not one family: black disc, flat on gold, and grey.
2. **The apps have no store listing at all.**
   - `apps/` has no fastlane, no metadata, no subtitle or keywords, and no store screenshots.
   - The only screenshots are test goldens in `apps/android/app/src/test/screenshots`.

**SHOULD**
3. **The Discord server icon is the same black disc** (`design/community/discord/assets/icon.png`). The bot avatar is painted and the favicon is flat. Use the painted seal.
4. **Two Discord banners.**
   - `design/community/discord/assets/banner.png` has tracked forced caps ("FOREVER SIXTY", "PLAY YOUR CLASS BETTER") in a Cinzel-like face, a tiny black-disc seal and a row of class icons.
   - `design/assets/art/discord-banner-2026-11-09.jpg` matches the web: the painted seal over "Forever Sixty" in Marcellus.
   - Keep the second and delete the first.
5. **The listing gallery is stale.**
   - `design/listing/gallery/*` was last committed `6b4e85898` on 2026-11-03.
   - It shows the navy addon panel and "(needs bullets)", and predates the umber panel, the v4 minimap and the tour.
   - `RECAPTURE.md` itself says the plates come from lossy JPEG crops and need lossless in-game captures.
6. **`companion/ui/seal.svg` is the flat seal.** Confirm it is never drawn above 44 px. Large placements should use `web/public/seal-400.png`.

**NICE**
7. The CurseForge and Wago summaries do not share the site tagline "Play your class better." Their dashboards were filled on 2026-10-15 and may lag `design/listing/*.md`.

## OG images

These pass. `web/src/lib/og.ts` renders every per-page card with the road art under a veil, the painted seal, "Forever Sixty" in Marcellus and the fan-made line. Its colours are named token values, hardcoded because satori cannot read CSS variables. `og-pages.ts` tests that every page has a card. The bare `og-art.jpg` on the sheet is only the background.

## Seal variants

| Asset | Variant | Size | Rule |
|---|---|---|---|
| `web/public/favicon.svg`, `favicon.ico` | flat | 16-48 | ok |
| `web/public/apple-touch-icon.png` | painted | 180 | ok |
| `web/public/seal-400.png` | painted | 400 | ok |
| `addon/ForeverSixty/icon.png` | painted | 64+ | ok |
| `addon` `media/minimap.tga` | flat | 20-32 | ok |
| `companion/internal/icon/assets/seal-tray-*` | flat + dot | 16-32 | ok |
| `companion/ui/seal.svg` | flat | 32 in `install.html` | ok if never above 44 |
| iOS AppIcon light / dark / tinted (main) | black disc / flat on gold / grey | 1024 | **no** |
| Android `ic_launcher_foreground.png` (main) | black disc | adaptive | **no** |
| Android `ic_notification_seal.xml` | mono disc | 24 dp | system mono (allowed) |
| `design/community/discord/assets/icon.png` | black disc | server icon | **no** |
| `design/community/discord/assets/banner.png` | black disc, tiny | ~28 | **no** |
| `design/assets/art/discord-banner-2026-11-09.jpg` | painted | ~64 | ok |

## Generated vs hand-made

- **Generated:**
  - `design/logo/build.py` (marks, favicon, exports)
  - `web/src/lib/og.ts` (OG cards)
  - `design/listing/gen_gallery.mjs` and `md2html.mjs`
  - `web/scripts/build-app-icon.mjs` (on the app-brand-polish lane)
- **Hand-made one-offs:**
  - the iOS and Android icons on main
  - the Discord `icon.png`, `banner.png` and `crest-*.png`
  - the tray PNGs, built by `companion/tools/icons/render.py` from no shared source

VERDICT: DRIFT
