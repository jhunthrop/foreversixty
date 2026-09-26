// web/src/components/guides/GuideBuildTree.test.ts
// Server-rendered only (no fetch happens during SSR -- the component's own $effect runs
// client-side after hydration), so this asserts the loading state's shape: the same
// `render(Component, {props})` -> `.body` pattern every other planner/guide component test
// in this codebase already uses (see ImportBox.test.ts, TalentCell.test.ts).
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import GuideBuildTree from './GuideBuildTree.svelte';

describe('GuideBuildTree', () => {
  it('server-renders the loading skeleton (no fetch happens before hydration)', () => {
    const { body } = render(GuideBuildTree, { props: { code: 'FS1:1.60.1.69893:warrior:human:1/2/3:' } });
    expect(body).toContain('data-testid="guide-tree-skeleton"');
  });
});
