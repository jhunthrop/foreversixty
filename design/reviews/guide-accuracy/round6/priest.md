# Guide accuracy round 6: Priest

Audited on the round 6 audit tree (land/guides-r5, HEAD 367cbadb7), client 1.60.1.70755, engine pin d1d9f2775 (`sim/enginever/version.go`).
Files: `web/src/content/guides/priest/index.md`, `discipline.md`, `holy.md`, `shadow.md`, the card notes in `data/curated/apl/priest-*.json`, and every frontmatter source link. Read first: round 1 to round 5 priest reports (with every fix log), round4/priest-levels.md, round2 to round5 race-sims.md and the ledger.

What moved since the round 5 audit tree (42f10dda9): the nightly BiS re-rank de044125c (all three `bis/priest-*.json`), the data push 43dc0f074 (Lesser Heal 2050 icon, item effect text whitespace, no priest talent, trainer or spellconst change), and the round 5 commit 2504f4a69 (discipline.md:58 and :88, holy.md:88, shadow.md:95 and :99, the first Mind Blast note in `priest-shadow.json`). `addon-builds.json` and every FS1 string are unchanged since round 2.

## Decoded builds and point orders (unchanged, all legal on 70755)

- Discipline `00500503133001143/3350500303/0` (gnome): Twin Disciplines 5, Mental Strength 5, Mental Agility 3, Inner Focus 1, IPWS 3, Meditation 3, Soul Warding 1, Penance 1, Renewed Hope 4, Divine Aegis 3 (29); Twilight Focus 3, Improved Renew 3, Holy Specialization 5, Divine Fury 5, Inspiration 3, Improved Healing 3 (22). Order: TD 10-14, MS 15-19, Med 20, MA 21-23, IF 24, IPWS 25-27, Med 28-29, Penance 30, SW 31, RH 32-35, DA 36-38, TF 39-41, IR 42-44, HS 45-49, DF 50-54, Insp 55-57, IH 58-60.
- Holy `00500503103/33505003030121431/0` (human): Disc 17 (TD 5, MS 5, MA 3, IF 1, Med 3), Holy 34 (TF 3, IR 3, HS 5, DF 5, Insp 3, IH 3, Holy Binding 1, Litany 2, SoR 1, Spiritual Guidance 4, Spiritual Healing 3, PoM 1). Prayer of Mending's prerequisite is Spirit of Redemption 1 (`talents/priest.json`), so holy.md:62 "the step to Prayer of Mending" holds.
- Shadow leveling `02510503103/0/54342001201300051` and raid `00510503103/0/52352101201300251` (human), orders as in round 5. The raid trade (Wand Specialization 2 + Blackout 2 for Early Demise 2, Shadow Reach 1, IMB 5th point) balances, shadow.md:69 holds.
- Power Infusion swap (discipline.md:60): Power Infusion's prerequisite is Penance 1; Renewed Hope 5 brings Discipline to 30 for tier 7, and Holy tiers 1 to 3 keep 17 points after Improved Renew drops to 1, above the 15 that Improved Healing needs. Legal.

## Verified clean this round (no finding)

**Stat prose against the tables the pages render now.** `GuideStatTable.astro` and the rail card both take band 60, the raid preset, at the faction of the first `recommendedRaces` entry (`band60Weights`, `selectBand`), and `railStatRows` sorts per-point stats by scale factor, then per-1% haste, then "Not significant".
- Discipline (gnome, so Alliance): mp5 1.17, healing power 1.00, Intellect 0.62, crit 0.61, Spirit 0.44 (relative to healing power), haste per 1% after them. discipline.md:88 "mp5, then healing power. Intellect and critical strike sit close together after it, and Spirit is last" matches. Crit 0.61 against Holy's 0.45 (both relative to healing power) supports "worth more to you than to Holy".
- Holy (human, Alliance): mp5 2.17, Spirit 1.10, healing power 1.00 (error 0.25), Intellect 0.63, crit 0.45, haste not significant. Horde Undead has healing power ahead of Spirit. holy.md:88 "Spirit and healing power come next, close together, and gear decides which leads" holds on both factions.
- `haste_on_items` is false on every priest band, so prose that leaves haste out is right.
- Shadow: does NOT match, see PRIEST-R6-1.

