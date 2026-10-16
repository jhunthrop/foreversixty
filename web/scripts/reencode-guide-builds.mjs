#!/usr/bin/env node
// web/scripts/reencode-guide-builds.mjs
//
// Every spec guide's frontmatter `build:` and `raidBuild:` is an FS1 code
// (web/src/lib/planner/fs1.ts) stamped with a client build id. When the active build
// (web/src/data/active-build.json) moves, this script carries each code onto the new
// build's talent trees.
//
// decodeFS1 (and the addon) read a tree's digit string purely by array index, so a build
// laid out against one tree shape reads wrong against another. Between two client builds a
// talent can move tier or column, lose or gain ranks, be renamed in place (the same talent
// id with a different effect), or be dropped altogether; and a tree can reshuffle so that
// ids and names change partners (1.60.1.70374 -> 70425 swapped Tidal Focus and Totemic
// Focus's positions, and moved Lava Burst, Elemental Mastery and Earthbound between ids).
// Neither "the digits as written" nor "the talent id" survives all of that: the digits
// silently land on whatever now sits at that index, and the id silently lands on whatever
// talent now wears it. The one handle that means the same thing to the guide's prose, the
// planner and a player is the talent's NAME within its tree, so that is what this script
// carries: decode the code against its OWN stamped build into (tree, talent name) -> rank,
// then place every ranked talent onto the active build's tree by that name.
//
// A carry succeeds only when every name is found, no rank exceeds the new max_rank, the
// result is legal (no unmet prereq, every tier reached with 5 points per tier below it) and
// it still spends the full 51 points. Otherwise the tree changed in a way a person has to
// resolve: the script prints what it could not place, writes NOTHING, and exits 1, so the
// client watcher's activation stops before it can commit a build that reads wrong.
//
// Identical trees reduce this to a pure re-stamp: every name lands where it was.
import { readFileSync, readdirSync, writeFileSync } from 'node:fs';
import { dirname, join } from 'node:path';
import { fileURLToPath } from 'node:url';

const webRoot = join(dirname(fileURLToPath(import.meta.url)), '..');
const repoRoot = join(webRoot, '..');
const guidesRoot = join(webRoot, 'src/content/guides');

const ACTIVE_BUILD = JSON.parse(readFileSync(join(webRoot, 'src/data/active-build.json'), 'utf8')).build;
// web/src/lib/planner/types.ts's own constants, duplicated here -- this is a plain Node
// script with no loader for this repo's extensionless TS imports.
const POINTS_PER_TIER = 5;
const TOTAL_POINTS = 51;

/** `FS1:<build>:<class>:<race>:<t1>/<t2>/<t3>:<gear...>` -- gear is captured raw (it may
 * itself contain `:` for `item_id:enchant:suffix`) and never touched by this script. */
const CODE_RE = /^FS1:([^:]*):([^:]*):([^:]*):([^:]*):(.*)$/;

/** One base-36 digit per rank, trailing zeros trimmed -- byte-identical to fs1.ts's own
 * `encodeTree`. */
function encodeTree(ranks) {
  const digits = ranks.map((rank) => Math.max(0, Math.min(35, rank)).toString(36));
  const trimmed = digits.join('').replace(/0+$/, '');
  return trimmed === '' ? '0' : trimmed;
}

function parseTreeField(field) {
  return field.split('/').map((tree) => [...tree].map((digit) => Number.parseInt(digit, 36)));
}

const talentFileCache = new Map();
function talentsFor(build, classSlug) {
  const key = `${build}/${classSlug}`;
  const cached = talentFileCache.get(key);
  if (cached) return cached;
  const file = JSON.parse(
    readFileSync(join(repoRoot, 'data/builds', build, 'talents', `${classSlug}.json`), 'utf8'),
  );
  const trees = [...file.trees].sort((a, b) => a.position - b.position);
  // This script indexes a tree's digits straight into `tree.talents[i]` with no re-sort,
  // same as decodeFS1/orderFromRanks/builds.test.ts's talentAt. That is only correct if
  // the file itself already ships each tree's talents tier/column ascending.
  for (const tree of trees) {
    for (let i = 1; i < tree.talents.length; i++) {
      const prev = tree.talents[i - 1];
      const cur = tree.talents[i];
      if (cur.tier < prev.tier || (cur.tier === prev.tier && cur.column <= prev.column)) {
        throw new Error(`${build}/${classSlug}: ${tree.name} is not tier/column ascending at index ${i}`);
      }
    }
  }
  talentFileCache.set(key, trees);
  return trees;
}

