# Character history review: the History panel, the free owner's sentence, the account row

Date: 2026-10-25. Plan: `docs/superpowers/plans/2026-10-25-character-history.md`.
Experience spec: `design/specs/2026-10-24-character-history.md` (ux-designer; amended during the lane:
`difficulty` is the integer id, rings mark private reports only). Product spec:
`docs/superpowers/specs/2026-10-24-premium-experience-design.md` §4.4, §5.
Mock: `design/mocks/gen_character_history.py` → `design/mocks/renders/character-history-states.png` (git-ignored).

Captures taken from the fixture build with the API stubbed, at head cca3bc36 (110 files, `history-<state>-<width>.png`,
section crops as `history-<state>-section-<width>.png`):

- Premium owner at 390/1024/1440/2000: `owner`, `stepped` (rows marked after stepping back), `single-kill`,
  `show-all`, `loading`, `error`, `skeleton-hinted` (the page skeleton with the reserved box).
- Interaction states at 1024/1440/2000: `hover` (a viewport clip: the readout previews 1240 Rare with the warm
  guide), `focus-plot`, `focus-step`, `focus-pill`; `picker-hover-1440`, `picker-fade-390` (eight bosses),
  `step-disabled-1440`, `step-hover-1440`.
- Gaps at 390/1440: `gap-label` (in-band reason), `gap-note` (the narrow gap's under-chart note).
- Edge states at 390/1440: `empty`, `one-point`, `no-kills`, `ring-note-1440` (`Rings are private reports.`).
- Free owner, visitor, anonymised at every width: `free-owner`, `free-owner-focus`, `visitor`, `anonymised`.
- The account page's History row at 390/1440: `account-row-{entitled-main,entitled-nomain,off}`.

Measured CLS, hinted second visit: 0 at 390, 0.00276 at 1440; the spec's limit is 0.02.

## Verdicts

- ux-designer: SHIP WITH FIXES on the first pass (boss names hard-clipped beside the Private/Guild note; the
  loading skeleton left a void inside the reserve; the free-owner sentence orphaned `See plans` at 390; the hover
  captures showed no preview; serif sub-headings; the one-kill label misaligned). Fixed in commit cca3bc36 (the
  hover was a capture artefact) and re-captured; re-check: **SHIP** (two polish notes: 4 px above `See plans` at 390; a stronger picker hover border). Report: `ux-designer-review.md`.
- wow-player: SHIP WITH FIXES (13 items). One was real and shipped: the parse chart plotted the rating engine's
  output percentile while the row beneath showed the fight's Parse %, two numbers for one kill — the chart now
  plots the fight's own percentile (commit a6a0510f). The boss-name clipping was fixed with the ux-designer's
  item. The rest were fixture artefacts (two fixtures' numbers, a priest with dps rows, wipes in the score list),
  by spec (the y window around the data, no year inside one year, the bare free-owner sentence, the narrow gap's
  note) or later passes (an explanation of Execution and Parse %, an explicit "open fight" affordance, tier-coloured
  parse points, a locked-chart preview for free owners). Report: `wow-player-review.md`.

## Deferred, not this lane

- A one-line explanation of `Execution` and what `Parse %` is ranked against, site-wide (the fight table has none).
- An explicit action on a bad fight (open the fight, sim, logs) beyond the row link.
- Parse points coloured by percentile tier (the numbers are; the line is gold).
- The free owner's upsell as a locked-chart preview rather than one sentence.
- The skeleton reserve sits at the bottom of the skeleton while the real section sits between the rating panel
  and Gear; the fixture measures no shift, a taller rating panel may.
- `scoreDomain` and the tier-band labels at 390 (`Uncommon` touches the plot edge).
