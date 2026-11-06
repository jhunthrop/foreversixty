// web/src/lib/sim/cold-paste.test.ts
import { describe, expect, it } from 'vitest';
import { validateColdPaste } from './cold-paste';

// The same fixture tests/e2e/current-character.spec.ts and sim-run.spec.ts paste: a known-
// good FS1 v2 code, decoded here with a literal build id rather than the active one --
// decodeFS1 never reads that field for anything but presence (sources.test.ts's own FURY
// constant does the same).
const FURY = 'FS1:1.15.9.69722:warrior:orc:0/5530515/0:head=12640,main_hand=11726';

describe('validateColdPaste', () => {
  it('decodes the fixture addon export and hands back its code unchanged', () => {
    const result = validateColdPaste(FURY);
    expect(result).toEqual({ ok: true, code: FURY });
  });

  it('decodes the same code wrapped in a planner link’s ?code=', () => {
    const result = validateColdPaste(`https://foreversixty.gg/planner?code=${encodeURIComponent(FURY)}`);
    expect(result).toEqual({ ok: true, code: FURY });
  });

  it('tolerates leading/trailing whitespace', () => {
    expect(validateColdPaste(`  ${FURY}  `)).toEqual({ ok: true, code: FURY });
  });

  it('is silent, not an error, while the box is empty', () => {
    expect(validateColdPaste('')).toEqual({ ok: false, message: null });
    expect(validateColdPaste('   ')).toEqual({ ok: false, message: null });
  });

  it('rejects a bare saved-build id or link -- no local decode exists for that path', () => {
    const result = validateColdPaste('https://foreversixty.gg/b/abc123');
    expect(result.ok).toBe(false);
    expect((result as { message: string | null }).message).toBe(
      'That does not look like an addon export or a build code.',
    );
  });

  it('rejects unrecognisable text with the same message', () => {
    const result = validateColdPaste('not a real export');
    expect(result).toEqual({
      ok: false,
      message: 'That does not look like an addon export or a build code.',
    });
  });

  it('surfaces decodeFS1’s own message for a malformed FS1 string, never a generic one', () => {
    const result = validateColdPaste('FS1:1.15.9.69722:warrior');
    expect(result).toEqual({ ok: false, message: 'That code is missing its talent and gear fields.' });
  });
});
