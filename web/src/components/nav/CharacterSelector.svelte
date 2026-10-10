<!-- web/src/components/nav/CharacterSelector.svelte -->
<!-- The header's character selector (design/specs/2026-10-09-nav-character-selector.md),
     replacing the session chip and its account menu: "which character is this site showing
     me, and can I change it?" The closed control shows the current character (crest, full
     name, spec and level; crest alone on the phone bar), or "Sign in". It opens a popover on
     a desktop and a modal sheet on a phone, listing the account's characters; choosing one
     writes the site's one pointer (`lib/character-selector/choose.ts`) and every island that
     reads it follows. Server-rendered it is a link to /account#characters (`csel-pre`) that
     Base.astro's pre-paint script has already filled from the stored pointer, so the box, the
     crest and the name are where they will be when the island swaps in the button. -->
<script lang="ts">
  import { tick } from 'svelte';
  import { battlenetStartUrl } from '../../lib/account/api';
  import { browserChooseDeps, chooseCharacter, choosePointer } from '../../lib/character-selector/choose';
  import { selectorCopy } from '../../lib/character-selector/copy';
  import { OPEN_SELECTOR_EVENT } from '../../lib/character-selector/config';
  import {
    closedAriaLabel,
    closedSecondLine,
    markTooltip,
    ringMarkOf,
    type SelectorRow,
  } from '../../lib/character-selector/rows';
  import { createSelectorState } from '../../lib/character-selector/selector-state.svelte';
  import { classColorVar } from '../../lib/report/format';
  import SelectorCrest from './SelectorCrest.svelte';
  import SelectorMenuRow from './SelectorMenuRow.svelte';
  import SelectorPanel from './SelectorPanel.svelte';

  const SHEET_QUERY = '(max-width: 1023px)';
  const TIP_ID = 'character-selector-tip';

  const selector = createSelectorState();

  let open = $state(false);
  let sheet = $state(false);
  let retryingKey = $state<string | null>(null);
  let announcement = $state('');
  let root: HTMLElement | undefined = $state();
  let trigger: HTMLButtonElement | undefined = $state();
  let panel: SelectorPanel | undefined = $state();

  const model = $derived(selector.model);
  const current = $derived(model.current);
  const mark = $derived(ringMarkOf(model));

  $effect(() => {
    const query = window.matchMedia(SHEET_QUERY);
    sheet = query.matches;
    const onChange = (event: MediaQueryListEvent): void => {
      sheet = event.matches;
    };
    query.addEventListener('change', onChange);
    return () => query.removeEventListener('change', onChange);
  });

  $effect(() => {
    window.addEventListener(OPEN_SELECTOR_EVENT, openSelector);
    return () => window.removeEventListener(OPEN_SELECTOR_EVENT, openSelector);
  });

  $effect(() => {
    if (!open) return;
    const onPointerDown = (event: PointerEvent): void => {
      if (!sheet && event.target instanceof Node && root?.contains(event.target) !== true) close(false);
    };
    const onKeydown = (event: KeyboardEvent): void => {
      if (event.key === 'Escape') close(true);
    };
    document.addEventListener('pointerdown', onPointerDown);
    document.addEventListener('keydown', onKeydown);
    return () => {
      document.removeEventListener('pointerdown', onPointerDown);
      document.removeEventListener('keydown', onKeydown);
    };
  });

  $effect(() => {
    if (!open || !sheet) return;
    const previous = document.body.style.overflow;
    document.body.style.overflow = 'hidden';
    return () => {
      document.body.style.overflow = previous;
    };
  });

  function openSelector(): void {
    selector.refresh();
    open = true;
    void tick().then(() => panel?.focusInitial());
  }

  function close(returnFocus: boolean): void {
    open = false;
    if (returnFocus) void tick().then(() => trigger?.focus());
  }

  function openFromMenuRow(): void {
    window.dispatchEvent(new Event(OPEN_SELECTOR_EVENT));
  }

  function toggle(): void {
    if (open) close(false);
    else openSelector();
  }

  function onTriggerKeydown(event: KeyboardEvent): void {
    if (event.key !== 'ArrowDown' || open) return;
    event.preventDefault();
    openSelector();
  }

  function onFocusOut(event: FocusEvent): void {
    if (!open || sheet || !(event.relatedTarget instanceof Node)) return;
    if (root?.contains(event.relatedTarget) !== true) close(false);
  }

  function onChoose(row: SelectorRow): void {
    if (row.kind === 'account' && row.character !== undefined && !row.current) {
      chooseCharacter(row.character, browserChooseDeps());
      announcement = selectorCopy.nowUsing(row.name);
    }
    close(true);
  }

  async function onPaste(input: string): Promise<boolean> {
    const { resolvePastedExport } = await import('../../lib/character-selector/paste');
    const accountCharacters = model.rows.flatMap((row) =>
      row.character === undefined ? [] : [row.character],
    );
    const result = resolvePastedExport(input, accountCharacters, new Date());
    if (!result.ok) return false;
    choosePointer(result.pointer, result.account?.key ?? null, browserChooseDeps());
    announcement = selectorCopy.nowUsing(result.account?.name ?? result.pointer.label.split(' · ')[0]);
    close(true);
    return true;
  }

  function onRetry(row: SelectorRow): void {
    retryingKey = row.key;
    window.location.assign(battlenetStartUrl('/account?refreshed=1'));
  }
