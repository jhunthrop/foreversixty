# L5 Guild tab, brand consistency vs the web raid page

Reference: build-raid-1440.png, build-raids-1440.png, web/src/lib/guild/raid-copy.ts, RaidChip.svelte.
Captures reviewed: raid-day, queued, old-companion, no-companion, no-pack.

## ALLOWED (platform idiom or already matching)
- In/Late/Bench/Out as client buttons with a teal selected state, not the web's underlined segments: native control, same as the shipped Builds tab.
- Section headings in the title face at windowSection, sentence case (Tonight, Your prep, Sign-ups, Next raids, Officers sent): matches the web's display-face panel headings and N1.
- Two columns, hairline rule, no card chrome, flat plates: matches the web's dense system.
- Unsigned in gold, a coming answer in success green, warnings in wipe red: matches RaidChip.svelte (unsigned gold, coming success).
- Role groups Tanks / Healers / Melee / Ranged, "Your job:", "Not signed up", "Sign-ups", "Notes-style" muted prose, class-coloured names, tier-coloured scores, quality-coloured item names with the client tooltip: same words and roles as the web.
- Dates "Thu 10 Dec, 20:00" and the shared clock text (12 h on a 12 h client) follow the web's weekday-day-month shape; the client clock stands in for "your time".
- Every colour goes through Theme.HEX / tokens, every string through Locale `guild*`; no raw hex or literal copy in GuildHubTab.lua.
- Empty, old-companion and no-pack states leave sections out rather than showing blanks; the retired styles stay retired.

## DRIFT
1. SHOULD. Character rows in the standings (drawStanding, GuildHubTab.lua) use a 14 px square spec icon. The web shows every character as the circular ringed HD crest. Use the existing medallion (Widgets.lua crest/ring at ~16 px, Theme.crestPath(row.classToken), ring asset Theme crestRing) in place of BuildsPane.specIcon in that column. Keep the square spec icon on the Officers sent build row only (it is a build, as on the Builds tab).
2. SHOULD. Sign-up count wording. The web says "19 of 40 signed up" (raidCopy countLine) and "Nobody signed up yet"; the addon says "12 of 40 coming" (L.guildComing, L.guildComingAll). Change both keys to "%d of %d signed up" and "%d signed up", and keep "16 not signed up" for the rest.
3. NICE. The web shows comp shortfall as a Need / In / Short table with Short in wipe red. The addon's "Tanks 2 of 4 · Healers 4 of 10" line is all muted, so a short role reads as quiet as a full one. Colour a role warning (Theme.HEX.warning) when its count is under its target.
4. NICE. "Your job:" value is gold, which on the web means unsigned/needs action. The web prints the job name plain. Use "strong" for the job name so gold keeps its one meaning (the Not signed up answer and the Guild dot).
5. NICE. Column head "Att." abbreviates what the web writes out; use the full word if the 40 px column can take it, or keep and note it as the width exception.

VERDICT: FIX

## Round 2
- Drift 1 fixed: standings rows use the ringed class crest (Theme.buildCrest / paintCrest by classToken); the square spec icon stays only on the Officers sent build row. Verified in code and l5-raid-day.png.
- Drift 3 fixed: roles coloured short/met.
- Drift 2 (signed up vs coming), 4 (gold job) and 5 (Att.) are owner-approved spec copy/colour and a width exception (spec 11, 3.4, 3.5; 288 px column); accepted as ALLOWED, not must-fix.
- No new drift seen in the redrawn renders.

VERDICT: PASS
