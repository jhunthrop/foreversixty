# Officer tools gate: acceptance captures

Captures for the experience spec `design/specs/2026-10-25-officer-tools-gate.md` §8. They are
made by `web/tests/captures/officer-tools-gate/capture.spec.ts`, run against a fixture build
served on port 4431:

```
cd web
FOREVER_DATA=fixture npm run build
npx astro preview --outDir dist --port 4431   # in the background
npx playwright test -c tests/captures/officer-tools-gate/playwright.capture.config.ts
```

The stubs are shared with the e2e (`web/tests/e2e/support/officer-tools.ts`). The guild is the
seeded mock guild, Olympus XXVII. The settings and account states use `meWithGuildPlan` and
`scopedGrant` from `web/src/test-support/billing-fixtures.ts`.

## Capture list

Each state has a full-page PNG, `officer-tools-<state>-<width>.png`, and a panel crop,
`officer-tools-<state>-panel-<width>.png`. A long tab's crop is its top 520 px.

| # | State | File stem | Widths |
|---|---|---|---|
| 1 | Loot, officer, tools on | `01-loot-tools-on` | 390, 1440 |
| 2 | Loot, officer-read-only, drops listed | `02-loot-read-only` | 360, 390, 1440, 2000 |
| 3 | Loot, officer-read-only, no kills | `03-loot-read-only-empty` | 390, 1440 |
| 4 | Loot, officer-read-only, the awarded item (Undo disabled); the crop is the item card | `04-loot-read-only-awarded` | 390, 1440 |
| 5 | Loot, member | `05-loot-member` | 390, 1440 |
| 6 | Loot after the 403 on Award; there is also a `-card-` crop of the refused item | `06-loot-refused` | 390, 1440 |
| 7 | Readiness, officer-read-only (grid, and stacked cards on a phone) | `07-readiness-read-only` | 360, 390, 1440, 2000 |
| 8 | Readiness, member | `08-readiness-member` | 390, 1440 |
| 9 | Focus-visible on `Subscribe the guild`; the crop is the sentence | `09-link-focus` | 390, 1440 |
| 10 | Settings, scoped grant, claim holder | `10-settings-scoped-claim-holder` | 360, 390, 1440 |
| 11 | Settings, scoped grant, officer without the claim | `11-settings-scoped-officer` | 390, 1440 |
| 12 | Settings, granted full plan (regression) | `12-settings-granted` | 390, 1440 |
| 13 | Account, Guild panel: one scoped row, one paid row | `13-account-guild-rows` | 360, 390, 1440 |

## Heights against §3.5

Each "adds" value is the tab panel's height with the sentence minus its height without it,
at the same width and with the same data. The sentence's own box height is in brackets.
The raw numbers are in `heights-<tab>-<width>.json`.

| Width | Loot adds (cap) | Readiness adds (cap) |
|---|---|---|
| 360 | 56 px (56), sentence 40, 2 lines | 52 px (52), sentence 40, 2 lines |
| 390 | 36 px (≤ 56), sentence 20, 1 line | 32 px (≤ 52), sentence 20, 1 line |
| 1440 | 36 px (36), sentence 20 | 32 px (32), sentence 20 |
| 2000 | 36 px (36), sentence 20 | 32 px (32), sentence 20 |

Every width is at or under its cap. At 360 both tabs land exactly on the cap. At 390 the
sentence fits on one line.

## Noticed while capturing

- **The link breaks across lines at 360.** On both tabs, the second line reads `Subscribe the` /
  `guild.` (02 and 07 at 360). The heights are within the cap. A non-breaking space inside the
  link text would keep it whole; that needs a design call.
- **The phone focus ring is 44 px tall.** The link's `py-3` (44 px hit area) is inside the
  outline, so at 390 the focus ring is a tall box. It reaches down over the top of the boss
  picker (09 at 390). At 1440 the ring hugs the text. This is visible only on keyboard focus,
  which is rare on a phone, but the ring is larger than the line it marks.
- **The account row's `Guild settings` link is gold.** Spec §6 calls it a "muted inline link",
  `--text-strong` on hover. The capture (13) shows it gold and underlined. It is probably the
  global `a` colour, since the component sets no colour at rest. Worth a look.
- **No visible focus on the refused line.** In state 6 the refused line has focus
  (`tabindex=-1`, which the e2e asserts), but it shows no ring because the focus is
  programmatic. That matches the spec, which defines no focus style for the line. Screen
  readers get it through `role="alert"`.
- **Disabled and enabled Undo look alike.** On the awarded item at 1440, the disabled
  `Awarded · Undo` is muted text with no underline. Without the underline that marked it as a
  link, it reads as a plain label, which is the intent. Next to the `Already awarded` labels it
  is easy to miss that Undo exists at all.
- **A misleading doc comment.** In `web/src/lib/guild/api.ts`, `GuildViewerView.officer_tools`
  says "The guild is on a scoped officer-tools grant". The page (and spec §3.1) reads it the
  other way: `true` means the officer tools are available, `false` means read-only. The comment
  should say that.
- **Fixture oddities, not defects.** The awarded Helm of Wrath reads "Awarded to Kraggor … by
  Kraggor". The settings price line still says "progression tools", which spec §7 already
  flags as out of lane.
- **Not done here.** The ux-designer and wow-player reviews named in the task-9 brief, Step 4,
  were not run, because this task was dispatched with "never dispatch subagents". The
  side-by-side with the mock was not made either. Both are still needed before merge.
