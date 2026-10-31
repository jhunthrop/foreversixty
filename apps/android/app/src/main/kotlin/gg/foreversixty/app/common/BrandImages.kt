package gg.foreversixty.app.common

import androidx.annotation.DrawableRes
import androidx.compose.foundation.Image
import androidx.compose.foundation.background
import androidx.compose.foundation.border
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.shape.CircleShape
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.clip
import androidx.compose.ui.draw.shadow
import androidx.compose.ui.layout.ContentScale
import androidx.compose.ui.res.painterResource
import androidx.compose.ui.unit.Dp
import androidx.compose.ui.unit.dp
import gg.foreversixty.app.R
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.GameColors
import gg.foreversixty.app.designsystem.classColor

// The bundled brand images, copied from web/public by the forever.bundled-data plugin as
// `crest_<class>`, `faction_<faction>` and `brand_seal`, so none of them ever waits on the network.

/** The bundled crest of a class; null for an unknown one. The widgets draw it too. */
@DrawableRes
internal fun crestOf(className: String?): Int? = when (className?.lowercase()) {
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

@DrawableRes
private fun factionOf(slug: String): Int? = when (slug) {
    "horde" -> R.drawable.faction_horde
    "alliance" -> R.drawable.faction_alliance
    else -> null
}

/** The circular ringed HD crest (spec §4.3, one crest language): never a square class icon. */
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

/** The Horde or Alliance mark; a dot in the faction colour for an unknown slug. */
@Composable
fun FactionMark(faction: String, size: Dp) {
    val image = factionOf(faction)
    if (image != null) {
        Image(painterResource(image), contentDescription = null, modifier = Modifier.size(size))
    } else {
        Box(Modifier.size(size).background(GameColors.faction(faction), CircleShape))
    }
}

/** The LX seal. */
@Composable
fun Seal(size: Dp) {
    Image(painterResource(R.drawable.brand_seal), contentDescription = null, modifier = Modifier.size(size))
}
