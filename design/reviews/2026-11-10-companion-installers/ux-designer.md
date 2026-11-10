# Round 1

# Companion install experience: UX review

**Verdict: FIX.** The two install windows are close to shippable. Three must-fixes remain: one is missing evidence, one is a copy placement that contradicts the owner's brief, and one is a real state bug in the Mac window.

Reference surfaces: Mac, the stock "Move to Applications?" prompt (LetsMove) in the companion's own round-5 look. Windows, Inno's one-page wizard. The Windows wizard is native, so I judged only the title, subtitle, labels, button and seal.

## What passes

**Mac window (mac-1-move-window.png)**
- It answers one question with one button: the 24px Marcellus title, a two-line lede, one toggle and one teal primary.
- It reuses the shell's wordmark band, panel, `.switch` and `.action.primary.big`. No new control was invented.
- Hierarchy is right. The title is brightest, the button is the only filled element, and the toggle is plain strong text.
- Copy is honest and in game words. There are no em dashes, and "The download goes in the Trash." states the consequence plainly.

**Windows setup (win-1-setup.png)**
- The title, subtitle, two checkboxes and Install button read correctly, and the seal sits in the corner where Inno's box would be.
- "No admin needed." is true: `PrivilegesRequired=lowest` and a per-user directory.
- Defaults are correct: Start with Windows on, desktop shortcut off.

**Windows first run (win-2-first-run.png)**
- The setup lands on the existing first run, "Not connected / Sign in with Battle.net".
- No intermediate Finish page appears, as intended.

**Fallback disk image (mac-dmg-notarized-volume.png)**
- It shows the seal volume icon with the Applications link beside the app, and there is no list-view noise.
- It is a plain, correct fallback. The top-aligned icons leave a large empty window, which is acceptable for a fallback.

## Must fix

1. **The `/setup` Windows capture is a Vite error page, so it is not evidence.**
   - `design/reviews/2026-11-10-companion-installers/setup-windows-1440.png` shows `FailedToLoadModuleSSR … ../../data/generated/combos.json` in an unstyled white page.
   - `setup-mac-1440.png` is also a tight crop of one section, not the 1440 page.
   - Re-capture both as full 1440 pages (plus 360 and 2000, per the wide-screen rule) with the generated data present. Run the generator, or shoot from the built site.
   - I cannot judge the Windows visitor's download row, button label or first-run line until this exists. This blocks SHIP on the web half.

2. **The first-run line is not under the button, which the owner specified.**
   - In `web/src/components/setup/SetupCompanion.svelte` (lines 83-88), the order is `CompanionDownload`, then the `lead` paragraph ("The companion logs live from your desktop."), then the first-run `<p>`.
   - In `setup-mac-1440.png` the signing and install line sits about 80px below the button, with an unrelated sentence between them.
   - Move the `{#each firstRuns}` block above the `lead` paragraph, so the order is button, first-run line, lead, how-to rows.
   - On `unknown`, where both platforms' lines show, keep the "Windows: " and "Mac: " prefixes.

3. **The Mac window's checkbox stays live while moving.**
   - In `companion/ui/install.js` (click handler), the button is disabled and relabelled "Moving…", but `#move-login` is not disabled.
   - A toggle flipped mid-move is silently ignored, because `login` is already read.
   - Add `$('move-login').disabled = true;` next to `button.disabled = true`. Re-enable it in the `catch` block.
   - Add a disabled `.switch` style in `companion/ui/styles.css`, for example `.switch input:disabled { opacity: .5; }` with `cursor: default`, so the state is visible.

## Should fix

4. **The lede widows on one word.**
   - In `mac-1-move-window.png`, "in the Trash." wraps alone onto line 2. `.lede` has `max-width: 60ch` with `text-wrap: pretty`.
   - In `companion/ui/styles.css`, change `.prompt .lede` to `max-width: 52ch; text-wrap: balance;`.
   - That gives two even lines, "It runs from there and keeps itself up to date." / "The download goes in the Trash."

