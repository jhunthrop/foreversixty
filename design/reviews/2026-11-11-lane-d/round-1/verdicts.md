# Lane D, round 1 verdicts (captures of eca7ae85e / 93cf7857d)

| Reviewer | Verdict | Must-fix |
|---|---|---|
| ux-designer | FIX | 404 picker names in gold, not class colour; 404 listed Feral Bear where home did not |
| wow-player | FIX | Greenhammer black icon; "signed up" counts differ plaque vs raid page; "2 short" vs "2 missing" |
| ai-slop-reviewer | FAIL | Greenhammer black icon; "In" Tag repeated on 27 of 29 raid rows |
| info-density-reviewer | FAIL | "Last 5" head over an empty column; plaque does not say the role it signs as; "short" vs "missing" |
| guild-officer-reviewer | FIX | "Last 5" head empty; Ready dashes; "Can't come" left in Button.svelte; no test for the Move to In failure line |
| brand-consistency-reviewer | FAIL | notify/embed.go sim thumbnail still the bare web crest |

## What round 2 changed (c2b6f67a6)
- The black item and spec icons were the fixture build's 1x1 stubs, not a product bug: the capture now serves real icons (support/real-icons.ts), as the frame and guild captures do.
- PlannerRow: a Tag only for Late, Bench and Out; In is the default.
- PlannerRoster: the Last 5 head is blank when no row has nights.
- One `raidCopy.missing(n)` word on the raid page, the roster and the readiness board.
- HomeRaidPlaque: "as DPS" beside the four choices.
- ClassPicker: class names in their class colour; every picker lists one spec per tree (`treeSpecs`).
- Button.svelte comment reads In, Late, Bench, Out; `trySignup` test for the lock race behind `guild-move-error`.
- notify: the sim-finished thumbnail is `sheet.CrestURL(apiBaseURL, class)`; `crestPathFormat` deleted; a Go guard fails on any bare `icons/hd/crests` URL outside sheet.

## Out of lane, left as is
- "18 signed up" (home fixture) vs "29 of 40" (raid fixture) are two fixtures; the plaque's figure is `signupFigures`, unchanged by this lane.
- The Ready dash is the existing "not checked" mark.
