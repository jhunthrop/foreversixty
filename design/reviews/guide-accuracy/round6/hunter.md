# Guide accuracy round 6: Hunter

Auditor: hunter (class index, Beast Mastery, Marksmanship, Survival). Tree: gr6/audit at 367cbadb7 (land/guides-r5, all five rounds applied), client 1.60.1.70755, engine pin d1d9f2775 (sim/enginever/version.go).
Files: `web/src/content/guides/hunter/{index,beast-mastery,marksmanship,survival}.md`, card notes in `data/curated/apl/hunter-*.json`. Read first: round1 to round5 hunter reports with every fix log, round4/hunter-levels.md, round2 to round5 race-sims.md, ledger.md.

## Builds and leveling order

Build strings, recommendedRaces and point totals are unchanged since round 3. The decodes in round3/hunter.md hold: all five are legal, 51 points, tier gates and prerequisites met.

Shipped order (`addon-builds.json`, index 0 = level 10), re-decoded this round against `talents/hunter.json`:
- BM leveling: Deadly Aspects 10-14, Focused Fire 15-16, Endurance Training 17-20, Bestial Swiftness 21, Unleashed Fury 22-26, Ferocity 27-31, Intimidation 32, Bestial Discipline 33-34, Frenzy 35-39, Bestial Wrath 40, Lethal Attacks 41-45, Improved Stings 46-47, Efficiency 48-51, Careful Aim 52-56, Rapid Killing 57, Trueshot Aura 58, Mortal Shots 59-60.
- MM leveling: Lethal Attacks 10-14, Improved Stings 15, Efficiency 16-20, Careful Aim 21-25, Trueshot Aura 26, Mortal Shots 27-31, Barrage 32-34, Ranged Weapon Specialization 35-39, Deadly Aspects 40-44, Endurance Training 45-47, Focused Fire 48-49, Unleashed Fury 50-54, Ferocity 55, Improved Tracking 56-60.
- SV (leveling = raid): Improved Tracking 10-14, Savage Strikes 15-16, Survivalist 17-19, Clever Traps 20-21, Surefooted 22-24, Deterrence 25, Predator's Edge 26-29 and 31, Strider Kick 30, Counterattack 32, Resourcefulness 33-34, Expose Prey 35-36, Survivalist's Discipline 37-38, Lightning Reflexes 39 and 41-44, Lacerating Strikes 40, Deadly Aspects 45-49, Lethal Attacks 50-54, Careful Aim 55-59, Lone Wolf 60.

Every leveling sentence matches it (BM Deadly Aspects first five, Focused Fire 15 and 16, Endurance Training from 17, Bestial Wrath at 40 with mail; MM Lethal Attacks, Improved Stings at 15, Efficiency from 16; SV Improved Tracking, Savage Strikes 15 and 16, Survivalist from 17, Lone Wolf the last point).

## What was checked and holds

