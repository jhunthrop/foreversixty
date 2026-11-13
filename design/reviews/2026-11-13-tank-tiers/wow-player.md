# wow-player final review: Tanks tab (/tiers/tank), built fix

Verdict: SHIP.

Seen rendered: after-1440.png and after-390.png. Copy read in web/src/lib/tiers/tier-copy.ts (tankColumns).

## MUSTs from the proposal verdict
1. Health = real max health: done. Bear 9,171, paladin 7,330, warrior 7,484 sit in the range of a 60 character panel. Bear green, correct.
2. Leader reads "lowest": done. Head stays "Taken / s": accepted, the rail note defines it as damage taken and 96 px will not hold the long form.
3. Rail note defines Taken / s, Health, Threat / s, vs #1, the bar and the green rule: done in copy. Subtitle "vs a level 63 boss": done.
4. 351 vs 352: done, "=2" and "≈ tie with ...". Both rows read "+53% taken", so the tie is believable.
5. Accuracy answers accepted: no Salvation, Righteous Fury on, Thunder Clap 20% slow, warrior lead explained by crushes, crits, stance and parry. Bar widths check out (230/351 is the bar on the bear).

## Open, not blocking
- Bear threat 2,027 over warrior 1,677 will still be mocked by Classic players, and the green on it asserts it. The sim report leaves it open for a threat pass. Run that pass next, and until then do not write any prose that says bears out-threat warriors.
- The 1440 capture shows no rail note, only "Sign in to mark your spec". I could not see the note rendered, so confirm in the live page that it opens and that the bar's aria or tooltip text is present.
- 390: the only column head is "Taken / s", and there is room for "Damage taken / s" there.
- Still missing: avoidance and crush % (SHOULD from the proposal), and the page is empty below three rows.
- Health column stays hidden on live until the ranker's next run publishes the field; fine, nothing wrong shows meanwhile.
