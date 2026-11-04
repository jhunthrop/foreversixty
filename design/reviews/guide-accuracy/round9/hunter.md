# Guide accuracy round 9: Hunter

Auditor: hunter. Pages: class index, Beast Mastery, Marksmanship and Survival.
- Tree: gr9/audit at 645caa983 (land/guides-r8, all eight rounds applied).
- Client: 1.60.1.70755.
- Engine pin: d1d9f2775 (sim/enginever/version.go).
- Files: `web/src/content/guides/hunter/{index,beast-mastery,marksmanship,survival}.md` and the card notes in `data/curated/apl/hunter-*.json`.
- Read first: the round 1 to round 8 hunter reports with every fix log, round4/hunter-levels.md, the round 2 to round 5 race-sims files and ledger.md.

Since round 8 (`git diff 09b98d9e4 HEAD`), the hunter files changed only by the three round 8 fixes:
- beast-mastery.md:79, Mortal Shots.
- marksmanship.md:71, Mortal Shots.
- survival.md:74 and :102, Predator's Edge.

Builds, card notes, addon-builds.json and the hunter BiS files did not move.

## Builds and leveling order

I decoded all five strings again against `talents/hunter.json`, with talents ordered by tier, then column. All are legal at 51 points, and every prerequisite is met.

| Build | Split | Points |
|---|---|---|
| BM leveling | 31/20/0 | Deadly Aspects 5, Endurance Training 4, Focused Fire 2, Bestial Swiftness 1, Unleashed Fury 5, Ferocity 5, Intimidation 1, Bestial Discipline 2, Frenzy 5, Bestial Wrath 1; Lethal Attacks 5, Improved Stings 2, Efficiency 4, Careful Aim 5, Rapid Killing 1, Trueshot Aura 1, Mortal Shots 2 |
| BM raid | 31/20/0 | As leveling, except Efficiency 3, Rapid Killing 0 and Mortal Shots 4. This matches the band 60 raid BiS talent string `5420001505001251-0052350001400000`. |
| MM leveling | 16/30/5 | Deadly Aspects 5, Endurance Training 3, Focused Fire 2, Unleashed Fury 5, Ferocity 1; Lethal Attacks 5, Improved Stings 1, Efficiency 5, Careful Aim 5, Trueshot Aura 1, Mortal Shots 5, Barrage 3, Ranged Weapon Specialization 5; Improved Tracking 5 |
| MM raid | 5/33/13 | Deadly Aspects 5; Lethal Attacks 5, Improved Stings 3, Efficiency 5, Careful Aim 5, Lone Wolf 1, Trueshot Aura 1, Mortal Shots 5, Barrage 3, Ranged Weapon Specialization 5; Improved Tracking 5, Survivalist 5, Surefooted 3 |
| SV | 5/11/35 | Deadly Aspects 5; Lethal Attacks 5, Careful Aim 5, Lone Wolf 1; Improved Tracking 5, Savage Strikes 2, Survivalist 3, Clever Traps 2, Surefooted 3, Deterrence 1, Predator's Edge 5, Counterattack 1, Resourcefulness 2, Expose Prey 2, Survivalist's Discipline 2, Strider Kick 1, Lightning Reflexes 5, Lacerating Strikes 1 |

I re-decoded the shipped order from `addon-builds.json`, where index 0 is level 10:
- **BM:** Deadly Aspects 10-14, Focused Fire 15-16, Endurance Training 17-20, Bestial Swiftness 21, Unleashed Fury 22-26, Ferocity 27-31, Intimidation 32, Bestial Discipline 33-34, Frenzy 35-39, Bestial Wrath 40, Lethal Attacks 41-45, Improved Stings 46-47, Efficiency 48-51, Careful Aim 52-56, Rapid Killing 57, Trueshot Aura 58, Mortal Shots 59-60.
- **MM:** Lethal Attacks 10-14, Improved Stings 15, Efficiency 16-20, Careful Aim 21-25, Trueshot Aura 26, Mortal Shots 27-31, Barrage 32-34, Ranged Weapon Specialization 35-39, Deadly Aspects 40-44, Endurance Training 45-47, Focused Fire 48-49, Unleashed Fury 50-54, Ferocity 55, Improved Tracking 56-60.
- **SV:** Improved Tracking 10-14, Savage Strikes 15-16, Survivalist 17-19, Clever Traps 20-21, Surefooted 22-24, Deterrence 25, Predator's Edge 26-29, Strider Kick 30, Predator's Edge 31, Counterattack 32, Resourcefulness 33-34, Expose Prey 35-36, Survivalist's Discipline 37-38, Lightning Reflexes 39 and 41-44, Lacerating Strikes 40, Deadly Aspects 45-49, Lethal Attacks 50-54, Careful Aim 55-59, Lone Wolf 60.

