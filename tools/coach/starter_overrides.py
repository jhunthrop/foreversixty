"""Per-spec corrections to a generated priority list, each a short restatement of the spec's guide.

The generator reads the guide's rotation prose sentence by sentence. Where the guide states an order
in words ("goes down first", "sits ahead of", "comes last"), a condition the line must carry ("only
with a Windfury Totem in the group"), or a sentence too long to read at a glance, the reviewer's
correction is recorded here (wow-player read 4, guide-accuracy round 25). Every line restates its
guide: a test checks that each ability is named in the guide's rotation section and that no line
carries a digit.
"""

from __future__ import annotations

import re

from starter_text import short_sentences

Entry = tuple[str, str]

# ability -> line. A listed ability is added when the generator left it out. A line that names a
# spell, talent or rank the player gains during leveling says from when (round 25: the coach reads a
# boss pull at whatever level the log records, dungeon bosses from the low teens included).
GUIDE_LINES: dict[str, dict[str, str]] = {
    "druid-balance": {
        # balance.md: "Recast each one when it falls off"; Insect Swarm is the twenty-seventh point.
        "Moonfire": "Moonfire whenever it is not ticking on the target.",
        "Insect Swarm": "Insect Swarm from twenty-seven whenever it is not ticking on the target.",
        # balance.md: Moonkin Form is the fortieth point; "With no Eclipse charge banked, cast Wrath";
        # "If you fall behind, lean on Wrath"; Faerie Fire "only do it when no one else will".
        "Moonkin Form": "Moonkin Form from forty: shift into it before the pull and stay there.",
        "Wrath": (
            "Wrath whenever no Eclipse charge is banked, from thirty-five: one Wrath, then Starfires. "
            "Before twenty it is your filler, and when your mana falls behind the fight, lean on it."
        ),
        "Starfire": "Starfire everything else from twenty, spending the charges once you have Eclipse.",
        "Faerie Fire": "Faerie Fire only when no feral druid has it on the boss. In Moonkin Form it costs a global and mana.",
    },
    "druid-restoration": {
        # The card: Innervate below thirty percent mana; it trains at forty.
        "Innervate": "Innervate from forty when your mana drops below thirty percent.",
        # restoration.md (Swiftmend) and the card's gate: under three quarters health, mana keeping pace.
        "Swiftmend": "Swiftmend on the tank whenever he is under three quarters health and your mana is keeping pace.",
        # Gift of the Wild trains at fifty.
        "Gift of the Wild": "Gift of the Wild up on the raid from fifty, and Mark of the Wild on the group before then.",
    },
    "druid-feral": {
        # feral.md and the card: Faerie Fire is free in Cat Form; Rip at five points.
        "Faerie Fire": "Faerie Fire whenever the boss has none on it. In Cat Form it is free.",
        "Rip": "Rip at five combo points when it is off the boss or has under two seconds left.",
        # feral.md: Claw is the builder at twenty, Shred from twenty-one; Ravage trains at thirty-two;
        # Shifting Power is the fortieth point, pressed "at 60 energy or less, so all 40 fits under the cap".
        "Shred": "Shred fills every other global from twenty-one, building combo points. At twenty, Claw is the builder.",
        "Ravage": "Ravage from thirty-two: Prowl, walk behind the boss and open once the tank has two or three globals on it.",
        "Shifting Power": (
            "Shifting Power from forty, every time it comes off cooldown, but only at sixty energy or lower. "
            "Then the whole forty energy fits under the cap."
        ),
    },
    "druid-feral-bear": {
        # feral-bear.md: Bear Form until Dire Bear Form at forty; the numbered priority, Primal Bite
        # (forty) first, Lacerate (twenty-five), Faerie Fire, Demoralizing Roar, Maul, Swipe.
        "Dire Bear Form": "Dire Bear Form from forty, and Bear Form before then: be in it for the whole fight.",
        "Primal Bite": "Primal Bite on cooldown from forty, every six seconds.",
        "Lacerate": "Lacerate from twenty-five until it has five stacks, then again when it has under three seconds left.",
        "Faerie Fire": "Faerie Fire whenever the boss does not have it.",
        "Maul": (
            "Maul queued on your next swing whenever you have fifteen rage spare. "
            "From forty to fifty-three, wait for thirty rage, so Primal Bite seldom waits on rage."
        ),
        "Swipe": (
            "Swipe only with two or more targets in front of you, above forty rage. "
            "On a single boss, put the rage into Maul, and into Lacerate from twenty-five and Primal Bite from forty."
        ),
    },
    "hunter-beast-mastery": {
        # beast-mastery.md: mark first, then Aspect of the Hawk; "cast Aimed Shot every time the shared
        # six second cooldown comes up ... Multi-Shot is for packs and for when you have to move";
        # Feign Death before you reach the tank's threat.
        "Hunter's Mark": "Hunter's Mark first, before the pull, so the mana is back by the time the boss is pulled.",
        "Aspect of the Hawk": "Aspect of the Hawk goes up before the pull.",
        "Bestial Wrath": "Bestial Wrath from forty on the pull with Rapid Fire and your trinkets, then again every two minutes.",
        "Aimed Shot": (
            "Aimed Shot from twenty, every time the cooldown it shares with Multi-Shot comes up. "
            "Your auto shots keep firing while you cast it."
        ),
        "Multi-Shot": (
            "Multi-Shot on its cooldown at eighteen and nineteen. From twenty it is for packs and for moving, "
            "and never under a quarter of your mana."
        ),
        "Serpent Sting": "Serpent Sting when it is about to drop.",
        "Feign Death": "Feign Death before you reach the tank's threat, not after the boss turns. If it says Resisted, you still have all your threat.",
    },
    "hunter-marksmanship": {
        # marksmanship.md: as Beast Mastery; Sniper Shot is the raid build's talent at forty.
        "Hunter's Mark": "Hunter's Mark first, before the pull, so the mana is back by the time the boss is pulled.",
        "Aspect of the Hawk": "Aspect of the Hawk goes up before the pull.",
        "Aimed Shot": (
            "Aimed Shot from twenty, every time the cooldown it shares with Multi-Shot comes up. "
            "Your auto shots keep firing while you cast it."
        ),
        "Multi-Shot": (
            "Multi-Shot on its cooldown at eighteen and nineteen. From twenty it is for packs and for moving, "
            "and never under a quarter of your mana."
        ),
        "Feign Death": "Feign Death before you reach the tank's threat, not after the boss turns. If it says Resisted, you still have all your threat.",
    },
    "hunter-survival": {
        # survival.md: Aspect of the Beast trains at 30; Monkey until then.
        "Hunter's Mark": "Hunter's Mark on the boss first, so the mana is back by the pull. Wear Aspect of the Monkey until thirty, then Aspect of the Beast, and leave it up.",
        # survival.md: Feign Death "can miss or be resisted".
        "Feign Death": "Feign Death drops all of your threat if it lands. Check the message, because it can be resisted.",
        # survival.md: Counterattack after a parry costs mana (Resourcefulness cuts it, never to nothing);
        # Mongoose Bite after a dodge or from Expose Prey's window; Lacerate "is recast whenever its own
        # bleed has fallen off"; Wing Clip only with a Windfury Totem in the group; the damage racials named.
        "Counterattack": "Counterattack the moment it lights up after you parry: it cannot be blocked, dodged or parried.",
        "Lacerate": "Lacerate whenever its bleed has fallen off.",
        "Strider Kick": "Strider Kick on its own short cooldown, ahead of Immolation Trap.",
        "Mongoose Bite": (
            "Mongoose Bite the moment it lights up after a dodge. "
            "From thirty-five, Expose Prey also opens it off your hits on a marked target."
        ),
        "Rapid Fire": "Rapid Fire on the pull with your trinkets. Add Elune's Light, Blood Fury or Berserking if your race has one.",
        "Wing Clip": (
            "Wing Clip only with a Windfury Totem in the group, on a global with nothing else ready. "
            "Without the totem it loses damage."
        ),
    },
    "warlock-affliction": {
        # Bane of Doom trains at sixty; Siphon Life is the thirtieth point, Wrack the fortieth.
        "Bane of Doom": "Bane of Doom at sixty, while a minute or more of the fight is left.",
        "Siphon Life": "Siphon Life from thirty whenever it has fallen off.",
        "Wrack": "Wrack from forty with your DoTs up. Before forty, Shadow Bolt fills instead.",
        # affliction.md: Immolate restarted with about three seconds left; one curse per warlock, so the
        # raid leader's curse only when another warlock holds Curse of the Elements.
        "Immolate": (
            "Immolate while you level, with about three seconds left, so it lands a moment before the old one ends. "
            "The fresh hit is worth the tick you give up. The raid build sacrifices the Imp and drops it."
        ),
        "Life Tap": "Life Tap keeps you casting once mana runs low.",
        "Curse of Recklessness": (
            "Curse of Recklessness, Tongues or Weakness only if your raid leader asks. "
            "That is only when another warlock holds Curse of the Elements."
        ),
    },
    "warlock-demonology": {
        # Bane of Doom trains at sixty.
        "Bane of Doom": "Bane of Doom at sixty, the moment the last one fires, with a minute or more of the fight left.",
        # demonology.md: the five-step setup; "Your curse slot is then free" after another warlock's Elements.
        "Demonic Sacrifice": "Demonic Sacrifice only with the raid build's Demonic Pact, from forty. Sacrifice the Imp about a minute before the pull, then summon the Succubus. Without Pact, never sacrifice: keep the Succubus out.",
        "Soul Link": "Soul Link a couple of seconds before the pull, and again after every resummon.",
        "Curse of Recklessness": (
            "Curse of Recklessness, Tongues or Weakness only if your raid leader asks. "
            "That is only when another warlock holds Curse of the Elements."
        ),
    },
    "warlock-destruction": {
        # Bane of Doom trains at sixty; Shadowburn and Incinerate are talents (Incinerate at forty).
        "Bane of Doom": "Bane of Doom at sixty, the moment the last one fires, with a minute or more of the fight left.",
        "Shadowburn": "Shadowburn on cooldown once you have it, at any boss health.",
        "Incinerate": "Incinerate from forty for every other cast. Before forty, Shadow Bolt fills instead.",
        "Curse of Recklessness": (
            "Curse of Recklessness, Tongues or Weakness only if your raid leader asks. "
            "That is only when another warlock holds Curse of the Elements."
        ),
    },
    "warrior-arms": {
        # arms.md and the card notes: Bloodthrill (thirty-one raid, thirty-three leveling), Overpower from
        # twelve, Execute from twenty-four, Mortal Strike from forty, Recklessness from fifty, Reckless
        # Swings at fifty-seven in the raid build, Spearing Strike raid build only.
        "Bloodrage": "Bloodrage whenever Rage drops below eighty, since Rage limits everything else here.",
        "Rend": (
            "Rend kept ticking: the bleed ignores armor. "
            "Once Bloodthrill is in, your main-hand hits on a target bleeding from your Rend can light Overpower too."
        ),
        "Overpower": (
            "Overpower the moment it lights up, even ahead of Mortal Strike from forty. "
            "A dodge opens it for a few seconds, and a Bloodthrill proc does too once the talent is in."
        ),
        "Mortal Strike": "Mortal Strike on cooldown from forty, and from fifty-seven in the raid build Reckless Swings right behind it.",
        "Slam": "Slam only with Improved Slam, between Mortal Strikes from forty: the talent keeps your swing timer running through the cast.",
        "Spearing Strike": "Spearing Strike in the raid build, on any free global in Battle Stance. It stops at the execute swap.",
        "Heroic Strike": (
            "Heroic Strike only when you are about to cap at ninety or more. "
            "Below that, keep the rage for Overpower from twelve, Execute from twenty-four and Mortal Strike from forty."
        ),
        "Hamstring": (
            "Hamstring only with a Windfury Totem in the group, on a free global for ten rage. "
            "Never from twenty-five percent down: that rage is the first Execute."
        ),
        "Execute": (
            "Execute from twenty-four, at twenty percent, in Battle Stance first. "
            "It spends your whole bar, so that first Execute is your biggest hit of the fight."
        ),
        "Recklessness": (
            "Recklessness from fifty: once that first Execute has emptied the bar, swap to Berserker Stance and press it. "
            "Before fifty, stay in Battle Stance."
        ),
        "Sweeping Strikes": (
            "Sweeping Strikes only with two or more targets, from thirty-six. "
            "Press it in Battle Stance, then Berserker Stance and Whirlwind, and back for Overpower and Rend."
        ),
    },
    "mage-arcane": {
        # arcane.md: "Blast to 3 stacks ... Past 3, keep blasting only while your mana bar is fuller than the
        # boss's health bar and you are above 20% mana"; Missiles on Missile Barrage; the raid build's
        # Pyroblast on three Heating Up stacks and Arcane Wrath; Frostbolt the leveling filler.
        "Arcane Blast": (
            "Arcane Blast to three stacks. Past three, keep blasting only while your mana bar is fuller than the "
            "boss's health bar. Stop below a fifth of your mana."
        ),
        "Arcane Missiles": "Arcane Missiles the moment Missile Barrage lights up: the proc makes it free and halves the channel.",
        "Pyroblast": "Pyroblast with the raid build, on three Heating Up stacks, when it is a one and a half second cast.",
        "Arcane Wrath": "Arcane Wrath with the raid build, every time it is ready while you hold three Blast stacks.",
        "Frostbolt": "Frostbolt is the filler with the leveling build. In a raid, cast it only when you are too low on mana for a Fireball.",
    },
    "mage-fire": {
        # fire.md: "Everything else is Scorch" from thirty, with Fireball the filler below thirty.
        "Scorch": (
            "Scorch is the filler from thirty with the leveling build. Cast it on every global not spent on a three-stack Pyroblast, Fire Blast or Combustion. "
            "Every Scorch also renews your Improved Scorch stacks. Below thirty, Fireball is the filler."
        ),
    },
    "mage-frost": {
        # frost.md: Cone of Cold on a boss within ten yards, above half mana, on five Winter's Chill stacks.
        "Cone of Cold": (
            "Cone of Cold on a boss only within ten yards, above half mana, on five Winter's Chill stacks. "
            "At range, skip it."
        ),
    },
    "paladin-holy": {
        # holy.md: Oath of Light (twenty-five), Divine Favor (twenty-one), Holy Shock (thirty-two), Light's
        # Vigil (forty) are talents; "Flash of Light rank 5 on the tank below 90%. This is your staple".
        "Oath of Light": "Oath of Light from twenty-five, sworn on the tank a second or two before the pull.",
        "Divine Favor": "Divine Favor from twenty-one, then Holy Light, when the tank drops below forty percent.",
        "Holy Shock": "Holy Shock from thirty-two, only while your mana is keeping pace with the fight.",
        "Holy Light": (
            "Holy Light on the tank below seventy percent, and from thirty-two Holy Shock on anyone below seventy percent. "
            "Only with mana to spare: keep your mana no more than five points under the fight left."
        ),
        "Concentration Aura": "Concentration Aura from twenty-two if nobody else in your group runs it.",
        "Voice of Truth": (
            "Voice of Truth from twenty makes you immune to silences and interrupts for six seconds, every three minutes. "
            "Save it for the caster pack or boss that would lock you out."
        ),
        "Flash of Light": (
            "Flash of Light from twenty is your staple, on the tank below ninety percent. "
            "Everyone else gets it below eighty percent. When your mana falls behind the fight, heal with it alone until you catch up. "
            "At sixty in a raid that is rank five on everyone; without raid buffs, rank four on the party."
        ),
        "Light's Vigil": (
            "Light's Vigil from forty, only for a fight where your whole party takes big damage at once. "
            "On an ordinary boss its mana goes to the tank."
        ),
    },
    "paladin-protection": {
        # protection.md list: Seal of Fury, Holy Shield, Holy Strike, Hammer of the Righteous, Swift
        # Judgement, Judgement, Exorcism, Consecration. Holy Shield trains at thirty, Swift Judgement is the
        # twenty-ninth point, Hammer of the Righteous trains at forty, Greater Blessing of Kings at sixty.
        "Holy Shield": (
            "Holy Shield on cooldown from thirty, about two seconds before the boss reaches you and every time after. "
            "It is the first claim on your mana."
        ),
        "Hammer of the Righteous": (
            "Hammer of the Righteous on cooldown from forty, once Holy Shield's mana is safe. "
            "It hits your target and up to three enemies next to it."
        ),
        "Holy Strike": (
            "Holy Strike on cooldown, twenty mana at sixty. "
            "From forty Shieldsworn brings it back every six seconds, and each one keeps up your Iron Creed damage reduction."
        ),
        "Consecration": "Consecration from twenty, only above half mana, to spend mana you will not need.",
        "Exorcism": "Exorcism from twenty, on Undead and Demons.",
        "Templar's Bulwark": "Templar's Bulwark from twenty, first, at about a third of your health. It absorbs a full health bar for eight seconds.",
        "Judgement": (
            "Judgement to pull: under Seal of Fury it taunts, so the boss comes straight to you. "
            "After that, Judgement on cooldown, and from thirty only while Holy Shield's mana is safe."
        ),
        "Greater Blessing of Kings": (
            "Greater Blessing of Kings on yourself at sixty, Blessing of Kings from twenty before then. "
            "With a hunter's Heart of the Lion in the raid, take Might instead."
        ),
        "Swift Judgement": "Swift Judgement from twenty-nine whenever it is ready, unless a tank swap is coming: then hold it for the swap.",
        "Divine Shield": "Divine Shield from thirty-four and Divine Protection come after Templar's Bulwark, once Forbearance is gone.",
    },
    "paladin-retribution": {
        # retribution.md: with Unsealed Wrath, Seal of Command is the main Seal for the twist.
        "Seal of Command": (
            "Seal of Command once you have Unsealed Wrath, in a dungeon or a raid. "
            "Three swings store three Commands and the twist is under way. "
            "Twist only while your mana is at least four fifths of the fight left."
        ),
        # retribution.md: "Judgement on cooldown. It is off the global cooldown"; Seal of the Crusader up
        # before the pull, Seal of Righteousness straight after its Judgement; Exorcism and Consecration
        # train at twenty, Hammer of Wrath at forty-four, Holy Wrath at fifty.
        "Judgement": "Judgement on cooldown. It is off the global cooldown, so press it the moment it is ready.",
        "Seal of the Crusader": (
            "Seal of the Crusader about a second and a half before the pull, then Judge it on your first swing. "
            "Skip it if another paladin's Crusader Judgement is up."
        ),
        "Seal of Righteousness": (
            "Seal of Righteousness while leveling, at the top rank you have trained, straight after the opening Crusader Judgement. "
            "Recast it in its last second. With Unsealed Wrath it is your Seal whenever your mana falls behind the fight."
        ),
        "Exorcism": "Exorcism from twenty, on cooldown, against Undead and Demons only.",
        "Hammer of Wrath": "Hammer of Wrath from forty-four, below twenty percent target health, at its highest rank.",
        "Holy Strike": "Holy Strike on cooldown, every ten seconds. It spends the globals Seal upkeep leaves free.",
        "Consecration": (
            "Consecration from twenty and last, while the mana left after it still pays for your seal, Judgement "
            "and Holy Strike. Do not let your mana fall far behind the fight."
        ),
        "Holy Wrath": "Holy Wrath from fifty, only against a pack of Undead or Demons: on anything else it does nothing.",
    },
    "priest-discipline": {
        # Penance is the thirtieth point; Prayer of Healing trains at thirty.
        "Penance": "Penance from thirty on cooldown, on whoever is hurt: three heals one second apart every twelve seconds.",
        "Prayer of Healing": "Prayer of Healing from thirty, when three or more in a group are hurt.",
        # discipline.md: "Penance on cooldown"; "Shield the tank at the top rank when he is hurt, Weakened
        # Soul is off him and your mana is keeping pace"; Inner Focus (twenty-four) before Prayer of Healing
        # (thirty), Divine Aegis (thirty-six) and Greater Heal (forty); Dark Sacrifice for an Undead.
        "Power Word: Shield": (
            "Power Word: Shield on the tank at the top rank once he is hurt and Weakened Soul is off him. "
            "Your mana has to keep pace with the fight."
        ),
        "Inner Focus": (
            "Inner Focus on a Prayer of Healing from thirty or a Greater Heal from forty. "
            "It makes the spell free and adds twenty-five percent crit chance. "
            "A Greater Heal crit leaves the biggest Divine Aegis you can make."
        ),
        "Dark Sacrifice": "Dark Sacrifice if you are Undead, once your mana is down enough.",
        "Lightwell": "Lightwell from forty, dropped before the pull, and tell the raid it is up.",
    },
    "priest-shadow": {
        # shadow.md: "Keep Shadow Word: Pain up. Keep Devouring Plague up ... Mind Blast on cooldown. Shadow
        # Word: Death on cooldown. Mind Flay with everything else"; Vampiric Embrace (thirty-three leveling),
        # Improved Mind Flay (twenty-nine), Shadowform (forty-two) are talents.
        "Shadow Word: Pain": "Shadow Word: Pain kept up, unless the fight is close to ending.",
        "Inner Focus": "Inner Focus on Mind Blast every time it is up: the Mind Blast is free and more likely to crit.",
        "Mind Blast": "Mind Blast on cooldown, ahead of Mind Flay.",
        "Shadow Word: Death": (
            "Shadow Word: Death on cooldown, ahead of Mind Flay. "
            "Once you have Vampiric Embrace, keep it up so your own damage heals back the backlash."
        ),
        "Mind Flay": (
            "Mind Flay with everything else, and never clip it: a clipped channel loses the rest. "
            "Once you have Improved Mind Flay, the third tick is the biggest."
        ),
        "Shadowform": (
            "Shadowform from forty-two: buff the raid, then shift and drink back to full. "
            "The form costs mana, so never shift on the pull."
        ),
        "Shoot": "Shoot with your wand only when nothing else can be cast, such as out of mana.",
        "Dark Sacrifice": "Dark Sacrifice if you are Undead, on a long fight once your mana is well down.",
    },
    "rogue-assassination": {
        # assassination.md: Ambush lands into unshredded armor; Mutilate is the fortieth point.
        "Ambush": "Ambush hits harder on paper, but it lands into unshredded armor and gets none of Lethality's bonus.",
        "Mutilate": "Mutilate from forty is the builder: both daggers strike, two combo points a cast, harder against a poisoned target.",
        # assassination.md: Cold Blood (thirty-one raid, thirty-three leveling); poisons swapped at sixty;
        # Kidney Shot rank two at fifty.
        "Cold Blood": "Cold Blood from thirty-three, at four or more combo points, right before the Eviscerate.",
        "Deadly Poison": (
            "Deadly Poison on the main hand and Instant on the off hand from thirty to fifty-nine. "
            "At sixty swap them: the off hand then holds Deadly at five stacks on its own."
        ),
        "Kidney Shot": (
            "Kidney Shot costs twenty-five energy, sits on a twenty second cooldown and stuns a second longer per "
            "combo point. Five seconds at five, six once rank two trains at fifty. Use it on casters and on "
            "whatever is hitting your healer."
        ),
    },
    "rogue-combat": {
        # combat.md: Flash of Steel full only on the raid build at sixty; Adrenaline Rush (forty), Blade
        # Flurry (thirty-three); Kidney Shot rank two at fifty.
        "Shank": (
            "Shank on cooldown, right behind any damage racial and your trinket. "
            "It is an off-hand stab for five energy that cannot miss and pays back more than it costs."
        ),
        "Adrenaline Rush": "Adrenaline Rush from forty on cooldown, with your damage racial if your race has one, a trinket and a potion.",
        "Blade Flurry": "Blade Flurry from thirty-three, at any energy, for its twenty-five.",
        "Kidney Shot": (
            "Kidney Shot costs twenty-five energy, sits on a twenty second cooldown and stuns a second longer per "
            "combo point. Five seconds at five, six once rank two trains at fifty."
        ),
    },
    "rogue-subtlety": {
        # subtlety.md: Garrote after every Ghostly Strike (twenty-six); Rupture at four or more whenever it
        # is not ticking; Death Mark (forty) only while Rupture ticks; Backstab until Hemorrhage (thirty);
        # Premeditation again in the fight; Initiative full only from thirty-two.
        "Garrote": (
            "Garrote opens from Stealth once the tank has the boss, and a bleed does not care about missing Sunders. "
            "From twenty-six it follows every Ghostly Strike at four points or fewer: the strike lets it in without Stealth."
        ),
        "Ghostly Strike": "Ghostly Strike from twenty-six on cooldown, never at five points.",
        "Rupture": "Rupture at four or more combo points whenever it is not ticking: it is the bleed every mark feeds on.",
        "Death Mark": (
            "Death Mark from forty, at four or more combo points while Rupture ticks, never into a lapsed bleed. "
            "It comes back every ten seconds, so a mark now beats a bigger one later."
        ),
        "Slice and Dice": "Slice and Dice at three or more combo points while Rupture ticks, when it is down or nearly down.",
        "Cheap Shot": "Cheap Shot opens on trash that will die fast, three combo points once Initiative is full, then Eviscerate.",
        "Backstab": (
            "Backstab from behind with a dagger in your main hand is the builder until Hemorrhage at thirty. "
            "Sinister Strike only when you cannot get behind or hold no dagger."
        ),
        "Hemorrhage": (
            "Hemorrhage from thirty builds. "
            "It waits a moment when Ghostly Strike is about to come back, unless you are close to full energy."
        ),
        "Premeditation": (
            "Premeditation a couple of seconds before the pull for two combo points. "
            "Press it again every two minutes in the fight whenever you have three or fewer. Spend them inside twenty seconds."
        ),
        "Kidney Shot": (
            "Kidney Shot costs twenty-five energy, sits on a twenty second cooldown and stuns a second longer per "
            "combo point. Five seconds at five, six once rank two trains at fifty."
        ),
    },
    "shaman-elemental": {
        # elemental.md and the card: Mana Spring (twenty-six) a few seconds before the pull, skipped when a
        # Restoration shaman has one down; Lightning Mastery (forty) with the cooldowns; Flame Shock ahead of
        # everything but the cooldowns; Searing until Magma (twenty-six); Chain Lightning (thirty-two);
        # Fire Nova a pack spell; Lightning Bolt with a cheaper rank when mana falls behind.
        "Mana Spring Totem": (
            "Mana Spring Totem from twenty-six, a few seconds before the pull and whenever your Water slot is empty. "
            "Skip it if a Restoration shaman already has one down for your group."
        ),
        "Lightning Mastery": "Lightning Mastery from forty with your trinkets and any damage racial on the pull, and again every three minutes.",
        "Flame Shock": "Flame Shock the moment it falls off, ahead of everything but your cooldowns.",
        "Searing Totem": (
            "Searing Totem on a single target, and on packs until Magma Totem at twenty-six. "
            "Its bolts reach twenty yards, so drop it within twenty yards of the boss."
        ),
        "Magma Totem": "Magma Totem from twenty-six on two or more targets, in place of Searing Totem.",
        "Chain Lightning": (
            "Chain Lightning from thirty-two on two or more targets. "
            "Its first hit beats a Lightning Bolt, and each jump deals seventy percent of the hit before it."
        ),
        "Fire Nova": (
            "Fire Nova from twelve on two or more targets, bursting from the Fire totem you dropped in the pack. "
            "On a single boss, a Clearcasting proc goes to your next Lightning Bolt."
        ),
        "Lightning Bolt": (
            "Lightning Bolt on every global the lines above leave free. "
            "Cast your top rank while your mana is ahead of the fight, a cheaper one when it falls behind."
        ),
    },
    "shaman-enhancement": {
        # enhancement.md and the card: Clearcasting comes from Elemental Focus in the raid build.
        "Fire Nova": (
            "Fire Nova on a Clearcasting proc, with the raid build. "
            "It is free then, and in melee your Searing Totem stands within its ten yards of the boss."
        ),
        # enhancement.md: Windfury Weapon from thirty, Rockbiter before it and in a raid where a Windfury
        # Totem reaches you; Windfury Totem (thirty-two); Stormstrike (twenty-five); Rage of the Farseer
        # (forty); Totemic Call (twenty-two); Earth Shock above a quarter of your mana.
        "Windfury Weapon": (
            "Windfury Weapon on your weapon from thirty, before the pull and whenever it falls off. "
            "Before thirty, and in a raid where a Windfury Totem reaches you, use Rockbiter Weapon."
        ),
        "Windfury Totem": "Windfury Totem from thirty-two in your Air slot when no other shaman has it down near the melee.",
        "Totemic Call": "Totemic Call from twenty-two when the boss moves away from your totems: follow it, and it brings them to you.",
        "Rage of the Farseer": "Rage of the Farseer from forty, with your trinkets and any damage racial, on the pull and again on cooldown.",
        "Stormstrike": (
            "Stormstrike on cooldown from twenty-five. "
            "It makes your next Lightning Bolt, Chain Lightning or Earth Shock on the target hit twenty percent harder."
        ),
        "Earth Shock": "Earth Shock last, on cooldown while your mana is above a quarter. Keep that last quarter for Stormstrike and your totems.",
    },
    "shaman-restoration": {
        # restoration.md and the card: Water Shield and Healing Stream at twenty, Mana Tide (twenty-five)
        # under a quarter of your mana, Riptide (fifty), Chain Heal (forty), Healing Wave at sixty percent
        # never below rank eight at sixty, Lesser Healing Wave (twenty), Nature's Swiftness (forty-four).
        "Water Shield": (
            "Water Shield from twenty whenever it is down. "
            "Each globe returns two percent of your maximum mana when you are hit or a heal crits."
        ),
        "Healing Stream Totem": "Healing Stream Totem from twenty whenever your Water slot is empty: it heals your party every two seconds.",
        "Mana Tide Totem": (
            "Mana Tide Totem from twenty-five when your mana drops below a quarter. "
            "It takes the Water slot, so put Healing Stream back when it ends."
        ),
        "Riptide": (
            "Riptide from fifty on the tank when it is hurt and your last Riptide has run out. "
            "An early recast loses its last tick."
        ),
        "Chain Heal": (
            "Chain Heal from forty on the tank as soon as anyone near it is hurt. Your mana has to be ahead of the fight. "
            "On a target with your Riptide it heals twenty-five percent more."
        ),
        "Healing Wave": (
            "Healing Wave on the tank when it drops to sixty percent. "
            "When your mana falls behind, drop a rank, but never below rank eight at sixty."
        ),
        "Lesser Healing Wave": "Lesser Healing Wave from twenty when the tank is down to thirty percent and the next hit could kill it.",
        "Nature's Swiftness": (
            "Nature's Swiftness from forty-four when the tank drops below twenty-five percent, into your highest rank "
            "of Healing Wave."
        ),
    },
}

