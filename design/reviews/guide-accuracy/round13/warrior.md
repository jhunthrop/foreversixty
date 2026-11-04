# Guide accuracy round 13: Warrior

Audited on the round 13 audit tree (main 604b78a80: twelve guide rounds, sim round 15, guide-header-build), client 1.60.1.70755, engine pin 9f7f1f6e4 (`sim/enginever/version.go`). Files: `web/src/content/guides/warrior/{index,arms,fury,protection}.md` and every card note in `data/curated/apl/warrior-{arms,fury,protection}.json`. I read the round 1 to 12 warrior reports and fix logs, `round4/warrior-levels.md`, race-sims for rounds 2 to 5, and the ledger.

Changes since round 12's tree (a8b147a6f): `git diff a8b147a6f 604b78a80` touches no warrior guide, no warrior APL, no `addon-builds.json` entry, no warrior BiS file and no client data. Round 12 had no warrior fixes. This round is a fresh full line-by-line pass on unchanged text. Every claim was re-derived from the 70755 files rather than taken from earlier reports.

## Builds decoded (unchanged, all legal on 70755)

- Arms `build` `13323213032215201/0505/205`: IHS 1, Deflection 3, Improved Rend 3, Improved Charge 2, ITM 3, Improved Overpower 2, AM 1, Deep Wounds 3, 2H Spec 3, Impale 2, Bloodthrill 2, Sweeping Strikes 1, Weaponmaster 5, Improved Slam 2, MS 1 / Cruelty 5, Unbridled Wrath 5 / Improved Bloodrage 2, Iron Will 5.
- Arms `raidBuild` `30303213132215201/05050005012/1` (header card): IHS 3, Improved Rend 3, ITM 3, Improved Overpower 2, AM 1, Deep Wounds 3, Spearing Strike 1, 2H Spec 3, Impale 2, Bloodthrill 2, Sweeping Strikes 1, Weaponmaster 5, Improved Slam 2, MS 1 / Cruelty 5, Unbridled Wrath 5, Enrage 5, Reckless Swings 1, Improved Execute 2 / Improved Bloodrage 1.
- Fury `build` `30323013002/050530025122100231/2`: IHS 3, Improved Rend 3, Improved Charge 2, ITM 3, AM 1, Deep Wounds 3, Impale 2 / Cruelty 5, Unbridled Wrath 5, Furious Precision 3, Enrage 2, DW Spec 5, Reckless Swings 1, Improved Execute 2, Improved Berserker Rage 2, Death Wish 1, Improved Whirlwind 2, Flurry 3, Bloodthirst 1 / Improved Bloodrage 2.
- Protection `build` `35300003/0/250533120330001351`: IHS 3, Deflection 5, Improved Rend 3, Deep Wounds 3 / Improved Bloodrage 2, Shield Spec 5, Anticipation 5, Improved Revenge 3, Improved Thunder Clap 3, Last Stand 1, Master of Defense 2, Defiance 3, Improved Sunder Armor 3, Concussion Blow 1, Focused Rage 3, Bastion 5, Shield Slam 1.
- Prerequisites, re-read from `talents/warrior.json`, are all met: AM needs ITM 3, Deep Wounds needs Improved Rend 3, Impale needs Deep Wounds 3, MS needs Sweeping Strikes, Bloodthirst needs Death Wish, Master of Defense needs Shield Spec 5, and Shield Slam needs Concussion Blow.

Shipped leveling orders (`addon-builds.json`, index 0 = level 10) are unchanged and re-mapped this round:
- Arms: IHS 10, Deflection 11-13, Improved Rend 14-16, Improved Charge 17-18, ITM from 19.
- Fury: Cruelty 10-14, Unbridled Wrath 15-19, Furious Precision from 20.
- Protection: Improved Bloodrage 10-11, Shield Spec 12-16, Anticipation from 17.

## Findings

None. No statement on the four pages or in the warrior card notes is false on 70755. None contradicts its own page, the header build card, the cards or the shipped point order.

## Shoulds (wording, not findings; all carried, none newly false)

