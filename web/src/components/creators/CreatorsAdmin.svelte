<!-- web/src/components/creators/CreatorsAdmin.svelte -->
<!-- The owner's /admin/creators (design/specs/2026-11-08-creator-program.md 3a, 3g): a split Page of
     panels, the grammar of /admin/links. Main: the Requests panel (open Rows, only while any wait)
     above Your creators (a Row per creator, the status Tag and Edit in the trail, five Figures in the
     foot, Owed so far in gold). Rail: one form panel in three modes, Add a creator, Edit {name} (link
     and code fixed; invite and Stripe refresh) and Approve {channel} (prefilled from the request).
     Anyone but the owner sees one line: the API answers 403 for them whatever this page shows. -->
<script lang="ts">
  import { AccountError, fetchMeOnce } from '../../lib/account/api';
  import {
    addCreator,
    approveRequest,
    declineRequest,
    editCreator,
    listCreators,
    listRequests,
    makeInvite,
    refreshStripe,
  } from '../../lib/creators/api';
  import { adminCreatorsCopy as copy, requestsCopy, stripeStatusCopy } from '../../lib/creators/copy';
  import {
    emptyForm,
    formFromCreator,
    formFromRequest,
    inputFromForm,
    type CreatorFormValues,
  } from '../../lib/creators/form';
  import {
    STATUS_TONE,
    channelTitle,
    creatorFigures,
    dayLabel,
    requestSub,
    type Creator,
    type CreatorRequest,
  } from '../../lib/creators/model';
  import { createCopyFeedback, writeClipboard } from '../../lib/ui/copy-feedback';
  import { FIELD_LABEL } from '../../lib/ui/field';
  import { LINK } from '../../lib/ui/primitives';
  import SignInPrompt from '../SignInPrompt.svelte';
  import Button from '../ui/Button.svelte';
  import EmptyState from '../ui/EmptyState.svelte';
  import Field from '../ui/Field.svelte';
  import Figure from '../ui/Figure.svelte';
  import LoadError from '../ui/LoadError.svelte';
  import Page from '../ui/Page.svelte';
  import Row from '../ui/Row.svelte';
  import Section from '../ui/Section.svelte';
  import Skeleton from '../ui/Skeleton.svelte';
  import Tag from '../ui/Tag.svelte';

  type Phase = 'loading' | 'signed-out' | 'not-owner' | 'failed' | 'ready';
  type Mode =
    { kind: 'add' } | { kind: 'edit'; creator: Creator } | { kind: 'approve'; request: CreatorRequest };

  let { now = () => new Date() }: { now?: () => Date } = $props();

  let phase = $state<Phase>('loading');
  let creators = $state<Creator[]>([]);
  let requests = $state<CreatorRequest[]>([]);
  let stripeEnabled = $state(true);
  let connectDenied = $state(false);
  let mode = $state<Mode>({ kind: 'add' });
  let values = $state<CreatorFormValues>(emptyForm(new Date()));
  let saving = $state(false);
  let refreshing = $state(false);
  let inviting = $state(false);
  let fieldErrors = $state<Readonly<Record<string, string>>>({});
  let notice = $state('');
  let formError = $state('');

  // The invite's Copied clears after about two seconds, as every copy control's does.
  let inviteCopied = $state(false);
  const copier = createCopyFeedback((state) => (inviteCopied = state === 'copied'), writeClipboard);
  $effect(() => copier.dispose);

  function phaseForError(error: unknown): Phase {
    if (error instanceof AccountError && error.status === 401) return 'signed-out';
    if (error instanceof AccountError && error.status === 403) return 'not-owner';
    return 'failed';
  }

  async function reload(): Promise<void> {
    const [list, waiting] = await Promise.all([listCreators(), listRequests()]);
    creators = list.creators;
    stripeEnabled = list.stripe_enabled;
    connectDenied = list.connect_denied;
    requests = waiting;
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
      await reload();
      phase = 'ready';
    } catch (error) {
      phase = phaseForError(error);
    }
  }

  $effect(() => {
    void load();
  });

  function setMode(next: Mode): void {
    mode = next;
    fieldErrors = {};
    formError = '';
    notice = '';
    if (next.kind === 'edit') values = formFromCreator(next.creator);
    else if (next.kind === 'approve') values = formFromRequest(next.request, now());
    else values = emptyForm(now());
  }

  function failed(error: unknown, fallback: string): void {
    if (error instanceof AccountError && error.fields !== null) fieldErrors = error.fields;
    else formError = error instanceof AccountError && error.message !== '' ? error.message : fallback;
  }

  async function onSubmit(event: SubmitEvent): Promise<void> {
    event.preventDefault();
    saving = true;
    fieldErrors = {};
    formError = '';
    notice = '';
    const input = inputFromForm(values);
    try {
      if (mode.kind === 'edit') {
        const { slug: _slug, code: _code, ...editable } = input;
        await editCreator(mode.creator.id, editable);
        notice = copy.saved;
      } else {
        const made =
          mode.kind === 'approve' ? await approveRequest(mode.request.id, input) : await addCreator(input);
        setMode({ kind: 'add' });
        notice = copy.added(made.name);
      }
      await reload();
    } catch (error) {
      failed(error, copy.saveFailed);
    } finally {
      saving = false;
    }
  }

  async function onInvite(id: number): Promise<void> {
    notice = '';
    formError = '';
    inviting = true;
    try {
      const url = await makeInvite(id);
      await copier.copy(url);
      notice = copy.inviteCopied;
    } catch {
      formError = copy.inviteFailed;
    } finally {
      inviting = false;
    }
  }

  async function onRefresh(id: number): Promise<void> {
    notice = '';
    formError = '';
    refreshing = true;
    try {
      await refreshStripe(id);
      await reload();
      notice = copy.refreshed;
    } catch {
      formError = copy.refreshFailed;
    } finally {
      refreshing = false;
    }
  }

  async function onDecline(id: number): Promise<void> {
    try {
      await declineRequest(id);
      if (mode.kind === 'approve' && mode.request.id === id) setMode({ kind: 'add' });
      requests = requests.filter((r) => r.id !== id);
    } catch {
      formError = copy.saveFailed;
    }
  }

  const railTitle = $derived(
    mode.kind === 'edit'
      ? copy.editTitle(mode.creator.name)
      : mode.kind === 'approve'
        ? requestsCopy.approveTitle(channelTitle(mode.request.channel_url))
        : copy.addTitle,
  );
  const fixed = $derived(mode.kind === 'edit');
  const FIELDS = [
    { id: 'name', key: 'name', f: copy.fields.name },
    { id: 'channel_url', key: 'channel', f: copy.fields.channel },
    { id: 'slug', key: 'slug', f: copy.fields.slug },
    { id: 'code', key: 'code', f: copy.fields.code },
    { id: 'starts_on', key: 'startsOn', f: copy.fields.startsOn },
    { id: 'cap', key: 'cap', f: copy.fields.cap },
  ] as const;
