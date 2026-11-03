# Experience spec: class landing `/guides/<class>`

Date: 2026-11-02. Author: ux-designer. Before captures: design/reviews/2026-11-02-class-landing/before-*.png.

Scope: `web/src/pages/guides/[class]/index.astro`, `web/src/components/guides/SpecCard.astro`, the nine `web/src/content/guides/<class>/index.md` files, and the card-line helper. I make no change to `Content.astro` and no new mode. The page keeps `hideHeader wide`. The only shared-component edit is the one-line guard on `ClassHeader` (section 6).

---

## 1. The one question

**"Which of this class's specs is mine, and where is its guide?"**

Answer without scrolling:

| Viewport | Visible without scrolling |
|---|---|
| 360 x 640 | Nav (56), eyebrow, crest, h1 with the class name, 2 to 3 lines of identity, and spec card 1 in full with its "Read the guide →" row. |
| 390 x 844 | The same, plus spec card 2 in full (it ends near y 774). Card 3 shows from about y 786. |
| 1280 x 720 | Header (ends y 224), then every spec card in full (3 cards end near y 466, 4 cards near y 510). The h2 "What changed from 1.12" starts at about y 506 for 3 cards. |

Rules:
- The page has no h1 chips row.
- It has no figure.
- A visitor never scrolls to find a spec.

## 2. Reference surfaces

| Borrowed from | What we take |
|---|---|
| In-game talent frame / class picker | Three big spec tiles side by side: icon, name, role. The tile is the button. |
| Icy Veins class landing | Spec choice first. A "what changed" list below. Nothing else above the choice. |
| Wowhead class overview | Facts as short ruled rows, not paragraphs. Races in a side box. |
| Item tooltip | Label above name, name in colour, one body line. This is the card's three-line structure. |
| Site's own `/guides` index | The `ClassHeader` minimal head and `Read the guide`. |

## 3. Defects in the before captures

**Named by owner and coordinator (restated with captures):**
1. The h1 "HUNTER IN FOREVER" is set in the class colour, and the same colour is on the card names. That is two emphases.
   - Seen: hunter lime, shaman blue, paladin pink, druid orange. Priest is white, so the chrome differs per class.
   - The spec chips (hunter-2000 y 217 to 261, shaman-2000, priest-1440) repeat the card navigation.
2. The order is identity paragraph, cards, then a bullet list that restates the cards.
   - Seen: hunter-2000 y 600 to 745, shaman-2000, priest-1440 y 585 to 740.
3. The card body is the guide's SEO description, for example "How to play a Beast Mastery Hunter in Forever: talents for leveling and raids...". The same lead is repeated on all three cards, and it is 13px.
4. The cards print "658.7 DPS at 60, this set", and "375 damage taken per second" on a tank. This is a nightly number, and the sim owns numbers.
5. Prose is a 600px left column against a 1248px card row.
   - hunter-2000: x 368 to 970 is text and x 970 to 1591 is empty for about 900px of height.
   - paladin-1024: text is 640px of 928px.
   - shaman-2000 is the same.
6. Unbroken blocks, each with no subheading.
   - Warrior: one 8-sentence changes paragraph.
   - Priest: four bold-led paragraphs, about 1,100 characters of races alone.
   - Mage: a 3-sentence identity paragraph plus a 2-paragraph races block.
   - Shaman: 3 paragraphs before the bullets.
7. "READ THE GUIDE" is a 12px underlined gold link floating centred under each card.
   - The card itself does nothing.
   - The underline comes from `.prose-forever a`, so it is not a design choice.
8. Druid-390.
   - The chip row clips the fourth chip ("Restoration" cut at the right edge) with no scroll affordance.
   - The chip icons are `loading="lazy" fetchpriority="low"`, so they paint as empty dark boxes.
   - Feral and Feral Bear both take the tree's first-talent icon, so two cards can carry the same picture. The implementer must verify this.
   - Bullet lists have no markers, because Tailwind preflight resets `list-style` and `.prose-forever ul` only restores padding.

