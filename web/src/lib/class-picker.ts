// web/src/lib/class-picker.ts
// The one class picker model (component audit 2026-11-10 cluster 26): the nine classes in the
// game's order, each with its specs in talent-tree order, its guide links and its spec icons. The
// home and 404 grid (ui/ClassPicker), the /guides spec panels (guides/ClassSpecPicker) and the
// header's Guides panel all read it.
import { classDisplayName } from './sim/spec-label';
import { SPECS, type Spec } from './sim/specs';
import { CLASS_SLUGS, type ClassSlug } from './ui/primitives';

export interface PickerSpec {
  /** `<class>/<spec>`: what a page's answer matches. */
  key: string;
  name: string;
  href: string;
  icon: string;
  /** The talent tree: a role variant (Feral Bear) shares its spec's. */
  tree: number;
}

export interface PickerClass {
  slug: ClassSlug;
  name: string;
  /** The class's own page; none where a class has no page (/bis). */
  href?: string;
  specs: readonly PickerSpec[];
}

export interface PickerLinks {
  classHref?: (classSlug: ClassSlug) => string;
  specHref: (classSlug: ClassSlug, specSlug: string) => string;
  specIcon?: (spec: Spec) => string;
}

/** The guides: /guides/<class> and /guides/<class>/<spec>. */
export const GUIDE_LINKS: PickerLinks = {
  classHref: (classSlug) => `/guides/${classSlug}`,
  specHref: (classSlug, specSlug) => `/guides/${classSlug}/${specSlug}`,
};

export function pickerClasses(
  links: PickerLinks = GUIDE_LINKS,
  specs: readonly Spec[] = SPECS,
): PickerClass[] {
  return CLASS_SLUGS.map((slug) => ({
    slug,
    name: classDisplayName(slug),
    href: links.classHref?.(slug),
    specs: specs
      .filter((spec) => spec.class_slug === slug)
      .toSorted((a, b) => a.tree_index - b.tree_index)
      .map((spec) => ({
        key: `${slug}/${spec.spec_slug}`,
        name: spec.name,
        href: links.specHref(slug, spec.spec_slug),
        icon: links.specIcon?.(spec) ?? '',
        tree: spec.tree_index,
      })),
  }));
}

/** A class's specs, one per talent tree (Feral Bear folds into Feral): what every picker lists. */
export function treeSpecs(pickerClass: PickerClass): PickerSpec[] {
  return pickerClass.specs.filter(
    (spec, index, all) => all.findIndex((other) => other.tree === spec.tree) === index,
  );
}

/** A class's spec names, one per talent tree ("Balance, Feral, Restoration"). */
export function specSummary(pickerClass: PickerClass, separator: string): string {
  return treeSpecs(pickerClass)
    .map((spec) => spec.name)
    .join(separator);
}
