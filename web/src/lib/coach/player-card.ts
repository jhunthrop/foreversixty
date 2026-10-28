// web/src/lib/coach/player-card.ts
// The rail's player card line (spec 3.1): `Fury Warrior · Onyxia (40) · Kill 3:52`, built from the
// review's spec, the roster's class and the fight. Parts the report does not know are left out.
import { clock } from './chips';
import type { FightEntry } from '../report/types';

export function playerCardLine(
  spec: string,
  className: string | undefined,
  fight: Pick<FightEntry, 'name' | 'size' | 'kill' | 'duration_ms'> | null,
): string {
  const who = [spec, className].filter((part): part is string => part !== undefined && part !== '');
  const parts = [who.join(' ')];
  if (fight !== null) {
    parts.push(fight.size === undefined ? fight.name : `${fight.name} (${fight.size})`);
    parts.push(`${fight.kill ? 'Kill' : 'Wipe'} ${clock(fight.duration_ms)}`);
  }
  return parts.filter((part) => part !== '').join(' · ');
}
