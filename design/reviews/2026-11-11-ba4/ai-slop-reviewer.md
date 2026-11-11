# AI-slop review, lane BA4 (companion window, install steps, tray, DMG)

Reviewed: dmg-window, tray-menu, install-move, install-setup, reports, logging-credited, settings, first-run-device captures, plus apps/shared/copy/companion.json. Not opened one by one: the remaining state captures, which share the same chrome.

## First impression
Made by someone who plays WoW. Parchment, the ringed LX seal, serif titles on a painted teal band and a dense Status page read as a game-UI tool, not a template. No hero slogan, no gradients on controls, no glow, no stat tiles, no eyebrow labels.

## Tells found (most damaging first)
1. (15, risk) dmg-window.png: the big LX seal is glossy and over-lit, with a heavy dark outline. The 32 px crest in the same window's band is flat. Two renderings of the mark in one window read as "one is a render". If the large one is the brand asset (design/assets), keep it and use the same crest as the band and the app. If it was generated, replace it.
2. (22, 19) reports.png: page title "Your reports" with the link "Open your reports" directly across from it. The link repeats the heading. A player would write "Open on site" or "foreversixty.gg/logs". Tray "Open your reports" is fine because the menu has no heading.
3. (22) logging-credited.png: "Characters sent / 2 characters, 4 min ago" (label and value repeat the noun). Addon row: label "Addon", value line "Addon 2.0.16 · Data 2026.11.09 · up to date" (repeats the label). "3 / fights uploaded" is fine because the number is the point.
4. (7) settings.png: a grey explainer under every control (visibility, keep addon, credited-to). Keep the ones that carry information ("CurseForge and Wago users can leave this on; they get the same files" is good). Cut "Pick a character only if Automatic is wrong" (the dropdown says Automatic) and the visibility lines that restate the option (private: "Only you.", guild: "Members of your guild.").
5. (21) copy: "Uploading your fights as they end" under "Logging", "The first fight uploads when it ends", "This download goes in the Trash" (move lede, true but every Mac user knows it), "Sign in on foreversixty.gg so the companion can upload your fights and keep the addon up to date" (pitch-length for a device page whose title already says Not connected). Shorten or drop.
6. (12, low) Parchment background carries visible mottled noise behind every window. It reads as the in-game parchment, which is a specific reason, so it is not a fail. Yellow link text (Open your reports, Add a folder…, Having trouble?) on the brown parchment has weak contrast where it sits off the dark panel (reports.png header link).
7. (8, low) Dark panel with the same border on every block, ornate corner frame on the hero and device cards. Used by role (hero and primary get the ornament, lists plain), no cards inside cards. Acceptable. Do not add the ornament to a third level.
8. (4, low) dmg-window.png: cream field with serif title. The accent is teal, not terracotta, and the band carries the identity, so it does not match the tell. The cream appears nowhere in the app (parchment brown and near-black), so the DMG looks like a different product. Brand-consistency reviewer owns that call.
9. (mock only) dmg Applications target is a flat blue gradient square. Finder draws the real folder icon at ship, so this is a mock placeholder. Mark it as such in the capture so it is not judged as a design choice. Bottom third of the window is empty; fine, but the arrow is generic thin stroke (SHOULD be the only decoration, keep it).

No em dashes found in companion.json. No emoji. No fake proof. No Inter/Poppins: serif display plus a humanist sans body. Nav icons are consistent and functional.

## Authentic, keep
- Real report content: Molten Core, Blackwing Lair, Zul'Gurub, Ruins of Ahn'Qiraj, Scholomance, with day, time and fight counts. In-progress green dot, hollow dot for not on site, amber "not credited yet" are real states.
- "Advanced logging: on", the Options > System > Network > Advanced Combat Logging path, "Log out once in game with the addon loaded", /combatlog, the _forever_ folder name, the Program Files path in mono. Veteran copy.
- "Open your reports" in the tray with a plain native menu, tooltip "Forever Sixty · Logging". Restrained, nothing invented.
- Install prompts: one question, one toggle, one button. "Your reports on foreversixty.gg stay." is the right reassurance.
- Uneven hierarchy: Logging hero with the big fight count, plain rows below.
- Failure copy names the log file and the next step.

## Verdict: PASS WITH FIXES
No visual tell 1, 2, 5, 8, 12 or 13 present. Tell 15 depends on item 1.

MUST
1. Confirm the DMG seal is the brand asset (design/assets), not a generated render, and render the mark identically in the DMG band, DMG centre and app header. Fail tell 15 if it is generated.
2. Replace the reports-page link "Open your reports" with a non-repeating label (tray keeps it). Raise to the copy owner, since BA4 chose this name.

SHOULD
1. Remove the redundant words in "Characters sent" and "Addon" rows (item 3).
2. Cut the explainers in item 4 and item 5; keep the CurseForge/Wago line.
3. Give the off-panel yellow links a darker backing or move them onto the panel (contrast on parchment).
4. Label the DMG Applications square as a Finder placeholder in the capture.
5. Reconcile the DMG cream with the parchment/dark app palette, or note why it differs.
