# ux-designer verdict: header Premium tag for returning visitors

Verdict: SHIP

Bar balance (desktop, 1440):
- Without the tag, the search icon and the Sign in chip sit flush right.
- The right edge reads as a finished bar, with no gap and no orphan spacing.
- The nav items are unchanged.
- The tag is a secondary element, so removing it makes Sign in the single right-hand action. That is the better hierarchy for a first visit.

Phone (390):
- The menu head shows the bell crest and the hamburger.
- No Premium door is visible in the hidden state, so there is no empty slot.
- Only the top of the menu was captured, so I did not see whether a Premium row sits lower down.

Pointers to the tag:
- A grep of web/src for copy that refers to a "Premium button", "tag" or "top right" found none.
- The footer link and in-context offers are unchanged, so the paths to /premium remain.

Layout shift:
- The decision runs in the header's inline script, before the doors are parsed.
- data-session is stamped in Base's head script, which runs earlier, so signed-in free accounts correctly keep the tag.
- The tag uses display:none, not a post-paint removal, so no shift.

Rule:
- Three distinct local days is sensible.
- Keeping the last three days means the tag keeps showing once reached.
- Blocked or malformed storage counts as a first visit, which fails toward free-tool.

Minor (polish, not blocking):
- The API's /b/{id} page has no Base head script, so data-session is unset there. A signed-in free account sees no tag on that page. This is harmless.
- Confirm the phone menu's lower Premium row, if one exists, is hidden by the same rule. A bottom-of-menu capture, hidden and shown, would settle it.
