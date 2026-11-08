package gg.foreversixty.app.push

import gg.foreversixty.app.navigation.DeepLink
import java.net.URI

/**
 * One FCM message as the app reads it (app-v1.md "Push payloads"). FCM data values are strings.
 * Title and body come from the data (the data-only sends Android needs to draw its own actions),
 * else from the notification block of an older send.
 */
data class PushMessage(
    val kind: String,
    val group: PushGroup,
    val title: String,
    val body: String,
    val link: String?,
    val category: String?,
    val collapse: String?,
    val data: Map<String, String>,
) {
    /** A raid message's guild and raid, read from the data. */
    val raid: Pair<Long, Long>?
        get() {
            val guild = data[KEY_GUILD]?.toLongOrNull() ?: return null
            val raid = data[KEY_RAID]?.toLongOrNull() ?: return null
            return guild to raid
        }

    /** The app link a tap opens (spec F4's landing table), or null for Home. */
    fun landing(): String? = PushLanding.of(this)?.url()

    companion object {
        const val KEY_KIND = "kind"
        const val KEY_GROUP = "group"
        const val KEY_TITLE = "title"
        const val KEY_BODY = "body"
        const val KEY_LINK = "link"
        const val KEY_CATEGORY = "category"
        const val KEY_COLLAPSE = "collapse"
        const val KEY_GUILD = "guild_id"
        const val KEY_RAID = "raid_id"

        /** Null when the message has nothing to show. */
        fun parse(data: Map<String, String>, notificationTitle: String?, notificationBody: String?, collapse: String?): PushMessage? {
            val title = data[KEY_TITLE]?.takeIf { it.isNotBlank() } ?: notificationTitle?.takeIf { it.isNotBlank() } ?: return null
            val kind = data[KEY_KIND].orEmpty()
            return PushMessage(
                kind = kind,
                group = PushGroup.of(data[KEY_GROUP], kind),
                title = title,
                body = data[KEY_BODY] ?: notificationBody.orEmpty(),
                link = data[KEY_LINK]?.takeIf { it.isNotBlank() },
                category = data[KEY_CATEGORY]?.takeIf { it.isNotBlank() },
                collapse = data[KEY_COLLAPSE]?.takeIf { it.isNotBlank() } ?: collapse ?: data[KEY_RAID]?.let { "$kind-$it" },
                data = data,
            )
        }
    }
}

/**
 * Where a tap lands (spec 2026-11-08 F4): the raid page for every raid and officer event, the sim
 * result, the report; an app link the server sent opens as itself. Gear, the item sheet, Tiers and the
 * guild sims ladder land with the tabs that hold them (the app's navigation lanes); until then they open
 * Home.
 */
object PushLanding {
    fun of(message: PushMessage): DeepLink? {
        message.link?.let(DeepLink::parse)?.let { return it }
        message.raid?.let { (guild, raid) -> return DeepLink.Raid(guild, raid) }
        val path = message.link?.let { runCatching { URI(it).path }.getOrNull() }.orEmpty()
        return when (message.kind) {
            "sim.finished" -> lastSegment(path, "sim")?.let(DeepLink::Sim)
            "report.ready", "coach.ready" -> lastSegment(path, "reports")?.let(DeepLink::Report)
            else -> null
        }
    }

    /** `/sim/abc` with prefix `sim` is `abc`. */
    private fun lastSegment(path: String, prefix: String): String? {
        val parts = path.split('/').filter { it.isNotEmpty() }
        return parts.takeIf { it.size == 2 && it[0] == prefix }?.get(1)
    }
}