Every leveling sentence matches this order:
- Each "first five points" line.
- "Focused Fire at 15 and 16 ... Endurance Training starts at 17".
- "Improved Stings at 15 ... Efficiency from 16".
- "Savage Strikes at 15 and 16 ... Survivalist from 17".
- "Bestial Wrath comes the same level" as mail at 40.
- "Lone Wolf is one of the last points this build takes".

## What was checked and holds

**Round 9 lesson: crit damage.** Each talent that raises the crit damage bonus adds to the +100% part of a crit. 5 points of +30% bonus make a crit 230% of a normal hit.
- **Mortal Shots** (19485): aura 108, misc 15 (crit damage bonus), 30. Rank text: "critical strike damage bonus on all ranged abilities".
  - BM at 2 points: 212%. At 4 points: 224%.
  - MM at 5 points: 230%.
- **Predator's Edge** (1310627): the same aura 108, misc 15, 30. Blizzard's hunter deep dive calls it "Increases melee critical strike damage bonus", which confirms round 8's D2. 230% on melee specials is right.
- "Makes your ranged crits hit harder" (marksmanship.md:116) and "makes your special crits hit harder" (survival.md:116) make no number claim and hold.
- No Eureka! line appears on any hunter page, and no Gnome hunter exists (combos.json).

**Talent values.** I checked every value against the 70755 rank text in talents/hunter.json:
- Unleashed Fury 15%. Ferocity 10%. Frenzy 100%, 30% for 8 sec.
- Focused Fire 2%. Bestial Wrath 50% for 18 sec on 120 s (spellconst 19574). Bestial Discipline 50% regeneration while casting. Bestial Swiftness 30% pet speed.
- Intimidation 3 sec stun, 60 s cooldown.
- Deadly Aspects 2% a rank (10% at 5) for 30% attack speed over 12 sec. Under Hawk it procs from Auto Shot; under Beast, from melee autos.
- Lethal Attacks 5%. Improved Stings 6%, 13% and 20%. Efficiency 15% on shots, stings and melee. Careful Aim 100% of Intellect.
- Rapid Killing rank 1: 1 min off Rapid Fire and 10% (415405 and 415407 at max rank: 2 min and 20%).
- Lone Wolf 20%. Barrage 15%. Ranged Weapon Specialization 5%. Improved Tracking 5% for the seven tracked types.
- Survivalist 2% a point. Savage Strikes 4% melee ability crit. Surefooted 3% hit.
- Predator's Edge: off hand 50% (aura 122), so three quarters of a main-hand swing.
- Resourcefulness 60% plus the crit regeneration. Expose Prey 10% (aura 42, 10) for 10 sec. Survivalist's Discipline 40%. Lightning Reflexes 15%. Lacerating Strikes 40% over 21 sec.
- The Mongoose Bite windows are 5 sec after a dodge and 10 sec from Expose Prey (engine `sim/hunter/mongoose_bite.go`, spells 5302 and 1310726).
- Sniper Shot is the tier 6 capstone with a 4 sec cast, against Aimed Shot's 2 sec.

**Spells and trainer levels.** I checked `trainables/hunter.json` and `spelltips/hunter.json`.

| Level | Spells |
|---|---|
| 1 | Heart of the Lion, Raptor Strike (new ranks at 8 and 16), Track Beasts |
| 4 | Serpent Sting, Aspect of the Monkey |
| 6 | Arcane Shot, Hunter's Mark |
| 8 | Concussive Shot |
| 10 | Aspect of the Hawk, Tame Beast, Track Humanoids |
| 12 | Wing Clip, Mend Pet |
| 16 | Mongoose Bite ("Can only be performed after you dodge") |
| 18 | Multi-Shot, Track Undead |
| 20 | Aimed Shot, Aspect of the Cheetah ("dazed for 4 sec") |
| 26 | Rapid Fire ("ranged and melee attack speed by 40% for 15 sec", 300 s) |
| 30 | Feign Death, Aspect of the Beast, Lacerate |
| 40 | Volley |
| 60 | Aspect of the Falcon, Tranquilizing Shot |

Spell values also hold:
- Multi-Shot hits 3 targets and shares Aimed Shot's 6 s cooldown.
- Disengage is on 5 s and "Disables auto-attack".
- Freezing Trap: "Any damage caused will break the ice".
- Explosive Trap hits all within 10 yards.
- Volley is a 6 sec channel.
- Revive Pet costs 100% of base mana.
- Trueshot Aura rank 5 is 100. Hawk rank 6 is 110. Beast rank 4 is 110.
- Falcon matches "their highest rank of Aspect of the Hawk".
- Aspect of the Viper (415423) is a non-active row, and no page names it.

