// web/src/fixtures/me-addon.ts
// A GET /v1/me fixture exercising the home rebuild spec's §3.B signed-in hero now that
// `MeCharacter.build.gear`/`.talents` exist (the API lane's own extension to the contract,
// landed alongside this lane) -- for e2e specs to route **/v1/me to via page.route, and for
// screenshots of the signed-in home against the mock's own worked example (`SignedIn.png`).
//
// Three characters, matching the mock's own Switch character panel exactly:
//
//   - Zulmara: Horde Troll Marksmanship Hunter, level 24, with a worn-gear export that
//     produces exactly 4 upgrades against the real, committed
//     `data/builds/1.60.1.70009/bis/hunter-marksmanship.json` band 20 (horde) -- four slots
//     (head, neck, feet, ranged) wear a different, real, lower item than the band's own
//     pick (the same four items the home rebuild spec's own build brief names: Resilient
//     Cloth Headband, Lil Timmy's Peashooter, Footpads of the Fang, Erudite's Amulet), every
//     other known slot wears the band's own pick exactly (10 "already best in slot" ties).
//     Her talents differ from the band's own `"0000000000000000-35300000000000000-
//     000000000000000000"` build by exactly 2 points (one tree-1 cell at rank 1 instead of
//     3), so the Talents card reads "Unoptimized."
//   - Frostspine: Undead Frost Mage, level 17 (bandForLevel clamps below 20 up to the 20
//     band, `lib/bis/hover.ts`'s own `bandForLevel`), with a gear export that produces
//     exactly 2 upgrades (head, ranged) against `mage-frost.json` band 20 (horde), talents
//     matching the band exactly (an "Optimized" character, for contrast with Zulmara).
//   - Grokmar: Orc Warrior, level 9, no spec learned yet -- the honest "Pick a spec" state,
//     never a fabricated upgrade count for a character the addon has not captured a build
//     for.
import type { Me } from '../lib/account/api';

const now = Date.now();

export const meAddonFixture: Me = {
  user: { id: 42, battletag: 'Fixture#4242', email: null, role: 'user', anonymize: false },
  characters: [
    {
      key: 'us/normal/zulmara',
      region: 'us',
      ruleset: 'normal',
      name: 'Zulmara',
      class: 'Hunter',
      spec: 'Marksmanship',
      race: 'Troll',
      realm: 'Skyborne',
      level: 24,
      faction: 'horde',
      guild: { id: 9, name: 'Sample Guild', verified: true },
      source: 'export',
      build: {
        source: 'addon',
        captured_at: new Date(now - 4 * 60_000).toISOString(),
        level: 24,
        data_build: '1.60.1.70009',
        gear: {
          // Upgrades (4): worn differs from the band 20 (horde) pick.
          head: 211500, // Resilient Cloth Headband -- pick is Brawler's Leather Hood (252504)
          neck: 277204, // Erudite's Amulet -- pick is Scout's Medallion (20442)
          feet: 10411, // Footpads of the Fang -- pick is Feet of the Lynx (1121)
          ranged: 13136, // Lil Timmy's Peashooter -- pick is Ranger Bow (3021)
          // Already best in slot (10): worn is the band's own pick.
          shoulder: 5404,
          back: 6449,
          chest: 252490,
          wrist: 3202,
          hands: 5970,
          waist: 10403,
          legs: 10410,
          finger1: 285330,
          finger2: 20429,
          main_hand: 250603,
        },
        talents: {
          // Band 20 (horde) is "0000000000000000-35300000000000000-000000000000000000" --
          // this differs by exactly 2 points: tree 1's first cell is rank 1 here, rank 3 on
          // the band's own build (|1 - 3| = 2).
          trees: ['0000000000000000', '15300000000000000', '000000000000000000'],
          points: [0, 9, 0],
        },
      },
    },
    {
      key: 'us/normal/frostspine',
      region: 'us',
      ruleset: 'normal',
      name: 'Frostspine',
      class: 'Mage',
      spec: 'Frost',
      race: 'Undead',
      realm: 'Skyborne',
      level: 17,
      faction: 'horde',
      source: 'export',
      build: {
        source: 'addon',
        captured_at: new Date(now - 22 * 60_000).toISOString(),
        level: 17,
        data_build: '1.60.1.70009',
        gear: {
          // Upgrades (2): worn differs from the band 20 (horde) pick.
          head: 211500, // Resilient Cloth Headband -- pick is Pristine Circlet (253949)
          ranged: 5604, // Elven Wand -- pick is Cookie's Stirring Rod (5198)
          // Already best in slot: worn is the band's own pick.
          shoulder: 12998,
          back: 4311,
          chest: 253901,
          wrist: 251486,
          hands: 5970,
          waist: 253925,
          legs: 23173,
          feet: 4320,
          finger1: 20426,
          finger2: 1156,
          main_hand: 251534,
        },
        talents: {
          // Matches band 20 (horde)'s own build exactly -- "Optimized," for contrast with
          // Zulmara's "Unoptimized" state.
          trees: ['000000000000000000', '00000000000000000', '2531000000000000000'],
          points: [0, 0, 11],
        },
      },
    },
    {
      key: 'us/normal/grokmar',
      region: 'us',
      ruleset: 'normal',
      name: 'Grokmar',
      class: 'Warrior',
      race: 'Orc',
      realm: 'Skyborne',
      level: 9,
      faction: 'horde',
      source: 'export',
      // No spec learned yet (an early character the addon has synced basic facts for, but
      // whose talent split has not resolved to a recognised spec) -- the honest "Pick a
      // spec" state, never a fabricated upgrade count.
      build: { source: 'addon', captured_at: new Date(now - 90 * 60_000).toISOString(), level: 9 },
    },
  ],
  guilds: [],
  main_character_key: 'us/normal/zulmara',
};
