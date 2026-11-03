# Sim audit round 13: regression lane, casters and healers (regcast)

Scope: every round-12 change that reaches mages, priests, warlocks, Balance, Elemental and the five healers, re-checked against the 1.60.1.70753 client rows and in the sim, at every band and both factions; then a fresh re-audit of every caster and healer pre-pull line and opener.

Engine branch `lane/r13-regcast` and site branch `wt/r13-regcast`, both at the round-13 start (fork 076a5a517, site 93c5d2510). Client tables: round 11's 70753 copies (scratchpad r11/c70753, SpellMisc, SpellEffect, SpellCooldowns, SpellPower, SpellAuraOptions, ItemXItemEffect); tooltips from nether.wowhead.com/forever, read 2026-11-03.

**Result: 1 confirmed bug (old, latent, no published number), routed to the casters lane with a ready patch. No round-12 change caused a bug in this group. No commit of mine moves any number.**

## How it was checked

1. **Client rows.** A row printer over the cached tables (cast time, duration, missile speed, cooldowns, global, proc charges, every effect with its aura, points, misc value and class mask) for each spell named below.
2. **One-iteration combat logs.** A scratch test (never committed) logged every opener from the first pre-pull action to +4 s, for the 14 caster and healer specs at raid 60 and bare 20, 40 and 60, Horde (56 logs): start, completion, landing, mana spent and the pool at the pull.
3. **The core cast-timing fix isolated.** A second engine with only E 1e7e01e4b's `gcd.go` hunk reverted, against the lane engine. Both ran the 14 specs at bare 20 to 60 and raid 60, both factions, 3,000 iterations (168 pairs).
4. **Rank rows, with and without.** Each round-12 row was removed in its own build of spec-breakdown (the embedded rank-choices.json, restored after each build). Each build ran at 10,000 iterations, seeds 1 and 50001, on the factions each row keys. Healer rows ran on both the published entries' builds and the ladder builds the healprofile lane adopted them on (r12 healprofile/ladders-r12.json).
5. **The round-12 fork tests** for Evocation, Inner Fire, Lightwell, Vampiric Embrace, Mental Agility, item haste, consumables, Blood Fury, Command and bolt travel pass on the lane engine (`-tags=with_db`). The site's heal-profile tests pass (`go test ./request -run 'HealBand|HealProfile|CuratedHeal'`).

