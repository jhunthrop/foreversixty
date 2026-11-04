# Guide accuracy round 6: Rogue

- Class: rogue. I audited every line of `web/src/content/guides/rogue/index.md`, `assassination.md`, `combat.md` and `subtlety.md` in the round 6 audit tree (land/guides-r5 367cbadb7, all five rounds applied). I also read the rendered card notes (`rotation.priorityList[*].notes`) in `data/curated/apl/rogue-*.json`.
- Client 1.60.1.70755 data read: `talents/rogue.json` (every rank text), `trainables/rogue.json` (every rank level, cost and cooldown), `addon-builds.json`, `items.json`, `items/rogue.json`, and `bis/rogue-{assassination,combat,subtlety}.json` (band 60, raid preset, both factions: weights, `hit_to_cap`, slots). Engine pin d1d9f2775 (`sim/enginever/version.go`; `sim/core/spell_outcome.go` white table read for SIM-R6-1).
- Read first: the rogue reports for rounds 1 to 5 (with every fix log, including round 5's review pass), `round4/rogue-levels.md`, the race-sims files for rounds 2 to 5 (only round 3 has rogue rows) and `ledger.md` with the round 4 source-link rule.
- Rendering checked in code: the stat table, pills and rail card sort by band 60 raid scale factor at the faction of `recommendedRaces[0]` (`lib/guides/rail-stats.ts`, `lib/guides/class-dps.ts` `band60Weights`, `lib/bis/presets.ts` `selectBand`). All three rogue guides resolve to Alliance. The best in slot table has Alliance and Horde tabs (`GuideBisTable.astro`).
- Not re-filed:
  - SIM-1 to SIM-10 (SIM-3 is the Subtlety Slice and Dice line shadowed by Eviscerate), SIM-R2-1 to SIM-R2-3, SIM-R3-1, SIM-R3-2 and SIM-R5-1
  - the Mutilate off-hand engine change
  - the spelltips regen (R2-6, data lane)
  - card actions
  - the Combat leveling order (ladder-first)
  - the data-lane items in the round 6 brief (Alliance rows showing Frostwolf Clan, the "built in" hit baseline, BiS race captions)

## Decoded builds (re-verified, unchanged)

| Spec | build | raidBuild |
|---|---|---|
| Assassination | 38/11/2 | 36/10/5 |
| Combat | 14/34/3 (race field `gnome`) | 12/34/5 (race field `gnome`) |
| Subtlety | 12/8/31 (single build; its raid order equals it) | none |

The talent lists are identical to round 5's. Every build is legal on 70755 and equals the final state of its `addon-builds.json` order.

I decoded the shipped leveling orders again (index 0 = level 10). They are identical to rounds 2 to 5:
- **Assassination:** Improved Gouge 10-12, Remorseless Attacks 13-14, Malice 15-19, Ruthlessness 20-22, Improved Slice and Dice 23-24, Lethality 25-26, Vile Poisons 27-29 and 31-32, Venom 30, Cold Blood 33, Improved Poisons 34-38, Relentless Strikes 39 and 41-44, Mutilate 40, Seal Fate 45-47, Combat 48-58, Master of Deception 59, Opportunity 60.
- **Combat:** Coup de Grace 10-12, Efficient Strikes 13, Lightning Reflexes 14-18, Puncturing Wounds 19-21, Precision 22-24, Flawless Execution 25, Weapon Expertise 26-27, Aggression 28-30, Dual Wield Specialization 31-32, Blade Flurry 33, Hack and Slash 34-38, Flash of Steel 39 and 41-43, Adrenaline Rush 40, Improved Gouge 44-46, Remorseless Attacks 47-48, Malice 49-53, Ruthlessness 54-56, Improved Slice and Dice 57, Master of Deception 58-60.
- **Subtlety:** Master of Deception 10-12, Opportunity 13-17, Dirty Deeds 18, Camouflage 19, 24 and 27, Premeditation 20, Serrated Blades 21-23, Elusiveness 25 and 33, Ghostly Strike 26, Improved Ambush 28, Initiative 29 and 31-32, Hemorrhage 30, Preparation 34, Boot Knife 35-39, Death Mark 40, Assassination 41-52, Combat 53-60.

## The band 60 tables the stat paragraphs must follow

Raid preset, scale factors with Agility = 1.00 (haste is per 1%, outside the normalization):

| Spec (faction shown) | Order on the page |
|---|---|
| Assassination (Alliance) | Hit 1.99, Agility 1.00, Crit 0.82, Strength 0.59, Attack power 0.53, haste |
| Combat (Alliance; Horde identical) | Hit 2.39, Agility 1.00, **Strength 0.74, Attack power 0.67, Crit 0.52**, haste |
| Subtlety (Alliance) | Hit 1.05, Agility 1.00, Crit 0.83, Strength 0.57, Attack power 0.52, haste |

`hit_to_cap` (specials / white): Assassination 0.8 / 19.8 on both factions; Combat 1.8 / 20.8 on both; Subtlety Alliance **0 / 18.8**, Horde **6 / 25**.

## What checks out (fresh pass)

**"Levels 1 to 19", then "Levels 20 to 60" with the strip under it**
- Every training level matches `trainables/rogue.json`:
  - Sinister Strike, Eviscerate and Stealth at 1; Redirect at 1.
  - Backstab 4, Gouge 6, Evasion 8; Slice and Dice, Sprint and Sap 10; Kick 12; Garrote 14; Ambush and Poised 18.
  - Instant Poison 20, Shank 24 (rank 2 at 44), Cheap Shot 26, Deadly Poison 30, Kidney Shot 30, Mutilate 40.
- Every talent named at a level is in the shipped order at that level, with the rank text the prose quotes. Improved Gouge 3/3 is 1.5 s. Remorseless Attacks 2/2 is 40%. Malice is 5%. Coup de Grace is 20%. Efficient Strikes 1/2 is 3 energy. Lightning Reflexes is 5%. Puncturing Wounds names only Backstab and Mutilate. Master of Deception is "3 levels higher". Opportunity is 20%. Dirty Deeds 1/2 is 10 energy and removes the behind requirement. Camouflage 1/3 is 5% speed. Every row gate is met where the prose puts it.
- Poised 1328825 is 10 energy at full energy out of combat. Ambush costs 60 at every rank, so "Ambush from stealth costs 50" is right.
- Worn Dagger (2092) and Small Throwing Knife (2947) exist in `items.json`.
- Read in order:
  - Assassination's two blocks agree: dagger and Backstab first, Sinister Strike on mobs that face you, Mutilate at 40, Relentless Strikes at 39 to 44.
  - Combat's two blocks agree: Puncturing Wounds pays off only with a main-hand dagger, then you level on a sword with Sinister Strike, with Improved Gouge and Remorseless Attacks in the mid 40s.
  - Subtlety's two blocks are filed below (ROGUE-R6-5, R6-6).

**Talent, spell and build values elsewhere**
- Every value the prose quotes matches its 70755 rank or spell text. That covers Mutilate, Venom, Relentless Strikes, Lethality (200% to 220%), Vile and Improved Poisons, Vigor, Blood Splatter, Cold Blood, Coup de Grace (Eviscerate and Death Mark), Flash of Steel, Hack and Slash, Dual Wield Specialization, Weapon Expertise, Puncturing Wounds, Aggression, Flawless Execution, Blade Flurry, Adrenaline Rush, Serrated Blades, Hemorrhage, Ghostly Strike, Boot Knife, Death Mark, Initiative, Premeditation, Preparation and Elusiveness.
- Spell values match `trainables/rogue.json`, with costs and cooldowns at 60:
  - Kidney Shot: 25 energy, 20 s.
  - Kick: 25 energy. Gouge: 10 s.
  - Garrote: 50 energy. Shank: 5 energy, 10 s.
  - Death Mark: 25 energy, 10 s.
  - Ghostly Strike: 40 energy, 20 s. Hemorrhage: 35 energy.
  - Blade Flurry: 25 energy, 2 min. Adrenaline Rush: 5 min. Cold Blood: 3 min.
  - Premeditation: 2 min. Preparation: 10 min.
  - Blind: 5 min, so 3.5 min with Elusiveness 2/2.
- I recounted the build swaps:
  - Assassination trades 7 points for 7, Combat trades 11 for 11.
  - Combat's raid tier 3 gate is exactly 15. Subtlety's 12/8 split is legal (Lethality's tier 2 at 10 points, Precision's tier 1 at 5).
- Pull and reset arithmetic holds:
  - Combat: Garrote is 40 with Poised, plus a 60 Backstab.
  - Subtlety: the opener is four points and the Preparation reset is five.
  - Combat: a 3 minute kill gives two Blade Flurries and one Adrenaline Rush.

**Races, gear, sources, card notes, voice**
- The race prose matches `races.json` and `round3/race-sims.md:18-23`. Undead leads the Horde in every tree. Night Elf leads the Alliance for Assassination (a stated tie) and Subtlety, and Gnome leads it for Combat. Skyborne is named only as the paid option. No rogue spec changed builds since that run, so nothing needs re-measuring.
- Gear prose names pieces only for stable reasons, with one exception filed below:
  - Molten and Frozen Heart of the Mountain are `bonding: pickup` enchanting crafts.
  - Heartseeker is `bonding: equip` and blacksmith-made. Combat's line is conditional ("when it is on your list"), which is fine now that Heartseeker sits on the Assassination tables and not Combat's.
  - Devilsaur and Stormshroud are leatherworking crafts.
- Subtlety's weapon prose matches both tabs: sword main hand and mace off hand. Assassination's and Combat's "daggers at 60" match both tabs: every main and off hand on both factions is a dagger.
- The Assassination and Subtlety stat paragraphs give their tables' order. Combat's does not (R6-1).
- Source links: all six resolve (200), and each follows the round 4 rule. The Blizzard deep dive (24310968) still backs every claim the pages lean on it for: Expose Armor equal to 5 Sunders, poisons beside stones, Hack and Slash folding the weapon specializations together, Mutilate needing two daggers, Hemorrhage's dagger bonus and Rupture increase, and Dirty Deeds' Garrote from the front.
- Card notes: every rendered note states its own action truthfully. The Subtlety Slice and Dice note describes its action, and its shadowing by Eviscerate is SIM-3.
- Voice: no em or en dashes; no site, sim, APL, preset or band talk; no quoted nightly numbers.

## Findings

| id | status | kind | file:line | wrong claim | evidence | exact fix |
|---|---|---|---|---|---|---|
| ROGUE-R6-1 | CONFIRMED | stats (prose vs table) | combat.md:100 | "Crit is worth more to Combat than it looks, because Lethality makes your Backstab crits bigger." | `bis/rogue-combat.json`, band 60 raid, Alliance (the faction the page shows; Horde is the same): Strength 0.74 ± 0.001, Attack power 0.67 ± 0.001, **Crit 0.52 ± 0.017**. The page's table therefore reads Hit, Agility, Strength, Attack power, Crit, haste, with crit last of the per-point stats. The prose says the opposite. It also never says where Strength and attack power sit. The structural reason holds: Combat has no Seal Fate and none of the raid build's damage is a ticking bleed, so a crit only makes one hit bigger. Assassination (0.82, Seal Fate) and Subtlety (0.83, bleed ticks) rank crit third for exactly those reasons. | "Hit comes first: a missed Backstab is a wasted global and a lost combo point, and Precision only gets you part of the way. Agility gives attack power and crit together, which is why it beats everything after it. Strength and attack power come next, and crit after them: Combat has no Seal Fate and no bleeds, so a crit only makes one hit bigger, even with Lethality behind your Backstabs. Haste is rare on gear and good when you find it." Also reorder `statPriority` (combat.md:9) to `[Hit, Agility, Strength, Attack power, Critical strike, Melee haste]` so the no-weights fallback agrees. |
| ROGUE-R6-2 | CONFIRMED | stats (a round 5 edit) | subtlety.md:97 | "Your level 60 gear already puts Hemorrhage, Ghostly Strike and the finishers at the hit cap, so for them more hit adds nothing" | `bis/rogue-subtlety.json` band 60 raid `hit_to_cap`: Alliance specials 0, but **Horde specials 6** (white 25). The Horde tab of the same page's best in slot table leaves every special 6% short of the cap, so the sentence is false for a Horde reader, and Undead is the Horde pick. Whether a ranked set reaches the cap also moves with every nightly re-rank. The round 6 rule says stat prose keeps only the stable why. | "Hit first. Your white swings with two weapons miss far more often than Hemorrhage, Ghostly Strike and your finishers, so they keep wanting hit after your specials have stopped missing. Agility next, for attack power and crit together. Crit is better than it looks: your Rupture and Garrote ticks crit, and Death Mark's cap rises with your crit chance. Strength and attack power come after, and haste last." |
| ROGUE-R6-3 | CONFIRMED (minor) | gear (nightly claim) | assassination.md:115; combat.md:116 | "...bind-on-pickup trinkets only an enchanter can make, and both are pre-raid picks." | Band 60 raid Horde tables carry only one of the two Hearts: Assassination has Frozen plus Rune of the Guard Captain (19120), and Combat has Molten plus the Rune. The Alliance tables carry both. "Both are pre-raid picks" is a ranking that the nightly moves, and it is already false on the Horde tab. The bind-on-pickup and enchanter-only facts are stable (`items/rogue.json` 249469/249470 `bonding: pickup`). subtlety.md:113 ("bind-on-pickup pre-raid trinkets") makes no "both" claim and can stay. | assassination.md:115: "**Enchanting** makes Molten Heart of the Mountain and Frozen Heart of the Mountain, bind-on-pickup trinkets only an enchanter can make; the table above shows which of them your set wants." combat.md:116: "**Enchanting** is the one that changes your gear: Molten Heart of the Mountain and Frozen Heart of the Mountain are bind-on-pickup trinkets only an enchanter can make." (The rest of each line unchanged.) |
| ROGUE-R6-4 | LIKELY (minor) | stats (incomplete order) | assassination.md:99 | The paragraph stops at crit: "Hit first ... Agility next ... Crit is worth more to you than to most melee..." | The table reads Hit, Agility, Crit, Strength 0.59, Attack power 0.53, haste. The paragraph agrees as far as it goes, but the round 6 rule asks every stat paragraph to give the table's order. Subtlety's paragraph does. | Append: "Strength and attack power come after, and haste is rare enough that you take it when it comes." |
| ROGUE-R6-5 | LIKELY (minor) | leveling (block order) | subtlety.md:129 vs :117-125 | "### Levels 20 to 60 / Combat levels faster with a sword; respec at 60. If you level Subtlety, ..." | The new layout puts "Levels 1 to 19" first. It spends four paragraphs on how to level Subtlety from 1 (keep a dagger, Master of Deception 10-12, Opportunity 13-17, Dirty Deeds 18, Camouflage 19). The "level Combat instead" advice now arrives only at 20, after the reader has spent ten points in Subtlety. Read in order, the two blocks contradict each other. | subtlety.md:119 (first 1 to 19 paragraph) gets a lead sentence: "Combat levels faster with a sword, so if you only want Subtlety for raids, level Combat and respec at 60. If you level Subtlety, here is how it starts." subtlety.md:129 becomes: "Premeditation comes at 20, Ghostly Strike at 26, Hemorrhage at 30, Boot Knife from 35 and Death Mark at 40. Stealth and Camouflage keep you out of trouble on the way. Shank comes at 24 for every rogue: press it on cooldown. Open with Cheap Shot or Garrote, keep Slice and Dice up on anything that lives, and Eviscerate at five." |
| ROGUE-R6-6 | LIKELY (minor) | leveling (contradiction with the 60 text) | subtlety.md:119 | "Keep a dagger in the main hand: Backstab and Ambush need one, and they are what this tree builds on." | At 60 the same page says "Subtlety does not need daggers at 60" (:101) and "Ambush is not a boss opener for you" (:62, :70). Its builder is Hemorrhage, and its table holds a sword and a mace. The claim is true of levels 1 to 29 only, because Opportunity's Backstab, Garrote and Ambush bonus is what the early points buy, and Hemorrhage arrives at 30. | "Keep a dagger in the main hand while you level: Backstab and Ambush need one, and they are what your first talents build on until Hemorrhage at 30." |

## Decisions

- No build, race or card-action changes.
- Combat's stat order follows the table as measured: crit after Strength and attack power. The stated reason is structural (no Seal Fate, no bleeds), with no numbers. If tonight's re-rank put crit back above Strength, the sentence would need to change. SIM-R6-1 asks for the check that settles whether this is real or a modelling gap.
- Subtlety's hit paragraph keeps "Hit first", because Hit leads the table it shows. The reason becomes the stable one (white swings) instead of "your gear is capped", which is false on the Horde tab and moves nightly.
- Gear prose: profession lines keep the bind-on-pickup and enchanter-only facts and drop "both are pre-raid picks".

## Sims requested

- **SIM-R6-1 (check, Combat crit).** Combat raid 60, guide raid build, Alliance and Horde, using the band 60 set. Report the white-swing crit chance against the white crit cap (100 minus miss, dodge and glancing), and Backstab's crit chance. The goal is to confirm that crit's low scale factor (0.52, below attack power, while Assassination and Subtlety sit at about 0.82) comes from the white crit cap and the missing Seal Fate, and not from a gap in how Hack and Slash, Puncturing Wounds or Lethality are modelled. The prose fix above holds either way. Only the frontmatter order would follow a corrected engine.
- Carried for the sim round: the Mutilate off-hand penalty, SIM-1 to SIM-5, SIM-7, SIM-8, SIM-10, SIM-R2-1 to SIM-R2-3, SIM-R3-1, SIM-R3-2 and SIM-R5-1.

## Sim findings

- SIM-R6-1 above (a check, not yet a bug). The engine's white table (`sim/core/spell_outcome.go:345-373`, d1d9f2775) is a single roll: miss, dodge, glance, then crit. So a white crit cap is modelled, which is consistent with the low weight.
- None new on d1d9f2775 beyond that. SIM-3 (Subtlety Slice and Dice shadowed by the Eviscerate line, which fires only at 25 to 34 energy) is still open in `rogue-subtlety.json` priorities 5 and 6.

## Sources

- "World of Warcraft: Forever Class Deep Dives — Rogue and Warlock", Blizzard (read 2026-11-04, 200): https://news.blizzard.com/en-us/article/24310968/world-of-warcraft-forever-class-deep-dives-rogue-and-warlock
- "World of Warcraft: Forever Deep Dive Panel Recap", Blizzard (200): https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap
- "Skyborne First Look: New Neutral Race in World of Warcraft: Forever", Wowhead (200; body client-rendered): https://www.wowhead.com/forever/news/skyborne-first-look-new-neutral-race-in-world-of-warcraft-forever-382829
- foreversixty.gg/planner, /planner?class=rogue and /sim/weights (all 200)
- Client data (this tree): `data/builds/1.60.1.70755/bis/rogue-combat.json` and `rogue-subtlety.json` (band 60 raid weights and `hit_to_cap`, both factions), `bis/rogue-assassination.json`, `items/rogue.json` (249469, 249470, 12783, 19120), `talents/rogue.json`, `trainables/rogue.json`, `addon-builds.json`
- Round 3 race runs: `design/reviews/guide-accuracy/round3/race-sims.md:18-23` (this tree)

## Fix log (round 6)

- New stat rule, applied on every spec page: each `## Stat priority` paragraph now gives only lasting reasons (hit until your attacks stop missing and nothing past it, with the two-weapon white-swing reason; Agility for attack power and crit; the spec's crit mechanic; Strength and attack power as straight damage; haste rare on gear) and points to the table for the order. No ranking words. Prose test `-t rogue` passes.
- ROGUE-R6-1: fixed by the rule (combat.md:100). Crit's lasting reason stays without a rank: a crit makes one hit bigger, Lethality makes Backstab crits bigger, but Combat has no Seal Fate and no bleeds (true of the raid build, 12/34/5). `statPriority` fallback reordered to the current table, Hit, Agility, Strength, Attack power, Critical strike, Melee haste (combat.md:9).
- ROGUE-R6-2: fixed (subtlety.md:97). Hit keeps the white-swing reason; the "your gear already reaches the cap" claim is gone, and no line claims gear reaches the cap.
- ROGUE-R6-3: fixed (assassination.md:115, combat.md:116). Bind-on-pickup and enchanter-only facts kept; "both are pre-raid picks" dropped. Assassination points to the table for which Heart the set wants. subtlety.md:113 unchanged (no "both" claim).
- ROGUE-R6-4: fixed by the rule (assassination.md:99). Strength, attack power and haste now have their reasons, with no order.
- ROGUE-R6-5: fixed. The "level Combat and respec at 60" advice moved to the head of Levels 1 to 19 (subtlety.md:119); Levels 20 to 60 now starts at Premeditation and the doubled "Leveling Subtlety" is gone (subtlety.md:129).
- ROGUE-R6-6: fixed (subtlety.md:119). The dagger is for leveling, for the first talents, until Hemorrhage at 30.
- Other ranking by nightly value: none found outside Stat priority (poison pairing is a mechanic, left). Card notes have no stat ranking; none touched.
- `updated:` bumped to 2026-11-04 on assassination.md, combat.md and subtlety.md. index.md unchanged.
- SIM: no new findings. SIM-R6-1 (Combat crit check) and the carried list stand.
