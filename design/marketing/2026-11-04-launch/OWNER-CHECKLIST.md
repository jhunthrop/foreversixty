# Owner checklist, launch day 4 November 2026

Only the steps that need your logins or your say. Everything else was done and checked in
production on the evening of 3 November (main 153b9b5a0):

- Every one of the 97 sitemap pages has an og:image that returns a PNG.
- The tier lists have their own titles: "DPS tier list", "Tank tier list", "Healer tier list".
- The IndexNow key file is served, and IndexNow accepted all 97 sitemap URLs (HTTP 202).
- All 36 /go links exist and each one 302s to its target (see [tracked-links.md](tracked-links.md)).

Go in order. Times are US Eastern.

## 1. Search engines (morning, 10 minutes)

**Google Search Console**, https://search.google.com/search-console

1. Add property → **Domain** → paste `foreversixty.gg`.
2. Copy the `google-site-verification=...` TXT value. Cloudflare → foreversixty.gg → DNS → Records →
   Add record: Type `TXT`, Name `@`, Content the value, TTL Auto. Save, then press **Verify**.
3. Sitemaps → "Add a new sitemap" → paste `sitemap-index.xml` → **Submit**.
4. URL inspection → paste each of these and press **Request indexing**:
   `https://foreversixty.gg/guides`, `https://foreversixty.gg/bis`, `https://foreversixty.gg/tiers`

**Bing Webmaster Tools**, https://www.bing.com/webmasters

1. **Import from Google Search Console** and approve.
2. Sitemaps → Submit sitemap → paste `https://foreversixty.gg/sitemap-index.xml` (if the import
   did not list it).

Full detail: [seo-checklist.md](seo-checklist.md), "Submit the sitemap". Do not ping IndexNow again
today; it already has every URL.

## 2. Decide the r/wowforever AI disclosure (before any Reddit post)

r/wowforever rule 12 requires it; without it the post is removed. In
[reddit-launch-post.md](reddit-launch-post.md) replace each `[AI DISCLOSURE, owner's wording, see
below]` (three places: Version A, Version B, the r/classicwow variant) with your wording. The draft,
if every part is true for you:

```
We used AI coding tools to build the site and addon. Every guide claim carries its source, and players check them.
```

Also read the live sidebar of r/wowforever once (rules 2 and 8 were not captured).

## 3. Reddit posts

| When | Where | Title (paste) | Body |
|---|---|---|---|
| 4 Nov, about 10:00, once realms are up | r/wowforever | `[I Made] What changed from 1.12 for every class in Forever, one cheat sheet` | Version A body in reddit-launch-post.md, from "Forever launches today." to "...fix the guide today." |
| 11 Nov or later, same time of day | r/wowforever | `[I Made] Leveling BiS for every spec at 20, 30, 40, 50 and 60, and how to read it` | Version B body |
| Only once the posting account has real r/classicwow history | r/classicwow, flair Forever if offered | `Coming from Era: the Forever changes that will catch your main out` | The r/classicwow variant body |

Posting limits:

- r/wowforever rule 10: **one promotional post per 7 days.** Version A today, Version B no earlier
  than 11 November. Answers in other threads do not count as posts, but must not be ads.
- r/classicwow: about 1 promotional post per 9 normal ones. A fresh account gets removed.
- Each post carries **one** /go link (already in the bodies). Never ask for upvotes.
- Answers: use [reddit-answers.md](reddit-answers.md). Answer first, link second, say "we keep the
  list here", and never paste the same answer twice. Rewrite the first line per thread. Answers 1
  and 2 match production tonight; if the BiS refresh lands first, open the page and check the
  item names before posting.
- No guild-tool mentions on Reddit at all (r/wowforever rule 13, r/classicwow rule 8).
- Drop "we will fix the guide today" from Version A if nobody watches the thread.

## 4. CurseForge and Wago listings

