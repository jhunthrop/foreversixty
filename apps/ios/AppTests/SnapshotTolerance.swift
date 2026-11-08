/// Every image golden's tolerance (design/parity.md "Native", spec section 7): every pixel must
/// match within the perceptual threshold, so copy drift cannot pass silently (a signed-out
/// snapshot passed with changed text on 2026-11-07). Goldens are recorded on CI (ios.yml record).
enum SnapshotTolerance {
    static let precision: Float = 1
    static let perceptualPrecision: Float = 0.98
}
