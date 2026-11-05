# Guide accuracy round 20: Hunter

Audited on the round 20 audit tree (4234643f5: main after guide round 19 and sim round 20, engine pin d5c708834, source read from gr20/engine), against the active client 1.60.1.70790 (`data/builds/1.60.1.70790/`). Files: `web/src/content/guides/hunter/{index,beast-mastery,marksmanship,survival}.md`, every note in `data/curated/apl/hunter-{beast-mastery,marksmanship,survival}.json`, every rendered row of `addon-data.json` `rotations/hunter-*` (BM 18 bands, MM 19, SV 20, read band by band), the six orders in `addon-builds.json`, the hunter rows of `data/curated/rank-choices.json`, and `talents/`, `spelltips/`, `spellconst/hunter.json`, `gametables/basemp.txt`, `levels.json`, `items/hunter.json`, `bis/hunter-*.json`. Read first: round 19 hunter report, sim round 20 `hunter.md` and `integration.md`, the ledger's Round 19 section, and the earlier briefs.

Where a note renders: the addon row and BiS "Play it" panel carry the first sentence at every band; the guide card prints its first clause at the level 60 ranks; the sim drawer prints the whole note at any simmed level (bare 60 included). Every sentence is therefore read at every band.

## Changes since round 19's tree (6beb919c2)

