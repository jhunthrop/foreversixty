# Guide accuracy round 25: Shaman

Auditor: shaman. Audit tree 438f8ba15 (main after guide round 24 and the guild r1 merge; no sim round 22 in the log), engine pin 170a59c39 (`sim/enginever/version.go`), active client 1.60.1.70790 (`data/builds/1.60.1.70790/`).

Files audited:
- `web/src/content/guides/shaman/{index,elemental,enhancement,restoration}.md`, line by line (unchanged since round 24's tree: `git diff ea9adea1c..HEAD` touches no shaman page).
- Every note in `data/curated/apl/shaman-{elemental,enhancement,restoration}.json` and its rendered rows in `addon-data.json` / `addon/ForeverSixty/Data.lua`. Only the Elemental Searing Totem note changed (72c5e8361, then b04c8f5b4).
- New scope: the coach starters `data/curated/coach/shaman-{elemental,enhancement,restoration}.md`, their rule files `shaman-*.json`, the shaman entries of `data/pipeline/coachrules.py`, and the coach code that reads them (`api/internal/coach/f1_core_casts.go`, `f2_uptime.go`, `f5_ranks.go`, `builddata.go`, `rules.go`).
- Generator: `tools/coach/persona_starter.py`, `starter_overrides.py`, `starter_kit.py`, `starter_guide.py`, `starter_priority.py`, `test_persona_starter.py`.
- Orders decoded from `addon-builds.json` (index 0 = level 10), trainer levels from `trainables/shaman.json`, consumables from `sim/request/raid-consumes.json`.

Read first: round 24 shaman report, the ledger's Round 24 section, round 25 paladin and druid reports (coach finding format, shared roots).

## Coach regeneration check

`python3 persona_starter.py --spec shaman-<spec> --out <scratch>` for all three (scratch: `scratchpad/gr25/shaman-regen/`). Elemental and Enhancement: byte-identical to the committed files. Restoration: one line differs, the Riptide line (R25-11), which the regen fixes. Every other finding below survives a plain regen.

`test_persona_starter.py`: the shaman cases pass today (generated names below). `test_guide_order_statements_win` pins `elemental[:2] == ["Mana Spring Totem", "Searing Totem"]`, which encodes the wrong order (R25-1), so the fix changes the test too.

Generated priority names: Elemental Mana Spring Totem, Searing Totem, Lightning Bolt, Chain Lightning, Magma Totem, Lightning Mastery, Earth Shock, Fire Nova, Flame Shock. Enhancement Windfury Weapon, Windfury Totem, Strength of Earth Totem, Mana Spring Totem, Totemic Call, Rage of the Farseer, Stormstrike, Lightning Bolt, Searing Totem, Fire Nova, Earth Shock, Flame Shock, Magma Totem, Lightning Shield. Restoration Water Shield, Healing Stream Totem, Riptide, Chain Heal, Healing Wave, Lesser Healing Wave, Nature's Swiftness, Poison Cleansing Totem, Cure Poison, Ancestral Spirit.

## Round 24 fix, rechecked (new work)

- **Elemental card row 3 (Searing Totem), b04c8f5b4:** "Searing Totem on a single target, and in the pack until Magma Totem at 26: until then it is your only Fire totem that does damage, and from 12 it is the one Fire Nova bursts from."
  - Fire totems before 26 (`trainables`, `spellconst` school 4): Searing 10, Frost Resistance 24. Flametongue Totem comes at 28 and Fire Nova Totem's only trainer row at 52. So Searing is the only damaging one before 26. True.
  - Fire Nova trains at 12 (408341) and bursts from "your active Fire totem". With Searing down it is that totem. True at every band from 12, and the clause is scoped "from 12" for bands 10 and 11.
  - The condition `(targets == 1 or not known Magma Totem) and Searing not up` matches the first sentence.
  - The note is rendered on all 26 Elemental bands in `addon-data.json`, and in `Data.lua` and `sim/request/apl/shaman-elemental.apl.json`.
  - OK.

## Findings

| id | status | kind | where | quote | evidence | fix |
|---|---|---|---|---|---|---|
| SHAMAN-R25-1 | CONFIRMED | coach: priority order contradicts the page and the card | `coach/shaman-elemental.md:23-31`; `starter_overrides.GUIDE_ORDER["shaman-elemental"]`; `test_persona_starter.test_guide_order_statements_win` | The rendered order: Mana Spring, Searing Totem, "Lightning Bolt on every free global" (third), Chain Lightning, Magma, Lightning Mastery, Earth Shock, Fire Nova, then Flame Shock last | elemental.md:74: Flame Shock goes back on "the moment it falls off, ahead of everything but your cooldowns". The card agrees: row 0 Mana Spring, 1 cooldowns (Lightning Mastery), 2 Flame Shock, 3 Searing, 4 Magma, 5 Chain Lightning, 6 Fire Nova, 7 and 8 Lightning Bolt. The coach's header says the list is "in the order it gives them". In the coach, Lightning Bolt "on every free global" comes before Flame Shock, the pack spells and Lightning Mastery. GUIDE_ORDER's comment ("Mana Spring Totem first, the totems, Lightning Bolt the rest") does not match the page either. The test pins the wrong second entry. | Needs an override edit: GUIDE_ORDER `shaman-elemental` = ("Mana Spring Totem", "Lightning Mastery", "Flame Shock", "Searing Totem", "Magma Totem", "Chain Lightning", "Fire Nova", "Lightning Bolt", "Earth Shock"). Test: `assert elemental[:3] == ["Mana Spring Totem", "Lightning Mastery", "Flame Shock"]`. The Lightning Bolt line also needs an override, see R25-6. |
| SHAMAN-R25-2 | CONFIRMED | coach: a stale override restates a superlative round 21 removed, and drops the pack rule | `coach/shaman-elemental.md:24`; `GUIDE_LINES["shaman-elemental"]["Searing Totem"]` | "Searing Totem against a single target, for the higher fire damage." | SHAMAN-R21-1 found that Magma Totem out-damages Searing Totem per target at every band from 26 (rank 1 is 10 a second against Searing rank 2's 6.2; rank 4 is 36.5 against rank 6's 19.3), so the card note lost "strongest". The page never gives "higher fire damage" as the reason; the override text predates the 2026-11-02 rewrite (fdf6c87f4). The line also says "single target" only. The card (row 3) and elemental.md:118 drop Searing in packs too until Magma Totem at 26. | GUIDE_LINES "Searing Totem": "Searing Totem on a single target, and on packs until Magma Totem at twenty-six. It is cheap for its damage and its bolts reach twenty yards, so drop it within twenty yards of the boss." |
| SHAMAN-R25-3 | CONFIRMED | coach: a false damage claim | `coach/shaman-elemental.md:26`; `GUIDE_LINES["shaman-elemental"]["Chain Lightning"]` | "It hits every target for close to single-target Lightning Bolt damage." | Chain Lightning rank 4 (10605): "dealing 510 Nature damage ... Each jump reduces the damage by 30%. Affects 3 total targets." Against Lightning Bolt rank 10's 353 (15208), the hits are 510, 357 and 250: about 144%, 101% and 71% of a bolt. A fourth or fifth target takes nothing. The card note (row 5) says it right: "the first hit is bigger than a Lightning Bolt and each jump still deals 70% of the hit before it". The page never says "close to". | GUIDE_LINES "Chain Lightning": "Chain Lightning from thirty-two on two or more targets. Its first hit beats a Lightning Bolt, and each jump deals seventy percent of the hit before it." |
| SHAMAN-R25-4 | CONFIRMED | coach: a stale override contradicts the page and card | `coach/shaman-elemental.md:30`; `GUIDE_LINES["shaman-elemental"]["Fire Nova"]` | "Fire Nova on a Clearcasting proc only when your Fire totem is out and the boss stands close to it. Otherwise the proc goes to your next Lightning Bolt." | elemental.md:80: "Fire Nova is a pack spell for you ... Let a Clearcasting proc go to your next Lightning Bolt instead." Card row 6 casts Fire Nova only on two or more targets; its note says "on a single boss ... Lightning Bolt is the better cast". The override quotes "Spend it on Fire Nova only if your Fire totem is out and the boss is within 10 yards of it", text the page no longer has. The coach teaches a single-target Fire Nova the card never casts. It never says Fire Nova is the pack spell on its cooldown (elemental.md:84). | GUIDE_LINES "Fire Nova": "Fire Nova from twelve on two or more targets, bursting from the Fire totem you dropped in the pack. On a single boss, a Clearcasting proc goes to your next Lightning Bolt." |
| SHAMAN-R25-5 | CONFIRMED | coach: a stale override with a reason the page contradicts in a raid, and a dropped condition | `coach/shaman-elemental.md:23`; `GUIDE_LINES["shaman-elemental"]["Mana Spring Totem"]` | "Mana Spring Totem goes down first, one global before the cooldowns. The loop runs short of mana long before the fight ends." | elemental.md:82: "In a raid, cast rank 10 the whole fight: raid buffs carry its mana close enough to the end". Card row 7: "In a raid you never need to [drop]". So "runs short of mana long before the fight ends" is false in the raid the list describes. elemental.md:78 and card row 0: "Two Mana Springs do not stack: if a Restoration shaman already has one down for your group, skip yours". The coach line says "goes down first" with no exception, so a player who rightly skips it is taught the opposite. The override cites "starting it one global before the cooldowns"; the page now says "a few seconds before the pull" (elemental.md:72, card prepull at -4 s). | GUIDE_LINES "Mana Spring Totem": "Mana Spring Totem from twenty-six, a few seconds before the pull and whenever your Water slot is empty. Skip it if a Restoration shaman already has one down for your group: two do not stack." |
| SHAMAN-R25-6 | CONFIRMED | coach: buttons named before the player has them | `coach/shaman-elemental.md:23, :26, :27, :28` | ":23 Mana Spring Totem goes down first"; ":26 Chain Lightning only against two or more targets"; ":27 Magma Totem in place of Searing Totem against two or more targets"; ":28 Then Lightning Mastery with your cooldowns" | Trainer and order levels: Mana Spring 26, Magma 26, Chain Lightning 32 (`trainables`); Lightning Mastery 40 on both orders. At 12 to 25, :27 swaps out the Searing Totem the card keeps in packs (row 3) for a totem the player does not have. elemental.md:114 levels Elemental from 1, and a dungeon boss pull from 15 is logged against these lines. Same shape as PALADIN-R25-5, -12 and -15 (D1). ":25 Lightning Bolt on every free global" also drops the rank rule that elemental.md:82 and :118 and card rows 7 and 8 teach. | GUIDE_LINES: Mana Spring Totem as R25-5; Magma Totem: "Magma Totem from twenty-six on two or more targets, in place of Searing Totem."; Chain Lightning as R25-3; "Lightning Mastery": "Lightning Mastery from forty with your trinkets and Blood Fury on the pull, and again every three minutes."; "Lightning Bolt": "Lightning Bolt with every other global. Cast your top rank while your mana is ahead of the fight, a cheaper one when it falls behind." |
| SHAMAN-R25-7 | CONFIRMED | coach: a stale override contradicts the card and the page (sim round 21) | `coach/shaman-enhancement.md:33`; `GUIDE_LINES["shaman-enhancement"]["Earth Shock"]` | "Earth Shock last, and only while mana is above half." | Sim round 21 moved the gate to a quarter. Card row 7: `currentManaPercent > 25%`, note "while your mana is above a quarter". enhancement.md:77: "while you have more than a quarter of your mana". enhancement.md:87: "Stop shocking at a quarter of your mana". Round 24 confirmed no half-reserve text survived on the page or the card; this override is the one place it still does. | GUIDE_LINES "Earth Shock": "Earth Shock last, on cooldown while your mana is above a quarter. Keep that last quarter for Stormstrike and your totems." |
| SHAMAN-R25-8 | CONFIRMED | coach: two imbue lines contradict each other, one under the wrong button | `coach/shaman-enhancement.md:23, :24`; `GUIDE_LINES["shaman-enhancement"]["Windfury Weapon"]`; generator `starter_priority.with_name` | ":23 Windfury Weapon on your one weapon before the pull, and again whenever it falls off. Before level thirty, Rockbiter Weapon goes on it instead." / ":24 Windfury Totem: Imbue your weapon: Rockbiter Weapon in a raid where a Windfury Totem reaches you, Windfury Weapon everywhere else." | enhancement.md:66, :79 and :101 say: in a raid where a Windfury Totem reaches you, Rockbiter Weapon, because Windfury Weapon turns off the totem's bonus. The level 55 raid preset carries `main_hand_imbue:rockbiter_weapon`. The override at :23 says Windfury Weapon from 30 with no raid exception, and :24 then says the opposite. :24 is the page's imbue sentence filed under "Windfury Totem". Rockbiter Weapon is not among the generator's castable names, so `with_name` prefixed the sentence's first castable one. The list never teaches Windfury Totem itself (32; page :81). | GUIDE_LINES "Windfury Weapon": "Windfury Weapon on your weapon from thirty, before the pull and whenever it falls off. Before thirty, and in a raid where a Windfury Totem reaches you, use Rockbiter Weapon." and "Windfury Totem": "Windfury Totem from thirty-two in your Air slot when no other shaman has it down near the melee. Otherwise, Grace of Air Totem." Generator: `with_name` should not prefix a sentence with an ability that is not its subject. |
| SHAMAN-R25-9 | CONFIRMED | coach: a spell the page says not to cast, listed bare as a priority button | `coach/shaman-enhancement.md:34` | "Flame Shock." | enhancement.md:83: "Outside a proc, Fire Nova on its own cooldown costs more mana than its damage is worth on one target, and so does Flame Shock." The card never casts Flame Shock. The list header calls these "the buttons the guide's rotation names, in the order it gives them", so the bare name tells an Enhancement shaman to spend globals and mana on it. The rule-out wording ("costs more mana than its damage is worth ... and so does X") is not one `starter_guide.excluded_in` reads, and `starter_overrides` has no hook to drop a name. | Needs a generator change: `starter_guide` should read "costs more ... than its damage is worth" and "and so does X" as rule-outs. Or add an exclusion hook to `starter_overrides` (an EXCLUDED map applied in `corrected`) with `shaman-enhancement: ("Flame Shock",)`. Regenerate. |
| SHAMAN-R25-10 | CONFIRMED | coach: buttons named before the player has them | `coach/shaman-enhancement.md:24, :27, :28, :29` | ":29 Stormstrike on cooldown."; ":28 Rage of the Farseer, trinkets, Blood Fury and your potion go together on the pull"; ":27 ... press Totemic Call"; ":24" Windfury Totem | Stormstrike is the 25th-level point on both Enhancement orders. Rage of the Farseer is the 40th on both. Totemic Call trains at 22 and Windfury Totem at 32 (`trainables`). enhancement.md:121 calls Enhancement the shaman's leveling spec, so boss pulls from 15 are logged. At 15 to 24 the first damage button on the list does not exist. Same shape as R25-6. | GUIDE_LINES "Stormstrike": "Stormstrike on cooldown from twenty-five. It makes your next Lightning Bolt, Chain Lightning or Earth Shock on the target hit twenty percent harder." "Rage of the Farseer": "Rage of the Farseer from forty, with your trinkets and Blood Fury, on the pull and again on cooldown." "Totemic Call": "Totemic Call from twenty-two when the boss moves away from your totems: follow it, and it brings them to you." Windfury Totem as R25-8. |
| SHAMAN-R25-11 | CONFIRMED | coach: stale committed starter contradicts the page and card (round 23 change) | `coach/shaman-restoration.md:25` | "Riptide on the tank whenever it is ready and the tank is hurt." | restoration.md:71 and card row 6 (`not dotIsActive`): recast only when your last Riptide has run out, because it is ready at 12 s but heals for 15, so "an early recast loses its last tick". | Fixed by regen. The scratch output reads "Riptide on the tank when it is hurt and your last Riptide has run out. It is ready again at twelve seconds but heals for fifteen, so an early recast loses its last tick." (Its level is R25-16.) |
| SHAMAN-R25-12 | CONFIRMED | coach: the Chain Heal line is another heal's sentence, false for Chain Heal | `coach/shaman-restoration.md:26`; generator `starter_priority` (sentence to ability), `with_name` | "Chain Heal: No single-target heal gives more healing for its mana, and it sets up the Chain Heal." | This is restoration.md:71's Riptide sentence. In it, "No single-target heal ..." is about Riptide and "it sets up the Chain Heal" means Riptide sets up Chain Heal. Chain Heal is not a single-target heal, and the line says it sets itself up. The list never gives the page's Chain Heal rule (restoration.md:72, card row 16): on the tank as soon as anyone near it is hurt, while your mana is ahead of the fight, with Riptide's 25% on the target. The sentence names only Chain Heal, so it was filed under Chain Heal. It does not start with "It", so `leans_on_the_last` let it through. It survives regen. | GUIDE_LINES `shaman-restoration` "Chain Heal": "Chain Heal from forty on the tank as soon as anyone near it is hurt, while your mana is ahead of the fight. On a target with your Riptide it heals that target twenty-five percent more." Generator: do not assign a sentence to an ability it names only as an object after "sets up", "into" or "for". |
| SHAMAN-R25-13 | CONFIRMED | coach: the main tank heal is missing, and the fast heal's line has lost its subject | `coach/shaman-restoration.md:27, :28` | Healing Wave appears only as ":27 Your highest rank of Healing Wave right after Nature's Swiftness"; ":28 Lesser Healing Wave is the fast heal when it is down to thirty percent and the next hit could kill it." | restoration.md:73 item 3 and card row 17 (rendered 10 to 60): "Healing Wave when the tank is down to 60%. Drop a rank when your mana is behind the fight, but never below rank 8." That is the spec's staple single-target heal. The coach keeps only the Nature's Swiftness use, so the list reads as if Healing Wave is the emergency button. The generator took the first Healing Wave note (row 5) and dropped item 3. Item 3's next sentence became :28, where "it" now has no antecedent and reads as your mana down to thirty percent. Same shape as PALADIN-R25-13. | GUIDE_LINES "Healing Wave": "Healing Wave on the tank when it drops to sixty percent. When your mana falls behind, drop a rank, but never below rank eight at sixty." "Lesser Healing Wave": "Lesser Healing Wave from twenty when the tank is down to thirty percent and the next hit could kill it." "Nature's Swiftness": "Nature's Swiftness from forty-four when the tank drops below twenty-five percent, into your highest rank of Healing Wave." |
| SHAMAN-R25-14 | CONFIRMED | coach: a judged mana cooldown missing from the priority | `coach/shaman-restoration.md:23-32` (no Mana Tide line) | (none) | Card row 1: "Mana Tide Totem when you drop below 25% mana ... put Healing Stream back down when it ends". restoration.md:84: "Mana Tide replaces either for 12 seconds; drop it under 25% mana". `coachrules.py` lists Mana Tide Totem first under `efficiency`, the list F1 judges for a healer, and "Who they are" names it. The rotation prose says "Mana Tide", not "Mana Tide Totem", so the name match misses it (the shared short-name root, filed once here for shaman). | GUIDE_LINES "Mana Tide Totem": "Mana Tide Totem from twenty-five when your mana drops below a quarter. It takes the Water slot, so put Healing Stream back when it ends." The test's restatement check passes, since the card casts it (`rotation_cast_names`). GUIDE_ORDER `shaman-restoration` = ("Water Shield", "Healing Stream Totem", "Mana Tide Totem", "Riptide", "Chain Heal", "Healing Wave", "Lesser Healing Wave", "Nature's Swiftness"). |
| SHAMAN-R25-15 | CONFIRMED | coach: drill contradicts the page and the card | `coach/shaman-restoration.md:46` (drill.core_casts); `starter_kit.HEALER_CORE_CASTS_DRILL` | "When the mana bar drops below half, check your mana cooldowns before the next cast. Press each one as soon as the bar has room for its full return." | Mana Tide Totem is the shaman's mana cooldown (efficiency list). Rank 4 (1325256) returns 290 every 3 s for 12 s, 1,160 in all, so "room for its full return" comes at about 85% of a raid shaman's pool. The drill would have it pressed at half mana or above. restoration.md:84 and card row 1 say under 25%. Nature's Swiftness, the other efficiency entry, is not a mana cooldown at all (see R25-17). | Needs a kit change: per-spec healer drills, or an override hook for drills. Shaman: "When your mana drops below a quarter, drop Mana Tide Totem, then put Healing Stream back when it ends. Use mana potions and runes as soon as each fits under your cap." |
| SHAMAN-R25-16 | CONFIRMED | coach: buttons and talents named before the player has them | `coach/shaman-restoration.md:23, :24, :25, :26, :27, :28, :29` | Water Shield and Healing Stream Totem "whenever"; Riptide; Chain Heal; Nature's Swiftness; Lesser Healing Wave | Water Shield is the 20th-level point; Healing Stream Totem and Lesser Healing Wave train at 20; Chain Heal at 40; Nature's Swiftness is the 44th point and Riptide the 50th on both orders. restoration.md:116 sends a Restoration shaman to dungeons before 20 ("train each one before a dungeon", "Tremor Totem at 18 breaks the sleeps in Wailing Caverns"). A Wailing Caverns healer at 18 is coached on six buttons out of seven they do not have. | Prefix the level after the name, as in R25-12 to R25-14 (the test needs each line to start with the name): "Water Shield from twenty whenever it is down ...", "Healing Stream Totem from twenty whenever your Water slot is empty ...", "Riptide from fifty on the tank when it is hurt and your last Riptide has run out ...". |
| SHAMAN-R25-17 | LIKELY | coach rules: an emergency cooldown judged as an efficiency cooldown to keep rolling | `coachrules.py` `shaman-restoration` efficiency "Nature's Swiftness"; `api/internal/coach/f1_core_casts.go` | (rule entry) | restoration.md:74: "Nature's Swiftness into your highest rank of Healing Wave when the tank is under 25%. Keep it for exactly that". Card row 4 is gated on the tank below 25%, and card row 0's note says it is "saved for the emergency". F1 skips it only at zero casts (talent). After one use it counts every 180 s it sits ready (`cooldown_ms` 180000, inside F1's 4 s to 600 s band) as lost, and fires at 2 lost and 10%. A shaman who saves it as the page says, on a long fight where the tank dips once, is flagged "Nature's Swiftness: N casts lost". LIKELY because it needs a fight of about eight minutes or more with one early emergency. | Drop "Nature's Swiftness" from `shaman-restoration` `efficiency` and rerun `python -m pipeline coach-rules`. Mana Tide Totem stays. Its own under-25% rule rarely fires F1: a 300 s cooldown needs ten held minutes. |

