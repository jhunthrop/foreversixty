<!-- web/src/components/coach/CoachRowStrip.svelte -->
<!-- The row strip's face (spec 3.2 and 4.3): a full-width link under the viewer's own row, 36 px
     on desktop and 52 px on phone, the lead muted, the title strong and 600 with its 20 px spell
     icon, a `›` at the far right. The free line has two targets: the count and first title open the
     teaser, the gold underlined Premium link goes to /premium. Presentational: CoachRowLine owns
     the fetch. A strip that is loading holds its height so the list never moves. -->
<script lang="ts">
  import { COACH_SEPARATOR, coachCopy } from '../../lib/coach/copy';
  import { PREMIUM_LINKS } from '../../lib/billing/premium-links';
  import type { IconUrl } from '../../lib/coach/names';
  import { STRIP_RESERVE, type RowLine, type StripPhase, type StripReserve } from '../../lib/coach/row-line';
  import SpellIcon from './SpellIcon.svelte';

  let {
    phase,
    line = null,
    href,
    iconUrl,
    label,
    reserve,
    onOpen,
    onRetry,
  }: {
    phase: StripPhase;
    line?: RowLine | null;
    href: string;
    iconUrl: IconUrl;
    /** The strip's accessible name: the whole sentence, in order. */
    label: string;
    /** The viewer's entitlement, which sizes the phone slot before the answer lands. */
    reserve: StripReserve;
    onOpen: (event: MouseEvent) => void;
    onRetry: () => void;
  } = $props();

  const BASE = $derived(
    `flex ${STRIP_RESERVE[reserve]} items-center bg-[#0b0f17] shadow-[inset_0_1px_0_var(--color-line-soft)] pr-2 pl-12 text-[13px] leading-[18px] md:pl-[100px]`,
  );
  const FOCUS = 'focus-visible:outline-gold focus-visible:-outline-offset-2 focus-visible:outline-2';
  /** The Premium link's ring stands 2 px off its 4 px padding, clear of the underline. */
  const LINK_FOCUS = 'focus-visible:outline-gold focus-visible:outline-offset-2 focus-visible:outline-2';
</script>

{#if phase === 'loading'}
  <div class={BASE} role="status" aria-busy="true" data-testid="coach-row-loading">
    <span class="sr-only">{coachCopy.loading}</span>
    <span class="skeleton-block h-3 w-[260px] max-w-full" aria-hidden="true"></span>
  </div>
{:else if phase === 'error'}
  <button
    type="button"
    class="{BASE} {FOCUS} text-muted hover:bg-card-top w-full text-left"
    onclick={onRetry}
    data-testid="coach-row-error"
  >
    {coachCopy.rowError}&nbsp;<span class="underline">{coachCopy.rowRetry}</span>
  </button>
{:else if line !== null}
  {#if line.upgrade === ''}
    <a
      {href}
      class="{BASE} {FOCUS} hover:bg-card-top group"
      aria-label={label}
      onclick={onOpen}
      data-testid="coach-row-line"
    >
      <span class="text-muted">{line.lead}</span>
      {#if line.title !== ''}
        <SpellIcon url={iconUrl(line.spellId)} size={20} />
        <span class="text-strong ml-1 font-semibold group-hover:underline group-hover:underline-offset-[3px]"
          >{line.title}</span
        >
      {/if}
      <span class="text-muted ml-auto pl-2" aria-hidden="true">›</span>
    </a>
  {:else}
    <div
      class="{BASE} flex-col items-stretch gap-0 py-0 md:flex-row md:items-stretch"
      data-testid="coach-row-line"
    >
      <!-- The first target fills the strip's 36 px with 8 px each side, so its ring sits on the
           strip's edge as the Premium strip's does. -->
      <a
        {href}
        class="{FOCUS} hover:bg-card-top group flex min-h-11 items-center md:-ml-2 md:min-h-9 md:px-2"
        onclick={onOpen}
        data-testid="coach-row-open"
      >
        <span class="text-muted">{line.lead}</span>
        <SpellIcon url={iconUrl(line.spellId)} size={20} />
        <span class="text-strong ml-1 font-semibold group-hover:underline group-hover:underline-offset-[3px]"
          >{line.title}</span
        >{#if line.after !== ''}<span class="text-muted whitespace-pre">{line.after}</span>{/if}
      </a>
      <span class="text-muted hidden items-center whitespace-pre md:flex" aria-hidden="true"
        >{COACH_SEPARATOR}</span
      >
      <a
        href={PREMIUM_LINKS.coach}
        class="text-gold hover:text-gold-hover {LINK_FOCUS} flex min-h-11 items-center self-start underline md:min-h-0 md:self-center md:px-1"
        data-testid="coach-row-premium">{line.upgrade}</a
      >
    </div>
  {/if}
{/if}
