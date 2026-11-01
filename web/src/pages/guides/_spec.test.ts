// web/src/pages/guides/_spec.test.ts
// Underscore-prefixed for the same reason as every other _*.test.ts under src/pages/: Astro
// skips it, vitest still collects it.
//
// Builds its own `entry` by reading the guide file directly off disk and parsing it through
// the real `guideSchema`, the same way _sections.test.ts's own `loadGuides()` does -- NOT via
// `getCollection('guides')`, which returns an empty collection under plain `vitest run`
// unless Astro's dev-mode content-layer cache (`.astro/data-store.json`) happens to already be
// warm from a prior `astro dev` run, something nothing in this repo's CI produces before
// `npm test`. Running the real schema over the real file keeps this test honest about what
// the page actually receives (a real `CollectionEntry`-shaped object with `id`/`data`/`body`)
// without depending on Astro's content layer being primed.
import { readFileSync } from 'node:fs';
import { fileURLToPath } from 'node:url';
import type { experimental_AstroContainer as AstroContainer } from 'astro/container';
import { parse as parseYaml } from 'yaml';
import { beforeAll, describe, expect, it } from 'vitest';
import { guideSchema } from '../../content.config';
import SpecPage from './[class]/[spec].astro';
import { raceRows } from '../../lib/planner/reference';
import { racialEffects } from '../../lib/guides/race-table';
import { decodeBuildTrees } from '../../lib/guides/build-tree';
import { schoolHitTalent } from '../../lib/guides/hit-cap-line';
import { createPageContainer } from '../../test-support/page-container';
import activeBuild from '../../data/active-build.json';

let container: AstroContainer;

beforeAll(async () => {
  container = await createPageContainer();
});

function loadGuideEntry(id: string) {
  const path = fileURLToPath(new URL(`../../content/guides/${id}.md`, import.meta.url));
  const raw = readFileSync(path, 'utf8');
  const [, frontmatterBlock, ...bodyParts] = raw.split('---');
  const data = guideSchema.parse(parseYaml(frontmatterBlock));
  return { id, collection: 'guides' as const, data, body: bodyParts.join('---').trim() };
}

/**
 * The talent a guide's hit line names, read from the active build's trees the way the page
 * reads it (the raid build when it differs, else the leveling build; `schoolHitTalent`; the
 * copy's own "talents" when none), so a client build that renames the talent moves the
 * expectation with it instead of failing the move (Elemental Precision became Precision in
 * 1.60.1.70685).
 */
function guideHitTalent(id: string): string {
  const { build, raidBuild } = loadGuideEntry(id).data;
  const code = raidBuild !== undefined && raidBuild !== build ? raidBuild : build;
  const trees = code === undefined ? [] : (decodeBuildTrees(activeBuild.build, code) ?? []);
  return schoolHitTalent(trees) ?? 'talents';
}

/**
 * Client builds where the frost guide's hit line deliberately says "from talents": the
 * build's school-hit talent is not one `SCHOOL_HIT_TALENTS` names, and naming it was
 * decided against. Empty until someone decides that; a move to a build that renames the
 * talent (70685 and 70707 call it Precision) fails here, so the copy is a decision, never
 * a silent fallback.
 */
const HIT_LINE_FALLBACK_BUILDS: readonly string[] = [];

function escapeRegExp(text: string): string {
  return text.replace(/[.*+?^${}()|[\]\\]/g, '\\$&');
}

describe('/guides/warrior/fury', () => {
  it('embeds the tree, the build-action buttons, the stat and race pills, and the footer caveat', async () => {
    const entry = loadGuideEntry('warrior/fury');
    const html = await container.renderToString(SpecPage, { props: { entry } });
    // GuideBuildTree hydrates client:visible and fetches its talent data in a $effect that
    // only runs after hydration -- SSR always renders its loading skeleton
    // (GuideBuildTree.test.ts's own assertion), never the post-fetch tree. So this checks
    // what SSR actually produces: the island mounted with the right initial state and the
    // right code prop, and Astro's own client-directive attribute (not the literal string
    // "client:visible", which never appears -- Astro serializes it as `client="visible"` on
    // the <astro-island> element).
    expect(html).toContain('data-testid="guide-tree-skeleton"');
    expect(html).toContain('client="visible"');
    expect(html).toContain(entry.data.build as string);
    expect(html).toContain('data-testid="guide-load-build"');
    expect(html).toContain('data-testid="guide-sim-build"');
    expect(html).toContain('data-testid="stat-priority-pills"');
    expect(html).toContain('data-testid="guide-race-pills"');
    expect(html).not.toContain('confidence-footer-note');
    expect(html).not.toContain('Our inference');
    // The nine headings still appear, in order, exactly as SpecTableOfContents promises.
    const talentsIndex = html.indexOf('id="talents-and-builds"');
    const rotationIndex = html.indexOf('id="rotation-and-priority"');
    expect(talentsIndex).toBeGreaterThan(-1);
    expect(rotationIndex).toBeGreaterThan(talentsIndex);
  });
});

