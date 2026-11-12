<!-- web/src/components/logs/YourReports.svelte -->
<!-- /reports, "Your reports" (logs spec 2026-11-07 section 5): the PageHeader (the title and,
     signed in, `Upload a log`), then the ledger Page's one Plate holding
     every state: signed out, the sign-in sentence over the primary and the sample report; signed
     in, MyReports (rows, `No logs yet` with Upload a log and the sample, or the failed line); a
     session that will not load says so and retries in place.

     Signed out is the server markup's own state (`me` is undefined until the session answers), so
     a visitor with no session hint paints it at once; a hinted visitor's Plate keeps the list's
     reserve until the answer. -->
<script lang="ts">
  import { battlenetStartUrl, fetchMeOnce, type Me } from '../../lib/account/api';
  import { SAMPLE_REPORT_PATH } from '../../lib/report/sample';
  import { myReportsCopy, yourReportsPageCopy as copy } from '../../lib/reports/my-reports-copy';
  import { LINK } from '../../lib/ui/primitives';
  import MyReports from '../MyReports.svelte';
  import Button from '../ui/Button.svelte';
  import LoadError from '../ui/LoadError.svelte';
  import Page from '../ui/Page.svelte';
  import PageHeader from '../ui/PageHeader.svelte';
  import Plate from '../ui/Plate.svelte';

  const SAMPLE_LINK = { label: copy.sampleAction, href: SAMPLE_REPORT_PATH } as const;
  const UPLOAD_LINK = { label: myReportsCopy.uploadAction, href: '/logs#upload' } as const;

  /** Undefined until the session answers; null signed out. */
  let me = $state<Me | null | undefined>(undefined);
  let failed = $state(false);

  async function load(): Promise<void> {
    failed = false;
    try {
      me = await fetchMeOnce();
    } catch {
      failed = true;
      return;
    }
  }

  $effect(() => {
    void load();
  });

  const signedIn = $derived(me !== undefined && me !== null);
</script>

<Page layout="ledger" testid="your-reports">
  {#snippet header()}
    <PageHeader title={copy.title} actions={signedIn ? upload : undefined} />
  {/snippet}

  {#if signedIn}
    <MyReports
      plate={copy.title}
      guilds={me?.guilds ?? []}
      emptyMessage={copy.empty}
      emptyAction={UPLOAD_LINK}
      emptySecondary={SAMPLE_LINK}
    />
  {:else}
    <Plate label={copy.title}>
      <div class="reserve text-muted p-4 text-[14px] leading-5" class:hinted={me === undefined && !failed}>
        {#if failed}
          <LoadError message={copy.sessionFailed} onRetry={() => void load()} testid="your-reports-error" />
        {:else}
          <div class:ch-hint-out-only={me === undefined} data-testid="your-reports-signin">
            <p>{copy.signedOut}</p>
            <div class="mt-3 flex flex-wrap items-center gap-4">
              <Button tone="primary" href={battlenetStartUrl('/reports')}>{myReportsCopy.signInAction}</Button
              >
              <a class={LINK} href={SAMPLE_LINK.href} data-testid="your-reports-sample"
                >{SAMPLE_LINK.label} ›</a
              >
            </div>
          </div>
        {/if}
      </div>
    </Plate>
  {/if}
</Page>

{#snippet upload()}
  <Button tone="primary" href={UPLOAD_LINK.href} testid="your-reports-upload">{UPLOAD_LINK.label}</Button>
{/snippet}

<style>
  /* Until the session answers, a visitor whose hint says signed in gets the list's reserve. */
  :global(html[data-session='1']) .reserve.hinted {
    min-height: 200px;
  }
</style>
