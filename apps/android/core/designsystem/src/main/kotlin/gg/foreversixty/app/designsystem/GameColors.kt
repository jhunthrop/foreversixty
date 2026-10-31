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

    // Spec-only values (spec §4.1, F2.6, F3.2); lane T folds them into tokens.css when the site gains them.
    /** A picker sheet's ground. */
    val sheet = Color(0xFF10141D)

    /** The tooltip's inner panel, and the 1 dp line around it and around a spec icon. */
    val tooltipPanel = Color(0xFF0A0D14)
    val iconLine = Color(0xFF3A4256)

    /** A chip or secondary control's 4.5% white fill, and the pressed 10% gold wash. */
    val controlFill = Color.White.copy(alpha = 0.045f)
    val pressedWash = FsColor.gold.copy(alpha = 0.10f)

    /** The viewer's own row (a roster's own line): gold at 5%. */
    val rowWash = FsColor.gold.copy(alpha = 0.05f)

    /** The offline strip's ground (spec §5.0). */
    val offlineGround = Color(0xFF11151F)

    /** The parse ladder's pink as text (spec W3 §4 Parse number); the other text steps are tokens. */
    val parsePinkText = Color(0xFFF08CC0)

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
