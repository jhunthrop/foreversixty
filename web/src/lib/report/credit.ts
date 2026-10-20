// web/src/lib/report/credit.ts
// Who a report is credited to: the logging character the upload named. Without one the
// report's players are keyed under defaults and nothing is credited, which the subtitle says.
import { nameFromSlug } from '../characters';
import type { ReportMeta } from './types';

export const NOT_CREDITED = 'Not credited to a character';
export const NOT_CREDITED_TITLE =
  'Set the logging character in the companion’s Settings, or log out once with the addon so the companion can read who logged.';

export interface Credit {
  text: string;
  /** The hover; only the missing credit explains itself. */
  title: string | undefined;
  credited: boolean;
}

/** `logging_character` as the API sends it: a "region/ruleset/slug" key, or an object with the name. */
export function reportCredit(meta: Pick<ReportMeta, 'logging_character'> | null): Credit {
  const value = meta?.logging_character;
  let name = '';
  if (typeof value === 'string') name = nameFromSlug(value.split('/').pop() ?? '');
  else if (value && typeof value === 'object') name = value.name ?? '';
  if (name === '') return { text: NOT_CREDITED, title: NOT_CREDITED_TITLE, credited: false };
  return { text: `Credited to ${name}`, title: undefined, credited: true };
}
