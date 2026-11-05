# Guide accuracy round 26: Druid

Audited on the round 26 tree (`scratchpad/gr26/audit`, head ca06416cd: main after guide round 25, which carries sim round 22 through fd758eefc). No sim round 23 in the log. Active client 1.60.1.70790 (`data/builds/1.60.1.70790/`).

Files audited:

- `web/src/content/guides/druid/{index,balance,feral,feral-bear,restoration}.md`.
- Every note and row condition in `data/curated/apl/druid-{balance,feral,feral-bear,restoration}.json`, and every rendered band of `addon-data.json` `rotations/druid-*` with `race_spell_ids`, against both shipped orders in `addon-builds.json`, `trainables/druid.json`, `spellconst/druid.json`.
- Coach, read again from scratch:
  - `data/curated/coach/druid-{balance,feral,feral-bear,restoration}.md` and `.json`.
  - The druid entries in `data/pipeline/coachrules.py`.
  - The generator: `tools/coach/persona_starter.py`, `starter_overrides.py`, `starter_kit.py`.
  - The coach checks that read those rule files: `api/internal/coach/f1_core_casts.go`, `f2_uptime.go`, `f5_ranks.go`, `f6_preparation.go`, `f7_death.go`, `builddata.go`, `tables.go`, and `logs/engine/mechanics/consumables/{catalogue.json,atpull.go}`.
  - The raid preset `sim/request/raid-consumes.json`.

Read first: round 25 druid, the ledger's Round 25 section, sim round 22's `druid.md` and `integration.md`.

## Changes since round 25's tree (438f8ba15)

`git diff 438f8ba15 HEAD` on the druid files:

- **Balance card row [5] note.** This is round 25's R25-1 fix: "from Moonkin Form at 40 the loop still wins".
- **Feral engineering note 1.** Sim round 22 C2. It is not player-facing.
- **`trainables/druid.json`.** Sim round 22's platform talent-level fix gives five talent rows a level:
  - Insect Swarm 5570: 25.
  - Nature's Swiftness: 30.
  - Shifting Power: 40.
  - Swiftmend: 25.
  - Thick Hide: 15.

  With that fix, the rows now render from the shipped order:
  - Insect Swarm from 27.
  - Shifting Power from 40.
  - Nature's Swiftness from 30.
  - Swiftmend from 43.

  Round 25's routed "row-before-talent" items for these four rows are resolved. Two routed renderings remain: the Feral "Below level 20" rows past 20, and Restoration row [4], the save Healing Touch, which renders from 10.
- **All four coach `.md` files, `druid-restoration.json` and `coachrules.py`.** These carry round 25's coach fixes. Nature's Swiftness left Restoration's efficiency list.
- **No druid guide page changed.**

**Coach regeneration check.** I ran `persona_starter.py --spec druid-<spec> --out scratchpad/gr26/druid-coach/` for all four specs. All four come out byte-identical to the committed files. No druid persona is hand-written.

## Round 25 fixes, rechecked

- **R25-1, Balance card [5].** "From Moonkin Form at 40 the loop still wins, because the Wrath costs a fraction of a Starfire's mana." I checked it against both shipped orders.
  - **Leveling order.** One Eclipse point from 35 to 60. The lane measured the loop as identical at 36 and 39, and +0.9 to +1.0% at bare 40 to 60.
  - **Raid order.** Eclipse at 39, then 41 and 42. Sim round 22 G1 measured a tie at 39, +0.7% at 40 and +1.6% at 41.
  - **The mana clause.**
    - At 40 and 41, Wrath rank 6 costs 80, or 40 after Improved Wrath, against Starfire rank 3 at 180.
    - At 42 to 45, Wrath rank 6 against Starfire rank 4 at 230.
    - At 60, Wrath rank 8 at 60 against Starfire rank 6 at 315.
  - It matches balance.md:141. OK.
