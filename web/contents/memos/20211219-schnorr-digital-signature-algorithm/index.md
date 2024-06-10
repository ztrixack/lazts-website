---
title: Schnorr Digital Signature Algorithm
slug: 20211219-schorr-digital-signature-algorithm
excerpt:
featured_image:
published_at: 2021-12-19
last_updated_at: 2024-05-20
published: false
tags:
  - cryptography
---

## Digital Signature​ คืออะไร

- Use elliptic curve cryptography (ECC). (like ECDSA)​
- Was covered by U.S. Patent 4995082 since 1990 (expired in February 2008)​
- It is efficient and short signatures more than ECDSA​
- Multi-signatures feature​
- Advantages over ECDSA ​
  - computational efficiency​
  - Storage​
  - privacy​

##

- Verification is very close in performance to ECDSA.​
- Schnorr is linear but ECDSA is not; (more elegant and simpler).​
  - If we add 2 signatures together, the result is valid too! ​
- Public key recovery: reconstruct a public key from a signature and message.​

- Better privacy, by making different multisig spending policies indistinguishable on-chain.​
- Enabling simpler higher-level protocols, such as atomic swaps that are indistinguishable from normal payments. These can be used to build more efficient payment channel constructions.​
- Improving verification speed, by supporting batch validation of all signatures at a once.​
- Switching to a provably secure construction, perhaps preventing an exploit against ECDSA in the future.​

## Reference
