plugins {
    alias(libs.plugins.forever.android.library)
}

// Crash-report scrubbing and the on/off rule, with no SDK: the app links Sentry and calls into
// these, so the privacy rules are tested without it. iOS mirrors them in Sources/Telemetry.
dependencies {
    testImplementation(libs.junit)
}
