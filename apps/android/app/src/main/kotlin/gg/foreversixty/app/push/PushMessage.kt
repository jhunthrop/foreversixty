package gg.foreversixty.app.push

import gg.foreversixty.app.navigation.DeepLink
import gg.foreversixty.app.navigation.RAID_VIEW_UNANSWERED

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
 * Where a tap lands (spec 2026-11-08 F4): the raid page for every raid and officer event (its ids are
 * in the data), else the link the server sent, an app link or a site page MainActivity routes.
 */
object PushLanding {
    private val UNANSWERED_KINDS = setOf("raid.short", "roster.pending", "raid.late_change")

    fun of(message: PushMessage): DeepLink? {
        val linked = message.link?.let(DeepLink::parse)
        message.raid?.let { (guild, raid) ->
            // An officer's raid push opens the page filtered to Not answered (spec 2026-11-08 F3).
            val view = RAID_VIEW_UNANSWERED.takeIf { message.kind in UNANSWERED_KINDS }
            // The ids win; the server's app link still says what to press (`state=`, `spec=ask`).
            val asks = linked as? DeepLink.Raid
            return DeepLink.Raid(guild, raid, spec = asks?.spec, state = asks?.state, view = view)
        }
        return linked
    }
}
