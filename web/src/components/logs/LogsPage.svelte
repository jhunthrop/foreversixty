<!-- web/src/components/logs/LogsPage.svelte -->
<!-- /logs (logs spec 2026-11-07 section 4; directions b1-logs), built from the ui primitives: the
     title band, then a split Page of panels. Main: Upload a log (the page's lead panel), Your
     reports (signed in; All › to /reports), Recent public reports (the viewer's own left out, as
     they sit just above). Rail: the companion and, signed out, the Your reports line. One island,
     because every region follows one session read and one list read.

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
  import { logsBandCopy } from '../../lib/reports/copy';
  import { myReportsCopy } from '../../lib/reports/my-reports-copy';
  import MyReports from '../MyReports.svelte';
  import RecentReports from '../RecentReports.svelte';
  import Upload from '../Upload.svelte';
  import Page from '../ui/Page.svelte';
  import PageHeader from '../ui/PageHeader.svelte';
  import Section from '../ui/Section.svelte';
  import LogsCompanion from './LogsCompanion.svelte';

  const UPLOAD_TITLE = 'Upload a log';
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
  const ownIds = $derived(new Set((mine?.rows ?? []).map((row) => row.id)));
</script>

<div class="logs-body" class:logs-wait={!settled} data-testid="logs-page-body">
  <Page layout="split" panels testid="logs-grid">
    {#snippet header()}
      <PageHeader title={logsBandCopy.h1} />
    {/snippet}

    <Section title={UPLOAD_TITLE} anchor="upload" panel frame="lead" testid="logs-upload">
      <Upload />
    </Section>
    {#if signedIn}
      <Section
        title={REPORTS_TITLE}
        action={(mine?.total ?? 0) > 0 ? { label: 'All', href: '/reports' } : undefined}
        anchor="reports"
        panel
      >
        <MyReports guilds={me?.guilds ?? []} open />
      </Section>
    {/if}
    <RecentReports exclude={ownIds} />

    {#snippet rail()}
      <LogsCompanion {signedIn} {devices} />
      {#if !signedIn}
        <Section title={REPORTS_TITLE} anchor="reports" panel>
          <p class="text-muted py-2 text-[14px] leading-5" data-testid="reports-signed-out">
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