## Heal profile split (round-12 healprofile H1)

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| H1 | **The stated rule is what the code applies.** Bands 20, 30, 40 and 50 read their row's `tank_share` (0.679, 0.687, 0.677, 0.677). That share is each band's logged median rounded to the thousandth (30's median is 0.6873). Band 40 (6 fights in SM Armory) and band 50 (none) pool all 990 beta trash fights; bands 20 (218) and 30 (182) use their own dungeons. Level 60 is the reference band: it reads the profile's own 0.659 and states no logs. Loading refuses any row that breaks these (`validateSplit`), and every band's stated tank hit and pulse must equal what the rule derives (`validateBands`). | `sim/request/healprofile_bands.go` figuresFor, validateSplit; data/curated/heal-profile.json bands.rows | Aligned | - | - |
| H2 | **Every healer and every path reads its own band.** `Attach` and `AttachWeights` resolve the profile with `ForLevel` at the healer's level. Below 20 the band is 20; 60 and above take the reference. Every caller goes through these two: spec-breakdown, the in-process healing run, role metrics (`applyRoleFight`), dps-per-point, and the leveling-bis healer scorer. No path reads the top-level 60 figures at a lower band. | `rg Attach(` over sim/ | Aligned | - | - |
| H3 | Level 60 keeps 0.659 under ruling 6. The healers lane is deriving a provisional level-60 split from Era raid logs (coordinator, 2026-11-03), so this lane checked only the rule below 60. | - | Noted (healers lane owns 60) | - | - |

## Vampiric Embrace, Blood Fury, Command (round-12 healers A1, auras A2/A4)

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| V1 | **Vampiric Embrace matches the client.**<br>- 15286: one effect, aura 4, 20 points. Duration 30 s, 30 s cooldown, 1.5 s global, 40 mana.<br>- Its proc mask (0x800A0000) fires on harmful spell and periodic damage taken.<br>- Server text: "all party members ... healed for 20% of any Shadow spell damage you deal for 30 sec".<br>- The engine heals the priest's party, the priest included, for 20% of the priest's own landed Shadow damage, off both hits and ticks. Each tick fires one callback (`spell_result.go:506/509` route periodic and direct damage separately), so nothing heals twice.<br>- Another player's or a pet's damage heals nobody.<br>- No published rotation casts it (Shadow's list has no 15286), and Holy and Discipline cannot take it, so no number moves. | client rows; tooltip 15286 (2026-11-03); `priest/vampiric_embrace.go` | Aligned | - | - |
| V2 | **Blood Fury matches the client.** 20572 lasts 15 s on a 2 min cooldown, with no global (StartRecoveryTime 0). Its three effects are aura 166, 167 and 317 at 10 (attack power, ranged attack power, spell power), with no school or healing qualifier. The tooltip reads "Attack Power and Spell Power by 10% for 15 sec". The engine multiplies healing power with the rest. Items carry spell power and healing power as separate stats (`client_sets.go:475-485`), so the 10% is counted once. **No published healer is an orc** (bis races: Holy gnome/troll, Discipline and Holy paladin human/undead, druid night elf/tauren, shaman dwarf/tauren), so no healer number moves. The orc Elemental, Arcane and Fire entries use it at 0, after the pre-pull cast has completed (logs). | client rows; tooltip 20572; `core/racial_blood_fury.go` | Aligned | - | - |
| V3 | **Command matches 70753.** 20575 e0 is aura 107, +5, misc 8, mask bit 26 (the demon passive). e1 is aura 108, +5, misc 0, mask 0x3000. Every Firebolt and Lash of Pain id the engine registers (`imp.go:101`, `succubus.go:117`) is in `commandDemonSpellIDs`. Guardians are skipped. No published warlock is an orc (troll and gnome), so no number moves. | client rows; `core/racial_command.go` | Aligned | - | - |

## Missile speeds (round-12 warlock A1/X1; casters' list)

Client SpellMisc.Speed against the engine's MissileSpeed. All checked at the top rank and the ranks the rank rows cast:

| spell | client | engine |
|---|---|---|
| Shadow Bolt 686-25307 | 20 | 20 |
| Incinerate 412758, 1293812-13 | 20 | 20 |
| Soul Fire 17924 | 24 | 24 |
| Death Coil 17926 | 24 | 24 |
| Frostbolt 10181, 25304 | 28 | 28 |
| Fireball 10151, 25306 | 24 | 24 |
| Pyroblast 12526, 18809 | 24 | 24 |
| Arcane Missiles: channel 10212 / missile 10273 | 0 / 20 | missile 20 |
| Arcane Wrath 1326082 | 24 | 24 |
| Ice Lance 1240047 | 38 | 38 |
| Wrath 9912 | 20 | 20 |
| Lightning Bolt 15207, 15208 | 20 | 20 |
| Lava Burst 1326047 | 20 | 20 |
| Starfire, Chain Lightning, Mind Blast, Smite, Holy Fire, Searing Pain, Immolate, Scorch, Fire Blast | 0 | none |
| **Frostfire Bolt 401502, 1237312, 1237313** | **24** | **28** |
| Imp Firebolt 3110-11763 | 16 | none |

Mind Blast is not instant in Forever: 10947 is a 1.5 s cast with no travel, and the engine casts it in 1.5 s (logs).

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| T1 | **Frostfire Bolt flew at Frostbolt's 28 yards a second. The client says 24.** All three learnable ranks (trained at 40, 50 and 60) have SpellMisc.Speed 24, the same as Fireball and Pyroblast. The engine set 28 when the spell was rebuilt from the client (b039c21fb, 2026-10-07). From 30 yards a bolt lands 1.25 s after the cast, not 1.07 s. **No published rotation casts Frostfire Bolt**, so no number moves; a /sim custom rotation would. The test `TestFrostfireBoltTravelsAtTheClientSpeed` fails on int/r13 (ranks 2 and 3 at 28, want 24) and passes with the one-line fix. The mage suite passes. Patch for the casters lane, who own `sim/mage`: `$R/regcast/patches/frostfire-bolt-speed.patch`. | SpellMisc rows 401502, 1237312, 1237313; trainables/mage.json (Frostfire Bolt 40/50/60); `mage/frostfire_bolt.go:67` | CONFIRMED-BUG (latent) | old (b039c21fb; round 12's missile-speed fixes and casters M3 checked only the spells the APLs cast) | - (patch routed) |
| T2 | **The Imp's Firebolt has no travel; the client gives it 16.** The fix would move nothing: pets stand at 0 yards in the engine (`TravelTime` reads `Unit.DistanceFromTarget`, which no pet sets), and every warlock spec sims with the Succubus. The Demonology Imp is sacrificed before the pull. | SpellMisc 3110-11763 (Speed 16, global 1.0 s, engine global 1.0 s); `warlock/imp.go`; `request.go` warlockOptions | OPEN (latent; needs pet distance first) | old | - |

## The core cast-timing fix, Evocation, Inner Fire (round-12 casters H1-H3, regcast P1)

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| C1 | **The global fix (E 1e7e01e4b, `gcd.go`) now moves no caster or healer at any band.** With it reverted, all 168 pairs were byte-identical: the 14 specs, bare 20 to 60 and raid 60, both factions. Its only trigger in this group was the healing priest's Inner Fire spending the global as Lightwell ended. Since E 1d693e0da, Inner Fire goes up with no global, so the case is gone. | gcd-scan (lane dir); logs | Aligned (no side effect) | - | - |
| C2 | **The fix is safe for instant procs and channels.** It runs only when a hardcast that has not completed ends between now and the new global. A completed cast resets `Hardcast.Expires` before its OnComplete runs, so the next cast never re-queues it. Channels use ChanneledDot, not Hardcast. Off-global instants (Presence of Mind, Arcane Power, Inner Focus, Berserking, Blood Fury) never call SetGCDTimer. A pre-pull cast that completes at 0 is completed by the global's own action before the rotation runs. In the Arcane log, the pre-pull Frostbolt completes at 0.00, then Presence of Mind, then the instant Frostbolt, which pays Arcane Power's 30% (299 mana). Clearcasting and Nightfall change cost and cast time at completion only. Wands (Shoot 5019, Speed 20) are auto-attacks. | `core/cast.go:288-344`, `gcd.go`, `character.go:436`; logs | Aligned | - | - |
| C3 | **Evocation (E 7ccb794fd):** 12051 is an 8 s channel (aura 110 +1500, 8 min cooldown). The regen aura now ends at low priority just after the channel. An Evocation clipped early also ends its regen at once. The test passes. | SpellMisc/SpellCooldowns 12051; `mage/evocation.go` | Aligned | - | - |
| C4 | **Inner Fire before the pull** goes up at -1 s with no mana, global or cast slot. The client charges 315 (283 with Mental Agility), but a ten-minute armor is up long before the pull and refills before it. In Forever it is armor only: 10952 is one effect, 1395 armor, 20 charges taken by melee hits. So it gives no healing and no number depends on it beyond the opener timing. Holy and Discipline at raid 60: Lightwell 27871 from -2.5 completes at -1.0, Inner Fire at -1.0, the first heal at 0. At bare 40 the opener casts Lightwell rank 1 (724) and Inner Fire rank 1006; at bare 20 it casts Inner Fire 7128 with no Lightwell (the talent is not taken yet). Holy bare 60 casts Lightwell rank 1, its own Horde row. | tooltip 10952; logs | Aligned | - | - |

## Healer and caster rows (round-12 healers, healprofile, ranks)

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| R1 | **Mental Agility on Power Infusion:** 14520's mask names 0x80000000, and 10060 (15 s, 3 min, no global; Forever's is aura 136/79 at 20) now carries the engine mask. It is out of every default. | client rows; `priest/priest.go` | Aligned | - | - |
| R2 | **Idol of Swiftness:** 1306488 is aura 107, op 11, -3000 on Swiftmend's mask. With `-gear ranged:279250` at raid 60 Horde, the druid heals 833.7 → 845.6 HPS (1,000 iterations; round 12 measured +1.2 to +1.3%). It sits in no published set until the refresh (MR17). | ItemXItemEffect 279250; spec-breakdown | Aligned | - | - |
| R3 | **Reflection** reads 17/33/50 from the client curve. | `druid/talents_restoration.go` | Aligned | - | - |
| R4 | **The healer rows still win beyond error on the ladder builds they were adopted on** (10,000 iterations, seeds 1 / 50001, row against no row):<br>- Discipline 40-49 (Renew 1, Binding Heal 1, Shield 6, Prayer of Healing 1): Horde +1.42 / +1.42%, Alliance +1.31 / +1.38%.<br>- Discipline 50-59 Binding Heal 1: Horde +0.36 / +0.41%, Alliance +0.55 / +0.55%.<br>- Holy paladin 40 Flash of Light pair: Horde +3.07 / +3.07%, Alliance +2.77 / +2.77%.<br>- Discipline bare-60 Binding Heal 1, Alliance (published Human): +0.54 / +0.51%.<br>The dropped Shield rows stay dropped: Shield 2 at 20 loses 1.4 to 1.5%; Shield 5 at 30 sits at -0.23 to 0.00%. | rows.tsv and the ladder pairs (lane dir) | Aligned | - | - |
| R5 | **On the stale published builds, Discipline 30 would want Shield 5 back (+0.68 to +0.72%, both factions and seeds).** The published 30 entry predates the ladder build: it measures 73.5 HPS against the ladder's 87.6. At the refresh the BiS file takes the ladder build, where the row is neutral (R4). This is not a row error. | as R4 | Aligned (MR17) | - | - |
| R6 | **Restoration shaman band 20's Water Shield build** (5320000010) beats the round-11 build (5420000000) by +15.89% Horde and +12.41% Alliance at both seeds. It casts Water Shield at -3 s (no cost, the client has no power row) and refreshes it when down. The spec-breakdown default still reads the stale published 20 entry, which lacks the talent (MR17). | spec-breakdown `-talents`; log | Aligned | - | - |
| R7 | **The caster rank rows still win beyond error** (published entries, row against no row, seeds 1 / 50001):<br>- Elemental Lightning Bolt 15208 rank 7: bare 50 Horde +0.43 / +0.40%, Alliance +0.51 / +0.51%; bare 60 Horde +0.41 / +0.41%, Alliance +0.40 / +0.40%.<br>- Shadow Mind Blast 50-59: Horde rank 3 +0.40 / +0.37%; Alliance rank 2 +2.44 / +2.50%.<br>- Shadow Devouring Plague rank 2, bare 60, Alliance: +0.77 / +0.72%.<br>- Fire Blast rank 3 at 30, Horde: +0.76 / +0.77% (round 12 had +1.7%; still beyond error).<br>- Destruction Corruption rank 3 at 40, Horde: +0.16 / +0.16%, the same edge-of-error gain round 12 adopted (error 0.13). | rows.tsv | Aligned | - | - |
| R8 | **The other round-12 changes match the client.**<br>- dps-per-point (S 2d9d55d90): band requests name bare and the raid request names raid. `Preset` reaches only the rank choices (`rank_choices.go:131`), so nothing is layered twice.<br>- Consumable amounts: Lesser Mana Potion 438 is 320 ± 12.5% (280 to 360). Minor Recombobulator 18805 is 128 ± 25% (96 to 160). The healthstones 5509, 5510 and 9421 use spells 5720, 5723 and 11732: flat 600, 960 and 1440. The engine has every one.<br>- Item haste counts once through the stat path (HasteRatingPerHastePercent 1); its test passes. | client rows; `core/consumes.go`; `cmd/dps-per-point/main.go:367-388` | Aligned | - | - |

