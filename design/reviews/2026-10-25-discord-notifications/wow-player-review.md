# Discord notifications: veteran player review (2026-10-25)

Reviewed: README, spec 6.1 and 6.2, and the 1440 and 390 captures of paste-empty, paste-invalid, saved, saved-just-now, warning, warning-long, switched-off-long, lapsed, remove-confirm, replace, test-failed, free, feature-off and the failed-five-times row. Not every 360/768/1024/2000 capture was opened. The 6.2 Discord texts are judged as written; no real-channel capture exists.

**Verdict: SHIP WITH FIXES**

1. `discord-switched-off-long-panel-390.png`, `discord-warning-long-panel-1440.png`: the cut Discord error reads `...someone with Mana....`. It looks broken. Cut at a word boundary, drop the dangling clause, and do not put a full stop after an ellipsis. If Discord's text is long, show only the part before the colon: `Invalid Webhook Token`.
2. Same two captures: the switched-off notice does not say what to do next. Change `Sends a test message. If Discord accepts it, posts resume.` to `Fix it: make a new webhook in Discord and use "Use a different webhook", or press Test and switch on if you only changed the channel.` Put a "Use a different webhook" link inside the notice, because it currently sits 70 px below the buttons.
3. All saved-state captures: `Posts to channel 1198765432101234568` is a snowflake ID. Discord shows it only in developer mode, and no raid lead recognises it. Show the channel name (and server name) if the API can return them. If it cannot, change the line to `Posts to a Discord channel (ID 1198…4568)` and use a monospace font so it reads as an ID.
4. `discord-warning-*`: the 1 to 4 failure warning is muted small print, yet it is the one line that tells me my pings are dying. Give it the same coloured bar as switched-off, and add `Send test message` as the next step: `The last 2 posts failed. At 5 in a row we switch it off. Discord says: Unknown Webhook. Send a test message to check it.`
5. `discord-paste-empty-*` and `discord-replace-*`: the how-to is a single line of muted text. Put it first, as a numbered line: `1. In Discord, open the channel's settings, then Integrations, then Webhooks. 2. New Webhook, then Copy Webhook URL. 3. Paste it here.` Also close the gap of about 50 px between the input and the helper text.
6. Logs wording: `A log file you upload finishes parsing` is not how players speak. It also says "log file" while the Logs menu says reports. Use `A combat log you upload finishes parsing` or `A log report finishes parsing`. The same applies to the intro, the test message and /premium.
7. `Posts every log you upload, including private ones; each message says who can open it.` will make people nervous, and it should. The problem is that the message goes into a shared guild channel. Make it a plain warning: `Private logs post too. The message shows the log's name and "Only you"; anyone in the channel sees that it exists.` Also decide whether private logs should be a separate checkbox that is off by default (see owner question 1).
8. `discord-saved-1440.png`: the checkboxes are the browser's native bright blue, and the 70 px of empty space under the buttons looks like a missing row. Restyle the checkboxes in gold to match the site, and collapse the reserved status line when it is empty. Pre-ticked boxes on a first save are acceptable, but the Save button should say what it will do: `Save webhook`.
9. `discord-row-off-failed-five-times-*`: `Off · failed five times` is understandable, but it reads as a status with no hint of action. Use `Off · needs attention`, linked to #discord. Also `Off · set up` on the free-plan Plan panel reads as an invitation, so make it `Off · Premium`, or hide the row for non-members.
10. 6.2 Discord messages. `Took` should read `Run time`. `Fights` should read `Boss fights` if kills are trash-excluded. `Visible to / Only you` is fine and good. `Your Frost Mage sim finished` is fine, but include the character name when there is one, because a guild officer with three mages will otherwise not know which it was. `Result 1,204 DPS` needs a unit hint, such as `1,204 DPS (sim)`.
11. `discord-skeleton-*`: the Discord panel has no reserved height, so the page jumps on load (CLS 0.35 at 390). The lane deferred this; I would not block on it.

## What works