- Arms Bloodrage card [2] still adds "half again with the level 60 build's two" under the "For the raid build." caption. This is true. An optional trim is in round 12.
- arms.md:125 skips the level 19 point, the first point of Improved Tactical Mastery.
- arms.md:91 "Mortal Strike on cooldown and Execute on every other global" leaves out Reckless Swings, which sits above Execute on the raid cards and works in Berserker Stance.
- The Reckless Swings health-cost wording differs between Fury ("the healers cover without noticing") and Arms ("tell your healers").
- fury.md:54 "No healing cut or debuff of its own": Fury still brings Sunder Armor and Demoralizing Shout. Suggested: "No healing cut of its own for the raid".

## Checked and right (re-derived this round)

### Stances

From `SpellShapeshift` (70753 raw table, the latest stance table in the scratchpad; the 70755 hotfix push changed no warrior stance rows). Masks: Battle 65536, Defensive 131072, Berserker 262144.

- **Thunder Clap** (11581, 6343): Battle and Defensive. So index.md:30 "works in Defensive Stance" holds.
- **Defensive Stance only**: Revenge (25288, 11601), Shield Block (2565), Shield Wall (871) and Taunt (355). That is Protection card [1] "Revenge, Shield Block and Shield Wall all need it".
- **Berserker Stance only**: Recklessness (1719), Whirlwind (1680), Intercept (20617), Berserker Rage (18499) and Pummel. That matches arms.md:54 and :91 to :93, fury.md:96, and the Recklessness swap.
- **Battle Stance only**: Overpower (11585), Charge (11578), Sweeping Strikes (12292) and Spearing Strike (1310222). So:
  - "Sweeping Strikes in Battle Stance, swap to Berserker Stance, Whirlwind" holds.
  - "Overpower ... [does] not work in Berserker Stance" holds.
  - Spearing Strike "stops at the execute swap" holds.
  - Fury's ITM line "a Charge's rage come[s] with you when you swap back" holds.
- **Rend** (11574): Battle and Defensive, not Berserker.
- **Battle and Berserker**: Execute (20662, 5308), Hamstring (7373) and Heroic Strike's sibling rows.
- **No stance row**: Reckless Swings (1329555), Mortal Strike, Bloodthirst and Slam, so they work in any stance. That matches Arms card [1] "Mortal Strike and Execute still work there".

### Spell text and constants (`spelltips/warrior.json`, `spellconst/warrior.json`)

- **Slam** 11605: 1500 ms cast, 1500 ms global, 15000 ms cooldown. Improved Slam 2 is "-5.0 sec, the cast time and global cooldown by 0.50 sec, no longer interrupts or delays your melee swing". That gives "cast and global drop to one second" and "15 seconds to 10" (arms.md:73, index.md:29).
- **Mortal Strike** 21553: 50% healing cut, 6000 ms. **Bloodthirst** 23894: 45% AP, 6000 ms.
- **Overpower** 11585: "only useable after the target dodges", 5 s cooldown.
- **Execute** 20662: 150 cost, converts extra rage. Improved Execute 2 is "by 5", so 10 rage.
- **Sweeping Strikes** 12292: "next 5 melee attacks". **Whirlwind** 1680: "up to 4 enemies ... both melee weapons", 10000 ms. Improved Whirlwind 2 is "-2 sec, +30% to the closest enemy".
- **Death Wish** 12328: 20% physical, Fear immunity, 5% more damage taken, 30 s, 180000 ms, 10 rage.
- **Recklessness** 1719: 100% crit, 15 s, 1800000 ms.
- **Berserker Rage** 18499: 10 s, 30000 ms, level 30. Improved Berserker Rage 2 is "10 Rage, 100% remove movement impairing".
- **Reckless Swings** 1329555: 1000 ms global, 15000 ms cooldown, 8000 ms, "next 4 melee attacks generate 5 additional Rage".
- **Spearing Strike** 1310222: 150 cost, 20000 ms, 75% plus 75% against Giants and Dragonkin.
- **Bloodrage** 2687: 10 plus 10 over 10 s, 60000 ms, no global. Improved Bloodrage is 25% at rank 1 and 50% at rank 2, so 15 and 15 for Fury. With Improved Charge 2, Charge 11578 gives 21 rage, which is fewer, so fury.md:84 holds.
- **Battle Shout** 25289: "all party or raid members within 20 yards", 3 min, 10 rage.
- **Demoralizing Shout** 11556: 45 s. The 204 at 60 is the decision from rounds 4 and 8.
- **Shield Block** 2565: 75%, 7 s, 2 attacks, 5000 ms, no global. **Shield Wall** 871: 60%, 12 s, 900000 ms.
- **Last Stand** 12975: 30%, 20 s. **Concussion Blow** 12809: 5 s. **Challenging Shout** 1161: 10 yd, 6 s, 600000 ms.
- **Thunder Clap** 11581: 20% slow, 30 s, 4 targets. **Revenge** 25288: 5 rage, which Focused Rage takes to 2.
- **Taunt** 355: "no effect if the target is already attacking you". **Victory Rush** 402927: 10%, 20 s window, 30000 ms, level 20.
- **Tactical Mastery** 1310185: 10 rage at 14. ITM 3 adds 15, for 25.
- **Charge** 11578: 15 rage, "Cannot be used in combat".
- **Hamstring** 7373: 10 rage.

