// web/src/lib/chart/marked-casts.ts
// The life-saving casts a raider looks for beside a death (controls and charts spec §B5.6, fix
// round 1), matched by SPELL ID, not by name. Every id here was checked against the active
// client build's spell table (data/builds/<active>/spells.json); marked-casts.test.ts re-checks
// them against whichever build is active, so a build that drops one fails the suite instead of
// drawing a marker nothing in Forever can cast.
//
// Checked 2026-10-30 against 1.60.1.70658:
//   Rebirth                 20484 20739 20742 20747 20748 (ranks 1-5)
//   Soulstone Resurrection  20707 20762 20763 20764 20765 (ranks 1-5)
//   Power Infusion          10060
//   Mana Tide Totem         16190 17354 17359 (ranks 1-3), 1325256 (Forever's own)
//   Major Healing Potion    17534 (the client names the potion's effect "Healing Potion"),
//                           1244436 (Forever's own "Major Healing Potion")
//   Healthstone             6262 6263 5720 5723 11732 (Minor to Major), 23468-23477 (the
//                           Improved Healthstone ranks), 1326057 (Forever's own)
// All six exist in Forever; none was dropped.

export interface MarkedCast {
  /** The name a marker shows, whatever the client calls the effect. */
  label: string;
  ids: readonly number[];
  /** Icon stem under data/builds/<build>/icons, when the build ships one. */
  icon?: string;
}

export const MARKED_CASTS: readonly MarkedCast[] = [
  { label: 'Rebirth', ids: [20484, 20739, 20742, 20747, 20748], icon: 'spell_nature_reincarnation' },
  { label: 'Soulstone Resurrection', ids: [20707, 20762, 20763, 20764, 20765] },
  { label: 'Power Infusion', ids: [10060], icon: 'spell_holy_powerinfusion' },
  { label: 'Mana Tide Totem', ids: [16190, 17354, 17359, 1325256], icon: 'spell_frost_summonwaterelemental' },
  { label: 'Major Healing Potion', ids: [17534, 1244436] },
  {
    label: 'Healthstone',
    ids: [
      6262, 6263, 5720, 5723, 11732, 23468, 23469, 23470, 23471, 23472, 23473, 23474, 23475, 23476, 23477,
      1326057,
    ],
  },
];

const BY_ID = new Map(MARKED_CASTS.flatMap((cast) => cast.ids.map((id) => [id, cast] as const)));

/** The life-saving cast this spell id is, or undefined. */
export function markedCast(spellId: number): MarkedCast | undefined {
  return BY_ID.get(spellId);
}
