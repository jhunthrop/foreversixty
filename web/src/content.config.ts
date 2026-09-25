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
});

const guides = defineCollection({
  loader: glob({ pattern: '**/*.md', base: './src/content/guides' }),
  schema: guideSchema,
});

export const collections = { pages, changelog, guides };
