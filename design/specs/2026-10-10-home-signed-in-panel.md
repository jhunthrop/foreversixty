# Home page, signed-in panel without the Switch character column: experience spec

Author: ux-designer. 2026-10-10, revised the same day after the WoW player review (`design/reviews/2026-10-10-home-signed-in-panel-player-review.md`, SHIP with fixes; section 11 lists every change). Supersedes the signed-in hero regions of the lost home rebuild spec (2026-09-30) §3.B.1, §3.B.2 and §3.B.4. §3.B.3 ("Your upgrades") and §3.B.5 ("Another class") keep their numbering and are unchanged except for where they now sit on the page. Mock: `design/mocks/gen_home_panel.py` (run `python3 design/mocks/gen_home_panel.py [name ...]`; the renders directory is gitignored). Renders in `design/mocks/renders/`: `home-panel` (1440, alts), `home-panel-one` (1440, one character), `home-panel-2000`, `home-panel-1024`, `home-panel-phone` (390), `home-panel-before` (today, for the side-by-side), `home-panel-states`. The header in every render is the shipped bar from `gen_nav.py`. Every character is an example (invented names on real class and spec pairs); every item, source, gain, set DPS and talent-point total is read from the active build's BiS file and item table at render time.

Owner decision 2026-10-10: the header character selector (`design/specs/2026-10-09-nav-character-selector.md`) is the one way to change character. The home page has no switch control of its own.

