# Slop sweep 2026-11-09: grep hits and judgment pass

Lane slop-sweep. Patterns live in `tools/slop/patterns.txt` (one regex per tell), sources in `tools/slop/sources.txt` and the allowlist in `tools/slop/allowlist.txt`. Run `tools/slop/scan.sh`, with `--md`, `--counts`, `--check`, `--skip <prefix>` or `--probe <text>`. The guards are `web/src/slop-copy.test.ts` (everything outside api/) and `api/internal/slopguard` (the api trees).

## Hit counts per pattern

"Before" is origin/main 66c709939 scanned with the final patterns. "After" is this branch.

| pattern | before | after |
|---|---|---|
| em-dash | 32 | 2 |
| is-the-colon | 50 | 0 |
| nudge-prefix | 7 | 0 |
| here-is-what | 3 | 0 |
| lets | 0 | 0 |
| buzzword | 0 | 0 |
| stock-promise | 3 | 0 |
| simply-just | 0 | 0 |
| exclamation | 0 | 0 |
| grade-hype | 2 | 0 |
| explain-obvious | 1 | 0 |
| whether-youre | 0 | 0 |
| not-just | 0 | 0 |

The two "after" hits are `api/internal/discordbot/copy_coach.go:20,22` ("Fix first: {stem} — on ..."). The discordbot belongs to lane f-discord, and its branch (origin/lane/f-discord) already drops both em dashes. The api guard goes green once that lands.

