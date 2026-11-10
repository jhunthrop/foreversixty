# Round 2 response (lane h-doll)

## Fixed for round 3
- **One off-hand word everywhere:** "Two-handed" (`gearCopy.paperdoll.twoHanded`), shared by the frame's list form and the report list (slop 3, wow 2, density 1).
- **Long enchants:** from 560 px (two wide columns) a list row keeps one line and carries the whole line in its title. The report's side list does the same at every width; the item's tooltip holds the full enchant. A phone still wraps, so the enchant is never cut where there is no hover (slop 4, ux 1, density 2).
- **The stats note:** it now sits where the stats go and says what waits: "Your stats in the fight show after a run on a newer engine." (slop 1, wow 3).

## Kept, with the reason
- **"Item level" on each row (slop 2):** these are the client's own words ("Item Level 61" on every tooltip). A bare "61" after the slot reads as a count.
- **Enchant green (slop 5):** the client's enchant colour, kept as is.
- **"Item level 60.0" on the strip (wow 1, density 3):** the strip's identity line predates this lane, and the sim head prints the same figure. Routed.
