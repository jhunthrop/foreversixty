# Guide accuracy round 15: Hunter

Auditor: hunter. Pages: class index, Beast Mastery, Marksmanship and Survival, plus the card notes in `data/curated/apl/hunter-*.json`.
- Tree: gr15/audit at 4600b7d82 (main after guide round 14, merged with sim round 17).
- Client: 1.60.1.70755. Engine pin: b2c4a2d05 (sim/enginever/version.go); engine source read from gr15/engine.
- Read first: the round 1 to round 14 hunter reports with their fix logs, round4/hunter-levels.md, the race-sims files (rounds 2 to 5), ledger.md, and design/reviews/2026-11-04-sim-audit-round17/hunter.md.

## What moved since round 14

`git diff d7fc51fe7..4600b7d82` touches these hunter inputs:
- **Guide text.** Round 14's three profession paragraphs (BM:137, MM:143, SV:124). Sim round 17 also added one "Levels 20 to 60" paragraph each on Aimed Shot rank 1: BM:157 and MM:163 (S 7e395c864, ec74057de).
- **`data/curated/rank-choices.json`.** BM Aimed Shot rank 1 now covers 28-59 (it was 40-49), plus bare 60. MM rank 1 covers 28-49, plus bare 60.
- **Engine b2c4a2d05.** Sim round 17 H2: the hunter pet now takes its own hunter's Heart of the Lion (409583 is not flagged "players only") and no other player buff.
- **Not changed.** The hunter rows in `addon-builds.json`, `talents/hunter.json`, `trainables/hunter.json`, the APL card notes and the BiS files.

The round 14 profession fixes were re-read against the BiS rows and still hold:
- Frozen Heart (249469) and Molten Heart (249470) are BoP Enchanting crafts.
- Legionite Glaive (250619) is a BoP Blacksmithing craft.
- Heartseeker (12783) binds when equipped.

## Builds and leveling order

The decode is unchanged from rounds 13 and 14:
- BM: 31/20/0, leveling and raid.
- MM: 15/31/5 leveling and 5/33/13 raid.
- SV: 5/11/35.

Every build is legal at 51 points. The shipped orders still match the prose:
- BM: Focused Fire 15-16, Endurance Training from 17, Bestial Wrath 40.
- MM: Lethal Attacks 10-14, Efficiency 15-19, Rapid Killing 25-26, Rapid Recuperation 33-34.
- SV: Improved Tracking 10-14, Savage Strikes 15-16, Survivalist 17-19, Lone Wolf last.

## Findings

### HUNTER-R15-1 (CONFIRMED, leveling): Beast Mastery's Aimed Shot rank 1 paragraph gets the reason wrong. Rank 1 does not hit "almost as hard as the top rank", and it is not "a quarter of the mana" for most of the levels the paragraph covers
- **File:** web/src/content/guides/hunter/beast-mastery.md:157
- **Quote:** "Out in the world you go out of mana before the mob dies, and rank 1 hits almost as hard as the top rank, since most of the shot is your weapon, for about a quarter of the mana."
- **Evidence:**
  - **Client values** (`spelltips/hunter.json` 19434 to 20904; `trainables/hunter.json`). Aimed Shot is "increases ranged damage by" 40 / 70 / 115 / 185 / 260 / 350, for 75 / 115 / 160 / 210 / 260 / 310 mana, at 20 / 28 / 36 / 44 / 52 / 60.
  - **Engine formula** (b2c4a2d05 `sim/hunter/aimed_shot.go`): normalized ranged weapon damage (2.8 speed: weapon average + RAP x 2.8 / 14) + ammo bonus + the flat amount. Multi-Shot is the same shot with no flat amount (`multi_shot.go`).
  - **Level 60, bare band** (`bis/hunter-beast-mastery.json`, Alliance):
    - Gear: 204 Agility, 87 Intellect, 146 AP, 21 RAP; ranged 64-120 at 2.5.
    - Ranged attack power: Night Elf base, both 10% Heart of the Lion stat bonuses, Careful Aim, Hawk at 110 and Heart of the Lion's 280 come to about 1650.
    - The shot before its flat amount: about 92 weapon + 330 from attack power + 44 ammo = 466.
    - Rank 1 hits for about 506 and rank 6 for about 816, so rank 1 is about 62% of the top rank.
    - The raid band, with more attack power, is about 65%. At 52-59 against rank 5 it is about 62%.
    - Rank 1 lands near "almost as hard" (about 80%) only at 28-35, against rank 2.
  - **Mana:** rank 1 costs 24% of rank 6 and 29% of rank 5, but 47% of rank 3 (36-43) and 65% of rank 2 (28-35). "A quarter" holds only from 52.
  - **Why rank 1 wins:** damage per mana. At 60 it is about 6.7 damage a mana point for rank 1 against 2.6 for rank 6, on fights that run dry (sim round 17 B1: the bare bands run dry for 105 to 122 s).
  - **"Before the mob dies":** a single quest mob never drains a hunter. The dry spell comes over a chain of pulls or a long fight.
