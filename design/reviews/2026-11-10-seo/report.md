# SEO: the 30 Oct 2026 impression cliff

Investigated 2026-11-10 from Search Console (sc-domain:foreversixty.gg), URL inspections, a crawl of all 234 sitemap URLs, deploy history and Cloudflare audit logs.

## Most likely cause
A site-wide algorithmic re-score by Google on 29–30 Oct, not a technical fault.

1. **One-day drop:** impressions fell from 238 on 29 Oct to 2 on 30 Oct (PT), across every page type, country and device at once.
2. **Still indexed, ranked lower:** 62 of 102 inspected URLs are "Submitted and indexed" (home, /sim, top guides; crawled as late as 9 Nov). Positions fell: "sixtyupgrades forever" 2.8 → 6.3, /sim 7.0 → 27.2.
3. **Google is rationing the site:** 8 of 28 spec guides are indexed. Formerly ranking pages moved to "Crawled, not indexed" (/bis/paladin/retribution, /guides/rogue/subtlety). There are no crawl errors.
4. **External volatility matches:** unconfirmed volatility on 16, 20–21, 24, 27 and 29 Oct (Semrush 9.5/10 US desktop). The September spam update had finished on 8 Oct.
5. **Fragile demand:** of 703 query impressions on 15–29 Oct, 49% were competitor-intent ("classify forever", "sixtyupgrades forever"), 37% were "wow forever sim(s)" and 3% were "forever sixty". The 29 Oct peak was mostly /sources ranking #4 for "classify forever".
6. **Ruled out:**
   - head, robots, canonicals, noindex and Worker routes (unchanged 28–30 Oct);
   - the sitemap (all 234 URLs return 200 to Googlebot, with no redirects, noindex or duplicate titles);
   - server HTML on content pages;
   - other hosts;
   - blocklists (Safe Browsing, Spamhaus DBL, SURBL, URIBL);
   - DNS and Cloudflare (zone unchanged since 13 Sep);
   - scraped copies.
7. **Possible contributor:** the 29 Oct alignment stage 2 deploy (home, /bis, guides index, /premium, /setup) landed hours earlier. Unchanged guides and tiers fell the same day, so at most it added to the drop.

## Other findings
- **BiS band pages are not duplicates:** they share 6–20% of their text (5-word shingles). /60/pre-raid duplicates /60 and already canonicalises to it.
- **Thin server HTML** is limited to the tool pages: /sim/* (28 words), /planner (24) and /rankings (18).
- **The workers.dev twin** canonicalises to foreversixty.gg and isn't indexed.

## Fixes (lane/seo-fix, dcead254c)
1. The sitemap lists only pages that are their own canonical (web/src/lib/sitemap-canonical.mjs).
2. pageTitle no longer doubles the brand.
3. The workers.dev twin is noindex (web/public/_headers plus markPreviewNoindex in web/src/worker.ts).

## Owner decisions (2026-11-10: approved the fixes)
- **Links and brand demand:** CurseForge and Wago listings, the Forever Discords and subreddits, creator videos.
- **Server-render the tool pages' default state,** /sim first.
- **Keep home, /sim, guides and tiers stable for a few weeks.**
- **Request re-crawls in the Search Console UI:** the 20 unindexed spec guides, then /sim and home.
- **BiS bands:** revisit on 8 Dec. If most are still unindexed, fold the 30-39, 40-49 and 50-59 bands into the spec page.
- **Cleanups:**
  - the hidden second h1 on home;
  - the invisible future-dated "Updated" placeholder on /rankings;
  - 14 BiS titles of 66–69 characters.
