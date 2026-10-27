# Launch pricing: UX review on screenshots, 2026-10-27

Reviewer: ux-designer. Spec: design/specs/2026-10-27-launch-pricing.md (with the R-P8 ruling and the Task 16 amendment). Evidence: every PNG in this folder, cls-*.json, and the mock boards in design/mocks/renders/launch-pricing-*.png.

## Verdict: SHIP WITH FIXES

Fix the three HIGH items before merge, then recapture rows 8, 9 and 14. Fix the MED items in this lane. The LOW items can go to a follow-up issue.

## Findings

1. HIGH. `launch-pricing-14-built-1024.png`. At 1024 the Guild and Raid guild trial line wraps to 3 lines in its fixed 40 px slot, and the first feature divider strikes through "nothing is charged.". This breaks spec §7 (the trial line never collapses) and tenet 4. Fix: size the trial slot to the tallest card's trial line (shared grid rows or subgrid), not a fixed 40 px.
2. HIGH. `launch-pricing-08-confirm-loading-*.png`, `07-trial-*`, `08-size-line-*`. The disabled commit `Move to Raid guild` is pixel-identical to the enabled one (229,185,85). The disabled `Transfer the plan` looks enabled too. `PRIMARY_BUTTON_FIXED` and `SECONDARY_BUTTON_FIXED` have no disabled style, against spec §4 Disabled. Fix: add `disabled:opacity-50 disabled:cursor-not-allowed` (and the aria-disabled twin) to both _FIXED recipes.
3. HIGH. `launch-pricing-09-seats-*.png`. Row 2 reads "Kraggor / Mirelle WAITING" with a locked `Move seat`, then `Withdraw`, then two dates. That is not a §3.2.5 state, and an officer cannot tell who holds the seat. Fix: line 1 is the holder (Kraggor, Given date). The middle reads `Moving to Mirelle · offered 10 Nov. They pay for Premium, so they accept it first.` with `Waiting` beside Mirelle. Drop the second date and keep only `Withdraw`.
4. MED. `launch-pricing-09-seats-*.png`, `09-picker-payer-*.png`. The fixture is untrue to the product (tenet 8). The Raid guild reads `3 of 5 given`, but the plan has 3 seats. Every holder is Kraggor. The accepted seat "starts 1 Nov", which is in the past on the 20 Nov clock. Given 1 Nov reads `Can move from 2 Dec` on one row and `1 Dec` on another. Fix: 3 seats, distinct G1 roster names, a future start and coherent 30-day dates, then recapture.
5. MED. `launch-pricing-09-seats-crop-360.png`, `08-size-line-crop-360.png`. On phone every seat button stacks on its own line, about 170 px per row. Spec §7 says who and the button share line 1 and the middle text takes line 2, at most 112 px. Fix: below 768, a 2-row grid (`1fr auto`, the middle spanning row 2).
6. MED. `launch-pricing-08-confirm-ready-crop-360.png`, `09-picker-payer-crop-390.png`, `10-send-confirm-crop-390.png`. Confirm-block buttons sit inline below 768. Spec §7 says full width, commit first. Fix: below md, `flex-col` with `w-full` on both buttons.
7. MED. `launch-pricing-08-confirm-loading-crop-*.png`. While loading, the sentence stops at "Today you pay [skeleton]". Spec §3.2.4 says the Skeleton replaces only the amount. Fix: render the rest of the sentence (period, renewal date, price) around the 64 x 14 Skeleton.
8. MED. `launch-pricing-08-confirm-ready-crop-*.png`. The README says Keep Guild is focused, but no 2 px gold ring shows (spec §4). Fix: show the ring when focus moves in code, and recapture from a keyboard Enter.
9. MED. `launch-pricing-10-waiting-crop-*.png`. After `Send the offer`, `Download attendance (CSV)` renders in its gold hover/focus look, so focus fell through to Records. Fix: move focus to `Withdraw` when the waiting state renders.
10. MED. `launch-pricing-01-founder-yearly-1440.png`, `-2000.png`, `14-built-1024.png`. The band inner column starts about 30 px left of the card column at every width. The interval control overhangs the Raid guild card's right edge. The mock (`14-side-by-side-*`) shares one edge. Fix: one 1344 px container and gutter for the band inner and the body.
11. MED. `launch-pricing-01-founder-yearly-1440.png`. The Premium button bottom sits at about 545 px, against spec §1 (520 px at 1280 x 800). The band carries about 50 px of dead space between the h1 row and the lede. Fix: lede 8 px under the h1, with the interval control kept on the h1 row.
12. MED. `launch-pricing-04-hover-card-button-crop-1440.png`. The hover capture also carries the focus ring, and its ground stays `--raised`. Spec §4 asks for an `rgba(229,185,85,.08)` ground on hover. Fix: add the hover ground, then recapture hover with focus blurred.
13. MED. `launch-pricing-06-trial-available-*.png`, `06-trial-used-*`, `11-lapsed-records-*`. The chooser sets the whole "$120 a year" and the struck "$150 a year" in mono, so it reads like code and the strike runs through the words. Spec §3.2.2 makes only the price mono 20 px. Fix: mono figure, then ` a year` / ` a month` in Barlow 13 px muted, with only the list figure struck.
14. MED. `launch-pricing-05-two-guilds-*.png`. Both guild rows show an empty gold ring with no crest. That is a missing icon (tenet 13); §3.1.7 asks for the crest render. Fix: stub the crests in the harness, have each crest arrive with its row (reserved box, no flash-in), then recapture.
15. MED. `launch-pricing-12-seat-ended-crop-*.png`. "Not on Premium. See plans" shows above the ended sentence, so there are two competing links. The come-back link is white, not the gold inline link. "ended 24 Nov" is after the 20 Nov clock. Fix: the ended sentence replaces the plan line (§3.3), the link turns gold, and the fixture date moves on or before the clock.
16. MED. `launch-pricing-12-offer-*`, `12-accepted-*`, `12-founder-*`, `12-seat-holder-*`. A paying or seat-holding Premium account shows "Browser lane only / Sparkline only / Off · Premium / Not shown". The "Discord messages… See plans" upsell also shows. Fix: give the fixtures Premium entitlements and confirm the rows derive from the entitlement. Hide the upsell line for anyone on Premium (premium spec §2).
17. MED. `launch-pricing-10-receiver-offer-*`, `11-read-only-officer-*`, `12-seat-holder-*`. These read "Billed by Fixture#1234" and "Given by Fixture#1234", a BattleTag. Spec §3.2.1/§3.3 shows a character name (`Billed by Kraggor`). The receiver's offer comes from Thornwick, but Billed by names someone else. Fix: name the payer's main character, with PublicName as the fallback, and make the fixture coherent.
18. MED. `launch-pricing-09-picker-payer-*.png`. The option reads "Mirelle · priest · pays for Premium". The class is lowercase, and Mirelle already holds the waiting offer on row 2. Fix: write `{Class}` capitalised (`Priest`) and leave members with an open seat offer out of the picker.
19. MED. Copy `previewExpired`: "Prices changed. Review the new amount." is wrong and alarming, because only the prorated amount aged. Fix: `Today's amount was updated. Check it, then move.`
20. LOW. Copy `upgrade.unconfirmed`: "Stripe is still confirming the change. The plan below is up to date." contradicts itself. Fix: `Moved. Stripe is still confirming the charge.`
21. LOW. Copy `upgrade.unconfirmedStale`: "…Reload to see where it stands." is vague. Fix: `Stripe is still confirming the move. Reload in a minute to see the plan.`
22. LOW. Copy `whoPays.inProgress`: "An offer or a plan change…" makes the officer guess which one. Fix: when a transfer is open, `An offer to {Name} is already waiting. Withdraw it first.`; otherwise `A plan change is already under way. Reload to see it.`
23. LOW. Copy `picker.seatCount` ("This plan has no seat to give. Reload to see the current seats.") and `picker.offerPending` ("That seat already has an offer waiting. Withdraw it first."): accepted as written.
24. LOW. `launch-pricing-01-founder-yearly-*.png`. "until 6  January" puts a mono date in the 13 px prose with a visibly wide gap. Fix: trim the whitespace outside the mono span and match 13 px.
25. LOW. `launch-pricing-14-side-by-side-*.png`. How-it-works terms are 11 px muted labels, against the mock's 12 px `--strong` 600, which makes them hard to scan. Fix: match the mock.
26. LOW. `launch-pricing-06-trial-available-1440.png`. Each chooser card stretches to about 618 px with an empty right half. Fix: cap the chooser at 720 px.
27. LOW. `launch-pricing-10-receiver-offer-crop-1440.png`. "Only the guild master can … change billing" sits under an offer to take over billing. Fix: hide `guildMasterOnly` while an offer to the viewer stands.
28. LOW. `launch-pricing-09-picker-payer-crop-1440.png`, `08-confirm-*`. The opener (`Give seat`, `Move to Raid guild`) stays a live gold button beside a same-label commit. Fix: set `aria-expanded="true"` on the opener and draw it pressed (`--card-top`).
29. LOW. `launch-pricing-12-accepted-crop-*.png`. The accepted status keeps the offer's gold-rule block. Spec §3.3 asks for a 13 px muted status line. Fix: drop the rule and the block padding.
30. LOW. `launch-pricing-01-founder-yearly-360.png`. The Free summary row is about 80 px because its line wraps (spec 64 px). Fix: set min-height 64 and amend the spec height.
31. LOW (out of lane). `launch-pricing-06-*`. The settings page has a bare h1 and no guild band, but the mock and the owner's rule ask for a full-bleed band. Fix: follow-up issue.

