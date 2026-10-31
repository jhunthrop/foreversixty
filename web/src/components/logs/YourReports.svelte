<!-- web/src/components/logs/YourReports.svelte -->
<!-- The /reports page's one card ("Your reports" in the Logs menu): every report the visitor owns,
     newest first, through the same MyReports list /logs and /account mount. Signed out it is a
     sign-in prompt; with no reports, "No logs yet" and the way to upload one; either way the sample
     report is one click away. A session that will not load says so and retries in place.

     Signed out is the server markup's own state (`me` is undefined until the session answers), so a
     visitor with no session hint paints the prompt at once and keeps the same node; a hinted visitor
     sees the list's reserve instead (`.ch-hint-out-only` and `.reserve` below). -->
<script lang="ts">
  import { fetchMeOnce, type Me } from '../../lib/account/api';
  import { SAMPLE_REPORT_PATH } from '../../lib/report/sample';
  import { yourReportsPageCopy as copy } from '../../lib/reports/my-reports-copy';
  import MyReports from '../MyReports.svelte';
  import SignInPrompt from '../SignInPrompt.svelte';
  import Card from '../ui/Card.svelte';
  import LoadError from '../ui/LoadError.svelte';

  const SAMPLE_LINK = { label: copy.sampleAction, href: SAMPLE_REPORT_PATH } as const;

  /** Undefined until the session answers; null signed out. */
  let me = $state<Me | null | undefined>(undefined);
  let failed = $state(false);

  async function load(): Promise<void> {
    failed = false;
    try {
      me = await fetchMeOnce();
    } catch {
      failed = true;
    }
  }

  $effect(() => {
    void load();
  });
</script>

<Card glyph="chart" title={copy.cardTitle} testid="your-reports">
  <div class:reserve={me === undefined && !failed}>
    {#if failed}
      <LoadError message={copy.sessionFailed} onRetry={() => void load()} testid="your-reports-error" />
    {:else if me === undefined || me === null}
      <div class="flex flex-col items-start gap-4" class:ch-hint-out-only={me === undefined}>
        <SignInPrompt line={copy.signedOut} next="/reports" testid="your-reports-signin" />
        <a
          class="text-muted inline-flex min-h-11 items-center text-[13px] underline md:min-h-0"
          href={SAMPLE_LINK.href}
          data-testid="your-reports-sample">{SAMPLE_LINK.label}</a
        >
      </div>
    {:else}
      <MyReports
        heading={false}
        guilds={me.guilds ?? []}
        nightDates
        emptyMessage={copy.empty}
        emptySecondary={SAMPLE_LINK}
      />
    {/if}
  </div>
</Card>

<style>
  /* Until the session answers, a visitor whose hint says signed in gets the list's loading height
     (REPORTS_LOADING_MIN_H in lib/reports/layout.ts, which MyReports' skeleton then holds), so the
     rows land without moving the footer. */
  :global(html[data-session='1']) .reserve {
    min-height: 360px;
  }
</style>
