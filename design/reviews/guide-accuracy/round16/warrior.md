# Guide accuracy round 16: Warrior

Audited on the round 16 audit tree (land/guides-r15 2c3135a2e: fifteen guide rounds on launch-state main), client 1.60.1.70755, engine pin b2c4a2d05 (`sim/enginever/version.go`). Files: `web/src/content/guides/warrior/{index,arms,fury,protection}.md` and every card note in `data/curated/apl/warrior-{arms,fury,protection}.json`. I read the round 14 and 15 warrior reports and fix logs, `round4/warrior-levels.md` and the round 2 and 1 briefs, and I relied on the earlier rounds' sheets where no data has changed.

## Changes since round 15's tree (4600b7d82)

- `git diff 4600b7d82..HEAD` touches no warrior guide, no warrior APL, and no warrior file in `addon-builds.json`, `bis/warrior-*`, `talents/warrior.json`, `trainables/warrior.json` or `races.json`. It also leaves `web/src/lib/guides`, `web/src/components/guides` and `sim/enginever` alone. The engine pin is still b2c4a2d05.
- Round 15 filed no warrior findings and applied no warrior fixes, so this round has no fixes of its own to re-check. The whole page was re-audited line by line anyway.

## Builds decoded (re-decoded against `talents/warrior.json`, tier and column order)

- Arms `build` `13323213032215201/0505/205` (34 / 10 / 7):
  - Arms: IHS 1, Deflection 3, Improved Rend 3, Improved Charge 2, ITM 3, Improved Overpower 2, AM 1, Deep Wounds 3, Two-Handed Spec 3, Impale 2, Bloodthrill 2, Sweeping Strikes 1, Weaponmaster 5, Improved Slam 2, Mortal Strike 1.
  - Fury: Cruelty 5, Unbridled Wrath 5.
  - Protection: Improved Bloodrage 2, Iron Will 5.
- Arms `raidBuild` `30303213132215201/05050005012/1` (32 / 18 / 1):
  - Arms: as above, but IHS 3, Deflection 0, Improved Charge 0 and Spearing Strike 1.
  - Fury: Cruelty 5, Unbridled Wrath 5, Enrage 5, Reckless Swings 1, Improved Execute 2.
  - Protection: Improved Bloodrage 1.
- Fury `build` `30323013002/050530025122100231/2` (17 / 32 / 2):
  - Arms: IHS 3, Improved Rend 3, Improved Charge 2, ITM 3, AM 1, Deep Wounds 3, Impale 2.
  - Fury: Cruelty 5, Unbridled Wrath 5, Furious Precision 3, Enrage 2, Dual Wield Spec 5, Reckless Swings 1, Improved Execute 2, Improved Berserker Rage 2, Death Wish 1, Improved Whirlwind 2, Flurry 3, Bloodthirst 1.
  - Protection: Improved Bloodrage 2.
- Protection `build` `35300003/0/250533120330001351` (14 / 0 / 37):
  - Arms: IHS 3, Deflection 5, Improved Rend 3, Deep Wounds 3.
  - Protection: Improved Bloodrage 2, Shield Spec 5, Anticipation 5, Improved Revenge 3, Improved Thunder Clap 3, Last Stand 1, Master of Defense 2, Defiance 3, Improved Sunder Armor 3, Concussion Blow 1, Focused Rage 3, Bastion 5, Shield Slam 1.
- **Prerequisites** (from `prereq_talent_id`) are all met: AM needs ITM 3, Deep Wounds needs Improved Rend 3, Impale needs Deep Wounds 3, Mortal Strike needs Sweeping Strikes, Bloodthirst needs Death Wish, Master of Defense needs Shield Spec 5, Shield Slam needs Concussion Blow.
- **Tier gates:**
  - Fury's Arms points put 11 into tiers 0 and 1, which unlocks Deep Wounds. Without Improved Charge that falls to 9, so fury.md:78 holds.
  - Protection puts 30 points in its own tree before Shield Slam at 40. Fury puts 30 in its tree before Bloodthirst at 40.
- **Shipped orders** (`addon-builds.json`) are unchanged since round 15, and every leveling paragraph still matches them.

## Findings

