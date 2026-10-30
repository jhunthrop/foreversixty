# apps/shared

Assets and tokens both native apps (iOS in `apps/ios`, Android in `apps/android`) consume, so
the two never drift from the site.

## Design tokens

`web/src/styles/tokens.css` is the source of truth. `node web/scripts/gen-mobile-tokens.mjs`
writes `apps/ios/Shared/DesignTokens.swift` and
`apps/android/core/designsystem/src/main/kotlin/gg/foreversixty/app/designsystem/Tokens.kt`
(colours including class, item-quality and parse colours, type scale, radii, spacing). Never edit
the outputs; `web/src/styles/mobile-tokens.test.ts` fails CI when they differ from a fresh run.
The site is dark only, so the apps are dark only.

## Fonts (`fonts/`)

The OFL faces the site self-hosts, as TTF (converted from the site's own Latin subsets in
`web/src/assets/fonts/`, covering ASCII, Latin-1 and the typographic punctuation the site
renders): Barlow 400/600/700 (body), Barlow Condensed 700/800 (headings), Cinzel 700 (the
wordmark only). Licences sit beside them (`OFL-Barlow.txt` covers Barlow Condensed too). iOS
bundles them and lists them under `UIAppFonts` (`FSFont` uses the PostScript names); Android
copies them to `res/font` with lowercase names.

## Brand (`brand/`)

The app icon is the LX seal from `design/logo/foreversixty-mark.svg`, unmodified, scaled and
placed only. `node web/scripts/build-app-icon.mjs` regenerates it.

- `AppIcon.appiconset/`: the iOS asset catalog set (single 1024 px opaque image; the OS masks it).
  Drop it into the app's `Assets.xcassets`.
- `android/`: adaptive icon layers. `ic_launcher_foreground.png` (the seal inside the 66/108 safe
  zone, transparent elsewhere), `ic_launcher_background.png` (flat `#07090d`) and the
  `ic_launcher.xml` adaptive-icon wrapper. Put the PNGs in `mipmap-xxxhdpi` and the XML in
  `mipmap-anydpi-v26`, or regenerate density sizes from the 1024 masters.

## Crests and icons: one crest language

Wherever a class or character is pictured, both apps show the circular ringed HD crest, never a
square class icon, a letter square or a Blizzard avatar (`design/DESIGN-SYSTEM.md`, tenet 2).
All art is served by the site's own origin, `https://foreversixty.gg`, never from the API:

| Art | Path | Notes |
| --- | --- | --- |
| Class crest | `/icons/hd/crests/<class-slug>.webp` | 128 px WebP, the client's class icon upscaled 4x, cropped to the art. Nine slugs: warrior, paladin, hunter, rogue, priest, shaman, mage, warlock, druid. Source: `web/public/icons/hd/crests/`, path built by `web/src/lib/class-crest.ts`. |
| Faction emblem | `/icons/hd/faction/<faction>.webp` (72 px), `<faction>-logo-64.webp`, `<faction>-logo-512.webp`, `<faction>-512.webp` | `alliance` or `horde`; flat logo at 16 px in row descriptors. Source: `web/public/icons/hd/faction/`. |
| Spec icon | `/data/<build>/icons/<icon>.webp` | `<icon>` is the `icon` field of `data/curated/specs.json` (e.g. `spell_nature_starfall`); `<build>` is the active build id (e.g. `1.60.1.70009`). |
| Item and spell icons | `/data/<build>/icons/<icon>.webp` | The same tree as spec icons. |

The ring is drawn by the app, not baked into the image: the whole crest is a circle (clip
`Circle`) with a 2 px ring in the class colour (`FSColor.class*` / `FsColor.class*`) at 55%
opacity at rest, full opacity when focused or pressed, over a `FSColor.raised` (`#0d111a`)
backing so a loading image never flashes transparent. This is exactly `ClassCrest.astro`
(`web/src/components/ClassCrest.astro`); the sizes the site uses are 36, 44, 56, 64 and 84.
Cache images by URL; the crest files change only on a deploy.