Counts: CONFIRMED 16, LIKELY 1, NEEDS-SIM 0, NEEDS-OWNER 0. All 17 are the new coach scope. Page prose and card notes have no finding: round 24's Searing Totem fix holds.

By fix route:
- Fixed by regen: R25-11.
- Override (`starter_overrides.py` GUIDE_LINES / GUIDE_ORDER): R25-2, -3, -4, -5, -6, -7, -10, -13, -16.
- Override plus test: R25-1 (GUIDE_ORDER and `test_guide_order_statements_win`).
- Override plus generator: R25-8 and R25-12 (`starter_priority.with_name` / sentence assignment), R25-14 (short-name root).
- Generator or new override hook: R25-9 (`starter_guide` rule-out wording, or an exclusion map).
- Kit change: R25-15 (`starter_kit.HEALER_CORE_CASTS_DRILL` per spec).
- coachrules.py: R25-17.
- No coach Go code change is needed for shaman. Enhancement's only core entry, Stormstrike, is a talent, so F1's talent skip covers it before 25.

## Shared roots (filed elsewhere; named once for shaman, not counted)

- **Template meta talk:** all three "Who they are" paragraphs ("as the site's guide describes it ... not written down here yet ... a later full write"), and the Voice blocks' healer lines in the Elemental and Enhancement files.
- **"Top rank" drill:** Elemental and Enhancement drill.ranks ("every button on the priority list should show its top rank"). Elemental casts Lightning Bolt ranks 3 to 8 while leveling and rank 8 at bare 60 (`rank-choices.json`, elemental.md:82, :118). F5 honours those rows; the drill does not.
- **"Recast with 3 seconds left":** Elemental and Enhancement drill.uptime. Flame Shock is reapplied "the moment it falls off" (elemental.md:74, card row 2 `not dotIsActive`); a recast with 3 s left clips a tick.
- **Healer drills:** Restoration drill.ranks ("a mid rank of your filler") against "never below rank 8" at 60 (restoration.md:73, card row 17). drill.resource ("lean on your cheapest heal") against "Anything learned at 23 or earlier heals for its tooltip number and nothing more" (:76).
- **Combat potion by resource, and flask by style:** Enhancement lists Flask of the Titans and Major Mana Potion. The level 55 raid preset has no flask and carries Major Frenzy Potion. The coach's ":28 ... and your potion go together on the pull" then reads as a Major Mana Potion at full mana. Elemental and Restoration match their presets.
- **"No X" clauses and short names:** R25-14 is the short-name root ("Mana Tide"). R25-9 is the inverse of the "no X" root: a rule-out the parser does not read.

