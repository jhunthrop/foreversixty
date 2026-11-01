# Mobile Wave 3 F11: rankings and character pages (2026-11-01)

One lane: both apps and the API, from the same contract fixtures. Spec:
`design/specs/2026-10-31-mobile-wave3.md` §5.3 and §10 (W3-C5 to W3-C7). Boards:
`design/mocks/2026-10-31-mobile-wave3/renders/f11-rankings-{ios,and}`.

## API (W3-C5, W3-C6, W3-C7)

| Route | What it does |
|---|---|
| `GET /v1/rankings`, `/v1/rankings/percentile`, `/v1/phases` | The site's routes, unchanged, now app operations (X-App headers, 426, fixtures). |
| `GET /v1/rankings/encounters` | `api/internal/rankings/ranked_encounters.go`. Encounters with ranked kills (the boards' own rule), most kills first, with the newest log's zone and the usual raid size. |
| `GET /v1/me/watching`; `PUT`, `DELETE /v1/me/watching/{region}/{ruleset}/{name}` | `api/internal/rankings/watching.go`, migration **0065** `watched_characters`. Session or app. 50 an account (counted under a row lock), PUT idempotent and answers the row (200, not the spec's 204, so the app draws it without a second read), 404 for an unknown or anonymous character, 409 `limit`. Anonymous characters drop out of the list; level by the character page's rule. |
| Character card | Gains `component_caps` (rating's ceilings), so the apps draw the score bars as value over cap without the formula. |

## Fixtures: one story

`design/mocks/2026-10-31-mobile-wave3/f11_fixtures.py` writes every F11 fixture: the Elemental
Shaman Onyxia board on Thursday 17 December. Lightningbolltz is 14th of 212 at 398.0 DPS; his alt
Sparkwick is pinned 204th at 210.4. Figures fall off as a power of the rank. Voltaris (first) wears
nearly the whole Elemental raid list of the active build (one dot); Lightningbolltz is a week behind
in the Elements set (six dots), and his Home answer (`character.home.yours`) carries the same six
upgrades. Each card's Onyxia kill is its board row's fight. `cdn.items.shaman` holds their items.

## Apps

| | iOS (`apps/ios/App/Rankings`, `App/Common/Paperdoll*`) | Android (`app/.../rankings`, `app/.../paperdoll`) |
|---|---|---|
| Logic | Port of the web's `lib/rankings` rules: defaults by role (healer hps, tank damage taken), the default encounter, rank parse colour with the 99 cap, crest per row only on a mixed board, own row in place or pinned, paging deduped by rank; score bars; Paperdoll slot model | the same |
| Screens | Rankings (answer card, filter line, board, jump link, pin), filters sheet and encounter list, character page (score card with sparkline and five bars, bests, recent fights, Paperdoll gear, gear gap, builds seen, parse history for your own with `history`), no logs, not found | the same |
| Home | `Rankings` row under Logs; Your gear is the Paperdoll (F2.7) | the same |

- **Watching:** the bar star, optimistic, reverting with `{Name} was not saved. Try again.`; the
  row long-press menu (Watch, Open the report, Share).
- **Links:** `/app/rankings/{encounter}` and `/app/player/{region}/{ruleset}/{name}` in
  `apps/shared/links.json`; share links are the web's `/character/...` page.
- **Copy and switch:** `home.rankings*`, `rank.*`, `char.*`, `paperdoll.*` and `items.unknownItem`
  ("{Slot} · Item {id}", the owner's ruling) in `en.json`. The feature switch is `rankings`.
- **Codegen:** reserved words in endpoint parameters (`class`) are backticked in Swift and Kotlin;
  the Swift link parser is a matcher list (a chain of 12 matchers stalled the type checker).

## Reviews

Captures (default and largest text, both apps; the apps are dark only) are in
`design/reviews/2026-10-31-mobile-wave3-rankings/build/`.

- **Round 1: both HOLD.** ux-designer: iOS bar without star and share, Android filters under the
  tab bar, iOS chips and switch, the encounter sheet, the off-hand text. wow-player: identical gear on
  rank 1 and 14, Home naming the wrong class, the filters, the off-hand text, the unnamed answer
  card, the ordinal's colour. All fixed (`fix-round-1.md`), plus the two UX minors.
- **Round 2:** ux-designer SHIP; wow-player HOLD on Home mixing characters, fixed with Lightningbolltz's
  own Home answer.
- **Round 3:** see `player-verdict-build.md`.

## Gates

ROUND3_GATES

## Turning it on

`features.json` keeps `rankings` default false. To turn it on, add `rankings` to `APP_FEATURES` on
the API deploy (after the deploy that applies migration 0065). Not done from this lane.

## Gaps and follow-ups

1. **F13:** no `Open in planner` under Builds seen; the planner is not built.
2. **F12:** `Search again` on the not-found page goes back; Search is not merged. The Search
   `Rankings` quick link and Watching group belong to F12.
3. **F15:** rankings and player links open only for a signed-in player with a character (Home's
   sign-in gate, as F10). Spotlight entries for characters belong to F15.
4. **Device checks:** the real iOS long-press menu (the capture draws its items), the iOS star and
   share at the largest text size (`character-gear-large-ios` cannot draw the bar), the Android
   filters sheet scrolling with the tab bar visible, VoiceOver and TalkBack passes, haptics, budgets.
5. **Android bar:** it scrolls with the page as on F10, so its title stays empty under the head; a
   title on scroll needs a pinned bar on F10 too.
6. **BiS picker chips** (shared with signed-out Rankings) still show the stray border ticks fixed on
   the rankings filters.
7. **Deferred SHOULDs:** a unit a player knows for the gear gap, item level and talent split on
   rows, the Home fixture's older build, and the data items for the accuracy reviewer (the
   Leatherworking picks and Uther's Strength on an Elemental list, Totem of Thunder at 65, a five-man
   boss ranked beside Onyxia, the KR, TW and CN chips).
8. **Rating test hygiene:** two `rating` tests fail on a reused test database (rows left by earlier
   runs) and pass on a fresh one.
