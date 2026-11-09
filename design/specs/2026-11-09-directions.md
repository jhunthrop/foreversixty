# Design directions, 9 Nov 2026: round 5, soft corners and the owner's art set (on round 4's palette)

Mocks only; nothing here is built.

## Round 5 (supersedes round 4 where they differ)

**Owner, verbatim:** "make it have softer edges around the components to make it feel friendly and inviting like wow forever ... or try adding borders like they do on their square elements. I'd like it to fit the rounded nature of this nav bar element, though."

### Corners

Their nav bar's corner measures about 8 px in the captures. Ours is one scale (`style.py`: `R_LG`, `R_SM`, `R_IC`):

| Radius | Used on |
|---|---|
| 10 px | The nav, every panel, the raid plaque |
| 6 px | Buttons, the search control, the selected row |
| 3 px | Item icons |
| 4 px | The client tooltip only, the client's own |

### The frame

This is our rounded version of their bordered square elements.
- **`.framed`:** a 1 px cream line at 42% (`rgba(243,238,226,.42)`), inset 6 px, following the radius.
- **`.framed.ornate`:** adds small SVG corner flourishes (an inner arc, a scroll and a bead) to the page's one lead panel.

Where it goes, kept to two or three panels a page so it never becomes a stamped card:

| Page | Ornate (lead panel) | Plain frame | Plain rounded panels |
|---|---|---|---|
| Home | The raid plaque | Your upgrades | The feed and guild sims |
| BiS | The paperdoll | | The itinerary |
| Tier list | Your spec panel | | The table |

### The owner's art set

The owner explicitly chose this generated art on 2026-11-09, with the same written confirmation as the dusk road. It is used as background only and never captioned as a screenshot. Each file is loaded from `design/assets` (`art.py`, `assets.py`).

- **`hero-road-camp-wide.jpeg`, the home hero.**
  - It now has the band to itself: the floating nav at the top, the title at the foot.
  - It sits at 116% width so the dusk glow is centred and the lantern camp lands at the window edge, clear of the DPS figure.
  - It runs under a flat night veil, `rgba(7,29,46,.3)`.
  - Contrast: the title is 12:1. The art around the DPS figure measures 10.4:1 at the median and 7.4:1 at the 90th percentile.
  - On phones the crop shows the camp.
- **`raid-panel-cave.jpeg`, behind the Onyxia plaque.**
  - At 160% width, the cave mouth and its lava glow sit in the plaque's right edge, clear of the text.
  - It runs under a flat veil, `rgba(14,9,6,.36)`. Text contrast is 11.9:1.
  - The plaque moved out of the hero band to lead the page beneath it.
- **`blue-band.jpeg`, every section band.**
  - It runs navy at the top into soft teal cloud.
  - It is stretched with `cover` and never tiled, so its left and right edges never meet. That matters because its edge step is 3.1, against 0.6 between interior columns.
  - Its luminance stays at or below about 61/255, so the band's text clears AA (6.9:1 measured on BiS).
