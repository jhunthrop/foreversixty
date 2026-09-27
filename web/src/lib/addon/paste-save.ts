// web/src/lib/addon/paste-save.ts
// What the setup page's save-to-account form still needs before it can POST. Kept out of
// the component so the rule is a pure function with its own test: the button is never
// disabled for a missing field (a disabled secondary button looks identical to a live one,
// which read as "the button does nothing"); the click says which field is missing instead.

export type PasteSaveField = 'name' | 'region' | 'ruleset';

export interface PasteSaveValues {
  name: string;
  region: string;
  ruleset: string;
}

/** The first field, in form order, that is still empty; null when the form can be saved. */
export function missingPasteSaveField(values: PasteSaveValues): PasteSaveField | null {
  if (values.name.trim() === '') return 'name';
  if (values.region === '') return 'region';
  if (values.ruleset === '') return 'ruleset';
  return null;
}
