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
  import { battlenetStartUrl } from '../../lib/account/api';
  import type { AccountPanel } from '../../lib/character-selector/account-panel';
  import SelectorPaste from './SelectorPaste.svelte';
  import SelectorRow from './SelectorRow.svelte';

  const SWIPE_CLOSE_PX = 40;
  const SKELETON_ROWS = 3;
  /** Stroke paths on a 24px grid for the account rows' tiles. */
  const GLYPHS = {
    person: 'M12 12a4 4 0 100-8 4 4 0 000 8zM4 21c0-4 3.6-7 8-7s8 3 8 7',
    sliders: 'M4 7h9M17 7h3M4 17h3M11 17h9M15 4v6M9 14v6',
    crown: 'M4 18l-1-10 5 4 4-7 4 7 5-4-1 10zM4 21h16',
    signOut: 'M10 4H5v16h5M14 8l-4 4 4 4M10 12h10',
  } as const;

  let {
    model,
    loadStatus,
    sheet,
    retryingKey,
    account,
    signingOut,
    signOutFailed,
    restored,
    onsignout,
    onforget,
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
    /** The signed-in account's head and rows; null signed out or until `/v1/me` answered. */
    account: AccountPanel | null;
    signingOut: boolean;
    signOutFailed: boolean;
    /** The current character came back from an earlier visit: a muted line and Forget. */
    restored: boolean;
    onsignout: () => void;
    onforget: () => void;
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
    !showNone && (signedOutEmpty || (model.session === 'signed-in' && !model.hasAddonCharacter)),
  );
  const offerSignIn = $derived(model.session === 'signed-out' || model.session === 'expired');
  const showManage = $derived(!signedOutEmpty);
  const signInHref = battlenetStartUrl(window.location.pathname + window.location.search);

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

