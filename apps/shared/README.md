# apps/shared

Assets and tokens both native apps (iOS in `apps/ios`, Android in `apps/android`) consume, so
the two never drift from the site.

## Design tokens

`design/tokens/tokens.json` (W3C design tokens) is the one hand-edited token file for the site
and both apps (spec `design/specs/2026-11-07-mobile-redesign.md` section 2.1).
`node design/tokens/build.mjs` writes `web/src/styles/tokens.css`,
`apps/ios/Shared/DesignTokens.swift` and
`apps/android/core/designsystem/src/main/kotlin/gg/foreversixty/app/designsystem/Tokens.kt`:
colours (class, quality, quality text, faction, tier, parse, school), the faces, the type roles
(`FSType` / `FsType`: face, size, line, tracking, the Dynamic Type style), space, size, the touch
target, radius (and `FsShape`), stroke, elevation and motion. Never edit the outputs; each opens
with a GENERATED line, and `node --test design/tokens/tokens.test.mjs` (the web unit job and `mobile-codegen.yml`)
fails when one differs from a fresh run or lost that line. The site is dark only, so the apps are
dark only.

## Copy (`copy/`)

`copy/en.json` is the apps' catalog; `copy/shared.json` holds the words the site shows too (keys
`shared.*`, the same value grammar). `node web/scripts/gen-mobile-copy.mjs` generates both into
`Localizable.xcstrings` + `Copy.swift` (`Copy.shared*`) and `strings.xml` (`R.string.shared_*`);
the web copy modules read `shared.json` through `web/src/lib/copy/shared.ts`, so a word the site
and an app both show is edited once. A key lives in exactly one of the two files (the generator
fails otherwise).

## Fonts (`fonts/`)

The OFL faces the site self-hosts, as TTF (converted from the site's own Latin subsets in
`web/src/assets/fonts/`, covering ASCII, Latin-1 and the typographic punctuation the site
renders): Barlow 400/600/700 (body), Barlow Condensed 700/800 (headings), Cinzel 700 (the
wordmark only). Licences sit beside them (`OFL-Barlow.txt` covers Barlow Condensed too).
Marcellus 400 (`Marcellus-Regular.ttf`, the unmodified upstream file from google/fonts, licence
`OFL-Marcellus.txt`) is the screen title's face from the flat redesign
(`design/specs/2026-11-07-mobile-redesign.md`); the apps wire it in that spec's lane T, which also retires
Barlow Condensed and Cinzel from the apps. iOS
bundles them and lists them under `UIAppFonts` (`FSFont` uses the PostScript names); Android
copies them to `res/font` with lowercase names.

## Brand (`brand/`)

The app icon is the LX seal from `design/logo/foreversixty-mark.svg`, unmodified, scaled and
placed only. `node web/scripts/build-app-icon.mjs` regenerates it.

- `ios/Brand.xcassets/AppIcon.appiconset/`: the iOS app icon (single 1024 px opaque image; the
  OS masks it). `apps/ios/project.yml` compiles the catalog in place.
- `android/`: adaptive icon layers. `ic_launcher_foreground.png` (the seal inside the 66/108 safe
  zone, transparent elsewhere), `ic_launcher_background.png` (flat `#07090d`) and the
  `ic_launcher.xml` adaptive-icon wrapper. The `forever.app-icon` Gradle plugin copies the PNGs
  to `mipmap-nodpi` and the XML to `mipmap-anydpi-v26` at build time.

## Crests and icons: one crest language

Wherever a class or character is pictured, both apps show the circular ringed HD crest, never a
square class icon, a letter square or a Blizzard avatar (`design/DESIGN-SYSTEM.md`, tenet 2).
All art is served by the site's own origin, `https://foreversixty.gg`, never from the API:

| Art | Path | Notes |
| --- | --- | --- |
| Class crest | `/icons/hd/crests/<class-slug>.webp` | 128 px WebP, the client's class icon upscaled 4x, cropped to the art. Nine slugs: warrior, paladin, hunter, rogue, priest, shaman, mage, warlock, druid. Source: `web/public/icons/hd/crests/`, path built by `web/src/lib/class-crest.ts`. |
| Faction emblem | `/icons/hd/faction/<faction>.webp` (72 px), `<faction>-logo-64.webp`, `<faction>-logo-512.webp`, `<faction>-512.webp` | `alliance` or `horde`; flat logo at 16 px in row descriptors. Source: `web/public/icons/hd/faction/`. |
| Spec icon | `/data/<build>/icons/<icon>.webp` | `<icon>` is the `icon` field of `web/src/data/specs.json` (e.g. `spell_nature_starfall`); `<build>` is the active build id (e.g. `1.60.1.70009`). |
| Item and spell icons | `/data/<build>/icons/<icon>.webp` | The same tree as spec icons. |

The ring is drawn by the app, not baked into the image: the whole crest is a circle (clip
`Circle`) with a 2 px ring in the class colour (`FSColor.class*` / `FsColor.class*`) at 55%
opacity at rest, full opacity when focused or pressed, over a `FSColor.raised` (`#0d111a`)
backing so a loading image never flashes transparent. This is exactly `ClassCrest.astro`
(`web/src/components/ClassCrest.astro`); the sizes the site uses are 36, 44, 56, 64 and 84.
Cache images by URL; the crest files change only on a deploy.
