import { describe, expect, it } from 'vitest';
import { FakeHTMLRewriter } from './html-rewriter';

const HEAD = `<!doctype html><html><head><title>Forever Sixty</title>
<meta name="description" content="old" data-og="description" />
<link rel="canonical" href="https://foreversixty.gg/reports" data-og="canonical" />
<meta property="og:title" content="old" data-og="og-title" />
</head><body><div id="report"></div></body></html>`;

async function transform(rewriter: FakeHTMLRewriter, html: string): Promise<string> {
  return rewriter.transform(new Response(html, { headers: { 'content-type': 'text/html' } })).text();
}

describe('FakeHTMLRewriter', () => {
  it('replaces an element’s inner content', async () => {
    const out = await transform(
      new FakeHTMLRewriter().on('title', { element: (el) => el.setInnerContent('A raid night') }),
      HEAD,
    );
    expect(out).toContain('<title>A raid night</title>');
  });

  it('sets an attribute on the element an attribute selector picks out', async () => {
    const out = await transform(
      new FakeHTMLRewriter().on('meta[data-og="description"]', {
        element: (el) => el.setAttribute('content', 'three fights'),
      }),
      HEAD,
    );
    expect(out).toContain('content="three fights"');
    expect(out).toContain('<meta property="og:title" content="old" data-og="og-title" />');
  });

  it('reads an attribute back and leaves untouched elements byte for byte', async () => {
    let seen: string | null = null;
    const out = await transform(
      new FakeHTMLRewriter().on('link[data-og="canonical"]', {
        element: (el) => {
          seen = el.getAttribute('rel');
          el.setAttribute('href', 'https://foreversixty.gg/reports/fixture2abcd');
        },
      }),
      HEAD,
    );
    expect(seen).toBe('canonical');
    expect(out).toContain('href="https://foreversixty.gg/reports/fixture2abcd"');
    expect(out).toContain('<div id="report"></div>');
  });

  it('escapes what it writes', async () => {
    const out = await transform(
      new FakeHTMLRewriter()
        .on('title', { element: (el) => el.setInnerContent('<script>x</script>') })
        .on('meta[data-og="og-title"]', { element: (el) => el.setAttribute('content', 'a "quoted" & raid') }),
      HEAD,
    );
    expect(out).toContain('<title>&lt;script&gt;x&lt;/script&gt;</title>');
    expect(out).toContain('content="a &quot;quoted&quot; &amp; raid"');
    expect(out).not.toContain('<script>x');
  });

  it('keeps the response status and headers', async () => {
    const response = new FakeHTMLRewriter()
      .on('title', { element: (el) => el.setInnerContent('x') })
      .transform(
        new Response(HEAD, { status: 200, headers: { 'content-type': 'text/html', 'x-kept': '1' } }),
      );
    expect(response.headers.get('x-kept')).toBe('1');
    expect(response.status).toBe(200);
  });

  it('leaves the document alone when nothing matches', async () => {
    const out = await transform(new FakeHTMLRewriter().on('h1', { element: () => {} }), HEAD);
    expect(out).toBe(HEAD);
  });

  it('refuses a selector it does not implement, rather than silently matching nothing', () => {
    expect(() => new FakeHTMLRewriter().on('head > title', { element: () => {} })).toThrow(
      'unsupported selector',
    );
  });
});
