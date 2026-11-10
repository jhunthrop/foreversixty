# Round 2: what changed for round 3 (Android)

The sheets are now full-colour JPEGs: `android/<group>-app-N.jpg` and `-reference-N.jpg`. The round-2 PNGs were packed to 256 colours, which turned small hues grey. That is why the kill dots read grey: the code passes `FsColor.kill`, and it renders green in the goldens.

## Fixed

**Ink on the ground.** `mutedGround` is now on every remaining line that sits on the ground:
- FlowNote by default, which covers the raid night aside and Not answered's notes;
- the raid page's count line and state notes;
- the report meta and footers;
- the talents legends;
- the tier note and the guides "not saved" line;
- the BiS set line;
- the Your builds note.

**One answer control and one primary.**
- Raid night, Home and the raid page all draw FlowSegmented.
- FlowButton's primary is the teal button, never the gold gradient.
- Not answered's filters are FlowChips.
- Reset and Cancel are the outline (Quiet) button.

**The raid flows' black body.** The test harness drew them inside the tab Scaffold's opaque background. It now paints the ground, as the routes always did.

**BiS:**
- the failed and uncached states are the LoadError Plate;
- the picker title is SectionTitle;
- the search and roster chips use the control radius.

**Crests.** Tier and roster rows use the ringed class crest. The tier list has a DPS column head.

**Dates** read as the web writes them: "28 Oct" and "Updated 1 Nov".

**Account:**
- Signed out is a band head, then one Plate holding the one teal primary, with Pick your class as the outline button.
- Plan shows the price as a gold Figure with its period, then the renewal.
- Delete puts the confirmation (field, failure, button) on a Plate.
- Disabled controls are at 50%.
- The duplicate Premium pill on You is gone.

**Explainers cut:**
- the devices footer;
- the founder note;
- the lock note;
- the guides footer's second sentence.

The search placeholder is shorter.

**Captures:**
- You at 2x scrolled to its end (`premium_you-premium-ax5-end`);
- Offline data clears the bar.

## Answered, not changed
- **The GB plan's "$40 a year".** A web subscriber pays the web's price, in USD, wherever they live. The app shows what they pay. A store-priced plan would be a product change.
- **The guild's faction across captures.** Each test suite has its own fixture (the raid page, the applicant, the Gear tab). In the app, all of them read the one guild from the API.
- **"1 0 1 · kill wipes new best" against the report's wipes.** These are two fixtures for two different nights.
- **Your guide (Elemental) against Your spec (Fire Mage).** The two screenshot states use two test characters. In the app both come from the one current character through `Guides.yourGuide`.
- **The tooltip's "+28 Crit".** The item text comes from the item data's stat vocabulary (data follow-up).
- **Fire Nova's blank tile.** The guide data carries a spell id but no icon (guides pipeline follow-up). The neutral tile is an earlier ruling ("a spell with no icon keeps the neutral tile").
- **The default band "opens Dec 9".** The selection logic follows the product's band rules, not this restyle.

## Declared drift
`android-raid-night-live` and `-signup`: the sections under the answers sit 11 to 14 dp lower than in the mock. The mock still draws the old small buttons (`ws_app_flows._signup_buttons`), and that generator is shared with the iOS Home mocks that lane ios-restyle owns. The drift is dated and named in the test, and the test fails once it is gone.
