<!-- web/src/components/admin/AppTestersAdmin.svelte -->
<!-- The owner's /admin/testers (mock web-admin-testers, design/mocks/ws_app_beta.py): a ledger
     Page with one Section, the copy buttons (one per platform, its count in the label), then the
     Grid of address, phone and day. Anyone but the owner sees one line saying so: the API answers
     401 or 403 for them whatever this page shows. -->
<script lang="ts">
  import EmptyState from '../ui/EmptyState.svelte';
  import { AccountError, fetchMeOnce } from '../../lib/account/api';
  import {
    adminTestersCopy as copy,
    countsByPlatform,
    emailsFor,
    PLATFORM_ORDER,
    testerRows,
  } from '../../lib/app-beta/admin';
  import { listAppTesters, type AppTester } from '../../lib/app-beta/api';
  import type { Platform } from '../../lib/app-beta/platform';
  import { smallButton } from '../../lib/account/small-button';
  import { createCopyFeedback } from '../../lib/ui/copy-feedback';
  import SignInPrompt from '../SignInPrompt.svelte';
  import Grid from '../ui/Grid.svelte';
  import LoadError from '../ui/LoadError.svelte';
  import Page from '../ui/Page.svelte';
  import Section from '../ui/Section.svelte';
  import Skeleton from '../ui/Skeleton.svelte';

  type Phase = 'loading' | 'signed-out' | 'not-owner' | 'failed' | 'ready';

  let phase = $state<Phase>('loading');
  let testers = $state<AppTester[]>([]);
  const rows = $derived(testerRows(testers));
  const counts = $derived(countsByPlatform(testers));

  let copied = $state<Platform | null>(null);
  let copyFailed = $state(false);
  let copying: Platform | null = null;
  const copier = createCopyFeedback((state) => {
    copied = state === 'copied' ? copying : null;
    copyFailed = state === 'failed';
  });
  $effect(() => copier.dispose);

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
      testers = await listAppTesters();
      phase = 'ready';
    } catch (error) {
      phase = phaseForError(error);
    }
  }

  $effect(() => {
    void load();
  });

  function copyEmails(platform: Platform): void {
    copying = platform;
    void copier.copy(emailsFor(testers, platform));
  }
</script>

{#if phase === 'loading'}
  <div data-testid="testers-loading"><Skeleton lines={6} rowHeight="h-6" minHeight="min-h-[320px]" /></div>
{:else if phase === 'signed-out'}
  <div data-testid="testers-signed-out">
    <SignInPrompt line={copy.signedOut} next="/admin/testers" testid="testers-signin" />
  </div>
{:else if phase === 'not-owner'}
  <p class="text-[15px]" data-testid="testers-not-owner">{copy.notOwner}</p>
{:else if phase === 'failed'}
  <div data-testid="testers-error">
    <LoadError message={copy.loadFailed} onRetry={() => void load()} testid="testers-load-error" />
  </div>
{:else}
  <Page layout="ledger" testid="testers-page-body">
    <Section title={copy.tableTitle} panel testid="testers-table">
      <div class="flex flex-col gap-4">
        {#if rows.length === 0}
          <EmptyState message={copy.empty} testid="testers-empty" />
        {:else}
          <div class="flex flex-wrap items-center gap-3">
            {#each PLATFORM_ORDER as platform (platform)}
              <button
                type="button"
                class={smallButton()}
                disabled={counts[platform] === 0}
                onclick={() => copyEmails(platform)}
                data-testid={`testers-copy-${platform}`}
              >
                {copied === platform ? copy.copied : copy.copyEmails(platform, counts[platform])}
              </button>
            {/each}
            {#if copyFailed}<p class="text-wipe text-small m-0" role="alert">{copy.copyFailed}</p>{/if}
          </div>
          <Grid
            caption={copy.caption}
            labelHeader={copy.emailHeader}
            columns={copy.columns}
            {rows}
            testid="testers-grid"
          />
        {/if}
      </div>
    </Section>
  </Page>
{/if}
