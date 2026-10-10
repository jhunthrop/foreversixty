<!-- web/src/components/nav/SelectorPanel.svelte -->
<!-- What the selector opens (spec 3.C to 3.F): the heading, the character list, and the
     action rows -- the popover on desktop and the sheet on a phone are the same content. It is
     a pure render of the model it is given plus the callbacks; the island owns open/close,
     focus return and the pointer write. Keyboard inside the list: ArrowUp/Down, Home/End and a
     letter move between rows (one roving tab stop); Enter or Space chooses. -->
<script lang="ts">
  import { selectorCopy } from '../../lib/character-selector/copy';
  import { FILTER_AFTER_CHARACTERS, VISIBLE_ROWS } from '../../lib/character-selector/config';
  import { filterRows, type SelectorModel, type SelectorRow as Row } from '../../lib/character-selector/rows';
  import type { LoadStatus } from '../../lib/character-selector/selector-state.svelte';
  import SelectorPaste from './SelectorPaste.svelte';
  import SelectorRow from './SelectorRow.svelte';

  const SWIPE_CLOSE_PX = 40;
  const SKELETON_ROWS = 3;

  let {
    model,
    loadStatus,
    sheet,
    retryingKey,
    onchoose,
    onpaste,
    onretry,
    ontryagain,
    onclose,
  }: {
    model: SelectorModel;
    loadStatus: LoadStatus;
    sheet: boolean;
    retryingKey: string | null;
    onchoose: (row: Row) => void;
    onpaste: (input: string) => Promise<boolean>;
    onretry: (row: Row) => void;
    ontryagain: () => void;
    onclose: () => void;
  } = $props();

  let query = $state('');
  let pasteOpen = $state(false);
  let rovingIndex = $state(0);
  let scrolledToEnd = $state(false);
  let panelEl: HTMLDivElement | undefined = $state();
  let listEl: HTMLUListElement | undefined = $state();
  let swipeStartY: number | null = null;

  const signedOutEmpty = $derived(model.session === 'signed-out' && model.rows.length === 0);
  const showSkeleton = $derived(model.session === 'loading' && loadStatus === 'loading');
  const showLoadError = $derived(model.session === 'loading' && loadStatus === 'failed');
  const showNone = $derived(model.session === 'signed-in' && model.rows.length === 0);
  const showList = $derived(!signedOutEmpty && !showSkeleton && !showLoadError && model.rows.length > 0);
  const heading = $derived(signedOutEmpty ? selectorCopy.signedOutHeading : selectorCopy.listHeading);
  const visibleRows = $derived(filterRows(model.rows, query));
  const scrolls = $derived(model.rows.length > VISIBLE_ROWS);
  const showFilter = $derived(model.rows.length > FILTER_AFTER_CHARACTERS);
  const offerAddon = $derived(
    signedOutEmpty || showNone || (model.session === 'signed-in' && !model.hasAddonCharacter),
  );
  const offerSignIn = $derived(model.session === 'signed-out' || model.session === 'expired');
  const showManage = $derived(!signedOutEmpty);

  function rowButtons(): HTMLButtonElement[] {
    return listEl === undefined ? [] : [...listEl.querySelectorAll<HTMLButtonElement>('[data-row-main]')];
  }

  function focusRow(index: number): void {
    const buttons = rowButtons();
    const target = buttons[Math.max(0, Math.min(index, buttons.length - 1))];
    if (target === undefined) return;
    rovingIndex = buttons.indexOf(target);
    target.focus();
  }

  function indexOfNameStarting(letter: string, from: number): number {
    const names = rowButtons().map((button) => (button.dataset.name ?? '').toLowerCase());
    for (let step = 1; step <= names.length; step += 1) {
      const index = (from + step) % names.length;
      if (names[index].startsWith(letter.toLowerCase())) return index;
    }
    return -1;
  }

  function onListKeydown(event: KeyboardEvent): void {
    const buttons = rowButtons();
    const at = buttons.indexOf(event.target as HTMLButtonElement);
    if (at === -1) return;
    const moves: Record<string, number> = {
      ArrowDown: at + 1,
      ArrowUp: at - 1,
      Home: 0,
      End: buttons.length - 1,
    };
    if (event.key in moves) {
      event.preventDefault();
      focusRow(moves[event.key]);
      return;
    }
    if (event.key.length !== 1 || event.ctrlKey || event.metaKey || event.altKey || event.key === ' ') return;
    const next = indexOfNameStarting(event.key, at);
    if (next === -1) return;
    event.preventDefault();
    focusRow(next);
  }

  function onListScroll(): void {
    if (listEl === undefined) return;
    scrolledToEnd = listEl.scrollTop + listEl.clientHeight >= listEl.scrollHeight - 2;
  }

  function onGrabDown(event: PointerEvent): void {
    swipeStartY = event.clientY;
  }

  function onGrabUp(event: PointerEvent): void {
    if (swipeStartY !== null && event.clientY - swipeStartY > SWIPE_CLOSE_PX) onclose();
    swipeStartY = null;
  }

  function onPanelKeydown(event: KeyboardEvent): void {
    if (event.key !== 'Tab' || !sheet || panelEl === undefined) return;
    const focusable = [
      ...panelEl.querySelectorAll<HTMLElement>(
        'button:not([disabled]):not([tabindex="-1"]), a[href], input, textarea',
      ),
    ];
    const first = focusable[0];
    const last = focusable[focusable.length - 1];
    if (first === undefined || last === undefined) return;
    if (event.shiftKey && document.activeElement === first) {
      event.preventDefault();
      last.focus();
    } else if (!event.shiftKey && document.activeElement === last) {
      event.preventDefault();
      first.focus();
    }
  }

  /** Puts focus where a keyboard user starts: the current row, else the first row, else the first action. */
  export function focusInitial(): void {
    const buttons = rowButtons();
    const current = buttons.findIndex((button) => button.getAttribute('aria-current') === 'true');
    if (buttons.length > 0) {
      focusRow(current === -1 ? 0 : current);
      return;
    }
    panelEl?.querySelector<HTMLElement>('.csel-action, .csel-button')?.focus();
  }
