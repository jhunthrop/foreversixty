# Round 23, lane warrior (Arms, Fury, Protection)

Site `wt/r23-warrior` on int/r23 (77fc2cd56). Engine `lane/r23-warrior` on fork int/r23 (90176092d). Client 1.60.1.70790, still the newest Forever build on wago (`/api/builds`, product wowf, read 5 Nov 13:40 UTC).

**Result: 0 confirmed warrior sim bugs fixed; no engine, APL, rank-row, build or guide change.** One new lead no round had examined: Blizzard's 30 Sep class deep dive says Rend and Thunder Clap "now gain significant damage scaling from Attack Power". Neither the client nor the engine has a coefficient for either. The brief's rule for a 0 client coefficient applies: two independent log readings settle it, and there are none, so it is recorded OPEN (N2) with the exact log it needs. The rank sweep finds no lower rank that could win, and the engine casts the newest trainer rank at every level after a new one.

## Seed items (round 22 open)

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| S1 | **Defiance 7/13/20 against the client's 5/10/15.** Re-read 5 Nov 13:40 UTC. Wowhead's Forever tooltip for 12792 still reads "an additional 15%", and the one for 7376 still reads "[Defiance: 107)% / 80%]". There is no newer client build. Blizzard's thread still has 10 posts; the last is 4 Nov 19:35 UTC. | https://nether.wowhead.com/forever/tooltip/spell/12792, /7376, /71; wago `/api/builds`; https://us.forums.blizzard.com/en/wow/t/2360696.json | OPEN (carried: next client build or a threat capture) | - | - |
| S2 | **Threat constants with no client column.** Revenge is 2.25 x damage plus 2.25 x 2 x level, Shield Slam is a flat 508, and Thunder Clap is 2.5 x damage. All three came from SoD in upstream commit d30d87d9c (2024-05-06). At raid 60, Revenge is 545 of 1,453 threat/s, about 3,260 threat a cast against Shield Slam's 1,870. Blizzard's 30 Sep deep dive says Thunder Clap's threat was "rebalanced (reduced)" but gives no number. forever-bugs #63, #64 and #278 hold the only Forever threat captures (Sunder in beta, shouts at 0, rage gains at 5 threat a rage). The engine matches them where they overlap: block/dodge/parry rage gains come to 25.3 threat an event (`prot-raid-1.txt`). No capture of Revenge, Shield Slam or Thunder Clap exists. | revenge.go:15, shield_slam.go:16, thunder_clap.go:90; fork `git log -S` d30d87d9c; gh issues 63, 64, 278 (read 5 Nov 13:40 UTC) | OPEN (carried; a log with threat events for the three) | - | - |
| S3 | **Deep Wounds' attacker multiplier, applied twice** (`deep_wounds.go` OnTick). Nothing new: #434 has no answer, and no log exists. | forever-bugs #234, #434 | OPEN (carried; a 40+ log with Deep Wounds ticks beside a known weapon and Death Wish/Enrage windows) | - | - |
| S4 | **forever-bugs since round 22.** #449 (5 Nov) is a Cooldown Manager display request, not combat. #439 (Eureka! and Execute) is unchanged; the platform's P1b already holds its 30 Oct "neither effect" comment, and Gnome is not a published race. | gh issue list (5 Nov 13:40 UTC) | Aligned / carried | - | - |

