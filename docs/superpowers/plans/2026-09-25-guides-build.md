# Guides that end in a build — Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Every spec guide's "Talents and builds" section embeds the planner's own talent
tree (read-only) lit to a real, decodable FS1 build; "Stat priority" and "Races" render as
pill rows; "Load this build" / "Sim this build" buttons hand the code to the planner/sim;
the "may change" caveat moves to a quiet line above Sources.

**Architecture:** No second tree renderer and no MDX (not installed, and `npm install` is
forbidden in this worktree). `TreeGrid.svelte`/`TalentCell.svelte` grow a `readOnly` mode.
A new Svelte island (`GuideBuildTree.svelte`) decodes a guide's `build:` FS1 code, fetches
that build's talent file, reconstructs a legal point order with the planner's own
`orderFromRanks`, and renders three read-only `TreeGrid`s. Because a spec guide's markdown
body needs an island and two new Astro components planted at specific points *inside* its
nine-section flow, and this project's Markdown files cannot embed Astro/Svelte components
(no MDX), `[spec].astro` splits the entry's raw body into its nine `## `-delimited sections
(`splitSpecSections`, extracted from the existing test's own regex) and renders each
section's markdown independently through `@astrojs/markdown-satteri`'s own
`createSatteriMarkdownProcessor` (the same processor Astro's content pipeline uses
internally; already installed as one of Astro's own dependencies, confirmed importable
directly, so this needs no new package). This lets the page interleave real HTML fragments
with real Astro/Svelte components while every section keeps the exact heading markup
(`<h2 id="...">`, matching `slugify()`) the table of contents already depends on.

**Tech Stack:** Astro 7, Svelte 5 (runes), Tailwind v4 (`@theme static` tokens), Vitest,
Playwright, `@astrojs/markdown-satteri` (already a transitive Astro dependency).

**Spec:** `docs/superpowers/specs/2026-09-25-product-not-wiki-design.md`, section 5 (binding
sections 1 and 2 apply to every lane; section 8 names ownership; section 9 names the tests).

## Global Constraints

- Design system, states model, `createQueryState`-only session reads, `characterDescriptor`
  vocabulary, honesty (no invented numbers), Lighthouse budgets held and quoted against
  `main`, old URLs redirect, lane rules — spec section 1 items 1-9, verbatim.
- Lane ownership (spec section 8): this lane owns `content/guides/**` (frontmatter `build:`),
  `pages/guides/**`, new `components/guides/*`, and a read-only mode on
  `components/planner/TreeGrid.svelte`. Lane B owns `CurrentCharacterBar.svelte`,
  `CurrentCharacterChip.svelte`, `lib/current-character*.ts`, the pages' bar mounts,
  `Rankings.svelte`'s prefilter, the logs "Your reports" order, sim bootstrap — none of
  these are touched here. Lane D owns copy modules, empty states, `ReportView.svelte`,
  `SummaryBar.svelte`, `SimResults` — none of these are touched here. Lane A has already
  merged (confirmed: `web/src/pages/classes` is gone, `SETUP_NAV_ITEM` exists).
- Additional files this plan must touch that no lane exclusively owns (ledgered as rulings
  in Task list below): `web/src/components/planner/TalentCell.svelte` (the read-only mode
  cannot be built in `TreeGrid.svelte` alone — clicking and hovering are `TalentCell`'s own
  event handlers), `web/src/layouts/Content.astro` (the "may change" line's new position is
  shared markup also used by the `pages` collection and must not move for those), and
  `web/lighthouserc.json` (a new URL must be added for the guide page to be measured at all).
- Real build data for the FS1 codes comes from `data/builds/1.60.1.69893/talents/<class>.json`
  at the repo root (checked into git, confirmed via `git ls-files`), not from
  `web/public/data/` (gitignored, sync-generated). `data/builds/1.60.1.69893/races.json` is
  the source for race slugs/factions.
- Web toolchain (from `web/`): `export NVM_DIR="$HOME/.nvm"; . "$NVM_DIR/nvm.sh"; nvm use
  22.12` then `FOREVER_DATA=fixture npm run sync` once. The SSR test in Task 3 reads the
  real (non-fixture) `data/builds/1.60.1.69893/talents/*.json` directly off disk with
  `node:fs`, exactly as `_sections.test.ts` already reads guide markdown off disk — it does
  **not** need `public/data/` populated, so it works under the fixture sync too. E2E port is
  4387.
- Never `npm install`; `web/node_modules` is a symlink — check `git status` before every
  commit. Commit messages: write to a file, `git commit -F <file>` as its own Bash command.
  Trailer: `Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>` then
  `Claude-Session: https://claude.ai/code/session_01DAHb7A9Ytd6qpZu2UujXgY`.
- Coordinator's hard rules (verbatim): every spec guide's `build:` FS1 code decodes, names
  the guide's own class, and its point totals match the guide's own "Talents and builds"
  prose (an SSR test decodes every guide and checks class + total points ≤ 51); the embedded
  tree is `TreeGrid.svelte`'s own read-only mode, never a second renderer; "Load this
  build"/"Sim this build" behave as spec 5 describes; stat priority is an ordered pill row,
  races are a pill row with recommended ones marked (no invented race art — a name pill in
  the faction colour); the "may change" caveat sits in a quiet line above Sources; the guide
  page is in Lighthouse's strict bucket and the tree embed hydrates `client:visible` and is
  never the LCP element.

---

## Pre-computed data (read once, used across several tasks)

### FS1 `build:` codes

Computed with a throwaway script (not part of this plan's tree) that used the real
`indexTalents`/`orderFromRanks`/`encodeFS1`/`decodeFS1` from `web/src/lib/planner/{rules,fs1}.ts`
against `data/builds/1.60.1.69893/talents/<class>.json`: every code below decodes, names the
right class, and round-trips through `decodeFS1` → `orderFromRanks` with **zero dropped
talents** (every point is legally reachable — tier gates and prerequisites all satisfied) and
a total ≤ 51. Every talent spent is either explicitly named in that guide's own "Talents and
builds" prose, or is a real, data-backed prerequisite of one that is (e.g. Fury's "Deep
Wounds" needs 3 points in Arms's "Improved Rend" first — a real talent, never invented).
Where a guide's own prose under-specifies a secondary tree (naming no specific talent, e.g.
"a few utility points into Survival"), the code spends the leftover budget bottom-up in that
same named tree rather than inventing a pick to name.

| Guide | `build:` (quote it — the string contains colons) | Points (primary/secondary/…) |
|---|---|---|
| druid/balance | `FS1:1.60.1.69893:druid:night-elf:5532220005501001/0/5:` | Balance 31, Restoration 5 |
| druid/feral | `FS1:1.60.1.69893:druid:night-elf:0/5523232020032010001/055:` | Feral Combat 31, Restoration 10 |
| druid/restoration | `FS1:1.60.1.69893:druid:night-elf:503/0/5552005003113001:` | Balance 8, Restoration 31 |
| hunter/beast-mastery | `FS1:1.60.1.69893:hunter:dwarf:5420001505001251/3551/51:` | Beast Mastery 31, Marksmanship 14, Survival 6 |
| hunter/marksmanship | `FS1:1.60.1.69893:hunter:dwarf:5522/35305500115003/51:` | Beast Mastery 14, Marksmanship 31, Survival 6 |
| hunter/survival | `FS1:1.60.1.69893:hunter:dwarf:0/35534/5552230010512:` | Marksmanship 20, Survival 31 |
| mage/arcane | `FS1:1.60.1.69893:mage:gnome:255225200000011501/2305/0:` | Arcane 31, Fire 10 |
| mage/fire | `FS1:1.60.1.69893:mage:gnome:230005/23552200030003051/0:` | Arcane 10, Fire 31 |
| mage/frost | `FS1:1.60.1.69893:mage:gnome:0/2305/2555100300000301051:` | Fire 10, Frost 31 |
| paladin/holy | `FS1:1.60.1.69893:paladin:human:255321003025101001/5/0:` | Holy 31, Protection 5 |
| paladin/protection | `FS1:1.60.1.69893:paladin:dwarf:253000003/5532410301001051/0:` | Holy 13, Protection 31 |
| paladin/retribution | `FS1:1.60.1.69893:paladin:human:0/50005/552253311010000021:` | Protection 10, Retribution 31 |
| priest/discipline | `FS1:1.60.1.69893:priest:gnome:52533310130010103/33554/0:` | Discipline 31, Holy 20 |
| priest/holy | `FS1:1.60.1.69893:priest:gnome:500003/33555000030021031/0:` | Discipline 8, Holy 31 |
| priest/shadow | `FS1:1.60.1.69893:priest:gnome:5230000003/0/555100500001300251:` | Discipline 13, Shadow 33 |
| rogue/assassination | `FS1:1.60.1.69893:rogue:night-elf:32500000551501051/3252/51:` | Assassination 33, Combat 12, Subtlety 6 |
| rogue/combat | `FS1:1.60.1.69893:rogue:night-elf:32531/32531300000515201/51:` | Assassination 14, Combat 31, Subtlety 6 |
| rogue/subtlety | `FS1:1.60.1.69893:rogue:night-elf:321/32003/532322131000300105:` | Assassination 6, Combat 8, Subtlety 31 |
| shaman/elemental | `FS1:1.60.1.69893:shaman:dwarf:5532300500103031/0/55334:` | Elemental 31, Restoration 20 |
| shaman/enhancement | `FS1:1.60.1.69893:shaman:dwarf:05/254030030005102051/0:` | Elemental 5, Enhancement 31 |
| shaman/restoration | `FS1:1.60.1.69893:shaman:dwarf:0/005/5532500010513001:` | Enhancement 5, Restoration 31 |
| warlock/affliction | `FS1:1.60.1.69893:warlock:gnome:25552000030201051/235523/0:` | Affliction 31, Demonology 20 |
| warlock/demonology | `FS1:1.60.1.69893:warlock:gnome:055/2355003001200001351/0:` | Affliction 10, Demonology 31 |
| warlock/destruction | `FS1:1.60.1.69893:warlock:gnome:055/0/2555005100101051:` | Affliction 10, Destruction 31 |
| warrior/arms | `FS1:1.60.1.69893:warrior:human:35325213032010001/0505/5005:` | Arms 31, Fury 10, Protection 10 |
| warrior/fury | `FS1:1.60.1.69893:warrior:human:35310003002/555000005050010051/0:` | Arms 17, Fury 32 |
| warrior/protection | `FS1:1.60.1.69893:warrior:dwarf:003/055/552531233000010001:` | Arms 3, Fury 10, Protection 31 |

Some of these differ slightly from a guide's own loose prose ("roughly 31", "a few points"):
where the real tier-gate cost of a named capstone or secondary pick is higher or lower than
the guide's own rough words implied, Task 11 adjusts that one sentence to the real number
(e.g. "roughly 32 points" or naming one fewer secondary talent) — never the other way
around; the code is authoritative because it is the thing the tree actually renders. Task 11
lists the exact wording change per guide.

### `recommendedRaces` (slugs, in the order the guide names them — first is the code's own `raceSlug`)

| Guide | recommendedRaces |
|---|---|
| druid/balance, druid/feral, druid/restoration | `[night-elf, tauren]` |
| hunter/beast-mastery, hunter/marksmanship, hunter/survival | `[dwarf, troll]` |
| mage/arcane, mage/fire | `[gnome, orc]` |
| mage/frost | `[gnome, troll]` |
| paladin/holy | `[human, undead]` |
| paladin/protection | `[dwarf, undead]` |
| paladin/retribution | `[human, dwarf, undead]` |
| priest/discipline | `[human, undead]` |
| priest/holy | `[gnome, troll]` |
| priest/shadow | `[gnome, undead]` |
| rogue/assassination, rogue/combat, rogue/subtlety | `[night-elf, troll]` |
| shaman/elemental, shaman/enhancement | `[dwarf, orc]` |
| shaman/restoration | `[dwarf, tauren]` |
| warlock/affliction, warlock/demonology, warlock/destruction | `[gnome, troll]` |
| warrior/arms | `[human, orc]` |
| warrior/fury | `[human, troll]` |
| warrior/protection | `[dwarf, tauren]` |

### `statPriority` (ordered stat-name strings, taken verbatim from each guide's existing bold terms)

| Guide | statPriority |
|---|---|
| druid/balance | `[Spell power, Intellect, Critical strike, Hit, Spell haste, Spell penetration, "Nature damage and Arcane damage"]` |
| druid/feral | `[Attack power, "Feral-specific attack power", Strength, Agility, Critical strike, Hit, Melee haste]` |
| druid/restoration | `[Healing power, Spell power, Spirit, MP5, Intellect, Critical strike]` |
| hunter/beast-mastery, hunter/marksmanship, hunter/survival | `[Attack power, "Ranged attack power", Agility, Critical strike, Hit, Melee haste]` |
| mage/arcane | `[Spell power, Intellect, Critical strike, Hit, Spell haste, Spell penetration, "Arcane damage"]` |
| mage/fire | `[Spell power, Intellect, Critical strike, Hit, Spell haste, Spell penetration, "Fire damage"]` |
| mage/frost | `[Spell power, Intellect, Critical strike, Hit, Spell haste, Spell penetration, "Frost damage"]` |
| paladin/holy, priest/discipline, priest/holy, shaman/restoration | `[Healing power, Spell power, Spirit, MP5, Intellect, Critical strike]` |
| paladin/protection, paladin/retribution, shaman/enhancement, warrior/arms, warrior/fury, warrior/protection | `[Attack power, Strength, Agility, Critical strike, Hit, Melee haste]` |
| priest/shadow | `[Spell power, Intellect, Critical strike, Hit, Spell haste, Spell penetration, "Shadow power"]` |
| rogue/assassination, rogue/combat, rogue/subtlety | `[Attack power, Agility, Critical strike, Hit, Melee haste]` |
| shaman/elemental | `[Spell power, Intellect, Critical strike, Hit, Spell haste, Spell penetration, "Nature power"]` |
| warlock/affliction | `[Spell power, Intellect, Critical strike, Hit, Spell haste, Spell penetration, "Shadow damage"]` |
| warlock/demonology, warlock/destruction | `[Spell power, Intellect, Critical strike, Hit, Spell haste, Spell penetration, "Shadow damage and Fire damage"]` |

---

## Task 1: `splitSpecSections` — one shared section splitter

**Files:**
- Modify: `web/src/lib/guides/sections.ts`
- Modify: `web/src/content/guides/_sections.test.ts` (use the new function instead of its own inline regex)
- Test: `web/src/lib/guides/sections.test.ts` (new)

**Interfaces:**
- Produces: `splitSpecSections(body: string): Map<SpecSection, string>` — one entry per
  `SPEC_SECTIONS` heading actually present, keyed by the exact heading text, value is the
  raw markdown from that `## Heading` line (inclusive) up to (not including) the next `##`
  heading, insertion-ordered to match the body's own heading order. A heading in the body
  that is not one of `SPEC_SECTIONS` is skipped (its text stays folded into the preceding
  section's slice, exactly as today's TOC/order test already tolerates).

- [ ] **Step 1: Write the failing test**

```typescript
// web/src/lib/guides/sections.test.ts
import { describe, expect, it } from 'vitest';
import { SPEC_SECTIONS, splitSpecSections } from './sections';

const BODY = `## Overview

Some overview text.

## Talents and builds

Pick these talents.

- one
- two

## Stat priority

1. **Attack power** — text.
`;

describe('splitSpecSections', () => {
  it('slices each heading through to the next one, in document order', () => {
    const sections = splitSpecSections(BODY);
    expect([...sections.keys()]).toEqual(['Overview', 'Talents and builds', 'Stat priority']);
    expect(sections.get('Overview')).toBe('## Overview\n\nSome overview text.');
    expect(sections.get('Talents and builds')).toBe(
      '## Talents and builds\n\nPick these talents.\n\n- one\n- two',
    );
    expect(sections.get('Stat priority')).toBe('## Stat priority\n\n1. **Attack power** — text.');
  });

  it('returns an empty map for a body with no recognised heading', () => {
    expect(splitSpecSections('just some text').size).toBe(0);
  });

  it('never returns a key outside SPEC_SECTIONS', () => {
    const sections = splitSpecSections('## Not a real section\n\nx\n\n## Overview\n\ny');
    expect([...sections.keys()]).toEqual(['Overview']);
  });
});
```

- [ ] **Step 2: Run it to see it fail**

Run: `cd web && npx vitest run src/lib/guides/sections.test.ts`
Expected: FAIL — `splitSpecSections` is not exported.

- [ ] **Step 3: Implement `splitSpecSections`**

Add to `web/src/lib/guides/sections.ts`, after `slugify`:

```typescript
/**
 * Every `## `-prefixed section a guide's raw markdown body carries, sliced from its own
 * heading line through to (not including) the next `## ` heading, keyed by heading text and
 * insertion-ordered to match the body. `[spec].astro` renders each slice through its own
 * markdown pass and plants a component between two of them; `_sections.test.ts` uses this
 * same split so the page and the test can never drift on what counts as "a section".
 * A heading that is not one of SPEC_SECTIONS is not a key here -- its text stays folded into
 * whichever recognised section precedes it, the same tolerance the order test already had.
 */
export function splitSpecSections(body: string): Map<SpecSection, string> {
  const headingMatches = [...body.matchAll(/^##[ \t]+.+$/gm)];
  const sections = new Map<SpecSection, string>();
  for (const [index, match] of headingMatches.entries()) {
    const heading = match[0].replace(/^##[ \t]+/, '').trim();
    if (!(SPEC_SECTIONS as readonly string[]).includes(heading)) continue;
    const start = match.index ?? 0;
    const end = headingMatches[index + 1]?.index ?? body.length;
    sections.set(heading as SpecSection, body.slice(start, end).trimEnd());
  }
  return sections;
}
```

- [ ] **Step 4: Run it to see it pass**

Run: `cd web && npx vitest run src/lib/guides/sections.test.ts`
Expected: PASS (3 tests).

- [ ] **Step 5: Refactor `_sections.test.ts` to use it (no behaviour change, less duplication)**

Replace the body of the `'renders all nine sections in order'` test in
`web/src/content/guides/_sections.test.ts`:

```typescript
  it.each(specGuides.map((g) => [g.id, g] as const))('%s renders all nine sections in order', (id, guide) => {
    const sections = splitSpecSections(guide.body);
    for (const section of SPEC_SECTIONS) {
      expect(sections.has(section), `missing section "${section}" in ${id}`).toBe(true);
    }
    expect([...sections.keys()], `sections out of order in ${id}`).toEqual(
      SPEC_SECTIONS.filter((section) => sections.has(section)),
    );
  });
```

Add `import { SPEC_SECTIONS, splitSpecSections } from '../../lib/guides/sections';` next to
the existing `SPEC_SECTIONS` import (replace it, don't duplicate).

- [ ] **Step 6: Run the full guides test suite**

Run: `cd web && npx vitest run src/lib/guides/sections.test.ts src/content/guides/_sections.test.ts`
Expected: PASS, same pass count as before the refactor (no regression).

- [ ] **Step 7: Scoped checks and commit**

```bash
cd web && npx astro check && npm run lint && npx prettier --check src/lib/guides/sections.ts src/lib/guides/sections.test.ts src/content/guides/_sections.test.ts
```

```bash
printf 'refactor(web): share the guide section splitter between the TOC test and the page\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\nClaude-Session: https://claude.ai/code/session_01DAHb7A9Ytd6qpZu2UujXgY\n' > /tmp/commit-msg-1.txt
```

```bash
git add web/src/lib/guides/sections.ts web/src/lib/guides/sections.test.ts web/src/content/guides/_sections.test.ts
git commit -F /tmp/commit-msg-1.txt
```

---

## Task 2: `renderGuideMarkdown` — markdown-to-HTML for one section at a time

**Files:**
- Create: `web/src/lib/guides/markdown.ts`
- Test: `web/src/lib/guides/markdown.test.ts`

**Interfaces:**
- Consumes: nothing from earlier tasks.
- Produces: `renderGuideMarkdown(markdown: string): Promise<string>` — renders one markdown
  string to an HTML string via the same processor Astro's own content pipeline uses
  (`@astrojs/markdown-satteri`), memoizing the (expensive, async) processor creation so 9
  sections × 27 pages share one instance per build. `''` in, `''` out (no processor call for
  an empty section, e.g. a heading-only split with no prose after it).

- [ ] **Step 1: Write the failing test**

```typescript
// web/src/lib/guides/markdown.test.ts
import { describe, expect, it } from 'vitest';
import { renderGuideMarkdown } from './markdown';

describe('renderGuideMarkdown', () => {
  it('renders a heading with the same id slugify() would produce', async () => {
    const html = await renderGuideMarkdown('## Stat priority');
    expect(html).toContain('<h2 id="stat-priority">Stat priority</h2>');
  });

  it('renders prose and a numbered list', async () => {
    const html = await renderGuideMarkdown('In priority order:\n\n1. **Attack power** — text.\n');
    expect(html).toContain('<strong>Attack power</strong>');
    expect(html).toContain('<ol>');
  });

  it('returns an empty string for empty input without calling the processor', async () => {
    expect(await renderGuideMarkdown('')).toBe('');
    expect(await renderGuideMarkdown('   \n')).toBe('');
  });
});
```

- [ ] **Step 2: Run it to see it fail**

Run: `cd web && npx vitest run src/lib/guides/markdown.test.ts`
Expected: FAIL — module does not exist.

- [ ] **Step 3: Implement it**

```typescript
// web/src/lib/guides/markdown.ts
// Renders one guide section's raw markdown to an HTML string, independent of Astro's whole-
// entry `render(entry)`. [spec].astro needs this because a spec guide's markdown carries no
// MDX (not installed, and this worktree may never `npm install`) -- so a Svelte island or an
// Astro component cannot live inside the markdown itself. Splitting the raw body into its
// nine `## ` sections (sections.ts's own splitSpecSections) and rendering each one separately
// through this is what lets the page plant a real component between two of them while every
// section keeps the exact heading markup (id, tag) Astro's own pipeline would have produced,
// since this is the very processor that pipeline uses internally.
import { createSatteriMarkdownProcessor } from '@astrojs/markdown-satteri';

type Processor = Awaited<ReturnType<typeof createSatteriMarkdownProcessor>>;
let processor: Promise<Processor> | null = null;

function sharedProcessor(): Promise<Processor> {
  if (processor === null) processor = createSatteriMarkdownProcessor();
  return processor;
}

export async function renderGuideMarkdown(markdown: string): Promise<string> {
  if (markdown.trim() === '') return '';
  const { render } = await sharedProcessor();
  const { code } = await render(markdown);
  return code;
}
```

- [ ] **Step 4: Run it to see it pass**

Run: `cd web && npx vitest run src/lib/guides/markdown.test.ts`
Expected: PASS (3 tests).

- [ ] **Step 5: Scoped checks and commit**

```bash
cd web && npx astro check && npm run lint && npx prettier --check src/lib/guides/markdown.ts src/lib/guides/markdown.test.ts
```

```bash
printf 'feat(web): render one guide section at a time for component-planting\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\nClaude-Session: https://claude.ai/code/session_01DAHb7A9Ytd6qpZu2UujXgY\n' > /tmp/commit-msg-2.txt
```

```bash
git add web/src/lib/guides/markdown.ts web/src/lib/guides/markdown.test.ts
git commit -F /tmp/commit-msg-2.txt
```

---

## Task 3: Schema fields + the SSR decode test

**Files:**
- Modify: `web/src/content.config.ts`
- Modify: `web/src/content/guides/_sections.test.ts`

**Interfaces:**
- Produces: `guideSchema` gains three fields, all present only on a spec guide (a class
  landing page omits them, same convention as `spec`/`role`): `build: z.string().optional()`,
  `recommendedRaces: z.array(z.string()).default([])`, `statPriority: z.array(z.string()).default([])`.
- Produces: a new `it.each` block in `_sections.test.ts` — the SSR test spec section 9 asks
  for — that decodes every spec guide's `build:` with the real `decodeFS1`, checks its class
  matches the guide's own `classSlug`, and checks the whole build is *legally reachable*
  (every point's tier gate and prerequisite satisfied — a stronger, `dropped`-based version
  of the coordinator's "total points ≤ 51", since a code that decodes with unreachable ranks
  is exactly the kind of broken build the coordinator's rule exists to catch).

- [ ] **Step 1: Write the failing test**

Add to `web/src/content/guides/_sections.test.ts` (new imports:
`import { decodeFS1, orderFromRanks } from '../../lib/planner/fs1';`,
`import { indexTalents } from '../../lib/planner/rules';`,
`import { MAX_POINTS, type TalentFile } from '../../lib/planner/types';`,
`import { readFileSync } from 'node:fs';`, `import { fileURLToPath } from 'node:url';`,
`import { join } from 'node:path';`):

```typescript
const DATA_BUILD = '1.60.1.69893';
const dataRoot = join(dirname(fileURLToPath(import.meta.url)), '../../../../data/builds', DATA_BUILD, 'talents');
const talentCache = new Map<string, TalentFile>();
function talentsFor(classSlug: string): TalentFile {
  const cached = talentCache.get(classSlug);
  if (cached) return cached;
  const file = JSON.parse(readFileSync(join(dataRoot, `${classSlug}.json`), 'utf8')) as TalentFile;
  talentCache.set(classSlug, file);
  return file;
}

describe('guide build codes', () => {
  it.each(specGuides.map((g) => [g.id, g] as const))(
    '%s has a build: that decodes, names its own class, and is legally reachable within 51 points',
    (id, guide) => {
      const code = guide.frontmatter.build;
      expect(typeof code, `${id}: no build: frontmatter`).toBe('string');
      const decoded = decodeFS1(code as string);
      expect(decoded.ok, `${id}: ${decoded.ok ? '' : decoded.message}`).toBe(true);
      if (!decoded.ok) return;
      expect(decoded.build.classSlug, `${id}: build: is for the wrong class`).toBe(guide.frontmatter.classSlug);
      const index = indexTalents(talentsFor(decoded.build.classSlug));
      const { order, dropped } = orderFromRanks(index, decoded.build.treeRanks);
      expect(dropped, `${id}: unreachable talents in build:`).toEqual([]);
      expect(order.length, `${id}: build: spends more than the ${MAX_POINTS}-point cap`).toBeLessThanOrEqual(
        MAX_POINTS,
      );
    },
  );
});
```

- [ ] **Step 2: Run it to see it fail**

Run: `cd web && npx vitest run src/content/guides/_sections.test.ts`
Expected: FAIL — every spec guide's `build.frontmatter` is `undefined` (schema has no field
yet, and no guide file has been touched yet either); the `typeof code` assertion fails first.

- [ ] **Step 3: Extend the schema**

In `web/src/content.config.ts`, extend `guideSchema`:

```typescript
export const guideSchema = factSchema.extend({
  classSlug: z.string(),
  spec: z.string().optional(),
  role: z.enum(['dps', 'healer', 'tank']).optional(),
  /** An FS1 code for this spec's recommended build (spec guides only). `[spec].astro`
   *  decodes it to light the embedded read-only tree and to build the "Load this build" /
   *  "Sim this build" links; `_sections.test.ts`'s own SSR test decodes every one and checks
   *  it names this guide's own class and stays legally reachable within 51 points. */
  build: z.string().optional(),
  /** Race slugs (races.json's own `slug`) this guide calls a strong pick, in the order the
   *  guide's own Races prose names them. RacePillRow.astro marks these among the class's
   *  full legal race list; an empty array (the default, and every class landing page's
   *  value) marks none. */
  recommendedRaces: z.array(z.string()).default([]),
  /** Stat names in priority order, exactly as this guide's own Stat priority prose already
   *  names them (its bold terms) -- StatPriorityPills.astro renders these as an ordered pill
   *  row; the guide's own prose stays underneath as the reasoning, unchanged. */
  statPriority: z.array(z.string()).default([]),
});
```

- [ ] **Step 4: Run it again to confirm the *shape* of the failure changes**

Run: `cd web && npx vitest run src/content/guides/_sections.test.ts`
Expected: still FAIL (no guide file has `build:` yet), but now on the `typeof code`
assertion for a reason consistent with "no data yet", not a schema/type error. This confirms
the schema change alone didn't break anything else in the file (the earlier "schema-valid
frontmatter" test in the same file should still pass, since every field just added is
optional/defaulted).

- [ ] **Step 5: Scoped checks and commit**

```bash
cd web && npx astro check && npm run lint && npx prettier --check src/content.config.ts src/content/guides/_sections.test.ts
```

```bash
printf 'feat(web): schema fields and the SSR decode test for guide build codes\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\nClaude-Session: https://claude.ai/code/session_01DAHb7A9Ytd6qpZu2UujXgY\n' > /tmp/commit-msg-3.txt
```

```bash
git add web/src/content.config.ts web/src/content/guides/_sections.test.ts
git commit -F /tmp/commit-msg-3.txt
```

(This test will only turn green once Task 11 fills in every guide's frontmatter — that is
expected and fine to leave red across intermediate commits in this plan; the ledger notes it,
and Task 11's own step re-runs this exact test as its acceptance check.)

---

## Task 4: `TreeGrid`/`TalentCell` read-only mode

**Files:**
- Modify: `web/src/components/planner/TreeGrid.svelte`
- Modify: `web/src/components/planner/TalentCell.svelte`
- Test: `web/src/components/planner/TreeGrid.test.ts` (new)
- Test: `web/src/components/planner/TalentCell.test.ts` (new)

**Interfaces:**
- Produces: `TreeGrid` gains an optional `readOnly?: boolean` prop (default `false`,
  preserving today's planner behaviour exactly); when `true` it passes `readOnly` to every
  `TalentCell` and does not attach its own roving-tabindex keyboard handler (nothing is
  focusable in read-only mode, so there is nothing to move focus between).
- Produces: `TalentCell` gains an optional `readOnly?: boolean` prop (default `false`); when
  `true` it renders a plain, non-interactive `<div>` in place of the `<button>` (no
  `onclick`/`oncontextmenu`/`onpointerdown`/`onpointerup`/`onpointercancel`/hover handlers, no
  tooltip, no tabindex), showing exactly the same icon-and-rank-pill visual the interactive
  cell shows, driven by the same `store.ranks`/`talent.max_rank` the interactive path already
  reads — never a second source of truth for what "rank N of M" means.

- [ ] **Step 1: Write the failing tests**

```typescript
// web/src/components/planner/TalentCell.test.ts
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import fixtureTalents from '../../fixtures/planner/talents/warrior.json';
import { createPlannerStore } from '../../lib/planner/store.svelte';
import { indexTalents } from '../../lib/planner/rules';
import type { TalentFile } from '../../lib/planner/types';
import TalentCell from './TalentCell.svelte';

const talents = indexTalents(fixtureTalents as TalentFile);
const arms = talents.trees[0];
const talent = arms.talents[0];

function storeWithRank(rank: number) {
  const store = createPlannerStore({ treeVersion: 'test', classSlug: 'warrior', raceSlug: 'human' });
  store.setTalents(fixtureTalents as TalentFile);
  for (let i = 0; i < rank; i++) store.addPoint(talent.id);
  return store;
}

describe('TalentCell read-only mode', () => {
  it('renders a div, not a button, when readOnly', () => {
    const store = storeWithRank(1);
    const { body } = render(TalentCell, {
      props: { store, talent, focused: false, onfocuscell: () => {}, readOnly: true },
    });
    expect(body).not.toContain('<button');
    expect(body).toContain(`data-testid="talent-${talent.id}"`);
  });

  it('still shows the real rank and max rank when readOnly', () => {
    const store = storeWithRank(2);
    const { body } = render(TalentCell, {
      props: { store, talent, focused: false, onfocuscell: () => {}, readOnly: true },
    });
    expect(body).toContain(`${2}/${talent.max_rank}`);
  });

  it('renders the normal interactive button when readOnly is left out', () => {
    const store = storeWithRank(0);
    const { body } = render(TalentCell, { props: { store, talent, focused: false, onfocuscell: () => {} } });
    expect(body).toContain('<button');
  });
});
```

```typescript
// web/src/components/planner/TreeGrid.test.ts
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import fixtureTalents from '../../fixtures/planner/talents/warrior.json';
import { createPlannerStore } from '../../lib/planner/store.svelte';
import type { TalentFile } from '../../lib/planner/types';
import TreeGrid from './TreeGrid.svelte';

const talentFile = fixtureTalents as TalentFile;
const tree = talentFile.trees[0];

function readyStore() {
  const store = createPlannerStore({ treeVersion: 'test', classSlug: 'warrior', raceSlug: 'human' });
  store.setTalents(talentFile);
  return store;
}

describe('TreeGrid read-only mode', () => {
  it('renders no <button> elements when readOnly', () => {
    const { body } = render(TreeGrid, { props: { store: readyStore(), tree, readOnly: true } });
    expect(body).not.toContain('<button');
  });

  it('still renders one cell per talent', () => {
    const { body } = render(TreeGrid, { props: { store: readyStore(), tree, readOnly: true } });
    for (const talent of tree.talents) {
      expect(body).toContain(`data-testid="talent-${talent.id}"`);
    }
  });
});
```

- [ ] **Step 2: Run them to see them fail**

Run: `cd web && npx vitest run src/components/planner/TalentCell.test.ts src/components/planner/TreeGrid.test.ts`
Expected: FAIL — `readOnly` prop does not exist yet, `<button>` always renders.

- [ ] **Step 3: Implement `TalentCell.svelte`'s read-only branch**

Add `readOnly = false` to the props destructure:

```svelte
  let {
    store,
    talent,
    focused,
    onfocuscell,
    readOnly = false,
  }: {
    store: PlannerStore;
    talent: Talent;
    focused: boolean;
    onfocuscell: () => void;
    readOnly?: boolean;
  } = $props();
```

Change the `available` derivation so a read-only cell never claims to be "available to add a
point to" (nothing is addable when nothing is clickable):

```svelte
  const available = $derived(
    !readOnly && store.talentIndex !== null && canAddPoint(store.talentIndex, store.order, talent.id).ok,
  );
```

Wrap the markup so the two render paths share the icon/pill snippet (DRY — one source for
what a cell's face looks like) and only the outer element and its interactivity differ:

```svelte
{#snippet face()}
  {#if iconBroken}
    <span class="text-muted font-display text-[13px] font-bold" aria-hidden="true">
      {talent.name.slice(0, 2)}
    </span>
  {:else}
    <img
      src={iconSrc}
      alt=""
      width="40"
      height="40"
      loading="lazy"
      decoding="async"
      class="rounded-control h-10 w-10 object-cover"
      onerror={() => (iconBroken = true)}
    />
  {/if}
  <span
    class={`tabular rounded-pill bg-bg absolute -right-1 -bottom-1 border px-1 font-mono text-[11px] leading-[14px] ${CELL_PILL[state]}`}
  >
    {rank}/{talent.max_rank}
  </span>
{/snippet}

<div class="relative">
  {#if readOnly}
    <div
      aria-label={`${talent.name}, rank ${rank} of ${talent.max_rank}`}
      data-testid={`talent-${talent.id}`}
      data-rank={rank}
      data-state={state}
      class={`rounded-control bg-card-top relative flex h-11 w-11 items-center justify-center border md:h-12 md:w-12 ${CELL_BORDER[state]}`}
    >
      {@render face()}
    </div>
  {:else}
    <button
      type="button"
      tabindex={focused ? 0 : -1}
      aria-describedby={open ? tooltipId : undefined}
      aria-label={`${talent.name}, rank ${rank} of ${talent.max_rank}`}
      data-testid={`talent-${talent.id}`}
      data-rank={rank}
      data-state={state}
      class={`rounded-control bg-card-top relative flex h-11 w-11 items-center justify-center border md:h-12 md:w-12 ${CELL_BORDER[state]}`}
      onclick={add}
      oncontextmenu={removeOnContextMenu}
      onfocus={() => {
        open = true;
        onfocuscell();
      }}
      onblur={() => (open = false)}
      onmouseenter={() => (open = true)}
      onmouseleave={() => {
        open = false;
        abandonPress();
      }}
      onpointerdown={startPress}
      onpointerup={disarmPress}
      onpointercancel={abandonPress}
    >
      {@render face()}
    </button>

    {#if open}
      <div
        id={tooltipId}
        role="tooltip"
        bind:this={tip}
        style:left={`${shift}px`}
        class="border-line bg-raised rounded-panel absolute top-full z-30 mt-2 flex w-[260px] flex-col gap-2 border p-3 shadow-[0_12px_30px_rgba(0,0,0,.45)]"
      >
        <span class="text-strong font-display text-[14px] font-bold">{talent.name}</span>
        <span class="tabular text-muted font-mono text-[12px]">
          Rank {rank} of {talent.max_rank}
        </span>
        {#if rank > 0}
          <p class="text-text text-[13px] leading-snug">{talent.ranks[rank - 1].description}</p>
        {/if}
        {#if rank < talent.max_rank}
          <p class="text-muted text-[13px] leading-snug">
            <span class="label text-muted">Next rank</span>
            {talent.ranks[rank].description}
          </p>
        {/if}
      </div>
    {/if}
  {/if}
</div>
```

(Keep every existing `<script>` block line as-is — `open`/`tip`/`shift`/`iconBroken`/press
handlers are untouched; only the props destructure, `available`, and the markup below change.)

- [ ] **Step 4: Implement `TreeGrid.svelte`'s read-only branch**

Add `readOnly = false` to its props and thread it through, and only attach the keyboard
handler when interactive (nothing is focusable in read-only mode, so there is nothing to
move focus between — keeping the roving-tabindex `focusedId`/`focused` state harmless but
unused is fine and simpler than deleting it, since it is a plain `$state`/`$derived` with no
side effect on its own):

```svelte
  let { store, tree, readOnly = false }: { store: PlannerStore; tree: TalentTree; readOnly?: boolean } = $props();
```

```svelte
    <div
      bind:this={root}
      role="grid"
      tabindex={-1}
      aria-label={`${tree.name} talents`}
      data-testid={`tree-${tree.id}`}
      class="relative grid gap-2"
      style={`grid-template-columns: repeat(${size.columns}, minmax(0, max-content));`}
      onkeydown={readOnly ? undefined : onKeyDown}
    >
      {#each Array.from({ length: size.tiers }, (_, tier) => tier) as tier (tier)}
        <div role="row" class="contents">
          {#each Array.from({ length: size.columns }, (_, column) => column) as column (column)}
            {@const cell = cells.find((c) => c.tier === tier && c.column === column)}
            <div role="gridcell" class="flex h-11 w-11 md:h-12 md:w-12">
              {#if cell}
                <TalentCell
                  {store}
                  talent={cell.talent}
                  focused={focused?.talent.id === cell.talent.id}
                  onfocuscell={() => (focusedId = cell.talent.id)}
                  {readOnly}
                />
              {/if}
            </div>
          {/each}
        </div>
      {/each}
    </div>
```

- [ ] **Step 5: Run the tests to see them pass**

Run: `cd web && npx vitest run src/components/planner/TalentCell.test.ts src/components/planner/TreeGrid.test.ts`
Expected: PASS (5 tests total).

- [ ] **Step 6: Run the existing planner suite to confirm no regression**

Run: `cd web && npx vitest run src/components/planner/ src/lib/planner/`
Expected: PASS, same count as on `main` for these paths (the interactive `else` branch is
untouched markup, just re-indented under `{#else}`).

- [ ] **Step 7: Scoped checks and commit**

```bash
cd web && npx astro check && npm run lint && npx prettier --check src/components/planner/TreeGrid.svelte src/components/planner/TalentCell.svelte src/components/planner/TreeGrid.test.ts src/components/planner/TalentCell.test.ts
```

```bash
printf 'feat(web): a read-only mode for TreeGrid/TalentCell, for the guides tree embed\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\nClaude-Session: https://claude.ai/code/session_01DAHb7A9Ytd6qpZu2UujXgY\n' > /tmp/commit-msg-4.txt
```

```bash
git add web/src/components/planner/TreeGrid.svelte web/src/components/planner/TalentCell.svelte web/src/components/planner/TreeGrid.test.ts web/src/components/planner/TalentCell.test.ts
git commit -F /tmp/commit-msg-4.txt
```

**Ruling:** `TalentCell.svelte` is not in this lane's ownership list (only "a read-only mode
on `TreeGrid.svelte`" is named in spec section 8), but the click/hover/tooltip behaviour spec
5 says the embed must not have ("no clicks, no hover state") lives entirely in `TalentCell`,
not `TreeGrid` — there is no way to satisfy spec 5 by touching `TreeGrid` alone. No other
lane's file list claims `TalentCell.svelte` either. Cost if wrong: a merge conflict with
whichever lane next touches this specific file, resolved the normal way; low risk since the
change is a strictly additive, default-`false` prop.

---

## Task 5: `lib/guides/race-pills.ts` and `lib/guides/build-links.ts`

**Files:**
- Create: `web/src/lib/guides/race-pills.ts`
- Test: `web/src/lib/guides/race-pills.test.ts`
- Create: `web/src/lib/guides/build-links.ts`
- Test: `web/src/lib/guides/build-links.test.ts`

**Interfaces:**
- Produces: `interface GuideRacePill { slug: string; name: string; faction: string; recommended: boolean }`
  and `racePillsFor(classSlug: string, recommendedRaces: readonly string[]): GuideRacePill[]`
  — every race legal for the class (`racesForClass`, already exported from
  `lib/planner/reference.ts`), each marked `recommended` when its slug is in
  `recommendedRaces`. Returns `[]` for an unknown class slug (fails safe, never throws on
  guide content).
- Produces: `factionColorVar(faction: string): string` — `var(--color-alliance)` for
  `'alliance'`, `var(--color-horde)` for `'horde'`, `var(--color-text)` otherwise (mirrors
  `report/format.ts`'s own `classColorVar` fallback pattern, not reused directly since that
  file is a different lane's -- `ReportView.svelte`'s -- and already near its own 800-line
  ceiling per that file's own header comment).
- Produces: `loadBuildHref(code: string): string` — `/planner?code=<encoded>`.
- Produces: `simBuildHref(code: string): string` — built from `lib/sim/tabs.ts`'s
  `SIM_TABS`/`tabHref` and `lib/sim/url.ts`'s `defaultSimState`/`withSimState` (that module's
  own comment: a `?code=` link is "never built by hand"), landing on the quick-sim tab.

- [ ] **Step 1: Write the failing tests**

```typescript
// web/src/lib/guides/race-pills.test.ts
import { describe, expect, it } from 'vitest';
import { factionColorVar, racePillsFor } from './race-pills';

describe('racePillsFor', () => {
  it('lists every race legal for the class, marking the recommended ones', () => {
    const pills = racePillsFor('warrior', ['human', 'troll']);
    const slugs = pills.map((p) => p.slug);
    expect(slugs).toContain('human');
    expect(slugs).toContain('orc');
    expect(pills.find((p) => p.slug === 'human')?.recommended).toBe(true);
    expect(pills.find((p) => p.slug === 'orc')?.recommended).toBe(false);
  });

  it('returns an empty list for an unknown class slug rather than throwing', () => {
    expect(racePillsFor('not-a-class', [])).toEqual([]);
  });
});

describe('factionColorVar', () => {
  it('maps alliance and horde to their tokens, and anything else to plain text', () => {
    expect(factionColorVar('alliance')).toBe('var(--color-alliance)');
    expect(factionColorVar('horde')).toBe('var(--color-horde)');
    expect(factionColorVar('neutral')).toBe('var(--color-text)');
  });
});
```

```typescript
// web/src/lib/guides/build-links.test.ts
import { describe, expect, it } from 'vitest';
import { loadBuildHref, simBuildHref } from './build-links';

describe('loadBuildHref', () => {
  it('links to /planner?code=, URL-encoded', () => {
    expect(loadBuildHref('FS1:1.60.1.69893:warrior:human:1/2/3:')).toBe(
      `/planner?code=${encodeURIComponent('FS1:1.60.1.69893:warrior:human:1/2/3:')}`,
    );
  });
});

describe('simBuildHref', () => {
  it('links to the quick-sim tab with ?code=', () => {
    const href = simBuildHref('FS1:1.60.1.69893:warrior:human:1/2/3:');
    expect(href.startsWith('/sim?')).toBe(true);
    expect(new URLSearchParams(href.split('?')[1]).get('code')).toBe('FS1:1.60.1.69893:warrior:human:1/2/3:');
  });
});
```

- [ ] **Step 2: Run them to see them fail**

Run: `cd web && npx vitest run src/lib/guides/race-pills.test.ts src/lib/guides/build-links.test.ts`
Expected: FAIL — neither module exists.

- [ ] **Step 3: Implement `race-pills.ts`**

```typescript
// web/src/lib/guides/race-pills.ts
// A spec guide's Races section renders a row of name pills (no race portraits exist --
// design/DESIGN-SYSTEM.md and the current-character spec both say so) rather than a second
// prose enumeration of the class's legal races; this is the one place that list is derived,
// from the same build-time reference data reference.ts already exposes to the classes page.
import { classRows, racesForClass } from '../planner/reference';

export interface GuideRacePill {
  slug: string;
  name: string;
  faction: string;
  recommended: boolean;
}

export function racePillsFor(classSlug: string, recommendedRaces: readonly string[]): GuideRacePill[] {
  const classRow = classRows.find((row) => row.slug === classSlug);
  if (!classRow) return [];
  return racesForClass(classRow.id).map((race) => ({
    slug: race.slug,
    name: race.name,
    faction: race.faction,
    recommended: recommendedRaces.includes(race.slug),
  }));
}

const FACTION_COLORS: Record<string, string> = {
  alliance: 'var(--color-alliance)',
  horde: 'var(--color-horde)',
};

export function factionColorVar(faction: string): string {
  return FACTION_COLORS[faction] ?? 'var(--color-text)';
}
```

- [ ] **Step 4: Implement `build-links.ts`**

```typescript
// web/src/lib/guides/build-links.ts
// The two links under a spec guide's embedded tree (spec 5): "Load this build" is a plain
// /planner?code= link (current-character.ts's own plannerHrefFor builds the identical shape
// for a 'code' pointer); "Sim this build" goes through lib/sim/url.ts's own state builder
// (that module's header comment: a `?code=` link there is "never built by hand"). Neither
// link writes the current-character pointer itself -- Planner.svelte's writePlannerPointer
// and sim's sources.ts already do that unconditionally on any `?code=` load, so this needs no
// session read of its own (Global Constraint 3 stays satisfied by having nothing to read).
import { SIM_TABS, tabHref } from '../sim/tabs';
import { defaultSimState, withSimState } from '../sim/url';

export function loadBuildHref(code: string): string {
  return `/planner?code=${encodeURIComponent(code)}`;
}

export function simBuildHref(code: string): string {
  return tabHref(SIM_TABS[0].href, withSimState(defaultSimState(), { code }));
}
```

- [ ] **Step 5: Run the tests to see them pass**

Run: `cd web && npx vitest run src/lib/guides/race-pills.test.ts src/lib/guides/build-links.test.ts`
Expected: PASS (4 tests).

- [ ] **Step 6: Scoped checks and commit**

```bash
cd web && npx astro check && npm run lint && npx prettier --check src/lib/guides/race-pills.ts src/lib/guides/race-pills.test.ts src/lib/guides/build-links.ts src/lib/guides/build-links.test.ts
```

```bash
printf 'feat(web): race-pill and build-link helpers for the guides lane\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\nClaude-Session: https://claude.ai/code/session_01DAHb7A9Ytd6qpZu2UujXgY\n' > /tmp/commit-msg-5.txt
```

```bash
git add web/src/lib/guides/race-pills.ts web/src/lib/guides/race-pills.test.ts web/src/lib/guides/build-links.ts web/src/lib/guides/build-links.test.ts
git commit -F /tmp/commit-msg-5.txt
```

**Ruling:** spec 5 says "Load this build" should "also write the pointer" for a signed-in
visitor "so Sim and Logs follow." `Planner.svelte`'s own `writePlannerPointer` (called from
its mount effect) and `lib/sim/sources.ts`'s `recordCurrentCharacter` already write the
current-character pointer unconditionally whenever their own page loads with a `?code=` —
confirmed by reading both call sites — regardless of sign-in state (the pointer is a
`localStorage` convenience, not session-gated). So a plain link already produces the
"also writes the pointer" behaviour once the destination page loads; no session read or
sign-in branch is needed on the guide page itself, and `lib/current-character.ts` (Lane B's
file) is not touched. Cost if wrong: the pointer would need writing here too, which would mean
adding a small client-side click handler and importing `writeCurrent`/`readCurrent`-adjacent
helpers — a contained, low-risk follow-up if Task 12's e2e spec (which does click "Load this
build" and check what landed) turns up a gap.

---

## Task 6: Guide copy module

**Files:**
- Create: `web/src/lib/guides/copy.ts`
- Test: `web/src/lib/guides/copy.test.ts`

**Interfaces:**
- Produces: `guidesCopy` — every user-visible string Tasks 7-9's new components need
  (Global Constraint: "every visible string in a copy module").

- [ ] **Step 1: Write the failing test**

```typescript
// web/src/lib/guides/copy.test.ts
import { describe, expect, it } from 'vitest';
import { guidesCopy } from './copy';

describe('guidesCopy', () => {
  it('has no blank strings', () => {
    for (const [key, value] of Object.entries(guidesCopy)) {
      expect(value.trim(), `${key} is blank`).not.toBe('');
    }
  });
});
```

- [ ] **Step 2: Run it to see it fail**

Run: `cd web && npx vitest run src/lib/guides/copy.test.ts`
Expected: FAIL — module does not exist.

- [ ] **Step 3: Implement it**

```typescript
// web/src/lib/guides/copy.ts
// Every user-visible string the guides lane's own components (GuideBuildTree,
// BuildActionButtons, RacePillRow, StatPriorityPills) render, reference voice throughout
// (design/DESIGN-SYSTEM.md: state the thing and stop).
export const guidesCopy = {
  loadThisBuild: 'Load this build',
  simThisBuild: 'Sim this build',
  treeLoadFailed: 'This build did not load',
  recommendedRaceLabel: 'Recommended',
} as const;
```

- [ ] **Step 4: Run it to see it pass**

Run: `cd web && npx vitest run src/lib/guides/copy.test.ts`
Expected: PASS.

- [ ] **Step 5: Scoped checks and commit**

```bash
cd web && npx astro check && npm run lint && npx prettier --check src/lib/guides/copy.ts src/lib/guides/copy.test.ts
```

```bash
printf 'feat(web): a copy module for the guides lane'"'"'s own new UI strings\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\nClaude-Session: https://claude.ai/code/session_01DAHb7A9Ytd6qpZu2UujXgY\n' > /tmp/commit-msg-6.txt
```

```bash
git add web/src/lib/guides/copy.ts web/src/lib/guides/copy.test.ts
git commit -F /tmp/commit-msg-6.txt
```

---

## Task 7: New guide components

**Files:**
- Create: `web/src/components/guides/GuideBuildTree.svelte`
- Test: `web/src/components/guides/GuideBuildTree.test.ts`
- Create: `web/src/components/guides/BuildActionButtons.astro`
- Create: `web/src/components/guides/StatPriorityPills.astro`
- Create: `web/src/components/guides/RacePillRow.astro`
- Test: `web/tests/astro/guides-components.test.ts` — **skip this file if the project has no
  existing Astro-component-render test convention**; check first with
  `grep -rl "AstroContainer" web/src/components web/tests` — if none of the three `.astro`
  components already have direct render tests elsewhere, cover them instead through the SSR
  page test in Task 9 (which renders the whole page and can assert on their output), and
  delete this bullet from the task rather than inventing a new test pattern for three small
  presentational files.

**Interfaces:**
- Consumes: `racePillsFor`/`factionColorVar` (Task 5), `loadBuildHref`/`simBuildHref` (Task
  5), `guidesCopy` (Task 6), `TreeGrid` in read-only mode (Task 4), `decodeFS1`/
  `orderFromRanks` (existing `lib/planner/fs1.ts`), `indexTalents` (existing
  `lib/planner/rules.ts`), `loadTalents` (existing `lib/planner/load.ts`),
  `createPlannerStore` (existing `lib/planner/store.svelte.ts`).
- Produces: `GuideBuildTree` props `{ code: string }`; renders one column per tree via
  `TreeGrid` in read-only mode, `data-testid="guide-tree-columns"` once ready,
  `data-testid="guide-tree-skeleton"` while loading, `data-testid="guide-tree-error"` (with a
  retry) on failure.
- Produces: `BuildActionButtons` props `{ code: string }`; two links,
  `data-testid="guide-load-build"` / `data-testid="guide-sim-build"`.
- Produces: `StatPriorityPills` props `{ stats: string[] }`; `data-testid="stat-priority-pills"`,
  one `data-testid="stat-pill-<index>"` per entry.
- Produces: `RacePillRow` props `{ classSlug: string; recommendedRaces: string[] }`;
  `data-testid="guide-race-pills"`, one `data-testid="race-pill-<slug>"` per legal race,
  `data-recommended="true"/"false"` on each.

- [ ] **Step 1: Write the failing test for `GuideBuildTree`**

```typescript
// web/src/components/guides/GuideBuildTree.test.ts
// Server-rendered only (no fetch happens during SSR -- the component's own $effect runs
// client-side after hydration), so this asserts the loading state's shape: the same
// `render(Component, {props})` -> `.body` pattern every other planner/guide component test
// in this codebase already uses (see ImportBox.test.ts, TalentCell.test.ts).
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import GuideBuildTree from './GuideBuildTree.svelte';

describe('GuideBuildTree', () => {
  it('server-renders the loading skeleton (no fetch happens before hydration)', () => {
    const { body } = render(GuideBuildTree, { props: { code: 'FS1:1.60.1.69893:warrior:human:1/2/3:' } });
    expect(body).toContain('data-testid="guide-tree-skeleton"');
  });
});
```

- [ ] **Step 2: Run it to see it fail**

Run: `cd web && npx vitest run src/components/guides/GuideBuildTree.test.ts`
Expected: FAIL — module does not exist.

- [ ] **Step 3: Implement `GuideBuildTree.svelte`**

```svelte
<!-- web/src/components/guides/GuideBuildTree.svelte -->
<!-- The read-only talent tree a spec guide embeds under "Talents and builds" (spec 5):
     decodes the guide's own `build:` FS1 code, fetches that build's talent data, reconstructs
     a legal point order with the planner's own orderFromRanks, and renders one column per
     tree through TreeGrid's own read-only mode -- never a second tree renderer. States model
     (2026-09-22 spec §1): a skeleton sized to the ready three-column grid, a retry on
     failure, one .reveal on success; nothing moves between them. -->
<script lang="ts">
  import { guidesCopy } from '../../lib/guides/copy';
  import { decodeFS1, orderFromRanks } from '../../lib/planner/fs1';
  import { loadTalents } from '../../lib/planner/load';
  import { indexTalents } from '../../lib/planner/rules';
  import { createPlannerStore, type PlannerStore } from '../../lib/planner/store.svelte';
  import LoadError from '../ui/LoadError.svelte';
  import Skeleton from '../ui/Skeleton.svelte';
  import TreeGrid from '../planner/TreeGrid.svelte';

  let { code }: { code: string } = $props();

  type Status = 'loading' | 'ready' | 'failed';
  let status = $state<Status>('loading');
  let store = $state<PlannerStore | null>(null);
  let attempt = $state(0);

  async function load(): Promise<void> {
    status = 'loading';
    const decoded = decodeFS1(code);
    if (!decoded.ok) {
      status = 'failed';
      return;
    }
    try {
      const talentFile = await loadTalents(decoded.build.dataBuild, decoded.build.classSlug);
      const index = indexTalents(talentFile);
      const { order } = orderFromRanks(index, decoded.build.treeRanks);
      const next = createPlannerStore({
        treeVersion: decoded.build.dataBuild,
        classSlug: decoded.build.classSlug,
        raceSlug: decoded.build.raceSlug,
        order,
        readOnly: true,
      });
      next.setTalents(talentFile);
      store = next;
      status = 'ready';
    } catch {
      status = 'failed';
    }
  }

  $effect(() => {
    void attempt;
    void load();
  });

  function retry(): void {
    attempt += 1;
  }
</script>

<!-- Always-present anchor so client:visible's observer has a real element from mount
     (HomeGuildLink.svelte's own pattern, for the same reason). -->
<span aria-hidden="true"></span>
{#if status === 'loading'}
  <Skeleton lines={3} rowHeight="h-12" minHeight="min-h-[360px]" testid="guide-tree-skeleton" />
{:else if status === 'failed'}
  <LoadError message={guidesCopy.treeLoadFailed} onRetry={retry} testid="guide-tree-error" />
{:else if status === 'ready' && store !== null && store.talentIndex !== null}
  {@const talentIndex = store.talentIndex}
  <div class="reveal grid grid-cols-1 gap-4 md:grid-cols-3" data-testid="guide-tree-columns">
    {#each talentIndex.trees as tree (tree.id)}
      <div class="flex flex-col gap-2">
        <h3 class="section-title text-[14px]">{tree.name}</h3>
        <TreeGrid {store} {tree} readOnly />
      </div>
    {/each}
  </div>
{/if}
```

- [ ] **Step 4: Run it to see it pass**

Run: `cd web && npx vitest run src/components/guides/GuideBuildTree.test.ts`
Expected: PASS.

- [ ] **Step 5: Implement `BuildActionButtons.astro`**

```astro
---
// web/src/components/guides/BuildActionButtons.astro
// The two links under a spec guide's embedded tree (spec 5): "Load this build" and "Sim this
// build". Plain links, not buttons with client handlers -- Planner.svelte's own
// writePlannerPointer and sim's sources.ts already write the current-character pointer
// unconditionally whenever their page loads with a ?code=, so there is nothing for a click
// handler here to do that the destination page does not already do on arrival.
import { guidesCopy } from '../../lib/guides/copy';
import { loadBuildHref, simBuildHref } from '../../lib/guides/build-links';

interface Props {
  code: string;
}
const { code } = Astro.props;
---

<div class="flex flex-wrap gap-3" data-testid="guide-build-actions">
  <a
    href={loadBuildHref(code)}
    class="rounded-control border-line-warm text-text inline-flex min-h-11 items-center border px-4 text-[14px] font-semibold"
    data-testid="guide-load-build"
  >
    {guidesCopy.loadThisBuild}
  </a>
  <a
    href={simBuildHref(code)}
    class="rounded-control border-line-warm text-text inline-flex min-h-11 items-center border px-4 text-[14px] font-semibold"
    data-testid="guide-sim-build"
  >
    {guidesCopy.simThisBuild}
  </a>
</div>
```

- [ ] **Step 6: Implement `StatPriorityPills.astro`**

```astro
---
// web/src/components/guides/StatPriorityPills.astro
// Spec 5: "Stat priority renders as an ordered row of pills with the reasoning beneath" --
// the guide's own existing prose (a numbered list or a sentence, depending on the guide) IS
// that reasoning and is left exactly as it is; this renders only the ordered pill row above
// it, from the guide's own statPriority frontmatter.
interface Props {
  stats: string[];
}
const { stats } = Astro.props;
---

<ol class="m-0 flex list-none flex-wrap gap-2 p-0" data-testid="stat-priority-pills">
  {
    stats.map((stat, index) => (
      <li
        class="rounded-pill bg-raised border-line inline-flex items-center gap-2 border px-3 py-1 text-[13px] font-semibold"
        data-testid={`stat-pill-${index}`}
      >
        <span class="tabular text-gold font-mono text-[11px]">{index + 1}</span>
        <span>{stat}</span>
      </li>
    ))
  }
</ol>
```

- [ ] **Step 7: Implement `RacePillRow.astro`**

```astro
---
// web/src/components/guides/RacePillRow.astro
// Spec 5 / the dispatch's own instruction: "race portraits do not exist: use the race name in
// a pill with the faction colour, no invented art." Every race legal for the class, the
// guide's own recommended ones marked both visually (a filled pill vs. an outline) and for
// screen readers (recommendedRaceLabel).
import { guidesCopy } from '../../lib/guides/copy';
import { factionColorVar, racePillsFor } from '../../lib/guides/race-pills';

interface Props {
  classSlug: string;
  recommendedRaces: string[];
}
const { classSlug, recommendedRaces } = Astro.props;
const pills = racePillsFor(classSlug, recommendedRaces);
---

<ul class="m-0 flex list-none flex-wrap gap-2 p-0" data-testid="guide-race-pills">
  {
    pills.map((race) => {
      const color = factionColorVar(race.faction);
      const style = race.recommended
        ? `background-color: color-mix(in srgb, ${color} 18%, transparent); border-color: ${color}; color: ${color}`
        : `border-color: var(--color-line); color: var(--color-muted)`;
      return (
        <li>
          <span
            class="rounded-pill inline-flex items-center gap-1 border px-3 py-1 text-[13px] font-semibold"
            style={style}
            data-testid={`race-pill-${race.slug}`}
            data-recommended={race.recommended}
          >
            {race.name}
            {race.recommended && <span class="sr-only">{` (${guidesCopy.recommendedRaceLabel})`}</span>}
          </span>
        </li>
      );
    })
  }
</ul>
```

- [ ] **Step 8: Scoped checks and commit**

```bash
cd web && npx astro check && npm run lint && npx prettier --check src/components/guides/
```

```bash
printf 'feat(web): the guide-tree embed, build-action buttons, and stat/race pill rows\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\nClaude-Session: https://claude.ai/code/session_01DAHb7A9Ytd6qpZu2UujXgY\n' > /tmp/commit-msg-7.txt
```

```bash
git add web/src/components/guides/
git commit -F /tmp/commit-msg-7.txt
```

---

## Task 8: `Content.astro` — the confidence caveat moves for guides only

**Files:**
- Modify: `web/src/layouts/Content.astro`
- Test: `web/src/layouts/Content.test.ts` (new, if no such test exists yet — check first with
  `find web/src/layouts -name "*.test.ts"`; if a test harness for `.astro` layouts already
  exists elsewhere in the project, follow that pattern instead of introducing a new one)

**Interfaces:**
- Produces: `Content.astro` gains an optional `confidencePlacement?: 'header' | 'footer'`
  prop, default `'header'` (today's exact behaviour, unchanged for `pages/[...slug].astro`,
  the only other consumer, which passes nothing). `'footer'` removes the confidence sentence
  from under the title and instead renders it as a quiet line directly above `SourcesList`.

- [ ] **Step 1: Check for an existing `.astro` layout test pattern**

Run: `find web/src -name "*.test.ts" | xargs grep -l "AstroContainer" 2>/dev/null | head -5`

If this prints file(s) (e.g. `_setup.test.ts`'s own pattern from `astro/container` +
`@astrojs/svelte/container-renderer`), write `Content.test.ts` using that exact harness. If
it prints nothing for a plain page (not the `[...slug]`/`[spec]` dynamic routes, which are
covered end-to-end in Task 9), skip a dedicated `Content.astro` unit test and instead assert
the two placements directly from the Task 9 page-level SSR test (one guide page for
`'footer'`, one existing `pages/[...slug].astro`-rendered page — e.g. `/about` — for
`'header'`, both already exercised elsewhere in the suite). Note the choice made in the
ledger.

- [ ] **Step 2: Implement the prop**

In `web/src/layouts/Content.astro`:

```typescript
interface Props {
  title: string;
  description?: string;
  path: string;
  updated: Date;
  confidence: Confidence;
  sources: Source[];
  /** 'header' (default, unchanged): the confidence sentence sits under the title, as every
   *  `pages` collection entry still shows it. 'footer' (guides only, spec 5): it moves to a
   *  quiet line directly above Sources instead -- "under the title" was the caveat's most
   *  prominent spot on the page and spec 5 asks for the least prominent one instead. */
  confidencePlacement?: 'header' | 'footer';
}
const {
  title,
  description = '',
  path,
  updated,
  confidence,
  sources,
  confidencePlacement = 'header',
} = Astro.props;
```

Change the header block to only show the confidence sentence when `confidencePlacement ===
'header'`:

```astro
      <div class="flex flex-wrap items-center gap-3 text-[13px] text-muted">
        <UpdatedStamp date={updated} />
        {confidencePlacement === 'header' && (
          <>
            <span>·</span>
            <span>{confidenceCopy[confidence]}</span>
          </>
        )}
      </div>
```

Add the footer line directly above `<SourcesList sources={sources} />`:

```astro
    {confidencePlacement === 'footer' && (
      <p class="text-muted text-[13px]" data-testid="confidence-footer-note">{confidenceCopy[confidence]}</p>
    )}
    <SourcesList sources={sources} />
```

- [ ] **Step 3: Verify `pages/[...slug].astro`'s own tests still pass unchanged**

Run: `cd web && npx vitest run src/pages/_classes.test.ts` (or whichever existing test
exercises a `pages`-collection page through `Content.astro` — grep first:
`grep -rl "layouts/Content\|Content.astro" web/src/pages/*.test.ts`). It calls `Content`
with no `confidencePlacement`, so its output must be byte-identical to before this task.
Expected: PASS, no diff.

- [ ] **Step 4: Scoped checks and commit**

```bash
cd web && npx astro check && npm run lint && npx prettier --check src/layouts/Content.astro
```

```bash
printf 'feat(web): Content.astro can move its confidence caveat to a footer line\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\nClaude-Session: https://claude.ai/code/session_01DAHb7A9Ytd6qpZu2UujXgY\n' > /tmp/commit-msg-8.txt
```

```bash
git add web/src/layouts/Content.astro
git commit -F /tmp/commit-msg-8.txt
```

**Ruling:** `Content.astro` is not listed under any lane in spec section 8, but it is the
shared layout `pages/[...slug].astro` (lane A's list, already merged, not touched again here)
and both guide pages (this lane's own `pages/guides/**`) render through. The new prop is
opt-in and defaults to today's exact behaviour, so lane A's already-merged pages render
identically. Cost if wrong: a merge conflict on this one file with whichever lane next
touches it; low risk given the additive, default-preserving shape of the change.

---

## Task 9: `[spec].astro` — assemble the page

**Files:**
- Modify: `web/src/pages/guides/[class]/[spec].astro`
- Modify: `web/src/pages/guides/[class]/index.astro` (confidence placement only)
- Test: `web/src/pages/guides/_spec.test.ts` (new)

**Interfaces:**
- Consumes: `splitSpecSections` (Task 1), `renderGuideMarkdown` (Task 2), the guide schema's
  `build`/`recommendedRaces`/`statPriority` (Task 3), `GuideBuildTree`/`BuildActionButtons`/
  `StatPriorityPills`/`RacePillRow` (Task 7), `Content.astro`'s `confidencePlacement` (Task 8).

- [ ] **Step 0: Confirm `entry.body` is readable off a `getCollection('guides')` entry**

Before writing the page, confirm the data source works as assumed:
`cd web && node -e "
import('./src/content.config.ts');
" 2>&1 | head -5` is not a real check (content collections need Astro's own loader) — instead
add a one-line temporary log inside `getStaticPaths` (`console.log(typeof guides[0]?.body)`),
run `npx astro build 2>&1 | head -20` once, confirm it prints `'string'`, then remove the log.
If `astro check` instead reports a type error on `entry.body` once Step 3 is written, apply
the fallback noted inline in Step 3's code rather than restructuring the page around a
different data source (the raw body is exactly what `_sections.test.ts` already validates
against, so it must stay the single source both read).

- [ ] **Step 1: Write the failing SSR test**

Uses the same `experimental_AstroContainer` + `@astrojs/svelte/container-renderer` pattern
`_setup.test.ts` already uses, so this needs no dev server. Renders the Fury Warrior guide
(a real spec guide, already in the repo) once Task 11 has filled in its frontmatter — write
this test now, expect it red until Task 11 lands, and re-check it green as Task 11's own
final step (same intermediate-red pattern as Task 3).

```typescript
// web/src/pages/guides/_spec.test.ts
// Underscore-prefixed for the same reason as every other _*.test.ts under src/pages/: Astro
// skips it, vitest still collects it.
import { experimental_AstroContainer as AstroContainer } from 'astro/container';
import { getContainerRenderer } from '@astrojs/svelte/container-renderer';
import { loadRenderers } from 'astro:container';
import { getCollection } from 'astro:content';
import { beforeAll, describe, expect, it } from 'vitest';
import SpecPage from './[class]/[spec].astro';

let container: AstroContainer;

beforeAll(async () => {
  const renderers = await loadRenderers([getContainerRenderer()]);
  container = await AstroContainer.create({ renderers });
});

describe('/guides/warrior/fury', () => {
  it('embeds the tree, the build-action buttons, the stat and race pills, and the footer caveat', async () => {
    const guides = await getCollection('guides');
    const entry = guides.find((g) => g.id === 'warrior/fury');
    if (!entry) throw new Error('warrior/fury guide not found');
    const html = await container.renderToString(SpecPage, { props: { entry } });
    expect(html).toContain('data-testid="guide-tree-columns"');
    expect(html).toContain('client:visible');
    expect(html).toContain('data-testid="guide-load-build"');
    expect(html).toContain('data-testid="guide-sim-build"');
    expect(html).toContain('data-testid="stat-priority-pills"');
    expect(html).toContain('data-testid="guide-race-pills"');
    expect(html).toContain('data-testid="confidence-footer-note"');
    // The nine headings still appear, in order, exactly as SpecTableOfContents promises.
    const talentsIndex = html.indexOf('id="talents-and-builds"');
    const rotationIndex = html.indexOf('id="rotation-and-priority"');
    expect(talentsIndex).toBeGreaterThan(-1);
    expect(rotationIndex).toBeGreaterThan(talentsIndex);
  });
});
```

- [ ] **Step 2: Run it to see it fail**

Run: `cd web && npx vitest run src/pages/guides/_spec.test.ts`
Expected: FAIL — the page still renders the whole body as one `<Body />`, none of the new
`data-testid`s exist yet.

- [ ] **Step 3: Rewrite `[spec].astro`**

```astro
---
// web/src/pages/guides/[class]/[spec].astro
// One spec guide: content/guides/<class>/<spec>.md, id `<class>/<spec>`. The nine-section
// structure (SPEC_SECTIONS) still renders in order with the same jump-list table of contents
// at the top; three of the nine sections now also carry a planted component (spec 5): the
// embedded read-only tree and its two buttons under "Talents and builds", an ordered pill
// row before the "Stat priority" prose, and a race pill row before the "Races" prose. No MDX
// is installed in this project, so a markdown file cannot embed a component directly --
// splitSpecSections/renderGuideMarkdown (lib/guides/) render each of the nine sections
// independently instead of the whole body as one `<Body />`, which is what makes planting
// possible while every heading keeps its exact id (renderGuideMarkdown uses the very
// processor Astro's own content pipeline does).
import { getCollection } from 'astro:content';
import BuildActionButtons from '../../../components/guides/BuildActionButtons.astro';
import GuideBuildTree from '../../../components/guides/GuideBuildTree.svelte';
import RacePillRow from '../../../components/guides/RacePillRow.astro';
import StatPriorityPills from '../../../components/guides/StatPriorityPills.astro';
import SpecTableOfContents from '../../../components/SpecTableOfContents.astro';
import Content from '../../../layouts/Content.astro';
import { renderGuideMarkdown } from '../../../lib/guides/markdown';
import { splitSpecSections, type SpecSection } from '../../../lib/guides/sections';
import { plannerHref } from '../../../lib/planner/reference';

export async function getStaticPaths() {
  const guides = await getCollection('guides');
  const specGuides = guides.filter((g) => g.data.spec !== undefined);
  return specGuides.map((entry) => {
    const [classSlug, spec] = entry.id.split('/');
    return { params: { class: classSlug, spec }, props: { entry } };
  });
}

const { entry } = Astro.props;
// entry.body is the collection entry's own raw markdown source (the same field
// _sections.test.ts reads off disk directly); if `astro check` reports it missing from the
// public CollectionEntry<'guides'> type in this Astro version, fall back to
// `(entry as unknown as { body?: string }).body ?? ''` with a one-line comment saying why,
// rather than switching to a different data source -- the raw body is what splitSpecSections
// is built to consume, and it must be byte-identical to what _sections.test.ts validates.
const sections = splitSpecSections(entry.body ?? '');

/** A section's raw markdown, minus its own `## Heading` line -- for the two sections that
 *  plant a component between the heading and the guide's own prose. */
function proseOnly(section: SpecSection): string {
  const raw = sections.get(section) ?? '';
  const [, ...rest] = raw.split('\n');
  return rest.join('\n').replace(/^\n+/, '');
}

const overviewHtml = await renderGuideMarkdown(sections.get('Overview') ?? '');
const talentsHtml = await renderGuideMarkdown(sections.get('Talents and builds') ?? '');
const rotationHtml = await renderGuideMarkdown(sections.get('Rotation and priority') ?? '');
const statPriorityHeadingHtml = await renderGuideMarkdown('## Stat priority');
const statPriorityProseHtml = await renderGuideMarkdown(proseOnly('Stat priority'));
const gearHtml = await renderGuideMarkdown(sections.get('Gear') ?? '');
const enchantsHtml = await renderGuideMarkdown(sections.get('Enchants and consumables') ?? '');
const racesHeadingHtml = await renderGuideMarkdown('## Races');
const racesProseHtml = await renderGuideMarkdown(proseOnly('Races'));
const professionsHtml = await renderGuideMarkdown(sections.get('Professions') ?? '');
const levelingHtml = await renderGuideMarkdown(sections.get('Leveling') ?? '');

const buildCode = entry.data.build;
---

<Content
  title={entry.data.title}
  description={entry.data.description}
  path={`/guides/${entry.id}`}
  updated={entry.data.updated}
  confidence={entry.data.confidence}
  sources={entry.data.sources}
  confidencePlacement="footer"
>
  <p>
    <a href={plannerHref(entry.data.classSlug)}>Open the planner for this class</a>
  </p>
  <SpecTableOfContents />
  <Fragment set:html={overviewHtml} />
  <Fragment set:html={talentsHtml} />
  {
    buildCode !== undefined && (
      <>
        <GuideBuildTree client:visible code={buildCode} />
        <BuildActionButtons code={buildCode} />
      </>
    )
  }
  <Fragment set:html={rotationHtml} />
  <Fragment set:html={statPriorityHeadingHtml} />
  <StatPriorityPills stats={entry.data.statPriority} />
  <Fragment set:html={statPriorityProseHtml} />
  <Fragment set:html={gearHtml} />
  <Fragment set:html={enchantsHtml} />
  <Fragment set:html={racesHeadingHtml} />
  <RacePillRow classSlug={entry.data.classSlug} recommendedRaces={entry.data.recommendedRaces} />
  <Fragment set:html={racesProseHtml} />
  <Fragment set:html={professionsHtml} />
  <Fragment set:html={levelingHtml} />
</Content>
```

- [ ] **Step 4: `[class]/index.astro`'s own caveat placement**

Class landing pages carry no talents/stats/races to plant components around, but spec 5's
"each guide's caveat moves" reads naturally as every guide page, not spec guides only. Add
the same prop to `web/src/pages/guides/[class]/index.astro`'s existing `<Content>` call:

```astro
<Content
  title={entry.data.title}
  description={entry.data.description}
  path={`/guides/${entry.data.classSlug}`}
  updated={entry.data.updated}
  confidence={entry.data.confidence}
  sources={entry.data.sources}
  confidencePlacement="footer"
>
```

(No other change to that file — its `<Body />` rendering stays exactly as today, since a
landing page has no per-section components to plant.)

- [ ] **Step 5: Run the SSR test again**

Run: `cd web && npx vitest run src/pages/guides/_spec.test.ts`
Expected: still FAIL until Task 11 fills in `warrior/fury.md`'s frontmatter (`buildCode` is
`undefined`, so the tree/buttons never render) — this is the same expected-red-until-Task-11
pattern as Task 3. Confirm the *reason* it fails is exactly that (no `guide-tree-columns`,
because `buildCode !== undefined` is false), not a crash — read the test output.

- [ ] **Step 6: Scoped checks (accepting the one known-red test) and commit**

```bash
cd web && npx astro check && npm run lint && npx prettier --check src/pages/guides/
```

```bash
printf 'feat(web): assemble the guide spec page from split sections and the new components\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\nClaude-Session: https://claude.ai/code/session_01DAHb7A9Ytd6qpZu2UujXgY\n' > /tmp/commit-msg-9.txt
```

```bash
git add web/src/pages/guides/
git commit -F /tmp/commit-msg-9.txt
```

---

## Task 10: `lighthouserc.json` — add the guide page to the strict bucket

**Files:**
- Modify: `web/lighthouserc.json`

**Interfaces:** none (config only).

- [ ] **Step 1: Add the URL**

In `web/lighthouserc.json`'s `ci.collect.url` array, add
`"http://localhost/guides/warrior/fury.html"` (after the existing `"http://localhost/guides.html"`
entry). Its pattern matches none of the `planner|logs|setup`, `reports`, or `sim*` bucket
regexes in `assertMatrix`, so it automatically falls into the first (strictest, 0.95
performance) bucket alongside `/guides.html` and `/index.html` — Global Constraint 7 / the
coordinator's "guide pages are in the strict bucket" needs no new bucket, only this one line.

- [ ] **Step 2: Confirm the JSON is still valid**

Run: `node -e "JSON.parse(require('fs').readFileSync('web/lighthouserc.json', 'utf8'))"`
Expected: no output (valid JSON, no exception).

- [ ] **Step 3: Commit**

```bash
printf 'chore(web): measure /guides/warrior/fury in the strict Lighthouse bucket\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\nClaude-Session: https://claude.ai/code/session_01DAHb7A9Ytd6qpZu2UujXgY\n' > /tmp/commit-msg-10.txt
```

```bash
git add web/lighthouserc.json
git commit -F /tmp/commit-msg-10.txt
```

(The actual `npm run lhci` run, and the `main`-comparison numbers, happen once in the final
verification step, after every page task has landed — running it once per task would waste
the numbered-runs budget on unfinished pages.)

**Ruling:** `web/lighthouserc.json` is not listed under any lane in spec section 8, but
Global Constraint 7 ("Lighthouse budgets... hold... before every final report") requires the
new page type to actually be measured, and no lane owns this file exclusively. Cost if wrong:
a merge conflict on one JSON array entry, trivial to resolve.

---

## Task 11: Fill in every spec guide's frontmatter

**Files:**
- Modify: all 27 files under `web/src/content/guides/*/[a-z]*.md` except each class's own
  `index.md` (the 9 landing pages are untouched — they have no `build:`/`recommendedRaces:`/
  `statPriority:`, matching the schema's optional/defaulted shape).

**Interfaces:** none new — this is data entry validated by Tasks 3 and 9's own tests.

- [ ] **Step 1: Add the three frontmatter fields to every spec guide**

For each of the 27 files, add `build:`, `recommendedRaces:`, and `statPriority:` to its YAML
frontmatter (anywhere after `role:` is a natural spot), using the exact values from this
plan's "Pre-computed data" section above. Quote the `build:` string (it contains `:`). Example
for `web/src/content/guides/warrior/fury.md` (its frontmatter today ends at `sources:` — add
these three keys before the closing `---`, after `role: dps`):

```yaml
role: dps
build: 'FS1:1.60.1.69893:warrior:human:35310003002/555000005050010051/0:'
recommendedRaces: [human, troll]
statPriority: [Attack power, Strength, Agility, Critical strike, Hit, Melee haste]
```

Repeat for the other 26 files using the plan's own tables (FS1 codes table,
`recommendedRaces` table, `statPriority` table). Do not touch anything else in any of these
27 files' frontmatter or body in this step.

- [ ] **Step 2: Adjust the six prose sentences that now describe a slightly different number**

The real, legally-reachable point totals (from actually building each tree, tier gates and
prerequisites included) differ from a few guides' own loose "roughly N points" wording. Fix
only the sentence that states the number, in exactly these six files (every other guide's
existing wording already matches its own code's total closely enough that "roughly" already
covers the gap):

  - `warrior/fury.md`: "A build reaching Bloodthirst spends roughly 31 points in Fury" →
    "roughly 32 points in Fury" (Bloodthirst's own tier-6 gate needs 30, plus the point in
    Bloodthirst itself and Death Wish's own prerequisite chain lands at 32).
  - `rogue/assassination.md`: "roughly 31 points are needed to reach Venom at the bottom of
    the tree" → "roughly 33 points".
  - `mage/arcane.md`: "remaining 20 points usually going into Fire for Ignite and Critical
    Mass" → "remaining 10 points usually going into Fire for Ignite" (Critical Mass's own
    tier-4 gate made both together too expensive to fit this build's own 51-point cap
    alongside Arcane Power; Ignite alone still shows Fire's own school-power idea).
  - `mage/fire.md`: "remaining 20 points usually going into Arcane for Arcane Concentration's
    Clearcasting and Arcane Mind's crit damage bonus" → "remaining 10 points usually going
    into Arcane for Arcane Concentration's Clearcasting" (Arcane Mind's own tier-4 gate made
    the same trade-off as above).
  - `mage/frost.md`: "remaining 20 points usually going into Fire for Ignite and Critical
    Mass" → "remaining 10 points usually going into Fire for Ignite".
  - `rogue/subtlety.md`: "a handful of points in Combat for utility (Deflection, Dual Wield
    Specialization)" → "a handful of points in Combat for utility (Deflection)" (Dual Wield
    Specialization's own tier-3 gate and its own prerequisite made keeping both too
    expensive alongside a full Subtlety tree).
  - `warrior/protection.md`: "a couple of Arms points for Deep Wounds" → "a couple of early
    Arms points" (Deep Wounds's own tier-2 gate made it too expensive to fit alongside a full
    Fury and Protection spend in this build specifically; `warrior/fury.md` and
    `warrior/arms.md` still spend on Deep Wounds directly, where it fits).
  - `druid/restoration.md`: "remaining 16 usually going into Balance for Nature's Grace's
    faster recast on a crit and Moonglow's mana discount" → "remaining 8 usually going into
    Balance for Moonglow's mana discount" (Nature's Grace's own tier-4 gate made both too
    expensive alongside a full Restoration spend).

- [ ] **Step 2: Run the two SSR tests written earlier**

Run: `cd web && npx vitest run src/content/guides/_sections.test.ts src/pages/guides/_spec.test.ts`
Expected: PASS — every `it.each` in `_sections.test.ts` (schema validity, section order, and
the new build-code decode/legality check) green for all 27 spec guides; `_spec.test.ts`'s
`warrior/fury` page-level assertions green now that `buildCode` is defined.

- [ ] **Step 3: Run the whole guides+planner scoped suite**

Run: `cd web && npx vitest run src/content/guides/ src/pages/guides/ src/lib/guides/ src/components/guides/ src/components/planner/ src/lib/planner/`
Expected: PASS, no regressions.

- [ ] **Step 4: `astro check`, lint, prettier, and a real build**

```bash
cd web && npx astro check && npm run lint && npx prettier --check src/content/guides/
FOREVER_DATA=fixture npm run sync && npm run build
```

Expected: the build succeeds and emits `dist/guides/warrior/fury.html` (and all 26 other
spec guide pages) with no errors. (`npm run build` here is the standard way this repo
verifies a dynamic-route page tree actually generates for every static path — check
`package.json`'s own `build` script if this differs; the goal is confirming all 27 pages
build clean, not necessarily this exact command.)

- [ ] **Step 5: Commit**

```bash
printf 'feat(web): every spec guide ends in a real, legal, decodable build\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\nClaude-Session: https://claude.ai/code/session_01DAHb7A9Ytd6qpZu2UujXgY\n' > /tmp/commit-msg-11.txt
```

```bash
git add web/src/content/guides/
git commit -F /tmp/commit-msg-11.txt
```

---

## Task 12: E2E — "Load this build" lands in the planner with the points lit

**Files:**
- Create: `web/tests/e2e/guides.spec.ts`

**Interfaces:** none new — exercises the shipped page and the existing planner.

- [ ] **Step 1: Check how an existing planner e2e spec asserts "points lit" and the fixture data build id**

Run: `grep -rl "data-testid=\"talent-" web/tests/e2e/*.spec.ts | head -3` and read one hit for
the exact assertion shape (a filled-state `data-state="filled"`/`"maxed"` attribute, per
`styles.ts`'s `CELL_BORDER`/`cellState`). Also check `grep -n "FOREVER_DATA\|fixture" web/tests/e2e/support/*.ts`
for how e2e's own server is started (fixture data vs. real data) — the guide pages' `build:`
codes are real-build (`1.60.1.69893`) FS1 codes, so this spec needs the e2e server serving
that real build's talent files under `/data/1.60.1.69893/talents/warrior.json`, not the
fixture set. Follow whatever existing spec (if any) already depends on real, non-fixture
planner data for its own `?code=` test; if none does, this spec is the first, and Step 2's
`beforeAll` must run `npm run sync` (no `FOREVER_DATA=fixture`) before starting the e2e
preview server, then restore the fixture sync afterward so it doesn't leak into other specs'
run — ledger this as a ruling either way, since it changes what `beforeAll`/`afterAll` do for
the whole file.

- [ ] **Step 2: Write the spec**

```typescript
// web/tests/e2e/guides.spec.ts
// Spec section 9 (lane C): "a spec that Load this build lands in the planner with the
// points lit."
import { test, expect } from '@playwright/test';

test('Load this build on the Fury Warrior guide lands in the planner with points already spent', async ({
  page,
}) => {
  await page.goto('/guides/warrior/fury');
  const loadLink = page.getByTestId('guide-load-build');
  await expect(loadLink).toBeVisible();
  const href = await loadLink.getAttribute('href');
  expect(href).toMatch(/^\/planner\?code=/);

  await loadLink.click();
  await expect(page).toHaveURL(/\/planner\?code=/);

  // Bloodthirst (Fury's own named capstone) is lit in the loaded build.
  const bloodthirstCell = page.locator('[data-testid^="talent-"][aria-label^="Bloodthirst"]');
  await expect(bloodthirstCell).toHaveAttribute('data-rank', '1');
});
```

- [ ] **Step 3: Run it**

Run: `E2E_PORT=4387 npx playwright test tests/e2e/guides.spec.ts`
Expected: PASS. If the real-build data isn't being served (Step 1's finding), fix the
`beforeAll`/global setup per what Step 1 found, not the test itself.

- [ ] **Step 4: Scoped checks and commit**

```bash
cd web && npx prettier --check tests/e2e/guides.spec.ts
```

```bash
printf 'test(web): e2e -- Load this build on a guide lands in the planner with points lit\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\nClaude-Session: https://claude.ai/code/session_01DAHb7A9Ytd6qpZu2UujXgY\n' > /tmp/commit-msg-12.txt
```

```bash
git add web/tests/e2e/guides.spec.ts
git commit -F /tmp/commit-msg-12.txt
```

---

## Task 13: Whole-branch review, Lighthouse, whole e2e suite, final report

Not a code task — the closing verification pass `lane-common-web.md` requires before the
final report. Steps:

- [ ] Read every file this plan touched or added, end to end, checking: no `console.log`, no
  hardcoded copy strings outside `lib/guides/copy.ts`, no file over 800 lines, no function
  over 50 lines, immutable updates throughout, every new test asserts something real (no
  placeholder `expect(true).toBe(true)`).
- [ ] Full scoped check across everything this lane touched:
  `cd web && npx vitest run src/lib/guides/ src/components/guides/ src/content/guides/ src/pages/guides/ src/components/planner/TreeGrid.test.ts src/components/planner/TalentCell.test.ts src/layouts/ && npx astro check && npm run lint && npx prettier --check .`
- [ ] `FOREVER_DATA=fixture npm run sync`, then `npm run build`, then
  `npx astro preview --port 4387 &`, then `npm run lhci` — quote the guide page's own
  performance/LCP/TBT/CLS numbers.
- [ ] Compare against `main`: `git worktree list` to confirm no stray worktree is needed;
  build and preview `main`'s own `/guides/warrior/fury` (a plain content page there today,
  no tree/pills) under the same `lhci` config and quote its numbers alongside this branch's,
  per Global Constraint 7. Stop the preview server afterward
  (`npx astro preview stop`), and confirm with
  `ps -axo command | grep "[a]stro.*guides-build"` that nothing is left running.
- [ ] Run the **whole** e2e suite once, not just `guides.spec.ts`:
  `E2E_PORT=4387 npx playwright test`. A failure reproducible on `main` is reported, not
  fixed; anything this lane's changes caused is fixed here.
- [ ] Stop every server this lane started; confirm with the same `ps`/`pkill`-scoped-to-
  worktree-path check `lane-common-web.md` requires.
- [ ] Write the final report (under 25 lines): branch head sha; what shipped per spec
  bullet; the SSR decode test's result over all 27 guides; test counts (vitest, astro check,
  lint, prettier, e2e, lhci with `main`'s numbers alongside); every ruling from this plan
  (Task 4's `TalentCell.svelte`, Task 8's `Content.astro`, Task 10's `lighthouserc.json`,
  and Task 12's fixture-vs-real-data e2e setup decision); every copy string added (the whole
  of `lib/guides/copy.ts`); files touched that lane B or D might also touch (`Content.astro`
  is the one to flag explicitly); confirmation every server is stopped.
