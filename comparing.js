var Jimp = require("jimp");
var convert = require("color-convert");
var DeltaE = require("delta-e");
const R1 = 255;
const G1 = 0;
const B1 = 0;

const R2 = 184;
const G2 = 0;
const B2 = 0;

var lab1 = convert.rgb.lab.raw(R1, G1, B1);
var lab2 = convert.rgb.lab.raw(R2, G2, B2);
var color1 = { L: lab1[0], A: lab1[1], B: lab1[2] };
var color2 = { L: lab2[0], A: lab2[1], B: lab2[2] };

console.log(DeltaE.getDeltaE00(color1, color2));
