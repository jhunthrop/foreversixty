# Sim-accuracy round 23: warlock (2026-11-05)

Specs: Affliction, Demonology, Destruction. Client 1.60.1.70790. Site `wt/r23-warlock` and fork `lane/r23-warlock` from int/r23 (platform, druid and hunter merged). Baseline: spec-breakdown built from int/r23 before any change (`$R/warlock/sb-int`, `out/base60.tsv`). Raid 60 runs use `-bis-dir $R/bis-frozen` and the guide raid build (`$R/rt`).

**Totals: 2 CONFIRMED-BUGs (both `old`), fixed with failing tests first. Raid 60 moves Demonology +0.3 to +0.7% and Destruction +0.1 to +0.4%; Affliction (no demon) is unchanged.**

## Seeds

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| S1 | **Demons took none of the warlock's crit; in Forever they take the crit on his gear.** <br>**What was wrong:** the engine gave each demon its own measured crit (Succubus melee 0.7%, Lash of Pain 6.4%, Imp Firebolt 2.2%) and nothing from the warlock. A level-60 raid warlock wears 5 to 8% crit (frozen BiS: Affliction and Demonology 6% horde and 8% alliance, Destruction 5% and 7%), and none of it reached the Succubus. <br>**What Forever does:** Blizzard (AidanZMoon) on forever-bugs #331, 2026-10-14 (https://github.com/ClassicWoWCommunity/forever-bugs/issues/331): "Warlock pets currently only scale with their master's +% critical strike gear and have erroneous base crit." Blizzard intends the demon to mirror the warlock's spell crit, but "we aren't enabling the scaling until the baseline critical strike chance is removed". #371 (Human and Orc crit racials missing from the pet sheet, 2026-10-16) was closed as #331's duplicate. Blizzard's notes thread (t/2360696, every post through 4 Nov, re-fetched today with no new posts) has no line that enables it. The 16 Oct line "pets were only ever inheriting the player's Melee Critical Strike chance instead of whichever is higher, melee or ranged" sits under Hunter > Pets, and #331's thread says it is the hunter-pet fix ("will use Ranged Crit starting Thursday"). The client gives Warlock Pet Scaling 416189 melee and spell crit channels (effects 13 aura 52, 14 aura 57) with server-computed amounts, so the client cannot settle the amount. <br>**Answer to the seed:** a demon's crit should NOT follow the warlock's full crit yet (unlike the hunter pet's ranged crit). It takes his gear crit only. The raid's 3% druid aura reaches the demon neither directly (SpellMisc Attributes_3 0x100 "only target players", `RefusePlayerBuffs`) nor through the warlock (it is a buff, not gear), so nothing counts twice. <br>**Fix:** `makePet` adds the warlock's `EquipStats()` crit (gear plus the request's bonus stats, which stand for gear, so stat weights see it) to the demon once at the summon; Intellect, racial, raid-buff and proc crit stay out. Test `TestDemonTakesOnlyTheWarlocksGearCrit` (Human warlock with 6% gear crit under both druid auras): red at 0.70/6.40% Succubus and 0.00/2.20% Imp, then green at 6.70/12.40% and 6.00/8.20%. Breakdown check (alliance, 8% gear): Succubus melee crits 4.0% on the boss (0.7 + 8 - 4.8 suppression) and Lash of Pain 13.2%, both 0% and 5.9% before. Destruction keeps the Succubus: with an Imp it is 7% lower (719.8 against 775.8, horde seed 1). The Affliction raid build (Imp sacrificed) still beats the old Succubus build by 1.02 to 1.06% horde and 1.59 to 1.62% alliance after the change. | `sim/warlock/pet.go`, `pet_crit_test.go`; forever-bugs #331, #371; client SpellEffect 416189; `out/crit60.tsv` | CONFIRMED-BUG (single source: Blizzard's own statement on its server; no log can be had) | old (the engine comment read the same reply and left it unwired as "unconfirmed") | E ac1c1762a; guide S aa3506292, S ed1895130 |
| S2 | Round 22's open items: C1 (a row on an unregistered aura runs unconditionally) is platform's and is not reopened here; S8 (Demonic Brand damage, one source) stays OPEN (no log). The Demonology Horde race re-run stays pending owner. | round 22 report | Open / pending | - | - |
| S3 | Guide round 24's optional decomposition (R24-2): Destruction leveling order, bare band gear, horde, 10,000 iterations, seed 1, Conflagrate on cooldown against held, with Shadow and Flame's Shadow buff zeroed (a) or its keep-Immolate chance zeroed (b), via a throwaway env-gated build (not committed). On cooldown against held: 39 +1.11% (a +0.49, b +0.64), 42 +0.76% (a -0.23, b -1.31), 44 +2.89% (a +1.17, b -0.76). The keep-Immolate chance carries most of the on-cooldown lead (it reverses without it from 42), the Shadow buff some of it. The guide's reason holds; no card change. | `out/saf_ab.txt` | Aligned | - | - |
| S4 | Guide round 24's other optional (raid order at 36) was decided in round 22 (S6: -0.3%, not gateable). | | Aligned | - | - |

