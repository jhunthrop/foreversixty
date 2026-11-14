# Lane L8 (Panel and tour fold), round 1

Renders: `l8-first-open.png`, `l8-steps-done.png` (design/mocks/addon_hub_l8.py, from Onboarding.steps and MeTab.model).

| Reviewer | Verdict |
|---|---|
| ux-designer | SHIP |
| wow-player | SHIP |
| ai-slop-reviewer | FIX (1 step 1's spec list, 2 the "3 steps" note, 7 the dead tour keys) |
| info-density-reviewer | FIX (1 the header's "no spec chosen", 2 the Preparation reason, 3 the empty left column) |
| brand-consistency-reviewer | FIX (1 the header's "no spec chosen", 2 the Preparation reason, 3 gold links on done rows, 4 the footer, 5 the Gear dot) |

## Fixed

- Header: Me's figures say `no spec chosen` while no build is followed (MeTab.figures, `L.meNoSpec`), as the board.
- The `3 steps` note went: the numbered rows say how many; the note shows only progress, `1 of 3 done`.
- Every `tour*` and `panel*` Locale key the deleted tour window and Panel used is gone (93 keys).
- The tooltip's hit line reads as the client prints it, `3.0%` (`%1$.1f%%`), wow-player's check.
- Step 1's second line runs under the link, so the three spec names fit whole.

## Kept, with reasons

- Step 1 keeps the specs on its second line: spec 4.1 says "`Pick your spec` with the specs"; info-density called it richer than the board.
- The Preparation reason (`2 flasks missed`) comes from the inbox's `last_raid`. On a first open with no companion there is no last raid, so there is no reason to show; the board drew both, which cannot be true at once. With a companion, Last raid and the reason both show (`l8-steps-done.png`).
- The empty left column under the score is spec 4.1: Last raid is left out until the companion brings one, so the app address shows once.
- Done rows keep their gold tab links: a link is gold (spec 3.5) and still goes somewhere; the board keeps them gold; ux-designer approved.
- The footer says `/fs tour` while Set up shows (the board); the companion's freshness returns with Right now.
- The Gear dot is the real upgrade data's (Gloves of Holy Might beats the worn Trueaim Gauntlets).
- Done rows keep one row height, so the tooltip never moves when step 3's address goes.
- Step 2 keeps the approved words `Hover an item you wear` (spec 11 board copy); its tooltip carries the point.
