<!-- web/src/components/guild/GuildWeekOneCard.svelte -->
<!-- Before raids open on the Overview (design/specs/2026-11-05-guild-week-one.md §3.1, redrawn in
     the home B direction 2026-11-07): the first raids' countdown, then an open Section with the
     guild's level spread as a Tally, who has which profession and the viewer's one next step (a
     sentence, one Button, and an officer's invite link), then the guild's activity. No box of its
     own: the Overview's one Plate is the next raid. The words and branches are lib/guild/week-one.ts;
     `loading` and `error` hold the section's height so the page does not shift. -->
<script lang="ts">
  import CharacterName from '../ui/CharacterName.svelte';
  import EmptyState from '../ui/EmptyState.svelte';
  import activeBuild from '../../data/active-build.json';
  import { minuteClock } from '../../lib/ui/minute-clock.svelte';
  import {
    levelSpread,
    levelTally,
    professionFacts,
    raidsCountdownTarget,
    spreadIsEmpty,
    weekOneCopy,
    weekOneMeta,
    weekOneStep,
    weekOneTitle,
    type ProfessionHolder,
    type WeekOneAction,
    type WeekOneCharacter,
  } from '../../lib/guild/week-one';
  import type { GuildRosterEntry, NextRaidView } from '../../lib/guild/api';

  import { REGION } from '../../lib/ui/regions';
  import { viewerLocale } from '../../lib/dates';
  import RaidsCountdown from '../RaidsCountdown.svelte';
  import ActionLink from '../ui/ActionLink.svelte';
  import Button from '../ui/Button.svelte';
  import LoadError from '../ui/LoadError.svelte';
  import Section from '../ui/Section.svelte';
  import Skeleton from '../ui/Skeleton.svelte';
  import Tally from '../ui/Tally.svelte';
  import FactLine from './FactLine.svelte';
  import GuildActivity from './overview/GuildActivity.svelte';

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
</script>

{#snippet primary(action: WeekOneAction)}
  {#if action.kind === 'link'}
    <Button tone="primary" size="sm" href={action.href} testid="guild-week-one-action">{action.label}</Button>
  {:else}
    <Button
      tone="primary"
      size="sm"
      onclick={action.kind === 'approve' ? onApprove : onPlan}
      testid="guild-week-one-action">{action.label}</Button
    >
  {/if}
{/snippet}

<div data-region={REGION.countdowns}>
  <RaidsCountdown
    raidsOpen={raidsCountdownTarget(clock.now)}
    build={activeBuild.build}
    testid="guild-week-one-countdown"
  />
</div>

<Section
  panel
  title={weekOneTitle(clock.now)}
  meta={weekOneMeta(clock.now)}
  testid="guild-overview-card-week-one"
>
  {#if status === 'error'}
    <div class="min-h-[180px]">
      <LoadError message={weekOneCopy.error} {onRetry} testid="guild-week-one-error" />
    </div>
  {:else if status === 'loading'}
    <Skeleton lines={4} minHeight="min-h-[180px]" testid="guild-week-one-loading" />
  {:else}
    <div class="flex flex-col gap-4">
      {#if spreadIsEmpty(spread)}
        <EmptyState message={weekOneCopy.empty} testid="guild-week-one-empty" />
      {:else}
        <div class="flex flex-col gap-2">
          <Tally items={levelTally(spread)} testid="guild-week-one-spread" />
          {#if spread.noLevel > 0}
            <p class="text-muted text-small" data-testid="guild-week-one-no-level">
              {weekOneCopy.noLevel(spread.noLevel)}
            </p>
          {/if}
        </div>
      {/if}

      {#if professions.length > 0}
        <p class="text-muted text-small" data-testid="guild-week-one-professions">
          <span class="text-strong font-semibold">{weekOneCopy.professions}</span>
          <FactLine facts={professions} />
        </p>
      {/if}

      <div
        class="flex flex-wrap items-center justify-between gap-x-6 gap-y-3"
        data-testid="guild-week-one-step"
      >
        <p
          class={step.action === null ? 'text-muted text-[14px]' : 'text-[15px] leading-6'}
          data-testid="guild-week-one-sentence"
        >
          {#if step.name !== null && character !== null}<CharacterName
              name={step.name}
              cls={character.class}
            />{/if}{step.text}
        </p>
        <div class="flex flex-wrap items-center gap-4">
          {#if step.invite && inviteHref !== ''}
            <ActionLink href={inviteHref} data-testid="guild-week-one-invite">{weekOneCopy.invite}</ActionLink
            >
          {/if}
          {#if step.action !== null}{@render primary(step.action)}{/if}
        </div>
      </div>
    </div>
  {/if}
</Section>

{#if status === 'ready'}<GuildActivity now={fixedNow} />{/if}
