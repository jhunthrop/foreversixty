# Player verdict: BUILT report charts (2026-10-30)

Reviewer: veteran Classic raider (Warcraft Logs daily). Seen: threat, damage, damage tooltip, resources, timelines at 1440; threat at 390. 2000 and the other 390 captures not opened. Players control not judged (temporary).

## Verdict: HOLD

The threat tooltip and the damage tooltip are right. Death rows say who, what hit them and the overkill ("Brokk died - Melee from Ragnaros - 3,862 (1,104 overkill)"). The tanks (Thandril, Brokk) are drawn on the 1440 threat chart. Several things a raider would screenshot and mock remain.

## Blockers

1. "Top 8 of 3". The Timelines and Resources boards (3-player fight) print "Players: Top 8 of 3" and the damage head reads the same. Clamp it ("All 3"), and hide the Top-N wording when N exceeds the roster.
2. The phone threat chart does not name the tank. At 390 the legend is "Kargath 150.8k, Durgan 147k, +38". The two tank lines sit unlabelled near the bottom, and there is no Players control on that board. The always-show-the-tank promise holds at 1440 only. The phone legend must pin the tanks (Thandril, Brokk) ahead of "+N". It should also keep the estimate wording; the phone foot reads "Damage and healing only", not the 1-per-damage / 0.5-per-heal line.
3. Damage defaults still hide the low performers, and the median line cannot be seen. The legend is Raid median plus the top 8 plus YOU, so every named line is at or above the median (377 against 410 to 444). Nobody below the median is drawn or named, and no "lowest 5" or "below median" preset exists. In the 1440 capture with Ashka hovered ("Only"), the dashed median is not distinguishable among the noisy lines. Make the median a heavier white dashed line, and add a "Lowest 8" choice so the slackers are one click away.
4. The damage numbers do not reconcile on sight. Head: "Raid 11,015 DPS". Legend: median 377 with the top at 444. Forty players at about 377 is about 15,100. It only adds up if the legend is in-combat DPS and the head is whole-fight DPS (the 90 s submerge at zero accounts for the gap). Neither the legend nor the tooltip column has a unit ("377", "672"), so a raider will call it broken. Label the unit ("DPS, 5 s average") and say which basis each figure uses.
5. The Resources tab draws the Damage chart ("DAMAGE DONE Raid 136 DPS") above the resource rows, and the rows have no title or column header. Rows read "Baelgrim Mana ... peak 8,000 low 8,000 at 8.0s at cap 0.0% of the fight" with a bare COPY CSV under each. Spec says one chart per tab. Put the resource chart or a header there, not a damage chart.
6. Death-to-Deaths link is not demonstrated. The static captures show the markers and tooltips but nothing that shows a click goes to the death (the earlier verdict required this). Confirm and capture the hover/focus state with "Open in Deaths", or this stays open.

## Shoulds

- Threat Share column is a percent of all 40 players ("4.0%" for the top, "1.2%" for Thandril). It never answers whether DPS are ahead of the tank. Add "vs tank" (for example 330%, with the 110%/130% pull lines at melee and range). Classic raiders think in those numbers.
- Every DPS is shown at about 3x the tank's estimated threat, so the chart implies the boss should be on Kargath all fight. The "Estimated" label and foot line cover it, but add "Tank threat is understated: no Defensive Stance 1.3x, Sunder, Righteous Fury" next to the tank pins, not only in the foot.
- Threat legend values (95,971 and so on) are end-of-window figures while the cursor sits at 2:07 (tooltip 77,145). Dim or label them "end of window".
- Sidebar says the selected pull is "Warden Kelthas 1:00", but the damage and threat boards show Ragnaros 5:40 and "40 rows - 5:40". Fixture mismatch; fix it so the captures can be trusted.
- Timelines: "Phase 1 / Phase 2" again; use encounter names. The damage lines on the timelines board are spiky triangles on a 1:00 fight, which is fine, but the "Elyra Duskvale 11" grey line is barely visible.
- Classic: Ragnaros submerge is correctly a flat gap; add a "no damage to Ragnaros" label there, since the raid is killing Sons of Flame.

## Reconciles
- Threat legend leader equals table leader (Kargath 95,971 both); Brokk 22,820 in tooltip equals the table. Taunt row (2:07) matches the Taunts list under the table.
- Brokk's death at 2:06 matches the red line on the damage chart and the "1 death" in the report head.

## Re-check (fix round 2)

Seen: threat 390, damage 1440, players menu 1440, resources 1440, damage tooltip with death (before), Deaths tab (after). Not opened: 2000 captures, timelines.

Verdict: SHIP.

1. "All M" clamping: fixed. Damage head reads "All 3" on the 3-player fight and "Top 4, bottom 4 of 40" on the raid.
2. Phone tanks: fixed. Legend reads "Thandril 45,325, Brokk 22,820, +38" with tank markers, and the foot keeps "An estimate: 1 per damage, 0.5 per heal. What it leaves out". Title now says "Estimated threat, every enemy".
3. Damage defaults: fixed. Top four (444, 432, 428, 426) and bottom four (327, 326, 322, 318) are named around a 377 median; the median is a visible white dashed line. Presets "Top 4 + bottom 4 / Top 8 / Bottom 8 / Custom" sit in the Players menu. Healers/tanks are separate groups; a "28 damage dealers" median is stated.
4. Units: fixed. "Raid 11,015 DPS over the whole fight", every chip "DPS", tooltip "DPS, 5 s average". Sums plausibly (28 dealers around 377).
5. Resources: fixed. No damage chart; header row Player / Power / Over the fight / Time empty.
6. Death link: demonstrated. Tooltip shows "Thalgrit died - Anima Lash from Warden Kelthas - 2,400 (100 overkill)"; the after capture lands on the Thalgrit death card with the 20s-before and Sim buttons.

Nothing new that blocks. Minor, not blockers:
- The before-capture tooltip has no visible "Open in Deaths" cue; a raider will not know it is clickable.
- Resources rows: Baelgrim appears twice (Mana, Rage) with tiny grey stat text; the "wasted" dagger footnote is dense.
- Still open from my earlier shoulds: "vs tank" share, tank-understated note, end-of-window legend label, sidebar/fixture mismatch (Warden Kelthas 1:00 vs Ragnaros 5:40 in the raid captures).
