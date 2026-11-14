# Wow-player gate: four guild specs (spec text only, no screenshots exist)
Reviewer: veteran Classic raider, loot council and raid leader. Judged: copy verbatim, game facts, fit with how a night runs
(invite, ready check, pull, wipe, 60 seconds of talking, loot). Nothing here is a render review. No spec is approved on visuals.
Verdicts
- loot-plan: SHIP WITH FIXES
- live-pulls: SHIP WITH FIXES (web, Discord, app); the in-game pull block (L5) is NOT SHIPPABLE as written
- prep-reminder: SHIP WITH FIXES
- guild-later: SHIP WITH FIXES (comp and bench); recruiting search NOT SHIPPABLE until the action copy and spec filter are fixed

## 1. Loot plan (2026-11-15-loot-plan.md)
The question is right and the answer is first. Officers will accept a "recommendation, the council decides" line because it is true.
Must fix
1. Section 5.1 and 6 fixture: `replaces Ashkandi`, `Ashkandi, Greatsword of the Brotherhood`. Ashkandi is a Nefarian drop and is not in
   Onyxia's table in loot.json (Onyxia's items there include Vis'kag the Bloodletter, Deathbone Guardian, Ancient Cornerstone Grimoire,
   Shard of the Scale). A Fury Warrior with no raid gear would not be wearing Ashkandi either. A raider screenshots this and the guild
   laughs. Use a real Onyxia drop and a real pre-raid item for `replaces`.
2. Sections 5.1, 12: `Drop` as the officer button next to the heading `Drops tonight`. In game "drop" is what the boss does. Rename the
   button (`Skip`, `Pass over`, `Remove`) and its toast `Dropped Brenn from Ashkandi.` and row `Dropped by Karl`.
3. Rule 3, 5.1, 12: `Won 2 this tier`. In Classic "tier" is the armor set (Tier 1, Tier 2); a raider reads this as set pieces. Say
   `this phase` or `since reset`, whichever the data actually is. The caption `Ties go to whoever has won least this tier` has the same fault.
4. Section 5.1 header `22 drops`, `Drops tonight`, raider `Your drops tonight`. Onyxia drops a few items per kill. loot.json shows 8 percent
   on Deathbone Guardian, Vis'kag and the Grimoire, 40 percent on the Sinew. The page promises 22 items. Show the drop chance beside each
   item (Wowhead does), call it `Possible drops` / `Could drop tonight`, and never print `22 drops` as a figure of what is coming.
5. The 22 includes non-gear (Head of Onyxia, Onyxia Hide Backpack, Draconic for Dummies, the Scale). They cannot be ranked. Say how they
   are filtered; a ranked list for a book would be mocked.
6. Candidate row has one `replaces X`. Rings, trinkets and weapons have two slots: show the one it would replace (the worse), and for a
   shaman (never dual wields in Forever) compare against the single weapon, never an off-hand sword. Survival hunters compare as melee.
   Add one line to rule 6 stating both.
7. Section 5.2 `Likely upgrade` for a raider on a 8 percent drop reads like a promise. Add the chance to the raider row.
Should fix
- Date format: `Wed 9 Dec · 20:00 server` here, `Wed 8:00 PM` in prep. One clock (the unified-formats rule).
- Discord raid post `Onyxia drops tonight: Ashkandi (Damage, Tanks)` lists tanks for a two-hander. Only list roles that can use the item.
- `Tanks`: `Upgrade` / `Small upgrade` is fine; no number is correct.
Game facts that hold: three role groups never mixed, recommendation wording, awards and attendance on the row, names to officers only,
directory gear dark by default.

## 2. Live pulls (2026-11-15-live-pulls.md)
The question ("what cost us that pull, and who do I tell") is what a raid leader asks. Cost lines in lost seconds is sensible.
Must fix
1. Section 1 sample: `Pull 4 · Onyxia · wipe at 31% · 3:12` / `6 deaths in phase 3, 4 to Fireball` / `died first at 2:41 to Deep Breath`.
   Deep Breath and Fireball are phase 2 (air) abilities; phase 3 starts at 40 percent. A wipe at 31 percent at 2:41 cannot have a
   Deep Breath first death. Phase 3 causes are Bellowing Roar fear into whelps and lava, Flame Breath, Knock Away. Fix the fixture in
   sections 1, 5.2 and 5.3 or a raid leader stops trusting the page on sight.
2. Section 2, cost 2 (downtime): active ms against fight length will flag every melee player on Onyxia, because she is untargetable
   in the air phase and everyone moves for Deep Breath. Downtime must exclude phase 2 and boss-untargetable windows (the engine knows
   phases) or the line is false for the whole raid. Section 5.2 `downtime, 5 raiders over 15 s` would be 40 raiders on Onyxia.
3. Section 2, cost 3: `Death Wish used once in 3:12, could be twice`. Death Wish is a 3 minute cooldown; the example is wrong, and
   naming a warrior's rotation to a raid leader between pulls is the kind of thing that starts arguments. Hold rotation lines (D7) until
   audited per spec; at launch ship deaths only.
4. Section 6 (addon): `under 0.3 ms per event average`. A 40 player fight throws thousands of combat log events a second;
   0.3 ms each is most of a core's frame time and would freeze the leader's client. The budget must be in microseconds (single digit),
   and the acceptance gate must be measured in a 40 man pull, not assumed. Until that holds, section 6 does not ship.
5. Section 6 block placement: "between pulls in the existing panel (the Raid tab area)". Between pulls the leader is looking at
   raid frames and talking. The block must appear by itself on ENCOUNTER_END as a small frame that fades, not hide behind a tab.
