# Round 1: info-density-reviewer (captures in round-1/)

The Status and Reports pages carry real data but could carry more. Device is the thin one: about half the page is empty.

## What the window can reach but doesn't show

- **`pipeline.last_append`** (drain.go:298) is how long ago the combat log last grew. Status has no "log last written 12s ago" row. That row is how a raider checks that the log is being read mid-raid.
- **Report age in the Current report row.** `reports[0].started_at` is in the snapshot, but the row only reads "In progress · credited to …".
- **`pipeline.queued_fights`** is the fight count behind the Queued row, which today shows only the count of queued uploads.
- **`last_sync.count` and `last_sync.at`** feed "2 characters, 4 minutes ago". The names behind that count are not shown. `characters[]` (`addon.Export`) carries `name`, `ruleset` and `region`.
- **`addon_seen_at`** is already in the snapshot. It is the only signal that the addon was seen in game. It is used only as a tooltip (`addon_seen_title`), never as a visible figure.
- **`builds_checked.new_builds` and `builds_checked.for`** feed the checked line. Keep that line. `builds_waiting` is already wired in.
- **`installs[].last_played`** shows only as a time, "Last played 18:40". Add the date or a relative age when it is not today.
- **`sim_request_state`** has no visible surface. Add it only if the site feature is live.

## Panels that say too little

- **Device (connected)** is one card with an empty lower half. Fill it with:
  - companion and engine version (`app_version`, `engine_version`); the footer repeats the version, so show the engine version here and drop the footer's duplicate;
  - visibility (`visibility`);
  - credited as (`credit`, `logging_character`);
  - last sync (`last_sync.title`);
  - addon last seen (`addon_seen_at`);
  - reports this device uploaded (`reports.length`, the sum of `reports[].fights`).
- **Status, Current report row:** make it a row of facts, not a sentence:
  - the zone (`reports[key].zone`, looked up in app.js:295 but not rendered);
  - the fight count;
  - the report's age (`started_at`);
  - the credit as a state dot (`credit.state`).

  The whole row should open the site report (`report_id`).
- **Reports rows:** every row repeats "credited to Bow Jackzon".
  - Show the credit only when it differs from `credit.name`, or when it is missing (amber).
  - Use the freed width for the date.
  - Keep "Copy link", and consider an "Open" icon beside it, since the zone title already links.

## What to drop

- Two grey sub-lines that explain the obvious. Keep the hint only as a tooltip or an amber-row action.
  - "A private store in your user profile." (`token_backend_note`).
  - The Reports hint "Not credited yet? Log out once…".
- The "uploads when the site is reachable" tail on the unsynced row. A ring and "not on the site yet" is enough.
- The second Details subpanel on Device (device name and API URL). It is low value for a raider: move it into one small mono line under the heading.
- The height of the "Check for builds now" row on Status. Put it beside the Addon line as an icon button. That frees 60 px, so "Game folders" is not pushed under the fold at 800 x 620.

## Phone, tray and window size

The collapse rules hold at 560 x 480. The three-column Device facts can fold to one column.

## Needs data (not in the snapshot)

- Upload byte or speed figures.
- The last upload error. A failing fight holds the queue, but no snapshot field names the failure for Status to show.
- Per-report boss kills or wipes.
- A site-side `last_seen` for the device.

VERDICT: DENSE WITH ADDITIONS — Status and Reports are mostly sound, but Device leaves half its card empty, and `last_append`, `addon_seen_at`, report zone and age, and the names behind the characters-sent row are already in the snapshot but not rendered.
