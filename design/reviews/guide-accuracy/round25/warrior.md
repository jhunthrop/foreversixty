# Guide accuracy round 25: Warrior

Audited on the round 25 tree (head 438f8ba15: main after guide round 24; no sim round 22 in the log). Engine pin 170a59c39 (`sim/enginever/version.go`). Active client 1.60.1.70790 under `data/builds/1.60.1.70790/`.

Files audited:

- `web/src/content/guides/warrior/{index,arms,fury,protection}.md`, line by line.
- Every note and condition in `data/curated/apl/warrior-{arms,fury,protection}.json`, against both shipped orders (`addon-builds.json`, decoded below) and `trainables/warrior.json`.
- New this round: the coach files `data/curated/coach/warrior-{arms,fury,protection}.md` and their rule files `warrior-*.json`, the warrior entries of `data/pipeline/coachrules.py`, and the coach code that reads them (`api/internal/coach/f1_core_casts.go`, `f2_uptime.go`, `f7_death.go`, `tables.go`).
- Generator: `tools/coach/persona_starter.py` with `starter_guide.py`, `starter_priority.py`, `starter_overrides.py` and `starter_kit.py`.
- Presets: `sim/request/raid-consumes.json` and `data/curated/presets.json`. The guide consumables table is `web/src/lib/guides/gear-extras.ts`.

Read first: round 24 warrior, the ledger's Round 24 section, and the round 25 rogue, paladin and warlock reports for the shared coach roots.

Coach scope follows PRIEST-R25 D1. The coach reviews a boss pull at the level the log records, dungeon bosses from the low teens included. A coach line must therefore hold at every level it can be read.

## Changes since round 24's tree (ea9adea1c)

- `git diff ea9adea1c HEAD` on the warrior pages, cards, coach files, `coachrules.py`, `tools/coach`, `addon-builds.json`, warrior BiS and `web/src/lib/guides` shows one line: fury.md:112, round 24's R24-1 fix.
- The coach files last moved in the round 21 regen. Arms is a generated starter (`status: starter`). Fury and Protection are hand-written full personas (`status: full`): the generator refuses them ("is a full persona; not generated"), so every fix there is a hand edit.

## Round 24 fix, re-checked

- **fury.md:112 (R24-1).** "Skyborne's Wind Blessed, 2% haste that never turns off, is level with Night Elf, and it needs the paid Skyborne Heroic Pack, so take it only if you already own it."
  - This matches sim round 21 F2: +0.02% and +0.03% at seeds 1 and 50001, inside the ±0.5% error.
  - It matches arms.md:109's "level" and the Horde sentence after it.
  - Skyborne is named only as the paid option. OK.

## Coach regeneration check

- `persona_starter.py --spec warrior-arms --out gr25/warrior-regen/warrior-arms.md` is byte-identical to the committed file, so no Arms finding is fixed by a regen.
- Fury and Protection are not generated.

## Orders (index 0 = level 10)

Arms:

| Talent | Leveling | Raid |
|---|---|---|
| Improved Heroic Strike | 10 | 10-12 |
| Deflection | 11-13 | |
| Improved Rend | 14-16 | 13-15 |
| Improved Charge | 17-18 | |
| Improved Tactical Mastery | 19-21 | 16-18 |
| Improved Overpower | 22-23 | 19-20 |
| Anger Management | 24 | 21 |
| Deep Wounds | 25-27 | 22-24 |
| Spearing Strike | | 25 |
| Two-Handed Weapon Specialization | 28-30 | 26-28 |
| Impale | 31-32 | 29-30 |
| Bloodthrill | 33-34 | 31-32 |
| Sweeping Strikes | 35 | 33 |
| Weaponmaster | 36-41 | 34-38 |
| Mortal Strike | 40 | 40 |
| Improved Slam | 42-43 | 39-41 |
| Cruelty | 44-48 | 42-46 |
| Unbridled Wrath | 49-53 | 47-51 |
| Improved Bloodrage | 54-55 | 60 |
| Iron Will | 56-60 | |
| Enrage | | 52-56 |
| Reckless Swings | | 57 |
| Improved Execute | | 58-59 |

Fury (the two orders are identical):

| Talent | Levels |
|---|---|
| Cruelty | 10-14 |
| Unbridled Wrath | 15-19 |
| Furious Precision | 20-22 |
| Enrage | 23-24 |
| Dual Wield Specialization | 25-29 |
| Death Wish | 30 |
| Reckless Swings | 31 |
| Improved Execute | 32-33 |
| Improved Berserker Rage | 34-35 |
| Improved Whirlwind | 36-37 |
| Flurry | 38-41 |
| Bloodthirst | 40 |
| Improved Heroic Strike | 42-44 |
| Improved Rend | 45-47 |
| Improved Charge | 48-49 |
| Improved Tactical Mastery | 50-52 |
| Anger Management | 53 |
| Deep Wounds | 54-56 |
| Impale | 57-58 |
| Improved Bloodrage | 59-60 |

