# Round 1: the first gate on the boards

Boards rendered by `design/mocks/addon_tour.py` from Locale.lua and Data.lua, 2x, into this folder. Six reviewers read them with the spec's section 6.

| Reviewer | Verdict | Blocking items |
|---|---|---|
| ux-designer | SHIP WITH FIXES | the druid board's paladin crest (mock), the panel board's rotation block (mock), the guild board's overflow (mock), the footer line's alignment (Lua offset 4 to 0) |
| wow-player | SHIP WITH FIXES | "Plate healer" at level 13 (plate at 40), "Twist of Light" and "Berserk" absent from the shipped talents, the paladin crest on the druid row (mock), raid builds totalling 50 points at 60 (data, not this lane) |
| ai-slop-reviewer | PASS WITH FIXES | the leads on pages 1, 2 and 3 (shorten, name a thing), the repeated subtitle after page 1, "loads its leveling build" and the chip card's footer, the chip card's solid gold bars |
| info-density-reviewer | THIN WITH FIXES | the split beside the question, the next talent on own rows, tiles and Pick rows, the chip card's footer |
| brand-consistency-reviewer | DRIFT | the wrong crest (mock), "Custom build" over site-sent rows, the "only gold" claim (the title and tab are gold), the chip card's gold bars; the dialog's older chrome noted for a later pass |
| guild-officer-reviewer | CLOSE | an empty spec cell beats a guess; wants a hover with the split on others' rows (needs data the hello does not carry) and a "Pick your path" target on the own row (follow-up); the caption and a level 60 row on the board |

The verdicts are in `<reviewer>-round-1.md`. Round 2 answers the blocking items; `round-2-notes.md` lists what changed and what was declined with its reason.
