// Both apps' deep-link parsers and the site's app-link path list are generated from
// apps/shared/links.json (web/scripts/gen-mobile-links.mjs). This fails when a checked-in output
// differs from a fresh generation: after editing links.json run
// `node web/scripts/gen-mobile-links.mjs` and commit the result.
import { readFileSync } from 'node:fs';
import { describe, expect, it } from 'vitest';
import {
  OUTPUTS,
  appLinkPaths,
  buildRoutes,
  codeName,
  excludedPaths,
  generate,
  renderKotlin,
  renderSwift,
  sitePagePaths,
} from '../../../scripts/gen-mobile-links.mjs';

const HINT = 'run `node web/scripts/gen-mobile-links.mjs` and commit the result';
const table = (routes: object[]) => JSON.stringify({ host: 'foreversixty.gg', routes });

describe('mobile deep-link codegen', () => {
  const generated = generate() as Record<string, string>;

  it.each(Object.values(OUTPUTS) as string[])('%s matches the generator', (file) => {
    expect(readFileSync(file, 'utf8'), HINT).toBe(generated[file]);
  });

  it('serves every route, under /app/ except the site pages', () => {
    const model = buildRoutes(
      readFileSync(new URL('../../../../apps/shared/links.json', import.meta.url), 'utf8'),
    );
    const paths = appLinkPaths(model) as string[];
    expect(paths).toContain('/app/auth');
    expect(paths).toContain('/app/raid/*/*');
    expect(paths).toContain('/app/planner/*');
    expect(sitePagePaths(model)).toEqual([
      '/b/*',
      '/bis/*/*',
      '/guides/*',
      '/guides/*/*',
      '/sim/*',
      '/reports/*',
      '/character/*/*/*',
      '/guild/*/*/*/raids/*',
      '/tiers',
    ]);
    const outside = paths.filter((p) => !p.startsWith('/app/'));
    expect(outside).toEqual(sitePagePaths(model));
  });

  it('refuses a path outside /app/, a duplicate screen and two routes for one link', () => {
    const route = { path: '/app/sim/{id}', screen: 'sim', example: '/app/sim/a' };
    expect(() => buildRoutes(table([{ ...route, path: '/raid/{id}' }]))).toThrow(/under \/app\//);
    expect(() => buildRoutes(table([{ ...route, path: '/b/{id}', site: true }]))).not.toThrow();
    expect(() => buildRoutes(table([{ ...route, site: true }]))).toThrow(/outside \/app\//);
    expect(() => buildRoutes(table([route, { ...route, path: '/app/sims/{id}' }]))).toThrow(
      /two routes open/,
    );
    expect(() => buildRoutes(table([route, { ...route, path: '/app/sim/{x}', screen: 'other' }]))).toThrow(
      /same links/,
    );
  });

  it("keeps a route's excluded pages on the site, in both parsers and the path list", () => {
    const route = {
      path: '/sim/{id}',
      site: true,
      screen: 'simPage',
      example: '/sim/a',
      exclude: ['/sim/gear'],
    };
    const model = buildRoutes(table([route]));
    expect(excludedPaths(model)).toEqual(['/sim/gear']);
    expect(renderSwift(model, '')).toContain('static let siteOnly: Set<String> = ["/sim/gear"]');
    expect(renderKotlin(model, '')).toContain('val SITE_ONLY: Set<String> = setOf("/sim/gear")');
    expect(() => buildRoutes(table([{ ...route, exclude: ['/reports/gear'] }]))).toThrow(
      /not a path this route/,
    );
    expect(() => buildRoutes(table([{ ...route, exclude: ['/sim/{id}'] }]))).toThrow(/not a path this route/);
    expect(() => buildRoutes(table([{ ...route, exclude: ['/sim/a/b'] }]))).toThrow(/not a path this route/);
  });

  it('renames a keyword param in the code and keeps it in the URL', () => {
    expect(codeName('class')).toBe('className');
    expect(codeName('spec')).toBe('spec');
    const model = buildRoutes(
      table([{ path: '/app/guides/{class}/{spec}', screen: 'guide', example: '/app/guides/a/b' }]),
    );
    const kotlin = renderKotlin(model, '') as string;
    const swift = renderSwift(model, '') as string;
    expect(kotlin).toContain('data class Guide(val className: String, val spec: String)');
    expect(kotlin).toContain('link("/app/guides/${enc(className)}/${enc(spec)}")');
    expect(kotlin).not.toMatch(/val class:/);
    expect(swift).toContain('case guide(className: String, spec: String)');
    expect(swift).toContain('Self.link("/app/guides/\\(className)/\\(spec)", [])');
    expect(appLinkPaths(model)).toEqual(['/app/guides/*/*']);
  });
});
