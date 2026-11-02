# Sim audit round 10: auras lane

Scope: every raid buff, debuff, consumable, party aura and racial-free outside effect in the raid and bare presets of every spec at bands 20/30/40/50/60 (data/curated/presets.json, sim/request, the engine's core buffs, debuffs and consumables).

Sources:
- The 1.60.1.70707 client tables with the hotfix push layered in (every table's sha256 matches `data/builds/1.60.1.70707/hotfix-push.json`). Read with the lane's `tools/sp.py` and `tools/it.py`.
- Wowhead Forever server tooltips (`nether.wowhead.com/forever/tooltip/...`) and item pages, read 2026-11-02. A tooltip and its client row count as one source.
- Blizzard's Forever druid deep dive, news.blizzard.com article 24301515 (forum copy us.forums.blizzard.com/t/2367361, 2026-09-30).
- forever-bugs issues #20, #109 and #419 (read 2026-11-02).
- wow-professions.com's page for Nina Surefire, the Azeroth Commerce Authority alchemy vendor (read 2026-11-02).

Engine branch `lane/r10-auras` (from 4d3129cae). Site branch `wt/r10-auras`.

**Found: 2 CONFIRMED-BUG rows (P1, F1), both fixed, plus one note fix (N1).** Both move published numbers at the level-60 raid preset:
- P1: Forever's 30-second potions were unmodelled. Ten specs now drink one: the three rogues and the cat had no potion, and six others drink theirs instead of a mana potion.
- F1: the engine kept Era's shapeshift rules. A feral druid got no Windfury Totem, no weapon stone, no potion and no Thistle Tea. Forever allows all four.

## Consumables

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| P1 | **Forever's 30-second potions were missing, so ten specs drank a weaker potion or none.**<br>- The recipes are launch content: Wowhead Forever says each "can be purchased in Redridge Mountains (2) and The Barrens (2)". wow-professions.com lists the Azeroth Commerce Authority alchemy vendor Nina Surefire's stock. Frenzy Potion and Spellblasting Potion need no reputation. Major Frenzy Potion and Major Spellblasting Potion need Honored. This settles round 9's open C5 for the potions.<br>- The client amounts:<br>&nbsp;&nbsp;- Frenzy 250940: 28 attack power and 28 ranged attack power (spell 1251936, aura 99 and aura 124).<br>&nbsp;&nbsp;- Major Frenzy 250943: 80 (1251940).<br>&nbsp;&nbsp;- Spellblasting 250934: 16 spell damage (1251923, aura 13, mask 126).<br>&nbsp;&nbsp;- Major Spellblasting 250937: 47 (1251932).<br>&nbsp;&nbsp;- Each lasts 30 s in SpellCategories category 4 with a 120000 ms SpellCooldowns row. That is the potion timer a mana or rage potion uses.<br>- The engine now has these four potions. Each spec carries the potion that measured best:<br>&nbsp;&nbsp;- The rogues drank no potion. Major Frenzy is +1.0 to +1.2%, with Frenzy from level 25.<br>&nbsp;&nbsp;- Major Frenzy beats Major Mana for Survival (+0.9%) and Enhancement (+0.8 to +0.9%).<br>&nbsp;&nbsp;- Major Spellblasting beats Major Mana for Balance (+0.8 to +1.2%), Frost (+1.4 to +1.5%), Shadow (+0.9 to +1.0%) and the paladin tank (+0.8 to +0.9% threat).<br>- These specs keep their old potion, which measured ahead:<br>&nbsp;&nbsp;- Beast Mastery and Marksmanship (-0.1 to -0.2% with Frenzy).<br>&nbsp;&nbsp;- Retribution (-0.8 to -0.9%).<br>&nbsp;&nbsp;- Arcane and Fire (-0.2 to -0.7% with Spellblasting).<br>&nbsp;&nbsp;- Elemental (-4.8%, out of mana).<br>&nbsp;&nbsp;- The warlocks (-0.7 to -1.2%).<br>&nbsp;&nbsp;- Arms, Fury and the warrior tank (-0.6 to -1.9% against Mighty Rage). | Client rows above; tooltips "Increases Attack Power by 80 for 30 sec.", "Increases Spell Damage by 47 for 30 sec." (items 250943, 250937). Swap measurements: band 60 raid, horde, 3,000 iterations, seeds 1 and 50001 (`swap.py`). Tests first: `TestBurstPotionsMatchTheClient`, `TestBurstPotionIsADPSCooldownOnThePotionTimer` (engine, failed to build before the enum existed), and new cases in `TestEachConsumableSlotCarriesTheFirstUsableEntry` (site, 8 failures before the preset change). | CONFIRMED-BUG (two sources for the vendor; the reputation is single-source) | E 61734a1cd, S 3df48dfd6 |
| F1 | **A feral druid's consumables and weapon buffs followed Era's shapeshift rules.**<br>- Blizzard's druid deep dive: "All potions and almost all similar consumables are now useable while shapeshifted and do not cancel the Druid's shapeshift. Rage potions and Thistletea are also now usable by Druids. One notable exception is Engineering bombs." It also says: "Weapon enchantments and procs now all work while shapeshifted, including effects like weightstones, Crusader enchant, and Windfury Totem."<br>- forever-bugs #109: Blizzard fixed the last case, a mana potion refused in Cat Form, on 2026-10-08: "We can drink mana and health potions in forms now".<br>- What the engine did:<br>&nbsp;&nbsp;- Every potion, rune and healthstone cancelled the form. The rage potion waited for the form to end, so a cat or bear never drank one in combat.<br>&nbsp;&nbsp;- Thistle Tea lived in the rogue package.<br>&nbsp;&nbsp;- Windfury Totem and Flametongue Totem skipped a feral druid (`core/buffs.go` weaponTotemReaches).<br>&nbsp;&nbsp;- An Elemental Sharpening Stone gave a feral druid nothing. The site resolver also stripped every imbue from the cat (`NoWeaponImbueSpecs`).<br>- Now potions and conjured items keep the form; the explosives still cancel it. Thistle Tea is a core consumable for any energy bar. The totems, the stone and the oils reach a feral druid.<br>- The preset gives the cat the melee group's stone, Major Frenzy Potion (Frenzy from 25) and Thistle Tea. Measured on top of the engine fix: the stone +1.8%, the potion +0.8%, the tea +1.5%.<br>- Windfury Totem was already in the preset and now reaches the cat (+8.4%) and the bear (+11% threat). | Blizzard article 24301515 (two quotes above); forever-bugs #109 (Blizzard reply). Tests first: `TestACatUsesItsConsumablesWithoutLeavingCatForm` (feral: no potion, no tea and no Windfury proc before), `TestABearDrinksARagePotionInForm` (tank, `-tags=with_db`), `TestWeaponConsumablesReachAShapeshiftedDruid` and `TestPotionsDoNotBreakAShapeshift` (core), and `TestRaidPresetGivesACatItsWeaponStoneAndThistleTea` (site; it replaces the test that pinned "a cat's claws take no stone"). The two totem-reach tables now expect a feral druid to get the totem. | CONFIRMED-BUG (Blizzard article plus a Blizzard bug reply) | E d5c560d31, S c07e320cd |
| N1 | Juju Power's preset note said "Elixir of Giants is not in the 1.60 client's item table". It is there: item 9206, renamed Elixir of Greater Strength, 25 Strength (spell 11405). Juju Power (30, spell 16323) is still the stronger, so no number moves. | ItemSparse 9206; SpellEffect 11405, 16323 | note fix | S 3df48dfd6 |
| C1 | **Forever's new elixirs** (Grizzly, Ferocity, Cunning, Owl, Whale, Nature Power) are sold by the same vendors at Honored, so they are launch content too. They stay out of the preset, because whether they stack with Mongoose, Greater Arcane and Juju Power is unknown.<br>- In Classic no Battle or Guardian category exists. Every elixir shares only the 3 s category 2593 cooldown.<br>- The preset already carries three of the same vendor's Honored and Friendly recipes for healers and tanks: Elixir of Sages, Greater Mageblood and Greater Cleric's. They have no rival in their slot, so they raise no stacking question. | wow-professions.com (Nina Surefire's 28 recipes and their reputations); Wowhead Forever 250387, 250377 ("Redridge Mountains (2) and The Barrens (2)"); ItemEffect category 2593 on every elixir | OPEN: a beta character sheet with Mongoose plus Grizzly settles the stacking | - |
| C2 | Every other preset consumable was re-read against the client:<br>- **Flasks:** Supreme Power 150 spell damage (aura 13, mask 126), Distilled Wisdom 2000 mana, the Titans 1200 health.<br>- **Elixirs and Zanza:** Greater Arcane 35. Mongoose 25 Agility with 2% crit (aura 290, every table: `stats.Crit`). Juju Power 30, Juju Might 40 and 40 ranged, R.O.I.D.S., Assay, Cortex and Gizzard Gum each 25 (category 103). Greater Defense 450 (aura 674), Fortitude 200, Greater Stoneshield 2000.<br>- **Healer elixirs:** Greater Mageblood 20, Greater Cleric's 40, Elixir of Sages 25 Spirit with 2% crit (aura 290).<br>- **Foods:** Baked Salmon 28 (Well Fed 1249520), Poached Sunscale Salmon 40 and 40 (1249519), Smoked Desert Dumplings 20 Strength (1248422) from level 45.<br>- **Potions and runes:** Major Mana 1350 to 2250 with a 2-minute cooldown (tooltip), Mighty Rage 45 to 75 rage and 60 Strength for 20 s, Demonic Rune and Thistle Tea in category 1153, apart from the potion timer.<br>- **Stones:** Elemental Sharpening Stone aura 52 (melee table), Brilliant Wizard Oil aura 57 (spell table). | `tools/it.py` over the layered ItemSparse, ItemEffect, SpellEffect, SpellCategories and SpellCooldowns; `core/consumes.go` | Aligned | - |

## Raid buffs

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| B1 | **Every `BuffRanks` table re-checked from scratch** against the layered SpellEffect and SpellLevels. The script matches each row's spell, learn level and amount:<br>- Arcane Intellect 31 at 56. Mark of the Wild 385 armor, 16 stats, 27 resistances at 60. Power Word: Fortitude 70. Divine Spirit 40 (from 30).<br>- Blessing of Might 112 at 52 and Blessing of Wisdom 36 at 54. The level-60 ranks 25291 and 25290 are the Ahn'Qiraj books (`IncludeAQ = false`).<br>- Battle Shout 111 + 0.6 a level to 61 (115 at 60, aura 99 melee only). Trueshot Aura 50 (aura 124 ranged only). Heart of the Lion 40 + 4 a level (280 at 60, tooltip "Attack Power by 280").<br>- Devotion Aura 735.<br>- The totems, on their aura spells: Strength of Earth 42 (10441), Grace of Air 77 (10626), Mana Spring 10 every 2 s (10494). Windfury Totem's attack buff is 95/179/246 (8516/10608/10610). Its proc auras (8515/10609/10612) carry ProcChance 20.<br>- The reactive auras by rank: Thorns, Retribution Aura, Stoneskin, Mana Tide. | `tools/ranks.py` over `engine/sim/core`; the "mismatches" it prints are the totems (stated on their aura spells) and the Judgements (stated on their triggered spells), each checked by hand | Aligned | - |
| B2 | **Which crit table each crit buff feeds after round 9's split.**<br>- Leader of the Pack (24932) and Moonkin Aura (24907) are both aura 290. Forever's tooltips read "Increases Critical Strike chance by 3%" for both, so each reaches every table: melee, ranged and spell. That is a Forever change from Era's melee-only and spell-only auras.<br>- The engine adds `stats.Crit` for both, from the talent tiers' levels 30 and 40. | SpellEffect 24932, 24907; Wowhead Forever 17007 ("increases the Critical Strike chance of all party or raid members"), 24858, 24907 | Aligned (client plus tooltip) | - |
| B3 | **World buffs are refused** by the request layer (`isWorldBuff`), and the owner's ruling says Forever has none. Their engine code is unreachable from a request:<br>- Rallying Cry gives `Crit` 10 (Forever's tooltip: "critical chance of spells and attacks by 10%"). The engine lacks its "+80 spell damage and healing".<br>- Songflower gives `Crit` 5 ("melee, ranged, or spell critical by 5%"). Slip'kik's gives 3. Spirit of Zandalar gives 15% to every stat. | Wowhead Forever 22888, 15366, 22820, 24425; `sim/request/vocabulary.go` isWorldBuff | Aligned (unreachable; Rallying Cry's spell damage noted) | - |
| B4 | **Blessing of Kings with Heart of the Lion.** Both are aura 137 at 10 with misc -1. Neither Forever tooltip says the two do not stack, and the client has no exclusivity row. Unchanged from round 7. | SpellEffect 20217, 25898, 409580, 409583; Wowhead Forever tooltips of all four (read 2026-11-02) | OPEN: a beta character sheet with both | - |
| B5 | **Salvation and Tranquil Air** reduce threat only. The engine does not cap a damage dealer's threat, so neither changes a number. The tanks carry no Salvation. | preset notes; engine has no threat cap | Aligned | - |
| B6 | Round 9's Winter's Chill fix has no side effect. Frost's raid run shows its own talent stacks only, and the raid flag stays inert. | spec-breakdown mage-frost aura table; `core/debuffs.go` | Aligned | - |

## Raid debuffs

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| D1 | **Armor debuffs and aura 674 (round 9 open), settled from the client.**<br>- Forever moved every flat armor effect from Era's aura 22 to aura 674. That covers 283 spells, positive and negative:<br>&nbsp;&nbsp;- Sunder, Expose Armor, Faerie Fire and Curse of Recklessness;<br>&nbsp;&nbsp;- Mark of the Wild, Devotion Aura, Inner Fire, Demon Armor and Ice Armor;<br>&nbsp;&nbsp;- the item "Increased Armor N" equip spells, the 200-armor set bonuses and the pet scaling auras.<br>- Its test spell is named "(DNT) Mod Armor No Mods" (1322297). Wowhead Classic shows Era's Sunder as "Mod Resistance (Physical)" (aura 22).<br>- So 674 is "flat armor not scaled by percentage modifiers", not a non-stacking aura. If 674 kept only the strongest source, a player's item armor bonuses, Inner Fire and Mark of the Wild could not all count, and the client gives them all to one character.<br>- forever-bugs #20 agrees. Bear Form's multiplier was being applied to Mark of the Wild and Elixir of Minor Defense, and Blizzard fixed that (2026-10-14).<br>- The engine already models all of this:<br>&nbsp;&nbsp;- Buff armor is `BonusArmor`, which no form multiplies.<br>&nbsp;&nbsp;- Sunder and Expose share the major category, so the larger one counts, as in Era.<br>&nbsp;&nbsp;- Faerie Fire is in the minor category, and Curse of Recklessness adds on its own.<br>&nbsp;&nbsp;- Raid total at 60: 2250 + 505 + 505. | SpellEffect (283 rows with EffectAura 674), SpellName 1322297; https://www.wowhead.com/classic/spell=11597; forever-bugs #20; `core/debuffs.go` Sunder, Expose, minor armor reduction | Aligned (client rows plus a Blizzard bug reply). Exclusivity between spells is server-side; no source says Forever changed Era's. | - |
| D2 | Curse of the Elements rank 1311680 at 50: aura 87 +10% and aura 22 -75 over mask 126, which includes Holy. The engine scales all six magic schools, so Retribution's and the paladin tank's Holy damage too. | SpellEffect 1311680 | Aligned | - |
| D3 | Judgement of the Crusader is 161 Holy damage taken (aura 14) at 52 for 40 s. Judgement of Wisdom and Judgement of Light are unchanged. Their debuffs (20355, 20346) trigger script spells 1826 and 5373 at ProcChance 100, so the 50% chance stays the engine's constant. | SpellEffect, SpellAuraOptions 20355/20346/1826/5373 | Aligned (the 50% is single-source, unchanged) | - |
| D4 | Shadow Weaving (15258), Fire Vulnerability (22959) and Shadow Vulnerability (17794) are aura 270, "school damage from caster". Winter's Chill (12579) and Stormstrike (17364) are aura 271 on the caster's own spells. All five stay out of the raid flags, as round 8 left them. | SpellEffect rows; TrinityCore SpellAuraDefines.h (270 MOD_SCHOOL_MASK_DAMAGE_FROM_CASTER, 271 MOD_SPELL_DAMAGE_FROM_CASTER) | Aligned | - |

## Round 9 fixes, side effects

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| R1 | Survival's main-hand Elemental Sharpening Stone (round 9 S1) now reaches its pet's melee too. The pet copies the hunter's `MeleeCrit` (`hunter/pet.go`), and only the hunter's ranged shots take the 2 back off. The crit lane noted this as open while "no published hunter preset carries a stone". Round 9 S1 put one on Survival in the same round, so it now moves a published number, by about 2% pet crit. Hunter Pet Scaling 415429 carries aura 52 and 57 channels at bp 0 (server-scripted), so the client does not say whether a weapon enchant on the owner counts. | SpellEffect 415429; `hunter/pet.go:215` | OPEN (crit lane; the open item is now live) | - |
| R2 | The other round 9 fixes have no side effect:<br>- The Dumplings' level gate holds at 45.<br>- The Blasted Lands slot holds one buff per spec. `ZanzaBuff` is one field.<br>- Thorns, Retribution Aura, Stoneskin and Mana Tide are custom-panel only.<br>- Orc Command reaches pets only. | `TestEachConsumableSlotCarriesTheFirstUsableEntry`, `TestItemLevelsReadTheClientsRequiredLevel`; engine core suite | Aligned | - |

## Concentration Aura and forever-bugs #419

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| A1 | Concentration Aura (19746) is unchanged: aura 149 at 35, all schools, trainable at 22. Round 9 settled the mechanic: mangos sums aura 149 with the spell's own "keep casting time" chance, so Spiritual Focus plus the aura means Holy Light is never pushed back. The open question is whether a raid's Holy paladin runs it, which is a kit choice. About +0.4% for the Holy paladin. | SpellEffect 697291; round 9 healers/aud-paladin.md P2 | OPEN (owner: kit choice; the mechanic is settled) | - |
| A2 | forever-bugs #419 (shadow wands ignore Shadow Weaving, Darkness and Shadowform): still open with no Blizzard reply (updated 2026-10-21). The client does not explain it. Shadowform 15473 is aura 79, a plain +10% to Shadow damage done, which would reach a wand. So the behaviour is server-side or a bug. No published rotation shoots a shadow wand. | gh issue 419 (read 2026-11-02); SpellEffect 15473 | OPEN (no published number) | - |

## Before/after

Band 60, raid preset, 10,000 iterations. "Before" is `$R/sb-before` (fork 4d3129cae with site origin/main bb619a188). "After" is lane engine d5c560d31 with site c07e320cd. Each cell is the change in DPS, effective HPS for healers, or threat per second for tanks.

| spec | alliance s1 / s50001 | horde s1 / s50001 | horde s1 absolute |
|---|---|---|---|
| druid-balance | +0.91 / +0.78 | +1.13 / +1.15 | 565.2 -> 571.6 |
| druid-feral | +13.02 / +12.87 | +12.98 / +12.97 | 593.0 -> 670.0 |
| druid-feral-bear | +11.59 / +11.70 | +11.55 / +11.54 | 836.3 -> 932.9 |
| druid-restoration | +0.00 / +0.00 | +0.00 / +0.00 | 779.2 -> 779.2 |
| hunter-beast-mastery | +0.00 / +0.00 | +0.00 / +0.00 | 633.4 -> 633.4 |
| hunter-marksmanship | +0.00 / +0.00 | +0.00 / +0.00 | 623.0 -> 623.0 |
| hunter-survival | +0.88 / +0.88 | +0.90 / +0.90 | 734.9 -> 741.5 |
| mage-arcane | +0.00 / +0.00 | +0.00 / +0.00 | 600.2 -> 600.2 |
| mage-fire | +0.00 / +0.00 | +0.00 / +0.00 | 660.8 -> 660.8 |
| mage-frost | +1.43 / +1.43 | +1.45 / +1.52 | 594.1 -> 602.7 |
| paladin-holy | +0.00 / +0.00 | +0.00 / +0.00 | 641.4 -> 641.4 |
| paladin-protection | +0.86 / +0.89 | +0.88 / +0.83 | 993.8 -> 1002.5 |
| paladin-retribution | +0.00 / +0.00 | +0.00 / +0.00 | 621.3 -> 621.3 |
| priest-discipline | +0.00 / +0.00 | +0.00 / +0.00 | 624.9 -> 624.9 |
| priest-holy | +0.00 / +0.00 | +0.00 / +0.00 | 693.4 -> 693.4 |
| priest-shadow | +0.99 / +0.99 | +1.01 / +0.88 | 712.2 -> 719.4 |
| rogue-assassination | +1.01 / +1.01 | +0.99 / +0.98 | 613.6 -> 619.7 |
| rogue-combat | +1.06 / +1.05 | +1.06 / +1.06 | 611.8 -> 618.3 |
| rogue-subtlety | +1.15 / +1.15 | +1.15 / +1.15 | 538.3 -> 544.5 |
| shaman-elemental | +0.00 / +0.00 | +0.00 / +0.00 | 466.9 -> 466.9 |
| shaman-enhancement | +0.83 / +0.83 | +0.87 / +0.87 | 575.8 -> 580.8 |
| shaman-restoration | +0.00 / +0.00 | +0.00 / +0.00 | 838.0 -> 838.0 |
| warlock-affliction | +0.00 / +0.00 | +0.00 / +0.00 | 573.2 -> 573.2 |
| warlock-demonology | +0.00 / +0.00 | +0.00 / +0.00 | 642.6 -> 642.6 |
| warlock-destruction | +0.00 / +0.00 | +0.00 / +0.00 | 638.9 -> 638.9 |
| warrior-arms | +0.00 / +0.00 | +0.00 / +0.00 | 720.8 -> 720.8 |
| warrior-fury | +0.00 / +0.00 | +0.00 / +0.00 | 757.8 -> 757.8 |
| warrior-protection | +0.00 / +0.00 | +0.00 / +0.00 | 1051.4 -> 1051.4 |

What moved:
- **Feral cat and bear:** Windfury Totem now reaches them (F1). The cat also gains its stone, Major Frenzy Potion and Thistle Tea.
- **Rogues:** Major Frenzy Potion, which had no rival.
- **Survival, Enhancement, Balance, Frost, Shadow and the paladin tank:** the 30-second potion in place of Major Mana Potion.
- Every other spec is unchanged.

**Bare bands:** no change. The bare kit carries no potion, totem or stone. Spot check at bands 20 and 40 (horde, seed 1, 3,000 iterations), before and after identical: Feral 45.9 and 114.9, Combat 40.6 and 109.2, bear 56.4 and 170.7 threat.

**Raid-buffed /sim below 60:** rogues and the cat drink Frenzy Potion from 25. The other potion switches start at 55; below that, the mana potion keeps the slot.

## For other lanes

- **Integrator:**
  - S 3df48dfd6 and S c07e320cd need E 61734a1cd (the four new Potions values 23 to 26) pinned with them. Without it, the site's vocabulary refuses the new ids.
  - `data/proto/common.proto` and `common_pb2.py` need those values at the pin.
  - `sim/request/IDS.md` is regenerated against the lane engine.
  - No fork golden moved: the core, druid, rogue and full `-tags=with_db` suites pass. `sim/web` fails setup at the pin too.
  - No `TestRotationLadder` golden was regenerated.
  - Specs outside this lane that move: the three rogues, Survival, Enhancement, Balance, Frost, Shadow, the paladin tank, the cat and the bear (table above).
- **Hybrids lane (feral):**
  - The bear's rotation (`druid-feral-bear.json`, fork `forever_feral_bear.apl.json`) has no autocast line. So the preset's Greater Stoneshield Potion is never drunk, and with F1 a bear could drink a Mighty Rage Potion (rage plus 60 Strength) in form. The potion choice for the bear needs the autocast line first.
  - The flat damage stones and weightstones still skip a feral druid. The paw weapon is rebuilt from the item weapon, so a +8 stone would need the paw model to carry it. The deep dive says weightstones work in a form.
- **Crit lane:** R1, Survival's stone reaching the pet, is now live. F1 removed the feral gate on the Elemental Sharpening Stone (`core/consumes.go`); the table it writes (MeleeCrit) is unchanged.
- **Platform lane:**
  - The custom /sim panel and simbuffs.json (held) need names and icons for the four new potions (items 250940, 250943, 250934, 250937) and for Thistle Tea on a druid.
  - The coach catalogue (`logs/engine/mechanics/consumables/catalogue.json`) still names the old potion picks.
- **Ranker lane:** the alchemy recipes come from the Azeroth Commerce Authority and Durotar Supply and Logistics vendors, which `loot.json` does not list as sources.

## Still open

1. C1: whether Forever's new elixirs stack with Mongoose, Greater Arcane and Juju Power. They are launch content at Honored.
2. B4: Blessing of Kings with Heart of the Lion (round 7).
3. R1: a hunter's weapon stone reaching its pet.
4. A1: Concentration Aura in the Holy paladin's kit (owner).
5. A2: forever-bugs #419.
6. D3: the 50% Judgement of Wisdom and Light chance is server-side and single-source.

## Follow-up: the tanks' weapon imbue (integrator request)

The hybrids lane found the bear best with Mighty Rage Potion and an Elemental Sharpening Stone. Neither the warrior tank nor the paladin tank carried any main-hand imbue. Measurement setup:
- Fork: hybrids 6f3cb68fa merged with the crit lane's E 56a5e5134, which makes the stone a per-spell bonus on the owner's own melee attacks. The conflict was resolved locally for measurement only, without the feral gate, as F1 requires.
- Site: hybrids f3728d435.
- Band 60 raid, threat per second, 10,000 iterations, both factions, seeds 1 and 50001.
- The published main hand for both tanks is Greenhammer 279261, a mace (Item SubclassID 4), so the dense stone is a weightstone.

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| T1 | **The warrior and paladin tanks carried no weapon imbue.** Threat with each imbue against none:<br>- Protection warrior: Elemental Sharpening Stone +1.48 to +1.67%, Dense Weightstone +0.47%. It now carries the Elemental stone.<br>- Protection paladin: Brilliant Wizard Oil +2.32 to +2.40%, Elemental stone +1.29 to +1.34%, Dense Weightstone +0.97 to +0.99%. Its threat is mostly Holy spell damage, so it now carries the oil.<br>- Bear, checked with the crit lane's stone: Elemental +2.05 to +2.19%, either dense stone +0.98 to +0.99%. That confirms the hybrids lane's pick. | `imbue.py` runs above; enchant 2506 (aura 52), enchant 2628 (aura 13 and 135 at 36, aura 57 at 1); Item 279261 | CONFIRMED-BUG (preset) | S f8ad6d401 |

Absolute values, horde seed 1, none to the adopted imbue: warrior 1051.4 -> 1067.5, paladin 994.6 -> 1017.8, bear 971.1 -> 992.4 (with the hybrids lane's bear changes).

For the integrator:
- E 56a5e5134 conflicts with E d5c560d31 in `core/consumes.go`. Keep the crit lane's `addMeleeWeaponCrit` for the Elemental stone and drop its `FeralCombatEnabled` gate, which F1 removed. Keep both new functions, `makeThistleTeaMCD` and `addMeleeWeaponCrit`.
- `TestProtection.results` conflicts too and needs regenerating at the pin.
- `sim/request/raid-consumes.json` will conflict with the hybrids lane's copy; regenerate it after merging.
