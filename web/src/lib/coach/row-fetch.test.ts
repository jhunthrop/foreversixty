import { describe, expect, it } from 'vitest';
import type { Me } from '../account/me-types';
import type { FightEntry } from '../report/types';
import { lineRequestFor, officerButtonFor } from './row-fetch';
import { SIGNED_OUT_VIEWER, viewerFromMe } from './viewer';

const boss = {
  index: 3,
  kind: 'encounter',
  duration_ms: 200_000,
} as FightEntry;
const report = { reportId: 'abc123', fight: boss, firstFight: 1 };
const owner = viewerFromMe({
  user: { id: 1 },
  characters: [{ key: 'us/normal/bow-jackzon', name: 'Bow Jackzon' }],
  guilds: [],
  entitlements: { log_coach: true, officer_views: true },
} as unknown as Me);

describe('lineRequestFor', () => {
  it("asks for the viewer's own character on a reviewable pull", () => {
    expect(lineRequestFor(owner, 'Bow-Jackzon-ClassicBetaPvP-', report)).toEqual({
      reportId: 'abc123',
      fightIndex: 3,
      playerKey: 'us/normal/bow-jackzon',
    });
  });
  it('asks for nobody else, the signed-out, trash or the whole night', () => {
    expect(lineRequestFor(owner, 'Simfury', report)).toBeNull();
    expect(lineRequestFor(SIGNED_OUT_VIEWER, 'Bow Jackzon', report)).toBeNull();
    expect(
      lineRequestFor(owner, 'Bow Jackzon', { ...report, fight: { ...boss, kind: 'trash' } as FightEntry }),
    ).toBeNull();
    expect(lineRequestFor(owner, 'Bow Jackzon', { ...report, fight: null })).toBeNull();
  });
});

describe('officerButtonFor', () => {
  it("shows on another player's row for an officer only", () => {
    expect(officerButtonFor(owner, 'Simfury', report)).toBe(true);
    expect(officerButtonFor(owner, 'Bow Jackzon', report)).toBe(false);
    expect(officerButtonFor(SIGNED_OUT_VIEWER, 'Simfury', report)).toBe(false);
  });
});