**Gear prose.** shadow.md:103, discipline.md:92 and holy.md:92 name no item; they give rules only (healing items carry a third of the bonus as damage, only the healing number heals, Meditation keeps Spirit working mid-cast). Nothing moves nightly.

**Leveling order on the page (1 to 19, then 20 to 60, strip under).** Shadow reads right: the 1 to 19 paragraph ends on Blackout to 18, and "From 19 the build takes Spirit Tap to 21" picks up the 19th point, with no gap and no repeat; the trainer levels (Smite 1, Shadow Word: Pain 4, Power Word: Shield 6, Renew 8, Mind Blast 10, Inner Fire 12, Psychic Scream 14, Heal 16, Dispel Magic 18, Holy Fire 20) and the wands (11287 required 5, 11288 required 13, Elven Wand from Bashal'Aran 957, Cinder Wand from Arugal's Folly 99) are as round 5 verified. On the two healer pages the first sentence of "Levels 20 to 60" was written before the 1 to 19 block existed and now reads against it (PRIEST-R6-2).

**Round 5 fixes re-checked.** shadow.md:95 (Eureka! fills with Mind Flay, skips Shadow Word: Pain and Devouring Plague, per R5-D1) holds. discipline.md:58 "most raid hits no longer push back your casts" is the 70% rank text. The new `priest-shadow.json` Mind Blast note ("Inner Focus makes that Mind Blast free, with 25% more crit chance") matches Inner Focus's rank text and the page.

