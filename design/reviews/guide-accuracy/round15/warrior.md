# Guide accuracy round 15: Warrior

Audited on the round 15 audit tree (main 4600b7d82: fourteen guide rounds, sim round 17, launch state), client 1.60.1.70755, engine pin b2c4a2d05 (`sim/enginever/version.go`). Files: `web/src/content/guides/warrior/{index,arms,fury,protection}.md` and every card note in `data/curated/apl/warrior-{arms,fury,protection}.json`. I read the round 1 to 14 warrior reports and fix logs, `round4/warrior-levels.md`, race-sims for rounds 2 to 5, the ledger, and the sim round 16 and 17 warrior reports (`design/reviews/2026-11-04-sim-audit-round1{6,7}/warrior.md`).

## Changes since round 14's tree (d7fc51fe7)

- `arms.md:111` (R14-1, R14-2) and `arms.md:115` (R14-3), `fury.md:118` (R14-4): round 14's own fixes. Re-checked below as new work.
- `sim/enginever/version.go`: f637ec32c to b2c4a2d05. Sim round 17 made no warrior engine commit; its before/after table shows every warrior value unchanged (Arms 820.4 / 826.4, Fury 879.4 / 873.3, Protection 1097.0 / 1088.0 threat/s).
- `addon-builds.json`: only rogue-subtlety and warlock-affliction orders changed. Every warrior order is unchanged.
- No warrior APL, BiS file, talent, trainable or race data changed.

## Builds decoded (unchanged, all legal on 70755)

Same strings as round 13 and 14, re-decoded against `talents/warrior.json` (tier, column order):

- Arms `build` `13323213032215201/0505/205`: 34 / 10 / 7 = 51.
- Arms `raidBuild` `30303213132215201/05050005012/1`: 32 / 18 / 1 = 51. Fury tier 3 (Reckless Swings, Improved Execute) reached with Cruelty 5, Unbridled Wrath 5, Enrage 5.
- Fury `build` `30323013002/050530025122100231/2`: 17 / 32 / 2 = 51.
- Protection `build` `35300003/0/250533120330001351`: 14 / 0 / 37 = 51.
- Prerequisites met: Anger Management (ITM 3), Deep Wounds (Improved Rend 3), Impale (Deep Wounds 3), Mortal Strike (Sweeping Strikes), Bloodthirst (Death Wish), Master of Defense (Shield Spec 5), Shield Slam (Concussion Blow).

Shipped orders (`addon-builds.json`, index 0 = level 10), re-mapped this round:

- **Arms leveling:** IHS 10, Deflection 11-13, Improved Rend 14-16, Improved Charge 17-18, ITM 19-21, Improved Overpower 22-23, AM 24, Deep Wounds 25-27, 2H Spec 28-30, Impale 31-32, Bloodthrill 33-34, Sweeping Strikes 35, Weaponmaster 36-39, Mortal Strike 40, Weaponmaster 41, Improved Slam 42-43, Cruelty 44-48, Unbridled Wrath 49-53, Improved Bloodrage 54-55, Iron Will 56-60.
- **Fury:** Cruelty 10-14, Unbridled Wrath 15-19, Furious Precision 20-22, Enrage 23-24, DW Spec 25-29, Death Wish 30, Reckless Swings 31, Improved Execute 32-33, Improved Berserker Rage 34-35, Improved Whirlwind 36-37, Flurry 38-39, Bloodthirst 40, Flurry 41, IHS 42-44, Improved Rend 45-47, Improved Charge 48-49, ITM 50-52, AM 53, Deep Wounds 54-56, Impale 57-58, Improved Bloodrage 59-60.
- **Protection:** Improved Bloodrage 10-11, Shield Spec 12-16, Anticipation 17-21, Improved Revenge 22-24, Improved Thunder Clap 25-27, Last Stand 28, Master of Defense 29-30, Defiance 31-33, Improved Sunder Armor 34-36, Concussion Blow 37, Focused Rage 38-39, Shield Slam 40, Focused Rage 41, Bastion 42-46, IHS 47-49, Deflection 50-54, Improved Rend 55-57, Deep Wounds 58-60.

Every leveling paragraph on the three pages matches these orders: arms.md:125 and :131 (Improved Overpower 22, Bloodthrill 33 after Deep Wounds and Impale), fury.md:67 and :128, protection.md:129 and :135 (Shield Slam the 31st point at 40, Focused Rage 41, then Bastion, IHS, Deflection, Improved Rend, Deep Wounds).

## Findings

