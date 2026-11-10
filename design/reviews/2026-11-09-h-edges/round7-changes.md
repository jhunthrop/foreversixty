# Round 6 to 7: coordinator round (captures: round7/)

## API additions (data the DB already held; no migration)
- Each award's decision record (ms_os, reason, votes_yes, votes_no, votes_officers from loot_awards) on the night report's loot and the member profile's loot (award_record.go, raids_report_view.go, member_profile.go).
- loot.guild_median, the median award count across the guild's verified raiders; loot.nights_since, raids they came to since their last item; attendance.streak (member_profile_history.go).
- upgrades: the readiness gear gap named slot by slot, with the BiS pick, its source and its gain (bis.GapPiece; the guild-officer reviewer's one flipping item).
- openapi.yaml, docs/contracts/fixtures/app/guild.member.profile.json, web types and Go tests updated.

## Needs data (no table or column holds it)
- A raider's trial status and rank history (guild_trials is the guild's billing trial).
- A guild attendance or at-risk threshold.
- Officer actions from the profile (message, bench, demote): endpoints and Discord DM consent.
- Week over week per raider: report.moves exists but the fixture guild has one night.

## Shared label
- "Joined from the bench or without a sign-up" lives in web/src/lib/guild/attendance-copy.ts line 135 (attendanceCopy.attendance.groups.joined), a file lane f-guild changes; routed to the coordinator. The report's officer names read lane copy instead ("Not in the log", "Late", "Joined").
