# chainbridge-utils

This repository is fork from [ChainSafe/chainbridge-utils](https://github.com/ChainSafe/chainbridge-utils)

## Requirements

- [Rust](https://www.rust-lang.org/tools/install)
- [TON-SDK](https://github.com/tonlabs/TON-SDK) - important version `1.1.0` and compile it via `cargo build --release`

One need to specify compiled DLL directory path:

```sh
export CGO_LDFLAGS="-L//${TON_SDK_PATH}/target/release/deps/ -lton_client"
```