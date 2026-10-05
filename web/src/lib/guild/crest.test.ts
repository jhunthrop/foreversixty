// web/src/lib/guild/crest.test.ts
import { describe, expect, it } from 'vitest';
import { CREST_MAX_BYTES, precheckCrestFile } from './crest';

function fakeFile(type: string, size: number): File {
  return new File([new Uint8Array(size)], 'crest', { type });
}

describe('precheckCrestFile', () => {
  it('passes a small PNG/JPEG/WebP file', () => {
    expect(precheckCrestFile(fakeFile('image/png', 1024))).toBeNull();
    expect(precheckCrestFile(fakeFile('image/jpeg', 1024))).toBeNull();
    expect(precheckCrestFile(fakeFile('image/webp', 1024))).toBeNull();
  });

  it('rejects a file whose type is not one of the three', () => {
    expect(precheckCrestFile(fakeFile('image/gif', 1024))).toBe(
      'That file needs to be a PNG, JPEG or WebP image.',
    );
  });

  it('rejects an oversize file with the exact server wording for the size', () => {
    const file = fakeFile('image/png', 3_250_585); // 3.1 MB
    expect(precheckCrestFile(file)).toBe('That file is 3.1 MB; the limit is 2 MB.');
  });

  it('passes a file exactly at the limit', () => {
    expect(precheckCrestFile(fakeFile('image/png', CREST_MAX_BYTES))).toBeNull();
  });
});
