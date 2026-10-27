# Launch pricing: player review (wow-player persona), 2026-10-27

Reviewer: veteran Classic raider and guild officer. Read every state on the captures in this folder (1440 crops and full pages, 360/390 phone pages). Spec: design/specs/2026-10-27-launch-pricing.md.

## Verdict: BLOCK

The /premium page is good. I understood what I get and what it costs in about five seconds, there is nothing hidden, and the founder deal is plain. The rules table ("How the guild plans work") is what I'd want to read before putting a guild plan on my card. Behind the paywall it falls apart. A Founder paying for Premium is told by the account page that they don't have Premium. A Raid guild sold with 3 seats shows "3 of 5 given". A trial or plan bought inside the founder window is charged list price. Officers are told "Only the guild master can subscribe" on a plan another officer is being offered. Any one of these ends up screenshotted in guild chat with "lol, they charged me and it says I'm free".

## Findings

1. HIGH, launch-pricing-12-founder-crop-1440.png (also 12-seat-holder, 12-offer, 12-accepted): the panel says "Premium · renews 19 Dec 2026 · $40 a year FOUNDER" and right under it "Server sims: Browser lane only", "History charts: Sparkline only", "Discord notifications: Off · Premium", "Supporter mark: Not shown", plus a "See plans" upsell. I'm paying and the site says I'm on the free tier. Fix: on an own sub or a seat, the rows show the Premium values (server sims on, full charts, notifications on, mark shown) and the See plans line is hidden.
2. HIGH, launch-pricing-09-seats-crop-1440.png / -390 / picker-payer: "3 of 5 given" on a Raid guild that /premium sells with 3 seats. Five rows, every holder is "Kraggor", the waiting row is titled Kraggor with "Mirelle" under it, and Kraggor both holds seats and "left the guild". I can't tell who has my seats. Fix: total = the plan's seat count (3; returned rows replace an empty slot, not add one). The waiting row's title is the offeree. The fixture uses distinct names and classes.
3. HIGH, launch-pricing-09-seats-crop-1440.png: the waiting (offered) row says "Can move from 1 Dec" and shows a disabled Move seat. The rules say a withdrawn offer starts no lock and the lock starts at acceptance. Fix: an offered row shows only the offer sentence and Withdraw, with no lock date and no Move seat.
4. HIGH, launch-pricing-07-trial-crop-1440.png, 10-*, 11-read-only-officer, 08-confirm-ready: on 20 Nov, inside the founder window, the guild shows "first charge 27 Nov 2026: $15" / "$15 a month" with no Founder mark, and the upgrade quotes "$25 a month". The page promised $12/$20 until 6 January. Feels like a bait and switch. Fix: fixture and plan line use the founder amount and show Founder for anything bought in the window, and the upgrade quotes the founder Raid guild price.
5. HIGH, launch-pricing-11-read-only-officer-crop-1440.png and 10-receiver-offer-crop-1440.png: "Only the guild master can subscribe the guild, release the claim or change billing." This contradicts the page's own rules (subscriber or guild leader), and it sits under an offer telling this same officer to take over billing. Fix: replace it with "The subscriber or the guild leader manages billing and seats." and drop it on the receiver's view.
6. HIGH, launch-pricing-11-read-only-officer, 10-receiver-offer, 12-seat-holder: "Billed by Fixture#1234" / "Given by Fixture#1234". Other officers see a BattleTag, and the receiver's offer comes "from Thornwick" while billing says someone else. Fix: show the subscriber's main character name (as spec: "Billed by Kraggor"), never a BattleTag, and make sure offerer and biller agree.
7. MED, launch-pricing-08-confirm-ready-crop-1440.png: plan "renews 30 Nov 2026", but the upgrade says "From 1 Dec 2026 it is $25 a month". An off-by-one date on a money sentence makes me not press the button. Fix: the upgrade uses the same renewal date the plan line shows.
8. MED, launch-pricing-10-send-confirm-crop-1440.png / 10-receiver-offer: as a GM handing the guild over, I'm not told what happens to the seats (do Kraggor and Mirelle keep Premium?), the records, or whether I keep officer access. The receiver isn't told they get the seats to manage. Fix: add one line to both. "Seats, raids and records stay with the guild; {Name} manages seats from then."
9. MED, launch-pricing-05-two-guilds-1440.png: on the checkout chooser, guild crests render as empty gold rings, and there's no price or interval on the page. I picked "$200 a year" two clicks ago and now nothing confirms it. Fix: render the guild crest, and add a line under "Which guild?" with the chosen plan and price (e.g. "Raid guild · $200 a year").
10. MED, launch-pricing-05-*, 09-seats-1440.png: the header reads "Sign in / or paste an export" on checkout and on a guild settings page I can only reach signed in. Fix: run the captures with the signed-in header so the review shows what a payer sees (and check the real header does too).
11. LOW, launch-pricing-01-founder-yearly-1440.png / 02-monthly: "Founder prices until 6 January." has no year, and the mono "6  January" sits on an odd baseline mid-sentence. Fix: "until 6 January 2027", same font as the sentence.
12. LOW, launch-pricing-01-founder-yearly-1440.png: the band text (PLANS, h1, founder row) starts left of the card column, so the header and cards don't line up at 1440. Fix: put band content and cards on the same 1344 column edge.
13. LOW, launch-pricing-12-*: "Browser lane only" and "Off · Premium" read like a SaaS feature matrix, not a WoW tool. Fix: plain words, e.g. "Server sims: In your browser" / "Discord notifications: Premium only".
14. LOW, launch-pricing-09-picker-payer-crop-1440.png: "Mirelle · priest · pays for Premium". The class is lowercase where the client always writes "Priest". Fix: capitalise class names.
15. LOW, launch-pricing-10-waiting-crop-1440.png: "Download attendance (CSV)" draws gold-outlined at rest next to a plain twin (a leaked focus or hover state in the capture). Fix: recapture with no focus, or check that the button doesn't keep focus after Withdraw.
16. LOW, launch-pricing-11-lapsed-records-crop-1440.png: "Guild plan ended 1 Oct 2026" is earlier than launch day. Fix: fixture date inside the possible range.

