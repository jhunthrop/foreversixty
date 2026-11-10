# h-audit review brief (2026-11-09)

Fresh-eyes audit of pages reviewed during the redesign before verdicts were saved. Owner goals tonight: "fix any and all AI slop that still exists" and "update ANY page with the new style that hasn't been updated."

- Repo checkout: /Users/jh/code/forever/.claude/worktrees/agent-a4344aef401fed698 (branch lane/h-audit). Read code there, not in the main checkout.
- Captures: design/reviews/2026-11-09-h-audit/<round>/<page>-<1440|390>.png, full-page, fixture build (fixture characters, fixture data; blank spell icons or odd fixture names are build artefacts, not findings).
- The new style: docs/design-system.md, design/specs/2026-11-09-directions.md, design/specs/2026-11-09-feedback-ai-look.md, the direction mocks in design/mocks/directions/png/.
- Owner rules: compact title band (no big header); class or quality-coloured lists sit on a plate; a number takes its unit and nothing else; no em dashes; player voice and game words; dense but clean; never a thin page or an empty icon square; one job per page (BiS owns gear, planner talents, sim numbers).
- For every finding name the element, where it is (page, width, region) and, if you can find it, the source file that renders it (grep web/src).
