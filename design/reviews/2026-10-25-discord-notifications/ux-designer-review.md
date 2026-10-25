# Discord notifications: ux-designer review

Reviewed 2026-10-25 on the captures at head 8822bf07, against `design/specs/2026-10-25-discord-notifications.md` (sections 3, 4, 5, 6.1, 7, 9; section 10 and 4.1 as amended) and `design/mocks/renders/discord-notifications-states.png`. Every panel crop, the full pages at 1440, 2000, 1024, 768, 390, 360, the Premium row crops, the focus and hover crops and both skeletons were viewed. Known and ruled, not re-raised: skeleton reserves no block height; Plan "Not on Premium" beside a live panel and "Your reports did not load" are fixture artefacts; native checkboxes.

VERDICT: SHIP WITH FIXES

Fixes (both small; no design change):
1. `discord-free-1440.png`, `discord-free-panel-1440.png`, `discord-free-panel-360.png`, `discord-focus-see-plans-1440.png`: the free sentence is not 44 px high and its two lines are 31 px apart. Spec 3.2 (Free) and 5: one sentence, min height 44 px, `inline-flex min-h-11` "on phone". Smallest fix: put `min-h-11` on the paragraph (flex, items-center), not on the `See plans` link inside running text; the link keeps normal line height from md up.
2. `discord-saving-*.png` and `discord-sending-*.png`: they do not show the pending states. `saving` shows the post-save view with `Saved. Send a test message to see it in your channel.`; `sending` shows `Sent. Check the channel.` with the button hovered. Spec 9 items 2 and 3 require `Saving…` and `Sending…`. Smallest fix: re-run the capture so the screenshot is taken while the held route is pending (the spec file waits for the text, so the PNGs on disk are stale or the shoot ran after the release); attach new files at 390 and 1440.

## Must fix before merge

1. Free sentence line box (fix 1 above). Shipped at 1440: text at y 921 and y 952, the second line holds only `done. See plans` (orphan). Focus ring on `See plans` (`discord-focus-see-plans-1440.png`) is a 44 px box that overprints the end of line one (`messages` is clipped by it). Spec 3.2, 4.2 (`See plans` focus: 2 px gold outline, 2 px offset). At 360 the same 31 px line gap shows.
2. Missing pending-state evidence (fix 2). Spec 9 items 2 and 3, 4.2 (`Saving…`, `Sending…`, disabled, `aria-busy`).

## Should fix this lane

1. Gap above `Use a different webhook` is about 74 px from the buttons (`discord-saved-panel-1440.png`: buttons end y 321, link text y 397; `discord-side-by-side-1440.png`: mock 53 px, shipped 75 px). Cause: buttons, 16 px gap, 21 px reserved status line, 16 px gap, 44 px link. The mock governs: 8 px between buttons and status line, 0 between status line and link gives about 51 px. Same in `discord-switched-off-panel-1440.png` (255 to 329) and the warning crops. Spec 3.2 Saved 6 and 7 ("directly under the status line"), 5 whitespace. The status line stays reserved, so no shift.
2. `Mana….` in `discord-warning-long-panel-1440.png` and `-360.png`: ellipsis then full stop. Spec 6.1 Warning rows: one trailing full stop. Smallest fix: when the text was truncated, end with the ellipsis only. Also cut at the last space at or before 140 characters, not mid-word (`Mana…`); spec 6.1 "shortened to 140 characters" allows it. The switched-off notice already ends `Mana…` with no stray full stop; the warning must match it.
3. Failure warning weight (`discord-warning-panel-1440.png`, `-long-`): 13 px `--muted` is what 3.2 says, and it reads as small print while the panel is one failure from switching off. Ruling (amends 3.2 row 3): split into two lines in the switched-off notice grammar without a bar: line one `The last 4 posts failed. At 5 in a row it switches off.` in 13 px `--text`; line two `Last error from Discord: {last_error}.` in 13 px `--muted`. Same words, same colours, no new token. Copy in 6.1 is unchanged.
4. Switched-off height at 390: `discord-switched-off-panel-390.png` is 524 px tall (572 minus the 48 px crop padding); spec 7 allows 520. The helper `Sends a test message. If Discord accepts it, posts resume.` wraps to two lines. Fixing item 1 above (about 22 px) brings it to about 502. At 360 with the 140-character error (`discord-switched-off-long-panel-360.png`, 582 px) the limit is exceeded by 62 px; accept this as the worst case but record it in the README.
5. Channel id font: README says body face; by advance width (19 digits in 147 px, 7.7 px each) it is the same monospace as the mock (148 px), and the code sets `font-mono text-[13px] tabular-nums`. I could not prove the family from pixels. Check computed `font-family` is `JetBrains Mono` (only the 500 subset is bundled) and close the item; if it resolves to the system mono, fix the token mapping. Spec 3.2.

