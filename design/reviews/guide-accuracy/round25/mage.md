# Guide accuracy round 25: Mage

Audited on the round 25 audit tree (438f8ba15: main after guide round 24 and the guild r1 merge; no sim round 22 in the log; engine pin 170a59c39 in `sim/enginever/version.go`), against the ACTIVE client 1.60.1.70790 under `data/builds/1.60.1.70790/`.

Files read line by line: `web/src/content/guides/mage/{index,arcane,fire,frost}.md`; every row, gate and note in `data/curated/apl/mage-{arcane,fire,frost}.json` (notes identical to `sim/request/apl/mage-*.apl.json`, checked note by note); the rendered rows per band in `addon-data.json` (`rotations.mage-*`, first sentences match the APL notes); all six orders in `addon-builds.json` decoded level by level (index 0 = level 10) against `talents/mage.json`; the top-rank and low-rank talent texts the pages quote; `items.json` (Major Mana Potion 49, Baked Salmon 45); `sim/request/raid-consumes.json` and `data/curated/presets.json` (the raid consumables the page tables render).

New scope: `data/curated/coach/mage-{arcane,fire,frost}.md` and `.json`, and the generator `tools/coach/persona_starter.py` with `starter_guide.py`, `starter_priority.py`, `starter_kit.py`, `starter_food.py` and `starter_overrides.py`. I regenerated the three mage starters into scratch (`persona_starter.py --spec <slug> --out <scratch>`, nothing written to the tree) and diffed them against the committed files, so each coach finding says whether regenerating fixes it. `test_persona_starter.py` run read-only: 4 failures, none mage-specific beyond the committed-file match (survival aspect, protection order, feral-bear Swipe).

Read first: round 24's mage report and the ledger's Round 19 to 24 sections. Nothing routed there is re-filed.

## Round 24's fixes (new work)

All six landed and hold (50a97908b, dba977b2e):
- **R24-1 Evocation (Arcane and Fire):** "once a fight: spend your mana gem from 28 and a mana potion first, and at 60 a Demonic Rune too". Conjure Mana Agate trains at 28; a mana potion exists below 49; the three cards now agree; matches arcane.md:78 and fire.md:78. OK.
- **R24-2 Missile Barrage Missiles:** "free, or 30% of its cost inside Arcane Power"; 40%/20% with both ranks, half with one (talent ranks: 20/10 at rank 1). Matches arcane.md:53. OK.
- **R24-3 raid Frostbolt:** a quarter less than Fireball rank 10 (260 against 350) at 57 to 59, a third less than rank 11 (395) at 60. The raid build has no Frost Channeling, Arcane Power and Eureka! scale both sides alike. OK.
- **R24-4 leveling Frostbolt:** Missile Barrage 20% "once both ranks are in" (L23, R24). OK.
- **R24-5 Scorch stack note:** "From 30 your filler Scorches refresh it anyway ...; below 30 Fireball fills, so this Scorch is what keeps the stacks up." Filler Scorch gated on Heating Up (30, both orders); Improved Scorch 3/3 at 29. True at 27 to 60. OK (R24-7, the prose against the unconditional stack line, stays with the sim lane).
- **R24-6 racials:** Arcane Power and Combustion notes "your damage racial if your race has one"; arcane.md:72 "any damage racial (Eureka!, Blood Fury or Berserking)". OK. The addon rows carry the new first sentences.

The coach starters were never regenerated after round 24, so two of these fixes do not reach the coach (MAGE-R25-1).

## Findings

