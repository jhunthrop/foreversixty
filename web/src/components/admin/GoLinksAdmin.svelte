<!-- web/src/components/admin/GoLinksAdmin.svelte -->
<!-- The owner's /admin/links: every tracked link (foreversixty.gg/go/<slug>) with its funnel,
     clicks to the accounts it brought to the accounts with a character to Premium, each step
     with its share of the one before; and the form that makes a new link. A split Page: the
     links panel holds a Row per link (its four Figures and Copy in the trail, under the name
     below 768) and the All links Row; the rail is the New link panel (Fields). Anyone but the
     owner sees one line saying so: the API answers 403 for them whatever this page shows. -->
<script lang="ts">
  import EmptyState from '../ui/EmptyState.svelte';
  import { AccountError, fetchMeOnce } from '../../lib/account/api';
  import { adminLinksCopy as copy } from '../../lib/golinks/admin-copy';
  import { createGoLink, listGoLinks, type GoLinkFunnel } from '../../lib/golinks/api';
  import {
    countLabel,
    funnelTotals,
    shareOf,
    shortLink,
    sortedByClicks,
    type FunnelFigures,
  } from '../../lib/golinks/funnel';
  import { smallButton } from '../../lib/account/small-button';
  import { button } from '../../lib/ui/button';
  import { createCopyFeedback } from '../../lib/ui/copy-feedback';
  import SignInPrompt from '../SignInPrompt.svelte';
  import Field from '../ui/Field.svelte';
  import Figure from '../ui/Figure.svelte';
  import LoadError from '../ui/LoadError.svelte';
  import Page from '../ui/Page.svelte';
  import Row from '../ui/Row.svelte';
  import Section from '../ui/Section.svelte';
  import Skeleton from '../ui/Skeleton.svelte';

  type Phase = 'loading' | 'signed-out' | 'not-owner' | 'failed' | 'ready';

  let phase = $state<Phase>('loading');
  let links = $state<GoLinkFunnel[]>([]);
  const rows = $derived(sortedByClicks(links));
  const totals = $derived(funnelTotals(links));

  const EMPTY_FORM = { slug: '', target: '/', label: '' };
  let values = $state<Record<'slug' | 'target' | 'label', string>>({ ...EMPTY_FORM });
  let saving = $state(false);
  let fieldErrors = $state<Readonly<Record<string, string>>>({});
  let formError = $state('');
  let notice = $state('');

  let copiedSlug = $state<string | null>(null);
  let copyingSlug: string | null = null;
  const copier = createCopyFeedback((state) => (copiedSlug = state === 'copied' ? copyingSlug : null));
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
      links = await listGoLinks();
      phase = 'ready';
    } catch (error) {
      phase = phaseForError(error);
    }
  }

  $effect(() => {
    void load();
  });

  async function onCreate(event: SubmitEvent): Promise<void> {
    event.preventDefault();
    saving = true;
    fieldErrors = {};
    formError = '';
    notice = '';
    try {
      const made = await createGoLink({
        slug: values.slug.trim(),
        target: values.target.trim(),
        label: values.label.trim(),
      });
      notice = copy.created(made.slug);
      values = { ...EMPTY_FORM };
      links = await listGoLinks();
    } catch (error) {
      if (error instanceof AccountError && error.fields !== null) fieldErrors = error.fields;
      else formError = error instanceof AccountError ? error.message : copy.createFailed;
    } finally {
      saving = false;
    }
  }

  function copyLink(rowSlug: string): void {
    copyingSlug = rowSlug;
    void copier.copy(shortLink(rowSlug));
  }

  interface Step {
    key: string;
    label: string;
    value: string;
    /** The line under the figure: the bots left out of the clicks, or the share of the step before. */
    note: string;
  }

  function shareNote(part: number, whole: number, basis: string): string {
    const share = shareOf(part, whole);
    return whole > 0 ? `${share} ${basis}` : share;
  }

  /** The four figures a row and the totals line show, each over the line that explains it. */
  function steps(f: FunnelFigures): Step[] {
    const nothingYet = f.clicks === 0 && f.bot_clicks === 0;
    return [
      {
        key: 'clicks',
        label: copy.columns.clicks,
        value: countLabel(f.clicks),
        note: nothingYet ? copy.noClicks : copy.bots(f.bot_clicks),
      },
      {
        key: 'sign-ins',
        label: copy.columns.signIns,
        value: countLabel(f.sign_ins),
        note: shareNote(f.sign_ins, f.clicks, copy.shareOf.signIns),
      },
      {
        key: 'linked',
        label: copy.columns.linked,
        value: countLabel(f.linked),
        note: shareNote(f.linked, f.sign_ins, copy.shareOf.linked),
      },
      {
        key: 'premium',
        label: copy.columns.premium,
        value: countLabel(f.premium),
        note: shareNote(f.premium, f.linked, copy.shareOf.premium),
      },
    ];
  }

  const FIELDS = [
    { id: 'slug', label: copy.slugLabel, hint: copy.slugHint, placeholder: 'reddit-classicwow' },
    { id: 'target', label: copy.targetLabel, hint: undefined, placeholder: '/' },
    { id: 'label', label: copy.labelLabel, hint: undefined, placeholder: 'Reddit: r/classicwow' },
  ] as const;
