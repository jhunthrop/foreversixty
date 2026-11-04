# Guide accuracy round 7: Hunter

Auditor: hunter (class index, Beast Mastery, Marksmanship, Survival). Tree: gr7/audit at 394562f20 (land/guides-r6, all six rounds applied), client 1.60.1.70755, engine pin d1d9f2775 (sim/enginever/version.go).
Files: `web/src/content/guides/hunter/{index,beast-mastery,marksmanship,survival}.md`, card notes in `data/curated/apl/hunter-*.json`. Read first: round1 to round6 hunter reports with every fix log, round4/hunter-levels.md, round2 to round5 race-sims.md, ledger.md.

Since round 6 (`git diff 367cbadb7 HEAD`), the only hunter changes are the round 6 fixes: the stat paragraphs rewritten to explain without ranking, the BM and MM bracer reason, the Survival enchant paragraph, the burst-line Elune's Light sentence deleted, Heart of the Lion "a flat attack power bonus" in the Brings lines, and `updated:` bumped. No card notes, builds, addon-builds.json or BiS files moved for hunter.

## Builds and leveling order

All five build strings were decoded again against `talents/hunter.json` (tab order = tier, column). All are legal at 51 points, and every prerequisite is met: Intimidation needs Bestial Swiftness, Bestial Wrath needs Intimidation, Frenzy needs Ferocity 5, Mortal Shots needs Careful Aim 5, Counterattack needs Deterrence, and Lacerating Strikes needs Expose Prey 2.
- BM leveling 31/20/0: Deadly Aspects 5, Endurance Training 4, Focused Fire 2, Bestial Swiftness 1, Unleashed Fury 5, Ferocity 5, Intimidation 1, Bestial Discipline 2, Frenzy 5, Bestial Wrath 1; Lethal Attacks 5, Improved Stings 2, Efficiency 4, Careful Aim 5, Rapid Killing 1, Trueshot Aura 1, Mortal Shots 2.
- BM raid: the same, except Efficiency 3, no Rapid Killing and Mortal Shots 4. This matches "one point out of Efficiency and the point in Rapid Killing into Mortal Shots".
- MM leveling 16/30/5 and MM raid 5/33/13: these match the Talents paragraph, including "Survivalist is the way to Surefooted" (Improved Tracking 5 plus Survivalist 5 opens tier 2).
- SV 5/11/35: Lethal Attacks, Careful Aim and Lone Wolf in Marksmanship. Rapid Killing sits on Lone Wolf's row, and Lethal Attacks plus Careful Aim are exactly the ten points that open it, so "exactly the ten points that open it" holds.

The shipped order (`addon-builds.json`, index 0 = level 10) is unchanged since round 6. Every leveling sentence still matches it.

## What was checked and holds

**Round 6 edits re-checked.**
- Stat paragraphs: all three now give reasons without ranking words. Every reason is true on 70755:
  - The hit cap is 9%. The Survival special cap lines up with `hit_to_cap` specials 5 from baseline 4 (Surefooted 3 plus Dwarf 1).
  - The pet copies all of your crit (engine `makeStatInheritance`, Warcraft Tavern).
  - Careful Aim gives "Attack Power by 100% of your Intellect". It is generic AP, so it counts for melee too.
  - Mortal Shots raises the "critical strike damage bonus on all ranged abilities".
  - Lightning Reflexes gives 15% Agility.
  - For a hunter, Strength is one melee attack power per point.
  - The haste factor is large per 1% (BM 5.28, MM 6.21, SV 8.81 agility points), and `haste_on_items` is false, so "rare on gear but strong when you find it" holds.