Protection: as round 24 decoded it. Shield Slam is at 40, Last Stand at 28, and the build has no Improved Shield Wall.

Trainer levels: Whirlwind 36, Berserker Stance 30, Recklessness 50, Revenge 14 and Shield Block 16.

## Findings

| id | status | kind | where | quote | evidence | fix |
|---|---|---|---|---|---|---|
| WARRIOR-R25-1 | CONFIRMED | page: prose against the page's own consumables table | protection.md:109 | "Pre-pull, eat and drink for the health bar; a tank that dies holds no threat." | The page's table renders the Protection raid preset (`raid-consumes.json` `warrior-protection`, from 45; `gear-extras.ts` `consumableRows`). Its food is Smoked Desert Dumplings, which `CONSUMABLE_GIVES` labels "Strength (well fed)". `presets.json` by_spec `warrior-protection` gives the reason: "the dumplings measured 0.1% more threat than Poached Sunscale Salmon ... The engine models no stamina food a Phase 1 tank can cook". The drink (Rumsey Rum Black Label, Stamina) and Elixir of Fortitude are the health bar; the food is threat. Round 7 checked this line before the round 9 food measurement. | "Mighty Rage Potion is threat at the pull. Pre-pull, the rum and Elixir of Fortitude are your health bar, and the dumplings' Strength is threat and block value; a tank that dies holds no threat." |
| WARRIOR-R25-2 | LIKELY | page: prose against the page's own consumables table | arms.md:105 | "Pre-pull, take Elixir of the Mongoose, R.O.I.D.S., Winterfall Firewater and Smoked Desert Dumplings" | The Arms table (`warrior-arms` from 50) lists Elixir of the Mongoose, R.O.I.D.S., Juju Power, Juju Might, Smoked Desert Dumplings, Mighty Rage Potion and an Elemental Sharpening Stone. Winterfall Firewater is not in it. The prose leaves out Juju Power and Juju Might, and names an attack power drink in their place. It is LIKELY, not CONFIRMED, because I could not settle from a primary source whether Firewater stacks with Juju Might (Sources). Either way, the prose names a list the table does not show. | "Pre-pull, take Elixir of the Mongoose, R.O.I.D.S., Juju Power, Juju Might and Smoked Desert Dumplings: crit, Strength and attack power all stack with the raid's buffs." |
| WARRIOR-R25-3 | CONFIRMED | coach: lines that need a level, taken from the page's finished-build rotation | `coach/warrior-arms.md:25, :26, :27, :30, :32, :33` | ":33 Then swap to Berserker Stance and press Recklessness."; ":27 Mortal Strike goes on cooldown otherwise, and Reckless Swings right behind it in the raid build"; ":25 ... Bloodthrill only lights Overpower on a bleeding target."; ":26 ... even ahead of Mortal Strike ... after a dodge or a Bloodthrill proc"; ":30 keep the rage for Mortal Strike, Overpower and Execute" | Levels:<br>- Recklessness trains at 50, and the card's swap row asks `spellIsKnown 1719` ("From 50"). From 30 to 49, :33 sends the player into Berserker Stance at 20% with no Recklessness, where Overpower and Rend stop working (arms.md:91). The card does not swap there.<br>- Mortal Strike comes at 40 on both orders.<br>- Reckless Swings comes at 57 on the raid order.<br>- Bloodthrill comes at 33 on the leveling order and 31 on the raid order.<br>- Execute trains at 24.<br>The card notes already carry each level: "once Bloodthrill is in", "even ahead of Mortal Strike from 40", "Overpower from 12, Execute from 24 and Mortal Strike from 40", "From 50, below 20%". The generator reads the page's unqualified rotation instead. | Overrides in `starter_overrides.GUIDE_LINES["warrior-arms"]`, from the card notes:<br>- "Rend": "Keep Rend ticking: the bleed ignores armor, and once Bloodthrill is in, your main-hand hits on a target bleeding from your Rend can light Overpower too."<br>- "Overpower": "Overpower the moment it lights up, even ahead of Mortal Strike from forty: a dodge opens it for a few seconds, and a Bloodthrill proc does too once the talent is in."<br>- "Mortal Strike": "From forty, Mortal Strike on cooldown, and from fifty-seven in the raid build Reckless Swings right behind it."<br>- "Heroic Strike": "... keep the rage for Overpower from twelve, Execute from twenty-four and Mortal Strike from forty."<br>- "Recklessness": "From fifty, once that first Execute has emptied the bar, swap to Berserker Stance and press Recklessness. Before fifty, stay in Battle Stance."<br>Generator option (the shape of ROGUE-R25-8): prefix "From N," to any line whose spell or talent comes later than the earliest dungeon level. Not fixed by a regen. |
| WARRIOR-R25-4 | CONFIRMED | coach: a multi-target line lost its target condition | `coach/warrior-arms.md:34-35` | "Sweeping Strikes in Battle Stance, swap to Berserker Stance, Whirlwind." / "Whirlwind hits up to four targets ..." | On the page these sit under "**Two or more targets.**" (arms.md:93). The coach drops the bold lead and lists them inside "The priority, in words", after Recklessness, as a single-target step. A player who follows it on one boss leaves Battle Stance (no Overpower, no Rend) for a Whirlwind the page never casts single-target. They also need Sweeping Strikes (35 leveling, 33 raid) and Whirlwind (trains at 36). | Override "Sweeping Strikes": "With two or more targets, from thirty-six: Sweeping Strikes in Battle Stance, swap to Berserker Stance, Whirlwind, and swap back for Overpower and Rend once Whirlwind is on cooldown." Generator: `starter_guide` should carry a paragraph's bold lead ("Two or more targets.") as the condition of the lines under it. This is not the "only with N targets" root; the condition here is dropped, not misread. |
| WARRIOR-R25-5 | CONFIRMED | coach: short or missing spell names (shared root) | `coach/warrior-arms.md:28, :29` | ":28 Only with Improved Slam, and between Mortal Strikes from forty: the talent keeps your swing timer running through the cast."; ":29 Spearing Strike." | :28 never names Slam. It is the Slam card note with the action dropped, so it reads as an instruction about nothing. :29 is a bare name: it loses "raid build only", "on a free global" and "in Battle Stance". The leveling build never has Spearing Strike (raid order 25), and it stops at the execute swap (arms.md:87). | Overrides:<br>- "Slam": "Slam only with Improved Slam, between Mortal Strikes from forty: the talent keeps your swing timer running through the cast."<br>- "Spearing Strike": "In the raid build, Spearing Strike on any free global in Battle Stance: it stops at the execute swap."<br>The shared generator root as filed. |
| WARRIOR-R25-6 | CONFIRMED | coach: a rule dropped from a card condition | `coach/warrior-arms.md:31` | "Hamstring only with a Windfury Totem in the group, on a free global for ten rage." | The card's condition and note: "Only with a Windfury Totem on you, on a free global, and never from 25% down". arms.md:91: "From 25% stop Hamstring and let the rage build." The coach's execute line (:32) does not carry the stop either, so a Windfury group's Arms warrior is told to keep Hamstringing into the full-bar Execute. | Override "Hamstring": "Hamstring only with a Windfury Totem in the group, on a free global for ten rage, and never from twenty-five percent down: that rage is the first Execute." |
| WARRIOR-R25-7 | CONFIRMED | coach: site, sim, data and coach meta talk (voice rule; shared root, HUNTER-R25-13) | Arms :17, :21, :48 (generated); Fury :21, :37; Protection :21, :23, :25, :37, :39 (hand) | Arms: "as the site's guide describes it ... This file says what the guide says"; "the catalogue looks for". Fury: "This is the order the site's Fury rotation runs"; "the catalogue looks for". Protection: "every line in it was kept because taking it out made the tank worse"; "the simulated boss has none"; "Devastate is in the client's spell list"; "the catalogue looks for"; "the coach cannot judge your share of it" | The owner's guide-voice rule, which the round 25 brief applies to coach files. Protection :23 also undercuts its own page: protection.md:93 says Demoralizing Shout "takes a real slice off every swing" and keeps it on the boss. | Arms: the shared generator root (`persona_starter.who_section`, the `know_section` lead-in, `starter_kit.consumable_paragraph`).<br>Fury and Protection are hand edits:<br>- Fury :21 "This is the order they play."<br>- Protection :21 "This is the order they play."<br>- Protection :23 "Demoralizing Shout takes 204 attack power off the boss at sixty; keep it up unless a druid's Roar or another warrior's Shout is already there."<br>- Protection :25 "Devastate is not learnable in Forever."<br>- :37: "the consumables to bring are ...".<br>- Protection :39: "your share of the Sunders is lost among everyone's". |
| WARRIOR-R25-8 | CONFIRMED | coach: the "recast with 3 seconds left" drill (shared root) | `coach/warrior-arms.md:53` | "Put the timers of your buffs and debuffs where you can see them. Recast each with 3 seconds left" | Arms judges two auras (`coachrules.py`): Battle Shout and Rend. The card casts Battle Shout "Whenever it drops" (`not auraIsActive`) and Rend when it is not ticking (`not dotIsActive`). Rend recast with 3 s left clips a bleed tick for 10 rage. | The shared root: `starter_kit.DRILLS["dps"]["uptime"]`. Not fixed by a regen. |
| WARRIOR-R25-9 | CONFIRMED | coach: consumables against the spec's raid preset (shared root for Arms; hand edit for Fury) | Arms :44, :48, :57 (generated); Fury :17, :37, :48 (hand) | Arms: "the catalogue looks for a flask (Flask of the Titans) ... a weapon enchant (Elemental Sharpening Stone); a combat potion (Mighty Rage Potion)"; "Pull preparation: a missing flask, food, weapon enchant and combat potion." Fury: "never go into a boss without the flask, food and stone"; "Flask of the Titans, Smoked Desert Dumplings ..."; "check four things: flask, food buff, sharpening stone on both weapons, a Mighty Rage Potion" | The Arms and Fury raid presets carry no flask at any band (`raid-consumes.json`). They carry Elixir of the Mongoose, Juju Power, R.O.I.D.S. and Juju Might, which the Arms kit never names. The stone is Dense below 50 and Elemental from 50. Mighty Rage Potion is from 46 (`presets.json` by_class warrior, "client level 46"). So both files mark a correctly prepared Arms or Fury warrior as missing a flask. Below 46 they ask for a potion and a stone the player cannot have. Protection's kit (flask from 50) matches its preset at 60. | Arms: the shared root, `starter_kit.kit()` read from the top-band preset. Fury hand edit:<br>- :17 "the food, elixirs and stones the pull asks for".<br>- :37 "Smoked Desert Dumplings, Elixir of the Mongoose, Juju Power, R.O.I.D.S. and Juju Might, a sharpening stone on each weapon (Elemental from fifty), and from forty-six the Mighty Rage Potion".<br>- :48 "food, elixirs, a stone on both weapons, and from forty-six a Mighty Rage Potion". |
| WARRIOR-R25-10 | CONFIRMED | coach: "What changed for Forever" states changes that are not changes | `coach/warrior-fury.md:25` | "Unbridled Wrath gives its rage with any weapon, not only a two-hander. Bloodthirst scales harder with attack power than it once did" | Unbridled Wrath on 70790: "a 60% chance to generate 1 additional Rage when you deal melee damage with a weapon". In Classic it was also any weapon (Wowhead Classic 13002: proc on any weapon damage, 40% chance at 5/5). The Forever change is the chance, 40% to 60%. Bloodthirst 23894 on 70790: "45% of your Attack Power plus 48". Classic 23894 is 45% too (Wowhead Classic, effect value 45). The attack power scaling is unchanged; Forever added a small flat amount and 10% movement speed for 10 seconds. | Hand edit: "Unbridled Wrath's chance is up to sixty percent. Bloodthirst keeps its 45% of attack power, adds a little flat damage, and gives 10% movement speed for ten seconds." |
| WARRIOR-R25-11 | CONFIRMED | coach: a superlative and an "every time" the card's own order breaks | `coach/warrior-fury.md:21, :17` | ":21 Otherwise Bloodthirst whenever it is off cooldown, since it is the highest damage per rage in the kit."; ":17 Bloodthirst goes out the moment it is ready, every time" | Per rage:<br>- Execute 20662 is 600 damage for 10 rage with Improved Execute (Fury order 32 to 33), which is 60 a point before any extra rage.<br>- Bloodthirst is 45% of attack power plus 48 for 30 rage: at about 2,400 raid attack power, 1,128, or about 38 a point.<br>- Whirlwind, two weapons with Improved Whirlwind's +30% for 25 rage, is level with or above Bloodthirst.<br>So Bloodthirst is not the kit's best per rage. The page's reason is "the button you lose the most by delaying" (card note, fury.md:96). For :17, the card puts Execute above Bloodthirst below 20% (fury.md:90; the note "from 40 it goes ahead of Bloodthirst too"), so Bloodthirst does not go out "every time". | Hand edit:<br>- :21 "... since it is the button you lose the most by delaying."<br>- :17 "Bloodthirst goes out the moment it is ready, until Execute takes over at twenty percent". |
| WARRIOR-R25-12 | CONFIRMED | coach: card buttons missing from the priority, and the adherence drill tells the player to leave them | `coach/warrior-fury.md:21, :50` | The priority (:21) has no Berserker Rage and no Overpower. ":50 play only Bloodrage, Death Wish, Battle Shout, Execute, Bloodthirst, Whirlwind, Reckless Swings and Heroic Strike for three minutes. Anything else on your bars waits." | The card has two rows the coach leaves out:<br>- Berserker Rage from 36 under 80 rage ("Improved Berserker Rage makes it 10 rage, back every 30 seconds"). fury.md:86 puts it in the priority right behind Reckless Swings.<br>- Overpower before 36 ("Only before level 36, while you still fight in Battle Stance"), and fury.md:134.<br>The drill tells the player to leave both waiting. | Hand edit:<br>- In :21, after Reckless Swings: "From thirty-six, Berserker Rage right behind it whenever you are under eighty rage, held for the fear on Magmadar, Onyxia and Nefarian. Before thirty-six you fight in Battle Stance, so Overpower whenever it lights."<br>- In :50, add Berserker Rage to the list. |
| WARRIOR-R25-13 | CONFIRMED | coach: talent and trainer buttons taught with no level (the shape of ROGUE-R25-8) | `coach/warrior-fury.md:17, :21, :23, :43, :45, :47` | "Bloodthirst goes out the moment it is ready"; "Whirlwind on cooldown right behind it, even on a single target"; "Death Wish and Recklessness for the execute phase"; drill "press Bloodthirst the moment it comes off cooldown"; "the first button after a reposition is Bloodthirst" | Levels: Bloodthirst 40, Whirlwind 36, Death Wish 30, Reckless Swings 31, Recklessness 50. Fury pulls are logged from the low teens. The persona is built around Bloodthirst with no level anywhere, while the page carries the levels (fury.md:134). | Hand edit: one sentence in :21, "Before forty you have no Bloodthirst and before thirty-six no Whirlwind: Overpower in Battle Stance, Heroic Strike above sixty, and Execute from twenty-four carry you", and "from forty" on the Bloodthirst drills. |
| WARRIOR-R25-14 | CONFIRMED | coach: card buttons missing from the priority and from the adherence drill | `coach/warrior-protection.md:21, :50` | The priority has no Demoralizing Shout and no Battle Shout. ":50 play only Shield Block, Revenge, Shield Slam, Thunder Clap, Sunder Armor and Heroic Strike ..., with Bloodrage when rage is low. Anything else waits." | The card has two rows the coach leaves out:<br>- Demoralizing Shout, after Thunder Clap ("Keep it on the boss ... Refresh it in its last 3 seconds"). protection.md:91 says "Demoralizing Shout comes next".<br>- Battle Shout "Whenever it drops".<br>The drill tells the tank to leave both waiting. Demoralizing Shout is also one of this spec's three judged uptime auras (`coachrules.py`), so the persona tells the player to skip a button the coach then grades. | Hand edit:<br>- In :21, after Thunder Clap: "Demoralizing Shout next, refreshed the same way, unless a druid's Roar or another warrior's Shout already holds it. Battle Shout whenever it drops and nobody else's is up."<br>- In :50, add Demoralizing Shout. |
| WARRIOR-R25-15 | CONFIRMED | coach: buttons taught with no level | `coach/warrior-protection.md:17, :21, :30, :43, :45, :46` | "Shield Slam goes out on cooldown"; "Shield Slam on cooldown, the biggest hit you have"; "Shield Slam ready and unpressed"; drills "press Shield Slam the moment it is ready", "Your first button after the move is Shield Slam or Revenge" | Shield Slam is the 31st point, at 40 (protection.md:135). Revenge trains at 14 and Shield Block at 16. Tank pulls from the low teens read a persona whose first button they lack. | Hand edit: "From forty" on the Shield Slam lines. In :21: "Before forty, Revenge and Sunder Armor carry your threat." |
| WARRIOR-R25-16 | CONFIRMED | coach rules: a core ability judged before it is trained (shared root, PALADIN-R25-7) | `coachrules.py` `warrior-fury` core "Whirlwind"; `f1_core_casts.go` `assessAbility` | (rule entry) | Whirlwind 1680 is a trainer spell (36), `talent: false`. F1 skips a zero-cast entry only when it is a talent. A Fury pull at 20 to 35 gets "Whirlwind: N casts lost" for every 10 s with 25 rage in the bar, for a button the player cannot have. Bloodrage (trains at 10) has the same gap in all three specs, but only on logs below 10. | The shared code fix: `assessAbility` skips a non-talent entry with no casts when `trainableAtLevel` finds no rank at the player's level. Nothing to change in `coachrules.py`. |
| WARRIOR-R25-17 | CONFIRMED | coach rules: an uptime the page says to leave to someone else | `coachrules.py` `warrior-protection` auras `("Demoralizing Shout", DEF)`; `f2_uptime.go` `judgeAura` | (rule entry) | protection.md:93 and the card note say not to press it when "a druid or another warrior already holds it: two do not stack". F2 handles the two cases differently:<br>- Another warrior's Demoralizing Shout is the same family, so it reads as "another player applied it" (`appliedByOthers`). Correct.<br>- A druid's Demoralizing Roar is a different family. The tank's own track is nil, `appliedByOthers` does not match it, and the uptime is measured as 0% against the pool. The tank who follows the page is flagged.<br>The bear entry judges Demoralizing Roar, so the case is common in raids. | `coachrules.py`: give the Demoralizing Shout rule the Demoralizing Roar ids as "same slot" ids that `appliedByOthers` reads (needs a small `AuraRule` field and an F2 check). Or drop the aura from the tank list and leave it to adherence. Rerun `python -m pipeline coach-rules`. |

