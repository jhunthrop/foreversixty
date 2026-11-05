# Guide accuracy round 22: Druid

Audited on the round 22 audit tree (`scratchpad/gr22/audit`, head 987d08a7c: main after guide round 21 landed; no sim round 21 commit in the log, so there is no sim round 21 druid report). Engine pin d5c708834 (`sim/enginever/version.go`), source read from `scratchpad/gr22/engine`. Active client 1.60.1.70790.

Files: `web/src/content/guides/druid/{index,balance,feral,feral-bear,restoration}.md`, every row note in `data/curated/apl/druid-{balance,feral,feral-bear,restoration}.json`, and every rendered band of `addon-data.json` `rotations/druid-*` (balance 22 bands, feral 24, bear 20, restoration 28), against both shipped orders in `addon-builds.json` and `trainables/druid.json`.

Read first: the round 1, 2 and 22 briefs, `round21/druid.md`, the ledger's Round 21 section.

## Changes since round 21's tree (e039e1490)

None for druid. `git diff e039e1490 HEAD` touches no druid page or APL file, and the four `rotations/druid-*` arrays in `addon-data.json` match the round 21 tree exactly. `addon-builds.json` is unchanged. The shipped orders in `round21/druid.md` still apply. They were re-decoded from the frontmatter strings this round.

One correction to the round 21 report, which does not affect any page: the Feral leveling totals are 9 Balance, 31 Feral and 11 Restoration, not 6/39/6. The raid build is 9/32/10. The page's "Your 31 Feral points are done at 40" is right, and so is the 4-point respec at 60 (Brutal Impact 2, Feral Charge 1 and Natural Shapeshifter 1 go into Rend and Tear +2 and Feral Instinct +2).

## Findings

### DRUID-R22-1: feral.md:119 names Feral Charge's leap without its talent level (CONFIRMED)

