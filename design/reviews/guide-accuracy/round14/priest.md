# Guide accuracy round 14: Priest

Audited on the round 14 audit tree (main d7fc51fe7: thirteen guide rounds, sim round 16), client 1.60.1.70755, engine pin f637ec32c (`sim/enginever/version.go`; engine source read at scratchpad/gr14/engine).
Files: `web/src/content/guides/priest/index.md`, `discipline.md`, `holy.md`, `shadow.md`; the card notes in `data/curated/apl/priest-*.json`; every frontmatter and body link; and, new this round, the live best-in-slot files `data/builds/1.60.1.70755/bis/priest-{shadow,discipline,holy}.json` (de044125c picks), read band by band for both factions against item stats in `items/priest.json`. Read first: the round 1 to 13 priest reports and fix logs, round4/priest-levels.md, the round 2 to 5 race-sims files and the ledger through round 13.

## What moved since round 13

On the priest side, nothing. `git diff 604b78a80 d7fc51fe7` over the four guides, `data/curated/apl/priest-*.json`, `talents/`, `spelltips/`, `races.json` and `bis/` is empty, and the only `addon-builds.json` hunk is Demonology. The engine pin moved from 9f7f1f6e4 to f637ec32c. What did change is the rule: the provisional-BiS carve-out is over, so the prose is now held to the live BiS tables on every band and both factions.

## Decoded builds

These are the same as round 13 (see that report), and I re-derived them from `addon-builds.json`: Shadow leveling `02510503103/0/52542001201300051`, Shadow raid `00510503103/0/50052201211312251`, Discipline `00500503133001143/3350500303/0`, Holy `00500503103/33505003030121431/0`. The BiS files were ranked on older talent strings (Shadow raid `...-52352101201300251`, Discipline 40 band and below on a different order). That is the sim lane's open re-rank, not a prose item.

## Findings

| id | status | kind | file:line | claim | evidence | fix |
|---|---|---|---|---|---|---|
| PRIEST-R14-1 | CONFIRMED | gear | shadow.md:107 | "It still loses to a real caster piece: take healing gear only when nothing with spell power drops." | The page's own Shadow BiS table often picks a healing piece over a caster piece that has more spell damage. At 40, both factions: Papal Fez (22 healing, 8 damage) over Augural Shroud (11 damage) and Electromagnetic Gigaflux Reactivator (8). Inquisitor's Shawl (22 / 7) over Bloodmage Mantle (9 damage) and Green Silken Shoulders (8). Glowing Eye of Mordresh (20 / 7) over Triune Amulet (8). At 50 Horde, Glowing Eye of Mordresh again. At 60 bare Alliance: Torch of Light (17 healing, 6 damage, 11 Spirit) over Brilliant Wand (9 damage) and Greater Eternal Wand (9), and Mooncloth Boots (24 / 8) in the feet slot (`bis/priest-shadow.json` slots and alternatives; stats from `items/priest.json`). These healing pieces win on Intellect and Spirit, which is exactly what the line denies. The table renders every band (`GuideBisTable.astro`), so a reader at 40 sees the line contradicted. | "Forever puts a third of a healing item's bonus on it as spell damage, so a healer's cloth piece still does some damage. Read it by that damage number and its Intellect, never by the healing: a healing piece with good Intellect can beat a thin caster piece, but a real caster piece with the same Intellect always wins." |

Counts: CONFIRMED 1, LIKELY 0, NEEDS-SIM 0, NEEDS-OWNER 0.

## SHOULDs (wording, not findings)

| id | file:line | note |
|---|---|---|
| PRIEST-R14-S1 | shadow.md:103 | The raid band-60 table puts Spell haste first (both factions), and the paragraph gives a reason for every other stat it lists but not haste. That is not false, but a reader sees the top row go unexplained. Suggested addition, with the reason and no rank: "Spell haste shortens every Mind Blast and Mind Flay and every global, so it adds casts across the whole fight." |
| Carried | R13-S1 (index.md:48 Skyborne per faction), R13-S2 (shadow.md:91 15 s vs 6 s), R12-S1/S2 (healer "would overheal" card notes), R11-S2, R8-S1/S2, R9-S1 | Unchanged and still optional. |

## Verified clean this round