## Polish for a later pass

1. `discord-saved-1024.png`: Premium row value wraps `On · 2` / `events` and the underline breaks across lines. Use non-breaking spaces in the value so it wraps whole-value or not at all. Spec 3.1 accepts wrapping; this wrap is ugly. The Plan panel's other values (`Browser` / `lane only`) wrapping at 288 px is the existing panel, not this lane.
2. Native checkboxes are `#0075ff` blue beside gold and ink panels (all saved and form captures). Ruled as the anonymize row; when that row moves, set `accent-color` to `--gold` on both so the page stops mixing a browser blue into the palette.
3. `discord-paste-empty-panel-1440.png`: about 40 px of dead space between the input and the `In Discord:` helper (reserved 21 px error line plus gaps). Pull the helper up 8 px by taking the gap off the reserved line's top.
4. `discord-remove-confirm-*.png`: `Use a different webhook` stays visible under a pending remove question, 94 px below the buttons. Hide it while the question is open.
5. README `Verdicts` and `Deferred, not this lane` sections are empty; fill them (skeleton height, checkbox accent, item 9 real Discord messages).
6. Premium row crops (`discord-row-*`) show three rows with no panel edge; add the row label column so they read at a glance. Not blocking.

## What matches the spec well

- `discord-side-by-side-1440.png`: rail position 16 px under Plan (y 894 to 911), panel width 395 px, label row, destination line, `POST WHEN`, both rows, note, button sizes and the underlined link all line up with the mock within 2 px.
- 2000: page centred, content 376 to 1624, rail right column, panel 395 px, nothing stretched.
- Every string in 6.1 that appears is verbatim: destination, helpers, `Saved. Send a test message to see it in your channel.`, `Sent. Check the channel.`, `Discord said: Missing Permissions.`, `Wait a minute before sending another test.`, `Saved.`, `Removed.`, remove question, the three notices, the free sentence.
- Primary element per state: gold `Save webhook` only in the form and replace form; no filled button when saved; `Test and switch on` is the one gold fill when switched off; lapsed uses a 2 px gold bar and a text `See plans`.
- Switched off: 2 px `--wipe` bar, 14 px `--strong` sentence, 13 px muted last error, test helper, events after the replace link, as 3.2 specifies.
- Lapsed: boxes disabled, no note, no test, no replace; only `Remove`.
- Focus-visible: gold 2 px rings on the input, save, both checkboxes, test, remove, replace; hover borders and text on test and remove, fill on save.
- Invalid URL message sits under the input and does not repeat the typed value; the 402 line carries `See plans`; the removed state shows `Removed.` under the label row with the input focused.
- Premium row: all seven values render and link correctly; `Off` is muted and unlinked; `feature-off` renders no block and an `Off` row.
- Heights: saved 408 px at 360 and 390 (limit 420); paste form 532 px at 390 (limit 560). CLS after member action is 0 at 390 and 1440.
- 768 and 1024 stack and wrap as 3.1 and 7 describe; no horizontal overflow.

## Re-check (6f493595)

VERDICT: SHIP

