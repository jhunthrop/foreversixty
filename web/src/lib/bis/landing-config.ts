// web/src/lib/bis/landing-config.ts
// What the /bis landing's browser script needs to know about classes and specs, serialised into
// the page by index.astro: the lead reads a character's class and spec NAMES from /v1/me and has
// to turn them into the spec's page, crest colour and label without importing the build-time
// catalogue.
export interface LandingSpec {
  classSlug: string;
  specSlug: string;
  name: string;
}

export interface LandingClass {
  slug: string;
  name: string;
  color: string;
  /** The class's first spec, the page /bis lists first for it. */
  defaultSpecKey: string;
}

export interface LandingConfig {
  build: string;
  /** Keyed by class slug. */
  classes: Record<string, LandingClass>;
  /** Keyed by spec key, "warrior-fury". */
  specs: Record<string, LandingSpec>;
}

/** The class a /v1/me class name ("Warrior") names, or undefined when the site does not know it. */
export function classFor(config: LandingConfig, className: string | undefined): LandingClass | undefined {
  return className === undefined ? undefined : config.classes[className.toLowerCase()];
}

/** The spec key for a class and the spec name the character reports ("Fury"). */
export function specKeyFor(
  config: LandingConfig,
  className: string | undefined,
  specName: string | undefined,
): string | undefined {
  const classSlug = classFor(config, className)?.slug;
  if (classSlug === undefined || specName === undefined) return undefined;
  return Object.entries(config.specs).find(
    ([, spec]) => spec.classSlug === classSlug && spec.name === specName,
  )?.[0];
}
