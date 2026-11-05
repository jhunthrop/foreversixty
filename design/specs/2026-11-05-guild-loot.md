# Guild Loot tab: right candidates, game words (experience spec)

Author: ux-designer, 2026-11-05. Round-1 guild audit fixes UX-2, UX-6, WP-2, WP-3.
Amends `design/specs/2026-10-04-guild-page.md` §4.F and §5 (Loot row) and `design/specs/2026-10-25-officer-tools-gate.md` §3.2 (Loot placement only; gate words move to `2026-11-05-guild-gating-and-invites.md`).
Captures audited: `scratchpad/guild-r1/shots/{paid,free,nologs}-*-loot-*`.
Code read: `api/internal/guilds/loot.go` (`onyxiaLoot`, `lootCandidates`, `bisMatch`, `fallbackMatch`), `web/src/components/guild/{GuildLoot,RaidLootItem}.svelte`, `web/src/lib/guild/loot-view.ts`, `data/builds/1.60.1.70790/items.json` (`class_id`, `subclass_id`, `inventory_type`, `required_level`).
No new control and no new token (tenet 10).

## 1. The one question, and the answer above the fold

| Viewer | Question | Answer with no scroll, 360 x 640 and 1280 x 800 |
|---|---|---|
| Officer at the council | "This dropped. Who should get it?" | The first item card's top candidate with its gain, and the officer's `Award` beside it. At 1280 the first two item cards are fully visible under the boss picker. |
| Raider | "Which drops are mine to want?" | The same list. The raider's own candidate rows carry the own-row wash and `(you)`, so they find themselves in one scan. |

What the audit found on that screen: a hunter ranked for a caster book, characters under 60 ranked for level-60 items, the skinned Scale of Onyxia as the "next pick", raw `OFF_HAND` slot words, eight "No one on the roster can use this" cards, and `1 of 1 raid night` beside every name on a guild with one log. Each is fixed below.

## 2. Reference, and what we borrow

- **The in-game item tooltip.** Borrowed: the slot and type words exactly as the client prints them (`Head`, `Leather`, `One-Hand`, `Axe`, `Held In Off-hand`, `Trinket`, `Finger`, `Back`, `Neck`), and the item's icon and quality colour.
- **The in-game Master Looter dropdown.** Only characters who are eligible for the item appear. Borrowed: a candidate list is a list of people who can actually take the item.
- **A loot-council spreadsheet** (Gargul/RCLootCouncil exports): columns name, what they wear, upgrade size, attendance. Borrowed: the column order. Not borrowed: free-text "need/greed" notes.

## 3. Layout, region by region

### 3.1 Header (unchanged structure)

`Card` with `glyph="star"`, title `Loot`, the boss `Select` at the right (`Onyxia · on farm` or `Onyxia · not killed yet`), the `Read-only` label for members. The gate sentence for officers without the plan sits under it (spec C owns the words). The old two lines under the header (`Gain here is item level. The simulator has not checked these.` and the footnote) move: see §3.4.

### 3.2 Who is a candidate (the rule, API side)

A roster character is a candidate for an item only when every test passes. The tests run in `lootCandidates` before tiering; a failed test drops the character from the item, never shows a greyed row.

