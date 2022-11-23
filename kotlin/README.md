## Kotlin WebAssembly (WASM)

### How-to
- Download last release and unpack from https://github.com/JetBrains/kotlin/releases/tag/v1.7.21 or install with Homebrew: https://formulae.brew.sh/cask/kotlin-native
- Create file **hello.kt**:
```kotlin
fun main() {
	println("Hello World!")
}
```
- Execute `kotlinc-native -list-targets` and find **wasm32** in the list, this will be our target.
- Execute `kotlinc-native hello.kt -target wasm32 -o hello` and the output will contain a **hello.wasm** and a **hello.wasm.js**.
  - On Apple Macbook you might see: **libcallbacks.dylib**: App Can't Be Opened Because Apple Cannot Check It.
  - Watch this video: https://www.youtube.com/watch?v=T7LXprbqm3E
  - In directory: `/kotlin/konan/nativelib/` open all files `*.dylib` in terminal by Allowing each file in System Preference -> Security & Privacy: Allow apps downloaded from
- Create **index.html**
```html
<!DOCTYPE html>
<meta charset="utf-8">
<title>WebAssembly Kotlin Demo</title>
<script wasm="hello.wasm" src="hello.wasm.js"></script>
```
- To see result you need start http server in current directory. For example(Rust): `cargo install https` and then start: `http`
- This how-to from: https://blog.jdriven.com/2021/04/running-kotlin-in-the-browser-with-wasm/

### Do not recommend use Kotlin for WASM yet 
- I didn't find how to call kotlin function from JS. https://stackoverflow.com/questions/54138204/how-to-execute-kotlin-webassembly-function-from-javascript
- Seems Kotlin creators doesn't pay attention to WASM. https://discuss.kotlinlang.org/t/wasm-wasi-where-do-you-go-kotlin/12175 (**darksnake** comments):
- Still, I do not really think that WASM or WASI will have any significant impact anytime in near future.
  - Whatever Mozilla does, it is still almost impossible to reach the level of sophistication and duplicate the effort put into JDK/JVM.
  - Kotlin relies a lot on Java libraries and that is why it have grown so fast. There are no such libraries for WASM, so it will be years before one can write with similar level of comfort.
  - If you are searching for a platform that could eat and run anything and give polyglot interop functionality, there is GraalVM. It is also not yet production ready, but if it works as it supposed to do (and there are some indications that it will), it gives the same functionality with addition of JVM languages and will be ready much faster.
- There is no such thing as a “speed of C++”. There are fast programs in C and there are very slow ones. People usually compare average JVM program with optimized C code and do wrong conclusions. JVM engineers invested a lot of effort to optimize average code, so average code in JVM is super-fast. Even without tools in the language to make specialized optimizations. Part of that speed is achieved by JIT smart optimization. WASM is much more low-level and is designed to work without runtime deoptimizations. It means that it relies a lot on compile-time optimizations done by LLVM. In order to make the most from LLVM one need to optimize the LLVM IR beforehand. And even in the most optimized case, I still think, that average code in JVM will be much faster.

### Maybe in future
- https://blog.jetbrains.com/kotlin/2021/11/k2-compiler-kotlin-wasm-and-tooling-announcements-at-the-2021-kotlin-event/

### Reference
- https://kotlinlang.org/docs/native-command-line-compiler.html#compile-the-code-from-the-console
