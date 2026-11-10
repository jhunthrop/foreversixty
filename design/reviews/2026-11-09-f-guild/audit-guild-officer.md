# Audit (round 0): guild-officer-reviewer on the current guild pages

Captures: audit/build-<state>-<width>.png (390, 1440; every parity state but raids-plan, which failed to open).
Reviewer run as a general-purpose agent acting from .claude/agents/guild-officer-reviewer.md (the agent type is not registered in this session).

Fixture caveat (reviewer): GuildRosterTable already has Rating, Attendance, Best parse, professions and main/alt columns, but the fixture sends rating null, best_parse null, no professions, addon null and no attendance_30, so none render.

## Findings (condensed; ranks as given)
- Officer overview. BLOCKER: Action needed is counts with no context; no go/no-go read for Wednesday (sign-ups vs 40, comp shortfalls, readiness failing, attendance trend). BLOCKER: comp "Short" names no one; expand to fillers from bench/roster. SHOULD: cover gaps should name who can fill; next raid needs boss list, last night's kills/wipes; no "who's slipping / new" strip (attendance_30, logged_at, last_report_at). NICE: empty gutters.
- Member overview. BLOCKER: Your numbers is two figures; add best parse + percentile, rating components, gear gap and next upgrade, what I won, attendance vs guild median. SHOULD: my line from the last raid (DPS vs spec, deaths, drops for my class).
- Roster. BLOCKER: name list with item level and "1 of 1"; needs a dense row: ilvl vs median, gear gap, rating, best parse, attendance 30, professions, last logged, readiness dot, role groups. BLOCKER: counts disagree (24 raiders / 27 on roster / 19 of 40). SHOULD: collapse the approval block; phone rows waste space.
- Raids list. BLOCKER: trash pulls shown as "no named encounter published". SHOULD: night-over-night trend; red bar under target.
- Raid planner. SHOULD: sign-up rows carry no figures (ilvl, rating, attendance, readiness); bench swap has no recommendation; flag assignees not In; add cover row.
- Report. BLOCKER: DPS/HPS shows three people; needs the whole raid vs spec medians, deaths. SHOULD: consumables apology panel; group deaths by ability; render the Discord post.
- Progression. BLOCKER: one boss, page empty; attempts per boss, first kills, wipe causes, kill time trend.
- Readiness. BLOCKER: 4 short rows + ready names as a paragraph; needs a matrix of every raider by check (ilvl, gear gap, enchants, flask/food, attunement, talents, fire resistance, last logged), Nudge all. SHOULD: FR without target.
- Loot. BLOCKER: no fairness view (awards per raider, attendance-weighted, last item won). SHOULD: candidates as a comparison table; quest item noise; show voters.
- Profile. SHOULD: thin; add rating components, parse, trial/tenure, trend.
- Applications/recruit/public/landing. SHOULD: comp fit, gear vs roster median; public visitor dead end; directory bars carry no information.
- Settings. NICE: a usage line beside the plan.
- Cross-cutting. BLOCKER: every page is one or two cards over empty ground; no per-raider drill-down. SHOULD: no deltas vs last week; 390 tab strip clips.

## Value test (verbatim)
A 40-man guild would not pay $15 to $25 a month for the screens as captured. The single addition that flips it: a per-raider table (roster plus readiness plus attendance plus loot-received) as one sortable screen, with the numbers the code already supports, ending in "Nudge" and "Bench" actions. The loot-fairness history is the second flip: it is the part of the spreadsheet a council cannot give up.

## Needs data (reviewer)
Spec-median DPS/HPS per boss joined to guild reports; avoidable damage; weekly trend series for rating and parse; consumables by pull when the log has no sheet; next week's drops per planned raid; Discord reachability; attunement beyond the board raid.

VERDICT: NOT YET