**Heart of the Lion.** 409583 is aura 99 and aura 124 (melee and ranged attack power) at 40, plus 4 a level to max level 60. The engine scales it from BaseLevel 0 (`sim/core/buff_ranks_test.go` TestHeartOfTheLionAttackPowerFollowsClientRule), which gives 280 at 60. The index's "280 melee and ranged at 60" holds.

**Pets.** Warcraft Tavern's Forever pets guide says pets receive "10% of the Hunter's total Attack Power and 100% of their Critical Strike Chance". Its hunter guide adds "whichever is highest" of melee or ranged. Blizzard: "Pets can no longer receive player buffs that increase their stats." The index, BM and MM pet lines hold. The taming chain "teaches Tame Beast" (Warcraft Tavern: "Once you have completed all three quests, you will learn Tame Beast").

**Races.** Every racial named matches races.json:
- Elune's Light: 10% crit, 15 sec, 3 min.
- Ruthless: 1% crit.
- Artisan's Precision: 1% hit with spells and attacks.
- Big Game Hunter: 5% to beasts.
- Blood Fury: 10% attack power, ranged attack power and spell power, 15 sec, 2 min.
- Command: 5% pet damage.
- Berserking: 10% casting and attack speed, 15 sec, 3 min.
- Wind Blessed: 2% haste.

The orders follow the guide-build race sims in round2/race-sims.md, seeds 1 and 50001:
- BM: Night Elf > Human > Dwarf, and Orc > Troll.
- MM: Night Elf > Human > Dwarf, then Troll with Orc close behind.
- SV: Night Elf > Human > Dwarf, with Orc and Troll level (818.80|818.27 against 818.59|818.99).

Skyborne is named only as the paid option.

**Survival's "two fewer points of Deadly Aspects" line.** The integrator measured it in round 2: 3/005005201 came in 1.9% under the published build, beyond error. It holds.

**"What the raid adds" (survival.md:65).** The 70755 Strength of Earth Totem rank 5 is 53 Strength (Era 77). Curse of Recklessness rank 4 is 505 armor (Era 640). Both clauses hold. Battle Shout (139) and Blessing of Might (133) are also below Era. Grace of Air rank 3 (89 Agility) is above Era, but it does not turn the net over (see SHOULDs).

**Enchants.** I re-scored `enchants.json` with the band 60 raid Alliance weights, using the same rule as gear-extras.ts pickEnchants:
- BM and MM: Greater Agility on gloves and boots, Lesser Agility on the cloak, Greater Stats on the chest, Greater Intellect on bracers. No bracer enchant has Agility beyond Minor.
- SV: Superior Strength on bracers and Agility on one-hand weapons.

All of it matches the enchant prose on all three pages.

**Consumables, ammo and quivers.** These are unchanged since round 7 and still match the data:
- Major Mana Potion 49, Mongoose 46, Salmon 35, Major Frenzy Potion 55.
- Jagged Arrow 40, Ice Threaded 51, Thorium Headed and Thorium Shells 52 (Schematic: Thorium Shells exists).
- Quickdraw Quiver 40, Harpy Hide 55, Lamina 60.

**Stat paragraphs.** None of them ranks stats, and each reason is true on 70755:
- The hit cap is 9%.
- SV `hit_to_cap` specials is 5 from a baseline of 4, which is Surefooted 3 plus Dwarf 1.
- The white cap is 24.
- `haste_on_items` is false.

The pills render from data (`rankedStatLabels`), so the frontmatter order is only the set of stats.

**Card notes.** I re-read all BM, MM and SV notes. Nothing in them is false. The Survival priority prose matches the APL order.

**Voice.** No em or en dashes. No sim, site or pipeline words. No quoted nightly numbers. No removed spells or SoD runes. No ranking words in Stat priority.

## Source links

