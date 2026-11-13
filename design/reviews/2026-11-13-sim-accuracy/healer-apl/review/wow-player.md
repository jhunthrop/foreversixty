# wow-player review, healer APL prose and cards, 2026-11-13

Verdicts: Paladin SHIP WITH FIXES, Shaman SHIP WITH FIXES, Holy priest SHIP WITH FIXES, Discipline priest SHIP WITH FIXES, Druid SHIP WITH FIXES. No HOLD. The behaviour matches the logged casts (paladin 91% Flash of Light, priest 60% Flash Heal, shaman 56% Chain Heal, druid 51% Rejuvenation). The fixes are contradictions inside the guides and a few wording slips.

## Paladin, Holy: SHIP WITH FIXES
True to play: Flash of Light as the raid staple is exactly what logged paladins do (22.6 casts a minute).
Fixes:
1. Prose: drop "and some of it will land on someone another healer already topped." Every raid healer knows overheal, and it reads as an apology. Keep "it is most of what you cast."
2. Prose: "a hit or two down" is fine, but "as soon as one hit lands" repeats the card. Suggest: "In a raid, Flash of Light goes on the tank after his first hit and on anyone in your party a hit or two down: it is most of what you cast."
3. The "Pace, not health bars" heading now sits over a paragraph that gives bars (97%, 93%). Rename it "Pace." or "Mana pace, then Flash of Light."
4. Card line 11 note "your staple, and in a raid the heal you cast most": ok. Card lines 13 to 16 have no note, so the party lines read thin. Fine to ship.

## Shaman, Restoration: SHIP WITH FIXES
Fixes:
1. The new sentence is under "Everyone else" but talks about the tank. Move it to "Tank first" or reword the paragraph: "Otherwise Chain Heal heals them. Cast it on the tank as soon as anyone near him is down a hit, not when they are low."
2. "a hit down" is not how anyone says it. Use "down a hit" or "even a little hurt".
3. "Tank first" says Riptide "never answers a big hit on its own", but the card now has Riptide on the tank below 95%. Add: "so keep it rolling on the tank".
4. Note, not a blocker: logged shamans put only 25% of Chain Heals on the main tank and 57% on others. "Cast on the tank" is still the right way to aim a chain, but "anyone near him" is the part that matters.

## Priest, Holy: SHIP WITH FIXES
Fixes:
1. "Renew goes back on the tank once he takes a hit": "goes back" is wrong for the first cast (Before the pull says Renew goes on as he engages). Use "Renew goes on the tank once he takes a hit and again when it falls off, at rank 7 or higher".
2. Contradiction: "It heals the least for its mana of anything you cast" is followed by Flash Heal as the top-off filler. Add "which is why it only tops off while your mana is ahead".
3. Contradiction with Mana: "Every heal you skip early is mana for the last minute." Change to "Every heal you skip once you are behind the fight is mana for the last minute."
4. "Flash Heal at the top rank in a raid" next to a top-off at 85% overheals. Logged priests mostly use r7 and lower on small holes. Say "top rank on the tank, a lower rank for the small holes".
5. Card note "On a full tank every tick is wasted." Over-explained, cut it.
6. Panels "while mana is on pace": "on pace" is not a game phrase and is not defined on the card. Use the existing shaman wording: "while mana % is over boss health %" (56 characters for the tank line, fits).

## Priest, Discipline: SHIP WITH FIXES
Fixes:
1. Contradiction in Fillers: "Greater Heal rank 1 on everyone else" versus Flash Heal on anyone under 85%. Write: "Greater Heal rank 1 on everyone else who is really hurt; the small holes are Flash Heal."
2. "Renew goes back on the tank": same fix as Holy.
3. The card claim "in a raid most of your casts are this one" is true of the logged priests (a Holy-heavy sample). Disc also has shield and Penance casts. Soften to "in a raid it is a large share of your casts".
4. Cut "On a full tank every tick is wasted."
5. Panel wording "mana is on pace": same fix as Holy.
6. Mana paragraph "stop topping off the party and heal only real holes" agrees with the new text. Good.

## Druid, Restoration: SHIP WITH FIXES
Fixes:
1. Contradiction: "Before the pull. Rejuvenation on the tank a few seconds before he pulls" versus "Rejuvenation goes on the tank once he takes a hit" and the card rule below 97%. Either keep the pre-pull Rejuvenation as an opener (say so in The tank: "after the opening Rejuvenation, Regrowth goes on once he is under 90%") or drop it from Before the pull.
2. "and from then on they stay up" conflicts with "recasts only below 90% / 97%" and with letting the HoT run out for the Bloom. Use "and you recast each only after it has run out".
3. The party Rejuvenation sentence is buried in Ranks. Move it to "The party": "With your mana level with the fight or ahead of it, rank 7 Rejuvenation goes on anyone in your party under 90%, you included."
4. "that Rejuvenation" is vague; name "rank 7".
5. Card line 15 note ends "The same rank as on a party member under 70%." This points at another line the player cannot see. Replace with "Rank 7 in a raid."
6. Tank Rejuvenation card note says "and he has taken a hit"; match the panel: "once he is below 97%".
7. Panels "mana on pace": same fix as the priests ("mana % over boss health %"); the yourself line is 59 characters now, so check the limit.

# Pass 2

Read the rotation sections of the five guides and the changed notes and panel lines in data/curated/apl/*.json. All panel lines I checked are within 60 characters (longest: "on yourself under 90% with no Rejuvenation while mana holds", 59). "while mana holds" is accepted: it is the phrase the other lists already use.

- Paladin, Holy: SHIP. Heading is "Pace.", the overheal apology is gone, "after his first hit" reads right.
- Shaman, Restoration: SHIP. Riptide is "keep it rolling on the tank", the Chain Heal sentence sits in "Tank first", "down a hit" is fixed.
- Priest, Holy: SHIP. Renew "goes on ... and again when it falls off", the Mana line is conditioned on being behind, and the "dearest heal, only tops off while you are ahead" line resolves the Flash Heal contradiction. Top-rank Flash Heal on the top-off is accepted because the list really casts it; the logged priests mostly use lower ranks (r7 40%), so a player may still downrank by hand. Not a blocker.
- Priest, Discipline: SHIP. Fillers now says Flash Heal for the small holes, and the "full tank every tick is wasted" line is gone. Note: the Discipline card says "a large share of your casts" and the Holy card says "most of your casts"; that is right for the logs (Holy 60%).
- Druid, Restoration: SHIP. The pre-pull Rejuvenation is removed so it no longer contradicts the new rule, "Recast each only after it has run out" replaces "stay up", the party Rejuvenation sits in "The party" with rank 7 named, and the card note ends "Rank 7 in a raid."
