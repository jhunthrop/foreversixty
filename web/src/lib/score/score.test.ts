// web/src/lib/score/score.test.ts
import { describe, expect, it } from 'vitest';
import labels from '../../../../docs/contracts/fixtures/score-labels.json';
import type { MeCharacter } from '../account/api';
import { scoreCopy } from './copy';
import {
  COMPONENT_MAX,
  TIER_FROM,
  foreverOf,
  hasScore,
  lastAddonExportAt,
  lastSeenText,
  shortDate,
  tierColor,
  tierLabel,
  SCORE_MAX,
  mostToGain,
  ownCharacter,
  tierStrip,
  tierTextColor,
  weakestPart,
} from './score';

describe('shortDate', () => {
  it('reads YYYY-MM-DD as "4 Nov", the addon’s Compat.shortDate', () => {
    expect(shortDate('2026-11-04')).toBe('4 Nov');
    expect(shortDate('2026-09-30')).toBe('30 Sep');
  });

  it('is null for anything else', () => {
    expect(shortDate('')).toBeNull();
    expect(shortDate(null)).toBeNull();
    expect(shortDate('2026-13-01')).toBeNull();
    expect(shortDate('4 Nov')).toBeNull();
  });
});

describe('scoreCopy', () => {
  it('formats the line and the date with the fixture’s own strings', () => {
    expect(labels.line).toBe('Forever Score %d');
    expect(scoreCopy.line(1830)).toBe('Forever Score 1830');
    expect(scoreCopy.asOf('4 Nov')).toBe('as of 4 Nov');
    expect(scoreCopy.noLogs).toBe(labels.noLogs);
  });

  it('labels the five components in the addon’s order', () => {
    expect(scoreCopy.components.map((c) => c.key)).toEqual([
      'gear',
      'experience',
      'parses',
      'attendance',
      'discipline',
    ]);
    expect(scoreCopy.components.map((c) => c.label)).toEqual(labels.components.map((c) => c.label));
  });
});

describe('tiers and ceilings', () => {
  it('carries the formula’s ceilings (api/internal/rating/score_v2.go)', () => {
    expect(COMPONENT_MAX).toEqual({
      gear: 500,
      experience: 700,
      parses: 600,
      attendance: 200,
      discipline: 100,
    });
  });

  it('carries the tier thresholds (score_v2.go’s *From constants)', () => {
    expect(TIER_FROM).toEqual({ common: 0, uncommon: 500, rare: 1000, epic: 1500, legendary: 2000 });
  });

  it('names and colours a known tier, and neither for an unknown one', () => {
    expect(tierColor('epic')).toBe('#a335ee');
    expect(tierLabel('epic')).toBe('Epic');
    expect(tierColor('mythic')).toBeNull();
    expect(tierLabel(null)).toBeNull();
  });
});

describe('foreverOf', () => {
  it('reads the five score fields off a character page', () => {
    const score = foreverOf({
      score: 1830,
      tier: 'epic',
      as_of: '2026-11-04',
      components: { gear: 412, experience: 640, parses: 510, attendance: 180, discipline: 88 },
      experience: { dungeons: 4, raid_kills: 12, alts: 1 },
    });
    expect(score.score).toBe(1830);
    expect(hasScore(score)).toBe(true);
  });

  it('a page that predates the fields, or carries nulls, has no score: never a zero', () => {
    expect(hasScore(foreverOf({}))).toBe(false);
    expect(
      hasScore(foreverOf({ score: null, tier: null, as_of: null, components: null, experience: null })),
    ).toBe(false);
    expect(hasScore(null)).toBe(false);
  });
});

const CHARACTER: MeCharacter = {
  key: 'us/hardcore/elyra-duskvale',
  region: 'us',
  ruleset: 'hardcore',
  name: 'Elyra Duskvale',
  build: { source: 'addon', captured_at: '2026-11-04T14:03:00Z' },
};
const PATH = { region: 'us' as const, ruleset: 'hardcore' as const, slug: 'elyra-duskvale' };

describe('lastAddonExportAt', () => {
  it('is the captured time of the viewer’s own character’s addon export', () => {
    expect(lastAddonExportAt([CHARACTER], PATH)).toBe('2026-11-04T14:03:00Z');
  });

  it('is null for a character the viewer does not own, a Battle.net-only build, or no build', () => {
    expect(lastAddonExportAt([CHARACTER], { ...PATH, slug: 'someone-else' })).toBeNull();
    expect(lastAddonExportAt([CHARACTER], { ...PATH, ruleset: 'pvp' })).toBeNull();
    expect(
      lastAddonExportAt(
        [{ ...CHARACTER, build: { source: 'blizzard', captured_at: '2026-11-04T14:03:00Z' } }],
        PATH,
      ),
    ).toBeNull();
    expect(lastAddonExportAt([{ ...CHARACTER, build: undefined }], PATH)).toBeNull();
    expect(lastAddonExportAt(null, PATH)).toBeNull();
  });
});

