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
