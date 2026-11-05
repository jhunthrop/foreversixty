// web/src/lib/guild/week-one.test.ts
import { describe, expect, it } from 'vitest';
import type { GuildRosterEntry, NextRaidView } from './api';
import { firstKillLine, progressionEmpty } from './progression-view';
import {
  characterSpecKey,
  levelSpread,
  professionFacts,
  raidsOpenDay,
  raidsOpenLine,
  spreadIsEmpty,
  weekOneStep,
  weekOneTitle,
  type WeekOneInput,
} from './week-one';

const BEFORE = new Date('2026-11-05T00:00:00Z');
const AFTER = new Date('2026-12-10T00:00:00Z');

function entry(level: number | undefined, state: GuildRosterEntry['state'] = 'approved'): GuildRosterEntry {
  return { name: 'X', state, rank: 'member', level };
}

function raid(overrides: Partial<NextRaidView> = {}): NextRaidView {
  return {
    id: 7,
    zone: 'deadmines',
    zone_name: 'The Deadmines',
    title: 'Deadmines run',
    starts_at: '2026-11-08T02:00:00Z',
    size: 5,
    guild_timezone: 'America/Chicago',
    counts: { in: 2, late: 1, tanks: 1, healers: 1, dps: 1 },
    own: null,
    ...overrides,
  };
}

function input(overrides: Partial<WeekOneInput> = {}): WeekOneInput {
  return {
    officer: false,
    waiting: 0,
    nextRaid: null,
    raidHref: null,
    character: { name: 'Rootbound', class: 'hunter', level: 52, spec: 'BeastMastery' },
    now: BEFORE,
    locale: 'en-US',
    viewerTz: 'America/Chicago',
    ...overrides,
  };
}

describe('the raids line', () => {
  it('reads phases.json for the date and counts down in the raid words', () => {
    expect(raidsOpenDay('en-US')).toBe('Wed 9 Dec');
    expect(raidsOpenLine(new Date('2026-12-07T00:00:00Z'), 'en-US')).toBe(
      "Barrow Deeps, Hyjal Summit and Onyxia's Lair open Wed 9 Dec, in 2 days.",
    );
    expect(raidsOpenLine(new Date('2026-12-07T20:00:00Z'), 'en-US')).toContain('in 1 day 4 h.');
  });

  it('turns over on 9 Dec, title and line both', () => {
    expect(weekOneTitle(BEFORE)).toBe('Before raids open');
    expect(weekOneTitle(AFTER)).toBe('Your first raid');
    expect(raidsOpenLine(AFTER)).toBe(
      'Raids are open. Your first logged night starts the progression board.',
    );
  });
});

describe('levelSpread', () => {
  it('buckets the roster, leaving out bank alts and unclaimed in-game names', () => {
    const spread = levelSpread([
      entry(60),
      entry(60),
      entry(59),
      entry(50),
      entry(49),
      entry(40),
      entry(39),
      entry(10),
      entry(9),
      entry(undefined),
      entry(60, 'no_account'),
      entry(undefined, 'no_account'),
    ]);
    expect(spread).toEqual({ at60: 2, fifties: 2, forties: 2, under40: 2, noLevel: 1 });
    expect(spreadIsEmpty(spread)).toBe(false);
  });

  it('is empty when nobody counted has a level', () => {
    const spread = levelSpread([entry(undefined), entry(5)]);
    expect(spreadIsEmpty(spread)).toBe(true);
    expect(spread.noLevel).toBe(1);
  });
});

describe('professionFacts', () => {
  it('lists primary professions alphabetically with who has them, top three then +n', () => {
    const facts = professionFacts([
      { name: 'Cindergale', level: 40, professions: ['alchemy', 'cooking'] },
      { name: 'Rootbound', level: 52, professions: ['alchemy', 'first-aid', 'fishing'] },
      { name: 'Ashenveil', level: 60, professions: ['enchanting', 'alchemy'] },
      { name: 'Duskwind', level: 12, professions: ['alchemy'] },
      { name: 'Bankalt', level: 5, professions: ['tailoring'] },
      { name: 'Elder', level: 11, professions: ['alchemy'] },
    ]);
    expect(facts).toEqual(['Alchemy: Ashenveil, Cindergale, Duskwind +2', 'Enchanting: Ashenveil']);
  });

  it('is empty when no counted row carries one', () => {
    expect(professionFacts([{ name: 'A', level: 30 }])).toEqual([]);
  });
});

describe('characterSpecKey', () => {
  it('reads a log spelling, a name and a key', () => {
    expect(characterSpecKey('hunter', 'BeastMastery')).toBe('hunter-beast-mastery');
    expect(characterSpecKey('hunter', 'Beast Mastery')).toBe('hunter-beast-mastery');
    expect(characterSpecKey('hunter', 'hunter-survival')).toBe('hunter-survival');
    expect(characterSpecKey('hunter', null)).toBeNull();
    expect(characterSpecKey('mage', 'Holy')).toBeNull();
  });
});

