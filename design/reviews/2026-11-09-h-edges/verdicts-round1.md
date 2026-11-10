# Round 1 verdicts (captures: round1/)

- ux-designer: VERDICT: FIX (blocker: the admin links 390 stacked-metric layout in GoLinksAdmin.svelte)
- wow-player: VERDICT: FIX
- ai-slop-reviewer: VERDICT: FAIL
- info-density-reviewer: VERDICT: THIN
- guild-officer-reviewer (profile, report): VERDICT: CLOSE

## Blockers taken into round 2
- Admin links at 390: the four metrics stack in one column (ux).
- Top guilds: an identical full bar and "1 boss" on every row measures nothing (wow-player, ai-slop 13, density).
- Report deaths: every killing blow 4,210 (1,050 overkill), Deep Breath at 0:30 (wow-player; fixture data).
- Report: who needs a word is not surfaced for officers; parses carry no comparison (guild-officer).
- Profile: no comparison on item level or the top parse; Raid ready has no roll-up (guild-officer, density).
- Consumables panel spends a panel on two identical no-data lines (wow-player, ai-slop, guild-officer).
- Qualifier and explainer text: admin heads and metas, profile notes said four times, the gain line (ai-slop 7, 21, 22).
- Signed-in landing still pitches Run a guild? to a guild officer (ux, wow-player).

## Declined, with reason
- Loot sources (wow-player): Eskhandar's Collar, Ring of Binding and Sapphiron Drape are Onyxia drops (data/builds/1.60.1.70790/raid-bosses.json lists Sapphiron Drape, 17078, on Onyxia's table).
- The framed inset line and the painted ground (ai-slop 8, 2, 12): the round-5 design system's own (docs/design-system.md: two or three framed panels a page); not a lane decision.
- 40-man sizes in the shared night fixture (wow-player): the fixture serves every guild parity state; out of this lane.
