<!-- web/src/components/report/ActorRow.svelte -->
<!-- One table row, exactly the shape spec section 4 lists: parse percentile, class-coloured
     name, amount bar segmented by ability, active time, per-second figure, expander. On
     phone the row becomes a card: the same fields, stacked, with the bar full width.

     Active time: `actor.active_ms` is exact even under a brushed window (window.ts measures
     it from the one-second series rather than scaling it), but that series is one-second
     grained, so a percentage of a short window quantises hard -- a 3s window can only ever
     read 0%, 33%, 67% or 100%. Below ACTIVITY_SECONDS_BELOW_MS this column reads seconds
     instead of a percentage: "2s active" is honest at that grain, "67%" implies precision
     the data does not have.

     The Amount column is `actor.effective`, which window.ts computes independently from
     the one-second series -- exact regardless of the window -- and which filters.ts leaves
     alone unless a target or boss filter is active (ReportView.svelte's `approximate`
     already folds that in). It never carries the `~` mark. The expanded Abilities and
     Targets tables below the row DO carry it: `ability.total`/`effective` and
     `target.total` are the per-ability and per-target splits that both the window and a
     target/boss filter scale by a ratio. -->
<script lang="ts">
  import { characterHref, splitUnitName } from '../../lib/characters';
  import {
    approximateAriaLabel,
    approximateMark,
    approximateTitle,
    classColorVar,
    formatAmount,
    formatDuration,
    formatPercent,
    formatPerSecond,
    parseTitle,
    percentileToken,
    schoolName,
    schoolToken,
  } from '../../lib/report/format';
  import type { Placement } from '../../lib/report/percentile';
  import { abilityKey, type Ability, type Actor } from '../../lib/report/types';
  import CopyCsv from './CopyCsv.svelte';
  import type { ExactSplit } from '../../lib/report/exact';
  import AbilityBar from './AbilityBar.svelte';
  import ClassIcon from './ClassIcon.svelte';

  let {
    rank,
    actor,
    peak,
    durationMs,
    percentile = null,
    parseFallback = '',
    pairsLabel = 'Targets',
    share = 0,
    approximate = false,
    amountApproximate = false,
    splitUnavailable = false,
    deadSince = null,
    measure = undefined,
    characterLink = null,
  }: {
    rank: number;
    actor: Actor;
    peak: number;
    durationMs: number;
    percentile?: Placement | null;
    parseFallback?: string;
    /** The word over the per-unit split: Targets for damage done and healing, Sources for damage taken. */
    pairsLabel?: string;
    /** This row's share of its table's total, 0..100. */
    share?: number;
    approximate?: boolean;
    /** True when the amount itself is prorated: a window met by a target or boss filter. */
    amountApproximate?: boolean;
    /** Over the night under a target or boss filter: the split by ability cannot be had, so say so. */
    splitUnavailable?: boolean;
    /** When this player died before the window's end without coming back: the row is a corpse's. */
    deadSince?: number | null;
    /** Measures this row's split inside the window from the fight's own events. */
    measure?: (actor: Actor) => Promise<ExactSplit>;
    characterLink?: { region: string; ruleset: string } | null;
  } = $props();

  /** Below ten one-second buckets, a percentage rounds to steps of 10% or coarser. */
  const ACTIVITY_SECONDS_BELOW_MS = 10_000;

  let open = $state(false);
  /** The exact split once measured; the prorated one shows until then. */
  let exact = $state<ExactSplit | null>(null);
  let measuring = $state(false);
  let measureError = $state('');
  // A new window or filter invalidates a measurement taken under the old one.
  $effect(() => {
    void [actor, approximate];
    exact = null;
    measureError = '';
  });
  async function runMeasure(): Promise<void> {
    if (measure === undefined) return;
    measuring = true;
    measureError = '';
    try {
      exact = await measure(actor);
    } catch (thrown) {
      measureError = thrown instanceof Error ? thrown.message : 'The measurement did not run.';
    } finally {
      measuring = false;
    }
  }
  // The row's own measure answers for one actor and window; a new actor object (the table
  // re-measured, the window moved) is a new question.
  $effect(() => {
    void actor;
    exact = null;
    measureError = '';
  });
  // Once the table is measured the engine is warm, so an opened row measures itself.
  $effect(() => {
    if (
      open &&
      actor.measured &&
      exact === null &&
      !measuring &&
      measureError === '' &&
      measure !== undefined
    )
      void runMeasure();
  });
  /** What the detail tables show: the exact split when measured, the summary's otherwise. */
  const shownAbilities = $derived(exact?.abilities ?? actor.abilities);
  const shownTargets = $derived(exact?.targets ?? actor.targets);
  /** The row's amount: measured when its split was, or when the table was; prorated otherwise. */
  const shownEffective = $derived(
    exact === null ? actor.effective : exact.abilities.reduce((sum, ability) => sum + ability.effective, 0),
  );
  const amountMark = $derived(exact !== null || actor.measured || !amountApproximate ? '' : mark);
  // Overheal is a per-ability figure the summary prorates under any window, so it carries the mark
  // until the row or the table is measured, unlike the amount, which the series keep exact.
  const overhealMark = $derived(exact !== null || actor.measured || !approximate ? '' : mark);
  const shownOverheal = $derived(
    exact === null
      ? (actor.overheal ?? 0)
      : exact.abilities.reduce((sum, ability) => sum + (ability.overheal ?? 0), 0),
  );
  /** Hits and ticks together: a dot's ticks are its hits. */
  const landed = (ability: Ability): number => ability.hits + ability.ticks;
  /** The abilities table as lines, with the notes column as words. */
  function abilityCsv(): string[][] {
    return [
      ['Ability', 'Spell id', 'Via', 'Amount', 'Share %', 'Hits', 'Crit %', 'Avg', 'Max', 'Notes'],
      ...detailRows.map((ability) => {
        const hits = landed(ability);
        return [
          ability.name,
          String(ability.spell_id),
          ability.via ?? '',
          String(ability.effective),
          detailTotal === 0 ? '0' : ((ability.effective / detailTotal) * 100).toFixed(1),
          String(hits),
          hits === 0 ? '0' : ((ability.crits / hits) * 100).toFixed(1),
          hits === 0 ? '0' : Math.round(ability.effective / hits).toString(),
          String(ability.max),
          abilityNotes(ability).join('; '),
        ];
      }),
    ];
  }
  /** The abilities worth a line, largest first; a row that only missed still says so. */
  const detailRows = $derived(
    [...shownAbilities]
      .filter(
        (ability) =>
          ability.total > 0 ||
          landed(ability) > 0 ||
          (ability.misses !== undefined && Object.keys(ability.misses).length > 0),
      )
      .sort((a, b) => b.effective - a.effective),
  );
  const detailTotal = $derived(detailRows.reduce((sum, ability) => sum + ability.effective, 0));
  const targetsTotal = $derived(targetsByName.reduce((sum, target) => sum + target.total, 0));
  /** What a line's last cell says: overhealing for a heal, otherwise what did not land. */
  function abilityNotes(ability: Ability): string[] {
    // A heal says how much of it was over; a shield says what it absorbed; a heal with
    // nothing over says nothing, measured or not, so one cell means one thing.
    if (ability.overheal !== undefined && ability.total > 0) {
      const notes: string[] = [];
      if (ability.absorbed) notes.push(`${formatAmount(ability.absorbed)} absorbed`);
      if (ability.overheal > 0)
        notes.push(
          `${formatPercent((ability.overheal / ability.total) * 100)} over · ${formatAmount(ability.overheal)}`,
        );
      return notes;
    }
    const notes: string[] = [];
    if (ability.absorbed) notes.push(`${formatAmount(ability.absorbed)} absorbed`);
    if (ability.blocked) notes.push(`${formatAmount(ability.blocked)} blocked`);
    // Alphabetical, like the table's Mitigated line: the same numbers read in one order
    // whatever filter or measure produced them.
    for (const [type, count] of Object.entries(ability.misses ?? {}).sort(([a], [b]) => a.localeCompare(b))) {
      notes.push(`${count} ${type.toLowerCase()}`);
    }
    return notes;
  }
  const detailMark = $derived(exact === null ? mark : '');
  const detailTitle = $derived(exact === null ? title : 'Measured from the fight’s events for this window');
  const color = $derived(classColorVar(actor.class));
  const display = $derived(splitUnitName(actor.name));
  const activitySeconds = $derived(Math.round(actor.active_ms / 1000));
  const activityPct = $derived.by(() => {
    const over = actor.time_ms ?? durationMs;
    // Never past the whole: a cast that straddles the fight's end can count a hair over.
    return over === 0 ? 0 : Math.min((actor.active_ms / over) * 100, 100);
  });
  const showActivitySeconds = $derived(durationMs > 0 && durationMs < ACTIVITY_SECONDS_BELOW_MS);
  /** One source for the figure, which the desktop column and the phone card both read. */
  const activeText = $derived(showActivitySeconds ? `${activitySeconds}s` : formatPercent(activityPct));
  const mark = $derived(approximateMark(approximate));
  const title = $derived(approximateTitle(approximate));
  /** Overhealing as a share of the raw total, for healing rows; null where there is none. */
  const overhealPct = $derived.by(() => {
    if (exact !== null && exact.abilities.some((ability) => ability.overheal !== undefined)) {
      const total = exact.abilities.reduce((sum, ability) => sum + ability.total, 0);
      const over = exact.abilities.reduce((sum, ability) => sum + (ability.overheal ?? 0), 0);
      return total <= 0 ? null : (over / total) * 100;
    }
    return actor.overheal === undefined || actor.total <= 0 ? null : (actor.overheal / actor.total) * 100;
  });
  /**
   * Targets merged by name: a trash pack is six "Gluttonous Tick" GUIDs, and six rows of
   * the same name tell nobody anything the one row with a count does not.
   */
  /** Ability names two different spell ids share, shown with the id to tell them apart. */
  const sameName = $derived.by(() => {
    // eslint-disable-next-line svelte/prefer-svelte-reactivity
    const seen = new Map<string, Set<number>>();
    for (const ability of actor.abilities) {
      // Distinct spell ids, not rows: a pet's Melee and its owner's are one spell twice.
      // eslint-disable-next-line svelte/prefer-svelte-reactivity
      const ids = seen.get(ability.name) ?? new Set<number>();
      ids.add(ability.spell_id);
      seen.set(ability.name, ids);
    }
    return new Set([...seen.entries()].filter(([, ids]) => ids.size > 1).map(([name]) => name));
  });

  /** The row's amount by spell school: how much was physical, how much magic. */
  const schoolSplit = $derived.by(() => {
    // eslint-disable-next-line svelte/prefer-svelte-reactivity
    const totals = new Map<string, { name: string; token: string; total: number }>();
    for (const ability of shownAbilities) {
      const name = schoolName(ability.school) || 'Physical';
      const found = totals.get(name);
      if (found === undefined)
        totals.set(name, { name, token: schoolToken(ability.school), total: ability.effective });
      else found.total += ability.effective;
    }
    const sum = [...totals.values()].reduce((acc, part) => acc + part.total, 0);
    return [...totals.values()]
      .sort((a, b) => b.total - a.total)
      .map((part) => ({ ...part, pct: sum === 0 ? 0 : (part.total / sum) * 100 }));
  });

  const targetsByName = $derived.by(() => {
    // A plain Map: built once inside the derived and never read reactively by key.
    // eslint-disable-next-line svelte/prefer-svelte-reactivity
    const merged = new Map<string, { name: string; total: number; count: number }>();
    for (const target of shownTargets) {
      const name = splitUnitName(target.name).name;
      const found = merged.get(name);
      if (found === undefined) merged.set(name, { name, total: target.total, count: 1 });
      else {
        found.total += target.total;
        found.count += 1;
      }
    }
    return [...merged.values()].sort((a, b) => b.total - a.total);
  });