- Card notes (sim round 20 plus round 19 fixes re-applied at merge): BM Rapid Fire (R19-1 fix), MM Rapid Fire (R19-2 fix), BM and MM Multi-Shot (R19-3 parenthetical plus the new quarter-mana gate and "worst damage for its mana"), BM and MM Serpent Sting and Arcane Shot (ranged attack power shares, damage-per-mana ranking; MM's first sentences now say "the shared cooldown" where they said "Multi-Shot"), SV Immolation Trap (attack power sentence). APL: Multi-Shot gated `currentManaPercent > 25%`; Arcane Shot's 25% gate removed.
- Prose: beast-mastery.md:61, 91, 99, 113 and marksmanship.md:96, 103, 105, 119.
- Orders and point counts unchanged for hunter (the addon-builds diff touches mage and shaman only). Rank rows unchanged for hunter.

## Shipped orders (addon-builds.json, index 0 = level 10)

- BM leveling: Deadly Aspects 10-14, Focused Fire 15-16, Endurance Training 17-20, Bestial Swiftness 21, Unleashed Fury 22-26, Ferocity 27-31, Intimidation 32, Bestial Discipline 33-34, Frenzy 35-39, Bestial Wrath 40, Lethal Attacks 41-45, Improved Stings 46-47, Efficiency 48-51, Careful Aim 52-56, Rapid Killing 57, Trueshot Aura 58, Mortal Shots 59-60. Raid: same to 47, Efficiency 48-50, Careful Aim 51-55, Trueshot Aura 56, Mortal Shots 57-60.
- MM leveling: Lethal Attacks 10-14, Efficiency 15-19, Careful Aim 20-24, Rapid Killing 25-26, Trueshot Aura 27, Mortal Shots 28-32, Rapid Recuperation 33-34, Barrage 35, Ranged Weapon Specialization 36-40, Deadly Aspects 41-45, Endurance Training 46-48, Focused Fire 49-50, Unleashed Fury 51-55, Improved Tracking 56-60. Raid: Lethal Attacks 10-14, Improved Stings 15-16, Efficiency 17-21, Careful Aim 22-26, Lone Wolf 27, Trueshot Aura 28, Mortal Shots 29-33, Barrage 34-36, Ranged Weapon Specialization 37-39 and 41-42, Sniper Shot 40, Deadly Aspects 43-47, Improved Tracking 48-52, Survivalist 53-57, Surefooted 58-60.
- SV (both): Improved Tracking 10-14, Savage Strikes 15-16, Survivalist 17-19, Clever Traps 20-21, Surefooted 22-24, Deterrence 25, Predator's Edge 26-29 and 31, Strider Kick 30, Counterattack 32, Resourcefulness 33-34, Expose Prey 35-36, Survivalist's Discipline 37-38, Lightning Reflexes 39 and 41-44, Lacerating Strikes 40, Deadly Aspects 45-49, Lethal Attacks 50-54, Careful Aim 55-59, Lone Wolf 60.

Rendered Aimed Shot rows: rank 1 (19434) at BM 20-58 and MM 20-48, 52, 58; rank 4 (20902) at MM 50; rank 6 (20904) at 60. Rank rows also cast rank 1 at bare 60 (sim drawer).

## Findings

| id | status | kind | file | wrong claim | evidence | fix |
|---|---|---|---|---|---|---|
| HUNTER-R20-1 | CONFIRMED | card note (value/ranking) | `hunter-beast-mastery.json` Arcane Shot (priorityList[7]) and `hunter-marksmanship.json` Arcane Shot (priorityList[7]); later sentence, rendered in the sim drawer at every band | "every hit carries 15% of your ranged attack power, so it is second only to Serpent Sting for damage per mana" | The same cards cast rank 1 Aimed Shot from 20 to 58 (rank-choices rows, `addon-data` rows above). Rank 1 is normalized weapon damage + ammo + 40 for a flat 75 mana (`spellconst` 19434: cost 75, effect 121 amount 40; `aimed_shot.go` FlatCost 75, `CalculateNormalizedWeaponDamage(RAP)`, i.e. 0.2 damage per point of ranged attack power). Arcane Shot is base + 15% RAP for 80/105/135/160 mana (14283 to 14286). On each band's own bare best in slot gear (horde, `bis/hunter-*.json` slots summed with `levels.json` base stats, Heart of the Lion 40+4/level, Aspect of the Hawk and Hunter's Mark by rank, Careful Aim on its order), raw damage per mana before armor: BM 40 Aimed r1 4.4 against Arcane 2.5, BM 50 4.4 against 2.6; MM 40 4.6 against 2.6, MM 50 (rank 4 row) Arcane ahead. Arcane wins only if the target's armor removes more than 42 to 44% of physical damage at 40 to 50 (59% at bare 60); a same-level mob's armor removes about 33 to 38% (Classic armor formula). Below about 30 the two are even or Arcane leads. So on BM bands from the mid 30s to 58, MM 34 to 48, 52 and 58, and bare 60, rank 1 Aimed Shot does more for its mana than Arcane Shot, and the note's ranking is false. The page's own leveling paragraphs (BM:157, MM:163) keep rank 1 precisely because it is cheap damage. At band 60 (rank 6, 310 mana) the ranking holds (sim round 20 B3: Sting 6.4, Arcane 4.5, Aimed 3.8, Multi 3.1 on BM). | Both cards: "Arcane Shot to fill a gap, and from 18 only when Multi-Shot is more than a global away. It is instant and never holds an auto shot, and every hit carries 15% of your ranged attack power, so it does far more for its mana than Multi-Shot: keep it going even when the mana runs low." (MM keeps its first sentence as decided under the Shoulds.) |
| HUNTER-R20-2 | CONFIRMED | card note (contradiction) | `hunter-beast-mastery.json` Aimed Shot (priorityList[4]); second sentence, rendered in the sim drawer at bands 20 to 60 | "For Beast Mastery it is only a small extra over Multi-Shot, and Multi-Shot on every cooldown is fine." | Round 20 gated Multi-Shot at a quarter of your mana (APL `currentManaPercent > 25%`), and the Multi-Shot note on the same card now says "while you are above a quarter of your mana ... it is the first shot to go when the mana runs low". The page says the same (BM:61 weakness "if you keep firing Multi-Shot on every cooldown"; BM:91 "Multi-Shot on cooldown is fine while your mana is healthy"; BM:99 "hold Multi-Shot first"). This sentence was not touched in round 20 and now contradicts both the card and the page; the wow-player MUST on BM:91 in sim round 20 was this exact contradiction in the prose, and the note kept it. | "Aimed Shot when the shared cooldown comes up right after an auto shot fires. For Beast Mastery it is only a small extra over Multi-Shot, so Multi-Shot on that cooldown is fine while you are above a quarter of your mana. If you are partway to your next auto shot, take Multi-Shot and keep the auto rolling." |
| HUNTER-R20-3 | LIKELY | card note (value/ranking) | `hunter-marksmanship.json` Serpent Sting (priorityList[6]); later sentence, sim drawer | "Every tick carries 5% of your ranged attack power on top of its own damage, so it is the best damage for its mana you fire" | Same arithmetic as R20-1 against Serpent Sting rank 7 and 8 (190 and 230 mana, base 75 and 100 a tick + 5% RAP). Before talents, rank 1 Aimed Shot beats the sting only when armor removes less than 21 to 24% at 50 (so the sting wins there) but less than 42% at bare 60. The Marksmanship leveling build adds Barrage (1 point, 5%), Ranged Weapon Specialization (5%) and Mortal Shots 5/5 (+30% crit bonus) to Aimed Shot and nothing to the sting (no Improved Stings on that order), which lifts the break-even to about 45% at 58 and bare 60, above a level 58 to 63 target's 38 to 40%. So at MM 58 and bare 60 the rank 1 Aimed Shot the rows cast very likely out-does the sting per mana. On Beast Mastery Improved Stings (46-47) keeps the sting ahead. Gear dependent, so LIKELY; sim 1 settles it. | "Every tick carries 5% of your ranged attack power on top of its own damage, so it does more for its mana than Arcane Shot or Multi-Shot: keep it rolling for the whole fight." (True at every band; apply to the BM card too for one wording.) |

