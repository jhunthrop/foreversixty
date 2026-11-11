# Round 4 findings: what we change in round 5

Round 4: guild-officer WORTH PAYING FOR (passes); info-density DENSE WITH ADDITIONS; brand DRIFT;
ux-designer NO SHIP; ai-slop FAIL. Codes: ID, GO, BR, UX, AS as before (round-4 sections).

## Both apps
| # | Change | From |
|---|---|---|
| 1 | The officers' not-synced line: primary text colour (not warning red), a trailing chevron, a 44 pt / 48 dp row, the same on both | UX1 |
| 2 | The filtered roster says it is filtered: the not-synced tap opens the roster with a selected chip `Not synced this week · 4` and a clear control (`×`, 44 / 48), `All` deselected, the stat reading `4 of 38`; the change is the filter state only (the roster screen itself is another lane's) | UX2, GO P4' |
| 3 | The officer push results show on iOS as on Android: after Mention, Accept or Decline the notification is replaced by the result (`Mentioned 3 healers in Discord.`, `Application accepted`, the failure line) | UX evidence, GO P2' |
| 4 | Covered: the `and N more` tail is a tap target to the raid page's cover list; the one icon sits on the text baseline | ID N2, BR N3 residue |
| 5 | Buttons from the shared tokens: primary `teal` fill, `tealEdge` border, `strong` text, no gradient; secondary `blue` with `blueEdge`; Android's `FlowButton` gains the secondary kind and loses the gold gradient (both apps then draw Preview mention and Accept alike) | BR R4-1, R4-2 |
| 6 | ~~Activity's trailing figures carry one decimal~~ REVERTED after the main merge: the house format (apps/shared/formats.json, owner 2026-11-10) prints figures grouped and whole from 100 (`204 DPS`) | AS polish, ID N3 |
| 7 | The raid head's size reads `40-man` (`Invites 6:45 pm CST · 40-man`) | AS polish |
| 8 | Copy: delete the unused one of `raid.noRaid` / `raid.tab.noRaid`; fold `FlowRows.kt:348`'s line-height override into a role | AS polish, BR N4 residue |

## iOS
| # | Change | From |
|---|---|---|
| i1 | The Activity tab badge loads at shell start (as the Raid badge does), so it shows on every tab | BR R4-3 |
| i2 | The Raid badge counts by the same rule as Android (the officer state read 1 on iOS and 6 on Android) | BR R4-4 |

## Android
| # | Change | From |
|---|---|---|
| a1 | The drawer's Notifications row shows its real value from the per-group state (`Raids, gear and sims, logs on`), built by the same function the golden uses | ID N1 |
| a2 | The drawer's disabled `Refreshing…` draws no ripple patch | UX polish |

## Server and fixtures
| # | Change | From |
|---|---|---|
| s1 | The ready check's enchant row is `Boots not enchanted` with nothing under it: no restated detail, no where-to-get line (server copy and the fixtures) | AS must-fix 1 |
| s2 | Fresh sync stamps in the captures: the fixture character's `synced_at` / `captured_at` / `logged_at` move to the day before the capture date, so the app captures read hours, not `12 d`; one stale capture is kept from round 4 | AS must-fix 2 |

## Captures (round 5)
Android: close the notification shade before every section (the r4 gear and activity frames were the shade); Gear with the stale tone (kept from round 4) and fresh. iOS: the officer push results after the action.

## KEEP
- The age on its own meta line under feed rows (AS polish): round 3 decision 3, asked for by ux-designer and brand to stop `· 3 d ago` wrapping; the web keeps it inline only where the row is wide.
- Covered icons for Might and Wisdom: the fixture build has no art; live builds do.
