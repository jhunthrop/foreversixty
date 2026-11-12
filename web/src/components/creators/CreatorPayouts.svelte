<!-- web/src/components/creators/CreatorPayouts.svelte -->
<!-- The owner's /admin/creators/payouts (design/specs/2026-11-08-creator-program.md 3b): a split Page
     of panels. Main: the `{n} creators` panel with the month Select in its head, the payout switch's
     notice when it is off, and a Row per creator (the status line, a `Stripe not ready` Tag when
     blocked, Approve when the month is ready and owes money, Lines; five Figures in the foot, Total
     first in gold). Rail: the selected creator's lines. Approve asks once in the shared ConfirmDialog
     with the Stripe balance it is checked against; nothing is ever paid without it. -->
<script lang="ts">
  import { tick } from 'svelte';
  import { AccountError, fetchMeOnce } from '../../lib/account/api';
  import { approveStatement, listStatements, payoutBalance } from '../../lib/creators/api';
  import { adminCreatorsCopy, payoutsCopy as copy } from '../../lib/creators/copy';
  import {
    canApprove,
    lineText,
    money,
    monthLabel,
    recentMonths,
    statementFigures,
    statementStatusLine,
    toPay,
    type Statement,
  } from '../../lib/creators/model';
  import SignInPrompt from '../SignInPrompt.svelte';
  import Button from '../ui/Button.svelte';
  import ConfirmDialog from '../ui/ConfirmDialog.svelte';
  import EmptyState from '../ui/EmptyState.svelte';
  import Figure from '../ui/Figure.svelte';
  import LoadError from '../ui/LoadError.svelte';
  import Page from '../ui/Page.svelte';
  import Row from '../ui/Row.svelte';
  import Section from '../ui/Section.svelte';
  import Select from '../ui/Select.svelte';
  import Skeleton from '../ui/Skeleton.svelte';
  import Tag from '../ui/Tag.svelte';

  const LINES_HEADING = 'payout-lines-title';

  type Phase = 'loading' | 'signed-out' | 'not-owner' | 'failed' | 'ready';

  let { now = () => new Date() }: { now?: () => Date } = $props();

  const months = recentMonths(now());
  const options = months.map((m, i) => ({
    value: m,
    label: i === 0 ? copy.openMonth(monthLabel(m)) : monthLabel(m),
  }));

  let phase = $state<Phase>('loading');
  /** The month on show: the API's answer, so the page and the API never disagree on which. */
  let month = $state('');
  let statements = $state<Statement[]>([]);
  let payoutsEnabled = $state(true);
  let connectDenied = $state(false);
  let selected = $state<number | null>(null);
  let confirming = $state<Statement | null>(null);
  let paying = $state(false);
  let payError = $state('');
  /** The available balance for the open dialog, in cents; null while it loads; NaN when it failed. */
  let available = $state<number | null>(null);
  let dialogProblem = $state('');

  const current = $derived(statements.find((s) => s.creator_id === selected) ?? statements[0] ?? null);
  const canPay = $derived(payoutsEnabled && !connectDenied);

  function phaseForError(error: unknown): Phase {
    if (error instanceof AccountError && error.status === 401) return 'signed-out';
    if (error instanceof AccountError && error.status === 403) return 'not-owner';
    return 'failed';
  }

  async function read(m: string): Promise<void> {
    const page = await listStatements(m);
    month = page.month;
    statements = page.statements;
    payoutsEnabled = page.payouts_enabled;
    connectDenied = page.connect_denied;
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
      await read(month);
      phase = 'ready';
    } catch (error) {
      phase = phaseForError(error);
    }
  }

  $effect(() => {
    void load();
  });

  async function onMonth(event: Event): Promise<void> {
    const next = (event.target as HTMLSelectElement | null)?.value;
    if (next === undefined || next === month) return;
    month = next;
    payError = '';
    try {
      await read(next);
    } catch (error) {
      phase = phaseForError(error);
    }
  }

  async function showLines(id: number): Promise<void> {
    selected = id;
    await tick();
    const heading = document.getElementById(LINES_HEADING);
    heading?.scrollIntoView({ block: 'nearest' });
    heading?.focus();
  }

  /** Opens the dialog for one statement and reads the balance it is checked against. */
  async function confirm(s: Statement): Promise<void> {
    confirming = s;
    available = null;
    dialogProblem = '';
    try {
      available = await payoutBalance();
    } catch {
      available = Number.NaN;
    }
  }

  const short = $derived(confirming !== null && available !== null && available < confirming.total_cents);
  const balanceLine = $derived(
    confirming === null || available === null
      ? copy.balanceChecking
      : Number.isNaN(available)
        ? copy.balanceFailed
        : copy.balanceLine(money(available), money(confirming.total_cents)),
  );
  const blocked = $derived(available === null || Number.isNaN(available) || short);
  const problem = $derived(dialogProblem !== '' ? dialogProblem : short ? copy.balanceShort : '');

  async function pay(): Promise<void> {
    if (confirming === null) return;
    paying = true;
    dialogProblem = '';
    const target = confirming;
    try {
      const paid = await approveStatement(target.creator_id, target.month);
      statements = statements.map((s) => (s.creator_id === paid.creator_id ? { ...paid, name: s.name } : s));
      confirming = null;
    } catch (error) {
      if (error instanceof AccountError && error.code === 'balance_insufficient') {
        // Nothing was paid or recorded: the dialog stays open and says so.
        dialogProblem = copy.transferShort;
      } else {
        confirming = null;
        payError = copy.approveFailed(
          error instanceof AccountError && error.message !== '' ? error.message : '',
        );
      }
    } finally {
      paying = false;
    }
  }