## What already works (keep)

- Four cards, price first, one button each, no "most popular" push. Monthly and yearly swap cleanly, with the struck list price beside the founder price.
- "This list never shrinks" and the history promise ("records stay readable for good, CSV any time") are the two lines that make an officer trust this. Keep both above the fold on phone (360 does).
- The trial line says up front that a card is taken and that cancelling before day 8 charges nothing. The 7-day renewal email and the 14-day refund rule are stated. Nothing feels like an auto-renew trap.
- The seat picker's payer note ("They pay for Premium, so they get an offer... seat starts when it ends") and the offer's "You are not charged twice" line answer the exact question a member would ask.

## Re-review 1

Captures at 1d7aee9c7 (fixes 62facb8a6, 0c67dfa43). Same persona. I read the fresh 1440 crops and full pages, and the 390 seats.

### Verdict: SHIP WITH FIXES

Every original finding is fixed. The trust problems are gone: a Founder sees their Premium, seats add up, money lines agree, and the hand-over says what happens to seats. The fixes left two layout regressions on money lines, and those are the kind of thing a player screenshots: the guild Founder mark is glued to the price, and the chooser prices run together. Fix N1 and N2 before merge. N3 to N6 are small.

### Original findings

1. Fixed. launch-pricing-12-founder-crop-1440.png and 12-seat-holder-crop-1440.png: Server sims On, History charts On, Supporter mark "Shown on your characters", and no See plans.
2. Fixed. launch-pricing-09-seats-crop-1440.png / -390: "3 of 3 given", Kraggor / Lightbrand / Sagebloom in class colours, and the pending row titled with the current holder ("Moving to Mirelle"). The empty and returned rows are in 09-picker-payer-crop-1440.png ("Free since 15 Nov: Thistle left the guild.").
3. Fixed. launch-pricing-09-seats-crop-1440.png: the waiting row shows only Withdraw, with no Move seat and no lock date.
4. Fixed. launch-pricing-07-trial-crop-1440.png reads "$12"; 08-confirm-ready-crop-1440.png says "$5 ... From 1 Dec 2026 it is $20 a month", with Founder on the plan lines. (The mark itself is broken; see N1.)
5. Fixed. launch-pricing-11-read-only-officer-crop-1440.png: "The subscriber or the guild leader manages billing and seats." It's gone from the receiver's view (10-receiver-offer-crop-1440.png).
6. Fixed. launch-pricing-11-read-only-officer-crop-1440.png and 10-receiver-offer-crop-1440.png read "Billed by Thornwick", and the offer is from Thornwick. 12-seat-holder-crop-1440.png reads "Given by Thornwick".
7. Fixed. launch-pricing-08-confirm-ready-crop-1440.png: "renews 1 Dec 2026" and "From 1 Dec 2026" agree.
8. Fixed. launch-pricing-10-send-confirm-crop-1440.png: "Seats, raids and records stay with the guild; Mirelle manages seats from then." (See N4 for the receiver's wording.)
9. Fixed. launch-pricing-05-two-guilds-1440.png: real faction crests, and "Raid guild · $200 a year" under Which guild?
10. Fixed. launch-pricing-05-two-guilds-1440.png: the header shows Fixture#1234 signed in.
11. Fixed. launch-pricing-01-founder-yearly-1440.png: "Founder prices until 6 January 2027.", in the sentence's font.
12. Fixed. launch-pricing-01-founder-yearly-1440.png: the band and the cards share the column edge.
13. Fixed. launch-pricing-12-*: "On", "Off · set up", "Shown on your characters". No more "Browser lane".
14. Fixed. launch-pricing-09-picker-payer-crop-1440.png: "Bramblewood · Druid · pays for Premium".
15. Fixed. launch-pricing-10-waiting-crop-1440.png: the CSV buttons are plain. The ring is now on Withdraw, the expected focus after sending.
16. Fixed. launch-pricing-11-lapsed-records-crop-1440.png: "Guild plan ended 15 Nov 2026."

### New

N1. HIGH, launch-pricing-07-trial-crop-1440.png, 08-confirm-ready, 10-*, 11-read-only-officer: on every guild plan line the Founder mark renders as "$12Founder" / "$12 a monthFounder". There's no gap, no gold border and no capitals; it's a grey box glued to the price, nothing like the account's FOUNDER pill (12-founder-crop-1440.png). Fix: use the same `pill-site` pill as the account panel, 8 px after the price.
N2. MED, launch-pricing-11-lapsed-records-crop-1440.png (and the not-on-a-plan chooser): the option prices now read "$120$150 a year" and "$200$250 a year", with the struck price glued to the founder price. At a glance that's $120150. In the first round it was "$120 a year  $150 a year". Fix: put a space between them as on /premium ("$120  ~~$150~~ a year").
N3. MED, launch-pricing-09-seats-crop-1440.png: Lightbrand was "Given 29 Oct", so the seat can't move before 28 Nov, but a move was "offered 10 Nov". Either the data breaks the 30-day rule or the rule doesn't cover offers. Fix: the fixture gives Lightbrand a date at least 30 days before the offer, and the API refuses a move offer while the seat is locked.
N4. LOW, launch-pricing-10-receiver-offer-crop-1440.png: the receiver is told "Mirelle manages seats from then". That's third person to Mirelle herself. Fix: on the receiver's view, "you manage seats from then".
N5. LOW, launch-pricing-05-two-guilds-1440.png: under "Raid guild · $200 a year", Olympus XXVII's row says "Guild" on the right. I can't tell whether clicking it buys Raid guild or Guild. Fix: when a plan is already chosen, the right-hand hint only flags a mismatch ("12 members; Guild would cover them").
N6. LOW, launch-pricing-01-founder-yearly-1440.png: the "How the guild plans work" terms changed from small uppercase labels to bold mixed case. It's fine, but now it doesn't match the other label rows on the site. Fix: pick one and keep it.