Counts: CONFIRMED 2, LIKELY 1, NEEDS-SIM 0, NEEDS-OWNER 0.

## Shoulds (wording, not findings)

- **MM Serpent Sting and Arcane Shot first sentences at band 18 and 19:** round 20 changed "Multi-Shot" to "the shared cooldown". Nothing shares it before Aimed Shot at 20, but the client's own Multi-Shot tooltip says "Multi-Shot shares its cooldown with Aimed Shot", so a level 18 reader can follow it; not filed (unlike R19-3, the sentence gates on no spell the player lacks). "Multi-Shot" is what the APL condition reads (`spellTimeToReady 2643`) and is exact at every band; I would keep BM's wording and put MM back to it, against integration's Should to move BM the other way.
- **BM Serpent Sting, "When both are ready, Multi-Shot goes first":** under a quarter of your mana Multi-Shot is held and does not go at all; "while your mana is above a quarter" would close it.
- **SV Mongoose Bite:** "your hits on a target with Hunter's Mark ... open it for 10 seconds" leaves out that it is a 5% (35) then 10% (36) chance per hit; survival.md:72 says "a 10% chance". "give each of your hits a 10% chance to open it" matches the talent text.
- **SV Immolation Trap:** "a share of your attack power" is ranged attack power in the engine and the wiki's 2% a tick; Blizzard's 16 October note says only "Attack Power". Strength raises a Survival hunter's melee attack power but not the trap if it is ranged; "ranged attack power" would be exact once a log confirms it.
- **Sources:** both ranged pages now quote 15% and 5% from Blizzard's 3 November notes; neither lists that source. Consider adding it (see Sources below).
- **Carried from round 19:** BM Serpent Sting "pushes back the shared cooldown" at 18; SV Raptor Strike "never hold it back for mana" beside the 46-47 rank 5 row; SV Counterattack row at 30 before the talent at 32; survival.md:140, :124, :146; the rank-choices round tag; MM:159 "most hard pulls"; SV:65; "throws away a tick"; "the Molten Core tome"; BM:81 Deflection; SV:118; BM:157; SV:80. Integration's Shoulds (Rapid Fire row against paragraph wording, cooldown-only rows, Sniper raid-only note) also stand.

## Checked and right this round

- **Round 19's three fixes:** BM Rapid Fire (Bestial Wrath "from 40", Rapid Killing at 57 makes four minutes at band 58, five at 60 raid) and MM Rapid Fire (Rapid Killing 25 and 26 before Rapid Fire at 26, three minutes on all leveling bands, five on the raid build) hold at every band; the Multi-Shot parenthetical covers band 18.
- **Round 20 values:** Arcane Shot +15% RAP and Serpent Sting +5% RAP a tick (Blizzard, 3 November; forever-xp.com quotes both lines verbatim); fire traps scale with attack power (Blizzard, 16 October). `attack_power_scaling.go` matches. Multi-Shot is the worst damage per mana at every band it renders: 13.9% of base mana (`basemp.txt` hunter: 54 mana at 18, 154 at 40, 239 at 60) for normalized weapon damage with no flat bonus, below rank 1 Aimed, Arcane, the sting and Sniper Shot at every band computed (18, 30, 40, 50, 60). Sting above Arcane at every band (bands 10 to 60).
- **Prose:** BM:61, 91, 99, 113 and MM:96, 103, 105, 119 read in the raid context of their sections (rank 6 Aimed Shot), where sting, then Arcane, then Sniper Shot, Aimed Shot and Multi-Shot is the damage-per-mana order; no stat is ranked. Every other line is unchanged since round 19 and its checks stand; the class page needs no change.
- **Note sweep:** every row below.

