# Guild outreach

Drafts for the owner to send by hand. Facts are from web/src/lib/billing/plans.ts and
plans-page-copy.ts as of 3 November 2026.

## Where these can go

- **Not as Reddit comments on other people's posts.** r/wowforever rule 13 and r/classicwow rule 8
  keep guild talk in the weekly megathread, and r/wowforever says "do not advertise in unrelated
  comment sections". A reply to a guild's recruitment post on Reddit counts as that.
- **Fine:** the official forums' guild recruitment section, a guild's own Discord when an officer
  has invited you in, and server community Discords that have a tools or addons channel. Read each
  server's rules first.
- One message per guild. No follow-up unless they answer.

## What a guild gets

| Plan | List price | Founder price (to 6 Jan 2027) | What it adds |
|---|---|---|---|
| Free | $0 | $0 | Guides, BiS, planner, browser sims, logs and parses, rankings, the tier list, the addon's guild tab |
| Guild | $15 a month or $150 a year | $12 a month or $120 a year | Raids posted on the site and in Discord with one-tap sign-ups, attendance and a report for every raid night from its log, the Discord bot (raid posts with buttons, /bis and /char), 1 Premium seat, up to 25 verified members |
| Raid guild | $25 a month or $250 a year | $20 a month or $200 a year | Everything in Guild, plus loot council, the readiness board, 3 Premium seats and any number of members |

Founder prices stay for as long as the guild stays subscribed. Every guild gets 7 days free, once.
Members never pay to sign up. If a plan ends, raids, attendance, loot and reports stay readable,
and officers can export attendance and loot as CSV at any time.

## Proposed first-50-guilds offer (NEEDS OWNER APPROVAL)

> **Status: proposal only. Do not send any message that mentions it until the owner approves the
> offer, its end date and how it is granted.**

**The offer:** the first 50 guilds that post a raid on Forever Sixty get **the Guild plan free
until 6 January 2027**, the day founder prices end. After that they can keep Guild at the founder
price ($12 a month or $120 a year) for as long as they stay subscribed, or drop to free and keep
everything they logged.

Why this shape:

- It ends on a date already public on /premium, so there is one deadline to explain.
- "Post a raid" is a real step. A guild that does it has already used the tool, and it is easy to
  count.
- It costs nothing until a guild pays, and the founder price is the one already on offer.

What the owner must decide:

1. **Yes or no, and the number.** 50, or fewer.
2. **Guild or Raid guild.** The draft gives Guild. Raid guild free would also hand out 3 Premium
   seats per guild.
3. **How it is granted.** There is no coupon flow on the site. The grant CLI already does this:
   `api grant --guild <id> --plan guild --until 2027-01-06 --note "first 50 guilds"`, run from the
   deployed sha only, since the grant CLI migrates at startup.
4. **How a guild claims it.** The draft has the guild leader reply with their guild's page link
   after posting a raid. The owner runs the grant and replies once.
5. **What happens at guild 51.** The draft says "while places last" and nothing more.

## Message to a guild leader

Use the first line to show you know the guild. The bracketed parts change per guild.

> Hi [name], I saw [guild] is recruiting for [raid night / progression goal] on [realm].
>
> I run Forever Sixty (foreversixty.gg), a free class helper for Forever: guides, leveling BiS and
> a sim. We also built guild tools, and I think they would save your officers time:
>
> - Post a raid on the site and in your Discord. Members sign up with one tap.
> - Attendance and a report for every raid night, straight from the log.
> - A Discord bot for raid posts with buttons, plus /bis and /char lookups.
> - The addon's guild tab shows each member's level, spec, gear gap and score in game.
>
> [IF APPROVED: The first 50 guilds that post a raid get the Guild plan free until 6 January.]
> Otherwise there is a 7 day free trial, and founder pricing is $12 a month until 6 January.
>
> Here is the guild page if you want a look: foreversixty.gg/go/guild-dm
>
> No reply needed if it is not for you. Good luck at launch.

Word count: about 140. Keep it under 160 when you personalise it.

## Reply on a recruitment thread (forums or a community Discord, never Reddit)

Only where the channel allows tool mentions. Reply to the thread, not the poster's DMs.

> Good luck with the roster. If you want sign-ups and attendance handled for you, we built raid
> sign-ups, attendance from logs and a Discord bot for Forever guilds at foreversixty.gg/go/guild-reply.
> There is a free week to try it. (Disclosure: it is ours.)

## Follow-through

- Write down each guild contacted (name, realm, date, where) in the owner's own sheet, so nobody
  gets a second message.
- If a guild says no, do not contact it again.