Counts: CONFIRMED 16, LIKELY 1, NEEDS-SIM 0, NEEDS-OWNER 0.

By location and fix:

- Page findings: 2 (R25-1 and R25-2). Both are guide edits.
- Coach findings: 15, routed as follows:
  - Fixed by a regen: 0.
  - Overrides in `starter_overrides.py`: R25-3, R25-4, R25-5 and R25-6.
  - Generator or kit changes: R25-4 (bold lead as condition), R25-5 (short names), R25-7 Arms (template), R25-8 (`DRILLS` uptime) and R25-9 Arms (`kit()`).
  - Hand edits (Fury and Protection are full personas): R25-7, R25-9 Fury, R25-10 to R25-15.
  - `coachrules.py`: R25-17.
  - Coach code: R25-16 (shared) and R25-17.

## Shared roots: which apply to warrior

| Root | Warrior status |
|---|---|
| Template meta talk | R25-7 (Arms generated; Fury and Protection carry their own hand-written meta lines) |
| "Recast with 3 seconds left" | R25-8 (Arms). The hand-written Fury "10 seconds left" is a Should (S2) |
| Tank uptime drill | Not applicable. Protection is hand-written, and its drill (:44) watches the Shield Block buff, which matches the card |
| Consumables by resource, not preset | R25-9 (Arms). Fury repeats the flask in hand-written text |
| "Top rank" drill | Not filed (D3) |
| "no X" / RULED_OUT_NEXT | Did not fire for warrior. The Arms priority keeps every card row (I traced it against the card) |
| Short or missing names | R25-5 |
| "only with N targets" read as a group condition | Not present. The Arms multi-target condition is dropped, not misread (R25-4) |
| Values true only at 60 stated at every level | R25-3, R25-9, R25-13, R25-15 |
| "Only as an opener" against the card | Not present |
| F1 level gate (PALADIN-R25-7) | R25-16 |