## Decisions

- **D1.** R20-1 is CONFIRMED on arithmetic from client values and the bands' own gear. The margin from 40 to 58 (break-even armor 42 to 44% against a real 33 to 38%) does not depend on the gear choice within any plausible leveling set. The fix drops the comparison with Serpent Sting rather than adding an Aimed Shot exception, so it reads right at every band.
- **D2.** R20-2 is CONFIRMED as a self-contradiction created by round 20's gate: same card, same page.
- **D3.** R20-3 is LIKELY: on Marksmanship it turns on talent multipliers and crit, so it waits for sim 1. The fix is true either way, so the integrator can apply it now.
- **D4.** The Survival Hunter's Mark MUST is routed to round 21 and not re-filed; the Expose Prey wording is a Should.
- **D5.** Earlier decisions stand (rows that render before their talent are the component lane's; Heart of the Lion 280; the pet rule).

## Sims requested

1. Damage per mana by spell (spec-breakdown damage divided by casts times mana) for rank 1 Aimed Shot, Serpent Sting and Arcane Shot: MM leveling build at bands 40, 50, 52, 58 and bare 60; BM leveling build at 40, 50, 58. Decides R20-3 and confirms R20-1's margins.
2. BM and MM at bands 18 and 19 (and 20) on a long fight that runs below a quarter of mana: the shipped list against Serpent Sting and Arcane Shot without the Multi-Shot readiness check while Multi-Shot is held (see Sim findings).
3. Carried: MM race order on the raid build; Survival Crusader against Agility; BM and SV race re-runs; the best non-BoP two-hander; BM picks with the pet on Heart of the Lion; the 46-47 Raptor Strike row on re-ranked gear; the low-rank Aimed Shot below 20% mana at raid 60 (round 16 D2).

## Sim findings

- **New (APL):** below a quarter of your mana at 18 and 19 the BM and MM cards cast nothing but Hunter's Mark. Multi-Shot is held by its mana gate, so its cooldown never starts; Serpent Sting needs Multi-Shot 0.5 s or more from ready and Arcane Shot needs it more than 1.5 s away, and with no Aimed Shot to spend the cooldown neither condition is ever met. From 20 Aimed Shot spends it, so the stall is short. The notes ("keep it rolling for the whole fight", "keep it going even when the mana runs low") describe the right play; the APL conditions should skip the Multi-Shot check while Multi-Shot is held.
- **Carried, open with the sim lane:** the hawk's level 60 swing; Heart of the Lion on the pet; the Focused Fire snapshot report; Ranged Weapon Specialization on the sting; Lacerating Strikes and the queued Raptor Strike; the single-source Immolation Trap 2%; stale top-level APL notes (not rendered).

## Engine questions

- Immolation Trap's share is read from ranged attack power (`immolationTrapTickAPCoeff` on RAP). Blizzard says "Attack Power"; for a melee Survival hunter, is it melee or ranged? A level 16+ log with Strength gear swapped would settle it.
- Serpent Sting and Arcane Shot read the share when cast or applied; Blizzard and the wiki say ticks recalculate. The difference shows only during trinket and racial windows; noted for completeness.

## Sources

**Client data, 1.60.1.70790:** `spellconst/hunter.json` (19434-20904, 14281-14287, 13549-13555, 2643, 1310687-1310786, 3045, 19574, 13795-14305, 1130-14325, Aspect of the Hawk 13165-25296, Heart of the Lion 409583), `spelltips/hunter.json` (same ids, 1495-14271, 24118-1299332, 1317257, 2974, 20910, 1242634), `talents/hunter.json`, `addon-builds.json`, `addon-data.json`, `gametables/basemp.txt`, `levels.json`, `items/hunter.json`, `bis/hunter-beast-mastery.json` and `hunter-marksmanship.json` (bare band slots), `data/curated/rank-choices.json`.

**Engine d5c708834 (gr20/engine):** `sim/hunter/attack_power_scaling.go`, `aimed_shot.go`, `arcane_shot.go`, `serpent_sting.go`, `multi_shot.go`, `sniper_shot.go`, `hunter.go` (ammo), `talents.go`.

**Repository:** `design/reviews/guide-accuracy/round19/hunter.md`, `ledger.md` (Round 19), `design/reviews/2026-11-04-sim-audit-round20/hunter.md` and `integration.md`.

**Web (fetched 2026-11-04):** forever-xp.com, "Hunter changes", https://forever-xp.com/changes/hunter/ (quotes Blizzard's 3 November lines on Arcane Shot 0.15 and Serpent Sting 5% a tick, and the 16 October fire-trap line). Blizzard's WoW Forever Beta Development Notes as cited in the sim report: https://us.forums.blizzard.com/en/wow/t/major-class-final-tuning/2399448.

## Note sweep

One line per row note: spec, row, bands it renders, each named talent, spell, rank or value with the level the player has it, then the verdict.

- BM Hunter's Mark (10-60): spell 6; "one mark per target" (no stacking) - OK.
- BM Bestial Wrath (40-60): talent 40; Rapid Fire 26; 50%, 18 s, fear immunity, 120 s (19574) - OK.
- BM Rapid Fire (26-60): Multi-Shot 18; Bestial Wrath "from 40"; five minutes (3045), four from Rapid Killing at 57 (band 58), five on the raid build at 60; trinkets and racial - OK.
- BM Aimed Shot (20-60): spell 20; shared cooldown (tooltip); rank 1 to 58, rank 6 at 60; "small extra over Multi-Shot" (+40 to +350) OK; "Multi-Shot on every cooldown is fine" - **HUNTER-R20-2**.
- BM Multi-Shot (18-60): spell 18; Aimed Shot 20 (parenthetical); quarter-mana gate matches the APL; instant; worst damage per mana at every band - OK.
- BM Serpent Sting (10-60): ranks 2 to 8 by trainer; "from 18" Multi-Shot; 5% RAP a tick (Blizzard); "best damage for its mana you fire": true on BM at every band on the arithmetic (Improved Stings 46-47 keeps it above rank 1 Aimed Shot at 50 to 58), but see R20-3 for one wording; "pushes back the shared cooldown" at 18 - carried Should; "Multi-Shot goes first" under a quarter mana - Should.
- BM Arcane Shot (10-60): ranks 1 to 8 by trainer; "from 18"; 15% RAP (Blizzard); "second only to Serpent Sting for damage per mana" - **HUNTER-R20-1** (bands from the mid 30s to 58 and bare 60).
- MM Hunter's Mark (10-60): as BM - OK.
- MM Rapid Fire (26-60): Multi-Shot 18; Rapid Killing 25-26 so three minutes on leveling bands 26-58; raid build five minutes at 60 - OK.
- MM Aimed Shot (20-60): spell 20; rank 1, rank 4 at 50, rank 6 at 60, each above Multi-Shot per cast (flat 40/185/350, Barrage on both) - OK.
- MM Multi-Shot (18-60): as BM - OK.
- MM Sniper Shot (40-60): raid talent at 40 (leveling bands render the row: component lane, D5); 4000 ms cast, 15 s cooldown, hardest hit (weapon + 320/450/590) - OK.
- MM Serpent Sting (10-60): "the shared cooldown" at 18-19 - Should; 5% RAP OK; "best damage for its mana you fire" at 58 and bare 60 - **HUNTER-R20-3** (LIKELY).
- MM Arcane Shot (10-60): "the shared cooldown" at 18-19 - Should; 15% RAP OK; "second only to Serpent Sting" - **HUNTER-R20-1** (MM 34-48, 52, 58, bare 60).
- SV Raptor Strike (10-60): ranks 2 to 8 (46-47 rank 5 row); no global, replaces the next swing; Resourcefulness 33-34, 60% at rank 2 - OK (46-47 wording carried Should).
- SV Counterattack (30-60): talent 32 (row at 30: component lane); cannot be dodged, blocked or parried (client text) - OK.
- SV Mongoose Bite (16-60): spell 16; after you dodge, 5 s window; Expose Prey 35-36, 10 s, any hunter's mark; one press per window - OK (chance wording Should).
- SV Lacerate (30-60): spell 30 (ranks 30/40/50/60); Mongoose Bite 16 - OK.
- SV Strider Kick (30-60): talent 30; 8 s cooldown, 100% weapon damage, 5.81% of base mana; Lacerate 30 - OK.
- SV Immolation Trap (16-60): ranks 16/26/36/46/56; Clever Traps 20-21 and Survivalist's Discipline 37-38 as a "once in" clause; attack power share (Blizzard 16 October) - OK (ranged wording Should).
- SV Wing Clip (12-60): ranks 12/38/60; Windfury Totem conditional; 5/25/50 damage; main hand hit with weapon procs - OK.