{#snippet glyph(path: string)}
  <svg
    width="18"
    height="18"
    viewBox="0 0 24 24"
    fill="none"
    stroke="currentColor"
    stroke-width="1.9"
    stroke-linecap="round"
    stroke-linejoin="round"
    aria-hidden="true"
  >
    <path d={path}></path>
  </svg>
{/snippet}

{#snippet chevron(gold = false)}
  <svg
    class={['csel-arow-chev', gold && 'is-gold']}
    width="16"
    height="16"
    viewBox="0 0 24 24"
    fill="none"
    stroke="currentColor"
    stroke-width="2.2"
    stroke-linecap="round"
    stroke-linejoin="round"
    aria-hidden="true"
  >
    <path d="M9 6l6 6-6 6"></path>
  </svg>
{/snippet}

{#snippet signInButtons()}
  <div class="csel-signin" data-testid="selector-sign-in">
    <a class="csel-cta" href={signInHref} data-testid="selector-sign-in-bnet">{selectorCopy.signInBnet}</a>
    <a class="csel-cta is-secondary" href={selectorCopy.signInHref} data-testid="selector-sign-in-email"
      >{selectorCopy.signInEmail}</a
    >
  </div>
{/snippet}

{#snippet accountRows(panel: AccountPanel)}
  {@const words = selectorCopy.account}
  <div class="csel-arows" data-testid="selector-account-rows">
    {#if panel.profile !== null}
      <a class="csel-arow" href={panel.profile.href} data-testid="selector-row-profile">
        <span class="csel-tile is-small">{@render glyph(GLYPHS.person)}</span>
        <span class="csel-arow-text">
          <span class="csel-arow-title">{words.profile}</span>
          <span class="csel-arow-sub">{words.profileSub}</span>
        </span>
        {@render chevron()}
      </a>
    {/if}
    <a class="csel-arow" href={words.accountHref} data-testid="selector-row-account">
      <span class="csel-tile is-small">{@render glyph(GLYPHS.sliders)}</span>
      <span class="csel-arow-text">
        <span class="csel-arow-title">{words.account}</span>
        <span class="csel-arow-sub">{words.accountSub}</span>
      </span>
      {@render chevron()}
    </a>
    <a class="csel-arow" href={panel.plan.href} data-testid="selector-row-plan">
      <span class="csel-tile is-small">{@render glyph(GLYPHS.crown)}</span>
      <span class="csel-arow-text">
        <span class="csel-arow-line">
          <span class="csel-arow-title">{words.plan}</span>
          {#if panel.plan.founder}<span class="csel-tag is-outline">{words.founderTag}</span>{/if}
        </span>
        <span class="csel-arow-sub">{panel.plan.sub}</span>
        {#if panel.plan.sub2 !== null}<span class="csel-arow-sub">{panel.plan.sub2}</span>{/if}
      </span>
      {#if panel.plan.seePlans}
        <!-- Text inside the one link, not a second control: the name reads Plan, Free, See plans. -->
        <span class="csel-see-plans" data-testid="selector-plan-see-plans">{words.seePlans}</span>
      {/if}
      {@render chevron(panel.plan.seePlans)}
    </a>
    <button
      type="button"
      class="csel-arow is-quiet"
      disabled={signingOut}
      aria-busy={signingOut}
      data-testid="selector-row-sign-out"
      onclick={onsignout}
    >
      <span class="csel-tile is-small is-quiet">
        {#if signingOut}<i class="csel-spinner" aria-hidden="true"></i>{:else}{@render glyph(
            GLYPHS.signOut,
          )}{/if}
      </span>
      <span class="csel-arow-text">
        <span class="csel-arow-title">{signingOut ? words.signingOut : words.signOut}</span>
      </span>
    </button>
    {#if signOutFailed}
      <p class="csel-arow-error" role="alert" data-testid="selector-sign-out-failed">{words.signOutFailed}</p>
    {/if}
  </div>
{/snippet}

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
  {#if account !== null}
    <div class="csel-acct-head" data-testid="selector-account-head">
      <span class="csel-tile">{@render glyph(GLYPHS.person)}</span>
      <span class="csel-acct-id">
        <span class="csel-acct-name" data-testid="selector-account-name">{account.head.name}</span>
        <span class="csel-acct-via">{selectorCopy.account.signedInWith[account.head.via]}</span>
      </span>
      {#if account.head.premium}
        <span class="csel-tag" data-testid="selector-premium-tag"
          ><i aria-hidden="true"></i>{selectorCopy.account.premiumTag}</span
        >
      {/if}
    </div>
  {/if}
  <div class="csel-head" class:is-signout={signedOutEmpty}>
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

  {#if signedOutEmpty}
    <p class="csel-signout-body" data-testid="selector-signed-out-body">{selectorCopy.signedOutBody}</p>
    {@render signInButtons()}
  {/if}

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
      <div class="csel-message-body" data-testid="selector-none-body">{selectorCopy.accountPanelBody}</div>
      <div class="csel-message-links">
        <a href={selectorCopy.accountSetUpHref} data-testid="selector-get-set-up"
          >{selectorCopy.accountSetUp}</a
        >
        <a href={selectorCopy.accountLinkHref} data-testid="selector-account">{selectorCopy.accountLink}</a>
      </div>
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

  {#if restored}
    <div class="csel-note csel-restored" data-testid="selector-restored">
      <span>{selectorCopy.restoredLine}</span>
      <button
        type="button"
        class="csel-manage"
        aria-label={model.current === null
          ? selectorCopy.forget
          : selectorCopy.forgetAria(model.current.name)}
        onclick={onforget}
        data-testid="selector-forget"
      >
        {selectorCopy.forget}
      </button>
    </div>
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
  {#if offerSignIn && !signedOutEmpty}
    {@render signInButtons()}
  {/if}
  {#if account !== null}
    {@render accountRows(account)}
  {/if}
</div>

<style>
  /* The free Plan row's one action (spec 2026-10-29-premium-entry section 6): the only gold
     below the character list, turning to the hover gold with the row. */
  .csel-see-plans {
    flex-shrink: 0;
    font-size: 13px;
    font-weight: 600;
    color: var(--color-gold);
  }
  .csel-arow-chev.is-gold {
    color: var(--color-gold);
  }
  :global(.csel-arow:hover) > .csel-see-plans,
  :global(.csel-arow:hover) > .csel-arow-chev.is-gold {
    color: var(--color-gold-hover);
  }
</style>
