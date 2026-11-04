# Guide accuracy round 1: Priest

Audited on origin/main 64168181f, client 1.60.1.70755, engine pin 2eaae03c6.
Files: `web/src/content/guides/priest/index.md`, `discipline.md`, `holy.md`, `shadow.md`.

## Decoded builds (tab order = array order of `data/builds/1.60.1.70755/talents/priest.json`, all 51 points, every tier gate and prerequisite met on 70755)

**Discipline `build:`** `00500503133001143/3350500303/0` (29/22/0)
- Discipline: Twin Disciplines 5, Mental Strength 5, Mental Agility 3, Inner Focus 1, Improved Power Word: Shield 3, Meditation 3, Soul Warding 1, Penance 1, Renewed Hope 4, Divine Aegis 3
- Holy: Twilight Focus 3, Improved Renew 3, Holy Specialization 5, Divine Fury 5, Inspiration 3, Improved Healing 3
- Power Infusion swap in the prose (Improved Renew -2, Renewed Hope +1, Power Infusion +1) is legal: Disc reaches 30 before tier 6, Penance is PI's prerequisite, Holy keeps 14 before Improved Healing.

**Holy `build:`** `00500503103/33505003030121431/0` (17/34/0)
- Discipline: Twin Disciplines 5, Mental Strength 5, Mental Agility 3, Inner Focus 1, Meditation 3
- Holy: Twilight Focus 3, Improved Renew 3, Holy Specialization 5, Divine Fury 5, Inspiration 3, Improved Healing 3, Holy Binding 1, Litany of Light 2, Spirit of Redemption 1, Spiritual Guidance 4, Spiritual Healing 3, Prayer of Mending 1

**Shadow `build:` (leveling)** `02510503103/0/54342001201300051` (20/0/31)
- Discipline: Wand Specialization 2, Twin Disciplines 5, Silent Resolve 1, Mental Strength 5, Mental Agility 3, Inner Focus 1, Meditation 3
- Shadow: Shadow Focus 5, Blackout 4, Spirit Tap 3, Improved Mind Blast 4, Improved Shadow Word: Pain 2, Mind Flay 1, Improved Mind Flay 2, Vampiric Embrace 1, Shadow Weaving 3, Darkness 5, Shadowform 1
- Addon/strip order (`addon-builds.json` priest-shadow.leveling): Shadow Focus 10-14, Blackout 15-18, Spirit Tap 19-21, Improved Mind Blast 22-25, Improved SW:P 26-27, **Mind Flay 28**, Improved Mind Flay 29-30, VE 31, Shadow Weaving 32-34, Darkness 35-39, Shadowform 40, Wand Specialization 41-42, then the rest of Discipline.

**Shadow `raidBuild:`** `00510503103/0/52352101201300251` (18/0/33)
- Discipline: Twin Disciplines 5, Silent Resolve 1, Mental Strength 5, Mental Agility 3, Inner Focus 1, Meditation 3
- Shadow: Shadow Focus 5, Blackout 2, Spirit Tap 3, Improved Mind Blast 5, Improved Shadow Word: Pain 2, Shadow Reach 1, Mind Flay 1, Improved Mind Flay 2, Vampiric Embrace 1, Shadow Weaving 3, Early Demise 2, Darkness 5, Shadowform 1

## Verified clean (no finding)

