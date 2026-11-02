# Round 11 items lane, appendix: the 165 item ids

These are the ids that round 10's local data run lost from `items.json` (25,853 to 25,688). Sources: wago.tools build 1.60.1.70707, read 2026-11-02 22:09 UTC (`/db2/Item/csv` and `/db2/ItemSparse/csv`, each with and without `useHotfixes=1`); the layer that `python -m pipeline fetch-wago-hotfixes` builds (`raw/wago-hotfixes/`, every wowf build of the 1.60.1 line); build 1.60.1.70753 (listed on wago 2026-11-02 21:35 UTC); the committed hotfix dump (`raw/hotfixes/`, seeded from 70685); and nether.wowhead.com/forever/tooltip/item/<id>.

Columns:
- **Item 07** and **Sparse 07**: the row is in the shipped 70707 Item or ItemSparse export.
- **Sparse 07+hf**: the row is in wago's 70707 ItemSparse with that build's recorded pushes applied.
- **push**: the build whose recorded push carries the ItemSparse row in the layer. "70707" means a 70707 push row.
- **Sparse 53**: the row is in the shipped 70753 ItemSparse.
- **dump**: the row is in the committed `raw/hotfixes/ItemSparse.csv`.
- **WH**: Wowhead Forever has a page for the id.
- **used**: the id is named in a published BiS set (`bis/*.json`), in loot.json, in a guide or curated file (`web/src`, `data/curated`), or in the addon. "simitems" means only the nightly `simitems.json` candidate list names it.

Totals:
- Item 07: 145 of 165. The other 20 Item rows come from the layer.
- Sparse 07: 0. Sparse 07+hf: 39.
- push: 39 from 70707 and 126 from 70532. Sparse 53: 39, the same 39 as the 70707 push.
- dump: 0. WH: 67.
- used: no BiS set, 2 in loot, none in a guide, curated file or the addon.

