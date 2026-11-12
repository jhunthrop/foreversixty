<!-- web/src/components/creators/CreatorStats.svelte -->
<!-- The creator's own /creators/stats (design/specs/2026-11-08-creator-program.md 3d, landing 0G): the
     creator's name as the page title, a split Page of panels. Main: Your figures (the open month's
     estimate in gold, then the all-time counts with their rates), then Your players, one Row per week
     with its three counts in the trail, never a name, "still filling" while the week is under 14 days
     old. Rail: Payouts, the Stripe status Tag in its head, the last closed month and the day it is
     approved by, every approved month, Set up payouts while Stripe is not ready and Payouts and tax
     forms (a fresh Express dashboard link per press) once it is. -->
<script lang="ts">
  import { AccountError } from '../../lib/account/api';
  import { dashboardLink, fetchCreatorMe } from '../../lib/creators/api';
  import { creatorStatsCopy as copy, onboardingCopy, stripeStatusCopy } from '../../lib/creators/copy';
  import {
    STATUS_TONE,
    approvedByDay,
    money,
    monthLabel,
    previousMonth,
    statsFigures,
    weekRow,
    type CreatorMe,
  } from '../../lib/creators/model';
  import SignInPrompt from '../SignInPrompt.svelte';
  import Button from '../ui/Button.svelte';
  import EmptyState from '../ui/EmptyState.svelte';
  import Figure from '../ui/Figure.svelte';
  import LoadError from '../ui/LoadError.svelte';
  import Page from '../ui/Page.svelte';
  import PageHeader from '../ui/PageHeader.svelte';
  import Row from '../ui/Row.svelte';
  import Section from '../ui/Section.svelte';
  import Skeleton from '../ui/Skeleton.svelte';
  import Tag from '../ui/Tag.svelte';

  type Phase = 'loading' | 'signed-out' | 'not-creator' | 'failed' | 'ready';

  let {
    go = (url: string) => window.location.assign(url),
    now = () => new Date(),
  }: { go?: (url: string) => void; now?: () => Date } = $props();

  let phase = $state<Phase>('loading');
  let me = $state<CreatorMe | null>(null);
  let opening = $state(false);
  let dashError = $state('');

  async function load(): Promise<void> {
    phase = 'loading';
    try {
      me = await fetchCreatorMe();
      phase = 'ready';
    } catch (error) {
      if (error instanceof AccountError && error.status === 401) phase = 'signed-out';
      else if (error instanceof AccountError && error.status === 403) phase = 'not-creator';
      else phase = 'failed';
    }
  }

  $effect(() => {
    void load();
  });

  async function openDashboard(): Promise<void> {
    opening = true;
    dashError = '';
    try {
      go(await dashboardLink());
    } catch {
      dashError = copy.dashboardFailed;
      opening = false;
    }
  }

  const paidSoFar = (mine: CreatorMe): number => mine.payouts.reduce((sum, p) => sum + p.amount_cents, 0);
</script>

{#if phase === 'ready' && me !== null}
  {@const mine = me}
  {@const closed = previousMonth(mine.month)}
  <PageHeader title={mine.name}>
    {#snippet meta()}<span data-testid="stats-meta">{copy.meta(mine.slug, mine.code)}</span>{/snippet}
  </PageHeader>
  <Page layout="split" panels testid="stats-page-body">
    <Section title={copy.summaryLabel} meta={copy.summaryMeta} panel testid="stats-summary">
      <div class="flex flex-wrap items-start gap-x-6 gap-y-3" data-testid="stats-figures">
        {#each statsFigures(mine) as fig (fig.key)}
          <span class={fig.lead ? 'w-[240px]' : 'w-[150px]'}
            ><Figure
              value={fig.value}
              label={fig.label}
              caption={fig.note}
              tone={fig.lead ? 'gold' : 'strong'}
            /></span
          >
        {/each}
      </div>
    </Section>
    <Section title={copy.playersTitle} meta={copy.playersMeta} panel flush testid="stats-players">
      {#if mine.weeks.length === 0}
        <EmptyState message={copy.playersEmpty} inset testid="stats-players-empty" />
      {:else}
        <div role="group" aria-label={copy.playersTitle}>
          {#each mine.weeks as week (week.week_of)}
            {@const row = weekRow(week, now())}
            <Row
              title={row.title}
              sub={row.filling ? copy.filling : undefined}
              density="slim"
              wrapTrail
              testid={`stats-week-${week.week_of}`}
            >
              {#snippet trail()}
                <span class="flex shrink-0 items-start gap-x-5">
                  {#each row.figures as fig (fig.key)}
                    <span class="w-[84px]"><Figure value={fig.value} label={fig.label} size="sm" /></span>
                  {/each}
                </span>
              {/snippet}
            </Row>
          {/each}
        </div>
      {/if}
    </Section>
    {#snippet rail()}
      <Section title={copy.payoutsTitle} panel testid="stats-payouts">
        {#snippet head()}
          <Tag tone={STATUS_TONE[mine.stripe_status]} testid="stats-stripe-status"
            >{stripeStatusCopy[mine.stripe_status]}</Tag
          >
        {/snippet}
        <div class="flex flex-col gap-4">
          <p class="text-[14px] leading-5" data-testid="stats-closed-month">
            {copy.closedMonth(monthLabel(closed), approvedByDay(mine.month))}
          </p>
          {#if mine.stripe_status !== 'ready'}
            <div>
              <Button tone="primary" href="/creators/onboarding" testid="stats-setup">{copy.setup}</Button>
            </div>
          {/if}
          {#if mine.payouts.length === 0}
            <EmptyState message={copy.payoutsEmpty} testid="stats-payouts-empty" />
          {:else}
            <div role="group" aria-label={copy.payoutsTitle}>
              {#each mine.payouts as p (p.month)}
                <Row title={monthLabel(p.month)} density="slim">
                  {#snippet trail()}<span class="tabular-nums">{money(p.amount_cents)}</span>{/snippet}
                </Row>
              {/each}
            </div>
            <p class="text-muted text-[14px]" data-testid="stats-paid-so-far">
              {copy.paidSoFar(money(paidSoFar(mine)))}
            </p>
          {/if}
          {#if mine.stripe_status === 'ready'}
            <div>
              <Button
                busy={opening}
                busyLabel={copy.dashboardOpening}
                onclick={() => void openDashboard()}
                testid="stats-dashboard">{copy.dashboard}</Button
              >
            </div>
            {#if dashError !== ''}<p class="text-wipe text-[14px]" role="alert">{dashError}</p>{/if}
          {/if}
        </div>
      </Section>
    {/snippet}
  </Page>
{:else}
  <PageHeader title={onboardingCopy.pageTitle} />
  <Page layout="form">
    {#if phase === 'loading'}
      <div data-testid="stats-loading"><Skeleton lines={5} rowHeight="h-6" /></div>
    {:else if phase === 'signed-out'}
      <SignInPrompt line={copy.signedOut} next="/creators/stats" testid="stats-signin" />
    {:else if phase === 'not-creator'}
      <Section title={onboardingCopy.stripeTitle} panel>
        <p data-testid="stats-not-creator">{onboardingCopy.notCreator}</p>
      </Section>
    {:else}
      <LoadError message={copy.loadFailed} onRetry={() => void load()} testid="stats-load-error" />
    {/if}
  </Page>
{/if}