The gallery screenshots stay as they are: the new plates are held until the in-game recaptures
([design/listing/RECAPTURE.md](../../listing/RECAPTURE.md)). Paste
text only. All blocks are in [listing-copy.md](listing-copy.md).

**Forever Sixty addon, CurseForge** (project → Edit → Description)

1. Summary field, paste:
   `Play your class better in Forever. Your next upgrade, your next talent and your Forever Score on the game's own frames. No setup.`
2. Description: replace from `## Play your class better, in game` down to (not including)
   `### Install in two minutes` with the Markdown block under "Long description: new top section".
3. Replace the last line with the "Footer line" block.
4. Check the "What's new" build number (1.60.1.70755, addon 2.0.23) is still current.

**Forever Sixty addon, Wago**: the same three pastes, with every `https://foreversixty.gg/go/curseforge`
changed to `https://foreversixty.gg/go/wago`. The tagline is the same summary line.

**Forever Sixty Data, both sites**

1. Short description, paste:
   `Forever Scores and guild standings from foreversixty.gg, updated every night. Install it with the Forever Sixty addon.`
2. Add the quoted block under "Long description: one line at the top" above `## The Forever Score, in game`
   (`/go/wago` on Wago).

Afterwards, copy the same changes into design/listing/curseforge-description.md and
data-addon-description.md so the repo matches the dashboards.

## 5. Creator DMs

Template and list: [creator-dm.md](creator-dm.md). Every `yt-<name>` link already exists.

1. Decide first: Premium length (open-ended, a year, or to the end of Phase 1); what "early sim
   data" means; nothing asked in return.
2. Send through each creator's public business contact only (YouTube About → business email, a
   Twitch panel, their own Discord). Start with rows 1 to 12. Check rows 17 to 19 for Forever videos
   first. Note the NOHITJEROME caution.
3. Replace every bracket. In the link line, paste the slug from the table, for example
   `foreversixty.gg/go/yt-kargoz`. Never send two creators the same first line.
4. When a creator accepts, they sign in once with Battle.net, then run from a checkout at the
   **deployed api sha** (the CLI migrates at startup):
   `api grant --user <battletag> --plan premium --note "creator: <name>"` (add `--until 2027-11-04`
   for a year).

## 6. Guild offer: your decision, then the grant

Proposal in [guild-outreach.md](guild-outreach.md), "Proposed first-50-guilds offer". Decide: yes or
no, how many (50 or fewer), Guild or Raid guild, and the claim step (leader replies with the guild
page link after posting a raid).

- **If no:** delete the `[IF APPROVED: ...]` line from the leader message and send it as it is.
- **If yes:** keep that line without its brackets. For each guild that claims, look up its guild id,
  then from a checkout at the deployed api sha run:
  `api grant --guild <id> --plan guild --until 2027-01-06 --note "first 50 guilds"`
- Send only on the official forums' recruitment section or Discords that allow it, never Reddit.
  One message per guild. Keep your own sheet of who you contacted.

## 7. X and Bluesky

Five posts in [social-posts.md](social-posts.md); paste the quoted text as it is. On Bluesky change
`x-` to `bsky-` in the link (`bsky-launch`, `bsky-bis`, `bsky-changes`, `bsky-tiers`, `bsky-guild`).

| When | Post | Link in the text |
|---|---|---|
| Morning | 1, Launch | `foreversixty.gg/go/x-launch` |
| Midday | 2, Leveling BiS | `foreversixty.gg/go/x-bis` |
| Afternoon | 3, What changed | `foreversixty.gg/go/x-changes` |
| Evening | 4, Tier list (check the ranks on the day) | `foreversixty.gg/go/x-tiers` |
| Late, after raids form | 5, Guilds | `foreversixty.gg/go/x-guild` |

Images: your own screenshots of the pages named under each post (1440 wide, dark mode, cropped to
1600 by 900), never AI images. Keep "updated every night" in post 4 only while the nightly tier run
is green.

## 8. End of day

Open https://foreversixty.gg/admin/links and look at clicks and sign-ups per link.