| id | status | kind | file:line | wrong claim | evidence | fix |
|---|---|---|---|---|---|---|
| MAGE-R25-1 | CONFIRMED | coach: stale card text (the round 24 racial fix not regenerated) | `data/curated/coach/mage-arcane.md`:27, `mage-fire.md`:27 | "Arcane Power on cooldown, together with Presence of Mind, your on-use trinkets and your racial." / "Combustion once the boss carries five Improved Scorch stacks, with your on-use trinkets and your racial." | Same error as MAGE-R24-6: the recommended Fire races (Human, Undead) and the recommended Horde Arcane race (Undead) have no damage active; "your racial" tells them to spend Will to Survive or Will of the Forsaken with their cooldowns. The cards were fixed in round 24; the starters copy the cards' first sentence and were not regenerated. | **Regenerating fixes it** (`persona_starter.py --all`): the scratch regen reads "your damage racial if your race has one" on both lines and changes nothing else in the mage files. |
| MAGE-R25-2 | CONFIRMED | coach: rule false below 30, contradicts its own list (stale override) | `mage-fire.md`:25 (from `tools/coach/starter_overrides.py` GUIDE_LINES["mage-fire"]["Scorch"]) | "Scorch on every global not spent on a three-stack Pyroblast, Fire Blast or Combustion. Every Scorch also renews the Improved Scorch debuff." | Since sim round 21 the card's filler Scorch is gated on Heating Up (30, both orders) and Fireball fills below 30 (`mage-fire.json` row 7 note, fire.md:116 "Fireball stays your nuke through the 20s"). Four lines further down the same list says "Below thirty, Fireball is the filler." Below 27 there is no Improved Scorch debuff to renew. The override's own comment quotes "Scorch is the filler for every global that isn't spent on a three-stack Pyroblast", a sentence fire.md no longer contains, so the correction no longer restates its guide. | **Regenerating does not fix it; the override must change.** GUIDE_LINES["mage-fire"]["Scorch"] → "From thirty, Scorch on every global not spent on a three-stack Pyroblast, Fire Blast or Combustion. Every Scorch also renews your Improved Scorch stacks." and its comment to quote fire.md:72 ("Everything else is Scorch") and the card's row 6 note. |
| MAGE-R25-3 | CONFIRMED | coach: a talent no Arcane build has listed as a button; the raid build's core buttons left out | `mage-arcane.md`:21 to 35 | The list "These are the buttons the guide's rotation names" includes "Ice Block." and omits Arcane Missiles, Arcane Wrath and Pyroblast; Arcane Blast sits twelfth, after Ice Block. | Neither Arcane build takes Ice Block (index.md:29 "only Frost mages carry Ice Block", arcane.md:42 "No Ice Block", arcane.md:82 "the build has no spare point for Ice Block"; decoded orders: no Frost point past Precision, Ice Shards, Improved Frostbolt and Frost Channeling). The raid loop (arcane.md:74) names Missiles on Missile Barrage, Pyroblast on 3 Heating Up stacks and Arcane Wrath; all three are card rows, and the card note calls the Missiles "a big share of your damage". Generator causes, read from `ruled_out_of`/`guide_order`: (a) "On the way to 60 there is no Arcane Wrath and no Pyroblast" hits RULED_OUT_NEXT ("there is no", "and no") and rules both out for the whole spec, though the clause is leveling-only; (b) the guide writes "Missiles", the name pattern only knows "Arcane Missiles"; (c) "has no spare point for Ice Block" is not a rule-out phrase; (d) Arcane Blast's first full-name mention is the Movement paragraph ("Blast to 3 stacks" does not match). | **Regenerating does not fix it** (scratch regen: same list). Generator: treat "On the way to 60" / "With the leveling build" clauses as conditions (CONDITION), not rule-outs; add "no spare point for" to RULED_OUT_NEXT; let the name pattern match a spell's last word when the guide uses it alone ("Missiles" for Arcane Missiles). Overrides: GUIDE_LINES["mage-arcane"] add "Arcane Missiles": "Arcane Missiles the moment Missile Barrage lights up: the proc makes it free and halves the channel.", "Pyroblast": "Pyroblast on three Heating Up stacks, when it is a one and a half second cast, with the raid build.", "Arcane Wrath": "Arcane Wrath every time it is ready while you hold three Blast stacks, with the raid build."; GUIDE_ORDER["mage-arcane"] = ("Arcane Blast", "Arcane Missiles", "Pyroblast", "Fire Blast", "Arcane Wrath", "Fireball", "Frostbolt", "Arcane Power", "Presence of Mind"). Overrides cannot remove a line, so Ice Block needs the generator change (or a removal list). |
| MAGE-R25-4 | LIKELY | coach: Missile Barrage's Missiles missing from the Frost list | `mage-frost.md`:21 to 32 | The priority omits Arcane Missiles. | frost.md:71 "Missiles the moment Missile Barrage lights up"; card row 4 casts Arcane Missiles on the proc from 51. Same generator cause as R25-3(b). LIKELY rather than CONFIRMED: the starter's own Overview line still says "fires free Missiles on Missile Barrage", so the coach is not left without it. | Fixed by R25-3(b)'s name-pattern change on regeneration; or GUIDE_LINES["mage-frost"]["Arcane Missiles"] = "Arcane Missiles the moment Missile Barrage lights up: the proc makes it free and halves the channel." |
| MAGE-R25-5 | CONFIRMED | coach: the Blast rule the page and card gave up (stale override) | `mage-arcane.md`:34 (GUIDE_LINES["mage-arcane"]["Arcane Blast"]) | "Past three, keep blasting while your mana is ahead of the fight." | The card (row 11, condition `currentManaPercent > 0.2`) and arcane.md:74 both stop past 3 stacks below 20% mana: "keep blasting only while your mana bar is fuller than the boss's health bar and you are above 20% mana". With 15% mana on a boss at 10% the coach says blast, the card and page say Fireball. The override comment quotes "Otherwise cast Arcane Blast until its buff carries three stacks. Past three, keep blasting while your mana is ahead of the fight", which arcane.md no longer contains. | **Override must change.** "Arcane Blast to three stacks. Past three, keep blasting only while your mana bar is fuller than the boss's health bar and you are above a fifth of your mana." Comment to quote arcane.md:74. |
| MAGE-R25-6 | CONFIRMED | coach: a conditional button listed bare (generator reads a heading as a line) | `mage-frost.md`:31 | "Cone of Cold on a boss." | frost.md:79 and the card's row 5 allow it only within 10 yards (12 with Arctic Reach, which the build does not take), above half mana, on 5 Winter's Chill stacks, and say "At 30 yards it never comes up" / "from 30 yards, skip it". As a priority line it tells a ranged Frost mage to Cone of Cold the boss. Cause: the bold run-in heading "**Cone of Cold on a boss.**" has six words, one over BOLD_TOPIC's MAX_TOPIC_WORDS (5), so it is kept as a sentence. | **Regenerating does not fix it.** GUIDE_LINES["mage-frost"]["Cone of Cold"] = "Cone of Cold on a boss only when you can stand within ten yards, your mana is above half and the boss carries five Winter's Chill stacks. At range, skip it."; and in `starter_guide.py` drop a bold run-in heading that ends in a full stop and starts with an ability name followed by a qualifier, or raise MAX_TOPIC_WORDS to 6. |
| MAGE-R25-7 | CONFIRMED | prose contradicting the page's own consumables table and the measurement | frost.md:95 | "A Major Mana Potion is enough for most bosses." | The Frost page's consumables table renders the raid preset (`gear-extras.ts` reads `sim/request/raid-consumes.json`): Frost's level 60 row drinks **Major Spellblasting Potion** and no mana potion (presets.json, mage-frost: "measured about 1.5% ahead of Major Mana Potion for Frost (round 10)"; commit 3df48dfd6, Nov 2). The two share the 2 minute potion timer, so a Frost mage who follows the prose drinks the potion that measured worse and contradicts the table above the line. Arcane and Fire keep Major Mana Potion (measured 0.2 to 0.7% ahead), so their lines stand. | "Frost rarely needs a mana potion, so the potion timer goes to a Major Spellblasting Potion on the pull." (Its 30 seconds line up with the opener; no number from the measurement in prose.) |
| MAGE-R25-8 | CONFIRMED | coach: the same potion, hard-coded by resource | `mage-frost.md`:44 | "a combat potion (Major Mana Potion)" | Same evidence as R25-7. `starter_kit.kit()` takes the potion from `{"rage": RAGE_POTION, "mana": MANA_POTION}[traits.resource]`, not from the preset, while the food beside it comes from the preset (`starter_food.preset_food`). The catalogue's dps role lists Major Spellblasting Potion, so the name validates. | **Regenerating does not fix it.** Generator: read the combat potion from the top raid band's consumes in `raid-consumes.json`, as `preset_food` does for food (Frost → Major Spellblasting Potion; Arcane and Fire stay Major Mana Potion). This also covers Balance, Shadow and the paladin tank, which the same preset change moved. |