- **R25-2 to R25-17, coach.** Each fix is in the regenerated file. Every line is re-read in Appendix B. Three fixes left a new error or met a coach check they do not cover: R25-7 (Swipe, now R26-5), R25-8, and the rule-file pass of 210c63d93 (R26-3 and R26-4).

## Findings

Counts:

- **Guide pages and cards:** CONFIRMED 1, LIKELY 1, NEEDS-SIM 0, NEEDS-OWNER 0.
- **Coach:** CONFIRMED 4, LIKELY 0, NEEDS-SIM 0, NEEDS-OWNER 0.

| id | status | kind | where | quote | evidence | fix |
|---|---|---|---|---|---|---|
| DRUID-R26-1 | CONFIRMED | page: claim false at the levels and pack sizes it covers. The sim now decides it. | `feral-bear.md:83` (Adds) | "Swipe moves to the top on a pack, with Lacerate on the target the raid will kill last." | Sim round 22 G2 (`2026-11-05-sim-audit-round22/druid.md`) ran Swipe first against the shipped order, 10,000 iterations, both seeds:<br>- **Raid 60:** −0.9 to −1.0% on 2 targets; +8.2 to +8.5% on 3 and 4.<br>- **Bare 30:** −5.3% / −3.0% / −3.1% (2 / 3 / 4 targets).<br>- **Bare 50:** −13.6% / −5.2% / −4.7%.<br>- **Bare 55 to 59, 3+ targets:** a gain (+0.4 to +4.0%), after Swipe rank 5 at 54.<br>- **Bare 60:** a loss of 3.7 to 4.3%.<br>A two-target pack and every leveling band below 54 lose threat with Swipe on top. The line names no level and no pack size. Round 25 carried it as a Should waiting for this sim. The sim has now resolved it, so it is filed. | feral-bear.md:83: "**Adds.** In a raid at 60, on three or more targets, Swipe goes to the top. On two targets, and while leveling, keep it below Primal Bite, Lacerate and Maul. Put Lacerate on the target the raid will kill last." The rest of the paragraph stays. Pages only, no regen. |
| DRUID-R26-2 | LIKELY | card note: a superlative that holds on one band only | `druid-feral-bear.json` row [8] note, renders 16 to 60 | "It hits up to three, so on a pack of three it is your main threat" | Sim round 22's guide-lane section: "Bear card [8], 'on a pack of three it is your main threat', holds at raid 60 only." It gives no per-spell threat breakdown, so this stays LIKELY. On three targets at bare 30 and 50, the card's own order (Swipe last, above 40 rage) beats Swipe first. That makes Swipe the main threat source there unlikely. | Note: "On two or more targets, above 40 rage. It hits up to three; on a single boss, put the rage into Maul instead, and into Lacerate from 25 and Primal Bite from 40." This drops the superlative. Regenerate `addon-data.json` (note text only). See the sim request. |
| DRUID-R26-3 | CONFIRMED | coach rule file: the core check counts casts that the card and page hold back | `coachrules.py` `druid-feral` `core=["Shifting Power"]`, read by `f1_core_casts.go` | F1 detail: "{lost} of a possible {possible} went unused while you had the mana." Title "Shifting Power: {n} casts lost". | **The mana half.**<br>- spellconst 1322605 has `cost: 0`, `cost_pct: 55.0`. `ParseSpellConst` treats a percent-only cost as unknown (`CostKnown` false), so `newResourceCheck` returns `unverified`.<br>- `readyMSInGap` then counts every second off cooldown as ready, whether or not the cat had the mana.<br>- feral.md:81: "Shifting Power costs 55% of your base mana, and mana decides how many you get."<br>- A cat's 8 s cooldown would need about 85 mana a second to keep up. Every raid cat runs dry and is then told it lost the casts it had no mana for, "while you had the mana". The appended "Couldn't check whether you had the mana" contradicts that sentence; it does not soften it.<br>**The energy half.** Card row [5] and feral.md:77 cast it only at 60 energy or less, and F1 does not read energy for a mana spell.<br>The finding fires at 2 lost casts and 10%, and goes major from `CoreMajorPct`. | **Rule file:** `coachrules.py` `druid-feral` `core=[]`, with a comment: "Shifting Power is mana-limited (55% of base mana) and held above 60 energy; F1 cannot read either (round 26)." Restamp `druid-feral.json`.<br>Or coach code: let F1 read `cost_pct` against the class's base mana at `Row.Level` and gate on the card's energy condition.<br>The persona's core finding text then becomes "could not check", which is honest. |
| DRUID-R26-4 | CONFIRMED | coach rule file: uptime judged against the coach's own line and the page | `coachrules.py` `druid-feral-bear` `auras=[("Demoralizing Roar", DEF)]`, read by `f2_uptime.go` | Coach line: "Demoralizing Roar whenever the boss does not have it ... unless a warrior's Demoralizing Shout is already on it." | feral-bear.md:79 and :85, card [6] ("it does not stack with a warrior's Demoralizing Shout") and the coach line all tell the bear to leave the slot to a warrior's Shout.<br>`judgeAura` looks only for the rule's own Roar ids:<br>- `ownTrack` finds no Roar track.<br>- `appliedByOthers` finds no other Roar.<br>So it measures 0% uptime ("never applied: up for none of it") against the sim reference, and F2 fires.<br>This is the mirror of the routed Protection warrior item (Demoralizing Shout while a druid's Roar holds the slot). Round 25 fixed that one by dropping the aura (210c63d93) and left the bear side as it was. | **Rule file**, the same way: drop `("Demoralizing Roar", DEF)` from `druid-feral-bear` auras, with the warrior's comment mirrored. Restamp `druid-feral-bear.json`.<br>Or let the integrator's f2_uptime.go fix treat Roar and Shout as one slot in both directions. If it does, say so in the ledger and keep the aura. |
| DRUID-R26-5 | CONFIRMED | coach: a later clause names spells before they are trained (round 22 lens) | `coach/druid-feral-bear.md:31`; `starter_overrides.GUIDE_LINES["druid-feral-bear"]["Swipe"]` (the R25-7 fix) | "Swipe only with two or more targets in front of you, above forty rage. On a single boss, put the rage into Maul, Lacerate and Primal Bite." | Swipe trains at 16 (`trainables`). Lacerate is the 25th-level point and Primal Bite the 40th, on both orders. From 16 to 24 the line sends the rage into two buttons the player does not have, and from 25 to 39 into one. Card [8] scopes both ("and into Lacerate from 25 and Primal Bite from 40"); round 24 fixed the card for exactly this (R24-4). The coach copy dropped the scope. | **Override:** "Swipe only with two or more targets in front of you, above forty rage. On a single boss, put the rage into Maul, and into Lacerate from twenty-five and Primal Bite from forty." Regenerate with `persona_starter.py --spec druid-feral-bear`. |
| DRUID-R26-6 | CONFIRMED | coach check: preparation contradicts the persona (shared root) | `api/internal/coach/tables.go:185` (`cat.For(row.Role)`), `f6_preparation.go` `categoryIssues`, `consumables/atpull.go` `MissedAtPull`; `catalogue.json` `flask: complete` for every role | Feral persona: "Pull preparation: a missing food, weapon enchant and combat potion." Its consumables paragraph names no flask. | Round 25 R25-17 made each persona's kit follow its spec's raid preset. The preset gives feral no flask at any band. Balance, bear and Restoration get one only from 50.<br>F6 still picks the catalogue by role. It has no spec or level gate (`grep Row.Level` finds none in f6), and the flask category is "complete" for dps, tank and healer. So:<br>- Every feral cat, at every level and raid size, gets "no flask". It is major at raid size.<br>- Every druid of any spec below 50 gets "no flask" on a dungeon boss, although no flask can be drunk below 50.<br>The persona and the check disagree on the same pull. | **Coach code (shared root):**<br>- F6 reads the expected flask, food, oil and potion from the spec's preset band at `Row.Level`, the way `starter_kit.kit()` now does for the persona.<br>- A category the preset band leaves empty is never "missing".<br>This touches every spec whose preset skips a flask (rogues and hunters, as well as feral) and every sub-50 log. Test first: a feral fixture with no flask, and a level-30 Balance fixture. |