5. **The Mac title asks a question with no visible "no".**
   - The title is "Move Forever Sixty to Applications?" but only the yes button exists. The Go side's behaviour on window close is not shown.
   - Either make the title declarative, "Move Forever Sixty to Applications", or add a quiet text link "Not now" under the button, using the existing muted link style.
   - Whichever you choose, state in `companion/ui/copy.js` what close does, and have the implementer attach a capture of it.

6. **The Mac primary button has no initial focus.**
   - In `companion/ui/install.html`, add `autofocus` to `#move` so Return confirms. A one-button prompt must work from the keyboard.
   - Add `:focus-visible` to the screenshot set (the shared `.action` focus ring should apply).

7. **Evidence is missing for the Mac window's non-default states.**
   - Attach three captures: Moving… (button and toggle disabled), the failed state (the `.reason` alert showing `C.failed`), and the state with no login-item support (toggle row hidden, so the button should not jump).
   - Also check that the failed copy fits in 660 pt without clipping the panel.

8. **The Windows wizard page has about 140px of dead white.**
   - In `win-1-setup.png`, the checkboxes sit in the top third and the lower half of the page is empty above the buttons.
   - In `companion/packaging/windows/companion.iss`, add to the `[Code]` section's `InitializeWizard` (new procedure):
     - shrink `WizardForm.ClientHeight` by about 110px;
     - then move the Next and Cancel buttons and the bevel up by the same amount.
   - Alternatively, add a third muted line to the page that says where it installs.
   - Re-capture after the change.

9. **The Mac and Windows labels differ.**
   - The Mac says "Open at login" and the Windows setup says "Start with Windows". Each is correct for its OS, so no change is needed.
   - The Mac button says "Move to Applications" where the Windows one says "Install". This is also correct for the platform.

10. **The Mac first-run line has a vague tail.**
    - In `web/src/data/companion.json` (mac `firstRun`), "…moves itself to Applications, then sets itself up." repeats the Windows phrasing.
    - Suggested replacement: "Signed by Commish LLC and notarized by Apple. Open the download and Forever Sixty moves itself to Applications, then asks you to sign in. macOS 12 or later."
    - Keep the Windows line. It is honest, and the SmartScreen instruction is the right place for it.

## Spec checklist

| Item | Status |
|---|---|
| Mac: one question, one primary, toggle on by default | pass |
| Mac: round-5 look, shared components | pass |
| Mac: no em dashes, game words | pass |
| Mac: disabled, failed and no-login-item states captured | no (item 7) |
| Mac: toggle locked while moving | no (item 3) |
| Windows: one screen, correct defaults, seal in corner | pass |
| Windows: no wasted space | no (item 8, polish) |
| Windows: opens straight into first run | pass |
| `/setup`: first-run line directly under the button | no (item 2) |
| `/setup`: renderable captures at 1440, 360, 2000 | no (item 1) |
| Honest copy (per-user, no admin, SmartScreen) | pass |

Files: `/Users/jh/code/forever/.claude/worktrees/agent-a401ca0c4adf37296/web/src/components/setup/SetupCompanion.svelte`, `/Users/jh/code/forever/.claude/worktrees/agent-a401ca0c4adf37296/companion/ui/install.js`, `/Users/jh/code/forever/.claude/worktrees/agent-a401ca0c4adf37296/companion/ui/install.html`, `/Users/jh/code/forever/.claude/worktrees/agent-a401ca0c4adf37296/companion/ui/styles.css`, `/Users/jh/code/forever/.claude/worktrees/agent-a401ca0c4adf37296/companion/packaging/windows/companion.iss`, `/Users/jh/code/forever/.claude/worktrees/agent-a401ca0c4adf37296/web/src/data/companion.json`.

# Round 2

# Round 2 re-review: companion install experience

**Verdict: SHIP WITH FIXES.** Nothing blocks the merge. All three round-1 must-fixes are closed. Two small items below should land with it or straight after.

## Round-1 items

