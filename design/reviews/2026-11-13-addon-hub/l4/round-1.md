# L4 Rotation tab: round 1 verdicts and what changed

Boards: `rotation-60.png`, `rotation-34-stepped.png`, `rotation-no-build.png` (design/mocks/addon_hub_rotation.py).

| Reviewer | Verdict |
|---|---|
| ux-designer | FIX: weights block conflicts with spec 4.3; raid name missing from the note; biggest-loss kinds not spell-specific; "1 boss" rows mix units in the warning colour; Golemagg cut; Me link not in the spec |
| wow-player | FIX: weights label read level 30 beside a Level 40 stepper; "No raid yet" pitch; Golemagg cut; Attack power and Ranged attack power duplicate rows; dimmed condition hard to read; demo night dated before raids open |
| ai-slop-reviewer | FAIL: "measured at level N" qualifier; "for Beast Mastery" doubled in the header; summary line restates the header figure; "No raid yet" pitch copy; hollow right column; Golemagg cut |
| info-density-reviewer | FAIL: parse column, loss seconds on every row, loss to ability mapping, history delta, current vs best per stat |
| brand-consistency-reviewer | FAIL: stat names not in the web's words and case; haste line not the web's "per 1% haste"; inconsistent "for Beast Mastery" |

Round 2 changes:
- Stat weights stay (the lane brief asks for them with their source) and follow the stepper's level: the note reads `shipped, level 40` at Level 40, or `your sim, 14 Dec`, the Gear spec's own words (`gearWeightsShipped`, `gearWeightsSim`).
- Stat names from the web's weights tables (`Hit`, `Crit`, `Attack power`, `Ranged attack power`); attack power folded into ranged attack power where the spec weighs them alike; at most six rows; a sim cap reads `Capped`; haste reads `Haste: 0.54 DPS per 1% haste`.
- The summary line is now the table's total row, `All 8 bosses 85%`, as a meter's total line.
- Lost across the night: the seconds kinds first in the warning colour, the kinds the API counts no seconds for (`1 boss`) muted below them.
- Boss column 136 px; a dimmed line dims its number, icon and name only, the condition stays readable.
- No build: the header drops the spec name; Priority keeps `for Beast Mastery` (spec 8).
- No raid with the companion running: `No raid logged yet.`; the spec's `No raid yet. The companion app brings it in.` only without the companion.
- The Me link reads `See parses and deaths`, shown only while the Me tab exists; the boss hover adds the parse.
- Demo night moved to Tue 15 Dec.

Declined, with the reason:
- Raid name: `last_raid` carries none and the addon has no boss to raid table; a data follow-up (add `raid` to the message), not faked here.
- Spell-specific losses (`Aimed Shot late`): the API's biggest-loss phrases are downtime, dead early, died to, standing in fire, missing consumables, missed mechanics, low output and utility (api/internal/rating/last_raid.go); seconds exist only for downtime and time dead.
- Parse column: Me owns parses (one job per tab); the hover and the Me link carry them.
- Kill order: spec 4.4 says least followed first.
- History delta, current vs best per stat, loss to ability mapping: no data in the addon today.

## Round 2 verdicts and round 3 changes

| Reviewer | Round 2 |
|---|---|
| ux-designer | SHIP with fixes (weights accepted as the integrator's call; Golemagg cut; total row reads as a ninth boss; raid name and loss kinds are spec follow-ups) |
| wow-player | FIX (Golemagg cut; "shipped" is dev talk; Hit and Crit need their rating unit; the 34 header; the Arcane Shot wrap) |
| ai-slop-reviewer | FAIL narrowly ("shipped"; "All 8 bosses"; Golemagg; hollow right rail on the no-raid board; "for Beast Mastery" intent) |
| info-density-reviewer | PASS (optional additions only) |
| brand-consistency-reviewer | FAIL close ("shipped"; the rating note; attack power merge to explain; one colour for the losses) |

Round 3:
- Boss cells print the name before its epithet (`Golemagg`, `Vaelastrasz`), as raid frames do; the hover keeps the whole name.
- Weights note: `at level 60` (the web's words) or `your sim, 14 Dec`.
- The web's rating note under the table: `Hit and Crit are per rating point: 10 hit rating is 1% hit, 14 crit rating is 1% crit.`
- Total row: `8 bosses` muted, the mean strong.
- Every loss value in the warning colour (spec 3.5: a loss), seconds kinds first.
- Kept, with the reason: the 34 header is `Follow.buildName`, the Builds tab's own name for that build; `for Beast Mastery` shows only with no build, as spec 8 asks, while the header then names no spec; the no-raid right column is the spec 8 single line; attack power is folded into ranged attack power because a hunter's attack power counts for ranged attacks and the spec weighs them alike.

## Round 3 and 4

Round 3: ux-designer FIX (untimed losses red again), wow-player SHIP, ai-slop FAIL (the rating footnote, a caption that repeats the stepper), brand PASS.

Round 4 changes: Lost across the night lists only the timed kinds (downtime, time dead), all in the warning colour; the untimed ones stay on their boss rows. Rating stats read `Hit rating`, `Crit rating`, the conversion in the row's hover (the web's ratingConversion); the footnote is gone. The weights caption shows only when it adds something: `your sim, 14 Dec`, or the band when it differs from the stepper's level.

## Final

| Reviewer | Verdict |
|---|---|
| ux-designer | SHIP (round 4) |
| wow-player | SHIP (round 3) |
| ai-slop-reviewer | PASS (round 4) |
| info-density-reviewer | PASS (round 2) |
| brand-consistency-reviewer | PASS (round 4) |

For the integrator: add `raid` to `last_raid` (spec 4.4's "Molten Core · Tue 8 Dec"); amend 4.3/4.4 for the weights the Rotation tab also shows (the lane brief); amend 4.4's loss kinds to the API's phrases. Polish left: the Arcane Shot condition wraps to "less" (shipped data).

After the verdicts: the stepper labels became plain `<` and `>` (the client font draws the angle quotes as dots; settings_spec guards it); boards re-rendered.
