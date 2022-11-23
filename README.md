## WebAssembly (WASM)

### leibniz calc PI number: rounds = 100000000
- **JS**: Call to doSomething took **115.6** milliseconds.
- **Rust**: Call to doSomething took **223.5** milliseconds.
- **Go**: Call to doSomething took **351.1** milliseconds.

### Language comparison: better chose Rust
#### [Rust WebAssembly (WASM)](rust/README.md): recommended
#### [Go WebAssembly (WASM)](go/README.md): less performant than Rust
#### [Kotlin WebAssembly (WASM)](kotlin/README.md): not recommended yet

### WASM isn't necessarily faster than JS
- [WASM isn't necessarily faster than JS](https://www.reddit.com/r/webdev/comments/uj8ivc/wasm_isnt_necessarily_faster_than_js/)
- [Zaplib post-mortem](https://zaplib.com/docs/blog_post_mortem.html?ck_subscriber_id=1715213923)
- [WebAssembly vs Javascript](https://ianjk.com/webassembly-vs-javascript/?ck_subscriber_id=1715213923)
- [Typescript as fast as Rust: Typescript++](https://zaplib.com/docs/blog_ts++.html)

### Reference
- https://github.com/mbasso/awesome-wasm
- [When to use Rust and when to use Go](https://blog.logrocket.com/when-to-use-rust-when-to-use-golang/)
- https://github.com/niklas-heer/speed-comparison
