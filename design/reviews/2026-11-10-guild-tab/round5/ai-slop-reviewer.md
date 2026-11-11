# Round 5 AI-slop review: Guild tab (iOS and Android)

**Verdict: PASS WITH FIXES**

Read the two full-page officer captures and the loot crop. "Assigning is on the site." is gone from both captures.

## 1. First impression
Made by someone who plays: real roster names, real items, real encounter slot names, real item levels; dense and uneven the way a player would build it.

## 2. Round-4 items
- **Item 1, "Assigning is on the site.": resolved.** The header carries a gold "Assign in planner ›"; it names where the action happens and looks tappable.
- **Item 2, Needs you lines and log-out stamp: accepted as a tracked follow-up.** "From your last log-out, 9 h ago" is still tell 22; the follow-up should reduce it to "9 h ago".
- **Gear-gap unit: good.** "41 HPS" for the healer, "37 DPS" and "22 DPS" for the others.

## 3. Remaining tells (none are blockers)
1. **Tell 22, "item level" repeated on roster rows** ("60 · Protection Warrior · item level 66"); use "ilvl 66". The longer label wraps on Thalrin on iOS.
2. **Tell 1/12, glow on the loot icons:** the purple quality borders in `parity/ios-crop-loot-yours.png` have a soft outer halo; a game icon border is a crisp 1 px line. Borderline tell 12.
3. **Tell 8, one dark rounded card on every block:** acceptable here; do not add more card chrome.
4. **Tell 11, the 31 / 4 / 24 strip:** defensible on a guild header.
5. **Parity copy drift:** Guild sims reads "You are #7 of 14 DPS" on Android and "#7 of 14 in guild sims" on iOS (use the iOS wording); Recently on times bold on Android, not iOS; iOS truncates the Healers line ("(Lat…") where Android wraps.

## 4. Keep
The plain gold links; assignment rows with real slot names and fill levels; class-coloured names; missing buffs as spell icons; Before you pull's tick, cross and warning with the Auction house; the invite block and the In / Late / Bench / Out control; the roster sections; no em dashes, emoji, gradients, button glow, uppercase eyebrows or invented social proof.

## 5. Fixes before this counts as PASS
1. Remove the glow halo on the loot icon borders.
2. Cut "item level" to "ilvl".
3. Align the Guild sims line across platforms using the iOS wording.
4. For the raid-lane follow-up, trim "From your last log-out, 9 h ago" to the time alone.
