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
import { RAID_BUFFS } from './settings';

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
 * Fourteen of RAID_BUFFS' 54 slugs are deliberately absent:
 *   - The twenty debuffs (curse_of_elements .. winters_chill) land on the TARGET, never the
 *     player, and sim/adapter/adapter.go's own comment says a sim result carries aura
 *     metrics for the player only -- a target debuff has no row to resolve in the first
 *     place, on this build or any other.
 *   - blessing_of_sanctuary: buffs.go's own implementation of it is commented out; the buff
 *     never applies.
 *   - blessing_of_wisdom, leader_of_the_pack, moonkin_aura: each is a flat
 *     character.AddStat(s) call in buffs.go with no RegisterAura at all, so none of the
 *     three ever produces an aura row either.
 * 54 - 20 - 1 - 3 = 30, this table's own size.
 */
const RAID_BUFF_AURA_SPELLS: Readonly<Record<string, number>> = {
  blessing_of_kings: 20217,
  blessing_of_might: 25291,
  fengus_ferocity: 22817,
  moldars_moxie: 22818,
  rallying_cry_of_the_dragonslayer: 22888,
  sayges_fortune: 23735,
  slipkiks_savvy: 22820,
  songflower_serenade: 15366,
  spirit_of_zandalar: 24425,
  warchiefs_blessing: 16609,
  arcane_brilliance: 23028,
  battle_shout: 25289,
  blood_pact: 11767,
  devotion_aura: 10293,
  divine_spirit: 27841,
  fire_resistance_aura: 19900,
  fire_resistance_totem: 10538,
  frost_resistance_aura: 19898,
  frost_resistance_totem: 10479,
  gift_of_the_wild: 21850,
  grace_of_air_totem: 25359,
  mana_spring_totem: 10497,
  nature_resistance_totem: 10601,
  power_word_fortitude: 23948,
  retribution_aura: 10301,
  sanctity_aura: 20218,
  shadow_protection: 16874,
  strength_of_earth_totem: 25361,
  thorns: 9910,
  trueshot_aura: 20906,
};

describe('RAID_BUFF_AURA_SPELLS', () => {
  it('is exactly the RAID_BUFFS slugs this task can actually walk (30 of the preset’s 54)', () => {
    for (const slug of Object.keys(RAID_BUFF_AURA_SPELLS)) {
      expect(RAID_BUFFS, `${slug} must be a real RAID_BUFFS slug`).toContain(slug);
    }
    expect(Object.keys(RAID_BUFF_AURA_SPELLS)).toHaveLength(30);
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
