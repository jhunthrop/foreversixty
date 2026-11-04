# Guide accuracy round 2: Hunter

Auditor: hunter (class index, Beast Mastery, Marksmanship, Survival). Tree: land/guides-r1 609eb471c, client 1.60.1.70755, engine pin d1d9f2775.
Files: `web/src/content/guides/hunter/{index,beast-mastery,marksmanship,survival}.md`.

## Builds and leveling order (re-checked)

Build strings are unchanged since round 1 and still decode as listed in `round1/hunter.md` (all legal on 70755, 51 points, gates and prerequisites met). Leveling order from `data/builds/1.60.1.70755/addon-builds.json` (index 0 = level 10):

- **BM**: Deadly Aspects 10-14, Focused Fire 15-16, Endurance Training 17-20, Bestial Swiftness 21, Unleashed Fury 22-26, Ferocity 27-31, Intimidation 32, Bestial Discipline 33-34, Frenzy 35-39, Bestial Wrath 40, Lethal Attacks 41-45, Improved Stings 46-47, Efficiency 48-51, Careful Aim 52-56, Rapid Killing 57, Trueshot Aura 58, Mortal Shots 59-60. Prose ("first five points into Deadly Aspects", "Focused Fire at 15 and 16", "Bestial Wrath comes the same level" as mail at 40) matches.
- **MM**: Lethal Attacks 10-14, Improved Stings 15, Efficiency 16-20, Careful Aim 21-25, Trueshot Aura 26, Mortal Shots 27-31, Barrage 32-34, Ranged Weapon Specialization 35-39, Deadly Aspects 40-44, Endurance Training 45-47, Focused Fire 48-49, Unleashed Fury 50-54, Ferocity 55, Improved Tracking 56-60. Prose names no order; nothing contradicts it.
- **Survival**: Improved Tracking 10-14, Savage Strikes 15-16, Survivalist 17-19, Clever Traps 20-21, Surefooted 22-24, Deterrence 25, Predator's Edge 26-29 and 31, Strider Kick 30, Counterattack 32, Resourcefulness 33-34, Expose Prey 35-36, Survivalist's Discipline 37-38, Lightning Reflexes 39 and 41-44, Lacerating Strikes 40, Deadly Aspects 45-49, Lethal Attacks 50-54, Careful Aim 55-59, Lone Wolf 60. Prose ("Lone Wolf is one of the last points ... dismiss it the moment you take Lone Wolf") matches.

## What was checked and holds

