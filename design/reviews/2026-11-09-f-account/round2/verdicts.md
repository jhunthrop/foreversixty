# Lane f-account, review round 2 (2026-11-09)

Captures: the build at 509b1e305 (overwritten by round 3).

## ux-designer
Verdict: SHIP WITH FIXES
- MUST 1 /setup signed out 1024/1440: screenshot frames empty in the capture. The capture now decodes every image before the shot (loaded is not painted); the frames reserve their box (ui/Capture width/height), so nothing shifts in the product.
- SHOULD 4 /connect check line orphan: shortened ("...matches the one the companion shows.").
- Not taken: Switch 8 px past the edge (shared ui/Switch hit area), Plan label wraps in the rail, /setup 1024 step-3 and three-column grid (shared RowGrid/Steps), login reserved status line.

## wow-player
Verdict: SHIP WITH FIXES (MUST: the Settings screenshot empty in the signed-out capture; fixed as above)
- SHOULD 1: "The addon never talks to the network..." cut from the handshake panel.

## ai-slop-reviewer
Verdict: FAIL
- 1 empty Settings frame: capture fix.
- 2 the phone fold read as a placeholder: the fold is now a secondary button, "Show them, with screenshots from the game".
- Tell 8 cleared ("reads as the design system's framed panel").

## info-density-reviewer
Verdict: DENSE WITH ADDITIONS (the optional additions are listed in its reply; none are blocking)
