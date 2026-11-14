# Wow-player re-gate, round 2 (spec text only, no renders)
Checked: each round-1 must-fix, then new game-fact errors against tables/1084.json and loot.json 1.60.1.70938.

## loot-plan: SHIP WITH FIXES
Resolved: Ashkandi gone (Vis'kag, Dal'Rend's Sacred Charge as the replaced piece), Drop became Skip, tier became phase, chance on every row,
non-gear listed as Not ranked, two-slot, shaman and survival rule 7, Discord lists only roles that can use the item.
New error: sections 6 fixture (Discord officer post, line ~133) and fixture text print `Ancient Cornerstone Grimoire (8%)`. loot.json
item_chances for Onyxia: 8% is Shard of the Scale (17064), Deathbringer (17068), Vis'kag (17075). Grimoire (17067) has no chance on file,
so it must read `Chance unknown` by the spec's own rule 8. Fix the sample, nothing else.
Check: "16 gear, 6 not ranked" matches 22 entries; confirm the 6 are Head x2, Hide Backpack, Scale, Sinew, Draconic for Dummies.

## live-pulls: SHIP WITH FIXES
Resolved: phases fixed (Phase 2 from 65%, Phase 3 from 40%; Fireball and Deep Breath are phase 2), downtime and rotation held, microsecond
budget, self-appearing frame, Name them officer channel only, Death Wish line gone.
New error: section 1 says `Pull 4 ... wipe at 31% · 3:12` with `Thandrel died first at 1:40`; the Discord sample (line ~134) has the same
pull 4, 3:12, with `Thandrel was first, at 2:41` and `2 to Deep Breath`. Same pull, two first-death times. Pick one. Also 6 deaths all in
phase 2 on a wipe that ended at 31% is plausible (air phase wipes) but then `Deaths by phase` for pull 4 shows nothing in phase 3; fine, say it.
Keep the coverage note; the Discord post still must not name raiders.

## prep-reminder: SHIP
Resolved: no `all slots enchanted` (enchantable slots only), attunement first, Resistance only where encounter data names a need (absent for
Onyxia), stale bags dropped, reagents and ammo stated out of v1 with verdict `All set on what we check`, DM default off, DM skipped when the
in-game check ran, Late raiders get their own time or nothing, send button disabled in combat, ready check, loot session, summon, mid-pull.
No new game-fact errors. Onyxia and Drakefire Amulet attunement wording is right.

## guild-later: SHIP WITH FIXES (comp, bench); recruiting SHIP
Resolved: Strength of Earth contradiction stated as an open check (K1 before K3) and the bench example no longer depends on it; missing
buffs and debuffs added to K1; `Copy whisper` replaces the invite button; spec filter is valid (directory carries spec); G1 privacy
sentence and rate limit stay gating.
Note: the 24-entry catalogue list in section 1 still omits Improved Scorch, Expose Armor, Shadow Weaving, Curse of Shadow, Mana Tide.
The fix says K1 adds them; confirm that line names them, else a raider asks "where is Expose Armor".
Nothing else wrong.
