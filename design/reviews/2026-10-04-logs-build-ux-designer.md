# /logs build review — ux-designer (2026-10-04)

Spec: `design/specs/2026-10-04-logs-landing.md` (incl. §12 amendments). Lane report:
`.worktrees/web-logs-rebuild/design/reviews/2026-10-04-logs-build-report.md`. Reviewed on the
build's own captures, `design/mocks/renders/logs-build/*.png` (33 files), against the approved
boards. Method: visual read of every capture plus pixel measurement (`PIL`) of gutters, band
full-bleed, control heights, and button/hero states, and a source read of
`web/src/pages/logs.astro`, `Panel.astro`, `lib/planner/styles.ts`, `MyReports.svelte` to confirm
what a screenshot alone can't settle (hover CSS, 44px tokens).

## Ruling on the two flagged differences from the boards

1. **Panel.astro's chrome (section title outside the panel) vs. the board's title-inside-panel
   drawing.** ACCEPTABLE, ship as built. `Panel.astro` (read directly) is the one, unmodified,
   already-shipped component — `section-title` sits above a `bg-raised border` box on every other
   page on the site, and the spec itself names `Panel.astro`'s chrome explicitly four times
   (§2, §4.C, §4.D) as the thing to reuse. The board's inside-panel title was an illustration;
   the build followed the real component per tenet 10 ("reuse the tokens... never invent a
   control"). No defect.
2. **`CurrentCharacterBar` (`spine`) unchanged vs. the board's own drawing of it.** ACCEPTABLE,
   ship as built. §4.B is explicit: "kept, unchanged... mounted the same way." The real
   component's abbreviated door labels and the absence of a visible "Forget" at this width are
   that component's own existing, already-shipped behavior, not something this spec touched. The
   board's spine row was an unlabeled illustration, not literal output. No defect.

## Findings

### 1. BLOCKER — the two most important 360px captures do not show the states they're named for
`signed-in-reports-360.png` (§11 state 2, signed-in-has-reports, at the spec's own "narrowest the
site designs for" viewport, §1) shows the **sample** hero ("Sanguine Depths, sample log," with the
`Sample` pill and the static hook sentence) instead of the visitor's own newest report — even
though the same capture's `Your reports` panel correctly lists the real "Molten Core, week 3" row
beneath it. `signed-in-zero-360.png` (§11 state 3) shows the mirror fault: its hero **and** its
`Your reports` list both show the real "Molten Core, week 3" report, so it never demonstrates the
zero-reports state (sample hero + "No reports yet. Upload a log or run the desktop companion.") at
all. Confirmed correct at every other width (390, 1024, 1280, 1440, 1920) for both states — this
is isolated to 360px. Whether the root cause is a capture-script fixture race or a genuine
viewport-triggered bug in the hero-selection fetch, the delivered artifact fails §11.2/§11.3's
explicit requirement and violates §1's own no-scroll promise (the visitor's own report must be the
hero at 360px when they have one). **Fix:** re-run both captures cleanly at 360px and, if the
wrong hero still renders, debug the hero-selection fetch for a 360px-specific regression (the
8 other widths all resolve correctly, so this looks capture-side, but it must be proven, not
assumed).

### 2. BLOCKER (carried forward, not introduced here) — sample report roster still not Classic-valid
Unchanged from the spec's own filed defect (§4.C.3/§10 finding 5): the canonical sample report
(`5lop7n5kwysf`) still carries a Mistweaver roster row, not a Classic Forever class
(`forever-class-rules`). The build lane filed this correctly and did not fake it away, per the
task's own instruction — but the spec's own words make this "a blocking fix needed before this is
a reference example." It is still open. Not a regression from this build; still gates ship of the
sample content itself.