- **`brown-ground-tile.jpeg`, the page ground.**
  - The raw tile's seams measured 1.1 left-right and 1.9 top-bottom, against 0.7 between interior columns.
  - Made seamless by offset-and-feather (blending the tile with itself rolled by half): the seams are now 0.69 and 0.72.
  - **Depth (owner, 2026-11-08: "like the wow forever site. I think that adds a lot of depth").** Round 5 had halved the contrast after a stains note, and the ground read flat (sd 0.9 over 200 px windows). Blizzard's clean ground patches measure sd 2.1 to 4.9 over 200 px windows at 1x. A first pass (contrast 1.8 at 1400 px) matched their sd, but ux-designer and ai-slop both blocked it. Dark patches with edges read as stains, and the dark blotches repeated on a 700 px grid. That repeat was the offset-and-feather method itself: the tile's centre reappears at its edges.
  - **Now (`art.ground_tile`):** the tile is made periodic (Moisan's periodic component: no seam, nothing repeated inside the tile) at 1600 px. Its luminance is then rebuilt in three bands:
    - the large soft field, gained and soft-clamped both ways so no patch sits far from the ground around it;
    - the brush-edge band, damped;
    - the fine grain, lifted, since their ground has a fine surface.
  - **Measured at 1x:**
    - sd 2.8 over 200 px windows, inside Blizzard's 2.1 to 4.9;
    - darkest 200 px window 6.1 below the mean (8.8 before);
    - half-tile correlation -0.07 (0.88 with offset-and-feather);
    - seam steps half the interior step.
  - The side-by-side is `blizzard/texture-compare.png`: theirs, theirs, ours, ours, 300 px each at 1:1. If stains come back, the levers are the field cap and the stroke gain, not flattening. One tile size, `art.GROUND_TILE_PX`, serves web and mobile. The 2 x 2 check is `png/ground-seam-check.png`.
- **`social-banner-valley.jpeg`, the link preview.**
  - It is byte-identical to `hero-road-dusk.jpeg`, so the round-4 preview is unchanged: the road-and-glow crop under the veil, with our seal, "Forever Sixty" and "Fan-made player tool for WoW: Forever".
- **`discord-banner-tavern.jpeg`:** the Discord banner, and a candidate for guild pages (`png/discord-banner.png`).
- **`app-icon-background.jpeg`:** the app-icon field. The sample `png/app-icon.png` has the flat seal at 62% inside the hammered-gold frame.
- **`blue-band-tile-REJECT.jpeg`:** not used.

### Logo

- The flat seal (`lx-seal-flat-64.png`, shown at 40 px) in the nav.
- The painted seal (`lx-seal-256.png`) at 88 px in the footer and in the link preview, and in the paperdoll medallion when there is no portrait.
- **Owner decision, 2026-11-08: the logo stays as is** ("keep the logo").

### Premium

**Owner decision, 2026-11-08: "just copy the subscribe button - keep the logo".** Premium is Blizzard's teal Subscribe exactly. This supersedes round 4's gold outline and the three in-palette options tried the same day (cream outline, seal plus cream, gold inset), which are dropped. Measured from their bar: flat `#408076` (`--teal-edge`), 47 px tall in the 64 px bar, 8 px inset from the bar's top, bottom and right, 6 px radius, white Open Sans semibold 16 px. Their label face looks closer to our Archivo nav face than to Open Sans; that is one token to switch if the owner prefers it. The official-page risk ux-designer raised in round 4 is now the owner's accepted choice; the seal, the name and the not-affiliated footer still carry the distinction.

The real account name (Frickin Laserbeams 31) made the bar 25 px too wide, so the menu gap is 22 px rather than 26, and the bar folds into Menu below 1320 px (it needs 1318).

### Paperdoll (owner, 2026-11-09: "it should look like the ingame paperdoll... can we overlay the items on this?")

The frame is `design/assets/art/paperdoll-frame-source.jpeg`, generated by the owner; using it is the owner's explicit choice. It is cleaned by `paperdoll_asset.py` into `design/mocks/directions/paperdoll/frame@{1x,2x}.{webp,png}`:
- **Alpha:** the baked checkerboard (near-grey, light, connected to the edge) is cut to real alpha with a one-pixel feather, and the canvas is cropped to the frame's outline.
- **Centre:** every teal pixel inside the knot border, which includes the fake writing and map scratches, is replaced by the owner's blue band, colour-matched to the panel. The gold knot corners stay.

**Measurements.** `paperdoll/slots.json` holds every slot's inner rectangle, the medallion, the title bar and the centre, as percentages of the frame. The slots were measured on gridded crops and refined to each slot's teal edges; a glyph that crosses an edge takes its row from its neighbours. The frame's order is the client's:
- **Left:** head, neck, shoulder, back, chest, shirt, tabard, wrist.
- **Right:** hands, waist, legs, feet, finger, finger, trinket, trinket.
- **Bottom:** main hand, off hand, ranged.

**The overlay** (`paperdoll.py`):
- **Slots:** each item's real icon sits inside its slot, inset 6%, with a 2 px border in the client's quality colour.
- **Empty slots:** they keep the frame's own glyph. At 20 to 29 these are the shirt, tabard and both trinkets.
- **Two-hander:** it shows greyed in the off hand, as the client draws it.
- **Medallion (owner, 2026-11-08):** the character's Battle.net portrait, the head render (`characters.avatar_url`, the `-avatar.jpg` media; `-inset.jpg` returns 403). With no portrait, signed out included, the LX seal. Never the class crest.
- **Title bar:** "Beast Mastery Hunter 20 to 29 · Alliance" on one line, with "91.0 DPS" at the right. The page's h1 is the bar, so the blue band above holds only the nav and the controls.
- **Centre, public:** the selected slot's client tooltip and Open in the sim.
- **Centre, signed in:** the character's Battle.net render (`characters.render_url`, migration 0026, captured on import) stands in the panel as the client's pane shows the model. The name in class colour, "Level 31 Night Elf Hunter" and "As last logged out" sit above it. Only the tooltip stands over the model. Where it drops, what it replaces and Open in the sim go in a bar under the frame, so the model's legs show (wow-player). Hover or tap on a slot swaps the tooltip.
  - The mock uses the owner's own character, Frickin Laserbeams (render `forever-us/character/204/6928844`, level 31 from production). Blizzard's media are referenced by URL and never committed. The signed-in captures, which show them, are git-ignored and live only in the gallery.
  - The mock keeps the 20 to 29 list selected although the character is 31: the repo's BiS facts have no 30 to 39 list. A build opens on the character's own band.
- **Beside the frame:** What to stack and Notes. The itinerary by dungeon stays below.
- **Phone:** under 36 px slots (a frame narrower than 424 px, so every phone) the frame gives way to its one-line header strip (medallion, title, DPS) and a two-column slot list on its bronze and teal. The render does not fit the strip without bringing back a second line, so the phone shows the portrait in the medallion and no model.

**Mobile Gear tab:** the same header strip and list for Thandrel's Fury list (1004.2 DPS, the bis file's set), then Where to go tonight. Its medallion shows Thandrel's portrait, the same one as the profile button.