## Fresh re-audit: pre-pull lines and openers

Logs at raid 60 and bare 20/40/60, Horde. Every pre-pull spell's cast time, cost and cooldown was read against SpellMisc, SpellPower and SpellCooldowns.

| # | spec | finding | confirmed? |
|---|---|---|---|
| F1 | Fire | Pyroblast (18809 at 60, 12523 at 40) from -6 s completes at 0 for 440 / 240 mana and lands at +1.25 s (24 yd/s). Bare 20 has no Pyroblast talent and no pre-pull line. | Aligned |
| F2 | Frost | Ice Barrier at -4 s: 13033 costs 480 × 0.85 with Frost Channeling = 408; 11426 at 40 costs 259. A 30 s category, a 60 s shield, the 5% / 5% modifiers. Frostbolt from -2.5 s completes at 0. Cold Snap (12472, no global) fires at 0 as a major cooldown and resets nothing that is on cooldown. It does no harm, and held for Ice Lance's 8 s category it is worth at most one cast. | Aligned |
| F3 | Arcane | The pre-pull Frostbolt completes at 0 at raid and bare 60 (Improved Frostbolt). At bare 20 (rank 4, 2.6 s) it spills to +0.1 s; at bare 40 (3.0 s) to +0.5 s. This is round 12's M4 note, still unchanged; it is not a correctness bug. A miss on the pre-pull bolt shows as damage 0 in the log, as it should. | Aligned (timing note, casters) |
| F4 | Shadow | Shadowform at -20 s (550 at 60, 40% of base mana) is regenerated to full by the pull (6013 of 6013). Bare 20 has no Shadowform yet. | Aligned |
| F5 | Elemental | Lightning Bolt from -2.5 s: 2.0 s with Alacrity (rank 9 at raid 60, rank 4 at bare under its row), completing at -0.5 s and landing at +1.0 s. Mana Spring at -4 s at bare bands (10495 at 40, 10497 at 60). Elemental Mastery at 0 (15 s, 3 min, no global; -0.5 s on Lightning Bolt and Lava Burst) followed by Flame Shock: Flame Shock is outside its masks, so the next Lava Burst still takes the cut (1.5 s cast in the log). The half-second idle before the pull is round 12's note. | Aligned |
| F6 | Balance | No pre-pull lines; the druid starts in Moonkin Form (`balance.go:78`). Wrath costs 30 = 120 × (1 - 0.50 Improved Wrath - 0.25 Moonglow): the two aura-108 cost rows add, as the client's same-op modifiers do. Nature's Grace (16886) is 10% cast speed plus -10% global (misc 21) for 3 s; the engine models both. Clearcasting gained in the middle of a cast is spent by that cast at completion. | Aligned |
| F7 | Affliction | Amplify Curse at -5 s (30 s, off the global, 3 min). | Aligned |
| F8 | Destruction | Immolate from -1.5 s (1.5 s with Bane) lands at 0, with the band's rank (11668, 11665, 1094). | Aligned |
| F9 | Demonology | Imp at -60 s, sacrifice at -45 s, Succubus at -44 s (6 s with Master Summoner; 10 s in the client and at bare 20), Soul Link at -2 s. **The warlock starts the pull 8 to 12% under full mana** (raid 60 about 5,630 of 6,152; bare 40 about 2,614 of 2,982). The summons fit into 60 s, and Soul Link at -2 s is paid with no time to regenerate. A player summons earlier and tops up. Moving Soul Link to -30 s: raid 60 +0.24 / +0.14%; bare 20, 40 and 60 within ±0.1%. Round 12's W2 said "regenerated to full"; that holds for the other two warlocks only. | OPEN (play assumption, below 0.25%; warlock lane) |
| F10 | Holy paladin | Oath of Light (1328548, 8% base mana, no cast time, 1.5 s global) at -1.5 s; the global ends at the pull. | Aligned |
| F11 | Restoration shaman | Water Shield at -3 s, then the band's Healing Stream at -1.5 s (1 s totem global). | Aligned |
| F12 | Restoration druid | No pre-pull lines; Regrowth at 0. | Aligned (pre-HoTs are a rotation choice) |
| F13 | Holy and Discipline | Inner Focus at 0 on a free Prayer of Mending (Holy) or Renew (Discipline). | Aligned |

