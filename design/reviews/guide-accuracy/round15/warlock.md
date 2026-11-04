# Guide accuracy round 15: Warlock

Auditor: warlock lane, round 15. Tree: gr15/audit at 4600b7d82 (fourteen guide rounds, sim round 17, launch state), client 1.60.1.70755, engine pin b2c4a2d05 (`sim/enginever/version.go`).
Audited line by line: `web/src/content/guides/warlock/index.md`, `affliction.md`, `demonology.md`, `destruction.md` (frontmatter, every body line, Levels 1 to 19 and Levels 20 to 60), the curated APLs `data/curated/apl/warlock-*.json`, the rendered card clauses in `addon-data.json` `rotations`, `addon-builds.json`, and every source link.
Read first: `round1/warlock.md` to `round14/warlock.md`, `round4/warlock-levels.md`, the race-sims files, `ledger.md`, and the sim-accuracy round 17 reports for warlock, hunter (H2) and integration, plus round 16 hunter H1.

## What changed since round 14 (d7fc51fe7..4600b7d82)

Sim round 17 made the only warlock changes. I checked each one as new work.
- **Affliction leveling build** (round 17 A1, commits 8a4114ad9 and e18d0bb92): `25522300130201351/200522/005` became `25522300132001351/200522/005`. Nightfall 2 became Malevolence 2. affliction.md:6, :48, :54, :59 and :115 were rewritten.
- **Heart of the Lion on demons** (round 17 C4, engine b2c4a2d05, commits 27d777138 and afc7c11e9): index.md:29 and demonology.md:64 now name Heart of the Lion as the one raid buff a demon takes.
- `addon-builds.json`: in the Affliction leveling order, levels 34 and 35 are now Malevolence. `addon-data.json`: the warlock card clause text has not changed (I diffed every warlock path). Only `default_builds` (the same two order slots) and the Affliction `dps_per_point` bands moved.
- APL: only the Affliction `notes` and `engineering_notes` text changed ("largest damage source", in place of "#2 behind Shadow Bolt"). The rotation is unchanged. This closes the SHOULD carried since round 10.
- No change to `talents/warlock.json`, trainables, spelltips, `races.json`, `stat-weights.json` or `bis/warlock-*.json`.

## Builds (tier, column order)

- **Affliction leveling (new)** `25522300132001351/200522/005`: ILT 2, Suppression 5, IC 5, Grim Reach 2, Soul Harvest 2, Improved Drains 3, Amplify Curse 1, Pandemic 3, Malevolence 2 (tier 4), Siphon Life 1, Soul Siphon 3, Shadow Mastery 5, Wrack 1 (35); IHF 2, Unholy Power 5, Demonic Aegis 2, Improved Voidwalker 2 (11); Bane 5. Tier 4 needs 15 points and has 23 below it, so the build is legal.
- Affliction raid `25520020135001351/00052/005004`, Demonology `25522/203511310122000135/005` and `05/2035103021220001351/0550005`, Destruction `25522/0/2053235103101351` and `25500000033/0/0050355103101351`: unchanged and still legal (decoded again).

Shipped orders, re-derived from `addon-builds.json` (index 0 = level 10):
- Affliction leveling: ILT 10-11, Suppression 12-16, IC 17-21, Grim Reach 22-23, Soul Harvest 24-25, Improved Drains 26-28, Amplify 29, Siphon Life 30, Pandemic 31-33, **Malevolence 34-35**, Soul Siphon 36-38, Shadow Mastery 39 and 41-44, Wrack 40, IHF 45-46, Unholy Power 47-51, Aegis 52-53, Improved Voidwalker 54-55, Bane 56-60.
- Demonology and Destruction: the same as round 14.

Every level that the Levels 1 to 19 and Levels 20 to 60 paragraphs name still matches these orders. Affliction no longer mentions Nightfall anywhere except the bullet that explains why it is out.

## The round 17 edits, checked as new work

