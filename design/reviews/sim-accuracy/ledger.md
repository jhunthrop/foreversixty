# Sim-accuracy ledger

The owner's goal (2026-10-31): repeat audit and fix rounds over every spec until a whole round finds zero confirmed sim bugs. Each round has three or more audit groups, and every confirmed bug is fixed end to end before the next round starts.

| Round | Dates | Reports | Confirmed bugs | Engine pin after | Zero-bug round? |
|---|---|---|---|---|---|
| 1 | 2026-10-31 | design/reviews/2026-10-31-theorycraft-audit/ (casters, hybrids, physical) | every row in the reports' Fixed columns; the rest held under the adoption rule or open | 17d31ecfa (via 31fbe10e2, c99fb94b4, f923e1ad1, 09ca529f3 and the lane follows) | no |
| 2 | 2026-10-31 to 2026-11-01 | design/reviews/2026-11-01-sim-audit-round2/ (warlock, casters, hybrids, physical) | about 35 fixed; 2 held (warlock demons refusing player buffs, and the Destruction build that depends on it) | c8d29ea33 | no |
| 3 | 2026-11-01 | design/reviews/2026-11-01-sim-audit-round3/ (warlock, physical, hybrids, casters) | 22 fixed (12 from the fresh re-audit, 1 from the wow-player review); round 2's two holds landed | 640da215f | no |
| 4 | 2026-11-01 | design/reviews/2026-11-01-sim-audit-round4/ (warlock, physical, hybrids, casters, integration) | 33 fixed (warlock 6, physical 6, hybrids 8, casters 7, integrator 6); 1 from the wow-player review | f8a5f94ce | no |
| 5 | 2026-11-01 | design/reviews/2026-11-01-sim-audit-round5/ (warlock, physical, hybrids, casters, ranker, page, integration) | 30 fixed (warlock 1, casters 8, hybrids 4, physical 6, ranker 3, page 5, integrator 1); 2 more from the wow-player review | 8baffe5fc | no |
| 6 | 2026-11-01 to 2026-11-02 | design/reviews/2026-11-02-sim-audit-round6/ (warlock, casters, physical, hybrids, healers, ranker, pipeline, wasm, page, integration) | 71 fixed (warlock 5, casters 11, physical 10, hybrids 17, healers 10, pipeline 5, ranker 4, wasm 3, page 5, integrator 1) | b57765fb9 | no |
| 7 | 2026-11-02 | design/reviews/2026-11-02-sim-audit-round7/ (warlock, casters, physical, hybrids, healers, pipeline, ci, page, ranker, integration) | 55 fixed (warlock 3, casters 5, physical 8, hybrids 11, healers 16, ranker 1, pipeline 1, ci 1, page 7 with 2 from the wow-player review, integrator 2) | cd4dfdcc0 | no |
| 8 | 2026-11-02 | design/reviews/2026-11-02-sim-audit-round8/ (warlock, casters, physical, hybrids, healers, auras, items, ranker, pipeline, page, wasm, integration) | 67 fixed (warlock 5, casters 5, physical 5, hybrids 4, healers 14, auras 11, items 13, ranker 5, pipeline 3, page 1, wasm 2, integrator 2; Innervate, shaman regen and Curse of Recklessness were each found by two lanes) | 40021374a | no |
| 9 | 2026-11-02 | design/reviews/2026-11-02-sim-audit-round9/ (warlock, casters, physical, hybrids, healers, auras, items, crit, ranks, ranker, platform, integration) | 67 fixed (warlock 1, casters 4, physical 10, hybrids 6, healers 6, auras 11, items 5, crit 8, ranks 3, ranker 6, platform 3, integrator 4) | 4d3129cae | no |
| 10 | 2026-11-02 | design/reviews/2026-11-03-sim-audit-round10/ (warlock, casters, physical, hybrids, healers, healprofile, ranks, crit, auras, items, ranker, platform, integration) | 88 fixed (warlock 3, casters 5, physical 6, hybrids 5, healers 11, healprofile 3, ranks 3, crit 5, auras 3, items 30, ranker 4, platform 9, integrator 1) | cfcb76e80 | no |
| 11 | 2026-11-02 | design/reviews/2026-11-03-sim-audit-round11/ (regression, warlock, casters, physical, hybrids, healers, healprofile, crit, auras, items, ranks, ranker, platform, integration) | 55 fixed (regression 5, warlock 3, casters 6, physical 4, hybrids 1, crit 5, items 6, auras 5, healers 4, ranks 2, ranker 1, platform 9, integrator 1, plus 3 side effects of round-11 fixes caught before merge); by origin: 6 caused by round 9, 6 by round 10, 1 mixed, 39 older; plus 11 healprofile rows that follow ruling 1 | 47df4a89d (client 1.60.1.70753) | no |
| 12 | 2026-11-03 | design/reviews/2026-11-03-sim-audit-round12/ (regmelee, regcast, warlock, casters, physical, hybrids, healers, healprofile, auras, itemscrit, holypower, ranks, trainables, guides, platform, integration) | 69 fixed (guides 11, regmelee 7, ranks 6, auras 5, healers 5, healprofile 5, hybrids 5, trainables 5, physical 4, casters 4, itemscrit 3, platform 3, holypower 2, warlock 1, regcast 1, coordinator 2); by origin: 23 caused by the landing (11 of them guide prose), 1 by round 11, 2 by round 10, 1 by round 9, 37 older, 5 new data | not pinned: fork int/r12 (batch landing pending) | no |
| 13 | 2026-11-03 | design/reviews/2026-11-03-sim-audit-round13/ (core, regmelee, regcast, warlock, casters, physical, hybrids, healers, auras, ranks, trainables, guides, platform, hotfix, walk, review, integration) | 67 fixed (guides 25, walk 10, hotfix 6, casters 5, ranks 5, healers 4, platform 3, trainables 2, hybrids 2, physical 2, auras 2, regmelee 1); by origin: 6 caused by round 12, 10 by the landing, 1 by round 9, 44 older, 6 new data (5 from the 2026-11-03 hotfix push) | not pinned: fork int/r13 on int/r12 (batch landing pending) | no |
| 14 | 2026-11-03 to 11-04 | design/reviews/2026-11-04-sim-audit-round14/ (core, regmelee, regcast, warlock, casters, physical, hybrids, tanks, healers, data, guides, platform, build, review-physical, review-casters, integration) | 72 fixed (guides 39, healers 7, platform 6, data 5, tanks 4, warlock 3, core 2, casters 2, regmelee 1, regcast 1, hybrids 1, physical 1); by origin: 17 caused by round 13, 2 by round 12, 2 by the landing, 48 older, 3 new data; 42 copy and 30 sim/data/tool | not pinned: fork int/r14 on 5195a8889 (batch landing pending; client 1.60.1.70755 activation in a separate lane) | no |
| 15 | 2026-11-04 | design/reviews/2026-11-04-sim-audit-round15/ (one lane per class, platform, integration) | about 77 fixed; low-rank penalty (17 levels then -5%/level) and weapon DPS in forms confirmed; guide rounds 1-9 seeds folded in | 9f7f1f6e4 (client 1.60.1.70755 + late hotfix re-layer) | no |
| 16 | 2026-11-04 | design/reviews/2026-11-04-sim-audit-round16/ (one lane per class, platform, rank rows wave, integration) | about 43 fixed; from this round a build/APL/card-rank change ships its guide prose in the same landing | f637ec32c | no |
| 17 | 2026-11-04 | design/reviews/2026-11-04-sim-audit-round17/ (one lane per class, platform, integration) | about 15 fixed; the addon/API gear-score rating bug (5-15x) fixed alongside | b2c4a2d05 | no |
| 18 | 2026-11-04 | design/reviews/2026-11-04-sim-audit-round18/ (one lane per class, platform, integration) | about 26 fixed, mostly priest rank rows re-judged level by level under the penalty | 6ea675594 | no |
| 19 | 2026-11-04 | design/reviews/2026-11-04-sim-audit-round19/ (one lane per class, platform, integration) | about 39 fixed, all older; Ice Lance 0.107 from beta logs (Frost -9%), Heroic Strike/Cleave off the proc table and Deep Wounds rollover on Forever logs (Arms/Fury +4%); rogue the first zero lane; three method rules (rank rows on any gain beyond error, client coefficient 0 means check the logs, confirm a lower rank was cast) | da54b7b06 | no |
| 20 | 2026-11-04 to 2026-11-05 | design/reviews/2026-11-04-sim-audit-round20/ (one lane per class, platform, integration) | 31 fixed; hunter attack power shares from Blizzard's 3 Nov notes (BM +14.5%, MM +16.6%); 70790 activated on the way | d5c708834 | no |
| 21 | 2026-11-05 | design/reviews/2026-11-05-sim-audit-round21/ (one lane per class, platform, items, integration) | 41 fixed: 32 old, 7 new-data, 1 R20, 1 R21; Bear Form 115% and Defensive Stance 80% threat (+45%, +38%), Auto Shot through casts (MM +5.8%); Blizzard's notes read line by line against the engine | 170a59c39 | no |
| 22 | 2026-11-05 | design/reviews/2026-11-05-sim-audit-round22/ (platform, items, one lane per class, integration) | 33 fixed: 27 old, 1 new-data, 1 70753, 2 R20, 1 R21, 1 R22; Leader of the Pack and Moonkin Aura exclusive (3% raid crit, every spec about -2.5%), talent spells at their tier level, rogue one-hand axes, Fire/Affliction/Assassination raid builds; weapon damage ranges match the client | bff826207 | no |
| 23 | 2026-11-05 | design/reviews/2026-11-05-sim-audit-round23/ (platform, items, one lane per class, integration) | 19 fixed: 16 old, 3 R22; priest rank rows right after a new rank, Assassination raid build parse (landed alone first), demon gear crit, Robe of the Archmage and Ankh cooldowns, manaIncomePerSecond and the Retribution Consecration budget; raid crit exclusivity sourced to Blizzard's 3 Nov notes | 3b4d6b054 | no |
| 24 | 2026-11-05 to 2026-11-06 | design/reviews/2026-11-05-sim-audit-round24/ (platform, items, data, healbisect, one lane per class or pair, integration) | 70 fixed; the measuring tool judged rank rows on band-floor talents (fixed; 193 rows re-judged); the 70790 hotfix push layered for the first time (Defiance 7/13/20, Natural flasks); manaSpentPerSecond; healer bare-band drops traced to round 15's level-gap penalty (correct) | bf727699f | no |
| 25 | 2026-11-06 | design/reviews/2026-11-06-sim-audit-round25/ (coach, platform, race, shadow, healers-bare, integration) | 23 fixed: 14 old, 7 R21, 2 R24; below 20 the band gear had dropped weapons, bows and wands (now refilled from the client tables, faction-filtered); coach F1 counts on the player's own talent cooldowns and F7 offers melee-only defensives only at physical deaths; the ranker sweeps every legal race (Protection warrior Horde is Undead on threat); healer bare-band drops are round 15's penalty (correct); Shadowform at 40 loses to the current order | f871a65ae | no |
| 26 | 2026-11-06 | design/reviews/2026-11-06-sim-audit-round26/ (platform, melee, ranks, integration) | 19 fixed: 11 old, 6 R24, 1 R21, 1 R25; below-20 gear faction-filtered with armor refilled, the caster and healer rank rows at 10 to 19 re-judged on it; Ghostly Strike on normalized weapon damage (Subtlety raid -0.57%); Combat holds Slice and Dice until 60; Subtlety and Arms leveling orders | 227d95bd1 | no |
| 27 | 2026-11-06 to 2026-11-07 | design/reviews/2026-11-07-sim-audit-round27/ (orders, tools, integration) | 9 fixed: 7 old, 2 R26; four melee leveling orders stop spending early points on talents that do nothing (Subtlety, Arms, Assassination, Combat); Skyborne druid rows below 20; spec-breakdown refuses impossible races; every engine-only spell id has an icon source (icon regen waits on wago) | 978e4c8aa | no |
| 28 | 2026-11-07 | design/reviews/2026-11-07-sim-audit-round28/ (casters, hybrids, integration) | 17 fixed, all old: six leveling orders (Destruction, Demonology, Frost, Feral, Survival, Protection paladin) stop spending points on talents that do nothing ahead of live ones; Fingers of Frost six levels sooner, Conflagrate with Fire and Brimstone, Swift Judgement at 25 | 276ecc361 | no |
| 29 | 2026-11-07 | design/reviews/2026-11-07-sim-audit-round29/ (walks-dps, walks-heal, integration) | 23 fixed: 22 old, 1 R28; the leveling-walk sweep finishes over the last 18 specs (every spec's order now walked); Survival raid drops Deterrence and Counterattack; tanks between bands fight the lower band's boss; tank leveling orders take a threat talent only if the sim chance of death rises under 5 points (coordinator ruling) | fe29dfbd2 | no |
| 2026-11-11 warlock | 2026-11-11 | design/reviews/2026-11-11-sim-accuracy/warlock.md | 6 fixed, all new data from Blizzard's 6 and 10 Nov hotfixes (live client 70888 and its push, which the site's 70790 tables lack): Wrack -30% with no DoT vulnerability; Siphon Life, Corruption and Agony tick twice as fast while Wrack is channelled (new core Dot clock rate); Soul Siphon counts those three only; Imp Firebolt ranks 1-4; Orc Command once on Firebolt and Lash. The Affliction card keeps Agony over Doom at 60 and, from 40, recasts the three under 3 s (raid 60 +5.1 to +5.4% over the old card on the new engine). Raid 60 against the old pin: Affliction -11%, Demonology and Destruction -0.1% Horde | 278876977 | no |
| 2026-11-11 hotfixes, other classes | 2026-11-11 | design/reviews/2026-11-11-sim-accuracy/hotfixes-other-classes.md | 11 fixed, all new data from Blizzard's 9 and 10 Nov hotfixes (server side or client 70888, which the site's 70790 tables lack): shaman 6 (Stormstrike 100% on 8 s, Mental Dexterity 33/66/100, Maelstrom Weapon 20% a stack at every rank and 6 to 10 a minute, weapon imbues with weapon totems, Lightning Overload half the bolt), hunter 3 (Arcane Shot and Summon Hawk 11% of base mana, the hawk's swing +80%), mage 1 (Arcane Concentration split over a Missiles channel), druid 1 (Nature's Swiftness on Wild Growth); warrior, paladin, priest, rogue nothing modelled moved. Enhancement keeps Windfury Weapon under the raid's totem (+10.8% at raid 60 over Rockbiter), every Arcane Shot rank row is gone, the five-stack bolt fires from the first Maelstrom point (+4.8 to +9.1% at 35 to 38). Raid 60 against the old pin: Enhancement -9%, Elemental +2.3%, Arcane -1.3%; bare: Enhancement about -19%, hunters -2 to -7%, Arcane bare 20 -13% | 7d0e129ad | no |
| 2026-11-11 round30-open | 2026-11-11 | design/reviews/2026-11-11-sim-accuracy/round30-open.md | 6 fixed, all old: the tank fight below 60 scales by the heal profile's dungeon table (it was 1.4 to 2x; a bare paladin at 30 to 39 now dies in 6 to 49% of fights, not 70 to 97%); Protection warrior takes Sunder, Thunder Clap and Focused Rage before Anticipation (+1 to +11% threat at 25 to 37, death 0 either way); the bear's Maul gate retuned by level (+0.4 to +6.3% threat at 25 to 59, +0.5 to +0.8% raid 60); rank rows on round 29's moved levels: Arcane Frostbolt, Fire Scorch and Pyroblast, Balance Wrath, Holy and Discipline Renew, Discipline Prayer of Healing, Healing Wave, both Rejuvenations. Held: Arcane Precision before 40 (gains at every level but delays Arcane Power to 43; adoption rule, owner call). Open: a bare bear at 30 to 35 still dies in 93 to 97% (healing below its damage taken until Frenzied Regeneration at 36) | 4a3601e1f | no |
| 2026-11-11 physical logs | 2026-11-11 | design/reviews/2026-11-11-sim-accuracy/physical-logs.md | 0 fixed. Warrior, rogue, hunter and Feral cat against our own advanced-logging report (8 Nov) and all 23 foreverlogs.gg reports: normalized rage exact on a one-hander (7.6 a 2.2 swing, 15.2 on a crit), Energy 10 a second continuous, crit multipliers, Shank (cost, refund, no procs, 25% of the off-hand weapon), off-hand 50%, Windfury about 20%, Touch of the Grave's 15 s, Deadly Aspects within 2% a rank, pet share and rage from damage taken all aligned. Held: Sinister Strike 6% over the client's normalized +15 (Aggression-sized, talents not in the log), Instant Poison's rate (poison uptime), glancing below 60 at +2/+3 (level unknown). No usable Arms, Fury or melee-Survival log exists | 2b10bae6c (unchanged) | no |

