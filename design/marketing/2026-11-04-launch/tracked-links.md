# Tracked links for the launch kit

## How /go links work

- `foreversixty.gg/go/<slug>` is answered by the Worker (web/src/lib/golinks/redirect.ts). It
  sends a 302 to the slug's target with `?golink=<slug>` added, then counts the click: one counter
  per slug per day, people and bots counted apart, nothing stored about the visitor.
- An unknown slug sends the visitor to the home page, uncounted, so **create every slug before you
  post it.**
- On landing, the page keeps the slug for an hour. If the visitor signs up in that time, the
  account records it once as its first link (`users.first_touch_slug`). That is the funnel on
  /admin/links.
- A plain page URL with `?golink=<slug>` added (the Reddit answers use this) records sign-ups the
  same way, but counts no click.

**Where slugs are made:** https://foreversixty.gg/admin/links, signed in as the owner. Each needs
a slug (lowercase letters, digits and dashes, 1 to 40 characters, starting with a letter or
digit), a target (a path on foreversixty.gg only, starting with one `/`, up to 200 characters) and a
label (1 to 80 characters). Migration api/internal/db/migrations/0066_go_links.up.sql seeded
eleven links; everything else is made on the admin page. Do not add a migration for these.

## Already there (seeded by 0066)

| Slug | Target | Used in this kit |
|---|---|---|
| `curseforge` | `/` | listing-copy.md, the CurseForge descriptions |
| `wago` | `/` | listing-copy.md, the Wago descriptions |
| `reddit` | `/` | no |
| `reddit-classicwow` | `/` | no (the variant has its own slug, below) |
| `reddit-wowforever` | `/` | no (each post has its own slug, below) |
| `reddit-wowclassic` | `/` | no |
| `reddit-guides` | `/guides` | no |
| `reddit-sim` | `/sim` | no |
| `discord` | `/` | no |
| `youtube` | `/` | no |
| `twitter` | `/` | no |

Production may differ if the owner has edited these; check the table on /admin/links.

## To create on /admin/links (36 links)

### Reddit (4)

| Slug | Target | Label | Used in |
|---|---|---|---|
| `rwf-changes` | `/guides` | Reddit r/wowforever: what changed, launch post | reddit-launch-post.md, Version A |
| `rwf-bis` | `/bis` | Reddit r/wowforever: leveling BiS post | reddit-launch-post.md, Version B |
| `rcw-changes` | `/guides` | Reddit r/classicwow: coming from Era | reddit-launch-post.md, variant |
| `reddit-answer` | `/` | Reddit answers (signups only, no clicks) | reddit-answers.md, as `?golink=` |

### Guild outreach (2)

| Slug | Target | Label | Used in |
|---|---|---|---|
| `guild-dm` | `/guild` | Guild leader DM | guild-outreach.md |
| `guild-reply` | `/guild` | Guild recruitment-thread reply | guild-outreach.md |

### Social (10)

| Slug | Target | Label |
|---|---|---|
| `x-launch` | `/` | X: launch post |
| `x-bis` | `/bis` | X: leveling BiS post |
| `x-changes` | `/guides` | X: what changed post |
| `x-tiers` | `/tiers` | X: tier list post |
| `x-guild` | `/guild` | X: guild tools post |
| `bsky-launch` | `/` | Bluesky: launch post |
| `bsky-bis` | `/bis` | Bluesky: leveling BiS post |
| `bsky-changes` | `/guides` | Bluesky: what changed post |
| `bsky-tiers` | `/tiers` | Bluesky: tier list post |
| `bsky-guild` | `/guild` | Bluesky: guild tools post |

### Creators (20, create each just before its DM goes out)

| Slug | Target | Label |
|---|---|---|
| `yt-kargoz` | `/guides/hunter/beast-mastery` | Creator: Kargoz |
| `yt-nicktew` | `/guides/shaman/enhancement` | Creator: NickTew |
| `yt-wille` | `/guides` | Creator: WillE |
| `yt-crix` | `/guides/warlock` | Creator: Crix Forever |
| `yt-tiqqle` | `/guides/shaman` | Creator: Tiqqle |
| `yt-bajheera` | `/guides/warrior/protection` | Creator: Bajheera |
| `yt-venruki` | `/guides/mage` | Creator: Venruki |
| `yt-hammerdance` | `/guides` | Creator: hammerdance |
| `yt-signsofkelani` | `/guides` | Creator: SignsOfKelani |
| `yt-supatease` | `/guides` | Creator: Supatease |
| `yt-frostadamus` | `/bis` | Creator: Frostadamus |
| `yt-icosiol` | `/bis` | Creator: Icosiol |
| `yt-savix` | `/planner` | Creator: Savix |
| `yt-tipsout` | `/planner` | Creator: Tips Out |
| `yt-shobek` | `/` | Creator: Shobek |
| `yt-nohitjerome` | `/tiers` | Creator: NOHITJEROME |
| `yt-aeghis` | `/guides/mage` | Creator: Aeghis |
| `yt-hydramist` | `/guides/priest` | Creator: HydraMist |
| `yt-jokerd` | `/guides` | Creator: Jokerd |
| `yt-taliesin` | `/` | Creator: Taliesin and Evitel |

Every target above is a page in the live sitemap (checked 3 November 2026). Every slug matches the
slug rule.

## Total

- 2 existing links used: `curseforge`, `wago`.
- 16 to create before launch posts: the Reddit, guild and social tables.
- 20 creator links to create one by one as DMs go out.
