# Sim audit round 8: hybrids (2026-11-02)

This report covers Feral cat, Feral bear (threat), Retribution, Protection paladin (threat) and Enhancement.

**Setup:**
- Client 1.60.1.70707. Client rows are the layered 70707 tables (hotfix push included) from the 2026-11-01 21:42Z fetch, plus SpellCategories from `scratchpad/wago/` (2026-10-31).
- Published sets and builds are the committed `data/builds/1.60.1.70707/bis/` (8baffe5fc nightly). No ranker run was made; nothing nightly-owned was committed.
- Engine: pin cd4dfdcc0 plus `lane/r8-hybrids`. Site: `wt/r8-hybrids` on origin/main 170b7b84c.
- Second sources: Wowhead Forever tooltips (read 2026-11-01), Blizzard's 15 October beta tuning notes as reprinted by wow.gg (https://wow.gg/guides/wow-forever-beta-tuning-15-10-26, read 2026-11-01), Warcraft Tavern's Forever feral guide (https://www.warcrafttavern.com/forever/guides/feral-druid/, read 2026-11-01).
- Method, besides reading code: a rank sweep that ran every spec at bands 20 to 60 (bare, and raid at 60) and compared every cast and every aura with uptime to the class trainables' top rank at that level (`scratchpad/r8/hybrids/rankaudit.py`), and an item sweep that listed every on-use, chance-on-hit and proc effect on every published set item, with the engine files that name it (`items.py`, `items_all.py`).

Fix-sha column: "E" is a fork commit on `lane/r8-hybrids`, "S" a site commit on `wt/r8-hybrids`.

**In short: 3 CONFIRMED-BUGs, all fixed.** None moves a published raid number at 60. They move the bare preset (every band below 60, and the bare 60 entry):
- Shred rank 1 is learned at 20, not 22.
- Stormstrike normalizes its weapon damage.
- Cat and bear can cast Faerie Fire in form, free on a 6 s cooldown; the engine could not, so a bare cat or bear never had its own armor reduction.

## Item 1: side effects of round 7's hybrid fixes

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| S1 | Strength of Earth 42 / Grace of Air 77: the raid preset's totem and Enhancement's own totem share one aura (`core.StrengthOfEarthTotemAura`, GetOrRegister), so they never stack (2 activations, one aura, in the Enhancement breakdown). The raid preset turns the buff on at its first value (TristateEffect Regular, multiplier 1); Forever has no Enhancing Totems, so the 1.15 "improved" value is never reached (`sim/request/buffs.go onEnumValue`). Band ranks: the rank sweep shows 8075 at 20, 8160 at 30, 8161 at 40 and 50, 10442 at 60, each the top non-book rank. | `core/buffs.go:327`, `shaman/earth_totems.go`; rank sweep | Aligned | - |
| S2 | Mark of the Chosen: 21 stats on its 2-minute proc cooldown, worn by the bear at 50 and 60 and Protection at 50 and 60; the item sweep finds it handled in `common/item_effects.go`. | item sweep | Aligned | - |
| S3 | Judgement of Fury on the melee table: Protection's breakdown shows it at 4.2% miss with no dodge or parry, and Judgement of Command (Retribution) likewise. | breakdowns, raid 60 | Aligned | - |
| S4 | Iron Creed opens only on a landed Holy Strike (69% uptime from 28 Holy Strikes, 19.6 landed, 6 s each). | breakdown | Aligned | - |
| S5 | The shared immunity cooldown: Divine Shield 0.01 casts and Templar's Bulwark 0.23 (Category 37 vs 0) behave as round 7 left them. | breakdown | Aligned | - |
| S6 | Hammer of the Righteous at 4 targets: single target unchanged (27.4 casts on its 6 s cooldown, 594 average). | breakdown | Aligned | - |
| S7 | Held items replacing the main hand: none of the five specs publishes a held item; Enhancement's band-60 two-hander leaves the off hand empty, as round 7 N3 said. | bis json | Aligned | - |
| S8 | Feral attack power on items: Cat Form's attack power is the client's 3025 effect 0 (12 + 2 a level from level 6, which is 2 x level), plus Agility and the items' feral attack power through `stats.FeralAttackPower` (`forms.go:registerCatFormSpell`). | client 3025, code read | Aligned | - |

## Druid: Feral (cat)

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| F1 | **Shred rank 1 was learned at 22. The client teaches it at 20.**<br>- What the client says: Shred 5221 has BaseLevel 20 (SpellLevels 98857) and a SkillLineAbility row on Feral Combat. Wowhead Forever 5221 reads "Requires level 20". Blizzard's 15 October notes: "Rank 1 Shred is learned at level 20 instead of 22."<br>- What was wrong: the engine's `shredLearnLevels` kept Era's 22, and the site's spellranks held 5221 at 22 to match, so a band-20 cat had no Shred.<br>- What changed: Shred rank 1 registers at 20 and the site override is gone. Rank 1 Shred (155% plus 22 for 60 energy) does less than rank 1 Claw on a level-20 weapon, so the rotation keeps Claw until Ferocious Bite's first rank (24) and Shreds from there. Claw first measured about 4.6% ahead at band 20; Shred is about 5% ahead at 30. | `druid/shred.go:12`; `sim/internal/spellranks/spellranks.go rankLevelOverrides`. Tests first: `TestShredRankOneIsLearnedAtTwenty` in the fork (rank 0 at 20, then 1) and in the site's spellranks (unlearned at 20, then 5221). | CONFIRMED-BUG (client rows, Wowhead, Blizzard notes) | E 47b00171b, S 1c8452558 |
| F2 | **A cat could not cast Faerie Fire.**<br>- What the client says: Cat Form (Passive) 3025 effects 2 to 4 take 100% off Faerie Fire's cost (aura 108, op 14, mask 0x400), add 6000 ms to its cooldown (aura 107, op 11) and take 500 ms off its global cooldown (op 21). Blizzard's 15 October notes: "Omen of Clarity no longer procs from a free Faerie Fire cast while using Cat Form, Bear Form, or Dire Bear Form". Warcraft Tavern's Forever cat rotation: "Use Faerie Fire to maintain the debuff."<br>- What was wrong: the engine registered only the caster spell, rank 4 alone, castable in caster and Moonkin form. A bare cat (every band below 60, and bare 60) never had an armor reduction.<br>- What changed: a form spell (`FaerieFireFeral`, tag 1) at the level's rank (175/285/395/505 armor, client 770/778/9749/9907 effect 0), free, 6 s cooldown, 1 s global cooldown in Cat Form, no swing reset, no proc mask (so no Omen roll). Rank 4 shares the raid preset's aura, so the raid number at 60 does not move. The rotation casts it when the target has no Faerie Fire. | `druid/faerie_fire.go`. Test first: `TestFaerieFireIsAFreeSixSecondCastInCatForm` (build failure: no such spell; then rank 2 at 30, no mana, 285 armor, 6 s and 1 s). | CONFIRMED-BUG (client rows, Blizzard notes, Warcraft Tavern) | E 528467049, E 9340e9743, S 1c8452558 (cat line) |
| F3 | Ferocious Bite casts rank 5 (22829) at 60; rank 6 (31018) is the Ahn'Qiraj book, held back by IncludeAQ. Forever adds rank 1 at 24 (1328618), and the ranks shift (22568 at 32 ... 22829 at 56); the engine's ladder is the generated one and matches. | rank sweep; trainables | Aligned (book rank: round 7 F2, still open) | - |
| F4 | Checked and aligned, fresh against the 70707 rows:<br>- Spells: Claw ranks 20/28/38/48/58, Rip 20 to 60 (12 s, 2 s ticks, 15 + 25.5 a point at rank 6), Ravage 32/42/50/58, Prowl 20/40/60, Shred 155% plus 22/32/44/64/80, Ferocious Bite 35 energy, Shifting Power.<br>- Cat Form: -29% threat (3025 effect 1, the engine's 0.71), attack power 2 a level.<br>- Talents, against the client masks: Natural Instinct (0x839800 names Claw, Rake, Shred, Ravage, Pounce, Ferocious Bite, Rip, Maul, Swipe, Lacerate and Primal Bite; the engine's mask matches), Savage Fury (0x9800), Shredding Attacks (Shred -6 a rank), Ferocity (Maul, Swipe, Primal Bite, Claw and Rake only), Genesis (Rip and Rake are in its mask), Nature's Majesty, Sharpened Claws, Predatory Strikes, Blood Frenzy, Rend and Tear (2% a rank, direct damage only), Naturalist (2% a rank on all damage), Heart of the Wild (Strength in Cat Form), Leader of the Pack (one raid flag, never doubled with the raid preset's).<br>- Gear: Idol of the Dream (Improved Rip 446212), Shadowcraft 5-piece energy proc and the Defias 4-piece (both in `client_dungeon_sets.go`), Haste (Manual Crowd Pummeler), Chthonic Power, Molten Fury. Shadowcraft has no class restriction in the 70707 ItemSparse rows (AllowableClass -1), so a druid may wear it. | rank sweep, client rows, code read | Aligned | - |

