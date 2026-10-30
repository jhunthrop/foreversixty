package gg.foreversixty.app.items

import androidx.compose.foundation.background
import androidx.compose.foundation.border
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.runtime.Composable
import androidx.compose.runtime.staticCompositionLocalOf
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.clip
import androidx.compose.ui.draw.shadow
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.layout.ContentScale
import androidx.compose.ui.unit.Dp
import androidx.compose.ui.unit.dp
import coil3.compose.AsyncImage
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsRadius
import gg.foreversixty.app.designsystem.GamePalette
import gg.foreversixty.app.network.StaticDataClient

/**
 * Where game icons come from: `/data/<build>/icons/<icon>.webp` on the site (contract "Static
 * files"). Null in screenshot tests and previews, so the tile draws without a request.
 */
data class IconSource(val build: String, val siteUrl: String = StaticDataClient.SITE_URL) {
    fun url(icon: String): String = "${siteUrl.trimEnd('/')}/data/$build/icons/$icon.webp"
}

val LocalIconSource = staticCompositionLocalOf<IconSource?> { null }

/** A square game icon from the CDN; Coil fades it in and keeps it on disk (spec §11). */
@Composable
fun RemoteIcon(icon: String?, source: IconSource?, modifier: Modifier = Modifier) {
    if (icon.isNullOrEmpty() || source == null) return
    AsyncImage(model = source.url(icon), contentDescription = null, contentScale = ContentScale.Crop, modifier = modifier)
}

/**
 * The quality-ringed item tile (spec §4.3): a 1.5 dp quality border, a 1 dp inner black line and
 * a soft quality glow. It draws at once at its real size; the icon fades in.
 */
@Composable
fun ItemIcon(icon: String?, quality: Int, size: Dp = 44.dp, source: IconSource? = LocalIconSource.current) {
    val shape = RoundedCornerShape(FsRadius.pill)
    val border = GamePalette.qualityBorder(quality)
    Box(
        Modifier
            .size(size)
            .shadow(5.dp, shape, ambientColor = border.copy(alpha = 0.22f), spotColor = border.copy(alpha = 0.22f))
            .clip(shape)
            .background(FsColor.cardTop),
    ) {
        RemoteIcon(icon, source, Modifier.fillMaxSize())
        Box(Modifier.fillMaxSize().padding(1.5.dp).border(1.dp, Color.Black, shape))
        Box(Modifier.fillMaxSize().border(1.5.dp, border, shape))
    }
}

/** A spec's talent-tab icon (spec §4.3): square, the pill radius, a 1 dp line. */
@Composable
fun SpecIcon(icon: String, size: Dp = 28.dp, source: IconSource? = LocalIconSource.current) {
    val shape = RoundedCornerShape(FsRadius.pill)
    Box(Modifier.size(size).clip(shape).background(FsColor.cardTop).border(1.dp, GamePalette.iconLine, shape)) {
        RemoteIcon(icon, source, Modifier.fillMaxSize())
    }
}
