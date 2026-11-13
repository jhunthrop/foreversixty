# wow-player review: addon dialogs fix (commit 8d7a719a1)

Verdict: SHIP WITH ONE CHANGELOG EDIT. No must-fix in the render itself.

## Reasons
- The bug a player would screenshot is gone. In-game 2.0.44 shows page 4 with a blank blue box bottom-left. In both pixel-grid captures (1440p and 1080p, footer crops included) the right-hand "fixed" side reads "Skip tour" in the same face and weight as Back and Next. A button with no label is the worst kind of addon bug, and it is fixed.
- Borders and button outlines on the fixed side are hairline-crisp at both scales. The old side has soft, doubled edges that match the "a bit blurry" complaint. Text edges are cleaner too.
- Layout matches the design mock: Skip tour left, Back and Next right, Next in teal as the primary.

## Caveats (not blockers)
- The captures show a sword icon, not the class crest. "Class crests at 128 px" is therefore not shown by these renders. Check one crest in game before claiming it.
- Nothing here shows the panel or the Builds window. The "Change" button is not rendered anywhere in this folder.
- The in-game screenshot differs from the mock on page 4 and this fix does not touch it. These are cheap follow-ups:
  - "Tonight 08:00 pm" has a leading zero. The mock says "8:00 pm".
  - "Job: Whelps, west cave" is plain white in game. The mock has it gold.
  - "In game and on the phone app coming" has "coming" jammed against the line with no gap or tag styling.
  - The right column rows sit a few pixels off the left column's baselines.

## Changelog line
Proposed: "Every button in the tour, the panel and the Builds window shows its label again (Skip tour, Change), and the addon's windows, borders and class crests draw sharp instead of soft."

- "Every button ... panel and the Builds window" claims more than the renders prove. Only Skip tour is visible, and the fix loads fonts at login, so the claim may be true. Confirm it in game before it ships.
- "Change" means nothing to a player reading patch notes. Name the window.
- "Draw sharp instead of soft" is fine, plain language.

Suggested: "Fixed buttons showing no label (Skip tour in the welcome tour, Change in the Builds window). Windows, borders and class crests now draw sharp instead of soft."
