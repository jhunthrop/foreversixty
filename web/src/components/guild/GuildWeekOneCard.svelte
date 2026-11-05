<!-- web/src/components/guild/GuildWeekOneCard.svelte -->
<!-- The Before raids open card (design/specs/2026-11-05-guild-week-one.md §3.1): the first card on
     Overview while nothing is logged in the tier. The guild's level spread, when the raids open,
     who has which profession, and the viewer's one next step (a quiet line, no button, when the
     band above already holds the sign-up segments for the next event). Below 768 the next step leads the
     card and professions close it, so the primary button sits above the 640 px fold. The words and branches are
     lib/guild/week-one.ts; `loading` and `error` hold its height so the page does not shift. -->
<script lang="ts">
  import { minuteClock } from '../../lib/guild/minute-clock.svelte';
  import {
    levelSpread,
    professionFacts,
    raidsOpenLine,
    spreadIsEmpty,
    weekOneCopy,
    weekOneStep,
    weekOneTitle,
    type ProfessionHolder,
    type WeekOneAction,
    type WeekOneCharacter,
  } from '../../lib/guild/week-one';
  import type { GuildRosterEntry, NextRaidView } from '../../lib/guild/api';
  import { viewerLocale } from '../../lib/guild/raid-time';
  import { classColorVar } from '../../lib/report/format';
  import { button } from '../../lib/ui/button';
  import ActionLink from '../ui/ActionLink.svelte';
  import Card from '../ui/Card.svelte';
  import LoadError from '../ui/LoadError.svelte';
  import StatTile from '../ui/StatTile.svelte';
  import FactLine from './FactLine.svelte';

  let {
    status,
    officer,
    waiting,
    entries,
    rows,
    character,
    nextRaid,
    raidHref,
    inviteHref,
    onRetry,
    onApprove,
    onPlan,
    bandSignsUp = false,
    planGated = false,
    fixedNow,
  }: {
    status: 'loading' | 'error' | 'ready';
    officer: boolean;
    waiting: number;
    entries: readonly GuildRosterEntry[];
    rows: readonly ProfessionHolder[];
    character: WeekOneCharacter | null;
    nextRaid: NextRaidView | null;
    raidHref: string | null;
    /** The Settings page's invite link; empty hides the secondary link. */
    inviteHref: string;
    onRetry: () => void;
    onApprove: () => void;
    onPlan: () => void;
    /** The guild band shows the sign-up segments for `nextRaid`: the raider's step is a quiet line. */
    bandSignsUp?: boolean;
    /** The guild's plan gates planning: the officer's empty-calendar sentence names the plan. */
    planGated?: boolean;
    fixedNow?: Date;
  } = $props();

  const clock = minuteClock(() => fixedNow);
  const locale = viewerLocale();

  const spread = $derived(levelSpread(entries));
  const professions = $derived(professionFacts(rows));
  const step = $derived(
    weekOneStep({
      officer,
      waiting,
      nextRaid,
      raidHref,
      character,
      now: clock.now,
      locale,
      bandSignsUp,
      planGated,
    }),
  );
  const tiles = $derived([
    { value: spread.at60, label: 'AT 60', gold: spread.at60 > 0 },
    { value: spread.fifties, label: '50 TO 59', gold: false },
    { value: spread.forties, label: '40 TO 49', gold: false },
    { value: spread.under40, label: 'UNDER 40', gold: false },
  ]);
  /** Below 768 the primary leads its row, above the secondary link (§3.1 Row 4, guild round 2). */
  const actionClass = `${button('primary', 'md')} max-md:order-first max-md:w-full`;
</script>

{#snippet primary(action: WeekOneAction)}
  {#if action.kind === 'link'}
    <a href={action.href} class={actionClass} data-testid="guild-week-one-action">{action.label}</a>
  {:else}
    <button
      type="button"
      class={actionClass}
      onclick={action.kind === 'approve' ? onApprove : onPlan}
      data-testid="guild-week-one-action">{action.label}</button
    >
  {/if}
{/snippet}

<Card glyph="target" title={weekOneTitle(clock.now)} testid="guild-overview-card-week-one">
  {#if status === 'error'}
    <div class="min-h-[460px] md:min-h-[300px]">
      <LoadError message={weekOneCopy.error} {onRetry} testid="guild-week-one-error" />
    </div>
  {:else if status === 'loading'}
    <div
      class="flex min-h-[460px] flex-col gap-4 md:min-h-[300px]"
      role="status"
      aria-busy="true"
      data-testid="guild-week-one-loading"
    >
      <div class="grid grid-cols-2 gap-3 md:grid-cols-[repeat(4,132px)]">
        {#each [0, 1, 2, 3] as tile (tile)}<span class="skeleton-block h-[76px]" aria-hidden="true"
          ></span>{/each}
      </div>
      <span class="skeleton-block h-[15px] w-[70%]" aria-hidden="true"></span>
      <span class="skeleton-block h-[13px] w-[70%]" aria-hidden="true"></span>
      <span class="skeleton-block h-11" aria-hidden="true"></span>
    </div>
  {:else}
    <div class="flex flex-col gap-4">
      {#if spreadIsEmpty(spread)}
        <p class="text-muted text-[14px]" data-testid="guild-week-one-empty">{weekOneCopy.empty}</p>
      {:else}
        <div class="flex flex-col gap-2">
          <div
            class="grid grid-cols-2 gap-3 md:grid-cols-[repeat(4,132px)]"
            data-testid="guild-week-one-tiles"
          >
            {#each tiles as tile (tile.label)}
              <StatTile value={String(tile.value)} label={tile.label} tone={tile.gold ? 'gold' : 'strong'} />
            {/each}
          </div>
          {#if spread.noLevel > 0}
            <p class="text-muted text-[13px]" data-testid="guild-week-one-no-level">
              {weekOneCopy.noLevel(spread.noLevel)}
            </p>
          {/if}
        </div>
      {/if}

      <p class="text-text text-[15px]" data-testid="guild-week-one-raids-line">
        {raidsOpenLine(clock.now, locale)}
      </p>

      {#if professions.length > 0}
        <p class="text-muted text-[13px] max-md:order-last" data-testid="guild-week-one-professions">
          <span class="quiet-label">{weekOneCopy.professions}</span>
          <FactLine facts={professions} />
        </p>
      {/if}

      <div
        class="border-line-soft flex flex-wrap items-center justify-between gap-4 max-md:order-first max-md:flex-col max-md:items-stretch max-md:border-b max-md:pb-4 md:border-t md:pt-4"
        data-testid="guild-week-one-step"
      >
        <p
          class={step.action === null ? 'text-muted text-[14px]' : 'text-[15px]'}
          data-testid="guild-week-one-sentence"
        >
          {#if step.name !== null && character !== null}<span
              class="font-semibold"
              style={`color:${classColorVar(character.class)}`}>{step.name}</span
            >{/if}{step.text}
        </p>
        <div class="flex flex-wrap items-center gap-4 max-md:flex-col max-md:items-stretch">
          {#if step.invite && inviteHref !== ''}
            <ActionLink href={inviteHref} data-testid="guild-week-one-invite">{weekOneCopy.invite}</ActionLink
            >
          {/if}
          {#if step.action !== null}{@render primary(step.action)}{/if}
        </div>
      </div>
    </div>
  {/if}
</Card>
