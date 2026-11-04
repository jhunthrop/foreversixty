# Guide accuracy round 9: Priest

Audited on the round 9 audit tree (land/guides-r8, HEAD 645caa983), client 1.60.1.70755, engine pin d1d9f2775 (`sim/enginever/version.go`).
Files: `web/src/content/guides/priest/index.md`, `discipline.md`, `holy.md`, `shadow.md`, the card notes in `data/curated/apl/priest-*.json`, every frontmatter and body link. Read first: round 1 to round 8 priest reports and fix logs, round4/priest-levels.md, the race-sims files and the ledger through round 8.

What moved since the round 8 audit tree (09b98d9e4): only the round 8 commit on the priest pages. `git diff 09b98d9e4..HEAD` over the priest guides, `data/curated/apl/priest-*`, `addon-builds.json`, `talents/`, `trainables/`, `spelltips/` and `races.json` shows index.md:5 (`updated: 2026-11-04`), index.md:38 and shadow.md:115 (the PRIEST-R8-S3 Eureka! rewrites). Every FS1 string and point order is unchanged since round 2; the decoded builds in round8/priest.md hold.

## Round 9 focus checks

**Crit-damage claims** (the +100% bonus rule). Shadowform's 70755 text: "increasing the critical strike damage bonus of your Shadow spells by 100%". A spell crit is 150% of a hit, so Shadowform doubles the +50% bonus and a Shadow crit becomes 200% of a hit, a third more than a normal crit, not double a crit. Engine agrees (`sim/priest/talents.go:218` shadowformCritDamageBonus 1.0 on the bonus; `shadow/shadowform_mind_blast_test.go`). shadow.md:54 "double the critical damage bonus" and shadow.md:111 "doubles every crit's bonus" are right; shadow.md:99 is not (R9-3). No other priest talent in either build raises a crit bonus; Divine Aegis and Inspiration are crit-triggered effects and are stated correctly.

**Eureka! lines against the client description.** Wowhead Forever tooltip, Eureka! (priest) 1259823, read 2026-11-04: "Your non-periodic damaging or healing abilities have their Mana cost reduced by 10% and deal 10% more damage or healing for 8 sec." Seven priest lines checked:
- index.md:35: names the racial only. Holds.
- index.md:38 (R8 rewrite): "heals and damage spells ... never a DoT, Renew or a shield". Holds against the tooltip.
- shadow.md:95: Eureka! on Mind Blast and Shadow Word: Death, fill with Mind Flay, skips the DoTs. Rotation advice; it claims no boost for Mind Flay. Holds.
- shadow.md:115 (R8 rewrite): names Mind Flay as boosted. Mind Flay is a channelled periodic spell, outside "non-periodic" (R9-1).
- discipline.md:100 and holy.md:104: "your heals", which takes in Renew, the heal both pages keep on the tank; index.md:38 says the opposite (R9-2).
- holy.md:84: Eureka! on a string of Prayers of Healing. Holds (direct heal; in the client damage and cost masks).

## Verified clean this round (no finding)

Every other line, re-read against 70755:
- **Talent text**: Twin Disciplines "instant cast spells" (Penance and Renew are in its client heal mask, engine `PriestSpellMaskInstantHealing`), Mental Agility, Inner Focus "non-periodic" 25%, Soul Warding 4 s and 15%, Penance "instantly and every 1 sec for 2 sec", Renewed Hope, Divine Aegis "non-periodic critical heals" 15%, Improved Healing, Divine Fury, Holy Binding "Flash Heal" 50% on others, Litany of Light "previous heal was a different spell", Spirit of Redemption 15 s free heals, Spiritual Guidance, Spiritual Healing 10%, Prayer of Mending "Jumps up to 5 times", Twilight Focus 70%, Improved Renew 15%, Shadow Weaving, Improved Mind Flay 20%, Spirit Tap, Blackout 2% a point, ISWP 6 s, Early Demise 30%, Silent Resolve 10% a point, Shadow Reach, Wand Specialization, Vampiric Embrace 20% for 30 s, Meditation 50%, Mental Strength 10%, Power Infusion 20% for 15 s.
- **Leveling** (1 to 19, then 20 to 60, strip under): talent sentences match `addon-builds.json` orders (Shadow SF 10-14, Blackout 15-18, Spirit Tap 19-21, IMB to 25, ISWP 26-27, Mind Flay 28, then IMF, VE, Shadow Weaving, Darkness, Shadowform 40; Discipline TD 10-14, MS 15-19, Penance 30; Holy TF 10-12, IR 13-15, HS from 16). Trainer levels, wands and quests as round 8 verified.
- **Races, consumables, professions, low-rank prose, stat prose reasons, card notes**: unchanged since round 8 and still true. No ranking words in Stat priority beyond the table.
- **Source links** (fetched 2026-11-04): Blizzard Deep Dive recap 200; Power Up Gaming 200; Talents Forever data.json 200 (labelled as beta client data); foreversixty.gg/sim/weights 200 (Shadow only); /planner and /planner?class=priest 200; Icy Veins talent calculator news 403 (unknown, not dead, R5-D2).
- No em or en dashes; no site or sim words outside the planner links.

