// web/src/test-support/assignment-fixtures.ts
// An Onyxia assignments block as the API answers it (Go guilds.RaidAssignments), for the view,
// actions and planner tests. Names and picks are illustrative test data, never shipped.
import type {
  AssignedName,
  AssignmentRaider,
  AssignmentSlot,
  GuildRaidAssignments,
} from '../lib/guild/assignments-api';

export function raider(name: string, extra: Partial<AssignmentRaider> = {}): AssignmentRaider {
  return {
    character_key: `us/pvp/${name.toLowerCase()}`,
    name,
    class: 'warrior',
    group: 'tank',
    state: 'in',
    own: false,
    ...extra,
  };
}

export function picked(r: AssignmentRaider, offRole = false): AssignedName {
  return { ...r, off_role: offRole };
}

export const TANK = raider('Bulwark');
export const OWN_HEALER = raider('Mendara', { class: 'priest', group: 'healer', own: true });
export const MAGE = raider('Frostquill', { class: 'mage', group: 'ranged', state: 'late' });

export function slot(extra: Partial<AssignmentSlot> = {}): AssignmentSlot {
  return {
    key: 'main_tank',
    label: 'Main tank',
    purpose: 'Picks her up and holds her side-on to the raid in phases 1 and 3.',
    count: 1,
    role: 'tank',
    names: [],
    ...extra,
  };
}

export function onyxia(extra: Partial<GuildRaidAssignments> = {}): GuildRaidAssignments {
  return {
    encounters: [
      {
        encounter_id: 1084,
        name: 'Onyxia',
        raid_zone: "Onyxia's Lair",
        slots: [
          slot({ names: [picked(TANK)] }),
          slot({
            key: 'main_tank_heals',
            label: 'Main tank healers',
            purpose: 'Heal the main tank only.',
            count: 4,
            role: 'healer',
            names: [picked(OWN_HEALER)],
          }),
          slot({
            key: 'whelps_west',
            label: 'Whelps, west cave',
            purpose: 'Pick up the whelps.',
            count: 3,
            role: 'any',
          }),
        ],
        raid_warning: ['MT: Bulwark', 'MT heals: Mendara'],
      },
    ],
    not_in_client: [],
    candidates: [TANK, OWN_HEALER, MAGE],
    posted_at: null,
    ...extra,
  };
}