# Round 26: hunter buttons named with the level they arrive (the generator's own rule).
GUIDE_LINES["hunter-beast-mastery"].update(
    {
        "Rapid Fire": (
            "Rapid Fire from twenty-six on the pull, as an auto shot leaves. "
            "Press it as the shared cooldown is about to come back, and from forty with Bestial Wrath."
        ),
        "Feign Death": (
            "Feign Death from thirty, before you reach the tank's threat, not after the boss turns. "
            "If it says Resisted, you still have all your threat."
        ),
        "Volley": "Volley from forty is a mana-heavy channel, worth it only on a big pack that lives through it.",
        "Freezing Trap": "Freezing Trap from twenty is your crowd control, laid before the pull. Any damage breaks it.",
        "Explosive Trap": (
            "Explosive Trap from thirty-four under a pack the tank is holding hits everything within ten yards. "
            "At sixty it costs more than Volley, so save it for the same big packs."
        ),
    }
)
GUIDE_LINES["hunter-marksmanship"].update(
    {
        "Rapid Fire": (
            "Rapid Fire from twenty-six on the pull, as an auto shot leaves. "
            "Press it as the shared cooldown is about to come back."
        ),
        "Sniper Shot": "Sniper Shot with the raid build from forty, whenever it is ready and the shared cooldown is not.",
        "Feign Death": (
            "Feign Death from thirty, before you reach the tank's threat, not after the boss turns. "
            "If it says Resisted, you still have all your threat."
        ),
        "Volley": "Volley from forty is a mana-heavy channel, worth it only on a big pack that lives through it.",
        "Freezing Trap": "Freezing Trap from twenty is your crowd control, laid before the pull.",
        "Explosive Trap": (
            "Explosive Trap from thirty-four under a pack the tank is holding hits everything within ten yards. "
            "At sixty it costs more than Volley, so save it for the same big packs."
        ),
    }
)
GUIDE_LINES["hunter-survival"].update(
    {
        "Counterattack": (
            "Counterattack from thirty-two the moment it lights up after you parry. It cannot be blocked, dodged or parried."
        ),
        "Lacerate": "Lacerate from thirty whenever its bleed has fallen off.",
        "Strider Kick": "Strider Kick from thirty on its own short cooldown, ahead of Immolation Trap.",
        "Rapid Fire": (
            "Rapid Fire from twenty-six on the pull with your trinkets. "
            "Add Elune's Light, Blood Fury or Berserking if your race has one."
        ),
        "Disengage": (
            "Disengage from twenty trims some threat on a five second cooldown. "
            "It turns off your auto attack, so start swinging again the moment it lands."
        ),
        "Feign Death": (
            "Feign Death from thirty drops all of your threat if it lands. Check the message, because it can be resisted."
        ),
        "Deterrence": "Deterrence from twenty-five is your emergency button when something turns on you.",
        "Explosive Trap": "Explosive Trap from thirty-four under the pack is your AoE, and Engineering's explosives are the rest.",
        "Freezing Trap": "Freezing Trap from twenty is your crowd control, laid before the pull.",
    }
)