- The panel asks one thing at a time. Paste, save, send a test, and the result appears right under the buttons.
- A wrong URL produces a specific message and the typed value is not echoed (`paste-invalid`). `Discord said: Missing Permissions.` (`test-failed`) is honest and short.
- Remove asks once, in words I understand: `Remove this webhook? Posts stop and we forget the URL.` Lapsed says we kept the webhook and what brings it back, which is good.
- Encrypted-storage and never-shown-again copy builds trust. No marketing words anywhere.
- The Premium row has seven real states. The free panel is one sentence and a link to plans, without a pitch.
- The 390 layouts stack cleanly and keep the buttons together.

## What a player would trip on

- The channel ID (fix 3) and the cut error (fix 1).
- Switched off gives the cause but not the cure (fix 2). Most players hit this after deleting a webhook, and the remedy is a new webhook, not Test.
- Warning text is small print (fix 4).
- "log file you upload" reads like a raw WoWCombatLog.txt (fix 6). Players say "log" or "report".
- Private logs posting into a guild channel (fix 7).
- Native blue checkboxes, grey in lapsed (`discord-lapsed-panel-390.png`), look unfinished next to the gold UI.
- Free-plan page: the free sentence in `discord-free-panel-390.png` has loose line spacing, so it looks like two separate lines.
- When the feature is off for a deployment (`feature-off`), the panel disappears but the Plan row still says `Off`. The row should vanish too.

## Copy I would change

| Current | Proposed |
|---|---|
| `Last error from Discord: ...someone with Mana....` | cut at a word, no trailing full stop; or only the headline `Invalid Webhook Token` |
| `Sends a test message. If Discord accepts it, posts resume.` | `Fix the webhook in Discord, or use a different one. Test and switch on posts again once Discord accepts it.` |
| `Posts to channel 1198765432101234568` | `Posts to #channel-name` (or `Posts to a Discord channel (ID ...4568)`) |
| `A log file you upload finishes parsing` | `A combat log you upload finishes parsing` |
| `Posts every log you upload, including private ones; each message says who can open it.` | `Private logs post too. The message shows the log's name and "Only you".` |
| `In Discord: Channel settings, Integrations, Webhooks, New Webhook, Copy Webhook URL.` | numbered steps (fix 5) |
| `The last 2 posts failed. At 5 in a row it switches off. Last error from Discord: Unknown Webhook.` | `The last 2 posts failed. At 5 in a row we switch it off. Discord says: Unknown Webhook. Send a test message to check.` |
| `Off · failed five times` | `Off · needs attention` |
| Embed field `Took` | `Run time` |
| `Your Frost Mage sim finished` | `Your Frost Mage sim finished` plus character name in the Run field |

## Questions for the owner

1. Should private logs post by default? A third checkbox for private logs, off by default, would remove the nervousness.
2. Can the API return the channel and server name when the webhook is saved? Discord's webhook GET returns `channel_id` and `guild_id` only, so the name may need a bot. If not, the ID line has to be honest about being an ID.
3. After 5 failed posts, do we also tell the member by email or on the site's next visit? The switch-off is silent otherwise.
4. Do live-logging uploads never post, and does the Premium row say so? Players who live-log will wait for a ping that does not come.
5. Does the sim message include a link to the character's BiS gap, so the ping leads to a better character and not just a number?

## Follow-up lane re-check (a71bbb90)

**Verdict: SHIP WITH FIXES.** Nothing here would be mocked in guild chat; the two things I would fix before launch are copy, not layout.

