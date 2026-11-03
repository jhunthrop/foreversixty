# Sim-accuracy round 14: integration (2026-11-03 to 11-04)

Round 14 runs on client 1.60.1.70753 with the 2026-11-03 14:14 UTC hotfix push layered. It starts from the landing of rounds 12 and 13 (site origin/main 72471fcb3, fork pin 5195a8889). The lane reports sit beside this file:
- core, regmelee and regcast: the regression lanes over rounds 12, 13 and the landing;
- warlock, casters, physical, hybrids, tanks and healers: the class re-audits;
- data and platform;
- guides, the guide-prose lane;
- build, a read-only lane added mid-round for client 1.60.1.70755;
- review-physical and review-casters, the two wow-player reviews.

**Not landed.** Following the batch-landing practice, nothing is pushed or pinned. Site `int/r14` and fork `int/r14` wait for the integrator (see "For the integrator").

**Not a zero round.**

## Owner rulings this round, and where they landed

| # | ruling | outcome | where |
|---|---|---|---|
| 1 | Protection paladin band 20: Templar's Bulwark over the second Sacred Duty point | Done. The band-20 build becomes `-15301100`; bands 19 and 25 to 60 are unchanged. Bare 20: chance of death 75.6 to 62.6% (horde), 73.5 to 58.5% (alliance), threat -2.8%. Recorded as a stated exception to the threat metric for band 20 only. The guide's Leveling section now says so (review fix). | tanks.md, S on wt/r14-tanks; S 39136b254 |
| 2 | Same-instant ties break fairly | The engine already settles both exact ties the way the server does: a cooldown that ends on the instant a cast ends is ready (cmangos), and a pulse on a tick keeps the tick in either order. The real bug was nearby: a hit landing a fraction of a millisecond **before** a cast ended pushed it back a full second or cut the last tick. The server handles events in 10 ms batches and runs player casts and ticks before creatures, so such a hit now lands after them (E cbdc60fd5). Discipline's 0% peak falls from about 1% to 0.24-0.45% over +0.01%. The rest is Penance's cooldown ending on the instant a 1.5/2.5 s cast does at exactly 0% haste, which any haste breaks by about 1 ms (OPEN, C4). | core.md C1, C4; healers.md |
| 3 | Restoration druid heal scale: retune by the coverage rule | Done: 1.85 to 2.3, the smallest 0.05 step where every bare set covers under half on both factions. The druid's raid set drops from 95.5% to 83% coverage. The rule's floor ("weakest about 60%") cannot be met at the same time: at 2.3 the paladin and Discipline cover 54-55%. The profile's notes say the "under half" clause decided. The one scale moves every healer (see Before/after). Three rank rows then followed it (healers F1-F3). | healers.md A1, S 78d5b20f7 |
| 4 | The tank boss gets attack power | Done. The engine used a Wrath formula for creature melee (270 AP was under 9% of a swing). Classic's is swing x (1 + AP/630) (Blizzard forums "Boss Attacks and You", 2020-03-03; vmangos `Creature::UpdateDamagePhysical`). The boss takes cmangos classic-db's level-63 warrior-class 268 AP. The mean swing is held, so this alone moves no number. Demoralizing Shout and Roar at the client's 204 cut a swing by 22.7%. The raid debuff options used Era's 146/138 and now use the client's ranks. Protection warrior and bear rotations now keep Shout/Roar up (see owner call O1). | tanks.md, E bfa9c25cd, a3800aac6, 75338e087 |
| 5 | Round 13's in-game checks stay pending | Recorded, not reopened: Arcane Shot's tooltip with spell power (E 05b72a3b5 kept), level-60 trainer book ranks, the six Wowhead-baseline talents, Heart of the Lion 280/276, Rejuvenation rank 1's tooltip. The review adds Renew 1, Heal 1, Lesser Heal 1 and Healing Wave 2 to that tooltip check (see O2). | data.md, healers.md |