- **Fix** (replace that sentence):
  > Out in the world you run dry over a chain of pulls, and rank 1 still hits for well over half of what the newest rank does, since most of the shot is your weapon, for far less mana. More damage for every point of mana keeps you shooting instead of drinking.

### HUNTER-R15-2 (LIKELY, leveling): the same reason on Marksmanship overstates rank 1 in the 40s
- **File:** web/src/content/guides/hunter/marksmanship.md:163
- **Quote:** "Out in the world you go out of mana before the mob dies, and rank 1 hits almost as hard as the higher ranks, since most of the shot is your weapon, for a fraction of the mana."
- **Evidence:** the same client values and formula as R15-1, run on the MM bare bands:
  - At 46 (band 40 gear, 134 Agility, 70 Intellect, ranged 47-88 at 2.7, Hawk rank 4, Heart of the Lion 224) the shot before its flat amount is about 327.
  - Rank 1 is about 367 against 512 for rank 4: about 72%.
  - At 28-35 it is about 80% against rank 2.
  - "Almost as hard" fits the 30s and not the mid 40s. "A fraction of the mana" is true at every level.
  - Barrage scales both ranks equally. "Before the mob dies" is wrong for the reason given in R15-1.
- **Fix** (replace that sentence):
  > Out in the world you run dry over a chain of pulls, and rank 1 still hits for most of what the newer ranks do, since most of the shot is your weapon, for a fraction of the mana.

### HUNTER-R15-3 (NEEDS-SIM, leveling): Marksmanship's "From 50 to 59 the newest rank ties with rank 1" was measured only at 50
- **File:** web/src/content/guides/hunter/marksmanship.md:163
- **Quote:** "From 50 to 59 the newest rank ties with rank 1, so cast the newest. At 60, drop back to rank 1 out in the world"
- **Evidence:**
  - Sim round 17 M1 swept MM at 28, 30, 33, 36, 39, 40, 43, 46 and 49 (rank 1 wins) and at 50 (tie), so the row ends at 49. It reports no MM result at 53, 56 or 59.
  - BM, on the same shot, finds rank 1 ahead at 50-59 (+0.1 to +0.6%). MM's own bare 60 row picks rank 1.
  - So 53-59 may favour rank 1 too, and the page then sends a player to the newest rank for nine levels and back to rank 1 at 60.
  - Not a finding until measured. If rank 1 wins anywhere in 53-59, extend the row and change the sentence to match.
- **Fix if 53-59 ties:** none. **Fix if rank 1 wins from 5x:**
  > ...keep rank 1 on your bar through 59. ... At 60, keep rank 1 out in the world, and cast the top rank in a raid, with Judgement of Wisdom on the boss and a mana potion.

### Re-checked this round and still true

- **The pet and Heart of the Lion (engine H2).** index.md:32, BM:52, BM:60, BM:101, MM:88 and SV:71 say the pet "takes no stat buffs from other players". That is still true: the only player buff it takes is its own hunter's Heart of the Lion. No line says Heart of the Lion misses the pet.
- **Lone Wolf with the stronger pet.**
  - MM:88 holds in a raid (sim round 17 M2: Lone Wolf 801.1 against the pet build 787.7).
  - SV:71 holds in a raid (S1: 861.0 against 839.9).
