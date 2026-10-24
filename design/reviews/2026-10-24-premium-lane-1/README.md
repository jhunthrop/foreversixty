# Premium lane 1 review: the Supporter mark and the Premium block

Date: 2026-10-24. Plan: `docs/superpowers/plans/2026-10-24-premium-lane-1-page-truth-and-supporter-mark.md`.
Spec: `docs/superpowers/specs/2026-10-24-premium-experience-design.md` §4.1, §4.2, §4.5, §5.

Captures, each at 390, 1440 and 2000, taken from the fixture build with the API stubbed:

- `character-public-supporter-*`: a visitor's view of a supporter's public character page.
- `home-header-chip-*`: the signed-in home page, the shared "Your character" header.
- `account-premium-block-*`: the account page: hero chip and the Premium block under Plan.
- `premium-page-*`: `/premium` after the copy change.

## Verdicts

- ux-designer, first pass: FIX. The chip and the `/premium` edit could ship; the Premium block
  needed the plan line first, a two-column key/value layout, muted "ships" states, links on the
  live rows, the chip off the account character rows, and one chip border on every surface.
- wow-player, first pass: FIX. The chip needed a tooltip saying what it means; the `/premium`
  line should say "on your characters" (there are no profiles); the plan line should come
  first; the rest of the page read honest and worth the price.
- Both lists were fixed in one round (commit "the Premium block as a key/value panel with live
  links, one chip style, one ships month") and the captures here are from after that round.

## Deferred, not this lane

- The Monthly/Yearly toggle shows no price change in a static capture: it is the CSS `:target`
  toggle, covered by `premium.spec.ts`.
- The server-sims line "(5,000 combinations, any precision, no tab left open)" is developer
  talk; a copy pass on `/premium` can reword it.
- `/premium` has no page title above "What is free, forever".
- The home page's two empty bands at 1440 and 2000 and the 390 profile's gap under the ruleset
  chip predate this lane (island reserved heights in the fixture).
- "Supporter through <guild>" on the account page for guild-plan members, once `/v1/me` says
  which guild pays.
