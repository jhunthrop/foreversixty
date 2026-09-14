// web/src/planner-island.ts
// Entry point for the standalone bundle the Go API loads on its server-rendered /b/:id
// page: <div id="planner" data-build='…' data-tree-version="…"> plus
// <script type="module" src="https://foreversixty.gg/planner-island.js">.
//
// The two stylesheet imports are deliberate, and load-bearing rather than incidental:
// planner-island.css is the only stylesheet that page links, and it renders the site's
// header and footer alongside the planner. The emitted file therefore has to carry the same
// tokens, base rules and utilities the site uses (global.css) and the same self-hosted faces
// (fonts.css, shared with src/layouts/Base.astro), so the API page matches the site instead
// of keeping a second copy of the design system or falling back to Georgia and Arial.
// src/planner-island.test.ts holds it to both.
import { mount } from 'svelte';
import Planner from './components/planner/Planner.svelte';
import { DEFAULT_CLASS_SLUG } from './lib/planner/config';
import { loadReference } from './lib/planner/load';
import type { BuildRecord } from './lib/planner/types';
import './styles/fonts.css';
import './styles/global.css';

const MOUNT_ID = 'planner';

function readRecord(element: HTMLElement): BuildRecord | null {
  const raw = element.dataset.build;
  if (!raw) return null;
  try {
    return JSON.parse(raw) as BuildRecord;
  } catch (error) {
    console.error('planner island: data-build is not valid JSON', error);
    return null;
  }
}

async function boot(): Promise<void> {
  const target = document.getElementById(MOUNT_ID);
  if (!target) return;

  const record = readRecord(target);
  const treeVersion = target.dataset.treeVersion ?? record?.tree_version ?? '';

  // The record names its class and race by id; the planner works in slugs. Resolving them
  // here keeps Planner identical on both mounts. A failed reference fetch still mounts, and
  // Planner renders its own "Talent data did not load" state with a retry.
  let classSlug = DEFAULT_CLASS_SLUG;
  let raceSlug: string | undefined;
  if (record) {
    try {
      const reference = await loadReference(treeVersion);
      classSlug = reference.classes.find((row) => row.id === record.class_id)?.slug ?? classSlug;
      raceSlug = reference.races.find((row) => row.id === record.race_id)?.slug;
    } catch {
      /* Planner shows the load failure and offers a retry. */
    }
  }

  mount(Planner, { target, props: { treeVersion, classSlug, raceSlug, record } });
}

if (document.readyState === 'loading') {
  document.addEventListener('DOMContentLoaded', () => void boot(), { once: true });
} else {
  void boot();
}
