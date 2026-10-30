# Sample chart: player verdict, 2026-10-30

Reviewer: veteran Classic raider. Seen: 1440, 2000, 390 and the 1440 10 s pull captures, plus the lane report.

## Verdict: HOLD

The full chart does the job of looking like real tooling. It has class-coloured lines, a legend with DPS, death and cooldown markers, a brush and a log switch. A player who has used Details! or Warcraft Logs reads it in two seconds. The 1 s change is right in principle. Two things on screen would get mocked, and one of them is a known bug in the exact feature being shipped.

## Question this surface answers
"What will my own report give me if I upload?" The chart is the first thing under the header, so it answers quickly. It answers "a nice chart" and not "what should I fix", which is the site's job.

## Blockers

1. **Brush handle on a short pull stops short while Window says "Whole fight".** On the 9.9 s pull, the right handle sits at about x=1245 and the track ends at about x=1331. That is about 85 px of the fight, read as cut off. The lane report lists it under Concerns and leaves it unfixed. The 9.9 s pull is the second row in the Fights list, so the first player who tries the new 1 s detail will hit it. Fix it before shipping the short-pull change.
2. **Headcounts that cannot be true in Classic.**
   - Chart: "Top 4, bottom 4 of 90".
   - Table: "265 players".
   - Fights list: "79 kills" on a 6:19 pull and "203 kills" on 12:43.
   - A raid is 40 players at most, and the top player did 1,374 damage in total.
   - A raider reads this as a broken parser in the first five seconds.
   - Either filter the sample to the actual group, or say what the count is (for example "everyone in range of the fight"). Label "kills" as what it counts.
   - Accuracy reviewer: please check the 265, the 90 and the kills-to-damage ratio. At 203 kills in 12 minutes, "Raid 11 DPS" does not add up.

## Shoulds

- **The default view reads as dead.**
  - "Raid median 0 DPS" and the top player at "2 DPS".
  - The bottom 4 are four flat lines on the baseline, drawn under the dashed median.
  - A prospect sees a chart where nobody does anything.
  - The disclaimer ("so the numbers are small") helps. For the sample, show the top 8 and drop the bottom 4, or drop the median line when it is 0.
- **Raid numbers disagree.** "Raid 11 DPS" in the subtitle and "Raid median 0 DPS" in the legend sit one line apart. Name them "Raid total" and "Median player".
- **The y-axis has no unit.** It shows 15 / 10 / 5 / 0 with nothing next to them. Put "DPS" on the axis or the top tick. The same goes for the phone strip, which shows no unit.
- **Markers have no key.** The skull, "3" circles and lightning icons mean nothing until someone hovers. Add a one-line key under the legend, or a tooltip a player can find.
- **"0 pulls" in the Fights header** sits next to a list of 79 pulls. If it is a selected-pulls count, label it "0 selected".
- **1 s on a 10 s pull is useful in a limited way.**
  - It shows burst timing, for example the 20 DPS spike at 0:04.
  - The line is drawn as straight ramps through zero between points. That makes triangles and implies damage ramping, which did not happen. Use steps, or dots on the points.
  - Between 30 s and about 2.5 min the 2 s points are noisier than the old 5 s average (the report says so). That is fine for burn windows. Keep the 5 s average on long trash, as shipped.
- **Nothing on the sample says what to do next.** The site's line is "Play your class better". Under the chart, add one finding of the kind a real report gives, for example "Soren is active 40% of the pull". Without it the sample sells the chart and not the help.
- **Phone strip.** There are no names, no legend and no unit, so the player cannot tell whose line is whose. A one-line legend under the strip fixes that. The flask and link icons have no labels.
- **Duplicate class colours.** Three orange lines (druids) and two cyan lines (mages) look identical. That is the Details! convention, so it is acceptable. Keep the tooltip name obvious.

## Data I do not believe
See blocker 2. Also: no sample fight has a boss, so every sample view is trash. The sample therefore never shows what a raider uploads for.

## Re-check (both blockers)

Seen: sample-chart-1440, -1440-10s-pull, -1440-everyone, and the fix-round section of the lane report.

Verdict: SHIP. No blockers remain.

1. **Brush end handle: fixed.** On the 9.9 s pull the right handle sits on the end of the track, and the end label reads "0:09.9". Window reads "Whole fight". The report's cause (range inputs with step 1000) is a believable reason for the old 9000 snap, and the e2e test checks the browser's own value against max.
2. **Counts say what they count: fixed.**
   - Default Source is "Your group". The Players control reads "All 4" and the subtitle reads "Group 3 DPS over the whole fight". The legend says "Group median".
   - Under "Everyone in range" the control reads "Top 4, bottom 4 of 90 players seen" (12:43 pull) and "of 26 players seen" (9.9 s pull). "Players seen" is honest for an open-world log, and the default no longer shows 90.
   - The Players table header reads "4 players · 0 deaths · 12:43".
3. **Trash kill counts: hidden.** The Fights list shows only the length and start time, and "79 pulls" is on the group row. No "kills" figure appears anywhere I looked.

New, not blocking:
- The Fights header still reads "0 pulls" beside "79 pulls". Label it "0 selected" or drop it.
- "Median player 0 DPS" remains under Everyone in range. Rename it "Median player" only when it is non-zero, or leave it.
- The 9.9 s pull still draws straight ramps through zero between 1 s points (player should from the first round).
- Y-axis unit, marker key, a next-step finding under the chart and the phone legend are still open. The report lists them as not done. They are shoulds, not blockers.
- The default group view is four players on one pull. That is honest, but it is the thinnest view the sample shows. Fine for launch.
- Copy still says "raid" in the Debuffs note under a group scope (the report's own concern). Fix when that tab is next touched.
