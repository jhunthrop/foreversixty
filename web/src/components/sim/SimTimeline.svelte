<!-- web/src/components/sim/SimTimeline.svelte -->
<!-- The anchor and the loading skeleton are the page's (LazySection, id `timeline`); this is the card. -->
<!-- The saved sim's fight timeline (spec 2026-11-04-sim-results 3.7): the median run, one lane per
     ability, buff and debuff bars, the resource line, the Execute wash, a brush, a crosshair with
     its readout, tick tooltips and the cast list. A lazy chunk: lib/sim/timeline.ts holds the
     rules and lib/sim/timeline-draw.ts the marks, drawn as one SVG string so a mount stays cheap.
     Pointer, keyboard and touch all move the same crosshair; on a phone the plot scrolls sideways
     under a sticky icon column and the readout is pinned under the brush. -->
<script lang="ts">
  import { onMount } from 'svelte';
  import { clockLabel } from '../../lib/chart/time-axis';
  import { withShardOverlay, type ActionNames } from '../../lib/sim/action-names';
  import {
    actionVisuals,
    loadVisualTables,
    pendingVisual,
    type ActionVisuals,
  } from '../../lib/sim/action-visuals';
  import {
    buildTimeline,
    caretAt,
    caretLanes,
    closeUpWindow,
    COLLAPSED_CAST_ROWS,
    ROW_HEIGHT,
    layoutRows,
    moveCaret,
    readoutAt,
    tickNear,
    windowIncluding,
    type Caret,
    type CaretKey,
    type DrawnTick,
    type PlotRow,
    type TickLane,
    type ZoomWindow,
  } from '../../lib/sim/timeline';
  import { timelineCopy } from '../../lib/sim/timeline-copy';
  import { timelineHeadMeta } from '../../lib/sim/timeline-lanes';
  import { drawnTicks, plotMarkup, xOf, timeOf, type PlotView } from '../../lib/sim/timeline-draw';
  import type { SimResult } from '../../lib/sim/types';
  import { yieldToMain } from '../../lib/yield';
  import ActionLink from '../ui/ActionLink.svelte';
  import Card from '../ui/Card.svelte';
  import Switch from '../ui/Switch.svelte';
  import ActionIcon from './ActionIcon.svelte';
  import SimCastList from './SimCastList.svelte';
  import TimelineBrush from './TimelineBrush.svelte';
  import TimelineReadout from './TimelineReadout.svelte';
  import TimelineTickTip from './TimelineTickTip.svelte';

  let {
    result,
    names,
    onOpenInSim,
  }: {
    result: SimResult;
    names: ActionNames | null;
    /** The band's "Open in the sim" handoff, for the older-result line's link. */
    onOpenInSim?: () => void;
  } = $props();

  /** The widths the spec fixes (3.7, Plot and Phone). */
  const LABEL_WIDTH = 200;
  const LABEL_WIDTH_LONG = 240;
  const LABEL_WIDTH_PHONE = 44;
  /** A name wider than the 200 px column's text room widens the column, by 14 px Barlow's average advance. */
  const LABEL_TEXT_ROOM_PX = 142;
  const NAME_ADVANCE_PX = 6.6;
  const PHONE_QUERY = '(max-width: 767px)';
  const PHONE_WHOLE_PX_PER_S = 4;
  const PHONE_CLOSE_PX_PER_S = 12;
  /** Before the first measure: the page column less the label column. Also what a server render draws. */
  const DEFAULT_PLOT_WIDTH = 1144;
  const READOUT_WIDTH = 280;
  const TIP_WIDTH = 250;
  const TIP_HEIGHT = 140;
  const READOUT_TOP = 36;
  const PANEL_GAP = 12;
  const HIT_TOLERANCE_PX = 5;
  const CAST_LIST_RESERVE = ROW_HEIGHT.castFoot + ROW_HEIGHT.castRow * COLLAPSED_CAST_ROWS;
  const BROWSER = typeof window !== 'undefined';

  // An item proc or a racial no class table names is read from the build's shards, once.
  let shardedNames = $state.raw<ActionNames | null>(null);
  const liveNames = $derived(shardedNames ?? names);
  $effect(() => {
    const base = names;
    if (base === null) return;
    const keys = [
      ...(result.sample ?? []).map((event) => event.action),
      ...(result.sample_auras ?? []).map((aura) => aura.action),
    ];
    let cancelled = false;
    withShardOverlay(base, new Set(keys)).then(
      (named) => {
        if (!cancelled && named !== base) shardedNames = named;
      },
      () => {
        // The shards failing leaves the class tables' names: "An unnamed spell" is the honest read.
      },
    );
    return () => {
      cancelled = true;
    };
  });

  const model = $derived(buildTimeline(result, liveNames));

  // --- The build's icons and ranks: names are the only input the first paint needs. ---
  let tables = $state.raw<Awaited<ReturnType<typeof loadVisualTables>> | null>(null);
  const visuals = $derived<ActionVisuals>(tables === null ? pendingVisual : actionVisuals(tables, liveNames));
  const visualsLoaded = $derived(tables !== null);

  // The card draws in three tasks (spec 8, under 50 ms of blocking at 4x CPU): the head and the
  // toolbar over reserved space, then the plot, then the brush and the cast list. A server render
  // has no tasks to spread, so it draws them all.
  const ALL_STAGES = 2;
  let drawnStage = $state(BROWSER ? 0 : ALL_STAGES);

  onMount(() => {
    let cancelled = false;
    void loadVisualTables(result).then((loaded) => {
      if (!cancelled) tables = loaded;
    });
    void (async () => {
      while (drawnStage < ALL_STAGES) {
        await yieldToMain();
        if (cancelled) return;
        drawnStage += 1;
      }
    })();
    const phoneQuery = window.matchMedia(PHONE_QUERY);
    const syncPhone = (): void => {
      phone = phoneQuery.matches;
    };
    phoneQuery.addEventListener('change', syncPhone);
    return () => {
      cancelled = true;
      phoneQuery.removeEventListener('change', syncPhone);
    };
  });

  // --- The view ---
  let phone = $state(BROWSER && window.matchMedia(PHONE_QUERY).matches);
  let rootWidth = $state(0);
  /** The phone notice's sticky line: the scroller less its 48px label gutter, its 8px padding
   *  each side and the scrollbar allowance (the card's own width, never the viewport's). */
  const NOTICE_INSET = 98;
  let closeUp = $state(false);
  let closeWindow = $state<ZoomWindow | null>(null);
  let swings = $state(false);

  const whole = $derived<ZoomWindow>(
    model === null ? { startMs: 0, endMs: 1 } : { startMs: model.startMs, endMs: model.endMs },
  );
  const viewWindow = $derived(closeUp && closeWindow !== null ? closeWindow : whole);

  const layout = $derived(model === null ? null : layoutRows(model, swings, phone));
  const rows = $derived<readonly PlotRow[]>(layout?.rows ?? []);
  const plotHeight = $derived(layout?.height ?? 0);

  const labelWidth = $derived.by(() => {
    if (phone) return LABEL_WIDTH_PHONE;
    const widest = rows.reduce(
      (most, row) => (row.kind === 'lane' ? Math.max(most, row.lane.name.length) : most),
      0,
    );
    return widest * NAME_ADVANCE_PX > LABEL_TEXT_ROOM_PX ? LABEL_WIDTH_LONG : LABEL_WIDTH;
  });
  const plotWidth = $derived.by(() => {
    if (phone) {
      const perSecond = closeUp ? PHONE_CLOSE_PX_PER_S : PHONE_WHOLE_PX_PER_S;
      return Math.round(((viewWindow.endMs - viewWindow.startMs) / 1000) * perSecond);
    }
    if (rootWidth > 0) return Math.max(0, rootWidth - labelWidth);
    // Nothing is drawn in a browser until the first measure; a server has none to wait for.
    return BROWSER ? 0 : DEFAULT_PLOT_WIDTH;
  });
  const brushWidth = $derived(phone ? Math.max(0, rootWidth - LABEL_WIDTH_PHONE) : plotWidth);
  const view = $derived<PlotView>({ window: viewWindow, width: plotWidth, closeUp });

  const drawable = $derived(drawnStage >= 1 && plotWidth > 0);
  const ticks = $derived(drawable ? drawnTicks(rows, view) : new Map<string, DrawnTick[]>());
  const markup = $derived(
    model === null || !drawable ? '' : plotMarkup(model, rows, plotHeight, ticks, view),
  );

  // --- The crosshair, its readout and the tooltip ---
  let crosshairMs = $state<number | null>(null);
  let pinned = $state(false);
  let caret = $state<Caret | null>(null);
  let hoverTip = $state<{ lane: TickLane; tick: DrawnTick; y: number } | null>(null);
  let announcement = $state('');
  let readoutHeight = $state(0);
  let lastPointerType = '';

  const lanes = $derived(model === null ? [] : caretLanes(model, swings));
  const readout = $derived(model === null || crosshairMs === null ? null : readoutAt(model, crosshairMs));

  const caretTip = $derived.by(() => {
    if (caret === null) return null;
    const lane = lanes[caret.lane];
    const event = lane?.events[caret.index];
    if (lane === undefined || event === undefined) return null;
    const row = rows.find((candidate) => candidate.kind === 'lane' && candidate.lane.id === lane.id);
    return {
      lane,
      index: caret.index,
      x: xOf(event.atMs, view),
      y: row?.y ?? 0,
      height: row?.height ?? 0,
    };
  });
  const tip = $derived(
    hoverTip === null
      ? caretTip
      : {
          lane: hoverTip.lane,
          index: hoverTip.tick.index,
          x: hoverTip.tick.x,
          y: hoverTip.y,
          height: 24,
        },
  );

  const crosshairX = $derived(crosshairMs === null ? null : xOf(crosshairMs, view));
  const crosshairInView = $derived(crosshairX !== null && crosshairX >= 0 && crosshairX <= plotWidth);
  const readoutLeft = $derived.by(() => {
    if (crosshairX === null) return 0;
    const right = crosshairX + PANEL_GAP;
    return right + READOUT_WIDTH > plotWidth ? Math.max(0, crosshairX - PANEL_GAP - READOUT_WIDTH) : right;
  });
  const readoutFlipped = $derived(crosshairX !== null && readoutLeft < crosshairX);
  const tipLeft = $derived.by(() => {
    if (tip === null) return 0;
    const left = tip.x - PANEL_GAP - TIP_WIDTH;
    const right = tip.x + PANEL_GAP;
    // The side away from the readout, so the two never sit on each other around one line.
    const wanted = crosshairX !== null && readoutFlipped ? right : left;
    return Math.max(0, Math.min(plotWidth - TIP_WIDTH, wanted));
  });
  const tipTop = $derived.by(() => {
    if (tip === null) return 0;
    const atLane = Math.max(0, Math.min(plotHeight - TIP_HEIGHT, tip.y - 8));
    // No room beside the readout: the tooltip drops below it instead of sitting on it.
    const overReadout =
      readout !== null && tipLeft < readoutLeft + READOUT_WIDTH && readoutLeft < tipLeft + TIP_WIDTH;
    return overReadout ? Math.max(atLane, READOUT_TOP + readoutHeight + PANEL_GAP) : atLane;
  });

  function clamp(atMs: number): number {
    return Math.max(viewWindow.startMs, Math.min(viewWindow.endMs, atMs));
  }

  function pointIn(event: MouseEvent): { x: number; y: number } {
    const box = (event.currentTarget as HTMLElement).getBoundingClientRect();
    return { x: event.clientX - box.left, y: event.clientY - box.top };
  }

  function tickAt(x: number, y: number): { lane: TickLane; tick: DrawnTick; y: number } | null {
    const row = rows.find((candidate) => y >= candidate.y && y < candidate.y + candidate.height);
    if (row?.kind !== 'lane' || row.lane.kind !== 'ticks') return null;
    const tick = tickNear(ticks.get(row.lane.id) ?? [], x, HIT_TOLERANCE_PX);
    return tick === null ? null : { lane: row.lane, tick, y: row.y };
  }

  function onPointerMove(event: PointerEvent): void {
    if (event.pointerType === 'touch') return;
    const { x, y } = pointIn(event);
    hoverTip = tickAt(x, y);
    // On a tick the line sits on the tick, so its readout and its tooltip read the same moment.
    if (!pinned) crosshairMs = hoverTip === null ? clamp(timeOf(x, view)) : hoverTip.tick.event.atMs;
  }

  function onPointerLeave(): void {
    hoverTip = null;
    if (!pinned && caret === null) crosshairMs = null;
  }

  function onPlotClick(event: MouseEvent): void {
    if ((event.target as Element).closest('[data-timeline-panel]') !== null) return;
    const { x, y } = pointIn(event);
    if (lastPointerType === 'touch') {
      const hit = tickAt(x, y);
      hoverTip = hit;
      crosshairMs = hit === null ? clamp(timeOf(x, view)) : hit.tick.event.atMs;
      pinned = hit === null;
      return;
    }
    if (pinned) {
      pinned = false;
      return;
    }
    crosshairMs = clamp(timeOf(x, view));
    pinned = true;
  }

  function unpin(): void {
    pinned = false;
    hoverTip = null;
    if (caret === null) crosshairMs = null;
  }

  function moveTo(next: Caret): void {
    caret = next;
    hoverTip = null;
    const event = lanes[next.lane]?.events[next.index];
    if (event === undefined || model === null) return;
    crosshairMs = event.atMs;
    if (closeUp && closeWindow !== null) closeWindow = windowIncluding(closeWindow, event.atMs, model);
    announce(event.atMs);
  }

  function announce(atMs: number): void {
    if (model === null) return;
    const line = readoutAt(model, atMs);
    const state = line.resource === null ? '' : `${line.resource.label} ${line.resource.value}, `;
    announcement = `${line.time}. ${state}${timelineCopy.readoutOnYou(line.onYou)}`;
  }

  const CARET_KEYS: ReadonlySet<string> = new Set([
    'ArrowLeft',
    'ArrowRight',
    'ArrowUp',
    'ArrowDown',
    'Home',
    'End',
  ]);

  function onPlotKey(event: KeyboardEvent): void {
    if (CARET_KEYS.has(event.key)) {
      event.preventDefault();
      const next = moveCaret(lanes, caret, event.key as CaretKey);
      if (next !== null) moveTo(next);
      return;
    }
    if (event.key === 'Enter' && crosshairMs !== null) {
      pinned = true;
    } else if (event.key === 'Escape') {
      pinned = false;
    } else if (event.key === '+' || event.key === '=') {
      setZoom(true);
    } else if (event.key === '-') {
      setZoom(false);
    }
  }

  function onPlotFocus(event: FocusEvent): void {
    const plot = event.currentTarget as HTMLElement;
    if (caret !== null || !plot.matches(':focus-visible')) return;
    const next = crosshairMs === null ? moveCaret(lanes, null, 'Home') : caretAt(lanes, crosshairMs);
    if (next !== null) moveTo(next);
  }

  function onPlotBlur(): void {
    caret = null;
    if (!pinned) crosshairMs = null;
  }

  function setZoom(next: boolean): void {
    if (model === null || next === closeUp) return;
    hoverTip = null;
    if (next) closeWindow = closeUpWindow(model, crosshairMs);
    closeUp = next;
  }

  function setSwings(next: boolean): void {
    swings = next;
    hoverTip = null;
    caret = null;
  }

  function openAt(atMs: number): void {
    if (model === null) return;
    closeWindow = closeUpWindow(model, atMs);
    closeUp = true;
  }

  // --- The legend ---
  const damageToken = $derived(model?.casts.find((lane) => lane.schoolToken !== null)?.schoolToken ?? null);
  const hasCooldown = $derived(model?.casts.some((lane) => lane.schoolToken === null) ?? false);
  /** The same meta the loading head printed (SavedRun), so nothing pops in on load. */
  const meta = $derived(timelineHeadMeta(result));
