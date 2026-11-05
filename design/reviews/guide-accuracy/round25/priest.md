# Guide accuracy round 25: Priest

Audited on the round 25 tree (head 438f8ba15: main after guide round 24; no sim round 22 in the log). Engine pin 170a59c39 (`sim/enginever/version.go`). Active client 1.60.1.70790 under `data/builds/1.60.1.70790/`.

Files audited, line by line:

- `web/src/content/guides/priest/{index,discipline,holy,shadow}.md`.
- Every note and row condition in `data/curated/apl/priest-{discipline,holy,shadow}.json`, and every rendered row of `addon-data.json` `rotations/priest-*` with `race_spell_ids` (Holy 23 bands, Discipline 25, Shadow 22).
- New this round: the coach files `data/curated/coach/priest-{discipline,holy,shadow}.md`, with the generator `tools/coach/persona_starter.py`, `starter_overrides.py`, `starter_kit.py`, `starter_priority.py`, `starter_guide.py`, `starter_food.py`, and `sim/request/raid-consumes.json`.
- `addon-builds.json` (all six priest orders, decoded against `talents/priest.json`), `trainables/priest.json`, `spelltips/priest.json`, `spellconst/priest.json`.
- Coach scope: `docs/superpowers/specs/2026-10-27-log-coach-design.md` §6.1 and F5.

Read first: round 24 priest, the ledger's Round 24 section, round 25 hunter (for the shared coach roots).

## Changes since round 24's tree (ea9adea1c)

- One priest commit, 43957d431: the Holy Heal note (`priest-holy.json` priorityList[30]) and holy.md:122, round 24's three fixes.
- `addon-data.json` priest rotations are byte-identical to ea9adea1c (diffed row by row). The addon shows a note's first sentence, and the Heal note's first sentence did not change.
- Cards, `rank-choices.json`, `addon-builds.json` and the priest BiS files are unchanged.

## Coach regeneration check

`persona_starter.py --spec priest-discipline` and `--spec priest-shadow` were run into the scratchpad (`gr25/priest-coach/`) and diffed against the committed files: **no difference for either**. A regen fixes none of the findings below. `priest-holy` is a full persona (`FULL_SPECS`), not generated, so its fixes are hand edits to `data/curated/coach/priest-holy.md`. The generator test has four red tests (`test_persona_starter.py`); none fails on a priest spec.

## Round 24 fixes, re-checked

- **Holy Heal note** (priorityList[30]), every band 16 to 60, against the rows (`race_spell_ids`, Heal 2054/2055/6063/6064):
  - 28: rank 3, Human and Undead rank 2.
  - 29: rank 3, Undead rank 2.
  - 30 to 37 (bands 30, 32, 33): rank 3.
  - 38: rank 4, Troll and Undead rank 3.
  - 39: rank 4, Undead rank 3.
  - 40 on: rank 4.
  - "keep rank 3 from 34 to 37, to 38 as a Troll and to 39 as an Undead" and "A Human or Undead also keeps rank 2 at 28, and an Undead at 29" match every row. Rank 4 trains at 34, so band 33 (33 to 37) casting rank 3 needs no override.
  - The mana reason now sits on rank 4 only. It holds per cast: at 38, rank 3 (543, 255 mana) beats rank 4 (669, 305 mana) above 116 healing, and Holy wears 233 to 244 at 30 to 39.
  - The first sentence, the one the addon shows, is unchanged. OK.
