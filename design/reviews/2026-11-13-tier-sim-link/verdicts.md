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