0. **The item's own class restriction.** The client's `ItemSparse.AllowableClass` mask for the active build (wago.tools, fetched with the build's other tables; `raw/hotfixes/ItemSparse.csv` overrides a row it carries). Until the pipeline writes that column into the build, the builder reads `allowable_class` from `raw/classicdb/item_template.json` (classic-db, already in the repo for every equippable item). A mask that is not all classes keeps only the classes in it: Dragonstalker's Helm lists hunters only.
1. **Verified** and class known (today).
2. **Level.** `level >= item.required_level` (items.json). Level is the row's export level, else its snapshot level (spec A §4.2). Unknown level: not a candidate. Every Onyxia drop with a slot needs 60 today.
3. **Can equip the type.** One table, `api/internal/guilds/equip_rules.go`, generated, never typed. Source: the client's own tables for the active build, the same wago.tools source that decides what a Forever player learns (`SkillLineAbility`): `SkillLine` (category Weapon Skills and Armor Proficiencies) joined to `SkillRaceClassInfo` (the class mask, and the level an armour proficiency is learned) and `SkillLineAbility` (the training spell). The pipeline writes `data/builds/<build>/equip-skills.json` beside `trainables/`; the API copies it into Go with a test pinning the two. Forever changes the Classic answers (shamans train fist weapons, rogues train maces), so no Classic table is used. The builder sanity-checks the generated file against these Classic facts and explains any difference in the lane report: rogues have no two-handed weapon of any kind (Forever rogues do use one-hand axes: the client's Axes mask grants them and the Hack and Slash talent, spell 13960, reads Axe/Sword; amended 2026-11-05, guild round 1); priests have no swords; hunters have no maces; paladins have no daggers, fist weapons or staves; plate is warrior and paladin from 40; mail is hunter and shaman from 40.
   Two rules sit on top of the table:
   - A one-hander in the off hand needs Dual Wield from the same tables. Shamans never dual wield in Forever.
   - Held In Off-hand (`inventory_type` 23) is open to every class; rule 4 decides who wants it.
4. **The spec wants the stats.** For the item-level fallback tier only (tier 1). Each spec belongs to one stat family in `equip_rules.go`: **physical** (Warrior, Rogue, Hunter, Feral, Enhancement, Retribution), **caster** (Mage, Warlock, Shadow, Balance, Elemental), **healer** (Holy and Discipline priest, Holy paladin, Restoration druid and shaman). An item qualifies when it carries a stat of the spec's family (physical: Strength, Agility, attack power; caster: Intellect, spell damage; healer: Intellect, Spirit, healing, mana per 5) or hit or crit (which every family values), or carries no family stat at all (Stamina, armour, resistances, Defense only). A weapon with no family stat (Deathbringer and Vis'kag the Bloodletter carry only a chance on hit) counts only for a character who fights in melee with it: warriors, rogues, Survival hunters (melee in Forever), Enhancement shamans, Retribution and Protection paladins, and any logged tank but a bear. Beast Mastery and Marksmanship hunters, Feral cats, casters and healers never swing it; with no such character on the roster the item goes to the collapsed line (§3.5). Tanks (Protection warrior and paladin, Feral bear) read as physical plus Defense: Defense draws a tank in but shuts nobody out, so Ring of Binding (armour, resistances, Defense) is open to every class at 60 (amended 2026-11-05, guild round 1). This removes a Beast Mastery hunter from Ancient Cornerstone Grimoire and a Holy paladin from Deathbringer. Tier 0 (the item is on the spec's BiS list) skips this test: the list already chose it. The rule is a table, not the curated stat weights, because `data/curated` is not in the API image.
5. **Slotless items have no candidates** (today), and are not drawn as item cards (§3.5).

Known-good fixtures (tenet 8), each a failing test until true:
- Ancient Cornerstone Grimoire never lists a hunter, warrior or rogue.
- Deathbringer and Vis'kag list melee specs that can wield them, never a Beast Mastery or Marksmanship hunter.
- Ring of Binding lists every class at 60.
- Deathbringer never lists a Holy paladin or any priest.
- A level 58 character is on no Onyxia item.
- Dragonstalker's Helm lists hunters only; Judgement Crown paladins only; Stormrage Cover druids only.
- A shaman is never a candidate for a one-hander in the off hand.

### 3.3 Item card (the slotted drops)

One card per drop that has at least one candidate. Order: the boss's loot order as today.

**Head of the card.** `RaidLootItem` as today: 40 px icon with the quality border, the name in its quality colour (epic text `#b866f5`), the shared item tooltip on hover, focus and tap (tenet 2). The eyebrow over the name (11 px `label`, muted) is the game's own slot line, not the enum:

| Drop | Eyebrow |
|---|---|
| Tier helms | `HEAD · CLOTH`, `HEAD · LEATHER`, `HEAD · MAIL`, `HEAD · PLATE` (each from its armour subclass) |
| Shard of the Scale | `TRINKET` |
| Ancient Cornerstone Grimoire | `HELD IN OFF-HAND` |
| Deathbringer | `ONE-HAND · AXE` |
| Vis'kag the Bloodletter | `ONE-HAND · SWORD` |
| Sapphiron Drape | `BACK` |
| Eskhandar's Collar | `NECK` |
| Ring of Binding | `FINGER` |

Weapon and hand words keep the client's casing in sentences: `Main Hand`, `Off Hand`, `One-Hand`, `Held In Off-hand`. The words come from one map `web/src/lib/items/slot-words.ts` (`inventory_type` and subclass to the client's tooltip words). The API sends `inventory_type`, `item_class`, `item_subclass` on each item; the raw `slot` stays for logic only and is never printed.

