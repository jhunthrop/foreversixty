# Client-unconfirmed items: what the client backs and what it does not

Build 1.60.1.70790, client_unconfirmed lane, 5 November 2026 UTC. Rule: `data/pipeline/client_evidence.py`. Ranker margin: `sim/cmd/leveling-bis/confirmed.go`.

## Two separate questions

**"The item exists in the client" and "its stats are verified" are different claims.** This lane settles the first one only.

| Claim | Status for the 1,250 re-admitted items | Evidence |
|---|---|---|
| The item exists in Forever, in classic-db's shape | Confirmed | The client's `Item` row matches class, subclass and inventory type, and its item effects agree with classic-db's |
| Its stats, armor and weapon damage are Forever's | **Not verified** | They are 1.12's. The client's `ItemSparse` row has not reached us |

So a re-admitted row has `client_unconfirmed: false` and keeps `stats_source: "classic-db"`. Two safeguards follow from that:

- **Ranker:** such an item must beat a client-stated alternative by more than the 1% adoption margin (`statsInDoubt`).
- **Site:** the "1.12 stats" tag shows on these rows too, so clearing the flag does not hide the doubt.

## Why 1,354 items were flagged

The client's `ItemSparse` carries name, stats and required level. The server sends that row on demand, when a client first meets the item, so no export or scrape names the item until somebody loots one.

| Where the client has it | Count |
|---|---|
| `Item.csv` only, no `ItemSparse` row | 1,345 |
| In 70790 `ItemSparse`, but as inventory type 0 (Dimensional Ripper, Ultrasafe Transporter) | 2 |
| Hotfix dump (Orb of Deception, a toy in Forever) | 1 |
| The owner's `DBCache` only (Piccolo; build 70338, inventory type 0) | 1 |
| Absent from every client path | 5 |

Other sources add little. Wago's hotfix table for 70790 names 3 of these items and the live Wowhead gear planner names 1. Wowhead's Forever tooltip endpoint had no row for any of the 207 highest-level items a polite scan reached before the disk filled, so the scan was not resumed. Its results are not committed.

## What the rule decided

| Result | Count | Examples |
|---|---|---|
| Re-admitted | 1,250 | Hand of Justice, Burst of Knowledge, Felstriker, Mark of Fordring |
| Still flagged: Forever rework, the client names new spells | 84 | Briarwood Reed, Blackhand's Breadth, Ironfoe, five Royal Seals |
| Still flagged: another shape | 15 | Toys, one-hand to main-hand changes |
| Still flagged: not in the `Item` table | 5 | |

Briarwood Reed and Blackhand's Breadth stay out on purpose. Briarwood Reed is +15 spell power doubled in swamps, not 1.12's +29, and Blackhand's Breadth has a new mark effect.

## How often 1.12 numbers are wrong: 113 of 624

The population is classic-db items of item level 55+ and uncommon or better that have no shipped `ItemSparse` row, but whose Forever numbers reached us through a server push or Wowhead's planner. Those are the items most like the re-admitted set. The rule was then applied to them, and the numbers below are for the ones that pass.

| Where the Forever numbers came from | Pass the rule | Differ from 1.12 | Rate |
|---|---:|---:|---:|
| Server push (hotfix dump, wago) | 573 | 79 | 14% |
| Wowhead gear planner | 51 | 34 | 67% |
| **Total** | **624** | **113** | **18%, about 1 in 6** |

- **The rule does not catch this.** Before the rule the same populations hold 630 items with 116 differing. It removes 6 items, 3 of which differ.
- **The server-push row is the closer comparison.** The Wowhead rows are a small group whose numbers may disagree for reasons of their own, so 1 in 7 is the better estimate and 1 in 6 the figure with everything counted. Neither is a measurement of the re-admitted items themselves.
- **What differs:** mostly primary stats and which rating stats an item has. Weapon damage or speed differs on 4 of 57 server-push weapons.

## Launch gates the re-admission exposed