## Blizzard's Forever notes since round 22, every warlock line

The thread (t/2360696) was re-fetched 2026-11-05: 10 posts, identical to round 22's cached copy (last post 4 Nov 19:35 UTC, UI only). One line round 22 marked "does not touch damage" is still an engine error and is fixed here:

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| N1 | **Demons still took 7.5 health a point of the warlock's Stamina.** Blizzard, 3 Nov (post 9, Warlock): "Pets now gain 2 Health per point of Master's Stamina (was 7.5 Health per point of Master's Stamina)." The engine gave a demon 75% of the warlock's Stamina (7.5 health a point); it now gives a fifth (2 health a point). Player threads from 3 to 5 Nov quote the same line (forums t/2399417, t/2402048: Voidwalker health complaints). No damage path reads a demon's health: DPS is the same to the decimal (Demonology 708.0, Destruction 776.3 horde seed 1, same as before the change). Tests `TestDemonInheritsForeversShares` and `TestSummonedSuccubusFollowsTheWarlocksSheet` were changed to 2 health a point first (red at 300 and 284.25), then green. | `sim/warlock/pet.go:233` | CONFIRMED-BUG (no DPS effect) | old (round 22 read the note, did not change the share) | E 771953ed5 |

Forums: player threads from 3 to 5 Nov ("Warlock Pet Nerf in the Nov 2nd Update?", "Warlock Pets need a Hotfix immediately") have no blue post. The forum API rate-limited the other three threads ("Recent Warlock Nerfs", "Please look into warlock pet HP", "Latest warlock changes"); their titles concern the same health change.

## Affliction

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| A1 | Round 22's raid build re-checked: Burning Shadow 18789 (aura 79, 15, mask 32 Shadow), Shadow Trance 17941 (aura 108 misc 10, -100% cast on mask 0x1), Shadow Vulnerability with no charges (no SpellAuraOptions row), and the Nightfall proc list are all as the client says. After S1 the build still clears the adoption bar against the old Succubus build (S1). Raid 60 unchanged. | client 70790 rows | Aligned | - | - |
| A2 | Wrack's vulnerability covers only Corruption and Bane of Agony (aura 271 on mask 0x402), which matches forever-bugs #446 (3 Nov: the tooltip says "other Shadow damage over time effects" but Wowhead's list excludes Siphon Life and Doom). | `wrack.go:37`; #446 | Aligned | - | - |
| A3 | Mana costs on the card's spells against SpellPower 70790 (Shadow Bolt 370, Immolate 370, Corruption 290, Agony 215, Doom 300, Siphon Life 365, Wrack 200, Curse of the Elements 200, Soul Link 20% base, summons 80/100% base): the breakdown's per-cast spend matches. | SpellPower.csv; breakdown | Aligned | - | - |

## Demonology

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| D1 | S1 is this spec's change (+0.34% horde, +0.66% alliance). | | | | |
| D2 | Raid build decoded again: every point acts, or is the toll to Pact (Improved Health Funnel, Demonic Embrace, Demonic Aegis) or a keep (Suppression). Master Demonologist (Succubus, 23825) is up all fight; Soul Link 25228 is +3% damage for both (aura 79, misc 127) and 30% damage moved. Round 22's Soul Fire 56 to 59 row is unaffected by S1 (demon crit does not touch the warlock's casts). | breakdown; client rows | Aligned | - | - |