</script>

<div class="csel" data-testid="selector-root" bind:this={root} onfocusout={onFocusOut}>
  {#if selector.closed === 'pre'}
    <a
      class="csel-btn csel-pre"
      href={selectorCopy.manageHref}
      aria-label={selectorCopy.dialogLabel}
      data-testid="character-selector-pre"
    >
      <span class="csel-pre-char">
        <span class="csel-crestbox">
          <span class="csel-crest csel-pre-crest"></span>
          <span class="csel-badge"></span>
        </span>
        <span class="csel-name-bar" data-csel-name></span>
        <span class="csel-text">
          <span class="csel-name" data-csel-name></span>
          <span class="csel-sub" data-csel-sub></span>
          <span class="csel-sub-short" data-csel-sub-short></span>
        </span>
        <svg
          class="csel-chev"
          width="14"
          height="14"
          viewBox="0 0 24 24"
          fill="none"
          stroke="currentColor"
          stroke-width="2.4"
          stroke-linecap="round"
          stroke-linejoin="round"
          aria-hidden="true"
        >
          <path d="M6 9l6 6 6-6"></path>
        </svg>
      </span>
      <span class="csel-pre-out">
        <span class="csel-ghost">
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
        <span class="csel-text">
          <span class="csel-name">{selectorCopy.closedSignedOut}</span>
          <span class="csel-sub">{selectorCopy.closedSignedOutSub}</span>
        </span>
      </span>
      <span class="csel-pre-skel">
        <span class="csel-skeleton csel-skeleton-ring"></span>
        <span class="csel-text">
          <span class="csel-skeleton" style="width:96px;height:13px"></span>
          <span class="csel-skeleton" style="width:130px;height:11px;margin-top:5px"></span>
        </span>
      </span>
    </a>
  {:else if selector.closed === 'character' && current !== null}
    <button
      type="button"
      class="csel-btn"
      aria-haspopup="dialog"
      aria-expanded={open}
      aria-label={closedAriaLabel(current, mark)}
      aria-describedby={mark === null ? undefined : TIP_ID}
      title={current.name}
      bind:this={trigger}
      onclick={toggle}
      onkeydown={onTriggerKeydown}
      data-testid="character-selector"
    >
      <span class="csel-crestbox">
        <SelectorCrest classSlug={current.classSlug} size={32} />
        {#if mark !== null}<i class="csel-mark" data-mark={mark} data-testid="selector-mark"></i>{/if}
        <span class="csel-badge">
          <svg
            width="8"
            height="8"
            viewBox="0 0 24 24"
            fill="none"
            stroke="currentColor"
            stroke-width="4"
            stroke-linecap="round"
            stroke-linejoin="round"
            aria-hidden="true"
          >
            <path d="M6 9l6 6 6-6"></path>
          </svg>
        </span>
      </span>
      <span class="csel-name-bar" style:color={classColorVar(current.classSlug)}>{current.name}</span>
      <span class="csel-text">
        <span class="csel-name" style:color={classColorVar(current.classSlug)} data-testid="selector-name">
          {current.name}
        </span>
        <span class="csel-sub" data-testid="selector-sub">{closedSecondLine(current, true)}</span>
        <span class="csel-sub-short">{closedSecondLine(current, false)}</span>
      </span>
      <svg
        class="csel-chev"
        width="14"
        height="14"
        viewBox="0 0 24 24"
        fill="none"
        stroke="currentColor"
        stroke-width="2.4"
        stroke-linecap="round"
        stroke-linejoin="round"
        aria-hidden="true"
      >
        <path d="M6 9l6 6 6-6"></path>
      </svg>
    </button>
    {#if mark !== null}
      <span class="csel-tip" role="tooltip" id={TIP_ID}>{markTooltip(mark, current)}</span>
    {/if}
  {:else}
    <button
      type="button"
      class="csel-btn is-signed-out"
      aria-haspopup="dialog"
      aria-expanded={open}
      aria-label={selectorCopy.signedOutAria}
      bind:this={trigger}
      onclick={toggle}
      onkeydown={onTriggerKeydown}
      data-testid="character-selector"
    >
      <span class="csel-ghost">
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
      <span class="csel-text">
        <span class="csel-name">{selectorCopy.closedSignedOut}</span>
        <span class="csel-sub">{selectorCopy.closedSignedOutSub}</span>
      </span>
      <svg
        class="csel-chev"
        width="14"
        height="14"
        viewBox="0 0 24 24"
        fill="none"
        stroke="currentColor"
        stroke-width="2.4"
        stroke-linecap="round"
        stroke-linejoin="round"
        aria-hidden="true"
      >
        <path d="M6 9l6 6 6-6"></path>
      </svg>
    </button>
  {/if}

  <SelectorMenuRow closed={selector.closed} {current} onopen={openFromMenuRow} />

  {#if open}
    {#if sheet}
      <div class="csel-scrim" role="presentation" onclick={() => close(true)}></div>
    {/if}
    <SelectorPanel
      bind:this={panel}
      {model}
      loadStatus={selector.loadStatus}
      {sheet}
      {retryingKey}
      onchoose={onChoose}
      onpaste={onPaste}
      onretry={onRetry}
      ontryagain={() => selector.reload()}
      onclose={() => close(true)}
    />
  {/if}
  <div class="sr-only" aria-live="polite" role="status" data-testid="selector-live">{announcement}</div>
</div>
