# Round 10 UX verdicts: character, planner, rankings-you, report-deaths (1440 and 390)
Handed-off items excluded (Upgrades card, doll item level, client build footer, yellow dash marker, tier gates, Rankings tab collapse, "Your spec" pill).
Fixed since round 9: Lucifron row no longer wraps at 390 (date and Execution on one line, delta and sparkline on a second line); Report column is on rankings-you at 1440.

## character-visitor-60: SHIP
- Pass: paperdoll leads, compact title band, Upgrades plate, Builds seen plated, one score colour language, no em dashes.
- Polish, 1440: right rail ends y 1200, left y 1415; about 215 px empty under Activity.
- Polish, 1440, Lucifron: "+82 vs 27 Oct" and the sparkline float mid-row, detached from both the name and the score; right-align them against the score or move to a second line as at 390.

## character-owner-60-premium: SHIP WITH FIXES
- Must fix, 1440 and 390, Best per encounter footer: "History 7 fights" and "Coach" are still bare disclosure rows (tiny triangle, no count or state on Coach, no divider treatment). Third round open. Give Coach a count or state (for example "Coach 2 reviews") and render both as full-width rows with chevrons.
- Pass: Upgrade banner, Standing plate and Activity (Simmed 912 HPS) carry the next step; rails balanced at 1440 (both end y ~1580); Lucifron wrap fixed.
- Polish, 390: Upgrade banner is yellow text on dark with no chevron; make the whole banner a link with a chevron.

## planner: SHIP
- Pass: empty state is a real tree; Load the level 60 build is primary in the right rail at 1440; In the game plate leads with Import from addon.
- Polish, 390: Load the level 60 build still sits after Arms guide order (about y 1450 of 2480); the primary action should follow Talents.
- Polish, 1440: about 80 px dead under In the game against the Talents plate.

## planner-you: SHIP
- Pass: Send this build to Felsnap and Copy for the game lead In the game; Share and Import are full-width plates; Point order border solid.
- Polish, 1440: about 150 px dead in the right rail under In the game against Talents (was 90 px; the rail got shorter because Share and Import moved down).
- Polish, 390: Load the levels 30 to 39 build is below Point order (about y 1490); "26 points off" and the delta still drop under the name and lose the right-aligned scan column.

## planner-guide: SHIP
- Pass: Loaded state, Copy for the game, 51 of 51, Sim this build link, point deltas, Survival row marked Loaded.
- Polish, 390 and 1440: Loaded pill still grey on dark, low contrast; give it the teal state colour.
- Polish, 1440: Show DPS toggle sits about 7 px past the content column (right edge 1330 vs 1323). Third round open.
- Polish, 1440: here the rail holds In the game, Share and Import in one plate, while planner-you puts Share and Import full-width below; pick one layout.

## rankings-you: SHIP
- Pass: Report column present at 1440; the You row has a left accent and a You pill; Your rank 14th leads on phone above the board; percentile colours and the Executed column read as the Details! table they borrow from.
- Polish, 1440: right rail holds only Your rank (ends y 410) and leaves about 890 px beside the board.
- Polish, 390: the column header shows only "DPS" while rows carry an unlabelled percentile and date.

## report-deaths: SHIP WITH FIXES
- Must fix, 390: "This pull" has only Started; Sim this fight and Copy link, present at 1440, are gone. Keep both actions on phone.
- Polish, 390: tab strip clips "Damage ta" with no fade or chevron cue, and a stray "..." sits beside "5 deaths".
- Polish, 390: chart legend is six stacked text lines (about 180 px) before the chart appears; wrap two per row.
- Polish, 1440: rail holds only This pull (ends y 475); about 880 px empty beside the chart.
- Pass: first death and its last hits and heals lead above the fold, health-after bars, Killing blow tag, class-coloured names, no em dashes.
