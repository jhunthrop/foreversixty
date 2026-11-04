# Guide accuracy round 6: Druid

Auditor: druid persona. Tree: gr6/audit at 367cbadb7 (land/guides-r5, all five rounds applied; READ-ONLY), client 1.60.1.70755, engine pin d1d9f2775 (sim/enginever/version.go).

Specs audited: class page (`index.md`), Balance (`balance.md`), Feral cat (`feral.md`), Feral bear (`feral-bear.md`), Restoration (`restoration.md`). Before starting I read the round 1 to 5 druid reports with every fix log, round4/druid-levels.md, round2 to round5 race-sims.md and ledger.md. Since round 5's audit head (42f10dda9) the druid changes are the round 5 fixes (commit 2504f4a69: bear BiS source, bear Thick Hide and stat paragraph, bear Furor line, cat Feral Swiftness line, class-page capstone line, Restoration Swiftmend, Improved Moonfire and Elune's Light lines) and the nightly BiS re-rank (de044125c), which rewrote all four `bis/druid-*.json` files. No commit touched addon-builds.json, talents/druid.json, trainables/druid.json, spelltips or the druid APLs.

## Decoded builds (unchanged since round 1, all legal, 51 points)

Same as round4/druid.md. Balance `5232220115101151/0/555301`, raid `5232220115501351/0/505003`. Feral `01/50212032123012123001/535033`, raid `05002/50230032023011125001/05503`. Bear `0/5523033202310212551/051`. Restoration `05312001/0/55530051531132`. Shipped leveling orders as decoded in round5/druid.md (unchanged file); every Levels 1 to 19 paragraph and every 20 to 60 talent claim still matches them.

## What the stat tables render now

The guide page's table and pills read `band60Weights` (web/src/lib/guides/class-dps.ts): band 60, the faction of the first `recommendedRaces` entry (Night Elf, so Alliance), raid preset; rows sorted by scale factor (rail-stats.ts), per-1% haste after the per-point stats, "Not significant" last. From the re-ranked files (generated 2026-11-04T02:22 to 02:56Z):

| Spec | Rendered order |
|---|---|
| Balance | Spell power 4.97, Hit 4.41, Arcane power 3.19, Crit 2.92, Nature power 1.78, Intellect 1.00, Spell haste (per 1%), then MP5, Spirit, Spell penetration not significant |
| Feral | Hit 1.02, Agility 1.00, Strength 1.00, Crit 0.71, Attack power 0.41, Feral attack power 0.41, Melee haste (per 1%) |
| Bear | Stamina 1.66, Agility 1.00, Dodge 0.99, Hit 0.62, Armor 0.39, Strength 0.16, Attack power 0.07, then Defense and Crit not significant |
| Restoration | Spirit 1.27, MP5 1.13, Intellect 1.00, Healing power 0.91, Crit 0.61 |

Before the re-rank (round 1 to 5 audits), Balance had Nature above Crit, the cat had Agility first, and Restoration had MP5 first. Round 5's review pass rewrote only the bear paragraph.

## Findings

