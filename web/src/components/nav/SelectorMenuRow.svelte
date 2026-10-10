<!-- web/src/components/nav/SelectorMenuRow.svelte -->
<!-- The phone Menu's first row (spec section 7): the current character -- crest, the full name
     in class colour, `{Spec} {Class} · {level}`, a chevron -- or, signed out, "Sign in / or
     paste an export". The bar shows only the crest below 560px, so this row is where the
     full name lives there. Tapping it opens the selector's sheet. Server-rendered it
     carries the same pre-paint hooks as the bar's link (`csel-pre`), so the row is filled
     before the island hydrates. It is rendered by CharacterSelector.svelte (one island, one
     state) and CSS places it in the Menu beneath the bar. -->
<script lang="ts">
  import { selectorCopy } from '../../lib/character-selector/copy';
  import { closedSecondLine, type SelectorRow } from '../../lib/character-selector/rows';
  import type { ClosedKind } from '../../lib/character-selector/selector-state.svelte';
  import { classColorVar } from '../../lib/report/format';
  import SelectorCrest from './SelectorCrest.svelte';

  let { closed, current, onopen }: { closed: ClosedKind; current: SelectorRow | null; onopen: () => void } =
    $props();
</script>

{#snippet chevron()}
  <svg
    width="14"
    height="14"
    viewBox="0 0 24 24"
    fill="none"
    stroke="currentColor"
    stroke-width="2.4"
    stroke-linecap="round"
    stroke-linejoin="round"
    aria-hidden="true"
    style="color: var(--color-muted)"
  >
    <path d="M6 9l6 6 6-6"></path>
  </svg>
{/snippet}

{#snippet ghost()}
  <span class="csel-ghost" style="width:36px;height:36px">
    <svg
      width="16"
      height="16"
      viewBox="0 0 24 24"
      fill="none"
      stroke="currentColor"
      stroke-width="2"
      stroke-linecap="round"
      aria-hidden="true"
    >
      <circle cx="12" cy="8" r="4"></circle>
      <path d="M4 21c0-4 3.6-7 8-7s8 3 8 7"></path>
    </svg>
  </span>
{/snippet}

{#if closed === 'pre'}
  <button
    type="button"
    class="site-menu-character csel-pre"
    data-testid="menu-character"
    aria-label={selectorCopy.dialogLabel}
    onclick={onopen}
  >
    <span class="csel-pre-char">
      <span class="csel-crest csel-pre-crest" style="width:36px;height:36px"></span>
      <span class="csel-menu-text">
        <span class="csel-menu-name" data-csel-name></span>
        <span class="csel-menu-sub" data-csel-sub></span>
      </span>
      {@render chevron()}
    </span>
    <span class="csel-pre-out">
      {@render ghost()}
      <span class="csel-menu-text">
        <span class="csel-menu-name">{selectorCopy.closedSignedOut}</span>
        <span class="csel-menu-sub">{selectorCopy.menuSignInSub}</span>
      </span>
      {@render chevron()}
    </span>
    <span class="csel-pre-skel">
      <span class="csel-skeleton csel-skeleton-ring" style="width:36px;height:36px"></span>
      <span class="csel-menu-text">
        <span class="csel-skeleton" style="width:110px;height:14px"></span>
      </span>
      <span></span>
    </span>
  </button>
{:else if closed === 'character' && current !== null}
  <button
    type="button"
    class="site-menu-character"
    data-testid="menu-character"
    aria-haspopup="dialog"
    onclick={onopen}
  >
    <SelectorCrest classSlug={current.classSlug} size={36} />
    <span class="csel-menu-text">
      <span class="csel-menu-name" style:color={classColorVar(current.classSlug)}>{current.name}</span>
      <span class="csel-menu-sub">{closedSecondLine(current, true)}</span>
    </span>
    {@render chevron()}
  </button>
{:else}
  <button
    type="button"
    class="site-menu-character"
    data-testid="menu-character"
    aria-haspopup="dialog"
    onclick={onopen}
  >
    {@render ghost()}
    <span class="csel-menu-text">
      <span class="csel-menu-name">{selectorCopy.closedSignedOut}</span>
      <span class="csel-menu-sub">{selectorCopy.menuSignInSub}</span>
    </span>
    {@render chevron()}
  </button>
{/if}
