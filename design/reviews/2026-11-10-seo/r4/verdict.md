# Round 4: /rankings at 390 only (shots-r4)

**Change:** with nothing ranked, the "How to get your logs here" panel follows the board below 1280 instead of arriving above it from the rail. From 1280 it stays in the rail. lhci phone CLS dropped from 0.093 to 0.0008.

## ux-designer: SHIP

1. **Order matches the mock:** the Encounters answer comes first, then the panel 16 px below it, then the footer. Nothing is inserted above the board.
2. **Polish:** two filled buttons now sit close together (Upload a log, Get the companion). A secondary button style would suit the panel's button if the design system has one. Not blocking.
3. **No JavaScript:** only the Encounters skeleton shows. The panel arrives after hydration below the board, so only the footer moves, and that is covered by the 0.0008 figure.
4. **Two copies of the panel:** `max-xl:hidden` and `xl:hidden`, with different testids (`rankings-companion`, `rankings-companion-below`). No e2e spec queries either testid.
5. **Not reviewed here:** 1280 and wider, which were not in scope for this round. The rail copy is unchanged from earlier rounds.