None. No statement on the four pages or in the warrior card notes is false on 70755, and none contradicts its own page, the header build card, the cards, the live BiS rows, the stat tables, the race table or the shipped point order.

Counts: CONFIRMED 0, LIKELY 0, NEEDS-SIM 0, NEEDS-OWNER 0.

## Round 14 fixes re-checked

- **arms.md:111 (R14-1, R14-2).** "Dwarf comes out level with Night Elf, a hair ahead, since a Dwarf starts with more Strength, and Human's Ruthless ... and Gnome's Eureka! ... sit just behind them. On the Horde, Orc and Undead are close."
  - Sim round 16 A4 (raid 60, guide raid build): Dwarf 824.2, Night Elf 821.7, Human 820.2, Gnome 819.3; Orc 826.8, Undead 824.6, Troll 822.5, Tauren 818.2. The order the line states holds on both factions. Troll "trails Undead and Orc" and Tauren last hold.
  - The reason holds: engine b2c4a2d05 `sim/core/base_stats.go` RaceOffsets give Dwarf Strength +2 against Night Elf's -3 (the Classic base stats). The raid rows put Arms hit past the cap (`hit_to_cap` specials 0, hit weight 0, insignificant), so Artisan's Precision adds nothing and is rightly no longer named as Dwarf's case.
  - Night Elf / Undead stay the picks under the 1% race rule, as arms.md:109 and `recommendedRaces` say.
