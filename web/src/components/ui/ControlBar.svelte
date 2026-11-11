<!-- web/src/components/ui/ControlBar.svelte -->
<!-- The one control bar (web layout spec 2026-11-05 §3.4): a full-bleed 56 px row, sticky under the
     header, hairlines above and below, content in the page column. Left to right: segmented groups,
     then selects, then the `end` content. Group labels show from 1280 (or 1920, `labelsFrom`). Below 768 it is one 48 px line,
     the choices in words and "Change", that expands the controls in place, stacked and full width,
     and collapses on a choice. A segmented group is a radio group of links (`?param=value`, or the
     option's own href for a path choice such as the spec), so it works before any script; the URL
     is the truth (lib/ui/control-bar.ts). Static pages use ControlBar.astro, which boots it; an
     island passes `onchange`. A `compact` group (the class picker, ruling P10) shows its option
     icons (round ringed crests with `ring`) and its labels only from 1280 and in the open phone
     stack, where it lays out as a 3 × 3 grid. -->
<script lang="ts">
  import Crest from './Crest.svelte';
  import SpellIcon from './SpellIcon.svelte';
  import type { Snippet } from 'svelte';
  import { REGION } from '../../lib/ui/regions';
  import { checkRadio } from '../../lib/ui/control-bar';
  import { controlSummary, type ControlGroup, type ControlOption } from '../../lib/ui/primitives';
  import Icon from './Icon.svelte';
  import { ICON_CHEVRON_DOWN } from '../../lib/ui/ui-icons';
  import './segmented.css';
  import Select from './Select.svelte';

  let {
    groups,
    label,
    summary,
    end,
    action,
    busy = false,
    busyTitle,
    tail,
    sticky = true,
    toggleLabel = 'Change',
    labelsFrom = 1280,
    staticField = false,
    inline = false,
    onchange,
    id = 'control-bar',
    testid,
  }: {
    groups: readonly ControlGroup[];
    /** The bar's accessible name ("Best in slot choices"). */
    label: string;
    /** The phone line; defaults to each group's chosen label joined with a middle dot. */
    summary?: string;
    /** One short segmented group that fits a phone line (the rankings landing's Characters / Guilds):
     *  below 768 it shows as itself, with no summary line and no Change. */
    inline?: boolean;
    end?: Snippet;
    /** The page's one primary (the sim's Run). One element: on phone it sits on the summary line
     *  beside "Change", from 768 after the groups and `end`, so a run is never behind the toggle. */
    action?: Snippet;
    /** A run is in flight: every group and `end` at 50% and inert; `action` stays live (Stop). */
    busy?: boolean;
    /** The busy groups' title ("Stop the run to change it"). */
    busyTitle?: string;
    /** Groups the bar's own kinds do not cover (a multi-select), after the groups. */
    tail?: Snippet;
    sticky?: boolean;
    /** The phone toggle's word, "Change" by default; a bar that is mostly a jump menu says so. */
    toggleLabel?: string;
    /** The width the group labels show from: 1280 by default; 1920 for a bar whose choices name
     *  themselves and need the room to stay on one row (the BiS bar). */
    labelsFrom?: 1280 | 1920;
    /** Server-rendered for ControlBar.astro: its script wires the radios and selects. */
    staticField?: boolean;
    onchange?: (param: string, value: string) => void;
    id?: string;
    testid?: string;
  } = $props();

  let open = $state(false);
  const line = $derived(summary ?? controlSummary(groups));

  function href(group: ControlGroup, option: ControlOption): string {
    return option.href ?? `?${new URLSearchParams({ [group.param]: option.value }).toString()}`;
  }

  function choose(event: MouseEvent, option: ControlOption): void {
    if (onchange === undefined || (option.href !== undefined && option.inPlace !== true)) return;
    // A modified click on a page's own link (a new tab, a new window) keeps the link's own way.
    if (option.inPlace === true && (event.metaKey || event.ctrlKey || event.shiftKey || event.altKey)) return;
    event.preventDefault();
    if (option.disabled) return;
    const change = checkRadio(event.currentTarget as HTMLElement);
    if (change) onchange(change.param, change.value);
    open = false;
  }
</script>

<div
  class={['ui-cbar ui-band ui-band-end full-bleed', sticky && 'ui-cbar--sticky']}
  data-control-bar
  data-open={open ? '' : undefined}
  data-inline={inline ? '' : undefined}
  data-labels-from={labelsFrom}
  role="group"
  aria-label={label}
  data-testid={testid}
  data-region={REGION.controls}
>
  <div class="page-column ui-cbar-inner" data-has-action={action ? '' : undefined}>
    <button
      type="button"
      class="ui-cbar-toggle"
      aria-expanded={open}
      aria-controls={`${id}-groups`}
      data-control-bar-toggle
      onclick={staticField ? undefined : () => (open = !open)}
    >
      <span class="min-w-0 truncate">{line}</span>
      <span class="text-gold ml-auto flex items-center gap-1 text-[13px]"
        >{toggleLabel} <Icon icon={ICON_CHEVRON_DOWN} size={14} /></span
      >
    </button>
    <div
      class="ui-cbar-groups"
      id={`${id}-groups`}
      data-busy={busy ? '' : undefined}
      inert={busy}
      title={busy ? busyTitle : undefined}
    >
      {#each groups as group (group.param)}
        <div class="ui-cbar-group" data-kind={group.kind}>
          <span class="ui-cbar-label text-label text-muted font-semibold" id={`${id}-${group.param}-label`}
            >{group.label}</span
          >
          {#if group.kind === 'segmented'}
            <div
              class={['ui-seg', group.compact && 'ui-seg--compact']}
              role="radiogroup"
              aria-labelledby={`${id}-${group.param}-label`}
              data-param={group.param}
              data-default={group.defaultValue}
              data-testid={`control-${group.param}`}
            >
              {#each group.options as option (option.value)}
                {@const checked = option.value === group.value}
                <a
                  class="ui-seg-item"
                  role="radio"
                  href={href(group, option)}
                  aria-checked={checked}
                  aria-label={group.compact ? option.label : undefined}
                  aria-disabled={option.disabled ? 'true' : undefined}
                  title={option.disabled
                    ? (option.disabledReason ?? 'No data yet at this level')
                    : group.compact
                      ? option.label
                      : undefined}
                  tabindex={checked ? 0 : -1}
                  data-value={option.value}
                  data-navigate={option.href !== undefined && option.inPlace !== true ? 'true' : undefined}
                  data-testid={`control-${group.param}-${option.value}`}
                  onclick={(event) => choose(event, option)}
                >
                  {#if option.crest !== undefined}<Crest
                      cls={option.crest}
                      size={22}
                    />{:else if option.icon}<SpellIcon url={option.icon} size={22} />{/if}
                  {#if group.compact}<span class="ui-seg-lbl">{option.label}</span>{:else}<span
                      >{option.label}</span
                    >{/if}
                  {#if option.count !== undefined}<span class="ui-seg-count tabular-nums">{option.count}</span
                    >{/if}
                </a>
              {/each}
            </div>
          {:else}
            <div
              class="ui-cbar-select"
              data-param={group.param}
              data-fixed-width={group.fieldWidth === undefined ? undefined : 'true'}
              style:--cbar-field-w={group.fieldWidth === undefined ? undefined : `${group.fieldWidth}px`}
            >
              <Select
                id={`${id}-${group.param}`}
                label={group.label}
                size="dense"
                fit={group.fieldWidth === undefined}
                {staticField}
                value={group.value}
                options={group.options.map((option) => ({
                  value: option.value,
                  label: option.label,
                  description: option.description,
                  disabled: option.disabled,
                }))}
                testid={`control-${group.param}`}
                onchange={staticField || onchange === undefined
                  ? undefined
                  : (event) => {
                      onchange(group.param, (event.currentTarget as HTMLSelectElement).value);
                      open = false;
                    }}
              />
            </div>
          {/if}
        </div>
      {/each}
      {#if tail}{@render tail()}{/if}
      {#if end}<div class="ui-cbar-end">{@render end()}</div>{/if}
    </div>
    {#if action}<div class="ui-cbar-action">{@render action()}</div>{/if}
  </div>
</div>

<style>
  /* The controls sit on the band (round 5: their tab groups on the section band), continuing the
     title band's image from about where it ends, under a faint light rule. */
  .ui-cbar {
    --band-y: calc(var(--site-top-h, 80px) + 110px);
    border-top: 1px solid rgb(255 255 255 / 0.14);
    z-index: 30;
  }
  .ui-cbar--sticky {
    position: sticky;
    top: 0;
  }
  /* The phone inset only: from 768 .page-column's own gutter puts the bar on the content edge. */
  @media (max-width: 767px) {
    .ui-cbar-inner {
      padding-inline: var(--page-inset);
    }
  }
  .ui-cbar-toggle {
    display: flex;
    align-items: center;
    gap: 10px;
    width: 100%;
    min-height: var(--bar-h-phone);
    font-size: 14px;
    line-height: 20px;
    font-weight: 600;
    color: var(--color-strong);
    background: none;
    border: 0;
    padding: 0;
    text-align: left;
  }
  .ui-cbar-groups {
    display: none;
    flex-direction: column;
    gap: 12px;
    padding-bottom: 12px;
  }
  .ui-cbar[data-open] .ui-cbar-groups {
    display: flex;
  }
  @media (max-width: 767px) {
    .ui-cbar[data-inline] .ui-cbar-toggle,
    .ui-cbar[data-inline] .ui-cbar-label {
      display: none;
    }
    .ui-cbar[data-inline] .ui-cbar-groups {
      display: flex;
      padding-block: 6px;
    }
  }
  .ui-cbar-group {
    display: flex;
    flex-direction: column;
    gap: 4px;
    min-width: 0;
  }
  .ui-cbar-groups[data-busy] {
    opacity: 0.5;
  }
  /* The open phone bar wraps a segmented group that does not fit (guides ruling 0A MUST 5): the
     chosen item is never cut off. Items keep their equal share while they fit. */
  .ui-seg {
    flex-wrap: wrap;
  }
  .ui-seg-item {
    min-width: max-content;
  }
  /* With a primary: the summary line and the primary share the phone's first row, the groups open
     under them. */
  .ui-cbar-inner[data-has-action] {
    display: grid;
    grid-template-columns: minmax(0, 1fr) auto;
    column-gap: 12px;
    align-items: center;
  }
  .ui-cbar-inner[data-has-action] > .ui-cbar-toggle {
    grid-column: 1;
    grid-row: 1;
  }
  .ui-cbar-inner[data-has-action] > .ui-cbar-action {
    grid-column: 2;
    grid-row: 1;
  }
  .ui-cbar-inner[data-has-action] > .ui-cbar-groups {
    grid-column: 1 / -1;
    grid-row: 2;
  }

  @media (min-width: 768px) {
    .ui-cbar-inner {
      display: flex;
      align-items: center;
      min-height: var(--bar-h);
      padding-block: 6px;
    }
    .ui-cbar-toggle {
      display: none;
    }
    .ui-cbar-groups,
    .ui-cbar[data-open] .ui-cbar-groups {
      display: flex;
      flex-direction: row;
      flex-wrap: wrap;
      align-items: center;
      gap: 12px 20px;
      padding-bottom: 0;
      width: 100%;
    }
    .ui-cbar-group {
      flex-direction: row;
      align-items: center;
      gap: 8px;
    }
    .ui-cbar-label {
      display: none;
    }
    .ui-seg {
      display: inline-flex;
      flex-wrap: nowrap;
    }
    .ui-seg-item {
      flex: none;
      height: 40px;
    }
    /* With its labels hidden, a bar's groups sit closer: the gap was there to part a label from
       the group before it. */
    .ui-cbar[data-labels-from='1920'] .ui-cbar-groups {
      column-gap: 12px;
    }
    .ui-cbar-select[data-fixed-width='true'] {
      width: var(--cbar-field-w);
    }
    .ui-cbar-inner[data-has-action] {
      display: flex;
    }
    .ui-cbar-inner[data-has-action] > .ui-cbar-groups {
      flex: 1;
      width: auto;
      min-width: 0;
    }
    .ui-cbar-action {
      margin-left: 20px;
    }
    .ui-cbar-end {
      margin-left: auto;
      display: flex;
      gap: 16px;
      align-items: center;
    }
  }
  @media (min-width: 1280px) {
    .ui-cbar[data-labels-from='1280'] .ui-cbar-label {
      display: block;
    }
  }
  @media (min-width: 1920px) {
    .ui-cbar[data-labels-from='1920'] .ui-cbar-label {
      display: block;
    }
    .ui-cbar[data-labels-from='1920'] .ui-cbar-groups {
      column-gap: 20px;
    }
  }
</style>
