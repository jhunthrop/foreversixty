import { describe, expect, it } from 'vitest';
import fixtureMeta from '../../fixtures/report/meta.json';
import type { ReportMeta } from './types';
import { characterShellMeta, guildShellMeta, rankingsShellMeta, reportShellMeta, titleize } from './og-meta';

const API = 'https://api.foreversixty.test';
const meta = fixtureMeta as ReportMeta;

describe('shell unfurl values', () => {
  it('titles a report and describes its night in plain numbers', () => {
    const shell = reportShellMeta(meta, API);
    expect(shell.title).toBe('Sanguine Depths, fixture night · Forever Sixty');
    expect(shell.description).toBe('3 fights in Sanguine Depths, 1 boss kill, logged 2026-09-26.');
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
    expect(shell.description).toBe('3 bosses down over 52 pulls.');
    expect(shell.canonical).toBe('https://foreversixty.gg/guild/eu/normal/the-last-watch');
  });
});