## Findings

| id | status | kind | file:line | wrong claim | evidence | exact fix |
|---|---|---|---|---|---|---|
| PRIEST-R9-1 | CONFIRMED | consumable/race | shadow.md:115 | "Eureka! makes Mind Blast, Shadow Word: Death and Mind Flay 10% cheaper and 10% stronger for 8 seconds every 2 minutes" | Round 8 fix (R8-S3) wrote this from the engine's mask table. The client description (Wowhead Forever 1259823) covers "non-periodic damaging or healing abilities"; Mind Flay is a channel whose damage comes in periodic ticks (Improved Mind Flay: "Each instance of damage of Mind Flay"). The ledger's round 8 ruling: the guide follows the tooltip, and the warlock page was cut back the same way (WARLOCK-R8-1; its fix log refused to give channels the boost). The class mask for 1259823 does name Mind Flay (sim finding below), so the safe line names only the direct spells. | "**Gnome** trails them a little: Eureka! makes your direct damage spells, Mind Blast and Shadow Word: Death, 10% cheaper and 10% stronger for 8 seconds every 2 minutes, and Expansive Mind adds 5% mana." |
| PRIEST-R9-2 | CONFIRMED | consumable/race | discipline.md:100, holy.md:104 | Discipline: "Eureka! makes your heals 10% cheaper and 10% stronger for 8 seconds every 2 minutes"; Holy: "Eureka! makes your heals 10% stronger and 10% cheaper for 8 seconds every 2 minutes" | Tooltip: "non-periodic ... healing abilities". Renew is a heal over time, and both pages keep it on the tank all fight (discipline.md:70, holy.md:68). The client row's masks leave out Renew, Power Word: Shield and Prayer of Mending (`sim/core/eureka_auto_gen.go`, priest row: none of the three in the damage, periodic or cost lists). The class page says the opposite: index.md:38 "never a DoT, Renew or a shield". The page set contradicts itself. | discipline.md:100: "Eureka! makes your direct heals 10% cheaper and 10% stronger for 8 seconds every 2 minutes, never Renew or a shield: put it on a string of Greater Heals and Flash Heals." holy.md:104: "Eureka! makes your direct heals 10% stronger and 10% cheaper for 8 seconds every 2 minutes, never Renew, and Expansive Mind adds 5% mana." |
| PRIEST-R9-3 | CONFIRMED | stats | shadow.md:99 | "in Forever your DoT ticks and Mind Flay crit, and Shadowform doubles every crit." | Shadowform (15473, 70755): "increasing the critical strike damage bonus of your Shadow spells by 100%". The bonus doubles (+50% to +100%), so a crit goes from 150% to 200% of a normal hit. That is a third more than a normal crit, not double. Engine: `sim/priest/talents.go:218` adds 1.0 to the crit damage bonus. The same page states it correctly at :54 ("double the critical damage bonus") and :111 ("doubles every crit's bonus"). This is the round 7 and 8 lesson (Impale, Mortal Shots). | "Critical strike counts for more than it looks: in Forever your DoT ticks and Mind Flay crit, and Shadowform doubles the crit bonus, so a Shadow crit hits for twice a normal hit." (No STAT_RANKING word.) |

Counts: CONFIRMED 3, LIKELY 0, NEEDS-SIM 0, NEEDS-OWNER 0.

## SHOULDs (wording, not findings)

Round 8's S1 (Twin Disciplines and Mental Agility reach Prayer of Mending, holy.md:59), S2 (Spiritual Healing "on everything", holy.md:51) and S4 ("five heals", holy.md:30) still apply and were not taken in round 8. One more:

