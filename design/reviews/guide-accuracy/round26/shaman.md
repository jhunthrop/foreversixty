# Guide accuracy round 26: Shaman

Auditor: shaman. Audit tree ca06416cd (main after guide round 25, which merged sim round 22 at fd758eefc; no sim round 23 in the log), engine pin per `sim/enginever/version.go` (engine source `scratchpad/gr26/engine`), active client 1.60.1.70790.

Files audited:
- `web/src/content/guides/shaman/{index,elemental,enhancement,restoration}.md`, line by line. Changed since round 25's tree (438f8ba15): elemental.md:76 and enhancement.md:68 (8e9503456, "any damage racial").
- Every note in `data/curated/apl/shaman-{elemental,enhancement,restoration}.json` and its rendered rows in `addon-data.json`. Changed: four Elemental notes (sim round 22 E2: Flame Shock, Chain Lightning, both Lightning Bolt rows). Rows moved by sim round 22's talent-level fix (platform T1/U1): Enhancement Stormstrike now 25 to 60, Restoration Mana Tide 25 to 60, Nature's Swiftness 44 to 60, Riptide 50 to 60.
- Coach, from scratch: `data/curated/coach/shaman-{elemental,enhancement,restoration}.{md,json}`, the shaman entries of `data/pipeline/coachrules.py`, the generator (`tools/coach/persona_starter.py`, `starter_overrides.py`, `starter_kit.py`) and the coach checks that read them (`api/internal/coach/f1_core_casts.go`, `f2_uptime.go`, `copy.go`).
- Orders decoded from `addon-builds.json` (index 0 = level 10; shaman orders unchanged by sim round 22), trainer levels from `trainables/shaman.json`, client values from `spelltips/shaman.json` and `talents/shaman.json`.

Read first: round 25 shaman report, the ledger's Round 25 section, sim round 22 `shaman.md`, `platform.md` (T1, U1) and `integration.md`, and the round 26 druid, hunter, paladin, priest, rogue, mage and warlock reports (shared roots).

## Coach regeneration check

`python3 persona_starter.py --spec shaman-<spec> --out scratchpad/gr26/shaman/regen-<spec>.md` for all three: byte-identical to the committed files. Every coach finding below survives a plain regen, so each needs an override, kit or rule-file change.

## New work, rechecked

**Sim round 22 (shaman):**
- **Stormstrike from 25 (X2, T1, U1).** The Enhancement row now renders at bands 25 to 60. Both orders take Stormstrike at 25, and the engine casts it from 25. The note "it is your hardest hit" holds from 25. At 25 Earth Shock rank 4 is 130 base (8046), and a level-25 two-hander with Rockbiter rank 4 and Mental Dexterity puts 150% weapon damage around 200. At 36 to 39 Earth Shock rank 5 is 242, and Lightning Bolt rank 7 with Stormstrike's 20% is 233, both under a 380-ish Stormstrike. The near tie at 40 is still a Should. The Earth Shock note's "once you have Stormstrike" now matches the rendered rows exactly. enhancement.md:121 ("you can press it the moment you learn it ... at 25") now agrees with the card. OK.
- **Restoration rows (T1, U1).** Mana Tide 25, Nature's Swiftness 44 and Riptide 50 now match the shipped order (identical on both orders). Their notes need no other level. OK.
- **Elemental notes (E2).**
  - Flame Shock: "from 25, with all three points of Call of Flame, it ticks for 21 seconds". Call of Flame is points 23 to 25 on both orders, and 12 + 9 = 21. OK.
  - Chain Lightning: "for the same cast time as a Lightning Bolt, until Lightning Mastery makes the bolt faster". Every bolt rank from 3 up is 2.5 s, and both spells lose 0.5 s to Elemental Alacrity (20 to 22). True. "until" reads as permanent from 40, though, so it is a Should (S1).
  - Top-rank bolt: "your mana percent above the share of the fight still to go" matches `currentManaPercent > remainingTimePercent`. OK.
  - Filler: "at the rank your mana can carry to the end of the fight". Rank 10 in a raid at 60, rank 8 bare at 60, and the leveling list match `rank-choices.json` (unchanged). The filler ids exclude rank 9. OK.
