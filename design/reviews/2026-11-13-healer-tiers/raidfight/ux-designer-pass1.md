SHIP WITH FIXES

Reviewed: healer-gear and healer-raid at 390 fold and 1440 fold. Not reviewed: 1024, 2000 and the full-page captures (same header block, same 5-row table; the 1440 measure of about 60 characters is unchanged).

## What passes
- Both states say which fight they rank on, as the first sentence under the title, before any caveat. A fallback state that admits it is a fallback is correct.
- Layout holds: lead in the PageHeader `lead` slot, left aligned with the title, no new control, no pill, one "Updated 13 Nov" stamp. Nothing clips at 390 or 1440.
- Answer visible at 390 without scrolling: the table ends at y 737 (gear) and y 713 (raid) in an 844 px fold. At 1440 the table is fully in the fold.
- The ranking itself reads right in both states: 1 to 5 with vs #1 in raid, the "=1" ties with partners named in gear.
- Hierarchy is intact: title, lead, meta, then the table as the dominant object.

## Fixes
1. Phone lead is too long (must fix). At 390 the lead is six lines in the gear state (y 108 to 280, about 175 css px) and five in the raid state. The table starts at y 440 and 416. The earlier review asked to hold phone leads to short; this lead is now twice the earlier three lines. The shield and Power Infusion caveat is already in the "What it leaves out" column of the notes (HEALER_LEAD is reused there), so the lead repeats it. Cut the lead to the fight sentence plus one short caveat clause (target: three lines at 390).

2. Duplicate and inconsistent length (must fix). Gear state: the lead says "300 seconds", the meta line says "5 min", on the same screen. Raid state: lead says "110 seconds long", meta says "110 s". The length is stated twice in two formats. The meta line already owns the length; take it out of the lead in both states.

3. Gear sentence talks about the schedule (must fix). "until the next ranking adds a logged raid kill" is site meta talk and promises a date. Say what the list is on and what replaces it, with no "ranking" wording. "heavy damage" is vague and the number it sits next to is dropped by fix 2.

4. Raid sentence does not name the fight (should fix). "a logged raid kill" is the thing a Molten Core healer wants named, and the names are stable, not nightly. Name them.

5. Meta line wording (polish). "Best in slot, raid buffs" then length is fine. Keep "5 min" and "110 s" only if one formatter produces both; otherwise a 300 s fight reads "5 min" and a 90 s fight would read "90 s" on the same page family. Use one rule: under 2 min in seconds, otherwise minutes (so 110 s and 5 min stay as they are).

## Replacement copy (verbatim)

healerFightSentence, raid:
"Ranked on one healer's share of a Molten Core or Onyxia kill."

healerFightSentence, gear fallback:
"Ranked on the BiS fight until a logged raid kill replaces it."

HEALER_LEAD as used in the lead only (the notes column keeps the current long form):
"HPS is effective healing. Shields count for what they soak, and buffs on others are not counted."

Result, raid at 390 (three lines expected):
"Ranked on one healer's share of a Molten Core or Onyxia kill. HPS is effective healing. Shields count for what they soak, and buffs on others are not counted."

Result, gear at 390:
"Ranked on the BiS fight until a logged raid kill replaces it. HPS is effective healing. Shields count for what they soak, and buffs on others are not counted."

(Use a straight or curly apostrophe to match the house style; no em or en dashes anywhere. The Power Infusion example stays in the notes column.)

## Checklist
- [x] Page says which fight it ranks on, fallback state
- [x] Page says which fight it ranks on, raid state
- [x] Answer visible without scrolling at 390 and 1440
- [x] No overflow, clipping or layout shift between states
- [x] No nightly figure in the lead
- [ ] Phone lead at three lines (fix 1)
- [ ] Length stated once, in the meta line (fix 2)
- [ ] No schedule wording in the fallback (fix 3)
- [ ] Raid fight named (fix 4)
- [ ] One length formatter (fix 5)
