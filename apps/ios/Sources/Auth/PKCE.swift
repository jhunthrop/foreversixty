import CryptoKit
import Foundation
import Security

/// A PKCE verifier and its S256 challenge (RFC 7636).
public struct PKCE: Sendable, Equatable {
    public static let method = "S256"
    /// 32 random bytes encode to the 43-character verifier, RFC 7636's minimum length.
    static let verifierByteCount = 32

    public let verifier: String
    public let challenge: String

    public init(verifier: String) {
        self.verifier = verifier
        self.challenge = Self.challenge(for: verifier)
    }

    /// A fresh pair from the system's secure random source.
    public static func generate(random: any RandomBytes = SystemRandomBytes()) throws(AuthError) -> PKCE {
        PKCE(verifier: try random.bytes(verifierByteCount).base64URLEncoded())
    }

    public static func challenge(for verifier: String) -> String {
        Data(SHA256.hash(data: Data(verifier.utf8))).base64URLEncoded()
    }
}

/// Secure random bytes, injectable so tests are deterministic.
public protocol RandomBytes: Sendable {
    func bytes(_ count: Int) throws(AuthError) -> Data
}

public struct SystemRandomBytes: RandomBytes {
    public init() {}

    public func bytes(_ count: Int) throws(AuthError) -> Data {
        var buffer = [UInt8](repeating: 0, count: count)
        let status = SecRandomCopyBytes(kSecRandomDefault, count, &buffer)
        guard status == errSecSuccess else { throw .randomUnavailable }
        return Data(buffer)
    }
}

extension Data {
    /// Base64url without padding (RFC 4648 §5), as PKCE and OAuth state use.
    func base64URLEncoded() -> String {
        base64EncodedString()
            .replacingOccurrences(of: "+", with: "-")
            .replacingOccurrences(of: "/", with: "_")
            .replacingOccurrences(of: "=", with: "")
    }
}
