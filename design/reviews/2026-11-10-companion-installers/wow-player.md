# Round 1

# wow-player review: companion install experience

**Verdict: FIX.** The Mac window and the Windows installer both read well. I would trust either in one glance. Two things stop me from saying SHIP: the /setup capture is a build error page rather than the page, and the first-run screen has no way to confirm the install worked.

## What a player sees

**Mac move window** (mac-1-move-window.png). The question "Move Forever Sixty to Applications?" is the first thing on screen. It has one toggle and one button, and the crest and parchment frame match the companion. I would click it without thinking. "Open at login" is on by default, which is what an uploader wants. It is a single honest toggle, so it is not naggy.

**Windows setup** (win-1-setup.png). It is a single page with two ticks and Install. "The companion for your Windows account. No admin needed." is the right reassurance for a veteran who has been burned by installers. The crest is in the corner and nothing is bundled. "Start with Windows" is on and "Desktop shortcut" is off, which is the correct split.

**First run** (win-2-first-run.png). "Not connected" with one Battle.net button is clear. "Opens your browser. Your Battle.net password never reaches this app." is a good trust line.

## Numbered fixes

1. **/setup was not rendered, so I cannot approve it.** setup-windows-1440.png is an Astro "FailedToLoadModuleSSR: Could not import ../../data/generated/combos.json" error page. This is the generated data missing in the worktree, not a copy problem. Re-capture `web/src/pages/setup.astro` for Windows at 1440 after the data is generated, and re-run it. setup-mac-1440.png did render, and that is the only /setup surface I can judge.

2. **The Mac line under the button is a pitch and reads as a sequence.** File: `web/src/data/companion.json`, mac `firstRun`. It currently says "Signed by Commish LLC and notarized by Apple. Open the download and Forever Sixty moves itself to Applications, then sets itself up. macOS 12 or later." On screen it sits in grey under the line "The companion logs live from your desktop." Change it to: "Signed by Commish LLC and notarized by Apple. Open the download and it moves itself to Applications. macOS 12 or later." Players do not need "then sets itself up", because step 1 already covers the sign-in. Also drop "The companion logs live from your desktop." It tells a veteran nothing the header does not.

3. **The Windows line promises something the screen does not show.** File: `web/src/data/companion.json`, windows `firstRun`. It says "Open the download and Forever Sixty sets itself up" and then explains SmartScreen. The player actually meets a Setup window with two ticks and an Install button. Change it to: "Signed by Commish LLC. Run the download, leave the ticks as you like, press Install. Your Windows account only, no admin. If SmartScreen asks, choose More info, then Run anyway." Keep the SmartScreen line, because that is the thing that makes people quit.

4. **The first-run lede still says "up to date" in two senses.** File: `companion/ui/copy.js`, `install.lede`: "It runs from there and keeps itself up to date. The download goes in the Trash." The first half is fine. "The download goes in the Trash." is the line a player will stop on. A veteran asks whether the zip gets deleted. Change it to: "It runs from there and keeps itself up to date. The downloaded copy goes to the Trash." Using "copy" makes clear that nothing else is touched.

5. **The Windows title bar and header repeat the version without a reason.** File: `companion/packaging/windows/companion.iss`. The title "Setup - Forever Sixty 0.2.45" is Inno's default and is fine. The header "Install Forever Sixty" is also fine. However, the empty lower half of the window, roughly 60 percent of its height, looks unfinished. Either shrink the wizard height in `[Setup]` (`WizardSizePercent` or `WizardStyle=modern` with a smaller size) or add one line of what Install does, for example "Installs to your user folder and opens the companion." Choose the shrink and do not add copy.

6. **There is no confirmation that the move or install happened.** File: `companion/ui/copy.js`. After Move or Install the player lands on "Not connected", which is the same screen as a first launch from anywhere. A player cannot tell it is now running from Applications or its install folder. This is a minor gap. Add a one-line toast or a line in the Device footer such as "Installed. Opens at login." (`install.done`), shown once. Otherwise it is fine to ship without.

7. **The failure line is a dead end.** File: `companion/ui/copy.js`, `install.failed`: "That did not work. Close this window and open Forever Sixty again to try once more." "That did not work" is vague. Say what to check: "Could not move it. Drag Forever Sixty into Applications yourself, then open it from there." A veteran can always drag an app. The retry advice may fail repeatedly.

8. **The uninstall question is good but one clause is odd.** File: `companion/packaging/windows/companion.iss`, `[Code]` message. "Also delete the companion's settings, logs and any fights still waiting to upload? Your reports on foreversixty.gg stay either way." Defaulting to No (`MB_DEFBUTTON2`) is correct. "logs" may be read as WoW combat logs. Change "logs" to "its own logs", or drop it, so no player fears it deletes their WoWCombatLog files.

## Missing