## Shoulds (wording, not findings)

- **DRUID-R26-S1, Balance coach :27, Eclipse "from thirty-five".**
  - This is the leveling order. The raid order takes Eclipse at 39, and card [5] says both.
  - balance.md tells you to level on the leveling build and respec at 60, so a pull at 35 to 38 is on the leveling order in practice.
  - Optional: "from thirty-five (thirty-nine on the raid build)".
- **DRUID-R26-S2, `RANKS_DRILL` on Feral :51 and bear :39.** "Sometimes the right one is a lower, cheaper rank."
  - The feral and bear cards cast the newest trainer rank at every band (checked in Appendix A).
  - The sentence is generic, so it is not filed.
  - Optional: give the second sentence only to specs whose card casts below the newest rank (`DRILL_LINES` or a trait).
- **DRUID-R26-S3, bear drill.uptime.** "the buff on you" and "before it drops".
  - The bear's rule file judges one debuff.
  - Faerie Fire's line (card [5]) recasts it after it drops.
  - "Refresh each where its line says" already carries the rule.
- **DRUID-R26-S4, Restoration drill.core_casts.** "Press each one as soon as the bar has room for its full return."
  - The persona's Innervate line and card [1] say below 30%.
  - It is latent: Restoration's efficiency list is empty, so F1 never fires and the drill is never shown.
  - Optional: a per-spec line naming the potion by room and Innervate at 30%.
