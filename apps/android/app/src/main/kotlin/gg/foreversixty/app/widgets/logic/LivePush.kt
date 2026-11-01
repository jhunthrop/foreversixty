package gg.foreversixty.app.widgets.logic

/**
 * A live-surface FCM data message (W3-C14, fixtures surface.*.fcm): Android has no Live
 * Activity, so the app draws the ongoing notification itself from these strings.
 */
sealed interface LivePush {
    /** `widget.refresh`: fetch the snapshot and redraw every surface. */
    data object WidgetRefresh : LivePush

    /** `live.raid.start|update|end`: the raid countdown; its figures come from the refreshed snapshot. */
    data class Raid(val event: Event, val guildId: Long, val raidId: Long) : LivePush

    /** `live.sim.update|end`: a server run's progress. */
    data class Sim(val event: Event, val sim: SimLive) : LivePush

    enum class Event { Start, Update, End }

    companion object {
        const val WIDGET_REFRESH = "widget.refresh"
        private const val RAID_PREFIX = "live.raid."
        private const val SIM_PREFIX = "live.sim."

        /** The message's kind decides; a malformed or unknown one is null and ignored. */
        fun parse(data: Map<String, String>): LivePush? {
            val kind = data["kind"] ?: return null
            return when {
                kind == WIDGET_REFRESH -> WidgetRefresh
                kind.startsWith(RAID_PREFIX) -> {
                    val event = event(kind.removePrefix(RAID_PREFIX)) ?: return null
                    Raid(event, data["guild_id"]?.toLongOrNull() ?: return null, data["raid_id"]?.toLongOrNull() ?: return null)
                }
                kind.startsWith(SIM_PREFIX) -> {
                    val event = event(kind.removePrefix(SIM_PREFIX)) ?: return null
                    Sim(event, SimLive.parse(data) ?: return null)
                }
                else -> null
            }
        }

        private fun event(word: String): Event? = when (word) {
            "start" -> Event.Start
            "update" -> Event.Update
            "end" -> Event.End
            else -> null
        }
    }
}

/** A server run's live state (SimActivityAttributes and SimActivityState). [total] is null for a plain run. */
data class SimLive(
    val simId: String,
    val title: String,
    val state: State,
    val done: Int,
    val total: Int?,
    val etaSeconds: Int?,
    val dps: Double?,
) {
    enum class State(val wire: String) { Running("running"), Finished("finished"), Failed("failed") }

    companion object {
        fun parse(data: Map<String, String>): SimLive? {
            val id = data["sim_id"]?.takeIf { it.isNotEmpty() } ?: return null
            val state = State.entries.firstOrNull { it.wire == data["state"] } ?: return null
            return SimLive(
                simId = id, title = data["title"].orEmpty(), state = state, done = data["done"]?.toIntOrNull() ?: 0,
                total = data["total"]?.toIntOrNull(), etaSeconds = data["eta_sec"]?.toIntOrNull(), dps = data["dps"]?.toDoubleOrNull(),
            )
        }
    }
}