References: `docs/tenets.md` (14 first, then 4, 9, 11, 13), `design/DESIGN-SYSTEM.md` (the home hero h1 exception: 34 px 700, the character's name in class colour signed in; one crest language; "Updated" stamp only), nav spec §1 and §3.B (the full-name rule: the name is never cut), `web/src/components/HomeAccountPanel.svelte`, `character/HomeHeroCards.svelte`, `character/HomeUpgradesPanel.svelte`, `web/src/pages/index.astro`, `web/src/lib/home/*`.

Reference surface: the **in-game character frame** (crest, name in class colour, "Level 27 Human Fury Warrior" under it) for the identity block, and the **Details! title bar** habit of naming the current segment once and leaving the picker in one place. Not borrowed: a second picker on the page. The next-action cards and the upgrade table are existing site components and are unchanged.

## 1. The one question and the one job

**"Which character is this page about, and what do I do next to make it better?"**

Visible with no scroll:

| | 360 / 390 px phone | 1280 / 1440 px desktop |
|---|---|---|
| Which character | crest 64 px, full name 34 px in class colour (wraps to two lines, never cut), level race spec class, faction, guild, realm | crest 84 px, full name 34 px on one line, same facts |
| What next | the three cards, Best in slot first: "11 UPGRADES, in your band, 20 to 29 · +11.2 DPS together" with the Talents and Simulator cards under it, all inside the first screen of 390 x 844 | the three cards sit beside the identity; the first six upgrade rows are inside 900 px (at 1440 the table starts at y 527, today y 593) |

The job: tell a player who has several characters which one the whole page is about, then lead with the gear gap. Changing character is not this panel's job.

## 2. Decisions, with a recommendation for each

1. **What takes the column's place: nothing that switches, and nothing new. The hero re-flows into one band.** At 1280 and up the identity block takes 5 of 12 columns and the three next-action cards take 7 of 12, bottom-aligned (`items-end`, the grid the hero already uses). Below 1280 the identity stacks over the cards. The hero is the same stack as before, so the cards are the same 230 px wide at 1440 and nothing in them is re-set. Measured on the mock: the hero band shrinks from 272 px to 205 px of content (the 205 includes the reserved two-line name, decision 2), the "Your upgrades" table moves up 66 px (y 593 to y 527 at 1440), and no half-empty sky is left on the right.
   - **Rejected, leave 5/12 empty ("the panel narrows"):** at 1440 that is a 500 px hole in the sky with the answer still 593 px down; at 2000 it reads as a failed load.
   - **Rejected, widen the upgrades panel:** it is already the full 1248 px content width. Wider is not a thing; moving it up is, and the re-flow does that.
   - **Rejected, a second row of next actions (Tier list, Guide):** "Another class" already carries the Tier list link two sections down, a fourth card breaks the Best in slot / Talents / Simulator trio, and a card that does not improve this character dilutes the one primary (tenet 11, tenet 14).
   - **Rejected, a guild card:** the guild is already in the descriptor and links to its page; a card would restate it (one job per page).
   - **Gone with the column:** "Add one" (the selector's "Paste an export" row and `/setup` do that job, always present) and the per-alt "N upgrades" stat (see section 8, open question 1).
2. **How the hero names the character** (all of it, in this order, all already in `/v1/me`):
   - Eyebrow `Current character` when the account has two or more characters, `Your character` when it has one. The word "current" is the nav's own word for the pointer, so an alt player connects the page to the header selector.
   - The crest (84 px, circular ringed, class ring at 55 %, the one crest language) and the **full name as the page's only h1, 34 px Cinzel 700 in class colour**. The shipped hero uses 22 px; the design system says the home hero h1 is 34 px, and the signed-out h1 is 34 px in the same grid cell, so the swap moves nothing. Two-word names are the rule in Forever: the name wraps to a second line. **The name box reserves two lines (75 px) always, with the name sitting on its bottom edge, and the h1 is capped at two lines, then an ellipsis** (`line-clamp: 2`; the full name stays in the `title` and the accessible name). A header switch from a one-line to a two-line name therefore moves nothing, at the cost of 37 px of quiet space above a one-line name. A name that needs a third line is not a real Forever name.
   - Descriptor, two lines on purpose (the column is 500 px, so a one-line shape would wrap by accident): line 1 `Level 27 Human Fury Warrior`; line 2 the real faction emblem at 16 px (the flat Alliance lion and Horde crest, `factionLogoSrc` in `lib/faction-mark.ts`, `/icons/hd/faction/{faction}-logo-512.webp`, width and height 16; not the 72 px unit-frame shield and disc, which read as a blob and a generic shield at that size) and the faction word in faction colour, `· <Guild>` (a link to the guild page), `· Living Flame-US`. The realm stays because two alts can share a name across realms.
   - Sync line, 12 px muted: `12 minutes ago from the addon · gear and talents in sync`.
   - **One quiet line, only with two or more characters: `Change character: top right`.** 12 px muted, plain text under the sync line, not a link, not a button, no hover, no focus stop. "Top right" is where the selector is at every width (the bar's right end on desktop, the crest beside the Menu button on a phone), so it needs no word the player has to decode. It exists because returning players will look for the column that used to be here. It is not repeated anywhere else on the page. Revisit after two weeks of live use; delete it if nobody misses the column.
   - Not added: a "Current" pill (the eyebrow says it), a gold ring (gold marks the eyebrow only), a "not you?" link (nothing on this page switches).
3. **Signed-out panel: unchanged.** Its copy ("Pick your class for its best-in-slot list, talents and rotation, or sign in and your own characters arrive with their gear, talents and guild.") names no switching, the right cell is the Example panel, and the sign-in button and class picker stay. No edit in `index.astro`'s signed-out block, `homeHeroCopy` or `homeHeroStatePanelCopy`.
4. **Phone order (below 768, signed in): header; eyebrow; identity (crest 64 px beside the name); the three cards one-up (Best in slot, Talents, Simulator); Your upgrades; the dates timeline; Five doors; Another class.** The answer ("11 upgrades, +11.2 DPS together") is inside the first screen at 390 x 844 and is the first card under the identity at 360 x 640. The dates timeline is reference, so on a phone it sits under the table that answers the question, not between the answer and its detail (the table now starts about 640 px down instead of 1071). **Signed out, and signed in at 768 and up, the strip stays directly under the hero.** Implementation: a CSS-only reorder keyed on the existing `html[data-session='1']` hint below 768 (for example the content wrapper becomes `display: contents` there and the strip takes `order` after `#upgrades`), so a signed-out page and the pre-paint state do not move. Passed dates are muted at every width (decision 5).
5. **Passed dates are muted at every width:** the whole cell at 50 % opacity and its title line in `--muted`, not `--strong` (today only the opacity, 60 %, and the value stays bright). The next date keeps the gold dot; future dates are unchanged.

## 3. Anatomy by region

Tokens (all existing): `--bg #07090d`, `--raised #0d111a`, `--line #262e40`, `--line-soft #1c2230`, `--gold #e5b955`, `--gold-deep #a8762a`, `--muted #9a9484`, `--nav #b9b3a4`, `--strong #f2eee4`, body `#c9c2b2`, `--kill #7fd48a`, class colours from `tokens.css`, rarity text variants (`#3d94f0` rare, `#b866f5` epic). Type: Cinzel 700 for the name and card figures, Barlow for text, JetBrains Mono for the sync age, set DPS and gains. Components reused, none invented: `SkyBand`, `CharacterIdentity` (size `xl`, `heading`), `CharacterPortrait` (circular ringed crest), the three cards in `HomeHeroCards`, `Skeleton`, `UpgradeRow`, `ItemHover`, `HomeTimeline`. No new shared component, so no design-system review is owed.

Page container: `max-w 1344` including padding, gutter 48 px from 768 up, 18 px below. At 1440 the content is 1248 px wide; at 2000 the 1344 container is centred and the sky band runs full-bleed.

### 3.A Hero band (inside `SkyBand`, padding 32 px top and bottom, 28 on phone)

| Width | Layout |
|---|---|
| 1280 and up | 12-column grid, gap 32 px, `items-end`. Identity: `col-span-5`. Cards: `col-span-7`, three equal columns, gap 12 px (230 px each at 1440, 218 px at 1280). |
| 768 to 1279 | One column, gap 22 px: identity, then the three cards (three equal columns from 640 up; one-up below 640). Crest 84 px at 1024 and up, 64 px below. |
| below 768 | One column, gap 16 px, cards one-up. |

The signed-in island takes the whole hero grid. `index.astro` gives the island cell `lg:col-span-12` under the existing pre-paint `html[data-session='1']` hint and removes the right cell (`home-example-block`, `home-switch-character-slot`) from the grid under the same hint, so the Example panel does not leave a hole. If the hint was wrong (an expired session), the island's existing recovery that un-hides the signed-out block also restores the two-cell grid; one shift in that rare case is accepted. The `html[data-session='1'] .home-switch-character-slot { min-height: 168px }` rule in `global.css` is deleted.

### 3.B Identity (primary element of the hero; the h1)

Order top to bottom: eyebrow (11 px 700, .14em, gold, with the 28 px rule), 16 px gap, then a row: crest, 18 px gap (14 px on phone), text column with 2 px gaps: h1, descriptor line 1 (15 px `#c9c2b2`), descriptor line 2 (14 px `#c9c2b2`, emblem 16 px), sync line (12 px muted, age in mono), hint (12 px muted, alts only), and in the rare no-build case the Blizzard sentence (13 px muted) in place of the sync line. A rating figure line, when one exists, keeps its place after the sync line, unchanged. Stale sync (more than 24 hours) adds `Reopen the addon to refresh your gear.` under the sync line, no alarm colour.

States (nothing here is interactive except the guild link):

| Element | Default | Hover | Focus-visible | Active | Disabled | Loading | Empty | Error |
|---|---|---|---|---|---|---|---|---|
| Guild link | `--nav`, no underline | `--strong`, underline | 2 px gold outline, 2 px offset | `--strong` | n/a | n/a | absent when the character has no guild | n/a |
| Name (h1) | class colour, not a link | none | none | none | n/a | skeleton, see below | n/a | n/a |
| Hint line | muted text | none (it is not a control) | none | none | n/a | n/a | absent with one character | n/a |
| Whole block | as above | n/a | n/a | n/a | n/a | session hinted, `/v1/me` pending: the session snapshot paints it when one exists, else a `Skeleton` sized to the block (84 px ring, a 240 x 34 name inside the same 75 px reserved name box, 300 x 15 and 200 x 12 lines) | `/v1/me` answers with no characters: the island renders nothing, the signed-out block un-hides (existing behaviour) | `/v1/me` fails: same fall-back to the signed-out block |

Character changed from the header: the page does not reload and the hero keeps its height (the two-line name box). The identity swaps in place from `/v1/me` data already held (no skeleton), the cards and the table show their own skeletons at their reserved heights and fade in once with the existing 160 ms `.reveal`. The selector's live region already says "Now using {Name}."; this panel adds no second announcement.

### 3.C Next-action cards (secondary to the name, equal to each other)

Unchanged from `HomeHeroCards` except the Simulator card's copy (below): three links, each card `p 14/16`, radius 6, 1 px `--line`, gradient `#131824` to `#0d111a`, 1 px gold hairline along the bottom, label 11 px gold, the answer as a 22 px Cinzel uppercase figure, one 12 px muted line. Equal weight on purpose: they are three doors. The figure may wrap to two lines (a long simulator headline); the grid stretches the three cards to one height.

**Simulator card, two figures, both with their unit.** Figure: `{n} DPS now`, the character's latest saved sim, in the figure style. Line: `{x.x} DPS at band best in slot`, the age-free comparison figure in mono. **What the second figure compares against:** the simulator's DPS for the full best-in-slot set of this character's own spec, band and faction (the band entry's `set_dps`, run with the band's own talents, race and preset); the first figure is the player's own run. The two are only comparable when the saved sim used the same preset as the band entry, so the line keeps its `title` naming the preset (`presetLabelFor`, as shipped) and the build lane must confirm that rule before this ships (tenet 8). The unit word follows the band (DPS, HPS or score). Also: the "Run" label stays plain text inside the link.

| State | Default | Hover | Focus-visible | Active | Disabled | Loading | Empty | Error |
|---|---|---|---|---|---|---|---|---|
| Card (link) | as above | border `--gold-deep`, 120 ms | 2 px gold outline, 2 px offset | border `--gold-deep`, fill `#0f131d` to `#0a0d15` | n/a (a card with nothing to say still links) | three 16 px `Skeleton` lines (Best in slot, Talents), two 12 px lines (Simulator): the ready height | the honest-gap sentences in section 4 | Simulator fetch failed: `No sim yet.` |

Targets: Best in slot `#upgrades` (the table below, now higher on the page), Talents `/planner?spec={spec}` (bare `/planner` with no spec), Simulator the character's armory sim. All 44 px or taller (cards are 118 px).

### 3.D Your upgrades (§3.B.3, unchanged) and the strip above it

The dates timeline strip stays between the sky band and the table on desktop. The table is the existing `HomeUpgradesPanel`: heading `Your upgrades` (Cinzel 18 px, .10em), the "Full list for {Spec} {band}" link in gold 12 px 700 uppercase, a raised panel with a Slot / You wear / Best in slot / Gain header, `UpgradeRow`s (icon 36 px with rarity border, name in rarity colour, one source line under each item, mono gain in `--kill`), and the "Already best in slot" footer. **Worn rows drop the "you wear this" sub-line** (the column header already says "You wear"): the sub-line is the worn item's own source in the same 12 px muted style when known (a world drop, a vendor, a crafting profession and "crafted"), and nothing when not, in which case the name is vertically centred beside its icon. The `youWearThis` copy key is deleted. **Below 768 the "Full list for {Spec} {band}" link sits on its own line under "Your upgrades"** (4 px gap, 12 px, no wrapping beside the heading). All its states (loading skeleton, no spec, no gear export, no list yet, all match) are unchanged; two are drawn in `home-panel-states`. The reserved height of its slot stays `min-h-[220px]`.

### 3.E Removed

`HomeSwitchCharacterPanel.svelte` (and its test), the `home-switch-character-slot` div in `index.astro`, the dynamic import and `switchTo`/`onswitch` wiring in `HomeAccountPanel.svelte`, and the copy keys `switchCharacterLabel`, `addOneCharacter` (`home-panel-copy.ts`) and `switchUpgradesStat` (`current-character-copy.ts`). `switchCurrentMarker` stays (the sim landing list uses it); `switchAction` has no other user and goes. `homeHero.switchTo` is deleted from `createHomeHero` if nothing else calls it; the hero still follows the pointer through `CURRENT_CHARACTER_CHANGED`, which the header selector already fires. Side effect worth the lane's note: the home page no longer fetches and parses one 4 to 6 MB class item table per alt class to count upgrades.

## 4. Copy, verbatim

| Where | Text |
|---|---|
| Eyebrow, two or more characters | `Current character` |
| Eyebrow, one character | `Your character` |
| Hint (two or more characters only) | `Change character: top right` |
| Descriptor line 1 | `Level {n} {Race} {Spec} {Class}` (any part the API did not send is left out, never a placeholder) |
| Descriptor line 2 | `{Faction}` · `<{Guild}>` · `{Realm}-{REGION}` |
| Sync line | `{relative time} from the addon · gear and talents in sync` |
| Stale sync, added line | `Reopen the addon to refresh your gear.` |
| No build, Blizzard has nothing | `Blizzard serves no data for this realm type yet.` |
| Best in slot card | label `Best in slot`; figure `{n} upgrade(s)`; line `in your band, {band} · +{x.x} DPS together` (`, one slot not sim-checked` or `, {n} slots not sim-checked` appended when it applies; the unit word is the band's own: DPS, HPS or score) |
| Best in slot, nothing to gain | figure `Best in slot`; line `every slot matches the {band} list` |
| Talents card | label `Talents`; figure `Optimized` or `Unoptimized`; line `{n} of the {p} points in the {band} build differ · compare in the planner` |
| Simulator card | label `Simulator`; figure `{n} DPS now`; line `{x.x} DPS at band best in slot` (unit word per band: DPS, HPS, score). No sim: `No sim yet.` and `Run` |
| No spec (Best in slot, Talents) | `Pick a spec` / `Set a spec in the planner to see this.` |
| No gear export | `Not available yet: no gear export for this character. Open the addon once to send it.` |
| No talent export | `Not available yet: no talent export for this character. Open the addon once to send it.` |
| No list published | `No best in slot list published yet for this spec and band.` / `No talent build published yet for this spec and band.` |
| Table heading, link, headers | `Your upgrades`, `Full list for {Spec} {band}`, `Slot`, `You wear`, `Best in slot`, `Gain` |
| Worn row sub-line | the worn item's source line, or nothing (no "you wear this") |
| Table footer | `Already best in slot:` then up to three items, then `{n} more slot(s)` |

`{band}` reads `20 to 29` and plain `60` (never "60 to 60").

## 5. Hierarchy

Identity block: the name is the one primary element (34 px, class colour, the only large type on the page above the fold). Everything under it is secondary and visibly so: descriptor in body colour, sync and hint in 12 px muted. Cards region: one level of emphasis, the gold labels and 22 px figures; the three are peers. Gold appears only on the eyebrow, the card labels and the table's link; the hint is the quietest text on the page. Passed dates are the quietest of all. Whitespace: 32 px sky padding, 16 px between eyebrow and identity, 22 px stacked gap, 32 px band to strip.

## 6. Phone and desktop behaviour

- **Never collapses:** the full name (it wraps, it never cuts), the crest, the Best in slot figure and its line, the sync line, the guild link.
- **Collapses:** the right-hand cards drop under the identity below 1280; one-up below 640; the table header row hides below 768 and each row stacks (worn item, down arrow, best item, gain) as shipped.
- Hint line: same string at every width. "Top right" is the selector's place on both bars (the crest by the Menu button on a phone).
- 2000 px: the 1344 container is centred, the sky is full-bleed, nothing stretches (`home-panel-2000`).
- Touch and keyboard parity: the only controls in the panel are the three cards and the guild link; tab order is guild link, Best in slot, Talents, Simulator, then the table. No hover-only behaviour. 44 px targets throughout.

## 7. Performance and polish limits

- No layout shift on hydration. Reserved heights, measured on the mock (implementer re-measures and pins them in the test): hero content 205 px at 1280 and up (eyebrow, the two-line name box, identity with the hint; the 118 px cards sit inside it), about 345 px at 768 to 1279, about 560 px on a 390 phone. The island root keeps `[grid-area:1/1]` under the signed-out block so signed-out to signed-in swaps in the same cell.
- The h1 is 34 px in both states, so the swap does not resize the heading.
- The crest `<img>` carries width and height; no missing-icon flash (the class crest is a local WebP). Upgrade icons are `loading=lazy` with a fixed 36 px box.
- Signed-out LCP is unchanged (the hero font subset and the signed-out h1 are untouched). Signed-in, the island still mounts `client:idle-after-load`; removing the Switch panel removes one dynamic import and N item-table fetches. `web/lighthouserc.json` budgets stay as they are; the lane runs the `index.html` shard.
- Reduced motion turns the 160 ms reveal off, as everywhere.

## 8. What it is NOT, and open questions

Not: a character switcher of any kind (no list, no button, no link that sets the pointer); a per-alt comparison; an account page (no battletag, no manage links); a place that explains the selector (one 12 px line, once); a marketing hero; a provenance header (the sync line is the one freshness fact, no pills).

Open questions for the owner:
1. **Alt upgrade counts.** The column was the only place a player saw "which alt needs gear most" at a glance. The selector deliberately shows no gear. Recommendation: put the same mono "N upgrades" stat on the `/account` character rows (the design system's character-row stat form already allows it), as a separate lane. `DESIGN-SYSTEM.md` still says the signed-in home's Switch character panel uses the stat form; that sentence needs editing by whoever owns the file (outside this task's files).
2. **Footer edge pinning at 2000 (site-wide, outside this spec).** At 2000 the footer's logo and legal line pin to the far viewport edges while every page's content is centred in the 1344 container. The footer should use the same container. Follow-up for whoever owns `Footer.astro`; no change here.
3. **Other switch controls.** `web/src/components/sim/LandingState.svelte` lists characters with a per-row "Sim it" button and a "Current" marker (not a pointer switch, but the same shape), and `CurrentCharacterBar.svelte` historically carried a Switch popover (nav spec open question 4). Not touched here; audit both against the owner's rule.

## 9. Acceptance screenshots the implementer must attach

Built page, `/v1/me` stubbed (as `design/mocks/shoot_nav_built.mjs` does), same example characters:

1. 1440 x 900 viewport and full page, signed in with four characters (compare to `home-panel.png` side by side, then to `home-panel-before.png`: the Switch column is gone, the table starts near y 527).
2. 1440 signed in with one character: eyebrow `Your character`, no hint line (`home-panel-one.png`).
3. 2000: container centred, sky full-bleed (`home-panel-2000.png`).
4. 1280 (side-by-side still holds, cards 218 px, no figure clipped) and 1279 (stacked).
5. 1024 (`home-panel-1024.png`).
6. 390 x 844 full page and 360 x 640 first screen: the Best in slot card is on screen; the timeline sits under the table; "Full list for ..." is on its own line (`home-panel-phone.png`). Also 768 and 1024 to prove the strip stays under the hero there.
7. States: a two-word name of 24 characters wrapping at 34 px, and a header switch between a one-word and a two-word name with the table's y position unchanged (attach both captures); a name that would need a third line clamped with an ellipsis; sync older than 24 hours; no build and Blizzard has no data; no spec; session hinted before `/v1/me` answers (skeleton, then the character, no shift; attach a CLS reading of 0 for the swap); a character switched from the header (cards and table skeleton then reveal); card hover, focus-visible and active; Best in slot loading, all-match, no gear export; Simulator no sim yet.
8. Signed out at 1440 and 390, proving it is unchanged (the Example panel still in the right cell, no hole when the hint is absent).
9. A grep result showing no `Switch character`, `Add one`, `home-switch-character` or `HomeSwitchCharacterPanel` left under `web/src` or `web/tests`; the e2e cases at `web/tests/e2e/home-panel.spec.ts` lines near 301 and 418 rewritten to assert the column is absent and that selecting another character in the header changes the h1.

## 10. Checklist for the review

- [ ] No control on the page changes the current character.
- [ ] Eyebrow, name (34 px, full, class colour), crest 84/64 px, descriptor in two lines, sync line, as in section 3.B.
- [ ] Hint line only with two or more characters, plain text, once.
- [ ] Cards 230 px at 1440, equal height, all states in 3.C.
- [ ] Table starts higher than today's y 593 at 1440; first six rows inside 900 px.
- [ ] Name box reserves two lines; a one-to-two-line switch moves nothing.
- [ ] Hint reads `Change character: top right`, alts only.
- [ ] Descriptor emblems are the 16 px lion and Horde crest.
- [ ] Simulator card reads `{n} DPS now` and `{x.x} DPS at band best in slot`.
- [ ] Worn rows carry a source or nothing; no "you wear this".
- [ ] Phone: timeline under the table; passed dates muted at every width.
- [ ] No source pills, no battletag, no extra gold.
- [ ] Signed-out panel pixel-identical to today.
- [ ] No layout shift at hydration or on a header switch; CLS 0.

## 11. Changes in the 2026-10-10 revision (player review)

1. Name box reserves two lines and clamps at two (decision 2, 3.B, section 7). Hero content 168 px becomes 205 px; the table gain over today is 66 px, not 103.
2. Simulator card: `{n} DPS now` and `{x.x} DPS at band best in slot`, with what the second figure compares against and a build-lane check on presets (3.C).
3. Phone, signed in: timeline below the table; passed dates muted at every width (decisions 4 and 5).
4. Hint is `Change character: top right`.
5. Descriptor emblems are the real 16 px faction logos. `DESIGN-SYSTEM.md` still says the row-descriptor emblem is the unit-frame shield and disc; that line needs the owner of the file to update it, and the nav's 14 px silhouette marks are a separate, deliberate variant.
6. Phone: "Full list for ..." on its own line under the heading.
7. Footer edge pinning at 2000 noted as a site-wide follow-up (section 8, question 2).
8. Worn rows: source or nothing, no "you wear this".

Data the player doubted, for the accuracy loop and not decided here: a level-24 hunter at 105 DPS against a level-27 warrior at 26 DPS (pet included?); Defias leather over mail for a warrior at 27; a melee weapon swap worth +0.6 DPS for a Beast Mastery hunter. The mock reads these from the build as it stands.
