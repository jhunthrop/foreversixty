# Wave 3 mobile design: wow-player verdict (after the HOLD)

Reviewer: the wow-player agent run headless on the rendered boards (`design/mocks/2026-10-31-mobile-wave3/renders/`), four rounds, 2026-10-31. The first HOLD (the MUSTs on the Elemental tier hero, the rotation, the Binds line and the raid history) is quoted in the fix commit's message; this file holds the re-reviews.

**Final verdict: SHIP.** No MUSTs and no design defect remain on the boards. What travels with the lanes, written into the spec as ship gates:
- F9: `tiers.json` does not ship until the accuracy reviewer has verified every DPS band (Elemental 20th, 51.0% behind Fury; Frost Mage 44.1%; Balance Druid 47.8%).
- F12 and F2.6: the Binds line prints from the client's data; item 11623 is in no ItemSparse row of build 70658.
- F9.3: Mana Spring Totem's duration (client data says 300 s, Era is 60 s) stays unprinted until settled.
- The Back slot reads two builds (F12: 70658, Home and widgets: 70532); the lanes read one build; no store screenshot shows both.
- "+14 Spell Power" against the client's own Equip wording, and the first-raid names on F9.5, are for the accuracy reviewer.

---

# Round 1

Reviewer: wow-player, 2026-10-31. This follows my HOLD. I read the iOS boards F9, F10, F11, F12, F14, F15 and F16, and spot-checked Android F9 and F16. The renders are small, so a few strings are read at low resolution; I say so where it matters.

The structure still works. Each surface answers its question first: Your guide, Your rank, Your parses, Your Back pick and Next raid. Every number I checked has a unit.

## Earlier MUSTs

| Earlier MUST | Now |
|---|---|
| Tier list hero was Elemental, dead last at "51.0% behind Fury Warrior." | **Fixed.** The sample is a Fire Mage, "Fire Mage is 6th of 20 DPS specs. 9.1% behind Fury Warrior." Fury, Destruction Warlock, Affliction, Beast Mastery and Marksmanship sit above it, and that reads as believable. The caption keeps the Elemental accuracy gate. |
| Rotation had no Flame Shock, and Mana Spring Totem said "five minutes". | **Fixed.** Flame Shock Rank 6 is in the list, marked "Not in this sim's loop. In our sim it costs more mana than it adds." Mana Spring Totem now says "Down first. It returns mana for the rest of the fight." with no duration. Elemental Mastery is in the talent list (Android shows it at the bottom). |
| Spritecaster Cape had no bind line and no armor type. | **Fixed on the board.** It now reads "Binds when equipped", then "Back" with "Cloth" right-aligned, as the client prints it. "Binds when equipped" is the mock's word; see the gate below. |
| Raid history had five-man dungeons as 38 raiders, every night 3 h 12 m, and dates before the raid opening. | **Fixed.** BRD and Lower Blackrock Spire are "5 raiders". Lengths vary (1 h 47 m, 1 h 38 m, 2 h 26 m, 1 h 12 m). The Onyxia nights are 10 and 17 Dec, after the 9 Dec opening. The weekdays are correct for 2026. |

## F9 Guides, tier lists, what changed

The player's question is "How do I play my spec, and where does it stand?" Your guide is the first card on the Guides tab, so the answer is first.

MUST
- None.

SHOULD
- **Strengths versus the loop.** The overview says "Lava Burst hits harder with Flame Shock already on the target". The rotation then says Flame Shock is not in the loop. The reason is printed, but a raider reading both will argue with it. Add one line under Flame Shock, such as "Lava Burst still counts it if you apply it", or drop the strength line.
- **Chain Lightning.** "Chain Lightning Rank 4 · Two or more targets." is still there. If that is the APL's own wording, fine. A single-target raider casts it on cooldown too, so say so.
- **What changed.** Every row is a "SITE" pill (log coach, plans, sign-up). A player still has no game hotfixes or talent changes in the list. One row reads "Premium, Guild and Raid guild: the launch plans and founder prices." That reads as an ad inside a helper app. Cut it or move it to the You tab.
- **Talent page.** It lists key talents but never shows the build's point split or a "Build in the planner" action. The planner link only peeks out at the top of the rotation page.
- **Date style.** "Sept 17" sits beside "Oct 27" and "Nov 4". Use one abbreviation style.
- **Guide date.** On iOS F9.2 the header reads "Updated 24 Sep 2025" at the size I could read, while F9.1 says "24 Sep 2026". Confirm the year.
- **Sample player.** Emberquill the Fire Mage replaces Lightningbolltz in the header only on the tier board. The caption says so. Keep it.

