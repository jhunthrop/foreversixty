package gg.foreversixty.app.common

import androidx.compose.foundation.background
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.heightIn
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.material3.Icon
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.drawBehind
import androidx.compose.ui.geometry.Offset
import androidx.compose.ui.res.painterResource
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.semantics.LiveRegionMode
import androidx.compose.ui.semantics.liveRegion
import androidx.compose.ui.semantics.semantics
import androidx.compose.ui.unit.dp
import gg.foreversixty.app.R
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.GamePalette
import java.time.Instant
import java.time.ZoneId
import java.time.format.DateTimeFormatter
import java.time.temporal.ChronoUnit
import java.util.Locale

/**
 * When the shown answer was fetched, in the strip's words (spec §5.0): `9:41 pm` today,
 * `Thu 9:41 pm` this week, else `10 Dec, 9:41 pm`.
 */
object OfflineStamp {
    private const val DAYS_IN_WEEK = 7
    private val time = DateTimeFormatter.ofPattern("h:mm", Locale.US)
    private val meridiem = DateTimeFormatter.ofPattern("a", Locale.US)
    private val weekday = DateTimeFormatter.ofPattern("EEE", Locale.US)
    private val date = DateTimeFormatter.ofPattern("d MMM", Locale.US)

    fun format(savedAtMillis: Long, nowMillis: Long = System.currentTimeMillis(), zone: ZoneId = ZoneId.systemDefault()): String {
        val saved = Instant.ofEpochMilli(savedAtMillis).atZone(zone)
        val now = Instant.ofEpochMilli(nowMillis).atZone(zone)
        val days = ChronoUnit.DAYS.between(saved.toLocalDate(), now.toLocalDate())
        val clock = "${time.format(saved)} ${meridiem.format(saved).lowercase(Locale.US)}"
        return when (days) {
            0L -> clock
            in 1 until DAYS_IN_WEEK -> "${weekday.format(saved)} $clock"
            else -> "${date.format(saved)}, $clock"
        }
    }
}

/**
 * The offline strip (spec §5.0): 36 tall under the bar, a glyph and the stamp. Shared by every
 * screen that draws a saved answer; TalkBack reads it once as a polite live region.
 */
@Composable
fun OfflineStrip(savedAtMillis: Long, nowMillis: Long = System.currentTimeMillis()) {
    val line = FsColor.line
    Box(
        Modifier
            .fillMaxWidth()
            .heightIn(min = 36.dp)
            .background(GamePalette.offlineGround)
            .drawBehind { drawLine(line, Offset(0f, size.height), Offset(size.width, size.height), strokeWidth = 1.dp.toPx()) }
            .semantics(mergeDescendants = true) { liveRegion = LiveRegionMode.Polite },
        contentAlignment = Alignment.CenterStart,
    ) {
        Row(
            Modifier.padding(horizontal = 16.dp, vertical = 8.dp),
            horizontalArrangement = Arrangement.spacedBy(8.dp),
            verticalAlignment = Alignment.CenterVertically,
        ) {
            Icon(painterResource(R.drawable.ic_wifi_off), contentDescription = null, tint = FsColor.nav, modifier = Modifier.size(16.dp))
            Text(
                stringResource(R.string.offline_strip, OfflineStamp.format(savedAtMillis, nowMillis)),
                style = MaterialTheme.typography.labelMedium,
                color = FsColor.nav,
            )
        }
    }
}
