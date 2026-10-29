# Discord bot amendment: the gear panel and the character sheet: experience spec

Author: ux-designer, 2026-10-29. Revision 2 the same day applies the coordinator's rulings on the persona gates (ux and wow-player, both FIX, direction right); they are listed under "Revision 2" below. Amends `design/specs/2026-10-27-discord-bot.md` (the bot spec) §3.3.3 and §3.3.4, and `design/specs/2026-10-28-discord-guided-cards.md` (the guided-card spec) §G.4.2, §G.4.3, §G.5 and §G.6. Tenets: `docs/tenets.md` 1, 2, 4, 7, 8, 10, 11, 13, 14.

Source (the owner, 2026-10-29, on a production `/bis` answer of five stacked item embeds, verbatim): "can we return a generated output that looks like the character screen with hover over or the items names beside them? Does discord rendering allow that?" The coordinator's answer, which the owner approved with "do it": Discord has no hover, so do both. Part A replaces the `/bis` and `/char` item answers with one compact Components V2 panel. Part B posts a server-rendered character-sheet PNG on `Share in channel`.

## Revision 2 (2026-10-29), what changed

1. **Blocker fixed.** §A.7 now matches §0: a type 7 update whose source message lacks flag 32768 sends `flags` with 32768, `content: null` and `embeds: []`. An update to a message that already has the flag never sends those keys. Test 1 and capture D23 cover a legacy card pressed after deploy. Test 2 asserts at most 38 components, the live-share failure line included.
2. **Sheet.**
   - Meta lines are 16/20.
   - Names sit on one line, shrinking from 20 px to 14 px, at a fixed row height and an even 72 px pitch.
   - Nothing is cut with "…". A source that cannot fit falls back to a short form built only from real abbreviations, then to no meta line.
   - Shirt and Tabard are gone; the left column's six slots centre on the right column's eight.
   - The sheet is now 600 x 844 logical (1200 x 1688 PNG).
3. **Panel.**
   - Each name is led by a quality-coloured dot, from six new application emoji (§A.4.1).
   - A gain shows only when there is one.
   - An unmodelled effect reads `Gain not modelled yet`.
   - Quest sources read `requires level {min_level}`.
   - Slot labels match between panel and sheet.
4. **Two rulings checked against the data and deliberately not applied as worded.**
   - "Masons Fraternity Ring" is the item's real name: item 9533 in `items.json` and in the class table, with no apostrophe. It stays.
   - The ruled legend "Gain over the next best in the level 60 Phase 1 raid" would misstate the `/char` sheet's gold figures. They are the pick's gain over what is worn, measured by the band, not over the next best. The legend says that instead (ruling B6).
5. **Order.** Character-screen order is kept (owner question 4 is closed by the coordinator).

## Mock

`design/mocks/2026-10-29-discord-gear-panel/gen_gear_panel.mjs` (with `data.mjs`, `payloads.mjs`, `sheet.mjs`). Run it with Node 22 or later.

The reviewers' render kit (`web/tests/captures/discord-bot/discord-render.ts`) draws legacy embeds only, so the mock carries its own Components V2 drawing in the kit's dark-theme colours.

Every item, icon, quality, item level, source line, unmodelled flag and DPS figure is real. The data is the active build `1.60.1.70658`, Beast Mastery Hunter, level 60 Phase 1 raid, Alliance, read from `data/builds/1.60.1.70658/bis/hunter-beast-mastery.json`, `items/hunter.json` and `loot.json`. The source lines go through a port of the bot's own `describeSource`, with this spec's quest change. The example character Zulmara wears the band's own alternatives, so each gain is the band's measured delta.

The reference payloads are in `payloads/*.json`, each with its counted `budget`. The generator fails if any payload exceeds 38 components or 3,800 text characters. Boards are drawn at 640 (desktop client) and 390 (phone client).

| File | What it shows |
|---|---|
| `bis-panel-p1-{desktop,phone}.png` | `/bis`, Beast Mastery Hunter, Phase 1 raid: Head to Legs, 9 slots |
| `bis-panel-p2-{desktop,phone}.png` | After `Feet to Ranged`: 8 slots, the off hand empty because the main hand is two-handed |
| `bis-panel-p1-sharefail-{desktop,phone}.png` | Page 1 after a failed share: the failure line on top, Share live again (38 components) |
| `bis-panel-slot-{desktop,phone}.png` | `/bis slot:legs`: the pick and the next two |
| `char-panel-{desktop,phone}.png` | `/char` for a linked member: header, five upgrades (an empty slot first), `4 more upgrades` |
| `char-panel-nogear-{desktop,phone}.png` | `/char` with no gear shared |
| `share-bis-{desktop,phone}.png` | The public message after `Share in channel` on `/bis`, with the sheet attached |
| `share-char-{desktop,phone}.png` | The same for `/char` |
| `share-bis-nosheet-{desktop,phone}.png` | The fallback when the sheet cannot render: the panel, public, without the sheet |
| `sheet-bis.png` | The `/bis` sheet at its real size, 1200 x 1688 (492 KB) |
| `sheet-char.png` | The `/char` sheet at its real size, 1200 x 1688 (453 KB) |

No new site token and no new site control (tenet 10). The Discord side uses only Discord's own components and application emoji. The sheet is drawn only from the site's tokens, fonts, quality colours, HD crest and faction mark.

# 0. Discord facts, verified

Source: `discord/discord-api-docs`, branch `main`, fetched 2026-10-29. The last commit to `developers/components/reference.mdx` is dated 2026-10-28T22:31Z. The files read were `developers/components/reference.mdx`, `overview.mdx`, `using-message-components.mdx`, `developers/resources/message.mdx` and `developers/change-log.mdx`. These files are the source of https://discord.com/developers/docs/components/reference.

