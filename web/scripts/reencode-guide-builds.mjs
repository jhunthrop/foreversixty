#!/usr/bin/env node
// web/scripts/reencode-guide-builds.mjs
//
// Every spec guide's frontmatter `build:` is an FS1 code (web/src/lib/planner/fs1.ts),
// stamped with a client build id. The stamp on all 27 guides currently reads
// 1.60.1.69893; web/src/data/active-build.json's active build is 1.60.1.70009. Between
// those two builds the PALADIN and SHAMAN talent trees changed shape (a talent dropped, two
// talents swapping tier/column) -- the other seven classes' talent files are byte-identical
// between the builds. decodeFS1 (and the addon) read a tree's digit string purely by array
// index, so a build whose digits were laid out against one tree shape reads wrong against a
// different one.
//
// BEFORE assuming every guide's digits are genuinely laid out against the stamped
// 1.60.1.69893 shape and blindly remapping them, this script checks: are the digits, taken
// exactly as written, already a legal build (no over-rank, no unmet prereq, no tier spent
// before it unlocks) against the ACTIVE build's tree? If so, the digits are already correct
// for 1.60.1.70009 -- re-deriving them from the stamped build's tree would be re-reading
// already-correct ids as if they were positions in a tree they no longer describe, which
// corrupts rather than fixes (verified by hand against three of the six paladin/shaman
// guides below: reading their digits against the OLD 69893 tree overflows a max_rank,
// fails a prereq, or spends a tier before it unlocks -- impossible in-game -- while reading
// the SAME digits against the ACTIVE 70009 tree is clean, matches guide prose word for word
// (e.g. paladin/holy's "Reverence / Illumination / Infusion of Light / Divine Favor"
// priority list, paladin/retribution's "Instrument of Law ... Twist of Light" capstone
// discussion), and spends the full 51-point cap). In that case the only fix needed, and the
// only one this script makes, is the dataBuild stamp -- the tree-rank digits are untouched.
//
// Only when the digits are NOT already legal against the active build does this script fall
// back to the task's originally-specified algorithm: decode against the code's OWN stamped
// build, carry every ranked talent forward by its stable id onto the active build's tree,
// and report (never silently drop) any id the active build no longer has. That path is kept
// as a genuine fallback -- current data never takes it -- so a future guide whose digits
// really were authored against an older tree shape is still handled correctly rather than
// assumed away.
import { readFileSync, readdirSync, writeFileSync } from 'node:fs';
import { dirname, join } from 'node:path';
import { fileURLToPath } from 'node:url';

const webRoot = join(dirname(fileURLToPath(import.meta.url)), '..');
const repoRoot = join(webRoot, '..');
const guidesRoot = join(webRoot, 'src/content/guides');

const ACTIVE_BUILD = JSON.parse(readFileSync(join(webRoot, 'src/data/active-build.json'), 'utf8')).build;
const POINTS_PER_TIER = 5; // web/src/lib/planner/types.ts's own constant, duplicated here --
// this is a plain Node script with no loader for this repo's extensionless TS imports.

/** `FS1:<build>:<class>:<race>:<t1>/<t2>/<t3>:<gear...>` -- gear is captured raw (it may
 * itself contain `:` for `item_id:enchant:suffix`) and never touched by this script. */
const CODE_RE = /^FS1:([^:]*):([^:]*):([^:]*):([^:]*):(.*)$/;

