# Holy Paladin guide, sim-accuracy lane: wow-player review (2026-11-12)

Surface: prose and card notes only (web/src/content/guides/paladin/holy.md, data/curated/apl/paladin-holy.json). Judged as text, read start to finish.

## 1. The question and whether the text answers it

Question: "What do I cast on the tank and the group at my level, and does my mana last?" The new rank lists answer it at last: a leveling Holy paladin can now bind rank 1 / 2 / 4 by level instead of guessing. The Illumination and Concentration Aura corrections are right for a player and cost nothing to read. The rank answer is, however, the weakest-presented part of the page (see MUST 1 and 2).

Checks run: no em or en dashes anywhere in the guide; no "the sim", "measured" or "re-judged"; no mention of the site owner. One meta phrase found ("live server", line 54).

## 2. Verdict: SHIP WITH FIXES

### MUST

**M1. The group Flash of Light line stops at 49.** The new paragraph (line 112) gives the rest of the group rank 2 "from 44 to 49" and then nothing. The lane's own note says rank 4 from 50, and the top of the page promises rank 5 at 60 in a raid. A player at 50 to 60 reading this paragraph has no answer for the group. Also the Human, Dwarf and Undead clauses are in a different order from the tank sentence and the Undead clause ("rank 1 until 43 if you are Undead") reads as a contradiction of "rank 1 until 41" two clauses earlier.

**M2. "until 42" is ambiguous, and the paragraph is a wall in the wrong section.** "Rank 1 until 42" reads as "up to but not at 42"; the data says rank 1 at 42 included. Use "through". The paragraph also sits under "Levels 1 to 19", but Flash of Light is learned at 20 and the paragraph runs to 60. Move it under "Levels 20 to 60" (before the talent order) and set it as a list like the Holy Light list above it, so a player can find their level by eye. Replacement:

> Flash of Light by level, with the rank to bind:
>
> - **On the tank:** rank 1 through 42 (a Dwarf keeps rank 1 at 43), rank 2 from 43, rank 3 from 50, rank 4 from 54 (an Undead takes rank 4 at 53). At 60 it is rank 4 on your own buffs and rank 5 in a raid.
> - **On the rest of the group:** rank 1 through 41. Rank 2 from 42 as a Human, from 43 as a Dwarf, from 44 as an Undead. Rank 4 from 50.

