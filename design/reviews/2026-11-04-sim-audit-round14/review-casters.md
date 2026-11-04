# Round 14 wow-player review: healers, casters, platform text, data, guides

Reviewer: wow-player (veteran Classic raider and healer). Read-only pass on int/r14 (merged tree). Read: healers.md, regcast.md, casters.md, platform.md, data.md, guides.md; the rendered priest, druid, shaman guides; the notes in the priest, druid-restoration, shaman-restoration, mage-fire and mage-frost APLs; data/curated/coach/priest-shadow.md. No page was screenshotted here, so these are text and data verdicts only.

## Verdict: SHIP WITH FIXES (five text fixes, one regeneration, one owner check before launch)

## The two questions

**Would a real Holy priest raid on rank 1 Heal and rank 1 Renew?** Not in Era. Every Era healer knows the sub-20 penalty, and nobody raided on a level 8 Renew. But the Forever rule (a heal learned below 20 keeps its whole spell power bonus, round 11 C1) is the only reason it works, and with that rule it is the old downranking trick pushed to its end: a big bonus on a tiny mana bill. I believe it only if the client shows it. The whole row set (Renew 1, Heal 1, Lesser Heal 1, Rejuvenation 1, Healing Wave 2) rests on that one rule, and the owner's tooltip check (healers R3, ruling 5) is still pending. The guides already state it as fact ("Forever keeps a low rank's whole healing bonus"). If a player opens Renew rank 1 on a geared priest and the tooltip says otherwise, that is the screenshot that gets us mocked in guild chat. Owner call 1 below.

**Is a rank-2 Healing Wave tank heal believable?** Same answer. In Era, rank 2 (level 6) would be a joke on a raid tank. Under the Forever rule it is believable, and +0.7% is a small gain to hang a guide sentence on. The real problem is the text: see G2.

## Per row

### healers.md
| row | verdict |
|---|---|
| A1 heal scale 1.85 to 2.3 | SHIP. The player-visible result is that healer HPS figures and the tier order move (druid up, shaman down). Do not quote any number in prose. The profile notes already say the strict clause (bare under half) won over the floor; keep that. |
| A2 level-60 split from logs (blocked) | SHIP as is. Level 60 keeps 0.659; nothing should claim a logged tank share. One kill is not a sample. |
| A3 Discipline 0% haste tie | Not mine. F5 (open, core) is invisible to a player. |
| P1 Holy tank Flash Heal rank 7 | SHIP the row. FIX the note and the guide, see G1 (the "mana outlasts the fight" wording is false after this very change). |
| P2 Renew rank 1 at 60 | SHIP, conditional on owner call 1. |
| P3 Holy Heal rank 1 in a raid, rank 2/3/4 rows kept for bare bands | SHIP, conditional on owner call 1. Rows keep what each bare level really casts, which is what a player sees on the bar. |
| P4 fresh read (PoM, Lightwell, Binding Heal) | SHIP. Matches what I know of the client: PoM 5 charges, 10 s; Lightwell 5 clicks, 10 minute cooldown. |
| P5 Inner Focus on Penance pulses | SHIP as OPEN. One Penance under Inner Focus in game settles it; under 0.1%. |
| D1/F1 Discipline Renew rank 1, Lesser Heal rank 1 | SHIP, conditional on owner call 1. The Discipline guide already says "a heal learned below level 20 keeps all of its spell power bonus", and Lesser Heal rank 1 is level 1. |
| D2-D4, H1-H3 | SHIP. |
| F2 Restoration shaman tank Healing Wave rank 2 | SHIP the row; FIX the note and guide, G2. |
| F3 Holy Shock rank 4 at bare 60 | SHIP. |
| F4 not adopted (shield rank 10, Flash Heal 2, Holy Light 3) | SHIP. The shield reason ("a full-rank shield on a warrior tank eats the rage he needs") is a real raid reason a player will nod at. |
| S1-S3, R1-R4 | SHIP. |
| X3 does the coverage rule apply per band? | OWNER (call 3). |

### regcast.md
| row | verdict |
|---|---|
| H1 remove Holy's band-50 Heal rank 1 row | SHIP. A 50-59 priest casting rank 4 Heal is exactly what the guide has always said. |
| H2 other Holy rows | SHIP. |
| X1/L5 Shadow coach starter | FIX. The committed `data/curated/coach/priest-shadow.md` on int/r14 still has the old text: "Inner Focus into Devouring Plague outside the execute window", "Mind Blast only when Mind Flay cannot be cast", "Shadow Word: Death first below twenty percent". Platform C2 fixed the generator, but platform's own "Files to regenerate" says the file is not regenerated yet. Run `python tools/coach/persona_starter.py --all`, then restamp `reviewed:` (it says wow-player 2026-10-28 and will be wrong after the regeneration). The new text must say what shadow.md says: Inner Focus on Mind Blast, Mind Blast and Shadow Word: Death on cooldown, Mind Flay with everything else, never clipped. |

