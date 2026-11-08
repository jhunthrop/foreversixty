package gg.foreversixty.app.auth.session

import gg.foreversixty.app.model.Device
import java.time.Duration
import java.time.Instant

/** One row of Devices (mock `devices`): signed in on a date, last used an age ago; the view supplies the words. */
data class DeviceRow(
    val id: String,
    /** The API's name (`Pixel 9`); the view adds `This device` when [isThisDevice]. */
    val name: String,
    val isThisDevice: Boolean,
    /** When it signed in; null when the API sent no date. */
    val signedIn: Instant?,
    /** How long since it was last used; this device is used now. */
    val lastUsed: RelativeAge?,
    /** The desktop companion. */
    val isCompanion: Boolean = false,
) {
    /** Every device but this one carries Revoke. */
    val canRevoke: Boolean get() = !isThisDevice
}

object DeviceList {
    private const val COMPANION = "companion"

    /** This device first, then the others by last use, newest first. */
    fun rows(devices: List<Device>, thisDevice: String?, now: Instant): List<DeviceRow> {
        val (mine, others) = devices.partition { it.id == thisDevice }
        return mine.map { row(it, isThis = true, now) } +
            others.sortedByDescending { lastUse(it) }.map { row(it, isThis = false, now) }
    }

    private fun lastUse(device: Device): Instant =
        parse(device.lastSeenAt ?: device.createdAt) ?: Instant.EPOCH

    private fun row(device: Device, isThis: Boolean, now: Instant): DeviceRow {
        val lastUsed = if (isThis) RelativeAge.Now else parse(device.lastSeenAt ?: device.createdAt)?.let { RelativeAge.of(it, now) }
        return DeviceRow(device.id, device.name, isThis, parse(device.createdAt), lastUsed, isCompanion = device.kind == COMPANION)
    }

    private fun parse(stamp: String?): Instant? = stamp?.let { runCatching { Instant.parse(it) }.getOrNull() }
}

/** How long ago, in the one unit the meta lines use (`2 h ago`). */
sealed interface RelativeAge {
    data object Now : RelativeAge
    data class Minutes(val n: Int) : RelativeAge
    data class Hours(val n: Int) : RelativeAge
    data class Days(val n: Int) : RelativeAge

    companion object {
        fun of(at: Instant, now: Instant): RelativeAge {
            val seconds = Duration.between(at, now).seconds.coerceAtLeast(0)
            return when {
                seconds < MINUTE -> Now
                seconds < HOUR -> Minutes((seconds / MINUTE).toInt())
                seconds < DAY -> Hours((seconds / HOUR).toInt())
                else -> Days((seconds / DAY).toInt())
            }
        }

        private const val MINUTE = 60L
        private const val HOUR = 3_600L
        private const val DAY = 86_400L
    }
}