</script>

{#snippet figures(c: Creator)}
  <span class="grid min-w-0 grid-cols-2 items-start gap-x-5 gap-y-2 md:flex md:flex-wrap">
    {#each creatorFigures(c) as fig (fig.key)}
      <span class={fig.lead ? 'md:w-[220px]' : 'md:w-[104px]'}
        ><Figure
          value={fig.value}
          label={fig.label}
          caption={fig.note}
          size="sm"
          tone={fig.lead ? 'gold' : 'strong'}
        /></span
      >
    {/each}
  </span>
{/snippet}

{#if phase === 'loading'}
  <div data-testid="creators-loading"><Skeleton lines={6} rowHeight="h-6" minHeight="min-h-[320px]" /></div>
{:else if phase === 'signed-out'}
  <div data-testid="creators-signed-out">
    <SignInPrompt line={copy.signedOut} next="/admin/creators" testid="creators-signin" />
  </div>
{:else if phase === 'not-owner'}
  <p class="text-[15px]" data-testid="creators-not-owner">{copy.notOwner}</p>
{:else if phase === 'failed'}
  <div data-testid="creators-error">
    <LoadError message={copy.loadFailed} onRetry={() => void load()} testid="creators-load-error" />
  </div>
{:else}
  <Page layout="split" panels testid="creators-page-body">
    {#if requests.length > 0}
      <Section
        title={requestsCopy.title}
        meta={requestsCopy.meta(requests.length)}
        panel
        flush
        testid="creators-requests"
      >
        <div role="group" aria-label={requestsCopy.title}>
          {#each requests as r (r.id)}
            <Row title={channelTitle(r.channel_url)} sub={requestSub(r)} wrapTrail testid={`request-${r.id}`}>
              {#snippet trail()}
                <span class="flex shrink-0 items-center gap-2.5">
                  <Button
                    tone="primary"
                    size="sm"
                    onclick={() => setMode({ kind: 'approve', request: r })}
                    testid={`request-approve-${r.id}`}>{requestsCopy.approve}</Button
                  >
                  <Button
                    tone="ghost"
                    size="sm"
                    onclick={() => void onDecline(r.id)}
                    testid={`request-decline-${r.id}`}>{requestsCopy.decline}</Button
                  >
                </span>
              {/snippet}
              {#snippet foot()}
                <span class="flex flex-col gap-0.5 text-[13px] leading-5">
                  <span>{r.about}</span>
                  <span class="text-muted">
                    <a class={LINK} href={r.channel_url} rel="noopener" target="_blank">{r.channel_url}</a>
                    {#if r.contact_email}<span data-testid={`request-email-${r.id}`}>
                        · {r.contact_email}</span
                      >{/if}
                  </span>
                </span>
              {/snippet}
            </Row>
          {/each}
        </div>
      </Section>
    {/if}
    <Section
      title={copy.tableTitle}
      meta={copy.tableMeta(creators.length)}
      panel
      flush
      testid="creators-table"
    >
      <div>
        {#if connectDenied || !stripeEnabled}
          <p class="text-wipe px-4 pt-3 text-[14px] leading-5" role="alert">
            {#if connectDenied}<span data-testid="creators-connect-denied">{copy.connectDenied}</span>{/if}
            {#if !stripeEnabled}<span data-testid="creators-stripe-off">{copy.stripeOff}</span>{/if}
          </p>
        {/if}
        {#if creators.length === 0}
          <EmptyState message={copy.empty} inset testid="creators-empty" />
        {:else}
          <div role="group" aria-label={copy.tableTitle}>
            {#each creators as c (c.id)}
              <Row
                title={c.name}
                sub={copy.sub(c.slug, c.code, dayLabel(c.starts_on))}
                selected={mode.kind === 'edit' && mode.creator.id === c.id}
                wrapTrail
                testid={`creator-row-${c.slug}`}
              >
                {#snippet trail()}
                  <span class="flex shrink-0 items-center gap-2.5">
                    <Tag tone={STATUS_TONE[c.stripe_status]} testid={`creator-status-${c.slug}`}
                      >{stripeStatusCopy[c.stripe_status]}</Tag
                    >
                    <Button
                      size="sm"
                      label={copy.editLabel(c.name)}
                      onclick={() => setMode({ kind: 'edit', creator: c })}
                      testid={`creator-edit-${c.slug}`}>{copy.edit}</Button
                    >
                  </span>
                {/snippet}
                {#snippet foot()}{@render figures(c)}{/snippet}
              </Row>
            {/each}
          </div>
        {/if}
        <p class="px-4 py-3 text-[14px]">
          <a class={LINK} href="/admin/creators/payouts" data-testid="creators-statements-link"
            >{copy.statementsLink} ›</a
          >
        </p>
      </div>
    </Section>
    {#snippet rail()}
      <Section title={railTitle} panel testid="creators-form-card">
        <div class="flex flex-col gap-3">
          <p class="text-muted text-[14px]">
            {mode.kind === 'approve' ? requestsCopy.approveJob : copy.formJob}
          </p>
          <form class="flex flex-col gap-3" onsubmit={onSubmit} novalidate data-testid="creators-form">
            {#each FIELDS as field (field.id)}
              {#if fixed && (field.id === 'slug' || field.id === 'code')}
                {#if field.id === 'slug'}
                  <div class="flex flex-col gap-1.5">
                    <span class={FIELD_LABEL}>{copy.fields.slug.label} · {copy.fields.code.label}</span>
                    <span class="text-strong text-[15px]" data-testid="creators-fixed"
                      >{copy.fixed(values.slug, values.code)}</span
                    >
                  </div>
                {/if}
              {:else}
                <Field
                  id={`creator-${field.id}`}
                  label={field.f.label}
                  hint={field.f.hint}
                  placeholder={field.f.placeholder}
                  type={field.id === 'starts_on' ? 'date' : field.id === 'channel_url' ? 'url' : 'text'}
                  inputmode={field.id === 'cap' ? 'numeric' : undefined}
                  error={fieldErrors[field.id === 'cap' ? 'cap_cents' : field.id]}
                  autocomplete="off"
                  bind:value={values[field.id]}
                  testid={`creators-${field.id}`}
                />
              {/if}
            {/each}
            <div class="flex flex-wrap items-center gap-3">
              <Button
                type="submit"
                tone="primary"
                busy={saving}
                busyLabel={copy.saving}
                testid="creators-save">{mode.kind === 'edit' ? copy.save : copy.add}</Button
              >
              {#if mode.kind !== 'add'}
                <Button onclick={() => setMode({ kind: 'add' })} testid="creators-cancel"
                  >{copy.cancelEdit}</Button
                >
              {/if}
            </div>
            {#if mode.kind === 'edit'}
              {@const editing = mode.creator}
              <div class="flex flex-wrap items-center gap-2.5">
                <Button
                  size="sm"
                  busy={inviting}
                  busyLabel={copy.inviting}
                  onclick={() => void onInvite(editing.id)}
                  testid="creators-invite">{inviteCopied ? copy.inviteCopiedShort : copy.invite}</Button
                >
                <Button
                  size="sm"
                  busy={refreshing}
                  busyLabel={copy.refreshing}
                  onclick={() => void onRefresh(editing.id)}
                  testid="creators-refresh">{copy.refresh}</Button
                >
              </div>
            {/if}
            <div class="min-h-[22px]" aria-live="polite">
              {#if notice !== ''}<p class="text-[14px]" data-testid="creators-notice">{notice}</p>{/if}
              {#if formError !== ''}<p
                  class="text-wipe text-[14px]"
                  role="alert"
                  data-testid="creators-form-error"
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
