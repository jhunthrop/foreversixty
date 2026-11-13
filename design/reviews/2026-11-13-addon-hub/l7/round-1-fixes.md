# L7 Officer tab: round 1 fixes

Round 1 verdicts: ux-designer FIX, wow-player FIX, ai-slop-reviewer FIX, guild-officer-reviewer FIX, info-density-reviewer PASS, brand-consistency-reviewer PASS. Round 1 files: `*-round-1.md` here.

Fixed (renders redrawn from `design/mocks/addon_hub_l7.py`):

1. Armed Give covered Won and Att. (ux, wow MUST). The session's columns now leave Give 104 px (Raider 88, gain 40, Wears 64 with a 4 px inset, Won 28, Att. 34); Won and Att. always show. A name too long for the column reads `Give?` and the footer names the raider and the item while Give is armed (`Click Give again: Ilvar gets Striker's Mark`).
2. Enchants cut mid-name (wow MUST, ux). Readiness rows draw as many whole names as fit, then a muted `+N`; every name stays in the row's hover.
3. Enchants figures disagreed in the nobody-addon render (slop MUST). The officer herself is Out in that render (her own addon is always live, so "Nobody coming runs the addon" only happens when she is not coming); readiness counts only those coming.
4. Raid buff and debuff coverage lost with the group panel (officer MUST). New readiness row `Not covered`: the needs of tonight's bosses no class coming can bring (ReadyCheck's coverage over the sign-ups, before anyone is grouped), the hover lists every need and who covers it. The last ready check is its own row (`Ready check`, the time in its hover, what each member lacks).
5. The in-raid member view (officer MUST 2): by the spec (4.6 and 6) the pre-pull look at other members is the officers' roster; a raider keeps `/fs ready` in chat and the chip's ready-check line, and Guild's Your prep reads their bags. No change.
6. Unknown Gap, Ench. and iLvl were blank beside `Best` and `0` (wow, officer): a muted `-` now, as Score without the pack.
7. `Online (24)` read as the guild online: `Raiders online (24)`.
8. Right column: `open` plain (not client gold); an item nobody answered yet reads `no answers yet`.
9. States added: `l7-loot-waiting` (before 30 s: gain a muted dash, `Waiting for answers.`, a Give master loot will not hand out drawn muted with the client's reason in the footer), `l7-sent` (`Sent to guild` in the footer for 2 s).
10. Won heading hover: `Items they won this week in your raids, from your own loot record.`
11. Raid build hover names the build each member follows; pushes heard from other officers are kept per class, so the row counts every class with a known push.
12. Long rosters page: `16 of 31 shown · Next 15`, then `17 to 31 of 31 shown · Back to the top`.
13. Without sign-ups by name the roster note says `sign-ups come with the companion`.
14. Copy cut: `No guild roster yet.`, `Send %s to the guild? It reaches them as a build to choose, never applied.`

Kept, with reasons:
- Role counts in the warning colour when short: the Guild tab's `GuildRaid.signups`, approved in L5 round 2; both tabs read one function.
- `Opens here when you loot as master looter.`, `Click Give twice: master loot hands it out`, `6 to the biggest gain`, `Orwen without`, `Dravik off`: the owner-approved board's and spec 4.6's words.
- Keyboard focus ring: the client's buttons and the hub's tabs draw it (L1); the tab adds no control of its own kind.
- A roster scroll: spec 3.1 says the window never scrolls; the list pages instead.