On the "top rank" drill (D3): every warrior row casts the top trained rank except Battle Shout rank 6, Heroic Strike rank 8 and Revenge rank 5 at 60. Those three are the routed W23-1 engine question (trainer rows against the cards). Every warrior rank row is strictly better than the one below, so the drill is true for this class.

## Coach rules (`coachrules.py`), otherwise

- **Arms core.**
  - Bloodrage: its cost is health, so it gets no resource check. The card holds it at 80 or more rage, but the gaps are far below one 60 s cooldown, so it cannot reach CoreFiresLost 2.
  - Mortal Strike: a talent, so it is skipped until it is cast. OK.
- **Arms auras.**
  - Battle Shout: a self buff, read on the player, so another warrior's shout counts. OK.
  - Rend: drops in the execute swap by design; it is judged against the pool. OK.
- **Fury core.**
  - Death Wish: needs 2 lost casts (over 6 min ready). The page's rule presses it at the pull whenever a second fits, so the hold never reaches that. OK.
  - Bloodthirst: talent. OK.
  - Whirlwind: R25-16.
- **Protection core.** Bloodrage under 60 and Shield Slam. OK.
- **Protection auras.**
  - Shield Block and Thunder Clap: OK.
  - Demoralizing Shout: R25-17.
- **Defensives.**
  - Protection: Shield Wall and Last Stand. OK.
  - Arms: Last Stand, which is a Protection talent in neither Arms order. It is never judged (talents count only once cast), so S1.

