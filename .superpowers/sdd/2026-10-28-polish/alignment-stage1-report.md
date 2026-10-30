# Site alignment, stage 1 (the shared pieces): report

Spec: `design/specs/2026-10-29-site-alignment.md`, amended in this stage with the coordinator's and both persona gates' rulings (commit `fdf131ad5`). The per-route lanes build on the APIs below exactly as written.

## Rulings written into the spec

- D1: no band anywhere. Every `PageHead` sits on `--bg`. §3.1.2 and §3.1.9 drop the band, S2 has no `band` prop, §6.3 says no legacy navy band survives on any page, every "on the band" in §5 is struck, the guild faction art goes with the band (§5.10), `ClassHeader` and `ArtPanel` lose their navy gradient (§5.3, §5.4), and /account is untouched.
- D2: IconTile stays 36 px. The new compact head (§3.3.1, 48 px) serves dense data cards: Stat weights, This set, Play it, the guide rail, the planner's spec trees and right-column cards (compact on both, for consistency). It supersedes the old §3.4.5 phone-head rule.
- §5.5: the duplicated "your level's range" line on the planner's build card is dropped.
- D3: one shared `TabStrip` (S7), built in this stage.
- D4: the About copy is adopted.
- §5.17 blocker order: Guild settings groups, then the tenet-2 item icons (guild loot, /bis "Where to get it"), then the rest. The home launch line is "verify after merge".

## Component APIs

All live in `web/src/components/ui/`, and their class recipes in `web/src/lib/ui/`. Each Astro twin ships no client JavaScript, except `TabStrip.astro`'s fade script.

### PageHead (`PageHead.svelte`, `PageHead.astro`; recipe `lib/ui/page-head.ts` `PAGE_HEAD`)

Svelte props:

