# Forever Sixty BiS landing: experience spec

Author: ux-designer. 2026-10-10, **round 3 after the player's round-two review (SHIP WITH FIXES) in `design/reviews/2026-10-10-bis-landing-player-review.md`; round 2 followed the NOT YET review and the coordinator's twelve decisions. Section 12 lists every change.** Route `/bis` (`web/src/pages/bis/index.astro`, copy `web/src/lib/bis/copy.ts`). Mock: `design/mocks/gen_bis_landing.py`; renders `design/mocks/renders/bis-landing{,-signed-out,-2000,-1024,-768,-phone,-phone-signed-out,-360,-states,-before}.png`.

References: `docs/tenets.md` (14 first, then 11, 10, 9, 8, 7, 4), the one-job-per-page rule (BiS owns gear, planner owns talents, sim owns numbers), `design/DESIGN-SYSTEM.md` (principles 1, 2, 4; the `BandTabs`, class tile, class page header and character row entries), `design/specs/2026-10-09-tier-list.md` (the band header, the 1344 column, `Your spec` treatment, the 1% and tier rules), `design/specs/2026-10-09-nav-character-selector.md` (the header selector is the only way to change character; ember means stale), `design/specs/2026-10-10-home-signed-in-panel.md` (the word "upgrades", the descriptor line), `web/src/components/bis/BisClassCard.astro` (what this replaces), `web/src/lib/bis/hover.ts` (`bandForLevel`, `bisPageHref`: `/bis/{class}/{spec}?faction=…#band-{faction}-{band}`).

## 1. The one question and the one job

**"What should I wear right now?"** The page gets a player from the front door to their own spec's list at their own band, in one click, and says nothing else about gear. It does not rank specs (tier list), show a spec's picks (the spec page), or show talents or numbers.

Visible with no scroll:

| | 390 x 844 phone | 1440 x 900 desktop |
|---|---|---|
| Signed in, character synced | Eyebrow, h1, job line, tier pointer, Updated, the whole lead (identity, `6 of 15 slots`, button) | Hero, the lead, the heading, the strip and the first row of class cards |
| Signed in, no addon sync | Same, the count replaced by one sentence and one link | Same |
| Signed out, or no character | Hero, heading, the strip; the first card starts on screen | Hero, heading, the strip and the first row of three class cards (Warrior, Paladin, Hunter) in full, above the fold |

Today's page opens with a plain "BiS" heading, a paragraph that explains what a band is, and nine identical stacked cards. `bis-landing-before.png` is that page as built; the answer to the question is on it 28 times as a DPS number and nowhere as an action.

## 2. Decisions, with a recommendation for each

### 2.1 The hero: one line, the pointer, the stamp, no pitch

**Recommendation.** The tier list's band: eyebrow `Levels 20 to 60`, h1 `Best in slot` (Cinzel 30 px, 26 on phone), one 14 px job line in the body colour, **one 14 px line under it with the tier list pointer** (`Which spec is strongest? Tier list →`, round 2: the player review found it tucked at 13 px beside the grid heading), and the one `Updated Oct 10` stamp (12 px muted) on the h1's baseline at the right edge (under the pointer on phone). The job line is the question answered and the instruction in one breath: `The best gear you can wear at your level, for every spec. Open yours.`

- The "what is a band" paragraph leaves the hero and becomes the strip's two lines (2.3). No slogan ("Play your class better." is the home page's line). The stamp is the newest `generated_at` across every spec file; no build, engine or pills in the header.
- Rejected: an h1 that asks the question (reads as a pitch, duplicates the job line).

### 2.2 The class grid, not a list; and no per-spec figure

**Recommendation: a 3 x 3 grid of class cards, and no figure on any spec row.**

