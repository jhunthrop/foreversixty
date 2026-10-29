# Discord bot amendment: the gear panel and the character sheet: experience spec

Author: ux-designer, 2026-10-29. Amends `design/specs/2026-10-27-discord-bot.md` (the bot spec) §3.3.3 and §3.3.4, and `design/specs/2026-10-28-discord-guided-cards.md` (the guided-card spec) §G.4.2, §G.4.3, §G.5 and §G.6. Tenets: `docs/tenets.md` 1, 2, 4, 7, 8, 10, 11, 13, 14.

Source (the owner, 2026-10-29, on a production `/bis` answer of five stacked item embeds, verbatim): "can we return a generated output that looks like the character screen with hover over or the items names beside them? Does discord rendering allow that?" The coordinator's answer, approved by the owner with "do it": Discord has no hover, so do both. Part A replaces the `/bis` and `/char` item answers with one compact Components V2 panel. Part B posts a server-rendered character-sheet PNG on `Share in channel`.

Mock: `design/mocks/2026-10-29-discord-gear-panel/gen_gear_panel.mjs` (with `data.mjs`, `payloads.mjs`, `sheet.mjs`). Run it with Node 22 or later. The reviewers' render kit (`web/tests/captures/discord-bot/discord-render.ts`) draws legacy embeds only, so the mock carries its own Components V2 drawing in the kit's dark-theme colours. Every item, icon, quality, item level, source line and DPS figure on the boards is real: the active build `1.60.1.70658`, Beast Mastery Hunter, level 60 Phase 1 raid, Alliance, read from `data/builds/1.60.1.70658/bis/hunter-beast-mastery.json`, `items/hunter.json` and `loot.json`. The source lines go through a port of the bot's own `describeSource`. The example character Zulmara wears the band's own alternatives, so each of her gains is the band's measured delta. The reference payloads are in `payloads/*.json`, each with its component and text count. Boards are drawn at 640 (desktop client) and 390 (phone client), as `capture-messages.spec.ts` does.

| File | What it shows |
|---|---|
| `bis-panel-p1-{desktop,phone}.png` | `/bis`, Beast Mastery Hunter, Phase 1 raid: Head to Legs, 9 slots |
| `bis-panel-p2-{desktop,phone}.png` | The same answer after `Feet to Ranged`: 8 slots, with the off hand empty because the main hand is two-handed |
| `bis-panel-slot-{desktop,phone}.png` | `/bis slot:legs`: the pick and the next two |
| `char-panel-{desktop,phone}.png` | `/char` for a linked member: header, five upgrades (an empty slot first), `4 more upgrades` |
| `char-panel-nogear-{desktop,phone}.png` | `/char` with no gear shared |
| `share-bis-{desktop,phone}.png` | The public message after `Share in channel` on `/bis`: the sheet attached |
| `share-char-{desktop,phone}.png` | The same for `/char` |
| `share-bis-nosheet-{desktop,phone}.png` | The fallback when the sheet cannot render: the panel, public, without the sheet |
| `sheet-bis.png` | The `/bis` sheet PNG at its real size, 1080 x 1616 |
| `sheet-char.png` | The `/char` sheet PNG at its real size, 1080 x 1616 |

No new site token and no new site control (tenet 10). The Discord side uses only Discord's own components. The sheet is drawn only from the site's tokens, fonts, quality colours, HD crest and faction mark.

# 0. Discord facts, verified

Source: `discord/discord-api-docs`, branch `main`, fetched 2026-10-29 (the last commit to `developers/components/reference.mdx` is dated 2026-10-28T22:31Z). The files read were `developers/components/reference.mdx`, `developers/components/overview.mdx`, `developers/components/using-message-components.mdx`, `developers/resources/message.mdx` and `developers/change-log.mdx`. These files are the source of https://discord.com/developers/docs/components/reference. Each line below cites its file.

| Fact | Value | Source |
|---|---|---|
| The flag | `IS_COMPONENTS_V2 = 1 << 15` (32768), set per message | message.mdx, message flags table; reference.mdx intro |
| Content and embeds | Forbidden once the flag is set. "Providing `content`, `embeds`, `sticker_ids`, `poll`, or `shared_client_theme` will fail with a 400 BAD REQUEST" | message.mdx, Create Message note ** |
| The flag is one-way | "Once a message has been sent with this flag, it can't be removed from that message." An edit can set it, never unset it. Setting it on an edit needs `content: null` and `embeds: []` | reference.mdx intro; message.mdx, Edit Message note * |
| Component limit | "Messages allow up to 40 total components." The changelog says "total components", and removed both the top-level limit and the limit on a Container's children (2025-04-29). The spec therefore counts every component, nested ones included | reference.mdx intro; change-log.mdx 2025-04-29 |
| Legacy messages | Still allow up to 5 action rows | change-log.mdx 2025-04-29 |
| Container (17) | `components`, `accent_color` 0x000000 to 0xFFFFFF (nullable), `spoiler`. Children may be Action Row, Text Display, Section, Media Gallery, Separator or File | reference.mdx, Container |
| Section (9) | "One to three child components", which may only be Text Display. `accessory` is one Button or one Thumbnail. The docs' own example draws the accessory to the right of the text | reference.mdx, Section, and its example image `section.webp` |
| Text Display (10) | `content`: markdown "similar to a message". The reference states no per-component length limit | reference.mdx, Text Display |
| Thumbnail (11) | `media` (unfurled media item: an `https` URL or `attachment://<filename>`), `description` (alt text, max 1024 characters), `spoiler`. Images only, GIF and WebP included. Usable only as a Section accessory. **The reference states no displayed size** | reference.mdx, Thumbnail |
| Media Gallery (12) | 1 to 10 items, each with `media`, `description` (alt text, max 1024) and `spoiler` | reference.mdx, Media Gallery |
| Separator (14) | `divider` (default `true`), `spacing` 1 small or 2 large (default 1) | reference.mdx, Separator |
| Unfurled media item | Only `url` is settable. Discord fills in `width`, `height`, `proxy_url` and `content_type` | reference.mdx, Unfurled Media Item |
| Attachments | "Attachments won't show by default - they must be exposed through components." Upload as `multipart/form-data`. An attachment's `description` (alt) is max 1024 characters | reference.mdx intro and "Uploading a file"; message.mdx, attachment object |
| Button | `label` max 80 characters. The design note asks for at most 38 without an icon. `custom_id` 1 to 100. `url` max 512 | reference.mdx, Button |
| String select | Up to 25 options. Label, value and description max 100. Placeholder max 150 | reference.mdx, String Select |
| Where the flag works | Message Create, Execute Webhook, and interaction responses and follow-ups | using-message-components.mdx |

