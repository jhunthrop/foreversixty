# Lane G2 AI-slop review

**First impression:** Reads as made by a player. The rankings rows are real names, real DPS and class-coloured bars. The home foot is honest SEO prose but still reads like a person wrote it, apart from the few lines below.

## Tells, most damaging first
1. **Tell 13, rankings-1440.png.** "Leaders by spec / Relic Guardian" shows "No ranked kill on this board yet." directly under a table of ranked kills. It is a placeholder that contradicts the page and pretends to be content. Fix: render the real leaders server-side, or render nothing until they load.
2. **Tell 7, PageHeader lead (rankings, planner, sim-weights).** A lead sentence under every title is now the page pattern. Several explain the obvious:
   - Planner: "Plan every point up to level 60 in the order you take them, check it against the guide builds, then copy it into the game or share the link."
   - Rankings: "ranked by DPS and HPS for each spec, class, ruleset and guild".
   Fix: keep the lead only where it says something a veteran does not know (the rankings uploads sentence). Cut the planner one to a line or remove it.
3. **Tell 21, home foot paragraph 1.** "with the quest, dungeon or vendor each piece comes from" and "a guide written for Forever's changes" are pitch-flavoured. Say "BiS per ten-level band, 20 to 60, with the source of each piece."
4. **Tell 22, home foot paragraph 4.** "Data from client build 1.60.1.70790, read from the game files on 11 Nov. BiS lists ranked 10 Nov." has the qualifier "read from the game files". It also repeats the Updated stamps used elsewhere. Fix: "Client build 1.60.1.70790" and one date.
5. **Tell 19/11 (mild), home foot paragraph 2.** It is a link soup of 13 yellow links in one sentence run. Fix: split into the sim tools and the tier list, or drop the links the nav above already carries.
6. **Tell 6 (mild), "Tanks / Healers / Damage" labels** in Leveling BiS by spec are tiny, small-caps labels. Acceptable as a real grouping; make them normal-case body weight.
7. **Rankings table, minor.** The "Executed" column is all em dashes, which reads as an empty column. "Report" is repeated on every row. Hide the Executed column when it has no data, and show Report on hover or in a row menu.

## Authentic, keep
- Rankings rows: class-coloured names, crest icons, percentile colour ladder, real 0 DPS rows left honest.
- Home spec lists in class colour, full names with "(Bear)", by role.
- Concrete copy: 9 Dec, Barrow Deeps, Hyjal Summit, Onyxia's Lair, boss names linked, build number.
- Server-rendered rankings match the JS render (nojs variants).
- No gradients, glows, emoji or icon-per-bullet grids in the new parts.

## Verdict
FAIL (tell 13: the "Leaders by spec" empty state beside populated rows). Fixing item 1 gives PASS WITH FIXES for items 2 to 7.
