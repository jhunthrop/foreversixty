# Guide accuracy round 8: Hunter

Auditor: hunter (class index, Beast Mastery, Marksmanship, Survival). Tree: gr8/audit at 09b98d9e4 (land/guides-r7, all seven rounds applied), client 1.60.1.70755, engine pin d1d9f2775 (sim/enginever/version.go).
Files: `web/src/content/guides/hunter/{index,beast-mastery,marksmanship,survival}.md`, card notes in `data/curated/apl/hunter-*.json`. Read first: round1 to round7 hunter reports with their fix logs, round4/hunter-levels.md, the race-sims files and ledger.md.

Since round 7 (`git diff 394562f20 HEAD`), the only hunter changes are the round 7 fixes: the Survival crit sentence in Stat priority (survival.md:102), the Survival Night Elf line (survival.md:116), and `updated:` on index.md. Builds, card notes, addon-builds.json and BiS files for hunter did not move.

## Builds and leveling order

I decoded all five strings again against `talents/hunter.json`, with talents ordered by tier, then column. All five are legal at 51 points, and every prerequisite is met.
- BM: 31/20/0. The raid build moves one Efficiency point and the Rapid Killing point into Mortal Shots (2 to 4).
- MM leveling: 16/30/5.
- MM raid: 5/33/13.
- SV: 5/11/35.

The shipped order is unchanged. I re-decoded it from `addon-builds.json`, where index 0 is level 10:
- BM: Deadly Aspects 10-14, Focused Fire 15-16, Endurance Training 17-20, ... Bestial Wrath 40.
- MM: Lethal Attacks 10-14, Improved Stings 15, Efficiency 16-20.
- SV: Improved Tracking 10-14, Savage Strikes 15-16, Survivalist 17-19, ... Lone Wolf 60.

Every leveling sentence matches it, and so does each talent's first trainable level:
- Intimidation (30) at 32.
- Bestial Wrath (40) at 40.
- Trueshot Aura (25) at 26.
- Strider Kick (30) at 30.
- Counterattack (30) at 32.

## What was checked and holds

- **Round 7 edits.**
  - survival.md:116 now says "makes your special crits hit harder". That is true: Predator's Edge covers melee abilities only.
  - survival.md:102's Resourcefulness clause and its Savage Strikes clause both hold against the client text ("50% of your Mana Regeneration to continue while casting for 30 sec"; "critical strike chance of all your melee abilities by 4%").
  - The "30% harder" clause in the same sentence is wrong (HUNTER-R8-3).
  - index.md `updated: 2026-11-04` is right: round 5 changed the index's content on 2026-11-04 (2504f4a69).
- **Spells and trainer levels.** I checked all of these against `trainables/hunter.json` and `spelltips/hunter.json`:
  - Trainer levels: Heart of the Lion 1, Raptor Strike 1/8/16, Serpent Sting and Monkey 4, Arcane Shot and Hunter's Mark 6, Concussive Shot 8, Hawk and Tame Beast 10, Track Humanoids 10, Wing Clip and Mend Pet 12, Mongoose Bite 16, Multi-Shot and Track Undead 18, Cheetah 20, Feign Death, Beast and Lacerate 30, Falcon and Tranquilizing Shot 60.
  - Spell values: Aimed Shot is a 2 s cast on a 6 s cooldown it shares with Multi-Shot. Multi-Shot hits 3 targets. Rapid Fire is 40% ranged and melee speed on a 300 s cooldown. Trueshot Aura is 100. Hawk is 110, with 120 as the book rank, and Beast is 110. Disengage is on 5 s and "Disables auto-attack". Wing Clip rank 3 deals 50.
  - Heart of the Lion's buff 409583 is a 15 s aura that the hunter's 409580 keeps refreshing, so "keep it up" holds.
- **Talent values.** I checked every talent value against the 70755 rank text. Rapid Killing rank 1 is 1 min and 10%, and Improved Stings is 6%, 13% and 20%. Everything else holds except the crit-damage wording below.
- **Races.** Every racial named on the pages matches its text in races.json.
- **Card notes.** I re-read every BM, MM and Survival note. Survival's priority prose matches the APL order: Raptor Strike, Counterattack, Mongoose Bite, Lacerate, Strider Kick, Immolation Trap, then Wing Clip under Windfury only.
- **Consumables, ammo and enchant prose.** All unchanged since round 7, and they still match the tables, which did not move.
- **Voice.** No em or en dashes, no sim, site or pipeline words, no quoted nightly numbers, no ranking words in Stat priority, and no removed spells or SoD runes. Hydra Shot, Wyvern Strike and Widow Bite are unlearnable class_spell rows, and no page names them.

