# Guide accuracy round 24: Paladin

Audited on the round 24 audit tree (ea9adea1c: main after guide round 23, which merged sim round 21 at d5cc4d6b4). Engine pin 170a59c39 (`sim/enginever/version.go`; source read at `scratchpad/gr24/engine`). Active client 1.60.1.70790 under `data/builds/1.60.1.70790/`.

Files audited:
- `web/src/content/guides/paladin/{index,holy,protection,retribution}.md`, line by line.
- Every card note in `data/curated/apl/paladin-{holy,protection,retribution}.json`, with each row's condition read, not only its note.
- Every rendered row of `addon-data.json` (`rotations/paladin-retribution` 19 bands including the new 53, `paladin-protection` 18, `paladin-holy` 16), with every `race_spell_ids` split.
- Both shipped point orders per spec in `addon-builds.json` (decoded by tier and column against `talents/paladin.json`).
- `talents/paladin.json`, `spelltips/paladin.json`, `spellranks.json`, `data/curated/rank-choices.json`, engine `sim/core/spell.go` (level-gap rule), `spell_max_level_auto_gen.go`, `sim/paladin/consecration.go`, `sim/api/envelope.go` (default fight 180 s, 20% variation).

Read first: the round 23 paladin report, the ledger's Round 19 to 23 sections, `design/reviews/2026-11-05-sim-audit-round21/paladin.md` and its integration notes.

Order of work, per the brief: sim round 21's paladin changes first, then round 23's paladin fixes, then the full sweep. Lenses: the round 20 note sweep (every row note, every band, both orders); every superlative, comparison and quoted cost, duration and cooldown on both sides across all bands; spells and ranks named before the player has them; "every", "only", "on cooldown" and "every time" claims against the card's own conditions and against every race's rows; the reason each line gives, not only its fact; class index pages against the cards.

## Changes since round 23's tree (24ee457d0)

- **Sim round 21** (963c5ddcc, f84045998, 7716885f4, 4ffeaac1b, 6d5983239):
  - Retribution card: the twist rows 6, 7 and 8 dropped the 1200 attack power gate and the flat 30% mana gate; all three now need mana percent at least 0.8 of the fight's remaining percent. Notes on rows 2, 6, 8, 9 and 11 rewritten. Consecration at 53 split by race (`rank-choices.json`: Human rank 1, Dwarf and Undead default to rank 4; new band 53 in addon-data).
  - retribution.md: Strengths and Weaknesses bullets, the opener (two paragraphs now), priority step 4, "The Seal twist" paragraph (was "The raid twist"), the Consecration paragraph, the Levels 20 to 60 Consecration line, and a new common mistake.
- **Guide round 23** (2b5731536, 9577f4bd3): Holy row 3's last sentence ("From 25, Holy Light is the only heal you can keep casting that outheals a tank by itself"), holy.md:33 and :59 ("as long as the rank is close to your level").
- Protection page and card, and the Holy rendered rows, are unchanged (compared in code against `git show 24ee457d0:...`).

## Findings

