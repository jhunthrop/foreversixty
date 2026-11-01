// web/src/lib/sim/raid-buff-names.test.ts
// 2026-09-21 result-page review, Defect 2: the one-sentence summary read "Spell 9910 is up
// 100% of the fight" (9910 is Thorns, from the raid-buffed preset) and the BUFFS tab was
// mostly rows named "Spell <id>" -- a raid buff from a class other than the player's own
// has no name in that class's own simnames/<class>.json (scripts/sync-data.mjs), because
// that file is pruned to the spells the SIMMED class's own spellconst carries.
// simnames/_shared.json (writeSimNames, same file) is the fix: every class's own spell
// table unioned into one, plus the handful of buffs no class owns at all. This walks every
// RAID_BUFFS (settings.ts) slug that can actually reach the player as an aura against the
// REAL checked-in build data -- not a synthetic fixture -- so a real gap in the union or a
// wrong id in the table below fails here, not on a saved sim page.
import { mkdtempSync, readFileSync, rmSync } from 'node:fs';
import { tmpdir } from 'node:os';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { describe, expect, it } from 'vitest';
import { writeSimNames } from '../../../scripts/sync-data.mjs';
import { resolveActionName, type ActionNames } from './action-names';
import raidPreset from '../../data/generated/raid-preset.json';
import { RAID_BUFFS } from './settings';

const debuffIds = new Set<string>(raidPreset.debuffs);

const REPO_ROOT = path.resolve(path.dirname(fileURLToPath(import.meta.url)), '../../../..');
const BUILD_DIR = path.join(REPO_ROOT, 'data/builds/1.60.1.69893');

/**
 * Every RAID_BUFFS slug that actually reaches the player as an aura, mapped to one real
 * spell id for it. Verified against wowsims-forever's sim/core/buffs.go (read there
 * 2026-09-21, not reproduced or edited here -- that checkout is read-only) and, for a
 * class-owned buff, cross-checked against this build's own
 * data/builds/1.60.1.69893/spellconst/<class>.json, which is what the test below actually
 * resolves against.
 *
 * The rest of RAID_BUFFS is deliberately absent:
 *   - The debuffs (the curated raid preset's "debuffs" list, curse_of_elements ..
 *     curse_of_recklessness) land on the TARGET, never the player, and
 *     sim/adapter/adapter.go's own comment says a sim result carries aura metrics for the
 *     player only -- a target debuff has no row to resolve in the first place.
 *   - blessing_of_wisdom, leader_of_the_pack, moonkin_aura: each is a flat
 *     character.AddStat(s) call in buffs.go with no RegisterAura at all, so none of the
 *     three ever produces an aura row either.
 */
const RAID_BUFF_AURA_SPELLS: Readonly<Record<string, number>> = {
  arcane_brilliance: 23028,
  gift_of_the_wild: 21850,
  power_word_fortitude: 23948,
  divine_spirit: 27841,
  blessing_of_might: 25291,
  battle_shout: 25289,
  trueshot_aura: 20906,
  heart_of_the_lion: 409583,
  strength_of_earth_totem: 25361,
  windfury_totem: 10610,
  mana_spring_totem: 10497,
};

describe('RAID_BUFF_AURA_SPELLS', () => {
  it('is exactly the RAID_BUFFS slugs that reach the player as an aura', () => {
    const noAura = ['blessing_of_wisdom', 'leader_of_the_pack', 'moonkin_aura'];
    const onPlayer = RAID_BUFFS.filter((slug) => !debuffIds.has(slug) && !noAura.includes(slug));
    expect(new Set(Object.keys(RAID_BUFF_AURA_SPELLS))).toEqual(new Set(onPlayer));
  });
});

describe('every raid buff aura the raid-buffed preset can put on the player resolves to a real name', () => {
  it('resolves through a Mage sim’s own merged name table -- the exact case that broke (spell 9910, "Spell 9910")', async () => {
    const dir = mkdtempSync(path.join(tmpdir(), 'raid-buff-names-'));
    try {
      const written = await writeSimNames(BUILD_DIR, dir);
      expect(written).toContain('simnames/mage.json');
      expect(written).toContain('simnames/_shared.json');

      const own = JSON.parse(readFileSync(path.join(dir, 'simnames', 'mage.json'), 'utf8')) as ActionNames;
      const shared = JSON.parse(
        readFileSync(path.join(dir, 'simnames', '_shared.json'), 'utf8'),
      ) as ActionNames;
      // loadActionNames' own merge (action-names.ts): the shared cross-class table, with
      // the player's own class table layered on top.
      const merged: ActionNames = {
        spell: { ...shared.spell, ...own.spell },
        item: { ...shared.item, ...own.item },
      };

      for (const [slug, spellId] of Object.entries(RAID_BUFF_AURA_SPELLS)) {
        const resolved = resolveActionName(`spell:${spellId}`, merged);
        expect(resolved, `${slug} (spell:${spellId}) resolved to ${JSON.stringify(resolved)}`).not.toMatch(
          /^Spell \d+$/,
        );
      }
    } finally {
      rmSync(dir, { recursive: true, force: true });
    }
  });
});