Viewed again at the re-taken captures: the recomposed `discord-side-by-side-1440.png`, `discord-free-1440.png`, `discord-free-panel-390.png`, `discord-focus-see-plans-1440.png`, the saved, saving, sending, warning, warning-long, switched-off, switched-off-long, lapsed, remove-confirm and test-failed panel crops at 1440 and 390, and `discord-saved-1024.png`.

### Earlier findings, closed

- Must 1, free sentence: now two lines 20 px apart at 1440 (y 923 and 943), `See plans` ends line two, the focus ring hugs the link (about 60 x 20 px, no overprint). Closed.
- Must 2, pending states: `discord-saving-panel-1440.png` shows `SAVING…` on a disabled button with the field filled, the boxes greyed and the typed value kept. `discord-sending-panel-390.png` shows `SENDING…` with gold border and the boxes disabled. Real holds. Closed.
- Should 1, gap: buttons to `Use a different webhook` text is 52 px at 1440 (`discord-saved-panel-1440.png`: 321 to 373), the mock's 53 px. In `discord-test-failed-panel-390.png` the red status sits 8 px under the buttons with the link right below. Closed.
- Should 2, `Mana….`: gone; the long error ends `someone with…` at a word boundary, one ellipsis, no full stop. Closed.
- Should 3, warning weight: the first sentence is visibly brighter than the last-error line (`discord-warning-panel-1440.png`: `The last 2 posts failed. At 5 in a row it switches off.` in `--text`, `Last error from Discord: Unknown Webhook.` in `--muted`, 2 px apart). The capture agent's "both muted" is wrong; the two greys differ. Closed.
- Should 5, channel id: settled. The `1` in the id has a foot serif and a flag, and every digit is the same 7.7 px advance; the Barlow `1` in `Fixture#1234` (header) has neither. It is the monospace face, matches the mock's glyph shapes. Weight reads heavier than the surrounding 14 px Barlow only because mono at 13 px `--strong` is denser. Closed; no change.
- Copy: `A combat log you upload finishes parsing` is verbatim from the amended spec in the form, saved, remove-confirm, lapsed and replace crops.

### Rulings

1. Free sentence at two lines at 1440: accepted. The rail is 395 px (inner width 395, no panel padding) and the 82-character sentence at 13 px cannot fit one line; two lines of 20 px with the link ending line two is the right outcome. The spec's "44 px" is a minimum tap height, met by `min-h-11`. Amend §3.2 Free: "min height 44 px; two lines where the rail is narrower than the sentence".
2. `switched-off-long` at 390 (606 px crop, so about 558 px panel; 140-character three-line error): accepted as the worst case, no shorter cap. The typical switched-off panel at 390 is now about 500 px (`discord-switched-off-panel-390.png`, 548 px crop), inside the 520 px limit. Real Discord errors are one line (`Unknown Webhook`, `Missing Permissions`); capping the shown text below 140 characters would hide the one thing the member can act on. Amend §7: the 520 px limit holds for a one-line error; with a 140-character error the panel may reach 560 px at 390 and 590 px at 360.
3. `Saving…` and `Sending…` captures: confirmed real holds (see above). The blank band between the input and `In Discord:` while saving is the reserved 21 px error line plus its gaps, by spec §3.2 row 3 and §9 ("a result never moves the buttons"); it is about 40 px at 1440 and 390. Keep it. If it is trimmed later, remove only the gap above the reserved line, never the line.

### Remaining Must fix

None.

### Polish for a later pass

1. `discord-saved-1024.png`: the Premium row value still breaks as `On · 2` / `events` with the underline split. Join the value with non-breaking spaces.
2. `discord-remove-confirm-panel-1440.png`: `Use a different webhook` stays under an open remove question. Hide it while the question shows.
3. Native blue checkboxes: set `accent-color` to `--gold` on this row and the anonymize row together.
4. At 360 and 390 the form's gap between the input and the `In Discord:` helper (see ruling 3) could drop 8 px.
5. Record rulings 1 and 2 in the spec's §3.2 and §7 as written above, and fill the README's empty `Verdicts` and `Deferred, not this lane` sections.