**Affliction Malevolence (affliction.md:48, :54, :59, :115)**
- :54 "1% crit a point on your Shadow spells": `talents/warlock.json` Malevolence gives 1% to 5%, "critical effect chance of your Shadow spells". "Pandemic doubles the bonus on those": Pandemic 3 gives +100% critical strike damage bonus on Corruption, both banes, Drain Soul, Drain Life, Siphon Life and Wrack, which are this page's Shadow DoT ticks and Wrack.
- "two points at 34 and 35 instead of Nightfall": matches the order. Nightfall sits in the same tier (tier 4, column 1).
- "Between 34 and 39 ... about level with Nightfall; from 40 they pull ahead": round 17 measured +0.2% to +0.5% at 35, 37 and 39, and +1.5% at 40 and 45, at both seeds. The prose gives no number.
- "a Nightfall bolt is instant but still costs full mana, while Wrack does more for less": Shadow Trance 17941 is a single aura 108 SPELLMOD_CASTING_TIME -100, with no cost modifier. Shadow Bolt r10 costs 380 and Wrack 200 (`spelltips`). The claim holds.
- :59: the leveling build has Soul Harvest 2, Improved Drains 3, IHF 2, IVW 2 and Malevolence 2. The raid build has none of the first four, Malevolence 5, Improved Bane of Agony 2 and Aftermath 4. So "drops ... takes the last three points of Malevolence" is exact, and the 9 points out match the 9 points in.
- :115: the Nightfall sentence is gone, and nothing else in the paragraph leaned on it.

**Heart of the Lion (index.md:29, demonology.md:64)**
- The exception itself holds. Client 409583 has Attributes_3 0 (no "can only target players" bit), aura 137/99/124 on ImplicitTarget 31, and the tooltip (409580) says "increasing total stats by 10% and Attack Power by 40 for all nearby allies". "10% more stats and its attack power" matches. "a hunter in your group" is a narrower but true condition. See D2.
- The exclusivity ("no raid buffs except", "with one exception") does not hold. See WARLOCK-R15-1.

## Findings

### WARLOCK-R15-1 (CONFIRMED, kind: stale-70755 / new edit, round 17 C4)
- **Where:** index.md:29 and demonology.md:64.
- **Quotes:** index.md:29 "**Your demon scales from you** and takes no raid buffs except Heart of the Lion, the hunter aura, which reaches everyone near the hunter, your demon included." demonology.md:64 "raid buffs do not land on it, with one exception: Heart of the Lion".
- **Evidence:** Forever keeps a group buff off pets with SpellMisc Attributes_3 0x100 ("can only target players"). That bit is set on Prayer of Fortitude, Gift of the Wild, Arcane Brilliance, Prayer of Spirit, the Greater Blessings (25916, 25898), Battle Shout (25289), Devotion Aura (10293), the paladin resistance auras (19896, 19898, 19900), Strength of Earth (25362), Grace of Air (25360), Windfury Totem (10612), Leader of the Pack (24932) and Moonkin Aura (24907).
  - Several other party auras do not have the bit (wago SpellMisc and SpellEffect 70755, read from `scratchpad/r17/warlock/wago`):
    - Retribution Aura 10301 (area party aura 15, a damage shield).
    - Concentration Aura 19746.
    - Blood Pact 11767 (area party aura 29, Stamina).
    - Mana Spring 10494.
    - The Fire, Frost and Nature Resistance Totem auras 10535, 10477 and 10599.
  - The beta logs confirm that this reaches pets. Sim round 17 hunter H2 found Retribution Aura and Blood Pact landing on pets in the beta timelines.
  - So a paladin's Retribution Aura, another warlock's Blood Pact and a shaman's resistance totems do land on your demon.