6. Section 5.3 `Name them` posting `Thandrel, Karl, Mira died to Fireball` to the guild Discord. Names after a wipe in a channel is the
   drama post. Officer channel only, or drop the checkbox in v1.
Should fix
- Phase 2 and 3 are the words a raider uses; show phase on each death line.
- Section 3: raiders see `a raider` for others. Fine. Keep the lock screen without names.
- One pull row per pull in Discord is off; correct, keep it off.
Game facts that hold: protection potion line suppressed until a COMBATANT_INFO sheet exists (never printed as missing), one companion
chosen for duplicate logs, no sim or parse numbers mixed into pull lines.

## 3. Prep reminder (2026-11-15-prep-reminder.md)
Question is right: am I ready, what do I fix in the next hour. Honest source and age stamps are the best part; keep them.
Must fix
1. Section 4.1 `Item level 61 · all slots enchanted`. Not every slot takes an enchant in Classic (neck, rings, belt, trinkets,
   shirt). `all slots enchanted` is false for every player. Say `Enchants on file: all that apply` or list only the slots that can take
   one (chest, bracers, gloves, boots, cloak, weapon, shoulders and head by the hand-in kinds).
2. Section 4.1 field `Resistance` : `Fire 21 from gear` and the web row order Enchants, Resistance, Flask or elixir, Food, Attunement.
   Onyxia does not need fire resistance from a raider; tanks may want a Fire Protection Potion, nobody checks gear resist for her. Make the
   Resistance row appear only when the raid's encounter data names a resistance need (Molten Core fire tank, later Nature/Shadow/Frost
   bosses), otherwise it is noise printed in a DM. Per class it is wrong too (a priest does not care).
3. Same order: put Attunement first. For Onyxia it is the one fix that stops you entering (Drakefire Amulet / the chain). It is the
   only row that can end a raider's night.
4. Section 4.1 Bags field `Last upload 4 days ago: 2 flasks, 1 food` with `Not a live check`. A four day old bag list is not
   information, it is clutter in a DM. Drop the Bags field when the upload is over 24 h and say only `Open the game, the addon checks
   at 7:00 PM`.
5. Missing from the check: class reagents and ammo (Hunter ammo, Shaman Ankhs, Paladin symbols, Priest candles, Rogue poisons,
   Warlock shards, Mage runes) and weapon stones and oils. A ready check that says `All set` for a hunter with no arrows will be
   mocked. State they are out of v1 and make the verdict read `All set on what we check` or add the rows.
6. Section 4.1/4.2 sends a DM, a phone push, an Activity row and an in-game toast for one fact. A raider already logged in with the
   addon gets the toast and a DM. Skip the DM and push when the in-game check ran in the last hour (the raider is plainly online).
7. Section 13.1 default on for DMs. Raiders hate bot DMs. Keep the setting prominent in the first DM and in Settings, and send the
   first DM only after the raider has linked Discord on purpose (already true), never to a Late sign-up an hour early.
Should fix
- `Onyxia in 1 hour` for a raider who signed `Late` is wrong; send at their own expected time or skip.
- Relay addon traffic: `relay!` on GUILD every 300 s and 76 `prep!`/`snapshot!` messages at T-1h. Invisible to chat, acceptable, but
  raiders in game at that hour are loading consumables; the jitter and throttle language is right, keep it.
- `Send tonight's prep` forces a reload at T-1h; correct to disable it in combat, ready check, loot session. Add: not while a
  summon is running or the raid is mid-pull, and the button must say it will reload (it does, in the tooltip).
- Ages: `31 of 38 checked` in the raid-facing line is good. Keep the Cannot check group neutral grey as written.

## 4. Guild later (2026-11-15-guild-later.md)
Comp and bench call
Must fix
1. Section 1 sample `Bench Brenn: lose Strength of Earth Totem (melee group 2)`. The spec itself says Strength of Earth is raid-wide
   in Forever; the example copy contradicts it. Fix the example, and check every totem and aura, not one: Windfury, Grace of Air, Mana
   Spring, Trueshot, Leader of the Pack, Moonkin share the same question (G2 names only Strength of Earth).
2. `Missing: Windfury Totem, Sunder Armor` is a good line, but the catalogue lacks the debuffs and buffs raiders actually count:
   Improved Scorch, Expose Armor, Shadow Weaving, Curse of Shadow, Judgement of the Crusader, Mana Tide, Prayer of Fortitude by name.
   Add or say they are out.
3. A bench call names a person. Keep it officer only (as written) and never post `Brenn sits` even in the officer post without
   the officer pressing send. One line is fine.
Ship K3 (coverage only, no DPS band) first; agreed with the spec.
Recruiting search
Must fix
1. Section 2 first screen offers `Invite on the site` as the only action, but the same section says no contact path exists and the
   action is a `/w` copy line. A raider reads that as a built-in invite. Make the button `Copy /w`.
2. Question `a 60 Fury Warrior`: the directory carries an equipment summary only, no talents. A spec filter cannot be right. Filter
   by class and level and item level, or say spec is a guess from gear (it would be wrong for a Fury warrior in tank gear).
3. Cold whispers to level 60s from a site are the recruitment spam players already hate. Keep the rate limit, the opt-out and the paid
   plan gate (G4), and do not ship before the privacy sentence (G1) is in.
No other game facts in either brief look wrong.

## Facts I could not confirm
- Whether Onyxia's Fireball and the air phase exactly match Forever's client; treat the phase list in live-pulls as needing the
  encounters.json read, not my memory.
- Forever Onyxia item chances beyond loot.json (8, 40, 100 percent) are as shown there.
