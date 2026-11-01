# Sim audit round 6: hybrids (2026-11-02)

This report covers Druid (Balance, Feral cat, Feral bear, Restoration), Paladin (Holy, Protection, Retribution) and Shaman (Elemental, Enhancement, Restoration).

**Setup:**
- Client 1.60.1.70707. Published sets and builds are the 8baffe5fc nightly (site 21d97955f).
- Engine: pin 8baffe5fc plus fork branch `lane/r6-hybrids`. The branch also carries the pipeline lane's follow (E 6d8bb01a8, here f3e2c4453), as the integrator asked.
- Site: `wt/r6-hybrids` (this report only; every fix is in the fork).
- Sources: the 70707 shipped CSVs and the server's hotfix push (`scratchpad/r5/hotfix70707/`), Wowhead Forever tooltips (read 2026-11-01), TrinityCore master and cmangos classic master (`scratchpad/r5/hybrids/src/`, plus cmangos `HostileRefManager.cpp`, `ThreatManager.cpp` and `sql/base/mangos.sql` at 5731249195, read 2026-11-01), Warcraft Tavern's Classic threat reference table, and Magey's classic-warrior wiki "Threat Mechanics".
- The fresh re-audit of Balance, Feral cat, Restoration druid, Retribution, Holy paladin and the three shaman specs ran as three read-only sub-audits. I re-checked every row they called confirmed against the client rows before writing a test.

Fix-sha column: "E" is a fork commit on `lane/r6-hybrids`.

**In short:**
- **Protection paladin's +50% stands.** Righteous Fury alone is +50.7%. Two threat bugs found this round take 3.4% back off.
- **The bear was in Bear Form all along.** It was undertuned because Maul and Swipe lacked Classic's 1.75 threat. Fixed: +29.5% threat.
- **17 CONFIRMED-BUGs fixed in the fork:**
  - 2 from Item 1.
  - 1 from the bear item.
  - 1 routed by the integrator.
  - 13 from the fresh re-audit.
- One more confirmed data bug (P5) is routed to the data lane.

## Item 1: Protection paladin threat +50% (round 5's biggest move)

### Decomposition

Each round-5 change applied alone, or cumulatively, onto the f8a5f94ce engine, with today's site and simdb. Raid preset, band 60, 10,000 iterations, threat a second. Alliance (dwarf) / horde (undead), seed 1; seed 50001 is within 0.1%.