## SHOULDs (wording, not findings)

- **coach/shaman-enhancement.md:17** (Who they are, from enhancement.md:30): "a free instant Lightning Bolt whenever Maelstrom Weapon reaches five stacks" is true only at 5/5 (42 leveling, 39 raid). :30 gets it right.
- **coach/shaman-enhancement.md:32** Fire Nova on Clearcasting: Clearcasting comes from Elemental Focus in the raid build only (52). Card row 6 says so and the coach line does not.
- **coach/shaman-enhancement.md:36** Lightning Shield: drops the page's "Keep it for leveling and for adds that turn on you" (enhancement.md:89).
- **coach/shaman-enhancement.md:27** "Totemic Call: If the boss moves ... press Totemic Call": the name is doubled and capitalised after the colon (`with_name` prefix). R25-10's line replaces it.
- **coach/shaman-enhancement.md** never names the Air totem (card row 2, loop item 1).
- **coach/shaman-elemental.md:31** "nothing else wants the shock cooldown" sits two lines under ":29 Earth Shock stays on your bar for interrupts". It is true of the card (no other shock row), and the same words are on card row 2. Carried as the card's Should.
- **coach/shaman-restoration.md:27 and :29** repeat the Nature's Swiftness rule; R25-13 merges them. The "Everyone else" rules (restoration.md:78) and Chain Heal ranks (:80) are absent.
- **Kit:** Restoration lists no weapon enchant, though restoration.md:100 and the preset carry Brilliant Mana Oil (as PRIEST-R25-S5). Elemental's consumables omit the Demonic Rune that elemental.md:98 and the preset carry.
- **coach/shaman-elemental.md:34 and coach/shaman-restoration.md:35** "Core abilities under-used": Elemental's `core` is empty, so F1 can only say it cannot check. The line describes a check the coach does not run for this spec.
- Carried from round 24, still true as SHOULDs: the Elemental filler first sentence at band 60; restoration.md:120 "far more"; enhancement.md:74 "costs you nothing"; enhancement.md:87 (optional sim); Enhancement Searing "would barely fire"; elemental.md:90 "Ancestral Knowledge"; Stormstrike "hardest hit" near a tie at 40 (also coach :29); Grace of Air "rank 3 at 60"; elemental.md:74 "about twice"; elemental.md:66 against :114; Healing Wave note "17 levels past"; enhancement.md:55; index.md:28; Elemental Flame Shock "with Call of Flame" at 10 to 22 (also coach :31); Chain Lightning "same cast time" outside Lightning Mastery.

