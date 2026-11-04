# Guide accuracy round 9: Warlock

Auditor: warlock lane, round 9. Tree: gr9/audit at 645caa983 (land/guides-r8, all eight rounds applied), client 1.60.1.70755, engine pin d1d9f2775 (`sim/enginever/version.go`).
Audited line by line: `web/src/content/guides/warlock/index.md`, `affliction.md`, `demonology.md`, `destruction.md` (frontmatter, every body line, Levels 1 to 19 and Levels 20 to 60), the rendered card notes in `data/curated/apl/warlock-*.json`, and every source link.
Read first: `round1/warlock.md` to `round8/warlock.md` (every fix log and decision), `round4/warlock-levels.md`, the race-sims files, `ledger.md`. Nothing below re-files a fixed, sim-round, accepted, component-lane or data-lane item. Round 8's two edits (diff 09b98d9e4..645caa983: the Destruction Gnome line at destruction.md:110 and the Affliction opener count at affliction.md:37) were re-checked against the current text and data.

## Builds and shipped orders

No warlock build string, `talents/warlock.json`, `addon-builds.json`, `bis/warlock-*.json` or warlock APL changed since round 6 (`git log`). All six frontmatter strings decode again (tier, column order) to 51 points with every gate and prerequisite held:

- Affliction leveling `25522300130201351/200522/005`: Affliction 35 (ILT 2, Suppression 5, Improved Corruption 5, Grim Reach 2, Soul Harvest 2, Improved Drains 3, Amplify Curse 1, Pandemic 3, Nightfall 2, Siphon Life 1, Soul Siphon 3, Shadow Mastery 5, Wrack 1), Demonology 11 (IHF 2, Unholy Power 5, Demonic Aegis 2, Improved Voidwalker 2), Destruction 5 (Bane 5).
- Affliction raid `2552001013500105/00052/0550004`: Affliction 30 (ILT 2, Suppression 5, IC 5, Grim Reach 2, Improved Bane of Agony 1, Amplify Curse 1, Pandemic 3, Malevolence 5, Siphon Life 1, Shadow Mastery 5), Demonology 7 (Unholy Power 5, Demonic Aegis 2), Destruction 14 (ISB 5, Bane 5, Ruin 4). The page's trade list (15 out, 15 in, Bane 5 kept) matches.
- Demonology leveling `25522/203511310122000135/005` and raid `05/2035113011220001351/0550005`: the raid trade (ILT 2, IC 5, Grim Reach 2, Soul Harvest 2, Demonic Energies 1 out; Pact 1, Improved Sayaad 1, ISB 5, Ruin 5 in) matches demonology.md:60; Demonic Embrace 3 (9%), Demonic Aegis 1 (15%), Improved Voidwalker 1 (10%), Improved Sayaad 1 (10%) match :62.
- Destruction leveling `25522/0/2053235103101351` and raid `25500000033/0/0050355103101351`: Destructive Reach 2 and Molten Skin 3 out, Aftermath +2 and Cataclysm +1 in; Grim Reach 2 and Soul Harvest 2 out, Pandemic 3 and Malevolence 3 in (tier gates: 12 points before Pandemic, 15 before Malevolence). Matches destruction.md:50.

Every leveling sentence on the three pages still matches the shipped orders round 8 lists (index 0 = level 10).

## Checked and correct (no finding)

