// The parity reference-mock report (scripts/parity-mocks-stale.mjs, design/parity.md).
import { mkdirSync, mkdtempSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import path from 'node:path';
import { describe, expect, it } from 'vitest';
import { formatStaleReport, normalizeMock, staleMocks } from '../../../scripts/parity-mocks-stale.mjs';

const ICON =
  '<img class="ic q4" src="data:image/webp;base64,AAAA" alt="" width="40" height="40" style="width:40px">';
const SOCKET = '<span class="ic q4" style="width:40px;height:40px"></span>';

function folders(files: { rendered: Record<string, string>; committed: Record<string, string> }) {
  const root = mkdtempSync(path.join(tmpdir(), 'parity-stale-'));
  const rendered = path.join(root, 'rendered');
  const committed = path.join(root, 'committed');
  for (const [dir, entries] of [
    [rendered, files.rendered],
    [committed, files.committed],
  ] as const) {
    mkdirSync(dir, { recursive: true });
    for (const [name, html] of Object.entries(entries)) writeFileSync(path.join(dir, name), html);
  }
  return { rendered, committed };
}

describe('normalizeMock', () => {
  it('reads an icon and its empty socket as the same block', () => {
    expect(normalizeMock(`<li>${ICON}</li>`)).toBe(normalizeMock(`<li>${SOCKET}</li>`));
  });

  it('reads a public image the same from any depth the mock was written at', () => {
    const at = (prefix: string) =>
      `<img class="site-panel-crest" src="${prefix}/icons/hd/crests/warrior.webp" srcset="${prefix}/a.webp 1x, ${prefix}/b.webp 2x">` +
      `<link href="${prefix}/favicon.svg"><i style="background-image:url(${prefix}/icons/x.webp)"></i>`;
    expect(normalizeMock(at('../../../web/public'))).toBe(normalizeMock(at('../../../../public')));
  });

  it('keeps a different public image apart', () => {
    expect(normalizeMock('<img src="../../../web/public/icons/a.webp">')).not.toBe(
      normalizeMock('<img src="../../../../public/icons/b.webp">'),
    );
  });

  it('keeps an icon of another quality apart', () => {
    expect(normalizeMock(ICON)).not.toBe(normalizeMock(SOCKET.replace('q4', 'q3')));
  });
});

describe('staleMocks', () => {
  it('finds nothing when the committed mock matches but for icon files', () => {
    const dirs = folders({
      rendered: { 'web-bis-a.html': `<p>Trust</p>${ICON}` },
      committed: { 'web-bis-a.html': `<p>Trust</p>${SOCKET}` },
    });
    expect(staleMocks(dirs.rendered, dirs.committed, 'web-bis-')).toEqual([]);
  });

  it('names a mock whose text moved on, and one never committed', () => {
    const dirs = folders({
      rendered: { 'web-bis-a.html': '<p>Ranked by a longer trust line</p>', 'web-bis-b.html': '<p>b</p>' },
      committed: { 'web-bis-a.html': '<p>Ranked by trust</p>' },
    });
    const stale = staleMocks(dirs.rendered, dirs.committed, 'web-bis-');
    expect(stale.map((s) => [s.file, s.missing])).toEqual([
      ['web-bis-a.html', false],
      ['web-bis-b.html', true],
    ]);
    expect(stale[0]).toMatchObject({ rendered: expect.stringContaining('longer trust line') });
  });

  it('only reads the page its prefix names', () => {
    const dirs = folders({ rendered: { 'bis-facts.json': '{}', 'web-bis-a.html': 'x' }, committed: {} });
    expect(staleMocks(dirs.rendered, dirs.committed, 'web-bis-').map((s) => s.file)).toEqual([
      'web-bis-a.html',
    ]);
  });

  it('fails when the parity specs have not rendered the mocks', () => {
    expect(() => staleMocks('/no/such/folder', '/no', 'web-bis-')).toThrow(/run the parity specs first/);
  });
});

describe('formatStaleReport', () => {
  it('is empty when nothing is stale', () => {
    expect(formatStaleReport({ bis: [], sim: [] })).toBe('');
  });

  it('names the stale files and the exact regen commands of their page', () => {
    const report = formatStaleReport({ bis: [{ file: 'web-bis-a.html', missing: true }], sim: [] });
    expect(report).toContain('web-bis-a.html: not committed');
    expect(report).toContain(
      'node design/mocks/bis_facts.mjs http://localhost:4500 && python3 design/mocks/ws_bis.py',
    );
    expect(report).not.toContain('ws_sim.py');
  });
});
