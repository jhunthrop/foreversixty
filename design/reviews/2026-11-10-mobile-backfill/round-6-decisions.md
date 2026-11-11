# Round 5 findings: what we change in round 6

Round 5: guild-officer WORTH PAYING FOR; brand CONSISTENT; ux-designer SHIP; info-density DENSE WITH
ADDITIONS; ai-slop FAIL (one must-fix). Codes: ID, GO, BR, UX, AS (round-5 sections).

## Both apps
| # | Change | From |
|---|---|---|
| 1 | The drawer's Notifications value says what a player misses: `All on` when every shown group is on, `All off` when none is, otherwise the groups that are off (`Logs, tiers off`). Shared copy keys `menu.notifications.allOn`, `menu.notifications.someOff` (`{groups} off`); `menu.notifications.on` goes | ID S2, AS polish, UX nit |
| 2 | The filtered roster's stat plate counts one population: while a filter is on it shows `4 of 38 on the roster` alone, without the guild-wide `in game, no account` and `waiting` figures (the filter state only; the roster screen is another lane's) | GO S1, UX S4 |
| 3 | Push result failure frames: the Mention and Accept/Decline failure lines are captured once on each platform (a fixture route that fails) | GO S2 |

## Captures (round 6)
- Re-shoot the Android drawer on a fresh debug build: round 5's `Updated 12 d ago` was an APK built before the fixture's `bnet_imported_at` moved to 10 Nov (the drawer reads `bnet_imported_at`, Home reads the character's `synced_at`; both are fresh in the fixtures). AS must-fix 1, UX S3.
- Re-shoot Activity on both apps in the house format (`204 DPS`, `+12 over your 191 (Fury Warrior, 60).`). ID N4.
- One stale Gear frame on each app (the clock moved past the stale age). UX S1.
- The iOS flow snapshots from CI record run 42953910502. UX S2.

## KEEP
- The age on its own meta line (round 3 decision 3); the one duplicate day on a raid row stays, as the row's day and its age are different facts at a glance.
- Loot rows carry no trailing DPS figure: `loot.awarded` has no per-character gain on the server today (needs data; the parity golden's `+12.6 DPS` is a mocked body).
- The iOS Gear `Main hand · Equipped` wrap at the phone width (UX nit): the cell keeps its type role.
- The emulator's `Serial console enabled` line in the shade is the emulator's, not the app's (GO nit).
