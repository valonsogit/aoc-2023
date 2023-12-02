set file [open "input/day1.txt" r]
set input [read -nonewline $file]
set lines [split $input "\n"]
proc parse {lines} {
    set i 0
    set total 0
    foreach line $lines {
        set first ""
        set last ""
        foreach char [split $line ""] {
            if { [string is integer $char] } {
                if { ![string length $first] } {
                    set first $char
                } else {
                    set last $char
                }
            }
        }
        if { ![string length $last] } {
            set last $first

        }
        set calibration [expr {int("$first$last")}]
        puts $i
        incr total $calibration
        incr i
    }
    return $total
}
puts "______"
puts [parse $lines]