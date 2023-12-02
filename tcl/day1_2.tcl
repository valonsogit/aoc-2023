set file [open "input/day1.txt" r]
set input [read -nonewline $file]
set lines [split $input "\n"]

proc parse {lines} {
    set possibleValues [list "one" "two" "three" "four" "five" "six" "seven" "eight" "nine" "1" "2" "3" "4" "5" "6" "7" "8" "9"]
    set i 1
    set total 0
    foreach line $lines {
        set firstIdx 9999999
        set lastIdx -1
        set first 0
        set last 0
        set currentN 0
        foreach n $possibleValues {
            incr currentN
            set fidx [string first $n $line]
            set lidx [string last $n $line]

            if { $fidx == -1 } {
                continue
            }
            if { $fidx < $firstIdx } {
                if { $currentN < 10} {
                    set first $currentN
                } else {
                    set first [ expr {int($n)} ]
                }
                set firstIdx $fidx

            }
            if { $lidx > $lastIdx } {
                if { $currentN < 10} {
                    set last $currentN
                } else {
                    set last [ expr {int($n)} ]
                }
                set lastIdx $lidx
            }
        }
        if { $last < 0 } {
            set last $first
        }
        set calibration [expr {int("$first$last")}]
        puts "$calibration"
        incr total $calibration
        incr i
    }
    return $total
}
puts [parse $lines]