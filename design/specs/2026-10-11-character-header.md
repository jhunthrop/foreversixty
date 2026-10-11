# The shared character header: experience spec

Author: ux-designer. 2026-10-11. Owner instruction: "the character header needs to look the same(ish) on every page." One component, one fixed anatomy, two slots. Mock: `design/mocks/gen_character_header.py` (also imported by `gen_rankings.py` and `gen_logs.py`, so those boards use this code, not a copy). Render: `design/mocks/renders/character-header-states.png` (every state at 1440 and 390, the controls, what each page fills, the two card layouts).

Merges: the BiS landing lead (`2026-10-10-bis-landing.md` section 2.4, as built in `gen_bis_landing.py`) and the home signed-in hero (`2026-10-10-home-signed-in-panel.md` section 3.B). Existing code it replaces: `web/src/components/CharacterCard.svelte`, `web/src/components/planner/PlannerCharacterCard.svelte`, the identity block in `HomeAccountPanel.svelte`/`character/HomeHeroCards.svelte`, and the `CurrentCharacterBar spine` on Rankings and Logs. References: `docs/tenets.md` 9, 10, 11, 14; the nav selector spec (the header selector is the only way to change character; ember means stale).

## 1. What it is, in one paragraph

A full-width panel at the top of a page's content, signed in only: **who** (crest, name, descriptor, sync age), **the page's answer about that character** (one figure and one line), and **the page's one action** (one gold button). It follows the header selector in place and has no switcher of its own. Every page that mounts it looks the same except for the words in the middle and on the button.

## 2. Fixed, and what each page fills

