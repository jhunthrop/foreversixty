// web/src/lib/guides/key-talents.test.ts
import { describe, expect, it } from 'vitest';
import type { DecodedTalentView, DecodedTreeView } from './build-tree';
import { keyTalentsFor, parseTalentBullets, withoutTalentBullets } from './key-talents';

function talent(id: number, name: string, tier: number, rank: number, description = ''): DecodedTalentView {
  return {
    id,
    name,
    icon: `icon_${id}`,
    max_rank: 5,
    tier,
    column: 0,
    prereq_talent_id: null,
    prereq_rank: null,
    ranks: Array.from({ length: 5 }, () => ({ spell_id: id, description })),
    spell_id: id,
    rank,
  };
}

const TREE: DecodedTreeView = {
  name: 'Frost',
  position: 1,
  talents: [
    talent(1, 'Improved Frostbolt', 0, 5, 'Cuts the cast time.'),
    talent(2, 'Ice Shards', 4, 5, 'More crit damage.'),
    talent(3, "Winter's Chill", 5, 5, 'Stacks a debuff.'),
    talent(4, 'Unspent', 6, 0, 'Never shown.'),
  ],
  points: 15,
};

const MARKDOWN = `Intro sentence.

- **Improved Frostbolt** (5 points) — up to 0.5 seconds off the cast.
- **Ice Shards** and **Winter's Chill** — the biggest multipliers.
- **Feral Instinct**, **Savage Fury** and **Rend and Tear** — more damage from Swipe.
- Not a talent bullet at all.

Closing paragraph.`;

describe('parseTalentBullets', () => {
  it('reads the names and the why from each bold-led bullet with a dash', () => {
    const bullets = parseTalentBullets(MARKDOWN);
    expect(bullets.map((b) => b.names)).toEqual([
      ['Improved Frostbolt'],
      ['Ice Shards', "Winter's Chill"],
      ['Feral Instinct', 'Savage Fury', 'Rend and Tear'],
    ]);
    expect(bullets[0]!.why).toBe('Up to 0.5 seconds off the cast');
  });

  it('skips a bullet with no dash or no bold name', () => {
    expect(parseTalentBullets('- **Role:** Caster\n- plain text — with a dash')).toEqual([]);
  });
});

describe('withoutTalentBullets', () => {
  it('cuts the parsed bullets and keeps everything else', () => {
    const out = withoutTalentBullets(MARKDOWN);
    expect(out).toContain('Intro sentence.');
    expect(out).toContain('Not a talent bullet at all.');
    expect(out).toContain('Closing paragraph.');
    expect(out).not.toContain('Ice Shards');
    expect(out).not.toMatch(/\n{3,}/);
  });

  it('returns the text unchanged when it has no such bullets', () => {
    expect(withoutTalentBullets('Just prose.')).toBe('Just prose.');
  });
});

describe('keyTalentsFor', () => {
  it('attaches the build icon and rank to each named talent', () => {
    const entries = keyTalentsFor(MARKDOWN, [TREE]);
    expect(entries[0]).toEqual({
      talents: [{ name: 'Improved Frostbolt', icon: 'icon_1', rank: 5, maxRank: 5 }],
      why: 'Up to 0.5 seconds off the cast',
    });
    expect(entries[1]!.talents.map((t) => t.icon)).toEqual(['icon_2', 'icon_3']);
  });

  it('keeps a name the build does not have, without an icon', () => {
    const entries = keyTalentsFor(MARKDOWN, [TREE]);
    expect(entries[2]!.talents[0]).toEqual({ name: 'Feral Instinct' });
  });

  it('falls back to the deepest talents of the build, with their own description, when the guide has no bullets', () => {
    const entries = keyTalentsFor('Only prose here.', [TREE]);
    expect(entries.map((e) => e.talents[0]!.name)).toEqual([
      "Winter's Chill",
      'Ice Shards',
      'Improved Frostbolt',
    ]);
    expect(entries[0]!.why).toBe('Stacks a debuff');
  });

  it('is empty without a build', () => {
    expect(keyTalentsFor('Only prose here.', undefined)).toEqual([]);
  });
});