- **DRUID-R26-S5, Restoration drill.resource.** "If you are behind it, let the HoTs and Swiftmend carry the tank."
  - This sits beside the Swiftmend line's "and your mana is keeping pace". The card gate is mana above 0.8 of the fight's pace, so a druid slightly behind still casts it, and the two lines are compatible.
  - They read as opposed. This is the carried restoration.md :73 and :75 gate Should.
- **DRUID-R26-S6, sim round 22 X1.** Leader of the Pack and Moonkin Aura are exclusive (3 Nov notes).
  - balance.md:49 and :63, feral.md:51 and feral-bear.md:48 could add "it does not stack with Leader of the Pack" (or "with Moonkin Aura").
  - The Balance coach's "a form that gives your group crit" is true but gives nothing extra beside a cat.
  - Nothing stated is false (sim lane: "a Should").
- **Carried, unchanged:**
  - Round 25 S1 (verbless card fragments in the coach lists).
  - S2 (Hurricane as a fact).
  - S3 (Pounce dropped; now excluded, done).
  - S4 (Barkskin and Frenzied Regeneration at the bottom of the bear list).
  - S5 (party Rejuvenation missing from the Restoration list).
  - S7 (healer Voice clauses on the DPS personas).
  - S8 (no opening potion on balance.md).
  - Round 24's carried page Shoulds.
- **New omission.** The feral coach list has no line for levels 10 to 19 (Moonfire, then Wrath, card [3] and [4]). F8 reads the card, so no check misfires.

## Checked and right this round

- **Every rendered rank against the page's rank prose.** All match:
  - Balance:
    - Moonfire 2 / 3 / 4 at 20–28 / 29–35 / 36–39.
    - Wrath 3 at 22–23, 4 at 24–39, 6 at 46–48.
    - Insect Swarm 1 at 27–38, 3 at 54–58.
    - Starfire newest everywhere.
  - Restoration:
    - Tank Rejuvenation 2 / 4 / 5 / 6 / 8 / 10, with the night elf rank 3 at 29.
    - Healing Touch filler 2 / 3 / 4 / 5 / 6 and top, with the tauren rank 2 to 23.
    - Party Rejuvenation 1, with 2 at 27–29 and 40–42, and 7 at 60.
    - Regrowth top.
  - Feral and bear: newest trainer rank at every band.
