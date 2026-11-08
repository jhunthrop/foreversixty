// web/src/lib/copy/shared.ts
// The words the site and both apps show (apps/shared/copy/shared.json, mobile redesign spec 2.2).
// Each word is its own export in shared/<group>.generated.ts (web/scripts/gen-mobile-copy.mjs), one
// module per key group, so a page bundles only the groups it reads; the copy modules fill their placeholders here and keep their
// own names and exports, so no call site changes.

type Values = Readonly<Record<string, string | number>>;

/** A plural entry: its `one` form, and the `other` form naming the count placeholder. */
export interface SharedPlural {
  readonly count: string;
  readonly one: string;
  readonly other: string;
}

const PLACEHOLDER = /\{([a-zA-Z][a-zA-Z0-9]*)\}/g;

/** `{name}` placeholders filled from `values`; a placeholder without a value fails loudly. */
export function fillShared(template: string, values: Values): string {
  return template.replace(PLACEHOLDER, (_, name: string) => {
    const value = values[name];
    if (value === undefined) throw new Error(`shared copy "${template}": no value for {${name}}`);
    return String(value);
  });
}

/** A plural's `one` form at 1, else `other` with the count filled in. */
export function pluralShared(entry: SharedPlural, count: number): string {
  return fillShared(count === 1 ? entry.one : entry.other, { [entry.count]: count });
}
