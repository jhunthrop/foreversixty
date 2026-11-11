# ai-slop-reviewer, Android secondary screens, round 1 (2026-11-10)

Evidence: the six capture sheets in `android/` and `strings.xml`.

| Group | Verdict | Findings |
|---|---|---|
| guides | FAIL | **SHOULD:** grey and salmon text on the tan ground is barely legible (the inactive segments, the section labels, the error line). **SHOULD:** the explainers say the obvious: "Guides update by themselves...", and the tier legend beyond the ≈ meaning. **SHOULD:** the loading tile is blank. **NICE:** the large-text Your guide card wraps badly. **NICE:** a "Site" chip on every What changed row. |
| character | FAIL | **BLOCKER:** qualifiers the context already gives: "As of", "percentile", "points", and the Gold dot footnote. **SHOULD:** a loading tile without the real head. **SHOULD:** every Forever Score bar is the same yellow. **SHOULD:** the large-text head wraps to five lines. **NICE:** the Epic pill. |
| bis-search | FAIL | **SHOULD:** a skeleton loading state. **SHOULD:** the grey summary sentence wraps, and "Runner-up" repeats on every row. **SHOULD:** a "Yours" pill on a heading. **SHOULD:** three stacked filter chips take a third of the screen. **NICE:** the picker selects Alliance while the page says Horde. |
| account | PASS | **SHOULD:** a selling footer on Plan. **SHOULD:** the Founder and Premium pills repeat the row text. **SHOULD:** contrast on the muted lines. **NICE:** a redundant devices note. |
| guild-raids | FAIL | **BLOCKER:** cards inside cards on the raid page (roster, Your sign-up, composition, buffs). **SHOULD:** the buff row is abbreviations in empty squares, with no icons. **SHOULD:** an empty checkbox on every sign-up row. **SHOULD:** a dim chip row, and an explainer on locking. **NICE:** the same "In" dropdown on every row. |
| sims-reports-talents | FAIL | **SHOULD:** a blank loading state. **SHOULD:** grey, unreadable lines on the tan ground. **SHOULD:** explainers ("Start a sim...", "Server runs keep going..."). **SHOULD:** a light DPS unit, and "Tuesday" as a sim label. **SHOULD:** colour the progress by state. |

Across all six: there are no gradients, glows or glass. The main fix everywhere is the text contrast on the tan ground. In the copy, "Unlock sign-ups" should read "Open sign-ups".