### Tab-bar icons

The four icons are the owner's, traced to single-colour SVGs: `design/assets/icons/nav/{home,raid,gear,activity}.svg`, `fill="currentColor"`.
- **Home:** a rune stone.
- **Raid:** crossed swords on a banner.
- **Gear:** a T-visor helm.
- **Activity:** a war horn.

They are used at 24 pt on iOS and 24 dp on Android: gold `#ffd36b` when selected, muted white otherwise.

At 24 px the swords' grips and the horn's strap broke up, so the mock adds a same-colour stroke of 22 units (about two-thirds of a pixel) at embed time (`nav_icons_check.THICKEN`). The traced files stay as they are; the build should bake that stroke in.

**Owner decision, 2026-11-09:** the spiral home stone ships as is ("let's ship it - I like it - it's a quick change if it draws issue"). Each icon stays one SVG file, so swapping one later is a one-file change.

### Mobile (iOS and Android in the round-5 look)

The flows, tap counts and every element come from `design/specs/2026-11-08-app-flows.md`: its tap table and its purpose audit. The content is imported from that spec's mock (`design/mocks/ws_app_flows.py`), so the two stay one source. Only the visual language changes:
- the Forever palette;
- Open Sans and Marcellus;
- rounded panels, with the ornate frame on the raid card (the lead) and the plain frame on Next upgrades and the Gear summary;
- the owner's art: the camp-road hero on Home, the cave behind the raid card, the painted blue band behind the other tabs' bars, and the brown ground;
- the flat seal in notifications.

