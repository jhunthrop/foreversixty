# Lane D, round 2 verdicts (captures of 47aa6d93d and cc9723dbe)

Captures: `after/<page>-{1440,390}.jpg` (home-raid, home-out, guild-raids, guild-raid, guild-raid-officer, guild-roster, rankings, tiers, guides, character, 404), `discord/*.png`. Captures serve real icons.

| Reviewer | Round 2 | What it asked | Final |
|---|---|---|---|
| ux-designer | SHIP | "as DPS" wrapped at 390; the planned rows' In chip; then the missed raiders' commas | SHIP |
| wow-player | FIX | Absent and Bench names not in class colour; the count line; "Feral Bear"; the phone spec select; the leaders' names | SHIP (holds noted below) |
| ai-slop-reviewer | PASS | Tag repeating its group head; "as DPS" at 390; the commas | PASS |
| info-density-reviewer | FAIL | A rail panel for the member's raid page; Last 5 and Ready on the raid rows | PASS (the rail panel is out of lane) |
| guild-officer-reviewer | FIX | Ready dashes against the roster's marks | SHIP (member view by design; the officer capture shows the marks) |
| brand-consistency-reviewer | FAIL | The build card drew the bare crest art | PASS |

## What changed (3090d2d15, 47aa6d93d, cc9723dbe, and the mock regen)
- The build card draws the ringed crest: `sheet.CrestImage` is the one exported crest image, the bare art (`embeddedCrest`) is unexported, `TestRenderDrawsTheRingedCrest` checks the ring and the clear corner.
- A sign-up row under a Late, Bench or Out group head carries no state Tag (`stateNamed`); the Class grouping keeps it.
- The raid nights' Absent and Bench names are `CharacterName` in class colour, ", " between them in the label's grey.
- The home plaque keeps "as DPS" on the choices' line at every width.
- The capture set adds the officer's raid page (`guild-raid-officer`), which shows the Last 5 strips and the Ready marks on every row; a member sees only their own.
- The chrome snapshot and the shared-build fixtures regenerated from the lane's header; the pages facts list one spec per talent tree, so the 404 and /b mocks match the pickers.

## Out of lane, recorded for the owning lanes
- The raid page's count line reads In plus Late ("29 of 40 signed up · 27 In · 2 Late · 2 Bench"); wow-player would rather it said "in or late" (raid comp spec).
- "Feral Bear" is the bear guide's name; wow-player would rather "Feral (Bear)" or "Feral Tank" (guide catalogue).
- The member's raid page leaves its right rail short; info-density proposes a "Cover the gaps" panel from the Bench, Late and Out rows against the Short roles (guild lane).
- The officer's spec select and some row sub lines truncate at 1440 ("Beast Mastery · …", "Restoration Shaman · Hea…"); the Comp table's Short column is red on every row.
- Home: the Latest feed panel's empty band on home-out; the tier letter on the spec card against the tier list; "Warden Kelthas", "Hyjal Summit" and "Barrow Deeps" in the fixture's first raids.
