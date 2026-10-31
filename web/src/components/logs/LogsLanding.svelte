<!-- web/src/components/logs/LogsLanding.svelte -->
<!-- The /logs landing under its band (design/specs/2026-10-11-logs.md): the shared character
     header, then two columns split by verb. Left, OPEN AND READ: your reports, the recent public
     reports (and the one sample). Right, WAYS IN: the companion above Upload when a device is
     paired, Upload first otherwise and signed out. One island, because the page's three decisions
     share one session read: whether the header shows, which way in comes first, and where the
     header's action points (the file picker on a desktop, the newest report on a phone).

     Signed out, the header is the shared band and every panel says in plain words that it needs a
     signed-in account; the page has no sign-in button. For a visitor with a session hint the right
     column keeps its box but is hidden until the device list has answered, so a paired visitor's
     panels never visibly swap places. -->

<script lang="ts">
  import {
    battlenetStartUrl,
    fetchMeOnce,
    listCompanions,
    type Device,
    type Me,
  } from '../../lib/account/api';
  import { PRIMARY_BUTTON_FIXED, SECONDARY_BUTTON_FIXED } from '../../lib/planner/styles';
  import type { HeaderAction } from '../../lib/character-header/types';
  import { logsAnswerSpec, newestOwnReport, type NewestReport } from '../../lib/logs/latest-kill';
  import { waysOrder } from '../../lib/logs/panel-order';
  import { logsCompanionCopy } from '../../lib/reports/copy';
  import { myReportsCopy } from '../../lib/reports/my-reports-copy';
  import type AccountIsland from '../Account.svelte';
  import CharacterHeader from '../character/CharacterHeader.svelte';
  import CompanionStatus from '../CompanionStatus.svelte';
  import MyReports from '../MyReports.svelte';
  import RecentReports from '../RecentReports.svelte';
  import Upload from '../Upload.svelte';
  import LogsPanel from './LogsPanel.svelte';

  /** Below `md`: the header's action reads the newest report; Upload lives only in its panel. */
  const PHONE_QUERY = '(max-width: 767px)';
  const PHONE_ACTION_LABEL = 'Open your latest report →';
  const UPLOAD_ACTION_LABEL = 'Upload a log';

  const answer = logsAnswerSpec();

  /** `undefined` while the session is being asked for; `null` signed out. */
  let me = $state<Me | null | undefined>(undefined);
  /** `null` until a signed-in visitor's device list has answered. */
  let devices = $state<Device[] | null>(null);
  let newestReport = $state<NewestReport | null>(null);
  /** The pairing block (Account.svelte's 'pairing' mode), imported only for a signed-in
   *  visitor: the account island is the site's heaviest chunk and a signed-out /logs never
   *  draws it, so it stays off the page's first-paint module graph. Loaded alongside the
   *  device list, so it is ready before the hidden right column is revealed. */
  let PairingAccount = $state<typeof AccountIsland | null>(null);
  let phone = $state(false);
  let upload = $state<Upload | undefined>(undefined);

  $effect(() => {
    const query = window.matchMedia(PHONE_QUERY);
    const sync = (): void => {
      phone = query.matches;
    };
    sync();
    query.addEventListener('change', sync);
    return () => query.removeEventListener('change', sync);
  });

  $effect(() => {
    let cancelled = false;
    void (async () => {
      const session = await fetchMeOnce().catch(() => null);
      if (cancelled) return;
      me = session;
      if (session === null) return;
      const [list, newest, pairingAccount] = await Promise.all([
        listCompanions().catch((): Device[] => []),
        newestOwnReport(),
        // A chunk that fails to load leaves the block out rather than the column hidden.
        import('../Account.svelte').then((imported) => imported.default).catch(() => null),
      ]);
      if (cancelled) return;
      PairingAccount = pairingAccount;
      devices = list;
      newestReport = newest;
    })();
    return () => {
      cancelled = true;
    };
  });

  const signedIn = $derived(me !== undefined && me !== null);
  const paired = $derived(devices !== null && devices.length > 0);
  const waysPending = $derived(me === undefined || (me !== null && devices === null));
  const order = $derived(waysOrder({ signedIn, paired }));

  const headerAction = $derived<HeaderAction>(
    !phone
      ? { kind: 'picker', label: UPLOAD_ACTION_LABEL, href: '#upload' }
      : newestReport === null
        ? { label: UPLOAD_ACTION_LABEL, href: '#upload' }
        : { label: PHONE_ACTION_LABEL, href: `/reports/${newestReport.id}` },
  );
