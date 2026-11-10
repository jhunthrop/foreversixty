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
