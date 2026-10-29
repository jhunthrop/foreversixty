# Discord captures D13 to D17: the owner's checklist

Needs the deployed build with the commands registered to `DISCORD_COMMANDS_GUILD_ID` (re-register after this deploy: the command set changed) and your Discord client. Save every screenshot under `design/reviews/2026-10-28-discord-guided-cards/discord/` as `<name>-desktop.png` (the desktop client, window about 640 px wide for the message) and `<name>-phone.png` (the phone app). Never paste a token, a webhook URL or an interaction token into a screenshot; crop them out.

Set-up once: the test server has the bot, a raid channel set, a webhook is not needed until D17. You need three views: (A) a Discord account not linked to Forever Sixty, (B) a linked member with two verified characters (one the main), (C) a linked officer. B and C may be one account if that account is an officer.

Expected look of every frame: `payloads-vs-boards.md` and the `*-real` PNGs beside the boards in `design/mocks/2026-10-28-discord-guided-cards/`.

## D13 `/bis`, and the 15-minute question (decides Ruling 12)

1. As A, type `/bis` and send with nothing filled in. Screenshot the card: `d13-bis-unlinked-1-card`.
2. Open the class list, pick `Hunter`: `d13-bis-unlinked-2-spec`. Open the spec list, pick `Beast Mastery`: `d13-bis-unlinked-3-ready`. Press `Show BiS`: `d13-bis-unlinked-4-result`. Press `Change picks`: `d13-bis-unlinked-5-change-picks`.
3. As A, send `/bis` again, pick only a class, press `Show BiS` before a spec: `d13-bis-nudge`.
4. As B, type `/bis` bare: `d13-bis-linked-1-main`. Change the spec and press `Use <main name>`: `d13-bis-linked-2-changed`. Press `Show BiS`: `d13-bis-linked-3-result`.
5. As B, send `/bis class:hunter` (a typed class): the card opens with Hunter chosen: `d13-bis-typed-class`.
6. The 20-minute test, the one that matters. As B, send bare `/bis` and write down the clock time. Leave the card alone for 20 full minutes (set a timer; do not touch the card or reply in the channel). Then open the spec list and pick a spec.
   - It worked (the card changed, no error): screenshot `d13-late-pick-accepted`. Result: presses are accepted at 20 minutes; Ruling 12 stays "no footer". Tell the coordinator to record "D13: presses accepted at 20 min, no footer" in the §9 row of `design/specs/2026-10-27-discord-bot.md`.
   - Discord drew `This interaction failed` or a red error: screenshot `d13-late-pick-failed`. Result: add `card.footer.window` to every card (plan Task 15 step 4).
   - Repeat once with a second card at 14 minutes, to be sure the limit is the 15 minutes of an interaction token and not something else: `d13-late-pick-14min-<accepted|failed>`.

## D14 `/char`

1. As A, `/char` bare: `d14-char-unlinked` (`Type a name` and `Link Discord`).
2. Press `Type a name`: the dialog: `d14-char-modal`. Submit `Nobody`: `d14-char-unknown` (the sentence and the card stay).
3. As B, `/char` bare: the main's card with `Share in channel`, `Someone else`, `Character page` and the own-characters select: `d14-char-own`.
4. Open the select, pick the other character: `d14-char-pick`.
5. Press `Someone else`, submit a real public character's name: `d14-char-someone-else`.
6. Press `Share in channel`: the public post (`d14-char-shared-public`) and your own copy, which must still show `Someone else` and the select (`d14-char-shared-own`).
7. Typed, for comparison: `/char name:<main name>` (`d14-char-typed`). The answer must read the same as step 3 apart from the controls.

## D15 `/signup`

Set-up: two planned raids in the guild, B has two verified characters.
1. As B, `/signup` bare: `d15-signup-card` (raid select, character select, `In` green, `Late`, `Bench`, `Out`, `Raid page`).
2. Open the raid select, pick the other raid: `d15-signup-raid-pick`. Open the character select, pick the other character: `d15-signup-char-pick`.
3. Press `In`: the card becomes the outcome sentence: `d15-signup-in`. Check the raid page shows the signup.
4. Lock the raid on the site, then press a status on an older card: the lock sentence in place: `d15-signup-locked`.
5. `/signup status:in` typed, for comparison: `d15-signup-typed`.

## D16 `/attendance`, `/readiness`, `/report`

1. As C, `/attendance` bare: the own record and the `Look up a member` select, `Type a name` before `Roster`: `d16-attendance-officer`. Open the select: `d16-attendance-select-open` (lowest first, chips as descriptions). Pick a member: `d16-attendance-pick`. Press `Type a name`: `d16-attendance-modal`; submit a name that is not a member: `d16-attendance-unknown`.
2. As B (not an officer), `/attendance`: no select: `d16-attendance-member`.
3. As C, with two planned raids, `/readiness` bare: `d16-readiness-select`; pick the other raid: `d16-readiness-pick`.
4. As C, with two done raids, `/report` bare: `d16-report-select`; pick the other night: `d16-report-pick`. Press `Share in channel`, then check your own copy still has `Another night`: `d16-report-shared-own`.
5. Demotion check (optional): demote C on the site while a card is open, pick again: the officers-only sentence in place: `d16-demoted-pick`.

## D17 Raid posts and reminders

Use the test server's raid channel, bot bound. Timing is by hand: run the job instead of waiting ten minutes.

    gcloud run jobs execute raid-posts --region us-east1 --wait

Do not print the job's environment, and crop any log line that holds a URL.

1. Within 7 days: on the site, create a raid that starts in 3 days, `Post to Discord` and `Remind the day before` left at their defaults (both on). The post appears at once in the channel: `d17-post-scheduled`.
2. Over 7 days: create a raid 9 days out. Nothing posts: `d17-nothing-yet` (the channel). Run the job when it is 7 days out, or move the raid to 6 days out and run the job: the post appears: `d17-post-deferred`.
3. Reminder as an edit: with the raid 24 hours out (edit its start to about 23 hours from now) and the post still the newest message in the channel, run the job twice: the post becomes `Tomorrow: ...`: `d17-tomorrow-edit`.
4. Reminder as a reply: in a second raid, post anything in the channel after the post, then bring the raid inside 24 hours and run the job: the post edits to `Tomorrow:` and a reply appears under it: `d17-reminder-reply`.
5. Webhook: in guild settings, switch the destination to a webhook (the bot's install off), create a raid inside 24 hours, run the job: `d17-reminder-webhook` (a content line with the masked link).
6. Withdraw: cancel the raid that has a reminder: the reminder reply is deleted and the post shows Cancelled: `d17-cancel-deletes-reminder`. For a webhook reminder, cancel and check the message is gone: `d17-cancel-deletes-webhook-reminder`.
7. Move: move a raid after its reminder: the old reply is deleted, the post shows `Moved:`: `d17-move`.

Note: the roster block and the `Which spec tonight?` card are not built yet (Task 10b, owner gate); expect the posts to show `Signed up` as a count only, as the report says.

## Reporting back

Return the file names saved, the D13 20-minute result in one line (accepted or failed), and anything that differs from the `*-real` frames.
