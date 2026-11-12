<!-- web/src/components/nav/SelectorMenuRow.svelte -->
<!-- The phone Menu's first row (spec section 7): the current character -- crest, the full name
     in class colour, `{Spec} {Class} · {level}`, a chevron -- or, signed out, "Sign in / or
     paste an export". The bar shows only the crest below 560px, so this row is where the
     full name lives there. Tapping it opens the selector's sheet. Server-rendered it
     carries the same pre-paint hooks as the bar's link (`csel-pre`), so the row is filled
     before the island hydrates. It is rendered by CharacterSelector.svelte (one island, one
     state) and CSS places it in the Menu beneath the bar. -->
<script lang="ts">
  import CharacterName from '../ui/CharacterName.svelte';
  import Icon from '../ui/Icon.svelte';
  import { ICON_ACCOUNT, ICON_CHEVRON_DOWN } from '../../lib/ui/ui-icons';
  import { selectorCopy } from '../../lib/character-selector/copy';
  import { closedSecondLine, type SelectorRow } from '../../lib/character-selector/rows';
  import type { ClosedKind } from '../../lib/character-selector/selector-state.svelte';
  import Crest from '../ui/Crest.svelte';

  let {
    closed,
    current,
    accountName = '',
    onopen,
  }: { closed: ClosedKind; current: SelectorRow | null; accountName?: string; onopen: () => void } = $props();
</script>

{#snippet chevron()}
  <Icon icon={ICON_CHEVRON_DOWN} size={14} class="text-muted" />
{/snippet}

{#snippet ghost()}
  <span class="csel-ghost" style="width:36px;height:36px">
    <Icon icon={ICON_ACCOUNT} size={20} />
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
      <span class="crest csel-crest csel-pre-crest" style="width:36px;height:36px"></span>
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
    <Crest cls={current.classSlug} size={36} eager class="csel-crest" />
    <span class="csel-menu-text">
      <CharacterName name={current.name} cls={current.classSlug} weight="inherit" class="csel-menu-name" />
      <span class="csel-menu-sub">{closedSecondLine(current, true)}</span>
    </span>
    {@render chevron()}
  </button>
{:else if closed === 'account'}
  <button
    type="button"
    class="site-menu-character"
    data-testid="menu-character"
    aria-haspopup="dialog"
    onclick={onopen}
  >
    {@render ghost()}
    <span class="csel-menu-text">
      <span class="csel-menu-name">{accountName}</span>
      <span class="csel-menu-sub">{selectorCopy.accountSub}</span>
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