The flag had been hiding items from content that is not open at launch. A probe through the ranker's own `sourceFor` found these, now closed in `data/curated/content-phases.json` and pinned in `sim/cmd/leveling-bis/testdata/launch-banned.json`:

| Gate | Added | Items it closes |
|---|---|---|
| `dire-maul` | Quests 7498 to 7506 | Royal Seals of Eldre'Thalas (four had been re-admitted) |
| `dire-maul` | Quest 7509 | Quel'Serrar |
| `dire-maul` | Quest 5526 | Milli's Shield, Milli's Lexicon |
| `dire-maul` | Quests 7462, 7877 | Bonecrusher, Backwood Helm, Sedge Boots |
| `scourge-invasion` | Bone Witch, Lumbering Horror, Spirit of the Damned | Bracers and Wristwraps of Undead Slaying |

After these, 544 of the 1,250 have a source open to a level 60 at launch.

Other things checked:

- **Andonisus, Reaper of Souls** (legendary, item level 100, 600 attack power) has an open `world` source but cannot be picked: `legendaryGatedLater` drops every legendary. It does set `_launch_phase_item_level_cap` to 100, which leaves the vendor cost gate's overleveled check nearly inert. Nothing here depends on that.
- **Azuregos and Kazzak drops** (Snowblind Shoes, Eskhandar's Left Claw and others) are already closed by the `world-bosses` gate.
- **Not gated, phase not verified:** Ironvine Belt (22764, Blacksmithing) and Crystalline Threaded Cape (20697, Chillwind Ravager) look like later-patch items. I could not confirm that from the repo, and neither is on the likely-pick list below.

## Pix Xizzix: an inferred judgement call

Pix Xizzix (npc 227853) moved from `forever-new-vendors` to a new `undermine-reals` gate that opens at launch, confidence `inferred`. **Please check this one.** The other four traders stay gated.

**What supports it:**

- Beta vendor records on zockify.com and questing.wiki (read 2026-10-09 per the gate's note) price her stock in Tarnished Undermine Reals, currency 3469, client cap 165: belts 100, rings and cloaks 75, trinkets 125.
- Blizzard's 17 September Q&A brings the Reals back as "kind of like that emblem of justice system", a reason to re-run dungeons, and dungeons are open at launch.

**What is inferred:**

- No Blizzard post names this vendor or where Reals drop.
- The vendor records are from the beta, and nobody had reached 60 on live when this was set.
- Wowhead's rows for her stock state no cost at all, so `loot.json` cannot show the price.

**What it opens:** 36 items, mostly item level 60 epic belts and item level 65 rings, cloaks and relics. If the Reals turn out to be scarce, these become picks a fresh 60 cannot quickly buy. Moving her back is one line in the gate file.

### Does the nightly pick the gate up?

Yes, by reading `bis.yml` and `pipeline/loot/__init__.py`. `loot.json` is deliberately not regenerated or committed here, so the committed file still says `opens: later`.

1. `bis.yml` runs `pipeline loot-merge` with no `if:`, in `data/`, so `merge_loot_files`' default `curated_dir` is `data/curated`.
2. `_build_and_write_loot` rebuilds from the engine fork and the raw caches. It never reads the committed `loot.json`, so the stale `later` does not carry over. No file in `data/curated/loot` or `pipeline/` names her.
3. `apply_content_phases` writes only closed gates. A `launch` gate writes nothing, so her source keeps `opens: null`.
4. `apply_vendor_cost_gate` then looks at her rows: no cost items and no money. The overleveled rule needs an item above the cap, and her epics are item level 60.
5. The next step commits `loot.json` and `manifest.json`, and the ranking hash is computed after that commit, so the ranking sees the new file.

A read-only, in-memory run of steps 3 and 4 on the committed file with her `opens` cleared agrees: all 36 items stay on `vendor:227853` with no `opens`, and the vendor gate moves nothing. That run is not a full rebuild.

**One timing point:** the ranker takes the strictest `opens` of `loot.json` and the gates, so she stays closed for any ranking that reads the old `loot.json`. A plain `bis.yml` run regenerates it first, so that is only a problem for a ranking started by hand from the current tree.

## `items*.json` was patched, not regenerated

A local regeneration also pulled in unrelated changes, so only the flag on the 1,250 rows and the 10 matching manifest hashes changed. The next normalize run produces the same flags from the code.

## Settling the numbers in game

**These lists have not been run in a client.** They need a logged-in Forever character. `GetItemStats` comes from an earlier review's macro; the `C_Item` fallback and the setup line are untested.

- **What the output shows:** stats, armor and weapon DPS. It does not show procs, so read the tooltip for those.
- **Compare against:** the row in `data/builds/1.60.1.70790/items/<class>.json`.
- **The durable fix:** asking the client for an item should make the server send its `ItemSparse` row. Log out, copy `Cache/ADB/enUS/DBCache.bin`, then run `python -m pipeline hotfixes --build <build> --cache <path>` and `normalize`. That promotes each row to client stats, and the doubt margin and "1.12 stats" tag then drop off by themselves. This is the documented hotfix path, but nobody has yet tried it on these items.

### Re-admitted items likely to be picked at band 60

107 items, the five named staples first. The ranker was not run for this. An item is listed if either test holds, among items with a source open at launch:

- It is in the top 5 for its slot (top 10 for rings and trinkets) for any spec, by the weighted stat score using that spec's published band 60 weights: 54 items.
- It is rare or better, item level 55+ or required level 50+, and has an effect, because procs and uses are sim-ranked whatever their stats: 53 items.

"Specs in reach" counts the first test only, out of 28 specs. Stat weights come from the published picks, which predate this change, so the real picks may differ.

**Staples first, as plain `/dump`** (run twice; the first call may return nil while the server sends the row):

```
/dump GetItemStats("item:11815") -- Hand of Justice
/dump GetItemStats("item:11832") -- Burst of Knowledge
/dump GetItemStats("item:12590") -- Felstriker
/dump GetItemStats("item:15411") -- Mark of Fordring
/dump GetItemStats("item:19991") -- Devilsaur Eye
```

**The other 102, in order of how many specs have them in reach** (paste the setup line once per session):

```
/run local g=C_Item.GetItemStats or GetItemStats;FSD=function(...)for _,i in ipairs{...}do local t={i}for k,v in pairs(g("item:"..i)or{})do t[#t+1]=_G[k]or k;t[#t+1]=v end print(unpack(t))end end
/run FSD(10455,22408,13938,15867,12939,22403,16058,18723,22334,13178,11819,12554,12936,15873,13396,19992,14558,12966,13098,13346,18695,18761,13969,14553,11928,13210,12110,12103,13345,22247,22269,13951,14502,18696,18735,13208,14616,18738,1443,12548)
/run FSD(13253,12653,13376,18700,12555,11662,19990,20036,20082,10624,12583,12588,12592,12709,12969,13348,13361,13505,13984,14555,14557,15418,13243,13246,13937,13983,14487,18758,13143,13148,13209,13218,13393,13401,14024,14541,19812,22321,1263,11810)
/run FSD(12582,13183,13198,13204,13212,13375,15853,1168,11669,11808,12621,14531,13285,13286,647,13408,833,11744,11803,11809,11920,12532)
```

| Item | ID | Slot | ilvl | Specs in reach | Source | Effect |
|---|---|---|---:|---:|---|---|
| Hand of Justice | 11815 | trinket | 58 | 11 | Blackrock Depths: Emperor Dagran Thaur | Chance on hit (2%): Gain 1 extra attack. |
| Burst of Knowledge | 11832 | trinket | 58 | 16 | Blackrock Depths: Ambassador Flamelash | Reduces the Mana cost of all spells by 150 for 10 sec. |
| Felstriker | 12590 | main_hand | 63 | 0 | Blackrock Spire: Warchief Rend Blackha | Chance on hit: +100% Critical Strike, +100% Hit. |
| Mark of Fordring | 15411 | neck | 63 | 11 | In Dreams |  |
| Devilsaur Eye | 19991 | trinket | 52 | 3 | The Green Drake | Use: +150 Attack Power, +2% Hit. |
| Chained Essence of Eranikus | 10455 | trinket | 60 | 16 | The Essence of Eranikus | Poison Cloud |
| Ritssyn's Wand of Bad Mojo | 22408 | ranged | 63 | 9 | Stratholme: Baron Rivendare |  |
| Bonecreeper Stylus | 13938 | ranged | 62 | 9 | Scholomance: Darkmaster Gandling |  |
| Prismcharm | 15867 | trinket | 60 | 9 | Urok Doomhowl | Use: +20 Fire Resistance, +20 Nature Resistance, +20 Frost Resistance, |
| Dal'Rend's Tribal Guardian | 12939 | off_hand | 63 | 7 | Blackrock Spire: Warchief Rend Blackha |  |
| Diana's Pearl Necklace | 22403 | neck | 61 | 6 | Stratholme: Cannon Master Willey |  |
| Fordring's Seal | 16058 | finger | 63 | 5 | In Dreams |  |
| Animated Chain Necklace | 18723 | neck | 62 | 5 | Stratholme: Ramstein the Gorger |  |
| Band of Mending | 22334 | finger | 62 | 5 | Stratholme: Balnazzar |  |
| Rosewine Circle | 13178 | finger | 60 | 5 | Blackrock Spire: Urok Doomhowl |  |
| Second Wind | 11819 | trinket | 59 | 5 | Blackrock Depths: Golem Lord Argelmach | Restores 63 Mana every 1 sec for 10 sec. This effect is doubled in Mou |
| Hands of the Exalted Herald | 12554 | hands | 59 | 4 | Blackrock Depths: Princess Moira Bronz |  |
| Battleborn Armbraces | 12936 | wrist | 63 | 3 | Blackrock Spire: Warchief Rend Blackha |  |
| Ragged John's Neverending Cup | 15873 | trinket | 60 | 3 | Mother's Milk | Use: +-5 Strength, +-5 Agility, +10 Stamina. |
| Skul's Ghastly Touch | 13396 | ranged | 57 | 3 | Stratholme: Skul (rare) |  |
| Devilsaur Tooth | 19992 | trinket | 52 | 3 | The Green Drake | Use: +100% Critical Strike. |
| Lady Maye's Pendant | 14558 | neck | 64 | 2 | World drop |  |
| Blackmist Armguards | 12966 | wrist | 63 | 2 | Blackrock Spire: The Beast |  |
| Painweaver Band | 13098 | finger | 63 | 2 | Blackrock Spire: General Drakkisath |  |
| Robes of the Exalted | 13346 | chest | 63 | 2 | Stratholme: Baron Rivendare |  |
| Spellbound Tome | 18695 | off_hand | 62 | 2 | Scholomance: Ras Frostwhisper |  |
| Oblivion's Touch | 18761 | ranged | 62 | 2 | Scholomance: Death Knight Darkreaver |  |
| Loomguard Armbraces | 13969 | wrist | 61 | 2 | Scholomance: Kirtonos the Herald |  |
| Sash of Mercy | 14553 | waist | 61 | 2 | World drop |  |
| Thaurissan's Royal Scepter | 11928 | off_hand | 60 | 2 | Blackrock Depths: Emperor Dagran Thaur |  |
| Pads of the Dread Wolf | 13210 | feet | 60 | 2 | Blackrock Spire: Halycon |  |
| Raincaster Drape | 12110 | back | 58 | 2 | The Rise of the Machines |  |
| Star of Mystaria | 12103 | neck | 63 | 1 | Stratholme: Balnazzar |  |
| Seal of Rivendare | 13345 | finger | 63 | 1 | Stratholme: Baron Rivendare |  |
| Faith Healer's Boots | 22247 | feet | 63 | 1 | Blackrock Spire: Warchief Rend Blackha |  |
| Shadow Prowler's Cloak | 22269 | back | 63 | 1 | Blackrock Spire: General Drakkisath |  |
| Vigorsteel Vambraces | 13951 | wrist | 62 | 1 | Scholomance: Darkmaster Gandling |  |
| Frostbite Girdle | 14502 | waist | 62 | 1 | Scholomance: Ras Frostwhisper |  |
| Intricately Runed Shield | 18696 | off_hand | 62 | 1 | Scholomance: Ras Frostwhisper |  |
| Maleki's Footwraps | 18735 | feet | 62 | 1 | Stratholme: Maleki the Pallid |  |
| Bleak Howler Armguards | 13208 | wrist | 61 | 1 | Blackrock Spire: Gizrul the Slavener |  |
| Bloodmail Boots | 14616 | feet | 61 | 1 | Scholomance |  |
| Carapace Spine Crossbow | 18738 | ranged | 61 | 1 | Stratholme: Nerub'enkan |  |
| Jeweled Amulet of Cainwyn | 1443 | neck | 60 | 1 | World drop |  |
| Magni's Will | 12548 | finger | 60 | 1 | The Princess's Surprise |  |
| Hands of Power | 13253 | hands | 60 | 1 | Blackrock Spire: Quartermaster Zigris |  |
| Riphook | 12653 | ranged | 59 | 1 | Blackrock Spire: Shadow Hunter Vosh'ga |  |
| Royal Tribunal Cloak | 13376 | back | 59 | 1 | Stratholme: Magistrate Barthilas |  |
| Malefic Bracers | 18700 | wrist | 58 | 1 | Scholomance: Ragged Ghoul |  |
| Battlechaser's Greaves | 12555 | feet | 55 | 1 | Blackrock Depths |  |
| Ban'thok Sash | 11662 | waist | 54 | 1 | Blackrock Depths: Ok'thor the Breaker |  |
| Blessed Prayer Beads | 19990 | trinket | 52 | 1 | Blood of Morphaz | Use: +190 Healing. |
| Fire Ruby | 20036 | trinket | 52 | 1 | Destroy Morphaz | Use: +100 Fire Spell Damage. |
| Woestave | 20082 | ranged | 52 | 1 | Blood of Morphaz |  |
| Stinging Bow | 10624 | ranged | 47 | 1 | Sunken Temple: Atal'ai Warrior |  |
| Blackhand Doomsaw | 12583 | main_hand | 63 | 0 | Blackrock Spire: Warchief Rend Blackha | Wound |
| Bonespike Shoulder | 12588 | shoulder | 63 | 0 | Blackrock Spire: Warchief Rend Blackha | Deals 75 damage when you are the victim of a critical melee strike. |
| Blackblade of Shahram | 12592 | main_hand | 63 | 0 | Blackrock Spire: General Drakkisath | Summons the infernal spirit of Shahram. |
| Finkle's Skinner | 12709 | main_hand | 63 | 0 | Blackrock Spire: The Beast | Finkle's Skinner +45 Attack Power against Beasts. |
| Seeping Willow | 12969 | main_hand | 63 | 0 | Blackrock Spire: The Beast | Deals 51 Nature damage every 2 sec to all enemies within an 8 yard rad |
| Demonshear | 13348 | main_hand | 63 | 0 | Stratholme: Balnazzar | Sends a shadowy bolt at the enemy causing 150 Shadow damage and dealin |
| Skullforge Reaver | 13361 | main_hand | 63 | 0 | Stratholme: Baron Rivendare | Drains target for 21 Shadow damage every 1 sec and transfers it to the |
| Runeblade of Baron Rivendare | 13505 | main_hand | 63 | 0 | Stratholme: Baron Rivendare | Increases movement speed by 8% and restores 60 health every 5. |
| Darrowspike | 13984 | main_hand | 63 | 0 | The Lich, Ras Frostwhisper | Blasts a target for 105 Frost damage. |
| Alcor's Sunrazor | 14555 | main_hand | 63 | 0 | World drop | Blasts a target for 105 Fire damage. |
| The Lion Horn of Stormwind | 14557 | trinket | 63 | 0 | World drop | When struck in combat has a 5% chance of increasing all party member's |
| Shimmering Platinum Warhammer | 15418 | main_hand | 63 | 0 | In Dreams | Blasts a target for 252 Nature damage. |
| Argent Defender | 13243 | off_hand | 62 | 0 | The Argent Hold | Has a chance when struck in combat to increase your chance to block by |
| Argent Avenger | 13246 | main_hand | 62 | 0 | The Argent Hold | Increases Attack Power by 100, and Attack Power against Undead by an a |
| Headmaster's Charge | 13937 | main_hand | 62 | 0 | Scholomance: Darkmaster Gandling | Gives 20 additional intellect to party members within 30 yards. |
| Gravestone War Axe | 13983 | main_hand | 62 | 0 | Scholomance: Kirtonos the Herald | Diseases target enemy for 42 Nature damage every 3 sec for 12 sec. |
| Bonechill Hammer | 14487 | main_hand | 62 | 0 | Scholomance: Ras Frostwhisper | Blasts a target for 105 Frost damage. |
| Specter's Blade | 18758 | main_hand | 62 | 0 | Scholomance: Death Knight Darkreaver | +45 Attack Power against Undead. |
| Mark of the Dragon Lord | 13143 | finger | 61 | 0 | Blackrock Spire: Overlord Wyrmthalak | A protective mana shield surrounds the caster absorbing 500 damage. Wh |
| Chillpike | 13148 | main_hand | 61 | 0 | Blackrock Spire: Overlord Wyrmthalak | Blasts a target for 205 Frost damage. |
| Seal of the Dawn | 13209 | trinket | 61 | 0 | The Active Agent | +78 Attack Power against Undead.  Allows the acquisition of Scourgesto |
| Fang of the Crystal Spider | 13218 | main_hand | 61 | 0 | Blackrock Spire: Crystal Fang (rare) | Slows target enemy's casting speed and increases the time between mele |
| Malown's Slam | 13393 | main_hand | 61 | 0 | Stratholme: Postmaster Malown | Knocks target silly for 2 sec and increases Strength by 53 for 20 sec. |
| The Cruel Hand of Timmy | 13401 | main_hand | 61 | 0 | Stratholme: Timmy the Cruel | Chance on hit: +-15 Strength, +-15 Agility, +-15 Stamina, +-15 Intelle |
| Frightalon | 14024 | main_hand | 61 | 0 | Scholomance: Kirtonos the Herald | Chance on hit: +-10 Strength, +-10 Agility, +-10 Stamina, +-10 Intelle |
| Barovian Family Sword | 14541 | main_hand | 61 | 0 | Scholomance: Jandice Barov | Deals 49 Shadow damage every 1 sec for 5 sec. All damage done is then  |
| Rune of the Dawn | 19812 | trinket | 61 | 0 | The Active Agent | Increases damage done to Undead by magical spells and effects by up to |
| Heart of Wyrmthalak | 22321 | trinket | 61 | 0 | Blackrock Spire: Overlord Wyrmthalak | Melee and Ranged attacks have a chance to deal 140 Fire damage. Deals  |
| Brain Hacker | 1263 | main_hand | 60 | 0 | World drop | Chance on hit: +-25 Intellect. |
| Force of Will | 11810 | trinket | 60 | 0 | Blackrock Depths: General Angerforge | When struck in combat has a chance to reduce all damage taken by 51 fo |
| Keris of Zul'Serak | 12582 | main_hand | 60 | 0 | Blackrock Spire: War Master Voone | Inflicts numbing pain that deals 10 Nature damage every 1 sec and incr |
| Venomspitter | 13183 | main_hand | 60 | 0 | Blackrock Spire: Mother Smolderweb | Poisons target for 7 Nature damage every 1 sec for 15 sec. |
| Hurd Smasher | 13198 | main_hand | 60 | 0 | Blackrock Spire: Ghok Bashguud (rare) | Knocks target silly for 1 sec. Orcs and Ogres are stunned for 2 times  |
| Bashguuder | 13204 | main_hand | 60 | 0 | Blackrock Spire: Ghok Bashguud (rare) | Punctures target's armor lowering it by 100. Can be applied up to 3 ti |
| Halycon's Spiked Collar | 13212 | neck | 60 | 0 | Blackrock Spire: Halycon | +48 Attack Power against Beasts. |
| Crest of Retribution | 13375 | off_hand | 60 | 0 | Stratholme: Ramstein the Gorger | Deals 22 Physical damage every time you block. |
| Windreaper | 15853 | main_hand | 60 | 0 | Dawn's Gambit | Inflicts 21 Nature damage every 2 sec for 10 sec. |
| Skullflame Shield | 1168 | off_hand | 59 | 0 | World drop | When struck in melee combat, has a 3% chance of stealing 270 life from |
| Naglering | 11669 | finger | 59 | 0 | Blackrock Depths: Golem Lord Argelmach | When struck in combat inflicts 3 Arcane damage to the attacker. |
| Circle of Flame | 11808 | head | 59 | 0 | Blackrock Depths: Ambassador Flamelash | Channels 79 health into mana every 1 sec for 10 sec. |
| Demonfork | 12621 | main_hand | 59 | 0 | Blackrock Spire: Bannok Grimaxe (rare) | Transfers 10 health every 1 seconds from the target to the caster for  |
| Frightskull Shaft | 14531 | main_hand | 59 | 0 | Scholomance: Rattlegore | Deals 15 Shadow damage every 2 sec for 14 sec and lowers their attack  |
| The Nicker | 13285 | main_hand | 58 | 0 | Blackrock Spire: Spirestone Battle Lor | Wounds the target for 100 damage and deals an additional 13 damage eve |
| Rivenspike | 13286 | main_hand | 58 | 0 | Blackrock Spire: Spirestone Butcher (r | Punctures target's armor lowering it by 100. Can be applied up to 3 ti |
| Destiny | 647 | main_hand | 57 | 0 | World drop | Chance on hit: +200 Strength. |
| Soul Breaker | 13408 | main_hand | 57 | 0 | Stratholme: The Unforgiven | Target enemy loses 10 health and mana every 1 sec for 10 sec. |
| Lifestone | 833 | trinket | 56 | 0 | World drop | Restores 800 Health, increased by 1% for each 1% of your Health missin |
| Bloodfist | 11744 | main_hand | 56 | 0 | Blackrock Depths: Phalanx | Wounds the target for 35 damage. |
| Force of Magma | 11803 | main_hand | 56 | 0 | Blackrock Depths: Bael'Gar | Blasts a target for 182 Fire damage. |
| Flame Wrath | 11809 | main_hand | 56 | 0 | Blackrock Depths: Ambassador Flamelash | Shoots a ring of fire dealing 154 Fire damage to all nearby enemies, a |
| Wraith Scythe | 11920 | main_hand | 56 | 0 | Blackrock Depths | Steals 49 life from target enemy. Deals 2 times as much damage to Unde |
| Spire of the Stoneshaper | 12532 | main_hand | 56 | 0 | Blackrock Depths | Increases armor by 4000 for 10 sec but cannot cast spells or attack fo |

### Still flagged

All 104 class-file rows, Briarwood Reed and Blackhand's Breadth first. They are not candidates today. A row here matters if its stats turn out to match 1.12, or if the rework is worth modelling. Uses the same setup line.

```
/run FSD(12930,13965,14022,14023,18984,13968,13966,22458,18470,12344,13544,18411,18986,23168,10760,10627,5616,10625,17752,17744,11603,12400,11625,12791,10828,11922,11923,18308,13379,18340,18321,13399,12556,1973,11684,18462,18425,18345,13515,13213)
/run FSD(13382,18371,13385,20666,18370,13964,13353,18537,22268,13360,17067,23056,19760,17064,13982,10567,10758,9465,10628,17754,9478,2824,10805,11702,11816,18296,10837,10847,12779,2099,18324,14576,13255,18498,18463,18047,13952,18396,13953,18388)
/run FSD(18755,18538,17068,17075,18467,9483,18483,13959,18472,12423,18313,18326,12413,18383,18351,13955,20130,20517,23051,19982,20503,18471,20035,18469)
```