## Follow-up lane re-check (a71bbb90)

VERDICT: SHIP. No must-fix items. Two cap amendments and one capture request below; none blocks the merge.

Viewed: paste-steps, paste-invalid, replace, destination (named, unnamed, long, markup), events (private on, logs-off/private-disabled, lapsed-all-disabled), warning one and four-long, switched-off (1440, 360 long), remove-confirm (saved and switched-off), lapsed, the Premium row crops at 390, 1024 and 1440 (all seven values), feature-off, the warm loading captures beside their `-ready` twins (all six states at 390 and 1440), cold at 1440, the checkbox captures (Discord focus and disabled, rankings checked and focus, sim) and the side-by-side. Only what is on disk was judged; the fixture artefacts the README lists are not re-raised.

### Ruling 1: the 1440 caps (amend, no cut)

The extra 8 px is `StatePanel`'s own `md:py-5` against `py-4` (spec 3 tokens), so the 7 px and 1 px overshoots are the padding, not the content. A cut would mean changing every panel on the page. Amend section 7, 1440 column, to: paste form 650 px (measured 647), switched off with a one-line error 565 px (561), saved 470 (measured 457, unchanged), plan lapsed 490 (478, unchanged), switched off with the 140-character error 620 (600, unchanged). Rule for the future: a cap at 768 and up is the phone cap plus 8 px unless the wider line length removes a wrapped line. The phone caps stand: form 639 of 640 at 390, switched off 553 of 580. One addition the spec lacked: the 1-to-4 warning with a 140-character error (`discord-warning-four-long-panel-390.png`, 584 px; 1440 573 px). Amend the saved-with-warning cap to 540 px with a one-line error and 590 px at 390 and 360 with the longest error, for the reason already ruled for switched off (the error is the one thing the member can act on).

### Ruling 2: the loading reserve (accepted)

Warm skeletons sit where the ready panels sit. At 1440 the rail blocks start at y 185 (Devices), 450 (You), 675 (Plan) and 910 (Discord, saved); the ready panels start at 185, 450, 675 and 911. The same holds in the form, off, lapsed, free and none montages. At 390 the Devices, You, Plan and Discord skeletons start at the ready panels' y within a pixel in all six states, and the left skeleton ends where the Characters and Your reports pair begins to be absorbed by the stored height. Page heights (full-page PNGs) are equal within 1 px for all twelve pairs (for example saved 2188 against 2187 at 390, form 1809 against 1809 at 1440). The footer top matches. Visible differences, all accepted by 9.1: the ready header's identity line and `Sign out` are absent from the skeleton (a blank band of the same height, so nothing moves), and the skeletons draw bars without panel chrome, as 9.1 specifies. Cold is acceptable: it is the recorded, ungated case (CLS 0 at 390, at most 0.006 at 1440, inside the 0.15 and 0.08 allowances), and the 656 and 663 px footer growth sits below the first screen, which is why CLS stays near zero. `discord-loading-cold-1440.png` shows the three-block rail and footer at y 987 as 9.1 describes.

### Ruling 3: the checkboxes (accepted; one capture requested, not blocking)

Checked: gold fill, dark tick (`discord-events-private-on-panel-390.png`, `rankings-checkbox-checked-1440.png`, `checkbox-sim-1440.png`). Disabled: muted grey box with a grey tick when checked, label `--muted` (`discord-events-logs-off-private-disabled-panel-1440.png`, `discord-events-lapsed-all-disabled-panel-1440.png`). Focus-visible: 2 px gold ring at 2 px offset on the box, not clipped (`discord-checkbox-focus-1440.png`, `rankings-checkbox-focus-1440.png`). Unchecked: the dark-scheme outline. The source rule in `global.css` is the one 9.2 asks for, with the disabled label colour set outside the layers, and no local `accent-*` remains on any checkbox (the two `accent-gold` hits are range sliders in `TimeChart`, not checkboxes). The account `anonymize` box is judged from that generality: the unchecked state is visible and correct in `discord-feature-off-1440.png`. Polish, not a gate: attach one checked `anonymize` capture and one disabled sim box the next time the capture spec runs, since the sim capture was force-enabled by a throwaway spec.