describe('the officer step', () => {
  const officer = { officer: true };

  it('approvals come first, with the invite link beside', () => {
    const step = weekOneStep(input({ ...officer, waiting: 2, nextRaid: raid(), raidHref: '/r' }));
    expect(step.text).toBe('2 waiting for approval. Approve them so they can sign up.');
    expect(step.action).toEqual({ kind: 'approve', label: 'Approve 2' });
    expect(step.invite).toBe(true);
  });

  it('then an empty calendar', () => {
    const step = weekOneStep(input(officer));
    expect(step.text).toBe('Nothing on the calendar yet. Post a dungeon run or a guild event.');
    expect(step.action).toEqual({ kind: 'plan', label: 'Plan an event' });
  });

  it('then the event planned, counting In and Late', () => {
    const step = weekOneStep(input({ ...officer, nextRaid: raid(), raidHref: '/r/7' }));
    expect(step.text).toBe('Deadmines run Saturday 8 pm, 3 signed up.');
    expect(step.action).toEqual({ kind: 'link', href: '/r/7', label: 'Open the sign-ups' });
  });

  it('never prints 0 signed up', () => {
    const counts = { in: 0, late: 0, tanks: 0, healers: 0, dps: 0 };
    const step = weekOneStep(input({ ...officer, nextRaid: raid({ counts }), raidHref: '/r/7' }));
    expect(step.text).toBe('Deadmines run Saturday 8 pm. Nobody signed up yet.');
  });
});

describe('the raider step', () => {
  it('no export asks for one', () => {
    const step = weekOneStep(input({ character: null }));
    expect(step.text).toBe('Sync the addon or import from Battle.net so the guild sees your level and gear.');
    expect(step.action).toEqual({ kind: 'link', href: '/setup', label: 'Import your character' });
    expect(step.invite).toBe(false);
  });

  it('an event planned and not signed up asks for the sign-up', () => {
    const step = weekOneStep(input({ nextRaid: raid(), raidHref: '/r/7' }));
    expect(step.text).toBe('Deadmines run Saturday 8 pm. You have not signed up yet.');
    expect(step.action).toEqual({ kind: 'link', href: '/r/7', label: 'Sign up' });
  });

  it('signed up, a leveling character reads its level and its guide', () => {
    const own = { character_key: 'k', state: 'in' as const, role: 'dps' as const };
    const step = weekOneStep(input({ nextRaid: raid({ own }), raidHref: '/r/7' }));
    expect(step.name).toBe('Rootbound');
    expect(step.text).toBe(' is level 52, 8 levels to 60.');
    expect(step.action).toEqual({
      kind: 'link',
      href: '/guides/hunter/beast-mastery#leveling',
      label: 'Level as Beast Mastery',
    });
  });

  it('with no spec yet, the class guide says which spec to level as', () => {
    const step = weekOneStep(input({ character: { name: 'A', class: 'warrior', level: 20, spec: null } }));
    expect(step.action.kind === 'link' && step.action.href).toMatch(/^\/guides\/warrior\/\w+#leveling$/);
    expect(step.action.label).toMatch(/^Level as \w+$/);
  });

  it('at 60 reads the gear check and its best in slot', () => {
    const step = weekOneStep(
      input({ character: { name: 'A', class: 'hunter', level: 60, spec: 'Survival' } }),
    );
    expect(step.text).toBe(' is level 60. Check your gear before Wed 9 Dec.');
    expect(step.action).toEqual({
      kind: 'link',
      href: '/bis/hunter/survival',
      label: 'Open your best in slot',
    });
  });
});

describe('progression view', () => {
  it('the zero-pull tab says when, then asks for a log', () => {
    expect(progressionEmpty(BEFORE)).toEqual({
      message:
        "Barrow Deeps, Hyjal Summit and Onyxia's Lair open Wed 9 Dec. Progression starts with your first logged pull.",
      action: { label: 'How to log a raid', href: '/logs' },
    });
    expect(progressionEmpty(AFTER)).toEqual({
      message: 'No pulls logged yet. Upload your first raid night and it shows here.',
      action: { label: 'Upload a log', href: '/logs' },
    });
  });

  it('a kill line needs both figures: an omitted field never prints NaN', () => {
    const now = new Date(2026, 11, 20);
    expect(firstKillLine({ first_kill_at: '2026-12-16T19:00:00Z', best_kill_ms: 225_000 }, now)).toBe(
      'First kill Wed 16 Dec · best kill 3:45',
    );
    expect(firstKillLine({ first_kill_at: '2026-12-16T19:00:00Z' }, now)).toBeNull();
    expect(firstKillLine({ best_kill_ms: 225_000 }, now)).toBeNull();
    expect(firstKillLine({ first_kill_at: null, best_kill_ms: null }, now)).toBeNull();
  });
});