| # | Round-1 item | Status | Evidence |
|---|---|---|---|
| 1 | Re-capture the `/setup` pages | Closed | `setup-windows-1440.png` and `setup-mac-360.png` are real pages. The set at 360, 1440 and 2000 is present. |
| 2 | First-run line directly under the button | Closed | In `SetupCompanion.svelte` the `firstRuns` block now comes before `lead`. In `setup-windows-1440.png` the line sits under "Download for Windows". |
| 3 | Toggle disabled while moving | Closed | `install.js` has a `busy()` helper that sets `login.disabled`. In `mac-state-moving.png` the toggle is dimmed and the button reads "Moving…" and is dimmed. |
| 4 | Lede widow | Closed | The lede is one line, "The download goes in the Trash." |
| 5 | Question title with no visible "no" | Accepted | The owner fixed the wording. `install.js` documents that closing the window quits and the next launch asks again. |
| 6 | Initial focus on the button | Closed | The focus ring is visible on the button in `mac-1-move-window.png`. |
| 7 | Non-default state captures | Closed | `mac-state-moving.png`, `mac-state-failed.png` and `mac-state-no-login-item.png` are present. |
| 8 | Dead white in the Windows wizard | Closed | In `win-1-setup.png` the page is about 110px shorter and the buttons sit close under the checkboxes. |
| 10 | Vague Mac first-run line | Closed | The Mac line is now "Notarized by Apple. Open the download and Forever Sixty moves itself to Applications. macOS 12 or later." |

## What shipped, region by region

- **Mac window.** There is one question, one filled button and one toggle. The copy is honest and in game words, with no em dashes. The failed state shows a red `.reason` strip with "Move failed. Reopen Forever Sixty to try again." It fits at 660pt and the button does not jump. With no login item support, the toggle row is hidden cleanly and the panel closes up.
- **Windows setup.** The subtitle "Installs for your Windows account." is accurate, with the per-user guarantee carried by the button and the web line. The seal sits in the corner, the checkboxes read correctly, and the button says Install.
- **`/setup` at 1440.** The download row, then the muted first-run line, then the lead, then the ranked how-to rows. The reading order follows the install journey.
- **`/setup` at 360.** There is no overflow. The first-run line wraps in four lines and the Mac button is full width.

## Remaining items

1. **Should: on `/setup` at 360 the first-run line floats away from the button.**
   - In `setup-mac-360.png` there is a ~48px blank band between "Other platforms" and "Notarized by Apple…", so the line reads as detached.
   - Cause: the `.setup-body` flex gap, plus the download component's bottom margin.
   - In `web/src/components/setup/SetupCompanion.svelte`, give the first-run `<p>` a `margin-top: -16px` at narrow widths, or reduce the gap between the download block and the first-run paragraph to 8px on every width.
   - The same band shows on desktop in `setup-windows-1440.png` (about 25px, so milder).

2. **Should: the Windows line is two sentences too long for a muted aside.**
   - It reads "Signed by Commish LLC. Open the download and press Install. It installs for your account only, no admin needed. If SmartScreen asks: More info, then Run anyway."
   - It is accurate and honest. The cap is the four-line wrap at 360.
   - In `web/src/data/companion.json`, a tighter version drops one clause: "Signed by Commish LLC. Open the download and press Install. For your account only, no admin needed. If SmartScreen asks: More info, then Run anyway."
   - This is optional.

## Notes, no action

- The Mac window title bar still shows the OS title "Forever Sixty" above the wordmark band. That is the native frame, so leave it.
- The disk image fallback is unchanged and fine.
- `win-2-first-run.png` is unchanged and correct.

## Checklist

- [x] First-run line directly under the button.
- [x] Mac states captured: default, focus, moving, failed, no login item.
- [x] Toggle locked while moving.
- [x] Windows wizard trimmed, copy honest.
- [x] `/setup` captures at 360, 1440 and 2000.
- [x] No em dashes, no slop, game words.
- [ ] Tighten the spacing under the button on `/setup` (item 1).

