# Task 24 - wow-player read 4 of the 25 starter coach personas (2026-10-27, commit bf55fca68)

Verdict: 15 SHIP / 10 FIX.

## Patterns from read 3, checked file by file
1. Bare line for an out/conditional ability: FIXED for Moonfire (balance), Moonfire/Wrath/Claw (cat), Feral Charge (bear), Flame Shock (elemental), Frostfire Bolt (fire), Heroic Strike (arms), Sinister Strike (assassination), Seal of Command (ret). New form of the same defect: conditional abilities that DO belong now appear bare or not at all (see rulings).
2. Changelog and sim voice: FIXED ("this week", "now the only", "is back in the rotation", "raid-ready run", "tracked by slot" are gone). Left: "The curated rotation ..." opens Holy Paladin and Discipline (dev word; say "The rotation").
3. Orphan openers: FIXED (no "The instant Healing Wave.", no Elemental sim line, Fire "the buff" now has Pyroblast as subject).
4. Duplicates: Frenzied Regeneration, Hammer of the Righteous, Water Shield, Divine Favor, Disc Heal are FIXED. Remaining lighter form: a sentence followed by the same ability as a bare line (Rejuvenation/Tranquility/Innervate in restoration druid, Ferocious Bite in cat, Prayer of Mending/Power Word: Shield in Disc, Divine Protection in Prot Paladin). Harmless, but it reads like a list printed twice.
5. Guide order: FIXED for bear (Primal Bite, Lacerate, Maul, Swipe, Barkskin, Frenzied, Berserk = bear guide 72-78) and balance. NOT fixed for cat and Prot Paladin (below).
6. Healer drill.resource: FIXED. Every healer now reads "lean on your cheapest heal and save the fast, dear heal", true for Holy Paladin.
7. Forever-only claims (Stormstrike 8 s, Hemorrhage 145%/15%, Devouring Plague "biggest nuke", Riptide "most mana-efficient heal", Flash of Light "most efficient heal per mana"): all are in the guide or APL notes, so not generator errors. Still tooltip-check Stormstrike (Classic 10 s), Flash of Light vs Holy Light, Riptide. Shadowfiend/Dark Sacrifice on Disc is gone. Unsourced no longer.

## Previously-SHIP nine
- hunter-beast-mastery SHIP: leads with Aspect of the Hawk pre-pull and first Aimed Shot, which is exactly what a hunter does. Nits: "hence the reserve floor" is bot jargon (say "so keep a mana reserve"); the Rapid Fire / auto shot sentence is still hard to parse.
- hunter-marksmanship SHIP: Aimed Shot first, pre-pull line second. Same two nits.
- mage-arcane, mage-frost (loses "Shoot." correctly; Frostbolt / Ice Lance on a Fingers of Frost charge is enough), rogue-combat, rogue-subtlety, warlock-affliction, warlock-destruction: SHIP unchanged. Subtlety Hemorrhage numbers still a tooltip check.
- hunter-survival: was SHIP, now FIX only because of the Wing Clip ruling below.