Renders: `design/mocks/directions/mobile.py` and `shoot_mobile.cjs`, producing `png/m-<ios|android>-<screen>.png`. iOS is the iPhone 15 (393 x 852) and Android is the Pixel 7 (412 x 915).

**Screens:**
- the four tabs: Home (raid day), Raid, Gear, Activity;
- the profile drawer;
- the raid sign-up: Late with the 15/30/60 chips inline under the answers, on the card, as spec F1 says;
- the item sheet (Deathbringer);
- the lock screen Live Activity on iOS, the ongoing notification on Android.

**Platform conventions:**
- **iOS:** large titles, the tab bar with badges, sheets with a grabber, an 85% modal drawer, 44 pt answer buttons.
- **Android (Material 3):** the top app bar carries the title, so the title is never doubled. The navigation bar uses the pill indicator. The answers are a 48 dp segmented button. The drawer is `ModalNavigationDrawer`, and the item sheet is a modal bottom sheet with 28 dp corners.

**One state across the set.** Thandrel is In everywhere: Home, Raid, the Live Activity, and Activity's "You're In for Onyxia". New for you is Activity's three newest rows, with the same times.

**Greenhammer** is left out because it unlocks with the first raids, so the Gear off hand shows Persuader worn.

**The lock screen** sits on a neutral system wallpaper. The wallpaper is the user's, never ours.

**Kept on purpose, against reviewer notes, because each is a binding app-flows rule:**
- The Battle.net portrait on the profile button and the menu rows. This is the owner's written exception to the crest rule (spec section 3).
- "as Thandrel, Damage" on every raid answer (rule E7).
- The raid-day moment on the app's own mock clock (Wed 6 Jan), which differs from the web home's day 32 days before the first raids.
- No "over <item>" on the axe's gain. The owner's rule keeps "over <item>" for rings and trinkets.

### Logs, reports and rankings (batch 1)

Every logs surface in the round-5 look: `/logs`, `/reports`, the report's tabs (Summary, Damage done, Healing, Damage taken, Deaths, Mechanics, Timeline, Threat, Resources, Queries, one player's view, Coach) and `/rankings` (encounters, the character board, the guild board). Each is drawn at 1440 and 390 as `png/b1-*.png`.

```
python3 design/mocks/directions/report.py && python3 design/mocks/directions/logs.py && python3 design/mocks/directions/rankings.py
node design/mocks/directions/shoot.cjs <web/node_modules> b1-logs b1-reports b1-report-summary ...   # every html/b1-*.html
python3 design/mocks/directions/logs_gallery.py                                                     # the gallery's Logs section
```

