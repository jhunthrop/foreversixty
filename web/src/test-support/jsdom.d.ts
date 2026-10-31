// jsdom ships no type declarations; tests that parse rendered Astro markup (Header.test.ts)
// use only this slice of it.
declare module 'jsdom' {
  export class JSDOM {
    constructor(html?: string, options?: { url?: string });
    readonly window: { readonly document: Document };
  }
}
