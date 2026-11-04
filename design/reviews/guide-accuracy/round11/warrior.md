# Guide accuracy round 11: Warrior

Audited on the round 11 audit tree (main d54de48aa: guide rounds 1 to 10, sim round 15, leveling-order), client 1.60.1.70755, engine pin 9f7f1f6e4 (`sim/enginever/version.go`; engine source read from scratchpad/gr10/engine). Files: `web/src/content/guides/warrior/{index,arms,fury,protection}.md` and every card note in `data/curated/apl/warrior-{arms,fury,protection}.json`. I read the round 1 to 10 warrior reports and fix logs, the ledger and the sim round 15 warrior lane report.

Since round 10's audit tree, the only warrior change is the round 10 fix commit eabba8f84 (arms.md, fury.md, index.md, Arms and Fury card notes). Every line it touched was re-derived from the 70755 client data, and every other line on the four pages was re-read against the round 1 to 10 checks.

## Builds decoded (all legal on 70755, unchanged since round 10)

- Arms `build` `13323213032215201/0505/205`: IHS 1, Deflection 3, Improved Rend 3, Improved Charge 2, ITM 3, Improved Overpower 2, AM 1, Deep Wounds 3, 2H Spec 3, Impale 2, Bloodthrill 2, Sweeping Strikes 1, Weaponmaster 5, Improved Slam 2, MS 1 (34) / Cruelty 5, Unbridled Wrath 5 (10) / Improved Bloodrage 2, Iron Will 5 (7).
- Arms `raidBuild` `30303213132215201/05050005012/1`: IHS 3, Improved Rend 3, ITM 3, Improved Overpower 2, AM 1, Deep Wounds 3, Spearing Strike 1, 2H Spec 3, Impale 2, Bloodthrill 2, Sweeping Strikes 1, Weaponmaster 5, Improved Slam 2, MS 1 (32) / Cruelty 5, Unbridled Wrath 5, Enrage 5, Reckless Swings 1, Improved Execute 2 (18) / Improved Bloodrage 1 (1). Fury tier 3 is reached with exactly 15 (Cruelty, Unbridled Wrath, Enrage). Mortal Strike's prerequisite, Sweeping Strikes, is taken.
- Fury `build` `30323013002/050530025122100231/2`: IHS 3, Improved Rend 3, Improved Charge 2, ITM 3, AM 1, Deep Wounds 3, Impale 2 (17) / Cruelty 5, Unbridled Wrath 5, Furious Precision 3, Enrage 2, DW Spec 5, Reckless Swings 1, Improved Execute 2, Improved Berserker Rage 2, Death Wish 1, Improved Whirlwind 2, Flurry 3, Bloodthirst 1 (32) / Improved Bloodrage 2 (2).
- Protection `build` `35300003/0/250533120330001351`: IHS 3, Deflection 5, Improved Rend 3, Deep Wounds 3 (14) / Improved Bloodrage 2, Shield Spec 5, Anticipation 5, Improved Revenge 3, Improved Thunder Clap 3, Last Stand 1, Master of Defense 2, Defiance 3, Improved Sunder Armor 3, Concussion Blow 1, Focused Rage 3, Bastion 5, Shield Slam 1 (37).

Shipped orders (`addon-builds.json`, index 0 = level 10), re-derived this round:
- Arms leveling: IHS 10, Deflection 11-13, Improved Rend 14-16, Improved Charge 17-18, ITM 19-21, Improved Overpower 22-23, AM 24, Deep Wounds 25-27, 2H Spec 28-30, Impale 31-32, Bloodthrill 33-34, Sweeping Strikes 35, Weaponmaster 36-39, MS 40, Weaponmaster 41, Improved Slam 42-43, Cruelty 44-48, Unbridled Wrath 49-53, Improved Bloodrage 54-55, Iron Will 56-60.
- Fury (leveling and raid identical): Cruelty 10-14, Unbridled Wrath 15-19, Furious Precision 20-22, Enrage 23-24, DW Spec 25-29, Death Wish 30, Reckless Swings 31, Improved Execute 32-33, Improved Berserker Rage 34-35, Improved Whirlwind 36-37, Flurry 38-39, Bloodthirst 40, Flurry 41, IHS 42-44, Improved Rend 45-47, Improved Charge 48-49, ITM 50-52, AM 53, Deep Wounds 54-56, Impale 57-58, Improved Bloodrage 59-60.
- Protection: Improved Bloodrage 10-11, Shield Spec 12-16, Anticipation 17-21, Improved Revenge 22-24, Improved Thunder Clap 25-27, Last Stand 28, Master of Defense 29-30, Defiance 31-33, Improved Sunder Armor 34-36, Concussion Blow 37, Focused Rage 38-39, Shield Slam 40, Focused Rage 41, Bastion 42-46, IHS 47-49, Deflection 50-54, Improved Rend 55-57, Deep Wounds 58-60.