- **Sim round 22 C1 (Rip and Bite +25% attack power) and C2 (no Cat Form below 20).** No page or card line quotes an attack-power share or a sub-20 cat. Nothing to change.
- **Coach checks that hold for druid:**
  - **F1 on bear Primal Bite:** category cooldown 6 s read through `max(cooldown, category)`; the rage cost is known; talent with no casts is skipped.
  - **F2 on Balance Insect Swarm:** a talent aura, skipped when never up.
  - **F2 on Feral Rip.**
  - **F5:** the site's own lower ranks are exempt through `sitesOwnRank`, so Balance's and Restoration's downranks are never filed as stale.
  - **F7 defensives:** Barkskin at 44 and Frenzied Regeneration at 36 are level-gated.
  - **F8:** reads the card, including the sub-20 rows.
- **Coach consumables paragraphs.** Each matches its spec's top preset band:
  - Balance: Supreme Power, Baked Salmon, Brilliant Wizard Oil, Major Spellblasting Potion.
  - Feral: Smoked Desert Dumplings, Elemental Sharpening Stone, Major Frenzy Potion.
  - Bear: Titans, Dumplings, Elemental Sharpening Stone, Mighty Rage Potion.
  - Restoration: Distilled Wisdom, Sage's Tea, Major Mana Potion. The oil is still omitted; this is round 25's carried Should.

## Decisions

- **D1.** Round 25 D1 stands: coach lines are judged at every level a boss pull can be logged. A line that only names a button the player lacks is not filed. A line whose instruction is wrong at that level is filed (R26-5).
- **D2. R26-1 is CONFIRMED and R26-2 only LIKELY.** The sim measured the page's claim (a pack ordering) directly. The card's "main threat" is a threat-share claim that G2 did not break down.
- **D3. R26-3, R26-4 and R26-6 are coach findings although they live in rule files and coach code.** The round 26 brief puts the rule file and core list in scope, and each one makes the coach state something false that the persona and the page contradict. For R26-3 and R26-4 I recommend the rule-file change: it is the same shape round 25 used for Nature's Swiftness and the warrior's Shout. R26-6 is a shared root, filed once here.
- **D4. R26-4 is not a re-file of the routed warrior item.** That item is the warrior's Shout read as a gap while a Roar holds the slot. This is the bear's Roar read as a gap while a Shout holds it: a different rule file and spec. Round 25 dropped only the warrior's aura. If the integrator's f2_uptime.go fix makes the two one slot in both directions, R26-4 closes with it.

## Sims requested

- **Bear threat breakdown on 3 targets, shipped order (Swipe last, 40 rage).** Bare 30, bare 50 and raid 60, both seeds. Report each ability's share of threat; this settles R26-2. Same setup as sim round 22 G2 with the `-targets 3` flag.

## Engine questions

None new. Carried and held:

- Bear threat constants: Maul, Swipe and Lacerate (sim round 22 B3, needs a capture).
- Eclipse on a missed Wrath.
- The exact Rip and Bite attack-power share.
- Level 60 book ranks (pending owner).

## Sim findings

None new.

## Routed items, not re-filed

- Row-before-talent renderings that remain: the Feral "Below level 20" Moonfire and Wrath rows past 20; the Balance Eclipse Wrath row before 35; Restoration row [4] (the save Healing Touch) from 10.
- "Biggest threat hit" and the bear threat model.
- balance.md:105, a resisted Wrath and Eclipse.
- The Protection warrior Demoralizing Shout uptime (integrator, f2_uptime.go).

## Sources

