// web/src/components/home/HomeSignedIn.test.ts
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import { homeCopy } from '../../lib/home/home-view';
import HomeSignedIn from './HomeSignedIn.svelte';

describe('HomeSignedIn', () => {
  it('holds its place with a skeleton head that is no second h1 beside the signed-out page', () => {
    const { body } = render(HomeSignedIn, { props: { floor: [], raidsOpen: null, raids: [] } });
    expect(body).toContain('data-testid="home-head-loading"');
    expect(body).toContain(homeCopy.loadingTitle);
    expect(body).not.toMatch(/<h1[\s>]/);
  });
});