- Every talent value in the prose against `talents/hunter.json` (70755 rank text): Deadly Aspects (Hawk: Auto Shot; Beast: melee auto attacks; 10%/30%/12 s), Focused Fire 2%, Unleashed Fury 15%, Ferocity 10%, Frenzy, Bestial Wrath 50%/18 s, Bestial Discipline (50% regen while casting), Intimidation, Improved Stings 13%/20%, Efficiency 15% (shots, stings and melee), Careful Aim 100% Int, Rapid Killing rank 1 (1 min, 10%), Lone Wolf 20%, Mortal Shots 6%/rank, Barrage 15%, RWS 5%, Sniper Shot (tier 6, one point past 30), Improved Tracking (the seven types), Savage Strikes 4%, Surefooted 3%, Predator's Edge 30%/50%, Resourcefulness 60%, Expose Prey, Survivalist's Discipline 40%, Lightning Reflexes 15%, Lacerating Strikes 40%/21 s.
- Spells and levels against `trainables/hunter.json` and the client descriptions: Hunter's Mark 6, Tame Beast 10, Wing Clip 12, Mongoose Bite 16 ("Can only be performed after you dodge"), Aspect of the Cheetah 20, Aimed Shot 20 (2 s cast, shares Multi-Shot's cooldown), Rapid Fire 26 ("Increases ranged and melee attack speed by 40% for 15 sec", 5 min), Feign Death 30, Lacerate 30/40/50/60, Aspect of the Beast 30/40/50/60 (110 melee AP at 60), Volley 40 (channeled), Falcon 60, Tranquilizing Shot 60, Disengage 5 s, Multi-Shot "hitting 3 targets".
- Values the 70755 push moved, checked against Wowhead Forever's live tooltips (the lesson from round 1): Trueshot Aura rank 5 "by 100" (spell 20906), Aimed Shot rank 6 "increases ranged damage by 350" (20904), Heart of the Lion buff "Attack Power by 280" at 60 (409583). The prose is right on all three.
- Rotation prose matches the APL cards line for line (BM/MM: Hunter's Mark, Bestial Wrath, Rapid Fire, Aimed Shot when the next auto is more than 1.2 s away, Multi-Shot, Serpent Sting under 1.5 s, Arcane Shot above 25% mana; Survival: Raptor Strike, Counterattack, Mongoose Bite, Lacerate, Strider Kick, Immolation Trap, Wing Clip only under Windfury Totem).
- Raid buffs claim (survival.md:59): Forever Strength of Earth Totem rank 5 is 53 Strength (Era 77) and Curse of Recklessness rank 4 is 505 armor (Era 640) in `spelltips/shaman.json`, `spelltips/warlock.json`. True.
- Ammo and quivers in `items.json`: Jagged Arrow 40, Ice Threaded Arrow 51, Thorium Headed Arrow 52, Thorium Shells 52, Quickdraw Quiver 40, Harpy Hide Quiver 55, Ancient Sinew Wrapped Lamina 60, Sniper Scope 40, Biznicks 50, Major Mana Potion 49, Major Frenzy Potion 55, Elixir of the Mongoose 46, Poached Sunscale Salmon 35 (Forever: 40 melee and ranged attack power, engine consumes.go:384).
- statPriority frontmatter still agrees with the band 60 raid weights (BM Agi 1.00 ~ Int 1.01; MM Agi > Int > Crit; SV Crit > Agi > Int > Str > AP), hit first in each.
- No em or en dashes, no site/sim words, no quoted nightly numbers, no removed spells or SoD runes.

### Round 1 research answers re-checked against the description text

- **Falcon (Q1).** Deadly Aspects' description (talents/hunter.json 19552, and Wowhead Forever spell 19552) names only two cases: Auto Shot under Hawk, melee auto attacks under Beast. Falcon's description (469145) adds only that the talent "can now proc from and benefit melee attacks" under Falcon. Neither text gives Falcon a ranged proc, and the cited Blizzard forum thread is titled exactly that Falcon does not trigger it on ranged attacks. Holds.
- **Tranquilizing Shot (Q2).** Trainer row 19801 at 60, description "Attempts to remove 1 Frenzy effect from an enemy creature." Holds.
- **Cat (Q4), Aimed Shot timing.** No description claim involved. Holds.
- **Races.** Warcraft Tavern's Forever hunter guide does say "the best options for PvE are likely either Night Elf or Human ... due to their Crit racials" and "Orcs and Trolls remain strong options for Horde Hunters"; its Dwarf and Orc paragraphs still describe Earth Fury and Savage, which the 70755 client replaced (races.json: Artisan's Precision; Savage removed). Quoted correctly in round 1; now decided below on the measured race sims.

## Findings

