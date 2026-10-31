package gg.foreversixty.app.designsystem

import androidx.compose.foundation.background
import androidx.compose.foundation.border
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.heightIn
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.filled.Warning
import androidx.compose.material3.Icon
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.drawBehind
import androidx.compose.ui.geometry.Offset
import androidx.compose.ui.graphics.Brush
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.semantics.LiveRegionMode
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

/** A tag (`.pill`): 22 dp, 11 sp 700 uppercase, tinted ground and border. */
@Composable
fun Tag(text: String, color: Color) {
    Text(
        text.uppercase(Locale.ROOT),
        style = FsType.micro.copy(letterSpacing = 0.08.em),
        color = color,
        modifier = Modifier
            .background(color.copy(alpha = 0.12f), RoundedCornerShape(FsRadius.pill))
            .border(1.dp, color.copy(alpha = 0.35f), RoundedCornerShape(FsRadius.pill))
            .padding(horizontal = FsSpacing.space2, vertical = 4.dp),
    )
}

/** The offline strip (§5.0): 36 dp, `#11151f`, the glyph, then [text]; announced politely. */
@Composable
fun OfflineStrip(text: String) {
    Row(
        modifier = Modifier
            .fillMaxWidth()
            .heightIn(min = 36.dp)
            .background(Color(0xFF11151F))
            .drawBehind { drawLine(FsColor.line, Offset(0f, size.height), Offset(size.width, size.height), 1.dp.toPx()) }
            .padding(horizontal = FsSpacing.space4, vertical = FsSpacing.space2)
            .semantics { liveRegion = LiveRegionMode.Polite },
        verticalAlignment = Alignment.CenterVertically,
        horizontalArrangement = Arrangement.spacedBy(FsSpacing.space2),
    ) {
        Icon(Icons.Filled.Warning, contentDescription = null, tint = FsColor.nav, modifier = Modifier.size(16.dp))
        Text(text, style = FsType.meta, color = FsColor.nav)
    }
}

/** The primary card (`--surface-card`): the warm gradient and its gold border. */
fun Modifier.primaryCard(): Modifier = this
    .background(Brush.verticalGradient(listOf(Color(0xFF16130E), FsColor.raised)), RoundedCornerShape(FsRadius.panel))
    .border(1.dp, FsColor.lineCard, RoundedCornerShape(FsRadius.panel))

/** A quiet list card: raised ground, soft line. */
fun Modifier.quietCard(): Modifier = this
    .background(FsColor.raised, RoundedCornerShape(FsRadius.panel))
    .border(1.dp, FsColor.lineSoft, RoundedCornerShape(FsRadius.panel))

/** The hero band (`--surface-band`): the ember radial over `#14100b`, a soft bottom line. */
fun Modifier.heroBand(): Modifier = this
    .background(Color(0xFF14100B))
    .background(Brush.radialGradient(listOf(FsColor.ember.copy(alpha = 0.22f), Color.Transparent), radius = 900f, center = Offset(400f, 0f)))
    .drawBehind {
        drawLine(FsColor.lineSoft, Offset(0f, size.height), Offset(size.width, size.height), 1.dp.toPx())
    }
