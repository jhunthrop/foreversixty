import { readdir, readFile } from 'node:fs/promises';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { describe, expect, it } from 'vitest';
// A plain .mjs so scripts/sync-rotations.mjs can import the same parser this test drives:
// the script runs under node before vitest ever starts, so the parser cannot be a .ts --
// the same reason ids-md.test.ts drives scripts/ids-md.mjs this way.
import { engineeringLeakIn, rotationNotesOf } from '../../../scripts/rotation-apl.mjs';

const webRoot = path.join(path.dirname(fileURLToPath(import.meta.url)), '../../..');
const curatedAplDir = path.join(webRoot, '..', 'data', 'curated', 'apl');

describe('rotationNotesOf', () => {
  it('carries only steps with a non-empty notes string, in priority-list order', () => {
    const apl = {
      rotation: {
        priorityList: [
          { action: { autocastOtherCooldowns: {} } },
          { notes: 'Frostbolt is the whole rotation.', action: { castSpell: {} } },
          { notes: '', action: { castSpell: {} } },
          { notes: '   ', action: { castSpell: {} } },
          { action: { castSpell: {} } },
          { notes: 'Refresh the debuff before it falls off.', action: { castSpell: {} } },
        ],
      },
    };

    expect(rotationNotesOf(apl)).toEqual([
      'Frostbolt is the whole rotation.',
      'Refresh the debuff before it falls off.',
    ]);
  });

  it('trims surrounding whitespace from a note', () => {
    const apl = {
      rotation: { priorityList: [{ notes: '  Spend every global.  ', action: {} }] },
    };
    expect(rotationNotesOf(apl)).toEqual(['Spend every global.']);
  });

  it('never reads the file-level notes field -- that is developer prose about measurement, not a step', () => {
    const apl = {
      notes: 'the engine lane measured a 2.0 percent DPS loss versus Frostbolt alone',
      rotation: {
        priorityList: [{ notes: 'Frostbolt is the whole rotation.', action: {} }],
      },
    };
    expect(rotationNotesOf(apl)).toEqual(['Frostbolt is the whole rotation.']);
  });

  it('answers an empty list for a rotation with no steps, or a spec (e.g. a tank/healer) with none at all', () => {
    expect(rotationNotesOf({ rotation: { priorityList: [] } })).toEqual([]);
    expect(rotationNotesOf({})).toEqual([]);
  });
});

describe('engineeringLeakIn', () => {
  it('flags a Go source file path', () => {
    expect(engineeringLeakIn("tracks sim/core/debuffs.go's ImprovedScorchAura")).toMatch(/file path/);
  });

  it('flags a curated-data file path', () => {
    expect(engineeringLeakIn('maintained in ui/warlock/apls/rotation.apl.json already')).toMatch(/file path/);
  });

  it('flags a bare filename with no directory prefix', () => {
    expect(engineeringLeakIn("the same shape warrior-arms.json's own Overpower line uses")).toMatch(
      /file path/,
    );
    expect(engineeringLeakIn('a real, engine-implemented DoT (see druid.py for the walk)')).toMatch(
      /file path/,
    );
  });

  it('flags "fix round" and "smoke run", case-insensitively', () => {
    expect(engineeringLeakIn('Fix round 4: tracks the engine aura')).toMatch(/fix round/);
    expect(engineeringLeakIn('unresolved under this build’s no-talent smoke run')).toMatch(/smoke run/);
  });

  it('flags "audit finding" and "not yet implemented", case-insensitively', () => {
    expect(engineeringLeakIn('(audit finding druid.md #3 -- absent from the guide)')).toMatch(
      /audit finding/,
    );
    expect(engineeringLeakIn('the functional finisher while Venom is not yet implemented')).toMatch(
      /not yet implemented/,
    );
  });

  it('flags "engine\'s", straight or curly apostrophe', () => {
    expect(engineeringLeakIn("the engine's ExtraCastCondition only lets it land")).toMatch(
      /engineering phrase/,
    );
    expect(engineeringLeakIn('totemRemainingTime is the engine’s own primitive')).toMatch(
      /engineering phrase|camelCase/,
    );
  });

  it('flags a camelCase or PascalCase code identifier but not game shorthand', () => {
    expect(engineeringLeakIn('totemRemainingTime needs no aura or spell id to check')).toMatch(
      /code identifier/,
    );
    expect(engineeringLeakIn('gated by the engine ExtraCastCondition mechanism')).toMatch(/code identifier/);
    expect(engineeringLeakIn('loops at 1-3 combo points via currentComboPoints')).toMatch(/code identifier/);
    expect(engineeringLeakIn('a PvP rank 9 reward, correctly gated')).toBeNull();
    expect(engineeringLeakIn('an AoE DoT that ticks for GCD-free upkeep')).toBeNull();
  });

  it('flags an ISO date, the shape of a commit-log timestamp', () => {
    expect(engineeringLeakIn("Multi-Shot's real cooldown, corrected 2026-09-29, zeroed Arcane Shot")).toMatch(
      /ISO date/,
    );
  });

  it('flags a hex/commit-looking token but not a plain decimal spell or item id', () => {
    expect(engineeringLeakIn('resolved as 12873 is not in spellconst')).toBeNull();
    expect(engineeringLeakIn('pinned at commit 464d1a14a64')).toMatch(/hex\/commit/);
  });

  it('passes ordinary player-facing rotation prose', () => {
    expect(engineeringLeakIn('Frostbolt is the whole rotation.')).toBeNull();
    expect(
      engineeringLeakIn('Fire Blast on cooldown as a free extra nuke that does not compete.'),
    ).toBeNull();
  });
});

describe('rotationNotesOf throws on a leaking step instead of publishing it', () => {
  it('rejects a step whose notes contain a file path', () => {
    const apl = {
      rotation: {
        priorityList: [{ notes: "Fixed in sim/core/debuffs.go's ImprovedScorchAura." }],
      },
    };
    expect(() => rotationNotesOf(apl, 'mage-fire')).toThrowError(/mage-fire step 0/);
  });

  it('accepts the same step once the leak moves to engineeringNotes', () => {
    const apl = {
      rotation: {
        priorityList: [
          {
            notes: "Pop Combustion once Improved Scorch's debuff is fully stacked.",
            engineeringNotes: "Fix round 4: sim/core/debuffs.go's ImprovedScorchAura.",
          },
        ],
      },
    };
    expect(rotationNotesOf(apl, 'mage-fire')).toEqual([
      "Pop Combustion once Improved Scorch's debuff is fully stacked.",
    ]);
  });
});

describe('every curated rotation file publishes only player-facing notes', () => {
  it('has no step whose notes leak a file path, a hex/commit token, or "fix round"/"smoke run"', async () => {
    const entries = await readdir(curatedAplDir, { withFileTypes: true });
    const specFiles = entries.filter((e) => e.isFile() && e.name.endsWith('.json'));
    expect(specFiles.length).toBeGreaterThan(0);

    for (const entry of specFiles) {
      const spec = entry.name.slice(0, -'.json'.length);
      const raw = await readFile(path.join(curatedAplDir, entry.name), 'utf8');
      // Throws with the offending spec/step/reason if any step's notes leak -- see
      // assertNoEngineeringLeak in rotation-apl.mjs.
      expect(() => rotationNotesOf(JSON.parse(raw), spec)).not.toThrow();
    }
  });
});
