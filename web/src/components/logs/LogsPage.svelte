<!-- web/src/components/logs/LogsPage.svelte -->
<!-- /logs (logs spec 2026-11-07 section 4), built from the ui primitives: the PageHeader (signed in
     with a report, its answer is the latest report), then a split Page whose one Plate is the next
     step (decision L1): signed out or with no logs, the upload; with logs, Your reports (the upload
     open below it). Recent public reports follow in main; the rail holds the companion (signed in,
     with pairing) and, signed out, the Your reports line. The shared character band left /logs
     (L2). One island, because every region follows one session read and one list read.

     The server markup is the signed-out page, final for a visitor with no session hint; a hinted
     visitor's body stays hidden (its box kept) until the session and the list answer. -->
<script lang="ts">
  import {
    fetchMeOnce,
    listCompanions,
    listMyReports,
    type Device,
    type Me,
    type MyReport,
  } from '../../lib/account/api';
  import { relativeDay } from '../../lib/dates';
  import { ownReportRow } from '../../lib/logs/report-summary';
  import { logsBandCopy } from '../../lib/reports/copy';
  import { myReportsCopy } from '../../lib/reports/my-reports-copy';
  import { LINK } from '../../lib/ui/primitives';
  import MyReports from '../MyReports.svelte';
  import RecentReports from '../RecentReports.svelte';
  import Upload from '../Upload.svelte';
  import Figure from '../ui/Figure.svelte';
  import Page from '../ui/Page.svelte';
  import PageHeader from '../ui/PageHeader.svelte';
  import Plate from '../ui/Plate.svelte';
  import Section from '../ui/Section.svelte';
  import LogsCompanion from './LogsCompanion.svelte';

  const UPLOAD_TITLE = 'Upload a log';
  const UPLOAD_META = 'For a night already logged';
  const REPORTS_TITLE = 'Your reports';

  /** `undefined` while the session is asked for; `null` signed out. */
  let me = $state<Me | null | undefined>(undefined);
  /** The newest page of the visitor's own reports; `null` until it answers (or when it fails). */
  let mine = $state<{ rows: MyReport[]; total: number } | null>(null);
  let devices = $state<Device[]>([]);
  let settled = $state(false);

  $effect(() => {
    let cancelled = false;
    void (async () => {
      const session = await fetchMeOnce().catch(() => null);
      if (cancelled) return;
      me = session;
      if (session !== null) {
        const [list, companions] = await Promise.all([
          listMyReports(1).catch(() => null),
          listCompanions().catch((): Device[] => []),
        ]);
        if (cancelled) return;
        mine = list;
        devices = companions;
      }
      settled = true;
    })();
    return () => {
      cancelled = true;
    };
  });

  const signedIn = $derived(me !== undefined && me !== null);
  const newest = $derived(mine?.rows[0] ?? null);
  const hasReports = $derived(signedIn && newest !== null);
  const latest = $derived(
    !hasReports || newest === null
      ? null
      : {
          ...ownReportRow({
            title: newest.title,
            zone: newest.zone,
            startedAt: newest.started_at ?? newest.created_at,
            fightCount: newest.fight_count,
            killCount: newest.kill_count,
            bossCount: newest.boss_count,
          }),
          age: relativeDay(new Date(newest.created_at)),
          href: `/reports/${newest.id}`,
        },
  );
  /** "0 bosses · 42 fights": the row's own line without its credit. */
  const latestWhere = $derived(latest === null ? '' : latest.sub.split(' · ').slice(0, -1).join(' · '));
  const total = $derived(mine?.total ?? 0);
</script>

{#snippet latestAnswer()}
  {#if latest !== null}
    <!-- the answer column holds 360 px from 900, as the mock's -->
    <span class="hidden min-[900px]:block min-[900px]:w-[360px]" aria-hidden="true"></span>
    <span class="text-label tracking-label text-muted font-bold uppercase"
      >Your latest report · {latest.age}</span
    >
    <Figure value={latest.title} size="big" testid="logs-answer" />
    <span class="text-small text-muted">{latestWhere}</span>
    <a class={LINK} href={latest.href} data-testid="logs-answer-open">Open it ›</a>
  {/if}
{/snippet}

{#snippet uploadPlate()}
  <Section title={UPLOAD_TITLE} meta={UPLOAD_META} anchor="upload" testid="logs-upload">
    <Plate label={UPLOAD_TITLE}><Upload plate /></Plate>
  </Section>
{/snippet}

<div class="logs-body" class:logs-wait={!settled} data-testid="logs-page-body">
  <Page layout="split" testid="logs-grid">
    {#snippet header()}
      <PageHeader title={logsBandCopy.h1} answer={latest === null ? undefined : latestAnswer}>
        {#snippet meta()}<span>{logsBandCopy.jobLine}</span>{/snippet}
      </PageHeader>
    {/snippet}

    {#if hasReports}
      <Section
        title={REPORTS_TITLE}
        meta={`${total} report${total === 1 ? '' : 's'}, newest first`}
        action={{ label: 'All your reports', href: '/reports' }}
        anchor="reports"
      >
        <MyReports guilds={me?.guilds ?? []} />
      </Section>
      <Section title={UPLOAD_TITLE} meta={UPLOAD_META} anchor="upload" testid="logs-upload">
        <Upload />
      </Section>
    {:else if signedIn}
      {@render uploadPlate()}
      <Section title={REPORTS_TITLE} anchor="reports">
        <MyReports guilds={me?.guilds ?? []} />
      </Section>
    {:else}
      {@render uploadPlate()}
    {/if}
    <RecentReports />

    {#snippet rail()}
      <LogsCompanion {signedIn} {devices} />
      {#if !signedIn}
        <Section title={REPORTS_TITLE} anchor="reports">
          <p class="text-muted p-4 text-[14px] leading-5" data-testid="reports-signed-out">
            {myReportsCopy.signedOut}
          </p>
        </Section>
      {/if}
    {/snippet}
  </Page>
</div>

<style>
  /* A hinted (signed-in) visitor's body keeps its box but is hidden until the reads answer, so
     the signed-out server markup never flashes before the signed-in page replaces it. */
  :global(html[data-session='1']) .logs-wait {
    visibility: hidden;
  }
</style>