</script>

{#if phase === 'loading'}
  <div data-testid="payouts-loading"><Skeleton lines={6} rowHeight="h-6" minHeight="min-h-[320px]" /></div>
{:else if phase === 'signed-out'}
  <div data-testid="payouts-signed-out">
    <SignInPrompt line={adminCreatorsCopy.signedOut} next="/admin/creators/payouts" testid="payouts-signin" />
  </div>
{:else if phase === 'not-owner'}
  <p class="text-[15px]" data-testid="payouts-not-owner">{adminCreatorsCopy.notOwner}</p>
{:else if phase === 'failed'}
  <div data-testid="payouts-error">
    <LoadError message={copy.loadFailed} onRetry={() => void load()} testid="payouts-load-error" />
  </div>
{:else}
  <Page layout="split" panels testid="payouts-page-body">
    <Section
      title={copy.tableTitle(statements.length)}
      meta={copy.tableMeta(money(toPay(statements)))}
      panel
      flush
      testid="payouts-table"
    >
      {#snippet head()}
        <Select
          id="payouts-month"
          size="dense"
          fit
          label={copy.monthLabel}
          value={month}
          {options}
          onchange={onMonth}
          testid="payouts-month"
        />
      {/snippet}
      <div>
        {#if connectDenied || !payoutsEnabled}
          <p class="text-wipe px-4 pt-3 text-[14px] leading-5" role="status">
            {#if connectDenied}<span data-testid="payouts-connect-denied"
                >{adminCreatorsCopy.connectDenied}</span
              >{/if}
            {#if !payoutsEnabled}<span data-testid="payouts-off">{copy.payoutsOff}</span>{/if}
          </p>
        {/if}
        {#if statements.length === 0}
          <EmptyState message={copy.empty} inset testid="payouts-empty" />
        {:else}
          <div role="group" aria-label={monthLabel(month)}>
            {#each statements as s (s.creator_id)}
              <Row
                title={s.name}
                sub={statementStatusLine(s)}
                selected={current?.creator_id === s.creator_id}
                wrapTrail
                testid={`payout-row-${s.creator_id}`}
              >
                {#snippet trail()}
                  <span class="flex shrink-0 items-center gap-2.5">
                    {#if s.status === 'blocked'}<Tag tone="warn" testid={`payout-blocked-${s.creator_id}`}
                        >{copy.notReady}</Tag
                      >{/if}
                    {#if canApprove(s)}
                      <Button
                        tone={canPay ? 'primary' : 'secondary'}
                        size="sm"
                        disabled={!canPay}
                        onclick={() => void confirm(s)}
                        testid={`payout-approve-${s.creator_id}`}>{copy.approve(money(s.total_cents))}</Button
                      >
                    {/if}
                    <Button
                      size="sm"
                      label={copy.linesLabel(s.name)}
                      onclick={() => void showLines(s.creator_id)}
                      testid={`payout-lines-${s.creator_id}`}>{copy.lines}</Button
                    >
                  </span>
                {/snippet}
                {#snippet foot()}
                  <span class="grid min-w-0 grid-cols-2 items-start gap-x-5 gap-y-2 md:flex md:flex-wrap">
                    {#each statementFigures(s) as fig (fig.key)}
                      <span class={fig.lead ? 'md:w-[220px]' : 'md:w-[104px]'}
                        ><Figure
                          value={fig.value}
                          label={fig.label}
                          caption={fig.note || undefined}
                          size="sm"
                          tone={fig.lead ? 'gold' : 'strong'}
                        /></span
                      >
                    {/each}
                  </span>
                {/snippet}
              </Row>
            {/each}
          </div>
        {/if}
        {#if payError !== ''}<p
            class="text-wipe px-4 py-3 text-[14px]"
            role="alert"
            data-testid="payouts-error-line"
          >
            {payError}
          </p>{/if}
      </div>
    </Section>
    {#snippet rail()}
      {#if current !== null}
        <Section
          title={copy.linesTitle(current.name)}
          focusId={LINES_HEADING}
          panel
          flush
          testid="payout-lines"
        >
          {#if current.lines.length === 0}
            <EmptyState message={copy.linesEmpty} inset testid="payout-lines-empty" />
          {:else}
            <div role="group" aria-label={copy.linesTitle(current.name)}>
              {#each current.lines as line, i (i)}
                {@const shown = lineText(line)}
                <Row title={shown.text} density="slim">
                  {#snippet trail()}<span class="tabular-nums {shown.refund ? 'text-wipe' : ''}"
                      >{shown.amount}</span
                    >{/snippet}
                </Row>
              {/each}
            </div>
          {/if}
        </Section>
      {/if}
    {/snippet}
  </Page>
  {#if confirming !== null}
    {@const amount = money(confirming.total_cents)}
    <ConfirmDialog
      title={copy.confirmTitle(confirming.name, amount, monthLabel(confirming.month))}
      line={copy.confirmBody(amount)}
      confirmLabel={copy.confirmPay(amount)}
      busy={paying}
      busyLabel={copy.approving}
      disabled={blocked}
      onconfirm={() => void pay()}
      onclose={() => (confirming = null)}
      testid="pay"
    >
      <p class="text-muted m-0 mt-2 text-[14px] leading-5" data-testid="pay-balance">{balanceLine}</p>
      {#if problem !== ''}<p
          class="text-wipe m-0 mt-1 text-[14px] leading-5"
          role="alert"
          data-testid="pay-problem"
        >
          {problem}
        </p>{/if}
    </ConfirmDialog>
  {/if}
{/if}