## Source links

I fetched every URL with a browser user agent:
- 200 for the three Blizzard articles.
- 200 for the three Warcraft Tavern pages.
- 200 for foreversixty.gg combos.json, races.json, spellconst/hunter.json, /planner, /planner?class=hunter and /sim/weights.
- 403 for Wowhead spell=1495, spell=19306 and the Skyborne First Look article. Round 7 got a 200 on Skyborne. All three count as unknown, not dead.

Every entry still follows the source-link rule.

## Findings

All three findings are the error class round 7 fixed for warrior (WARRIOR-R7-1, Impale). The talent raises the critical strike damage **bonus**, which is the +100% part of a crit, so the prose doubles what the talent actually does.

Hard evidence:
- `spellconst/hunter.json` gives Mortal Shots (19485) and Predator's Edge (1310627) the same effect as Impale (warrior 16493): aura 108 (add-percent spell modifier), misc_value 15 (critical damage bonus), amount 30.
- The Mortal Shots client text reads "Increases the critical strike damage bonus on all ranged abilities by 30%".
- In the engine (d1d9f2775), both talents add 0.06 a point to the spell's CritDamageBonus (`sim/hunter/talents.go:322` `mortalShots`, `:332` `predatorsEdgeCritDamage`). `sim/core/spell_outcome.go:1085` turns that into a crit multiplier of `1 + (2.0 - 1) x CritDamageBonus`, which is 2.30 at 5 points.
- A crit therefore goes from 200% to 230% of a normal hit. That is 15% harder, not 30%.
- Classic gives the same answer: Mortal Shots 5/5 takes crits from 200% to 230% (Warcraft Wiki, Mortal Shots).

| id | status | kind | file:line | wrong claim | evidence | fix |
|---|---|---|---|---|---|---|
| HUNTER-R8-1 | CONFIRMED | talent | beast-mastery.md:79 | "**Mortal Shots** (2 points): 12% more crit damage on your shots, 24% in the raid build at 4 points" | At 2 points a crit is 212% of a normal hit (6% harder than 200%); at 4 points it is 224% (12% harder). See the evidence above. | "**Mortal Shots** (2 points): your shot crits hit for 212% of a normal hit instead of double, 224% with the 4 points in the raid build" |
| HUNTER-R8-2 | CONFIRMED | talent | marksmanship.md:71 | "**Mortal Shots** (5 points): 30% more crit damage on your shots" | At 5 points a crit is 230% of a normal hit, 15% harder. See the evidence above. | "**Mortal Shots** (5 points): your shot crits hit for 230% of a normal hit instead of double" |
| HUNTER-R8-3 | CONFIRMED | talent (line 102 is a round 7 edit) | survival.md:74, survival.md:102 | :74 "**Predator's Edge** (5 points): 30% more crit damage on your melee specials and 50% more off-hand damage"; :102 "Predator's Edge makes their crits hit 30% harder" | Predator's Edge's crit half is the same aura 108 / misc 15 / 30 effect as Mortal Shots and Impale, and the engine adds it to CritDamageBonus on Raptor Strike, Mongoose Bite, Counterattack, Lacerate and Lacerating Strikes. A melee special crit is 230%, not 260%. Its tooltip omits the word "bonus", but the spell data does not change. Round 7's fix (HUNTER-R7-1) turned this into an explicit "30% harder". | :74 "**Predator's Edge** (5 points): your melee special crits hit for 230% of a normal hit instead of double, and your off hand does 50% more damage". :102 "Crit pays off on your melee specials: Predator's Edge makes their crits hit for 230% of a normal hit instead of double, and every special crit keeps half your mana regeneration running through Resourcefulness." |

These lines hold and need no change:
- marksmanship.md:116, "makes your ranged crits hit harder".
- survival.md:116, "makes your special crits hit harder".