| id | name | ilvl / req | Item 07 | Sparse 07 | Sparse 07+hf | push | Sparse 53 | dump | WH | used |
|---|---|---|---|---|---|---|---|---|---|---|
| 7190 | Scorched Rocket Boots | 26 / 0 | yes | no | yes | 70707 | yes | no | yes | - |
| 9399 | Precision Arrow | 40 / 35 | yes | no | no | 70532 | no | no | yes | loot |
| 23192 | Tabard of the Scarlet Crusade | 1 / 0 | yes | no | no | 70532 | no | no | yes | loot |
| 211293 | Crimson Trophy Quill | 40 / 1 | yes | no | no | 70532 | no | no | yes | - |
| 212347 | Illari's Key | 1 / 0 | yes | no | no | 70532 | no | no | yes | - |
| 212982 | Squall-breakers Potion | 1 / 30 | yes | no | no | 70532 | no | no | yes | - |
| 215373 | Silver Hand Training Hammer | 31 / 26 | yes | no | no | 70532 | no | no | yes | - |
| 217497 | Narpas Sword | 42 / 34 | yes | no | no | 70532 | no | no | yes | simitems |
| 249519 | Formula: Frozen Heart of the Mountain | 53 / 0 | yes | no | yes | 70707 | yes | no | yes | - |
| 249520 | Formula: Molten Heart of the Mountain | 53 / 0 | yes | no | yes | 70707 | yes | no | yes | - |
| 249521 | Formula: Dormant Heart of the Mountain | 53 / 0 | yes | no | yes | 70707 | yes | no | yes | - |
| 250952 | Disorienting Smog Potion | 25 / 45 | yes | no | no | 70532 | no | no | yes | - |
| 251340 | Plans: Veteran's Chain Shirt | 20 / 0 | yes | no | yes | 70707 | yes | no | yes | - |
| 251342 | Plans: Protector's Chain Shirt | 20 / 0 | yes | no | yes | 70707 | yes | no | yes | - |
| 251343 | Plans: Acolyte's Chain Shirt | 20 / 0 | yes | no | yes | 70707 | yes | no | yes | - |
| 251344 | Plans: Crusader's Chain Shirt | 20 / 0 | yes | no | yes | 70707 | yes | no | yes | - |
| 251345 | Plans: Veteran's Chain Leggings | 25 / 0 | yes | no | yes | 70707 | yes | no | yes | - |
| 251346 | Plans: Guard's Chain Leggings | 25 / 0 | yes | no | yes | 70707 | yes | no | yes | - |
| 251347 | Plans: Protector's Chain Leggings | 25 / 0 | yes | no | yes | 70707 | yes | no | yes | - |
| 251348 | Plans: Acolyte's Chain Leggings | 25 / 0 | yes | no | yes | 70707 | yes | no | yes | - |
| 251349 | Plans: Crusader's Chain Leggings | 25 / 0 | yes | no | yes | 70707 | yes | no | yes | - |
| 251350 | Plans: Veteran's Chain Helm | 30 / 0 | yes | no | yes | 70707 | yes | no | yes | - |
| 251351 | Plans: Guard's Chain Helm | 30 / 0 | yes | no | yes | 70707 | yes | no | yes | - |
| 251352 | Plans: Protector's Chain Helm | 30 / 0 | yes | no | yes | 70707 | yes | no | yes | - |
| 251353 | Plans: Acolyte's Chain Helm | 30 / 0 | yes | no | yes | 70707 | yes | no | yes | - |
| 251354 | Plans: Crusader's Chain Helm | 30 / 0 | yes | no | yes | 70707 | yes | no | yes | - |
| 251395 | Plans: Shining Mithril Helm | 48 / 0 | yes | no | yes | 70707 | yes | no | yes | - |
| 252771 | Pattern: Brawler's Leather Armor | 20 / 0 | yes | no | yes | 70707 | yes | no | yes | - |
| 252772 | Pattern: Trapper's Leather Armor | 20 / 0 | yes | no | yes | 70707 | yes | no | yes | - |
| 252773 | Pattern: Defender's Leather Armor | 20 / 0 | yes | no | yes | 70707 | yes | no | yes | - |
| 252774 | Pattern: Totemic Leather Armor | 20 / 0 | yes | no | yes | 70707 | yes | no | yes | - |
| 252775 | Pattern: Stormrider's Leather Armor | 20 / 0 | yes | no | yes | 70707 | yes | no | yes | - |
| 252776 | Pattern: Wisdom's Leather Armor | 20 / 0 | yes | no | yes | 70707 | yes | no | yes | - |
| 252792 | Pattern: Brawler's Leather Pants | 25 / 0 | yes | no | yes | 70707 | yes | no | yes | - |
| 252793 | Pattern: Trapper's Leather Pants | 25 / 0 | yes | no | yes | 70707 | yes | no | yes | - |
| 252794 | Pattern: Defender's Leather Pants | 25 / 0 | yes | no | yes | 70707 | yes | no | yes | - |
| 252795 | Pattern: Totemic Leather Pants | 25 / 0 | yes | no | yes | 70707 | yes | no | yes | - |
| 252796 | Pattern: Stormrider's Leather Pants | 25 / 0 | yes | no | yes | 70707 | yes | no | yes | - |
| 252797 | Pattern: Wisdom's Leather Pants | 25 / 0 | yes | no | yes | 70707 | yes | no | yes | - |
| 252814 | Pattern: Brawler's Leather Helm | 30 / 0 | yes | no | yes | 70707 | yes | no | yes | - |
| 252815 | Pattern: Trapper's Leather Helm | 30 / 0 | yes | no | yes | 70707 | yes | no | yes | - |
| 252816 | Pattern: Defender's Leather Helm | 30 / 0 | yes | no | yes | 70707 | yes | no | yes | - |
| 252817 | Pattern: Totemic Leather Helm | 30 / 0 | yes | no | yes | 70707 | yes | no | yes | - |
| 252818 | Pattern: Stormrider's Leather Helm | 30 / 0 | yes | no | yes | 70707 | yes | no | yes | - |
| 252819 | Pattern: Wisdom's Leather Helm | 30 / 0 | yes | no | yes | 70707 | yes | no | yes | - |
| 273828 | Recoiling Gnarlybark Staff | 57 / 57 | yes | no | no | 70532 | no | no | yes | - |
| 273830 | Blackthorne Maul | 57 / 57 | yes | no | no | 70532 | no | no | yes | - |
| 285094 | Ravaged Fishing Pole | 30 / 20 | yes | no | no | 70532 | no | no | yes | - |
| 288358 | Gnoll Greataxe | 40 / 1 | yes | no | no | 70532 | no | no | no | simitems |
| 288823 | Arcanist Crown | 30 / 1 | yes | no | no | 70532 | no | no | no | simitems |
| 288824 | Arcanist Leggings | 30 / 1 | yes | no | no | 70532 | no | no | no | simitems |
| 288825 | Arcanist Mantle | 30 / 1 | yes | no | no | 70532 | no | no | no | simitems |
| 288826 | Arcanist Robes | 30 / 1 | yes | no | no | 70532 | no | no | no | simitems |
| 288827 | Arcanist Bindings | 30 / 1 | yes | no | no | 70532 | no | no | no | simitems |
| 288828 | Arcanist Boots | 30 / 1 | yes | no | no | 70532 | no | no | no | simitems |
| 288829 | Arcanist Gloves | 30 / 1 | yes | no | no | 70532 | no | no | no | simitems |
| 288830 | Arcanist Belt | 30 / 1 | yes | no | no | 70532 | no | no | no | simitems |
| 288831 | Felheart Slippers | 30 / 1 | yes | no | no | 70532 | no | no | no | simitems |
| 288832 | Felheart Bracers | 30 / 1 | yes | no | no | 70532 | no | no | no | simitems |
| 288833 | Felheart Gloves | 30 / 1 | yes | no | no | 70532 | no | no | no | simitems |
| 288834 | Felheart Belt | 30 / 1 | yes | no | no | 70532 | no | no | no | simitems |
| 288835 | Felheart Shoulder Pads | 30 / 1 | yes | no | no | 70532 | no | no | no | simitems |
| 288836 | Felheart Horns | 30 / 1 | yes | no | no | 70532 | no | no | no | simitems |
| 288837 | Felheart Robes | 30 / 1 | yes | no | no | 70532 | no | no | no | simitems |
| 288838 | Felheart Pants | 30 / 1 | yes | no | no | 70532 | no | no | no | simitems |
| 288839 | Boots of Prophecy | 30 / 1 | yes | no | no | 70532 | no | no | no | simitems |
| 288840 | Gloves of Prophecy | 30 / 1 | yes | no | no | 70532 | no | no | no | simitems |
| 288841 | Circlet of Prophecy | 30 / 1 | yes | no | no | 70532 | no | no | no | simitems |
| 288842 | Pants of Prophecy | 30 / 1 | yes | no | no | 70532 | no | no | no | simitems |
| 288843 | Robes of Prophecy | 30 / 1 | yes | no | no | 70532 | no | no | no | simitems |
| 288844 | Mantle of Prophecy | 30 / 1 | yes | no | no | 70532 | no | no | no | simitems |
| 288845 | Girdle of Prophecy | 30 / 1 | yes | no | no | 70532 | no | no | no | simitems |
| 288846 | Vambraces of Prophecy | 30 / 1 | yes | no | no | 70532 | no | no | no | simitems |
| 288847 | Nightslayer Chestpiece | 30 / 1 | yes | no | no | 70532 | no | no | no | simitems |
| 288848 | Nightslayer Cover | 30 / 1 | yes | no | no | 70532 | no | no | no | simitems |
| 288849 | Nightslayer Pants | 30 / 1 | yes | no | no | 70532 | no | no | no | simitems |
| 288850 | Nightslayer Shoulder Pads | 30 / 1 | yes | no | no | 70532 | no | no | no | simitems |
| 288851 | Nightslayer Boots | 30 / 1 | yes | no | no | 70532 | no | no | no | simitems |
| 288852 | Nightslayer Bracelets | 30 / 1 | yes | no | no | 70532 | no | no | no | simitems |
| 288853 | Nightslayer Gloves | 30 / 1 | yes | no | no | 70532 | no | no | no | simitems |
| 288854 | Nightslayer Belt | 30 / 1 | yes | no | no | 70532 | no | no | no | simitems |
| 288855 | Cenarion Belt | 30 / 1 | yes | no | no | 70532 | no | no | no | simitems |
| 288856 | Cenarion Boots | 30 / 1 | yes | no | no | 70532 | no | no | no | simitems |
| 288857 | Cenarion Bracers | 30 / 1 | yes | no | no | 70532 | no | no | no | simitems |
| 288858 | Cenarion Gloves | 30 / 1 | yes | no | no | 70532 | no | no | no | simitems |
| 288859 | Cenarion Vestments | 30 / 1 | yes | no | no | 70532 | no | no | no | simitems |
| 288860 | Cenarion Helm | 30 / 1 | yes | no | no | 70532 | no | no | no | simitems |
| 288861 | Cenarion Leggings | 30 / 1 | yes | no | no | 70532 | no | no | no | simitems |
| 288862 | Cenarion Spaulders | 30 / 1 | yes | no | no | 70532 | no | no | no | simitems |
| 288863 | Earthfury Boots | 30 / 1 | yes | no | no | 70532 | no | no | no | simitems |
| 288864 | Earthfury Belt | 30 / 1 | yes | no | no | 70532 | no | no | no | simitems |
| 288865 | Earthfury Gauntlets | 30 / 1 | yes | no | no | 70532 | no | no | no | simitems |
| 288866 | Earthfury Bracers | 30 / 1 | yes | no | no | 70532 | no | no | no | simitems |
| 288867 | Earthfury Vestments | 30 / 1 | yes | no | no | 70532 | no | no | no | simitems |
| 288868 | Earthfury Helmet | 30 / 1 | yes | no | no | 70532 | no | no | no | simitems |
| 288869 | Earthfury Legguards | 30 / 1 | yes | no | no | 70532 | no | no | no | simitems |
| 288870 | Earthfury Epaulets | 30 / 1 | yes | no | no | 70532 | no | no | no | simitems |
| 288871 | Giantstalker's Breastplate | 30 / 1 | yes | no | no | 70532 | no | no | no | simitems |
| 288872 | Giantstalker's Helmet | 30 / 1 | yes | no | no | 70532 | no | no | no | simitems |
| 288873 | Giantstalker's Leggings | 30 / 1 | yes | no | no | 70532 | no | no | no | simitems |
| 288874 | Giantstalker's Epaulets | 30 / 1 | yes | no | no | 70532 | no | no | no | simitems |
| 288875 | Giantstalker's Boots | 30 / 1 | yes | no | no | 70532 | no | no | no | simitems |
| 288876 | Giantstalker's Bracers | 30 / 1 | yes | no | no | 70532 | no | no | no | simitems |
| 288877 | Giantstalker's Belt | 30 / 1 | yes | no | no | 70532 | no | no | no | simitems |
| 288878 | Giantstalker's Gloves | 30 / 1 | yes | no | no | 70532 | no | no | no | simitems |
| 288879 | Lawbringer Chestguard | 30 / 1 | yes | no | no | 70532 | no | no | no | simitems |
| 288880 | Lawbringer Helm | 30 / 1 | yes | no | no | 70532 | no | no | no | simitems |
| 288881 | Lawbringer Legplates | 30 / 1 | yes | no | no | 70532 | no | no | no | simitems |
| 288882 | Lawbringer Spaulders | 30 / 1 | yes | no | no | 70532 | no | no | no | simitems |
| 288883 | Lawbringer Bracers | 30 / 1 | yes | no | no | 70532 | no | no | no | simitems |
| 288884 | Lawbringer Belt | 30 / 1 | yes | no | no | 70532 | no | no | no | simitems |
| 288885 | Lawbringer Boots | 30 / 1 | yes | no | no | 70532 | no | no | no | simitems |
| 288886 | Lawbringer Gauntlets | 30 / 1 | yes | no | no | 70532 | no | no | no | simitems |
| 288887 | Bracers of Might | 30 / 1 | yes | no | no | 70532 | no | no | no | simitems |
| 288888 | Sabatons of Might | 30 / 1 | yes | no | no | 70532 | no | no | no | simitems |
| 288889 | Gauntlets of Might | 30 / 1 | yes | no | no | 70532 | no | no | no | simitems |
| 288890 | Belt of Might | 30 / 1 | yes | no | no | 70532 | no | no | no | simitems |
| 288891 | Breastplate of Might | 30 / 1 | yes | no | no | 70532 | no | no | no | simitems |
| 288892 | Helm of Might | 30 / 1 | yes | no | no | 70532 | no | no | no | simitems |
| 288893 | Legplates of Might | 30 / 1 | yes | no | no | 70532 | no | no | no | simitems |
| 288894 | Pauldrons of Might | 30 / 1 | yes | no | no | 70532 | no | no | no | simitems |
| 288901 | Spinal Reaper | 30 / 1 | yes | no | no | 70532 | no | no | no | simitems |
| 288902 | Hyperthermically Insulated Lava Dredger | 30 / 1 | yes | no | no | 70532 | no | no | no | simitems |
| 288903 | Shadowstrike | 30 / 1 | yes | no | no | 70532 | no | no | no | simitems |
| 288904 | Bonereaver's Edge | 30 / 1 | yes | no | no | 70532 | no | no | no | simitems |
| 288905 | Obsidian Edged Blade | 30 / 1 | yes | no | no | 70532 | no | no | no | simitems |
| 288906 | Staff of Dominance | 30 / 1 | yes | no | no | 70532 | no | no | no | simitems |
| 288907 | Gutgore Ripper | 30 / 1 | yes | no | no | 70532 | no | no | no | simitems |
| 288908 | Aurastone Hammer | 30 / 1 | yes | no | no | 70532 | no | no | no | simitems |
| 288909 | Perdition's Blade | 30 / 1 | yes | no | no | 70532 | no | no | no | simitems |
| 288910 | Sorcerous Dagger | 30 / 1 | yes | no | no | 70532 | no | no | no | simitems |
| 288911 | Striker's Mark | 30 / 1 | yes | no | no | 70532 | no | no | no | simitems |
| 288913 | Blastershot Launcher | 30 / 1 | yes | no | no | 70532 | no | no | no | simitems |
| 288914 | Crimson Shocker | 30 / 1 | yes | no | no | 70532 | no | no | no | simitems |
| 288915 | Azuresong Mageblade | 30 / 1 | yes | no | no | 70532 | no | no | no | simitems |
| 288916 | Drillborer Disk | 30 / 1 | yes | no | no | 70532 | no | no | no | simitems |
| 288917 | Fire Runed Grimoire | 30 / 1 | yes | no | no | 70532 | no | no | no | simitems |
| 288918 | Malistar's Defender | 30 / 1 | yes | no | no | 70532 | no | no | no | simitems |
| 288919 | Band of Accuria | 30 / 1 | yes | no | no | 70532 | no | no | no | simitems |
| 288920 | Medallion of Steadfast Might | 30 / 1 | yes | no | no | 70532 | no | no | no | simitems |
| 288921 | Cloak of the Shrouded Mists | 30 / 1 | yes | no | no | 70532 | no | no | no | simitems |
| 288922 | Dragon's Blood Cape | 30 / 1 | yes | no | no | 70532 | no | no | no | simitems |
| 288923 | Choker of Enlightenment | 30 / 1 | yes | no | no | 70532 | no | no | no | simitems |
| 288924 | Seal of the Archmagus | 30 / 1 | yes | no | no | 70532 | no | no | no | simitems |
| 288925 | Fireproof Cloak | 30 / 1 | yes | no | no | 70532 | no | no | no | simitems |
| 289081 | Noggenfogger Elixir | 30 / 1 | no | no | no | 70532 | no | no | yes | - |
| 289082 | Elixir of Endless Intellect | 30 / 1 | no | no | no | 70532 | no | no | yes | - |
| 289083 | Elixir of Endless Power | 30 / 1 | no | no | no | 70532 | no | no | yes | - |
| 289084 | Elixir of Endless Healing | 30 / 1 | no | no | no | 70532 | no | no | yes | - |
| 289086 | Elixir of Endless Rapidity | 30 / 1 | no | no | no | 70532 | no | no | yes | - |
| 289087 | Elixir of Endless Opportunity | 30 / 1 | no | no | no | 70532 | no | no | yes | - |
| 289088 | Scroll of Endless Fortitude | 30 / 1 | no | no | no | 70532 | no | no | yes | - |
| 289292 | Hammer of Bestial Fury | 35 / 1 | no | no | no | 70532 | no | no | yes | simitems |
| 289294 | Thunderfury, Blessed Blade of the Windseeker | 45 / 1 | no | no | no | 70532 | no | no | yes | simitems |
| 289295 | Sulfuras, Hand of Ragnaros | 45 / 1 | no | no | no | 70532 | no | no | yes | simitems |
| 289296 | Benediction | 40 / 1 | no | no | no | 70532 | no | no | yes | simitems |
| 289297 | Anathema | 40 / 1 | no | no | no | 70532 | no | no | yes | simitems |
| 289302 | Quel'Serrar | 40 / 1 | no | no | no | 70532 | no | no | no | simitems |
| 289303 | Rhok'delar, Longbow of the Ancient Keepers | 40 / 1 | no | no | no | 70532 | no | no | yes | simitems |
| 289321 | Elixir of Endless Mana | 30 / 1 | no | no | no | 70532 | no | no | yes | - |
| 289329 | Elixir of Endless Growth | 30 / 1 | no | no | no | 70532 | no | no | yes | - |
| 289330 | Scroll of Endless Speed | 30 / 1 | no | no | no | 70532 | no | no | yes | - |
| 289334 | Elixir of Endless Experience | 30 / 1 | no | no | no | 70532 | no | no | yes | - |
| 289559 | Enchant Boots - Speed | 1 / 1 | no | no | yes | 70707 | yes | no | yes | - |
| 289560 | Formula: Enchant Boots - Speed | 50 / 0 | no | no | yes | 70707 | yes | no | yes | - |

The id 288887 also appears in `web/src/lib/gear/paperdoll.ts` and `web/src/lib/items/unknown-item-copy.ts` and their tests, but only as an example in doc comments. That is not a use.
