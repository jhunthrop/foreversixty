# Tier list: DPS figure links its ranking sim (2026-11-13)

Scope: each DPS row's DPS figure on /tiers is an anchor to its ranking sim (`simHref`); the
spec name still opens BiS. Tank and healer rows keep a plain figure: the browser sim does not
produce damage taken or healing, so no run reproduces their numbers. /bis still links its
ranking sim ("Open in the sim ›", BisSpecPage.svelte) for the level 60 raid band of a damage spec.

Quick pass ahead of a Reddit post: ux-designer and wow-player only. ai-slop-reviewer and
info-density-reviewer are a follow-up pass after it is live.

## Round 1
- ux-designer: SHIP WITH ONE FIX. The figure has no at-rest cue, so it is undiscoverable on a
  phone. Add a quiet dotted underline, solid on hover and focus.
- wow-player: FIX. Same at-rest cue; "sim run" reads as a dungeon run, use "Open in the sim".

Fixes: `.ui-rr-fig-link` 1px dotted underline at 40% of the text colour, solid on hover and
focus; tooltip "Open in the sim: <Spec>", accessible name "<value> DPS, open <Spec> in the sim".

## Round 2
- ux-designer: SHIP. The dotted cue sits clear of the phone's gap line; quiet by intent.
- wow-player: SHIP. Reads as the Wowhead "more info" cue; optional nit: the dots could be a
  touch darker at 390.

Captures: tiers-1440*.png, tiers-390*.png (local dev build, dev toolbar hidden).

## Round 3: live pass (ai-slop-reviewer, info-density-reviewer)
Captures of https://foreversixty.gg/tiers, Damage tab, 2x scale: live-tiers-1440*.png,
live-tiers-390*.png (the full-page captures stayed in the scratch dir).

- ai-slop-reviewer: SHIP. One dotted underline is the only cue, no chip, icon or arrow; the
  tooltip names the spec and does not restate the number; tank and healer rows stay plain, so
  the cue marks where something happens. Yes to the phone nit.
- info-density-reviewer: SHIP. The figure exposes `simHref` we already build for every DPS row;
  no duplicate sim link renders in the row. Yes to the phone nit. Follow-ups, not blockers:
  the talent build or band in the figure tooltip, a sort control and Copy CSV, and a figure on
  tank and healer rows that opens something real (BiS or the guide), never the sim.

Fix (wow-player's optional nit, both agreed): at 767 px and below the rest underline sits at
60% of the text colour instead of 40%; hover and focus stay solid at full colour. Also moved
the stray "Below 1280" comment back above its media query. after-tiers-390-row-rest.png shows
the darker dots (live page with the rule injected).
