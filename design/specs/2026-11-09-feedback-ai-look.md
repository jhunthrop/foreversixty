# Player feedback 2026-11-09: "looks super AI", blurry top-right text

Input to the whole-site design-direction lane. This lane fixes only the two plain bugs; no restyle,
no mocks.

Feedback, verbatim: "looks super ai Like the top right text is basically blurry ect, prot pally bis
is a wep you can only use like for a small period of time on a legendary staff quest". Three points:
the site reads as AI-made, the header's right side is blurry, and the Protection Paladin BiS weapon
(Task A, the data lane).

## Before captures

Production, signed out, 2x, in `design/reviews/2026-11-09-feedback-ai-look/before/`: home, /bis,
/tiers, /guides, /sim, /planner at 1440 and 390 (`<page>-<width>-dark.png`), each with a
`-header-right.png` crop, plus 1x crops of the home header (`home-1440-dark-1x-header-right.png`).
The site has one theme; the `prefers-color-scheme: light` captures were identical and were dropped.
`header-probe.json` holds the computed styles of every text node on the header's right side.

## Bug 1: the blurry top-right text (fixed)

No transform, filter, backdrop-filter or opacity on any header ancestor (header-probe.json). The
blur is the character selector's second line, "to see your characters": 12px, weight 400 Barlow in
`--color-muted` (#9a9484) under `-webkit-font-smoothing: antialiased`. At 1x on Windows that thins
to a grey smear next to the 14px semibold "Sign in" above it. The same rule draws the signed-in
class and level line.

Fix (web/src/styles/nav.css, `.csel-sub` and the phone `.csel-sub-short`): 13px, weight 500, in
`--color-nav` (#b9b3a4). After: `design/reviews/2026-11-09-feedback-ai-look/after-fix/`.

Still 12px in the bar, for the design lane to weigh: the Premium pill (12px semibold gold) and the
"/" key hint (12px mono, muted).

## Bug 2: "Forever right now" stuck on grey bars (not a code bug)

Measured on production (4 loads, 1440 and 390): `load` at about 0.3s, `networkidle` at about
0.85s, the GET /v1/feed request at 1.3 to 1.6s, rows rendered at 1.6 to 1.7s, every time. The feed
is not empty (the API returns BiS-refresh rows). A capture taken at `networkidle` lands before
the island hydrates: `client:idle-after-load` waits `load` plus 800 ms plus an idle callback for
signed-out visitors, an LCP guard (src/directives/idle-after-load.ts). A player sees the skeleton
for about 1.3s, then rows. Review captures of home must wait for `[data-testid=home-feed-rows]`.
Whether a signed-out visitor should get the build-time floor rows server-rendered instead of grey
bars is a design-lane call.

## What reads as AI-made (input to the design lane)

Ranked, from the captures. The reviewer pair (ux-designer, wow-player) still has to give its own
read on these captures before the design lane commits to a direction.

1. Glowing gold primary buttons. "Sign in with Battle.net" and "Look up" carry an outer gold glow
   and a gradient fill: the most common AI-template tell on the page.
2. The slogan hero. A big light serif "Play your class better." with a thin grey line under it says
   nothing a player can act on. The first screen is a slogan, a button and a class list.
3. The countdown card. A purple-to-navy gradient panel with a green status dot and an uppercase
   tracked kicker ("FIRST RAIDS OPEN · WED 9 DEC").
4. Uppercase letter-spaced micro-labels everywhere: CHARACTER NAME, REGION, RULESET, ADDON EXPORT
   OR BUILD CODE, BEST IN SLOT row tags, SHARE OF THE TOP SPEC'S DAMAGE. One label style repeated
   on every page reads as a generated design system.
5. Grey helper lines under every heading ("Pick your class.", "The published BiS gear, not yours",
   "Type the name, then pick the region and ruleset."). They restate the heading.
6. The class crests are glossy painted ringed art, not the game's class icons. Next to the real
   item icons in the feed and the countdown's raid icons they read as a different, generated hand.
   The one crest language is an owner tenet, so any change is a proposal for the owner.
7. The pill-shaped selector with a gold ring and a gold-ringed avatar placeholder at the top right,
   the first thing the feedback names.

What a veteran player trusts instead: real game icons at their native look, item names in quality
colours with Wowhead-style tooltips, numbers up front (the tier list is the strongest page for this
reason), plain sentence-case labels, flat buttons, and no hero slogan above the data.