Allowlisted: "Eureka!" (the Gnome racial's real name, in the guides and `race-table.ts`) and the em dash inside the regexes of the two guide note parsers (`lib/guides/key-talents.ts`, `rotation-lists.ts`). A lone `'—'` used as an empty table cell is notation, and the pattern does not match it.

## Full hit list before the sweep (origin/main)

### em-dash (32)

- `addon/ForeverSixty/Locale.lua:1427` L.diagSiteNever = "Site: not seen yet \226\128\148 is the companion running and signed in?"
- `api/internal/coach/copy.go:28` ThinPool          = "not enough logs yet — {n} {Spec} kills of {Boss} ({size}) on the site, the coach needs 20."
- `api/internal/discordbot/copy_coach.go:20` Line:      "Fix first: {stem} — on {n} of {m} bosses. [Open your coach's review]({url})",
- `api/internal/discordbot/copy_coach.go:22` LineAll:   "Fix first: {stem} — on all {m} bosses. [Open your coach's review]({url})",
- `web/src/content/changelog/2026-09-20-simulator-tools-are-live.md:18` None of it is graded yet. [Spec support](/sim/specs) still reads "Not yet" for every damage spec — nothing has been checked against a real parse, because there are no real parses yet. And Droptimizer can only rank what t
- `web/src/content/pages/privacy.md:15` never collect, see, or store your card number, its expiry, or your CVC — Stripe holds all of
- `web/src/content/pages/privacy.md:34` `fs_session` and `fs_csrf` are functional cookies only — they keep you signed in and protect
- `web/src/content/pages/refunds.md:19` the refund — access ends at that point, not automatically the instant the refund is issued,
- `web/src/content/pages/refunds.md:27` went wrong on our side — you were charged twice, or paid for something that did not work —
- `web/src/content/pages/terms.md:15` group. Nothing from Battle.net — your characters, your sign-in, anything the game itself
- `web/src/content/pages/terms.md:16` tells us about you — is ever sold or gated behind a subscription.
- `web/src/content/pages/terms.md:28` [refund policy](/refunds)). After that, the account works exactly like a free one — nothing is
- `web/src/lib/bis/copy.ts:407` firstBandEveryPickNew: 'This is the first level range — every pick here is new.',
- `web/src/lib/coach/privacy-section.ts:38` `When a Premium player or an officer opens a coach review, one fight's numbers — spec, class, role, boss, raid size, kill or wipe, duration, level and the spells cast — are sent through the Vercel AI Gateway, which keeps
- `web/src/lib/rating/copy.ts:74` absoluteBasis: 'Measured against the encounter’s own numbers — not enough logs yet to compare players.',
- `web/src/lib/rating/copy.ts:79` return `${componentLabel} — not scored. This fight’s encounter has no curated mechanics table yet.`;
- `web/src/lib/rating/copy.ts:81` return `${componentLabel} — not scored. The simulator does not model this spec yet; scores return once it does.`;
- `web/src/lib/rating/copy.ts:83` return `${componentLabel} — not scored. No curated utility table exists for this spec yet.`;
- `web/src/lib/rating/copy.ts:85` return `${componentLabel} — not scored. No consumable catalogue exists for this role yet.`;
- `web/src/lib/rating/copy.ts:87` return `${componentLabel} — no character sheet in this log.`;
- `web/src/lib/rating/copy.ts:89` return `${componentLabel} — not scored. This character is not signed in, and the bracket has too few logs to compare against.`;
- `web/src/lib/rating/copy.ts:91` return `${componentLabel} — not scored.`;
- `web/src/lib/rating/copy.ts:94` threatNotModeled: 'Threat — not modeled yet. This does not count for or against Utility.',
- `web/src/lib/rating/copy.ts:97` 'capped from a higher weighted average — a costly avoidable death outweighs the rest of the fight. See Survival.',
- `web/src/lib/report/glossary.ts:47` 'How many times an aura was put on that unit inside the window, refreshes not counted. Under Auras applied, Events adds a "refreshed" line wherever the client wrote its own refresh event; a spell that re-applies itself i
- `web/src/lib/sim/copy.ts:673` rotationNotSimulated: (name: string): string => `No rotation yet — ${name} is not simulated.`,
- `web/src/lib/sim/copy.ts:752` 'One target that becomes three at 40 seconds, five at 80, back to three at 130 and one at 160 — a fixed schedule, not a repeating pull. No execute phase at any point, and the count stays at one for the rest of the fight 
- `web/src/lib/sim/copy.ts:754` 'One target that takes no debuffs, no execute phase and no armor reduction — a plain damage check, not a boss fight.',
- `web/src/lib/sim/copy.ts:775` targetLevelHelp: 'The boss level the run’s numbers — armor, resistances — are drawn from.',
- `web/src/lib/sim/pool-quality-copy.ts:88` 'The build in the editor has not been added to the comparison yet — click “Add this build to the comparison” first, or RUN will leave it out.',
- `web/src/lib/sim/pool-quality-copy.ts:111` plannerDpsOwnGearNote: 'With the gear shown above — the comparison below always uses your current gear.',
- `web/src/lib/sim/pool-quality-copy.ts:118` talentsGearLockedNote: 'Every row above uses your current gear — only the talent tree changes.',

### is-the-colon (50)

- `web/src/content/changelog/2026-09-13-transmog-appearance-rules.md:19` When a bind-on-pickup uncommon or rare item drops in a dungeon, its appearance is granted to every eligible looter, so nobody has to roll against an upgrade just to collect the look. Epic raid drops are the exception: on
- `web/src/content/guides/druid/balance.md:45` Balance is the Moonkin: two DoTs on the boss, Wrath to bank Eclipse charges, and fast Starfires to spend them, all from a form that gives your group crit.
- `web/src/content/guides/druid/balance.md:115` Night Elf brings Elune's Light, 10% crit for 15 seconds on a 3 minute cooldown, and Vengeance makes that crit worth more to you. Racials work in Moonkin Form in Forever, so press it with your trinkets. Tauren is the Hord
- `web/src/content/guides/druid/feral.md:101` Night Elf brings Elune's Light, 10% crit for 15 seconds on a 3 minute cooldown, and it works in Cat Form. Racials are usable while shapeshifted in Forever, so line it up with your potion and trinkets. Tauren is the Horde
- `web/src/content/guides/druid/feral.md:105` Enchanting is the one profession your raid set asks for: Molten Heart of the Mountain, and on the Alliance Frozen Heart of the Mountain beside it, are bind-on-pickup trinkets only an enchanter can make. The Leatherworkin
- `web/src/content/guides/druid/restoration.md:101` Night Elf brings Elune's Light, 10% crit for 15 seconds, which is crit on your Regrowths and Healing Touches alike. Tauren is the Horde pick: Endurance is 5% more health, which keeps you alive through raid damage. Skybor
- `web/src/content/guides/druid/restoration.md:123` Restoration is the slowest druid to quest with, because it is built to keep others alive, not to kill. Level as Feral and switch at 60, or heal dungeons on the way up: Rejuvenation and Regrowth carry a group from the low
- `web/src/content/guides/hunter/beast-mastery.md:48` Beast Mastery is the pet spec: your cat does a big share of the damage, Bestial Wrath is the burst, and your shots fill every gap.
- `web/src/content/guides/hunter/beast-mastery.md:81` Endurance Training (4 points) keeps the pet alive through raid damage, Bestial Discipline (2 points) keeps half your mana regeneration running while you shoot, and Intimidation (1 point) is the way to Bestial Wrath and a
- `web/src/content/guides/hunter/beast-mastery.md:101` **The pet.** A cat is the safe raid pet: Claw and Bite on a fast attacker, steady damage with nothing to manage. Train it the top ranks of Claw, Bite and Growl as you level, then turn off Growl in raids (right-click it o
- `web/src/content/guides/hunter/beast-mastery.md:105` **Trash.** Multi-Shot hits three targets and is your AoE on most packs. Volley is a channel, costs a lot of mana, and is worth it only on a big pack that will live through the whole channel. Freezing Trap is the hunter's
- `web/src/content/guides/hunter/beast-mastery.md:137` Enchanting is the profession your best in slot list asks for. Frozen Heart of the Mountain is an Enchanting trinket that binds on pickup, so nobody can make it for you. Legionite Glaive is a Blacksmithing polearm that al
- `web/src/content/guides/hunter/marksmanship.md:143` Enchanting is the profession your best in slot list asks for. The Heart of the Mountain trinkets are Enchanting crafts that bind on pickup, so nobody can make them for you. Legionite Glaive is a Blacksmithing polearm tha
- `web/src/content/guides/hunter/survival.md:91` Then work through your melee priority: Raptor Strike, Mongoose Bite, Lacerate, Strider Kick, Immolation Trap. Raptor Strike replaces your next white hit instead of costing a global, so the only mistake is letting it sit 
- `web/src/content/guides/hunter/survival.md:149` At 31 you have only the first point of Resourcefulness, and for that one level Raptor Strike rank 1 is the better button: it costs almost nothing, and the mana goes into traps. Drag it onto the bar at 31 and put rank 5 o
- `web/src/content/guides/mage/arcane.md:98` **Gnome** is the best Alliance Arcane mage: Eureka! makes your damage spells, all but Arcane Wrath, hit 10% harder and cost 10% less for 8 seconds on a 2 minute cooldown, and Expansive Mind adds 5% mana. Press it with Ar
- `web/src/content/guides/mage/arcane.md:102` Tailoring and Enchanting are the standard pair: Tailoring makes the early cloth sets and Enchanting puts the spell power enchants on your own gear. Alchemy pays for itself once you raid every week, because you will drink
- `web/src/content/guides/mage/fire.md:101` **Human** is the Alliance pick for Fire: Ruthless adds 1% crit with all spells, which Fire turns into Ignite and Heating Up, and The Human Spirit adds 10% Spirit. **Gnome** is just behind: Eureka! makes your damage spell
- `web/src/content/guides/mage/frost.md:99` **Gnome** is the best Alliance Frost mage among the Classic races: Eureka! makes your damage spells hit 10% harder and cost 10% less for 8 seconds on a 2 minute cooldown, and Expansive Mind adds 5% mana. **Human** adds 1
- `web/src/content/guides/mage/index.md:36` - **Gnome** is the Alliance pick for Arcane and Frost: Eureka! makes your damage spells, all but Arcane Wrath, hit 10% harder and cost 10% less for 8 seconds, every 2 minutes. Arcane Power comes back every 3, so pair Eur
- `web/src/content/guides/paladin/retribution.md:60` - **Benediction** (5 points): your Seals, Judgement, Holy Strike, Consecration and Exorcism cost 10% less. Hammer of Wrath is the exception: Champion of the Light makes it instant, but Benediction still does not discount
- `web/src/content/guides/paladin/retribution.md:62` - **Conviction** (3 points): 6% crit on everything you do: your swings, Holy Strike, Seal procs, Judgements and Consecration's ticks. Exorcism is the exception: it never crits in Forever.
- `web/src/content/guides/paladin/retribution.md:117` Mining and Blacksmithing make your two-handers and plate. Engineering adds explosives for trash packs, and it is the profession that adds damage when your gear is done. Herbalism and Alchemy keep you in mana potions and 
- `web/src/content/guides/priest/discipline.md:104` **Undead** is the Horde pick for Dark Sacrifice: well over a thousand mana once a fight, paid in your own health while the raid is quiet. Will of the Forsaken breaks fears.
- `web/src/content/guides/priest/holy.md:26` Holy is the priest who heals everything: Prayer of Mending and Renew on the tank, Greater Heal as the filler, and Prayer of Healing for a hurt group.
- `web/src/content/guides/priest/holy.md:30` - Prayer of Mending is the cheapest healing you can cast all fight: up to five heals for one cast
- `web/src/content/guides/priest/holy.md:130` Greater Heal is the same story: keep rank 1 as the filler all the way to 60, even after rank 4 arrives at 58. A dungeon tank rarely needs the bigger ranks, so the mana they cost is better spent on more casts. At 60, rank
- `web/src/content/guides/rogue/assassination.md:30` Assassination is the Forever rogue built on Mutilate: two daggers, two combo points a strike, Seal Fate crits for a third and a Cold Blood Eviscerate whenever it is up.
- `web/src/content/guides/rogue/combat.md:30` Combat is the steadiest rogue in a raid: a dagger in each hand, Backstab from behind, Eviscerate at four points and Adrenaline Rush with Blade Flurry for the burst.
- `web/src/content/guides/rogue/combat.md:52` The raid build goes 12/34/5: Combat to Adrenaline Rush with all five points of Flash of Steel, Assassination for Malice, Ruthlessness, Improved Slice and Dice and Lethality, and five points of Opportunity for Backstab. T
- `web/src/content/guides/rogue/combat.md:116` **Enchanting** is the one that changes your gear: Molten Heart of the Mountain and Frozen Heart of the Mountain are bind-on-pickup trinkets only an enchanter can make. **Engineering** brings grenades and dynamite for tra
- `web/src/content/guides/rogue/index.md:35` Racials are reworked and none of them adds weapon skill. Undead is the Horde pick in every tree: **Touch of the Grave** gives your attacks a chance to drain up to 5% of your maximum health from the target, at most once e
- `web/src/content/guides/rogue/subtlety.md:29` Subtlety is the bleed rogue in Forever: a Premeditation opener, Hemorrhage and Boot Knife Shanks to build, Rupture as the main finisher and Death Mark to cash the bleeds in.
- `web/src/content/guides/rogue/subtlety.md:117` **Enchanting** makes Molten Heart of the Mountain and Frozen Heart of the Mountain, bind-on-pickup trinkets only an enchanter can make. **Engineering** is the PvP and trash pick: grenades, dynamite and gadgets. **Skinnin
- `web/src/content/guides/shaman/elemental.md:30` Elemental is the shaman's caster spec: Flame Shock kept ticking, Lightning Bolt for almost every cast, Lightning Mastery with your cooldowns, and Chain Lightning on more than one target.
- `web/src/content/guides/shaman/elemental.md:104` On the Horde, Orc is the damage pick: Blood Fury raises your spell power by 10% for 15 seconds every 2 minutes. Press it with Lightning Mastery on the pull and on its own cooldown after that. Troll's Berserking, 10% cast
- `web/src/content/guides/shaman/enhancement.md:30` Enhancement is the shaman's melee spec: one weapon and its imbue, Stormstrike on cooldown, and a free instant Lightning Bolt whenever Maelstrom Weapon reaches five stacks.
- `web/src/content/guides/shaman/enhancement.md:107` On the Horde, Orc is the pick: Blood Fury is 10% attack power for 15 seconds every 2 minutes. Press it with Rage of the Farseer on the pull and on its own cooldown after that. Troll's Berserking adds 10% attack speed for
- `web/src/content/guides/shaman/index.md:28` - **Lightning Mastery** is the 31-point Elemental talent: for 15 seconds every 3 minutes Lightning Bolt casts half a second faster and Chain Lightning has no cooldown. Call of Thunder, 5% crit on both lightning spells, i
- `web/src/content/guides/shaman/restoration.md:26` Restoration is the shaman's healing spec: Riptide on the tank, Chain Heal through the group, and Water Shield and Mana Tide to keep the mana coming.
- `web/src/content/guides/shaman/restoration.md:106` On the Horde, Orc is the healing pick: Blood Fury increases attack power and spell power by 10% for 15 seconds every 2 minutes, and its spell power covers your heals. Troll's Berserking, 10% casting speed for 15 seconds 
- `web/src/content/guides/shaman/restoration.md:110` Enchanting is the one profession your raid set asks for: Dormant Heart of the Mountain binds on pickup, so nobody can make it for you. Pair it with Herbalism or Alchemy for Major Mana Potions and flasks, which is most of
- `web/src/content/guides/warlock/affliction.md:72` Never clip a DoT. A recast throws away whatever was left on the old one, and Bane of Agony saves its biggest ticks for the end. Recast Corruption and Siphon Life when they fall off. While you level without the sacrifice,
- `web/src/content/guides/warlock/index.md:27` - **Your demon scales from you.** The raid's group buffs skip it: Prayer of Fortitude, Gift of the Wild, Arcane Brilliance, the Greater Blessings, Battle Shout, and the Strength of Earth, Grace of Air and Windfury totems
- `web/src/content/guides/warrior/arms.md:42` Arms is the two-handed warrior: Rend, Overpower and Mortal Strike in Battle Stance, a stance swap for Recklessness, and a full bar of rage spent on one huge Execute.
- `web/src/content/guides/warrior/arms.md:115` Enchanting is the one that changes your gear: Frozen Heart of the Mountain is a bind-on-pickup trinket only an enchanter can make. Blacksmithing and Mining are the other pair, but Arcanite Champion binds when equipped, s
- `web/src/content/guides/warrior/fury.md:41` Fury is the dual-wield warrior: Berserker Stance all fight, Bloodthirst and Whirlwind on cooldown, Heroic Strike to bleed off spare rage, and Death Wish and Recklessness held for the Executes.
- `web/src/content/guides/warrior/fury.md:118` Enchanting is the one that changes your gear: Frozen Heart of the Mountain, and on the Alliance Molten Heart of the Mountain beside it, are bind-on-pickup trinkets only an enchanter can make. Blacksmithing makes melee we
- `web/src/content/guides/warrior/protection.md:47` Protection is the shield tank: Defensive Stance all fight, Shield Slam and Revenge for threat, Shield Block and Thunder Clap against every hit, Sunder Armor in between.
- `web/src/lib/report/glossary.ts:140` 'How a boss pull ended, and which attempt it was when a boss took more than one. The percentage after a wipe is the boss’s health when the pull ended: how far the pull got.',

### nudge-prefix (7)

- `addon/ForeverSixty/CHANGELOG.md:140` - The Forever Sixty line beside the minimap is on screen at a first login and says your next step ("Next: Endurance Training (1 of 5)", or "Pick a build: /fs"); a saved spot that is off screen puts it back beside the min
- `addon/ForeverSixty/Locale.lua:101` talentStripNextTalent = "Next: %s",
- `addon/ForeverSixty/Locale.lua:970` L.dialogNowFollowingNext = "Now following %s. Next: %s."
- `web/src/lib/addon/copy.ts:242` alt: 'A hunter’s talent window, Beast Mastery, Marksmanship and Survival, with small order marks on the planned talents, one ringed in gold in Marksmanship, and the strip along the bottom: “Next: Lethal Attacks (1 of 5) 
- `web/src/lib/character/page-view.ts:174` text: `Next: spend your ${toSpend} talent ${toSpend === 1 ? 'point' : 'points'}${then}.`,
- `web/src/lib/character/page-view.ts:180` text: `Next: get ${top.name} from ${top.source}.`,
- `web/src/lib/sim/result-copy.ts:31` nextStep: 'Next: find the upgrade that adds the most.',

### here-is-what (3)

- `api/internal/coach/copy.go:50` RulesOpening = "Here is what your log shows for {Boss}."
- `web/src/content/guides/rogue/subtlety.md:123` Combat levels faster with a sword, so if you only want Subtlety for raids, level Combat and respec at 60. If you level Subtlety, here is how it starts. You start with Sinister Strike, Eviscerate and Stealth. Backstab tra
- `web/src/lib/sim/copy.ts:1206` sourcesAllGated: 'Nothing here has opened yet. Here is what is known about each:',

### stock-promise (3)

- `web/src/content/changelog/2026-10-25-discord-notifications.md:17` Private logs are your choice. New webhooks leave them off; a webhook saved before today keeps posting them until you untick Include private logs on the account page. You can also connect a channel in one click from the a
- `web/src/lib/billing/plans-page-copy.ts:234` sentence: `We email ${RENEWAL_REMINDER_DAYS} days before every yearly renewal, to the address you gave Stripe at checkout. Cancel in one click; the plan runs to the end of the period you paid for.`,
- `web/src/lib/guides/copy.ts:188` glance: 'At a glance',

### grade-hype (2)

- `apps/shared/copy/shared.json:120` "shared.standing.graded": "{spec} is graded {grade} this week",
- `web/src/content/changelog/2026-09-20-simulator-tools-are-live.md:18` None of it is graded yet. [Spec support](/sim/specs) still reads "Not yet" for every damage spec — nothing has been checked against a real parse, because there are no real parses yet. And Droptimizer can only rank what t

### explain-obvious (1)

- `web/src/lib/score/ratings-page.ts:70` "A guild's officers can change the weights and the cap for their own guild. This page shows the defaults.",

### allowlisted

- exclamation `web/src/content/guides/mage/fire.md:101`
- exclamation `web/src/content/guides/mage/index.md:36`
- exclamation `web/src/content/guides/priest/holy.md:84`
- exclamation `web/src/content/guides/priest/index.md:36`
- exclamation `web/src/content/guides/priest/shadow.md:99`
- exclamation `web/src/lib/guides/race-table.ts:24`
- exclamation `web/src/lib/guides/race-table.ts:61`
- exclamation `web/src/lib/guides/race-table.ts:86`
- exclamation `web/src/lib/guides/race-table.ts:105`
- em-dash `web/src/lib/guides/rotation-lists.ts:85`

## Judgment pass (lines grep did not flag)

Every copy module and content file in the sources was read as ai-slop-reviewer would. Each line below was filler, meta-talk about the site, explained the obvious, or used a generated cadence. All are fixed on this branch. `old -> new`, or `deleted`.

### Character page (the owner's flag)
- `web/src/lib/character/page-view.ts` "Next: spend your 1 talent point, then get Barbaric Iron Helm. Open the planner ›" -> two plain facts, each a link: "1 talent point unspent ›" (planner) and "Upgrade: Barbaric Iron Helm, <source> ›" (best in slot). "Nothing waiting." is deleted, so the line is left out when nothing waits.
- `apps/shared/copy/shared.json` `shared.standing.graded` "{spec} is graded {grade} this week" -> "{spec} tier {grade}" (still links to /tiers). `shared.standing.place` "{spec} is #{place} of {of} this week" -> "{spec} #{place} of {of} on the tier list". iOS and Android strings were regenerated.

### Web copy modules
- `sim/copy.ts` specValidatedNote "… Numbers from this spec are trustworthy." -> "Within 5% of the top 50 parses."
- `sim/copy.ts` specInProgressNote "…as a direction, not a figure." -> "…as a direction."
- `sim/copy.ts` gearDescription / talentsDescription: dropped "actually".
- `sim/copy.ts` weightsDescription "…, with the caveat that comes with it." -> "What one point of each stat is worth for your damage spec."
- `sim/copy.ts` fightStyleOptions.dungeon / dummy: dropped "a fixed schedule, not a repeating pull" and "a plain damage check, not a boss fight".
- `sim/copy.ts` emptyPrompt, dummyNote, targetLevelHelp: deleted (unused).
- `sim/copy.ts` noscript line "…runs the engine in your browser rather than on our servers. The simulator says the same." -> "`${tool}` runs in your browser and needs JavaScript."
- `sim/result-copy.ts` nextStep / nextStepLink / nextStepHref: deleted (unused "Next:" line).
- `sim/pool-quality-copy.ts` "…click “Add this build to the comparison” first…" -> "The build in the editor is not in the comparison yet. Add it first, or RUN will leave it out."
- `report/glossary.ts`: deleted three sentences that explain the obvious ("A damage-over-time debuff at 95% uptime was on the boss nearly the whole fight.", "A cancelled hardcast is time spent for nothing…", "A rage or energy bar at its cap is generating nothing…").
- `report/format.ts` "No parse for trash; parses are made for boss fights." -> "No parse for trash." and "Damage taken has no parse: taking more is not doing better" -> "Damage taken has no parse".
- `score/copy.ts` "Every upgrade for your spec moves it up." -> deleted.
- `score/ratings-page.ts` "…broad coverage is what counts." -> deleted. "This page shows the defaults." -> "These are the default weights and cap. A guild's officers can change them for their own guild."
- `tiers/tier-copy.ts` "Gap to the top is less healing in this fight, not a verdict on the class." -> deleted.
- `billing/plans-page-copy.ts` "Everything a player needs is free…" -> "The free list never shrinks. …". "Cancel in one click" -> "Cancel in Manage billing". The no-ads answer was cut to "No, for anyone, paying or not. Premium and the guild plans pay for the site." and ", and we would want it that way regardless" was deleted.
- `billing/checkout-guild-copy.ts` "Here is the next step for each of your guilds." / "Each card below has its next step." / "Here is where each of your guilds stands." -> deleted.
- `billing/copy.ts` "You're all set." -> "Payment confirmed."
- `bis/copy.ts` "--" pseudo dashes -> colons. "The picks below are still real sim results." -> "The picks are still ranked by the sim." The unused "Hover or tap an item for its stats" is deleted.
- `bis/landing-copy.ts` description "The best gear you can wear at every level range, from the simulator: …" -> deleted (tell 23, unused).
- `guides/copy.ts` "At a glance" -> "Summary".
- `items/item-sources.ts` "Place — Boss" -> "Place: Boss". The reputation line reads "Argent Dawn: Exalted".
- `app-beta/copy.ts` "…, in your pocket." -> "Raid sign-ups and loot when you are away from the PC."
- `addon/copy.ts` screenshot alt no longer quotes "Next: Lethal Attacks".
- `home-landing-copy.ts` + `CharacterSearchForm.astro` "Type the name, then pick the region and ruleset." -> deleted.

### Web pages and components
- `pages/b-unavailable.astro` "…In the meantime, the build planner works without it." -> "The Forever Sixty API is not answering. The build is not lost: the link works again once the API is back."
- `pages/sim.astro` noscript -> "The simulator runs in your browser and needs JavaScript. The build planner works without it."
- `components/report/RaidCooldowns.svelte` "Hover or tap a lane for …" -> deleted.
- `components/report/MechanicsMode.svelte` "A swing lands on whoever holds the enemy, so no table judges it; …" -> "No table judges melee; the Damage taken tab splits it by source."

### Guides (web/src/content/guides)
- Every spec opener in the form "<Spec> is the <Y>: …" (tell 23) now starts on the play itself. Fury: "Berserker Stance all fight, Bloodthirst and Whirlwind on cooldown, …". The same applies to Arms, Protection, Balance, Beast Mastery, Holy priest, Retribution, Assassination, Combat, Subtlety, Elemental, Enhancement and Restoration shaman.
- "[Open the planner](/planner?class=…) to change the build." (12 guides) -> deleted (site meta talk).
- "The full lists by level are on the … best in slot page" / "The full list … is on the … BiS page" (6 guides) -> deleted.
- "Enchanting is the profession your best in slot list asks for." -> "Your raid set asks for Enchanting." Similar race and profession "X is the pick:" lines now lead with the choice ("On the Horde, take Orc for Blood Fury: …").
- `priest/shadow.md` "See the rotation for what it costs you." -> deleted. `rogue/subtlety.md` "If you level Subtlety, here is how it starts." -> deleted. `druid/balance.md` "You do not need to do anything about it." -> deleted.
- `paladin/retribution.md` "a new Seal simply replaces the old one" -> "replaces". `warrior/protection.md` "unlocks Shield Slam" -> "opens Shield Slam".

### Changelog and legal pages
- Em dashes in privacy, refunds and terms became a semicolon, colon, full stop or parentheses. Legal meaning is unchanged, and the asserted phrases are untouched.
- `2026-10-27-discord-bot.md` "The Forever Sixty Discord bot is here." -> "Raid sign-up buttons and slash commands in your guild Discord."
- `2026-10-28-log-coach.md` "The log coach is here." -> "Premium log coach: the casts, buffs and consumables to fix first."
- `2026-10-27-guild-raids.md` "sign up in one tap" -> "sign up on the raid page".
- Three "The point is …" / "The distinction matters because …" / "Both follow from the same rule." filler sentences -> deleted.

### Apps (apps/shared/copy/en.json)
- `bis.next` / `bis.nextSame` "Next: …" -> "Runner-up: …".
- `gear.intro`, `talents.legend`, `plan.cmp.legend`, `roster.footer` and `report.footer`: deleted the "Tap X for Y" sentences.
- `home.answer.noGearNote`, `raids.noGuild.body` ("in one tap") and `raid.cancelledYou` ("Nothing else to do."): filler deleted.
- `signin.pickClassNote`, `home.next.talentsNone`, `talents.pickSpec`, `bis.tankLine`, `devices.footer`, `reports.footer` and `widget.gallery.upgrade` were tightened. `guides.glance` "At a glance" -> "Summary".

### Addon (Locale.lua)
- `talentStripNextTalent` "Next: %s" -> "%s" (the strip reads "Endurance Training (1 of 5)"). `dialogNowFollowingNext` "Now following %s. Next: %s." -> "Now following %s. Take %s." `dialogNextStep` "Your next step: %s" -> "%s".
- `diagSiteNever` escaped em dash -> full stop. "Click to open or close." (2 hints) -> deleted. "by our (stat) weights" (2 lines) -> deleted.
- Seven `specPlaystyle` hype tails replaced with game facts. Examples: "Bestial Wrath makes it deadly" -> "Your pet fights, Bestial Wrath", and "Shield and plate, Shield Slam, holds attention" -> "Shield tank, Shield Slam".
- `CHANGELOG.md` 2.0.16 no longer quotes "Next: …".

### API (player-facing)
- `coach/copy.go` ThinPool "not enough logs yet — …" -> "Not enough logs yet: …". RulesOpening "Here is what your log shows for {Boss}." -> "Your {Boss} pull, from the log."
- `notify/report_embed.go` report_ready description "Kill times, loot and the biggest changes on Forever Sixty." -> deleted (it restated the fields below it).
- `companion/ui/copy.js` and `companion/internal/app/copy.go`: read and clean, so nothing ships in the next companion release.

## Left for other lanes (not edited here)

### Discord bot (lane f-discord; api/internal/discordbot)
- `copy_coach.go:20,22` "Fix first: {stem} — on {n} of {m} bosses." -> "Fix first: {stem}, on {n} of {m} bosses." (already fixed on origin/lane/f-discord)
- `copy_coach.go:25` "Coach reviews for tonight's pulls are ready; open the report or use /report." -> "Coach reviews for tonight's pulls are in the report and /report." (this also lives in `notify/report_embed.go` CoachSentence and the raidNight row of design/specs/2026-10-28-log-coach.md, so change all three together)
- `copy_guide.go:75` "Choose a class, a spec and a range, then Show BiS. Link Discord and this starts from your character." -> "Link Discord to start from your character."
- `copy_officer.go:86` "The readiness board is part of the Raid guild plan. An officer can move the guild to the Raid guild plan." -> "The readiness board is part of Raid guild."
- `commands.go:60` "A character: item level, readiness and the top 3 upgrades." -> "Item level, readiness and top 3 upgrades."
- `commands.go:69` "The last logged raid: kills and loot." -> "Kills and loot from the last logged raid."

### Guild (lanes f-guild and f-guild-onboard; web/src/lib/guild, web/src/components/guild, api/internal/guilds)
api/internal/guilds: no slop found. Web, most damaging first:
1. `lib/guild/copy.ts:35` "A guild's progression per boss, pull counts, kill dates and roster bests. Its roster, raid nights and loot sit in the tabs once it has loaded." -> delete (tell 23; renders twice in Guild.svelte)
2. `components/guild/GuildRoute.astro:21` "A World of Warcraft: Forever guild: progression per boss, …" -> "Progression per boss, pull counts, kill dates and roster bests."
3. `lib/guild/overview-role-view.ts:228-230` "Ready to raid" / "Almost there" / "Getting started" -> "Ready for Thursday" / "2 to fix before Thursday"
4. `lib/guild/overview-role-view.ts:245-246` "Almost there: {n} things before {day}." -> "{n} to fix before {day}."
5. `lib/guild/raid-copy.ts:200` "…members sign up in one tap, from the raid page or the guild page." -> delete (unused key, stock promise)
6. `lib/guild/attendance-copy.ts:160` "Added attendance and raid-night reports. …" -> delete (unused, meta)
7. `lib/guild/copy.ts:52-54` "You're the only member the site knows about. Share the invite link…" -> delete (unused)
8. `lib/guild/copy.ts:166` "The site knows a character of yours in {guild}." -> "You have a character in {guild}."
9. `lib/guild/copy.ts:169` "The site does not know a character of yours in this guild. …" -> "None of your characters is in this guild. Add one on Get set up."
10. `lib/guild/copy.ts:23` "Spelled right? Then it is not on Forever Sixty yet. …" -> "Not on Forever Sixty yet. It shows up once a member imports a character or syncs the addon."
11. `lib/guild/recruit-copy.ts:154` "Same units as the BiS page" -> delete
12. `lib/guild/overview-role-view.ts:124` "From your gear and buffs when you last synced." -> delete
13. `lib/guild/overview-role-view.ts:144` "The BiS list shows the swaps." -> delete
14. `lib/guild/readiness-view.ts:195` "Best in slot or no upgrades, enchanted, stocked, every point spent" -> "BiS, enchanted, flask and food, talents spent"
15. `lib/guild/discord-copy.ts:15` "Discord asks for a server and a text channel, not a voice channel, then lets Forever Sixty post…" -> "Pick a server and a text channel. You need Manage Webhooks there. We store only the webhook, encrypted."
16. `lib/guild/discord-bot-copy.ts:36` "…Discord asks which server; you need Manage Server in it." -> "Raid posts get In, Late, Bench and Out buttons; /raid, /bis and /char work in the server. Needs Manage Server."
17. `components/guild/GuildRosterTable.svelte:753` "The addon saw this character in the guild; no account on the site has claimed it. …" -> "In the guild per the addon; no account yet."
18. `lib/guild/copy.ts:208` "Copy it now. For safety we keep only a fingerprint, so it cannot be shown again after you leave this page." -> "Copy it now. It cannot be shown again."
19. `lib/guild/copy.ts:215` "Use this link to join our guild on Forever Sixty." -> delete (repeats the share title)
20. `lib/guild/copy.ts:159` "…Claim the guild first; then the plan is one step." -> "…Claim the guild first, then subscribe."
21. `lib/guild/progression-view.ts:17` "No pulls logged yet. Upload your first raid night and it shows here." -> "No pulls logged yet."
22. `lib/guild/progression-view.ts:22` "…Progression starts with your first logged pull." -> delete that sentence
23. `components/guild/GuildRaids.svelte:250` "No raid nights logged yet. Upload a raid log to get started." -> "No raid nights logged yet."
24. `lib/guild/raid-copy.ts:60-61`, `lib/guild/week-one.ts:240`: "Nothing planned yet. Officers post …" -> "Nothing planned yet." / "Nothing on the calendar yet."
25. `lib/guild/assignments-view.ts:23` "…Its jobs show here once the client carries them." -> delete that sentence
26. `lib/guild/roster-gate.ts:22` "Your character is on this guild, waiting for an officer to verify it. …" -> "Waiting for an officer to verify your character."
27. `lib/guild/raid-copy.ts:162` "Markdown: bold, lists and links. Members see this on the raid page." -> "Markdown: bold, lists, links."
28. `lib/guild/overview-view.ts:91` "{n} raiders with gear on the site" -> "{n} raiders with gear shared". `lib/guild/roster-facts.ts:88` "{spec} on the site" -> "{spec}".
29. `lib/guild/roster-best.ts:29` "Top {unit} in any {encounter} kill" -> delete (the tooltip restates its label)
30. `lib/guild/readiness-view.ts:131`, `lib/guild/standing.ts:53` "…every readiness check passes." -> "{name}: ready."
31. `lib/guild/recruit-copy.ts:172-176` three "officers only" notes -> keep one: "Officers only"
32. `lib/guild/loot-council-view.ts:29` "Tell the officers how you want it. Only officers see it." -> "Only officers see your call."
33. `lib/guild/discord-bot-copy.ts:95` "Discord did not answer, so the post is tried again on its own." -> "Discord did not answer. Retrying."
34. `lib/guild/raid-copy.ts:153` "Posted to {channel}. The post updates itself." -> "Posted to {channel}. Sign-ups update live."
35. `lib/guild/copy.ts:126` "Nobody has claimed this guild. Claiming lets you approve members, …" -> "Nobody has claimed this guild."
36. `lib/guild/copy.ts:331` "Seen at {clock} on the member's clock" -> "Seen {clock}, their time"
37. `lib/guild/week-one.ts:34` "Your first logged night starts the progression board." -> delete
- About 15 minor lines remain. Error punctuation is inconsistent ("That did not save; try again." vs ". Try again.", which should be one shared constant). Other examples: the `recruit-copy.ts:40` placeholder tricolon, the `overview/BuffCover.svelte:34` meta line and the `overview/MemberChecklist.svelte:12` "Private to you" qualifier.

### Not copy, noted
- `docs/contracts/fixtures/app/cdn.changelog.json` still has the old changelog text "connect a channel in one click". It is a contract fixture.
- `design/mocks/data/*.json` and `design/mocks/web/*` still quote the old rating and changelog lines (mock data).
- Possible fact problem, left as is: `score/ratings-page.ts` and `rating/copy.ts` list "world buffs at the pull" under Preparation, but a `sim/copy.ts` comment says Forever has no world buffs.
