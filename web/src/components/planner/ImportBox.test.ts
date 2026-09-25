// web/src/components/planner/ImportBox.test.ts
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
// Same fixture and pattern src/lib/planner/live-dps.test.ts already uses: the checked-in
// warrior talent file under src/fixtures/planner/, indexed with the planner's own
// indexTalents -- not a hand-built TalentIndex, and not the generated public/data/<build>/
// one, whose content is whatever npm run sync last pruned it to.
import fixtureTalents from '../../fixtures/planner/talents/warrior.json';
import { currentCharacterCopy } from '../../lib/current-character-copy';
import { indexTalents } from '../../lib/planner/rules';
import type { TalentFile } from '../../lib/planner/types';
import ImportBox from './ImportBox.svelte';

const talents = indexTalents(fixtureTalents as TalentFile);

describe('ImportBox', () => {
  it('links to /addon, with the shared "get the addon" copy', () => {
    const { body } = render(ImportBox, { props: { talents, activeBuild: '1', onimport: () => {} } });
    expect(body).toContain('href="/setup"');
    expect(body).toContain(currentCharacterCopy.getTheAddon);
  });

  it('gives the /addon link a 44px hit target', () => {
    const { body } = render(ImportBox, { props: { talents, activeBuild: '1', onimport: () => {} } });
    const match = /<a[^>]*href="\/setup"[^>]*>/.exec(body);
    if (match === null) throw new Error('no /addon anchor rendered');
    expect(match[0]).toContain('min-h-11');
  });
});