| id | status | kind | where | quote | evidence | fix |
|---|---|---|---|---|---|---|
| PALADIN-R24-1 | CONFIRMED | page prose (a reason sim round 21 disproved, left behind by its change) | retribution.md:70 | "**Seal of Command** is a required point on the way to Vengeance. Below 60 you never cast it: Seal of Righteousness does more with leveling gear." | The card never casts Seal of Command because of Unsealed Wrath, not gear: rows 2, 7 and 8 open with `auraIsKnown(Unsealed Wrath)`, and the leveling build has no Unsealed Wrath (decoded order: Repentance at 31, Champion of the Light 35 to 39, Divine Strength from 40, no Unsealed Wrath). Sim round 21 T1 retired exactly this reason: "the old note's reason (with little attack power Seal of Command does less than Righteousness) does not hold on this build at any buff state measured", and at bare 60, fresh gear with no raid buffs, twisting with Seal of Command as the main Seal beat staying in Righteousness by +1.7% Human and +0.5% Undead. The gear clause tells a reader that better gear is what makes Seal of Command worth casting, which is the opposite of the measurement, and it contradicts the rewritten opener and twist paragraph on the same page (:78, :93), which tie Seal of Command to Unsealed Wrath alone. Second, smaller problem: the shipped raid order takes Unsealed Wrath at 40, so for a player who follows it, the card casts Seal of Command from 40 and "below 60 you never cast it" is false for them. | "**Seal of Command** is a required point on the way to Vengeance. The leveling build never casts it: without Unsealed Wrath, Seal of Righteousness does more on its own. With Unsealed Wrath, Seal of Command becomes your main Seal for the twist below." |
| PALADIN-R24-2 | CONFIRMED | page prose ("only" claim false on the same page, and out of step with the card row it describes) | holy.md:76 | "**Divine Favor**, then **Holy Light** (rank 8), when the tank drops below 40%. It is the only heal that outheals a tank by itself, and the crit is guaranteed." | Lay on Hands, on the same page at :91, "heals your target for as much as your own maximum health" (client 633/2800/10310: "Heals a friendly target for an amount equal to the Paladin's maximum health"). At 60 that is several thousand points, far past a rank 8 Holy Light. So "the only heal that outheals a tank" is false as written. Round 23's review recognised this and qualified the card row this line describes: Holy row 3 now reads "From 25, Holy Light is the only heal you can keep casting that outheals a tank by itself." The prose still carries the unqualified version. | "...when the tank drops below 40%. It is the only heal you can keep casting that outheals a tank by itself, and the crit is guaranteed." |
| PALADIN-R24-3 | LIKELY | page prose ("every time" against the card's own gates) | retribution.md:97 and :135 | :97 "Rank 1 is cheap enough to put down every time it is ready, so on one target it does more than the higher ranks for most of your leveling." :135 "a cheap rank you can put down every time it is ready beats a big one you have to hold back" | Both sit in the paragraph that teaches the pace (":97 Fall under that and Consecration waits until you catch up"), and the card's row 11 gates every rank, rank 1 included, on the mana reserve and the four-fifths pace. Sim round 21 C1 measured rank 1 at 50 to 52 on the paced line: about 19 casts a fight and 11 to 17 s out of mana. With the default 180 s fight (20% variation, `sim/api/envelope.go`) and Consecration's 8 s cooldown (`consecration.go`), about 22 casts are available, so rank 1 goes down roughly five times in six, not every time. The mechanism (rank 1 is cast far more often, 19 against rank 4's 9) is right; "every time" is what the card's own gate breaks. LIKELY rather than CONFIRMED because the available-cast count depends on fight length and first-cast timing; see Sims requested 1. | :97 "Rank 1 is cheap enough to put down almost every time it is ready, so on one target it does more than the higher ranks for most of your leveling." :135 "a cheap rank you can put down nearly every time it is ready beats a big one you have to hold back" |

Counts: CONFIRMED 2, LIKELY 1, NEEDS-SIM 0 (3 optional measurements requested), NEEDS-OWNER 0.

## SHOULDs (wording, not findings)

