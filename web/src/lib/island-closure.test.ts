// web/src/lib/island-closure.test.ts
// scripts/island-closure.mjs, the startup-closure measure check-island-size.mjs budgets.
import { mkdtempSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import path from 'node:path';
import { describe, expect, it } from 'vitest';
import { startupClosure, staticImports } from '../../scripts/island-closure.mjs';

describe('an island’s startup closure', () => {
  it('counts static imports and re-exports, never a dynamic import', () => {
    const source =
      'import{a as b}from"./a.js";import"./side.js";export{c}from"../c.js";import d from "./d.js";' +
      'const e=()=>import("./lazy.js");';
    expect(staticImports(source)).toEqual(['./a.js', './side.js', '../c.js', './d.js']);
  });

  it('follows static imports transitively and counts a shared chunk once', async () => {
    const dir = mkdtempSync(path.join(tmpdir(), 'closure-'));
    writeFileSync(path.join(dir, 'entry.js'), 'import"./a.js";import"./shared.js";import("./lazy.js");');
    writeFileSync(path.join(dir, 'a.js'), 'import"./shared.js";');
    writeFileSync(path.join(dir, 'shared.js'), 'export const x=1;');
    writeFileSync(path.join(dir, 'lazy.js'), 'export const y=2;');
    const { files } = await startupClosure(path.join(dir, 'entry.js'), dir);
    expect([...files].sort()).toEqual(['a.js', 'entry.js', 'shared.js']);
  });
});
