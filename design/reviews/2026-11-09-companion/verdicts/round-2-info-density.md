# Round 2: info-density-reviewer (captures in round-2/)

Round 2 closes most of the round-1 gaps.

## What landed

### Status
- The fights-uploaded figure in the headline.
- The Current report row reads state · start · visibility, with the log-written age under it.
- A held queue names the failing fight and its tries.
- "Fetch my builds" sits under the Addon label.
- Game folders are in view on the first screen.

### Device
- It is a real fact sheet. The empty ground beneath the card is acceptable for a small window.

### Reports
- The repeated credit name and the hint are dropped, which frees the rows.

## Remaining additions, ranked

1. **The Current report row has no zone.** `reports[key].zone` is in the snapshot but not rendered. Put the zone first.
2. **The Device uploaded row disagrees with Status.** Device says "1 report, 0 fights" while Status shows 3 fights uploaded. Make the two pages agree.
3. **Device has room for two cheap rows.**
   - "Addon last seen in game" from `addon_seen_at`. It may only show when the field is set.
   - "Last sync" from `last_sync.title`.

## Optional

- The held line is long. Cut it to "Fight 4 failed twice · retrying".
- "Characters sent" without names is correct as it is.

## Needs data

- The names in the last sync, as `last_sync.names`.
- Per-report kills and wipes.
- The last upload error.

VERDICT: DENSE WITH ADDITIONS — Status and Device now use most of what the snapshot carries; add the report zone to the Current report row and fix the Device fights count that disagrees with Status.
