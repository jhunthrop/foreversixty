# Guide accuracy round 12: Hunter

Auditor: hunter. Pages: class index, Beast Mastery, Marksmanship and Survival, plus the card notes in `data/curated/apl/hunter-*.json`.
- Tree: gr12/audit at a8b147a6f (main after guide round 11 and guide-header-build).
- Client: 1.60.1.70755.
- Engine pin: 9f7f1f6e4 (sim/enginever/version.go). Engine source read from gr10/engine.
- Read first: the round 1 to round 11 hunter reports with every fix log, round4/hunter-levels.md, the race-sims files and ledger.md through the round 11 section.

## What moved since round 11

`git log d54de48aa..HEAD` touches no hunter guide, no hunter APL, and none of `addon-builds.json`, `talents/hunter.json`, `trainables/hunter.json`, `ladder-order.json` or the hunter BiS files. The hunter BiS files are still de044125c (05:46 UTC), so the final re-rank has not landed in this tree.

The one change that reaches the hunter pages is 3bd24d3a2 (guide-header-build). Beast Mastery and Marksmanship both have a `raidBuild`, so on those two pages:
- The rail's build card now shows and loads the raid build, titled "Raid build".
- Key-talent chips show "Level 60 n/m" and "Raid n/m" where the two builds differ, and the raid rank otherwise.
- "For the raid build." sits under Rotation and priority.
- The rotation and cooldown cards now read the raid build's talents only, so Rapid Killing is no longer applied. Rapid Fire reads 5 minutes on both pages.

Survival has no `raidBuild`, so its page is unchanged.

## Builds and leveling order

Decoded again against `talents/hunter.json` (tier, then column). All five strings are legal at 51 points, and every prerequisite and row gate is met. The decode is the same as round 11:

| Build | Split | Points |
|---|---|---|
| BM leveling | 31/20/0 | Deadly Aspects 5, Endurance Training 4, Focused Fire 2, Bestial Swiftness 1, Unleashed Fury 5, Ferocity 5, Intimidation 1, Bestial Discipline 2, Frenzy 5, Bestial Wrath 1; Lethal Attacks 5, Improved Stings 2, Efficiency 4, Careful Aim 5, Rapid Killing 1, Trueshot Aura 1, Mortal Shots 2 |
| BM raid | 31/20/0 | As leveling, except Efficiency 3, Rapid Killing 0, Mortal Shots 4 |
| MM leveling | 15/31/5 | Deadly Aspects 5, Endurance Training 3, Focused Fire 2, Unleashed Fury 5; Lethal Attacks 5, Efficiency 5, Careful Aim 5, Rapid Killing 2, Trueshot Aura 1, Mortal Shots 5, Rapid Recuperation 2, Barrage 1, Ranged Weapon Specialization 5; Improved Tracking 5 |
| MM raid | 5/33/13 | Deadly Aspects 5; Lethal Attacks 5, Improved Stings 2, Efficiency 5, Careful Aim 5, Lone Wolf 1, Trueshot Aura 1, Mortal Shots 5, Barrage 3, Ranged Weapon Specialization 5, Sniper Shot 1; Improved Tracking 5, Survivalist 5, Surefooted 3 |
| SV | 5/11/35 | Deadly Aspects 5; Lethal Attacks 5, Careful Aim 5, Lone Wolf 1; Improved Tracking 5, Savage Strikes 2, Survivalist 3, Clever Traps 2, Surefooted 3, Deterrence 1, Predator's Edge 5, Counterattack 1, Resourcefulness 2, Expose Prey 2, Survivalist's Discipline 2, Strider Kick 1, Lightning Reflexes 5, Lacerating Strikes 1 |

The shipped orders in `addon-builds.json` are unchanged, and every leveling sentence matches them:
- BM: Deadly Aspects 10-14; Focused Fire 15-16, when tier 1 opens; Endurance Training from 17; Bestial Wrath at 40, with mail.
- MM: Lethal Attacks 10-14; Efficiency from 15; Rapid Killing 25-26, with Rapid Fire trained at 26; Rapid Recuperation 33-34.
- SV: Improved Tracking 10-14; Savage Strikes 15-16; Survivalist from 17; Lacerating Strikes at 40, after 30 Survival points; Lone Wolf is the level 60 point.