None. No statement on the four pages or in the warrior card notes is false on 70755. None contradicts its own page, the header build card, the cards, the live BiS rows, the stat tables, the race table or the shipped point order.

Counts: CONFIRMED 0, LIKELY 0, NEEDS-SIM 0, NEEDS-OWNER 0.

## Checked and right this round (re-derived from client data)

- **Talent rank text** (`talents/warrior.json`):
  - Arms: Improved Rend 35%; Improved Charge +6; ITM +15; Improved Overpower 50%; Deep Wounds 60% over 12 s; Two-Handed Spec 6%; Impale +20% critical strike damage bonus, which is 220% of a normal hit; Bloodthrill 20% on main-hand hits against your Rend.
  - Weaponmaster 5: axe or polearm 5% crit, mace or staff 15% armor ignored, sword 5% extra attack. Improved Slam 2: -5 s cooldown, -0.5 s cast and global, no swing interrupt.
  - Fury: Cruelty 5%; Unbridled Wrath 60%; Furious Precision 10% off-hand hit; Enrage at rank 2 gives a 30% chance at 4% for 12 s; DW Spec 25% damage and 50% off-hand rage; Improved Execute -5; Improved Berserker Rage 10 rage plus a 100% slow removal; Improved Whirlwind -2 s and 30%; Flurry 3 gives 30% for 3 swings.
  - Protection: Improved Bloodrage 25% and 50%; Shield Spec 5% and a 100% chance at 5 rage on block at rank 5 only (protection.md:129 "from the fifth point" holds); Iron Will 15%; Anticipation 20 Defense; Improved Revenge 60%; Improved Thunder Clap -6; Master of Defense; Defiance 15% with a shield; Improved Sunder Armor -3; Improved Shield Wall -11 min; Concussion Blow 5 s; Focused Rage -3; Bastion 10%.
- **Spell values** (`spelltips/warrior.json`, `spellconst/warrior.json`):
  - Rage and cooldowns: Charge rank 3 15 rage; Battle Shout 10 rage, 20 yards, 3 min; Bloodrage 10 + 10 over 10 s, 60 s, no global; Mortal Strike 6 s and 50% healing cut.
  - Shield Block: gcd 0, 5 s cooldown, 7 s, 2 blocks, 75%.
  - Reckless Swings: gcd 1000, 15 s cooldown, 4 swings, 5 rage, 8 s. The 25% base-health cost is SpellPower PowerCostPct 25, power type health (engine `reckless_swings.go`).
  - Spearing Strike: 15 rage, 20 s, 75% plus 75% against giants and dragonkin. Execute: 15 rage. Death Wish: 10 rage, 20%, +5% damage taken, 30 s, 3 min.
  - Recklessness: 100% crit, 15 s, 30 min. Berserker Rage: 30 s, 10 s fear immunity. Whirlwind: 4 targets, both weapons, 10 s.
  - Thunder Clap: 20% attack-speed slow, 4 targets. Shield Wall: 60%, 12 s, 15 min. Challenging Shout: 6 s, 10 min. Last Stand: 30% for 20 s.
  - Victory Rush: 10%, 20 s window, 30 s. Demoralizing Shout 11556: -196 at base level 54 with points per level -1.4, so 204 at 60 (protection.md:93 and card [9]).
- **Stances** (engine b2c4a2d05 `RegisterSpell` masks):
  - Battle Stance only: Overpower, Spearing Strike, Sweeping Strikes and Charge.
  - Berserker Stance only: Whirlwind and Berserker Rage (also Recklessness, per `recklessness.go`).
  - Rend: Battle and Defensive. Execute: Battle and Berserker. Mortal Strike, Slam and Bloodthirst: any stance.
  - Shield Block, Revenge and Shield Wall: Defensive.
  - So arms.md:54, :87, :91, :93 and protection card [1] hold.
- **Trainer levels** (`trainables/warrior.json`):
  - Charge and Rend 4, Thunder Clap 6, Hamstring 8.
  - Bloodrage, Defensive Stance, Sunder Armor and Taunt 10.
  - Overpower and Shield Bash 12, Revenge, Demoralizing Shout and Tactical Mastery 14, Shield Block 16.
  - Victory Rush and Cleave 20, Berserker Stance, Berserker Rage and Intercept 30, Whirlwind 36, Recklessness 50.
