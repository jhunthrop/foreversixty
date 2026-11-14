# Addon brand skin: the hub window carries the site

Status: spec for a later build lane. Nothing in addon/ changes in this lane. Spec author: ux-designer. Previous hub spec: design/specs/2026-11-13-addon-hub.md (window 740 px wide, header 56, strip 28, footer 28). This spec keeps every tab's content and layout. It changes the chrome (ground, header, strip, plates, rows, footer), the iconography and the crests.

## 0. The one question and the finding

The question a player has when the window opens: "Is this the Forever Sixty thing I use on the site?" The answer must be visible in the first frame, with no tab clicked: blue-teal painted header with the ringed crest, a brown painted ground showing between umber plates, a nav-bar strip with the seal and site icons.

Finding (reading design/reviews/2026-11-14-addon-brand/site/*-1440.png against design/listing/captures/window-me.png, window-gear.png, window-builds.png): the shipped window is one flat dark slate (base.bg #100a06) with gold. Theme.lua reads base.bg, raised, line, gold, text, and forever.cream/teal/blue only on buttons. It never reads forever.ground, band, navBar, navy, accent or mutedGround. There is no painted art, no site icon, a 64 px crest, and a plain gold underline on tabs.

Reference surfaces and what is borrowed:
- The site's own page chrome (bis-1440.png, home-1440.png): ground, band, nav bar, umber plates, Marcellus titles, teal and blue buttons, accent underline. Borrowed whole, at addon scale.
- The in-game character panel: the window frames a character, so the crest sits in the header the way the portrait sits in the character frame.
- A third-party journal addon the owner likes (ideas only): layered tinted surfaces (frame < content < row < selected row at 90 to 99 percent alpha), a painted band per section, round masked portraits, strong title type, a quest-log row highlight, one pulse at the next action. We take the ideas in our site's colours. We copy no art or code.
- Not borrowed: native Blizzard frames. The window stays a custom frame. Only tooltips and item/spell icons stay native (section 6).

## 1. What makes the site the site

1. The earth: painted ground tile, token forever.ground #67563a (web/public/art/ground.webp, 1600 x 1600).
2. The band: painted navy-into-teal cloud at the top of every page, forever.band #054562, veilBand rgba(5,58,84,0.38) (web/public/art/band-1200.webp, 1200 x 442).
3. The nav bar plate: forever.navBar #0b2f3b with the seal and a Marcellus wordmark; the active item is underlined (forever.accent #0093b2), labels forever.tabOff rgba(255,255,255,0.82), active white.
4. Dark umber plates on the ground: base.raised #170e09, edge forever.panelEdge rgba(130,112,91,0.55), 1 px.
5. Ringed HD crests (web/public/icons/hd/crests/*.webp, web/src/styles/crest.css).
6. Type: Marcellus titles in sentence case, Open Sans body (both already in addon/ForeverSixty/media/fonts).
7. Buttons: teal primary forever.teal #0f6064 / forever.tealEdge #408076; blue secondary forever.blue #114f72 / forever.blueEdge #1d6a96.
8. One selected state: base.selected rgba(0,147,178,0.16).
9. The corner filigree on the lead plaque (design/assets/ornaments/corner-filigree.svg), the paperdoll's painted teal frame (web/public/art/paperdoll-frame@2x.webp), the cave art behind raid surfaces (web/public/art/cave-800.webp).
10. The site's custom filled icon set (design/assets/icons/ui/*.svg) and line family (design/assets/icons/ui/line/*.svg).

## 2. Surface map and layer order

All colours are tokens from design/tokens/tokens.json, read through addon/ForeverSixty/Tokens.lua (generated; never hand-edit). Window width 740 px. Height unchanged from the shipped hub.

### 2.1 Layer order, back to front

| Layer | Content | Colour / alpha |
|---|---|---|
| L0 | Drop shadow, Theme.SIZES.shadow 5 px | addon.shadow, as shipped |
| L1 | Window ground: painted tile, whole window, BACKGROUND | texture alpha 1.0, no tint (vertex white) |
| L2 | Window edge: 1 px outer line, then 1 px inner line | outer forever.edge #000000 alpha 1.0, inner forever.panelEdge |
| L3 | Header band art (section 5) with veil | art alpha 1.0, veilBand over it |
| L4 | Tab strip plate | forever.navBar alpha 0.97 |
| L5 | Content plates | addon.plate rgba(23,14,9,0.84) (new token), 1 px panelEdge border |
| L6 | Row fills: zebra, hover, selected | see 2.5 |
| L7 | Text, icons, crests, item icons | tokens by role |
| L8 | Footer plate | forever.foot #170e09 alpha 0.85, so the earth shows under it |

Contrast rule: no text sits on bare ground. Text sits on L3 (white on navy-teal, veil guarantees 7:1 minimum), L4, L5 or L8. The ground shows in the 10 px side margins, the 6 px gaps between plates and under the footer, and warms each plate through its 16 percent transparency. A control that must sit on bare ground uses forever.mutedGround #e2dcd3 text. There is no such control in this spec.

### 2.2 Window ground

- Chosen: the painted ground tile exported to a texture, tiled, plates on top. Flat forever.ground is the fallback when tiling is not available (capability check, same pattern as Theme's mask check near line 451), and the plate rule is the same either way.
- Texture: media/ground.tga, 256 x 256, 32-bit uncompressed, 262,144 bytes. Generator downsamples ground.webp with sharp (Lanczos), then makes it seamless by a 32 px cross-fade on the wrap edges. The generator fails if the wrap seam exceeds a mean colour delta of 6 on any 8-bit channel.
- Tiled with SetHorizTile(true) and SetVertTile(true), tex coord 0..(740/256) by 0..(height/256). Draw layer BACKGROUND, sub-level 0.
- Window margin: 10 px left and right, 6 px top and bottom of the content region. Between plates: 6 px. Plate padding 8 px on every side. At 740 px this gives a 405 px left column and a 307 px right column (inner 389 and 291 px), so no string the shipped window fits is truncated (the mock checks the Rotation and Me tables).
- Window height grows by 18 to 43 px over the shipped hub (plate borders and title rules): Me 527, Builds 519, Gear 645, Rotation 613, Guild 514 px in the mock. The tallest, Gear at 645, fits a 768 px screen at UI scale 1.0 with room for the client's own bars.
- Ground tile cut: a 640 px window of ground.webp downsampled to 256 px (the whole 1600 px tile at 256 px reads as noise).

### 2.3 Header band (56 px)

- Region: x 0 to 740, y 0 to 56, directly under the 1 px window edge.
- Art: media/band-<mood>.tga, 512 x 64, stretched to 740 x 56 (tex coord 0..1 by 0..0.875). Soft painted cloud, so the 1.45x horizontal stretch holds. Moods in section 5.
- Veil: forever.veilBand over the art, plus on the left 260 px a linear fade of addon.inset from 0.0 at x 260 to 0.35 at x 0, so the crest and name sit on a deeper ground. The fade uses Theme.verticalGradient's horizontal counterpart if SetGradient exists, otherwise 4 stacked 65 px strips at alpha 0.09, 0.18, 0.27, 0.35.
- Crest: 40 px ring crest at x 12, y 8 (section 4), replacing the shipped 36 px.
- Name: Marcellus 18 px, base.strong, at x 62, y 9. One line, truncated at 380 px.
- Figures line: Open Sans 11 px, forever.inkSoft rgba(255,255,255,0.72), at x 62, y 31. Shipped text unchanged ("Level 60 Hunter . guild . build").
- Close: 22 x 22 at x 710, y 17, line icon close.tga 12 px, tinted tabOff. Hover: tint white and a base.hover plate at alpha 0.5. Focus ring: 1 px accent. Active: pressed plate base.pressed.
- Bottom edge of the band: a 1 px forever.accent line at alpha 0.5 (the join to the nav bar), across 740.

### 2.4 Tab strip as the site nav bar (28 px)

- Plate: x 0 to 740, y 56 to 84, forever.navBar, with a 1 px forever.panelEdge line at its bottom.
- Seal: media/seal.tga 32 x 32 drawn at 20 x 20, x 8, y 60, tinted none (the art is forever.seal #2c8a9e on transparent). It is a button: click does nothing in 2.x; hover shows the tooltip "Forever Sixty" (tooltip frame native). Decorative for keyboard: not focusable.
- Tabs start at x 36. Each tab has a fixed width of 117 px (704 px over six tabs), icon 14 px, 6 px gap, label Open Sans SemiBold 12 px, centred. The width is fixed so the strip never moves when the active tab turns bold or a dot appears. Order and labels are shipped: Me, Builds, Gear, Rotation, Guild, Officer (strip shows only the tabs the hub shows, as shipped).
- Default: icon and label forever.tabOff.
- Hover: label and icon white; plate rgba(255,255,255,0.06) (addon.tabHover, new).
- Active: label and icon base.strong white, Open Sans Bold 12 px, plate forever.band alpha 0.55, and a 2 px forever.accent underline across the tab's full width at the plate's bottom edge. This replaces the shipped gold underline and forever.tabOn gold label. Gold stays only for the waiting dot.
- Focus-visible: 1 px forever.accent rectangle inset 1 px, no fill change.
- Pressed: plate forever.band alpha 0.8.
- Waiting dot: 6 px, base.gold, after the label, as shipped.
- Disabled tab (a tab whose data is absent but that still shows): icon and label tabOff at alpha 0.4, no hover, tooltip as shipped.
- Keyboard and touch parity: Left and Right move between tabs, Enter or Space selects, as shipped. The WoW client has no touch; the 28 px strip height is kept.

### 2.5 Panels, titles, rows

Panel (every shipped card or column becomes one):
- Fill addon.plate (L5), border 1 px forever.panelEdge, no corner radius (the client frame is square; the site's 10 px plate radius is not reproduced).
- Inner padding 8 px top and bottom, 10 px left and right (section 2.2).
- Title: Marcellus 15 px, base.strong, sentence case ("Forever Score", "Last raid", "Right now", "Your sim"), baseline-aligned with the shipped right-side note (Open Sans 11 px, base.muted). Under the title row: a 1 px forever.veilRule hairline, 4 px below the title, then 4 px to the content.
- Lead plate (the plate holding the tab's first answer: Me "Right now", Builds the first class block, Gear the slot list, Rotation the priority, Guild "Tonight", Officer the roster): the site's filigree corners (web/src/components/ui/FrameCorners.svelte), the two top ones only: media/filigree.tga 64 x 64 drawn at 28 x 28 top left and top right (mirrored by tex coords), inset 3 px, tinted forever.cream at alpha 0.8 as the site's panel.css does. The lead plate's top padding is 15 px so the title clears the curls; the other sides stay 8 px. The bottom corners are left out because the addon's plates are too tight for them to clear text (the site's plates have room for four).
- Guild and Officer lead plate: cave art behind it, media/cave.tga 256 x 128 (the cave mouth and its glow, cut from cave-1600.webp at x 480 to 1600, y 60 to 620), anchored right, veilCave over it, then a horizontal fade of base.raised from alpha 0.90 at the left (to 40 percent of the width) to 0.50 at the right, so the glow reads behind the right edge and the text side stays dark.

Rows (every list row: logs table, gains, roster, build rows, slot rows):
- Default: no fill. Row height as shipped.
- Zebra: odd rows addon.zebra rgba(255,255,255,0.03) (new). Rows under 24 px tall skip zebra.
- Hover: base.hover #241b15 at alpha 0.9 and cursor stays; no border.
- Selected (the one selected state): base.selected rgba(0,147,178,0.16) fill, a 2 px forever.accent bar on the row's left edge, full row height. The followed build on Builds is drawn at addon.selectedActive so it reads at 1x; its "Following" word is forever.accent. Selected text stays base.text. The quest-log highlight: when the selected row's tooltip or detail is open, the fill rises to rgba(0,147,178,0.24) (addon.selectedActive, new).
- Focus-visible: 1 px forever.accent rectangle inset 1 px; with selected it sits over the fill.
- Pressed: base.pressed #2e241d.
- Disabled: text alpha 0.4 (shipped ALPHA.dim), no hover.
- Row divider: none inside zebra tables; lists without zebra use forever.veilRule 1 px.

Buttons (Widgets.lua button constructors; no new control):
- Primary (one per region at most): forever.teal fill, 1 px forever.tealEdge, label base.strong Open Sans Bold 12 px, height as shipped. Hover: tealEdge lightened 8 percent via fill #12727a (addon.tealHover, new). Pressed: fill forever.tealEdge at alpha 0.6 over forever.teal. Focus-visible: 1 px forever.cream outer ring. Disabled: alpha 0.4. Loading: spinner line icon spinner.tga 12 px beside the shipped busy label, rotating 1 turn per 1.0 s (a texture rotation through an Animation, capability-checked; else the label alone).
- Secondary: forever.blue fill, 1 px forever.blueEdge, same type. Hover fill #17608a (addon.blueHover, new). Other states as primary.
- Tertiary links (text only, "Full log"): forever.accent text, base.strong on hover, underline 1 px on hover and focus. The shipped gold link colour is retired; gold is for the dot, the score, and item-level emphasis only.
- Shipped primary and secondary assignment per tab is kept. Sim again stays secondary blue as shipped (spec 2026-11-13-addon-hub.md); the next action on the Me tab is the first "Right now" row (see section 9).

Progress tracks (Forever Score bars, build progress):
- Track: base.lineSoft #2c241f, 6 px (Theme.SIZES.progressHeight), 1 px inner shadow addon.inset at the top edge.
- Fill: tier colour as shipped (the site's rarity text colours, Theme.PANEL_TIER). A 1 px forever.frameLine rgba(243,238,226,0.42) highlight along the fill's top edge at alpha 0.5 gives the painted lip. Gain deltas use forever.gain #32ae88.

Footer (28 px, y bottom):
- Plate forever.foot at 0.85, a 1 px forever.panelEdge line on top.
- Left: "Esc to close" in base.muted 11 px, as shipped. Right: sync text as shipped, base.muted. A 12 px seal at the far left is NOT added (the strip has it).

## 3. Icon set

Source: the site's SVGs. Generator: a new web/scripts/gen-addon-icons.mjs, sitting beside web/scripts/gen-ui-icons.mjs and importing its parsing helpers (gen-ui-icons.mjs exports its svg loading and the line-family grid, LINE_BOX 24; the build lane exports what it needs). It exports a list `ADDON_ICONS`, the same pattern as `COMPANION_ICONS`. Rasteriser: @resvg/resvg-js (already in web/package.json). Outputs: addon/ForeverSixty/media/icons/<name>.tga, 32 x 32, 32-bit uncompressed, no RLE, white glyph on transparent, 4,096 bytes each plus an 18 byte header. They are drawn at 14 px (tabs) or 12 px (row glyphs) with the 32 px source for crisp minification, and tinted at runtime with SetVertexColor from tokens (tabOff, tabOn white, accent, gain, warn, muted).

ADDON_ICONS (initial, 18 glyphs, about 74 KB):

| Addon spot | Icon | Family |
|---|---|---|
| Tab Me | character | filled |
| Tab Builds | planner | filled |
| Tab Gear | bis | filled |
| Tab Rotation | sim | filled |
| Tab Guild | guild | filled |
| Tab Officer | manage | filled |
| Header close | close | line |
| Row opens a detail (Right now rows, build rows) | chevron-right | line |
| Disclosure open | chevron-down | line |
| Verified, on plan, flask up | check | line |
| Upgrade in bags (gear rows) | upgrade | line |
| Died, dead early (logs) | skull | line |
| Slot target, "Aim here" | target | line |
| Refresh (re-read, Sim again glyph) | refresh | line |
| Haste and procs (rotation rows) | bolt | line |
| Pinned build | pin | line |
| Search box | search | filled |
| Busy | spinner | line |

Also written by the same generator: media/seal.tga (32 x 32, from web/public/art/seal-flat-small.svg, tinted by its own colour forever.seal), media/filigree.tga (64 x 64, from design/assets/ornaments/corner-filigree.svg), media/ground.tga, media/band-*.tga, media/cave.tga, media/crests/*.tga (generator replaces addon/tools/make_class_crests.py; delete that script in the build lane; no duplicate path).

Guard test plan (Lua spec under addon tests, plus a node check in CI):
1. Every `Interface\\AddOns\\ForeverSixty\\media\\...` path referenced in addon/ForeverSixty/**/*.lua exists in the generator's manifest (addon/ForeverSixty/media/MANIFEST.json, written by the generator, committed).
2. Every file in media/ except fonts/ and legacy minimap.tga and arrow.tga is in the manifest, and is byte-identical to a fresh generator run (a stale-output test, same shape as design/tokens/tokens.test.mjs).
3. Any other `Interface\\` path in addon Lua fails the test unless it is on the allowlist: Interface\\Icons\\ (item and spell icons), Interface\\Buttons\\UI-CheckBox-Check, Interface\\CharacterFrame\\TempPortraitAlphaMask, Interface\\Minimap\\MiniMap-TrackingBorder, Interface\\Minimap\\UI-Minimap-Background, Interface\\RaidFrame\\ReadyCheck-Ready and ReadyCheck-NotReady (until replaced by check and close glyphs; the build lane retires the last two and removes them from the list), Interface\\TargetingFrame\\UI-Classes-Circles (retired, see section 4), and Interface\\Tooltips\\* (native tooltip). The allowlist lives in one Lua table in Theme.lua (Theme.NATIVE_PATHS) and the test reads it.
4. Every texture is power of two on both sides, 32-bit, uncompressed (header byte 2 = 2, bits per pixel = 32), and the total media/ size (excluding fonts) is at most 2.2 MB.

## 4. Crests

- Source: web/public/icons/hd/crests/<class>.webp (128 x 128, circular on transparent).
- Output: media/crests/<class>.tga at 128 x 128, 32-bit, 65,554 bytes each, nine classes: 590 KB. Plus media/crests/ring.tga 128 x 128 (white 3 px ring at 128 scale, transparent inside and out), 65 KB, tinted per class by Theme.classColor as shipped.
- Where a class or character appears, always the ringed crest: header (40 px), Guild and Officer rows (24 px), player tooltip (the crest as an inline texture, 16 px, `|T...:16:16|t`), the "Right now" and gains rows where a class shows, the minimap button remains the seal. The shipped 64 px files are replaced, so 40 px draws at 3.2x oversample and 24 px at 5.3x.
- The retired Theme.MEDIA.classIcons (UI-Classes-Circles) is removed from every draw path: no square, no Blizzard avatar (one-crest-language rule). The guard test greps for it.
- Crest states: default full colour; offline guildmate alpha 0.5 (shipped ALPHA.dim); no loaded-crest flash: the textures are preloaded when the hub frame is built, each in a hidden 1 px texture, so the first show never draws a blank ring.

## 5. Header art per tab

Each tab swaps the band texture and its veil. The swap is a SetTexture on one header texture, instant, no fade (a fade over 56 px reads as flicker). Art is 512 x 64 TGA, 131,090 bytes each. Five textures, 640 KB.

| Tab | File | Source and crop | Mood | Veil over |
|---|---|---|---|---|
| Me | band-me.tga | band-2400.webp, full width, y 30 to 52 percent (navy into teal cloud) | the site's own band | veilBand 0.38 |
| Builds | band-builds.tga | band-2400.webp, mirrored, y 48 to 70 percent | band, cooler | veilBand 0.38 |
| Gear | band-gear.tga | paperdoll-frame@2x.webp, x 22 to 78 percent, y 42 to 50 percent (the painted teal inside the frame, no knotwork) | paperdoll teal | veilNight rgba(7,29,46,0.3); raise to 0.4 if the white figures line falls under 4.5:1 in game |
| Rotation | band-rotation.tga | band-2400.webp, y 62 to 84 percent (teal deepest) | band, deeper teal | veilBand 0.38 |
| Guild and Officer | band-cave.tga | cave-1600.webp, x 30 to 100 percent, y 12 to 40 percent (dusk sky and the cave mouth's glow at the right) | cave, raid plaque | veilNight rgba(7,29,46,0.3), so the navy keeps the brand blue |

Guild and Officer share band-cave.tga. Officer is told apart by a 2 px forever.ruleWarm line at the header's bottom edge instead of the accent line. These are the crops the approved mock uses (design/mocks/addon_brand.py BAND_CROPS); the generator copies that table.

Where a header or icon gives a panel more to say (no new information architecture, only room for what we already hold):
- Me: the "Right now" lead plate carries the corners and a 10 px chevron-right after each row's tab word, in forever.accent, so rows read as openable. The pulse bar sits on the plate's inner left edge, 9 px clear of the row's icon.
- Gear: the Sources ticks are forever.accent with a white check (the shipped gold squares were the loudest thing in the window). The slot list header row gets the target icon beside "Aim here" when the shipped sentence names a slot.
- Rotation: rows carry bolt for procs and cooldowns; the shipped text is unchanged.
- Guild and Officer: roster rows carry the 24 px crest, the shipped rank text, and the check icon in `forever.gain` where the shipped data says verified.

## 6. What stays native

- The client's tooltip frame, border, font and background. We never skin GameTooltip, ItemRefTooltip or the comparison tooltips.
- Item quality colours on item names (the client's ITEM_QUALITY_COLORS, via Theme.QUALITY_HEX), and the 1 px quality border on item icons.
- Item and spell icons: `Interface\Icons\*` with the client's own border art. Icons keep the shipped 20 to 24 px size. Item icons get no brand frame.
- The client's own cursor, scroll bars and edit boxes keep native behaviour, with our colours on the parts we draw (shipped).

## 7. Item tooltip: our block inside the native frame

The frame stays native. Our lines (ItemTooltip.lua) are text added to it, so the brand can only arrive through text colour, one inline texture and line order. Per-line font changes are impossible on the client's tooltip; do not attempt them.

Block, appended after the item's own lines, one blank spacer line first:
1. Header line: inline seal (media/seal.tga, `|T<path>:12:12|t`) then " Forever Sixty" in forever.accent #0093b2 (about 5:1 on the client's dark tooltip ground; forever.seal #2c8a9e is lower, so it is used only in the seal art, which keeps its own colour).
2. Our lines, in this colour grammar (site tooltip tokens, forever.tooltip* #060914 ground, #8b8f9c edge, #2b3040 inner, #ffd100 yellow, as the site's bis-tooltip-1440.png):
   - Verdict gain, the shipped words kept: "Upgrade +2.7 DPS over Trueaim Gauntlets": the arrow glyph and the gain in forever.gain #32ae88 (the site's green; replaces addon.tooltipGreen #1aff1a, which is the client's own item-green and so collides with the Equip lines and green item names), "Upgrade" in addon.tooltipText, "over" in addon.tooltipMuted #b8b8b8, the compared item name in its quality colour.
   - The attribution line drops its "Forever Sixty ·" prefix (the header line now says it): "Beast Mastery, your sim" in addon.tooltipMuted.
   - A loss: forever.warn #d66e28.
   - Labels ("Source", "Slot"): addon.tooltipMuted. Values: addon.tooltipText #d0d0d0.
   - The next step ("Alt-click: open Gear"): forever.tooltipYellow #ffd100, as the site's own tooltip draws "Open in the sim".
   - Font: the client's own tooltip face throughout. The mock has no copy of it and draws Marcellus as a stand-in, labelled on the board.
3. No rules, boxes or background fills added. Maximum added lines is the shipped count; the brand adds the one header line only if the block has at least one other line. An item with nothing to say gets nothing (no empty "Forever Sixty" line).

Player tooltip (PlayerTooltip.lua): the same header line. The 16 px crest is not put on the name line (it would break the client's name colouring); it goes on the second line, `|T<crest>:16:16|t` before the "Level 60 Hunter" text, which the shipped tooltip already prints.

## 8. New addon-only tokens

Add to the color.addon group in design/tokens/tokens.json (type color, native false, no css). Regenerate with node design/tokens/build.mjs; never hand-edit Tokens.lua.

| Token | Value | Why |
|---|---|---|
| color.addon.plate | rgba(23,14,9,0.84) | base.raised at the layer-5 alpha, so the earth warms every plate and the window reads brown, not slate |
| color.addon.zebra | rgba(255,255,255,0.03) | odd row tint on dense tables |
| color.addon.tabHover | rgba(255,255,255,0.06) | nav hover plate |
| color.addon.selectedActive | rgba(0,147,178,0.24) | selected row while its detail is open |
| color.addon.tealHover | #12727a | primary button hover (teal lifted) |
| color.addon.blueHover | #17608a | secondary button hover (blue lifted) |
| color.addon.headerFade | rgba(11,7,5,0.35) | the 260 px left fade behind crest and name (addon.inset at alpha) |

Size constants are Lua, not tokens: they go in Theme.SIZES (window side margin 10, plate gap 6, plate padding 8, lead top padding 15, tab width 117, crest 40, tab icon 14, row glyph 12, chevron 10, strip seal 20, filigree 28). Theme.SIZES.windowWidth still reads 720 in the file I read; the hub spec set 740, so the build lane confirms the hub reads 740 and removes any stale 720.

## 9. Motion

One pulse, for the next action, on the Me tab and the tab strip together.
- Where: the 2 px accent bar on the left edge of the first "Right now" row (the shipped top item), and the waiting dot on the tab that row belongs to.
- How: forever.accent at alpha 1.0 to 0.35, BOUNCE, 0.6 s each way (Theme.GLOW_PULSE, the shipped constants), one AnimationGroup per texture, both started in the same frame.
- Duration and stop: 4 full cycles (4.8 s) from the hub's first show in a session, then it settles to alpha 0.6 steady. It also stops at once on any of: the row is clicked, the tab is opened, the window hides. It does not restart in the same session.
- Only one pulse at a time anywhere in the hub. The talent and slot glows outside the hub are unchanged.
- No other motion except the shipped 0.15 s window fade and the spinner. A client without animations shows the steady state (alpha 0.6): every step is guarded, as Theme does at line 1667.

## 10. Phone and desktop

Not applicable to the client: the window is one size (740 px) and scales with the client's UI scale. At UI scale 0.64 to 1.0 the layout holds. What never collapses: the header crest, the strip, the lead plate and the footer. What the strip does when tab labels overflow 704 px (six tabs, long locale): labels drop and the 14 px icons stay, with each tab's name as its tooltip (native tooltip).

## 11. Performance and polish limits

- Texture memory: about 1.76 MB new media (table in section 12). Load: every texture is created once at hub build, never per show.
- No layout shift: header, strip and footer heights are fixed (56, 28, 28); art swaps do not resize. No missing-icon flash: the generator manifest is the only source of paths, preload as in section 4.
- No per-frame work: no OnUpdate for the skin. The pulse is an AnimationGroup.
- Frame count: the skin adds at most 3 frames (header, strip seal button, plate backdrop host per plate reuses the shipped card frames). Textures only otherwise.
- Web Lighthouse budgets do not apply to this surface. The addon equivalent is: hub open under 50 ms on a 2019 laptop, measured with debugprofilestop around Hub.show, shipped baseline plus at most 15 ms.

## 12. Build-lane plan

Files touched:
- design/tokens/tokens.json (7 tokens above), regenerate addon/ForeverSixty/Tokens.lua and the other token outputs.
- web/scripts/gen-addon-icons.mjs (new): ADDON_ICONS, BAND_CROPS, crest, ground, filigree, seal, cave, manifest. A `--check` mode for CI.
- addon/ForeverSixty/media/** (new files below), media/MANIFEST.json.
- addon/ForeverSixty/Theme.lua: HEX roles (plate, zebra, tabHover, selectedActive, tealHover, blueHover), SIZES, MEDIA, NATIVE_PATHS, a Theme.plate(parent) helper (fill + border + title rule) and Theme.ground(frame) with the tiling capability check.
- addon/ForeverSixty/Widgets.lua: tab (icon, underline, hover), button states, row (zebra, hover, selected bar), progress track.
- surfaces/HubFrame.lua, HubDraw.lua, Hub.lua: ground, header band swap, strip seal, footer plate.
- surfaces/MeTab.lua, MeFrame.lua, GearTab.lua, GearTabView.lua, BuildsPane.lua, RotationTab.lua, GuildTab.lua, GuildHubTab.lua, OfficerTab.lua, PanelTab.lua: replace card creation with Theme.plate; row constructors; lead plate ornament; no content change.
- surfaces/ItemTooltip.lua, PlayerTooltip.lua: header line, colour grammar.
- Delete addon/tools/make_class_crests.py (generator replaces it) and Theme.MEDIA.classIcons.

New textures (32-bit TGA, bytes = w x h x 4 + 18):

| File | Size | KB |
|---|---|---|
| ground.tga | 256 x 256 | 256 |
| band-me, band-builds, band-gear, band-rotation, band-cave | 512 x 64 each | 128 each, 640 |
| cave.tga (plate backdrop) | 256 x 128 | 128 |
| filigree.tga (one corner, mirrored to the two top corners) | 64 x 64 | 16 |
| seal.tga | 32 x 32 | 4 |
| icons/*.tga (18) | 32 x 32 | 4 each, 72 |
| crests/*.tga (9, replacing 64 px) | 128 x 128 | 64 each, 576 |
| crests/ring.tga | 128 x 128 | 64 |
| Total new | | about 1,760 KB (net of the replaced 64 px crests about 1.65 MB) |

Budget ceiling for media/ excluding fonts: 2.2 MB. If a driver reports a hitch, the fallbacks in order: ground to 128 x 128 (64 KB), bands to 256 x 32.

Tests (touched only, per the team's rule):
- addon: the guard tests in section 3; a Widgets spec per state of tab, button, row; a Hub spec for the band swap per tab and the pulse start and stop (animation group double); ItemTooltip spec for the header line absent when no other line.
- design/tokens/tokens.test.mjs stays green after regeneration.
- web: gen-addon-icons `--check` stale-output test.
- Lua tests run through the existing addon suite entry; do not run full suites locally.

Reviewer gates before merge (all on rendered screenshots, not the diff): ux-designer, wow-player, brand-consistency-reviewer against the live site page, ai-slop-reviewer, info-density-reviewer. Guild and Officer surfaces also pass guild-officer-reviewer.

Acceptance screenshots the implementer attaches (in-client at 740 px, and the mock at the same scale): Me, Builds, Gear, Rotation, Guild, Officer tabs default; the Me tab with the pulse at its low and high frames; hover and selected on one row; focus-visible on one tab and one button; primary and secondary buttons default, hover, disabled; a long-locale strip with icons only; one item tooltip with the block and one without; a player tooltip; the client at UI scale 0.64 and 1.0; the window with a failed texture load (fallback flat ground and no art) to show it still reads.

## 13. Copy

No new visible copy except tooltips and the tooltip header line.
- Seal tooltip: "Forever Sixty"
- Tooltip block header: "Forever Sixty"
- Icon-only tab tooltips: the shipped tab labels ("Me", "Builds", "Gear", "Rotation", "Guild", "Officer").
- Close button tooltip: "Close" (shipped footer line "Esc to close" is kept).
- Failed-art states have no copy: the window falls back silently to flat ground and a flat header.

## 14. Mock (this lane's second deliverable)

Built by design/mocks/addon_brand.py over the store fixture (design/listing/gen_captures.py), output in design/reviews/2026-11-14-addon-brand/mock/: brand-{me,gear,builds,rotation,guild}.png at 2x, tooltip.png, band-*.png, key.png, and side-by-side.png (Me, Gear, Builds) and side-by-side-2.png (Rotation, Guild) at 1x CSS scale with the colour key under each. The mock must show, at the real 740 px width and at 1x and 2x:
1. The hub on Me, Gear and Builds tabs with every section 2 surface: ground tile in the margins, header band with crest and name, nav strip with seal, icons and accent underline, plates with Marcellus titles, filigree on the lead plate, zebra, one hover row, one selected row, teal and blue buttons, progress tracks, footer.
2. One item tooltip (client-style frame, our block with the seal header line, gain green, loss orange, yellow next step) next to the shipped one.
3. Each composition placed side by side at the same scale: left the shipped window (design/listing/captures/window-me.png, window-gear.png, window-builds.png), middle the mock, right the live site page (design/reviews/2026-11-14-addon-brand/site/bis-1440.png, home-1440.png scaled so the nav bar heights match the strip's 28 px band where compared).
4. A second sheet with the Guild tab (cave header, lead plate) and Rotation (deeper band) so every band crop is seen.
5. A colour key under each sheet listing the tokens used per region.
The mock's crops feed BAND_CROPS: the build lane takes them from the approved mock, not from this spec's starting percentages.

## 15. Open decisions for the owner

1. Ground: painted tile (this spec) or flat forever.ground. The tile costs 256 KB and is mostly hidden by plates; flat is cleaner and cheaper.
2. Seal alone in the strip, or the Marcellus "Forever Sixty" wordmark beside it (costs about 110 px of the 704 px strip and pushes the Officer tab to an icon-only label at six tabs).
3. Gold: kept only for the waiting dot, the score and item-level emphasis (this spec), or kept for links as shipped.
4. Whether the player tooltip gets the crest line (the spec adds it on the second line).
5. Whether Guild and Officer share one cave band (this spec) or Officer gets its own.

## 16. Owner rulings (2026-11-14)

The owner approved the round 3 mock ("it looks good"). The five open calls in section 15 are settled:
1. Ground: the painted tile, as in the mock.
2. Strip: the seal plus the Marcellus "Forever Sixty" wordmark while every tab label fits beside it at the live window width, else the seal alone. The live width is 740 (Hub.width; Theme.SIZES.windowWidth 720 is an unused legacy constant). With tabs as wide as their glyph, label and dot slot, the wordmark fits.
3. Gold only for the waiting dot and the score; links, the active tab, a sorted column and the followed build use the site's accent.
4. The player tooltip gets a crest line: our lines open on the player's ringed crest and "Forever Sixty" in the accent.
5. Officer shares Guild's cave band; its header join is the warm rule so the two read apart.

## 17. As built (2.0.55)

Where the build differs from the sections above, the build is the record:
- Tokens: every alpha the skin paints is a token in the addon group (plate 0.84, footer 0.85, zebra, rowHover, selectedActive, tabHover, tabActive, headerFade, headerJoin, corner, caveFadeFrom 0.9, caveFadeTo 0.5, titleRule and headerNote referencing the companion's veilRule and inkSoft), plus tealHover and blueHover. Tokens.lua carries each alpha under Tokens.alpha and Theme.paint applies it (Theme.HEX_ALPHA). The client's tooltip green (addon.tooltipGreen) is gone: gains on tooltips, reward and roll lines are forever.gain.
- Geometry (Theme.SIZES): hubSide 11, hubMargin 6, plateGap 6, platePad 8, platePadFirst 15 on every column's first plate, plateSectionGap 24, corners 20 px flush on the plate edge (the two top ones on the lead plate only; 28 px inset 3 crossed the first line), groundRepeat 640 (the 256 px tile is the whole site tile, so it repeats without a seam). The columns keep the shipped 388 and 288 px, so no shipped string is cut. Section titles stay 14 px (15 cut the Guild leaderboard's note); the window title is 18.
- Plates: surfaces/HubPlates.lua. A tab's draw marks each section's start (BuildsWindow.drawSection, HubDraw.heading, each Builds spec block, the Gear level row); Hub.fitHeight lays the plates out from the marks still showing, the last one to the body's bottom. Lead plates: Me Right now, Builds the first spec block, Gear the slot list, Rotation Priority, Guild Tonight and Officer Roster or a loot session (the raid ones over the cave art).
- Strip tabs are as wide as their glyph, label and dot slot (fixed per tab, one label weight in every state), not 117 px each.
- Item tooltip: one blank row (ItemTooltip.BLOCK_GAP), then the header line, after the client's lines; the Shift attribution drops its "Forever Sixty ·" ("Beast Mastery, your sim"), so the name shows once; the Alt-click hint is the client tooltip gold; the upgrade arrow is the site's upgrade glyph tinted forever.gain through the texture escape (media/arrow.tga and addon/tools/make_arrow.py are retired).
- Retired client art: the ready-check tick (Onboarding's done step is the check glyph in forever.gain), UI-Classes-Circles and Theme.classIcon, and addon/tools/make_class_crests.py (the generator writes the crests).
- Generator: web/scripts/gen-addon-icons.mjs writes 37 textures, 1,769 KB, and media/MANIFEST.json. Its --check (and web/src/lib/ui/addon-media.test.ts) compares the sha256 of every source, the generator and each committed texture with the manifest, so it holds on any platform. The seal is 64 px. addon/tests/media_guard_spec.lua is the Lua guard with its own planted-path check.
- Forever Score bars: every part's fill is forever.accent on the base.lineSoft track; the tier colour stays on the score figure (build gate: the client lime read as the default UI).
- Motion: Me's first Right now row bar and its tab's dot pulse four times on the first show in a session (Theme.pulseNext) and hold at 0.6; leaving the tab or clicking the row ends it.
- Captures: design/mocks/addon_brand_build.py renders the boards with the committed textures and the Lua's own Theme.SIZES and Hub.LAYOUT to design/reviews/2026-11-14-addon-brand/build/.
