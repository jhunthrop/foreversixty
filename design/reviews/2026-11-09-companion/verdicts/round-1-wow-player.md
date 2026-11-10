# Round 1: wow-player (captures in round-1/)

## 1. The question

The question this window answers: is my raid being uploaded right now, and if not, what do I do? Status answers it first: a coloured dot, "Logging", and the report link. That is better than the WarcraftLogs Uploader. The first run (`first-run-device`) is right: one button and no pitch.

## 2. What a raider would roll their eyes at, worst first

1. **`ready`, `reports-empty`.** The /combatlog toggle explanation is printed twice. Cut it to "Waiting for a combat log. /combatlog in game."
2. **`offline-status`.**
   - "Right-click its icon in the system tray (bottom right, under the ^ arrow)..." explains the obvious, and the wording is Windows-only.
   - A dead companion mid-raid offers only a manual three-step restart.
   - "may not be uploading" is vague.
3. **`update-ready`.**
   - "Pick a quiet moment" is hand-holding.
   - A manual quit and restart is a chore. Add a Restart now.
4. **`uploading-held`.**
   - Nothing says a fight is stuck: no failing fight, no retry and no reason.
   - The amber dot does double duty for "not credited" and "stuck".
5. **`logging-not-credited`, `reports`.** "Log out once in game with the addon loaded" appears up to three times on one screen. Show it once.
6. **`logging-credited`.** "logging as Bow Jackzon" and "credited to Bow Jackzon" say the same thing. Drop one.
7. **`sign-in-waiting`.** The countdown "Expires in 4:58" sits beside copy that says links expire "after 10 minutes", and the bar fill matches neither.
8. **Fixtures.**
   - `logging-credited` shows Fights uploaded 14 while `reports` shows the open Molten Core report with 3 fights.
   - The fixture dates run 14 to 18 October under Molten Core and Naxxramas, which makes the Reports capture look fake.
9. **Ellipsis styles are mixed.** Both "…" and "..." are used. Pick one.
10. **Stray explaining.**
    - The "Automatic reads who is logging..." sentence and the three-clause CurseForge/Wago note.
    - The password line can stay.
11. **`icons`.** The tray tones differ by colour only.
12. The update banner and the Save strip cover the lower folder rows at 800 x 620. Check that the rows scroll.

No em dashes and no "X is the Y:" blurbs found.

## 3. Missing

- A live last-fight line: boss, kill or wipe, uploaded at.
- Per-report kills and wipes.
- An Open action per row.
- A pause control.
- A restart action.
- "Check for builds now": a raider reads "build" as a client build.
- "Last synced" against "Characters sent" reads as two sync times.
- The current visibility on Status.

## 4. Data doubts

- Is the live Forever folder still `_classic_beta_` after the 4 November launch?
- "Sim needs premium" reads as an upsell. Confirm it shows only after a real request.

VERDICT: SHIP WITH FIXES — the layout and first-run flow are good, but a dead companion offers only a manual restart, a held upload looks like a normal one, the sign-in expiry contradicts itself (5 versus 10 minutes), and the copy repeats and over-explains in places.
