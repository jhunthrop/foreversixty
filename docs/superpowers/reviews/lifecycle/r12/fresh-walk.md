# Fresh-account walk — 21 Oct 2026, evening (production)

The owner's account was deleted from the production database (users row 71 and its six characters, after a CSV backup; 8 other accounts, 3 reports and both guild rows untouched), the addon's saved variables were backed up and removed from the Classic Beta client, and the released companion 0.2.16 was run from an empty home as a first run. The beta realms were closed, so the in-game half (first-login welcome, `/fs`, logout) could not be walked; the companion's upload was exercised by putting the addon's saved export back as a logout writes it.

| Stop | What happened | Capture |
|---|---|---|
| Site as a stranger, before anything | `/character/us/pvp/bow-jackzon` says "No character called Bow Jackzon on PvP US yet." with the lookup and "Your character? Get the addon · Upload logs". `/setup` reads three steps, sign-in first. | — |
| Companion first run | "Not connected · Your sign-in is no longer valid. Sign in again to keep uploading." — Windows still held the old device token, so a brand-new machine would read the plain connect state instead. One gold "Sign in with Battle.net" with "Opens your browser. Your Battle.net password never reaches this app." | `fresh-01-companion-first-run-565.jpg` |
| Device sign-in (owner's clicks) | Battle.net sign-in in the owner's browser, approve → "this device is connected" at 19:49:52; production has user 78 and device JH-DESKTOP. | — |
| Status after connect | "Ready · Waiting for a combat log…", "Addon · Not synced yet · Install ForeverSixty, then log out once.", "Addon 2.0.15 · up to date", "Data 2026.10.21 · up to date", "Characters sent · No characters sent yet. In game, log out once." | `fresh-02-companion-connected-status-565.jpg` |
| The addon's export arrives | Saved export put back → "uploaded character exports … characters 6" within a second; Status reads "Last synced 1 minute ago" and "Characters sent · 6 characters, 1 minute ago"; production has the six characters under user 78. | — |
| Character page as a stranger | Level 24 High Order Skyborne, "Beast Mastery · from your talents", "4 of 14 slots best in slot", every slot with its best pick and source, Ring 1 "Worn item not in our list → Protector's Band · Not scored" (the heirloom ring, W.13), "No logs yet · Log a raid with the companion". | `fresh-03-character-stranger-1280.jpg` |

## Findings

1. **No signed addon or data release exists** (`addon-v2.0.15` and every `data-v*` predate the signing step), so the companion's updater logs "the update check found no signed addon release" and a player without CurseForge gets no addon from the companion. One line per unsigned data tag is logged too (24 lines on every check). Fix: tag `addon-v2.0.16` from main; the nightly data release signs from tomorrow.
2. **Install offers for games the player does not play.** The Status card offers "Install the addon into Classic Anniversary" and "into Retail" to a beta-only player, and "Addon not found in Classic Anniversary / Retail" sit between the beta's own lines.
3. **The inbox is written into every WoW account folder on the machine** on connect and on every inbox change: nine `ForeverSixtyInbox.lua` files across `_anniversary_`, `_retail_` and `_classic_beta_`, including accounts that never ran the addon.
4. **A stale device token reads as "no longer valid"** on a first run that is only fresh on the site's side; harmless, and a truly new machine never sees it.
5. **"Install ForeverSixty, then log out once."** shows under Addon while the addon is already installed and up to date two lines below; the line means "no export yet", and could say so.

Still owed (owner): the in-game first-login welcome and `/fs help` when the beta reopens; the signed-in stops (the character as yours, the guild claim, the planner send, the sim) once the browser pane is signed in.

## Owner's ruling on findings 2 and 3 (21 Oct, 20:00)

"We should only support Forever variants — for right now that's just Classic beta." Lane C.23 (`briefs/task-C.23-forever-variants-only-brief.md`): the companion lists, watches, offers and writes to Forever installs only; Retail and Anniversary disappear from Game folders and from the inbox writer; the updater's unsigned-release lines fold into one.

## In game (22 Oct, 01:11 server time, the beta back up; the owner logged in, the walk drove the client)

| Stop | What happened | Capture |
|---|---|---|
| First login | The welcome printed in General: "Forever Sixty is on. The small Forever Sixty line beside your minimap shows your next step." / "/fs opens the build window; builds you send from the site wait there. /fs help lists the commands." / "The addon saves your character when you log out; the Forever Sixty companion on your PC sends it to foreversixty.gg." (re-printed with `/fs welcome` for the capture). | computer-use zooms in the tool results; see finding 6 |
| `/fs help` | Every command listed, one line each, with the two "For testers" lines last. | — |
| `/fs` | The build window: "Following: Beast Mastery build · Hunter · 51 points", "Next talent: Take Endurance Training (Beast Mastery, 1 of 4)", the paste box "Paste a build code from foreversixty.gg (Ctrl+V)", "Builds you pick on foreversixty.gg arrive here through the companion app. Run it on your PC.", "Tab to switch". | — |
| Unit tooltip | Hovering another player: "Forever Score: no logs yet" and "Target them, then /fs link: site page" under the class line. | — |
| `/fs diag` | "Export: built 01:11 pm, 15 slots, talents 50202105/0/0", "Site: seen yesterday 07:51 pm (6 characters)", "Inbox: no builds waiting · 1 guild message". | — |
| `/logout` | The character screen; the companion logged "uploaded character exports" within a second; the export row on production was unchanged because the export was identical to the one already stored (right). Enter World put the owner back. | — |

6. **The minimap line is not on screen.** The welcome's first sentence points at "the small Forever Sixty line beside your minimap", but nothing is drawn beside the minimap at level 24 in Darnassus; the saved settings say `shown = true`, `hiddenUntilLevel = 0`, not in combat, talent window closed (the only three reasons `Chip.state` hides it), and `/fs chip` twice answered "Forever Sixty line hidden until level 25" then "Forever Sixty line shown" without a line appearing either time. Only the round minimap button (LX) is there, and the icons above the minimap are Blizzard's buffs (Aspect of the Monkey), not the chip's idle icon. Lane A.1 (`briefs/task-A.1-chip-missing-at-first-login-brief.md`) finds the cause and makes the idle line say the next step. To find: the frame's anchor from the saved `x`/`y` (a position dragged off-screen survives), or a `Chip.disabled`/no-data path.
7. **The nightly data release is signed** (`data-v2026.10.22` carries a `.minisig`), the first signed release; the addon still has none.
8. **The nightly web run is red on the known Playwright list** (item-tooltip ×3, home-panel:226, bis-landing ×2, nav-selector:480, both projects) plus the Lighthouse LCP budget; the deploy itself passed. Those specs fail on every main push and should be fixed or retired.

**A.1 verified in the client (22 Oct, 01:31 server time, a normal login after `/logout`):** the line reads "Next: Endurance Training (1 of 4)" beside the minimap with the class icon (`fresh-04-chip-after-a1-login.jpg`). The lane's finding: the saved spot (`point = TOPRIGHT`, `x = -224.72`, `y = -41.94`) put the chip under the client's buff icons at the same strata, and the idle chip had no words at all; a saved spot off screen could also disable it for the session. Merged `c20050f8`; the welcome capture is `fresh-05-welcome-lines.jpg`.