## Shoulds (wording, not findings)

| id | where | note | suggested |
|---|---|---|---|
| WARRIOR-R25-S1 | `coachrules.py` `warrior-arms` defensives | Last Stand is in neither Arms build. It never fires, but it hides the "no defensives list" line, and it disagrees with Fury's comment ("Last Stand is a protection talent"). | `defensives=[]`, with the same comment as Fury. |
| WARRIOR-R25-S2 | `coach/warrior-fury.md:44` | "Recast it with 10 seconds left". The card casts Battle Shout "Whenever it drops". An early recast is 10 rage and a global, not an error. | "Recast it the moment it drops ...". |
| WARRIOR-R25-S3 | `coach/warrior-protection.md:49` | The drill says "wall" at half health; the card and the page say Shield Wall below 40%. | "... when your health drops toward forty percent". |
| WARRIOR-R25-S4 | `coach/warrior-protection.md:23` | Shield Wall "a once-a-fight button" is true of the shipped build (15 min). protection.md:83's Improved Shield Wall swap makes it 4 min. | "... once a fight, unless you took Improved Shield Wall". |
| WARRIOR-R25-S5 | `coach/warrior-fury.md:28` | "because Whirlwind or Heroic Strike went first". Heroic Strike costs no global and "never delays Bloodthirst" (fury.md:88); it can only starve it of rage. | "... or Heroic Strike spent the rage it needed". |
| WARRIOR-R25-S6 | Fury and Protection frontmatter | `sources` still name `1.60.1.70755` (Fury) and `1.60.1.70630` (Protection) spellconst. This is not player-facing. | Restamp to 70790 at the hand edit. |
| WARRIOR-R25-S7 | Fury :39 and Protection :39 | "A missing flask is a smaller deal on a ten-player boss" is an unsupported claim. Fury brings no flask at all (R25-9). | Drop the sentence. |
| WARRIOR-R25-S8 | world buffs (all three files) | As ROGUE-R25-S4 and MAGE-R25 D4. | As filed there. |
| WARRIOR-R25-S9 | carried | Every round 24 Should, unchanged lines. | As filed. |