# The guide's own order statements, which win over the order of first mention. Listed abilities lead.
GUIDE_ORDER: dict[str, tuple[str, ...]] = {
    # balance.md: the form, the two DoTs, then the Eclipse loop of Wrath and Starfire.
    "druid-balance": ("Moonkin Form", "Moonfire", "Insect Swarm", "Wrath", "Starfire", "Innervate"),
    # feral.md and the card: the Ravage opener and Faerie Fire, Shifting Power, Rip, the Bite while Rip
    # still has five seconds or more, then "Shred is everything else".
    "druid-feral": (
        "Ravage", "Faerie Fire", "Shifting Power", "Rip", "Ferocious Bite", "Shred", "Claw", "Cower",
    ),
    # feral-bear.md: the form and Enrage before the pull, then the numbered priority list.
    "druid-feral-bear": (
        "Dire Bear Form", "Enrage", "Primal Bite", "Lacerate", "Faerie Fire", "Demoralizing Roar", "Maul", "Swipe",
    ),
    # beast-mastery.md: the mark, then the aspect, the burst, then the shared cooldown and the fillers.
    "hunter-beast-mastery": (
        "Hunter's Mark", "Aspect of the Hawk", "Bestial Wrath", "Rapid Fire", "Aimed Shot", "Multi-Shot",
        "Serpent Sting", "Arcane Shot",
    ),
    "hunter-marksmanship": (
        "Hunter's Mark", "Aspect of the Hawk", "Rapid Fire", "Aimed Shot", "Sniper Shot", "Multi-Shot",
        "Serpent Sting", "Arcane Shot",
    ),
    # survival.md: "Raptor Strike, Counterattack, Mongoose Bite, Lacerate, Strider Kick, Immolation Trap".
    "hunter-survival": (
        "Hunter's Mark", "Raptor Strike", "Counterattack", "Mongoose Bite", "Lacerate", "Strider Kick",
        "Immolation Trap", "Wing Clip",
    ),
    # arcane.md, the raid loop.
    "mage-arcane": (
        "Arcane Blast", "Arcane Missiles", "Pyroblast", "Fire Blast", "Arcane Wrath", "Fireball", "Frostbolt",
        "Arcane Power", "Presence of Mind",
    ),
    # protection.md list items in order.
    "paladin-protection": (
        "Righteous Fury", "Seal of Fury", "Holy Shield", "Holy Strike", "Hammer of the Righteous", "Swift Judgement",
        "Judgement", "Exorcism", "Consecration",
    ),
    # retribution.md: Judgement, Exorcism, Hammer of Wrath, Seal, Holy Strike, Consecration.
    "paladin-retribution": (
        "Blessing of Might", "Seal of the Crusader", "Judgement", "Exorcism", "Hammer of Wrath",
        "Seal of Righteousness", "Seal of Command", "Holy Strike", "Consecration",
    ),
    # discipline.md: Penance on cooldown and the shield first after the pull, then the fillers.
    "priest-discipline": (
        "Lightwell", "Fear Ward", "Penance", "Power Word: Shield", "Renew", "Flash Heal", "Prayer of Healing",
        "Greater Heal", "Heal", "Lesser Heal", "Inner Focus", "Dark Sacrifice",
    ),
    # shadow.md: the form before the pull; "Keep Shadow Word: Pain up. Keep Devouring Plague up ... Mind
    # Blast on cooldown. Shadow Word: Death on cooldown. Mind Flay with everything else"; the card puts
    # Vampiric Embrace ahead of Devouring Plague.
    "priest-shadow": (
        "Shadowform", "Shadow Word: Pain", "Vampiric Embrace", "Devouring Plague", "Inner Focus", "Mind Blast",
        "Shadow Word: Death", "Mind Flay", "Shoot",
    ),
    # subtlety.md and the card: Garrote and Ghostly Strike, Rupture, Death Mark, Slice and Dice.
    "rogue-subtlety": (
        "Premeditation", "Garrote", "Ghostly Strike", "Rupture", "Death Mark", "Slice and Dice", "Eviscerate", "Shank",
        "Backstab", "Hemorrhage",
    ),
    # elemental.md and the card: Mana Spring, the cooldowns, Flame Shock, the totems, the pack spells,
    # Lightning Bolt for the rest.
    "shaman-elemental": (
        "Mana Spring Totem", "Lightning Mastery", "Flame Shock", "Searing Totem", "Magma Totem", "Chain Lightning",
        "Fire Nova", "Lightning Bolt", "Earth Shock",
    ),
    # enhancement.md: the imbue is "applied before the pull", ahead of every rotation line.
    "shaman-enhancement": ("Windfury Weapon",),
    # restoration.md and the card.
    "shaman-restoration": (
        "Water Shield", "Healing Stream Totem", "Mana Tide Totem", "Riptide", "Chain Heal", "Healing Wave",
        "Lesser Healing Wave", "Nature's Swiftness",
    ),
    # affliction.md: "Then Corruption, Bane of Doom, Siphon Life and Immolate, and Wrack".
    "warlock-affliction": (
        "Demon Armor", "Curse of the Elements", "Corruption", "Bane of Doom", "Amplify Curse", "Bane of Agony",
        "Siphon Life", "Immolate", "Wrack", "Life Tap",
    ),
    # demonology.md: the setup, then "Curse of the Elements ..., then Corruption, then Bane of Doom, then
    # Immolate, then Shadow Bolt".
    "warlock-demonology": (
        "Demonic Sacrifice", "Demon Armor", "Soul Link", "Curse of the Elements", "Corruption", "Bane of Doom",
        "Bane of Agony", "Immolate", "Shadow Bolt", "Life Tap",
    ),
}

