// A stand-in for the Workers runtime's HTMLRewriter, for vitest only. Nothing in src/
// imports it except the Worker's tests: it is tree-shaken out of every bundle because no
// production module references it.
//
// It implements exactly what src/worker.ts uses -- the selector forms `tag` and
// `tag[attr="value"]`, and getAttribute / setAttribute / removeAttribute /
// setInnerContent -- and throws on anything else, so a Worker change that reaches for more
// of the real API fails loudly here instead of passing against a fake that quietly did
// nothing. @cloudflare/vitest-pool-workers, which would give the real thing, currently
// peers on vitest ^4.1.0 while this project is on ^5.

const TAG = /<([a-zA-Z][\w-]*)((?:\s+[^\s=/>]+(?:\s*=\s*(?:"[^"]*"|'[^']*'|[^\s"'=<>`]+))?)*)\s*(\/?)>/g;
const ATTR = /([^\s=/>]+)(?:\s*=\s*(?:"([^"]*)"|'([^']*)'|([^\s"'=<>`]+)))?/g;
const SELECTOR = /^([a-zA-Z][\w-]*)(?:\[([^\]=]+)="([^"]*)"\])?$/;

function parseAttributes(raw: string): Map<string, string> {
  const attributes = new Map<string, string>();
  ATTR.lastIndex = 0;
  let match = ATTR.exec(raw);
  while (match !== null) {
    attributes.set(match[1], match[2] ?? match[3] ?? match[4] ?? '');
    match = ATTR.exec(raw);
  }
  return attributes;
}

function escapeAttribute(value: string): string {
  return value.replace(/&/g, '&amp;').replace(/"/g, '&quot;').replace(/</g, '&lt;');
}

function escapeText(value: string): string {
  return value.replace(/&/g, '&amp;').replace(/</g, '&lt;').replace(/>/g, '&gt;');
}

export class FakeElement {
  /** Set by setInnerContent; null means the element's children are untouched. */
  innerContent: string | null = null;

  constructor(
    readonly tagName: string,
    readonly attributes: Map<string, string>,
  ) {}

  getAttribute(name: string): string | null {
    return this.attributes.get(name) ?? null;
  }

  setAttribute(name: string, value: string): void {
    this.attributes.set(name, value);
  }

  removeAttribute(name: string): void {
    this.attributes.delete(name);
  }

  setInnerContent(text: string): void {
    this.innerContent = text;
  }
}

export interface FakeElementHandlers {
  element(element: FakeElement): void;
}

interface Rule {
  tag: string;
  attribute?: string;
  value?: string;
  handlers: FakeElementHandlers;
}

export class FakeHTMLRewriter {
  private readonly rules: Rule[] = [];

  on(selector: string, handlers: FakeElementHandlers): this {
    const parsed = SELECTOR.exec(selector);
    if (parsed === null) {
      throw new Error(`FakeHTMLRewriter: unsupported selector ${selector}`);
    }
    this.rules.push({
      tag: parsed[1].toLowerCase(),
      attribute: parsed[2],
      value: parsed[3],
      handlers,
    });
    return this;
  }

  transform(response: Response): Response {
    const rewritten = response.text().then((html) => this.rewrite(html));
    const body = new ReadableStream<Uint8Array>({
      async start(controller) {
        controller.enqueue(new TextEncoder().encode(await rewritten));
        controller.close();
      },
    });
    return new Response(body, {
      status: response.status,
      statusText: response.statusText,
      headers: response.headers,
    });
  }

  private rewrite(html: string): string {
    let out = '';
    let cursor = 0;
    TAG.lastIndex = 0;
    let match = TAG.exec(html);
    while (match !== null) {
      const [whole, tagName, rawAttributes, selfClosing] = match;
      const lower = tagName.toLowerCase();
      const attributes = parseAttributes(rawAttributes);
      const matching = this.rules.filter(
        (rule) =>
          rule.tag === lower &&
          (rule.attribute === undefined || attributes.get(rule.attribute) === rule.value),
      );
      if (matching.length > 0) {
        const element = new FakeElement(lower, attributes);
        for (const rule of matching) rule.handlers.element(element);

        let rebuilt = `<${tagName}`;
        for (const [name, value] of element.attributes) {
          rebuilt += value === '' ? ` ${name}` : ` ${name}="${escapeAttribute(value)}"`;
        }
        rebuilt += selfClosing === '/' ? ' />' : '>';

        out += html.slice(cursor, match.index) + rebuilt;
        cursor = match.index + whole.length;

        if (element.innerContent !== null && selfClosing !== '/') {
          const close = html.indexOf(`</${tagName}`, cursor);
          if (close !== -1) {
            out += escapeText(element.innerContent);
            cursor = close;
          }
        }
      }
      match = TAG.exec(html);
    }
    return out + html.slice(cursor);
  }
}