## Blizzard's Forever notes, every warrior line

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| N1 | The thread has no post after 4 Nov 19:35 UTC (10 posts; post 1 last edited 3 Nov). Round 22 read every warrior line (its N1); nothing has changed since. | thread JSON, 5 Nov 13:40 UTC | Aligned (S1 open) | - | - |
| N2 | **New lead, not in the notes thread: Blizzard's class deep dive (30 Sep 2026).** It says "Rend: Now gains significant damage scaling from Attack Power" and "Thunder Clap: Now gains significant damage scaling from Attack Power ... Threat rebalanced (reduced)". In 70790 the client rows (11574, 11581) have ap_coefficient 0 and no APRatio in the description. Wowhead's Forever tooltips show only the flat numbers. The engine scales neither (rend.go and thunder_clap.go roll the client base only). Round 19's coefficient check (its X) compared the engine with the client and so missed the deep dive. Player readings: phobomancer (21 Sep, beta, level 20 against level 1 mobs) measured "Revenge 25% of attack power, Thunder Clap ~9%". The client later confirmed his Revenge figure exactly (70707's APRatio, 0.25), which supports his method, but ~9% is still one approximate reading. Joey (13 Oct, level 29) saw Rend at ~14 a tick against a 9 base, which shows scaling but fixes no ratio. The brief's rule (0 client coefficient: settle on two or more independent log readings, or record OPEN) leaves it OPEN. Reach: Thunder Clap is about 1% of Protection's threat, and Arms' and Fury's cards never cast it. Rend is 1.4% of Arms raid DPS today and more while leveling, and it feeds Bloodthrill. | https://news.blizzard.com/en-gb/article/24301514/world-of-warcraft-forever-class-deep-dives-priest-and-warrior (30 Sep); https://us.forums.blizzard.com/en/wow/t/2356894 (21 Sep); https://us.forums.blizzard.com/en/wow/t/2380853 (13 Oct); spellconst/spelltips 11574, 11581; https://nether.wowhead.com/forever/tooltip/spell/11574 | OPEN (needs a Forever log with Rend ticks and Thunder Clap hits at two known attack powers, or two players' readings) | old (round 19 checked only the client) | - |

## Rank sweep (every lane rule): first two levels after each new rank

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| R1 | **No lower rank can win any warrior row.** I compared every trainer rank with the rank below it on the 70790 rows, at the level it arrives and the level after. The spells are Heroic Strike, Rend, Overpower, Execute, Mortal Strike, Bloodthirst, Sunder Armor, Revenge, Shield Slam, Thunder Clap, Battle Shout, Demoralizing Shout, Cleave, Slam, plus Hamstring, Charge, Pummel, Shield Bash and Mocking Blow, which the cards or a pack cast. At each step the cost, global, cooldown and cast time are the same, and each effect is equal or larger. Durations are longer: Rend 9 to 21 s, Thunder Clap 10 to 30 s. Battle Shout and Demoralizing Shout carry points per level, and with them evaluated (the lower rank at its level cap), the new rank still leads at every step, for example Battle Shout 83.4 against 111 at 52. A rage spell that costs the same and hits harder cannot lose. The low-rank penalty only widens the gap. | `$R/warrior/tools/dominance.py` on `spellconst/warrior.json` and `trainables/warrior.json` | Aligned | - | - |
| R2 | **The engine casts the newest trainer rank at every one of those levels.** I ran the shipped leveling card at every learn level and learn level + 1 from 10 to 60, for all three specs (bare, horde, frozen BiS, the ladder build at that level, 1,000 iterations). I read each cast id against trainables, with the book ranks 25286, 25288 and 25289 excluded. There were 0 mismatches in 150 level runs (`out/sweep-{arms,fury,protection}.txt`). Cleave appears on no card. | spec-breakdown per level | Aligned | - | - |
| R3 | **Why rows cannot be simmed:** the warrior engine registers only the learned rank. A forced `-rank` of a lower Rend, Overpower, Execute, Heroic Strike or Battle Shout at 41 removes that spell from the fight (for example, Rend rank 4 at 41 gives 0 Rend casts). R1 shows that registering the lower ranks would add nothing, so no engine change is warranted. | sims at 41, 3,000 iterations | Aligned (no row possible or needed) | - | - |

