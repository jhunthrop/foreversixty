VERDICT: SHIP (round 2; round 1 was FIX)

Page job: tell a raid leader that this addon, plus the guild's page on the site, replaces the sign-up bot and the readiness spreadsheet, and tell a raider why to install it before Thursday. Reviewed: curseforge-description.md (table, captions, Description), plates 01 and 02 in full, 03 to 11 by caption, against forever-gallery code (the main checkout lacks the guild addon files).

## Claims checked against code

- Discord sign-up: TRUE. api/internal/discordbot/cmd_signup.go is /signup status/raid/character; 98ed6f843 adds picking a character with no site account, at most 3 per account, marked with the Discord name for officers.
- "chest, back, boots and weapon are enchanted": TRUE. api/internal/guilds/readiness_core.go enchantableSlots = main_hand, chest, back, feet. The check is "carries any enchant", not the right one (its own comment says so). The listing never says "the right enchant", so it is true; do not let anyone add that.
- Officer tab: TRUE. Remind sends one SendChatMessage to GUILD (Officer.lua remind). Send to guild, the Give twice flow, whisper and shift-invite are all in OfficerTab.lua and Officer.lua. Readiness rows in plate 02 match the code's checks.
- Log rotation check: TRUE for DPS and tanks, healers and specs with no rotation file get nothing (api/internal/coach/f8_adherence.go). "(DPS and tanks)" is exact.
- Phone app "in beta": TRUE, web/src/pages/app.astro is a beta page with a tester form.
- Plate numbers agree between 01 and 02 (12 of 40, 11 In, 16 not signed up, Velisandre 9th of 15 at 1184, 2 missing enchants, no fire prot).
- Plate 02 note: the roster is "16 of 31 shown". Fine, but the readiness "10 of 12 addon live" only counts signed-up raiders; the listing's "who is ready" is therefore "of the people running the addon".

## Findings

1. BLOCKER. The listing never tells the raid leader what it takes to get the guild on it. Description, "### Guild": "Raids, sign-ups and jobs come from the guild's page on the site through the companion app; your answers go back the same way." and "### The companion app, optional": "...the raid parts of Guild and Officer need it." Optional is false for the officer: without the companion (Windows and Mac only) the Officer tab has no raids, sign-ups or roster rows. An officer installs it, sees an empty tab and uninstalls. Rewrite the second line of "### Guild" to: "Setup for the guild, once: an officer claims the guild on foreversixty.gg and posts the raid; everyone installs this addon and the companion app (Windows and Mac) so their sign-up and gear reach the roster. Raiders without the addon still sign up from Discord or the site and show on the roster as 'without addon'." And in the companion heading, drop "optional": "### The companion app (needed for the guild tabs)". Also say the cost in one line if the guild features are paid (not in the listing today; needs an owner decision, the pricing file is the source).

2. BLOCKER. The lead does not sound like a raid leader and buries the raid. Description first paragraph and summary: "Raid sign-ups and guild roster for Forever. Gear upgrades are measured by our sim for your spec and level, on Forever's rules." A roster is a list; the officer's win is "who is not ready at 19:30". The first bullet is still DPS gain, the raid is bullet two. Rewrite summary: "Run Forever raid night: sign-ups from Discord or in game, who has flasks, enchants and attunement, and which upgrade is worth it, sim-measured for your spec." Reorder bullets: sign-ups, then roster readiness, then upgrades, then log check. Rewrite first bullet pair: "- **Sign-ups where your guild already is**: In, Late, Bench or Out from Discord, the site, the addon or the phone app (beta). Officers see 11 In, 16 not signed up and one click pings them in guild chat." / "- **Know who is ready before the pull**: flasks, fire protection, attunement, and enchants on chest, back, boots and weapon, for every raider running the addon."

3. SHOULD. "shows who is ready" over-reaches slightly: readiness covers only raiders running the addon (plate 02 "Addon live 10 of 12"). Line: "Forever Sixty runs your guild's raid sign-ups and shows who is ready before the pull." Rewrite: "...and shows which raiders running the addon are ready before the pull, and who is not running it."

4. SHOULD. Officer value is hidden in prose. Plate 02 has Tonight's readiness, Push a build and Master loot, but the Description's Officer paragraph is one 150-word block. Split into three short lines: "Roster: sign-up, attendance, gear gap in DPS, missing enchants, item level." / "Readiness: addon, build, enchants, attunement, flask, fire protection, jobs open, last ready check." / "Loot: as master looter, ranked by what the item gains each raider, wins this week and attendance; Give on a second click." Lead with the one-click Remind, since the caption already does.

