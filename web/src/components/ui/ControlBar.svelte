<!-- web/src/components/ui/ControlBar.svelte -->
<!-- The one control bar (web layout spec 2026-11-05 §3.4): a full-bleed 56 px row, sticky under the
     header, hairlines above and below, content in the page column. Left to right: segmented groups,
     then selects, then the `end` content. Group labels show from 1280 (or 1920, `labelsFrom`). Below 768 it is one 48 px line,
     the choices in words and "Change", that expands the controls in place, stacked and full width,
     and collapses on a choice. A segmented group is a radio group of links (`?param=value`, or the
     option's own href for a path choice such as the spec), so it works before any script; the URL
     is the truth (lib/ui/control-bar.ts). Static pages use ControlBar.astro, which boots it; an
     island passes `onchange`. -->
<script lang="ts">
  import type { Snippet } from 'svelte';
  import { checkRadio } from '../../lib/ui/control-bar';
  import { controlSummary, type ControlGroup, type ControlOption } from '../../lib/ui/primitives';
  import Glyph from './Glyph.svelte';
  import './segmented.css';
  import Select from './Select.svelte';

  let {
    groups,
    label,
    summary,
    end,
    tail,
    sticky = true,
    labelsFrom = 1280,
    staticField = false,
    onchange,
    id = 'control-bar',
    testid,
  }: {
    groups: readonly ControlGroup[];
    /** The bar's accessible name ("Best in slot choices"). */
    label: string;
    /** The phone line; defaults to each group's chosen label joined with a middle dot. */
    summary?: string;
    end?: Snippet;
    /** Groups the bar's own kinds do not cover (a multi-select), after the groups. */
    tail?: Snippet;
    sticky?: boolean;
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
    if (onchange === undefined || option.href !== undefined) return;
    event.preventDefault();
    if (option.disabled) return;
    const change = checkRadio(event.currentTarget as HTMLElement);
    if (change) onchange(change.param, change.value);
    open = false;
  }
</script>

<div
  class={['ui-cbar full-bleed', sticky && 'ui-cbar--sticky']}
  data-control-bar
  data-open={open ? '' : undefined}
  data-labels-from={labelsFrom}
  role="group"
  aria-label={label}
  data-testid={testid}
>
  <div class="page-column ui-cbar-inner">
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
        >Change <Glyph name="chevron" size={14} class="rotate-90" /></span
      >
    </button>
    <div class="ui-cbar-groups" id={`${id}-groups`}>
      {#each groups as group (group.param)}
        <div class="ui-cbar-group" data-kind={group.kind}>
          <span
            class="ui-cbar-label text-label tracking-label text-muted font-bold uppercase"
            id={`${id}-${group.param}-label`}>{group.label}</span
          >
          {#if group.kind === 'segmented'}
            <div
              class="ui-seg"
              role="radiogroup"
              aria-labelledby={`${id}-${group.param}-label`}
              data-param={group.param}
              data-testid={`control-${group.param}`}
            >
              {#each group.options as option (option.value)}
                {@const checked = option.value === group.value}
                <a
                  class="ui-seg-item"
                  role="radio"
                  href={href(group, option)}
                  aria-checked={checked}
                  aria-disabled={option.disabled ? 'true' : undefined}
                  title={option.disabled ? (option.disabledReason ?? 'No data yet at this level') : undefined}
                  tabindex={checked ? 0 : -1}
                  data-value={option.value}
                  data-navigate={option.href !== undefined ? 'true' : undefined}
                  data-testid={`control-${group.param}-${option.value}`}
                  onclick={(event) => choose(event, option)}
                >
                  {#if option.icon}<img
                      src={option.icon}
                      alt=""
                      width="22"
                      height="22"
                      class="rounded-[4px]"
                    />{/if}
                  <span>{option.label}</span>
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
  </div>
</div>

<style>
  .ui-cbar {
    border-block: 1px solid var(--color-line-soft);
    background: color-mix(in srgb, var(--color-bg) 94%, transparent);
    backdrop-filter: blur(6px);
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
    min-height: 48px;
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
  .ui-cbar-group {
    display: flex;
    flex-direction: column;
    gap: 4px;
    min-width: 0;
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
