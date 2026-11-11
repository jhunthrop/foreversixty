# Round 1: what changed for round 2 (Android)

## Fixed

**Ink on the ground.** Secondary lines that sit on the painted ground now use `mutedGround`, the web's token for muted ink on the earth ground (AA on its lightest window); `muted` stays for panels. This covers:
- the raid head (guild, zone, the Invites/Pull/Ends labels);
- the account version and About lines, and "Signed in with Battle.net";
- the guides footers and not-cached lines;
- the devices and offline footers;
- the delete-email line;
- the item page's foot line.

**Heads continue the band under the bar,** as the web's page head does: the guide, the character, the raid page, and the BiS list (now a band head with the spec in its class colour, `Best in slot`, and a one-line control bar with `Change` in place of three pills).

**Errors.** One `LoadError` primitive (a Plate with wipe ink and Try again) replaces four local error blocks, so no error is red on the ground.

**Section heads** over the ground are SectionTitle (the display face, cream) on the raid, guild and guides screens.

**The raid flows** (raid night, applicant, unanswered) are on the ground with FsTopBar (FlowBackBar now draws it). Their tests paint the repo's game icons and raid icon build, as the tab roots' tests do.

**The answers** In, Late, Bench and Out on the raid page draw as the tab roots' FlowSegmented (teal edge), the same everywhere.

**Shapes.** Chips are on the control radius and tags on the field radius; there are no full pills. The crest glow is gone (the web's crest is a ring only).

**Segments.** Unselected segments sit on an opaque panel with body ink.

**Sim result.** The answer is a tabular Figure (figureXl) on a Plate with its error. The engine hash is cut. A Sentence answer uses `itemLg` (the retired `stat` style is gone from it).

**Tier list.** It follows the web's anatomy: the tier letter at the head of its group, rank, icon, name and race, and the DPS over the gap on the right. The bars are gone.

**Character.** The order follows R7: Gear, then the fights, then Forever Score. Percentiles are ordinal ("99th percentile"), here and on reports.

**Explainers cut:**
- the sims footers;
- the "Gold corner" and "Gold dot" legends;
- "Close the sheet and nothing is saved";
- "Raid sign-ups ... free for everyone";
- the talents footer.

**Copy.** "Unlock sign-ups" now reads "Open sign-ups".

**Captures.** The delete mismatch, ready and failed states are also captured scrolled to the end (`account_delete-*-end.png`).

## Checked and left as is (with the reason)

- **Sapphiron Drape from Onyxia:** correct. Item 17078 drops from Onyxia in Classic.
- **The applicant's faction:** the fixture guild is Alliance (the Gear capture reads "60 raid · Alliance"). Mirelle (the applicant) and Mirethil (a raider) are two people.
- **Warrior and hunter names in the tier list:** these are the class colours (#C69B6D, #AAD372), not white.
- **"Tuesday" as a sim's name:** that is the name the player gave the run. A row shows its title.

## Follow-ups (product or data, not this restyle)

Each is named here so it is not lost:
- **Stat weights on the guide.** The guide file carries only `stat_priority`. Weights live in the BiS file per band, so the guide reader needs that join and the web's band choice. Its order against the web is a data question for the guides pipeline.
- **The BiS paperdoll for a signed-out class (R6).** The signed-out Gear path shows the list. The paperdoll needs the Gear tab's composer without a character.
- **The character's head answer** (+DPS available, Sim this character). This needs `/v1/home` for a visited character.
- **A GB price in pounds** (the store's localized price), and item tooltip rating wording ("Critical strike rating"). Both are data.
- **Fire Nova's missing spell icon.** The build data has no icon for that spell id.
- **Info-density additions:** a Recently updated strip on Guides, a spec rank for Forever Score, sim deltas and stat weights in the app, drop rates on sources, and device platform. Each needs data the app does not load yet.
