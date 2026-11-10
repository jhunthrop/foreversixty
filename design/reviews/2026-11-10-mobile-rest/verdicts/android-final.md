# Android secondary screens: final gate (2026-11-10)

| Group | ux-designer | wow-player | ai-slop | info-density | brand-consistency | guild-officer |
|---|---|---|---|---|---|---|
| guides | SHIP (r3) | SHIP (r3) | PASS (r3) | DENSE (r3) | CONSISTENT (r5) | |
| character | SHIP (r3) | SHIP (r3) | PASS (r3) | DENSE (r3) | CONSISTENT (r3) | |
| bis-search | SHIP (r3) | SHIP (r3) | PASS (r3) | DENSE (r3) | CONSISTENT (r3) | |
| account | SHIP (r3) | SHIP (r3) | PASS (r3) | DENSE (r3) | CONSISTENT (r3) | |
| guild-raids | SHIP (r3) | SHIP (r4) | PASS (r3) | DENSE (r3) | CONSISTENT (r5) | WORTH PAYING FOR (r4) |
| sims-reports-talents | SHIP (r3) | SHIP (r3) | PASS (r3) | DENSE (r3) | CONSISTENT (r3) | |

Round 4 and 5 fixes:
- raid night's meta line uses `mutedGround`;
- the Guides root keeps its gutter without a lead (commit 1f4fd2e53).

Brand round 5 looked at `guides/f9-tiers-nochar-and.png` and `raids/raid-night-premium-and.png` as re-recorded at 14:29.

## Follow-ups named by the reviewers (none block a verdict)

**Restyle polish:**
- skeletons that draw the real head;
- crests on the raid page's sign-up rows;
- the roster sort menu anchored under its chip;
- the MotW buff box;
- the loading roster's figures;
- the Gear gap unit;
- the Forever Score bars by role;
- the Elemental Focus text wrapping in full;
- naming the Top Gear slots under the grid;
- the BiS loading rows;
- "Runner-up" once per list;
- the Mana Spring Totem rank value.

**Data and product:**
- stat weights on the guide;
- the signed-out BiS paperdoll;
- the character's head answer;
- a spec rank for Forever Score;
- sim deltas and weights in the app;
- drop rates on sources;
- device platform;
- the store's localized price;
- the tooltip's rating wording;
- Fire Nova's icon;
- Recently updated on Guides;
- comp targets for members;
- raid-ready flags on the roster;
- a loot-fairness view;
- raid assignments and the boss list.

**Mock:** `ws_app_flows._signup_buttons` should draw the round-5 segmented answers. That clears the declared raid-night drift.