</script>

{#snippet laneLabel(row: Extract<PlotRow, { kind: 'lane' }>)}
  {@const visual = visuals(row.lane.action)}
  <!-- The shared tooltip host opens on [data-game-tip]; a phone's tap reaches it the same way. -->
  <div
    class="lane-label"
    style:height="{row.height}px"
    data-game-tip={visual.tipKey ?? undefined}
    title={phone ? row.lane.name : undefined}
    data-testid="timeline-lane-label"
  >
    <span class="icon-cell"
      ><ActionIcon
        url={visual.icon}
        loaded={visualsLoaded}
        size={18}
        label={phone ? row.lane.name : ''}
      /></span
    >
    {#if !phone}<span>{row.lane.name}</span>{/if}
  </div>
{/snippet}

<Card
  glyph="clock"
  label={timelineCopy.label}
  title={timelineCopy.title}
  {meta}
  body="none"
  class="mx-[18px] md:mx-0"
  testid="sim-timeline"
>
  {#if model === null}
    <p class="text-text border-line-soft border-t px-5 py-4 text-[15px]" data-testid="timeline-no-sample">
      {timelineCopy.noSample}
    </p>
  {:else}
    <div
      class="border-line-soft flex min-h-14 flex-wrap items-center gap-x-4 gap-y-[10px] border-t px-4 py-[10px] md:px-5"
      data-testid="timeline-toolbar"
    >
      <div class="seg" role="group" aria-label={timelineCopy.zoomLabel} data-testid="timeline-zoom">
        <button
          type="button"
          aria-pressed={!closeUp}
          onclick={() => setZoom(false)}
          data-testid="timeline-whole">{timelineCopy.wholeFight}</button
        >
        <button
          type="button"
          aria-pressed={closeUp}
          onclick={() => setZoom(true)}
          data-testid="timeline-close">{timelineCopy.closeUp}</button
        >
      </div>
      <span class="inline-flex items-center gap-[18px]">
        <Switch
          checked={swings}
          onchange={setSwings}
          labelledby="timeline-swings-label"
          testid="timeline-swings"
        />
        <span id="timeline-swings-label" class="text-text text-[14px] font-semibold"
          >{timelineCopy.swingsAndTicks}</span
        >
      </span>
      <ul
        class="text-muted flex flex-wrap items-center gap-x-4 gap-y-1 text-[13px] md:ml-auto"
        data-testid="timeline-legend"
      >
        {#if damageToken !== null}
          <li>
            <i class="swatch" style:width="3px" style:height="12px" style:background="var({damageToken})"
            ></i>{timelineCopy.legend.damage}
          </li>
        {/if}
        {#if model.hasOutcomes}
          <li>
            <i class="swatch" style:width="4px" style:height="16px" style:background="var(--color-gold-light)"
            ></i>{timelineCopy.legend.crit}
          </li>
          <li>
            <i
              class="swatch"
              style:width="3px"
              style:height="12px"
              style:border="1px solid var(--color-muted)"
            ></i>{timelineCopy.legend.missed}
          </li>
        {/if}
        {#if hasCooldown}
          <li>
            <i class="swatch" style:width="3px" style:height="12px" style:background="var(--color-text)"
            ></i>{timelineCopy.legend.cooldown}
          </li>
        {/if}
        {#if model.onYou.length > 0}
          <li>
            <i
              class="swatch"
              style:width="16px"
              style:height="8px"
              style:background="var(--color-text)"
              style:opacity="0.35"
            ></i>{timelineCopy.legend.onYou}
          </li>
        {/if}
        {#if model.onTarget.length > 0}
          <li>
            <i
              class="swatch"
              style:width="16px"
              style:height="8px"
              style:background="var(--color-gold-deep)"
              style:opacity="0.78"
            ></i>{timelineCopy.legend.onTarget}
          </li>
        {/if}
        {#if model.resource !== null}
          <li>
            <i class="swatch" style:width="16px" style:height="2px" style:background="var(--color-gold)"
            ></i>{model.resource.label}
          </li>
        {/if}
      </ul>
    </div>

    <div class="relative" bind:clientWidth={rootWidth}>
      <div
        class="scroller"
        class:phone
        style:--notice-width={`${Math.max(0, rootWidth - NOTICE_INSET)}px`}
        data-testid="timeline-scroller"
      >
        <div class="flex" style:width={phone ? `${LABEL_WIDTH_PHONE + plotWidth}px` : undefined}>
          <div class="labels" style:width="{labelWidth}px" data-testid="timeline-labels">
            {#each drawnStage >= 1 ? rows : [] as row (row.kind === 'lane' ? row.lane.id : `${row.kind}-${row.y}`)}
              {#if row.kind === 'axis'}
                <div class="axis-label" style:height="{row.height}px">
                  {#if !phone}{timelineCopy.axisTime}{/if}
                </div>
              {:else if row.kind === 'group'}
                <div class="group-label" style:height="{row.height}px" data-testid="timeline-group">
                  <span>{row.label}</span>
                </div>
              {:else if row.kind === 'lane'}
                {@render laneLabel(row)}
              {:else if row.kind === 'notice' || row.kind === 'name'}
                <div style:height="{row.height}px"></div>
              {:else}
                <div class="resource-label" style:height="{row.height}px">
                  <b>{row.line.label}</b>
                  {#if !phone}
                    <span data-testid="timeline-resource-peak"
                      >{timelineCopy.resourcePeak(row.line.peak, row.line.cap)}</span
                    >
                  {/if}
                </div>
              {/if}
            {/each}
          </div>

          <!-- svelte-ignore a11y_no_noninteractive_tabindex, a11y_no_noninteractive_element_interactions -->
          <div
            class="plot"
            role="application"
            aria-roledescription={timelineCopy.plotRole}
            aria-label={timelineCopy.plotLabel(clockLabel(model.endMs), model.presses.length)}
            tabindex="0"
            style:width="{plotWidth}px"
            style:height="{plotHeight}px"
            onpointerdown={(event) => (lastPointerType = event.pointerType)}
            onpointermove={onPointerMove}
            onpointerleave={onPointerLeave}
            onclick={onPlotClick}
            onkeydown={onPlotKey}
            onfocus={onPlotFocus}
            onblur={onPlotBlur}
            data-testid="timeline-plot"
          >
            <svg
              class="absolute inset-0 block overflow-visible"
              width={plotWidth}
              height={plotHeight}
              viewBox={`0 0 ${plotWidth} ${plotHeight}`}
              font-family="var(--font-body)"
              aria-hidden="true"
            >
              <!-- eslint-disable-next-line svelte/no-at-html-tags -- numbers and constants only: lib/sim/timeline-draw.ts carries no text from the result -->
              {@html markup}
            </svg>

            {#if phone}
              <!-- Each name on its own row above its lane: it never sits over a bar or a tick (B1). -->
              {#each rows as row (`${row.kind}-${row.y}`)}
                {#if row.kind === 'name'}
                  <div
                    class="name-row"
                    style:top="{row.y}px"
                    style:height="{row.height}px"
                    aria-hidden="true"
                  >
                    <span class="name-chip" data-testid="timeline-lane-chip">{row.lane.name}</span>
                  </div>
                {/if}
              {/each}
            {/if}

            {#each rows as row (`notice-${row.y}`)}
              {#if row.kind === 'notice'}
                <div
                  class="notice-row"
                  style:top="{row.y}px"
                  style:height="{row.height}px"
                  data-timeline-panel
                >
                  <div class="notice" data-testid="timeline-auras-fallback">
                    <span class="text-muted text-[13px]">{timelineCopy.aurasFallback}</span>
                    <ActionLink
                      href="/sim"
                      onclick={(event: MouseEvent) => {
                        if (onOpenInSim === undefined) return;
                        event.preventDefault();
                        onOpenInSim();
                      }}>{timelineCopy.openInSim}</ActionLink
                    >
                  </div>
                </div>
              {/if}
            {/each}

            {#if tip !== null}
              <span
                class="tick-mark"
                style:left="{tip.x - 4}px"
                style:top="{tip.y + 1}px"
                style:height="{tip.height - 2}px"
                aria-hidden="true"
                data-testid="timeline-active-tick"
              ></span>
            {/if}

            {#if crosshairInView && crosshairX !== null && crosshairMs !== null}
              <span class="crosshair" style:left="{crosshairX}px" aria-hidden="true"></span>
              <span
                class="time-chip numeric"
                style:left="{crosshairX - 26}px"
                data-testid="timeline-time-chip">{clockLabel(Math.round(crosshairMs / 100) * 100)}</span
              >
            {/if}

            {#if !phone}
              {#if readout !== null && crosshairInView}
                <div
                  class="float"
                  style:width="{READOUT_WIDTH}px"
                  style:left="{readoutLeft}px"
                  style:top="{READOUT_TOP}px"
                  bind:clientHeight={readoutHeight}
                  style:pointer-events={pinned ? 'auto' : 'none'}
                  data-timeline-panel
                >
                  <TimelineReadout {readout} {pinned} {visuals} {visualsLoaded} onclose={unpin} />
                </div>
              {/if}
              {#if tip !== null}
                <div
                  class="float"
                  style:width="{TIP_WIDTH}px"
                  style:left="{tipLeft}px"
                  style:top="{tipTop}px"
                  style:pointer-events="none"
                  data-timeline-panel
                >
                  <TimelineTickTip
                    lane={tip.lane}
                    index={tip.index}
                    resourceLabel={model.resource?.label ?? null}
                    {visuals}
                    {visualsLoaded}
                  />
                </div>
              {/if}
            {/if}
          </div>
        </div>
      </div>
      {#if phone}<span class="fade" aria-hidden="true"></span>{/if}
    </div>

    <div class="border-line-soft flex min-h-11 border-t" data-testid="timeline-brush-row">
      <div class="brush-gutter" style:width="{labelWidth}px"></div>
      {#if drawnStage >= 2}
        <TimelineBrush
          {model}
          window={closeUp ? viewWindow : null}
          width={brushWidth}
          onWindow={(next) => (closeWindow = next)}
          onOpen={openAt}
        />
      {/if}
    </div>

    {#if phone && (tip !== null || (readout !== null && pinned))}
      <div class="flex flex-col gap-2 p-4" data-testid="timeline-phone-panel" data-timeline-panel>
        {#if tip !== null}
          <TimelineTickTip
            lane={tip.lane}
            index={tip.index}
            resourceLabel={model.resource?.label ?? null}
            {visuals}
            {visualsLoaded}
          />
        {:else if readout !== null}
          <TimelineReadout {readout} {pinned} {visuals} {visualsLoaded} onclose={unpin} />
        {/if}
      </div>
    {/if}

    {#if drawnStage >= 2}
      <SimCastList {model} window={viewWindow} {closeUp} names={liveNames} {visuals} {visualsLoaded} />
    {:else}
      <div style:height="{CAST_LIST_RESERVE}px"></div>
    {/if}

    <p class="sr-only" aria-live="polite" data-testid="timeline-live">{announcement}</p>
  {/if}
</Card>

<style>
  .seg {
    display: inline-flex;
    overflow: hidden;
    border: 1px solid var(--color-line);
    border-radius: var(--radius-control);
  }
  .seg button {
    display: inline-flex;
    align-items: center;
    height: 36px;
    padding: 0 14px;
    color: var(--color-nav);
    font-size: 14px;
    font-weight: 600;
    white-space: nowrap;
  }
  .seg button[aria-pressed='true'] {
    background: var(--tile);
    color: var(--tile-ink);
    box-shadow: inset 0 0 0 1px rgba(229, 185, 85, 0.35);
  }
  .seg button:focus-visible {
    outline: 2px solid var(--color-gold);
    outline-offset: -2px;
  }
  @media (hover: hover) and (pointer: fine) {
    .seg button[aria-pressed='false']:hover {
      color: var(--color-strong);
    }
  }
  @media (pointer: coarse), (max-width: 767px) {
    .seg button {
      height: 44px;
    }
  }
  li:has(> .swatch) {
    display: inline-flex;
    align-items: center;
    gap: 6px;
  }
  .swatch {
    display: inline-block;
    border-radius: 1px;
  }

  .scroller {
    border-top: 1px solid var(--color-line-soft);
  }
  .scroller.phone {
    overflow-x: auto;
    overscroll-behavior-x: contain;
    scrollbar-width: none;
  }
  .scroller.phone::-webkit-scrollbar {
    display: none;
  }
  .labels {
    flex: none;
    border-right: 1px solid var(--color-line-soft);
  }
  .phone .labels {
    position: sticky;
    left: 0;
    z-index: 3;
    background: #0d0f15;
  }
  .axis-label,
  .group-label {
    display: flex;
    align-items: flex-end;
    padding: 0 12px 6px 20px;
    color: var(--color-muted);
    font-size: 13px;
    font-weight: 600;
  }
  .group-label {
    padding-bottom: 4px;
  }
  .lane-label {
    display: flex;
    align-items: center;
    gap: 8px;
    padding: 0 12px 0 20px;
    color: var(--color-text);
    font-size: 14px;
    line-height: 1;
    white-space: nowrap;
  }
  /* The name keeps its column when a key has no icon: a blank cell, never an empty box (R1). */
  .icon-cell {
    display: inline-flex;
    flex: none;
    width: 18px;
    height: 18px;
  }
  .lane-label[data-game-tip] {
    cursor: help;
  }
  .resource-label {
    display: flex;
    flex-direction: column;
    justify-content: center;
    gap: 2px;
    padding: 0 12px 0 20px;
  }
  .resource-label b {
    color: var(--color-text);
    font-size: 14px;
    font-weight: 600;
  }
  .resource-label span {
    color: var(--color-muted);
    font-size: 13px;
  }
  .phone .axis-label,
  .phone .lane-label,
  .phone .resource-label {
    padding: 0 0 0 12px;
  }
  .phone .group-label {
    position: relative;
    padding: 0 0 4px 12px;
    overflow: visible;
  }
  /* The heading runs over the plot's left edge on the card's ground (spec 3.7, Phone). */
  .phone .group-label span {
    position: relative;
    z-index: 4;
    padding-right: 6px;
    background: #0d0f15;
    font-size: 12px;
    white-space: nowrap;
  }
  .phone .resource-label b {
    font-size: 11px;
  }

  .plot {
    position: relative;
    flex: none;
    cursor: crosshair;
    outline: none;
    touch-action: pan-x pan-y;
  }
  .plot:focus-visible {
    box-shadow: inset 0 0 0 2px var(--color-gold);
  }
  .crosshair {
    position: absolute;
    top: 26px;
    bottom: 0;
    width: 1px;
    background: var(--color-gold-hover);
    pointer-events: none;
  }
  .time-chip {
    position: absolute;
    top: 2px;
    z-index: 2;
    display: flex;
    align-items: center;
    justify-content: center;
    width: 52px;
    height: 20px;
    border-radius: 4px;
    background: var(--color-gold-hover);
    color: #1a1206;
    font-size: 13px;
    font-weight: 700;
    pointer-events: none;
  }
  .tick-mark {
    position: absolute;
    width: 8px;
    border: 2px solid var(--color-gold);
    border-radius: 2px;
    pointer-events: none;
  }
  .float {
    position: absolute;
    z-index: 4;
  }
  .notice-row {
    position: absolute;
    right: 0;
    left: 0;
    display: flex;
    align-items: center;
    background: #0d0f15;
    cursor: default;
  }
  .notice {
    display: flex;
    flex-wrap: wrap;
    align-items: center;
    gap: 2px 12px;
    padding: 0 16px;
  }
  /* Inside a sideways scroller, the line stays in the visible column. */
  .phone .notice {
    position: sticky;
    left: 48px;
    width: var(--notice-width);
    padding: 0 8px;
  }
  .name-row {
    position: absolute;
    right: 0;
    left: 0;
    display: flex;
    align-items: flex-end;
    pointer-events: none;
  }
  /* Its own row, so no backing: the name stays at the visible left edge as the plot scrolls. */
  .name-chip {
    position: sticky;
    left: 52px;
    color: var(--color-muted);
    font-size: 11px;
    line-height: 14px;
    white-space: nowrap;
  }
  .fade {
    position: absolute;
    top: 0;
    right: 0;
    bottom: 0;
    width: 24px;
    background: linear-gradient(90deg, rgba(13, 15, 21, 0), #0d0f15);
    pointer-events: none;
  }
  .brush-gutter {
    flex: none;
    border-right: 1px solid var(--color-line-soft);
  }
</style>