## F10 Your reports

The question is "How did I do last raid?" The answer is first: Onyxia, 398 DPS, 93 percentile.

MUST
- None.

SHOULD
- Resolved: 398 DPS, 14th of 212 and the 93rd percentile agree across F10.2, F10.3 and F11.1. The 412 DPS on Home is the sim, and the board says so.
- Resolved: the report dates are 10 Dec and later, and the roster names (Draknar, Skarn, Zekra) match the meter.
- **Kill players versus raiders.** F10.3 and F10.2 show 38 players on the 10 Dec kill. F16.3 shows 36 raiders for the whole 10 Dec night. A night cannot have fewer raiders than its kill. Make the night 38 or more.
- Meters still have no active time or percent of total, the way Details! does. This is optional on a phone.
- **Coach wording.** "Mana Spring Totem was down for 41 seconds of the kill" is ambiguous. Say "not up for 41 s of the 3:12 kill".

## F11 Rankings and characters

The question is "Where do I stand, and what does that player wear?" "14TH of 212 Elemental Shaman characters" with "Jump to your row, 14th" answers it first.

MUST
- None.

SHOULD
- **Ordinals.** "14TH" and "204TH" are in capitals. A player will mock it. Print "14th" and "204th".
- **Bare "40".** The row line "Hordebreakers · 40 · 3:12 · 1 d ago" has a number with no unit. Write "40-man" or "40 players".
- **Date before opening.** Voltaris's "Barrow Deeps · Normal · 9 d ago" on 17 Dec is 8 Dec. The first raids open on 9 Dec. Make it 8 d or less.
- The alt story works: Sparkwick is 204th at 210 DPS and the card names the character in both states. I agree the pinned row and the first row can differ.
- **Back-slot contradiction.** F12.4 calls Spritecaster Cape "Your Back pick" for Lightningbolltz. F10.8 and F14.1 say his next Back upgrade is Arcanoweave Cloak, +3.1 DPS, and F15 repeats the Home upgrade. A player will see one slot with two answers. Production data comes from one build, so make the sample agree before anyone screenshots it.

## F12 Search and the item page

The question is "What is this item and where do I get it?" The tooltip and the "Your Back pick" card come first.

MUST
- None. The bind line and Cloth are present.

SHOULD
- **Bind word.** "Binds when equipped" is a mock word. The item is in no `ItemSparse` row of our build. Before the build, the word has to come from data. A boss drop from Blackrock Depths that is BoE would also be unusual, so the accuracy reviewer should check it. If it is unknown, print nothing rather than a guess.
- **Vendor lines.** The prices are now printed (8g 63s, 12g 92s, 8g 23s). No standing is shown for Meilosh or Pix Xizzix, while the Warsong line does show "(Exalted)". An ilvl 65 vendor belt for 8g 63s reads as open to anyone. Where the standing is unknown, say "Standing not listed" so it looks deliberate.
- **Drop chance.** "Drop chance not known" is now printed. Good.
- **Spell Power wording.** "+14 Spell Power" is white with the other stats. If the client prints it differently, match the client. Accuracy to confirm.

## F14 Widgets and live surfaces

The question is "Am I in tonight, and what do I do about it?"

MUST
- None.

SHOULD
- Resolved: the lock screen now agrees with itself. The clock reads 7:07, the countdown 38:00 to invites at 7:45 pm and the compact island 38m. The maths holds.
- **Unitless rings.** The circular lock widgets read "38" and "53m". Which one is the signed-up count, and which is minutes? A bare "38" has no unit. The upgrade ring reads "53m" in the render (low resolution), which looks like a mangled "+5.3 DPS". Put the unit on both.
- **Timeline.** The caption puts all of F14 on the evening of 10 Dec. The home-screen and state widgets show "in 2 days 2 h" and "24 of 40 signed up", and the lock screen shows "38 of 40". The "No raid planned" widget says "Last: Onyxia's Lair, Thu 17 Dec." That date is a week after the evening in question. Either date the widgets to the earlier day, or change "Last" to a past raid.
- **Server time.** Still no guild server time on the medium widget. A raid leader says "8 server". Keep it as a SHOULD.

