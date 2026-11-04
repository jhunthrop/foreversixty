# Guide accuracy round 8: Warlock

Auditor: warlock lane, round 8. Tree: gr8/audit at 09b98d9e4 (land/guides-r7, all seven rounds applied), client 1.60.1.70755, engine pin per `sim/enginever/version.go`.
Audited line by line: `web/src/content/guides/warlock/index.md`, `affliction.md`, `demonology.md`, `destruction.md` (frontmatter, every body line, Levels 1 to 19 and Levels 20 to 60), the rendered card notes in `data/curated/apl/warlock-*.json`, and every source link.
Read first: `round1/warlock.md` to `round7/warlock.md` (every fix log and decision), `round4/warlock-levels.md`, the race-sims files, `ledger.md`. Nothing below re-files a fixed, sim-round, accepted, component-lane or data-lane item. Round 7's edits (diff 394562f20..09b98d9e4: three Intellect sentences, the Demonology Shadow and Fire power sentences, the Demonology 20 to 60 demon sentence, three weapon-enchant sentences, the Destruction setup bullet, the index date) were re-checked against the current text and data.

## Builds and shipped orders

No warlock build string, `talents/warlock.json`, `addon-builds.json` or `bis/warlock-*.json` changed since round 6. All six frontmatter strings decode again (tier, column order) to 51 points with every gate and prerequisite held (Wrack needs Siphon Life, Soul Link needs Demonic Sacrifice, Demonic Pact needs Soul Link, Incinerate needs Bane of Havoc, Fire and Brimstone needs Conflagrate):

- Affliction leveling `25522300130201351/200522/005`; shipped order Improved Life Tap 10-11, Suppression 12-16, Improved Corruption 17-21, Grim Reach 22-23, Soul Harvest 24-25, Improved Drains 26-28, Amplify Curse 29, Siphon Life 30, Pandemic 31-33, Nightfall 34-35, Soul Siphon 36-38, Shadow Mastery 39, Wrack 40, Shadow Mastery 41-44, Improved Health Funnel 45-46, Unholy Power 47-51, Demonic Aegis 52-53, Improved Voidwalker 54-55, Bane 56-60.
- Affliction raid `2552001013500105/00052/0550004` (Affliction 30, Demonology 7, Destruction 14 with Ruin 4).
- Demonology leveling `25522/203511310122000135/005`; Improved Health Funnel 10-11, Demonic Embrace 12-14, Unholy Power 15-19, Demonic Aegis 20, Improved Voidwalker 21, Fel Vitality 22-24, Demonic Energies 25, Demonic Sacrifice 26, Master Summoner 27-28, Decimation 29 and 31, Soul Link 30, Demonic Knowledge 32-34, Master Demonologist 35-39, then Affliction 40-55, Bane 56-60.
- Demonology raid `05/2035113011220001351/0550005` (Demonology 31, Pact the 31st point; trades Improved Corruption, Improved Life Tap, Grim Reach, Soul Harvest, Demonic Energies for Pact, Improved Sayaad 1, Improved Shadow Bolt 5, Ruin 5: 12 for 12).
- Destruction leveling `25522/0/2053235103101351`; Destructive Reach 10-11, Bane 12-16, Molten Skin 17-19, Cataclysm 20-21, Aftermath 22-24, Ruin 25-29, Shadowburn 30, Agonizing Flames 31-33, Conflagrate 34, Bane of Havoc 35, Fire and Brimstone 36-38, Shadow and Flame 39 and 41-44, Incinerate 40, Improved Life Tap 45-46, Suppression 47-51, Improved Corruption 52-56, Grim Reach 57-58, Soul Harvest 59-60.
- Destruction raid `25500000033/0/0050355103101351` (Affliction 18, Destruction 33).

Every leveling sentence on the three pages matches these orders.

## Checked and correct (no finding)

