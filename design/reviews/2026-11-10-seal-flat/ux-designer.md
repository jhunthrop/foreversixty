# ux-designer: SHIP (2026-11-10, lane/seal-flat)

Must-fix: none.

Checked: header 1440 (40 px) and 390 (34 px) at DPR 2, small seal crisp where the painted seal was muddy, no layout shift (width/height unchanged); sealFor thresholds (tiny <=24, small 25-44, painted above); Base.astro favicon links (ico declared 32x32 so Chrome prefers the SVG; Safari falls back to the ico; apple-touch stays painted); favicon-32.png, seal-64.png and FAVICON.png32 have no remaining references; companion wordmark clean at DPR 2; minimap button at 20 px beats the painted blob; tray 24-48 px clear, dot unchanged.

Shoulds (optional):
1. 16 px tray frame is still soft and the "on" dot hides most of the X: a physical limit; a heavier LX or smaller dot at 16 px only would help.
2. Check the minimap button in game beside LibDBIcon neighbours.
3. No captures on a light surface or forced-colours mode (hardcoded SVG fills are acceptable).
4. favicon.svg and seal-flat-tiny-bold.svg are one file copied twice; a test asserting equality would stop drift.
