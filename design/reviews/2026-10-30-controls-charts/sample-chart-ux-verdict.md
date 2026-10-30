# Sample chart: UX verdict, 2026-10-30

Reviewed: sample-chart-1440.png, -2000.png, -390.png, -1440-10s-pull.png, against the lane report and the controls-charts spec.

## Verdict: HOLD

One blocker (the brush handle). Everything else ships.

## Blockers

1. Brush right handle stops short of the track end on short fights.
   - Seen in sample-chart-1440-10s-pull.png: Window reads "Whole fight", the brush label reads 0:00 to 0:09, and the right handle sits at about 89% of the track (x 1245 of 1331).
   - The control says one thing and shows another. The window is the whole 9.9 s, the handle says 9 s.
   - Fight 2 is one click from the sample's default, so the owner will meet it.
   - The lane's diagnosis (brushShape rounding to whole seconds, not the fight's milliseconds) is plausible. Fix in the chart lane: handles and end label take the fight's exact end in ms; the default window is [0, end] and the handle sits at 100%.
   - Add a test for a non-integer-second fight.

## Judgements

- 1440 and 2000: the sample now shows the charts well. Title, units, legend, presets, log switch, event markers, 220 px plot and brush all draw like a real report. At 2000 the report stays centred and the chart keeps its width. Pass.
- 390: the strip is right. It keeps the damage rows on the first screen, the y labels are legible at 15/10/5/0, and the axis reads 0:00, 5:00, 10:00. Pass.
- 10 s pull at 1 s points: it reads. Ten points make a zigzag, but the peaks are real and the 2 s axis labels match the points. Every series touching 0 at alternate ticks is plausible for a few swings in a 10 s trash pull. Keep 1 s. Do not smooth it back to a 5 s average.

## Shoulds

1. The desktop fold now holds no damage row (owner's call, noted by the lane). Add a short caption under the sample line or accept it. Do not shrink the plot.
2. Legend on the sample at 1440: three wrapped rows plus "+82 MORE" before the plot. Cap the visible entries at one row on the sample so the plot starts higher.
3. The legend DPS values (0, 1, 2) look like a broken chart. The note already says the numbers are small. Consider "per second" in the head units line, which the 10 s capture already does ("Raid 49 DPS").
4. The 10 s capture's line is clipped at the right edge on a rising slope past the last label (0:08). Extend the axis to the end value or end the line on the last point.
5. The 30 s to 2.5 min range at 2 s points is noisier than the old average (lane concern). Capture a 60 s fight and judge it before this ships to real reports.
6. The sample-chart-390 capture is 780 px wide (2x). Say so in the filename or capture at 1x, so reviewers compare at the right scale.
7. The header font-swap CLS (0.0004) is outside this surface. Track separately.