## Arms

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| A1 | **Guide-lane sim: arms.md:53 "Rage starved without a raid around you".** Level 60 build, horde, 10,000 iterations, seeds 1 and 50001. Rage over a 180 s fight: raid 1,822 (10.1 a second) against bare 1,261 (7.0 a second), so the raid gives 44% more. By source (raid / bare): white swings 1,373 / 979 (Windfury Totem's extra swings and raid crit), refunds 124 / 70, Mighty Rage Potion 106 / 0, Unbridled Wrath 38 / 28, and Bloodrage, Anger Management and Charge the same. In a raid the spare rage goes to 34 Hamstrings for Windfury (342 rage) and 5 Heroic Strikes. Bare has no dump at all (0.3 Heroic Strikes). The line holds as written. | `out/arms-rage-{raid,bare}-{1,50001}.txt` | Aligned | - | - |
| A2 | Raid 60 on the guide raid build: 830.4 horde at seed 1, equal to `before.tsv`. | `out/arms-raid-1.txt` | Aligned | - | - |

## Fury

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| F1 | **Battle Stance from 30 to 35, before Whirlwind (`defaultStance`).** I tested a Berserker Stance opening with Berserker Rage from 30 and no Overpower against the card. Battle Stance wins by 5.1% at 30, 4.6% at 31, 2.4% at 33 and 1.1% at 35 (10,000 iterations, seeds 1 and 50001, both seeds agree). | `out/fury-bers30.txt` | Aligned | - | - |
| F2 | **Death Wish from 30** (round 22: talent spells at 10 + 5 x tier). The leveling build takes it at 30 (tier 5), and the card casts it 1.07 times a fight from 30. Reckless Swings comes at 31, at the point the build takes it. | sims at 29, 30, 31 and 36 | Aligned | - | - |
| F3 | **fury.md:112, Night Elf against Skyborne on the Alliance**, raid build, 10,000 iterations at seeds 1, 50001 and 100001: Night Elf 889.3, Skyborne 887.8 (Night Elf +0.17%). "Level" holds. | `out/fury-sky.txt` | Aligned | - | - |
| F4 | Raid 60: 878.2 horde at seed 1, equal to `before.tsv`. | `out/fury-raid-1.txt` | Aligned | - | - |

## Protection

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| P1 | **Guide-lane sim: protection.md:69, Shield Slam "Your biggest button".** Raid 60, guide build, 10,000 iterations, seeds 1 and 50001. An average landed Revenge (crits included) does 937 damage and an average landed Shield Slam 809, so Revenge is 16% bigger. Per cast, Revenge also brings about 1.7x Shield Slam's threat (545 against 313 threat/s on 30.1 against 30.2 casts), but S2 still leaves that ratio open. On damage alone the line is wrong. See "For the guide lane". | `out/prot-dmg-{1,50001}.txt`, `out/prot-raid-{1,50001}.txt` | Guide finding (no sim bug) | guides-handoff | - |
| P2 | **Guide-lane sim: protection.md:113, race TMI on one shared preset** (`-tank`, raid, 10,000 iterations, seeds 1 and 50001). On the published Alliance gear, Night Elf takes 259.0 damage a second against 263.8 to 265.3 for every other Alliance race, and its TMI is 22.85 against Dwarf's 23.1. On the Horde gear for all races, Tauren's TMI is 23.25 against 24.3 to 24.7 for every other race, Night Elf included. Both sentences hold. | `out/prot-races.txt` | Aligned | - | - |
| P3 | **Last Stand from 20** (round 22 rule). The build takes it at 28 (tier 3); the card's line is health-gated and never fires in the raid fight. Raid 60: 1,453.2 threat/s horde at seed 1, equal to `before.tsv`. | `out/prot-raid-1.txt` | Aligned | - | - |

## Cross-cutting

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | **Round 22's 3% raid crit as it touches warriors.** The tank sheet shows "crit (every table) 3.00" in the raid preset: Leader of the Pack once, with Moonkin Aura not added on top. Warriors take no spell crit, so neither aura's spell part matters. The scrolls fix (145b1fdd3) and the unknown-aura change (30bf8bdb0) move no warrior number: every value equals `before.tsv`, and no warrior APL condition reads a talent aura before its level. | sims; fork log 170a59c39..bff826207 | Aligned | - | - |
| X2 | **Single-rank trainer levels against the engine:** Berserker Rage 30, Berserker Stance 30, Retaliation 20, Shield Block 16, Shield Wall 28, Whirlwind 36, Recklessness 50, Sweeping Strikes 30, Bloodrage 10. All match. | trainables 70790; `constants_auto_gen.go`, `berserker_rage.go:13`, `retaliation.go:12` | Aligned | - | - |

## Before/after

No change in this lane. On this branch, horde seed 1 at 10,000 iterations: Arms 830.4, Fury 878.2, Protection 1,453.2 threat/s. Each equals `$R/before.tsv` to the decimal.

## Files to regenerate

None.

## Measure at refresh

- Carried: the Protection BiS on the 80% stance, and the warrior BiS after the CI re-rank.

## For other lanes

- **Platform:** Mighty Rage Potion's rage gives 0 threat in the engine (`item 13442`, `prot-raid-1.txt`). forever-bugs #278 shows potion rage making some threat, less than 5 a rage (69 threat from a 20 to 40 rage potion). It is worth about 0.1% of Protection threat, so I made no change.
- **Platform / data:** N2 is the same shape as Revenge's 25% before 70707: server-side attack-power scaling that the client leaves out. If a client build adds an APRatio to Rend or Thunder Clap, `spellconst` should carry it.

## For the guide lane

- **protection.md:69** "Shield Slam ... Your biggest button": on damage, an average landed Revenge (rank 5, Improved Revenge 3) hits 16% harder than Shield Slam at raid 60. Suggested: drop "Your biggest button", or say what Shield Slam is biggest at (it is the hit you control, on a 6 second cooldown, and it does not wait for a block, dodge or parry).
- arms.md:53 (A1), protection.md:113 (P2) and fury.md:112 (F3) hold as written; their carried sim requests are closed.

## Guide text not done

- None (no sim-side change).

## Still open

- S1 Defiance 7/13/20; S2 Revenge, Shield Slam and Thunder Clap threat; S3 Deep Wounds' double multiplier; N2 Rend and Thunder Clap attack-power scaling (Blizzard 30 Sep; ratio unknown); round 22's trainer book ranks and WARRIOR-13 (owner).
- **Logs, last looked 5 Nov 13:38 UTC:** api.foreversixty.gg `/v1/reports/recent` lists only the 20 Oct report. foreverlogs.gg `/api/reports/public` answers 403; its home page is a bare single-page app. Warcraft Logs v1 `/v1/zones` answers 401 "Invalid key specified" on www, classic, fresh, vanilla and sod. One Forever log with Rend ticks, Thunder Clap hits, Revenge/Shield Slam threat or Deep Wounds ticks would settle N2, S2 or S3.

## Sources

- Blizzard: https://us.forums.blizzard.com/en/wow/t/2360696 (10 posts, read 5 Nov 13:40 UTC); class deep dive Priest and Warrior, https://news.blizzard.com/en-gb/article/24301514 (30 Sep 2026; Wowhead copy https://www.wowhead.com/forever/news/priest-and-warrior-class-deep-dives-in-world-of-warcraft-forever-383192).
- Players: https://us.forums.blizzard.com/en/wow/t/2356894 (21 Sep), /t/2380853 (13 Oct), /t/2375342 (7 Oct, Thunder Clap rolls spell crit, which matches the engine's OutcomeMagicHitAndCrit).
- Client 1.60.1.70790: `spellconst`, `spelltips`, `trainables`, `talents`, `addon-builds` (warrior); wago `/api/builds`.
- Wowhead Forever tooltips 12792, 7376, 71, 11581, 11574, 6343 (5 Nov 13:40 UTC).
- forever-bugs: #449, #439, #434, #278, #63, #64, #66, #121, #324; warrior search sorted by update (5 Nov 13:40 UTC).
- Guide-accuracy rounds 22 to 25 (warrior), "Sims requested" and "Engine questions".
- Lane tools and outputs: `$R/warrior/tools/{dominance,levelsweep,ladder}.py`, `$R/warrior/out/`, `$R/warrior/rot/`.
