# Subscribe-and-join gate fixes (reviews 2026-10-28)

Inputs: `design/reviews/2026-10-27-subscribe-and-join/review-ux-2026-10-28.md` (G1 to G7, L2, rows) and
`review-player-2026-10-28.md` (A to H, rows), on top of P1 to P5 (task-28-fixes-report.md). Built and
tested from an isolated archive of HEAD at `/tmp/sjgate` (`FOREVER_DATA=fixture`), previewed from a
copy of its dist on port 4646 and killed by PID. Source edits in the worktree; commits through a
temporary index (only these paths).

## Commits

| sha | what |
|---|---|
| 130a0008b | G7 fixtures |
| e68e35852 | G1 + G2 |
| a1f059459 | G3 + G4 + G5 (+ row 16 field 640) |
| 23da2a2be | G6, row 7, row 8, possessives |
| be39b2cbe | rows 1, 2, 11, 16 |
| 1bc989ee8 | L2, PV rows 19 and 21, caption date |
| f677362b3 | capture spec (rows 3, 9, 15, 16) |
| abe426684 | G2 follow-up: the load error keeps the reserve (lhci LCP) |
| (next) | regenerated captures + this report |

## Per item

- **G1** `GUILD_CARD_MIN` removed from `GuildStepCard`; real cards size to their content. New
  `GUILD_SKELETON_CARD = 'h-[260px] md:h-[156px]'`; `CHECKOUT_FRAME_PX.card = { phone: 260, md: 156 }`.
  Without the floor the owner's two `ask_leader` cards outgrow their skeletons and pushed the visible
  site footer down: CLS 0.0421 at 1440 (fail). Fix: while the step loads, `checkout.astro`'s head
  script sets `html[data-checkout-loading]` and the site footer is `visibility: hidden` (keeps its
  box, not drawn); it is first drawn at its final place once rows render, so nothing visible moves.
  (Tried and rejected: a loading reserve of at least the viewport with the footer drawn. Chromium
  counts the footer moving up into view for a short ready card: 0.0902 at 390, 0.0294 at 1440.)
- **G2** `releaseGuildFrameReserve` (checkout-layout.ts) removes `--checkout-guild-cards` and the
  loading mark as soon as rows or an empty state render (`CheckoutGuildStep` effect). The min-height
  is then void, so the page is as tall as its content. **Deviation: the load error keeps the
  reserve** (now `max(calc, 100svh)`, so the footer sits below the fold) and only shows the footer.
  Releasing it on the error made the footer the late LCP element on lhci's signed-out run (which
  reaches the error state: no API): checkout LCP 3039 ms, perf 0.94, against 2400 / 0.95. With the
  error keeping the reserve: 1502 ms, perf 1.0. Spec §7 amended to say so.
- **G3/G4/G5** `CopyLinkField` gains `labelVisible`, `onCard`, `phoneFullWidth` (defaults keep
  `GuildInviteLink`'s render). Checkout card: sr-only label, `bg-bg` field, full-width buttons below
  768 (grid, 2 columns with Share). Claim page member hand-off: sr-only label, full-width phone
  buttons, field capped at 640. Spec §5.5, §6.2.
- **G6** `leadSentence`: any `ready` -> ready lead; else any `claim`, `start_claim`, `confirm_claim`,
  `claim_link_bnet`, `claim_blocked` -> the some-claimed form; else the table. Spec §5.1.
- **G7** fixtures: the owner's characters are derived from the guild entries (rank and verify mark
  move together); Bow Jackzon (BM Hunter 24) in Employed Millennials; guilds `faction: 'horde'`
  (crest ring shows the Horde logo); claim holders/payers are verified guild masters; claimed guilds
  by someone else carry `verified_count` 9 and `has_approver`; `seats_total` follows the tier
  (Guild 1, Raid guild 3); MIXED's on_plan guild and the new `ON_PLAN_ALL_SEATS_GIVEN` are Raid
  guild; `ON_PLAN_SEAT` carries `me.seat` (no Get Premium foot line); any Premium source builds
  `premiumEntitlements` (rows 3 and 15 show On rows); founder `amount_cents` = `PLANS.premium.founderYearly * 100` (4000), with a panel test asserting `$40 a year` and no `$50`.
- **Row 1** band row `lg:items-start` (CHECKOUT top-aligns with the summary). Spec §5.1.
- **Row 2** plan=premium band hugs its h1: head script sets `html[data-checkout-band='premium']`;
  `CheckoutBand` CSS min-height 51 px / 103 px from 768 (`PREMIUM_BAND_PX`). Static and hydrated band
  are the same box. Spec §4.2.
- **Row 7** confirm_claim with no `holder_name`: sentence `An officer has started a claim on {Guild}.
  As an officer you can confirm it.`, confirm `Confirm the claim on {Guild}? The officer who started it
  will run the guild here and can subscribe it.`, done `Confirmed. The officer who started the claim
  holds {Guild}.` `claimantName` no longer falls back to the rank word. Spec §9.1 rows added.
