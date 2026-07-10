// Memulai timer dengan label 'proses-loop'
console.time("proses-loop");

let total = 0;
for (let i = 0; i < 1000000; i++) {
  total += i;
}

// Menghentikan timer dan mencetak waktu ke console
console.timeEnd("proses-loop");
// Output di console: "proses-loop: 3.52ms"