describe('lastSeenText', () => {
  it('reads "Last seen in game 4 Nov 14:03" in the viewer’s own clock', () => {
    const local = new Date(2026, 10, 4, 14, 3).toISOString();
    expect(lastSeenText(local)).toBe('Last seen in game 4 Nov 14:03');
  });

  it('is null for a time it cannot read', () => {
    expect(lastSeenText('not a time')).toBeNull();
  });
});

// Fix round 1.
const CAPTURE = { gear: 412, experience: 620, parses: 540, attendance: 180, discipline: 78 };

describe('tier colours (fix round 1 item 1)', () => {
  it('common is white, the client’s quality 1; grey is Poor and is not a tier', () => {
    expect(labels.tiers.common.color).toBe('#ffffff');
    expect(tierColor('common')).toBe('#ffffff');
  });

  it('words at body size take the lightened rare and epic variants; the rest are the exact colour', () => {
    expect(tierTextColor('rare')).toBe('#3d94f0');
    expect(tierTextColor('epic')).toBe('#b866f5');
    expect(tierTextColor('common')).toBe('#ffffff');
    expect(tierTextColor('uncommon')).toBe('#1eff00');
    expect(tierTextColor('legendary')).toBe('#ff8000');
    expect(tierTextColor('mythic')).toBeNull();
  });

  it('the strip keeps the exact quality colour for each bar and the text variant for each name', () => {
    const strip = tierStrip();
    expect(strip.map((s) => s.tier)).toEqual(['common', 'uncommon', 'rare', 'epic', 'legendary']);
    expect(strip.map((s) => s.range)).toEqual(['0–499', '500–999', '1000–1499', '1500–1999', '2000+']);
    const epic = strip.find((s) => s.tier === 'epic');
    expect(epic?.color).toBe('#a335ee');
    expect(epic?.textColor).toBe('#b866f5');
  });

  it('says the note is true now', () => {
    expect(scoreCopy.tiersNote).toBe(
      'The tiers take the colours of item quality in game. The ranges may change as more players are logged.',
    );
  });
});

describe('mostToGain (fix round 1 item 3)', () => {
  it('is the part with the largest gap to its ceiling: Gear, 88 below its top', () => {
    expect(mostToGain(CAPTURE)).toEqual({ key: 'gear', gap: 88 });
    expect(scoreCopy.mostToGainRest(88)).toBe(', 88 below its top.');
  });

  it('weighs absolute points, not shares: 60 of 600 left beats 22 of 100', () => {
    expect(mostToGain({ ...CAPTURE, gear: 500, experience: 700 })).toEqual({ key: 'parses', gap: 60 });
  });

  it('takes the earlier part on a tie, and is null when every part is full', () => {
    expect(mostToGain({ gear: 400, experience: 600, parses: 600, attendance: 200, discipline: 100 })).toEqual(
      {
        key: 'gear',
        gap: 100,
      },
    );
    expect(mostToGain(COMPONENT_MAX)).toBeNull();
    expect(mostToGain(null)).toBeNull();
  });
});

describe('weakestPart (fix round 1 item 3)', () => {
  it('is the smallest share of its own ceiling: Discipline at 78 of 100', () => {
    expect(weakestPart(CAPTURE)).toBe('discipline');
    expect(weakestPart(null)).toBeNull();
  });
});

describe('the card’s words (fix round 1 items 3 and 4)', () => {
  it('reads the most a score can be: 2100', () => {
    expect(SCORE_MAX).toBe(2100);
    expect(scoreCopy.outOf(SCORE_MAX)).toBe('of 2100');
  });

  it('breaks Experience down with singulars where they belong', () => {
    expect(scoreCopy.experienceBreakdown(14, 0, 2)).toBe('14 dungeons · 0 raid bosses · 2 alts');
    expect(scoreCopy.experienceBreakdown(1, 1, 1)).toBe('1 dungeon · 1 raid boss · 1 alt');
  });

  it('captions the card with its date', () => {
    expect(scoreCopy.caption('4 Nov')).toBe('Your whole record, as of 4 Nov.');
    expect(scoreCopy.caption(null)).toBe('Your whole record.');
  });
});

describe('ownCharacter', () => {
  it('finds the viewer’s own row for the page’s character, else null', () => {
    expect(ownCharacter([CHARACTER], PATH)?.key).toBe('us/hardcore/elyra-duskvale');
    expect(ownCharacter([CHARACTER], { ...PATH, slug: 'someone-else' })).toBeNull();
    expect(ownCharacter(null, PATH)).toBeNull();
  });
});