- **Round 7 edits:** Intellect's crit (0.0165% a point at 60, about 1% per 61; "about 1% for every 60 points" holds), the Demonology Shadow/Fire power split (matches :84 and :94), the 20 to 60 demon sentence (Soul Link 30, Demonic Knowledge 32-34, Master Demonologist 35-39), the weapon-enchant sentences (Spellstone 17728 and Firestone 17953 descriptions; a shard each per the Wowhead Forever tooltips), and the Destruction "Little setup" bullet.
- **Talent text at the decoded ranks:** Molten Skin 6% at 3 (10% at 5), Demonic Embrace 9% at 3, Demonic Aegis 15% at 1, Improved Sayaad 10% at 1, Improved Voidwalker 10% at 1, Improved Bane of Agony 5% at 1, Malevolence 3% at 3 and 5% at 5, Ruin 80% at 4, Aftermath 30% at 3 and 50% at 5, Cataclysm 3, Shadow and Flame 20% / 20 s and its rank 1-4 partial chances (so "until 5 points Conflagrate can eat your Immolate" holds), Master Demonologist 2% a point, Soul Siphon 12% per other Affliction effect up to 36%, Decimation 90% cooldown passive / 6% / 40% / 10 s on Shadow Bolt or Searing Pain, Soul Link 30% / 3%, Demonic Knowledge 100% of level, Demonic Pact (only the sacrificed pet's resummon cancels), Bane timings, Improved Health Funnel 40% / 30% / no threat / any demon health, Master Summoner 4 s / 40%.
- **Spell data:** Amplify Curse has no global cooldown (client SpellCooldowns row 55294: StartRecoveryTime 0, 180 s); Soul Link triggers a global (row 55405) and costs 20% base mana; Shadowburn reagent Soul Shard, 15 s, 8 s shard rule (Wowhead Forever 18871); Shadow Trance lasts 10 s (Wowhead Forever 17941); Wrack 200 mana for a 6 s channel against Shadow Bolt's 380 per 3 s, so "costs far less mana" holds; Curse of the Elements r4 "Magic resistances by 75 ... Magic damage taken by 10%", 5 min; Curse of Recklessness r4 505; Bane of Doom 1742 after 1 min, 60 s; Bane of Agony 552 over 24 s; Corruption r7 2.0 s cast, 18 s; Immolate 2.0 s; Incinerate 2.5 s; Soul Fire 6 s / 60 s; Life Tap 420 plus Spirit; Drain Soul shard text; Death Coil heals 100%; Hellfire burns the caster; Rain of Fire Fire; summons 10 s.
- **Trainer levels:** Immolate 1, Corruption 4, Life Tap 6, Bane of Agony 8, Voidwalker and Drain Soul 10, Health Funnel 12, Drain Life 14, Succubus and Incubus 20, Felhunter and Siphon Life 30, Wrack and Incinerate 40, Death Coil 42, Soul Fire 48, Bane of Doom and Portal of Summoning 60; Subjugate Demon present, no Curse of Shadow; Unstable Affliction and Haunt are class_spell rows only.
- **Items:** Lesser Magic Wand req 5, Greater Magic Wand req 13; Elixir of Fire Power 6373 is the Fire elixir there is (21546 is Elixir of Holy Power in Forever; `presets.json` already says so).
- **Races:** `combos.json` warlock rows are Human, Orc, Undead, Gnome and Troll (new); no Skyborne, so no paid race to name. Racial text matches `races.json`. Picks match the race sims.
- **Card notes:** nothing false beyond the accepted Agony (Affliction) and Soul Fire (Demonology) actions.
- **Links** (fetched 2026-11-04): Blizzard Deep Dive recap 200, "Troll Warlock" on the page; But Why Tho 200; `https://foreversixty.gg/sim/weights` 200; `https://wago.tools/db2/SpellEffect?build=1.60.1.70755` 200 with currentVersion 1.60.1.70755.
- **Voice:** no em or en dashes, no site, sim, search, APL, preset or band talk, no quoted nightly numbers, no stat ranking in any Stat priority section.

## Findings

| id | status | kind | file:line | wrong claim (quoted) | evidence | exact fix |
|---|---|---|---|---|---|---|
| WARLOCK-R8-1 | CONFIRMED | consumable/race | destruction.md:110 | "Eureka! gives 8 seconds of 10% more damage and 10% lower cost on everything that is not a DoT tick, every two minutes" | Wowhead Forever tooltip, Eureka! (warlock) 1259821: "Your non-periodic damaging abilities have their Mana cost reduced by 10% and deal 10% more damage for 8 sec." Only damaging, non-periodic abilities qualify. Curse of the Elements, Corruption and Bane of Doom (periodic-only casts), Demon Armor, summons, stones and Rain of Fire or Hellfire ticks are all "not a DoT tick" and get nothing. The Affliction (:97) and Demonology (:110) lines say "your direct damage spells" correctly; round 7 fixed the same overstatement for rogue. Round 3 (R3-2) wrote this sentence. | "Gnome is a step behind. Eureka! gives 8 seconds of 10% more damage and 10% lower cost on your direct damage spells every two minutes, Immolate's first hit, Incinerate, Conflagrate and Shadowburn among them, but never on a DoT or channel tick: press it as Conflagrate and Shadowburn come off cooldown and fill the rest of it with Incinerate." |
| WARLOCK-R8-2 | LIKELY | rotation | affliction.md:37 | "Slow start: your first five globals go into putting DoTs up" | The page's own opener (:65) is Curse of the Elements only "if you hold it", then Corruption, Bane of Doom, Siphon Life and Immolate: four DoT globals, and the overview (:27) says "three DoTs and a bane". Curse of the Elements is a curse, not a DoT, and only one warlock in the raid holds it (:65, index.md:24). For every other Affliction warlock the count is four. | "Slow start: your first four globals, five when you hold Curse of the Elements, go into putting DoTs up" |

## SHOULDs (wording, not findings)

- demonology.md:84 and :128 call Decimation's Soul Fires "free". They cost no shard, but still 335 mana. "shard-free" reads truer; the card note already says "cost no shard".
- demonology.md:32 lists Soul Link among "Shadow bonuses"; it is 3% to all damage. Harmless, since it also covers Shadow.

## Decisions

| # | decision | reason and sources | guide change |
|---|---|---|---|
| D1 | Eureka! prose on all three pages follows the client description (non-periodic damaging abilities). | Wowhead Forever 1259821; round 3 D2 (covers Incinerate) stands; the engine mask stays a sim item. | WARLOCK-R8-1 |
| D2 | The opener count follows the page's own opener. | affliction.md:27, :65. | WARLOCK-R8-2 |
| D3 | Leveling talent orders stay as shipped. | Brief rule: prose follows `addon-builds.json`. | none |

## Sims requested

None new. Carried and still unrun: Destruction weapon enchant (Brilliant Wizard Oil vs Major Firestone vs Major Spellstone, band 60 raid preset, published gear and raid build, Human), and Destruction Gnome vs Human once the Eureka! mask covers Incinerate.

## Sim findings

None new. Still open, not re-filed: Curse of the Elements before Corruption in all three APLs; the Affliction card casting Agony and a prepull Amplify instead of Doom; the Demonology card's Soul Fire line; Wrack's aura mask (Siphon Life); the Wrack build on 70755; the Eureka! mask (Incinerate in, channel ticks out); Destruction band 60 raid BiS rows on the pre-round-2 raid build and the Gnome / Troll BiS captions (data lane); Wrack hover text at 36 a second against the talent's 86; the Demonology Spirit weight check (round 7).

## Sources

- Wowhead Forever tooltip, Eureka! (warlock) 1259821, https://nether.wowhead.com/forever/tooltip/spell/1259821 (read 2026-11-04): non-periodic damaging abilities, 10% less mana and 10% more damage for 8 sec, 2 min cooldown.
- Wowhead Forever tooltip, Shadow Trance 17941, https://nether.wowhead.com/forever/tooltip/spell/17941 (read 2026-11-04): next Shadow Bolt instant, 10 seconds.
- Wowhead Forever tooltip, Shadowburn 18871, https://nether.wowhead.com/forever/tooltip/spell/18871 (read 2026-11-04): 365 mana, 15 s cooldown, reagent Soul Shard, 8 s shard rule.
- Blizzard, World of Warcraft: Forever Deep Dive Panel Recap, https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap (read 2026-11-04).
- But Why Tho, "BLIZZCON 2026: The 'World Of Warcraft: Forever' Developers Talk Balancing Old And New", https://butwhytho.net/2026/09/wow-classic-devs-world-of-warcraft-forever/ (read 2026-11-04).
- Forever Sixty, stat weights, https://foreversixty.gg/sim/weights (read 2026-11-04).
- wago.tools, SpellEffect at build 1.60.1.70755, https://wago.tools/db2/SpellEffect?build=1.60.1.70755 (read 2026-11-04, 200, build on the page).
- Client data, 1.60.1.70755: `talents/warlock.json`, `trainables/warlock.json`, `spelltips/warlock.json` (18288, 1316697, 1311680, 11689, 11675, 19028, 18871, 603, 11713, 25309, 1293813, 18932, 25311, 18881, 17728, 17953, 11684, 11678, 712, 17924, 25307, 17923, 437169), `addon-builds.json`, `combos.json`, `races.json`, `items.json` (6373, 21546, 11287, 11288), `data/curated/presets.json`; client SpellCooldowns (rows 55294, 55405) in the repo root `SpellCooldowns_era.csv`.

## Count

CONFIRMED 1 (R8-1), LIKELY 1 (R8-2), NEEDS-SIM 0, NEEDS-OWNER 0. SHOULDs 2. Decisions 3. Verdict: not zero. The one confirmed error is a round 3 sentence: the Destruction Gnome line gives Eureka! to "everything that is not a DoT tick", while the client text limits it to non-periodic damaging abilities.

## Fix log (round 8)

- WARLOCK-R8-1 fixed, destruction.md:110: "your direct damage spells ..., Immolate's first hit, Conflagrate and Shadowburn among them, but your DoTs get nothing from it". The report's "never on a DoT or channel tick" was not used: the pinned engine's client-mask table (sim/core/eureka_auto_gen.go, warlock row 1259821) puts Rain of Fire and Hellfire in the damage modifier and Drain Life and Drain Soul in a 10% periodic modifier, so a channel claim is not settled by the 70755 data. Incinerate kept per round 3 D2.
- WARLOCK-R8-2 fixed, affliction.md:37: four globals, five with Curse of the Elements.
- D1 to D3 applied. SHOULDs (Decimation "free", Soul Link as a Shadow bonus) left: wording only.
- SIM finding: Eureka! client text says non-periodic only, while the warlock class masks cover Rain of Fire/Hellfire hits and add a periodic 10% to Drain Life, Drain Soul, Haunt and Wrack; Incinerate is in neither mask. Decide which the engine follows.
- Prose check: warlock clean.
