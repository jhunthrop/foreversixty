# Officer tools gate — experience spec

Status: spec, 2026-10-25. Implements the experience half of premium spec §4.7
(`docs/superpowers/specs/2026-10-24-premium-experience-design.md`). The product design is
`docs/superpowers/specs/2026-10-25-officer-tools-gate-design.md`. Base page:
`design/specs/2026-10-04-guild-page.md` §4.E, §4.F, §4.G. No new control and no new token (tenet 10).

## 1. The one question

An officer opens Loot or Readiness and asks: **"Why can't I press Award or Nudge, and how do I get them back?"**
The answer is one sentence at the top of the tab panel, above every row:
`Award and Nudge are part of the guild plan. Subscribe the guild.`
It is visible without scrolling at 360, 390, 1440 and 2000. It sits under the tab's h2, so it is
inside the first 160 px of the tab panel at every width.

A guild-settings visitor asks: **"What does this guild have, and until when?"** The answer is the plan
line, the first line under the `Billing` heading.

## 2. Reference

- **The in-game greyed action button.** An ability you cannot use stays on the bar, dimmed, in its slot.
  It is never removed. We borrow: the control keeps its place and label, dims, and stops responding.
- **The in-game red error text** ("Not enough rage"). One short line where you acted, then the UI
  reflects the new state. We borrow: the 403 says one line in the item card that was clicked.
- **Our own member read-only Loot tab** (guild spec §10.1, as built). We borrow its `Read-only`
  label at the h2's right edge.

## 3. Layout, region by region

### 3.1 The flag
- `Guild.svelte` derives `officerTools = (home.viewer.officer_tools ?? true) && !officerToolsEnded`.
- `?? true` is deliberate. The API ships the field first (product design §4.7, step 1). The server's
  403 is the real gate, and §4 below handles it.
- It passes `officerTools` to `GuildLoot` and `GuildReadiness` beside `officer`.
- The new state is **`officer-read-only`**: `officer === true && officerTools === false`. A member renders exactly as today: no sentence, no label, no controls.

### 3.2 The sentence (both tabs, once per tab)
- Once per tab, never once per control.
- Element: a `p` with `id="officer-tools-gate-{tab}"`, classes `text-text text-[13px] leading-5`,
  `data-testid="guild-{tab}-plan-gate"`. Barlow 400, 13 px, 20 px line, `--text` `#e9e4d8`.
- `Subscribe the guild` is an inline link to `guildSettingsHref(region, ruleset, name)`.
  It uses the inline-link recipe (DESIGN-SYSTEM.md, Accent): underlined, `--gold` `#e5b955` at rest,
  `--gold-hover` `#f5d27a` on hover and `:active`, 2 px gold focus ring at 2 px offset.
  The link is `whitespace-nowrap`, so it never splits across lines; at 360 the break falls after `plan.`.
  On phone a `::after` pseudo-element (`relative` on the link; `after:absolute after:inset-x-0 after:top-1/2
  after:h-11 after:-translate-y-1/2`, hidden from `md`) takes its hit area to 44 px: a fixed 44 px box centred
  on the link, because an inline link's box is its font's ~16 px content area, not the 20 px line. The line box stays 20 px and the focus
  ring stays on the link, hugging the text. (amended 2026-10-26, review round)
- **Loot tab:** directly under the header row (`Loot` h2 and `Read-only`), above the boss picker.
  It also renders above the empty state (`No bosses killed yet this tier…`). The tab's `gap-4` (16 px) spaces it.
- **Readiness tab:** directly under the `Readiness` h2, above the column header row (desktop) or the
  first stacked card (phone). The tab's `gap-3` (12 px) spaces it.
- The sentence names Nudge on the Loot tab too. That is §4.7's fixed wording; keep it verbatim.

### 3.3 The `Read-only` label
- Loot: the existing muted 12 px `Read-only` span at the h2's right edge now shows for
  `memberReadOnly` or `officer-read-only`.
- Readiness: the h2 gains the same header row (`flex items-center justify-between gap-3`) and the same
  span, in `officer-read-only` only.

### 3.4 The disabled controls (visible, never hidden — tenet 4)
Every control keeps its position, label, size and border. Each gets native `disabled` and
`aria-describedby="officer-tools-gate-{tab}"`.

| Control | Today | In `officer-read-only` |
|---|---|---|
| Award (each unawarded candidate row) | `border-[#3a3326]`, `text-strong`, 11 px uppercase, 44 px phone, 28 px from `md` | `disabled`; the existing `disabled:opacity-50`; add `disabled:cursor-default`. Never opens the confirm row. |
| Undo (winner row) | underlined `Undo`, muted 12 px, 44 px phone, 28 px from `lg` | `disabled`; add `disabled:no-underline disabled:cursor-default`. Reads `Awarded · Undo`, muted, plain. |
| Nudge (each row, both layouts) | `NUDGE_BTN`, 32 px | `disabled`; add `disabled:opacity-50 disabled:cursor-default` to `NUDGE_BTN`. Label stays `Nudge`. |

