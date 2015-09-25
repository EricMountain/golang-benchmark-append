set term png large size 1800,1100
set ylabel "ns/op" rotate by 90
set xlabel "len(slice)"

set output "polo-vs-append_struct.png"
set title "Index vs append with structs"
plot "plot.dat" using 1:4 with lines title "Index", "plot.dat" using 1:5 with lines title "Append"

set output "polo-vs-append_int.png"
set title "Index vs append with integers"
plot "plot.dat" using 1:2 with lines title "Index", "plot.dat" using 1:3 with lines title "Append"
