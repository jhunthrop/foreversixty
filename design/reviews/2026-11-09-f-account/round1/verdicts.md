# Lane f-account, review round 1 (2026-11-09)

Captures: the build at fd66c30c2, design/reviews/2026-11-09-f-account/build-*.png (overwritten by round 2).

## ux-designer
Verdict: SHIP WITH FIXES
- MUST 1 /connect: the code is a 12 px table row; make it the display-size answer above Device and Platform. Fixed: check line, then the code at 24 px, Device and Platform under it.
- MUST 2 /connect success: the Connected pill touches the code. Fixed: the pill sits 12 px after the big code.
- MUST 3 /account free and new: the Discord rows squeezed in the 340 px rail. Fixed: Discord moved to the main column (directions mock order).
- MUST 4 /setup done: empty screenshot frames. The images are lazy and the full-page capture never scrolled them in; the capture now scrolls the page first.
- MUST 5 /account: column imbalance. Fixed: main = Characters, Guild, Devices, Discord; rail = You, Plan, Delete account (the mock's order). The loading reserve follows (layout.ts, global.css re-measured).
- SHOULD 6 Plan rows 840 px apart: Plan is in the rail now.
- SHOULD 8 new account empty paragraph all gold: body colour now.
- SHOULD 9 /setup duplicates: the repeated Forever Sixty Data sentence is gone; the head, step 3 and The companion downloads are owner rulings (2026-11-08: "a download link at the top", and the download right under The companion heading).
- SHOULD 10 /login dead padding: trimmed.
- SHOULD 7 Switch 8 px past the text edge: the shared ui/Switch's -mr-2 hit area; not page code.

## wow-player
Verdict: NOT SHIPPABLE until the MUST items are fixed.
- MUST 1 empty setup screenshots: capture fix (above).
- MUST 2 /app slogan: "Your raid, your loot and your next upgrade, in your pocket." cut. The feature list stays: it says what the app does.
- MUST 3 /setup repeats: the duplicate data sentence cut; the three downloads are owner rulings.
- MUST 4 handshake: kept. It is the what-leaves-your-machine statement the setup spec pins (and the client's own "Site has your save from" line).
- MUST 5 two Unlinks in Discord: the webhook's button is now "Remove channel" (spec 2026-10-25 amended).

## ai-slop-reviewer
Verdict: FAIL
- Tell 8 cards in cards: the framed inset line is the design system's framed panel (docs/design-system.md, two or three a page), not a nested card.
- Tell 13 empty placeholders: capture fix.
- Tell 7/21/22/17: the app slogan cut; the Discord remove renamed; the rest are product lines kept on purpose (billing, privacy, delete).

## info-density-reviewer
Verdict: DENSE WITH ADDITIONS
- Added: the character row's item level (mock b4-account). Left column dead space: fixed by the mock's column order.
- Not taken this lane (needs new UI beyond the page restyle): a Recent sims/reports panel, guild roster counts, device kind icons, Discord failure chip.
