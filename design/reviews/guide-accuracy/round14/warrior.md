# Guide accuracy round 14: Warrior

Audited on the round 14 audit tree (main d7fc51fe7: thirteen guide rounds, sim round 16 merged), client 1.60.1.70755, engine pin f637ec32c (`sim/enginever/version.go`). Files: `web/src/content/guides/warrior/{index,arms,fury,protection}.md` and the warrior card notes in `data/curated/apl/warrior-{arms,fury,protection}.json`. I read the round 1 to 13 warrior reports, `round4/warrior-levels.md`, race-sims for rounds 2 to 5, the ledger, and the sim round 16 warrior report (`design/reviews/2026-11-04-sim-audit-round16/warrior.md`).

## Changes since round 13's tree (604b78a80)

- `fury.md:78`: Improved Charge is now "part of the points the rows above Deep Wounds ask for". This is checked below and holds.
- `data/curated/coach/warrior-arms.md`: this is coach text, not a guide page. It now matches arms.md:87.
- `addon-builds.json`: only warlock-demonology changed. Every warrior order is unchanged.
- No warrior APL, BiS file or client data changed.
- New evidence in the tree: sim round 16 measured every race on the Arms guide raid build and the Fury guide build (A4 and F3 in its report). Round 13 could not have that. Under round 14's rule, the BiS rows and race results are now final.

## Builds decoded (unchanged, all legal on 70755)

Same as round 13: Arms `build` and `raidBuild`, Fury `build`, and Protection `build`. Tier gates, re-read from `talents/warrior.json` (tier index, 5 points a tier):

- The Arms tree on 70755 puts Deep Wounds at tier 2, which needs 10 points in tiers 0 and 1. Impale is at tier 3 and needs 15.
- Fury's Arms points give tiers 0 and 1 a total of 11: IHS 3, Improved Rend 3, Improved Charge 2 and ITM 3. Without Improved Charge that falls to 9, and Deep Wounds is locked. With AM 1 and Deep Wounds 3 the total reaches the 15 Impale needs.
- So the new fury.md:78 wording is true.

Shipped orders (`addon-builds.json`, index 0 = level 10), re-decoded:

- **Arms leveling:** IHS 10, Deflection 11-13, Improved Rend 14-16, Improved Charge 17-18, ITM 19-21, Improved Overpower 22-23, AM 24, Deep Wounds 25-27, 2H Spec 28-30, Impale 31-32, Bloodthrill 33-34, Sweeping Strikes 35, Mortal Strike 40.
- **Fury:** Cruelty 10-14, Unbridled Wrath 15-19, Furious Precision 20-22, Death Wish 30, Bloodthirst 40.
- **Protection:** Improved Bloodrage 10-11, Shield Spec 12-16, Anticipation 17-21, Concussion Blow 37, Shield Slam 40, the last point of Focused Rage 41, Bastion 42-46, IHS 47-49, Deflection 50-54, Improved Rend 55-57, Deep Wounds 58-60.

Every leveling paragraph on the three pages matches these orders. Sim round 16 also fixed the level 41 ladder for Protection. The order is still Shield Slam at 40 and Focused Rage at 41, as protection.md:135 says.

## Findings

