# Lane B report (site-alignment leftovers)

1. Home "Your upgrades" void: the island stores its settled card height (`fs.home.upgrades.reserve`); an inline script in index.astro sets it on the slot before first paint. A first visit holds 147 px, the measured height of the smallest real state (was 220). home-panel pins no-void (short card) and the stored-reserve round trip; the signed-in CLS test still shows no late shifts at 390 and 1440.
2. /setup: Done / To do `Tag` in each step head (SetupStepTag, same hints and reads as the chips; global.css rules extended to `.setup-step-tag`). The character line uses CurrentCharacterBar with its `inCard` prop (added by lane A, uncommitted at the time), not a fork.
3. Report: muted "Not credited" Tag beside the unchanged credit sentence; "Show N trash fights" is a full-width row button with a chevron.
4. Beta gear wording: `lib/items/unknown-item-copy.ts` owns "Not in our item data yet". Home upgrades (plus "Item <id>" under it), the gear panel, and the planner and sim strips ("Item <id> · Not in our item data yet", via their copy modules only) read it.

Tests: vitest touched files green; Playwright (fixture build on 4672, desktop and mobile) home-panel, home, handoffs-setup, onboarding, character-gear/render/phone, one-character-journey, report-*, sample-report, coach-report, setup, current-character, logs-recent: 656 passed.
Captures: b-home, b-setup, b-report at 1440 and 390 in this folder; generator design/mocks/2026-10-30-leftovers.