**Levels 1 to 19, then Levels 20 to 60 (reading order).** On all three pages the two blocks now read in order without a repeat or a contradiction: trainer levels and Heart of the Lion, the taming quests at 10 and Mend Pet at 12, the first talent points, quest gear and leather; then the pet-first pull, Wing Clip and the dead zone, Feign Death at 30, Cheetah at 20, ammo and mail at 40. Survival's "Wear Aspect of the Monkey until 30 ... At 30 switch to Aspect of the Beast" follows "Train Aspect of the Monkey ... at 4" cleanly; BM and MM's "Aspect of the Hawk at 10, which replaces Monkey for good" is never contradicted below. Every level re-checked against `trainables/hunter.json` (Heart of the Lion 1, Raptor Strike 1/8/16, Serpent Sting and Monkey 4, Arcane Shot and Hunter's Mark 6, Concussive Shot 8, Hawk 10, Tame Beast 10, Wing Clip and Mend Pet 12, Mongoose Bite 16, Multi-Shot 18, Track Beasts 1, Humanoids 10, Undead 18, Cheetah 20, Feign Death 30, Lacerate 30, Beast 30).

**Stat paragraphs against the page tables.** The page table reads `bis/hunter-<spec>.json`, band 60, raid preset, at the faction of recommendedRaces[0] (night-elf, so Alliance; `rail-stats.ts`, `class-dps.ts`, `presets.ts selectBand`). After the nightly re-rank (de044125c, engine d1d9f2775) the scale factors are:
- BM: Hit 1.29 > Agility 1.00 > Crit 0.85 > Intellect 0.46 > RAP 0.30 (before the re-rank, 8baffe5fc: Hit 1.16 > Int 1.01 ~ Agi 1.00 > Crit 0.77). The prose still gives the old order: HUNTER-R6-1.
- MM: Agility 1.00 > Crit 0.82 > Intellect 0.56 > RAP 0.31, hit not significant (Alliance). The round 5 review rewrite matches.
- SV: Hit 1.74 > Crit 1.17 > Agility 1.00 > Intellect 0.52 > Strength 0.51 > AP 0.42. The prose order matches.
The frontmatter statPriority arrays only choose which rows show; the rows sort by scale factor (`railStatRows`), so their order is not a page claim.

**Enchant prose against the enchant table** (`gear-extras.ts pickEnchants`, raid band 60 weights, replicated for each page): BM and MM tables show Greater Agility gloves and boots, Lesser Agility cloak, Greater Stats chest, Greater Intellect bracers and 2H Agility on the two-hander. Their prose matches. Survival's table shows Superior Strength on bracers and Agility on both one-handers; its prose does not: HUNTER-R6-2.

**Gear prose.** No hunter page names a gear piece. The reasons given are stable: leather to 40 and mail from 40 (owner rule), the ranged weapon is the real weapon for BM and MM, a two-hander carries more stats than a one-hander (the BM and MM BiS main hand is the two-handed Legionite Glaive), Survival dual-wields with Predator's Edge's 50% off-hand bonus, and the ranged slot is a stat stick for Survival. Quivers checked in raw item data: Quickdraw Quiver 40 is 13%, Harpy Hide Quiver 55 and Ancient Sinew Wrapped Lamina 60 are 15%, so "10 to 15%" and "Quickdraw Quiver from 40 and a 15% bag near 60" hold.

**Consumables.** The table comes from sim/request/raid-consumes.json (unchanged since 2026-11-03 16:33): BM and MM at 49+ take Major Mana Potion, Elixir of the Mongoose, Ground Scorpok Assay, Poached Sunscale Salmon and Juju Might; Survival at 55+ takes Major Frenzy Potion with stones, Juju Power and Juju Might and no mana potion. The prose names a subset and states nothing the table contradicts.

**Talents and spells.** Every value re-read against the 70755 rank text and spelltips: Deadly Aspects (Hawk: Auto Shot; Beast: all melee autos; 2% a rank, 30%, 12 sec), Focused Fire 2%, Unleashed Fury 15%, Ferocity 10%, Frenzy 100%/30%/8 sec, Bestial Wrath 50%/18 sec/no fear, Bestial Discipline 50% regen while casting, Bestial Swiftness 30%, Intimidation (3 sec stun on the pet's next hit), Rapid Killing rank 1 (1 min, 10%), Lethal Attacks 5%, Improved Stings 6%/13%/20%, Efficiency 15% (shots, stings, melee), Careful Aim 100% of Intellect, Mortal Shots 6% a rank on ranged abilities (engine class mask includes Auto Shot), Barrage 15%, RWS 5%, Lone Wolf 20%, Sniper Shot tier 6 (4 sec cast, Volley's channel is 6 sec, Aimed Shot 2 sec), Improved Tracking (the seven types), Savage Strikes 4% melee ability crit, Survivalist 2% a rank, Surefooted 3%, Predator's Edge 30%/50%, Resourcefulness 60% plus the crit regen, Expose Prey on any Hunter's Mark, Survivalist's Discipline 40%, Lightning Reflexes 15%, Lacerating Strikes 40%/21 sec. Spelltips: Aimed Shot rank 6 350 and shared 6 sec cooldown, Multi-Shot "hitting 3 targets", Rapid Fire 40% ranged and melee for 15 sec on 300 s, Trueshot Aura rank 5 100, Aspect of the Hawk rank 6 110 and rank 7 120 (book), Aspect of the Beast rank 4 110, Falcon = highest Hawk, Mongoose Bite after a dodge, Counterattack after a parry, Disengage "Disables auto-attack" on 5 sec, Feign Death, Freezing Trap "Any damage caused will break the ice", Tranquilizing Shot trainer row at 60. Alliance has Dwarf shamans (combos.json), so "in a group without a shaman" in the Wing Clip line is right on both factions.

**Card notes.** All BM, MM and Survival notes re-read; none states anything false.

**Races.** Prose matches the guide-build reruns in round2/race-sims.md and the racial texts in races.json. The raid preset has not changed since those runs; the gear has (see Sims requested 2).

**Round 5 edits re-checked.** The racial in each burst line (holds, see HUNTER-R6-3 for the wording), the pet-first leveling pull (holds), "a pure damage class" (holds), the MM stat rewrite (holds against its table), the Improved Stings row (6% and 20%, holds), Wing Clip 12 and Multi-Shot 18 in Levels 1 to 19 (holds).

**Voice.** No em or en dashes; no sim, site, preset, band or pipeline words; no quoted nightly numbers; no removed spells or SoD runes.

## Source links

All resolve and match their labels: Blizzard Hunter and Druid deep dive, Deep Dive panel recap, What's Next recap (200); Warcraft Tavern Hunter and Druid Deep Dive, Forever Hunter Guide, Forever Hunter Leveling Guide (200); Wowhead Skyborne First Look (200); foreversixty.gg planner, planner?class=hunter, spellconst/hunter.json, races.json, combos.json, /sim/weights (200, damage specs only); /bis/hunter/<spec> body links (200). Wowhead spell=1495 and spell=19306 answer 403 to fetchers (unknown, not dead; round 5 matched them through the tooltip endpoint). Against the round 4 rule: no /sources links, no wago or Talents Forever entries, talent trees open the planner.

## Findings

| id | status | kind | file:line | wrong claim | evidence | fix |
|---|---|---|---|---|---|---|
| HUNTER-R6-1 | CONFIRMED | stats | beast-mastery.md:113 | "After hit, Agility and Intellect are close: Agility is attack power and crit, Intellect is attack power through Careful Aim and mana on top. Crit is better than it looks because the pet copies all of your crit chance." | The page's table (bis/hunter-beast-mastery.json, band 60, raid, Alliance, re-ranked de044125c) now reads Hit 1.29, Agility 1.00, Critical strike 0.85, Intellect 0.46, Ranged attack power 0.30. Intellect is under half of Agility and below crit. The prose kept the pre-re-rank order (Int 1.01 ~ Agi 1.00 > Crit 0.77, engine 8baffe5fc); round 5's review pass rewrote only the Marksmanship paragraph. | "Hit comes first until you reach 9% hit, the cap against a raid boss, because a missed auto shot is the most expensive thing that can happen to you. After hit, Agility leads: it is attack power and crit together. Crit comes next, and it is better than it looks because the pet copies all of your crit chance. Intellect follows: Careful Aim turns it into attack power and it is mana on top, so take it when it comes on good mail, but do not trade Agility or crit for it. Plain attack power is the weakest per point, and haste is rare on gear but strong when you find it." |
| HUNTER-R6-2 | CONFIRMED | gear (enchants) | survival.md:110 | "Agility on gloves, boots and cloak, the best stats enchant on the chest, and Intellect on bracers. Intellect is a fair weapon enchant too: through Careful Aim it is attack power for you." | The enchant table on the same page (gear-extras.ts pickEnchants on bis/hunter-survival.json band 60 raid, Alliance) picks Enchant Bracer - Superior Strength (+9 Str, scored 10.9) over Greater Intellect (+7 Int, 8.7), and Enchant Weapon - Agility (+15 Agi, 35.5) on both weapons over Mighty Intellect (+22 Int, 27.2). Per point Strength 0.51 and Intellect 0.52 are level (both are one melee attack power a point for a hunter, Intellect through Careful Aim), so the bigger Strength enchant wins the bracer. The round 5 audit's "Mighty Intellect beats Agility" held only on the pre-re-rank weights (Int 0.87). The prose now tells a Survival hunter to enchant both slots against the table beside it. | "Agility on gloves, boots, cloak and both weapons, and the best stats enchant on the chest. On bracers take Superior Strength: Strength and Intellect each give you one attack power a point, Intellect through Careful Aim, and the Strength enchant gives nine to the Intellect one's seven." |
| HUNTER-R6-3 | LIKELY | voice | beast-mastery.md:89, marksmanship.md:91, survival.md:88 | "Elune's Light raises your critical strike chance with all spells and attacks by 10% for 15 sec. That is your burst for the fight." (BM; MM and SV carry the same pasted sentence) | Round 5's fix inserted the client tooltip text verbatim ("15 sec") between the burst press and what follows. On BM, "That is your burst for the fight" now points at Elune's Light, which a Horde Orc or Troll does not have. Every page's Races section already says "Elune's Light is 10% crit for 15 seconds", so the sentence repeats it in client wording. | Delete the Elune's Light sentence on all three pages. BM then reads "...press Bestial Wrath, Rapid Fire, your trinkets and your racial (Elune's Light, Blood Fury or Berserking) together. That is your burst for the fight." |
| HUNTER-R6-4 | LIKELY | voice | beast-mastery.md:66, marksmanship.md:65 | "Heart of the Lion (10% stats and attack power for everyone near you)" | Client text (spelltips 409580): "increasing total stats by 10% and Attack Power by 40 for all nearby allies", 280 at 60 (round 5 D2). The parenthetical reads as 10% attack power. index.md:34 and the Levels 1 to 19 paragraphs state it right. | "Heart of the Lion (10% stats and 280 attack power at 60 for everyone near you)" |

## Decisions

- **D1. Survival weapon enchant.** Agility on both weapons, as the table shows: 15 Agility is attack power and crit, which outweighs 22 Intellect's attack power on the current weights by about a third, a margin the nightly ranks will not flip unless Intellect doubles again. Crusader is not scored by the table (proc enchants carry no stats), so the prose does not name it until Sims requested 1 settles it.
- **D2. statPriority frontmatter order.** Not a finding: `railStatRows` sorts the shown rows by scale factor, so the BM and MM arrays listing Intellect before crit never reach the page.
- **D3. BM Careful Aim lines** (Strengths "Careful Aim turns Intellect into attack power for you and the pet", talent row "a mail Intellect piece is an attack power piece"). Kept: both describe the mechanic and make no ranking claim; HUNTER-R6-1's fix puts Intellect in its place.
- **D4. Consumables prose** naming a subset of the table (no Juju Might, Juju Power or stones). Kept: nothing stated is false, and the table is the full list.

## Sims requested

1. **Survival weapon enchant (carried from round 1 S5, never run).** hunter-survival raid build `5/005005001/500230231051222151`, Night Elf, band 60 raid preset, seeds 1|50001: Crusader on both weapons against Agility (+15) on both. If Crusader wins beyond error, survival.md:110 becomes "Crusader on both weapons, Agility if you cannot get it; Superior Strength on bracers ...".
2. **Hunter races on the re-ranked gear.** The race picks were measured before the nightly re-rank (de044125c) changed the BiS gear and the weights. Re-run the guide builds at band 60 raid, seeds 1|50001, on the current gear: Marksmanship Troll against Orc (round 2 margin 745.6 against 743.5, the closest Horde call), Survival Orc against Troll (prose says "level"), and Beast Mastery Night Elf against Human and Orc against Troll as a check. Change recommendedRaces only on a result beyond error.
3. **Carried, blocked on the engine (unchanged):** Sniper Shot needs a rotation line before the MM comparison means anything; Summon Hawk is not modeled.

## Sim findings

None new. Carried: no Sniper Shot line in the Marksmanship rotation; Summon Hawk unmodeled; non-stat buffs on pets blocked by RefusePlayerBuffs; the APL top-level notes still call Falcon 55/55 (notes only).

## Sources

- Forever Sixty client data 1.60.1.70755: talents/hunter.json, trainables/hunter.json, spelltips/hunter.json (409580, 20904, 2643, 3045, 20906, 14322, 25296, 1299447, 469145, 14266, 14271, 20910, 781, 1499, 14311, 13813, 14317, 14305, 1510, 14295, 19801, 1317257, 1299332, 14325), enchants.json, combos.json, races.json, addon-builds.json, raw/wowhead/items.json (quiver haste), bis/hunter-*.json (band 60 raid, current and at 3d4712981)
- Site code (what the page renders): web/src/lib/guides/gear-extras.ts (pickEnchants, loadGearExtras), web/src/lib/guides/class-dps.ts (factionForFirstRace), web/src/lib/bis/presets.ts (selectBand, raid preset first), web/src/lib/guides/rail-stats.ts
- sim/request/raid-consumes.json (hunter raid consumables)
- Engine d1d9f2775: sim/hunter/talents.go (mortalShots class mask includes Auto Shot)
- Race sims: design/reviews/guide-accuracy/round2/race-sims.md
- Blizzard, World of Warcraft Forever Class Deep Dives: Hunter and Druid, https://news.blizzard.com/en-us/article/24301515/world-of-warcraft-forever-class-deep-dives-hunter-and-druid (link check)
- Warcraft Tavern, Hunter Class Guide (Forever), https://warcrafttavern.com/forever/guides/hunter/ ; Hunter Leveling Guide (Forever), https://www.warcrafttavern.com/forever/guides/hunter-leveling/ (link check)

## Fix log (round 6)

New stat rule applied on every spec page: each `## Stat priority` paragraph now opens "The table above gives the order for your gear; here is why each stat counts." and gives only lasting reasons (cap rule for hit, Careful Aim, the pet copying crit, Mortal Shots, Predator's Edge and Savage Strikes, Lightning Reflexes), with no ranking words. `updated:` bumped to 2026-11-04 on beast-mastery.md:11, marksmanship.md:11 and survival.md:10. Prose test `-t hunter` passes clean.

- HUNTER-R6-1 (BM stat order): fixed under the new rule, not by re-ranking. beast-mastery.md:113. Marksmanship (marksmanship.md:116) and Survival (survival.md:102) reworded the same way, though they had no finding: "leads", "next", "then", "first" and "weakest per point" are gone.
- HUNTER-R6-2 / D1 (Survival enchants): fixed. survival.md:110 gives the lasting reason for gloves, boots, cloak (Agility: attack power and crit, Lightning Reflexes) and chest (stats enchant lifts every stat), then says the Strength, Intellect and Agility enchants on bracers and weapons are all attack power for you and sends the reader to the enchant table on the page. It no longer ranks Strength, Intellect and Agility, and it does not name Crusader (Sims requested 1 is still open).
- BM and MM enchant prose (same rule, no finding): the bracer line now gives a reason that does not change nightly. No bracer enchant carries real Agility (Minor Agility, +1, is the only one in enchants.json), and Strength does nothing for shots, so Intellect through Careful Aim is the bracer enchant. beast-mastery.md:121, marksmanship.md:124. Gloves, boots, cloak and chest are unchanged and match the table.
- HUNTER-R6-3 (burst line): fixed. Deleted the pasted Elune's Light tooltip sentence at beast-mastery.md:89, marksmanship.md:91 and survival.md:88. On BM "That is your burst for the fight" now follows the racial list again.
- HUNTER-R6-4 (Heart of the Lion wording): fixed. "10% stats and 280 attack power at 60 for everyone near you" at beast-mastery.md:66 and marksmanship.md:65. Survival's vaguer "more stats and attack power" at survival.md:64 now uses the same wording. Trueshot Aura in the same Brings lines (100 ranged attack power at 60, rank 5) already matched the client and is unchanged.
- D2 (statPriority frontmatter): left, not a page claim.
- D3 (BM Careful Aim lines): left. They describe the mechanic and do not rank.
- D4 (consumables subset): left. Nothing stated is false.
- Races (integrator note): race prose left as it is. Sims requested 2 waits for the final best-in-slot re-rank.
- Card notes: no changes needed.

New SIM findings: none.