Two limits are **not in the docs**, and the spec does not guess either one:

1. **Total text length.** No file in the docs states a limit on text across Text Displays (searched for `4000`, `4,000` and `characters` in all five files). Community libraries validate 4,000 characters across a message's text. The builder guards at **3,800** (§A.6), and capture D21 sends 4,001 characters to the test server to record what Discord does.
2. **Thumbnail and gallery display sizes.** Discord sizes these in the client, and the docs give no figure. The mock draws a thumbnail at 72 px (desktop) and 64 px (phone), and the gallery image at the container's full inner width. Capture D20 measures both clients. The rules in this spec do not depend on the exact figure (§B.6).

Consequence for the guided cards (binding): because the flag is one-way, a message that becomes a V2 answer can never go back to an embed card. Today `Change picks` turns the `/bis` answer back into the card in place (guided-card spec §G.5.3). So the `/bis` card and the `/char` cards are V2 from their first frame (§A.7). Every edit in those flows writes `flags` with bit 15 set.

# Part A. The compact panel (`/bis`, `/char`)

## A.1 The one question, and the answer on the first screen

| Viewer, surface | Question | Visible with no scroll, phone client (390) and desktop client (640) |
|---|---|---|
| Anyone, `/bis` with a spec | "What is the best item in each slot, and where do I get it?" | The header (spec, class, band, faction, crest) and the first four slots with icon, linked name, slot, gain and source on the phone. All nine slots of the first page on the desktop. Today's five embeds held five slots in a taller message. |
| Anyone, `/bis slot:{slot}` | "What is best in this one slot, and what comes next?" | The pick and both alternatives on the phone (`bis-panel-slot-phone`). |
| A player, `/char` | "What should this character replace first?" | The character line, item level and readiness, and the first three upgrades on the phone, each saying what it replaces. |

Tenet 14: the panel leads with the class and the band (`/bis`) or the character (`/char`). The first item row is the next step.

## A.2 References, and what we borrow

- **The in-game character frame.** Borrowed: the slot order (left column Head, Neck, Shoulder, Back, Chest, Wrist; right column Hands, Waist, Legs, Feet, Ring 1, Ring 2, Trinket 1, Trinket 2; bottom Main hand, Off hand, Ranged). The panel's two pages follow it, so a player finds a slot where the game puts it.
- **Wowhead's BiS table and our own BiS page's "The list"** (`/bis/{class}/{spec}`). Borrowed: one row per slot holding the slot, the icon, the item name in its link, and the source line in the page's own words (`describeSource`, unchanged). Not borrowed: the runners-up column, which stays one tap away on the page (`All 17 slots`).
- **The game's item tooltip.** Borrowed: the quality word and the item level, kept as a quiet line under the name, the way the tooltip lists them under the title.
- **Our own `/char` card today** (bot spec §3.3.3). Borrowed unchanged: the character line `Level {n} {Spec} {Class} · {Guild}`, item level, readiness and synced. Added: what each upgrade replaces, the comparison the game's tooltip makes.
- **Discord's own Section pattern** (the docs' changelog example: text with a thumbnail beside it). Borrowed: one Section per item, with the icon as the Thumbnail accessory. Discord draws the accessory on the right. The text keeps the left edge, which is where a reader starts.

## A.3 The item row (one grammar, both commands)

Each item row is one Section (type 9). It holds one Text Display (type 10) and a Thumbnail (type 11) accessory. That is 3 components.

Text Display content, three lines (four for a `/char` upgrade):

```
**[{Item name}]({link})**
{Slot} · {gain}
-# {Quality} · item level {n} · {source line}
-# Replaces {worn name}, item level {n}          (/char only; absent for an empty slot)
```

