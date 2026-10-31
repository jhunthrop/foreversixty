# BiS phase audit: level-60 lists against Forever's launch content

2026-10-30. Read-only investigation. No data or code was changed.

Build audited: `1.60.1.70532` (`web/src/data/active-build.json`). Files: the 28 published `bis/<spec>.json` files on the CDN (`https://foreversixty.gg/data/1.60.1.70532/bis/`, generated 2026-10-26). These are newer than the copies committed in `web/public/data/` (2026-10-24), but every band-60 pick is the same in both. Scope: band 60, both presets (`raid`, labelled "Raid-ready, Phase 1", and `bare`), both factions, the chosen pick in every slot. Alternatives and runners-up were not classified.

## Headline

**Yes. The BiS lists include items that are not available at launch.**

- In the "Raid-ready, Phase 1" lists, **181 picks** (counting each faction separately) come from content that Forever has tied to undated, post-launch raids. That is **95 distinct spec/slot/item rows**, and **all 28 specs** have at least one. Counting the `bare` level-60 lists as well, the total is 359.
- A further **460 raid-preset picks are unknown.** Their sources are Tier 0.5, battleground reputation, the Darkmoon Faire, Season of Discovery-era vendor stock and Forever-new vendors. Blizzard has not said whether any of these is open on 4 November.
- All four items in the review are confirmed as problems: Chains of the Lich, Elemental Focus Band, Dryad's Wrist Bindings and Talisman of Ascendance. Three are in the not-available set. The fourth, Dryad's Wrist Bindings, is in the unknown set, because the battleground reputation it comes from may well be open at launch. Even so, it is an Exalted reward that no fresh level 60 will have on day one.
- The worst single item is not in the review. **Signet Ring of the Bronze Dragonflight**, an item-level 80 Ahn'Qiraj War Effort ring sold by Anachronos, is the finger1 pick in all 28 specs. It used to be gated. That gate was deleted on 2026-10-08, and the source came back in build 70532 with no gate (see finding 1c).

## 1. How the pipeline decides availability

The pipeline's only phase signal is a string field called `opens` on each loot source.

- `data/curated/phases.json` defines the phase boundaries: `pre-beta`, `beta`, `launch` (2026-11-04T23:00Z) and `raids-1` (2026-12-09). `data/pipeline/phases.py` emits a copy to `web/src/data/phases.json`. `later` is a sentinel meaning "no date known".
- **The ranker treats any non-empty `opens` as "not obtainable at any band".** This is `sourceObtainable` in `sim/cmd/leveling-bis/band.go:118-147`: `if s.Opens != "" { return false }`. An item stays in the pool if any one of its sources is ungated.
- `opens` gets set in only a few places:
  1. `data/curated/loot/forever-raid-phases.json`. This is the only curated loot overlay. It gates the seven Era raids (`later`) and Onyxia (`raids-1`).
  2. `data/pipeline/loot/overlay.py:140-157`. Any raid with no date defaults to `later`. The six named world bosses (Kazzak, Azuregos and the four Dragons of Nightmare) default to `raids-1`.
  3. `data/pipeline/loot/sources.py:apply_quest_opens_gate`. This covers quests that need Cenarion Circle (609) or Brood of Nozdormu (910) reputation (`REP_FACTION_RAID_PHASE_OPENS`), and quests whose turn-in item comes only from a gated source.
  4. `apply_crafted_opens_gate`: recipes whose source is gated.
  5. `data/pipeline/loot/vendor_cost.py`: vendor rows whose cost cannot be resolved are moved to `vendor:<id>:later`.
  6. Go fallbacks in `sim/cmd/leveling-bis/data.go:795-905`. These are `raidLockedQuestOpens` (AQ War Effort quests 8756, 8789 and 8790), `repFactionRaidPhaseOpens` (609 and 910) and `worldBossSources`. On top of that, `band.go:legendaryGatedLater` removes every item of quality 5.