Counts: CONFIRMED 7, LIKELY 1, NEEDS-SIM 0, NEEDS-OWNER 0.

Guide pages: one confirmed line (R25-7). Cards: none. Coach: six (one fixed by regenerating, five need override or generator changes).

## Shoulds (wording, not findings)

New:
- **Coach scaffolding speaks about the site and the file**, in all three starters (`persona_starter.who_section`): "as the site's guide describes it ... This file says what the guide says and leaves the rest for a later full write." It is addressed to the model, not relayed to the player (the API passes the persona as the system prompt), so not a finding, but it breaks the no-site-talk rule if the model echoes it. Better: "This is the Fire Mage. <overview> Habits beyond the guide come in a later full write."
- **Coach world buffs.** Every starter lists Rallying Cry, Warchief's Blessing, Songflower and Spirit of Zandalar. The repo's raid preset says "no world buffs (Forever has none)"; press reports say they exist in the open world and stop working in dungeons and raids (mein-mmo, gamewave; nothing official found). In a raid log they cannot appear, and the line says they are never held against anyone, so not a finding. Better: "In dungeons and raids Forever turns world buffs off, so none are expected."
- **Arcane coach Frostbolt line** gives only the raid stretch ("Frostbolt in a raid only when you are too low on mana for a Fireball"); the leveling build's Frostbolt filler is in the Overview only. Scoped "in a raid", so not false. It is the first note that leads with the name; GUIDE_LINES could give "Frostbolt fills with the leveling build; in a raid only when you are too low on mana for a Fireball."
- **Fire coach Pyroblast** "only on three stacks" leaves out the pre-pull Pyroblast and the leveling pull from range (fire.md:70, :116). Same words as fire.md:72's loop, so not a finding.
- **Frost coach "Fire Blast."** bare, last, from the Movement paragraph; not false.
- **arcane.md:74, leveling half**: "Blast to 3 stacks, Frostbolt as the filler." The card's Blast row has no build gate, so with the leveling build it also blasts past 3 while mana is ahead of the fight and above 20%. Reads as a summary, not a rule; "Blast to 3 stacks, past 3 only while your mana is ahead of the fight, Frostbolt as the filler" would match the card. Optional sim below.
- **Coach `.json` defensives** list Ice Block and Ice Barrier for Arcane and Fire, which take neither. `f7_death.defensivesAt` counts a talent defensive only if it was cast that fight, so the death review never names them; data hygiene for the coach lane.
- **Fire coach `.json` core includes Fire Blast**, which the card ranks below the stack Scorches and a 3-stack Pyroblast; if F1 reads "sitting ready while something lower went first" by the coach list's order rather than the card's, it can flag correct play. For the coach lane to check.

