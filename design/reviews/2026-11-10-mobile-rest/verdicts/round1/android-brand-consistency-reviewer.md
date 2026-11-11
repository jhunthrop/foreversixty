# brand-consistency-reviewer, Android secondary screens, round 1 (2026-11-10)

All six groups are DRIFT. Three problems run through all of them:
- **Muted text sits loose on the painted ground.** The web never does this. Move these lines onto a Plate, or use the `text` or `mutedGround` role, and keep `wipe` only inside a LoadError Plate.
- **Chips are full pills.** `pillSurface` defaults to 50% while the tokens say 3 to 4 px. Use `FsRadius.control`, or Segmented.
- **The crest glows.** `BrandImages.kt:56` draws a 7 dp class shadow, and the spec allows no glows.

| Group | Verdict | Findings |
|---|---|---|
| guides | DRIFT | **BLOCKER:** the stat priority is numbered chips with no weights, in a different order from the web. Use the rank rows with a Meter and the weights. **BLOCKER:** the tier list lacks the letter column and the DPS and % figures, and its meters are not on the web. **SHOULD:** a square spec icon where the web uses the crest. **SHOULD:** white class names in the tiers (TierLists.kt:199). **SHOULD:** the reader's section heads should be SectionTitle. **SHOULD:** the date formats. **SHOULD:** the unselected segments. **NICE:** a gutter, and Fire Nova's icon. |
| character | DRIFT | **BLOCKER:** the order breaks R7: no head answer (+DPS, Sim this character), and the score sits above the fights. **SHOULD:** the kill and wipe dot colours change between states. **SHOULD:** the error state is not a Plate. **SHOULD:** the rankings filters are pills. **NICE:** the Epic tag. |
| bis-search | DRIFT | **BLOCKER:** the BiS list is not round 5. It has no band, no Marcellus head (`BisScreen.kt:148`), no paperdoll, and three pills instead of a ControlBar. **SHOULD:** the picker title is in Barlow. **SHOULD:** the item meta on the band, and red error text on the ground. **NICE:** the hex in `Paperdoll.kt:165`. |
| account | DRIFT | **SHOULD:** signed out is all band with two filled buttons. **SHOULD:** the Plan price sits in a sentence instead of a gold Figure. **SHOULD:** illegible muted lines. **NICE:** a generic person glyph on You. |
| guild-raids | DRIFT | **BLOCKER:** In is green on the raid page and gold on raid night. **BLOCKER:** raid night, applicant and unanswered are without the ground and the bar. **BLOCKER:** empty buff, spec, loot and raid icons. **SHOULD:** the primaries clash (blue against gold). **SHOULD:** square spec icons on roster rows. **SHOULD:** muted labels on the ground. |
| sims-reports-talents | DRIFT | **BLOCKER:** the sim answer is in Marcellus and should be a tabular Figure on a Plate. **SHOULD:** Cancel and Reset are filled primaries. **SHOULD:** grey kill and wipe dots on reports. **SHOULD:** the talent trees sit on the bare ground. **SHOULD:** illegible meta on the ground. **NICE:** a warrior crest on zone reports. |