describe('/guides/rogue/assassination (a spec with a raid build)', () => {
  it('shows the leveling build and the raid build as two labelled trees with their own links', async () => {
    const entry = loadGuideEntry('rogue/assassination');
    const raidBuild = entry.data.raidBuild as string;
    expect(raidBuild).not.toBe(entry.data.build);
    const html = await container.renderToString(SpecPage, { props: { entry } });
    expect(html).toContain('data-testid="guide-build-leveling"');
    expect(html).toContain('data-testid="guide-build-raid"');
    expect(html).toContain('>Level 60 build</h3>');
    expect(html).toContain('>Raid build</h3>');
    expect(html).toContain(raidBuild);
    expect(html).toContain('data-testid="guide-load-raid-build"');
    expect(html).toContain('data-testid="guide-sim-raid-build"');
    expect(html).toContain(`href="/planner?code=${encodeURIComponent(raidBuild)}"`);
  });
});

describe('/guides/warrior/fury (a spec whose raid build is its leveling build)', () => {
  it('shows one tree under a single "Level 60 and raid build" heading', async () => {
    const entry = loadGuideEntry('warrior/fury');
    expect(entry.data.raidBuild).toBeUndefined();
    const html = await container.renderToString(SpecPage, { props: { entry } });
    expect(html).toContain('data-testid="guide-build-both"');
    expect(html).toContain('>Level 60 and raid build</h3>');
    expect(html).not.toContain('data-testid="guide-build-raid"');
  });
});

