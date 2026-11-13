# Tier share card review, ux-designer, 2026-11-13

Verdict: FIX (DPS card only; tank and healer ship as is)

Checked at 500 px wide (scale 0.417). Title 60 px renders at 25 px: legible. Spec names on the DPS card are about 18 px, so about 7.5 px at preview size: not legible. The tier letters (36 px, about 15 px) and the crests read fine.

## Must fix
1. DPS spec names: 18 px to 24 px, weight 500. Target at least 10 px in the preview.
2. DPS grid: 4 columns to 3 columns, each about 340 px wide. B and C become 2 rows of 3. The longest name, "Enhancement Shaman", is about 235 px at 24 px and fits.
3. DPS crests: 32 px to 40 px, with a 10 px gap to the name. Row pitch is 60 px (7 rows in about 420 px).
4. Meta line: 20 px to 24 px, colour to the muted-text token at 0.85 opacity or brighter. Shorten the DPS line to "Best in slot, raid buffs · level 63 boss · 3 min · 13 Nov".
5. "Damage" suffix: 30 px to 34 px, baseline-aligned with the title (keep), colour to --text-muted so it reads as secondary to the 60 px title.

## Nice to have
- Healer fight length is "110 s" and DPS is "3 min". Use one unit format. The tank card has no fight length.
- Tank rows are 138 px tall with empty space on the right. Add the sim score bar or a role icon.
- Tier C green looks neon next to the muted palette. Use the site's quality-green token.
- Panel top margin (165 px) is tight to the meta line. Add 8 px.
