# Web redesign rollout: close-out

The flat web system (`design/specs/2026-11-05-web-layout.md`) is on every player-facing page. Each page went mock first: an experience spec with a generator and mocks, a wow-player persona round recorded as the spec's section 0A, then the build, gated by a parity spec that renders the mocks at test time (`design/parity.md`). Main is at `ea09c458e` when this was written.

## Pages

| Page(s) | Spec | Parity spec | Main sha |
|---|---|---|---|
| Shared layout, header, footer, primitives | `design/specs/2026-11-05-web-layout.md` | `parity-*` (site chrome in every spec) | `420cacff3` (redesign 2) |
| Best in slot `/bis/[class]/[spec]` | `2026-11-05-web-layout.md` section 6 | `parity-bis.spec.ts` | `420cacff3`, final pass `5817a2107` (added parity) |
| Tier list `/tiers` | `2026-11-06-tiers.md` | none (pre-dates the parity gate) | `cc3ee6d2f` |
| Simulator `/sim`, results, saved | `2026-11-06-sim.md` | `parity-sim.spec.ts` | `103b0c576` |
| Home (A) `/` | `2026-11-06-home.md` | `parity-home.spec.ts` | `0672edbce` (added) |
| Home B: countdowns, feed, grades | `2026-11-07-home-b.md` | `parity-home.spec.ts` | `9776ce728` |
| Character `/character/...` | `2026-11-06-character.md` | `parity-character.spec.ts` | `102134418` (added) |
| Guild pages | `2026-11-05-guild-*.md` | none | guildx-1 `ca9f0c710`, -2 `be1cd0a40`, -3 `26abdb297`, -4 `70c4045b3`, raid route `556ed120f`, -5 `b0c3bab9f`, -6 `49357b28c`, -7 `c4e4a5de2` |
| Logs, reports list, rankings, ratings | `2026-11-07-logs-reports.md` | `parity-logs.spec.ts` | `e3a9f5d2c`, rankings LCP `8e9f4a519` |
| Setup `/setup`, `/connect` | `2026-11-07-setup.md` | `parity-setup.spec.ts`, `parity-connect.spec.ts` | `263536f9f` |
| Guides `/guides`, `/guides/[class]`, `/guides/[class]/[spec]` | `2026-11-08-guides.md` | `parity-guides.spec.ts` | `be1acd26a` |
| Talent planner `/planner` | `2026-11-08-planner.md` | `parity-planner.spec.ts` | `64577b1d2` |
| Premium, checkout, account, shared build `/b/:id`, build unavailable | `2026-11-08-premium-account.md` | `parity-premium.spec.ts`, `parity-account.spec.ts`, `parity-b.spec.ts` | `4fe3883bf` |
| Changelog, legal (terms, privacy, refunds), about, data sources, 404; `/roadmap` 301 to `/changelog` | `2026-11-08-pages.md` | `parity-pages.spec.ts` | `495b0cc08` |
| Follow-up: /account Lighthouse (0.84 to 0.96 signed in, CLS 0), Coming up rows expire, guides on `CharacterAnswer` | the two specs above | `parity-account`, `parity-guides`, `parity-pages` | `ea09c458e` |

"(added)" marks a sha that is the commit adding the page's parity spec, where no separate landing sha was recorded.

## What carried across rounds

- Rulings live in each spec as section 0A; the persona's factual flags were checked against the repo before they were applied (setup's `/fs job` flag was void).
- Mock facts readers bundle the page's own models and copy, and fall back to the fixture planner data when `src/data/generated` is absent (CI's parity job).
- Mock library classes are scoped: the step marks are `.smk` after a global `.sm` grew the BiS mocks.
- CI is in launch mode, so its build job runs only the guide prose test; each lane runs its touched unit tests itself.

## Open after the rollout

- The tier list and guild pages have no parity spec (a later follow-up).
- `/b/:id` is still rendered by the API's Go template; moving it to the web needs a Worker route.
- At 390 several parity specs carry expected drift for the head inset (the mock library draws the page inset outside the head region).
- `Coming up` season rows leave only when a changelog entry `covers` them.
- The two stat-weight questions from the guides round (Survival Strength over attack power, Fire crit under Intellect) are with the accuracy reviewer.