Files: `/Users/jh/code/forever/.claude/worktrees/agent-a401ca0c4adf37296/web/src/components/setup/SetupCompanion.svelte`, `/Users/jh/code/forever/.claude/worktrees/agent-a401ca0c4adf37296/web/src/data/companion.json`.

# After round 2

Applied: the first-run line sits 8 px under the button (one flex column with the download), in body ink, and the Windows line drops one clause. Recaptured setup-*-{360,1440,2000}.png.

# Round 3 (DMG window)

# Review: macOS DMG drag window

**Verdict: FIX.** The cream ground, ink, labels and seal-plus-wordmark are right, and both appearances read. Three items stop it matching the owner's spec and the reference: the window is shorter than designed, the icons are not centred, and the arrow is too long and thin.

I judged the screenshots only. The arrow lengths are my measurements from the 2x capture, not from a recompute of the render.

## What matches

- **Ground and ink:** the dark and light captures differ only in the title bar. Labels are dark ink on cream in both, so they read in both appearances.
- **Icons and labels:** 128 pt icons and 13 pt labels match the spec.
- **Wordmark position:** the seal sits at a 24 pt left margin and about 22 pt from the top of the content area. Discord's mark sits at about 24 pt in both directions, so the placement matches.
- **App icon:** the full blue tile with the centred seal reads as a finished macOS icon. The wordmark seal and the tile show the same art, so they are consistent.
- **Nothing extra:** there is no tagline, toolbar, sidebar or status bar.

## Numbered fixes

1. **Blocker: the window content area is 660x368, not 660x400.**
   - **Evidence:** the capture is 1320x800 px at 2x, and the title bar takes 64 px (32 pt). Finder's bounds include the title bar, so the bottom 32 pt of the 400 pt background is cropped.
   - **Change:** in `tools/package-mac-companion.sh` line 125, set `DMG_WINDOW=(660 432)`. The content area is then 660x400 and matches `WINDOW` in `tools/art/dmg_background.py`. Add a comment that the height is 400 plus the 32 pt title bar.
   - **Verify:** re-capture. The window should measure 1320x864 px, with a 736 px content area below the 64 px title bar.

2. **Must fix: the icons are not vertically centred.**
   - **Evidence:** the icon cells are centred at y=200 in content coordinates, which is 16 pt below the visible centre of 184. The icon-plus-label group centres around y=215 in the 368 pt content area, about 30 pt low. Once item 1 lands, y=200 is the true centre of the icon cell.
   - **Change:** after item 1, the icon-plus-label group is still bottom-heavy, because the label hangs about 30 pt below the icon cell. Set `ICON_Y` to 188 in `tools/art/dmg_background.py` and `DMG_APP=(180 188)`, `DMG_DROP=(480 188)` in `tools/package-mac-companion.sh`. This puts the group's visual centre at about 200.
   - **Keep in step:** `ICON_Y` and the shell numbers share one source of truth. Prefer having the shell read them from the Python module, or at least add a comment tying the two together.

3. **Must fix: the arrow is too long and too light.**
   - **Evidence:** it measures about 104 pt shaft, as long as an icon is wide, and about 36 pt clear of the folder. At 3.5 pt stroke it reads as a ruled line with a head. Discord's arrow is about 32 pt wide with a roughly 10 pt head, and it sits close to the icons.
   - **Change:** in `tools/art/dmg_background.py`, set `ARROW_AIR = 54` (about 68 pt shaft), `ARROW_STROKE = 4`, `ARROW_HEAD = 12`. That is a shorter, heavier arrow with a proportionate head.
   - **Why 54:** the visible tile is about 104 pt inside the 128 pt cell. 54 pt of air from the cell edge leaves about 66 pt clear of each tile.
   - **Keep:** the centre-line alignment is correct (arrow y=463 px against icon centre y=464 px).

