# Mobile apps (lane D): wow-player verdict

Reviewer: wow-player persona, 2026-10-30. Inputs: `design/specs/2026-10-30-mobile-apps.md`, every board in `design/mocks/2026-10-30-mobile-apps/renders/`, the lifecycle catalogue in `docs/superpowers/plans/2026-10-30-mobile-apps-lanes.md`, `docs/tenets.md`.

**Verdict: HOLD.** The raid flow is the best part. Signing up from the lock screen takes one tap, and so does signing up from the Raids tab. The roster is grouped by role with spec icons. Combined with "Classes coming", it beats a Raid-Helper post on a phone, and the Discord post updates from the app. Four things stop it shipping. Two boards show numbers a player would screenshot and mock. One state board teaches the wrong rollback. The catalogue is missing the one push that matters most on raid night.

## Answers to the brief

- **Sign-up vs a Discord button.** Faster from the lock screen, where a long-press and `In` replace opening Discord, finding the post and pressing the button. The Raids tab is just as fast. It is clearer on the raid page, which answers "am I in, as what, who else" in one place. `In / Late / Bench / Out` reads correctly. Raid-Helper people will look for "Tentative", and `Bench` covers it. The spec ask, `Which spec tonight?`, is exactly what a hybrid wants, asked once and kept. The roster reads like Raid-Helper done properly: role groups with counts, then Late, Bench and Out, notes under names, and `Not signed up 5` at the bottom.
- **Notifications.** `raid.posted`, `raid.signup_nudge`, `raid.changed`, `raid.roster_set`, `raid.starting`, `band.entered` and `sim.finished` are pushes a player wants at that moment. `loot.awarded` gets muted: you saw the loot drop in game, and the board delivers it a day later (`1d`). `ranking.best` per fight also gets muted, because a Molten Core night can send ten of them. The missing push is listed in blockers.
- **Gear, BiS and talents.** On par with the site: icons, quality borders, quality-coloured names, a runner-up on every BiS row, and DPS deltas with their unit. They share the site's known tooltip gaps (see shoulds). The talent board is wrong (blocker).

## Blockers

1. **`f6-talents-ios`, `f6-talents-and`: the point totals are wrong on a level-60 character.**
   - The tabs read `Elemental 30 · Enhancement 3 · Restoration 11`, which is 44 points. The build's own string in `data.json` (`5510331300001350-030000000000000000-5033520000000000`) spends 30/3/18 = 51. `Restoration 11` is hard-coded in `screens_more.py`.
   - F6.2 says `4 of the 51 points in the 60 build differ.`, but the drawn difference is 3 points: Reverberation 0→2 and Elemental Focus 1→0. The Elemental tab still reads 30 when it should read 31.
   - Any shaman will count the tabs, get 44 at level 60, and post it.
   - Fix: draw the totals from the string. Add an `{n} points unspent` line for the case where they really are under 51.
   - The tree also has no prerequisite arrows, which the talent frame and Wowhead both draw. Without them, `Elemental Mastery` and `Lava Burst` look free-floating.
2. **`f7-push-ios`, `f7-push-and` F7.3: `Loot: Talisman of Ascendance — Awarded to Lightningbolltz in Onyxia's Lair.`**
   - Talisman of Ascendance is not Onyxia loot. `loot.json` lists it as the Argent Dawn quest reward from `Epic Armaments of Battle - Friend of the Dawn` and `- Honored Amongst the Dawn` (quests 9222 and 9224).
   - A loot-council veteran screenshots that one immediately.
   - Use a real Onyxia drop, such as Onyxia Tooth Pendant or Mature Black Dragon Sinew.
   - Also make `loot.awarded` **off by default**, or fold it into `report.ready`. Arriving a day after you looted it in game, it is noise and gets the whole Raids channel muted.
3. **Catalogue gap: no push when your place on the roster changes after it is set.**
   - `raid.roster_set` fires once. The raid-night message that matters most comes later: an officer pulls you off Bench at 7:50 ("you're in, log on"), or moves you from In to Bench. Today that is a Discord @mention.
   - If the app is meant to replace Discord, that change must push to the affected player every time, as `time-sensitive` on iOS and on the high `Raid starting` channel on Android, and it must break quiet hours the way `raid.starting` does.
   - Add it to the plan's catalogue and to §6.2's copy: `You are In for {Title}` / `{when}. An officer called you in from Bench.` and the Bench mirror.
4. **`f4-states-ios`, `f4-states-and` F4.11 Save failed: the board shows `Late` still checked under `That did not save. Try again.`**
   - The spec says the press rolls back to the previous state (In).
   - On the flagship's error state, the board shows a raider who failed to save as Late. Builders follow boards.
   - Fix: draw `In` checked, with the wipe line under it.

## Shoulds

### Raids