- **Client 1.60.1.70790:**
  - `addon-builds.json`: druid orders unchanged since round 23.
  - `addon-data.json` `rotations/druid-*`: every band printed and compared.
  - `trainables/druid.json`: the talent-level rows above, Swipe 16, Cat Form 20, Ravage 32, Gift of the Wild 50.
  - `spellconst/druid.json`: 1322605 (cost_pct 55, cooldown 8 s), 1238070 (category cooldown 6 s, cost 200), 22812, 22842, 9898, 24977, 9896, and every Wrath, Starfire, Maul and Swipe rank.
- **Curated:**
  - `data/curated/apl/druid-*.json`.
  - `data/curated/coach/druid-*.md` and `.json`.
  - `data/pipeline/coachrules.py`.
  - `sim/request/raid-consumes.json` (every druid band).
  - `logs/engine/mechanics/consumables/catalogue.json`.
- **Coach code:**
  - `api/internal/coach/f1_core_casts.go` (`assessAbility`, `newResourceCheck`, `readyMSInGap`, `coreDetail`).
  - `f2_uptime.go` (`judgeAura`, `ownTrack`, `appliedByOthers`).
  - `f5_ranks.go`, `f6_preparation.go`, `f7_death.go`.
  - `builddata.go` (`ParseSpellConst`), `tables.go:185`, `copy.go` (CoreNoCost, DetailCoreCasts).
  - `logs/engine/mechanics/consumables/atpull.go`.
- **Generator:** `tools/coach/persona_starter.py`, `starter_overrides.py` (druid GUIDE_LINES, GUIDE_ORDER, DRILL_LINES), `starter_kit.py` (RANKS_DRILL, DRILLS). The read-only regeneration is in `scratchpad/gr26/druid-coach/`.
- **Reviews:** round 25 druid; the ledger's Round 25 section; `2026-11-05-sim-audit-round22/druid.md` (G1, G2, X1, C1, C2, guide-lane section); commit 210c63d93.
- **No web research needed.** The client data, the cards, the pages, the coach code and the sim round 22 measurements decided every call.

## Shipped orders used

Unchanged from round 25:

- **Balance leveling:** Improved Wrath 10–14, Genesis 15–16, Moonglow 17–19, Improved Moonfire 20–21, Insect Swarm 27, Improved Starfire 33, Nature's Grace 34, Eclipse 35, Moonkin Form 40.
- **Balance raid:** as leveling to 32, then Improved Starfire 33–37, Nature's Grace 38, Eclipse 39, 41 and 42, Moonkin Form 40.
- **Feral:** Ferocity 10–14, Shredding Attacks 20–22, Predatory Discipline 31 / 30, Rend and Tear from 37 / 36, Shifting Power 40.
- **Bear:** Ferocity 10–14, Furor 15, Feral Swiftness 16–17, Feral Instinct 18, 19 and 22, Lacerate 25, Primal Bite 40, Rend and Tear 41–45.
- **Restoration:** Bloom 20, Nature's Swiftness 30, Swiftmend 43, Improved Tranquility 47–48.

## Appendix A: Note sweep

The bands are the rendered bands in `addon-data.json` on this tree. Talent levels are given as leveling / raid where the orders differ.

**Balance**

- **[1] Moonfire, end of fight. Bands 10–60.** "Best instant" holds against Insect Swarm (no direct hit), Faerie Fire (no damage) and Hurricane (channelled) at every band. OK.
- **[2] Wrath, end of fight. Bands 10–60.** Starfire is scoped "at 20". OK.
- **[3] Moonfire upkeep. Bands 10–60.** Improved Moonfire 20 and Nature's Splendor are conditional. OK (carried Should).
- **[4] Insect Swarm. Bands 27–60** (now from the talent level). One global. OK.
- **[5] Eclipse Wrath. Bands 10–60, fires from 35 / 39.**
  - 0.17 s and 0.5 s match the talent text.
  - "Do not pay" with one point holds at every band.
  - "From Moonkin Form at 40 the loop still wins" holds at 40+ on both orders (lane and G1).
  - The mana clause holds at every band from 40 (Wrath 6 / 7 / 8 against Starfire 3 / 4 / 5 / 6). OK.
