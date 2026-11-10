# Addon first-run tour (2.0.38), info-density-reviewer, round 3 (final captures in round-3/)

**Verdict: DENSE.** No must-fix items.

## Why it passes

- **Page 2:** it shows a real tooltip on the character's own upgrade or Data pick, with the source and the drop chance.
- **Last page:** it reads the character's state:
  - the path, its points on plan and the next talent;
  - whether the companion is fresh, with a warning after 3 silent days, else the connect links;
  - the spec buttons, each with its playstyle line.
- **Pages 1 and 3:** single-screen maps. Every skip path takes one click. Nothing is invented, and "coming" keeps the truth rule.

## Optional shoulds, ranked

1. **Page 2, ring row:** show the real next point ("Level 25: Lethal Attacks") when a build is followed. The inputs already hold `progress.nextLevel` and `pointName`.
2. **Page 2, rotation row:** show the panel's first rotation button at this level. Only if it is cheap to read.
3. **Page 4:** show the guild's real posted raid and job when one exists, and keep the fixed example otherwise.
4. **Page 3:** show the character's Forever Score figure when one exists.
5. **"Open my panel":** already the right drill-down.

## Do not add

- A stats page, or a per-slot upgrade list. The panel owns those.
- Explainer copy.