### casters.md
| row | verdict |
|---|---|
| M1 Combustion stacks to 10 | SHIP. Nothing shows: three crits end it, and 10 stacks is 100% crit anyway. The Fire note ("until three non-periodic crits land") reads right. |
| S1 Inner Focus no longer crits Shadow Word: Death | SHIP. The Shadow note ("a free Mind Blast with 25% more crit chance") and the guide say Mind Blast, never Death, so no text moves. |
| S2 and the other aligned rows | SHIP. |

### platform.md
| row | verdict |
|---|---|
| F1-F3 stale fixtures | SHIP (not player-facing). |
| L1 land.sh | SHIP (not player-facing). |
| D1 Kindling Stave proc timer across iterations | SHIP. Max 0.03% on three mage and warlock numbers. |
| C1 coach book ranks | SHIP. Aspect of the Hawk rank 6, Serpent Sting rank 8, Corruption, Immolate, Battle Shout below the Ahn'Qiraj books: right, a fresh 60 cannot train a book rank. Regenerate the seven coach JSON files listed so the test passes. |
| C2 Shadow starter | SHIP the generator; the regeneration is still owed (see regcast X1). |
| S1-S5 | SHIP. S5 (is the addon's 60 raid or bare) is OWNER, call 4. |

### data.md
| row | verdict |
|---|---|
| P1 Trueshot Aura text "100" | SHIP. Matches the 11-03 push (rank 5, 20906). |
| K2 Devouring Plague rank 2 only for Dwarf, Night Elf, Troll at bare 60 | SHIP. A priest on rank 2 Plague at 60 is odd to Era eyes, but it is a mana row and it names its races. Wherever it shows, it must say "bare 60". |
| K3 Corruption rank 3 at 40 removed | SHIP. |
| K4 others | SHIP. |

### guides.md
The 40+ prose rows (T1-T12, G1-G7, W1, R1-R4, D1-D4, P1-P5, S1-S3, F1-F3) are SHIP. Specific mentions: T1/T2 Read Leyline wording now reads like a player wrote it; T3/T6 Vampiric Embrace ("heals your whole group for 20% of every Shadow spell you land; your own damage meter never shows that healing") is right and clear; H1/H2 Aimed Shot "doubled its bonus damage" is the honest wording; P1 Conviction on spells is what the client says; W1 Charge not usable in combat is the real reason (Bloodrage). Penance "comes at level 30": checked, Penance sits in tier 4 of Discipline (20 points), so the 21st point is level 30. Correct.

Meta-talk sweep: I searched every guide body and every `rotation` note for sim, simulator, ladder, ranker, tier list, published, "the guide", "the line", "damage figure", "custom rotation" and "the table". No violation in any class. The only hits are in test files and hunter "Engineering".

## Fixes required (exact wording)

**G1. Holy Flash Heal note and guide say something the engine no longer supports.** P1 made the fight run dry at about 115 s, yet:
- data/curated/apl/priest-holy.json, Flash Heal on the tank: "in a raid your mana outlasts the fight, so a big fast heal is worth its cost."
- web/src/content/guides/priest/holy.md, "Flash Heal is for emergencies": "your mana outlasts the fight, so a big fast heal is worth its cost."
These also contradict holy.md's own Mana section ("Every heal you skip early is mana for the last minute"). Replace both with: "Flash Heal at the top rank on a tank who is really low or on anyone about to die: it is your dearest heal, so spend it where the size matters." The guide sentence: "In a raid use the top rank, on a tank who is really low or anyone about to die: it is your dearest heal, so spend it there and let Prayer of Mending, Renew and Heal do the rest. While leveling, a lower rank keeps your bar alive."

**G2. Restoration shaman Healing Wave: the note contradicts the row, and the guide never mentions the rank.** data/curated/apl/shaman-restoration.json, Healing Wave at 60%: "your biggest efficient heal. Rank 9 is the top rank you can learn at launch; while leveling, cast the highest rank you have." Two faults: the row casts rank 2 (332), and rank 9 is an Ahn'Qiraj book (book-ranks.json, guides lane "Holy Light and Healing Wave rank 9 are books"), so the top trainer rank is 8. A shaman who reads this and sees rank 2 on the bar will assume a bug. New note: "Healing Wave when the tank drops below 60%, at rank 2 in a raid: Forever keeps a low rank's whole healing bonus, so it heals for most of a top rank's amount at a fraction of the mana. While leveling, cast the highest rank you have." Also the Nature's Swiftness follow-up, "the instant, full-size save", becomes "the instant save: rank 2 is cheap, and it still carries your whole healing bonus." In shaman/restoration.md, "Tank first" item 3, add after "Healing Wave when the tank is down to 60%.": "Cast it at rank 2 in a raid; a low rank keeps all of your healing bonus for a fraction of the mana, and Lesser Healing Wave is the fast heal for when he is falling." This sits beside the existing "Chain Heal ranks" paragraph the same way.

**G3. Holy filler wording.** holy.md and the Heal note say "without raid buffs and while leveling, rank 4". Reads as if raid buffs change the best rank. Say: "In a raid, cast rank 1 ... Without a raid's mana (bare gear, a dungeon group) and while leveling, rank 4 is the classic downrank." Keep the data as is.

**G4. Coach starter regeneration** (see regcast X1). Plus the other six coach JSONs from platform C1.

**G5. Coach starters carry meta-talk the new guard does not scan.** The starters say "This is the Shadow Priest as the site's guide describes it", "the buttons the guide's rotation names", "the guide keeps it up in a raid". If a coach starter is ever shown to a player or fed straight to the coach voice, that is the same defect round 14 just banned. Extend the `simTalk` guard (or a twin) to data/curated/coach/*.md bodies and reword the generator's template: "This is the Shadow Priest as the best players run it." If the files are only an internal prompt and never rendered, say so in the generator and skip this.

## Things I do not believe (for the accuracy reviewer)
- shaman/restoration.md "Chain Heal ranks. Rank 3 costs 405 mana for a 490 base heal; rank 1 costs 260 for 251." My Era memory is ~450-500 average for rank 3 and ~340 for rank 1, so the rank 1 figure looks low against Era. Forever beats my memory if it is the client row; please confirm 251 and 490 against spellconst, and that these are base numbers, not averages.
- The Fear Ward line ("Brings: ... Fear Ward") for every priest: Fear Ward was a Dwarf racial in Era. The guide's Dwarf line leaves Fear Ward out, which only fits if the client made it baseline. Please confirm against the trainer table.

## Owner calls I recommend
1. **Before launch tonight, one in-game check, five minutes.** On a character with healing gear, read these tooltips: Rejuvenation rank 1, Renew rank 1, Heal rank 1, Lesser Heal rank 1, Healing Wave rank 2. Compare the bonus-healing line to the top rank's. If each shows the full bonus, ship everything above. If any shows a penalty, revert P2, P3, F1, F2 and the druid's rank 1 row, and remove the sentences "Forever keeps a low rank's whole healing bonus" (priest holy.md, discipline.md; APLs priest-holy, priest-discipline; druid restoration.md "Ranks"; druid-restoration and shaman-restoration notes). Write the revert as one prepared commit so it can go in the first hour.
2. If the check cannot be done by launch, ship the rows and soften the prose to what the client tooltip will show: "Forever puts the full healing bonus on spells learned below level 20."
3. X3: do not apply the coverage rule per band before launch. The druid at 2.3 covers 50 / 77 / 85 / 67% of bands 20 / 30 / 40 / 50. That only affects the weights ranking at those bands, which nobody gears against on day one. Revisit after the first week's logs.
4. S5: the addon's level 60 should stand for the raid character (it is what a 60 plays at launch week); keep the bare 60 rows out of the addon rotation.
5. A2: leave 0.659 and say nothing about a logged split. Run the WCL steps after `gcloud auth login`.

## Release-note lines (player voice, no numbers that move nightly)
- Healers: the fight we rate healing gear against is heavier, so healing power counts for more and mana for less in the weights. Restoration druid gear ranks higher, Restoration shaman lower; no spec's pick list is empty.
- Holy priest: in a raid, Heal and Renew now use rank 1 and Flash Heal on the tank uses the top rank. Between level 22 and 59 the Heal rank follows what your level teaches.
- Discipline priest: Renew and the filler Lesser Heal now use rank 1.
- Restoration shaman: the tank's Healing Wave is rank 2 in a raid.
- Holy paladin: Holy Shock rank 4 for a freshly dinged 60.
- Shadow priest: the Coach no longer tells you to hold Mind Blast behind Mind Flay; Mind Blast and Shadow Word: Death go on cooldown. Inner Focus no longer adds crit to Shadow Word: Death, as in the client. Devouring Plague rank 2 at an ungeared 60 applies only to Dwarf, Night Elf and Troll.
- Mage: Combustion's crit bonus stacks to 10, as in the client; no damage change.
- Warlock: the Corruption rank 3 row at level 40 is gone.
- Hunter: Trueshot Aura now reads 100 attack power at 60; the Coach no longer tells hunters to train Ahn'Qiraj ranks.
- Guides: every class guide was reread for facts and rewritten where the game or the 3 November hotfix changed them.
