# Home body, round 1 (2026-11-09)

Captures: `r1-1440.png`, `r1-2000.png`, `r1-2576.png`, `r1-390.png`. Before: `before-*.png`.
State: signed in, level 34 Beast Mastery Hunter in Olympus XXVII with no raid planned, several guild sims, 14 upgrades (fixtures `*-34-guild.json`).

## ai-slop-reviewer: VERDICT: FAIL

Blocking:
1. Tell 8, card in card. Your upgrades sat in the ornate double-rule Frame, beside three more identical panels. The reviewer asked for a single border on upgrades and borderless Guild sims and Your spec this week.
2. Tell 19. "Best in slot changed for Beast Mastery Hunter 30 to 39: a new ranged pick" names no item.

Non-blocking: feed and Guild sims repeat the same sims; two "climbed 6 places to #3" rows look templated; "Climbing this week" repeats its heading; two "Not datamined yet" rows; the phone feed wraps to three lines; the phone "over" line truncates; a gap between the clock and the guild facts at 1440.

## ux-designer: VERDICT: SHIP WITH FIXES

Blocking:
1. Phone feed rows: the time column squeezes the sentence to three lines. Put the time under the sentence below 560 px.
2. Guild sims: the selected Felsnap row sits 16 px right of the others.
3. Guild sims splits 3+2 into two columns at 2000 and 2576, so Felsnap reads as #1. Keep one column.

Non-blocking: wide-screen empty space in the feed and the plaque; the raid list sits too close to the corner ornament; a tight gap between the "over" ellipsis and the gain; two links to BiS; the panel tops don't line up; raid rows have no affordance; on phone the plaque pushes upgrades below the fold; phone truncation.

## wow-player: VERDICT: SHIP WITH FIXES

Blocking:
1. "A level 34 hunter wearing mail": Highlander's Chain Girdle.
2. "3 of 4 raid nights attended" before any raid has opened.

Non-blocking: the vague "Best in slot changed" row; "Not datamined yet" reads as dev wording; an off-band Fury Warrior climb in the feed; with the guild name dropped, guildmates can't be told from strangers; a "+3.2" with no baseline; "+7.8 in 14 upgrades" is a sum; no level or drop chance on upgrades; phone truncation and wrapping.

## info-density-reviewer: VERDICT: DENSE WITH ADDITIONS

No blockers. Additions, ranked: Onyxia's drops as icons; more guild facts in the plaque (tier progress, waiting, needs before next raid); feed rows as drill-downs; Guild sims sim age and spec; links in Your spec this week; a rail filler; time under the sentence on phone; check the no-guild plaque.

## Changes for round 2

- **Upgrades:**
  - Moved onto the panel under its heading (`Section boxed`), like the feed beside it, so no framed card any more.
  - Guild sims and Your spec this week stay on panels. This is the owner's binding rule: class- and quality-coloured lists sit on a plate.
  - The phone wraps the name and the "over" line instead of truncating them.
- **Feed fixture:**
  - The fixture's BiS change row was not in the API's words. `api/internal/activity/bisdiff.go` names the item when one pick changes ("Golemsight Long Gun is the new ranged"); the fixture now does too.
- **Shared FeedRow (every feed):** on a phone the time drops under the sentence.
- **Shared Row:** an open selected row keeps its content on the column; the gold band reaches into the panel's padding. This fixes the Felsnap offset.
- **Guild sims:** one column.
- **Hunter "wearing mail":** not a data error. Highlander's Chain Girdle is leather in this client build (`items/hunter.json`: `armor_type: leather`, `subclass_id: 2`), so the worn piece is legal at 34.
- **Attendance:** the 34 guild fixture no longer carries raid nights. The code only shows attendance the account carries.
- **Raids not in the client:** "Not datamined yet" became "Loot not out yet". The raid list got 16 px of right padding off the ornament.