</script>

{#snippet copyButton(rowSlug: string)}
  <button
    type="button"
    class="{smallButton()} min-w-[76px] shrink-0"
    aria-label={copy.copyLabel(rowSlug)}
    onclick={() => copyLink(rowSlug)}
    data-testid={`links-copy-${rowSlug}`}
  >
    {copiedSlug === rowSlug ? copy.copied : copy.copy}
  </button>
{/snippet}

{#snippet figures(f: FunnelFigures)}
  <!-- Two by two on a phone, so a link is two figure rows, not a column of four. -->
  <span class="grid min-w-0 grid-cols-2 items-start gap-x-5 gap-y-2 max-md:flex-1 md:flex md:flex-wrap">
    {#each steps(f) as step (step.key)}
      <span class="md:w-[128px]"
        ><Figure value={step.value} label={step.label} caption={step.note} size="sm" /></span
      >
    {/each}
  </span>
{/snippet}

{#if phase === 'loading'}
  <div data-testid="links-loading"><Skeleton lines={6} rowHeight="h-6" minHeight="min-h-[320px]" /></div>
{:else if phase === 'signed-out'}
  <div data-testid="links-signed-out">
    <SignInPrompt line={copy.signedOut} next="/admin/links" testid="links-signin" />
  </div>
{:else if phase === 'not-owner'}
  <p class="text-[15px]" data-testid="links-not-owner">{copy.notOwner}</p>
{:else if phase === 'failed'}
  <div data-testid="links-error">
    <LoadError message={copy.loadFailed} onRetry={() => void load()} testid="links-load-error" />
  </div>
{:else}
  <Page layout="split" panels testid="links-page-body">
    <Section title={copy.tableTitle} meta={copy.tableMeta(rows.length)} panel flush testid="links-table">
      <div>
        {#if rows.length === 0}
          <EmptyState inset message={copy.empty} testid="links-empty" />
        {:else}
          <div role="group" aria-label={copy.tableTitle}>
            {#each rows as row (row.slug)}
              <Row title={row.label} wrapTrail testid={`links-row-${row.slug}`}>
                {#snippet subContent()}
                  <span class="text-text font-mono">/go/{row.slug}</span>
                  {#if row.target !== '/'}<span aria-hidden="true"> → </span><span>{row.target}</span>{/if}
                {/snippet}
                {#snippet trail()}
                  {@render figures(row)}
                  {@render copyButton(row.slug)}
                {/snippet}
              </Row>
            {/each}
            <Row title={copy.totals} wrapTrail testid="links-totals">
              {#snippet trail()}
                {@render figures(totals)}
                <span class="w-[76px] shrink-0" aria-hidden="true"></span>
              {/snippet}
            </Row>
          </div>
        {/if}
      </div>
    </Section>
    {#snippet rail()}
      <Section title={copy.formTitle} panel testid="links-form-card">
        <div class="flex flex-col gap-3">
          <p class="text-muted text-[14px]">{copy.formJob}</p>
          <form class="flex flex-col gap-3" onsubmit={onCreate} novalidate data-testid="links-form">
            {#each FIELDS as field (field.id)}
              <Field
                id={`link-${field.id}`}
                label={field.label}
                placeholder={field.placeholder}
                hint={field.hint}
                error={fieldErrors[field.id]}
                autocomplete="off"
                bind:value={values[field.id]}
                testid={`links-${field.id}`}
              />
            {/each}
            <button
              type="submit"
              class="{button('primary')} self-start"
              disabled={saving}
              aria-busy={saving}
              data-testid="links-create"
            >
              {saving ? copy.creating : copy.create}
            </button>
            <div class="min-h-[22px]" aria-live="polite">
              {#if notice !== ''}<p class="text-[14px]" data-testid="links-created">{notice}</p>{/if}
              {#if formError !== ''}<p
                  class="text-wipe text-[14px]"
                  role="alert"
                  data-testid="links-form-error"
                >
                  {formError}
                </p>{/if}
            </div>
          </form>
        </div>
      </Section>
    {/snippet}
  </Page>
{/if}