- **Row 8** no verify sub-line on `claim_blocked` (the viewer is the guild master); a claimed guild
  with no approver reads `... or once the claim holder approves you.` Spec §5.4, §9.1.
- **Possessives** `lib/possessive.ts`: a name ending in s takes the apostrophe alone (`Employed
  Millennials'`). Applied across checkout-guild-copy and `guildClaimCopy.backToClaim`. Spec §9.
- **Row 11** empty panels sit left in the 880 column (max 640) with the foot; `Get set up` full width
  below 768; neither empty state draws the foot's missing-guild line (the premium line still shows
  for a `none` account). Spec §5.7.
- **Row 16** claim page plan line only while the viewer does not hold the claim; the subscribe
  button's failure line is drawn only after a failure, so `Next` keeps 12 px gaps; hand-off field
  max 640; capture uses the stored `OLYMPUS XXVII` and an unclaimed home for the no-approver note.
- **Rows 3/15 captures** crop `account-guilds-plan` (PLAN label through button); last-raid and
  reports reads stubbed empty.
- **L2** plan-card action slot and button both 44 px from 768 (button `align-self: flex-start`);
  premium.spec equal-height at 1024 holds.
- **PV row 19** /sim example box `md:max-w-[480px]`, left in the card. LP spec B3.4.
- **PV row 21** failed example image: `img` set `visibility: hidden` (inline `onerror` on /premium,
  state on /sim). LP spec B4 and row 21 amended.
- **Caption date** `splitCaptionNumbers` treats `d Mon yyyy` as one mono run on both cards.
- **L1** owner-only, skipped.

## Not done, by ruling or on the evidence

- **"No TBC raid names in mock guilds" (player row 16 (4)) not applied.** Hyjal Summit is a Forever
  first-tier raid in this repo's own data: `web/src/data/dates.json` ("Barrow Deeps · Hyjal Summit ·
  Onyxia") and the roadmap ("Dec 9: Barrow Deeps, Hyjal Summit, and Onyxia's Lair open"). Swapping
  it out would make the mock wrong against Forever's real raid schedule. Needs a controller ruling
  if the reviewer's reading should win.
- Optional player suggestions not taken: row 2 second link `Run a server sim`; row 10 price in the
  switch note; row 1's "move the size line below the lead at 360" (G4 already removes the 34 to
  51 px label that pushed `Copy link` past the fold).
- The mock guild's own `GUILD.name` stays `Olympus XXVII` (87 uses across other lanes' tests); the
  subscribe-and-join capture overrides it to the stored name instead.

## CLS, guild step loading to ready (checkout-guild.spec.ts, cap 0.02)

| scene | before (HEAD 84ee42648, 376/532 floor) | after |
|---|---|---|
| owner, two ask_leader, 390 | 0.0003 | 0.0003 |
| one ready, 390 | 0.0003 | 0.0003 |
| owner, two ask_leader, 1440 | 0.0002 | 0.0002 (0.0421 without the footer hold) |
| one ready, 1440 | 0.0002 | 0.0002 |

Premium value CLS (cls-premium-value-*.json): 1024 0.0003, unchanged band.

## Suites (isolated copy of HEAD abe426684, fixture build)

- `astro check`: 0 errors, 0 warnings. eslint on every touched directory: clean.
- vitest full: 7249 passed, 3 skipped, 1 failed: `src/fixtures/report/fixture.test.ts`, which reads
  `logs/engine/session/session.go`, absent from the archive (passes in the worktree; not this lane).
- Playwright full (desktop + mobile): 3202 passed, 228 skipped, 2 failed: `home-character-search`
  desktop and mobile (pre-existing; fails on main too, see task-28-fixes-report). The coach-lane
  failures of the previous run are gone.
- lhci full (22 URLs x 3): one failing assertion, planner performance 0.86 (pre-existing: 0.88
  last run, main 0.89). Checkout passes: LCP 1502 ms, perf 1.0 on all three runs.

## Captures regenerated (into the worktree)

- `design/reviews/2026-10-27-subscribe-and-join/`: the full row 1 to 17 set, 120 tests (row 17 run
  on port 4483, the port its main baseline was captured on: the rotated URL contains the port; 0
  differing pixels at 390 and 1440).
- `design/reviews/2026-10-27-launch-pricing/`: premium value rows 15 to 22, 25 tests, plus the
  `cls-premium-value-*.json`.
- Spot-checked: row 1 at 360 (`Copy link` ends about y 604, above 640) and 1440; row 2 band 104 px;
  row 3 crop shows the plan line with every row On; row 5 Guild master/Verified and `1 of 3` on a
  Raid guild; row 16 after-claim has no plan line and 12 px gaps; LP 15 buttons 44 px; LP 19 box
  capped at 480; LP 21 plain box, date in mono.