## SHIP (15)
druid-balance ("Moonfire." gone; Insect Swarm, Starfire, Wrath weave per balance.md 66), druid-feral-bear, druid-restoration (Wild Growth/Tranquility line is verbatim restoration.md 62; the guide's "Tranquility when two are well down" is itself odd for a 10 s channel, accuracy reviewer should look), hunter-beast-mastery, hunter-marksmanship, mage-arcane, mage-frost, priest-discipline (Penance first, no Shadowfiend), rogue-assassination (Sinister Strike gone), rogue-combat, rogue-subtlety, shaman-restoration (order matches restoration.md 62), warlock-affliction, warlock-demonology (changelog and Shadow Bolt contradiction gone; Soul Fire, Searing Pain, Shoot, Life Tap), warlock-destruction.

## FIX (10), corrections quoted and sourced
- druid-feral (cat): "Open with Tiger's Fury..." then bare "Shifting Power." LAST. feral.md 68: Shifting Power "sits ahead of Tiger's Fury and every builder in the order", pressed "only when your energy is 60 or lower". Move it first with that condition; drop the bare "Ferocious Bite." duplicate.
- hunter-survival: add "Wing Clip, only with a Windfury Totem in the group, on a global with nothing else ready: it loses damage without the totem" (survival.md 85). "Mongoose Bite is a counterattack." should say "only after you dodge, or while Expose Prey is up" (survival.md 85).
- warrior-arms: add "Hamstring, only with a Windfury Totem in the group, on a free global (10 rage)" (arms.md 77). "Bloodrage." bare should carry "whenever Rage is below 80" (same line).
- mage-fire: "Scorch to apply or refresh the five stacks..." makes Scorch a debuff chore. fire.md 69: "Scorch is the filler for every global that isn't spent on a three-stack Pyroblast, Fire Blast, or Combustion". Rewrite to "Scorch on every free global; it also keeps Improved Scorch at five stacks". Pyroblast at three stacks is now correct.
- paladin-holy: bare "Holy Shock." drops holy.md 65 "only fire while your mana is keeping up with the share of the fight that is left"; add it ("Holy Shock only while your mana is on pace"). Replace "The curated rotation".
- paladin-protection: order follows the APL (Judgement before Holy Shield). prot.md 94-99: Righteous Fury, Seal of Fury, Holy Shield, Hammer of the Righteous and Holy Strike, Judgement, Consecration with spare mana. Consecration is missing. "Swift Judgement." is a talent, not a button: fold into the Judgement line. Delete bare "Divine Protection." duplicate.
- paladin-retribution: the ongoing Seal of Righteousness refresh is missing (retribution.md 65: "reapply Seal of Righteousness just before it expires", and after every Judgement), only the pre-pull cast remains. Hammer of Wrath line ("the execute-window button Retribution has besides seals, Judgement and Holy Strike") is garbled; say "Hammer of Wrath below twenty percent target health". Bare "Consecration." needs its condition (retribution.md 69: with mana support, only above 65 percent mana; a loss on a bare character).
- priest-shadow: the long Shadow Word: Death sentence is accurate to shadow.md 58 but unreadable, and bare "Mind Blast." in third place contradicts the last line ("Mind Flay on every free global comes before Mind Blast"). Rewrite as: "Shadow Word: Pain up. Inner Focus into Devouring Plague. Mind Flay on every free global. Mind Blast only when Mind Flay cannot be cast. Below twenty percent health, Shadow Word: Death first; while it is on its fifteen-second cooldown, Mind Blast, then Mind Flay." Add "Wand shots fill a global with nothing else ready" (shadow.md 58).
- shaman-elemental: lines say "Mana Spring Totem goes down first" and "Lava Burst is the first line" while sitting last and fifth, and bare "Magma Totem." right after "Single target: Searing Totem" reads as both. Reorder: Mana Spring Totem (pre-pull), Lava Burst on cooldown if talented, Searing Totem on one target, Lightning Bolt on every free global, Chain Lightning and Magma Totem only against two or more targets (elemental.md 56).
- shaman-enhancement: bare last "Earth Shock." drops "only while mana is above half" (enhancement.md 58), and the guide calls it a large gain, so unconditioned it becomes a mana-dump spam. Opening bare "Windfury Weapon." should read "Windfury Weapon on your one weapon before the pull and whenever it falls off (Rockbiter before 30)". Single weapon is correct. Stormstrike "eight second cooldown": tooltip check.

## Rulings on the six questions
1. Wing Clip / Hamstring: a conditional ability must appear WITH its condition, not be dropped. I approved Wing Clip with its condition in read 3; the condition is a plain player-checkable fact (a Windfury Totem in the group), the guide names the ability positively and measures it (+9%, +7%), and dropping it means the coach can never explain the hunter's or warrior's biggest raid filler. Amend P1: "named positively, possibly conditionally; carry the condition". Never bare.
2. Elemental in guide order: not acceptable. When the source text itself says "first" three times, the list must put those first; a player reading "Mana Spring Totem goes down first" in last place will stop trusting the list. Guide order governs ties, not the guide's own "first/last" statements.
3. Shadow's long SW:D sentence: faithful but FIX; split into the two short sentences above.
4. Enhancement's bare "Windfury Weapon.": acceptable as a bare line only if it is stated as a pre-pull step; a bare name leading the list tells nobody when. Add the clause.
5. Fire's Scorch note: wrong emphasis. Say Scorch is the filler for every global; "apply or refresh" is a subset.
6. Demonology without a pet line: acceptable in the priority list (the guide has no pet button). A sourced one-liner in "Who they are" is welcome (demonology.md 46-47: a dead pet must be re-summoned, costing mana and time; Demonic Knowledge needs the demon out) but not required to ship.

## Remaining generator-level patterns
A. P1 as written strips conditional abilities (Wing Clip, Hamstring, Consecration, Earth Shock, Holy Shock). Carry the condition instead of dropping or emitting bare.
B. Order: lines still come from the APL index for Prot Paladin and cat; when the guide's rotation text states an order or says "first", it must win.
C. Sentence plus bare duplicate (restoration druid, Disc, cat, Prot Paladin): skip the bare line when a sentence already names the ability.
D. Dev phrasing "The curated rotation" and "hence the reserve floor": block.

Files: /Users/jh/code/forever/.claude/worktrees/lane-discord-bot/data/curated/coach/*.md
