// web/src/test-support/sim-api.ts
// The API does not run in this repository's tests, so every route the simulator contract
// defines is answered here instead.
//
// This is the idiom the repository already has -- vi.stubGlobal('fetch', …) returning a
// real Response, as src/lib/account/api.test.ts and src/lib/rankings/api.test.ts do -- with
// the route table factored out so a route that changes shape changes in one place. An
// unmatched request throws rather than falling through to the network: a test that calls a
// route nobody wrote a handler for must fail loudly, not hang.
//
// lastBody(), lastUrl() and lastHeaders() exist because the assertions that matter are
// often about what the client *sent*: that the character travelled, that a character key
// was path-escaped rather than encoded, that the CSRF header was set, that the engine
// version travelled. Recording them here keeps every test from re-wrapping fetch.
import { readFileSync } from 'node:fs';
import path from 'node:path';
import { vi } from 'vitest';
import activeBuild from '../data/active-build.json';
import fixtureResultJson from '../fixtures/sim/result.json';
import fixtureSpecsJson from '../fixtures/sim/specs.json';
import { FIXTURE_SIM_ID } from '../lib/report/shell-paths';
import type { SimResult, SpecFidelity } from '../lib/sim/types';

const WEB_ROOT = path.resolve(import.meta.dirname, '..', '..');

/**
 * The build `FOREVER_DATA=fixture` publishes to public/data/. scripts/sync-data.mjs names
 * the published directory after src/data/active-build.json's own `build` field even under
 * FOREVER_DATA=fixture -- only the *content* underneath it comes from src/fixtures/planner
 * (whose own JSON files carry the build id they were captured against, '1.15.9.69722', which
 * is not the directory name once the planner has moved to a later active build -- see
 * src/lib/sim/character.test.ts's own note on the same fact). Reading the file rather than
 * repeating the id keeps this from going stale the next time active-build.json moves, the
 * way src/pages/_planner.test.ts already reads it for the same reason.
 */
export const FIXTURE_DATA_BUILD: string = activeBuild.build;
export const FIXTURE_BUILD_ID = 'bld123456789';

/** The planner's own data files, served from public/ the way the site serves them. */
function dataFile(relative: string): Response {
  try {
    const body = readFileSync(path.join(WEB_ROOT, 'public/data', FIXTURE_DATA_BUILD, relative), 'utf8');
    return new Response(body, { status: 200, headers: { 'content-type': 'application/json' } });
  } catch {
    return new Response(null, { status: 404 });
  }
}

// Twelve characters of [a-z2-7], because that is what a sim_id is -- the same alphabet
// and length as a report_id. It is not decoration: Task 22's Lighthouse entry for the
// saved-sim page matches on /sim/[a-z2-7]{12}\.html, and an id with a digit outside the
// alphabet silently falls into the catch-all at the wrong budget instead.
//
// Re-exported rather than declared here: src/lib/report/shell-paths.ts is what
// src/pages/sim/[id].astro prerenders against, so it is the one literal, and this module
// (Node-only, like shell-paths.ts) just carries the name testers already import from here.
export { FIXTURE_SIM_ID };
export const NEW_SIM_ID = 'simnew234567';
export const TEST_API = 'https://api.test';

export const fixtureResult = fixtureResultJson as unknown as SimResult;
export const fixtureSpecs = fixtureSpecsJson as unknown as SpecFidelity[];

/** One route: the method, a pattern over the path, and what it answers. */
export interface StubRoute {
  method: string;
  pattern: RegExp;
  respond: (match: RegExpExecArray, request: Request) => Response | Promise<Response>;
}

export interface SimApiStub {
  /** Installs the fetch stub. Call from `beforeEach`. */
  install(): void;
  /** Removes it and clears every recording. Call from `afterEach`. */
  reset(): void;
  /** Flips the premium flag POST /v1/sims/run checks. Reset by `reset()`. */
  setPremium(value: boolean): void;
  /** Adds a route ahead of the built-in ones, for a test that needs a different answer. */
  route(route: StubRoute): void;
  lastUrl(): string;
  lastBody(): unknown;
  lastHeaders(): Headers | null;
}

