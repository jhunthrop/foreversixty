# Round 5 final verdict: Guild tab, AI-slop review

**Verdict: PASS WITH FIXES**

None of the open items are visual tells 1, 2, 5, 8, 12, 13 or 15 on the Guild tab's own surface. The one routed item that is borderline tell 12 belongs to the design-system owner, not to this tab.

Re-checked the re-rendered Android capture and the copy: rows read "ilvl 66" and "ilvl 58"; `guild.tab.itemLevel` is "ilvl {n}"; Guild sims reads "#7 of 14 in guild sims". The iOS capture was not re-opened.

## Guild tab's own items

| Item | Status |
|---|---|
| Round 4, "Assigning is on the site." | Closed. "Assign in planner ›" is a plain gold link. |
| Fix 2, "item level" on roster and gear-gap rows | Closed. "ilvl N", the web roster's word. |
| Fix 3, "You are #7 of 14 DPS" | Closed. Android fixture text only; the shipped `raid.pages.simsRank` reads "#{position} of {of} in guild sims". |
| Gear-gap units ("41 HPS", "37 DPS", "22 DPS") | Good. |

## Routed items (not Guild-tab blockers)

- **Fix 4, "From your last log-out, 9 h ago" and the Needs you lines:** the raid-lane follow-up named in `round4-response.md`; the stamp should read "9 h ago" or go.
- **Fix 1, the 10 pt quality glow at 22% on item icons:** routing accepted. `ItemIcon.swift` is the app-wide icon from the one W3C token source the site and both apps share; changing it on one tab would split the brand. Judgement unchanged: replace the glow with a crisp 1 px quality border everywhere at once, via the design-system owner. Until then a known borderline tell 12, not a Guild-tab defect.

## Notes (no action required)
- Tell 8: the dark rounded card repeats across about 14 blocks; dividers and row variety keep it from being stamped. Do not add more card chrome.
- Tell 11: the 31 / 4 / 24 strip is defensible on a guild header.
- iOS truncates the Healers line ("Lat…") where Android wraps; not re-checked on iOS.

## Keep
Real slot names and fill states in Assignments; class-coloured names; missing buffs as spell icons; Before you pull with a real attunement, a real consumable and the Auction house; no em dashes, emoji, gradients, button glows or uppercase eyebrows; "Ping on Discord" as a plain solid button.
