# Forever Sixty community Discord

The server is "Forever Sixty" (guild 1548671158973628576), the one the site's invite
(`discord.gg/NZhBtPwvHt`) opens. The layout lives in [plan.json](plan.json), and
[apply.mjs](apply.mjs) applies it. Classes and specs come from `web/src/data/classes.json` and
`specs.json`, so the server always lists the same nine classes and the same specs as the site.

## Shape

One screen of channels for a new member. The class channels are opt-in through onboarding, and
the staff category is hidden.

| Category | Channels |
|---|---|
| Start Here | #welcome, #rules, #announcements (read-only), plus Discord's own Channels & Roles |
| Community | #general (the existing one, moved here), #lfg-and-recruitment, #screenshots, #off-topic |
| Classes | ⚔️・warrior, 🔨・paladin, 🏹・hunter, 🗡️・rogue, ✨・priest, ⚡・shaman, 🔮・mage, 😈・warlock, 🐾・druid. Each has one public thread per spec (Arms, Fury, Protection and so on; druid has Feral and Feral Bear). The hunter topic says Survival is melee |
| Theorycraft | #sims-and-bis, #logs |
| Forever Sixty | #site-feedback, #addon-support, #bug-reports (a forum), #premium-lounge (Premium only) |
| Voice | General (the existing one, moved here), Dungeon (5), Raid (40) |
| Staff | #mod-chat, #mod-log (hidden; AutoMod alerts and Discord's community updates) |

`#bug-reports` is a forum that requires a tag: Site, Addon, Companion, Discord bot or Data, plus
the moderator-only tags Confirmed, Fixed and Need info.

## Roles

From the top: Founder (gold), Moderator (crimson), then the nine class roles in the site's class
colours, each with its HD crest as the role icon, and Premium (light gold, hoisted, with the
Forever Sixty mark). Alliance and Horde roles sit between them, uncoloured. Class roles sit above Premium, so a name shows its class colour and crest,
and Premium members still get their own group in the member list.

There is **no automatic Premium role sync.** The API has no Discord role code. For now a
moderator adds the role by hand, and the welcome message says so.

## Onboarding

Discord's native onboarding (Community on):

1. **Which class?** Required, more than one allowed (for alts). Each option uses the class crest
   emoji, gives the class role and opens the class channel.
2. **Alliance or Horde?** Optional, both allowed. Each option gives the faction role, which LFG
   posts and moderators can read at a glance.
3. **What brings you here?** Optional: Leveling and dungeons, Raiding, Theorycraft and sims,
   Help with the site or addon. Each option opens #sims-and-bis, #logs or the support channels.

Default channels: Start Here, Community and the three public Forever Sixty channels. Forever has
no realms, so the LFG post format asks for the ruleset (Normal, PvP, Roleplay, Hardcore) instead.

There is no #bot-commands channel. The bot's /bis and /char only answer in a server linked to a
Forever Sixty guild, and this server is not one.

## Copy

The copy is in plan.json under `messages`, in the guide voice: no em dashes, no invented
features, and site links through `/go/discord` or `?golink=discord`.

- **#welcome:** what the server is for, three steps (rules, pick your class, say hello), what the
  site and addon do, where to report a bug, a class roster in crest emoji and role colours, and
  Premium with the manual role step. Pinned.
- **#rules:** seven rules (respect; no spam or self-promotion outside #lfg-and-recruitment; no
  real-money trading or account selling; follow the Blizzard and Discord ToS; no personal info;
  spoiler tags; moderators have the final say), then how enforcement works. Pinned.
- **#lfg-and-recruitment:** a pinned post format (class and spec, level, faction, ruleset, what
  you need) with two examples.
- **#announcements:** "Forever launches 4 November", with a Discord timestamp for 4 Nov 2026 15:00 PST, so
  every reader sees their own time and a countdown.

## Settings

- Community on. Rules channel #rules, system channel #general, updates and safety alerts to
  #mod-log.
- Verification: medium. Explicit content filter: all members. Default notifications: mentions only.
- AutoMod (alerts to #mod-log):
  - spam
  - mention spam (5 or more, with raid protection, plus a timeout)
  - invite links (exempt in #lfg-and-recruitment and for staff; our invite and the vanity allowed)
  - real-money-trading and power-leveling keywords
  - Discord's slur list
- Icon: the painted tile, the same image as the app icon. Banner: the tavern art with the painted
  seal and "Forever Sixty" (1920×1080, also the invite splash), allowed on this server's boost
  level. `build-assets.mjs` draws every image; none is placed by hand.
- Vanity URL `foreversixty`: set through the API if the server has the VANITY_URL feature,
  otherwise listed for the owner.

## Safety

apply.mjs never deletes a channel, role or member's message. The only messages it deletes are the "pinned a message" notices from its own pins. It finds everything by name (messages by the
bot's own embed title), creates only what is missing and edits the rest to match. Re-running it
creates no duplicates. It reports anything left over, such as the old empty "Text Channels" and
"Voice Channels" categories, for the owner to delete by hand.

## What the bot cannot do

Found on the first runs (3 Nov 2026):

- **Community** needs Administrator. The owner enabled it by hand (rules channel #rules, updates
  channel #mod-log), and the re-run then created #announcements and #bug-reports, posted the
  welcome and launch messages, turned onboarding on and set rules screening.
- **Role order:** the bot cannot move its own role, so the owner dragged it to the top. The
  invite also lacked **Pin Messages**, which the owner granted.
- **Vanity URL:** bots cannot set it, and it needs boost level 3. Set it in Server Settings >
  Vanity URL.
- **The safety alerts channel** is ignored by the API, so set it by hand if wanted.
- The old empty **"Text Channels" and "Voice Channels"** categories are left for the owner to
  delete.

## Reviews

- ux-designer: REVISE. Folded in: no timestamps in embed titles, the vanity on the AutoMod allow
  list, onboarding answers that open something, and a distinct Moderator colour. Not taken: the
  "timestamp is a day late" item. 1793833200 is 4 Nov 2026 23:00 UTC, which is 15:00 PST, so it
  is right.
- wow-player: REVISE. Folded in: the faction question, the LFG format, the 1.12 shaman claim
  rewritten, power-leveling in rule 3, no `*paypal*`, the 😈 warlock and the Survival note.
