import { describe, expect, it } from 'vitest';
import { linkProse, type ProseTarget } from './prose-tips';

const NAMES = new Map<string, ProseTarget>([
  ['Heroic Strike', { key: 'spell:25286' }],
  ['Improved Heroic Strike', { key: 'talent:124:3' }],
  ['Execute', { key: 'spell:20662' }],
  ["Nature's Grasp", { key: 'spell:16813' }],
  ['Blood Fury', { key: 'racial:Blood Fury' }],
  ['Eureka!', { key: 'racial:Eureka!' }],
  ['Flask of the Titans', { itemId: 13510, quality: 1 }],
]);

const keysIn = (html: string): string[] =>
  [...html.matchAll(/data-(?:game-tip|bis-item)="([^"]+)"/g)].map((match) => match[1]!);

describe('linkProse', () => {
  it('marks each known name with its tooltip key, keyboard focusable', () => {
    const html = linkProse('<p>Use Execute below 20%.</p>', NAMES);
    expect(html).toBe(
      '<p>Use <span class="gtip" tabindex="0" data-game-tip="spell:20662">Execute</span> below 20%.</p>',
    );
  });

  it('prefers the longest name where two overlap', () => {
    const html = linkProse('<p>Take Improved Heroic Strike, then Heroic Strike on rage.</p>', NAMES);
    expect(keysIn(html)).toEqual(['talent:124:3', 'spell:25286']);
  });

  it('matches only whole names in their own capitals', () => {
    expect(keysIn(linkProse('<p>execute the plan, Executes, Executed.</p>', NAMES))).toEqual([]);
  });

  it('matches every mention, not only the first', () => {
    expect(keysIn(linkProse('<p>Execute, then Execute again.</p>', NAMES))).toEqual([
      'spell:20662',
      'spell:20662',
    ]);
  });

  it('reads the apostrophe however the markdown printed it', () => {
    for (const text of ["Nature's Grasp", 'Nature’s Grasp', 'Nature&#39;s Grasp']) {
      expect(keysIn(linkProse(`<p>${text} roots.</p>`, NAMES)), text).toEqual(['spell:16813']);
    }
  });

  it('keeps a name ending in punctuation whole', () => {
    expect(keysIn(linkProse('<p>Gnomes get Eureka! for free.</p>', NAMES))).toEqual(['racial:Eureka!']);
  });

  it('colours an item by its quality and opens the item tooltip', () => {
    const html = linkProse('<li>Flask of the Titans before the pull</li>', NAMES);
    expect(html).toContain('data-bis-item="13510"');
    expect(html).toContain('class="gtip gtip-item text-rarity-common"');
  });

  it('leaves headings, links and code alone', () => {
    const html = linkProse(
      '<h3 id="execute">Execute</h3><p><a href="/x">Execute</a> <code>Execute</code> Execute</p>',
      NAMES,
    );
    expect(keysIn(html)).toEqual(['spell:20662']);
    expect(html).toContain('<h3 id="execute">Execute</h3>');
  });

  it('never marks inside a name it already marked', () => {
    const once = linkProse('<p>Blood Fury and Execute.</p>', NAMES);
    expect(linkProse(once, NAMES)).toBe(once);
  });

  it('escapes a key it writes into an attribute', () => {
    const names = new Map<string, ProseTarget>([
      ['Artisan "Precision"', { key: 'racial:Artisan "Precision"' }],
    ]);
    expect(linkProse('<p>Artisan "Precision" helps.</p>', names)).toContain(
      'data-game-tip="racial:Artisan &quot;Precision&quot;"',
    );
  });

  it('passes html through when no name is known', () => {
    expect(linkProse('<p>Nothing here.</p>', new Map())).toBe('<p>Nothing here.</p>');
  });
});
