# Companion screenshots: verdicts (2026-11-10)

Scope: /setup "The companion" step. `web/public/addon/companion-status.jpg` (1600x460) and
`companion-status-phone.jpg` (912x140) are now 2x crops of companion 0.2.44 (round-5 window, seal v4),
state `logging-credited` from `design/mocks/companion/shoot_ui.cjs`, with the fixture name Bow Jackzon.
The phone crop is now served under 640 px through `Capture.svelte`'s optional `narrow` source. Before this,
it was declared in copy.ts but nothing used it.

Captures (signed out, served from the built dist): `setup-1440-full.jpg`, `setup-1440-companion-step.jpg`,
`setup-390-full.jpg`, `setup-390-companion-step.jpg`. No other site page uses a companion image.

## ux-designer: SHIP (round 2)

- Round 1: SHIP. Shoulds: the 390 view scaled the full crop until its text was 6 to 7 px, the phone crop was
  unused, and a copy.ts comment was stale.
- Round 2: SHIP. At 390 the headline crop reads at about 12 px. At 1440 nothing changed. The Capture change
  holds up: the narrow source is optional, existing captures keep their inline aspect-ratio, and the
  `<picture>` source carries width and height. Should: the alt's trailing "as" was clumsy, so the alt now
  reads "…as they end, logging as Bow Jackzon." (applied).

## ai-slop-reviewer: PASS (round 2)

- Round 1: FAIL. At 390 the nav row was unreadable and the alt text was a pixel inventory.
- Round 2: PASS. At 390 the picture is just the Logging title and sub line, and it reads cleanly. At 1440
  it is the real window with no generated look. Minor and not blocking: the Example chip (pre-existing, from
  the spec) and the in-picture gold link, which looks clickable.
