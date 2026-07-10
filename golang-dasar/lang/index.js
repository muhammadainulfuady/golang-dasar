let array = [2, 3, 4, 1, 4];

jumlah = 0;
panjangArray = array.length;

array.forEach(function (e) {
  jumlah += e;
});

rataRata = jumlah / panjangArray;

console.log(jumlah);
console.log(rataRata);