### 3. MUST FIX — "paired, not seen yet" companion-status capture is a duplicate, not the state it claims
`companion-status-never-seen.png` and `companion-status-seen.png` are byte-identical
(confirmed via checksum) — both read "MacBook Pro · macOS · last seen 4 minutes ago." The required
§11.5 "both states" deliverable never actually shows "paired, not seen yet." The underlying
feature does work — `signed-in-zero-390.png`'s companion panel correctly shows "Gaming PC ·
Windows · paired, not seen yet" — so this reads as a capture-script mistake (wrong fixture
device cropped), not a component bug. **Fix:** re-capture `companion-status-never-seen.png`
against a device fixture with `last_seen_at: null`.

### 4. MUST FIX — "Open this report" has no visible hover state
`hover-open-this-report.png` is pixel-identical (background, border, and text colour all match
within anti-aliasing noise) to the button's default render in `signed-in-reports-1280.png`.
Confirmed in source: `SECONDARY_BUTTON_FIXED` (`web/src/lib/planner/styles.ts`) carries no
`hover:` class, and `LogsHero.svelte` adds none on top of it (`{SECONDARY_BUTTON_FIXED}
border-line-warm-strong text-gold w-fit px-4`, no hover override). Tenet 9 requires hover/focus
parity on every interactive element; the focus-visible capture on the guild pill does show a real
focus ring (`focus-visible-guild-pill.png`, a visible gold offset outline around "Sample Guild"),
so the gap is specific to hover on this button. **Fix:** add a hover affordance (e.g.
`hover:border-gold` or a background tint) to the button, and consider adding it to
`SECONDARY_BUTTON_FIXED` itself so every caller of the shared token benefits, not just this page.

## Verified correct (worth recording, not re-litigating)

- **Gutters and full-bleed band**, pixel-measured: band background reaches both edges (x=0 and
  x=viewport-1) at 360, 390, 1024, 1440 and 1920px. Content sits on the sitewide 48px-desktop /
  18px-phone gutter — the 96px measured at 1440px from the true browser edge to content is the
  same `max-w-[1344px] mx-auto` (auto-margin) plus the container's own `px-12` (48px) padding
  every other page on the site uses verbatim (`character.astro`, `planner.astro`,
  `account.astro`, etc., grep-confirmed); not a regression specific to this page.
- **44px controls**: the guild tab pills are `h-11` (44px) with `border-gold` +
  `background: #e5b95514` when active, `border-line` otherwise — exact match to spec and to
  `BandTabs`' own active-pill convention (confirmed in `MyReports.svelte` source).
  `SECONDARY_BUTTON_FIXED` is `h-11` on every breakpoint, confirmed in source.
- **No CLS between the header's skeleton and its ready/failed states**: band bottom edge lands
  at the identical pixel row (y=365 at 1280px, y=625 at 360px) in both the skeleton and the ready
  capture — matches the build report's own CLS-fix claim.
- **Hook line, format and colour**: `Last kill, Lucifron: Zulmara · 1,250 DPS · Gear in the
  planner | Sim` — matches the §12-amended format exactly, thousands-separator at ≥1,000 DPS
  confirmed, `Zulmara` renders in the Hunter class colour (~#AAD372, matches `#ABD473` within
  anti-aliasing).
- **Copy, verbatim**: header description, facts line, empty states (`Your reports`, `Recent
  public reports`, guild-empty), pairing-success line, and the signed-in/signed-out sign-in
  prompts all match §6's table exactly, including the §12 ruling that sign-in actions use the
  outline Secondary button, never filled gold (confirmed on `signed-out-1920.png` and
  `signed-in-zero-1920.png`).
- **Pairing flow**: code-shown and polling captures are intentionally identical (the design
  defines no visible polling indicator); the success line replaces the code block in place with
  the exact copy, `{device.name} paired. It starts uploading as soon as you are logging.`

## Checklist (from spec §11)

- [x] Signed out, 360/390/1024/1280/1440/1920 — sample hero, sign-in prompt, honest empty recent
  reports