- **Round 9 crit-damage lesson.** Every crit-damage claim was re-derived from the 70755 talent text as a multiplier on the crit bonus: Ruin "critical strike damage bonus ... by 100%" at 5 (80% at 4), Pandemic "by 100%" at 3. "Doubles the crit bonus" (index.md:25, affliction.md:32/53/85, demonology.md:60/110, destruction.md:34/53/61/98/110) and "raise the crit bonus ... by 80%" (affliction.md:59) are exact statements of a +100% / +80% multiplier on the bonus, true whatever the base spell crit bonus is. No page states a crit's total multiplier.
- **Eureka! lines** against "non-periodic damaging abilities" (Wowhead Forever 1259821): affliction.md:97, demonology.md:110 and destruction.md:110 all limit it to direct damage spells, cost and damage both, and say DoTs get nothing. Round 8's Destruction fix holds; Incinerate is a non-periodic damaging ability, so "fill the rest of it with Incinerate" holds.
- **Round 8 opener count** (affliction.md:37): four DoT globals (Corruption, Doom, Siphon Life, Immolate), five with Curse of the Elements, matches :65.
- **Spell data** (`spelltips/warlock.json`): Curse of the Elements r4 misc value 126 (all magic schools, Holy included, per the engine comment on spellconst), 75 and 10% for 5 min, so index.md:24 holds; Spellstone 17728 and Firestone 17953 are weapon enchants (Shadow haste 2% and up to 21; Fire crit 2% and up to 21); Siphon Life, both banes and Corruption (with Improved Corruption) are instant; Soul Fire 6 s cast, 60 s cooldown (Bane 5 makes it 4 s, Decimation 40% makes it 2.4 s, 6 s cooldown); Immolate 2.0 s (1.5 s with Bane); Incinerate 2.5 s; Corruption r1 2 s cast; Drain Soul "non-trivial targets damaged ... have a chance"; Shadowburn 15 s and the 8 s shard rule; Life Tap base plus Spirit; Death Coil heals 100%; Hellfire burns the caster; summons 10 s; Portal of Summoning one hour, caster plus two, two members summon.
- **Talent text at decoded ranks**: Soul Siphon 12% per other Affliction effect up to 36% at 3; Malevolence 3% at 3, 5% at 5; Improved Shadow Bolt 20% for 12 s at 5; Decimation rank 2 text (90%, 6%, 40%, 10 s); Amplify Curse; Demonic Pact; Master Demonologist; Soul Link; Shadow and Flame; Aftermath; Bane; Destructive Reach; Molten Skin 6% at 3.
- **Races** (`races.json` Forever text): Ruthless, The Human Spirit, Eureka!, Expansive Mind, Blood Fury (10% spell power, 15 s, 2 min), Command (5% warlock pets), Berserking (10%, 15 s, 3 min), Touch of the Grave (up to 5% of maximum health, once every 15 s). Picks match the race sims; no paid race is named (no Skyborne warlock in `combos.json`).
- **Card notes**: nothing false beyond the accepted Agony (Affliction) and Soul Fire (Demonology) actions.
- **Links** (fetched 2026-11-04): Blizzard Deep Dive recap 200 ("Troll Warlock" on the page); `https://wago.tools/db2/SpellEffect?build=1.60.1.70755` 200 with 1.60.1.70755 on the page; But Why Tho 200 ("Ana Resendez"); `https://foreversixty.gg/sim/weights` 200 ("Stat weights"). All match their labels and the source-link rule.
- **Voice**: no em or en dashes, no site, sim, search, APL, preset or band talk, no quoted nightly numbers, no stat ranking in any Stat priority section.

## Findings

None. No statement on the four pages is false on 70755 or contradicts its own page.

## SHOULDs (wording, not findings)

- destruction.md:106 "its 2% crit lands for double through Ruin": it is the crit's bonus that doubles, not the 2%. "and the crits it adds land for double the bonus through Ruin" reads truer.
- demonology.md:84 and :128 still call Decimation's Soul Fires "free" (no shard, but 335 mana); carried from round 8.

## Decisions

