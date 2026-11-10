# Batch F, lane premium: review round 3 (2026-11-09)

Captures: build/ at 9716c238d. The verdicts are quoted word for word, followed by what the lane did.

- ux-designer: VERDICT: SHIP WITH FIXES
- wow-player: VERDICT: SHIP
- ai-slop-reviewer: VERDICT: PASS WITH FIXES (no visual tell 1, 2, 5, 8, 12, 13 or 15 remains)
- info-density-reviewer: VERDICT: DENSE WITH ADDITIONS

## Applied in round 4

- ux M1 (contrast): "How the guild plans work" goes back on its panel. Measured on the round 3 capture, the open section put muted text (#8b8784) on the painted ground (about #70623f), roughly 1.3:1, which fails. On the panel it clears the bar, as in the directions mock and as on Questions. ai-slop accepted the panels for the plan comparison, and its remaining note was about prose panels. Here contrast decides.
- ux S4: the confirmed panel's See your plan row has its own line, "Renewal date, receipts and cancel".

## Not applied, with the reason

- ux S2 (the Free and Raid guild voids): the columns stretch so every buy button sits level across the row, which ux accepted. The space stays.
- ux S3 and S5 (/creators at 2000, the contents rail): the shared ReadingPanel's prose and lane grid, and the shared Toc. These are page-column and component rules, not this page's.
- ai-slop 1 to 5 (the per-month line, Billing interval, the audience lines, the Sample tag, the creators wrapper and rail): non-blocking. These are shared lines and primitives, or required by spec (the Sample tag), or drawn by the directions mock.
- density 3 to 5: optional. They need the account read or a spec change (the rules stay open).
- wow 1 to 4: these belong to other lanes (the coach fixture, the shared RowGrid, the account's Discord state, the footer).
