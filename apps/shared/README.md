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

The OFL faces, as TTF: Open Sans 400/600/700 and Archivo 600 (`make brand-fonts` converts the
site's own Latin subsets in `web/src/assets/fonts/`; licences `OFL-OpenSans.txt`,
`OFL-Archivo.txt`) and Marcellus 400 (`Marcellus-Regular.ttf`, the unmodified upstream file from
google/fonts, so the OFL's Reserved Font Name clause holds; licence `OFL-Marcellus.txt`), the
titles' face. These five are the apps' faces (`design/tokens/tokens.json` `font`, round 5): iOS
bundles them and lists them under `UIAppFonts` in the app and the widget extension (`FSFont`
names them by PostScript name, which is also each file's name); Android's `forever.shared-fonts`
plugin copies exactly these to `res/font` with lowercase names (`FsFont`). Barlow 400/600/700,
Barlow Condensed 700/800 and Cinzel 700 (licences `OFL-Barlow.txt`, `OFL-Cinzel.txt`) stay here
only for the older design mocks (`design/mocks/2026-11-04-sim-results`, `2026-10-31-paperdoll`)
and the Discord assets (`design/community/discord`); neither app ships them.

## Brand (`brand/`)

The app icon is a painted tile: the owner's blue ground (`design/assets/art/app-icon-background.jpeg`,
inside its brass frame, as the companion's macOS icon) with the painted seal v4
(`design/assets/logo/lx-seal-v4.png`) centred at 78% of the visible tile, scaled and placed only.
`node web/scripts/build-app-icon.mjs` regenerates every file below; never place one by hand.

- `ios/Brand.xcassets/AppIcon.appiconset/`: the iOS app icon (single 1024 px full-bleed image, no
  alpha channel; the OS masks it). `apps/ios/project.yml` compiles the catalog in place. The Dark
  and Tinted appearances are the flat seal (`design/assets/brand/seal-flat/seal-flat-small.svg`) on
  a transparent ground, Tinted in greyscale; the light appearance is the painted tile.
- `android/`: adaptive icon layers at 432 px (108 dp at xxxhdpi). `ic_launcher_background.png`
  (the painted ground over the whole canvas), `ic_launcher_foreground.png` (the seal inside the
  66/108 safe zone, transparent elsewhere), the `ic_launcher.xml` adaptive-icon wrapper and its
  `ic_launcher_round.xml` alias, and `legacy/mipmap-<density>/` 48 dp square and round PNGs. The
  `forever.app-icon` Gradle plugin copies the layers to `mipmap-nodpi`, the XML to
  `mipmap-anydpi-v26` and the legacy buckets as they are at build time. The wrapper's
  `<monochrome>` themed layer and the notification small icon are vector drawables in the app's
  `res/drawable`, both the flat tiny-bold seal with the LX knocked out, emitted by
  `uv run tools/gen_seal_silhouette.py`.

## Crests and icons: one crest language

Wherever a class or character is pictured, both apps show the circular ringed HD crest, never a
square class icon, a letter square or a Blizzard avatar (`design/DESIGN-SYSTEM.md`, tenet 2).
One scoped exception (owner, 2026-11-08): the profile button on the tab roots and the menu's character rows
show the character's Battle.net portrait (`avatar_url` on `/v1/me` characters) in a ring of the faction colour
with the faction emblem as a badge; without a portrait, the person glyph in the same ring and badge, never the
class crest. Everywhere else, Home's header, widgets and the Live Activity included, stays the crest.
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