## Destruction

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | S1 moves it +0.13% horde, +0.39% alliance (the Succubus). destruction.md:67 "her melee and Lash of Pain out-damage the Imp's Firebolt" still holds (S1). Bane of Havoc stays Incinerate's prerequisite (rounds 14 to 21). | | Aligned | - | - |

## Cross-cutting

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| C1 | Demonic Tactics 1243030 (aura 290, +10% crit while a demon is out) is in the client but has no SkillLineAbility row and no talent: a warlock cannot learn it. Not modelled, correctly. | client SkillLineAbility, SpellName | Aligned | - | - |

## Before/after

Raid 60, guide raid builds, frozen BiS, 10,000 iterations, seeds 1 / 50001, int/r23 baseline against this branch:

| spec | horde | alliance |
|---|---|---|
| Affliction | 745.0 / 744.7 unchanged | 753.7 / 754.0 unchanged |
| Demonology | 705.6 / 705.2 to 708.0 / 707.6 (+0.34 / +0.34%) | 715.6 / 716.6 to 720.4 / 721.3 (+0.67 / +0.66%) |
| Destruction | 774.7 / 775.1 to 775.8 / 776.1 (+0.14 / +0.13%) | 778.3 / 776.8 to 781.3 / 779.8 (+0.39 / +0.39%) |

Bare band (horde, seeds 1 / 50001): Demonology 60 and 40 unchanged; Destruction 60 +0.02%, 40 unchanged (bare gear carries almost no crit).

## Files to regenerate

- None for this lane's own text beyond what the next regen does anyway. The engine change moves Demonology and Destruction raid numbers, so the integrator's regen should refresh the sim/request ladder goldens for warlock-demonology and warlock-destruction rungs whose gear carries crit, and the nightly numbers. No fork golden needed regenerating (the warlock dps goldens are skipped).

## Measure at refresh

- Demonology and Destruction BiS and stat weights: gear crit now also feeds the Succubus, so their crit weight rises a little.

## For other lanes

- **Moves DPS outside warlock: nothing.** Both engine commits are in `sim/warlock/pet.go` only.
- Hunter: confirmed from #331 that the "whichever is higher, melee or ranged" fix of 16 Oct is the hunter pet's; warlock demons are on a different (gear-only) path.

## For the guide lane

- Done here: demonology.md demon paragraph (:64) and stat priority (:98): the demon takes the crit on your gear, not Intellect, racial or raid-buff crit.

## Guide text not done

- None.

## Guide review

- wow-player on the rendered Demonology page (dev build of this branch, 1440 and 390): SHIP with one SHOULD ("Intellect ... pays the same way" could read as reaching the Succubus). Taken (S ed1895130). `_prose.test.ts` passes.

## Logs

Last looked 2026-11-05 12:17 UTC: `api.foreversixty.gg/v1/reports/recent` still lists only z2opt2x2gs5u (2026-10-20); foreverlogs.gg answers 200 on its page (its report API answered "API access not allowed" in round 22); WCL not tried (key is 401, owner). What would settle S1 at 60: any level-60 Demonology or Destruction log with Succubus melee and Lash of Pain crit counts, next to the warlock's gear crit.

## Still open

- S1 is Blizzard's statement only; when Blizzard enables the full mirror (and removes the demons' base crit), the inheritance becomes the warlock's spell crit for both melee and spells (#331).
- Round 22's S8 (Demonic Brand damage) and C1 (platform). Pending owner, not reopened: the Demonology Horde race, level-60 trainer book ranks, Siphon Life tick crit, Heart of the Lion 280/276, warlock Orc.

## Sources

- forever-bugs #331 (2026-10-12 to 10-14), #371 (2026-10-16), #394 (2026-10-20), #446 (2026-11-03), #435: https://github.com/ClassicWoWCommunity/forever-bugs/issues/
- Blizzard Forever notes, https://us.forums.blizzard.com/en/wow/t/2360696 (re-fetched 2026-11-05, posts 1 to 10); forum threads t/2399417 and t/2402048 (3 to 5 Nov, no blue posts).
- Client 1.60.1.70790 tables (SpellEffect, SpellMisc, SpellAuraOptions, SpellPower, SkillLineAbility; `scratchpad/b70790/70790`).
