# Guide accuracy round 14: Hunter

Auditor: hunter. Pages: class index, Beast Mastery, Marksmanship and Survival, plus the card notes in `data/curated/apl/hunter-*.json`.
- Tree: gr14/audit at d7fc51fe7 (main after guide round 13 and sim round 16).
- Client: 1.60.1.70755. Engine pin: f637ec32c (sim/enginever/version.go).
- Read first: the round 1 to round 13 hunter reports with their fix logs, round4/hunter-levels.md, race-sims (rounds 2 to 5), ledger.md through round 13, and design/reviews/2026-11-04-sim-audit-round16/hunter.md.

## What moved since round 13

`git diff 604b78a80..d7fc51fe7` touches no hunter guide, no hunter APL, no hunter BiS file, and no hunter row in `addon-builds.json` (the one change there is warlock-demonology). Sim round 16 hunter found 0 bugs and changed no build, card rank or rotation. Its H1 confirms with client data (SpellMisc Attributes_3 0x100 on every group buff) the page lines that say pets take no stat buffs from other players. The only hunter text change is the coach file `data/curated/coach/hunter-marksmanship.md`, which is not a guide page.

So this round's new work is the round 14 rule: the BiS files in the tree (de044125c) are what ships, and the "provisional BiS" carve-out is over. Every gear, profession, race and stat line was re-read against the live band 60 rows for both factions and both presets.

## Builds and leveling order

The decode is unchanged from round 13 (same strings, same `talents/hunter.json`): BM 31/20/0 leveling and raid, MM 15/31/5 leveling and 5/33/13 raid, SV 5/11/35. Every build is legal at 51 points. The shipped orders (BM Focused Fire 15-16 then Endurance Training from 17, Bestial Wrath at 40; MM Lethal Attacks 10-14, Efficiency 15-19, Rapid Killing 25-26, Rapid Recuperation 33-34; SV Improved Tracking 10-14, Savage Strikes 15-16, Survivalist 17-19, Lone Wolf last) still match BM:147/157, MM:153/159 and SV:134/140.

## Findings

### HUNTER-R14-1 (CONFIRMED, gear/professions): Beast Mastery's profession advice ignores the two bind-on-pickup crafts in its own best in slot list
- **File:** web/src/content/guides/hunter/beast-mastery.md:137
- **Quote:** "Engineering is the raid hunter's profession: scopes, Thorium Grenades for trash, and the target dummies and explosives that save wipes. Skinning and Leatherworking pay for your leveling, since you kill beasts all day, and Leatherworking makes strong leather and mail pieces before raid gear."
- **Evidence:** `bis/hunter-beast-mastery.json` band 60, raid and bare, Alliance and Horde: the main hand is Legionite Glaive (250619), source `crafted:blacksmithing` (loot.json source 3141), `bonding: pickup` (items/hunter.json). Wowhead Forever tooltip: "Binds when picked up". On all four panels a trinket is Frozen Heart of the Mountain (249469), source `crafted:enchanting` (loot.json source 3146), `bonding: pickup`. A bind-on-pickup craft can only be worn by the player who made it, so the page's own level 60 list needs a blacksmith and an enchanter. The section recommends neither. The other crafts in the set bind when equipped and can be bought: Black Dragonscale Helm, Dawn Armor and Mourning Handguards (Leatherworking), and Hyper Deluxe Sniper Rifle Mk XVII (Engineering). Other classes' pages already say this: rogue combat.md:116, shaman enhancement.md:111.
- **Fix:**
  > Enchanting is the profession your best in slot list asks for. Frozen Heart of the Mountain is an Enchanting trinket that binds on pickup, so nobody can make it for you. Legionite Glaive is a Blacksmithing polearm that also binds on pickup. If you are not a blacksmith, take the next two-hander on your best in slot page. Every other craft in your set binds when equipped, the Engineering rifle and the Leatherworking mail included, so buy them. Engineering is the trash and utility pick: Thorium Grenades, and the target dummies and explosives that save wipes. Skinning pays for your leveling, since you kill beasts all day.

### HUNTER-R14-2 (CONFIRMED, gear/professions): the same error on Marksmanship
- **File:** web/src/content/guides/hunter/marksmanship.md:143
- **Quote:** "Engineering is the raid hunter's profession: scopes, Thorium Grenades for trash, and the explosives that save a bad pull. Skinning and Leatherworking pay for your leveling, and Leatherworking makes strong leather and mail pieces before raid gear."
- **Evidence:** `bis/hunter-marksmanship.json` band 60, all four panels: Legionite Glaive (BoP Blacksmithing) in the main hand. Alliance raid and bare, and Horde bare, carry Frozen Heart of the Mountain (249469). Horde raid carries Molten Heart of the Mountain (249470), which Wowhead Forever's tooltip shows as "Binds when picked up", Unique-Equipped. Both are Enchanting crafts (loot.json 3146).
- **Fix:**
  > Enchanting is the profession your best in slot list asks for. The Heart of the Mountain trinkets are Enchanting crafts that bind on pickup, so nobody can make them for you. Legionite Glaive is a Blacksmithing polearm that also binds on pickup. If you are not a blacksmith, take the next two-hander on your best in slot page. Every other craft in your set binds when equipped, the Engineering rifle and the Leatherworking mail included, so buy them. Engineering is the trash pick: Thorium Grenades and the explosives that save a bad pull. Skinning pays for your leveling.

