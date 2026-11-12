# ux-designer, round 2

**Verdict: SHIP WITH FIXES.** All four round-1 must-fix items are closed (the druid crest at 32 px matching the heading's; no rotation block on the panel; the guild board wider with its cells inside and a level 60 undecided row; the footer line at the buttons' edge, symmetric with Done). The four-tile page holds at exactly 450. One new defect before merge.

## Must fix (new)

1. `panel-ask`: the question wraps to two lines and the split wraps inside its own phrase ("2 Holy · 2 / Retribution"). Put the split on its own muted line under the question at panel width.

## Should fix

- The dialog's bottom hints sit about 8 px above the border (round 1, neither fixed nor declined).
- No board shows a hovered tile or a focused Pick (round 1); name the pure test or add a capture.
- Done's right edge sits past the tile row's and the QR column's right edge.
- The chip card's button labels are left-aligned while the tour tiles centre them.
- Feral and Feral Bear tiles both say "Next: Ferocity"; the playstyle lines separate them, so correct the notes' claim.
- "Your path: Holy build" under a lead that says "the build you pick": accepted as the owner's wording; record it as the one exception.

## Declines accepted

"Tab to switch" and the companion line; "Pick a path" and "Your path"; no spec icons; the 50-point raid defaults, sign-up counts, others' splits and the dialog chrome (data or a later pass).

## The additions

The split reads right on the tour, the dialog heading and the chip card's foot; the next talent is muted everywhere and Following is the only gold among the rows; the four-tile page fits at 450 with about 50 px over the footer; the chip card has the client buttons and the crest on the pill; `tour-aside` shows the dialog opaque over the world with the tour hidden.