| # | decision | reason and sources | guide change |
|---|---|---|---|
| D1 | Keep "Shadow Bolt below 35% hits 6% harder" (demonology.md:33, :84). | The 70755 talent text (`talents/warlock.json`, Blizzard's per-rank text from the Wowhead Forever talent payload, still current on 2026-11-04) gives 3% / 6%. Against it: the client spell row 440870 carries that SPELLMOD_DAMAGE effect (index 3, aura 108) at amount 0 in every build from 70338 to 70755, and the live Wowhead Forever spell tooltip renders it as "0% increased damage". The two client sources disagree, the page follows the talent text, so it is not a hard finding. Filed as a sim finding below. | none |
| D2 | Crit-damage wording stays as "doubles the bonus". | Ruin and Pandemic raise the bonus by 100%; the pages never claim a total multiplier. | none |
| D3 | Leveling talent orders stay as shipped. | Brief rule: prose follows `addon-builds.json`. | none |

## Sims requested

None new. Carried and still unrun: Destruction weapon enchant (Brilliant Wizard Oil vs Major Firestone vs Major Spellstone, band 60 raid preset, published gear and raid build, Human), and Destruction Gnome vs Human once the Eureka! mask covers Incinerate.

## Sim findings

- **New:** the pinned engine's `applyDecimation` (sim/warlock/soul_fire.go:125-169 at d1d9f2775) models the cooldown cut and the 10 s Soul Fire window but no damage bonus on Shadow Bolt or Searing Pain below 35%. The talent text says 6% at rank 2; the client effect amount is 0 and the live Wowhead spell tooltip reads 0%. Decide which the server applies (a level 60 log of Shadow Bolt below and above 35% settles it). If it is 0%, demonology.md:33 and :84 must drop "6% harder". If it is 6%, the engine should add it.
- Still open, not re-filed: Curse of the Elements before Corruption in all three APLs; the Affliction card casting Agony and a prepull Amplify instead of Doom; the Demonology card's Soul Fire line; Wrack's aura mask; the Wrack build on 70755; the Eureka! class masks (Rain of Fire/Hellfire in, a periodic 10% on the drains and Wrack, Incinerate in neither); Destruction band 60 raid BiS rows and Gnome / Troll BiS captions (data lane); Wrack hover text at 36 a second against the talent's 86; the Demonology Spirit weight check.

## Sources

- Wowhead Forever tooltip, Decimation 440870, https://nether.wowhead.com/forever/tooltip/spell/440870 (read 2026-11-04): "they deal 0% increased damage", 90% cooldown, 40% cast time, 10 sec.
- Wowhead Forever talent payload, https://nether.wowhead.com/forever/data/talents-classic?dv=100 (read 2026-11-04): Decimation rank text 3% / 6%.
- Wowhead Forever tooltip, Wrack 1316697, https://nether.wowhead.com/forever/tooltip/spell/1316697 (read 2026-11-04): channeled 6 sec, 86 a second.
- Wowhead Forever tooltip, Eureka! (warlock) 1259821, https://nether.wowhead.com/forever/tooltip/spell/1259821 (round 8 read): non-periodic damaging abilities.
- Blizzard, World of Warcraft: Forever Deep Dive Panel Recap, https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap (read 2026-11-04).
- But Why Tho, "BLIZZCON 2026: The 'World Of Warcraft: Forever' Developers Talk Balancing Old And New", https://butwhytho.net/2026/09/wow-classic-devs-world-of-warcraft-forever/ (read 2026-11-04).
- Forever Sixty, stat weights, https://foreversixty.gg/sim/weights (read 2026-11-04).
- wago.tools, SpellEffect at build 1.60.1.70755, https://wago.tools/db2/SpellEffect?build=1.60.1.70755 (read 2026-11-04, 200, build on the page).
- Client data, 1.60.1.70755: `talents/warlock.json`, `spelltips/warlock.json`, `spellconst/warlock.json` (440870, 440873; same 440870 effects in 70338 to 70755), `races.json`, `combos.json`, `addon-builds.json`; engine d1d9f2775 `sim/warlock/soul_fire.go`, `sim/warlock/curses.go:236-241`.

## Count

CONFIRMED 0, LIKELY 0, NEEDS-SIM 0, NEEDS-OWNER 0. SHOULDs 2. Decisions 3. New sim finding 1 (Decimation's damage bonus). Verdict: zero confirmed errors for warlock.
