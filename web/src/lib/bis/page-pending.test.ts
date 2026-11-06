// @vitest-environment jsdom
import { afterEach, describe, expect, it } from 'vitest';
import { BIS_PENDING_SCRIPT, PENDING_ATTRIBUTE } from './page-pending';

function run(url: string): boolean {
  window.history.replaceState(null, '', url);
  document.documentElement.removeAttribute(PENDING_ATTRIBUTE);
  new Function(BIS_PENDING_SCRIPT)();
  return document.documentElement.hasAttribute(PENDING_ATTRIBUTE);
}

describe('the BiS head script', () => {
  afterEach(() => localStorage.clear());

  it('marks a deep link with state params or the legacy hash', () => {
    expect(run('/bis/hunter/marksmanship?band=60')).toBe(true);
    expect(run('/bis/hunter/marksmanship#band-horde-40')).toBe(true);
  });

  it('leaves the default URL alone', () => {
    expect(run('/bis/hunter/marksmanship')).toBe(false);
  });

  it('marks a remembered Sources choice', () => {
    localStorage.setItem('fs.bis.sources.leveling', 'dungeon');
    expect(run('/bis/hunter/marksmanship')).toBe(true);
  });
});
