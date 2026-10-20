// web/src/lib/report/credit.ts
// Who a report is credited to: the logging character the upload named. Without one the
// report's players are keyed under defaults and nothing is credited; the subtitle says so
// in plain text (a phone has no hover).
import { nameFromSlug } from '../characters';
import type { ReportMeta } from './types';

export const NOT_CREDITED = 'Not credited yet. Log out once in game; the companion credits it then.';

export interface Credit {
  /** Plain text before the name ("Credited to "), or the whole not-credited sentence. */
  prefix: string;
  /** The character's name; empty when not credited. */
  name: string;
  /** The character page, when the key is known. */
  href: string | undefined;
  credited: boolean;
}

/** `logging_character` as the API sends it: `{ key, name }` or null; an older "region/ruleset/slug" string still reads. */
export function reportCredit(meta: Pick<ReportMeta, 'logging_character'> | null): Credit {
  const value = meta?.logging_character;
  let key = '';
  let name = '';
  if (typeof value === 'string') key = value;
  else if (value && typeof value === 'object') {
    key = value.key ?? '';
    name = value.name ?? '';
  }
  if (name === '' && key !== '') name = nameFromSlug(key.split('/').pop() ?? '');
  if (name === '') return { prefix: NOT_CREDITED, name: '', href: undefined, credited: false };
  const href = key.split('/').length === 3 ? `/character/${key}` : undefined;
  return { prefix: 'Credited to ', name, href, credited: true };
}
