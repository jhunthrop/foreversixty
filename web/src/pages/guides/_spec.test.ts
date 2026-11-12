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
//
// The page is the flat system's reading page (design/specs/2026-11-08-guides.md §4.3): the
// PageHeader with its lead and the signed-in answer, the ControlBar, the one Plate (At a
// glance), the nine SPEC_SECTIONS as ReadingSections (lane before prose), Sources, the data line.
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
import { guidesCopy } from '../../lib/guides/copy';
import { SPEC_SECTIONS, slugify } from '../../lib/guides/sections';
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

const rendered = new Map<string, Promise<string>>();

/** One render per guide for the whole file: the page is static, so every test reads the same HTML. */
function pageHtml(id: string): Promise<string> {
  const cached = rendered.get(id);
  if (cached !== undefined) return cached;
  const html = container.renderToString(SpecPage, { props: { entry: loadGuideEntry(id) } });
  rendered.set(id, html);
  return html;
}

/** The page's `<main>`, without its scripts, styles and Svelte's hydration comments. */
function mainOf(html: string): string {
  return html
    .slice(html.indexOf('<main'), html.lastIndexOf('</main>'))
    .replace(/<script[\s\S]*?<\/script>|<style[\s\S]*?<\/style>|<!--[^>]*-->/g, '');
}

/** One ReadingSection's markup, from its opening tag to the next section. */
function sectionOf(html: string, anchor: string): string {
  const start = html.indexOf(`<section id="${anchor}"`);
  if (start < 0) throw new Error(`no section #${anchor}`);
  const next = html.indexOf('<section', start + 1);
  return html.slice(start, next < 0 ? undefined : next);
}

/** A data-testid block's markup up to the next block of the same kind, or `length` characters. */
function blockOf(html: string, testid: string, length = 40_000): string {
  const start = html.indexOf(`data-testid="${testid}"`);
  if (start < 0) throw new Error(`no ${testid}`);
  return html.slice(start, start + length);
}

const ANCHORS = SPEC_SECTIONS.map(slugify);