/** One base-36 digit per rank, trailing zeros trimmed -- byte-identical to fs1.ts's own
 * `encodeTree`. Reimplemented here rather than imported: fs1.ts's extensionless TS imports
 * (`./rules`, `./types`) only resolve under a bundler, not a bare `node scripts/*.mjs`. */
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
  // Sanity check: this script indexes a tree's digits straight into `tree.talents[i]` with
  // no re-sort, same as decodeFS1/orderFromRanks/builds.test.ts's talentAt. That is only
  // correct if the file itself already ships each tree's talents tier/column ascending.
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
function legalityIssues(trees, digitsPerTree) {
  const issues = [];
  const rankOf = new Map();
  const byId = new Map();
  const pointsInTierByTree = trees.map(() => new Map());

  trees.forEach((tree, treeIndex) => {
    const digits = digitsPerTree[treeIndex] ?? [];
    tree.talents.forEach((talent, digitIndex) => {
      const raw = digits[digitIndex] ?? 0;
      byId.set(talent.id, talent);
      if (raw <= 0) return;
      rankOf.set(talent.id, raw);
      pointsInTierByTree[treeIndex].set(
        talent.tier,
        (pointsInTierByTree[treeIndex].get(talent.tier) ?? 0) + raw,
      );
      if (raw > talent.max_rank) {
        issues.push(`${tree.name}/${talent.name}: rank ${raw} > max_rank ${talent.max_rank}`);
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

function pointsPerTree(trees, digitsPerTree) {
  return trees.map((tree, treeIndex) => {
    const digits = digitsPerTree[treeIndex] ?? [];
    return { name: tree.name, points: digits.reduce((sum, r) => sum + (r ?? 0), 0) };
  });
}

/**
 * Re-stamps one guide's code onto `toBuild`. Prefers keeping the digits exactly as written
 * (stamp-only) whenever they already form a legal build against `toBuild`'s tree; only
 * remaps by talent id, from the code's own stamped build, when they do not.
 */
function remapCode(code, toBuild) {
  const match = CODE_RE.exec(code);
  if (!match) throw new Error(`not an FS1 code: ${code}`);
  const [, fromBuild, classSlug, raceSlug, treeField, gearTail] = match;

  const fromTrees = talentsFor(fromBuild, classSlug);
  const toTrees = talentsFor(toBuild, classSlug);
  const treeRanks = parseTreeField(treeField);

  const fromBuildIssues = fromBuild === toBuild ? [] : legalityIssues(fromTrees, treeRanks);
  const asIsIssues = legalityIssues(toTrees, treeRanks);

  if (asIsIssues.length === 0) {
    const newCode = `FS1:${toBuild}:${classSlug}:${raceSlug}:${treeField}:${gearTail}`;
    return {
      newCode,
      mode: 'stamp-only',
      before: pointsPerTree(fromTrees, treeRanks),
      after: pointsPerTree(toTrees, treeRanks),
      dropped: [],
      staleStampWasIllegal: fromBuildIssues,
    };
  }

  // Fallback: the digits are not legal as written against the active build. Decode them
  // against the code's OWN stamped build (where they must be legal -- if they are not
  // legal there either, something is wrong with the guide itself and this throws rather
  // than guessing), then carry every ranked talent forward by stable id.
  if (fromBuildIssues.length > 0) {
    throw new Error(
      `${classSlug} code is illegal against both its own build and ${toBuild}: ${fromBuildIssues.join('; ')}`,
    );
  }

  const rankById = new Map();
  fromTrees.forEach((tree, treeIndex) => {
    const ranks = treeRanks[treeIndex] ?? [];
    tree.talents.forEach((talent, digitIndex) => {
      const rank = ranks[digitIndex] ?? 0;
      if (rank > 0) rankById.set(talent.id, rank);
    });
  });

  const toById = new Map();
  toTrees.forEach((tree, treeIndex) => {
    tree.talents.forEach((talent, digitIndex) => {
      toById.set(talent.id, { treeIndex, digitIndex, maxRank: talent.max_rank });
    });
  });

  const newTreeRanks = toTrees.map((tree) => new Array(tree.talents.length).fill(0));
  const dropped = [];
  for (const [id, rank] of rankById) {
    const target = toById.get(id);
    if (!target) {
      const name = fromTrees.flatMap((t) => t.talents).find((t) => t.id === id)?.name ?? `#${id}`;
      dropped.push({ id, name });
      continue;
    }
    newTreeRanks[target.treeIndex][target.digitIndex] = Math.min(rank, target.maxRank);
  }

  const newTreeField = newTreeRanks.map(encodeTree).join('/');
  const newCode = `FS1:${toBuild}:${classSlug}:${raceSlug}:${newTreeField}:${gearTail}`;

  return {
    newCode,
    mode: 'id-remap',
    before: pointsPerTree(fromTrees, treeRanks),
    after: pointsPerTree(toTrees, newTreeRanks),
    dropped,
    staleStampWasIllegal: [],
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

function main() {
  const guides = findGuides();
  const report = [];
  for (const guide of guides) {
    const result = remapCode(guide.oldCode, ACTIVE_BUILD);
    const raw = readFileSync(guide.path, 'utf8'); // re-read: a guide has up to two codes
    const line = new RegExp(`^${guide.key}: '[^']*'$`, 'm');
    const newRaw = raw.replace(line, `${guide.key}: '${result.newCode}'`);
    if (!newRaw.includes(`${guide.key}: '${result.newCode}'`)) {
      throw new Error(`${guide.id}: failed to rewrite ${guide.key}: line`);
    }
    writeFileSync(guide.path, newRaw);
    report.push({ id: guide.id, ...result });
  }

  for (const r of report) {
    console.log(`\n${r.id} (${r.mode})`);
    for (let i = 0; i < r.before.length; i++) {
      const before = r.before[i];
      const after = r.after[i];
      const changed = before.points !== after.points ? '  <-- changed' : '';
      console.log(`  ${before.name.padEnd(12)} ${before.points} -> ${after.points}${changed}`);
    }
    if (r.dropped.length > 0) {
      console.log(`  dropped: ${r.dropped.map((d) => `${d.name} (#${d.id})`).join(', ')}`);
    }
    if (r.staleStampWasIllegal.length > 0) {
      console.log(
        `  note: these digits would be illegal against the stale stamp's own build: ${r.staleStampWasIllegal.join('; ')}`,
      );
    }
  }
  const idRemapped = report.filter((r) => r.mode === 'id-remap').length;
  console.log(
    `\nRewrote ${report.length} guide(s) to build ${ACTIVE_BUILD} (${idRemapped} by id-remap, ${report.length - idRemapped} stamp-only).`,
  );
}

main();