The landing items:
- **Weapon DPS in Cat Form.** The engine is right; the landing's workaround was the bug. Blizzard's druid deep dive (forum t/2367361, 2026-09-30) says a form's auto attack now uses the equipped weapon's DPS, and abilities that deal weapon damage use the same values. 3e3a14ff9's `formIgnoresWeapon` is deleted (S ffdb90ce4); the cat's weapon term returns at the dps-per-point regen. hybrids.md.
- **Searing Totem.** The attack count was right (22 at rank 6), but after round 13's 2.437 s cadence the totem counted as gone after its last attack (53.6 s), not at the client's 55 s, so rotations re-dropped it 1.4 s early (E 2fe671caa; R13 5ca93cfa3). core.md C5.
- **The two app fixtures with the 70707 Elemental build:** fixed, with the Fire fixture's old raid build, and a test that catches a stale guide code (S d0391912d). platform.md.
- **Level-60 heal split from Warcraft Logs: BLOCKED.** gcloud's login has expired ("Reauthentication failed", checked 2026-11-03). 60 keeps 0.659. Steps to finish are in healers.md A2: `gcloud auth login`, raidsplit.py, then the heal-profile change that lets the reference band take a logged split.

## Confirmed and fixed, by lane, with origin

Origin key: **R13** caused by a round-13 fix; **R12** caused by a round-12 fix; **landing** caused by the land2 landing itself; **old** predates round 12 and was not found before; **new-data** data changed under us. Each bug is counted once, where it was fixed:
- the Beast Mastery Arcane Shot rows under data (regmelee found them too);
- the Survival Falcon line under regmelee (guides found it too);
- the Shadow coach starter under platform (regcast found it).

| lane | confirmed | R13 | R12 | landing | old | new-data | bugs |
|---|---|---|---|---|---|---|---|
| guides | 39 | 10 | 2 | | 26 | 1 | prose and rotation notes the client contradicts, across all nine classes; site meta-talk in guides and notes (a new `simTalk` guard in _prose.test.ts); Aimed Shot "about twice as hard"; Conviction on every table; missing stat pills; Premeditation without Stealth |
| healers | 7 | 1 | | | 6 | | the heal scale outside its rule (R13 383d137f1); Holy's tank Flash Heal rank 2 with mana unspent; Renew rank 3 at 60; Holy's raid Heal rank 1; Discipline Lesser Heal rank 1; Restoration shaman tank Healing Wave rank 2; Holy paladin Holy Shock rank 4 at bare 60 |
| platform | 6 | 1 | | 1 | 4 | | stale Elemental (R13) and Fire app fixtures; land.sh counting three failed `gh run view` calls as green (landing 2030bef5e); Kindling Stave's proc timer carried across iterations, so results depended on sharding (E b08e7cf68); coach book ranks for Hawk, Serpent Sting, Corruption, Immolate, Battle Shout; the Shadow starter's priority |
| data | 5 | 2 | | | 1 | 2 | Gnome Expansive Mind did nothing for rage/energy; Beast Mastery Arcane Shot downrank rows lose since the push; Devouring Plague rank 2 for Human/Gnome at bare 60 (R13 175d4ff96); Corruption rank 3 at 40 (R13 e88a33c97); Trueshot Aura text 50 to 100 |
| tanks | 4 | | | | 4 | | creature melee's attack power formula; the demoralize debuff options' Era values; Frenzied Regeneration 0.3% instead of 1% per rage; Demoralizing Roar threat 104 instead of 48.4 |
| warlock | 3 | | | | 3 | | Immolate ignored cast speed; Orc beats Troll beyond 1% for Horde Affliction and Destruction; Demonology's leveling Improved Imp 3 (dead with the Sayaad) to Unholy Power 5 |
| core | 2 | 1 | | | 1 | | a hit just before a cast end pushed it back or cut the last tick; Searing Totem's early re-drop (R13) |
| casters | 2 | | | | 2 | | Combustion capped at 20 stacks, client 10; Inner Focus crit on Shadow Word: Death |
| regmelee | 1 | 1 | | | | | Survival's "Falcon gives 10 more attack power" (R13 13497f823) |
| regcast | 1 | 1 | | | | | Holy's Heal rank 1 row at 50-59, which round 13's free Prayer of Mending turned into a -12 to -14% loss (R13 0dcd484d5) |
| hybrids | 1 | | | 1 | | | the cat's weapon term removed by the landing's tool workaround (3e3a14ff9) |
| physical | 1 | | | | 1 | | Premeditation needs no Stealth in Forever; Subtlety now uses it in the fight every 2 minutes |
| build | 0 | | | | | | 70755: every row OPEN until its pin (see below) |
| **total** | **72** | **17** | **2** | **2** | **48** | **3** | |