- No way to see where it installs on Windows before pressing Install, since the one-screen setup hides the folder. That is acceptable for a per-user install, but "Your Windows account only" in the web line should say where it lives only if a player asks. I would not add it.
- The Mac toggle gives no hint that it can be changed later. If Settings has "Open at login", the lede could say so, but it is not needed.

## Data and claims I do not believe

- "notarized by Apple" and "No admin needed" are claims about the build, not the screens. I cannot verify them from screenshots, so whoever checks accuracy should confirm the notarization ticket and the `PrivilegesRequired=lowest` setting.
- "macOS 12 or later" should match the build's `LSMinimumSystemVersion`.

## Summary

Items 1 to 3 are needed before shipping. Items 4 and 7 are copy fixes worth making now. Items 5, 6 and 8 are small and can follow. No em dashes or slop phrasing appeared in the copy I read. The tone matches how a veteran would write it.

# Round 2

# Round 2 verdict: SHIP WITH FIXES

All round-1 blockers are closed. I re-viewed mac-1-move-window.png, mac-state-failed.png, win-1-setup.png and setup-windows-1440.png. I did not re-open the Mac /setup captures or the 360 and 2000 versions, so the Mac /setup line is judged on the copy you quoted, not on a render.

## Round-1 items

1. **/setup not rendered: closed.** setup-windows-1440.png is now the real page. The companion card reads "Signed by Commish LLC. Open the download and press Install. It installs for your account only, no admin needed. If SmartScreen asks: More info, then Run anyway." That matches what the player sees, and the SmartScreen line is still there.
2. **Mac pitch line: closed.** "Notarized by Apple. Open the download and Forever Sixty moves itself to Applications. macOS 12 or later." is plain and true. Dropping "Signed by Commish LLC" on the Mac is fine, since notarization is the thing Gatekeeper shows.
3. **Windows line promises vs. screen: closed.** "Press Install" matches the one button.
4. **Lede: closed.** "The download goes in the Trash." is short and honest.
5. **Empty window: closed.** The Setup window is visibly tighter with 110px trimmed, and "Installs for your Windows account." is enough.
6. **No done toast: accepted.** The first run follows at once, so the player sees it worked. I agree.
7. **Failure line: accepted.** "Move failed. Reopen Forever Sixty to try again." reads fine in the red strip, and the button stays live. I accept the owner's no-drag rule.
8. **Uninstall "its own logs": closed.** This is not a screenshot item. I take it from your message, and it removes the fear that WoWCombatLog files get deleted.

## Remaining fixes

1. **`web/src/pages/setup.astro` (the /setup header, step 3 and companion card): three "Download for Windows" buttons plus two "Other platforms" links on one screen.** A player scrolling the Windows /setup at 1440 sees the same button in the header, in step 3 and in the card. Keep the header one and the card one and make step 3 a plain "Install the companion below" row. This is minor and was there before this lane.
2. **`web/src/data/companion.json`, both `firstRun` lines: low contrast.** The grey text under the button on the dark card is hard to read on the Windows capture, and that is the line with the SmartScreen instruction. Raise it to the body text colour, or at least the same grey as the step text, so the "More info, then Run anyway" part is readable. This is the line players need when SmartScreen blocks them.
3. **"The companion logs live from your desktop." still sits between that line and step 1.** I accept that it belongs to another lane (addonCopy), but it still tells a veteran nothing. Log it as a follow-up for that lane.

## Not verified

"Notarized by Apple", "no admin needed" and "macOS 12 or later" remain build claims for the accuracy reviewer, as in round 1.

Fix 2 is worth doing before release. Fixes 1 and 3 can follow.

# After round 2

Applied fix 2: the first-run line is in body ink, not muted. Fixes 1 and 3 belong to the /setup page and addonCopy lanes and are left as follow-ups.

# Round 4 (both platforms, the Windows setup in our own window)

**Verdict: FIX.** It is close. A player would trust the signed, one-window, no-admin story, and nothing here nags. Item 1 is a blocker until it is recaptured, and items 2 to 4 are real defects.

**Capture limits.** The Windows captures sit on a CI runner with a console behind them, which I ignored. I could not check the Mac states I wasn't given (`mac-state-*`). I did not open the `.html` or `.go` sources, so I name them from the copy comments rather than from reading them.

1. **`win-2-installing.png` shows a blank brown window.**
   - There is no header, no progress bar and no step line. `copy.js` `install.steps` ("Copying", "Adding the Start menu shortcut" and so on) is nowhere on screen. A player sees an empty box and assumes it hung.
   - Check whether the capture caught the window between states. If not, render the card, a bar and the current `steps` line from the first frame of the install state.
   - Recapture after the change, showing at least "Copying".

2. **`setup-mac-1440.png`: the manual-install path uses Windows backslashes on the Mac page.**
   - It reads `World of Warcraft\_forever_\Interface\AddOns`. The site's setup page template needs the path per OS. On Mac it should read `/Applications/World of Warcraft/_forever_/Interface/AddOns`.
   - A Mac player will not take a backslash path seriously.

