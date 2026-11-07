package gg.foreversixty.app.testing

import com.github.takahirom.roborazzi.RoborazziOptions

/**
 * Every screenshot golden compares exactly (design/parity.md "Native", spec section 7): a changed
 * pixel fails, so copy drift cannot pass silently. Goldens are recorded on CI (android.yml
 * record), so the local and CI renderers never meet in one compare.
 */
val STRICT_SNAPSHOT = RoborazziOptions(
    compareOptions = RoborazziOptions.CompareOptions(changeThreshold = 0F),
)
