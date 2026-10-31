package gg.foreversixty.app.home

import android.content.res.Resources
import gg.foreversixty.app.auth.session.RelativeAge
import gg.foreversixty.app.common.OfflineStamp
import gg.foreversixty.app.phrase
import java.time.Instant
import java.time.ZoneId

/**
 * The two time phrases Home prints: the sync age (`2 h ago`) and the offline stamp (`9:41 pm`
 * today, `Thu 9:41 pm` this week, else `10 Dec, 9:41 pm`; spec §5.0).
 */
class HomeTime(private val now: Instant, private val zone: ZoneId = ZoneId.systemDefault()) {
    fun age(since: Instant, res: Resources): String = RelativeAge.of(since, now).phrase(res)

    fun stamp(at: Instant): String = OfflineStamp.format(at, now, zone)
}

/** The API's RFC 3339 timestamps; null when absent or unreadable. */
fun parseInstant(text: String?): Instant? = text?.let { runCatching { Instant.parse(it) }.getOrNull() }