## Findings

None. No statement on the four pages or in the warrior card notes is false on 70755, and none contradicts the page it is on.

## Shoulds (wording, not findings)

- arms.md:125 (Levels 1 to 19) covers the points at 10 to 18 and skips the one at 19, which is the first point of Improved Tactical Mastery. The 20 to 60 block does not name ITM either. Suggested ending: "...and Improved Charge at 17 and 18 for 6 more Rage on every opener. Improved Tactical Mastery starts at 19, so more of your Rage survives a stance swap."
- arms.md:91 (Execute): "From there, Mortal Strike on cooldown and Execute on every other global". The Arms list keeps Reckless Swings on cooldown above Execute, and it works in Berserker Stance (no SpellShapeshift row, round 15 A2). Suggested: "From there, Mortal Strike and, in the raid build, Reckless Swings on cooldown, and Execute on every other global until the boss dies."
- fury.md:71 says the Reckless Swings health cost is something "the healers cover without noticing", while arms.md:87 says "so tell your healers". Both are true for their specs, but the two pages read differently about the same 25% cost. Optional: use the same wording on both.

## Checked and right

- **Round 10's fixes, re-derived from the client.**
  - arms.md:71: Improved Charge 2 is in the level 60 `build` only. Charge rank 3 is 15 rage (11578), so 2 points make it 21.
  - arms.md:74 and the Arms Bloodrage card [2]: Improved Bloodrage rank 1 is "by 25%" and rank 2 is "by 50%". Bloodrage 2687 is "10 rage ... and then an additional 10 rage over 10 sec".
  - arms.md:79, the raid build paragraph:
    - It drops Iron Will, Deflection and Improved Charge, and IHS goes from 1 to 3.
    - Spearing Strike 1310222 costs 150 (15 rage) on a 20000 ms cooldown. It does 75% weapon damage, plus 75% more against Giants and Dragonkin, which is twice the damage. Battle Stance only (round 15 A1).
    - Reckless Swings and Improved Execute sit in Fury tier 3, reached with 15 points.
    - Reckless Swings 1329555: "next 4 melee attacks generate 5 additional Rage. Lasts 8 sec". It costs `cost_pct` 25 of health (`cost_type` -2), has a 1000 ms global and a 15000 ms cooldown. Only white swings count (ProcTypeMask 4, round 15 row 3). That makes "two or three swings with a slow two-hander" right for a 3.3 to 3.8 speed weapon over 8 s.
    - Improved Execute rank 2 is "by 5", so Execute (150) costs 10 rage.
  - arms.md:87 and Arms cards [7] and [10]: the order (Reckless Swings after Mortal Strike, Spearing Strike after Slam) matches the list. The opener's "shout, Rend, Mortal Strike in three globals" still adds up without Improved Charge: 15 rage from Charge, minus 10 for the shout, plus 12.5 from Bloodrage, minus 10 for Rend, then about 17 from one landed 3.8-speed swing (engine `rage.go` 4.50 per speed for a two-hander), plus Bloodrage ticks and Anger Management, comes to about 30 by the third global.
  - fury.md:67 and :128 match the shipped order. Unbridled Wrath is 15-19, and Furious Precision starts at 20, with Dual Wield at 20 as decided in round 4.
  - fury.md:78: Improved Berserker Rage 2 is "instantly generate 10 Rage and has a 100% chance to remove all movement impairing effects". Improved Charge is what reaches Impale: IHS 3, IR 3, IC 2, ITM 3, AM 1 and DW 3 come to 15 Arms points.
  - fury.md:86 and :92 and Fury card [9]: Berserker Rage 18499 has a 30000 ms cooldown, trains at 30, and gives "immune to Fear and Incapacitate ... Lasts 10 sec". Its card sits right behind Reckless Swings under 80 rage, and the fear exceptions are kept.
  - fury.md:64: "Outside the execute phase" agrees with Fury card [5] (Execute above Bloodthirst).
  - index.md:34: Battle Shout 25289 reaches "all party or raid members within 20 yards ... Lasts 3 min". In 1.12 it reached party members only and lasted 2 min, so the bullet belongs in "What changed".
  - index.md:35: Spearing Strike is a Forever talent (1310222).