- **Officer tools in v1 (§17 Q1).** On a phone, Raid-Helper lets officers move people and close sign-ups from Discord. If officers must go to the web to lock or bench someone on raid night, guilds keep Raid-Helper and the app becomes a second door. Pull `lock` and `move to In / Bench` into v1. Planning can stay on the web.
- **Invite time.** Raids start at the pull, but players log on for invites. `Invites at 7:45 pm` lives only in the notes on F4.6. Make the invite time a field on the raid. Anchor `raid.starting` to it when set, and show it on the raid page under the time line.
- **iOS lock screen (F7.1, F7.3).** The guild subtitle the spec promises in §6.2 is missing. Android has it (`Olympus XXVII`). It matters for anyone with alts in two guilds.
- **`raid.starting` title.** `Onyxia's Lair starts in 60 min` wraps badly on iOS. Use `in 1 hour`, and `in 30 min` below an hour.
- **F7.3 Android `Pick a spec for Barrow Deeps`.** The example uses Lightningbolltz, who every other board shows with a saved Elemental spec in Olympus XXVII (`Change spec` is drawn). Use an alt with no saved spec. Also draw the `Out` action that §6.1 promises next to `Choose a spec`.
- **F4.12 Moved.** The other character is Grimtusk, a level-41 warrior, signed up for Onyxia. Use a level-60 alt. Draw both character blocks, with `Grimtusk moved to Bench.` on Grimtusk's own block, as §5.4 F4.3 says.
- **Raids tab.** `Planned 3` lists two rows, because the next raid is counted but drawn above. Count only the rows listed, or say `3 planned` in the head.
- **F4.6 notes copy.** `Bring Fire Resistance gear for phase 3` should say phase 2. Fire resistance on Onyxia is for Deep Breath in phase 2; phase 3 is fears and whelps.
- **Soft reserves (v1.x).** On raid night most Classic guilds want a soft-reserve link (softres.it or similar) as a raid field next to the notes.

### Notifications

- **Batch the after-raid pushes.** Fold `ranking.best` into `report.ready` (`Your log from Molten Core is ready · 2 new personal bests`). Otherwise a raid night sends three to ten pushes after the last boss.
- **`band.entered`.** Name the first upgrade in the body (`First upgrade: {item} from {source}`). That is the line that sends a levelling player somewhere.
- **`game.milestone`.** Keep the phase and raid-unlock dates. Drop `Forever launches today` for accounts that already have a synced character; they know.

### Gear, BiS and the tooltip

- **F2.6 tooltip.** These are site parity gaps, and the client tooltip has all of them:
  - no armor material (Cloth, Leather, Mail);
  - no `Binds when picked up` or `Binds when equipped` line;
  - no `Unique`;
  - `+10 Hit` and `+25 Spell Power` instead of the client's wording;
  - the item level repeated in the header and the panel.

  A 56-armor epic waist with no material shown leaves a shaman asking "is this cloth?". Have the data lane publish class and subclass ids, and keep the gap visible as a follow-up.
- **Reputation sources.**
  - `Vendor: Meilosh` (Knowledge of the Timbermaw, Rugged Mantle of the Timbermaw) needs the faction and standing: `Timbermaw Hold, Revered · Meilosh, Felwood`. A player cannot act on a bare vendor name.
  - `Warsong Outriders` on Dryad's Wrist Bindings needs its standing too.
- **F2.7 Your gear.** The pick is drawn only as an icon. At 36 pt, Warlord's Mail Helm (worn) and the Head pick look alike. Name the pick under its icon, or show `→ {pick}` in meta.
- **F3.1.** The faction chip is clipped to `Hor` (tenet 4). Wrap the chips or shorten them to the mark plus `H` or `A`.
- **F3.2.** `Back to Lightningbolltz's list` shows while the list already is his. The spec says it shows only when the picks differ.
- **F5.3 `The gear it ran`.** 17 icons with no slot order cue. Keep character-frame order and add a long-press slot label, or draw two rows split left and right like the character frame.

### Phone uses v1 lacks

- Item lookup beyond the 17 rows. "What is this green I just looted, and is it better?" is the most common phone use while levelling. BiS search filters only the current list.
- Quest sources with the giver and zone, so a player knows where to go.
- Raid night: soft reserves and the invite time, covered above.

## Data I do not believe (for the accuracy reviewer)

The `60 · Phase 1 raid` Elemental list and home answer pick or offer several items whose Classic sources open after Phase 1. Check each against Forever's own source opening; the item names come from the boards.

| Item | Classic source | Opened in Classic |
|---|---|---|
| Chains of the Lich | Balzaphon in Stratholme, a Scourge Invasion summon | 1.11 |
| Elemental Focus Band | Prince Skaldrenox, Abyssal Council in Silithus | 1.8 |
| Dryad's Wrist Bindings | Warsong Outriders, Warsong Gulch reputation | Not at launch in Era |
| Talisman of Ascendance (the trinket in F5.3's gear grid) | Argent Dawn `Epic Armaments` quests | Naxx era |

`Blue Dragonscale Helm` and `Arcanoweave Cloak` are not Classic items I know. If they are Forever's own, they are fine.

## Already right

- One tap on the lock screen, with the answer posted back in place.
- The spec is asked once per guild.
- R3, the moved-to-Bench line, is shown rather than hidden.
- Live refetch in the last 24 h.
- Calendar export.
- No queued writes offline, with a sentence that says why.
- The reminder ask comes in context, after the first In, never at launch.
- No Premium pushes.
- `Free` with no nag outside link regions.