| id | status | kind | file:line | claim | evidence | fix |
|---|---|---|---|---|---|---|
| WARRIOR-R14-1 | CONFIRMED | race | arms.md:111 | "On the Alliance, Human's Ruthless ... and Gnome's Eureka! ... sit just behind Night Elf, and Dwarf's Artisan's Precision, 1% hit, trails them." | The sim round 16 report (A4) measured the Arms guide raid build at raid 60 (3000 iterations, seed 1): **Dwarf 824.2**, Night Elf 821.7, Human 820.2, Gnome 819.3. Dwarf leads every Alliance race. It is 4.0 DPS over Human and 4.9 over Gnome, well beyond the error at that count (about 0.8 a run). The reason is on the client side, not the racial. Engine `sim/core/base_stats.go` RaceOffsets give Dwarf Strength +2 and Night Elf -3, against Human 0 and Gnome -5. The raid BiS rows show hit at 0 for Arms (gear at the cap), so Artisan's Precision adds nothing, and Strength decides it. Under the race rule (a pick moves only beyond 1%), Night Elf stays the pick. The order the line states is still false. | "If you want a different look, the gap is small. On the Alliance, Dwarf is level with Night Elf: Artisan's Precision, 1% hit, does nothing once your gear reaches the hit cap, but a Dwarf starts with more Strength than any other Alliance race. Human's Ruthless, a flat 1% crit, and Gnome's Eureka!, 10% more damage and 10% less rage on your non-periodic damaging abilities for 8 seconds every 2 minutes, sit just behind." |
| WARRIOR-R14-2 | LIKELY | race | arms.md:111 | "On the Horde, Orc is next" (after Undead, the pick) | Sim round 16 A4, same run: **Orc 826.8**, Undead 824.6, Troll 822.5, Tauren 818.2. Orc is above Undead by 2.2 DPS, about 2 standard errors on a single seed, so this needs the two-seed run to confirm. Orc also has Strength +3 against Undead's -1 (RaceOffsets). The pick stays Undead under the 1% rule, but "next" puts Orc behind it. Troll and Tauren hold. | "On the Horde, Orc is level with Undead: Blood Fury is 10% attack power for 15 seconds every 2 minutes, so press it on the pull and every time it comes back, and it still adds damage inside Recklessness." (The rest of the line stays.) |
| WARRIOR-R14-3 | CONFIRMED | gear/professions | arms.md:115 | "Blacksmithing and Mining. Blacksmithing makes Arcanite Champion, and Mining feeds it. Engineering is the other pick" | The page's BiS carries **Frozen Heart of the Mountain** (249469, source Enchanting) in the raid 60 rows on both factions (Alliance trinket1, Horde trinket2), and in the bare rows too. `items/warrior.json` 249469 says `"bonding": "pickup"`, and the Wowhead Forever tooltip (nether.wowhead.com/forever/tooltip/item/249469) says "Binds when picked up". A bind-on-pickup craft can only be made by an enchanter. Every Blacksmithing piece in the Arms rows binds when equipped, including Arcanite Champion 12790 and Lionheart Helm 12640, so they can be bought. The professions line steers a warrior away from the one profession the page's own gear needs. The rogue Combat page already says this about the same trinkets (rogue/combat.md:116). | "Enchanting is the one that changes your gear: Frozen Heart of the Mountain is a bind-on-pickup trinket only an enchanter can make. Blacksmithing and Mining are the other pair, but Arcanite Champion binds when equipped, so buy one rather than train it for the sword. Engineering brings Goblin Sapper Charges and Dense Dynamite for every trash pack, and grenades in PvP." |
| WARRIOR-R14-4 | CONFIRMED | gear/professions | fury.md:118 | "Blacksmithing and Mining. Blacksmithing makes melee weapons and armor, and Mining feeds it. Engineering is the other pick" | The Fury raid 60 BiS rows carry **Frozen Heart of the Mountain** (249469) on both factions. The Alliance row also carries **Molten Heart of the Mountain** (249470). Both are Enchanting crafts with `"bonding": "pickup"` (`items/warrior.json`). The Blacksmithing pieces on the page all bind when equipped: Ironwood Blade 279259, Persuader 22384, Annihilator 12798 and Lionheart Helm 12640. | "Enchanting is the one that changes your gear: Molten Heart of the Mountain and Frozen Heart of the Mountain are bind-on-pickup trinkets only an enchanter can make. Blacksmithing makes melee weapons and armor, but the ones on your list bind when equipped, so buy them rather than train it. Engineering brings Goblin Sapper Charges and Dense Dynamite for every trash pack, and grenades in PvP." |

Counts: CONFIRMED 3, LIKELY 1, NEEDS-SIM 0, NEEDS-OWNER 0.

## Shoulds (wording, not findings)

- fury.md:100: "Weapon skill from gear, Edgemaster's Handguards and the like, is worth chasing." The page's hand rows pick Theramore Gauntlets (Alliance) and Raider Gauntlets (Horde). Edgemaster's is not even among the alternatives, and its stats are `client_unconfirmed` classic-db values that carry no stat key. The line explains, not ranks, so it is not a finding. A softer wording would be: "Weapon skill from gear, Edgemaster's Handguards and the like, takes the sting out of glancing blows and trims your misses, and no race brings it now."
- Carried from round 13: the Arms Bloodrage card [2] trim, arms.md:125 skipping the level 19 point, arms.md:91 leaving out Reckless Swings, the Reckless Swings health-cost wording, and fury.md:54 "debuff".

## Checked and right this round

- **Gear and stat prose against the live BiS rows (band 60, bare and raid, both factions):**
  - Arms: Arcanite Champion is the raid main hand on both factions, from Blacksmithing.
  - Arms: Obsidian Edged Blade 18822 is "Increased Two-handed Swords +8". Edgemaster's 14551 is "Axes, Daggers, Swords +7", all one-handed skills.
  - Arms raid hit shows 0. That fits "once your gear reaches the cap it adds nothing".
  - Fury raid hit is above 0. That fits "past the special cap every point still trims your white misses".
  - Protection: every Blacksmithing piece binds when equipped, and "Blacksmithing makes plate and shields" stays true.
  - Protection expertise: "only a few late sets carry" holds. Waywatcher Gauntlets 240081 is ilvl 98 set gear (Wowhead Forever tooltip), and fa02e9c6d names 71 catalogue items.