1. Changelog: `The account page now shows the name of the webhook, so you can tell which channel it is.` overpromises. The page shows the name the member gave the webhook (`Posts through “Raid pings”`), not the channel, and the same changelog's first paragraph still says `the page only shows which channel it posts to`. Say `shows the webhook's name and a short channel ID`.
2. Changelog: the private-logs sentence is buried at the end of a long paragraph. A player who set up a webhook yesterday and skims will miss that their private logs are still posting. Move it to its own paragraph and say what it means (see Copy).
3. `discord-events-logs-off-private-disabled-panel-1440.png`: `Include private logs` is greyed but ticked, directly under `A combat log you upload finishes parsing` (unticked), and the note under it says `Off, private logs never post.` The box says on, the note says off. Either leave a disabled box unticked, or when disabled make the note read `Kept on, but nothing posts while logs are off.`
4. `discord-row-on-2-events-1024.png`: the Premium row value breaks as `On ·` then `2 events` on the next line, and the line above wraps to `only`. The non-breaking space the lane listed as polish is needed; the row looks broken at 1024.
5. `discord-warning-four-long-panel-1440.png`, `discord-switched-off-long-panel-390.png`: the Discord error is cut at `someone with…`, which is exactly where the useful part would be. Cut at the last comma, or show `Invalid Webhook Token: the token in this webhook URL was changed, or the webhook was deleted.`

## What landed well

- Numbered steps (`discord-paste-steps-1440.png`) are right: `Integrations, then Webhooks`, `Choose New Webhook, name it something you will recognize, then Copy Webhook URL`, `Paste it below.` Step 2 asking for a name is what makes the saved line meaningful. The trust line sits 21 px under the input and reads fine.
- `Posts through “Raid pings”` is better than the snowflake ever was. Even the unnamed variant (`Posts to a Discord channel` / `Channel 1198…4568`) is honest about what it is. Not getting the channel name is Discord's limit, and I accept it.
- The three event rows read as one decision with a child: `Include private logs` indented under the log row, with a note that says who sees what (`everyone in the channel sees the log's name and “Only you”`). Off by default for new webhooks is the right call.
- Switched-off notice (`discord-switched-off-long-panel-390.png`): red bar, cause, `Fixed it in Discord? Test and switch on. If the webhook is gone, make a new one.`, and `Use a different webhook` inside the notice, above `TEST AND SWITCH ON`. That is the cure I wanted, in the order a player does it.
- The 1-to-4 warning now ends with `Send a test message to check.` and the first line is brighter than the error under it.
- Premium row: `Off · needs attention` is linked and clear. `Off · Premium` on the free plan (`discord-free-1440.png`) is quiet, with no pitch, only the one sentence under the Plan panel. `On · 2 events` is correct.
- Gold checkboxes are gold on the Discord rows and on the account `You` row, and look like the rest of the site.
- Loading, warm: the skeleton bars sit where the panels land (footer within 0.5 px in all six states). That fixes the jump.
- Messages as written in 6.2: `Run time`, `Boss fights` `9 · 7 kills`, `Frostbolt's Frost Mage sim finished`, `Visible to` `Only you`, `Test message from Forever Sixty`, and `finishes parsing, private logs included.`. All read like something a raid lead would be glad to see. `Result` `1,204 DPS` without `(sim)` is fine now that the title says it is a sim; I withdraw that item.

## Still trips a player

- Cold load (`discord-loading-cold-1440.png`): the skeleton shows three bars per right-hand panel and no Discord frame, and the footer sits at about y 990. When the data arrives the page grows by about 660 px and the footer drops. Warm is solved; cold (first visit, new tab) is not. Not a blocker; a player will see the footer fall once.
- A webhook that still carries Discord's default name will show that default in `Posts through “…”`. The step asks for a name, but a pasted old webhook will read oddly. Acceptable, but know it.
- `Channel 1198…4568` is in the body font, bold. A monospace face would say "this is an ID" at a glance.
- Panel heights: the paste form is 647 px against its 640 cap and the one-line switched-off panel 561 against 560. Nobody will notice; update the cap.
- The saved panel still leaves about 50 px between the buttons and `Use a different webhook`, which looks like a missing row (`discord-saved-unnamed-panel-1440.png`). Better than 70; not fixed.
- Live logging never posts (6.2 says so). The event label says "upload", but a player who live-logs and sees nothing will blame the site. A one-line note under the log row would head that off.

## Copy I would change

| Current | Proposed |
|---|---|
| `The account page now shows the name of the webhook, so you can tell which channel it is.` | `The account page now shows the webhook's name (the one you gave it in Discord) and a short channel ID.` |
| `...the page only shows which channel it posts to.` | `...the page only shows the webhook's name and a short channel ID.` |
| `New webhooks leave private logs off; a webhook saved before today keeps posting them until you untick Include private logs.` | Own paragraph: `Private logs. New webhooks do not post private logs. If you saved a webhook before today, it still posts them to your channel until you untick Include private logs on your account page.` |
| `Off, private logs never post.` (under a disabled, ticked box) | `Kept on, but nothing posts while logs are off.` |
| `...deleted from the channel settings by someone with…` | `...the webhook was deleted.` (cut at the last comma, not mid-clause) |
| `On · 2 events` wrapping at 1024 | non-breaking spaces so it stays on one line |
| `Fixed it in Discord? Test and switch on. If the webhook is gone, make a new one.` | Fine as is. |
| (none) | Under the log row: `Live logging does not post.` |

