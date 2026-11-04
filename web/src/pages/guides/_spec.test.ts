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
import { bandEntry, readSpecCatalog } from '../../lib/bis/load';
import { guideLadderSteps } from '../../lib/guides/leveling-builds';
import { rankedBisFile } from '../../lib/guides/class-dps';
import { guideHitCapLine } from '../../lib/guides/hit-cap-line';
import { bandTalentHits } from '../../lib/guides/talent-hit';
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

  it('keeps sim talk out of the reading flow: one data-source line in the footer, no site pill', async () => {
    const html = await frostHtml();
    expect(html).not.toContain('data-testid="guide-status-note"');
    // The guide's own reading flow (the site nav's "Simulator" link is not the guide speaking).
    const main = html.slice(html.indexOf('<main'), html.lastIndexOf('</main>'));
    const visible = main.replace(/<script[\s\S]*?<\/script>|<style[\s\S]*?<\/style>|<[^>]+>/g, ' ');
    const talk = [...visible.matchAll(/.{0,60}(?:simulator|simulated|this site).{0,40}/gi)];
    expect(talk.map((m) => m[0])).toEqual([]);
    expect(html).toMatch(
      /data-testid="content-footnote"[^>]*>\s*Numbers from client build 1\.60\.1\.\d+, checked against launch logs/,
    );
    expect(main).not.toContain('This site');
    expect(html).toMatch(/DPS · Updated \d{1,2} [A-Z][a-z]{2}</);
    // The header summary is the Overview's own sentence, not the frontmatter description.
    expect(html).toContain('Frost casts Frostbolt behind Ice Barrier');
  });

  it('lays the Overview out as strengths, weaknesses and the raid role', async () => {
    const html = await frostHtml();
    expect(html).toContain('data-testid="guide-overview-strengths"');
    expect(html).toContain('data-testid="guide-overview-weaknesses"');
    expect(html).toContain('data-testid="guide-overview-raid"');
  });

  it('names the raid build’s difference in one line and lists the key talents with icons', async () => {
    const html = await frostHtml();
    // Frost plays one build at 60 in and out of raids; Arcane's raid build differs.
    const arcane = await container.renderToString(SpecPage, {
      props: { entry: loadGuideEntry('mage/arcane') },
    });
    expect(arcane).toMatch(/data-testid="guide-build-diff"[^>]*>The raid build moves \d+ points? from /);
    expect(html).toContain('data-testid="guide-build-both"');
    expect(html).toContain('data-testid="guide-key-talents"');
    expect(html).toContain('Improved Frostbolt');
    expect(html).toMatch(/guide-key-talent-0[\s\S]*?<img src="\/data\/[^"]+\/icons\/[^"]+\.webp"/);
    // Two full guide renders in one test.
  }, 20_000);

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
    // The trinket line names the racial cooldowns the sim casts for the mage's races.
    expect(html).toMatch(
      /data-testid="guide-rotation-cooldowns-\d+"(?:(?!<\/li>)[\s\S])*?Trinkets, racials and potions(?:(?!<\/li>)[\s\S])*?Eureka! \(Gnome\)/,
    );
  });

  it('gives the stat block pills, a weights table and the hit line as its headline, in one order', async () => {
    const html = await frostHtml();
    expect(html).toContain('data-testid="stat-priority-pills"');
    expect(html).toContain('data-testid="guide-stat-table"');
    // The line names every source the band's figure counts: the BiS gear and buffs, and each
    // talent of the weights character's own build that adds hit (Precision, for Frost).
    const file = rankedBisFile(activeBuild.build, 'mage-frost');
    const band = file === undefined ? undefined : bandEntry(file, 60, 'alliance');
    const line = guideHitCapLine(
      band?.hit_to_cap,
      band === undefined ? [] : bandTalentHits(activeBuild.build, 'mage', band.talents),
    );
    expect(line).toBeDefined();
    expect(html).toContain(`data-testid="guide-stat-table-hit-cap">${line!.text}<`);
    expect(line!.text).toContain('from Precision');
    expect(html).toContain('DPS per point');
    // The unit line names the top per-point stat at 60, and the nightly sim decides which one
    // that is: for Frost, hit and spell power sit within error of each other and swap places
    // between nightlies. Pin the shape, a lowercase stat label and then dpsText's figure (two
    // decimals under 10, one at 10 or more), never which stat wins.
    expect(html).toMatch(
      /data-testid="guide-stat-table-unit"[^>]*>1 (?:point of [a-z][a-z ]*|[a-z][a-z ]* rating) = (?:\d\.\d\d|\d{2,}\.\d) DPS at level 60\.?</,
    );
    // Hit and crit are priced per rating point; the note under the table says what a point is.
    expect(html).toMatch(
      /data-testid="guide-stat-table-ratings"[^>]*>[^<]*per rating point: [^<]*10 hit rating is 1% hit/,
    );
    // No item at 60 carries haste: its row reads "Not on any item" and the note under the
    // table gives its per-1% worth.
    expect(html).toMatch(/Spell haste: \d+\.\d per 1%, but no item at level 60 has it\./);
    expect(html).not.toMatch(/\d per 1%</);
    expect(html).toContain('Spell penetration shows as &quot;Not significant&quot;');
    expect(html).not.toContain('from school');
    const pills = [...html.matchAll(/data-testid="stat-pill-(\d+)"[\s\S]*?<span>([^<]+)<\/span>/g)].map(
      (m) => m[2],
    );
    // The pills rank exactly the rows the table gives a number, in the table's order.
    const rankedRows = [
      ...html.matchAll(
        /data-testid="guide-stat-table-row-\d+"[^>]*><td class="stat-name">([^<]+)<\/td>[\s\S]*?<\/tr>/g,
      ),
    ]
      .filter((m) => !m[0].includes('stat-muted'))
      .map((m) => m[1]);
    expect(pills.length).toBeGreaterThan(3);
    expect(pills).toEqual(rankedRows);
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
    // A Classic flask stacks with the elixirs: listed beside them, never instead of them.
    expect(html).toContain('Flask of Supreme Power');
    expect(html).not.toContain('replaces both elixirs');
    expect(html).not.toContain('Without a flask');
    expect(html).not.toContain('The caster meal');
    expect(html).toContain('data-testid="guide-race-table"');
    expect(html).toContain('data-testid="guide-race-row-gnome"');
    expect(html).toMatch(/data-testid="guide-race-row-gnome" data-recommended="true"/);
    // The guide's own Races prose names both picks, so no pick line repeats it under the table.
    expect(html).not.toContain('data-testid="guide-race-picks"');
    // The racials come from the race data of the sync that ran: the real build's sentences, or
    // none under the unit-test fixture.
    const gnome = raceRows.find((race) => race.slug === 'gnome');
    if (gnome !== undefined && racialEffects(gnome).has('Eureka!')) {
      expect(html).toContain(
        'Makes your damaging abilities cost 10% less Mana and deal 10% more damage for 8 seconds',
      );
      expect(html).not.toContain('next 3 abilities');
      // The prose names Eureka! as the client has it (8 s, 10% damage, 10% cheaper), never the
      // BlizzCon demo's reading.
      expect(html).not.toMatch(/Eureka![^.<]*50%/);
    }
    // The leveling tree, server-rendered at level 60: the caption, the control and the link.
    expect(html).toContain('data-testid="guide-leveling-tree"');
    expect(html).toMatch(/data-testid="guide-leveling-caption"[\s\S]*?Level 60: [^<]+, rank \d/);
    expect(html).toContain('51 points spent, all placed');
    expect(html).toMatch(
      /href="\/planner\?spec=mage-frost&amp;talents=[0-9-]+"[^>]*data-testid="guide-leveling-load"|data-testid="guide-leveling-load"[^>]*href="\/planner\?spec=mage-frost&amp;talents=/,
    );
    expect(html).not.toContain('guide-leveling-strip');
  });

  it('opens #leveling on the Levels 1 to 19 text, then the tree under Talents by level, then Levels 20 to 60', async () => {
    const html = await frostHtml();
    const at = (needle: string): number => html.indexOf(needle);
    expect(at('id="leveling"')).toBeGreaterThan(-1);
    expect(at('id="levels-1-to-19"')).toBeGreaterThan(at('id="leveling"'));
    expect(at('id="talents-by-level"')).toBeGreaterThan(at('id="levels-1-to-19"'));
    expect(at('data-testid="guide-leveling-tree"')).toBeGreaterThan(at('id="talents-by-level"'));
    expect(at('id="levels-20-to-60"')).toBeGreaterThan(at('data-testid="guide-leveling-tree"'));
  });

  it('renders no per-section hedge', async () => {
    const html = await frostHtml();
    for (const hedge of ['is a projection', 'not tested', 'cannot be named', 'inferred from 1.12']) {
      expect(html, hedge).not.toContain(hedge);
    }
  });
});

