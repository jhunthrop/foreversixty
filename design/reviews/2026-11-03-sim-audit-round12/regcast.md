# Sim audit round 12: regression lane, casters and healers (regcast)

Scope: everything the combined landing (fork 47df4a89d..2606bc05e, site 8d0f78ea0..c6bdc855c) and round 11 changed for mages, priests, shamans, warlocks, druid casters and healers and the Holy paladin; then a fresh re-audit of every pre-pull line in those APLs against the client.

Engine branch `lane/r12-regcast` (from 2606bc05e). Site branch `wt/r12-regcast` (from c6bdc855c). Client 1.60.1.70753 (data/builds/1.60.1.70753 spellconst, trainables, spellranks, talents; round 11's wago copies in scratchpad r11/c70753 for SpellMisc and SpellCooldowns).

**Result: 1 confirmed bug, fixed (old). No landing or round-11 change caused a bug in this group.**

## How it was checked

1. Every caster and healer rotation diffed across the landing as structure (notes stripped): the curated file, the sim/request copy and the fork's ui copy are identical for all 15 specs.
2. One-iteration combat logs (a scratch test, never committed) of every opener at raid 60 and bare 20/40/60: when each pre-pull cast starts, completes and lands, what it costs, and what the rotation does at 0.
3. Client rows for every pre-pull spell: cast time, cost, cooldown and category (SpellCooldowns), missile speed (SpellMisc), talent text.
4. The rank rewrite run over every caster and healer APL at levels 10 to 60, both presets (scratch test): every rewritten id belongs to the authored family's client chain (Arcane Blast, Ice Lance, Penance, Binding Heal, Dark Sacrifice ranks); no Season of Discovery copy leaks in.
5. Every rank-choice row checked against spellranks: family, rank learned by `from_level`, no book rank, no overlap. The curated heal-profile, rank-choices and ladder-first files equal their embedded copies.

## Mage

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| M1 | Fire's pre-pull Pyroblast is legal. Rank 8 (18809) casts 6.0 s from -6 s, completes at 0, the missile (speed 24) lands at +1.25 s, and its 12 s DoT ticks from +4.25 s. Mana (440) is spent at completion, from a full pool; nothing lands before the pull. | spellconst 18809 (cast 6000, 440 mana, DoT 53 every 3 s); SpellMisc 18809 Speed 24; log | Aligned | - | - |
| M2 | Frost's Ice Barrier at -4 s and Frostbolt at -2.5 s are legal. Ice Barrier takes its 1.5 s global to -2.5 s, the 2.5 s Frostbolt (Improved Frostbolt 5) completes at 0 and lands at +1.07 s (speed 28). The 15% Frost damage for a minute plus 5% while shielded match the talent text and the pushed rows (11426: aura 108 at 5 and 5; Frigid Power 1326891 at 15 and 15); 30 s cooldown (SpellCooldowns 54730). The shield is assumed never to break (nothing hits a simmed mage; documented in ice_barrier.go). | talents/mage.json Ice Barrier; spellconst 11426, 13033, 1326891; log | Aligned | - | - |
| M3 | Arcane's pre-pull Frostbolt and Presence of Mind Frostbolt are legal at raid 60: the pre-pull Frostbolt (2.5 s with the build's Improved Frostbolt) completes at 0; at 0 Arcane Power and Presence of Mind (off the global; 180 s category, SpellCooldowns 54721) make the next Frostbolt instant ("your next Mage spell with a casting time less than 10 sec"). 2606bc05e's `spellIsKnown` gate keeps bands without the talent off the line. Mana and cooldowns are not counted twice: the pre-pull cast pays 221 at completion, the PoM cast pays 299 under Arcane Power's +30%. | talents/mage.json Presence of Mind; log at raid 60, bare 20 and 40 | Aligned | - | - |
| M4 | Side note: at bare 40 Arcane's build has no Improved Frostbolt, so the -2.5 s Frostbolt takes 3.0 s and completes at +0.5 s, holding the opening cooldowns half a second. This is legal; the line just starts late for that build. | log, bare 40 | Aligned (APL timing, see For other lanes) | - | - |
| M5 | Fresh: Pyroblast's 1.5 s cast after Combustion is Heating Up at three stacks (-25% a stack), not an error. Hot Streak's aura id carries Forever's Heating Up. | talents/mage.json Heating Up; log | Aligned | - | - |
| M6 | Round 11's mage fixes re-read for side effects: Clearcasting -1000% reaches damage spells only and moves neither Fire nor Frost; Improved Scorch only on a hit; Arcane Missiles rank 1; the Robe of the Archmage's shared minute. The rank rewrite never picks the book ranks (Frostbolt 25304, Arcane Missiles 25345). | round-11 casters.md M1-M3, M11; spellranks book flags | Aligned | - | - |

