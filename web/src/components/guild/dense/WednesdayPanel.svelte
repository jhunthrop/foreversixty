<!-- web/src/components/guild/dense/WednesdayPanel.svelte -->
<!-- The officer's next raid (dense spec 5.1), the Overview's framed panel: the raid's Row, the answers
     in figures with Ready of those coming, who is Late, on the Bench and Out, then each raid group
     against the comp target: In, Late and the seats open, with the raiders who could fill them as
     one-click chips (a not-answered raider: Ping, the raid-wide Discord ping that names every
     non-signer; a benched one: Move to In), and how many seats stay open if everyone says yes. -->
<script lang="ts">
  import {
    GuildApiError,
    type GuildRaidSignup,
    type GuildRaidView,
    type NextRaidView,
  } from '../../../lib/guild/api';
  import { pingRaid } from '../../../lib/guild/overview-api';
  import { pingAllowed } from '../../../lib/guild/overview-role-view';
  import { trySignup } from '../../../lib/guild/signup-save';
  import { compGaps, FILLERS_SHOWN, type CompGapRow, type Filler } from '../../../lib/guild/dense/week-board';
  import { GROUP_WORD } from '../../../lib/guild/dense/roster-board';
  import { clock, viewerLocale } from '../../../lib/dates';
  import { withRetry } from '../../../lib/ui/copy';
  import Button from '../../ui/Button.svelte';
  import CharacterName from '../../ui/CharacterName.svelte';
  import Figure from '../../ui/Figure.svelte';
  import Section from '../../ui/Section.svelte';
  import StatRow from '../../ui/StatRow.svelte';
  import StatTable from '../../ui/StatTable.svelte';
  import NextRaidRow from '../overview/NextRaidRow.svelte';

  let {
    guildId,
    next,
    raid,
    ready,
    planHref,
    viewerTz,
    now,
    onChanged,
  }: {
    guildId: number;
    next: NextRaidView;
    raid: GuildRaidView;
    /** Ready and coming (In and Late) from the readiness board for this raid; null until it answers. */
    ready: { ready: number; coming: number } | null;
    planHref: string;
    viewerTz: string;
    now: Date;
    /** A chip changed the raid: read it again. */
    onChanged: () => void;
  } = $props();

  const columns = [
    { id: 'gap', label: 'Gap' },
    { id: 'need', label: 'Need' },
    { id: 'in', label: 'In' },
    { id: 'late', label: 'Late' },
  ] as const;

  let pingedAt = $state<Date | null>(null);
  let pingError = $state('');
  let moving = $state<string | null>(null);
  /** A Move to In the API refused: the line the officer sees beside the ping line. */
  let moveError = $state('');
  const rows = $derived(compGaps(raid));
  const stillShort = $derived(rows.filter((r) => r.stillShort > 0));
  const notAnswered = $derived(raid.not_signed_raiders?.length ?? raid.not_signed.length);
  const canPing = $derived(pingedAt === null && pingAllowed(raid, now));
  const pingedText = $derived(
    pingedAt === null ? '' : `Pinged ${clock(pingedAt, viewerTz, { locale: viewerLocale() })}`,
  );

  function named(state: GuildRaidSignup['state']): GuildRaidSignup[] {
    return raid.signups.filter((s) => s.state === state);
  }

  async function ping(): Promise<void> {
    if (!canPing) return;
    try {
      await pingRaid(guildId, raid.id);
      pingedAt = now;
      pingError = '';
    } catch (error) {
      pingError = error instanceof GuildApiError ? error.message : withRetry('That did not go through.');
    }
  }

  async function moveIn(filler: Filler): Promise<void> {
    const row = raid.signups.find((s) => s.character_key === filler.key);
    if (row === undefined || moving !== null) return;
    moving = filler.key;
    moveError = '';
    const saved = await trySignup(guildId, raid.id, row.character_key, {
      state: 'in',
      role: row.role,
      note: row.note,
    });
    moving = null;
    if (saved.ok) onChanged();
    else moveError = saved.line;
  }

  function figures(row: CompGapRow) {
    return [
      row.gap > 0
        ? { value: `${row.gap} short`, tone: 'wipe' as const }
        : { value: '✓', tone: 'kill' as const },
      { value: String(row.need), muted: true },
      { value: String(row.in), tone: row.in >= row.need ? ('kill' as const) : ('gold' as const) },
      row.late > 0 ? { value: String(row.late) } : { value: '–', muted: true },
    ];
  }
</script>