## F15 System integration

MUST
- None.

SHOULD
- The Siri upgrade answer is good: icon, purple quality colour, "Waist · item level 65 · Vendor: Meilosh · 8g 63s", "+5.3 DPS for Lightningbolltz. Replaces Belt of the Archmage." It matches F10.8 and the widget. The standing is not printed, as the spec decided.
- Offline sizes add up: 2.4 + 3.1 + 0.6 + 0.4 + 41 comes to about 47 MB.
- F15.5 prints "(v1 strip, unchanged)" in a heading. That is a build note. Take it off the board.

## F16 Guild extras

The question is "Who is in my guild, and how did our nights go?"

MUST
- None.

SHOULD
- Resolved: the roster now says "Protection Warrior", "Holy Priest", "Fury Warrior", "Elemental Shaman", "Restoration Shaman", "Combat Rogue" and "Affliction Warlock". The raid history is believable.
- **Same date, two roles.** F16.3 lists Thu 10 Dec as a finished night with 2 wipes. F16.7 on the same board shows Thu 10 Dec as the next raid, "Invites in 2 days 2 h". The spec explains two timelines, but one board reads as a contradiction. Move F16.4 to F16.7 to a Thursday after 17 Dec, or caption the board with its date.
- **Class chips.** Tanks, Healers and DPS have no class chips. This is still open and optional.
- **Raiders in a dungeon.** "5 raiders" on a five-man run reads oddly. Use "5 players".
- Android F16 matches iOS.

## Data I do not believe (for the accuracy reviewer)

- "First raids: Barrow Deeps · Hyjal Summit · Onyxia" in `dates.json`. I do not know the first two as Classic raids. If they are Forever's own, fine. Confirm the names and that they are Phase 1.
- Paperdoll: "Signet Ring of the Bronze Dragonflight", ilvl 80, on a character in the first week. Is it obtainable in Phase 1? Also check "Serenity Field" and "Ironfeather Leggings" for existence and level.
- "Lava Burst Rank 4" in the rotation, next to a one-rank Lava Burst talent. Check the rank against `spellranks.json`.
- Mana Spring Totem lasts 300 s in the client, which is a Forever fact. I accept it. The board prints no duration, which is right.
- Stat priority "Intellect, Spell power, Hit, Nature power, Critical strike, Spell haste, Spell penetration". Hit below Spell power is unusual for an Elemental shaman. Nature power is rare as a stat. Check the sim.
- The tier board's Fire Mage at 9.1% below Fury sits level with Marksmanship Hunter. Fine as a tie, but confirm the order.

**Verdict:** SHIP WITH FIXES

Remaining MUSTs: none. The fixes before build are the Back-slot contradiction (F12.4 against F10.8/F14), the bare "38" and "53m" rings, the 36 against 38 raiders, "14TH" and "204TH", the 9 d ago before the raid opening, the "Last: Thu 17 Dec" widget date, and the Flame Shock reconciliation line. The accuracy gates are the Elemental band file, the Binds word, the first-raid names and the ilvl 80 ring.

---

# Round 2


I read all seven iOS boards (F9, F10, F11, F12, F14, F15, F16) at 1.5x. I did not open the Android renders, so the spot-check of `-and` is not done. Strings are small, so I mark the two I read at low resolution.

## Round-1 fixes, checked on the renders

| Fix | On the board |
|---|---|
| Ordinals | **Fixed.** "14th" and "204th" print in lower case on F11.1 and F11.3. |
| Rank rows | **Fixed.** They read "Hordebreakers · 40 players · 3:12 · 1 d ago". |
| Barrow Deeps | **Fixed.** F11.4 says "Normal · 8 d ago". |
| F16.3 night size | **Fixed.** Thu 10 Dec is 39 raiders, and the kill on F10 is 38 players. |
| Dungeons | **Fixed.** Lower Blackrock Spire and Blackrock Depths read "5 players". |
| Board day | **Fixed.** The F16 caption says F16.3 is Thu 17 Dec and F16.4 to F16.7 are Tue 8 Dec. F16.7 "Invites in 2 days 2 h" now agrees with that. |
| Lock-screen rings | **Fixed.** They read "38 of 40" and "53 min". |
| No raid planned widget | **Fixed.** It reads "Last: Lower Blackrock Spire, Sun 6 Dec." |
| Flame Shock | **Fixed.** The row says "Not in this sim's loop… it costs more mana than it adds, even counting Lava Burst's 20% bonus." That reconciles with the Strengths line. |
| Coach wording | **Fixed.** "Mana Spring Totem was not up for 41 seconds of the 3:12 kill." |
| F15.5 build note | **Fixed.** It is gone. |