## Shaman

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| S1 | Elemental's pre-pull Lightning Bolt is legal: rank 9 casts 2.0 s from -2.5 s, completes at -0.5 s and lands at +1.0 s (client Speed 20 at the caster's 30 yd). Nothing lands before the pull. The half second before the pull is idle. | SpellMisc 15207 Speed 20; log | Aligned | - | - |
| S2 | Elemental's pre-pull Mana Spring (-4 s, 1 s totem global) is cast at every bare band and skipped under the raid preset, where round 11's E1 makes it uncastable: the raid's improved Mana Spring (31.25 mp5) beats the shaman's own 25. | log raid and bare 60; water_totems.go | Aligned | - | - |
| S3 | Elemental's Fire Nova only at two or more targets: right. Fire Nova hits "within 10 yd of your active Fire totem", and a ranged Elemental's totem stands at its feet, so the old single-target Clearcasting line hit a boss 30 yd away. Enhancement's Clearcasting Fire Nova is legal in melee. Elemental Focus procs on 10% of casts of any Fire, Frost or Nature damage spell and frees the next one (talent text); Fire Nova's blast does not proc it a second time. | fire_nova.go; talents/shaman.json Elemental Focus | Aligned | - | - |
| S4 | Lightning Bolt rank 10 while mana is ahead of the clock, rank 9 (or the band's row) after: rank 10 (15208) trains at 56 (no book flag). Below 60 the rank-10 line rewrites to the band's top rank and the rank-9 line to the row's rank 4; that is the pacing the guide lane measured, not a row conflict. | spellranks; rank-choices.json | Aligned | - | - |
| S5 | Enhancement's talent swap: Shamanistic Focus is -45% on Shocks and Lightning Shield (aura 108, misc 14, -45), stacking additively with Convection; Earth's Grasp and Guardian Totems do nothing in a DPS sim. Pre-pull Strength of Earth (-3 s) and Grace of Air (-1.5 s) are 1 s totem globals and end before the pull. | spellconst 1223030; talents/shaman.json | Aligned | - | - |
| S6 | The Stormstrike spellranks fix moves no sim number: the rewrite never tracked 17364, and every caster chain it touched (Water Shield, Grounding Totem, the early-copy drops) leaves the rewritten ids unchanged. | scratch rewrite scan; commit 270943fe0 | Aligned | - | - |
| S7 | Restoration shaman's pre-pull Water Shield (-3 s) and Healing Stream rank 5 (-1.5 s) end before the pull. | log | Aligned | - | - |

## Priest

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| P1 | **The healing priest cast Inner Fire in the middle of its pre-pull Lightwell.** The engine's Inner Fire option was a cast with a global at -1 s. Holy's and Discipline's rotations cast Lightwell (1.5 s) at -2.5 s, which finishes at that same instant. The engine let Inner Fire start first (the cast had not been completed yet), so Lightwell, whose completion rides on the global, finished only at +0.5 s, and the fight's first heal waited for Inner Fire's global at every band. In game you cannot start a spell while another is casting, and a priest has Inner Fire (ten minutes) up long before the pull, as every armor is. Inner Fire now goes up at -1 s with no cast slot or global. | combat log: `[-2.50] Casting 27871`, `[-1.00] Casting 10952`, `[0.50] Completed cast 27871`; `sim/priest/healing/healing_priest.go` castInnerFireBeforeThePull; Lightwell 27871 cast 1500 ms (spellconst). `TestInnerFireDoesNotCutIntoThePrepullCasts` failed first (Lightwell completed at 0, want -1 s; global running to 0.5 s at the pull) and passes. | CONFIRMED-BUG | old (Inner Fire pre-pull E a6bd9de7e, 2026-10-07; Lightwell pre-pull S b70e67d2b / E e0aff7e5a, 2026-10-08; no round looked at the healer openers' timing) | E 1d693e0da |
| P2 | Shadowform at -20 s: legal and better. Shadowform costs 40% of base mana (550 at 60) and the 20 s before the pull regenerate it, so the priest enters the fight full in form, as players do. At -1 s it started the fight 550 short. All five Shadowform effects are modelled (10% Shadow damage, -15% physical taken, -50% Shadow cost, +100% crit bonus). | spellconst 15473 (cost_pct 40); log: 6013 at -20 s, full at 0 | Aligned | - | - |
| P3 | Shadow's Dark Sacrifice line (round 11 S1): Undead rank 5 (1277328), 320 health to 320 mana a 3 s tick plus Spirit, 10 min cooldown; paced so raid 60 stays within error. Round 11's Devouring Plague row interaction was fixed by the ranks lane (S 5766c8ce7). | spellconst 1277324-1277328; mana_abilities.go:74 | Aligned | - | - |
| P4 | Discipline's party Binding Heals ahead of party Flash Heals: a reorder of lines that fire below 30% party health; raid 60 moves 0.0%. Binding Heal ranks 1-6 match the client (rank 6: 825 twice, 0.429, 380 mana, 1.5 s; rank 1 at 25, so no sub-20 penalty applies to the bare-60 rank-1 row). | binding_heal.go; spellconst 1240774, 401937 | Aligned | - | - |
| P5 | Round 11's sub-20 heal coefficient (E 074edfadf): the healing proc mask exempts heals only. The damage spells that carry it (Holy Nova's mixed mask, Drain Life's and Death Coil's heal halves) are learned at 20 or later or are passive, so no damage rank lost its penalty by accident. Lesser Heal rank 3 (2053: 163, 0.714, 75 mana, level 10) keeps 0.714. | spell.go:803; the mask grep; spellconst 2053 | Aligned | - | - |
| P6 | Holy and Discipline's pre-pull Lightwell rank 3 (27871: 1.5 s, 365 mana, 3 min, 10 min category) is legal once P1 is fixed: it lands at -1 s and lasts to 179 s. | spellconst; SpellCooldowns | Aligned after P1 | - | - |

