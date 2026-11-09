import UIKit

/// Every image golden's tolerance (design/parity.md "Native", spec section 7): every pixel must
/// match within the perceptual threshold, so copy drift cannot pass silently (a signed-out
/// snapshot passed with changed text on 2026-11-07). Goldens are recorded on CI (ios.yml record).
enum SnapshotTolerance {
    static let precision: Float = 1
    static let perceptualPrecision: Float = 0.98

    /// The same threshold as a per-channel byte delta, for the key-window snapshots compared on the
    /// CPU (KeyWindowSnapshot.swift): (1 - 0.98) of 255 is 5 levels. Anti-aliasing noise between
    /// two GPUs is a level or three; a changed glyph or colour moves far more.
    static let maxChannelDelta = Int(((1 - perceptualPrecision) * 255).rounded(.down))

    /// Whether every channel of every pixel of the two images is within maxChannelDelta. Images of
    /// different pixel sizes never are.
    static func withinChannelDelta(_ lhs: UIImage, _ rhs: UIImage) -> Bool {
        guard let a = rgba(lhs), let b = rgba(rhs), a.width == b.width, a.height == b.height else { return false }
        // A plain loop over the raw buffers: a test target builds without optimisation, and a
        // full-screen image is twelve million bytes.
        return a.bytes.withUnsafeBufferPointer { left in
            b.bytes.withUnsafeBufferPointer { right in
                for i in 0..<left.count where abs(Int(left[i]) - Int(right[i])) > maxChannelDelta {
                    return false
                }
                return true
            }
        }
    }

    /// The image's pixels as 8-bit RGBA, drawn into one known format so two images compare byte for byte.
    private static func rgba(_ image: UIImage) -> (bytes: [UInt8], width: Int, height: Int)? {
        guard let cg = image.cgImage else { return nil }
        let width = cg.width, height = cg.height
        var bytes = [UInt8](repeating: 0, count: width * height * 4)
        let drawn = bytes.withUnsafeMutableBytes { buffer -> Bool in
            guard let context = CGContext(
                data: buffer.baseAddress, width: width, height: height, bitsPerComponent: 8, bytesPerRow: width * 4,
                space: CGColorSpaceCreateDeviceRGB(), bitmapInfo: CGImageAlphaInfo.premultipliedLast.rawValue
            ) else { return false }
            context.draw(cg, in: CGRect(x: 0, y: 0, width: width, height: height))
            return true
        }
        return drawn ? (bytes, width, height) : nil
    }
}
