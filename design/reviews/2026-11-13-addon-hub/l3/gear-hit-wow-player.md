# Gear hit lines and Conditions row: wow-player

Verdict: SHIP, with shoulds.

Checked against data/builds/1.60.1.70938/bis/*.json hit_to_cap.

Figures:
- Hunter BM, last band: specials 2, spell 9. The render reads "Hit: 2% to the 9% cap" and "Spell hit: 9% to the 16% cap". Matches.
- Fury, Bare band: specials 3, white 22. The render reads "Hit: 3% to the 9% cap" and "White swings: 22% to the 28% cap". Matches.
- The caps are right for a level-63 boss: 9 yellow, 28 dual-wield white, 16 spell.

Musts: none.

Shoulds:
1. The weights table is "per rating point" while the lines are in percent, so the player has to know 10 hit rating = 1%. Add that to the hover: "10 hit rating = 1%."
2. The lines are the list's character, not the player's own gear. A player reads "Hit: 2% to the 9% cap" as their own gap. The hover says "Counted with this list worn", but the visible line doesn't. Consider "with this list: Hit 2% to the 9% cap".
3. The hover quotes "Improves your chance to hit by 1.0%.". The client writes "Equip: Improves your chance to hit by 1%." Check the exact string against the Forever client and add the "Equip:" prefix if the client shows it.
4. The hit lines sit tight under the last weight row and look like a table row. A small gap or a divider would separate them.
5. A hunter's "Hit" is the ranged-special line. "Ranged hit" would match how hunters say it.
