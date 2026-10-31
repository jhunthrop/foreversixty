# UX verdict: remove the Jump to bar

Verdict: SHIP. No blockers.

Reviewed on rendered captures only: sim, sim-gear, planner at 1440 and 390, plus the selector popover (1440) and phone sheet (390).

## Bar removal
- /sim 1440: tab underline rule, then about 33 px, then the current-character card. No empty band, no doubled rule.
- /sim 390: tab row, about 20 px, then the card. Spacing matches the gap between the other stacked cards. No orphan.
- /sim/gear 1440 and 390: tab rule, then the "The simulator runs damage specs only" note at normal paragraph spacing. Clean.
- Planner 1440 and 390: the header block runs straight into the page; no leftover strip above the title or between header and body.
- The header selector is the only switch control on every page. The one-way-to-switch rule holds.

## Restored line and Forget
- Placement: a full-width row between the character list and "Paste an export", with a hairline above and below. It reads as a footnote to the list. It does not compete with the character rows or the Paste action.
- Style: the muted 11 px "Restored your last character" on the left, and "Forget" at the right end, underlined in the same treatment as "Manage". It matches the panel's existing quiet-link style at both widths.
- 1440: row is about 45 px tall, aligned to the panel's 18 px side padding; Forget's right edge lines up with Manage above it.
- 390: Forget sits on the right edge of the sheet and is readable. It is clear of the sheet's close button, which only affects the header row.
- Copy is verbatim to the request. Visible label "Forget" is contained in the accessible name "Forget {name}", so label-in-name is satisfied.

## Blockers
None.

## Shoulds
1. Forget text is 11 px muted on the phone sheet. If the 44 px button already makes the target large (as the change states), consider raising the line to 12 px so it is not the smallest text in the sheet. Polish, not required.
2. Not verifiable from stills: the 44 px hit height, focus-visible ring, and focus returning to the trigger after Forget. Attach a focus-visible capture of Forget and confirm in the existing test that focus lands on the trigger.
3. The line does not say what happens next (header falls back to the main). Acceptable as a quiet menu line. If players ask, add a title or aria-description "Switches back to your main"; no visible copy.
4. No before captures were supplied, so "no gap" is judged against the surrounding spacing rhythm, not a diff.

## Out of scope, noted only
- Empty tool areas, "That sim did not load.", and the planner's "Talent data did not load" are from the missing API stub.
- The sim tab row clips "Weights" and "Spec support" at 390 (scrolls); not from this change.
