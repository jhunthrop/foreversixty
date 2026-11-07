// The served content and copy the changelog, content-page and 404 mocks render (design/specs/2026-11-08-pages.md):
//   node design/mocks/pages_facts.mjs [out.json]  ->  design/mocks/data/pages-facts.json (or out.json)
// Reads, by bundling the TypeScript modules themselves (esbuild from web/node_modules), never by retyping them:
//  - the collections the pages read: web/src/content/changelog/*.md and web/src/content/pages/*.md (frontmatter through
//    `yaml`, the body rendered by lib/guides/markdown.ts, the processor Astro's own markdown pipeline uses, so the heading
//    ids are the built page's); the changelog in changelog.astro's own order (date, newest first; ties in collection order);
//  - the privacy page's Log coach section: lib/coach/privacy-section.ts (coachPrivacySection, splitBeforeHeading), placed
//    as LegalBody.astro places it, in CoachPrivacySection.astro's markup;
//  - copy: lib/guides/copy.ts (the answer's label, the guide callout, no character, the contents label),
//    lib/account/answer-copy.ts, lib/addon/copy.ts (setupPageTitle), lib/home-landing-copy.ts (the character lookup),
//    lib/footer-copy.ts (the legal links), lib/sources.ts (pillLabelFor), lib/dates.ts (formatDate, latest), lib/characters.ts;
//  - strings that live in markup (gap PG1): changelog.astro, 404.astro, b-unavailable.astro, read from the files verbatim;
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
const read = (p) => readFileSync(join(WEB, p), "utf8");
// Markup literals, verbatim from their files (gap PG1: no copy module carries them yet).
const attr = (text, name) => new RegExp(`${name}="([^"]*)"`).exec(text)?.[1];
const changelogAstro = read("src/pages/changelog.astro");
const notFoundAstro = read("src/pages/404.astro");
const unavailableAstro = read("src/pages/b-unavailable.astro");
const markup = {
  changelog: { title: attr(changelogAstro, "title"), description: attr(changelogAstro, "description"),
    h1: /<h1[^>]*>([^<]+)<\/h1>/.exec(changelogAstro)[1].trim() },
  notFound: { docTitle: attr(notFoundAstro, "title"), description: attr(notFoundAstro, "description"), eyebrow: attr(notFoundAstro, "eyebrow"),
    h1: /<PageHead[\s\S]*?title="([^"]+)"/.exec(notFoundAstro)[1], job: attr(notFoundAstro, "job"), pick: /<Card[^>]*title="([^"]+)"/.exec(notFoundAstro)[1] },
  unavailable: { docTitle: attr(unavailableAstro, "title"), description: attr(unavailableAstro, "description"),
    h1: /<h1[^>]*>([^<]+)<\/h1>/.exec(unavailableAstro)[1].trim(),
    paragraphs: [...unavailableAstro.matchAll(/<p[^>]*>([\s\S]*?)<\/p>/g)].map((m) => m[1].replace(/\s+/g, " ").trim()) },
};

