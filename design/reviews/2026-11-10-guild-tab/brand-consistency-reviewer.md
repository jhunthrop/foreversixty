# brand-consistency-reviewer: the Guild tab (iOS, Android) vs the web guild pages at 390 px

Date 2026-11-10. Reference: web-390/*.png, web/src/lib/guild/assignments-view.ts, web/src/lib/nav.ts, design/tokens.
Reviewed: ios/dark, android/dark, parity/ios-full-*, GuildTabViews.swift, GuildTabParts.kt. Dark only (the apps and the site are dark).
Limit: I read the Assignments rows from code and the parity render; the officer Assignments block is not in a still I could open.

## Passes
- Tab icon: `tab-guild` is baked from design/assets/icons/ui/guild.svg (the glyph nav.ts gives Guild, icon 'guild') by gen-ui-icons.mjs, 24 pt / 24 dp. Thicken 52 on a 1886 box is 0.66 dp, the same weight as Activity (22 on 823). Selected gold, unselected muted, badge in the same place as Activity's, on both platforms. One family with Home, Gear, Activity. `tab-raid` is gone.
- Type: guild name in Marcellus (FSFont.title) in cream, mixed case as authored; section titles Marcellus; body and labels sentence case; no forced caps. Figures tabular.
- Colour: FSColor tokens throughout (cream, muted, gold, text, class colours); no hex in either file. Quality borders, spell icons and ringed crests as on the web.
- Copy: Your job, Assignments, `{boss} assignments`, `{label} with {names}`, `{n} of {count}`, `No one is picked yet.` all come from shared.assign.* (the same keys web's assignments-view.ts reads). Section order Before you pull, Roster, Loot, Recent logs, Progression follows the web's job order. No em dashes.
- Surfaces: hairline panels, no gradients, glows or shadows beyond the token set. Native tab bar and large title are rule 8.

## Findings
1. SHOULD. Assignments names (iOS `AssignedNames.text`, GuildTabViews.swift ~269-283; Android GuildTabParts.kt ~215-219) colour every pick in class colour and own in gold. The web planner's Tag is gold for yours, warn for an off-role pick, muted for the rest (`nameTone` in assignments-view.ts). In the apps a full slot reads as a rainbow, and the web's quiet hierarchy (yours stands out) is lost. Fix: others in FSColor.muted (the Tag's muted tone), own in gold. Off-role is not carried in the app models (`off_role` does not appear in apps/); either add it with the warn tone or record the omission in the spec, since read-only members rarely need it.
2. NICE. Your job is a top-level section ("Your job") above Who is coming on the Guild tab. On the web member view it is one quiet line, `Your job: Whelps, west cave`, inside the Onyxia assignments block, and the page leads with the sign-up. Rule 7 (do not promote what the web keeps quiet). Spec 2.3a (owner addition) rules for the placement, so this is ALLOWED as a ruling; the row text carries the web's words. Suggest the title row stays but nothing louder than the row's current weight.
3. NICE. Level spread (SpreadFigures): figure above label, `30 at 60 / 5 55 to 59 / 58 you`. The web's figure blocks (Your week) put the label above, the figure large, the caption below, and the spec asked for the line `12 at 60 · 9 close · you're 36`. The app uses `55 to 59`, not the spec's `close`. Fix: pick one of spec or build and keep the web's label-over-figure order for figures (Copy.guildTabYou etc.).
4. NICE. Guild head line `38 members · Horde` vs the web head's `US PvP · 40 raiders`: the web leads with ruleset and counts raiders. The parity render adds `US PvP`; the officer and android captures do not (fixture world differs, 3 members). Fix: confirm both surfaces read the same count and word (`raiders` vs `members`) from guild home, and always include the ruleset as the web does.
5. NICE. Hardcoded sizes outside the type scale: `ScaledType.sectionTitle.copy(fontSize = 22.sp)` and `FsType.figureSm.copy(fontSize = 20.sp, lineHeight = 24.sp)` (GuildTabParts.kt), and `size: 22` / `size: 20` TypeRoles (GuildTabViews.swift). They match each other today but are not token steps. Fix: add a title-3 and a figure role in design/tokens/tokens.json and read them.
6. NICE. The tab glyph's ink fills about 80 percent of its box while Home and Gear are drawn with more air, so Guild looks a touch larger and solid in the bar. Fix only if the owner sees it: scale the art in gen-ui-icons.mjs for `guild`.

## Verdict
VERDICT: DRIFTS
(One SHOULD, five NICE. Fix finding 1 and the gate passes on a re-check; 2 is a spec ruling.)