## Before/after

No number moves. The lane has no commits that change behaviour. Spot-checked against $R/before.tsv (raid 60, Horde, seed 1, 10,000 iterations): mage-fire 673.0, priest-shadow 720.0, warlock-destruction 636.2, shaman-elemental 479.6, all equal. The routed Frostfire Bolt patch moves no published number (no rotation casts the spell).

## Files to regenerate

None.

## Measure at refresh

- Nothing new. MR17 already covers Discipline 30's stale published build (R5) and Restoration shaman 20's stale entry (R6).

## For other lanes

- **Casters (sim/mage):** apply `/private/tmp/claude-501/-Users-jh-code-forever/54c69daf-a689-41fc-ad89-a21df2102a2f/scratchpad/r13/regcast/patches/frostfire-bolt-speed.patch` (T1). It makes `MissileSpeed` 24 in `frostfire_bolt.go` and adds `frostfire_bolt_speed_test.go`. Run `go test -tags=with_db ./sim/mage`; no golden moves. Arcane's pre-pull Frostbolt still spills past the pull at bare 20 and 40 (F3).
- **Warlock:** Demonology's pull mana (F9): Soul Link earlier (about -30 s) is the player's line. The Imp's Firebolt speed 16 (T2) waits on pet distance.
- **Coordinator:** nothing of mine to merge besides this report.

## Still open

- T2 (pet distance), F9 (Demonology pre-pull mana).
- Carried, with no new evidence: Prayer of Mending hop travel (Speed 40 on 1240827), Elemental Mastery's charge, Ice Barrier's never-broken shield, Inner Focus on Power Infusion.