### SHOULDs (wording, not findings)

- The BM and MM Serpent Sting card notes say "Refreshing any earlier throws away a tick". With under 1.5 s left, a refresh can still cost the last tick, depending on how a refresh handles a pending tick. Nothing in the 70755 data settles it, so it is not filed. "Refresh it as it drops" would avoid the claim.
- beast-mastery.md:109 and marksmanship.md:112 say "Lucifron's tome". The Tome of Tranquilizing Shot drops across Molten Core, and "the Molten Core tome" would be more exact. It does not change the advice.

## Decisions

- **D1. Crit-damage wording.** A crit's size goes on the page as a percent of a normal hit ("230% of a normal hit instead of double"), the same form the warrior pages now use for Impale. That keeps the talents consistent across classes and cannot be misread as a multiplier.
- **D2. Predator's Edge despite the tooltip.** Its tooltip says "critical strike damage by 30%" with no "bonus". Following the round 1 lesson, I checked the effect itself (aura 108, misc 15). It is the same spell modifier as Mortal Shots and Impale, and the engine applies it the same way, so the page follows the spell data.
- **D3. Round 7's D1 to D4** stand as written.

## Sims requested

None new. Carried items, none of them claimed in prose:
1. Survival weapon enchant: Crusader on both weapons against Agility on both, on the guide build `5/005005001/500230231051222151`, Night Elf, band 60 raid, seeds 1|50001.
2. Hunter race re-runs after the final best-in-slot re-rank: MM Troll against Orc, SV Orc against Troll, BM Night Elf against Human, and BM Orc against Troll.
3. Blocked on the engine: Sniper Shot needs a rotation line, and Summon Hawk is unmodeled.

## Sim findings

None new. The engine's crit-damage-bonus handling for Mortal Shots and Predator's Edge matches the client, so R8-1 to R8-3 are guide errors only. Carried:
- No Sniper Shot line in the Marksmanship rotation.
- Summon Hawk is unmodeled.
- Pet non-stat buffs are blocked by RefusePlayerBuffs.
- The APL top-level notes still call Falcon 55/55 (notes only).

## Sources

- Forever Sixty client data 1.60.1.70755:
  - `spellconst/hunter.json`: 19485 Mortal Shots and 1310627 Predator's Edge, both aura 108 / misc 15 / 30; 409583 Heart of the Lion, 15 s.
  - `spellconst/warrior.json`: 16493 Impale, the same effect.
  - `talents/hunter.json`, `trainables/hunter.json`, `spelltips/hunter.json`, `races.json`, `addon-builds.json`.
- Engine d1d9f2775:
  - `sim/hunter/talents.go:319-334`: mortalShots and predatorsEdgeCritDamage.
  - `sim/hunter/hunter.go:307`: Auto Shot takes Mortal Shots.
  - `sim/hunter/{aimed_shot,arcane_shot,mongoose_bite,counterattack,lacerate,lacerating_strikes}.go`: CritDamageBonus.
  - `sim/core/spell_outcome.go:1083-1085`: the crit multiplier.
- Round 7 warrior precedent: design/reviews/guide-accuracy/round7/warrior.md, WARRIOR-R7-1.
- Warcraft Wiki, Mortal Shots: https://warcraft.wiki.gg/wiki/Mortal_Shots (crit damage bonus, 200% to 230% at 5/5).
- Wowhead Classic, Mortal Shots: https://www.wowhead.com/classic/spell=19485.
- Link check: the frontmatter URLs on all four pages, as listed above.

## Fix log (round 8)

- HUNTER-R8-1 fixed, beast-mastery.md:79: 212% of a normal hit, 224% at 4 points (talents/hunter.json Mortal Shots rank text "critical strike damage bonus ... by 12%/24%").
- HUNTER-R8-2 fixed, marksmanship.md:71: 230% of a normal hit.
- HUNTER-R8-3 fixed, survival.md:74 and :102: 230% of a normal hit; off hand 50% kept (D2: same aura 108 / misc 15 as Mortal Shots).
- D1 to D3 applied as above. SHOULDs (Serpent Sting refresh note, "Lucifron's tome") left: wording only, not findings.
- Prose check: hunter clean. No new SIM findings.