- **index.md:29** "Retribution's capstone, Unsealed Wrath, brings a new kind of swap back for raiders." Since sim round 21 the card twists in a dungeon or a raid, with or without raid buffs, whenever the mana keeps pace. "For raiders" stays defensible (Unsealed Wrath is only in the raid build, which retribution.md:68 calls "the raid capstone"), but "brings a new kind of swap back with the raid build" would not echo the retired raid-buff gate.
- **Retribution row 11, first sentence** (rendered on every band 20 to 60): "covers a Seal, Judgement and Holy Strike (rank 4 at 60 without raid buffs)". The parenthetical sits after "Holy Strike" and can read as Holy Strike's rank (Holy Strike is rank 8 at 60). The row icon makes Consecration clear, and the second sentence states the reserve correctly ("The mana left after it must still pay for those three"), so not false. "Last, rank 4 at 60 without raid buffs: while your mana percent..." would read cleanly.
- **Retribution row 2 and retribution.md:78** "in a dungeon or a raid": the row has no location condition, so a raid-build paladin out in the world does the same. Harmless scoping.
- **retribution.md:135** "move up to rank 4 at 53 as a Dwarf or Undead, from about 55 as a Human, when rank 1 has lost nearly all of its spell damage bonus": the "when" clause can read as covering 53 too, where rank 1 still keeps 20% (rank 1 MaxLevel 25, `1 - (53 - 25 - 12) x 0.05`). :97 attaches it to the Human only; match that structure.
- **retribution.md:97 and :135, Human "from about 55"**: the rows give every race rank 4 from band 54 (only band 53 carries the Human rank 1 row). Round 19 measured Human level at 54 and 55 and ahead from 56, so "about 55" is within the measurement; "from 54 or 55" would match the row exactly.
- Carried, still true as SHOULDs:
  - retribution.md:43 "turns Seal swapping into real damage" (the measured gains are a few percent).
  - Retribution row 6 "180 to 190 at 60" (exact: 180 and 189).
  - retribution.md:72 Holy Conduit omits Holy Wrath; :123 "both come back every 10 seconds"; :36 "Judge every 8 seconds"; :133 "refunds most of each Seal" (round 23 D3).
  - Holy row 12 "rank 1 for most of your leveling" and "a small rank" at 50 to 54; holy.md:81 "half as much".
  - protection.md:111 block chance; Holy Shield ranks 2 and 3 tooltips missing "4 charges" (data lane).

## Checked and right this round

**Sim round 21's Retribution card (new work, every row, every band, both orders)**