| change (fork commit) | alone | cumulative | TPS at seed 1 |
|---|---|---|---|
| f8a5f94ce (the before pin) | - | - | 686.0 / 692.4 |
| Righteous Fury Holy +60% → +90% only (f942519c3, half) | +13.3 / +13.0 | +13.3 / +13.0 | 777.0 / 782.6 |
| Righteous Fury +33% on all threat only (f942519c3, half) | +33.0 / +33.0 | | 912.3 / 920.9 |
| Both halves, multiplied (f942519c3) | +50.7 / +50.3 | +50.7 / +50.3 | 1033.5 / 1040.9 |
| Holy Shield mana 100/150/195 (7156faea2) | -0.1 / -0.2 | +50.5 / +50.1 | 1032.6 / 1039.2 |
| Healing gives no hidden damage third (b03f7e479) | -0.4 / -0.4 | +49.8 / +49.4 | 1027.9 / 1034.7 |
| Champion of the Light spell damage only (fb8acebd9) | 0.0 / 0.0 (with b03f7e479; the build does not take it) | +49.8 / +49.4 | 1027.9 / 1034.7 |
| Racials (e074b0415): Touch of the Grave every 15 s | 0.0 / +1.0 | +49.8 / +50.7 | 1027.9 / 1043.2 |
| The rest of the diff (Demonic Rune, Hunter's Mark, hunter, priest, mage) | 0.0 | +49.8 / +50.7 | = 8baffe5fc |
| Shield block value (site S 69e0fddf9, already in today's simdb) | +0.0 / -0.0 (shield reverted to its pre-fix block, before and after) | | 686.0 vs 686.0; 1027.9 vs 1027.2 |
| **This round:** Shield Specialization's mana makes no threat (E c979e0c74) | -2.6 / -2.5 | +46.0 / +46.9 | 1001.5 / 1017.0 |
| **This round:** mana gains take no threat multiplier (E f6232bd9d) | -0.9 / -0.9 | **+44.7 / +45.6** | 992.7 / 1008.2 |

- Shield block value adds nothing to a paladin's threat: no paladin ability reads block value.
- It reaches the published numbers already: d40f4d6a4 regenerated simdb.bin after the fix, and the 8baffe5fc nightly ranked on it.

### Every input, with two sources

| input | engine | source 1: 70707 client and push (with Wowhead, one source) | source 2 | verdict |
|---|---|---|---|---|
| Righteous Fury spell, cost, duration, level | 25780, 30% base mana, 30 min, level 10 | SpellPower 189813 PowerCostPct 30; SpellMisc DurationIndex 30 (1,800,000 ms); Wowhead "Requires level 10" | wowforevertalents, 70450 hotfix cache: "learned at level 10 instead of 16" | Aligned |
| Holy threat +90% | `righteous_fury.go:28` 0.90, a Holy-school spell mod | SpellEffect 703848 (aura 10, misc 2, 90); Wowhead "...Holy attacks by an additional 90%" | wowforevertalents 31 October: "...an additional 90%" | Aligned |
| All threat +33% | `righteous_fury.go:29`, the unit's threat multiplier | Push SpellEffect 1367626 (effect 4, aura 10, misc 127, 33); Wowhead text | wowforevertalents 31 October (DBCache, an independent capture of the same server data) | Aligned. The amount has only the server's data behind it; Blizzard published no note. |
| The two multiply (1.33 × 1.9) | unit multiplier × spell mod | Two aura 10 effects on one spell | TrinityCore `UpdateMySpellSchoolModifiers` (AddPct per effect); cmangos `HandleModThreat` (ApplyPercentModFloatVar) | Aligned. A launch threat log stays OPEN (round-5 T4). |
| Righteous Fury's other effects | not threat | Effect 1 aura 87 (damage taken, base 0, which Improved Righteous Fury feeds); effect 2 aura 108 -100% on Instrument of Law (push misc 12 → 23); effect 3 aura 668 | - | Aligned (no threat) |
| Improved Righteous Fury | 3/6/9% damage taken, no threat | 20468: aura 107 misc 12 on Righteous Fury's effect index 1 (the aura 87); talent text "all damage taken is reduced by 9%" | wowforevertalents: "...reduced by 6% 9%" | Aligned. Era's 16/33/50% threat talent is gone. |
| Iron Creed, the one Forever threat talent | +30% Holy Strike threat (`tank_talents.go:94`) | 1311034 effect 0: aura 108, misc 2 (threat), 30, mask_3 0x8000 (Holy Strike); talent text | wowforevertalents "15% 30%" | Aligned |
| Two-Handed Weapon Specialization | -20% threat with a two-hander | Push 1367246 (aura 10, -20) | wowforevertalents | Aligned (the build has a shield) |
| Holy Shield damage, levels, mana | 75/110/153 at 30/40/50; 100/150/195 | Push SpellEffect 698886, 699487, 697660; SpellLevels 103883, 104071, 103515; SpellPower 185892, 186383, 184913; Wowhead 20925/20927/20928 | wowforevertalents 31 October (221 → 153, 240 → 195 mana) | Aligned |
| Holy Shield coefficient, charges, block, cooldown | 0.08, 4 charges, +30% block, 10 s | Push coefficient 0.0800; SpellAuraOptions 131453 ProcCharges 4; effect 0 aura 51, 30; SpellCooldowns 55507 10,000 ms | Wowhead "4 charges... 30%" | Aligned |
| Holy Shield +20% threat | `holy_shield.go:48` 1.2 | Talent and tooltip text "causes 20% additional threat" | Warcraft Tavern Classic table: "40×1.2 (for damage caused)" | Aligned on the 1.2. cmangos instead adds a flat 20/30/40 per damage (`spell_threat` 20925/20927/20928). The flat part is OPEN; it is worth about +3% if Forever kept it. |
| 1239570, a second "Holy Shield" in the push (221 damage, "until canceled") | not registered | Push name and effects; shipped SkillLineAbility 53829. No SpellMisc, SpellLevels or cost, so it cannot be cast or trained. | Wowhead shows no level and no cost | Aligned (not learnable) |
| Block value on the shield | Evergreen Shield 279262: 44 | The client states none (no ItemSparse column; empty game table) | Wowhead feed `blockamount` 44 (one source) | OPEN, single source, as in round 5. It moves 0 threat for the paladin. |
| **Shield Specialization's mana** | was 0.5 threat a mana × 1.33 | 20150 aura 42 triggers 1310925 (effect 137, 4%). SpellMisc 857517 Attributes_1 **0x400 = SPELL_ATTR1_NO_THREAT**; the push leaves it alone. | TrinityCore `ForwardThreatForAssistingMe` and cmangos `threatAssist` both return early for NO_THREAT | **CONFIRMED-BUG**, fixed E c979e0c74 |
| **Mana gain threat multiplier** | was × the unit's multiplier (Righteous Fury 1.33) | 0.5 a mana (all sources agree) | Not multiplied: TrinityCore `EnergizeBySpell` (ignoreModifiers true); Magey "Power gains are not affected... by threat multipliers"; Warcraft Tavern "These values are not ever multiplied by modifiers". cmangos alone multiplies. The engine's own rage gain does not. | **CONFIRMED-BUG**, fixed E f6232bd9d |
| Seal of Fury, Judgement of Fury, Hammer of the Righteous, Consecration, Holy Strike | Holy school, threat multiplier 1 | SpellMisc SchoolMask 2 on every one; no aura 10 or threat spell mod in the shipped tables or the push targets them; Consecration 12 + 27 a second matches Wowhead's 96 and 216 | - | Aligned |
| Shieldsworn | Holy Strike -4 s; strikes trigger Holy Shield and Shield Specialization | Push 1367611 (aura 107, misc 11 cooldown, -4000, Holy Strike mask) | Wowhead text | Aligned |
| Blessing of Salvation | (raid presets never carry it on tanks) | Removed in Forever (see X3) | | Aligned |

### Does +50% stand?

Yes. Every round-5 component has client rows plus a second source, and Righteous Fury alone accounts for +50.7%. This round's two fixes take 3.4% back off. That leaves Protection paladin **+44.7% (alliance) / +45.6% (horde)** above the f8a5f94ce engine: 992.7 / 1008.2 threat a second.

**Plausibility against the other tanks** (same preset, today's engine):

| tank | threat/s (alliance / horde) | of the warrior's |
|---|---|---|
| Protection warrior | 1036.0 / 1024.6 | 100% |
| Protection paladin | 992.7 / 1008.2 | 96% / 98% |
| Feral bear (after G1) | 877.0 / 885.3 | 85% / 86% |

The one public Forever comparison I found is a MythicSim snapshot reported by dving.net (undated; MythicSim does not publish it itself). It has the paladin 4.2% (44 threat a second) behind the warrior, which matches ours, and the bear far ahead (1569 against 1096). Its bear number follows its own Lacerate and Primal Bite model (G3). Tier lists (mobalytics, warcrafttavern, rankedboost, read 2026-11-01) rate the paladin the strongest raid tank, without numbers. No source puts a Forever paladin 50% below a warrior any longer.

## Item 2: owner ruling 3

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| O3 | Restoration druid's lead over the other healers is left alone, blocked on launch logs (owner ruling 3). Not reopened. R1 below (Genesis on Tranquility, -0.8%) is a spell-mask fix, not a profile or ranking change. | Brief, owner rulings | Recorded | - |

## Item 3: post-nightly checks

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| N-1 | Retribution horde is off The Needler. Band 60 raid, both factions: Arcanite Champion (12790); bare: Legionite Glaive. | `data/builds/1.60.1.70707/bis/paladin-retribution.json` (engine 8baffe5fc, 2026-11-01T20:45Z) | Aligned | - |
| N-2 | Healing-stat pieces left in hybrid sets all state spell damage of their own and won their slot at verify on 8baffe5fc. Examples: Living Crown (Balance head, 25 spell damage), Gloves of the Greatfather (Balance and Elemental hands, +1.5 / +2.5 over the runner-up), Hand of Edward the Odd (90 / 90). Enhancement and Retribution raid sets carry none. | simdb.bin stats against the published slots and their verified alternatives | Aligned | - |

## Cross-cutting

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| X1 | Mana gains made threat times the unit's threat multiplier (Item 1). Threat only, every class with a mana bar. | `sim/core/mana.go:139` | CONFIRMED-BUG (fresh, three sources against one) | E f6232bd9d |
| X2 | `ResourceMetrics.NoThreat` now marks a NO_THREAT energize; the mana and rage gain threat skip it. Only Shield Specialization sets it. I checked Judgement of Wisdom (20268/20352/20353), Furor, Blood Frenzy, Bloodrage, Enrage, Unbridled Wrath and Master of Defense: none carries 0x400. | SpellMisc Attributes_1 | (mechanism for T1) | E c979e0c74 |
| X3 | **Integrator routing (pipeline lane P5, P6): Blessing of Light, Greater Blessing of Light, Salvation and Greater Salvation are removed in Forever.** The engine registered all of them, and the fork's Holy APL copy precast Greater Blessing of Light. The engine side and the fork APL copy are removed. The site's `paladin-holy.json` precast is already gone on `wt/r6-healers` (healers lane). No preset carries either blessing (`data/curated/presets.json`). | Push SpellName status 2 (19977, 19978, 25895 in the 10:59 dump; 19979, 25890 and 1038 per the pipeline lane's newer dump); Wowhead 404 for 19979, 25890, 1038; wowforevertalents 31 October: "Removed Revoke, Salvation, Mass Salvation, Blessing of Light, Greater Blessing of Light" | CONFIRMED-BUG (routed) | E 5b79f29fd, 25ea37277 |
| X4 | Pipeline P7 to P9 for this group: Holy Shield and Shield Specialization proc masks gain taken spells (only Shieldsworn's 6 s spell block could use it; the boss swings in melee, so 0). Echo of Fury and Righteousness stack to 3 with no charges (Unsealed Wrath already stacks to 3). Templar's Bulwark at 30 (`defensive_cooldowns.go:29` already 30). | Pipeline lane report | Aligned | - |
| X5 | The pipeline follow E 6d8bb01a8 is cherry-picked (f3e2c4453). Fork suite green with it. | | (carried) | E f3e2c4453 |

## Druid: Balance

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| B1 | **Moonkin crit counted twice in the raid preset.** The form added 3%, and the preset's `moonkin_aura` added a static 3% (`core/buffs.go:250`). The form is one aura (24907, aura 290, 3); 24858 itself has no crit effect, and an aura does not stack with itself. The form now adds its 3% only when the raid buffs do not already carry it. | SpellEffect rows of 24858 and 24907. Test first: `TestMoonkinCritIsGrantedOnce` (0.03 extra with the raid aura up, then 0). | CONFIRMED-BUG (fresh, client rows) | E 3ca90b1e4 |
| B2 | **Vengeance skipped Insect Swarm.** 16909 effect 0 (aura 108, misc 15) mask_0 0x600207 includes 0x200000, Insect Swarm (SpellClassOptions 24977); the text says "Arcane and Nature spells". | `talents.go:497`. Test first: `TestVengeanceRaisesInsectSwarmCrits` (0, then 1.0). | CONFIRMED-BUG (fresh, mask plus text) | E 90d8ca056 |
| B3 | **Nature's Reach's 4% hit was limited to Wrath, Starfire and Moonfire.** 16819's hit is effects 1 and 2 (aura 54 and 55, 4 each) with no class mask; only the range effect names spells. Wowhead's text is unscoped. | Test first: `TestNaturesReachAddsHitToEverySpell` (Insect Swarm 0, then 4). | CONFIRMED-BUG (fresh, client row; Wowhead agrees) | E 9af6d73ac |
| B4 | Moonglow and Improved Wrath's cost cuts add (Wrath 25%); a multiplying server gives 37.5%. Out of mana 0.1 s of 180, so no published effect. | 16845, 16814 | OPEN (a logged Wrath cost settles it) | - |
| B5 | Checked and aligned: Moonfire, Wrath, Starfire (25298 is an Ahn'Qiraj rank), Insect Swarm amounts, coefficients, costs and ticks; Improved Wrath, Improved Moonfire, Nature's Majesty, Nature's Splendor, Improved Starfire, Nature's Grace, Eclipse, Moonfury, Naturalist, Genesis; Idol of the Moon, Frozen Heart of the Mountain, Ironfeather 2-piece; DoT crit flags. | Sub-audit, client rows | Aligned | - |

## Druid: Feral (cat)

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| F1 | **Natural Instinct's +20% crit damage missed Rip.** 1223242 effect 0 (aura 108, misc 15) mask_0 0x839800 includes 0x800000, Rip (SpellClassOptions 9896). Bleeds always crit (owner ruling). | `druid.go:112`, `talents.go:709`. Test first: `TestNaturalInstinctRaisesRipCrits` (0, then 0.20). | CONFIRMED-BUG (fresh, client mask; the server's periodic crit through misc 15 follows TrinityCore's tick path) | E 6e2647d84 |
| F2 | Checked and aligned: Shred, Ravage, Rip, Ferocious Bite (the pushed rank 1 arrives with the pipeline follow), Claw, Shifting Power; Ferocity, Savage Fury, Sharpened Claws, Predatory Strikes, Blood Frenzy, Leader of the Pack (one raid flag, not doubled), Rend and Tear, Predatory Discipline, Genesis on Rip; Idol of the Dream, Molten Heart of the Mountain, Shadowcraft 5-piece. | Sub-audit | Aligned | - |

## Druid: Feral (bear)

**The bear was in Dire Bear Form for the whole fight. It was undertuned because Maul and Swipe made too little threat: Maul made plain damage threat where Classic gives 1.75 times, and Swipe had a Season of Discovery multiplier no Forever druid has. With both fixed, bear threat is +29.5% (alliance) / +29.7% (horde), 85 to 86% of a Protection warrior's.**

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| G0 | **Form check (integrator's item).** All of these hold:<br>- The tank starts in Dire Bear Form (BuildPhase base): aura uptime 100.1%, one gain, no drops or shifts in the fight, so Furor never rolls.<br>- Paw: 2.5 s swings at the equipped weapon's DPS (Twig of the World Tree, 49.7 DPS), the deep dive's rule (Blizzard, 2026-09-30, `forms.go:30`).<br>- Form numbers at 60: +180 attack power and +1240 health (9635: 120 + 3 a level and 600 + 32 a level from 40). Wowhead's 210 / 1560 are rendered at MaxLevel 70.<br>- Armor: +360% of item armor (9635 aura 142 and 466). Item armor 1028 × 4.6, plus 2 × Agility, plus Thick Hide 180 × 4.6, plus 385 bonus armor, gives 6318 bare. The engine shows 6318.<br>- Threat: +50% (21178 aura 10, 50).<br>- Defense 17 from items, so the 15% crushing blow chance is right for a bear.<br>- Forever has no Savage Defense: no trainable or talent.<br>- Demoralizing Roar is not in the rotation.<br>- Lacerate is the talent (1235827, 15 rage, 75 over 15 s, Wowhead agrees). | `bear_form.go`, `forms.go`; SpellEffect 687085/6/7, 1231421, 700189; spec-breakdown `-tank` | Aligned | - |
| G1 | **Maul made plain damage threat, and Swipe twice its damage.** Classic gives both 1.75 times their damage, and no client row states it. Swipe's 2.0 was Season of Discovery's tuning passive 436895 (+100% on Swipe's mask). No 70707 SkillLineAbility row teaches that passive, and the engine applies none of its other effects. Maul was 41% of the bear's threat. | Warcraft Tavern Classic table ("Maul x1.75", "Swipe x1.75"); wowsims' Era-based SoD engine (`sim/druid/maul.go` ThreatMultiplier 1.75; Swipe 3.5 = 1.75 × the SoD passive). cmangos classic gives Maul a flat +207 (9881) instead, so the form is disputed but the extra threat is not. Test first: `TestMaulAndSwipeCarryClassicThreat` (1 and 2, then 1.75 and 1.75). | CONFIRMED-BUG (two sources; owner item) | E 15bb5b66f |
| G2 | Bear damage taken is 611 a second against the warrior's 357, and the chance of death in the tank fight is 99%. Same armor as the warrior (6722 against 7327), no block (the warrior blocks 59% of swings) and 15% crushing blows. That is Era's bear trade, not a form bug. | spec-breakdown `-tank` | Aligned (no bug found) | - |
| G3 | Lacerate's and Primal Bite's "high amount of threat" is the engine's unconfirmed 2.0 (`threat.go:12`). Its comment derives it from the same SoD passive, which Forever does not teach. No client number exists; TBC's Lacerate used flat threat. | Client text only | OPEN. A launch threat log, or Blizzard's number, settles it. | - |
| G4 | Demoralizing Roar's flat threat is 2 × the rank's level (`demoralizing_roar.go:76`). cmangos gives 9898 a flat 42 and Warcraft Tavern 39. Not cast by the rotation. | | OPEN (no published effect) | - |

## Druid: Restoration (spell amounts only; builds and APLs belong to the healers lane)

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| R1 | **Genesis raised Tranquility.** 1223081 effect 0 (aura 108, misc 22) mask_0 0x220252 names Moonfire, Rejuvenation, Regrowth, Pounce and Insect Swarm (with Wild Growth and Lacerate in mask_1, Rip in mask_2). Tranquility is 0x80 and is not in it. | `talents.go:603`. Test first: `TestGenesisAddsToTheHealsOverTimeItsMaskNames` (1.1, then 1.0). | CONFIRMED-BUG (fresh, client mask) | E b28eb9490 |
| R2 | Genesis effect 1 and Gift of Nature name Swiftmend; whether the server builds Swiftmend from the raw or the modified heal-over-time amount is unread. | | OPEN (tiny) | - |
| R3 | Checked and aligned: Healing Touch, Rejuvenation, Regrowth, Tranquility (Improved Tranquility, crit flag), Swiftmend, Nature's Swiftness, Innervate, Bloom; Gift of Nature, Tranquil Spirit, Improved Rejuvenation, Living Spirit, Reflection. | Sub-audit | Aligned | - |

## Paladin: Protection

The Item 1 tables are this spec's audit. Two more rows:

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| T1 | Shield Specialization's mana made threat (Item 1). | 1310925 SpellMisc 857517 Attributes_1 0x400; TrinityCore and cmangos. Test first: `TestShieldSpecializationManaAddsNoThreat` (2446, then 0). | CONFIRMED-BUG (fresh) | E c979e0c74 |
| T2 | Mana gain threat took Righteous Fury's +33% (X1). Test first: `TestManaGainThreatIgnoresRighteousFury` (2394 for 3600 mana, then 1800). | | CONFIRMED-BUG (fresh) | E f6232bd9d |

The tank talent search was not re-run. Tanks have no raid keeps, and the tank score orders effective health and risk before threat, so a threat-only change cannot flip the build.

## Paladin: Retribution

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| P1 | **Judgement of Righteousness, and Judgement of Command's hit check, rolled the spell table.** The 70707 client moved both to the melee table with no active defense:<br>- SpellCategories DefenseType 2 on 20286 (and its ranks) and 20968; Era 1.15.9 has 1.<br>- SpellMisc Attributes_0 gains 0x200000, SPELL_ATTR0_NO_ACTIVE_DEFENSE (20286 0x50000 → 0x250000; 20968 329728 → 2426880).<br>The engine missed about 9% of Judgements and crit Judgement of Righteousness for 1.5x. | `sor.go`, `soc.go`. Test first: `TestJudgementsRollTheMeleeTableWithoutActiveDefenses` (magic, then melee). | CONFIRMED-BUG (fresh, client tables; the Era comparison shows the change was deliberate) | E c7a9f1fd1 |
| P2 | **Hammer of Wrath took Precision's hit back.** Precision 20189 is aura 54 (melee and ranged) plus a new aura 55 (spell), with no SpellEquippedItems row; Wowhead: "Improves your chance to hit by 1%". | `hammer_of_wrath.go:76`. Test first: `TestHammerOfWrathKeepsPrecisionsHit` (-3, then 0). | CONFIRMED-BUG (fresh, two sources) | E 4a2abc879 |
| P3 | **Benediction discounted Seals only.** 20101's mask (0x1be11fbf / 0x26018242 / 0x20212100 / 0x8800) also holds Judgement 0x800000, Consecration 0x20, Exorcism mask_2 0x2000, Holy Strike mask_3 0x8000 and Holy Shock 0x200000, but not Hammer of Wrath (mask_1 0x80); the text says "all instant cast spells and abilities". | Test first: `TestBenedictionDiscountsEveryInstantItsMaskNames` (100%, then 90%). | CONFIRMED-BUG (fresh, mask plus text) | E a3e569f13 |
| P4 | Arcanite Champion's proc rate is a guessed 1 PPM. It fires from both Seals' procs and Judgement of Command; 70707's Seal of Righteousness damage spell 25713 lacks Attributes_3 0x200, which 20424 and 20966 carry. | `common/item_effects.go:220` | OPEN (a level-60 proc log settles it) | - |
| P5 | Rune of the Guard Captain counts its Forest/Grassland tripling as a flat +28 (42 attack power where the client gives 14). | ItemXItemEffect 101515 → 1318000 (+14, "tripled in Forest and Grassland"); 237884 → 1287704; Wowhead item 19120 | CONFIRMED (data pipeline, two sources). **Routed to the data/ranker lane**; not fixed here. | - |
| P6 | 25713 carries base points 35 and no coefficient; the damage is server-scripted from 20293. | | OPEN (low) | - |
| P7 | Consecration's damage comes from an area trigger as direct damage, so its crits may count toward Vengeance; Vengeance is already up 99.3%. | 1280349 | OPEN (negligible) | - |
| P8 | Holy Strike's flat bonus under the weapon percent (round 5 P3): no new evidence. | | OPEN (carried) | - |
| P9 | Checked and aligned: Seal of Command (70% weapon, 0.29, 210 mana, 1 s ICD, swing-only proc), Seal of Righteousness, Judgement (6%, 8 s with Improved Judgement), Judgement of Command and Righteousness amounts, Holy Strike, Hammer of Wrath, Consecration; Improved Seals, Holy Conduit, Divine Strength, Divine Intellect, Precision, Conviction, Sanctified Judgement, Sacred Arbiter, Two-Handed Weapon Specialization, Vengeance, Champion of the Light, Unsealed Wrath; Ruthless, Touch of the Grave; Molten and Frozen Heart of the Mountain, Lightforge 2-piece, Cloak of Flames. | Sub-audit | Aligned | - |

## Paladin: Holy (spell amounts only)

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| L1 | **Illumination's 20272 also gives 15 s of 50% regeneration while casting**, new in Forever. The engine gave the mana only. | SpellEffect 1237075 (aura 134, 50), SpellMisc DurationIndex 8 (15 s); Wowhead buff "Regenerating at 50% of normal rate while casting spells. 15 seconds remaining"; Era 20272 has only the energize. Test first: `TestIlluminationOpensFifteenSecondsOfCastingRegen` (0 procs, then 1). | CONFIRMED-BUG (fresh, client plus Wowhead) | E 26088c6c6 |
| L2 | Benediction on Holy Shock (P3). | | CONFIRMED-BUG (P3) | E a3e569f13 |
| L3 | Blessing of Light is gone (X3). The Holy rotation's Greater Blessing of Light precast falls with it. | | CONFIRMED-BUG (X3) | E 5b79f29fd |
| L4 | Holy Power's crit-heal bonus: the push turns 5923 into aura 50 and removes the 15% crit effect (1340809). The engine adds +50% to the crit bonus (×1.75). TrinityCore's `SpellCriticalHealingBonus` would multiply the whole crit (×2.25). The talent text supports the engine. | Push 683158, 1340809 | OPEN. One in-game crit against a non-crit Flash of Light settles it (about +12% raw HPS if multiplied). | - |
| L5 | Checked and aligned: Holy Light 1 to 9, Flash of Light 1 to 6, Holy Shock, Light's Vigil, Divine Favor; Healing Light, Infusion of Light, Boundless Light, Illumination's mana half. | Sub-audit | Aligned | - |

## Shaman: Elemental

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| E1 | Elemental Mastery (16166) carries ProcCharges 1 with proc mask 87376, the shape of Clearcasting. The engine keeps it for its full 15 s; the text says "Lasts 15 sec". | SpellAuraOptions 128637; `talents.go:327` | OPEN, single source. Cast Elemental Mastery and then two Lightning Bolts on beta: if the second is not instant, the charge ends it (-1.7 to -1.9%). | - |
| E2 | Orc Command (21563) has a class mask that includes shaman, but its effect row is not in the shipped table, and the engine applies it to hunter and warlock pets only. If it reaches Searing Totem, that is at most +0.3%. | Push SkillLineAbility 59881 | OPEN | - |
| E3 | Checked and aligned:<br>- Talents: Convection, Concussion, Elemental Alacrity, Call of Flame, Elemental Focus, Lightning Overload, Elemental Fury, Elemental Devastation.<br>- Spells: Lightning Bolt, Lava Burst, Flame Shock, Searing Totem, Fire Nova.<br>- Gear and totems: Mana Spring, Totem of Thunder, Frozen Heart of the Mountain. | Sub-audit | Aligned | - |

## Shaman: Enhancement

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| N1 | **Windfury Weapon took Elemental Weapons twice** (333 × 1.4 × 1.4 = 652.7, where the client gives 466.2). | 16266 effect 2: +13/27/40 on 16361's attack power (mask 0x800000). Text: "Windfury Weapon effect by 40%". Test first: `TestWindfuryWeaponAttackPowerTakesElementalWeaponsOnce` (652.68, then 466.2). | CONFIRMED-BUG (fresh, client plus text) | E ad38419e5 |
| N2 | **Stormstrike was Era's debuff**: +20% to all Nature damage taken, with two charges nothing ever set, so it never ran out. Forever's is one charge on the shaman's next Lightning Bolt, Chain Lightning or Earth Shock. The shaman now owns a single-charge debuff; `core.StormstrikeAura` (the raid-debuff option, in no preset) is untouched. | 17364 effect 1: aura 271, +20, mask 0x100003; SpellAuraOptions 129538 ProcCharges 1; Wowhead "your next Lightning Bolt, Chain Lightning, or Earth Shock spell by 20% for 12 sec". Test first: `TestStormstrikeRaisesOnlyTheNextShock` (Nature taken 1.32 and a second Earth Shock at 1.2, then 1.1 and 1.0). | CONFIRMED-BUG (fresh, client plus Wowhead) | E a2dc3e757 |
| N3 | Flurry's buff never expired; the client's is 15 s (DurationIndex 8). No measurable effect. | Test: `TestFlurryBuffLastsFifteenSeconds` (failed to build without the constant) | CONFIRMED-BUG (fresh, negligible) | E d14e24b46 |
| N4 | Windfury Weapon's extra swings are yellow specials (no glancing, no Flurry charge spend). 16361 is effect 19 (extra attacks), the shape of Windfury Totem, which the engine swings white. Worth -2 to -2.5% more if white. | `windfury_weapon.go:103` TODO | OPEN. A beta log showing whether Windfury swings glance settles it. | - |
| N5 | The APL drops its own Windfury Totem, which a Windfury Weapon shaman cannot use: one GCD and 250 mana a fight. | `data/curated/apl/shaman-enhancement.json` | Rotation note (no bug) | - |
| N6 | Improved Stormstrike's 50% casting regen is modelled; its cooldown reset on dodge or parry never fires, because the boss attacks the tank. | `talents.go:703` | Aligned | - |

## Shaman: Restoration (spell amounts only)

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| H1 | **Purification raised Healing Stream.** 16178's masks (448; mask_2 16) hold Healing Wave, Lesser Healing Wave, Chain Heal and Riptide; Healing Stream is 0x2000, which only Restorative Totems names. | `water_totems.go:77`. Test first: `TestRestorativeTotemsRaiseHealingStreamAndPurificationDoesNot` (31.68, then 29.70). | CONFIRMED-BUG (fresh, client masks) | E 00788a21c |
| H2 | **Route to healers lane:** the APL casts Mana Tide rank 3 (17359). Rank 4 (1325256, level 58, 290 a tick, SkillLineAbility 59647) is learnable and registered. | `data/curated/apl/shaman-restoration.json` | Route | - |
| H3 | Tauren Endurance's aura 50 row 1349372 (+2%) survives the push (round-5 open). | | OPEN (carried) | - |

## Before/after

Raid preset, band 60, 10,000 iterations, seeds 1 / 50001.
- **Before:** pin 8baffe5fc.
- **After:** the lane at 5b79f29fd, with `wt/r6-hybrids`. The two later commits move no number: d14e24b46 (Flurry; goldens unchanged) and 25ea37277 (a conformance golden).
- Healers are scored in effective HPS on the heal profile, tanks in threat a second.

| spec | changes | alliance % | horde % | absolute, seed 1 (alliance; horde) |
|---|---|---|---|---|
| Balance | B1, B2, B3 (they net out) | -0.1 / -0.2 | -0.3 / -0.3 | 537.2 → 536.8; 540.3 → 538.6 |
| Feral cat | F1 | +1.4 / +1.4 | +1.3 / +1.3 | 615.6 → 624.1; 616.5 → 624.7 |
| Feral bear (threat) | G1 | +29.5 / +29.5 | +29.7 / +29.6 | 677.2 → 877.0; 682.8 → 885.3 |
| Restoration druid | R1 | -0.8 / -0.8 | -0.8 / -0.8 | 741.5 → 735.4; 739.2 → 733.1 |
| Holy paladin | L1, P3, X3 | +4.5 / +4.5 | +3.7 / +3.9 | 553.8 → 578.7; 544.2 → 564.5 |
| Protection paladin (threat) | T1, T2 | -3.4 / -3.4 | -3.4 / -3.4 | 1027.9 → 992.7; 1043.2 → 1008.2 |
| Retribution | P1, P2, P3 | +2.1 / +2.1 | +1.6 / +1.7 | 652.6 → 666.2; 658.8 → 669.3 |
| Elemental | - | 0.0 | 0.0 | 485.9; 494.6 |
| Enhancement | N1, N2 | -3.0 / -3.0 | -2.9 / -2.9 | 630.0 → 611.3; 634.3 → 615.7 |
| Restoration shaman | H1 | +0.1 / +0.0 | +0.1 / +0.1 | 641.9 → 642.5; 640.2 → 641.1 |

**Outside the group:**
- f6232bd9d moves threat for every mana user, never damage or healing; TestP1Mage's golden moved on two threat rows. No other class's published number is threat.
- The pipeline follow (f3e2c4453) is the pipeline lane's, and that lane reports it.

**Not done:**
- Nothing nightly-owned was committed.
- `TestRotationLadder` was not regenerated.

**Fork goldens regenerated:** TestProtection, TestFeralTank, TestEnhancement, TestP1Mage (threat only), and the paladin conformance golden and SUMMARY (the removed blessings). The full fork suite passes with `-tags=with_db`, except `sim/web`, which needs `binary_dist` as before.

## Still open

1. **Holy Shield's flat Era threat** (cmangos +20/30/40 per damage on top of the 1.2): about +3% paladin threat if Forever kept it. Settled by a launch threat log.
2. **Righteous Fury's multiplied threat** (round-5 T4) and the +33% amount: server data and Wowhead only. Settled by a launch threat log.
3. **G3.** Lacerate and Primal Bite's "high threat" multiplier. **G4.** Demoralizing Roar's flat threat.
4. **E1.** Elemental Mastery one charge or 15 s (-1.8% if one charge).
5. **N4.** Windfury Weapon's swings: white or yellow (-2 to -2.5% if white).
6. **L4.** Holy Power ×1.75 or ×2.25.
7. **P4.** Arcanite Champion PPM. **P6.** 25713 base points. **P8.** Holy Strike's flat bonus. **R2.** Swiftmend and Genesis. **B4.** Moonglow and Improved Wrath stacking. **E2.** Command on totems. **H3.** Tauren 1349372.
8. **Routed:**
   - P5, Rune of the Guard Captain's biome tripling, to the data and ranker lane.
   - H2, Mana Tide rank 4 for Restoration shaman, to the healers lane.
   - N5, Enhancement's own Windfury Totem, a rotation note.
9. **Block value** remains single-source (Wowhead's feed); it moves 0 paladin threat.
