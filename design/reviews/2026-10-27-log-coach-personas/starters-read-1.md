# Task 24 - wow-player read of the 25 starter coach personas (2026-10-27)

Verdict: 9 SHIP (once the generator fixes land), 16 FIX. Every file also inherits the generator-level defects below.

## Generator-level fixes (tools/coach/persona_starter.py, voice.md)

G1. Consumables are per role, not per spec (19 dps files carry one identical line). Wrong for most:
 - "a weapon enchant (Elemental Sharpening Stone or Brilliant Wizard Oil); a combat potion (Mighty Rage Potion)" lists both oils to everyone. Casters (mage x3, warlock x3, shadow, balance, elemental) want Wizard Oil and a mana potion; Mighty Rage is a warrior rage potion and means nothing to them. Melee want the stone, not the oil.
 - "Grilled Squid" is an agility fish; no caster eats it (Nightfin Soup / Runn Tum Tuber for mana users).
 - "Flask of the Titans or Flask of Supreme Power": Titans is a health flask for tanks/melee, Supreme Power a caster flask. Offer the one that fits the spec.
 - Feral (cat and bear) cannot use stones or oils in form; Enhancement uses its Windfury/Rockbiter imbue (Forever: one weapon), not a stone; Elemental Sharpening Stone is edged-only (Prot Paladin on a mace, mace rogues).
 - Hunters (BM/MM) want a mana potion and a scope, not Mighty Rage.
 - Tank and healer lines offer "Flask of Chromatic Resistance", a resistance flask for specific fights, not a default (healers: Distilled Wisdom is right).