- **Twist conditions against the notes.** Rows 6, 7 and 8 each carry `currentManaPercent >= remainingTimePercent x 0.8` and no attack power term; row 9 carries none. Row 6's "while your mana percent is at least four fifths of the fight left", row 8's "while your mana keeps pace with the fight" and row 9's "any time your mana falls behind the fight" describe them exactly. Row 9 also fires for a twisting paladin only once the current Seal is in its last second, which is what "keep ... up ... recast it in its last second" says.
- **Row 7's timer rule.** `auraRemainingTime(Unsealed Wrath) + 15 s > currentSealRemainingTime` is false right after Righteousness goes up (its 30 s against a stale timer) and true the moment the swing lands and Unsealed Wrath resets to 15 s. "Unsealed Wrath restarts at 15 seconds when the swing lands, so its timer tells you" holds; the talent text (1310735) gives 15 s and three stacks.
- **Costs.** Seal of Righteousness rank 8 200 and Seal of Command rank 5 210 (spelltips 20293, 20920), Benediction 10% at 14: 180 and 189. Seal of Command ranks 65/110/140/180/210 at 20/30/40/50/60. No note quotes a twist cost below 60.
- **Row 11 Consecration.** Condition: (mana minus Consecration's cost) at least Seal plus Judgement plus Holy Strike, and the four-fifths pace. The second sentence states the reserve; "halfway through ... 40% mana or more" is 0.8 x 50. "(rank 4 at 60 without raid buffs)": `rank-choices.json` paladin-retribution band 60 bare = rank 4, raid row casts rank 5 (20924). Both races, no race split at 60.
- **Consecration at 53, every race.** Band 53 renders rank 4 (20923) with `race_spell_ids {human: rank 1}`; Dwarf and Undead take rank 4. retribution.md:97 ("a Dwarf or Undead paladin moves up at 53") and :135 ("at 53 as a Dwarf or Undead") match. Bands 50 and 52 render rank 1 for every race ("keep it to 52", "go back to rank 1 at 50"). Band 54 renders rank 4 for every race (SHOULD above on "about 55").
- **Spell damage share behind the rank advice.** Consecration rank 1 is learned at 20, MaxLevel 25 (no client exception), so it keeps all of its coefficient to 37, 50% at 47 ("about half ... mid to late 40s"), 10% at 55 ("nearly all" lost for the Human), none from 57.
- **Opener, both builds.** Pre-pull rows 1 and 2 split on `auraIsKnown(Unsealed Wrath)`; row 2 has no pace gate, which matches retribution.md:78 (Seal of Command straight away, mana is full at the pull). Seal of the Crusader at -1.5 s matches "about a second and a half before the pull".
- **The Seal twist paragraph (retribution.md:93).** "80% at the pull, 40% halfway and 20% with a quarter left" is 0.8 x 100, 50 and 25. "The same pace as Consecration": row 11 uses the same term. "Each with its normal chance to fire" is the talent's "giving the stored Seal its normal chance to activate once per stack". "Let three swings land under Seal of Command": each swing stores one copy, and the release swing under Righteousness stores one Righteousness copy, so the next Command swing releases it and the cycle is one Righteousness swing and three Command swings.
- **Priority step 4 and the new common mistake (:87, :140)** match the card's order (twist rows 6 to 8 sit after Hammer of Wrath and before Holy Strike) and its gate.
- **Strengths and Weaknesses (:43, :47).** True as stated; "real damage" stays a SHOULD.
- **Raid order side effects.** Unsealed Wrath at 40 on the raid order, after Repentance at 31 (its prerequisite). Every note that names Unsealed Wrath is conditional on having it, so no note credits it early on either order. Only retribution.md:70 states a level ("below 60"), filed as R24-1.

**Round 23's fixes**

- Holy row 3 "From 25 ... you can keep casting": at 25 to 29 rank 2 with the Oath's share heals 233 to 218 against Flash of Light rank 1's 94 to 102 (round 23's worked numbers, unchanged rows); from 30 Holy Light is the largest castable heal on every band, including rank 4 at 40 to 42 and 47 to 49. Lay on Hands and Light's Vigil are what "you can keep casting" excludes. The prose counterpart is R24-2.
- holy.md:33 and :59 now carry the rank limit, in line with :78, :95, :127 and Holy rows 1 and 11.

**Full sweep**

- **Class index (index.md).** "Put a Seal up once every 30 seconds and Judge on cooldown" holds for Protection (row 5: Seal of Fury, last 2 s, 30 s) and for Retribution without Unsealed Wrath; the next sentence names the Unsealed Wrath swap (SHOULD on "for raiders"). Holy Strike at 6 with 10 s; Consecration and Blessing of Kings at 20; Hammer of the Righteous at 40; Seal of Fury and Righteous Fury at 10; one-point talents at 11, 16, 21, 31; races Human, Dwarf, Undead (`combos.json`).
- **Point orders against prose.** Retribution 14/7/30 and raid 12/8/31, "moves Reverence's two points into the third rank of Precision and Unsealed Wrath"; Levels 1 to 19 (Benediction 10 to 14, Improved Judgement 15 and 16, Conviction 17 to 19); Champion of the Light 5/5 at 39 ("by level 40"); Repentance 31. Holy 31/12/8, Spiritual Focus 10 to 14, Healing Light 15 to 17, Reverence 18 and 19, Oath of Light 25, Holy Shock 32. Protection 10/33/8, Divine Intellect 10, Redoubt 11 to 15, Precision 16 to 18, Sacred Duty 19 and 21, Templar's Bulwark 20, Swift Judgement 29, Holy Shield 30, Shieldsworn 40. All match.
- **"Only", "every", "on cooldown" and "whenever ready" claims against conditions.** Protection: Holy Shield and Holy Strike rows have no condition ("on cooldown"); Hammer of the Righteous, Judgement, Exorcism carry the Holy Shield reserve their notes and :95 to :98 state; Swift Judgement has none ("whenever it is ready"); Consecration above half mana plus reserve (:99 "only above half mana"). Holy: rows 4 to 10 carry the five-point pace (:77); Divine Favor and row 3 at 40% (:76). Retribution: Judgement `spellCanCast` (on cooldown), Holy Strike and Exorcism on cooldown. Only R24-3 is broken by a gate.
- **Race rows.** Retribution band 53 is the only `race_spell_ids` split for any paladin card; Protection and Holy have none. Race prose (retribution.md:113, protection.md:119, holy.md:103) matches `races.json` (Ruthless, The Human Spirit, Will to Survive, Stoneform 10% for 8 s, Artisan's Precision, Big Game Hunter, Will of the Forsaken without immunity, Cannibalize both bars, Touch of the Grave once every 15 s). No measured race result contradicts the lines; see Sims requested 2 for the post-T1 re-check.
- **Quoted values, unchanged since round 23 and re-read:** Lay on Hands 20 minutes, 550 mana at rank 3, drains your mana without stopping regeneration (10310); Righteous Fury 30% base mana, 30 minutes, 33% plus 90% Holy (25780); Consecration costs 135/235/320/435/565 and 8 s; Holy Strike 18 (Retribution) and 20 (Protection) at 60; Light's Vigil 1,340 against ten rank 6 Flash of Lights at 140; Hammer of Wrath 6 s; Divine Shield 10 and 12 s.
- **Sources** in frontmatter: unchanged since round 23's check (all nine distinct URLs 200 on 2026-11-04).
- **Voice:** no em or en dashes, no site or sim talk, no nightly numbers, no stale 70755 reference in the four pages.

## Decisions

- **D1. R24-1 is CONFIRMED, not a SHOULD.** The sentence gives a reason (gear) that sim round 21 measured and retired for this exact choice, and the page's own rewritten opener and twist paragraph now teach the real rule (Unsealed Wrath). A reason contradicted by the measurement is the brief's "reasons that do not match the numbers" shape, and the line is a leftover of sim round 21's change, which round 20's shape 3 makes a finding.
- **D2. R24-2 is CONFIRMED.** Lay on Hands on the same page heals for far more than a Holy Light. Round 23's review already qualified the card row; the prose that describes the same row did not follow.
- **D3. R24-3 is LIKELY.** C1's 19 casts against about 22 available is measured, but the available count is my estimate from the default fight and the cooldown. Either way the fix ("almost every time") is safe, since the mechanism stays right.
- **D4. Not re-filed, per the brief:** whole rows rendering before their talent (Seal of Command rows at 20 to 25, the Unsealed Wrath release row at 10 to 19, Swift Judgement at 10 to 28, Divine Favor at 20, Holy Shock at 30 and 31); card actions casting measured ranks; the Holy Light rank 9 book; Exorcism never crits; Benediction not discounting Hammer of Wrath; BiS race captions; the Protection APL's non-rendered top-level "+60 percent threat" note (round 23 housekeeping).

## Sims requested

1. Retribution Consecration rank 1 on the paced line at 30, 40 and 50 (bare band, both factions, seeds 1 and 50001): casts per fight against cooldowns available. Settles R24-3's wording ("almost every time" against "every time"). Optional.
2. Retribution race re-check after sim round 21 T1 (the paced twist): Human against Dwarf at raid 60 and bare 60 on the raid build. retribution.md:113 says "Dwarf matches Human on the raid build in a raid, and Human is ahead solo"; T1 moved the raid build +0.3% (raid) and +0.5 to +1.7% (bare) and changed attack power's value, and no Dwarf number was taken. Optional; race re-runs on the new builds are already routed.
3. Retribution raid order at 40 to 58: twist (Unsealed Wrath from 40) against Seal of Righteousness only, bare band. Would decide whether R24-1's fix should also tell a raid-order leveler to twist before 60. Optional; the fix as written is true either way.

## Engine questions

- None. The twist conditions, Unsealed Wrath (`unsealed_wrath.go`, 15 s, three stacks), Consecration's 8 s cooldown and the level-gap rule (`sim/core/spell.go:777-850`, grace 12, 5% a level) agree with the page's lines.

## Sim findings

- None new.

## Sources

No web research was needed this round: the client data, the engine and sim round 21's measurements decided every call.

- **Client data 1.60.1.70790** (`data/builds/1.60.1.70790/`): `talents/paladin.json` (Unsealed Wrath 1310735, Sanctified Judgement, Holy Conduit, Benediction), `spelltips/paladin.json` (Lay on Hands 633/2800/10310, Seal of Command 20375 to 20920, Consecration 26573 to 20924, Righteous Fury 25780), `spellranks.json`, `addon-builds.json`, `addon-data.json` (`rotations/paladin-*`, every band and race split), `races.json`, `combos.json`.
- **Curated:** `data/curated/apl/paladin-{holy,protection,retribution}.json`, `data/curated/rank-choices.json`.
- **Engine 170a59c39:** `sim/core/spell.go:777-850`, `sim/core/spell_max_level_auto_gen.go`, `sim/paladin/consecration.go`, `sim/api/envelope.go:398`.
- **Reports:** `design/reviews/2026-11-05-sim-audit-round21/paladin.md` (C1, O1, T1), `design/reviews/guide-accuracy/round23/paladin.md`, `ledger.md` Rounds 19 to 23.
- **Git:** 963c5ddcc, f84045998, 7716885f4, 4ffeaac1b, 6d5983239 (sim round 21 paladin), 2b5731536, 9577f4bd3 (round 23 paladin), diffed against 24ee457d0.

## Counts

| status | count |
|---|---|
| CONFIRMED | 2 (R24-1 retribution.md:70, R24-2 holy.md:76) |
| LIKELY | 1 (R24-3 retribution.md:97 and :135) |
| NEEDS-SIM | 0 (3 optional measurements requested) |
| NEEDS-OWNER | 0 |
| Card notes with a finding | 0 of 33 (Retribution 11, Protection 12 plus the auto-cooldown note, Holy 9 plus the auto-cooldown note) |

## Note sweep

One line per row note: spec, row action, bands it renders on, each named talent, spell, rank or value with the level the player has it (leveling / raid order where they differ), verdict.

**Retribution** (bands 10, 12, 18, 20, 26, 28, 30, 34, 36, 38, 40, 42, 44, 50, 52, 53, 54, 58, 60; the raid order differs from 40: Unsealed Wrath 40, Divine Strength 41 to 45, Improved Seals 46 to 50, Holy Conduit 51 and 52, Divine Intellect 53 to 57, Precision 58 to 60, no Reverence)

- **Row 1 Judgement, 10-60:** Judgement 4; off the global cooldown; 6% of base mana; Benediction full at 14 on both orders; never consumes the Seal; Sanctified Judgement 3/3 at 22 on both orders, 60% (1311074). OK.
- **Row 2 Seal of Command (opener, Unsealed Wrath), 20-60:** conditional on Unsealed Wrath (none on leveling, 40 on raid). Seal of the Crusader 6, Judgement 4, Seal of Command talent 26 (render at 20 to 25 is D4). "In a dungeon or a raid" (SHOULD). Three swings, three stacks (1310735). Crusader refresh by melee. OK.
- **Row 3 Seal of Righteousness (opener, no Unsealed Wrath), 10-60:** top rank trained on every band (rank 2 at 10 to rank 8 at 58). Seal of the Crusader 6. OK.
- **Row 4 Exorcism, 20-60:** Exorcism 20; Undead and Demons only. OK.
- **Row 5 Hammer of Wrath, 44-60:** trains 44; 6 s on all three ranks; Champion of the Light 5/5 at 39 on both orders. OK.
- **Row 6 Seal of Righteousness (release), 10-60:** Unsealed Wrath conditional; four-fifths pace matches the condition; "same pace as Consecration" (row 11's term); "180 to 190 at 60" is 180 and 189 (SHOULD). Renders before Seal of Command exists at 10 to 25 (D4). OK.
- **Row 7 Seal of Command (back after one swing), 20-60:** Unsealed Wrath restarts at 15 s; the timer condition reads as the note says. OK.
- **Row 8 Seal of Command (keep up), 20-60:** Unsealed Wrath conditional; pace; last second (`currentSealRemainingTime <= 1`). OK.
- **Row 9 Seal of Righteousness (otherwise), 10-60:** top rank trained; last second; "every fight without Unsealed Wrath, and any time your mana falls behind" matches rows 6 to 8's gates. OK.
- **Row 10 Holy Strike, 10-60:** trains 6; 18 at 60 with Benediction (10333, 20 less 10%); Sacred Arbiter 27 on both orders. OK.
- **Row 11 Consecration, 20-60:** Consecration 20; reserve and pace match the condition; "(rank 4 at 60 without raid buffs)" matches the bare-60 rank row, all races; 40% at halfway. Placement of the parenthetical is a SHOULD. OK.

**Protection** (bands 10, 12, 18, 20, 26, 28, 30, 34, 36, 38, 40, 42, 43, 44, 50, 52, 58, 60; raid order identical; notes and rows byte-identical to round 23, re-read against the conditions)

- **Row 0 auto cooldowns:** Lay on Hands 10, below 10%. OK.
- **Row 1 Templar's Bulwark, 20-60:** talent 20; below 35%; whole health bar, 8 s; Forbearance; Divine Protection 6; "Divine Shield from 34". OK.
- **Row 2 Divine Shield, 34-60:** below 25%; 10 s rank 1 (34 to 44 rows), 12 s rank 2 at 50; half damage dealt. OK.
- **Row 3 Divine Protection, 10-60:** below 20%; Forbearance. OK.
- **Row 4 Stoneform (Dwarf):** below 70%; 10% physical, 8 s, global, 3 minutes. OK.
- **Row 5 Seal of Fury, 10-60:** last 2 s; Holy Shield reserve "once you have Holy Shield at 30"; 30 s Seal, not consumed by Judgement. OK.
- **Row 6 Holy Shield, 30-60:** no condition ("on cooldown"); 10 s cooldown and duration; four charges. OK.
- **Row 7 Holy Strike, 10-60:** no condition; twenty at 60 (no Benediction in this build); Iron Creed 39, 41 and 42; Shieldsworn 40, 6 s. OK.
- **Row 8 Hammer of the Righteous, 40-60:** trains 40; reserve matches "once Holy Shield's mana is safe"; target plus three; Righteous Fury 10. OK.
- **Row 9 Swift Judgement, 10-60:** no condition ("whenever it is ready"); one minute, free, off the global cooldown; taunt under Seal of Fury (render before 29 is D4). OK.
- **Row 10 Judgement, 10-60:** reserve from Holy Shield at 30; 6% base mana; taunt. OK.
- **Row 11 Exorcism, 20-60:** Undead and Demons; reserve. OK.
- **Row 12 Consecration, 20-60:** above half mana plus reserve; Holy Shield from 30; Hammer of the Righteous from 40 far ahead per mana against rank 1 (40 to 58) and rank 5 (60). OK.

**Holy** (bands 10, 14, 20, 25, 30, 38, 40, 43, 44, 46, 47, 50, 54, 55, 56, 60; raid order identical; Benediction 53 to 57 changes no quoted value)

- **Row 0 auto cooldowns:** mana potions and Demonic Runes; Divine Favor kept off. OK.
- **Row 1 Oath of Light, 25-60:** talent 25; Holy Light, Flash of Light, Holy Shock while the rank is close to your level; 5 minutes; condition `remainingTime > 10 s`. OK.
- **Row 2 Divine Favor, 20-60:** talent 21 (render at 20 is D4); below 40%; 2 minutes. OK.
- **Row 3 Holy Light, tank below 40%, 10-60:** Divine Favor from 21; rank 8 at 60 (10329, trains 54); rank 9 book (D4); "From 25 ... you can keep casting": holds at 25 to 60 (round 23's fix, re-checked); 20 to 24 now excluded by "From 25". OK.
- **Row 4 Holy Light, tank below 70%, 10-60:** `currentManaPercent > remainingTimePercent - 0.05`; "no more than 5 points under", 45% at halfway. OK.
- **Row 5 Holy Shock, below 70%, 30-60:** Holy Shock 32 (render at 30 and 31 is D4); Boundless Light from 34; Illumination 5/5 at 26. OK.
- **Row 6 Holy Shock, self, 30-60:** "Yourself included". OK.
- **Rows 7 to 10 Holy Shock, party:** no notes.
- **Row 11 Flash of Light, tank below 90%, 20-60:** trains 20; Oath from 25 with the rank limit; "At 60 without raid buffs, rank 4" matches the bare-60 rank row. OK.
- **Row 12 Flash of Light, everyone else below 80%, 20-60:** rank 1 at 20 to 43, rank 2 at 44 to 49 (SHOULD), rank 4 from 50 (19941, trains 42). OK.
- **Rows 13 to 16 Flash of Light, party:** no notes.