## Checked and right this round

- **Page prose, full sweep.** Unchanged since round 24; every superlative, "only", "every", "never" and quoted cost re-read against the unchanged rows and the decoded orders. These hold:
  - Lightning Bolt 198 / 171 with Convection.
  - Totemic Call at 25% of base mana and 30 s (437009) on enhancement.md:42.
  - Stormstrike 8% base mana and 6 s (17364).
  - Riptide 445 / 575 / 805 over 15 s, 12 s cooldown.
  - Mana Spring 10 mana every 2 s at rank 4.
  - Lightning Mastery 15 s, 3 minutes (16166).
  - Rage of the Farseer 30% for 25 s, 3 minutes (425336).
  - Mana Tide 12 s, every 5 minutes.
  - Chain Heal 247 / about 300 / about 385.
  - Elemental Focus at 26 on both Elemental orders, so the talent is in the build at the levels elemental.md:60 describes.
- **Coach lines that hold:**
  - Elemental :29 (Earth Shock for interrupts, trains at 4) and :31 (the card's Flame Shock note).
  - Enhancement :25 (Strength of Earth, 30 yards), :26 (Mana Spring in the Water slot with no other shaman), :30 (Lightning Bolt at five stacks "once you have all five points", the card's gate `tag 5`), :31 (the 20 s Searing guard), :35 (packs).
  - Restoration :23 (Water Shield 2% a globe), :24 (Healing Stream every 2 s for 5 minutes), :27 (Nature's Swiftness into the top rank), :30 to :32 (cleansing totems 22 and 38, Cure Poison 16, Cure Disease 22, Ancestral Spirit 12, all conditional).
  - Consumables: Elemental (Supreme Power, Baked Salmon, Brilliant Wizard Oil, Major Mana Potion) and Restoration (Distilled Wisdom, Sage's Tea, Major Mana Potion) match their top presets.
- **Coach rule files:**
  - Elemental tracks Flame Shock as a DoT (all six ranks) and has no core entry, per the comment.
  - Enhancement's core is Stormstrike (talent flag, so it is skipped before it is learned).
  - Restoration tracks Water Shield as a self buff (talent; F2 says "can't check" when it is never up).
  - Restoration's efficiency is Mana Tide Totem (300 s from `category_cooldown_ms`, read by `builddata.go:47`) and Nature's Swiftness (R25-17).
  - F5: Elemental Lightning Bolt and Restoration Healing Wave / Chain Heal rank rows are honoured through `rank-choices.json` (`sitesOwnRank`) or the healer overheal branch, so the page's downranking is not flagged.

## Decisions

- **D1. Coach lines are judged at every level a pull can be logged** (brief; PALADIN-R25 D1, druid). A button named with no level is a finding when a common logged band lacks it (R25-6, R25-10, R25-16). One finding per file, every button listed.
- **D2. R25-3 is CONFIRMED**, not a Should. "every target" fails at four or more targets (three in total), and the third hit is 71% of a bolt, which the card's own note states.
- **D3. R25-15 is filed for shaman**, though the drill text is shared kit. No shared root covers it, and for this spec it contradicts a page and card threshold (25%) with a different one (half, and room for the full return).
- **D4. R25-17 is LIKELY**: it needs a long fight with one early emergency.
- **D5. Proposed override lines start with the ability name, carry no digits and keep each sentence to 20 words or fewer**, so `test_every_correction_restates_its_guide` and `test_guide_lines_are_short` stay green.
- **D6. Not re-filed:** round 24's D3 list (rows rendering before their talent, card action ranks, BiS race captions, leveling tree rendering, the below-20 healing penalty); the site card's per-race rank issue (sim round 22).

## Sims requested

- None new. Round 24's optional run stands: Enhancement Earth Shock 25% gate against no gate on the doubled Maelstrom chance (enhancement.md:87).

## Engine questions

- None.

## Sim findings

- None new. Carried from sim round 21 "Still open": Maelstrom Weapon flat 100% against 20 a minute on a fast weapon; Lightning Overload rank spells; Healing Wave 2 tooltip; level 60 book ranks.

## Sources

No web research was needed: the client data, the card rows, the presets and the coach code decided every call.

- Client 1.60.1.70790:
  - `spelltips/shaman.json`: 10605, 421, 15208, 1325256 and Mana Tide ranks 16190 / 17354 / 17359, 16188, 437009, 17364, 16166, 425336, 408521 / 1239242 / 1239243, 5675 / 10497.
  - `spellconst/shaman.json`: school and `category_cooldown_ms` of 3599, 8181, 8190, 16188, the Mana Tide ranks and 408510.
  - `trainables/shaman.json`; `addon-builds.json` (six shaman orders decoded); `addon-data.json`; `addon/ForeverSixty/Data.lua`.
- Curated: `data/curated/apl/shaman-*.json`, `data/curated/rank-choices.json`, `data/curated/coach/shaman-*.{md,json}`, `data/pipeline/coachrules.py`, `sim/request/raid-consumes.json`.
- Coach: `api/internal/coach/f1_core_casts.go`, `f2_uptime.go`, `f5_ranks.go`, `builddata.go`, `rules.go`; `tools/coach/*` (scratch regen at `scratchpad/gr25/shaman-regen/`).
- Reviews: round 24 and round 21 shaman reports (SHAMAN-R21-1 Magma against Searing), ledger Round 24, round 25 paladin and druid reports.
- Git: 72c5e8361 and b04c8f5b4 (round 24 shaman); fdf6c87f4 (2026-11-02 rewrite that the Elemental overrides predate).

## Appendix A: Note sweep

The cards are unchanged since round 24 apart from Elemental row 3. Each note was re-read against its condition, both orders (decoded above) and the trainer. Bands are as in round 24.

**Elemental** (orders identical apart from Ancestral Knowledge / Totemic Focus; Elemental Focus 26, Call of Flame 23 to 25, Trial by Fire 29 to 30, Lightning Mastery 40)
- [0] Mana Spring, 26-60: T 26; skip-if-stacked clause; many times its cost. OK.
- [1] Lightning Mastery / cooldowns: 40; 15 s; Berserking after. OK.
- [2] Flame Shock, 10-60: Call of Flame conditional (carried Should). OK.
- [3] Searing Totem, 10-60: only damaging Fire totem before 26 (Frost Resistance 24 does no damage, Flametongue 28); Fire Nova from 12; 20 yards; Magma from 26. OK (round 24 fix holds).
- [4] Magma Totem, 26-60: 20 s, 40 from 30. OK.
- [5] Chain Lightning, 32-60: first hit above the top bolt; 70% a jump. OK.
- [6] Fire Nova, 12-60: Searing until 26, Magma after; 10 yd; one-boss clause. OK.
- [7] Lightning Bolt top rank, 10-60: 198 / 171 at 60; about 20%. OK.
- [8] Lightning Bolt filler, 10-60: rank list matches `rank-choices.json` (3 at 26 and 27, 3 at 30 to 33, 4 to 38, 5 to 44, 6 to 50, 7 to 56, 8 from 57, 8 bare at 60). OK (carried first-sentence Should).

**Enhancement** (Stormstrike 25 both; Maelstrom 37 to 42 leveling, 35 to 39 raid; Rage of the Farseer 40 both; Elemental Focus raid 52)
- [0] autocast: no note.
- [1] Strength of Earth, 10-60: 30 yards. OK.
- [2] Grace of Air, 42-60: T 42, rank 2 at 56; Windfury Totem 32. OK.
- [3] Stormstrike, 40-60: 6 s; 20% (17364). OK (carried "hardest hit").
- [4] Lightning Bolt at five stacks, 10-60: `tag 5` = 42 leveling, 39 raid; instant and free only at 5/5. OK.
- [5] Searing Totem, 10-60: 20 s guard. OK.
- [6] Fire Nova on Clearcasting, 12-60: "in the raid build" (52). OK.
- [7] Earth Shock, 10-60: 25% gate; Stormstrike "once you have it". OK.

**Restoration** (both orders identical: Water Shield 20, Mana Tide 25, Nature's Swiftness 44, Riptide 50)
- [0] autocast: potions, runes; Nature's Swiftness saved. OK.
- [1] Mana Tide, 10-60: under 25%; 12 s; rank 4 at 58. OK.
- [2] Water Shield, 20-60: three globes, 2%. OK.
- [3] Healing Stream, 20-60: every 2 s for 5 minutes. OK.
- [4] Nature's Swiftness, 10-60: tank below 25%. OK.
- [5] Healing Wave after Nature's Swiftness: top rank. OK.
- [6] Riptide, 40-60: 12 / 15 s; recast when run out; 25%. OK.
- [7] Lesser Healing Wave tank, 20-60: 1.5 s. OK.
- [8]-[15] party rows: no note.
- [16] Chain Heal, 40-60: ranks 2 (46) and 3 (54); 247 / about 300 / about 385; rank 2 bare at 60. OK.
- [17] Healing Wave tank, 10-60: rank 8 floor at 60; rank list matches `rank-choices.json`. OK (carried "17 levels past").

## Appendix B: Coach sweep

One line per coach line: file:line, claim, the level the player has it, verdict.

**shaman-elemental.md**
- :17 Who they are: overview copy. Template root.
- :23 Mana Spring (26), "goes down first", "runs short of mana". R25-5, R25-6.
- :24 Searing Totem (10), "for the higher fire damage". R25-2.
- :25 Lightning Bolt (1), "every free global". R25-1 (order), R25-6 (rank rule).
- :26 Chain Lightning (32), "every target ... close to". R25-3, R25-6.
- :27 Magma Totem (26), in place of Searing on two or more. R25-6.
- :28 Lightning Mastery (40) with cooldowns. R25-6; order R25-1.
- :29 Earth Shock (4) for interrupts. OK.
- :30 Fire Nova (12) on Clearcasting at a boss. R25-4.
- :31 Flame Shock (10), Call of Flame (23). Order R25-1; carried Should.
- :34 to :41 finding list: "Core abilities under-used" can't be checked (SHOULD); ranks line is the top-rank root.
- :43 consumables: match the preset. OK (SHOULD: Demonic Rune).
- :47 core_casts OK. :48 uptime: 3-second root. :49 idle OK. :50 resource OK. :51 ranks: top-rank root. :52 preparation OK. :53 death OK. :54 adherence OK.
- :58 to :67 Voice: healer lines in a damage file (template root).

**shaman-enhancement.md**
- :17 Who they are: five-stack bolt "free instant" (SHOULD); template root.
- :23 Windfury Weapon (30), no raid exception. R25-8.
- :24 "Windfury Totem: Imbue your weapon ..." R25-8; Windfury Totem (32) R25-10.
- :25 Strength of Earth (10). OK.
- :26 Mana Spring (26), party condition. OK.
- :27 Totemic Call (22). R25-10; SHOULD (doubled name).
- :28 Rage of the Farseer (40), Blood Fury, potion. R25-10; potion root.
- :29 Stormstrike (25). R25-10; carried "hardest hit".
- :30 Lightning Bolt at 5/5 (42 leveling, 39 raid). OK.
- :31 Searing Totem (10), unless the boss is about to die. OK.
- :32 Fire Nova on Clearcasting (raid 52). SHOULD.
- :33 Earth Shock above half. R25-7.
- :34 Flame Shock, bare. R25-9.
- :35 Magma Totem (26) and Fire Nova (12) placement. OK.
- :36 Lightning Shield. SHOULD.
- :39 to :46 finding list: ranks line is the top-rank root.
- :48 consumables: Titans and Major Mana Potion against a preset with no flask and Major Frenzy Potion (shared root).
- :52 core_casts OK. :53 uptime: 3-second root. :54 idle OK. :55 resource OK. :56 ranks: top-rank root. :57 preparation: flask root. :58 death OK. :59 adherence OK.
- :63 to :72 Voice: template root.
- Rule file: core Stormstrike (talent, guarded). OK.

**shaman-restoration.md**
- :17 Who they are: Riptide (50), Chain Heal (40), Water Shield (20), Mana Tide (25). Template root.
- :23 Water Shield (20). R25-16.
- :24 Healing Stream Totem (20). R25-16.
- :25 Riptide "whenever it is ready" (50). R25-11 (regen), R25-16.
- :26 "Chain Heal: No single-target heal ..." (40). R25-12, R25-16.
- :27 Healing Wave after Nature's Swiftness (44). R25-13 (only Healing Wave line), R25-16.
- :28 Lesser Healing Wave (20), orphan "it". R25-13, R25-16.
- :29 Nature's Swiftness (44). R25-16; SHOULD (repeat).
- :30 to :32 cleansing and Ancestral Spirit. OK.
- (missing) Mana Tide Totem. R25-14.
- (missing) Healing Wave at 60%. R25-13.
- :35 to :40 finding list: generic. OK.
- :42 consumables: match the preset (SHOULD: oil).
- :46 core_casts drill. R25-15. :47 uptime OK. :48 resource: cheapest-heal root. :49 ranks: mid-rank root. :50 preparation OK. :51 death OK.
- :55 to :64 Voice: OK for a healer (template root for the header).
- Rule file: Water Shield aura OK; efficiency Mana Tide OK, Nature's Swiftness R25-17.
