/// <reference types="astro/client" />

// Type augmentation for the custom `client:interaction` directive registered in
// astro.config.mjs (see src/directives/interaction.ts for the runtime behavior).
declare namespace Astro {
  interface ClientDirectives {
    'client:interaction'?: boolean;
  }
}

interface ImportMetaEnv {
  /** Origin of the Go API. Defaults to https://api.foreversixty.gg in lib/planner/config.ts. */
  readonly PUBLIC_API_BASE_URL?: string;
  /** 'fake' | 'wasm'. Defaults to 'fake' in lib/sim/engine.ts until sim.wasm exists. */
  readonly PUBLIC_SIM_ENGINE?: string;
}

interface ImportMeta {
  readonly env: ImportMetaEnv;
}
