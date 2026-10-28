# Task 33 - WoW-player read 2 (2026-10-27)
Verdict: SHIP WITH FIXES (copy only). All 31 string changes landed word for word in design/specs/2026-10-28-log-coach.md §8; mock fixes landed (footer "Reviewed 10 Dec" matches the report date; Death Wish praise; flask line corrected; "9 of a possible 40 went unused" matches the "9 casts lost" title).
Remaining corrections:
- rowError / error | `The coach couldn't read this fight just now.` | `...couldn't read this pull just now.` (pull is the vocabulary everywhere else; also DetailIdle `% of the fight` -> `% of the pull`, StrengthAdherence is fine)
- CannotLevelSim | `The simulated comparison: this log doesn't tell us your level.` | `The simulated rotation check: ...` (or `Our guide's simulated rotation: ...`, matching StrengthCastsSim)
- footer | `Evidence: {labels}.` | `Based on: {labels}.` (system-speak); historyRules `Rules` | `Log only` (a tag players will not parse)
- lastRaidPremium*, discordLine* use `What to fix first: {stem}` while the strip says `Fix first · {title}`; pick one (`Fix first: {stem}`), shorter and one voice.
Cleared: Coach in More after Rating, "no log in the last 14 days", "{n} of a possible {m} went unused", spell icon on the teaser. No exclamation marks, no "heal more", no "findings" in player text (board cell labels only).
