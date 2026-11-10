# Round 1: ux-designer (captures in round-1/)

The build is faithful to round 5 in the main:

- **Shape:** 10 px panels and 6 px buttons.
- **Type:** Marcellus headlines on Open Sans body.
- **The frame:** a cream 1 px inset with flourishes on the one lead panel.
- **Primary action:** a teal Sign in.
- **Focus:** one gold ring.
- **Tray:** the seal v3 tray tones and the tray menu match the inventory.

## Blockers

1. **Min size has no capture.** There is no capture at 560 x 480, and the nav may not fit. Attach status (logging), device (waiting) and settings-unsaved at 560 x 480.

## Must fix

2. **Status loses its answer block at 800 x 620.** Game folders is cut off by the footer in every populated state.
   - The mock puts the fight count as the big figure top right of the lead panel. Do that and drop the Fights uploaded row.
   - Put "Check for builds now" inline. It takes its own line and makes the Addon row about 100 px tall.
3. **The update-ready bar.**
   - It covers the folder row.
   - Its text sits off the page gutter.
   - Its 16 px yellow text reads as a link.
   - It has no Dismiss.

   Use the tray-notice pattern: left at the gutter, cream text, a Dismiss ghost button.
4. **The held upload is not shown.** `uploading-held` reads only "Uploading 3 fights...". Add the held line in amber and a Queued row.
5. **The offline bar.**
   - It is off the gutter and wraps to two lines.
   - "Check for builds now" in the dimmed table looks clickable.

## Polish

6. **Band.** A hard 1 px rule under the band, where the mock melts into the ground: soften it or drop the rule. The h2 page titles on the ground are not in the mock.
7. **Destructive links.** "Forget this device" and "Remove" use the harmless gold ghost. Use the wipe colour.
8. **Bars.** The tray-notice bar is centred and the offline bar is left-aligned. Pick one, left-aligned at the gutter.
9. **Settings strip.**
   - The Save strip merges with the footer. Add a 1 px line-warm top rule.
   - "No changes" contrast.
10. **Copy link.** Seven solid blue "Copy link" buttons shout over the zone names. Use a ghost with a hairline.
11. **Menu-bar dot.** At 16 pt the dot overlaps the seal face. Move it to the corner.
12. **Nav.** The nav's right edge is 2 px off the gutter.

## Checklist

- [x] Round-5 corners, frame, fonts and teal primary
- [x] Focus ring
- [x] Tray and app icons per OS
- [ ] Min-size capture
- [ ] Status answer visible without scrolling
- [ ] Bars aligned
- [ ] Held-upload state
- [ ] Destructive colour, band edge, strip rule

VERDICT: SHIP WITH FIXES — the round-5 look is faithful, but there is no 560 x 480 capture, Status loses its Game folders section at 800 x 620, and the held-upload state is not shown, so recapture after fixing items 1 to 5.
