<!-- web/src/components/ui/Prose.svelte -->
<!-- Rendered guide markdown's typography (guides spec 2026-11-08 §3A): 16/26 (15/24 `lane`),
     list markers, h3 18/24, `strong`, the game-tooltip names (dotted underline, --color-strong
     500; items keep their quality colour), `glabel` (a bold-only paragraph as a 12 px caps label),
     and the modifiers `ruled` (a fact list with gold marks and hairlines), `cols` (two columns from
     1024), `plus` / `minus` (Strengths and Weaknesses marks). The html is the build's own render
     of committed markdown (lib/guides/markdown.ts), never user input. -->
<script lang="ts">
  type ProseMod = 'lane' | 'ruled' | 'cols' | 'plus' | 'minus';

  let {
    html,
    mods = [],
    testid,
  }: {
    html: string;
    mods?: readonly ProseMod[];
    testid?: string;
  } = $props();
</script>

<!-- eslint-disable-next-line svelte/no-at-html-tags -- the build's own markdown render -->
<div class={['ui-prose', ...mods.map((mod) => `ui-prose--${mod}`)]} data-testid={testid}>{@html html}</div>

<style>
  .ui-prose {
    font-size: 16px;
    line-height: 26px;
    color: var(--color-text);
    max-width: var(--measure-reading);
    overflow-wrap: break-word;
  }
  .ui-prose--lane {
    font-size: 15px;
    line-height: 24px;
  }
  .ui-prose :global(p) {
    margin: 0 0 14px;
  }
  .ui-prose > :global(:last-child) {
    margin-bottom: 0;
  }
  .ui-prose :global(ul),
  .ui-prose :global(ol) {
    margin: 0 0 14px;
    padding-left: 20px;
    list-style: disc;
  }
  .ui-prose :global(ol) {
    list-style: decimal;
  }
  .ui-prose :global(li) {
    margin: 0 0 8px;
  }
  .ui-prose :global(li::marker) {
    color: var(--color-muted);
  }
  .ui-prose :global(h3) {
    margin: 28px 0 10px;
    font-size: 18px;
    line-height: 24px;
    font-weight: 700;
    color: var(--color-strong);
  }
  .ui-prose > :global(h3:first-child) {
    margin-top: 0;
  }
  .ui-prose :global(strong) {
    color: var(--color-strong);
    font-weight: 700;
  }
  .ui-prose :global(a) {
    font-weight: 600;
    color: var(--color-gold);
  }
  .ui-prose :global(.gtip) {
    position: relative;
    cursor: pointer;
    text-decoration: underline dotted;
    text-decoration-color: color-mix(in srgb, currentColor 40%, transparent);
    text-underline-offset: 3px;
  }
  .ui-prose :global(span.gtip:not(.gtip-item)) {
    color: var(--color-strong);
    font-weight: 500;
  }
  .ui-prose :global(.gtip:focus-visible) {
    outline: 2px solid var(--color-gold);
    outline-offset: 2px;
    border-radius: 3px;
  }
  .ui-prose :global(.gtip [role='tooltip']) {
    color: var(--color-text);
    text-decoration: none;
    cursor: default;
  }
  .ui-prose :global(.glabel) {
    margin: 18px 0 4px;
    font-size: 12px;
    line-height: 16px;
    font-weight: 700;
    letter-spacing: 0.12em;
    text-transform: uppercase;
    color: var(--color-muted);
  }
  .ui-prose :global(.glabel:first-child) {
    margin-top: 0;
  }
  .ui-prose--ruled :global(ul) {
    list-style: none;
    padding: 0;
    margin: 0;
  }
  .ui-prose--ruled :global(li) {
    position: relative;
    margin: 0;
    padding: 12px 0 12px 20px;
    border-bottom: 1px solid var(--color-line-soft);
    break-inside: avoid;
  }
  .ui-prose--ruled :global(li::before) {
    content: '';
    position: absolute;
    left: 0;
    top: 21px;
    width: 6px;
    height: 6px;
    border-radius: 50%;
    background: var(--color-gold);
  }
  .ui-prose--plus :global(ul),
  .ui-prose--minus :global(ul) {
    list-style: none;
    padding: 0;
  }
  .ui-prose--plus :global(li),
  .ui-prose--minus :global(li) {
    position: relative;
    padding-left: 20px;
  }
  .ui-prose--plus :global(li::before),
  .ui-prose--minus :global(li::before) {
    position: absolute;
    left: 0;
    font-weight: 700;
  }
  .ui-prose--plus :global(li::before) {
    content: '+';
    color: var(--color-kill);
  }
  .ui-prose--minus :global(li::before) {
    content: '\2212';
    color: var(--color-wipe);
  }
  /* The ruled table and inline code (pages spec 2026-11-08 §3A): 12 px caps heads over a hairline,
     15/22 cells, the first column 600 --strong; below 768 each row stacks, heads hidden. Code in
     the one mono face 600 14/20 --strong, no box. */
  .ui-prose :global(table) {
    width: 100%;
    border-collapse: collapse;
    margin: 0 0 18px;
    font-size: 15px;
    line-height: 22px;
  }
  .ui-prose :global(th) {
    text-align: left;
    font-size: 12px;
    line-height: 16px;
    font-weight: 700;
    letter-spacing: 0.12em;
    text-transform: uppercase;
    color: var(--color-muted);
    padding: 0 16px 8px 0;
    border-bottom: 1px solid var(--color-line-soft);
  }
  .ui-prose :global(td) {
    padding: 10px 16px 10px 0;
    border-bottom: 1px solid var(--color-line-soft);
    vertical-align: top;
  }
  .ui-prose :global(td:first-child) {
    width: 1%;
    white-space: nowrap;
    font-weight: 600;
    color: var(--color-strong);
  }
  .ui-prose :global(code) {
    font-family: var(--font-mono);
    font-weight: 600;
    font-size: 14px;
    line-height: 20px;
    color: var(--color-strong);
  }
  @media (max-width: 767px) {
    .ui-prose :global(thead) {
      display: none;
    }
    .ui-prose :global(tr) {
      display: grid;
      gap: 2px;
      padding: 10px 0;
      border-bottom: 1px solid var(--color-line-soft);
    }
    .ui-prose :global(td) {
      padding: 0;
      border: 0;
    }
    .ui-prose :global(td:first-child) {
      width: auto;
      white-space: normal;
    }
  }
  @media (min-width: 1024px) {
    .ui-prose--cols {
      max-width: none;
    }
    .ui-prose--cols :global(ul) {
      column-count: 2;
      column-gap: 48px;
    }
  }
</style>
