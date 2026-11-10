# Round 6 verdicts: final (captures: round6/)

- ux-designer: VERDICT: SHIP
- wow-player: VERDICT: SHIP
- ai-slop-reviewer: VERDICT: PASS WITH FIXES
- info-density-reviewer: VERDICT: DENSE WITH ADDITIONS
- guild-officer-reviewer (profile, report; last read round 4): VERDICT: CLOSE

## Open, for the coordinator
- ai-slop holds PASS WITH FIXES only on the painted ground and band (tells 2 and 12, the owner's system art, not this lane's) and one shared attendance group label ("Joined from the bench or without a sign-up", attendanceCopy, used by the raid page too).
- guild-officer CLOSE: wants loot reasons, streaks, at-risk flags, loot fairness, trials and officer actions; none of those fields are in the member-profile or report API yet.
- Every guild parity state (on main too) fails only on the site-footer box (mock about 60 px above the build): the shared GuildShell view height, not a lane file.
- The report's Discord summary is still an open section on the ground: ReportDiscord belongs to lane f-content.

## Density additions to consider later
- DPS and HPS share bars, sort, Copy CSV and fight-log links; profile last_report_at, best-parse log link, loot night links; a Join rail (recruit needs, officers); admin testers sort and filter.
