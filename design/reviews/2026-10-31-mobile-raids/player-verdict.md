# Player verdict: mobile raid sign-up (F4), 2026-10-31

Reviewed: Android emulator captures (real icons), Roborazzi and iOS snapshots (empty icon tiles, offline), and the F4 report.

Question this answers: "Am I in for Saturday, and who else is?" The In / Late / Bench / Out row is on the raid card on the Raids tab, so signing up is one tap from the tab and no page push. That is faster than a Raid-Helper button in a Discord channel.

## Verdict: HOLD

## Blockers
1. Nothing real works without the API. The report lists what no handler sends yet: `invite_at`, `signups[].spec_name` and `spec_icon`, `spec_options`, `viewer.can` (lock/move), `tools_ended`, `raid_night`, `waiting_for_approval`, the whole `GET /v1/me/raid-nights/{raid}`, and the push `answers` payloads. Without them there are no spec icons in the roster, no invite time, no officer tools, no Your raid night and no notification answers. Every capture is a fixture. Do not call this the headline feature until the handlers send these and a live raid has been checked on a device.
2. iOS fixture mode is likely broken (report concern 9: fixtures land in a folder named `app`, so the bundled transport finds nothing). There is no live iOS capture at all, so nobody has seen the iOS roster with real icons. An iOS reviewer cannot sign off on blank tiles.
3. The officer move menu on iOS was never captured (SwiftUI Menu), and the notification Bench/Move paths were never driven. Officer tools are half of what a raid leader needs on raid night; unverified on iOS.
4. Notification answers need lane F7 (Android channel `raids` does not exist yet; iOS has no token registration). The "answer from the notification" promise does not work until then.
5. Roster read is not good enough to replace Raid-Helper for a leader. There is a class strip with counts (5 / 0 / 2 ...) but it has no class names, and a zero class looks like a greyed icon with no label. The Android capture shows the roster rows (class-coloured name, spec icon, spec name) but nothing tells me "2 of 3 tanks, 6 healers, need X". The page has no comp read by role against targets, and no raid buff gaps (no Mage, no Priest, no Paladin or Shaman warning) which is what a Classic leader checks first.

## Shoulds
- Times use a wide monospace stand-in font (JetBrains Mono not bundled): "6:45 pm" and "in 50 days 1 h" read huge. Also "Invites" appears twice (the times row and the countdown row), which reads like a typo. Show "Invites 6:45 pm, Pull 7:00 pm" once, then "Invites in 50 days 1 h" as a plain line.
- The countdown is to invites only; a raider wants "Pull" too (7:00 server time). Say which clock the times are in (the "guild time" line shows only when offsets differ; I would always say "server time" or the realm clock).
- Next-raid card has a dead gap under the In/Late/Bench/Out row (report concern 7), and the sign-up card has a big empty band between the segments and the note field. Tighten.
- "Approve them on foreversixty.gg" sends an officer out of the app for the thing they do most before raid night. Add approve in-app or say it is coming.
- Officer toast "Sign-ups locked. 23 In and 1 on the Bench were told." is good. The locked card text "Members cannot change theirs; officers can" is good. But Late is not in the toast count; say how many Late were told.
- Notification actions offer In / Late / Out only. Fine for raiders, but add Bench for signed-in officers or say why not.
- The mono type on iOS vs Android differs; check the roster rows at AX5.
- Raids tab shows "Barrow Deeps" as the 22 Dec raid. I do not know a Classic or Forever raid by that name. It is a fixture, but a fixture a player screenshots; use a real raid (Molten Core, Blackwing Lair, Zul'Gurub, Ruins of Ahn'Qiraj) in captures.
- Fixture note "Phase 2 prep: bring fire resistance" for Onyxia's Lair: Onyxia is a first-phase raid; use phase wording the client uses or drop "Phase 2".
- Your raid night: "Replaces Stoneskin Gargoyle Cape, ilvl 61": Classic 1.12 tooltips show no item level. Show the slot and the item name only, or confirm the Forever client shows ilvl. Also the coach note "the safe lane is along the side walls" for Deep Breath is not how Classic raiders describe it (you move out of the breath path or use Fire Resistance / the wall); have the accuracy reviewer check coach text. Premium-only coach text should not be in a free fixture capture as an example of the base feature.