- BM and MM bracer reason: `enchants.json` has only Minor Agility (+1) on bracers. The engine pet takes 10% of the hunter's higher attack power (`higherAttackPower`), which is ranged AP for a ranged hunter, so "Strength does nothing for your shots" holds.
- Survival enchant paragraph: Greater Stats lifts every stat. The paragraph names Strength, Intellect and Agility as attack power without ranking them and defers to the table.
- Burst lines: the pasted tooltip sentence is gone. "That is your burst for the fight" follows the racial list again.
- Heart of the Lion "10% stats and a flat attack power bonus": the client text is "Attack Power by 40", scaling to 280 at 60, which is flat and not a percentage. The wording holds.

**Ranking by nightly values elsewhere on the page.**
- The quest-gear line "Agility first" holds at every leveling band (20 to 50). Agility is the reference 1.00 and every other stat is lower on all three specs. Survival Intellect is 0.02 to 0.31 before Careful Aim at 55.
- The race orders are measured choices (round2/race-sims.md and the guide-build reruns). Their re-run waits for the final best-in-slot re-rank, as the brief says.
- "A two-hander with Agility or Intellect beats a fast one-hander with damage" is a stat-budget rule, not a nightly value.

**Spells and talents** were re-read against 70755 spelltips and rank text:
- Multi-Shot is instant on 6 s, shares its cooldown with Aimed Shot and hits 3 targets. That supports "fire Arcane Shot and Multi-Shot on the run".
- Aimed Shot is a 2 s cast, rank 6, 350.
- Rapid Fire is 40% ranged and melee speed for 15 s on 300 s. Bestial Wrath is 120 s.
- Disengage is on 5 s and "Disables auto-attack".
- Feign Death and Freezing Trap ("Any damage caused will break the ice") hold.
- Explosive Trap rank 3 hits all within 10 yd. Immolation Trap is rank 5. Volley is a channel.
- Sniper Shot is a 4 s cast. Trueshot Aura rank 5 is 100. Hunter's Mark rank 4 is ranged AP for all attackers.
- Aspect of the Hawk rank 6 is 110 and rank 7 (book) is 120. Beast rank 4 is 110, and Falcon equals the highest Hawk rank.
- Mend Pet is a channel. Revive Pet costs 100% of base mana, so "Mend Pet between pulls is cheaper" holds.
- Rapid Killing rank 1 is 1 min and 10%. Improved Stings is 6%, 13% and 20%. Mortal Shots is 6% a rank. Endurance Training, Bestial Discipline (50% at 2), Expose Prey (10% at 2, any Hunter's Mark), Predator's Edge (30% melee ability crit damage, 50% off hand), Savage Strikes (4% melee ability crit), Resourcefulness and Lacerating Strikes all hold.
- Trainer levels hold: Aspect of the Beast 30, Falcon 60, Cheetah 20, Feign Death 30, Lacerate 30, Tranquilizing Shot 60, Multi-Shot 18 and Wing Clip 12.
- Pet training still exists in Forever: Warcraft Tavern's pets guide says Beast Training unlocks after the taming chain and "Training Points ... can be spent teaching it new abilities". So "train it the top ranks of Claw, Bite and Growl" holds.

**Races.** The race texts in races.json hold: Elune's Light, Ruthless, Artisan's Precision 1% hit, Big Game Hunter, Blood Fury 2 min, Command 5% pet damage, Berserking 10% for 15 s on 3 min, and Skyborne Wind Blessed 2% haste. Hunter races in combos.json are Human, Orc, Dwarf, Night Elf, Tauren, Troll and both Skyborne, so "Humans can be Hunters for the first time" holds. Horde paladins exist (Undead, combos.json), so Judgement of Wisdom on the boss holds on both factions.

**Consumables and items.** The prose matches sim/request/raid-consumes.json:
- BM and MM at 49 take Major Mana Potion, Mongoose, Scorpok and Salmon.
- Survival at 55 takes Major Frenzy Potion (items.json 250943, required level 55), Mongoose and Salmon, and no mana potion.

Ammo and scopes in items.json:
- Jagged Arrow 40.
- Ice Threaded Arrow 51.
- Thorium Headed Arrow and Thorium Shells 52.
- Sniper Scope 40.
- Accurascope 50.
- Quickdraw Quiver 40.

**Card notes.** All BM, MM and Survival notes were re-read against the same spell data. None states anything false.

**Voice.** No em or en dashes, no sim, site, preset, band, pipeline or APL words, no quoted nightly numbers, and no removed spells or SoD runes.

## Source links

Every frontmatter URL was fetched this round (browser user agent):
- 200 for the Blizzard Hunter and Druid deep dive, the Deep Dive panel recap and the What's Next recap.
- 200 for Warcraft Tavern's Forever Hunter Guide, Leveling Guide and Hunter and Druid Deep Dive.
- 200 for Wowhead's Skyborne First Look.
- 200 for foreversixty.gg combos.json, races.json, spellconst/hunter.json, /planner, /planner?class=hunter and /sim/weights.
- 403 for Wowhead spell=1495 and spell=19306. These are unknown, not dead; round 5 matched both labels through the nether tooltip endpoint.

The body links /bis/hunter/<spec> follow the source-link rule. No /sources, wago or Talents Forever entries. Talent trees open the planner, and stat weights appear on damage pages only.

## Findings

| id | status | kind | file:line | wrong claim | evidence | fix |
|---|---|---|---|---|---|---|
| HUNTER-R7-1 | LIKELY | stats | survival.md:102 | "Crit is strong because Predator's Edge and Savage Strikes build on it." | Savage Strikes (talents/hunter.json) "Increases the critical strike chance of all your melee abilities by 4%". It adds crit chance, but it does not make a point of crit worth more, so it is not a reason crit is strong. The talents that do raise crit's value are Predator's Edge ("melee ability critical strike damage by 30%") and Resourcefulness ("melee and ranged ability Critical Strikes a 100% chance to allow 50% of your Mana Regeneration to continue while casting for 30 sec"). Round 6's rewrite turned "which Predator's Edge and Savage Strikes make strong" into a causal claim. | "Crit is strong because Predator's Edge makes your special crits hit 30% harder, and every special crit keeps half your mana regeneration running through Resourcefulness." |
| HUNTER-R7-2 | LIKELY | race | survival.md:116 | "Elune's Light is 10% crit for 15 seconds, and Predator's Edge makes every crit count" | Predator's Edge raises melee **ability** crit damage only ("Increases your melee ability critical strike damage by 30%"). White swings, which are most of a dual-wielder's hits, get nothing from it. "Every crit" overstates it. | "Elune's Light is 10% crit for 15 seconds, and Predator's Edge makes your special crits hit harder, so press it with Rapid Fire on the pull." |

No CONFIRMED finding this round. Every round 6 edit holds against the 70755 data and the page's own tables.

## Decisions

- **D1. "The enchant table on this page" (survival.md:110) and "The table above" (stat paragraphs).** Kept. These follow the round 6 owner-approved rule that the table ranks and the prose explains. Neither is sim or pipeline talk.
- **D2. Tauren is never named in the Horde race lines.** Kept. Tauren can be a hunter (combos.json race 6), but it trails every measured Horde run (round2/race-sims.md), and no line says Orc and Troll are the only choices.
- **D3. Survival "Strider Kick is a third melee special".** Kept. It reads as third after Raptor Strike and Mongoose Bite, the two the build is built around, and the priority list right below orders all of them.
- **D4. Heart of the Lion "a flat attack power bonus" in the Brings lines.** Kept. It is true (40 plus 4 a level, 280 at 60), and the Levels 1 to 19 paragraphs give the 280.

## Sims requested

1. **Carried from round 6, still unrun: Survival weapon enchant.** hunter-survival raid build `5/005005001/500230231051222151`, Night Elf, band 60 raid preset, seeds 1|50001, Crusader on both weapons against Agility (+15) on both. If Crusader wins beyond error, survival.md:110 names it.
2. **Carried, waiting for the final best-in-slot re-rank:** hunter race re-runs on the re-ranked gear:
   - MM Troll against Orc.
   - SV Orc against Troll.
   - BM Night Elf against Human, and Orc against Troll.
3. **Carried, blocked on the engine:** Sniper Shot needs a rotation line, and Summon Hawk is unmodeled.

## Sim findings

None new. Carried:
- No Sniper Shot line in the Marksmanship rotation.
- Summon Hawk is unmodeled.
- Non-stat buffs on pets are blocked by RefusePlayerBuffs.
- The APL top-level notes still call Falcon 55/55 (notes only, not rendered).

## Sources

- Forever Sixty client data 1.60.1.70755:
  - talents/hunter.json (Savage Strikes, Predator's Edge, Resourcefulness, Careful Aim, Rapid Killing, Improved Stings, Mortal Shots and Expose Prey rank text).
  - trainables/hunter.json and spelltips/hunter.json (2643, 20904, 3045, 19574, 781, 14273, 5384, 20906, 14325, 409580, 19801, 15632, 14295, 14317, 14311, 14305, 13163, 1310786, 1317257, 1299332, 20910, 14271, 14266, 14287, 25295, 469145, 14322, 25296, 1299447, 982, 13544, 5118).
  - races.json, combos.json, enchants.json (bracer enchants), items.json (ammo, scopes, quivers, potions, Major Frenzy Potion 250943), addon-builds.json, and bis/hunter-*.json (all bands; band 60 raid weights, haste_scale_factor, hit_to_cap).
- sim/request/raid-consumes.json (hunter raid consumables by level)
- Engine d1d9f2775: sim/hunter/pet.go (petAttackPowerShare 0.10 of the higher attack power; crit inherited)
- Race sims: design/reviews/guide-accuracy/round2/race-sims.md
- Warcraft Tavern, Hunter Pets Guide (Forever), https://www.warcrafttavern.com/forever/guides/hunter-pets/ (Beast Training after the taming chain; Training Points teach new abilities)
- Link check: Blizzard, World of Warcraft Forever Class Deep Dives: Hunter and Druid, https://news.blizzard.com/en-us/article/24301515/world-of-warcraft-forever-class-deep-dives-hunter-and-druid ; Warcraft Tavern, Hunter Class Guide (Forever), https://warcrafttavern.com/forever/guides/hunter/ ; Warcraft Tavern, Hunter Leveling Guide (Forever), https://www.warcrafttavern.com/forever/guides/hunter-leveling/

## Fix log (round 7)

Client text checked (talents/hunter.json, 70755): Savage Strikes rank 2 "Increases the critical strike chance of all your melee abilities by 4%"; Predator's Edge rank 5 "Increases your melee ability critical strike damage by 30% and your Off Hand weapon damage by 50%"; Resourcefulness rank 2 "grants your melee and ranged ability Critical Strikes a 100% chance to allow 50% of your Mana Regeneration to continue while casting for 30 sec". The Survival build takes all three at max rank.

- HUNTER-R7-1: fixed, survival.md:102. Predator's Edge is now the reason special crits pay off (30% harder), with Resourcefulness keeping half the mana regeneration on every special crit; Savage Strikes is stated as what it is, 4% added crit chance on melee specials, not a reason a point of crit is worth more.
- HUNTER-R7-2: fixed, survival.md:116. "makes every crit count" is now "makes your special crits hit harder"; white swings get nothing from Predator's Edge.
- Stat priority rule (no ranking words): BM, MM and SV Stat priority sections re-read and clean; prose test `-t hunter` passes.
- D1 to D4: kept as written in Decisions.
- `updated:` is 2026-11-04 on all three spec pages; index.md unchanged (2026-11-03).
- SIM findings: none new; the carried list above stands (Survival weapon enchant, race re-runs after the final re-rank, Sniper Shot line, Summon Hawk).