describe('/guides/priest/shadow leveling tree', () => {
  // Sim-accuracy round 16: the ladder opens on Wand Specialization at 10 and 11 and Spirit Tap
  // from 12, which the addon spends too; the tree's steps are that order.
  it('opens on Wand Specialization, then Spirit Tap', () => {
    const steps = guideLadderSteps(
      activeBuild.build,
      'priest',
      'priest-shadow',
      readSpecCatalog().find((row) => row.class_slug === 'priest' && row.spec_slug === 'shadow')!.tree_index,
      loadGuideEntry('priest/shadow').data.build!,
    );
    expect(steps.slice(0, 3).map((step) => [step.level, step.name, step.rank])).toEqual([
      [10, 'Wand Specialization', 1],
      [11, 'Wand Specialization', 2],
      [12, 'Spirit Tap', 1],
    ]);
  });
});

describe('/guides/warrior/arms rotation rows for the raid build', () => {
  // Sim-accuracy round 16: the priority is the level 60 raid list; Reckless Swings and
  // Spearing Strike are talents only the raid build takes, so their rows say so. The rail card
  // describes the raid build (guide round 12), so it keeps them.
  it('tags the raid build’s own talent rows and no other', async () => {
    const html = await container.renderToString(SpecPage, {
      props: { entry: loadGuideEntry('warrior/arms') },
    });
    const priority = html.split('data-testid="guide-rotation-priority"')[1]?.split('</ol>')[0] ?? '';
    const rows = priority.split('data-testid="guide-rotation-priority-').slice(1);
    const tagged = rows
      .filter((row) => row.includes('data-testid="guide-rotation-raid-build"'))
      .map((row) => /class="rotation-name"[^>]*>\s*([^<]+?)\s*</.exec(row)?.[1]);
    expect(tagged.sort()).toEqual(['Reckless Swings', 'Spearing Strike']);
    expect(html).toContain('data-testid="guide-rotation-raid-build-legend"');
  });
});
