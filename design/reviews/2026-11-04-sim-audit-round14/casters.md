# Sim audit round 14: casters, 2026-11-03

Specs: mage-arcane, mage-fire, mage-frost, druid-balance, shaman-elemental, priest-shadow.

## Method

- **Branches.** Fork `lane/r14-casters` from int/r14 (5195a8889). Site `wt/r14-casters` from int/r14 (72471fcb3), with `make simdb`.
- **Client.** 1.60.1.70753 rows (`r11/c70753/.../raw`, read with round 12's `healers/rows.py`) plus the 2026-11-03 hotfix push. The push touches only hunter player spells (round 13 hotfix X1: of its 91 spell ids, only family 9 is a player class). Nothing in this group moves with it.
- **Re-audit.** Three read-only passes, one each for mage, Balance plus Elemental, and Shadow. They read every APL cast, every talent in the published raid-60 and bare-60 strings (decoded against the client tree order) and the raid-60 BiS gear, and compared each against the client rows. I checked every candidate by hand, against a client row or with a run.
- **Probes.** Two engine probes were built into separate binaries and reverted before any commit:
  - a cast-speed multiplier set from an environment variable;
  - Elemental Mastery as a one-charge buff.
- **Runs.** spec-breakdown, 10,000 iterations, seeds 1 and 50001, one at a time, `nice -n 15`, GOMAXPROCS=9.
- **Tests.** The fork's `sim/priest/...` and `sim/mage` packages (`-tags=with_db`). Both goldens are unchanged.

## Mage (Arcane, Fire, Frost)

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| M1 | **Combustion's buff capped at 20 stacks; the client caps it at 10.** 28682 gives +10% Fire crit a stack. Three crits end it long before ten stacks, so no number moves. | SpellAuraOptions 135400 (28682) CumulativeAura 10; `mage/talents.go` `MaxStacks: 20`. `TestCombustionStacksToTen` failed, then passed; TestP1Mage unchanged. | CONFIRMED-BUG (no published number) | old (upstream; no round read the buff's stack cap) | E eff1d0491 |
| M2 | **Fire's negative haste weight is stale, not a model fault.** On the current engine Fire is flat across haste: raid 60 horde, 3,000 iterations, 673.7 / 673.8 / 673.6 / 673.9 / 674.5 at 0 / 1 / 3 / 5 / 10% cast speed. Scorch casts stay at 70.8. Every Fire cast already sits on the 1.5 s global: Scorch, three-stack Pyroblast and the instant Fire Blast. The global is not hasted. The published −3.5 (raid) and −6.6 (bare) come from engine 8baffe5fc (round 5). | haste probe; `core/cast.go:244` | Aligned (MR25) | - | - |
| M3 | Ignite's ticks skip the target's damage-taken modifiers. The research pass argued from SpellMisc attributes (no 0x100 bit on 412545) that they should take them, which would be about +4%. The engine follows Blizzard's beta note "Ignite no longer double dips on % damage increase modifiers" (round 2 F1, round 10 A2). A patch note about Forever beats a missing attribute bit. | `mage/ignite.go:55-61` | Aligned (patch note) | - | - |
| M4 | Fresh, aligned with the client rows:<ul><li>**Arcane:** Arcane Blast 1239700 (236 + 1.6/level, 0.429, 15% of base mana, stack 400573: 12% of Intellect, +100% cost, 4 stacks, 18 s from the first cast); Arcane Missiles 10212/10274; Missile Barrage 400588/400589 (40% / 20%, −50% channel, free, 15 s); Arcane Power (+30% damage and cost, 15 s, 180 s); Presence of Mind (bit 30); Arcane Concentration → Clearcasting 12536.</li><li>**Fire:** Combustion's mask, 3 crits and its cooldown from the buff's end; Heating Up 400625 (−25% a stack); Improved Scorch 22959; Fire Power; Critical Mass; Incineration; Accelerant; Improved Fire Blast; Master of Elements (30% of base cost); Ignite (40% over two 2 s ticks, rolling, attacker multiplier 1).</li><li>**Frost:** Winter's Chill 12579 (+20% Cone of Cold a stack, +7% spell power); Fingers of Frost and Shatter; Ice Lance ×4 frozen; Ice Barrier (15% + 5%, 30 s); Evocation (not hasted).</li><li>**Shared:** mana gems and Demonic Rune on category 1153.</li><li>**Raid BiS:** Magister's 3-piece; Robe of the Archmage; Frozen Heart (+29 Frost/Shadow use); Rune of the Guard Captain (stats only); Kindling Stave's Ignition 1300782 (70 ± 15%).</li></ul> | engine against SpellEffect, SpellMisc, SpellPower, SpellAuraOptions | Aligned | - | - |
| M5 | **"Ignition" (round 13's open log item) is not a mage talent.** It is Kindling Stave's equip proc (1300781 → 1300782), modelled at one a minute (`common/item_effects/kindling_stave.go:47`). Only raid Arcane wears the stave, and it is 0.3% of Arcane's damage. Its rate is server-side. | client rows; breakdown | OPEN (proc rate, carried) | old | - |
| M6 | Arcane Subtlety's "−15 resistance to all your spells" is not modelled. Against the boss's 0 resistance it cannot act, and the level-based part is not penetrable. | `mage/talents.go:706`; `core/spell_resistances.go:146` | Aligned (no effect) | - | - |
| M7 | Arcane Concentration 11213 has no class mask and ProcTypeMask 0x15550, which includes ranged auto. The engine procs only from mage spells, so Shoot and the stave's Ignition never give Clearcasting. That is about 0.4 Clearcastings a fight on Arcane, mana only. The mask is the only source, and vanilla talents often carry this broad mask while a server rule narrows it. | SpellAuraOptions 11213; `mage/talents.go:238` | OPEN (negligible) | - | - |
| M8 | Carried open items, with no new evidence: <ul><li>cost modifiers add (Arcane Blast stacks plus Arcane Power; Arcane Power on a free Missile Barrage);</li><li>Master of Elements on a Clearcasting crit;</li><li>Ice Lance's coefficient: the client column is 0, the dummies are 300/20/5, and Wowhead's text shows no spell-power term. Round 6 measured 0 at −9.7% for Frost.</li></ul> | `arcane_blast.go:102`, `talents.go:341,412`, `constants_auto_gen.go:2752` | OPEN | old | - |
| M9 | Frost never casts Cone of Cold at raid 60. The caster stands at 30 yd and the cone reaches 10 yd; the APL note says so. The `expected_idle` reason "From 38 on it casts" is out of date: it is not player text, but it misleads. | breakdown; `data/curated/apl/mage-frost.json` | Aligned (note stale) | - | - |

## Balance

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| B1 | **Eclipse timing (round 13 B4) settled by the closest analogue.** Eclipse 408248's ProcTypeMask is 0x10000, the same single bit as Missile Barrage 400588. The logs show Missile Barrage procs at the cast, misses included (round 13 M1, M3). Nature's Grace, which procs on impact, carries a different mask (0x14000). The engine banks Eclipse at the end of Wrath's cast, hit or miss. No Eclipse log exists yet. | SpellAuraOptions 223856, 222941, 128854; `druid/talents.go:712` | Aligned (identical-mask analogue) | - | - |
| B2 | Fresh, aligned with the client rows: <ul><li>**Spells:** Wrath 9912 (0.571, 120 mana); Starfire 9876 (1.0, 3.5 s − Improved Starfire); Moonfire 9835 (0.15 direct, 0.13 a tick, 4 + 1 ticks); Insect Swarm 24977; Innervate.</li><li>**Talents:** Vengeance (2.0× crit); Nature's Grace (10%, −150 ms global, 3 s, on impact); Moonfury; Genesis; Nature's Majesty; Nature's Reach; Naturalist; Eclipse 170/330/500 ms.</li><li>**Moonkin Form and Omen of Clarity:** Moonkin's +3% (aura 290) is not doubled with the raid's Moonkin Aura; Omen is doubled in Moonkin.</li><li>**Raid BiS:** Idol of the Moon (+33); Ironfeather 2-piece (+20); Hand of Edward the Odd's proc needs melee (never fires).</li></ul> | engine against client rows; Wowhead 23197 | Aligned | - | - |
| B3 | **Raid build: Nature's Focus 5 and Subtlety 3 do nothing in the fight.** Heart of the Wild 5 and Reflection 3 instead (`.../0/055300`) measure +1.58 / +1.55% (horde, 574.9 → 584.0 at seed 1, 575.1 → 584.0 at 50001, ±0.3). Both points are on the raid-keeps list (pushback and threat matter in a real raid), and the adoption rule keeps unmodelled talents. **Not adopted; owner decision.** | spec-breakdown `-talents`; `design/reviews/talent-search/raid-keeps.json:4` | Held (owner) | - | - |
| B4 | Casters fight at 30 yd (`sim/leveling/distance.go:45`), so a speed-20 Wrath flies 1.5 s. Round 13's logged Wrath flights have a median of 0.67 s, but those are beta leveling pulls. A raid caster's distance is a positioning choice, not a client value. | round 13 B1 data | OPEN (raid distance) | - | - |
| B5 | Carried: Improved Moonfire on base damage only (the client mods are aura 108 misc 0 and misc 22, which emulators apply to the whole damage; that is a single non-client source); the Moonglow and Improved Wrath cost cuts add; Starfire rank 7 is a book. | `druid/talents.go:509,568` | OPEN (carried) | old | - |

## Elemental

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| E1 | **Elemental Mastery's charge (round 13 E4): the size is now measured.** The client has one ProcCharge (SpellAuraOptions 128637, mask 0x15550). The Forever tooltip reads "Lasts 15 sec" with no charge and "removing the cooldown from your Chain Lightning", which reads as a window. If the buff is spent by the first Lightning Bolt, Lava Burst or Chain Lightning, raid 60 horde goes from 480.3 / 480.0 to 468.6 / 468.6, **−2.4%**. The engine keeps the 15 s window. | probe binary (not committed); https://nether.wowhead.com/forever/tooltip/spell/16166 (read 2026-11-03) | OPEN (a level-40+ log or an in-game test) | old | - |
| E2 | Fresh, aligned with the client rows: <ul><li>**Spells:** Lava Burst 1326047 (220 + 1.3/level, 0.714, 2.5 s, 10 s, +20% on Flame Shock); Lightning Bolt; Chain Lightning; Flame Shock; Mana Spring 10497.</li><li>**Elemental Mastery's mods:** −500 ms on Lightning Bolt and Lava Burst (mask word 1 0x1000), and no Chain Lightning cooldown.</li><li>**Talents:** Elemental Focus → Clearcasting 16246; Convection; Concussion; Call of Flame; Elemental Alacrity; Lightning Overload; Elemental Fury; Thundering Strikes; Ancestral Knowledge; Tidal Focus; Mindfulness; Totemic Focus.</li><li>**Totems:** no Wrath of Air or Totem of Wrath is trainable.</li><li>**Raid BiS:** Totem of Thunder 228176 (+1% Lightning Bolt crit); Ironfeather 2-piece.</li></ul> | engine against client rows | Aligned | - | - |
| E3 | The Searing Totem duration item belongs to lane core; not touched. | - | - | - | - |

## Shadow priest

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| S1 | **Inner Focus gave Shadow Word: Death +25% crit; the client does not.** The crit modifier (14751 effect 1, SpellEffect 692357, class mask word 1 0x00098434) leaves out Death's bit (1309636, word 1 0x2). The cost modifier (692356, word 1 0x10940436) includes it, so Death under Inner Focus is free but no likelier to crit. Latent: the rotation spends Inner Focus on Mind Blast. | `TestInnerFocusCritBonusSkipsShadowWordDeath` failed (25), then passed; `priest/talents.go` innerFocusAddsCrit | CONFIRMED-BUG (no published number) | old (the 1 October periodic rule, 2026-10; no round read the mask) | E ab73d2083 |
| S2 | Fresh, aligned with the client rows: <ul><li>**Spells:** Mind Blast 10947 (485 + 2.6/level, 0.429, 350 mana, 8 s − Improved Mind Blast); Mind Flay 18807 (130 a tick, 0.167, ticks crit: Attributes_8 4608 has 0x200, and it is on the periodic-crit list); Improved Mind Flay compounding (single source, as round 13 SH3); Shadow Word: Pain 10894 (8 ticks with Improved Shadow Word: Pain 2/2); Devouring Plague 19280 (every race, ticks crit); Shadow Word: Death 1309636 (448 + 2.5/level, 18 s, 30% health backlash); Vampiric Embrace; Dark Sacrifice; the wand (Torch of Light is Arcane, so Shadow mods skip it).</li><li>**Shadowform:** +10% Shadow, −50% cost, 2.0× crits.</li><li>**Talent masks:** Twin Disciplines, Mental Agility (not Mind Flay or Death), Shadow Focus 5% hit, Darkness, Shadow Weaving (2% × 5).</li><li>**Not learnable or not taken:** Shadowfiend (SkillLineAbility 49886, RaceMasks 0); Vampiric Touch (a rune); Power Infusion (0 in both builds).</li><li>**Raid BiS:** no set; Frozen Heart's +29 Shadow use; Mark of the Chosen needs melee hits taken (never fires).</li></ul> | engine against client rows | Aligned | - | - |
| S3 | `priest/mana_abilities.go:24` says Shadowfiend is "learned at level 1"; the client says it is not learnable. Nothing casts it. | SkillLineAbility 49886 | Aligned (comment only) | - | - |

## Cross-cutting

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | **Spell hit and partial resists.** Against a level-63 boss: 17% base miss with a 1% floor, so a 16% cap; a 6% average level-based partial resist on non-binary spells (0.02 × 3); none on binary spells. This is the royalgiraffe Era model. Forever has no row or log against it. | `core/target.go:470`, `core/spell_resistances.go:155` | Aligned (Era model) | - | - |
| X2 | **The published BiS talent strings are older than the guides:** <ul><li>Shadow raid (`325000031…`, against the guide's `00510503103`, about +0.2% by talent-search);</li><li>Shadow, Arcane and Frost bare 60;</li><li>Balance bare 60 (pre-round-13 B3);</li><li>Elemental bare 60 (pre-Tidal Focus);</li><li>Fire raid (Accelerant/Impact against Burning Soul; no DPS difference).</li></ul> Every BiS file was made on engine 8baffe5fc. | `bis/<spec>.json` generated 2026-11-01 | Measure at refresh | - | - |
| X3 | Lane core has not finished (no DONE file when this report was written), so the effect of the same-instant tie fix on these specs is not measured here. | - | - | - | - |

## Before/after

10,000 iterations, raid 60. Before is `$R/before.tsv`; after is `lane/r14-casters` eff1d0491 on `wt/r14-casters`.

| spec | horde s1 | horde s50001 | alliance s1 | alliance s50001 |
|---|---|---|---|---|
| mage-fire | 673.0 → 673.0 | 673.5 → 673.5 | 666.4 → 666.4 | 667.1 → 667.1 |
| priest-shadow | 691.8 → 691.8 | 691.6 → 691.6 | 684.2 → 684.2 | 684.9 → 684.9 |

Both fixes are latent. Nothing moves to the decimal, and the fork goldens (TestP1Mage, TestShadow) are unchanged. The other four specs have no code change.

## Files to regenerate

None.

## Measure at refresh

- Fire and every caster's haste weight on the current engine (M2; MR25). Fire should come out near 0, not −3.5.
- The stale BiS talent strings (X2; MR19, MR24).

## For other lanes

- **Healers:** Inner Focus's client crit mask (692357) also differs from the engine's rule for other spells. It names some Penance and Prayer of Mending rows and Desperate Prayer, and leaves out Renew. The engine excludes every channel. Check Penance against the mask.
- **Owner:** <ul><li>B3, the Balance raid build: Heart of the Wild 5 and Reflection 3 for Nature's Focus 5 and Subtlety 3 gain +1.6%. Both dropped talents are on the keep list.</li><li>E1, Elemental Mastery as one charge (−2.4% if so): an in-game check of whether the second Lightning Bolt inside Mastery is still 0.5 s faster.</li></ul>
- **Guides:** the Frost `expected_idle` reason for Cone of Cold (M9) is out of date.

## Still open

- E1 Elemental Mastery's charge (now sized at −2.4%).
- M8 cost-mod stacking, Master of Elements on free crits, Ice Lance's coefficient.
- M5 Kindling Stave's Ignition rate; M7 Arcane Concentration off wand shots.
- B4 raid caster distance; B5 Improved Moonfire on base damage, cost-cut stacking, Starfire rank 7.
- Warcraft Logs at 60 is blocked (gcloud login expired).
