// web/src/lib/report/shell-paths.test.ts
import { afterEach, beforeEach, describe, expect, it } from 'vitest';
import {
  FIXTURE_GUILD_CLAIM_PATH,
  FIXTURE_GUILD_INVITE_TOKEN,
  FIXTURE_GUILD_PATH,
  FIXTURE_GUILD_SETTINGS_PATH,
  fixtureGuildPaths,
} from './shell-paths';

const ORIGINAL = process.env.FOREVER_DATA;
beforeEach(() => {
  process.env.FOREVER_DATA = 'fixture';
});
afterEach(() => {
  process.env.FOREVER_DATA = ORIGINAL;
});

describe('fixtureGuildPaths', () => {
  it('includes the plain guild path and the three new sub-routes under FOREVER_DATA=fixture', () => {
    const paths = fixtureGuildPaths().map((entry) => entry.params.path);
    expect(paths).toContain(FIXTURE_GUILD_PATH);
    expect(paths).toContain(FIXTURE_GUILD_CLAIM_PATH);
    expect(paths).toContain(FIXTURE_GUILD_SETTINGS_PATH);
    expect(paths).toContain(`invite/${FIXTURE_GUILD_INVITE_TOKEN}`);
  });

  it('returns nothing outside fixture mode', () => {
    process.env.FOREVER_DATA = 'production';
    expect(fixtureGuildPaths()).toEqual([]);
  });
});