| id | status | kind | file:line | wrong claim | evidence | fix |
|---|---|---|---|---|---|---|
| HUNTER-R2-1 | CONFIRMED | spell | beast-mastery.md:103, marksmanship.md:106 | "Distracting Shot pulls a loose add onto you for a few seconds so your tank can grab it" | Distracting Shot rank 6 (15632) is a single threat effect: spellconst effect 63, amount 600, duration 0; description "Distract the target, causing threat." (spelltips and Wowhead Forever 15632, no duration). It is not a timed taunt: the add stays on you until someone passes your threat. | "Distracting Shot puts a big lump of threat on a loose add, enough to pull it off a healer and onto you. Bring it to your tank and let him take it from you." |
| HUNTER-R2-2 | CONFIRMED | stats | marksmanship.md:110, survival.md:96 | "Surefooted gives you 3% of that cap and 1% is built in, so you need 5% from gear" | Hunters have no built-in hit. The only 1% hit source is the Dwarf racial Artisan's Precision (races.json; engine racials.go:299). The line is wrong for every other race, and with the race decision below the guides no longer lead with Dwarf. | MM: "In the raid build Surefooted gives you 3% of that cap, so you need 6% from gear, or 5% as a Dwarf." SV: "Surefooted gives you 3% of it, so you need 6% from gear, or 5% as a Dwarf." |
| HUNTER-R2-3 | CONFIRMED | race | beast-mastery.md:8,123,125 | recommendedRaces [dwarf, troll]; "Dwarf for the Alliance"; "Troll or Orc for the Horde, and the two are close" | Race sims (design/reviews/guide-accuracy/round2/race-sims.md, raid preset band 60, error about ±0.23): Night Elf 664.1 > Skyborne 661.0 ~ Human 660.8 > Dwarf 659.0; Orc 668.5 > Windshaper 663.7 ~ Troll 663.4. Night Elf beats Dwarf by 0.8% and Orc beats Troll by 0.8%, both far outside error. | See Decisions D1, D2. |
| HUNTER-R2-4 | CONFIRMED | race | marksmanship.md:8,126,128 | recommendedRaces [dwarf, troll]; "Dwarf for the Alliance"; "Troll for the Horde ... Orc's Blood Fury is the other good pull cooldown" | Race sims: Night Elf 709.3 > Skyborne 706.6 ~ Human 706.0 > Dwarf 703.7; Orc 711.0 > Skyborne 709.0 > Troll 707.9 (error about ±0.28). | See Decisions D3. |
| HUNTER-R2-5 | CONFIRMED | race | survival.md:7,110,112 | recommendedRaces [dwarf, troll]; "Artisan's Precision is 1% hit, which is worth more to a dual-wielder than to anyone"; "Troll for the Horde" | Race sims: Night Elf 755.1 > Human 751.8 > Skyborne 749.9 > Dwarf 747.7; Orc 753.5 > Troll 751.4 > Skyborne 750.6 (error about ±0.3-0.4). Dwarf is last of the four Alliance races for the dual-wielder, so the 1% hit claim is false as a ranking reason. | See Decisions D4. |
| HUNTER-R2-6 | LIKELY | gear | survival.md:100 | "Predator's Edge's 50% off-hand bonus makes the off hand worth almost as much as the main hand" | Off-hand swings do half damage; Predator's Edge raises that by 50% (talents/hunter.json: "Off Hand weapon damage by 50%"; engine talents.go:346 multiplies off-hand damage by 1.5), so the off hand lands at three quarters of the main hand, and Raptor Strike, Mongoose Bite and Strider Kick are main hand only. "Almost as much" oversells it. | "Survival dual-wields, and Predator's Edge's 50% off-hand bonus lifts the off hand to three quarters of a main-hand swing, so it is a real weapon: pick it for damage and speed, not just stats." |
| HUNTER-R2-7 | NEEDS-SIM | talent | survival.md:76 | "Leave Rapid Killing out there too: a shorter Rapid Fire does not pay for the points it costs." | Unmeasured. Rapid Killing 2 takes 2 minutes off Rapid Fire (talents/hunter.json 415405), and in Forever Rapid Fire hastes melee swings by 40% (spell 3045), so a 3 minute Rapid Fire is a real melee cooldown for Survival. The points would come from Lethal Attacks (5 to 3, the only movable Marksmanship points; Careful Aim and Lone Wolf stay). | Keep the line until S-R2-1 runs. If Rapid Killing loses: "Leave Rapid Killing out: two more points of Lethal Attacks are worth more than a second Rapid Fire in a long fight." If it wins, move the build and write: "Two points in Rapid Killing bring Rapid Fire back every three minutes, and since it speeds your swings, that is a second melee cooldown in most boss fights." |

## Decisions (owner rule: decide, Forever sources first, then evidence)

No weapon-type racial touches a hunter on 70755 (races.json: Dwarf Artisan's Precision and Big Game Hunter, Night Elf Elune's Light and Quickness, Human Ruthless, Orc Blood Fury and Command, Troll Berserking and Beast Slaying, Skyborne Wind Blessed and Elemental Insight), so the sims' weapon caveat does not apply.

