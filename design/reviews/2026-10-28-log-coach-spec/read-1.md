# Task 33 - WoW-player read of the log coach spec and boards (2026-10-27)

Verdict: FIX. The shape is right; 31 string keys to change, 4 mock/data notes. Nothing blocks the layout.

## 1. Would I open it after a pull
Yes. Order (opening, Fix first with icon, Practise, Then, Doing well, Couldn't check) is what a good raid lead would say, and "Bloodthirst: 9 casts lost" is the first thing under the header. The strip under my own row is the right door.
Lines that grate (mock text from the product 5.3 examples; fix before the model prompt is written):
- Opening "Solid kill: you were alive the whole fight and Battle Shout was up 100% of it." and Doing well repeats "Alive the whole fight." and "Battle Shout up 100%". Same praise twice, and Battle Shout on a warrior is a trivial buff. Praise something hard (Death Wish or Flurry uptime, cast rate vs top logs).
- "A flask is the cheapest damage you will ever buy" is false in Classic (flasks are the expensive consumable) and reads as a lecture. Use: "A flask is a big gain for one boss; bring one to every 40-player kill."
- Footer "Reviewed 28 Oct" on a report dated "Thu 10 Dec" (mock only; a player screenshots that).
- Fix-first detail in the mock says "9 of a possible 40 went unused" but DetailCoreCastsOne says "{actual} of a possible {possible}", which would print "31 of a possible 40" under a title "9 casts lost". The detail must carry the title's number.
No "heal more", no grade, no exclamation marks found.

## 2. Strings to change (key | current | proposed)
"findings" is tool-speak; players say things to fix.
- rowPremium | `What to fix first: {title}` | `Fix first · {title}` (title already has a colon, so it prints two)
- teaserCount | `{n} findings` | `{n} things to fix`
- teaserCountOne | `1 finding` | `1 thing to fix`
- rowFree | `{n} findings · First: {title} · See your coach's review with Premium` | `{n} things to fix · First: {title} · Read the full review with Premium`
- rowFreeOne | `1 finding · First: ...` | `1 thing to fix · First: {title} · Read the full review with Premium`
- rowFreeNoFirst | `{n} findings · See your coach's review with Premium` | `{n} things to fix · Read the full review with Premium`
- rowFreeNoFirstOne | `1 finding · See ...` | `1 thing to fix · Read the full review with Premium`
- rowPremiumNoFirst | `{n} findings · Open your coach's review` | `{n} things to fix · Open your coach's review`
- rowPremiumNoFirstOne | `1 finding · Open ...` | `1 thing to fix · Open your coach's review`
- upgrade | `See your coach's review with Premium` | `Read the full review with Premium` (repeats on every free row, tab and account line; the long form is a nag and wraps mid-phrase at 360 px)
- lastRaidFree | `... {count} findings · First: {stem} · See your coach's review with Premium` | `... {count} things to fix · First: {stem} · Read the full review with Premium`
- lastRaidFreeOne | `... 1 finding ...` | `... 1 thing to fix · First: {stem} · Read the full review with Premium`
- lastRaidFreeNoFirst | `... {count} findings · See ...` | `... {count} things to fix · Read the full review with Premium`
- lastRaidPremium | `... on {n} of {m} bosses · Open` | `... on {n} of {m} bosses · Open review` (bare "Open" tells nobody what opens)
- lastRaidPremiumOne | `... {stem} · Open` | `... {stem} · Open review`
- lastRaidPremiumAll | `... on all {m} bosses · Open` | `... on all {m} bosses · Open review`
- lastRaidClean | `... Nothing to fix that the log can show · Open` | `... Nothing to fix that the log can show · Open review`
- lastRaidNone | `Last raid · none in the last 14 days` | `Last raid · no log in the last 14 days` (the site only knows uploaded logs; "none" claims you did not raid)
- lastRaidPending | `... Reading the fights...` | `... Reading your pulls...`
- loading | `Reading the fight...` | `Reading the pull...`
- rowPreparing | `Your {Spec} comparison is still being prepared.` | `Top {Spec} logs for this boss are still loading.`
- preparing | `Your {Spec} comparison is still being prepared. Check again in a few minutes.` | `Top {Spec} logs for this boss are still loading. Check again in a few minutes.`
- labelPractise | `Practise` | `Practice` (US client)
- starterHint | `This coach knows the rotation from our guide. Deeper class knowledge is on the way.` | `This coach reads your log against our guide's {Spec} rotation.` (drop the roadmap promise)
- night | `Reviews are one pull's. Pick a boss pull from the list to see one.` | `Coach reviews cover one pull at a time. Pick a boss pull from the list.`
- writingLate | `Still writing. Check later.` | `Still writing. Check back in a few minutes.`
- notReviewable | `Your coach reads boss pulls long enough to show a rotation. This pull isn't one.` | `This pull is too short to review. The coach reads boss pulls long enough to show a rotation.`
- StrengthCasts | `{Spell}: {you} a minute. Top {n}: {ref}.` | `{Spell}: {you} casts a minute. Top {n}: {ref}.` (number without unit)
- StrengthCastsSim | `{Spell}: {you} a minute. The written rotation, simulated: {ref}.` | `{Spell}: {you} casts a minute. Our guide's rotation, simulated: {ref}.`
- DetailCoreCastsOne | `{actual} of a possible {possible} with {resource} in your bar. Longest gap {gap1}.` | `{lost} of a possible {possible} went unused while you had the {resource}. Longest gap {gap1}.`
- DetailNoFood | `You pulled with {items}, and no food.` | `You pulled with {items} but no food.`
- raidNight | `Premium raiders: your coach's review of tonight is ready on Forever Sixty and in /report.` | `Coach reviews for tonight's pulls are ready on Forever Sixty and in /report.` (opening a guild channel post with "Premium raiders:" is an ad to everyone who is not one; the guild mutes the bot)

## 3. Free teaser
Count + first title + one button is tempting: it names a real fix and holds back the why and the drill. Annoying only because the button text is long and repeats on strip, tab and account line. "11 findings" on the account line is a sum over bosses with no unit; say "across 8 bosses" or show the worst boss, or it reads as a scolding.

## 4. Row strip and account line
Strip directly under my own row, 36 px, muted lead and bold title: right place, never expands the row. Coach hides in More beside Rating; the strip is the real door so acceptable, but show Coach as a visible tab for owners. Account line "none in the last 14 days": keep (reserved slot, no jump), reworded; mildly noisy for non-loggers, fine otherwise. The free account line wraps to two lines at 1440 in the board with the gold link broken mid-phrase; the shorter upgrade text fixes it.

## 5. Discord
The /report field "Your coach" above Loot is one useful line; alt lines with bold names are clear. I would read it. The raid-night sentence names nobody, which is right, but it only survives if it stops sounding like a Premium ad (above). It posts once per raid night, so frequency will not get it muted.

## Data I do not believe (for the accuracy reviewer)
- DetailAdherence example `Heroic Strike (31), Slam (12) and Rend (9)` as "outside the written rotation" for a Fury warrior: Heroic Strike is the Fury rage dump and Rend is not a Fury cast. Check the F8 rotation lists.
- "Cheapest damage you will ever buy" for a flask (above).