4. **Polish: the wordmark is heavier than "small".**
   - **Evidence:** "Forever Sixty" in Marcellus 19 pt is about 110 pt wide. Discord's wordmark is about 85 pt wide.
   - **Change:** in `tools/art/dmg_background.py`, set `NAME_PT = 16` and `NAME_GAP = 8`. Keep `SEAL_PT = 28`.
   - **Hierarchy:** this keeps the wordmark secondary to the icons and does not compete with the labels.

5. **Polish: the bottom of the window looks empty.**
   - **Evidence:** there is about 73 pt of bare cream under the labels, against a tighter top. This is largely a side effect of items 1 and 2.
   - **Change:** apply items 1 and 2 and re-check. If it still reads unfinished, leave it. Do not add a tagline.

6. **Polish: the dark-appearance title bar over a cream slab.**
   - This is a deliberate trade-off, because Finder draws labels dark in both appearances. No change is needed.
   - Add one line to the `tools/art/dmg_background.py` docstring saying the dark title bar over cream is expected.

## Acceptance re-capture (attach)

- Light and dark appearance of the final window, at 2x.
- The window measured at 660 pt wide and 432 pt tall including the title bar.
- A side-by-side with `design/reviews/2026-11-10-companion-installers/reference-discord-dmg.png` at the same scale.
- Both labels fully legible, with no crop at the bottom edge.

Files touched by the fixes:
- `/Users/jh/code/forever/.claude/worktrees/agent-a401ca0c4adf37296/tools/package-mac-companion.sh`
- `/Users/jh/code/forever/.claude/worktrees/agent-a401ca0c4adf37296/tools/art/dmg_background.py`

`companion/tools/icons/render.py` (`mac_seal`) needs no change. After the fixes, re-run `tools/art/dmg_background.py` to regenerate both background PNGs.

# Round 4 (both platforms, the Windows setup in our own window)

**Verdict: FIX.** The shared design is right. Both platforms use one `install.html` in one band-plus-lead-panel layout. The DMG is a faithful take on the Discord reference. The facts check out: per-user, no UAC, a Start menu shortcut, an Apps & Features entry from Commish LLC, and the same-window uninstall with "Keep my settings". Items 1 to 5 must be fixed before ship.

**Must fix**

1. **Blank frame while installing (blocker). `companion/internal/install/setup_window.go`, `runInWindow`, and `companion/ui/install.js`.**
   - `win-2-installing.png` shows the window client area as flat brown. The band, panel and bar are all missing.
   - The 0.5 s run goes from the last step straight to `win.Close()`, so the bar is never seen and the window can drop to an empty frame.
   - Change: hold the window for at least 1200 ms after the click. After the last step, call `fsProgress(of, of, 'done')` so the bar fills before anything closes.
   - Close only after `Start(exe)` has returned.
   - Re-capture `win-2` showing a mid-run state: the bar part-filled with a step line under it.
   - Today `fsProgress` sets the bar at each step's start, so "Starting Forever Sixty" shows 100% while it is still working. Use `(step-1)/of` while a step runs and 100% only on done.

2. **The title keeps asking its question while it works. `companion/ui/install.js`, `button` click handler.**
   - `page-setup-installing.png` shows "Install Forever Sixty?" over a running bar.
   - On click, set `#move-title` to `C.busy` ("Installing…", "Removing…") and hide `#move-lede`.
   - Keep the toggle visible but disabled, as it is now.

3. **An upgrade silently turns "Start with Windows" back on. `setup.go`, `setup_window.go`, `install.js`.**
   - The toggle is always checked on load. Someone who turned startup off, then runs the setup over the install, gets it re-enabled by `SetStartup`.
   - Add `StartupEnabled() bool` to `WinSystem`.
   - Return it from `fsSetupInfo` as `{upgrade, startup}`.
   - In `install.js`, set `toggle.checked = info.startup` when `info.upgrade` is true.
   - Add a test in `setup_test.go`.

