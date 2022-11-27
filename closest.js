var Jimp = require("jimp");
var convert = require("color-convert");
var DeltaE = require("delta-e");
const R = 75;
const G = 129;
const B = 145;

Jimp.read("gradient.jpg", function (err, image) {
  i = 0;
  var rgbarray = [];
  while (i < 200) {
    var rgb = JSON.stringify(Jimp.intToRGBA(image.getPixelColor(i, 0)))
      .slice(1, -1)
      .split(",");

    rgb = rgb.map((s) => s.slice(4));
    rgbarray.push(rgb);
    i++;
  }
  console.log(rgbarray);
  var lab1 = convert.rgb.lab.raw(R, G, B);
  var color1 = { L: lab1[0], A: lab1[1], B: lab1[2] };
  console.log(color1);

  var color2;
  var lowestD = 0;
  var pos = 0;
  var closestcol;
  for (let i = 0; i < rgbarray.length; i++) {
    tempcol = rgbarray[i][0] + ", " + rgbarray[i][1] + ", " + rgbarray[i][2];
    lab2 = convert.rgb.lab.raw(
      parseInt(rgbarray[i][0]),
      parseInt(rgbarray[i][1]),
      parseInt(rgbarray[i][2])
    );
    var color2 = { L: lab2[0], A: lab2[1], B: lab2[2] };
    if (DeltaE.getDeltaE00(color1, color2) < lowestD || lowestD == 0) {
      lowestD = DeltaE.getDeltaE00(color1, color2);
      closestcol = tempcol;
      pos = i;
    }
  }
  console.log(
    "The closest color was: " +
      closestcol +
      " which is in position " +
      (pos + 1) +
      " with a delta e of: " +
      lowestD +
      "%"
  );
});