**Everything else, line by line**, against `talents/priest.json` rank text, `races.json`, `trainables/priest.json` and the earlier rounds' spell checks: Spirit Tap 3/5 60% and Blackout 4/5 8%, 2/5 4% (re-read rank by rank this round); Renewed Hope 4/5 8% and 4 s; Spiritual Guidance 4/5 20% of Spirit; Twin Disciplines masks (Penance and Renew, not the shield; round 4); Mental Agility "Smite, Holy Fire, and instant cast spells"; Soul Warding removes the shield's 4 s cooldown; Divine Aegis 15%; Improved Healing names Lesser Heal, Heal, Greater Heal, Penance and Prayer of Mending; Holy Binding 50% on Flash Heal on others; Litany of Light; Spirit of Redemption 15 s of free heals; Shadowform, Shadow Weaving, Darkness, IMF, IMB, ISWP, Early Demise, Silent Resolve 10% a point, VE 20%. Racials (`races.json`): Ruthless 1% spell and attack crit, The Human Spirit 10%, Elune's Light 10% crit for 15 s on 3 min, Berserking 10% cast speed for 15 s on 3 min, Artisan's Precision 1% hit, Touch of the Grave up to 5% max health once per 15 s, Will of the Forsaken removes Charm, Fear and Sleep, Expansive Mind 5% mana, Eureka! 8 s on 2 min, Engineering Specialization 20% less likely to fail. Consumables exist in `items.json` (Major Spellblasting Potion 250937, Major Mender's Potion 250949). Race order prose against round 3 race-sims (held in round 4). Low-rank arithmetic as in round 5. No em or en dashes; no site or sim words outside the planner links.

**Card notes.** No new false note. The healer rank-1 notes ("Lesser Heal rank 1 ... for very little mana", "Heal is the filler: rank 1 in a raid") stay with the rank move (known and accepted, R3-D1, R4-D2).

**Source links.** Blizzard Deep Dive recap 200; Power Up Gaming 200; Talents Forever data.json 200 (labelled as beta client data); foreversixty.gg/sim/weights 200 (Shadow only, damage spec); /planner and /planner?class=priest 200; Icy Veins talent calculator news 403 to automated fetch (unknown, not dead, R5-D2). Every entry fits the round 4 source-link rule.

## Findings

| id | status | kind | file:line | wrong claim | evidence | exact fix |
|---|---|---|---|---|---|---|
| PRIEST-R6-1 | CONFIRMED | stats (re-rank, missed by the round 5 review pass) | shadow.md:99 (and frontmatter :9) | "Spell power and Shadow power are the same stat to you. Hit is next until you reach the 16% spell hit cap against a raid boss" | The table on the same page reads `bis/priest-shadow.json` band 60, raid, Alliance (Human is first in `recommendedRaces`). Scale factors: hit 4.71 (error 0.21), spell power and Shadow power 4.24 (error 0.004), crit 3.49, Intellect 1.00, then spell haste per 1% and spell penetration not significant. Weights relative to spell power: hit 1.11 ± 0.05 against 1.00 ± 0.001, so hit leads beyond its error. `railStatRows` sorts by scale factor, so the table, the pills and the rail card all show Hit first, then Spell power and Shadow power. The prose puts spell power first. Before the re-rank (round 5 audit) spell power led at 5.75 against hit 4.61, which is why this line was right then. The round 5 review pass rewrote the Discipline and Holy paragraphs for the re-rank but only touched Shadow's hit wording. Hit ahead of spell power while you are under the cap is also the stable Classic answer, so the new order will not move nightly. | "Hit comes first until you reach the 16% spell hit cap against a raid boss: Shadow Focus gives 5%, your gear finds the rest, and every miss is a lost Mind Blast. Spell power and Shadow power are the same stat to you and come right after it. Critical strike is worth more than it looks: in Forever your DoT ticks and Mind Flay crit, and Shadowform doubles every crit. Intellect is a little crit and a deeper bar; with raid buffs the mana rarely matters. Spell penetration does nothing against raid bosses." Frontmatter, so the fallback order matches: `statPriority: [Hit, Spell power, Shadow power, Critical strike, Intellect, Spell haste, Spell penetration]`. |
| PRIEST-R6-2 | LIKELY | leveling (page order since round 5) | holy.md:122, discipline.md:124 | Holy: "Level as Shadow and respec at 60, or level in a group as Holy." Discipline: "Level as Shadow and respec, or level in dungeons." | Both lines open "Levels 20 to 60" directly under a "Levels 1 to 19" block that has just spent points 10 to 19 in this spec's own order (Holy: Twilight Focus 10-12, Improved Renew 13-15, Holy Specialization 16-19; Discipline: Twin Disciplines 10-14, Mental Strength 15-19, per `addon-builds.json`), and the strip under them shows this spec's build at every band. Read in that order, "level as Shadow" tells a level 20 reader to drop the ten points the page just told them to take, without saying a respec at 20 is what that means. Not a false fact, so not CONFIRMED; it reads as a contradiction now that the two blocks sit together. | Holy: "From here you can keep leveling Holy in groups, or respec to Shadow at the trainer for faster solo questing and come back to Holy at 60." then the rest unchanged. Discipline: "From here you can keep leveling Discipline in dungeons, or respec to Shadow at the trainer for faster solo questing and come back at 60." then the rest unchanged. |

Counts: CONFIRMED 1, LIKELY 1, NEEDS-SIM 0, NEEDS-OWNER 0.

## Decisions (no owner questions)

- **R6-D1 Shadow stat order follows the rendered table** (PRIEST-R6-1): hit, then spell power and Shadow power, then crit, Intellect, spell penetration. Haste stays out of the prose: `haste_on_items` is false on every Shadow band, as for the healers (round 5 review pass).
- **R6-D2 Healer leveling openers** (PRIEST-R6-2): keep the "level as Shadow" option, since it is the faster solo route on both pages, but say it is a respec to Shadow from this point so it does not read against the points just given.
- **R6-D3 Card notes**: as R4-D2 and R5-D3. No change.

## Sims requested

None. Both findings are page text against committed data.

## Sim findings (engine vs game)

None new. Still open from earlier rounds, not re-filed: Starshards cooldown and a Shadow line (S-R3-1/2); Power Infusion on yourself (S-R2-1); the Forever low-rank healing curve and the healer card ranks with their notes; Shadow leveling Mind Flay at 20 and Spirit Tap 5/5; the Shadow raid Blackout 2 + Spirit Tap 3 move; the SW:P cutoff `remainingTime >= 10` to `>= 6`. The healer haste weights the ledger listed for the sim round are now large and positive (Discipline raid haste 5.0 per 1%), but haste is on no priest item, so no prose depends on it.

## Sources

- Repo data (70755): `data/builds/1.60.1.70755/bis/priest-{shadow,discipline,holy}.json` (band 60 weights, `haste_on_items`, `hit_to_cap`), `talents/priest.json`, `trainables/priest.json`, `races.json`, `items.json`, `addon-builds.json`; `data/curated/apl/priest-*.json`
- Site code: `web/src/components/guides/GuideStatTable.astro`, `web/src/lib/guides/rail-stats.ts`, `web/src/lib/guides/class-dps.ts` (`band60Weights`), `web/src/lib/bis/presets.ts` (`selectBand`), `web/src/lib/bis/weight-reading.ts`
- Commits: de044125c (nightly BiS re-rank), 43dc0f074 (data push), 2504f4a69 (round 5 fixes)
- Blizzard, World of Warcraft: Forever Deep Dive panel recap: https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap (200)
- Power Up Gaming, New WoW Forever talents: changes vs Classic: https://powerupgaming.co.uk/2026/09/14/new-wow-forever-talents-changes-vs-classic-guide/ (200)
- Talents Forever data export: https://talentsforever.com/data.json (200)
- Forever Sixty, stat weights: https://foreversixty.gg/sim/weights (200); character planner: https://foreversixty.gg/planner and https://foreversixty.gg/planner?class=priest (200)
- Icy Veins, Forever talent calculator news: https://www.icy-veins.com/wow-forever/news/start-planning-your-wow-forever-character-with-our-new-talent-calculator/ (403, unknown)

## Fix log (round 6)

- **Round 6 stat rule (all three spec pages).** Each `## Stat priority` paragraph now opens "The table above gives the order for your gear; here is why each stat counts." and gives only lasting reasons, no ranking words. Prose check `-t priest` passes.
  - discipline.md:88: mp5 (mana while casting), healing power (size of shield, Penance, heal), Intellect (Mental Strength +10%, rank text), crit (non-periodic crits leave Divine Aegis, rank text says non-periodic, so "direct heal"), Spirit (Meditation 50% at 3/3, rank text). Dropped "worth more to you than to Holy", which compared nightly weights.
  - holy.md:88: mp5, healing power, Spirit (Spiritual Guidance and Meditation 3/3, both in the Holy build), Intellect, crit (Inspiration only adds armor, rank text). Dropped "a point of it buys more healing than a point of healing power", a nightly comparison.
  - shadow.md:99: hit "matters up to the 16% spell hit cap and adds nothing past it" (Shadow Focus 5/5 in both builds), spell power and Shadow power, crit, Intellect, spell penetration unchanged in substance.
- **PRIEST-R6-1** (CONFIRMED): resolved by the new stat rule at shadow.md:99, not by re-ranking. Frontmatter `statPriority` left as is per integrator note.
- **PRIEST-R6-2** (LIKELY, verified against the page order): fixed as written. holy.md:122, discipline.md:124 now say keep leveling this spec, or respec to Shadow at the trainer and come back at 60.
- **R6-D1**: superseded by the stat rule (order lives in the table; haste stays out of prose). **R6-D2**: applied as above. **R6-D3**: no card note change.
- **Elsewhere on the pages**: no other stat or gear ranking by nightly value. Kept "on a long fight the mana is worth more" (mana potion against Mender's, a shared-cooldown rule) and shadow.md:103 "a healer's cloth piece still loses to a real caster piece" (the third-as-damage rule). index.md has no stat prose, unchanged.
- `updated: 2026-11-04` on discipline.md, holy.md, shadow.md. No APL notes changed. No new SIM findings.