- None of the three has a hover style today. None gains one. Disabled shows no hover, focus or active state.
- Native `disabled` takes them out of the tab order. A loot list can hold 80 or more Award buttons;
  80 dead tab stops would be a defect.
- Keyboard order in `officer-read-only`: the tab strip, then `Subscribe the guild`, then whatever follows.
  Screen readers in browse mode read each button as dimmed, with the sentence as its description.
- axe: disabled controls are exempt from contrast (WCAG 1.4.3). The sentence is `#e9e4d8` on `--bg`.
  The link is gold and underlined, so `link-in-text-block` passes. The e2e runs axe on both tabs in this state.
- The Readiness footnote keeps `Nudge copies a message you can paste to them.` for officers in both
  states. It still describes a control that is on screen.

### 3.5 Height
The panel gains one line. Nothing else moves.

| Width | Lines | Loot adds | Readiness adds |
|---|---|---|---|
| 360 (324 px content) | 2, cap 2 | 56 px | 52 px |
| 390 (354 px content) | 1 or 2, cap 2 | at most 56 px | at most 52 px |
| 1440 and 2000 | 1, cap 1 | 36 px | 32 px |

`home` has answered before either tab mounts, so the sentence is in the tab's first render.
No layout shift on hydration. The tab skeletons (`min-h-[240px]` for Loot) are unchanged.

## 4. The 403 path (`guild_plan_required`)

The grant ended between page load and the click.
1. The officer confirms Award, or presses Undo. The API answers 403 with code `guild_plan_required`.
2. `Guild.svelte` sets `officerToolsEnded = true`. Both tabs flip to `officer-read-only` in that same
   render, with no reload and no refetch.
3. The clicked item's card shows one line, under its candidate rows. It reuses the tab's error-line
   recipe (`text-[13px]`, `role="alert"`) plus `leading-5`, `tabindex="-1"` and
   `data-testid="guild-loot-plan-refused-{itemId}"`.
   Focus moves to this line, because the button that held focus is now disabled.
4. The line stays until the page reloads. The tab-bottom `guild-loot-error` line is **not** set for
   this code: one message, not two.
5. Any other error keeps today's path (the bottom `guild-loot-error` line).

Needs: `GuildApiError` gains `readonly code: string` from the envelope. Branch on `code`, never on
`status` alone; the non-officer 403 is a different case.

## 5. Guild settings (`GuildSettings.svelte`, `settings-view.ts`)

### 5.1 A scoped grant (`plan.scope === "officer_tools"`)
`billingState` gains `'officer-tools'`. Inside the existing `Billing` section, top to bottom:
1. Plan line, 14 px `--text`, slot `guild-billing-plan`: `Officer tools · until 12 Dec 2026`.
   The date comes from `utcDayMonthYear(current_period_end)`.
2. Explanation, 13 px muted, in the `billedLine` slot (`guild-billing-contact`):
   `Granted by Forever Sixty for the first raid nights: Award, Undo and Nudge only, until 12 Dec 2026. After that they need the guild plan.` (amended 2026-10-26, review round)
3. Price line, 13 px muted: today's `notOnPlanLines().detail`, unchanged.
4. `Subscribe the guild`, `PRIMARY_BUTTON_FIXED w-fit px-4`, today's checkout href
   (`plan=guild&interval=monthly&guild_id=`). It is the same link as the not-on-plan state.
   A scoped guild can buy the plan (product design §4.1). There is no Manage billing.

The full-plan states `on-plan`, `granted` (`Guild plan · until 31 Dec 2026`) and `not-on-plan` are unchanged in words and layout.

### 5.2 Officers who do not hold the claim (`mayBill === false`)
- Today they see only `Only the guild master can subscribe the guild, release the claim or change billing.` (amended 2026-10-26, review round)
- New: when `guildBilling !== null`, the `Billing` section renders for them too. It holds the plan line
  and its explanation (5.1 rows 1 and 2, or the `on-plan` or `granted` rows 1 and 2), then the
  existing `guildMasterOnly` line in place of any button.
- Reason: the read-only sentence sends every officer here. They must find "until 12 Dec 2026", not a dead end.

### 5.3 An expired grant
`/v1/me` attaches only an active plan (`attachGuildPlans`), so an expired grant arrives as `plan: null`.
It reads as today's `not-on-plan` state. There is no `ended` line.

