<!-- web/src/components/admin/GoLinksAdmin.svelte -->
<!-- The owner's /admin/links: every tracked link (foreversixty.gg/go/<slug>) with its funnel,
     clicks to the accounts it brought to the accounts with a character to Premium, each step
     with its share of the one before; and the form that makes a new link. From 1024 the form
     is a rail beside the table; below it the form follows. Below 768 each link is a block of
     four figures instead of a table row. Anyone but the owner sees one line saying so: the API
     answers 403 for them whatever this page shows. -->
<script lang="ts">
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
  import { button } from '../../lib/ui/button';
  import { createCopyFeedback } from '../../lib/ui/copy-feedback';
  import SignInPrompt from '../SignInPrompt.svelte';
  import Card from '../ui/Card.svelte';
  import LoadError from '../ui/LoadError.svelte';
  import Skeleton from '../ui/Skeleton.svelte';

  type Phase = 'loading' | 'signed-out' | 'not-owner' | 'failed' | 'ready';

  let phase = $state<Phase>('loading');
  let links = $state<GoLinkFunnel[]>([]);
  const rows = $derived(sortedByClicks(links));
  const totals = $derived(funnelTotals(links));

  let slug = $state('');
  let target = $state('/');
  let label = $state('');
  let saving = $state(false);
  let fieldErrors = $state<Readonly<Record<string, string>>>({});
  let formError = $state('');
  let notice = $state('');

  let copiedSlug = $state<string | null>(null);
  let copyingSlug: string | null = null;
  const copier = createCopyFeedback(
    (text) => navigator.clipboard.writeText(text),
    (state) => (copiedSlug = state === 'copied' ? copyingSlug : null),
  );
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
      const made = await createGoLink({ slug: slug.trim(), target: target.trim(), label: label.trim() });
      notice = copy.created(made.slug);
      slug = '';
      target = '/';
      label = '';
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

  const FIELD =
    'border-line-warm bg-bg rounded-control text-text h-11 w-full min-w-0 px-3 text-[15px] aria-invalid:border-wipe';
</script>