</script>

<!-- svelte-ignore a11y_no_noninteractive_element_interactions -->
<div
  class="csel-panel"
  role="dialog"
  aria-label={selectorCopy.dialogLabel}
  aria-modal={sheet ? 'true' : undefined}
  data-testid="selector-panel"
  bind:this={panelEl}
  onkeydown={onPanelKeydown}
>
  <div
    class="csel-grab"
    role="presentation"
    onpointerdown={onGrabDown}
    onpointerup={onGrabUp}
    onpointercancel={() => (swipeStartY = null)}
  >
    <span></span>
  </div>
  <div class="csel-head">
    <span class="csel-head-title">{heading}</span>
    {#if showManage}
      <a class="csel-manage" href={selectorCopy.manageHref} data-testid="selector-manage">
        {selectorCopy.manage}
      </a>
    {/if}
    <button type="button" class="csel-close" aria-label={selectorCopy.close} onclick={onclose}>
      <svg
        width="20"
        height="20"
        viewBox="0 0 24 24"
        fill="none"
        stroke="currentColor"
        stroke-width="2"
        stroke-linecap="round"
        aria-hidden="true"
      >
        <path d="M6 6l12 12M18 6L6 18"></path>
      </svg>
    </button>
  </div>

  {#if showFilter}
    <div class="csel-filter">
      <input
        type="search"
        placeholder={selectorCopy.filterPlaceholder}
        aria-label={selectorCopy.filterLabel}
        bind:value={query}
        data-testid="selector-filter"
      />
    </div>
  {/if}

  {#if showSkeleton}
    <ul class="csel-list" aria-busy="true" data-testid="selector-loading">
      {#each Array.from({ length: SKELETON_ROWS }, (_, index) => index) as index (index)}
        <li class="csel-row" aria-hidden="true">
          <span class="csel-row-main">
            <span class="csel-skeleton" style="width:36px;height:36px;border-radius:999px"></span>
            <span class="csel-row-text">
              <span class="csel-skeleton" style="width:120px;height:14px"></span>
              <span class="csel-skeleton" style="width:210px;height:12px;margin-top:5px"></span>
              <span class="csel-skeleton" style="width:150px;height:11px;margin-top:5px"></span>
            </span>
          </span>
        </li>
      {/each}
    </ul>
  {:else if showLoadError}
    <div class="csel-message" data-testid="selector-load-error">
      <div class="csel-message-title">{selectorCopy.loadErrorTitle}</div>
      <div class="csel-message-body">{selectorCopy.loadErrorBody}</div>
      <button type="button" class="csel-button" style="margin-top:12px" onclick={ontryagain}>
        {selectorCopy.tryAgain}
      </button>
    </div>
  {:else if showNone}
    <div class="csel-message" data-testid="selector-none">
      <div class="csel-message-title">{selectorCopy.noneTitle}</div>
      <div class="csel-message-body">{selectorCopy.noneBody}</div>
    </div>
  {:else if showList}
    <div class="csel-list-wrap">
      <!-- svelte-ignore a11y_no_noninteractive_element_interactions -->
      <ul
        class="csel-list"
        class:is-scroll={scrolls}
        bind:this={listEl}
        onkeydown={onListKeydown}
        onscroll={onListScroll}
        data-testid="selector-list"
      >
        {#each visibleRows as row, index (row.key)}
          <SelectorRow
            {row}
            tabindex={index === Math.min(rovingIndex, visibleRows.length - 1) ? 0 : -1}
            retrying={retryingKey === row.key}
            {onchoose}
            {onretry}
          />
        {/each}
      </ul>
      {#if scrolls && !scrolledToEnd}<div class="csel-fade" data-testid="selector-fade"></div>{/if}
    </div>
  {/if}

  {#if model.session === 'signed-out' && model.rows.length > 0}
    <div class="csel-note">{selectorCopy.pastedNotSaved}</div>
  {:else if model.session === 'expired'}
    <div class="csel-note is-ember">{selectorCopy.sessionEnded}</div>
  {/if}

  {#if pasteOpen}
    <SelectorPaste {onpaste} oncancel={() => (pasteOpen = false)} />
  {:else}
    <button type="button" class="csel-action" data-testid="selector-paste" onclick={() => (pasteOpen = true)}>
      <span class="csel-action-icon">
        <svg
          width="18"
          height="18"
          viewBox="0 0 24 24"
          fill="none"
          stroke="currentColor"
          stroke-width="2"
          stroke-linecap="round"
          stroke-linejoin="round"
          aria-hidden="true"
        >
          <rect x="6" y="4" width="12" height="17" rx="2"></rect>
          <path d="M9 4h6v3H9zM9 12h6M9 16h4"></path>
        </svg>
      </span>
      <span>
        <span class="csel-action-title">{selectorCopy.pasteTitle}</span>
        <span class="csel-action-sub">{selectorCopy.pasteSub}</span>
      </span>
    </button>
  {/if}
  {#if offerAddon}
    <a class="csel-action" href={selectorCopy.addonHref} data-testid="selector-addon">
      <span class="csel-action-icon">
        <svg
          width="18"
          height="18"
          viewBox="0 0 24 24"
          fill="none"
          stroke="currentColor"
          stroke-width="2"
          stroke-linecap="round"
          stroke-linejoin="round"
          aria-hidden="true"
        >
          <path d="M12 4v11M7 11l5 5 5-5M5 20h14"></path>
        </svg>
      </span>
      <span>
        <span class="csel-action-title">{selectorCopy.addonTitle}</span>
        <span class="csel-action-sub">{selectorCopy.addonSub}</span>
      </span>
    </a>
  {/if}
  {#if offerSignIn}
    <a class="csel-action is-quiet" href={selectorCopy.signInHref} data-testid="selector-sign-in">
      <span class="csel-action-icon">
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
      <span>
        <span class="csel-action-title">{selectorCopy.signInTitle}</span>
        <span class="csel-action-sub">{selectorCopy.signInSub}</span>
      </span>
    </a>
  {/if}
</div>
