# Server sims: "Your server runs" and "done while away": experience spec

Author: ux-designer. 2026-10-24. Lane 3 of the premium experience. Product spec: `docs/superpowers/specs/2026-10-24-premium-experience-design.md` section 4.3 (the promise, the two additions, the rulings), section 5 (the account page is the Premium console; the upsell sentence), section 6 (order). Mock: `design/mocks/gen_server_runs.py`, render `design/mocks/renders/server-runs-states.png` (1440 and 390: the sim page with every row state, the empty, loading and error states, the character band with the done line, the account row). References: `docs/tenets.md` 4, 9, 10, 11, 13, 14; `design/specs/2026-10-11-character-header.md` (the band's fixed anatomy, which this spec adds one footer row to); `design/specs/2026-10-10-home-signed-in-panel.md`.

Surfaces changed: `web/src/components/sim/SimView.svelte` (lines 596-640, the signed-in landing composition), a new `web/src/components/sim/ServerRuns.svelte` (built only from existing parts, section 3), `web/src/components/character/CharacterHeader.svelte` (one footer row), `web/src/components/account/AccountPremiumBlock.svelte` (one value), `web/src/lib/billing/copy.ts` and a new `web/src/lib/sim/server-runs-copy.ts` (new copy goes in a new small module: `sim/copy.ts` is far past the 800-line ceiling).

## 0. The experience in one paragraph

A payer who started a run on our servers and closed the tab comes back to a page that says what happened. On any page with the character band, one line sits at the bottom of the band: `Your Frost Mage sim finished: 1,204 DPS`, with `Open result →` and a dismiss. On `/sim`, a list called `Your server runs` is the first thing under the page heading: each run's spec, title, state and result, a running row saying `Running · takes about 3 minutes`, a failed row saying why in plain words with `Run again`. No queue position, no cancel, no percentage appear anywhere, because none exist.

## 1. The one question and the answer that must be visible

**"I started a sim on your servers and left. Is it done, and what did it say?"**

| | 360 / 390 px phone | 1280 / 1440 px desktop |
|---|---|---|
| Any page with the band, a run finished while away | The band's footer line, directly under the band's action button, inside the first screen (the stacked band ends near y 400 at 390). | The footer line inside the band, y about 230 at 1440. The DPS figure is in the sentence. |
| `/sim`, a payer with runs | `Server runs` card head (alignment 2026-10-29) and the newest rows. The newest row (state and result) starts above y 560 at 390 x 844 and is the first thing under the static `Simulator` heading and scope note. At 360 x 640 the first row's state line is inside the screen. | Heading and the newest three rows inside 900 px at 1440; the `Run sim` block follows directly. |
| `/sim`, a payer with no runs | One line of 64 px: the empty sentence. The `Run sim` block is under it, inside the first screen. | Same. |

The job is to answer "did it finish" and hand over the number. Opening the result is one tap. Everything else on the page (the run block, the character list, Your sims) keeps its job and its place.

## 2. The reference, and exactly what we borrow

**The in-game mailbox and its login notice.** A player who logs in sees `You have unread mail` and a minimap icon; opens the mailbox and sees one row per item, unread rows marked, each with who, what and when, and a control on the row (`Take`, `Return`). That is the same job: something happened while you were away; here is a row for it; the unread ones are marked until you look.

Borrowed:
- The login notice becomes the band's one footer line (a single sentence, dismissed by looking at the thing or by one click), not a modal, a toast or a badge.
- The mailbox's row grammar: one row per run, the unread mark (a gold dot, the mailbox's unread glow), the muted time on the right, one control on the row where there is one (`Run again`, the mailbox's `Return`).
- The unread mark clears when the member opens the item. Opening `/sim` does not clear it, as opening the mailbox window does not read the letters.

Also cited, the incumbent we beat: Raidbots' "My Sims". Not borrowed from it: a queue position and a percentage bar. Ours cannot honestly show either (section 8). Row grammar also follows our own `SimHistory.svelte` (title, API headline sentence, date; whole row is the link) so the two lists on one page read as siblings. Details!'s segment list is not the reference: it lists finished fights, not jobs with states.

## 3. Layout, region by region

Tokens (all existing, `web/src/styles/tokens.css`): `--color-bg #07090d`, `--color-raised #0d111a`, `--color-card-top #131824`, `--color-line #262e40`, `--color-line-soft #1c2230`, `--color-gold #e5b955`, `--color-gold-hover #f5d27a`, `--color-strong #f2eee4`, `--color-text #e9e4d8`, `--color-muted #9a9484`, `--color-wipe #f0736b` (the failure text, AA on raised). Finished is not green: the number is the answer. Type: Barlow for text, JetBrains Mono tabular for the DPS and ages, `section-title` 15 px for the list heading (as `SimHistory` does). Radius `rounded-panel` for the list, `rounded-control` for the button. No new colour, no new control.

Shared components used: `ClassCrestRing` (the one crest language; 32 px in the list), `EmptyState`, `LoadError` (Try again), `Skeleton`, `SECONDARY_BUTTON_FIXED` (the `Run again` button style, 44 px high on phone, `md:h-9` on desktop as `SimHistory`'s filter does), `formatAmount`, `specLabel` (`lib/sim/spec-label.ts`, `Frost Mage`), and the header spec's relative-date rule (relative inside 14 days, else `Nov 25`). `pill-sample` is NOT used (ruling 4).

### 3.1 Page order on `/sim`, signed-in payer (`SimView.svelte` lines 596-640)

Top to bottom: static `Simulator` h1 and scope sentence (unchanged, in the astro page); **`ServerRuns`** (new); `SimRunBlock`; `LandingState`; the rest unchanged (`SimHistory` last).

`ServerRuns` is rendered only when `me.entitlements.server_sims === true`. A free member never sees the list, the heading or an empty state: the existing page, byte for byte. A lapsed payer (flag false) is a free member here; their old runs stay in `Your sims`. The upsell sentence is not added to `/sim` by this lane (the page already carries `simCopy.premiumRequired` where the server lane is chosen).

### 3.2 `ServerRuns` region (section element, `data-testid="server-runs"`, id `server-runs`)

- Card head (alignment 2026-10-29 §5.6): compact `server` head, h2 `Server runs`. Nothing else on the row: no filter, no count, no refresh control (the list refreshes itself, 3.5).
- Panel: `border-line bg-raised rounded-panel` 1 px, the panel `SimHistory` draws. Rows separated by 1 px `--line-soft`. At most **5 rows**, newest first by `created_at`. If the API returned more than five rows, a 12 px muted footer line inside the panel: `Older finished runs are in Your sims below.`
- Margins: `mx-[18px] md:mx-0`, gap 12 px under the heading, 24 px under the panel (before `SimRunBlock`), as the sibling sections.

### 3.3 One row (anatomy)

Desktop (768 and up): CSS grid `32px minmax(0,1.4fr) minmax(0,1.2fr) auto`, gap 14 px, padding 10 x 14, min height 64 px. Phone (below 768): `32px minmax(0,1fr)`, the state block drops under the title in column 2, padding 12 x 14.

| Cell | Content | Style |
|---|---|---|
| Crest | The spec's class crest, `ClassCrestRing` 32 px, class-colour ring. | One crest language. Decorative (`alt=""`). |
| Title (primary) | The member's title, else `specLabel(spec)` (`Frost Mage`), with the unread dot before it when unseen. | 15 px 600 `--strong`, single line, truncates with the whole title in `title`. Unread dot: 8 px circle `--gold`, 8 px gap, sr-only text `Not opened yet`. |
| Meta (secondary) | `{specLabel} · {Kind}` when the run is titled, else `{Kind}` alone. Kind words are `simCopy.kindLabel` (`Sim`, `Gear`, `Talents`, `Drops`, `Weights`). | 13 px `--muted`. |
| State block | Two lines, right-aligned on desktop, left-aligned on phone. Line 1 is the state or the result; line 2 is the time. | Line 1 14 px; line 2 12 px `--muted`, JetBrains Mono for the age. |
| Action | `Run again` on failed rows only. A muted `›` (16 px) on rows that are links. | Button: `SECONDARY_BUTTON_FIXED`, min width 112. |

Row states:

| State | Line 1 | Line 2 | Link? |
|---|---|---|---|
| `queued` | gold dot (static), `Starting` | `Started just now` | No |
| `running`, estimate present, not overdue | gold dot (pulsing), `Running` in `--strong`, then ` · takes about 3 minutes` in `--muted` | `Started 2 minutes ago` | No |
| `running`, estimate present, elapsed more than 1.5 times the estimate | gold dot, `Running` then ` · taking longer than usual` in `--muted` | `Started 9 minutes ago` | No |
| `running`, no estimate (fewer than three finished runs for the spec) | gold dot, `Running` alone | `Started 2 minutes ago` | No |
| `done` | The API headline in `--strong`, JetBrains Mono 15 px (`1,204 DPS`; `+41 DPS from Vis'kag`; `3 upgrades on Ragnaros`), falling back to `{dps} DPS` as `headlineOf` already does | `Finished 14 minutes ago` | Yes: whole row to `/sim/<id>` |
| `done`, unseen | As done, plus the unread dot on the title and the row ground `color-mix(--gold 8%, --raised)` | as done | Yes |
| `error` | `Failed: stopped before completion` (the API's plain words after `Failed: `), `--wipe`; with no reason sent, `Failed` alone | `Ended 2 hours ago` | No; the `Run again` button is the control |

The pulse is a 1.6 s opacity animation on the dot only, off under `prefers-reduced-motion` (the dot is then static). It is not a progress indicator and no bar, ring or percentage exists.

Estimate words: `estimate_seconds` rounded to the nearest minute; 1 minute or less reads `about a minute`; 2 to 89 minutes read `about {n} minutes`; 90 minutes and over read `about {n} hours` (rounded).

### 3.4 The done line in the character band (`CharacterHeader.svelte`)

A footer row inside the band's gold-bordered panel, rendered only when `/v1/me` carries one or more `finished_sims` the member has not dismissed. The panel becomes a column: the existing three-column grid on top (unchanged anatomy, padding and the 136 px minimum), then the footer.

| Property | Value |
|---|---|
| Position | Last child of the panel, full width, edge to edge inside the 1 px gold border; 1 px `--line-soft` top border; bottom corners follow the panel radius (5 px inside the border). |
| Ground | `color-mix(in srgb, --color-gold 8%, --color-raised)`. The only tint in the band; the band's filled gold button stays the one filled gold element. |
| Height | Desktop: 44 px, one line, padding 0 22 px. Phone (below 600): min height 56 px, up to two lines, padding 6 x 14 px. |
| Grid | `minmax(0,1fr) auto`: the link (sentence), the dismiss. |
| Link | One anchor to `/sim/<newest sim_id>` containing the sentence (13 px `--text`; the DPS figure in JetBrains Mono `--strong`), then on desktop ` Open result →` (13 px 700 `--gold`; the arrow never wraps). On phone `Open result` is dropped and only a gold `→` follows the sentence (the whole text is the 44 px target). |
| Spec name | `specLabel(spec)` (`Frost Mage`); the member's own title is in the anchor's `title` only. |
| More than one unseen | On desktop after the link: ` and 2 more`, a muted underlined link to `/sim#server-runs`, 44 px tall. On phone it is a second line under the sentence, 12 px muted underline, 44 px tall; the footer is then 80 px. |
| Dismiss | 44 x 44 button at the right edge (negative margin so the glyph sits on the 22 px padding), an `×` glyph 16 px `--muted`. `aria-label` `Dismiss: your Frost Mage sim finished`. |

States of the interactive parts:

| Element | Default | Hover | Focus-visible | Active | Disabled | Loading | Empty | Error |
|---|---|---|---|---|---|---|---|---|
| Line link | as above | the sentence underlines, 4 px offset; `Open result →` goes `--gold-hover` | 2 px gold outline, inset by 2 px (`outline-offset: -2px`), radius 4 | `--gold-hover`, no underline change | n/a | n/a | line absent | n/a |
| `and N more` | muted, underlined | `--strong` | 2 px gold outline | `--strong` | n/a | n/a | absent with one unseen | n/a |
| Dismiss | `×` in `--muted` on the tint | `×` `--strong`, 32 px circle ground 6 % white | 2 px gold outline, inset by 2 px (`outline-offset: -2px`, so the panel never clips it) | circle ground 12 % | while its seen call is in flight: `aria-busy`, the line stays until the call answers | n/a | n/a | the call fails: the line stays, and its text is replaced for 4 s by `Could not dismiss that. Try again.`, then returns |
| Whole line | one sentence | n/a | n/a | n/a | n/a | absent while `/v1/me` is pending (the hint reserves the space, 3.7) | absent when `finished_sims` is empty | `/v1/me` failed: the shared signed-out band shows, no line |

Behaviour: Dismiss posts the seen call (section 10) for **every run the line names** (the newest and the `N more`), removes the line in place, and the panel returns to its ready height. The line does not move keyboard focus when it appears; on dismiss, focus goes to the band's action button. Following the link also marks that run seen (3.6); the other unseen runs show on the next load. The line is announced once (`role="status"`, polite) when it first appears.

### 3.5 Where the line lives on `/sim` (decision)

`/sim` has no `CharacterHeader`. **It carries no separate line: the list is the line.** Unseen finished rows carry the gold dot and the tinted ground (3.3), directly under the page heading. A second sentence above the list would say the same thing twice. Rows leave the unseen state when opened (the result page marks it seen) or when they finish while this tab is visible: the page re-reads the list every 10 s while any row is `queued` or `running` and the tab is visible; a row that flips to `done` during a visible tab is marked seen at once (the member watched it finish), so no dot and no band line follows them to the next page. If the tab was hidden when it flipped, it stays unseen.

### 3.6 Opening a result

The existing saved-sim page (`/sim/` plus the id) marks that sim seen on mount, once, silently (no UI, failures ignored: the line simply stays and the member can dismiss it).

### 3.7 No layout shift when the list and the line appear

**The list.** `ServerRuns` sits above the run block, so a late list would push the run button. It reserves its box before data: heading 24 px plus the panel's last rendered height in px, which the list stores in `localStorage` (`fs.serverRuns.reserve`) after every paint; on a first visit, or when the stored value is missing or absurd, the fallback is one row (64 px desktop, 100 px phone). (Amended 2026-10-24 after the build: rows measure 99.5 px done, 118 px running and about 161 px failed on a phone, so a count × fixed-height reserve could not be exact; the stored height is.) The reserved box holds skeleton rows while loading. The empty state is one panel at the same reserve so it fits with no shift; a reserve from a longer list that resolves to fewer rows collapses once, after the member's own state changed since the last visit, which is accepted. The reserve is applied only under the existing `html[data-session='1']` hint plus a second pre-paint hint `data-server-sims='1'` set from the cached `/v1/me` snapshot's `entitlements.server_sims`, so a free member or a signed-out visitor reserves nothing and the static HTML of `/sim` is unchanged for them.

**The band line.** The band is hydrated from a cached `/v1/me` and its ready height is fixed. The line adds a row, so it is reserved by hint: the pre-paint script sets `html[data-finished='1']` when the cached snapshot held at least one `finished_sims` entry **or** the flag `fs.serverRuns.live` is set. That flag is written when a server run is started from this browser, or when the list reads a `queued` or `running` row, and expires 2 hours after it was last set. While the hint is on, the band's skeleton and its ready state reserve the footer row (44 px desktop, 56 px phone). If the hinted reserve resolves to no line, the reserved row collapses inside the band; that is one shift, only for a member who started a run in the last two hours and has nothing finished yet. The `and N more` phone layout (80 px) is not reserved: it grows once, inside the same paint that fills it. The one unavoidable case is a run that finished while away on a browser whose hint expired or never existed: the line appears with the band's own fill, a single 44 px shift. Lighthouse runs signed out, so neither case is in `lighthouserc.json`; the Playwright check in section 9 asserts CLS at most 0.02 for the hinted case at 1440 and 390.

**The rows.** State changes after hydration (running to done, done to seen) swap content inside a row whose min height is fixed. The one change that alters a row's height is a failure reason wrapping to a second line on a phone: a live update the member is watching, accepted, never at load.

## 4. States (every interactive element)

| Element | Default | Hover | Focus-visible | Active | Disabled | Loading | Empty | Error |
|---|---|---|---|---|---|---|---|---|
| Done row (link, whole row) | as 3.3 | ground `--card-top`, title underlines at 4 px offset, `›` goes `--gold` | 2 px gold outline, `outline-offset: -2px` (inset: the row is flush with its panel) | ground `#171d2b` | n/a | n/a | n/a | n/a |
| Running / starting / failed row | not a link; no hover ground, no pointer cursor | none | not a stop; `Run again` is | none | n/a | n/a | n/a | n/a |
| `Run again` (failed row) | outline button, `--text` on `--raised`, 1 px `--line-warm-strong` | border `--gold`, text `--gold-hover` | 2 px gold outline, 2 px offset | ground `--card-top` | after success: label `Started`, `--muted`, until reload | after press: label `Starting…`, `aria-busy`, disabled, until the API answers | n/a | the start failed: a 13 px `--wipe` line under the row's state block, `Could not start it again. Try again.`, the button returns to rest; a 402 (plan lapsed) reads `Server sims are part of Premium.` with the upsell sentence's `See plans` secondary link |
| List (whole) | up to 5 rows | n/a | n/a | n/a | n/a | the reserved box with three `Skeleton` rows (title 180 x 15, meta 120 x 12, state 140 x 14), `aria-busy`, no text | payer with no runs: `EmptyState`, no action | the read failed: `LoadError` with `Try again` (re-reads in place); the run block and everything below render as usual |
| Unread dot | gold dot | n/a | n/a | n/a | n/a | n/a | n/a | n/a |
| Band line, dismiss, `and N more` | section 3.4 | | | | | | | |

Touch and keyboard parity: every control is 44 px high or tall on phone (`Run again` 44 px, the line 44 px+, the dismiss 44 x 44, `and N more` 44 px); the list rows are 64 px or more. Tab order in the band: identity link, action button, line link, `and N more`, dismiss. Enter and Space activate buttons, Enter the links. Escape does nothing.

## 5. Hierarchy

- **Band:** the band's answer and its filled gold button stay primary. The line is secondary, but it is the only tinted row, so a member who came back for it finds it at a glance. Its one gold element is `Open result →`; the DPS figure is `--strong` mono. Nothing else on the line is coloured.
- **`/sim`:** the primary element for a payer with a run in flight or unread is the list; with none it is the `Run sim` block, so the empty state is one quiet line and the block follows at once. Per row: the title is primary (15 px 600), the state block is the answer (14 px), the meta and ages are secondary (12 and 13 px muted). Gold appears once per unread or running row (the dot) and on focus and the line's link.
- Failure is the only red on the list, as text, never a red ground.

## 6. Copy, verbatim

New strings live in `web/src/lib/sim/server-runs-copy.ts`; the account strings in `billing/copy.ts`. Functions take the data and return the string; no number is typed into prose.

| Where | String |
|---|---|
| List heading | `Server runs` |
| Empty (payer, no runs) | `No server runs yet. Choose Run on our servers when you start a sim, and it keeps going after you close this tab.` One `EmptyState`, no action. `Run on our servers` is `simCopy.runOnServers`, not retyped. |
| Loading | none visible (skeleton); sr-only `Reading your server runs…` |
| List error | `Your server runs did not load.` + `Try again` (`uiCopy.retry`) |
| Overflow footer | `Older finished runs are in Your sims below.` |
| `queued` | `Starting` / `Started just now` |
| `running`, estimate | `Running` + ` · takes about 3 minutes` / `Started 2 minutes ago`. Variants: `takes about a minute`, `takes about 2 hours`. |
| `running`, overdue | `Running` + ` · taking longer than usual` / `Started 9 minutes ago` |
| `running`, no estimate | `Running` / `Started 2 minutes ago` |
| `done` | the API headline / `Finished 14 minutes ago` |
| `error` | `Failed: ` + the API's `failure_reason` verbatim (`Failed: stopped before completion`, `Failed: the job could not be started`), or `Failed` alone / `Ended 2 hours ago` |
| Run again | `Run again`; pending `Starting…`; success `Started`; failure `Could not start it again. Try again.`; plan lapsed `Server sims are part of Premium.` + `See plans` |
| Unread, sr-only | `Not opened yet` |
| Band line | `Your Frost Mage sim finished: 1,204 DPS` + `Open result →` (desktop) or `→` (phone); with others `and 2 more` |
| Band line, dismiss | visible `×`; `aria-label` `Dismiss: your Frost Mage sim finished`; failure `Could not dismiss that. Try again.` |
| Account row, on | `On · 4 runs this month` |
| Account row, one | `On · 1 run this month` |
| Account row, none | `On · no runs yet this month` |
| Account row, no count sent (older `/v1/me`) | `On` |
| Account row, off | `In your browser` (amended 2026-10-27, launch pricing persona round, player 13; was `Browser lane only`) |

Numbers: `toLocaleString('en-US')` (1,204; 1,000 runs). Ages: relative inside 14 days (`14 minutes ago`, `2 hours ago`, `3 days ago`), else `Nov 25`; the absolute stamp is in each time's `title`.

## 7. Phone and desktop

**Never collapses, at any width:** each row's state and its result or reason; `Run again`; the band line's sentence and its DPS; the dismiss; the unread dot; the age on each row.

**Collapses:** below 768 the row's state block drops under the title (title, meta, state with age), `Run again` goes on its own full-width 44 px line under a failed row; the meta line may truncate before the title does. Below 600 the band line's `Open result` word is dropped, leaving the arrow, and `and N more` takes its own line. Desktop holds one line.

- 360: row padding 12 x 14, crest 32 px, title truncates with an ellipsis; a failure reason wraps to two lines (never truncates).
- 1024 to 1279: the desktop grid; no overflow at 1024 (the longest state line, `Failed: the job could not be started`, wraps to two lines at 14 px inside the 1.2fr column).
- 1440 and **2000**: the list is as wide as its siblings in the centred 1344 column (`SimHistory`'s width), never wider; at 2000 the container is centred and the band is the same 1344 band the header spec fixes. Columns do not stretch past the grid.
- Account row at 390: the key column is 10 rem, so `On · 4 runs this month` (22 characters, 14 px) fits the value column (about 190 px) without wrapping; at 360 it may wrap to two lines, and the underline wraps with it.

## 8. Rulings this spec makes

1. **No queue position, no cancel, no percentage, anywhere.** Runs start at once. A `queued` row is `Starting`, lasts seconds, and has no copy about waiting. The pulse on the running dot is not progress.
2. **The estimate is a duration, not a countdown.** It is the median of the last ten finished server runs for that spec, so `takes about 3 minutes` is honest at minute 1 and at minute 2. A countdown (`about 1 minute left`) would be wrong on a long run. Past 1.5 times the estimate the row says `taking longer than usual` rather than keep a number that has failed.
3. **Failed rows keep their row, their plain-words reason and `Run again`.** No error code, no stack, no job id.
4. **No kind pill on the row** (a difference from `SimHistory`'s row grammar). The class crest is the leading mark and the kind is a word in the meta line. Reason: a crest plus a pill plus a title is three leading marks on a 360 px row. `SimHistory` may adopt the crest in its own pass; this spec does not touch it.
5. **Running and failed rows are not links.** A running sim has no result to open. (Contingent: if the saved-sim page already renders a running sim, the API lane says so and the row can link.)
6. **The `/sim` list is the done line there**; no separate sentence is drawn on `/sim` (3.5).

## 9. Performance, polish, and acceptance

Limits: `ServerRuns` is lazy-loaded the way `SimHistory` is (`createLazyComponent`), and only for payers; no new font, no image beyond the existing crest sprite; the band line adds no network request (it reads `/v1/me`, which the band already fetches). Lighthouse (`web/lighthouserc.json`): `/sim.html` and every signed-out page are unchanged by construction (a free or signed-out visitor renders none of this), so the existing budgets hold: CLS at most 0.05, TBT at most 100 ms, accessibility and performance at least 0.95 (0.9 where the file says so). No missing-icon flash: the crest ring is the `ClassCrestRing` already used by the run block. No font swap.

Tests the lane adds: the e2e in the product spec (queue two runs, finish one, reload, see the line and the list states) with the API stubbed to the contract below; a Playwright CLS check of `/sim` and `/` at 390 and 1440 with the hints on; a unit test per row state's strings (estimate rounding, the overdue rule, a missing `failure_reason`); an axe pass on the list and on the band with the line. The product spec's e2e mentions cancelling one run: there is no cancel (section 8), so that step becomes "fail one run".

**Acceptance screenshots the implementer attaches** (fixture build, API stubbed; 390, 1440 and 2000 unless stated):

1. `/sim` payer: one running row with an estimate, one running without, one done unseen, one done seen, one failed with `Run again`.
2. `/sim` payer, empty (no runs), run block directly under it.
3. `/sim` payer, loading (skeleton box) and error (`Try again`), 1440 and 390.
4. Failed row: `Run again` rest, hover (1440), focus-visible, pending, started, and the could-not-start line.
5. Done row: hover and focus-visible at 1440.
6. Home signed in with the done line: one run, then three (`and 2 more`), then dismissed (band back at its ready height), at 390, 1440 and 2000. BiS, Rankings and Logs at 1440 with the line.
7. The band at 390 with focus on the line link and on the dismiss.
8. Account page Premium block: `On · 4 runs this month`, `1 run`, `no runs yet`, no count sent, and off, at 390 and 1440.
9. The CLS number printed by the Playwright check, hinted case, 390 and 1440.
10. Same-scale side-by-side of each shipped surface with the mock's matching panel, at 1440, with 1024 and 2000 captures.

## 10. API contract this spec assumes

The API lane builds exactly this; if it differs, change the contract here first.

- **List:** `GET /v1/sims?mine=1&lane=server` returns the first page, newest first, in the same page shape as the saved-sim list (`SimListPage`). Payers only (`entitlements.server_sims === true`); for others the web never calls it.
- **Row fields:** `sim_id`, `spec` (e.g. `mage-frost`), `kind` (`run|gear|talents|drops|weights`), `title` (may be empty), `state` (`queued|running|done|error`), `created_at`; for done rows the API headline sentence (`1,204 DPS`, `+41 DPS from Vis'kag`, `3 upgrades on Ragnaros`) and `dps`; **new:** `finished_at` for done and error rows, `failure_reason` (plain words) for error rows, `estimate_seconds` for running rows (median of the last ten finished server runs of that spec; absent with fewer than three), and, for the unread mark, **`seen: boolean`** on done rows (not in the lane brief; decision 10).
- **No queue and no cancel:** a run starts the moment `POST /v1/sims/run` is accepted. `queued` lasts a few seconds until the first progress tick flips it to `running`. There is no position field and no cancel route.
- **Done while away:** `GET /v1/me` carries `finished_sims: [{sim_id, spec, dps, title, finished_at}]`, the finished runs the member has not seen, at most 3, newest first. A run is seen when the member opens its result page or dismisses the line; both call `POST /v1/sims/{id}/seen`. The band is fed from the cached `/v1/me` on every page that mounts `CharacterHeader`.
- **Account row:** `GET /v1/me` carries `server_sims: {runs_this_month}` (a non-negative integer; calendar month; the API lane states the time zone, the copy says `this month`).
- **Run again (assumption the API lane must confirm or amend):** a failed run can be started again by id, `POST /v1/sims/{id}/rerun`, which returns the new `sim_id`. The list does not carry the run's request, so the web cannot re-post `POST /v1/sims/run` itself.

## 11. Decisions the product spec left open (for the controller to rule)

1. **Where the list sits on `/sim`: above `SimRunBlock`**, first under the static heading (the product spec: "the first thing a returning payer sees above the landing state"). Cost: the `Run sim` block moves down one list's height for a payer with runs; the empty state is one 64 px row to keep it small. Alternative: between the run block and `LandingState`, which keeps the button fixed. Recommended: above.
2. **No separate line on `/sim`; the list's unread rows are the line** (3.5). A finish watched in a visible tab is marked seen at once.
3. **Dismiss marks every run the line names as seen**, including `N more`. Alternative: one at a time. Chosen because three dismisses to clear a line is a chore and the line already told them.
4. **Estimate wording `takes about 3 minutes`** (a duration) rather than the bare `about 3 minutes`, and `taking longer than usual` past 1.5 times the estimate.
5. **A pre-paint hint to reserve the band line** (`data-finished`), fed by the cached snapshot or the `fs.serverRuns.live` flag (2-hour expiry). One 44 px shift remains for a run that finished on a browser with no hint.
6. **Count words:** `On · 4 runs this month`, `On · 1 run this month`, `On · no runs yet this month`.
7. **The list shows five rows, not a page.** Older finished runs are in `Your sims`; older failed runs are shown nowhere. Alternative: ten rows.
8. **`Run again` needs a rerun route** (section 10) and, after success, leaves the old failed row with `Started` on its button; the new run is a new top row. Alternative: the API marks the old row superseded and the list hides it.
9. **Done runs appear twice on `/sim`** (in `Your server runs` and in `Your sims`), because both are the saved-sim set. Alternative: `SimHistory` hides `lane=server` done rows; that changes a shipped list and is not done here.
10. **Unread uses a new `seen` field on list rows** (section 10), beyond the fields the lane brief listed.
11. **The product spec's acceptance step "cancels one"** cannot exist (no cancel); the e2e fails one run instead.