| id | file:line | text | suggested wording |
|---|---|---|---|
| PRIEST-R9-S1 | shadow.md:65 | "your instants hit harder and cost less" | Mental Agility's client cost mask leaves out Mind Flay and Shadow Word: Death (`sim/priest/priest.go` PriestSpellMaskMentalAgility comment, SpellEffect 687728), so "cost less" covers Shadow Word: Pain, Devouring Plague and Vampiric Embrace only. True as a summary. Optional: "your instants hit harder, and your DoTs and Vampiric Embrace cost less". |

## Decisions (no owner questions)

- **R9-D1 Eureka! prose follows the client description** (non-periodic damaging or healing abilities), as the ledger ruled in round 8 for warlock. Name only the direct spells; never claim a boost for a channel, DoT, Renew, shield or Prayer of Mending, and never claim Mind Flay is excluded either while the client masks and the tooltip disagree.
- **R9-D2 Crit wording**: Shadow crits are 200% of a normal hit in Shadowform. Say "doubles the crit bonus" or "twice a normal hit", never "doubles every crit".
- **R9-D3 Frontmatter `statPriority`**, card notes, and the healer rank-1 notes: as R6-D2 to R8-D3. No change.

## Sims requested

None needed for the fixes. Optional, wording only: Shadow Gnome vs Human on the raid build with the Eureka! periodic modifier (Mind Flay, misc 22) turned off, to confirm "trails them a little" still holds if the tooltip is the real behaviour. The direction can only get worse for Gnome, so the current wording stays true either way.

## Sim findings (engine vs game)

- **S-R9-1 Eureka! (priest) masks against the tooltip.** The client row 1259823 carries three aura 108 effects: misc 14 cost -10%, misc 0 damage or healing +10%, misc 22 periodic damage +10%. The engine (`sim/core/eureka_auto_gen.go`, priest row) puts Mind Flay, Penance and Starshards in the periodic list as well as the damage and cost lists, and leaves out Prayer of Mending, a non-periodic heal by the tooltip. The tooltip says non-periodic only. This is the same open item as the warlock row (ledger round 8). Decide which one the engine follows; the guide follows the tooltip.
- Still open, not re-filed: Starshards cooldown and a Shadow line (S-R3-1/2); Power Infusion on yourself (S-R2-1); the Forever low-rank healing curve and the healer card ranks; Shadow leveling Mind Flay at 20 and Spirit Tap 5/5; the Shadow raid Blackout 2 + Spirit Tap 3 move; the SW:P cutoff `remainingTime >= 10` to `>= 6`; healer haste weights.

## Sources

- Repo data (70755): `data/builds/1.60.1.70755/{talents,spellconst,trainables}/priest.json`, `spellranks.json`, `races.json`, `addon-builds.json`; `data/curated/apl/priest-*.json`
- Engine d1d9f2775: `sim/core/eureka_auto_gen.go` (priest row 1259823), `sim/core/racials.go:433`, `sim/priest/talents.go:218`, `sim/priest/shadow/shadowform_mind_blast_test.go`, `sim/priest/priest.go` (Twin Disciplines and Mental Agility masks)
- Wowhead Forever tooltip, Eureka! (priest) 1259823: https://nether.wowhead.com/forever/tooltip/spell/1259823 (read 2026-11-04): "Your non-periodic damaging or healing abilities have their Mana cost reduced by 10% and deal 10% more damage or healing for 8 sec."
- Ledger round 8 and round8/warlock.md (WARLOCK-R8-1 and its fix log): the Eureka! ruling
- Blizzard, World of Warcraft: Forever Deep Dive panel recap: https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap (200)
- Power Up Gaming, New WoW Forever talents: changes vs Classic: https://powerupgaming.co.uk/2026/09/14/new-wow-forever-talents-changes-vs-classic-guide/ (200)
- Talents Forever data export: https://talentsforever.com/data.json (200)
- Forever Sixty, stat weights: https://foreversixty.gg/sim/weights (200); character planner: https://foreversixty.gg/planner and https://foreversixty.gg/planner?class=priest (200)
- Icy Veins, Forever talent calculator news: https://www.icy-veins.com/wow-forever/news/start-planning-your-wow-forever-character-with-our-new-talent-calculator/ (403, unknown)

## Fix log (round 9)

Applied by the integrator; see the round 9 ledger entry and the commit for the exact lines.
