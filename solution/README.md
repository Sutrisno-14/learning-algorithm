# SOAL 
[
Write a function :
function solution($A);
that, given an array A of N integers, returns the smallest
positive integer (greater than 0) that does not occur in A.
For example, 
- given A = [1, 3, 6, 4, 1, 2], the function should
return 5.
- Given A = [1, 2, 3], the function should return 4.
- Given A = [−1, −3], the function should return 1.

Write an efficient algorithm for the following assumptions:
- N is an integer within the range [1..100,000];
- each element of array A is an integer within the range [−1,000,000..1,000,000
]

## Untuk membuat function solution(A) ini kita membutuhkan sebuah perulangan untuk menyelesaikannya. Perulangan ini digunakan untuk mencari angka yang paling kecil dan tidak ada di dalam array A. Kita akan membandingkan setiap elemen yang ada didalam array A dengan variabel result. Berikut algortima untuk menyelesaikan nya:
- (1). Mengurutkan array A secara ascanding dengan function sort.Ints(A) -> dari yang terkecil hingga ke yang terbesar.
- (2). Membuat variabel result := 1 (inisialisasi 1 karena angka yang paling kecil yang mungkin tidak ada di dalam array A/bilangan bulat positif (lebih besar dari 0) yang tidak terjadi di A).
- (3). Melakukan perulangan dari i := 0 sampai len(A).
- (4). Jika A[i] > result, maka akan keluar dari perulangan dan menandakan perulangan selesai.
- (5). Jika A[i] == result,berarti angka yang di cari sudah ada di array dan nilai result akan bertambah 1, yang menjadikan result sebagai angka yang lebih besar dari angka result sebelumnya.

## Kesimpulan
- Dari task yang diberikan kita dapat menyelesaikannya dengan menggunakan perulangan yang mana perulangan ini bertujuan untuk menemukan angka yang paling kecil bilangan bulat positif (lebih besar dari 0) yang tidak ada di dalam array A.