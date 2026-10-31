# Theorycraft audit: hybrids (2026-10-31)

This audit covers Druid (Balance, Feral, Restoration), Shaman (Elemental, Enhancement, Restoration) and Paladin (Retribution, Protection, Holy). It compares public Forever theorycraft with what our sim produces.

**What was compared:**
- Engine pin `816ec57c0`, read with `git show 816ec57c0:<path>` in `/Users/jh/code/wowsims-forever`, not from the checkout.
- Active client 1.60.1.70658.
- Our side: the APLs in `data/curated/apl/`, the guide builds in `web/src/content/guides/`, and the level-60 alliance raid bands in `data/builds/1.60.1.70658/bis/`.
- **The bands are stale.** They were generated on 2026-10-31 by engine `1f542def3`, which is older than caster round 2. Elemental Mastery and Eclipse are not in their talents yet.

**How claims were checked:**
- Client tables were fetched from wago for 1.60.1.70658: `SkillLineAbility`, `SpellEffect`, `Spell` and `SpellName`. `SpellName` was also fetched for 1.60.1.70707, the newest Forever build on wago.
- Wowhead's live Forever talent payload (`nether.wowhead.com/forever/data/talents-classic?dv=100`) was fetched the same day.

**Sources and their limits:**
- No source has a tested level-60 build. The beta was capped at 30 and there are no level-60 logs.
- Every outside build below is a projection or a simulation: kami-labs, SixtySim, wowforeverprogress, Endgame Tools.
- Warcraft Tavern's spec guides cover the level-30 beta only.
- Icy Veins and Mobalytics refused direct fetches (403). Mobalytics was read through a search index.
- Single-source claims are flagged as such.

**Classification:**
- CONFIRMED-BUG: our sim is wrong, with evidence attached.
- FOREVER-CHANGE: the web source describes Era only, and the client shows Forever differs, so we are right.
- OPEN: cannot be settled yet; each says what would settle it.

## Cross-cutting findings (affect more than one spec)

| # | Finding | Class | Evidence | Fixed |
|---|---|---|---|---|
| X1 | **The raid preset grants Sanctity Aura (+10% Holy damage). No Forever paladin can learn it.** It inflates every alliance paladin's Holy damage: seals, Judgement, Holy Strike, Consecration, Exorcism. | CONFIRMED-BUG | Spell 20218 has no `SkillLineAbility` row in 70658. It is absent from `trainables/paladin.json`, from the 70658 Retribution tree, and from Wowhead's live trees. `data/curated/presets.json` lists `sanctity_aura`. `sim/request/request.go:725` sets `Aura: PaladinAura_SanctityAura` for Retribution. The engine applies ×1.1 Holy (`sim/core/buffs.go:463`). | 40b761b36, 448d5a096 |
| X2 | **The raid preset leaves out Blessing of Kings because "the live paladin trees do not grant it". That reason is false: Kings is a baseline trainer spell in Forever.** | OPEN | `SkillLineAbility`: 20217 Blessing of Kings (skill line 267, AcquireMethod 0, level 20) and 25898 Greater Blessing of Kings (level 60). Both are listed in `trainables/paladin.json`. Its effect is aura 137, +10% total stats. That is the same aura and amount as Heart of the Lion (409583), which the preset already carries. **Settle:** do Kings and Heart of the Lion stack? `SpellEffect` cannot say; a beta character sheet or a Blizzard note can. If they stack, every spec is missing about 10% stats, which shifts tank and melee more than casters. | Open (note corrected in 40b761b36) |
| X3 | **Wowhead's live talent payload (fetched 2026-10-31) has reworked hybrid trees that no client on wago contains.** | OPEN (high impact) | Changes include:<br>- Retribution's capstone is **Unsealed Wrath**; Twist of Light is gone. wowtbc.gg independently names "Unsealed Wrath".<br>- Conviction is 6% crit to spells and attacks; Instrument of Law is merged into Champion of the Light; Holy Conduit moves to Holy.<br>- Divine Intellect and Toughness move to Protection; Holy adds Oath of Light and Boundless Light; Protection's capstone is **Shieldsworn**.<br>- Feral loses **Berserk** and Improved Shifting Power. Primal Bite and Shifting Power become 31-point talents. Lacerate and Predatory Discipline are new talents. Heart of the Wild moves to Restoration and Furor to Feral.<br>- Restoration gains **Bloom**.<br>The new spell ids (1328544–1328621) are absent from `SpellName` in 70658 and in 70707, the newest wago build (2026-10-30). WFDB (client "70707", undated) reports the same Feral change; it is a single source. **Settle:** this is either a hotfix or a launch build. A `DBCache.bin` from a current client, or the next wago build, decides. If it lands, every Retribution, Protection, Feral, Balance (Heart of the Wild) and Restoration build in this report must be re-searched. | Open: no newer client exists (see Fixes) |
| X4 | **Guide `statPriority` front matter disagrees with our own published weights.** For example, Elemental lists Intellect first, while its weights put spell power at 2.69× Intellect. Retribution lists attack power second, below Strength but above crit and hit; its weights put attack power last. | CONFIRMED-BUG (presentation only, no DPS impact) | `web/src/content/guides/{shaman/elemental,paladin/retribution,druid/feral,shaman/enhancement}.md` against the `weights` in `bis/<spec>.json`. The lists were written on 2026-09-24. | - |

