# Home body, round 4: final (2026-11-09)

Final captures: `r4-1440.png`, `r4-2000.png`, `r4-2576.png`, `r4-390.png`. Before: `before-1440.png`, `before-2000.png`, `before-2576.png`, `before-390.png`.

## Final verdicts

- ai-slop-reviewer: `VERDICT: PASS`
- ux-designer: `VERDICT: SHIP`
- wow-player: `VERDICT: SHIP`
- info-density-reviewer: `VERDICT: DENSE WITH ADDITIONS`

Blocking findings: none from any reviewer.

## Logged non-blocking follow-ups

### ai-slop-reviewer
- **Panels look alike.** The four panels share one treatment. A reviewer suggested taking the panel off a small list, but class-coloured lists stay on a plate by owner rule.
- **"of you" on every row.** The Guild sims gaps repeat it.
- **Source wording.** "Excavation Site: Wetlands: Relic Guardian" is the BiS file's own source wording.
- **"is the new ranged".** These are the API's words, from `bisdiff.go`.
- **Fixture rank moves.** Two "climbed 6 places to #3" fixture rows.
- **Phone head line.** It breaks after a dangling separator. The head band belongs to the other lane.

### ux-designer
- **Uneven panel bottoms.** The Guild sims and Your spec this week bottoms differ by about 10 px.
- **Selected-row band.** It reaches the panel's rounded edge.
- **Bar distance.** At 2000 and up the bar sits far from the name.
- **Phone "over" line.** It wraps on long names.
- **Tier tile.** It has no key.
- **"Climbing this week" spacing.** It sits close to the line above it.

### wow-player
- **Plaque title.** "First raids open" beside a countdown could read "open in".
- **Upgrade rows.** They show no level requirement or drop chance; the BiS payload carries neither.
- **Source detail.** Quest and reputation sources don't say which they are.
- **Feed pinning.** Pin the viewer's own class and band rows.
- **Raid row states.** Mixed states on the raid rows.
- **Accuracy check.** Golemsight Long Gun, the Relic Guardian source, Ironweaver, Cartilage Shapers and Armor Piercer.

### info-density-reviewer
- **Guild sims.** It has no drill-down links, no header link, and no spec on the rows.
- **Movers.** They show places but no figure.
- **Phone "over" truncation.** Truncate the "over" line on the phone.
- **Feed filter.** A Mine / Guild / All filter.
- **Needs data.** A per-character sim history, and drop rates.

## The stray top-right element ("...found.")

Not reproduced, and not fixed.
- The owner's capture was 2x, so the viewport was about 1288 CSS px. It showed the bottom 12 px of a dark rounded box at x 1025 to 1288, above the bar.
- No fixture build reproduces it (signed in at 34 in a guild, at 1440, 2000 and 2576).
- Live production signed out doesn't show it either.
- No site or API copy ends in "found." (grep of web/src, apps/shared/copy and api/internal).
- Ruled out:
  - the selector's restored note, its live region and its messages;
  - the header search status and typeahead ("No character named ... yet");
  - the deleted-account notice;
  - the Founder tag;
  - the item tooltip host, which home does not mount.
- A browser extension's overlay, such as a password manager's inline menu, fits the shape and the wording.
- Next step: the owner's DevTools element picker on the box when it shows again.