- **Race table rows** (`race-table.ts` melee and tank picks) against the prose:
  - Tauren shows Endurance only, so "brings nothing for damage" holds.
  - Protection shows Night Elf Elune's Light + Quickness, Dwarf Stoneform + Artisan's Precision, and Tauren Endurance. That matches protection.md:113-115.
- **Fury race lines against sim round 16 F3** (Human 880.6, Night Elf 880.1, Gnome 878.0, Dwarf 876.5; Undead 878.1, Orc 875.6, Troll 872.8, Tauren 866.8):
  - Human over Night Elf is 0.5, inside the error.
  - Every other ordering the prose states holds.
- **Trainer levels:**
  - Victory Rush 20, Berserker Stance 30, Intercept 30, Cleave 20, Challenging Shout 26, all from `trainables/warrior.json`.
  - Dual Wield at 20, per the round 4 levels note.
- **Consumables named** are all in `simconsumes.json`: Mighty Rage Potion, Elixir of the Mongoose, R.O.I.D.S., Winterfall Firewater and Smoked Desert Dumplings.
- **Source links**, fetched this round:
  - All return 200 except Icy Veins, which returns 403 (unknown, not dead).
  - Every entry follows the round 4 source-link rule.
- **Round 13's full sheet** (stances, costs, cooldowns, rank values, racial tooltips) still applies. No client data changed.

## Decisions

- **Race picks stay** (Night Elf / Undead on both DPS pages): every measured gap is inside 1%. Only the order the Arms prose states changes.
- **The professions lines follow the page's BiS.** The bind-on-pickup Enchanting trinkets are a stable reason, the same rule the rogue page uses. Arcanite Champion and the Fury blacksmith weapons are named only as buy, not train.
- **BiS race captions** (Human / Orc / Troll / Dwarf) are not filed, per the brief.

## Sims requested

- **Arms race order on the guide raid build:** raid 60, every race on both factions including both Skyborne, 10,000 iterations, seeds 1 and 50001. This confirms R14-1 and R14-2 and the line "Skyborne edges both by a hair on either faction" (arms.md:109).
- **Fury, the same run.** It settles Human against Night Elf and the Skyborne lines at fury.md:112 ("edges Night Elf by a hair", "On the Horde, Undead stays ahead of it").
- **Protection races on the current build**, tank fight. The round 3 numbers were measured on an older build.
- Carried: Protection Improved Shield Wall build, the Dwarf against Human breakdown, Blood Fury held timing, and the Stoneform rotation line.

## Sim findings

None new. Sim round 16's F5 (whether Heroic Strike counts as a white swing for Unbridled Wrath) stays open with the sim lane. No page claims it.

## Sources

- **Client data 1.60.1.70755:**
  - `talents/warrior.json` (tiers and columns)
  - `addon-builds.json`
  - `trainables/warrior.json`
  - `items/warrior.json` (bonding for 249469, 249470, 12790, 279259, 22384, 12798, 12640, 18822 and 14551)
  - `bis/warrior-{arms,fury,protection}.json` (band 60 rows)
  - `simconsumes.json`
- **Engine f637ec32c:** `sim/core/base_stats.go` RaceOffsets.
- **Race measurements:** `design/reviews/2026-11-04-sim-audit-round16/warrior.md`, A4 and F3.
- **Wowhead Forever tooltips:**
  - Frozen Heart of the Mountain: https://nether.wowhead.com/forever/tooltip/item/249469
  - Waywatcher Gauntlets: https://nether.wowhead.com/forever/tooltip/item/240081
- **Precedent:** `web/src/content/guides/rogue/combat.md:116`.

## Fix log (round 14)

- WARRIOR-R14-1: fixed, arms.md:111. Order only: Dwarf "comes out a hair ahead of Night Elf", Human and Gnome "just behind them". No Strength reason added; Night Elf stays the pick.
- WARRIOR-R14-2 (LIKELY): fixed, arms.md:111, "Orc and Undead are close", true whichever the two-seed run shows. Troll "trails Undead and Orc" still holds.
- WARRIOR-R14-3: fixed, arms.md:115. Frozen Heart (BoP Enchanting) is on every band 60 row; Arcanite Champion binds on equip.
- WARRIOR-R14-4: fixed, fury.md:118. Worded as "Frozen Heart, and on the Alliance Molten Heart beside it": the Alliance raid row wears both and the Horde raid row wears Frozen only.
- SHOULDs: not applied.

All files were already at updated: 2026-11-04. The prose and sections tests pass (145/145).