- **The rest of the new Aimed Shot paragraphs.**
  - Aimed Shot trains at 20 and rank 2 at 28.
  - It shares Multi-Shot's cooldown (spell text).
  - "Keep rank 1 ... all the way to 60" and "In a raid at 60 ... cast the top rank" match the BM rows (28-59, bare 60 rank 1; the raid at 60 runs the authored rank 6).
  - MM "through 49" matches its row.
  - "Judgement of Wisdom on the boss" is in the raid preset's debuffs.
  - Rank 1 (Multi-Shot's damage + 40) still outhits Multi-Shot, which has no flat amount, so index.md:33 and the leveling advice to cast Aimed Shot when the cast fits stay consistent.
  - The paragraphs do not contradict BM:83 ("you will run out of mana long before you run out of targets").
  - They read in order between the pull paragraph and the ammo paragraph.
- **Card notes.** The Aimed Shot and Multi-Shot notes on both cards say nothing about ranks, so no note contradicts the rank rows.
- **Professions (round 14 fixes), gear, enchants, consumables, races, stats.** Their inputs are unchanged since round 14's full re-read against the live BiS and race data, and that read stands.
- **Source links.** All 15 were fetched again with a browser user agent:
  - 200: the six foreversixty.gg URLs, the three Blizzard articles and the three Warcraft Tavern pages.
  - 403 (unknown under the brief, labels confirmed in round 9): Wowhead Skyborne First Look, spell=1495 and spell=19306.
- **Voice.** No em or en dashes and no site, sim, APL, band or preset words in the new paragraphs.

### SHOULDs (wording, not findings)

- **New, BM:157.** Beast Mastery never casts ranks 2 to 5, so a top player would say "skip ranks 2 to 5 at the trainer and buy rank 6 at 60" rather than "leave it in your spellbook". That saves gold, and nothing needs dragging back. (Marksmanship does cast its newer ranks at 50-59, so its wording fits.)
- **New, from sim round 17's guide note.** Add "Your own Heart of the Lion still reaches it." to BM:60 and index.md:32.
- **Carried:** MM:159 "so you have it for most hard pulls"; MM:94 five minutes on the leveling build; SV:65 "smaller than Era's"; the Serpent Sting cards' "throws away a tick"; "The Molten Core tome"; BM:81 naming Deflection; SV:118 Orc and Troll "level" with only Orc tagged.

## Decisions

- **D1. A rank paragraph's reason has to hold at every level it covers.** Rank 1 wins on damage per mana in fights that run dry, not because it hits nearly as hard. The page should say the true reason.
- **D2. Sim round 17's pet change (H2) needs no prose change.** Every pet sentence on the pages is still true as written. The addition about Heart of the Lion stays a SHOULD.
- **D3.** All earlier decisions stand.

## Sims requested

1. **New:** MM bare 53, 56 and 59, Aimed Shot rank 1 against the newest rank, 10000 iterations, seeds 1|50001, both factions (HUNTER-R15-3).
2. **New:** SV bare 60, Lone Wolf against the pet, at 10000 iterations, seeds 1|50001. Sim round 17 S1 had the pet ahead by 1.9% at 3000 iterations, seed 1. If that holds, SV:140 ("Dismiss it the moment you take Lone Wolf") should say what to do out in the world, under the owner's one-build-at-60 ruling.
3. **Carried:**
   - MM race order on the raid build.
   - Survival Crusader against Agility on both weapons.
   - BM and SV race re-runs on the re-ranked gear.
   - BM and MM with the best non-BoP two-hander in Legionite Glaive's place.
   - BM picks and weights with the pet on Heart of the Lion (sim round 17 "Measure at refresh").

## Sim findings

- **Carried, LIKELY:**
  - Falcon's melee proc does not need the talent in the engine.
  - Summon Hawk is unmodeled (owner check, sim round 17 H1).
  - The stale top-level APL notes (not rendered).
- **Closed by sim round 17 B3:** Sniper Shot's haste scaling. It has Aimed Shot's SpellMisc attributes.

## Sources

**Forever Sixty client data, 1.60.1.70755**
- `spelltips/hunter.json`: 19434, 20900 to 20904 (Aimed Shot cost and flat damage), 2643 (Multi-Shot).
- `trainables/hunter.json`: Aimed Shot ranks and levels.
- `bis/hunter-beast-mastery.json` and `bis/hunter-marksmanship.json`: band 30 to 60 rows, and the raid preset debuffs.
- `items/hunter.json`: item stats and ranged weapon damage.
- `addon-builds.json`, `talents/hunter.json`.

**Repository**
- `data/curated/rank-choices.json` (hunter rows); `data/curated/apl/hunter-{beast-mastery,marksmanship}.json` (card notes).
- design/reviews/2026-11-04-sim-audit-round17/hunter.md (H2, B1, B3, M1, M2, S1).
- Engine b2c4a2d05: `sim/hunter/aimed_shot.go`, `sim/hunter/multi_shot.go`, `sim/hunter/client_damage.go` (AimedShotDamage), `sim/hunter/hunter.go:262` (ammo bonus), `sim/core/attack.go:130` (normalized weapon damage).

**Web** (fetched 2026-11-04, browser user agent)
- The 15 frontmatter URLs, with the status codes above.

## Fix log (round 15)

- HUNTER-R15-1 fixed, beast-mastery.md:157: the reason is now damage per mana over a chain of pulls.
- HUNTER-R15-2 fixed, marksmanship.md:163: same fix.
- HUNTER-R15-3 (NEEDS-SIM): marksmanship.md:163 now reads "From 50 to 59 the newest rank and rank 1 come out close, so either serves, and rank 1 leaves more mana for the next pull. At 60, use rank 1 out in the world". This holds whether 53 to 59 ties or favours rank 1. The 53 to 59 MM sweep is still open as a SIM request.
- SHOULD (Heart of the Lion) fixed: index.md:32 and beast-mastery.md:60 now say your own Heart of the Lion still reaches the pet.
- BM:157 trainer-skip SHOULD left.