# Buttons the guide tells this spec not to press, in words the generator cannot read as a rule-out
# (a pronoun sentence, "costs more than it is worth"), or that sit outside the fight. Never listed.
EXCLUDED: dict[str, frozenset[str]] = {
    # beast-mastery.md: Mend Pet is between pulls; Aspect of the Falcon is named only to say Hawk wins.
    "hunter-beast-mastery": frozenset({"Mend Pet", "Aspect of the Falcon"}),
    "hunter-marksmanship": frozenset({"Aspect of the Falcon"}),
    # feral.md: Pounce is for trash; on a boss, Ravage.
    "druid-feral": frozenset({"Pounce"}),
    # holy.md: the paladin's own blessing goes on the raid, never a rotation button.
    "paladin-holy": frozenset({"Blessing of Wisdom"}),
    # survival.md: "From 30 to 59 Beast is the only melee aspect you have".
    "hunter-survival": frozenset({"Aspect of the Falcon"}),
    # arcane.md: "the build has no spare point for Ice Block".
    "mage-arcane": frozenset({"Ice Block"}),
    # enhancement.md: "Fire Nova ... costs more mana than its damage is worth on one target, and so does Flame Shock".
    "shaman-enhancement": frozenset({"Flame Shock"}),
    # The raid leader's curse is one line with its condition (Curse of Recklessness).
    "warlock-affliction": frozenset({"Curse of Tongues", "Curse of Weakness"}),
    # demonology.md: Soul Fire "Skip it anyway ... never cast it in a raid"; Searing Pain "Keep it off your bars in raids".
    "warlock-demonology": frozenset({"Soul Fire", "Searing Pain", "Curse of Tongues", "Curse of Weakness"}),
    # destruction.md: Searing Pain off the raid bars; Hellfire burns you.
    "warlock-destruction": frozenset({"Searing Pain", "Hellfire", "Curse of Tongues", "Curse of Weakness"}),
}

