// web/src/lib/report/shell-paths.test.ts
import { afterEach, beforeEach, describe, expect, it } from 'vitest';
import {
  FIXTURE_GUILD_CLAIM_PATH,
  FIXTURE_GUILD_INVITE_TOKEN,
  FIXTURE_GUILD_PATH,
  FIXTURE_GUILD_SETTINGS_PATH,
  LOCAL_GUILD_SHELL_PATH,
  fixtureGuildPaths,
} from './shell-paths';

const ORIGINAL = process.env.FOREVER_DATA;
const ORIGINAL_SHELL = process.env.FS_LOCAL_GUILD_SHELL;
beforeEach(() => {
  process.env.FOREVER_DATA = 'fixture';
  delete process.env.FS_LOCAL_GUILD_SHELL;
});
afterEach(() => {
  process.env.FOREVER_DATA = ORIGINAL;
  if (ORIGINAL_SHELL === undefined) delete process.env.FS_LOCAL_GUILD_SHELL;
  else process.env.FS_LOCAL_GUILD_SHELL = ORIGINAL_SHELL;
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

  it('prerenders the one local-stack shell path, outside fixture mode, only when asked', () => {
    process.env.FOREVER_DATA = 'production';
    process.env.FS_LOCAL_GUILD_SHELL = '1';
    expect(fixtureGuildPaths()).toEqual([{ params: { path: LOCAL_GUILD_SHELL_PATH } }]);
  });
});