/** The Phase 0 envelope every route of ours answers in. */
export function envelope(data: unknown, status = 200): Response {
  return new Response(JSON.stringify({ ok: status < 400, data, error: null, request_id: 'req-test' }), {
    status,
    headers: { 'content-type': 'application/json' },
  });
}

export function failure(message: string, status: number): Response {
  return new Response(JSON.stringify({ ok: false, data: null, error: { message }, request_id: 'req-test' }), {
    status,
    headers: { 'content-type': 'application/json' },
  });
}

export function createSimApi(): SimApiStub {
  let premium = true;
  let lastUrl = '';
  let lastBody: unknown = null;
  let lastHeaders: Headers | null = null;
  let extra: StubRoute[] = [];

  const builtIn: StubRoute[] = [
    {
      method: 'POST',
      pattern: /\/v1\/sims$/,
      respond: () => envelope({ sim_id: NEW_SIM_ID }, 201),
    },
    {
      method: 'POST',
      pattern: /\/v1\/sims\/run$/,
      respond: () => (premium ? envelope({ sim_id: NEW_SIM_ID }, 202) : failure('premium required', 402)),
    },
    {
      method: 'GET',
      pattern: /\/v1\/sims\/([a-z2-7]{12})\/progress$/,
      respond: () => envelope({ state: 'running', iterations_done: 4200, dps: 1559.7 }),
    },
    {
      method: 'GET',
      pattern: /\/v1\/sims\/([a-z2-7]{12})$/,
      respond: (match) =>
        match[1] === FIXTURE_SIM_ID ? envelope(fixtureResult) : failure('no such sim', 404),
    },
    {
      method: 'GET',
      pattern: /\/v1\/sims(\?|$)/,
      respond: (_match, request) => {
        const wanted = new URL(request.url).searchParams.get('kind');
        const rows = [
          {
            sim_id: FIXTURE_SIM_ID,
            spec: 'warrior-fury',
            dps: fixtureResult.dps.mean,
            engine_version: fixtureResult.engine_version,
            created_at: '2026-09-14T10:02:00Z',
            title: 'Raid-buffed, 3:00, single target',
            kind: 'run',
            headline: '1,204 DPS',
          },
          {
            sim_id: 'gearaaaaaaaa',
            spec: 'warrior-fury',
            dps: 1245,
            engine_version: fixtureResult.engine_version,
            created_at: '2026-09-15T10:02:00Z',
            title: '',
            kind: 'gear',
            headline: '+41 DPS from Vis’kag',
          },
          {
            sim_id: 'dropsaaaaaaa',
            spec: 'warrior-fury',
            dps: 1210,
            engine_version: fixtureResult.engine_version,
            created_at: '2026-09-16T10:02:00Z',
            title: '',
            kind: 'drops',
            headline: '3 upgrades on Ragnaros',
          },
          {
            sim_id: 'talentsaaaaa',
            spec: 'warrior-fury',
            dps: 1222,
            engine_version: fixtureResult.engine_version,
            created_at: '2026-09-17T10:02:00Z',
            title: '',
            kind: 'talents',
            headline: '+18 DPS with ‘Deep Fury’',
          },
          {
            sim_id: 'weightsaaaaa',
            spec: 'warrior-fury',
            dps: 0,
            engine_version: fixtureResult.engine_version,
            created_at: '2026-09-18T10:02:00Z',
            title: '',
            kind: 'weights',
            headline: 'Crit 1.00 · Agility 0.87',
          },
        ];
        const filtered = wanted === null ? rows : rows.filter((row) => row.kind === wanted);
        return envelope({ rows: filtered, total: filtered.length, page: 1, per_page: 100 });
      },
    },
    {
      method: 'GET',
      pattern: /\/v1\/specs$/,
      respond: () => envelope({ specs: fixtureSpecs }),
    },
    // Three path segments, not one key: the contract spells this route the way the existing
    // character route is spelled. `source` is "addon" or "fight" -- Armory is not a source
    // yet. Shape matches `api/internal/sims/input_test.go`'s own fixtures, not a guess:
    // `gear` is the source's opaque JSON (`{"slots": […]}` there, for an addon read) and
    // `talents` is `fight_metrics.talent_split`, "31/0/20". No `race`: the real API never
    // sends one (final whole-branch review, H3) -- fromStoredCharacter always refuses on
    // this default route the same way it does against the real API today, and the one test
    // that needs a successful read adds its own `race` to exercise "the day the API starts
    // sending it".
    {
      method: 'GET',
      pattern: /\/v1\/characters\/[^/]+\/[^/]+\/[^/]+\/sim-input$/,
      respond: () =>
        envelope({
          spec: 'warrior-fury',
          gear: { slots: [12640, 11726] },
          talents: '31/0/20',
          // IDS.md ids, not spell ids: the API does the mapping (amended contract).
          buffs: ['battle_shout', 'blessing_of_kings'],
          captured_at: '2026-09-14T09:40:00Z',
          source: 'addon',
        }),
    },
    // fromPlannerBuild's GET /v1/builds/{id} read (Task 7). The class and race ids are the
    // fixture's own -- warrior is 1 and orc is 2 in src/fixtures/planner/classes.json and
    // races.json -- so a test reads them from those files rather than repeating the numbers.
    {
      method: 'GET',
      pattern: /\/v1\/builds\/([A-Za-z0-9]+)$/,
      respond: (match) =>
        match[1] !== FIXTURE_BUILD_ID
          ? failure('no such build', 404)
          : envelope({
              id: FIXTURE_BUILD_ID,
              class_id: 1,
              race_id: 2,
              tree_version: FIXTURE_DATA_BUILD,
              point_order: [2001, 2001, 2001, 2001, 2001, 2002, 2002, 2002, 2002, 2002, 2003, 2003, 2003],
              gear: { head: 12640 },
              title: 'Fury opener',
              created_at: '2026-09-13T20:00:00Z',
              views: 4,
            }),
    },
    // fromPlannerBuild and fromAddonExport (Task 7) both read /data/<build>/… files with
    // lib/planner/load.ts's own loaders. Last and broadest on purpose: its capture is the
    // whole remaining path, so talents/warrior.json and classes.json are both served by
    // this one route rather than by two that could disagree, and it must stay ordered
    // after every /v1/… route above or it would swallow them too.
    {
      method: 'GET',
      pattern: /\/data\/[^/]+\/(.+)$/,
      respond: (match) => dataFile(match[1]),
    },
  ];

  async function handle(input: RequestInfo | URL, init?: RequestInit): Promise<Response> {
    // Every /v1/… caller builds an absolute url itself (requestEnvelope, report/load.ts's
    // apiGet). Task 7's /data/<build>/… reads do not: the planner's own fetchJson(url) calls
    // fetch() with dataUrl()'s root-relative path exactly as the browser would, resolving it
    // against the page's own origin. Node's fetch has no page to resolve against, so a bare
    // relative string is resolved against a fixed placeholder origin first -- routing below
    // reads only the resulting pathname and search, so the origin itself is never observed.
    const request =
      input instanceof Request ? input : new Request(new URL(String(input), 'https://fixture.test'), init);
    lastUrl = request.url;
    lastHeaders = new Headers(request.headers);
    lastBody = null;
    if (request.method !== 'GET') {
      try {
        lastBody = await request.clone().json();
      } catch {
        lastBody = null;
      }
    }
    const path = new URL(request.url).pathname + new URL(request.url).search;
    for (const route of [...extra, ...builtIn]) {
      if (route.method !== request.method) continue;
      const match = route.pattern.exec(path);
      if (match !== null) return route.respond(match, request);
    }
    throw new Error(`unhandled request: ${request.method} ${request.url}`);
  }

  return {
    install() {
      vi.stubGlobal('fetch', vi.fn(handle));
    },
    reset() {
      vi.unstubAllGlobals();
      premium = true;
      lastUrl = '';
      lastBody = null;
      lastHeaders = null;
      extra = [];
      document.cookie = 'fs_csrf=; Max-Age=0; path=/';
    },
    setPremium(value) {
      premium = value;
    },
    route(route) {
      extra = [route, ...extra];
    },
    lastUrl: () => lastUrl,
    lastBody: () => lastBody,
    lastHeaders: () => lastHeaders,
  };
}
