# Discord bot: player and officer review, 2026-10-27

Reviewer: wow-player persona (Classic raider, guild officer). Read from the PNGs in this folder: the Discord message renders (faithful renders of the real payloads, not the real client) at phone and desktop, and the site captures for guild settings, the account link row and the raid form. Spec: design/specs/2026-10-27-discord-bot.md.

## Verdict: SHIP WITH FIXES

The core loop works. The raid post reads right on a phone between pulls: the title, `When` in my own time, `33 of 40 · 3 tanks · 8 healers · 20 DPS · 2 late`, then a green `In` and three grey buttons, all on one screen. The character picker is plain (`Choose a character`, main first, spec and role under each name). Nothing leaks: no names in the public post, readiness and attendance stay ephemeral with no Share, and other people's attendance is officer-only. The officer setup is the clearest I have seen for a bot: one gold button, a sentence that tells the GM they need Manage Server, and a notice that names the channel and the exact Discord permission to fix. The wording sounds like a raid tool, not a SaaS bot. My raiders would use the buttons instead of the site.

Not approved, because I have not seen them rendered: `/char`, `/raid`, `/report`, `/signup`, `/link`, and the full `/bis` answer (text line plus two alternatives). The real-client captures D1 to D12 are still owed. Ship them dark, or ship once they are captured and pass.

## Findings

1. MED, discord-msg-press-moved-phone.png: "Zulmara is In as DPS for Onyxia's Lair. 40 of 40 signed up. Mendwhisper moved to Bench." reads in Discord as though I just bumped another raider to the bench. Officers will get DMs about it. Fix: "Your other character Mendwhisper moved to Bench."
2. MED, discord-msg-press-in/bench/out/already/late-phone.png: every reply says only "for Onyxia's Lair". A guild running two Ony or MC nights a week cannot tell which raid the press landed on. Fix: add the day, "for Onyxia's Lair, Thu 10 Dec."
3. MED, discord-msg-post-open-phone.png and discord-bot-raid-form-post-to-discord-bot-only-crop-390.png: the post can never ping a role (`allowed_mentions` is empty and the form has no ping option), so it lands silently. Raid-Helper guilds ping @Raiders on every post, and without a ping half the roster misses it. Fix: add an optional "Ping a role" choice to the channel settings and send it in the message content.
4. MED, discord-bot-04-warning-3-failures-crop-390.png: "Last error from Discord: Missing Access." is a raw API error, and a non-technical GM cannot act on it. "At 5 in a row it switches off" does not say what switches off. Fix: translate it the way the 2b notice does ("The bot cannot see #raid-signups. Give the Forever Sixty role View Channel there.") and say "posting to Discord switches off".
5. MED, discord-msg-bis-slot-phone.png: the capture is not a real /bis answer. It has no item icon thumbnail, no "Main hand for a … Hunter, …" text line, no two alternatives, and the author line reads "Zulmara" instead of "Best in slot". Its fixture is also Quel'Serrar, a tank sword from an Onyxia chain, sitting in a Beast Mastery hunter's main hand with the source given only as "Dire Maul". Anyone would mock that screenshot. Fix: re-capture from the real BiS data with the icon, the text line and the alternatives; the source should read as the quest chain does.
6. MED, discord-msg-attendance-phone.png: "Came 7 of 8 · 26 of 30 raids" has two numbers and no units. In Discord, without the roster column around it, a raider cannot tell that 7 of 8 means the last 8. Fix: "Came 7 of the last 8 · 26 of 30 raids in all".
7. MED, discord-msg-readiness-phone.png: "Kraggor: gear 12 below the raid line" is a number with no unit. Fix: "gear 12 item levels below the raid line".
8. MED, discord-bot-12-permissions-no-webhook-picker-open-crop-390.png: right after "The bot cannot post in #raid-signups", the picker opens with #raid-signups preselected under the lead "Bot added to Olympus Discord". The channel it just said was broken is the default, and the lead says the bot was just added. Fix: leave the unusable channel out of the list (or clear the select to "Choose a channel"), and use "Raid channel in Olympus Discord" as the lead outside the install return.
9. MED, discord-msg-press-select-phone.png (missing): the role always follows the character's saved spec, so a resto druid or holy paladin who is asked to DPS or tank tonight cannot sign as that role from Discord. Raid-Helper allows it. Fix: add an optional `role` on /signup, or a role select after a hybrid's press.
10. LOW, discord-msg-press-in-phone.png vs discord-msg-post-open-phone.png vs discord-msg-post-done-phone.png: the fixtures disagree ("31 of 40" in the reply, "33 of 40" on the post, "Came 29 of 31" on the Done post), and the Done post still says "(in 5 hours)". Players will compare the numbers. Fix: use one consistent fixture for the real-client captures.
11. LOW, discord-msg-readiness-phone.png: "The 31 signed In or Late." is missing a noun. Fix: "31 signed up In or Late; 3 have something to fix."
12. LOW, discord-msg-press-select-phone.png: "· your main" sits in the option label, not the description, and the pressed state "In" is not bold as the spec says. A class icon beside each name would help someone with eight alts. Fix: move "your main" to the description and bold the state.
13. LOW, discord-bot-raid-form-post-to-discord-bot-only-crop-390.png: "Post to Discord" does not say where it posts or that the post gets buttons. Fix: "Post to #raid-signups with sign-up buttons".
14. LOW, discord-bot-02-state-4-offer-crop-1440.png: "Channel 1198…4568" shows a raw channel ID to a GM. This is G1's existing webhook state and is not new here. Fix: show the channel name where the webhook API gives it, else "a Discord channel" alone.

