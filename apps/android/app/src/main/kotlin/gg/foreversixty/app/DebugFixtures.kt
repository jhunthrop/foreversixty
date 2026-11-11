package gg.foreversixty.app

import gg.foreversixty.app.network.FixtureVariant

/**
 * Debug builds only (round 4 a2): the fixture states the review walk asks for, so the emulator can
 * show a surface the default fixtures do not reach (the officer's Raid tab). Set by
 * `adb shell am start ... --es fs.debug.fixtures guild.raid-badge=officer`; the fixture engine reads
 * it on every request, and [latencyMillis] (`--el fs.debug.latency 1500`). Live builds never read either.
 */
object DebugFixtures {
    const val EXTRA = "fs.debug.fixtures"

    /** `--el fs.debug.latency 1500`: every fixture answer waits this long (the in-flight frames). */
    const val LATENCY_EXTRA = "fs.debug.latency"

    @Volatile
    var latencyMillis: Long = 0L

    @Volatile
    var variants: Map<String, FixtureVariant> = emptyMap()

    /** `stem=variant,stem=variant`: each a route stem and its named success variant. */
    fun parse(spec: String): Map<String, FixtureVariant> = spec.split(',')
        .map { it.trim() }
        .filter { it.isNotEmpty() }
        .associate { pair ->
            val (stem, variant) = pair.split('=', limit = 2).map { it.trim() }.also {
                require(it.size == 2 && it.all(String::isNotEmpty)) { "fs.debug.fixtures takes stem=variant pairs, not $pair" }
            }
            stem to FixtureVariant.Named(variant)
        }
}
