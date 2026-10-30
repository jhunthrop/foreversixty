package gg.foreversixty.app.auth.session

import gg.foreversixty.app.model.Device
import java.time.Duration
import java.time.Instant

/** One row of You › Devices (spec F1.6), ordered and dated; the view supplies the words. */
data class DeviceRow(
    val id: String,
    /** The API's name; the view shows `This phone` instead when [isThisDevice]. */
    val name: String,
    val isThisDevice: Boolean,
    val meta: Meta,
) {
    /** Every device but this one carries Revoke. */
    val canRevoke: Boolean get() = !isThisDevice

    sealed interface Meta {
        /** This device: `Signed in {date}`. */
        data class SignedIn(val at: Instant) : Meta

        /** Another device: `Last used {age}`. */
        data class LastUsed(val age: RelativeAge) : Meta

        data object None : Meta
    }
}

object DeviceList {
    /** This device first, then the others by last use, newest first. */
    fun rows(devices: List<Device>, thisDevice: String?, now: Instant): List<DeviceRow> {
        val (mine, others) = devices.partition { it.id == thisDevice }
        return mine.map { row(it, isThis = true, now) } +
            others.sortedByDescending { lastUse(it) }.map { row(it, isThis = false, now) }
    }

    private fun lastUse(device: Device): Instant =
        parse(device.lastSeenAt ?: device.createdAt) ?: Instant.EPOCH

    private fun row(device: Device, isThis: Boolean, now: Instant): DeviceRow {
        val meta = if (isThis) {
            parse(device.createdAt)?.let { DeviceRow.Meta.SignedIn(it) }
        } else {
            parse(device.lastSeenAt ?: device.createdAt)?.let { DeviceRow.Meta.LastUsed(RelativeAge.of(it, now)) }
        }
        return DeviceRow(device.id, device.name, isThis, meta ?: DeviceRow.Meta.None)
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