Rounds stopped by owner after r29; open items listed in design/reviews/2026-11-07-sim-audit-round29/integration.md ("Open, for round 30"). The round 29 re-rank shipped as addon 2.0.34.

## Loop structure

The owner set this structure on 2026-11-02 (round 11). It replaces the round-by-round practice of re-ranking locally.

1. **Bug-fix rounds run no regens.** A round uses sims, unit tests and small fixtures only. Each ranker bug is proven with a targeted fixture test (one spec, one band, low iterations). No local or CI BiS ranking run, data run or simdb rebuild happens in a bug-fix round. Anything that needs a full re-rank to see goes on the "Measure at refresh" list below. Spec DPS and HPS runs are fine; they run one at a time under `nice -n 15` with GOMAXPROCS at half the cores.
2. **A zero round triggers one full regen, on CI.** When a whole round confirms zero bugs, CI runs the full regen once: data, then simdb.bin, then bis. It never runs on the local Mac.
3. **A verification round audits the regenerated output.** It checks:
   - slot legality;
   - weapon picks;
   - that no set scores below its own check;
   - tier order;
   - every row on the "Measure at refresh" list.

   Fixes from the verification round get a regen of the affected specs only.

## Measure at refresh

Picks and deltas that only a full re-rank can show. Each row is checked in the verification round after the zero-round regen.