- **[6] Wrath filler and mana line. Bands 10–60.**
  - Starfire 20 and Innervate 40 are scoped.
  - "A fraction of a Starfire's mana for half the cast" holds from 20: Wrath rank 3, 40 halved to 20, against Starfire rank 1 at 95. OK.
- **[7] Starfire. Bands 20–60.** Eclipse at 35 / 39 is scoped. OK.

**Feral**

- **[1] Ravage. Bands 32–60.** Prowl; Poised at 32. OK.
- **[2] Faerie Fire. Bands 18–60.** Free in Cat Form from 20; 6 s. OK (band 18 and 19 Should carried).
- **[3] Moonfire below 20. Bands 10–60.** Routed rendering past 20. OK.
- **[4] Wrath below 20. Bands 10–60.** Routed. OK.
- **[5] Shifting Power. Bands 40–60** (now from the talent level). 40 energy, mana, 8 s match spellconst. OK.
- **[6] Rip. Bands 20–60.** 2 s; Rend and Tear at 37 / 36 is scoped. OK.
- **[7] Ferocious Bite. Bands 24–60.** Rip 5 s; Predatory Discipline is conditional. OK.
- **[8] Shred. Bands 21–60.** From 21; Shredding Attacks' second point at 21 on both orders. OK.
- **[9] Claw. Band 20.** 40 against 54 energy. OK.

**Feral bear**

- **[0] Cooldowns. Page only.** OK.
- **[1] Barkskin. Bands 44–60.** 20%, 15 s, 60 s. OK.
- **[2] Frenzied Regeneration. Bands 36–60.** Half health, 20 rage. OK.
- **[3] Primal Bite. Bands 40–60.** 6 s; "biggest threat hit" is routed. Not re-filed.
- **[4] Lacerate. Bands 25–60.** Five stacks; 3 s; Rend and Tear from 41 on both orders. OK.
- **[5] Faerie Fire. Bands 18–60.** Free in Bear Form, 6 s, threat. OK.
- **[6] Demoralizing Roar. Bands 10–60.** "204 at 60" is scoped; does not stack with the warrior's Shout. OK (the coach rule is R26-4).
- **[7] Maul. Bands 10–60.** 15 / 30 rage gate verified; Ferocity by 14; Lacerate from 25 is scoped. OK.
- **[8] Swipe. Bands 16–60.** Two targets and 40 rage; up to three; Lacerate and Primal Bite are scoped. "On a pack of three it is your main threat": **R26-2**.

**Restoration**

- **[1] Innervate. Bands 40–60.** Under 30%. OK.
- **[2] Read Leyline. Alliance Skyborne; no addon row.** OK.
- **[3] Nature's Swiftness. Bands 30–60** (now from the talent level). Tank under 35%. OK.
- **[4] Healing Touch save. Bands 10–60.** Routed rendering below 30. The newest trainer rank at every band. OK.
- **[5] Swiftmend. Bands 43–60** (now from the talent level). 75%, mana pace, keeps the HoT, 15 s, 20% of base mana. OK.
- **[6] Tranquility. Bands 30–60.** Two under 70%; every 2 s. OK.
- **[7] Regrowth, tank. Bands 12–60.** Top rank; Bloom from 20 is scoped. OK.
- **[8] Rejuvenation, tank. Bands 10–60.** Ranks as listed under "Checked and right"; they match the rows, including the race overrides at 29. OK.
- **[9] Healing Touch filler. Bands 10–60.** Ranks match, including the tauren override at 21–23. OK.
- **[10] Rejuvenation, party. Bands 10–60.** Under 70%; no mana gate. OK.
- **[11] to [14].** No note.

## Appendix B: Coach sweep

One line per coach line, judged at every level a boss pull can be logged.

