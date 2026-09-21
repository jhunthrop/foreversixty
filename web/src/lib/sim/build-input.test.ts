import { describe, expect, it } from 'vitest';
import { parseBuildInput } from './build-input';

const CODE = 'FS1:1.60.1.69893:warrior:orc:3/5/0:head=12640';

describe('parseBuildInput', () => {
  it('reads a saved build id from a bare id', () => {
    expect(parseBuildInput(' abc123 ')).toEqual({ kind: 'id', id: 'abc123' });
  });

  it('reads a saved build id from a /b/ link, ignoring the query and a trailing slash', () => {
    expect(parseBuildInput('https://foreversixty.gg/b/abc123/?utm=x')).toEqual({ kind: 'id', id: 'abc123' });
  });

  it('reads the code from an unsaved planner link', () => {
    const link = `https://foreversixty.gg/planner?code=${encodeURIComponent(CODE)}`;
    expect(parseBuildInput(link)).toEqual({ kind: 'code', code: CODE });
  });

  it('reads a pasted code itself', () => {
    expect(parseBuildInput(CODE)).toEqual({ kind: 'code', code: CODE });
  });

  it('answers nothing for an empty box or a planner link with no code', () => {
    expect(parseBuildInput('  ')).toBeNull();
    expect(parseBuildInput('https://foreversixty.gg/planner')).toBeNull();
  });
});
