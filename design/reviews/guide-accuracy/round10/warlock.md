# Guide accuracy round 10: Warlock

Auditor: warlock lane, round 10. Tree: gr10/audit at 84460cec8 (all nine guide rounds plus sim round 15), client 1.60.1.70755, engine pin 9f7f1f6e4 (`sim/enginever/version.go`).
Audited line by line: `web/src/content/guides/warlock/index.md`, `affliction.md`, `demonology.md`, `destruction.md` (frontmatter, every body line, Levels 1 to 19 and Levels 20 to 60), the curated APL notes in `data/curated/apl/warlock-*.json` and the rendered cards in `data/builds/1.60.1.70755/addon-data.json` `rotations` (bands 10 to 60), and every source link.
Read first: `round1/warlock.md` to `round9/warlock.md` (every fix log and decision), `round4/warlock-levels.md`, the race-sims files, `ledger.md`, and sim round 15's `design/reviews/2026-11-04-sim-audit-round15/warlock.md` (waves 1 to 10).

## What changed since round 9 (645caa983..84460cec8)

- **affliction.md:7 raidBuild** is now `25520020135001351/00052/005004` (sim round 15 A1). Only the frontmatter line changed; no Affliction prose moved. Sim round 15's "For the guide lane" list (its replacements for the raid-trade, Wrack and filler sentences) was not applied when the merge took int/r15's build line.
- **warlock-affliction.json**: prepull Amplify Curse removed; Curse of the Elements before Corruption; new Bane of Doom line (60 s or more left), new Amplify Curse line (no Doom, no Agony, under 60 s), Bane of Agony gated on no Doom; new Wrack line above Shadow Bolt. The three new lines carry no note; the Agony and Shadow Bolt notes were not touched.
- **warlock-demonology.json**: Soul Fire gated on no Demonic Pact (sim round 15 G3); Curse of the Elements before Corruption. **warlock-destruction.json**: Curse of the Elements first.
- `addon-builds.json` warlock-affliction raid order regenerated for the new build; leveling orders unchanged for all three specs. `talents/warlock.json`, trainables and spellconst unchanged.

## Builds

New Affliction raid `25520020135001351/00052/005004` (tier, column order), 51 points, every gate held (Malevolence after 20 Affliction points, Wrack after 34 with Siphon Life, Aftermath after Bane 5):
Affliction 35: Improved Life Tap 2, Suppression 5, Improved Corruption 5, Grim Reach 2, Improved Bane of Agony 2, Amplify Curse 1, Pandemic 3, Malevolence 5, Siphon Life 1, Soul Siphon 3, Shadow Mastery 5, Wrack 1. Demonology 7: Unholy Power 5, Demonic Aegis 2. Destruction 9: Bane 5, Aftermath 4.
Against the leveling build `25522300130201351/200522/005`: out Soul Harvest 2, Improved Drains 3, Nightfall 2, Improved Health Funnel 2, Improved Voidwalker 2 (11); in Improved Bane of Agony 2, Malevolence 5, Aftermath 4 (11). Soul Siphon, Wrack, Grim Reach and Bane 5 stay.
Talent text at these ranks (`talents/warlock.json`): Improved Bane of Agony 2 "by 10%"; Aftermath 4 "initial damage of your Immolate spell by 40%"; Bane 5 0.5 s off Shadow Bolt, Immolate and Incinerate; Soul Siphon 12% per other Affliction effect, up to 36%; Pandemic covers Wrack.
The other five strings are unchanged and decode as round 9 lists them.

## Checked and correct (no finding)

- Demonology and Destruction prose already give Curse of the Elements first (demonology.md:78, destruction.md:69) and match the reordered APLs. demonology.md:84 already skips Soul Fire in the raid build, which the new Pact gate now matches.
- Decimation's 6% (demonology.md:33, :56, :84, :110): sim round 15 M1 found the rank curve (TraitDefinitionEffectPoints 135428, curve 110390: 3 / 6), so round 9's sim finding is closed and the prose stands.
- affliction.md:52, :65 (last sentence), :67: Amplify Curse on the last-minute Agony and Doom while a minute or more is left now match the APL exactly.
- affliction.md:54 Nightfall: still dropped by the raid build, and Malevolence is still in it.
- Leveling: every Levels 1 to 19 and 20 to 60 sentence on the three pages still matches the unchanged leveling orders (index 0 = level 10), as round 9 listed.
- Races: no Skyborne warlock in `combos.json`; index.md:35 holds. Racial text unchanged and right.
- Links (fetched 2026-11-04): Blizzard Deep Dive recap 200, wago.tools SpellEffect 70755 200, But Why Tho 200, foreversixty.gg/sim/weights 200; labels match.
- Voice: no dashes, no site or sim talk, no nightly numbers, no stat ranking.