/**
 * Every legality rule the planner (rules.ts) and the content test (builds.test.ts) check:
 * no talent above its max_rank, every prereq met, every tier reached with 5 points per tier
 * below it. Returns one message per violation, empty when the ranks form a build the game
 * could actually produce.
 */
function legalityIssues(trees, ranksPerTree) {
  const issues = [];
  const rankOf = new Map();
  const byId = new Map();
  const pointsInTierByTree = trees.map(() => new Map());

  trees.forEach((tree, treeIndex) => {
    const ranks = ranksPerTree[treeIndex] ?? [];
    tree.talents.forEach((talent, index) => {
      const rank = ranks[index] ?? 0;
      byId.set(talent.id, talent);
      if (rank <= 0) return;
      rankOf.set(talent.id, rank);
      pointsInTierByTree[treeIndex].set(
        talent.tier,
        (pointsInTierByTree[treeIndex].get(talent.tier) ?? 0) + rank,
      );
      if (rank > talent.max_rank) {
        issues.push(`${tree.name}/${talent.name}: rank ${rank} > max_rank ${talent.max_rank}`);
      }
    });
  });

  for (const talent of byId.values()) {
    const have = rankOf.get(talent.id) ?? 0;
    if (have <= 0 || talent.prereq_talent_id === null) continue;
    const need = talent.prereq_rank ?? 0;
    const prereqHave = rankOf.get(talent.prereq_talent_id) ?? 0;
    if (prereqHave < need) {
      const prereqName = byId.get(talent.prereq_talent_id)?.name ?? talent.prereq_talent_id;
      issues.push(`${talent.name} needs ${need} in ${prereqName}, has ${prereqHave}`);
    }
  }

  trees.forEach((tree, treeIndex) => {
    let cumulativeBelow = 0;
    for (const tier of [...pointsInTierByTree[treeIndex].keys()].sort((a, b) => a - b)) {
      const tierPoints = pointsInTierByTree[treeIndex].get(tier) ?? 0;
      if (tierPoints > 0 && cumulativeBelow < POINTS_PER_TIER * tier) {
        issues.push(
          `${tree.name} tier ${tier}: ${tierPoints} spent, only ${cumulativeBelow} below (needs ${POINTS_PER_TIER * tier})`,
        );
      }
      cumulativeBelow += tierPoints;
    }
  });

  return issues;
}

function pointsPerTree(trees, ranksPerTree) {
  return trees.map((tree, treeIndex) => ({
    name: tree.name,
    points: (ranksPerTree[treeIndex] ?? []).reduce((sum, rank) => sum + (rank ?? 0), 0),
  }));
}

/** The code's ranked talents as the guide means them: `${tree name}/${talent name}` -> rank. */
function rankedByName(trees, ranksPerTree) {
  const ranked = new Map();
  trees.forEach((tree, treeIndex) => {
    const ranks = ranksPerTree[treeIndex] ?? [];
    tree.talents.forEach((talent, index) => {
      const rank = ranks[index] ?? 0;
      if (rank > 0) ranked.set(`${tree.name}/${talent.name}`, rank);
    });
  });
  return ranked;
}

/**
 * Carries one guide's code onto `toBuild` by talent name. `problems` is empty when the
 * carry is clean; otherwise it names everything a person has to resolve.
 */
