// web/src/lib/ui/select-options.ts
// The option shapes several dropdowns share (design/specs/2026-10-30-controls-charts.md §A10):
// a player or class with its ringed crest and class colour, and a spec with its icon. One
// place builds them so every list reads the same.
import { classCrestSrc, hasClassCrest } from '../class-crest';
import { specIconSrc } from '../guild/raid-spec';
import { SPECS } from '../sim/specs';
import { classColorToken, classSlug } from '../report/format';
import type { SelectOption } from './listbox-model';

export interface ClassOptionInput {
  value: string;
  label: string;
  /** The class name or slug (`Mage`, `mage`). */
  className: string | undefined;
  group?: string;
  you?: boolean;
  meta?: string;
  description?: string;
}

/** A player or class: the crest and the class colour on the label. A player whose class the log
 *  never named keeps the crest's place as an empty neutral ring, so every player row reads the
 *  same and the names stay in one column. */
export function classOption(input: ClassOptionInput): SelectOption {
  const slug = classSlug(input.className);
  return {
    value: input.value,
    label: input.label,
    group: input.group,
    you: input.you,
    meta: input.meta,
    description: input.description,
    colour: classColorToken(input.className),
    icon: { kind: 'crest', src: slug !== undefined && hasClassCrest(slug) ? classCrestSrc(slug) : '' },
  };
}

/** A spec, or anything else drawn with a square icon. */
export function iconOption(
  value: string,
  label: string,
  src: string | null | undefined,
  kind: 'spec' | 'spell' = 'spec',
): SelectOption {
  return src === null || src === undefined || src === ''
    ? { value, label }
    : { value, label, icon: { kind, src } };
}

/** A class's spec icon by the spec's display name (`Shadow`), else undefined. */
export function specIconByName(className: string | undefined, specName: string): string | undefined {
  const slug = classSlug(className);
  const spec = SPECS.find((entry) => entry.class_slug === slug && entry.name === specName);
  return spec === undefined ? undefined : specIconSrc(spec);
}
