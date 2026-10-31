# Player verdict on the built F8 Premium screens (iPhone and Pixel 9)

VERDICT: SHIP (re-check below closed the one must-fix)

One must-fix, and it is missing evidence, not a visible defect: I have not seen See plans open the site.

## Must-fix
1. device-see-plans-browser-and.png is Chrome's first-run screen ("Make Chrome your own", "Add account to device", "Use without an account"). It does not show foreversixty.gg. I cannot confirm that See plans lands on the plans page. Pre-dismiss the Chrome first run on the emulator and re-capture the page the tap opens. It should be the /premium page (or the account page), with no Premium prompt added in the app.

## What a player sees, and it is right
- Free, US (you-free-us, device-you-free-us[-large]-and): "Plan / Free" with one small gold "See plans" link and an external-link arrow. Quiet, one link, no banner, no badge. Not pushy.
- Free, GB and unreadable country (you-free-gb, device-you-free-gb/none[-large]-and): "Plan / Free" only, no chevron, no link, no Premium word anywhere. This matches the owner rule. Free does not look broken; it looks finished.
- Premium (plan-us, plan-gb): "Premium · renews 27 Oct 2027 · $40 a year", the Founder tag, and "Founder price, kept while your plan stays active." This matches account-plan-copy.ts (planLine.renews, founder.title). Dates and prices read like the site.
- Seat (plan-seat): "Premium · a seat from OLYMPUS XXVII", "Given by Hunthrop#1894. It lasts while the guild's plan does; nothing is billed to you." This is verbatim planLine.seat and givenBy. It answers "am I being charged?". There is no billing button on the seat card, which is correct.
- Billing: US gets "Manage billing on foreversixty.gg" as a quiet outlined button on both platforms. GB gets the plain sentence "Billing is managed on foreversixty.gg." and no button. There is no purchase button, no price prompt and no in-app purchase anywhere. This matches the owner rules.
- The five "What your plan includes" lines are the website's Premium column words, in the website's order: Coach, server sims, Forever Score chart, Discord message, Supporter mark. This is the earlier must-fix, now closed.

## Shoulds
1. The app lists 5 of the site's 7 Premium lines. It leaves out the lead server-sim line, "Try up to N gear, talent or drop combinations in one sim and find your best DPS. Your browser stops at N.", and the iterations line. It also trims "Close the tab. Server sims keep running..." to "Server sims keep running, and the site tells you when they finish." The shortening is the app's own wording, not the site's. Either use the site's lines exactly or keep the cut to the five and note it in the spec. A Premium player will wonder where the 'more combinations' line went, since it is the headline benefit.
2. The footer "Raid sign-ups, best in slot, talents and your character are free for everyone, in the app and on the web." is not on the site. The site says "Everything a player needs is free, and the free list never shrinks." and lists "BiS lists", "Talent planner" and so on. "Raid sign-ups" is a Guild-plan item on the site (members never pay to sign up). Use the site's free sentence so the app and the site make the same promise.
3. Largest text, iPhone you-premium-ax5: the You rows break. The person avatar and the crown icon are pinned low or small and misaligned against the wrapped text, and the PREMIUM chip drops under the name. Everything stays readable and nothing is clipped, so it does not block. Centre the icons vertically against the row at large text. The Android large-text captures of the free rows are clean.
4. Seat holders see the Supporter mark and Discord message lines. The site lists the Supporter mark under Premium; confirm that seats really get it and the Discord message (the site says a seat holder "has Premium").
5. Founder tag, "Founder price..." line and the PREMIUM heading are three Premium cues on one card. It is harmless; keep it.
6. Android Plan row uses a star and iPhone uses a crown. The site's plan icon is not at stake; accept as platform style.

## Data I do not dispute
Prices ($40 a year) come from the site's plans table. A GB buyer sees dollars, the same as the site's "Prices in US dollars" foot. Fine.

Re-check: once the browser capture shows the real destination page, this is SHIP. The shoulds are not blocking.

---

# Re-check (device-see-plans-browser-and.png re-captured)

VERDICT: SHIP

Must-fix 1 is closed. The capture now shows Chrome on foreversixty.gg/premium ("What is free, and what the plans add", Yearly/Monthly, Founder line, the free list), reached by tapping See plans in the real app on a US storefront. The words on the page match what the app promises, and the app added no prompt of its own.

On the shoulds:
- The cut sim-cap and iteration lines are accepted. Those are numbers that change, and spec 5.8.1 records the cut. The site's free list ("BiS lists", "Talent planner", "Browser sims") does not clash with anything the app says.
- The footer is the owner-approved spec section 13 line. Accepted, given that /premium says members never pay to sign up for raids.
- Still open and not blocking: the large-text icon alignment on iPhone You rows, and confirming that seat holders get the Supporter mark and Discord message.
