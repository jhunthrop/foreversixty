# Round 6 findings: what we change in round 7

Round 6: guild-officer WORTH PAYING FOR; brand CONSISTENT; ux-designer SHIP; info-density DENSE WITH
ADDITIONS; ai-slop FAIL (one must-fix). Codes: ID, GO, BR, UX, AS (round-6 sections).

## Both apps
| # | Change | From |
|---|---|---|
| 1 | Gear's header figure in the house format (apps/shared/formats.json: grouped and whole from 100, one decimal below): `182 DPS`, `+18.0 DPS from 13 upgrades`. One formatter for figures per app, used by Gear and Activity | AS must-fix 1, BR S1 |
| 2 | Officer push failures say what failed and how to retry. Mention: title `Could not mention the healers in Discord` (`push.short.failed`, `{who}` the short role's plural as the action names it), body `Onyxia's Lair. Open the raid to try again.` (`push.short.retry`). Accept/Decline: title `Could not accept Mirelle` / `Could not decline Mirelle` (`push.applicant.acceptFailed` / `.declineFailed`), body `Mirelle, Holy Priest 60. Open Forever Sixty to try again.` (`push.applicant.failed`, `{applicant}` the applicant line). The tap opens the raid page / the applications list, as the original push does | GO S1, UX S1-S3 |
| 3 | Gear's item sheet shows where the selected piece comes from (one `Where` line from the loot data the app already loads: boss and instance, drop chance, or profession), as Home's Best run line does | ID S1 (round 5 S3) |
| 4 | Android's drawer Notifications value reads the account's per-group settings (GET /v1/me/notifications, as iOS does) AND the channel state: a group is on only when both are. Same fixture, same line on both apps | ID S2 |

## Captures (round 7)
- Gear (fresh) and the Gear sheet with the Where line, both apps.
- The two officer push failures, both apps.
- The drawer Notifications value in three states, both apps: some off (the fixture), All on, All off. UX S4.

## KEEP
- Gear cells carry no source sub-line: at 390 a cell holds the name, the slot and the gain; the source is one tap away in the sheet (row 3).
- `Updated … from Battle.net` (drawer) and `Synced …` (Home, Gear) are different facts from different fields (`bnet_imported_at`, the character's `synced_at`) and keep their own verbs (AS polish).
- The duplicate day on the Molten Core feed row (round 6 KEEP).
- With `Show older` and one fixture page, the empty ground under three Activity rows is the honest end of the list (ID S3).