| id | status | kind | file:line | wrong claim | evidence | exact fix |
|---|---|---|---|---|---|---|
| DRUID-R6-1 | CONFIRMED | stats | feral.md:89 | "Agility is crit and attack power at once, which is why it leads." | The page's table now puts Hit first, then Agility and Strength level with it, then Crit (bis/druid-feral.json band 60 alliance raid: hit 1.020, agility 1.000, strength 0.998, crit 0.708). The paragraph also gives hit last and never places crit. | Keep the weapon sentences, then: "Hit, Agility and Strength sit together at the top. Hit is first: a missed or dodged Shred hands back most of its energy, but you still lose the global and the combo point, and once your gear reaches the hit cap more adds nothing. Agility is crit and attack power at once. Strength is attack power, and Heart of the Wild adds 6% to it (10% with the raid build's five points). Crit comes next, and attack power from gear is last point for point." |
| DRUID-R6-2 | CONFIRMED | stats | restoration.md:86 | "...which is why MP5 and Spirit lead. Healing power is every heal, and your HoTs take a big share of it per tick. Spirit is mana ... MP5 is mana that never stops. Intellect is your pool and a little crit." | Table order: Spirit, MP5, Intellect, Healing power, Crit (bis/druid-restoration.json band 60 alliance raid: spirit 1.274, mp5 1.129, intellect 1.000, healing_power 0.909, crit 0.606). The paragraph puts MP5 before Spirit and Healing power before Intellect, and never places crit. | "A long tank fight runs you out of mana before it runs out of time, so the mana stats lead. Spirit is first: it is mana through Reflection and Innervate, and it is your Bloom heal. MP5 is next, mana that never stops. Intellect is your pool and a little crit. Healing power comes after them: it is every heal, and your HoTs take a big share of it per tick, but on a long fight an empty bar heals nothing. Crit is last, and mostly lands on your Regrowths. Spell power with no healing on it adds nothing to your heals." |
| DRUID-R6-3 | CONFIRMED | stats | balance.md:105 | "Arcane damage is nearly as good as spell power, because Starfire and Moonfire are Arcane; Nature damage only feeds Wrath and Insect Swarm. Crit pays twice, through Vengeance and Nature's Grace." | Table order: Spell power, Hit, Arcane power, Crit, Nature power, Intellect (bis/druid-balance.json band 60 alliance raid: 4.97, 4.41, 3.19, 2.92, 1.78, 1.00). Since the re-rank Crit sits above Nature power, so the paragraph names Nature before Crit against the table. "Nearly as good as spell power" is a size claim that moves nightly (Arcane is now below Hit as well). | "Spell power is damage on every cast. Hit is worth more to you than to most casters: a resisted Starfire is a 3 second cast for nothing, and a resisted Wrath banks no Eclipse charges. Arcane damage comes next, because Starfire and Moonfire are Arcane. Crit pays twice, through Vengeance and Nature's Grace, so it sits above Nature damage, which only feeds Wrath and Insect Swarm. Intellect is your mana pool and a little crit." |
| DRUID-R6-4 | LIKELY | stats (round 5 edit) | feral-bear.md:90 | "Defense lowers your chance to be crit and gives Thick Hide a little armor, but at 60 neither moves your survival much" | The order (Defense last) matches the rendered table, but the reason does not hold at 60: Defense is the second stat on the bare row (1.89, above Stamina 1.82) and the first on the Horde raid row (1.66), both significant; only the Alliance raid row reads "Not significant" (0.89, error 0.30). The verdict on Defense swings from row to row, so "neither moves your survival much" is a nightly-moving claim, not a stable reason. | "Defense and crit come last. Defense lowers your chance to be crit and gives Thick Hide a little armor, and how much that is worth swings with your gear and your healers, so take it when it comes on a piece you want anyway." |
| DRUID-R6-5 | LIKELY | gear | feral.md:93 | "It is a level 40 blue with little on it besides the bonus, so it is a small upgrade while leveling" | Wolfshead Helm (items/druid.json 8345) is 109 armor and 10 Spirit, BoE. The engine knows the item (sim/druid/items.go `WolfsheadHelm = 8345`), and the page's own leveling BiS never takes it: band 40 and 50 pick White Bandit Mask, band 60 Outlaw's Collar (bis/druid-feral.json, Alliance). "An upgrade" is a ranking claim that moves nightly, and it does not match the page today. The druid-only effect and the Leatherworking pattern are the stable reasons to name it. | "Wolfshead Helm is druid-only in Forever and now gives 5 more energy on every Shifting Power (and 5 more rage on Enrage). Leatherworkers make it. It is a level 40 blue with only 10 Spirit beside the bonus, so you trade a helm's real stats for the energy, and at 60 helms with real stats beat it." |

Checked and right (no finding):

**Round 5 edits re-checked**
- Bear stat paragraph order (Stamina, Agility, Dodge, Hit, Armor, Strength and attack power, Defense and crit last) still matches the re-ranked table; only the Defense reason is filed (R6-4). Thick Hide "for every point of Defense over five times your level" matches the rank text (0.67 to 2.00 base armor a point). The Furor line matches the rank texts (20% to 100% chance of 10 Rage). The bear BiS source opens the Feral Bear BiS page (200).
- Feral Swiftness: 4% dodge at 2 points and 30% Cat Form speed; Cat Form trains at 20 (trainables/druid.json 768).
- "the cat's capstone in Feral" (index.md:30): Shifting Power and Primal Bite are both tier 6.
- Restoration: Swiftmend "an extra heal" (it costs 20% base mana, stated at :72); Improved Moonfire named a dead point with Improved Regrowth 1 (10% Regrowth crit) as the small gain; Elune's Light on Regrowths and Healing Touches.

