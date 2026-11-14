# brand-consistency-reviewer, round 2 (Me with no site activity)

Verdict: PASS

The round-2 changes bring the Me tab closer to the hub and the web. I found no must-fix items. I checked the re-rendered `me-empty-l25.png` and `me-empty-l60.png` and spec section 12. I did not open the `-l5` or `-point` captures this round, or the code.

**Must-fix**
None.

**What now matches**
1. **Gear upgrade rows:** they now use the Right now anatomy: a 20 px icon, the name in its quality colour, the gain in green at the right, and the muted source line under it. The source line (slot, source, drop chance) reads like the web's BiS "where it comes from" detail.
2. **Enchant rows:** they are the Gear tab's own shared `GearEnchants` rows, so Me and Gear show the same `No enchant: <enchant>` string with the enchant's spell icon. That is the same component doing the same job.
3. **Your sim:** "Not simmed yet" is muted, with a primary button reading "Run a sim". That is consistent with the button states ("Sim again" once a result exists).
4. **Level 60:** dropping the Experience row is right, since the header already says "Level 60". Build and Professions fill the column, and the one-line site copy is shorter and calmer.
5. **Site copy:** the site line is plain game words, has no em dashes, and keeps the gold link as the only gold in the Me column.
6. **Colour, gain format and "best in slot":** unchanged and correct.
7. **Row anatomy:** label, value and right-aligned note are unchanged across levels.

**Shoulds**
1. **Truncated enchant rows (Level 25):** `No enchant: Enchant Chest - Minor Stats · Che...` and `... Minor Agility · ...` cut off the slot name, which is the part the player acts on. The enchant name is long and the slot comes last. Consider truncating the enchant name first so the slot always survives, or dropping the repeated slot when the enchant name already contains it ("Chest", "Bracer"). Do this in the shared `GearEnchants` row so the Gear tab gets the same fix.
2. **Red text on whole rows:** both the enchant name and the slot are warning red, so the full line is red. This is the shared Gear tab styling, so it is consistent in the addon. If the web's missing-enchant row keeps the enchant name in body text and only the "No enchant" label in warning colour, align the shared row to that.
3. **Professions row:** `Skinning 112, Leatherworking 105` is a good fact, but a long list will wrap. The spec caps the section at 5 rows. Confirm the value is cut to fit like the Build name.
4. **Spec note:** section 12 records that the hit-cap row, the Build hover and the talent tooltip were dropped by decision, and that the tab does not redraw live on XP. These are acceptable. Row 3 ("Next point") still shows a talent icon but, per section 12, no spell tooltip. A player who hovers it will notice, so it is worth adding when there is time.
