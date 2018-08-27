<?php
/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

fscanf(STDIN, "%d",
    $N // Number of elements which make up the association table.
);
fscanf(STDIN, "%d",
    $Q // Number Q of file names to be analyzed.
);
for ($i = 0; $i < $N; $i++)
{
    fscanf(STDIN, "%s %s",
        $EXT[], // file extension
        $MT[] // MIME type.
    );
}
foreach ($EXT as $k=>$v) {
    $EXT[$k] = strtolower($v);
}
for ($i = 0; $i < $Q; $i++)
{
    $FNAME = stream_get_line(STDIN, 500 + 1, "\n"); // One file name per line.
    if (strpbrk($FNAME, ".") === FALSE) {
        echo("UNKNOWN\n");
    } else {
        $fileExt = strtolower(array_reverse(explode(".", $FNAME))[0]);
        $k = array_search($fileExt, $EXT);
        if ($k === FALSE) {
            echo("UNKNOWN\n");
        } else {
            echo($MT[$k]."\n");
        }
    }
}

// Write an action using echo(). DON'T FORGET THE TRAILING \n
// To debug (equivalent to var_dump): error_log(var_export($var, true));

// For each of the Q filenames, display on a line the corresponding MIME type. If there is no corresponding type, then display UNKNOWN.

?>