## Checklist (spec §9)

- [x] Row 1 at 360/390/1440/2000. The founder row, the struck prices and Yearly checked all render. At 360 the Premium button sits above y 600.
- [x] Row 2. Monthly prices and buttons are right, and the Free summary row stays first.
- [x] Row 3. After the window there is no founder row and no strike. Launch day matches founder-open (R-P8 ruling).
- [~] Row 4. The focus rings are right; the hover capture is contaminated and has no ground (12).
- [~] Row 5. The rows and the none line are right; the crests are missing (14).
- [x] Row 6. The fitting plan is preselected. With the trial used, the trial line is gone and the button carries the price.
- [~] Row 7. The trial plan line and its cancel-before date are right; the disabled Transfer looks enabled (2).
- [~] Row 8. The size line and the single Move to Raid guild are right. Open: loading (2, 7), ready (8), phone (6).
- [ ] Row 9. Findings 3, 4, 5 and 18.
- [~] Row 10. The copy is right. Open: focus after send (9) and the fixture names (17).
- [x] Row 11. Lapsed shows the read-only line, Subscribe again and both CSV buttons. The read-only officer view has no buttons.
- [~] Row 12. The states render. Open: findings 15, 16 and 17.
- [ ] Row 13. The five emails are not captured; they still need real mail clients before launch.
- [~] Row 14. The layout is centred at 2000 and the cards keep their width. Open: misalignment (10) and the 1024 clipping (1).
- [x] CLS for a warm #billing load is 0.0001 at 1440 and 0 at 390, under the 0.02 cap.
- [x] Provenance stays quiet, each page keeps one job, and the seat crests are ClassCrestRing.
