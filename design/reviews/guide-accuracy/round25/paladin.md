# Guide accuracy round 25: Paladin

Audited on the round 25 audit tree (438f8ba15: main after guide round 24 and the guild r1 merge; no sim round 22 in the log). Engine pin 170a59c39 (`sim/enginever/version.go`). Active client 1.60.1.70790 under `data/builds/1.60.1.70790/`.

Files audited:
- `web/src/content/guides/paladin/{index,holy,protection,retribution}.md`, line by line.
- Every card note in `data/curated/apl/paladin-{holy,protection,retribution}.json` (unchanged since round 24's tree: `git diff ea9adea1c..HEAD` is empty for all three), re-read against row conditions, both point orders and `rank-choices.json`.
- New scope: the coach starters `data/curated/coach/paladin-{holy,protection,retribution}.md`, their rule files `paladin-*.json`, the paladin entries of `data/pipeline/coachrules.py`, and the coach code that reads them (`api/internal/coach/f1_core_casts.go`, `f2_uptime.go`, `f5_ranks.go`, `f7_death.go`, `prompt_rules.txt`).
- Generator: `tools/coach/persona_starter.py`, `starter_overrides.py`, `starter_kit.py`, `starter_guide.py`, `starter_priority.py`, `starter_food.py`, `test_persona_starter.py`.

Read first: round 24 paladin report, the ledger's Round 24 section, round 25 druid and warlock reports (for the coach finding format and to avoid duplicates).

## Changes since round 24's tree (ea9adea1c)

- 28e3b4125 (round 24 paladin): retribution.md:70 (Seal of Command by Unsealed Wrath), holy.md:76 ("the only heal you can keep casting"), retribution.md:97 and :135 ("almost" / "nearly every time").
- No paladin card, rank row, build order, coach file or coach rule changed.

## Coach regeneration check

`python3 persona_starter.py --spec paladin-<spec> --out <scratch>/paladin-<spec>.md` for all three (scratch: `scratchpad/gr25/paladin-regen/`): each output is byte-identical to the committed file. Nothing below is fixed by a plain regen. `test_persona_starter.py` (read-only run): 4 failures; the paladin one is `test_guide_order_statements_win` (`prot.index("Holy Shield")` raises: Holy Shield is not in the Protection list, R25-1).

## Round 24 fixes, rechecked

- **retribution.md:70.** "The leveling build never casts it: without Unsealed Wrath, Seal of Righteousness does more on its own." Card pre-pull row 1 and rows 2, 7, 8 all require `auraIsKnown(Unsealed Wrath)`; the leveling order (decoded) has no Unsealed Wrath. "With Unsealed Wrath, Seal of Command becomes your main Seal for the twist below" matches :93 and the pace gate. Holds on the raid order (Unsealed Wrath at 40). OK.
- **holy.md:76.** Item 2 names rank 8 (60). Lay on Hands (20 minute cooldown) and Light's Vigil (one charge per Vigil) are what "you can keep casting" excludes. Matches Holy row 3. OK.
- **retribution.md:97, :135.** "almost / nearly every time" matches sim round 21 C1 (19 casts against about 22 cooldowns). OK.

## Findings

| id | status | kind | where | quote | evidence | fix |
|---|---|---|---|---|---|---|
| PALADIN-R25-1 | CONFIRMED | coach: core ability missing from the priority | `coach/paladin-protection.md:23-35` | (no Holy Shield line) | protection.md:93 item 2 "Holy Shield on cooldown ... the first claim on your mana", card row 6 (no condition), pre-pull Holy Shield at -2 s, and `coachrules.py` lists Holy Shield first in `core` and as the tank's tracked aura. The generator drops it: `starter_guide.excluded_abilities` reads the Lay on Hands sentence at protection.md:107 ("until it comes back you have no Holy Shield and little threat") as a rule-out. This is the shared "no X" clause root (filed for the leveling-only case); here the clause is not even a leveling line. It is also why `test_guide_order_statements_win` fails. | Needs an override until the root is fixed: GUIDE_LINES `paladin-protection` "Holy Shield": "Holy Shield on cooldown from thirty. Its cooldown is its duration, so it is up whenever its charges last, and it is the first claim on your mana." Generator: `excluded_in` should not treat "you have no X" inside a consequence clause as excluding X. |
| PALADIN-R25-2 | CONFIRMED | coach: priority order contradicts page and card | `coach/paladin-protection.md:25-32`; `starter_overrides.GUIDE_ORDER["paladin-protection"]` | Order rendered: Righteous Fury, Seal of Fury, Hammer of the Righteous, Holy Strike, Judgement, Consecration, Greater Blessing of Kings, Devotion Aura, Swift Judgement, Exorcism | The page's list (protection.md:92-99) and the card (rows 5 to 12) are Seal of Fury, Holy Shield, Holy Strike, Hammer of the Righteous, Swift Judgement, Judgement, Exorcism, Consecration. GUIDE_ORDER (commented "list items 1-6") predates the current list: it puts Hammer of the Righteous ahead of Holy Strike and leaves Swift Judgement and Exorcism behind Consecration, which the page says to cast "only above half mana" and to cut first. The file's own header says the list is "in the order it gives them". | Needs an override edit: GUIDE_ORDER `paladin-protection` = ("Righteous Fury", "Seal of Fury", "Holy Shield", "Holy Strike", "Hammer of the Righteous", "Swift Judgement", "Judgement", "Exorcism", "Consecration"). |
| PALADIN-R25-3 | CONFIRMED | coach: override drops the page's condition and adds a superlative the page does not make | `coach/paladin-protection.md:25`; `starter_overrides.GUIDE_LINES["paladin-protection"]["Hammer of the Righteous"]` | "Hammer of the Righteous and Holy Strike on their cooldowns. Hammer of the Righteous is the largest single threat gain in this list." | protection.md:95 and card row 8: "on cooldown, once Holy Shield's mana is safe"; :101 "If you are going dry, cut Consecration first, then Hammer of the Righteous". The coach line drops the reserve. The page never calls it the largest threat gain, and per cast it is not: Hammer of the Righteous deals 3 x main-hand DPS (407632), about 180 Holy with a 60 DPS tank one-hander, while Holy Shield rank 4 (1239570) deals 221 Holy per block over four charges with 20% extra threat, and Shieldsworn fires that 221 again on every Holy Strike. The override's comment cites "list items 4 and 6", an older numbering. | Needs an override edit: "Hammer of the Righteous on cooldown from forty, once Holy Shield's mana is safe. It hits your target and up to three enemies next to it, so it is where your threat on a pack comes from." |
| PALADIN-R25-4 | CONFIRMED | coach: level 60 value at every level | `coach/paladin-protection.md:26` | "Holy Strike on cooldown: every six seconds with Shieldsworn, twenty mana, and your damage reduction from Iron Creed." | Holy Strike costs 5/9/12/14/16/17/19/20 at ranks 1 to 8, trained 6/12/20/28/36/44/52/60 (`spelltips`, `trainables`). Twenty is rank 8 at 60 only; card row 7 says "twenty at 60". The coach reads the level from the log and applies this line at every level. The line is copied from protection.md:94 ("20 mana"), which sits in the level 60 rotation (SHOULD below). | Hand edit protection.md:94 to "20 mana at 60", then regen (the generator carries the words). Or GUIDE_LINES "Holy Strike": "Holy Strike on cooldown, twenty mana at sixty: every six seconds with Shieldsworn from forty, and your damage reduction from Iron Creed." |
| PALADIN-R25-5 | CONFIRMED | coach: buttons named before the player has them | `coach/paladin-protection.md:25, :29, :31` | ":25 Hammer of the Righteous and Holy Strike on their cooldowns."; ":29 Get Greater Blessing of Kings on yourself..."; ":31 Then Swift Judgement, Holy Strike and Hammer of the Righteous in the first few globals." | Hammer of the Righteous trains at 40, Greater Blessing of Kings at 60 (Blessing of Kings at 20), Swift Judgement is the 29th-level point on the Protection order, Holy Shield (R25-1) the 30th (`trainables/paladin.json`, `addon-builds.json`). protection.md:139 tells Protection players they tank dungeons from 20. A Shadowfang Keep or Scarlet Monastery boss pull at 20 to 39 is coached against all three lines. | GUIDE_LINES `paladin-protection`: Hammer of the Righteous as in R25-3 ("from forty"); "Greater Blessing of Kings": "Blessing of Kings on yourself from twenty, Greater Blessing of Kings at sixty, unless a hunter's Heart of the Lion already covers you..."; "Swift Judgement": "Then Swift Judgement from twenty-nine, Holy Strike, and Hammer of the Righteous from forty, in the first few globals." |
| PALADIN-R25-6 | CONFIRMED | coach: drill contradicts the page and card thresholds | `coach/paladin-protection.md:57` (drill.death) | "Put your emergency cooldown on a key you can hit without looking, and press it on a dummy pull the moment your health falls past half." | protection.md:107 "Templar's Bulwark first, at about a third of your health ... All three put Forbearance on you, so plan one per bad moment"; card rows 1 to 3 fire below 35%, 25% and 20%. A drill that trains pressing the emergency cooldown at half spends Forbearance on a hit the card would not answer, and locks out the next two. (A Dwarf's Stoneform, row 4, is under 70%, also not half.) The drill is the shared tank text in `starter_kit.DRILLS["tank"]["death"]`. | Needs a kit change: per-spec death drills (or an override hook for drills). Paladin: "Put Templar's Bulwark on a key you can hit without looking, and press it on a dummy pull when your health falls to about a third. Divine Shield and Divine Protection wait for Forbearance to fade." |
| PALADIN-R25-7 | CONFIRMED | coach rules: a core ability judged before it is trained | `coachrules.py` `paladin-protection` core; `api/internal/coach/f1_core_casts.go` `assessAbility` | core: "Hammer of the Righteous" (`talent: false`) | F1 skips a core entry at zero casts only when it is a talent (`if e.Talent && len(seq) == 0`). Hammer of the Righteous is a trainer spell (`trainables`: level 40), so `pickSpell` resolves it from spellconst at any level and `coreGaps` counts every 6 s of the fight with the mana in the bar as lost. A Protection tank at 30 to 39 who cannot have it gets "Hammer of the Righteous: N casts lost". F7 already gates defensives on `trainableAtLevel`; F1 does not. The same gap applies to Holy Strike below 6 and Judgement below 4 on both melee specs (rare logs). | Needs a coach code change (owning lane): in `assessAbility`, skip a non-talent entry with no casts when `trainableAtLevel(in, e)` finds no rank at the player's level, and report it under could-not-check when the log has no level. |
| PALADIN-R25-8 | LIKELY | coach rules: core ability the page tells you to hold | `coachrules.py` `paladin-protection` core "Swift Judgement" | (rule entry) | protection.md:103 "On a fight with a tank swap, hold one of them for the swap instead of spending it on threat", :146 common mistake "Burning both Judgements on threat right before a tank swap", card row 9 "hold it when a tank swap is coming". F1 counts every minute Swift Judgement sits ready as lost (fires at 2 lost and 10%). A tank who holds it two or three minutes for a swap, as the page says, is flagged. LIKELY because it fires only on swap fights long enough to lose two casts. | Drop "Swift Judgement" from `paladin-protection` core and rerun `python -m pipeline coach-rules`, or leave it and accept the flag on swap fights (then the coach md needs a hold line). Prefer the drop: Judgement (row 10) already carries the taunt cadence. |
| PALADIN-R25-9 | CONFIRMED | coach: stale override contradicts the line above it, the card and the page | `coach/paladin-retribution.md:25`; `starter_overrides.GUIDE_LINES["paladin-retribution"]["Seal of Righteousness"]` | "Seal of Righteousness about a second and a half before the pull. Reapply it just before it expires so it never lapses." | :24 (and retribution.md:78, card pre-pull row 0) puts Seal of the Crusader up at -1.5 s and Seal of Righteousness only after the first-swing Judgement. The card casts Righteousness before the pull only when Crusader is already on the target and you lack Unsealed Wrath (pre-pull row 2); with Unsealed Wrath it is Seal of Command (row 1). The override's comment quotes "reapply Seal of Righteousness just before it expires so uptime never lapses", text the page no longer has. | Needs an override edit: "Seal of Righteousness at the top rank you have trained, straight after the opening Crusader Judgement, and recast in its last second. With Unsealed Wrath it is the one-swing release in the twist, and your Seal whenever your mana falls behind the fight." |
| PALADIN-R25-10 | CONFIRMED | coach: priority order contradicts page and card | `coach/paladin-retribution.md:26-28`; no GUIDE_ORDER for `paladin-retribution` | Holy Strike (:26) listed ahead of Exorcism (:27) and Hammer of Wrath (:28) | retribution.md:84-89: Judgement, Exorcism, Hammer of Wrath, Seal, Holy Strike, Consecration; :99 "Below 20% Hammer of Wrath goes ahead of everything but Judgement and Exorcism". Card rows 4 and 5 sit above Holy Strike (row 10). The generator orders by first mention, and :80 names Holy Strike ("Your swings and Holy Strike keep the Crusader debuff up") before the numbered list. | Needs an override: GUIDE_ORDER `paladin-retribution` = ("Blessing of Might", "Judgement", "Exorcism", "Hammer of Wrath", "Seal of Righteousness", "Seal of Command", "Holy Strike", "Consecration"). |
| PALADIN-R25-11 | CONFIRMED | coach: the top button's line teaches only the opener; a reason the page contradicts | `coach/paladin-retribution.md:24, :26` | :24 (the Judgement entry) "If no other paladin has Judgement of the Crusader on the target, put up Seal of the Crusader..."; :26 "It spends the globals Judgement and seal upkeep leave idle." | Judgement's entry is the opener sentence (it leads with "Judgement of the Crusader"), so the list never says what card row 1 and page item 1 say: "Judgement on cooldown. It is off the global cooldown, so press it the moment it is ready." Judgement is a `core` entry the coach judges. The Holy Strike override then says Judgement leaves globals idle, which treats it as an on-global button; it is off the global cooldown (row 1 note, :84). | Needs an override: GUIDE_LINES `paladin-retribution` "Judgement": "Judgement on cooldown. It is off the global cooldown, so press it the moment it is ready." (the opener sentence moves to "Seal of the Crusader"); "Holy Strike": "Holy Strike on cooldown, every ten seconds. It spends the globals Seal upkeep leaves free." |
| PALADIN-R25-12 | CONFIRMED | coach: buttons named before the player has them | `coach/paladin-retribution.md:27, :28, :29, :32` | ":27 Exorcism on cooldown..."; ":28 Hammer of Wrath below twenty percent target health, at its highest rank."; ":29 Consecration last..."; ":32 ...Holy Wrath hits all of them" | Exorcism and Consecration train at 20, Hammer of Wrath at 44, Holy Wrath at 50 (`trainables`). Retribution is the page's leveling spec (:133), so boss pulls from Ragefire Chasm on are logged against these lines; at 15 to 43 the list tells the player to press a Hammer of Wrath they do not have "at its highest rank". | GUIDE_LINES `paladin-retribution`: "Hammer of Wrath": "Hammer of Wrath from forty-four, below twenty percent target health, at its highest rank."; "Exorcism": "Exorcism from twenty, on cooldown, against Undead and Demons only."; "Holy Wrath": prefix "From fifty, "; "Consecration": prefix "From twenty, ". |
| PALADIN-R25-13 | CONFIRMED | coach: the staple heal reduced to its fallback | `coach/paladin-holy.md:27` | "Flash of Light: Once you fall more than five points under, heal with Flash of Light alone until you catch up." | holy.md:78-79: "Flash of Light rank 5 on the tank below 90%. This is your staple", and rank 4 on everyone else below 80%; card rows 11 and 12. The coach's only Flash of Light line is the last sentence of item 3, so the list teaches Flash of Light as the heal for when you are behind on mana and never as the main heal. | Needs an override: GUIDE_LINES `paladin-holy` "Flash of Light": "Flash of Light from twenty on the tank below ninety percent is your staple, at a rank close to your level; on everyone else below eighty percent, yourself included. When your mana falls behind the fight, heal with it alone until you catch up." |
| PALADIN-R25-14 | CONFIRMED | coach: a button the page tells you to save, listed bare in the priority | `coach/paladin-holy.md:29` | "Light's Vigil." | holy.md:85 "Keep it for a fight where your whole party takes big damage at once. On an ordinary boss it is mana you need for the tank", and :130 lists "Casting Light's Vigil on an ordinary boss" as a mistake. The card never casts it. The bare name sits in "the buttons the guide's rotation names, in the order it gives them", between Holy Shock and Voice of Truth. The generator skipped :85's reason because it opens with "It" (`leans_on_the_last`). | Needs an override: GUIDE_LINES `paladin-holy` "Light's Vigil": "Light's Vigil from forty, only for a fight where your whole party takes big damage at once. On an ordinary boss its mana goes to the tank." |
| PALADIN-R25-15 | CONFIRMED | coach: talents named before the player has them | `coach/paladin-holy.md:23, :25, :26, :28` | ":23 Swear Oath of Light on the tank..."; ":25 Divine Favor, then Holy Light..."; ":26 ...Holy Shock on anyone below seventy percent"; ":28 Holy Shock only while..." | On the Holy order Divine Favor is the 21st-level point, Oath of Light the 25th, Holy Shock the 32nd, Light's Vigil the 40th (decoded `addon-builds.json` `paladin-holy`); Flash of Light trains at 20, Concentration Aura at 22. holy.md:117 and :123 say most Holy paladins level as Retribution, which has none of these talents. A dungeon healer at 20 to 31 is coached to swear an Oath and open Divine Favor they do not have. | GUIDE_LINES `paladin-holy`: "Oath of Light": "From twenty-five, swear Oath of Light on the tank a second or two before the pull."; "Divine Favor": "Divine Favor from twenty-one, then Holy Light, when the tank drops below forty percent."; "Holy Shock": "Holy Shock from thirty-two, only while your mana is keeping pace with the fight." |
| PALADIN-R25-16 | LIKELY | coach rules: uptime of one Seal on a spec whose raid build swaps Seals | `coachrules.py` `paladin-retribution` auras `("Seal of Righteousness", SELF)` | (rule entry) | With Unsealed Wrath, Seal of Command is "your main Seal while you can afford the twist" (retribution.md:93) and Righteousness is up for one swing in four. F2 judges self-buff uptime against the top-log median, else the sim. A raid-build paladin twisting as the page teaches shows Righteousness up a fraction of the fight, and is flagged whenever the pool's median comes from paladins who stay in Righteousness. The page's actual rule is "Letting the Seal drop" (:141), any Seal. LIKELY because it depends on the pool. | `coachrules.py`: judge one aura family covering the damage Seals (Seal of Righteousness and Seal of Command ids) as "your Seal", or drop the aura and leave the Seal to F8 adherence. Rerun `python -m pipeline coach-rules`. |

Counts: CONFIRMED 14, LIKELY 2, NEEDS-SIM 0, NEEDS-OWNER 0. Page prose and card notes: no finding (their round 24 fixes hold). All 16 are the new coach scope.

## Shared roots (filed elsewhere; named once for paladin, not counted)

- **Template meta talk:** all three "Who they are" paragraphs ("as the site's guide describes it ... not written down here yet ... a later full write"); the Voice block's "For a healer, the subject is lasting the fight" and "Never tell a healer to heal more" in the tank and damage files.
- **"Top rank" drill:** Retribution and Protection drill.ranks ("every button on the priority list should show its top rank"). The Retribution card casts Consecration rank 1 at 20 to 43 and 50 to 52 (Human to 53), rank 4 at bare 60; the Protection card casts rank 1 Consecration at 40 to 58.
- **"Recast with 3 seconds left":** Retribution drill.uptime. The Seals are recast in their last second (card rows 8 and 9, retribution.md:141).
- **Healer drills:** Holy drill.ranks ("a mid rank of your filler") and drill.resource ("lean on your cheapest heal"). For a Holy paladin the cheapest heal is rank 1 Flash of Light, which holy.md:78 and :127 call a trap at 60.
- **Combat potion by resource:** Protection lists Major Mana Potion; its raid preset (`sim/request/raid-consumes.json`) carries Major Spellblasting Potion. Retribution and Holy match their presets.
- **"No X" clauses excluding a spell spec-wide:** the cause of R25-1 (here a defensive sentence, not a leveling one).

## SHOULDs (wording, not findings)

- **protection.md:94** "20 mana": rank 8 at 60; the rotation is read at 60 on the page, but "20 mana at 60" would match card row 7 and fix R25-4 on regen.
- **coach/paladin-protection.md:28** "Consecration only with spare mana, where it pays for itself in threat": the page says it makes "far less threat per mana than Hammer of the Righteous" and is how you spend mana you will not need (:99). "Only above half mana, to spend mana you will not need" would restate it.
- **coach/paladin-protection.md:52** drill.uptime "The moment one drops, press it again": Holy Shield often drops on its fourth block before its 10 s cooldown is back.
- **coach/paladin-retribution.md:29** "Do not let your mana fall far behind the fight": the page's pace is exact (four fifths of the fight left). The override comment quotes older page text.
- **coach/paladin-retribution.md:30** "Seal of Command: Once you have..." (capital after the colon, `with_name` prefix) and no twist line: the list never says to swap to Righteousness for one swing.
- **coach/paladin-holy.md:26 and :28**: Holy Shock's mana rule appears twice. :31 Lay on Hands drops the page's "It drains all of your own mana" and "late in the fight".
- **coach/paladin-holy.md** drill.uptime "Recast your self buff" and the "Key buff and debuff uptime" and "Core abilities under-used" lines: the Holy rule file has empty `auras` and `efficiency`, so the coach can only report "doesn't know your spec's list yet" for both; a Holy paladin has no self buff to recast.
- **coachrules.py Protection** Judgement and Hammer of the Righteous are judged at base cost only; the page holds both for Holy Shield's reserve (:95, :97, :101). A narrow window, so not filed.
- **Kit:** Retribution and Protection list no weapon enchant, though their raid presets carry Elemental Sharpening Stone and Brilliant Wizard Oil.
- Carried from round 24, still true as SHOULDs: index.md:29 "for raiders"; Retribution row 11 parenthetical placement; "in a dungeon or a raid" (row 2, :78, now also coach :30); retribution.md:135 "when" clause; Human "from about 55" (rows say 54); retribution.md:43 "real damage"; row 6 "180 to 190"; retribution.md:72, :123, :36 ("Judge every 8 seconds", also the Retribution coach "Who they are"), :133; Holy row 12 and holy.md:81; protection.md:111 block chance and the Holy Shield ranks 2 and 3 tooltips (data lane).

## Checked and right this round

- **Page prose, full sweep.** Unchanged since round 24 apart from the three fixes above; every superlative, "only", "every" and quoted cost re-read against the unchanged rows (Holy Light the only castable heal from 25; Consecration costs and 8 s; Hammer of Wrath 6 s on all ranks; Lay on Hands 550 mana at rank 3 and 20 minutes; Light's Vigil 1,340 at rank 3; Righteous Fury 30% base mana; Holy Shield ranks at 30/40/50/60 match :139; Greater Blessing of Kings and Heart of the Lion both 10% to every stat, 409580).
- **Coach lines that hold:** Protection :23, :24, :27, :30, :32 to :35 (Righteous Fury 30 minutes; Seal of Fury at -3.5 s; Judgement taunts under Seal of Fury; Templar's Bulwark at about a third matches row 1's 35%; Forbearance; Lay on Hands to full and every point of mana, 633). Retribution :23, :24, :31. Holy :24, :30 (Voice of Truth 6 s, 3 minutes, 1310897, trains at 20), :31, :32. Consumables: flasks match the role; foods match each raid preset (Sage's Tea, Baked Salmon, Smoked Desert Dumplings).
- **Coach rule files:** Retribution core Judgement and Holy Strike are on-cooldown buttons on page and card. Holy downranks Holy Shock and Flash of Light match the page's rank advice; Holy Light is rightly not a downrank family. Defensive lists are gated on level in F7 (`trainableAtLevel`). F5 for the damage and tank specs honours `rank-choices.json`, so Consecration rank 1 is not flagged as an old rank.

## Decisions

- **D1. Coach lines are judged at every level a pull can be logged** (brief), as the druid report did: a button named with no level is a finding when a common logged band lacks it (R25-5, R25-12, R25-15). Lines conditional on the target (Exorcism on Undead) still name a spell the player may not have.
- **D2. R25-3 is CONFIRMED.** The dropped Holy Shield reserve contradicts :95 and :101 outright; the superlative is not in the page, and Holy Shield's per-cast damage at 60 is larger.
- **D3. R25-7 is CONFIRMED from the code**, not a log: the talent-only skip is explicit, and nothing upstream filters `Rules.Core` by level (`grep` of `tables.go`, `specrules.go`, `input.go`).
- **D4. R25-8 and R25-16 are LIKELY**: both need a fight shape (a tank swap; a pool median from non-twisters) to fire.
- **D5. Not re-filed:** the D4 list from round 24 (rows rendering before their talent, card action ranks, rank 9 Holy Light book, Exorcism never crits, Benediction and Hammer of Wrath, BiS race captions); the site card's per-race rank issue (sim round 22).

## Sims requested

- None new. Round 24's three optional runs stand (Consecration rank 1 casts against cooldowns at 30, 40 and 50; Human against Dwarf after T1; raid order twist at 40 to 58).

## Engine questions

- None. The coach-code item (R25-7) is for the coach lane, not the engine.

## Sim findings

- None.

## Sources

No web research was needed: the client data, the card rows and the coach code decided every call.

- Client 1.60.1.70790: `spelltips/paladin.json` (Holy Strike 679 to 10333, Holy Shield 20925 to 1239570, Hammer of the Righteous 407632, Light's Vigil 1310911 to 1311595, Voice of Truth 1310897, Lay on Hands 633/2800/10310, Seal of Fury ranks, Hammer of Wrath, Holy Wrath, Exorcism), `trainables/paladin.json`, `talents/paladin.json` (Benediction, Shieldsworn), `spelltips/hunter.json` (Heart of the Lion 409580), `addon-builds.json` (six paladin orders).
- Curated: `data/curated/apl/paladin-*.json`, `data/curated/rank-choices.json`, `data/curated/coach/paladin-*.{md,json}`, `data/pipeline/coachrules.py`, `sim/request/raid-consumes.json`, `logs/engine/mechanics/consumables/catalogue.json`.
- Coach: `api/internal/coach/f1_core_casts.go`, `f2_uptime.go`, `basis.go`, `f5_ranks.go`, `f7_death.go`, `rules.go`, `prompt.go`, `prompt_rules.txt`; `tools/coach/*` (scratch regen at `scratchpad/gr25/paladin-regen/`).
- Git: 28e3b4125 (round 24 paladin), diffed against ea9adea1c.

## Appendix A: Note sweep

Cards byte-identical to round 24's tree; each note re-read against its condition, both orders and the rank rows. Bands and talent levels as in round 24's appendix.

**Retribution** (bands 10 to 60, 19 bands; raid order differs from 40)
- Row 1 Judgement, 10-60: Judgement 4; off GCD; 6% base mana; Benediction full at 14; Sanctified Judgement 22, 60%. OK.
- Row 2 Seal of Command opener, 20-60: Unsealed Wrath conditional (none leveling, 40 raid); Crusader 6; three stacks. OK (SHOULD "dungeon or raid").
- Row 3 Seal of Righteousness opener, 10-60: top rank trained. OK.
- Row 4 Exorcism, 20-60: Undead and Demons. OK.
- Row 5 Hammer of Wrath, 44-60: 6 s all ranks; Champion of the Light 5/5 at 39. OK.
- Row 6 Seal of Righteousness release, 10-60: four-fifths pace; 180 and 189 at 60. OK (SHOULD).
- Row 7 Seal of Command back, 20-60: 15 s reset. OK.
- Row 8 Seal of Command keep up, 20-60: pace; last second. OK.
- Row 9 Seal of Righteousness otherwise, 10-60: matches rows 6 to 8 gates. OK.
- Row 10 Holy Strike, 10-60: 18 at 60 with Benediction (20 less 10%); Sacred Arbiter 27. OK.
- Row 11 Consecration, 20-60: reserve and pace; rank 4 at bare 60; 40% at halfway. OK (SHOULD placement).

**Protection** (bands 10 to 60, 18 bands; one order)
- Row 0 auto cooldowns: Lay on Hands below 10%, trains 10. OK.
- Row 1 Templar's Bulwark, 20-60: talent 20; 35%; 8 s; Forbearance; Divine Shield from 34. OK.
- Row 2 Divine Shield, 34-60: 25%; 10 s, 12 s from rank 2 at 50. OK.
- Row 3 Divine Protection, 10-60: 20%. OK.
- Row 4 Stoneform (Dwarf): 70%; 10% physical, 8 s, 3 minutes. OK.
- Row 5 Seal of Fury, 10-60: last 2 s; Holy Shield reserve from 30; 30 s. OK.
- Row 6 Holy Shield, 30-60: 10 s; four charges. OK.
- Row 7 Holy Strike, 10-60: "twenty at 60"; Iron Creed; Shieldsworn from 40. OK.
- Row 8 Hammer of the Righteous, 40-60: reserve; target plus three. OK.
- Row 9 Swift Judgement, 10-60: one minute, free, off GCD; hold for a swap. OK.
- Row 10 Judgement, 10-60: reserve from 30; 6% base mana; taunt. OK.
- Row 11 Exorcism, 20-60: reserve from 30. OK.
- Row 12 Consecration, 20-60: above half mana; Holy Shield from 30; Hammer of the Righteous from 40. OK.

**Holy** (bands 10 to 60, 16 bands; one order)
- Row 0 auto cooldowns: potions, Demonic Runes; Divine Favor kept off. OK.
- Row 1 Oath of Light, 25-60: talent 25; rank limit; 5 minutes. OK.
- Row 2 Divine Favor, 20-60: talent 21 (render at 20 is D5); 40%; 2 minutes. OK.
- Row 3 Holy Light below 40%, 10-60: Divine Favor from 21; rank 8 at 60; "From 25 ... keep casting". OK.
- Row 4 Holy Light below 70%, 10-60: 5-point pace, 45% halfway. OK.
- Row 5 Holy Shock, 30-60: Boundless Light from 34. OK.
- Row 6 Holy Shock self: OK.
- Row 11 Flash of Light tank below 90%, 20-60: Oath from 25; rank 4 at bare 60. OK.
- Row 12 Flash of Light others below 80%, 20-60: rank 1, rank 4 from 50. OK (SHOULD).

## Appendix B: Coach sweep

One line per coach line: file:line, what it claims, the level the player has it, verdict.

**paladin-retribution.md**
- :17 Who they are: overview copy; Judge every 8 s from 16, Hammer of Wrath from 44. Template root; SHOULD carried.
- :23 Blessing of Might (4). OK.
- :24 Judgement entry = Crusader opener (Crusader 6). R25-11.
- :25 Seal of Righteousness before the pull. R25-9.
- :26 Holy Strike, 10 s (6). Order R25-10; "globals Judgement leaves idle" R25-11.
- :27 Exorcism (20). R25-10, R25-12.
- :28 Hammer of Wrath (44). R25-10, R25-12.
- :29 Consecration (20). R25-12; SHOULD on "far behind".
- :30 Seal of Command with Unsealed Wrath (talent 26, Unsealed Wrath 40 raid). OK; SHOULD.
- :31 Seal of the Crusader description, 40 s. OK.
- :32 Holy Wrath (50). R25-12.
- :35 to :42 finding list: generic; ranks line is the top-rank root.
- :44 consumables: flask, Smoked Desert Dumplings, Major Mana Potion match the preset. OK.
- :48 drill.core_casts. OK. :49 drill.uptime: 3-second root. :50 idle OK. :51 resource OK. :52 ranks: top-rank root. :53 preparation OK. :54 death OK. :55 adherence OK.
- :59 to :68 Voice: healer lines in a damage file (template root).

**paladin-protection.md**
- :17 Who they are: Holy Shield (30), Holy Strike every 6 s (40). Template root.
- :23 Righteous Fury (10), 30 minutes. OK.
- :24 Seal of Fury at -3 s (10). OK.
- (missing) Holy Shield. R25-1.
- :25 Hammer of the Righteous (40). R25-2, R25-3, R25-5.
- :26 Holy Strike, twenty mana. R25-2 (order), R25-4.
- :27 Judgement taunt under Seal of Fury. OK.
- :28 Consecration. Order R25-2; SHOULD.
- :29 Greater Blessing of Kings (60). R25-5.
- :30 Devotion Aura (1). OK.
- :31 Swift Judgement (29), Hammer of the Righteous (40). R25-2, R25-5.
- :32 Exorcism (20). Order R25-2.
- :33 Templar's Bulwark (20), about a third. OK.
- :34 Divine Shield (34) and Divine Protection after Forbearance. OK.
- :35 Lay on Hands. OK.
- :38 to :45 finding list: ranks line is the top-rank root.
- :47 consumables: Major Mana Potion against the preset's Major Spellblasting Potion (combat potion root).
- :51 core_casts OK. :52 uptime SHOULD. :53 idle OK. :54 resource OK. :55 ranks: top-rank root. :56 preparation OK. :57 death R25-6. :58 adherence OK.
- :62 to :71 Voice: template root.
- Rule file: core Hammer of the Righteous R25-7; Swift Judgement R25-8; Holy Shield aura OK.

**paladin-holy.md**
- :17 Who they are: Oath (25), Flash of Light (20), Divine Favor (21). Template root.
- :23 Oath of Light (25). R25-15.
- :24 Concentration Aura (22). OK as a group-aura line.
- :25 Divine Favor (21) into Holy Light. R25-15.
- :26 Holy Light and Holy Shock (32) with the 5-point pace. R25-15 (Holy Shock).
- :27 Flash of Light fallback only. R25-13.
- :28 Holy Shock (32). R25-15; SHOULD (repeat).
- :29 Light's Vigil bare. R25-14.
- :30 Voice of Truth (20), 6 s, 3 minutes. OK.
- :31 Lay on Hands (10), 550 mana at rank 3, 20 minutes. OK; SHOULD.
- :32 Blessing of Protection. OK.
- :35 to :40 finding list: core and uptime lines unreachable (empty rule lists), SHOULD.
- :42 consumables: Flask of Distilled Wisdom, Sage's Tea, Major Mana Potion match the preset. OK.
- :46 core_casts (potions and runes, as holy.md:83). OK. :47 uptime SHOULD. :48 resource: cheapest-heal root. :49 ranks: mid-rank root. :50 preparation OK. :51 death OK.
- Rule file: downranks OK; core and efficiency empty (SHOULD).
