# Raid enchants: evidence, 2026-11-04 (client 1.60.1.70790)

What the BiS ranker's raid pass at level 60 wears (`data/curated/raid-enchants.json`) and why each
one is believed to be obtainable by a fresh 60 at launch.

## Confidence, stated plainly

- **Client tables (certain for this build):** enchant row (SpellItemEnchantment), teaching spell
  (SpellEffect Effect 53), recipe item (ItemEffect + ItemXItemEffect), profession and skill
  (SkillLineAbility, ItemSparse).
- **Vendor, cost and standing (third party, beta era):** the client tables here hold no vendor rows.
  The Merchant's Favor rows come from warcraftforever.games (states standing, built on beta 70009,
  and warns it may change) and stillsixty.com (states price and vendors, marks every row
  "Availability unconfirmed"). They agree on price and vendor. Wowhead returned 403. Nobody in-game
  has confirmed these on the live realm from this repo. Check one vendor in-game.
- **Classic recipe sources (Era data, assumed to carry over):** classic-db (cmangos) via
  `raw/classicdb/sources.json`. This is what Classic Era does, not a statement about Forever.
- **Launch timing:** no gate in `content-phases.json` covers these vendors, the Mathredis quest or
  the drops. The trading companies have no Blizzard launch statement here; a guide says it was
  "updated for WoW Forever release". Inferred, not confirmed.

## Findings beyond the brief

1. **The engine already knew all four enchants.** `simdb.bin` is built from every
   SpellItemEnchantment row, so 7662, 7665, 8218 and 8492 carried their stats all along. Only
   `enchants.json`, which starts from the fork's UI table, lacked them. The competitor run's
   `-bonus` stats were not needed.
2. **The fork's table is stale for Forever's re-tuned enchants.** 21 fork rows name an id the client
   spell no longer applies, for example Gloves - Greater Agility (20012) applies 8206, +10, where the
   fork says 1887, +7. `data/pipeline/loot/client_enchants.py` now rebinds every such row to the
   client's id and stats.