- **Line:** "Feral Charge's Cat Form leap lands you behind a mob for a Shred, and in Bear Form it is the charge you pull dungeon packs with."
- **Problem:** This sits in "Levels 20 to 60", a paragraph that gives the level for everything else it names (Cat Form 20, Rake 24, Shredding Attacks' 42 energy, Ravage 32, Pounce 36, Ferocious Bite 24). The talent comes at 25 on the leveling order (index 15, `[1,9,1]`). The line reads as usable from 20.
- **Shape:** round 22's shape 1, a spell named in the 20 to 60 text without its level where the line tells the player to use it.
- **Fix:** "Feral Charge, your talent at 25, gives a Cat Form leap that lands you behind a mob for a Shred, and in Bear Form it is the charge you pull dungeon packs with."

### DRUID-R22-2: balance.md:135 gives Insect Swarm's rank across the 20s without saying it is the talent at 27 (CONFIRMED)

- **Line:** "Mana runs out before mobs do in your 20s and 30s, so ... Insect Swarm stays at rank 1 until 39."
- **Problem:** Insect Swarm is the talent at 27 on both Balance orders. The rank 1 trainer row is level 20, but you learn it through the talent. The sentence tells a player in their 20s which rank to cast, so from 20 to 26 it points at a spell they do not have.
- **Shape:** the same shape as R22-1. The rendered row from 20 is the known whole-row component issue (round 21 D3) and is not re-filed; this finding is the prose only.
- **Fix:** "Insect Swarm, your talent at 27, stays at rank 1 through 38." This also takes the carried "until 39" Should: band 39 casts rank 2 (24974).

### DRUID-R22-3: restoration.md:91 says "Healing power is every heal", but Bloom takes none of it (CONFIRMED)

- **Line:** "Healing power is every heal, and your HoTs take a big share of it per tick."
- **Problem:** Bloom heals for 100% of your Spirit (talent 1328621). Its heal spell 1328622 has `sp_coefficient 0.0` in `spellconst/druid.json`. The engine's 0.2143 in `constants_auto_gen.go` is the generator's convention fill, and `bloom.go` adds no healing power either.
- The page itself calls Bloom "the best single point you take", with "HoTs doing most of your healing" (:55). So the heal that healing power misses is a large one, and "every heal" is false.
- **Shape:** round 22's "every heal" shape.
- **Fix:** "Healing power is every heal but Bloom, and your HoTs take a big share of it per tick."

Counts: CONFIRMED 3, LIKELY 0, NEEDS-SIM 0, NEEDS-OWNER 0. All three are low severity (an omitted level and a one-heal exception). None is a wrong number.

## Round 22 shapes, checked on every page and note

- **(1) A spell named in Levels 20 to 60 without its level, where the line has you use it early.**
  - All four 20 to 60 paragraphs were read against `trainables/druid.json` and the orders. Feral Charge and Insect Swarm are R22-1 and R22-2.
  - Everything else is levelled or already yours. Moonkin Form 40 and Starfire 20, which the section starts at. Claw, Prowl and Rip 20; Rake and Ferocious Bite 24; Ravage 32; Pounce 36. Maul, Swipe and Demoralizing Roar before 20; Frenzied Regeneration 36; Lacerate 25; Primal Bite and Dire Bear Form 40. The Rejuvenation, Regrowth and Healing Touch ranks are all levelled.
  - The 1 to 19 paragraphs name a level for each spell they introduce.
- **(2) Later clauses of a note comparing to a spell on bands before it is trained.**
  - Each comparison is scoped. Balance [6] has Starfire "once ... at 20" and Innervate "from 40"; Balance [2] has Starfire at 20.
  - Feral [8] and [9] compare to Claw and Shred, both live from 20. Feral [7]'s "keep Shredding" is on bands 24-60, where Shred renders.
  - Bear [7] has Lacerate "from 25" and Primal Bite "from 40"; bear [4] has Rend and Tear "from 41".
  - Restoration [7] and [8] have Bloom "from 20". See the sweep.
- **(3) A "cheap" list that includes the top rank.** None. Restoration :81 puts rank 7 (not the top) on the party "for far less mana". Balance :135-137 lists only ranks below the newest, and every one matches the rendered rows.
- **(4) A rounded cost quoted as exact.** Every quoted cost is the exact client value less the build's discounts:
  - Shred 42 (60 - 18), Claw 40 (45 - 5), Rake 35 (40 - 5), Shred 54/48 at 20/21.
  - Primal Bite 15 rage (20 - 5), Lacerate 3 cheaper, Maul 10 by 14 (15 - 5).
  - Shifting Power 55% of base mana, with no Natural Shapeshifter in the raid build. Swiftmend 20% of base mana; Moonglow covers damaging spells only and Tranquil Spirit covers Healing Touch and Tranquility only.
  - Moonglow 25%, Improved Wrath half. Eclipse 0.17 s is the client's own rank 1 text. Demoralizing Roar's 204 is 193 + 1.4 x 8 at 60.
- **(5) "Most for your mana" true only among one spell's ranks.**
  - Restoration :81 says "rank 1 on the party is the most healing your mana buys". I checked it against every spell, not only Rejuvenation ranks. Rank 1 Rejuvenation is 32 over 12 s for 25 mana. With Nature's Splendor (+3 s), Improved Rejuvenation, Genesis and Gift of Nature it heals about 56 (2.2 a mana), and Bloom adds your full Spirit when it runs out, so the total is several times anything else.
  - Healing Touch rank 1 (44 for 25, 22.5 after Tranquil Spirit) is about 2.15 a mana before Bloom. Regrowth rank 1 is about 3.4 a mana without Bloom, but at 70 mana its Bloom share is far lower, so rank 1 Rejuvenation leads once Bloom counts, and the build has Bloom from 20.
  - Tranquility rank 1 is party-wide on a 2-minute cooldown, a different use. The claim holds.
  - Balance :135 and Restoration :85 are scoped to one spell's ranks ("the newest rank is not always the best", "a lower rank often heals more for your mana than the newest one"). OK.
- **(6) "Every heal" claims a talent covers only some of.**
  - R22-3 is the stat-prose case.
  - Talents checked against `talents/druid.json`. Genesis (periodic healing) gives "10% more on every HoT tick". Gift of Nature carries its Tranquility exception. Naturalist covers Healing Touch and Wild Growth only, and the page says so. Tranquil Spirit covers Healing Touch and Tranquility, as stated. Nature's Grace covers "non-periodic spell criticals" (balance :67, a moonkin's crits are spells). Blood Frenzy gives a combo point "on every builder crit", which matches its text. Moonglow's "every damage spell" matches "damaging spells". OK.