### Ruling 4: the changelog sentence (approved verbatim)

`New webhooks leave private logs off; a webhook saved before today keeps posting them until you untick Include private logs.` It is true to spec section 10 (the one-time `upload_private = true` for old webhooks), names the control exactly as the label reads, uses the date-relative `today` correctly for a same-day entry, and has one idea on each side of the semicolon. Keep it.

### Ruling 5: the amendments, item by item (all pass)

- Numbered steps: three `list-decimal` items, 13 px muted, markers muted, 4 px between items, in the paste form and the replace form; the intro is absent in the replace form (`discord-replace-panel-390.png`). Pass.
- Error line: input bottom to the `We store the URL` line is 21 px (README measurement; the crops agree). With the message present the error starts directly under the border in `--wipe` and the focused input shows the gold focus ring, not the `--wipe` border (the field keeps focus after a failed save), which is acceptable. The invalid-URL message wraps to two lines at 390 and 1440 (395 px rail), so it pushes the trust line down 21 px and touches it with no gap, which reads as one paragraph in two colours. Member-caused shift, allowed; polish: 4 px above the trust line when an error shows. Pass.
- Destination two lines: `Posts through “Raid pings”` over `Channel 1198…4568`, name `--strong` 600, id mono tabular; unnamed reads `Posts to a Discord channel`; the 40-character name wraps and ends in an ellipsis inside the quotes; `@everyone <b>x</b>` renders as text. Pass. The cut glyph looks like three dots in the crop and the README writes `...`; confirm it is U+2026 as spec 3.2 says (not verifiable from pixels).
- Private row: indented by the box plus gap (about 28 px), disabled and muted while the combat-log row is off with its value kept, the note shown in every state except lapsed (spec: lapsed has no note). Pass.
- Four-line notice: 2 px `--wipe` bar, 14 px `--strong` line, muted error, muted cure line, `Use a different webhook` inside, `Test and switch on` the only gold fill directly below. Pass. With the remove question open the link is absent (`discord-switched-off-remove-confirm-panel-390.png`) and also absent in the saved remove question (`discord-remove-confirm-panel-390.png`). Pass.
- Two-paragraph warning: `--text` first paragraph with `Send a test message to check.`, muted second paragraph, 2 px apart, no bar. Pass for one, two and four failures.
- Premium row: `Off · needs attention` fits one line at 390 and wraps at 1024; `Off · Premium` is muted and unlinked; `2 events` and `1 event` never split at 1024 (the wrap falls after the dot: `On ·` / `2 events`). Pass. Cosmetic: at 1024 `Off · no events chosen` and `Paused · plan lapsed` take three lines in the 90 px value column; accepted by 3.7, but a later pass could widen the value column at `lg` to 2fr. The 1024 crops start mid-row (the previous row's `only` shows); crop 4 px lower next time.
- Feature-off: no Discord block and no Discord row in Plan (`discord-feature-off-1440.png`). Pass.
- Side-by-side: destination, post-when rows, indent, note, buttons and link align with the mock within 2 px; the mock's gap above `Use a different webhook` is 53 px, shipped about 52 px.

### Must fix

None.

### Spec edits for the controller (I did not touch the spec file)

1. Section 7 caps as in ruling 1.
2. Section 3.2 and 7: saved with the warning, 590 px at 360 and 390 with the longest error.

### Polish

1. Capture a checked `anonymize` box and a real disabled sim box (ruling 3).
2. 4 px above the trust line when the invalid-URL message shows.
3. Confirm the truncation ellipsis is U+2026.