## Prose, line by line

The pages are unchanged since round 24 except fury.md:112 (above). I re-ran round 24's sweep on this tree, against the cards, both orders, `trainables` and the client rows, and it stands:

- index.md :22 to :35.
- arms.md :42 to :131.
- fury.md :41 to :134.
- protection.md :47 to :135.

This round I added the consumables lines against the rendered preset tables. That check gives R25-1 and R25-2. fury.md:108 ("Enchant and sharpen both weapons") matches the off-hand and main-hand imbues in `warrior-fury`.

Voice checks: no em dashes, no site or sim meta talk, and no page names 70755.

## Decisions

- **D1.** Coach scope by level follows PRIEST-R25 D1. R25-3, R25-13 and R25-15 are findings although each line is true at 60. As with ROGUE-R25 D4, each is one finding per file, because the lines share one shape.
- **D2.** R25-2 is LIKELY, not CONFIRMED. The prose does not match the table, but I could not settle whether Winterfall Firewater stacks with Juju Might from a primary source, so the "it does not stack" reason is not asserted.
- **D3.** The top-rank drill is not filed for warrior (Shared roots).
- **D4.** R25-16 is the shared F1 level gate. It is filed once here for Whirlwind; the code fix is PALADIN-R25-7's.
- **D5.** Nothing routed in the ledger's Round 19 to 24 sections is re-filed: W23-1 trainer ranks, S2 Defiance, the Defensive Stance hover tooltip, and the site card's per-race ranks.

