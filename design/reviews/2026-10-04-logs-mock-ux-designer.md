# Logs landing mock review — ux-designer

Spec: `design/specs/2026-10-04-logs-landing.md`
Boards: `design/mocks/renders/logs-signed-in.png` (1440), `logs-signed-out.png` (1440),
`logs-phone-signed-in.png` (390)
Generators: `design/mocks/gen_logs.py`, `design/mocks/mocklib.py`, `design/mocks/data/logs-sample-report.json`

## Verdict: FIX

One blocker (garbled, overlapping text on every phone report row) and several must-fix
deviations from the spec's own verbatim copy and token rules. Nothing here requires a
new layout concept — all fixes are inside `gen_logs.py`.

## Region-by-region findings

1. **Blocker — `Your reports` rows collide on phone (390px), every row.**
   Element: `report_row()` in `gen_logs.py` (`grid-template-columns:1fr auto`).
   Spec said: §7, every region either stays one column or stacks cleanly; nothing specs
   a title running into its own facts line.
   What shipped: at 390px the title column is too narrow, so multi-word titles wrap to
   a second line and the mono facts text (same grid row, `align-items:baseline`) lands
   flush against that wrapped second line with no space — `Sanguine` / `Depths2026-09-15
   · 63 fights · 12 kills`, `Deadmines,` / `first` / `clear2026-09-28 · 24 fights · 6
   kills`. This is on all three rows of the signed-in phone board, confirmed by pixel
   crop. It reads as broken, not as design.
   Fix: on phone, stack the row as title (own line, can wrap to two lines) then the
   facts line below it (own line), i.e. drop the `auto` column below the `lg`-equivalent
   breakpoint and go `flex-direction:column` the way `ReportRow.svelte`'s real responsive
   rule must already do — match that, don't invent a new phone rule in the mock.

2. **Must fix — `Your reports`' sign-in button uses the wrong button token (signed-out board).**
   Element: `mine` panel's `btn("Sign in with Battle.net", "#signin", fill=True)`.
   Spec said: §4.A explicitly bars the filled gold Account button for anything but true
   sign-in on the signed-out home, quoting the design system: "Used only for 'Sign in
   with Battle.net' on the signed-out home... never for a tool or a link." The real
   component this panel wraps, `SignInPrompt.svelte` (non-compact), already renders this
   exact button with `SECONDARY_BUTTON_FIXED` (gold-bordered outline, `border-line-warm-
   strong`, `text-strong`) — not filled.
   What shipped: a fully filled gold button (`.btn-fill`), identical in weight to the
   home's one sign-in CTA, competing with the hero's own gold-bordered "Open this
   report" button for primary-action status on the same screen.
   Fix: render it as the outline Secondary button (border `#3a3326`/warm, text
   `--text-strong`, no fill), matching `SignInPrompt.svelte` line for line.

