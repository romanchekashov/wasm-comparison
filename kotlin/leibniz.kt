fun main() {
    val rounds: Int = 100000000

    var pi = 1.0
    var x = 1.0

    for (i in 2 until rounds + 2) {
        x *= -1.0
        pi += x / (2 * i - 1)
    }

    pi *= 4.0

    println(pi)
}
