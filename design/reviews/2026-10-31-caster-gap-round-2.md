# Caster gap, round 2 (2026-10-31)

Level-60 raid band, alliance raid-preset BiS, `sim/cmd/spec-breakdown`, 10000 iterations at seeds 1 and 50001. Seeds closer than the iteration count share iterations, because iteration i runs at seed + i. Every error is under ±0.1%. Engine 816ec57c0, client 1.60.1.70658. Relative figures only.

## Fixes

| Spec | Root cause | Change | vs round one |
|---|---|---|---|
| Elemental | The engine had Elemental Mastery commented out. Forever's version (spell 16166) takes 0.5 s off Lightning Bolt and Lava Burst and removes Chain Lightning's cooldown for 15 s, on a 3 min cooldown. | The fork implements it. The raid build takes it (31st point) and moves Totemic Focus into Thundering Strikes and Ancestral Knowledge: the talent-search winner that keeps the raid keep list. | about +3.7% |
| Balance | The engine already modeled Eclipse, but the build had no points in it and the rotation never cast Wrath for charges. | Wrath fires whenever no charge is banked. Fight-end clips key off the real cast times. The raid build takes Eclipse 3/3 and Improved Wrath 5/5 in place of Heart of the Wild. | about +8.6% |
| Frost | Cone of Cold had no reach. A raid frost mage stands at range, so round one's Cone of Cold gain was an upper bound. | The fork refuses the cast beyond 10 yd (12 with Arctic Reach). The ranker stands casters at 30 yd, and spec-breakdown now does too. `winters_chill` leaves the raid preset: in Forever it only feeds the applying mage's own Cone of Cold, and here it fed ours for free. | about −3.4% |

Fury and Fire are unchanged. As a share of Fury: Elemental goes from about 55% to 57%, Balance from 57% to 62%, Frost from 62% to 60%. Fire stays at about 95%.

## What the remaining gap is

Checked and correct in the engine:
- Spell power coefficients, crit and resistance. Lightning Bolt's average hit reproduces from the client's base, 0.714 coefficient, Concussion, the 10% Curse of the Elements and a 2.0x crit (Elemental Fury). Frostbolt's reproduces with Ice Shards' 2.0x crit and no partial resist, because Frostbolt is binary.
- Boss resistance. Every school takes the same level-based mitigation against a level-63 boss, and no Frost- or Nature-specific resistance is applied.
- Hit. Casters sit about 7% short of the 16% spell cap, which costs them about 10% of their damage. The pre-raid item pool causes this, not the scoring: the ranker weighs 1% hit at about 10 spell power. Fire and Destruction have the same shortfall and still match Fury, so hit does not explain the spec gap.
- Raid buffs. Casters receive their buffs and consumables: Arcane Intellect, Divine Spirit, Wisdom, Moonkin Aura, Curse of the Elements, Supreme Power, Greater Arcane Elixir, Wizard Oil, mana potions and Demonic Rune.

Forever-specific (client data, confirmed against Forever tooltips):
- **Lower base damage on caster nukes.** Lightning Bolt rank 10 is 191 to 213 at a 0.714 coefficient (Classic: 428 to 477 at 0.857). Starfire rank 7 is 364 to 426 (Classic: 496 to 584). Frostbolt rank 11 is 470 to 506 (Classic: 515 to 555). Elemental and Balance carry the largest cuts.
- **Melee raid buffs dwarf caster buffs.** Removing the melee-only buffs costs Fury 31%. Removing every caster-relevant buff costs Frost 14%. Heart of the Lion alone (a raid-wide 10% stats and 276 attack power) costs Fury 14% and Frost 1%.
- **Winter's Chill is no longer +10% Frost crit for the raid.** It only empowers the applying mage's Cone of Cold, which a ranged frost mage does not cast.
- **No Ahn'Qiraj book ranks.** `spellranks.json` marks Frostbolt rank 11 and Starfire rank 7 as books, and the engine stops one rank short (`core.IncludeAQ`). If Forever trainers teach them, Frost and Balance gain. SkillLineAbility cannot tell a book rank from a trainer rank, so this stays as data says.

Era ratio check: phase 1 to 3 parses put mages within 10% to 25% of top warriors, with Balance and Elemental further behind. Fire mage at about 95% of Fury fits. Frost at about 60% sits below Era frost because of the Winter's Chill rework, Frostbolt's lower base, and Heart of the Lion. Balance and Elemental fall where Era puts them, deeper because of their nukes' base-damage cuts. No further modeling error was found.
