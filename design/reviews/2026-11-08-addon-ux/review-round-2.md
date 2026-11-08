# Addon UX, persona review round 2

Round-2 boards and spec (commit 7bb65be77) with `review-round-1.md`'s dispositions. Same three `wow-player` personas, resumed with their round-1 context.

| Persona | Verdict |
|---|---|
| Leveller | **SHIP**, one copy must-fix |
| Raider | **SHIP**, one must-fix |
| Player who hates clutter | **SHIP**, shoulds only |

All three SHIP. The two must-fixes landed in the final commit, with the dispositions below.

## Leveller: SHIP

| # | Finding | Disposition |
|---|---|---|
| M1 | Pick your path lines end in an ellipsis ("Lacerating Stri…") | **Fixed.** Spec lines wrap to two lines; the spec forbids ellipses on panel rows (§4.3). |
| S1 | Confirm a Beast Mastery hunter learns Aimed Shot by 22 in Forever | **Accepted as a build-lane gate.** §4.3: check `SkillLineAbility` before ship; drop the row if it is an untaken Marksmanship talent. |
| S2 | Name the shared cooldown | **Accepted.** `When its cooldown, shared with Multi-Shot, is up`. |
| S3 | Pet row first in the follow-up | Noted for the follow-up spec. |
| S4 | `Preparation (flasks, food, enchants)` stretches the Shift tooltip | **Kept.** It is Shift-only, and "Consumables" would drop enchants, a third of what it measures. |
| S5 | Grey small gains in the panel | **Declined.** Every row is already above the 1.0 floor; greying would add a fifth tooltip colour meaning. |

## Raider: SHIP

| # | Finding | Disposition |
|---|---|---|
| M1 | No visible pressed-but-unsent state | **Fixed.** Middle board of `panel-raider-60.png`: `In, not sent yet`, buttons gone, `Sends when the companion app runs` on its own line; spec §4.3 updated. |
| S1 | "with Thornhide" orphaned on wrap | **Fixed.** `with Thornhide` is its own muted line under the job. |
| S2 | Striker's Mark from Magmadar; a Disenchant button | Striker's Mark is Magmadar's drop in the Molten Core loot table; the Classic-era group roll frame has Need, Greed and Pass only. Gain stays labelled a layout figure. |
| S3 | Phase 2 wording | Out of scope (API slot templates); passed to the coordinator. |
| S4 | `Sim again` reads as a status | **Fixed.** Drawn as a small grey button. |

## Player who hates clutter: SHIP

| # | Finding | Disposition |
|---|---|---|
| S1 | Gold and the arrow | Accepted by the reviewer; no change. |
| S2 | Drop the Shift sentence from the welcome | **Kept.** Shift is the only door to everything §2.3 moved off the default hover; the welcome is the one place it is taught. |
| S3 | Tonight row is heavy | Accepted by the reviewer as raid-day only. |
| S5 | The raid-day chip line for 18 hours | **Accepted.** Once the panel has been opened on raid day, the chip drops the raid back to normal priority for the session until one hour before the pull. Spec §3.2 carries it. |
| S6 | `Tooltips only` first in Settings, silencing chip, welcome and marks | Already so (§3.3). |