## Druid: Balance

**Outside consensus:**
- **Build:** about 40/0/11 (wowforeverprogress 41/0/10; SixtySim 38/2/11 to 40/0/11).
- **Rotation:** Moonkin Form, keep Moonfire and Insect Swarm up, then one Wrath into two Starfires through Eclipse. Warcraft Tavern, Icy Veins, SixtySim and wowtbc.gg agree.
- **Stats:** spell damage > hit > crit (Warcraft Tavern). SixtySim notes casters sit far below the hit cap.
- **Position:** SixtySim puts Balance at 62.7% of the top spec (19th of 25).

**Ours:**
- **Build:** the guide raid build `5232220115501351/0/5053` is 38/0/13.
- **Rotation:** Moonfire and Insect Swarm, then Wrath whenever no Eclipse charge is banked, then Starfire.
- **Weights:** spell power > arcane power > hit > crit > Intellect.
- **Position:** about 62% of Fury after round 2.

**Rotation, tier position and weights order are aligned.**

| Difference | Class | Evidence | Fixed |
|---|---|---|---|
| **Omen of Clarity is not modeled at all.** It is a baseline trainer spell in Forever, and Moonkin Form doubles its proc chance. The free spell matters for mana: casters run out in long fights. | CONFIRMED-BUG | Client: `SkillLineAbility` 16864, skill line 574, AcquireMethod 0, level 20. Spell 16864's text: "Your spells and attacks have a chance to grant you Clearcasting… not consumed by Wrath". Moonkin Form (24858) effect 4 is +100% to Omen of Clarity's chance. Engine: `sim/druid/talents.go` `applyOmenOfClarity` has its whole body commented out, still gated on a vanilla talent Forever does not have. Endgame Tools (10-16) also says Omen of Clarity is baseline. | 317a8d7ec (fork), e4071215b |
| The build spends 8 points that do nothing in this sim (Nature's Focus 5, Subtlety 3) to reach Naturalist. | Aligned | Naturalist's +5% to all damage costs 10 points. The outside 11-point Restoration splits do the same. | - |
| Starfire rank 7 is an Ahn'Qiraj book rank. | OPEN | Carried over from round 2. **Settle:** do Forever trainers teach it? | - |
| Wrath and Starfire base damage: Endgame Tools says Wrath went up about 50%; SixtySim says both were cut. | FOREVER-CHANGE | Round 2 rebuilt both from client base damage and coefficients. Our numbers stand. | - |

## Druid: Feral (cat)

**Outside consensus:**
- **Build:** 0/38/13 (wowforeverprogress), 9/33/9 (SixtySim), 9/40/2 (Warcraft Tavern).
- **Rotation:**
  - Keep Rake up; it feeds Rend and Tear.
  - Rip at 5 combo points, Ferocious Bite at 5 while Rip is up.
  - Shifting Power at 50 energy or below; Shred as the filler.
  - Tiger's Fury, King of the Jungle and powershifting are gone (Endgame Tools, Warcraft Tavern, Icy Veins, WFDB).
- **Stats:** Strength > Agility (SixtySim, WFDB). Three sources say cat damage now scales with weapon DPS (Icy Veins, Mobalytics, SixtySim).
- **Position:** 73% of the top spec (SixtySim).

**Ours:**
- **Build:** the band build is `0500200000000000-45211031021032212001-5053000000000000`.
- **Rotation:** Shifting Power, **Tiger's Fury**, Rip at 5 under 6 s, Bite at 5, Shred. Rake was dropped by an earlier search.
- **Weights:** Strength 1.06 > Agility 1.00 > hit > crit.
- **Position:** about 77% of Fury.

| Difference | Class | Evidence | Fixed |
|---|---|---|---|
| **Omen of Clarity is missing.** It is a baseline Forever spell, and for a cat it means free Shreds. This is likely the largest single gap in this report, about +5–15%; measure it. | CONFIRMED-BUG | Same client and engine evidence as Balance. `sim/druid/feral/apl_values.go:112` also has the Clearcasting check commented out. | 317a8d7ec (fork), e4071215b |
| **Rend and Tear does nothing for a cat.** "Increases damage done by your melee abilities on Bleeding targets by 2/4/6/8/10%". Rip keeps the target bleeding, yet the engine applies the talent only to bear abilities. The build skips it as a result, and its 31 points include zero-damage talents: Feral Swiftness 2, Feral Instinct 1, Brutal Impact 1. | CONFIRMED-BUG | Engine `sim/druid/feral_talents.go` around lines 80–105 says "the Cat Form abilities (claw.go, rake.go, rip.go, shred.go, ravage.go and ferocious_bite.go) do not read it yet". Its bleed check looks only at Lacerate or `AssumeBleedActive`, never Rip or Rake. Client text: talent node, spell 1223246. Up to +10% on Shred and Ferocious Bite once the build takes 5/5. | 317a8d7ec (fork), e4071215b |
| **Tiger's Fury is cast, but a Forever druid cannot learn it.** | CONFIRMED-BUG (about −1 to −2%) | `SkillLineAbility` 5217 has AcquireMethod **3**, a Season of Discovery rune, not learnable (see `forever-trainables-source`). It is absent from `trainables/druid.json`. Engine `sim/druid/tigers_fury.go` registers it at level 24; the APL casts it every 30 s. Four web sources say it was removed. | 317a8d7ec (fork), e4071215b |
| Weapon DPS scaling: the engine fixes the cat's paw at 43.84–65.76 dmg, 1.0 s, and ignores the equipped weapon (`sim/druid/forms.go:31`). The old feral-attack-power code is commented out. Shred is 80% weapon damage + 155 (`SpellEffect` 9830). | OPEN | Three sources say Forever cat damage scales with weapon DPS. Mobalytics says 1 weapon DPS ≈ 4 AP. The tables do not say whose weapon "weapon damage" means in Cat Form. **Settle:** a beta log or tooltip of Shred damage with two different weapons. If true, cat gear, BiS and weights all change. | - |
| Rake is not kept up. | OPEN | Re-measure once Rend and Tear applies to cats; until then Rake has no bonus to feed. | - |
| Berserk and Improved Shifting Power are in our build and tree. | OPEN | See X3: the live Wowhead tree removes both. | - |

## Druid: Restoration

**Outside consensus:**
- **Build:** 11/0/40 with Wild Growth (wowforeverprogress); no other source has a level-60 build.
- **Priority:** Swiftmend; Rejuvenation and Regrowth on the tank; Wild Growth for party damage; a Healing Touch rank as the filler.
- **Stats:** healing power > Spirit > haste > crit > Intellect (Warcraft Tavern).
- **Position:** SixtySim puts it 2nd of 5 healers.

**Ours:** the band build is `4532200100000000-0-5050035153113200`, with no Wild Growth. Weights: mp5 > Intellect > Spirit > healing power.

| Difference | Class | Evidence | Fixed |
|---|---|---|---|
| Wild Growth (31 points) is skipped, though 34 points sit in Restoration. | OPEN | The guide says the sim measured Wild Growth against those points elsewhere and the points won under our heal profile (180 s, `data/curated/heal-profile.json`). Raid party healing is where Wild Growth earns its place. **Settle:** a heal profile with raid-wide or party damage. | - |
| Omen of Clarity is missing (free heals). | CONFIRMED-BUG | As for Balance. For a healer it is mana only. | 317a8d7ec (fork), e4071215b |
| Mana-heavy weights (mp5 at the top) against the outside healing-power-first order. | OPEN | This follows from the heal profile's mana pressure. Settle with the profile above. | - |

## Shaman: Elemental

**Outside consensus:**
- **Build:** kami-labs `5501331020011351-055000001-05301` is a projection.
- **Rotation:** keep Flame Shock up; Lava Burst on cooldown; Lightning Bolt as the filler; Chain Lightning on 2 or more targets; fire totem down; Elemental Mastery with potions. Warcraft Tavern and kami-labs agree.
- **Stats:** spell damage > crit > hit > Intellect (Warcraft Tavern only).
- **Position:** kami-labs' sim has Elemental at 68% of its Retribution figure.

**Ours:**
- **Build:** the guide raid build `5510331300001351/052/053302`, with round 2's Elemental Mastery.
- **Rotation:** the same as the outside rotation.
- **Weights:** spell power > nature power > hit > Intellect > crit.
- **Position:** about 57% of Fury.

**Rotation, Elemental Mastery and Lava Burst are aligned after rounds 1 and 2.**

| Difference | Class | Evidence | Fixed |
|---|---|---|---|
| 5 points sit in zero-damage Restoration talents while Ancestral Knowledge is at 2/5 and Shamanistic Focus at 0. | CONFIRMED-BUG (minor, about 0.5–1%) | The zero-damage points are Natural Grace 3 (threat) and Improved Reincarnation 2. Kami-labs takes Ancestral Knowledge 5/5 and Shamanistic Focus 1/1 in Enhancement instead. Talent search should be rerun with those two talents removed from its keep list. | 943646d1c (Natural Grace; Improved Reincarnation held) |
| Fire Nova on Clearcasting (kami-labs only). | OPEN | Our APL notes measured Fire Nova losing to Lightning Bolt at every mana threshold, but never on a free Clearcasting proc. Single source. | - |
| Elemental Mastery: Blizzard's Oct 7 post says removed; Conquest Capped and kami-labs say reworked. | FOREVER-CHANGE | The client has 16166 as the capstone, and round 2 modeled the client's version. | - |

## Shaman: Enhancement

**Outside consensus:**
- **Build:** kami-labs `453003102-055030031005102251` (projected) takes Maelstrom Weapon 5/5 and Improved Stormstrike 2/2.
- **Rotation:**
  - Windfury Weapon is the imbue (Warcraft Tavern, wowtbc.gg; kami-labs uses Flametongue).
  - Totems down, Rage of the Farseer with on-use effects, Stormstrike on cooldown.
  - **Lightning Bolt at 5 Maelstrom stacks, timed into the Stormstrike debuff.**
  - Flame Shock, else Earth Shock.
- **Weapon:** a slow two-hander.
- **Stats:** weapon > hit > Strength > crit.
- **Position:** noobtoboss puts it in S tier; kami-labs' sim ranks it below its own Retribution figure.

**Ours:**
- **Build:** the band build `3230031000000000-255030031005102031-0530000000000000` takes Maelstrom 3/5, Improved Stormstrike 0 and Earth's Grasp 2 (no damage).
- **Rotation:** Strength of Earth, Windfury Totem, Stormstrike, Searing Totem, Earth Shock. **No Lightning Bolt.**
- **Weapon and imbue:** a two-hander with Windfury Weapon, aligned.
- **Position:** about 80% of Fury.

| Difference | Class | Evidence | Fixed |
|---|---|---|---|
| **Maelstrom Weapon is taken and modeled, but the APL never casts Lightning Bolt.** The stacks are wasted, and the talent search under-values Maelstrom and Improved Stormstrike as a result. Estimated +4–6% with Maelstrom at 5/5. | CONFIRMED-BUG | Engine `sim/shaman/talents.go:782` `applyMaelstromWeapon`: 50% chance per melee hit, 5 stacks; per-rank cast and cost cuts go to Lightning Bolt. Stormstrike's +20% applies to the next Lightning Bolt. `data/curated/apl/shaman-enhancement.json` has no 15208 line. Even our own guide says Maelstrom "lets Enhancement weave in a ranged nuke". Three outside sources put Lightning Bolt at 5 stacks in the priority. | 943646d1c |
| Imbue: Windfury against Flametongue. | OPEN (minor) | Two sources say Windfury, one says Flametongue. Our kit's Windfury is the majority call; one sim comparison would settle it. | - |

## Shaman: Restoration

**Outside consensus:**
- **Build:** level 30 only (Warcraft Tavern); key talents are Mindfulness, Water Shield, Mana Tide, Riptide.
- **Priority:** Riptide, then Chain Heal on the Riptide target, then Lesser Healing Wave and Healing Wave.
- **Stats:** healing power > Spirit > Intellect > crit.

**Ours:** the band build is `5300000000000000-0-5532003315513151`, with Riptide.

| Difference | Class | Evidence | Fixed |
|---|---|---|---|
| 8 points go to Convection 5 and Concussion 3, which only affect shocks and Lightning Bolt. Enhancement's Thundering Strikes 5 (+5% crit to all spells, heals included) and Ancestral Knowledge 3 are open. Natural Grace 2 (threat) is also dead weight. | CONFIRMED-BUG (build, about +2–3% healing) | Client talent text (70658 shaman tree). Crit and Intellect both carry positive weight for this spec (crit 0.44, Intellect 1.00). | 943646d1c, 448d5a096 |

## Paladin: Retribution

**Outside consensus:**
- **Build:** kami-labs `52003-503-05225231001330321`, published 14 Sep and simmed 26–28 Oct. It takes **Precision 3/3** (behind Toughness 5) and **Twist of Light**.
- **Rotation:** open with Seal of the Crusader and judge it; Holy Strike and Judgement on cooldown; Hammer of Wrath under 20%; alternate Seal of Command and Seal of Righteousness for Twist of Light's Echo; Consecration above about 70% mana. All three sources agree.
- **Position:** kami-labs puts Retribution at 84% of Fury; tier lists rank it from A to C.

**Ours:**
- **Build:** the band build `55003003100100000-0000000000000000-05225331001330320` has no Precision and no Twist of Light. It takes Seal of Command but never casts it.
- **Rotation:** Seal of Righteousness only (hardcoded), Judgement, Exorcism, Hammer of Wrath, Holy Strike, Consecration above 65% mana.
- **Position:** 73.5% of Fury, which includes the false Sanctity Aura from X1.

| Difference | Class | Evidence | Fixed |
|---|---|---|---|
| **Sanctity Aura counted (X1). Retribution is overstated by roughly 4%** (+10% on its Holy share of damage). | CONFIRMED-BUG | See X1. | 40b761b36, 448d5a096 |
| **The build skips Precision.** The engine models it (`sim/paladin/protection/talents_test.go:72`). Our weights value 3% hit at about 24 DPS, about +4% on a 570 band (2.11 scale factor × 0.2666 DPS per Strength × 14 rating per 1% × 3). The build instead spends 10 Holy points on Divine Intellect 5, Reverence 3, Purifying Power 1 and Divine Favor 1. Purifying Power only shortens the Exorcism and Holy Wrath cooldowns, and Exorcism cannot land on our Humanoid boss. Divine Favor is a heal. | CONFIRMED-BUG (talent choice; verify with a talent-search rerun) | Hit to cap: 9% for special attacks (band `hit_to_cap`). Kami-labs takes Toughness 5 + Precision 3. | 40b761b36 |
| **The APL's notes are stale on Twist of Light.** They say Twist of Light "is never read anywhere", but the pinned engine implements it (`sim/paladin/talents.go:415`, the hook at `paladin.go:357`). No seal-twist rotation has been measured since, so the talent search never takes Twist of Light. | OPEN | Measure a Seal of Command / Seal of Righteousness alternation with Twist of Light. But X3: the live tree replaces Twist of Light with Unsealed Wrath, so settle X3 first. | - |
| Seal of Command is talented but the rotation always casts Seal of Righteousness (`sim/request/request.go:724`). | OPEN | In Era, a slow two-hander favors Seal of Command. **Settle:** measure Seal of Command as the primary seal. | - |
| Judgement of the Crusader debuff, and Seal of the Crusader as the opener. | OPEN | Not in the raid preset debuffs. `SkillLineAbility` lists Judgement of the Crusader ranks with AcquireMethod 3, which is ambiguous for a triggered spell. **Settle:** does Forever's Seal of the Crusader judgement still add Holy damage taken? | - |

## Paladin: Protection

**Outside consensus:**
- **Build:** kami-labs `50003-0530213301301551-50203` takes Reckoning 5, Conviction and Improved Judgement for threat.
- **Rotation:** Seal of Fury, Holy Shield, Holy Strike, Judgement, Consecration, Templar's Bulwark (Warcraft Tavern, wowtbc.gg, Gfinity). Kami-labs' sim uses Seal of Righteousness instead.
- **Stats:** hit to 3%, then defense and block, then Stamina and armor.

**Ours:** the band build `50003000000000000-5530513321301051-50000000000000000` takes Toughness, Anticipation and Sacred Duty for survival. The APL holds Seal of Fury and runs Holy Shield, Holy Strike, Hammer of the Righteous, Swift Judgement, Judgement, Exorcism and Consecration. **The rotation is aligned.**

| Difference | Class | Evidence | Fixed |
|---|---|---|---|
| Sanctity Aura counted (X1): it inflates threat and damage. | CONFIRMED-BUG | See X1. | 40b761b36, 448d5a096 |
| Survival talents against threat talents. | OPEN | Our tank score is a different objective from kami-labs' threat-per-second figure. | - |
| Shieldsworn and the tree rework. | OPEN | See X3. | - |

## Paladin: Holy

**Outside consensus:**
- **Build:** a 48/0/3 Benediction split (wow.gg, 19 Sep) or kami-labs' overview.
- **Priority:** Light's Vigil, then Holy Shock; Holy Light on Infusion of Light procs.

**Ours:** the band build is `05320003225111051-5532500000000000`.

| Difference | Class | Evidence | Fixed |
|---|---|---|---|
| 20 points in Protection do nothing for healing (Toughness, Redoubt, Precision, Guardian's Favor, Anticipation). | Aligned, negligible | Retribution's only healer talent is Benediction 5; any 20-point split carries about 15 dead points either way. | - |
| The Holy tree rework (Oath of Light, Boundless Light, Holy Power). | OPEN | See X3. | - |

## Ranked confirmed bugs (estimated DPS impact)

1. **Feral: Omen of Clarity not modeled.** Likely +5–15% for Feral; smaller, mana-only gains for Balance and Restoration Druid. Engine `sim/druid/talents.go` `applyOmenOfClarity`. **Fixed:** 317a8d7ec (fork), e4071215b.
2. **Feral: Rend and Tear ignored by cat abilities.** Up to +5–10% once the build takes 5/5. `sim/druid/feral_talents.go`. **Fixed:** 317a8d7ec (fork), e4071215b.
3. **Enhancement: the APL never casts Lightning Bolt,** so Maelstrom Weapon is wasted. Estimated +4–6% with a re-search to Maelstrom 5/5. **Fixed:** 943646d1c.
4. **Retribution: Sanctity Aura counted though not learnable.** About −4% for Retribution (also inflates Protection Paladin). `data/curated/presets.json`, `sim/request/request.go:725`. **Fixed:** 40b761b36.
5. **Retribution: Precision skipped.** About +4%; needs a talent-search rerun. **Fixed:** 40b761b36.
6. **Feral: Tiger's Fury cast though not learnable.** About −1 to −2%. **Fixed:** 317a8d7ec (fork), e4071215b.
7. **Restoration Shaman:** Convection and Concussion picked over Thundering Strikes and Ancestral Knowledge. About +2–3% healing. **Fixed:** 943646d1c, 448d5a096.
8. **Elemental:** dead Restoration points over Ancestral Knowledge and Shamanistic Focus. About +0.5–1%. **Fixed:** 943646d1c.

The biggest open item is **X3**: Wowhead's live trees are ahead of every wago client. Then **X2** (does Blessing of Kings stack with Heart of the Lion), then Feral weapon-DPS scaling.

## Fixes (2026-10-31)

Measured with `sim/cmd/spec-breakdown` on the level-60 alliance raid band (raid preset, or `-heal` against the heal profile), 10000 iterations at seeds 1 and 50001. Every error is under ±0.1%. Relative figures only. "Before" is engine 816ec57c0 with the origin/main rotations, presets and guide builds; "after" is engine 317a8d7ec with the new rotations and builds.

| Spec | Change | Before → after |
|---|---|---|
| Feral | Omen of Clarity baseline, Rend and Tear on cat abilities, Tiger's Fury removed (engine) | about +4.2% |
| Feral | Raid build takes Rend and Tear 5/5 (talent-search winner) | about +9.8% total |
| Balance | Omen of Clarity (mana only) | about +0.6% |
| Restoration Druid | Omen of Clarity on heals | about +4.0% effective HPS |
| Enhancement | Lightning Bolt at five Maelstrom stacks; Maelstrom 5/5 raid build | about +3.2% |
| Elemental | Natural Grace into Ancestral Knowledge | about +1.3% |
| Restoration Shaman | Thundering Strikes 5 and Ancestral Knowledge 5 for Convection, Concussion and Natural Grace | about +4.4% effective HPS |
| Retribution | Sanctity Aura removed | about −5.4% |
| Retribution | Precision 3/3 behind Toughness 5 | about +2.1% (net about −3.4%) |
| Protection Paladin | Sanctity Aura removed | about −5.5% DPS |

**How the builds were chosen.** Feral and Elemental are the talent-search winners that keep every unmodeled guide talent. The search cannot raise a talent past the guide's rank, nor take a no-damage prerequisite, so Enhancement's Maelstrom 5/5 and Retribution's Toughness-then-Precision were hand-built, then a search centred on each found no clean variant ahead. Elemental could gain about 3% more by also dropping Improved Reincarnation. That is held, because its +4% health is unmodeled and the adoption rule keeps unmodeled guide talents. Restoration Shaman was chosen with `sim/cmd/heal-builds`, because healers have no talent search. Rake upkeep was re-measured after Rend and Tear started working for cats: it still loses about 9% in the raid run, because Rip alone keeps the target bleeding.

**Omen of Clarity model.** The client row has a 10 s proc cooldown and a ProcChance of 100, which is not the server's rate.
- Melee keeps the 2 procs a minute the engine inherited, at weapon speed for yellow hits.
- Spells and heals use the same 2 a minute taken over the cast time, with an instant counted as the global cooldown. This rate is marked unconfirmed.
- Moonkin Form doubles the chance and halves the cooldown (24858 effects 4 and 5).
- Wrath neither benefits nor consumes it: it is outside Clearcasting's class mask.

**X2, Blessing of Kings: still open.**
- It is learnable: `SkillLineAbility` 20217 has AcquireMethod 0 at level 20, and 25898 (Greater Blessing of Kings) is level 60.
- The client cannot say whether it stacks with Heart of the Lion. Both use aura 137 at +10%. Kings carries the blessing `SpellLabel` 26 and Heart of the Lion carries none. No table ties them as exclusive, and none says they stack.
- The preset keeps Kings out, and its notes now give the real reason.
- If the two stack, Kings is worth about 4% to Fury and Retribution and about 1.5% to Fire.
- A beta character sheet showing both buffs settles it.

**X3, the reworked hybrid trees: no newer client exists.**
- 70707 (wago, 2026-10-30 03:20) is the newest Forever build on wago.tools under any product.
- Blizzard's version service (`us.patch.battle.net:1119/wowf/versions`) also serves 1.60.1.70707. It serves 70450 for `wow_cn_beta` and 70532 for `wow_classic_beta`.
- The new talent spells (1328544 Shieldsworn, and the others) are in no wago `SpellName`.
- Wowhead's Forever tooltip endpoint already serves them.
- So Wowhead is showing a tree from server-side data: a hotfix, or a build ahead of every shipped client. This audit cannot fetch it from a client table.

**The client watcher saw 70685 and 70707; it did not miss them.** `client.yml` failed three nights running:
- **29 Oct, run 40857191821 (70685):** activation stopped at the guide re-encode, by design. Fire/Improved Fireball and Frost/Elemental Precision are gone from the mage trees, so the mage builds no longer fit.
- **30 Oct, run 41040748220 (70707):** stopped for the same reason.
- **31 Oct, run 41190692973:** stopped earlier, at `pipeline build-floors`: "1.60.1.70707 takes 396 items' stats from wowhead, below 90% of 1.60.1.70658's 688".

Nothing is activated past 70658. Activation needs two things:
- The mage guides re-encoded by hand for 70707's mage trees. A fork branch `follow/70707` (ab382f6df) already follows them.
- The build-floors shortfall in 70707's Wowhead item stats settled through `data.yml`. The run's own hint covers `allow_shrink`.

## Sources

**Blizzard and news**
- Blizzard, Mage and Shaman Deep Dive (2026-10-07): https://worldofwarcraft.blizzard.com/news/24302097
- Blizzard, Deep Dive recap: https://news.blizzard.com/en-us/article/24303313
- Icy Veins, Countdown to Forever 8, Druids (2026-10-27): https://www.icy-veins.com/wow-forever/news/countdown-to-forever-8-how-druids-are-shaping-up/
- Endgame Tools, druid changes (2026-10-16, client 70532): https://endgametools.com/en/wow-forever/news/wow-forever-druid

**Warcraft Tavern**
- Spec guides, level-30 beta builds: https://www.warcrafttavern.com/forever/guides/ (retribution-paladin, protection-paladin, holy-paladin, elemental-shaman, enhancement-shaman, restoration-shaman, balance-druid, feral-druid, restoration-druid)
- Paladin deep dive (2026-10-15): https://www.warcrafttavern.com/forever/news/seal-twisting-shockadins-paladin-deep-dive-for-wow-forever/
- Beta patch notes (2026-09-24): https://www.warcrafttavern.com/forever/news/wow-forever-beta-update-bug-fixes-class-changes-gamepad-support/

**Builds and simulations**
- kami-labs builds and sims (published 2026-09-14, sim data 2026-10-26 to 10-28):
  - https://kami-labs.fr/en/wow-forever/build-paladin-vindicte-wow-forever/
  - https://kami-labs.fr/en/wow-forever/build-paladin-protection-wow-forever/
  - https://kami-labs.fr/en/wow-forever/build-paladin-sacre-wow-forever/
  - https://kami-labs.fr/en/wow-forever/build-chaman-amelioration-wow-forever/
  - build-chaman-elementaire-wow-forever, on the same site
- Conquest Capped, shaman talents (2026-10-22): https://conquestcapped.com/guides/wow-forever/wow-forever-shaman-talents/
- SixtySim simulations (2026-10-28): https://sixtysim.com/guides/druid/balance/, https://sixtysim.com/guides/druid/feral/, https://sixtysim.com/guides/druid/restoration/
- wowforeverprogress (2026-10-20): https://wowforeverprogress.com/en/siniflar/druid/balance
- WFDB, feral cat (undated): https://wowfdb.com/classes/druid/feral-cat/

**Other guides and tier lists**
- wowtbc.gg class guides: https://wowtbc.gg/warcraftforever/class-guides/ (retribution-paladin, enhancement-shaman, balance-druid, …)
- Gfinity, Protection Paladin (2026-10-28): https://www.gfinityesports.com/article/world-of-warcraft-forever-protection-paladin-guide-build-rotation
- Mobalytics guides (2026-10-28 to 10-30), read through a search index: https://mobalytics.gg/wow-forever/classes/
- noobtoboss tier list (2026-10-19): https://noobtoboss.com/wowforever-dps-tierlist/
- rankedboost tier list: https://rankedboost.com/world-of-warcraft/forever/tier-list/specs/

**Client data**
- Wowhead live Forever talent payload, fetched 2026-10-31: https://nether.wowhead.com/forever/data/talents-classic?dv=100
- wago.tools DB2 tables `SkillLineAbility`, `SpellEffect`, `Spell` and `SpellName` for build 1.60.1.70658, plus `SpellName` for 1.60.1.70707.