**Levels 1 to 19, then 20 to 60 (reading order)**
- Balance, Restoration: the 1 to 19 block ends on gear, the 20 to 60 block starts on how the spec quests. No repetition, no contradiction.
- Feral: the 1 to 19 block ends as a caster with Bear Form; the 20 to 60 block opens "Cat Form arrives at 20", which follows the Feral Swiftness line ("once you learn Cat Form at 20").
- Bear: "Dire Bear Form comes at 40" no longer repeats Bear Form at 10; "until then" after "Lacerate arrives at 25 and Primal Bite at 40" reads right.
- Trainer levels re-read: Wrath 1/6/14, Moonfire 4, Entangling Roots 8, Nature's Grasp 10, Faerie Fire 18, Healing Touch 1, Rejuvenation 4, Regrowth and Revive 12, Cure Poison 14, Bear Form, Maul, Growl and Demoralizing Roar 10, Enrage 12, Bash 14, Swipe 16, Maul 2 at 18, Aquatic Form 16, Frenzied Regeneration 36, Lacerate 25, Primal Bite and Dire Bear Form 40, Innervate 40, Abolish Poison 26, Remove Curse 24.

**Talent and spell texts re-read (talents/druid.json 70755)**
- Nature's Focus covers Arcane and Nature spells (70%; 28% at 2), so it protects Healing Touch and Regrowth. Natural Instinct is melee crit damage plus spell healing from Intellect. Nature's Swiftness is "your next Nature spell". Improved Tranquility removes all threat and 60% of the cooldown. Swiftmend heals "the full duration of the periodic effect". Predatory Strikes is 150% of level. Natural Shapeshifter is 30% and reaches Shifting Power. Blood Frenzy 2/2 is 5 rage on every bear crit and a combo point on every builder crit. Bloom is 100% Spirit on a natural expiry. Wild Growth heals the target's party over 7 s; the Wild Growth swap still meets every tier gate (37 Restoration points before Wild Growth).
- Balance: Vengeance covers Arcane and Nature spells; Moonkin Form doubles Omen and blocks only healing spells; Nature's Reach is 4% hit and 20% range; Eclipse 0.5 s on 2 Starfires, 4 charges.

**Gear and card notes**
- Bear Wolfshead line is the stable reason only (druid-only, crafted). Every druid APL note re-read; none states anything false.

**Sources**
- Blizzard class deep dive and panel recap: 200. Talents Forever data.json: 200. /sim/weights, /sim/drops, /bis/druid/feral-bear, /bis/druid/restoration: 200. Wowhead Skyborne First Look: 403 to curl (fetched and title-matched in rounds 4 and 5), so unknown, not dead. Every entry follows the ledger's source-link rule: weights and boss drops only on Balance and Feral, own BiS page on the bear and Restoration, no /sources, no "how X plays", Talents Forever labelled as beta client data.
- No em or en dashes, no site, sim, APL, preset or band words in prose, no quoted nightly numbers.

## Decisions

1. **Stat prose follows the rendered table, Alliance raid row.** The table reads the first recommended race's faction (Night Elf), so the prose follows that row. Where two stats sit together at the top (cat Hit, Agility and Strength), the prose says so and gives the table's order, with no numbers.
2. **Defense on the bear stays last, with a reason that does not move** (R6-4). The rows disagree on Defense, so the line says its worth depends on gear and healing instead of claiming it does little.
3. **Wolfshead Helm keeps its place in the gear prose** for its stable reasons (druid-only effect, Leatherworking). The ranking claim goes (R6-5).
4. **Frontmatter statPriority is not touched.** The pills and table follow the weights; the frontmatter is only the label list.

## Sims requested

1. Carried from round 4 (S10): druid-feral leveling, Claw to Shred switch, ladder levels 20 to 25 with the Shred line gated on Shredding Attacks 3/3 (level 22), against the current Ferocious Bite rank 5 gate.
2. Carried from round 3: druid-restoration Improved Moonfire 1 into Improved Regrowth 1 (`05302001/0/555300515311321`), heal profile at band 60 raid, seeds 1 and 50001, with a ladder-first row that keeps Improved Regrowth.

## Sim findings

- S1 (carried): `rendAndTearDamagePerRank = 0.02` against the client's 3% a rank.
- S7, S8, S9, S10 and SIM-D1 (carried): Nature's Reach melee hit for the feral raid re-run, Feral Charge (Cat) unmodeled, the Pounce behind-target check, the Claw gate, the feral leveling Nature's Focus and Subtlety points.
- S11 (new, data lane): the bear weights still carry expertise at the top of every band 60 row (15.12 Alliance raid), though no druid item has it; the page does not render it because the frontmatter has no Expertise label. Already the round 1 S4 item; noted only because the re-rank kept it.
- S12 (new, noise): bear Defense flips between first (Horde raid, bare) and not significant (Alliance raid) at band 60. Worth a higher-iteration weights run for the tank rows before the prose ever ranks Defense again.

