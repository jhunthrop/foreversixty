# SEO checklist for launch

Audited 3 November 2026 against production (https://foreversixty.gg). I fetched every /guides,
/bis and /tiers URL in the live sitemap (70 pages, all 200) and read each page's `<title>`, meta
description, canonical, og:url, og:title, og:description, og:image and Twitter tags. Then I fetched
every og:image. Where a page's head comes from, in code: `web/src/layouts/Base.astro` and
`web/src/lib/seo.ts`. Share cards come from `web/src/pages/og/[...path].png.ts`.

## Sitemap, confirmed

- **Sitemap URL to submit: `https://foreversixty.gg/sitemap-index.xml`** (200, application/xml).
  It lists one child, `https://foreversixty.gg/sitemap-0.xml`, with 97 URLs.
- `https://foreversixty.gg/robots.txt` allows everything and names the same sitemap.
- `/sitemap.xml` is a 301 to `/sitemap-index.xml`, so a tool that guesses the default name still
  finds it.
- `www.foreversixty.gg` and `foreversixty.com` both 301 to `https://foreversixty.gg/`.

## Defects found

| # | Severity | Defect | Pages | Status |
|---|---|---|---|---|
| 1 | High | **og:image is a 404.** Base.astro points every page at `/og<path>.png`, but the share-card route only built cards for the home page, guides, BiS, the content pages and a few others. Every link shared to Discord, Reddit, X or Bluesky showed no picture. | /tiers, /tiers/tank, /tiers/healer, plus 18 more sitemap pages outside the audit: /sim and its 5 sub-pages, /premium, /premium/checkout, /guild, /logs, /rankings, /ratings, /reports/sample, /setup, /character, /connect, /account, /login | **Fixed on this branch** |
| 2 | High | **Duplicate titles and descriptions.** All three tier lists had the title "Tier list · Forever Sixty · WoW Forever" and the same description. | /tiers, /tiers/tank, /tiers/healer | **Fixed:** "DPS tier list", "Tank tier list", "Healer tier list", each with its own description |
| 3 | Medium | **Titles over 60 characters**, which Google cuts off. The suffix " · Forever Sixty · WoW Forever" is 30 characters on its own. | 11 BiS pages (e.g. "Hunter Beast Mastery leveling BiS · Forever Sixty · WoW Forever", 63) and /guides/hunter/beast-mastery (61) | **Fixed:** a title that would run past 60 keeps "WoW Forever" and drops "Forever Sixty", which og:site_name still carries |
| 4 | Medium | **Meta descriptions over 160 characters**, cut off in results. | /guides/mage/frost (183), /guides/rogue/combat (177), /guides/rogue/subtlety (170) | **Fixed:** shortened to 158, 154 and 152. The guide schema now rejects a description over 160, so the build catches the next one |
| 5 | Low | No `twitter:title` or `twitter:image` on any page. X and Bluesky fall back to og:title and og:image, so this only matters once defect 1 is live. | all 70 | Not changed; the fallback is fine |
| 6 | Low | BiS descriptions say "ranked by the simulator and verified with Top Gear". "Top Gear" means nothing to a searcher who has not used the sim. | 28 BiS spec pages | **Owner decision** (copy lives in `web/src/lib/bis/copy.ts`, `metaDescription`) |
| 7 | Low | Sign-in and utility pages are in the sitemap: /account, /login, /premium/checkout, /connect, /character. They are not pages anyone searches for. | 5 | **Owner decision:** add them to the sitemap filter in web/astro.config.mjs, and a `noindex` like /admin/links has |
| 8 | Low | /bis and /tiers sitemap entries have no `lastmod` (guides do), and BiS and tier pages carry no JSON-LD. | 32 | Not changed |

Checked and fine on all 70 pages: exactly one `<title>`, a canonical equal to the page URL,
og:url equal to the canonical, og:title equal to the title, og:description equal to the
description, `twitter:card` = `summary_large_image`, exactly one `<h1>`, no `noindex`, and every
other og:image returns a 1200 by 630 PNG.

## What changed in code (this branch)

- `web/src/lib/og-pages.ts` (new): one table of share cards for every static page. The OG route
  reads it. `og-pages.test.ts` walks `src/pages` and fails if a static page has no card, so a new
  page cannot ship a broken preview again.
- `web/src/lib/tiers/tier-copy.ts`: `pageTitle(role)` and `metaDescription(role)`, used by
  `TierPage.astro`.
- `web/src/lib/seo.ts`: `pageTitle` drops the brand from a title that would pass 60 characters.
- `web/src/content.config.ts`: guide descriptions are capped at 160 (`META_DESCRIPTION_MAX`).
- Three guide descriptions shortened.
- `web/public/ee556cf79f11bf298b87463b0558a4e4.txt`: the IndexNow key file (below).

None of this is live until main deploys. **Defects 1 to 4 stay on production until then.** Check
after the deploy:

```sh
for p in tiers tiers/tank tiers/healer sim premium guild; do
  curl -s -o /dev/null -w "%{http_code} /og/$p.png\n" "https://foreversixty.gg/og/$p.png"
done
curl -s https://foreversixty.gg/tiers/tank | grep -o '<title[^<]*</title>'
curl -s https://foreversixty.gg/ee556cf79f11bf298b87463b0558a4e4.txt
```

Expect six 200s, "Tank tier list · Forever Sixty · WoW Forever", and the key echoed back.

## Submit the sitemap

Do this after the deploy, so the crawlers see the fixed titles.

### Google Search Console

1. Go to https://search.google.com/search-console and sign in with the Google account that should
   own the property.
2. **Add property → Domain**, and enter `foreversixty.gg`. A Domain property covers http, https
   and www in one.
3. Google shows a TXT record (`google-site-verification=...`). In the Cloudflare dashboard, open
   foreversixty.gg → **DNS → Records → Add record**: Type `TXT`, Name `@`, Content the full string,
   TTL Auto. Save.
4. Back in Search Console, press **Verify**. It usually passes within minutes; DNS can take up to
   an hour.
5. Open **Sitemaps** in the left menu. Under "Add a new sitemap", type `sitemap-index.xml` (the
   box already shows `https://foreversixty.gg/`) and press **Submit**. The status should read
   "Success" and, after a while, "Discovered pages: 97".
6. Optional, for launch day: **URL inspection** → paste `https://foreversixty.gg/guides` →
   **Request indexing**. Repeat for `/bis` and `/tiers`. Google allows only a few of these a day,
   so spend them on the hub pages.

Google retired the anonymous "sitemap ping" URL in 2023; submitting in Search Console is the only
way.

### Bing Webmaster Tools

1. Go to https://www.bing.com/webmasters and sign in.
2. Choose **Import from Google Search Console** and approve access. Bing copies the verified site
   and its sitemap, so nothing more is needed. (If you would rather add it by hand: **Add site** →
   `https://foreversixty.gg` → verify with the DNS option, a CNAME Bing gives you, added in
   Cloudflare the same way as above.)
3. Open **Sitemaps → Submit sitemap**, enter `https://foreversixty.gg/sitemap-index.xml`, and
   press **Submit** if the import did not already list it.

### IndexNow (Bing, Yandex, Seznam, Naver and others)

IndexNow needs a key file at the site root, so the engines know the ping comes from the site's
owner. This branch adds it: `web/public/ee556cf79f11bf298b87463b0558a4e4.txt`, which holds the
key and nothing else. The key is public by design; it is not a secret.

After the deploy, and once the curl above echoes the key, send every launch page in one request:

```sh
curl -s -o /dev/null -w "%{http_code}\n" https://foreversixty.gg/sitemap-0.xml
curl -s https://foreversixty.gg/sitemap-0.xml \
  | grep -o '<loc>[^<]*</loc>' | sed 's/<[^>]*>//g' \
  | python3 -c 'import json,sys; print(json.dumps({"host":"foreversixty.gg","key":"ee556cf79f11bf298b87463b0558a4e4","keyLocation":"https://foreversixty.gg/ee556cf79f11bf298b87463b0558a4e4.txt","urlList":[l.strip() for l in sys.stdin if l.strip()]}))' \
  > /tmp/indexnow.json
curl -s -o /dev/null -w "%{http_code}\n" -X POST https://api.indexnow.org/indexnow \
  -H 'Content-Type: application/json; charset=utf-8' --data @/tmp/indexnow.json
```

A `200` or `202` means it was accepted. `403` means the key file is not live yet, and `422` means
a URL in the list is not on foreversixty.gg. One request per change is enough; never send the same
URLs on a loop. After launch, the pages worth a fresh ping are the ones a new client build changes:
the guides and the BiS lists.

## After launch

- In Search Console, check **Pages** a week after submitting: every guide, BiS and tier URL should
  be "Indexed". Look at any under "Crawled, currently not indexed" first.
- Paste one guide, one BiS and one tier URL into Discord and X once the deploy is live, to see the
  share card.