- **Level 10 quest and Dual Wield:**
  - Warcraft Tavern Forever leveling lists the level 10 class quest as teaching Defensive Stance, Taunt and Sunder Armor, and Dual Wield from the trainer at 20.
  - Method's Forever leveling guide describes the follow-up quest rewarding a one-handed weapon with several weapon types.
  - So the "Levels 1 to 19" lines hold.
- **Card notes** (all three APLs): re-read against their actions, conditions and the values above. Nothing false.
- **Races, BiS-linked gear and professions:** no data changed since round 15's checks (sim round 16 A4 and F3; bonding of 249469, 249470, 12790 and the Blacksmithing pieces). Every line still holds.
- **Source links**, fetched this round:
  - Every link returns 200 except the Icy Veins racial rework, which returns 403. That is unknown, not dead.
  - Every link follows the round 4 rule. The labels match the pages.
- **Voice:** no em or en dashes, no quoted nightly numbers, no site or sim talk outside the planner links.

## Shoulds (wording, not findings)

- **arms.md:79.** "Two more points of Improved Heroic Strike make your rage dump cost 3 less." The total at 3 points is 3 less, but the two extra points over the level 60 build add only 2. Suggested: "Two more points of **Improved Heroic Strike** take your rage dump to 3 less."
- Carried from round 15:
  - fury.md:74 "Fury never casts Rend" next to fury.md:124's Rend advice.
  - arms.md:115's "other pair" sentence.
  - The Arms Bloodrage card [2] trim.
  - arms.md:125 skipping the level 19 point.
  - arms.md:91 leaving out Reckless Swings.
  - The Reckless Swings health-cost wording.
  - fury.md:54 "debuff".
  - fury.md:100 Edgemaster's wording.

## Decisions

- **Race picks stay** (Night Elf / Undead on Arms and Fury, Night Elf / Tauren on Protection): every measured gap is inside 1%. Skyborne is named only as the paid option.
- **Professions lines stand** on the live raid rows.
- **Not filed, per the brief:** the BiS race captions (data lane) and the leveling strip rendering (land/leveling-tree).

## Sims requested

All carried, none run yet:

- **Arms and Fury race order on the guide raid builds:** raid 60, every race on both factions including both Skyborne, 10,000 iterations, seeds 1 and 50001. This checks arms.md:109 and :111 and fury.md:112.
- **Protection races on the current build:** tank fight, every race, seeds 1 and 50001, reporting damage taken, TMI and threat. This checks protection.md:113.
- **Protection Improved Shield Wall build**, the Dwarf against Human breakdown, and the Stoneform rotation line.

## Sim findings

None new. Sim round 17 seed 1 (Heroic Strike as an auto attack for Unbridled Wrath, Flurry and Reckless Swings) stays with the sim lane. No page claims either way.

## Sources

- **Client data 1.60.1.70755:** `talents/warrior.json`, `spelltips/warrior.json`, `spellconst/warrior.json`, `trainables/warrior.json`, `addon-builds.json`, `races.json`, `bis/warrior-*.json`.
- **Engine b2c4a2d05:** `sim/warrior/*.go` RegisterSpell stance masks, `reckless_swings.go`, `recklessness.go`, `constants_auto_gen.go`.
- **Warcraft Tavern, Forever Warrior Leveling Guide:** https://www.warcrafttavern.com/forever/guides/warrior-leveling/
- **Method, WoW Forever Warrior Leveling Guide and Talents:** https://www.method.gg/wow-forever/wow-forever-warrior-leveling-guide-and-talents
- **Frontmatter links checked:** Blizzard Deep Dive recap, talentsforever.com/data.json, Warcraft Tavern (Classic PvE DPS warrior, Forever warrior, Forever Fury), Icy Veins racial rework (403), Wowhead Skyborne First Look, Power Up Gaming talent changes, foreversixty.gg /sim/weights, /sim/drops, /planner, /bis/warrior/protection.

## Fix log (round 16)

- arms.md:79 SHOULD: fixed, "take your rage dump to 3 less" (Improved Heroic Strike 1/2/3 rage in 70755).