## Findings

None. Nothing on the four hunter pages or in the hunter card notes is false on 70755 or contradicts the page, the header card or the new raid-build rotation cards.

### The guide-header-build change against the prose

- **Beast Mastery.** The header now shows the raid build (Efficiency 3, Mortal Shots 4, no Rapid Killing). The prose agrees:
  - BM:81 names Rapid Killing as "the level 60 build's point".
  - BM:83 says "The raid build moves one point out of Efficiency and the point in Rapid Killing into Mortal Shots".
  - The Mortal Shots bullet gives both ranks: 212% at 2 points, "224% with the 4 points in the raid build". Its chips read Level 60 2/5 and Raid 4/5.
  - BM:89 says "Rapid Fire's five minute cooldown", which now matches the card. The Rapid Fire card note says the same.
- **Marksmanship.** The header now shows the raid build. Every key-talent bullet names its build:
  - Barrage, 3 points in the raid build and 1 point while leveling (chips 1/3 and 3/3).
  - Lone Wolf, Improved Stings, Surefooted and Sniper Shot are labelled "raid build".
  - Rapid Killing and Rapid Recuperation are labelled "leveling build".
  - MM:94 says five minutes, and MM:96 says "With the raid build, Sniper Shot comes next". Both match the raid-build card.
  - The card order matches the prose: Hunter's Mark, Rapid Fire, Aimed Shot (next auto more than 1.2 s away), Multi-Shot, Sniper Shot, Serpent Sting, Arcane Shot above 25% mana.
- No hunter page says "the build above" or refers to the header card as the level 60 build.

### Re-checked this round and still true

- **Talent text, 70755 rank values:**
  - Deadly Aspects: Hawk on Auto Shot, Beast on melee autos, 10%, 30%, 12 sec.
  - Focused Fire 2%, Unleashed Fury 15%, Ferocity 10%, Frenzy 100% for 30% over 8 sec, Bestial Wrath 50% for 18 sec on 120 s, Bestial Discipline 50%, Intimidation 3 sec.
  - Lethal Attacks 1% a point. Efficiency 3% a point on shots, stings and melee. Careful Aim 100% of Intellect. Lone Wolf 20%.
  - Rapid Killing: rank 1 is 1 min and 10%, rank 2 is 2 min and 20%.
  - Rapid Recuperation rank 2: 50%, or 100% once Rapid Killing is consumed.
  - Barrage 5% a point. Improved Stings rank 2 is 13%. Mortal Shots raises the crit bonus by 30%. Ranged Weapon Specialization 5%.
  - Sniper Shot rank 3 (1310786): "ranged damage plus 590", 4000 ms cast, 15000 ms cooldown.
  - Improved Tracking covers the seven types the pages list. Savage Strikes 2% a point. Survivalist 2% health a point. Surefooted 3%.
  - Predator's Edge: 30% crit bonus and 50% off hand. Resourcefulness: 60%, and on crits 50% regen for 30 sec.
  - Expose Prey 10% for 10 sec. Lightning Reflexes 15%. Lacerating Strikes 40% over 21 sec.
- **Tier gates** (tier and column in `talents/hunter.json`):
  - Focused Fire is tier 1. Lone Wolf and Rapid Killing are tier 2, opened by Lethal Attacks plus Careful Aim. Surefooted is tier 2, opened by Improved Tracking plus Survivalist. Sniper Shot and Lacerating Strikes are tier 6.
  - Counterattack follows Deterrence. Intimidation leads to Bestial Wrath.
- **Trainer levels** (`trainables/hunter.json`):
  - Heart of the Lion, Raptor Strike, Track Beasts and Auto Shot at 1. Raptor Strike ranks at 8 and 16.
  - Serpent Sting and Monkey at 4. Arcane Shot and Hunter's Mark at 6. Concussive Shot at 8. Hawk, Tame Beast and Track Humanoids at 10. Wing Clip and Mend Pet at 12. Mongoose Bite at 16. Multi-Shot and Track Undead at 18.
  - Aimed Shot, Cheetah, Disengage and Freezing Trap at 20. Rapid Fire at 26. Feign Death, Beast and Lacerate at 30. Explosive Trap at 34. Volley at 40. Falcon and Tranquilizing Shot at 60.
  - Wing Clip rank 3 (14268) is a level 60 row with 50 damage, which matches the card's "50 damage for a global".
