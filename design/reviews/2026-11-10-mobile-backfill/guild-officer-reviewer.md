# guild-officer-reviewer: mobile backfill (Raid tab, both platforms)

## Round 1 (2026-11-10, base ac2f37dec)

# Guild-persona review, round 1: Raid tab (iOS + Android), Home raid card, raid push

**Page job:** answer "when is raid, am I in, are we full, am I ready" in two taps for a raider. For an officer, answer "who is short and who hasn't answered" and let them act on it from a phone.

**What I judged from:** I read the Android parity PNGs and `.txt` dumps, the Android flow steps, the iOS simulator and golden captures, and the push fixtures. I read the code only where the captures left a question open. I could not run either app. I did not exercise the web mocks beyond reading the README. The non-officer iOS captures show an empty "Guild sims" row, which I treat as the fixture limit. I have not re-reported anything in the BRIEF's fix-in-progress list.

## Jobs, step by step

**Raider jobs**

| Job | Result |
|---|---|
| When is the raid, am I signed up | Pass. The Raid tab head shows the raid, the pull countdown, "As Thandrel, Damage" and a 4-way control, all at once. The Home card adds a big countdown and role bars. |
| Change my answer | Good. Late opens 15, 30 and 60 min chips in place. The Android flow shows the saving state, the failed state and an offline banner. |
| Back after changing | Android `raid-link-02` to `-03` returns to the tab with the answer kept. |
| Before you pull | The best raider feature in the app. The rows are attunement, Elixir "3 in your bags", Flask "None in your bags · Auction house", and "Boots not enchanted · Where to get it", with a "from your last log-out, 9 h ago" stamp. It is honest about staleness and every miss has a next step. |
| Raid-night live | Raid night page, ongoing notification and green dot all exist. |
| Push from the lock screen | `push.raid-posted.apns.json` carries In/Late/Bench/Out actions. iOS `notif-03` shows a "How late?" follow-up with 15/30/60. This is the strongest raider flow. |

**Officer jobs**

| Job | Result |
|---|---|
| Who is short | The tab shows "Healers 5 of 8 · 3 short" and Home shows a 5/8 bar. |
| Chase non-answerers | "5 have not answered, 2 of them healers" leads to a Not answered screen. It lists healers first. Each row has Mark In, Bench and Discord DM, and Mark In confirms with Undo. A role mention comes with a preview, and the bench list is under it ("If Mirethil, Selka and Aurelith are In, healers are full"). That is real spreadsheet-replacing work and it is a phone-native workflow. |
| Applicants | The applicant screen shows Forever Score, sim HPS, item level, enchants "6 of 8 · cloak and boots missing", BiS gap, guild history and their note. Accept and Decline are on the same screen. This is strong. |
| Readiness of the roster | Missing. See finding 1. |
| Loot, attendance, performance, assignments | Not on the mobile Raid tab. Only "Loot history · Last raid: 4 pieces" and "Guild sims #7 of 14" appear, as link rows. |

## Findings

### BLOCKER

1. **An officer cannot see raid readiness, only their own.** "Before you pull" is the viewer's own flask, attunement and enchant check. An officer with 5 unanswered and 3 healers short also wants "6 raiders have no flask or enchants". Neither platform shows this, and the web mock has "Before the pull" as a raid-wide concept. The data to build it exists. The guild home roster rows carry `addon`, `logged_at`, `logged_recently` and `sim_export` (openapi `GuildHome.roster[]`). The raid view's `signups` rows already carry names and state. `RecentAttendance.slipping` is also on the guild home. Add a single "Not ready" line under Who is coming for officers, such as "4 signed In have not synced this week". If per-raider consumable and enchant state is not in the roster payload, label it **needs data**, but the sync staleness line alone is already actionable. Files: `RaidTabScreen.kt` near line 170 and `RaidRootView.swift` near line 120.