## What is still wrong

No MUSTs. These are small and quick to fix.

1. **F14 has two days on one board.**
   - F14.1 and F14.2 say "in 2 days 2 h", which is Tuesday.
   - F14.3 says "Thursday 10 December", a 7:07 clock and "Tonight 8 pm". The lock-screen countdown also assumes Thursday.
   - The F14 caption does not name a day. Caption each phone with its day, or make the F14.1 and F14.2 widgets read "Tonight 8:00 pm" to match F14.3.
   - "24 of 40" on the home screen against "38 of 40" on the lock screen is fine across a day and a half, but only once the days are named.
2. **"Last:" on the No raid planned widget is a five-man dungeon.** A "Next raid" widget that says "Last: Lower Blackrock Spire" is odd. Pick a raid night from the history, or say "Last night out".
3. **The Back slot still has two answers on screen, by decision.**
   - F12.4 shows Spritecaster Cape as "Your Back pick" with a "YOURS" badge.
   - F10.8, F14.1 and the Home widget say the next Back upgrade is Arcanoweave Cloak, +3.1 DPS.
   - I accept that this comes from two builds. The F12.4 caption does carry "Build 70658". A player screenshotting the pair will still ask which cape. If the upgrade rows can be regenerated from 70658 before anyone captures them, do it. Otherwise leave the pair out of marketing and store screenshots.
4. **F9.2 guide date, low-res read.** The header looks like "Updated 24 Sep 2025", while F9.1 says "24 Sep 2026". Confirm the year.
5. **F9.5 date style and ad row.**
   - "Sept 17" still sits beside "Oct 27" and "Nov 4".
   - "Premium, Guild and Raid guild: the launch plans and founder prices." is still an ad inside a helper app. I raised both in round 1.
6. **F11.4 score date, low-res read.** "As of 4 Nov" sits on a board dated 17 Dec, before the first raid on 9 Dec. Use a date after the 10 Dec kill.

## Data still gated (unchanged by decision)

These go to the accuracy reviewer.
- "Binds when equipped" on the cape. If the word is unknown, print no bind line.
- The Elemental band, which puts it last of 20 at 51%.
- "+14 Spell Power" against the client's "Equip: Increases damage and healing done by magical spells and effects by up to 14." wording.
- "Signet Ring of the Bronze Dragonflight", ilvl 80, on F11.5 for a first-week character.
- "Barrow Deeps · Hyjal Summit · Onyxia" as the first raids.
- Vendor standing for Meilosh and Pix Xizzix is not printed. That is acceptable, but "Standing not listed" would look deliberate.

## What I liked

Each board puts its answer first: your guide, your rank, your parses, your Back pick and your next raid. Every number I checked now carries its unit. The meters, roster, history and Siri answer are things a raider would use.

**Verdict:** SHIP WITH FIXES. Fix the F14 day labels, the "Last:" raid wording, the F9.2 year and the F11.4 "As of 4 Nov" date before build. Do the F9.5 cleanup in the same pass. Clear the accuracy gates before launch.

---

# Round 3


I read F14 on iOS and Android, F11 and F9 on iOS, and F16 on Android. The F9 and F11 strings are small, so I mark the ones I read at low resolution.

## Round-2 fixes, checked on the renders

| Fix | On the board |
|---|---|
| F14 is one day | **Fixed on iOS.** The caption reads "7:07 pm on raid day". Home, widget states and the Dynamic Island all say "Tonight 8:00 pm", "in 53 min" and "38 of 40". The lock screen says "7:07" and shows the "53 min" ring. |
| No raid planned widget | **Fixed.** It reads "Last night out: Lower Blackrock Spire, Sun 6 Dec." The Android board matches. |
| Forever Score date | **Fixed.** F11.4 reads "As of 17 Dec", after the 10 Dec kill. |
| F16 | **Consistent.** F16.3 is the Thu 17 Dec history, and F16.7 reads "Invites in 2 days 2 h" on its Tue 8 Dec board. |

