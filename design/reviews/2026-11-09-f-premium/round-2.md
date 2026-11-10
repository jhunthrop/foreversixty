# Batch F, lane premium: review round 2 (2026-11-09)

Captures: build/ at 78e60150f. Verdicts word for word, then what the lane did.

- ux-designer: VERDICT: SHIP WITH FIXES
- wow-player: VERDICT: SHIP WITH FIXES
- ai-slop-reviewer: VERDICT: FAIL
- info-density-reviewer: VERDICT: DENSE WITH ADDITIONS

## Applied in round 3

- Checkout back from Stripe: the finished three-step list is gone. One framed panel reads "Premium is on." with the renewal line, then rows that go to what Premium turned on (Upload a log for the coach, Run a server sim, See your plan) (wow M1, ai-slop 1; CheckoutConfirmed.svelte).
- Checkout for an account that already has Premium: the head draws no price; the answer keeps its box so nothing shifts (wow M1).
- Coach sample: rows with no icon draw no lead column, so they sit flush; Fix first is an open row, no box inside the panel (ux S2, ai-slop 2); the panel stays in view beside the longer column (ux S3).
- How the guild plans work is an open section on the ground; Questions stays the one prose panel (ai-slop 2).
- Copy: Verified members leads with "Each player counts once" (wow S1); Are there ads? is "No, for anyone. The site is paid for by Premium and the guild plans." (wow S4); the scope line is "Sims cover damage specs. The coach reviews every spec." (wow S7, ai-slop 4); the Questions meta is "Prices in US dollars." (ai-slop 3).
- /creators worked examples: a Guild row, and each row's list-price figure after 6 Jan 2027 under it (density 1, 2; wow S5).
- ux M1: the 1:1 crop of the Premium card at 390 (free state) is build/premium-free-390-premium-card-1to1.png; the rows are clean.

## Not applied, with the reason

- ai-slop 3 "Billing interval": the shared ControlBar's group label (not page code; batch E owns the shared header controls).
- ai-slop 3 "$3.33 a month, billed yearly" on the checkout head: the plan summary's second line (shared CheckoutHead model, priceCell) that /premium shows too.
- ai-slop 5 /creators wrapper and contents rail: the directions mock b4-creators draws exactly this reading panel and rail.
- wow S2 (RowGrid row tops): the shared RowGrid's alignment.
- wow S3 (the coach fixture's Bloodthirst sentence): the logs lane's fixture text.
- wow S6 (/creators in the nav): it is not in the nav; nothing links it yet (the footer is the shared chrome).
- density 9 (rules as expanders at 390): the premium spec keeps the rules and questions open (test "opens the rules and the questions").
