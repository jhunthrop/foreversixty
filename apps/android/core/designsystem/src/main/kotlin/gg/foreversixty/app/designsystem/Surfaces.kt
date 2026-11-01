package gg.foreversixty.app.designsystem

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
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.filled.Check
import androidx.compose.material3.Icon
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.clip
import androidx.compose.ui.draw.drawBehind
import androidx.compose.ui.geometry.Offset
import androidx.compose.ui.graphics.Brush
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.graphics.Shape
import androidx.compose.ui.res.painterResource
import androidx.compose.ui.semantics.LiveRegionMode
import androidx.compose.ui.semantics.Role
import androidx.compose.ui.semantics.heading
import androidx.compose.ui.semantics.liveRegion
import androidx.compose.ui.semantics.semantics
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.em
import java.util.Locale

/** The site's uppercase section label (`micro`, .14em), read as a heading. */
@Composable
fun MicroLabel(text: String, modifier: Modifier = Modifier, color: Color = FsColor.muted) {
    Text(text.uppercase(Locale.ROOT), modifier.semantics { heading() }, style = FsType.micro, color = color)
}

/** A tag (`.pill`): 22 dp, 11 sp 700 uppercase, tinted ground and border; [checked] leads with a tick. */
@Composable
fun Tag(text: String, color: Color, modifier: Modifier = Modifier, checked: Boolean = false) {
    val shape = RoundedCornerShape(FsRadius.pill)
    Row(
        modifier
            .heightIn(min = 22.dp)
            .background(color.copy(alpha = 0.12f), shape)
            .border(1.dp, color.copy(alpha = 0.35f), shape)
            .padding(horizontal = FsSpacing.space2, vertical = 4.dp),
        horizontalArrangement = Arrangement.spacedBy(4.dp),
        verticalAlignment = Alignment.CenterVertically,
    ) {
        if (checked) Icon(Icons.Filled.Check, contentDescription = null, tint = color, modifier = Modifier.size(11.dp))
        Text(text.uppercase(Locale.ROOT), style = FsType.micro.copy(letterSpacing = 0.08.em), color = color)
    }
}

/** The offline strip (§5.0): 36 dp, `#11151f`, the wifi-off glyph, then [text]; announced politely. */
@Composable
fun OfflineStrip(text: String, shape: Shape? = null) {
    // Full width it is a band with a bottom rule; given a [shape] (inside a list) it is a card,
    // its line all the way round so every corner rounds.
    val ground = if (shape == null) {
        Modifier
            .background(GameColors.offlineGround)
            .drawBehind { drawLine(FsColor.line, Offset(0f, size.height), Offset(size.width, size.height), 1.dp.toPx()) }
    } else {
        Modifier.background(GameColors.offlineGround, shape).border(1.dp, FsColor.line, shape)
    }
    Row(
        modifier = Modifier
            .fillMaxWidth()
            .heightIn(min = 36.dp)
            .then(ground)
            .padding(horizontal = FsSpacing.space4, vertical = FsSpacing.space2)
            .semantics { liveRegion = LiveRegionMode.Polite },
        verticalAlignment = Alignment.CenterVertically,
        horizontalArrangement = Arrangement.spacedBy(FsSpacing.space2),
    ) {
        Icon(painterResource(R.drawable.ic_wifi_off), contentDescription = null, tint = FsColor.nav, modifier = Modifier.size(16.dp))
        Text(text, style = FsType.meta, color = FsColor.nav)
    }
}

/**
 * The new-data strip (W3 §5.7): 36 dp on a gold 10% ground, [text] and a trailing gold [action]
 * (a 48 dp target). It only ever appears above content; nothing under it changes until [onAction].
 * Lane F15 extends this same component with the stale state.
 */
@Composable
fun NewDataStrip(text: String, action: String, onAction: () -> Unit, failed: Boolean = false) {
    Row(
        modifier = Modifier
            .fillMaxWidth()
            .heightIn(min = 36.dp)
            .background(FsColor.gold.copy(alpha = NEW_DATA_GROUND_ALPHA))
            .drawBehind { drawLine(FsColor.line, Offset(0f, size.height), Offset(size.width, size.height), 1.dp.toPx()) }
            .padding(start = FsSpacing.space4)
            .semantics { liveRegion = LiveRegionMode.Polite },
        verticalAlignment = Alignment.CenterVertically,
        horizontalArrangement = Arrangement.spacedBy(FsSpacing.space2),
    ) {
        Text(text, style = FsType.meta, color = if (failed) FsColor.wipe else FsColor.text, modifier = Modifier.weight(1f).padding(vertical = FsSpacing.space2))
        Text(
            action,
            style = FsType.bodyStrong,
            color = FsColor.gold,
            modifier = Modifier
                .heightIn(min = 48.dp)
                .clickable(role = Role.Button, onClick = onAction)
                .padding(horizontal = FsSpacing.space4)
                .wrapContentHeight(Alignment.CenterVertically),
        )
    }
}

private const val NEW_DATA_GROUND_ALPHA = 0.10f

/** The primary card (`--surface-card`): the warm gradient and its gold border. */
fun Modifier.primaryCard(): Modifier = this
    .clip(RoundedCornerShape(FsRadius.panel))
    .background(Brush.verticalGradient(listOf(Color(0xFF16130E), FsColor.raised)), RoundedCornerShape(FsRadius.panel))
    .border(1.dp, FsColor.lineCard, RoundedCornerShape(FsRadius.panel))

/** A quiet list card: raised ground, soft line. */
fun Modifier.quietCard(): Modifier = this
    .clip(RoundedCornerShape(FsRadius.panel))
    .background(FsColor.raised, RoundedCornerShape(FsRadius.panel))
    .border(1.dp, FsColor.lineSoft, RoundedCornerShape(FsRadius.panel))

/**
 * The band ground (`--surface-band`, spec §4.1): the ember radial from the top centre over
 * `#14100b`, with the soft bottom line when [bottomLine]. Every band uses it: Home's hero bands
 * (with the line) and the signed-out screens' full-bleed ground (without). iOS: `surfaceBand`.
 */
fun Modifier.surfaceBand(bottomLine: Boolean = true): Modifier = drawBehind {
    drawRect(BAND_GROUND)
    drawRect(
        Brush.radialGradient(
            listOf(FsColor.ember.copy(alpha = BAND_GLOW_ALPHA), Color.Transparent),
            center = Offset(size.width / 2, 0f),
            radius = BAND_GLOW_RADIUS.toPx(),
        ),
    )
    if (bottomLine) drawLine(FsColor.lineSoft, Offset(0f, size.height), Offset(size.width, size.height), 1.dp.toPx())
}

private val BAND_GROUND = Color(0xFF14100B)
private const val BAND_GLOW_ALPHA = 0.22f
private val BAND_GLOW_RADIUS = 320.dp

/** A chip or pill's ground: [fill] and a 1 dp [line] border in [shape]. */
fun Modifier.pillSurface(fill: Color, line: Color, shape: Shape = RoundedCornerShape(percent = 50)): Modifier =
    background(fill, shape).border(1.dp, line, shape)