## Warlock

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| W1 | The +11 lines on each warlock APL are round 11's own Curse of the Elements line (S ac4937071, merged in the landing). At raid 60 it never casts (the raid's curse is on the target); at bare 60 Affliction casts it 1.12 times and Destruction 1.13. Amplify Curse is consumed only by Bane of Agony (or Weakness/Exhaustion), never by the curse placed ahead of it. | breakdowns; curses.go; talents/warlock.json Amplify Curse | Aligned | - | - |
| W2 | Pre-pull lines: Affliction's Amplify Curse at -5 s (30 s, off the global, 180 s cooldown); Destruction's Immolate rank 7 (11668, the top trainer rank; 25309 is a book) at -1.5 s with Bane's 1.5 s cast, landing at 0; Demonology's Imp, Demonic Sacrifice, Succubus and Soul Link from -60 s, the sacrifice kept through the new summon by Demonic Pact. Each ends before the pull with mana regenerated to full. | spellconst 18288, 11668; SpellCooldowns 55294; log | Aligned | - | - |

## Druid casters and healers, Holy paladin

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| D1 | Balance and Restoration druid have no pre-pull lines and the landing did not touch them. | APL diff | Aligned (nothing to audit) | - | - |
| D2 | Holy paladin's pre-pull Oath of Light (1328548, 8% base mana, 5 min, instant) on the tank at -1.5 s ends its global at 0. Greater Blessing of Light, removed from 70753, is gone from the APL. | spellconst 1328548; log | Aligned | - | - |
| D3 | Round 11's crit lane on this group: Conviction (aura 290) is paladin-only and the Holy paladin's heals keep the same 6%; Atiesh, Blackhand's Breadth and Eye of the Beast are in no published caster or healer set. | round-11 crit.md | Aligned | - | - |

