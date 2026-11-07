// web/src/components/ui/pages-primitives.test.ts
// The pages round's primitives (pages spec 2026-11-08 §3A): Dated (open and folded), DatedGroup,
// CharacterAnswer, the reading Page's `tocFrom`, Prose's table and inline code, Disclosure `eager`.
import { createRawSnippet, type Component } from 'svelte';
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import CharacterAnswer from './CharacterAnswer.svelte';
import Dated from './Dated.svelte';
import DatedGroup from './DatedGroup.svelte';
import Disclosure from './Disclosure.svelte';
import Page from './Page.svelte';
import Prose from './Prose.svelte';

const text = (html: string) => createRawSnippet(() => ({ render: () => html }));

function html<P extends Record<string, unknown>>(component: Component<P>, props: P): string {
  return render(component, { props }).body.replace(/<!--[^>]*-->/g, '');
}

const ENTRY = { anchor: '2026-11-05-guild-events', iso: '2026-11-05', date: '5 Nov', title: 'Guild events.' };

describe('Dated', () => {
  it('is an article anchored by its id, the date its permalink, the title, the body, then the lane', () => {
    const body = html(Dated, {
      ...ENTRY,
      show: 'site',
      tags: text('<span>TAG</span>'),
      lane: text('<p>LANE</p>'),
      children: text('<p>BODY</p>'),
    });
    expect(body).toMatch(
      /<article class="ui-dated ui-dated--o[^"]*" id="2026-11-05-guild-events" data-show="site"/,
    );
    expect(body).toMatch(
      /<a class="ui-dated-d[^"]*" href="#2026-11-05-guild-events"><time datetime="2026-11-05">5 Nov<\/time><\/a>/,
    );
    expect(body.match(/<h3[^>]*>Guild events\.<\/h3>/g)).toHaveLength(1);
    const order = ['TAG', 'Guild events.', 'BODY', 'LANE'].map((part) => body.indexOf(part));
    expect(order).toEqual([...order].sort((a, b) => a - b));
  });

  it('folds into a title row that opens in place (date, title, tags), the body and lane inside', () => {
    const body = html(Dated, {
      ...ENTRY,
      folded: true,
      tags: text('<span>TAG</span>'),
      lane: text('<p>LANE</p>'),
      children: text('<p>BODY</p>'),
    });
    expect(body).toMatch(/<details class="ui-dated ui-dated--f[^"]*" id="2026-11-05-guild-events"/);
    expect(body).toMatch(
      /<summary[^>]*><time[^>]*datetime="2026-11-05">5 Nov<\/time> <h3[^>]*>Guild events\.<\/h3>/,
    );
    expect(body).not.toContain('href="#2026-11-05-guild-events"');
    expect(body.indexOf('</summary>')).toBeLessThan(body.indexOf('BODY'));
    expect(body).not.toMatch(/ open/);
  });

  it('draws no surface', () => {
    expect(html(Dated, { ...ENTRY, children: text('<p></p>') })).not.toMatch(
      /bg-raised|rounded-plate|card-surface/,
    );
  });
});

describe('DatedGroup', () => {
  it('is a section region named for its month, anchored, with its count at the right of the head', () => {
    const body = html(DatedGroup, {
      title: 'November 2026',
      anchor: '2026-11',
      meta: '3 changes',
      children: text('<i>E</i>'),
    });
    expect(body).toMatch(/<section id="2026-11"[^>]*data-region="section-november-2026"/);
    expect(body).toMatch(/<h2[^>]*>November 2026<\/h2> <span class="[^"]*ml-auto[^"]*">3 changes<\/span>/);
    expect(body).toContain('<i>E</i>');
  });

  it('takes a meta snippet and the attributes a page filters it by', () => {
    const body = html(DatedGroup, {
      title: 'October 2026',
      anchor: '2026-10',
      metaContent: text('<b>M</b>'),
      attrs: { 'data-game': '0' },
      children: text(''),
    });
    expect(body).toContain('<b>M</b>');
    expect(body).toContain('data-game="0"');
  });
});

describe('CharacterAnswer', () => {
  it('is reserved at its final size with its label and skeleton, and carries every field the painter fills', () => {
    const body = html(CharacterAnswer, {
      label: 'Your character',
      id: 'nf-answer',
      page: { a: 1 },
      testid: 'ans',
    });
    expect(body).toMatch(
      /id="nf-answer"[^>]*data-signed-in-only[^>]*data-state="loading"[^>]*data-page="\{&quot;a&quot;:1\}"/,
    );
    expect(body).toContain('>Your character</span>');
    expect(body).toContain('data-answer-skeleton');
    for (const field of ['who', 'crest', 'name', 'detail', 'message', 'next', 'links', 'action']) {
      expect(body).toContain(`data-field="${field}"`);
    }
    // the crest comes before the name: it sits at the name's left (ruling SHOULD)
    expect(body.indexOf('data-field="crest"')).toBeLessThan(body.indexOf('data-field="name"'));
  });
});

describe('Page tocFrom', () => {
  it('marks a document whose rail shows from 1024 and names the rail', () => {
    const body = html(Page, {
      layout: 'reading',
      tocFrom: 1024,
      tocLabel: 'On this page',
      toc: text('<ol></ol>'),
      children: text('<p></p>'),
    });
    expect(body).toContain('data-toc-from="1024"');
    expect(body).toMatch(/<nav class="ui-page-toc[^"]*" aria-label="On this page">/);
  });

  it('draws no rail for a document with none, and leaves the guides default unmarked', () => {
    const none = html(Page, { layout: 'reading', tocFrom: 0, toc: text('<ol></ol>'), children: text('') });
    expect(none).toContain('data-toc-from="0"');
    expect(none).not.toContain('ui-page-toc');
    const guides = html(Page, { layout: 'reading', toc: text('<ol></ol>'), children: text('') });
    expect(guides).not.toContain('data-toc-from');
    expect(guides).toContain('aria-label="Contents"');
  });
});

describe('Prose table and code, Disclosure eager', () => {
  it('passes a table and inline code through', () => {
    const body = html(Prose, {
      html: '<table><tr><td>Dec 9</td></tr></table><p><code>fs_session</code></p>',
    });
    expect(body).toContain('<td>Dec 9</td>');
    expect(body).toContain('<code>fs_session</code>');
  });

  it('renders its body closed when eager, for a page with no runtime', () => {
    const lazy = html(Disclosure, { summary: 'On this page', children: text('<i>ROWS</i>') });
    const eager = html(Disclosure, {
      summary: 'On this page',
      meta: '8 sections',
      eager: true,
      children: text('<i>ROWS</i>'),
    });
    expect(lazy).not.toContain('ROWS');
    expect(eager).toContain('ROWS');
    expect(eager).not.toMatch(/<details[^>]* open/);
    expect(eager).toContain('8 sections');
  });
});