**Shared primitives.** Every page is drawn from `prims.py`; nothing is page-local beyond layout:
- `panel`, `table` (sticky header, tabular figures, one fixed name and meter column so the bars never move between tabs), `player_cell`, `meter`, `facts`, `row_list`, `chips` (a chip can carry its mark's colour, so the filter is its own legend), `code_block`, `parse_colour` and `parse_legend`, plus batch 2's `tag`, `ability_icon`, `item_row` and `lanes` (the Timeline's lanes).
- **One chart axis** (`Axis`) for the line, area and column charts:
  - round minute ticks from `TICK_STEPS` (at most 8, 4 on a phone);
  - a y top from `NICE`, so the four gridlines are round;
  - the unit above the axis;
  - the faint grid and crisp 1 px rules;
  - deaths as red ticks on the axis.
- **Series colour:** class colour, and a repeat of a class steps lighter or darker (`shades`) and is dashed.
- **Phone drawing:** each chart is drawn a second time at 380 wide (`responsive_chart`).
- **Phone band:** three short lines (title, meta, one control strip that scrolls with an edge fade and opens on the selected tab).

**Page rules.**
- **The band's figure follows the tab:** group DPS, HPS or DTPS; deaths on Deaths and Mechanics; the player's DPS on their view; none on Timeline, Threat, Resources, Queries and Coach.
- **This pull** keeps only what the band does not say: when it started, the scope's count and its deaths.
- **Long tables** show their top 15 and "All 265".

**Data.**
- **Pulls:** the real 17 Sept sample (trash only).
- **Mechanics:** Onyxia's table (`logs/engine/mechanics/tables/1084.json`) and the guild report fixture's wipe. Deep Breath, which killed three, is missing from the table, so it gets its own row until the table has it.
- **Coach:** the Fury fixture, with proposed copy (`report.COACH_COPY`).
- **The player's gear:** a stand-in, the Beast Mastery 20 to 29 list. The sample has none, because the game writes gear at boss pulls.
- **Rankings:** the rankings spec's example board, with the e2e stub's encounter and 212.

**Icons.** No question marks. A spell without a build icon takes another same-name spell's icon, then its school's. Onyxia's physical abilities take her own icon. Auto Shot is `ability_whirlwind`, as the client draws it.

**Proposed product changes the mocks carry.**
- The Queries preset sums `amount - greatest(overkill, 0)`, so it matches Damage done.
- The public feed leaves out the viewer's own reports.

### Everything else

Everything else is round 4 (below):
- the palette, fonts and floating nav;
- no big headers, units-only numbers;
- the paperdoll-led BiS with the itinerary;
- the client tooltip and the medallion.

---

## Round 4

**How to rebuild.** Everything lives in `design/mocks/directions/`:

```
python3 design/mocks/directions/pages.py                           # -> html/g-home, g-bis, g-tiers
node design/mocks/directions/shoot.cjs <web/node_modules>          # -> png/<page>-1440|390.png, home also -2000
node design/mocks/directions/shoot_now.cjs <web/node_modules>      # -> now/ (the live site)
node design/mocks/directions/shoot_blizzard.cjs <web/node_modules> # -> blizzard/ captures + study.json
node design/mocks/directions/palette_blizzard.cjs <web/node_modules> # -> blizzard/palette.json (every computed colour)
python3 design/mocks/directions/study_sheets.py && node design/mocks/directions/shoot_sheet.cjs <nm> blizzard/fonts.html blizzard/fonts.png
node design/mocks/directions/shoot_sheet.cjs <nm> blizzard/ground.html blizzard/ground.png && python3 design/mocks/directions/study_sheets.py compare
python3 design/mocks/directions/gallery.py                         # -> ~/Desktop/forever-design-directions/index.html
```

**Captures.** `blizzard/*.png` and the two study sheets contain Blizzard's pixels, so they are git-ignored. They travel only to the local gallery.

**History.** Rounds 1 to 3 are retired and kept in git:
- Rounds 1 and 2: `7156c5f23`.
- Round 3: `0edfd08a1`. This was the owner-approved artifacts style, which the owner judged "less AI" but still muddy and dark-on-dark.

## 1. Owner, round 4 (verbatim)

- "I think we can copy the wow forever colors exactly with no issue? ... I want the colors to still feel crisp and clean like the wow forever site."
- "I also don't like the journey map - we can remove that. It can be used as silhouette type of item though."
- "see how their backgrounds on the site look like they have texture? Can we do that too?"
- "use the same fonts and everything I even like their navigation bar - the entire site is clean"

**Still binding:**
- No big headers.
- A number takes its unit and nothing else (ai-slop-reviewer tell 22).
- No Blizzard logo, art or ornament.
- No generated images.
- No em dashes.
- We must never be mistakable for an official Blizzard page. Our seal and name lead the bar, "World of Warcraft" is never our branding, and the footer keeps the not-affiliated line.

## 2. Colours, exactly theirs

**Sources.**
- **c:** computed CSS from their live page (`blizzard/palette.json`, every element including shadow roots).
- **s:** the mean of a flat region of a capture, or the coordinator's median sample where it agreed.

The tokens live in `style.py` (`TOKENS`), each with its source.

| Token | Hex | Source |
|---|---|---|
| nav | `#0a2933` | c, the navigation colour |
| nav-2 (the floating bar) | `#0b2f3b` | c, `rgb(11,47,59)` |
| band | `#054562` | s, their section band runs `#053d58` to `#044c6c`; flat midpoint |
| navy | `#001845` | c |
| night | `#071d2e` | s, the hero sky |
| earth | `#67563a` | s, the brown ground (`#645337` to `#6a593c`) |
| card, every list surface | `#170e09` | c, their footer and body brown-black. Their `#1e1e1e` article card read grey on our pages (ai-slop-reviewer). |
| foot | `#170e09` | c, the footer and page body, also the card surface |
| cream | `#f3eee2` | c, display type |
| cream-2 | `#fcedcf` | c, footer heading |
| text | white at 0.78 | c, their body is white at 0.7; 0.78 for our denser 16 px |
| muted | white at 0.5 | c |
| gold, links | `#ffec71` | c, their "Learn More" |
| gold-hi, hover | `#f7b90a` | c |
| teal, primary button | `#0f6064` | c, "Buy Now" |
| teal-edge, bar button | `#408076` | c, "Subscribe" |
| blue, active tab | `#0093b2` | s, the tab underline |
| blue-btn, secondary button | `#114f72` | c, "Upgrade Now" |
| rule | `#5a5d70` | c, their border |
| rule-warm | `#82705b` | c |
| gain | `#32ae88` | c |

### Where the coordinator's sample differs, and why

- **Nav `#0e314b`.** That is their bar at 0.8 alpha over art. The CSS value is `rgba(11,47,59,.8)`, so we use the solid `#0b2f3b`.
- **Band `#06324b` to `#053d59`.** My region of the same band measured `#053d58` to `#044c6c`. Both are near-flat, and the flat midpoint is `#054562`.
- **Heading cream `#dccdbe`.** Their headings are gradient-filled text, `#f3eee2` to `#e2c2ab` (`study.json`), and `#dccdbe` is the median of that gradient. Gradient text is slop tell 2, so ours are flat `#f3eee2`.

## 3. Texture

**Measured on their page.**
- The brown ground has a luminance standard deviation of 2.1 to 5.3 around a mean of 87 to 94: large soft mottling with a faint fine stone surface.
- The blue band measures under 1. It is flat, a slow vertical shift with no texture.

**Ours.**
- Round 4a matched their spread with a large mottle plus a fine layer (sd 3.1 to 4.4). Two reviewers read that as blotchy stains, and the slop reviewer failed home on it.
- Round 4b keeps only one large, soft, colour-matched variation (feTurbulence baseFrequency 0.003, no fine layer). It is near flat, at sd 2.1 around a mean of 89.5, which is still inside their range.
- The blue band is solid with no vignette, as theirs measures flat.
- Every list sits on a flat surface.

**The limit.** Procedural noise matches their hue and spread but not the painted speckle of their stone (`blizzard/texture-compare.png`, theirs left, ours right).

If the owner wants that last step, the route that keeps the no-generated-imagery rule is a CC0 photographed surface, tinted to `#67563a`: Poly Haven or ambientCG, public domain, real photographs. A generated tile would break the owner's own rule, so it is not recommended.

## 4. Type

| Use | Font | Detail |
|---|---|---|
| Body | Open Sans | Exactly theirs. 16 px for density (theirs is 18 to 20), white at 0.78. |
| Navigation | Archivo 600, 17 px, white | Exactly theirs. |
| Display (titles only) | Marcellus | Stand-in for their BlizQuadrata Web, Blizzard's own cut of ITC Friz Quadrata, which we cannot use. Mixed case, tracking -0.015em, cream with a soft shadow. |

**Display font test.** `blizzard/fonts.png` sets their heading crops ("Undiscovered Vistas Await" at 60 px, "Latest Updates" at 48 px) against eight Google Fonts:
- Marcellus, Philosopher, Rosarivo, Cinzel, Quattrocento, Forum, Bellefair and Cormorant Garamond.

The verdict:
- **Marcellus** is closest in build and flare, once tracked tighter. Their stroke is heavier.
- **Quattrocento** is second.
- **Philosopher** and **Rosarivo** are too calligraphic.
- **Cinzel** is caps only.

**Owner decision, 2026-11-08:** Marcellus stays as the display face; no Friz Quadrata licence is bought.

**Licensed path.**
- ITC Friz Quadrata (Monotype). MyFonts lists single styles from US$43.99 and the family of four at US$174.99 for desktop.
- The web licence is priced by monthly pageviews at checkout. The page rate-limited my fetch, so the owner should read the exact tier there.
- In `style.py` the swap is one line, `DISPLAY_FONT`, plus its `@font-face` in `DISPLAY_FACE`.

## 5. Navigation bar, in their form

- Their floating rounded bar, inset over the section band and aligned to the content width at every size: 64 px tall, `#0b2f3b`, 6 px radius, 1 px light edge.
- **Left:** our LX seal at 40 px and "Forever Sixty" at 25 px, divided from the menu by a rule, as their logo block is.
- **Menu:** Archivo white labels, chevrons on the menus (Guides, Logs, Guild), and a cyan underline under the current page.
- **Right:** search, then the account on one line (medallion, name in class colour, level, chevron), then Premium as their teal Subscribe (owner, 2026-11-08; see Premium in round 5, which supersedes the gold outline).
- **Phone:** the bar insets 8 px and shows the seal, name, account medallion and Menu.
- The seal comes from `design/assets/logo`, committed with the owner's art and nav icons. A missing seal stops the build.

## 6. Pages

### Home (Thandrel, Fury Warrior 60, Olympus XXVII)

**Decision, 2026-11-09 (owner, in writing).** The owner explicitly chose to use their generated dusk-road image as home background art.
- The owner's words: "this isn't a real in-game screen shot - can you use it?"
- The answer: yes, as background, never captioned as a screenshot.
- This is the written confirmation the ai-slop-reviewer asked for under its tell 15.

**The owner's art.**
- **The image:** `design/assets/art/hero-road-dusk.jpeg`, a landscape generated in the Classic style.
- **Where it goes:** background only, on home only, behind the title band and the raid card. It is never captioned or implied to be a screenshot.
- **The veil:** one flat night-blue veil, `rgba(7,29,46,.5)`.
- **Contrast:** 9.8 to 12:1 for the title, sampled conservatively against the brightest art behind the text. AA needs 4.5.
- **Desktop crop:** positioned at 70% 20% so the dusk glow sits right of the raid card. The card is capped at 960 px so the glow stays clear at 2000.
- **Phone crop:** a 130 px text-free strip between the title and the card, at x 58 to 90% and y 38 to 58%. It shows the glow and the head of the road, clear of the smeared cliff.
- **Artefacts at full size:** the smeared rock face left of the road, melted foliage on the right-edge pine, and soft trees in the mist. The veil and the crops keep them out of sight.
- **Link preview:** `png/og-home.png`, 1200 x 630. It is the road-and-glow crop under the veil, with our seal, "Forever Sixty" and "Fan-made player tool for WoW: Forever", so a shared link can't read as an official promo.

- **The band** holds the bar, then the compact title: medallion, name in class colour, "Level 60 Human Fury Warrior · <Olympus XXVII>", and "948.6 DPS".
- **No silhouette.** The route line was tried as a faint drawing in the band and cut. Two reviewers read it as decoration.
- **The raid card leads,** on charcoal:
  - Onyxia's Lair, with the guild, day and time.
  - Days and hours at 44 px, the largest thing on the page, then "signed up" and "late".
  - Role counts.
  - I'm in (teal), Late (blue, pressed) and Can't come (blue).
  - Before the pull, in the right half.
- **Under the band:** the feed on the left and Your upgrades in the rail (`+14.5 DPS`).
- **Guild sims:** full width below, in two columns (`1012.3 DPS`), then "Tier A · #2 of 20 · Tier list".
- **Latest on Forever.** The guild raid post is left out, because the raid card already shows it. Guild sims read "#2 of 14".
- **The footer** is their brown-black, with our seal large at 88 px and the not-affiliated line.

### Best in slot (Beast Mastery 20 to 29)

- **The band:** "Beast Mastery Hunter", "Best in slot", "91.0 DPS", and the spec, level and faction tabs.
- **The paperdoll stands alone** on charcoal: left column, right column, weapons below.
  - Where the character would stand, the selected item's client tooltip, ending "+0.3 DPS" ("over <item>" appears only for a ring or trinket), and the sim link.
  - Empty sockets are dimmed. The harpoon reads Two-Hand, and there is no off-hand socket.
- **The itinerary** sits below on charcoal, with ruled sections in route order:
  - Wailing Caverns 15 to 24, The Deadmines 17 to 26, Crafted, Silverwing Sentinels reputation, Quest, World / BoE.
  - Each row gives the item, what it adds to its heading, and `+0.3 DPS`.
- **On a phone:** the doll is a grid of 44 px sockets (six per row, so the trinkets sit together), with the selected item's tooltip right under it.
- **The road map is removed** (owner, round 4).

### Tier list at 60

- **The band:** "Tier list at 60", "Best in slot, raid buffs · level 63 · 3 min", and the Damage/Tanks/Healers and faction tabs.
- **The list,** on charcoal: grouped S to D. Each row has rank, medallion, spec in class colour, a cyan relative bar that starts at 90% of the lowest spec, DPS and "vs #1". Your spec is tinted.
- **Your spec panel:** "Tier A · #2 · 1004.2 DPS", with Best in slot (teal) and Guide (blue). It sits first on a phone.
- **Data:** the shipped ranking (`design/mocks/ws_tiers.py`), with 20 damage specs.

## 7. Reviews

Verdicts are in `design/mocks/directions/reviews.json` and in the gallery.

| Page | ux-designer | wow-player | ai-slop-reviewer |
|---|---|---|---|
| Home | SHIP | SHIP | PASS WITH FIXES |
| BiS | SHIP WITH FIXES | SHIP WITH FIXES | PASS WITH FIXES |
| Tier list | SHIP | SHIP | PASS WITH FIXES |

The fixes named in round 4b were applied after it.

**"Does it feel as crisp and clean as Blizzard's Forever page?"**
- ux-designer: yes.
- wow-player: yes.
- ai-slop-reviewer: yes as a flat, clean tool page. Their page alternates full-bleed bands; ours keeps cards for lists, which a dense tool needs.

**Round 4a, for the record:**
- ai-slop-reviewer failed home on the blotchy texture, the four stamped boxes and the band reading as a gradient.
- All of that is fixed.

**Owner decision: the LX seal.** It is the owner's logo, so it was not changed. ux-designer flags it as a conscious risk: a teal ringed medallion with a gold rim, close to Blizzard's Forever badge.

## 8. Data gaps (named, not built)

- **Item tooltips** lack durability, sell price, weapon speed and set lines. The Forever wording for rating stats is unknown.
- **Crafted items** have no recipe source or skill level.
- **Zone level ranges** are missing from `zones.json`. They are only needed again if the route returns as content.

## 9. Screenshots the owner needs to take in game

1. Westfall, Redridge or Duskwood at dusk from a road, UI hidden. This is the framed plate on the signed-out home and the guides.
2. The guild at Onyxia's Lair on 9 Dec, UI hidden.
3. The character sheet with the addon tab open, full gear, 1x UI scale.
4. A tooltip for an item with crit or hit rating, and one weapon tooltip.
5. The Wailing Caverns and Deadmines entrances.
