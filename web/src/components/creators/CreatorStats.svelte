<!-- web/src/components/creators/CreatorStats.svelte -->
<!-- The creator's own /creators/stats (design/specs/2026-11-08-creator-program.md 3d): the creator's
     name as the page title, a split Page. Main: the one Plate of four Figures (Earned so far in gold),
     then Your players, one row per week and never a name. Rail: Payouts, with Set up payouts while
     Stripe is not ready and Payouts and tax forms (a fresh Express dashboard link per press) once it is. -->
<script lang="ts">
  import { AccountError } from '../../lib/account/api';
  import { dashboardLink, fetchCreatorMe } from '../../lib/creators/api';
  import { creatorStatsCopy as copy, onboardingCopy } from '../../lib/creators/copy';
  import { money, monthLabel, statsFigures, weekText, type CreatorMe } from '../../lib/creators/model';
  import { button } from '../../lib/ui/button';
  import SignInPrompt from '../SignInPrompt.svelte';
  import Figure from '../ui/Figure.svelte';
  import LoadError from '../ui/LoadError.svelte';
  import Page from '../ui/Page.svelte';
  import PageHeader from '../ui/PageHeader.svelte';
  import Plate from '../ui/Plate.svelte';
  import Row from '../ui/Row.svelte';
  import Section from '../ui/Section.svelte';
  import Skeleton from '../ui/Skeleton.svelte';

  type Phase = 'loading' | 'signed-out' | 'not-creator' | 'failed' | 'ready';

  let { go = (url: string) => window.location.assign(url) }: { go?: (url: string) => void } = $props();

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
</script>

{#if phase === 'ready' && me !== null}
  {@const mine = me}
  <PageHeader title={mine.name}>
    {#snippet meta()}<span data-testid="stats-meta">{copy.meta(mine.slug, mine.code)}</span>{/snippet}
  </PageHeader>
  <Page layout="split" testid="stats-page-body">
    <div class="flex flex-col gap-8">
      <Plate label={copy.summaryLabel}>
        <div class="flex flex-wrap items-start gap-x-5 gap-y-2 p-4" data-testid="stats-figures">
          {#each statsFigures(mine) as fig (fig.key)}
            <span class={fig.lead ? 'w-[220px]' : 'w-[104px]'}
              ><Figure
                value={fig.value}
                label={fig.label}
                caption={fig.note}
                size="sm"
                tone={fig.lead ? 'gold' : 'strong'}
              /></span
            >
          {/each}
        </div>
      </Plate>
      <Section title={copy.playersTitle} meta={copy.playersMeta} testid="stats-players">
        {#if mine.weeks.length === 0}
          <p data-testid="stats-players-empty">{copy.playersEmpty}</p>
        {:else}
          <div>
            {#each mine.weeks as week (week.week_of)}
              <Row title={weekText(week)} density="slim" open testid={`stats-week-${week.week_of}`} />
            {/each}
          </div>
        {/if}
      </Section>
    </div>
    {#snippet rail()}
      <Section title={copy.payoutsTitle} testid="stats-payouts">
        <div class="flex flex-col gap-4">
          {#if mine.stripe_status !== 'ready'}
            <div>
              <a class={button('primary')} href="/creators/onboarding" data-testid="stats-setup"
                >{copy.setup}</a
              >
            </div>
          {/if}
          {#if mine.payouts.length === 0}
            <p class="text-muted text-[14px]" data-testid="stats-payouts-empty">{copy.payoutsEmpty}</p>
          {:else}
            <div>
              {#each mine.payouts as p (p.month)}
                <Row title={monthLabel(p.month)} density="slim" open>
                  {#snippet trail()}<span class="tabular-nums">{money(p.amount_cents)}</span>{/snippet}
                </Row>
              {/each}
            </div>
          {/if}
          {#if mine.stripe_status === 'ready'}
            <div>
              <button
                type="button"
                class={button('secondary')}
                disabled={opening}
                aria-busy={opening}
                onclick={() => void openDashboard()}
                data-testid="stats-dashboard">{opening ? copy.dashboardOpening : copy.dashboard}</button
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
      <p data-testid="stats-not-creator">{onboardingCopy.notCreator}</p>
    {:else}
      <LoadError message={copy.loadFailed} onRetry={() => void load()} testid="stats-load-error" />
    {/if}
  </Page>
{/if}