# The short names the guides use for a button, by full name. The generator reads the short name
# as the button wherever the full name is not written ("Missiles" for Arcane Missiles).
SHORT_NAMES: dict[str, str] = {
    "Arcane Missiles": "Missiles",
    "Mana Tide Totem": "Mana Tide",
    "Power Word: Shield": "Shield",
}

# code -> drill, for a spec whose guide plays a drill differently from its role's shared text.
DRILL_LINES: dict[str, dict[str, str]] = {
    # restoration.md: "keep the top rank on the tank and put rank 7 on a hurt party member"; "If your
    # mana is behind, let the HoTs and Swiftmend carry it".
    "druid-restoration": {
        "ranks": (
            "In a raid, keep the top rank on the tank and bind Rejuvenation rank seven for a hurt party member. "
            "While leveling, use the ranks your level calls for."
        ),
        "resource": (
            "Before the pull, note the mana you should have at the halfway mark. If you are behind it, let the HoTs "
            "and Swiftmend carry the tank and keep Nature's Swiftness for the save."
        ),
    },
    # holy.md: Flash of Light rank five on the tank, rank four on everyone else, rank one a trap at sixty;
    # "if it is behind, heal with Flash of Light alone".
    "paladin-holy": {
        "ranks": (
            "Bind the ranks you use: at sixty, Flash of Light rank five on the tank, and in a raid on everyone else too; "
            "without raid buffs, rank four on the party. Holy Light at its top rank. Rank one is a trap at sixty."
        ),
        "resource": (
            "Before the pull, note the mana you should have at the halfway mark. If you are behind it, heal with "
            "Flash of Light alone until you catch up."
        ),
    },
    # protection.md: "Templar's Bulwark first, at about a third of your health"; Forbearance.
    "paladin-protection": {
        "death": (
            "From twenty, put Templar's Bulwark on a key you can hit without looking, and press it on a dummy pull when your health "
            "falls to about a third. Divine Shield and Divine Protection wait for Forbearance to fade."
        ),
    },
    # discipline.md: Greater Heal top rank on the tank, rank one on everyone else; "When your bar falls
    # behind the fight, stop topping off the party and heal only real holes."
    "priest-discipline": {
        "ranks": (
            "Before forty, bind the Heal rank your level calls for. From forty, bind Greater Heal rank one: "
            "it is the party's small-hole heal, and the tank's filler too until you learn rank four at fifty-eight. "
            "From then the tank gets the top rank."
        ),
        "resource": (
            "Before the pull, note your halfway mana. If you are behind it, keep Penance on cooldown, "
            "hold the shield until you are back on pace, let the filler wait and save Flash Heal for the emergency."
        ),
    },
    # restoration.md: Mana Tide under a quarter of your mana; Healing Wave never below rank eight at sixty.
    "shaman-restoration": {
        "core_casts": (
            "When your mana drops below a quarter, drop Mana Tide Totem, then put Healing Stream back when it ends. "
            "Use mana potions and runes as soon as each fits under your cap."
        ),
        "ranks": (
            "Bind Healing Wave at the rank your level calls for, and at sixty never below rank eight. "
            "Keep your top rank for Nature's Swiftness."
        ),
        "resource": (
            "Before the pull, note the mana you should have at the halfway mark. If you are behind it, drop Healing "
            "Wave a rank and let Riptide and Healing Stream carry the small holes."
        ),
    },
}