2. **There is no officer push at all, and the officer problem is a time problem.** `guild.raid-badge.officer.json` already computes `unanswered: 5` and `applicants: 1`. The only push fixtures are `raid-posted` (to raiders) and `raid-starting` ("Starts in 60 minutes.", no actions, no roster state). Nothing tells an officer the night before "Healers 5 of 8, 3 short, 5 have not answered", and nothing tells them an applicant arrived. Without that, an officer opens the app only if they remember. They keep the Discord bot and the sheet. Fix: an officer-only push at a configurable lead time. It is built from the badge and `composition.targets` data, with a "Mention healers" action reusing `RaidPingResult` and `RaidPingState.next_allowed_at`. This needs a scheduler. The **data exists**. The push kind is **needs server work**.

### SHOULD

3. **For officers, Needs you sits below the fold.** The Android officer capture ordering puts the raider's own sign-up card, then Who is coming, then Before you pull (4 rows plus a stamp), then Needs you. The officer's job is below the member's own readiness. Move Needs you to directly under the head when `officer` is true, and show the count as a badge there too. The tab badge "6" already exists, so a tap should land on the thing it counts. Files: `RaidTabScreen.kt:170` and `RaidRootView.swift:120`.

4. **Comp is roles only.** The member tab reads "Tanks 2 of 2 · Healers 5 of 8 · Damage 11 of 30", but the raid payload carries `composition.classes`, `composition.buffs[]` (Fort, MotW, Int, with `present` and `count`) and `composition.roles` and `targets`. iOS already has `CompositionRead.swift` with a buff summary, and Android has `CompositionRead.kt`, so the work is done and not surfaced on the tab. At minimum show a single "Missing: Fort, Int" line to officers on the Raid tab. Today you reach it only through the raid page, which is outside this lane.

5. **The Late minutes are not on Who is coming.** The live capture shows "Brightwen (Late · 15 min)", but the pre-raid list does not. Wanting "who is Late and by how much" the day before is exactly the thing an officer asks. It is already in the signup rows (`late_minutes`).

6. **The `raid.starting` push has no actions and no readiness.** "Starts in 60 minutes." should open Before you pull. It could carry the viewer's own result, for example "2 things to fix: flask, boots", because the raider is on the lock screen with 60 minutes left. The iOS tap-through (`notif-04`) goes to the raid, which is right, but the body is generic. The raiders get a push for the raid being posted and then nothing between, so there is also no "you have not answered" reminder. `you_unanswered` is already in the badge payload (`guild.raid-badge.json`), so a push for unanswered-with-24h-to-go can be computed.

7. **The "posted" push body lacks context.** "Sign-ups are open: Thu 5 Nov, 20:00 CST." is correct but says nothing about what they are being asked. The `answers` blob already holds "24 of 40 signed up" for the confirmation. Put "24 of 40 signed up, healers short by 3" in the first push, since that is the reason to answer now. The server has the counts.

8. **The iOS profile button draws the person glyph in the Raid captures.** It is on the BRIEF's already-found list, so I only confirm here that the fix is needed on every tab head, including the Raid tab, not just Home.

9. **Chasing is Discord-only and has a 6-hour mention cooldown.** `RaidPingState.next_allowed_at` exists and `NotAnsweredModel.kt` and `RaidFlowModels.swift` reference it. I could not confirm from the captures that the button disables or counts down. Please verify that "Mention healers" shows "Mentioned 2 healers. Next ping in 5 h 40 min" and does not just fail. If there is no "no_discord_names" line ("Selka has no Discord linked, DM not possible"), add one, because `RaidPingResult.no_discord_names` has it.

### NICE

10. **No attendance, streaks or lates on mobile.** `RecentAttendance` (`slipping`, cells per raid) and `GuildAttendance` already exist, so an officer line "Varro: 2 missed" is a free fit. Not-answered rows already say "was In last week" and "missed last week", which is a good start.

11. **No loot context for the raider.** "Last raid: 4 pieces" is a bare link. `guild.loot.json` and `guild.item-needs.json` exist. "You won nothing in 3 raids, 5 of 14 DPS have" is the fairness number that gets raiders to open the tab. This lives on a secondary screen, so only mention it here.

12. **Cold-state wording.** Android `none` state ("Guilds recruiting a Fury Warrior · Apply") is a good empty state. The `no-raid` state ("Your officers plan raids on the site") is honest but is a dead end for an officer, who gets no create-raid link. Link to the site's raid planner for officers.

