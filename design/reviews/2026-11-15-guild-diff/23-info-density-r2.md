# Information density re-gate, round 2 (2026-11-15)

Checked only the round-1 must-fixes against the four specs and their Review fixes sections. Privacy held: names officer-only,
raiders see their own rows. Nothing here asks for a name to reach a raider or for data the repo lacks.

| Spec | Verdict |
|---|---|
| loot-plan | SHIP (optional: Copy CSV) |
| live-pulls | SHIP (optional: raider active time) |
| prep-reminder | SHIP |
| guild-later | SHIP (briefs; recruiting still blocked on G1 and R1) |

## loot-plan
Resolved: drop chance on drop and raider rows, `replaces` and item level change on the raider row, per-role counts on the drop list,
vote tally and LootCandidate link, per-source split of the gear count.
Open, minor: group-head Copy CSV is not in the spec (no `csv` anywhere). Optional, not a block.

## live-pulls
Resolved: death phase, last hits, overkill and auras, interrupts, dispels and avoidable hits in v1, boss % trend strip, Deaths by phase,
repeat deaths rail line, last upload stamp, Copy CSV and sort, raider own view (deaths with cause and phase, damage taken, trend strip).
Open, minor: raider own active time against fight length is still unspecified; downtime is held behind D7, so this is consistent.

## prep-reminder
Resolved: unspent talents, weapon enchant, protection potions, gear gap and encounter items on the member block; raid tiles; sort,
class filter and copy for officers; Attuned and a where per fix in the DM; item level delta in the addon panel; resistance shown only
where a raid needs it (the per-boss number stays under needs data).

## guild-later
Resolved: bench rows gain Adds, before and after cover counts and a coverage figure; G2 stated and K1 checked in game before K3;
recruiting row gains item level, rare piece count and `Fills:`; sort, result count and a saved filter; `Copy whisper`.

## Needs data (unchanged)
Resistance need per boss, marginal DPS per bench player (K2), group assignments, consumables at pull (COMBATANT_INFO log).
