<!-- web/src/components/nav/SelectorRow.svelte -->
<!-- One character in the selector's list (spec 3.C): crest, name in class colour with a
     faction mark, `{Spec} {Class} · {level} · {Realm}`, then source and age. Gold marks the
     current row only; stale and failed use text colour. The row button is one roving tab
     stop; Retry on a failed row is its own sibling button, never nested. -->
<script lang="ts">
  import { FACTION_MARK_SIZE, factionLogoSrc, factionName } from '../../lib/faction-mark';
  import { selectorCopy } from '../../lib/character-selector/copy';
  import { rowLineThree, rowLineTwo, type SelectorRow } from '../../lib/character-selector/rows';
  import { classColorVar } from '../../lib/report/format';
  import SelectorCrest from './SelectorCrest.svelte';

  let {
    row,
    tabindex,
    retrying = false,
    onchoose,
    onretry,
  }: {
    row: SelectorRow;
    tabindex: 0 | -1;
    retrying?: boolean;
    onchoose: (row: SelectorRow) => void;
    onretry: (row: SelectorRow) => void;
  } = $props();

  const lineThree = $derived(rowLineThree(row));
</script>

<li
  class="csel-row"
  class:is-current={row.current}
  style:--c={classColorVar(row.classSlug)}
  data-testid={`selector-row-${row.key}`}
>
  <button
    type="button"
    class="csel-row-main"
    {tabindex}
    data-row-main
    data-name={row.name}
    aria-current={row.current ? 'true' : undefined}
    onclick={() => onchoose(row)}
  >
    <SelectorCrest classSlug={row.classSlug} size={36} cssClass="csel-row-crest" lazy />
    <span class="csel-row-text">
      <span class="csel-r1">
        <span>{row.name}</span>
        {#if row.faction !== undefined}<img
            class="csel-faction"
            src={factionLogoSrc(row.faction)}
            alt={factionName(row.faction)}
            title={factionName(row.faction)}
            width={FACTION_MARK_SIZE}
            height={FACTION_MARK_SIZE}
            loading="lazy"
            decoding="async"
            data-testid={`faction-mark-${row.faction}`}
          />{/if}
      </span>
      {#if rowLineTwo(row) !== ''}<span class="csel-r2">{rowLineTwo(row)}</span>{/if}
      {#if retrying}
        <span class="csel-r3">{selectorCopy.refreshing}</span>
      {:else if lineThree.text !== ''}
        <span class="csel-r3" class:is-failed={lineThree.failed}>
          {lineThree.text}{#if lineThree.tail !== ''}&nbsp;·&nbsp;<span class="is-stale"
              >{lineThree.tail}</span
            >{/if}
        </span>
      {/if}
    </span>
    {#if row.current}
      <svg
        class="csel-check"
        width="18"
        height="18"
        viewBox="0 0 24 24"
        fill="none"
        stroke="currentColor"
        stroke-width="2.6"
        stroke-linecap="round"
        stroke-linejoin="round"
        aria-hidden="true"
      >
        <path d="M5 12.5l4.5 4.5L19 7.5"></path>
      </svg>
    {/if}
  </button>
  {#if row.status === 'failed'}
    <button
      type="button"
      class="csel-retry"
      disabled={retrying}
      data-testid={`selector-retry-${row.key}`}
      onclick={() => onretry(row)}
    >
      {retrying ? selectorCopy.retrying : selectorCopy.retry}
    </button>
  {/if}
</li>