| # | what | specs / bands / slots | why | from |
|---|---|---|---|---|
| MR2 | PvP source lines in the BiS files read the in-game rank | every published BiS file | label fix; the page moves at deploy | round 11 ranker K23 |
| MR3 | the 18 specs not re-ranked locally in round 11 | Retribution, Protection paladin, Holy paladin (Libram of Infusion in the relic slot), the priests, rogues, shamans, warlocks, warriors; every band | the ranking hold stopped the local re-rank | round 11 ranker |
| MR4 | Horde paladin gear after Touch of the Grave | Holy, Retribution, Protection paladin, Horde, every band | round 10's Touch of the Grave change; Horde paladins are undead | rounds 10 and 11 |
| MR5 | healer picks on the pull-size heal profile | the five healers, bands 20 to 50 every slot; trinket and relic slots at 60 | round 11 healprofile pull multiplier | round 11 healprofile |
| MR6 | Holy priest picks for the new races | Holy priest, both factions, every band | races are now human and undead | round 11 healers P1c |
| MR7 | classic-db items restated from the client | Royal Seals of Eldre'Thalas, Counterattack Lodestone, Briarwood Reed; ranged-only hit on Trueaim, Marksman Bands, Infernal Trickster, Stoneshatter | client stats replace Era's key by key | round 11 items F1/F2 |
| MR8 | the crit lane's item and enchant mapping | Blackhand's Breadth, Eye of the Beast, Atiesh; scopes 8720 and 8057, Biznicks 2523 | crit tables and use effects from the client | round 11 crit CR2 to CR5 |
| MR9 | simdb's held bonus armor and rounding | 337 items' bonus armor; Blackcrow's hit, Savage Gladiator Chain's crit | the committed simdb predates S 946dde232 | round 11 items |
| MR10 | bare warlock picks with Curse of the Elements in the rotation | the three warlock specs, bare bands | separate bane and curse slots | round 11 warlock |
| MR11 | melee set procs on auto-attacks only | the cat, rogues, Retribution, Arms and Fury at the bands that wear Defias, Shadowcraft, Valor or Bloodmail | proc flags from the client | round 11 items S1 |
| MR12 | round 10's carry-overs | Frost after Fingers of Frost; healer trinkets and relics; the Undead race picks | listed at round 10 | round 10 |
| MR13 | the refresh's data run | run fetch with the item hotfix layer; do NOT record a 165-item drop; `--allow-shrink` is still needed for the wowhead-stats gate | the 165 ids live in the 70532 and 70707 item pushes | round 11 items, ranker |
| MR14 | hunter gear without spell power; Juju Might in the bow sets | Beast Mastery, Marksmanship, every band | Arcane Shot, Serpent Sting and Volley take no spell power (E 05b72a3b5); Juju Might gives ranged attack power (S 0206582a6) | round 12 physical, itemscrit |
| MR15 | paladin raid picks with Elixir of Holy Power | Retribution, Protection paladin, raid 60 (picks, weights, tiers) | new consumable in the preset (E 46be83bcd, S 408d35cde) | round 12 holypower |
| MR16 | the new band builds | cat bands 50 and 60 (leveling build) and raid 60 (c0f5432b2's raid build); Protection paladin band 30 (Holy Shield); Enhancement on b2e4a3bc5's builds | builds change; the published entries carry the old ones | round 12 hybrids |
| MR17 | sub-60 healer picks on the logged split | the five healers, bands 20 to 50; Restoration shaman 20 and Holy paladin 40 above all; the druid idol slot at 60 (Idol of Swiftness) | tank share from beta trash logs; rank rows; Water Shield build; idol registered | round 12 healprofile, healers |
| MR18 | Fury and Frost talents in the BiS files | Fury bare and raid 60; mage-frost bare bands | Fury's build moves dead points; Frost's leveling build kept on audit | round 12 regmelee, casters |
| MR19 | haste scale factors, and guide builds against the stale published builds | every spec's haste; Holy, Discipline, Restoration druid builds | dps-per-point band preset and the core cast-timing fix; the published builds predate the guides | round 12 casters, healers |
| MR20 | hunters on the 2026-11-03 hotfix push | Beast Mastery, Marksmanship, Survival, every band; Marksmanship's Lone Wolf raid build | Aimed Shot, Arcane Shot, Serpent Sting, Sniper Shot, Aspect of the Hawk, Trueshot Aura raised (single source: the push; Wowhead still shows the old text) | round 13 hotfix |
| MR21 | healer picks after the rotation fixes; the druid's raid scale | Restoration druid (coverage about 95%: owner retune of 1.85), Holy priest | Rejuvenation ranks and pacing, Prayer of Mending whenever ready | round 13 healers |
| MR22 | the walk's band builds | Arms, Fury, Survival, Assassination, Subtlety, Demonology, Affliction, Fire, Restoration shaman, the bear; bands 20 to 40 | cast talents that fit are now taken | round 13 walk, hybrids |
| MR23 | Subtlety and Shadow | Subtlety picks and weights (armor ignore as a share); Shadow with Vampiric Embrace | E b08a1c83c; Vampiric Embrace upkeep | round 13 physical, healers |
| MR24 | the round's guide builds | Elemental bare 50/60 (Tidal Focus), Balance bare 30 to 60, Affliction raid (Improved Bane of Agony), Fury raid (Raging Blows) | owner rulings 1 and 2; adoption rule | round 13 casters, warlock, regmelee |
| MR25 | dps-per-point and haste weights | every spec | random regen phase; Elemental's Lava Burst tie | round 13 core, casters |
| MR26 | tanks on the boss with attack power and Demoralizing Shout/Roar in the rotations; Protection paladin band 20 Templar's Bulwark | the three tanks, every band | owner ruling 4 (boss AP 268 at 63, swing x (1 + AP/630)) and ruling 1 | round 14 tanks |
| MR27 | healers on the 2.3 heal profile and the new rank rows | the five healers, every band; the heal-profile block in every healer BiS file | owner ruling 3 retune; Flash Heal 7, Renew 1, Heal 1, Lesser Heal 1, Healing Wave 2, Holy Shock 4 | round 14 healers |
| MR28 | Horde Affliction and Destruction on Orc; Demonology bare bands on Unholy Power | the three warlocks | race pick beyond 1%; dead Improved Imp replaced | round 14 warlock |
| MR29 | the cat's weapon DPS term in weights and tooltips | druid-feral, every band | Forever forms use the weapon's DPS; the landing's workaround removed | round 14 hybrids |
| MR30 | Subtlety with Premeditation in the fight; Gnome rogues and warriors | rogue-subtlety; Gnome rogues/warriors | Premeditation needs no Stealth; Expansive Mind | round 14 physical, data |
| MR31 | Beast Mastery bare 40/50/60, Destruction 40, Shadow bare 60 after the rank rows | those bands | rows removed | round 14 data, regmelee |
| MR32 | the ranker's race sweep in every band report | every published BiS file | best race, margin and error next to the pick; no free race may beat a guide race beyond error by over 1% at 60 | round 25 race |
| MR33 | Protection warrior Horde on Undead | warrior-protection, Horde, every band; then drop its knownStaleBISFiles entry | guide race moved on threat (Touch of the Grave) | round 25 race |
| MR34 | Shadow's published band 20 to 40 talents and Alliance race | priest-shadow | they should follow the leveling order (Shadowform at 42) and Human | round 25 shadow |
| MR35 | Combat bare bands and Subtlety picks | rogue-combat bare 20 to 50; rogue-subtlety every band; warrior-arms band 20 | Slice and Dice waits for 60; Ghostly Strike normalized; Arms leveling order | round 26 ranks, melee |
| MR36 | Combat and Assassination band 20 on the round 27 orders; the 54 engine-only icons | rogue-combat, rogue-assassination band 20; spellicons.json | leveling orders changed; icon regen held for wago | round 27 orders, tools |
| MR37 | the round 28 band talents | warlock-destruction 20 to 40, warlock-demonology 20, mage-frost 20 and 30, hunter-survival 30 | leveling orders changed | round 28 casters, hybrids |
| MR38 | the round 29 band talents | 19 specs (see round 29 integration.md), Survival raid 60 too | leveling orders changed; Survival raid build | round 29 walks-dps, walks-heal |
| MR39 | Affliction on the 10 Nov Wrack hotfix | warlock-affliction bands 40 to 60 and raid 60: BiS, weights, races, the raid build's talents (Soul Siphon, Siphon Life, Improved Bane of Agony gain; Bane of Doom left the card), the leveling order's path to Wrack | the card and engine moved about -11% at raid 60 | 2026-11-11 warlock |
| MR40 | Orc warlocks after Command's second 5% left Firebolt and Lash | warlock-demonology and warlock-destruction race sweeps, Horde | -0.1% for Orc | 2026-11-11 warlock |
| MR41 | the client move to 70888 (issue #21) | `wrack.go` and `imp.go` carry push 113766 by hand; once the generated Wrack/Firebolt tables carry it, read them again (tests unchanged) | the site's data is 70790 | 2026-11-11 warlock |
| MR42 | Enhancement on the 10 Nov hotfix | shaman-enhancement every band and raid 60: BiS, weights, the stat order (Strength against Intellect after Mental Dexterity's cut; the guide's statPriority still ranks Intellect second), races, talents (Maelstrom Weapon ranks now buy procs), leveling order, the 25% Earth Shock reserve | card and engine moved about -9% at raid 60 and -19% bare | 2026-11-11 hotfixes |
| MR43 | Elemental on the 9 Nov overload rows | shaman-elemental Lightning Bolt rank rows, weights, BiS | overload half the bolt again (+2.3% at raid 60) | 2026-11-11 hotfixes |
| MR44 | hunters on Arcane Shot's 11% | hunter-beast-mastery and hunter-marksmanship Serpent Sting rank rows (judged jointly with the removed Arcane Shot rows), leveling mana lines; Beast Mastery builds with Summon Hawk after +80% | bare -2 to -7%, raid 60 unchanged | 2026-11-11 hotfixes |
| MR45 | Arcane on Arcane Concentration's split | mage-arcane leveling card at 20 to 29 (Missiles against Frostbolt), weights, BiS; mage-fire and mage-frost weights | Arcane -1.3% raid 60, bare 20 -13% | 2026-11-11 hotfixes |
| MR46 | the client move to 70888 (issue #21), other classes | overload rows, Nature's Swiftness mask, Arcane Shot and Summon Hawk costs back to the generated tables (tests unchanged); Stormstrike, Mental Dexterity, Maelstrom Weapon, the imbue-and-totem rule, Arcane Concentration's split and the hawk's swing stay hand values until a client or push row states them (then drop the conformance server-hotfix entry); trainables check for the new 70888 class rows | the site's data is 70790 | 2026-11-11 hotfixes |
| MR47 | the tanks on the dungeon-table boss | warrior-protection, paladin-protection, druid-feral-bear, bands 20 to 50: picks, tank weights, threat, chance of death, TMI | the boss below 60 scales by the heal profile's band table | round 30 open T1 |
| MR48 | the round 30 band talents and rotations | warrior-protection band 30 (talents; in knownStaleBISFiles); druid-feral-bear every band and raid 60 (Maul gate) | leveling order and APL changed | round 30 open T2, B1 |
| MR49 | the round 30 rank rows | mage-arcane band 20, mage-fire band 50, druid-balance band 40, priest-holy band 30, priest-discipline bands 30 to 50, shaman-restoration band 40, druid-restoration bands 40 and 50 | rows moved on round 29's orders | round 30 open R1, R2 |

## Round 13 at a glance

Client 1.60.1.70753, built on round 12's unlanded heads (site int/r12 93c5d2510, fork int/r12 076a5a517). Nothing is pinned or pushed: site `int/r13` and fork `int/r13` wait for the batch landing (integration.md "For the integrator").

The owner's rulings this round:
- **Tidal Focus** for Elemental at bare 50/60: adopted (+5.7 to +6.1%).
- **Improved Bane of Agony** for Affliction raid: adopted on the Shadow-build precedent (+0.72% on the guide build).
- **Rank rows keyed by race:** built end to end; the faction key is deleted.
- **The mana regen tick's phase is random per fight:** done; every spec within ±0.3%. Energy has no tick in Forever.
- **Arcane Shot spell power (E 05b72a3b5): kept, pending the owner's in-game tooltip check.**
- **The level-60 heal split:** the owner's Warcraft Logs key arrived mid-round. The fetcher (vanilla.warcraftlogs.com MC/Onyxia/BWL as a provisional proxy) stopped when gcloud's login expired; 60 stays 0.659 until `gcloud auth login`.

**The 2026-11-03 hotfix push** (wago, 14:14 UTC; the committed hotfix-push.json from 11-02 is empty) buffs hunters: Aimed Shot 166 to 350, Arcane Shot 217 to 280, Serpent Sting 83 to 100 a tick, Sniper Shot 295 to 590, Aspect of the Hawk rank 6 55 to 110, Trueshot Aura 50 to 100. **Single source: the push itself.** Wowhead's Forever tooltips still show the old text (checked 2026-11-03). Layering the push is the first regen step at the landing.

67 bugs fixed:
- **6 caused by round 12:** two guide fixes, the Lightwell row, two mobile-planner leftovers of the strip-order change, and the int/r12 merge that dropped four preset reasons. No round-12 sim fix caused a sim bug.
- **10 caused by the landing**, all guide prose.
- **51 newly found:** 44 older, 6 new data (5 from the hotfix push), 1 from round 9.

26 are guide or landing copy; 41 are sim, data or tool bugs. Two lanes added mid-round found 16 (hotfix 6, walk 10).

Fixed, by group (rows and evidence in each report):
- **Hunters (hotfix 6):** the push's values; Aimed Shot on the shared cooldown; Marksmanship keeps Serpent Sting and raids with Lone Wolf.
- **Leveling bands (walk 10, hybrids 1):** bands that left out a cast talent that fits (Arms 40 Mortal Strike, Survival 30, Assassination 30, Subtlety 20/30, Fury 30, Demonology 30, Fire 20, Restoration shaman 30, Affliction 30, the bear 30/40).
- **Healers (4):** Restoration druid Rejuvenation ranks and pacing; Holy's Prayer of Mending; Shadow's Vampiric Embrace.
- **Casters (5):** Elemental's Lava Burst tie and Searing Totem; Balance's Nature's Grace and pre-pull Wrath; Frostfire Bolt's speed.
- **Physical (2), regmelee (1), hybrids (1):** Serrated Blades as a share of armor; Demoralizing Shout's threat; Fury's Raging Blows; Templar's Bulwark's gate.
- **Ranks (5), trainables (2), platform (3), auras (2), guides (25):** race-keyed rows; auto-learned spells; illegal talent strings refused; the mobile planners; Read Leyline; guide facts.

Before and after (raid preset, band 60, 10,000 iterations, seeds 1 and 50001, both factions):
- **Restoration druid +12.0** (effective HPS), **Marksmanship +12.7 to +12.9**, **Beast Mastery +6.2 to +6.4**, **Holy priest +1.4 to +1.6**.
- **Shadow -3.9 to -4.0** (Vampiric Embrace), **Subtlety -1.1**.
- Balance +0.5; every other spec within ±0.2.
- Below 60: Restoration druid bare 60 +40%, Holy bare 60 +22%, bear band 30 +41% (threat), Arms 40 +20%, Survival 30 +10%, Assassination 30 +8%, Elemental bare 50/60 +6%.

Open going into round 14:
- **Owner decisions:** Protection paladin band 20 Templar's Bulwark; same-instant ties in core (Discipline at exactly 0% haste); the Restoration druid's heal scale; the tank boss's attack power.
- **Verify in game:** Arcane Shot's tooltip with spell power up; Era book ranks at a level-60 trainer; six Wowhead-baseline talents; Heart of the Lion; the push's values once Wowhead re-scrapes; gcloud reauth for the 60 split.
- **Regression lane next round:** everything this round changed (the race key, the regen phase, the push's values, the walk rows, the healer rotations, the legality check).

## Round 12 at a glance

Client 1.60.1.70753, from the combined landing c6bdc855c (round 11 plus the launch-day guide lanes) and pin 2606bc05e. Nothing is pinned or pushed: site `int/r12` and fork `int/r12` wait for one batch landing (integration.md "For the integrator" lists the pin and the regen order).

The owner's rulings this round:
- **Rank rows are adopted beyond error.** The 1% bar is for builds. The six sub-1% rows stay; three rows held for missing 1% are adopted.
- **A faction key on rank rows,** only where factions differ beyond error. Eight rows are keyed; every one turns on the published race (a race key is an open owner question).
- **The heal split below 60 comes from beta trash logs.** The tank's share is the median share of group damage the tank took, about 0.68 at 20 to 50; 60 stays 0.659 until a level-60 log is readable. Line of sight stays out of the pull rule.

69 bugs fixed:
- **23 caused by the landing:** 11 guide prose, plus rotations, rank rows and builds.
- **4 caused by earlier rounds:** 1 by round 11, 2 by round 10, 1 by round 9.
- **37 older.**
- **5 from new data:** Poised for rogues and cats, Primal Bite's removed rank, Command's demon row, Voice of Truth's move.

The count rose from 55. Two new lanes (guides, trainables) found 16, and the landing caused 23. Earlier rounds' fixes caused only 4, down from 13.

Fixed, by group (rows and evidence in each report):
- **Regression (8).** The landing's pre-pull Bloodrage recast the raid's shout; Arms charged after Bloodrage; Arms swapped stance for an unlearned Recklessness; Improved Charge rage; Dire Bear Enrage armor; Fury's dead points; shout copy; the healing priest's Inner Fire during Lightwell.
- **Negative haste explained (casters 4).** Evocation's lost tick; a cast ending on the global finished late; Inner Fire skipped under slowing; dps-per-point on raid ranks with bare gear.
- **Physical (4) and hybrids (5).** Poised for rogues and the cat; Thunder Clap's raid debuff; missile speeds (hunters, wands, Hammer of Wrath, and the warlock lane's bolts); Arcane Shot without spell power; Primal Bite rank 3; Protection paladin band 30's Holy Shield; Shred note.
- **Auras (5).** Windfury Totem ends with the totem; Blood Fury heals; Orc Command on demons and guardians.
- **Healers (5) and healprofile (5).** Vampiric Embrace 20%; Idol of Swiftness; Reflection; Power Infusion's cost; the logged split, Discipline and Holy paladin rows, Restoration shaman 20's Water Shield.
- **Items (5).** Juju Might back for bow hunters; consumable amounts; item haste once; Elixir of Holy Power for paladins.
- **Ranks (6).** Rows the landing's rotations needed, faction rows, the Destruction Corruption row.
- **Trainables (5), guides (11), platform (3), coordinator (2).** The pipeline agrees with the live tree (all five questioned spells are baseline); guide facts the client contradicts; talent strips in spend order; the weights headline; the coach's rank rule; guide cooldowns; the paladin landing.

Before and after (raid preset, band 60, 10,000 iterations, seeds 1 and 50001, both factions):
- **Protection paladin +2.2 to +2.3** (threat) and **Retribution +1.6**: Elixir of Holy Power.
- **Marksmanship +1.0 to +1.1**: Juju Might.
- **Beast Mastery -0.5**: Juju Might, less Arcane Shot's lost spell power.
- Discipline and Holy +0.2 to +0.4; rogues +0.1 to +0.24; warlocks -0.03 to -0.44; the rest within 0.2.
- Below 60: Protection paladin band 30 +38 to +40% (threat), Restoration shaman 20 +12 to +16%, Arms bare 30/40 +2.4 to +3.6%, Holy paladin 40 +2.8 to +3.1%.

Open going into round 13:
- **Owner decisions** (all answered in round 13): Tidal Focus at bare 50/60 (the reviewer would adopt it); a race key on rank rows; regen-tick phase randomisation; a warcraftlogs key for the level-60 split; Improved Bane of Agony under the Shadow-build precedent.
- **Verify:** Arcane Shot's tooltip with spell power up, before launch; Bloodrage's in-combat flag; Prayer of Mending hops; Blood Fury heals; Lightwell and Binding Heal rank 1; the round-11 log list.
- **Regression lane next round:** everything this round changed (Poised, faction rows, the split, travel times, the Holy Power elixir, the Arms opener).

## Round 11 at a glance

Client 1.60.1.70753 went live during the round (engine follow 9f42ff1b3); the pin 47df4a89d carries it, and every number is measured on it.

The owner's rulings this round:
- **The sub-60 heal profile adds a hard pull.** A band fights the harder of its median last boss and its 75th-percentile trash pull, weighted by mobs (assistance-radius chains from cmangos classic-db). Multipliers 20 ×1.61, 30 ×2.06, 40 ×3.05, 50 ×1.85. Bare coverage below 60 falls to 42-90%, and gear picks separate. Rank rows and five band builds are re-measured; the pick check is at the refresh.
- **Conviction keeps its crit on heals.** The hotfix push made its aura crit on every table with no school or class mask, and 70753 ships it. Holy damage spells get it back.
- **Retribution bare 20's small loss to the paced Consecration is accepted.**
- **The 165 item ids are not gone.** They live in the item-hotfix layer round 10's local run skipped. The refresh must not drop them, and `fetch` now fetches the layer.
- **PvP defaults stop at in-game rank 6.** Rank 7 and above stay alternatives. The cap reads the in-game rank, and no pick moves.
- **No regens in bug-fix rounds.** See "Loop structure".

New this round, the regression lane re-audited every round-9 and round-10 change. Of the 55 bugs fixed:
- 13 were caused by earlier rounds' fixes: 6 by round 9, 6 by round 10, and 1 mixed (the stone fix: old for most melee specs, round 9 for Survival).
- 39 were older bugs no round had found.
- 3 were side effects of other round-11 fixes, caught before merge.

The count is falling: 88 in round 10, 55 in round 11, or 66 counting the 11 healprofile rows that follow ruling 1.

Confirmed and fixed, by group (rows and evidence in each report):
- **Regression (5).** The search tools sent no preset; the coach flagged Elemental's own rank; a fork UI buff id; Healing Stream's school crit; the bear's potion notes.
- **Warlock (3).** Banes and curses in separate slots (bare bands gain); Curse of the Elements stacking on the raid's; Conflagrate on a miss.
- **Casters (6).** Arcane Clearcasting at -1000% (Arcane gains at raid); Improved Scorch on a hit; Arcane Missiles rank 1; Undead Shadow's Dark Sacrifice; Elemental's Mana Spring; Robe of the Archmage's shared minute.
- **Physical (4).** Execute per rage by rank; Whirlwind's off hand; rage from damage taken at the tank's level; a Survival note.
- **Hybrids (1).** Benediction on Holy Shield.
- **Crit (5).** Conviction; Blackhand's Breadth and Eye of the Beast; Atiesh; the scopes and Biznicks.
- **Items (6).** Melee set procs on auto-attacks only; Boot Flask crit; client stats over classic-db's; ranged-only hit; Blackfathom stone; the item push fetch.
- **Auras (5).** Improved Mana Spring; the warrior tank's Battle Shout and its zero threat; a note; the Dense stone at 35-49.
- **Healers (4).** Heals learned below 20 keep their coefficient (beta logs); Mender's potions; healer potions fire; Holy priest's races (human, undead).
- **Ranks (2).** Lava Burst rank 2 at 50-59; Restoration druid's Rejuvenation 10 row.
- **Ranker (1).** PvP labels name the in-game rank.
- **Platform (9).** The drawer's preset; the addon's rotation ids and talent spells; the Enhancement guide's totems; the weights tool's SimC form; server encounter defaults; the coach's book ranks, Affliction rules and Swiftmend starter.
- **Integrator (1).** The fork's CI had been red since 40021374a: a golden rounded differently on amd64 and arm64. It is green again.
- **Round-11 side effects (3).** Shadow's Devouring Plague row after Dark Sacrifice; Restoration shaman's Healing Wave rank 2 and Discipline's Lesser Heal rank 3 after the heal coefficient fix.

Before and after, raid preset, band 60, 10,000 iterations, seeds 1 and 50001, both factions, on 70753: Discipline +4.3 to +4.5, Arcane +1.7 / +1.8, bear +1.2 / +1.3 (threat), cat -1.3 to -1.5; Protection paladin and Restoration shaman about +0.8, Elemental +0.7, Fury -0.3 to -0.5; the rest within 0.6. Below 60 the warlocks gain 2 to 8% at bare bands and Arms and Fury lose up to 6.8% at 30 to 50; integration.md has the table.

Open going into round 12:
- **New in 70753:** Poised for rogues and the cat; Holy Shield rank 4 in the Protection paladin rotation; the Orc Command rows.
- **Owner decisions:**
  - Whether rank rows also need the 1% bar (six kept rows under it).
  - A faction key on rank rows (Alliance Shadow at bare 60).
  - Elemental's bare 50/60 Tidal Focus.
  - The heal profile's tank-versus-group split and line of sight, and coverage as the sub-60 healer headline.
- **Verify on logs:**
  - Conviction on Holy damage after the push.
  - The heal coefficient fix at 40 and above.
  - Improved Moonfire, Searing Totem, Elemental Mastery, Ice Lance, Fire Power on DoT ticks.
  - Raptor Strike's off-hand penalty; Healing Stream crit; spell modifier stacking.
  - Round 10's log-blocked list.
- **Held under the adoption rule:** Marksmanship, Improved Bane of Agony, Destruction's Immolate line, Fire's Improved Scorch line, Frost's search winner, Frost bare 20's Ice Shards, the band builds the reservation cannot express.
- **Healers:** round 12 gives them a full client-row sweep (this round's lane could not).

## Launch-day guide lanes (2026-11-02)

- **Priest, Shadow builds.** Both Shadow builds move the Discipline points that do nothing for a Shadow priest (Power in Light and Holy Precision touch only Smite, Penance and Holy spells; the raid build's Improved Power Word: Shield and Wand Specialization are never used) into Mental Strength 5 and Silent Resolve 1. Leveling build +2.5% / +2.9% at bare 60, raid build +0.5% at raid 60 (alliance / horde, 10,000 iterations, seeds 1 and 50001). Coordinator ruling: the raid build is accepted below the 1% adoption threshold, because moving points out of talents with no effect for the spec fixes a build defect; it is not an optimization the rule gates.

## Round 10 at a glance

The owner's rulings this round:
- **The heal profile scales by band.** Each band below 60 fights its own content: the open dungeons whose last boss is from the band's level to five above, their bosses' melee and spells (cmangos), the client's x1.2 dungeon modifier, and the client's expected health at the level. Every sub-60 healer rank row was re-measured on its band, and healer band builds can now take curated ranks first: Restoration druid, Discipline and Holy adopt band builds. 60 is unchanged.
- **Retribution's Consecration gate was not needless**: it kept the paladin from running dry. The flat 65% line became a pace line, and rank 1 at 30 to 59 still wins beyond error with no gate at all, so it stays; rank 4 is new at bare 60.
- **The Watchers stays "later"**: Blizzard's raid overview and the 70707 client tie it to the Barrow Deeps prison (raid, 9 December); nothing names a launch zone.
- **sim.wasm goes out as brotli** to browsers that take it (about 1.2 MB less), gzip otherwise; CI budgets both.
- **Page fixes:** the Survival line is time-proof; the coach persona test is green and in CI; "Apply to the page", Top Gear and the weights tool send the preset; the custom panel's icons and weapon rows are fixed in the pipeline (live at the refresh); Warsong Gulch's faction ids follow the client.

Confirmed and fixed, by group (rows and evidence in each report):
- **Crit sweep (owner's extra attention).** Every crit roll in every published spec traced to its client defense type. Conviction is melee and ranged crit, not Holy damage spells (138k logged Consecration ticks); Exorcism never crits (0 of 641 logged); the stone is a bonus on the owner's melee attacks, so no pet inherits it; Weaponmaster and Hack and Slash reach weapon attacks only; Gloves of Holy Might and Thornspike roll the ranged table; heals take their school's crit; Emerald Dragon Whelp; Combustion off Ignite; Inner Focus off Renew; Healing Stream crits on every-table crit only.
- **Pets.** Demons crit 0.7% in melee (a 1,718-swing test, a Blizzard reply, the logs), Lash of Pain 6.4%, Firebolt 2.2%; hunter pets take the hunter's spell crit; Web rolls ranged.
- **Rotations.** Retribution's twist waits for the Righteousness swing (+9.4% raid); Arcane Missiles no longer spend Fingers of Frost (+3.4%); Swiftmend keeps the HoT and is cast under 75%; the bear drinks its potion.
- **Forever rules.** 30-second potions; druids keep potions, Thistle Tea, Windfury Totem and weapon enchants in forms; Touch of the Grave can miss; Spellstone; tanks carry a weapon imbue.
- **Items (30).** Forever's use effects and retuned procs, Mindtap Talisman, Blazing Rapier, the WSG ids, Totem of Urgency and Libram of Infusion.
- **Ranker (4).** Gated vendor stock, the WSG ids, the healer re-measure's iteration cap, a stale report line.

Before and after, pin 4d3129cae against cfcb76e80, raid preset, band 60, 10,000 iterations, seeds 1 and 50001, both factions:
- **Bear +18.6 to +20.4** (threat), **cat +12.9 to +13.0**: forms keep consumables and Windfury Totem; Mighty Rage and a stone on the bear.
- **Retribution +10.4 to +10.6**, **Frost +4.9 to +5.1**, **Restoration druid +6.4 / +6.5**.
- **Protection paladin +2.0 to +2.5**, **Protection warrior +1.5 to +1.7**, **Survival +1.4**, **rogues, Balance, Shadow, Enhancement +0.8 to +1.2** (potions, imbues).
- **Warlocks -0.3 / -0.4**, **Restoration shaman -0.4 / -0.5**, **Fire -0.1**; every other spec within 0.1.

Every item-data fix, re-rank, panel icon and the WSG ids reach published numbers only at the owner's single refresh after a zero round (simdb.bin, loot.json, simbuffs.json, Data.lua, the BiS re-rank).

Open going into round 11:
- **Owner decisions.** A pull-size or band multiplier for the sub-60 heal profile (bare coverage 72 to 98% below 60 against 29 to 40% at 60), and showing sub-60 healers as coverage; Conviction on paladin heals (5.6% of Holy paladin's number); Retribution bare 20's 0.3 to 0.5% under the pace line.
- **Found late, unverified:** Blackhand's Breadth and Eye of the Beast crit tables; Restoration druid's band-20 Furor against 5051 on the merged ladder.
- **Blocked on logs:** a stone on a form's swing; Holy Power; Oath of Light; Holy Shield crit; Master of Elements on Clearcasting; Lightning Overload; demons and the warlock's crit; Thornspike and poisons; normalized rage; Lacerating Strikes' refresh; one-a-minute procs; Force of Will; base-crit columns.
- **Held under the adoption rule:** Marksmanship, Improved Bane of Agony, Destruction's Immolate, Fire's Improved Scorch, Discipline's Inspiration, Frost's search winner, Holy paladin band 20.
- **Refresh checklist:** the data run needs `--allow-shrink` (165 ids gone from wago's export; confirm first); then simbuffs.json, loot.json's WSG ids, the re-rank (Frost, healer trinkets and relics, Undead race picks), and the brotli curl check on the first deploy.

## Round 9 at a glance

The owner's rulings this round:
- **Battlegrounds and Darkspear Islands open at launch.** Both content-phases gates are "launch"; the held refresh gives 347 equippable items an open source.
- **The rank-10 PvP cap stays on `title_rank`**, a realism rule; the code records that it should move to honor cost once the client links honor costs to items.
- **Rotations may downrank when mana-bound.** Below 60 a spell casts the highest learned rank not above the authored one, so curated downranks hold; 43 per-band and per-preset choices sit in data/curated/rank-choices.json, and five raid-60 rotations change rank. Every adopted rank deals its own client amount and costs its own mana.
- **Demonology's sacrifice then summon was a rotation bug**: the Imp is sacrificed only under Demonic Pact. Builds keep their talents.
- **CI:** actionlint accepts `xcode-27`; `real-engine` gates deploy.
- **Melee and spell crit are separate.** Agility gives melee and ranged crit, Intellect spell crit; stones melee, wizard oil spell; item crit and aura-290 buffs every table. Forever beta logs support two pools.
- **Tank weight sweeps were already deterministic**: round 8's failures came from a probe that swept tanks with the damage sweep. **The ranker runs 36% faster** with identical picks.

Confirmed and fixed, by group (rows and evidence in each report):
- **Warlock (1).** Demonology's prepull.
- **Casters (4).** Evocation for Arcane and Frost; Clearcasting spent at the proc's moment; Curse of Shadow gone; base mana from BaseMP.
- **Physical (10).** Hunter pets at their level; ammo by level; the cat's Bite; Savage Strikes; Mortal Shots on Serpent Sting; Serpent Sting's clip; Fury's stance to 36; Ghostly Strike pooling; the execute weapons; Expose Armor 450.
- **Hybrids (6).** Omen and Maelstrom at the proc's moment; imbue attack power by level; Enhancement's Windfury Totem; totem ranks' own amounts; display ids.
- **Healers (6).** Twin Disciplines on Prayer of Mending; racial heal masks; Nature's Swiftness off Wild Growth; the addon's rank table; two notes. Plus Restoration druid's Heart of the Wild build (+1.4%, adopted).
- **Auras (11).** Winter's Chill; Thorns, Retribution Aura, Stoneskin and Mana Tide by rank; Dumplings' level; two salmon foods; Blasted Lands per spec; Survival's second stone; Orc Command.
- **Items (5).** Aegis' cooldown; Dormant Heart; the stale engine-effects list; terrain-only bonuses counted in the item files; the Theramore rewards' missing source (client reputation sources by side).
- **Crit (8).** The split, stones, oil, Cruelty, spell-only bonuses, simdb's crit enchants.
- **Ranks (3).** Curated downranks lifted below 60, in the sim and the addon; a fork test's Holy Light cost.
- **Ranker (6).** The Darkspear vendor's standing and side; Unique-Equipped pairs; Alliance legs on Horde rows; the bare pass's preset; set completion's one combination; wands past the runner-up. Powderfuse opens at launch; The Watchers is gated "later".
- **Platform (3).** The Discord bot's PvP line; the coach's Well Fed and foods.
- **Integrator (4).** Rotation notes that broke the web build; `melee_crit`/`spell_crit` in the weights vocabulary; a stale addon-builds.json; the coach catalogue's missing salmon foods (logs CI).
- Also adopted: Assassination's leveling build trades Venom for Seal Fate (+1.6% bare 60).

Before and after, pin 40021374a against 4d3129cae, raid preset, band 60, 10,000 iterations, alliance / horde, seed 1 (seed 50001 within 0.2):
- **Retribution** -5.7 / -5.7, **hunters** -3.2 to -4.7, **Enhancement** -3.2 / -3.1: the crit split, part won back by the hunter fixes and the Windfury Totem fix.
- **Balance, cat, bear, Protection paladin** -2.3 to -2.9; **warlocks** -1.6 to -2.1; **Shadow, Elemental, Fire, Frost, Holy paladin** -1.0 to -2.1; **Arcane** -0.2 to -0.6.
- **Holy priest** +2.7 / +2.9 (tank Flash Heal rank 2, Twin Disciplines); **Protection warrior** +1.3 / +1.4; **rogues** +0.7 to +1.1 (Blasted Lands buffs); **Discipline** +0.5 / +0.8.
- **Arms, Fury, Restoration druid and shaman** within 0.3.

The bare bands move far more: Holy paladin +20 to +58% and Elemental +10 to +46% from the rank choices, Frost and Arcane up to +28% from Evocation, hunters -2 to -28% from pets and ammo by level. integration.md has them.

Every item-data fix and re-rank reaches published numbers only at the owner's single refresh after a zero round. The refresh must rebuild simdb.bin first (held off-hand typing; round 8's item fixes; round 9's crit enchants), then re-rank; simbuffs.json then names the two salmon foods.

Open going into round 10:
- **Owner decisions.**
  - Should the heal profile scale below 60? Every sub-60 healer rank choice is measured on the level-60 fight, and the notes say so.
  - Consecration rank 1 for Retribution at 30 and 40, adopted behind a mana line though Retribution never runs dry.
  - Brotli or no-inlining for sim.wasm; the time-bound Survival line after launch.
- **After the owner's refresh, check:** the round 8 list, plus Glimmering Staff's spell power, the open battleground sources, the bare-60 rank choices in the published bare sets, and the simbuffs foods.
- **Blocked on level-60 logs:** Holy Shield crit; Faerie Fire's form threat; armor-debuff stacking; the one-a-minute proc rate; Holy Power, Swiftmend, Oath of Light; Era's base-crit columns; the Succubus's crit after the split.
- **Single source or open:** warlock demon sheets against Era's table; aura-52 enchant hands; Concentration Aura; forever-bugs #419; Forever's new elixirs and potions; Lacerating Strikes, normalized-rage off hand, Unbridled Wrath; Toughness on bonus armor; Mindtap Talisman, Force of Will.
- **Held under the adoption rule:** Marksmanship +0.9%, Affliction +0.7%, Destruction, Fire's Improved Scorch line, Discipline's Inspiration.
- **Page and pipeline:** simbuffs.json's two foods (held output; the data gate deselects that one test); The Watchers' launch status and the banned test's AWAITING pass for it; loot's swapped Warsong Gulch faction ids (label only today); the coach persona test's 7 standing failures; "Apply to the page" drops the preset; custom panel icons; `allow_shrink`.

## Round 8 at a glance

The owner's rulings this round:
- **Survival has one build at 60.** Its raid build was +0.7% over leveling, under the 1% bar; raids use the leveling build.
- **Two copy notes.**
  - /tiers/healer says no level-60 log is behind the healer order yet.
  - Survival's guide says why its raid number fell: Forever's Strength of Earth Totem and Curse of Recklessness, on either faction, not the build. The page lane's first draft said "Horde"; the integrator corrected it.
- **One loot generator.**
  - `loot` and `loot-merge` build through one function from `raw/client/loot-classification.json` and write byte-identical loot.json.
  - Mid-round the owner ruled that battlegrounds and PvP ranks exist at launch. Alterac Valley and Arathi Basin spawns are sources on every run.
  - The rank 5/10 tabards keep their client rank.
  - 467 PvP listings whose rank came only from their Era title are now unranked "PvP reward" sources. **The client does not say how Forever sells them**: ItemExtendedCost has honor costs but no item link. They keep a `title_rank` the ranker reads, so BiS eligibility is unchanged.
  - Darkspear Islands has no loot in the client; its vendor is a source, gated "later".
- **Headroom.**
  - The character chunk went from 18,427 to 17,260 of 18,432 bytes (the coach history list loads lazily).
  - sim.wasm.gz went from 95.0% to 90.6% of 4.5 MiB with a byte-identical module (pigz zopfli, and the budget now gates the shipped file).
- **The ranker's treatment of PvP and battleground items is recorded, unchanged.** Rank 1 to 10 rewards are open; rank 11+ are alternatives only. No battleground item is eligible: the "battleground-reputation" and "darkspear-islands" gates are "later".

Confirmed and fixed, by group (rows and evidence in each report):
- **Warlock (5).**
  - Kindling Stave's Ignition was unmodelled.
  - Necropile Drain fired on every cast.
  - A shared 10 ms recovery ate procs.
  - Display only: Demonic Sacrifice's buff ids, and prepull aura uptime.
- **Casters (5).**
  - Arcane Blast's rank from 40 to 49.
  - Innervate below 40.
  - Balance's bare filler ran dry.
  - Shaman regen 17 + Spirit/5.
  - The rank table trusting an NPC copy.
- **Physical (5).**
  - Curse of Recklessness 505 armor, no attack power.
  - 26 weapon procs at their client damage.
  - Whirlwind before 36 and Berserker Stance before 30.
  - Mutilate rank 1 at 30.
  - The stale Raptor Strike note.
- **Hybrids (4).** Shred rank 1 at 20; Faerie Fire in Cat and Bear Form; Stormstrike normalized.
- **Healers (14).**
  - "Damage and healing" bonuses heal (Devout 3-piece).
  - The sub-20 level penalty.
  - Innervate at 40.
  - Shaman regen.
  - Lightwell ticked three renews.
  - Twin Disciplines on Prayer of Mending.
  - Divine Grace, Contingency Plan.
  - Desperate Prayer is dwarf-only.
  - The Holy paladin kit's Wisdom.
  - The shared HoT-aura guard.
  - Two note fixes.
- **Auras (11).**
  - Fortitude, Divine Spirit, Blood Pact and Windfury Totem by rank.
  - Leader of the Pack and Moonkin Aura at their talent levels.
  - Elemental's own Mana Spring stacking with the raid's.
  - Curse of Recklessness.
  - Every raid debuff by the raid's level.
  - Caster-only debuffs out of the raid flags.
  - Six foods and two elixirs at Forever's values.
- **Items (13).**
  - Stat 45 heals.
  - Wowhead percentages to the decimal.
  - Stat 50 is bonus armor.
  - Manual Crowd Pummeler 3 min.
  - Fiery War Axe, Bow of Searing Arrows, Gut Ripper.
  - 14 unmodelled effects on picked items.
  - Wizard oils heal.
- **Ranker (5).**
  - Paladins train axes.
  - Weights measured on a naked character (bands now rank under both weight sets).
  - An empty stat trinket.
  - "PvP rank 0".
  - Ankh of Life in both trinket slots.
- **Pipeline (3).** The loot flip, the merge losing client ranks, the lopsided Stone Guard's rank 10.
- **Page (1).** "PvP rank 0" labels.
- **Wasm (2).** A plain gzip; the budget measured a different file than the one deployed.
- **Integrator (2).** wasm-gzip.sh lacked its exec bit (deploy would stop); the Survival line's "Horde".

Before and after, pin cd4dfdcc0 against 40021374a, raid preset, band 60, 10,000 iterations. Alliance / horde, seed 1; seed 50001 is within 0.2:
- **Every physical attacker** falls 1.1 to 2.0%: the curse, and Survival -1.7 / -1.7.
- **Retribution** -0.2 / -0.1 (Smoked Desert Dumplings offset the curse).
- **Damage casters** rise 1.4 to 2.2%: foods, debuff ranks, Ignition. Balance +2.2, Frost +2.0, warlocks +1.6 to +1.9.
- **Restoration shaman:** +2.2 / +2.2 HPS.
- **Discipline:** +1.6 / +0.9 HPS.
- **Holy paladin:** +0.9 / +1.0 HPS.
- **Restoration druid:** +0.9 / +0.9 HPS.
- **Elemental:** +0.7 / +0.8.
- **Tanks:** Protection paladin +0.6 / +0.7, Protection warrior -0.6 / -0.7, bear -1.5 / -1.4.

The bare bands move far more. integration.md has the largest moves, among them Balance bare 40 and 50 +23 and +26%, Restoration druid bare 20 -24%, and Holy paladin bare 20 to 50 +11 to +20%.

The item-data fixes and re-ranks reach published numbers only at the owner's single refresh after a zero round:
- this round's: stat 45, decimals, stat 50, paladin axes, the Pummeler and the other item effects at the leveling bands, geared weights, the healer sets, the Survival single build, the loot relabel
- round 7's: held items, feral attack power, supplement wands, one-hand casters' off hands, the priest builds, the coverage check, /tiers/healer's order

Open going into round 9:
- **Owner decisions.**
  - Open the "battleground-reputation" and "darkspear-islands" gates at launch?
  - Keep the `title_rank` cap?
  - How Forever sells the 467 "PvP reward" items: not in the client data.
  - Rotation policy: always the top learned rank. Elemental runs dry bare; downranking measured +14% at band 50. The rewrite lifts the band-60 Renew and Shield downranks below 60.
  - Demonology's wasted sacrifice at bands 30 to 50.
  - Brotli or no-inlining for sim.wasm.
  - The time-bound Survival line after launch.
- **After the owner's refresh, check:**
  - the published picks against local re-ranks, especially the healer sets with stat 45 and the Pummeler bands
  - heal-profile coverage
  - /tiers/healer's order
  - the Raptor Strike note on the guide's Rotation card (it reads nightly addon-data)
  - the Survival single build
  - paladin axes
- **Blocked on level-60 logs** (launch is 2026-11-04): every round-6 and round-7 row, plus the following.
  - The one-a-minute rate for "chance on" procs with no client rate.
  - Armor-debuff stacking under aura 674.
  - The sub-20 level penalty.
  - Swiftmend's amount.
  - Faerie Fire's form threat.
- **Single source or open:**
  - One crit stat for melee and spell.
  - Mana Tide's tick.
  - Retribution Aura, Thorns and Stoneskin ranks.
  - Tank weight sweeps that disagree between seeds.
  - The Satchels; Red Dragonscale Leggings.
  - Concentration Aura, Healing Wave rank 10, Chain Heal's party-only jumps, Power in Light, forever-bugs #419.
- **Held under the adoption rule:** Marksmanship +0.9%, Affliction +0.7%, Destruction, Fire's Improved Scorch line +0.4%, Assassination leveling +3.6%, Discipline's Inspiration.
- **Page and pipeline:**
  - `real-engine` does not gate deploy.
  - `allow_shrink` while data.yml passes it to `loot`.
  - The coach catalogue's old foods.
  - The four dead caster-only flags on the custom /sim panel.
  - The Discord bot skips an unranked PvP line.

## Round 7 at a glance

The owner's rulings this round:
- **Survival:** Hawk Eye's points go to Deadly Aspects. In the raid build, which already had 5/5, they went to Lightning Reflexes and Rapid Killing.
- **The Wowhead refresh carries saved records by itself.** The floor is on the carried catalogue, and only a record whose item left the client is dropped, loudly.
- **CI builds the wasm and runs the four real-engine browser tests** (`real-engine` in web.yml).
- **/sim's bundle has room again** from trimming code: sim-island went from 87,518 to 74,834 bytes, with the budget unchanged.
- **Every action a published result carries has a name.** Ignite and Shoot were already in their true schools; a test holds them.
- **Mid-round, the owner held every BiS and data refresh until a zero round.** bis run 41363895819 was cancelled, and the bis, data and engine workflows are off. Picks were measured locally and never committed.
- **The 14 dropped loot sources are recorded, not acted on.** Our two loot generators disagree about battlegrounds, so loot.json flips every pair of runs.

Confirmed and fixed, by group (rows and evidence in each report):
- **Warlock (3).**
  - The mana potion waited for the execute phase.
  - Destruction lacked Elixir of Shadow Power beside Fire Power.
  - Spirit was weighed at zero although Life Tap scales with it.
- **Casters (5).**
  - Shadowform's cost cut was summed with Mental Agility.
  - Arcane Instability's all-crit missed wands and procs.
  - Silent Resolve and Searing Light missed schools.
  - Three Wowhead-supplement wands were Physical.
- **Physical (8).**
  - Rend crits and Impale; Bloodthrill 5 s.
  - Dual Wield Specialization on Whirlwind's and Mutilate's off hand.
  - Hemorrhage normalized.
  - Feral attack power, in simdb and in the item files.
  - The warrior Deadmines addon spec.
- **Hybrids (11).**
  - Strength of Earth 42 and Grace of Air 77.
  - Mark of the Chosen.
  - Moonkin Aura counted twice, and Insect Swarm's crit.
  - Judgement of Fury on the melee table, Iron Creed, the shared immunity cooldown, Hammer of the Righteous at 4 targets.
  - The planner's equip stats.
  - Held items replacing the main hand.
- **Healers (16).**
  - Mana Spring cancelled by Blessing of Wisdom, and Mana Spring's rank below 60.
  - Oath of Light.
  - Healing Stream uncounted; Water Shield on pulses.
  - Pushback on raid pulses: the resist chance, the delay cap.
  - The Focus talents.
  - Clearcasting from HoTs, Bloom's crit.
  - Ankh of Life.
  - Spiritual Healing on Binding Heal.
  - The heal profile retune, 1.65 to 1.85.
- **Ranker (1).** A one-hander that replaced a two-hander left the off hand empty.
- **Pipeline (1).** The manual `allow_shrink` for the Wowhead feed.
- **CI (1).** No job ran the real-engine specs.
- **Page (7).**
  - Unnamed spells.
  - Raid-buffed without the class kit.
  - Crit rates rounded per worker.
  - Zero-damage rows counted as Physical.
  - From the wow-player review: the "<1"-hit wand row, and internal spell names.
- **Integrator (2).** data's APL rank test and spellranks' dead Greater Blessing of Light override had both been red on main since c0bd0bd3d.

Before and after, pin b57765fb9 against cd4dfdcc0, raid preset, band 60, 10,000 iterations. Alliance / horde, seed 1; seed 50001 is within 0.2:
- **Restoration shaman:** +26.0 / +26.1 HPS.
- **Holy paladin:** +13.2 / +13.9 HPS.
- **Restoration druid:** +4.0 / +4.0 HPS.
- **Subtlety:** -2.9 / -2.9.
- **Demonology:** +1.9 / +2.0.
- **Bear threat:** -1.4 / -1.9.
- **Affliction:** +1.7 / +1.7.
- **Destruction:** +1.6 / +1.6.
- **Discipline:** +1.5 / +0.8 HPS.
- **Arms:** -1.4 / -1.4.
- **Enhancement:** -1.3 / -1.3.
- **Elemental:** +1.2 / +1.2.
- **Protection warrior threat:** -1.1 / -1.2.
- **Holy priest:** +0.4 / +1.1 HPS.
- **Every other spec:** within 1.

The healer columns include the profile retune; healers.md splits it out.

The item-data fixes and re-ranks reach published numbers only at the owner's single refresh after a zero round. Those are:
- the supplement wands' school (band 50 bare casters +1 to +4.6%)
- held items
- feral attack power
- the one-hand casters' off hands (+1.2 to +3.3% locally)
- the new Holy and Discipline priest and Survival builds
- the coverage re-check
- /tiers/healer's new order (shaman first)

Open going into round 8:
- **Owner decisions.**
  - The 14 loot sources: which loot generator is right about battlegrounds.
  - Survival's separate raid build at +0.7% over leveling.
  - Display wording for about 413 client spells with internal names (none in a published result).
  - Copy: /tiers/healer's model caveat, and why Survival's number fell.
- **After the owner's refresh, check:**
  - published picks against the local re-rank
  - heal-profile coverage on the re-picked sets
  - /tiers/healer
  - Torch of Light out of warlock raid sets
  - Survival on Deadly Aspects
  - the real-engine CI job's first run
- **Blocked on level-60 logs** (launch is 2026-11-04): every round-6 row, plus the following.
  - Oath of Light: scaled or flat.
  - Gift of Nature on Tranquility.
  - Same-op percent stacking.
  - Fire Power on the Fireball and Pyroblast burns.
  - A Fire wand crit and Ignite.
  - The wand hit table.
  - Rage from damage taken by level.
  - Heroic Strike as a swing.
  - Battle Shout's rounding.
  - Eviscerate's attack power.
  - Holy Shield crits, Toughness's armor, Holy Strike's flat damage.
  - Barkskin in bear, Searing Totem's attack, Fire Nova's range.
  - The Ahn'Qiraj book ranks at trainers.
- **Single source or open:**
  - Concentration Aura (unmodelled, at most +1.3%).
  - Healing Wave rank 10.
  - Chain Heal's party-only jumps.
  - Power in Light.
  - forever-bugs #419.
- **Held under the adoption rule:**
  - Marksmanship +0.9%.
  - Affliction +0.7%.
  - Destruction +0.1 to +0.2%.
  - Fire's Improved Scorch line +0.4%.
  - Assassination leveling +3.6% (drops Murder and Venom).
  - Discipline's Inspiration.
- **Page and pipeline:**
  - The character-page chunk has 2 bytes of headroom.
  - sim.wasm.gz is at 94.7% of its budget.
  - `real-engine` does not gate deploy.
  - normalize's `allow_shrink` re-baseline input and client.yml's hint.

## Round 6 at a glance

The owner's rulings this round:
- **/sim below 60 defaults to the preset the band publishes** (bare below 60, raid at 60). The ranker's rule lives once in `sim/request.PublishedPreset`; Raid-buffed stays a choice.
- **Horde race picks within 1% stay as published.** The adoption rule's bar applies to race.
- **Restoration druid's lead stays,** blocked on launch logs.

The owner's priorities:
- **Protection paladin +50% stands.** Righteous Fury alone is +50.7%, and every input has two sources. This round's two threat bugs take 3.4% back off, leaving the paladin at 96 to 98% of a Protection warrior.
- **Torch of Light leaves the raid sets.** The ranker credited wand DPS that raid warlocks never shoot, and wand shots were Physical.
- **The hotfix push is layered into every generator** and re-fetched when it moves.
- **sim.wasm was at 94% of Cloudflare's limit.** It ships gzipped now, under a 20 MiB budget check.
- **The header layout shift is fixed at its two causes,** with the limit unchanged.

The owner's questions, answered:
- **Survival's Improved Concussive Shot did nothing.** It now goes to Deadly Aspects (+4% leveling, +8% new raid build). Efficiency is worth nothing because Resourcefulness already cuts melee mana by 60%.
- **Wild Growth is modelled right but loses to Rejuvenation on mana.** Raid damage now reaches the healer.
- **The bear was in form all along.** Maul and Swipe lacked Classic's 1.75x threat (+29.5% now).

Confirmed and fixed, by group (rows and evidence in each report):
- **Warlock (5).**
  - Wand shots were Physical.
  - Classic-db items lost single-school spell damage.
  - Forever's dynamic DoTs for warlock, druid and shaman.
  - Balance and Feral registered Moonfire and Wrath twice, casting the talent-less copies.
- **Casters (11).**
  - Ignite rolls into one burn and is registered once.
  - Improved Scorch benefits only its mage, so it is out of the raid preset.
  - Ice Lance's Frozen x4 covers spell power.
  - Arcane Blast lands on cast.
  - Robe of the Archmage's use.
  - Mage and priest DoTs are dynamic.
  - The Mental Agility and Spiritual Healing masks.
  - The addon's stale default builds.
- **Physical (10).**
  - The build search protected Improved Concussive Shot.
  - The Resourcefulness proc mask.
  - Aspect of the Beast.
  - Bestial Discipline's mana half.
  - Claw ranks.
  - Defensive Stance and Defiance multiply.
  - Mana regen reads its rates each tick.
  - Rune of the Guard Captain is 14 attack power, and terrain riders are dropped.
  - Classic-db shield block.
  - Half-up stat rounding.
- **Hybrids (17).**
  - Maul and Swipe threat.
  - Moonkin crit counted twice, Vengeance on Insect Swarm, Nature's Reach.
  - Natural Instinct on Rip.
  - Genesis off Tranquility.
  - Illumination's casting regen.
  - Benediction on Holy Shock.
  - Shield Specialization mana threat, and mana-gain threat multipliers.
  - Judgements on the melee table, Hammer of Wrath's hit, the Benediction mask.
  - Windfury's Elemental Weapons, Stormstrike's charge, Flurry's 15 s.
  - Purification off Healing Stream.
  - Blessing of Light and Salvation removed.
- **Healers (10).**
  - Raid damage reaches the healer.
  - HoTs are dynamic.
  - Discipline casts Greater Heal; Binding Heal is trainable.
  - The Greater Blessing of Light precast is removed.
  - Chain Heal pacing.
  - Mana Tide rank 4.
  - Coach book ranks and priest rank notes.
  - The rotations heal the healer.
- **Pipeline (5).**
  - Generators read only the shipped tables.
  - The hand dump was incomplete.
  - Removed spells stayed as nameless constants.
  - Nothing re-fetched on a new push.
  - The follow hash ignored the push.
- **Ranker (4).**
  - Wand credit by Shoot time.
  - Verify's flat 1% gate.
  - Empty trinket slots.
  - Doubled ranged attack power.
- **Wasm (3).**
  - The deployed module was at 94% of the limit.
  - The local toolchain differed from CI's.
  - The dev replace built the fork checkout silently.
- **Page (5).**
  - /sim defaulted to Raid-buffed below 60.
  - Header layout shift.
  - Stale real-engine e2e specs.
  - From the wow-player review: an off hand under a two-hander drew an empty socket.
  - From the wow-player review: /sim's Damage tab opened collapsed on a nameless "Build" row.
- **Integrator (1).** web.yml had been red on prettier since f95ff8e77, so round 5's pin never deployed. It was pushed alone (bef2cb599), and web has deployed since.

Before and after, pin 8baffe5fc against b57765fb9, raid preset, band 60, 10,000 iterations. Alliance / horde, seed 1; seed 50001 is within 0.2:
- **Bear threat:** +29.5 / +29.7.
- **Discipline:** +11.1 / +11.0 HPS.
- **Frost:** +7.8 / +7.9.
- **Destruction:** -7.8 / -7.8.
- **Restoration shaman:** +7.6 / +7.4 HPS.
- **Holy priest:** +6.1 / +6.3 HPS.
- **Fire:** -5.7 / -6.1.
- **Elemental:** -5.5 / -5.5.
- **Balance:** +5.3 / +5.0.
- **Holy paladin:** +4.5 / +3.7 HPS.
- **Enhancement:** -3.5 / -3.4.
- **Protection paladin threat:** -3.4 / -3.4.
- **Protection warrior threat:** +3.1 / +3.1.
- **Demonology:** -2.9 / -2.8.
- **Retribution:** +2.1 / +1.6.
- **Beast Mastery:** +1.6 / +1.7.
- **Affliction:** -1.6 / -1.6.
- **Feral:** +1.4 / +1.3.
- **Every other spec:** within 0.5.

The item-data fixes and the re-ranks reach published numbers only after data.yml's fetch and bis.yml run on the pin. Those are:
- wand school
- classic-db spell damage
- the doubled ranged attack power
- Guard Captain
- classic-db shield block
- rounding
- Torch out of the raid sets
- trinket slots
- the new Survival and Beast Mastery builds

Open going into round 7:
- **Owner decisions.**
  - Hawk Eye to Deadly Aspects for Survival: +5% bare, at the cost of range utility, the same class of choice as Grim Reach.
  - The gzipped sim.wasm transfer is at 95% of its 4.5 MB budget.
- **After this round's nightly, check:**
  - Wowhead's live feed lists only 58% of the seeded supplement. Run 41363024342 needed `allow_shrink`, and every fetch will until the feed recovers or the floor is revisited: the owner decides which.
  - Loot lost 14 sources at c0bd0bd3d: Alterac Valley world drops and PvP ranks 5 and 10. Confirm they should go.
  - A real /sim run in production loads the gzipped module.
  - The warrior Deadmines addon spec still hardcodes its pick.
  - /tiers/healer reorders, with Discipline passing Holy paladin.
  - The warlock and raid caster ranged slots leave Torch of Light.
  - The new Survival and Beast Mastery builds publish.
  - Horde hunters drop about 1.2% from the ranged attack power fix.
- **Blocked on level-60 logs** (launch is 2026-11-04):
  - Ice Lance's coefficient: the client says 0, the engine 0.4286. At 0, Frost is 9.7% lower.
  - Illumination's casting regen, client rows only.
  - Lacerate and Primal Bite's high-threat factor.
  - Holy Shield's flat threat.
  - Elemental Mastery: one charge, or 15 s.
  - Windfury swings: white or not.
  - Holy Power's crit-heal bonus.
  - Arcanite Champion's proc rate.
  - Ignite's logged 0.73 ratio.
  - Missile Barrage on a miss.
  - Arcane Concentration per missile.
  - Master of Elements on Clearcasting.
  - Prayer of Mending's bounce mask.
  - Binding Heal's crit rolls.
  - Wild Growth under burst damage.
  - Revenge, Shield Slam and Thunder Clap threat.
  - Venom and Rupture attack power.
  - Improved Corruption: add or multiply.
  - Serpent Sting's 0.2 spell power, where the client row has none.
  - Righteous Fury's multiplication.
  - Round 5's demon, pet and rage items.
- **Single source or open:**
  - Shadowfiend at launch.
  - The heal profile's coverage sits outside its 60 to 85% rule: druid about 86%, paladin about 57%.
- **Held under the adoption rule:**
  - Marksmanship's search winner, +0.9%.
  - Fire dropping its Improved Scorch upkeep line, +0.4%.
  - Affliction, +0.7%.
- **Pipeline and page follow-ups:**
  - Items still use the useHotfixes CSV without removals.
  - ChrRaceRacialAbility and the other unfetched tables are hand-written.
  - 24 aura-99-only spells still mirror into ranged attack power.
  - Benediction's planner row lacks its healing line.
  - The real-engine e2e specs need a wasm build in CI (build-real has none).
  - sim-island has 12 bytes of headroom.
  - Raid-buffed doesn't add the class kit (Protection paladin's Devotion Aura at 60).
  - Copy: why Torch loses in raid on BiS.
  - 36 action ids read "An unnamed spell" on /sim's Damage tab: simnames pruning, plus engine ids like Ignite 12654 that the client's SpellName lacks.
  - Ignite and Shoot show as Physical: the result carries no school for them. The engine deals Ignite as Fire, and Shoot's school arrives with damage_school at the data run.

## Round 5 at a glance

The owner's rulings this round:
- **/tiers/healer orders healers only by gaps beyond error at every tuning of the heal profile.** The ranker publishes `hps_error` and `hps_at_bounds` (×1.5 and ×1.65, now configuration in heal-profile.json). With no error published, the page shows no ranks.
- **Power Infusion stays out of every default.** No default carried it; tests now hold that.
- **/sim consumables are each spec's published set,** resolved by level, cut to the character's RequiredLevel and loaded lazily.
- **The ranker's guard covers every role,** and the root causes are fixed.
- **Rhok'delar** is not obtainable in Forever's phase 1. Its leaf drops only in Molten Core, and its sinew comes from Onyxia, which opens 9 December.
- **Ice Barrier's 15% / 5% has three sources:** the 31 October hotfix push on wago.tools, the hotfix cache wowforevertalents.com reads, and Wowhead Forever. Frigid Power carries the 15%, Ice Barrier's own effects the 5%, so the engine stands.

Confirmed and fixed, by group (rows and evidence in each report):
- **Warlock (1).** Healing on items gave a hidden third as spell damage, on top of the damage half the client states. The sources that grant both now state both halves. This is the largest cross-class move of the round.
- **Casters (8).**
  - Mages wore no armor: Molten Armor is an SoD rune a Forever mage cannot learn, so they now wear Mage Armor. Fire also adopts Improved Fire Blast 2/2, +3.3%.
  - Racials follow Blizzard's 26-27 October notes and the hotfix push: Orc Command (Savage removed), Berserking 15 s, Touch of the Grave every 15 s, The Human Spirit 10%, Wind Blessed 2%.
  - Holy Binding is modeled.
  - A Demonic Rune costs 800 health.
  - The race text follows the client.
- **Hybrids (4).**
  - Righteous Fury is +33% all threat and +90% Holy threat, multiplied.
  - Holy Shield costs 100/150/195 mana.
  - Champion of the Light is spell damage only.
  - Healing Stream ticks crit, as periodic heals.
- **Physical (6).**
  - Hunter's Mark rank 4 is 14325 at 143 ranged attack power, not SoD's 1213268 at 110.
  - Deadly Aspects procs Quick Strikes from melee.
  - The item hotfix dump was frozen at build 70009: 201 rows wrong and 569 items missing at 70707. data.yml now layers wago.tools' hotfix pushes.
  - Shields carry block value.
  - Client items take Wowhead's on-equip stats where the client lost the link (Benediction among 153).
  - Item stats 36 and 37 are haste and expertise rating.
- **Ranker (3).**
  - `weight_stats` missed stats Forever's talents turn into damage: Intellect for Enhancement, the hunters and Retribution; spirit and mp5 for mana casters.
  - The stat runner-up never reached verify.
  - Only healers had the published-set guard.
- **Page (5).**
  - The /sim consumables were the page's own lists.
  - They were the level-60 set at every level.
  - They included items above the character's level.
  - A rogue's stone was filed with the potions.
  - /tiers/healer ordered healers by any margin and by the tuned profile.
- **wow-player review (2).** A run in flight when another character loaded wrote its result onto the new character, and the after-sim line read an unowned Droptimizer upgrade.
- **Integrator (1).** The addon showed 20 rotation lines bands before the player can train them, from client copies with no SkillLineAbility row (Recklessness at 30 for Fury, among others).

Also fixed, not counted as sim bugs:
- **data.yml had been red on every push since round 4.** The addon data derived from curated rotations that only the nightly regenerated. A push job now owns it.
- **Two data tests pinned generated state.**
- **The bis.yml conflict handler,** which my push job broke for run 41332442067: that run's picks were lost.

Before and after, pin f8a5f94ce against 8baffe5fc on today's published data. Raid preset, band 60, 10,000 iterations, alliance / horde, seed 1 (seed 50001 within 0.3):
- Protection paladin threat +50.0 / +51.1.
- Fire +5.4 / +5.2 on the published build; the new raid build adds +3.3 at the bis run.
- Balance -4.6 / -3.9.
- Demonology -4.2 / -2.6.
- Elemental -3.4 / -4.1.
- Arcane -2.7 / -3.4.
- Affliction -3.0 / -2.3.
- Destruction -3.0 / -2.3.
- Frost -2.6 / -1.6.
- Shadow -2.2 / -1.8.
- Discipline HPS +1.8 / +0.5.
- Enhancement -0.5 / -1.4.
- Marksmanship +0.9 / +1.1.
- Beast Mastery +0.8 / +1.0.
- Restoration druid HPS +1.0 / +1.0.
- Arms 0.0 / -0.8.
- Retribution 0.0 / +0.8.
- Every other spec within 0.5.

The item-data fixes, the ranker fixes and the re-ranks reach published numbers only after data.yml's fetch and bis.yml run on the pin. The integrator dispatches both after the push. The ranker lane's local re-ranks (verify against the published set, alliance / horde):
- Survival +19.6 / +20.1 (Heartseeker in the off hand).
- Retribution horde +11.6 (Arcanite Champion).
- Enhancement +2.7 / +2.3.

Open going into round 6:
- **Owner decisions.**
  - Horde race picks. Orc beats the published Troll by about 1.0% on all three warlocks (Command). Undead beats the published Orc by 0.5 to 0.7% on Arcane and Fire.
  - /sim defaults to Raid-buffed below 60 while the published bands below 60 are bare. At band 40, buffs give +65 to +181%.
  - The Restoration druid's lead over the other healers rests on the heal profile.
- **After this round's nightly, check:**
  - /tiers/healer shows 1, =2, =2, 4, 5 on the published error.
  - Survival's off hand is Heartseeker; re-shoot /sim.
  - The warlock sets drop the healing-hybrid pieces.
  - Retribution horde is off The Needler.
  - Fire publishes its new build.
- **Blocked on level-60 logs** (launch is 2026-11-04):
  - Demon crit, haste and hit.
  - The demon level-60 sheets.
  - Pet haste.
  - The rage-crit factor and Sunder's AP threat.
  - Holy Strike's flat bonus.
  - Omen's 4% at 60.
  - Master of Elements on a Clearcasting cast.
  - Righteous Fury's two threat bonuses, multiplied or added (about 12% of Holy threat).
  - Fire traps scaling with attack power.
  - Rogue Armor Energize.
  - Raptor Strike lifting the dual-wield penalty.
- **Single source or open:**
  - Tauren row 1349372, +2% critical healing, which survives the push while the notes and tooltip say health only (0.2 to 0.35% HPS).
  - Forever DoTs may read spell power each tick (forever-bugs #398).
  - Demonic Brand: its client rows carry no amounts.
  - Shadowfiend at launch.
  - Kings with Heart of the Lion.
  - Pet family.
  - Siphon Life crit, held for the owner.
- **Pipeline follow-ups:**
  - The spell-table generators still read only the shipped tables, not the hotfix push, so the lanes read the push by hand.
  - About 17 items keep a stat gap Wowhead lists.
  - Keen Machete's quality.
  - 24 rotation lines carry placeholder icons until the next icons run.
  - The local sim.wasm build is 40.4 MB, over Cloudflare's 25 MiB, and sim/go.mod points at the owner's fork checkout rather than the pin.
  - header-finished-sim's CLS check is flaky on origin/main (2 failures in 6).
  - The comment-only fork commit e203aca7e merges with the next pin.
- **Held under the adoption rule:**
  - Affliction +0.7%.
  - Arcane +0.9%.
  - Survival's Deadly Aspects build, +8.6% on a legal dual-wield set, which needs Hawk Eye and Improved Concussive Shot freed from the raid keeps.

## Round 4 at a glance

Confirmed and fixed, by group (rows and evidence in each report):
- **Warlock (6).**
  - Demons take Forever's 17% attack power and 10% spell damage shares, not Classic's 56.5% and 15%. Sources: the Forever UI's `WARLOCK_PET_BONUS`, wowsims/forever's beta readings and the beta log fits.
  - Demon melee takes 8/17 of the warlock's hit and spells all of it. This rests on one Blizzard statement, forever-bugs #394.
  - Lash of Pain takes Malevolence's crit.
  - Shadow Vulnerability raises only the warlock's own Shadow damage.
  - Bane of Doom needs a minute left to be cast.
  - Eureka! follows each Gnome class's own client row: 8 s, no charges, class masks. A new generator, `make eureka`, writes it.
- **Physical (6).**
  - Weapon damage: the item file was Era's `db.json`, and the client curve missed QualityModifier and the caster-weapon bit. 1,027 of 1,036 melee weapons now match Wowhead Forever.
  - Items now carry ranged attack power (stat 39) and "+N Weapon Damage" (stat 83, Molten Heart).
  - Hunter pets inherit hit.
  - Wound Poison is outside Venom and Vile Poisons.
- **Hybrids (8).**
  - Heals over time crit; the client flags every trained one.
  - Touch of the Grave, Orc Savage and Tauren Endurance follow the 70707 client.
  - Maul spends Clearcasting.
  - Bloom's heal rolls no cast procs.
  - Seal of Righteousness takes spell power once. The Retribution list now keeps Seal of Command and twists to Righteousness, gated on mana and attack power.
- **Casters (7).**
  - Arcane Concentration's Clearcasting is one charge.
  - Arcane Blast's stacks and Missile Barrage keep the other cost mods.
  - Arcane Missiles rank 8 is an Ahn'Qiraj book.
  - Fire casts Evocation below 5% mana. It also adopts a raid build with Arcane Meditation, +5.2%.
  - Items keep holy, nature, frost, shadow and arcane spell damage and spell penetration (ItemSparse 84 to 89 and 47).
  - Wowhead-only items keep their flat crit, hit, dodge, parry, block and defense.
- **Integrator (6).**
  - The /sim page defaults to the published numbers' type-neutral target.
  - Its Raid-buffed buffs and debuffs come from the curated raid preset.
  - The coach's sim reference fights the type-neutral target.
  - A rage potion on a class with no rage bar no longer panics the engine.
  - The /sim armor field names the armor of the target the character fights; found by the wow-player review.
  - The BiS ranker kept main-hand-only weapons out of the off hand. The bis.yml run after round 3's weapon-DPS fix had put Teebu's Blazing Longsword in Survival's off hand, where the engine swings nothing: 653.7 against 768.6 on the re-rank. Otherwise that run's picks are sane: Feral takes Hammer of the Titans.

The heal profile is retuned to ×1.65 by its own rule. The Restoration druid covered 91% after heal-over-time crits; raid sets now cover 60.4 to 84.0%.

Before and after, pin 640da215f against f8a5f94ce on today's published data. Raid preset, band 60, 10,000 iterations, alliance / horde, seed 1 (seed 50001 within 0.2):
- Arcane -9.3 / -9.0.
- Fire -12.4 / -12.9 on the published build. The new build adds +5.2.
- Frost -2.2 / -2.5.
- Retribution +7.3 / -10.2. Horde is on Blade of Hanna; on Arcanite Champion it is +10.1.
- Protection paladin threat 0.0 / -4.5.
- Restoration shaman HPS +4.7 / +4.7.
- Restoration druid HPS +2.8 / +2.4.
- Holy paladin HPS +2.0 / +2.0.
- Holy priest HPS -2.1 / -1.4.
- Discipline HPS +0.7 / +0.5.
- Affliction -3.6 / -3.1.
- Demonology -1.4 / -0.7.
- Destruction -1.7 / -1.2.
- Balance 0.0 / -1.0.
- Feral 0.0 / -0.6.
- Feral bear threat -0.9 / -1.1.
- Beast Mastery +0.7 / +0.6.
- Marksmanship +0.5 / +0.7.
- Shadow -0.4 / -0.1.
- Elemental 0.0 / +0.8.
- Enhancement 0.0 / +0.9.
- Arms 0.0 / +0.8.
- Protection warrior threat 0.0 / -1.2.
- Every other spec 0.0.

The item-data fixes reach published numbers only after data.yml and bis.yml regenerate simdb.bin and the BiS sets. integration.md lists the lane measurements.

Open going into round 5:
- **Owner decisions.**
  - /tiers/healer ranks healers side by side. The ×1.65 retune puts Restoration shaman 0.8% above Holy priest, where ×1.5 had the priest 7% ahead. The wow-player review asks that healer scores stay out of a cross-class order.
  - Power Infusion for healers is policy: it goes to a DPS. One self-cast would add at most about 1.7%.
- **The /sim page's consumables** are its own list, not the per-spec curated groups plus the class kit. Combat reads 7.6% low on the page with no poisons, and Survival 4.1% high (integration I6). The page needs per-spec resolution.
- **The BiS ranker's verify losses.** After the item-stat fixes, the ranker's weights pick sets its own verification scores lower: Survival, Arcane (Kindling Stave), Enhancement, Destruction (Elemental Mage Staff) and the tanks (defense items). Rhok'delar is not among Marksmanship's candidates.
- **Blocked on level-60 logs** (launch is 2026-11-04):
  - Demon crit (Blizzard's stated intent, not live; +1.7 to +2.6%) and haste.
  - The demon level-60 sheets.
  - Pet haste for hunters.
  - The rage-crit factor and Sunder's AP threat.
  - Holy Strike's flat bonus and Unsealed Wrath on special attacks.
  - Omen's 4% at 60.
- **Single source, one in-game look each:**
  - Ice Lance's 8 s on ranks 2 to 6.
  - Ice Barrier's +15% / +5%. The shipped tables say 3% / 3%; worth 9% of Frost.
  - Touch of the Grave's cooldown, 25 s or 15 s.
  - Blood Fury, 2 or 3 min.
  - Endurance's 2% crit.
  - The Human Spirit, 5 or 10%.
  - Pet hit from forever-bugs #394.
- **Carried over:**
  - Kings with Heart of the Lion (moot for the preset and the page now).
  - Shadowfiend at launch.
  - Pet family.
  - Demonic Brand unmodeled.
  - Siphon Life crit held for the owner.
  - Healing Stream tick crit.
  - Master of Elements on a Clearcasting cast.
  - Nine weapons on hotfix rows newer than the seed.
  - Expertise (37) and haste (36) item stats unconfirmed.
  - The fork's Settings suite stands melee at 30 yd.
- **Pre-existing:** five data tests (`test_addondata`, `test_audit_addon`) already fail on origin/main.

## Round 3 at a glance

Confirmed and fixed, by group:
- **Warlock (7).**
  - Demons refuse player buffs (round 2's held dfbb63ae3, as E e3bdac097).
  - Demons inherit Classic's shares: 75% Stamina, 30% Intellect, 35% armor, 56.5% of spell power as attack power and 15% as spell damage (E 94a45a77b). Two sources each: wowsims/sod for the same client spell 416189, and warcraft.wiki.gg or TBC. Resistances (40%) and hit (100%) are single-source. Crit is not inherited: the 29 low-level logs show pet crit far below the warlock's.
  - Lash of Pain takes Forever's 50, not Era's 99. Firebolt uses the client's range. Unholy Power is one pet-wide multiplier (E f6d77f02a).
  - Demons at bands 20 and 30 had no stats or weapon. They now have level-interpolated stats and their top legal rank (E d1e2536b9).
  - Destruction adopts Pandemic and Cataclysm for Unholy Power: +1.55% with inheritance (round 2's held build, S e4db33928).
- **Physical (6).**
  - Fury holds Death Wish and Recklessness for execute, unless that costs a use. It gains at every fight length from 30 to 300 s. Round 2's 60 s loss was the fork suite standing melee at 30 yd.
  - talent-search and rotation-search fought a Humanoid, while every other tool fights a target with no type. That was the Marksmanship gap (647.0 against 619.5).
  - Improved Tracking gave nothing on the published target, which has no creature type. Every Phase 1 boss is trackable, so it now counts (Marksmanship +4.2%, Survival +5.0%).
  - Murder joins the rogue raid keeps.
- **Hybrids (5).**
  - Shred was Era's 225% weapon damage; Forever's is 155% (client, Wowhead Forever, three beta log fits).
  - Claw is 110%.
  - Savage Fury follows its client mask.
  - Omen of Clarity is a flat 4% a spell cast, doubled in Moonkin Form, measured against 12,402 logged casts.
  - Judgement of the Crusader is in the raid preset (owner ruling).
- **Casters (3).**
  - No rotation casts Shadowfiend. A Forever priest cannot learn it: about 25 beta priests over 1,103 encounters never cast it, and its skill-line row has no learn level.
  - Every Ice Lance rank waits Fire Blast's 8 s category cooldown (client 70685 change, plus Wowhead Forever's text).
- **Ranker (1, from the wow-player review).** The BiS ranker priced weapon DPS from `items/<class>.json`, while the sim swings `simdb.bin`'s damage. About 200 weapons disagree. Elemental Mage Staff is 57.5 DPS in the item file and 42.7 in the engine, so it won the cat's main hand on damage the sim never deals. Candidates now take the engine's weapon row before scoring (S 425147a6f). In a local cat run, Hammer of the Titans replaces the staff: +4.8%. Every melee, hunter and hybrid weapon pick can move at the next bis.yml run.

Settled without a code change:
- Penance (0.19 a pulse) and Power Word: Shield (0.10): the client tables match the engine.
- Revenge rank 6 is an Ahn'Qiraj book rank, so not learnable in Phase 1.
- Venom with Vile Poisons stays additive.
- Elemental keeps Improved Reincarnation (owner ruling).
- Siphon Life crit is held for the owner, unchanged.

Before and after, pin c8d29ea33 against 640da215f on the same published data. Raid preset, band 60, mean of seeds 1 and 50001, 10000 iterations, alliance / horde:
- Feral -11.3 / -11.3.
- Restoration druid HPS -1.4 / -1.4.
- Marksmanship +4.2 / +4.2.
- Survival +5.0 / +5.0.
- Fury +0.8 / +0.4.
- Frost -0.9 / -0.9.
- Discipline HPS -8.1 / -7.8.
- Holy priest HPS -15.6 / -17.3.
- Retribution +9.3 / +11.3.
- Protection paladin threat +9.8 / +9.3.
- Affliction -7.0 / -7.0.
- Demonology -6.5 / -6.3.
- Destruction -6.4 / -6.3.
- Every other spec 0.0.

The published BiS files still carry the old Destruction build and the old weapon picks, and bis.yml regenerates them; the lane measured Destruction at -4.8% with its new build. The new order is Destruction, then Demonology, then Affliction.

Open going into round 4:
- Pet crit, haste, hit with Suppression, and Lash of Pain's own spell-power share need a level-60 warlock log with its sheet. The demon stats at bands 20 and 30 are interpolated.
- Kings with Heart of the Lion: the client cannot tell. A Forever character sheet with both buffs settles it.
- Ice Lance's 8 s on every rank: one source. Two Ice Lances inside 8 s in game settle it.
- Shadowfiend: if a launch trainer shows it, revert S 09848d599 and E 72aefbb89.
- Which weapon damage is true for Forever, `items/<class>.json` or `simdb.bin`? The tooltips probably show the item file's numbers.
- The /sim page defaults to a Humanoid target, so rogues show 4% above their published number. Its "raid-buffed" list (`web/src/lib/sim/settings.ts`) carries Kings, Sanctity Aura, Winter's Chill and Shadow Weaving and no Heart of the Lion, against the curated preset.
- The fork's Settings suite stands melee at 30 yd, so it cannot judge a melee rotation.
- Single-source rows:
  - Sunder's 5% AP threat.
  - The rage-on-crit factor.
  - Omen's 4%, measured at levels 20 to 30 only.
- Holy Strike's flat bonus under its weapon share.
- Power Infusion for healers.
- The Restoration druid covers 91% of incoming damage, above the heal profile's 85% ceiling.
- Demonic Brand is unmodeled.
- Wound Poison takes Venom and Vile Poisons against its client masks; no rotation casts it, so it moves no number today.
- Unsealed Wrath's reading.
- Wild Growth under a party-damage profile.
- Pet family.

## Round 3 wow-player review

SHIP: Balance, Feral bear, Restoration druid, the three hunters, Arcane, Fire, Holy paladin, Protection paladin, Retribution, Discipline, Holy priest, Shadow, Combat, Subtlety, Enhancement, Restoration shaman, and the three warriors. The reviewer held six, and each was settled before the push:
- **Feral, "586 DPS on a caster staff".** This is the ranker bug above (S 425147a6f). The next bis.yml run puts a real weapon on the band.
- **Assassination, "about 260 DPS on a mace and sword".** That came from a stale lane note. The published set is already on daggers after the nightly re-rank: 601 alliance, 615 horde.
- **Affliction, "the published build is the pre-round-2 build".** Also stale. The published band-60 build is round 2's Malevolence build.
- **Frost, "the 8 s Ice Lance rests on one source".** It is kept, because the client table is the truth, and it is listed under Open above. It costs 0.9% if wrong.
- **Demonology, "label pet scaling as an estimate".** The guide now says so.
- **Destruction, "-6.4% against the lane's -4.8%".** Both are right. -6.4% is measured on the published build, and -4.8% includes the new build, which bis.yml publishes.

The reviewer also flagged the Fury guide, which still said Death Wish goes on cooldown; it now describes the held cooldowns. Survival's 62 Wing Clips a fight are by design: the line fires only under a Windfury Totem, as a cheap special that rolls the totem's extra attacks, and its note records the measurement.

## Round 2 at a glance

Confirmed and fixed, by group:
- **Warlock (8).**
  - Wrack is a 6 s channel, and its +10% reaches only Corruption and Agony.
  - Nightfall and Shadow Mastery follow the client masks.
  - Affliction adopts Malevolence and Ruin.
  - Decimation applies only below 35% health, its cast cut is 40% of 4 s, and it no longer cuts mana.
  - Blood Fury is off the global cooldown.
- **Priest and mage (7, plus Shadowfiend).**
  - Darkness covers all Shadow damage.
  - Shadow Word: Death adds Shadow Weaving, takes its 30% backlash, and is cast on cooldown.
  - Dark Sacrifice is Undead-only.
  - Prayer of Mending spends a charge when its target is hit.
  - Ignite no longer double-dips target modifiers.
  - Shadowfiend is added below 30% mana.
- **Hybrids (5).**
  - The Feral form attack uses the weapon's DPS, for cat and bear.
  - The Ravage opener counts.
  - Fire Nova uses the Forever totem damage.
  - Retribution takes Unsealed Wrath and swaps seals.
  - Two client amounts are corrected: Heart of the Lion is 280 AP at 60, and Judgement of the Crusader is 161.
- **Physical (11).**
  - Mutilate needs daggers in the BiS ranker.
  - A stone sits beside a poison on rogues.
  - Revenge and Sunder scale with attack power.
  - The pre-pull cast lands at 0 and counts.
  - Lacerating Strikes no longer double-dips.
  - Subtlety adopts Quietus; Beast Mastery adopts Careful Aim 5.

Held:
- Demons refusing player buffs (lane E dfbb63ae3). Hunter pets already refuse them, and the client's warlock pet scaling spell matches the hunter one, but demons inherit no owner stats yet, so landing it alone would push pet damage further from the logs. It lands with a pet-inheritance model.
- Destruction's Pandemic and Cataclysm build (lane S 5d013a9c0). It passes the adoption rule only once demons refuse buffs.

Open going into round 3 (each report's "Still open" has the detail):
- Warlock pet inheritance.
- Siphon Life never crit in 772 low-level logged ticks, against the owner's DoT crit rule.
- Kings stacking with Heart of the Lion.
- Judgement of the Crusader in the raid preset.
- Omen of Clarity's rate.
- Fury cooldowns held for the execute phase.
- Venom with Vile Poisons.
- Murder's target type.
- Revenge rank 6.
- The rage-on-crit factor.
- Penance and Power Word: Shield coefficients.
- Ice Lance's cooldown on lower ranks.
- Shadowfiend's learnability at launch.
- talent-search and spec-breakdown give different baselines for Marksmanship.
- Elemental's Improved Reincarnation hold.

Real logs: no level-60 Forever logs exist before launch (2026-11-04). The 29 warlock logs at levels 13 to 30 on foreverlogs.gg agree that DoT ticks crit and that Life Tap scales with Spirit. Our API holds one public beta trash report. Round 3 starts on the first level-60 logs after launch, and warlocks are checked against real logs every round.

## Round 2 wow-player review

SHIP: Demonology, Destruction, Fire, Arcane, Frost, Protection warrior, Combat, Subtlety, the hunters, Feral, Elemental, Enhancement, Retribution and Protection paladin. The reviewer held three, and each was settled before the push:
- **Affliction, "Wrack is never cast".** The client data settles it. Wrack (1316697) is 36 a tick at 0.143 a tick over a 6 s channel, 0.143 coefficient a second, against Shadow Bolt's 0.343. Every Wrack rotation measured below the shipped one, wowtbc's Forever rotation verbatim included (-2.4%). Row A14 of warlock.md has the numbers.
- **Shadow, "Shadow Word: Death on cooldown kills the priest".** Two Forever guides (Warcraft Tavern Forever, wowtbc Forever) rank Shadow Word: Death above Mind Flay. The 30% backlash costs the healers, not the priest's damage. The rotation stays.
- **Assassination, "about 260 DPS on the page".** The ranker now needs daggers for Mutilate. The bis.yml run dispatched after the push re-ranks the set (576 on daggers in the lane run).
