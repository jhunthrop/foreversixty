# Talent search 2: classifier fix + full re-run

Branch `talent-search-2`, engine 90f9325b0, client 1.60.1.70009. `sim/cmd/talent-search` classified
every talent by a static source scan (`modeled.go`) that only matched a literal `Talents.<GoName>`
receiver. Every package that aliases the proto - mage's `applyDeclarativeTalents` does `t :=
mage.Talents` and reads `t.FirePower` - was invisible to that scan, so the aliased fields were
mislabeled "unmodeled" even when their own probe row showed a real DPS swing (Fire Power: +33.1 ±
8.0 DPS in the old mage-fire report). That miscount fed the "keeps every unmodeled guide talent"
winner and the credit estimate, so the old reports' adoption recommendations were built on a bad
classifier.

## Fix

`sim/cmd/talent-search/modeled.go`:

- `talentRefRE` now matches `\b\w+\.(?:Get)?([A-Z][A-Za-z0-9]*)\b` - any identifier before the dot,
  not only `Talents` - so an aliased receiver (`t := mage.Talents; t.FirePower`) counts exactly as
  `mage.Talents.FirePower` would. Comment-stripping and the test/generated-file skip are unchanged.
- A new `classify(staticModeled bool, c credit) talentClass` applies the precedence rule: a probe
  that moves DPS beyond the combined error (`|Diff| > Err`) is `damage` whatever the static scan
  found; short of that, the scan finding any read is at worst `modeled, no damage`; `unmodeled` only
  when the scan finds nothing **and** the probe is within error.
- `report.go`'s `engineStatus`, `removedUnmodeled`, `unmodeledIn` and the "keeps every unmodeled
  guide talent" winner (`bestClean`) all now go through `classify` instead of the raw static-scan
  map, so the fix reaches every place the old miscount leaked into: the credits table label, which
  guide talents a candidate is allowed to drop "for free", and the winner picked under that rule.
- Every report now prints an **Engine gaps** section: the talents the *final* classification still
  calls unmodeled that sit in the spec's own tree or that the guide takes - the actual engine-lane
  work list for that spec, as opposed to every utility talent in another role's tree.

Unit tests in `sim/cmd/talent-search/modeled_test.go`: a fixture source with `t := x.Talents` +
`t.FirePower` counts as modeled; the comment/test/generated skip still holds under the new regex;
`classify`'s precedence table covers scan-unmodeled-but-beyond-error (damage wins), scan-modeled-but-
beyond-error (still damage), scan-modeled-within-error (modeled, no damage), scan-unmodeled-within-
error (unmodeled), a negative diff beyond error, and the exact-boundary case (not beyond error).

```
cd sim && go vet ./cmd/talent-search/... && go test ./cmd/talent-search/...
ok  	github.com/jhunthrop/foreversixty/sim/cmd/talent-search	0.295s
```

Confirmed against the real engine: mage-fire's Fire Power now reports `+33.8 ± 7.7` and `damage`
(was `unmodeled`); the spec's Engine gaps list drops to the four Fire-tree talents that are actually
never credited (Flame Throwing, Impact, Improved Fire Ward, Hot Streak).

## Re-run

All 20 written-rotation DPS specs re-run with the tool's existing defaults (300 screening
iterations, top 10, 800 confirmation iterations - `100 * api.WeightsIterationsFactor`, seed 7),
overwriting `design/reviews/talent-search/*.md`. Each spec already fans its own sims out to 14
concurrent workers, so specs were run back-to-back rather than in parallel processes (14 cores,
already saturated per spec); total time was under 3 minutes for all 20.

## Table 1: per-spec results

"Keep-winner" is the best finalist that drops no talent the final classification calls unmodeled -
the recommendation the engine's own numbers can actually back (`report.go`'s `bestClean`). Gain % and
Beyond error are the keep-winner against the guide. "Own tree still max" asks whether the spec's own
tree still holds the most points in the keep-winner build. "Engine gaps in own tree" is this spec's
own Engine gaps section filtered to its own tree (full list, including off-role-tree gaps, is in each
report).

| Spec | Guide DPS | Best DPS (variant) | Best code | Keep-winner DPS | Keep-winner code | Gain % | Beyond error | Keep-winner trees (own tree) | Own tree still max | Engine gaps in own tree |
|---|---|---|---|---|---|---|---|---|---|---|
| Balance (druid-balance) | 157.7 | 179.1 (guide, all non-damage points re-spent) | `FS1:1.60.1.70009:druid:night-elf:523221101550105/05/550003:` | 161.8 | `FS1:1.60.1.70009:druid:night-elf:523221101540105/0/55333:` | +2.6% | yes | 32/0/19 (Balance) | yes | Improved Wrath, Genesis, Improved Moonfire, Nature's Majesty, Nature's Reach, Improved Entangling Roots, Nature's Splendor, Overgrowth, Nature's Grace, Eclipse |
| Feral (druid-feral) | 189.4 | 190.8 (guide, modeled non-damage points re-spent) | `FS1:1.60.1.70009:druid:night-elf:0/552322212103201/55532:` | 190.8 | `FS1:1.60.1.70009:druid:night-elf:0/552322212103201/55532:` | +0.7% | yes | 0/31/20 (Feral Combat) | yes | Feral Swiftness, Feral Instinct, Brutal Impact, Thick Hide, Feral Charge, Shredding Attacks, Primal Bite, Blood Frenzy, Predatory Instincts, King of the Jungle, Natural Reaction, Rend and Tear |
| Beast Mastery (hunter-beast-mastery) | 200.4 | 204.7 (guide, all non-damage points re-spent) | `FS1:1.60.1.70009:hunter:dwarf:5420001515121251/0/5500020301:` | 200.4 | `FS1:1.60.1.70009:hunter:dwarf:5420001505001251/3551/51:` | +0.0% | no | 31/14/6 (Beast Mastery) | yes | Deadly Aspects, Endurance Training, Focused Fire, Improved Aspect of the Monkey, Pathfinding, Improved Revive Pet, Bestial Swiftness, Improved Mend Pet, Summon Hawk, Spirit Bond, Intimidation |
| Marksmanship (hunter-marksmanship) | 182.0 | 204.2 (guide, all non-damage points re-spent + 1 Sniper Shot -> Bestial Discipline) | `FS1:1.60.1.70009:hunter:dwarf:55000005050001/1310550001510305/0:` | 182.0 | `FS1:1.60.1.70009:hunter:dwarf:5522/35305500115003/51:` | +0.0% | no | 14/31/6 (Marksmanship) | yes | Hawk Eye, Improved Concussive Shot, Lethal Attacks, Improved Stings, Careful Aim, Rapid Killing, Improved Arcane Shot, Lone Wolf, Trueshot Aura, Rapid Recuperation, Scatter Shot, Sniper Shot |
| Survival (hunter-survival) | 238.1 | 248.1 (deep Beast Mastery + 1 Deadly Aspects -> Strider Kick) | `FS1:1.60.1.70009:hunter:dwarf:450022050500025/0/5002302300500001:` | 239.9 | `FS1:1.60.1.70009:hunter:dwarf:0/32005500005/500230231050120151:` | +0.8% | yes | 0/20/31 (Survival) | yes | Improved Tracking, Deflection, Entrapment, Survivalist, Improved Wing Clip, Deterrence, Survival Tactics, Resourcefulness, Expose Prey, Survivalist's Discipline |
| Arcane (mage-arcane) | 456.4 | 498.5 (guide, modeled non-damage points re-spent) | `FS1:1.60.1.70009:mage:gnome:253225113100011531/032023/005:` | 498.5 | `FS1:1.60.1.70009:mage:gnome:253225113100011531/032023/005:` | +9.2% | yes | 36/10/5 (Arcane) | yes | Wand Specialization, Improved Channeling, Magic Absorption, Arcane Concentration, Arcane Resilience, Arcane Geometry, Arcane Shielding, Improved Counterspell, Arcane Meditation |
| Fire (mage-fire) | 381.4 | 419.9 (guide, all non-damage points re-spent) | `FS1:1.60.1.70009:mage:gnome:221005/23552110030033051/0050002:` | 414.8 | `FS1:1.60.1.70009:mage:gnome:2050151/23552100030023051/005:` | +8.8% | yes | 14/32/5 (Fire) | yes | Flame Throwing, Impact, Improved Fire Ward, Hot Streak |
| Frost (mage-frost) | 218.3 | 378.9 (guide, all non-damage points re-spent + 2 Shatter -> Arcane Meditation) | `FS1:1.60.1.70009:mage:gnome:2250050001003/0/255511133000010105:` | 325.4 | `FS1:1.60.1.70009:mage:gnome:2030050001/113023/253511130000030105:` | +49.1% | yes | 11/10/30 (Frost) | yes | Frost Warding, Permafrost, Frostbite, Improved Blizzard, Arctic Reach, Ice Block, Shatter, Fingers of Frost |
| Retribution (paladin-retribution) | 194.0 | 194.3 (guide, all non-damage points re-spent) | `FS1:1.60.1.70009:paladin:human:551002102001/323/0502533010133:` | 194.0 | `FS1:1.60.1.70009:paladin:human:54/3232/0502533121133021:` | +0.0% | no | 9/10/32 (Retribution) | yes | Deflection, Holy Conduit, Sanctified Judgement, Seal of Command, Pursuit of Justice, Eye for an Eye, Repentance, Champion of the Light, Instrument of Law, Twist of Light |
| Shadow (priest-shadow) | 247.8 | 277.3 (deep Shadow + 1 Improved Mind Flay -> Mental Strength) | `FS1:1.60.1.70009:priest:gnome:325000031302/0/555320001001300151:` | 252.8 | `FS1:1.60.1.70009:priest:gnome:5241110013/0/543110401201300251:` | +2.0% | yes | 18/0/33 (Shadow) | yes | Blackout, Spirit Tap, Shadow Affinity, Shadow Reach, Improved Psychic Scream, Improved Mind Flay, Improved Fade, Silence, Devouring Contagion, Early Demise |
| Assassination (rogue-assassination) | 188.5 | 238.3 (deep Combat) | `FS1:1.60.1.70009:rogue:night-elf:125020104003/32533301200510201/002:` | 206.0 | `FS1:1.60.1.70009:rogue:night-elf:3250001055150105/3252/51:` | +9.3% | yes | 33/12/6 (Assassination) | yes | Improved Gouge, Remorseless Attacks, Ruthlessness, Improved Slice and Dice, Improved Expose Armor, Improved Kidney Shot, Seal Fate |
| Combat (rogue-combat) | 227.9 | 269.7 (deep Combat) | `FS1:1.60.1.70009:rogue:night-elf:005323101005/32533300000510231/0:` | 227.9 | `FS1:1.60.1.70009:rogue:night-elf:32531/32531300000515201/51:` | +0.0% | no | 14/31/6 (Combat) | yes | Lightning Reflexes, Puncturing Wounds, Deflection, Endurance, Riposte, Improved Sprint, Improved Kick, Flawless Execution, Hack and Slash |
| Subtlety (rogue-subtlety) | 183.1 | 267.8 (deep Combat) | `FS1:1.60.1.70009:rogue:night-elf:005323101005/32533300000510231/0:` | 183.1 | `FS1:1.60.1.70009:rogue:night-elf:005/325131/5322210310013011051:` | +0.0% | no | 5/15/31 (Subtlety) | yes | Camouflage, Master of Deception, Opportunity, Setup, Elusiveness, Dirty Tricks, Improved Ambush, Initiative, Improved Distract, Heightened Senses, Dirty Deeds, Quietus, Cutthroat, Thousand Cuts |
| Elemental (shaman-elemental) | 143.8 | 176.6 (guide, all non-damage points re-spent) | `FS1:1.60.1.70009:shaman:dwarf:553233130010305/055/052:` | 161.0 | `FS1:1.60.1.70009:shaman:dwarf:453231130010305/0/553322:` | +12.0% | yes | 31/0/20 (Elemental) | yes | Elemental Warding, Reverberation, Elemental Devastation, Elemental Alacrity, Improved Fire Nova, Eye of the Storm, Call of Thunder, Elemental Reach, Lightning Overload, Earthbound |
| Enhancement (shaman-enhancement) | 182.0 | 195.4 (guide, all non-damage points re-spent) | `FS1:1.60.1.70009:shaman:dwarf:5500331/155030030005/0530202:` | 182.0 | `FS1:1.60.1.70009:shaman:dwarf:553322/253130030005102051/0:` | +0.0% | no | 20/31/0 (Enhancement) | yes | Earth's Grasp, Ancestral Knowledge, Guardian Totems, Mental Dexterity, Improved Ghost Wolf, Improved Lightning Shield, Shamanistic Focus, Anticipation, Toughness, Stormstrike, Spirit Weapons, Mental Quickness, Improved Stormstrike, Maelstrom Weapon, Rage of the Farseer |
| Affliction (warlock-affliction) | 351.0 | 381.6 (guide, all non-damage points re-spent) | `FS1:1.60.1.70009:warlock:gnome:25552000120201051/0005/0550005:` | 351.0 | `FS1:1.60.1.70009:warlock:gnome:25552000130201051/235522/0:` | +0.0% | no | 32/19/0 (Affliction) | yes | Malediction, Soul Harvesting, Improved Drains, Improved Bane of Agony, Fel Concentration, Pandemic, Malevolence, Curse of Exhaustion, Soul Siphon |
| Demonology (warlock-demonology) | 283.0 | 375.3 (deep Demonology) | `FS1:1.60.1.70009:warlock:gnome:205/230523323012000005/0550003:` | 283.0 | `FS1:1.60.1.70009:warlock:gnome:255323/2352113101200001351/0:` | +0.0% | no | 20/31/0 (Demonology) | yes | Improved Health Funnel, Improved Imp, Demonic Embrace, Demonic Aegis, Improved Voidwalker, Fel Vitality, Demonic Energies, Demonic Sacrifice, Master Summoner, Fel Domination, Demonic Brand, Improved Felhunter, Soul Link, Demonic Knowledge, Demonic Pact |
| Destruction (warlock-destruction) | 278.3 | 315.9 (deep Demonology + 1 Agonizing Flames -> Suppression) | `FS1:1.60.1.70009:warlock:gnome:22/235521323000000005/20503051:` | 279.3 | `FS1:1.60.1.70009:warlock:gnome:255323/0/2353325100001051:` | +0.4% | yes | 20/0/31 (Destruction) | yes | Destructive Reach, Improved Shadow Bolt, Molten Skin, Aftermath, Intensity, Agonizing Flames, Pyroclasm, Bane of Havoc, Fire and Brimstone, Shadow and Flame |
| Arms (warrior-arms) | 201.4 | 225.6 (deep Arms) | `FS1:1.60.1.70009:warrior:human:05325213032015001/055500000012/0:` | 210.3 | `FS1:1.60.1.70009:warrior:human:05325213032013001/0505/5005:` | +4.4% | yes | 31/10/10 (Arms) | yes | Deflection, Improved Charge, Improved Tactical Mastery, Spearing Strike, Bloodthrill, Sweeping Strikes, Improved Slam, Improved Hamstring |
| Fury (warrior-fury) | 259.5 | 279.9 (deep Fury) | `FS1:1.60.1.70009:warrior:human:3/554500005052310051/052:` | 260.8 | `FS1:1.60.1.70009:warrior:human:35311103002/353310005050010051/0:` | +0.5% | yes | 19/32/0 (Fury) | yes | Booming Voice, Iron Will, Blood Craze, Boundless Rage, Raging Blows, Enrage, Improved Intercept, Improved Berserker Rage |

## Table 2: engine work list — unmodeled talents that plausibly affect DPS

Per class, every talent whose final classification is `unmodeled` in at least one of that class's
reports **and never reaches `damage` in any of them** - i.e. the engine genuinely never credits it in
any of this class's specs, not just a talent that happened to sit outside one spec's build. Tooltip
text is Forever's own (`data/raw-forever/wowhead-talents-2026-09-14.json`), not vanilla's, since
Forever rewrote several talents' meaning (notably four PvP resist-reduction talents into global hit
talents - `research/08-stats.md` §1.2). Obvious non-DPS talents (healing, threat, pure CC, movement,
pure defense/parry/dodge/armor) are left out even when they sit in a spec's own tree; a few marked
"marginal" touch damage only indirectly or need a shield/stance/rotation choice the guide may not make.

Flagged specially where Forever's rewrite makes an off-role-tree talent globally relevant: **Nature's
Reach** (Druid Balance), **Tidal Focus** (Shaman Restoration), **Divine Precision** (Paladin Holy) and
**Holy Precision** (Priest Discipline) all grant global or school hit chance under Forever's talent
rewrite, not vanilla's resist reduction - and Paladin Retribution's own tree hides **Seal of Command**,
the spec's core damage seal, among the unmodeled set.

| Class | Tree | Talent | Why it plausibly affects DPS |
|---|---|---|---|
| Druid | Balance | Eclipse | Wrath crit reduces next 2 Starfire cast times |
| Druid | Balance | Genesis | +5% periodic damage (and healing) |
| Druid | Balance | Improved Moonfire | +10% Moonfire damage and crit |
| Druid | Balance | Improved Wrath | Wrath cast time -0.5s |
| Druid | Balance | Nature's Grace | +10% cast speed for 3s on spell crit |
| Druid | Balance | Nature's Majesty | +4% crit, spells and melee |
| Druid | Balance | Nature's Reach | +4% hit, all spells/attacks (Forever-redefined, was range+resist) |
| Druid | Balance | Nature's Splendor | +duration on Moonfire/Insect Swarm (and Regrowth) |
| Druid | Balance | Improved Entangling Roots | +75% Entangling Roots damage (marginal, CC-primary) |
| Druid | Feral Combat | Blood Frenzy (= Primal Fury) | combo point / Bear rage on crit |
| Druid | Feral Combat | Feral Instinct | +30% Swipe damage |
| Druid | Feral Combat | King of the Jungle | Tiger's Fury grants 60 Energy |
| Druid | Feral Combat | Predatory Instincts | +20% melee crit damage bonus |
| Druid | Feral Combat | Primal Bite (= Mangle) | direct-damage finisher the Forever rename hides |
| Druid | Feral Combat | Rend and Tear | +10% melee damage vs Bleeding targets |
| Druid | Feral Combat | Shredding Attacks | Shred Energy / Lacerate Rage cost cut |
| Druid | Restoration | Furor | Energy/Rage preserved on shapeshift (Cat/Bear rotation resource) |
| Druid | Restoration | Nature's Swiftness | instant-cast next Nature spell (burst Starfire/Wrath window) |
| Hunter | Beast Mastery | Deadly Aspects | 10% proc: +30% ranged/melee attack speed (haste) |
| Hunter | Beast Mastery | Focused Fire | +2% all damage while pet active |
| Hunter | Beast Mastery | Summon Hawk | pet hawk deals direct Physical damage |
| Hunter | Marksmanship | Careful Aim | Attack Power from 100% Intellect |
| Hunter | Marksmanship | Improved Arcane Shot | Arcane Shot cooldown -1.5s |
| Hunter | Marksmanship | Improved Stings | +20% Serpent Sting damage |
| Hunter | Marksmanship | Lethal Attacks | +5% crit, all attacks |
| Hunter | Marksmanship | Lone Wolf | +20% damage with no active pet |
| Hunter | Marksmanship | Rapid Killing | Rapid Fire CD -2min + on-kill Shot damage proc |
| Hunter | Marksmanship | Rapid Recuperation | mana regen while casting (resource) |
| Hunter | Marksmanship | Sniper Shot | +160 ranged damage |
| Hunter | Marksmanship | Trueshot Aura | +30 ranged Attack Power (party) |
| Hunter | Survival | Expose Prey | 10% chance to proc Mongoose Bite on Hunter's Mark target |
| Hunter | Survival | Improved Tracking | +5% damage vs tracked creature type |
| Hunter | Survival | Resourcefulness | -60% trap/melee mana cost + regen-while-casting on crit |
| Hunter | Survival | Survival Tactics | +10% hit, Trap/Feign Death (marginal, trap damage only) |
| Mage | Fire | Hot Streak | crit stacks cut Pyroblast cast time |
| Mage | Frost | Fingers of Frost | free "Frozen" procs for next 2 spells (enables Shatter) |
| Mage | Frost | Shatter | +50% spell crit vs Frozen targets |
| Paladin | Holy | Divine Precision | +18% Holy spell hit (Forever-redefined, was new) |
| Paladin | Holy | Improved Seals | +15% Seal/Judgement damage - Ret's core scaling |
| Paladin | Holy | Holy Power | +15% Holy Shock crit, +5% crit all other spells |
| Paladin | Holy | Holy Shock | direct Holy damage option |
| Paladin | Holy | Divine Favor | 100% crit proc usable on Holy Shock |
| Paladin | Holy | Purifying Power | -33% cooldown on Exorcism and Holy Wrath (both damage spells) |
| Paladin | Holy | Light's Vigil | can deal 175-189 Holy damage instead of healing |
| Paladin | Holy | Consecrated Ground | +10% Holy spell damage near Consecration |
| Paladin | Protection | One-Handed Weapon Specialization | +10% damage, one-hand melee weapons |
| Paladin | Protection | Reckoning | extra-attack proc off Block/being crit |
| Paladin | Protection | Swift Judgement | resets Judgement cooldown, free next cast |
| Paladin | Protection | Holy Shield | 110 Holy damage per block while active (marginal, needs shield) |
| Paladin | Retribution | Seal of Command | **core Ret damage seal** - proc + Judgement damage |
| Paladin | Retribution | Champion of the Light | spell damage/healing from 100% Intellect |
| Paladin | Retribution | Holy Conduit | -40% mana cost on Consecration/Holy Wrath/Exorcism/Hammer of Wrath |
| Paladin | Retribution | Instrument of Law | Hammer of Wrath cast time -1.0s |
| Paladin | Retribution | Sanctified Judgement | 60% mana refund on Judgement (resource) |
| Paladin | Retribution | Twist of Light | seal-twisting damage-application mechanic |
| Priest | Discipline | Holy Precision | +18% Holy spell hit (Forever-redefined, was resist) |
| Priest | Discipline | Power Infusion | +20% spell damage and healing, 15s cooldown |
| Priest | Discipline | Twin Disciplines | +5% damage/healing on instant-cast spells |
| Priest | Discipline | Power in Light | +10% Smite/Penance vs Holy Fire targets (marginal) |
| Priest | Discipline | Penance | direct Holy damage option (marginal) |
| Priest | Holy | Divine Fury | Smite/Holy Fire cast time -0.5s |
| Priest | Holy | Holy Specialization | +5% crit, Holy spells |
| Priest | Holy | Searing Light | +5% Holy damage + free-Holy Nova proc |
| Priest | Holy | Holy Nova | AoE direct Holy damage (marginal, mostly a heal) |
| Priest | Shadow | Devouring Contagion | -50% Devouring Plague mana cost + spreads it on kill |
| Priest | Shadow | Early Demise | +30% Shadow Word: Death crit on execute targets |
| Priest | Shadow | Improved Mind Flay | +20% Mind Flay damage |
| Priest | Shadow | Spirit Tap | +100% Spirit and mana-regen-while-casting after a kill (resource) |
| Rogue | Assassination | Improved Expose Armor | combo point refund (resource) |
| Rogue | Assassination | Improved Kidney Shot | +10% damage to targets stunned by Kidney Shot |
| Rogue | Assassination | Remorseless Attacks | +40% crit on next finisher after a kill |
| Rogue | Combat | Hack and Slash | weapon-type damage/crit/armor-pen bonus |
| Rogue | Combat | Puncturing Wounds | +30%/+15% crit Backstab/Mutilate + combo point chance |
| Rogue | Combat | Riposte | bonus-damage attack proc off a parry |
| Rogue | Subtlety | Cutthroat | enables stealth-free Ambush (marginal, opener flexibility) |
| Rogue | Subtlety | Dirty Deeds | Garrote energy cost cut (marginal, bundled with Cheap Shot) |
| Rogue | Subtlety | Improved Ambush | +45% crit, Ambush |
| Rogue | Subtlety | Initiative | free combo point on opener (Ambush/Garrote/Cheap Shot) |
| Rogue | Subtlety | Quietus | +10% damage vs targets below 35% health |
| Rogue | Subtlety | Setup | free combo point on Dodge/resist (resource) |
| Rogue | Subtlety | Thousand Cuts | Rupture ticks cut Hemorrhage/Backstab energy cost (resource) |
| Shaman | Elemental | Call of Thunder | +3% crit, Lightning Bolt/Chain Lightning |
| Shaman | Elemental | Elemental Alacrity | -0.5s cast time, LB/CL/Lava Burst (haste) |
| Shaman | Elemental | Elemental Devastation | spell crit procs +9% melee crit |
| Shaman | Elemental | Improved Fire Nova | +20% Fire Nova damage, -4s cooldown |
| Shaman | Elemental | Lightning Overload | 10% chance to cast a free second LB/CL |
| Shaman | Elemental | Reverberation | Shock spells cooldown -1s |
| Shaman | Enhancement | Improved Lightning Shield | +15% Lightning Shield orb damage |
| Shaman | Enhancement | Improved Stormstrike | mana regen + cooldown reset on Dodge/Parry |
| Shaman | Enhancement | Maelstrom Weapon | melee-damage proc cuts Lightning Bolt cast time/cost |
| Shaman | Enhancement | Mental Dexterity | Attack Power from 100% Intellect |
| Shaman | Enhancement | Mental Quickness | spell damage from Intellect |
| Shaman | Enhancement | Rage of the Farseer | +30% melee and cast speed for 25s (haste) |
| Shaman | Enhancement | Shamanistic Focus | -45% Shock/Lightning Shield mana cost (resource) |
| Shaman | Enhancement | Stormstrike | **core Enhancement ability** - direct damage + 20% Nature damage debuff |
| Shaman | Restoration | Tidal Focus | +5% hit, all spells/attacks (Forever-redefined, was mana cost only) |
| Warlock | Affliction | Improved Bane of Agony | +10% Bane of Agony damage |
| Warlock | Affliction | Improved Drains | up to +18% Drain Life/Soul damage, tripled under 20% health |
| Warlock | Affliction | Malediction | +5% periodic damage, all Warlock spells |
| Warlock | Affliction | Malevolence | +5% crit, Shadow spells |
| Warlock | Affliction | Pandemic | +100% crit damage bonus on Corruption/Bane/Drain DoTs |
| Warlock | Affliction | Soul Harvesting | mana regen on a Drain Soul kill (resource) |
| Warlock | Affliction | Soul Siphon | +50% Drain Life/Soul damage rate |
| Warlock | Demonology | Demonic Brand | pet deals bonus Fire/Shadow damage (pet damage) |
| Warlock | Demonology | Demonic Knowledge | +spell damage while a Demon pet is active |
| Warlock | Demonology | Demonic Pact | sustains the Demonic Sacrifice damage buff across resummons |
| Warlock | Demonology | Demonic Sacrifice | +15% Shadow/Fire damage - a core caster cooldown |
| Warlock | Demonology | Improved Imp | +30% Imp Firebolt damage (pet damage) |
| Warlock | Demonology | Soul Link | +3% damage, both pet and master |
| Warlock | Destruction | Aftermath | +50% Immolate initial damage |
| Warlock | Destruction | Agonizing Flames | +10% Searing Pain crit, +10% all Destruction damage |
| Warlock | Destruction | Bane of Havoc | clones 15% of all damage dealt onto a second target |
| Warlock | Destruction | Fire and Brimstone | +25% crit, Conflagrate |
| Warlock | Destruction | Shadow and Flame | +10% Shadow/Fire damage procs + Shadowburn shard refund |
| Warrior | Arms | Bloodthrill | free Overpower proc off Rend ticks |
| Warrior | Arms | Improved Charge | +6 Rage from Charge (resource) |
| Warrior | Arms | Improved Slam | Slam cast time/GCD -0.5s |
| Warrior | Arms | Improved Tactical Mastery | +15 Rage retained on stance change (resource) |
| Warrior | Arms | Spearing Strike | direct-damage ability |
| Warrior | Arms | Sweeping Strikes | next 5 melee attacks cleave an extra target |
| Warrior | Fury | Boundless Rage | +30 max Rage (resource) |
| Warrior | Fury | Enrage | 30% chance: +10% physical damage for 12s after being hit |
| Warrior | Fury | Improved Berserker Rage | +10 Rage on Berserker Rage (resource) |
| Warrior | Fury | Raging Blows | Whirlwind hits with off-hand too; Cleave cost -2 Rage |
| Warrior | Protection | Bastion | +10% damage with a shield equipped (marginal, needs shield) |
| Warrior | Protection | Focused Rage | -3 Rage cost, offensive abilities (resource) |
| Warrior | Protection | Improved Revenge | +60% Revenge damage (marginal, tank ability) |
| Warrior | Protection | Master of Defense | +5 Rage on Dodge/Parry with a shield (marginal, resource) |

Not scored (no matching tooltip in the Forever Wowhead snapshot, so its effect could not be verified
one way or the other): Rogue Combat's **Flawless Execution**. Worth a direct look since everything
else around it in that tier is damage-relevant.

## Reports

`design/reviews/talent-search/*.md` (all 20 overwritten). Supporting files: 20-row data at
`/private/tmp/claude-501/-Users-jh-code-forever/54c69daf-a689-41fc-ad89-a21df2102a2f/scratchpad/parsed.json`
and `by_class.json` (not committed - scratchpad only).
