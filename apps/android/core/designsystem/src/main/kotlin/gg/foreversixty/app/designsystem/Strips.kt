package gg.foreversixty.app.designsystem

import androidx.annotation.DrawableRes
import androidx.compose.foundation.background
import androidx.compose.foundation.border
import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.heightIn
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.layout.wrapContentHeight
import androidx.compose.material3.Icon
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.drawBehind
import androidx.compose.ui.geometry.Offset
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.graphics.Shape
import androidx.compose.ui.res.painterResource
import androidx.compose.ui.semantics.LiveRegionMode
import androidx.compose.ui.semantics.Role
import androidx.compose.ui.semantics.liveRegion
import androidx.compose.ui.semantics.semantics
import androidx.compose.ui.unit.dp

/**
 * The one strip (v1 §5.0, W3 §5.7, board F15.5), in three states: offline, stale and new data. 36 dp
 * at least, a glyph, the sentence and, for stale and new data, a trailing gold text button with a
 * 48 dp target. Announced politely once when it appears; the action is a button to TalkBack.
 */

/** Offline: `#11151f`, the wifi-off glyph, then [text]. Given a [shape] (inside a list) it is a card. */
@Composable
fun OfflineStrip(text: String, shape: Shape? = null) {
    StripRow(GameColors.offlineGround, shape, R.drawable.ic_wifi_off, text, FsColor.nav)
}

/** Stale: the offline ground with the cloud glyph, [text] and a trailing gold [action] (`Retry`). */
@Composable
fun StaleStrip(text: String, action: String, onAction: () -> Unit, shape: Shape? = null) {
    StripRow(GameColors.offlineGround, shape, R.drawable.ic_cloud, text, FsColor.nav, StripAction(action, onAction))
}

/**
 * New data: a gold 10% ground, [text] and a trailing gold [action] (`Update`). It only ever appears
 * above content; nothing under it changes until [onAction]. [failed] draws the sentence in wipe.
 */
@Composable
fun NewDataStrip(text: String, action: String, onAction: () -> Unit, failed: Boolean = false) {
    StripRow(
        FsColor.gold.copy(alpha = NEW_DATA_GROUND_ALPHA),
        null,
        R.drawable.ic_download,
        text,
        if (failed) FsColor.wipe else FsColor.text,
        StripAction(action, onAction),
    )
}

private class StripAction(val label: String, val onClick: () -> Unit)

@Composable
private fun StripRow(ground: Color, shape: Shape?, @DrawableRes icon: Int, text: String, ink: Color, action: StripAction? = null) {
    // Full width it is a band with a bottom rule; given a [shape] it is a card, its line all the way round.
    val surface = if (shape == null) {
        Modifier
            .background(ground)
            .drawBehind { drawLine(FsColor.line, Offset(0f, size.height), Offset(size.width, size.height), 1.dp.toPx()) }
    } else {
        Modifier.background(ground, shape).border(1.dp, FsColor.line, shape)
    }
    Row(
        modifier = Modifier
            .fillMaxWidth()
            .heightIn(min = 36.dp)
            .then(surface)
            .padding(start = FsSpacing.space4, end = if (action == null) FsSpacing.space4 else 0.dp)
            .semantics { liveRegion = LiveRegionMode.Polite },
        verticalAlignment = Alignment.CenterVertically,
        horizontalArrangement = Arrangement.spacedBy(FsSpacing.space2),
    ) {
        Icon(painterResource(icon), contentDescription = null, tint = FsColor.nav, modifier = Modifier.size(16.dp))
        Text(text, style = FsType.meta, color = ink, modifier = Modifier.weight(1f).padding(vertical = FsSpacing.space2))
        action?.let {
            Text(
                it.label,
                style = FsType.bodyStrong,
                color = FsColor.gold,
                modifier = Modifier
                    .heightIn(min = 48.dp)
                    .clickable(role = Role.Button, onClick = it.onClick)
                    .padding(horizontal = FsSpacing.space4)
                    .wrapContentHeight(Alignment.CenterVertically),
            )
        }
    }
}

private const val NEW_DATA_GROUND_ALPHA = 0.10f