13. **Parity.** Both platforms follow the same content spec (identical rows and strings), which is rare. Visual differences are styling only: iOS uses 4 separate blue buttons, Android a segmented control. The iOS buttons give no stronger selected state than the Android segmented fill, but both read.

## Value test

Would a guild pay $15 or $25 a month for this exact screen? **Not for the screen as it is, but the mechanism is close.** The Not answered screen plus applicant review plus a one-tap lock-screen sign-up are things Discord and a sheet do badly on a phone, and a raider who has never opened the app gets a reason to on raid day from the Before you pull list. What is missing is anything that makes the officer open it unprompted. They get no push, they cannot see roster readiness, and the roster-level numbers (comp buffs, attendance, loot fairness) that the sheet gives them are not here.

**The single addition that flips it:** finding 2 plus finding 1. That is an officer push ("Healers 5 of 8, 5 have not answered, 4 not synced") with a Mention action, and a raid-wide readiness line. All the underlying data exists today.

VERDICT: CLOSE

The two additions: (1) an officer-only push built from `raid-badge.officer` and `composition.targets` with a one-tap "Mention healers"; (2) a raid-wide readiness line for officers on the Raid tab, from roster `addon`, `logged_at` and `RecentAttendance`. Smaller items to ship with them: Needs you moved above Before you pull for officers (finding 3), the Late minutes on Who is coming (5), and the "raid.starting" push carrying the viewer's missing flask and enchant (6).

Key files: `apps/android/app/src/main/kotlin/gg/foreversixty/app/raids/flows/RaidTabScreen.kt`, `apps/ios/App/Raids/Flow/RaidRootView.swift`, `apps/ios/App/Raids/Flow/NotAnsweredView.swift`, `apps/android/app/src/main/kotlin/gg/foreversixty/app/raids/flows/NotAnsweredScreen.kt`, `docs/contracts/fixtures/app/push.raid-starting.apns.json`, `docs/contracts/fixtures/app/guild.raid-badge.officer.json`.

## Round 2 (2026-11-10, base 9bbf15e37)

**Round 2 verdict: guild-officer-reviewer, Raid tab, Home raid card, raid and officer pushes (iOS and Android)**

**Page job:** Tell a raider when the raid is, whether they're in, and what they're missing. Tell an officer who is short and who hasn't answered, and let them act from a phone. Earn an unprompted open from the officer.

**Judged from:**
- Android parity `android-raid-officer.png` and `.txt`, plus `android-raid.txt`.
- iOS `raid-01-top.png`, a non-officer capture.
- Android `notif-01-shade.png` and `notif-02-expanded.png`.
- The push fixtures, `api/internal/guilds/raids_events.go`, and `ApplicantPush.kt`.

I could not run either app. I did not open the flow PNGs or the Roborazzi PNGs for the Home card. I took the Late-chip, saving, failed and offline states from the code, the file names and your brief. I saw no iOS officer capture, so iOS officer behaviour is code-only. I did not find the officer readiness line in the Swift or Kotlin sources, only in the strings and in `RaidFlowContent.swift`/`RaidTabScreen.kt` comments. I did not read the Home raid card at all.

## Round-1 findings