| Region | Fixed by the component | Filled by the page |
|---|---|---|
| Panel | Raised ground, 1 px gold border, 6 px radius, padding 18 x 22, min height **136** at 1024 and up (round 2: the tallest state, a two-line stale sync under a two-line-capable identity, so no state changes the band's height). Three columns `1.15fr / 1fr / auto`, gap 28, inside the 1344 column. Below 1024, and in any container under 900 px, it stacks (identity, answer, full-width button), padding 16, gap 16. | Nothing. No tint, no size variants. |
| Identity | Crest 64 px (56 stacked), circular ringed HD crest, the one crest language. Label above the name: **always `Your character`** (round 2: it must not change meaning when an alt is added). Name: Cinzel 700, **24 px (22 stacked)**, class colour, full name, clamps at two lines with the whole name in `title`; the crest and name are one link. Descriptor: `{Race} {Spec} {Class} · Level {n} · {flat faction logo 16 px} {Faction}`, 13 px, e.g. `Gnome Fury Warrior · Level 47 · Alliance`; the race is left out when the API has not sent it, never a placeholder. In a narrow container (below 1280 or a stacked card) the descriptor is two fixed lines, `{Race} {Spec} {Class}` then `Level {n} · {logo} {Faction}`, so a separator can never start or end a line; wide, it is one line. Sync line, 12 px muted: `{Source} · {age}` (Addon, Battle.net or Pasted export), capped at two lines. On a phone the name box reserves two lines so a header switch moves nothing. | The link target of the identity link (default: the page's own action target). Nothing else. |
| Answer (slot) | Grammar: label (11 px upper, muted), then **either** a figure line (JetBrains Mono 26 px, unit 15 px muted; **a primary figure and an optional second figure at the same size**, e.g. home's `9 upgrades` and `+11.2 DPS`; optional 6 px bar only when the figure is a fraction of a whole), **or** one 14 px sentence; then at most two 12 px muted note lines (the second only where a page has two facts, Logs); the first may carry one gold inline link whose arrow never wraps. Max four lines. | The label, figure, unit, note, optional bar, and the strings for the states that apply to it (section 3). The figure colour may be a parse colour only where the page's board uses that ladder (Rankings). |
| Action (slot) | One button, 48 px, filled gold gradient, min width 250, label `{Verb} {object} →`, always the only filled gold element on the page. States: rest, hover, focus-visible (2 px gold outline, 2 px offset), active, loading. Never a mutating action: it opens a page, jumps to an anchor, or opens the file picker. | The label and the target. |
| Emphasis | `primary` (landing pages) is the filled button. `secondary` (tool-page context cards: planner, BiS spec page) is a 44 px outline gold button. This is the only variant, chosen by whether the header is the page's lead. | Which one. |

The panel never carries the Updated stamp, a source pill, a switcher, a "not you?" link or a second button. The hint `Change character: top right` (home) is **dropped** here (the selector is at the same place on every bar).

## 3. States (fixed by the component, words by the page)

| State | What shows | Who supplies words |
|---|---|---|
| Ready | As section 2. | Page: answer, action. |
| Loading, nothing known (`/v1/me` pending) | The ready box rendered hidden, with skeletons drawn over it, so the height is exactly the ready height at every width (390 included). | None. |
| Loading the answer only | Identity and the action are real (both come from `/v1/me`; labels are static); the answer slot is a skeleton of its own size. This is also exactly what a header-selector switch does: **identity swaps in place from held data, no skeleton; the answer re-fetches under a skeleton; no reload, no announcement of its own** (the selector's live region says "Now using {Name}."). | None. |
| Stale sync | Only the words past the age go ember: `Addon · 19 days ago · Log out in the game to sync` (nav spec 2.11 threshold). Name and ring are never recoloured. A page whose answer reads gear changes its label (BiS: `Slots you wore at the last sync`); a page that does not read gear (Rankings, Logs) is unchanged. | Page: the stale label if it reads gear. |
| No gear on file | The sync line goes. A page that reads gear (home, BiS, BiS spec card) replaces the answer with one 13 px sentence and a 44 px gold text link (`Add the addon to see how many of these slots you already wear.` / `Get the addon →`); the action still works. A page that does not (Rankings, Logs) changes nothing else. No `0 of 15`: an unknown is not a zero. | Page: sentence and link, if it reads gear. |
| Under 20 | Identity unchanged. A page whose answer starts at a level (home, BiS: 20) shows its sentence (`Best in slot starts at 20.`) in the slot and points its action at the first band. A page with no level floor (Rankings, Logs: logs exist at any level) is identical to ready. | Page: sentence and the first-band action label, if it has a floor. |
| Empty answer | One 14 px sentence in body colour; the action stays and is the way out. | Page. |
| Answer failed | Identity and action stay; the slot is one 13 px sentence and a 44 px `Try again`. Nothing says "failed" on its own. | Page: the sentence. |
| Signed out, no character, or `/v1/me` failed | **One shared signed-out band on every page that mounts the header:** a single 44 px muted 13 px line between hairlines, `Sign in or paste an export to see your character here`. No button, no link: the nav selector is the only sign-in control and **pages add no sign-in buttons or prose of their own**. | None. |
| Long name | Clamps at two lines; the panel grows once, and only for a name that cannot fit 24 characters at 24 px. | None. |

Controls (rendered): action rest, hover, focus-visible, active, secondary rest/hover/focus; identity link hover (name underlines, 4 px offset) and focus-visible. Hover never moves anything. Touch: every target is 44 px or taller (button 48, text links 44, identity link at least 56).

## 4. Pages

All figures are examples; real ones come from the sources named.

| Page | Layout | Answer (label / figure / note) | Action | Source |
|---|---|---|---|---|
| **Home** (signed in) | band, primary | `Upgrades in your band` / `9 upgrades` and `+11.2 DPS` at the same size / `Best in slot for 60, together` | `See your upgrades →` (`#upgrades`) | `HomeUpgradesPanel` data, unchanged |
| **BiS landing** | band, primary | `Best-in-slot slots you wear` / `6 of 17 slots` with bar / (a note only where slots are excluded) | `Fury gear for 60 →` | bis-landing spec 2.4, unchanged |
| **Rankings** | band, primary | `Your rank` / `14th of 212 Fury Warrior kills` (parse-ladder colour) / `On Onyxia, Normal ruleset` | `See your parses →` (`/character/{region}/{ruleset}/{name}`) | rankings spec |
| **Logs** | band, primary | `Latest kill` / `612 DPS` / `Onyxia kill · 3 days ago · Open the fight →`, then `Last report 1 day ago` | `Upload a log` (opens the file picker; on a phone `Open your latest report →`) | logs spec |
| **BiS spec page** | card (440 px), secondary | the BiS count for this spec and band, same strings | `Send this list to the addon` | replaces `CharacterCard.svelte` |
| **Planner** | card (440 px), secondary, no answer | none (the planner's job is the tree, not a figure about the character) | `Send this build to the addon` | replaces `PlannerCharacterCard.svelte` |

**Who converges and who does not, with the reason.**

- **Home converges, and is the page that loses the most.** Its 34 px name and 84 px crest become 24 px and 64 px; the DPS gain is no longer a footnote but a second figure at figure size beside the count; the three next-action cards give up the Best in slot card (now the answer and the button) and become a two-card row, Talents and Simulator, directly under the panel. Why: the owner's instruction, and the home spec's own point that the name should be what the page is about, which the gold-bordered panel still says. This is an owner call (open question 1).
- **BiS landing, Rankings, Logs: converge fully** (landing layout, primary action).
- **BiS spec page and planner cards converge on identity and every state** (they already were this component's two earlier copies; their identical crest/name/descriptor markup is the DRY case), in the stacked card layout with a secondary action. They are context on a tool page whose own controls are the point (tenet 11), so their button is outline. The old 40 px crest, 15 px name and `Switch` aside go. The planner card keeps no answer slot.
- **Do not mount (kept as they are):** the Simulator (its `LandingState` list and spine are a separate audit, home spec open question 3; a candidate for the next pass), Tier list (not about a character; it keeps its `Your spec` row treatment), Guides, `/character/...` and `/account` (they are the character's own page, with the Blizzard render), guild pages, `/setup`.
- The `CurrentCharacterBar spine` is removed from Rankings and Logs (this component carries identity; the door links and `Forget` live in the selector and nav).

## 5. Implementation notes for the lane

One Svelte component `CharacterHeader.svelte` in `web/src/components/character/` with slots `answer` and the props `action {label, href, kind: 'link'|'picker'}`, `emphasis`, `layout` (auto by container width), `needsGear`, `levelFloor`. It reads `createCharacterCardState` (existing) and `CURRENT_CHARACTER_CHANGED`. The inline crest and faction recipe in the two old cards becomes the one copy here. Reserve 112 px (signed-in hint on `html[data-session='1']`, as BiS reserves 108 today) so nothing shifts; signed out reserves nothing. Run the full vitest and Playwright suites plus the lhci shard (shared layout rule). Tests: identity string for every state, no pill, one filled gold element per page, follows the selector without reload, omitted when signed out.

## 6. Open questions for the owner

1. Home's 34 px name becomes 24 px (section 4). The alternative (a bigger name only on home) is a size variant and breaks "the same".
2. One filled gold primary per page (BiS spec open question 2 already asked); the planner and BiS spec cards use outline.
3. Name size 24 px everywhere moves BiS's 22 px up by 2.

## 7. Round 2 (after the WoW player's round-one review, 2026-10-11)

Applied as decided: (1) second figure at figure size; (2) distinct labels (`Upgrades in your band`, `Best-in-slot slots you wear`, `Your rank`, `Latest kill`); (3) stale copy `Log out in the game to sync`, sync line capped at two lines, band height 136 so no state changes it; (4) the 390 loading skeleton is the ready box hidden under skeletons, exact height by construction; (5) one shared signed-out band, no page sign-in controls; (6) race in the descriptor when known; (7) the label is always `Your character`; (8) the planner and BiS spec card variant is unchanged (the new copy flows into it).

Chosen not to do, with the reason:
- **Home's own signed-out hero keeps its sign-in button and class picker.** It is the front door, outside this component, and the home spec makes that button the home's one filled gold. Decision 5 says pages add no sign-in buttons; including home needs its own hero redesign (open question 4).
- **BiS landing signed out gains the 44 px band**, pushing its first row of cards 44 px lower (the bis-landing spec made that row above the fold at 1440, binding). The BiS lane must re-measure; if the row falls below 900 the band drops its hairlines and shrinks. Not re-rendered here (the BiS generator is out of scope).
- **The 16 px faction logo stays 16 px**: the descriptor also carries the faction word, which is the readable part. Open to 18 px.
- The example is now a level 60 Gnome Fury Warrior; the BiS `6 of 17 slots` and home `9 upgrades, +11.2 DPS` are invented and appear only on the states sheet.

4. Open question: should home's signed-out hero join the shared band?

## 8. Round 3

The descriptor rule above replaces the wrapping segments (a lone `·` appeared at 390 and 1024). Rendered in `character-header-states.png`, `rankings-phone.png`, `logs-1024.png`. The Logs answer's second note, `Last report 1 day ago`, is a link to that report (underlined muted text, not gold, so the one gold link stays `Open the fight →`).
