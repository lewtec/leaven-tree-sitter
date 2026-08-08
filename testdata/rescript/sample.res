// Single line
1

/*
 * Multiline
 */
2

/* /* Nested
  /* multiple */ times
  // */ */
3

4 // post single-line

5 /* post multi-line */

/* / */
6

/* */ 7

/** **/ 8

/* comment //*/ */*/ 9

foo
  // in-pipe
  ->bar

switch foo {
| 1 => 1
// in-switch
}

switch foo {
| 1 => 1
/* block comment in switch */
| 2 => 2
}

switch l {
| "a" => letterAll
/* | "b" => [||] */
/* | "g" => [||] */
/* | "h" => [||] */
/* | "i" => [||] */
}
