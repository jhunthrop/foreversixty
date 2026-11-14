# Sim accuracy 2026-11-13: Windfury follow-ups

Lane `wf-followup`, brief `design/reviews/sim-accuracy/brief.md`, after the warriors lane (`warriors-70938.md`, engine pin f0f90b5f9). Two jobs: the red Enhancement card-note guard on main, and the warriors lane's O5 (do Retribution's Seals react to Windfury Totem's strike since push 113887).

**Totals: 0 engine bugs; 1 card-note fix.**

## N1: Enhancement's Flame Shock note named Stormstrike at band 10 (fixed, note)

`data/tests/test_card_note_levels.py` failed on origin/main for `shaman-enhancement` band 10: the Flame Shock line added by the Enhancement lane (0f77fad48) ended "Below a third, keep the mana for Stormstrike and Earth Shock.", and Stormstrike is a level 25 talent. The data is right (the line runs from 10, the mana floor protects Earth Shock there), the note was unscoped. Now "Below a third, keep the mana for Earth Shock and for Stormstrike once you have it." in both apl copies (notes only, so no sync or pin). `test_card_note_levels.py`, `test_apl.py`, `test_addonrotation.py`: 114 passed. Pushed as 692113211.

## O5 of the warriors lane: Seals and the Windfury Totem strike (CONFIRMED-CORRECT, no change)

**Client** (wago.tools 1.60.1.70938, `useHotfixes=1`, fetched 14 Nov ~04:20 UTC, so push 113887 and later included):

| spell | SpellAuraOptions | reads |
|---|---|---|
| Seal of Command 20375, 20915, 20918, 20919, 20920 | ProcTypeMask 0x4, 100%, ProcCategoryRecovery 1000 | melee auto attack done only |
| Seal of Righteousness 20154 to 20293, 21084 | ProcTypeMask 0x4 | melee auto attack done only |
| Seal of Fury 20163, 20419 to 20423, 1311649, 1311656 | ProcTypeMask 0x4 | melee auto attack done only |
| Unsealed Wrath 1310735 | ProcTypeMask 0x4, 100 ms | melee auto attack done only |
| Windfury Totem aura 8515, 10609, 10612 (for contrast) | ProcTypeMask 0x14, 20%, 100 ms | auto attack and melee abilities |

The strike itself (8516, 10608, 10610) is effect 58 (weapon damage), DefenseType 2 (melee), Attributes_3 0x400 (main hand) and 0x200 (a triggered spell that may proc). It is a melee spell, so it raises the melee ability flag (0x10), never the auto attack flag (0x4). The 0x200 attribute lets it proc auras that listen for 0x10 (the totem's own mask carries 0x10, so it is why a strike can in turn feed things like Flurry's trigger); it does not make it an auto attack. Every Seal listens for 0x4 alone, so no Seal rolls on the strike, and Unsealed Wrath neither stores nor releases on it. Before the push the totem granted SPELL_EXTRA_ATTACKS, a real white swing, which is why Seals did react to it then.

**Logs.** No report since push 113887 has a Retribution paladin in a Windfury group (the only post-push report, ju2t73py6pib, is two Enhancement shamans; checked by the warriors lane at 14 Nov 03:00 UTC). The same 0x4 rule is already confirmed on our launch logs for a melee ability: 0 of 33 Holy Strikes added an Unsealed Wrath stack while 62 of 62 stacks came from white swings (`sim/paladin/unsealed_wrath.go`), and Holy Strike raises the same 0x10 flag as the strike.

**Engine.** `core.ApplyWindfury` casts the strike as `ProcMaskMeleeMHSpecial`; `isSealSwing` gates every Seal and Unsealed Wrath on `ProcMaskMeleeWhiteHit`. That is the client's 0x10 against 0x4 exactly. No fork change, no pin; the pin stays f0f90b5f9. Retribution raid 60's +0.1 to +0.2% from the warriors lane (the strike's own damage, no Seal procs) stands.

Open: a post-push log with a Retribution paladin under a Windfury Totem would confirm it from the fight side (SPELL_DAMAGE 10608 with no Seal of Command 20424 at the same timestamp).