- **arms.md:115 (R14-3).** Frozen Heart of the Mountain (249469) is on all four Arms band 60 rows (bare and raid, both factions), source Enchanting. `items/warrior.json` bonding `pickup`; the Wowhead Forever tooltip (fetched this round) reads "Binds when picked up". Arcanite Champion 12790 (the raid main hand on both factions) is bonding `equip`, tooltip "Binds when equipped". Every other crafted piece on the Arms rows binds when equipped: Lionheart Helm 12640 (Blacksmithing), Timbermaw Tunic 252484, Scalegut Treaders 275618 and Might of the Timbermaw 19044 (Leatherworking). So Enchanting is the only profession the gear needs, as the line says.
- **fury.md:118 (R14-4).** The page shows the raid entry at 60 (`lib/bis/load.ts` selectBand: raid unless asked otherwise). The Alliance raid row wears Frozen Heart 249469 and Molten Heart 249470; the Horde raid row wears Frozen Heart only. Both hearts are bonding `pickup` (tooltip: "Binds when picked up"). The Blacksmithing pieces on every Fury band 60 row bind when equipped: Lionheart Helm 12640, Ironwood Blade 279259, Persuader 22384, Annihilator 12798. The line holds. (The Horde bare row wears Molten Heart instead of Frozen; that table is not the one the page opens, and "Molten Heart ... on the Alliance" is about the raid list, so it is noted, not filed.)
- **protection.md:119** (not changed, re-checked under round 14's rule): every crafted piece on the Protection band 60 rows is Blacksmithing and binds when equipped (Enchanted Thorium Helm 12620, Heavy Thorium Gauntlets 279264, Depleted Thorium Sabatons 279263, Azerothium Legplates 279265, Greenhammer 279261). No Enchanting or Engineering bind-on-pickup craft is on them, so "Blacksmithing makes plate and shields" stays true.

## Checked and right this round (re-derived)

- **Stat prose against the live band 60 tables:**
  - Arms raid: hit weight 0 (insignificant), specials cap met, so "once your gear reaches the cap it adds nothing" holds. The only two-handed skill item in `items/warrior.json` is Obsidian Edged Blade 18822 ("Increased Two-handed Swords +8"), so "for a two-hander it only comes on the weapon itself" holds.
  - Fury raid: hit weight 2.59, specials cap met, white cap 17.8 away, so "past the special cap every point still trims your white misses" holds.
  - Protection: expertise and the avoidance stats carry weight; block chance is insignificant. The prose explains each stat without ranking, and names no stat order. No ranking words anywhere in the three Stat priority sections.
- **Talent text** (`talents/warrior.json` rank tables): Enrage 2 = 30% / 4% / 12 s; Impale 2 = +20% critical strike damage bonus, so 220% of a normal hit (round 9 rule); Improved Overpower 2 = 50%; Deep Wounds 3 = 60% over 12 s; Improved Bloodrage 1 = 25%, 2 = 50%; Improved Shield Wall 2 = 11 min, so 15 to 4; Improved Thunder Clap 2 a rank, so two points out costs 4 rage; Weaponmaster 5 = 5% / 15% / 5%; Improved Slam 2 = -5 s, -0.5 s cast and global, swing kept; Improved Whirlwind 2 = -2 s and 30% to the closest target; Improved Berserker Rage 2 = 10 rage and slow removal; Focused Rage 3; Defiance 15%; Bastion 10%; Master of Defense 5 rage on dodge and parry with a shield; Shield Spec 5 = 5% and 5 rage on every block.
- **Spell text** (`spelltips/warrior.json`, re-read): Charge rank 3 15 rage, out of combat only; Bloodrage 10 + 10 over 10 s, 60 s; Battle Shout "all party or raid members within 20 yards", 3 min; Victory Rush 10% heal, 20 s window, 30 s; Shield Block 75% / 7 s / 2 attacks / 5 s; Shield Wall 60% / 12 s / 15 min; Last Stand 30% / 20 s; Challenging Shout 6 s / 10 min; Thunder Clap 20% slow, 4 targets; Whirlwind 4 targets, both weapons, 10 s; Sweeping Strikes 5 attacks; Death Wish 10 rage, 20%, 5% damage taken, 30 s, 3 min; Recklessness 100% crit, 15 s, 30 min; Berserker Rage 10 s, fear immunity; Reckless Swings 15 s, 4 attacks, 5 rage, 8 s, 25% base health; Spearing Strike 15 rage, 20 s, 75% plus 75% against giants and dragonkin; Execute 15 rage, converts extra rage; Tactical Mastery 10 rage; Demoralizing Shout 196 at 54, 204 at 60 (rounds 4 and 8).
- **Trainer levels** (`trainables/warrior.json`): Charge and Rend 4, Thunder Clap 6, Hamstring 8, Bloodrage, Defensive Stance, Sunder Armor and Taunt 10, Overpower and Shield Bash 12, Revenge, Demoralizing Shout and Tactical Mastery 14, Shield Block 16, Victory Rush 20, Berserker Stance and Berserker Rage 30, Whirlwind 36, Recklessness 50.
- **Card notes** against their actions and conditions (all three APLs): Arms Berserker swap at rage 25 or less after the first Execute, Bloodrage under 80, Hamstring only with Windfury and never from 25%, Heroic Strike at 90; Fury Death Wish at 210 s or more remaining, in the execute phase or the last 30 s, Recklessness in the execute phase or the last 15 s, Heroic Strike from 60 and not from 25% once Execute is known; Protection Last Stand under 35%, Shield Wall under 40%, Bloodrage under 60, Thunder Clap and Demoralizing Shout refreshed under 3 s, Heroic Strike from 40. Nothing false.
- **Racials** against `races.json` and the race table: Elune's Light 10% crit / 15 s / 3 min; Touch of the Grave drain, once every 15 s; Wind Blessed 2% haste; Blood Fury 10% attack power / 15 s / 2 min; Berserking 10% attack speed / 15 s / 3 min; Ruthless 1% crit; Artisan's Precision 1% hit; Eureka! (warrior row) 10% less rage and 10% more damage on non-periodic damaging abilities, 8 s / 2 min; Quickness 1% dodge and 2% movement; Endurance 5% health; Stoneform. `combos.json`: all eight original races and both Skyborne can be warriors (index.md:39). Skyborne is named only as the paid option.
- **Fury race order** against sim round 16 F3 (Human 880.6, Night Elf 880.1, Gnome 878.0, Dwarf 876.5; Undead 878.1, Orc 875.6, Troll 872.8, Tauren 866.8): every ordering fury.md:114 states holds (Human over Night Elf is 0.5, inside the error).
- **New source since round 14:** sim round 17 cites Icy Veins Forever (Abide, Fury stat priority, 2026-11-03): queuing Heroic Strike no longer removes the extra off-hand miss chance in Forever. No warrior page claims it does: fury.md:88 says only that Heroic Strike replaces the next main-hand swing off the global cooldown, and fury.md:100 treats white-swing hit on its own.
- **Source links**, fetched this round: all return 200 except the Icy Veins racial rework, which returns 403 (unknown, not dead, per the brief). Every entry follows the round 4 source-link rule: /sim/weights and /sim/drops only on the damage pages, Protection links its own best in slot page, no /sources link, no "how the spec plays" entry, Talents Forever labelled as beta client data.
- **Voice:** no em or en dashes, no quoted nightly numbers, no site or sim talk outside the planner links.

## Shoulds (wording, not findings)

- **fury.md:74 against fury.md:124.** The talent line says "Fury never casts Rend", while Levels 1 to 19 on the same page says to put Rend on anything that will live through it. The talent line speaks of the 60 build (Improved Rend at 45 to 47, when Fury lives in Berserker Stance, where Rend cannot be cast), so it is not false, but a reader can trip on it. Suggested: "with Improved Rend as the toll to reach Deep Wounds; by the time you take it you live in Berserker Stance and no longer Rend".
- **arms.md:115.** "Blacksmithing and Mining are the other pair, but Arcanite Champion binds when equipped, so buy one rather than train it for the sword" offers the pair and talks the reader out of it in the same sentence. Suggested: "Blacksmithing and Mining are the classic pair, but Arcanite Champion binds when equipped, so you can buy the sword without them."
- **Round 14 fix log, R14-1** says "No Strength reason added", but the shipped line carries "since a Dwarf starts with more Strength". The reason is true (above), so only the log is off.
- Carried from rounds 13 and 14: the Arms Bloodrage card [2] trim, arms.md:125 skipping the level 19 point (ITM), arms.md:91 leaving out Reckless Swings, the Reckless Swings health-cost wording difference, fury.md:54 "debuff", and fury.md:100 Edgemaster's wording.

## Decisions

- **Race picks stay** (Night Elf / Undead on Arms and Fury, Night Elf / Tauren on Protection). Every measured DPS gap is inside 1%.
- **Professions lines stand.** They follow the raid rows the page opens. The Fury Horde bare row's Molten Heart is not on the default table and is a BiS-side mix the post-launch re-rank owns.
- **BiS race captions** and the Fury BiS row's older talent string are not filed, per the brief (data lane).
- **Skyborne lines** (arms.md:109, fury.md:112) are unmeasured on the current builds and contradict no measured result or client value, so they are not filed (round 11 rule).

## Sims requested

- Carried from round 14, still not run: **Arms and Fury race order on the guide raid builds**, raid 60, every race on both factions including both Skyborne, 10,000 iterations, seeds 1 and 50001. This confirms the Dwarf and Orc order at arms.md:111 and the Skyborne lines at arms.md:109 ("edges both by a hair on either faction") and fury.md:112 ("edges Night Elf by a hair", "On the Horde, Undead stays ahead of it").
- Carried: **Protection races on the current build**, tank fight, every race, seeds 1 and 50001, reporting damage taken, TMI and threat. protection.md:113 ("takes less damage than any other Alliance race", "its health swings less than a Dwarf's", "no race's health swings less") rests on round 3 numbers from an older build.
- Carried: Protection Improved Shield Wall build, the Dwarf against Human breakdown, the Stoneform rotation line. (The Blood Fury held timing lead is closed by sim round 17 A6: on cooldown, as the prose says.)

## Sim findings

None new. Sim round 17 seed 1 (Heroic Strike counting as an auto attack for Unbridled Wrath, Flurry and Reckless Swings) stays open with the sim lane until a level 60 Forever log; no page claims either way. fury.md:71 and :86 say "white swings" for Reckless Swings, which is the client's proc mask (done melee auto attack) whichever way that question settles.

## Sources

- **Client data 1.60.1.70755:** `talents/warrior.json`, `addon-builds.json`, `trainables/warrior.json`, `spelltips/warrior.json`, `items/warrior.json` (bonding for 249469, 249470, 12790, 12640, 252484, 275618, 19044, 279259, 22384, 12798, 12620, 279261, 279263, 279264, 279265, 18822), `bis/warrior-{arms,fury,protection}.json` (band 60 rows, weights, `hit_to_cap`), `races.json`, `combos.json`.
- **Engine b2c4a2d05:** `sim/core/base_stats.go` RaceOffsets.
- **Site code:** `web/src/lib/bis/load.ts` (raid entry at 60), `web/src/lib/guides/race-table.ts`.
- **Race measurements:** `design/reviews/2026-11-04-sim-audit-round16/warrior.md` A4 and F3; `design/reviews/2026-11-04-sim-audit-round17/warrior.md` (before/after, A6, seed 1, 1b).
- **Wowhead Forever tooltips** (fetched 2026-11-04):
  - Frozen Heart of the Mountain: https://nether.wowhead.com/forever/tooltip/item/249469
  - Molten Heart of the Mountain: https://nether.wowhead.com/forever/tooltip/item/249470
  - Arcanite Champion: https://nether.wowhead.com/forever/tooltip/item/12790
- **Icy Veins Forever**, Fury Warrior Stat Priority (Abide, 2026-11-03), as cited by sim round 17: https://www.icy-veins.com/wow-forever/fury-warrior-melee-dps-pve-stat-priority