## Questions for the owner

1. The changelog tells people to untick a box. Do you also want to tell the people it affects? Anyone who saved a webhook yesterday is still posting private-log names into their channel today. A one-time line on the account page (`This webhook was saved before private logs became optional; it still posts them`) would reach them, and a changelog entry will not.
2. Will the messages ever link to the character's BiS gap (my question 5 from yesterday)? It is the one part of the ping that moves a player toward a better character, and it is not in 6.2.
3. After five failures the member is told only if they open the account page. Do we want an email or a notice on the next visit? (Question 3 from yesterday, still open.)
4. Real channel: no capture shows the three embeds, so `Frostbolt's Frost Mage sim finished` with `Run`, `Run time` and `Result` has not been seen in Discord. Do not call spec item 9 done until it is.

## Connect lane re-check (3b6cec94)

**Verdict: SHIP WITH FIXES.** The flow is clear and honest, and the remove questions are the best copy in the block. Three small fixes before I would send it to a guild.

1. `discord-connected-destination-390.png`, `-360`, `-1440` and `discord-connected-saved-panel-390.png` still show `Channel 1198…4568· Open in Discord` with no space before the dot. You said it is fixed in source; the committed PNGs are stale, so recapture them. I will not call it fixed until I see a capture without the glitch.
2. The explanation says nothing about permissions. `Discord asks which server and channel, then lets Forever Sixty post there. We keep only that webhook, encrypted.` Add that you need Manage Webhooks on that channel. Today a raid leader learns it only from the red failure line in `discord-return-failed-panel-390.png`, after a trip to Discord and back. Proposed text is below.
3. `Connected to a Discord channel` over `Channel 1198…4568` does not tell the player which channel. `Open in Discord` helps only if the player is in that server and can see the channel. Use the channel name or the webhook name Discord returns at connect time, if the API gives one, and keep the id as the small second line. If it cannot be done, fix the first line so it does not read like a name: `Connected to a Discord channel` looks like a placeholder someone forgot to fill in.

### What landed well

- `discord-connect-1440.png` and `discord-connect-panel-390.png`: connect is the one gold button, paste is a quiet underlined link. It is obvious which to do. The gold button is the only call to action, and the account page is easy to read around it.
- `discord-connect-opening-panel-390.png`: `Opening Discord…` replaces the label and nothing jumps. That is the right amount of feedback for a full-page redirect.
- Return outcomes: `Cancelled in Discord. Nothing was connected.` is plain and calm. `Could not connect. Check that you can manage webhooks in that channel, then try again.` is red, tells me the likely cause and what to do. The gold button stays focused so I can retry at once. Good.
- `discord-return-connected-panel-390.png`: `Connected. Send a test message to see it in your channel.` with focus on `Send test message` is the right next step. It also answers "did it work?" without a guess.
- `discord-connect-reconnect-panel-390.png`: `Connect a different channel` is a link, not a button, so nobody clicks it by habit. The note `Connecting again replaces this channel and removes its webhook from Discord.` reads as information, not alarm. It tells me the old channel will go quiet, which is what I need to know.
- Remove questions are honest and different on purpose. Connected: `Remove this connection? Posts stop and we delete its webhook from Discord.` Pasted: `Remove this webhook? Posts stop and we forget the URL. It stays in Discord until you delete it there.` A player who pasted their own webhook is told it is still theirs to clean up. That is correct and rare.
- `discord-connect-paste-view-panel-390.png`: the way back, `Connect a Discord channel instead`, is there. Nobody is stuck in the paste form.
- `discord-connect-replace-view-panel-390.png`: replacing a pasted webhook offers connect first, paste second, `Cancel` last. Same order as a fresh setup, so there is one thing to learn.

### Still trips a player