## Re-review 1

Same reviewer, after 705c0093d. This round I read the handler-backed renders: every D1 post state, every D3 press reply, plus /raid, /signup, /char (upgrades, no gear, unknown), /bis (slot, top, shared, after share), /report (member, officer, shared, after share), /attendance (own, officer, refusal), /readiness (officer, member, unlinked, lapsed, Guild tier, nobody), /link (both states) and /help. I also re-read the site crops for 4, 8 and the account row.

### Findings from the first round

1. Fixed. discord-msg-press-moved-phone.png reads "Your other character Zulmara moved to Bench."
2. Fixed. discord-msg-press-in-phone.png and discord-msg-signup-phone.png both say "for Onyxia's Lair, Thu 10 Dec."
3. Deferred to the owner (role ping). I accept that; it is still the first thing a Raid-Helper guild will ask for.
4. Fixed. discord-bot-04-warning-3-failures-crop-390.png names #raid-signups and the three missing permissions in plain words.
5. Fixed. discord-msg-bis-slot-phone.png now has real icons, the "Best in slot" author line, the text line, two alternatives with deltas, and Share / Open on Forever Sixty.
6. Kept by ruling (it is the site's attendance chip, word for word). Accepted.
7. Kept by ruling. discord-msg-readiness-officer-phone.png now uses the board's own words; the unit-less "gear 12 below" line is gone. Accepted.
8. Fixed. discord-bot-12-permissions-no-webhook-picker-open-crop-390.png reads "Raid channel in Olympus Discord". #raid-signups is still preselected, but that comes from the capture's stubbed list: the real list comes from PostableChannels and drops the broken channel. Optional: re-capture with a list that matches the real one.
9. Deferred to the owner (role choice for hybrids). Accepted.
10. Fixed. "33 of 40" is the same on the post, /raid and the In reply; the Done post says "30 of 33"; the relative stamps match each frame's clock ("in 4 hours", "13 hours ago").
11. Fixed. "33 signed In or Late. 4 have something to fix."
12. Fixed. discord-msg-press-select-open-phone.png has "your main" in the description, and "In" is bold.
13. Kept by ruling. Accepted.
14. Deferred to the owner. Accepted.

### New

N1. MED, discord-msg-bis-slot-phone.png, discord-msg-bis-top-phone.png, discord-msg-char-upgrades-phone.png: the source line drops how you get the item. "Anachronos", "Pix Xizzix", "Illiyana Moonblaze", "Divino-matic Rod" and "Stormpike Guard" are bare names, so a raider can't tell a vendor from a quest from a mob, and the reputation level is missing. Wowhead writes "Vendor: Pix Xizzix", "Quest: The Divino-matic Rod" and "Stormpike Guard, Exalted". The BiS data already carries the kind ([vendor], [quest], [rep]). Fix: put the kind (and the reputation standing) in front of the source in ItemEmbed. "Stratholme: Baron Rivendare", "Leatherworking" and "to Zulmara · Onyxia" already read right.
N2. MED, for the BiS accuracy owner (not a bot defect), discord-msg-bis-top-phone.png and discord-msg-bis-slot-phone.png: "level 60 Phase 1 raid" lists no raid drops at all; every pick is vendor, crafted, reputation or dungeon gear. A raider reads "Phase 1 raid" as Molten Core and Onyxia loot. If the band means "raid buffs, launch gear", the label should say that. Separately, "Signet Ring of the Bronze Dragonflight · Anachronos" reads to every Classic player as the AQ40 ring. Id 234202 looks like Forever's own vendor version, but please confirm it can really be bought at launch. Discord makes this more visible: Share puts it in the channel.
N3. LOW, discord-msg-raid-phone.png: "You are In on Zulmara as DPS." is awkward and doesn't match the press replies. Fix: "Zulmara is In as DPS."
N4. LOW, discord-msg-report-member-phone.png and discord-msg-report-officer-phone.png: "4 items. The first 4:" says the same thing twice when every item is shown. Fix: "4 items:" when nothing is cut off.
N5. LOW, discord-msg-readiness-guild-tier-phone.png: "part of Raid guild ... move the guild to Raid guild" reads as if Raid guild were a guild's name. Fix: "the Raid guild plan".

Good, and worth saying: the report's Onyxia loot is right (Dragonstalker's Helm, Sapphiron Drape, Ring of Binding, Eskhandar's Collar). The shared report drops loot, deaths and consumables. /attendance refuses a member who looks up someone else, and /readiness refuses members without naming anyone. /char with no gear says so plainly. The unlinked and link replies are one sentence with one button. Nothing leaks, and nothing shames.

### Verdict: SHIP WITH FIXES

Fix N1 before ship. Send N2 to the BiS accuracy owner, since it is the site's data and the same problem is on the BiS page. N3 to N5 are wording only.