**Additional defects I found:**
9. No gap between the cards and the next block. It is 12px (hunter-2000, y 597 to 609), 15px (paladin-1024) and 12px (shaman-2000). The paragraph or list butts the card shadow.
10. The vertical rhythm is inverted. The identity sits 27px above the cards but 53px below the chips. It groups with the cards, not the header.
11. Lists look broken. Items have a 25px pitch equal to the line height, a 20px indent and no marker. Hunter-2000's 6 changes read as one wall. Shaman-2000's 5 bullets read as a paragraph. Paladin-1024 is the same.
12. Spec-by-spec sentences are inlined in 7 identities.
    - Shaman: three bold sentences in paragraph 1.
    - Rogue, priest and mage: bold or bullet sentences repeating the cards.
    - Hunter, warlock, warrior and druid: the 3 or 4 spec bullets.
13. Navigation sentences that point at the page. Examples: "The Beast Mastery, Marksmanship and Survival guides below cover...", "Each spec guide above covers...", "Pick a spec below...". The heading and cards already say this, and the paladin one says "above" for a block that is below.
14. The card name repeats "Hunter" on every card ("Beast Mastery Hunter" x3) under an h1 that already says Hunter.
15. The card is `<div>` plus a nested link. The padding area does nothing on click and there is no hover state.
16. `SpecCard` hand-writes `.panel` (surface, border, radius, shadow). That duplicates the `card-surface-quiet` utility.
17. The "Your spec" pill is injected into a 13px label row at hydration. Its 22px height shifts the card by about 9px.
18. Labels carry the wrong heading vocabulary: "What changed from Classic:", "...for every druid:", "What changes for every rogue in Forever:", "What else changed from 1.12:", "in short:", "in the parts you will feel". They are plain paragraphs, not headings, and no two classes match.
19. Priest and mage carry caveats ("The beta stops at level 30...") at body weight, as if they were a section.
20. 2000px. The 1344 column is centred, but the content is left-flush inside it and nothing balances the right side (defect 5).

## 4. Target layout

### 4.1 Page order (all widths)

1. Nav (unchanged).
2. Head: eyebrow, crest, h1 (class name), identity line.
3. Spec cards.
4. Body: "What changed from 1.12", then "Races", then "In a raid" (only where content exists).
5. Optional note (priest only), confidence line, Sources, footer (unchanged).

### 4.2 Head (`ArtPanel` + `ClassHeader minimal`)

Props:
- `specTabs=[]`.
- `titleOverride={className}`, which gives h1 text "Hunter".
- `summaryOverride={identityText}`.
- `eyebrow="Guides"`.

`<title>`, meta description and JSON-LD stay `entry.data.title` ("Hunter in Forever").

**Required `ClassHeader` edit (one line):** render the `.class-header-spec-tabs` row only when `specTabs.length > 0`. Otherwise it leaves an empty div plus a 16px flex gap. No other caller passes an empty array.

Tokens:

| Item | Spec |
|---|---|
| Eyebrow | unchanged (existing `.label` gold plus 28px rule) |
| Crest | `ClassCrest size=64 priority`, eager |
| h1 | `heading-page`, **colour `var(--color-strong)`**, not the class colour. Set via an h1-colour prop, or via a page-scoped `:global(.class-header-h1){color:var(--color-strong)}`. 34px at <768, 44px at >=768. |
| Identity | existing `.class-header-summary`: 14/20, `--color-nav`, max-width 62ch |
| Class colour | only the crest ring and the card names |
| Padding | `ClassHeader` keeps its own: 22px top below 1024, 28px top from 1024 |
| Below the head | 24px (the cards' `margin-top`) |

The crest stacks above the h1 below 1024. That is `ClassHeader`'s existing rule, kept for consistency with every other guide page. It costs 74px at 390, which section 1 already budgets.

Identity rule:
- It is the first paragraph of `index.md`.
- It is plain text, at most 2 sentences and at most 300 characters.
- The page passes it as a string. Bold or links are build-time errors.

### 4.3 Spec cards grid

`.class-landing-cards` is a CSS grid, `margin: 24px 0 0`.

| Viewport | 3 specs | 4 specs (druid) | Gap | Card layout |
|---|---|---|---|---|
| 390 (<640) | 1 col, 354w | 1 col, 354w | 12 | Stacked |
| 768 (640-1023) | 1 col, 672w | 1 col, 672w | 16 | **Row** |
| 1024 | 3 cols, 298w each (content 928) | **2 x 2**, 456w each | 16 | Stacked |
| 1440 (content 1248) | 3 cols, 405w | **4 cols**, 300w | 16 | Stacked |
| 2000 (column capped, content 1248, 376 to 1624) | 3 cols, 405w | 4 cols, 300w | 16 | Stacked |

The 4-up switch is at >=1200. Between 1024 and 1199, druid is 2 x 2.

Grid rule for druid: `repeat(2,1fr)` from 1024, `repeat(4,1fr)` from 1200. For 3 cards: `repeat(3,1fr)` from 1024.

## 5. The card (`SpecCard.astro`, rebuilt)

### 5.1 Structure

```html
<a class="spec-card card-surface-quiet" href=... aria-label="Read the Beast Mastery Hunter guide"
   data-testid="guide-spec-card-beast-mastery">
  <img class="spec-card-icon" width=44 height=44 loading=eager decoding=async alt="">
  <span class="spec-card-text">
    <span class="label spec-card-role" data-spec-slug="beast-mastery">DPS [Your spec pill]</span>
    <span class="spec-card-name">Beast Mastery</span>
    <span class="spec-card-line">...</span>
  </span>
  <span class="spec-card-cta">Read the guide <span aria-hidden="true">→</span></span>
</a>
```

Use `card-surface-quiet`. Delete the local `.panel`.

Remove these props and code:
- `setMetric`, `setMetricLine`, `classLandingSetMetric` (and their imports).
- `description`, `guidesCopy.setDpsLineSuffix`, `guidesCopy.setMetricLine` (if unused elsewhere).
- The `className` suffix on the name.

All inner elements are `span`s. A `<p>` would pick up `.prose-forever p` margins and max-width. A nested link is not allowed.

### 5.2 Card line (source)

Build-time helper `specOverviewLine(body)`:
1. Take the first paragraph under `## Overview` in the spec guide.
2. Split it into sentences.
3. Drop any sentence that contains `](` (a link).
4. Join what remains.
5. Strip `*` and `_` emphasis markers to plain text.

The build fails if `## Overview` is missing, the result is empty, or it is longer than 200 characters. No fallback to `description`.

Warrior/fury.md: "Fury is the dual-wield warrior: Berserker Stance all fight, Bloodthirst and Whirlwind on cooldown, ...".

### 5.3 Card tokens

| Part | Spec |
|---|---|
| Surface | `card-surface-quiet`: `--surface-card`, 1px `--color-line-card`, `--radius-panel` 12px, `--shadow-card` |
| Padding (stacked) | 20px 22px at >=1024, 16px 18px below 640 |
| Stack gap | 12px between head, line, CTA |
| Icon | 44x44, radius `--radius-control` 8px, 1px `--color-line`, `background: var(--color-bg)` so a late image never shows an empty or broken box, `object-fit: cover` |
| Head | grid, 44px icon, 12px gap, then the role/name block |
| Role | `.label` (11px), `--color-muted`, `min-height: 22px`, `display:flex; align-items:center; gap:8px` |
| Pill | `pill pill-site`, 22px tall, sits in the role row |
| Name | `heading-name` 20px/1.1, **class colour** (`classes.json` color), sentence case as in the data ("Beast Mastery") |
| Line | 15px/22px `--color-text`, weight 400, no clamp, no truncation |
| CTA | 12px/700, uppercase, `--tracking-button` (.07em), `--color-gold`, `min-height: 44px`, `margin-top: auto`, left-aligned, no underline |
| Equal heights | the grid rows stretch; the CTA pins to the bottom |
| Underline | `text-decoration:none` on `.spec-card` and its children. Scoped CSS wins over `.prose-forever a`. |

Row layout, 640 to 1023 (3 and 4 cards alike):
- `grid-template-columns: 44px minmax(0,1fr) auto`.
- `align-items: center`, `column-gap: 16px`, padding 16px 20px.
- The icon is in column 1. The role, name and line stack in column 2. The CTA is in column 3 and right-aligned.
- The card is about 130 to 150px tall. It is still one link.

### 5.3a Card heights

| Width | Approx. card height |
|---|---|
| 390 | 228 |
| 768 | 150 |
| 1024 | 232 |
| 1440, 3-up | 218 |
| 1440, 4-up | 262 |

### 5.4 Hierarchy inside the card

1. **Primary:** the spec name in class colour, 20px.
2. The line, in body text.
3. **Secondary:** the role label (muted, small), then the CTA.
4. One accent colour apart from the class colour: gold on the CTA only.

## 6. The body

Wrapper `.class-landing-body`. All list, paragraph and heading styles are scoped to it. Spec guides keep `.prose-forever`.

The page splits the rendered markdown by `h2`:
- Everything before the first `##` is the identity.
- Each `##` becomes a section.
- Only these headings are allowed: `What changed from 1.12`, `Races`, `In a raid`. Any other heading fails the build.
- A trailing `---` rule followed by one paragraph is the note.

Section spacing:
- 40px above the body (32px at <640).
- 32px between sections.
- 40px between the body and the confidence line.

### 6.1 Section heading

`<h2 class="section-title">`:
- 18px.
- Barlow Condensed 800.
- Uppercase.
- `--color-strong`.
- `margin: 0 0 12px`.

An 18px secondary heading is deliberate. The cards are primary.

### 6.2 "What changed from 1.12"

A ruled fact list, not bullets.

| Part | Spec |
|---|---|
| `ul` | `list-style:none; padding:0; margin:0` |
| `li` | `position:relative; padding:12px 0 12px 20px; border-top:1px solid var(--color-line-soft)` |
| Marker | `li::before`, 6x6 circle, `--color-gold`, `left:0; top:20px`. It is visible at 390. |
| Last row | the last `li` has `border-bottom: 1px solid var(--color-line-soft)` |
| Text | 15px/22px `--color-text`. `strong` is `--color-strong` 700. |
| Columns at >=1024 | `column-count:2; column-gap:48px; column-rule:1px solid var(--color-line-soft)`. `li { break-inside: avoid }`. |
| Columns at <1024 | 1 column |
| Mage only | the first paragraph in the section is `.landing-lead`: 15/22 `--color-nav`, `margin: 0 0 12px`, `max-width: 68ch` |

### 6.3 "Races" and "In a raid"

A panel (`card-surface-quiet`, padding 20px 22px). The panel `h2` is the same `section-title`.

| Case | Layout at >=1024 | Layout at <1024 |
|---|---|---|
| Races only (6 classes) | One panel, full width. Items in `column-count:2`, gap 48. | One column |
| Races and In a raid (mage, warlock) | `grid-template-columns:1fr 1fr; gap:16px`. Races on the left, In a raid on the right, equal height. | Stacked, Races first, gap 16 |

- Items inside a panel are `li`s: 15/22, 12px gap, `border-top: 1px solid var(--color-line-soft)` except the first.
- A single-paragraph panel renders the paragraph without a list.
- Race and spell names stay bold from the source.
- Reference: the Wowhead class overview side box. This is a full-width band under the list because an aside beside a short list leaves an empty column at 1440 and 2000. A short Races panel under the list is only about 80px of void.

### 6.4 Note

The priest note renders directly above the confidence line:
- 13px.
- `--color-muted`.
- Same style as the confidence line, with 8px between them.

### 6.5 Centring and widths

- Every region sits inside `.page-column` (1344 max, 48px gutter).
- At 2000 the content is 1248px wide with 328px of margin on each side.
- No region sets its own max-width wider than the column. The identity keeps 62ch, left aligned inside the column.

## 7. Wireframes

### 7.1 390 (content 354, 18px inset)

```
[nav 56]
 ─ GUIDES                                  eyebrow 22 below nav top
 (crest 64)
 HUNTER                                     h1 34
 Hunter is Forever's physical damage        identity 14/20, ~3 lines
 class with a pet and a bow. ...
 [24]
 ┌─────────────────────────────┐ 228
 │ [44] DPS  [Your spec]       │
 │      Beast Mastery          │
 │ <line, 4-5 lines>           │
 │ READ THE GUIDE →            │
 └─────────────────────────────┘
 [12] card 2, card 3 ...
 [32]
 WHAT CHANGED FROM 1.12
 ─────────────────────
 • Pets are reworked. ...
 ─────────────────────
 • ...
 [32] RACES panel
 [40] confidence line, Sources, footer
```

### 7.2 1024 (content 928)

```
 ─ GUIDES
 [crest 64] HUNTER  44px
            identity (62ch)
 [24]
 ┌──────298──────┐ ┌──────298──────┐ ┌──────298──────┐   232 tall
 │ [44] DPS      │ │ [44] DPS      │ │ [44] DPS      │
 │  Beast Mastery│ │  Marksmanship │ │  Survival     │
 │  line (5 ln)  │ │               │ │               │
 │  READ THE ... →│ │               │ │               │
 └───────────────┘ └───────────────┘ └───────────────┘
 [40]
 WHAT CHANGED FROM 1.12
 • item          │ • item                      2 cols 440 / 48 / 440
 • item          │ • item
 [32]
 RACES panel, 928 wide
```

### 7.3 1440 (content 1248, x 96 to 1344)

2000 is the same with margins 328.

```
 ─ GUIDES
 [crest] HUNTER
         identity
 [24]
 ┌────405────┐ ┌────405────┐ ┌────405────┐       218 tall
 [40]
 WHAT CHANGED FROM 1.12
 col 1 (600)           │ col 2 (600)
 [32]
 RACES panel (1248)
```

For druid at 1440 the cards row is 4 x 300 and 262 tall.

## 8. States

**Card (the whole `<a>`):**

| State | Spec |
|---|---|
| Default | as 5.3 |
| Hover (`@media (hover:hover)`) | `translateY(-2px)`; border `--color-gold-deep`; background `linear-gradient(var(--row-hover),var(--row-hover)), var(--surface-card)`; CTA colour `--color-gold-hover`; arrow `translateX(3px)`; 140ms ease-out on transform, border-color and background |
| Focus-visible | `outline:2px solid var(--color-gold); outline-offset:3px` (follows the 12px radius). Same hover look. No transform on focus. |
| Active | `translateY(0)`; border `--color-gold`; background `--row-hover` x2 |
| Disabled | none. A spec without a guide is not a card. |
| Loading | none. Static HTML. The icon tile is already sized with a bg. |
| Empty | build fails if a class has zero cards |
| Error | the `/v1/me` read fails: no pill. Nothing else changes. |
| Touch | no hover transform (`hover:none`); active as above; target is the whole card, at least 88px |
| Keyboard | Tab order is the cards in tree order, then the Sources links. Enter follows the link. |
| Reduced motion | `prefers-reduced-motion`: no transform or transition |

**Your spec pill** (unchanged behaviour):
- It is injected into `[data-spec-slug]` after the `/v1/me` read when the visitor's character is on this class.
- It also sets the card `aria-label` to "Read the {Spec} {Class} guide, your spec".
- Role-row `min-height: 22px` makes injection shift-free.

**Lists and panels:** static. No interactive state. Links inside panel text keep `.prose-forever a` underline behaviour (these are none today).

## 9. Copy (verbatim)

| Element | Copy |
|---|---|
| Eyebrow | Guides |
| h1 | `{Class name}`, for example "Hunter" |
| Hidden h2 above cards (`sr-only`) | Specs |
| Role label | DPS / Healer / Tank |
| Pill | Your spec |
| CTA | Read the guide → |
| Card aria-label | Read the {Spec} {Class} guide |
| Card aria-label, your spec | Read the {Spec} {Class} guide, your spec |
| Section headings | What changed from 1.12 / Races / In a raid |
| Confidence footer and Sources | unchanged |
| Empty and error | none rendered |

No "pick a spec below" line. No new prose.

## 10. Content map (per class)

Terms:
- IDENTITY: the first paragraph of the file, plain text.
- CHANGES: `## What changed from 1.12`.
- RACES: `## Races`.
- RAID: `## In a raid`.
- DELETE: removed (the cards or headings carry it).
- Bold leads are kept. Sentences are not rewritten.

**Druid**
- IDENTITY: paragraph 1.
- DELETE: the 4 spec bullets and the label "What changed for every druid:".
- CHANGES: the 7 bullets, unchanged.
- RACES: the final paragraph ("Druids are Night Elf or Tauren...").

**Hunter**
- IDENTITY: paragraph 1.
- DELETE: the 3 spec bullets, the label "What changed from Classic:", and the last paragraph ("The Beast Mastery, Marksmanship and Survival guides below...").
- CHANGES: the 6 bullets.
- RACES: "**Human Hunter** is one of the six new race and class pairs...".

**Mage**
- IDENTITY: sentence 1 of paragraph 1 (ends "...every caster in the raid.").
- Sentence 2 ("Forever keeps all of that and rebuilds the three trees underneath, ...") moves to the top of CHANGES as `.landing-lead`.
- DELETE: "The three specs:", the 3 bullets, the label "What changed from 1.12, in short:", and the final sentence ("Pick the spec guide below...").
- CHANGES: the 5 bullets.
- RACES: the 3 sentences from the closing races paragraph, as 3 items:
  - "**Orc Mage** is new for Forever..."
  - "**Skyborne** can be mages too..."
  - "**Gnome** is still the strongest Alliance pick..."
- RAID: "All three are raid specs. Fire wants long fights and a healthy mana bar, Arcane puts the most damage into its Arcane Power windows, and Frost is the easiest to level and the safest in a bad pull."

**Paladin**
- IDENTITY: paragraph 1.
- DELETE: the label "What changed from 1.12, for every paladin:" and the final paragraph ("Each spec guide above covers...").
- CHANGES: the 7 bullets.
- RACES: "**Undead Paladin** is new in Forever..." (2 sentences, one paragraph).

**Priest**
- IDENTITY: paragraph 1 (both sentences).
- DELETE: the whole paragraph that starts "**Discipline** is the tank healer...", and the sentence "Each spec guide says which ones matter to it."
- CHANGES: the "**Every priest now has**" paragraph split into 4 items, one per existing sentence.
  - Item 1: "**Every priest now has** Fear Ward, ... Lightwell from the trainer."
  - Item 2: "Mind Flay is a Shadow talent instead of a trainer spell, ...".
  - Item 3: "Shadowfiend is not learnable at launch."
  - Item 4: "In Forever a heal learned below level 20 keeps its full healing power bonus, ...".
- RACES, in this order, with each group's existing bold lead kept as its label:
  - "**Race picks.**" split into 3 items (Healers: / Shadow: / "Dwarf's 1% hit, ...").
  - "**The Gnome priest** is new...", one item.
  - "**Each race has its own two priest spells.**", split by race into 5 items (Human, Dwarf, Night Elf, Undead, Troll).
- NOTE: "The beta stops at level 30. ..." goes after a final `---`.

**Rogue**
- IDENTITY: sentence 1 ("Rogue is still energy, combo points and stealth, but Forever moves all three trees onto daggers at 60 and gives each one a clear job.") plus the existing sentence "If you want one character for launch, level Combat and pick your raid tree at 60."
- DELETE: the 3 bold spec sentences, the label "What changes for every rogue in Forever:", and "Each spec guide below has the talents...".
- CHANGES: the 6 bullets.
- RACES: "Skyborne can be rogues on both factions."

**Shaman**
- IDENTITY: sentence 1 of paragraph 1 ("Shaman is the totem class: ... whatever their spec.").
- DELETE: the 3 bold spec sentences and the label "What else changed from 1.12:".
- CHANGES: 7 items, in this order.
  - Item 1: "**One weapon, never two.**" paragraph.
  - Item 2: "**Totems reach further.**" paragraph.
  - Items 3 to 7: the 5 bullets.
- RACES: "**Both factions get shamans now.**" paragraph.

**Warlock**
- IDENTITY: paragraph 1 (2 sentences).
- DELETE: "The three specs:", the 3 bullets, the label "What changed from 1.12:", and "Every spec guide below covers..." (the trailing sentence of the last paragraph).
- CHANGES: the 8 bullets, plus the sentence "Unstable Affliction and Haunt are in the client files, but no trainer teaches either one." as a 9th item.
- RACES: "Troll warlocks are new for the Horde, ... Skyborne cannot be warlocks."
- RAID: "In a raid you bring a curse when assigned (...), Healthstones for everyone, a Soulstone on a healer, and summons."

**Warrior**
- IDENTITY: paragraph 1 (2 sentences).
- DELETE: "The three specs:", the 3 bullets, the lead-in sentence "What changed from 1.12, in the parts you will feel.", and "Pick a spec below...".
- RACES: "Every one of the original eight races can be a warrior, and so can the Skyborne, on either faction."
- CHANGES: the changes paragraph split into 8 items, one per sentence:
  1. Tactical Mastery.
  2. Victory Rush.
  3. Thunder Clap.
  4. Shield Block.
  5. Sweeping Strikes.
  6. Shield Wall.
  7. Death Wish.
  8. "New talents include ...".
  - These items have no bold leads. The list renders the same.

**File shape after the edit:** identity paragraph, `## What changed from 1.12`, `## Races`, `## In a raid` (mage and warlock only), and optional `---` plus note (priest only).

## 11. Phone and desktop behaviour

| | Phone (<640) | 640 to 1023 | >=1024 |
|---|---|---|---|
| Collapses | Cards go 1 column; list and panels single column | Cards become full-width rows; list single column | Nothing |
| Never collapses | Identity, every card in full, CTA | Same | Same |
| Hidden | nothing | nothing | nothing |

There is no accordion. Nothing is truncated. At 390 the lists are long (a warlock has 9 items) and that is accepted, because the body is below the answer.

## 12. Performance and polish

- No client island. The pill script is unchanged.
- No layout shift on hydration:
  - Role row `min-height: 22px`.
  - Fixed icon `width/height` attributes (44).
  - `ClassCrest priority` (eager).
- All card icons are `loading="eager"`, `decoding="async"`, with explicit width and height. The `dataUrl(build, icons/<icon>.webp)` helper stays. A card whose icon resolves to the same file as another card's must fall back to the spec guide's own icon. This applies to Feral and Feral Bear. Verify before merge. No placeholder square.
- Spec chips are gone. That removes 3 to 4 lazy, low-priority icon requests above the fold.
- Web fonts are unchanged, so no new swap.
- Page must pass its existing `web/lighthouserc.json` entry. No new budgets. CLS must be at most 0.05.
- `.prose-forever` link underline does not apply to cards (section 5.3).
- Remove the `class-landing-identity` block and its CSS.
- Tests: update `SpecCard` tests and any h1-text assertion. The h1 now reads "Hunter", not "Hunter in Forever".
- `Content.astro` needs no change.

## 13. Acceptance (screenshots to attach)

Full-page captures, signed out, real data. Use the same viewports as `design/reviews/2026-11-02-class-landing/before-*`:

1. All nine classes at 390, 1024, 1440 and 2000 (36 captures).
2. Druid at 768 and 1280 (4-up 2 x 2, and the row layout).
3. Hunter at 768 (row cards) and 1280, with a viewport-height crop at 1280x720 and 360x640 to prove section 1.
4. Card states at 1440 on hunter: hover and focus-visible. Plus an active (mouse down) capture. Druid hover at 1440 for the 4-up width.
5. "Your spec" pill: warrior signed in as Fury at 1440 and 390. Show the same card before and after the pill with no shift (record the card box height before and after).
6. Same-scale side-by-side with the mock at 1440, plus the 1024 and 1920 captures (the owner's rule).

Checklist for the review:
- [ ] No spec chips, no figure, no nightly number anywhere.
- [ ] h1 is `--color-strong`, "{Class}".
- [ ] Each card line is the first Overview paragraph minus link sentences.
- [ ] The whole card is the link, with no underline and one CTA.
- [ ] Gap above the body is at least 40px. Gap between the cards and the head is 24px.
- [ ] List markers are visible at 390.
- [ ] Only the three headings appear. Each class matches its map in section 10.
- [ ] 1024 and up uses the full column width. Wide-screen content is centred at 2000.
- [ ] Confidence line and Sources are still present.
- [ ] Touch target at least 44px (the card is at least 88px).
- [ ] No console errors, no broken icon boxes, CLS under the budget.

## 14. wow-player spec review (SHIP WITH CHANGES) and how the build answers it

1. Mage's "Frost is the easiest to level" sat under "In a raid". The mage panel is titled **"Raids and leveling"**, a fourth allowed heading used by mage only.
2. "Plate" stays in the warrior and paladin identity lines (both identities are kept whole).
3. The Fury card line (192 characters) fits 354px in five lines at 15/22. The Overview text is approved prose and is not trimmed; the build fails a card line over 200 characters.
4. Warrior's eight change items get a bold ability lead (markup only, no words changed).
5. Priest's class colour is white: the card name is the 20px heading face in `--color-strong`-white against 15px `--color-text` body, which separates by size and face. No extra mark.
6. Feral and Feral Bear must never share an icon. Card icons come from `data/curated/specs.json`'s `icon` (the spec icons players know), with Feral drawn from `ability_druid_catformattack` (the override the raid roster already uses, owner ruling 2026-10-28); `web/src/lib/bis/spec-icon.test.ts` fails if two specs of a class resolve the same icon.

## 15. Coordinator rulings (2026-11-02)

- The landing shows no confidence sentence ("Reported by one outlet, not yet confirmed"): sourcing meta text does not belong in the reading flow, and the rewritten prose is client-verified. Sources stay in the footer. The page passes `hideConfidenceNote`; Content.astro is unchanged.
- `web/tests/e2e/guides-redesign.spec.ts` follows this design: cards are links carrying the Overview line, no figure, the two headings, no confidence line, distinct Feral and Feral Bear icons.

## 16. Mock

`design/mocks/2026-11-02-class-landing/mock.html` draws sections 4 to 7 for Hunter with the site's tokens (the same copy the page renders); `design/mocks/2026-11-02-class-landing/render.mjs` writes `mock-<width>.png` at 390, 1024, 1440 and 2000.