### HUNTER-R14-3 (CONFIRMED, gear/professions): Survival's professions leave out the Enchanting trinkets its list is built on
- **File:** web/src/content/guides/hunter/survival.md:124
- **Quote:** "Engineering gives Survival the AoE it lacks: grenades and explosives on trash. Skinning and Leatherworking pay for your leveling, and Blacksmithing makes melee weapons worth carrying before raid drops."
- **Evidence:** `bis/hunter-survival.json` band 60. Alliance raid wears both Frozen Heart of the Mountain (249469) and Molten Heart of the Mountain (249470). Horde raid, Horde bare and Alliance bare each wear one of them. Both are BoP Enchanting crafts. Heartseeker (12783), the raid off hand on both factions, is a Blacksmithing craft that binds when equipped (`bonding: equip`), so the set never needs a blacksmith. The section names Engineering, Skinning, Leatherworking and Blacksmithing, but not Enchanting, the one profession the list actually needs.
- **Fix:**
  > Enchanting is the profession your best in slot list asks for. Molten Heart of the Mountain and Frozen Heart of the Mountain are Enchanting trinkets that bind on pickup, so nobody can make them for you. Heartseeker, the blacksmith's dagger in your off hand, binds when equipped, so buy one and do not train Blacksmithing for it. Engineering gives Survival the AoE it lacks: grenades and explosives on trash. Skinning pays for your leveling.

### Re-checked against the live tables this round and still true

- **Gear prose against BiS rows (both factions, all bands):**
  - BM:117 "a two-hander with Agility or Intellect beats a fast one-hander": every BM and MM band from 20 to 60 picks a two-hander. At 60 that is Legionite Glaive (16 Agility).
  - SV:106: the raid off hand is Heartseeker (1.9 speed dagger) behind Teebu's Blazing Longsword. Predator's Edge's 50% makes a 50% off hand 75%. "The ranged slot is a stat stick": Precisely Calibrated Boomstick (Alliance) and Blackcrow (Horde) are never fired.
  - SV:136, a two-hander before 20, then two weapons "from 20" (SV:144): band 20 SV already dual-wields.
  - Leather to 40 and mail from 40: no leather or mail row breaks the owner rule.
- **Enchants (pickEnchants re-derived against band 60 raid weights, both factions):**
  - BM and MM: Lesser Agility on the cloak, Greater Agility on gloves and boots, Greater Stats on the chest and Greater Intellect on the bracers. That matches BM:121 and MM:127.
  - SV: Agility on the cloak, gloves and boots, Greater Stats on the chest, Superior Strength on the bracers and Agility on both weapons. That matches SV:110.
- **Consumables (sim/request/raid-consumes.json at 60):**
  - BM and MM: Major Mana Potion, Elixir of the Mongoose, Ground Scorpok Assay and Poached Sunscale Salmon, as BM:125 and MM:131 say.
  - SV: Major Frenzy Potion, Elixir of the Mongoose and Poached Sunscale Salmon, as SV:112 says.
- **Stat prose against band 60 raid weights:**
  - No paragraph ranks stats. Haste: `haste_on_items` is false on every hunter band, so "rare on gear" holds.
  - MM Alliance raid has hit at the cap (baseline 9.3, hit insignificant) and MM Horde has 5% to go. MM:119 covers both ("past the cap hit adds nothing; if your character sheet is still short of it, hit matters until you reach it").
  - SV: Strength scale factor 0.511 against attack power 0.422 is 1.21 attack power a point, so "Strength is attack power directly" holds. Its special hit cap of 9% with Surefooted 3 gives "6% from gear, or 5% as a Dwarf".
- **Race lines:**
  - recommendedRaces is BM [night-elf, orc], MM [night-elf, troll] and SV [night-elf, orc]. Each matches its prose.
  - The measured order (round2/race-sims, guide builds) holds: BM Night Elf, Human, Dwarf and Orc ahead of Troll; MM Night Elf, Human, Dwarf and Troll 745.5 ahead of Orc 743.6; SV Orc 818.8/818.3 against Troll 818.6/819.0, which is "level".
  - Racial values are unchanged in races.json.
  - Skyborne is named only as the paid option.