## What is still wrong

No MUSTs. Three small items:

1. **The F14 lock screen and shade still say Thursday 10 December while the Home board says "Tonight".**
   - Thu 10 Dec was the first Onyxia kill, and F16.3 and F16.7 put the Onyxia event on that same Thursday. So "Thursday 10 December" on the F14 lock screen matches "Tonight 8:00 pm" if raid day is the 10th. That works.
   - The conflict is F16, where the board day is Tue 8 Dec but the Onyxia event is 10 Dec. That is fine, because it reads "in 2 days 2 h".
   - I am only noting it because the F11 and F16 captions use 17 Dec as the "today" date, which is a different day again. Each board names its own day, so nobody will be confused.
2. **F14.3 on Android says "Invites for Onyxia's Lair in 38 minutes".** The widgets on the same board say "in 53 min". Invites are at 7:45 pm and the raid is 8:00, so 7:07 pm gives 38 minutes to invites and 53 to the pull. Both are right. A player may still read it as a mismatch, so name the event in the notification ("Invites in 38 minutes, pull in 53").
3. **F9.2 header year.** At zoom it reads "Updated 24 Sep 2025", while F9.1 says "24 Sep 2026". You decided on 2026 and the guide file says 2026-09-24. If the build actually renders 2025, that is a bug, and I read it at low resolution. The accuracy reviewer should check the string in the build, not in the mock.

## Decided and unchanged

- The F9.5 date style ("Sept 17" beside "Oct 27") and the changelog rows, including "Premium, Guild and Raid guild: the launch plans and founder prices.", are site data. I raised the ad row twice. It is still an ad inside a helper app, but I accept your decision.
- The Back slot shows Spritecaster Cape on F12 and Arcanoweave Cloak on Home and F14. I still say to keep that pair out of store screenshots. F11.5 also shows "Arcanoweave Cloak" in the Back slot of Voltaris's gear, which makes the mismatch more visible on a single board.

## Data still gated (the accuracy reviewer owns these)

- The "Binds" word on the cape. If it is unknown, print no bind line.
- The Elemental band on F9.4, which puts it last of 20. The caption already says "ships only after its band is verified."
- "+14 Spell Power" against the client's "Equip: Increases damage and healing done by magical spells and effects by up to 14."
- "Signet Ring of the Bronze Dragonflight", ilvl 80, for a first-week character on F11.5.
- "Barrow Deeps · Hyjal Summit · Onyxia" as the first raids. F9.5 shows "Dec 9 First raids".

## What I liked

Every board puts its answer first. The widget states are complete. A raider would use the 7:07 lock screen as it stands.

**Verdict:** SHIP WITH FIXES. The design is shippable. Name "pull in 53 min" in the Android invite notification, confirm the F9.2 year in the build, and clear the accuracy gates before launch.

---

# Round 4

## Round 4 verdict

**F14 Android (confirmed).** The notification now reads "Pull at 8:00 pm, in 53 minutes. You are In as DPS on Lightningbollttz. 38 of 40 signed up." The title says "Invites for Onyxia's Lair in 38 minutes". Invites are before pull, so 38 minutes against 53 is consistent. The widgets, the "Next raid Thu 8 pm · In" tile and the server-sim progress and finished rows all agree with each other. No defect on this board.

**F9.2 header (not confirmed at full size).** I could not crop the PNG here, because shell access was denied. I also could not open the HTML at the path I tried. The board only reaches me downscaled from 5361 px to 2000 px. At that size the header line under ELEMENTAL SHAMAN reads "Updated 24 Sep 2025 · Recommended races: Dwarf, Orc". The last digit is only a few pixels tall and could be a 6, so this is a resolution read, not a finding. Your grep count of 0 for "2025" in the HTML and the guide file's date of 2026-09-24 outweigh it. I'm withdrawing the item on that evidence, not on a full-size look. The F9.1 card on the same board reads "Updated 24 Sep 2026", which agrees with you.

The remaining items from round 3 are data gates and build-time gates owned by the lanes and the accuracy reviewer. They are the Mock section and F9.4 and F12.4 of the spec, including the Elemental tier-list band that ships only after verification. No design defect is left on the boards.

**Verdict:** SHIP
