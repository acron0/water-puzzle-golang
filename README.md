water-puzzle-golang
===================

Water Puzzle (aka 'Water Logic') solver written in Go. The common implementation of this puzzle goes something like...

"Given two buckets - one holding 4 litres, one holding 7 litres - place exactly 2 litres of water into the 4 litre bucket."

###Usage

-target &lt;target mount, e.g. 2 litres (-target 2)&gt;

-buckets &lt;list of buckets. the first is always the target bucket (-buckets 4 7)&gt;

-moves &lt;the maximum number of moves (defaults to 10). some variantions of the puzzle specify a cap on the number of moves allowed&gt;

-results &lt;the maximum number of results to return&gt;

-quick &lt;returns the first result under the max moves count, so doesn''t necessarily find the least number of moves&gt;

###Examples

<pre>
$ go run *.go -target 3 -buckets 9 6 3

Water Puzzle Solver
----------
Puzzle: 3 in 9
Buckets: [9 6 3]
Moves Allowed: 10

Route:
=&gt; 01 [9l: 9, 6l: 0, 3l: 0] (Filled bucket 0)
=&gt; 02 [9l: 3, 6l: 6, 3l: 0] (Transferred bucket 0 to bucket 1)
Route:
=&gt; 01 [9l: 0, 6l: 0, 3l: 3] (Filled bucket 2)
=&gt; 02 [9l: 3, 6l: 0, 3l: 0] (Transferred bucket 2 to bucket 0)
</pre>

<pre>
$ go run *.go -target 3 -buckets 4 8 19 -quick

Water Puzzle Solver
----------
Puzzle: 3 in 4
Buckets: [4 8 19]
Moves Allowed: 10

Route:
=&gt; 01 [4l: 4, 8l: 0, 19l: 0] (Filled bucket 0)
=&gt; 02 [4l: 0, 8l: 4, 19l: 0] (Transferred bucket 0 to bucket 1)
=&gt; 03 [4l: 4, 8l: 4, 19l: 0] (Filled bucket 0)
=&gt; 04 [4l: 0, 8l: 8, 19l: 0] (Transferred bucket 0 to bucket 1)
=&gt; 05 [4l: 0, 8l: 0, 19l: 8] (Transferred bucket 1 to bucket 2)
=&gt; 06 [4l: 0, 8l: 0, 19l: 19] (Filled bucket 2)
=&gt; 07 [4l: 0, 8l: 8, 19l: 11] (Transferred bucket 2 to bucket 1)
=&gt; 08 [4l: 0, 8l: 0, 19l: 11] (Emptied bucket 1)
=&gt; 09 [4l: 0, 8l: 8, 19l: 3] (Transferred bucket 2 to bucket 1)
=&gt; 10 [4l: 3, 8l: 8, 19l: 0] (Transferred bucket 2 to bucket 0)
</pre>
