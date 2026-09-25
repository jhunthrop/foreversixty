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

const dungeons = defineCollection({
  loader: glob({ pattern: '**/*.md', base: './src/content/dungeons' }),
  schema: factSchema.extend({
    zone: z.string(),
    levelMin: z.number().int().min(1).max(60).optional(),
    levelMax: z.number().int().min(1).max(60).optional(),
    order: z.number().int(),
  }),
});

const zones = defineCollection({
  loader: glob({ pattern: '**/*.md', base: './src/content/zones' }),
  schema: factSchema.extend({
    continent: z.enum(['Eastern Kingdoms', 'Kalimdor', 'Zephras Isle']),
    levelMin: z.number().int().optional(),
    levelMax: z.number().int().optional(),
    isNew: z.boolean().default(true),
  }),
});

const guides = defineCollection({
  loader: glob({ pattern: '**/*.md', base: './src/content/guides' }),
  schema: factSchema.extend({
    classSlug: z.string(),
    // Present on every spec guide (e.g. `warrior/fury`), absent on a class landing page
    // (`warrior/index`), which covers the whole class rather than one tree.
    spec: z.string().optional(),
    role: z.enum(['dps', 'healer', 'tank']).optional(),
  }),
});

export const collections = { pages, changelog, dungeons, zones, guides };