- **Guild channels and permissions.** The raid leader who sets up the guild's channel needs Manage Webhooks on it, and an officer who does not have it gets Discord's own error page or a bare cancel. The lane's only hint is the failure line after the fact. If Discord returns the user to us with the same flag for "no permission" as for "cancelled", the player sees `Cancelled in Discord. Nothing was connected.` when they did not cancel anything. Please check what Discord sends back in that case.
- **Which server.** A player in six servers, two of them with a `#raids` channel, cannot tell from `Channel 1198…4568` which one they picked. See fix 3.
- **Authorise step.** Discord shows its own `Forever Sixty wants to access your Discord account` style screen and asks the player to pick a channel in a drop-down. We do not warn them. A short `You will see a Discord screen asking you to authorise Forever Sixty` takes the surprise out of it. I could not check the real screen (capture 29 is manual), so this is from how Discord's webhook OAuth works.
- **Nothing happens.** There is no popup blocker, but a slow or blocked Discord (work network, adblock list) leaves `Opening Discord…` on screen with a shrunken button and no way out except reloading. After about ten seconds I would expect the label to go back and say `Discord did not open. Try again, or paste a webhook URL instead.` I did not see such a state in the captures.
- **The 40 px gap.** In `discord-connect-panel-390.png` there is a large empty band between the explanation and `Paste a webhook URL instead`. It is the reserved status line, and it looks like a missing paragraph. The captures README says the reserve is deliberate for layout shift; fine, but at 40 px it reads as a hole. Tighten if the shift gate still passes.
- **The destination dot spacing**, see fix 1.
- **Reconnect note stays on screen.** `Connecting again replaces this channel and removes its webhook from Discord.` is shown all the time under the link. It is accurate, but a quiet note under a link that most players will never click is clutter beside the private-logs line. Acceptable; I would not hold the ship for it.

### Copy I would change

| Current | Proposed |
|---|---|
| `Discord asks which server and channel, then lets Forever Sixty post there. We keep only that webhook, encrypted.` | `Discord asks you to pick a server and channel, and to authorise Forever Sixty. You need Manage Webhooks on that channel. We keep only the webhook, encrypted.` |
| `Connected to a Discord channel` | The channel or webhook name when we have one (`Connected to #raid-pings`); otherwise `Connected to a Discord channel (name not provided)` is worse, so keep the current line and lean on `Open in Discord`. |
| `Channel 1198…4568· Open in Discord` | `Channel 1198…4568 · Open in Discord` (space before the dot) |
| `Cancelled in Discord. Nothing was connected.` | Fine as is. |
| `Could not connect. Check that you can manage webhooks in that channel, then try again.` | `Could not connect. You need Manage Webhooks on that channel. Ask a server admin, or paste a webhook URL they give you.` (the second half is the raid leader's real way out) |
| `Connect a different channel` | Fine as is. |
| `Connecting again replaces this channel and removes its webhook from Discord.` | Fine as is. Keep the word `replaces`. |
| `Remove this connection? Posts stop and we delete its webhook from Discord.` | Fine as is. |
| `Remove this webhook? Posts stop and we forget the URL. It stays in Discord until you delete it there.` | Fine as is. |
| `Opening Discord…` | Fine as is, with a timeout fallback (see above). |

### Questions for the owner

1. What does Discord's webhook response give us at connect time? It normally includes the webhook name, channel id and guild id. Can we show the channel name or the server name, or at least keep the webhook name and show that? Fix 3 depends on it.
2. When Discord returns without a webhook because the member lacks Manage Webhooks, do we get a different error from a cancel? If not, the `Cancelled in Discord` line will be wrong for exactly the raid-leader case, and I would merge the two messages: `Nothing was connected. If you did not cancel, check that you can manage webhooks in that channel.`
3. Is a channel id plus `Open in Discord` link a link to the channel (discord.com/channels/guild/channel)? If the guild id is not stored the link cannot work. I could not verify it from captures.
4. What happens to the old webhook in Discord if the delete call fails during `Connect a different channel`? The note promises it is removed. The row is deleted first (commit 9b51612f), so the player never sees a failure. Do we want a line that says `Could not remove the old webhook from Discord; delete it there`?
5. Is the 40 px reserve gap staying, or can it come down?
