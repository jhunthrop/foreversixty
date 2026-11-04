# Sim-accuracy round 17: Druid

Lane `druid`. Site wt/r17-druid, fork lane/r17-druid, both from int/r17 (site d7fc51fe7, fork f637ec32c). Client 1.60.1.70755 with the late re-layer. Specs: Balance, Feral cat, Feral bear (threat), Restoration. Measured on guide builds (`$R/rt`) and the frozen BiS (`$R/bis-frozen`).

**Summary: 0 confirmed sim bugs.** Every round-16 druid fix holds. The re-audit found nothing that moves a number. One guide-text change (seed 3) and one internal APL note fixed. Two leads stay open for lack of a log (Eclipse on a resisted Wrath; the form spread). No launch logs exist yet: foreverlogs.gg's newest report is 46 (2026-10-31), and api.foreversixty.gg/v1/reports/recent holds one report from 2026-10-20.

## Seed items

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| S1 | **Cat Form ±20% spread against the weapon's spread** (round 16 S1). No Forever launch log yet. Every paw ability is linear in the paw's average, which both readings share, so no published number depends on it. | foreverlogs.gg `/api/reports/47` and up: 404 (2026-11-04); api.foreversixty.gg recent: 1 report (2026-10-20) | OPEN (a launch cat log's white-hit min/max at a known weapon settles it) | - | - |
| S2 | **Unarmed 54.8 DPS fallback.** No log. Every published band of druid-feral and druid-feral-bear still wears a weapon (frozen BiS). | as S1 | OPEN (no published case) | - | - |
| S3 | **Rip under 2 s** (round 16 C1). The cat list recasts Rip at 5 points only when it is down or under 2 s, and Bites only while Rip has 5 s or more. `TestCatLetsRipRunToItsLastTick` green. Raid 60 Rip: 13.49 casts, 76.08 ticks. That is 7 ticks a landed Rip, less the fight-end cut, because the raid BiS idol, Idol of the Dream (220606), adds a tick. Its client spell 446212 is aura 107 misc 1 (duration) +2000 ms on Rip's mask {0x800000, 0, 0x200000}. | `sim/request/druid_audit_rotation_test.go:58`; wago SpellEffect 1143480, SpellClassOptions 44311 (70755) | Holds | - | - |
| S4 | **Genesis on Lacerate** (round 16 B1). `TestGenesisRaisesLaceratesBleed` green; `DruidSpellMaskPeriodicDamage` holds Lacerate (`druid.go:109`). | fork `sim/druid/tank/genesis_lacerate_test.go` | Holds | - | - |
| S5 | **Per-level rows** (round 16 W2-2): Moonfire r2 20-29 / r3 30-35 / r4 36-39, Wrath r3 22-26 / r4 30-39, Insect Swarm r1 30-38, tank Rejuvenation r2 20-37 / r4 38-49 / r6 50-52 / r8 58-59, party Rejuvenation r1 10-59 and bare 60. All present in `rank-choices.json` as wave 2 reported. `TestBalanceRankRowsByLevel` and `TestRestorationDruidRejuvenationKeepsItsHealingBonusInARaid` green. Wave 2 was judged after round 16's shared penalty change (fork 18815822c, rank MaxLevel; it is an ancestor of the wave-2 base 65c8e36cf), so the rows already price it in. | `data/curated/rank-choices.json`; `git merge-base --is-ancestor 18815822c 65c8e36cf` | Holds | - | - |
| S6 | **Read Leyline for Alliance Skyborne Restoration** (round 16 S3). The line casts 1259705 whenever it is ready (`druid-restoration.json`). `TestDefaultRotationRunsWithoutWarnings` (all three druid races) green. The engine matches the client: SpellEffect 1268891 triggers 1270842, aura 110 at 50 (mana regeneration %) and aura 88 at 50 (health); SpellCooldowns 120000 ms recovery, 1500 ms global. The engine puts it on spirit regeneration (Innervate's lever), which is where aura 110 acts on a Classic server. | `sim/core/racial_read_leyline.go`; wago SpellEffect/SpellCooldowns 1259705, 1270842 (70755) | Holds | - | - |
| S7 | **Restoration Leveling points at its Ranks paragraph** (round 16 wow-player SHOULD). "Levels 20 to 60" now ends: "Which Rejuvenation rank to keep on the tank changes as you level, from rank 2 at 20 to rank 4 at 38 and on up: the Ranks paragraph under [Rotation and priority](#rotation-and-priority) lists it by level range. Rank 1 stays on the party throughout." Guide tests green (532). wow-player SHIP on the rendered page (see "For the guide lane"). | `restoration.md:117` | Guide text | - | S d479bbf99, S 59affb460 |

## Feral cat (re-audit)

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| C1 | Ferocious Bite: the client roll plus 30/45/74/115/160/181 a point, 3% of attack power a point, and 0.7..3.3 an energy left over. Predatory Discipline keeps the per-energy damage but spends no extra energy. That is the 70707 talent text. No energy refund on a miss, the same as Eviscerate's (`rogue/eviscerate.go:53`). Rank 5 (22829) is the level-60 rank (rank 6 is pending owner). | `ferocious_bite.go`; talents/druid.json Predatory Discipline | Aligned | - | - |
| C2 | Rip: 6 ticks of 2 s, the client per-point table, 1% of attack power a point a tick (4 at 5 points), bleed crit (owner ruling). Shred 60 less 6 a Shredding Attacks rank (42 at 3, the talent's 18), Rake 40 less Ferocity. Shred, Ravage and Claw refund 80% on a miss; finishers refund nothing. | `rip.go`, `shred.go:55`, `rake.go:42` | Aligned | - | - |
| C3 | Pounce and Ravage both need Prowl and the target's back (`pounce.go:74`, `ravage.go:60`). This was the guide-round-1 lead ("Pounce behind-target check"). | engine | Aligned | - | - |
| C4 | Breakdown at raid 60, 3000 it.: 886.4 (before 886.1 at 10,000). Shred 40%, white 25%, Rip 22%, Windfury 7%, Bite 4%, Ravage 1%. A Shred lands 1.45 combo points: Blood Frenzy's cat clause, "your non-periodic critical strikes from Cat Form abilities that generate Combo Points have a 100% chance to add an additional Combo Point", against 40.8% crit. Clearcasting 8.5 a fight. | `cat-base.txt` (lane dir) | Aligned | - | - |
| C5 | Builds, decoded against talents/druid.json. Raid `050022/50230032023012125001/055` and leveling `050022/50212032123012123001/05501`: every point acts for a cat or is a tier gate. Feral Charge, Brutal Impact and Feral Instinct are the bear/Pounce/Prowl points the guide explains. Natural Shapeshifter 1 cheapens Shifting Power ("reduced by effects that reduce the cost of Shapeshifting"). Feral Instinct 3 in the raid build is round 16 W3's taste call. | `builds.txt` (lane dir) | Aligned | - | - |

## Feral bear (re-audit)

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| B1 | Client against engine:<br>- Barkskin 22812: aura 87 -20% physical (misc 1), 15 s (DurationIndex 8), 60 s.<br>- Frenzied Regeneration 22842: 10 s, 3 min.<br>- Enrage 5229: 10 rage up front, 2 a second for 10 s, 60 s.<br>- Primal Bite 1238070: 100% weapon plus 77, 20 rage, 6 s.<br>All match `barkskin.go:17-19` and the engine's tables. | wago SpellEffect/SpellCooldowns/SpellMisc/SpellPower (70755) | Aligned | - | - |
| B2 | Tank fight, raid 60, 3000 it.: threat 1202.6 (before 1200.3 at 10,000). Boss swings: 15.0% crush (63 vs 300 defense) and 0% parry or block, as a form gives. Primal Bite 28.8 casts in 180 s, every cooldown. Maul replaces 73 of the swings. Faerie Fire never casts because the raid debuff is already up (correct). | `bear-base.txt` | Aligned | - | - |
| B3 | Build `0/5123033202310212551/055`: every point is threat, survival or a gate. Feral Swiftness gives dodge in Bear Form; Feral Instinct is Swipe's 30%; Furor 1 is round 16 B2's kept point. | decode | Aligned | - | - |

## Balance (re-audit)

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| L1 | **Eclipse on a resisted Wrath** (guide round 13 lead). The engine banks two charges on any finished Wrath, hit or miss (`talents.go:713-718`, OnCastComplete). The client cannot settle it. Eclipse 408248's four effects are all aura 4 (dummy), so the proc is the server's script. Its SpellAuraOptions row 223856 is ProcChance 100 with ProcTypeMask 0x10000. Emulators name that bit "successful negative spell hit" (cmangos) or "done harmful magic spell", and default it to hits only (TrinityCore). The engine's own reading of it is split: `core/flags.go:27` says "Deal spell damage", while `core/harmful_spell_proc.go` counts landed or missed. A beta log was checked: report 2769, a 25-point Balance druid in Razorfen Kraul (4 boss timelines). That druid casts Wrath only and has no Eclipse aura in the log. The impact is small either way: a missed Wrath is about 11% of Wraths, and the list just casts another. | wago SpellAuraOptions 223856, SpellEffect 1072239-1343591 (70755); foreverlogs report 2769 encounters 9733/9734/9735/9740 | OPEN (a launch Balance log with a resisted Wrath followed by Starfire settles it; single-source either way) | - | - |
| L2 | Insect Swarm shows 0% miss in the breakdown, but its application does roll spell hit. `OutcomeMagicHitNoHitCounter` only keeps the miss out of the table (`insect_swarm.go`). | engine | Aligned | - | - |
| L3 | Omen of Clarity: 5.9 Clearcastings in 180 s. That fits the beta-log rate of 4% a spell, doubled in Moonkin Form, with the 10 s cooldown halved (`talents.go:322-341`, calibrated on 34 caster druids' logs). | `bal-base.txt` | Aligned | - | - |
| L4 | Nature's Grace from a Wrath crit starts when the bolt lands, not at cast end. The beta log agrees: report 2769 shows Nature's Grace applied 8-13 ms after each crit's damage event, and the cast ended about 0.8 s earlier. | foreverlogs 2769 / 9733-9740 | Aligned (log confirms a86de2495) | - | - |
| L5 | Builds `5232220115501351/0/505003` (raid) and `5232220115101151/0/555301` (leveling). Nature's Focus 5 is the gate for Naturalist, and Subtlety is threat; both are unmodelled real effects, so the adoption rule keeps them. For information: Heart of the Wild 5 in place of Nature's Focus 5 measures the same (3000 it.). | decode; `sb -talents ...-055003` | Aligned (rule keeps unmodelled talents) | - | - |

## Restoration (re-audit)

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| R1 | **Moonglow 3 in the build** (`05302001/0/555300515311321`). The 70755 text is "Reduces the Mana cost of your damaging spells by 25%", and the engine scopes it to Balance damage (`talents.go:563-573`), so it does nothing for heals. It is the filler that opens Nature's Splendor (tier 2, 10 Balance points). No other tier 0-1 Balance talent heals: Genesis and Nature's Majesty are already full. The only legal alternative drops Splendor too. Moving all four points to Improved Regrowth 5 (`05002/0/555300515311325`) measures 872.0 against 939.3 HPS (raid 60, 3000 it., -7.2%). The filler stays. | talents/druid.json; spec-breakdown -heal | Aligned (gate filler; no legal point that acts) | - | - |
| R2 | **Genesis and Gift of Nature on Swiftmend.** The 70755 client adds Genesis effect 1291838 (aura 108, misc 0, mask_1 0x2 = Swiftmend). Gift of Nature effect 0 also names Swiftmend (mask_1 0x2). The engine builds Swiftmend from the consumed HoT's snapshot, which already carries Genesis and the healing multiplier, and Swiftmend ignores attacker modifiers (`swiftmend.go`). So each talent applies once, as the client's direct effect would. | wago SpellEffect 1291838, 695243 (70755) | Aligned | - | - |
| R3 | **Bloom** (1328622) heals 100% of Spirit when a HoT runs out, and crits. Its SpellClassOptions row 85277 carries no mask, so no talent (Gift of Nature included) reaches it, and the engine registers it with no class mask. 77 Blooms a fight. | wago SpellClassOptions 85277 (70755); `bloom.go` | Aligned | - | - |
| R4 | Raid 60 breakdown, 3000 it.: 939.3 HPS (before 939.5). Mana lasts 249.6 s of 300. Swiftmend 13.6 casts (15 s cooldown), Tranquility 3 (2 min with Improved Tranquility), Nature's Swiftness 2. Regrowth crit 30.5% (Improved Regrowth 1: +10). | `resto-base.txt` | Aligned | - | - |
| R5 | The APL's internal notes still gave the tank's rank as "rank 4 from 28 to 49" (guide round 10 handed this to the sim lane). They now list the wave-2 rows. These notes are not rendered and no action changed. | `data/curated/apl/druid-restoration.json` | Note fixed (not a sim bug) | - | S d479bbf99 |

## Cross-cutting

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | Fork `./sim/druid/` package (client-conformance tests among them) green on the round-start engine. | `go test -tags=with_db ./sim/druid/` | - | - | - |
| X2 | Read Leyline with Innervate: the engine adds the two (Innervate +4, Energized +0.5 on `SpiritRegenMultiplier`). A Classic server multiplies aura-110 percentages (cmangos `GetTotalAuraMultiplierByMiscValue`): 7.5 against 5.5 times base while both are up. This reaches only Alliance Skyborne (not a published race), and only when Innervate (under 30% mana, once a fight) overlaps a 15 s Energized. | `racial_read_leyline.go`; `buffs.go:1175` | OPEN (single source, emulator behaviour; tiny and unpublished) | - | - |

## Before/after

No commit in this lane changes the engine, a rotation action, a rank row or a build, so no number moves. Before is `$R/before.tsv` (raid 60, 10,000 it.). My 3000-iteration runs on this branch sit within noise of it, seed 1, alliance:

| Spec | before.tsv | this branch (3000 it.) |
|---|---|---|
| druid-balance | 652.8 | 654.2 ± 0.7 |
| druid-feral | 886.1 | 886.4 ± 0.6 |
| druid-feral-bear (threat) | 1200.3 | 1202.6 |
| druid-restoration (HPS) | 939.5 | 939.3 |

## Files to regenerate

- `data/builds/1.60.1.70755/addon-data.json` / `addon/ForeverSixty/Data.lua` only if the addon export carries the APL's top-level `notes` (it reads per-action notes; d479bbf99 touched only the top-level string). Nothing else.
- Fork goldens: none.

## Measure at refresh

- Nothing new from this lane.

## For other lanes

- **Platform:** `/Users/jh/code/forever/web/node_modules` on main is empty. A worktree that symlinks it cannot run vitest. Run `npm ci` in the worktree instead.

## For the guide lane

- Done here: `restoration.md:117` (S7, S d479bbf99 then 59affb460).
- wow-player on the rendered page (1440, captures in `$R/druid/cap/`):
  - First pass: HOLD. MUST: the bold "Ranks" read as emphasis, not a link. Fixed: the sentence now links `#rotation-and-priority` (the click lands there).
  - Both SHOULDs taken: the sentence no longer repeats itself, and it names the first ranks (2 at 20, 4 at 38).
  - Second pass: **SHIP**.
- Open SHOULDs (platform, not text):
  - The link lands at the top of Rotation and priority, not on the Ranks lead-in. A paragraph anchor needs a markdown/renderer feature that no guide uses yet.
  - The priority card's "Rejuvenation rank 10 in a raid" can read as the leveling tank rank. The Ranks paragraph and the new sentence cover the leveling case.
- Carried, not mine to change without a log: `balance.md:105` "a resisted Wrath banks no Eclipse charges" (L1). The engine does the opposite, and neither side is proven.

## Still open

- S1, S2: form spread and unarmed paw. Need a Forever launch log.
- L1: Eclipse on a resisted Wrath. Needs a launch Balance log.
- X2: Read Leyline and Innervate stacking.
- Pending owner, not reopened: the low-rank penalty parameters, the Rejuvenation rank 1 tooltip, and Ferocious Bite rank 6 / Healing Touch and Rejuvenation 11 as level-60 book ranks.
- Warcraft Logs blocked (gcloud login); not used.