3. **8492 is not a necklace enchant.** It is the Mystic Rugged Armor Kit: chest, legs, hands, feet.
4. **The earlier +9-11% caster gain is overstated.** `competitor/tools/enchant_bis.py` passed 7662's
   +65 spell power as a flat bonus whatever weapon our BiS held, and our casters held a one-hander and
   an off-hand (all nine caster DPS specs' published raid sets, Sageblade or Sageclaw). Real sims must give up the off-hand to get it. Measured below: Fire +6-7% against that run's +9.5%.
5. **Mod 45 "spell power" gives no healing in our engine.** `STAT_BY_MODIFIER_ID[45]` is `spell_power`
   and a heal reads `HealingPower` alone, so healers wear the healing rows, not 7665 or 7662. If mod 45
   does heal in Forever, that is a wider item question than this lane.
6. **Modifier 124 on 8203 (Lesser Spirit) maps to nothing,** so the rebound row has empty stats. No plan
   wears it.

## Merchant's Favor recipes

Vendors: Alynsia (Three Corners, Redridge, Alliance), Beneris (outside the Crossroads, the Barrens,
Horde). Leatherworking pattern: Daniel Stitchsong (Three Corners), Pawani.

| Enchant | Effect | Spell | Recipe | Cost | Standing | Skill |
|---|---|---|---|---|---|---|
| 7662 | +65 spell power, 2H | 1219580 | 274394 | 240 | Honored | Enchanting 300 |
| 7665 | +16 spell power, bracer | 1220624 | 274401 | 240 | Honored | Enchanting 300 |
| 8218 | +6 spell power, necklace | 1249057 | 249502 | 120 | Friendly | Enchanting 210 |
| 8492 | +6 spell power, +40 armor, kit | 1254770 (item 252487) | 252941 | 90 | Honored | Leatherworking 235 |
| 8212 | +100 healing, +34 spell damage, 2H | 1248636 | 249526 | 240 | Honored | Enchanting 300 |
| 8219 | +11 healing, +4 spell damage, necklace | 1249058 | 249503 | 120 | Friendly | Enchanting 210 |
| 1901 | +9 Intellect, bracer | 1248661 | 249538 | 240 | Honored | Enchanting 300 |
| 2563 | +15 Strength, gloves | 1248640 | 274398 | 240 | Honored | Enchanting 300 |
| 7667 | +5 Agility, cloak | 1219587 | 274399 | 240 | Honored | Enchanting 300 |
| 904 | +5 Agility, necklace | 1249059 | 249504 | 120 | Friendly | Enchanting 210 |
| 856 | +5 Strength, necklace | 1249019 | 249496 | 120 | Friendly | Enchanting 210 |
| 8213 | +20 Nature damage, gloves | 1248645 | 249535 | 240 | Honored | Enchanting 300 |
| 7647 | +20 Arcane damage, gloves | 1213626 | 249534 | 240 | Honored | Enchanting 300 |
| 7646 | +20 Holy damage, gloves (table only) | 1213622 | 249536 | 240 | Honored | Enchanting 300 |
| 8211 | +55 spell power, 2H (table only) | 1248607 | 249525 | 240 | Honored | Enchanting 300 |

Costs are Merchant's Favor. The standing was read from warcraftforever.games for all fifteen. The
kit's pattern page shows "90" beside the currency with no label, so 90 is a reading.

## Classic recipes

| Enchant | Effect | Spell | Recipe | Source (classic-db) |
|---|---|---|---|---|
| 1900 | Crusader | 20034 | 16252 | Scarlet Spellbinder, 0.75%, Western Plaguelands |
| 1891 | +4 all stats, chest | 20025 | 16253 | low-rate drops; a rare world recipe, in practice bought or a paid craft |
| 1885 | +9 Strength, bracer | 20010 | 16246 | Deadwind Warlock, 0.9% |
| 8206 | +10 Agility, gloves (+7 in Era) | 20012 | 16219 | Legashi Rogue, 0.9%, Azshara |
| 1887 | +7 Agility, boots | 20023 | 16245 | world drops |
| 929 | +7 Stamina, boots | 20020 | 16215 | world drops |
| 1886 | +9 Stamina, bracer | 20011 | 16251 | world drops |
| 1886 | +9 Stamina, shield (+7 in Era) | 20017 | 16217 | Mythrin'dir, Daniel Bartlett, vendors |
| 1889 | +70 armor, cloak | 20015 | 16224 | Lorelae Wintersong, vendor, Moonglade |
| 1883 | +7 Intellect, bracer | 20008 | 16214 | Vilebranch Shadowcaster, 1.96% |
| 664 | Sniper Scope, +7 ranged damage (engine) | 12460 | Schematic 10608 | world drops |
| 1506-1509 | +8 stat, Lesser Arcanum, head or legs | 15397-15404 | 11645-11648 | Libram of Voracity quest; the librams drop from 179 sources incl. Azshara and Blackrock Depths. Mathredis Firestar, Burning Steppes, is from memory of Classic, not this data |

## Left out on purpose

| Enchant | Why |
|---|---|
| Arcanum of Focus, Rapidity, Protection (2544, 2543, 2545) | Dire Maul, gated `later` in content-phases.json |
| Enchant Weapon - Spell Power, Healing Power (2504, 2505) | Molten Core drops: raid |
| Mighty Intellect, Weapon Strength (2568, 2563 weapon) | Lokhtos, Thorium Brotherhood standing; not a fresh 60 |
| Weapon Agility, 2H Agility (2564, 2646) | Meilosh, gated in `season-of-discovery-vendor-stock` |
| Bracer Healing Power (2566) | Argent Quartermasters, same gate |
| Gloves Fire, Frost, Shadow Power (2614-2616), Healing Power (2617) | Ahn'Qiraj |
| Zandalar, Scourge and class head enchants | Zul'Gurub and Naxxramas |
| Bracer Superior Agility 7656, Off-Hand Superior Intellect 7659, Shield Excellent Stamina 7663, Grand Crusader 7940, 2H Strength 8215, 2H Lesser Agility 2618 | Real client recipes, but no source found. Good candidates once one is confirmed; 7656 would replace the agility specs' bracer |
| Biznicks Accurascope 2523 | Molten Core schematic |

## Measured gain

Band 60 raid pass, the ranker run end to end (`-spec <spec> -bands 60`) on this branch, once with
every plan emptied and once with the curated plan. The set may re-pick between the two runs, so this
is enchants plus the gear they move, on one run each. Treat a point or two as noise.

| Spec | Faction | No enchants | Enchants | Gain | Main hand without, with |
|---|---|---|---|---|---|
| mage-fire | Alliance | 738.5 | 784.8 | +6.3% | Sageblade, Ironbark Staff |
| mage-fire | Horde | 726.5 | 775.8 | +6.8% | Sageblade, Ironbark Staff |
| druid-balance (Natural Power) | Alliance | 660.1 | 722.1 | +9.4% | Sageclaw, Ironbark Staff |
| druid-balance (Natural Power) | Horde | 647.3 | 711.0 | +9.8% | Hand of Edward the Odd, Ironbark Staff |
| warrior-fury | Alliance | 915.3 | 1002.9 | +9.6% | Ironwood Blade, Persuader |
| warrior-fury | Horde | 925.6 | 995.5 | +7.5% | Persuader, Persuader |
| priest-holy (healing per second) | Alliance | 803.4 | 858.7 | +6.9% | Glowing Brightwood Staff, same |
| priest-holy (healing per second) | Horde | 815.5 | 873.5 | +7.1% | Glowing Brightwood Staff, same |

Balance's gloves, same runs: Arcane Power 726.2 / 714.8, Natural Power 722.1 / 711.0, the kit
721.4 / 710.4 (Alliance / Horde). Arcane leads by about 0.5% on both sides, so the plan wears it. The
other 23 specs were not run; the next forced run is their first measurement.