- **holy.md:122**: the same ranks by race in the same words, and "Renew steps down too" (R24-S2's wording). OK.
- Round 24's carried Shoulds stand as listed there.

## Findings

All sixteen are in the coach files. No guide page or card note is false this round.

| id | status | kind | where | quote | evidence | fix |
|---|---|---|---|---|---|---|
| PRIEST-R25-1 | CONFIRMED | coach: a spell no Forever priest has; a racial given to every Holy priest | `coach/priest-holy.md:26, :28, :44` | ":26 Shadowfiend returns a share of the whole mana bar over its fifteen seconds, so it waits until the bar is below half"; ":28 Shadowfiend and Dark Sacrifice are in the Holy priest's kit"; ":44 When the mana bar drops below half, check Shadowfiend before your next cast." | index.md:28: "Shadowfiend, the Season of Discovery rune, is not in Forever." The APL header (round 3) says the same: no learn level on its skill-line row, and no beta priest cast it. Dark Sacrifice is the Undead priest spell (index.md:45; card note "Undead only"). "In the Holy priest's kit" gives it to every race. Round 24 named this for the coach lane. | Hand edit (`priest-holy.md` is a full persona). :26: drop the Shadowfiend sentence, and say "Dark Sacrifice, if you are Undead, trades your own health for mana ...". :28: "Spiritual Guidance turns Spirit into healing, and an Undead priest has Dark Sacrifice." :44: "When the mana bar drops below half, check your Major Mana Potion and Demonic Rune, and Dark Sacrifice if you are Undead, before your next cast." |
| PRIEST-R25-2 | CONFIRMED | coach: the filler reversed | `coach/priest-holy.md:18, :22, :47` | ":18 the filler is a mid rank of Heal"; ":22 Heal at a mid rank is the filler: once Greater Heal is learned, a mid rank of Heal shares its healing coefficient for far less mana, the classic downrank."; ":47 Bind a mid rank of Heal and use it for every small hole on the tank." | holy.md:70: "The filler is Greater Heal at your top rank. When it would overheal or your mana runs short, step down to rank 1 ... leave ranks 2 and 3 off your bars. Heal rank 4 is cheaper still, but it keeps only about half your bonus, so it is the last step down." The card agrees: rows [10] to [14] are the Greater Heal filler, [20] to [24] Greater Heal rank 1, and Heal rank 4 ([30]) only "when you are short of mana for Greater Heal rank 1, since rank 4 ... keeps only about half your bonus at 60". Heal rank 4 has MaxLevel 39, so at 60 it keeps 55% of the bonus and does not "share" Greater Heal's healing in any useful sense. From 40 to 60 the coach teaches the step the page puts last as the first one. | Hand edit. :18 "the filler is Greater Heal, stepped down to rank 1 for small holes"; :22 "Greater Heal is the filler at your top rank. Step down to rank 1 when it would overheal or your mana runs short; Heal rank 4 is the last step down, since it keeps only about half your healing bonus at sixty."; :47 "Bind Greater Heal rank 1 and use it for every small hole. Save the top rank for a hole as deep as it heals." |
| PRIEST-R25-3 | CONFIRMED | coach: false client mechanic | `coach/priest-holy.md:22` | "Prayer of Healing heals your own party" | holy.md:31 and :76: "In Forever it heals the target and the target's party ... you are not limited to your own party." The card note says "it heals the target and the target's whole party". Round 24 named this. | Hand edit: "Prayer of Healing heals the target's party, so target someone in the group that took the hit, and it earns its cost only when three of the four are hurt." |
| PRIEST-R25-4 | CONFIRMED | coach: contradicts the card and itself | `coach/priest-holy.md:28` | "its first rank arrives at forty and heals little for its price, so at that level it waits for the mana pace like the filler." | The Prayer of Mending row ([2]) has no condition at any band (rank 1 at 40 to 45, rank 2 from 50, rank 3 at 60), and its note says "It is not held back by the mana pace the costlier heals follow". The coach's own :24 says "Prayer of Mending ... never paced." | Hand edit: "Prayer of Mending is a Forever spell and the capstone of the tree. From forty it goes on the tank every time it is ready, whatever your mana." |
| PRIEST-R25-5 | CONFIRMED | coach: rule narrower than the card | `coach/priest-holy.md:22` | "Flash Heal is the fast heal and the dear one, only for a tank who is genuinely low." | The card also casts Flash Heal on any party member under 30% (rows [5] to [8]; note: "when someone is about to die, the speed matters more than the mana"). holy.md:74 says the same: "on a tank who is really low or anyone about to die". Read as written, the coach counts a Flash Heal that saves a dying party member as a mistake. | Hand edit: "Flash Heal is the fast heal and the dear one, for a tank who is really low or anyone about to die." |
| PRIEST-R25-6 | CONFIRMED | coach: rank reason false at 60 | `coach/priest-holy.md:22` | "Renew stays on the tank all fight, at a cheap rank, because a cheaper rank returns more health per mana." | holy.md:68: "Renew always up at rank 7 or higher." holy.md:72: Renew ranks 1 to 4 "heal for little more than their base" at 60, and ranks 5 and 6 "keep only part of it". At 60, with Holy's 554 raid healing:<br>- Rank 1 (45 over 15 s, 30 mana, no bonus) returns about 1.5 health a point of mana.<br>- Rank 7 (370, 250 mana, full bonus) returns about 3.7.<br>So "a cheaper rank returns more health per mana" fails on every rank below 7. The card keeps rank 7 at 60. | Hand edit: "Renew stays on the tank all fight at rank seven or higher: it still keeps your healing bonus and costs far less than the top rank." |
| PRIEST-R25-7 | CONFIRMED | coach: list order against the page and card; core button missing | `coach/priest-discipline.md:21 to 33` | "These are the buttons the guide's rotation names, in the order it gives them." List: Lightwell, Fear Ward, Renew, Greater Heal, Lesser Heal, Heal, Flash Heal, Prayer of Healing, Inner Focus, Penance, Dark Sacrifice. | discipline.md:66 and :68 put "Penance on cooldown ... the heal you never let sit" and "The shield is a switch" first after the pull. The card runs Dark Sacrifice, Penance, Power Word: Shield, Renew, Flash Heal, Prayer of Healing, Greater Heal, Heal, Greater Heal rank 1, Lesser Heal. The coach lists Penance tenth, behind every filler, and leaves out Power Word: Shield, the spec's core button. The overview at :17 names the shield, but the list does not. Cause: the rotation prose says "Shield the tank", never "Power Word: Shield", so `starter_priority` never sees the name. | `starter_overrides.py`:<br>- GUIDE_LINES `priest-discipline`: add "Power Word: Shield": "Power Word: Shield on the tank at the top rank once he is hurt, Weakened Soul is off him and your mana keeps pace with the fight." It is on the card, so `test_every_correction_restates_its_guide` passes.<br>- GUIDE_ORDER `priest-discipline`: ("Lightwell", "Fear Ward", "Penance", "Power Word: Shield", "Renew", "Flash Heal", "Prayer of Healing", "Greater Heal", "Heal", "Lesser Heal", "Inner Focus", "Dark Sacrifice").<br>Regenerating alone does not fix it. |
| PRIEST-R25-8 | CONFIRMED | coach: list order against the page and card | `coach/priest-shadow.md:21 to 33` | "in the order it gives them": Shadow Word: Pain, Inner Focus, Mind Blast, Shadow Word: Death, Mind Flay, Shoot, Shadowform, Devouring Plague, Vampiric Embrace, Fade, Dark Sacrifice. | shadow.md:81: "Keep Shadow Word: Pain up. Keep Devouring Plague up ... Mind Blast on cooldown. Shadow Word: Death on cooldown. Mind Flay with everything else." The card runs Shadow Word: Pain, Vampiric Embrace, Devouring Plague, Inner Focus with Mind Blast, Mind Blast, Shadow Word: Death, Mind Flay, Shoot. The coach puts Devouring Plague and Vampiric Embrace below Mind Flay and the wand, and Shadowform, the before-the-pull step (:77), in the middle. The "core abilities under-used" finding reads list order ("a button the priority list wants sitting ready while something lower went first"), so a Devouring Plague cast while Mind Blast was ready, as the card does it, reads as a mistake. The GUIDE_ORDER comment quotes :81 without its Devouring Plague sentence. | `starter_overrides.py` GUIDE_ORDER `priest-shadow`: ("Shadowform", "Shadow Word: Pain", "Vampiric Embrace", "Devouring Plague", "Inner Focus", "Mind Blast", "Shadow Word: Death", "Mind Flay", "Shoot"), with the full :81 quote in the comment. Regenerating does not fix it. |
| PRIEST-R25-9 | CONFIRMED | coach: contradicts the page (shared root, HUNTER-R25-2) | `coach/priest-shadow.md:41, :45, :54` | "a combat potion (Major Mana Potion)" | shadow.md:111: "Open with Forever's Major Spellblasting Potion; a Major Mana Potion and a Demonic Rune come out on a long fight." Shadow's top raid band in `sim/request/raid-consumes.json` (from 55) carries `major_spellblasting_potion` and no mana potion. Discipline's carries `major_mana_potion`, so its line is right. | Generator, the shared root: `starter_kit.kit()` picks the potion from `Traits.resource`. Read the combat potion from the spec's top-band preset, the way `starter_food.preset_food` reads the food. Regenerating alone does not fix it. |
| PRIEST-R25-10 | CONFIRMED | coach: drill contradicts the card (shared root, HUNTER-R25-8) | `coach/priest-shadow.md:50` (drill.uptime) | "Recast each with 3 seconds left" | Devouring Plague's cooldown equals its duration (24 s at every rank), and Vampiric Embrace's does too (30 s), so neither can be recast with 3 seconds left. The Shadow Word: Pain, Vampiric Embrace and Devouring Plague rows all cast on `not dotIsActive`, that is once the effect is gone. Shadow Word: Pain ticks every 3 s, so a recast at 3 s left throws away its last tick. | Generator, the shared root: `starter_kit.DRILLS["dps"]["uptime"]`: "... Recast each as the priority list says, no earlier, and again the moment you are back up after a death or a phase change." Regenerating does not fix it. |
| PRIEST-R25-11 | CONFIRMED | coach: drill contradicts the rank advice (shared root, HUNTER-R25-7) | `coach/priest-shadow.md:53` (drill.ranks) | "Visit your trainer every time you level and read your spellbook after: every button on the priority list should show its top rank." | shadow.md:137 keeps cheaper ranks while leveling. The card rows agree:<br>- Shadow Word: Pain rank 3 at 30 to 38.<br>- Devouring Plague rank 1 at 30 to 43 and rank 2 at 44 to 59.<br>- Mind Blast rank 4 at 40 to 43 for every race but Gnome.<br>- Shadow Word: Death rank 2 at 50 to 58.<br>The coach reviews any logged boss pull at the player's level (design §6.1 "Level ... F5's rank check"), and F5 itself skips the site's downranks. Only the drill tells the player to undo them. At 60 every Shadow row casts the top rank, so the drill is true there. | Generator, the shared root: `starter_kit.DRILLS["dps"]["ranks"]` (and `["tank"]`): "Visit your trainer every time you level, then check every button on the priority list shows the rank the guide gives for your level: sometimes that is a lower one." |
| PRIEST-R25-12 | CONFIRMED | coach: healer drill contradicts the filler (shared root) | `coach/priest-discipline.md:50` (drill.ranks); the same text in Holy :47 is R25-2 | "Bind a mid rank of your filler heal and use it for every small hole on the tank. Save the top rank for a hole as deep as it heals." | discipline.md:72: "Greater Heal at the top rank on the tank, where the big hits land. Greater Heal rank 1 on everyone else." The card agrees: tank filler Greater Heal at the top rank, party Greater Heal rank 1. :74 ("Downrank within the ranks you learned from your forties on") leaves no mid rank. The drill moves the tank off the top rank and onto a rank the page does not use. | Generator: `starter_kit.DRILLS["healer"]["ranks"]` is shared by every healer starter, so a per-spec drill is needed. For Discipline: "Bind Greater Heal rank 1 for the party's small holes, and keep the top rank for the tank, who takes the big hits." Regenerating does not fix it. |
| PRIEST-R25-13 | CONFIRMED | coach: healer drill against the page's behind-pace rule (shared root) | `coach/priest-discipline.md:49` (drill.resource) | "If you are behind it, lean on your cheapest heal and save the fast, dear heal for the real emergency." | discipline.md:84: "When your bar falls behind the fight, stop topping off the party and heal only real holes." :74 drops the cheapest casts at 60: "keep Lesser Heal off your bars for good ... wait the few seconds for your regeneration rather than cast a heal that no longer grows with your gear." The card's behind-pace rule is the same: the filler waits for a tank under 50% and a member under 35%. At 60 the cheapest heals by cost are Lesser Heal and the low Renew and Heal ranks, the ones the page takes off the bar. | Generator: `starter_kit.HEALER_RESOURCE_DRILL`, shared by every healer starter, needs a per-spec text. For Discipline: "If you are behind it, keep Penance and the shield on the tank, let the filler wait for a deeper hole, and save Flash Heal for the real emergency." |
| PRIEST-R25-14 | CONFIRMED | coach: talent named before the order gives it | `coach/priest-shadow.md:26, :27, :29` | ":26 Shadow Word: Death on cooldown ... Keep Vampiric Embrace up so your own damage heals back its backlash."; ":27 never clip it: the third tick is the biggest."; ":29 then shift into Shadowform" | The card notes carry the conditions ("once you have Vampiric Embrace", "once you have Improved Mind Flay"); the overrides drop them.<br>- Shadow Word: Death trains at 32, Vampiric Embrace is the leveling order's point at 33: at 32 the line names a talent the player does not have.<br>- Mind Flay comes at 28 (leveling) and 21 (raid), Improved Mind Flay at 29 and 22: at 28 and 21 the three ticks are equal.<br>- Shadowform is 42 on both orders, so the line is false on a boss pull at 40 to 41. | `starter_overrides.py` GUIDE_LINES `priest-shadow`:<br>- "Shadow Word: Death": "... Once you have Vampiric Embrace, keep it up so your own damage heals back its backlash."<br>- "Mind Flay": "Mind Flay with everything else, and never clip it: a clipped channel loses the rest, and once you have Improved Mind Flay the third tick is the biggest."<br>- "Shadowform": "From forty-two, buff the raid, then shift into Shadowform and drink back to full ..." |
| PRIEST-R25-15 | CONFIRMED | coach: spells named before the player has them | `coach/priest-discipline.md:31` | "Inner Focus on a Greater Heal or a Prayer of Healing. It makes the spell free and adds twenty-five percent crit chance, and a Greater Heal crit leaves the biggest Divine Aegis you can make." | Inner Focus is the order's point at 24. Prayer of Healing trains at 30, Divine Aegis is 36, Greater Heal is 40. From 24 to 29 the line holds Inner Focus for two spells the player does not have, so it is never pressed. The card autocasts it on the next priest spell. The Divine Aegis clause is false to 35, and the Greater Heal clause to 39. | `starter_overrides.py` GUIDE_LINES `priest-discipline` "Inner Focus": "Inner Focus on your most expensive heal: from forty a Greater Heal, where a crit leaves the biggest Divine Aegis you can make, or a Prayer of Healing. It makes the spell free and adds twenty-five percent crit chance." |
| PRIEST-R25-16 | CONFIRMED | coach: site, file and data meta talk (voice rule; shared root, HUNTER-R25-13) | `coach/priest-discipline.md:17, :21, :43`; `coach/priest-shadow.md:17, :21, :45`; `coach/priest-holy.md:22, :38` | "This is the Discipline Priest as the site's guide describes it ... This file says what the guide says and leaves the rest for a later full write."; "These are the buttons the guide's rotation names ... Where the guide or the rotation data gives a reason"; "For this spec the catalogue looks for ..."; Holy :22 "This is the order the site's Holy rotation runs."; Holy :38 "the catalogue looks for" | Owner rule (guide voice), which the round 25 brief applies to the coach files: no site or sim meta talk. | Generator, the shared root: `persona_starter.who_section` and the `know_section` lead-in, and `starter_kit.consumable_paragraph`. Speak as the player ("Your buttons, in order:"; "Before a boss you want a flask, food and ..."). Holy is a hand edit: ":22 The order: the tank first."; ":38 For a healer: Flask of Distilled Wisdom ...". |

Counts: CONFIRMED 16, LIKELY 0, NEEDS-SIM 0, NEEDS-OWNER 0.

- Guide pages and card notes: 0.
- Coach, Holy (hand edits): 6 (R25-1 to 6), plus Holy's part of R25-16.
- Coach, Discipline and Shadow: 9. Regenerating fixes none.
  - Four need `starter_overrides.py`: R25-7, R25-8, R25-14 and R25-15.
  - Five need a generator or kit change, all shared roots: R25-9 `starter_kit.kit`, R25-10 and R25-11 `DRILLS["dps"]`, R25-12 `DRILLS["healer"]["ranks"]` and R25-13 `HEALER_RESOURCE_DRILL` (per-spec healer drills), and R25-16 `persona_starter.who_section`/`know_section` plus `consumable_paragraph`.

## Shoulds (wording, not findings)

- **PRIEST-R25-S1, Holy coach :24, "at least as full as the share of the fight that is left".** The card's pace gate is 90% of that share (rows [9] to [14], [20] to [24], [30] to [34]). Optional: "while the mana bar is close to the share of the fight that is left".
- **PRIEST-R25-S2, Holy coach front matter** lists `data/builds/1.60.1.70630/spellconst/priest.json`. The active build is 1.60.1.70790. This is not player-facing.
- **PRIEST-R25-S3, Shadow coach :28, "Shoot with your wand on any global where nothing else is ready."** The GUIDE_LINES comment quotes a shadow.md sentence ("Wand shots fill any global ...") that the page no longer has. The card's note reads "only when nothing else can be cast ... it never fits between spells." Optional: "Shoot only when nothing else can be cast, such as out of mana."
- **PRIEST-R25-S4, Shadow coach Voice (`tools/coach/voice.md`, shared).** A DPS persona carries "For a healer ..." and "Never tell a healer to heal more". These are instructions to the coach, and they are conditional, but they read oddly on Shadow.
- **PRIEST-R25-S5, Discipline coach :40, :43, :51 (pull preparation).** The list omits the weapon oil that the page (:96 Brilliant Mana Oil) and the preset (`main_hand_imbue:brilliant_mana_oil`) name. `starter_kit.kit` gives an oil to casters only. This is an omission, not a false statement.

Carried, unchanged and optional: R24-S1, R24-S3, R24-S4, R23-S1 to S12 (R23-S4 Discipline Renew by race stays a Should), R22-S1, R22-S2, R21-S1 to S4, R20-S1, R20-S2, R18-S1 to S5, R17-S2.

## Checked and right this round

**Guide pages, the full sweep:**
- Every line of the four pages was read again against the rows, the orders and the client. They are unchanged since round 24 apart from holy.md:122, which is re-checked above. Each comparison, cost, cooldown and rank claim that rounds 20 to 24 verified still holds on identical data:
  - Shadow Word: Death: 18 s cooldown, 30% backlash.
  - Mind Blast: 8 s cooldown.
  - Fade: 10 s, 30 s cooldown.
  - Lightwell: 3 minutes or 5 clicks.
  - Penance: 12 s cooldown, a pulse every 1 s.
  - Inner Focus: 3 minutes, 25%.
  - Shadowform: 40% of base mana, which supports "the form costs mana".
- No page or rendered note names 70755.

**Coach lines, OK at every level they apply to.** The Coach sweep appendix has one line each.
- Discipline: Lightwell, Fear Ward, Renew, Greater Heal at 60, Lesser Heal before 40, Heal to 40, Flash Heal, Prayer of Healing, Penance, Dark Sacrifice.
- Shadow: Shadow Word: Pain, Inner Focus, Mind Blast, Devouring Plague, Vampiric Embrace, Fade, Dark Sacrifice.
- Holy: the other lines.

## Decisions

- **D1. Coach scope by level.** The coach reviews any logged boss pull, kill or wipe, of 30 s or more. It reads the player's level from the log for the rank check, and F5 skips the site's band rank rows (`api/internal/coach/f5_ranks.go`, design §6.1). So a coach line must be true at every level a boss pull can be logged, dungeon bosses included. That is why R25-11, R25-14 and R25-15 are findings, although each line is true at 60.
- **D2. The shared roots are filed once each for priest, as the brief asks:**
  - R25-9 (potion, HUNTER-R25-2).
  - R25-10 (uptime drill, HUNTER-R25-8).
  - R25-11 (top-rank drill, HUNTER-R25-7).
  - R25-16 (meta talk, HUNTER-R25-13).
  - R25-12 and R25-13 are new healer-drill roots. Hunter has no healer, so they are not in its report.
- **D3. R25-13 is CONFIRMED, not a Should.** "Cheapest heal" could be read as cheapest per mana (Penance). But the page's behind-pace advice is the opposite move: let the filler wait and heal only real holes. At 60 the heals cheapest by cost are the ones the page takes off the bar. Misleading at the level the coach mostly reviews.
- **D4. R25-6 is CONFIRMED.** The coach does not name rank 7, and its stated reason ("a cheaper rank returns more health per mana") is a general rule. holy.md:72 says that rule fails at 60 for every Renew rank below 7.
- **D5. The Prayer of Mending "cheapest healing per mana" clause (Holy coach :22)** restates the page line routed in round 23 (sim 1). It is not re-filed.

## Sims requested

None new. Carried and already routed: Holy and Discipline Heal rank 1 against rank 2 at 22 to 27; Holy Heal rank 2 against rank 3 at 28 and 29 (overheal and mana); Discipline tank Lesser Heal at 10 to 15 (optional); Prayer of Mending heals per cast; Shadow Word: Pain at 6 against 10 seconds; the Discipline shield with no health gate; Holy Greater Heal rank 1 behind the pace.

## Engine questions

None new. P2 (heals scaling with the target's level) stays routed.

## Sim findings

None new. Still open, not re-filed: P2, P3, P4, S2.

## Sources

- **Client data 1.60.1.70790:** `spelltips/priest.json`, `spellconst/priest.json` (Shadowform `cost_pct` 40), `trainables/priest.json`, `talents/priest.json`, `addon-builds.json`, `addon-data.json` (rows and `race_spell_ids`, diffed against ea9adea1c).
- **Curated:** `data/curated/apl/priest-{discipline,holy,shadow}.json`, `data/curated/rank-choices.json`, `data/curated/coach/priest-*.md`, `data/curated/presets.json`.
- **Generator:** `tools/coach/persona_starter.py`, `starter_overrides.py` (GUIDE_LINES and GUIDE_ORDER for `priest-discipline` and `priest-shadow`), `starter_kit.py` (`kit`, DRILLS, HEALER_RESOURCE_DRILL), `starter_priority.py`, `starter_guide.py`, `starter_food.py`, `test_persona_starter.py` (four red, none priest). `sim/request/raid-consumes.json`: Shadow top band `major_spellblasting_potion`; Discipline and Holy `major_mana_potion`. Read-only regen output: `scratchpad/gr25/priest-coach/`.
- **Coach design:** `docs/superpowers/specs/2026-10-27-log-coach-design.md` §6.1 and F5; `api/internal/coach/f5_ranks.go`.
- **Reviews:** round 24 priest, ledger Round 24, round 25 hunter.
- **No web research needed:** the client data, the cards and the pages decided every call.

## Appendix: Note sweep

The card rows in `addon-data.json` are byte-identical to round 24's tree, and the curated notes differ only in the Holy Heal note. Every line below was re-read against the current rows. Bands are as rendered. Levels are when the player has each named item, on the shipped order or from the trainer. The healer raid orders equal their leveling orders.

### Holy (bands 10, 16, 20, 22, 26, 28, 29, 30, 32, 33, 38, 39, 40, 44, 45, 50, 51, 52, 53, 56, 58, 59, 60)

- Renew (10 to 60): "at 60 that means rank 7 or higher". OK.
- Lesser Heal rank 3, tank (band 10): "filler until you learn Heal at 16" (16); rank 3 (10); the tank before the party; the 0.9 pace. OK.
- Lesser Heal, other rows (band 10): no note.
- Heal, tank (16 to 60): round 24's text, re-checked above at every band and race:
  - "filler before ... 40": OK.
  - "Rank 4 ... first levels, keep rank 3 from 34 to 37, to 38 as a Troll and to 39 as an Undead": OK.
  - "A Human or Undead also keeps rank 2 at 28, and an Undead at 29": OK.
  - "From 40 ... short of mana for Greater Heal rank 1": OK.
  - "cheaper" (305 against 370): OK.
  - "about half your bonus at 60" (55%): OK.
- Heal, deeper hole (16 to 60): "a tank under 50% or a party member under 35%". OK.
- Dark Sacrifice (20 to 60): Undead only, the 70% gates, 15 s, Spirit, 10 minutes. OK.
- Flash Heal, tank (20 to 60): top rank, "heals less for its mana than your slower heals" at every band. OK.
- Flash Heal, party (20 to 60): under 30%. OK.
- Prayer of Healing (30 to 60): three of four under 85%, the target's party. OK.
- Prayer of Mending (40 to 60):
  - "every time it is ready" (no gate): OK.
  - "up to five" and 10 s: OK.
  - "next hit or direct heal": OK.
  - "not held back by the mana pace": OK.
  - "cheapest ... all fight": routed.
- Greater Heal, filler (40 to 60): 0.9 pace; "rank 1 until 57, then rank 3 or your top trainer rank". OK.
- Greater Heal, deeper hole (40 to 60): OK.
- Greater Heal rank 1 (40 to 60): "once you have higher ranks" (46); "costs far less"; "most of your healing bonus at 60" (85%). OK.
- Greater Heal rank 1, deeper hole (40 to 60): OK.

### Discipline (bands 10, 12, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 45, 48, 50, 54, 56, 58, 59, 60)

- Power Word: Shield (10 to 60): the under-90% gate, Weakened Soul off, the 0.9 pace, "costs a warrior no rage", "once you have Renewed Hope" (32) and Divine Aegis (36). OK.
- Renew (10 to 60): "at 60 that means rank 7 or higher". OK.
- Lesser Heal rank 3, party (10 to 38): "before 40", "late forties ... lost your healing bonus". OK.
- Heal, tank (16 to 60): "until you learn Greater Heal at 40"; "your biggest heal until then", Undead rank 2 at 30 to 33 included (605 against Penance rank 1's 547). OK.
- Dark Sacrifice (20 to 60): as Holy. OK.
- Flash Heal, tank (20 to 60): "the fast heal, and the dear one". OK.
- Flash Heal, party (20 to 60): the 0.7 gate. OK.
- Penance, tank (30 to 60): 12 s; three pulses 1 s apart; "on whoever is hurt"; "cheap for all it heals". OK.
- Prayer of Healing (30 to 60): "Rank 4 is the top rank a trainer teaches". OK.
- Greater Heal, tank (40 to 60): "at the top rank ... at 60". OK.
- Greater Heal rank 1, party (40 to 60): "missing a third"; "once you have higher ranks" (46); "far less mana". OK.
- Greater Heal, deeper hole, tank, and Greater Heal rank 1, deeper hole (40 to 60): OK.

### Shadow (bands 10, 16, 18, 20, 22, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 51, 52, 58, 60; leveling and raid orders)

- Shadow Word: Pain (10 to 60): "Always up; stop refreshing it in the closing seconds" (the 10 s gate). OK.
- Mind Blast with Inner Focus, and the Inner Focus row (10 to 60): "whenever it is up" (57 on both orders); free; 25%. OK.
- Mind Blast (10 to 60): "ahead of Mind Flay once you have it" (28 and 21); "half cost in Shadowform" (42). OK.
- Dark Sacrifice (20 to 60): below a fifth and behind the fight; Undead only. OK.
- Devouring Plague (20 to 60): 24 s at every rank; every priest; ticks crit; heals you. OK.
- Mind Flay (20 to 60): "a clipped channel loses the rest"; "once you have Improved Mind Flay" (29 and 22). OK.
- Vampiric Embrace (30 to 60): the talent at 33 and 28; 30 s; 20%; "until you have the talent". OK.
- Shadow Word: Death (32 to 60): 30%; "once you have Vampiric Embrace". OK.
- Shoot (curated note, no rendered row): OK.

## Appendix: Coach sweep

One line per coach line, at every level a boss pull can be logged (D1).

### priest-discipline.md

- :17 meta talk: R25-16. The overview sentence matches discipline.md:26. OK.
- :21 lead-in: R25-16; the order: R25-7.
- :23 Lightwell (40): OK.
- :24 Fear Ward (20): OK.
- :25 Renew rank 7 or higher at 60: OK.
- :26 Greater Heal top rank, tank filler, at 60: OK.
- :27 Lesser Heal before 40: OK.
- :28 Heal until 40, the biggest heal: OK.
- :29 Flash Heal on a tank really low: OK.
- :30 Prayer of Healing, three or more hurt: OK.
- :31 Inner Focus: R25-15.
- :32 Penance, 1 s, 12 s: OK.
- :33 Dark Sacrifice, Undead: OK.
- (missing) Power Word: Shield: R25-7.
- :36 to :41 mistakes: OK (S5 for the oil).
- :43 consumables: names OK (Distilled Wisdom, Sage's Tea, Major Mana Potion match the preset); wording R25-16.
- :47 drill.core_casts: OK.
- :48 drill.uptime (self buff): OK.
- :49 drill.resource: R25-13.
- :50 drill.ranks: R25-12.
- :51 drill.preparation: OK (S5).
- :52 drill.death: OK.
- :56 to :65 voice: OK.

### priest-shadow.md

- :17 meta talk: R25-16. The overview matches shadow.md:30. OK.
- :21 lead-in: R25-16; the order: R25-8.
- :23 Shadow Word: Pain: OK.
- :24 Inner Focus on Mind Blast "every time it is up" (57): OK.
- :25 Mind Blast ahead of Mind Flay: OK.
- :26 Shadow Word: Death and Vampiric Embrace: R25-14 (32 on the leveling order).
- :27 Mind Flay third tick: R25-14 (28 and 21).
- :28 Shoot: S3.
- :29 Shadowform: R25-14 (40 and 41). The form costs mana (40% of base). OK otherwise.
- :30 Devouring Plague 24 s, every priest, crits, heals: OK.
- :31 Vampiric Embrace "whenever it is down": OK.
- :32 Fade 10 s on a 30 s cooldown: OK.
- :33 Dark Sacrifice: OK.
- :36 to :43 mistakes: OK; "Ranks" describes what F5 flags, which skips the site's downranks.
- :45 consumables: potion R25-9; wording R25-16; Supreme Power, Baked Salmon and Brilliant Wizard Oil match the preset.
- :49 drill.core_casts: OK.
- :50 drill.uptime: R25-10.
- :51 drill.idle: OK.
- :52 drill.resource: OK.
- :53 drill.ranks: R25-11.
- :54 drill.preparation: potion R25-9.
- :55 drill.death: OK.
- :56 drill.adherence: OK.
- :60 to :69 voice: S4.

### priest-holy.md (hand-written)

- :18 mid-rank Heal filler: R25-2. The rest (Prayer of Mending and Renew on the tank, Flash Heal for the really low, mana cooldowns when there is room): OK.
- :22:
  - "the site's Holy rotation": R25-16.
  - Prayer of Mending every time it is ready, on the most-hit target: OK. "Cheapest per mana" is routed (D5).
  - Renew "cheap rank" reason: R25-6.
  - Flash Heal "only for a tank": R25-5.
  - Prayer of Healing "your own party": R25-3; "three of the four": OK.
  - Mid-rank Heal: R25-2.
- :24 pacing: S1. "Prayer of Mending, Renew on the tank and the emergency Flash Heal are never paced": OK.
- :26:
  - Inner Focus, 3 minutes, on Prayer of Healing or Greater Heal: OK.
  - Shadowfiend: R25-1.
  - Dark Sacrifice "waits for room on the bar and for you to be healthy" (the 70% gates): OK, apart from R25-1's race clause.
  - Potion and Rune with room: OK.
  - Lightwell before the pull: OK.
- :28:
  - Prayer of Mending pace at 40: R25-4.
  - Litany of Light: OK.
  - Spiritual Guidance: OK.
  - Shadowfiend and Dark Sacrifice in the kit: R25-1.
  - Spell damage on healing items: OK.
- :31 to :36 mistakes: OK.
- :38 consumables: names OK, including Forever's Greater Mageblood, Cleric's and Greater Cleric's, Elixir of Sages and Sage's Tea (`presets.json`); "the catalogue looks for": R25-16.
- :40 ten against forty: OK.
- :44 drill.core_casts: R25-1 (Shadowfiend). The Inner Focus pairing with Prayer of Healing is OK.
- :45 drill.uptime, Inner Fire: OK.
- :46 drill.resource: OK.
- :47 drill.ranks: R25-2.
- :48 drill.preparation: OK.
- :49 drill.death: OK.
- :53 to :62 voice: OK.
- Front matter: S2.
