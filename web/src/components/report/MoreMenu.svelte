<!-- web/src/components/report/MoreMenu.svelte -->
<!-- The tab row's More menu (design/specs/2026-10-30-controls-charts.md §A10 row 3, §C): the
     tabs past the first six, on the shared popover core, so the menu renders in the top layer
     and no chart or sticky row can paint over it. The dropdown list's surface and 36 px rows,
     `role="menu"`, the current tab checked. Arrows move between entries, Escape closes it and
     returns to the button, one menu open at a time. -->
<script lang="ts">
  import type { Tab } from '../../lib/report/url';
  import { popoverSupported } from '../../lib/ui/listbox-model';
  import type { PopoverHandle } from '../../lib/ui/popover';
  import Glyph from '../ui/Glyph.svelte';

  let {
    tabs,
    current,
    label,
    menuLabel,
    active,
    buttonClass,
    onselect,
  }: {
    tabs: readonly { id: Tab; label: string }[];
    current: Tab;
    /** The button's text: the current tab's label when it lives in this menu, else "More". */
    label: string;
    menuLabel: string;
    /** The current tab is one of this menu's. */
    active: boolean;
    buttonClass: string;
    onselect: (tab: Tab) => void;
  } = $props();

  let wrap = $state<HTMLElement | null>(null);
  let button = $state<HTMLButtonElement | null>(null);
  let panel = $state<HTMLElement | null>(null);
  let open = $state(false);
  let handle: PopoverHandle | null = null;
  const fallback = !popoverSupported();

  const items = (): HTMLElement[] => [
    ...(panel?.querySelectorAll<HTMLElement>('[role="menuitemradio"]') ?? []),
  ];

  function focusItem(index: number): void {
    const list = items();
    if (list.length === 0) return;
    list[(index + list.length) % list.length].focus();
  }

  async function show(focus: 'current' | 'first' | 'last' | 'none'): Promise<void> {
    if (open || wrap === null || button === null || panel === null) return;
    open = true;
    if (!fallback) {
      const { openPopover } = await import('../../lib/ui/popover');
      if (!open) return;
      handle = openPopover({
        anchor: button,
        panel,
        container: wrap,
        align: 'start',
        matchWidth: false,
        onDismiss: () => {
          handle = null;
          open = false;
        },
      });
    }
    const list = items();
    const currentIndex = list.findIndex((item) => item.getAttribute('aria-checked') === 'true');
    if (focus === 'current') focusItem(currentIndex === -1 ? 0 : currentIndex);
    if (focus === 'first') focusItem(0);
    if (focus === 'last') focusItem(-1);
  }

  function hide(refocus: boolean): void {
    if (!open) return;
    handle?.close();
    handle = null;
    open = false;
    if (refocus) button?.focus();
  }

  function onButtonKey(event: KeyboardEvent): void {
    if (event.key === 'ArrowDown' || event.key === 'Enter' || event.key === ' ') {
      event.preventDefault();
      void show(event.key === 'ArrowDown' ? 'first' : 'current');
    } else if (event.key === 'ArrowUp') {
      event.preventDefault();
      void show('last');
    }
  }

  function onMenuKey(event: KeyboardEvent): void {
    const list = items();
    const index = list.indexOf(document.activeElement as HTMLElement);
    const moves: Record<string, () => void> = {
      ArrowDown: () => focusItem(index + 1),
      ArrowUp: () => focusItem(index - 1),
      Home: () => focusItem(0),
      End: () => focusItem(-1),
    };
    if (event.key === 'Escape') {
      event.preventDefault();
      event.stopPropagation();
      hide(true);
    } else if (event.key === 'Tab') {
      hide(true);
    } else if (moves[event.key] !== undefined) {
      event.preventDefault();
      moves[event.key]();
    }
  }

  function pick(tab: Tab): void {
    hide(true);
    onselect(tab);
  }
</script>

<div class="relative" bind:this={wrap}>
  <button
    bind:this={button}
    type="button"
    class={buttonClass}
    class:border-gold={active || open}
    class:text-strong={active || open}
    class:border-transparent={!active && !open}
    class:text-nav={!active && !open}
    class:bg-row-hover={open}
    aria-haspopup="menu"
    aria-expanded={open}
    aria-controls="report-more-menu"
    data-testid="tab-more"
    onclick={() => (open ? hide(true) : void show('none'))}
    onkeydown={onButtonKey}
  >
    {label}
  </button>
  <div
    bind:this={panel}
    id="report-more-menu"
    class="sel-pop sel-menu"
    popover={fallback ? undefined : 'manual'}
    data-fallback={fallback && open ? 'true' : undefined}
    role="menu"
    tabindex="-1"
    aria-label={menuLabel}
    data-testid="tab-more-menu"
    onkeydown={onMenuKey}
  >
    <div class="sel-scroll">
      {#each tabs as tab (tab.id)}
        <button
          type="button"
          role="menuitemradio"
          aria-checked={current === tab.id}
          class="sel-row sel-menu-item"
          data-testid={`tab-${tab.id}`}
          onclick={() => pick(tab.id)}
        >
          {tab.label}
          {#if current === tab.id}<Glyph name="check" class="sel-check" />{/if}
        </button>
      {/each}
    </div>
  </div>
</div>