### Talent text at the build's ranks (`talents/warrior.json`)

- **Arms**:
  - IHS: 1 / 3 rage.
  - Improved Charge 2: +6. ITM 3: +15.
  - Improved Overpower 2: 50%. Bloodthrill 2: 20% on main-hand attacks against your Rend.
  - Weaponmaster 5: 5% crit for axe and polearm, 15% armor for mace and staff, 5% extra attack for sword.
  - Improved Rend 3: 35%. Deep Wounds 3: 60% over 12 s. Impale 2: +20% bonus, which is 220%.
  - 2H Spec 3: 6%. Iron Will 5: 15% duration.
- **Fury**:
  - Cruelty 5: 5%. Unbridled Wrath 5: 60%.
  - Enrage: 2% a rank, 30% chance, 12 s. Furious Precision 3: 10% off-hand hit.
  - DW Spec 5: 25% damage and 50% off-hand rage. Flurry 3: 30% for 3 swings.
- **Protection**:
  - Shield Spec 5: 5% block and a 100% chance at 5 rage; ranks 1 to 4 are only a chance. Anticipation 5: 20 Defense. Master of Defense 2: 100% chance at 5 rage on a dodge or parry with a shield.
  - Defiance 3: 15% in Defensive Stance with a shield. Bastion 5: 10% with a shield. Focused Rage 3: -3.
  - Improved Revenge 3: 60%. Improved Thunder Clap 3: -6. Improved Sunder Armor 3: -3.
  - Improved Shield Wall 2: "by 11.0 min", so 4 minutes. Vanguard: "Charge usable while in Defensive Stance". Improved Shield Bash 2: 3 s silence.
- **Improved Shield Wall swap** (protection.md:83): Improved Thunder Clap 3 to 1 and Improved Shield Wall 0 to 2.
  - The build stays legal: 25 points before tier 4, 31 before Bastion, 36 before Shield Slam.
  - Thunder Clap costs 4 more.
  - A 4 minute cooldown gives two or three Shield Walls on a long fight.

### Racials (`spelltips/shared.json`)

- Elune's Light: 10% crit, 15 s, 180000 ms. Blood Fury: 10% AP, 15 s, 120000 ms. Berserking: 10% attack speed, 180000 ms.
- Touch of the Grave: a 20% chance to drain health. The 15 s internal cooldown is from the engine, per round 9.
- Ruthless: 1% crit. Artisan's Precision: 1% hit. Wind Blessed: 2% haste.
- Quickness: 1% dodge and 2% speed. Endurance: 5% health.
- Stoneform: removes and grants immunity to bleed, poison and disease, -10% physical, 8 s, 180000 ms.
- No racial carries weapon skill, so "no race brings it any more" holds.
- Tauren's four racials carry no damage.

### Items

- Obsidian Edged Blade (18822) is a two-handed sword, subclass 8, slot 17.
- Edgemaster's Handguards (14551) and Arcanite Champion are in `items.json`.
- The consumables named on the pages are in `simconsumes.json` (round 1 check, unchanged data).

### Card notes against their actions, thresholds and the client values

- **Arms**: Recklessness swap, Bloodrage under 80, Battle Shout, Rend, Overpower, Mortal Strike, Reckless Swings, Execute, Slam (cast under 1.5 s), Spearing Strike, Hamstring (Windfury and not from 25%), and Heroic Strike at 90 or more.
- **Fury**: Bloodrage under 80, Death Wish (210 s, execute, or 30 s), Recklessness (execute or 15 s), Battle Shout, Execute first, Bloodthirst, Whirlwind, Reckless Swings, Berserker Rage under 80, Overpower before 36, and Heroic Strike at 60 or more but not from 25%.
- **Protection**: Defensive Stance, Last Stand under 35%, Shield Wall under 40%, Bloodrage under 60, Shield Block, Revenge, Shield Slam, Thunder Clap and Demoralizing Shout with a 3 s refresh, Battle Shout, Sunder, and Heroic Strike at 40 or more.

