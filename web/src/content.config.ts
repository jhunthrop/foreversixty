import { defineCollection } from 'astro:content';
import { z } from 'astro/zod';
import { glob } from 'astro/loaders';

export const sourceKinds = ['blizzard', 'datamined', 'community', 'site'] as const;
export const confidences = ['confirmed', 'single-source', 'inferred'] as const;

export const sourceSchema = z.object({
  label: z.string().min(1),
  url: z.url(),
  kind: z.enum(sourceKinds),
});

export const factSchema = z.object({
  title: z.string().min(1),
  description: z.string().min(1).max(200).optional(),
  updated: z.coerce.date(),
  confidence: z.enum(confidences).default('confirmed'),
  sources: z.array(sourceSchema).min(1),
});

const pages = defineCollection({
  loader: glob({ pattern: '**/*.md', base: './src/content/pages' }),
  schema: factSchema,
});

const changelog = defineCollection({
  loader: glob({ pattern: '**/*.md', base: './src/content/changelog' }),
  schema: factSchema.extend({
    date: z.coerce.date(),
    kind: z.enum(sourceKinds),
    note: z.string().optional(),
  }),
});

// Exported (not just used inline below) so a plain Node test can validate a guide's raw
// frontmatter against the exact same schema without going through Astro's content layer --
// `getCollection` needs the dev/build pipeline and returns nothing under plain `vitest run`,
// for every collection in this file, not only this one (see src/content/guides/_sections.test.ts).
export const guideSchema = factSchema.extend({
  classSlug: z.string(),
  // Present on every spec guide (e.g. `warrior/fury`), absent on a class landing page
  // (`warrior/index`), which covers the whole class rather than one tree.
  spec: z.string().optional(),
  role: z.enum(['dps', 'healer', 'tank']).optional(),
  /** An FS1 code for this spec's recommended build (spec guides only). `[spec].astro`
   *  decodes it to light the embedded read-only tree and to build the "Load this build" /
   *  "Sim this build" links; `_sections.test.ts`'s own SSR test decodes every one and checks
   *  it names this guide's own class and stays legally reachable within 51 points. */
  build: z.string().optional(),
  /** An FS1 code for the raid build (spec guides only): the build the raid-ready preset's
   *  level-60 entry is simmed on, which the BiS ranker reads for that entry alone. Absent when
   *  the raid search found nothing better than `build`, so the leveling build is also the
   *  raid build. Same legality rules as `build`; `builds.test.ts` checks it the same way. */
  raidBuild: z.string().optional(),
  /** Race slugs (races.json's own `slug`) this guide calls a strong pick, in the order the
   *  guide's own Races prose names them. RacePillRow.astro marks these among the class's
   *  full legal race list; an empty array (the default, and every class landing page's
   *  value) marks none. */
  recommendedRaces: z.array(z.string()).default([]),
  /** The set of stats this guide's own Stat priority section covers. StatPriorityPills.astro
   *  still renders these as a numbered pill row in this array's own order, but that order is
   *  no longer authoritative (live-numbers lane, 2026-10-06): `GuideStatTable.astro` renders
   *  the real per-point order from this band's own sim weights (`railStatRows`), which can
   *  and does disagree with this array -- the sim re-ranks nightly, this array does not. Only
   *  the SET of stats named here still matters; do not read this array's order as a ranking. */
  statPriority: z.array(z.string()).default([]),
});

const guides = defineCollection({
  loader: glob({ pattern: '**/*.md', base: './src/content/guides' }),
  schema: guideSchema,
});

export const collections = { pages, changelog, guides };