- 70755: no Binding Heal anywhere on main (its trainer rows are gone: `diff 70753 70755 trainables/priest.json` removes only Binding Heal). Voice of Truth is a paladin trainable, not priest; nothing to check here.
- Baseline trainer spells: Fear Ward @20, Devouring Plague @20 (24 s cd), Shadow Word: Death @32 (18 s cd), Divine Spirit @30, Lightwell @40 (3 min or 5 charges, being attacked cancels) all in `trainables/priest.json`. Penance @30, Prayer of Mending @40, Mind Flay @20 are talent-sourced with trainer ranks.
- Shadowfiend "not learnable": `data/curated/class-spell-limits.json` notLearnable (row 401977 has no learn level). See O3.
- Dark Sacrifice (undead, `class-spell-limits.json` raceLocked): r5 "Cannibalize 1600 of your own Health over 15 sec to gain 1600+Spirit Mana", 10 min cd. "Well over a thousand mana once a fight" holds.
- Race spells (raceLocked from SkillLineAbility RaceMasks): Human Feedback/Divine Grace, Dwarf Desperate Prayer/Chastise, NE Elune's Grace/Starshards, Undead Touch of Weakness/Dark Sacrifice, Gnome Contingency Plan/Confounding Flash, Troll Hex of Weakness/Shadowguard. Index matches. Skyborne cannot be priests (`combos.json`).
- Racials (`races.json`): The Human Spirit 10% Spirit, Ruthless 1% crit, Artisan's Precision 1% hit, Elune's Light 10% crit 15 s / 3 min, Berserking 10% 15 s / 3 min, Expansive Mind 5% mana, Engineering Specialization. All match.
- Eureka! priest row 1259823 (`sim/core/eureka_auto_gen.go`): covers Mind Blast, Mind Flay, SW:D, Penance, Smite, Holy Fire, Heal/Greater/Flash/Lesser Heal, Prayer of Healing; not SW:P, Devouring Plague, Renew, Power Word: Shield. Shadow's "it skips Shadow Word: Pain and Devouring Plague" is right.
- Prayer of Healing "heals the target and their party" (spelltips 25316). Heal r4 and Greater Heal both 3.0 s with sp coefficient 0.857 (`spellconst/priest.json` 6064, 10965), so the Heal r4 downrank claim is right.
- Shadow Word: Death backlash 30% max health, Vampiric Embrace 20% / 30 s / 30 s cd, Shadowform 40% base mana cost (cost_pct 40) so "the form costs mana" is right; Shadow Weaving "from you"; Mind Flay channel crits (flagged in `periodic_crit_auto_gen.go`, 18807); Twin Disciplines mask includes Penance and Renew, Mental Agility mask includes Renew and PW:S (`sim/priest` masks).
- Healing items carry a third of their healing as spell damage (Hide of the Wild: healing 42, spell_power 14).
- All named consumables exist in `items.json` (Flask of Distilled Wisdom, Greater Mageblood Elixir, Greater Cleric's Elixir, Elixir of Sages, Sage's Tea, Brilliant Mana/Wizard Oil, Flask of Supreme Power, Elixir of Shadow Power, Greater Arcane Elixir, Demonic Rune, Major Mana Potion, Mender's/Spellblasting families).
- Stat priority prose (all three) matches the published 60 raid scale factors (`bis/priest-*.json`): Disc mp5 > Spirit > healing > Int > crit; Holy mp5 > healing > Spirit > Int > crit; Shadow spell power > hit > haste > crit > Int, spell pen 0. No quoted numbers.
- Prose test rules pass (priest is in ENFORCED_CLASSES); no dashes found.

## Findings

| id | status | kind | file:line | wrong claim | evidence | exact fix | also on r14 |
|---|---|---|---|---|---|---|---|
| PRIEST-1 | CONFIRMED | leveling / talent | shadow.md:121, shadow.md:34 | "Spirit Tap after every kill lets you chain pulls without drinking." / "Spirit Tap mean you rarely drink" | Leveling build has Spirit Tap 3/5; rank 3 text: "Gives you a 60% chance to increase your Spirit by 100%" (`talents/priest.json`, Spirit Tap ranks[2]) | Make the leveling build Spirit Tap 5/5 and Blackout 2/5 (Shadow tree `54342001201300051` becomes `52542001201300051`; see PRIEST-2), then the sentence is true. If the build stays, write: "Spirit Tap procs on most kills and lets you chain pulls with little drinking." | not fixed |
| PRIEST-2 | CONFIRMED | leveling / talent | shadow.md:41, shadow.md:123; data `addon-builds.json` priest-shadow.leveling | "No filler until Mind Flay at level 20" and "so Mind Flay is your point at 20" | The leveling order players import and see in the strips takes Mind Flay as the 19th point, level 28 (order entry [3,3,3] at index 18), after Shadow Focus 5, Blackout 4, Spirit Tap 3, Improved Mind Blast 4, Improved SW:P 2. Wand Specialization lands at 41 although the prose says it is there because "you wand a lot on the way up". The guide then tells the reader to ignore its own strips. | Reorder the leveling build (data, not prose): Spirit Tap 10-14, Improved Shadow Word: Pain 15-16, Improved Mind Blast 17-19, Mind Flay 20, Improved Mind Flay 21-22, Blackout 23-24, Vampiric Embrace 25, Shadow Weaving 26-28, Improved Mind Blast 29, Shadow Focus 30-34, Darkness 35-39, Shadowform 40, Wand Specialization 41-42, then Discipline. Then replace line 123 with: "Take Spirit Tap first at 10, then Improved Shadow Word: Pain and Improved Mind Blast, so Mind Flay is your point at 20. Improved Mind Flay, Vampiric Embrace and Shadow Weaving come next, Shadow Focus once you start running dungeons, and every point in Shadow gets you to Shadowform at 40." | not fixed |
| PRIEST-3 | CONFIRMED | leveling / voice | shadow.md:123 | "Each strip below shows the build for that level range ... Those builds buy Shadow Focus and Blackout early, which is right for dungeon bosses." | Blackout's stun does nothing on a boss; the same guide says so at shadow.md:67 ("bosses cannot be stunned"). "Strip" is page furniture, not something a player says. | Delete the two sentences once PRIEST-2 lands (the replacement in PRIEST-2 covers the order). | not fixed |
| PRIEST-4 | CONFIRMED | consumable | shadow.md:105 | "Flask of Supreme Power, or Elixir of Shadow Power with Greater Arcane Elixir when you do not have a flask." | The flask and both elixirs stack (separate consume slots in `sim/core/consumes.go`; the raid kit `sim/request/raid-consumes.json` priest-shadow from 55 carries flask_of_supreme_power + elixir_of_shadow_power + greater_arcane_elixir together, plus major_spellblasting_potion and cerebral_cortex_compound). In 1.12 they stack too. | "Flask of Supreme Power, Elixir of Shadow Power and Greater Arcane Elixir all stack, so run all three. Brilliant Wizard Oil on the weapon. Open with a Major Spellblasting Potion; a Major Mana Potion and a Demonic Rune come out on a long fight." | not fixed |
| PRIEST-5 | CONFIRMED | consumable | shadow.md:105, discipline.md:97, holy.md:96 | "Forever's Spellblasting Potion" / "Forever's Mender's Potion(s)" | Spellblasting Potion (250934) and Mender's Potion (250946) are the required-level-25 ranks; the level 60 items are Major Spellblasting Potion (250937, req 55) and Major Mender's Potion (req 55) in `simconsumes.json`. | Name the Major ranks: "Forever's Major Spellblasting Potion", "Forever's Major Mender's Potion". | not fixed |
| PRIEST-6 | NEEDS-OWNER | spell / stale | index.md:29, discipline.md:75 | "a heal learned below level 20 keeps its full healing power bonus, so low ranks of Renew and Lesser Heal are real mana savers at 60" | Unverified in game: memory forever-refresh-hold lists "Pending owner: tooltips Rejuv 1/Renew 1/Heal 1/Lesser Heal 1/Healing Wave 2 on a geared 60 (rows revert together if penalised)". The engine states no penalty by construction (APL notes). | Hold until the owner reads a geared 60's Lesser Heal r3 and Renew r1 tooltips. If penalised, replace with: "Below level 20 a heal's rank still loses part of its healing bonus, so the cheap ranks that pay off are the ones learned at 20 or later: Renew rank 3, Heal rank 4." | r14 goes further (Renew rank 1, Heal rank 1 in Disc/Holy) on the same unverified premise |
| PRIEST-7 | CONFIRMED | stale-70755 | r14 only: index.md:26, discipline.md (strengths, emergencies, description), all three spec `build:` lines | r14 adds "Binding Heal" to the baseline list and to Discipline's rotation, and its build strings read `FS1:1.60.1.70753:` | 70755 removed Binding Heal's trainer rows (`diff 1.60.1.70753 1.60.1.70755 trainables/priest.json`); main already reads 70755 in every build string. | Integrator: when merging int/r14, drop every Binding Heal line and keep main's `1.60.1.70755` build prefixes. Emergencies stays: "Flash Heal at the top rank on anyone about to die. It is your fastest heal and your dearest, so it never becomes the filler." | REGRESSION on r14 |
| PRIEST-8 | LIKELY | rotation | shadow.md:73 | Opener "Shadow Word: Pain, Devouring Plague, then Mind Blast ..." leaves out Vampiric Embrace, which shadow.md:81 says to keep up on the boss. | APL `priest-shadow.json` casts SW:P, then VE, then DP, then Inner Focus + Mind Blast. | "Shadow Word: Pain, Vampiric Embrace, Devouring Plague, then Mind Blast with Inner Focus, Shadow Word: Death and Mind Flay." | fixed |
| PRIEST-9 | LIKELY | rotation | shadow.md:83 | "Drop Shadow Word: Pain in the last 10 seconds of a fight; it will not tick out." | SW:P r8 ticks 127 + 0.2 SP every 3 s (`spellconst` 10894); a Mind Flay global is 1.5 ticks of 130 + 0.167 SP (18807). With 9 s left a refresh lands 3 ticks, about twice a Mind Flay global; it stops paying only near one tick left. Same 10 s cutoff in the APL. | "Stop refreshing Shadow Word: Pain once fewer than two ticks (6 seconds) of the fight are left." Pending S2. | not fixed |
| PRIEST-10 | LIKELY | talent | shadow.md:67, raidBuild | Raid build keeps "two points of Blackout" and Spirit Tap 3 | Blackout cannot stun a boss; Spirit Tap needs a kill or a VE target dying. Neither gates anything (Shadow Focus alone opens tier 1; Mind Flay's tier has 18 points before it). Five dead points in the raid build. The guide itself says threat is a problem (shadow.md:89). | Owner call O2. The five points must stay in Shadow (Shadowform needs 30 before it). Legal re-spend checked against the tier gates: Blackout 2 + Spirit Tap 3 into Devouring Contagion 2 (half-cost Devouring Plague, spreads on add deaths), Silence 1 (an interrupt for casters and adds), Shadow Reach +1 (2/2) and Dark Inspiration 1; tiers 0-3 keep 22, tiers 0-4 keep 25, Shadowform still has 32 before it. Prose then: "The raid build trades Wand Specialization, Blackout and Spirit Tap for Early Demise, Devouring Contagion, Silence, full Shadow Reach and the last point of Improved Mind Blast." | not fixed |
| PRIEST-11 | LIKELY | stats | discipline.md:8, holy.md:8, shadow.md:9 | `statPriority` frontmatter orders: Disc/Holy [Healing power, Intellect, Spirit, MP5, Crit]; Shadow [Intellect, Hit, Spell power, ...] | Published 60 raid scale factors: Disc mp5 2.52, Spirit 1.30, healing 1.17, Int 1.00, crit 0.67; Holy mp5 2.73, healing 1.69, Spirit 1.45, Int 1, crit 0.56; Shadow spell power 5.75, hit 4.61, haste 4.44, crit 3.83, Int 1. The pills sort by the data when it loads (`[spec].astro:211`), so the frontmatter order shows only as the fallback and in `guide-file.ts:339`. | Disc: `[MP5, Spirit, Healing power, Intellect, Critical strike]`; Holy: `[MP5, Healing power, Spirit, Intellect, Critical strike]`; Shadow: `[Spell power, Shadow power, Hit, Spell haste, Critical strike, Intellect, Spell penetration]`. | not fixed |
| PRIEST-12 | LIKELY | talent | discipline.md:60 | "Improved Healing (3 points): ... cheaper direct heals." | Improved Healing names Lesser Heal, Heal, Greater Heal, Penance and Prayer of Mending (`talents/priest.json`), not Flash Heal or Prayer of Healing. | "...faster Heal and Greater Heal, and 15% off Lesser Heal, Heal, Greater Heal and Penance." | not fixed |
| PRIEST-13 | LIKELY | race | shadow.md:111 | "Touch of the Grave procs Shadow damage off your hits" | `races.json` Undead: "gives spells and attacks a chance to drain Health from the target, up to 5% of your maximum Health, at most once every 15 seconds". | "Touch of the Grave drains a little health from your target every so often, which helps against Shadow Word: Death's backlash." | not fixed |
| PRIEST-14 | LIKELY | voice / talent | shadow.md:52 | "Nothing in the Discipline tree's damage row helps you: Power in Light only touches Smite and Penance, and Holy Precision only Holy spells." | Twin Disciplines sits in that row and is the guide's own first Discipline pick (engine mask covers Mind Flay, SW:P, DP, SW:D). | "Skip Power in Light and Holy Precision: one only touches Smite and Penance, the other only Holy spells." | not fixed |
| PRIEST-15 | LIKELY | voice | holy.md:36 | "Lightwell and Holy Binding heal while you are busy with something else" | Holy Binding heals you only when you cast Flash Heal on someone else (talent text). | "Lightwell heals the melee while you heal the tank, and Holy Binding heals you every time you Flash Heal someone else." | not fixed |
| PRIEST-16 | LIKELY | stats / voice | shadow.md:97 | Crit sentence names only Shadowform's doubled bonus | In Forever Shadow Word: Pain and Devouring Plague ticks crit and Mind Flay's channel crits (forever-dots-crit; `periodic_crit_auto_gen.go`), and Shadowform doubles every one of them. A top player says why crit is better here than in 1.12. | "Critical strike is worth more than it looks: in Forever your DoT ticks and Mind Flay crit, and Shadowform doubles every crit." | not fixed |
| PRIEST-17 | NEEDS-OWNER | voice | discipline.md:63, holy.md:64, shadow.md:67 | "Open the [planner](/planner?class=priest) to make the swap / adjust" | Owner rule: guides never mention the site; the prose test allows planner links. | Owner call O5; if links stay, keep them as bare links without "Open the planner". | not fixed |

## Sim findings (engine vs real game, next sim round)

- **S1 Healer haste negative.** 60 raid haste scale: Disc -8.0, Holy -5.4 per 1%; bare 60 Holy -29.3 (`bis/priest-discipline.json`, `priest-holy.json`). Faster casts cannot lower healing in the game; this is the mana-pacing model (already on the round-12 list). Guides do not mention haste for healers, so no prose change.
- **S2 SW:P end-of-fight cutoff.** APL `priest-shadow.json` skips the SW:P refresh with under 10 s left. Question: at 60 raid, does the cutoff at `remainingTime >= 6` (two ticks) or `>= 3` beat `>= 10` beyond error?
- **S3 Shadow raid dead points.** Question for a talent run: raid build with Blackout 2 + Spirit Tap 3 moved to Devouring Contagion 2 / Silence 1 / Shadow Reach 2 / Dark Inspiration 1 (no DPS loss expected; confirms they are free to move and Dark Inspiration's probe was -0.6/point on 70707, so check that point).
- **S4 Shadow bare 60 weights** rank mp5 and Spirit above spell power (bare-gear mana starvation). Not a guide issue; flagging so the BiS refresh does not leak it into prose.

## Owner questions

- **O1** Read Lesser Heal rank 3, Renew rank 1 and Heal rank 1 tooltips on a geared 60 (PRIEST-6). r14 builds rank-1 Renew/Heal advice on this.
- **O2** Shadow raid build: OK to move Blackout 2 and Spirit Tap 3 (PRIEST-10)?
- **O3** `trainables/priest.json` lists Shadowfiend with a rank-0 row at level 1 while `class-spell-limits.json` says not learnable. Guides say not learnable; confirm on a live trainer.
- **O4** Integrator for r14: drop Binding Heal and keep 70755 build prefixes (PRIEST-7).
- **O5** Are "Open the planner" links acceptable under the no-site-talk rule (PRIEST-17)?

## Fix log

Branch gr1/priest @ 49f95aed3 (on new origin/main with r14). Builds untouched.

- PRIEST-1: fixed (shadow.md:34, shadow.md:123; Spirit Tap 3/5 now described as "procs on most kills", plus a Spirit Tap talent line at shadow.md:61 stating the 60% rank and that on a boss it only fires off a dying VE add)
- PRIEST-2: fixed in prose (shadow.md:41, shadow.md:125): prose now says the build takes Shadow Focus, Blackout, Spirit Tap, Improved Mind Blast and Improved SW:P first, so Mind Flay waits until 28 (verified in 70755 addon-builds.json), and gives the legal questing path to Mind Flay at 20 (Spirit Tap 3, Blackout 2, ISWP 2, IMB 3 = 10 points, tier 2 open). Wand Specialization line reworded to "finishes with" (shadow.md:69; it lands at 41-42). Order change goes to SIM.
- PRIEST-3: fixed (shadow.md:125): "strip" and the Blackout-for-dungeon-bosses reason are gone; Blackout talent line (shadow.md:62) says it is useless on a boss.
- PRIEST-4: fixed (shadow.md:107)
- PRIEST-5: fixed (shadow.md:107, discipline.md:97, holy.md:96)
- PRIEST-6: softened (index.md:29, discipline.md:73, discipline.md:75, holy.md:70, holy.md:72; notes-only in priest-discipline.json Renew and party filler lines, priest-holy.json Renew and Heal filler lines). r14 had gone further (Renew rank 1, Heal rank 1 as fact); every line now states the 1.12 below-20 rule and says to read the tooltip on a geared 60. Counted confirmed (r14 regression on an unverified premise). APL actions/ranks unchanged; owner O1 still decides.
- PRIEST-7: fixed by r14 + 70755 activation (no Binding Heal in any priest guide; all build strings read 1.60.1.70755)
- PRIEST-8: fixed by r14 (shadow.md:75 opener has VE)
- PRIEST-9: fixed (shadow.md:85): verified SW:P r8 period 3000 ms (spellconst 10894), so "it will not tick out" in 10 s was false; prose now says keep refreshing until the last few seconds. APL note line drops "when it cannot tick out" (condition untouched). Exact cutoff = S2.
- PRIEST-10: left (build change). Prose already truthfully describes the raid build's trade; added Spirit Tap / Blackout talent lines that state what those points do on a boss. Proposal below.
- PRIEST-11: fixed (discipline.md:8, holy.md:8, shadow.md:9), verified against bis/priest-*.json band 60 raid weights on 70755 (data/builds/.../stat-weights.json is a community hand-weight file, not used)
- PRIEST-12: fixed (discipline.md:60), verified against Improved Healing text in talents/priest.json
- PRIEST-13: fixed (shadow.md:113), verified races.json Undead text
- PRIEST-14: fixed (shadow.md:52), Twin Disciplines verified tier 0 beside Power in Light
- PRIEST-15: fixed (holy.md:36), verified Holy Binding talent text
- PRIEST-16: fixed (shadow.md:99), per owner ruling forever-dots-crit and periodic_crit flag on Mind Flay
- PRIEST-17: left (NEEDS-OWNER O5; prose test allows planner links)

### SIM findings / build proposals (next sim round)

- B1 Shadow leveling build: Spirit Tap 5/5 and Blackout 2/5 (Shadow `54342001201300051` to `52542001201300051`), so "Spirit Tap after every kill" becomes true (PRIEST-1).
- B2 Shadow leveling order: Spirit Tap first at 10, ISWP, IMB, Mind Flay at 20, then Improved Mind Flay, VE, Shadow Weaving, Shadow Focus once dungeons start, Darkness, Shadowform 40, Wand Specialization (PRIEST-2). If B1 is not adopted, the legal 3/5 path is Spirit Tap 3, Blackout 2, ISWP 2, IMB 3, Mind Flay 20.
- B3 Shadow raid build: move Blackout 2 + Spirit Tap 3 into Devouring Contagion 2 / Silence 1 / Shadow Reach 2/2 / Dark Inspiration 1 (PRIEST-10, S3, owner O2).
- S2 SW:P end-of-fight cutoff (`remainingTime >= 10` vs `>= 6` / `>= 3`); prose now says "last few seconds" pending this.
- Low-rank heal rows (Disc Renew r1 / Lesser Heal r3, Holy Renew r1 / Heal r1) rest on O1; if the tooltip shows a penalty, revert the rank rows and tighten the prose to Renew rank 3 / Heal rank 4.
- S1, S4 unchanged from the audit.

## Fix pass 2

Integration tree only; no actions, conditions, ids, ranks or build strings touched.

- MUST 1/2 (healer ranks vs prose): fixed. Prose now names what the cards cast and states the below-20 caveat once per page as one instruction ("Check your low ranks once": read the tooltip gear on/off; if it barely moves, cast Renew rank 3 and Heal rank 4). discipline.md:73,75,77; holy.md:70,72,74; index.md:29 (one line). Card notes match, hedge removed: priest-discipline.json Renew ("Renew rank 1 on the tank...") and party filler ("Lesser Heal rank 1 on a party member..."); priest-holy.json Renew and Heal filler. Lesser Heal "?" icon and the duplicated Greater Heal / Lesser Heal rows on the Discipline card: left (card component / APL structure, logged for the integrator). Note: Discipline's top-level note still says Lesser Heal rank 3 (2053) while the rows cast 2050 rank 1; flagged, not touched.
- MUST 3 (Shadow SW:P cutoff): fixed. shadow.md:85 drops the "10 seconds still lands three ticks" claim; says keep it up until the closing seconds, then spend the global on Mind Blast / SW:D / Mind Flay. priest-shadow.json SW:P note: "Always up; stop refreshing it in the closing seconds of a fight." Neither names a number; exact cutoff stays S2.
- MUST 4 (hedge repeated five times, capital "In"): fixed; once per page (index.md:29, discipline.md:77, holy.md:74).
- MUST 5 (Shadow leveling, one plan): fixed. shadow.md:125 now gives only the build's order, verified against addon-builds.json priest-shadow leveling: Shadow Focus 10-14, Blackout 15-18, Spirit Tap 19-21, IMB 22-25, ISWP 26-27, Mind Flay 28, IMF, VE, Shadow Weaving, Darkness, Shadowform 40. shadow.md:41 "level 20 is the earliest" replaced with "the leveling build reaches it at 28, so until then you wand between casts" (also the filler SHOULD). shadow.md:34 Spirit Tap "from level 21".
- MUST 6 (raid Blackout / Spirit Tap called useless): fixed. shadow.md:61 Spirit Tap says what it does leveling and in a raid (trash kills, VE adds); shadow.md:62 Blackout states 8% / 4% stun from talent text and that in the raid build those points and Spirit Tap's are part of the 30 Shadow points Shadowform needs and pay out on trash; shadow.md:69 drops the "(bosses cannot be stunned)" aside.
- MUST 7 (planner links): left (R1).
- SHOULDs: hit cap fixed shadow.md:99 (16% spell cap, Shadow Focus 5%, 11% from gear; from bis/priest-shadow.json hit_to_cap spell 16 / school hit 5 / to_cap 11). Renewed Hope fixed discipline.md:53,71 (8% crit at 4 points, 4 s off Weakened Soul, talent text). Penance channel fixed discipline.md:33,52 (402261 duration 2000 ms). Skyborne context fixed index.md:48 (combos.json: Skyborne races 95/96 classes 1,3,4,7,8,11). "Fear Ward is no longer theirs alone" removed index.md:43. Flask/elixir stacking kept. Holy filler reason for rank 4 in dungeons rewritten (holy.md:72) without a mana claim.

### New SIM findings

- SIM-P1 (O1) Healer low ranks: Disc Renew r1 (139) / Lesser Heal r1 (2050) and Holy Renew r1 (139) / Heal r1 (2054). If the geared-60 tooltip shows the below-20 penalty, move the rows to Renew r3 (6075) and Heal r4 (6064); the prose already names that fallback.
- SIM-P2 Shadow SW:P end-of-fight cutoff (`remainingTime >= 10`): test >= 6 / >= 3. Prose and note say "closing seconds" until settled (replaces S2).
- SIM-P3 Shadow leveling order: Mind Flay at 20 (Spirit Tap / Blackout / ISWP / IMB to 10 points, then Mind Flay) vs the build's 28 on the ladder (B2).
- SIM-P4 Shadow raid build: Blackout 2 + Spirit Tap 3 are not a hard toll (SF5 + IMB5 open tier 3); test moving them to Devouring Contagion 2 / Silence 1 / Shadow Reach 2 / Dark Inspiration 1 (B3, O2).

## Owner questions resolved by research

Sources used throughout:
- WoW Forever Beta Development Notes (Blizzard, Kaivax), entries Oct 1, Oct 8, Oct 16, Nov 3: https://us.forums.blizzard.com/en/wow/t/forever-beta-development-notes/2360696
- World of Warcraft: Forever Class Deep Dives — Priest and Warrior: https://worldofwarcraft.blizzard.com/news/24301514
- ForeverChanges, "WoW Forever downranking: the level penalty and which ranks to keep": https://foreverchanges.pro/downranking ; downrank calculator: https://foreverchanges.pro/downrank-calculator
- ForeverWisp, "WoW Forever downranking: how low spell ranks lose power and procs": https://www.foreverwisp.com/guides/wow-forever-downranking-spell-ranks
- Master of Warcraft, "WoW Forever Just Nerfed Spell Downranking, and Healers May Feel It Most": https://www.masterofwarcraft.net/2026/10/wow-forever-spell-downranking-nerf-healers.html
- SpellCoeffsForever (CurseForge): https://www.curseforge.com/wow/addons/spellcoeffsforever
- Warcraft Tavern, Forever Priest Leveling Guide: https://www.warcrafttavern.com/forever/guides/priest-leveling/ ; Forever Priest Guide: https://www.warcrafttavern.com/forever/guides/priest/
- wow.gg, WoW Forever Shadow Priest (beta): https://wow.gg/guides/wow-forever-shadow-priest-beta
- Kami-labs, Shadow Priest Build — WoW Forever: https://kami-labs.fr/en/wow-forever/build-pretre-ombre-shadow-magic-wow-forever/
- wowtbc.gg, Shadow Priest PvE Guide - WoW Forever: https://wowtbc.gg/warcraftforever/class-guides/shadow-priest/
- 70755 data: `trainables/priest.json` (rank learn levels), `talents/priest.json` (Shadow tree tiers).

### O1 / PRIEST-6 / SIM-P1: low-rank heals at 60 (Renew r1, Heal r1, Lesser Heal r1)
**Answer: RESOLVED. Forever penalizes them, harder than 1.12.** Blizzard's Oct 1 note: on the character sheet, Spell Damage and Spell Healing now explain that "spells cast using spell ranks vastly below your current level receive less benefit from your Spell Damage and Spell Healing". The Nov 3 note removes the low-rank penalty only for talent procs ("Various talents no longer have a penalty to their chance to trigger"; still applies to Blizzard), not for the healing bonus. Community measurement (ForeverChanges, beta 1.60.1.70338/70532 to level 30, echoed by ForeverWisp): a rank keeps its full bonus for 17 levels after it is learned, loses 5% per level after, and has none 37 levels later; Lesser Heal keeps none at 60, Heal at most 55%. Not measured above 30, so exact percentages stay unquoted in prose; the direction is official. The old fixed below-20 rule is not how Forever works (no level-20 clause in the measured curve). By 70755 learn levels: Renew r7 (44) and up, Greater Heal r2 (46) and up, Flash Heal r5 (44) and up keep everything at 60; Greater Heal r1 (40) keeps most; Heal (top rank learned 34), Renew r1 (8), Lesser Heal (1-10) keep little or nothing.
- Prose changed: index.md:29; discipline.md:73, 75, 77 (Renew rank 7+, party filler Greater Heal rank 1 or 2, "Low ranks" paragraph); holy.md:29, 70, 72, 74, 76, 80, 82, 84 (Renew rank 7+, raid filler downranked Greater Heal, Heal rank 4 only on bare gear or in dungeons, "Low ranks" paragraph, every "Heal" filler mention now Greater Heal).
- **RESOLVED, needs sim round:** (a) engine: implement the Forever low-rank healing-bonus curve (full for 17 levels after the rank's learn level, minus 5% per level, zero at 37; apply to healing; damage spells likely the same, unverified) and drop "no downrank penalty" from the Disc/Holy top-level notes once it lands; (b) Disc card: Renew r1 (139) to r7+ (10927 or higher), party filler Lesser Heal r1 (2050) / top-level note's Lesser Heal r3 (2053) to Greater Heal r1/r2 (2060/10963); (c) Holy card: Renew r1 (139) to r7+, raid Heal r1 (2054) filler to Greater Heal r1/r2 (2060/10963); rerun the rank ranker with the curve on. Card notes still describe the r1 casts; rewrite them with (b)/(c). Until then prose and cards disagree on these lines; prose is the game-correct side.

### SIM-P3 / PRIEST-2 / B2: Shadow leveling order, Mind Flay at 20 vs 28
**Answer: RESOLVED, Mind Flay at 20.** 70755 tree: Shadow Focus, Blackout and Spirit Tap are tier 0, Improved Mind Blast / Improved SW:P / Shadow Reach tier 1, Mind Flay tier 2 (10 points), so the 11th point at level 20 is legal. Warcraft Tavern's Forever leveling guide plays Mind Flay "after level 20"; Classic consensus has always been Spirit Tap first, Mind Flay at 20. A 1.5 s channel filler from 20 instead of eight more levels of wanding is not a taste call.
- **RESOLVED, needs sim round:** leveling order in the build/addon-builds: Spirit Tap 10-14, Improved Shadow Word: Pain 15-16, Improved Mind Blast 17-19, Mind Flay 20, Improved Mind Flay 21-22, Vampiric Embrace from tier 3 as points allow, then Shadow Weaving, Shadow Focus, Darkness, Shadowform 40, Blackout/Wand Specialization last. Prose not changed: shadow.md:41 and :125 describe the build players import (fix pass 2 MUST 5), so they flip with the build. Replacement for :125 when it lands: "Spirit Tap from 10 to 14, Improved Shadow Word: Pain and Improved Mind Blast to 19, and Mind Flay at 20. Improved Mind Flay, Vampiric Embrace and Shadow Weaving follow, Shadow Focus once you run dungeons, then Darkness, and Shadowform is your point at 40." For :41: "Mind Flay is a talent in Forever, your point at 20; until then you wand between casts." For :34: "from level 10 Spirit Tap ...".

### PRIEST-1 / B1: Spirit Tap 3/5 vs 5/5 leveling
**Answer: RESOLVED, 5/5.** wow.gg's Forever build takes Spirit Tap 5/5 by 20 and calls it "the main leveling talent"; Classic consensus is the same. 70755: rank 5 is a 100% chance, rank 3 60%.
- **RESOLVED, needs sim round:** leveling build Shadow `54342001201300051` to `52542001201300051` (Spirit Tap 5, Blackout 2), together with the order above. Prose (shadow.md:34, :61, :123) describes 3/5 truthfully today; on the flip, :61 becomes "Spirit Tap (5 points): after every kill your Spirit doubles for 15 seconds..." and :123 "Spirit Tap after every kill lets you chain pulls without drinking."

### O2 / PRIEST-10 / SIM-P4 / B3: Shadow raid Blackout 2 + Spirit Tap 3
**Answer: RESOLVED, move them.** No Forever raid build is published yet (wowtbc.gg and Kami-labs builds are placeholders; Kami-labs notes Devouring Contagion halves Devouring Plague's cost and spreads it). Forever's tree is not 1.12's: Shadow Focus alone opens tier 1 and SF5 + IMB5 open tier 2, so these five points are not a toll; the only gate they serve is Improved Psychic Scream's Blackout prerequisite (Deep Dive), which a raid build does not take. Blackout cannot stun a boss; Spirit Tap fires only on kills. Devouring Contagion's 50% cheaper Devouring Plague is mana every 24 seconds on every boss.
- Prose changed: shadow.md:62 drops the false "toll toward the 30 Shadow points" claim ("on a boss the raid build's two points do nothing; they pay out only on trash").
- **RESOLVED, needs sim round:** raid build Blackout 2 + Spirit Tap 3 into Devouring Contagion 2, Shadow Reach 2/2, Silence 1, Dark Inspiration 1 (legal per the audit's tier check); confirm Dark Inspiration's point is not negative, else Silence/other. Then shadow.md:61-62, :69 per PRIEST-10.

### S2 / SIM-P2 / PRIEST-9: Shadow Word: Pain end-of-fight cutoff
**Answer: RESOLVED, 6 seconds.** SW:P r8 ticks every 3 s; two ticks (2 × (127 + 0.2 SP)) out-damage the Mind Flay global it replaces (1.5 ticks of 130 + 0.167 SP with Improved Mind Flay's ramp), one tick does not. Kami-labs' Forever sim APL casts SW:P only while "the remaining fight time is at least 6 s".
- Prose changed: shadow.md:85 ("until the last 6 seconds ... it takes two ticks to out-damage the Mind Flay you would cast instead").
- **RESOLVED, needs sim round:** `priest-shadow.json` SW:P condition `remainingTime >= 10` to `>= 6`; then note to "Always up; stop refreshing it in the last 6 seconds of a fight." (note left as "closing seconds" so it matches the current card).

### O3: Shadowfiend learnability
**Stays with the owner (fact, not taste; the web cannot settle it).** No Forever source (Blizzard dev notes, Deep Dive, Warcraft Tavern, Kami-labs, wowtbc.gg) mentions Shadowfiend; 70755 has a rank-0 row at level 1 with no skill-line learn level (`class-spell-limits.json` notLearnable), and no beta priest cast it. Recommendation: keep "cannot be learned at launch"; check a level 60 trainer on launch day (11-04).

### O4 / PRIEST-7: r14 Binding Heal and build prefixes
**Resolved already.** Blizzard Nov 3 notes: "Binding Heal removed", replaced by the Holy Binding talent. Main has no Binding Heal and reads 70755 in every build string. No change.

### O5 / PRIEST-17: "Open the planner" links
**Stays with the owner (taste/house rule).** Options: (a) keep as is; (b) keep bare links, drop "Open the"; (c) remove (discipline.md:63, holy.md:64, shadow.md:69). Recommendation (c): the guide-voice rule says no site talk, and the sources block already links the planner.
