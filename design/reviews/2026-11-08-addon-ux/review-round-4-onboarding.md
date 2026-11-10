# Onboarding, persona review round 2 (overall round 4)

Onboarding v2 (commit 92003b0c3) with `review-round-3-onboarding.md`'s dispositions. The owner's brief fixes the loop pages, first-login auto-open, the QR and the `coming` marks; the personas judged inside those constraints.

| Persona | Verdict |
|---|---|
| Clutter-hater | **SHIP** (withdrew HOLD), two small must-fixes |
| Leveller | **SHIP** |
| Raider | **SHIP** |

## Clutter-hater: SHIP

| # | Finding | Disposition |
|---|---|---|
| M1 | Page 1 still shows the guild node and "and to your guild" without a guild | **Fixed.** `onboarding-1-loop-solo.png`: no guild node, no clause, four dots (§3.1). |
| M2 | Three exits on the last page | **Fixed.** No `Skip tour` on the last page; `Done` and `Open my panel` remain. |
| S1 | Shorten the last-page footnote | **Kept.** The owner asked for `Tooltips only` on the done step; page 1's button and this line serve the player who did not skip. |
| S2, S3 | The QR is the largest block and points at a page that is not live | **Raised to the owner** in the report; the brief asks for it, tagged `coming`. The web lane is building `foreversixty.gg/app` (coming soon and a tester sign-up), so the code never leads nowhere. |

## Leveller: SHIP

| # | Finding | Disposition |
|---|---|---|
| S1 | Check the path-tile copy (Barrage, Sniper Shot, Bestial Wrath, Lacerating Strikes) against Forever's talents; check the logging menu path | The tile lines are the shipped `Locale.lua` `specPlaystyle` strings, written from the client's own talent data (its comment says so). The build lane re-checks them against `talents/hunter.json` before the tour ships (§3.1 truth rule). The Advanced Combat Logging path is the companion README's. |
| S2 | Guild node without a guild | Fixed (clutter-hater M1). |
| S3 | Page 3 mentions attendance to a solo leveller | **Kept.** The five parts are what the score is; one line names them. |
| S4 | Skip next to Done | Fixed (clutter-hater M2). |
| S5 | Give a leveller a reason for the companion | Noted for the copy pass; the row already says reports and the coach. |

## Raider: SHIP

| # | Finding | Disposition |
|---|---|---|
| S1 | At 60 the tiles still read as leveling playstyles (BM raids without the pet in Forever's guide) | **Accepted for the build lane:** at 60 the tiles read a raid playstyle line, a new Locale table beside `specPlaystyle`, taken from the class guides' raid sections and verified (tenet 8). |
| S2 | Hiding your own score | Open owner question (§3.1). |
| S3 | `/app` must exist before the tour ships | Release gate for the build lane. |
| S4 | Page 4 is short for guilds that never post raids | Accepted; it only shows in a guild. |