- [ ] Signed in w/ reports, 360 — **fails, finding 1**
- [x] Signed in w/ reports, 390/1024/1280/1440/1920
- [ ] Signed in, zero reports, 360 — **fails, finding 1**
- [x] Signed in, zero reports, 390/1024/1280/1440/1920
- [x] Guild tab active, 390/1440
- [ ] Companion status, both states — **fails, finding 3 ("paired, not seen yet" not actually shown)**
- [x] Pairing: code shown / polling / success
- [x] Upload: idle / chosen / progress / error
- [x] Hero skeleton, 1280 and 360 — reserved height matches ready height, no CLS
- [~] Hover / focus-visible — focus-visible pill passes; **hover on Open this report fails,
  finding 4**
- [x] 360px full page top to bottom, region order matches §7

## Verdict: FIX

1. Re-capture (and if the underlying bug is real, fix) `signed-in-reports-360.png` and
   `signed-in-zero-360.png` so the hero and the `Your reports` list agree with each other and
   with the state the filename claims — blocker.
2. Resolve or re-point `data/sample-report.json`'s canonical sample so its roster carries only
   Classic Forever class/spec rows — blocker, carried from the spec's own filed defect, still
   open.
3. Re-capture `companion-status-never-seen.png` against a null-`last_seen_at` device so it
   actually shows "paired, not seen yet" — must fix.
4. Add a visible hover affordance to the `Open this report` button (and consider
   `SECONDARY_BUTTON_FIXED` itself) — must fix.

---

## Fix round 1 re-check (2026-10-04)

Commits `57ff52ea`, `4bad371c`. Re-checked items 1, 3, 4 on the re-shot files in
`design/mocks/renders/logs-build/` (filenames unchanged, content/mtime updated 12:19). Item 2
(sample roster) stays filed for the owner, unchanged.

1. **Fixed.** `signed-in-reports-360.png` now shows the signed-in hero ("Molten Core, week 3"
   with the resolved hook line, "Last kill, Lucifron: Zulmara · 1,250 DPS · Gear in the planner |
   Sim") agreeing with its own `Your reports` list. `signed-in-zero-360.png` now shows the sample
   hero ("Sanguine Depths, sample log") with `Your reports`' own empty state ("No reports yet.
   Upload a log or run the desktop companion.") and the "Gaming PC · Windows · paired, not seen
   yet" companion line — both captures now internally consistent and correct for their named
   state. Capture-script root cause (shared browser context carrying over the previous state's
   cache) matches what was observed.
2. **Fixed.** `companion-status-never-seen.png` (checksum now `78ca7bd2...`, previously identical
   to `companion-status-seen.png`'s `e905291f...`) now reads "Gaming PC · Windows · paired, not
   seen yet," distinct from `-seen.png`'s "MacBook Pro · macOS · last seen 4 minutes ago." Both
   states of §11.5 are now demonstrated.
3. **Fixed.** `hover-open-this-report.png` now shows a visibly brighter/bolder gold border and
   text versus the default render (pixel-sampled: dominant gold fill brightened from
   `(229,185,85)`/border `(74,64,48)` pre-fix to `(245,210,122)`/border `(105,94,64)` post-fix,
   with the bright-pixel count in the button box roughly tripling) — a real, visible hover state
   now exists, matching the described gold-border-and-text, 120ms treatment.
   `focus-visible-guild-pill.png` is unchanged and still correct (gold offset ring on "Sample
   Guild").
4. **Side-by-side unchanged.** `_sbs.png`'s mtime is untouched by this round (still the original
   12:03 capture); spot-checking `signed-in-reports-1440.png` (regenerated 12:19 alongside the
   fix round) against the pre-fix version shows identical content region-by-region — no
   incidental regression from the fix round.

### Updated checklist

- [x] Signed in w/ reports, 360 — now passes
- [x] Signed in, zero reports, 360 — now passes
- [x] Companion status, both states — now passes
- [x] Hover (Open this report) / focus-visible (guild pill) — both now pass

Only open item: finding 2, the sample report's Mistweaver roster row, filed for the owner, not a
build-lane defect.

## Verdict: SHIP

Ship with the sample-roster data defect (finding 2) tracked separately as a content fix, not a
layout/build blocker — the coordinator's own routing of it to the owner is correct and it was
never this build's to fix.