| # | Finding | Status |
|---|---|---|
| 1 | Officer can't see raid readiness | **FIXED in part, and not enough to flip it.** The officer capture shows `4 signed In have not synced this week` as the last line of Who is coming. It is a count with no names and no tap. It is also below "Needs you", the sign-up card and the comp list. An officer who sees it can't tell who to chase. |
| 2 | No officer push | **FIXED in part.** The fixtures exist for `raid.short` and `guild.applicant`, and the body names the short role and the not-answered count. Neither push has a "Mention healers" action, and `raid.short` has no actions at all. See N1 and N2. |
| 3 | Needs you below the fold | **FIXED.** In the Android officer capture it is directly under the head, with "6 open", and it matches the tab badge 6. iOS has the same ordering in code (`RaidRootView.swift` line 106). |
| 4 | Comp is roles only | **FIXED.** Android reads `Missing Fortitude · Windfury 1 of 2` and iOS reads `Missing Blessing of Kings, Blessing of Might, Blessing of Wisdom and 2 more`. The iOS line is long and ends in "and 2 more", which hides the names an officer needs. |
| 5 | Late minutes missing | **FIXED.** `Brightwen (Late · 15 min)`. |
| 6 | `raid.starting` body is generic | **FIXED.** `Starts in 60 minutes. No flask in your bags.` It has no actions or deep link to Before you pull, only the raid link. That's acceptable. |
| 7 | `raid.posted` has no counts | **FIXED in the payload, but the shade clips it.** Android's collapsed shade reads `Sign-ups are open: Thu 5 Nov, 8:00 pm CS…`, so "24 of 40 signed" is cut off. `notif-02-expanded.png` is a blank spinner on the Home tab, not a notification, so I can't confirm it is expanded. See N3. |
| 8 | iOS profile glyph | **KEPT (ruling accepted).** The owner exception for portraits holds, and `raid-01-top.png` shows the stand-in ring with the faction badge. I accept it. |
| 9 | Mention cooldown line | **STILL OPEN, and agreed it's another lane's.** I can see `raid_na_noDiscord` and a recently-mentioned plural in the code, which is good. I could not see `next_allowed_at` rendered as a countdown. |
| 10 | Attendance on mobile | **STILL OPEN.** It is a NICE item and nothing changed. |
| 11 | Loot fairness | **KEPT.** It lives on a secondary screen. |
| 12 | `no-raid` dead end for an officer | **STILL OPEN.** The `android-raid-no-raid` text is unchanged from round 1. |
| 13 | iOS control | **FIXED.** It is now one connected four-way group with the chosen answer in teal, on par with Android. |

## New findings

### BLOCKER

**N1. The officer push has no officer action.**
- **What's wrong:** `push.raid-short.*.json` has no `actions` array. Every officer reply happens after opening the app, then Needs you, then Preview mention, then confirm. The brief says officer speed means one or two clicks from where the officer is looking.
- **Fix:** Add a `raid_mention_short` action. It reuses the `RaidPingResult` call and honours `RaidPingState.next_allowed_at`. Android has room for it (three-action limit) because `raid.short` needs only one button.
- **Status:** The data exists. This is server work in `api/internal/guilds/raids_events.go` plus a handler in `apps/android/.../push/PushActions.kt` and the iOS category.
- **Impact:** This is what finding 2 asked for, and it is the one thing that moves an officer off Discord. Leaving it out is a decision I disagree with. It is the difference between a notification and a tool.

### SHOULD

**N2. The applicant push offers Decline but not Accept.**
- **What's wrong:** `ApplicantPush.kt` registers only `application_decline`, so the shade offers the destructive choice and withholds the positive one.
- **Fix:** Offer Accept in the shade too, or neither. Declining a Holy Priest 60 with full T1 from a lock screen is a worse failure than accepting one slowly. If Accept is withheld on purpose, keep Decline off the shade as well and open the applicant sheet.
- **Where:** `push.guild-applicant.apns.json` and `ApplicantPush.kt`.

**N3. Push bodies carry the counts too late for a collapsed shade.**
- **What's wrong:** Android's collapsed raid-posted notification clips at about 40 characters, ahead of the counts.
- **Fix:** Reorder to lead with the signed count (`24 of 40 signed. Thu 5 Nov, 8:00 pm CST.`) and set `BigTextStyle` for the expanded view. `raid.short` is similar: the short role is at the end of a 60-character body and will clip before "Healers 5 of 8".
- **Where:** `raids_events.go` `raidEventCopy`.

**N4. The readiness line has no names and no destination.**
- **What's wrong:** `4 signed In have not synced this week` cannot be acted on.
- **Fix:** Make it tappable to a list of the unsynced raiders with a Discord mention or DM, reusing the Not answered rows. The data is the guild home roster `logged_recently` against the sign-ups. If per-raider consumable or enchant state isn't in the roster payload, that part is **needs data**.
- **Where:** `RaidTabFacts.kt`, `RaidFlowContent.swift` line 138.