## Sims requested

- None new. Carried from round 24, only if the wording is kept:
  - Protection Revenge against Shield Slam per landed cast (protection.md:69).
  - Protection race TMI on one shared preset.
  - Arms rage per second, raid against bare.

## Engine questions

- None new. Carried: S2, W23-1/W22-1, W21-1, X1, P5/S3, AC, C2/S4, and whether a missed Heroic Strike spends a Reckless Swings charge.

## Sim findings

- None.

## Sources

- Client 1.60.1.70790:
  - `spelltips/warrior.json`: 23881 to 23894 Bloodthirst, 20662 Execute, 1680 Whirlwind, 1719 Recklessness, 2687 Bloodrage, 11551 and 25289 Battle Shout.
  - `talents/warrior.json`: Unbridled Wrath, Improved Whirlwind, Enrage.
  - `spellconst/warrior.json`, `trainables/warrior.json`, `addon-builds.json`.
- Presets and tooling:
  - `sim/request/raid-consumes.json` (`warrior-*`, every band).
  - `data/curated/presets.json` (by_class warrior, by_spec warrior-*).
  - `web/src/lib/guides/gear-extras.ts`.
  - `data/pipeline/coachrules.py`, `data/curated/coach/warrior-*.{md,json}`.
  - `api/internal/coach/{f1_core_casts,f2_uptime,f7_death,tables,rules}.go`.
  - `tools/coach/*`; scratch regen at `scratchpad/gr25/warrior-regen/warrior-arms.md`.