- **(7) "Your biggest hit" that another button beats at 60 with raid buffs.**
  - The only druid superlative of this kind is bear [3] and feral-bear :76, "your biggest threat hit". It is routed as an Engine question in the ledger's Round 21 section (which threat model Forever uses), so it is not re-filed. A new web search found no Forever threat numbers (see Sources).
  - Feral :71 "Ravage hits much harder" (than Pounce) holds. Balance [1] "Moonfire's hit is the best instant you have" holds: Insect Swarm has no direct hit and Faerie Fire does no damage.
- **Full sweep of every superlative, "only", "every", "never" and quoted value** (grep over all five pages plus every note): every other hit was re-checked against the client rows and the decisions from rounds 1 to 21 (Shifting Power "the only energy button", Dodge "your only avoidance", Clearcasting "not consumed by Wrath", Barkskin, Enrage, Frenzied Regeneration, Tranquility 5 to 2 minutes, Wild Growth 2.5 s on 6 s, Starfire 3 s). No new error.

## Round 21 decisions re-checked

- D1 (biggest threat hit) is still routed. D2 (routed sims) and D3 (rows that render before their talent: Shifting Power 10-39, Insect Swarm 20-26, Nature's Swiftness 10-29, Swiftmend 10-42, the Feral "Below level 20" rows past 20) stand.
- No round 21 druid fix exists to re-check, because round 21 found nothing.

## Shoulds (wording, not findings)

- **New: balance.md:105** "Spell power is damage on every cast". Faerie Fire (balance :97) is a cast with no damage. Suggested: "Spell power is damage on every damage spell."
- **New: balance.md:91** "it always lands on a Starfire or a DoT". A Faerie Fire cast while Clearcasting is up would take it. Suggested: "so it lands on your next Starfire, DoT or other paid spell."
- **Carried from round 21 (none applied):**
  - feral-bear :76 and bear [3] "biggest threat hit" (routed).
  - restoration :121 should add the tauren 21 to 23.
  - Balance [6] "With Improved Wrath" at band 10.
  - restoration :85 "Below 20, from 21".
  - feral-bear :70: Thorns before Dire Bear Form.
  - feral :67 "Both builds skip Furor".
  - Feral Bite "With Predatory Discipline" and Balance Moonfire "With Improved Moonfire and Nature's Splendor", both conditional.
  - Bear Demoralizing Roar "204 ... at 60" on rank 1 bands.
  - Feral and bear Faerie Fire "free" on band 18.
  - balance :135 "until 39" (folded into R22-2's fix).
  - "for good" at balance :115 and restoration :101.
  - restoration :57 "next Healing Touch" should be "next Nature spell".
  - The lines carried from round 19.

## Decisions

- **D1. R22-1 and R22-2 are CONFIRMED, not Shoulds.** Each sentence is literally true from the level the talent arrives. But each sits in the 20 to 60 text, names a talent button with no level, and tells the player how to use it (pull with it, cast this rank). That is the shape the other round 22 auditors filed as findings (Lacerate at 30, Death Coil at 42, Repentance at 31), and the round is judged across classes on one standard.
- **D2. R22-3 is CONFIRMED.** The client gives Bloom's heal 1328622 a zero spell-power coefficient. The amount comes from the talent aura's Spirit percentage (1328621, effect amount 100), and the engine models the same. A stat line that says healing power is "every heal" is false for a heal the page itself ranks as the build's best point. Bloom does crit (`bloom.go`, no CANT_CRIT), so "Crit makes your direct heals land bigger" stands.
- **D3. Primal Bite against Maul threat stays routed** (ledger Round 21 Engine questions). No client value or Forever source settles it.

## Sims requested

- None new. Carried and already routed: Feral Claw against Shred at 21 to 23; Balance leveling Eclipse Wrath on and off at 36, 40 and 50.

## Sim findings

- None new.

## Engine questions

- **Bloom coefficient.** `constants_auto_gen.go` lists `BloomSpellCoeff 0.2143`, a convention fill: the client's 1328622 coefficient is 0.0. `bloom.go` correctly ignores it. The generator could emit 0 for triggered heals whose client column is zero, so no later ability file picks up the 0.2143 by mistake.
- Carried: the bear threat model (Maul 1.75x against Primal Bite 2.0x, both unconfirmed); Rend and Tear on any bleeding target.

## Sources

- **Client 1.60.1.70790** (`data/builds/1.60.1.70790/`):
  - `talents/druid.json` (every rank text; Eclipse, Genesis, Gift of Nature, Furor, Natural Shapeshifter ranks; the Shifting Power prerequisite)
  - `trainables/druid.json` (every trainer level and base cost)
  - `spelltips/druid.json` (all 218 druid rows: costs, cast times, cooldowns, tooltips)
  - `spellconst/druid.json` (1328621 and 1328622 Bloom, 774 and 9841 Rejuvenation)
  - `addon-builds.json`, `addon-data.json` (`rotations/druid-*`, every band, race keys)
- **Curated:** `data/curated/apl/druid-{balance,feral,feral-bear,restoration}.json`.
- **Engine d5c708834:** `sim/druid/bloom.go`, `sim/druid/constants_auto_gen.go`.
- **Reports:** `round21/druid.md`, `ledger.md` Round 21, `round2/druid.md`, `round9`, `round13`, `round16` to `round18` druid (rank 11 book ranks).
- **Web (Forever first; none settles threat numbers):**
  - Kami Labs, "WoW Forever Druid: Every Reworked Spell and Talent": https://kami-labs.fr/en/wow-forever/druide-wow-forever-sorts-et-talents-modifies/
  - RankedBoost, "WoW Forever Druid Feral Combat Macros": https://rankedboost.com/world-of-warcraft/forever/macros/druid/feral-combat/
  - warcraft.wiki.gg, "Primal Bite": https://warcraft.wiki.gg/wiki/Primal_Bite

## Note sweep

One line per card row note: spec, row and action, the bands it renders on, each named talent, spell, rank or value with the level the player has it (leveling / raid order where they differ), and the verdict.

- **Balance [1] Moonfire, end of fight. Bands 10-60, ranks 2 to 10.** Wrath cast time; "best instant": Insect Swarm has no hit (from 27), Faerie Fire no damage, Hurricane is channelled (40). OK.
- **Balance [2] Wrath, end of fight. Bands 10-60.** Starfire at 20, scoped; before 20 "the last few seconds". OK.
- **Balance [3] Moonfire upkeep. Bands 10-60.** Improved Moonfire 20-21 and Nature's Splendor 26 on both orders, as a conditional. OK (Should carried).
- **Balance [4] Insect Swarm. Bands 20-60.** Talent at 27 on both orders (whole-row issue, D3 r21); one global a refresh. OK.
- **Balance [5] Eclipse Wrath. Bands 10-60.** Eclipse at 35 / 39; 0.17 s at one point (client rank 1 text); half a second at three (raid 42). OK.
- **Balance [6] Wrath while behind on mana. Bands 10-60.** Innervate 40; Improved Wrath 10-14; Starfire 20; Wrath 60 against Starfire 236 mana at 60 after Moonglow. OK (band 10 Should carried).
- **Balance [7] Starfire. Bands 20-60, ranks 1 to 6.** Eclipse at 35 / 39. OK.
- **Feral [1] Ravage. Bands 32-60.** Prowl 20, Poised 32 (10 off 60). OK.
- **Feral [2] Faerie Fire. Bands 18-60.** Cat Form 20; free on a 6 s cooldown. OK (band 18-19 Should carried).
- **Feral [3] Moonfire, below 20. Bands 10-60.** Cat Form 20. OK (component row issue).
- **Feral [4] Wrath, below 20. Bands 10-60.** As [3]. OK.
- **Feral [5] Shifting Power. Bands 10-60.** Talent at 40 on both orders (whole-row issue); 40 energy, mana, 8 s. OK.
- **Feral [6] Rip. Bands 20-60, ranks 1 to 6.** 5 combo points, under 2 s; Rend and Tear at 37 / 36. OK.
- **Feral [7] Ferocious Bite. Bands 24-60.** Shred live from 20; Predatory Discipline at 31 / 30, as a conditional. OK (Should carried).
- **Feral [8] Shred. Bands 20-60, ranks 1 to 5.** Shredding Attacks 20-22 on both orders; 54, 48 and 42 against Claw's 40; "harder at every rank" holds on every pairing the bands cast. OK; the gate is routed.
- **Feral [9] Claw. Bands 20-60, ranks 1 to 5.** 40 against Shred's 54 at 20 and 48 at 21; "more for its energy" at 21 is in the routed sim. OK.
- **Bear [0] Cooldowns. Page only.** Mighty Rage Potion in form. OK.
- **Bear [1] Barkskin. Bands 44-60.** Trainer 44; 20%, 15 s, 60 s (22812). OK.
- **Bear [2] Frenzied Regeneration. Bands 36-60.** Trainer 36; half health, 20 rage. OK.
- **Bear [3] Primal Bite. Bands 40-60, ranks 1 to 3.** Talent 40; 6 s cooldown; "biggest threat hit" is routed (D3). Not re-filed.
- **Bear [4] Lacerate. Bands 25-60, ranks 1 to 5.** Talent 25; five stacks; 3 s; Rend and Tear from 41 (index 31). OK.
- **Bear [5] Faerie Fire. Bands 18-60.** Bear Form 10; free on a 6 s cooldown; carries threat. OK.
- **Bear [6] Demoralizing Roar. Bands 10-60, ranks 1 to 5.** 204 at 60 (193 + 1.4 x 8); the last 3 s; no stacking with Demoralizing Shout. OK (rank 1 bands Should carried).
- **Bear [7] Maul. Bands 10-60, ranks 1 to 7.** 15 rage, 10 by 14 (Ferocity 10-14); Lacerate 25, Primal Bite 40. OK.
- **Bear [8] Swipe. Bands 16-60.** Trainer 16; above 40 rage; 3 targets ("Swipe 3 nearby enemies"). OK.
- **Restoration [1] Innervate. Bands 40-60.** Trainer 40; under 30%; +400% regeneration. OK.
- **Restoration [2] Read Leyline. Alliance Skyborne, no addon row.** 50% for 15 s; not under Innervate. OK.
- **Restoration [3] Nature's Swiftness. Bands 10-60.** Talent 30 (whole-row issue); tank under 35%. OK.
- **Restoration [4] Healing Touch save. Bands 10-60.** The newest trainer rank at every band (5186 to 9889; rank 11 is a book). OK.
- **Restoration [5] Swiftmend. Bands 10-60.** Talent 43 (whole-row issue); under 75%; keeps the HoT; 15 s; 20% of base mana (cost_pct 20, no discount in the build). OK.
- **Restoration [6] Tranquility. Bands 30-60, ranks 1 to 4.** Trainer 30; two of the party under 70%; every 2 s. OK.
- **Restoration [7] Regrowth, tank. Bands 12-60.** The top rank at every band (8936 to 9858); Bloom at 20 (index 10). OK.
- **Restoration [8] Rejuvenation, tank. Bands 10-60.** Rank 10 at 60; Bloom 20. Leveling ranks match the rows: 2 at 20-39 (night elf rank 3 at 29, the default 2091 at 29 overridden for all four druid races), 4 at 40-43, 5 at 44-48, 6 at 49-56, 8 at 57-59. OK.
- **Restoration [9] Healing Touch filler. Bands 10-60.** Rank 2 at 20 (tauren 21-23), 3 at 30-33, 4 at 34-39, 5 at 40-45, 6 at 46-49, the top rank from 50. Matches the rows. OK.
- **Restoration [10] Rejuvenation, party. Bands 10-60.** Under 70%; no mana-pace gate. OK. Rows [11] to [14] carry no note.
