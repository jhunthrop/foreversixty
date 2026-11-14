# Store listing round 1 (addon 2.0.48 page text and the ten mock plates)

| Reviewer | Verdict | Must-fix, in short |
|---|---|---|
| wow-player | FIX | the "no messages sent" and "window closes in combat" claims (Remind posts to guild chat; no combat close in code); the repeated Striker's Mark roll; raid dates (last raid before the 9 Dec opening, Rotation's 15 Dec listed as a future raid); the Beaststalker's Gloves tooltip (bind type, same armor as the worn gloves); Gear and Rotation weights labels and units disagree |
| ai-slop-reviewer | FIX | the tagline "Play your class better."; the gloves gain +11.2 on the line and Me against +2.7 on Gear; the countdown 1 h 12 against 3 h 12; Rotation's Tue 15 Dec; Guild and Officer role targets disagree; the blank minimap disc on plate 1 |
| brand-consistency-reviewer | FIX | the character contradicts itself across plates (gain, countdown, sign-up, roles, dates, level 42 on Builds); the caption block moves between one-line and two-line captions; plate 1 unfinished (olive ground, blank disc, small) |
| ux-designer | FIX | the same cross-plate contradictions; plate 1 weak as the featured image; dead stage on plates 8 to 10; caption 8 against plate 8 (gold corner, set noise); the lead JPEG shows an old surface; "What's new" before the tabs |

## Fixed

- `gen_captures.py` gained a store fixture: one evening of one character (raid day Thu 17 Dec, In, 3 h 12 min, last raid Tue 15 Dec with the Rotation lane's boss results, the bag upgrade Beaststalker's Gloves +2.6 everywhere, Me's sim gains and the Officer gap from the Gear tab's model, Builds at 60 on the Beast Mastery raid build, the copy code's guild Thornwake Company).
- Guild and Officer read one wanted-roles table (`addon_hub.WANTED`) in `addon_hub_l5.py` and `addon_hub_l7.py`.
- Plate 1 is the line and its hover only; plate 8 drops the bag item's set tooltip; plate 9 drops the repeated roll; plate 10 shows the stale record and the Shift view.
- `gen_gallery.mjs` pins the caption block with a two-line minimum.
- The page: the closing claim names Remind and Send to guild and drops the combat claim; Gear "from level 10" and "More upgrades"; Officer whisper and invite; the lead JPEG moved to the marks section; "What's new" after Settings; the line's example uses the plates' item and gain.

## Kept

- The tagline is the owner's positioning line, outside this lane.
- The Gear weights header and labels (Attack Power and Ranged Attack Power both listed), the spell hit cap line and Rotation's Aimed Shot condition are the shipped 2.0.48 addon's output; reported to the owner as addon follow-ups, not changed in a store screenshot.
- Item details (bind type, armor, level) come from the build's item data.