</script>

<li class="border-line-soft border-b" data-testid={`actor-${actor.guid}`}>
  <button
    type="button"
    class="grid min-h-11 w-full grid-cols-[28px_auto_minmax(0,1fr)_auto] items-center gap-x-3 gap-y-1 px-2 py-3 text-left text-[14px] md:grid-cols-[28px_40px_minmax(120px,1.4fr)_52px_minmax(0,3fr)_92px_80px_64px] md:py-2"
    aria-expanded={open}
    onclick={() => (open = !open)}
  >
    <span class="text-muted tabular font-mono text-[12px]"
      ><span class="mr-1 inline-block w-3 transition-transform" class:rotate-90={open} aria-hidden="true"
        >›</span
      >{rank}</span
    >

    <span
      class="tabular font-mono text-[12px]"
      class:text-muted={percentile === null}
      style={percentile === null ? undefined : `color: ${percentileToken(percentile.percentile)}`}
      title={percentile === null
        ? parseTitle(parseFallback)
        : parseTitle(percentile.percentile, percentile.ranked)}
      data-testid="row-percentile"
    >
      {#if percentile === null}<span class:text-muted={true}
          >{parseFallback === 'none' || parseFallback === 'night' ? '' : parseFallback}</span
        >{:else}<span class="label font-body mr-1 md:hidden">Parse</span>{Math.round(
          percentile.percentile,
        )}<span class="text-muted ml-1 md:hidden">among {percentile.ranked}</span>{/if}
    </span>

    <span
      class="flex min-w-0 flex-wrap items-center gap-x-2 font-semibold [&>*]:truncate"
      style={`color: ${color}`}
      data-testid="row-name"
    >
      <ClassIcon className={actor.class} />
      {#if characterLink}
        <a
          href={characterHref(characterLink.region, characterLink.ruleset, display.name)}
          style={`color: ${color}`}
        >
          {display.name}
        </a>
      {:else}
        {display.name}
      {/if}
      {#if deadSince !== null}
        <span
          class="text-death label basis-full text-[10px] whitespace-nowrap"
          title="Dead before this window ended and not raised: what is here ticked on a corpse"
          data-testid="row-dead">dead since {formatDuration(deadSince)}</span
        >
      {/if}
    </span>

    <span class="text-muted tabular hidden text-right font-mono text-[12px] md:inline" data-testid="row-share"
      >{share.toFixed(1)}%</span
    >
    <span class="col-span-4 row-start-2 md:col-span-1 md:row-auto">
      <AbilityBar abilities={actor.abilities} total={actor.effective} {peak} {color} />
    </span>

    <span class="tabular flex flex-col text-right font-mono leading-tight" data-testid="row-amount">
      <span title={exact === null && !actor.measured ? undefined : 'Measured from the fight’s events'}
        >{amountMark}{formatAmount(shownEffective)}</span
      >
      {#if overhealPct !== null}
        <span
          class="text-muted text-[11px]"
          title={overhealMark === ''
            ? 'Healing that landed on a full health bar, and how much'
            : 'Overhealing, scaled to the window in proportion to the total until it is measured'}
          data-testid="row-overheal"
          >{overhealMark}{formatPercent(overhealPct)} over · {overhealMark}{formatAmount(shownOverheal)}</span
        >
      {/if}
    </span>

    <!-- The card's third line, below the bar: Active on the left, Per sec on the right.
         It sits here in source order rather than after the two cells it replaces because a
         grid places its row-3 items in source order, and it is display:none above `md`, so
         where it sits costs the desktop grid nothing. -->
    <span class="text-muted label col-span-2 row-start-3 md:hidden" data-testid="phone-labels">
      <span class="tabular font-mono">{activeText}</span> active
    </span>
    <!-- One element, both layouts: a column under ActorTable's "Per sec" heading above
         `md`, and the same figure saying what it is once that heading is gone. -->
    <span
      class="text-muted tabular col-span-2 row-start-3 text-right font-mono text-[13px] md:col-span-1 md:row-auto"
      data-testid="row-per-second"
    >
      <span class="flex flex-col leading-tight">
        <span title={actor.time_ms === undefined ? undefined : 'Per second of the pulls this player was in'}
          >{formatPerSecond(shownEffective, actor.time_ms ?? durationMs)}<span
            class="label font-body ml-1.5 md:hidden">per sec</span
          ></span
        >
        {#if actor.active_ms > 0 && actor.active_ms < durationMs}
          <span
            class="text-[11px]"
            title="Per second over the time this row was active, not the whole window"
            data-testid="row-active-per-second"
            >{formatPerSecond(shownEffective, actor.active_ms)} while active</span
          >
        {/if}
      </span>
    </span>
    <span
      class="text-muted tabular hidden text-right font-mono text-[13px] md:inline"
      data-testid="row-active"
    >
      {activeText}
    </span>
  </button>

  {#if open}
    <div
      class="bg-card-top flex flex-col gap-4 px-2 py-3 md:flex-row md:flex-wrap md:items-start"
      data-testid="row-detail"
    >
      {#if approximate && measure !== undefined}
        <div class="flex flex-wrap items-center gap-3 md:basis-full" data-testid="row-measure">
          {#if exact === null}
            <button
              type="button"
              class="border-line-warm rounded-control text-text inline-flex h-11 items-center border px-3 text-[12px] font-bold tracking-[0.06em] uppercase md:h-9"
              disabled={measuring}
              onclick={() => void runMeasure()}
            >
              {measuring ? 'Measuring…' : 'Measure this window exactly'}
            </button>
            <span class="text-muted text-[12px]"
              >The split below is prorated (~). Measuring reads the fight’s events.</span
            >
          {:else}
            <span class="text-kill text-[12px]" data-testid="row-measured"
              >Measured exactly from the fight’s events.</span
            >
          {/if}
          {#if measureError !== ''}<span class="text-wipe text-[12px]" role="alert">{measureError}</span>{/if}
        </div>
      {/if}
      {#if splitUnavailable}
        <p class="text-muted text-[13px] md:basis-full" data-testid="row-split-unavailable">
          Over the whole night the split by ability under a target or boss filter can only be guessed from
          each pull’s whole, so it is not shown. Open a pull to read it from that fight’s events.
        </p>
      {/if}
      <p class="text-muted text-[11px] md:hidden" hidden={splitUnavailable}>
        Swipe the table sideways for the other columns, a column at a time; the ability column stays put.
      </p>
      <p
        class="label text-muted pb-1 text-left"
        id="row-abilities-caption-{actor.guid}"
        hidden={splitUnavailable}
      >
        Abilities{#if schoolSplit.length > 1}
          <span class="ml-3 tracking-normal normal-case" data-testid="school-split"
            >{#each schoolSplit as part (part.name)}<span class="mr-3 inline-flex items-center gap-1"
                ><span
                  class="inline-block h-[8px] w-[8px] rounded-[2px]"
                  style={`background: ${part.token}`}
                  aria-hidden="true"
                ></span>{part.name} <span class="tabular font-mono">{part.pct.toFixed(0)}%</span></span
              >{/each}</span
          >{/if}
      </p>
      <!-- Its own scroller, so the hint and heading above stay put while the table moves;
           separate borders, so the pinned column's shadow can paint (a collapsed table
           drops cell shadows); one snap stop per column, so a number is whole or absent. -->
      <div
        class="-mx-2 snap-x snap-mandatory scroll-pl-[150px] overflow-x-auto px-2 md:mx-0 md:flex-1 md:snap-none md:overflow-visible md:px-0"
        hidden={splitUnavailable}
      >
        <table
          class="w-max border-separate border-spacing-0 text-[13px] md:w-full"
          data-testid="row-abilities"
          aria-labelledby="row-abilities-caption-{actor.guid}"
        >
          <thead>
            <tr class="label text-muted border-line-soft border-b">
              <th
                scope="col"
                class="bg-bg border-line-soft sticky left-0 w-[150px] max-w-[150px] min-w-[150px] border-r py-1 pr-3 pl-2 text-left font-normal shadow-[6px_0_8px_-4px_rgba(0,0,0,0.6)] md:static md:w-auto md:max-w-none md:min-w-0 md:border-r-0 md:pl-0 md:shadow-none"
                >Ability</th
              >
              <th scope="col" class="py-1 pr-3 text-right font-normal" title="Effective amount in this window"
                >Amount</th
              >
              <th scope="col" class="py-1 pr-3 text-right font-normal" title="Share of this row's total"
                >Share</th
              >
              <th scope="col" class="w-[16%] min-w-[72px] py-1 pr-3 font-normal" aria-label="Share, drawn"
              ></th>
              <th scope="col" class="py-1 pr-3 text-right font-normal" title="Hits and ticks that landed"
                >Hits</th
              >
              <th
                scope="col"
                class="py-1 pr-3 text-right font-normal"
                title="Share of the hits that were critical">Crit</th
              >
              <th scope="col" class="py-1 pr-3 text-right font-normal" title="Amount per hit">Avg</th>
              <th scope="col" class="py-1 pr-3 text-right font-normal" title="Largest single hit">Max</th>
              <th
                scope="col"
                class="py-1 text-left font-normal whitespace-nowrap md:min-w-[220px] md:text-right"
                aria-label="Notes"
              ></th>
            </tr>
          </thead>
          <tbody>
            {#each detailRows as ability (abilityKey(ability))}
              {@const hits = landed(ability)}
              {@const school = schoolToken(ability.school)}
              <tr class="border-line-soft border-b">
                <td
                  class="bg-bg border-line-soft sticky left-0 w-[150px] max-w-[150px] min-w-[150px] border-r py-1.5 pr-3 pl-2 shadow-[6px_0_8px_-4px_rgba(0,0,0,0.6)] md:static md:w-auto md:max-w-none md:min-w-0 md:border-r-0 md:pl-0 md:shadow-none"
                  >{ability.name}{#if ability.via}
                    <span
                      class="text-muted ml-1 text-[11px]"
                      title="Cast by this pet or guardian, counted on its owner's row">· {ability.via}</span
                    >{/if}{#if sameName.has(ability.name)}
                    <span
                      class="text-muted ml-1 font-mono text-[11px]"
                      title={ability.spell_id === 0
                        ? 'Two things share this name; this one is the auto-attack swing, which has no spell id'
                        : `Two spells share this name; this is spell id ${ability.spell_id}`}
                      >{ability.spell_id === 0 ? 'swing' : `#${ability.spell_id}`}</span
                    >{/if}{#if schoolName(ability.school)}
                    <span class="text-muted ml-1 text-[11px]">{schoolName(ability.school)}</span>{/if}</td
                >
                <td
                  class="tabular py-1.5 pr-3 text-right font-mono whitespace-nowrap"
                  title={detailTitle}
                  aria-label={approximateAriaLabel(approximate, formatAmount(ability.effective))}
                >
                  {detailMark}{formatAmount(ability.effective)}
                </td>
                <td class="text-muted tabular py-1.5 pr-3 text-right font-mono whitespace-nowrap"
                  >{detailTotal === 0 ? '' : formatPercent((ability.effective / detailTotal) * 100)}</td
                >
                <td class="py-1.5 pr-3">
                  {#if detailTotal > 0 && ability.effective > 0}
                    <span class="bg-line-soft block h-[6px] w-full"
                      ><span
                        class="block h-full"
                        style={`width: ${(ability.effective / detailTotal) * 100}%; background: ${school}`}
                      ></span></span
                    >
                  {/if}
                </td>
                <td class="text-muted tabular py-1.5 pr-3 text-right font-mono whitespace-nowrap"
                  >{detailMark}{hits}</td
                >
                <td class="text-muted tabular py-1.5 pr-3 text-right font-mono whitespace-nowrap"
                  >{#if hits > 0}{formatPercent((ability.crits / hits) * 100)}{/if}</td
                >
                <td class="text-muted tabular py-1.5 pr-3 text-right font-mono whitespace-nowrap"
                  >{#if hits > 0 && ability.effective > 0}{formatAmount(ability.effective / hits)}{/if}</td
                >
                <td class="text-muted tabular py-1.5 pr-3 text-right font-mono whitespace-nowrap"
                  >{#if ability.max > 0}{formatAmount(ability.max)}{/if}</td
                >
                <!-- Wraps: a long list of what did not land must not push the table past its box. -->
                <td class="text-muted tabular min-w-[220px] py-1.5 text-right font-mono text-[12px]"
                  >{abilityNotes(ability).join(' · ')}</td
                >
              </tr>
            {/each}
          </tbody>
        </table>
        <CopyCsv lines={abilityCsv} />
      </div>
      <table class="min-w-[280px] flex-1 self-start text-[13px]" data-testid="row-targets">
        <caption class="label text-muted pb-1 text-left">{pairsLabel}</caption>
        <thead>
          <tr class="label text-muted border-line-soft border-b">
            <th scope="col" class="py-1 pr-3 text-left font-normal">Name</th>
            <th scope="col" class="py-1 pr-3 text-right font-normal" title="Effective amount in this window"
              >Amount</th
            >
            <th scope="col" class="py-1 text-right font-normal" title="Share of this row's total">Share</th>
          </tr>
        </thead>
        <tbody>
          {#each targetsByName as target (target.name)}
            <tr class="border-line-soft border-b">
              <td class="py-1.5 pr-3"
                >{target.name}{#if target.count > 1}
                  <span class="text-muted tabular font-mono text-[12px]">×{target.count}</span>{/if}</td
              >
              <td
                class="tabular py-1.5 pr-3 text-right font-mono whitespace-nowrap"
                title={detailTitle}
                aria-label={approximateAriaLabel(approximate, formatAmount(target.total))}
              >
                {detailMark}{formatAmount(target.total)}
              </td>
              <td class="text-muted tabular py-1.5 text-right font-mono whitespace-nowrap"
                >{targetsTotal === 0 ? '' : formatPercent((target.total / targetsTotal) * 100)}</td
              >
            </tr>
          {/each}
        </tbody>
      </table>
    </div>
  {/if}
</li>

<style>
  /* One snap stop per column of the phone's abilities table (see the scroller above). */
  [data-testid='row-abilities'] th {
    scroll-snap-align: start;
  }
</style>