| Fact | Value | Source |
|---|---|---|
| The flag | `IS_COMPONENTS_V2 = 1 << 15` (32768), per message | message.mdx flags table; reference.mdx intro |
| Content and embeds | Forbidden with the flag: "Providing `content`, `embeds`, `sticker_ids`, `poll`, or `shared_client_theme` will fail with a 400 BAD REQUEST" | message.mdx, Create Message note ** |
| The flag is one-way | "Once a message has been sent with this flag, it can't be removed from that message." An edit can set it, never unset it. **Setting it on an edit requires `content: null` and `embeds: []`**, or Discord answers 400 | reference.mdx intro; message.mdx, Edit Message note * |
| Component limit | "Messages allow up to 40 total components." Nested components count. The top-level and Container-children limits were removed on 2025-04-29 | reference.mdx; change-log.mdx 2025-04-29 |
| Container (17) | `components`, `accent_color` 0x000000 to 0xFFFFFF, `spoiler`. Children may be Action Row, Text Display, Section, Media Gallery, Separator or File | reference.mdx, Container |
| Section (9) | 1 to 3 Text Displays, plus `accessory`: one Button or one Thumbnail. The docs' example draws the accessory to the right of the text | reference.mdx, Section, and `section.webp` |
| Text Display (10) | `content`, markdown "similar to a message", custom emoji included. No per-component length is stated | reference.mdx, Text Display |
| Thumbnail (11) | `media` (`https` URL or `attachment://<filename>`), `description` (alt, max 1024), `spoiler`. Images only. **No displayed size is stated** | reference.mdx, Thumbnail |
| Media Gallery (12) | 1 to 10 items, each with `media`, `description` (alt, max 1024) and `spoiler` | reference.mdx, Media Gallery |
| Separator (14) | `divider` (default true), `spacing` 1 or 2 | reference.mdx, Separator |
| Attachments | Shown only when a component references them. Upload is multipart. Attachment `description` max 1024 | reference.mdx; message.mdx |
| Button | `label` max 80 (the design note: 38 without an icon), `custom_id` 1 to 100, `url` max 512 | reference.mdx, Button |
| Where the flag works | Message Create, Execute Webhook, interaction responses and follow-ups | using-message-components.mdx |

Not in the docs, so not guessed:
1. **Total text length.** None of the five files states a limit on text across Text Displays. Community libraries validate 4,000. The builder guards at 3,800 (§A.6), and capture D21 records what Discord answers to 4,001.
2. **Displayed sizes of thumbnails and gallery images.** The mock draws a thumbnail at 72 px (desktop) and 64 px (phone), and the gallery at the container's full inner width. Capture D20 measures both. No rule in this spec depends on the exact figure.
3. **Text colour.** Discord markdown has no colour syntax. Only ANSI code blocks colour text, and those cannot hold links; the coordinator rules them out. §A.4.1 solves colour with application emoji.

# Part A. The compact panel (`/bis`, `/char`)

## A.1 The one question, and the answer on the first screen

| Viewer, surface | Question | Visible with no scroll, phone client (390) and desktop client (640) |
|---|---|---|
| Anyone, `/bis` with a spec | "What is the best item in each slot, and where do I get it?" | Phone: the header (spec, class, band, faction, crest) and the first three to four slots, each with its dot, bold name, slot, gain and source. Desktop: all nine slots of page 1. |
| Anyone, `/bis slot:{slot}` | "What is best in this slot, and what comes next?" | The pick and both alternatives |
| A player, `/char` | "What should this character replace first?" | The character line, item level and readiness, and the first three upgrades, each saying what it replaces |

Tenet 14: the panel leads with the class and the band, or the character. The first item row is the next step.

## A.2 References, and what we borrow

- **The in-game character frame.** Borrowed: the slot order. Left column Head, Neck, Shoulder, Back, Chest, Wrist. Right column Hands, Waist, Legs, Feet, Ring 1, Ring 2, Trinket 1, Trinket 2. Bottom Main hand, Off hand, Ranged. The two pages follow it.
- **Wowhead's BiS table and our BiS page's "The list".** Borrowed: one row per slot with the slot, the icon, the name and the page's own source words. Not borrowed: runners-up, which stay on the page (`All 17 slots`).
- **The game's item tooltip.** Borrowed: the name led by its quality colour, with quality and item level as a quiet line under it.
- **Our `/char` card today.** Borrowed unchanged: the character line, item level, readiness and synced. Added: what each upgrade replaces.
- **Discord's Section pattern.** Borrowed: one Section per item, with the icon tile as the Thumbnail accessory, drawn on the right by Discord.

## A.3 The item row (one grammar, both commands)

Each item row is one Section (type 9) holding one Text Display and a Thumbnail accessory: 3 components.

```
{quality dot} **{Item name}**
{Slot}[ · {gain}]
-# {Quality} · item level {n} · {source line}
-# Replaces {worn name}, item level {n}          (/char only; absent for an empty slot)
```