## Druid: Feral (bear)

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| G1 | **A bear could not cast Faerie Fire either.** Bear Form (Passive) 1178 effects 4 and 5 make it free with a 6000 ms cooldown (no global cooldown change, so 1.5 s); Wowhead Forever's 1178 lists "Modifies Power Cost -99%" and "Modifies Cooldown +6001" on Faerie Fire. Blizzard's 15 October notes: "Faerie Fire now generates threat when used in Bear Form and Dire Bear Form." Warcraft Tavern's Forever bear rotation: "Use Faerie Fire to maintain the debuff or to pull from range." The bear now casts it (same form spell, 1.5 s global cooldown, Era's 108 flat threat) when the target has none. The raid preset's Faerie Fire covers the raid number. | `druid/faerie_fire.go`; the cat test above covers the spell; TestFeralTank gains only the new spell's rows. | CONFIRMED-BUG (client rows, Blizzard notes, Warcraft Tavern) | E 528467049, E 9340e9743, S a0ac4c31c |
| G2 | The Faerie Fire threat amount (108, Era's Faerie Fire (Feral)) is not in the client tables (effect 3 is Effect 63 with no points). | 9907 effect 3 | OPEN (a beta log of a bear's Faerie Fire threat; about 0.1% of the bear's threat) | - |
| G3 | Checked and aligned: Bear Form threat +50% (21178 aura 10, 50), Maul 128 flat at rank 7 (15 rage, -5 Ferocity), Swipe 83 (20 rage), Lacerate 1235827 (15 rage, 5 stacks, 3 s ticks), Primal Bite 42/62/77 flat plus 100% weapon at 40/50/60 (20 rage, 6 s category cooldown), Frenzied Regeneration (3 min category), Barkskin (60 s, -20%), Natural Reaction (5% dodge, 100% chance of 5 rage at 5/5), Thick Hide, Furor, Blood Frenzy (rage on a crit). Mark of the Chosen (S2). Arena Master and Arena Grand Master (aura 87, damage taken) and Aura of Protection are defensive and unmodelled, as are Savior's Sacrifice and Guarded. | rank sweep, client rows | Aligned | - |

## Paladin: Retribution

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| P1 | Holy Forgefire, Verigan's Fist's (6953, band 20) chance on hit, is not modelled: 39 to 59 Holy (1322218 effect 0, 49 with 0.42 variance; Wowhead Forever "dealing 39 to 59 Holyfire damage"). The client row carries no proc rate. At an assumed 1 PPM it is about 2% of a band-20 Retribution paladin's damage. | ItemEffect trigger 2 on 6953; `items.py` | OPEN (the proc rate; a beta log of the mace) | - |
| P2 | Checked and aligned, fresh against the 70707 rows at every band:<br>- Seal of Righteousness procs 25739/25738/25737/25735/25713 at 20/30/40/50/60, each the top rank's; Judgement of Righteousness 20281 at 20.<br>- Seal of Command 20375/20915/20918/20919/20920 at 20/30/40/50/60 (65/110/140/180/210 mana), 1 s proc cooldown (SpellAuraOptions ProcCategoryRecovery 1000), 70% weapon (20424 effect 31).<br>- Holy Strike's eight ranks (6 to 60): mana 5 to 20, weapon percent 25 to 50, flat 12 to 93 under effect 121, 10 s category cooldown; normalized in the engine.<br>- Consecration 20/30/40/50/60 (135 to 565 mana, 8 s), Exorcism 20 to 60 (85 to 345 mana, 15 s; Undead and Demon only, so never cast on the published target), Hammer of Wrath 44/52/60 (295/360/425 mana, 6 s).<br>- Blessing of Might 19834 at 20, 19838 at 60 (25291 is the Ahn'Qiraj book).<br>- Talents: Benediction, Improved Judgement, Conviction, Sanctified Judgement, Vindication, Seal of Command, Sacred Arbiter, Two-Handed Weapon Specialization, Vengeance (1% a rank a stack, 3 stacks, 30 s), Champion of the Light, Unsealed Wrath, Divine Strength, Improved Seals, Holy Conduit, Divine Intellect, Precision.<br>- Gear: Arcanite Champion's Strength of the Champion, the Lightforge 4-piece and 5-piece, the Seal cost libram (249442).<br>- From Blizzard's 15 October notes: Consecration's tick coefficient 0.042 and Righteous Fury's 90% are in the engine. | rank sweep, client rows, code read | Aligned | - |

## Paladin: Protection

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| T1 | Scarlet Retribution (7726, band 40: 13 to 15 Holy on a block, 5 s cooldown) and Toxic Sporidium (6629, band 20: 1 Nature when struck) are unmodelled. Scarlet Retribution is at most 3 damage a second before Righteous Fury's threat. | 1293539, 1291698; Wowhead Forever | OPEN (small; routed to the shared item-proc list below) | - |
| T2 | Bone Armor (1979, band 50, Skeletal ring: 3% on a hit taken, an absorb through 18828) is unmodelled; defensive. | 19409 | OPEN (defensive only) | - |
| T3 | Checked and aligned, fresh: Seal of Fury's seven ranks (10/18/26/34/42/50/58, 40 to 200 mana) and their procs at every band, Holy Shield 30/40/50 (100/150/195 mana, 75/110/153 a block, 10 s, 4 charges), Hammer of the Righteous at 40 (6% base mana, 6 s), Righteous Fury (90% Holy, 33% all; 25780 aura 10), Holy Strike with Shieldsworn (-4 s), Shield Specialization (4%, 3 s), Swift Judgement, Templar's Bulwark, Redoubt, Sacred Duty, Toughness (item armor; round 7 T6 open), Improved Righteous Fury, Anticipation, Iron Creed, Deflection, Conviction. Devotion Aura 10293 at 60. Deathbone 2-piece to 5-piece in `client_sets_gen.go`. | rank sweep, client rows, code read | Aligned | - |

## Shaman: Enhancement

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| N1 | **Stormstrike used the weapon's own speed. The client normalizes it.**<br>- What the client says: Stormstrike 17364 effect 0 is Effect 121, normalized weapon damage, with no percent effect (1.60.1.70707 SpellEffect). It is the same effect round 7 normalized for Hemorrhage, and Backstab, Sinister Strike and Mutilate use it.<br>- What was wrong: `stormstrike.go` called `MHWeaponDamage`, so a slow weapon added its attack power at its own speed (3.5 s or more) instead of 3.3 s for a two-hander or 2.4 s for a one-hander.<br>- What changed: Stormstrike calls `MHNormalizedWeaponDamage`. The band-60 Forest Defender's Axe swings at 3.3 s, so the published band-60 number does not move; bands 40 and 50 fall (table below), and slow weapons lose a little in the ranker. | `shaman/stormstrike.go:51`. Test first: `TestStormstrikeUsesNormalizedWeaponDamage` (707.5 against the normalized 686.0 on a 3.5 s weapon; then equal). TestEnhancement goldens move about -1% on their own gear. | CONFIRMED-BUG (client row; the engine's own rule for Effect 121) | E 339cd8104 |
| N2 | The shaman carries a Stoneskin Totem aura (10408, rank 6 at every band) beside its Strength of Earth Totem. It only lowers melee damage taken, so it changes no published DPS; it would matter only if a shaman were simmed as a tank. | rank sweep (aura with 100% uptime at bands 20 to 60) | Aligned for DPS (cosmetic; noted) | - |
| N3 | Checked and aligned, fresh: Stormstrike's charge (one, +20%, Lightning Bolt, Chain Lightning and Earth Shock; not spent on a miss; 8% base mana; 8 s), Windfury Weapon 8232/8235/10486/16362 at 30/40/50/60, Rockbiter at 20, Earth Shock 8045/8046/10412/10413/10414 at 20 to 60, Searing Totem and its Attack ranks, Strength of Earth (S1), Windfury Totem 10614. Talents: Concussion, Elemental Devastation (9% melee crit for 10 s from a spell crit, taken back off spells), Reverberation (-0.6 s), Elemental Focus, Thundering Strikes, Ancestral Knowledge, Mental Dexterity (100% of Intellect), Elemental Weapons, Shamanistic Focus (-45% on shocks), Flurry, Mental Quickness, Maelstrom Weapon, Rage of the Farseer (+30% for 25 s, 3 min: 425336), Tidal Focus, Mindfulness. Gear: Chthonic Power and Molten Fury, Fiery War Axe's Fireball, Manual Crowd Pummeler's Haste, Ironfeather 2-piece. | rank sweep, client rows, code read | Aligned | - |

## Cross-cutting

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| X1 | **Unmodelled item effects on published sets, all specs** (routed; most are other lanes' specs at leveling bands). The item sweep finds no engine code for: Gloves of Holy Might 867 (Survival 40, 50), Thornspike 6681 (Assassination and Combat 30), the Shadow Bolt of 6831 (rogues 40), Lifeforce 10750 (rogues 50), Blood Siphon 10844 (Balance and mages 50), Uther's Light 11302 (Balance, mages), Ignition 11750 (mages, Elemental 50), Breath of Fire 744 (Frost 50), the Satchels of Bombs 285275/285277 (Protection warrior 20, 50), Holy Forgefire 6953 (Retribution and Holy 20), Scarlet Retribution 7726 and Toxic Sporidium 6629 (tanks 20, 40), Bone Armor 1979. Most are Forever-authored ids (1290000 and up) with no proc rate in the client rows. | `scratchpad/r8/hybrids/items_all.py` output | OPEN (needs each proc's rate; the engine's convention elsewhere is "1 PPM assumed") | routed |
| X2 | Feral Faerie Fire's rank aura: ranks 1 to 3 use their own aura labels in the minor armor category, so the strongest reduction on a target wins (`core.MinorArmorReductionAura`); rank 4 is the raid preset's aura. | `druid/faerie_fire.go` | Design note | - |

## Before/after

Before: pin cd4dfdcc0 on origin/main 170b7b84c. After: `lane/r8-hybrids` at 9340e9743 with `wt/r8-hybrids`. 10,000 iterations; tanks in threat a second.

Raid preset, band 60: **no change for any of the five specs** (0.00% at both seeds and both factions). Absolute, seed 1, alliance; horde: Feral cat 620.9; 616.3. Bear 864.6; 868.6. Retribution 662.4; 659.6. Protection 995.3; 1011.5. Enhancement 600.2; 599.7.

Bare preset (the published preset below 60), horde, seed 1 / seed 50001, and absolute at seed 1:

| spec | band | % seed 1 / 50001 | absolute, seed 1 |
|---|---|---|---|
| Feral cat (F1, F2) | 20 | +3.3 / +3.3 | 45.9 → 47.4 |
| Feral cat (F1, F2) | 30 | +4.4 / +4.4 | 97.4 → 101.7 |
| Feral cat (F1, F2) | 40 | +3.3 / +3.3 | 131.4 → 135.7 |
| Feral cat (F1, F2) | 50 | +3.6 / +3.5 | 173.8 → 180.0 |
| Feral cat (F1, F2) | 60 | +3.0 / +3.0 | 270.2 → 278.2 |
| Feral bear, threat (G1) | 20 | +14.0 / +14.2 | 50.1 → 57.1 |
| Feral bear, threat (G1) | 30 | +12.4 / +12.6 | 68.7 → 77.2 |
| Feral bear, threat (G1) | 40 | +5.6 / +5.6 | 166.7 → 176.1 |
| Feral bear, threat (G1) | 50 | +4.8 / +4.6 | 315.1 → 330.2 |
| Feral bear, threat (G1) | 60 | +4.5 / +4.6 | 429.2 → 448.5 |
| Enhancement (N1) | 20 | +0.0 / +0.0 | 33.7 → 33.7 |
| Enhancement (N1) | 30 | +0.0 / +0.0 | 85.3 → 85.3 |
| Enhancement (N1) | 40 | -0.7 / -0.7 | 121.8 → 120.9 |
| Enhancement (N1) | 50 | -0.7 / -0.7 | 161.4 → 160.3 |
| Enhancement (N1) | 60 | +0.0 / +0.0 | 217.3 → 217.3 |

The bear's large gain at 20 and 30 is Faerie Fire's armor cut against the low-armor targets of those bands (175 and 285 armor), with its flat threat on top. The cat's band-20 gain is Faerie Fire alone (it keeps Claw there, F1). Alliance was not run separately: no change here depends on race.

## For other lanes

- **Integrator / ranker:** after the pin, the bare-preset feral cat, bear and Enhancement numbers move (table above); a re-rank may move Enhancement's band-40 and band-50 weapons (slow weapons lose Stormstrike attack power).
- **Item procs (X1):** the list above, for whichever lane owns item effects. Nothing in it was changed here.
- `TestRotationLadder` goldens were not regenerated; the feral cat and bear rungs move with F1, F2 and G1, and the Enhancement rungs at 40 and 50 with N1.

## Still open

1. **P1, T1, T2 and X1:** unmodelled item procs (rates unknown).
2. **G2:** Faerie Fire's threat in Bear Form.
3. Carried from round 7: the Ahn'Qiraj book ranks Forever lists at 60 (Ferocious Bite 6, Strength of Earth 5, Grace of Air 3), Holy Shield crits, Toughness's bonus armor, Holy Strike's flat damage, Barkskin in bear, Searing Totem's 2.2 s attack, Lacerate and Primal Bite's threat, Righteous Fury's multiplication, Windfury white or yellow, Arcanite Champion's proc rate.