**N5. The Who is coming counts contradict the Needs you counts.**
- **What's wrong:** The Android officer capture says `5 have not answered, 2 of them healers`, and the line below the roles says `14 not answered`.
- **Why it matters:** If these are different populations, such as roster members vs raiders who normally attend, the screen needs to say so. A tab that disagrees with itself loses an officer's trust. It may be a fixture artefact, but the parity golden is the contract.
- **Also:** The Out count differs between `round-2-decisions.md` (`Out 4`) and the capture (`Out 3`).

**N6. Fixture date and time text contradict each other.**
- **What's wrong:** `push.raid-posted.apns.json` has the title line `Thu 5 Nov, 8:00 pm CST` while `answers` says `Sunday 7 pm`.
- **Why it matters:** The sign-up confirmation is the line a raider reads to trust the tap, and a wrong time on the confirmation is a no-show. The answers blob must come from the same formatter as the body.
- **Where:** The `answers` blob in `push.raid-posted.*.json`.

**N7. iOS Missing line truncates the names that matter.**
- **What's wrong:** Android's `Missing Fortitude · Windfury 1 of 2` reads in one glance, but iOS lists three blessings and "and 2 more".
- **Fix:** Order by raid impact and cap at two names, or show a count.

### NICE

- **N8.** The Android `moved` answer body in `push.raid-posted.*.json` begins with a space and has an empty title. Trim it.
- **N9.** I saw no iOS officer capture in `ios/`. Add `raid-officer-*` steps so the iOS officer ordering is reviewable.
- **N10.** `notif-02-expanded.png` should be recaptured with the expanded notification.
- **N11.** The `no-raid` state still has no officer link to create a raid.

## Value test

Would a guild pay $15 or $25 a month for this screen? **Closer, but no, not yet.** Round 2 fixed the screen itself: ordering, comp cover, Late minutes, and a push that names what is short. An officer opening the app sees their job first. What's missing is the push-to-action loop. The officer is told "Healers 5 of 8, 5 not answered" and has to open the app and click through to do anything. The unsynced line is a number they can't act on. Discord and the sheet still win for chasing people.

**Two additions that flip it:**
1. A one-tap "Mention healers" action on `raid.short`, using the existing `RaidPingResult` and `next_allowed_at`. The data exists. The push action is server and client work.
2. Make the unsynced line tappable, to named raiders with a mention. The data exists in the roster payload (`logged_recently`).

Fix N5 and N6 along with these. Both are consistency bugs that cost trust.

Key files:
- `/Users/jh/code/forever/.claude/worktrees/agent-aee50bfb43330b1e0/docs/contracts/fixtures/app/push.raid-short.apns.json`
- `/Users/jh/code/forever/.claude/worktrees/agent-aee50bfb43330b1e0/docs/contracts/fixtures/app/push.guild-applicant.apns.json`
- `/Users/jh/code/forever/.claude/worktrees/agent-aee50bfb43330b1e0/api/internal/guilds/raids_events.go`
- `/Users/jh/code/forever/.claude/worktrees/agent-aee50bfb43330b1e0/apps/android/app/src/main/kotlin/gg/foreversixty/app/raids/push/ApplicantPush.kt`
- `/Users/jh/code/forever/.claude/worktrees/agent-aee50bfb43330b1e0/apps/android/app/src/main/kotlin/gg/foreversixty/app/raids/flows/RaidTabFacts.kt`
- `/Users/jh/code/forever/.claude/worktrees/agent-aee50bfb43330b1e0/apps/ios/App/Raids/Flow/RaidFlowContent.swift`

VERDICT: CLOSE

## Round 3 (2026-11-10, base 7911bf87e)

# Round 3 verdict: guild-officer-reviewer, Raid tab (iOS and Android), Home raid card, raid and officer pushes

**Page job:** tell a raider when the raid is, whether they're in, and what they're missing. Tell an officer who is short and who hasn't answered, and let them act from the phone. Earn an unprompted open from the officer.