- **Spelltips:**
  - Multi-Shot hits 3 targets and shares its cooldown with Aimed Shot. Aimed Shot rank 6 has a 2 s cast and a 6 s cooldown.
  - Rapid Fire speeds ranged and melee by 40% for 15 sec on 300 s.
  - Disengage: 5 s cooldown, "Disables auto-attack". Freezing Trap: "Any damage caused will break the ice".
  - Hawk rank 6 is 110 and rank 7 is 120 (a book). Beast rank 4 is 110. Falcon equals your highest Hawk.
  - Trueshot Aura rank 5 gives 100 to "party or raid members". Heart of the Lion is 40 + 4 a level, which is 280 at 60.
- **Racials** (races.json, 70755): Elune's Light 10% crit for 15 s on 3 min; Ruthless 1% crit; Artisan's Precision 1% hit with spells and attacks; Big Game Hunter 5% to Beasts; Blood Fury 10% AP for 15 s on 2 min; Command 5% pet damage; Berserking 10% for 15 s on 3 min.
- **Consumables and gear** (items.json):
  - Jagged Arrow 40. Ice Threaded Arrow 51, so "the early 50s" holds. Thorium Headed Arrow and Thorium Shells 52. Quickdraw Quiver 40.
  - Sniper Scope and Biznicks 247x128 Accurascope exist.
  - Major Frenzy Potion (250943) is level 55. Major Mana Potion, Elixir of the Mongoose and Ground Scorpok Assay exist.
  - Poached Sunscale Salmon gives 40 attack power and ranged attack power in Forever (presets.json notes, the sim audit rounds 8 and 9), so it is a fair hunter food.
  - The only bracer Agility enchant is Minor Agility (+1), so "no bracer enchant carries real Agility" holds.
- **Voice.** No em or en dashes on the pages or in the card notes. No sim, APL, preset, band, site or pipeline words in the prose, and no quoted nightly numbers.
- **Source links.** All 15 URLs were fetched again with a browser user agent.
  - 200: the three Blizzard articles, the three Warcraft Tavern pages and the six foreversixty.gg URLs.
  - 403: Wowhead Skyborne First Look, spell=1495 and spell=19306. Under the brief's rule that is unknown, not dead. Round 9 confirmed both spell labels through the nether tooltip endpoint.
  - Every entry follows the round 4 source-link rule.

### SHOULDs (wording, not findings)

- **BM:81, "Improved Stings (2 points) is the way down the Marksmanship tree".** This is exact for the raid build: Lethal Attacks 5, Efficiency 3 and Careful Aim 5 make 13, and Improved Stings brings it to the 15 that open Trueshot Aura and Mortal Shots. On the level 60 build, Efficiency 4 and Rapid Killing 1 already make 15 without it. Suggested wording: "Improved Stings (2 points) adds Serpent Sting damage and, in the raid build, is the way down to Mortal Shots."
- **Carried from round 11:**
  - MM:159, "so you have it for most hard pulls".
  - MM:94, five minutes for a player who raids on the leveling build.
  - survival.md:65, "smaller than Era's".
  - The Serpent Sting card's "throws away a tick".
  - "The Molten Core tome".
  - BM:81 naming Deflection.

## Decisions

- **D1. The guide-header-build change needs no prose edit on any hunter page.** Every talent claim names the build it belongs to, and every cooldown claim now matches the raid-build cards.
- **D2. The statPriority frontmatter order is not a finding.** The BM and MM band 60 raid factors in this tree put Critical strike above Intellect, while the frontmatter lists Intellect first. But `railStatRows` uses the frontmatter as a set and sorts the rows by scale factor, so the page renders the data's order. The prose explains each stat and does not rank them.
- **D3.** Round 11's D1 to D3 stand: "the hardest hit you have", "Hawk's Ahn'Qiraj rank", and the earlier rounds' decisions.

## Sims requested