## Data I do not believe (for the accuracy reviewer)
- "Dragon's Blood Cape" awarded by Onyxia kill at 61 ilvl, purple: verify it is an Onyxia drop in Forever (Onyxia's cloak is Onyxia Scale Cloak).
- Onyxia 4:28 kill with a fury warrior at 712.4 DPS is plausible but check the percentile source.
- "Barrow Deeps" as a raid name.
- The class strip shows nine classes, and the 0 class (rogue slot appears greyed). Check composition fixture is a plausible 40-man (3 tanks, 6 healers, 14 DPS out of 24 In, 40 capacity fine).

## What is right
- One-tap sign-up with optimistic save, "Saved." and rollback with the API's own sentence; Out never asks for spec; spec sheet only when no spec is saved, and "Close the sheet and nothing is saved." is clear.
- Spec lists are correct for Forever: Lightningbolltz offers Elemental / Enhancement / Restoration; no shaman dual wield claimed. Roster shows "Tidecaller · Restoration · Can swap to Elemental", which is how raid leaders actually think.
- Lock, Unlock, Copy the list, Move to Bench, Remove from the raid, and "Kraggor · In as Tank" in the menu header.
- Roster groups Tanks, Healers, DPS, Late, Bench, Out, Not signed up, with the player's own row highlighted.
- Spec icon plus class-coloured name per row (Android).

## Re-check (2026-10-31, after the fix round)

Verdict: HOLD (one blocker left)

Cleared:
- iOS fixture mode: iOS captures exist and show the full roster. Cleared.
- iOS officer move menu: `officer-move-menu-ios` shows "Kraggor · In as Tank", Move to Bench, Remove from the raid. Cleared.
- Notification answer: `notification-answer-in-ios` and `notification-answer-in-and` show "You are In for Onyxia's Lair". The action path works; push registration waits for F7 as ruled.
- Composition: class names with counts and a muted "0" for Paladin, "Tanks 3/4, Healers 6/8, DPS 14/28" against targets, and "Missing: Blessings". The counts match the roster rows (5 Warrior, 4 Priest, 4 Shaman, 3 Druid, 24 total with one Late). No Paladin gives a correct Blessings gap. The targets add up to 40.

Remaining blocker:
1. The buff and utility row is seven blank tiles on both platforms. In `raid-page-officer-ios`, `officer-move-menu-ios` and `raid-page-officer-and` you cannot tell Fortitude from Soulstone, and the muted Blessings tile is just a darker box. The only readable cue is the "Missing: Blessings" line. The report admits the icons do not load on the live iOS simulator, and the CDN export lacks `spell_nature_regeneration` and `spell_holy_fistofjustice`. There is no capture of this row with real icons on any platform: the Android emulator roster capture is scrolled past it. A leader checks this row first, so a row of empty boxes is a screenshot-and-mock item. Fix: capture the composition block on the Android emulator with real icons, get the iOS icons loading, and add the missing two icons to the export. Or give each tile a short text label (Fort, MotW, Int, Bless, ShadowProt, SS, BR) until icons load.

New Classic points (should, not blockers):
- Buff list is right for Classic (Fortitude, Mark of the Wild, Arcane Intellect, Blessings, Shadow Protection, Soulstone, Rebirth). It omits shaman totems (Windfury, Strength of Earth, Mana Spring), and Forever has shamans on both sides, so the "utility" row is incomplete. The fixture raid has 4 shamans and the row says nothing about them.
- Role targets of 4 tanks, 8 healers and 28 DPS for 40 are a sane Classic spread. They are fine as a fixture.
- Label "Battle Rez" is the brief's wording. The tile is Rebirth, so name it "Rebirth" (the Classic name) in the label.
