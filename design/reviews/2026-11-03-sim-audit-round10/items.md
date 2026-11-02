# Sim audit round 10: items lane

Scope:
- Owner ruling 6's Warsong Gulch faction ids.
- Every item effect on the published BiS gear at every band (`data/builds/1.60.1.70707/bis/*.json`, 719 distinct items) and on the ranker's candidate pool's engine-modelled items. That is every item with an engine effect: 189 items equip in the engine's own tests.
- Side effects of round 9's item fixes.

**Method.**
- **Client tables.** The 70707 raw tables with the hotfix push layered in (`push-check`: moved=false, 2026-11-02). Copied into this worktree, and git ignores them. SpellCategories, SpellProcsPerMinute and SpellClassOptions were fetched from wago.tools for build 70707 on 2026-11-02.
- **Server text.** Wowhead Forever's tooltips: nether.wowhead.com/forever/tooltip/spell/<id>, read 2026-11-02. A tooltip plus the client row counts as one source.
- **Published sets.** Every one of the 336 published entries was built through spec-breakdown's own request path. Every spell its character registers was dumped and compared with the client's SpellCategories.DefenseType and SpellMisc.SchoolMask.
- **Candidate pool.** Every engine item effect was equipped on a level-60 character. The spells it registers were compared with the item's ItemXItemEffect → ItemEffect rows (cooldown, spell id). The literal amounts in each hand-written proc were compared with that spell's SpellEffect rows.
- **Scratch.** Scripts and dumps are in `scratchpad/r10/items/` (`scripts/`, `out/`). The temporary dump tests were never committed.

## 1. Warsong Gulch faction ids (owner ruling 6)

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| W1 | **loot.json carried Warsong Gulch's two reputations under each other's id.** The client has Faction 889 = Warsong Outriders (parent 892, Horde Forces) and 890 = Silverwing Sentinels (parent 891, Alliance Forces). The fork's `db.json` mines its rewards on the client's ids: all 46 Horde-only WSG rewards sit under 889 and all 46 Alliance-only ones under 890. Only the fork's `factions` table names the two the wrong way round, and its `atlasloot.go` comments do the same. `pipeline/loot/sources.py` trusted those names. It "corrected" every restricted reward to the other id, so loot.json had `rep:warsong-outriders:*` under 890 and `rep:silverwing-sentinels:*` under 889. **Fix:** ids stay as mined. The two names come from the client (`_CLIENT_FACTION_NAMES`). Sides are keyed by the client's ids for all six battleground factions. An item that only one side can equip, mined under the other side's battleground reputation, now stops the build (`RepFactionSideError`) instead of being swapped. | Faction.csv 70707 rows 889/890; db.json census (889: 46 horde_only + 16 unrestricted; 890: 46 alliance_only + 16 unrestricted). Four tests failed first: `test_wsg_rep_sources_keep_the_clients_ids_and_take_the_clients_names`, `test_an_unrestricted_wsg_reward_is_named_by_the_clients_faction`, `test_a_battleground_reward_mined_under_the_other_sides_faction_fails_loudly` and `test_a_client_warsong_gulch_reward_joins_the_forks_source_on_its_own_side` (890 ≠ 889). All four pass now, along with 297 loot tests. | CONFIRMED-BUG | S 34e994332 |
| W2 | **Offline build (`loot-merge`, not committed).** The run was made before and after the fix on the same inputs. Only the 8 WSG rep sources change, and only their `faction_id`. Names, item lists, every other source, `factions` and `quests` are identical. All 124 WSG placements sit under the right side: the 92 restricted ones match their own restriction (0 wrong) and 32 are unrestricted. The 70707 ItemSparse has no MinFactionID 889 or 890 row, so no client reputation source names either id today. 155 rows name 509/510. The new test shows that such a source would join `rep:warsong-outriders:<standing>` under 889 and be sided Horde. | `out/` regen diff, 2026-11-02 | Aligned after W1 | - |
| W3 | **The ranker still keys the reverse.** `sim/cmd/leveling-bis/data.go:843-865` (`repSide`: 890 horde / 889 alliance; `correctedRepSource`) and the tests at `data_test.go:1228-1281`, `reputation_test.go:88-90` and `contentphases_test.go` use the old ids. `data/curated/content-phases.json:371-372` labels them the same way. Its gate is by id and both ids share one gate, so this is label only. **The two must land together:** once loot.json is regenerated with W1, today's `correctedRepSource` would relabel Horde rewards "Silverwing Sentinels", and an Alliance character would see unrestricted rewards under "Warsong Outriders". | file:line above | Routed (ranker lane) | - |