- **Why a grid.** Nine classes make a 3 x 3 at the 1344 column (card 416 px at 1440). A stacked list of nine full-width rows (today) puts 900 px of empty space between a spec name and its number; a grid lets the eye find a crest and a colour, which is how a player finds their class. The order is the game's (Warrior, Paladin, Hunter, Rogue, Priest, Shaman, Mage, Warlock, Druid), fixed: it never reorders for the signed-in character.
- **Card anatomy.** A card (6 px radius, 1 px `--border`) with a 3 px class-colour stroke on the top edge and a 7% class-colour wash fading out by 120 px. Header: the circular ringed `ClassCrest` at 56 px (44 on phone), the class name in Cinzel 18 px in the class colour. Then one **48 px row per spec** (44 on phone): the spec's **tab icon at 28 px** (round 2: the same icon the planner's `TreeTabs` shows, `specs.json` `icon`; round 1 used each tree's first talent icon, which players do not know as the spec), the spec name in Barlow 600 15 px `--strong`, a role word on the right only for tanks and healers, and an arrow that appears on hover. A class is never pictured with a square class icon; the spec icon is a spell icon, as in the planner and class header tabs.
- **Cards size to content and a grid row aligns by top** (`align-items: start`; no `stretch`, no min-height). Round 2 and 3: Mage and Warlock end where their third row ends; Druid's card is 48 px taller because it has four specs. No filler, no stretched empty strip.
- **Role word.** `Tank` or `Healer` (11 px uppercase, muted) because Feral Bear sits beside Feral and two specs are called Restoration and two Holy. DPS is the default and carries no word (tenet 11).
- **The per-spec figure is removed, all 28.** (Today there are 28, not 27: Druid has four.) (1) The tier list owns rankings and gives the figure its meaning (units, gap, ties, the raid-ready conditions); a bare "Level 60: 778.2 DPS" invites the comparison without the caveats (tenet 7). (2) It answers the wrong question for a level 34 player. (3) Three units (DPS, HPS, tank health and damage) cannot sit in one column. (4) 28 equal-weight mono strings are what make the page flat (tenet 11). No figure stays anywhere on the cards; the one number on the page is the lead's slot count (2.4). The only defensible alternative, if the owner wants a number on rows, is a worn-slot count on the character's own row, never a DPS.

### 2.3 The band story: one compact row, the grid's own header

**Recommendation: one compact row between the lead and the cards, visually part of the grid (the heading sits above both), built from the five existing `BandTabs` pills. It tells where the spec links open and explains bands once.**

- **Anatomy (900 and up).** A `panel` row, 14 px padding, two zones. Left (440 px): `Spec links open at this band` (13 px 600, `--strong`) and under it the note at **13 px, two lines at 1440** (`Each list is the best you can wear inside one band; gear changes about every ten levels. 60: full raid-ready, Phase 1 gear.`). Right: the five pills `20 to 29`, `30 to 39`, `40 to 49`, `50 to 59`, `60` (mono 13 px, 44 px tall, equal width; the top band is plain `60`, never "60 to 60"). **Every pill has a 12 px caption** (round 3, so none is bare): `Top sources: {kind}, {kind}.` in the body colour where the band's top two source kinds differ from the band before it (the first band always, as the baseline), and `Same sources as {previous band}` in the muted colour where they match. On the 1.60.1.70291 build: `20 to 29` `Top sources: dungeons, crafting.`; `30 to 39` `Same sources as 20 to 29`; `40 to 49` `Same sources as 30 to 39`; `50 to 59` `Top sources: vendors, dungeons.`; `60` `Top sources: vendors, quests.`
- **What it does, said in every state.** The selected pill sets the `#band-{faction}-{band}` hash on every spec link on the page, so Open Fury lands on Fury at that band. The row says so in its own head line, `Spec links open at this band`, signed in, signed out and while loading. It changes nothing else.
- **What it does not do (binding): the lead's button and count never change with the strip.** The lead is the character's own band, always. Picking `50 to 59` retargets the 28 spec rows and nothing above them; that is why the strip sits below the lead, with the grid it controls, not above it.
- **One gold highlight.** The selected pill is the only gold element in the row (gold border, gold text, 8% tint). The character's own band is marked **inside its pill**, never as a badge: at 480 px and up with the small word `you` (10 px, 75% of the pill's own colour, on the label's line); **below 480 px with a 7 px gold dot in the pill's top-right corner** (round 3: the word wrapped at 390 and 360). When the player selects another pill, the character's pill keeps its mark and stays neutral. The dot is the one gold mark besides the selected pill, and it is the same gold.
- **Default selection.** Signed in: the character's band (`bandForLevel(level)`; levels under 20 clamp to `20 to 29`). **Signed out, or no character: `20 to 29`** (round 2, binding: launch players are leveling, and a level 27 warrior who opens Arms must not land on the 60 list; 60 itself is a correct list, just not the first one a new player needs). The choice is kept in the URL as `?band=40` and, per viewer, in `localStorage` in try/catch.
- **What `60` means.** The site's 60 band is the full raid-ready, Phase 1 list (the spec page's and the tier list's own preset, and what the data holds), so the strip's note says so: `60: full raid-ready, Phase 1 gear.` **Not done, with the reason:** the round-two review and the coordinator asked for `(Molten Core, Onyxia)` after it. The band-60 raid-ready files hold no pick or runner-up sourced from either (grep of all 28 files: none), so naming them would state a fact the data does not carry (tenet 8). If the owner wants the raids named, the BiS pipeline has to source raid items first, or someone supplies a primary source that Phase 1 here means those two; then the parenthesis is one string. The leveling bands (20 to 29 through 50 to 59) are the ones with no raid drops among their picks; the note does not say that, because the pill captions already show each band's sources.
- **Stacked form (below 900).** The same panel, vertical: head line, the five pills in one 5-column row (**44 px tall, label on one line**, mono 12 px at -0.04em tracking, 3 px gaps, 10 px panel padding below 480 so the 58 px cell at 360 holds the 54 px label), one 13 px caption line for the **selected** band (a top-sources caption or `Same sources as …`), then the note. At 480 to 899 the pill carries the `you` word on the label's line; below 480 the gold dot.
- **Why not a bigger ladder.** A richer step graphic would be a new control (tenet 10). Equal-width stretching of `BandTabs` is a layout variant for the design-system owner (open question 4).