Bugs by origin, as the owner asked:
- **Caused by round 13: 17.**
  - 10 guide lines, among them Survival's Falcon line;
  - two rank rows (Devouring Plague, Corruption) and Holy's Heal row;
  - the heal scale, which round 13's party Rejuvenation pushed out of its rule;
  - Searing Totem's early re-drop, from round 13's cadence;
  - the stale Elemental app fixture.
- **Caused by round 12: 2,** both guide lines (the Fury stance swap 70aff3eaa and the Shred note 7ea0ba759).
- **Caused by the landing: 2.** The cat's weapon workaround 3e3a14ff9, and land.sh's green-on-error 2030bef5e.
- **Newly found: 51.** That is 48 older, plus 3 from new data: the push's Arcane Shot rows, the Trueshot Aura text, and Conviction in 70753.

Of the 72, 42 are guide, note or coach copy, and 30 are sim, data or tool bugs.

Not counted:
- The review fixes (S 39136b254 and the guides lane's review rows).
- Guides F1 to F3, which followed this round's own healer and tank changes.
- The two healer fork tests the rank syncs broke, caught before merge (healers F6).
- Regmelee's three stale code comments.

**Is the count falling? No: 67 in round 13, 72 in round 14.** But the sim share fell: 30 sim/data/tool bugs against 41 in round 13, and 42 copy bugs against 26. The new meta-talk ban alone produced 12 copy rows.

## Merge notes

| # | conflict | resolution |
|---|---|---|
| M1 | sim/request/rank_choices_remeasured_test.go (data's and regmelee's round-14 rank tests for the same Arcane Shot rows) | kept data's broader test; the rank file auto-merged to the same rows |
| M2 | web/src/content/guides/hunter/survival.md (regmelee's and guides' Falcon sentence) | kept guides' wording; the review then split the sentence |
| M3 | data/curated/apl/warrior-protection.json (tanks' Demoralizing Shout line, guides' Battle Shout note) | took tanks' file and re-applied guides' note; `make apl-check` matches all 56 copies |
| M4 | fork sim/priest/healing and sim/shaman/restoration rotation tests pinned the old rank ids after the healers' syncs | fixed on the healers lane (E b852f9127) |

## Before/after

Before: the landing (site 72471fcb3, fork 5195a8889). After: site int/r14 with fork int/r14, before the last review commits, which are prose only. Raid preset, band 60, 10,000 iterations, both factions. Healers are in effective HPS on the heal profile; tanks in threat per second. The table's % column gives seed 1 / seed 50001.

| Spec | Faction | seed 1 | % (seed 1 / 50001) |
|---|---|---|---|
| druid-balance | alliance | 569.5 to 569.5 | +0.00 / +0.00 |
| druid-balance | horde | 574.9 to 574.9 | +0.00 / +0.00 |
| druid-feral-bear | alliance | 1020.2 to 964.2 | -5.49 / -5.39 |
| druid-feral-bear | horde | 1011.2 to 955.7 | -5.49 / -5.39 |
| druid-feral | alliance | 664.9 to 664.9 | +0.00 / +0.00 |
| druid-feral | horde | 660.4 to 660.4 | +0.00 / +0.00 |
| druid-restoration | alliance | 934.2 to 1006.4 | +7.73 / +7.75 |
| druid-restoration | horde | 933.7 to 1005.3 | +7.67 / +7.70 |
| hunter-beast-mastery | alliance | 659.1 to 659.1 | +0.00 / +0.00 |
| hunter-beast-mastery | horde | 663.5 to 663.5 | +0.00 / +0.00 |
| hunter-marksmanship | alliance | 697.2 to 697.2 | +0.00 / +0.00 |
| hunter-marksmanship | horde | 701.3 to 701.3 | +0.00 / +0.00 |
| hunter-survival | alliance | 741.0 to 741.0 | +0.00 / +0.00 |
| hunter-survival | horde | 744.0 to 744.0 | +0.00 / +0.00 |
| mage-arcane | alliance | 614.2 to 614.2 | +0.00 / +0.00 |
| mage-arcane | horde | 619.2 to 619.2 | +0.00 / -0.02 |
| mage-fire | alliance | 666.4 to 666.4 | +0.00 / +0.00 |
| mage-fire | horde | 673.0 to 673.0 | +0.00 / +0.00 |
| mage-frost | alliance | 628.0 to 628.0 | +0.00 / +0.00 |
| mage-frost | horde | 632.3 to 632.3 | +0.00 / +0.00 |
| paladin-holy | alliance | 656.8 to 676.5 | +3.00 / +3.06 |
| paladin-holy | horde | 644.6 to 662.0 | +2.70 / +2.67 |
| paladin-protection | alliance | 1113.8 to 1113.8 | +0.00 / +0.00 |
| paladin-protection | horde | 1125.5 to 1125.5 | +0.00 / +0.00 |
| paladin-retribution | alliance | 704.2 to 704.2 | +0.00 / +0.00 |
| paladin-retribution | horde | 701.0 to 701.0 | +0.00 / +0.00 |
| priest-discipline | alliance | 652.1 to 664.7 | +1.93 / +1.99 |
| priest-discipline | horde | 655.7 to 669.6 | +2.12 / +2.12 |
| priest-holy | alliance | 720.8 to 807.3 | +12.00 / +11.83 |
| priest-holy | horde | 708.7 to 790.4 | +11.53 / +11.44 |
| priest-shadow | alliance | 684.2 to 684.2 | +0.00 / +0.00 |
| priest-shadow | horde | 691.8 to 691.8 | +0.00 / +0.00 |
| rogue-assassination | alliance | 626.4 to 626.4 | +0.00 / +0.00 |
| rogue-assassination | horde | 622.1 to 622.1 | +0.00 / +0.00 |
| rogue-combat | alliance | 627.0 to 627.0 | +0.00 / +0.00 |
| rogue-combat | horde | 628.2 to 628.2 | +0.00 / +0.00 |
| rogue-subtlety | alliance | 544.7 to 547.2 | +0.46 / +0.53 |
| rogue-subtlety | horde | 544.0 to 546.9 | +0.53 / +0.51 |
| shaman-elemental | alliance | 472.4 to 471.8 | -0.13 / -0.06 |
| shaman-elemental | horde | 480.3 to 479.7 | -0.12 / -0.08 |
| shaman-enhancement | alliance | 586.1 to 585.2 | -0.15 / -0.03 |
| shaman-enhancement | horde | 587.0 to 585.9 | -0.19 / -0.07 |
| shaman-restoration | alliance | 842.3 to 749.9 | -10.97 / -11.05 |
| shaman-restoration | horde | 840.0 to 748.6 | -10.88 / -10.94 |
| warlock-affliction | alliance | 569.9 to 569.9 | +0.00 / -0.02 |
| warlock-affliction | horde | 572.0 to 571.7 | -0.05 / -0.02 |
| warlock-demonology | alliance | 635.4 to 635.3 | -0.02 / +0.03 |
| warlock-demonology | horde | 639.3 to 639.4 | +0.02 / +0.05 |
| warlock-destruction | alliance | 633.8 to 633.8 | +0.00 / +0.00 |
| warlock-destruction | horde | 636.8 to 636.9 | +0.02 / -0.03 |
| warrior-arms | alliance | 717.5 to 717.5 | +0.00 / +0.00 |
| warrior-arms | horde | 729.7 to 729.7 | +0.00 / +0.00 |
| warrior-fury | alliance | 755.3 to 755.3 | +0.00 / +0.00 |
| warrior-fury | horde | 758.2 to 758.2 | +0.00 / +0.00 |
| warrior-protection | alliance | 1081.2 to 1053.1 | -2.60 / -2.49 |
| warrior-protection | horde | 1068.8 to 1042.7 | -2.44 / -2.63 |

Where the moves over 1% at raid 60 come from:
- **Holy priest +11.4 to +12.0%:** the tank's Flash Heal at rank 7 (+10.3 to +10.9%), with Renew rank 1 and Heal rank 1, against the 2.3 scale's -1.3 to -2.2%.
- **Restoration druid +7.7%:** the heavier fight at 2.3 leaves it less overheal.
- **Restoration shaman -10.9 to -11.1%:** the heavier fight runs it dry sooner (mana lasts 268 s, then 162 s); the tank's Healing Wave rank 2 wins back +0.7 to +0.8%.
- **Holy paladin +2.7 to +3.1%:** the scale.
- **Discipline +1.9 to +2.1%:** Renew rank 1 (+0.9 to +1.2%), Lesser Heal rank 1 (+0.3 to +0.4%), the scale.
- **Bear threat -5.4 to -5.5%** and **Protection warrior threat -2.4 to -2.6%:** Demoralizing Roar and Shout now in the rotation. Damage taken falls by 20.2% and 22.8%, and the bear's chance of death goes from 98% to 10% (see O1).

Under 1% at raid 60:
- Subtlety +0.5% (Premeditation in the fight).
- Elemental and Enhancement -0.03 to -0.19% (Searing Totem's full duration).
- Arcane, Affliction and Demonology within ±0.03% (Kindling Stave).
- Protection paladin is identical to the decimal: it has no attack power reduction, and the boss's AP holds the mean swing.
- Every other spec is identical.

Moves over 1% below 60 (the lane reports have every band):
- **Holy priest:** bare 50 +12.5 to +13.8% (the Heal row removed).
- **Discipline:** bare 60 +4.3 to +4.5%.
- **Demonology:** band 20 +1.24%, band 30 +1.10% (Unholy Power); +0.72 to +0.95% at 40 to 60.
- **Healers' bands, from the scale:**
  - shaman 20 -1.9%, 30 -3.4%, 40 -2.3%, 50 -7.0%;
  - Holy 20 -1.3%, 40 +1.2%;
  - Restoration druid coverage at bands 20 to 50 falls from 61 / 92 / 93 / 79% to 50 / 77 / 85 / 67%.
- **Protection paladin band 20:** chance of death -13 to -15 points, threat -2.8%.
- **Tanks' bare bands:** -2 to -6% threat for 8 to 21% less damage taken.

## wow-player review

Two reviewers: physical, tanks, hybrids and core; casters, healers, guides, data and platform.
- **SHIP:**
  - every tank, physical, regmelee, hybrids, core, warlock and data row;
  - casters M1 and S1;
  - platform;
  - every guides row;
  - healer A1.
- **SHIP after fixes:**
  - Protection paladin's Leveling section now names Templar's Bulwark at 20 (S 39136b254).
  - Survival's Falcon sentence is split (S 39136b254).
  - Premeditation's 20 s combo points are confirmed by the Forever tooltip (14183).
  - Holy's "mana outlasts the fight" line, the Restoration shaman rank notes (the AQ rank 9 line, Nature's Swiftness) and Holy's filler wording are fixed by the guides lane (G1 to G3, S 3ca54e3da). The reviewer's two queries check out against the client: Chain Heal's 251/490 base heals, and Fear Ward as a baseline priest trainer spell at 20.
- **Conditional SHIP:** healers P2, P3, F1 and F2, the rank 1 heals. They hold only under Forever's rule that heals learned below 20 keep their whole bonus, which rests on the owner's pending tooltip check (O2).
- **Release-note lines** are in review-physical.md and review-casters.md.

## Owner calls (new)

- **O1. Tanks keep Demoralizing Shout / Roar up.** Adopted under ruling 4's intent and the ruling-1 survival precedent; both reviewers' Classic read agrees. It fails the plain threat metric (bear -5.4%, warrior -2.5%) for -20 to -23% damage taken. To hold the threat metric instead, revert S 27edbc55e and E 75338e087.
- **O2. Low-rank heals.** Read the tooltips for Rejuvenation 1, Renew 1, Heal 1, Lesser Heal 1 and Healing Wave 2 on a geared level-60 character before launch. If any shows a downranking penalty, revert healers P2, P3, F1 and F2 and the guides' "keeps a low rank's whole healing bonus" lines together.
- **O3. Balance raid build.** Heart of the Wild 5 and Reflection 3 in place of Nature's Focus 5 and Subtlety 3 gains +1.6% at both seeds. Both dropped talents are on the keep list, so it is held (casters.md B3).
- **O4. Elemental Mastery as one charge** would cost Elemental 2.4%; an in-game check settles it (casters.md E1). Superseded if 70755 removes the talent (see below).
- **O5. Warlock races.** Orc is now the Horde pick for Affliction and Destruction (beyond 1%). Demonology stays Troll (+0.8% for Orc, under the race rule). The reviewer suggests Orc for all three, for consistency.
- **O6. Band coverage.** Should the heal-scale rule apply per band? The reviewer says not before launch.
- **O7. The addon at 60** shows raid-60 ranks with bare BiS and bare dps-per-point; the reviewer says 60 should be the raid character.

## Client 1.60.1.70755 (build lane)

A separate lane is activating 70755. The build lane did a read-only diff:
- **Client files.** All 42 sim-relevant tables match 70753 exactly. Every change is a server hotfix push (wago 2026-11-03 20:06 UTC), and the rogue and warrior pushes also reached 70707 clients (17:34 and 18:29 UTC). These changes are probably live on servers whatever client build players run.
- **Shaman:**
  - large shock and Lightning Bolt buffs (Lightning Bolt rank 10 base 196 to 353);
  - Chain Lightning gets a higher base and coefficient, a 2.5 s cast and more mana;
  - Stormstrike's cooldown drops from 8 to 6 s and it gains a 150% weapon-damage effect;
  - Riptide stops being a direct heal, on a 12 s cooldown.
  - **Lava Burst's trainer rows are removed, and its talent node is now Call of Thunder.** That contradicts the activation lane's "Lava Burst becomes trainable"; recheck.
- **Binding Heal's trainer rows are removed,** but the Discipline and Holy APLs still cast it.
- **Rogue:**
  - Mutilate and poisons up about 40 to 50%;
  - Venom reworked;
  - Shank and Death Mark added as baseline abilities.
  - All 17 published rogue talent strings become illegal.
- **Warrior:** Slam's cooldown goes from 18 to 15 s, and Improved Slam from -3 to -5 s. Raging Blows becomes Improved Whirlwind. Arms raid's 5/5 picks in Improved Tactical Mastery and Bloodthrill exceed the new caps of 3 and 2.
- **Other classes:** Wrack's coefficient doubles; Vengeance stacks to 5; Two-Handed Weapon Specialization's weapon-type value changes; Sniper Shot becomes a trainer spell.
- **Unchanged:** racials, and every published BiS item's stats and effects.

Round-14 fixes that touch the moving ground, to re-measure on 70755: Subtlety's Premeditation (E e01d5596a) and the rogue rows; Elemental and Enhancement's Searing Totem (E 2fe671caa); the Holy and Discipline rank rows next to Binding Heal; Affliction (Wrack). Fury has no round-14 change.

## For the integrator

**Branches.**
- Site `int/r14`: every lane branch (core, tanks, healers, regmelee, regcast, warlock, casters, physical, hybrids, data, guides, platform, build) plus the coordinator's merge and review commits, on origin/main 72471fcb3.
- Fork `int/r14`: 18 lane commits on 5195a8889.
- No fork master fast-forward.
- If 70755 lands first, rebase or merge int/r14 onto it and re-measure the moving-ground fixes listed above.

**Pin first:** data/proto/ENGINE_SHA, sim/enginever/version.go, web/src/lib/sim/version.ts and the web sim fixtures, at the fork int/r14 head.

**Regenerate, in this order.**
1. **apl-sync** (`make apl-sync`, then `make apl-check`): the 12+ APLs whose notes changed (guides), the tank and healer APLs (already hand-synced on both sides; the sync is a no-op check), and the fork ui copies of the note text.
2. **addon-builds.json** (`go run ./sim/cmd/addon-builds -repo-root .`): Protection paladin band 20 (Templar's Bulwark), Demonology's leveling build (Unholy Power). TestCommittedAddonBuildsFollowTheGuides fails until then.
3. **sim/request TestRotationLadder goldens** at the pin. These move:
   - the three tanks;
   - the five healers;
   - hunter-beast-mastery (bands 40, 50, bare 60);
   - priest-shadow (bare 60), priest-holy (50);
   - rogue-subtlety (40 to 60);
   - Elemental and Enhancement (Searing);
   - warlock-destruction (40), Demonology's bands, warlock bands below 60 (Immolate cast speed).
4. **addon-data** (`pipeline addon-data` → addon-data.json and Data.lua): rank rows, changed rotations and notes.
5. **Coach:**
   - pipeline.coachrules: hunter-beast-mastery, hunter-marksmanship, the three warlocks, warrior-arms, warrior-fury;
   - `python tools/coach/persona_starter.py --all`, which reaches priest-shadow (still the old priority on int/r14), the hunters, the priests, the tanks and every changed guide page;
   - restamp `reviewed:`.
   test_committed_files_match_the_active_build and the starter files-match test fail until then.
6. **dps-per-point.json**, then Data.lua: druid-feral (weapon term back, about 0.8 to 1.1 at the bands and 2.3 at 60), rogue-subtlety, the five healers, Elemental and Enhancement, priest-holy band 50, Beast Mastery bare.
7. **Nightly-owned, at the refresh:** see "Measure at refresh".
8. **App fixtures after the refresh:** cdn.bis.warrior-fury and cdn.bis.hunter-beast-mastery.

**Gates run here.** No lane ran a full suite. On int/r14 I ran:
- the fork's conformance, core, druid, encounters, hunter, mage, paladin, priest, rogue, shaman, warlock and warrior packages, with `--tags=with_db` (green after M4);
- site sim/request, sim/leveling, sim/rankchoices, cmd/dps-per-point and cmd/spec-breakdown (skipping only TestRotationLadder, TestCommittedAddonBuildsFollowTheGuides and TestTrueshotAuraRanksMatchTheClient, which wait on regens);
- the guide vitest set, src/lib/mobile and app-files;
- data test_apl.py;
- `make apl-check`.

Not run: iOS `xcodebuild test`.

Heads at the end of round 14 are on the last line of this section.

## Measure at refresh (added this round)

- **MR26:** the three tanks on the boss with attack power and the Shout/Roar lines: picks, weights, tiers (damage taken now moves). Protection paladin band 20 (Templar's Bulwark).
- **MR27:** every healer on the 2.3 profile: picks at every band, the heal-profile block in each BiS file, tiers; the new rank rows (Flash Heal 7, Renew 1, Heal 1, Lesser Heal 1, Healing Wave 2, Holy Shock 4 at bare 60).
- **MR28:** Horde Affliction and Destruction ranked on Orc; Demonology bare bands on Unholy Power.
- **MR29:** the cat's weapon DPS term in BiS weights and the addon's tooltips.
- **MR30:** Subtlety picks and weights with Premeditation in the fight; Gnome rogues and warriors (Expansive Mind, within 1%, race picks stay).
- **MR31:** Beast Mastery bare 40/50/60 and Destruction 40 after the rank rows; Shadow bare 60.

## Still open

- **Owner:** O1 to O7 above; round 13's pending in-game checks (ruling 5).
- **Blocked:** the level-60 heal split from Warcraft Logs (gcloud reauth). Deep Wounds' double multiplier and other level-60 log items wait on it too.
- **Engine, OPEN:**
  - Discipline's 0% haste step (C4).
  - Cooldowns under the same 10 ms batch rule, which would move every spec (core proposal).
  - Maelstrom Weapon's 50% against 10 a minute (a fast one-hander log settles it; +0.14% at raid).
  - Penance's pulses under Inner Focus' crit mask (under 0.1%).
  - Lacerating Strikes losing unticked damage on refresh, and the Raptor Strike queue's off-hand miss gap (Survival logs).
  - Ranged Weapon Specialization on Serpent Sting (at most 0.3%).
  - Aspect of the Falcon's own proc.
  - Whether Shout and Roar can miss a raid boss.
  - Wowhead's 210/207 against the client's 204.
  - Holy Shield crit.
  - Holy paladin Holy Light rank 3 at bare 60 (+4.2%, but it ends with mana unspent).
  - Discipline Shield rank 10 (+0.4 to +0.6%, held for the unmodeled rage cost).
  - The "Endless" elixirs.
  - Kings with Heart of the Lion.
  - The Judgement of Wisdom/Light proc roll.
  - Fire's stale haste weight in the BiS files (refresh).
  - Kindling Stave's proc rate.
- **Coach starters** still use "the site's guide" phrasing, which the meta-talk guard does not scan (review G5).
- **Guides:** Swiftmend consuming the HoT, Faerie Fire free in forms, form swing timers, Portal of Summoning's shards.
- **70755:** every build.md row, pending the activation lane's pin.
- **The round was not zero:** round 15 follows, on 70755 once it is pinned.

Heads at the end of round 14: fork `int/r14` b852f9127; site `int/r14` e33b3a4c9 (this report), plus the one commit that records this line.
