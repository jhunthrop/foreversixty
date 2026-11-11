# Round 8 findings: what we change in round 9

Round 8: guild-officer WORTH PAYING FOR; brand CONSISTENT; ux-designer SHIP; info-density DENSE WITH
ADDITIONS; ai-slop FAIL (one must-fix). Codes: ID, GO, BR, UX, AS (round-8 sections).

## Both apps (the Gear item sheet's runner-up row)
| # | Change | From |
|---|---|---|
| 1 | A runner-up with no DPS difference reads `Same DPS` (`items.sameDps` is now `Same {unit}`), a figure-like tail like its neighbours (`−5.0 DPS`) | AS must-fix 1 |
| 2 | The runner-up row names its source only when it differs from the pick's source row above it (`Runner-up · −5.0 DPS` under `Crafted: Blacksmithing (275)`; `Runner-up · World drop · −0.3 DPS` under a vendor pick); when it is named, both apps use the same shared source words as the source row (`bis.source.*`, the boss shape `Boss · 8%`) | ID S1, UX S1, AS polish |
| 3 | A PvP source on the sheet drops the faction word: the rank title already belongs to one faction (`PvP rank 11 · Blood Guard`) | AS must-fix 1 |

## Captures (round 9)
The three item sheets on both apps (crafted pick with a PvP runner-up, crafted upgrade, vendor or world-drop upgrade), and the iOS item-sheet golden from the round-8 CI run for the boss-drop row (ID S2).

## KEEP
- `Replaces … / Equipped` on the sheet keeps its icon: it is the item the pick replaces, drawn as an item (AS polish).
- `gear.sheet.where` stays as the iOS parity region's section id (no visible text).
- Failure bodies keep `Open Forever Sixty to try again.` for applicants (GO nit): the tap opens the applications list, and the app's name is what a lock-screen reader needs.

## After the round-9 reviews (all pass)
Round 9: info-density DENSE, ai-slop PASS, ux-designer SHIP, brand CONSISTENT; guild-officer WORTH
PAYING FOR stands from round 8 (round 9 touched no officer surface).
- KEEP, native-only: the Gear sheet drops the PvP faction word (`PvP rank 11 · Blood Guard`); the web's
  `pvpSourceLabel` keeps it (BR nit 1). Carrying it to the web is the web lane's call.
- Vis'kag the Bloodletter drops from Onyxia like the pick, so its runner-up names no place (UX S2).
- Each app also captures the other's third sheet (Android the world-drop cape, iOS the vendor
  leggings), so both naming rules show on both apps (UX S1).
- KEEP: the Gear header's `Horde` names the band's faction list (BiS lists are per faction), not a
  qualifier (AS polish 3); the iOS tooltip golden is re-recorded by the round-9 CI run (AS polish 1).