- **Line 1, the row's one primary element.** It is the quality dot (§A.4.1), a space, then the name in plain bold text, **never a link** (ruling A10: Discord draws every link in its link blue #00a8fc, which hid the quality colour). The colour is carried by the dot beside it and the ring on the tile. When the dot emoji is not loaded, line 1 is the name alone. The panel's link button (`All 17 slots`, `Open on Forever Sixty`, `Character page`) opens the page.
- **Line 2, body size.** The slot (the `/bis` slot choice's name: `Ring 1`, `Ring 2`, `Trinket 1`, `Trinket 2`, the same words the sheet's alt text uses), then ` · ` and the gain **only when there is one**:
  - `/bis`: `signedDPS(-alternatives[0].dps_delta)` when its absolute value is at least 0.05. A tie or a slot with no alternative shows the slot alone. The retired `ties the next best` never appears.
  - `/char`: the band's measured gain over the worn item. An empty slot shows its real gain, the pick's stat-weight score in DPS (`score x reference_dps_per_point`, the site's `upgrades.ts` rule with nothing worn), or `Gain not modelled yet` where no figure exists: a sim-decided pick with no score (an on-use trinket), a weapon, or an unmodelled effect (ruling A11).
  - Tank: `+{x.x} points`.
  - **A pick with `effect_unmodelled: true`** (its proc, on-use or set effect is not in the sim) shows `Gain not modelled yet` in place of any figure, because the figure would leave out the effect that makes it the pick. In this band that is Head (Beastmaster's Cap), Feet (Beastmaster's Treads) and Ring 1 (Signet Ring of the Bronze Dragonflight). In the slot form, an alternative carrying the flag reads the same.
- **Line 3, Discord subtext** (`-# `, 13 px, #949ba4): quality word, item level, and the source line from `sourceLine`, in full, never cut. Each part is left out when unknown.
- **Quest source lines change** (here and on the site, together; `item_source.go`'s header says to change both together). The new line is `Quest: {name} · requires level {min_level}`.
  - `{min_level}` is the level a player must be to take the quest. It is loot.json's quest `min_level`, from classic-db.
  - It is **not** `level`, the quest's own level, which the line showed as `Level 60`. Players read a bare "Level 60" as the requirement, and it is not: Saving the Best for Last is quest level 60, requires 58.
  - The data holds both fields. `lootQuest` gains `MinLevel int \`json:"min_level"\``.
  - A quest with no `min_level` shows `Quest: {name}`.
  - The site's `questSourceLabel` changes in the same lane, so the panel and the page read alike.
- **Line 4 (`/char`), subtext:** `Replaces {worn name}, item level {n}`. Without a known level: `Replaces {worn name}`. Without a name: no line.
- **Escaping.** Names go through `memberText`, so `*`, `_`, `[` and `]` in a name stay inside the bold.
- **No item links.** The rows carry no link; the message's link button opens the page.

The `ItemEmbed` grammar is retired for `/bis` and `/char`. `/report` keeps it until its own amendment.

## A.4 The tile: the icon with its quality ring

The Thumbnail URL is `{API_BASE_URL}/v1/discord/tiles/{build}/{item id}.png`.
- A 72 x 72 PNG: a 4 px ring in the game's quality colour (`--color-rarity-*`: #9d9d9d, #ffffff, #1eff00, #0070dd, #a335ee, #ff8000) around the icon at 64 x 64, 1:1, never resampled.
- An item unknown to the build gets `inv_misc_questionmark` in a #262e40 ring.
- Alt text: `{Item name}, {Quality} {slot} item`. The header's accessory is the class's ringed HD crest, alt `{Class} crest`.

### A.4.1 The quality dot (ruling A7)

**Choice:** six application emoji, one per quality, at the start of line 1.
- Names: `fs_q_poor`, `fs_q_common`, `fs_q_uncommon`, `fs_q_rare`, `fs_q_epic`, `fs_q_legendary`.
- Markup: `<:fs_q_epic:{id}> **[Beastmaster's Cap](…)**`.

**Why the dot:**
- Discord markdown cannot colour text in a Text Display. The only coloured text form, an ANSI code block, cannot carry a link and draws as a monospace box; the coordinator rejected it.
- The game teaches quality by colour before the name is read, in bags, loot rolls and tooltips. A dot gives the colour at the same glance, at the line's start, where the eye enters.
- The words alone (`Epic Beastmaster's Cap`) would add a word to every name, and colour is the faster read.
- The dot costs no component, only text: 37 characters per row, which §A.6 counts.
- It is the same mechanism as the roster's spec emoji (guided-card spec §P.9.1), so no new setup path is needed.

**Images:**
- Six transparent PNGs at Discord's 128 px emoji size: the 64 px drawing at scale 2, one filled circle of diameter 72 centred, in the quality colour (`--color-rarity-*`), with a 4 px #000000 outline at 45% opacity centred on its edge, so Common's white dot holds on Discord's light theme.
- Drawn in Go by `sheet.QualityDots` (`api/internal/sheet/dot.go`) when `api discord-emoji` runs; nothing is committed. `sheet.QualityEmojiName(q)` names them.
- Discord draws an inline emoji at about 22 px, so the dot reads at about 12 px.

**Setup:** `api discord-emoji` uploads them with the 28 spec emoji: 34 of the application's 2,000. Missing emoji follow §P.9.1's rule: the bot re-reads at most hourly, logs `emoji_missing`, and draws line 1 without the dot. The quality word on line 3 always carries the quality, so colour is never the only carrier.

## A.5 Layouts

### A.5.1 `/bis` with a spec (`bis-panel-p1`, `bis-panel-p2`)

Ephemeral. `flags` = 64 | 32768 = 32832.

1. **Container**, `accent_color` the class colour (`classColors`).
   1. **Header Section.** One Text Display:
      ```
      ### {Spec} {Class}
      Best in slot · {band label} · {Faction}
      -# {Page range}, {n} of 17 slots. Each gain is over the next best item.
      ```
      The accessory is the class crest.
   2. **Separator**, divider, spacing 1.
   3. **The page's slot rows** (§A.3).
      - Page 1 (`Head to Legs`): Head, Neck, Shoulder, Back, Chest, Wrist, Hands, Waist, Legs (9).
      - Page 2 (`Feet to Ranged`): Feet, Ring 1, Ring 2, Trinket 1, Trinket 2, Main hand, Off hand, Ranged (8).
   4. A slot with nothing to show is one Text Display:
      - Off hand with `empty_reason: two_hand_equipped`: `**Off hand** · none: {main hand name} is two-handed.`
      - No pick: `**{Slot}** · no pick for this range.`
2. **Action row** (outside the Container): `Share in channel` (Secondary), the page button (Secondary: `Feet to Ranged` on page 1, `Head to Legs` on page 2), `Change picks` (Secondary), `All 17 slots` (Link).

The page button's id is `gbis:1:pg:{class}:{spec}:{phase}:{slot}:{page}`; the longest is 32 characters. A press answers type 7 with the other page. It copies the message's `Share in channel` button from the interaction's `message.components` unchanged and mints no token.

### A.5.2 `/bis slot:{slot}` (`bis-panel-slot`)

1. **Container**, class accent.
   1. **Header Section**: `### {Slot} · {Spec} {Class}` and `Best in slot · {band label} · {Faction}`, with the crest.
   2. **Separator**.
   3. **The pick's row.**
   4. **Separator**.
   5. **Text Display**: `-# The next best`.
   6. **The first two alternatives.** Line 2 is `{Slot}` plus ` · {delta}` only when the absolute delta is at least 0.05 (`−7.4 DPS`, U+2212). An unmodelled pick or alternative shows `Gain not modelled yet` instead.
2. **Action row**: `Share in channel`, `Change picks`, `Open on Forever Sixty` (Link).

### A.5.3 `/char` (`char-panel`, `char-panel-nogear`)

Ephemeral, flags 32832.

1. **Container**, class accent.
   1. **Header Section**:
      ```
      ### [{Name}]({character URL})
      Level {n} {Spec} {Class} · {Guild}
      -# Item level {x.x} · Readiness: {word} · Synced <t:{unix}:R>
      ```
      The accessory is the crest. Each line-3 clause appears only when known. Readiness and Synced appear only for a verified member.
   2. **Separator**.
   3. **Text Display**: `**Top upgrades** · {band label}`, then `-# By DPS gained, against the guide build.` (tank: `-# By Forever Score gained, against the guide build.`).
   4. **Up to five upgrade rows** (§A.3; `maxUpgrades` = 5), ordered:
      1. Empty slots that have a pick, in frame order, each with its real gain or `Gain not modelled yet` (§A.3).
      2. Measured gains, largest first.
      3. Picks whose gain is not modelled (`Gain not modelled yet`), in frame order.

      Zulmara's list: Trinket 2 (empty, `Gain not modelled yet`: Molten Heart of the Mountain is sim-decided, no score), Ranged +23.1, Chest +11.0, Legs +9.3, Wrist +4.7, then four more.
   5. When more exist: `-# {n} more upgrades on the character page.` (`-# 1 more upgrade on the character page.`).
2. **Action row 1**: `Share in channel`, `Someone else`, `Character page` (Link).
3. **Action row 2**, with two or more verified characters: `Another of your characters`.

An empty slot whose band names a pick is kept even though `VerdictFor` gives no `GainDps`. A test pins Zulmara's empty Trinket 2 first.

The `/bis` panel's lead keeps saying each gain is over the next best item; the `/char` sheet's gold legend says its figures are the gain over what is worn (ruling B6).

### A.5.4 Fully geared, no gear, unknown

| State | Body under the header |
|---|---|
| No gear shared | `**Top upgrades** · {band label}`, newline, `No gear shared yet. Upgrades show once the character is synced.` |
| No upgrade | `**Top upgrades** · {band label}`, newline, `Every slot already wears the guide's pick for {band label}.` |
| Unknown or private name | The `char.unknown` card (§G.4.3), V2 (§A.7) |

## A.6 The component and text budget

Every component counts, nested ones included. The counts come from `payloads/*.json`.

| Message | Count | Total | Text |
|---|---|---|---|
| `/bis` page 1 | Container 1 + header Section 3 + Separator 1 + 9 x 3 + row 1 + 4 buttons | 37 | 1,355 |
| `/bis` page 1 after a failed share | the above + the failure Text Display | **38** | 1,386 |
| `/bis` page 2 | 1 + 3 + 1 + 7 x 3 + 1 + 1 + 4 | 32 | 1,070 |
| `/bis slot:` | 1 + 3 + 1 + 3 + 1 + 1 + 6 + 1 + 3 | 20 | 493 |
| `/char` | 1 + 3 + 1 + 1 + 15 + 1 + 4 + 2 | 28 | 1,172 |
| `/char` no gear | | 12 | 241 |
| Shared `/bis`, shared `/char` | | 5, 6 | 72, 185 |

Rules:
- **The ceiling is 38, not 40.** The largest message (page 1 with the failure line) is 38. Every handler test asserts at most 38. The 2 spare are held for Discord changes, never spent.
- **`slotsPerPage` = 9.**
- **Text guard.** While the sum of Text Display contents is over 3,800, the builder drops line 3's source from rows, last row first. The quality dots count: 37 characters a row. Tests:
  - The longest real band in the build never triggers the guard.
  - A synthetic band of 120-character names ends at 3,800 or below.

## A.7 Every message in the `/bis` and `/char` flows is V2

The flag is one-way and an edit that sets it must clear `content` and `embeds` (§0). Two rules follow, and they are one rule from either end:

1. **Every frame these flows draw has flag 32768 from its first draw.** This covers the `/bis` card, answer and pages; the `/char` answer and picks; the `/char` start and unknown cards; and every failure and stale line on them.
2. **Every type 7 update decides by its source message.** The handler reads `interaction.message.flags`.
   - **Source lacks bit 15** (a legacy message: a card drawn before this deploy, or any legacy message a press arrives from): the update body is `{"flags": 32768, "content": null, "embeds": [], "components": […]}`. Discord keeps the ephemeral state of an interaction message on update, so 64 is not sent again.
   - **Source already has bit 15:** the update body carries `flags: 32768` and `components`, and never the keys `content` or `embeds`, because sending them on a V2 message is itself a 400.

   Deferred updates (type 6, then the edit through the press's token, `EditOriginal`) follow the same two cases. `Message.edit` gains a `fromLegacy bool` set from the source's flags.

The guided cards keep their layout and words; only the encoding changes:

| Legacy part | V2 |
|---|---|
| Embed (author, title, description, colour, thumbnail) | A Container with that `accent_color`, holding one Section: Text Display `-# {author}`, `### {title}` and the description, with the crest accessory. With no class, a Text Display alone. |
| Footer and timestamp | Dropped. V2 has neither. |
| Action rows | Unchanged, outside the Container |
| A content line (`char.lead.*`, `card.failure`, `card.stale`) | One top-level Text Display |

The `/bis` card counts 14 components.

`withLiveShare` (share.go) puts a Text Display with `Could not load that. Try again.` first in `components`, never `Content` (`bis-panel-p1-sharefail`, 38).

## A.8 States and interaction parity

| State | What the player sees |
|---|---|
| Default | §A.5 |
| A link button pressed | Discord's external-link confirm, then the page (band, faction, preset). Keyboard focus reaches the buttons in order. Item names are not links (ruling A10). |
| Page button | Discord's spinner, type 7 within 3 s, the other page in place, Share unchanged |
| A press on a legacy card after deploy | The V2 frame replaces it in place (§A.7 rule 2); never a 400 |
| Press after 15 minutes | As the guided cards (§G.6) |
| Tile fails to load | Discord's broken-image square. The row's text still carries everything but the icon. Logged `op tile`. |
| Quality emoji missing | Line 1 without the dot. The quality word stays on line 3. |
| Band missing | `Could not load that. Try again.` |
| Item unknown to the build | Question-mark tile, the data's name, no dot, no quality word or item level |

There is no hover. The link button opens our page, whose shared tooltip is the real item (tenet 2).

# Part B. The character sheet on `Share in channel`

## B.1 The question, and the answer in the channel

The channel asks: "What is this build, or this character, wearing, at a glance?" The answer is one image laid out like the in-game character frame: every slot's icon in its quality ring, the name beside it in its quality colour, the class crest in the middle with the band and faction, and the site's mark. Under it is one link to the full page.

## B.2 References, and what we borrow

- **The in-game character frame (PaperDollFrame).** Borrowed: two slot columns around the centre, and the weapons along the bottom, in the game's order.
  - Shirt and Tabard are not drawn. The BiS data and the addon export have no rows for them (17 export slots), so their sockets could only ever be empty, and empty sockets read as missing gear.
  - The left column's six sit centred against the right column's eight, on the same 72 px pitch, so every icon lines up with an even slot.
  - The round portrait becomes the circular ringed HD crest (one crest language), in the centre where the frame draws the model.
- **Gear-inspection addons that print names beside the frame's slots.** Borrowed: the name in its quality colour beside each icon, the left column's to the right and the right column's to the left, with a quiet meta line under it.
- **Our BiS page header and the OG card.** Borrowed: the gold small-caps kicker and the Cinzel title in the class colour.

## B.3 The layout (`sheet-bis.png`, `sheet-char.png`)

**600 x 844 logical, drawn at scale 2: a 1200 x 1688 PNG** (the mock's files are 492 and 453 KB). Fonts: Cinzel 700, Barlow 400 and 600.

| Region | Box (x, y, w, h), logical px | Content |
|---|---|---|
| Surface | 0, 0, 600, 844 | `--surface-card`, 1 px `--color-line-warm` #3a3326 |
| Kicker | 16, 16, auto, 15 | Barlow 600 12/15, uppercase, tracking 0.12em, `--color-gold` #e5b955. `/bis`: `Best in slot`. `/char`: `Level {n} {Spec} {Class} · {Guild}` |
| Title | 16, 34, auto, 30 | Cinzel 700 24/30, the class colour. `/bis`: `{Spec} {Class}`. `/char`: `{Name}` |
| Site mark | right 16, top 22 | `foreversixty.gg` Barlow 400 13/16 `--color-muted` #9a9484, then the seal at 28 x 28, gap 8 |
| Rule | 16, 76, 568, 1 | `--color-line-warm` |
| Right column | x 344, rows at y = 92 + 72i, i 0 to 7, each 240 x 64 | Hands, Waist, Legs, Feet, Ring 1, Ring 2, Trinket 1, Trinket 2 |
| Left column | x 16, rows at y = 92 + 72i, i 1 to 6, each 240 x 64 | Head, Neck, Shoulder, Back, Chest, Wrist (centred against the right's eight) |
| Centre column | x 264, w 72, y 92 to 660, contents vertically centred | Crest 64 x 64 (native 128 px). `/bis`: `Level {n}` Barlow 600 14/17 `--color-strong`, then `{range}` Barlow 400 13/16 muted. `/char`: `ITEM LEVEL` Barlow 600 11/13 tracking 0.12em muted, then `{x.x}` Cinzel 700 22/26 gold. Then the faction mark 32 x 32 and `Alliance` or `Horde`, Barlow 600 13/16, `--color-alliance` #6fb1ff or `--color-horde` #ff6b5c |
| Weapons, row 9 | y 668: Main hand at x 16 (left style), Off hand at x 344 (right style) | |
| Weapons, row 10 | y 740: Ranged centred, x 180, w 240 (left style) | |
| Footer | 16, 814, 568, 16 | Barlow 400 13/16 muted. `/bis` left: `Updated {d Mon yyyy}` (the band file's `generated_at`). `/char` left: `Gold: gain over what is worn, {band label}`; right: `Synced {d Mon yyyy}` |

**A cell** is 240 x 64, a fixed height. It holds a socket and a text block, gap 8, so the text block is 196 wide.
- Socket: 36 x 36, radius 4, black fill, 2 px ring in `--color-rarity-*`, icon 32 x 32 logical (its native 64 px).
- Name: Barlow 600 20/24 in the quality text colour: `--color-rarity-rare-text` #3d94f0, `--color-rarity-epic-text` #b866f5, the rest their `--color-rarity-*`.
- Meta: Barlow 400 **16/20** `--color-muted`.
- Gold line: Barlow 600 16/20 `--color-gold`.
- Right-style cells right-align the text and put the socket on the outer edge.

**The fit rule** (one rule for the mock's script and the Go renderer; no "…" anywhere):
1. **Name, one line.** Draw at 20 px. While it is wider than 196, shrink 1 px, down to 14 px.
   - Measured on the build's 756 BiS pick names: 94.7% fit one line by 14 px; 92% of all 10,361 wearable items.
   - A name still wider takes two lines at 14/17 (34 px), word-wrapped. On these boards: `Signet Ring of the Bronze Dragonflight` and `Hyper Deluxe Sniper Rifle Mk XVII`.
   - The row height does not change.
2. **Meta lines available.** 2 under a one-line name. 1 under a two-line name. 1 less again when the cell has a gold line.
3. **Meta candidates,** tried in order; the first that fits the lines available wins:
   1. The full source line, word-wrapped.
   2. The short form, built only from real abbreviations and dropped titles:
      - Quest: `Quest · requires level {min_level}` (the quest title is dropped).
      - Instance names by the abbreviations players use: `Blackrock Depths` BRD, `Lower Blackrock Spire` LBRS, `Upper Blackrock Spire` UBRS, `Stratholme` Strat, `Scholomance` Scholo, `Molten Core` MC, `Blackwing Lair` BWL, `Onyxia's Lair` Ony, `Zul'Gurub` ZG, `Ruins of Ahn'Qiraj` AQ20, `Temple of Ahn'Qiraj` AQ40, `Naxxramas` Naxx, `Sunken Temple` ST, `Zul'Farrak` ZF, `Maraudon` Mara, `Scarlet Monastery` SM, `Razorfen Downs` RFD, `Razorfen Kraul` RFK, `Shadowfang Keep` SFK, `Blackfathom Deeps` BFD, `Wailing Caverns` WC, `Ragefire Chasm` RFC, `Gnomeregan` Gnomer. Dire Maul and Uldaman stay whole: "DM" is also the Deadmines.
      - PvP: `PvP rank {n} · {Faction}` (the title is dropped; the bot's own untitled form).
      - Crafted: `Crafted: {Profession}` (the recipe name is dropped).
   3. No meta line.
4. **`/char` meta** is `item level {n}`, one line, and always fits.

On the `/bis` board:
- Head shows the full `Quest: Saving the Best for Last · requires level 58` over two lines.
- Neck falls back to `Quest · requires level 55`, because `1200 Tickets - Amulet of the Darkmoon` does not fit.
- Back shows `Stratholme · Baron Rivendare` in full.

**Cell states:**

| State | Socket | Text |
|---|---|---|
| `/bis` pick | Quality ring, icon | Name, then the source (fit rule) |
| `/char` worn, no gain | Quality ring, icon | Name, then `item level {n}` |
| `/char` worn, gain | Quality ring, icon | Name, `item level {n}`, gold `+{x.x} DPS` (tank `+{x.x} points`) |
| `/char` worn, pick not modelled | Quality ring, icon | Name, `item level {n}`, gold `Gain not modelled yet` |
| `/char` empty slot | Empty: #0d111a fill, #262e40 ring | `Empty` Barlow 400 16/20 muted; gold `Upgrade available` when the band has a pick, or `Gain not modelled yet` when that pick's effect is unmodelled (ruling B9) |
| Off hand, two-handed main hand | Not drawn (ruling B8): the off-hand box is left out, not greyed. The alt text still says `Off hand: none, two-handed main hand` | |
| No pick for the range (`/bis`) | Empty socket | `No pick`, 16/20 muted |
| Item unknown to the build | `inv_misc_questionmark`, #262e40 ring | The data's name in `--color-text` |
| Shirt, Tabard | Not drawn. If a future export carries a worn shirt or tabard, `/char` draws it as left rows 6 and 7 at the in-game position, and the left column grows to eight | |

Emphasis: the quality-coloured names are the sheet's one primary layer. The kicker, title and gold figures are the only other coloured text. Meta, footer and mark are muted.

## B.4 The shared message (`share-bis`, `share-char`)

The follow-up is public: `flags` 32768, multipart, `files[0]` `sheet.png`, `attachments: [{ "id": 0, "filename": "sheet.png", "description": "{alt}" }]`.

1. **Container**, class accent.
   1. **Text Display.** `/bis`: `### {Spec} {Class}` and `Best in slot · {band label} · {Faction}`. `/char`: `### [{Name}]({character URL})` and `Level {n} {Spec} {Class} · {Guild}`.
   2. **Media Gallery**, one item, `attachment://sheet.png`, with the alt text.
   3. `/char` only: `-# Gold figures: the {band label} pick's gain over what is worn.` (tank: `-# Gold figures: the {band label} pick's Forever Score gain over what is worn.`).
2. **Action row**: `All 17 slots` (Link) for `/bis`, `Character page` (Link) for `/char`. Nothing else.

The `/bis` sheet always shows all 17 slots. A shared `/bis slot:{slot}` answer posts the slot panel publicly with no sheet.

## B.5 Flow, timing, failure

1. The press answers type 7 with the disabled `Shared` within 3 s, as today.
2. In `rep.After`, the bot builds the sheet model from the recomputed visitor answer, renders (p95 at most 250 ms, recorded in the lane report), and sends the multipart follow-up.
3. **Render fails:** log `op sheet_render`, then post the visitor's panel publicly with no image (`share-bis-nosheet`). It has one Link button and never Share, the page button or Change picks.
4. **Follow-up fails:** `postShare` as today. The live button is restored with the failure line as a Text Display (§A.7).
5. Size: the PNGs are about 450 to 500 KB. A test asserts under 1.5 MB with the build's longest names.

## B.6 Phone legibility

- The phone client draws the gallery at the container's inner width, about 270 css px on the 390 client (the mock's measure; D20 records the truth). The 600-wide sheet draws there at 0.45:
  - names at 20 px read about 9 css px (14 px names about 6.3);
  - meta at 16 px about 7.2;
  - sockets about 16 css px; the crest about 29.
- The 1200 px file is sharp on a 3x phone, about 810 device px for 270 css.
- Inline on the phone, every icon, quality ring, the crest, the faction, the title and the names read. Meta reads on a tap: Discord's viewer opens the 1200 px image for pinch zoom.
- The trade, stated: a two-column frame of 17 full names cannot also be fully legible at about 270 css px. The fit rule keeps every name whole and on one line for 95% of picks. The sheet stays portrait so the names are as large as the frame allows.
- If D20 shows the desktop client capping the gallery under 440 css px wide, the sheet stays as specified (desktop opens the full image on click), and the lane report records it.

## B.7 Accessibility

- **Alt text.** The Gallery item and the attachment carry the same alt, at most 1024 characters.
  - `/bis`: `Best in slot sheet, {Spec} {Class}, {band label}, {Faction}. {Slot}: {Item}; …` over 17 slots in frame order, with Off hand as `Off hand: none, two-handed main hand`.
  - `/char`: `Character sheet, {Name}, level {n} {Spec} {Class}, {Guild}. {Slot}: {Item}, {+x.x DPS} available; {Slot}: empty; …`. Unmodelled gains are left out of the alt.
  - The mock's alts are 624 and 701 characters.
  - Over 1024: drop the gain clauses first, then cut at the last whole slot and end with `.`. The alt never uses "…".
- **Tile and crest alts** per §A.4.
- **Colour is never the only carrier.** The quality word is on line 3, gains are figures, empty is a word, and the quality dot's emoji name (`fs_q_epic`) is what a screen reader announces.
- **Contrast on #0d0f15:** names use the lightened rare and epic text tokens; muted #9a9484 is above 4.5:1; gold above 7:1.

# Part C. The render path

**Recommendation: one Go renderer inside the API for the sheet and the tiles. The sheet is attached to the follow-up; the tiles are served from a cached endpoint.**

| Option | Verdict | Why |
|---|---|---|
| **Go renderer in the API** (`api/internal/sheet`: `golang.org/x/image/font/opentype`, `golang.org/x/image/webp`, `image/draw`, `image/png`) | **Chosen** | One drawing code for `/bis` and a player's actual gear (`/char`), which exists only at request time. It reads the same data the bot reads (`BisDataDir`, `loot.json`, `BuildItem`), so the sheet cannot disagree with the panel (tenet 8). No second runtime. The tile is the sheet's socket. |
| Build-time static sheets (28 specs x 6 bands x 2 factions = 336) | Rejected | It cannot draw `/char`, so a second renderer would exist anyway. About 160 MB per deploy. A site rebuild would gate every data change. |
| Node `satori` + `resvg` (`web/src/lib/og.ts`) as a sidecar or Pages Function | Rejected | A second runtime and service, and a hop per share. |
| A public image URL for the sheet | Rejected | The attachment is a snapshot that cannot drift, survives our outages, and exposes no player's gear at a guessable URL. Public URLs are used only for the tiles, which are public item art. |

Assets:
- **Fonts:** Cinzel 700 and Barlow 400 and 600 TTF (OFL), `go:embed`, fetched by `make sheet-fonts` from the site's `@fontsource` packages. The site's woff2 subsets are cut to the site's glyphs, so the renderer embeds full Latin TTFs.
- **Crests** (9 x 128), the faction marks (72 px) and the seal: copied by `make sheet-assets` from `web/public` into `api/internal/sheet/assets` and embedded. Copied, never redrawn; a test compares hashes.
- **Item icons:** read from `data/builds/{build}/icons/` (1,704 files, 6.7 MB). The image already carries them: `api/Dockerfile`'s `COPY data/builds /data` puts every build's `icons/` at `/data/{build}/icons`, and `Dockerfile.dockerignore` does not exclude them. A missing icon becomes the question mark (`op sheet_icon_missing` on the tiles route).
- **No game art is resampled up.** Icons draw at 64 px; the crest at 128 px (64 logical); the faction mark 72 to 64; the seal 400 to 56.

Caching:
- **Sheets:** an in-process LRU keyed by SHA-256 of the sheet model's canonical JSON (build, spec, band, preset, faction, per-slot id, name, meta candidates and gold, header strings, date). 32 entries, at most 48 MB.
- **Tiles:** `GET /v1/discord/tiles/{build}/{item}.png`. `{build}` must be a known build folder and `{item}` digits, else 404. `Cache-Control: public, max-age=31536000, immutable`. An in-process LRU of 2,048. Discord's media proxy caches in front. Rate-limited like every public route.

# Copy, verbatim

Strings live in `lookupCopy` and `guideCopy`; tests read each string back from this table. No string uses `soon`, an exclamation mark or `unlock`.

| Key | String |
|---|---|
| `bis.head.title` | `### {Spec} {Class}` |
| `bis.head.line` | `Best in slot · {band label} · {Faction}` |
| `bis.head.page` | `-# {Page range}, {n} of 17 slots. Each gain is over the next best item.` |
| `bis.page.one` | `Head to Legs` |
| `bis.page.two` | `Feet to Ranged` |
| `bis.slot.title` | `### {Slot} · {Spec} {Class}` |
| `bis.slot.next` | `-# The next best` |
| `bis.offhand.twoHand` | `**Off hand** · none: {main hand name} is two-handed.` |
| `bis.slot.none` | `**{Slot}** · no pick for this range.` |
| `item.notModelled` | `Gain not modelled yet` |
| `item.facts` | `-# {Quality} · item level {n} · {source line}` (each part only when known) |
| `item.replaces` | `-# Replaces {worn name}, item level {n}`; without a level `-# Replaces {worn name}` |
| `source.quest` | `Quest: {name} · requires level {min_level}`; with no `min_level` `Quest: {name}` (replaces `Quest: {name} · Level {n}`, here and on the site) |
| `source.quest.short` | `Quest · requires level {min_level}` (sheet only, fit step 2) |
| `source.pvp.short` | `PvP rank {n} · {Faction}` (sheet only) |
| `source.crafted.short` | `Crafted: {Profession}` (sheet only) |
| `char.upgrades.title` | `**Top upgrades** · {band label}` |
| `char.upgrades.byDPS` | `-# By DPS gained, against the guide build.` |
| `char.upgrades.byScore` | `-# By Forever Score gained, against the guide build.` |
| `char.upgrades.more` | `-# {n} more upgrades on the character page.`; for one `-# 1 more upgrade on the character page.` |
| `char.upgrades.none` | `Every slot already wears the guide's pick for {band label}.` |
| `char.noGear` | `No gear shared yet. Upgrades show once the character is synced.` (unchanged) |
| `char.facts` | `-# Item level {x.x} · Readiness: {word} · Synced <t:{unix}:R>` |
| `share.char.gold` | `-# Gold figures: the {band label} pick's gain over what is worn.`; tank `-# Gold figures: the {band label} pick's Forever Score gain over what is worn.` |
| `tile.alt` | `{Item name}, {Quality} {slot} item`; without a quality `{Item name}, {slot} item` |
| `crest.alt` | `{Class} crest` |
| `sheet.bis.kicker` | `Best in slot` |
| `sheet.bis.level` | `Level {n}` |
| `sheet.bis.range` | `Phase 1 raid`; `No buffs or consumables` (level 60 bare); `Levels {n} to {n+9}` (20 to 50) |
| `sheet.char.ilvl` | `Item level` (drawn uppercase) |
| `sheet.empty` | `Empty` |
| `sheet.upgrade` | `Upgrade available` |
| `sheet.none` | `No pick` |
| `sheet.mark` | `foreversixty.gg` |
| `sheet.updated` | `Updated {d Mon yyyy}` |
| `sheet.char.gold` | `Gold: gain over what is worn, {band label}` |
| `sheet.char.synced` | `Synced {d Mon yyyy}` |
| `sheet.alt.bis` | `Best in slot sheet, {Spec} {Class}, {band label}, {Faction}. {Slot}: {Item}; …` |
| `sheet.alt.char` | `Character sheet, {Name}, level {n} {Spec} {Class}, {Guild}. {Slot}: {Item}, {gain} available; …` |
| `sheet.alt.offhand` | `Off hand: none, two-handed main hand` |

Slot words are the `slotChoices` names everywhere: panel line 2, the alt texts, and the page buttons' ranges (`Ring 1`, `Ring 2`, `Trinket 1`, `Trinket 2`, `Main hand`, `Off hand`). The sheet draws no slot words.

Unchanged and reused: `Share in channel`, `Shared`, `Change picks`, `All 17 slots`, `Open on Forever Sixty`, `Character page`, `Someone else`, `Another of your characters`, `Could not load that. Try again.`

Retired:
- `lookupCopy.BisTopLead`, `BisSlotLead`, `BisFooter`, and `BisAuthor` on items
- `ItemLineOne` and `ItemLineTwo` for `/bis` and `/char`
- `TopUpgrades`, `UpgradesBand`, `UpgradesByDPS` and `UpgradesByScore`
- the revision-1 strings `ties the next best` and `same DPS as the pick`, and `fills an empty slot` (ruling A11)
- `topSlots = 5` and `maxUpgrades = 3`

# Rulings

A1. Two pages in the character frame's order (page 1 Head to Legs, page 2 Feet to Ranged). Seventeen rows at 3 components each are 51. The coordinator closed owner question 4: the owner asked for the character screen.

A2. No hover in Discord. The link button opens our page, whose tooltip is the real item.

A3. Quality on the tile's ring, because V2 has one accent colour per container.

A4. Slot and gain at body size; item facts as subtext.

A5. `/char` shows five upgrades: empty slots first, then measured gains, then unmodelled.

A6. Every `/bis` and `/char` message is V2 from its first frame. Every update from a legacy source sends 32768 with `content: null` and `embeds: []`; never on an already-V2 source.

A7. Quality colour on the name comes from an application-emoji dot (§A.4.1). Rejected: ANSI code blocks (no links, monospace, ruled out), and the quality word alone (slower than colour, and already on line 3).

A8. A gain shows only when there is one. A pick whose effect the sim cannot model says `Gain not modelled yet`, never a figure that leaves the effect out.

A9. Quest lines show the level a player must be (`min_level`). The data holds both fields; the quest's own level misled.

A10. (Coordinator, after the persona round.) Item names are plain bold text, not links: Discord draws every link blue, which hid the quality the dot gives. The panel's link button opens the page.

A11. (Coordinator, after the persona round.) An empty `/char` slot shows its real gain, or `Gain not modelled yet` for a pick the sim cannot model (an on-use trinket); `fills an empty slot` is retired.

B1. The sheet goes only on the public share.

B2. No Shirt or Tabard sockets: no data can ever fill them.

B3. Portrait, 600 x 844 logical, a fixed 64 px cell on a 72 px pitch. Names on one line by shrinking to 14 px, two lines at 14/17 only past that. Meta 16/20. No ellipsis anywhere.

B4. Game art is drawn at native size or smaller, never upscaled.

B5. A sheet that fails to render degrades to the panel, public, and is logged.

B6. The `/char` gold legend says what the figure is: `Gold: gain over what is worn, {band label}`. The ruled wording, "Gain over the next best in the level 60 Phase 1 raid", describes the `/bis` panel's figure (which says so in its header line), not the sheet's, and tenet 8 forbids a label that misstates a number.

B7. Item 9533 is `Masons Fraternity Ring` in `items.json`, the class table and the client. It is drawn as the data names it.

B8. (Coordinator, after the persona round.) When the main hand is two-handed the sheet leaves the off-hand box out; it is not drawn greyed. `/bis` draws no Shirt or Tabard. Ring and trinket labels are `Ring 1`, `Ring 2`, `Trinket 1`, `Trinket 2` everywhere.

B9. (Coordinator, after the persona round.) An empty `/char` slot whose pick's gain is not modelled reads `Gain not modelled yet` in gold, not `Upgrade available`.

C1. One Go renderer in the API: attachment for the sheet, cached endpoint for the tiles.

# Acceptance

## Handler tests (`api/internal/discordbot`, strings read back from the copy table)

1. **V2 everywhere, both sources.** Every `/bis` and `/char` frame carries flag 32768 (ephemeral 32832). A type 7 update or deferred edit:
   - whose interaction `message.flags` lacks bit 15 sends `flags: 32768`, `content: null` and `embeds: []`;
   - whose source has bit 15 sends neither key.

   One case each for the `/bis` card (legacy fixture: the guided-card spec's `bis-card-main` payload), a `/char` select pick, a page press, and `Change picks`.
2. **`/bis` pages.**
   - Header lines; rows in frame order; the dot emoji per quality; gain only when the absolute delta is at least 0.05; `Gain not modelled yet` on Head, Feet and Ring 1 of this band.
   - The off-hand two-hander line; tile URL and alt; four buttons in order; page id under 100.
   - A page press answers type 7 with the other page and the same Share id.
   - **Every payload, the live-share failure line included, at or under 38 components** and at or under 3,800 text characters.
3. **`/bis slot:`.** The pick, `-# The next best`, two alternatives with U+2212 deltas or none, `Gain not modelled yet` for an unmodelled pick or alternative, three buttons.
4. **`/char`.** Header clauses present and absent, member-only clauses, the order empty, measured, unmodelled; at most five; `Replaces` with and without a level; `more` singular and plural; no gear and no upgrade; tank in points.
5. **Sources.** `Quest: Saving the Best for Last · requires level 58` (`min_level`, not `level` 60). No `min_level` gives `Quest: {name}`. The site's `questSourceLabel` test changes in step.
6. **Escaping.** Brackets, parentheses and `*` in names stay inside the bold, and no row carries a link.
7. **Text guard.** The longest real band never triggers it; a synthetic band does and ends at 3,800 or below.
8. **Emoji.** With the six `fs_q_*` loaded, line 1 leads with the dot. With one missing, that quality's lines have no dot, and `emoji_missing` names it. `api discord-emoji` uploads only missing names and prints no token.
9. **Share.** Multipart with `sheet.png` and the alt; the gallery `attachment://sheet.png`; a page-2 share renders all 17; a `/bis slot:` share posts the slot panel with no file; a render error posts the no-sheet panel and logs `sheet_render`; a follow-up error restores the live button with the failure Text Display.
10. **Alt texts.** At or under 1024 on the longest real band; the over-length rule on a synthetic case; never "…".

## Renderer tests (`api/internal/sheet`)

11. **Golden PNGs** of the mock's exact data at 1200 x 1688 (`api/internal/sheet/testdata/sheet-{bis,char}.golden.png`, `-update` rewrites them), compared with a per-pixel tolerance for rasterising. Goldens set beside `sheet-*.png` in the lane report. They differ from the mock's sheets only where ruling B8 removes the greyed off-hand box.
12. **Fit rule.**
    - `Raider Gloves` at 20 px, one line.
    - `Signet Ring of the Bronze Dragonflight` on two lines at 14/17, with one meta line.
    - Head's quest source in full over two lines; Neck's short form `Quest · requires level 55`; a synthetic source that fits no candidate drops its meta.
    - No output string contains U+2026.
    - Every icon top sits on the 72 px pitch.
13. **Every cell state of §B.3**, one each.
14. **Icons 1:1.** Asset hashes match `web/public`.
15. **Tiles endpoint:** 72 x 72, ring per quality, unknown item gives the question mark, bad build gives 404, cache headers present.
16. **Budget:** render p95 at or under 250 ms over 50 renders; PNG under 1.5 MB with the longest names.

## Boards the reviewers compare

The mock boards beside the same frames from the real handlers' payloads, through the kit. The kit gains a Components V2 drawing, ported from this mock's `component()` and `markdown()`, with the quality dot. The two sheets go beside their goldens. Any difference in words, order, colour, dot, ring, button style or cell state is a defect; Discord's own chrome is not.

## Discord captures (test server, production application), added to the bot spec's §9 table

| # | What | Desktop client | Phone client |
|---|---|---|---|
| D20 | `/bis` page 1, page 2 and slot form; `/char` own and no gear. Measure the drawn thumbnail, the dot emoji and the gallery's width and height | x | x |
| D21 | A test-only message of 4,001 text characters: record Discord's answer | x | |
| D22 | `Share in channel` on `/bis` from page 2 and on `/char`: the sheet inline, opened full size, pinch zoom on the phone; a forced render failure posting the no-sheet panel | x | x |
| D23 | **A legacy `/bis` card drawn by the build before this one, pressed after deploy: it becomes the V2 answer in place, no 400.** `Change picks` on a V2 answer reopening the V2 card; a page press 20 minutes after the answer; the share-failure restore | x | x |
| D24 | VoiceOver on iOS and NVDA reading a shared sheet's alt, a tile's alt and a dot | x | x |

# Questions only the owner can answer

1. **The sheet on the ephemeral answer too?** Ruling B1 keeps it to the public share, so the panel stays on the first screen. A `Show sheet` button on the ephemeral answer would cost one render per press.
2. **256 px crests?** The crests ship at 128 px, so the sheet's crest is 64 logical, the largest without upscaling. A larger crest needs 256 px art made the way the 512 px faction marks were.