3. **Must fix — dungeon framing line is in two different places on phone vs. desktop.**
   Element: `logsFraming` string ("Logs are for group content at any level: a dungeon
   run logs the same way a raid does.").
   Spec said: §7 lists phone region order as Header → spine → Your reports → Recent
   public reports → Companion → Upload, and doesn't mention the framing line separately
   — implying it keeps whatever position it has on desktop.
   What shipped: on the 1440 board it sits directly under the spine, above the two-panel
   grid (first thing read after "who you are"). On the 390 board it's placed dead last,
   after the Upload panel, just above the footer — the opposite end of the page.
   Fix: move the phone framing line to sit directly under the spine, same relative
   position as desktop, and add it explicitly to §7's region list so this can't drift
   again.

4. **Must fix — per-player hook line doesn't match the spec's verbatim copy (signed-in hero, both boards).**
   Element: `hook` string in `gen_logs.py`, §4.A.1 / §6.
   Spec said (§6, verbatim): `{row.name} - {row.dps.toFixed(1)} DPS` then the two links,
   nothing else.
   What shipped differs three ways:
   - A lead-in that isn't in the spec anywhere: `Last kill, {fight.name}:`.
   - No hyphen between name and DPS — `REGLITCH` and `2,364 DPS` sit side by side on a
     flex gap with no `-`.
   - DPS formatted `2,364 DPS` (comma thousands, no decimal) instead of the spec's
     `toFixed(1)` form, `2364.0 DPS`.
   Fix: either add the lead-in to the spec's copy table and keep it, or strip it to match
   the spec exactly; restore the literal ` - ` separator; format the number with
   `.toFixed(1)`, no thousands comma.

5. **Must fix — §4.D.2's one new interaction, the pairing confirmation, is unreviewable.**
   Element: `pairing` block, companion panel.
   Spec said: §4.D.2 is the one genuinely new piece of UI this rebuild adds — the code
   block swaps in place for `{device.name} paired. It starts uploading as soon as you
   are logging.` once a device appears — and §11's acceptance list calls for all three
   pairing states (code shown, polling, success) plus both companion-status copy
   variants (`last seen ...` / `paired, not seen yet`).
   What shipped: all three boards show only the code-shown state and only the "last
   seen 4 min ago" status copy. The success line and the "paired, not seen yet" line
   appear nowhere, so the one new region this spec exists to add is not actually
   checkable against a rendered board.
   Fix: add a fourth board (or a cropped state swap on the existing signed-in board)
   showing the success line in place of the code, and confirm the "paired, not seen
   yet" status copy once, anywhere.

6. **Polish — header description's two anchor links use the wrong colour token.**
   Element: `desc` string's `<a>` tags (`#companion`, `#upload`).
   Spec said: §4.A, "the existing underlined anchor links... (`text-nav underline`,
   unchanged mechanism)" — `--text-nav` is `#b9b3a4`.
   What shipped: inline style sets the link colour to `BODY` (`#c9c2b2`), the same shade
   as the surrounding sentence, so only the underline marks them as links.
   Fix: colour the two anchors `#b9b3a4` to match `text-nav`.

7. **Polish — one caption under `Your reports` isn't in the spec's copy table.**
   Element: "Newest first. A guild tab shows the reports your guild shared with its
   members."
   Spec said: §6 lists every verbatim string for this page; this sentence isn't one of
   them, and §4.C doesn't call for a caption here.
   Fix: either add it to §6 as a new verbatim string (it's harmless, accurate copy) or
   drop it from the board so the mock doesn't ship copy the spec never approved.

## What's right, worth saying

- Hero selection, Sample pill placement/suppression, facts-line formatting, and the
  Secondary-button (not filled) choice for "Open this report" all match §4.A exactly.
- The guild tab strip (`Mine` / `{guild}`, gold border + `#e5b95514` fill when active,
  44px) matches `BandTabs`' own convention, reused correctly.
- `Recent public reports`' empty copy is the exact, correct, already-shipped string, and
  it is reachable (not displaced by the sample) on both signed-in and signed-out boards —
  §4.C.2's fix reads correctly.
- The band ships flat `--bg-raised` with no art and still reads as an intentional header,
  not an unfinished one — the gradient overlay carries enough atmosphere.
- Upload panel and companion pointer/status (seen variant) match their copy and chrome
  verbatim.

## Checklist (from spec §11, against what's provided)

- [x] Signed-out hero with Sample pill and static hook sentence
- [x] Signed-in hero with resolved per-player hook line (copy wrong, see #4)
- [ ] Signed-in, zero-reports state — not provided in this board set
- [x] Guild tab strip with an active tab
- [x] Companion status line, "last seen" variant
- [ ] Companion status line, "paired, not seen yet" variant — not provided
- [x] Pairing block, code-shown state
- [ ] Pairing block, polling state — not provided
- [ ] Pairing block, success state — not provided
- [x] Upload panel, idle/signed-in and signed-out states
- [ ] Upload panel, file-chosen / mid-upload / error states — not provided
- [x] 360–390px full page top to bottom (region order has a defect, see #3; row layout
      has a defect, see #1)
- [ ] Hover/focus-visible states — not verifiable from static renders
