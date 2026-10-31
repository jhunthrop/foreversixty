# Player verdict on the F8 Premium mock (iOS and Android boards, spec 5.8.1)

VERDICT: HOLD

## What works
- A player reads the Plan row at a glance: "Plan / Free" or "Premium · renews 27 Oct 2027", with the Founder tag beside it.
- The free GB row ("Plan · Free", no chevron, no button) is complete and quiet, not broken. Nothing is pushy: See plans is one small gold link, and no Premium word shows outside a link region.
- The seat card ("Given by ... nothing is billed to you") answers "am I being charged?". Dates (27 Oct 2027) and "$40 a year" read like the site's.
- The free-for-everyone footer is a good reassurance line. Largest-text frame stacks cleanly.

## Must-fix
1. The "What your plan includes" lines are not the website's words. The Premium column on /premium (web/src/lib/billing/plans-page-copy.ts) says: the Coach "what to fix first" line, "Close the tab. Server sims keep running, and the site tells you when they finish.", "See if you are improving: your Forever Score and every boss parse on a chart over time...", the Discord message, and the Supporter mark. The mock invents "Coach notes on your raid logs", "Score and parse history charts", and "Sim results and logs kept for a year". Owner rule: the app's Premium rules match the website exactly. Reuse the site's wording (shortened only with the site's own terms), or drop the invented lines.
2. "Sim results and logs kept for a year" appears nowhere on the site's Premium page. A player will compare the two and ask which is true. Either put it on the site first or remove it from the app.
3. "Score and parse history charts" shows ticked in every Premium frame. The site's 2026-10-24 changelog marked history charts as coming in November. Confirm the history entitlement flag is genuinely live; if not, the tick promises a feature the player cannot open. The kill-check must be tied to a real, working surface.

## Shoulds
- Frame 5 (GB): "Billing is managed on the web." is tiny muted text at the bottom of the card. It is fine, but a Premium player in GB cannot tell where on the web. Consider "Billing is managed on foreversixty.gg." as plain text, not a link (no link is the owner's rule, a domain name is not a call to action).
- Seat frame: the card never says which guild plan or what the seat includes beyond the four lines; the site's Supporter mark and the Discord message lines are absent. Make the list match the site's order and completeness.
- Free players never see what free includes; the "free for everyone" footer only lives on the Plan screen they cannot reach. Acceptable per the no-nagging rule, but the free GB row gives no hint it is intentional. Accept as is.
- Founder line "Founder price, kept while your plan stays active." is a price fact, not a prompt; fine, but in GB the price is in dollars. Confirm that is what a UK buyer was actually charged.
- The "Premium" gold tag at the head plus "Founder" on the row plus "PREMIUM" h2 on the card say Premium three times in one flow. Harmless, but the head tag alone is enough.
- Manage billing button is all caps in a boxed style; it reads as the loudest element on the Plan card. Keep it secondary.

---

# Re-check (boards re-rendered, f8-build-ios.png and f8-build-and.png)

VERDICT: SHIP

Must-fix 1 and 2 closed: the includes list is now the five lines of /premium's Premium column, in its order (coach, server sims keep running, Forever Score chart, Discord message, Supporter mark). Wording matches plans-page-copy.ts; the "kept for a year" line is gone. Must-fix 3 closed on the evidence given: the live /premium page lists history with no month, and the tick only shows when the history entitlement is true. The GB managed line now names foreversixty.gg as plain text, and Manage billing is a quiet outline button. A player reading the site and the app sees the same promises.

Shoulds (not blocking):
- The Discord-message line shows on the seat card and on a Premium account with no Discord linked. The Discord row says "Linked as huntrhop" in the mock, so it reads true there; for an unlinked player add nothing, but make sure the Discord screen explains the link is what turns it on.
- A seat holder is shown the Supporter mark line; confirm seats really get the mark, since the site lists it for Premium only.
- Prices are Stripe USD for every buyer; fine, but a GB player sees "$40 a year" with no currency note. The site does the same, so this matches.
- Premium still reads three times in one flow (head tag, row, card heading); harmless.
