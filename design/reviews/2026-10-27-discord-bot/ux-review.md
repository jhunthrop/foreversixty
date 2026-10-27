# UX review: Discord bot (Lane G5), 2026-10-27

Reviewer: ux-designer. Spec: design/specs/2026-10-27-discord-bot.md (with the 2026-10-27 amendment and the controller's §6.2 additions). Evidence: the PNGs in this folder, cls-390.json and cls-1440.json, and web/tests/captures/discord-bot/payloads.json (read to check what the Discord renders were drawn from).

## Verdict: BLOCK

The site surfaces are close: SHIP WITH FIXES on their own (findings 3 to 9). The Discord renders are what block. payloads.json was not produced by the shipped builders. In several places it contradicts the code in api/internal/discordbot (finding 1). And five of the twelve Discord rows have no render at all (finding 2). So the bot's messages, including every controller ruling, cannot be judged on screenshots. Tenet 6 asks for SHIP on screenshots, never on a diff.

## Findings

1. HIGH · discord-msg-readiness-*, discord-msg-bis-slot-*, discord-msg-press-select-*, discord-msg-attendance-*, discord-msg-press-unlinked-*, discord-msg-press-not-member-* · README says the payloads are the Go builders' output. They are not, and they disagree with the shipped code. Readiness has no `Readiness board` button, its footer is `Olympus XXVII` (spec §3.3.0 says `{Guild} · Forever Sixty`), and its rows read `Name: phrase`, not the board's ` · ` words. /bis has author `Zulmara`, not `Best in slot`. It has no thumbnail (tenet 2), no lead text line, no two alternatives, no `Share in channel` / `Open on Forever Sixty` row, and its footer is `Forever Sixty`, not `{Spec} {Class} · {band label} · Forever Sixty`. In the select, `In` is not bold and ` · your main` sits in the label instead of the description. Roster links to the raid page. `Link Discord` goes to /account#you, not the start route. `Get set up` goes to /connect, not /setup. cmd_readiness.go, cmd_bis.go, press_select.go, cmd_attendance.go and press.go each do these right · Fix: dump through the Service command and press handlers on the fakes_test.go stores, never through the inner embed helpers or hand-written JSON, then re-render every discord-msg PNG.
2. HIGH · (missing) · Spec §9 rows D4 (/raid, /signup), D5 (/char with upgrades, no gear, unknown), D7 (/report member, officer, shared), D9 (/readiness member, lapsed, Guild tier, nobody signed) and D10 (/link both states) have no render. Without them, none of the controller rulings shows on a screenshot: one unit per /char card, Stocked / Short: … / Not shared, nameless /raid, "No one has signed In or Late yet.", Move to Raid guild · Fix: render these through the same handler-backed dump.
3. MED · discord-bot-03-loading-* · `Save channel` draws at full gold while the list loads. Spec §3.1 and the mock's loading frame have it disabled at opacity-50 · Fix: pass `disabled` to the button while channels load, so the shared disabled style applies.
4. MED · discord-bot-08-unlink-question-* · `Cancel` is an underlined text link. The mock draws it as an outline button beside `Unlink`, the same as the remove question (§3.1). The `Linked to Discord as zulmara` line is gone, though the mock keeps it above the question. No focus ring shows on `Cancel` · Fix: make `Cancel` a SECONDARY_BUTTON_FIXED, keep the linked line, and move focus to `Cancel` with focus-visible.
5. MED · discord-bot-08-linked-*, discord-bot-08-return-linked-*, discord-bot-10-account-* · `Unlink` sits inline to the right of `Linked to Discord as zulmara`. The mock and §3.2 stack the line, then `Unlink`, then the status line · Fix: put the linked row in the same `flex flex-col items-start gap-2` as the unlinked row.
6. MED · discord-bot-09-installed-focus-bot-open-crop-390 · The focus outline of `Open in Discord` (44 px by padding) runs through line one, `#raid-signups in Olympus Discord` · Fix: draw the outline on the inline text box only, radius 2 and 2 px offset as in §4.1, and keep the hit area in a transparent pseudo-element.
7. LOW · discord-bot-04-just-saved-* · Spec §3.1 says focus goes to `Send test message` after a save. No ring shows (Chromium hides programmatic focus after a mouse press) · Fix: call `focus({ focusVisible: true })`, or capture the save from the keyboard so the ring shows.
8. LOW · discord-bot-08-*, discord-bot-10-account-* · The page is signed in as Fixture#1234, but the header selector reads `Sign in / or paste an export`. The settings captures show the fixture account · Fix: stub the header's session read in the account captures and recapture.
9. LOW · discord-bot-03-open-list-* · The native select is not drawn open, so the grouping by Discord category (optgroups, uncategorised first) has no screenshot · Fix: state the optgroup order in the README with the e2e assertion that proves it, or capture the list as the mock does.
10. MED · discord-msg-bis-slot-* · The item facts are invented fixture data: Quel'Serrar at item level 69, source `Dire Maul`, `+9.2 DPS` against a tank sword. Tenets 7 and 8 forbid that, and the spec says the /bis boards read the active build's BiS file · Fix: render /bis from the active build's bis and item tables (icon, quality, item level, game source).
11. LOW · discord-msg-post-done-*, discord-msg-post-ended-* · `When` reads `(in 5 hours)` on raids that are over · Fix: set the renderer's clock after the raid's end for the ended and done frames.

## Checked and matching the spec

- State 5: one fill, then the explanation, then the webhook path as a text link. The section order under Billing is right at 360, 390, 1440 and 2000, and centred in the 1344 column at 2000.
- State 4: the offer sentence and the outline add button sit between the destination and the actions; hover shows a gold border and gold-hover text.
- Picker: the install return focuses the select with a gold ring. Empty, error and race states use the copy verbatim, with `--wipe` only on the error and race lines.
- Installed: the destination line, the linked-members line, the 3-failure warning, three outline actions wrapping to two rows at 360, the fallback note, and the remove question with its webhook variant.
- Returns: denied, failed, taken and taken-expired show the right copy, colour and focus on the add control.
- Removed in Discord: both variants, the 2 px gold bar, the one fill.
- Read-only: the gate sentence, `aria-disabled` add and webhook links, disabled test and change, a live `Remove`, and the members line. Frozen: every control disabled, with the page's frozen sentence.
- Amendment 2b: deleted with and without a webhook, and permissions with `#raid-signups` strong and `Open in Discord`. The picker opens from the notice, nothing preselected when deleted, and `Cancel` instead of `Remove the bot`.
- Account: the unlinked lead; the suggested lead as the panel's one fill; conflict in `--wipe` with the muted help line; the returns.
- CLS: 0 and 0.0001 against the 0.02 cap, at 390 and 1440.
- Raid post (render): titles, colours, descriptions, field order, the two-line counts with the first clause bold, `In` as the only Success button, disabled buttons when locked and ended, `Raid page` / `Open the night` links, and no row on the webhook post. On the phone the post fits one screen, with `Raid page` on row two.
- Press sentences (render): In, Late with `Add a note`, Bench, Out, already, the move, rate limit, pending, refusals. All verbatim.

## Re-review 1

Evidence: commit 705c0093d, the site captures retaken, every discord-msg PNG rebuilt by the bot's real handlers on the seeded test database, and payloads.json (I checked the link URLs, embed URLs and thumbnails). Spec as amended on 2026-10-27: `Raid channel in`, `{n} signed In or Late.`, the R3 wording and the press day.

### Verdict: SHIP WITH FIXES

All eleven findings are fixed. The three new findings below are LOW. N3 goes to the wow-player review.

### The eleven findings

1. Fixed · discord-msg-readiness-officer, bis-slot, press-select-open, attendance-own, press-unlinked, press-not-member · The renders now come from the real handlers. Footers read `{Guild} · Forever Sixty` and, on /bis, the band footer. `Readiness board` and `Roster` go to the right tabs. `Link Discord` opens the start route and `Get set up` opens /setup. The pressed state is bold, and ` · your main` sits in the option description. /bis shows the lead line, `Best in slot`, two alternatives with their deltas, real icons and the Share and Open row.
2. Fixed · discord-msg-raid, raid-unlinked, signup, char-upgrades/no-gear/unknown, report-member/officer/shared/after-share, readiness-member/lapsed/guild-tier/nobody/unlinked, link-linked/unlinked · All present. Every controller ruling shows on a screenshot: one unit per /char card (DPS), `Stocked` / `Not shared`, `No one has signed In or Late yet.`, `Move to Raid guild`, the R3 wording and the press day.
3. Fixed · discord-bot-03-loading-crop-360 · `Save channel` is disabled at 50 %.
4. Fixed · discord-bot-08-unlink-question-crop-360 · The linked line stays, `Cancel` is an outline button, and the focus ring shows on it.
5. Fixed · discord-bot-08-linked-crop-1440 · The line, then `Unlink`, are stacked.
6. Fixed · discord-bot-09-installed-focus-bot-open-crop-390 · The outline is tight to the link and clear of line one.
7. Fixed · discord-bot-04-just-saved-crop-1440 · The focus ring shows on `Send test message`.
8. Fixed · discord-bot-08-unlinked-1440 · The header reads Fixture#1234.
9. Fixed · discord-bot-03-open-list-* plus the README · The README names the e2e assertion on optgroup order (web/tests/e2e/guild-discord-bot.spec.ts, `the picker › ready`). That evidence is accepted.
10. Fixed · discord-msg-bis-slot, bis-top, char-upgrades, report-officer · The items come from the real build 1.60.1.70630, with real icons, qualities, item levels and sources.
11. Fixed · discord-msg-post-ended (`4 hours ago`), post-done (`13 hours ago`) · Each frame has its own clock.

### New

N1. LOW · discord-msg-bis-slot, bis-top · Every /bis item title, `Open on Forever Sixty` and `All 17 slots` go to `#band-alliance-60`. Spec §3.3.3 asks for the BiS page's slot anchor on item titles, and the BiS page has no slot anchors · Fix: give each gear row a `slot-{key}` id and link the item titles to it, or amend §3.3.3 to the band anchor.
N2. LOW · discord-msg-report-officer, report-member · `4 items. The first 4:` repeats itself when every item is shown · Fix: drop ` The first {k}:` when k equals n, and amend §6.2 to match.
N3. LOW (for wow-player) · discord-msg-char-upgrades vs discord-msg-bis-top · The same items show different gains: Sentinel's Chain Leggings `+3.0 DPS` on /char against `+7.4 DPS` on /bis, and the Signet Ring `+2.4` against `+5.0`. The two answers read different bands (bare level 60 against level 60 Phase 1 raid), and the /char card does not say which band. An item-level-80 Anachronos ring in a Phase 1 raid band also needs the phase-gate check · Fix: wow-player confirms the band and phase. If it holds, /char names its band in `By DPS gained, against the guide build.`, or uses the /bis default band.