## 6. Account page — ruling

**The guild line belongs in the account page's `Guild` panel, on that guild's row. It never goes in the Premium block.**
- The Premium block (`AccountPremiumBlock.svelte`) lists what *this account* gets.
- A scoped grant gives the account nothing (product design anti-goal 1). A guild line there would read as Premium.
- The `Guild` panel row is where the guild already is, and it already carries per-guild facts.

The row:
- Shown when the guild entry's `plan !== null`. The API attaches `plan` for verified officers and
  leaders only, so members never see it.
- Position: a `basis-full` line after the name, the consent select and Leave, before the 21 px status line.
- Element: a `p` with `text-muted text-[13px] leading-5 basis-full`, `data-testid="account-guild-plan"`.
- Text: the line guild settings shows, then `·`, then the muted inline link `Guild settings` to
  `guildSettingsHref`. The link is underlined, `--text-strong` on hover and `:active`, with the 2 px gold focus ring.
  - `Officer tools · until 12 Dec 2026 · Guild settings`
  - `Guild plan · renews 31 Dec 2026 · Guild settings`
- `guildPlanLine` is extended to return the scoped line, so both pages call one function.
- Height: the row grows by 32 px (a 20 px line and a 12 px gap) at every width. It does not wrap at 360.
- The account page's stored reserve (Discord spec §9.1) absorbs this when warm. Cold is recorded, not gated.

## 7. Copy, verbatim

New strings go in `web/src/lib/guild/copy.ts` (tabs) and `settings-view.ts` (`settingsWords`).

| Where | String |
|---|---|
| Gate sentence, Loot and Readiness | `Award and Nudge are part of the guild plan. Subscribe the guild.` (link text `Subscribe the guild`) |
| Header label (reused) | `Read-only` |
| 403 on Award | `Not awarded. This guild's officer tools have ended.` |
| 403 on Undo | `Not undone. This guild's officer tools have ended.` |
| Settings plan line, scoped | `Officer tools · until {d Mon yyyy}` |
| Settings explanation, scoped | `Granted by Forever Sixty for the first raid nights: Award, Undo and Nudge only, until {d Mon yyyy}. After that they need the guild plan.` (amended 2026-10-26, review round) |
| Settings price line (reused) | `notOnPlanLines().detail` |
| Settings button (reused) | `Subscribe the guild` |
| Settings, not the claim holder (reused) | `Only the guild master can subscribe the guild, release the claim or change billing.` (amended 2026-10-26, review round) |
| Account guild row | `{plan line} · Guild settings` (link text `Guild settings`) |

Out of lane (flag only): `notOnPlanLines().detail` still says "progression tools". §4.7 rules that the
page names only what ships. The premium lane that owns that copy fixes it.

## 8. Acceptance screenshots

The PR attaches each one. Stub `home.viewer.officer_tools` in the guild fixture and `plan.scope` in `/v1/me`.

| # | State | 360 | 390 | 1440 | 2000 |
|---|---|---|---|---|---|
| 1 | Loot, officer, `officer_tools: true` (controls unchanged) | | x | x | |
| 2 | Loot, `officer-read-only`, drops listed | x | x | x | x |
| 3 | Loot, `officer-read-only`, empty (no kills) | | x | x | |
| 4 | Loot, `officer-read-only`, an awarded item (Undo disabled) | | x | x | |
| 5 | Loot, member (no sentence, no controls) | | x | x | |
| 6 | Loot, after the 403 on Award: row line focused, sentence shown | | x | x | |
| 7 | Readiness, `officer-read-only` (grid and stacked cards) | x | x | x | x |
| 8 | Readiness, member | | x | x | |
| 9 | Focus-visible on `Subscribe the guild` (keyboard) | | x | x | |
| 10 | Settings, scoped grant, claim holder | x | x | x | |
| 11 | Settings, scoped grant, officer without the claim | | x | x | |
| 12 | Settings, `granted` full plan (regression) | | x | x | |
| 13 | Account, Guild panel: one scoped row and one paid row | x | x | x | |

Checks in the e2e (`guild-officer-tools.spec.ts`):
- axe clean on #2, #7, #10 and #13.
- In #2 and #7, every Award, Undo and Nudge is present and `disabled`; the counts equal #1's counts.
- Clicking a disabled control opens no confirm row and writes nothing to the clipboard.
- The sentence's link resolves to `/guild/{region}/{ruleset}/{slug}/settings`.
- The sentence's bounding box is inside the viewport on load at 360 and 1440.
- On #6 no reload happened (a `window` marker survives), and Readiness opened afterwards is read-only.
- Lighthouse budgets in `web/lighthouserc.json` are unchanged. These states are signed in, so they are not in the lhci set.