**druid-balance.md**

- **:18 who.** OK (S6).
- **:24 Moonkin Form.** OK.
- **:25 Moonfire.** OK.
- **:26 Insect Swarm from 27.** OK.
- **:27 Wrath.** OK on the leveling order (S1).
- **:28 Starfire.** OK.
- **:29 Innervate.** OK (trainer 40; names only).
- **:30 Faerie Fire.** OK.
- **:31 Hurricane.** OK (carried S2).
- **:34–:41 finding list.** OK.
- **:43 consumables.** OK (preset 55).
- **Drills:**
  - core_casts: OK (latent; core list empty).
  - uptime: OK.
  - idle: OK.
  - resource: OK.
  - ranks: OK (Balance downranks).
  - preparation: OK, with F6 per R26-6 below 50.
  - death: OK.
  - adherence: OK.
- **Rule file:** Insect Swarm DOT talent; Barkskin. OK.

**druid-feral.md**

- **:18 who.** OK.
- **:24 Ravage from 32.** OK.
- **:25 Faerie Fire.** OK.
- **:26 Shifting Power from 40.** The text is OK; the rule file is **R26-3**.
- **:27 Rip.** OK.
- **:28 Ferocious Bite.** OK (trainer 24).
- **:29 Shred and Claw at 20.** OK.
- **:30 Claw.** OK.
- **:31 Cower.** OK.
- **:39 preparation finding.** Matches the preset. The check disagrees: **R26-6**.
- **:43 consumables.** OK.
- **Drills:**
  - core_casts: OK.
  - uptime: OK. Rip is refreshed under 2 s, "when its line says".
  - idle: OK.
  - resource: OK.
  - ranks: S2.
  - preparation: OK.
  - death: OK.
  - adherence: OK.
- **Rule file:** core Shifting Power is **R26-3**; Rip DOT and Barkskin are OK.

**druid-feral-bear.md**

- **:18 who.** OK.
- **:24 forms.** OK.
- **:25 Enrage.** OK (trainer 12).
- **:26 Primal Bite from 40.** OK.
- **:27 Lacerate from 25.** OK.
- **:28 Faerie Fire.** OK (from 18).
- **:29 Demoralizing Roar.** The text is OK; the rule file is **R26-4**.
- **:30 Maul.** OK (R25-8 fix).
- **:31 Swipe:** **R26-5**.
- **:32 Challenging Roar.** OK (trainer 28).
- **:33 Barkskin.** OK.
- **:34 Frenzied Regeneration.** OK.
- **:46 consumables.** OK (preset 50).
- **Drills:**
  - core_casts: OK.
  - uptime: S3 (the R25-10 fix is in).
  - idle: OK.
  - resource: OK.
  - ranks: S2.
  - preparation: OK.
  - death: OK. Frenzied Regeneration under half matches card [2].
  - adherence: OK.
- **Rule file:** core Primal Bite OK; aura is **R26-4**; defensives Barkskin and Frenzied Regeneration OK.

**druid-restoration.md**

- **:18 who.** OK.
- **:24 Rejuvenation pre-pull.** OK (restoration.md:69).
- **:25 Gift of the Wild.** OK (R25-11 fix).
- **:26 Regrowth and Rejuvenation.** OK.
- **:27 Swiftmend.** OK (card gate; talent 43, names only).
- **:28 Healing Touch.** OK.
- **:29 Nature's Swiftness.** OK.
- **:30 Tranquility.** OK.
- **:31 Innervate.** OK.
- **:34–:39 finding list.** OK.
- **:41 consumables.** OK (oil omitted, carried).
- **Drills:**
  - core_casts: S4.
  - uptime: OK.
  - resource: S5.
  - ranks: OK (R25-15 fix; "top rank" as trainer ranks, decided before).
  - preparation: OK.
  - death: OK.
- **Rule file:** everything empty except Barkskin. OK; this is round 25's Nature's Swiftness removal.
