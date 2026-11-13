# L7 Officer tab, round 1: roster, readiness and master loot session as built

Checked: the nine L7 renders, both boards, spec 4.6/8/9, OfficerTab/Officer/OfficerLoot lua, the old GuildTab and GroupPanel from git, a grep of the ReadyCheck API.

## Verdict: FIX

An officer would open this on raid day. The roster by sign-up, readiness, Remind and the Give-twice loot session beat the spreadsheet. The deleted group panel took one thing with it.

1. MUST: raid buff and debuff coverage is gone (the group panel's Coverage region). Add a row to Tonight's readiness from `coverageEntries`, names in the hover.
2. MUST: the in-raid member view has no home for a non-officer; confirm Guild or chat covers the raider's own lack.
3. SHOULD: a button to call out the not-ready.
4. SHOULD: approvals are a count only (no regression).
5. SHOULD: a 40-raider guild needs a pager past the cap.
6. SHOULD: Gap column blanks: a muted `-` for unknown.
7. NICE: Give's refusal reason in the footer.
8. NICE: Won is your own record only; say so in its hover.

Sort, filters, push confirm and claim state are all present.