3. **Check the combat-logging path on the `/setup` companion steps.**
   - It says "Turn on Options > Network > Advanced Combat Logging once." I am not sure the Era-based client puts that checkbox under Options > Network. I recall it under Interface > Network.
   - Verify it in the Forever client. Fix the string in `web/src/data/companion.json`, or wherever the steps live, to the exact menu path. Also quote the checkbox label as the client spells it.

4. **`mac-1-move-window.png`: the lede "The download goes in the Trash." is ambiguous.**
   - A player does not know whether "the download" is the disk image or the app they are running. It also sounds like a warning. In `companion/ui/copy.js` `install.move.lede`, say what happens: "Moves the app to Applications and puts this download in the Trash."
   - Keep the title as it is.

5. **`copy.js` `install.setup.toggle` and `install.move.toggle`: both start-at-login toggles default on.**
   - That is acceptable because it is visible, one line and one click. It is the only thing that could read as naggy.
   - For trust, add a muted second line: "You can change this later in Settings." Confirm that the setting actually exists.
   - Keep the toggles as they are. Don't add a prompt after install.

6. **`win-7-removed.png` ends on a dead end.**
   - "Forever Sixty is removed." is good copy, but the window has no button. Either auto-close after a few seconds or add a Close button. Today the player has to find the X.
   - The same applies to `win-3-first-run.png`: it opens straight onto "Not connected" with the Battle.net sign-in. That is right. The sign-in is the only next step, and nothing else is pushed at the player.

7. **`web/src/data/companion.json` `firstRun` for Windows: "Signed by Commish LLC" means nothing to a player.**
   - It is the right fact, because SmartScreen shows that publisher. Reword it so the player can match it: "Publisher: Commish LLC (the company behind Forever Sixty)."
   - Keep "If SmartScreen asks: More info, then Run anyway." That is exactly what players need.

8. **`mac-dmg-window-*.png`: the drag window is clean and reads at a glance, but the lower third is empty.**
   - Light and dark both work. The cream body in the dark window is fine because it is a fixed background.
   - Optionally add one small line under the arrow, "Drag to Applications". Some players don't recognise the folder badge.
   - Also move the icons about 40px lower, or shorten the window, so the content is centred.

**What works and should not change.** The one-window flow on both platforms is a single decision with a single button. The upgrade state says "Replaces the copy already installed. Your settings stay." Uninstall reassures with "Your reports on foreversixty.gg stay." and "Keep my settings". The red Remove button is correctly alarming. Both setup pages mark the companion as "Optional", and the Download buttons are detected per OS.

# Round 5

**SHIP.** The blocker is fixed, and the declines are reasonable except that two of them need a ticket.

**What I checked in the fresh captures**
- **`win-2-installing.png`:**
  - It now shows the "Installing…" title, a greyed-out "Start with Windows" toggle, a progress bar and the step line "Starting Forever Sixty". A player sees something working instead of an empty box. The blocker is gone.
  - I only saw a single frame, with the bar part-full. I can't confirm the 1.2 s hold you describe.
  - The step list in `round-4-applied.md` ("Copying files", "Removing files") fits what the screen shows.
- **`win-4-upgrade.png`:** "Update Forever Sixty?" with an "Update" button and "Replaces the copy already installed. Your settings stay." This is what a player expects. Keeping their Start with Windows choice on update is the right call, since resetting it would have felt naggy.
- **`win-7-removed.png`:** "Forever Sixty is removed." with "Your reports on foreversixty.gg stay." is good. Minimize and maximize are gone from these dialogs, which is correct for a one-question window.
- **Copy:**
  - `companion/ui/copy.js` now reads "This download goes in the Trash." That is clear enough.
  - `web/src/data/companion.json` opens with "Verified publisher: Commish LLC." That is SmartScreen's own wording, so the player can match it. Good.

**Declines I accept**
- **No "drag" line on the DMG.** The owner has ruled on it, and the drag window with its arrow reads without text.
- **No "change it later in Settings" line.** Don't promise a setting that doesn't exist.
- **Brand items (the painted ground, the seal, the band).** Those are the owner's call, not mine.

**Two items I would not let drop**
1. **The `/setup` hand-install path shows Windows backslashes on the Mac page.** `World of Warcraft\_forever_\Interface\AddOns` appears under "Installing by hand" on `setup-mac-1440.png`. Mac players follow that text. Log it as a real follow-up for the lane that owns `addonCopy`, not a nice-to-have.
2. **The combat-log menu path "Options > Network > Advanced Combat Logging" is still unverified against the Forever client.** It sits in the companion's three-step list, so every new companion user reads it. If the checkbox lives elsewhere, the player is stuck at step 2. Someone should check the client and fix the string once.

**One small remaining nit, not a blocker:** the "removed" window still has only the X to close it. A Close button or an auto-close after a few seconds would finish it.

I did not reply to the other session. Everything above is for you.