- **Nothing gates content by the patch or event it belongs to.** No per-source rule covers the Scourge Invasion, the Silithus Abyssal Council, the Argent Dawn armaments quests, Zandalar (Zul'Gurub) reputation quests, Tier 0.5, the Darkmoon Faire or the battlegrounds. The model only knows "raid" versus "not raid".
- `sim/cmd/leveling-bis/phase60_test.go` (`TestTheBand60PicksComeOnlyFromContentOpenAtLaunch`) passes. It only checks that each pick has a source without `opens`, so it cannot catch a source that should have been gated and was not.

### 1a. Scourge Invasion and the Abyssal Council are not modelled at all

- Balzaphon is a boss inside `dungeon:stratholme`, so he inherits Stratholme's open status. `opens` exists per source, not per boss (`LootBoss` has no `opens` field), so an overlay cannot gate one boss today.
- Prince Skaldrenox and Lord Skwol come from classic-db as ordinary `world:` sources. The "world bosses are raid content" rule in `overlay.py` names only six NPCs, and these two are not among them.

### 1b. Duplicate quest entries leak through the gates that do exist

`loot.json` has two entries for the same `quest_id` on several items. One entry has the reputation gate and the other has nothing. Because one ungated way to get an item is enough, the gate never takes effect:

- Shadowsong's Sorrow (21522) and Ravencrest's Legacy (21520), quest 8745 "Treasure of the Timeless One". One entry has `required_rep_faction: 910, opens: later`. The other entry for the same quest is bare.
- Gnomish Turban of Psychic Might (21517), quest 8620 "The Only Prescription": the same pattern.
- The Eye of Zuldazar (19593, quest 8051) and Jewel of Kajaro (19601, quest 8104): one entry has Zandalar Tribe (270) at revered or exalted and the other has nothing. Faction 270 is also missing from `REP_FACTION_RAID_PHASE_OPENS`, so even the gated entry would not close.

### 1c. The Anachronos gate was removed and the vendor came back

Commit `3592e40a5` (2026-10-08) deleted `data/curated/loot/anachronos-war-effort.json` ("Wowhead no longer lists the vendor"). That file had gated `vendor:15192` to `later`. In build 70532 the source is back, from Wowhead's Forever item-source scrape (`raw/items/item-sources.json`: Anachronos, Tanaris, NPC level 9999, `cost_money: 0`), and it has no gate. Rings 234032, 234033, 234034, 234035 and 234202 are now the finger1 pick in every spec.

The overlay contract treats a stale instruction as an error, which is why the gate was deleted rather than kept for a source that might come back. The cost is that a gate disappears whenever its source drops out of a fetch for one night.

### 1d. Vendor rows from Wowhead carry no standing or currency

Every Wowhead-sourced vendor row checked has `cost_money: 0`, `cost_item_ids: []` and `required_faction_id: null`. This includes Meilosh, the Argent Quartermasters, Pix Xizzix, Creeg Bothunk, Illiyana Moonblaze's 2378xx legs and Mokvar. So the ranker sees Exalted-only and currency-only items as free walk-up purchases. This is not a phase problem, but it compounds one: none of these rows can be gated by reputation either.

## 2. Forever's own phase plan

| Fact | Source | Kind |
|---|---|---|
| Launch 4 Nov 2026, 15:00 PST; "On December 9, level 60 players can … [enter] three raids": Barrow Deeps, Hyjal Summit, Onyxia's Lair | [Blizzard, "Forever goes live globally 4 November"](https://news.blizzard.com/en-gb/article/24315820/world-of-warcraft-forever-goes-live-globally-4-november) | official |
| Raids unlock 9 Dec; then "more raids, dungeons, PvP updates, world content"; "a revamped iconic raid" (undated); "This is not a mode, a season, or a new version of Classic." | [Blizzard, What's Next panel recap](https://news.blizzard.com/en-gb/article/24303862/world-of-warcraft-forever-whats-next-panel-recap) | official |
| Setting is "year one of vanilla", after Warcraft III: Reforged's Forsaken Kingdom campaign and **before Molten Core** | [Blizzard 24301145](https://news.blizzard.com/en-us/article/24301145/), as recorded in `research/01-official-facts.md` §2 | official |
| PvP is seasonal from launch: Honor plus Rank Points, ranks 1 to 14 with rewards, blue and epic Honor sets; it mentions Alterac Valley ("compelled to choose Alterac Valley solely for Honor efficiency"); no battleground list or launch timing | [Blizzard, How PvP Progression in Forever Works](https://worldofwarcraft.blizzard.com/en-us/news/24303316) | official |
| New 15v15 Darkspear Islands battleground, no timing given | [Blizzard news roundup 24310655](https://news.blizzard.com/en-us/article/24310655/wow) | official |
| "Original battlegrounds remain" | What's Next recap, as recorded in `research/01-official-facts.md` §6 | official (via repo notes) |
| Zul'Gurub "has no date on Blizzard's roadmap … no raid opens at launch" | [Endgame Tools, Zul'Gurub](https://endgametools.com/en/wow-forever/raids/zulgurub) | community |
| Developers say Forever "is purposefully sticking to pre-Season of Discovery content for now"; major updates roughly every four months | Search summaries of But Why Tho? and GamerFuzion interviews ([butwhytho.net](https://butwhytho.net/2026/09/wow-classic-devs-world-of-warcraft-forever/), [gamerfuzion](https://www.gamerfuzion.com/world-of-warcraft-forever-update-schedule-every-four-months-northrend-pandaria/)); full articles not read | community, unverified |
| Four battlegrounds "anchor" PvP at launch | [ExitLag blog](https://www.exitlag.com/blog/wow-forever-pvp/) | community, weak |
| Alterac Valley not in the beta | [Warcraft Tavern](https://www.warcrafttavern.com/forever/news/pvp-progression-in-wow-forever-ranks-honor-seasons/) | community |
| Pix Xizzix's stock (for example Serenity Field 272439) is labelled "New in Forever", sold in Stranglethorn Vale; Pix Xizzix is the "Undermine Trader" | [Wowhead Forever item 272439](https://www.wowhead.com/forever/item=272439), [NPC 227853](https://www.wowhead.com/forever/npc=227853) | community database |
| Anachronos exists in the Forever database (Tanaris) | [Wowhead Forever NPC 15192](https://www.wowhead.com/forever/npc=15192) | community database |

**No official Blizzard source says anything** about the Scourge Invasion, Silithus or the Abyssal Council, the Argent Dawn armaments quests, Zul'Gurub or the Zandalar Tribe, Tier 0.5 or the Darkmoon Faire in Forever. The searches returned only Classic, Anniversary and Season of Discovery material for these topics.

So the classification rests on two official facts:

- **No raid is open at launch, and only three raids are dated.**
- **The setting is "before Molten Core".**

Content that exists only to lead into Zul'Gurub, Ahn'Qiraj or Naxxramas is therefore classed as not available:

| Content | Original patch | Why it is classed this way |
|---|---|---|
| Zandalar Tribe quests | 1.7 | Zul'Gurub content |
| Abyssal Council | 1.8 | The Silithus Twilight's Hammer lead-in to Ahn'Qiraj, summoned by a raid group |
| Ahn'Qiraj War Effort and Brood of Nozdormu | 1.9 | Ahn'Qiraj content |
| Scourge Invasion and the Argent Dawn armaments quests | 1.11 | Naxxramas lead-in |

Content with no raid tie is classed as unknown:

| Content | Original patch |
|---|---|
| Tier 0.5 | 1.10 |
| Darkmoon Faire | |
| Battlegrounds | |

The 2019 Classic phase for each item is given only as context. Forever does not follow those phases.

## 3. Classification of every band-60 pick

| Class | Raid preset (A+H) | Bare preset (A+H) | Distinct items |
|---|---|---|---|
| Not available | 181 | 178 | 24 |
| Unknown | 460 | 499 | 149 |
| Available | 261 | 225 | 69 |

Available means dungeon bosses other than Balzaphon, crafted items, world drops, ordinary leveling and world quests, and Sunken Temple. Crafted recipes were not audited one by one; two of them are listed as unknown below.

### Rules used

**Not available**

| Reason key | Sources | Basis |
|---|---|---|
| `scourge` | Stratholme: Balzaphon (Chains of the Lich, 23125) | Scourge Invasion boss (1.11, Naxxramas lead-in; Classic 2019 phase 6). No Forever date. |
| `abyssal` | Prince Skaldrenox (Elemental Focus Band 20682), Lord Skwol (Wavefront Necklace 20685) | Abyssal Council Lords, Silithus (1.8 Ahn'Qiraj lead-in). Summoned with Twilight set pieces and signets, a raid-sized fight. By the pipeline's own rule ("world bosses are raid content") these should be at least `raids-1`, and Silithus itself has no date. |
| `ad-armaments` | "Epic/Superior Armaments of Battle - Friend of the Dawn" (Talisman of Ascendance 22678, Band of Piety/Resolution, Bracers of Hope, The Purifier, Amulet/Medallion of the Dawn) | Argent Dawn Insignia and Valor turn-ins added in 1.11 for the Naxxramas lead-in (Classic 2019 phase 6). Note that Argent Dawn reputation itself is fine at launch. |
| `aq-wareffort` | Vendor Anachronos (Signet Ring of the Bronze Dragonflight 23403x and 234202, item level 80) | Ahn'Qiraj War Effort reward; was gated `later` until 2026-10-08. |
| `aq40-quest` | "Treasure of the Timeless One" (Ravencrest's Legacy, Shadowsong's Sorrow), "The Only Prescription" (Gnomish Turban) | Brood of Nozdormu or Ahn'Qiraj 40 quests; already gated `later` on one duplicate entry. |
| `zg` | "A Collection of Heads", "The Eye of Zuldazar", "The Jewel of Kajaro", "The Rage of Mugamba" | Zandalar Tribe and Zul'Gurub turn-ins. Zul'Gurub is gated `later` in this repo and has no date. |
| `silithus-cc` | "Armaments of War" (Sandworm Skin Gloves 20716) | Cenarion Hold quest in Silithus (1.8 to 1.9 era). Medium confidence. |

**Unknown**

| Reason key | Raid-preset picks | Sources | Note |
|---|---|---|---|
| `t05` | 110 | Tier 0.5 quests ("Saving the Best for Last", "Anthion's Parting Words", "Just Compensation", "An Earnest Proposition"), Mokvar vendor 230317, "Falrin's Vendetta" | 1.10 content (Classic 2019 phase 5). Forever is silent. The 226xxx ids are Season of Discovery re-itemisations kept in the client. Likely later. |
| `bg` | 89 | Silverwing Sentinels / Warsong Outriders (Dryad's Wrist Bindings 19595), League of Arathor / Defilers, Frostwolf / Stormpike, Samuel Hawke, Korrak and Eye of Command quests | Blizzard says the original battlegrounds remain and PvP seasons start at launch, but gives no launch list. Most of these picks are Revered or Exalted, which a launch-day character does not have. |
| `forever-new` | 83 | Pix Xizzix (Stranglethorn Vale, 2724xx stock) | New in Forever; currency and timing unknown. |
| `sod-rep-vendor` | 66 | Meilosh (Timbermaw), Argent Quartermaster Hasana | The reputations exist at launch, but this 2278xx stock is Season of Discovery-era, and the rows carry no standing. |
| `forever-new-bg` | 36 | Creeg Bothunk (Darkspear Islands rewards, 2720xx–2721xx) | Battleground timing is not stated. |
| `dmf` | 28 | Darkmoon Faire (Darkmoon Card: Maelstrom and Heroism, Amulet of the Darkmoon) | No statement on the Faire schedule in Forever. |
| `bg-epic` | 26 | Illiyana Moonblaze, Sentinel's * Leggings 2378xx | Item level 78 epics, well above anything a launch character can otherwise reach. Suspect. |
| `late-recipe` | 16 | Stormshroud Gloves 21278, Titanic Leggings 22385 | Recipe ids from the 1.9 to 1.11 range. Not verified. |
| `pvp-rank` | 6 | Captain Dirgehammer (Knight-Captain's 227xxx) | Forever's rank and Honor vendor stock is not confirmed to be these items. |

### The four items in the review

| Item | Source in the data | Class | Why |
|---|---|---|---|
| Chains of the Lich (23125) | `dungeon:stratholme`, boss Balzaphon | not available | Scourge Invasion boss. The dungeon is open, so the boss inherits "open". |
| Elemental Focus Band (20682) | `world:prince-skaldrenox` (classic-db, 34% chance) | not available | Abyssal Council, Silithus; not covered by the world-boss rule. |
| Dryad's Wrist Bindings (19595) | `rep:silverwing-sentinels:exalted`, vendors 14753 and 14754, `pvp:rank-12:horde` | unknown | Warsong Gulch may be open at launch. Exalted standing is not realistic on day one, and the ranker allows any standing at band 60 (`band.go:146`). |
| Talisman of Ascendance (22678) | Quest 9222 (Argent Dawn, friendly) | not available | 1.11 Naxxramas lead-in turn-in. |

### Per-spec counts, "Raid-ready, Phase 1" preset (both factions)

| Spec | Raid-preset picks (A+H) | Not available | Unknown | Available | Bare-preset not available |
|---|---|---|---|---|---|
| druid-balance | 32 | 8 | 16 | 8 | 8 |
| druid-feral | 32 | 2 | 20 | 10 | 2 |
| druid-feral-bear | 30 | 4 | 22 | 4 | 4 |
| druid-restoration | 30 | 7 | 15 | 8 | 6 |
| hunter-beast-mastery | 32 | 3 | 17 | 12 | 2 |
| hunter-marksmanship | 32 | 4 | 16 | 12 | 2 |
| hunter-survival | 34 | 8 | 18 | 8 | 4 |
| mage-arcane | 32 | 10 | 16 | 6 | 8 |
| mage-fire | 32 | 8 | 15 | 9 | 8 |
| mage-frost | 32 | 10 | 14 | 8 | 8 |
| paladin-holy | 32 | 8 | 12 | 12 | 8 |
| paladin-protection | 32 | 4 | 16 | 12 | 6 |
| paladin-retribution | 30 | 4 | 14 | 12 | 6 |
| priest-discipline | 32 | 6 | 19 | 7 | 4 |
| priest-holy | 32 | 4 | 18 | 10 | 6 |
| priest-shadow | 34 | 8 | 14 | 12 | 8 |
| rogue-assassination | 34 | 5 | 16 | 13 | 6 |
| rogue-combat | 34 | 4 | 18 | 12 | 4 |
| rogue-subtlety | 34 | 4 | 18 | 12 | 4 |
| shaman-elemental | 34 | 8 | 16 | 10 | 8 |
| shaman-enhancement | 30 | 3 | 22 | 5 | 2 |
| shaman-restoration | 30 | 8 | 12 | 10 | 12 |
| warlock-affliction | 32 | 10 | 12 | 10 | 10 |
| warlock-demonology | 32 | 10 | 13 | 9 | 10 |
| warlock-destruction | 32 | 10 | 12 | 10 | 10 |
| warrior-arms | 32 | 4 | 20 | 8 | 4 |
| warrior-fury | 34 | 10 | 18 | 6 | 10 |
| warrior-protection | 34 | 7 | 21 | 6 | 8 |

### Affected picks, "Raid-ready, Phase 1" preset

A+H means the pick is the same for Alliance and Horde.

| Spec | Slot | Item (id) | Source | Faction | Why |
|---|---|---|---|---|---|
| druid-balance | finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | A+H | aq-wareffort |
| druid-balance | finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | A+H | abyssal |
| druid-balance | neck | Chains of the Lich (23125) | Stratholme: Balzaphon [dungeon] | A+H | scourge |
| druid-balance | trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | A+H | ad-armaments |
| druid-feral | finger1 | Signet Ring of the Bronze Dragonflight (234034) | Anachronos [vendor] | A+H | aq-wareffort |
| druid-feral-bear | finger1 | Signet Ring of the Bronze Dragonflight (234035) | Anachronos [vendor] | A+H | aq-wareffort |
| druid-feral-bear | finger2 | Band of Resolution (22680) | Superior Armaments of Battle - Friend of the Dawn [quest] | A+H | ad-armaments |
| druid-restoration | finger1 | Signet Ring of the Bronze Dragonflight (234033) | Anachronos [vendor] | A+H | aq-wareffort |
| druid-restoration | finger2 | Band of Piety (22681) | Superior Armaments of Battle - Friend of the Dawn [quest] | A+H | ad-armaments |
| druid-restoration | neck | Wavefront Necklace (20685) | Lord Skwol [world] | A+H | abyssal |
| druid-restoration | wrist | Bracers of Hope (22667) | Epic Armaments of Battle - Friend of the Dawn [quest] | H | ad-armaments |
| hunter-beast-mastery | finger1 | Signet Ring of the Bronze Dragonflight (234202) | Anachronos [vendor] | A+H | aq-wareffort |
| hunter-beast-mastery | waist | Belt of Preserved Heads (20216) | A Collection of Heads [quest] | A | zg |
| hunter-marksmanship | finger1 | Signet Ring of the Bronze Dragonflight (234202) | Anachronos [vendor] | A+H | aq-wareffort |
| hunter-marksmanship | waist | Belt of Preserved Heads (20216) | A Collection of Heads [quest] | A+H | zg |
| hunter-survival | finger1 | Signet Ring of the Bronze Dragonflight (234202) | Anachronos [vendor] | A+H | aq-wareffort |
| hunter-survival | off_hand | Ravencrest's Legacy (21520) | Treasure of the Timeless One [quest] | A+H | aq40-quest |
| hunter-survival | ranged | The Purifier (22656) | Epic Armaments of Battle - Friend of the Dawn [quest] | A+H | ad-armaments |
| hunter-survival | waist | Belt of Preserved Heads (20216) | A Collection of Heads [quest] | A+H | zg |
| mage-arcane | finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | A+H | aq-wareffort |
| mage-arcane | finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | A+H | abyssal |
| mage-arcane | hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | A+H | silithus-cc |
| mage-arcane | neck | Chains of the Lich (23125) | Stratholme: Balzaphon [dungeon] | A+H | scourge |
| mage-arcane | trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | A+H | ad-armaments |
| mage-fire | finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | A+H | aq-wareffort |
| mage-fire | finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | A+H | abyssal |
| mage-fire | neck | Chains of the Lich (23125) | Stratholme: Balzaphon [dungeon] | A+H | scourge |
| mage-fire | trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | A+H | ad-armaments |
| mage-frost | finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | A+H | aq-wareffort |
| mage-frost | finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | A+H | abyssal |
| mage-frost | hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | A+H | silithus-cc |
| mage-frost | neck | Chains of the Lich (23125) | Stratholme: Balzaphon [dungeon] | A+H | scourge |
| mage-frost | trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | A+H | ad-armaments |
| paladin-holy | finger1 | Signet Ring of the Bronze Dragonflight (234033) | Anachronos [vendor] | A+H | aq-wareffort |
| paladin-holy | finger2 | Band of Piety (22681) | Superior Armaments of Battle - Friend of the Dawn [quest] | A+H | ad-armaments |
| paladin-holy | neck | Wavefront Necklace (20685) | Lord Skwol [world] | A+H | abyssal |
| paladin-holy | wrist | Bracers of Hope (22667) | Epic Armaments of Battle - Friend of the Dawn [quest] | A+H | ad-armaments |
| paladin-protection | finger1 | Signet Ring of the Bronze Dragonflight (234035) | Anachronos [vendor] | A | aq-wareffort |
| paladin-protection | finger1 | Signet Ring of the Bronze Dragonflight (234202) | Anachronos [vendor] | H | aq-wareffort |
| paladin-protection | main_hand | Ravencrest's Legacy (21520) | Treasure of the Timeless One [quest] | A+H | aq40-quest |
| paladin-retribution | finger1 | Signet Ring of the Bronze Dragonflight (234034) | Anachronos [vendor] | A+H | aq-wareffort |
| paladin-retribution | trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | A+H | ad-armaments |
| priest-discipline | finger1 | Signet Ring of the Bronze Dragonflight (234033) | Anachronos [vendor] | A+H | aq-wareffort |
| priest-discipline | finger2 | Band of Piety (22681) | Superior Armaments of Battle - Friend of the Dawn [quest] | A+H | ad-armaments |
| priest-discipline | neck | Wavefront Necklace (20685) | Lord Skwol [world] | A+H | abyssal |
| priest-holy | finger1 | Signet Ring of the Bronze Dragonflight (234033) | Anachronos [vendor] | A+H | aq-wareffort |
| priest-holy | neck | Wavefront Necklace (20685) | Lord Skwol [world] | A+H | abyssal |
| priest-shadow | finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | A+H | aq-wareffort |
| priest-shadow | finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | A+H | abyssal |
| priest-shadow | neck | Chains of the Lich (23125) | Stratholme: Balzaphon [dungeon] | A+H | scourge |
| priest-shadow | trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | A+H | ad-armaments |
| rogue-assassination | finger1 | Signet Ring of the Bronze Dragonflight (234202) | Anachronos [vendor] | A+H | aq-wareffort |
| rogue-assassination | main_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | H | aq40-quest |
| rogue-assassination | off_hand | Ravencrest's Legacy (21520) | Treasure of the Timeless One [quest] | A+H | aq40-quest |
| rogue-combat | finger1 | Signet Ring of the Bronze Dragonflight (234202) | Anachronos [vendor] | A+H | aq-wareffort |
| rogue-combat | main_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | A+H | aq40-quest |
| rogue-subtlety | finger1 | Signet Ring of the Bronze Dragonflight (234202) | Anachronos [vendor] | A+H | aq-wareffort |
| rogue-subtlety | main_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | A+H | aq40-quest |
| shaman-elemental | finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | A+H | aq-wareffort |
| shaman-elemental | finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | A+H | abyssal |
| shaman-elemental | neck | Chains of the Lich (23125) | Stratholme: Balzaphon [dungeon] | A+H | scourge |
| shaman-elemental | trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | A+H | ad-armaments |
| shaman-enhancement | finger1 | Signet Ring of the Bronze Dragonflight (234202) | Anachronos [vendor] | A+H | aq-wareffort |
| shaman-enhancement | waist | Belt of Preserved Heads (20216) | A Collection of Heads [quest] | A | zg |
| shaman-restoration | finger1 | Signet Ring of the Bronze Dragonflight (234033) | Anachronos [vendor] | A+H | aq-wareffort |
| shaman-restoration | finger2 | Band of Piety (22681) | Superior Armaments of Battle - Friend of the Dawn [quest] | A+H | ad-armaments |
| shaman-restoration | neck | Wavefront Necklace (20685) | Lord Skwol [world] | A+H | abyssal |
| shaman-restoration | wrist | Bracers of Hope (22667) | Epic Armaments of Battle - Friend of the Dawn [quest] | A+H | ad-armaments |
| warlock-affliction | finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | A+H | aq-wareffort |
| warlock-affliction | finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | A+H | abyssal |
| warlock-affliction | hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | A+H | silithus-cc |
| warlock-affliction | neck | Chains of the Lich (23125) | Stratholme: Balzaphon [dungeon] | A+H | scourge |
| warlock-affliction | trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | A+H | ad-armaments |
| warlock-demonology | finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | A+H | aq-wareffort |
| warlock-demonology | finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | A+H | abyssal |
| warlock-demonology | hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | A+H | silithus-cc |
| warlock-demonology | neck | Chains of the Lich (23125) | Stratholme: Balzaphon [dungeon] | A+H | scourge |
| warlock-demonology | trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | A+H | ad-armaments |
| warlock-destruction | finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | A+H | aq-wareffort |
| warlock-destruction | finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | A+H | abyssal |
| warlock-destruction | hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | A+H | silithus-cc |
| warlock-destruction | neck | Chains of the Lich (23125) | Stratholme: Balzaphon [dungeon] | A+H | scourge |
| warlock-destruction | trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | A+H | ad-armaments |
| warrior-arms | finger1 | Signet Ring of the Bronze Dragonflight (234034) | Anachronos [vendor] | A+H | aq-wareffort |
| warrior-arms | ranged | The Purifier (22656) | Epic Armaments of Battle - Friend of the Dawn [quest] | A+H | ad-armaments |
| warrior-fury | finger1 | Signet Ring of the Bronze Dragonflight (234034) | Anachronos [vendor] | A+H | aq-wareffort |
| warrior-fury | main_hand | Ravencrest's Legacy (21520) | Treasure of the Timeless One [quest] | A+H | aq40-quest |
| warrior-fury | off_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | A+H | aq40-quest |
| warrior-fury | ranged | The Purifier (22656) | Epic Armaments of Battle - Friend of the Dawn [quest] | A+H | ad-armaments |
| warrior-fury | waist | Belt of Preserved Heads (20216) | A Collection of Heads [quest] | A+H | zg |
| warrior-protection | finger1 | Signet Ring of the Bronze Dragonflight (234035) | Anachronos [vendor] | H | aq-wareffort |
| warrior-protection | finger1 | Signet Ring of the Bronze Dragonflight (234202) | Anachronos [vendor] | A | aq-wareffort |
| warrior-protection | finger2 | Band of Resolution (22680) | Superior Armaments of Battle - Friend of the Dawn [quest] | H | ad-armaments |
| warrior-protection | main_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | A+H | aq40-quest |
| warrior-protection | neck | Rage of Mugamba (19577) | The Rage of Mugamba [quest] | A+H | zg |

Six more not-available items appear only in the `bare` lists:

- Belt of Tiny Heads 20217
- The Eye of Zuldazar 19593
- Jewel of Kajaro 19601
- Gnomish Turban of Psychic Might 21517
- Amulet of the Dawn 22657
- Medallion of the Dawn 22659

The classifier is `classify.py` in the session scratchpad. It is not committed, and it reads the published `bis/*.json`. Its rule table is the "Rules used" table above.

## 4. Fix proposal

The fix belongs in the pipeline's rules and curated inputs. Nothing in `loot.json` or `bis/*.json` should be edited by hand; the nightly job regenerates both. The steps are in order of value.

1. **A curated content-phase file, kept separate from the raid overlay.** For example `data/curated/content-phases.json`, with `sources`, `notes` and gates keyed by stable game ids rather than by generated source ids:
   - **NPCs:**
     - 15192 Anachronos: `later`.
     - The Abyssal Council: the four Lords (Skaldrenox 15203, Skwol 15305, Kazum 15205, Whirlaxis 15204) and the Dukes and Templars: `later`.
     - The Scourge Invasion bosses (Balzaphon, Lord Blackwood, Revanchion, Scorn, Sever, Lady Falther'ess): `later`.
   - **Quests:**
     - The Argent Dawn armaments line (9221 to 9232 and its siblings): `later`.
     - Zul'Gurub turn-ins (8201 and the 8xxx Zandalar quests): `later`.
     - 8381 "Armaments of War": `later`.
     - The three AQ War Effort quests that today live only in Go (`raidLockedQuestOpens`): `later`.
   - **Reputation factions:** add 270 Zandalar Tribe next to 609 and 910, and move `REP_FACTION_RAID_PHASE_OPENS` here so that Python and Go read one table.

   `apply_overlays` or `apply_quest_opens_gate` would apply this file after the generator. Unlike the source-id overlay, a gate whose NPC or quest is missing from tonight's fetch should be skipped and logged, not raised as an error. Otherwise, as with Anachronos, the gate gets deleted and the source returns ungated. The contract still holds as long as one test asserts that every listed id is either present and gated, or absent. The Go hand lists (`raidLockedQuestOpens`, `repFactionRaidPhaseOpens`, `worldBossSources`) then become dead code to delete once `loot.json` carries the gates.

2. **Per-boss gating inside open dungeons.** Balzaphon lives inside `dungeon:stratholme`. Either add `opens` to `LootBoss` and teach `loadLootIndex` to read it, or have `pipeline.loot.classicdb` emit the invasion bosses as their own source (say `event:scourge-invasion`) so that rule 1 can gate it. The second option needs no Go change.

3. **Merge duplicate quest entries.** Collapse `loot.json` `quests[item]` entries by `quest_id`, keeping the strictest gate and the reputation fields, before `apply_quest_opens_gate` runs. Today a bare duplicate overrides a gated one (Treasure of the Timeless One, The Only Prescription, the Zandalar quests).

4. **A test that cannot be satisfied by an empty `opens`.** Extend `phase60_test.go`, or add a data test, with a curated deny-list of item ids that must never be a band-60 pick at launch. Start with the four items in the review plus 23403x and 234202. This turns the veteran review into a regression guard.

5. **For the owner, not the pipeline.** The ranker treats any reputation standing as reachable at band 60 (`band.go:146`), and the vendor rows from Wowhead carry no standing at all. An Exalted battleground or Timbermaw item is therefore a day-one pick. A "fresh 60" cap on standing (Honored, say) belongs in the same rule set as `pvpRankCap`.

Steps 1 to 3 remove all 181 raid-preset picks listed here. Each slot then falls back to the next ungated candidate the next time the nightly BiS job runs. Build 70707 is already being worked on (worktree `mage-70707`), so the fix has to sit upstream of the regeneration, which is where these steps put it.

## 5. Uncertain, needs the owner or an official source

1. **The Scourge Invasion, the Abyssal Council, the Argent Dawn armaments and Zul'Gurub in Forever.** No Blizzard statement either way. The classification above infers "not at launch" from two official facts: no raid opens at launch, and the setting is before Molten Core. The owner should confirm this as the house rule. If Blizzard announces any of this as launch world content, its gate flips to `launch`.
2. **Battlegrounds at launch:** Warsong Gulch, Arathi Basin, Alterac Valley and Darkspear Islands. Blizzard says the original battlegrounds remain and PvP seasons run, but gives no day-one list. This decides 89 + 36 + 26 unknown picks, including Dryad's Wrist Bindings.
3. **Tier 0.5** (1.10: the Anthion and Deliana/Mokvar quests, Mokvar's vendor stock). Silent; 110 picks. Owner call: `later` by default, or open?
4. **The Darkmoon Faire schedule** in Forever (28 picks).
5. **Season of Discovery-era stock on Wowhead's Forever vendors** (Meilosh, the Argent Quartermasters, Captain Dirgehammer, Illiyana Moonblaze's item level 78 legs, Mokvar 230317). Wowhead's vendor data may be inherited from Season of Discovery rather than observed in Forever. The beta was capped at level 30, so no level-60 vendor could have been sniffed. All of these rows have `cost_money: 0` and no currency.
6. **Forever-new vendors** Pix Xizzix and Creeg Bothunk: what currency they take, and whether they are open at launch.
7. **"Phase 1" in the preset label** ("Raid-ready, Phase 1") describes the buff context, but players read it as a gear phase. Worth renaming, or making it true.
8. **Two late recipes**, Stormshroud Gloves 21278 and Titanic Leggings 22385, were flagged only by their id range and were not verified.
