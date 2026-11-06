// web/src/lib/planner/fs1-drops.ts
// The FS1 export's drops= section (guild slice 7), written by the addon's RaidDrops.lua through
// Codec.encodeDropCorpse and read by the API's api/internal/fs1/drops.go: each boss corpse the
// character's loot window read in a raid, `;`-joined:
//
//   <unix>:<instance_id>:<encounter_id>:<method>:<source>:<item_id>x<count>[@<winner>+<winner>],...
//
// The planner has no use for it beyond knowing the section, so a pasted export never reports it
// as unknown. Read leniently: a malformed corpse or item is dropped, never the code.

export interface FS1DropItem {
  itemId: number;
  count: number;
  winners: string[];
}

export interface FS1DropCorpse {
  at: number;
  instanceId: number;
  encounterId: number;
  method: string;
  source: string;
  items: FS1DropItem[];
}

const CORPSE = /^(\d+):(\d+):(\d+):([a-z]*):([A-Za-z0-9-]*):(.*)$/;
const ITEM = /^(\d+)x(\d+)(?:@(.*))?$/;

function decodeWinner(raw: string): string {
  try {
    return decodeURIComponent(raw);
  } catch {
    return raw;
  }
}

function dropItem(text: string): FS1DropItem[] {
  const match = ITEM.exec(text);
  if (match === null) return [];
  const [, id, count, winners] = match;
  const names = (winners ?? '')
    .split('+')
    .filter((name) => name !== '')
    .map(decodeWinner);
  return [{ itemId: Number(id), count: Number(count), winners: names }];
}

function dropCorpse(entry: string): FS1DropCorpse[] {
  const match = CORPSE.exec(entry);
  if (match === null) return [];
  const [, at, instanceId, encounterId, method, source, list] = match;
  const items = (list ?? '')
    .split(',')
    .filter((text) => text !== '')
    .flatMap(dropItem);
  return [
    {
      at: Number(at),
      instanceId: Number(instanceId),
      encounterId: Number(encounterId),
      method: method ?? '',
      source: source ?? '',
      items,
    },
  ];
}

/** The drops= section's corpses, a malformed one dropped. */
export function decodeDrops(field: string): FS1DropCorpse[] {
  return field === '' ? [] : field.split(';').flatMap(dropCorpse);
}