**Candidate rows.** Up to five, then `Show {n} more` (an underlined 44 px text button) when more qualify. A row, desktop grid `minmax(160px,1.2fr) minmax(0,1.6fr) 120px 112px auto`, gap 16, min 44 px, divided by `--border-soft`:

| Cell | Content |
|---|---|
| Who | `ClassCrestRing` 24, the name 14 px 600 in class colour, `(you)` for the viewer; under it 12 px muted the spec (`Beast Mastery Hunter`). |
| Replaces | 13 px: `Replaces ` then the worn item as the shared item link (16 px icon, quality colour, tooltip), or `Empty slot` muted. For a ring, a trinket or a one-hander a dual wielder could put in either hand, it names the weaker of the two worn items (lower gain verdict, else lower item level), since that is the one the drop pushes out. |
| Gain | `numeric` 14 px. Tier 0 with a sim gain: `+12.4 DPS` (healers `HPS`, tanks `TPS`, the unit from the BiS file; no unit, no figure: `On their BiS list`). Tier 1: `+14 item levels`, 13 px muted; when the candidate wears nothing in the slot (`fills_empty_slot`), `Fills an empty slot` instead, since the gap would only be the item's own level (amended 2026-11-05, guild round 3). |
| Attendance | 13 px muted `{3} of {4} raid nights` (`1 of 1 raid night` singular). Hidden, cell and header both, while the guild has fewer than 1 counted raid night (`attendance.nights === 0` for every row). |
| Action | Officer: `Award` (today's control and gate). Member: nothing. |

**Two tiers, visibly two.** Tier-0 rows come first. When an item has tier-1 rows, a 12 px muted sub-label sits above the first one. The first card on the tab that shows a tier-1 row says it in full, `Ranked by item level only. Not on their best-in-slot list.` (one line, wraps on phone); every later card's tier-1 group reads the short `Item level only`, so the sentence is read once per tab, not once per card (amended 2026-11-05, guild round 2). Tier-1 rows never sit above a tier-0 row. Both tiers sort as today.

**Row states.** Hover (fine pointer): `--row-hover`. The viewer's row: the own-row wash `rgba(229,185,85,.06)` at rest. Awarded: today's winner row and `Awarded to {name}` lines. Focus: the item links and `Award` carry the 2 px gold ring.

### 3.4 Footnote

Moves to the bottom of the tab, one 13 px muted paragraph, max 720 px: `Candidates are characters who can wear the item at their level. Each is ranked by how much it adds over what they wear now. A drop stays open until an officer awards it.`

### 3.5 Drops nobody ranks, collapsed to one line

All slotless drops and every slotted drop with zero candidates collapse into one line at the bottom, above the footnote, drawn by `guild/RaidDisclosure.svelte` (the Svelte twin of `ui/Disclosure.astro` the raid page already uses). Tenet 4 allows it: these drops are not the council's job.

- Closed: 14 px `--text`: `{6} more drops nobody here ranks for` with the chevron. 44 px tall.
- Open: one compact row per drop (icon 24 px with quality border, name in quality colour with the item tooltip, then 13 px muted label):

| Drop | Label |
|---|---|
| Head of Onyxia (Alliance and Horde, listed once) | `Quest: Victory for the Alliance / Victory for the Horde` |
| Mature Black Dragon Sinew | `Quest: Rhok'delar (Hunter)` |
| Draconic for Dummies | the label from §3.6, or `Quest item` |
| Onyxia Hide Backpack | `Bag. Anyone can use it.` |
| Scale of Onyxia | `Skinned from Onyxia. Not a drop.` (this item is skinning loot; it never appears as a pick) |
| A slotted drop with no candidate | `Nobody here can use it yet.` |

The officer's `Award` still works on quest drops and the bag: each open row ends with `Award` for officers, which opens the existing award confirm with a `Select` of every verified level-60 character (the existing `Select` control, 44 px), because a quest item still goes to a person. The picker lists only characters the quest is for: Mature Black Dragon Sinew lists level-60 hunters only (the curated row's `class`); Head of Onyxia lists the guild's faction.

