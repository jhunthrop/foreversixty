package gg.foreversixty.app.designsystem

import androidx.compose.ui.graphics.Color

/** The class colours by the API's class word (`warrior`, `shaman`, ...). */
fun classColor(className: String?): Color = when (className?.lowercase()) {
    "warrior" -> FsColor.classWarrior
    "paladin" -> FsColor.classPaladin
    "hunter" -> FsColor.classHunter
    "rogue" -> FsColor.classRogue
    "priest" -> FsColor.classPriest
    "shaman" -> FsColor.classShaman
    "mage" -> FsColor.classMage
    "warlock" -> FsColor.classWarlock
    "druid" -> FsColor.classDruid
    else -> FsColor.text
}

/**
 * The game's colours beyond the class colours (item quality lives with the items, `ItemQuality`),
 * plus the spec-only values with no site token yet. Screens pass the data's own key and never a hex.
 */
object GameColors {
    fun faction(slug: String): Color = if (slug == "horde") FsColor.horde else FsColor.alliance

    // The surfaces the tokens name by another role, in round 5's palette (lane M0).
    /** A picker sheet's ground: the panel brown-black. */
    val sheet = FsColor.raised

    /** The tooltip's inner panel (the client's navy), and the 1 dp line around it and around a spec icon (their rule). */
    val tooltipPanel = FsColor.tooltip
    val iconLine = FsColor.rule

    /** A chip or secondary control's 4.5% white fill, and the pressed 10% gold wash. */
    val controlFill = Color.White.copy(alpha = 0.045f)
    val pressedWash = FsColor.gold.copy(alpha = 0.10f)

    /** The viewer's own row (a roster's own line): gold at 5%. */
    val rowWash = FsColor.gold.copy(alpha = 0.05f)

    /** The offline strip's ground: the deepest brown-black well. */
    val offlineGround = FsColor.bg

    /** A parse percentile's colour (`--color-parse-*`). */
    fun parse(percentile: Int): Color = when {
        percentile >= 100 -> FsColor.parseGold
        percentile == 99 -> FsColor.parsePink
        percentile >= 95 -> FsColor.parseOrange
        percentile >= 75 -> FsColor.parsePurple
        percentile >= 50 -> FsColor.parseBlue
        percentile >= 25 -> FsColor.parseGreen
        else -> FsColor.parseGreyText
    }
}
