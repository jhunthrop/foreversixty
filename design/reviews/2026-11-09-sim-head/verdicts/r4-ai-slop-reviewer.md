# ai-slop-reviewer, round 4 (captures in r4-after/)

VERDICT: PASS WITH FIXES

The round-3 findings are resolved. No head shows visual tells 1, 2, 5, 8, 12, 13 or 15.

1. Non-blocking, in the head: setup's "Updated 1 h ago from Battle.net" is accepted, because the sync date and the source change what the data means.
2. Non-blocking, in the head: drop "(95%)" from "± 0.3 DPS (95%)" on the browser, server and saved heads.
3. Non-blocking, in the head: the saved head's "Saved 4 Nov 2026" carries a year that the other heads leave out.
4. Non-blocking, in the head: setup is lopsided, with a lone title beside a six-line stack on the right.
5. Non-blocking, outside the head: small grey labels and explainer lines in the page bodies. This belongs to the design-system lane.
6. Non-blocking, outside the head: "Sample" and "Founder" on /premium are pills with real states behind them.

What reads as authentic:
- Heads that are a title plus an Updated stamp (changelog, privacy, 404, guides).
- Home's "58.4 DPS" with "Sim 8 h ago ›".
- The sim heads, where one figure dominates and the talent split and item level are real.
- The item and ability density below the heads.
