VERDICT: SHIP

Reviewed: `design/mocks/2026-10-30-mobile-apps/renders/f8-build-{ios,and}.png` (7 frames each) and `f8-you-{ios,and}.png`, against spec §5.8.1, after the wow-player HOLD of 2026-10-31.

## Player must-fixes, closed
1. Includes lines are /premium's own Premium-column words in /premium's order (Coach, server sims, history, Discord message, Supporter mark). The mock shows all five on the own and seat frames; the spec ties each to its entitlement flag and hides a line whose flag is false.
2. `Sim results and logs kept for a year` is gone from the spec (§5.8, §5.8.1, §13) and from every frame, including the older `f8-you` boards.
3. The history tick shows only when `entitlements.history` is true. The flag is the server's, so the app promises nothing the account cannot open. Whether the history surface is live in November is the owner's data question, not a layout one; if it is not, the API must send `history: false` and the line disappears with no app change.

## Player shoulds, taken
- Managed line reads `Billing is managed on foreversixty.gg.` as plain muted text, no link (frame F8.4 GB).
- Manage billing is now a quiet button: 1 pt outline, `nav` text, sentence case, 44 high. It is the quietest control on the card; the Founder tag and plan line outrank it (frame F8.4 US).
- The head `Premium` tag, row `Founder` tag and card `Premium` h2 stay. The head tag is the one signal that shows when the Plan row is scrolled away; the h2 names the card. Not changed.
- UK charge currency: not a layout question; the line renders what the API sends (`$40 a year` on the site's `formatCents`). Raised to the owner, not settled here.

## Region by region
- F8.1 free US: `Plan · Free · See plans` with the external glyph, one gold link, nothing else Premium. Pass.
- F8.2 free GB: `Plan · Free`, no chevron, no button, no other Premium word. Pass.
- F8.3 own founder: gold head tag, row sub without price, Founder tag and chevron, no overflow at 402 or 412. Pass.
- F8.4 US and GB: full site line with price, founder line, five includes lines that wrap without clipping, footer. Pass.
- F8.4 seat: no billing control, no founder, given-by line. Pass.
- F8.1 largest text: row stacks (Plan, Free, See plans own line), BattleTag wraps under the tile on iOS instead of cutting (fixed in this pass). Pass.

## Polish, not blocking
- The five includes lines are long at 402 pt (two to three lines each); the card runs about 520 pt tall and still fits above the tab bar on both devices.
- iOS AX5 frame clips Notifications under the tab bar by design: the screen scrolls.

## Build acceptance still owed
Real-device captures from spec §5.8.1 "Acceptance screenshots": both platforms, the seven frames, plus pressed and focus-visible on the See plans row and a 360 pt capture of F8.3 and F8.4 with the longest guild name.