1. **MM best in slot on the current builds (PROVISIONAL, carried).** `bis/hunter-marksmanship.json` (de044125c) still carries the old strings `0053550011503050` (raid) and `0051550001503050` (bare). Alliance raid hit there is 0 (insignificant), against 1.12 on Horde. Re-rank on `5/0052550011503051/50005003` and `532000050/005055200152105/5`, then re-read the MM stat table, the enchant table and the hit sentence at MM:119.
2. **MM race order on the new raid build (carried).** Night Elf against Human against Dwarf, and Troll against Orc, at seeds 1|50001, band 60 raid.
3. **Carried from earlier rounds:**
   - Survival weapon enchant: Crusader on both weapons against Agility on both, `5/005005001/500230231051222151`, Night Elf, seeds 1|50001.
   - Race re-runs after the final re-rank: SV Orc against Troll, BM Night Elf against Human, BM Orc against Troll.

## Sim findings

- **Carried, LIKELY:**
  - Aspect of the Falcon's melee proc does not need the talent in the engine (b86fa7d39).
  - Sniper Shot's cast time scales with ranged haste in the engine (`sniper_shot.go`), which the 70755 text does not state.
  - Summon Hawk is unmodeled.
  - Pet non-stat buffs are blocked by RefusePlayerBuffs.
  - The APL top-level notes (not rendered) still say "Aimed Shot is only the pre-pull opener".

## Sources

**Forever Sixty client data, 1.60.1.70755**
- `talents/hunter.json`: rank text, tier and column for all three trees.
- `trainables/hunter.json`: every trainer level above.
- `spelltips/hunter.json`: 781 to 14273 Disengage, 2643, 3045, 5384, 1499 to 14311, 1510 to 14295, 13813 to 14317, 14268, 14322, 25296, 469145, 1299447, 20904, 1310786, 20910, 14271, 14325, 19574, 19577, 20906, 409580, 1317257, 1299332.
- `races.json`: the racial rows for Human, Dwarf, Night Elf, Orc and Troll.
- `items.json` and `simconsumes.json`: the ammo, quiver, scope, potion, elixir and food rows above.
- `enchants.json`: the bracer Agility rows.
- `addon-builds.json`: the shipped orders.
- `bis/hunter-*.json`: band 60 talent strings and scale factors.

**Repository**
- Commit 3bd24d3a2: `describedBuildTrees`, the raid-build rail and rotation, and key-talent buildRanks.
- `web/src/lib/guides/key-talents.ts`, `web/src/lib/guides/rail-stats.ts` and `web/src/pages/guides/[class]/[spec].astro`.
- `data/curated/apl/hunter-{beast-mastery,marksmanship,survival}.json`: card notes and gates.
- `data/curated/presets.json`: the Forever food notes.

**Web** (fetched 2026-11-04, browser user agent)
- Blizzard, World of Warcraft: Forever Class Deep Dives: Hunter and Druid, https://news.blizzard.com/en-us/article/24301515/world-of-warcraft-forever-class-deep-dives-hunter-and-druid (200)
- Blizzard, Deep Dive panel recap, https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap (200)
- Blizzard, What's Next panel recap, https://news.blizzard.com/en-gb/article/24303862/world-of-warcraft-forever-whats-next-panel-recap (200)
- Warcraft Tavern, Hunter Class Guide (Forever), https://warcrafttavern.com/forever/guides/hunter/ (200)
- Warcraft Tavern, Hunter Leveling Guide (Forever), https://www.warcrafttavern.com/forever/guides/hunter-leveling/ (200)
- Warcraft Tavern, Hunter and Druid Deep Dive, https://www.warcrafttavern.com/forever/news/powershifting-is-back-sort-of-hunter-druid-deep-dive-for-wow-forever/ (200)
- Wowhead Forever: Skyborne First Look, spell=1495 and spell=19306 (403 to automated fetches, so unknown)

## Fix log (round 12)

- SHOULD BM:81 (Improved Stings): applied, beast-mastery.md:81: "Improved Stings (2 points) adds Serpent Sting damage and, in the raid build, is the way down to Trueshot Aura and Mortal Shots." Verified addon-builds.json hunter-beast-mastery: raid takes Lethal Attacks 5, Improved Stings 2, Efficiency 3, Careful Aim 5 (15) before tier 4; leveling reaches 15 without Improved Stings (Efficiency 4, Rapid Killing 1).
- Other hunter items: not in this lane's scope.