- **D1. Beast Mastery, Horde: Orc.** Orc 668.5 vs Troll 663.4 (+0.8%, about twenty times the error); Command is 5% on the pet all fight and Blood Fury's 2 minute cooldown comes up with every Bestial Wrath. Warcraft Tavern's Forever BM guide: Orc "benefits from a powerful Attack Power cooldown alongside a passive 5% pet damage bonus" and ranks among its strongest Horde picks, while Troll's speed cooldown is a burst that works best "when the cooldown can be used fully". Change beast-mastery.md:125 to: "Orc for the Horde. Command makes your pet hit 5% harder all fight, and Blood Fury's two minute cooldown comes up with every Bestial Wrath, so you press them together twice in most boss fights. Troll is a fine second: Berserking is 10% attack speed for 15 seconds and stacks with Rapid Fire on the pull."
- **D2. Beast Mastery, Alliance: Night Elf.** Night Elf 664.1 > Skyborne 661.0 ~ Human 660.8 > Dwarf 659.0. Warcraft Tavern: Night Elf or Human "due to their Crit racials". Change beast-mastery.md:123 to: "Night Elf for the Alliance: Elune's Light is 10% crit for 15 seconds, it lines up with Bestial Wrath and Rapid Fire on the pull, and your pet copies that crit. Human is next with Ruthless's 1% crit all fight, and Dwarf's 1% hit and Big Game Hunter come last. Humans can be Hunters for the first time in Forever." Frontmatter line 8: `recommendedRaces: [night-elf, orc]`. Build strings lines 6-7: race field `dwarf` to `night-elf` (`FS1:1.60.1.70755:hunter:night-elf:...`), so the planner opens on the recommended race.
- **D3. Marksmanship: Night Elf and Orc.** Night Elf 709.3 vs Dwarf 703.7; Orc 711.0 vs Troll 707.9 (both far outside ±0.28). With no pet, Orc wins on Blood Fury alone. Change marksmanship.md:126 to the BM Alliance line without the pet clause: "Night Elf for the Alliance: Elune's Light is 10% crit for 15 seconds and lines up with Rapid Fire on the pull. Human is next with Ruthless's 1% crit all fight, and Dwarf's 1% hit and Big Game Hunter come last. Humans can be Hunters for the first time in Forever." Change marksmanship.md:128 to: "Orc for the Horde: Blood Fury is 10% attack power for 15 seconds every two minutes, so it goes with Rapid Fire on the pull and again at two and four minutes. Command does nothing for you with the pet dismissed. Troll's Berserking is the close second." Frontmatter line 8: `recommendedRaces: [night-elf, orc]`; build strings lines 6-7 race `night-elf`.
- **D4. Survival: Night Elf and Orc.** Night Elf 755.1 vs Dwarf 747.7 (+1.0%); Orc 753.5 vs Troll 751.4 (+2.1 DPS, about four times the combined error in the raid preset; level with Troll in the bare 60 run, so the raid result decides). Change survival.md:110 to: "Night Elf for the Alliance: Elune's Light is 10% crit for 15 seconds, and Predator's Edge makes every crit count, so press it with Rapid Fire on the pull. Human's Ruthless is next, and Dwarf's 1% hit comes last even for a dual-wielder. Humans can be Hunters for the first time in Forever." Change survival.md:112 to: "Orc for the Horde: Blood Fury is 10% attack power for 15 seconds every two minutes, on the pull with Rapid Fire and again at two and four minutes. Troll is right behind: Berserking is 10% attack speed for 15 seconds and stacks with Rapid Fire on your swings." Frontmatter line 7: `recommendedRaces: [night-elf, orc]`; build line 6 race `night-elf`.
- **D5. Skyborne lines** (beast-mastery.md:127, marksmanship.md:130, survival.md:114) stay as written: true, and Skyborne lands mid-pack on both factions in every spec.

## Sim findings (for the next sim round; not prose)

- **S-R2-1 (Survival Rapid Killing).** Survival raid build with Rapid Killing 2 in place of Lethal Attacks 2 (MM 3/0/0/0/0/5/2/0/1 shape: Lethal Attacks 3, Careful Aim 5, Rapid Killing 2, Lone Wolf 1), band 60 raid preset, rotation unchanged (Rapid Fire on cooldown). Adopt under the talent rule only beyond error. Settles HUNTER-R2-7.
- **S-R2-2 (stale guide tooltips, component).** `data/builds/1.60.1.70755/spelltips/hunter.json`, which web/src/lib/guides/guide-tips.ts serves as the guide hover tooltips, still carries pre-push description text: Trueshot Aura rank 5 "by 50" (20906), Trueshot Aura talent rank 1 "by 15" (1299346), Aimed Shot rank 6 "by 166" (20904), Arcane Shot rank 8 "217" (14287). spellconst, talents/hunter.json and Wowhead Forever's live tooltips all have the pushed values (100, 20, 350, 280). A reader hovering Trueshot Aura in the In a raid bullet sees 50 next to prose that says 100. Regenerate spelltips from the hotfixed SpellEffect values.
- **S-R2-3 (bis files stale, carried from S9).** All three `bis/hunter-*.json` band 60 weights are engine 8baffe5fc, 2026-11-01, Troll, and the Survival weights character runs Marksmanship 3/2/5/0/0/5/0/0/1 with no Deadly Aspects, not the guide build. The prose quotes none of it; refresh measures it.

## Sims requested

