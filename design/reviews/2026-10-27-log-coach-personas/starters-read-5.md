# Task 24 - wow-player read 5 (final) of the 25 starter coach personas (2026-10-27, commit 36af4db34)

Verdict: 24 SHIP / 1 FIX. Priority lists re-read against the rendered files and the guide text (web/src/content/guides).

## Read-4 corrections: all landed
Cat leads with Shifting Power at sixty energy or lower; Wing Clip and Hamstring carry the Windfury Totem condition; Mongoose Bite says "only after you dodge, or while Expose Prey is up"; Bloodrage "below eighty"; Fire Scorch is the every-global filler and Improved Scorch is the rider; Holy Shock "only while your mana is keeping pace"; Prot Paladin follows the APL order with Consecration, Swift Judgement folded into Judgement, no bare Divine Protection; Ret has the Seal refresh, "Hammer of Wrath below twenty percent target health" and Consecration above sixty-five percent mana; Shadow is split into short sentences with Mind Blast after Mind Flay and Shoot added; Elemental is Mana Spring Totem, Lava Burst, Searing, Lightning Bolt, Chain Lightning, Magma; Enhancement has Windfury Weapon "before the pull" on one weapon and Earth Shock "only while mana is above half"; "The curated rotation" is gone from Holy Paladin and Discipline. No shaman carries a second weapon; Survival fights in melee; no mail on a shaman list.

## FIX (1)
hunter-survival, line: `- Counterattack.`
Replace with: `- Counterattack only after you parry: it is free and cannot be blocked, dodged or parried.`
Why: it is the same bare-conditional defect as the Mongoose Bite line already fixed. Counterattack is only usable in the parry window (survival.md 80, 85); a bare name tells a Survival hunter to press a greyed-out button. Shippable as a starter with that one change. Optional, not required: `- Lacerate.` to `- Lacerate whenever its bleed has fallen off.` (survival.md 85). Raptor Strike and Immolation Trap bare are fine (no condition).

## SHIP (24), notes only
- druid-balance, druid-feral-bear, druid-feral, druid-restoration, hunter-beast-mastery, hunter-marksmanship, mage-arcane, mage-fire, mage-frost, paladin-holy, paladin-protection, paladin-retribution, priest-discipline, priest-shadow, rogue-assassination, rogue-combat, rogue-subtlety, shaman-elemental, shaman-enhancement, shaman-restoration, warlock-affliction, warlock-demonology, warlock-destruction, warrior-arms.
- Prot Paladin "Righteous Fury is a setting, not a button" is the guide's own wording (protection.md 96); leave it.
- Not blockers, accuracy reviewer to tooltip-check: Stormstrike 8 s (guide says 8, Classic 10), Hemorrhage 145%/15% and Ghostly Strike 180% (subtlety.md 57), Flash of Light "most efficient heal per mana", Riptide, Disc Penance, Tranquility "when two are well down" for a channel.
- Clumsy but true (long sentences with no clean break): Arcane Blast cost line, Hemorrhage, Backstab, Affliction Immolate. Arcane's "Arcane Missiles: Whenever Missile Barrage's proc..." has a stray capital W. Enhancement lists the pre-pull Mana Spring Totem after Stormstrike; harmless because the line says "before the pull".
- priest-holy, warrior-fury, warrior-protection are not starters (full files); not read as starters.
