SHIP

Second pass. Seen rendered: healer-raid-1440-full, healer-raid-390-full, healer-gear-1440-full, healer-gear-390-full. Notes copy read in web/src/lib/tiers/tier-copy.ts (the notes columns are in no capture, so that part is judged from the text only).

Pass-1 fixes, checked:
1. Raid lead names the fight: "Ranked on one healer's share of a Molten Core or Onyxia kill, so it reads lower than your meter." Which raid is answered, and the share and the meter question are answered in the same line. Done.
2. Fallback lead: "Ranked on a stand-in fight, not a boss, that runs every healer dry. A Molten Core or Onyxia kill replaces it with the next update." A raider understands this is a placeholder and what replaces it. The four-figure HPS (1,028.3) now has a visible reason. "BiS fight" and "ranking adds" are gone. Done.
3. Units: the meta line owns the length ("110 s" raid, "5 min" fallback) and the lead carries none, so no mixed units remain. Done.
4. Notes raid column: "eleven healers" gone, "the healer's own group of five takes the rest", "Phase 1 gear, raid buffs, no world buffs", and "a share of the fight, so it reads lower than your own meter" are all there. Reads like a player wrote it. Done.
5. Order flip between states is explained by the fallback lead saying a kill replaces it.

Layout: lead is three lines at 1440 and four at 390 in both states; at 390 the table starts about 390 px down, well on the first screen. No clipping.

Data: the raid order (Shaman, Holy Priest, Disc, Druid, Paladin) matches logs at the top; the Paladin miss is the known one and is left alone. Druid over Paladin by 2.3 HPS against 0.27 combined error is beyond error, so no tie showing is right. Not a defect here.

Optional, not blocking:
- The short caveat dropped the "death it stopped" example; the long form with Power Infusion lives in the notes, which is the right place for it. If a Disc player complains, put "Power Infusion" back in the notes headline, not the lead.
- Raid lead "so it reads lower than your meter" is a little blunt as a cause ("Ranked on ... so ..."). Acceptable as is.
No em or en dashes, no sim talk, no nightly numbers in the lead.