const entry = `
import classes from ${src("src/data/classes.json")};
import specs from ${src("src/data/specs.json")};
import me34 from ${src("tests/e2e/fixtures/home/me-34.json")};
import { renderGuideMarkdown } from ${src("src/lib/guides/markdown.ts")};
import { coachPrivacySection, splitBeforeHeading, COACH_SECTION_BEFORE, PRIVACY_SLUG } from ${src("src/lib/coach/privacy-section.ts")};
import { formatDate, latest } from ${src("src/lib/dates.ts")};
import { pillLabelFor } from ${src("src/lib/sources.ts")};
import { guidesCopy } from ${src("src/lib/guides/copy.ts")};
import { answerCopy } from ${src("src/lib/account/answer-copy.ts")};
import { addonCopy } from ${src("src/lib/addon/copy.ts")};
import { homeCharacterSearchCopy } from ${src("src/lib/home-landing-copy.ts")};
import { FOOTER_LEGAL_LINKS, FOOTER_COLUMNS } from ${src("src/lib/footer-copy.ts")};
import { RULESETS, REGIONS } from ${src("src/lib/characters.ts")};
import { descriptorLines } from ${src("src/lib/character-header/descriptor.ts")};
const esc = (s) => s.replace(/&/g, '&amp;').replace(/</g, '&lt;').replace(/>/g, '&gt;').replace(/"/g, '&quot;');
// CoachPrivacySection.astro's markup (its h2 carries no id: gap PG3).
function coachHtml(s) {
  return '<h2>' + esc(s.heading) + '</h2>' + s.paragraphs.map((p) => '<p>' + esc(p) + '</p>').join('')
    + '<ul>' + s.links.map((l) => '<li><a href="' + esc(l.href) + '">' + esc(l.label) + '</a></li>').join('') + '</ul>';
}
const h2s = (html) => [...html.matchAll(/<h2(?: id="([^"]*)")?>([\\s\\S]*?)<\\/h2>/g)].map((m) => ({ id: m[1] ?? null, title: m[2].replace(/<[^>]+>/g, '') }));
export async function run(changelog, pages) {
  // changelog.astro: sorted by date, newest first; Updated is the newest entry's updated.
  const sorted = [...changelog].sort((a, b) => new Date(b.data.date).getTime() - new Date(a.data.date).getTime());
  const entries = [];
  for (const e of sorted) {
    const date = new Date(e.data.date);
    entries.push({ id: e.id, iso: date.toISOString().slice(0, 10), date: formatDate(date), title: e.data.title, kind: e.data.kind,
      kindLabel: pillLabelFor(e.data.kind), note: e.data.note ?? null, html: await renderGuideMarkdown(e.body),
      sources: e.data.sources.map((s) => ({ ...s, kindLabel: pillLabelFor(s.kind) })) });
  }
  const content = {};
  for (const p of pages) {
    let html = await renderGuideMarkdown(p.body);
    if (p.id === PRIVACY_SLUG) { const [head, tail] = splitBeforeHeading(html, COACH_SECTION_BEFORE); html = head + coachHtml(coachPrivacySection()) + tail; }
    content[p.id] = { id: p.id, title: p.data.title, description: p.data.description ?? '', updated: formatDate(new Date(p.data.updated)),
      confidence: p.data.confidence ?? null, sources: p.data.sources, html, headings: h2s(html) };
  }
  const c = me34.characters[0];
  const d = descriptorLines(c);
  return {
    changelog: { updated: formatDate(latest(changelog.map((e) => new Date(e.data.updated)))), entries },
    pages: content,
    copy: {
      answerLabel: guidesCopy.answerLabel, noCharacter: guidesCopy.noCharacter, tocBarLabel: guidesCopy.tocBarLabel, sources: guidesCopy.sources,
      callout: guidesCopy.signedInCallout(c.spec, c.class), answer: answerCopy, setupTitle: addonCopy.setupPageTitle,
      lookup: homeCharacterSearchCopy, legalLinks: FOOTER_LEGAL_LINKS, aboutLinks: FOOTER_COLUMNS.find((x) => x.heading === 'About').links,
      rulesets: RULESETS, regions: REGIONS, sourceLabels: Object.fromEntries(['blizzard', 'datamined', 'community', 'site'].map((k) => [k, pillLabelFor(k)])),
    },
    classes: classes.map((k) => ({ ...k, specs: specs.filter((s) => s.class_slug === k.slug).sort((a, b) => a.tree_index - b.tree_index).map((s) => ({ name: s.name, slug: s.spec_slug, icon: s.icon })) })),
    me: { name: c.name, class: c.class, spec: c.spec, level: c.level, identity: d.identity, level_line: d.level },
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
const facts = { ...(await run(collection("src/content/changelog"), collection("src/content/pages"))), markup };
writeFileSync(OUT, JSON.stringify(facts, null, 1) + "\n");
console.log(OUT.pathname);
