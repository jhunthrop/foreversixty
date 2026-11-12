// The served content and copy the changelog, content-page and 404 mocks render (design/specs/2026-11-08-pages.md):
//   node design/mocks/pages_facts.mjs [out.json]  ->  design/mocks/data/pages-facts.json (or out.json)
// Reads, by bundling the TypeScript modules themselves (esbuild from web/node_modules), never by retyping them:
//  - the collections the pages read: web/src/content/changelog/*.md and web/src/content/pages/*.md (frontmatter through
//    `yaml`, the body rendered by lib/guides/markdown.ts, the processor Astro's own markdown pipeline uses, so the heading
//    ids are the built page's); the changelog in changelog.astro's own order (date, newest first; ties in collection order);
//  - the changelog's model: lib/changelog/entries.ts (sortedEntries, changelogMonths, showCounts: the order, months, counts,
//    folded pre-launch entries, the sources lane and the site entries' action links) and lib/changelog/coming-up.ts;
//  - the privacy page's Log coach section: lib/coach/privacy-section.ts (withCoachSection, its h2 carries its id);
//  - copy: lib/guides/copy.ts (the answer's label, the guide callout, no character, the contents label),
//    lib/account/answer-copy.ts, lib/addon/copy.ts (setupPageTitle), lib/home-landing-copy.ts (the character lookup),
//    lib/footer-copy.ts (the legal links), lib/sources.ts (pillLabelFor), lib/dates.ts (formatDate, latest), lib/characters.ts,
//    lib/changelog/copy.ts and lib/not-found/copy.ts (the pages' own strings, gap PG1), lib/not-found/suggest.ts (knownPages);
//  - data: src/data/classes.json, specs.json; the signed-in fixture web/tests/e2e/fixtures/home/me-34.json (descriptorLines).
import { createRequire } from "node:module";
import { existsSync, mkdtempSync, readFileSync, readdirSync, writeFileSync } from "node:fs";
import { tmpdir } from "node:os";
import { basename, join } from "node:path";
import { pathToFileURL } from "node:url";

const WEB = new URL("../../web/", import.meta.url).pathname;
const require = createRequire(join(WEB, "package.json"));
const esbuild = require("esbuild");
const YAML = require("yaml");
const OUT = process.argv[2] ? pathToFileURL(process.argv[2]) : new URL("./data/pages-facts.json", import.meta.url);
const src = (p) => JSON.stringify(WEB + p);

function collection(dir) {
  return readdirSync(join(WEB, dir)).filter((f) => f.endsWith(".md")).sort().map((f) => {
    const raw = readFileSync(join(WEB, dir, f), "utf8");
    const m = /^---\n([\s\S]*?)\n---\n?([\s\S]*)$/.exec(raw);
    return { id: f.replace(/\.md$/, ""), data: YAML.parse(m[1]), body: m[2] };
  });
}