I fetched every URL this round with a browser user agent.
- 200: the three Blizzard articles (Hunter and Druid deep dive, Deep Dive panel recap, What's Next recap).
- 200: Warcraft Tavern's Forever Hunter Guide, Leveling Guide and Hunter and Druid Deep Dive.
- 200: foreversixty.gg combos.json, races.json, spellconst/hunter.json, /planner, /planner?class=hunter and /sim/weights.
- 403: Wowhead's Skyborne First Look, spell=1495 and spell=19306. These count as unknown, not dead. The nether tooltip endpoint returns "Mongoose Bite" for 1495 and "Counterattack" (Talent, level 30) for 19306, so both labels match.

Every entry follows the source-link rule:
- No /sources, wago or Talents Forever entries.
- The talent trees open the planner.
- Stat weights appear only on the damage pages, which is all of them for hunter.
- The body links go to /bis/hunter/<spec>.

## Findings

None. No statement on the four hunter pages or in the hunter card notes is false on 70755 or contradicts the page. The round 8 edits hold.

### SHOULDs (wording, not findings)

- **survival.md:65.** "A full raid lifts a Survival hunter less than it did in Era" is true net. Forever's Grace of Air (89 Agility at rank 3) is above Era's, though, and the sentence names only the buffs that went down. "Most of the melee buffs are smaller than Era's" would be safer.
- **Carried from round 8.**
  - The Serpent Sting card says "Refreshing any earlier throws away a tick"; "Refresh it as it drops" would be safer.
  - "Lucifron's tome" could read "the Molten Core tome".
  - beast-mastery.md:81 names Deflection, a Survival talent, in a list of what a 31/20/0 build skips. It is true but reads oddly.

## Decisions

- **D1. Predator's Edge stays "230% of a normal hit".** Blizzard's deep dive says it "Increases melee critical strike damage bonus". That settles the gap between the tooltip ("critical strike damage by 30%") and the spell data in favour of round 8's reading.
- **D2. "Orc and Troll are level" (survival.md:118) stays.** The guide-build sims are inside error at both seeds, so `recommendedRaces: [night-elf, orc]` is a tiebreak, not a contradiction.
- **D3.** Round 8's D1 to D3 and round 7's D1 to D4 stand.

## Sims requested

Nothing new. Carried, and none of it is claimed in prose:
1. **Survival weapon enchant.** Crusader on both weapons against Agility on both. Guide build `5/005005001/500230231051222151`, Night Elf, band 60 raid, seeds 1|50001.
2. **Hunter race re-runs after the final best-in-slot re-rank.** MM Troll against Orc, SV Orc against Troll, BM Night Elf against Human, and BM Orc against Troll.
3. **Blocked on the engine.** Sniper Shot needs a rotation line, and Summon Hawk is unmodeled.

## Sim findings

Nothing new. Carried:
- No Sniper Shot line in the Marksmanship rotation.
- Summon Hawk is unmodeled.
- Pet non-stat buffs are blocked by RefusePlayerBuffs.
- The APL top-level notes still call Falcon 55/55 (notes only, not rendered).

## Sources

**Forever Sixty client data, 1.60.1.70755**
- `talents/hunter.json`: rank text for all three trees.
- `trainables/hunter.json` and `spelltips/hunter.json`: every spell above, plus 409580 Heart of the Lion, 469145 Falcon and 415423 Viper.
- `spellconst/hunter.json`: 19485, 1310627, 409583, 19574, 3045, 1310532, 1310533, 415405, 415407.
- `races.json`, `combos.json`, `enchants.json`, `items.json`, `addon-builds.json`, and `bis/hunter-*.json` (band 60 raid weights, `hit_to_cap`, `haste_on_items`).
- `spelltips/shaman.json` (Strength of Earth, Grace of Air), `spelltips/warlock.json` (Curse of Recklessness), `spelltips/paladin.json` (Blessing of Might, Seal of Wisdom), `spelltips/warrior.json` (Battle Shout).

**Engine d1d9f2775**
- `sim/core/buff_ranks.go:135-141` and `buff_ranks_test.go:171-180`: Heart of the Lion scaling.
- `sim/hunter/mongoose_bite.go:45,84-119`: the Mongoose Bite windows.

**Sims and the site**
- Race and build sims: design/reviews/guide-accuracy/round2/race-sims.md and round2/hunter.md ("Integrator: S-R2-1 measured").
- Enchant pick rule: web/src/lib/guides/gear-extras.ts pickEnchants.
- Stat pills: web/src/pages/guides/[class]/[spec].astro (rankedStatLabels).

**Web**
- Blizzard, World of Warcraft: Forever Class Deep Dives: Hunter and Druid, https://news.blizzard.com/en-us/article/24301515/world-of-warcraft-forever-class-deep-dives-hunter-and-druid. Predator's Edge "Increases melee critical strike damage bonus"; pets "can no longer receive player buffs that increase their stats"; the 8 yard dead zone.
- Warcraft Tavern, Hunter Pets Guide (Forever), https://www.warcrafttavern.com/forever/guides/hunter-pets/. Pets take 10% of attack power and 100% of crit chance; Tame Beast after the three taming quests.
- Warcraft Tavern, Hunter Class Guide (Forever), https://warcrafttavern.com/forever/guides/hunter/. The pet takes 10% of melee or ranged attack power, whichever is highest.
- Warcraft Tavern, Hunter Leveling Guide (Forever), https://www.warcrafttavern.com/forever/guides/hunter-leveling/.
- Wowhead Forever tooltips: https://nether.wowhead.com/forever/tooltip/spell/1495 and https://nether.wowhead.com/forever/tooltip/spell/19306.
