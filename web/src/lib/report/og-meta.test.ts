import { describe, expect, it } from 'vitest';
import fixtureMeta from '../../fixtures/report/meta.json';
import { fixtureBulkResult, fixtureResult, fixtureWeightsResult } from '../../test-support/sim-api';
import type { ReportMeta } from './types';
import {
  characterShellMeta,
  guildShellMeta,
  rankingsShellMeta,
  reportShellMeta,
  simShellMeta,
  titleize,
} from './og-meta';

const API = 'https://api.foreversixty.test';
const meta = fixtureMeta as ReportMeta;

describe('shell unfurl values', () => {
  it('titles a report and describes its night in plain numbers', () => {
    const shell = reportShellMeta(meta, API);
    expect(shell.title).toBe('Sanguine Depths, fixture night · Forever Sixty');
    expect(shell.description).toBe('4 fights in Sanguine Depths, 2 boss kills, logged 2026-09-26.');
    expect(shell.image).toBe(`${API}/reports/fixture2abcd/card.png`);
    expect(shell.canonical).toBe('https://foreversixty.gg/reports/fixture2abcd');
  });

  it('says one fight and no boss kills without reading oddly', () => {
    const single: ReportMeta = {
      ...meta,
      title: 'One pull',
      fights: [{ ...meta.fights[0], kill: false }],
    };
    expect(reportShellMeta(single, API).description).toBe(
      '1 fight in Sanguine Depths, no boss kills, logged 2026-09-26.',
    );
  });

  it('falls back to the zone when a report has no title', () => {
    expect(reportShellMeta({ ...meta, title: '' }, API).title).toBe('Sanguine Depths · Forever Sixty');
  });

  it('titleizes an encounter slug', () => {
    expect(titleize('warden-kelthas')).toBe('Warden Kelthas');
    expect(titleize('the-drowned-city')).toBe('The Drowned City');
  });

  it('describes a rankings page from its totals', () => {
    const shell = rankingsShellMeta('warden-kelthas', { total: 812, updated_at: '2026-12-09T21:04:00Z' });
    expect(shell.title).toBe('Warden Kelthas rankings · Forever Sixty');
    expect(shell.description).toBe('812 ranked kills of Warden Kelthas, updated 2026-12-09.');
    expect(shell.canonical).toBe('https://foreversixty.gg/rankings/warden-kelthas');
    expect(shell.image).toBe('https://foreversixty.gg/og/index.png');
  });

  it('describes a character by ruleset and region', () => {
    const shell = characterShellMeta(
      { region: 'us', ruleset: 'hardcore', slug: 'elyra-duskvale' },
      {
        character: { name: 'Elyra Duskvale', region: 'us', ruleset: 'hardcore', class: 'Priest' },
        best: [{}, {}],
        history: [{}, {}, {}],
      },
    );
    expect(shell.title).toBe('Elyra Duskvale · Hardcore US · Forever Sixty');
    expect(shell.description).toBe('3 ranked fights across 2 encounters.');
    expect(shell.canonical).toBe('https://foreversixty.gg/character/us/hardcore/elyra-duskvale');
  });

  it('describes a guild by its progression', () => {
    const shell = guildShellMeta(
      { region: 'eu', ruleset: 'normal', slug: 'the-last-watch' },
      {
        guild: { name: 'The Last Watch', region: 'eu', ruleset: 'normal' },
        progression: [
          { kills: 3, pull_count: 40 },
          { kills: 0, pull_count: 12 },
        ],
      },
    );
    expect(shell.title).toBe('The Last Watch · Normal EU · Forever Sixty');
    // Only the first row has been killed at all (kills: 3 > 0); the second (kills: 0) has
    // not. "Bosses down" counts distinct bosses defeated, not total kill events, so this is
    // one boss down, not three -- see the next test for the case that distinguishes them.
    expect(shell.description).toBe('1 boss down over 52 pulls.');
    expect(shell.canonical).toBe('https://foreversixty.gg/guild/eu/normal/the-last-watch');
  });

  it('counts a repeatedly-farmed boss once, not by its kill count', () => {
    const shell = guildShellMeta(
      { region: 'eu', ruleset: 'normal', slug: 'the-last-watch' },
      {
        guild: { name: 'The Last Watch', region: 'eu', ruleset: 'normal' },
        progression: [{ kills: 3, pull_count: 12 }],
      },
    );
    expect(shell.description).toBe('1 boss down over 12 pulls.');
  });
});

describe('simShellMeta', () => {
  const result = fixtureResult;

  // The fixture's own numbers, not the plan draft's: dps.mean is 1038.660353207436, which
  // rounds to 1,039, and dps.error is 2.4, so the 95% band is
  // round(1.96 * 2.4) = round(4.704) = 5.
  it('names the spec and the figure in the title, and the run in the description', () => {
    const meta = simShellMeta(result);
    expect(meta.title).toBe('Fury Warrior, 1,039 DPS · Forever Sixty');
    expect(meta.description).toBe(
      `Simulated on engine ${fixtureResult.engine_version}: 1,039 DPS ± 5 over 3,000 iterations, raid-buffed, 3:00, single target.`,
    );
    expect(meta.canonical).toBe('https://foreversixty.gg/sim/simfixtureab');
    expect(meta.image).toBe('https://foreversixty.gg/og/index.png');
  });

  it('uses the spec slug itself when the spec list does not know it', () => {
    const meta = simShellMeta({ ...result, request: { ...result.request, spec: 'warrior-gladiator' } });
    expect(meta.title).toBe('warrior-gladiator, 1,039 DPS · Forever Sixty');
  });

  it('counts the iterations the run actually asked for', () => {
    const meta = simShellMeta({ ...result, request: { ...result.request, iterations: 10000 } });
    expect(meta.description).toContain('10,000 iterations');
  });

  it('says solo rather than raid-buffed when the summary records no buff on the player', () => {
    const meta = simShellMeta({ ...result, summary: { ...result.summary, auras: [] } });
    expect(meta.description).toContain('solo, 3:00, single target');
    expect(meta.description).not.toContain('raid-buffed');
  });
});

describe('simShellMeta by kind', () => {
  it('names Top Gear and its headline, item included', () => {
    const meta = simShellMeta({ ...fixtureBulkResult, sim_id: 'simfixtureab' });
    expect(meta.title).toBe('Top Gear · Fury Warrior · Forever Sixty');
    // The fixture's own combo count, not a typed-out literal: Task 9's report already
    // documents this same fixture carrying 7 combos (a 7th was added by a fix round after
    // the plan's own draft was written against a 6-combo fixture), and hardcoding either
    // number here would drift the moment the fixture changes again.
    expect(meta.description).toContain(`${fixtureBulkResult.combos.length} combinations`);
    expect(meta.description).toContain('+41 DPS from Helm of Wrath');
  });

  it('says "no combinations" for an empty result, as the API’s headline rule does', () => {
    const meta = simShellMeta({
      ...fixtureBulkResult,
      sim_id: 'simfixtureab',
      combos: [],
    });
    expect(meta.description).toContain('no combinations');
  });

  it('names stat weights and lists the top three', () => {
    const meta = simShellMeta({ ...fixtureWeightsResult, sim_id: 'simfixtureab' });
    expect(meta.title).toBe('Stat weights · Fury Warrior · Forever Sixty');
    expect(meta.description).toContain('Attack power 1.00');
  });

  it('leaves a plain run exactly as it was', () => {
    const meta = simShellMeta(fixtureResult);
    expect(meta.title).toContain('DPS · Forever Sixty');
  });
});