- **Gear and professions against the live BiS rows (both factions, every band).** Discipline and Holy BiS prose names no item. "Only the healing number heals" holds for every row. Shadow professions (:123): the raid 60 table takes six Tailoring pieces (Felcloth Hood, Mantle of the Timbermaw, Fel Cape, Robe of Winter Night, Gloves of Spell Mastery, Earthen Silk Slippers) and two Enchanting pieces (Frozen Heart of the Mountain, Brilliant Wand), so "a lot of the cloth you wear before raids" and "Enchanting makes its own caster pieces" hold. The Discipline and Holy Tailoring lines (:110, :108) hold too: Mooncloth Circlet, Mooncloth Shoulders, Truefaith Vestments, Wisdom of the Timbermaw, Swarmtender's Footpads, the Gilded and Nethergeld pieces, Mooncloth Boots.
- **Stat prose against the live band-60 tables.** Shadow raid (both factions): Spell haste, hit, spell power = Shadow power, crit, Intellect, with spell penetration, Spirit and mp5 at 0. That matches "Spell penetration does nothing against raid bosses", "with raid buffs the mana rarely matters", and spell power and Shadow power being the same stat. Discipline and Holy: every reason is still true, and no paragraph ranks stats. Holy Alliance and Horde order differently; the prose ranks neither.
- **Race lines against the race data.** `recommendedRaces` matches the prose picks: Shadow and Holy [human, undead], Discipline [gnome, undead]. The prose agrees with round3/race-sims.md, which is the latest priest measurement (round 4 and 5 race-sims have no priest rows). Racials match `races.json` and the race-table racial picks (Gnome caster and healer: Eureka! and Expansive Mind). No paid race is recommended.
- **Eureka! "never ... a channel" (index.md:38), re-checked.** The client damage and healing mask on 1259823 includes Penance's helper ids (`sim/core/eureka_auto_gen.go`). The heal helper 1316991's Forever tooltip reads "Healing for 425 every 1 sec for 2 sec", a periodic heal, so the "non-periodic" tooltip leaves Penance out. The line holds. The engine mask is the open S-R9-1 item.
- **Everything round 13 verified**, re-read line by line against unchanged data: talent text, spell values, leveling against `addon-builds.json`, low-rank arithmetic, card notes, opener, consumables, and the Skyborne line.
- **Source links** (fetched 2026-11-04): Blizzard Deep Dive recap 200; Power Up Gaming 200; Talents Forever data.json 200; foreversixty.gg /sim/weights, /planner, /planner?class=priest 200; Icy Veins 403 (unknown, R5-D2).
- No em or en dashes. No site or sim words outside the planner links.

## Decisions

- **R14-D1.** shadow.md:107 is a finding, not a BiS-side mismatch. It is prose making a rule about gear that the page's own table breaks, which the round 14 note says to file. The BiS race captions were not filed, per the brief.
- **R14-D2.** The Penance/Eureka! wording stays. Penance's heal helper is periodic by its client tooltip.

## Sims requested

Carried, no prose change waits on them:

- **Discipline race pick with the top-rank shield**: Gnome against Human, Undead against Troll; band 60, published gear, 10,000 iterations, seeds 1 and 50001.
- **Shadow race lines on the raid build** `50052201211312251`: Human, Night Elf, Dwarf, Gnome, Undead, Troll, same method.
- **BiS re-rank on the shipped builds** (post-launch). The Shadow table was ranked on `...-52352101201300251`, not the guide's raid build.

## Sim findings

None new. Still open: S-R2-1 (Power Infusion on self), S-R9-1 (the Eureka! masks: the engine puts Penance and Mind Flay in the damage list as well as the periodic list), X1, D2, W1, and the Shadow Word: Pain cutoff (APL 10 s against the page's 6 s).

## Sources

- Repo data (70755): `bis/priest-*.json`, `items/priest.json`, `talents/priest.json`, `spelltips/priest.json`, `addon-builds.json`, `races.json`; `web/src/components/guides/GuideBisTable.astro`, `RaceTable.astro`, `web/src/lib/guides/race-table.ts`
- Engine f637ec32c: `sim/core/eureka_auto_gen.go` (priest row), `sim/priest/penance.go`, `sim/priest/heal_ranks.go:200`
- Wowhead Forever tooltips: Penance helper https://nether.wowhead.com/forever/tooltip/spell/1316991 ; Eureka! https://nether.wowhead.com/forever/tooltip/spell/1259823
- Race measurements: `design/reviews/guide-accuracy/round3/race-sims.md`

## Fix log (round 14)

- PRIEST-R14-1: fixed, shadow.md:107. Shortened from the report text: "Judge it by that damage number and its Intellect, never by the healing: a healing piece with good Intellect can beat a thin caster piece." The "always wins" clause was dropped because Spirit can also tip a slot.
- PRIEST-R14-S1 (haste reason): not applied. Channel haste was not verified quickly.

All files were already at updated: 2026-11-04. The prose and sections tests pass (145/145).
