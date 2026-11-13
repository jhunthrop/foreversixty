# L5 Guild tab: wow-player verdict

Question the tab answers: am I raiding tonight, am I in, what is my job, am I ready? Yes: Tonight, the sign-up row and the job are the first things on screen, prep sits right under them, and the teal "Late" plus the "Sends when the companion runs" footer is honest. The blue In/Late/Bench/Out buttons read like a raid-helper bot. The numbers add up (11+1+2+1+16 signed/unsigned; tanks 2+4... 12 coming, wanted 4+10+12+14 = 40). Dec 10, 13, 15 2026 are Thu, Sun, Tue. Item ids check out: Drakefire Amulet 16309, Onyxia Scale Cloak 15138, Greater Fire/Frost/Nature/Shadow/Arcane Protection 13457/13456/13458/13459/13461. MC and BWL boss lists are complete and correct.

## Verdict: FIX (small, no blocker)

1. Time zone. "Thu 10 Dec, 20:00" has no clock named; the approved board says "20:00 server" and the build dropped "server". Every raider asks "whose 20:00?". Restore "server" (or the local-time marker the Unified formats rule uses) in Tonight and in Next raids.
2. "Fire resistance 42" is a number with no target. Onyxia's breath: say what is wanted, or say nothing about the school. A raid leader wants "42" next to what counts as enough (guild call) or, failing that, the row should read as information ("Fire resistance 42, from gear and buffs") not as a check. As built it reads like a pass/fail it cannot judge.
3. Raid prep is thin for a raid night. Missing, as every Classic raider checks: Molten Core and Blackwing Lair attunement (MC "Attunement to the Core", BWL "Blackhand's Command"/Orb of Command) are not in the prep at all, only Onyxia's amulet is. A level-60 with no BWL attunement sees "Nothing missing" for Blackwing Lair. Either add those two door checks (the quest-done reader exists for 6502/6602) or say "Attunement not checked" for MC and BWL; never "Nothing missing" over an unchecked door. Also absent: durability/repair, hunter ammo, class reagents (Ankh, Symbol of Kings, Soulstone) . At least ammo for hunters and low-durability gear are what real raid-prep addons warn about; list these as follow-ups if not built now.
4. "Nothing missing" and the "Consumables" packed label: "Nothing missing" is fine, but with only "a flask or elixir, buff food" checked it overclaims. Use "Flask and food packed" or keep "Nothing missing" only under the heading "Consumables". Confirm the packed row prints under that heading (the render only shows the missing case).
5. The missing-potion row names the Greater Fire Protection Potion in red but is plain text: a name where the client shows an icon with a tooltip. Give the potion (and the amulet) the icon and item tooltip used elsewhere in the hub, and show how many you carry when you have some.
6. "Your job" is the owner's spec word and I will not fail it, but a raider says "assignment" or "role"; "1 of 2 jobs filled" reads strange. Cosmetic, owner call.
7. The old-companion render says "Sign up on the site" under "In as DPS, sent Tue 19:04": the answer is shown as sent and then tells you to sign up. Show "Change it on the site" when an answer exists, "Sign up on the site" only when none.
8. Guild line: "you are 9th of 15" while Sign-ups counts 31 raiders answered or not of 40: the ranked pack lists 15 and the raid roster 31+. Label the ranking "9th of 15 rated raiders" so the two numbers do not look like a bug.
9. Phase 2 job text "Fire on the whelps from the west wall when she lifts off" is fine for ranged, but hunters in Phase 2 stand and shoot Onyxia, not "fire on the whelps"; confirm the example copy is made-up and not shipped as default text.

## Data I do not fully believe (for the accuracy reviewer)
- Onyxia's door: in Classic the entry check is the Drakefire Amulet in the bags, so item-based is right; verify that Forever behaves the same, since attunement quests normally persist and the quest flag (6502/6602) is the safer second check.
- fire_protection kind: confirm the catalogue includes only the Greater potion (13457) or the row asks for the Greater while the lesser (Fire Protection Potion) also satisfies it; the label says "Greater" either way.
- Onyxia Scale Cloak for Nefarian: the Shadow Flame need is the community norm; confirm it matches Data.encounterNeeds [11583] rather than only the addon's own BOSS_NAMES.

## Round 2

Verdict: SHIP. Checked l5-raid-day, l5-old-companion and l5-queued.

- Time: "Thu 10 Dec, 20:00 local" is honest given the client clock; accepted over "server".
- Old companion: "Change it on the site" now sits under "In as DPS, sent Tue 19:04". Fixed.
- Guild line: "you are 9th of 15 raiders" with 3rd of all guilds and 41 raid nights. Fixed.
- Prep: item icon and tooltip on the amulet, "In your bags  Flask 2 · Food 20" row, red missing potion. Fixed.
- Roles are coloured against the target (Tanks 2 of 2 green, the short ones red). Good.
- Phase 2 job text confirmed example data, not shipped copy.
- MC and BWL left out of the prep rather than faked "Nothing missing": acceptable, since the prep is for Onyxia. Follow-ups, not blockers: MC/BWL attunement, durability, hunter ammo, resistance target (still a bare "Fire resistance 42" but it reads as information, not a check).