- **Builds.** No shaman order, talent string or rank row changed in sim round 22 (the diff of `addon-builds.json` and `rank-choices.json` touches mage, rogue, warlock, hunter, priest and paladin only). The Elemental page's build paragraph was re-decoded anyway and holds:
  - Leveling: 31 Elemental by 40, Thundering Strikes 41 to 45, Tidal Focus 46 to 50, Mindfulness 51 to 53, Totemic Focus 54 to 58, Improved Reincarnation 59 to 60.
  - Raid: Ancestral Knowledge 46 to 50 in place of Totemic Focus.

**Round 25 fixes:**
- **Coach R25-1 to R25-17.** Each fix is in the regenerated files. Two left a gap: the level rule of R25-10 and R25-16 missed lines the overrides do not cover (R26-3, R26-5).
- **Page "any damage racial" (8e9503456).** Elemental: R26-1. Enhancement: OK, since Berserking and Rage of the Farseer are both 3 minutes and the Races section says Berserking "lines up with Rage of the Farseer just as well".

## Findings

| id | status | kind | where | quote | evidence | fix |
|---|---|---|---|---|---|---|
| SHAMAN-R26-1 | CONFIRMED | page: a round 25 edit contradicts its own paragraph, the card and the Races section | `elemental.md:76` | "Press it on the pull with your trinkets and any damage racial (Blood Fury or Berserking), and again every 3 minutes; the racial comes back sooner, so use it whenever it is ready. ... A Troll saves Berserking for after it ends." | **Self-contradiction.** The paragraph tells a Troll to press Berserking with Lightning Mastery, then three sentences later to save it for after.<br>**Contradicts the rest of the page and the card.** elemental.md:94 says "Troll's Berserking ... goes outside Lightning Mastery: inside it your bolts are already a 1.5 second cast, and casting speed never shortens the 1.5 second global cooldown". Card row 1 says "A Troll presses Berserking after Mastery ends".<br>**"comes back sooner" is false for Berserking.** Its cooldown is 3 minutes (races.json: "10% for 15 seconds on a 3 minute cooldown"), the same as Lightning Mastery's 180 s (16166). Only Blood Fury (2 minutes) comes back sooner. | Restore round 24's wording: "Press it on the pull with your trinkets and Blood Fury, and again every 3 minutes; Blood Fury comes back sooner, so use it whenever it is ready. ... A Troll saves Berserking for after it ends." Do not touch enhancement.md:68 (true there). |
| SHAMAN-R26-2 | CONFIRMED | coach: the same Berserking error | `coach/shaman-elemental.md:25`; `GUIDE_LINES["shaman-elemental"]["Lightning Mastery"]` | "Lightning Mastery from forty with your trinkets and any damage racial on the pull, and again every three minutes." | For a Troll this says Berserking goes with Mastery. The card (row 1) and elemental.md:76 and :94 say after it ends: under Mastery the bolt is already at the 1.5 s global floor. Same root as R26-1 (8e9503456 changed both). | Override edit: "Lightning Mastery from forty with your trinkets and Blood Fury on the pull, and again every three minutes. A Troll presses Berserking after it ends." Regenerate. |
| SHAMAN-R26-3 | CONFIRMED | coach: buttons named without the level they arrive (the generator's own rule; R25-10 left these) | `coach/shaman-enhancement.md:27, :34, :35` | ":27 If your party has no other shaman, put Mana Spring Totem in your Water slot as well."<br>":34 Earth Shock last ... Keep that last quarter for Stormstrike and your totems."<br>":35 Magma Totem and Fire Nova ..." | **Trainer and order levels.** Mana Spring Totem and Magma Totem train at 26. Stormstrike is the 25th-level point on both orders.<br>**Where it bites.** enhancement.md:121 calls Enhancement the shaman's leveling spec. A dungeon boss at 15 to 25 is coached to drop a Mana Spring the player does not have, and at 10 to 24 to save mana for a Stormstrike they do not have.<br>**The card scopes both.** Earth Shock row 7: "and for Stormstrike once you have it". The Mana Spring row is Elemental-only, from 26.<br>**The rule.** `starter_overrides.py`'s GUIDE_LINES header says "A line that names a spell, talent or rank the player gains during leveling says from when". The Elemental file already does it for both totems ("from twenty-six"). | Overrides, then regenerate:<br>- "Mana Spring Totem": "Mana Spring Totem from twenty-six in your Water slot when your party has no other shaman."<br>- "Earth Shock": "Earth Shock last, on cooldown while your mana is above a quarter. Keep that last quarter for your totems, and from twenty-five for Stormstrike."<br>- Magma Totem: as R26-4. |
| SHAMAN-R26-4 | CONFIRMED | coach: a pack sentence listed as a rotation button, with its condition dropped | `coach/shaman-enhancement.md:35` (generated from enhancement.md:85, no override) | "Magma Totem and Fire Nova both hit within a few yards of your Fire totem, and in melee you are standing at it." | The list header says "Your buttons, in the order you press them. Where a button has a reason or a condition, it follows the name." This line names two buttons with no condition. The page's sentence is the lead-in to "On a big pack, drop Magma Totem, use Fire Nova on its 10 second cooldown and Chain Lightning ...". On one target the page says the opposite (enhancement.md:83): "Outside a proc, Fire Nova on its own cooldown costs more mana than its damage is worth on one target". The card never casts Magma Totem for Enhancement, and casts Fire Nova only on Clearcasting (row 6). The same file's :33 limits Fire Nova to a Clearcasting proc, so :35 contradicts it on a single boss. | GUIDE_LINES `shaman-enhancement` "Magma Totem": "Magma Totem from twenty-six on a big pack, with Fire Nova on its own cooldown and Chain Lightning while your mana lasts. Both totems' bursts reach a few yards, and in melee you stand at the totem." (Each sentence is under 20 words, with no digits.) Regenerate. |
| SHAMAN-R26-5 | CONFIRMED | coach: cleansing buttons named without the level they arrive (R25-16 left these) | `coach/shaman-restoration.md:32, :33` | ":32 Poison Cleansing Totem and Disease Cleansing Totem take the Water slot and clear your party every few seconds; drop one on fights that poison or disease."<br>":33 Cure Poison and Cure Disease reach anyone the totems do not." | **Trainer levels** (`trainables`): Cure Poison 16, Poison Cleansing Totem 22, Cure Disease 22, Disease Cleansing Totem 38.<br>**Where it bites.** restoration.md:116 sends a Restoration shaman to Wailing Caverns at 18. There, and on any poison or disease boss at 16 to 37, the coach names a totem or a cure the player does not have yet.<br>R25-16 put levels on the other Restoration lines. The file's own rule (as R26-3) covers these too. | Overrides:<br>- "Poison Cleansing Totem": "Poison Cleansing Totem from twenty-two, or Disease Cleansing Totem from thirty-eight, in the Water slot on fights that poison or disease."<br>- "Cure Poison": "Cure Poison from sixteen and Cure Disease from twenty-two reach anyone the totems do not."<br>Regenerate. |
| SHAMAN-R26-6 | CONFIRMED | coach kit: the healer "core abilities" mistake teaches the opposite of the spec's own mana-cooldown rule | `coach/shaman-restoration.md:37`; `starter_kit.HEALER_MISTAKES["core_casts"]` | "Core abilities under-used: a mana or efficiency cooldown sitting ready long after the bar had room for it." | Restoration's only efficiency entry is Mana Tide Totem (`coachrules.py`). Its rule is: drop it under a quarter of your mana. That rule is in restoration.md:84, card row 1, the persona's own :26 and its drill at :48.<br>Rank 4 (1325256) returns 290 every 3 s for 12 s, 1,160 in all. So "the bar had room for it" comes at about 80% of a raid shaman's pool.<br>The "what to look for" line calls a shaman who holds Mana Tide from 80% down to 25%, as the page says, a mistake.<br>Round 25 fixed the drill text (R25-15) through `DRILL_LINES`, but not this line, and no override hook reaches it. The same text sits in `druid-restoration.md` (Innervate under 30%), `paladin-holy.md` and `priest-discipline.md`. Filed once here as the shared root. | Kit change: word `HEALER_MISTAKES["core_casts"]` for what F1 measures, for example "a mana cooldown left unused through a long fight when its line says to press it". Or add a per-spec MISTAKE override like `DRILL_LINES`. For shaman: "Mana Tide Totem left unused through a long fight, or dropped while your mana was still above a quarter." Regenerate the four healer starters. |

Counts, guide pages and cards: CONFIRMED 1 (R26-1), LIKELY 0, NEEDS-SIM 0, NEEDS-OWNER 0.

Counts, coach: CONFIRMED 5 (R26-2 to R26-6), LIKELY 0, NEEDS-SIM 0, NEEDS-OWNER 0.

By fix route:
- Page edit: R26-1.
- Override (`starter_overrides.py` GUIDE_LINES): R26-2, R26-3, R26-4, R26-5.
- Kit (`starter_kit.HEALER_MISTAKES`, or a new per-spec hook): R26-6.
- No rule-file or coach Go change is needed for shaman:
  - Elemental `core` is empty; it tracks the Flame Shock DoT, which both builds keep up at every band.
  - Enhancement `core` is Stormstrike. It is a talent, now learned at 25 in `spellranks`, 6 s, and no shaman talent in the builds changes its cooldown.
  - Restoration: Water Shield aura (point 20, both orders); efficiency Mana Tide 300 s (no talent changes it).

## Shared roots (filed elsewhere; named once for shaman, not counted)

- **DRUID-R26-6 (F6 flask by role).** The Enhancement persona names no flask, matching its preset, but F6 still judges a flask for every dps. The same goes for flasks below 50 on any shaman spec's dungeon logs. The Restoration persona omits the weapon enchant its preset carries (Brilliant Mana Oil), which F6 may judge.
- **DRUID-R26-3 / PALADIN-R26-5 (percent-of-base-mana costs).** Stormstrike costs 8% of base mana (`cost_pct` 8.0, `cost` 0), so F1's resource check is unverified. The detail then says "went unused while you had the mana" next to "Couldn't check". The brief says percent costs are now handled in code. Stormstrike is never held by the page (the quarter reserve exists to pay for it), so no shaman-specific harm.
- **Template Voice block.** The healer lines sit in the Elemental and Enhancement files (round 25 root).

## SHOULDs (wording, not findings)

- **S1. Elemental card row 5 (Chain Lightning), 32 to 60.** "until Lightning Mastery makes the bolt faster" reads as a permanent change from 40. Outside its 15 s, both spells are still 2.0 s casts. "except while Lightning Mastery is up" says it exactly.
- **S2. coach/shaman-enhancement.md:25.** The Windfury Totem line still gives no Air totem for when another shaman has Windfury down. The card's only Air row is Grace of Air (42 to 60), and enhancement.md:81 names it. Suggest adding "Otherwise, from forty-two, Grace of Air Totem."
- **S3. coach/shaman-enhancement.md:33.** "with the raid build" is true, but Elemental Focus is the raid order's 52nd point. "with the raid build, from fifty-two" would follow the file's level rule. The line is a conditional, so it is never false.
- **S4. coach/shaman-enhancement.md:36.** Lightning Shield gives no instruction ("which behind a boss is never") and drops "Keep it for leveling and for adds that turn on you" (carried from round 25).
- **S5. coach/shaman-restoration.md.** Tremor Totem (18, restoration.md:116 for Wailing Caverns) and the "Everyone else" rules are absent (carried).
- **Carried as Shoulds from round 25:**
  - Elemental persona: "Core abilities under-used" against an empty core list. Consumables omit Demonic Rune and Greater Arcane Elixir.
  - Enhancement "Who they are": five-stack bolt "free instant".
  - Stormstrike "hardest hit" near a tie at 40.
  - Grace of Air "rank 2 at 56" with rank 3 at 60.
  - elemental.md:74 "about twice".
  - Healing Wave note "17 levels past".
  - enhancement.md:74 "costs you nothing".
  - Elemental Flame Shock "nothing else wants the shock cooldown".
  - Fire Nova and Lightning Bolt rows rendering before their talents (round 24 D3).

## Checked and right this round

- **Page prose, full sweep.** Every superlative, "only", "every", "never" and quoted value was re-read against the unchanged rows, the decoded orders and the client. These hold:
  - Lightning Bolt 198 / 171 with Convection; rank 9 "about 16% less" (295 / 353).
  - Chain Lightning 2.5 s / 6 s / 495.
  - Elemental Alacrity 0.5 s.
  - Call of Flame 12 + 9 = 21 s, "three and a half shock cooldowns".
  - Trial by Fire 10 - 4 = 6 s and Magma 20 + 20 = 40 s.
  - Elemental Warding 2 points = 7%.
  - Ancestral Knowledge 10%.
  - Searing Totem rank 6: 55 s, 20 yards.
  - Mana Spring 10 every 2 s for 5 minutes.
  - Fire Nova 10 yd.
  - Lightning Mastery 15 s / 3 min.
  - Rage of the Farseer 30% / 25 s / 3 min.
  - Totemic Call 25% base mana, 30 s.
  - Stormstrike 150% weapon damage, 6 s, 20% for 12 s.
  - Shamanistic Focus 45%.
  - Elemental Weapons 40% / 20%.
  - Mental Dexterity 150% Int, Mental Quickness 30%.
  - Maelstrom Weapon 20% a stack.
  - Mana Tide 12 s every 5 minutes, ranks 25 / 38 / 48 / 58.
  - Riptide 445 / 575 / 805 over 15 s, 12 s, ranks 50 and 60 trained.
  - Nature's Swiftness 3 minutes.
  - Chain Heal 247 / about 300 / about 385, 490 / 335.
  - The 17-level falloff list (Healing Wave 8 to 10 full, 7 "most", 6 "under half", Lesser Healing Wave 3 "about two thirds", 23 or earlier none).
  - Trainer levels: Call of the Elements 20, Totemic Call 22, Ghost Wolf 20, Chain Heal 40, Lesser Healing Wave and Healing Stream 20, Tremor 18.
- **Leveling prose against the orders.**
  - Enhancement: Thundering Strikes 10 to 14, Ancestral Knowledge 15 to 19, Mental Dexterity 20 to 22, Elemental Weapons 23 to 24, Stormstrike 25, Elemental Weapons 26, Shamanistic Focus 27, Flurry by 32, Improved Ghost Wolf 33 and 34, Maelstrom 37 to 42.
  - Elemental: Convection 10 to 14, Concussion 15 to 19.
  - Restoration: Improved Healing Wave 10 to 14, Mindfulness 15 and 16, Tidal Focus from 17 and full at 22, Water Shield 20, Mindfulness 23, Mana Tide 25, the rest by 30, Riptide 50, then Thundering Strikes and Ancestral Knowledge.
- **Coach lines that hold** (Appendix B), including every round 25 fix not listed above.

## Decisions

- **D1. R26-1 and R26-2 are CONFIRMED, not NEEDS-SIM.** The page contradicts itself, the card and its own Races section. "Comes back sooner" is false for Berserking on the client's cooldowns. Whether the engine should hold Berserking is a separate engine question (below). The fix restores the text the card and the Races section already agree with.
- **D2. Coach level findings are one per file** (round 25 D1, HUNTER-R26 D5). R26-4 is separate from R26-3 because its fault is the dropped pack condition, not only the level.
- **D3. R26-6 is filed for shaman as the shared root.** The other three healer files carry the same line. No round 26 report has filed it (searched for "had room" in round26/*.md).
- **D4. Not re-filed:** the ledger's Round 19 to 25 routes; rows rendering before their talent; card action ranks (Healing Wave 9 at 60 with rank 10 at the trainer, Grace of Air 2, Strength of Earth 4); BiS race captions; the Leveling strip; the Assassination and Protection items named in the brief.

## Sims requested

- **Troll Elemental Berserking timing.** Raid 60, guide raid build, frozen BiS, Troll, 10,000 iterations. Compare the current card (autocast fires Berserking with Lightning Mastery on the pull) with a card that holds Berserking until Mastery ends (for example `not auraIsActive(16166)` on a Berserking line ahead of the autocast). Two outcomes:
  - The hold wins or ties: R26-1's fix stands and the engine follows the note.
  - Pressing it with Mastery wins by more than the error: rewrite card row 1, elemental.md:76 and :94, and the coach line together.
- Carried: the Enhancement Earth Shock gate run is done (sim round 22 N2, aligned).

## Engine questions

- **Berserking under Lightning Mastery.** Card row 1 is an `autocastOtherCooldowns` action. The engine registers Berserking as an ordinary racial cooldown (`sim/core/racials.go:466-480`) with no hold, so a simulated Troll presses it with Mastery on the pull. The card's note says to press it after Mastery ends. The published Troll Elemental numbers therefore come from a different line than the one the page teaches. The sim above settles which is right.

## Sim findings

- None new. Carried from sim round 22 "Still open": Maelstrom Weapon flat chance against 20 a minute; the sub-20 damage penalty (E1, OPEN-low; the Elemental rank 3 rows at 26 and 27 and 30 to 33 depend on it).

## Sources

No web research was needed. The client data, the card rows, the presets and the coach code decided every call.

- **Client 1.60.1.70790:**
  - `spelltips/shaman.json`: 421, 930, 2860, 10605; 529 to 15208; 8046, 10412, 10414; 17364; 16166; 425336; 437009; 16188; 16190, 17354, 17359, 1325256; 1239242, 1239243; 10497; 10587; 10438; 408345; 29228; 1064, 10623; 10396, 25357; 10468; 10463.
  - `talents/shaman.json`: all three trees.
  - `trainables/shaman.json`.
  - `races.json`: Berserking 3 minutes, Blood Fury 2 minutes.
  - `addon-builds.json`: six shaman orders, decoded.
  - `addon-data.json`: shaman rotations, by band.
- **Curated:** `data/curated/apl/shaman-*.json`, `data/curated/rank-choices.json`, `data/curated/coach/shaman-*.{md,json}`, `data/pipeline/coachrules.py`, `sim/request/raid-consumes.json`.
- **Coach:** `tools/coach/starter_overrides.py`, `starter_kit.py`, `persona_starter.py` (scratch regen at `scratchpad/gr26/shaman/`); `api/internal/coach/f1_core_casts.go`, `f2_uptime.go`, `copy.go`.
- **Engine:** `sim/core/racials.go`.
- **Reviews:**
  - Round 25 shaman report and the ledger's Round 25 section.
  - Sim round 22 `shaman.md` (E2, N2, N5, X2), `platform.md` (T1, U1) and `integration.md`.
  - Round 26 druid, hunter, paladin, priest, rogue and warlock reports.
- **Git:** 8e9503456 (round 25 racial lines), 0a7c63a7a and 6903cd50b (sim round 22 Elemental notes), b4a82d5ad and 6c639ea7f (talent learn levels and card rows).

## Appendix A: Note sweep

Orders:
- Elemental: leveling and raid are identical to 45. Call of Flame 23 to 25, Elemental Focus 26, Trial by Fire 29 and 30, Lightning Mastery 40.
- Enhancement: Stormstrike 25 on both. Maelstrom Weapon 37 to 42 leveling, 35 to 39 raid. Rage of the Farseer 40 on both. Elemental Focus raid only, at 52.
- Restoration: both orders identical. Water Shield 20, Mana Tide 25, Nature's Swiftness 44, Riptide 50.

Bands are the rendered `addon-data.json` levels.

**Elemental**
- [0] Mana Spring, 26-60: trains at 26; skip-if-stacked; many times its cost. OK.
- [1] cooldowns (autocast): Lightning Mastery 40; 15 s; Blood Fury; Troll after. OK. (The page contradicts it: R26-1.)
- [2] Flame Shock, 10-60: "from 25" (Call of Flame 23 to 25 on both orders); 21 s. OK.
- [3] Searing Totem, 10-60: only damaging Fire totem before 26; Fire Nova from 12; 20 yd; Magma from 26. OK.
- [4] Magma Totem, 26-60: 20 s; 40 from 30 (Trial by Fire 29 and 30). OK.
- [5] Chain Lightning, 32-60: rank 1 215 > top bolt 148 at 32; rank 4 510 > 353; 70% a jump; same cast time (Alacrity on both). OK, with S1.
- [6] Fire Nova, 12-60: Searing until 26, Magma after; 10 yd. OK.
- [7] Lightning Bolt top rank, 10-60: condition text matches; 198 / 171; "about 20%". OK.
- [8] Lightning Bolt filler, 10-60: matches `rank-choices.json`; rank 10 raid 60, rank 8 bare 60. OK.

**Enhancement**
- [0] autocast: no note.
- [1] Strength of Earth, 10-60: 30 yards. OK.
- [2] Grace of Air, 42-60: trains at 42, rank 2 at 56; Windfury Totem 32. OK (rank 3 Should).
- [3] Stormstrike, 25-60 (new bands 25 to 39): point 25 on both orders; 6 s; 20%; hardest hit at 25 to 39 (see New work). OK.
- [4] Lightning Bolt at five stacks, 10-60: `tag 5` = 42 leveling, 39 raid; "Cast it right after Stormstrike when you can" (25). OK.
- [5] Searing Totem, 10-60: 20 s guard. OK.
- [6] Fire Nova on Clearcasting, 12-60: "in the raid build" (52). OK.
- [7] Earth Shock, 10-60: quarter gate; "once you have Stormstrike" (25). OK.

**Restoration**
- [0] autocast: potions, runes; Nature's Swiftness saved. OK.
- [1] Mana Tide, 25-60: point 25; under 25%; 12 s; rank 4 at 58. OK.
- [2] Water Shield, 20-60: three globes, 2%. OK.
- [3] Healing Stream, 20-60: every 2 s for 5 minutes. OK.
- [4] Nature's Swiftness, 44-60: point 44; tank under 25%. OK.
- [5] Healing Wave after Nature's Swiftness, 10-60: highest rank (D4 for the bands before 44). OK.
- [6] Riptide, 50-60: point 50; 12 / 15 s; recast when it has run out; 25%. OK.
- [7] Lesser Healing Wave tank, 20-60: 1.5 s. OK.
- [8]-[15] party rows: no note.
- [16] Chain Heal, 40-60: ranks 2 (46) and 3 (54); 247 / about 300 / about 385; rank 2 bare at 60. OK.
- [17] Healing Wave tank, 10-60: rank 8 floor at 60; rank list matches `rank-choices.json`; "ranks 1 to 4 keep none". OK.

## Appendix B: Coach sweep

One line per coach line: file:line, claim, the level the player has it, verdict.

**shaman-elemental.md** (regen identical)
- :18 Who they are: page overview. OK.
- :24 Mana Spring (26), skip clause. OK.
- :25 Lightning Mastery (40) with any damage racial. **R26-2**.
- :26 Flame Shock (10), ahead of everything but cooldowns. OK.
- :27 Searing Totem (10), packs until Magma at 26, 20 yd. OK.
- :28 Magma Totem (26). OK.
- :29 Chain Lightning (32), first hit beats a bolt, 70%. OK.
- :30 Fire Nova (12), packs; Clearcasting to the bolt. OK.
- :31 Lightning Bolt, top rank while ahead, cheaper behind. OK.
- :32 Earth Shock (4) for interrupts. OK.
- :35 to :42 finding list: core_casts against an empty core (Should); ranks text honours cheaper ranks. OK.
- :44 consumables: Supreme Power, Baked Salmon, Brilliant Wizard Oil, Major Mana Potion. These match the preset (Should: Demonic Rune, Greater Arcane Elixir).
- Drills:
  - :48 core_casts, :50 idle, :51 resource, :52 ranks, :53 preparation, :54 death, :55 adherence. OK.
  - :49 uptime ("when its line says, never earlier"). OK.
- :57 to :68 Voice: template root.
- Rule file: core []; Flame Shock DoT, all six ranks. OK.

**shaman-enhancement.md** (regen identical)
- :18 Who they are: Should (five stacks).
- :24 Windfury Weapon (30), Rockbiter before 30 and under a raid Windfury Totem. OK.
- :25 Windfury Totem (32). OK (S2).
- :26 Strength of Earth (10), 30 yd. OK.
- :27 Mana Spring, no level (26). **R26-3**.
- :28 Totemic Call (22). OK.
- :29 Rage of the Farseer (40), trinkets and any damage racial. OK (Berserking and the Rage are both 3 minutes).
- :30 Stormstrike (25), 20%. OK.
- :31 Lightning Bolt at 5/5 (42 leveling, 39 raid). OK.
- :32 Searing Totem (10). OK.
- :33 Fire Nova on Clearcasting with the raid build (52). OK (S3).
- :34 Earth Shock quarter gate; Stormstrike reason before 25. **R26-3**.
- :35 Magma Totem and Fire Nova, no condition, no level. **R26-3, R26-4**.
- :36 Lightning Shield. S4.
- :39 to :46 finding list. OK ("food and combat potion" matches the preset; F6 is the shared root).
- :48 consumables: Smoked Desert Dumplings, Major Frenzy Potion (preset). OK.
- Drills :52 to :59. OK.
- Voice: template root.
- Rule file: core Stormstrike (talent; learned at 25 in `spellranks`). OK.

**shaman-restoration.md** (regen identical)
- :18 Who they are. OK.
- :24 Water Shield (20). OK.
- :25 Healing Stream (20). OK.
- :26 Mana Tide (25), under a quarter. OK.
- :27 Riptide (50), after it runs out. OK.
- :28 Chain Heal (40), mana ahead, Riptide 25%. OK.
- :29 Healing Wave at 60%, rank 8 floor at 60. OK.
- :30 Lesser Healing Wave (20). OK.
- :31 Nature's Swiftness (44) into the highest rank. OK.
- :32 Poison Cleansing Totem (22) and Disease Cleansing Totem (38), no level. **R26-5**.
- :33 Cure Poison (16) and Cure Disease (22), no level. **R26-5**.
- :34 Ancestral Spirit (12). OK.
- :37 core_casts mistake: "long after the bar had room for it". **R26-6**.
- :38 to :42 finding list. OK.
- :44 consumables: Distilled Wisdom, Sage's Tea, Major Mana Potion. OK (Should: oil; F6 shared root).
- Drills:
  - :48 core_casts (Mana Tide under a quarter). OK.
  - :49 uptime, :52 preparation, :53 death. OK.
  - :50 resource. OK.
  - :51 ranks (rank 8 floor, top rank for Nature's Swiftness). OK.
- Voice: OK.
- Rule file: Water Shield self buff (talent); efficiency Mana Tide only. Nature's Swiftness was dropped per R25-17. OK.
