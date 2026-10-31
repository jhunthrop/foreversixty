# Mobile Wave 1, F8 Premium (2026-10-31)

Branch `wt/mobile-f8-premium`, one lane building both platforms. Spec: `design/specs/2026-10-30-mobile-apps.md` §5.8 and the new §5.8.1 build notes. Boards: `design/mocks/2026-10-30-mobile-apps/renders/f8-build-{ios,and}.png`. Captures and verdicts: `design/reviews/2026-10-31-mobile-premium/`.

## What shipped

- **The plan, in the site's words (decision G).** `PremiumSource` on both platforms ports the site's `premiumSource` (`web/src/lib/account/plan-line.ts`): own, own then seat, grant, seat, fallback and none, with the same live statuses. `PlanSummary` words each source the way the account page does (`accountPlanCopy`, `premiumSourceCopy`). The You row carries the line without its price. The Plan card (F8.4) carries the full line with the price (`formatCents`: `$40 a year`), the second line for grants and seats, and Founder only for an own subscription.
- **What the plan includes.** These are the entitlement flags that are true, worded as /premium's Premium column and in its order: coach, server sims, history, Discord message, Supporter mark. `web/src/lib/mobile/premium-copy.test.ts` fails if an app line is not a sentence of a /premium line, or if the grant and seat lines drift from the account page.
- **See plans only where the store allows it (decision P).** `PlansLink` resolves to "may link" only when all three hold:
  - `premium_link` is on;
  - `GET /v1/app/config` has answered, from this launch or the cache;
  - the storefront country is in `web_checkout_regions`.
  The storefront is StoreKit `Storefront.current` on iOS (alpha-3, converted with Foundation's `Locale` region) and Play Billing `getBillingConfigAsync` on Android (alpha-2, with a 5 s timeout). Anything missing fails closed. Free in a link region shows `See plans`, and the whole row opens foreversixty.gg/premium in the browser. Free elsewhere shows `Plan · Free`: not a button, and no other Premium word. Premium elsewhere shows `Billing is managed on foreversixty.gg.` in place of the quiet Manage billing button.
- **The account head** carries the gold Premium tag for any source. At the largest sizes the tag drops under the name, and the Founder tag drops under the card heading.
- **Fixture.** `me.seat.json` now resolves the Premium entitlements a real seat gets. They were all false, which emptied the seat card.
- **Capture hooks (debug builds only).** iOS reads `-FSStorefront USA|GBR|none`. Android reads the `fs_storefront` launch extra.

## Process

- **Design.** The ux-designer wrote §5.8.1 and the boards. The wow-player held on the includes list, which was app-only wording plus "kept for a year", a line /premium never says. After the fix both reviewers gave SHIP (`ux-verdict-mock.md`, `player-verdict-mock.md`).
- **Tests first.** The logic list is the same on both platforms (iOS `PremiumTests`, Android `PremiumTest`, 29 cases). It covers each source and its words, prices, includes, row actions, billing lines, and screen-reader labels (iOS). Fail-closed is tested with no config, an unreachable config endpoint (the state before A3), no storefront, the feature off, and the US and GB storefronts against the contract fixture.
- **Screenshots.** Both platforms cover You free US, free GB, a seat, You free US and Premium at the largest text size, and the Plan screen in the US, in GB, for a seat and at the largest size.

## Gates (on the tree merged with origin/main; each run alone, exit codes checked)

| Gate | Result |
|---|---|
| `make -C apps/ios test` (shared simulator, iOS 26.5) | exit 0, TEST SUCCEEDED: app-hosted 267 in 32 suites, packages 115 |
| `./gradlew test :app:verifyRoborazziDebug lint assembleDebug` (JDK 21) | exit 0, 380 JVM tests, 0 failures |
| `npm run check`, `lint`, `format:check` | exit 0 each |
| `npx vitest run src/lib/mobile src/styles/mobile-tokens.test.ts` | exit 0, 35 passed (7 new) |
| `node web/scripts/gen-mobile-{api,copy,features,links,tokens}.mjs` | no drift |
| `go test ./internal/contract/ ./internal/server/` on `forever_f8_premium` (created, then dropped) | ok, ok |

## Verdicts

- Mock: ux-designer SHIP; wow-player HOLD (includes wording), then SHIP.
- Build captures: ux-designer SHIP; wow-player HOLD (the See plans capture showed Chrome's first run), then SHIP on the re-capture of foreversixty.gg/premium.
- Open shoulds are listed in `ux-verdict-build.md` and `player-verdict-build.md`. They cover the Founder tag position at the largest size across platforms, a line starting with `·` at Android 2.0, tick weight parity, and crown vs star tile icons.

## After merging F5

F5 (sims) left a fail-closed `mayLinkToPlans` hook for its lapsed-plan line. On both platforms it now reads `PlansLink`, so the second `See plans` (spec §5.8: at most two) follows the same rule. F5's copy of the plans URL is replaced by the one in `AccountLinks`.

## Concerns

1. **Play Billing dependency.** Adding it puts `com.android.vending.BILLING` in the merged manifest, even though the app only reads the country. The Play Console may ask about in-app products, and the Data safety form should say none are sold.
2. **No device-level Premium capture.** Fixture mode serves the free `me.json`. The Premium, seat and Plan states are rendered by the screenshot tests from the real views. The emulator captures cover free US, GB and an unreadable storefront, plus the browser opening.
3. **iOS emulator-style captures.** simctl cannot tap Sign in, so the iOS review images are the snapshot goldens. They are rendered on the shared simulator's iOS 26.5 runtime, the same one CI uses.
4. **Owner questions still open.**
   - Spec §17 Q2: should `plans_link.android` ship empty until the Play programme enrolment?
   - Spec §17 Q3: Apple 3.1.3(b) outside link regions.
   - Whether seat holders get the Supporter mark. The app shows whatever the API resolves.
5. **The You rows owned by F7 and F1** (Notifications, Characters, Discord) are drawn on the boards but are not part of this lane.
