# UX verdict: mobile follow-ups, 2026-10-31

Verdict: SHIP. No blockers.

## 1. Band mismatch (spec F2 Home, F3 BiS) - pass
- Obnoxious Yell, level 58: Home hero note reads "Best in slot for 50 to 59", the list header right label reads "50 to 59", BiS chip reads "50 to 59". iOS and Android agree. The "60 · raid" label is gone.
- Yellfury, level 31: Home reads "30 to 39" at both the hero and the list header. Same rule, a second data point.
- Not judged: the fixture's generic figures (12 upgrades, +102.5 DPS) repeat across characters.

## 2. Selector truncation (spec 3.4) - pass
- Default size: "Hunter · 31" is whole on iOS (raids) and Android (raids).
- Largest size, iOS AX5 (Raids, Home) and Android 2.0x (Raids, Home): the selector holds at its capped size, name and description both whole, pill hugs the text, nothing cut. The page content grows and the selector does not, as 3.4 now says.
- Spec 3.4 text matches what shipped (description line stays and is never cut).

## 3. Buff tiles (spec F4 raids composition) - pass
- Every tile shows its short name under the icon at a constant size, bold. Present buffs are strong white, the missing Blessings tile is muted with a dim border and matches "Missing: Blessings" below. Present and missing read apart at a glance.
- iOS wraps 7 + 3, Android 8 + 2. Tiles align on a grid, labels centred, no clipping, no layout jump between rows.
- Raid composition list below is unchanged.

## Blockers
None.

## Shoulds (optional)
1. MotW and Bless show an empty dark square. MotW is a present buff, so its green-bordered empty square reads as a failed image load. Export both icons in the pipeline and track it; until then the name under the tile carries it. Do not ship these two as a permanent state.
2. iOS BiS summary line wraps "Updated 30 / Oct 2026" and orphans "Oct 2026". Android fits it on one line. Keep the date together (no break inside the date) or shorten to "Updated 30 Oct".
3. Android Home at 2.0x: the "+21.2 DPS" figure drops below the detail text in the upgrade row, while iOS AX5 and normal sizes keep it at the right. Not caused by these fixes; log it for the large-text pass.
4. iOS selector at AX5 leaves a little more vertical padding than at default (pill 115 vs 100 px high in the capture). Acceptable; no action.

## Checklist
- [x] Home and BiS band labels agree for 58 (iOS, Android)
- [x] Home band for 31 reads 30 to 39
- [x] "Hunter · 31" whole at default, both platforms
- [x] Selector whole at AX5 / 2.0x, both platforms, Home and Raids
- [x] Spec 3.4 text updated
- [x] Every buff tile named; missing buff muted
- [ ] MotW and Bless art (should 1)