describe('/guides/warrior/fury (the page structure)', () => {
  it('runs the nine sections of SPEC_SECTIONS as ReadingSections, in order, each with its anchor and region', async () => {
    const main = mainOf(await pageHtml('warrior/fury'));
    const sections = [...main.matchAll(/<section id="([^"]+)" class="ui-rsec[^"]*" data-region="([^"]+)"/g)];
    expect(sections.map((m) => m[1])).toEqual(ANCHORS);
    expect(sections.map((m) => m[2])).toEqual(ANCHORS.map((anchor) => `section-${anchor}`));
    for (const section of SPEC_SECTIONS) {
      expect(main, section).toContain(`<h2 class="ui-rsec-title`);
      expect(main, section).toMatch(new RegExp(`>${section}</h2>`));
    }
  });

  it('links the contents rail to exactly those anchors, in order, the first current', async () => {
    const main = mainOf(await pageHtml('warrior/fury'));
    const toc = blockOf(main, 'guide-toc', 4000).split('</ol>')[0]!;
    expect([...toc.matchAll(/href="#([^"]+)"/g)].map((m) => m[1])).toEqual(ANCHORS);
    expect(toc).toMatch(/href="#overview" data-toc-link="overview" aria-current="location"/);
    expect(main).toMatch(/<nav class="ui-page-toc[^"]*" aria-label="Contents">/);
  });

  it('keeps the spec regions in the spec’s order: head, controls, plate, the sections, Sources, the data line', async () => {
    const main = mainOf(await pageHtml('warrior/fury'));
    const regions = [...main.matchAll(/data-region="([^"]+)"/g)].map((m) => m[1]);
    expect(regions).toEqual([
      'head',
      'head-title',
      'head-meta',
      'head-answer',
      'controls',
      'plate',
      // The Best in slot section's doll (the shared PaperdollFrame) carries its own three regions.
      ...ANCHORS.flatMap((anchor) =>
        anchor === 'best-in-slot'
          ? [`section-${anchor}`, 'slots', 'paperdoll-medallion', 'paperdoll-centre']
          : [`section-${anchor}`],
      ),
      'section-sources',
    ]);
  });

  it('draws one Plate, Summary, with the described build’s trees and its Load and Sim buttons', async () => {
    const entry = loadGuideEntry('warrior/fury');
    const html = await pageHtml('warrior/fury');
    const main = mainOf(html);
    expect(main.match(/data-region="plate"/g)).toHaveLength(1);
    expect(main).toMatch(/role="region" aria-label="Summary" data-testid="guide-glance"/);
    const glance = blockOf(main, 'guide-glance-build', 200_000).split(
      'data-testid="guide-glance-rotation"',
    )[0]!;
    // The trees are server-rendered: a grid per tree, a button per talent naming its rank.
    expect(glance.match(/role="grid"/g)).toHaveLength(3);
    expect(glance).toMatch(/class="tt-cell[^"]*" aria-label="[^"]+, rank \d of \d"/);
    const code = encodeURIComponent(entry.data.build as string);
    expect(glance).toMatch(
      new RegExp(
        `href="/planner\\?code=${code}(?:&amp;spec=[a-z-]+)?"[^>]*data-testid="guide-load-build"|data-testid="guide-load-build"[^>]*href="/planner\\?code=${code}(?:&amp;spec=[a-z-]+)?"`,
      ),
    );
    expect(glance).toContain('data-testid="guide-sim-build"');
  });

  it('sets the PageHeader (one title in the class colour, the meta line, the answer reserved) and the lede under it', async () => {
    const main = mainOf(await pageHtml('warrior/fury'));
    expect(main.match(/<h1/g)).toHaveLength(1);
    expect(main).toMatch(/<h1[^>]*style="color: var\(--color-class-warrior\);">Fury Warrior<\/h1>/);
    expect(main).toMatch(/data-region="head-meta"[^>]*><span[^>]*>Updated \d{1,2} [A-Z][a-z]{2}</);
    expect(main).not.toContain('ui-page-header-lead');
    // No lede under the band (owner 2026-11-09: the "<Spec> is the" sentences were struck from every guide).
    expect(main).not.toContain('data-testid="guide-summary"');
    const answer = blockOf(main, 'guide-answer', 2000);
    expect(main.indexOf('data-region="head-answer"')).toBeLessThan(main.indexOf('id="guide-answer"'));
    expect(main).toMatch(
      /id="guide-answer"[^>]*data-signed-in-only[^>]*data-state="loading"[^>]*data-page="[^"]*classSlug&quot;:&quot;warrior&quot;,&quot;specSlug&quot;:&quot;fury&quot;,&quot;ladder&quot;:\[\{&quot;level&quot;:10/,
    );
    expect(answer).toContain(guidesCopy.answerLabel);
  });

  it('retires the old head and rails: no ClassHeader, ArtPanel, SpecTableOfContents or aside', async () => {
    const html = mainOf(await pageHtml('warrior/fury'));
    for (const retired of [
      'class-header',
      'art-panel',
      'guide-header-meta',
      'bis-spec-tab-',
      'guide-toc-inline',
      'guide-toc-rail',
      'content-aside',
      'guide-action-rail',
      'stat-priority-pills',
      'guide-race-pills',
      'confidence-footer-note',
      'Our inference',
    ]) {
      expect(html, retired).not.toContain(retired);
    }
  });

  it('puts each lane before its prose in the DOM; Professions and Leveling read in one column', async () => {
    const main = mainOf(await pageHtml('warrior/fury'));
    for (const anchor of ANCHORS.filter((a) => a !== 'professions' && a !== 'leveling')) {
      const section = sectionOf(main, anchor);
      expect(section, anchor).toMatch(/class="ui-rd-lane[^"]*"[\s\S]*class="ui-rd-prose/);
    }
    for (const anchor of ['professions', 'leveling']) {
      const section = sectionOf(main, anchor);
      expect(section, anchor).toContain('data-solo=""');
      expect(section, anchor).not.toContain('ui-rd-lane');
    }
    // The leveling tree sits above its prose, full width: no rail beside nine paragraphs.
    expect(sectionOf(main, 'leveling')).toMatch(/data-testid="guide-talents-by-level"[\s\S]*class="ui-prose/);
  });
});

describe('/guides/rogue/assassination (a spec with a raid build)', () => {
  it('lists the leveling build and the raid build as two labelled rows with their own links', async () => {
    const entry = loadGuideEntry('rogue/assassination');
    const raidBuild = entry.data.raidBuild as string;
    expect(raidBuild).not.toBe(entry.data.build);
    const main = mainOf(await pageHtml('rogue/assassination'));
    const builds = blockOf(main, 'guide-builds', 6000);
    expect(builds).toMatch(/data-testid="guide-build-leveling"[\s\S]*?>Level 60 build</);
    expect(builds).toMatch(/data-testid="guide-build-raid"[\s\S]*?>Raid build</);
    // The planner link may carry `&amp;spec=` after the code (the guide's own levelling order).
    expect(builds).toContain(`href="/planner?code=${encodeURIComponent(raidBuild)}`);
    expect(builds).toContain(`href="/sim?code=${encodeURIComponent(raidBuild)}"`);
    expect(builds).toContain('data-testid="guide-build-load-raid"');
    expect(builds).toContain('data-testid="guide-build-sim-raid"');
    expect(builds).not.toContain('data-testid="guide-build-both"');
    // The plate describes the raid build, under its own label.
    expect(blockOf(main, 'guide-glance-build', 400)).toContain(`>${guidesCopy.raidBuild}</h2>`);
    expect(main).toMatch(
      new RegExp(
        `href="/planner\\?code=${encodeURIComponent(raidBuild)}(?:&amp;spec=[a-z-]+)?"[^>]*data-testid="guide-load-build"|data-testid="guide-load-build"[^>]*href="/planner\\?code=${encodeURIComponent(raidBuild)}(?:&amp;spec=[a-z-]+)?"`,
      ),
    );
  });
});

describe('/guides/warrior/fury (a spec whose raid build is its leveling build)', () => {
  it('leaves one build to the band: no Builds block repeats its Load and Sim', async () => {
    const entry = loadGuideEntry('warrior/fury');
    expect(entry.data.raidBuild).toBeUndefined();
    const main = mainOf(await pageHtml('warrior/fury'));
    expect(main).not.toContain('data-testid="guide-builds"');
    expect(main).not.toContain('data-testid="guide-build-raid"');
    expect(blockOf(main, 'guide-glance-build', 400)).toContain(`>${guidesCopy.buildRailLabel}</h2>`);
  });
});

describe('/guides/mage/frost (the class-guide blocks)', () => {
  let frost = '';
  beforeAll(async () => {
    frost = await pageHtml('mage/frost');
  });

  it('keeps sim talk out of the reading flow: one data-source line, provenance only in Sources', () => {
    expect(frost).not.toContain('data-testid="guide-status-note"');
    const main = mainOf(frost);
    // The guide's own reading flow, Sources aside (its kind tags name "This site", spec §4.3).
    const reading = main.slice(0, main.indexOf('data-testid="guide-sources"'));
    const visible = reading.replace(/<[^>]+>/g, ' ');
    const talk = [...visible.matchAll(/.{0,60}(?:simulator|simulated|this site).{0,40}/gi)];
    expect(talk.map((m) => m[0])).toEqual([]);
    expect(main).not.toContain('Numbers from client build');
    expect(main.indexOf('data-testid="guide-sources"')).toBeGreaterThan(
      main.indexOf('<section id="leveling"'),
    );
    expect(frost).toMatch(/>Updated \d{1,2} [A-Z][a-z]{2}</);
    expect(frost).not.toContain('DPS · Updated');
    // No lede: neither the Overview's old sentence nor the frontmatter description shows under the band.
    expect(frost).not.toContain('Frost casts Frostbolt behind Ice Barrier');
  });

  it('lays the Overview out as strengths and weaknesses beside the raid role lane', () => {
    const overview = sectionOf(mainOf(frost), 'overview');
    expect(overview).toMatch(
      /class="ui-rd-lane[^"]*"[^>]*><div class="ui-lane-block[^"]*" data-testid="guide-overview-raid"/,
    );
    expect(overview).toContain(`>${guidesCopy.inARaid}</h3>`);
    expect(overview).toMatch(/class="ui-prose ui-prose--plus[^"]*" data-testid="guide-strengths"/);
    expect(overview).toMatch(/class="ui-prose ui-prose--minus[^"]*" data-testid="guide-weaknesses"/);
  });

  it('names the raid build’s difference in one line and lists the key talents with icons', async () => {
    // Arcane's raid build differs from its leveling build, and since sim-accuracy 2026-11-12
    // (casters) so does Frost's: the raid build keeps Incineration where the level 60 build
    // takes Arcane Meditation and the Improved Channeling toll.
    const arcane = await pageHtml('mage/arcane');
    expect(arcane).toMatch(/data-testid="guide-raid-diff"[^>]*>The raid build moves \d+ points? from /);
    const talents = sectionOf(mainOf(frost), 'talents-and-builds');
    expect(talents).toContain('data-testid="guide-builds"');
    expect(talents).toMatch(/data-testid="guide-raid-diff"[^>]*>The raid build moves 5 points from /);
    const keys = blockOf(talents, 'guide-key-talents', 20_000);
    expect(keys).toContain('Improved Frostbolt');
    expect(keys).toMatch(/<img src="\/data\/[^"]+\/icons\/[^"]+\.webp"/);
    expect(keys).toMatch(/data-game-tip="talent:\d+:\d+"/);
    // Two full guide renders in one test.
  }, 20_000);

  it('numbers a priority with spell icons, names the real cooldowns, and prints the pre-pull Ice Barrier as the opener', () => {
    const rotation = sectionOf(mainOf(frost), 'rotation-and-priority');
    const priority = blockOf(rotation, 'guide-rotation-priority').split(
      'data-testid="guide-rotation-cooldowns"',
    )[0]!;
    expect(priority).toMatch(/<img[^>]+icons\/spell_frost_frostblast\.webp/);
    expect(priority).toContain('Ice Lance');
    expect(priority).toMatch(/tabular-nums">1<\/span>/);
    const opener = blockOf(rotation, 'guide-rotation-opener').split(
      'data-testid="guide-rotation-priority"',
    )[0]!;
    const firstOpener = opener.split('data-testid="guide-rotation-opener-row"')[1]!;
    expect(firstOpener).toContain('Ice Barrier');
    expect(rotation).not.toContain('Then work down it from the top');
    const cooldowns = blockOf(rotation, 'guide-rotation-cooldowns').split('class="ui-rd-prose')[0]!;
    for (const name of ['Evocation', 'Cold Snap', 'Ice Block'])
      expect(cooldowns, name).toContain(`>${name}</span>`);
    expect(cooldowns).toContain('8 min cooldown');
    // Evocation has its icon from the spell icon data, so every spell row shares the left edge.
    expect(cooldowns).toMatch(
      /icons\/spell_nature_purge\.webp[^>]*>(?:(?!guide-rotation-cooldowns-row)[\s\S])*?>Evocation</,
    );
    // The trinket line names the racial cooldowns the sim casts for the mage's races.
    expect(cooldowns).toMatch(
      /Trinkets, racials and potions(?:(?!guide-rotation-cooldowns-row)[\s\S])*?Eureka! \(Gnome\)/,
    );
  });

  it('gives the stat lane the hit line as its headline and a weights table; the plate ranks the same order', () => {
    const main = mainOf(frost);
    const lane = blockOf(main, 'guide-stat-table', 60_000).split('class="ui-rd-prose')[0]!;
    // The line names every source the band's figure counts: the BiS gear and buffs, and each
    // talent of the weights character's own build that adds hit (Precision, for Frost).
    const file = rankedBisFile(activeBuild.build, 'mage-frost');
    const band = file === undefined ? undefined : bandEntry(file, 60, 'alliance');
    const line = guideHitCapLine(
      band?.hit_to_cap,
      band === undefined ? [] : bandTalentHits(activeBuild.build, 'mage', band.talents),
    );
    expect(line).toBeDefined();
    // The cap is the headline; the breakdown follows it, muted (round 20).
    expect(lane).toMatch(
      /data-testid="guide-stat-table-hit-cap"[^>]*><b class="text-strong"[^>]*>([^<]+)<\/b>/,
    );
    expect(lane).toContain(`>${line!.cap}</b> · ${line!.detail}</p>`);
    expect(line!.detail).toContain('from Precision');
    expect(lane).toContain('DPS per point');
    // No "1 point of hit = 0.79 DPS at level 60" line: the per-point column already says it.
    expect(lane).not.toMatch(/>1 (?:point of [a-z][a-z ]*|[a-z][a-z ]* rating) = /);
    // No footnotes under the table, and no row for a stat no item carries: gear is the only
    // choice the table speaks to; a veteran knows what a rating point is (h-guide-spec 2026-11-10).
    expect(lane).not.toMatch(/>[^<]*per rating point/);
    expect(lane).not.toContain('but no item at level 60 has it');
    expect(lane).not.toContain('from school');
    expect(lane).not.toContain('Not on any item');

    // The plate ranks exactly the rows the table gives a number, in the table's order.
    const glance = blockOf(main, 'guide-glance-stats', 60_000).split(
      'data-testid="guide-glance-hit-cap"',
    )[0]!;
    const ranked = [
      ...glance.matchAll(
        /guide-glance-rank[^>]*>(\d+)<\/span>[\s\S]*?class="ui-st-name" role="cell"><span[^>]*>([^<]+)</g,
      ),
    ].map((m) => m[2]);
    const numberedRows = [
      ...lane.matchAll(
        /data-testid="guide-stat-table-row-\d+">([\s\S]*?)(?=data-testid="guide-stat-table-row-|$)/g,
      ),
    ]
      .filter((m) => !m[1]!.includes('data-muted'))
      .map((m) => /class="ui-st-name" role="cell"><span[^>]*>([^<]+)</.exec(m[1]!)?.[1]);
    expect(ranked.length).toBeGreaterThan(3);
    expect(ranked).toEqual(numberedRows);
    // A row a player cannot gear for is in neither.
    expect(glance).not.toContain('Not on any item');
    // The plate keeps the hit cap under its ranks.
    expect(blockOf(main, 'guide-glance-hit-cap', 400)).toContain(`>${line!.cap}</b>`);
  });

  it('gives the plate a Rotation column: the first four priority lines and a Full priority link', () => {
    const main = mainOf(frost);
    const column = blockOf(main, 'guide-glance-rotation', 60_000).split(
      'data-testid="guide-glance-stats"',
    )[0]!;
    expect(column).toContain(`>${guidesCopy.rotationRailTitle}</h2>`);
    const rows = column.split('data-testid="guide-full-priority"')[0]!;
    expect(rows.match(/class="ui-spell-icon/g)).toHaveLength(4);
    expect(column).toMatch(/href="#rotation-and-priority"[^>]*data-testid="guide-full-priority"/);
    expect(column).toContain(guidesCopy.fullPriorityLink);
    // Each line reads at its level 60 clause: no rank swap a leveller makes.
    expect(rows).not.toMatch(/\b(?:at|to|from) \d{2}\b/);
  });

  it('puts the level 60 best in slot in the lane, every other band and faction ready as a template', () => {
    const bis = sectionOf(mainOf(frost), 'best-in-slot');
    expect(bis).toMatch(/data-testid="guide-bis" data-guide-class="mage" data-default-faction="alliance"/);
    expect(bis).toMatch(
      /data-bis-current data-panel="60-alliance"[^>]*><div class="ui-lane-block[^"]*" data-testid="guide-bis-panel"/,
    );
    const current = bis.slice(bis.indexOf('data-bis-current'), bis.indexOf('<template'));
    expect(current).toContain('>Level 60</h3>');
    expect(current).toContain('data-testid="guide-bis-head"');
    const templates = [...bis.matchAll(/<template data-bis-panel="([^"]+)"/g)].map((m) => m[1]).sort();
    const expected = ['alliance', 'horde']
      .flatMap((faction) => [20, 30, 40, 50, 60].map((band) => `${band}-${faction}`))
      .filter((key) => key !== '60-alliance')
      .sort();
    expect(templates).toEqual(expected);
    expect(bis).toMatch(
      /href="\/bis\/mage\/frost\/60"[^>]*data-testid="guide-bis-full-list"[^>]*>Full list and alternatives/,
    );
    // Every band's own BiS page and the role's tier list (SEO lane 2026-11-08).
    for (const href of [
      '/bis/mage/frost',
      '/bis/mage/frost/30-39',
      '/bis/mage/frost/40-49',
      '/bis/mage/frost/50-59',
    ])
      expect(bis).toContain(`href="${href}"`);
    expect(bis).toMatch(/href="\/tiers"[^>]*data-testid="guide-tier-list"/);
    expect(bis).not.toContain('named honestly');
    // The lane's list is server-rendered; only the doll in the prose column hydrates.
    expect(bis.split('class="ui-rd-prose')[0]).not.toMatch(/<astro-island/);
    expect(bis).toContain('data-testid="guide-bis-doll"');
  });

  it('adds the enchant and consumable choices and the race lane', () => {
    const main = mainOf(frost);
    const extras = sectionOf(main, 'enchants-and-consumables');
    expect(extras).toContain('data-testid="guide-enchants"');
    expect(extras).toContain('data-testid="guide-consumables"');
    expect(extras).toContain('Elixir of Frost Power');
    // A Classic flask stacks with the elixirs: listed beside them, never instead of them.
    expect(extras).toContain('Flask of Supreme Power');
    expect(frost).not.toContain('replaces both elixirs');
    expect(frost).not.toContain('Without a flask');
    expect(frost).not.toContain('The caster meal');
    const races = sectionOf(main, 'races');
    expect(races).toContain('data-testid="guide-race-table"');
    // No pill and no unexplained highlight: the prose names the picks.
    expect(races).not.toContain('>Recommended<');
    expect(races).not.toContain('aria-current="true"');
    // The guide's own Races prose names both picks, so no pick line repeats it in the lane.
    expect(races).not.toContain('data-testid="guide-race-picks"');
    // The racials come from the race data of the sync that ran: the real build's sentences, or
    // none under the unit-test fixture.
    const gnome = raceRows.find((race) => race.slug === 'gnome');
    if (gnome !== undefined && racialEffects(gnome).has('Eureka!')) {
      expect(frost).toContain(
        'Makes your damaging abilities cost 10% less Mana and deal 10% more damage for 8 seconds',
      );
      expect(frost).not.toContain('next 3 abilities');
      // The prose names Eureka! as the client has it (8 s, 10% damage, 10% cheaper), never the
      // BlizzCon demo's reading.
      expect(frost).not.toMatch(/Eureka![^.<]*50%/);
    }
  });

  it('server-renders the leveling tree at level 60 in the Leveling lane: caption, slider, trees and link', () => {
    const leveling = sectionOf(mainOf(frost), 'leveling');
    expect(leveling).toMatch(/data-testid="guide-talents-by-level"[\s\S]*?>Talents by level<\/h3>/);
    expect(leveling).toContain('data-testid="guide-leveling-tree"');
    expect(leveling).toContain('client="visible"');
    expect(leveling).toMatch(/data-testid="guide-leveling-caption"[\s\S]*?Level 60: [^<]+, rank \d/);
    // Every point placed at 60: the caption says only the level's talent.
    expect(leveling).not.toContain('points spent');
    expect(leveling).toMatch(/value="60" aria-label="Level" aria-valuetext="Level 60"/);
    expect(leveling).toMatch(/data-testid="guide-leveling-trees"[\s\S]*?role="grid"/);
    expect(leveling).toMatch(
      /href="\/planner\?spec=mage-frost&amp;talents=[0-9-]+(?:&amp;spec=[a-z-]+)?"[^>]*data-testid="guide-leveling-load"|data-testid="guide-leveling-load"[^>]*href="\/planner\?spec=mage-frost&amp;talents=/,
    );
    expect(leveling).not.toContain('guide-leveling-strip');
  });

  it('opens #leveling on the tree lane, then the Levels 1 to 19 text, then Levels 20 to 60', () => {
    const at = (needle: string): number => frost.indexOf(needle);
    expect(at('id="leveling"')).toBeGreaterThan(-1);
    expect(at('data-testid="guide-talents-by-level"')).toBeGreaterThan(at('id="leveling"'));
    expect(at('id="levels-1-to-19"')).toBeGreaterThan(at('data-testid="guide-leveling-tree"'));
    expect(at('id="levels-20-to-60"')).toBeGreaterThan(at('id="levels-1-to-19"'));
  });

  it('renders no per-section hedge', () => {
    for (const hedge of ['is a projection', 'not tested', 'cannot be named', 'inferred from 1.12']) {
      expect(frost, hedge).not.toContain(hedge);
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
  // Sim-accuracy round 16: the priority is the level 60 raid list; Reckless Swings is a
  // talent only the raid build takes, so its row says so. Spearing Strike was raid-only
  // until the 2026-11-12 melee lane put it in the level 60 leveling build too.
  it('tags the raid build’s own talent rows and no other', async () => {
    const rotation = sectionOf(mainOf(await pageHtml('warrior/arms')), 'rotation-and-priority');
    const priority = blockOf(rotation, 'guide-rotation-priority').split(
      'data-testid="guide-rotation-cooldowns"',
    )[0]!;
    const rows = priority.split('data-testid="guide-rotation-priority-row"').slice(1);
    const tagged = rows
      .filter((row) => row.includes(`>${guidesCopy.raidBuildTag}</span>`))
      .map((row) => /data-game-tip="[^"]+"[^>]*>([^<]+)</.exec(row)?.[1]);
    expect(tagged.sort()).toEqual(['Reckless Swings']);
    // The tag on the row says it; no grey legend explains it above the lists (h-guide-spec).
    expect(rotation).not.toContain('data-testid="guide-raid-build-legend"');
    // Tagged rows say whose list it is: no "For the raid build." line beside them.
    expect(rotation).not.toContain(`data-testid="guide-rotation-build"`);
  });
});