</script>

<!-- Signed out the page has no character to show, and the sentence the shared band would say
     repeats the nav's: the header is left out (whenSignedOut="none"). -->
<CharacterHeader action={headerAction} {answer} whenSignedOut="none" onPick={() => upload?.openPicker()} />

<div class="logs-grid" data-testid="logs-grid">
  <div class="logs-column" data-testid="logs-reading">
    <LogsPanel glyph="chart" title="Your reports" id="reports">
      <div class="min-h-[88px]">
        {#if me === undefined || me === null}
          <!-- One node from the server markup through the session answer: a visitor with no session
               hint is signed out, so the sentence is the page's largest first paint, and keeping the
               same element (only the hint class drops) keeps it the stable LCP candidate. A hinted
               visitor sees the reserved box until the session answers. -->
          <div
            class="flex flex-col items-start gap-3"
            class:ch-hint-out-only={me === undefined}
            data-testid="reports-signed-out-block"
          >
            <p class="text-[14px]" data-testid="reports-signed-out">{myReportsCopy.signedOut}</p>
            <a
              class="{PRIMARY_BUTTON_FIXED} px-5"
              href={battlenetStartUrl('/logs')}
              data-testid="reports-sign-in"
            >
              {myReportsCopy.signInAction}
            </a>
          </div>
        {:else}
          <MyReports heading={false} guilds={me.guilds ?? []} uploadAction={false} />
        {/if}
      </div>
    </LogsPanel>
    <LogsPanel glyph="news" title="Recent public reports">
      <RecentReports heading={false} />
    </LogsPanel>
  </div>

  <div class="logs-column" class:logs-pending={waysPending} data-testid="logs-ways">
    {#each order as way (way)}
      {#if way === 'companion'}
        <LogsPanel glyph="device" title="The companion" aside="Live, while you raid" id="companion">
          <div class="flex flex-col items-start gap-3">
            {#if devices !== null}<CompanionStatus {devices} />{/if}
            <p class="text-[14px]" data-testid="logs-companion-pointer">{logsCompanionCopy.pointer}</p>
            <a
              class="{SECONDARY_BUTTON_FIXED} border-line-warm-strong text-strong px-4"
              href={logsCompanionCopy.pointerHref}
              data-testid="logs-companion-link"
            >
              {logsCompanionCopy.pointerLink}
            </a>
          </div>
          <!-- The pairing code is for an account that has one to pair: signed out, the card is
               the sentence and the button, sized to its content, and the code block is not drawn. A signed-in visitor gets it once the session read answers, while
               the column is still hidden (.logs-pending), so nothing visibly moves. -->
          {#if signedIn && PairingAccount !== null}
            <div class="border-line-soft flex flex-col gap-3 border-t pt-3">
              <h3 class="label text-muted">
                {paired ? logsCompanionCopy.pairAnotherLabel : logsCompanionCopy.pairLabel}
              </h3>
              <PairingAccount mode="pairing" />
            </div>
          {/if}
        </LogsPanel>
      {:else}
        <LogsPanel glyph="upload" title="Upload a log" aside="For a night already logged" id="upload">
          <Upload bind:this={upload} />
        </LogsPanel>
      {/if}
    {/each}
  </div>
</div>

<style>
  .logs-grid {
    display: grid;
    grid-template-columns: minmax(0, 1fr);
    gap: 22px;
    align-items: start;
  }
  .logs-column {
    display: flex;
    flex-direction: column;
    gap: 22px;
    min-width: 0;
  }
  /* A hinted (signed-in) visitor's right column keeps its box but is hidden until the device list
     answers, so a paired visitor's panels never visibly swap. A visitor with no hint is signed out,
     whose order is already final, so their panels paint with the server markup. */
  :global(html[data-session='1']) .logs-pending {
    visibility: hidden;
  }
  @media (min-width: 1024px) {
    .logs-grid {
      grid-template-columns: minmax(0, 7fr) minmax(0, 5fr);
      gap: 24px;
    }
    .logs-column {
      gap: 24px;
    }
  }
</style>