- **Source links:** all 15 fetched again with a browser user agent. The six foreversixty.gg URLs, the three Blizzard articles and the three Warcraft Tavern pages return 200. The Wowhead Skyborne article, spell=1495 and spell=19306 return 403, which counts as unknown; their labels were confirmed in round 9.
- **Card notes, talent rank text, trainer levels, engine behaviours:** no input changed since round 13's line-by-line pass, and that pass stands.

### SHOULDs (wording, not findings)
- Carried from round 13: MM:159 "so you have it for most hard pulls"; MM:94 five minutes on the leveling build; SV:65 "smaller than Era's"; the Serpent Sting cards' "throws away a tick"; "The Molten Core tome"; BM:81 naming Deflection.
- New: SV:118 calls Orc and Troll "level", but only Orc carries the Recommended tag in the race table. That is true to the measurement, but a reader may ask why Troll is not tagged. "Orc, with Troll level behind it" would read cleaner.

## Decisions

- **D1. A BoP craft in the shipped BiS list makes its profession part of the guide's advice.** This follows the round 14 rule and the precedent on the rogue and shaman pages. A BoE craft does not: Hyper Deluxe Sniper Rifle, the Leatherworking mail and Heartseeker can all be bought.
- **D2. For BM and MM, Enchanting is the profession to recommend, and the Blacksmithing glaive is named with a fallback rather than a second profession.** The trinket is a raid-long pick on every panel. The glaive is a melee stat stick, and the BiS page lists alternatives for it.
- **D3.** All earlier decisions stand, including round 13 D2 (the client's Savage Strikes text over the deep dive's).

## Sims requested

1. **MM BiS on the shipped builds (carried, no longer provisional).** `bis/hunter-marksmanship.json` still ranks on `0053550011503050` (no Sniper Shot). No prose contradicts it this round. When the post-launch re-rank lands, re-read the MM stat table, the enchant table and MM:119.
2. **Carried:**
   - MM race order on the raid build: Night Elf against Human against Dwarf, and Troll against Orc, seeds 1|50001, band 60 raid.
   - Survival Crusader on both weapons against Agility on both.
   - BM and SV race re-runs on the re-ranked gear.
3. **New, low priority:** BM and MM band 60 raid with the best non-BoP two-hander in Legionite Glaive's place, so the fallback in HUNTER-R14-1/2 can name a piece if the gap is large.

## Sim findings

Carried as LIKELY, all from earlier rounds:
- Falcon's melee proc does not need the talent in the engine.
- Sniper Shot's cast time scales with haste in the engine, which the client text does not state.
- Summon Hawk is unmodeled (sim round 16 H2, an owner check in game).
- The stale top-level APL notes, which are not rendered.

## Sources

**Forever Sixty client data, 1.60.1.70755**
- `bis/hunter-{beast-mastery,marksmanship,survival}.json` (de044125c picks): band 60 rows for both factions and both presets, weights, hit_to_cap, haste_on_items.
- `items/hunter.json`, `items.json`: bonding, two_hand, inventory_type for 250619, 249469, 249470, 279273, 252605, 252483, 279257, 275618, 12783, 1728, 2100, 12651.
- `loot.json` sources 3141 (crafted:blacksmithing), 3146 (crafted:enchanting), 3147 (crafted:engineering).
- `enchants.json`, with the pick logic in web/src/lib/guides/gear-extras.ts and enchant-fit.ts re-run in Python.
- `races.json`, `combos.json`, `talents/hunter.json`, `addon-builds.json`.

**Repository**
- sim/request/raid-consumes.json; design/reviews/guide-accuracy/round2/race-sims.md; design/reviews/2026-11-04-sim-audit-round16/hunter.md.
- web/src/content/guides/rogue/combat.md:116 and shaman/enhancement.md:111 (precedent).

**Web** (fetched 2026-11-04)
- Wowhead Forever tooltip, Legionite Glaive: https://nether.wowhead.com/forever/tooltip/item/250619 ("Binds when picked up").
- Wowhead Forever tooltip, Molten Heart of the Mountain: https://nether.wowhead.com/forever/tooltip/item/249470 ("Binds when picked up", Unique-Equipped).
- The 15 frontmatter URLs listed above (status codes as given).

## Fix log (round 14)

- HUNTER-R14-1: fixed, beast-mastery.md:137 (report text). Verified on band 60 bare and raid, both factions: Frozen Heart (Enchanting, BoP) and Legionite Glaive (Blacksmithing polearm, BoP) on all four rows; rifle and Leatherworking pieces bind on equip.
- HUNTER-R14-2: fixed, marksmanship.md:143 (report text). Horde raid wears Molten Heart in place of Frozen; "The Heart of the Mountain trinkets" covers both.
- HUNTER-R14-3: fixed, survival.md:124 (report text). Each row wears one or both Hearts; Heartseeker binds on equip.
- SHOULDs: not applied.

All files were already at updated: 2026-11-04. The prose and sections tests pass (145/145).
