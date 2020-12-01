# goc

goc converts numbers between bases.

this naively wraps golang's strconv for my convenience.

# usage

`./goc [-f fromBase] [-t toBase] [-p] <number>`

# notes

this tool defaults to converting hex to decimal. it ignores prefixes like `0x` or `0b` if they're included in input.

the `-p` flag adds base-specific prefixes to the output for base 16 and base 2.

supports bases 2 through 36, and values up to `64_BIT_INT_MAX`

# building

`go build .`