## Cross-cutting

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | Pre-pull timing in the engine: a pre-pull hardcast that finishes at a pull-relative instant completes then, and its damage is computed at completion; every landing caster pre-pull lands at or after 0. Pre-pull mana spends are real and regeneration before the pull is modelled. The only engine-registered pre-pull action was the healing priest's Inner Fire (P1). | log; core/apl.go, environment.go RegisterPrepullAction | Aligned | - | - |
| X2 | Rank rewrite and rank rows: every rewritten id at levels 10 to 60 in both presets belongs to the authored client chain; every row is learned by its `from_level`, no row names a book rank, no rows overlap. | scratch scan; spellranks | Aligned | - | - |
| X3 | guide-cooldowns.json (new in the landing) states its rule as "baseline class trainer cooldowns, three minutes or longer, no talent needed", but leaves out Inferno (1122, 1 h, trained at 50) for the warlocks and Reincarnation (20608, 1 h, trained at 30) for the shamans. Ice Block is rightly out: it is a Frost talent in Forever. Guide copy only, no sim number. | trainables/warlock.json, shaman.json | OPEN (guides) | - | - |

## Before/after

Engine lane/r12-regcast (1d693e0da) with site wt/r12-regcast, against $R/sb-before from $R/site-before/sim. Effective HPS from `-heal`, 10,000 iterations.

| spec, raid 60 | horde s1 | horde s50001 | alliance s1 | alliance s50001 | horde s1 absolute |
|---|---|---|---|---|---|
| priest-holy | +0.32% | +0.40% | +0.23% | +0.20% | 696.7 -> 698.9 |
| priest-discipline | +0.37% | +0.41% | +0.40% | +0.38% | 653.5 -> 655.9 |

Bare bands, horde, seed 1 / seed 50001:

| spec | 20 | 30 | 40 | 50 | 60 |
|---|---|---|---|---|---|
| priest-holy | -0.29% / -0.29% (34.8 -> 34.7) | 0.00% / -0.14% | +0.14% / +0.14% | +1.48% / +1.43% | +1.60% / +1.60% |
| priest-discipline | +0.61% / +0.61% | 0.00% / 0.00% | +0.69% / +0.76% | +0.60% / +0.65% | +0.69% / +0.79% |

Holy bare 20 is 0.1 HPS lower: the priest now starts healing at 0 rather than 0.5 s, and that mana-bound band runs dry 0.6 s sooner (mana lasts 154.7 s against 155.3 s). It has the same mana: before, the 65-mana Inner Fire was regenerated before the pull. Holy at 50 and 60 gains most because Lightwell now lands before the pull. No DPS spec moves; the change is in the healing priest package only.

## Files to regenerate

None. No fork golden moved (the `-tags=with_db` priest suites pass unchanged). The site's ladder goldens for priest-holy and priest-discipline (sim/request TestRotationLadder) will move at the pin, with the healer numbers above.

## Measure at refresh

- Holy and Discipline priest picks at every band once 1d693e0da is pinned. Holy at bare 50 and 60 moves by about 1.5%, but that moves every set the same way; pick swaps are unlikely.

## For other lanes

- **Casters (mage APL):** Arcane's pre-pull Frostbolt is written for a 2.5 s cast. At bands whose build lacks Improved Frostbolt it casts 3.0 s and spills 0.5 s past the pull (bare 40). A -3 s line, or one keyed to the cast time, would start it on time. Elemental's -2.5 s Lightning Bolt (2.0 s at 60) leaves 0.5 s idle before the pull; -2.0 s would be exact. Neither is a correctness bug.
- **Guides / class lanes:** X3, Inferno and Reincarnation in guide-cooldowns.json.
- **Integrator:** E 1d693e0da moves priest-holy and priest-discipline only (tables above). No commit here moves a spec outside this group.

## Still open

- Ice Barrier's shielded 5% assumes the shield never breaks; raid-wide damage on a real boss would end it early. A log of shield uptime would settle it.
- Round 11's heal coefficient fix at level 40 and above still rests on logs from levels 27 to 35 (unchanged from round 11's open list).