{#snippet copyButton(rowSlug: string)}
  <button
    type="button"
    class="{button('secondary', 'sm')} min-w-[76px] shrink-0"
    aria-label={copy.copyLabel(rowSlug)}
    onclick={() => copyLink(rowSlug)}
    data-testid={`links-copy-${rowSlug}`}
  >
    {copiedSlug === rowSlug ? copy.copied : copy.copy}
  </button>
{/snippet}

{#snippet linkCell(row: GoLinkFunnel)}
  <div class="flex min-w-0 flex-col gap-1">
    <span class="text-strong text-[15px] font-semibold">{row.label}</span>
    <span class="text-muted flex min-w-0 flex-wrap items-center gap-x-2 text-[13px]">
      <span class="text-text font-mono">/go/{row.slug}</span>
      {#if row.target !== '/'}
        <span aria-hidden="true">→</span>
        <span class="truncate">{row.target}</span>
      {/if}
    </span>
  </div>
{/snippet}

{#if phase === 'loading'}
  <Card glyph="share" title={copy.tableTitle} testid="links-loading">
    <Skeleton lines={6} rowHeight="h-6" minHeight="min-h-[320px]" />
  </Card>
{:else if phase === 'signed-out'}
  <Card glyph="lock" title={copy.tableTitle} testid="links-signed-out">
    <SignInPrompt line={copy.signedOut} next="/admin/links" testid="links-signin" />
  </Card>
{:else if phase === 'not-owner'}
  <Card glyph="lock" title={copy.tableTitle} testid="links-not-owner">
    <p>{copy.notOwner}</p>
  </Card>
{:else if phase === 'failed'}
  <Card glyph="share" title={copy.tableTitle} testid="links-error">
    <LoadError message={copy.loadFailed} onRetry={() => void load()} testid="links-load-error" />
  </Card>
{:else}
  <div class="grid items-start gap-[22px] lg:grid-cols-[minmax(0,1fr)_360px] lg:gap-8">
    <Card
      glyph="share"
      title={copy.tableTitle}
      meta={copy.tableMeta(rows.length)}
      body={rows.length === 0 ? 'text' : 'none'}
      testid="links-table"
    >
      {#if rows.length === 0}
        <p data-testid="links-empty">{copy.empty}</p>
      {:else}
        <!-- From 768: one table row per link. -->
        <table class="border-line-soft hidden w-full border-t text-[14px] md:table">
          <thead>
            <tr class="border-line-soft border-b">
              <th scope="col" class="quiet-label px-5 py-3 text-left">{copy.columns.link}</th>
              {#each steps(totals) as step (step.key)}
                <th scope="col" class="quiet-label px-3 py-3 text-right last:pr-5">{step.label}</th>
              {/each}
              <th scope="col" class="w-[84px] pr-5"><span class="sr-only">{copy.copy}</span></th>
            </tr>
          </thead>
          <tbody>
            {#each rows as row (row.slug)}
              <tr class="border-line-soft border-b" data-testid={`links-row-${row.slug}`}>
                <td class="px-5 py-3 align-top">{@render linkCell(row)}</td>
                {#each steps(row) as step (step.key)}
                  <td class="numeric px-3 py-3 text-right align-top">
                    <span class="text-strong block text-[15px] font-semibold">{step.value}</span>
                    <span class="text-muted block text-[12px]">
                      {step.note}
                    </span>
                  </td>
                {/each}
                <td class="py-2 pr-5 text-right align-top">{@render copyButton(row.slug)}</td>
              </tr>
            {/each}
          </tbody>
          <tfoot>
            <tr data-testid="links-totals">
              <th scope="row" class="text-strong px-5 py-3 text-left text-[14px] font-semibold"
                >{copy.totals}</th
              >
              {#each steps(totals) as step (step.key)}
                <td class="numeric px-3 py-3 text-right align-top">
                  <span class="text-strong block text-[15px] font-semibold">{step.value}</span>
                  <span class="text-muted block text-[12px]">{step.note}</span>
                </td>
              {/each}
              <td></td>
            </tr>
          </tfoot>
        </table>

        <!-- Below 768: each link is a block, its four figures in a row under it. -->
        <ul class="border-line-soft flex flex-col border-t md:hidden" data-testid="links-list">
          {#each rows as row (row.slug)}
            <li
              class="border-line-soft flex flex-col gap-3 border-b px-[18px] py-4"
              data-testid={`links-item-${row.slug}`}
            >
              <div class="flex items-start justify-between gap-3">
                {@render linkCell(row)}
                {@render copyButton(row.slug)}
              </div>
              <dl class="grid grid-cols-4 gap-2">
                {#each steps(row) as step (step.key)}
                  <div class="flex flex-col gap-1">
                    <dt class="quiet-label text-[10px]">{step.label}</dt>
                    <dd class="numeric">
                      <span class="text-strong block text-[15px] font-semibold">{step.value}</span>
                      <span class="text-muted block text-[12px]">{step.note}</span>
                    </dd>
                  </div>
                {/each}
              </dl>
            </li>
          {/each}
          <li class="flex flex-col gap-3 px-[18px] py-4" data-testid="links-totals-phone">
            <span class="text-strong text-[14px] font-semibold">{copy.totals}</span>
            <dl class="grid grid-cols-4 gap-2">
              {#each steps(totals) as step (step.key)}
                <div class="flex flex-col gap-1">
                  <dt class="quiet-label text-[10px]">{step.label}</dt>
                  <dd class="numeric">
                    <span class="text-strong block text-[15px] font-semibold">{step.value}</span>
                    <span class="text-muted block text-[12px]">{step.note}</span>
                  </dd>
                </div>
              {/each}
            </dl>
          </li>
        </ul>
      {/if}
    </Card>

    <Card glyph="send" title={copy.formTitle} testid="links-form-card">
      <p class="text-muted">{copy.formJob}</p>
      <form class="flex flex-col gap-4" onsubmit={onCreate} novalidate data-testid="links-form">
        {#each [{ id: 'slug', name: copy.slugLabel, hint: copy.slugHint, placeholder: 'reddit-classicwow' }, { id: 'target', name: copy.targetLabel, hint: copy.targetHint, placeholder: '/' }, { id: 'label', name: copy.labelLabel, hint: copy.labelHint, placeholder: 'Reddit: r/classicwow' }] as field (field.id)}
          <div class="flex flex-col gap-1">
            <label class="quiet-label" for={`link-${field.id}`}>{field.name}</label>
            {#if field.id === 'slug'}
              <input
                id="link-slug"
                class="{FIELD} font-mono"
                autocapitalize="none"
                autocomplete="off"
                spellcheck="false"
                placeholder={field.placeholder}
                aria-invalid={fieldErrors.slug !== undefined}
                aria-describedby="link-slug-hint"
                bind:value={slug}
                data-testid="links-slug"
              />
            {:else if field.id === 'target'}
              <input
                id="link-target"
                class={FIELD}
                autocapitalize="none"
                autocomplete="off"
                spellcheck="false"
                placeholder={field.placeholder}
                aria-invalid={fieldErrors.target !== undefined}
                aria-describedby="link-target-hint"
                bind:value={target}
                data-testid="links-target"
              />
            {:else}
              <input
                id="link-label"
                class={FIELD}
                autocomplete="off"
                placeholder={field.placeholder}
                aria-invalid={fieldErrors.label !== undefined}
                aria-describedby="link-label-hint"
                bind:value={label}
                data-testid="links-label"
              />
            {/if}
            <p
              id={`link-${field.id}-hint`}
              class="text-[13px] {fieldErrors[field.id] === undefined ? 'text-muted' : 'text-wipe'}"
              data-testid={`links-${field.id}-hint`}
            >
              {fieldErrors[field.id] ?? field.hint}
            </p>
          </div>
        {/each}
        <button
          type="submit"
          class="{button('primary', 'md')} self-start"
          disabled={saving}
          aria-busy={saving}
          data-testid="links-create"
        >
          {saving ? copy.creating : copy.create}
        </button>
        <div class="min-h-[22px]" aria-live="polite">
          {#if notice !== ''}<p class="text-[14px]" data-testid="links-created">{notice}</p>{/if}
          {#if formError !== ''}<p class="text-[14px]" role="alert" data-testid="links-form-error">
              {formError}
            </p>{/if}
        </div>
      </form>
    </Card>
  </div>
{/if}
