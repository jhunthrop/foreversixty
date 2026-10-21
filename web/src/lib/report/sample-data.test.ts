// web/src/lib/report/sample-data.test.ts
// The bundled sample is a real beta night cut and pseudonymised by api/cmd/samplecut. This pins
// what ships: every player is a fantasy two-word name on the report's own realm under a synthetic
// GUID, no owner or guild travels with it, and the whole directory stays inside the size budget.
// The pseudonymiser itself (deterministic, no original survives) is tested in Go beside the tool.
import { readdirSync, readFileSync, statSync } from 'node:fs';
import path from 'node:path';
import { describe, expect, it } from 'vitest';

const ROOT = path.resolve(__dirname, '../../../public/sample-report/sample');
const MAX_BYTES = 8 * 1024 * 1024;
const PLAYER_GUID = /^Player-4184-[0-9A-F]{8}$/;
const PLAYER_NAME = /^[A-Z][a-z]+-[A-Z][a-z]+-ClassicBetaPvP-$/;

interface Unit {
  guid: string;
  name: string;
  kind: string;
}
interface Actor {
  guid: string;
  name: string;
}

const read = <T>(file: string): T => JSON.parse(readFileSync(path.join(ROOT, file), 'utf8')) as T;

function sizeOf(dir: string): number {
  return readdirSync(dir, { withFileTypes: true }).reduce(
    (sum, entry) =>
      sum +
      (entry.isDirectory() ? sizeOf(path.join(dir, entry.name)) : statSync(path.join(dir, entry.name)).size),
    0,
  );
}

describe('the bundled sample report data', () => {
  const meta = read<{
    title: string;
    owner: unknown;
    guild?: unknown;
    created_at: string;
    zone: string;
    fights: { index: number; players: string[] }[];
    players: string[];
  }>('meta.json');
  const report = read<{ units: Unit[]; fights: { index: number }[] }>('report.json');

  it('is the beta night, dated from its own log, with no owner or guild', () => {
    expect(meta.title).toBe('Beta night, 17 Sept 2026');
    expect(meta.created_at.startsWith('2026-09-17')).toBe(true);
    expect(meta.owner).toBeNull();
    expect(meta.guild).toBeUndefined();
    expect(meta.fights.length).toBeGreaterThan(30);
    expect(meta.fights.every((fight) => fight.players.length > 0)).toBe(true);
    expect(report.fights.map((fight) => fight.index)).toEqual(meta.fights.map((fight) => fight.index));
  });

  it('names every player with a pseudonym and a synthetic GUID', () => {
    const players = report.units.filter((unit) => unit.kind === 'player');
    expect(players.length).toBeGreaterThan(100);
    for (const unit of players) {
      expect(unit.guid).toMatch(PLAYER_GUID);
      expect(unit.name).toMatch(PLAYER_NAME);
    }
    expect(new Set(players.map((unit) => unit.name)).size).toBe(players.length);
    expect(meta.players.every((guid) => PLAYER_GUID.test(guid))).toBe(true);
  });

  it('carries no real player name into any fight summary', () => {
    for (const fight of meta.fights) {
      const summary = read<{ damage_done: Actor[]; healing: Actor[] }>(`fights/${fight.index}/summary.json`);
      for (const actor of [...summary.damage_done, ...summary.healing]) {
        if (actor.guid.startsWith('Player-')) {
          expect(actor.guid).toMatch(PLAYER_GUID);
          expect(actor.name).toMatch(PLAYER_NAME);
        }
      }
    }
  });

  it('stays under the static asset budget', () => {
    expect(sizeOf(ROOT)).toBeLessThan(MAX_BYTES);
  });
});

describe('the sample classes, read from the log', () => {
  const report = read<{
    units: { guid: string; kind: string; class?: string; class_source?: string; spec?: string }[];
  }>('report.json');
  const players = report.units.filter((unit) => PLAYER_GUID.test(unit.guid));

  it('gives most players a class, read from the spells they cast', () => {
    const classed = players.filter((unit) => unit.class);
    expect(classed.length / players.length).toBeGreaterThan(0.85);
    expect(classed.every((unit) => unit.class_source === 'inferred')).toBe(true);
    expect(new Set(classed.map((unit) => unit.class)).size).toBeGreaterThanOrEqual(5);
  });

  it('has no spec to show, so the Spec column hides (no talent spell was cast)', () => {
    expect(players.some((unit) => unit.spec)).toBe(false);
  });

  it('shows a class on the opening fight roster rows', () => {
    const summary = read<{ roster: { class?: string }[] }>('fights/6/summary.json');
    expect(summary.roster.filter((row) => row.class).length / summary.roster.length).toBeGreaterThan(0.85);
  });
});
