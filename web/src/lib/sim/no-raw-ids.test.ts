// web/src/lib/sim/no-raw-ids.test.ts
// 2026-09-21 result-page review round 2: raid-buff-names.test.ts pinned RAID_BUFFS itself,
// but production still showed four MORE kinds of raw id once that pass was fixed -- a
// world buff variant this repo had not named yet, a talent-granted passive proc
// (spellconst is the CASTABLE spellbook only), a racial (no class and no RAID_BUFFS entry
// owns one) and an ITEM action (a potion, a mana gem -- items/<class>.json is gear only).
// This is the guarantee made testable and general, rather than id by id: every action and
// aura key sim/adapter can emit for five raid-buffed test characters across five different
// classes, read from two checked-in golden fixtures (sim/adapter/testdata, a real full
// gear/talent Frost Mage and Fury Warrior) plus three minimal-but-valid native runs
// (web/src/fixtures/sim/action-coverage/*.json, whose own `provenance` field says how), resolved
// through the exact pipeline the page uses (loadActionNames' merge) and asserted never to
// render as `Spell <n>`/`Item <n>` -- and, short of that hard floor, resolved by REAL name
// for the large majority of what a raid-buffed run can produce.
import { mkdtempSync, readFileSync, rmSync } from 'node:fs';
import { tmpdir } from 'node:os';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { describe, expect, it } from 'vitest';
import { writeSimNames } from '../../../scripts/sync-data.mjs';
import paladinCoverage from '../../fixtures/sim/action-coverage/paladin.json';
import priestCoverage from '../../fixtures/sim/action-coverage/priest.json';
import rogueCoverage from '../../fixtures/sim/action-coverage/rogue.json';
import { resolveActionName, type ActionNames } from './action-names';

const REPO_ROOT = path.resolve(path.dirname(fileURLToPath(import.meta.url)), '../../../..');
const BUILD_DIR = path.join(REPO_ROOT, 'data/builds/1.60.1.69893');
const ADAPTER_TESTDATA = path.join(REPO_ROOT, 'sim/adapter/testdata');

interface GoldenSummary {
  damage_done: { abilities: { name: string }[] }[];
  casts: { spell_name: string }[];
  auras: { name: string }[];
}

/** Every action/aura key one adapter summary carries, player and pets, abilities, casts
 *  and auras alike -- the same three lists namedSummary (sentence.ts) resolves. */
function keysOf(summary: GoldenSummary): string[] {
  const keys = new Set<string>();
  for (const actor of summary.damage_done) {
    for (const ability of actor.abilities) keys.add(ability.name);
  }
  for (const cast of summary.casts) keys.add(cast.spell_name);
  for (const aura of summary.auras) keys.add(aura.name);
  return [...keys];
}

function readGolden(spec: string): GoldenSummary {
  const file = path.join(ADAPTER_TESTDATA, `${spec}.summary.json.golden`);
  return JSON.parse(readFileSync(file, 'utf8')) as GoldenSummary;
}

/** Five test characters, five different classes: two checked-in raid-buffed goldens (a
 *  real Frost Mage and a real Fury Warrior, full gear and talents) plus three minimal
 *  native runs across three more classes (coverage/*.json). */
const CHARACTERS: { classSlug: string; keys: string[] }[] = [
  { classSlug: 'mage', keys: keysOf(readGolden('mage-frost')) },
  { classSlug: 'warrior', keys: keysOf(readGolden('warrior-fury')) },
  { classSlug: 'priest', keys: priestCoverage.keys },
  { classSlug: 'rogue', keys: rogueCoverage.keys },
  { classSlug: 'paladin', keys: paladinCoverage.keys },
];

/** A row that still reads as the raw id, spelled out -- exactly the shape this task exists
 *  to end, in either kind. */
const RAW_ID_SHAPE = /^(Spell|Item) \d+$/;

/** Below this share of a class's own distinct ids resolving to a real name (not the
 *  generic "An unnamed spell/item" fallback), something in the shared-table union broke --
 *  this is a floor, not a target: a handful of ids are genuinely engine-internal with no
 *  client spell behind them at all (12577, Clearcasting's proc aura, is not in spells.json
 *  or any talents.json -- confirmed 2026-09-21 against wowsims-forever's own source -- and
 *  is expected to read as the honest fallback, not a regression). */
const MIN_RESOLVED_SHARE = 0.85;

describe('no raw spell/item id reaches the player, across five raid-buffed characters', () => {
  for (const { classSlug, keys } of CHARACTERS) {
    it(`${classSlug}: never renders "Spell <n>" or "Item <n>", and resolves most ids by name`, async () => {
      const dir = mkdtempSync(path.join(tmpdir(), 'no-raw-ids-'));
      try {
        await writeSimNames(BUILD_DIR, dir);
        const own = JSON.parse(
          readFileSync(path.join(dir, 'simnames', `${classSlug}.json`), 'utf8'),
        ) as ActionNames;
        const shared = JSON.parse(
          readFileSync(path.join(dir, 'simnames', '_shared.json'), 'utf8'),
        ) as ActionNames;
        const merged: ActionNames = {
          spell: { ...shared.spell, ...own.spell },
          item: { ...shared.item, ...own.item },
        };

        let resolved = 0;
        let total = 0;
        const rawShaped: string[] = [];
        for (const key of keys) {
          const name = resolveActionName(key, merged);
          if (RAW_ID_SHAPE.test(name)) rawShaped.push(`${key} -> ${name}`);
          // "other:" and "unknown" keys (move, mana_gain, an attack tag...) are prose by
          // construction (resolveActionName never looks them up in a table at all) and are
          // not part of the resolved/unresolved question this share is asking.
          if (key.startsWith('spell:') || key.startsWith('item:')) {
            total += 1;
            // startsWith, not equality: a tagged/ranked key's fallback still carries
            // resolveActionName's own "(2)"/"(Rank 3)" suffix (simCopy.actionVariant)
            // after the unnamed prefix.
            if (!name.startsWith(simCopyUnnamed(key))) resolved += 1;
          }
        }

        expect(rawShaped, `raw-id-shaped names: ${rawShaped.join(', ')}`).toEqual([]);
        const share = total === 0 ? 1 : resolved / total;
        expect(
          share,
          `${classSlug} resolved ${resolved}/${total} spell/item ids by name (${(share * 100).toFixed(1)}%)`,
        ).toBeGreaterThanOrEqual(MIN_RESOLVED_SHARE);
      } finally {
        rmSync(dir, { recursive: true, force: true });
      }
    });
  }
});

/** Which of the two unresolved-fallback strings a key's own kind would produce, so the
 *  resolved-share count above does not need to import simCopy just to compare against it. */
function simCopyUnnamed(key: string): string {
  return key.startsWith('item:') ? 'An unnamed item' : 'An unnamed spell';
}