## 2. Item effects

### 2a. Which table a proc rolls after the crit split

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| T1 | **Gloves of Holy Might's Smite rolled the spell table.** 1325229 has SpellCategories.DefenseType 3 (ranged). Round 9's law says a melee or ranged defense type rolls melee crit for double damage. The engine built it as a magic spell, so it rolled the wearer's spell crit at 150%, and a Survival hunter has almost no spell crit. It now rolls the ranged table: melee hit and crit, 200% crits. Published for Survival at 40 bare (both factions) and 50 bare alliance. | SpellCategories 1325229 DefenseType 3 (single source: the client row). `TestGlovesOfHolyMightRollsTheRangedTable` failed (defense type 1), then passed. | CONFIRMED-BUG | E c2eed2ba6 |
| T2 | **Thornspike's strike was a melee special.** 1293183 is DefenseType 3 (ranged), so it can be neither dodged nor parried. Its crit table is the same either way. A side effect for a Classic player to judge: a melee-special proc also set off the rogue's poisons and Devious Strike, and a ranged strike does not. Net −0.3 to −0.4% for the band-30 rogues. | SpellCategories 1293183 (single source). New test `TestWeaponProcsRollTheClientsDefenseType` checks all 33 `clientWeaponProcs`; only Thornspike failed. | CONFIRMED-BUG | E c2eed2ba6 |
| T3 | Every other spell the 336 published characters register matches the client's defense type and school where it deals damage. The other mismatches are class spells: dummy parents (Holy Shock, Penance, Prayer of Mending, Blizzard, Hurricane, Volley channels) and two racial or class leeches. All are listed under "For other lanes". The weapon procs in `weapon_proc_rolls.go` (Fatal Wounds, Wounds, Shadow and Frost bolts, Firebolts) are magic or melee, as the client says. | `out/dtcmp.txt` | Aligned | - |
| T4 | **Crit sources on items.** Felstriker 16551 is aura 290 plus aura 54 and maps to `stats.Crit` (every table). Bloodvine 18382 is aura 57 → `SpellCrit`. Blood Tiger is 7597 (aura 52) + 18384 (aura 57) → one `Crit` write. The client set decoder (`core/client_sets.go:395-488`) sends 52 to melee, 57 to spell and 290 to both. No set that a published band reaches has a crit bonus: the only reached bonuses are Ironfeather 14799 (aura 13/135) and Devilsaur 460230 (aura 54/55). No 70707 set bonus uses aura 71. | ItemSetSpell × picks census | Aligned | - |

### 2b. Rates (PPM, one a minute)

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| P1 | **The one-a-minute rule cannot be settled from the client.** SpellProcsPerMinute 70707 has 11 rows: base rates 1 to 10 and 2.3. Exactly one SpellAuraOptions row points at one: 1326019 Holy Smite (item 5028, not in ItemSparse) at 2 a minute. Chance-on-hit weapon procs (ItemEffect trigger 2) and enchant combat spells (SpellItemEnchantment Effect 1, EffectPointsMin 0) carry no rate column; the server holds the rate. No cached beta log (`scratchpad/fl/rep`, 238 reports) contains any of the nine weapon-proc spell ids. Round 9's V8 rows stand. | wago SpellProcsPerMinute, SpellAuraOptions 70707 | OPEN (needs Blizzard data or level-60 logs) | - |

