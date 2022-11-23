function jsCalcPi(rounds) {
    let x = 1.0;
    let pi = 1.0;

    for (let i = 2; i < rounds + 2; i++) {
        x *= -1;
        pi += x / (2 * i - 1);
    }

    pi *= 4;

    console.log(pi);

    return pi;
}

function arrayManipulations(num) {
    const run = (num) => {
        let arr = new Array(num).fill('a');

        for (let i = 0; i < num; i++) {
            arr = arr.map(x => x + '1');
        }

        while (arr.length > 0) {
            arr.splice(0, 1);
        }
    }

    let count = num;

    while (count-- > 0) {
        run(num);
    }
}