None beyond S-R2-1. The race sims sent by the integrator settle D1 to D4.

## Owner questions

None. Every open item is decided above or waits on S-R2-1 with the current line standing as the sensible default.

## Sources

- Wowhead Forever, Trueshot Aura (rank 5), https://www.wowhead.com/forever/spell=20906 (tooltip: "by 100")
- Wowhead Forever, Aimed Shot (rank 6), https://www.wowhead.com/forever/spell=20904 (tooltip: "by 350")
- Wowhead Forever, Heart of the Lion buff, https://www.wowhead.com/forever/spell=409583 (buff: Attack Power by 280)
- Wowhead Forever, Deadly Aspects, https://www.wowhead.com/forever/spell=19552 (names Hawk and Beast only)
- Wowhead Forever, Distracting Shot (rank 6), https://www.wowhead.com/forever/spell=15632 ("Distract the target, causing threat", no duration)
- Warcraft Tavern, Forever Hunter Guide, https://warcrafttavern.com/forever/guides/hunter/ (Night Elf or Human for Alliance PvE; Orc and Troll strong for Horde)
- Warcraft Tavern, Forever Beast Mastery Hunter Guide, https://www.warcrafttavern.com/forever/guides/beast-mastery-hunter/ (Orc for BM)
- Race sims: design/reviews/guide-accuracy/round2/race-sims.md (this tree)

## Fix log (round 2)

Integrator race decision (guide builds, raid band 60, seeds 1|50001): BM [night-elf, orc], MM [night-elf, troll] (Troll beyond error over Orc), SV [night-elf, orc] (Orc/Troll tie, Warcraft Tavern consensus).

- HUNTER-R2-1 fixed: beast-mastery.md:103, marksmanship.md:106 (Distracting Shot is a threat lump, no timed taunt).
- HUNTER-R2-2 fixed: marksmanship.md:110, survival.md:96 (no built-in hit; 6% from gear, 5% as a Dwarf).
- HUNTER-R2-3 / D1 / D2 fixed: beast-mastery.md:6-8 (race field night-elf on both FS1 strings, recommendedRaces [night-elf, orc]), :123 Night Elf line, :125 Orc line with Troll second.
- HUNTER-R2-4 / D3 fixed with the integrator's Horde override: marksmanship.md:6-8 (night-elf strings, recommendedRaces [night-elf, troll]), :126 Night Elf line, :128 Troll first (Berserking stacks with Rapid Fire), Orc close behind with Command noted as dead with the pet dismissed.
- HUNTER-R2-5 / D4 fixed: survival.md:6-7 (night-elf string, recommendedRaces [night-elf, orc]), :110 Night Elf line, :112 Orc line with Troll right behind.
- HUNTER-R2-6 fixed: survival.md:100 (off hand at three quarters of a main-hand swing with Predator's Edge).
- HUNTER-R2-7 fixed without a verdict: survival.md:76 now states the structure only. Rapid Killing is tier 2 beside Lone Wolf (talents/hunter.json), Lethal Attacks 5 + Careful Aim 5 are the ten points that open the row, so its two points must come out of Survival or Deadly Aspects. S-R2-1 as written was illegal (Lethal Attacks 3 leaves Careful Aim's row short), per the integrator.
- D5 left as written (Skyborne lines make no ranking claim).
- Source labels: talents/spellconst/races/combos .json labels renamed to "Forever Sixty, talent trees / spell values / racial abilities / race and class pairs" in all four files. No planner links appear in the prose, so the planner-link form had nothing to change.
- Card notes: re-read all hunter APL card notes; none states a race, Distracting Shot or hit claim, nothing false found; unchanged. Actions untouched.
- S-R2-2 (stale spelltips hover text) handed to the data lane per the integrator; not chased.

New SIM findings: none. If wanted, a legal Rapid Killing test for Survival is Lethal Attacks 5, Careful Aim 5, Lone Wolf 1, Rapid Killing 2 with two points taken from Deadly Aspects (5 to 3) or from the Survival tail; adopt only beyond error under the talent rule.

### Integrator: S-R2-1 measured
Survival on the guide build (Night Elf, raid band 60, 10,000 iterations, seeds 1|50001): published 5/005005001/500230231051222151 820.44|819.41; Rapid Killing 2 from Deadly Aspects 2 (3/005005201/500230231051222151, the legal swap) 804.56|804.56, about 1.9% lower, beyond error. survival.md:76 now says the swap loses.