### 2c. Use effects and cooldowns

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| U1 | **Mindtap Talisman (18371) had no use** (round 9 R5). Forever's ItemEffect 235401 is 1302205 on a 7 min cooldown: aura 24 at 125 mana every 1 sec for 10 sec, plus 1302206 Tapped Mind, −60 Intellect for 2 min. Server text: "restoring 1250 mana over 10 sec, but reducing Intellect by 60 for 2 min." It is now a mana cooldown, used once the 1250 and the 900 mana the lost Intellect costs both fit. | Test failed (no spell), then passed. | CONFIRMED-BUG (candidate pool) | E 3bf3a1d1f |
| U2 | **Burst of Knowledge (11832) kept Era's 15 min cooldown.** ItemEffect 99371 gives 6 min. The amount, −150 cost, matches. | Test failed (15m), then passed. | CONFIRMED-BUG (candidate pool) | E 3bf3a1d1f |
| U3 | **Sanctified Orb (20512) gave Era's 3% crit.** Forever's 24865 "Restores 319 to 361 Mana" (Effect 30, 340, variance 0.125), on a 5 min item cooldown and the 1 min conjured category 30 shared with mana gems and Robe of the Archmage. It is now a mana cooldown. TestProtection's Sanctified Orb item row was regenerated (it moves +0.2%). | Test failed (3m, 0 mana, crit), then passed. | CONFIRMED-BUG (candidate pool) | E caa405a7d |
| U4 | **Diamond Flask (20130) gave Era's 75 Strength.** Forever's 363881 is a 5 sec channel that restores 1120 health, then gives 20 Strength for 1 min (6 min cooldown). The channel costs more than it returns in a fight, and the sim runs no pre-pull channel, so the item now has no effect. | Test failed, then passed. | CONFIRMED-BUG (candidate pool) | E caa405a7d |
| U5 | **Smolderweb's Eye (13213) dealt Era's poison.** Forever's 1308917 Mother's Watch is "Remove 1 Disarm effect from target ally". It now has no effect. | same test | CONFIRMED-BUG (candidate pool) | E caa405a7d |
| U6 | **Nat Pagle's Broken Reel (19947)** had a 90 sec cooldown; ItemEffect gives 75 sec. | Test failed, then passed. | CONFIRMED-BUG (candidate pool) | E caa405a7d |
| U7 | The other 30 engine-modelled uses agree with ItemEffect.CoolDownMSec: Hearts, Robe of the Archmage, the hauberks, Earthstrike, Kiss of the Spider, Eye of Moam and so on. Their spell amounts agree with SpellEffect: Earthstrike 280, Kiss 20%, Ephemeral Power 175, Robe 375 to 625, Second Wind 63. Category timers: the Hearts' 1141 is 15 sec, the engine's offensive-trinket timer. Mindtap's category 2554 (10 sec) is shared with uses that no published set pairs with it. | `out/uses.txt` | Aligned | - |
| U8 | **Force of Will (11810)** (round 9 R5). Forever's 15594 is aura 42 on being struck (ProcTypeMask 0x28, ProcChance 100, 30 sec recovery), triggering 15595: −51 damage taken per hit for 10 sec, all schools. The amounts are stated. The rate is not: its text is "has a chance" with ProcChance 100. The one-a-minute rule is built for swings and casts, not for being struck. The item has no ItemSparse row, so simdb gives it Era stats (round 9 D5). | rows 15594, 15595 | OPEN (rate) | - |

### 2d. Chance-on-hit and equip effects Forever retuned

