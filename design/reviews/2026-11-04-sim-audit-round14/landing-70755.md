# Round 14 on client 1.60.1.70755: the landing re-measure

Round 14 was written on 70753. It lands on main's 70755 activation (126b9c287, engine 2eaae03c6). The fork merge is d1d9f2775: fork int/r14 b852f9127 merged onto 2eaae03c6.

**Before:** origin/main 126b9c287 with engine 2eaae03c6.
**After:** the merged site with engine d1d9f2775.
Both use spec-breakdown at the raid preset, band 60, 10,000 iterations, seeds 1 and 50001, both factions, all 28 specs.

Six specs' published BiS talents (carried from 70753) are illegal on 70755's trees: the three rogues, Enhancement, Arms and Fury. They were run on their guide raid builds (`-talents`) on both sides.

## Moves over 1%

| Spec | Faction | Seed | Before | After | % |
|---|---|---|---|---|---|
| druid-feral-bear | alliance | 1 | 1020.2 | 964.2 | -5.49% |
| druid-feral-bear | alliance | 50001 | 1020.2 | 965.2 | -5.39% |
| druid-feral-bear | horde | 1 | 1011.2 | 955.7 | -5.49% |
| druid-feral-bear | horde | 50001 | 1010.8 | 956.3 | -5.39% |
| druid-restoration | alliance | 1 | 934.2 | 1006.4 | +7.73% |
| druid-restoration | alliance | 50001 | 934.3 | 1006.7 | +7.75% |
| druid-restoration | horde | 1 | 933.7 | 1005.3 | +7.67% |
| druid-restoration | horde | 50001 | 933.3 | 1005.2 | +7.70% |
| paladin-holy | alliance | 1 | 656.8 | 676.5 | +3.00% |
| paladin-holy | alliance | 50001 | 656.9 | 677.0 | +3.06% |
| paladin-holy | horde | 1 | 644.6 | 662.0 | +2.70% |
| paladin-holy | horde | 50001 | 645.2 | 662.4 | +2.67% |
| priest-discipline | alliance | 1 | 593.5 | 584.4 | -1.53% |
| priest-discipline | alliance | 50001 | 594.1 | 584.1 | -1.68% |
| priest-discipline | horde | 1 | 585.7 | 569.8 | -2.71% |
| priest-discipline | horde | 50001 | 585.6 | 569.8 | -2.70% |
| priest-holy | alliance | 1 | 720.8 | 807.3 | +12.00% |
| priest-holy | alliance | 50001 | 721.7 | 807.1 | +11.83% |
| priest-holy | horde | 1 | 708.7 | 790.4 | +11.53% |
| priest-holy | horde | 50001 | 709.2 | 790.3 | +11.44% |
| shaman-restoration | alliance | 1 | 770.8 | 719.8 | -6.62% |
| shaman-restoration | alliance | 50001 | 771.1 | 719.8 | -6.65% |
| shaman-restoration | horde | 1 | 769.2 | 718.8 | -6.55% |
| shaman-restoration | horde | 50001 | 768.8 | 719.0 | -6.48% |
| warrior-protection | alliance | 1 | 1081.2 | 1053.1 | -2.60% |
| warrior-protection | alliance | 50001 | 1080.3 | 1053.4 | -2.49% |
| warrior-protection | horde | 1 | 1068.8 | 1042.7 | -2.44% |
| warrior-protection | horde | 50001 | 1069.4 | 1041.3 | -2.63% |

Healers are in effective HPS and tanks in threat per second.

These match round 14's own 70753 table except for two specs:
- **Discipline turns from +1.9 to +2.1% (70753) into -1.5 to -2.7%.** On the merged tree, main's Discipline list in place of round 14's measures 579.4 / 566.1 at seed 1 (alliance / horde). So round 14's rank rows (Renew 1, Lesser Heal 1) still gain +0.7 to +0.9%. The loss is the 2.3 heal scale, which costs Discipline more once 70755 took Binding Heal away. The rows are kept.
- **Restoration shaman is -6.5 to -6.7%,** against -10.9 to -11.1% on 70753. It is the same scale, on 70755's shaman retune.

## Under 1%

- Subtlety: +0.57 to +0.63% (Premeditation in the fight, on 70755's tree).
- Enhancement: -0.04 to +0.11%; Elemental: -0.02 to -0.05% (Searing Totem's full duration).
- Arcane, the three warlocks: within ±0.05%.
- Every other spec is identical. Wrack, which 70755 raised, is not in the Affliction raid build.

## Kept and dropped

- **Kept:** every fork fix (18 commits) and every site fix.
- **Dropped as moot on 70755:**
  - the Binding Heal rank rows' tests (Discipline at 45/55, the Undead bare-60 race row);
  - the Lava Burst coach-starter correction;
  - casters O4 (Elemental Mastery as one charge), since 70755 took the talent out.
- **Lightning Overload:** the engine already deals each overload's own client row (LightningBoltOverload* and ChainLightningOverload*). No change.
