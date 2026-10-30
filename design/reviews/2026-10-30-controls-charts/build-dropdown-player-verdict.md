# Dropdown build: player verdict (2026-10-30)

Verdict: HOLD

## Blockers
1. Rankings filter labels are cut off in the closed fields at 1440: "Every cl...", "Every s...", "All rules...", "All regio...", "No encounter...". The old selects showed the full word. A player cannot tell the Class field from the Spec field without the label above it. Widen the fields or shorten the text to "Any class", "Any spec", "All". At 390 the filters collapse into one summary line, which reads fine.
2. The sim Fight style list runs off the bottom of the screen. "Cleave, 2 targets" is clipped and the rows after it (any further fight styles, and the third and later options) cannot be seen without scrolling the page. The list is shown opening downward with about 300 px of room. The flip rule or the list's height cap is not working for tall description rows. The list must fit or scroll inside itself.
3. The report Source list is inconsistent. Baelgrim, Morrowlyn and Sunwick get a class crest, but Thalgrit (under Tanks) and "Elyra Duskwalk · Hardcore" have none and no class colour. Names are coloured only for the two Damage rows (gold and cyan), while Sunwick (Healers) is white. The chart legend shows the same players as "Elyra Duskvale" (the list spells it "Duskwalk"), which is a data or label mismatch and looks like a typo to a player. Either every player has a crest and class colour, or none does.
4. Planner Race list for a Warrior shows "Skyborne", which is not a Classic race and is not in the Forever race list I was given. If it is a real Forever race, fine; if not, it is a wrong option. The accuracy reviewer must confirm.

## Shoulds
- Fight style descriptions: "no parse has measured the cost of that yet" in Light and Heavy movement is hedging a player does not need in a picker. Say "Melee leaves range 5 s every 45 s." Also the check mark overlaps the description on the selected row.
- Report Source: the options are "All friendlies" and "All enemies" first, then Tanks, Healers, Damage. Fine. The "You" tag is not visible in the capture, so I cannot confirm it. Enemies group is missing (noted in the lane report); a player will ask for boss as a source.
- Planner Race list has no faction grouping or crests, so faction is hidden until you read "Faction follows your race". Wowhead orders Alliance then Horde; here they interleave. Acceptable but add faction colour dots.
- Region list order: All regions, US, EU, KR, TW, CN. Fine for Classic. Ruleset and Phase option lists were not captured open; check them.
- The More menu is plain text, 7 items, no grouping. Fine, but "Rating" and "Resources" first reads odd to a player who expects Buffs and Casts first.
- The Window select on the report chart still looks like a plain native control in the captures (not the new face), and so does "Enemy: Every enemy" under the threat table. The lane report says those two are the chart lane's, but a player will see two different dropdown styles on one page.
- At 390 the capture shows only closed fields, so the phone picker was not reviewed.
- The home search Region and Ruleset fields are fine (US, Normal), but the list was not captured open.

## Re-check (2026-10-30, after fix round 1)

Verdict: SHIP

Fixed, seen in the captures:
- Rankings labels are whole at 1440: "Every class", "Every spec", "All rulesets", "All regions", "Beta", "Both", "Damage". The Class list opens with crests and class colours.
- Fight style list: it now opens upward, stays on screen and scrolls inside itself (the "Cleave, 3 targets" row is cut at the panel edge, which reads as a scrolling list).
- Source list: every player row keeps the crest slot; Thalgrit and "Elyra Duskvale · Hardcore" show an empty neutral ring (agreed with the coordinator). Duskwalk was a misread; Skyborne is accepted as a real Forever race.
- Chart Players control: Compare group with descriptions ("Median DPS of the 28 damage dealers"), Tanks/Healers groups with crests and class colours, field reads "Top 4, bottom 4 of 40", presets in the foot. Reads like a real raid-log picker.

Remaining shoulds (not blockers):
- Players list figures ("133", "67", "30") have no unit in the rows; the legend says "DPS". Add the unit or a "DPS" column head.
- Fight style copy still says "no parse has measured the cost of that yet" (sim's copy); say "Melee leaves range 5 s every 20 s."
- Source list has no Enemies group (a player will want the boss as a source).
- Fight style list cuts a row mid-text with no scroll cue.
- Planner race faction dots remain a should.

Remaining blocker: none.