- Line 1 is the only primary element in the row: the name, bold, a masked link (Discord draws it in its link blue #00a8fc). Markdown cannot colour text, so the item's quality colour is carried by the tile (§A.4) and the quality word.
- Line 2 sits at body size. It holds the slot (the `/bis` slot choice's name, `slotLabel`) and the gain. `/bis`: the pick's gain over the next best, `signedDPS(-alternatives[0].dps_delta)`. A gain that rounds to 0.0 reads `ties the next best`. This replaces today's `+0.0 DPS`, which `signedDPS(-0)` prints. A slot with no alternatives shows the slot alone. `/char`: the band's measured gain over the worn item, or `fills an empty slot` (§A.5.2). A tank's gain reads in points (`+12.4 points`, unchanged).
- Line 3 is Discord subtext (`-# `, 13 px, #949ba4): the quality word, the item level, and the source line from `sourceLine`, unchanged. Each part is left out when unknown, as `joinPresent` does today.
- Line 4 (`/char`) is subtext: `Replaces {worn name}, item level {n}`. When the worn item's level is unknown: `Replaces {worn name}`. When its name is unknown too, the line is absent.
- Escaping: every name goes through `memberText`. Inside a masked link, the characters `[`, `]`, `(` and `)` in an item name are escaped with a backslash. A test pins `Ashjre'thul, Crossbow of Smiting` and an item name holding brackets.
- `{link}`: `/bis` uses `itemLink(page, id)` as today. Each link stays unique for consistency, although V2 does not fold embeds. `/char` uses `{character URL}#gear-slot-{slot}`, as today.

The old `ItemEmbed` grammar (bot spec §3.3.3) is retired for `/bis` and `/char`. `/report` keeps it until its own amendment.

## A.4 The tile: the icon with its quality ring

The Thumbnail's URL is a tile the API draws: `{API_BASE_URL}/v1/discord/tiles/{build}/{item id}.png`.

- The tile is a 72 x 72 PNG. A 4 px ring in the item's quality colour (the game's own colours: `--color-rarity-*` #9d9d9d, #ffffff, #1eff00, #0070dd, #a335ee, #ff8000, never the lightened `-text` variants) surrounds the item's icon at 64 x 64, 1:1, never resampled (tenet 2, the faithful-icons rule). Discord rounds the corners itself.
- The icon comes from `data/builds/{build}/icons/{icon}.webp`, the same file the site serves.
- An item the build does not know gets `inv_misc_questionmark` (the game's own icon, present in the build) inside a #262e40 ring. The row's text then leaves out the quality word and the item level, as today.
- Alt text (`description`): `{Item name}, {Quality} {slot} item` (`Beastmaster's Cap, Epic head item`). Without a known quality: `{Item name}, {slot} item`.
- The header Section's accessory is the class's ringed HD crest, `{PUBLIC_BASE_URL}/icons/hd/crests/{class}.webp` (one crest language), with the alt text `{Class} crest`.

Why a drawn tile and not the bare icon URL used today: legacy embeds carried quality in the colour bar, one item per embed. A V2 panel has one accent colour for the whole container. Tenet 2 needs the quality colour on every item, and the ring is where the game itself draws it (the bag slot and the character frame).

## A.5 Layouts

### A.5.1 `/bis` with a spec (`bis-panel-p1`, `bis-panel-p2`)

Ephemeral. `flags` = 64 | 32768 = 32832. Top to bottom:

1. **Container**, `accent_color` the class colour (`classColors`).
   1. **Header Section**. One Text Display:
      ```
      ### {Spec} {Class}
      Best in slot · {band label} · {Faction}
      -# {Page range}, {n} of 17 slots. Each gain is over the next best item.
      ```
      The accessory is the class crest. `{band label}` comes from `bandLabels` (`level 60 Phase 1 raid`). `{Faction}` is `Alliance` or `Horde`, the band's faction. Today the faction is silent, but the picks differ by faction.
   2. **Separator**, divider, spacing 1.
   3. **The slot rows of the page**, each per §A.3. Page 1 (`Head to Legs`): Head, Neck, Shoulder, Back, Chest, Wrist, Hands, Waist, Legs (9). Page 2 (`Feet to Ranged`): Feet, Ring 1, Ring 2, Trinket 1, Trinket 2, Main hand, Off hand, Ranged (8).
   4. A slot with nothing to show is one Text Display with no Section, so there is no tile to fake:
      - Off hand with `empty_reason: two_hand_equipped`: `**Off hand** · none: {main hand name} is two-handed.`
      - Any other slot with no pick (`item_id` 0): `**{Slot}** · no pick for this range.`
2. **Action row** (outside the Container), in this order:
   - `Share in channel` (Secondary)
   - the page button (Secondary): `Feet to Ranged` on page 1, `Head to Legs` on page 2
   - `Change picks` (Secondary)
   - `All 17 slots` (Link to the band page)

   No Success button: the panel is an answer and has no forward action (guided-card ruling G7).

The page button's custom id is `gbis:1:pg:{class}:{spec}:{phase}:{slot}:{page}`. The longest, `gbis:1:pg:druid:feral-bear:p1::2`, is 32 characters. A press answers type 7 with the other page. It reads the message's existing `Share in channel` button from the interaction's `message.components` and keeps it as it is. A page press mints no Share token, and Share always shares all 17 slots (Part B).

Ruling A1 (why two pages, not the top N): 17 rows at 3 components each are 51, over the 40. Nine rows fit (§A.6). The owner asked for the character screen, so the split follows the frame's own halves: the left column and the start of the right on page 1, then the rest of the right column and the weapons on page 2. Today's ranking ("the five slots where the best item gains most") is dropped from the panel. Its numbers stay on each row's line 2, and the slot order is what a player scans (Wowhead's table, our page). Rejected: the top 9 by gain with the other 8 as plain text lines. That would leave 8 items with no icon and no quality (tenet 2).

### A.5.2 `/bis slot:{slot}` (`bis-panel-slot`)

1. **Container**, class accent.
   1. **Header Section**: `### {Slot} · {Spec} {Class}` and `Best in slot · {band label} · {Faction}`, with the crest.
   2. **Separator**.
   3. **The pick's row** (§A.3).
   4. **Separator**.
   5. **Text Display**: `-# The next best`.
   6. **The first two alternatives**, each per §A.3. Line 2 is `{Slot} · {delta}`, where the delta is the alternative's own `signedDPS(dps_delta)` against the pick (`−7.4 DPS`, U+2212). A delta that rounds to 0.0 reads `same DPS as the pick`.
2. **Action row**: `Share in channel`, `Change picks`, `Open on Forever Sixty` (Link). There is no page button.

### A.5.3 `/char` (`char-panel`, `char-panel-nogear`)

Ephemeral, flags 32832.

1. **Container**, `accent_color` the class colour.
   1. **Header Section**, one Text Display:
      ```
      ### [{Name}]({character URL})
      Level {n} {Spec} {Class} · {Guild}
      -# Item level {x.x} · Readiness: {word} · Synced <t:{unix}:R>
      ```
      The accessory is the class crest. Line 2 is `lookupCopy.CharLine`, unchanged. In line 3, each clause appears only when known. Readiness and Synced appear only for a verified member of the guild, as `memberFields` does today. With no clause known, line 3 is absent.
   2. **Separator**.
   3. **Text Display**: `**Top upgrades** · {band label}`, then on its own line `-# By DPS gained, against the guide build.` (`-# By Forever Score gained, against the guide build.` for a tank). With no band known: `**Top upgrades**` alone.
   4. **Up to five upgrade rows** (§A.3). Order: empty slots that have a pick first (a slot left empty is the largest gap a player has), in frame order. Then the measured gains, largest first.

      Today's top three becomes five (`maxUpgrades` = 5). The budget allows nine. Five keeps the first screen to the header plus three rows on the phone (tenet 11).
   5. When more upgrades exist, a **Text Display**: `-# {n} more upgrades on the character page.` (`-# 1 more upgrade on the character page.`).
2. **Action row 1**: `Share in channel`, `Someone else`, `Character page` (Link). Unchanged (§G.4.3).
3. **Action row 2**, only with two or more verified characters: the `Another of your characters` select. Unchanged.

An empty slot as an upgrade: `bis.VerdictFor(slot, 0, false)` may give no `GainDps`. When the band names a pick for a slot the character leaves empty, the row is drawn with `fills an empty slot` on line 2 and no line 4. Today such a slot drops out of the list, because only positive measured gains are kept. This amendment keeps it. A test pins an empty Trinket 2 drawn first.

### A.5.4 Fully geared, no gear, unknown

| State | Body under the header (replaces items 3 to 5) |
|---|---|
| No gear shared (`len(card.Gear) == 0`) | Text Display: `**Top upgrades** · {band label}`, newline, `No gear shared yet. Upgrades show once the character is synced.` (`lookupCopy.NoGear`, unchanged) |
| Gear shared, no upgrade | Text Display: `**Top upgrades** · {band label}`, newline, `Every slot already wears the guide's pick for {band label}.` |
| Unknown or private name | The guided-card spec's `char.unknown` card (§G.4.3), drawn as V2 (§A.7) |

## A.6 The component and text budget (asserted in tests)

Discord counts every component, nested ones included (§0). The numbers below are counted from the mock's reference payloads (`payloads/*.json`, the `budget` field).

| Message | Count | Total | Text characters |
|---|---|---|---|
| `/bis` page 1 | Container 1 + header Section 3 + Separator 1 + 9 rows x 3 + Action row 1 + 4 buttons | **37** of 40 | 2,069 |
| `/bis` page 2 | 1 + 3 + 1 + 7 rows x 3 + 1 off-hand Text Display + 1 + 4 | **32** | 1,656 |
| `/bis slot:` | 1 + 3 + 1 + 3 + 1 + 1 + 2 x 3 + 1 + 3 | **20** | 724 |
| `/char` | 1 + 3 + 1 + 1 + 5 x 3 + 1 + (1 + 3) + (1 + 1) | **28** | 1,341 |
| `/char` no gear | 1 + 3 + 1 + 1 + 4 + 2 | **12** | 241 |
| Shared `/bis` | Container 1 + Text 1 + Gallery 1 + Action row 1 + Link 1 | **5** | 72 |
| Shared `/char` | 1 + 1 + 1 + 1 + 1 + 1 | **6** | 189 |

Rules:
- `/bis` never puts more than 9 rows on a page. This is a constant (`slotsPerPage = 9`), and a test fails if a page payload exceeds 40.
- Text guard: the builder sums every Text Display's `content`. While the sum is over 3,800, it drops line 3's source line from rows, last row first. The source stays one tap away on the page. A test builds the longest band in the build (the longest names and source lines over all 28 specs and 6 bands, at both factions) and asserts the guard never fires on real data today. A second test builds a synthetic band of 120-character names and asserts the guard holds.
- Masked-link URLs count toward the text. The longest item link on the boards is 106 characters.

## A.7 Every message in the `/bis` and `/char` flows is V2

Because the flag cannot be removed (§0), these messages are built with flag 32768 from their first frame:
- the `/bis` card (§G.4.2)
- the `/bis` answer and its pages
- the `/char` answer and its select picks
- the `/char` start and unknown cards (§G.4.3)
- every in-place failure and stale line on them

Each guided card keeps its layout and words. Only its encoding changes:

| Legacy part | V2 encoding |
|---|---|
| The embed (author, title, description, colour, thumbnail, footer) | A Container with that `accent_color`, holding one Section whose Text Display reads `-# {author}`, `### {title}` and the description, with the crest as its accessory. Before a class is known there is no crest, so it is a plain Text Display, not a Section. |
| The footer and timestamp | Dropped. V2 has neither field. The `Forever Sixty` app name already heads the message. |
| Action rows with selects and buttons | Unchanged, outside the Container. |
| A content line (`char.lead.*`, `card.failure`, `card.stale`, `This card is out of date…`) | One Text Display at the top level |

The `/bis` card counts 1 + 3 + 4 rows + 3 selects + 3 buttons = **14**.

Edits: `Message.edit` writes `flags` with 32768 and never writes `content` or `embeds`. A V2 edit body has only `components`.

`withLiveShare` (share.go) sets `Content` to the failure line today. On a V2 message it must instead put a Text Display holding `commonCopy.Failure` (`Could not load that. Try again.`) first in `components`. A test asserts no V2 payload ever carries `content` or `embeds`. Discord answers 400 when one does.

## A.8 States and interaction parity

| State | What the player sees |
|---|---|
| Default | As §A.5 |
| A link pressed (item name, `All 17 slots`, `Character page`) | Discord's own external-link confirm, then the page at the band, faction and preset (and `#gear-slot-{slot}` for `/char`). Keyboard: Discord focuses masked links and buttons in order. Touch: every link is the whole name. |
| Page button pressed | Discord's spinner on the button. Type 7 within 3 s. The panel changes in place to the other page. The Share button is unchanged. |
| A page press after 15 minutes | As the guided cards: each press carries its own fresh token (§G.6). |
| Tile fails to load (endpoint down) | Discord draws its own broken-image square. The text row still carries the name, slot, gain, quality word and source, so nothing is lost but the icon. The endpoint's error rate is logged (`op tile`). |
| Band missing (data not regenerated) | The existing failure: `Could not load that. Try again.` |
| Item unknown to the build | Question-mark tile, the name from the BiS data, no quality word or item level |
| Share pressed | Part B |

There is no hover, by Discord's design. The masked link is the "real item": it opens the item on our page, where the shared tooltip shows stats, required level and source (tenet 2's site half). Ruling A2.

# Part B. The character sheet on `Share in channel`

## B.1 The question, and the answer in the channel

The channel's question is "What is this build (or this character) wearing, at a glance?" The answer is one image laid out like the in-game character frame: every slot's icon in its quality ring with the item's name beside it, the class crest in the middle, then spec, band and faction, and the site's mark. Under it is one link to the full page.

## B.2 References, and what we borrow

- **The in-game character frame (PaperDollFrame).** Borrowed: the three slot groups, two columns of eight and a bottom row of three, in the game's order. The left column's icons sit on the outside edge, as do the right column's. Shirt and Tabard keep their sockets, drawn empty, so both columns read eight and the shape is the frame players know. The round portrait becomes the circular ringed HD crest (one crest language). The crest sits in the centre column, where the frame draws the model, with the band and faction under it.
- **Gear-inspection addons that print names beside the frame's slots** (the item-level overlays of ElvUI and similar). Borrowed: the item name in its quality colour beside each icon, the left column's name to the icon's right and the right column's to its left. A one-line quiet meta under each name.
- **Our BiS page header** (`ClassHeader`) and the OG card (`web/src/lib/og.ts`). Borrowed: the gold small-caps kicker and the Cinzel title in the class colour.

## B.3 The layout (`sheet-bis.png`, `sheet-char.png`)

Size: 540 x 808 logical, drawn at scale 2: **a 1080 x 1616 PNG**. The mock's files are 433 to 465 KB. All coordinates below are logical px. Fonts: Cinzel 700, Barlow 400 and Barlow 600, the site's own faces (`web/src/assets/fonts`).

| Region | Box (x, y, w, h) | Content |
|---|---|---|
| Surface | 0, 0, 540, 808 | `--surface-card` (gold-tint gradient over #0d0f15), 1 px border `--color-line-warm` #3a3326 |
| Kicker | 16, 16, 360, 14 | Barlow 600 11/14, uppercase, tracking 0.12em, `--color-gold` #e5b955. `/bis`: `Best in slot`. `/char`: `Level {n} {Spec} {Class} · {Guild}`. One line, ellipsis. |
| Title | 16, 32, 360, 28 | Cinzel 700 22/28, the class colour. `/bis`: `{Spec} {Class}`. `/char`: `{Name}`. One line, ellipsis. |
| Site mark | right 16, top 20 | `foreversixty.gg` Barlow 400 12/16 `--color-muted` #9a9484, then the seal (`web/public/seal-400.png`) at 28 x 28, gap 8 |
| Rule | 16, 72, 508, 1 | `--color-line-warm` |
| Left column | x 16, rows at y = 88 + 72i, i 0 to 7 | Head, Neck, Shoulder, Back, Chest, Shirt, Tabard, Wrist |
| Right column | right edge x 524, the same rows | Hands, Waist, Legs, Feet, Ring 1, Ring 2, Trinket 1, Trinket 2 |
| Centre column | x 226, w 88, y 88 to 664, contents vertically centred | Crest 64 x 64. `/bis`: `Level {n}` Barlow 600 13/16 `--color-strong` and `{Phase or range}` Barlow 400 12/15 muted. `/char`: `ITEM LEVEL` Barlow 600 10/12 tracking 0.12em muted, then `{x.x}` Cinzel 700 20/24 gold. Then the faction mark 32 x 32 (`/icons/hd/faction/{faction}.webp`, the `FactionMark` art) and `Alliance` or `Horde` Barlow 600 12/15 in `--color-alliance` #6fb1ff or `--color-horde` #ff6b5c |
| Bottom row | y 680, cells at x 16, 192, 368, each w 156 | Main hand, Off hand, Ranged |
| Footer | 16, 780, 508, 14 | Barlow 400 11/14 muted. `/bis` left: `Updated {d Mon yyyy}` (the band file's `generated_at`). `/char` left: `Gold: DPS the {band label} pick adds`, right: `Synced {d Mon yyyy}` |

**A cell** is a socket and a text block, gap 8. Cell width is 210 in the columns and 156 in the bottom row.
- Socket: 36 x 36, radius 4, black fill, 2 px border in the quality colour (`--color-rarity-*`), icon 32 x 32 inside. At scale 2 that is the icon's native 64 px, never resampled.
- Name: Barlow 600 20/23 in the quality's text colour. Rare is `--color-rarity-rare-text` #3d94f0, Epic `--color-rarity-epic-text` #b866f5, the others their `--color-rarity-*`. The lightened text variants are used because the base blue and purple fail contrast on #0d0f15. At most 2 lines.
- Meta: Barlow 400 14/18 muted, one line, ellipsis.
- Gold figure: Barlow 600 14/18 `--color-gold`. It never truncates. In the columns it follows the meta after ` · `. In the bottom row it takes its own line under the meta.
- The right column right-aligns its text and puts the socket on the outside.

**The fit rule** (the Go renderer and the mock share it):
1. Wrap the name at word boundaries within the text width.
2. If it needs more than 2 lines at 20/23, draw it at 17/20.
3. If it still needs more than 2 lines, cut whole words from the end of line 2 and end it in `…` (U+2026).

The full name is always in the alt text (§B.7) and on the page behind the link. Real cases on the boards: `Signet Ring of the Bronze Dragonflight` fits at 17 px. `Hyper Deluxe Sniper Rifle Mk XVII` in a bottom cell becomes `Hyper Deluxe Sniper Rifle…`.

**Cell states:**

| State | Socket | Text |
|---|---|---|
| `/bis` pick | Quality ring, icon | Name. Meta: the source line (`sourceLine`, the panel's own) |
| `/char` worn, no gain | Quality ring, icon | Name. Meta: `item level {n}` |
| `/char` worn, gain | Quality ring, icon | Name. Meta: `item level {n}` and the gold `+{x.x} DPS` (or `+{x.x} points` for a tank) |
| Empty slot, `/char` | Empty socket: #0d111a fill, #262e40 border, no icon | `Empty`, Barlow 400 14/18 muted. When the band has a pick for it, the gold `Upgrade available` |
| Two-handed main hand (off hand) | The main hand's icon at 35% opacity in an empty socket, as the game greys it | `Two-handed main hand`, Barlow 400 14/18 muted |
| No pick for the range, `/bis` | Empty socket | `No pick`, 14/18 muted |
| Shirt, Tabard | Empty socket | No text |
| Item unknown to the build | `inv_misc_questionmark` in a #262e40 ring | The name from the data, in `--color-text` |

Emphasis: the names in quality colour are the sheet's one primary layer. The kicker, title and gold figures are the only gold or class-coloured text. Meta lines, the footer and the mark are muted (tenet 11).

## B.4 The shared message (`share-bis`, `share-char`)

The follow-up is public. `flags` = 32768. It is multipart: `payload_json` plus `files[0]` named `sheet.png`, and `attachments: [{ "id": 0, "filename": "sheet.png", "description": "{alt}" }]`.

1. **Container**, class accent.
   1. **Text Display**. `/bis`: `### {Spec} {Class}` and `Best in slot · {band label} · {Faction}`. `/char`: `### [{Name}]({character URL})` and `Level {n} {Spec} {Class} · {Guild}`. A visitor's view, so no readiness.
   2. **Media Gallery**, one item: `media.url` `attachment://sheet.png`, `description` the alt text.
   3. `/char` only, a **Text Display**: `-# Gold figures are the DPS the {band label} pick adds in that slot.` (`-# Gold figures are the Forever Score points the {band label} pick adds in that slot.` for a tank.)
2. **Action row**: `All 17 slots` (Link to the band page) for `/bis`, `Character page` (Link) for `/char`. One link, nothing else: the public message never carries Share, the page button or Change picks (§G.5.3).

The `/bis` sheet always shows all 17 slots, whichever page the sharer was on. A shared `/bis slot:{slot}` answer has no sheet: it posts the slot panel (§A.5.2) publicly, with its `Open on Forever Sixty` link. One item does not make a character frame.

## B.5 Flow, timing, failure

1. The press is answered as today: type 7, the ephemeral's button turns into the disabled `Shared` within 3 s.
2. In `rep.After`, the bot builds the sheet model from the same recomputed visitor answer (`visitorAnswer`). It renders the PNG (§C) and sends the multipart follow-up. Render budget: p95 at most 250 ms on the API's Cloud Run instance, measured and recorded in the lane report.
3. **The render fails** (an asset missing, a draw error): the bot logs `op sheet_render` with the error, and posts the visitor's panel publicly with no image (`share-bis-nosheet`). That is the Container of §A.5.1 or §A.5.3 and one Link button, never Share, the page button or Change picks. Nothing fails silently, and the channel still gets the answer.
4. **The follow-up fails**: unchanged (`postShare`). The token is freed and the live button is restored with the failure line, as a V2 Text Display (§A.7).
5. The upload is about 450 KB, far under Discord's per-file limit for a bot. A test asserts that a sheet with the longest names in the build stays under 1.5 MB.

## B.6 Phone legibility

- The phone client draws the gallery image at the container's inner width. That is about 270 css px on the 390 client (the mock's measure, `share-*-phone`). Capture D20 records the true width.
- At 270 css px the 540-wide sheet is drawn at 0.5. Names at 20 px render at about 10 css px, and the 17 px fit size at 8.5. Sockets render at 18 css px, the crest at 32. The 1080 px file gives a retina phone (3x, about 810 device px) a sharp image.
- What reads inline on the phone: every icon, every quality ring, the crest, the faction, the title, and the names at 10 css px.
- What needs a tap: the meta lines (7 css px). Discord's viewer opens the full 1080 px image for pinch zoom. The meta repeats what the ephemeral panel and the page already say, which is why it is the secondary layer.
- The layout is portrait (0.67) on purpose. A landscape frame drawn at 270 css px would put the names under 7 px.
- If capture D20 shows the desktop client capping the gallery's height so the sheet draws under 440 css px wide, the sheet stays as specified. Desktop opens the full image on click, and the phone is the binding case. Record it as a known limit in the lane report.

## B.7 Accessibility

- Media Gallery item `description` and the attachment's `description` are the same alt text, at most 1024 characters. `/bis`: `Best in slot sheet, {Spec} {Class}, {band label}, {Faction}. {Slot}: {Item}; …` over all 17 slots in frame order. The off hand reads `Off hand: none, two-handed main hand`. `/char`: `Character sheet, {Name}, level {n} {Spec} {Class}, {Guild}. {Slot}: {Item}, {+x.x DPS} available; {Slot}: empty; …`.
- If the alt text is over 1024, it drops the `, {gain} available` clauses first, then cuts at the last whole slot and ends in `…`. The mock's alt texts are 624 (`/bis`) and 742 (`/char`) characters.
- Thumbnail `description`: the tile per §A.4, the crest `{Class} crest`.
- Colour is never the only carrier. Quality is also the quality word in the panel, gains are figures, and empty is a word.
- Contrast on #0d0f15: names use the lightened rare and epic text variants. The muted text #9a9484 is above 4.5:1. Gold #e5b955 is above 7:1.

# Part C. The render path

**Recommendation: one Go renderer inside the API, used for both the sheet and the tiles. The sheet is attached to the follow-up message; the tiles are served from a cached endpoint.**

| Option | Verdict | Why |
|---|---|---|
| **Go renderer in the API** (`api/internal/sheet`): `golang.org/x/image/font/opentype` for text, `golang.org/x/image/webp` to decode icons, crest and mark, `image/draw` and `image/png` | **Chosen** | One drawing code serves `/bis` and a player's actual gear (`/char`), which only exists at request time. It reads the same data the bot already reads (`BisDataDir`, `loot.json`, `BuildItem`), so the sheet can never disagree with the panel beside it (tenet 8). There is no second runtime on Cloud Run. The tile is the sheet's socket, so tenet 10 holds in code too. |
| Build-time static sheets per spec x band x faction x preset (28 specs x 6 bands x 2 factions = 336 PNGs, served from Cloudflare Pages) | Rejected | It cannot draw `/char`, so a second renderer would exist anyway. It adds about 150 MB to every deploy. Nightly data changes would need a site rebuild before Discord shows the change. |
| Node `satori` + `resvg` (what `web/src/lib/og.ts` does at build time) as a sidecar or a Pages Function | Rejected | It is a second runtime and service to deploy, and a network hop on every share. Its only advantage is reusing `og.ts`, which draws one title. |
| A public image URL in the gallery instead of an attachment | Rejected for the sheet | The attachment is a snapshot that cannot drift from the message's words, survives our outages, and puts no player's gear behind a guessable public URL. It is used for the tiles, because those are public item art. |

Assets the renderer embeds or reads:
- **Fonts.** Cinzel 700 and Barlow 400 and 600 TTF, OFL, `go:embed`. Fetched by `make sheet-fonts` from the same `@fontsource` packages the site uses. The site's woff2 subsets are cut to the site's glyphs, so the renderer uses full Latin TTFs.
- **Crests** (9 x 128 px), the **faction marks** (`alliance.webp`, `horde.webp`, 72 px) and the **seal**. Copied by `make sheet-assets` from `web/public` into `api/internal/sheet/assets/` and embedded. They are copied, never redrawn. A test compares their hashes to the `web/public` files.
- **Item icons.** Read at runtime from `data/builds/{active build}/icons/{icon}.webp` (1,704 files, 6.7 MB). `api/Dockerfile` adds `COPY data/builds/${ACTIVE_BUILD}/icons/ …` beside its `simdb.bin` line. A missing icon is the question-mark icon, logged once per icon (`op sheet_icon_missing`).
- **No resampling of game art up.** Icons draw at 64 px, crests at 128 px (64 logical at scale 2), the seal down from 400. The faction mark draws at 64 px from its 72 px source, a downscale (Catmull-Rom). This is the faithful-icons rule. It is also why the crest sits at 64 logical, not larger: a 256 px crest source would let it grow (owner question 3).

Caching:
- **Sheets.** An in-process LRU keyed by SHA-256 of the sheet model's canonical JSON: build, spec, band, preset, faction, and per slot the item id, name, meta and gold figure, plus header strings, crest, faction and date. 32 entries, capped at 48 MB. A share of the same `/bis` sheet by ten players renders once. `/char` sheets change whenever the export changes, so they key apart.
- **Tiles.** `GET /v1/discord/tiles/{build}/{item}.png`. `{build}` must be a known build folder and `{item}` digits, or the answer is 404. The answer carries `Cache-Control: public, max-age=31536000, immutable`, keyed by build, because an item's icon and quality are fixed within a build. An in-process LRU holds 2,048 tiles. Discord's media proxy caches in front, so the API sees about one fetch per tile per build. The tile is about 6 KB. The endpoint has rate limits like every public route (bot spec §8).

# Copy, verbatim

Strings live in `lookupCopy` (`api/internal/discordbot/copy_lookup.go`) and `guideCopy`, keyed as below. Tests read each string back from this table. `{…}` is escaped data. No string uses `soon`, an exclamation mark or `unlock`. Retired strings are listed at the end.

| Key | String |
|---|---|
| `bis.head.title` | `### {Spec} {Class}` |
| `bis.head.line` | `Best in slot · {band label} · {Faction}` |
| `bis.head.page` | `-# {Page range}, {n} of 17 slots. Each gain is over the next best item.` |
| `bis.page.one` | `Head to Legs` (page 1's range, and page 2's button) |
| `bis.page.two` | `Feet to Ranged` (page 2's range, and page 1's button) |
| `bis.slot.title` | `### {Slot} · {Spec} {Class}` |
| `bis.slot.next` | `-# The next best` |
| `bis.tie` | `ties the next best` |
| `bis.slot.tie` | `same DPS as the pick` |
| `bis.offhand.twoHand` | `**Off hand** · none: {main hand name} is two-handed.` |
| `bis.slot.none` | `**{Slot}** · no pick for this range.` |
| `item.facts` | `-# {Quality} · item level {n} · {source line}` (each part only when known) |
| `item.replaces` | `-# Replaces {worn name}, item level {n}`; without a level `-# Replaces {worn name}` |
| `char.upgrades.title` | `**Top upgrades** · {band label}` |
| `char.upgrades.byDPS` | `-# By DPS gained, against the guide build.` |
| `char.upgrades.byScore` | `-# By Forever Score gained, against the guide build.` |
| `char.upgrades.empty` | `fills an empty slot` |
| `char.upgrades.more` | `-# {n} more upgrades on the character page.`; for one `-# 1 more upgrade on the character page.` |
| `char.upgrades.none` | `Every slot already wears the guide's pick for {band label}.` |
| `char.noGear` | `No gear shared yet. Upgrades show once the character is synced.` (unchanged) |
| `char.facts` | `-# Item level {x.x} · Readiness: {word} · Synced <t:{unix}:R>` (each clause only when known) |
| `share.char.gold` | `-# Gold figures are the DPS the {band label} pick adds in that slot.`; tank `-# Gold figures are the Forever Score points the {band label} pick adds in that slot.` |
| `tile.alt` | `{Item name}, {Quality} {slot} item`; without a quality `{Item name}, {slot} item` |
| `crest.alt` | `{Class} crest` |
| `sheet.bis.kicker` | `Best in slot` |
| `sheet.bis.level` | `Level {n}` |
| `sheet.bis.range` | `Phase 1 raid` for `p1`; `No buffs or consumables` for the level 60 bare band; the band's words for 20 to 50 (`Levels 20 to 29`) |
| `sheet.char.ilvl` | `Item level` (drawn uppercase by the renderer) |
| `sheet.empty` | `Empty` |
| `sheet.upgrade` | `Upgrade available` |
| `sheet.twoHand` | `Two-handed main hand` |
| `sheet.none` | `No pick` |
| `sheet.mark` | `foreversixty.gg` |
| `sheet.updated` | `Updated {d Mon yyyy}` |
| `sheet.char.gold` | `Gold: DPS the {band label} pick adds`; tank `Gold: points the {band label} pick adds` |
| `sheet.char.synced` | `Synced {d Mon yyyy}` |
| `sheet.alt.bis` | `Best in slot sheet, {Spec} {Class}, {band label}, {Faction}. {Slot}: {Item}; …` |
| `sheet.alt.char` | `Character sheet, {Name}, level {n} {Spec} {Class}, {Guild}. {Slot}: {Item}, {gain} available; …` |
| `sheet.alt.offhand` | `Off hand: none, two-handed main hand` |

Unchanged and reused: `Share in channel`, `Shared`, `Change picks`, `All 17 slots`, `Open on Forever Sixty`, `Character page`, `Someone else`, `Another of your characters`, `Could not load that. Try again.`

Retired:
- `lookupCopy.BisTopLead` (`{Spec} {Class}, {band label}. The five slots where the best item gains the most over the next best:`)
- `lookupCopy.BisSlotLead`
- `lookupCopy.BisFooter`
- `lookupCopy.BisAuthor` on items
- `lookupCopy.ItemLineOne` and `ItemLineTwo`, for `/bis` and `/char` (`/report` keeps them)
- `lookupCopy.TopUpgrades`, `UpgradesBand`, `UpgradesByDPS` and `UpgradesByScore`, replaced by the `char.upgrades.*` keys
- `topSlots = 5` and `maxUpgrades = 3`, replaced by `slotsPerPage = 9` and `maxUpgrades = 5`

# Rulings

A1. Two pages in the frame's order instead of a top N. See §A.5.1.

A2. No hover exists in Discord. The masked link on every item name opens the item on our page, where the shared tooltip is the "real item" (tenet 2). The owner asked whether Discord allows hover: it does not, and this is the closest Discord surface.

A3. The quality colour lives in a drawn tile, because V2 has one accent colour per container (§A.4).

A4. Slot and gain at body size, and the item facts as subtext. On the phone the old two-line grammar wrapped to four lines per item. Subtext puts the reference facts one level down (tenet 14) and keeps a row about 3 lines.

A5. `/char` shows five upgrades, not three. An empty slot sorts first and says so.

A6. Every `/bis` and `/char` message is V2 from its first frame, forced by the one-way flag (§0).

B1. The sheet goes only on the public share. The ephemeral answer stays the panel, because a 1616 px image above it would push every item below the first screen.

B2. Shirt and Tabard are drawn as empty sockets, keeping the frame's shape. They carry no text, and the alt text leaves them out.

B3. The sheet is portrait, 540 x 808 logical, because the phone is the binding case (§B.6).

B4. The crest is drawn at its native 128 px (64 logical). No game or crest art is ever upscaled.

B5. A sheet that fails to render degrades to the panel, public, and is logged. A share never fails because of the image.

C1. One Go renderer in the API, attachment for the sheet, cached endpoint for the tiles (Part C).

# Acceptance

## Handler tests (`api/internal/discordbot`, strings read back from the copy table)

1. **V2 everywhere.** Every `/bis` and `/char` frame (cards, answers, pages, picks, unknown, failure, stale, the live-share restore) carries flag 32768. It never carries `content` or `embeds`. Ephemeral frames carry 32832.
2. **`/bis` page 1 and page 2.** The header lines. Nine rows in frame order, then eight. The off-hand two-hander line. `ties the next best` for a delta under 0.05. The tile URL and alt text. The four buttons in order. The page id under 100 characters. A page press answers type 7 with the other page and the same Share custom id. No page payload over 40 components or 3,800 text characters.
3. **`/bis slot:`.** The pick, `-# The next best`, two alternatives with U+2212 deltas, `same DPS as the pick`, three buttons, no page button.
4. **`/char`.** The header with each clause present and absent. Member clauses only for a verified member. Up to five upgrades, an empty slot first. The `Replaces` line with and without a level. The `more` line singular and plural. The no-gear and no-upgrade bodies. A tank in points. Both rows unchanged.
5. **Escaping.** An item name with `[`, `]`, `(`, `)` or `*` stays inside its masked link. A guild name with markdown is escaped.
6. **Text guard.** The longest real band in the build (all specs, bands and factions) never triggers the guard. A synthetic band with 120-character names drops source lines last row first and ends at 3,800 or below.
7. **Share.** The follow-up is multipart, with `attachments[0].filename` `sheet.png` and the alt description. The gallery's `media.url` is `attachment://sheet.png`. A `/bis` share from page 2 renders all 17. A `/bis slot:` share posts the slot panel with no file. A render error posts the no-sheet panel and logs `sheet_render`. A follow-up error restores the live button with the failure line as a Text Display.
8. **Alt texts.** Both alt forms at or under 1024 characters on the longest real band. The over-length rule (gains dropped first) on a synthetic case.

## Renderer tests (`api/internal/sheet`)

9. **Golden PNGs.** The `/bis` and `/char` sheets of the mock's exact data (Beast Mastery Hunter, level 60 Phase 1 raid, Alliance, build 1.60.1.70658, Zulmara's gear per `data.mjs`) render at 1080 x 1616. Each is compared to a committed golden with a per-pixel tolerance for font rasterising, and the golden is set beside `sheet-*.png` in the lane report.
10. **Fit rule.** `Signet Ring of the Bronze Dragonflight` in a column fits at 17 px in two lines. `Hyper Deluxe Sniper Rifle Mk XVII` in a bottom cell ends `Rifle…`. A gold figure never truncates.
11. **Every cell state of §B.3**, one each: pick, worn, worn with gain, empty, empty with upgrade, two-handed, no pick, shirt and tabard, unknown item.
12. **Icons 1:1.** The decoded icon's pixels appear unscaled in the socket. Asset hashes match `web/public`.
13. **Tiles endpoint.** 72 x 72, ring colour per quality, unknown item gives the question-mark tile, bad build gives 404, the cache headers are present.
14. **Budget.** Sheet render p95 at or under 250 ms over 50 renders. The PNG is under 1.5 MB with the build's longest names.

## Boards the reviewers compare

The mock boards above, beside the same frames rendered from the real handlers' payloads through the reviewers' kit. The kit gains a Components V2 drawing, ported from this mock's `component()` and `markdown()`: Container, Section with accessory, Text Display with `###` and `-#`, Separator, Media Gallery. The two sheets go beside their goldens. A difference in words, order, colour, quality ring, button style or any cell state is a defect against this spec. Discord's own chrome is not.

## Discord captures (test server, production application), added to the bot spec's §9 table

| # | What | Desktop client | Phone client |
|---|---|---|---|
| D20 | `/bis` page 1, page 2 and the slot form. `/char` own and no gear. Measure the drawn thumbnail size and the gallery image's width and height on both clients | x | x |
| D21 | A test-only message of 4,001 text characters across Text Displays: record Discord's answer (decides whether the 3,800 guard stands) | x | |
| D22 | `Share in channel` on `/bis` (from page 2) and `/char`: the public message with the sheet, the image opened full size, pinch zoom on the phone; a forced render failure posting the no-sheet panel | x | x |
| D23 | `Change picks` on a V2 answer reopening the V2 card in place; a page press 20 minutes after the answer | x | x |
| D24 | A screen reader (VoiceOver on iOS, NVDA on desktop) reading a shared sheet's alt text and a tile's alt text | x | x |

# Questions only the owner can answer

1. **Shirt and Tabard sockets.** The sheet draws them empty to keep the in-game frame's shape (ruling B2). Keep them, or close the left column up to six slots?
2. **The sheet on the ephemeral answer too.** Ruling B1 keeps it to the public share, so the panel stays on the first screen. Do you also want a `Show sheet` button on the ephemeral answer? It would be one more render per press.
3. **A 256 px crest.** The crests ship at 128 px, so the sheet's centre crest is 64 logical (128 device px), the largest it can be without upscaling. If you want the crest to lead the sheet at 96 or 128 logical, it needs 256 px crest art made the same way as the 512 px faction marks.
4. **`/bis` order.** The panel lists slots in the frame's order on two pages (ruling A1). Today's answer led with "the five slots where the best item gains the most". Is the frame order right for a BiS answer, or should page 1 lead with the largest gains?