**Judged from:**
- Android `android-raid-officer.png` and `.txt`, `android-home-raid-day-officer.txt`, and `android-raid-no-raid.txt`.
- Android `notif-01-shade.png`.
- Push fixtures `push.raid-short.*`, `push.guild-applicant.apns.json` and `push.raid-posted.apns.json`.
- `raids_events.go` (`raidShortEvent`, `shortestRole`, `mentionAction`).
- `RaidShortPush.kt`, `OfficerPushActions.swift`, `RaidTabFacts.kt`, `RaidTabScreen.kt` and `RaidsFeature.kt`.

**Not verified:**
- I could not run either app.
- I did not open the iOS flow tests; I saw only that Flows tests exist.
- I did not open the Android flow PNGs; I listed them and took the signup, late and offline states from the file names and my round-2 reading.

## Round-2 findings

| # | Round-2 finding | Status |
|---|---|---|
| N1 | `raid.short` has no officer action | **FIXED.** Both fixtures carry `raid_mention_short`, "Mention healers", POST `/v1/guilds/2/raids/7/ping` with `{"role":"healer"}`, category `RAID_SHORT`. The server leaves the action out when the ping is in its cooldown or Discord isn't connected (`canPing`). Android's `RaidShortPushAnswerer` posts it and replaces the notification with the result line (`Mentioned 3 healers in Discord.`). iOS parses the same action into `OfficerPushAnswer.mention`. This is the one that moves an officer off Discord. See new findings P1 and P2 for what is still loose. |
| N2 | Applicant push offers Decline only | **FIXED.** The fixture carries Accept then Decline, with Accept first. |
| N3 | Push body clips before the counts | **FIXED.** The Android shade (`notif-01-shade`) reads `24 of 40 signed. Thu 5 Nov, 8:00 pm CST.` without clipping. `raid.short` leads with the short roles: `Healers 5 of 8, 5 not answered. 21 of 40 signed, 24 hours out.` |
| N4 | Unsynced line has no names or destination | **FIXED.** Android reads `Not synced this week: Kaelra, Thalrin, Brakka, Mirethil` (three names, then `and N more`). It is tappable (`onRosterOf`) and pushes `RaidsDestination.Roster(guild, keys)`. iOS wires the same tap (`openUnsynced`, `RaidsTab.swift:247`). I have no capture of the filtered roster. See P4. |
| N5 | Counts contradict | **FIXED.** Needs you says "5 have not answered", the foot says "5 not answered", and Bench 1 · Out 3 matches round-2-decisions. |
| N6 | Fixture time and `answers` disagree | **FIXED.** Body and `answers` both read `Thu 5 Nov, 8:00 pm CST`. `moved` now has a title and `cancelled` reads `Onyxia's Lair was cancelled.` |
| N7 | iOS Missing line truncated names | **FIXED in code, not seen.** The Covered line is capped at three names with `and N more`, as round-3 row 4 says. The Android capture reads `Missing Fort · WF 1 of 2`. I did not see an iOS capture of it, and `WF` is a data label I would spell out once. |
| N8 | `moved` leading space and empty title | **FIXED** (fixture reads `Sign-up moved` / `{name} moved to Bench.`). |
| N9 | No iOS officer captures | **STILL OPEN.** `ios/` has no `raid-officer-*` step and the officer-only round-3 content (Needs you, not-synced, Covered) is reviewable only from code on iOS. |
| N10 | `notif-02-expanded` was a blank spinner | **STILL OPEN.** It exists again in `android-r3`, and I did not re-open it. There is still no capture of a `raid.short` or `guild.applicant` notification, expanded or collapsed, on either platform. The shade capture is `raid.posted` only. |
| N11 | `no-raid` has no officer link | **STILL OPEN.** The `android-raid-no-raid` text is unchanged: `No raid posted yet.` |
| Earlier | Cooldown line (9), attendance (10), loot fairness (11) | **KEPT.** These stay with the other lanes' rulings. |

## New findings

### SHOULD

