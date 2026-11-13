# L7 Officer tab, brand consistency vs the web guild roster and the Guild tab, round 1

Checked: l7 renders, the L5 PASS, `OfficerTab.lua` and the `L.off*` copy, `web/src/lib/guild/raid-copy.ts`, `web/src/lib/score/copy.ts`.

## Verdict: PASS

Every colour goes through Theme and tokens (no raw hex in the tab); the ringed class crest on every raider, shared with the Guild tab through `HubDraw`; two columns, a hairline between, no card chrome; section headings in the display face; class names, tier scores and quality items in their colours; client buttons only; game words, no em dashes; empty states leave sections out.

1. NICE. "12 of 40 coming" vs the web's "signed up": settled in L5 round 2; change `L.guildComing` and the Officer header together if reopened.
2. NICE. "Att." and "Ench." abbreviations: the 40 px column exception L5 accepted; keep "Att." the same in Roster and Master loot.
3. NICE. Keep the "Gap in DPS" legend visible so the number is not unitless.
4. NICE. Pell's 920 is the tier colour from the tokens (no raw colours in the code).