## Sources

- Client 1.60.1.70755: bis/druid-balance.json, druid-feral.json, druid-feral-bear.json, druid-restoration.json (band 60 weights, generated 2026-11-04, de044125c; head, weapon and idol picks at every band), talents/druid.json (all rank texts), trainables/druid.json, items/druid.json (8345), addon-builds.json; APLs data/curated/apl/druid-*.json.
- Site code: web/src/lib/guides/class-dps.ts (`band60Weights`), web/src/lib/guides/rail-stats.ts (row order), web/src/lib/bis/presets.ts (`selectBand`, raid preset first), web/src/lib/bis/weight-reading.ts (not significant rule).
- Engine d1d9f2775: sim/druid/items.go (Wolfshead Helm modeled).
- Blizzard, "World of Warcraft: Forever Class Deep Dives: Hunter and Druid", https://news.blizzard.com/en-us/article/24301515/world-of-warcraft-forever-class-deep-dives-hunter-and-druid (200, 2026-11-04).
- Blizzard, "World of Warcraft: Forever Deep Dive Panel Recap", https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap (200).
- Wowhead, "Skyborne First Look", https://www.wowhead.com/forever/news/skyborne-first-look-new-neutral-race-in-world-of-warcraft-forever-382829 (403 to automated fetch; unknown).
- Talents Forever, https://talentsforever.com/data.json (200).
- Forever Sixty: https://foreversixty.gg/sim/weights, https://foreversixty.gg/sim/drops, https://foreversixty.gg/bis/druid/feral-bear, https://foreversixty.gg/bis/druid/restoration (all 200).

## Fix log (round 6)

New stat rule applied on all four spec pages: each Stat priority paragraph now opens by pointing to the table for the order and gives only each stat's lasting reason, with the cap rule for hit. No ranking words; `_prose.test.ts -t druid` passes. `updated:` bumped to 2026-11-04 on balance.md, feral.md, feral-bear.md, restoration.md (index.md untouched).

- DRUID-R6-1 fixed, feral.md:89. "which is why it leads" and the hit-last ordering removed. Agility, Strength (Heart of the Wild 6%/10%), Hit (lost global and combo point; counts up to the cap, nothing past it), Crit (bigger hits, combo points through Blood Frenzy, which both cat builds take per feral.md:60), gear attack power (no crit attached). Weapon sentences kept, reworded to drop "per second" and "a second" (test matches `second`): "Weapon DPS matters", "paw swings at a fixed 1.0 speed".
- DRUID-R6-2 fixed, restoration.md:86. "MP5 and Spirit lead" replaced by "the mana stats count so much for you"; reasons kept for Healing power, Spirit, MP5, Intellect; crit added as bigger direct heals (no HoT crit claim).
- DRUID-R6-3 fixed, balance.md:105. "nearly as good as spell power" and "worth more to you than to most casters" (matches the `worth more than` rule) removed; hit gains the spell hit cap rule; "3 second cast" reworded to "3 seconds of casting".
- DRUID-R6-4 fixed per integrator, feral-bear.md:90. "Stamina comes first", "Agility is next", "Defense and crit come last" and "neither moves your survival much" removed. Defense: "fewer crits taken, and Thick Hide turns it into a little armor", no worth claim. Crit: threat and rage through Blood Frenzy (feral-bear.md:61).
- DRUID-R6-5 fixed per integrator, feral.md:93. Druid-only effect and Leatherworking kept; "small upgrade while leveling; at 60 most helms with real stats beat it" replaced with "only 10 Spirit beside the bonus, so wearing it trades a helm's stats for the energy" (items/druid.json 8345). The same ranking ("a small bonus while leveling, outgrown on stats by 60") removed from feral-bear.md:94 the same way, ending "for the rage".
- Decision 1 superseded by the round 6 stat rule (prose no longer follows any row's order). Decision 2 applied as above. Decision 3 applied. Decision 4: frontmatter statPriority untouched.
- Other ranking spots checked: balance.md:50 "crit is worth more to you than to most casters" compares classes on a stable mechanic (Vengeance), outside Stat priority; left. No enchant or "your raid set wears" ranking on any druid page.
- No APL notes changed. No new SIM findings; S11 and S12 carried.
