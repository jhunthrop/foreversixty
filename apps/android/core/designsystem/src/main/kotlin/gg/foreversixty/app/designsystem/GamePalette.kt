package gg.foreversixty.app.designsystem

import androidx.compose.ui.graphics.Color

/**
 * The game's colours over the generated tokens: item quality (spec §4.1, §4.3), class and
 * faction, plus the spec-only values with no site token yet. Screens pass the data's own key
 * (a quality 0 to 5, a class slug) and never a hex.
 */
object GamePalette {
    /** The quality border: the game's own colour. [quality] is 0 (poor) to 5 (legendary). */
    fun qualityBorder(quality: Int): Color = when (quality) {
        0 -> FsColor.rarityPoor
        1 -> FsColor.rarityCommon
        2 -> FsColor.rarityUncommon
        3 -> FsColor.rarityRare
        4 -> FsColor.rarityEpic
        else -> FsColor.rarityLegendary
    }

    /** An item name: the site's AA variants for rare and epic on the dark ground. */
    fun qualityText(quality: Int): Color = when (quality) {
        3 -> FsColor.rarityRareText
        4 -> FsColor.rarityEpicText
        else -> qualityBorder(quality)
    }

    /** A class's text colour by slug (`shaman` is the site's AA blue); muted when unknown. */
    fun classColor(slug: String): Color = when (slug) {
        "warrior" -> FsColor.classWarrior
        "paladin" -> FsColor.classPaladin
        "hunter" -> FsColor.classHunter
        "rogue" -> FsColor.classRogue
        "priest" -> FsColor.classPriest
        "shaman" -> FsColor.classShaman
        "mage" -> FsColor.classMage
        "warlock" -> FsColor.classWarlock
        "druid" -> FsColor.classDruid
        else -> FsColor.muted
    }

    fun faction(slug: String): Color = if (slug == "horde") FsColor.horde else FsColor.alliance

    // Spec-only values (spec §4.1, §5.0, F2.6); lane T folds them into tokens.css when the site gains them.
    val offlineGround = Color(0xFF11151F)
    val sheet = Color(0xFF10141D)
    val tooltipPanel = Color(0xFF0A0D14)
    val iconLine = Color(0xFF3A4256)

    /** The game's green for Equip and Use lines, and its gold for Item Level. */
    val gameGreen = FsColor.parseGreen
    val itemLevel = FsColor.parseGold

    /** A chip or secondary control's 4.5% white fill, and the pressed 10% gold wash. */
    val controlFill = Color.White.copy(alpha = 0.045f)
    val pressedWash = FsColor.gold.copy(alpha = 0.10f)
}
