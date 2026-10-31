package gg.foreversixty.app.account

import androidx.annotation.DrawableRes
import androidx.compose.foundation.Image
import androidx.compose.foundation.background
import androidx.compose.foundation.border
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.shape.CircleShape
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.outlined.CheckCircle
import androidx.compose.material.icons.outlined.Info
import androidx.compose.material3.Icon
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.clip
import androidx.compose.ui.draw.shadow
import androidx.compose.ui.graphics.Brush
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.layout.ContentScale
import androidx.compose.ui.res.painterResource
import androidx.compose.ui.semantics.LiveRegionMode
import androidx.compose.ui.semantics.liveRegion
import androidx.compose.ui.semantics.semantics
import androidx.compose.ui.unit.Dp
import androidx.compose.ui.unit.dp
import gg.foreversixty.app.R
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsRadius
import gg.foreversixty.app.designsystem.FsSpacing
import gg.foreversixty.app.designsystem.FsType
import gg.foreversixty.app.designsystem.classColor

/** `@drawable/crest_<class>`, copied from web/public by the forever.brand-art plugin. */
@DrawableRes
private fun crestOf(className: String?): Int? = when (className?.lowercase()) {
    "warrior" -> R.drawable.crest_warrior
    "paladin" -> R.drawable.crest_paladin
    "hunter" -> R.drawable.crest_hunter
    "rogue" -> R.drawable.crest_rogue
    "priest" -> R.drawable.crest_priest
    "shaman" -> R.drawable.crest_shaman
    "mage" -> R.drawable.crest_mage
    "warlock" -> R.drawable.crest_warlock
    "druid" -> R.drawable.crest_druid
    else -> null
}

/** The circular ringed HD crest (spec §4.3): never a square class icon. */
@Composable
fun ClassCrest(className: String?, size: Dp, modifier: Modifier = Modifier) {
    val color = classColor(className)
    Box(
        modifier
            .size(size)
            .shadow(7.dp, CircleShape, ambientColor = color, spotColor = color)
            .clip(CircleShape)
            .background(FsColor.raised)
            .border(2.dp, color.copy(alpha = RING_ALPHA), CircleShape),
    ) {
        crestOf(className)?.let {
            Image(painterResource(it), contentDescription = null, contentScale = ContentScale.Crop, modifier = Modifier.fillMaxSize())
        }
    }
}

private const val RING_ALPHA = 0.6f

/** The LX seal. */
@Composable
fun Seal(size: Dp) {
    Image(painterResource(R.drawable.brand_seal), contentDescription = null, modifier = Modifier.size(size))
}

/** The hero band ground (`--surface-band`): the ember radial over the warm dark. */
fun Modifier.bandBackground(): Modifier = background(FsColor.bg).background(
    Brush.radialGradient(listOf(FsColor.ember.copy(alpha = 0.28f), Color.Transparent), radius = 1_400f, center = androidx.compose.ui.geometry.Offset(600f, 0f)),
)

/** The toast at the top of F1.1: failures in wipe red, a finished deletion neutral. */
@Composable
fun NoticeToast(text: String, modifier: Modifier = Modifier, isFailure: Boolean = true) {
    val border = if (isFailure) FsColor.wipe.copy(alpha = 0.45f) else FsColor.line
    Row(
        modifier
            .fillMaxWidth()
            .background(FsColor.raised, RoundedCornerShape(FsRadius.panel))
            .border(1.dp, border, RoundedCornerShape(FsRadius.panel))
            .padding(FsSpacing.space4)
            .semantics { liveRegion = LiveRegionMode.Assertive },
        horizontalArrangement = Arrangement.spacedBy(FsSpacing.space3),
    ) {
        Icon(
            if (isFailure) Icons.Outlined.Info else Icons.Outlined.CheckCircle,
            contentDescription = null,
            tint = if (isFailure) FsColor.wipe else FsColor.gold,
        )
        Text(text, style = FsType.body, color = FsColor.strong)
    }
}
