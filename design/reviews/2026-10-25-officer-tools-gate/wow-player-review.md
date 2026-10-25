# Officer tools gate: wow-player review (2026-10-25)

Reviewed from the PNGs in this folder (01-13, full pages and crops at 360/390/1440/2000) and spec §7.
Fixture names, specs and dates are not judged.

## 1. The question
"Why can't I press Award or Nudge, and how do I get them back?" The answer is the first thing in
the tab: one line under the `LOOT` / `READINESS` heading, above the boss picker and the roster grid,
at every width (02, 03, 07). Yes, it is first on screen.

## 2. What a raid officer would find wrong (most damaging first)
1. **The grant ends a raid night early for every guild in the Americas (data, see §4).** Settings and
   the account row read "Officer tools · until 11 Dec 2026". Every officer will read that as "through
   the 11th". It actually ends at 00:00 UTC on 11 Dec, which is Thursday 10 Dec at 7 pm Eastern and
   4 pm Pacific. Raids open Wed 9 Dec, so a US guild's Thursday raid loses Award at invite time.
   The 06 refusal ("Not awarded. This guild's officer tools have ended.") would land mid-Onyxia,
   on the night the grant was meant to cover. That is the screenshot that ends up in guild chat.
2. **An officer who is not the GM gets an order they cannot carry out.** The gate says "Subscribe the
   guild." The link lands on settings, and settings (11) says "Only the guild master can release the
   claim or change billing." That line never mentions subscribing. It is not a dead end, because it
   names who to ask. But it reads like a second refusal, not like "go ask your GM".
3. **The pair of words is wrong on both tabs.** Loot holds Award and Undo but no Nudge. Readiness
   holds Nudge but no Award. Settings says "Award, Undo and Nudge". A raid lead reads "Award and
   Nudge" on the loot tab and wonders whether Undo still works. Name all three, matching settings.
4. **The link breaks at 360** (02-360, 07-360): "Subscribe the" / "guild." The call to action is
   split over two lines. Tie it together with a non-breaking space.
5. **The refusal line in 06 is the quietest thing on the card.** It is muted grey, 13 px, under
   the rows. The reference the spec cites, the client's "Not enough rage", is red and sits where
   you are looking. In a raid you click Award and look at the master-loot window, not at grey text.
   Use the tab's existing error colour.
6. **The disabled Undo reads as a label** (04-1440). Plain grey "Awarded · Undo" beside the
   "Already awarded" lines looks like a typo, not a greyed button. It is minor because Undo is rare.
7. **The phone focus ring is a 44 px box** that runs over the top of the boss picker (09-390). You
   only see it with a keyboard on a phone, so it is cosmetic.

What works: the dimmed AWARD and NUDGE buttons keep their slot, their label and their border, like
a greyed-out action bar button. Together with "Read-only" at the heading's right edge, they read
as locked, not broken. Nothing is hidden. It is one line per tab with no modal, no banner, no lock
icons and no upsell per row, so it does not nag you during a raid. The refusal sits inside the
card you clicked, and both tabs flip without a reload. "Officer tools have ended" is the phrase a
raid lead would actually say.

## 3. Missing
- **Settings never says why the grant exists or why it stops on that date.** The brief says the
  grant runs "through the first raid night". Saying so would turn a countdown into a gift.
- **The price line still promises "progression tools"** (10, 13 context): "$15 a month, cancel any
  time. Officers get the loot, readiness and progression tools." No progression tool ships. An
  officer about to pay $15 will ask where it is. This copy belongs to another lane (spec §7), but
  it is on screen next to this gate.
- **The account row layout is uneven** (13-1440). On the first guild's row, LEAVE drops to its own
  line; on the second, it sits inline. This is cosmetic.

## 4. Data I do not believe
- **The grant end time.** The code uses `--until 2026-12-11`, read as `time.Date(2026,12,11,0,0,0,0,UTC)`
  in `api/cmd/api/grant_test.go` and `api/README.md`, and the label renders "until 11 Dec 2026".
  For the Americas that is the evening of 10 Dec, the second night raids are open. Either move the
  end to after the first full raid week (for example 2026-12-16 00:00 UTC, so it covers Wed to Tue
  everywhere), or end it at 11 Dec 23:59 Pacific. The label must name a day the player's own clock
  agrees with.

## 5. Verdict: SHIP WITH FIXES
1. **Grant end (data, 10/11/13):** move `until` so that "until {date}" holds for US raid nights.
   2026-12-16 00:00 UTC covers the first raid week. If the date stays, render the label in the
   viewer's local date, so Americas players see "until 10 Dec 2026".
2. **Gate sentence (02, 03, 07, 09):** replace
   `Award and Nudge are part of the guild plan. Subscribe the guild.`
   with `Award, Undo and Nudge are part of the guild plan. Subscribe the guild.`
   Put a non-breaking space in the link text: `Subscribe the guild`.
3. **The settings line for officers without the claim (11, and the expired not-on-plan state):**
   replace `Only the guild master can release the claim or change billing.`
   with `Only the guild master can subscribe the guild, release the claim or change billing.`
4. **The settings explanation (10, 11):** replace
   `Granted by Forever Sixty for Award, Undo and Nudge only. It ends on 11 Dec 2026; the guild plan keeps them after that.`
   with `Granted by Forever Sixty for the first raid nights: Award, Undo and Nudge only, until 11 Dec 2026. After that they need the guild plan.`
5. **Refusal line (06):** render it in the tab's error colour, not muted grey. Keep the words.
6. **Should fix, not blocking:** the 09-390 focus ring overlapping the boss picker, the plain "Undo" in 04,
   the gold `Guild settings` link (13; the spec calls for a muted link), and "progression tools" in
   the price line (another lane).
