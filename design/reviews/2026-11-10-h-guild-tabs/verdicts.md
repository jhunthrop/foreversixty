# h-guild-tabs: review verdicts (batch H)

Tabs: Raids (planned and logged nights), the raid planner (officer and member), Progression, Applications, Recruit, Settings.
States: `dense-raids`, `dense-planner`, `dense-raid-member`, `dense-progression`, `dense-recruit`, `dense-applications`, `dense-settings` (design/mocks/web/guild-dense.json), captured by web/tests/e2e/guild-captures.spec.ts with `GUILD_CAPTURE_CONFIG=guild-dense.json` at 390, 1024, 1440 and 2000.
Mocks: design/mocks/ws_guild_dense_tabs.py (fixtures design/mocks/guild_dense_tabs_fixtures.py), PNGs in mock-r1/ to mock-r3/. Baseline before the lane: r0/.

## Round 3 (build-r3)
- guild-officer-reviewer: VERDICT: CLOSE (member readiness, the cause of "short", empty tiers)
- ux-designer: VERDICT: FIX (note field overflow, planner rows without Last 5 and the cause)
- wow-player: VERDICT: FIX (implausible Onyxia kill times and hybrid DPS, pulls and deaths disagreeing between tabs)
- ai-slop-reviewer: VERDICT: FAIL (card inside a card, qualifiers, stat tiles, missing spaces)
- info-density-reviewer: VERDICT: DENSE WITH ADDITIONS

## Round 4 (build-r4)
- guild-officer-reviewer: VERDICT: CLOSE (member and officer short counts disagreed)
- ux-designer: VERDICT: SHIP
- wow-player: VERDICT: SHIP
- ai-slop-reviewer: VERDICT: FAIL (inner frame on Guild plan, applicant lines wrapping mid-phrase)
- info-density-reviewer: VERDICT: DENSE WITH ADDITIONS

## Round 5 (build-r5)
- guild-officer-reviewer: VERDICT: CLOSE (Absent was a count without names)
- ux-designer: VERDICT: FIX (availability overprinting the item level at 1440)
- wow-player: VERDICT: SHIP
- ai-slop-reviewer: VERDICT: PASS WITH FIXES
- info-density-reviewer: VERDICT: DENSE WITH ADDITIONS

## Round 6 (build-r6)
- guild-officer-reviewer: VERDICT: WORTH PAYING FOR
- ux-designer: VERDICT: FIX (applications at 1024)
- wow-player: VERDICT: SHIP
- ai-slop-reviewer: VERDICT: PASS WITH FIXES
- info-density-reviewer: VERDICT: DENSE WITH ADDITIONS

## Round 7 (build-r7)
- guild-officer-reviewer: VERDICT: WORTH PAYING FOR
- ux-designer: VERDICT: SHIP
- wow-player: VERDICT: SHIP
- ai-slop-reviewer: VERDICT: PASS WITH FIXES (one pull a line, "Level 60" at the cap)
- info-density-reviewer: VERDICT: DENSE WITH ADDITIONS

## Round 8 (build-r8): the gate, final
- guild-officer-reviewer: VERDICT: WORTH PAYING FOR
- ux-designer: VERDICT: SHIP
- wow-player: VERDICT: SHIP
- ai-slop-reviewer: VERDICT: PASS
- info-density-reviewer: VERDICT: DENSE WITH ADDITIONS

## Open polish (not gating)
- Absent names on an opened night could link to the member's roster row (attendance/recent has the account).
- Roster bests: a per-raider trend (best vs last night) and a % of sim bar; sortable heads.
- Raids: sortable nights, Copy CSV, a link from each night to its log report; the planned rows' short-by-role figures.
- Applications at 390: Decline beside Accept; a non-breaking space in "Top HPS 612". Recent decisions could carry dates.
- Settings: the billing, invite and Discord components' own lines ("Billed to you.", "Made 1 Dec by Dunmaro.", the Founder and Connected pills) belong to those components' lanes.
- The shared guild TabStrip clips at 390 (an edge fade or centring the active tab).
- The invite URL in the captures is the preview's origin; production builds it from foreversixty.gg.