G2. DPS drill.resource says "Spend it on the list's buttons before it can cap". Mana does not cap; false for mage, warlock, priest, balance, elemental, hunters. Use a mana drill (Life Tap, wand, Evocation, potion timing) for mana specs.
G3. DPS drill.preparation says "weapon enchant and a combat potion" for specs with neither (see G1). Healer drill.core_casts "efficiency cooldown next to your most expensive heal ... as a pair" is not a thing for Resto Shaman or Resto Druid; give it only where the spec has one (Disc, Holy Paladin Divine Favor).
G4. Priority lines are first sentences cut from APL notes, so many are orphans with no button name or refer to sim internals. Quoted offenders (require each line to open with the button, or fall back to the guide's rotation prose):
 - "The +twenty percent Flame Shock bonus is therefore never collected." (Elemental, line 1)
 - "The instant Healing Wave." (Resto Shaman)
 - "Held until sixty energy so a dagger build, which would have cast Backstab at that point, never spends the energy here." (Combat)
 - "The client states the cost and damage." (Arms) / "The client states its cost, cast time, cooldown and damage." (Demo)
 - "Worth about four percent more threat for no survival cost." (Bear)
 - "Cast your current highest learned rank." (Arcane)
 - "Seal of Fury kept up." (Prot Paladin)
 - "Lower ranks of Heal were tried and did not clear the rule." (Disc) - dev note.
 - "The aura-is-known guard keeps the line off for a build without the talent, where the stack check alone would read as always true and Pyroblast would be spammed." (Fire)
 - "... the cited guide's own rotation step, not skipped." (BM)
 - "the preset brings" (Holy Paladin, Resto Shaman), "the major-cooldown pass above" (Resto Druid), "paced like Healing Touch", "same pace", "mana is no more than five points behind the fight clock", "Behind the mana pace, the filler waits for a lower tank and member health", "a bare cast ... falls through", "Era fallback", "idle by design", "a tank hit for twelve percent a swing".
G5. Number-stripping leftovers: bare "s" for seconds ("under three s remain", "fifteen s cooldown", "every two s", "one hundred eighty s fight") and "one one hundred-mana drop" (Elemental). Spell out "seconds"; fix the join.
G6. The mistakes list uses raw finding ids ("core_casts:", "adherence:"); a player does not know them. Use plain names.
G7. The APL's else-branch builder/filler is often absent: Arcane never says "Arcane Blast", Combat never says "Sinister Strike", Fire never names the Fireball filler, Arms never names Whirlwind/Heroic Strike, Shadow lists Mind Blast only inside the fallback. Add the filler as the last line.
G8. Shadow: "Mind Flay fills every remaining global" followed by wand "whenever nothing else is ready" contradict.

## Per file

SHIP (after G1-G6): druid-balance, druid-feral, hunter-marksmanship, mage-frost, paladin-retribution, priest-shadow (G8), rogue-assassination, warlock-affliction, warlock-destruction.
 - druid-feral: role dps but also covers Bear; bear lives in druid-feral-bear. Fine.
 - warlock-affliction/destruction: "Bane of Doom is the one curse kept up" acceptable as the sim's pick; verify against the guide.

FIX:
 - druid-feral-bear: "Worth about four percent more threat for no survival cost." has no subject. "Maul on the next swing whenever fifteen rage is on hand" - Maul costs 10 rage in Classic; check the Forever tooltip. No weapon enchant for a bear.
 - druid-restoration: "Regrowth is the most healing per mana in the kit" - a Classic player will not believe it (Regrowth is the expensive HoT). Reword or verify. "the major-cooldown pass above" and "same pace" unreadable. "Swiftmend ... as the last-ditch button under thirty percent" - not how Classic players use Swiftmend; say what the guide says. "Tranquility ... every two seconds for ten" verify.
 - hunter-beast-mastery: "the cited guide's own rotation step, not skipped" is meta. Hunter's Mark and Mend Pet absent. "Rapid Fire timed to land with an incoming auto shot just as Aimed Shot is about to come off cooldown" is hard to parse.
 - hunter-survival: priority wrong for a melee spec. It lists only "Counterattack: a bare cast that only lands while its parry-triggered proc aura is active", "A target that never attacks gives no dodge, so without Expose Prey the line stays idle." (that is Mongoose Bite, name stripped) and "It costs eighty mana and does fifty damage; the client states both." (that is the Wing Clip filler, name stripped; guide: "Wing Clip (80 mana for 50 damage)", only with a Windfury Totem). Raptor Strike (first in the guide, queued on swing), Mongoose Bite, Lacerate, Strider Kick, Hunter's Mark are missing as named lines (web/src/content/guides/hunter/survival.md rotation paragraph). Not shippable as is.
 - mage-arcane: no "Arcane Blast" line (G7); orphan "Cast your current highest learned rank."; Evocation/mana gem absent.
 - mage-fire: contradiction "Every filler Scorch also renews the Improved Scorch debuff, so the refresh line above rarely has anything to do."; aura-guard sentence; Fireball filler unnamed.
 - paladin-holy: "preset", "five points behind the fight clock", "tank hit for twelve percent a swing" are sim talk. "Surplus mana buys throughput" breaks the voice rule (healers talk mana); say "spend spare mana on Holy Light".
 - paladin-protection: "Seal of Fury kept up." orphan; "Lay on Hands ... fires itself below ten percent health" - tell the player to press it. Righteous Fury absent. Templar's Bulwark, Divine Shield "twelve seconds ... half damage dealt" are Forever-specific; confirm in the guide. Stone on a mace tank (G1).
 - priest-discipline: "Lower ranks of Heal were tried and did not clear the rule."; "Prayer of Healing heals the healer's whole party, so it is cast on a party member" garbled.
 - rogue-combat: dagger-build orphan; Sinister Strike builder missing (G7).
 - rogue-subtlety: "the debuff that lifts Rupture damage by fifteen percent for fifteen seconds" - I do not believe it (Classic Hemorrhage raises physical damage taken); check the Forever tooltip, as with the 145% Hemorrhage and 180% Ghostly Strike figures.
 - shaman-elemental: first line orphan; Flame Shock and Lava Burst never appear as buttons; "Mana Spring ... five-minute life ... returns about nine hundred mana" is a number the pack must prove; Wizard Oil/mana potion (G1); "one one hundred-mana drop".
 - shaman-enhancement: no Flame Shock or weapon imbue line though "Who they are" names the imbue; stone advice conflicts with the imbue (G1). Stormstrike "eight second cooldown" - verify (Classic 10 s).
 - shaman-restoration: "The instant Healing Wave." orphan; "the preset brings". Riptide-first claim Forever-specific; confirm guide.
 - warlock-demonology: "Immolate follows the two curses" when only Bane of Doom is listed; no pet line in a "pet spec"; "The client states..." orphan.
 - warrior-arms: "same as Fury" cross-spec reference; "The client states the cost and damage." orphan; Whirlwind and Heroic Strike absent (G7).

## Voice check
No healer file says "heal more"; healer drills talk mana and survival. No exclamation marks. Many priority lines break "Short sentences. Name the button" (G4).
