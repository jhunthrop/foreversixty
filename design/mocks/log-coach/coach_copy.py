"""The log coach's final copy, for the mock boards. The experience spec's copy table
(design/specs/2026-10-28-log-coach.md, section 8) carries the same strings with their keys; the Go and web
tests read the spec. Every number here is the product spec's section 5.3 invented example, not a measurement."""

ROW_PREMIUM = 'Fix first · {title}'
ROW_FREE = '{n} things to fix · First: {title} · Read the full review with Premium'
ROW_FREE_ONE = '1 thing to fix · First: {title} · Read the full review with Premium'
ROW_FREE_NO_FIRST = '{n} things to fix · Read the full review with Premium'
ROW_PREMIUM_NO_FIRST = "{n} things to fix · Open your coach's review"
ROW_CLEAN = 'Nothing to fix on this pull that the log can show.'
ROW_OFFICER = 'Coach review'
ROW_PREPARING = 'Top {spec} logs for this boss are still loading.'
ROW_ERROR = "The coach couldn't read this pull just now."
ROW_LOADING = 'Reading the pull…'

LOADING = 'Reading the pull…'
WRITING = 'Your coach is writing…'
BUSY = ("Your coach is busy right now. This is the review from your log alone; the coach's version will be here "
        "when you next open it.")
ERROR = "The coach couldn't read this pull just now."
STARTER_HINT = "This coach reads your log against our guide's {spec} rotation."
RULES_NOTHING = 'Nothing to fix on this pull that the log can show.'
RULES_OPENING = 'Here is what your log shows for {boss}.'

LAST_RAID_PREMIUM = 'Last raid · {zone}, {date} · Fix first: {stem}, on {n} of {m} bosses · Open review'
LAST_RAID_FREE = 'Last raid · {zone}, {date} · {count} things to fix · First: {stem} · Read the full review with Premium'
LAST_RAID_NONE = 'Last raid · no log in the last 14 days'
LAST_RAID_PENDING = 'Last raid · {zone}, {date} · Reading your pulls…'

DISCORD_LINE = "Fix first: {stem} — on {n} of {m} bosses. [Open your coach's review]({url})"
DISCORD_PENDING = "[Your coach's review is on the site.]({url})"
RAID_NIGHT = "Coach reviews for tonight's pulls are ready on Forever Sixty and in /report."

# Fury example, product spec section 5.3 (invented values), rule-only mapping per section 5.2.
FURY_TITLE = 'Bloodthirst: 9 casts lost'
FURY_DETAIL = '9 of a possible 40 went unused while you had the rage. Longest gaps 1:12–1:25 and 2:40–2:51.'
FURY_DRILL = ('On a target dummy, press Bloodthirst the moment it comes off cooldown, before anything else; '
              'Whirlwind and Heroic Strike wait their turn.')
FURY_OPENING = 'This was a clean kill. Almost everything you left behind sits in one button.'
FURY_WHY = ('Bloodthirst is the button the Fury rotation is built around, and 9 of a possible 40 went unused while '
            'you had the rage, longest at 1:12–1:25. At 2:40–2:51 you sat on a full bar instead, so that one gap '
            'cost you twice.')
FURY_THEN = [
    ('No flask at the pull', 'You pulled with Grilled Squid and an Elemental Sharpening Stone, and no flask.',
     'A flask is a big gain for one boss; bring one to every 40-player kill.'),
    ('Rage full for 14 seconds', 'Your rage sat at its cap for 14 s outside the air phase, longest at 2:41–2:50.',
     'When the bar is full, Heroic Strike turns rage into damage instead of losing it.'),
]
FURY_WELL = ['Alive the whole fight.', 'Death Wish up 94% of the fight. Top 25: 91%.']
FURY_CANNOT = ["Talents and race: Forever logs don't record them.",
               "Range and movement: the log doesn't say where you stood."]
FURY_FOOT = 'Reviewed 10 Dec · Based on: Top 25 Fury warriors on Onyxia (40), your log.'

FURY_LINK = 'The Bloodthirst gap at 2:40–2:51 overlaps the rage cap at 2:41–2:50.'
