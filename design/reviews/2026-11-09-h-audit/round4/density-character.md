# Density review: character page (visitor-60, owner-60-premium), round 4

Verdict: DENSE WITH ADDITIONS. Round 3 MUSTs are met in kind: the score spark with +310 since 20 Oct sits in Forever Score, and Best per encounter carries a parse spark and "+82 since 27 Oct". Handed-off and needs-data items excluded.

Page job: one character's standing and gear, and what to do next.

## Additions, ranked
1. Best per encounter: only Lucifron has a spark and delta. Magmadar, Gehennas, Garr and Baron Geddon have none. Render a spark and "+n since <date>" on every boss with two or more kills in /history encounters[].points. File: character/CharacterFightTable.svelte.
2. Boss name drills down to the report (report_id, fight_index in the history payload); no link affordance is visible on any row. Same file.
3. Visitor 1440: the right rail ends at y~1180 and the left at y~1417, leaving ~240 px empty under Activity. Add Standing (rank, tier) if the public rating endpoint carries it (lib/character/standing.ts), or the History score chart.
4. Owner: History is still a heading over two closed rows (y~1530-1690). No longer a MUST, but it is the weakest block left (CharacterHistory.svelte, ScoreChart.svelte).
5. Forever Score parts (Gear, Experience, Parses, Attendance, Preparation): bars with no delta per part; HistoryScorePoint first vs last gives it (CharacterScore.svelte). The sparkline has no hover value.
6. Builds seen: spec and date only; add fights per build from /history, link to the planner.

## Phone (390)
- Lucifron meta wraps to three lines and the spark floats mid-row. Put the delta beside the parse, keep the meta on one line.
- Forever Score now sits directly after Upgrades, Standing follows. Good.
- Gear list with item names, kept.