4. **Upgrade still says "Install". `companion/ui/copy.js` and `install.js`.**
   - `win-4-upgrade.png` has the new lede under the old question and button.
   - Add `titleUpgrade: 'Update Forever Sixty?'`, `actionUpgrade: 'Update'` and `busyUpgrade: 'Updating…'`.
   - Apply all three in the `info.upgrade` branch, and make `showFailed` restore the right label.
   - In `win-4` the button is a lighter teal than in `win-5` and `win-1`. Re-capture with the pointer parked, to show whether that is hover or a transition caught mid-way.

5. **Failure copy contradicts the button on screen. `companion/ui/copy.js`.**
   - `showFailed` re-shows the button, but the text says to go elsewhere. Replace the three strings:
     - setup: `'Install failed. Press Install to try again.'`
     - uninstall: `'Remove failed. Press Remove to try again, or use Apps & Features.'`
     - move: `'Move failed. Press Move to Applications to try again.'`
   - Add `mac-state-failed.png` and a Windows failed capture to the acceptance set.

**Polish**

6. **The destructive button is autofocused. `install.js`.**
   - `win-6-uninstall.png` shows a gold focus ring on a red "Remove", so Enter removes the app.
   - In uninstall mode, remove `autofocus` and focus the toggle instead.
   - Keep the focus ring for keyboard users.

7. **The toggle does not say what it keeps. `copy.js`, `install.uninstall.lede`.**
   - The lede only says "Your reports on foreversixty.gg stay."
   - Change it to: `'Your reports on foreversixty.gg stay. Turn off Keep my settings to also forget this device's sign-in.'`
   - This matches `RunUninstall`, which calls `ForgetToken` only when the toggle is off.

8. **The Mac copy contradicts its own picture. `web/src/data/companion.json`, Mac `firstRun`.**
   - The DMG window shows an arrow from the app to Applications.
   - The page says "open Forever Sixty in it: it moves itself to Applications".
   - Change it to: `Notarized by Apple. Open the download and drag Forever Sixty onto Applications. macOS 12 or later.`
   - The move prompt is then the fallback for someone who double-clicks the app inside the image. `setup-mac-1440.png` should show the new line.

9. **The two Mac windows use different palettes. `tools/art/dmg_background.py`.**
   - The DMG window is cream with a brass arrow. The move prompt is dark with a teal band.
   - The cream ground is justified: Finder draws its labels in dark ink, so they need a light ground, and the Discord reference is white. Keep it.
   - Add the same teal band (about 60 pt, `base.raised` to teal) behind the wordmark so the DMG carries the brand's colour. Re-capture the light and dark DMG windows after.

10. **Setup progress step copy. `copy.js`, `install.steps`.**
    - "Copying" is the only step without an object, so change it to `'Copying files'`.
    - Match the Start menu wording to your fact list: `'Adding the Start Menu shortcut'`.

11. **The uninstall done state jumps. `install.js`, `fsDone`.**
    - The panel shrinks to one line for 1.5 s (`win-7-removed.png`), then the window closes.
    - Keep the lede visible under the title: "Your reports on foreversixty.gg stay."
    - Also set a `min-height` on the panel so it does not resize.

12. **The caption shows a disabled maximize and an enabled minimize. Window creation, `companion/internal/install/offer.go`.**
    - These are visible in `win-1`, `win-2` and `win-6`. For a fixed 660-wide dialog, remove both buttons so only close remains.
    - If the webview library cannot do that, clear the `WS_MAXIMIZEBOX` and `WS_MINIMIZEBOX` window styles.
    - I could not confirm the mechanism, because the window-creation code is not among the files you gave me.

**Accepted**
- Two wordmarks: the caption text and the band. Discord does the same.
- The dark caption in both system themes.
- 150% scale not captured.
- `win-5` matches `win-4` because the app ignores the system theme.

**Acceptance set for the re-review**
- Windows at 100%: setup, a mid-run state, upgrade (button parked), uninstall, uninstall done, failed.
- Mac: the move prompt and a failed state.
- Both DMG windows, light and dark.
- `/setup` at 1440 for both platforms.
