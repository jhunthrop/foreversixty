# ux-designer, round 1 (captures in r1-after/)

VERDICT: DO NOT SHIP

The /sim bug is fixed. The /bis answer is still four rows and breaks "figure plus one line".

1. **BLOCKING**, bis-landing-34-1440: the answer still stacks crest and name, the level line, "0 of 14 slots" with its Meter, and "Best-in-slot pieces you wear", about 100 px beside a one-line title. Fix: the figure and its Meter on one line, with at most one short line under them. Move who-it-is out (the nav chip and the "Yours" row already say it). Move the caption into the tooltip. Keep the action beside the answer.
2. **BLOCKING**, bis-landing-34-390: the same answer is five rows on a phone, and the empty Meter reads as a stray dark rule detached from the figure. Fix: a full-width 6 px track under the figure, with a visible track colour.
3. Non-blocking: the bis answer's right edge is ragged.
4. Non-blocking, sim 1440: the actions don't centre on the answer's last line, and that line sits at a different height with and without the Tag. Fix: a fixed 24 px caption row with the actions centred on it.
5. Non-blocking, sim 1440: the actions float mid-band with about 130 px of gap before the answer.
6. Non-blocking: the sim head is about 220 px against about 150 px on home.
7. Non-blocking: the Tag's sentence is hover-only on touch.
8. Non-blocking: the crumb is 12 px underlined text with a small target. Fix: 14 px, a 32 px hit area, no underline at rest.
9. Non-blocking: empty icon squares on the /bis spec rows.
10. Pass: home (no dead art, a linked context line), guild raid (crumb on the band, tabs and bar on one baseline), guild claim (flush), and sim title/label alignment at 1440 and 390.