const entry = `
import classes from ${src("src/data/classes.json")};
import specs from ${src("src/data/specs.json")};
import me34 from ${src("tests/e2e/fixtures/home/me-34.json")};
import { renderGuideMarkdown } from ${src("src/lib/guides/markdown.ts")};
import { coachPrivacySection, withCoachSection } from ${src("src/lib/coach/privacy-section.ts")};
import { formatDate, latest } from ${src("src/lib/dates.ts")};
import { pillLabelFor } from ${src("src/lib/sources.ts")};
import { guidesCopy } from ${src("src/lib/guides/copy.ts")};
import { answerCopy } from ${src("src/lib/account/answer-copy.ts")};
import { addonCopy } from ${src("src/lib/addon/copy.ts")};
import { homeCharacterSearchCopy } from ${src("src/lib/home-landing-copy.ts")};
import { FOOTER_LEGAL_LINKS, FOOTER_COLUMNS } from ${src("src/lib/footer-copy.ts")};
import { RULESETS, REGIONS, characterHref } from ${src("src/lib/characters.ts")};
import { descriptorLines } from ${src("src/lib/character-header/descriptor.ts")};
import { changelogCopy } from ${src("src/lib/changelog/copy.ts")};
import { changelogMonths, showCounts, sortedEntries, SHOWS } from ${src("src/lib/changelog/entries.ts")};
import { COMING_UP, COMING_UP_HEADS, COMING_UP_SOURCE } from ${src("src/lib/changelog/coming-up.ts")};
import { notFoundCopy } from ${src("src/lib/not-found/copy.ts")};
import { knownPages, suggestPage } from ${src("src/lib/not-found/suggest.ts")};
import { splitDocument } from ${src("src/lib/content-pages/sections.ts")};
import { contentPagesCopy } from ${src("src/lib/content-pages/copy.ts")};
import { accountSignInCopy } from ${src("src/lib/account/signin-copy.ts")};
import { myReportsCopy } from ${src("src/lib/reports/my-reports-copy.ts")};
const h2s = (html) => splitDocument(html).sections.map((s) => ({ id: s.id, title: s.title }));
const bySlug = (slug) => classes.find((c) => c.slug === slug);
export async function run(changelog, pages) {
  // changelog.astro: the collection through lib/changelog/entries.ts, each body rendered as the build renders it.
  const inputs = [];
  for (const e of changelog) inputs.push({ id: e.id, data: { ...e.data, date: new Date(e.data.date) }, html: await renderGuideMarkdown(e.body) });
  const entries = sortedEntries(inputs);
  const counts = showCounts(entries);
  const months = changelogMonths(entries).map((m) => ({ ...m, meta: Object.fromEntries(SHOWS.map((s) => [s, changelogCopy.changes(m.counts[s])])) }));
  const content = {};
  for (const p of pages) {
    const html = withCoachSection(p.id, await renderGuideMarkdown(p.body), coachPrivacySection());
    content[p.id] = { id: p.id, title: p.data.title, description: p.data.description ?? '', updated: formatDate(new Date(p.data.updated)),
      confidence: p.data.confidence ?? null, sources: p.data.sources, html, headings: h2s(html),
      sectionsMeta: contentPagesCopy.sections(h2s(html).length) };
  }
  const c = me34.characters[0];
  const d = descriptorLines(c);
  const specPages = specs.map((s) => ({ classSlug: s.class_slug, className: bySlug(s.class_slug)?.name ?? s.class_slug, specSlug: s.spec_slug, specName: s.name }));
  const known = knownPages(classes, specPages);
  const SUGGEST_FROM = '/guides/hunter/beastmastery';   // SAMPLE (spec 4.4 N3): a mistyped guide address
  return {
    changelog: { updated: formatDate(latest(changelog.map((e) => new Date(e.data.updated)))), counts, months,
      comingUp: { heads: COMING_UP_HEADS, rows: COMING_UP, source: { ...COMING_UP_SOURCE, kindLabel: pillLabelFor(COMING_UP_SOURCE.kind) } } },
    pages: content,
    copy: {
      answerLabel: guidesCopy.answerLabel, noCharacter: guidesCopy.noCharacter, tocBarLabel: guidesCopy.tocBarLabel, sources: guidesCopy.sources,
      callout: guidesCopy.signedInCallout(c.spec, c.class), answer: answerCopy, setupTitle: addonCopy.setupPageTitle,
      lookup: homeCharacterSearchCopy, legalLinks: FOOTER_LEGAL_LINKS, aboutLinks: FOOTER_COLUMNS.find((x) => x.heading === 'About').links,
      rulesets: RULESETS, regions: REGIONS, sourceLabels: Object.fromEntries(['blizzard', 'datamined', 'community', 'site'].map((k) => [k, pillLabelFor(k)])),
      changelog: { ...changelogCopy, summary: Object.fromEntries(SHOWS.map((s) => [s, changelogCopy.summary(s, counts[s])])) },
      notFound: { ...notFoundCopy },
      legalTabsLabel: contentPagesCopy.legalTabsLabel,
      login: { ...accountSignInCopy, action: myReportsCopy.signInAction },
    },
    // One spec per talent tree, as every picker lists them (lib/class-picker.ts treeSpecs: Feral Bear folds into Feral).
    classes: classes.map((k) => ({ ...k, specs: specs.filter((s) => s.class_slug === k.slug).sort((a, b) => a.tree_index - b.tree_index).filter((s, i, all) => all.findIndex((o) => o.tree_index === s.tree_index) === i).map((s) => ({ name: s.name, slug: s.spec_slug, icon: s.icon })) })),
    me: { name: c.name, class: c.class, spec: c.spec, level: c.level, identity: d.identity, level_line: d.level, href: characterHref(c.region, c.ruleset, c.name) },
    suggest: { from: SUGGEST_FROM, page: suggestPage(SUGGEST_FROM, known), known: known.length },
  };
}`;
const generatedFallback = {
  name: "generated-fallback",
  setup(build) {
    build.onResolve({ filter: /data\/generated\/[^/]+\.json$/ }, (args) => {
      const wanted = join(args.resolveDir, args.path);
      if (existsSync(wanted)) return { path: wanted };
      return { path: join(WEB, "src/fixtures/planner", basename(args.path)) };
    });
  },
};
const dir = mkdtempSync(join(tmpdir(), "pages-facts-"));
const outfile = join(dir, "facts.mjs");
const metaUrl = pathToFileURL(WEB + "src/lib/guides/build-files.ts").href;
await esbuild.build({ stdin: { contents: entry, resolveDir: WEB, loader: "ts" }, bundle: true, platform: "node", format: "esm", outfile,
  logLevel: "error", loader: { ".json": "json" }, plugins: [generatedFallback], external: ["@astrojs/markdown-satteri", "@playwright/test"],
  define: { "import.meta.url": JSON.stringify(metaUrl), "import.meta.env": "{}" } });
// The one external (the markdown processor, a native package) resolves from web/node_modules.
const bundled = readFileSync(outfile, "utf8").replace(/from "@astrojs\/markdown-satteri"/g, `from ${JSON.stringify(pathToFileURL(require.resolve("@astrojs/markdown-satteri")).href)}`);
writeFileSync(outfile, bundled);
const { run } = await import(pathToFileURL(outfile).href);
const facts = { ...(await run(collection("src/content/changelog"), collection("src/content/pages"))) };
writeFileSync(OUT, JSON.stringify(facts, null, 1) + "\n");
console.log(OUT.pathname);
