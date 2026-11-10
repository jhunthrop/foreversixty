# Seal v4: ux-designer, 2026-11-09

**Verdict: SHIP.** No MUSTs. The cut is clean, the rim whole, the seal round, and it reads at every size and on every ground.

Reviewed: sheet.png, header-after@2x.png, footer-after@2x.png, favicon-tab-after@2x.png, addon-icons-after.png, and the
full-res cut design/assets/logo/lx-seal-v4.png measured directly.

## The cut, measured

- Edge: the alpha ramp is 0, 75, 245, 255 (about 3 px at 1768 px): soft antialiasing, no stair-stepping.
- Halo: semi-transparent edge pixels average RGB (36, 24, 15); the solid rim beside them (38, 25, 15). Same dark brown:
  no ground bleeds in. Composited at 4x on #f3eee2 there is no tan fringe, only the dark rim against cream.
- Roundness: radius at alpha 128 runs 877 to 880 px (mean 878.6), under 0.4 % variation. Fills to 5 px from each edge
  with no clipping.
- Rim: intact all round. The rim-band denoise shows only as slight softness in a few brush streaks; the rune band and the
  inner dark ring are untouched.

## Sizes and grounds

| Surface | Result |
|---|---|
| Nav teal 512/128/64/32 | Pass. The dark rim holds the silhouette on #0b2f3b; LX reads at 64 and still at 32. |
| Footer brown 512/88 | Pass. Rim brighter and cleaner than v3, no ground-coloured fringe. |
| Light #f3eee2 | Pass. Crisp edge, clean shape; better than v3. |
| Header 40 px @2x | Pass. Balanced against the wordmark and the character crest. |
| Footer 88 px @2x | Pass. LX, rune band and teal field all legible. |
| Favicon tab, dark and light | Pass. Reads as a coin on both; LX just perceptible. |
| Addon icon 64 / minimap 32 | Pass. Ring and teal centre hold; at 4x nearest the minimap LX is slightly soft. |

## SHOULDs (non-blocking)

1. 16 px reads as a teal dot, not LX (parity with v3). A simplified 16 px mark for the .ico frame is an option if the
   owner wants the tab to carry the letters.
2. Contrast on teal is a touch lower than v3 (v3's drop shadow separated the lower-right rim). The dark rim carries it; if
   the header ever looks flat, `filter: drop-shadow(0 1px 2px rgba(0,0,0,.35))` on the 40 px seal.
3. Minimap 32 detail is soft; a sharpened downsample for the addon's 32/64 rasters only would crisp it.
4. A faint light fringe seen at 512 on brown on the sheet was not in the pixel data (display resampling); absent in the
   real footer @2x capture.