Carried from round 24 (still wording): the "your racial" lines followed by a list in arcane.md:76, fire.md:74, frost.md:73; fire.md:72 Fire Blast "every time it is off cooldown"; the Fireball note's per-mana reason against Frostbolt at 20 to 25; Arcane Presence of Mind on Frostbolt at 32 to 53 (raid); and the round 23 list carried there.

## Checked and right this round

- **Every superlative, comparison, quoted cost, duration and cooldown** on the four pages against the rows each band renders and the talent texts: Arcane Meditation 17/33/50% (Fire's two ranks "a third", Arcane's three "half"; with Mage Armor's half, "all" and "most"); Wand Specialization 13% at rank 1; Improved Channeling 80%/56% at four ranks ("most hits", "not every Blast"); Arcane Resilience 75% of Intellect as armor; Arcane Subtlety's resistance cut "to all your spells" (own only); Incineration covers Fire Blast, Ice Lance, Arcane Blast and Scorch (arcane.md:64 and :66, frost.md:63, fire.md:57 and :110); Critical Mass 6%; Master of Elements 30% (Fire and Frost; the page's "every Fire crit" is a true subset); Improved Fire Blast -2 s and +15 yards; Fire Power 3% an effect to 12%, mixed schools excluded (fire.md:84 Frostfire Bolt); Heating Up's four spells and the Fire Blast hit stack; Combustion; Fingers of Frost one charge at rank 1, two at rank 2 (37); Winter's Chill 40% at two ranks; Frost Warding 30% armor; Shatter 50%; Ice Shards +100% crit bonus; Arcane Mind "doubles the bonus"; Ice Barrier 15% for a minute plus 5% while shielded, no pushback.
- **Point counts and orders**, decoded: Arcane leveling 31/4/16 (Wand 10, Improved Channeling 11 to 14, Arcane Concentration 15 to 19, Blast 20, Resilience 21, Missile Barrage 22 and 23, Geometry 24, PoM 31, Arcane Power 40, Incineration 41 to 44, Improved Frostbolt 45 to 49, Precision 50 to 52, Ice Shards 53 to 57, Frost Channeling 58 to 60); Arcane raid 32/16/3 (Wrath 31, PoM 32, AP 41, Accelerant 47 and 48, Ignite 49 to 53, Pyroblast 54, Improved Fire Blast 55 and 56, Heating Up 57); Fire 17/31/3 both orders (toll Burning Soul + Accelerant, or Burning Soul 2/2; Magic Absorption moved to Arcane Resilience 3); Frost 15/5/31 one order (Ice Lance 30, Ice Block 31, Shatter 32 to 34, Cold Snap 35, Fingers of Frost 36 and 37, Winter's Chill 38 and 39, Ice Barrier 40, Missile Barrage 51 and 52). Every leveling paragraph matches.
- **Races**: unchanged prose; index.md, arcane.md, fire.md and frost.md agree with each other (Undead first on the Horde, Orc or Troll close; Gnome the free Alliance pick for Arcane and Frost, Human for Fire; Skyborne only as the paid option).
- **Enchants and consumables prose** against the rendered preset: Arcane and Fire carry Major Mana Potion and Demonic Rune (both prose lines right); Fire carries Elixir of Fire Power and Greater Arcane Elixir (carried Should); Frost's Elixir of Frost Power matches; the Frost potion is R25-7.
- **Coach lines not filed**: Evocation, Mage Armor/Arcane Brilliance, Fire Blast (Arcane and Fire), Fireball (Arcane), Presence of Mind, Arcane Explosion, Blizzard (both), Flamestrike, Frost Nova lines, Ice Barrier, Frostbolt (Frost), Ice Lance, Cold Snap, "Below thirty, Fireball is the filler", the drills and mistakes; consumables for Arcane and Fire; no em dashes, no exclamation marks.
- **Voice on the pages**: no em or en dashes, no site or sim words in body prose.

## Decisions

- **D1. Coach scope.** The priority lines, consumables paragraph and drills are what the coach relays, so they are audited as player-facing text with the guide rules. The "Who they are" scaffolding and the Voice rules address the model and are Shoulds where they talk about the site.
- **D2. R25-3 is CONFIRMED, R25-4 LIKELY.** R25-3 lists a button no Arcane build has and drops the three raid-only buttons that make up most of the raid loop beside Blast; R25-4 is one omission the same file's Overview covers.
- **D3. R25-7 is CONFIRMED** under round 14's rule (prose naming gear or consumables that contradicts the page's own tables); the preset's measurement (+1.4 to 1.5%) decides the direction, so no sim is needed.
- **D4. World buffs stay a Should.** The owner's rule and the preset say Forever raids have none; the press says they exist only outside instances. The coach never holds them against anyone, and a raid log cannot carry them.
- **D5. Regeneration.** Only R25-1 is fixed by `persona_starter.py --all`; R25-2, R25-5 and R25-6 need `starter_overrides.py`; R25-3, R25-4 and R25-8 need generator changes (`starter_guide.py`, `starter_priority.py` name pattern, `starter_kit.py` potion). Regenerate after all of them so `test_generator_is_deterministic_and_committed_files_match` goes green for mage.

## Sims requested

1. **Optional (Should on arcane.md:74).** Arcane leveling order, bare bands 30, 40 and 50, 10,000 iterations, seeds 1 and 50001, 60 s and 180 s: the Blast row as shipped (past 3 while mana is ahead and above 20%) against stopping at 3 stacks with the leveling build. Decides whether the leveling half of the loop paragraph should state the past-3 rule or the card should gate it.
2. Carried, not re-filed: MAGE-R24-7 (Improved Scorch stacking at 27 to 29), MAGE-R23-4 (Fire Blast mana gate when short).

## Engine questions

- None.

## Sim findings

- Carried from round 24: the Arcane prepull keys on Heating Up while both Presence of Mind rows key on Pyroblast known (54 raid).

## Note sweep

One line per row note: spec, row, bands it renders on (`addon-data.json`), each named talent / spell / rank / value with the level the player has it (L = leveling order, R = raid order), verdict.

**mage-arcane**
1. Evocation (12051), bands 20 to 60: Evocation 20; 8 s, +1500%, 8 min; gem from 28 (Agate), a mana potion, Demonic Rune "at 60". OK (R24-1 holds).
2. Arcane Power (12042), bands 10 to 60, live L40/R41: PoM L31/R32; 30% damage and cost, 15 s, off the global; "damage racial if your race has one". OK (R24-6 holds).
3. Presence of Mind (12043), bands 10 to 60, live L31/R32: Arcane Power "once you have it"; Pyroblast R54; Frostbolt before and on L; Blast 1.5 s. OK.
4. Pyroblast on PoM, bands 20 to 60, live R54 to 60: 6 s cast; "without Heating Up stacks" (Heating Up R57). OK.
5. Frostbolt on PoM, bands 10 to 60, live L31 to 60 and R32 to 53: Pyroblast R54; Blast 1.5 s. OK.
6. Missile Barrage Missiles, bands 10 to 60, live L22/R23: free, 30% inside Arcane Power (L40/R41); 40/20% with both ranks (L23/R24), half with one; rank 7 at 56, rank 8 a book. OK (R24-2 holds).
7. Missiles first until both ranks, bands 10 to 60, live L10 to 22, R10 to 23: second rank L23/R24; per mana over Blast once trained (20); per second over Frostbolt at every live band. OK.
8. Pyroblast on 3 Heating Up, bands 20 to 60, live R57: 1.5 s on 3 stacks. OK.
9. Fire Blast, bands 10 to 60, live R57: stack a hit, another on a crit; Improved Fire Blast R55 and 56, 6 s; leveling use. OK.
10. Arcane Wrath, bands 25 to 60, live R31: 3 Blast stacks; 5% of maximum mana. OK.
11. Arcane Blast, bands 20 to 60: 3 stacks; mana against the fight share and above 20%; 12% Intellect, +100% cost, 18 s, no refresh; Clearcasting L15/R15. OK.
12. Fireball in a raid, bands 10 to 60, live R57: Missile Barrage 20% (R24); Heating Up R57, Ignite R49 to 53. OK.
13. Frostbolt in a raid, bands 10 to 60, live R57: a quarter less than rank 10 at 57 to 59, a third less than rank 11 at 60; Missile Barrage 20%. OK (R24-3 holds).
14. Frostbolt filler on the way to 60, bands 10 to 60: stacks' spell damage (Blast 20; carried Should); free Missiles 20% once both ranks are in (L23/R24). OK (R24-4 holds).
15. Shoot (5019): no band renders it. Not rendered.

**mage-fire** (talents the notes name sit at the same levels on both orders: Ignite 5/5 L21/R19, Pyroblast 20, Improved Fire Blast 25 and 26, Improved Scorch 27 to 29, Heating Up 30, Critical Mass 34 to 36, Combustion 40)
1. Evocation, bands 20 to 60: as Arcane 1. OK.
2. Combustion (40), bands 40 to 60: 5 stacks (29); 10% crit a hit until 3 non-periodic crits; Ignite and Heating Up both in; "damage racial if your race has one". OK.
3. Scorch to 5 stacks (gated 27), bands 22 to 60: 3% a stack, 5 stacks, 30 s, own only, a miss adds none, 1.5 s refresh; "From 30 your filler Scorches refresh it anyway ... below 30 Fireball fills". OK (R24-5 holds; R24-7 routed).
4. Pyroblast on 3 Heating Up (30), bands 20 to 60: Fire Blast hit plus Scorch, Fireball, Fire Blast crits (talent text); "loses to Scorch" (carried Should at 20 to 29). OK.
5. Fire Blast, bands 10 to 60: instant; Heating Up 30; Improved Fire Blast by 26, 6 s; stack Scorches from 27 and a 3-stack Pyroblast from 30 ahead of it; "in raid gear" sentence (round 23). OK.
6. Scorch filler (gated 30), bands 22 to 60: two Scorches a Fireball; Ignite and Heating Up crits; Improved Scorch 3/3 by 29. OK.
7. Fireball, bands 10 to 60: "Below 30" scope; per mana against Scorch at 22 to 29; Improved Scorch 27. OK (carried Should).
8. Shoot: no band renders it. Not rendered.

**mage-frost** (one order)
1. Evocation, bands 20 to 60: as Arcane 1. OK.
2. Ice Barrier (40), bands 40 to 60: 15% for 1 min, 5% shielded; biggest damage buff; 30 s cooldown. OK.
3. Ice Lance (gated Fingers of Frost, 36), bands 20 to 60: 8 s; four times (+300%); Shatter 50% (34); two charges from 37. OK.
4. Missile Barrage Missiles (51), bands 10 to 60: free (no Arcane Power in the build); half channel, every 0.5 s; 20% with 2/2 (52). OK.
5. Cone of Cold on 5 Winter's Chill (38), bands 26 to 60: above half mana, 10 yards (12 with Arctic Reach); Frostbolt and Fingers of Frost Ice Lances feed it; double damage (+20% a stack), 35% spell damage (7% a stack). OK.
6. Frostbolt filler, bands 10 to 60: Fingers of Frost 36, Missile Barrage 51. OK.
7. Shoot: no band renders it. Not rendered.

## Coach sweep

One line per priority line of each starter, with the regen result.

**mage-arcane**: Mage Armor/Brilliance OK; Frostbolt (raid stretch) OK, Should; Fire Blast OK; Fireball OK; Arcane Power **R25-1** (regen fixes); Evocation OK; Arcane Explosion OK; Frost Nova, Blink bare OK; Blizzard OK; Ice Block **R25-3**; Arcane Blast **R25-5**; Presence of Mind OK; missing Missiles, Wrath, Pyroblast **R25-3**. Consumables OK. Drills OK.
**mage-fire**: Mage Armor/Brilliance OK; Pyroblast OK, Should; Scorch **R25-2**; Fire Blast OK; Combustion **R25-1** (regen fixes); Evocation OK; Fireball "Below thirty" OK; Flamestrike OK; Frost Nova and Blink OK. Consumables OK. Drills OK.
**mage-frost**: Mage Armor/Brilliance OK; Ice Barrier OK; Frostbolt OK; Ice Lance OK; Cold Snap OK (35); Evocation OK; Blizzard OK (about 1,200 = 1,400 less 15%); Frost Nova pack line OK; Cone of Cold **R25-6**; Fire Blast bare, Should; missing Missiles **R25-4**. Consumables potion **R25-8**. Drills OK.

## Sources

- data/builds/1.60.1.70790: talents/mage.json, addon-builds.json, addon-data.json, items.json.
- data/curated/apl/mage-{arcane,fire,frost}.json and sim/request/apl/mage-*.apl.json (identical notes).
- data/curated/coach/mage-*.md and .json; tools/coach/persona_starter.py, starter_guide.py, starter_priority.py, starter_kit.py, starter_food.py, starter_overrides.py, test_persona_starter.py (run read-only; regen to scratch only).
- sim/request/raid-consumes.json, data/curated/presets.json (mage-frost Major Spellblasting Potion reason), web/src/lib/guides/gear-extras.ts (the page table reads raid-consumes.json); commit 3df48dfd6.
- api/internal/coach/f7_death.go (defensives counted only if cast).
- git: 50a97908b, dba977b2e (round 24 mage fixes).
- design/reviews/guide-accuracy/round24/mage.md, ledger.md rounds 19 to 24.
- Web (D4): "WoW Forever entfernt eine super-nervige Mechanik aus Classic", mein-mmo, https://mein-mmo.de/wow-forever-keine-weltbuffs-fuer-raids/ ; "WoW Forever : Les world buffs sont-ils au programme ?", gamewave, https://gamewave.fr/world-of-warcraft-forever/wow-forever-les-world-buffs-sont-ils-au-programme/ ; "Blizzard, what is your official stance on world buffs in Forever?", Blizzard forums, https://eu.forums.blizzard.com/en/wow/t/question-world-buffs/628446