function carryCode(code, toBuild) {
  const match = CODE_RE.exec(code);
  if (!match) throw new Error(`not an FS1 code: ${code}`);
  const [, fromBuild, classSlug, raceSlug, treeField, gearTail] = match;

  const fromTrees = talentsFor(fromBuild, classSlug);
  const toTrees = talentsFor(toBuild, classSlug);
  const fromRanks = parseTreeField(treeField);

  const problems = [];
  const stampIssues = legalityIssues(fromTrees, fromRanks);
  if (stampIssues.length > 0) {
    problems.push(`illegal against its own stamped build ${fromBuild}: ${stampIssues.join('; ')}`);
  }

  const slotByName = new Map();
  toTrees.forEach((tree, treeIndex) => {
    tree.talents.forEach((talent, index) => {
      slotByName.set(`${tree.name}/${talent.name}`, { treeIndex, index, maxRank: talent.max_rank });
    });
  });

  const toRanks = toTrees.map((tree) => new Array(tree.talents.length).fill(0));
  for (const [name, rank] of rankedByName(fromTrees, fromRanks)) {
    const slot = slotByName.get(name);
    if (!slot) {
      problems.push(`${name} (${rank} point${rank === 1 ? '' : 's'}) is not in the ${toBuild} tree`);
      continue;
    }
    if (rank > slot.maxRank) {
      problems.push(`${name} has ${slot.maxRank} rank${slot.maxRank === 1 ? '' : 's'} in ${toBuild}, the guide spends ${rank}`);
    }
    toRanks[slot.treeIndex][slot.index] = Math.min(rank, slot.maxRank);
  }

  for (const issue of legalityIssues(toTrees, toRanks)) problems.push(issue);
  const after = pointsPerTree(toTrees, toRanks);
  const total = after.reduce((sum, tree) => sum + tree.points, 0);
  if (total !== TOTAL_POINTS) problems.push(`spends ${total} points, not ${TOTAL_POINTS}`);

  const newTreeField = toRanks.map(encodeTree).join('/');
  return {
    newCode: `FS1:${toBuild}:${classSlug}:${raceSlug}:${newTreeField}:${gearTail}`,
    changed: newTreeField !== treeField,
    before: pointsPerTree(fromTrees, fromRanks),
    after,
    problems,
  };
}

/** The frontmatter keys that carry an FS1 code: the leveling build and the raid build. */
const CODE_KEYS = ['build', 'raidBuild'];

function findGuides() {
  const guides = [];
  for (const classDir of readdirSync(guidesRoot, { withFileTypes: true })) {
    if (!classDir.isDirectory()) continue;
    for (const file of readdirSync(join(guidesRoot, classDir.name))) {
      if (!file.endsWith('.md')) continue;
      const path = join(guidesRoot, classDir.name, file);
      const raw = readFileSync(path, 'utf8');
      for (const key of CODE_KEYS) {
        const lineMatch = raw.match(new RegExp(`^${key}: '([^']*)'$`, 'm'));
        if (!lineMatch) continue; // class landing pages carry no build:
        guides.push({
          id: `${classDir.name}/${file.slice(0, -'.md'.length)}#${key}`,
          path,
          key,
          oldCode: lineMatch[1],
        });
      }
    }
  }
  return guides.sort((a, b) => a.id.localeCompare(b.id));
}

function describe(result) {
  const lines = [];
  for (let i = 0; i < result.before.length; i++) {
    const before = result.before[i];
    const after = result.after[i];
    const changed = before.points !== after.points ? '  <-- changed' : '';
    lines.push(`  ${before.name.padEnd(12)} ${before.points} -> ${after.points}${changed}`);
  }
  return lines.join('\n');
}

function main() {
  const guides = findGuides();
  const results = guides.map((guide) => ({ guide, result: carryCode(guide.oldCode, ACTIVE_BUILD) }));

  const failed = results.filter(({ result }) => result.problems.length > 0);
  if (failed.length > 0) {
    console.error(`${failed.length} guide build(s) do not carry cleanly onto ${ACTIVE_BUILD}; nothing written.`);
    console.error('Resolve each by hand in the guide (code and prose), then rerun.\n');
    for (const { guide, result } of failed) {
      console.error(`${guide.id}`);
      console.error(describe(result));
      for (const problem of result.problems) console.error(`  ! ${problem}`);
      console.error('');
    }
    process.exit(1);
  }

  for (const { guide, result } of results) {
    const raw = readFileSync(guide.path, 'utf8'); // re-read: a guide has up to two codes
    const line = new RegExp(`^${guide.key}: '[^']*'$`, 'm');
    const newRaw = raw.replace(line, `${guide.key}: '${result.newCode}'`);
    if (!newRaw.includes(`${guide.key}: '${result.newCode}'`)) {
      throw new Error(`${guide.id}: failed to rewrite ${guide.key}: line`);
    }
    writeFileSync(guide.path, newRaw);
    console.log(`\n${guide.id} (${result.changed ? 'digits moved' : 'stamp-only'})`);
    console.log(describe(result));
  }
  const moved = results.filter(({ result }) => result.changed).length;
  console.log(
    `\nRewrote ${results.length} guide code(s) to build ${ACTIVE_BUILD} (${moved} with digits moved, ${results.length - moved} stamp-only).`,
  );
}

main();