### 2.4 Signed in: the "your character" lead

**Recommendation: one full-width panel above the grid with the current character, how many of its best-in-slot slots it wears, and one primary action.**

- **Container.** `panel` on the neutral raised ground with a 1 px gold border (round 2: **no gold tint behind it**, so the class-colour name reads), 108 px min height, padding 18 x 22, three columns `1.15fr | 1fr | auto`, gap 28. Full 1344 column; stacks below 1024.
- **Identity.** The circular ringed crest at 64 px (56 below 1024) and, beside it, the label `Your character` (gold) and the character's name in Cinzel 22 px in the class colour. **The crest and the name are one link** to the same page as the button. Under them (indented to the name): `{Spec} {Class} · Level {n} · {faction mark} {Faction}` at 13 px, where the faction mark is the **flat Alliance lion or Horde crest at 16 px** (`factionLogoSrc`, the `{faction}-logo-512.webp` files), with 6 px gaps and a muted dot on each side (round 2: the stray gap and the grey shield are fixed), then the 12 px muted sync line `Addon · 12 minutes ago`.
- **The count.** `label` `Best-in-slot gear you wear`, the figure `6 of 15 slots` (JetBrains Mono 26 px, the `of 15 slots` 15 px muted), a 6 px gold progress bar, and **one 12 px muted line whenever slots are excluded**: `Trinkets have no pick at this band.` The 15 is the number of the 17 gear slots that have a sourced pick for the character's spec at its band (the same list the spec page shows); the other two, named in the note from the data, are the slots the sim has no pick for at this band (trinkets, in the 40 to 49 files). Round 2 rule: **"slots", never "picks", everywhere on the page.** At 60 all 17 slots have picks and there is no note. The coordinator's example "trinkets open at 41" is not used: the data says only that the 40 to 49 files hold no trinket pick; a level at which trinkets open is not in the BiS data and would be unverified (tenet 8).
- **The action.** The one primary gold button, **48 px**, naming the spec and the band: `Fury gear for 40 to 49 →` (`Fury gear for 60 →` at 60). It opens `/bis/{class}/{spec}?faction={faction}#band-{faction}-{band}`: the character's own spec, band and faction, and it never changes when the strip's selection changes. Filled gold deviates from the design system (the one filled gold button is the sign-in on the signed-out home, "never for a tool or a link"); this is the coordinator's decision for round 2 and needs the design-system owner to widen the rule to "one filled gold primary per page" (open question 2).
- **It follows the header selector in place** (`CURRENT_CHARACTER_CHANGED`): the lead, the `you` word, the selected pill (when it was on the character's band), the `Yours` pill and the heading re-render. No switcher on this page, no reload.
- **A character under 20.** The middle column is one sentence, `Best in slot starts at 20.`, with no count and no bar; the button reads `Fury gear for 20 to 29 →` (the first band) and the strip selects `20 to 29` with no `you` word (the character is not in a band yet).
- **The character's own spec row** in the grid gets the tier list's row treatment (2 px gold left edge, 6% gold tint) and a gold `Yours` pill in place of any role word. The heading reads `Another class or spec`.
- **Rejected:** a "your upgrades" table here (the home panel and the spec page own it); a DPS gain in the lead; a progress ring; a switcher.

### 2.5 Phone

One column: eyebrow, h1, job line, tier pointer, Updated, the lead (crest 56 beside the identity, the count, a full-width 48 px button), the heading, the strip (stacked form, 2.3), nine cards (crest 44, 44 px rows). Signed out the lead is gone and the heading and the strip start the page's content: `bis-landing-phone-signed-out.png`. 360 is rendered (`bis-landing-360.png`). Nine stacked cards are about 3 screens at 390, not the 2.0 the player review asked for: it would take disclosure or two-up cards, and tenet 4 forbids hiding spec rows; the page stays at nine full cards (open question 5).

## 3. Layout, region by region (1440)

Tokens (all existing): `--bg #07090d`, `--raised #0d111a`, `--line #262e40`, `--line-soft #1c2230`, `--gold #e5b955`, `--muted #9a9484`, body `#c9c2b2`, `--strong #f2eee4`, ember `#d66e28` for stale words only. Cinzel 700 for the h1 (30 px), the character name (22 px), class names (18 px) and the grid heading (18 px, uppercase, .10em); Barlow elsewhere; JetBrains Mono for every number and band label. Reused: the band, `ClassCrest` 44/56/64, `FactionMark` (flat logo, 16 px), `UpdatedStamp` wording, `panel`, `label`, `btn`, the `BandTabs` pill, the 6 px progress bar, `Skeleton`, the tier list's `Your row` treatment.

| # | Region | Height at 1440 | Notes |
|---|---|---|---|
| 1 | Header band (nav, hero) | 81 + about 150 | eyebrow, h1, job line, tier pointer, stamp |
| 2 | Lead (signed in only) | 108 | neutral panel, gold border |
| 3 | Grid heading | 44 | h2 |
| 4 | Strip | 110 | the grid's header row |
| 5 | Cards | rows of 236 (Druid's card 284), gap 16, aligned by top | |

**Signed out at 1440 x 900, the first row of cards is on screen above the fold** (round 2, binding): the nav and hero end at about 240, the heading and strip at about 400, and Warrior, Paladin and Hunter run from 415 to 650 (`bis-landing-signed-out.png`). Signed in, the lead takes 108 px more and the first row still ends at about 830. Signing in does not reflow the grid beyond the lead's height, which the layout reserves while the character resolves.

## 4. States

Every interactive element, rendered in `bis-landing-states.png`. **Hover never moves anything** (round 2: the round-1 board lifted the row; the board and this table now agree): hover is a background and colour change and an underline, never a translate or a size change.

| Element | Default | Hover | Focus-visible | Active | Disabled | Loading | Empty | Error |
|---|---|---|---|---|---|---|---|---|
| Spec row (link) | as 2.2 | `#101624` background, 2 px class-colour left edge, name underlined (4 px offset), arrow lights gold; nothing moves | 2 px gold outline, inset | `#0a0d15` ground, name white and underlined | n/a (all 28 pages exist) | n/a (static) | n/a | n/a |
| Spec row, the character's | gold edge, 6% tint, `Yours` | as above, gold edge kept | as above | as above | n/a | n/a | n/a | n/a |
| Band pill | `--line` border, `#b9b3a4` mono | text `--strong`, border `--line-warm-strong` | 2 px gold outline, 2 px offset | `#0a0d15` ground, text white | n/a | n/a | n/a | n/a |
| Band pill, selected | gold border, gold text, 8% tint, `aria-current="true"` | border and text `#f5d27a`, 14% tint | same ring | 4% tint | n/a | n/a | n/a | n/a |
| Band pill, the character's | as its state plus the word `you` inside (a gold corner dot below 480 px) | as its state | as its state | as its state | n/a | n/a | n/a | n/a |
| Lead button | filled gold gradient (`#f0cc6c` to `#c99a3a`), dark text, 48 px | lighter gradient | 2 px gold outline, 2 px offset | darker gradient | n/a | skeleton block, same size | n/a | n/a |
| Lead identity link (crest and name) | name in class colour | name underlined | 2 px gold outline, 4 px offset | n/a | n/a | skeleton | omitted | omitted |
| Tier list link | gold 14 px 700 | underline, `#f5d27a` | gold outline | n/a | n/a | n/a | n/a | n/a |

- **Mock scaffolding never renders.** The example character (Obnoxious Yell, `6 of 15 slots`, the sync ages) is invented for the boards; round 1 and 2 printed an "Example character… invented for this board" footnote under the grid. Round 3 removes it from every render. Nothing of the kind exists on the live page, in any state; the fact that the character is an example is stated here and in the acceptance notes (section 11) only.
- **Addon synced** (reference): `bis-landing.png`.
- **No addon sync.** The character is known (Battle.net, a pasted export, no gear on file). The sync line goes; the count column is one sentence and one 44 px gold text link: `Add the addon to see how many of these slots you already wear.` / `Get the addon →` (to `/setup`). The button still works. No `0 of 15`: an unknown is not a zero (tenet 8).
- **Stale sync.** The age is plain and the words past the limit go ember (nav spec 2.11: more than three times the player's own median gap, clamped 3 to 14 days; 19 days in the board): `Addon · 19 days ago · Log out or /reload in the game; the companion then syncs it`. **Wording checked against the addon:** the export is written to SavedVariables on logout or `/reload` (`addon/README.md`, `Locale.lua` `recorderStopped`), and the companion app, not the addon, reads that file and uploads it (about every ten minutes, `docs/superpowers/plans/2026-09-20-addon.md`). So the coordinator's "then sync in the addon" is replaced by `the companion then syncs it`. The count label becomes `Gear you wore at the last sync`; the figure is not ember; the name and the crest ring are never recoloured. The header selector carries the ring mark.
- **Level 60.** Band plain `60`, `Level 60` in the descriptor, all 17 slots, no note, `Fury gear for 60 →`.
- **Under 20.** See 2.4; rendered.
- **Pasted export.** Shown from the pointer like a synced character without a sync line; with gear in the export the count shows, otherwise the no-sync sentence.
- **Error or no character.** The lead is omitted and the page is the signed-out page (the strip falls to `20 to 29`); nothing says "failed". The grid and strip never wait for the character.
- **Keyboard.** Tab order: the lead identity link, the lead button, the five band pills (arrow keys move between them, a `role="radiogroup"`; each is a link to `?band=` for no-script), then each card's spec rows top to bottom. The tier list link comes first, in the hero. Card headers are not links. **Touch parity:** every target is 44 px or more (spec rows 44 on phone, pills 44, buttons 48), no state depends on hover.

## 5. Phone, tablet, wide

- **Phone (390 and 360).** See 2.5. The pills are 44 px tall with the label on one line (mono 12 px at -0.04em, 3 px gaps); the character's band is a gold corner dot, not the word, so nothing wraps at 390 or 360.
- **Tablet, 600 to 1023.** Phone bar (the nav's 56 px bar below 1024), gutter 32, the lead stacked (crest, identity, count, full-width button), the strip stacked, cards in **two columns** (Druid alone in the last row, left). `bis-landing-768.png`.
- **1024 to 1439.** Three columns (card 298 px at 1024), the strip's two zones with captions wrapping to two lines, `bis-landing-1024.png` (the mock's nav is the shared stand-in squeezed to fit; the shipped nav at 1024 is sized by the nav spec).
- **2000.** Everything is inside the 1344 column; cards keep 416 px; the nav content aligns to the same edge. `bis-landing-2000.png`.

## 6. Copy (verbatim; section 8 says which words are data)

| Where | Text |
|---|---|
| Page title | `BiS` (unchanged, `bisCopy.indexTitle`) |
| Eyebrow | `Levels 20 to 60` |
| h1 | `Best in slot` |
| Job line | `The best gear you can wear at your level, for every spec. Open yours.` |
| Tier pointer | `Which spec is strongest?` + link `Tier list →` |
| Stamp | `Updated {Mon d}` |
| Lead label | `Your character` |
| Lead descriptor | `{Spec} {Class} · Level {n} · {Faction}` (flat faction mark before the faction name) |
| Sync line | `Addon · {age}` (stale: `Addon · {age} · Log out or /reload in the game; the companion then syncs it`, the part after the last dot in ember) |
| Count label | `Best-in-slot gear you wear` (stale: `Gear you wore at the last sync`) |
| Count | `{n} of {total} slots` |
| Excluded slots note | `{Slots} have no pick at this band.` (`Trinkets have no pick at this band.`) |
| No sync | `Add the addon to see how many of these slots you already wear.` / `Get the addon →` |
| Under 20 | `Best in slot starts at 20.` |
| Lead button | `{Spec} gear for {band} →` (`Fury gear for 40 to 49 →`, `Fury gear for 60 →`, `Fury gear for 20 to 29 →`) |
| Grid heading | signed out `Pick a class and spec`; signed in `Another class or spec` |
| Strip head | `Spec links open at this band` |
| Strip note | `Each list is the best you can wear inside one band; gear changes about every ten levels. 60: full raid-ready, Phase 1 gear.` |
| Pills | `20 to 29`, `30 to 39`, `40 to 49`, `50 to 59`, `60` |
| Pill caption | `Top sources: {kind}, {kind}.` where they differ from the band before (kinds: `dungeons`, `crafting`, `vendors`, `quests`, `world drops`, `reputation rewards`, `PvP rewards`); `Same sources as {previous band}` where they match |
| The character's band | `you` inside the pill at 480 px and up; a 7 px gold dot in the pill's corner below 480 |
| Role words | `Tank`, `Healer` |
| Your spec row | `Yours` |
| Document description | `The best gear you can wear at every level band, from the simulator: one list per class and spec, built once and never per character.` (unchanged) |
| Gone | `indexIntro`, `tierListPointer`, `indexSpecDps60`, `tankCopy.indexSpecSummary` on this page, and the `noDataYet` row text (a spec with no ranked list keeps its row; its own page carries the empty state) |

## 7. Hierarchy

Primary element: signed in, the lead (the only gold-bordered block, the only filled gold button, the only 22 px name); signed out, the grid heading and the first row of cards. One level of emphasis below: the strip's selected pill. Everything else is quiet: the stamp is 12 px muted, role words are 11 px muted, the strip has one gold element (the selected pill) and nothing in the grid is gold except `Yours` and the arrow on hover. The page has one number (`6 of 15 slots`) where today it has 28.

## 8. Data and build notes (for the lane)

- **No figure is typed in copy.** Each pill caption is computed at build time from the BiS files (prose never quotes nightly numbers): the two most common `source_kind` values across all written specs' slots at that band (rows with no source are not counted), shown only when they differ from the band before (`band_story()` in `design/mocks/gen_bis_landing.py`). The strip's note carries no count.
- **The count.** `n of total slots` needs the character's worn gear at the band's slots from `/v1/me` (addon sync). `total` is the number of the spec's slots with a sourced pick at the character's band and faction (15 at 40 to 49, 17 at 60 for Fury); `n` is those whose item id is worn. The excluded slots are the ones with no sourced pick; the note names them from the data (`trinket1` and `trinket2` read as `Trinkets`). A count never mixes bands: a level 47 character is counted against the 40 to 49 list. The 6 and the 14 in the mock are invented.
- **Band for a character:** `bandForLevel(level)` (clamped 20 to 60). Level under 20 is the under-20 state, not a count.
- **Faction** for the link and band comes from the character; signed out the links carry no `?faction=` and the spec page's default stands.
- **What stays prerendered:** the hero, the heading, the strip (selected `20 to 29`), the grid. The lead is an island that reserves 108 px; the pill selection and the hash on spec links are applied on load from `?band=`, the pointer and `localStorage`, with no layout shift (links only change their `href`).
- **Spec icons:** `specs.json` `icon` per spec (what `TreeTabs.svelte` reads), served from the build's icon tree.
- **Removed work:** `BisClassCard.astro`'s DPS and tank-summary props, `band60Dps`, `band60TankSummary`, `band60PresetLabel`, `unit` and the loads that feed them in `index.astro`, `indexSpecDps60`, `tankIndexSummaryOf`'s use on this page, and the data-testids `bis-index-dps-*`. Page tests move from "a figure per spec" to "28 spec links, no figure, the strip, the lead". (`tankCopy.indexSpecSummary` and `rateUnitOf` are used elsewhere; leave them unless grep shows otherwise.)
- **The tier list note** in tier-list spec section 3.3 (the BiS index headlining effective health for tanks) is resolved by this page: the index no longer ranks tanks at all.

## 9. What this page is NOT

- Not a ranking: no DPS, no tiers, no sort order but the game's.
- Not a gear page: no picks, no sources of items, no stat weights. Source kinds appear only as the strip's two words per band.
- Not a character switcher: the header selector is the only one.
- Not a pitch: no feature list, no "why Forever Sixty".
- Not a class page: the class name is a heading, not a link; there is no `/bis/{class}` index.

## 10. Open questions (for the wow-player review, round 2)

1. **Dropping the per-spec figure entirely** (2.2). The risk is that some players read the cards as emptier.
2. **A filled gold button on a content page.** The design system reserves the one filled gold button for the signed-out home's sign-in. The lead's `Fury gear for 40 to 49 →` is a second; the rule should become one filled gold primary per page, or the lead goes back to the outlined button. Binding for round 2 by the coordinator; the design-system owner should record it.
3. **Equal-width `BandTabs`** (2.3). Same pill and states, stretched to a column so a caption can sit under it.
4. **Excluded slots.** The count excludes the slots the sim has no pick for at a band and says which, with no level at which they open (unverified). If the player-review wants "Trinkets open at 41", someone must supply the source.
5. **Phone length.** About three screens of nine cards at 390; the only shorter forms (two-up cards, accordions) hide spec rows. Not done, for tenet 4.

## 11. Acceptance screenshots the implementer must attach

(The character in these captures is the mock's example; do not copy its name, counts or ages.) `/bis` at 1440 signed in (synced), signed out, with no addon sync, with stale sync, at 2000 and at 390 and 360; hover, focus-visible and active on a spec row, a band pill and the lead button; a Druid card (four rows); the strip with a non-default pill selected and a spec link's `href` inspected; the lead resolving (skeleton to content, no shift); and a side-by-side with `bis-landing.png` at 1440 plus 1024 and 1920 captures (the owner's rule: full-bleed header, centred container). Tests: 28 spec links, no `DPS` or `HPS` or `effective health` string on the page, `Updated` appears once, no provenance pills, no square class icon (grep the old portrait paths), the lead follows the header selector without a reload, and an end-to-end test of the hover promise.

## 12. Changes in round 2

1. Signed-out default band is `20 to 29`; the 60 band is described as full raid-ready, Phase 1 gear, matching the spec page and the tier list (player findings 1 and 12; the signed-out default is `20 to 29` because launch players are leveling, not because 60 is wrong).
2. Band captions only where the top sources change (finding 2).
3. The strip is one compact row, the grid's own header, below the lead, with `Spec links open at this band` in every state; one gold highlight; the character's band carries the word `you` inside its pill; the lead never changes with the strip (findings 3 and 4).
4. Signed out, the first row of cards is above the fold at 1440 (finding 4).
5. Hover no longer lifts; active, selected-pill hover, and the under-20 lead added (finding 5).
6. Stale copy is `Log out or /reload in the game; the companion then syncs it`, checked against the addon (finding 6).
7. `6 of 15 slots` with a note for excluded slots; "picks" is gone (finding 7).
8. The lead's action is a 48 px filled gold button naming the spec and band; the crest and name link to the same page (finding 8).
9. Flat faction mark at 16 px with fixed spacing; the lead panel is neutral with a gold border (finding 9).
10. Phone pills are 44 px, mono 12 px, `you` inside; signed-out phone, 360, 768 and 1024 rendered (finding 10).
11. Planner tab icons; cards size to content and align by top; the tier pointer is 14 px under the job line (finding 11).

## 13. Changes in round 3

1. The example-character footnote is gone from every render; the spec states that nothing of the kind renders live (section 4, section 11).
2. Phone pills (below 480 px): 44 px, label on one line, the character's band is a small gold corner dot; the `you` word stays at 480 and up (rendered at 390 and 360, and in the states sheet).
3. No bare pill: `Same sources as {previous band}` in the muted caption style where the sources match the band before (2.3).
4. Mage and Warlock: cards size to their three rows and the grid aligns by top (`align-items: start`), stated in 2.2 (already how the round-two render drew it; now said).
5. The strip's note is 13 px and two lines at 1440 (the left zone is 440 px, the note reordered: `Each list is the best you can wear inside one band; gear changes about every ten levels.` then the 60 sentence).
6. The 60 note stays `60: full raid-ready, Phase 1 gear.` without `(Molten Core, Onyxia)`: the data carries no pick from either (2.3).
