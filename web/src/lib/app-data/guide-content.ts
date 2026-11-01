// web/src/lib/app-data/guide-content.ts
// The guide and changelog Markdown files read straight off disk, for the apps' static files
// (Wave 3 contract W3-C1). Astro's `getCollection` returns nothing under a plain vitest run, so
// the endpoints and their tests share this reader, and each file is validated against the very
// schema the site's collection uses (`content.config.ts`).
import { readFileSync, readdirSync } from 'node:fs';
import path from 'node:path';
import { parse as parseYaml } from 'yaml';
import type { z } from 'astro/zod';
import { guideSchema } from '../../content.config';
import { repoRoot } from '../data/build-files';

export type GuideFrontmatter = z.infer<typeof guideSchema>;

export interface GuideSource {
  /** `shaman/elemental`; a class landing is `shaman/index`. */
  id: string;
  data: GuideFrontmatter;
  /** The Markdown under the frontmatter. */
  body: string;
}

const GUIDES_DIR = 'web/src/content/guides';

/** `---\n{yaml}\n---\n{body}` split into the parsed YAML and the body. */
export function splitFrontmatter(raw: string): { frontmatter: unknown; body: string } {
  const [, block, ...rest] = raw.split('---');
  if (block === undefined) throw new Error('guide content: a file has no frontmatter');
  return { frontmatter: parseYaml(block) as unknown, body: rest.join('---') };
}

/** Every spec guide (class landing pages are left out), in directory then file order. */
export function readSpecGuides(root: string = repoRoot()): GuideSource[] {
  const dir = path.join(root, GUIDES_DIR);
  return readdirSync(dir, { withFileTypes: true })
    .filter((entry) => entry.isDirectory())
    .sort((a, b) => a.name.localeCompare(b.name))
    .flatMap((classDir) =>
      readdirSync(path.join(dir, classDir.name))
        .filter((name) => name.endsWith('.md') && name !== 'index.md')
        .sort()
        .map((name): GuideSource => {
          const id = `${classDir.name}/${name.replace(/\.md$/, '')}`;
          const { frontmatter, body } = splitFrontmatter(
            readFileSync(path.join(dir, classDir.name, name), 'utf8'),
          );
          const parsed = guideSchema.safeParse(frontmatter);
          if (!parsed.success)
            throw new Error(`guide content: ${id} fails the guide schema: ${parsed.error.message}`);
          return { id, data: parsed.data, body };
        }),
    );
}
