# h-report-tail, review round 1 (captures at 379baf8b0)

## ux-designer
Blockers: Queries at 390 drops hits and total; Timeline at 390 clipped at 0:10 with no scroll cue; Ratings at 390 Weights triple has no role labels; Ratings at 390 Forever Score bars dropped, figure repeated, lead corners overlap the title.
Shoulds: same placeholder icon on enemy spells; Events Time first and drop an all-dash Amount; Events footer missing space; Queries spell_name right-aligned and numbers with no commas; orphan "What these mean"; ratings prose long, flat sparkline; night "#1 #2" unlabeled; Events chips wrap alone at 390.

VERDICT: HOLD - Queries, Timeline and Ratings drop their data or labels at 390 px; fix those four blockers and recapture.

## wow-player
Blockers: Events Amount column empty and footer missing space; Events reads unlike the client log ("on X by X" for self buffs, "off" for fades, one apply row per target); same placeholder icon on every rail spell; Timeline casts anonymous, Casts and Taunts share yellow, clipped at 390; Ratings sample scores Parses and Attendance before raids open; Performance Rating sparkline flat against "Up 25".
Shoulds: Queries "with the load" meta talk, numbers with no commas, spell_name right-aligned; night "#1 #2" and "104 fights" against 6 boss pulls; Interrupts: say none were interruptible, "No interrupts"; Dispels rail "applications"; Ratings copy: simulator/threat status talk, "no leaderboard" is defensive.

VERDICT: HOLD - Events show no damage amounts, every spell in the Dispels and Interrupts rails shares one placeholder icon, the timeline is unlabeled and cut off at 390, and the ratings sample scores raid parts that can't exist yet.

## ai-slop-reviewer
Tells: inner frames and corner flourishes on some panels only (tell 8); empty-state cards over blank pages (tell 13); grey explainer paragraph under every ratings row (tell 7); qualifiers ("in this window", "with the load", "#1 #2"); flat sparkline (tell 16); tier colour note explains the obvious; Events footer spacing and all-dash Amount.

VERDICT: FAIL - panels stamped everywhere with inner frames (tell 8), empty-state cards over blank pages (tell 13), and explainer paragraphs under every ratings row (tell 7).

## info-density-reviewer
Adds: Events amounts or no dead column; Night share bars and Copy CSV; Queries Copy CSV and commas; Interrupts says which casts were interruptible; Dispels rail uptime; Night rail fastest kill and trash pulls; Ratings points to next tier and less prose; Timeline cast names.

VERDICT: THIN - Events' Amount column is empty, Night, Timeline, Dispels and Interrupts each leave half a screen of empty ground, tables lack the sort, share bars and Copy CSV that sibling pages already have, and Ratings is mostly explanation prose.