### Levels 1 to 19 and 20 to 60 on all three pages

- They match the shipped orders above and the trainer levels in `trainables/warrior.json`:
  - Charge and Rend 4, Thunder Clap 6, Hamstring 8.
  - Bloodrage, Defensive Stance, Sunder Armor and Taunt 10.
  - Overpower and Shield Bash 12.
  - Tactical Mastery, Revenge and Demoralizing Shout 14.
  - Shield Block 16, Victory Rush 20, Berserker Stance 30, Sweeping Strikes 35, Whirlwind 36.
  - Mortal Strike, Bloodthirst and Shield Slam 40.
- Arms Improved Overpower 22 and Bloodthrill 33, after Deep Wounds 25-27 and Impale 31-32.
- Protection Shield Slam is the 31st point at 40, and Focused Rage's last point is at 41.
- The two blocks read correctly in order and do not contradict each other.

### Header build card

Arms prose keeps the level 60 build and the raid build apart:
- Improved Charge is "level 60 build only".
- Improved Bloodrage is two points in the level 60 build and one in the raid build.
- Reckless Swings and Spearing Strike are "raid build only".

So nothing contradicts the raid-build header card. Fury and Protection have no `raidBuild`.

### Source links, fetched this round

- These return 200:
  - Blizzard Deep Dive recap
  - Wowhead Skyborne First Look
  - Power Up Gaming
  - Talents Forever data.json
  - Warcraft Tavern Classic PvE DPS warrior page
  - Warcraft Tavern Forever warrior and Fury pages
  - foreversixty.gg /sim/weights, /sim/drops, /planner and /bis/warrior/protection
- Icy Veins racial rework returns 403 (unknown, not dead).
- Every entry follows the round 4 source-link rule.

### Voice

No em or en dashes, no quoted nightly numbers, and no stat ranking inside Stat priority.

## Decisions

- **Zero confirmed findings for warrior this round**, for the third round running. The text is unchanged since round 11, and this round's fresh pass re-derived every stance, cost, cooldown, rank value, prerequisite, trainer level and racial from the 70755 files. Nothing is false.
- **The stance claims are now settled from the stance table itself**, not inferred. The sheet above can be reused by later rounds.
- **Race lines stay.** They are unmeasured on the new builds, and no measured result or client value contradicts them.

## Sims requested

Carried, still open (none is claimed in prose):
- Race re-runs on the Arms `raidBuild` and Fury `build`: raid 60, every race on both factions, 10,000 iterations, seeds 1 and 50001. Lines affected: arms.md:109-111 and fury.md:112-114.
- Dwarf against Human breakdown.
- Blood Fury held timing.
- Stoneform rotation line.
- Protection Improved Shield Wall build (Shield Wall is never cast in the tank sim).
- PROVISIONAL: any prose-vs-BiS-row mismatch after the BiS re-rank on the shipped raid builds. None found on this tree.

## Sim findings

None new.

## Sources

- Client data 1.60.1.70755:
  - `talents/warrior.json`: rank tables and prereq ids.
  - `trainables/warrior.json`.
  - `spelltips/warrior.json` and `spellconst/warrior.json`: 11605, 21553, 23894, 11585, 20662, 12292, 1680, 12328, 1719, 18499, 1329555, 1310222, 2687, 11578, 25289, 11556, 2565, 871, 12975, 12809, 1161, 11581, 25288, 355, 402927, 1310185 and 7373.
  - `spelltips/shared.json` racials.
  - `items.json` (18822, 14551, Arcanite Champion).
  - `addon-builds.json`.
- Stance table: wago SpellShapeshift at 1.60.1.70753 (scratchpad `c755/fetch/data/builds/1.60.1.70753/raw/SpellShapeshift.csv`).
- Engine pin 9f7f1f6e4 (`sim/enginever/version.go`).
- `git diff a8b147a6f 604b78a80` (no warrior changes).