describe('/guides/mage/frost (the class-guide blocks)', () => {
  async function frostHtml(): Promise<string> {
    return container.renderToString(SpecPage, { props: { entry: loadGuideEntry('mage/frost') } });
  }

  it('says the status once, at the top, and carries the header facts', async () => {
    const html = await frostHtml();
    expect(html.match(/data-testid="guide-status-note"/g)).toHaveLength(1);
    expect(html).toContain('The beta stops at level 30.');
    expect(html).toMatch(/DPS · Updated \d{1,2} [A-Z][a-z]{2}</);
    // The header summary is the Overview's own sentence, not the frontmatter description.
    expect(html).toContain('Frost is the Mage control and survivability spec');
  });

  it('lays the Overview out as strengths, weaknesses and the raid role', async () => {
    const html = await frostHtml();
    expect(html).toContain('data-testid="guide-overview-strengths"');
    expect(html).toContain('data-testid="guide-overview-weaknesses"');
    expect(html).toContain('data-testid="guide-overview-raid"');
  });

  it('names the raid build’s difference in one line and lists the key talents with icons', async () => {
    const html = await frostHtml();
    expect(html).toMatch(/data-testid="guide-build-diff"[^>]*>The raid build moves \d+ points? from /);
    expect(html).toContain('data-testid="guide-key-talents"');
    expect(html).toContain('Improved Frostbolt');
    expect(html).toMatch(/guide-key-talent-0[\s\S]*?<img src="\/data\/[^"]+\/icons\/[^"]+\.webp"/);
  });

  it('numbers a priority with spell icons, names the real cooldowns, and prints the pre-pull Ice Barrier as the opener', async () => {
    const html = await frostHtml();
    expect(html).toContain('data-testid="guide-rotation-priority"');
    expect(html).toMatch(/guide-rotation-priority-0[\s\S]*?<img[^>]+icons\/spell_frost_frostblast\.webp/);
    expect(html).toContain('Ice Lance');
    expect(html).toMatch(/data-testid="guide-rotation-opener-0"(?:(?!<\/li>)[\s\S])*?Ice Barrier/);
    expect(html).not.toContain('Then work down it from the top');
    for (const name of ['Evocation', 'Cold Snap', 'Ice Block']) {
      expect(html, name).toMatch(new RegExp(`data-testid="guide-rotation-cooldowns-\\d+"[\\s\\S]*?${name}`));
    }
    expect(html).toContain('8 min cooldown');
    // Evocation has its icon from the spell icon data, so every spell row shares the left edge.
    expect(html).toMatch(/guide-rotation-cooldowns-\d+"[^>]*>(?:(?!<\/li>)[\s\S])*?spell_nature_purge/);
  });

  it('gives the stat block pills, a weights table and the hit line as its headline, in one order', async () => {
    const html = await frostHtml();
    expect(html).toContain('data-testid="stat-priority-pills"');
    expect(html).toContain('data-testid="guide-stat-table"');
    const named = guideHitTalent('mage/frost');
    if (!HIT_LINE_FALLBACK_BUILDS.includes(activeBuild.build)) {
      expect(
        named,
        `the frost hit line names no talent on ${activeBuild.build}: add its school-hit talent to SCHOOL_HIT_TALENTS, or the build to HIT_LINE_FALLBACK_BUILDS`,
      ).not.toBe('talents');
    }
    // Precision adds hit to every Mage spell, so whether the band's cap still carries a
    // talent split is the nightly ranking's call. When it does, the split names the build's
    // talent, not the "talents" fallback.
    const talent = escapeRegExp(named);
    expect(html).toMatch(
      new RegExp(
        `data-testid="guide-stat-table-hit-cap"[^>]*>Hit cap: \\d+(\\.\\d)?%(?: · \\d+(\\.\\d)?% from ${talent}, \\d+(\\.\\d)?% from gear)?<`,
      ),
    );
    expect(html).toContain('DPS per point');
    // The unit line names the top per-point stat at 60, and the nightly sim decides which one
    // that is: for Frost, hit and spell power sit within error of each other and swap places
    // between nightlies. Pin the shape, a lowercase stat label and then dpsText's figure (two
    // decimals under 10, one at 10 or more), never which stat wins.
    expect(html).toMatch(
      /data-testid="guide-stat-table-unit"[^>]*>1 point of [a-z][a-z ]* = (?:\d\.\d\d|\d{2,}\.\d) DPS at level 60</,
    );
    expect(html).toMatch(/\d+\.\d per 1% haste/);
    expect(html).not.toMatch(/\d per 1%</);
    expect(html).toContain('Spell penetration shows as &quot;Not significant&quot;');
    expect(html).not.toContain('from school');
    const pills = [...html.matchAll(/data-testid="stat-pill-(\d+)"[\s\S]*?<span>([^<]+)<\/span>/g)].map(
      (m) => m[2],
    );
    const rows = [
      ...html.matchAll(/data-testid="guide-stat-table-row-\d+"[^>]*><td class="stat-name">([^<]+)</g),
    ].map((m) => m[1]);
    expect(pills.length).toBeGreaterThan(3);
    expect(pills).toEqual(rows);
  });

  it('puts the best-in-slot table in the page, one panel per faction and level range, with the BiS link', async () => {
    const html = await frostHtml();
    for (const faction of ['alliance', 'horde']) {
      for (const band of [20, 30, 40, 50, 60]) {
        expect(html).toContain(`data-testid="guide-bis-panel-${faction}-${band}"`);
      }
      expect(html).toContain(`data-testid="guide-bis-faction-${faction}"`);
    }
    for (const band of [20, 30, 40, 50, 60]) {
      expect(html).toContain(`data-testid="guide-bis-pill-${band}"`);
    }
    expect(html).toMatch(/id="gbis-band-60"[^>]*checked/);
    expect(html).toMatch(/id="gbis-f-alliance"[^>]*checked/);
    expect(html).toContain('href="/bis/mage/frost?faction=alliance#band-alliance-60"');
    expect(html).toContain('href="/bis/mage/frost?faction=horde#band-horde-60"');
    expect(html).toContain('Full list and alternatives');
    expect(html).toContain('data-testid="guide-bis-row-alliance-60-head"');
    expect(html).not.toContain('named honestly');
    expect(html).not.toMatch(/<astro-island[^>]*GuideBisTable/);
  });

  it('adds the enchant and consumable tables, the race table and the leveling talents', async () => {
    const html = await frostHtml();
    expect(html).toContain('data-testid="guide-enchants"');
    expect(html).toContain('data-testid="guide-consumables"');
    expect(html).toContain('Elixir of Frost Power');
    expect(html).toContain('replaces both elixirs');
    expect(html).toContain('Without a flask');
    expect(html).not.toContain('The caster meal');
    expect(html).toContain('data-testid="guide-race-table"');
    expect(html).toContain('data-testid="guide-race-row-gnome"');
    expect(html).toMatch(/data-testid="guide-race-row-gnome" data-recommended="true"/);
    // The racials and the recommendation line come from the race data of the sync that ran: the
    // real build's sentences, or none under the unit-test fixture.
    const gnome = raceRows.find((race) => race.slug === 'gnome');
    if (gnome !== undefined && racialEffects(gnome).has('Escape Artist')) {
      expect(html).toContain('Gnome is the Alliance pick: Escape Artist and Expansive Mind.');
      expect(html).not.toContain('Eureka!');
    }
    expect(html).toContain('data-testid="guide-leveling-row-talents"');
    expect(html).toMatch(/leveling-chip-rank[^>]*>\d\/\d</);
    expect(html).toContain('Level 20 to 29');
    expect(html).toContain('11 to 20 points');
    expect(html).toContain('Outlined: points added in this level range.');
  });

  it('renders no per-section hedge', async () => {
    const html = await frostHtml();
    for (const hedge of ['is a projection', 'not tested', 'cannot be named', 'inferred from 1.12']) {
      expect(html, hedge).not.toContain(hedge);
    }
  });
});