- Reviews: round24/warrior.md, ledger Round 24, round25/{rogue,paladin,warlock}.md.
- Web:
  - Wowhead Classic, Unbridled Wrath 13002, https://www.wowhead.com/classic/spell=13002/unbridled-wrath (aura proc chance 40%, proc on weapon damage).
  - Wowhead Classic, Bloodthirst 23894, https://www.wowhead.com/classic/spell=23894/bloodthirst (effect value 45, 30 rage, 6 s).
  - Warcraft Wiki, Winterfall Firewater, https://warcraft.wiki.gg/wiki/Winterfall_Firewater (stacking with Juju Might not stated; D2).

## Note sweep

No card note, row or order changed since round 24, so round 24's per-row sweep is re-affirmed at every band, both orders, under round 25's lenses. Every row is OK:

Arms:

- Berserker Stance swap (30 to 60).
- Bloodrage.
- Battle Shout.
- Rend.
- Overpower (12 to 60).
- Mortal Strike (40 to 60).
- Reckless Swings.
- Execute (24 to 60).
- Slam.
- Spearing Strike.
- Hamstring.
- Heroic Strike.

Fury:

- Bloodrage.
- Death Wish.
- Recklessness (50 to 60).
- Battle Shout.
- Execute.
- Bloodthirst (40 to 60).
- Whirlwind (36 to 60).
- Reckless Swings.
- Berserker Rage (36 to 60).
- Overpower (12 to 35).
- Heroic Strike.

Protection:

- Defensive Stance.
- Last Stand.
- Shield Wall (28 to 60).
- Bloodrage.
- Shield Block (16 to 60).
- Revenge (14 to 60).
- Shield Slam (40 to 60).
- Thunder Clap.
- Demoralizing Shout (14 to 60).
- Battle Shout.
- Sunder Armor.
- Heroic Strike.

The values each row was checked against are listed line by line in round24/warrior.md, "Note sweep".

## Coach line sweep

warrior-arms.md (generated):

| Line | Verdict |
|---|---|
| :17 | R25-7 |
| :21 | R25-7 |
| :23 Bloodrage under 80 | OK |
| :24 the pull | OK |
| :25 | R25-3 |
| :26 | R25-3 |
| :27 | R25-3 (the "two or three swings" Should is carried) |
| :28 | R25-5 |
| :29 | R25-5 |
| :30 | R25-3 |
| :31 | R25-6 |
| :32 | R25-3 ("biggest hit" Should is carried) |
| :33 | R25-3 |
| :34-35 | R25-4 |
| :36 Cleave | OK |
| :39 to :43, :45, :46 mistakes | OK |
| :44 | R25-9 |
| :48 | R25-7, R25-9 |
| :52 core_casts | OK |
| :53 | R25-8 |
| :54 idle | OK |
| :55 resource | OK |
| :56 ranks | OK (D3) |
| :57 | R25-9 |
| :58 death | OK |
| :59 adherence | OK (the list is the card's) |
| :63 to :72 voice | OK |

warrior-fury.md (hand):

| Line | Verdict |
|---|---|
| :17 | R25-9, R25-11, R25-13 |
| :21 | R25-7, R25-11, R25-12, R25-13. The Reckless Swings "healers" Should is carried; the Heroic Strike 25% stop is OK from 24 |
| :23 cooldowns | R25-13 for Recklessness at 50; otherwise OK (Death Wish 3 min and 5% damage taken; Bloodrage 60 s, 20% health) |
| :25 | R25-10. Whirlwind both weapons, Reckless Swings and plate at 40 are OK |
| :28 | S5 |
| :29 to :35 | OK |
| :37 | R25-7, R25-9 |
| :39 | S7 |
| :43 | R25-13 |
| :44 | S2 |
| :45 | R25-13 |
| :46 | OK |
| :47 | OK (D3) |
| :48 | R25-9 |
| :49 | OK |
| :50 | R25-12 |

warrior-protection.md (hand):

| Line | Verdict |
|---|---|
| :17 | R25-15 |
| :21 | R25-7, R25-14, R25-15. Revenge "most threat per rage" is OK; "biggest hit" Should is carried |
| :23 | R25-7, S4. Last Stand and Shield Wall order matches 35% and 40% |
| :25 | R25-7. Shield Specialization and Master of Defense, shield needs and plate at 40 are OK |
| :28 to :35 | OK |
| :37 | R25-7. The kit matches the preset at 60 |
| :39 | R25-7, S7 |
| :43 | R25-15 |
| :44 | OK |
| :45 | R25-15 |
| :46 | OK |
| :47 | OK (D3) |
| :48 | OK at 50 and above |
| :49 | S3 |
| :50 | R25-14 |