(That puts the group's 42 to 49 band as "rank 2" implicitly; add "rank 2 through 49" after the race clause if you want it spelled out: "Rank 2 stays through 49, and rank 4 takes over from 50.")

**M3. Line 54 breaks the voice rule.** "never lands on the live server" is bug-report talk, and a veteran reads "the live server" as meta. Say what the player sees. Replacement:

> - **Illumination** (5 points): every critical heal gives back half the spell's base cost. The tooltip also promises 15 seconds of Spirit regeneration while you cast; that part does not happen, so count Illumination as a mana refund and nothing more. It is the reason crit is a mana stat for you.

### SHOULD

**S1. Stat priority wording (line 85).** "Illumination gives you mana on a crit and no regeneration" reads as if Illumination gives nothing. Replacement: "Spirit only works once you have gone five seconds without a cast, apart from Reverence's 30%. Illumination pays you in mana on a crit, not in regeneration."

**S2. Devotion Aura clause (line 69).** "three hits in ten costing you a second of casting" overstates: only the first hit costs a full second. Replacement: "...where Devotion Aura lets three hits in ten through, the first costing you a second of casting." Also "Run Concentration Aura" now drops the old "if nobody else in your group runs it". If another paladin already has Concentration up on your group you are covered, so keep that condition: "Run Concentration Aura unless another paladin already has it on your group: its 35% and ..."

**S3. Strengths bullet (line 34).** "a hit while you cast never costs you casting time" is only true under Concentration Aura. Replacement: "Spiritual Focus and Concentration Aura together stop every pushback: under the aura, a hit while you cast never costs you casting time".

**S4. Oath wording versus rank 1.** Lines 31 and 57 say Oath of Light works "as long as the rank is close to your level", and line 73 says rank 1 is a trap at 60. The new list tells the tank healer to use Flash of Light rank 1 through 42, a rank learned 22 levels earlier, with the Oath on. A player will spot that in one read. Add one sentence to the Ranks paragraph, for example: "The further a rank sits below your level, the more of your healing power it loses, so the cheap rank only wins while the gap is small; past the breakpoints below, move up." (Accuracy reviewer: confirm the gap rule is what the engine does.)

**S5. "Holy Shock on the party ... gets a large part of its cost back" (line 71).** Half the base cost on a crit only, so on average it is a share, not a large part. Replacement: "gets a good share of its cost back through Illumination."

**S6. Why the races split.** "a Dwarf keeps it at 43", "an Undead takes rank 4 at 53" with no reason looks arbitrary. If the cause is the racial Intellect and mana pool (it must be), say it once: "The races differ by a level or two because their mana pools do." Only add if the lane confirms the cause.

### Consistency sweep (nothing else stale)

- Common mistakes: "Rank 4 or 5 under Oath of Light does the work for less mana than the top rank" is consistent with rank 4 bare, rank 5 raid. Fine.
- Ranks paragraph: "rank 4 does the tank's job for less again outside a raid" matches the APL note. Fine; the APL says "without raid buffs", either is ok.
- Holy Light list and the paragraph above it still say "rank 3's 110 mana" and "Rank 3 heals more per cast"; at 22 to 28 your top rank is 4, so "your top rank" would be safer than "rank 3" in the sentence at line 105. Minor, leave if the 110 figure is the point.
- Talents bullets, Strengths, Weaknesses, Races, Professions: nothing else mentions the dead Illumination regen buff.
- APL notes: Holy Light (18 to 28, rank 2 at 29), tank Flash of Light and group Flash of Light notes agree with the guide. The tank note leaves out the Undead 53 exception that the guide has; harmless.

### Data I do not believe (for the accuracy reviewer, not new in this lane)

- Line 73: Holy Shock rank 1 "is a level 30 rank" and "Use Holy Shock rank 3 or 4". Holy Shock is a level 40 spell in Classic (ranks at 40, 48, 56). The leveling section teaches it at 45, and the lane's fact 4 says rank 1 holds "40 to 49". If Forever really moved it to 30 and has four ranks, fine; if not, the sentence is wrong. Needs a client check.
- Line 93: "Ruthless" as a Human racial name is not one I know; unchanged by this lane.

## 3. Summary

The corrections that matter (Illumination is mana only; Concentration plus Spiritual Focus is full pushback immunity; rank breakpoints) are right in direction and read well. Fix the Flash of Light block (M1, M2: missing group 50+, "until" wording, wrong section, wall of text) and the meta phrase (M3), take S1 to S3 for accuracy of the wording, and this ships.

---

## Second pass (2026-11-12), re-read of `git diff aafe382b5 -- holy.md`

**Verdict: SHIP.** No MUST items remain.

Checked against the first pass:
- M1: the group Flash of Light line now runs Human 42, Dwarf 43, Undead 44, rank 2 through 49, rank 4 from 50. The tank line gives 60 as rank 4 on your own buffs and rank 5 in a raid. Complete, and it matches the APL notes.
- M2: "through" wording used, the block sits under "Levels 20 to 60" as two bullets.
- M3: the Illumination bullet now tells the player what the tooltip promises and that it does not happen. No "live server", "the sim", "measured" or "re-judged" anywhere in the guide.
- S1 to S5 applied as written. The Oath-versus-rank-1 question is answered by the new Ranks sentence; the numbers hold up (Flash of Light rank 1 is learned at 20, so at 42 it is five levels past the 17-level grace and keeps three quarters).
- Data notes withdrawn: Holy Shock's first rank at 30 and Ruthless as the Human racial are the client's own rows, so the line 73 "level 30 rank" and line 93 stand.
- No em or en dashes in the guide or in this file.
- S6 (why the races split) left out because the cause is unconfirmed. Agreed.

Optional polish, not blocking:
- Line 124 now says "Flash of Light at level 20 is enough to heal most dungeons at your level, even in a Retribution build. Flash of Light by level, with the rank to bind:". The repeated "Flash of Light" is clumsy. Suggested: "...even in a Retribution build. The rank to bind, by level:". The trailing "Questing as Holy is slow either way." after the list is fine.
- Line 105 still says "rank 3's 110" while your top Holy Light is rank 4 from 22 to 28; "your top rank's 110" is not true then either, so leave as is unless the figure is re-checked.