def with_lines(slug: str, entries: list[Entry]) -> list[Entry]:
    lines = GUIDE_LINES.get(slug, {})
    present = {name for name, _ in entries}
    replaced = [(name, lines.get(name, line)) for name, line in entries]
    return replaced + [(name, line) for name, line in lines.items() if name not in present]


def in_guide_order(slug: str, entries: list[Entry]) -> list[Entry]:
    leading = GUIDE_ORDER.get(slug, ())
    rank = {name: index for index, name in enumerate(leading)}
    return sorted(entries, key=lambda entry: rank.get(entry[0], len(leading)))


def without_repeated_bare(entries: list[Entry]) -> list[Entry]:
    """A bare name another line already names is a repeat (read 4, pattern C)."""
    lines = [line for _, line in entries]
    return [
        (name, line)
        for name, line in entries
        if line != f"{name}." or not any(re.search(rf"\b{re.escape(name)}\b", o) for o in lines if o != line)
    ]


def without_excluded(slug: str, entries: list[Entry]) -> list[Entry]:
    excluded = EXCLUDED.get(slug, frozenset())
    return [(name, line) for name, line in entries if name not in excluded]


def corrected(slug: str, entries: list[Entry]) -> list[Entry]:
    ordered = without_repeated_bare(in_guide_order(slug, with_lines(slug, without_excluded(slug, entries))))
    return [(name, short_sentences(line, name)) for name, line in ordered]
