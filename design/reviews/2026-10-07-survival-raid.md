# Survival raid DPS and the negative haste weight (2026-10-07)

Spec hunter-survival, band 60, alliance dwarf, published raid gear (Teebu's
Blazing Longsword, Ravencrest's Legacy, Riphook), build 1.60.1.70009.
Reproduced with a request built the way the ranker builds the raid band
(`bandCharacter` + `plainRequest` + `request.ResolvedPreset.Layer`), 1000 to
3000 iterations, seed 7. Published 847.0 reproduced as 846.0 +- 1.4.

## Result

| hunter-survival, band 60 alliance | before | after |
|---|---|---|
| bare set_dps | 265.2 | 237.2 |
| raid set_dps | 847.0 | 745.8 |
| raid / bare | 3.2x | 3.1x |

hunter-marksmanship is unchanged (bare 243.8, raid 637.8): neither fix touches
it. The raid haste weight for survival is unchanged (-9.0 before, -9.1 after);
it is not an engine defect, see below.

## Damage breakdown, raid, before (DPS share)

Mongoose Bite 274.4 (43.9 casts a fight), Raptor Strike 110.0, off-hand
autos 117.6, main-hand autos 82.4, Windfury extra main-hand attacks 51.7
(20.5 procs), Strider Kick 64.7, Lacerating Strikes 20.5, Immolation Trap 26.9,
Teebu's proc 5.9, pet (Cat) 91.4. Lacerate, Wing Clip and ranged shots are not
cast: the rotation never names them and the hunter stands in melee, inside the
12 yard ranged minimum. Mana is not the limiter (Judgement of Wisdom returns
about 5.3M against 6.1M spent). Boss never attacks: no parries, Counterattack
idle as the rotation file already records.

After: Mongoose Bite 183.1, Lacerating Strikes 13.7, Windfury extras 50.1,
everything else within noise.

## Checks against the client and the engine

- Windfury proc rule. The totem imbue procs only from landed main-hand hits
  (`ProcMaskMeleeMH`), 20%, 1.5s internal cooldown, one extra attack that is
  a main-hand swing (it resets the main-hand timer, as a real swing does). It
  cannot proc off the off hand. Matches; the extra attack is one cast, so
  Teebu's and other weapon procs roll once on it, never twice. The extra
  attack also carries the Windfury buff stacks as designed.
- Windfury values: CONTRADICTED. The client (spells 8516/10608/10610) gives
  attack power 95/179/246 for 1 second; the engine had vanilla's 122/229/315
  for 1.5 seconds. Fixed in `sim/core/buffs.go` (`WindfuryBuffBonusAP`,
  `WindfuryBuffDuration`). The 1.5s internal cooldown has no client value and
  stays (`extraAttackProcICD`). The dps_warrior golden moves (about -0.3%)
  only through this.
- Mongoose Bite: CONTRADICTED (engine bug). Its proc mask was
  `ProcMaskMeleeSpecial` (both hands) though the tooltip requires a main-hand
  weapon and the damage is main-hand only. Predator's Edge multiplies every
  off-hand-masked spell, so Mongoose Bite took +50% damage at 5/5 and fired
  off-hand procs. Now `ProcMaskMeleeMHSpecial`. This is the bulk of the
  847 -> 746 drop (Mongoose Bite 274 -> 183, its Lacerating Strikes bleed
  -33%).
- Mongoose Bite trigger. The tooltip (Wowhead Forever spell 1495) lists only
  "requires main hand weapon" and "cannot be used while shapeshifted", no dodge.
  The rotation casts it on cooldown, which is legal under that reading, so I
  left it. Open question for the owner: Expose Prey's client text says "chance
  to activate your Mongoose Bite for 5 sec", which reads as if an activation
  state still exists. If Mongoose Bite needs a dodge, a boss that never attacks
  can only grant it through Expose Prey (needs Hunter's Mark, 10% on landed
  hits), and survival would fall far below these numbers. The client spell
  tables carry no cast-requirement field, so they cannot settle it.
  Expose Prey resets the cooldown on 10% of landed hits when Hunter's Mark is
  up, which is why Mongoose Bite casts 43.9 a fight in the raid run (36 is the
  5s cap) and why dropping Hunter's Mark costs 59 DPS.
- Lacerate: 7 ticks of 3s, per-tick client amount, no attack power scaling
  (matches the client: effect 6/aura 3, ap coefficient 0). A recast refreshes
  the single dot (no stacking). Not in the rotation.
- Lacerating Strikes: 40% of the Mongoose Bite hit over 7 ticks, refresh not
  stack. Matches.
- Raptor Strike: weapon damage plus the client flat (70 at rank 8), 6s
  cooldown, queued on the main-hand swing and removing the dual-wield miss
  penalty while queued (the warrior Heroic Strike pattern). Matches.
- Dual wield: off-hand damage 50%, +19% miss on whites while dual wielding,
  hunter has Predator's Edge as the off-hand bonus. Matches (the Mongoose
  mask above was the leak).
- Trueshot Aura: one exclusive aura per unit, melee and ranged attack power
  granted once. Matches. Hunter's Mark is ranged attack power only.

## The negative haste weight

Not noise: -9.0 +- 2.6 at 100 iterations, -6.9 +- 0.8 at 1000. Not Windfury
(same weight with Windfury removed), not Hunter's Mark, not the Mongoose
elixir. It comes from the weights ladder character, which a hunter wears with
only a ranged weapon: it has no melee weapon, so it swings fists at 1.0s.
Raptor Strike has a 6s cooldown, exactly six fist swings. At exactly zero
bonus haste the cooldown expires on the sixth swing, so Raptor Strike fires
every 6.0s (30.5 a fight). At +0.5% haste six swings end just before the
cooldown and the strike waits a seventh swing (26.3 a fight, DPS 289.7 ->
286.9). Past that DPS climbs again with haste (+10 haste: 296.6, +20: 303.2).
The sweep measures at +-1 haste, straddling that step, and the +1 side is paired
with a deterministic loss, so the error bar understates the real uncertainty.
Give the same character Teebu's and Ravencrest's and the weight is +10.9. The
bare ladder shows the same quantisation, smaller and insignificant (+3.1).

Fix belongs to the ranker, not the engine or the hunter files: a melee
hunter's ladder character needs a melee main hand and off hand in
`sim/cmd/leveling-bis/character.go` (`ladderWeapon` picks a ranged weapon
only). Not changed here; that code is owned by the ranker lane.

## Tests

Fork: `TestMongooseBiteIsAMainHandSpecial`, `TestWindfuryTotemBuffMatchesClient`
(written first, failing). Fork goldens: dps_warrior moved by the Windfury
numbers only. Site: the survival ladder golden moves in its DPS column only.
