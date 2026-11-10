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