5. SHOULD. Raider reason to install is thin and not first. A raider opening plate 01 should read "when is the raid, am I in, what am I missing" in the Description's first screen. Add under the bullets: "Raiders: open Guild before raid for your sign-up, your job for the fight ('Phase 2 ranged, west side') and what you are missing (Greater Fire Protection Potion)." The job text is real (plate 01, jobs come from the site).

6. SHOULD. Discord is under-sold for a leader who lives there. Verified commands exist (cmd_raid, cmd_readiness, cmd_attendance, cmd_signup) but the listing says only "from Discord". Add: "The Discord bot posts the raid, takes /signup (a raider with no site account just picks their character) and answers readiness and attendance." Check each against the bot's help before pasting; cmd_help.go lists them.

7. SHOULD. Named add-ons. Description, "### Bags, paperdoll...": "The arrow also shows in Baganator, Bagnon, AdiBags, ArkInventory and BetterBags." Named bag addons are integrations, not rivals, but the brief says name no rival addon. Rewrite: "The arrow also shows in the popular bag addons; `/fs diag` names yours and says whether it shows." The pasted list can stay in the GitHub README.

8. SHOULD. Plate 01 caption and its name mismatch: caption "check your bags for what the fight asks" while the plate is titled by an officer's window. Fine for 01, but the raid-first order is wrong: plate 02 (officer) is the buyer's screen and should be second to featured plate 01 only if 01 shows the raid. It does; keep order, but star 01 and make 02's caption lead with the outcome: "**Officer.** Who is signed up, who is missing a flask, an enchant or an attunement, and one click to remind everyone not signed up."

9. NICE. Name "Forever Sixty: Guild Raids, Sim and BiS": "BiS" appears in no guild claim and "Sim" is the pitch's second half; fine. Keep.

10. NICE. What is not in the listing and not claimable: attendance history by night, loot award history and fairness per raider, recruiting and trial status, performance versus spec median per raider on the guild page. Do not add them to the listing without checking the addon or site; some exist on the site (api guild_raid_attendance, loot_awards tables), but the listing is addon-first. One line that points at the site is safe: "Attendance and loot history live on your guild's page on foreversixty.gg." Verify the page shows both before shipping.

11. NICE. No em dashes found, no rival addon beyond finding 7, no real person named (plate rosters are invented names; check none matches a real streamer before upload: Brannoc, Ysolde, Torvald, Kethra, Mirelle, Velisandre are fine).

## Does it make me install it

As a raider: partly. Plate 01 is strong (countdown, job, prep, missing potion). As a raid leader: not yet, because I cannot tell what I must do to move my guild, and the companion requirement is hidden behind "optional". Fix 1 and 2, then ship; 3 to 7 are one edit pass.

## Round 2

VERDICT: SHIP

Re-read curseforge-description.md. Checked against the code: the guild page's own heading and button are "Claim this guild" (web/src/lib/guild/copy.ts); the bot's command descriptions match the Discord line (copygen DiscordCommandsRaid, Signup, Attendance, Readiness); the readiness board's four enchant slots are weapon, chest, back, boots; the companion heading now reads "(needed for the guild tabs)". No em dashes, no named addon, no real person.

1. Round 1 finding 1 (BLOCKER): resolved. "Getting the guild on it" in Guild gives the officer the steps and tells raiders what they install; raiders without the addon can still sign up from Discord or the site. The opening paragraph and the companion heading say the guild parts in game need the companion.
2. Finding 2 (BLOCKER): resolved on the bullets, which now run sign-ups with Remind, readiness, upgrades, log check. The summary table line ("Raid sign-ups and guild roster for Forever...") is still roster-first. The owner chose it, so this is a NICE, not a blocker. If the owner revisits: "Run Forever raid night: sign-ups from Discord or in game, who has flasks, enchants and attunement, and which upgrade is worth it, sim-measured for your spec."
3. Finding 3: resolved. The enchant bullet names the site's readiness board, and Officer in game adds flasks, protection potions and attunement.
4. Finding 4: resolved. Officer leads with Remind, then Roster, Readiness, Push a build, Loot.
5. Finding 5: resolved. The "Raiders:" line is there.
6. Finding 6: resolved. The Discord line matches the bot's own descriptions, with /readiness marked officers only.
7. Finding 7: resolved. Bag addon names are gone.
8. Finding 8: resolved. Caption 2 uses the rewrite.

Open, not blocking:
- NICE: pricing is absent. A guild leader will ask what it costs; the owner's call, reported by the other session.
- NICE: "Raiders without them can still sign up" is true, but those raiders show in the officer roster as without the addon, so the officer cannot see their gear or enchants. The readiness-board bullet already limits the claim to raiders who have the addon in game; leave it.

For a raid leader: the setup path is now clear and every guild claim is true, so I would switch the sign-ups to it. A raider has a reason to open Guild before the pull. For the guild's page itself, the listing keeps to what the addon and bot do.
