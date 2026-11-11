# AI-slop review, lane C (unify-server)

**PASS WITH FIXES**

No blocking visual tell (1, 2, 5, 8, 12, 13, 15). The changes themselves read right: "8 pm CST", "4 min ago", "Today 6:40 pm" and "Sun 18 Oct" are how a player writes them, and the sign-in page now wears the site's seal, Marcellus and Open Sans. Companion and /b captures show no regressions; /b and /b-unavailable are unchanged as intended.

Fixes, none blocking:

1. `after/app-confirm-1440.png`, `app-confirm-390.png`: the backdrop is a bare blue-to-tan gradient with a visible band. The site and the companion show the same colours as painted sky and parchment, so here it reads as a decorative gradient (tell 2). Use the companion's background art or a flat site colour. Not a fail because the palette is the site's own.
2. `app-confirm-*`: "Sign in on Battle.net and the app picks it up from there." explains what the button already says (tells 19, 21). Drop it; the title and the "Continue to Battle.net" button carry the page. Keep "Didn't start this sign-in", it has a security reason.
3. `app-confirm-1440.png`: a narrow column with the lower 60% empty and a header with no nav; it feels unfinished next to the companion. Low priority.
4. `discord-post-full-desktop.png`, `discord-raid-answer-desktop.png`: "When 8:00 PM (in 4 hours)" sits next to "Guild time 8 pm CST", the same moment twice in two clock styles (tell 22). Discord's `<t:>` markup forces the first style, so fold the guild zone into one field, or show Guild time only when it differs from the reader's.
5. `discord-raid-answer-desktop.png`: "You have not signed up yet." restates "Signed up 1 of 40" and the In button (tell 21). The three "None yet" lines under 0 of N headers say what the count already says. The Buffs "Missing:" paragraph is real data but 20+ names in one block; trim to what the raid cares about.
6. `companion-logging-credited-en-US.png`: "Uploading your fights as they end", "Log written just now" and "up to date" explain what the green dot and stamps already show (tells 21, 22). Cut to "logging as Bow Jackzon" and keep the stamps.
7. `b-unavailable-1440.jpg`: "The build is safe and the link will work again soon. The build planner works in the meantime." is reassurance with no concrete noun (tells 18, 19). Unchanged by this lane; noted for the header-width owner.

Authentic, keep: the circular LX crest on the post, the seal on the sign-in header, real class icons and names, "Zul'Gurub / not on the site yet" row states, relative ages and locale-correct clocks in the companion (en-GB 24 h, en-US 12 h), the framed card used once per screen rather than stamped everywhere.
