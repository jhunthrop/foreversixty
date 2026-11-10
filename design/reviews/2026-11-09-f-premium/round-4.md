# Batch F, lane premium: review round 4 (2026-11-09)

Captures: build/ at c1b34f645. The verdicts are word for word; what the lane did follows.

- ux-designer: VERDICT: SHIP
- wow-player: VERDICT: SHIP
- ai-slop-reviewer: VERDICT: FAIL
- info-density-reviewer: VERDICT: DENSE WITH ADDITIONS

## Applied in round 5

- ai-slop 1 (tell 8, two stamped prose panels): "How the guild plans work" now sits inside the Questions panel as its second part, under its own heading. The page has the four plan cards, one prose panel and the coach sample. Contrast holds because the rules stay on a panel (ux's round 3 MUST).
- density 3: on a phone the coach sample comes before the long Questions panel.

## Not applied, with the reason

- ai-slop 2 ("billed yearly", "Billing interval"): these come from the shared plan summary (`perMonthLine`, used by account and checkout too) and the shared ControlBar label. They are handed to the shared-component pass.
- ai-slop 3 (audience lines, the Sample tag): the Sample tag is the coach sample spec's own label (polish 3.6, e2e "labelled Sample in its head and its caption"); the audience lines are PlanHead's required line. Both are non-blocking.