- **Why it matters:** the damage claim is right, because every damage and stat group buff skips the demon. But "no raid buffs except" is false. A Demonology player reading it would not stand the Succubus in a shaman's Fire Resistance Totem at Ragnaros.
- **Fix:**
  - index.md:29: "**Your demon scales from you.** The raid's group buffs skip it: Prayer of Fortitude, Gift of the Wild, Arcane Brilliance, the Greater Blessings, Battle Shout, and the Strength of Earth, Grace of Air and Windfury totems. Heart of the Lion, the hunter aura, is the exception: it reaches everyone near the hunter, your demon included. The Incubus is a summon of its own beside the Succubus, and Enslave Demon is now Subjugate Demon."
  - demonology.md:64 (first sentence): "Your demon scales from you in Forever, and the raid's damage and stat buffs skip it, with one exception: Heart of the Lion, the hunter aura, reaches every ally near the hunter, so when a hunter in your group runs it your demon gets its 10% more stats and its attack power too. Party auras that only protect still reach it, so park your Succubus in a shaman's Fire Resistance Totem on Fire fights."
- **Engine note:** the engine refuses every buff except Heart of the Lion, so it gives the demon no Blood Pact Stamina or resistance totem. None of these move damage, and the demon is never hit in the raid fight, so no sim finding.

No other statement on the four pages, and no rendered card clause, is false on 70755. None contradicts its own page, the header build card, the live BiS rows, the race table or the stat table.

## Checked and correct (no finding), fresh pass

- **Talent text at the ranks taken:** the same as round 14, plus Malevolence 2 = 2% and 5 = 5%. Destruction's Malevolence 3 = 3% ("Corruption ticks, Doom and Shadowburn" are its Shadow spells).
- **Spell facts and trainer levels** (`spelltips`, `trainables`): the same as round 14. Shadow Trance and Wrack were re-read above.
- **Rotation sections:** they match the APL priority lists (Affliction: CoE, Corruption, Doom while 60 s or more remain, then Amplify and Agony, Siphon Life, Immolate at 3 s, Wrack, Shadow Bolt before 40, wand last). The card clauses are unchanged.
- **Gear, professions, consumables, races and stats:** the BiS, weights and race data are unchanged since round 14. Every line still matches its table, as round 14 recorded.
- **Links** (curl, 2026-11-04, all 200): Blizzard Deep Dive recap, wago.tools SpellEffect 70755, But Why Tho, foreversixty.gg/sim/weights. The labels match their pages.
- **Voice:** no em or en dashes. No site, sim, APL, preset or band talk outside the source URLs. No nightly numbers and no stat ranking. The Malevolence bullet compares builds in words, with no figures.

## SHOULDs (wording, not findings)

- index.md:19 (carried from rounds 11 to 14): "Every spec plays out of the same core: Corruption, a bane and Immolate, with your talents deciding the filler: Wrack, Shadow Bolt or Incinerate."
- destruction.md:128 (carried): "Drain Life is your heal, and Death Coil joins it at 42."
- demonology.md:62 (carried, optional): one clause on why Improved Sayaad stops at 2 of 3.
- affliction.md:54: the bullet now carries a leveling comparison inside a key-talent list. The sentence "Between 34 and 39 ... more for less" could move to Levels 20 to 60, so the bullet reads as the talent.
- Closed: the APL `notes` "#2 damage source" SHOULD (fixed in round 17 A4).

## Decisions

| # | decision | reason and sources | guide change |
|---|---|---|---|
| D1 | The Affliction leveling Malevolence rewrite (affliction.md:6, :48, :54, :59, :115) stands. | Decoded build; `talents/warlock.json` Malevolence and Nightfall; `addon-builds.json` slots 24 and 25; Shadow Trance 17941 effects; sim round 17 A1 measurements (both seeds). | none |
| D2 | The demon does take Heart of the Lion. Blizzard's "Pets can no longer receive player buffs that increase their stats" is put into the client as the Attributes_3 0x100 bit on group buffs, and 409583 does not have it. Its tooltip says "for all nearby allies". The beta logs show unflagged Trueshot Aura, Retribution Aura and Blood Pact on pets, and the same spell buffed pets in Season of Discovery. The in-game pet buff bar check is still on the integrator's list (round 17 integration item 4). | wago SpellMisc/SpellEffect 70755; Blizzard, Hunter and Druid class deep dive; sim round 16 hunter H1, round 17 hunter H2 and warlock C4. | only the exclusivity wording (R15-1) |
| D3 | Leveling talent orders stay as shipped. Every prose level matches them. | `addon-builds.json`. | none |

