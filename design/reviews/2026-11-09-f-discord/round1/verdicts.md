# Round 1 verdicts (lane f-discord, 2026-11-09)

Captures: the commit f88044681 boards (`../discord/`, `../sheet-*.png`, `../card-*.png`).

## ux-designer: VERDICT: BLOCK
- BLOCKER: sheet text unreadable at Discord's display width (sub-lines and gains about 7 to 8 px on phone). Larger type relative to the canvas.
- SHOULD: the share post's text repeats the sheet's band; row holes; yellow "Gain not modelled yet" louder than item names; 63.5 on the sheet vs 63 on /char; /report loot stack long and "7 items:" ends on a colon; fallback seal bottom vs top on the other cards; scrim behind the lower card text.
- PASS: frame and corners only on the sheet's lead panel; teal, class, quality and kill-green stripes; seal v4 avatar and footers; raid post hierarchy; the band.

## wow-player: VERDICT: BLOCK
- BLOCKERS: bis-result fixture ("Pick head", eight "no pick" rows); report-pick fixture (Molten Core title over an Onyxia kill, every award item level 74, "7 items:"); webhook-report pitch and coach blurb, "1:14:00" with no unit, loot with no quality; BiS sheet quest rows without the quest name; char sheet "Upgrade available" on an empty slot, yellow "Gain not modelled yet", green legend vs the panel's wording, no "average" on item level.
- SHOULDS: raid post date said three times and "Sign up with the buttons below."; readiness "Short" says nothing; /char "Readiness: Ready", 63 vs 63.5; link previews with no crest or spec; /help lists itself.
- Data flagged (not this lane): Sentinel's vs Marshal's item levels, the Bronze Dragonflight signet in a Phase 1 band, the Mk XVII rifle. (Sapphiron Drape, Netherwind Crown and Bloodfang Hood do drop from Onyxia in 1.12.)

## ai-slop-reviewer: VERDICT: FAIL
- BLOCKERS: the fan-made line on every link preview; "Killed Onyxia" repeating the count; "item level" 17 times on the char sheet, the green legend, "Upgrade available" on an empty tile; /char panel "Epic" beside the purple dot, "By DPS gained, against the guide build.", "6 more upgrades on the character page."; webhook report pitch line and coach sentence; "1:14:00" vs "1h 14m".
- SHOULDS: raid post description and the long When stamp; ornate corners; BiS centre column; "Quest · requires level 58" wrapping; /help title and "/help This list."; lead lines that describe their own controls.

## info-density-reviewer: VERDICT: DENSE WITH ADDITIONS
- MUST: report card per-boss times and wipes; build card tree split with the class crest; char sheet upgrade summary (total gain over N slots); source on gain rows; one form for "Upgrade available" / "Gain not modelled yet"; BiS sheet gain per slot; raid post readiness; webhook report fields in one row; report post link button.
- LATER: BiS-vs-worn item level, per-raider item level on the roster, boss trends, guild leaderboard unfurl.

## guild-officer-reviewer: VERDICT: CLOSE
- BLOCKERS: raid post has no comp target ("Healers · 0 of 8") or buff cover line (guilds Cover); readiness "Short" does not say what; fixtures contradict each other (report title vs kills, 3 wipes vs 1, "tomorrow" vs "in 4 hours", raid-answer "None yet" vs the post's count).
- SHOULDS: loot winner's class mark and awarding officer; attendance pick last nights; tanks in top DPS parses; name who has not answered.
- Would flip: comp target and cover line on the raid post, readiness lines that name the fix, consistent fixtures.
