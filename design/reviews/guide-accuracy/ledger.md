# Guide-accuracy ledger

Rounds run until one finds zero confirmed errors. Each round: one expert auditor per class on the active client's data (hard evidence only), class fix lanes, a wow-player SHIP per class on rendered pages, one landing. Reports: `round<N>/<class>.md` (audit, fix log, fix passes).

Columns: **Audit** = CONFIRMED findings in the audit (a finding that touches several pages counts on each); **Review** = MUSTs the rendered-page reviews found and the fix passes fixed; **Fixed** = Audit + Review fixed this round (plus LIKELY findings the lanes proved with data, in the class report); **Left** = what waits on the owner or the next sim round; **sha** = the commit carrying the round's fixes (land/guides-r1 9acd0a78c, landed on main with the branch).

## Round 1 (client 1.60.1.70755, 2026-11-04)

| Spec | Audit | Review | Fixed | Left | sha |
|---|---|---|---|---|---|
| warrior class page | 0 | 0 | 0 | planner links (owner, kept by R1) | 9acd0a78c |
| warrior-arms | 1 | 0 | 1 | Arms opener Bloodrage vs Charge (owner) | 9acd0a78c |
| warrior-fury | 2 | 1 | 3 | Improved Berserker Rage (sim) | 9acd0a78c |
| warrior-protection | 1 | 2 | 3 | Improved Shield Wall (owner); Expertise cap footnote (component) | 9acd0a78c |
| shaman class page | 0 | 0 | 0 | | 9acd0a78c |
| shaman-elemental | 3 | 0 | 3 | | 9acd0a78c |
| shaman-enhancement | 1 | 3 | 4 | Air totem action (sim); Stormstrike level gate (sim) | 9acd0a78c |
| shaman-restoration | 0 | 3 | 3 | NS into rank 2 Healing Wave (sim); Chain Heal party or raid (owner); low-rank tooltip (owner) | 9acd0a78c |
| rogue class page | 0 | 0 | 0 | | 9acd0a78c |
| rogue-assassination | 4 | 1 | 5 | Mutilate off-hand penalty (owner) | 9acd0a78c |
| rogue-combat | 2 | 2 | 4 | Efficient Strikes on a Backstab raid build (sim) | 9acd0a78c |
| rogue-subtlety | 3 | 1 | 4 | Shank/Ghostly Strike/Eviscerate gating, Shank rank 1 (sim) | 9acd0a78c |
| priest class page | 1 | 1 | 2 | | 9acd0a78c |
| priest-discipline | 2 | 1 | 3 | low-rank heal ranks (owner tooltip, sim) | 9acd0a78c |
| priest-holy | 1 | 1 | 2 | low-rank heal ranks (owner tooltip, sim) | 9acd0a78c |
| priest-shadow | 5 | 3 | 8 | leveling order and Spirit Tap 5/5 (sim); SW:P cutoff (sim) | 9acd0a78c |
| mage class page | 2 | 0 | 2 | | 9acd0a78c |
| mage-arcane | 5 | 2 | 7 | Arcane Power timing (sim) | 9acd0a78c |
| mage-fire | 3 | 2 | 5 | Accelerant/Impact vs Burning Soul (sim); haste weight (sim) | 9acd0a78c |
| mage-frost | 2 | 2 | 4 | Accelerant point (sim) | 9acd0a78c |
| warlock class page | 1 | 0 | 1 | | 9acd0a78c |
| warlock-affliction | 4 | 0 | 4 | Wrack after 70755, Agony vs Doom (sim) | 9acd0a78c |
| warlock-demonology | 3 | 1 | 4 | Horde race (owner) | 9acd0a78c |
| warlock-destruction | 4 | 1 | 5 | Molten Skin/Destructive Reach points (sim); Curse of the Elements order (owner) | 9acd0a78c |
| hunter class page | 1 | 0 | 1 | | 9acd0a78c |
| hunter-beast-mastery | 2 | 1 | 3 | Focused Fire before Endurance Training (ladder-first, sim/owner) | 9acd0a78c |
| hunter-marksmanship | 5 | 3 | 8 | Sniper Shot, Rapid Killing (sim) | 9acd0a78c |
| hunter-survival | 4 | 1 | 5 | Deadly Aspects under Falcon (owner) | 9acd0a78c |
| druid class page | 1 | 0 | 1 | | 9acd0a78c |
| druid-balance | 1 | 0 | 1 | Rebirth in Moonkin Form (owner) | 9acd0a78c |
| druid-feral | 5 | 3 | 8 | cat leveling build's 8 Nature's Focus/Subtlety points (sim, SIM-D1) | 9acd0a78c |
| druid-feral-bear | 6 | 0 | 6 | Rend and Tear 2% in the engine (sim) | 9acd0a78c |
| druid-restoration | 2 | 1 | 3 | Rejuvenation rank 1 tooltip (owner) | 9acd0a78c |
| paladin class page | 0 | 0 | 0 | | 9acd0a78c |
| paladin-holy | 2 | 2 | 4 | Oath of Light at band 30 (sim) | 9acd0a78c |
| paladin-protection | 4 | 2 | 6 | Lay on Hands mana (owner) | 9acd0a78c |
| paladin-retribution | 5 | 1 | 6 | Dwarf vs Human (sim); Exorcism crit (owner) | 9acd0a78c |

Round 1 totals: 72 audit CONFIRMED by class (the per-spec sums above count multi-page findings once per page), about 40 rendered-review MUSTs, all fixed; two ad-review claims checked and rejected (Flame Shock 21 s, Fury opener). Verdict: NOT zero, round 2 needed.

Out of scope this round, for the site and sim lanes: stat-table Expertise cap footnote on melee specs (no Forever item carries expertise; engine weights still list it), hit-cap split footnotes, generated "raid build moves N points" summaries that drop talents (web/src/lib/guides/build-diff.ts), generated recommended-race line duplicating prose, the top Build card showing the leveling build, cooldown cards ignoring talent cuts (Sacred Duty, Rapid Killing), Discipline card "?" icon and duplicate rows, site-voice footers ("Numbers from client build ...", "Raid loot is not itemized ...", "enchants are not listed yet"), repo paths in source labels.