## Findings

| id | status | kind | file:line | wrong claim (quoted) | evidence | exact fix |
|---|---|---|---|---|---|---|
| WARLOCK-R10-1 | CONFIRMED | talent | affliction.md:48 | "The raid build is a Shadow Bolt build that keeps its DoTs up: it stops at 30 points in Affliction and buys Improved Shadow Bolt and Ruin in Destruction." | Decoded raidBuild (:7): Affliction 35, Destruction Bane 5 and Aftermath 4, no Improved Shadow Bolt, no Ruin. Sim round 15 A1 adopted it because the rotation casts no Shadow Bolt once Wrack is the filler. | "The raid build is a Wrack build too: it goes to 35 points in Affliction for Malevolence, Soul Siphon and Wrack, and spends its Destruction points on Bane and Aftermath for a faster, harder Immolate." |
| WARLOCK-R10-2 | CONFIRMED | talent | affliction.md:57 | "Cheap and strong solo. In a raid it trails Shadow Bolt without Soul Siphon and is close to it with Soul Siphon." | Both builds now take Soul Siphon and Wrack. 70755: Wrack 86 a second at 60 (spellconst 1316697: 36 + 2.5 a level, 0.286 a tick), 200 mana for 6 s, times 1.36 with three effects up; Shadow Bolt r9 370 mana a cast. Sim round 15 A1: Wrack beats Shadow Bolt per second of casting at about a quarter of the mana; the Wrack build gains 13.6% at both seeds. | "With three of your effects on the boss it out-damages Shadow Bolt for a fraction of the mana, so both builds take it and it is your filler from 40." |
| WARLOCK-R10-3 | CONFIRMED | talent | affliction.md:59 | "The raid build trades Nightfall, Soul Siphon, Wrack, Soul Harvest and the three points of Improved Drains ..., along with Improved Health Funnel and Improved Voidwalker, for Malevolence, Improved Shadow Bolt and Ruin. It also takes one point of Improved Bane of Agony for 5% more ... Improved Shadow Bolt makes ... Four points of Ruin raise the crit bonus of your Shadow Bolt and Immolate by 80%." | Decoded trade above: Soul Siphon and Wrack stay; Improved Bane of Agony is 2 points (10%); Aftermath 4 is in; Improved Shadow Bolt and Ruin are out. Aftermath 4 text: Immolate's initial damage +40%. | "The raid build drops Nightfall, Soul Harvest and the three points of Improved Drains (a raid rotation casts no drain), along with Improved Health Funnel and Improved Voidwalker. It takes Malevolence and both points of Improved Bane of Agony, for 10% more damage on the Agony that covers the last minute, and keeps Grim Reach for range. Malevolence gives your Shadow spells 5% crit, DoT ticks and Wrack included. In Destruction, Bane makes Immolate a 1.5 second cast and four points of Aftermath add 40% to its first hit. Its Demonology points are Unholy Power for your Succubus and Demonic Aegis for Demon Armor." |
| WARLOCK-R10-4 | CONFIRMED | rotation | affliction.md:71 | "This build does not take Wrack. Without Soul Siphon behind it, Wrack is a 6 second channel that Shadow Bolt beats, so Shadow Bolt is your filler. The Level 60 build above is the one that channels Wrack ... so in a raid take this build and hard cast Shadow Bolt." | Raid build takes Wrack 1 and Soul Siphon 3; `warlock-affliction.json` casts Wrack (1316697 r1) right above Shadow Bolt; the 60 card lists Wrack before Shadow Bolt. | "Wrack is your filler. Once Corruption, your bane and Siphon Life are up, Soul Siphon adds 36% to it, and while it channels your other Shadow DoTs take 10% more. It does more than Shadow Bolt for a fraction of the mana, so channel it whenever nothing needs refreshing." |
| WARLOCK-R10-5 | CONFIRMED | rotation | affliction.md:65 | "Then Corruption, Bane of Doom, Siphon Life and Immolate, and Shadow Bolt from then on." | APL order after the DoTs: Wrack, then Shadow Bolt (never reached at 60 with the talent). | "Then Corruption, Bane of Doom, Siphon Life and Immolate, and Wrack from then on." |
| WARLOCK-R10-6 | CONFIRMED | rotation | affliction.md:27, :43 | ":27 "and fills the gaps with Shadow Bolt"; :43 "Role: Ranged Shadow damage, DoTs plus Shadow Bolt"" | Same as R10-4: both published builds fill with Wrack at 60. | :27 "Affliction keeps three DoTs and a bane rolling on the boss and fills the gaps with Wrack, and every one of those ticks can crit in Forever." :43 "**Role:** Ranged Shadow damage, DoTs plus Wrack" |
| WARLOCK-R10-7 | LIKELY | rotation | affliction.md:79 | "Your last Bane of Agony wants most of its 24 seconds, since its big ticks come last. Cast Shadow Bolt instead." | With the DoTs that are still running on the boss, Soul Siphon keeps Wrack ahead of Shadow Bolt (86 + 0.286 spell power a second, times 1.12 per other effect, against Shadow Bolt's 370-mana 2.5 s cast); the page's filler is Wrack after R10-4. Only once the target carries nothing of yours does Shadow Bolt pull ahead. | "Your last Bane of Agony wants most of its 24 seconds, since its big ticks come last. Channel Wrack instead while your other DoTs are still on the boss." |
| WARLOCK-R10-8 | CONFIRMED | consumable/race | affliction.md:97 | "Gnome is a step behind: Eureka! gives ... on your direct damage spells every two minutes, but your DoTs get nothing from it, so press it as you start a run of Shadow Bolts" | The raid rotation casts no runs of Shadow Bolts (Wrack is the filler). Wrack's damage is periodic, and the Eureka! tooltip (Wowhead Forever 1259821) limits it to "non-periodic damaging abilities", which the pages follow (round 9). In this rotation its only target is Immolate's first hit. | "Gnome is a step behind: Eureka! gives 8 seconds of 10% more damage and 10% lower cost on your direct damage spells every two minutes, but your DoTs and Wrack get nothing from it, so press it with an Immolate inside it; Expansive Mind's 5% more mana helps a spec that taps." |
| WARLOCK-R10-9 | CONFIRMED | rotation (card note) | warlock-affliction.json, Bane of Agony line (card condition "Bane of Agony whenever it has fallen off.") | "Bane of Agony whenever it has fallen off. Never clip it, since its biggest ticks come last, and Amplify Curse makes the next one 50% stronger." | At band 60 the card shows it right under Bane of Doom, while the line itself is gated on no Doom and Doom covers every minute but the last; affliction.md:67 says Agony only in the last minute. At bands 20 to 50 (no Doom until 60, `trainables/warlock.json` 603 level 60) "whenever it has fallen off" is right, so the fix must hold at every band. | "Bane of Agony whenever no Bane of Doom is on the target: all fight before 60, and the last minute at 60, once Doom has fired. Never clip it, since its biggest ticks come last, and Amplify Curse makes it 50% stronger." |
| WARLOCK-R10-10 | CONFIRMED | rotation (card note) | warlock-affliction.json, Shadow Bolt line | "Shadow Bolt with every DoT up. It is your filler." | Both builds take Wrack at 40 (leveling order index 30 = level 40; raid build Wrack 1); the Wrack line sits above it on the band 40, 50 and 60 cards, so from 40 Shadow Bolt is not your filler. | "Shadow Bolt with every DoT up, until Wrack at 40 takes over as your filler." |
| WARLOCK-R10-11 | CONFIRMED | rotation (card note) | warlock-demonology.json and warlock-destruction.json, Bane of Agony line (card condition "Bane of Agony in the last minute, once Doom has fired.") | "Bane of Agony in the last minute, once Doom has fired. A new Doom would not land before the boss dies." | Rendered on the band 20, 30, 40 and 50 cards of both specs (`addon-data.json` rotations), where Bane of Doom is not trained (603, level 60). There Agony is the bane all fight, as demonology.md:128 and destruction.md:128 say. Present since before round 9; missed, not re-filed. Card notes are written against level 60, so this may route to the component lane, but the note itself can hold at every band. | "Bane of Agony whenever Bane of Doom is not on the target: all fight before 60, and the last minute at 60, once Doom has fired. A new Doom would not land before the boss dies." |

## SHOULDs (wording, not findings)

- warlock-affliction.json: the new Bane of Doom, Amplify Curse and Wrack lines have no note, so their card clause is empty. Proposed (sim round 15 wording, checked): Doom "Bane of Doom while a minute or more of the fight is left. One global for a big hit after a minute, and Pandemic doubles the bonus on its crits." Amplify "Amplify Curse just before the Bane of Agony that covers the last minute. It makes that Agony 50% stronger and costs no global." Wrack "Wrack with your DoTs up. With Soul Siphon it out-damages Shadow Bolt for a fraction of the mana, and your other Shadow DoTs take 10% more while it runs."
- warlock-demonology.json Soul Fire note: "in the raid build it does no more than the Shadow Bolt it replaces" is still true, but the line no longer fires in the raid build at all. "In the raid build, skip it: Soul Fire is Fire, and your Shadow bonuses make Shadow Bolt stronger." reads truer.
- demonology.md:84 and :128 still call Decimation's Soul Fires "free" (no shard, 335 mana); carried.
- destruction.md:106 "its 2% crit lands for double through Ruin"; carried from round 9.

## Decisions

| # | decision | reason and sources | guide change |
|---|---|---|---|
| D1 | Affliction prose follows the shipped raid build and APL: Wrack is the raid filler. | Decoded raidBuild; `warlock-affliction.json`; sim round 15 A1 (+13.6% at both seeds, adoption rule held). The sim lane's guide-lane list was written against older line numbers; the fixes above apply it to the current text. | R10-1 to R10-7 |
| D2 | Eureka! stays on the tooltip reading (non-periodic only), so Wrack gets nothing from it on the page. | Round 9 coordinator rule; Wowhead Forever 1259821. The client mask's periodic 10% on Wrack stays a sim item (sim round 15 E1). | R10-8 |
| D3 | Card notes that render at several bands must hold at each of them. | `web/src/lib/bis/rotation-view.ts` shows each entry from its own level through the band top; Doom is level 60. | R10-9, R10-11 |
| D4 | Leveling talent orders stay as shipped. | Brief rule. | none |

## Sims requested

- **Affliction races on the new raid build** (`25520020135001351/00052/005004`, the round 15 rotation, band 60 raid preset, 10,000 iterations, seeds 1 and 50001): Alliance Human vs Gnome; Horde Orc vs Undead vs Troll. affliction.md:97 ("Gnome is a step behind", "Troll ... a step behind Orc", "Undead is as good as Orc") was measured on the Shadow Bolt build; Ruthless, Blood Fury and Berserking now act on a Wrack rotation.
- **PROVISIONAL (BiS re-rank):** `bis/warlock-affliction.json` (generated 2026-11-04T05:00Z) predates the round 15 regen and the new raid build. Check affliction.md:9 statPriority and the :85 reasons (haste on a channel filler, Shadow Bolt no longer cast) against the 13:00 UTC re-rank.
- Carried and still unrun: Destruction weapon enchant (sim round 15 E3 measured it: oil, then Firestone, then Spellstone; the guide already says so, so this one can close).

## Sim findings

- None new. Still open: the Eureka! class masks (Rain of Fire and Hellfire in, a periodic 10% on the drains and Wrack, Incinerate in neither) against the tooltip (sim round 15 E1); Wrack hover text at 36 a second against the talent's 86 (data lane); the Demonology Spirit weight check. Closed by sim round 15: Curse of the Elements order (C1), the Affliction Agony card and prepull Amplify (A2), the Wrack build on 70755 (A1), Decimation's 6% (M1), the Demonology raid Soul Fire (G3).

## Sources

- Sim round 15, warlock: `design/reviews/2026-11-04-sim-audit-round15/warlock.md` (A1, A2, C1, E1, E3, G3, M1).
- Wowhead Forever tooltip, Eureka! (warlock) 1259821, https://nether.wowhead.com/forever/tooltip/spell/1259821 (round 8 read): non-periodic damaging abilities.
- Blizzard, World of Warcraft: Forever Deep Dive Panel Recap, https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap (200, 2026-11-04).
- wago.tools, SpellEffect at build 1.60.1.70755, https://wago.tools/db2/SpellEffect?build=1.60.1.70755 (200, 2026-11-04).
- But Why Tho, "BLIZZCON 2026: The 'World Of Warcraft: Forever' Developers Talk Balancing Old And New", https://butwhytho.net/2026/09/wow-classic-devs-world-of-warcraft-forever/ (200, 2026-11-04).
- Forever Sixty, stat weights, https://foreversixty.gg/sim/weights (200, 2026-11-04).
- Client data, 1.60.1.70755: `talents/warlock.json` (Improved Bane of Agony, Aftermath, Bane, Soul Siphon, Pandemic, Malevolence, Wrack), `spellconst/warlock.json` (1316697, 11661, 603, 18288), `trainables/warlock.json` (603 level 60), `addon-builds.json`, `addon-data.json` rotations; `data/curated/apl/warlock-*.json`; `web/src/lib/bis/rotation-view.ts`.

## Count

CONFIRMED 10 (R10-1 to R10-6, R10-8 to R10-11), LIKELY 1 (R10-7), NEEDS-SIM 0, NEEDS-OWNER 0. SHOULDs 4. Decisions 4. Verdict: not zero. Sim round 15 moved the Affliction raid build and rotation to Wrack, and only the build line reached the page: its talent paragraphs, rotation, filler, Gnome line and two card notes still describe the old Shadow Bolt build.

## Fix log (round 10)

All edits verified against `talents/warlock.json` (Wrack 86 a second, other Shadow DoTs +10% for 6 s; Soul Siphon 12% per other Affliction effect up to 36%; Improved Bane of Agony 2 = 10%; Bane 5 = 0.5 s off Immolate; Aftermath 4 = Immolate initial +40%; Pandemic covers Bane of Doom and Wrack; Amplify Curse 50% on Agony), engine `sim/warlock/wrack.go` at 9f7f1f6e4 (channel, 6 s) and the APL actions (unchanged).

| item | result |
|---|---|
| WARLOCK-R10-1 | fixed, affliction.md:48 (auditor wording) |
| WARLOCK-R10-2 | fixed, affliction.md:57 |
| WARLOCK-R10-3 | fixed, affliction.md:59 (raid trade rewritten for the shipped Wrack build: drops Nightfall, Soul Harvest, Improved Drains, Improved Health Funnel, Improved Voidwalker; takes Malevolence, Improved Bane of Agony 2, Bane, Aftermath 4; Improved Shadow Bolt and Ruin removed) |
| WARLOCK-R10-4 | fixed, affliction.md:71 (Wrack is the filler) |
| WARLOCK-R10-5 | fixed, affliction.md:65 ("and Wrack from then on") |
| WARLOCK-R10-6 | fixed, affliction.md:27 and :43 |
| WARLOCK-R10-7 (LIKELY) | fixed, affliction.md:79; verified from Soul Siphon's per-effect bonus and the APL order (Wrack above Shadow Bolt, no gate) |
| WARLOCK-R10-8 | fixed, affliction.md:97 (Eureka! with an Immolate; DoTs and Wrack get nothing) |
| WARLOCK-R10-9 | fixed, warlock-affliction.json Bane of Agony note (holds at every band) |
| WARLOCK-R10-10 | fixed, warlock-affliction.json Shadow Bolt note |
| WARLOCK-R10-11 | fixed, warlock-demonology.json and warlock-destruction.json Bane of Agony notes (hold at every band) |
| SHOULD: empty notes | fixed: Bane of Doom, Amplify Curse and Wrack lines in warlock-affliction.json now carry notes (auditor wording; Amplify holds at every band since its gate is the last minute) |
| SHOULD: Demonology Soul Fire note | fixed, warlock-demonology.json Soul Fire note ("In the raid build, skip it ...") |
| SHOULD: "free" Soul Fires | fixed, demonology.md:84 ("costs no shard") and :128 ("shard-free") |
| SHOULD: destruction.md:106 Ruin wording | fixed ("Ruin doubles the bonus on the crits its 2% adds") |
| Extra | affliction.md:81 Adds and AoE no longer compares DoTs on adds against Shadow Bolt (the filler is Wrack now; the comparison is unmeasured), it just says to dot each add |
| Stat priority rule | all three spec pages checked; the prose already explains without ranking; prose test clean for warlock |
| D1 | applied (R10-1 to R10-7) |
| D2 | applied (R10-8: Wrack gets nothing from Eureka! on the page) |
| D3 | applied (R10-9, R10-11 notes hold from 20 to 60) |
| D4 | applied, leveling orders and prose untouched |
| updated: | affliction.md, demonology.md, destruction.md already 2026-11-04; index.md not changed |

Left: `warlock-affliction.json` top-level `notes`/`engineering_notes` still call Wrack the "#2 damage source behind Shadow Bolt" and say the Wrack line is gone; these are engineering notes, outside this lane's card-notes scope (route to the sim lane).

### SIM findings (round 10)

- Still open from the audit: Affliction races on the Wrack raid build (Human vs Gnome; Orc vs Undead vs Troll). affliction.md:97 keeps the round 9 measured order ("Gnome is a step behind", "Troll ... a step behind Orc", "Undead is as good as Orc") until that sim runs.
- Wrack's vulnerability: the client tooltip says "your other Shadow damage over time effects"; the engine's class mask (0x402) covers only Corruption and Bane of Agony, so Siphon Life gets nothing in the sim. The page follows the tooltip; confirm the mask on the client.
- Eureka! periodic mask on Wrack (sim round 15 E1) unchanged.