{#snippet people(list: readonly GuildRaidSignup[])}
  {#each list as s, i (s.character_key)}<CharacterName name={s.name} cls={s.class} />{i < list.length - 1
      ? ', '
      : ''}{/each}
{/snippet}

<Section
  title="Wednesday"
  panel
  flush
  action={{ label: 'Edit target', href: planHref }}
  testid="guild-week-wednesday"
>
  <div class="px-5"><NextRaidRow {next} durationMinutes={raid.duration_minutes} {viewerTz} {now} /></div>
  <div class="flex flex-wrap gap-x-8 gap-y-3 px-5 pt-3 pb-2" data-testid="guild-week-figures">
    <Figure value={String(raid.counts.in)} label="In" tone="kill" />
    <Figure value={String(raid.counts.late)} label="Late" tone="gold" />
    <Figure value={String(raid.counts.bench)} label="Bench" />
    <Figure value={String(raid.counts.out)} label="Out" />
    <Figure value={String(notAnswered)} label="Not answered" tone="wipe" />
    {#if ready !== null}<Figure value={`${ready.ready}/${ready.coming}`} label="Ready" />{/if}
  </div>
  <p class="text-small text-muted px-5 pb-3" data-testid="guild-week-names">
    {#if named('late').length > 0}Late {@render people(named('late'))}{/if}
    {#if named('bench').length > 0}
      · Bench {@render people(named('bench'))}{/if}
    {#if named('out').length > 0}
      · Out {@render people(named('out'))}{/if}
  </p>
  <div class="no-gutter">
    <StatTable
      columns={[...columns]}
      nameHead="Role"
      bar={false}
      dense
      label="Seats against the target"
      testid="guild-week-comp"
    >
      {#each rows as row (row.group)}
        <StatRow bar={false} figures={figures(row)} testid={`guild-week-comp-${row.group}`}>
          {#snippet name()}<b>{GROUP_WORD[row.group]}</b>{/snippet}
          {#snippet sub()}
            {#if row.fillers.length > 0}
              <span class="fills">
                {#each row.fillers.slice(0, FILLERS_SHOWN) as f (f.key)}
                  {#if f.kind === 'ping' && pingedAt !== null}
                    <span class="fill done"><CharacterName name={f.name} cls={f.cls} />{pingedText}</span>
                  {:else}
                    <button
                      type="button"
                      class="fill"
                      disabled={f.kind === 'ping' ? !canPing : moving !== null}
                      onclick={() => void (f.kind === 'ping' ? ping() : moveIn(f))}
                      data-testid={`guild-fill-${f.name.toLowerCase()}`}
                      ><CharacterName name={f.name} cls={f.cls} /><span class="act"
                        >{f.kind === 'ping' ? 'Ping' : 'Move to In'}</span
                      ></button
                    >
                  {/if}
                {/each}
                {#if row.fillers.length > FILLERS_SHOWN}<span>+{row.fillers.length - FILLERS_SHOWN} more</span
                  >{/if}
              </span>
            {/if}
          {/snippet}
        </StatRow>
      {/each}
    </StatTable>
  </div>
  {#if stillShort.length > 0}
    <p class="text-muted px-5 pt-2 text-[14px]" data-testid="guild-week-still-short">
      With every filler in: <b class="text-wipe"
        >{stillShort.map((r) => `${GROUP_WORD[r.group]} ${r.stillShort} short`).join(', ')}</b
      >
    </p>
  {/if}
  <div class="flex flex-wrap items-center gap-x-4 gap-y-2 px-5 pt-3 pb-4">
    <Button
      tone="primary"
      size="sm"
      disabled={!canPing || notAnswered === 0}
      onclick={() => void ping()}
      testid="guild-ping-all">{pingedAt === null ? `Ping ${notAnswered} not answered` : pingedText}</Button
    >
    <a class="text-gold text-[14px] font-semibold" href={planHref}>Open the planner ›</a>
    {#if pingError !== ''}<span class="text-small text-wipe">{pingError}</span>{/if}
    {#if moveError !== ''}<span class="text-small text-wipe" role="alert" data-testid="guild-move-error"
        >{moveError}</span
      >{/if}
  </div>
</Section>

<style>
  .no-gutter :global(.ui-st) {
    --st-icon: 0px;
  }
  .fills {
    display: flex;
    flex-wrap: wrap;
    gap: 6px;
    margin-top: 4px;
  }
  .fill {
    display: inline-flex;
    align-items: center;
    gap: 6px;
    min-height: 28px;
    padding: 0 6px 0 0;
    font: 600 12px/1 var(--font-body);
    color: var(--color-text);
  }
  button.fill {
    cursor: pointer;
  }
  button.fill:hover:not(:disabled) .act {
    text-decoration: underline;
  }
  button.fill:disabled {
    opacity: 0.6;
    cursor: default;
  }
  .fill .act {
    color: var(--color-gold);
  }
  .fill.done {
    color: var(--color-muted);
  }
</style>
