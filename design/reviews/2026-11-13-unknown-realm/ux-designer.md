# Unknown-realm question: ux-designer verdict

Reviewed: question-390/1440/2000, question-close, answered-390/1440/2000, refused-close. Reference for the job: the in-game "choose your realm type" pick, i.e. one question, one row of equal choices, answer once.

## Verdict: FIX

The question is right: first thing in the plate, one sentence, four equal one-tap buttons, class-coloured crest name, no marketing. The answered state drops cleanly with no gap left behind (answered-390/1440/2000 start at "Sign in"). Two items must be fixed before ship.

## Must fix

1. Refusal sentence starts lowercase (refused-close.png): "another account already has a character of that name on that realm type. Try again." The component capitalises nothing. Capitalise the first letter of the API message in UnknownRealmQuestion.svelte before rendering (or have the API send a capitalised sentence). Result: "Another account already has a character of that name on that realm type. Try again."
2. Orphaned fourth button at 390 (question-390.png): Normal, PvP, Roleplay fill row one and Hardcore sits alone on row two. Make the group a 2 by 2 grid below the sm breakpoint (grid grid-cols-2 gap-2, buttons w-full), and keep the single row from sm up. Two even rows read as one choice set and give full-width 44 px-class touch targets.

## Should fix

1. Busy width jump: the pressed button swaps its label for the spinner plus "Saving…", which is wider than "PvP" and shifts its row neighbours. Give the four buttons a shared min-width (for example min-w-[96px]) or the 2 by 2 grid above, so nothing moves while saving.
2. The block has no mark that it is the next step. At 1440 and 2000 it reads as a fourth row of the plate, above a "2" step that carries the yellow active number. Add a 2 px left rule in the accent colour on the item (border-l-2, the active-step accent from Steps.svelte) so the eye lands on it first and it goes away with the answer.
3. Refusal placement: the error sits under the "RP PvP realms count as Roleplay." note, 40 px from the buttons. Put the error directly under the button group and the RP note last, so the failure is next to the tap that caused it.
4. Verify keyboard parity: Tab through the four buttons, focus-visible ring from Button on the dark plate, Enter and Space save. The captures show no focus state; attach one.

## Checklist

- [x] One question, answer visible without scrolling at 390 and 1280+ (it is the first element in the plate)
- [x] Reuses Button (size md, busy, busyLabel) and CharacterName; no new control
- [x] Four choices come from the one RULESETS list, site order
- [x] Wipe colour for the refusal, role="alert"
- [x] Answered state: block gone, no residual gap, 390/1440/2000
- [x] Centred at 2000
- [ ] Refusal copy capitalised
- [ ] 390 button layout without an orphan
- [ ] Focus-visible and busy captures attached

## Round 2

Re-read: question-390, question-close, refused-close, and UnknownRealmQuestion.svelte.

Verdict: SHIP.

- Refusal capitalised, now directly under the buttons, with the extra sentence "Check the realm type in game." (refused-close). Fixed.
- 390 layout is a 2 by 2 grid of equal buttons, no orphan (question-390). Fixed.
- Buttons are equal width (7.5 rem at desktop, 1fr on phone), so "Saving…" cannot move neighbours. Fixed.
- 2 px gold inset left rule marks the block as the next step. Fixed.
- RP PvP note folded into the muted line above the buttons. Fine; at 390 it wraps to three lines, acceptable.

Remaining should (not blocking): attach one focus-visible capture and one busy ("Saving…") capture from a real run; the static captures still show neither state.
