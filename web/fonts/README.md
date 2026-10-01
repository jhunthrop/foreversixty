# Font subsets

`glyphs.txt` in this directory is the explicit glyph inventory for the five self-hosted font
subsets committed under `web/src/assets/fonts/`: Cinzel 700, Barlow 400/600/700, JetBrains
Mono 500. `web/src/styles/fonts.css` declares `@font-face` rules against those subset files
directly (no more `@import '@fontsource/...'`); see that file's header comment for why.

## Why

Lighthouse's simulated-throttling LCP budgets in `web/lighthouserc.json` charge for every
byte fetched before the hero text paints. On every page that is these five font files, ~105
KB total, because `fonts.css` loads all five regardless of page. Fontsource's `latin`
subsets already drop non-Latin scripts, but still carry glyphs this English-only site never
renders (combining diacritics, the OE ligature, guillemets, the Euro and trademark signs,
dotless i, breve/ring/tilde as standalone characters). Subsetting to exactly the glyphs the
site renders removes that dead weight without touching hinting, kerning, or any other
OpenType layout feature.

## Inventory

`glyphs.txt` is built from three layers, each documented inline in the file itself:

1. Printable ASCII (`0x20`-`0x7E`).
2. The full Latin-1 Supplement block (`U+00A0`-`U+00FF`) — letters and punctuation, per the
   explicit spec this inventory was built against, independent of what today's content
   happens to use.
3. General punctuation used in copy (en/em dash, curly quotes, ellipsis, bullet, narrow
   no-break space, minus sign), plus every character
   `web/scripts/check-glyph-coverage.mjs` found when it scanned a fixture build of `dist/`
   (both the server-rendered HTML and the client JS bundles — several pages render their
   copy from Svelte components that only exist in the latter), `src/content/**/*.md`, the
   game-data JSON (`data/builds/<build>/*.json`, `data/curated/**/*.json`): the rightwards
   and leftwards arrows, the dagger, and the single angle quotes.

Run the checker after any copy or data change that might introduce a new character:

```sh
cd web
npm run build   # FOREVER_DATA=fixture, so dist/ matches what check-glyph-coverage.mjs expects
npm run check:glyphs
```

It exits non-zero and names the offending character and source file if something the site
renders is missing from `glyphs.txt`.

## Regenerating the subsets

Requires [fontTools](https://github.com/fonttools/fonttools) (`pyftsubset`). From a scratch
virtualenv (or any environment with `fonttools` installed — this repo's `data/` Python
project does not depend on it, so a venv is cleanest):

```sh
python3 -m venv /tmp/fonts-venv
/tmp/fonts-venv/bin/pip install fonttools==4.61.1 brotli
```

Then, from `web/`, for each of the five faces:

```sh
VENV=/tmp/fonts-venv/bin
$VENV/pyftsubset node_modules/@fontsource/cinzel/files/cinzel-latin-700-normal.woff2 \
  --text-file=fonts/glyphs.txt --flavor=woff2 --layout-features='*' \
  --output-file=src/assets/fonts/cinzel-700-subset.woff2

$VENV/pyftsubset node_modules/@fontsource/barlow/files/barlow-latin-400-normal.woff2 \
  --text-file=fonts/glyphs.txt --flavor=woff2 --layout-features='*' \
  --output-file=src/assets/fonts/barlow-400-subset.woff2

$VENV/pyftsubset node_modules/@fontsource/barlow/files/barlow-latin-600-normal.woff2 \
  --text-file=fonts/glyphs.txt --flavor=woff2 --layout-features='*' \
  --output-file=src/assets/fonts/barlow-600-subset.woff2

$VENV/pyftsubset node_modules/@fontsource/barlow/files/barlow-latin-700-normal.woff2 \
  --text-file=fonts/glyphs.txt --flavor=woff2 --layout-features='*' \
  --output-file=src/assets/fonts/barlow-700-subset.woff2

$VENV/pyftsubset node_modules/@fontsource/jetbrains-mono/files/jetbrains-mono-latin-500-normal.woff2 \
  --text-file=fonts/glyphs.txt --flavor=woff2 --layout-features='*' \
  --output-file=src/assets/fonts/jetbrains-mono-500-subset.woff2
```

Notes on the flags:

- `--layout-features='*'` keeps every OpenType layout feature (kerning, mark positioning,
  the GSUB closure that pulls in a handful of glyphs not explicitly requested — see
  `glyphs.txt`'s note on `‹`, pulled in alongside `›`). Dropping to a narrower feature set
  would save more bytes but was explicitly out of scope for this subsetting pass.
- No `--no-hinting`: Barlow's `glyf` table carries real per-glyph TrueType hinting
  instructions (verified with fontTools — 212 of 252 glyphs in the 400 subset have a
  non-empty instruction program); Cinzel and JetBrains Mono do not (0 instructed glyphs in
  either source file, confirmed before subsetting), so there is nothing for them to lose
  either way. `--no-hinting` would have silently discarded Barlow's.
- No explicit `--unicodes`: `--text-file` takes literal characters, which is what
  `glyphs.txt` is, and avoids a lossy round-trip through code-point lists for anyone
  reading or editing the inventory.

After regenerating, verify metrics didn't move (they shouldn't — subsetting never touches
`hhea`/`OS/2`) and that nothing fell out of the inventory:

```sh
npm run check:glyphs
npx vitest run src/assets/fonts/fonts.generated.test.ts
```

`web/src/assets/fonts/fonts.generated.test.ts` pins each subset's cmap against
`glyphs.txt`, with one documented exception list for characters absent from the _source_
`@fontsource` files themselves (`→` `←` `†` and the narrow no-break space `U+202F` — none of
Cinzel/Barlow/JetBrains Mono's `latin` files have ever had these glyphs, subsetting or not).