## Sims requested

- **Affliction races on the Wrack raid build with Heart of the Lion on the demon** (engine b2c4a2d05): `25520020135001351/00052/005004`, band 60 raid preset, 10,000 iterations, seeds 1 and 50001. Compare Orc, Undead and Troll, then Human and Gnome. affliction.md:97 says "Undead is as good as Orc" on round 16's 0.4% gap. Heart of the Lion raises the demon's share, and Command rides on that share. If Orc now leads Undead beyond error, the line becomes "Undead is close behind Orc", as on the other two pages.
- **Demonology and Destruction races** (carried, with the sim lane): the same comparison on `05/2035103021220001351/0550005` and `25500000033/0/0050355103101351`, now on b2c4a2d05.

## Sim findings

- None new. Still open, all carried: Eureka! on Incinerate (client mask against tooltip, round 17 S1), demon crit at 60, Soul Siphon's count, Wrack and Doom at 60 (all waiting for a launch log), and Decimation's 0% client effect (known engine gap).

## Sources

- Client data, 1.60.1.70755: `talents/warlock.json`, `spelltips/warlock.json` and `spellconst/warlock.json` (17941, 18094, 1316697, 25307), `spelltips/hunter.json` (409580, 19506, 20905, 20906), `trainables/hunter.json`, `addon-builds.json`, `addon-data.json`; wago SpellMisc, SpellEffect and SpellName at 70755 (local copy `scratchpad/r17/warlock/wago/*.layered.csv`): 409583, 10301, 19746, 11767, 10494, 10535, 10477, 10599 and the flagged group buffs listed in R15-1.
- Sim-accuracy round 17 reports: `design/reviews/2026-11-04-sim-audit-round17/warlock.md` (A1, A4, C4), `hunter.md` (H2, the beta-log evidence for Retribution Aura and Blood Pact on pets), `integration.md`; round 16 `hunter.md` H1.
- Blizzard, World of Warcraft: Forever Class Deep Dives, Hunter and Druid, https://news.blizzard.com/en-us/article/24301515/world-of-warcraft-forever-class-deep-dives-hunter-and-druid (the quote as recorded in sim round 15 hunter W7-1).
- Blizzard, World of Warcraft: Forever Deep Dive Panel Recap, https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap (200, 2026-11-04).
- wago.tools, SpellEffect at build 1.60.1.70755, https://wago.tools/db2/SpellEffect?build=1.60.1.70755 (200, 2026-11-04).
- But Why Tho, "BLIZZCON 2026: The 'World Of Warcraft: Forever' Developers Talk Balancing Old And New", https://butwhytho.net/2026/09/wow-classic-devs-world-of-warcraft-forever/ (200, 2026-11-04).
- Forever Sixty, stat weights, https://foreversixty.gg/sim/weights (200, 2026-11-04).

## Count

CONFIRMED 1, LIKELY 0, NEEDS-SIM 0, NEEDS-OWNER 0. SHOULDs 4 (plus 1 closed). Decisions 3. Sims requested 2 (1 new, 1 carried).

## Fix log (round 15)

- WARLOCK-R15-1 fixed.
  - index.md:29 names the group buffs that skip the demon, keeps Heart of the Lion as the exception, and adds that protective party auras (Blood Pact, resistance totems) still reach it. updated set to 2026-11-04.
  - demonology.md:64 gets the same rule, with Heart of the Lion as the one damage exception and the Fire Resistance Totem advice.