### 3.6 Quest labels are sourced, never typed from memory

New file `data/curated/loot/onyxia-quest-items.json`: `{item_id, label, quest_ids[], class|null, source}` per row. The lane fills it from classic-db (quest ids and names) and records the check in its report (tenet 8). The API copies it into a Go table the same way `onyxiaLoot` is copied (curated data is not in the API image), with a test pinning the two. A quest name is printed only when the lane verifies it on classic-db; otherwise the row keeps its short label without the name (`Quest: Hunter`, or `Quest item` when even the class is unverified). A quest name is printed only when the quest is also open in Forever today: a row whose quest id sits behind a `data/curated/content-phases.json` gate that is not open carries that gate and reads `Quest item` (Draconic for Dummies: The Only Prescription is the Ahn'Qiraj Scepter chain, gate `aq-war-effort`). A test pins the gates to content-phases.json (amended 2026-11-05, guild round 1).

### 3.7 Overview loot line (tenet 2)

The Overview Loot card body becomes one line:
- Before any kill and nobody at 60: `Onyxia's drops need level 60. Nobody here is 60 yet.`
- Before any kill, someone at 60: `Onyxia not killed yet. Top pick today: ` + item link + ` for ` + name.
- After a kill: `Next open drop: ` + item link + ` for ` + name.

The item link is `ItemHover` (16 px icon with quality border, name in quality colour, the shared tooltip). The name is in class colour with a 16 px `ClassCrestRing`. "Top pick" is the highest tier-0 gain across all slotted items, else the highest tier-1 row; the skinned Scale is never a pick.

## 4. States

| Element | Loading | Empty | Error | Gated |
|---|---|---|---|---|
| Tab | today's `Skeleton` (`min-h-[240px]`), plus two item-card skeletons at 168 px | see below | `LoadError` `Loot did not load.` + `Try again` | gate sentence (spec C), `Award` disabled |
| Nobody at 60 | n/a | No item cards. One `EmptyState`: `Onyxia's drops need level 60. Candidates show here as your raiders reach it.` Then the collapsed line. | n/a | sentence still shows |
| Some at 60, no candidates for any item | n/a | `EmptyState`: `Nobody at 60 can use these drops yet.` | n/a | as above |
| Candidate row | n/a | n/a | n/a | `Award` disabled, described by the gate |
| Disclosure | n/a | absent when every drop has candidates | n/a | `Award` disabled |

Before the first kill the list still shows (a council plans ahead), and the boss select reads `Onyxia · not killed yet`.

## 5. Copy, verbatim

| Where | String |
|---|---|
| Slot eyebrows | `HEAD · CLOTH`, `HEAD · LEATHER`, `HEAD · MAIL`, `HEAD · PLATE`, `TRINKET`, `HELD IN OFF-HAND`, `ONE-HAND · AXE`, `ONE-HAND · SWORD`, `BACK`, `NECK`, `FINGER` |
| Candidate | `Replaces `, `Empty slot`, `+{g} DPS`, `+{g} HPS`, `+{g} TPS`, `On their BiS list`, `+{n} item levels`, `Fills an empty slot`, `{p} of {n} raid nights`, `1 of 1 raid night`, `(you)` |
| Tier-1 label | `Ranked by item level only. Not on their best-in-slot list.` (first card); `Item level only` (later cards) |
| More rows | `Show {n} more` |
| Footnote | `Candidates are characters who can wear the item at their level. Each is ranked by how much it adds over what they wear now. A drop stays open until an officer awards it.` |
| Collapsed line | `{n} more drops nobody here ranks for` (`1 more drop nobody here ranks for`) |
| Collapsed labels | `Quest: Victory for the Alliance / Victory for the Horde`, `Quest: Rhok'delar (Hunter)`, `Quest: Hunter`, `Quest item`, `Bag. Anyone can use it.`, `Skinned from Onyxia. Not a drop.`, `Nobody here can use it yet.` |
| Empty, nobody at 60 | `Onyxia's drops need level 60. Candidates show here as your raiders reach it.` |
| Empty, no candidates | `Nobody at 60 can use these drops yet.` |
| Error | `Loot did not load.` |
| Overview | `Onyxia's drops need level 60. Nobody here is 60 yet.`; `Onyxia not killed yet. Top pick today: {item} for {name}`; `Next open drop: {item} for {name}` |

Deleted: `No one on the roster can use this.`, `Gain here is item level. The simulator has not checked these.`, the old footnote, `next pick: Scale of Onyxia`.

## 6. Hierarchy

- Per card the item leads (icon, quality name), then the top candidate. Gain is the one figure; attendance is muted. Tier 1 is visibly secondary (muted figure, its own label).
- `Award` is a secondary control, never filled. The tab has no filled gold button.
- The collapsed line sits last; it is reference, not the job.

## 7. Phone and desktop

- Below 768 a candidate row stacks: line 1 crest, name, gain right-aligned; line 2 `Replaces {item}`; line 3 attendance (when shown) and `Award` right (44 px). Max 120 px a row: 6 px top and bottom padding and no gap between the lines below 768 (round 1 measured 128 px at 8 px padding and a 4 px line gap; amended 2026-11-05, guild round 2).
- Never collapses: the item name and icon, the eyebrow, each candidate's name and gain, the tier-1 label, `Award`.
- 2000 px: the centred 1344 px column; item cards are full column width.

## 8. Performance

- The API answer grows by three small fields per item and the curated quest table. No new request. Item tooltips use the existing `fetchItemTooltipModel` cache.
- Icons reserve their 40 px and 24 px boxes; no missing-icon flash (the skeleton box rule of the profile Gear section).
- Lighthouse budgets unchanged (signed-in tab, not in the lhci set).

## 9. Acceptance screenshots (390, 1440; 2000 for row 1)

| # | State |
|---|---|
| 1 | Officer, Raid guild, five level-60s across classes, after a kill: tier-0 and tier-1 rows, attendance shown, collapsed line closed |
| 2 | Same, collapsed line open, the quest and bag labels, `Award` on the Head of Onyxia row |
| 3 | Guild with zero logs: attendance column absent everywhere |
| 4 | Nobody at 60 (the launch-week state) |
| 5 | Member read-only, own rows washed |
| 6 | Officer gated (spec C sentence, disabled `Award`) |
| 7 | Overview loot line in all three states, tooltip open on the item link (hover at 1440, tap at 390) |
| 8 | Loading and error |

E2E: the fixtures in §3.2 pass against the rendered tab; the Sinew award picker holds hunters only; a ring row names the weaker worn ring; no text node equals `OFF_HAND`, `MAIN_HAND` or `FINGER` in lower or upper snake case; Scale of Onyxia never appears outside the collapsed list.
