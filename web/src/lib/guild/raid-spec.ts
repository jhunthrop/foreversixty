// web/src/lib/guild/raid-spec.ts
// The spec catalogue as the raid page reads it: a class's specs and a spec's own icon. The picker's
// (spec, role) pairs live in raid-comp.ts.
import activeBuild from '../../data/active-build.json';
import { SPECS, type Spec } from '../sim/specs';
import { catalogueIconUrl } from '../ui/assets';

/** The class's specs in catalogue order. */
export function specsForClass(classSlug: string): readonly Spec[] {
  return SPECS.filter((spec) => spec.class_slug === classSlug);
}

/** The spec's talent-tab icon, from the active build. */
export function specIconSrc(spec: Spec): string {
  return catalogueIconUrl(activeBuild.build, spec.icon);
}
