use wasm_bindgen::prelude::*;
use web_sys::console;

#[wasm_bindgen(js_name = rustCalcPi)]
pub fn calc_pi(js_rounds: &str) -> f64 {
    fn calc(rounds: i64) -> f64 {
        let pi = (2..rounds).fold(1.0, |pi, i| {
            let x = -1.0f64 + (2.0 * (i & 0x1) as f64);
            pi + x / (2 * i - 1) as f64
        }) * 4.0;

        console::log_1(&JsValue::from_f64(pi));

        return pi;
    }

    return js_rounds.parse::<i64>().map_or(-1., calc);
}
