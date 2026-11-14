# AI-slop gate: loot-plan, live-pulls, prep-reminder, guild-later (spec-level, 2026-11-15)

No screenshots; judged on specified copy and layout. Em dashes: none found in any of the four. Emoji: none.
All-caps headings: none. Icon grids, stat tiles, gradients, blur: none specified. Numbers come from data, not invented.
Overall: reads as written by someone who raids (real item names, Onyxia, Details!/RCLootCouncil references, bands not figures).
The slop that exists is repeated disclaimers, explainer sentences and double labels, plus one glow.

## 1. Loot plan: PASS WITH FIXES
Offending lines (section, tell):
- 5.1 Notice `This is a recommendation. The loot council decides.`, 5.2/12 raider line `The council decides who gets it.`, 6 Discord
  `A recommendation. The council decides.` Same disclaimer three times; veterans know who decides (tell 21). Keep it once, on the page only.
- 12 group captions `Estimate in DPS. Ties go to whoever has won least this tier.` (x3, one per group). Grey explainer under every
  heading (tell 7), and rule 2 already puts `estimate` in the group head (22). Drop the tie sentence, keep `DPS` / `HPS` / `Tank score`.
- 3 rule 5 `Item level upgrades` caption plus label `Upgrade by item level`: same fact twice. Keep one.
- 5.1 `Gear read 14:02, 3 h ago`: time and age together (22). Use `Gear 3 h ago`.
- 5.4 `Nobody has signed up yet. The plan fills as raiders sign up.`: second sentence explains the obvious. First only.
- 5.4 `Some gear is older than your last raid.` Vague; name the count: `7 older than the last raid`.
- 5.1 candidate row carries BiS Tag, MS/OS Tag, source Tag, plus Old gear Tag: four chips per row (17). Fold source into the age text
  (`Addon 2 h`), show BiS Tag only when true, show MS/OS only when declared.
- 3 `No data` (raider status) and `No gear on file` (officer) both mean the same; pick one word.
- `Likely upgrade` / `Possible upgrade`: hedge-flavoured; ok only because tier 0 vs 1 is real. Prefer `BiS` / `Item level`.
Must-fix: remove repeated disclaimers, group-caption tie sentence, double tier-1 label, row chip stack.

## 2. Live pulls: PASS WITH FIXES
Strongest of the four. `Pull 4 · Onyxia · wipe at 31% · 3:12`, `Death Wish used once in 3:12, could be twice`,
`Thandrel was first, at 2:41` are what an addon author writes.
- 5.1 empty `No pulls yet. The first one shows here when it ends.` Second sentence is filler (21). `No pulls yet.`
- 5.1 no companion `No one is logging this raid. Run the companion on one PC and pulls show here as they end.` Explains the product.
  `No one is logging.` plus the `Get the companion ›` action.
- 5.2 `Most costly raider-side loss: downtime, 5 raiders over 15 s` is jargon, awkward ("raider-side"), and a third name for the
  same list (`Cost most`, `Most costly`, `What cost us`). Pick `Cost most` everywhere.
- 2 unit `lost seconds` is invented vocabulary shown to players; show seconds only, e.g. `Fireball: 9 deaths, 41 s`.
- 5.1 Reading state has two strings (`Reading the pull...` and `This pull is taking a while to read`); hedge-y. Keep the first, add `Retrying` only on failure.
- 5.4 `Couldn't load the pulls. Pull to retry.` vs web `Lost the connection. Retrying.`: contractions differ from the rest of
  the set (`Could not`). Match.
Must-fix: the empty and no-companion explainers, one name for the cost list.

## 3. Prep reminder: FAIL (tell 1)
- 4.4 toast: `A toast near the minimap button (Theme glow pulse, no chat output, no sound)`. Glow pulse on a notification is tell 1
  and cannot pass. Replace with a static toast that fades in and out, or a plain flash of the button border.
Other lines:
- 4.1 DM `Not a live check. Open the game and the addon checks your bags at 7:00 PM.` repeated in three variants plus
  `Last upload 4 days ago`: age stated three times, caveat explains the obvious (22, 21). Keep `Bags 4 days ago`; one footer line.
- 4.1 footer `Gear from Blizzard 3 h ago · Turn this off in Settings` repeats the Gear field's `(3 h old)`. Age once.
- 4.3 officer no-relay sentence: `No relay online. Nobody running the companion is in game, so bags are checked only by each raider
  in game. Ask a raider who runs the companion to log in.` Three sentences, "in game" twice, tells the officer what they know. 
  `No relay online.` with `Ask a raider with the companion to log in` as the action.
- 4.3 foot line `Gear from Blizzard 3 h ago · bags from your upload 4 days ago` plus row ages plus the chips: same ages three places.
- 4.3 chips `7:02 PM · Game`, `Bags 4 d · Upload`, `Gear 3 h · Blizzard` on every officer row: time plus source plus age on each
  row is chip density (17). Show the stamp only; put source in the row's tooltip.
- 4.2 iOS `Alerts on iPhone are not on yet. Prep still reaches you on Discord and in Activity.` Marketing-flavoured apology;
  `iPhone alerts are off.` is enough.
- 10 `Checked 7:00 PM in game` / `Offline in game` / `Bags not read` / `Not read` / `Attunement not read`: "in game" qualifier
  where the surface already says so (22); five phrasings for one unknown state (consistent with never-missing, but trim).
- 6.3 tooltip `Reloads your UI so the companion can send everyone's checks. Out of combat only. Takes a few seconds.` Fine in
  substance; drop `Takes a few seconds.`
Must-fix: remove glow; cut repeated ages and caveats; one no-relay sentence.

## 4. Guild later (comp/bench, recruiting): PASS WITH FIXES
Short, specific, real buff names. `Bench Brenn: lose Strength of Earth Totem (melee group 2)` and `Bench call: Brenn sits, no cover lost.` are good.
- 1 band `about 30 to 40 DPS raid-wide`: `about` is a hedge on top of a band, and `raid-wide` is a qualifier the row already implies (22). `30 to 40 DPS`.
- 1 `Healers: coverage only`, `Not tracked` vs the loot plan's `No data`: unknown wording differs across features. Align.
- 1 `The raid leader decides` and the loot plan's council line: disclaimer again (21). Drop here; the bench list is the answer.
- 2 answer line says `Invite on the site` is the only action, but the same section says there is no site message and the action is a
  `/w` copy line. Label the button `Copy /w` (or `Whisper`), not `Invite on the site`. Copy that promises a thing the product cannot do.
- 2 `Seen 2 d ago` while the rules say the age is up to 72 h; fine, but `Seen` reads as activity tracking. Use `Profile 2 d`.
Must-fix: the `Invite on the site` label; drop the leader disclaimer.

## Reads as authentic (keep)
- Real names everywhere: Ashkandi, Deathbone Guardian, Onyxia, Fireball, Deep Breath, Windfury Totem, Sunder Armor.
- Bands instead of false-precision numbers; unknown never printed as missing; names to officers only.
- References to RCLootCouncil, Details! death log, Recount, raid frame; "Not borrowed" lines.
- Lock/Drop/Swap with undo and keyboard D/S/L; 2 px gold selection rule used by role.
- Plain Discord posts (counts and mechanics, no names, no bands); the first-screen answer lines in live-pulls.