- **Every other line** re-read against the client:
  - Talent rank tables: Bloodthrill 20%, Improved Overpower 50%, Weaponmaster, Improved Rend 35%, Deep Wounds 60%, Impale 220%, 2H Spec 6%, ITM +15, Improved Slam, Flurry, DW Spec, Furious Precision, Enrage 2 = 4%, Improved Whirlwind, Death Wish, Bloodthirst 45% AP, Shield Spec, Master of Defense, Anticipation, Defiance, Bastion, Focused Rage, Improved Thunder Clap 2 a rank, Improved Shield Wall 11 min, Last Stand, Iron Will 15%.
  - Prerequisites: AM behind ITM 3, Deep Wounds behind Improved Rend 3, Impale behind Deep Wounds 3, MS behind Sweeping Strikes, Bloodthirst behind Death Wish, Master of Defense behind Shield Spec 5, Shield Slam behind Concussion Blow.
  - Trainer levels: Charge and Rend 4, Thunder Clap 6, Hamstring 8, Bloodrage, Defensive Stance and Taunt 10, Overpower and Shield Bash 12, Tactical Mastery, Revenge and Demoralizing Shout 14, Shield Block 16, Victory Rush 20, Berserker Stance 30, Whirlwind 36, Recklessness 50.
  - Spell text: Shield Block 75% / 7 s / 2 attacks, Shield Wall 60% / 12 s / 15 min, Demoralizing Shout 204 at 60 (round 4 and 8 decision), Thunder Clap 4 targets, Challenging Shout 6 s / 10 min, Whirlwind 4 targets with both weapons, Execute converts extra rage, Victory Rush 10% within 20 s, 30 s.
  - Racials in `spelltips/shared.json`: Elune's Light, Blood Fury, Berserking 15 s / 3 min, Eureka!, Touch of the Grave, Wind Blessed, Ruthless, Artisan's Precision, Stoneform, Endurance, Quickness.
  - Every warrior card note against its action and threshold.
  - Protection Levels 1 to 19 and 20 to 60 against the shipped order. Shield Slam is the 31st point at 40, and Focused Rage's last point is at 41.
- **Source links**, fetched this round. All return 200: the Blizzard recap, Wowhead Skyborne First Look, Power Up Gaming, Talents Forever data.json, both Warcraft Tavern Classic and Forever warrior pages, the Warcraft Tavern Fury guide, foreversixty.gg /sim/weights, /sim/drops, /planner and /bis/warrior/protection. Icy Veins returns 403 (unknown, not dead). Entries follow the round 4 source-link rule.
- No em or en dashes in the pages or card notes, no quoted nightly numbers, and no stat ranking in prose.

## Decisions

- **Zero confirmed findings for warrior this round.** Round 10's rewrites of the Arms raid build, the Fury build and the new card notes all hold against the 70755 client values. The three Shoulds are omissions or tone, not false statements.
- **Race lines stay as they are.** They were measured on the old builds, and the brief says not to file them unless they contradict a measured result or the client. Nothing contradicts them.

## Sims requested

- Still open, carried from round 10: race re-runs on the new Arms `raidBuild` and Fury `build`. Run raid 60 for every race on both factions, 10,000 iterations, seeds 1 and 50001. Lines affected: arms.md:109-111 and fury.md:112-114.
- PROVISIONAL: any prose-vs-BiS-row mismatch after the final re-rank. None found on this tree.
- Still open from earlier rounds: Dwarf against Human breakdown, Blood Fury held timing, Stoneform rotation line, and the Protection Improved Shield Wall build (Shield Wall is never cast in the tank sim).

## Sim findings

None new.

## Sources

- Client data 1.60.1.70755:
  - `talents/warrior.json`: rank tables and prerequisites.
  - `trainables/warrior.json`.
  - `spelltips/warrior.json`: 1329555, 1310222, 2687, 18499, 25289, 11578, 2565, 871, 11556, 20662, 402927 and others.
  - `spellconst/warrior.json`: 1329555 gcd 1000, cooldown 15000, duration 8000, cost_pct 25; 1310222 gcd 1500, cooldown 20000; 18499 gcd 1500, cooldown 30000.
  - `spelltips/shared.json` racials, `items.json` and `addon-builds.json`.
- Engine 9f7f1f6e4 `sim/core/rage.go` (normalized swing rage 4.50 per speed for a two-hander, 3.46 for a one-hander, crit x2).
- `design/reviews/2026-11-04-sim-audit-round15/warrior.md` (rows 2a, 2b, 3, A1, A2, F1).
- Commit eabba8f84 (round 10 warrior fixes).
