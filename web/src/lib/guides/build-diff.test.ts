// web/src/lib/guides/build-diff.test.ts
import { describe, expect, it } from 'vitest';
import { buildDiffLine, diffBuilds } from './build-diff';
import { readFileSync, readdirSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { decodeBuildTrees, type DecodedTalentView, type DecodedTreeView } from './build-tree';

function talent(id: number, name: string, rank: number): DecodedTalentView {
  return {
    id,
    name,
    icon: `icon_${id}`,
    max_rank: 5,
    tier: id,
    column: 0,
    prereq_talent_id: null,
    prereq_rank: null,
    ranks: [],
    spell_id: id,
    rank,
  };
}

function tree(...talents: DecodedTalentView[]): DecodedTreeView {
  return {
    name: 'Tree',
    position: 0,
    talents,
    points: talents.reduce((sum, t) => sum + t.rank, 0),
  };
}

describe('diffBuilds', () => {
  const leveling = [
    tree(talent(1, 'Improved Fireball', 2), talent(2, 'Impact', 3), talent(3, 'Incineration', 0)),
  ];
  const raid = [
    tree(talent(1, 'Improved Fireball', 0), talent(2, 'Impact', 0), talent(3, 'Incineration', 5)),
  ];

  it('lists what loses points, what gains them and how many moved', () => {
    expect(diffBuilds(leveling, raid)).toEqual({
      from: [
        { name: 'Improved Fireball', points: 2 },
        { name: 'Impact', points: 3 },
      ],
      to: [{ name: 'Incineration', points: 5 }],
      moved: 5,
    });
  });

  it('is undefined for the same build', () => {
    expect(diffBuilds(leveling, leveling)).toBeUndefined();
  });
});

describe('buildDiffLine', () => {
  it('names the move in one sentence', () => {
    const line = buildDiffLine({
      from: [
        { name: 'A', points: 2 },
        { name: 'B', points: 3 },
      ],
      to: [{ name: 'C', points: 5 }],
      moved: 5,
    });
    expect(line).toBe('The raid build moves 5 points from A (2) and B (3) to C (5).');
  });

  it('says "1 point" in the singular and names every talent, however long the list', () => {
    const line = buildDiffLine({
      from: [{ name: 'A', points: 1 }],
      to: ['C', 'D', 'E', 'F', 'G', 'H'].map((name) => ({ name, points: 1 })),
      moved: 1,
    });
    expect(line).toBe(
      'The raid build moves 1 point from A (1) to C (1), D (1), E (1), F (1), G (1) and H (1).',
    );
  });

  it('handles a build that only adds or only removes points', () => {
    expect(buildDiffLine({ from: [], to: [{ name: 'C', points: 2 }], moved: 0 })).toBe(
      'The raid build puts 2 more points into C (2).',
    );
    expect(buildDiffLine({ from: [{ name: 'A', points: 2 }], to: [], moved: 2 })).toBe(
      'The raid build takes 2 points out of A (2).',
    );
  });

  it('is undefined when there is no difference', () => {
    expect(buildDiffLine(undefined)).toBeUndefined();
  });
});

// The real guides: every talent the raid build changes is named in the line, and the points
// in brackets add up to the number the line states (guide-accuracy round 1 found Affliction
// reading 10 of 15 and Demonology 11 of 12 because "and N more" hid named talents).
const guidesRoot = path.resolve(path.dirname(fileURLToPath(import.meta.url)), '../../content/guides');

function frontmatterCode(text: string, key: string): string | undefined {
  return new RegExp(`^${key}: '([^']+)'`, 'm').exec(text)?.[1];
}

const guidePairs = readdirSync(guidesRoot, { withFileTypes: true })
  .filter((entry) => entry.isDirectory())
  .flatMap((dir) =>
    readdirSync(path.join(guidesRoot, dir.name))
      .filter((file) => file.endsWith('.md'))
      .map((file) => {
        const text = readFileSync(path.join(guidesRoot, dir.name, file), 'utf8');
        return {
          page: `${dir.name}/${file.replace(/\.md$/, '')}`,
          build: frontmatterCode(text, 'build'),
          raid: frontmatterCode(text, 'raidBuild'),
        };
      }),
  )
  .filter(
    (pair): pair is { page: string; build: string; raid: string } =>
      pair.build !== undefined && pair.raid !== undefined,
  );

function bracketTotal(text: string): number {
  return [...text.matchAll(/\((\d+)\)/g)].reduce((sum, match) => sum + Number(match[1]), 0);
}

describe('the raid-build line on the real guides', () => {
  it('covers the three warlock specs', () => {
    expect(guidePairs.map((pair) => pair.page)).toEqual(
      expect.arrayContaining(['warlock/affliction', 'warlock/demonology', 'warlock/destruction']),
    );
  });

  it.each(guidePairs)('$page names every changed talent and its points add up', ({ build, raid }) => {
    const dataBuild = build.split(':')[1]!;
    const before = decodeBuildTrees(dataBuild, build);
    const after = decodeBuildTrees(dataBuild, raid);
    expect(before).toBeDefined();
    expect(after).toBeDefined();
    const diff = diffBuilds(before!, after!);
    const line = buildDiffLine(diff);
    if (diff === undefined) return expect(line).toBeUndefined();
    for (const entry of [...diff.from, ...diff.to]) expect(line).toContain(`${entry.name} (${entry.points})`);
    const [fromPart, toPart] = line!.split(' to ');
    expect(bracketTotal(fromPart!)).toBe(diff.moved);
    expect(bracketTotal(toPart!)).toBe(diff.moved);
  });

  it('reads Affliction as 15 points out of seven talents into four', () => {
    const pair = guidePairs.find((entry) => entry.page === 'warlock/affliction')!;
    const line = buildDiffLine(
      diffBuilds(decodeBuildTrees('1.60.1.70755', pair.build)!, decodeBuildTrees('1.60.1.70755', pair.raid)!),
    );
    expect(line).toBe(
      'The raid build moves 15 points from Soul Harvest (2), Improved Drains (3), Nightfall (2), ' +
        'Soul Siphon (3), Wrack (1), Improved Health Funnel (2) and Improved Voidwalker (2) to ' +
        'Improved Bane of Agony (1), Malevolence (5), Improved Shadow Bolt (5) and Ruin (4).',
    );
  });
});
