# Launch pricing acceptance captures, 2026-10-27

Spec: design/specs/2026-10-27-launch-pricing.md section 9 (screenshot table). Fixture build at the lane HEAD, API stubbed, clock faked (8 Dec 2026 for /premium, 20 Nov 2026 for billing and account). Harness: web/tests/captures/launch-pricing/. Rows 1-12 and 14 are here; CLS for the warm #billing load is in cls-390.json and cls-1440.json (cap 0.02).

File names: `<prefix><state>-<width>.png`; widths are the viewport widths. `-crop` is a padded crop of the subject region from the same capture.

| File | Widths | State |
|---|---|---|
| `launch-pricing-01-founder-yearly-<w>.png` | 360, 390, 1440, 2000 | /premium on 8 Dec 2026: founder row open, Yearly checked, struck list prices |
| `launch-pricing-02-monthly-<w>.png` | 390, 1440 | /premium with Monthly selected |
| `launch-pricing-03-after-window-<w>.png` | 390, 1440 | /premium at the founder window close: no founder row, no strike, list prices |
| `launch-pricing-03-launch-day-<w>.png` | 390, 1440 | /premium on 27 Oct 2026 (launch day, founder window open; R-P8 ruling: there is no "before" state) |
| `launch-pricing-04-focus-card-button-crop-<w>.png` | 1440 | Focus-visible on the Premium card button (card crop) |
| `launch-pricing-04-focus-interval-<w>.png` | 1440 | Focus-visible on the Yearly/Monthly interval control (full page) |
| `launch-pricing-04-hover-card-button-crop-<w>.png` | 1440 | Hover on the Premium card button (card crop) |
| `launch-pricing-05-no-billable-guild-<w>.png` | 390, 1440 | Checkout guild chooser, no billable guild: none line and Find your guild |
| `launch-pricing-05-two-guilds-<w>.png` | 390, 1440 | Checkout guild chooser, two billable guilds (a third without billing is hidden), Raid guild fits line |
| `launch-pricing-06-trial-available-<w>.png` | 360, 390, 1440 | #billing not on a plan, 28 members, trial available (Raid guild preselected) |
| `launch-pricing-06-trial-available-crop-<w>.png` | 360, 390, 1440 | #billing not on a plan, 28 members, trial available (Raid guild preselected) (crop of the subject region) |
| `launch-pricing-06-trial-used-<w>.png` | 360, 390, 1440 | #billing not on a plan, trial used: Subscribe with price, no trial line |
| `launch-pricing-06-trial-used-crop-<w>.png` | 360, 390, 1440 | #billing not on a plan, trial used: Subscribe with price, no trial line (crop of the subject region) |
| `launch-pricing-07-trial-<w>.png` | 390, 1440 | #billing on a trial: first charge and cancel-before date |
| `launch-pricing-07-trial-crop-<w>.png` | 390, 1440 | #billing on a trial: first charge and cancel-before date (crop of the subject region) |
| `launch-pricing-08-confirm-loading-<w>.png` | 360, 390, 1440 | Upgrade confirm while the preview loads |
| `launch-pricing-08-confirm-loading-crop-<w>.png` | 360, 390, 1440 | Upgrade confirm while the preview loads (crop of the subject region) |
| `launch-pricing-08-confirm-ready-<w>.png` | 360, 390, 1440 | Upgrade confirm ready: preview amount, Keep Guild focused |
| `launch-pricing-08-confirm-ready-crop-<w>.png` | 360, 390, 1440 | Upgrade confirm ready: preview amount, Keep Guild focused (crop of the subject region) |
| `launch-pricing-08-size-line-<w>.png` | 360, 390, 1440 | #billing Guild tier with 28 members: size line and Move to Raid guild |
| `launch-pricing-08-size-line-crop-<w>.png` | 360, 390, 1440 | #billing Guild tier with 28 members: size line and Move to Raid guild (crop of the subject region) |
| `launch-pricing-09-picker-payer-<w>.png` | 360, 390, 1440, 2000 | The open seats with the picker open on the empty seat and a paying member chosen (payer note) |
| `launch-pricing-09-picker-payer-crop-<w>.png` | 360, 390, 1440, 2000 | Seats with the picker open on an empty seat and a paying member chosen (payer note) (crop of the subject region) |
| `launch-pricing-09-seats-<w>.png` | 360, 390, 1440, 2000 | #billing Raid guild, 3 of 3 seats: held and locked (Kraggor), moving to a payer (Lightbrand to Mirelle, Waiting), accepted and starting later (Sagebloom) |
| `launch-pricing-09-seats-crop-<w>.png` | 360, 390, 1440, 2000 | The same, cropped to the seats |
| `launch-pricing-09-seats-open-<w>.png` | 360, 390, 1440, 2000 | #billing Raid guild, 1 of 3 seats: held (Kraggor), empty, returned (Thistle left the guild) |
| `launch-pricing-09-seats-open-crop-<w>.png` | 360, 390, 1440, 2000 | The same, cropped to the seats |
| `launch-pricing-10-receiver-offer-<w>.png` | 390, 1440 | Who pays: the receiver's view of the offer, Take over the plan |
| `launch-pricing-10-receiver-offer-crop-<w>.png` | 390, 1440 | Who pays: the receiver's view of the offer, Take over the plan (crop of the subject region) |
| `launch-pricing-10-send-confirm-<w>.png` | 390, 1440 | Who pays: send-transfer confirm with its consequence line |
| `launch-pricing-10-send-confirm-crop-<w>.png` | 390, 1440 | Who pays: send-transfer confirm with its consequence line (crop of the subject region) |
| `launch-pricing-10-waiting-<w>.png` | 390, 1440 | Who pays: offer sent, waiting for them, with Withdraw |
| `launch-pricing-10-waiting-crop-<w>.png` | 390, 1440 | Who pays: offer sent, waiting for them, with Withdraw (crop of the subject region) |
| `launch-pricing-11-lapsed-records-<w>.png` | 390, 1440 | #billing lapsed: read-only line, Subscribe again, Records with both CSV buttons |
| `launch-pricing-11-lapsed-records-crop-<w>.png` | 390, 1440 | #billing lapsed: read-only line, Subscribe again, Records with both CSV buttons (crop of the subject region) |
| `launch-pricing-11-read-only-officer-<w>.png` | 390, 1440 | #billing read-only officer: plan line, Billed by, seats with no buttons, Records |
| `launch-pricing-11-read-only-officer-crop-<w>.png` | 390, 1440 | #billing read-only officer: plan line, Billed by, seats with no buttons, Records (crop of the subject region) |
| `launch-pricing-12-accepted-<w>.png` | 360, 390, 1440 | Account Plan panel after accepting: status line and hand-over line |
| `launch-pricing-12-accepted-crop-<w>.png` | 360, 390, 1440 | Account Plan panel after accepting: status line and hand-over line (crop of the subject region) |
| `launch-pricing-12-founder-<w>.png` | 360, 390, 1440 | Account Plan panel: Founder mark |
| `launch-pricing-12-founder-crop-<w>.png` | 360, 390, 1440 | Account Plan panel: Founder mark (crop of the subject region) |
| `launch-pricing-12-offer-<w>.png` | 360, 390, 1440 | Account Plan panel: seat offer to a payer (own end date and price) |
| `launch-pricing-12-offer-crop-<w>.png` | 360, 390, 1440 | Account Plan panel: seat offer to a payer (own end date and price) (crop of the subject region) |
| `launch-pricing-12-seat-ended-<w>.png` | 360, 390, 1440 | Account Plan panel: ended seat with its come-back link |
| `launch-pricing-12-seat-ended-crop-<w>.png` | 360, 390, 1440 | Account Plan panel: ended seat with its come-back link (crop of the subject region) |
| `launch-pricing-12-seat-holder-<w>.png` | 360, 390, 1440 | Account Plan panel: seat holder, no Manage billing |
| `launch-pricing-12-seat-holder-crop-<w>.png` | 360, 390, 1440 | Account Plan panel: seat holder, no Manage billing (crop of the subject region) |
| `launch-pricing-14-built-<w>.png` | 1024, 1440, 1920, 2000 | The built /premium at this width (8 Dec 2026) |
| `launch-pricing-14-mock-<w>.png` | 1024, 1440, 1920, 2000 | The mock board launch-pricing (rendered by the harness at this width) |
| `launch-pricing-14-side-by-side-<w>.png` | 1024, 1440, 1920, 2000 | Same-scale side by side, mock left, built right |

Not captured: row 13 (the five emails in Gmail web and iOS Mail) needs real mail clients. Row 14 is also taken at 1920 as the owner asked; the mock board has no 1024 or 1920 render, so the 1440-wide board is drawn at that viewport.

