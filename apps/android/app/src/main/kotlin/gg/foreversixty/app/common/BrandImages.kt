package gg.foreversixty.app.common

import androidx.annotation.DrawableRes
import androidx.compose.foundation.Image
import androidx.compose.foundation.background
import androidx.compose.foundation.border
import androidx.compose.foundation.layout.Box
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
import gg.foreversixty.app.designsystem.GamePalette

/**
 * The bundled brand images from web/public/icons/hd (the forever.bundled-data plugin copies them
 * in as `crest_<class>` and `faction_<faction>`), so a crest or a faction mark never waits.
 */
object BrandImages {
    @DrawableRes
    fun crest(classSlug: String): Int? = when (classSlug) {
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
    fun faction(slug: String): Int? = when (slug) {
        "horde" -> R.drawable.faction_horde
        "alliance" -> R.drawable.faction_alliance
        else -> null
    }
}

/**
 * The circular ringed HD crest (spec §4.3, "one crest language"): a 2 dp ring at 60% of the class
 * colour and a soft glow.
 */
@Composable
fun ClassCrest(classSlug: String, size: Dp = 44.dp) {
    val color = GamePalette.classColor(classSlug)
    Box(
        Modifier
            .size(size)
            .shadow(8.dp, CircleShape, ambientColor = color, spotColor = color)
            .clip(CircleShape)
            .background(FsColor.cardTop)
            .border(2.dp, color.copy(alpha = 0.6f), CircleShape),
    ) {
        BrandImages.crest(classSlug)?.let {
            Image(painterResource(it), contentDescription = null, contentScale = ContentScale.Crop, modifier = Modifier.size(size))
        }
    }
}

/** The Horde or Alliance mark. */
@Composable
fun FactionMark(faction: String, size: Dp = 24.dp) {
    val image = BrandImages.faction(faction)
    if (image != null) {
        Image(painterResource(image), contentDescription = null, modifier = Modifier.size(size))
    } else {
        Box(Modifier.size(size).background(GamePalette.faction(faction), CircleShape))
    }
}
