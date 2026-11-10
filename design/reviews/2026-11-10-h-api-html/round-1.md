# h-api-html review, round 1

Captures: after/ (side-b-* = parity mock left, build right; app-confirm-* and mock-app-confirm-*), before/.

## ux-designer: VERDICT: SHIP
No blockers. Nits: the mock's nav icons and the 390 tab style differ from the built chrome (mock library, not this lane);
balance the confirm title so "phone" is not orphaned; confirm the button has focus-visible and active states (it does);
the plate's dark fringe at 390 on the not-found page is faint.

## ai-slop-reviewer: VERDICT: PASS WITH FIXES
- Confirm page copy repeats "Forever Sixty app on your phone" three times: cut the body and the not-you line. FIXED.
- Not-found: "It may have been mistyped." states the obvious. FIXED (handler.go).
- "Paste in game with /fs." and the "Gear for 60" label on /b/{id}: the planner island's own panels (web planner
  components), not this lane's files. Follow-up for the planner owner.
- Inner frame on the single-paragraph plates: kept. It is the system's framed panel, the shape of the gated
  b-unavailable.astro and directions mock b4-unavailable, on the page's one featured panel (spec: frame two or three
  panels per page, the lead one first).

## wow-player: VERDICT: SHIP
Copy-level notes. Confirm page repeats "on your phone": FIXED (title "Sign in to the app" under the wordmark, body "Sign
in on Battle.net and the app picks it up from there.", not-you "Didn't start this sign-in from the app? Close this
page."). Planner notes (name the matched guide build, /fs before the addon link, 41 views at 390 only, point order
column heads 0..9) are planner-island follow-ups.
