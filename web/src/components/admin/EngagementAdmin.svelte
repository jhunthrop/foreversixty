<!-- web/src/components/admin/EngagementAdmin.svelte -->
<!-- The owner's /admin/activity: who signs in and who comes back. A ledger Page: this week's
     headline figures with what the numbers leave out, then the weeks, retention by signup week and
     the days, each a Grid on a flush panel. A figure from before the live record is a floor (≥).
     Anyone but the owner sees one line saying so: the API answers 401 or 403 for them whatever
     this page shows. -->
<script lang="ts">
  import { AccountError, fetchMeOnce } from '../../lib/account/api';
  import {
    adminActivityCopy as copy,
    cohortRows,
    dayRows,
    headlineFigures,
    weekRows,
  } from '../../lib/engagement/admin';
  import { fetchEngagement, type EngagementReport } from '../../lib/engagement/api';
  import SignInPrompt from '../SignInPrompt.svelte';
  import Figure from '../ui/Figure.svelte';
  import Grid from '../ui/Grid.svelte';
  import LoadError from '../ui/LoadError.svelte';
  import Page from '../ui/Page.svelte';
  import Section from '../ui/Section.svelte';
  import Skeleton from '../ui/Skeleton.svelte';

  /** Each Grid: capped so the figures stay near their labels on a wide screen, scrolled sideways
   *  inside its panel on a phone rather than widening the page (inline-size containment keeps the
   *  table's own width from pushing the panel and the page wider). */
  const GRID_BOX = 'w-full max-w-[760px] overflow-x-auto [contain:inline-size]';

  type Phase = 'loading' | 'signed-out' | 'not-owner' | 'failed' | 'ready';

  let phase = $state<Phase>('loading');
  let report = $state<EngagementReport | null>(null);

  function phaseForError(error: unknown): Phase {
    if (error instanceof AccountError && error.status === 401) return 'signed-out';
    if (error instanceof AccountError && error.status === 403) return 'not-owner';
    return 'failed';
  }

  async function load(): Promise<void> {
    phase = 'loading';
    try {
      const me = await fetchMeOnce();
      if (me === null) {
        phase = 'signed-out';
        return;
      }
      if (me.user.role !== 'admin') {
        phase = 'not-owner';
        return;
      }
      report = await fetchEngagement();
      phase = report === null ? 'failed' : 'ready';
    } catch (error) {
      phase = phaseForError(error);
    }
  }

  $effect(() => {
    void load();
  });
</script>

{#if phase === 'loading'}
  <div data-testid="activity-loading"><Skeleton lines={8} rowHeight="h-6" minHeight="min-h-[420px]" /></div>
{:else if phase === 'signed-out'}
  <div data-testid="activity-signed-out">
    <SignInPrompt line={copy.signedOut} next="/admin/activity" testid="activity-signin" />
  </div>
{:else if phase === 'not-owner'}
  <p class="text-[15px]" data-testid="activity-not-owner">{copy.notOwner}</p>
{:else if phase === 'failed' || report === null}
  <div data-testid="activity-error">
    <LoadError message={copy.loadFailed} onRetry={() => void load()} testid="activity-load-error" />
  </div>
{:else}
  <Page layout="ledger" testid="activity-page-body">
    <Section title={copy.nowTitle} panel testid="activity-now">
      <div class="flex flex-col gap-4">
        <div class="grid grid-cols-2 gap-x-6 gap-y-4 sm:grid-cols-5" data-testid="activity-figures">
          {#each headlineFigures(report).flat() as figure (figure.label)}
            <Figure value={figure.value} caption={figure.label} size="big" />
          {/each}
        </div>
        <div class="text-muted text-small flex flex-col gap-1">
          <p class="m-0" data-testid="activity-floor">{copy.floor(report.live_since)}</p>
          <p class="m-0" data-testid="activity-internal">{copy.internal(report.internal_accounts)}</p>
          <p class="m-0">{copy.surfaces}</p>
        </div>
      </div>
    </Section>
    <Section title={copy.weeksTitle} meta={copy.weeksMeta} panel flush testid="activity-weeks">
      <div class={GRID_BOX}>
        <Grid
          caption={copy.weeksCaption}
          labelHeader={copy.weekHeader}
          columns={copy.periodColumns}
          rows={weekRows(report.weeks)}
          edge
          testid="activity-weeks-grid"
        />
      </div>
    </Section>
    <Section title={copy.cohortsTitle} meta={copy.cohortsMeta} panel flush testid="activity-cohorts">
      <div class={GRID_BOX}>
        <Grid
          caption={copy.cohortsCaption}
          labelHeader={copy.cohortHeader}
          columns={copy.cohortColumns}
          rows={cohortRows(report.cohorts)}
          edge
          testid="activity-cohorts-grid"
        />
      </div>
    </Section>
    <Section title={copy.daysTitle} meta={copy.weeksMeta} panel flush testid="activity-days">
      <div class={GRID_BOX}>
        <Grid
          caption={copy.daysCaption}
          labelHeader={copy.dayHeader}
          columns={copy.periodColumns}
          rows={dayRows(report.days)}
          edge
          testid="activity-days-grid"
        />
      </div>
    </Section>
  </Page>
{/if}