**P1. The Mention button picks a different role than the screen leads with.**
- **Problem:** `shortestRole` takes the role with the largest absolute gap. In the officer capture, Damage is 11 of 30 (19 short) and Healers is 5 of 8 (3 short). The Needs you card says "Healers short by 3" and Preview mention is for healers. A server built from that same roster would pick `Mention DPS`, which pings 19 slots, and the fixture's "Mention healers" would never occur.
- **Fix:** Align the two. Either the push and the card both lead with the same role, or the push offers the role the officer can actually fill (tank, then healer, then DPS) rather than the largest raw gap.
- **Data:** `notify.ShortRoles`, in `raids_events.go`.

**P2. The officer's mention result is not captured.**
- **Problem:** Success, cooldown (action absent), no-Discord and failure all exist in code and tests, but I saw no capture of the officer shade. This is the feature that decides whether the app replaces Discord, and it should be reviewable like the sign-up flow is.
- **Fix:** Add `notif-officer-*` captures: the collapsed `raid.short`, the expanded one, the result after the tap, and the applicant push with Accept and Decline.

**P3. `Not synced this week` is built on `logged_recently`.**
- **Problem:** `logged_recently` records a character login, not an addon sync. An officer reads "not synced" as "no addon data". The roster payload also carries `addon` and `sim_export`.
- **Fix:** Either say "Haven't logged in this week", or match on `addon` and `sim_export` so the word is true.
- **Data:** `GuildHome.roster[]` in `api/openapi.yaml`.

**P4. The filtered roster result is unseen.**
- **Problem:** The tap pushes `Roster(guild, keys)`, but I saw no capture showing the filtered state: a header such as "Showing 4 raiders" with a way to clear the filter. Without that, an officer who taps from the tab lands on a roster that looks unfiltered or lands on a screen with no explanation.
- **Fix:** Capture it on both platforms and confirm a filter chip and a clear action exist.

### NICE

- **P5.** The Android shade shows In, Late and Bench only. Android's three-action limit drops Out, so a raider who is out has to open the app. The expanded notification or a link in the body could carry it.
- **P6.** The shade capture shows a different test notification ("warmup") beside the raid one. Retake it on a clean emulator before it goes in the record.
- **P7.** The Home officer card repeats the 5 not answered and the applicant card from the Raid tab. That is fine, but it is the same Needs you content twice at the top of two tabs; a single count and a tap through would be shorter.

## Value test

**Would a guild pay $15 or $25 a month for this exact screen?** For the officer's phone companion, yes, and the two flips I named are in. The loop is now complete. The server tells the officer a day out which role is short and how many have not answered. One tap in the shade pings Discord, and the notification says what happened. The applicant push accepts or declines from the lock screen. Inside the app, Needs you sits directly under the head and the unsynced line names raiders. A raider gets a count-first push with In, Late and Bench, a Late follow-up, a "Before you pull" list that says where to fix each miss, and honest sync stamps. The sheet still wins for loot fairness, attendance and per-raider performance, but those live on the site, and this is a raid-day and away-from-PC tool, so that gap doesn't block the switch.

**What would still cost trust:** P1. If the shade offers "Mention DPS" while the tab says "Healers short by 3", the officer stops believing the button. Fix that and show the officer notifications on a device, and it holds.

Key files:
- `/Users/jh/code/forever/.claude/worktrees/agent-aee50bfb43330b1e0/api/internal/guilds/raids_events.go`
- `/Users/jh/code/forever/.claude/worktrees/agent-aee50bfb43330b1e0/docs/contracts/fixtures/app/push.raid-short.apns.json`
- `/Users/jh/code/forever/.claude/worktrees/agent-aee50bfb43330b1e0/apps/android/app/src/main/kotlin/gg/foreversixty/app/raids/push/RaidShortPush.kt`
- `/Users/jh/code/forever/.claude/worktrees/agent-aee50bfb43330b1e0/apps/ios/App/Raids/Push/OfficerPushActions.swift`
- `/Users/jh/code/forever/.claude/worktrees/agent-aee50bfb43330b1e0/apps/android/app/src/main/kotlin/gg/foreversixty/app/raids/flows/RaidTabFacts.kt`
- `/Users/jh/code/forever/.claude/worktrees/agent-aee50bfb43330b1e0/apps/android/app/src/test/screenshots/parity/android-raid-officer.png`

VERDICT: WORTH PAYING FOR