All of these items are unpicked. Each fix changes what the ranker scores for a candidate. Every amount below is the 70707 SpellEffect row, and Wowhead Forever's server text matches each one.

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| R1 | **Blazing Rapier (12777)** (crit lane's lead): 16898 Blaze is 20 every 3 sec for 30 sec ("200 damage over 30 sec"). The engine dealt Era's 10 a tick. Its defense type is magic, so it now lands on the spell hit roll; before, it used the melee special table. | Test failed (0 snapshot), then passed. | CONFIRMED-BUG | E bc2ef10fa |
| R2 | **Gutwrencher (5616)**: 1294426 Rend, 11 every 1 sec for 7 sec. The engine had Era's 8 every 3 sec for 30 sec. **Barman Shanker (12791)**: 1300787 Rend, 14 every 1 sec for 7 sec, doubled when applied from behind (the sim stands behind). The engine had 10 × 10. | Test failed (no 1294426/1300787), then passed. | CONFIRMED-BUG | E d454d6b42 |
| R3 | **The Lobotomizer (19324)**: Forever's 1290950 Brain Damage is Shadow and magic, 200 to 300. The engine had it as physical (armor applied) on Era's 24388. Its 25% cast slow means nothing against a boss. | Test failed, then passed. | CONFIRMED-BUG | E d454d6b42 |
| R4 | **Razor Gauntlets (18326)**: 1302193 deals 7 Nature to a melee attacker. The engine had Era's 3 Arcane. Tank fights only. | Test failed, then passed. | CONFIRMED-BUG | E d454d6b42 |
| R5 | **Draconic Infused Emblem (22268)**: Forever has no use. The client links only the equip 1318931, a chance on a harmful spellcast, which triggers 1318930: +35 spell damage and healing for 10 sec. The engine had Era's use, +100 damage / +190 healing on 90 sec. It now procs at the unstated one-a-minute rate (P1). | Test failed, then passed. | CONFIRMED-BUG | E d454d6b42 |
| R6 | **Fourteen Era procs kept their ids but were retuned.** The engine had the old numbers for all of them: Bloodrazor 17504 (31 every 3 sec for 15 sec), Eskhandar's Left Claw 22639 (21 every 1 sec for 5 sec), Eskhandar's Right Claw 22640 (10% for 6 sec, was 30%), Hookfang Shanker 13526 (13 every 2 sec for 14 sec), Keris of Zul'Serak 16528 (10 every 1 sec, +11% attack time), Skullforge Reaver 17484 (21 every 1 sec for 10 sec, was 2 × 30), Volcanic Hammer 18082 (200 to 256 plus 12 every 2 sec), Lord General's Sword 15602 (84 AP for 10 sec, was 50 for 30), Argent Avenger 17352 (100 AP plus 100 against undead, was 200 against undead only), Ragehammer 12686 (102 AP and 6% speed, was +20 damage and 5%), The Jackhammer 13533 (15%, was 30%), The Cruel Hand of Timmy 17505 (steals 85 life, was −15 to the target's attributes), Armor Shatter 16928 (165 a stack, was 200) and Puncture Armor 17315 (100 a stack, was 200). | Five table tests failed (one row per item), then passed. The fork suite passes with no golden moved. | CONFIRMED-BUG | E 71a463e62 |
| R7 | Small leftovers, with no published or tank-measured effect. Fang of the Crystal Spider 17331 slows the target by 9%, not 10% (hunter weapon; target swing speed only). Gutgore Ripper is registered under SoD's 461682 but deals the client's 75 and −25. Pendulum of Doom's Forever equip, 1293881 (alternating swing speed +100% / −50%), is unmodelled. | rows 17331, 21151, 1293881 | OPEN (minor) | - |
| R8 | The weapon procs in `weapon_proc_rolls.go`, Kalimdor's Revenge, Shadowblade, Chillpike, Emerald Dragonfang, Teebu's and the Hearts all agree with SpellEffect amounts and variance. | `out/deflines.txt` | Aligned | - |

### 2e. Enchants

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| N1 | **No published set carries a permanent enchant.** The ranker sends none (`sim/request/request.go:369` passes only user-supplied enchants), so these matter only on custom /sim gear. Checked against SpellItemEnchantment 70707:<br>- Fiery Weapon 803 → 13897, 40 Fire, magic (spell table).<br>- Lifestealing 1898 → 20004, 30 Shadow leech.<br>- Crusader 1900 → 20007, +100 Strength for 15 sec and a self-heal of 75 to 125 (unmodelled; the sim measures no healing taken).<br>- Rates (Fiery 6, Lifestealing 6.66, Crusader 1 PPM) are the engine's own; P1 applies.<br>Forever adds a second Crusader effect, 458112: a druid-only equip aura (SpellClassSet 7) that procs 20007 for feral forms. It is not modelled separately, because the engine's cat autos already proc the enchant. Icy Weapon 1894 → 20005 slows the target's attacks by 25% (tank-only; unmodelled). Mongoose does not exist in the client. | rows above | Aligned / OPEN (rates; Icy Chill tank-only) | - |

## 3. Round 9's item fixes, side effects

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| S1 | **Aegis' 5 sec cooldown** (E 13b4768d3). Its test passes on the pin, and Protection paladin 40 bare is unchanged by this lane (0.00%). | fork test; measure | Aligned | - |
| S2 | **Dormant Heart** registers through `core.NewClientItemEffect`. The ranker now asks the engine for effects (`sim/cmd/leveling-bis/engine_effects.go:18`, `core.HasItemEffect`), so round 9's stale effects list (X1) is gone. Every new effect in this lane reaches the candidate pool the same way, with no list to regenerate. | code; `go test ./cmd/leveling-bis -run Effect` ok | Aligned | - |
| S3 | **Terrain riders** (S b63fc4cc9). The normalize tests pass. No published effect relies on a rider. | pytest | Aligned | - |
| S4 | **Client reputation sources by side** (S 3ef62621d). These sources are named from the client's Faction table. Since W1, they and the fork's battleground sources share ids and names, which W1's last test proves. | test above | Aligned | - |
| S5 | **Arena Grand Master (19024)**, raised by the physical lane. Its −6% damage taken applies "in the Battle Ring and The Maul" only (server text for 1318318 and the item). Its absorb use is a survival cooldown the sim never fires (round 9 R3). The physical lane changed its row to Aligned. | tooltips 2026-11-02 | Aligned | - |

## Before/after

10,000 iterations. Before: `sb-before` (engine 4d3129cae, origin/main). After: `lane/r10-items` 71a463e62 with `wt/r10-items`. Only T1 and T2 move a published set.

| spec | band | metric | seed 1 / 50001 | absolute, seed 1 |
|---|---|---|---|---|
| Survival hunter, alliance | 40 bare | DPS | +0.23% / +0.25% | 154.56 → 154.92 |
| Survival hunter, horde | 40 bare | DPS | +0.20% / +0.24% | 152.80 → 153.10 |
| Survival hunter, alliance | 50 bare | DPS | +0.19% / +0.22% | 216.96 → 217.38 |
| Assassination, alliance | 30 bare | DPS | −0.42% / −0.43% | 57.67 → 57.43 |
| Combat, horde | 30 bare | DPS | −0.40% / −0.32% | 65.37 → 65.11 |
| Fury 60 raid, Protection paladin 40 bare (controls) | | DPS | 0.00 / 0.00 | identical |

## For other lanes

- **Ranker lane (W3, must land with W1 before the refresh).** In `sim/cmd/leveling-bis/data.go`, flip `repSide` to 889 horde / 890 alliance, then drop or rewrite `correctedRepSource`: on corrected loot it is a no-op at best. Update the comments in `band.go:109` and `data.go:827`. Update the tests in `data_test.go:1228-1281`, `reputation_test.go:88-90` and `contentphases_test.go:91,153,169`. Swap the two names in `data/curated/content-phases.json:371-372`.
- **Ranker lane.** The candidate-pool effects in 2c and 2d change what the ranker scores. Re-rank at the refresh. Mindtap Talisman and the Lord General's Sword, Argent Avenger and Ragehammer numbers can move picks at 50 to 60.
- **Spec lanes (T3, class spells; not fixed here).** Death Coil 6789/17925/17926 registers as physical with no defense type; the client has Shadow, magic. Touch of the Grave 1260198 is a racial leech that the engine gives no defense type; the client has magic. Both are routed to the warlock lane and the lanes that own racials.
- **Fork.** `tools/database/atlasloot.go:405-409` and the fork's factions table name 889/890 backwards. The data is right; only the labels are wrong. Fixing it is optional now that the pipeline no longer trusts those names.

## Still open

1. P1: the one-a-minute rate (needs Blizzard data or level-60 logs).
2. U8: Force of Will's rate on being struck.
3. R7: Pendulum of Doom's alternating swing speed, and Fang of the Crystal Spider's 9%.
4. N1: Icy Weapon's slow and Crusader's self-heal (tank and healing-taken only).
5. Round 9 D4 (Toughness on bonus armor) and D5 (1,351 candidates on Era stats). Mindtap, Force of Will, Burst of Knowledge, Diamond Flask and most items in 2d have no ItemSparse row. Their effects now follow the client, but their stats are still Era's.

Sources: wago.tools db2 1.60.1.70707 (Faction, SpellCategories, SpellProcsPerMinute, SpellClassOptions, SpellAuraRestrictions), fetched 2026-11-02. The build's layered raw tables. nether.wowhead.com/forever/tooltip/spell, read 2026-11-02.
