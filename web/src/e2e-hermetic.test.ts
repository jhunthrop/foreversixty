// web/src/e2e-hermetic.test.ts
// Pins the three layers of tests/hermetic.ts in place, so no browser test, local or CI, can call
// the production API: every e2e spec imports the guarded `test`, every Playwright config
// launches with the resolver block, and CI's e2e dist is built against the dead API base.
import { readdirSync, readFileSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { describe, expect, it } from 'vitest';
import { API_ORIGIN_GLOBS, DEAD_API_BASE, HOST_BLOCK_RULE } from '../tests/hermetic';

const WEB = path.resolve(path.dirname(fileURLToPath(import.meta.url)), '..');
const E2E = path.join(WEB, 'tests/e2e');
const GUARDED_TEST = path.join(E2E, 'support/test.ts');
const PLAYWRIGHT_IMPORT = /^import \{([^}]*)\} from '@playwright\/test';$/gm;

function filesUnder(dir: string, pattern: RegExp): string[] {
  return readdirSync(dir, { recursive: true, encoding: 'utf8' })
    .filter((name) => pattern.test(name))
    .map((name) => path.join(dir, name));
}

/** The value bindings (not `type` ones) a file imports from `@playwright/test`. */
function playwrightValueImports(source: string): string[] {
  return [...source.matchAll(PLAYWRIGHT_IMPORT)].flatMap((match) =>
    match[1]
      .split(',')
      .map((part) => part.trim())
      .filter((part) => part !== '' && !part.startsWith('type ')),
  );
}

describe('the e2e suite is hermetic', () => {
  it('every spec imports test from support/test.ts', () => {
    const specs = filesUnder(E2E, /\.spec\.ts$/);
    expect(specs.length).toBeGreaterThan(100);
    const unguarded = specs.filter((file) => !/from '\.\/support\/test';$/m.test(readFileSync(file, 'utf8')));
    expect(unguarded).toEqual([]);
  });

  it('no file under tests/e2e takes test straight from @playwright/test', () => {
    const offenders = filesUnder(E2E, /\.ts$/)
      .filter((file) => file !== GUARDED_TEST)
      .filter((file) =>
        playwrightValueImports(readFileSync(file, 'utf8')).some(
          (binding) => binding.split(' ')[0] === 'test',
        ),
      );
    expect(offenders).toEqual([]);
  });

  it('every Playwright config launches with the production API unresolvable', () => {
    const configs = [
      ...readdirSync(WEB)
        .filter((name) => /^playwright\.[a-z.]*config\.ts$/.test(name))
        .map((name) => path.join(WEB, name)),
      ...filesUnder(path.join(WEB, 'tests/captures'), /playwright\.capture\.config\.ts$/),
    ];
    expect(configs.length).toBeGreaterThan(5);
    const unblocked = configs.filter((file) => !readFileSync(file, 'utf8').includes('HERMETIC_LAUNCH_ARGS'));
    expect(unblocked).toEqual([]);
  });

  it("Lighthouse's Chrome cannot resolve the production API either", () => {
    const config = JSON.parse(readFileSync(path.join(WEB, 'lighthouserc.json'), 'utf8')) as {
      ci: { collect: { settings: { chromeFlags?: string } } };
    };
    expect(config.ci.collect.settings.chromeFlags).toBe(`--host-resolver-rules="${HOST_BLOCK_RULE}"`);
  });

  it('CI builds the e2e and lhci dist against the dead API base', () => {
    const workflow = readFileSync(path.join(WEB, '../.github/workflows/web.yml'), 'utf8');
    const fixtureBuild = workflow.slice(workflow.indexOf('name: build the fixture dist'));
    const base = /PUBLIC_API_BASE_URL: (\S+)/.exec(fixtureBuild)?.[1];
    expect(base).toBe(DEAD_API_BASE);
  });

  it('guards the dead base and the production API by origin glob', () => {
    expect(API_ORIGIN_GLOBS).toEqual([`${DEAD_API_BASE}/**`, 'https://api.foreversixty.gg/**']);
  });
});
