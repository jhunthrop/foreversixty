# Info-density review, lane BA4 (companion copy/token unification)

Scope: screenshots in this folder (main window incl. -1100 and -560, install-*, tray-menu, dmg-window) against companion/ui/app.js, copy.js, testdata/*.json. BA4 intended no layout changes, so this judges regressions first, then unused data. Feature-sized asks are NICE.

## Verdict: PASS (DENSE WITH ADDITIONS)

No regression found. Status, Reports, Device and Settings keep every figure the pre-lane testdata carried. Units are bare ("3 fights uploaded", "2 characters, 4 min ago", "10 reports, 75 fights"). The Reports row carries zone, time, fight count and credit state, and shows the credited name only when it differs from the usual one (app.js reportRow). The 1100 capture centres the content column, and 560 collapses the headline sub-line to two lines with nothing dropped. The failed and upgrade install states keep their action visible.

## MUST
None.

## SHOULD (copy that explains the obvious; fits a BA4 follow-up, no layout change)
1. Settings carries four grey explainer lines: "Anyone can find it in rankings and search.", "Pick a character only if Automatic is wrong.", and the two-line "New versions go in when the game is closed. CurseForge and Wago users can leave this on; they get the same files." Veterans do not need them. Cut them or move each to a title tooltip. Source: copy.js settings strings.
2. Status "0 fights" state adds "The first fight uploads when it ends." under "0 so far", and the headline sub repeats "Uploading your fights as they end". Keep one. The "Addon 2.0.16 ready · installs when the game closes" and "Data 2026.10.21 ready · installs when the game closes" lines repeat the same tail twice, so fold them into "Addon 2.0.16 · Data 2026.10.21 · installs on game close". Source: status-addon-build-waiting.json, copy.js.
3. The update banner ("Update ready. Quit the companion from the tray icon and open it again to install it.") is a procedure. "Update 0.2.14 ready · Restart" would carry more. This needs the version, which the staged-update JSON (update_staged / updated_to) already has.
4. dmg-window.png has about a third of the window blank under the icons, and the Applications target is a plain blue rectangle with no folder glyph. If this is a mock stand-in for Finder's real icon, say so in the capture. Otherwise shrink the window or use the space for the version and the drag hint. Source: the DMG layout script.
5. Device (device-connected, sign-in-waiting) leaves about 25 percent of the window empty below one card. Fixed content is fine at this size, but two data fields already in the JSON are not used: last_sync (Characters sent, with count and time) and the addon-seen line. Both are on Status, so link to them from here, do not restate them.

## NICE (new feature or new data)
1. The Reports list is 10 rows with no sort, no filter and no "open in site" per row beyond the zone link. Add the credited character as a column, and a fights-in-progress share bar. The data is in reports[].credit, fights and started_at.
2. Status "Characters sent: 2 characters, 4 min ago" has no names, and the status JSON carries a characters array that is empty in the shipped snapshots. When it is populated, render class-crest chips (one crest language) with the names.
3. The tray menu shows only "Logging" as its header. Add the live fight count ("Logging · 3 fights") from pipeline.fights. The tooltip already says "Forever Sixty · Logging".
4. Game folders shows path and last played only. Add the addon version per install and the folder's watching state in the Settings list too. installs[] carries flavor and added.
5. Reports rows have no per-report duration, boss kills or wipes, since the API does not return them. This is a needs-data item, not a fail.

## Unused or barely used data (from testdata and app.js grep)
- token_backend_note: 0 uses in app.js ("A private store in your user profile."). Fine to drop from the payload; "Saved in Windows Credential Manager" already says it.
- advanced_logging_help: 0 uses (1 snapshot). The off-state paragraph hardcodes the Options path, so one source of truth is needed.
- attempts: 0 uses (1 snapshot).
- last_sync.title (full timestamp) is only used as a tooltip, if at all.
- pipeline.offset and queued appear only on the held state, which is correct.
- characters[] is empty in all snapshots, so the Settings credit dropdown and the Status chips are untested with real data.

## Phones (560 equivalent)
Holds. Nav stays one row with icons, the headline count stays top right, and the rows keep label-left, value-right. At 390 the Status headline sub-line should wrap (it already wraps at 560), the Reports Copy link button should stay 44 px tall, and the long folder path should truncate in the middle.

## Needs data
Per-report boss kills and wipes; credited-character class for crests; addon version per install.