- `eyebrow: string`
- `title: string` (the page's one `h1`, `heading-page`)
- `titleClass?: string`
- `titleColor?: string` (a CSS colour, e.g. `var(--color-class-warrior)`)
- `titleId?: string`
- `job?: string` (15/22, max 640)
- `lead?: Snippet` (64 px crest, at the title block's left)
- `meta?: Snippet` (13 px muted line; the Updated stamp)
- `aside?: Snippet` (the identity block or the page's one control; bottom-aligned from 1024, stacked below it)
- `below?: Snippet` (tabs or a controls row, 16 px under)
- `testid?: string`

The Astro twin takes the same props minus the snippets, and named slots `lead`, `meta`, `aside`, `below`. Test ids are `page-head-job`, `page-head-meta`, `page-head-aside`. There is no band and no card. /account renders it through `AccountTitle` (`aside` = identity block), pixel for pixel.

### Card (`Card.svelte`, `Card.astro`; recipe `lib/ui/card.ts`)

Svelte props:

- `title: string`
- `glyph?: GlyphName` (omit it and `lead` for a tile-less prose aside)
- `tone?: 'gold' | 'discord' | 'muted'`
- `lead?: Snippet` (replaces the tile: an item icon, a crest)
- `label?: string` (micro-label; the title then keeps its own case)
- `titleSnippet?: Snippet` (beside the title, e.g. a Founder tag)
- `meta?: string` (13 px muted, right)
- `action?: { label: string; href: string }` (an ActionLink, right)
- `aside?: Snippet` (a Tag or one sm button, right)
- `quiet?: boolean` (`card-surface-quiet`, no gold top edge)
- `body?: 'text' | 'rows' | 'table' | 'none'`, default `'text'`:
  - `text` is `px-5 py-4` with the top rule, `gap-3` and 14/22;
  - `table` is `px-5 pb-2` with the top rule;
  - `rows` is a full-bleed flex column (rows such as `SettingsRow` and `ROW` bring their own top rules);
  - `none` puts the children straight into the section.
- `foot?: Snippet` (`px-5 py-4`, top rule, `gap-3`)
- `headingLevel?: 2 | 3`
- `compact?: boolean` (48 px head, 28 px tile, `section-title` 18 px, no micro-label)
- `testid?: string`
- `id?: string`
- `class?: string`
- `children: Snippet`

The root is a `<section class="card-surface overflow-hidden">`. The Astro twin takes the same props minus the snippets, with slots default, `lead`, `aside` and `foot` (no `titleSnippet`). The `data-card-body` and `data-card-foot` attributes mark the body and foot. The head's meta test id is `card-head-meta`.

### CardHead (`CardHead.svelte`)

Props: `glyph?`, `tone?`, `lead?`, `label?`, `title`, `titleSnippet?`, `titleId?`, `meta?`, `action?`, `aside?`, `headingLevel?`, `compact?`, `testid?`. It is used on its own where a card's body is bespoke (the account cards). The existing calls are unchanged.

### IconTile

`size?: 36 | 32 | 28`. The glyph is 18, 16 or 14.

### card.ts

- `CARD = 'card-surface'`
- `CARD_QUIET = 'card-surface-quiet'` (it was `CARD_SURFACE`)
- `cardClass(quiet)`
- `CARD_HEAD`, `cardHeadTileSize(compact)`, `cardHeadTitleClass(compact, labelled)`
- `CARD_BODY[body]`
- `CARD_FOOT`
- `type CardBody`

### TabStrip (`TabStrip.svelte`, `TabStrip.astro`; recipe `lib/ui/tabs.ts`)

Svelte props:

- `tabs: readonly TabItem[]`, where `TabItem` is `{ id; label; href?; disabled?; suffix?; testid?; attrs?: Record<string,string> }`
- `active: string`
- `label: string` (the accessible name)
- `onSelect?: (id: string) => void`
- `idPrefix?: string`
- `testid?: string`

The two modes:

- **With `onSelect`:** a `role="tablist"` (horizontal) of `role="tab"`s with `aria-selected` and a roving tabindex (only the current tab is `tabindex=0`). ArrowLeft and ArrowRight wrap and skip disabled tabs, Home and End jump to the ends, and each selects as it moves (automatic activation, focus follows). A tab with an `href` stays an `<a>` (a click selects in place); otherwise it is a `<button>`. A disabled tab has `aria-disabled` and is out of the tab order. With `idPrefix`, the tab is `${idPrefix}-tab-${id}` and has `aria-controls="${idPrefix}-panel-${id}"`, and the consumer gives its panel that id plus `aria-labelledby`.
- **Without `onSelect`:** a `<nav aria-label>` of links, with the current one `aria-current="page"`.

Shared behaviour:

- `testid` names the strip element itself (the scroller).
- The fades are `${testid}-fade-start` and `${testid}-fade-end` (`tab-strip-fade-*` without a testid). They are 24 px, use `from-bg`, and are driven by `lib/scroll-fade.ts`.
- The current tab scrolls into view in the strip, never moving the page.
- The style is the report's underline: 44 px, 12 px 700 `.06em` caps, a 2 px gold rule on the current tab, a 1 px `--color-line-soft` base line, and a gold focus ring.
- `TabStrip.astro` is the link-mode twin (props `tabs`, `active`, `label`, `testid`).
- The helpers `tabIndexForKey`, `tabIdForKey`, `tabClass`, `scrollTabIntoView` and `tabStripFadeTestids` are exported.

### Glyphs

S4 adds `scale`, `bolt`, `list`, `pin`, `sword`, `tree`, `book`, `upload`, `download`, `share`, `layers`, `target`, `trophy`, `search` and `news` to `lib/ui/glyphs.ts`.

### CharacterHeader

The `.ch` surface is now the `card-surface` class. Its regions, states, API and heights are unchanged.

## What moved in this stage

- /account: `AccountTitle` → `PageHead`. "Your ratings" moved from the last Svelte `StatePanel` to `Card` (glyph `trophy`). `ui/StatePanel.svelte` and its test are deleted. Nothing else visible changed.
- S1b: `RatingCard`, `ForeverScoreCard`, `CoachPanel`, `GuildStepCard`, `CheckoutGuildStep`, `PremiumOwnedPanel` and `PlanCard` take the account surface. The guild plan cards are quiet.
- The four `.panel-box` copies are gone.
  - The /bis band body:
    - This set, Stat weights and Play it are compact Cards (`bolt`, `scale` with the band label as meta, `sword` with the band label as meta);
    - The list is a Card with a `table` body (`list`, meta "Hover or tap an item for its stats");
    - Where to get it is a Card (`pin`) with full-bleed group rows;
    - New at this level range is a Card (`star`).
  - The guide rail: Build, Rotation (meta = band) and Stat priority are compact Cards (`tree`, `sword`, `target`).
  - The leveling strip takes `CARD_QUIET`.
  - Copy: `statWeightsAtBand` became `statWeightsTitle`, `playItHeading` became `playItTitle`, and `rotationRailLabel` became `rotationRailTitle`. The band label moved to the meta.
- `Panel.astro`: /setup (three Cards: `user`, `download`, `device`; the titles are unchanged for now) and /ratings (seven h3 Cards) moved to `Card.astro`. Its only remaining user is `index.astro` (home lane), and a header note says so.
- Tabs: `GuildTabs` (in-page tabs on the guild page, link tabs on a raid page) and `SimTabs.astro` (link tabs, inside the phone gutter) draw `TabStrip`.
- /setup's phone capture runs to the card edge below 640 (the 20 px Card body took 4 px from the phone crop, under launch-polish's 11 px sub-line floor).

## Tests

- Unit: new suites for `PageHead` (Svelte and Astro twin), `Card` (both twins: surfaces, compact head, body kinds, lead, aside, foot, ActionLink, h3, decorative tile), `TabStrip` (keys, wrap and disabled-skip, roving tabindex, aria-selected, aria-controls, link mode with aria-current, the Astro twin, fades, attrs), `IconTile` 28, and glyphs. The pins that asserted the old surfaces (`PlanCard`, `RatingCard`, `CoachPanel`, `CharacterRatingPanel`, `GuildRaid` tab roles) were updated, each for the reason in its commit. The full web vitest suite passes: 661 files, 7552 tests, before the last two commits. The touched suites were re-run after them.
- Playwright on a fixture build (built outside `web/`, served on port 4517), desktop and mobile, 73 specs: the account, bis, guides, guild, sim tabs and phone, premium, ratings, checkout, coach, character, current-character, setup hand-offs, layout, planner, logs, rankings and header specs. Result: 2193 passed, 131 skipped, 2 failed. The 2 failures were the /setup phone picture floor; after the fix, launch-polish and handoffs-setup pass 44 of 44. The new `tests/e2e/tab-strip.spec.ts` (keyboard on the guild tabs, link semantics on the sim tabs) passes 8 of 8.
- Layout-reserve pins: none moved. The `PageHead` box equals the old account head, the `CharacterHeader` keeps its 1 px border and min-heights, and the account reserve specs pass.
- Captures: `design/reviews/2026-10-29-site-alignment/stage1-*.png` (generator `design/mocks/2026-10-29-site-alignment/capture-stage1.spec.ts`). /account at 1440 and 390 is pixel-identical to `before-account-*` (a PIL difference with an empty bbox). The 2000 capture is added.

## Concerns for the per-route lanes

- `Panel.astro` lives on for `index.astro` only. The home lane moves it onto `Card` and deletes the file.
- The guild tabs now draw `TabStrip`, but the guild head still sits on its navy and faction-art band until the guild lane removes it. The strip's fade runs to `--bg`, so on that band it reads faintly.
- /setup's step titles keep their copy ("1. Sign in with Battle.net"). The spec's numbered tile and "SIGN IN" titles are the /setup lane's.
- The `text` body is `py-4` (16 px), as /account ships. The spec's §3.4.1 says 16/20/16 bottom 20. The account wins as the reference; change `CARD_BODY.text` only with the owner's word.
- The report's `ModeBar` already is the underline recipe and was not moved onto `TabStrip` (its More menu and mode row are bespoke). The BiS band pills become a `SegmentedControl` per the spec, not tabs.
- The fixture sync replaced `web/public/data` (the other lane also builds on the fixture).
