package gg.foreversixty.app.home

import androidx.compose.foundation.background
import androidx.compose.foundation.border
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.heightIn
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.layout.width
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.filled.Refresh
import androidx.compose.material.icons.filled.Warning
import androidx.compose.material3.Button
import androidx.compose.material3.ButtonDefaults
import androidx.compose.material3.Icon
import androidx.compose.material3.OutlinedButton
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
import androidx.compose.ui.text.TextStyle
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.em
import androidx.compose.ui.unit.sp
import gg.foreversixty.app.designsystem.BodyFamily
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsRadius
import gg.foreversixty.app.designsystem.FsTextSize
import gg.foreversixty.app.designsystem.HeadingFamily
import java.util.Locale

/** From this font scale the rows restack for the largest text (spec §10). */
internal const val LARGE_TEXT_SCALE = 1.5f

/** The spec's type roles (§4.2), in sp so each scales with the font setting. */
object HomeType {
    val stat = TextStyle(fontFamily = HeadingFamily, fontWeight = FontWeight.ExtraBold, fontSize = FsTextSize.stat)
    val sectionTitle = TextStyle(fontFamily = HeadingFamily, fontWeight = FontWeight.ExtraBold, fontSize = FsTextSize.cardTitle)
    val pageTitle = TextStyle(fontFamily = HeadingFamily, fontWeight = FontWeight.ExtraBold, fontSize = FsTextSize.h1Phone)
    val lead = TextStyle(fontFamily = BodyFamily, fontWeight = FontWeight.SemiBold, fontSize = 17.sp, lineHeight = 24.sp)
    val body = TextStyle(fontFamily = BodyFamily, fontSize = FsTextSize.body, lineHeight = FsTextSize.bodyLineHeight)
    val bodyStrong = body.copy(fontWeight = FontWeight.SemiBold)
    val meta = TextStyle(fontFamily = BodyFamily, fontSize = FsTextSize.meta, lineHeight = 18.sp)
    val micro = TextStyle(fontFamily = BodyFamily, fontWeight = FontWeight.Bold, fontSize = FsTextSize.micro, letterSpacing = 0.14.em)
    val button = TextStyle(fontFamily = BodyFamily, fontWeight = FontWeight.Bold, fontSize = 13.sp, letterSpacing = 0.07.em)
    val tooltip = TextStyle(fontFamily = BodyFamily, fontSize = 16.sp, lineHeight = 22.sp)

    /** JetBrains Mono is not bundled yet (apps/shared/fonts), so mono figures use the system mono. */
    val mono = TextStyle(fontFamily = androidx.compose.ui.text.font.FontFamily.Monospace, fontWeight = FontWeight.SemiBold, fontSize = FsTextSize.body)
}

/** The site's uppercase section label (`micro`, .14em). */
@Composable
fun MicroLabel(text: String, modifier: Modifier = Modifier, color: Color = FsColor.muted) {
    Text(text.uppercase(Locale.ROOT), modifier.semantics { heading() }, style = HomeType.micro, color = color)
}

private val controlShape = RoundedCornerShape(FsRadius.control)

/** `PRIMARY_BUTTON_FIXED`: the gold gradient, 48 dp. One per screen at most. */
@Composable
fun PrimaryButton(title: String, onClick: () -> Unit) {
    Button(
        onClick = onClick,
        shape = controlShape,
        colors = ButtonDefaults.buttonColors(containerColor = Color.Transparent, contentColor = FsColor.bg),
        contentPadding = androidx.compose.foundation.layout.PaddingValues(0.dp),
        modifier = Modifier.heightIn(min = 48.dp),
    ) {
        Text(
            title.uppercase(Locale.ROOT),
            style = HomeType.button,
            modifier = Modifier
                .background(Brush.verticalGradient(listOf(FsColor.goldHover, FsColor.gold)), controlShape)
                .heightIn(min = 48.dp)
                .padding(horizontal = 18.dp, vertical = 15.dp),
        )
    }
}

/** `SECONDARY_BUTTON_FIXED`: 4.5% white, the gold 17% border. */
@Composable
fun SecondaryButton(title: String, onClick: () -> Unit, withRetryGlyph: Boolean = false) {
    OutlinedButton(
        onClick = onClick,
        shape = controlShape,
        border = androidx.compose.foundation.BorderStroke(1.dp, FsColor.lineCard),
        colors = ButtonDefaults.outlinedButtonColors(containerColor = Color.White.copy(alpha = 0.045f), contentColor = FsColor.strong),
        modifier = Modifier.heightIn(min = 48.dp),
    ) {
        if (withRetryGlyph) {
            Icon(Icons.Filled.Refresh, contentDescription = null, modifier = Modifier.size(16.dp))
            Spacer(Modifier.width(8.dp))
        }
        Text(title.uppercase(Locale.ROOT), style = HomeType.button)
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

/** A tag (`.pill`): 22 dp, 11 sp 700 uppercase. */
@Composable
fun Tag(text: String, color: Color) {
    Text(
        text.uppercase(Locale.ROOT),
        style = HomeType.micro.copy(letterSpacing = 0.08.em),
        color = color,
        modifier = Modifier
            .background(color.copy(alpha = 0.12f), RoundedCornerShape(FsRadius.pill))
            .border(1.dp, color.copy(alpha = 0.35f), RoundedCornerShape(FsRadius.pill))
            .padding(horizontal = 8.dp, vertical = 4.dp),
    )
}

/** The offline strip (§5.0): 36 dp, `#11151f`, the glyph, the stamp; announced politely. */
@Composable
fun OfflineStrip(text: String) {
    Row(
        modifier = Modifier
            .fillMaxWidth()
            .heightIn(min = 36.dp)
            .background(Color(0xFF11151F))
            .drawBehind { drawLine(FsColor.line, Offset(0f, size.height), Offset(size.width, size.height), 1.dp.toPx()) }
            .padding(horizontal = 16.dp, vertical = 8.dp)
            .semantics { liveRegion = LiveRegionMode.Polite },
        verticalAlignment = Alignment.CenterVertically,
        horizontalArrangement = Arrangement.spacedBy(8.dp),
    ) {
        Icon(Icons.Filled.Warning, contentDescription = null, tint = FsColor.nav, modifier = Modifier.size(16.dp))
        Text(text, style = HomeType.meta, color = FsColor.nav)
    }
}
